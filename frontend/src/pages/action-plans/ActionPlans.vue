<template>
  <div class="space-y-8 animate-fade-in">
    <!-- Header Section -->
    <div class="flex flex-col md:flex-row md:items-center md:justify-between gap-4 bg-white rounded-[2rem] border border-slate-100/80 shadow-md hover:shadow-lg transition-all duration-300 p-6 lg:p-8">
      <div>
        <h1 class="text-3xl font-bold text-slate-800 tracking-tight">Action Plans</h1>
        <p class="text-sm text-slate-400 mt-1">Manage and track improvement tasks derived from survey feedback.</p>
      </div>

      <!-- Filters & Actions -->
      <div class="flex flex-col sm:flex-row items-stretch sm:items-center gap-4">
        <!-- Survey Filter -->
        <div class="relative min-w-[200px]">
          <select 
            v-model="selectedSurveyId" 
            @change="fetchPlans"
            class="w-full px-4 py-2.5 bg-slate-50 border border-slate-200/80 rounded-xl text-sm font-semibold text-slate-700 focus:outline-none focus:bg-white focus:ring-2 focus:ring-blue-500/20 focus:border-blue-500 transition-all cursor-pointer"
          >
            <option :value="null">All Surveys</option>
            <option v-for="survey in surveys" :key="survey.id" :value="survey.id">
              {{ survey.title }}
            </option>
          </select>
        </div>

        <!-- Create Button -->
        <button 
          @click="openCreateModal"
          class="inline-flex items-center justify-center space-x-2 px-5 py-2.5 bg-blue-600 hover:bg-blue-700 active:bg-blue-800 text-white text-sm font-semibold rounded-xl transition-all shadow-sm shadow-blue-200"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
            <path fill-rule="evenodd" d="M10 3a1 1 0 011 1v5h5a1 1 0 110 2h-5v5a1 1 0 11-2 0v-5H4a1 1 0 110-2h5V4a1 1 0 011-1z" clip-rule="evenodd" />
          </svg>
          <span>New Action Plan</span>
        </button>
      </div>
    </div>

    <!-- Kanban Board Grid -->
    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      
      <!-- Pending Column -->
      <div class="bg-slate-50 rounded-3xl p-5 border border-slate-100/50 flex flex-col min-h-[500px]">
        <div class="flex items-center justify-between mb-4 px-2">
          <div class="flex items-center space-x-2">
            <span class="w-2.5 h-2.5 rounded-full bg-slate-400"></span>
            <h2 class="text-sm font-bold text-slate-600 uppercase tracking-wider">Pending</h2>
          </div>
          <span class="px-2 py-0.5 bg-slate-200/60 text-slate-600 font-bold text-xs rounded-full">
            {{ pendingPlans.length }}
          </span>
        </div>

        <!-- Cards list -->
        <div class="space-y-4 flex-1 overflow-y-auto max-h-[600px] pr-1">
          <div 
            v-for="plan in pendingPlans" 
            :key="plan.id" 
            class="bg-white p-5 rounded-2xl border border-slate-100 hover:border-slate-200 shadow-sm hover:shadow-md transition-all duration-300 space-y-4 group relative"
          >
            <div class="space-y-1">
              <span class="text-[10px] font-bold text-slate-400 uppercase tracking-wider block">
                {{ getSurveyTitle(plan.surveyId) }}
              </span>
              <h3 class="font-bold text-slate-800 text-sm group-hover:text-blue-600 transition-colors">
                {{ plan.title }}
              </h3>
            </div>
            <p class="text-xs text-slate-500 leading-relaxed line-clamp-3">
              {{ plan.description }}
            </p>
            <div class="flex items-center justify-between pt-2 border-t border-slate-50 text-[11px] font-semibold text-slate-400">
              <div class="flex items-center space-x-1.5">
                <div class="bg-slate-100 text-slate-600 w-5 h-5 rounded-full flex items-center justify-center font-bold text-[9px]">
                  {{ plan.assigneeInitials }}
                </div>
                <span>{{ plan.assigneeName }}</span>
              </div>
              <span :class="isOverdue(plan.targetDate) ? 'text-rose-500' : ''">
                Due: {{ formatDate(plan.targetDate) }}
              </span>
            </div>

            <!-- Action buttons to move -->
            <div class="flex justify-end pt-2">
              <button 
                @click="changePlanStatus(plan.id, 'in_progress')"
                class="w-full py-1.5 bg-blue-50/50 hover:bg-blue-50 active:bg-blue-100 text-blue-600 text-xs font-bold rounded-xl transition-all flex items-center justify-center space-x-1"
              >
                <span>Start Progress</span>
                <svg xmlns="http://www.w3.org/2000/svg" class="h-3.5 w-3.5" viewBox="0 0 20 20" fill="currentColor">
                  <path fill-rule="evenodd" d="M12.293 5.293a1 1 0 011.414 0l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414-1.414L14.586 11H3a1 1 0 110-2h11.586l-2.293-2.293a1 1 0 010-1.414z" clip-rule="evenodd" />
                </svg>
              </button>
            </div>
          </div>

          <div v-if="pendingPlans.length === 0" class="py-12 text-center text-slate-400 text-xs italic">
            No pending action plans.
          </div>
        </div>
      </div>

      <!-- In Progress Column -->
      <div class="bg-blue-50/30 rounded-3xl p-5 border border-blue-50/50 flex flex-col min-h-[500px]">
        <div class="flex items-center justify-between mb-4 px-2">
          <div class="flex items-center space-x-2">
            <span class="w-2.5 h-2.5 rounded-full bg-blue-500 animate-pulse"></span>
            <h2 class="text-sm font-bold text-blue-700 uppercase tracking-wider">In Progress</h2>
          </div>
          <span class="px-2 py-0.5 bg-blue-100 text-blue-700 font-bold text-xs rounded-full">
            {{ inProgressPlans.length }}
          </span>
        </div>

        <!-- Cards list -->
        <div class="space-y-4 flex-1 overflow-y-auto max-h-[600px] pr-1">
          <div 
            v-for="plan in inProgressPlans" 
            :key="plan.id" 
            class="bg-white p-5 rounded-2xl border border-slate-100 hover:border-slate-200 shadow-sm hover:shadow-md transition-all duration-300 space-y-4 group relative"
          >
            <div class="space-y-1">
              <span class="text-[10px] font-bold text-slate-400 uppercase tracking-wider block">
                {{ getSurveyTitle(plan.surveyId) }}
              </span>
              <h3 class="font-bold text-slate-800 text-sm group-hover:text-blue-600 transition-colors">
                {{ plan.title }}
              </h3>
            </div>
            <p class="text-xs text-slate-500 leading-relaxed line-clamp-3">
              {{ plan.description }}
            </p>
            <div class="flex items-center justify-between pt-2 border-t border-slate-50 text-[11px] font-semibold text-slate-400">
              <div class="flex items-center space-x-1.5">
                <div class="bg-blue-50 text-blue-600 w-5 h-5 rounded-full flex items-center justify-center font-bold text-[9px]">
                  {{ plan.assigneeInitials }}
                </div>
                <span>{{ plan.assigneeName }}</span>
              </div>
              <span :class="isOverdue(plan.targetDate) ? 'text-rose-500' : ''">
                Due: {{ formatDate(plan.targetDate) }}
              </span>
            </div>

            <!-- Action buttons to move -->
            <div class="flex gap-2 pt-2">
              <button 
                @click="changePlanStatus(plan.id, 'pending')"
                class="flex-1 py-1.5 border border-slate-200 text-slate-500 hover:text-slate-700 hover:bg-slate-50 text-xs font-bold rounded-xl transition-all"
              >
                Put Back
              </button>
              <button 
                @click="changePlanStatus(plan.id, 'completed')"
                class="flex-1 py-1.5 bg-emerald-600 hover:bg-emerald-700 active:bg-emerald-800 text-white text-xs font-bold rounded-xl transition-all flex items-center justify-center space-x-1"
              >
                <svg xmlns="http://www.w3.org/2000/svg" class="h-3.5 w-3.5" viewBox="0 0 20 20" fill="currentColor">
                  <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd" />
                </svg>
                <span>Complete</span>
              </button>
            </div>
          </div>

          <div v-if="inProgressPlans.length === 0" class="py-12 text-center text-slate-400 text-xs italic">
            No action plans currently in progress.
          </div>
        </div>
      </div>

      <!-- Completed Column -->
      <div class="bg-emerald-50/20 rounded-3xl p-5 border border-emerald-50/30 flex flex-col min-h-[500px]">
        <div class="flex items-center justify-between mb-4 px-2">
          <div class="flex items-center space-x-2">
            <span class="w-2.5 h-2.5 rounded-full bg-emerald-500"></span>
            <h2 class="text-sm font-bold text-emerald-700 uppercase tracking-wider">Completed</h2>
          </div>
          <span class="px-2 py-0.5 bg-emerald-100 text-emerald-700 font-bold text-xs rounded-full">
            {{ completedPlans.length }}
          </span>
        </div>

        <!-- Cards list -->
        <div class="space-y-4 flex-1 overflow-y-auto max-h-[600px] pr-1">
          <div 
            v-for="plan in completedPlans" 
            :key="plan.id" 
            class="bg-white/80 p-5 rounded-2xl border border-slate-100/80 shadow-sm opacity-90 hover:opacity-100 transition-all duration-300 space-y-4 group relative"
          >
            <div class="space-y-1">
              <span class="text-[10px] font-bold text-slate-400 uppercase tracking-wider block">
                {{ getSurveyTitle(plan.surveyId) }}
              </span>
              <h3 class="font-bold text-slate-500 text-sm line-through">
                {{ plan.title }}
              </h3>
            </div>
            <p class="text-xs text-slate-400 leading-relaxed line-clamp-3">
              {{ plan.description }}
            </p>
            <div class="flex items-center justify-between pt-2 border-t border-slate-50 text-[11px] font-semibold text-slate-400">
              <div class="flex items-center space-x-1.5">
                <div class="bg-emerald-50 text-emerald-600 w-5 h-5 rounded-full flex items-center justify-center font-bold text-[9px]">
                  {{ plan.assigneeInitials }}
                </div>
                <span>{{ plan.assigneeName }}</span>
              </div>
              <span class="text-emerald-600 font-bold flex items-center space-x-0.5">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-3 w-3" viewBox="0 0 20 20" fill="currentColor">
                  <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
                </svg>
                <span>Finished</span>
              </span>
            </div>

            <!-- Action buttons to move -->
            <div class="flex justify-start pt-2">
              <button 
                @click="changePlanStatus(plan.id, 'in_progress')"
                class="w-full py-1.5 border border-slate-200 hover:bg-slate-50 text-slate-600 text-xs font-bold rounded-xl transition-all"
              >
                Reopen Task
              </button>
            </div>
          </div>

          <div v-if="completedPlans.length === 0" class="py-12 text-center text-slate-400 text-xs italic">
            No completed action plans.
          </div>
        </div>
      </div>

    </div>

    <!-- Create Plan Modal -->
    <div v-if="isCreateModalOpen" class="fixed inset-0 z-50 overflow-hidden" role="dialog" aria-modal="true">
      <div class="flex items-center justify-center min-h-screen px-4 pt-4 pb-20 text-center sm:block sm:p-0">
        <!-- Backdrop -->
        <div @click="closeCreateModal" class="fixed inset-0 transition-opacity bg-slate-900/40 backdrop-blur-sm"></div>

        <!-- Modal Box (Wider 2-Column Layout) -->
        <div 
          class="inline-block align-bottom bg-white rounded-[2rem] text-left overflow-hidden shadow-2xl transform transition-all sm:my-8 sm:align-middle sm:max-w-4xl sm:w-full border border-slate-100"
        >
          <!-- Header -->
          <div class="p-6 border-b border-slate-100 flex items-center justify-between bg-slate-50/20">
            <div class="flex items-center space-x-2.5">
              <span class="w-2.5 h-2.5 rounded-full bg-blue-600"></span>
              <h3 class="text-lg font-extrabold text-slate-800 tracking-tight">Rencana Tindakan Baru (Action Plan)</h3>
            </div>
            <button 
              type="button" 
              @click="closeCreateModal" 
              class="p-2 text-slate-400 hover:text-slate-600 rounded-xl hover:bg-slate-100 transition-colors"
            >
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>

          <!-- Split Container -->
          <div class="flex flex-col md:flex-row divide-y md:divide-y-0 md:divide-x divide-slate-100">
            
            <!-- Left Column: Form Fields (2/3 width) -->
            <form @submit.prevent="saveActionPlan" class="flex-1 p-6 space-y-4 text-sm">
              <!-- Title -->
              <div class="space-y-1">
                <label class="block text-xs font-bold text-slate-500 uppercase tracking-wider">Judul Rencana (Plan Title)</label>
                <input 
                  v-model="newPlan.title"
                  type="text" 
                  placeholder="Contoh: Menyelenggarakan Sesi Konseling Mental"
                  required
                  class="w-full px-4 py-2.5 bg-slate-50 border border-slate-200/80 rounded-xl text-xs font-semibold text-slate-700 placeholder-slate-400 focus:outline-none focus:bg-white focus:ring-2 focus:ring-blue-500/20 focus:border-blue-500 transition-all"
                />
              </div>

              <!-- Description -->
              <div class="space-y-1">
                <label class="block text-xs font-bold text-slate-500 uppercase tracking-wider">Deskripsi Lengkap</label>
                <textarea 
                  v-model="newPlan.description"
                  placeholder="Jelaskan langkah kerja, tujuan, dan strategi pelaksanaan program..."
                  required
                  rows="4"
                  class="w-full px-4 py-2.5 bg-slate-50 border border-slate-200/80 rounded-xl text-xs font-semibold text-slate-700 placeholder-slate-400 focus:outline-none focus:bg-white focus:ring-2 focus:ring-blue-500/20 focus:border-blue-500 transition-all resize-none"
                ></textarea>
              </div>

              <!-- Linked Survey -->
              <div class="space-y-1">
                <label class="block text-xs font-bold text-slate-500 uppercase tracking-wider">Tautkan ke Kuesioner</label>
                <select 
                  v-model="newPlan.surveyId"
                  required
                  class="w-full px-4 py-2.5 bg-slate-50 border border-slate-200/80 rounded-xl text-xs font-semibold text-slate-700 focus:outline-none focus:bg-white focus:ring-2 focus:ring-blue-500/20 focus:border-blue-500 transition-all cursor-pointer bg-white"
                >
                  <option v-for="survey in surveys" :key="survey.id" :value="survey.id">
                    {{ survey.title }}
                  </option>
                </select>
              </div>

              <!-- Assignee and Due Date Grid -->
              <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
                <!-- Assignee -->
                <div class="space-y-1">
                  <label class="block text-xs font-bold text-slate-500 uppercase tracking-wider">Penanggung Jawab</label>
                  <select 
                    v-model="newPlan.assigneeIndex"
                    required
                    class="w-full px-4 py-2.5 bg-slate-50 border border-slate-200/80 rounded-xl text-xs font-semibold text-slate-700 focus:outline-none focus:bg-white focus:ring-2 focus:ring-blue-500/20 focus:border-blue-500 transition-all cursor-pointer bg-white"
                  >
                    <option v-for="(person, idx) in assignees" :key="idx" :value="idx">
                      {{ person.name }} ({{ person.role }})
                    </option>
                  </select>
                </div>

                <!-- Due Date -->
                <div class="space-y-1">
                  <label class="block text-xs font-bold text-slate-500 uppercase tracking-wider">Tanggal Tenggat</label>
                  <CustomDatePicker 
                    v-model="newPlan.targetDate"
                    placeholder="Pilih Tanggal Tenggat"
                  />
                </div>
              </div>

              <!-- Form Footer -->
              <div class="pt-4 flex justify-end space-x-3 border-t border-slate-100">
                <button 
                  type="button"
                  @click="closeCreateModal" 
                  class="px-5 py-2.5 bg-slate-100 hover:bg-slate-200 text-slate-700 text-xs font-bold rounded-xl transition-colors cursor-pointer"
                >
                  Batal
                </button>
                <button 
                  type="submit" 
                  class="px-5 py-2.5 bg-[#4647AE] hover:bg-[#383994] active:bg-[#2e2e7a] text-white text-xs font-bold rounded-xl shadow-md shadow-indigo-100 hover:-translate-y-0.5 transition-all cursor-pointer"
                >
                  Simpan Program Kerja
                </button>
              </div>
            </form>

            <!-- Right Column: AI Suggestions Panel (1/3 width) -->
            <div class="w-full md:w-80 p-6 bg-slate-50/50 flex flex-col space-y-4">
              <div class="flex items-center space-x-2">
                <div class="w-7 h-7 bg-indigo-50 text-indigo-600 rounded-lg flex items-center justify-center font-bold text-xs">
                  🤖
                </div>
                <div>
                  <h4 class="text-xs font-extrabold text-slate-700 uppercase tracking-wider">Asisten Asisten AI</h4>
                  <p class="text-[10px] text-slate-400 mt-0.5">Saran Rencana Program Kerja</p>
                </div>
              </div>

              <!-- Category Dropdown to view recommendations -->
              <div class="space-y-1">
                <label class="block text-[10px] font-bold text-slate-500 uppercase tracking-wider">Kategori Fokus</label>
                <select 
                  v-model="aiSelectedCategory"
                  class="w-full px-3 py-2 bg-white border border-slate-200 rounded-xl text-xs font-semibold text-slate-700 focus:outline-none focus:ring-2 focus:ring-blue-500/20 transition-all cursor-pointer"
                >
                  <option v-for="cat in Object.keys(aiRecommendations)" :key="cat" :value="cat">
                    {{ cat }}
                  </option>
                </select>
              </div>

              <!-- Recommendations List -->
              <div class="space-y-3 flex-1 overflow-y-auto max-h-[300px] pr-1">
                <button 
                  v-for="(rec, idx) in currentAiRecommendations" 
                  :key="idx"
                  @click="applyRecommendation(rec)"
                  type="button"
                  class="w-full text-left p-3.5 bg-white border border-slate-200 hover:border-indigo-300 hover:bg-indigo-50/10 rounded-2xl shadow-2xs hover:shadow-sm transition-all group flex flex-col space-y-1.5 cursor-pointer"
                >
                  <span class="text-xs font-bold text-slate-800 group-hover:text-indigo-600 transition-colors leading-snug">
                    {{ rec.title }}
                  </span>
                  <span class="text-[10px] text-slate-500 leading-normal line-clamp-3">
                    {{ rec.description }}
                  </span>
                  <span class="text-[9px] font-extrabold text-indigo-500 uppercase tracking-wide flex items-center space-x-1 pt-1">
                    <span>Gunakan Template</span>
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-3 w-3 group-hover:translate-x-0.5 transition-transform" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2.5">
                      <path stroke-linecap="round" stroke-linejoin="round" d="M14 5l7 7m0 0l-7 7m7-7H3" />
                    </svg>
                  </span>
                </button>
              </div>
            </div>

          </div>

        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { getActionPlans, createActionPlan, updateActionPlan } from '../../services/actionplan.service';
