<template>
  <div class="space-y-6 animate-fade-in">
    <!-- Respondent Log & Detail Split View Container -->
    <div class="flex flex-col lg:flex-row gap-6 items-start transition-all duration-300">
      
      <!-- Left Column (Full width when no detail is open, 2/3 width when detail panel is open) -->
      <div 
        class="flex-1 w-full space-y-4 transition-all duration-300"
        :class="selectedRespondent ? 'lg:w-2/3' : 'w-full'"
      >
        <!-- Standalone Header Card: Respondent Log Title & Search -->
        <div class="bg-white rounded-[2rem] border border-slate-100/80 shadow-md hover:shadow-lg transition-all duration-300 p-6 sm:p-7 flex flex-col gap-4.5">
          <div class="flex flex-col xl:flex-row xl:items-center xl:justify-between gap-4">
            <div>
              <h2 class="text-xl font-bold text-slate-800 tracking-tight">Respondent Log</h2>
              <p class="text-xs text-slate-400 mt-1">Review raw answers and submission timestamps from respondents.</p>
            </div>
            
            <!-- Filters & Search -->
            <div class="flex flex-wrap items-center gap-3">
              <!-- Select Survey Selector Dropdown -->
              <select 
                v-model="selectedSurveyId" 
                @change="handleSurveyChange" 
                class="px-4 py-2.5 bg-slate-50 border border-slate-200/80 rounded-xl text-xs font-semibold text-slate-700 focus:outline-none focus:bg-white focus:ring-2 focus:ring-blue-500/20 focus:border-blue-500 transition-all cursor-pointer"
              >
                <option v-for="survey in surveys" :key="survey.id" :value="survey.id">
                  {{ survey.title }}
                </option>
              </select>

              <!-- Score Filter Dropdown -->
              <select 
                v-model="scoreFilter"
                class="px-4 py-2.5 bg-slate-50 border border-slate-200/80 rounded-xl text-xs font-semibold text-slate-700 focus:outline-none focus:bg-white focus:ring-2 focus:ring-blue-500/20 focus:border-blue-500 transition-all cursor-pointer"
              >
                <option value="all">Semua Skor</option>
                <option value="above_60">Skor ≥ 60%</option>
                <option value="under_60">Skor &lt; 60%</option>
              </select>

              <!-- Method/Anon Filter Dropdown -->
              <select 
                v-model="anonFilter"
                class="px-4 py-2.5 bg-slate-50 border border-slate-200/80 rounded-xl text-xs font-semibold text-slate-700 focus:outline-none focus:bg-white focus:ring-2 focus:ring-blue-500/20 focus:border-blue-500 transition-all cursor-pointer"
              >
                <option value="all">Semua Metode</option>
                <option value="anon">Anonim</option>
                <option value="non_anon">Non-Anonim</option>
              </select>

              <!-- Search/Filter Input -->
              <div class="relative w-full sm:w-48">
                <span class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 text-slate-400" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                  </svg>
                </span>
                <input 
                  v-model="searchQuery"
                  type="text" 
                  placeholder="Cari nama..." 
                  class="w-full pl-9 pr-4 py-2.5 bg-slate-50 border border-slate-200/80 rounded-xl text-xs font-semibold text-slate-700 placeholder-slate-400 focus:outline-none focus:bg-white focus:ring-2 focus:ring-blue-500/20 focus:border-blue-500 transition-all"
                />
              </div>
            </div>
          </div>

          <!-- Second Row: Date Filter & Export Buttons -->
          <div class="flex flex-col md:flex-row items-stretch md:items-center justify-between gap-4 pt-3 border-t border-slate-100/60 no-print">
            <!-- Date Filter Inputs -->
            <div class="flex flex-wrap items-center gap-3">
              <select 
                v-model="datePreset"
                @change="handlePresetChange"
                class="px-3 py-2 bg-slate-50 border border-slate-200/80 rounded-xl text-xs font-semibold text-slate-700 focus:outline-none focus:bg-white focus:ring-2 focus:ring-blue-500/20 transition-all cursor-pointer shadow-2xs"
              >
                <option value="all">Semua Waktu</option>
                <option value="last_7">7 Hari Terakhir</option>
                <option value="last_30">30 Hari Terakhir</option>
                <option value="this_month">Bulan Ini</option>
                <option value="custom">Kustom Tanggal</option>
              </select>

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

            <!-- Export Buttons -->
            <div class="flex items-center space-x-2">
              <button 
                @click="exportToExcel"
                class="inline-flex items-center space-x-1.5 px-4.5 py-2.5 bg-emerald-600 hover:bg-emerald-700 active:bg-emerald-800 text-white text-xs font-bold rounded-xl shadow-md shadow-emerald-100 hover:-translate-y-0.5 transition-all cursor-pointer"
              >
                <svg xmlns="http://www.w3.org/2000/svg" class="h-4.5 w-4.5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                </svg>
                <span>Unduh Excel</span>
              </button>
              
              <button 
                @click="printReport"
                class="inline-flex items-center space-x-1.5 px-4.5 py-2.5 bg-blue-600 hover:bg-blue-700 active:bg-blue-800 text-white text-xs font-bold rounded-xl shadow-md shadow-blue-100 hover:-translate-y-0.5 transition-all cursor-pointer"
              >
                <svg xmlns="http://www.w3.org/2000/svg" class="h-4.5 w-4.5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M17 17h2a2 2 0 002-2v-4a2 2 0 00-2-2H5a2 2 0 00-2 2v4a2 2 0 002 2h2m2 4h6a2 2 0 002-2v-4a2 2 0 00-2-2H9a2 2 0 00-2 2v4a2 2 0 002 2zm8-12V5a2 2 0 00-2-2H9a2 2 0 00-2 2v4h10z" />
                </svg>
                <span>Cetak PDF</span>
              </button>
            </div>
          </div>
        </div>

        <!-- Standalone Table Card -->
        <div class="bg-white rounded-[2rem] border border-slate-100/80 shadow-md hover:shadow-lg transition-all duration-300 overflow-hidden p-2">
          <div class="overflow-x-auto rounded-[1.5rem]">
            <table class="w-full text-left border-collapse">
              <thead>
                <tr class="text-[11px] font-bold text-white uppercase tracking-wider bg-[#4647AE]">
                  <th class="py-4 px-4 text-center w-14 rounded-tl-[1.5rem]">No</th>
                  <th class="py-4 px-5 text-left">Respondent ID</th>
                  <th class="py-4 px-5 text-left w-36">Departemen</th>
                  <th class="py-4 px-4 text-center w-28">Skor</th>
                  <th class="py-4 px-4 text-center w-36">Score Rating</th>
                  <th class="py-4 px-4 text-center w-40">Completion Status</th>
                  <th class="py-4 px-5 text-center w-48">Submitted At</th>
                  <th class="py-4 px-5 text-center w-32 rounded-tr-[1.5rem]">Actions</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-slate-50 text-sm">
                <tr 
                  v-for="(res, index) in filteredRespondents" 
                  :key="res.id" 
                  class="hover:bg-slate-50/60 transition-colors group cursor-pointer"
                  :class="selectedRespondent?.id === res.id ? 'bg-blue-50/40 border-l-4 border-l-blue-600' : ''"
                >
                  <!-- No -->
                  <td class="py-4 px-4 text-center font-bold text-slate-400 text-xs" @click="openDetails(res)">
                    {{ index + 1 }}
                  </td>

                  <!-- Respondent ID -->
                  <td class="py-4 px-5 text-left" @click="openDetails(res)">
                    <div class="flex items-center space-x-3.5">
                      <img 
                        :src="getProfilePhoto(res)" 
                        class="w-9 h-9 rounded-xl object-cover flex-shrink-0 border border-slate-100 shadow-2xs" 
                        alt="Profile" 
                      />
                      <div class="min-w-0">
                        <span class="font-semibold text-slate-800 block truncate">{{ res.name }}</span>
                      </div>
                    </div>
                  </td>

                  <!-- Departemen -->
                  <td class="py-4 px-5 text-left" @click="openDetails(res)">
                    <span 
                      v-if="res.department && res.department !== '-'"
                      class="px-2.5 py-1 text-[11px] font-bold rounded-lg bg-blue-50 text-blue-700 border border-blue-100"
                    >
                      {{ res.department }}
                    </span>
                    <span v-else class="text-xs italic text-slate-400">-</span>
                  </td>

                  <!-- Skor -->
                  <td class="py-4 px-4 text-center font-extrabold text-slate-800" @click="openDetails(res)">
                    <span class="inline-flex items-center px-2.5 py-1 rounded-lg text-xs bg-slate-100 text-slate-700 font-bold border border-slate-200/60">
                      {{ Math.round((res.avgRating / 5) * 100) }}%
                    </span>
                  </td>

                  <!-- Score Rating -->
                  <td class="py-4 px-4 text-center font-bold" @click="openDetails(res)">
                    <span class="px-2.5 py-1 rounded-lg text-xs" :class="getScoreBadgeClass(res.avgRating)">
                      ★ {{ res.avgRating?.toFixed(1) || '0.0' }}
                    </span>
                  </td>

                  <!-- Completion Status -->
                  <td class="py-4 px-4 text-center" @click="openDetails(res)">
                    <span class="inline-flex px-2.5 py-0.5 text-xs font-semibold text-emerald-700 bg-emerald-50 rounded-full">Completed</span>
                  </td>

                  <!-- Submitted At -->
                  <td class="py-4 px-5 text-center text-slate-400 text-xs" @click="openDetails(res)">
                    {{ formatDate(res.submittedAt) }}
                  </td>

                  <!-- Actions -->
                  <td class="py-4 px-5 text-center">
                    <button 
                      @click="openDetails(res)" 
                      class="px-3 py-1.5 border border-slate-200 text-xs font-semibold text-slate-600 hover:text-blue-600 hover:border-blue-200 hover:bg-blue-50/50 active:bg-blue-50 rounded-xl transition-all whitespace-nowrap"
                    >
                      View Details
                    </button>
                  </td>
                </tr>
                <tr v-if="filteredRespondents.length === 0">
                  <td colspan="7" class="py-12 text-center text-slate-400 text-sm italic">
                    No respondents found for this survey.
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

      </div>

      <!-- Right Inline Detail Panel (Locked sticky in place while left table scrolls) -->
      <div 
        v-if="selectedRespondent" 
        class="w-full lg:w-1/3 bg-white rounded-[2rem] border border-slate-100 shadow-lg flex flex-col overflow-hidden animate-slide-left sticky top-24 h-[calc(100vh-7.5rem)]"
      >
        <!-- Header -->
        <div class="p-5 border-b border-slate-100 bg-slate-50/50 flex items-center justify-between">
          <div class="flex items-center space-x-3">
            <img 
              :src="getProfilePhoto(selectedRespondent)" 
              class="w-10 h-10 rounded-xl object-cover shadow-md border border-slate-100 flex-shrink-0" 
              alt="Profile" 
            />
            <div>
              <h3 class="text-sm font-bold text-slate-800 tracking-tight">{{ selectedRespondent.name }}</h3>
              <p class="text-[10px] text-slate-400 font-bold mt-0.5">{{ selectedRespondent.department || '-' }}</p>
            </div>
          </div>
          <button 
            @click="closeDetails" 
            class="p-1.5 text-slate-400 hover:text-slate-600 rounded-lg hover:bg-slate-100 transition-colors"
            title="Tutup Bilah Detail"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>

        <!-- Metadata Bar -->
        <div class="px-5 py-2.5 bg-white border-b border-slate-100 flex items-center justify-between text-xs">
          <div class="flex items-center space-x-1.5">
            <span class="text-slate-400 text-[11px]">Score Rating:</span>
            <span class="px-2 py-0.5 rounded-md font-bold text-xs" :class="getScoreBadgeClass(selectedRespondent.avgRating)">
              ★ {{ selectedRespondent.avgRating?.toFixed(1) || '0.0' }}
            </span>
          </div>
          <span class="text-slate-400 text-[10px]">{{ formatDate(selectedRespondent.submittedAt) }}</span>
        </div>

        <!-- Answers List (Scrollable Content) -->
        <div class="flex-1 overflow-y-auto overflow-x-hidden p-5 space-y-4">
          <div class="text-[11px] font-bold text-slate-400 uppercase tracking-wider">
            SUBMITTED ANSWERS ({{ selectedRespondent.answers?.length || 0 }})
          </div>

          <div 
            v-for="(qa, idx) in selectedRespondent.answers" 
            :key="idx" 
            class="p-3.5 bg-slate-50/70 border border-slate-100 rounded-xl space-y-2 hover:bg-slate-50 transition-colors"
          >
            <div class="flex justify-between items-start gap-2">
              <span class="text-[10px] font-bold text-blue-600 uppercase bg-blue-50 px-2 py-0.5 rounded-md flex-shrink-0">
                Soal #{{ idx + 1 }}
              </span>
              <span v-if="qa.score" class="px-2 py-0.5 bg-amber-50 text-amber-700 text-[10px] font-extrabold rounded-md border border-amber-100/50 flex-shrink-0">
                ★ {{ qa.score }} / 5
              </span>
            </div>
            <h4 class="font-semibold text-slate-700 text-xs leading-relaxed break-words">{{ qa.question }}</h4>
            <div class="text-xs text-slate-600 bg-white border border-slate-100 p-2.5 rounded-lg leading-relaxed break-words">
              {{ qa.answer || (qa.score ? qa.score + ' dari 5 bintang.' : '-') }}
            </div>
          </div>
        </div>

        <!-- Footer -->
        <div class="p-3.5 border-t border-slate-100 bg-slate-50/50 flex justify-end">
          <button 
            @click="closeDetails" 
            class="w-full bg-slate-800 hover:bg-slate-700 active:bg-slate-900 text-white font-semibold text-xs py-2.5 px-4 rounded-xl transition-all shadow-sm"
          >
            Tutup Bilah Detail
          </button>
        </div>

      </div>

    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue';
