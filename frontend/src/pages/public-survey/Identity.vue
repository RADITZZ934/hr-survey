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

    <!-- Error Notice Block -->
    <div v-if="surveyError" class="p-5 bg-rose-50 border border-rose-100 rounded-2xl text-center space-y-3">
      <div class="inline-flex p-2 bg-rose-100 text-rose-600 rounded-full">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
          <path stroke-linecap="round" stroke-linejoin="round" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
        </svg>
      </div>
      <div class="space-y-1 text-center">
        <h3 class="font-bold text-slate-800 text-sm">Survei Tidak Tersedia</h3>
        <p class="text-xs text-slate-500 leading-relaxed max-w-xs mx-auto">
          <span v-if="surveyError === 'expired'">Maaf, masa pengisian kuesioner survei ini telah berakhir (expired).</span>
          <span v-else-if="surveyError === 'not_started'">Survei ini belum dimulai.</span>
          <span v-else-if="surveyError === 'not_found'">Survei tidak ditemukan. Silakan periksa kembali tautan Anda.</span>
          <span v-else>Maaf, kuesioner survei ini sedang dinonaktifkan atau tidak tersedia saat ini.</span>
        </p>
      </div>
    </div>

    <!-- Form -->
    <form v-else @submit.prevent="startSurvey" class="space-y-5">
      <!-- Mode Selection Toggle -->
      <div v-if="surveyVisibility === 'internal'" class="space-y-1.5">
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

      <!-- Domicile Selection (Shown when survey is external) -->
      <div v-if="surveyVisibility === 'external' && selectedMode === 'identity'" class="space-y-4 animate-fade-in relative z-20">
        <!-- Province Dropdown -->
        <div class="space-y-1.5 relative z-20" ref="provinceRef">
          <label class="text-xs font-bold text-slate-500 uppercase tracking-wider block">Provinsi</label>
          <button
            type="button"
            @click="toggleProvinceDropdown"
            class="w-full border rounded-xl px-4 py-2.5 text-sm text-left flex items-center justify-between transition-colors bg-white"
            :class="showProvinceDropdown ? 'border-[#4647AE] ring-1 ring-[#4647AE]/20' : 'border-slate-200'"
          >
            <span :class="selectedProvince ? 'text-slate-700' : 'text-slate-400'">{{ selectedProvince?.name || 'Pilih Provinsi' }}</span>
            <svg class="w-4 h-4 text-slate-400 transition-transform" :class="showProvinceDropdown ? 'rotate-180' : ''" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
               <path stroke-linecap="round" stroke-linejoin="round" d="M19 9l-7 7-7-7" />
            </svg>
          </button>
          <input type="text" :value="selectedProvince?.name || ''" required class="sr-only" tabindex="-1" />
          
          <div v-if="showProvinceDropdown" class="absolute z-50 mt-1.5 w-full bg-white border border-slate-200 rounded-xl shadow-lg overflow-hidden">
            <!-- Search Province -->
            <div class="p-2 border-b border-slate-100">
              <input
                v-model="provinceSearch"
                type="text"
                placeholder="Cari provinsi..."
                class="w-full px-3 py-2 text-sm border border-slate-200 rounded-lg focus:outline-none focus:border-[#4647AE] transition-colors"
                @click.stop
              />
            </div>
            <ul class="max-h-48 overflow-y-auto py-1">
              <li
                v-for="prov in filteredProvinces"
                :key="prov.id"
                @click="selectProvince(prov)"
                class="px-4 py-2 text-sm cursor-pointer transition-colors"
                :class="selectedProvince?.id === prov.id ? 'bg-[#4647AE]/10 text-[#4647AE] font-semibold' : 'text-slate-700 hover:bg-slate-50'"
              >
                {{ prov.name }}
              </li>
            </ul>
          </div>
        </div>

        <!-- Regency Dropdown -->
        <div class="space-y-1.5 relative z-10" ref="regencyRef">
          <label class="text-xs font-bold text-slate-500 uppercase tracking-wider block">Kabupaten / Kota</label>
          <button
            type="button"
            @click="toggleRegencyDropdown"
            :disabled="!selectedProvince"
            class="w-full border rounded-xl px-4 py-2.5 text-sm text-left flex items-center justify-between transition-colors bg-white disabled:opacity-50 disabled:bg-slate-50 disabled:cursor-not-allowed"
            :class="showRegencyDropdown ? 'border-[#4647AE] ring-1 ring-[#4647AE]/20' : 'border-slate-200'"
          >
            <span :class="selectedRegency ? 'text-slate-700' : 'text-slate-400'">{{ selectedRegency?.name || 'Pilih Kabupaten/Kota' }}</span>
            <svg class="w-4 h-4 text-slate-400 transition-transform" :class="showRegencyDropdown ? 'rotate-180' : ''" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
               <path stroke-linecap="round" stroke-linejoin="round" d="M19 9l-7 7-7-7" />
            </svg>
          </button>
          <input type="text" :value="selectedRegency?.name || ''" required class="sr-only" tabindex="-1" />
          
          <div v-if="showRegencyDropdown" class="absolute z-50 mt-1.5 w-full bg-white border border-slate-200 rounded-xl shadow-lg overflow-hidden">
            <!-- Search Regency -->
            <div class="p-2 border-b border-slate-100">
              <input
                v-model="regencySearch"
                type="text"
                placeholder="Cari kabupaten/kota..."
                class="w-full px-3 py-2 text-sm border border-slate-200 rounded-lg focus:outline-none focus:border-[#4647AE] transition-colors"
                @click.stop
              />
            </div>
            <ul class="max-h-48 overflow-y-auto py-1">
              <li
                v-for="reg in filteredRegencies"
                :key="reg.id"
                @click="selectRegency(reg)"
                class="px-4 py-2 text-sm cursor-pointer transition-colors"
                :class="selectedRegency?.id === reg.id ? 'bg-[#4647AE]/10 text-[#4647AE] font-semibold' : 'text-slate-700 hover:bg-slate-50'"
              >
                {{ reg.name }}
              </li>
            </ul>
          </div>
        </div>

        <!-- Store / Toko Dropdown (External Survey) -->
        <div class="space-y-1.5 relative" ref="storeRef">
          <label class="text-xs font-bold text-slate-500 uppercase tracking-wider block">Toko / Lokasi Survey</label>
          <button
            type="button"
            :disabled="hasStoreFromUrl"
            @click="showStoreDropdown = !showStoreDropdown"
            class="w-full border rounded-xl px-4 py-2.5 text-sm text-left flex items-center justify-between transition-colors"
            :class="[
              showStoreDropdown ? 'border-[#4647AE] ring-1 ring-[#4647AE]/20' : 'border-slate-200',
              hasStoreFromUrl ? 'bg-slate-50 text-slate-500 cursor-not-allowed border-slate-200' : 'bg-white'
            ]"
          >
            <span :class="selectedStore ? 'text-slate-700 font-medium' : 'text-slate-400'">
              {{ selectedStore ? (selectedStore.name || selectedStore.nama_store || String(selectedStore.id_store || selectedStore.id)) : 'Pilih Toko' }}
            </span>
            <svg v-if="!hasStoreFromUrl" class="w-4 h-4 text-slate-400 transition-transform" :class="showStoreDropdown ? 'rotate-180' : ''" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
               <path stroke-linecap="round" stroke-linejoin="round" d="M19 9l-7 7-7-7" />
            </svg>
          </button>
          <input type="text" :value="selectedStore ? '1' : ''" required class="sr-only" tabindex="-1" />

          <div v-if="showStoreDropdown" class="absolute z-50 mt-1.5 w-full bg-white border border-slate-200 rounded-xl shadow-lg overflow-hidden">
            <div class="p-2 border-b border-slate-100">
              <input
                v-model="storeSearch"
                type="text"
                placeholder="Cari toko..."
                class="w-full px-3 py-2 text-sm border border-slate-200 rounded-lg focus:outline-none focus:border-[#4647AE] transition-colors"
                @click.stop
              />
            </div>
            <ul class="max-h-48 overflow-y-auto py-1">
              <li v-if="filteredStores.length === 0" class="px-4 py-3 text-xs text-slate-400 text-center">
                {{ stores.length === 0 ? 'Memuat data toko...' : 'Toko tidak ditemukan' }}
              </li>
              <li
                v-for="store in filteredStores"
                :key="store.id_store || store.id"
                @click="selectedStore = store; showStoreDropdown = false; storeSearch = ''"
                class="px-4 py-2 text-sm cursor-pointer transition-colors"
                :class="(selectedStore?.id_store || selectedStore?.id) === (store.id_store || store.id) ? 'bg-[#4647AE]/10 text-[#4647AE] font-semibold' : 'text-slate-700 hover:bg-slate-50'"
              >
                {{ store.name || store.nama_store || store.id_store || store.id }}
              </li>
            </ul>
          </div>
        </div>
      </div>

      <!-- Department Input (Shown when a mode is selected for internal) -->
      <div v-if="selectedMode && surveyVisibility === 'internal'" class="space-y-1.5 animate-fade-in relative z-20">
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
import { getSurveyDetail, getProvinces, getRegencies } from '../../services/survey.service';
import { inventoryService } from '../../services/inventoryService';

