<template>
  <div class="h-screen w-screen bg-[#332941] flex font-sans antialiased overflow-hidden select-none">
    
    <!-- 1. Dark Left Vertical Sidebar (#332941) -->
    <aside class="w-20 lg:w-24 bg-[#332941] flex flex-col items-center py-7 justify-between flex-shrink-0 z-30 h-full">
        
        <!-- Top: Cream Logo Pill Badge -->
        <router-link to="/dashboard" class="group">
          <div class="w-12 h-12 rounded-2xl bg-[#f6f2e9] flex items-center justify-center shadow-md group-hover:scale-105 transition-all duration-300">
            <img src="/hrsurvey-icon.png" class="w-8 h-8 object-contain" alt="HR Survey Logo" />
          </div>
        </router-link>

        <!-- Center: Vertical Nav Icons -->
        <nav class="flex flex-col space-y-6 items-center">
          
          <!-- Nav 1: Dashboard -->
          <router-link 
            to="/dashboard"
            class="relative p-3 rounded-2xl transition-all duration-200 group"
            :class="isRouteActive('/dashboard') ? 'bg-white/20 text-white shadow-inner' : 'text-slate-300 hover:text-white hover:bg-white/10'"
            title="Overview"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z" />
            </svg>
            <span v-if="isRouteActive('/dashboard')" class="absolute -right-1 top-1/2 -translate-y-1/2 w-1.5 h-5 bg-amber-300 rounded-l-full"></span>
          </router-link>

          <!-- Nav 2: Surveys -->
          <router-link 
            to="/dashboard/surveys"
            class="relative p-3 rounded-2xl transition-all duration-200 group"
            :class="isRouteActive('/dashboard/surveys') ? 'bg-white/20 text-white shadow-inner' : 'text-slate-300 hover:text-white hover:bg-white/10'"
            title="Surveys"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
            </svg>
            <span v-if="isRouteActive('/dashboard/surveys')" class="absolute -right-1 top-1/2 -translate-y-1/2 w-1.5 h-5 bg-amber-300 rounded-l-full"></span>
          </router-link>

          <!-- Nav 3: Reports -->
          <router-link 
            to="/dashboard/reports"
            class="relative p-3 rounded-2xl transition-all duration-200 group"
            :class="isRouteActive('/dashboard/reports') ? 'bg-white/20 text-white shadow-inner' : 'text-slate-300 hover:text-white hover:bg-white/10'"
            title="Reports"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
            </svg>
            <span v-if="isRouteActive('/dashboard/reports')" class="absolute -right-1 top-1/2 -translate-y-1/2 w-1.5 h-5 bg-amber-300 rounded-l-full"></span>
          </router-link>

          <!-- Nav 4: Action Plans -->
          <router-link 
            to="/dashboard/action-plans"
            class="relative p-3 rounded-2xl transition-all duration-200 group"
            :class="isRouteActive('/dashboard/action-plans') ? 'bg-white/20 text-white shadow-inner' : 'text-slate-300 hover:text-white hover:bg-white/10'"
            title="Action Plans"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
            </svg>
            <span v-if="isRouteActive('/dashboard/action-plans')" class="absolute -right-1 top-1/2 -translate-y-1/2 w-1.5 h-5 bg-amber-300 rounded-l-full"></span>
          </router-link>

          <!-- Nav 5: Employees -->
          <router-link 
            to="/dashboard/employees"
            class="relative p-3 rounded-2xl transition-all duration-200 group"
            :class="isRouteActive('/dashboard/employees') ? 'bg-white/20 text-white shadow-inner' : 'text-slate-300 hover:text-white hover:bg-white/10'"
            title="Manajemen Karyawan"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a3 3 0 11-6 0 3 3 0 016 0z" />
            </svg>
            <span v-if="isRouteActive('/dashboard/employees')" class="absolute -right-1 top-1/2 -translate-y-1/2 w-1.5 h-5 bg-amber-300 rounded-l-full"></span>
          </router-link>

        </nav>

        <!-- Bottom: Logout Action Icon -->
        <button 
          @click="handleLogout"
          class="p-3 rounded-2xl text-slate-300 hover:text-rose-300 hover:bg-white/10 transition-all duration-200"
          title="Logout"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
          </svg>
        </button>

      </aside>

      <!-- 2. Right Canvas Area (#332941 Background) -->
      <div class="flex-1 flex flex-col min-w-0 bg-[#332941] h-full overflow-hidden">
        
        <!-- Top Right Header (#332941 Background) -->
        <header class="px-6 sm:px-10 py-5 flex items-center justify-between bg-[#332941] text-white z-20">
          
          <!-- Page Title -->
          <div class="flex items-center space-x-3">
            <h1 class="text-2xl sm:text-3xl font-extrabold text-white tracking-tight">
              {{ pageTitle }}
            </h1>
          </div>

          <!-- Header Right Utilities (Search, Notification, Profile Pill) -->
          <div class="flex items-center space-x-4">
            
            <!-- Search Icon Input -->
            <div class="relative hidden sm:block w-48 lg:w-64">
              <span class="absolute inset-y-0 left-0 pl-3.5 flex items-center pointer-events-none text-slate-300">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                </svg>
              </span>
              <input 
                v-model="searchStore.searchQuery"
                type="text" 
                placeholder="Search..." 
                class="w-full pl-10 pr-4 py-2 bg-white/10 border border-white/20 text-xs font-semibold rounded-full text-white placeholder-slate-300 focus:outline-none focus:bg-white/20 transition-all"
              />
            </div>

            <!-- Notification Bell Icon -->
            <div class="relative">
              <button 
                @click="isAlertsOpen = !isAlertsOpen" 
                class="p-2.5 bg-white/10 hover:bg-white/20 text-white rounded-full transition-all border border-white/20 relative cursor-pointer"
                title="Notifikasi Peringatan Kritis"
              >
                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
                </svg>
                <span v-if="unreadAlertsCount > 0" class="absolute top-1 right-1 w-2 h-2 rounded-full bg-rose-500 ring-2 ring-[#332941]"></span>
              </button>

              <!-- Alerts Dropdown Backdrop -->
              <div v-if="isAlertsOpen" @click="isAlertsOpen = false" class="fixed inset-0 z-40"></div>

              <!-- Alerts Dropdown Menu Card -->
              <div 
                v-if="isAlertsOpen" 
                class="absolute right-0 mt-2 w-72 bg-white rounded-2xl shadow-xl border border-slate-100 py-3 z-50 animate-fade-in text-slate-800"
              >
                <div class="px-4 pb-2 border-b border-slate-100 flex items-center justify-between">
                  <span class="text-xs font-bold text-slate-800">Notifikasi Kritis</span>
                  <span class="text-[10px] font-semibold text-rose-500 bg-rose-50 px-2 py-0.5 rounded-md">
                    {{ unreadAlertsCount }} baru
                  </span>
                </div>

                <div class="max-h-60 overflow-y-auto divide-y divide-slate-50">
                  <div 
                    v-for="alert in alerts" 
                    :key="alert.id"
                    class="p-3 hover:bg-slate-50/80 transition-colors flex flex-col space-y-1 relative"
                    :class="!alert.is_read ? 'bg-blue-50/10' : ''"
                  >
                    <div class="flex items-center justify-between">
                      <span class="text-[10px] font-bold text-slate-400">
                        {{ alert.survey ? alert.survey.title : 'Survei' }}
                      </span>
                      <button 
                        v-if="!alert.is_read" 
                        @click="handleMarkRead(alert.id)"
                        class="text-[9px] font-extrabold text-blue-600 hover:text-blue-800 underline cursor-pointer"
                      >
                        Tandai Dibaca
                      </button>
                    </div>
                    <p class="text-xs font-semibold text-slate-700 leading-snug">
                      Responden <span class="text-slate-900 font-bold">"{{ alert.respondent }}"</span> memberikan kepuasan rendah:
                    </p>
                    <div class="flex items-center justify-between mt-1 text-[11px] font-bold">
                      <span class="text-rose-600 bg-rose-50 px-2 py-0.5 rounded-md">★ {{ alert.score.toFixed(1) }} / 5.0</span>
                      <span class="text-slate-400 text-[10px] font-medium">{{ formatTimeAgo(alert.created_at) }}</span>
                    </div>
                  </div>

                  <div v-if="alerts.length === 0" class="py-6 text-center text-xs font-medium text-slate-400">
                    Tidak ada peringatan kritis.
                  </div>
                </div>
              </div>
            </div>

            <!-- User Profile Dropdown Pill -->
            <div class="relative">
              <button 
                @click="isSettingsOpen = !isSettingsOpen" 
                class="flex items-center space-x-2.5 bg-white/10 hover:bg-white/20 p-1.5 pr-3 rounded-full border border-white/20 transition-all"
              >
                <!-- Avatar -->
                <div class="w-8 h-8 rounded-full bg-gradient-to-tr from-amber-400 to-purple-600 text-white font-bold text-xs flex items-center justify-center shadow-xs">
                  LB
                </div>
                <span class="text-xs font-bold text-white hidden md:inline-block">Admin</span>
                <svg xmlns="http://www.w3.org/2000/svg" class="h-3.5 w-3.5 text-slate-300" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M19 9l-7 7-7-7" />
                </svg>
              </button>

              <!-- Backdrop to close dropdown -->
              <div v-if="isSettingsOpen" @click="isSettingsOpen = false" class="fixed inset-0 z-40"></div>

              <!-- Dropdown Menu Card -->
              <div 
                v-if="isSettingsOpen" 
                class="absolute right-0 mt-2 w-56 bg-white rounded-2xl shadow-xl border border-slate-100 py-2 z-50 animate-fade-in"
              >
                <!-- User Profile Summary -->
                <div class="px-4 py-3 border-b border-slate-100">
                  <p class="text-xs font-bold text-slate-800">HR Admin Laskar Buah</p>
                  <p class="text-[11px] text-slate-500 mt-0.5">hradmin@laskarbuah.com</p>
                </div>

                <!-- Menu Items -->
                <div class="p-1 space-y-0.5">
                  <!-- Backup Database Button -->
                  <button 
                    @click="handleBackupDatabase" 
                    :disabled="isBackingUp"
                    class="w-full flex items-center space-x-2 px-3 py-2 text-xs font-semibold text-slate-700 hover:bg-amber-100/60 hover:text-purple-900 rounded-xl transition-all disabled:opacity-50"
                  >
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 text-purple-700" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                      <path stroke-linecap="round" stroke-linejoin="round" d="M4 7v10c0 2.21 3.582 4 8 4s8-1.79 8-4V7M4 7c0 2.21 3.582 4 8 4s8-1.79 8-4M4 7c0-2.21 3.582-4 8-4s8 1.79 8 4m0 5c0 2.21-3.582 4-8 4s-8-1.79-8-4" />
                    </svg>
                    <span>{{ isBackingUp ? 'Downloading Backup...' : 'Backup Database' }}</span>
                  </button>

                  <div class="border-t border-slate-100 my-1"></div>

                  <!-- Logout Button -->
                  <button 
                    @click="handleLogout" 
                    class="w-full flex items-center space-x-2 px-3 py-2 text-xs font-semibold text-rose-600 hover:bg-rose-50 rounded-xl transition-all"
                  >
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 text-rose-500" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                      <path stroke-linecap="round" stroke-linejoin="round" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
                    </svg>
                    <span>Logout</span>
                  </button>
                </div>
              </div>
            </div>

          </div>
        </header>

        <!-- Dynamic Main Page View Area with #FFF5E0 Background & Curved Top-Left Corner -->
        <main class="flex-1 bg-[#FFF5E0] rounded-tl-[2.5rem] px-6 sm:px-10 py-8 overflow-y-auto shadow-inner">
          <router-view></router-view>
        </main>

      </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useSearchStore } from '../stores/search';
import { getCriticalAlerts, markAlertAsRead } from '../services/employee.service';
import api from '../services/api';

const route = useRoute();
const router = useRouter();
const searchStore = useSearchStore();
const isSettingsOpen = ref(false);
const isBackingUp = ref(false);

// Alert Notification Center States
const isAlertsOpen = ref(false);
const alerts = ref([]);

const fetchAlertsData = async () => {
  try {
    const res = await getCriticalAlerts();
    alerts.value = res.data || [];
  } catch (err) {
    console.error('Failed to load critical alerts:', err);
  }
};

const unreadAlertsCount = computed(() => {
  return alerts.value.filter(a => !a.is_read).length;
});

const handleMarkRead = async (alertId) => {
  try {
    await markAlertAsRead(alertId);
    // Refresh alerts
    fetchAlertsData();
  } catch (err) {
    console.error('Failed to mark alert as read:', err);
  }
};

const formatTimeAgo = (dateStr) => {
  if (!dateStr) return '';
  const date = new Date(dateStr);
  const seconds = Math.floor((new Date() - date) / 1000);
  
  let interval = Math.floor(seconds / 31536000);
  if (interval >= 1) return `${interval} tahun lalu`;
  interval = Math.floor(seconds / 2592000);
  if (interval >= 1) return `${interval} bulan lalu`;
  interval = Math.floor(seconds / 86400);
  if (interval >= 1) return `${interval} hari lalu`;
  interval = Math.floor(seconds / 3600);
  if (interval >= 1) return `${interval} jam lalu`;
  interval = Math.floor(seconds / 60);
  if (interval >= 1) return `${interval} menit lalu`;
  return 'Baru saja';
};

const pageTitle = computed(() => {
  if (route.path === '/dashboard' || route.path === '/') return 'Overview';
  if (route.path.startsWith('/dashboard/surveys')) return 'Surveys';
  if (route.path.startsWith('/dashboard/reports')) return 'Respondent Log';
  if (route.path.startsWith('/dashboard/action-plans')) return 'Action Plans';
  if (route.path.startsWith('/dashboard/employees')) return 'Manajemen Karyawan';
  return 'Overview';
});

const isRouteActive = (path) => {
  if (path === '/dashboard') {
    return route.path === '/dashboard' || route.path === '/';
  }
  return route.path.startsWith(path);
};

const handleBackupDatabase = async () => {
  isBackingUp.value = true;
  try {
    const res = await api.get('/admin/backup', { responseType: 'blob' });
    const blob = new Blob([res.data], { type: 'application/json' });
    const url = window.URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    const now = new Date();
    const dateStr = now.toISOString().split('T')[0];
    a.download = `hr_survey_backup_${dateStr}.json`;
    document.body.appendChild(a);
    a.click();
    window.URL.revokeObjectURL(url);
    document.body.removeChild(a);
    isSettingsOpen.value = false;
  } catch (err) {
    console.error('Backup database failed:', err);
    alert('Gagal mengunduh backup database. Silakan coba kembali.');
  } finally {
    isBackingUp.value = false;
  }
};

const handleLogout = () => {
  localStorage.removeItem('token');
  localStorage.removeItem('user');
  isSettingsOpen.value = false;
  router.push('/auth/login');
};

onMounted(() => {
  fetchAlertsData();
  // Poll alerts every 20 seconds for real-time feel
  const intervalId = setInterval(fetchAlertsData, 20000);
  return () => clearInterval(intervalId);
});
</script>

<style scoped>
.animate-fade-in {
  animation: fadeIn 0.2s ease-out forwards;
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(6px); }
  to { opacity: 1; transform: translateY(0); }
}
</style>