import { getSurveys, getSurveyReport, getSurveyResponses } from '../../services/survey.service';
import { useSearchStore } from '../../stores/search';
import * as XLSX from 'xlsx';

const surveys = ref([]);
const selectedSurveyId = ref(null);
const searchQuery = ref('');
const selectedRespondent = ref(null);
const searchStore = useSearchStore();

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
    fetchReportDetails();
  }
};

const handleCustomDateChange = () => {
  if (customStartDate.value && customEndDate.value) {
    fetchReportDetails();
  }
};

const activeReport = ref({
  avgScore: 0,
  totalResponses: 0,
  completionRate: 0,
  actionPlansCount: 0,
  strengths: '',
  improvements: '',
  categories: []
});

const respondents = ref([]);

const fetchInitialData = async () => {
  try {
    const res = await getSurveys();
    surveys.value = res.data;
    if (surveys.value.length > 0) {
      selectedSurveyId.value = surveys.value[0].id;
      await fetchReportDetails();
    }
  } catch (error) {
    console.error('Failed to load reports initial data:', error);
  }
};

const fetchReportDetails = async () => {
  if (!selectedSurveyId.value) return;
  try {
    const params = {};
    if (startDate.value) params.startDate = startDate.value;
    if (endDate.value) params.endDate = endDate.value;

    const [reportRes, responsesRes] = await Promise.all([
      getSurveyReport(selectedSurveyId.value, params),
      getSurveyResponses(selectedSurveyId.value, params)
    ]);
    activeReport.value = reportRes.data;
    respondents.value = responsesRes.data;
  } catch (error) {
    console.error('Failed to load report breakdown details:', error);
  }
};