import { getSurveys } from '../../services/survey.service';
import CustomDatePicker from '../../components/CustomDatePicker.vue';

const surveys = ref([]);
const selectedSurveyId = ref(null);
const isCreateModalOpen = ref(false);
const actionPlans = ref([]);

// Mock Assignees to select
const assignees = [
  { name: 'Diana R.', role: 'HR Director', id: 2 },
  { name: 'John Doe', role: 'Operations Manager', id: 3 },
  { name: 'Sarah K.', role: 'Welfare Coordinator', id: 4 }
];

// New Plan Initial State
const newPlan = ref({
  title: '',
  description: '',
  surveyId: 1,
  assigneeIndex: 0,
  targetDate: ''
});

// AI Recommendation Assistant States & Database
const aiSelectedCategory = ref('Work-Life Balance');

const aiRecommendations = {
  'Work-Life Balance': [
    {
      title: 'Workshop Pengaturan Beban Kerja & Batasan Jam Kerja',
      description: 'Menyelenggarakan sesi edukasi bagi manajer dan staf mengenai pentingnya batasan waktu kerja, pembagian tugas yang adil, serta pencegahan burnout.'
    },
    {
      title: 'Evaluasi Kebijakan Kerja Fleksibel (WFA)',
      description: 'Meninjau ulang efisiensi jam kerja fleksibel dan memberikan fasilitas pendukung bagi staf yang bekerja jarak jauh agar tetap seimbang.'
    }
  ],
  'Team Collaboration': [
    {
      title: 'Program Team Building Lintas Departemen Bulanan',
      description: 'Mengadakan aktivitas santai dan kolaboratif antar divisi setiap bulan untuk mencairkan ketegangan komunikasi dan mempererat kerja sama.'
    },
    {
      title: 'Pembaruan Saluran Komunikasi Slack/Teams',
      description: 'Menyusun pedoman komunikasi tertulis yang jelas serta mengaktifkan ruang diskusi informal untuk mempercepat respon kerja tim.'
    }
  ],
  'Manager Support': [
    {
      title: 'Pelatihan Leadership & Coaching untuk Kepala Divisi',
      description: 'Membekali para kepala divisi dengan keahlian kepemimpinan empati, cara memberikan feedback konstruktif, serta penyelesaian konflik.'
    },
    {
      title: 'Sesi Pendampingan 1-on-1 Berkala',
      description: 'Mewajibkan sesi ngobrol pribadi antara manajer dan staf setiap dua minggu sekali untuk menampung keluh kesah secara langsung.'
    }
  ],
  'Compensation & Benefits': [
    {
      title: 'Transparansi Struktur Gaji & Review Bonus Tahunan',
      description: 'Menyusun dokumen panduan grading jabatan serta merumuskan skema bonus performa yang adil dan transparan.'
    },
    {
      title: 'Peningkatan Tunjangan Kesehatan & Kesejahteraan Karyawan',
      description: 'Menambahkan fasilitas wellness allowance (seperti subsidi keanggotaan gym, kelas yoga, atau layanan konseling kesehatan mental).'
    }
  ],
  'Career Growth': [
    {
      title: 'Penyusunan Peta Jalan Karier (Career Path Framework)',
      description: 'Membuat visualisasi alur promosi jabatan yang transparan lengkap dengan kompetensi teknis/non-teknis yang wajib dipenuhi staf.'
    },
    {
      title: 'Penyediaan Anggaran Pelatihan Karyawan Mandiri',
      description: 'Menyediakan tunjangan tahunan khusus bagi staf untuk mengambil sertifikasi profesional, kursus online, atau menghadiri seminar.'
    }
  ]
};

