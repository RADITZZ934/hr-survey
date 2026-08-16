<template>
  <div 
    @click="handlePageClick"
    class="min-h-screen w-full bg-[#f0edf9] flex items-center justify-center p-4 sm:p-6 lg:p-8 font-sans relative overflow-hidden cursor-pointer select-none"
  >
    <!-- Initial Page Load Celebration (Plays once) -->
    <div v-if="showInitialCelebration && celebrateDataUrl" class="fixed inset-0 pointer-events-none z-50 flex items-center justify-center">
      <lottie-player 
        :src="celebrateDataUrl" 
        background="transparent" 
        speed="1.2" 
        style="width: 100%; height: 100%; object-fit: cover;" 
        autoplay
      ></lottie-player>
    </div>

    <!-- Small Click-based Celebration Instances -->
    <div 
      v-for="click in clicks" 
      :key="click.id"
      class="fixed pointer-events-none z-50 transform -translate-x-1/2 -translate-y-1/2"
      :style="{ left: `${click.x}px`, top: `${click.y}px`, width: '150px', height: '150px' }"
    >
      <lottie-player 
        v-if="celebrateDataUrl"
        :src="celebrateDataUrl" 
        background="transparent" 
        speed="1.5" 
        style="width: 100%; height: 100%;" 
        autoplay
      ></lottie-player>
    </div>

    <div class="max-w-lg w-full bg-white rounded-3xl border border-slate-100 shadow-xl p-8 sm:p-10 text-center space-y-7 animate-fade-in relative z-20">
      
      <!-- App Brand & Logo Header -->
      <div class="flex flex-col items-center justify-center space-y-3">
        <img :src="isBazzar ? '/bz-icon.png' : '/laskar-corps.png'" class="h-24 w-auto object-contain" alt="HR Survey Icon" />
        <span class="text-base font-extrabold text-slate-800 tracking-tight">
          {{ 
            surveyVisibility === 'external' 
              ? (isBazzar ? 'LAYANAN KEPUASAN CUSTOMER | BAZZAR' : 'LAYANAN KEPUASAN CUSTOMER | LASKAR BUAH') 
              : (isBazzar ? 'HR SURVEY TOOLS | BAZZAR' : 'HR SURVEY TOOLS | LASKAR BUAH') 
          }}
        </span>
      </div>

            <!-- Score Result Summary Card -->
      <div v-if="resultData" class="p-6 bg-slate-50 border border-slate-100 rounded-2xl space-y-4 shadow-xs">
        <div class="text-xs font-bold text-slate-400 uppercase tracking-wider">
          Hasil Skor Penilaian Anda
        </div>

        <div class="flex items-center justify-center space-x-3">
          <span class="text-3xl sm:text-4xl font-extrabold text-slate-800">
            ★ {{ resultData.score ? resultData.score.toFixed(1) : '5.0' }}
          </span>
          <span class="text-xs font-semibold text-slate-400 self-end mb-1">
            / 5.0
          </span>
        </div>

        <!-- Progress Gauge Bar -->
        <div class="space-y-1.5">
          <div class="w-full bg-slate-200 rounded-full h-3 overflow-hidden">
            <div 
              class="h-3 rounded-full transition-all duration-1000 ease-out"
              :class="getProgressBarClass(resultData.score)"
              :style="{ width: `${resultData.percentage || 100}%` }"
            ></div>
          </div>
          <div class="flex justify-between items-center text-[11px] font-semibold text-slate-500">
            <span>Tingkat Kepuasan ({{ resultData.percentage || 100 }}%)</span>
            <span class="px-2.5 py-0.5 rounded-md text-xs font-bold" :class="getCategoryBadgeClass(resultData.category)">
              {{ resultData.category || 'Sangat Puas' }}
            </span>
          </div>
        </div>
      </div>


      <!-- Thank You Message -->
      <div class="space-y-2">
        <h1 class="text-2xl sm:text-3xl font-extrabold text-slate-800 tracking-tight">
          Terima kasih atas partisipasinya!
        </h1>
        <p class="text-xs sm:text-sm text-slate-400 leading-relaxed max-w-md mx-auto">
          <template v-if="surveyVisibility === 'external'">
            Tanggapan dan masukan Anda telah berhasil direkam. Penilaian Anda sangat berharga dalam meningkatkan kualitas pelayanan dan produk di {{ isBazzar ? 'Bazzar' : 'Laskar Buah' }}.
          </template>
          <template v-else>
            Tanggapan dan masukan Anda telah berhasil direkam. Penilaian Anda sangat berharga dalam meningkatkan kualitas dan kenyamanan kerja di {{ isBazzar ? 'Bazzar' : 'Laskar Buah' }}.
          </template>
        </p>
      </div>




    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';

