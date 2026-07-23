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
        <label class="text-xs font-bold text-slate-500 uppercase tracking-wider">Nama Lengkap / NIK</label>
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

      <!-- Department Input (Hidden until Identity selected) -->
      <div v-if="selectedMode === 'identity'" class="space-y-1.5 animate-fade-in">
        <label class="text-xs font-bold text-slate-500 uppercase tracking-wider">Departemen / Divisi</label>
        <select 
          v-model="department"
          required
          class="w-full border border-slate-200 rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:border-[#4647AE] transition-colors text-slate-700 bg-white"
        >
          <option value="" disabled>Pilih departemen Anda</option>
          <option v-for="dept in departments" :key="dept" :value="dept">{{ dept }}</option>
        </select>
      </div>

      <!-- Submit Button (Hidden until a mode is selected) -->
      <div v-if="selectedMode" class="pt-4 animate-fade-in">
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
import { ref } from 'vue';
import { useRouter, useRoute } from 'vue-router';

const router = useRouter();
const route = useRoute();
const selectedMode = ref(null); // 'identity', 'anonymous', or null
const employeeId = ref('');
const department = ref('');

const departments = [
  'AMDK',
  'BENGKEL',
  'BRAM PRINTING',
  'DISTRIBUTOR CENTER',
  'DISTRIBUTOR CENTER RECYCLE',
  'GYM SPORT CENTER',
  'HOLDING',
  'Information Technology',
  'KONSTRUKSI AND BRANDING',
  'KONVEKSI',
  'LEGAL FORMAL & TAX',
  'LIMBAH',
  'LOGISTIK',
  'MAIN DC',
  'MARKETING',
  'PACKING',
  'PENJUALAN',
  'PERTANIAN',
  'PRODUKSI',
  'PURCHASING'
];

const startSurvey = () => {
  const isAnon = selectedMode.value === 'anonymous';

  // Use 'ANONYMOUS' or the provided employee ID
  const finalId = isAnon 
    ? 'ANONYMOUS'
    : employeeId.value;

  const finalDept = isAnon
    ? 'ANONYMOUS'
    : department.value;

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
