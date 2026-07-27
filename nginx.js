const fs = require('fs');
const path = require('path');

// Target nama file konfigurasi
const fileName = 'hrd-survey.conf';
const outputPath = path.join(__dirname, fileName);

// Isi Konfigurasi Nginx
const nginxConfig = `# Config dikodekan secara otomatis via nginx.js
server {
    listen 1400;
    server_name _;

    # Logging
    access_log /var/log/nginx/hrd_survey_access.log;
    error_log /var/log/nginx/hrd_survey_error.log;

    # Dynamic Proxy Headers
    proxy_http_version 1.1;
    proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection 'upgrade';
    proxy_set_header Host $host;
    proxy_cache_bypass $http_upgrade;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header X-Forwarded-Proto $scheme;

    # Backend API Routing (/api/ atau /api)
    location /api {
        proxy_pass http://127.0.0.1:1402/;
    }

    # Frontend App Routing (Default /)
    location / {
        proxy_pass http://127.0.0.1:1401/;
    }
}
`;

// Tulis file konfigurasi
try {
  fs.writeFileSync(outputPath, nginxConfig, 'utf8');
  console.log(`\x1b[32m[SUCCESS]\x1b[0m File konfigurasi '${fileName}' berhasil dibuat di:`);
  console.log(`          ${outputPath}`);
  console.log('\n\x1b[34m[INFO]\x1b[0m Port Mapping Baru:');
  console.log('       - Nginx Listen    : 1400');
  console.log('       - Frontend Target : 1401');
  console.log('       - Backend Target  : 1402');
} catch (err) {
  console.error(`\x1b[31m[ERROR]\x1b[0m Gagal membuat file konfigurasi:`, err.message);
}