const CELEBRATE_CACHE_KEY = 'celebrate_json_cache';

const resultData = ref(null);
const surveyVisibility = ref(sessionStorage.getItem('survey_visibility') || 'internal');
const showInitialCelebration = ref(true);
const clicks = ref([]);
const celebrateDataUrl = ref(null);
const isBazzar = ref(false);

// Create a Blob URL from raw JSON text for lottie-player
const createBlobUrl = (jsonText) => {
  const blob = new Blob([jsonText], { type: 'application/json' });
  return URL.createObjectURL(blob);
};

// Load celebrate.json from localStorage cache, or fetch once and cache
const loadCelebrateAnimation = async () => {
  try {
    const cached = localStorage.getItem(CELEBRATE_CACHE_KEY);
    if (cached) {
      celebrateDataUrl.value = createBlobUrl(cached);
      return;
    }
    const res = await fetch('/celebrate.json');
    const text = await res.text();
    // Try to cache in localStorage for future visits
    try {
      localStorage.setItem(CELEBRATE_CACHE_KEY, text);
    } catch (storageErr) {
      // QuotaExceededError — skip caching, still works this session
      console.warn('localStorage quota exceeded, celebrate.json will not be cached:', storageErr);
    }
    celebrateDataUrl.value = createBlobUrl(text);
  } catch (e) {
    console.error('Failed to load celebrate animation:', e);
    // Fallback: just use the URL directly (no caching)
    celebrateDataUrl.value = '/celebrate.json';
  }
};

const handlePageClick = (e) => {
  const id = Date.now() + Math.random();
  clicks.value.push({
    id,
    x: e.clientX,
    y: e.clientY
  });
  // Auto-clean click items after the animation finishes
  setTimeout(() => {
    clicks.value = clicks.value.filter(click => click.id !== id);
  }, 1500);
};

onMounted(async () => {
  // Determine if survey is for Bazzar banner
  isBazzar.value = window.location.href.toLowerCase().includes('bazzar') || 
                   window.location.href.toLowerCase().includes('bazaar') || 
                   sessionStorage.getItem('is_bazzar') === 'true';

  const visibility = sessionStorage.getItem('survey_visibility') || 'internal';
  surveyVisibility.value = visibility;
  if (visibility === 'external') {
    document.title = isBazzar.value ? 'LAYANAN KEPUASAN CUSTOMER | BAZZAR' : 'LAYANAN KEPUASAN CUSTOMER | LASKAR BUAH';
  } else {
    document.title = isBazzar.value ? 'HR SURVEY TOOLS | BAZZAR' : 'HR SURVEY TOOLS | LASKAR BUAH';
  }

  // Clean up sessionStorage
  sessionStorage.removeItem('is_bazzar');
  sessionStorage.removeItem('survey_visibility');

  // Load celebrate animation from cache or network
  await loadCelebrateAnimation();

  // Hide initial big celebration after 3 seconds
  setTimeout(() => {
    showInitialCelebration.value = false;
  }, 3000);

  const cachedResult = sessionStorage.getItem('survey_result');
  if (cachedResult) {
    try {
      resultData.value = JSON.parse(cachedResult);
    } catch (e) {
      console.error('Failed to parse survey_result:', e);
    }
  } else {
    // Default fallback mock data for direct preview
    resultData.value = {
      score: 4.5,
      max_score: 5.0,
      percentage: 90,
      category: 'Sangat Puas'
    };
  }
});

const getProgressBarClass = (score) => {
  if (score >= 4.5) return 'bg-emerald-500';
  if (score >= 3.5) return 'bg-blue-500';
  if (score >= 2.5) return 'bg-amber-500';
  return 'bg-rose-500';
};

const getCategoryBadgeClass = (category) => {
  if (category === 'Sangat Puas') return 'bg-emerald-50 text-emerald-700 border border-emerald-200';
  if (category === 'Puas') return 'bg-blue-50 text-blue-700 border border-blue-200';
  if (category === 'Cukup') return 'bg-amber-50 text-amber-700 border border-amber-200';
  return 'bg-rose-50 text-rose-700 border border-rose-200';
};
</script>

<style scoped>
.animate-fade-in {
  animation: fadeIn 0.4s ease-out forwards;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(12px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