const handleSurveyChange = () => {
  searchQuery.value = '';
  selectedRespondent.value = null;
  scoreFilter.value = 'all';
  anonFilter.value = 'all';
  datePreset.value = 'all';
  fetchReportDetails();
};

const exportToExcel = () => {
  if (filteredRespondents.value.length === 0) {
    alert('Tidak ada data responden untuk diunduh.');
    return;
  }

  const data = filteredRespondents.value.map((res, index) => {
    const row = {
      'No': index + 1,
      'Nama Responden': res.name || 'Anonim',
      'Email': res.email || 'Anonim',
      'Departemen': res.department || '-',
      'Skor Rata-rata': (res.avgRating || 0).toFixed(2),
      'Metode': (res.name || '').toLowerCase() === 'anonim' ? 'Anonim' : 'Identitas Asli',
      'Waktu Pengisian': res.submittedAt,
    };

    (res.answers || []).forEach((ans, ansIdx) => {
      row[`Pertanyaan ${ansIdx + 1}: ${ans.question}`] = ans.score !== undefined && ans.score !== null ? ans.score : ans.answer;
    });

    return row;
  });

  const worksheet = XLSX.utils.json_to_sheet(data);
  const workbook = XLSX.utils.book_new();
  XLSX.utils.book_append_sheet(workbook, worksheet, 'Respondent Log');

  const activeSurvey = surveys.value.find(s => s.id === selectedSurveyId.value);
  const surveyTitle = activeSurvey ? activeSurvey.title.replace(/[^a-zA-Z0-9]/g, '_') : 'Survey';
  const dateStr = new Date().toISOString().split('T')[0];

  XLSX.writeFile(workbook, `Laporan_${surveyTitle}_${dateStr}.xlsx`);
};