const router = useRouter();
const route = useRoute();
const selectedMode = ref(null);
const employeeId = ref('');
const department = ref('');
const surveyError = ref(''); // '', 'expired', 'not_started', 'inactive', 'not_found'

const surveyVisibility = ref('internal');
const provinces = ref([]);
const regencies = ref([]);
const selectedProvince = ref(null);
const selectedRegency = ref(null);
const showProvinceDropdown = ref(false);
const showRegencyDropdown = ref(false);
const provinceSearch = ref('');
const regencySearch = ref('');
const provinceRef = ref(null);
const regencyRef = ref(null);

// Store selection state (untuk external survey)
const stores = ref([]);
const selectedStore = ref(null);
const storeSearch = ref('');
const showStoreDropdown = ref(false);
const storeRef = ref(null);

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

const filteredProvinces = computed(() => {
  if (!provinceSearch.value) return provinces.value;
  const q = provinceSearch.value.toLowerCase();
  return provinces.value.filter(p => p.name.toLowerCase().includes(q));
});

const filteredRegencies = computed(() => {
  if (!regencySearch.value) return regencies.value;
  const q = regencySearch.value.toLowerCase();
  return regencies.value.filter(r => r.name.toLowerCase().includes(q));
});

const hasStoreFromUrl = computed(() => !!route.query.store_id);

