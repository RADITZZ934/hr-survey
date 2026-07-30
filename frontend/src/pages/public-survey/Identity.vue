<template>
  <div class="max-w-md mx-auto my-12 bg-white rounded-2xl border border-slate-100 shadow-lg p-6 lg:p-8 space-y-6 animate-fade-in">
    <!-- Icon & Header -->
    <div class="text-center space-y-2">
      <div class="mx-auto bg-[#4647AE]/10 text-[#4647AE] p-3 rounded-full w-12 h-12 flex items-center justify-center">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
          <path stroke-linecap="round" stroke-linejoin="round" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
        </svg>
      </div>
      <h2 class="text-xl font-bold text-slate-800 tracking-tight">Identitas Peserta Survei</h2>
      <p class="text-xs text-slate-400">Silakan pilih metode pengisian dan lengkapi data Anda untuk memulai survei kepuasan karyawan Laskar Buah.</p>
    </div>

    <!-- Form -->
    <form @submit.prevent="startSurvey" class="space-y-5">
      <!-- Mode Selection Toggle -->
      <div class="space-y-1.5">
        <label class="text-xs font-bold text-slate-500 uppercase tracking-wider block">Metode Pengisian</label>
        <div class="grid grid-cols-2 gap-2 p-1 bg-slate-100 rounded-xl">
          <button 
            type="button"
            @click="selectedMode = 'identity'"
            class="py-2 text-xs font-bold rounded-lg transition-all"
            :class="selectedMode === 'identity' ? 'bg-white text-[#4647AE] shadow-xs' : 'text-slate-500 hover:text-slate-700'"
          >
            Isi Identitas
          </button>
          <button 
            type="button"
            @click="selectedMode = 'anonymous'"
            class="py-2 text-xs font-bold rounded-lg transition-all"
            :class="selectedMode === 'anonymous' ? 'bg-white text-[#4647AE] shadow-xs' : 'text-slate-500 hover:text-slate-700'"
          >
            Anonim
          </button>
        </div>
      </div>

      <!-- Name Input (Hidden until Identity selected) -->
      <div v-if="selectedMode === 'identity'" class="space-y-1.5 animate-fade-in">
        <label class="text-xs font-bold text-slate-500 uppercase tracking-wider">Nama Lengkap</label>
        <input 
          v-model="employeeId"
          type="text" 
          placeholder="Contoh: Budi Santoso"
          required
          class="w-full border border-slate-200 rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:border-[#4647AE] transition-colors"
        />
      </div>

      <!-- Anonymous Notice (Shown only when Anonymous selected) -->
      <div v-if="selectedMode === 'anonymous'" class="p-3 bg-amber-50 border border-amber-100 rounded-xl text-xs text-amber-700 leading-relaxed animate-fade-in">
        <strong>Pemberitahuan:</strong> Identitas dan nama Anda tidak akan dicatat. Survei ini akan dikirimkan secara anonim untuk menjaga privasi Anda.
      </div>

      <!-- Department Input (Shown when a mode is selected) -->
      <div v-if="selectedMode" class="space-y-1.5 animate-fade-in relative z-20">
        <label class="text-xs font-bold text-slate-500 uppercase tracking-wider">Departemen / Divisi</label>
        <div class="relative" ref="dropdownRef">
          <!-- Dropdown Trigger -->
          <button
            type="button"
            @click="showDropdown = !showDropdown"
            class="w-full border rounded-xl px-4 py-2.5 text-sm text-left flex items-center justify-between transition-colors bg-white"
            :class="showDropdown ? 'border-[#4647AE] ring-1 ring-[#4647AE]/20' : 'border-slate-200'"
          >
            <span :class="department ? 'text-slate-700' : 'text-slate-400'">{{ department || 'Pilih departemen Anda' }}</span>
            <svg class="w-4 h-4 text-slate-400 transition-transform" :class="showDropdown ? 'rotate-180' : ''" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M19 9l-7 7-7-7" />
            </svg>
          </button>
          <!-- Hidden required input for form validation -->
          <input type="text" :value="department" required class="sr-only" tabindex="-1" />
          <!-- Dropdown Panel -->
          <div v-if="showDropdown" class="absolute z-50 mt-1.5 w-full bg-white border border-slate-200 rounded-xl shadow-lg overflow-hidden">
            <!-- Search Input -->
            <div class="p-2 border-b border-slate-100">
              <div class="relative">
                <svg class="absolute left-3 top-1/2 -translate-y-1/2 w-3.5 h-3.5 text-slate-400" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                </svg>
                <input
                  ref="searchInputRef"
                  v-model="deptSearch"
                  type="text"
                  placeholder="Cari departemen..."
                  class="w-full pl-8 pr-3 py-2 text-sm border border-slate-200 rounded-lg focus:outline-none focus:border-[#4647AE] transition-colors"
                  @click.stop
                />
              </div>
            </div>
            <!-- Options List -->
            <ul class="max-h-48 overflow-y-auto py-1">
              <li
                v-for="dept in filteredDepartments"
                :key="dept"
                @click="selectDepartment(dept)"
                class="px-4 py-2 text-sm cursor-pointer transition-colors"
                :class="department === dept ? 'bg-[#4647AE]/10 text-[#4647AE] font-semibold' : 'text-slate-700 hover:bg-slate-50'"
              >
                {{ dept }}
              </li>
              <li v-if="filteredDepartments.length === 0" class="px-4 py-3 text-xs text-slate-400 text-center">
                Tidak ditemukan
              </li>
            </ul>
          </div>
        </div>
      </div>

      <!-- Submit Button (Hidden until a mode is selected) -->
      <div v-if="selectedMode" class="pt-4 animate-fade-in relative z-10">
        <button 
          type="submit"
          class="w-full bg-[#4647AE] hover:bg-[#383994] active:bg-[#2e2e7a] text-white font-semibold py-3 px-4 rounded-xl transition-all shadow-sm shadow-indigo-200 flex items-center justify-center space-x-2"
        >
          <span>Mulai Isi Kuesioner</span>
          <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M14 5l7 7m0 0l-7 7m7-7H3" />
          </svg>
        </button>
      </div>
    </form>
  </div>