const currentAiRecommendations = computed(() => {
  return aiRecommendations[aiSelectedCategory.value] || [];
});

const applyRecommendation = (rec) => {
  newPlan.value.title = rec.title;
  newPlan.value.description = rec.description;
};

const fetchInitialData = async () => {
  try {
    const res = await getSurveys();
    surveys.value = res.data;
    await fetchPlans();
  } catch (error) {
    console.error('Failed to load action plans initial data:', error);
  }
};

const fetchPlans = async () => {
  try {
    const res = await getActionPlans(selectedSurveyId.value);
    actionPlans.value = res.data;
  } catch (error) {
    console.error('Failed to fetch action plans:', error);
  }
};

// Filtered Lists per Swimlane
const pendingPlans = computed(() => {
  return actionPlans.value.filter(plan => plan.status === 'pending');
});

const inProgressPlans = computed(() => {
  return actionPlans.value.filter(plan => plan.status === 'in_progress');
});

const completedPlans = computed(() => {
  return actionPlans.value.filter(plan => plan.status === 'completed');
});

const getSurveyTitle = (id) => {
  const s = surveys.value.find(item => item.id === id);
  return s ? s.title : 'Survey';
};

const formatDate = (dateStr) => {
  if (!dateStr) return 'N/A';
  const options = { year: 'numeric', month: 'short', day: 'numeric' };
  return new Date(dateStr).toLocaleDateString('en-US', options);
};