const printReport = () => {
  window.print();
};

const scoreFilter = ref('all');
const anonFilter = ref('all');

const filteredRespondents = computed(() => {
  let list = respondents.value || [];

  // Filter based on Score
  if (scoreFilter.value === 'under_60') {
    list = list.filter(res => Math.round(((res.avgRating || 0) / 5) * 100) < 60);
  } else if (scoreFilter.value === 'above_60') {
    list = list.filter(res => Math.round(((res.avgRating || 0) / 5) * 100) >= 60);
  }

  // Filter based on Anon/Non-Anon
  if (anonFilter.value === 'anon') {
    list = list.filter(res => (res.name || '').toLowerCase() === 'anonim');
  } else if (anonFilter.value === 'non_anon') {
    list = list.filter(res => (res.name || '').toLowerCase() !== 'anonim');
  }

  const localQuery = searchQuery.value.toLowerCase().trim();
  const globalQuery = searchStore.searchQuery.toLowerCase().trim();
  const query = localQuery || globalQuery;

  if (!query) return list;
  return list.filter(res => 
    (res.name || '').toLowerCase().includes(query)
  );
});


const openDetails = (respondent) => {
  selectedRespondent.value = respondent;
};

const closeDetails = () => {
  selectedRespondent.value = null;
};

const getProgressBarColor = (score) => {
  if (score >= 4.0) return 'bg-emerald-500';
  if (score >= 3.0) return 'bg-blue-500';
  return 'bg-rose-500';
};