const filteredStores = computed(() => {
  if (hasStoreFromUrl.value && selectedStore.value) {
    return [selectedStore.value];
  }
  if (!storeSearch.value) return stores.value;
  const q = storeSearch.value.toLowerCase();
  return stores.value.filter(s =>
    (s.name || s.nama_store || '').toLowerCase().includes(q) ||
    String(s.id_store || s.id || '').toLowerCase().includes(q)
  );
});

const toggleProvinceDropdown = () => {
  showProvinceDropdown.value = !showProvinceDropdown.value;
  showRegencyDropdown.value = false;
  showStoreDropdown.value = false;
};

const toggleRegencyDropdown = () => {
  if (!selectedProvince.value) return;
  showRegencyDropdown.value = !showRegencyDropdown.value;
  showProvinceDropdown.value = false;
};

const selectProvince = async (prov) => {
  selectedProvince.value = prov;
  selectedRegency.value = null;
  showProvinceDropdown.value = false;
  provinceSearch.value = '';
  try {
    const res = await getRegencies(prov.id);
    regencies.value = res.data;
  } catch (err) {
    console.error('Failed to fetch regencies:', err);
  }
};

const selectRegency = (reg) => {
  selectedRegency.value = reg;
  showRegencyDropdown.value = false;
  regencySearch.value = '';
};

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
  if (provinceRef.value && !provinceRef.value.contains(e.target)) {
    showProvinceDropdown.value = false;
  }
  if (regencyRef.value && !regencyRef.value.contains(e.target)) {
    showRegencyDropdown.value = false;
  }
  if (storeRef.value && !storeRef.value.contains(e.target)) {
    showStoreDropdown.value = false;
  }
};