const isOverdue = (dateStr) => {
  if (!dateStr) return false;
  return new Date(dateStr) < new Date() && !actionPlans.value.find(p => p.targetDate === dateStr && p.status === 'completed');
};

const changePlanStatus = async (planId, newStatus) => {
  try {
    await updateActionPlan(planId, { status: newStatus });
    await fetchPlans();
  } catch (error) {
    console.error('Failed to update action plan status:', error);
  }
};

// Modal handlers
const openCreateModal = () => {
  newPlan.value = {
    title: '',
    description: '',
    surveyId: selectedSurveyId.value || (surveys.value.length > 0 ? surveys.value[0].id : 1),
    assigneeIndex: 0,
    targetDate: new Date(Date.now() + 14 * 24 * 60 * 60 * 1000).toISOString().split('T')[0] // default 2 weeks ahead
  };
  isCreateModalOpen.value = true;
};

const closeCreateModal = () => {
  isCreateModalOpen.value = false;
};

const saveActionPlan = async () => {
  const selectedAssignee = assignees[newPlan.value.assigneeIndex];
  try {
    await createActionPlan({
      surveyId: Number(newPlan.value.surveyId),
      title: newPlan.value.title,
      description: newPlan.value.description,
      targetDate: newPlan.value.targetDate,
      assigneeId: selectedAssignee.id
    });
    closeCreateModal();
    await fetchPlans();
  } catch (error) {
    console.error('Failed to create action plan:', error);
  }
};

onMounted(() => {
  fetchInitialData();
});
</script>

<style scoped>
.animate-fade-in {
  animation: fadeIn 0.4s ease-out forwards;
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
</style>
