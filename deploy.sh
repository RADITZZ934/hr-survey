#!/bin/bash

# Configuration - sesuaikan path sesuai kebutuhan server
APP_DIR="/root/hrd"
BACKEND_TARGET_DIR="$APP_DIR/backend"
FRONTEND_TARGET_DIR="/var/www/hr-survey" # Path frontend
BACKUP_DIR="$APP_DIR/backups"
LOG_DIR="$APP_DIR/logs"

# PM2 App Names & Ports
FRONTEND_APP_NAME="hr-survey-frontend"
BACKEND_APP_NAME="hr-survey-backend"
PORT_FRONTEND=1301
PORT_BACKEND=1302

# Print with styling
log_info() {
  echo -e "\e[34m[INFO]\e[0m $1"
}

log_success() {
  echo -e "\e[32m[SUCCESS]\e[0m $1"
}

log_error() {
  echo -e "\e[31m[ERROR]\e[0m $1"
}

# Input package validation
PACKAGE_PATH=$1
if [ -z "$PACKAGE_PATH" ]; then
  log_error "Please provide the path to the release package (e.g., ./deploy.sh ~/hrd/tmp/release.tar.gz)"
  exit 1
fi

if [ ! -f "$PACKAGE_PATH" ]; then
  log_error "Package file not found at: $PACKAGE_PATH"
  exit 1
fi

# Cek apakah PM2 terinstall
if ! command -v pm2 &> /dev/null; then
  log_error "PM2 is not installed. Install it first using: npm install -g pm2"
  exit 1
fi

log_info "Starting deployment using package: $PACKAGE_PATH"

# Create directories if they don't exist
mkdir -p "$BACKEND_TARGET_DIR"
mkdir -p "$FRONTEND_TARGET_DIR"
mkdir -p "$BACKUP_DIR"
mkdir -p "$LOG_DIR"
mkdir -p "$APP_DIR/tmp_extract"

# Cleanup temporary extraction folder
cleanup() {
  rm -rf "$APP_DIR/tmp_extract"
}
trap cleanup EXIT

# Extracting package
log_info "Extracting build package..."
tar -xzf "$PACKAGE_PATH" -C "$APP_DIR/tmp_extract"
if [ $? -ne 0 ]; then
  log_error "Failed to extract package."
  exit 1
fi

# Backup current release (Opsional)
TIMESTAMP=$(date +%Y%m%d_%H%M%S)
log_info "Creating backup of current deployment..."
tar -czf "$BACKUP_DIR/backup_$TIMESTAMP.tar.gz" -C "$APP_DIR" backend 2>/dev/null || true

# ---------------------------------------------------------
# Deploy Backend
# ---------------------------------------------------------
log_info "Deploying Backend..."
if [ -d "$APP_DIR/tmp_extract/backend" ]; then
  cp -r "$APP_DIR/tmp_extract/backend/." "$BACKEND_TARGET_DIR/"
  
  # Restart / Start Backend dengan PM2
  cd "$BACKEND_TARGET_DIR" || exit 1
  
  # Jalankan binary Go via PM2
  PORT=$PORT_BACKEND pm2 restart "$BACKEND_APP_NAME" --update-env 2>/dev/null || \
  PORT=$PORT_BACKEND pm2 start ./hrd-backend --name "$BACKEND_APP_NAME"
  
  log_success "Backend deployed and running on port $PORT_BACKEND"
else
  log_info "No backend folder found in package, skipping backend update."
fi

# ---------------------------------------------------------
# Deploy Frontend
# ---------------------------------------------------------
log_info "Deploying Frontend..."
if [ -d "$APP_DIR/tmp_extract/frontend" ]; then
  cp -r "$APP_DIR/tmp_extract/frontend/." "$FRONTEND_TARGET_DIR/"
  
  cd "$FRONTEND_TARGET_DIR" || exit 1

  # Menyajikan file statis (HTML/JS build hasil 'npm run build') via PM2 serve
  pm2 restart "$FRONTEND_APP_NAME" --update-env 2>/dev/null || \
  pm2 serve "$FRONTEND_TARGET_DIR" $PORT_FRONTEND --name "$FRONTEND_APP_NAME" --spa

  log_success "Frontend deployed and running on port $PORT_FRONTEND"
else
  log_info "No frontend folder found in package, skipping frontend update."
fi

# Save PM2 process list
pm2 save

log_success "Deployment completed successfully!"