onMounted(async () => {
  document.addEventListener('click', handleClickOutside);
  const surveyId = route.query.survey_id;
  if (surveyId) {
    try {
      const res = await getSurveyDetail(surveyId);
      const survey = res.data;
      
      // Perform timeline and status validation
      const now = new Date();
      const startDate = new Date(survey.start_date);
      const endDate = new Date(survey.end_date);
      // Set end date to the very end of that day (23:59:59)
      endDate.setHours(23, 59, 59, 999);

      if (survey.status !== 'active') {
        surveyError.value = 'inactive';
      } else if (now < startDate) {
        surveyError.value = 'not_started';
      } else if (now > endDate) {
        surveyError.value = 'expired';
      }

      surveyVisibility.value = survey.visibility || 'internal';
      if (surveyVisibility.value === 'external') {
        selectedMode.value = 'identity';
        // Fetch provinces
        const provRes = await getProvinces();
        provinces.value = provRes.data;
        // Fetch store list dari inventory
        try {
          stores.value = await inventoryService.getLocations();
        } catch (err) {
          console.error('Failed to fetch stores:', err);
        }
        // Pre-fill toko jika store_id sudah ada di URL (embedded dari QR code)
        const storeIdFromUrl = route.query.store_id;
        if (storeIdFromUrl) {
          const decodedStoreId = decodeURIComponent(storeIdFromUrl);
          const found = stores.value.find(s => 
            String(s.id_store || s.id) === String(decodedStoreId) ||
            String(s.name || s.nama_store || '').toLowerCase().trim() === String(decodedStoreId).toLowerCase().trim()
          );
          selectedStore.value = found || { id_store: decodedStoreId, name: decodedStoreId };
        }
      }
    } catch (err) {
      console.error('Failed to fetch survey details:', err);
      surveyError.value = 'not_found';
    }
  } else {
    surveyError.value = 'not_found';
  }
});

onBeforeUnmount(() => document.removeEventListener('click', handleClickOutside));

const startSurvey = () => {
  const isAnon = selectedMode.value === 'anonymous';

  // Use 'ANONYMOUS' or the provided employee ID
  const finalId = isAnon 
    ? 'ANONYMOUS'
    : employeeId.value;

  const surveyId = route.query.survey_id || '';

  // Save credentials/identity to session storage for the survey form step
  sessionStorage.setItem('respondent_id', finalId);
  sessionStorage.setItem('survey_id', surveyId);

  if (surveyVisibility.value === 'external') {
    if (!selectedStore.value) {
      alert('Silakan pilih toko terlebih dahulu.');
      return;
    }
    sessionStorage.setItem('respondent_province', selectedProvince.value ? selectedProvince.value.name : '');
    sessionStorage.setItem('respondent_regency', selectedRegency.value ? selectedRegency.value.name : '');
    const storeName = String(selectedStore.value.name || selectedStore.value.id_store || selectedStore.value.id || '');
    sessionStorage.setItem('respondent_dept', storeName);
    sessionStorage.setItem('id_store', storeName);
    sessionStorage.setItem('survey_visibility', 'external');
  } else {
    sessionStorage.setItem('respondent_dept', department.value);
    sessionStorage.setItem('respondent_province', '');
    sessionStorage.setItem('respondent_regency', '');
    sessionStorage.setItem('survey_visibility', 'internal');
  }
  
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
