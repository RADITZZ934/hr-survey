<template>
  <div class="space-y-8 animate-fade-in font-sans pb-6 select-none">
    
    <!-- Global Dashboard Date Range Filter Bar -->
    <div class="flex flex-col sm:flex-row items-stretch sm:items-center justify-between gap-4 bg-white border border-slate-100 p-4.5 rounded-[1.5rem] shadow-xs">
      <div class="flex items-center space-x-2.5">
        <div class="w-8 h-8 bg-blue-50 text-blue-600 rounded-lg flex items-center justify-center">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-4.5 w-4.5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
          </svg>
        </div>
        <span class="text-xs font-extrabold text-slate-700 uppercase tracking-wider">Rentang Analisis Data</span>
      </div>

      <div class="flex flex-col sm:flex-row items-stretch sm:items-center gap-3">
        <!-- Preset Range Dropdown -->
        <select 
          v-model="datePreset"
          @change="handlePresetChange"
          class="px-4 py-2 bg-slate-50 border border-slate-200/80 rounded-xl text-xs font-semibold text-slate-700 focus:outline-none focus:bg-white focus:ring-2 focus:ring-blue-500/20 transition-all cursor-pointer shadow-2xs"
        >
          <option value="all">Semua Waktu</option>
          <option value="last_7">7 Hari Terakhir</option>
          <option value="last_30">30 Hari Terakhir</option>
          <option value="this_month">Bulan Ini</option>
          <option value="custom">Kustom Tanggal</option>
        </select>

        <!-- Custom Date Inputs -->
        <div v-if="datePreset === 'custom'" class="flex items-center space-x-2">
          <input 
            v-model="customStartDate"
            type="date"
            @change="handleCustomDateChange"
            class="px-3 py-1.5 bg-slate-50 border border-slate-200/80 rounded-xl text-xs font-semibold text-slate-700 focus:outline-none focus:bg-white transition-all shadow-2xs"
          />
          <span class="text-xs text-slate-400 font-bold">s/d</span>
          <input 
            v-model="customEndDate"
            type="date"
            @change="handleCustomDateChange"
            class="px-3 py-1.5 bg-slate-50 border border-slate-200/80 rounded-xl text-xs font-semibold text-slate-700 focus:outline-none focus:bg-white transition-all shadow-2xs"
          />
        </div>
      </div>
    </div>

    <!-- Top Section: Portfolio / Hero Score Card (Left) + Trio Pastel Cards (Right) -->
    <div class="grid grid-cols-1 lg:grid-cols-12 gap-6 items-stretch">
      
      <!-- 1. Hero Blue Tint Card (Portfolio / Score Chart Card) - 6 Cols -->
      <div class="lg:col-span-6 bg-gradient-to-br from-[#eaf2fe] to-[#e1ebfa] rounded-[2rem] p-6 lg:p-7 border border-blue-100/80 shadow-xs flex flex-col justify-between relative overflow-hidden">
        
        <!-- Header Row -->
        <div class="flex justify-between items-start mb-4">
          <div>
            <div class="flex items-baseline space-x-2">
              <span class="text-3xl sm:text-4xl font-extrabold text-slate-900 tracking-tight">
                {{ stats.avgScore ? stats.avgScore.toFixed(1) : '4.8' }}
              </span>
              <span class="text-base font-bold text-slate-500">/ 5.0</span>
            </div>
            <p class="text-xs font-semibold text-slate-500 mt-1">Overall Satisfaction Score</p>
          </div>

          <!-- Trend Mode Selector Dropdown -->
          <div class="relative min-w-[160px]">
            <select 
              v-model="selectedTrendSurveyId"
              @change="updateChart"
              class="w-full px-3 py-2 bg-white/80 border border-blue-200/80 rounded-xl text-[11px] font-bold text-slate-700 focus:outline-none focus:ring-2 focus:ring-blue-500/20 transition-all cursor-pointer shadow-2xs"
            >
              <option value="all">Semua Survei</option>
              <option v-for="survey in surveys" :key="survey.id" :value="survey.id">
                {{ survey.title }}
              </option>
            </select>
          </div>
        </div>

        <!-- Trend Line Chart Area -->
        <div class="relative w-full h-36 my-2">
          <canvas ref="trendChartRef"></canvas>
        </div>

        <!-- Filter Time Pills Footer (1H, 24H, 1W, 1M, 1Y, ALL) -->
        <div class="flex items-center space-x-2 pt-2 text-[11px] font-bold text-slate-400">
          <button class="px-2.5 py-1 rounded-lg hover:text-slate-800 transition-colors">1H</button>
          <button class="px-2.5 py-1 rounded-lg hover:text-slate-800 transition-colors">24H</button>
          <button class="px-2.5 py-1 rounded-lg bg-white text-slate-900 shadow-xs">1W</button>
          <button class="px-2.5 py-1 rounded-lg hover:text-slate-800 transition-colors">1M</button>
          <button class="px-2.5 py-1 rounded-lg hover:text-slate-800 transition-colors">1Y</button>
          <button class="px-2.5 py-1 rounded-lg hover:text-slate-800 transition-colors">ALL</button>
        </div>

      </div>

      <!-- 2. Trio Pastel Asset Cards - 6 Cols (3 Cards Grid) -->
      <div class="lg:col-span-6 grid grid-cols-1 sm:grid-cols-3 gap-4">
        
        <!-- Pastel Card 1: Lavender / Purple (Total Submissions) -->
        <div class="bg-[#eee6f7] rounded-[2rem] p-5.5 border border-purple-100 flex flex-col justify-between shadow-xs hover:-translate-y-1 transition-all duration-300">
          <div>
            <div class="flex justify-between items-center text-slate-400">
              <span class="text-xs font-bold text-slate-700">Submissions</span>
              <button class="text-slate-400 hover:text-slate-600">⋮</button>
            </div>
            <div class="mt-3">
              <span class="text-2xl font-extrabold text-slate-900 block tracking-tight">
                {{ stats.totalResponses || 0 }}
              </span>
              <span class="text-[10px] font-medium text-slate-500">Total Responses</span>
            </div>
          </div>

          <div class="flex items-center justify-between mt-6 pt-2 border-t border-purple-200/40">
            <div class="bg-white/80 p-2 rounded-xl text-purple-700 shadow-2xs">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2.5">
                <path stroke-linecap="round" stroke-linejoin="round" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
              </svg>
            </div>
            <span class="text-[11px] font-bold text-purple-700">+0.14%</span>
          </div>
        </div>

        <!-- Pastel Card 2: Mint Green (Completion Rate) -->
        <div class="bg-[#e2f4ea] rounded-[2rem] p-5.5 border border-emerald-100 flex flex-col justify-between shadow-xs hover:-translate-y-1 transition-all duration-300">
          <div>
            <div class="flex justify-between items-center text-slate-400">
              <span class="text-xs font-bold text-slate-700">Completion</span>
              <button class="text-slate-400 hover:text-slate-600">⋮</button>
            </div>
            <div class="mt-3">
              <span class="text-2xl font-extrabold text-slate-900 block tracking-tight">
                {{ stats.completionRate || 0 }}%
              </span>
              <span class="text-[10px] font-medium text-slate-500">Target 80%</span>
            </div>
          </div>

          <div class="flex items-center justify-between mt-6 pt-2 border-t border-emerald-200/40">
            <div class="bg-white/80 p-2 rounded-xl text-emerald-700 shadow-2xs">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2.5">
                <path stroke-linecap="round" stroke-linejoin="round" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </div>
            <span class="text-[11px] font-bold text-emerald-700">+0.31%</span>
          </div>
        </div>

        <!-- Pastel Card 3: Warm Cream (Action Plans) -->
        <div class="bg-[#f9f2e3] rounded-[2rem] p-5.5 border border-amber-100 flex flex-col justify-between shadow-xs hover:-translate-y-1 transition-all duration-300">
          <div>
            <div class="flex justify-between items-center text-slate-400">
              <span class="text-xs font-bold text-slate-700">Action Plans</span>
              <button class="text-slate-400 hover:text-slate-600">⋮</button>
            </div>
            <div class="mt-3">
              <span class="text-2xl font-extrabold text-slate-900 block tracking-tight">
                {{ stats.actionPlansCount || 0 }}
              </span>
              <span class="text-[10px] font-medium text-slate-500">Active Tasks</span>
            </div>
          </div>

          <div class="flex items-center justify-between mt-6 pt-2 border-t border-amber-200/40">
            <div class="bg-white/80 p-2 rounded-xl text-amber-800 shadow-2xs">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2.5">
                <path stroke-linecap="round" stroke-linejoin="round" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-6 9l2 2 4-4" />
              </svg>
            </div>
            <span class="text-[11px] font-bold text-amber-800">+0.27%</span>
          </div>
        </div>

      </div>

    </div>

    <!-- Bottom Section: Active Surveys Table (Left - 8 Cols) + Dark Action Banner (Right - 4 Cols) -->
    <div class="grid grid-cols-1 lg:grid-cols-12 gap-8 items-start">
      
      <!-- Left: Active Surveys Table Container (8 Cols) -->
      <div class="lg:col-span-7 space-y-4">
        
        <!-- Table Header Control Bar -->
        <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-3">
          <div>
            <h2 class="text-xl font-extrabold text-slate-900 tracking-tight">Active Surveys</h2>
            <p class="text-xs text-slate-400">List of published surveys in database</p>
          </div>

          <!-- Filter Pill Dropdowns -->
          <div class="flex items-center space-x-2">
            <div class="bg-white border border-slate-200/60 px-3 py-1.5 rounded-full text-xs font-bold text-slate-600 flex items-center space-x-1 shadow-xs">
              <span>24h</span>
              <svg class="h-3.5 w-3.5 text-slate-400" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"/></svg>
            </div>
            <div class="bg-white border border-slate-200/60 px-3 py-1.5 rounded-full text-xs font-bold text-slate-600 flex items-center space-x-1 shadow-xs">
              <span>Top responses</span>
              <svg class="h-3.5 w-3.5 text-slate-400" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"/></svg>
            </div>
          </div>
        </div>

        <!-- Table Container -->
        <div class="bg-white rounded-[2rem] border border-slate-100/80 shadow-md hover:shadow-lg transition-all duration-300 overflow-hidden p-2">
          <div class="overflow-x-auto rounded-[1.5rem]">
            <table class="w-full text-left border-collapse">
              <thead>
                <tr class="text-[11px] font-bold text-white uppercase tracking-wider bg-[#E9B824]">
                  <th class="py-3.5 px-4 rounded-tl-[1.5rem]">Name</th>
                  <th class="py-3.5 px-3 text-center">Status</th>
                  <th class="py-3.5 px-4 text-right">Responses</th>
                  <th class="py-3.5 px-4 text-right rounded-tr-[1.5rem]">Start Date</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-amber-100/60 text-xs sm:text-sm font-medium">
                <tr v-for="survey in surveys" :key="survey.id" class="hover:bg-amber-100/40 transition-colors group">
                  
                  <!-- Title with Square Icon Badge -->
                  <td class="py-3.5 px-4">
                    <div class="flex items-center space-x-3">
                      <div class="w-9 h-9 bg-[#332941] text-white rounded-xl flex items-center justify-center font-bold text-xs shadow-xs group-hover:scale-105 transition-transform">
                        {{ survey.title.substring(0, 1).toUpperCase() }}
                      </div>
                      <span class="font-bold text-slate-800 group-hover:text-purple-700 transition-colors">
                        {{ survey.title }}
                      </span>
                    </div>
                  </td>

                  <!-- Status Badge -->
                  <td class="py-3.5 px-3 text-center">
                    <span 
                      class="inline-flex px-2.5 py-1 text-[11px] font-extrabold rounded-full capitalize"
                      :class="getStatusBadgeClass(survey.status)"
                    >
                      {{ survey.status }}
                    </span>
                  </td>

                  <!-- Dynamic Responses Count -->
                  <td class="py-3.5 px-4 text-right font-bold text-slate-700">
                    {{ survey.responses_count || 0 }}
                  </td>

                  <!-- Start Date -->
                  <td class="py-3.5 px-4 text-right text-slate-400 text-xs">
                    {{ formatDate(survey.start_date || survey.created_at) }}
                  </td>
                </tr>

                <tr v-if="surveys.length === 0">
                  <td colspan="4" class="py-12 text-center text-slate-400 text-xs italic">
                    No active surveys in database.
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

      </div>

      <!-- Right: Action Banner Card (#332941 Background with Laskar Buah Storefront Opacity Overlay) -->
      <div class="lg:col-span-5 bg-[#332941] text-white rounded-[2rem] p-7 shadow-lg relative overflow-hidden flex flex-col justify-between min-h-[260px] group">
        
        <!-- Storefront Background Image with Opacity & Blend Mode -->
        <img 
          src="/laskar_buah_storefront.jpg" 
          alt="Laskar Buah Storefront" 
          class="absolute inset-0 w-full h-full object-cover opacity-25 mix-blend-overlay pointer-events-none group-hover:scale-105 transition-transform duration-700"
        />

        <!-- Soft Overlay for text readability -->
        <div class="absolute inset-0 bg-gradient-to-t from-[#332941]/90 via-[#332941]/50 to-transparent pointer-events-none"></div>
        
        <!-- Top Title & Subtitle -->
        <div class="space-y-2.5 z-10">
          <div class="flex items-center space-x-2.5">
            <div class="p-2 bg-white/10 rounded-xl text-amber-300">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M4 7v10c0 2.21 3.582 4 8 4s8-1.79 8-4V7M4 7c0 2.21 3.582 4 8 4s8-1.79 8-4M4 7c0-2.21 3.582-4 8-4s8 1.79 8 4m0 5c0 2.21-3.582 4-8 4s-8-1.79-8-4" />
              </svg>
            </div>
            <h3 class="text-xl font-bold tracking-tight text-white">Backup Database</h3>
          </div>
          <p class="text-xs text-amber-100/80 leading-relaxed">
            Unduh cadangan data database PostgreSQL dalam format JSON kapan saja.
          </p>
        </div>

        <!-- Action Button -->
        <div class="pt-6 z-10">
          <button 
            @click="triggerBackup" 
            class="w-full sm:w-auto inline-flex items-center justify-center space-x-2 px-6 py-3 bg-[#FFF5E0] hover:bg-amber-100 active:bg-amber-200 text-[#332941] font-extrabold text-xs rounded-full transition-all shadow-md hover:scale-[1.02]"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 text-[#332941]" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2.5">
              <path stroke-linecap="round" stroke-linejoin="round" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4" />
            </svg>
            <span>Backup Database Now</span>
          </button>
        </div>

      </div>

    </div>

  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { getStats, getSurveyTrends } from '../../services/dashboard.service';
import { getSurveys } from '../../services/survey.service';
import api from '../../services/api';
import { Chart, registerables } from 'chart.js';

Chart.register(...registerables);

const stats = ref({
  avgScore: 0.0,
  totalResponses: 0,
  completionRate: 0,
  actionPlansCount: 0
});

const surveys = ref([]);
const trendChartRef = ref(null);
const selectedTrendSurveyId = ref('all');
const rawTrendsData = ref(null);
let chartInstance = null;

// Date Range Filter States
const datePreset = ref('all');
const customStartDate = ref('');
const customEndDate = ref('');

const startDate = computed(() => {
  if (datePreset.value === 'all') return '';
  const now = new Date();
  if (datePreset.value === 'last_7') {
    const d = new Date();
    d.setDate(now.getDate() - 6);
    return d.toISOString().split('T')[0];
  }
  if (datePreset.value === 'last_30') {
    const d = new Date();
    d.setDate(now.getDate() - 29);
    return d.toISOString().split('T')[0];
  }
  if (datePreset.value === 'this_month') {
    const d = new Date(now.getFullYear(), now.getMonth(), 1);
    return d.toISOString().split('T')[0];
  }
  if (datePreset.value === 'custom') return customStartDate.value;
  return '';
});

const endDate = computed(() => {
  if (datePreset.value === 'all') return '';
  if (datePreset.value === 'custom') return customEndDate.value;
  return new Date().toISOString().split('T')[0];
});

const handlePresetChange = () => {
  if (datePreset.value !== 'custom') {
    fetchDashboardData();
  }
};

const handleCustomDateChange = () => {
  if (customStartDate.value && customEndDate.value) {
    fetchDashboardData();
  }
};

const fetchDashboardData = async () => {
  try {
    const params = {};
    if (startDate.value) params.startDate = startDate.value;
    if (endDate.value) params.endDate = endDate.value;

    const [statsRes, surveysRes, trendsRes] = await Promise.all([
      getStats(params),
      getSurveys(),
      getSurveyTrends(params)
    ]);
    stats.value = statsRes.data;
    surveys.value = surveysRes.data;
    
    if (trendsRes && trendsRes.data) {
      rawTrendsData.value = trendsRes.data;
      initChart(trendsRes.data);
    }
  } catch (error) {
    console.error('Failed to load dashboard data:', error);
  }
};

const updateChart = () => {
  if (rawTrendsData.value) {
    initChart(rawTrendsData.value);
  }
};

const triggerBackup = async () => {
  try {
    const res = await api.get('/admin/backup', { responseType: 'blob' });
    const blob = new Blob([res.data], { type: 'application/json' });
    const url = window.URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    const dateStr = new Date().toISOString().split('T')[0];
    a.download = `hr_survey_backup_${dateStr}.json`;
    document.body.appendChild(a);
    a.click();
    window.URL.revokeObjectURL(url);
    document.body.removeChild(a);
  } catch (err) {
    console.error('Backup failed:', err);
    alert('Gagal mengunduh backup database.');
  }
};

const initChart = (trends) => {
  if (!trendChartRef.value) return;

  const ctx = trendChartRef.value.getContext('2d');
  if (chartInstance) {
    chartInstance.destroy();
  }

  // Beautiful curated colors for multiple surveys
  const colors = [
    { border: '#0284c7', fillStart: 'rgba(56, 189, 248, 0.35)', fillEnd: 'rgba(56, 189, 248, 0.01)' }, // Blue
    { border: '#7c3aed', fillStart: 'rgba(139, 92, 246, 0.35)', fillEnd: 'rgba(139, 92, 246, 0.01)' }, // Purple
    { border: '#ea580c', fillStart: 'rgba(249, 115, 22, 0.35)', fillEnd: 'rgba(249, 115, 22, 0.01)' }   // Orange
  ];

  let datasets = [];

  if (selectedTrendSurveyId.value === 'all') {
    // Mode Semua: Show all surveys as separate lines
    datasets = (trends.datasets || []).map((dataset, idx) => {
      const color = colors[idx % colors.length];
      const gradient = ctx.createLinearGradient(0, 0, 0, 150);
      gradient.addColorStop(0, color.fillStart);
      gradient.addColorStop(1, color.fillEnd);

      return {
        label: dataset.title,
        data: dataset.data,
        borderColor: color.border,
        backgroundColor: gradient,
        fill: true,
        tension: 0.4,
        borderWidth: 2,
        pointBackgroundColor: color.border,
        pointHoverRadius: 6,
      };
    });
  } else {
    // Mode Per Survey: Find and show only the selected survey
    const targetId = Number(selectedTrendSurveyId.value);
    const dataset = (trends.datasets || []).find(d => d.id === targetId);
    if (dataset) {
      const gradient = ctx.createLinearGradient(0, 0, 0, 150);
      gradient.addColorStop(0, 'rgba(56, 189, 248, 0.4)');
      gradient.addColorStop(1, 'rgba(56, 189, 248, 0.01)');

      datasets = [{
        label: dataset.title,
        data: dataset.data,
        borderColor: '#0284c7',
        backgroundColor: gradient,
        fill: true,
        tension: 0.4,
        borderWidth: 2.5,
        pointBackgroundColor: '#0284c7',
        pointHoverRadius: 6,
      }];
    }
  }

  chartInstance = new Chart(ctx, {
    type: 'line',
    data: {
      labels: trends.labels || ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'],
      datasets: datasets.length > 0 ? datasets : [{
        label: 'Trend',
        data: [0, 0, 0, 0, 0, 0, 0],
        borderColor: '#0284c7',
        tension: 0.4,
        borderWidth: 2.5,
      }],
    },
    options: {
      responsive: true,
      maintainAspectRatio: false,
      plugins: {
        legend: { 
          display: selectedTrendSurveyId.value === 'all',
          position: 'top',
          align: 'end',
          labels: {
            boxWidth: 8,
            boxHeight: 8,
            padding: 10,
            usePointStyle: true,
            font: { size: 9, weight: 'bold', family: 'sans-serif' },
            color: '#475569'
          }
        },
        tooltip: {
          backgroundColor: '#18191c',
          padding: 10,
          cornerRadius: 8,
        }
      },
      scales: {
        x: { display: false },
        y: { display: false }
      }
    }
  });
};

const getStatusBadgeClass = (status) => {
  if (status === 'active') return 'text-emerald-700 bg-emerald-100/80';
  if (status === 'closed') return 'text-rose-700 bg-rose-100/80';
  return 'text-slate-600 bg-slate-100';
};

const formatDate = (dateStr) => {
  if (!dateStr) return '';
  const date = new Date(dateStr);
  return date.toLocaleDateString('en-US', { month: 'short', day: 'numeric', year: 'numeric' });
};

onMounted(() => {
  fetchDashboardData();
});
</script>

<style scoped>
.animate-fade-in {
  animation: fadeIn 0.3s cubic-bezier(0.16, 1, 0.3, 1) forwards;
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(8px); }
  to { opacity: 1; transform: translateY(0); }
}
</style>
