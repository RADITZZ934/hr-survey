<script setup>
import { ref, onMounted, provide, watch } from 'vue';
import { useRoute } from 'vue-router';
import { getSurveyDetail } from '../services/survey.service';

const isBazzar = ref(false);
const surveyTitle = ref(sessionStorage.getItem('survey_title') || 'Employee Satisfaction Survey');
const route = useRoute();

const fetchSurveyTitle = async (surveyId) => {
  if (!surveyId) return;
  try {
    const res = await getSurveyDetail(surveyId);
    if (res.data && res.data.title) {
      surveyTitle.value = res.data.title;
      sessionStorage.setItem('survey_title', res.data.title);
    }
  } catch (err) {
    console.error('Failed to fetch survey details in layout:', err);
  }
};

provide('setSurveyTitle', (title) => {
  surveyTitle.value = title;
  sessionStorage.setItem('survey_title', title);
});

onMounted(() => {
  isBazzar.value = window.location.href.toLowerCase().includes('bazzar') || 
                   window.location.href.toLowerCase().includes('bazaar') || 
                   sessionStorage.getItem('is_bazzar') === 'true';

  const surveyId = route.query.survey_id || sessionStorage.getItem('survey_id');
  if (surveyId) {
    fetchSurveyTitle(surveyId);
  }
});

watch(() => route.query.survey_id, (newSurveyId) => {
  if (newSurveyId) {
    fetchSurveyTitle(newSurveyId);
  }
});
</script>

<template>
  <div class="min-h-screen bg-gradient-to-tr from-slate-100 via-slate-50 to-blue-50/50 flex flex-col relative">
    <!-- Ambient Background Blur Orbs Container -->
    <div class="absolute inset-0 overflow-hidden pointer-events-none z-0">
      <div class="absolute top-[-10%] left-[-10%] w-[600px] h-[600px] rounded-full bg-blue-400/15 blur-[120px]"></div>
      <div class="absolute bottom-[-10%] right-[-10%] w-[600px] h-[600px] rounded-full bg-indigo-400/15 blur-[120px]"></div>
      <div class="absolute top-[35%] right-[-15%] w-[500px] h-[500px] rounded-full bg-amber-300/15 blur-[100px]"></div>
    </div>

    <header class="h-16 bg-white/80 backdrop-blur-md border-b border-slate-200/60 flex items-center px-6 relative z-10 space-x-3">
      <img :src="isBazzar ? '/bz-icon.png' : '/laskar-corps.png'" class="h-8 w-auto object-contain" alt="HR Survey Logo" />
      <h1 class="text-lg font-bold text-slate-800">{{ surveyTitle }}</h1>
    </header>
    <main class="flex-1 max-w-3xl w-full mx-auto p-6 relative z-10">
      <router-view></router-view>
    </main>
  </div>
</template>