</template>

<script setup>
import { ref, computed, watch, nextTick, onMounted, onBeforeUnmount } from 'vue';
import { useRouter, useRoute } from 'vue-router';

const router = useRouter();
const route = useRoute();
const selectedMode = ref(null);
const employeeId = ref('');
const department = ref('');

// Searchable dropdown state
const showDropdown = ref(false);
const deptSearch = ref('');
const dropdownRef = ref(null);
const searchInputRef = ref(null);

const departments = [
  'HOLDING',
  'BENGKEL',
  'PURCHASING BUAH IMPOR',
  'DC BUAH',
  'MINUMAN',
  'ELECTRICAL',
  'LOGISTIK',
  'GUDANG KONSTRUKSI',
  'PENJUALAN',
  'BAKERY',
  'PURCHASING NON BUAH',
  'AKUNTING/FINANCE',
  'PURCHASING BUAH',
  'HRD & GA',
  'DC MAGETAN 1',
  'DC GRESIK',
  'PROYEK KONSTRUKSI',
  'GRAND OPENING',
  'DC PACKAGING',
  'DC SNACK',
  'DC GROSIR',
  'DC ROLLA',
  'SIPIL',
  'IT',
  'AMDK',
  'DC REMBANG 1',
  'R&D',
  'BRANCH MARKETING',
  'BRAM PRINTING',
  'KONVEKSI',
  'LEGAL FORMAL & TAX',
  'MARKETING',
  'DC JOMBANG 1',
  'SATPASPUR',
  'DC INDUK BUAH',
  'DUMP TRUCK',
  'LIMBAH',
  'DC PURWODADI',
  'PURCHASING SAYUR',
  'MAINTENANCE',
  'HSE',
  'DC KLATEN 1',
  'DC INDUK NON BUAH',
  'DAPUR LALA MIE',
  'MAKANAN',
  'DC REMBANG 2',
  'DC LAMONGAN',
  'DC FROZEN',
  'DC KLATEN 2',
  'DC JOMBANG 2',
  'DC MAGETAN 2',
  'SUPPLIER INTERNAL MAGETAN',
  'DC PURWODADI 2',
  'BRAM ELECTRIC',
  'LAINNYA',
  'PENJUALAN LALAMIE'
];

const filteredDepartments = computed(() => {
  if (!deptSearch.value) return departments;
  const q = deptSearch.value.toLowerCase();
  return departments.filter(d => d.toLowerCase().includes(q));
});

const selectDepartment = (dept) => {
  department.value = dept;
  showDropdown.value = false;
  deptSearch.value = '';
};

// Auto-focus search input when dropdown opens
watch(showDropdown, async (val) => {
  if (val) {
    await nextTick();
    searchInputRef.value?.focus();
  } else {
    deptSearch.value = '';
  }
});

// Close dropdown on outside click
const handleClickOutside = (e) => {
  if (dropdownRef.value && !dropdownRef.value.contains(e.target)) {
    showDropdown.value = false;
  }
};

onMounted(() => document.addEventListener('click', handleClickOutside));
onBeforeUnmount(() => document.removeEventListener('click', handleClickOutside));

const startSurvey = () => {
  const isAnon = selectedMode.value === 'anonymous';

  // Use 'ANONYMOUS' or the provided employee ID
  const finalId = isAnon 
    ? 'ANONYMOUS'
    : employeeId.value;

  const finalDept = department.value;

  const surveyId = route.query.survey_id || '';

  // Save credentials/identity to session storage for the survey form step
  sessionStorage.setItem('respondent_id', finalId);
  sessionStorage.setItem('respondent_dept', finalDept);
  sessionStorage.setItem('survey_id', surveyId);
  
  // Transition to survey questions page with survey_id query param
  router.push({ path: '/survey/form', query: { survey_id: surveyId } });
};
</script>

<style scoped>
.animate-fade-in {
  animation: fadeIn 0.4s ease-out forwards;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
