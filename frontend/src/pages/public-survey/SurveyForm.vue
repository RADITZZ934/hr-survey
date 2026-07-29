<template>
  <div class="max-w-2xl mx-auto my-8 bg-white rounded-2xl border border-slate-100 shadow-lg p-6 lg:p-8 space-y-8 animate-fade-in">
    <!-- Form Header & Progress (Sticky Header) -->
    <div class="sticky top-0 bg-white/95 backdrop-blur-md pt-2 pb-4 z-20 border-b border-slate-100 -mx-6 px-6 lg:-mx-8 lg:px-8 space-y-3.5 transition-all">
      <div class="flex justify-between items-center text-[10px] sm:text-xs font-bold text-slate-400 uppercase tracking-wider">
        <span>Kuesioner Kepuasan Karyawan Laskar Buah</span>
        <span v-if="questions.length > 0" class="text-[#4647AE] font-semibold bg-[#4647AE]/10 px-2 py-0.5 rounded-md">Terisi {{ answeredCount }} dari {{ questions.length }} Soal</span>
      </div>
      
      <!-- Progress Bar -->
      <div class="w-full bg-slate-100 h-2 rounded-full overflow-hidden">
        <div 
          class="bg-[#4647AE] h-full transition-all duration-300"
          :style="{ width: `${progressPercentage}%` }"
        ></div>
      </div>

      <!-- Section Indicator Header -->
      <div v-if="categories.length > 0" class="space-y-2">
        <div class="flex justify-between items-center bg-slate-50 p-2.5 rounded-xl border border-slate-100/50">
          <span class="text-xs font-bold text-slate-700">
            Bagian {{ activeCategoryIndex + 1 }} dari {{ categories.length }}: {{ activeCategory?.name }}
          </span>
          <span class="text-[10px] bg-[#4647AE]/10 text-[#4647AE] font-semibold px-2 py-0.5 rounded-full uppercase">
            {{ activeCategory?.questions.length }} Soal
          </span>
        </div>
        <!-- Star Legend Note -->
        <div v-if="activeCategory?.questions.some(q => q.type === 'star')" class="bg-slate-50 border border-slate-100 rounded-xl p-3 text-[11px] sm:text-xs text-slate-500 space-y-1.5 shadow-xs">
          <div class="font-bold text-slate-600">Petunjuk Pengisian Penilaian Bintang:</div>
          <div class="flex flex-wrap gap-x-6 gap-y-1 font-medium">
            <div class="flex items-center gap-1.5">
              <span class="bg-amber-400/10 text-amber-700 px-1.5 py-0.5 rounded text-[10px] font-bold">1</span>
              <span>Sangat tidak setuju</span>
            </div>
            <div class="flex items-center gap-1.5">
              <span class="bg-amber-400/10 text-amber-700 px-1.5 py-0.5 rounded text-[10px] font-bold">5</span>
              <span>Sangat setuju</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="py-12 text-center text-slate-400">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-[#4647AE] mx-auto mb-3"></div>
      <p class="text-sm">Memuat pertanyaan survei...</p>
    </div>

    <!-- Questions Page Content -->
    <div v-else-if="categories.length > 0" class="space-y-6">
      
      <!-- STEP: Questions by active category -->
      <div v-if="activeCategory" class="space-y-6">
        <div 
          v-for="(q, idx) in activeCategory.questions" 
          :key="q.id" 
          class="question-card p-6 border border-slate-100 rounded-2xl bg-white/60 backdrop-blur-xs space-y-4 hover:border-indigo-200/80 hover:shadow-md hover:bg-white transition-all duration-300 shadow-xs"
        >
          <div class="space-y-1">
            <h3 class="text-sm font-semibold text-slate-700 leading-relaxed">
              {{ idx + 1 }}. {{ q.text }}
              <span v-if="q.is_required" class="text-rose-500" title="Wajib Diisi">*</span>
            </h3>
          </div>

          <!-- Rating selection (Interactive animated 5-Star Rating) -->
          <div v-if="q.type === 'star'" class="flex items-center space-x-1.5 pt-2">
            <button 
              v-for="val in 5" 
              :key="val"
              type="button"
              @click="selectScore(q.id, val)"
              @mouseenter="hoveredRatings[q.id] = val"
              @mouseleave="hoveredRatings[q.id] = 0"
              class="transition-all duration-200 transform hover:scale-120 active:scale-95 focus:outline-none group"
            >
              <svg 
                xmlns="http://www.w3.org/2000/svg" 
                viewBox="0 0 24 24" 
                fill="currentColor" 
                class="star-icon w-10 h-10 transition-all duration-200 ease-out"
                :class="isStarActive(q, val) ? 'text-amber-400 drop-shadow-[0_0_6px_rgba(251,191,36,0.6)] scale-110' : 'text-slate-200 group-hover:text-slate-300'"
              >
                <path fill-rule="evenodd" d="M10.788 3.21c.448-1.077 1.976-1.077 2.424 0l2.082 5.006 5.404.434c1.164.093 1.636 1.545.749 2.305l-4.117 3.527 1.257 5.273c.271 1.136-.964 2.033-1.96 1.425L12 18.354 7.373 21.18c-.996.608-2.231-.29-1.96-1.425l1.257-5.273-4.117-3.527c-.887-.76-.415-2.212.749-2.305l5.404-.434 2.082-5.005Z" clip-rule="evenodd" />
              </svg>
            </button>
          </div>

          <!-- Essay Text Area -->
          <div v-else class="pt-2">
            <textarea
              v-model="answers[q.id]"
              @input="autoResize"
              rows="1"
              placeholder="Tuliskan jawaban Anda di sini..."
              class="w-full bg-white border border-slate-200 rounded-xl px-4 py-3 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500/20 focus:border-[#4647AE] transition-all resize-none text-slate-700 overflow-hidden"
              style="min-height: 48px; height: auto;"
            ></textarea>
          </div>
        </div>
      </div>
    </div>

    <!-- Empty State -->
    <div v-else class="py-12 text-center text-slate-400">
      <p class="text-sm">Tidak ada pertanyaan yang tersedia pada survei ini.</p>
    </div>

    <!-- Footer Actions -->
    <div class="flex gap-4 pt-4 border-t border-slate-100">
      <!-- Back button -->
      <button 
        v-if="activeCategoryIndex > 0"
        @click="activeCategoryIndex--"
        class="px-5 py-3 border border-slate-200 text-slate-600 rounded-xl hover:bg-slate-50 text-sm font-semibold transition-colors focus:outline-none"
      >
        Kembali
      </button>

      <!-- Next to next category button -->
      <button 
        v-if="activeCategoryIndex < categories.length - 1"
        @click="activeCategoryIndex++"
        :disabled="isActiveCategoryDisabled || loading || questions.length === 0"
        class="w-full bg-[#4647AE] hover:bg-[#383994] active:bg-[#2e2e7a] text-white font-semibold py-3.5 px-4 rounded-xl transition-all shadow-sm shadow-indigo-200 flex items-center justify-center space-x-2 disabled:opacity-40 disabled:cursor-not-allowed"
      >
        <span>Lanjutkan</span>
        <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
          <path stroke-linecap="round" stroke-linejoin="round" d="M9 5l7 7-7 7" />
        </svg>
      </button>

      <!-- Submit Response Button -->
      <button 
        v-else
        @click="submitSurvey"
        :disabled="isSubmitDisabled || loading || submitting || questions.length === 0"
        class="w-full bg-emerald-600 hover:bg-emerald-700 active:bg-emerald-800 text-white font-semibold py-3.5 px-4 rounded-xl transition-all shadow-sm shadow-emerald-200 flex items-center justify-center space-x-2 disabled:opacity-40 disabled:cursor-not-allowed"
      >
        <span>{{ submitting ? 'Mengirim...' : 'Kirim Kuesioner' }}</span>
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { getSurveyQuestions, submitSurveyResponse } from '../../services/survey.service';

const router = useRouter();
const route = useRoute();

const questions = ref([]);
const loading = ref(true);
const submitting = ref(false); // Controls submit button loading state
const answers = ref({});
const hoveredRatings = ref({}); // Maps question ID to active hover score
const activeCategoryIndex = ref(0);

const autoResize = (event) => {
  const el = event.target;
  el.style.height = 'auto';
  el.style.height = `${el.scrollHeight}px`;
};

// Group questions by category name
const categories = computed(() => {
  const map = {};
  questions.value.forEach(q => {
    const catName = q.category || (q.type === 'essay' ? 'Essay' : 'Rating Bintang');
    if (!map[catName]) {
      map[catName] = {
        name: catName,
        questions: []
      };
    }
    map[catName].questions.push(q);
  });
  
  // Sort categories: put categories with star questions first, essay only categories last
  return Object.values(map).sort((a, b) => {
    const aHasStar = a.questions.some(q => q.type === 'star');
    const bHasStar = b.questions.some(q => q.type === 'star');
    if (aHasStar && !bHasStar) return -1;
    if (!aHasStar && bHasStar) return 1;
    return a.name.localeCompare(b.name);
  });
});

const activeCategory = computed(() => {
  return categories.value[activeCategoryIndex.value] || null;
});

const isStarActive = (q, val) => {
  const hoverVal = hoveredRatings.value[q.id] || 0;
  if (hoverVal > 0) {
    return val <= hoverVal;
  }
  return val <= (answers.value[q.id] || 0);
};

const answeredCount = computed(() => {
  return questions.value.filter(q => {
    const ans = answers.value[q.id];
    if (q.type === 'essay') {
      return ans !== undefined && ans.trim() !== '';
    } else {
      return ans !== undefined;
    }
  }).length;
});

const progressPercentage = computed(() => {
  if (questions.value.length === 0) return 0;
  return (answeredCount.value / questions.value.length) * 100;
});

// Check if any required question in the active category is unanswered
const isActiveCategoryDisabled = computed(() => {
  if (!activeCategory.value) return false;
  return activeCategory.value.questions.some(q => {
    if (!q.is_required) return false;
    const ans = answers.value[q.id];
    if (q.type === 'essay') {
      return !ans || !ans.trim();
    } else {
      return ans === undefined;
    }
  });
});

// Check if any required question is unanswered in the entire survey
const isSubmitDisabled = computed(() => {
  return questions.value.some(q => {
    if (!q.is_required) return false;
    const ans = answers.value[q.id];
    if (q.type === 'essay') {
      return !ans || !ans.trim();
    } else {
      return ans === undefined;
    }
  });
});

onMounted(async () => {
  const surveyId = route.query.survey_id || sessionStorage.getItem('survey_id');
  if (surveyId) {
    try {
      const res = await getSurveyQuestions(surveyId);
      questions.value = res.data.map(q => ({
        id: q.id,
        text: q.text,
        category: q.category ? q.category.name : (q.category_id === 1 ? 'Rating Bintang' : 'Essay'),
        type: q.type === 'text' ? 'essay' : 'star',
        is_required: q.is_required
      }));
      
      // Initialize hover states for questions
      questions.value.forEach(q => {
        hoveredRatings.value[q.id] = 0;
      });
    } catch (error) {
      console.error('Failed to fetch survey questions:', error);
    } finally {
      loading.value = false;
    }
  } else {
    loading.value = false;
  }
});

const selectScore = (qId, score) => {
  answers.value[qId] = score;
};

const submitSurvey = async () => {
  const surveyId = route.query.survey_id || sessionStorage.getItem('survey_id');
  if (!surveyId) return;

  submitting.value = true;
  try {
    // Clean score JSON generation - just pass score as number or null in JS
    const payloadAnswers = questions.value.map(q => {
      const ans = answers.value[q.id];
      if (q.type === 'essay') {
        return {
          question_id: q.id,
          answer_text: ans || ""
        };
      } else {
        return {
          question_id: q.id,
          score: ans !== undefined ? Number(ans) : null
        };
      }
    });

    const payload = {
      respondent_id: sessionStorage.getItem('respondent_id') || 'ANONYMOUS',
      respondent_dept: sessionStorage.getItem('respondent_dept') || 'ANONYMOUS',
      answers: payloadAnswers
    };

    const res = await submitSurveyResponse(surveyId, payload);

    if (res.data?.data) {
      sessionStorage.setItem('survey_result', JSON.stringify(res.data.data));
    }

    // Clear cached respondent info after completing
    sessionStorage.removeItem('respondent_id');
    sessionStorage.removeItem('respondent_dept');
    sessionStorage.removeItem('survey_id');
    
    router.push('/survey/thanks');
  } catch (error) {
    console.error('Failed to submit survey answers:', error);
    alert('Gagal mengirimkan survey. Silakan coba kembali.');
  } finally {
    submitting.value = false;
  }
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

.question-card {
  content-visibility: auto;
  contain-intrinsic-size: auto 150px;
}

.star-icon {
  will-change: transform;
  transform: translate3d(0, 0, 0);
}
</style>