const getScoreBadgeClass = (score) => {
  if (score >= 4.0) return 'bg-emerald-50 text-emerald-700';
  if (score >= 3.0) return 'bg-blue-50 text-blue-700';
  return 'bg-rose-50 text-rose-700';
};

const getProfilePhoto = (respondent) => {
  if (!respondent) return '/hrsurvey-pp/1.png';
  const seed = respondent.id || (respondent.name ? respondent.name.length : 0);
  const index = (seed % 4) + 1;
  return `/hrsurvey-pp/${index}.png`;
};

const formatDate = (dateStr) => {
  if (!dateStr) return '';
  const date = new Date(dateStr);
  return date.toLocaleString('en-US', { month: 'short', day: 'numeric', year: 'numeric', hour: '2-digit', minute: '2-digit' });
};

onMounted(() => {
  fetchInitialData();
});
</script>

<style scoped>
.animate-fade-in {
  animation: fadeIn 0.4s ease-out forwards;
}

.animate-slide-left {
  animation: slideLeft 0.35s cubic-bezier(0.16, 1, 0.3, 1) forwards;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(8px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes slideLeft {
  from {
    transform: translateX(100%);
  }
  to {
    transform: translateX(0);
  }
}

@media print {
  .no-print, button, select, input, aside, header, .border-t {
    display: none !important;
  }
  .bg-\[\#FFF5E0\] {
    background-color: white !important;
  }
  main {
    padding: 0 !important;
    background-color: white !important;
  }
  .max-w-lg, .lg\:w-2\/3, .w-full {
    width: 100% !important;
    max-width: 100% !important;
    flex-basis: 100% !important;
  }
  .bg-white {
    border: none !important;
    box-shadow: none !important;
  }
}
</style>
