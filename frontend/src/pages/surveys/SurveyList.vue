<template>
  <div class="space-y-6 animate-fade-in">
    <!-- Breadcrumb & Header -->
    <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
      <div>
        <nav class="flex text-xs text-slate-400 space-x-2 mb-1">
          <router-link to="/dashboard" class="hover:text-slate-600 transition-colors">Dashboard</router-link>
          <span>/</span>
          <span class="text-slate-600 font-medium">Surveys</span>
        </nav>
        <h1 class="text-2xl font-bold text-slate-800 tracking-tight">Survey Management</h1>
      </div>
      <button 
        @click="toggleCreateWizard"
        class="inline-flex items-center space-x-2 px-5 py-2.5 bg-red-600 hover:bg-red-700 active:bg-red-800 text-white text-sm font-semibold rounded-xl transition-all shadow-sm shadow-red-200"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
          <path stroke-linecap="round" stroke-linejoin="round" d="M12 4v16m8-8H4" />
        </svg>
        <span>{{ isWizardOpen ? 'Close Panel' : 'Create Survey' }}</span>
      </button>
    </div>

    <!-- Inline Form Wizard Panel -->
    <div 
      class="transition-all duration-500 ease-in-out bg-white border border-slate-100 shadow-md hover:shadow-lg rounded-[2rem]"
      :class="isWizardOpen ? 'max-h-[1600px] opacity-100 p-6 sm:p-8 mb-6 overflow-visible relative z-30' : 'max-h-0 opacity-0 p-0 m-0 border-none overflow-hidden'"
    >
      <!-- Wizard Header -->
      <div class="pb-4 border-b border-amber-200/40 flex justify-between items-center mb-6">
        <div>
          <h2 class="text-lg font-bold text-slate-800">Create New Survey</h2>
          <p class="text-xs text-slate-400 mt-0.5">Step {{ currentStep }} of 3: {{ stepTitles[currentStep - 1] }}</p>
        </div>
        <div class="flex items-center space-x-2">
          <div class="flex space-x-1.5">
            <span class="w-6 h-1.5 rounded-full" :class="currentStep >= 1 ? 'bg-blue-600' : 'bg-slate-200'"></span>
            <span class="w-6 h-1.5 rounded-full" :class="currentStep >= 2 ? 'bg-blue-600' : 'bg-slate-200'"></span>
            <span class="w-6 h-1.5 rounded-full" :class="currentStep >= 3 ? 'bg-blue-600' : 'bg-slate-200'"></span>
          </div>
        </div>
      </div>

      <!-- Steps Content -->
      <div class="space-y-6">
        <!-- Step 1: General Info -->
        <div v-if="currentStep === 1" class="space-y-4">
          <div class="space-y-1.5">
            <label class="text-xs font-bold text-slate-500 uppercase tracking-wider">Survey Title</label>
            <input 
              v-model="wizardForm.title"
              type="text" 
              placeholder="e.g. Employee Engagement Survey 2026"
              class="w-full border border-slate-200 rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500/20 focus:border-blue-500 transition-all"
            />
          </div>
          <div class="space-y-1.5">
            <label class="text-xs font-bold text-slate-500 uppercase tracking-wider">Description</label>
            <textarea 
              v-model="wizardForm.description"
              rows="3"
              placeholder="Write a brief overview for respondents..."
              class="w-full border border-slate-200 rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500/20 focus:border-blue-500 transition-all resize-none"
            ></textarea>
          </div>
          <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <div class="space-y-1.5">
              <label class="text-xs font-bold text-slate-500 uppercase tracking-wider">Start Date</label>
              <CustomDatePicker 
                v-model="wizardForm.start_date"
                placeholder="Select start date"
                position="top"
              />
            </div>
            <div class="space-y-1.5">
              <label class="text-xs font-bold text-slate-500 uppercase tracking-wider">End Date</label>
              <CustomDatePicker 
                v-model="wizardForm.end_date"
                placeholder="Select end date"
                position="top"
              />
            </div>
          </div>
        </div>

        <!-- Step 2: Excel Upload & Dynamic Review List -->
        <div v-if="currentStep === 2" class="space-y-6">
          <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between p-4 bg-slate-50/50 border border-slate-100 rounded-2xl gap-4">
            <div>
              <h3 class="font-bold text-slate-800 text-sm">Upload Questions Template</h3>
              <p class="text-xs text-slate-400 mt-0.5">Prepare your survey questions offline using Excel, then import them here.</p>
            </div>
            <div class="flex items-center space-x-3">
              <a 
                href="#" 
                @click.prevent="downloadExcelTemplate" 
                class="px-4 py-2 bg-white border border-slate-200 text-xs font-semibold text-slate-600 rounded-xl hover:bg-slate-50 hover:text-blue-600 transition-all flex items-center space-x-1"
              >
                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4" />
                </svg>
                <span>Download Template</span>
              </a>
              <label class="px-4 py-2 bg-blue-600 text-white text-xs font-semibold rounded-xl hover:bg-blue-700 transition-all cursor-pointer flex items-center space-x-1">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-8l-4-4m0 0L8 8m4-4v12" />
                </svg>
                <span>Upload XLSX</span>
                <input type="file" accept=".xlsx, .xls" class="hidden" @change="handleExcelUpload" />
              </label>
            </div>
          </div>

          <!-- Questions Review Table List -->
          <div class="space-y-4">
            <div class="flex justify-between items-center">
              <h4 class="font-bold text-slate-800 text-sm">Questions Review List</h4>
              <button 
                @click="addQuestionRow" 
                class="px-3 py-1.5 border border-dashed border-blue-300 hover:border-blue-400 text-xs font-semibold text-blue-600 hover:bg-blue-50/50 rounded-xl transition-all"
              >
                + Add Question
              </button>
            </div>

            <div class="border border-slate-100 rounded-2xl overflow-y-auto max-h-[500px] shadow-sm">
              <table class="w-full text-left border-collapse">
                <thead class="sticky top-0 bg-slate-50 z-10 shadow-[0_1px_0_0_rgba(0,0,0,0.05)]">
                  <tr class="text-[10px] font-bold text-slate-400 uppercase tracking-wider">
                    <th class="py-3 px-4 w-12 text-center bg-slate-50">No</th>
                    <th class="py-3 px-4 bg-slate-50">Question Text</th>
                    <th class="py-3 px-4 w-36 bg-slate-50">Type</th>
                    <th class="py-3 px-4 w-44 bg-slate-50">Category</th>
                    <th class="py-3 px-4 w-28 text-center bg-slate-50">Required</th>
                    <th class="py-3 px-4 w-12 bg-slate-50"></th>
                  </tr>
                </thead>
                <tbody class="divide-y divide-slate-50 text-sm">
                  <tr v-for="(q, idx) in wizardForm.questions" :key="idx" class="hover:bg-slate-50/20 transition-colors">
                    <td class="py-3.5 px-4 text-center font-bold text-slate-400 text-xs">{{ idx + 1 }}</td>
                    <td class="py-3.5 px-4">
                      <input 
                        v-model="q.text"
                        type="text"
                        required
                        placeholder="Type question here..."
                        class="w-full bg-transparent border-b border-transparent focus:border-blue-500 focus:outline-none py-1 text-slate-700 text-sm font-semibold transition-all"
                      />
                    </td>
                    <td class="py-3.5 px-4">
                      <select 
                        v-model="q.type"
                        @change="handleTypeChange(q)"
                        class="w-full bg-slate-50 border border-slate-200/60 rounded-xl px-2 py-1.5 text-xs font-semibold text-slate-600 focus:outline-none focus:bg-white"
                      >
                        <option value="star">Rating Bintang</option>
                        <option value="essay">Essay</option>
                      </select>
                    </td>
                    <td class="py-3.5 px-4">
                      <input 
                        v-model="q.category"
                        type="text"
                        required
                        placeholder="Category (e.g. Work-Life Balance)"
                        class="w-full bg-transparent border-b border-transparent focus:border-blue-500 focus:outline-none py-1 text-slate-700 text-sm font-semibold transition-all"
                      />
                    </td>
                    <td class="py-3.5 px-4 text-center">
                      <button 
                        type="button"
                        @click="q.is_required = !q.is_required"
                        class="px-3 py-1 rounded-full text-[10px] font-bold tracking-wider uppercase transition-all"
                        :class="q.is_required ? 'bg-blue-50 text-blue-600 border border-blue-200' : 'bg-slate-100 text-slate-400 border border-slate-200/60'"
                      >
                        {{ q.is_required ? 'Yes' : 'No' }}
                      </button>
                    </td>
                    <td class="py-3.5 px-4 text-center">
                      <button 
                        @click="removeQuestionRow(idx)"
                        class="p-1 text-slate-300 hover:text-rose-600 rounded-lg hover:bg-slate-100 transition-colors"
                      >
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                          <path stroke-linecap="round" stroke-linejoin="round" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                        </svg>
                      </button>
                    </td>
                  </tr>
                  <tr v-if="wizardForm.questions.length === 0">
                    <td colspan="7" class="py-8 text-center text-slate-400 text-xs italic bg-slate-50/20">
                      No questions added yet. Download the template or click "+ Add Question" to begin.
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>

        <!-- Step 3: Review Details -->
        <div v-if="currentStep === 3" class="space-y-4 bg-slate-50 rounded-2xl p-6 border border-slate-100 text-slate-700 text-sm">
          <div class="space-y-1">
            <span class="text-xs font-bold text-slate-400 uppercase tracking-wider block">Survey Title</span>
            <p class="font-bold text-slate-800 text-base">{{ wizardForm.title || 'Untitled Survey' }}</p>
          </div>
          <div class="space-y-1">
            <span class="text-xs font-bold text-slate-400 uppercase tracking-wider block">Description</span>
            <p class="text-slate-500 leading-relaxed">{{ wizardForm.description || 'No description provided.' }}</p>
          </div>
          <div class="grid grid-cols-2 gap-4 pt-2 border-t border-slate-200/50">
            <div class="space-y-1">
              <span class="text-xs font-bold text-slate-400 uppercase tracking-wider block">Start Date</span>
              <p class="font-semibold text-slate-700">{{ wizardForm.start_date || '-' }}</p>
            </div>
            <div class="space-y-1">
              <span class="text-xs font-bold text-slate-400 uppercase tracking-wider block">End Date</span>
              <p class="font-semibold text-slate-700">{{ wizardForm.end_date || '-' }}</p>
            </div>
          </div>
          <div class="space-y-1 pt-2 border-t border-slate-200/50">
            <span class="text-xs font-bold text-slate-400 uppercase tracking-wider block">Total Questions</span>
            <p class="font-semibold text-slate-700">{{ wizardForm.questions.length }} custom questions configured</p>
          </div>
        </div>
      </div>

      <!-- Wizard Footer Actions -->
      <div class="px-6 py-4 border-t border-slate-100 bg-slate-50/50 flex justify-between mt-6 -mx-6 -mb-6 rounded-b-2xl">
        <button 
          @click="prevStep" 
          :disabled="currentStep === 1"
          class="px-4 py-2 border border-slate-200 text-slate-600 rounded-xl hover:bg-white text-xs font-semibold transition-colors disabled:opacity-40 disabled:cursor-not-allowed bg-white"
        >
          Back
        </button>
        
        <div class="flex space-x-2">
          <button 
            @click="closeWizard" 
            class="px-4 py-2 border border-transparent text-slate-500 hover:text-slate-700 text-xs font-semibold rounded-xl transition-all"
          >
            Cancel
          </button>
          
          <button 
            v-if="currentStep < 3"
            @click="nextStep"
            :disabled="!isStepValid"
            class="px-5 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-xl text-xs font-semibold transition-colors disabled:opacity-40 disabled:cursor-not-allowed"
          >
            Next
          </button>
          
          <button 
            v-else
            @click="submitSurvey"
            class="px-5 py-2 bg-emerald-600 hover:bg-emerald-700 text-white rounded-xl text-xs font-semibold transition-colors"
          >
            Publish Survey
          </button>
        </div>
      </div>
    </div>

    <!-- Filters & Search -->
    <div class="bg-white rounded-[2rem] border border-slate-100/80 shadow-md hover:shadow-lg transition-all duration-300 p-4 flex flex-col md:flex-row gap-4 justify-between items-center">
      <!-- Search Input -->
      <div class="relative w-full md:w-80">
        <span class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-slate-400" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
        </span>
        <input 
          v-model="searchQuery"
          type="text" 
          placeholder="Search survey title..." 
          class="w-full pl-10 pr-4 py-2 bg-slate-50 border border-transparent rounded-xl focus:outline-none focus:bg-white focus:border-slate-200 transition-colors placeholder-slate-400 text-sm"
        />
      </div>

      <!-- Dropdown Filters -->
      <div class="flex w-full md:w-auto gap-4">
        <!-- Status Filter -->
        <select 
          v-model="statusFilter"
          class="w-full md:w-44 px-3 py-2 bg-slate-50 border border-transparent rounded-xl focus:outline-none focus:bg-white focus:border-slate-200 text-sm text-slate-600 transition-colors"
        >
          <option value="all">All Statuses</option>
          <option value="active">Active</option>
          <option value="draft">Draft</option>
          <option value="closed">Closed</option>
        </select>

        <!-- Sort Filter -->
        <select 
          v-model="sortBy"
          class="w-full md:w-44 px-3 py-2 bg-slate-50 border border-transparent rounded-xl focus:outline-none focus:bg-white focus:border-slate-200 text-sm text-slate-600 transition-colors"
        >
          <option value="newest">Newest First</option>
          <option value="oldest">Oldest First</option>
        </select>
      </div>
    </div>

    <!-- Survey Table -->
    <div class="bg-white rounded-[2rem] border border-slate-100/80 shadow-md hover:shadow-lg transition-all duration-300 overflow-hidden p-2">
      <div class="overflow-x-auto rounded-[1.5rem]">
        <table class="w-full text-left border-collapse">
          <thead>
            <tr class="text-[11px] font-bold text-white uppercase tracking-wider bg-[#4647AE]">
              <th class="py-4 px-6 rounded-tl-[1.5rem]">Survey Details</th>
              <th class="py-4 px-4 text-center">Status</th>
              <th class="py-4 px-4 text-center">Timeline</th>
              <th class="py-4 px-4 text-right">Responses</th>
              <th class="py-4 px-6 text-center rounded-tr-[1.5rem]">Actions</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-50 text-sm">
            <tr 
              v-for="survey in filteredSurveys" 
              :key="survey.id" 
              class="hover:bg-slate-50/40 transition-colors group"
            >
              <!-- Info Column -->
              <td class="py-5 px-6">
                <div class="space-y-1">
                  <span class="font-semibold text-slate-800 text-base block">{{ survey.title }}</span>
                  <span class="text-xs text-slate-400 block line-clamp-1 max-w-md">{{ survey.description || 'No description provided.' }}</span>
                </div>
              </td>
              
              <!-- Status Badge -->
              <td class="py-5 px-4 text-center">
                <span 
                  class="inline-flex px-2.5 py-1 text-xs font-semibold rounded-lg capitalize"
                  :class="statusBadgeClass(survey.status)"
                >
                  {{ survey.status }}
                </span>
              </td>

              <!-- Dates Timeline -->
              <td class="py-5 px-4 text-center text-slate-500 whitespace-nowrap text-xs">
                <div class="font-medium">{{ formatDate(survey.start_date || survey.created_at) }}</div>
                <div class="text-[10px] text-slate-400 mt-0.5">to {{ formatDate(survey.end_date) }}</div>
              </td>

              <!-- Responses Progress -->
              <td class="py-5 px-4 text-right">
                <div class="font-bold text-slate-800">
                  {{ survey.responses_count || 0 }}
                </div>
                <div class="text-[10px] text-slate-400">filled responses</div>
              </td>

              <!-- Actions -->
              <td class="py-5 px-6 text-center">
                <div class="flex items-center justify-center space-x-2.5">
                  <!-- Copy Link button -->
                  <button 
                    @click="copySurveyLink(survey.id)"
                    class="p-2.5 rounded-xl border transition-all shadow-xs flex items-center justify-center"
                    :class="copiedSurveyId === survey.id ? 'bg-emerald-50 border-emerald-200 text-emerald-600' : 'bg-slate-50/80 border-slate-200/80 text-slate-500 hover:text-blue-600 hover:bg-blue-50 hover:border-blue-200'"
                    :title="copiedSurveyId === survey.id ? 'Link Copied!' : 'Copy Survey Link'"
                  >
                    <svg v-if="copiedSurveyId === survey.id" xmlns="http://www.w3.org/2000/svg" class="h-4.5 w-4.5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                      <path stroke-linecap="round" stroke-linejoin="round" d="M5 13l4 4L19 7" />
                    </svg>
                    <svg v-else xmlns="http://www.w3.org/2000/svg" class="h-4.5 w-4.5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                      <path stroke-linecap="round" stroke-linejoin="round" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z" />
                    </svg>
                  </button>

                  <!-- Preview button -->
                  <router-link 
                    :to="'/survey/identity?survey_id=' + survey.id"
                    class="p-2.5 rounded-xl border border-slate-200/80 bg-slate-50/80 text-slate-500 hover:text-blue-600 hover:bg-blue-50 hover:border-blue-200 transition-all shadow-xs flex items-center justify-center"
                    title="Preview Survey"
                  >
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-4.5 w-4.5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                      <path stroke-linecap="round" stroke-linejoin="round" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                      <path stroke-linecap="round" stroke-linejoin="round" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                    </svg>
                  </router-link>

                  <!-- Toggle activation -->
                  <button 
                    @click="toggleStatus(survey)"
                    class="p-2.5 rounded-xl border border-slate-200/80 bg-slate-50/80 text-slate-500 hover:text-amber-600 hover:bg-amber-50 hover:border-amber-200 transition-all shadow-xs flex items-center justify-center"
                    :title="survey.status === 'active' ? 'Pause Survey' : 'Activate Survey'"
                  >
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-4.5 w-4.5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                      <path stroke-linecap="round" stroke-linejoin="round" d="M14.752 11.168l-3.197-2.132A1 1 0 0010 9.87v4.263a1 1 0 001.555.832l3.197-2.132a1 1 0 000-1.664z" />
                      <path stroke-linecap="round" stroke-linejoin="round" d="M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                  </button>

                  <!-- Delete -->
                  <button 
                    @click="deleteSurvey(survey.id)"
                    class="p-2.5 rounded-xl border border-slate-200/80 bg-slate-50/80 text-slate-500 hover:text-rose-600 hover:bg-rose-50 hover:border-rose-200 transition-all shadow-xs flex items-center justify-center"
                    title="Delete Survey"
                  >
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-4.5 w-4.5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                      <path stroke-linecap="round" stroke-linejoin="round" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                    </svg>
                  </button>
                </div>
              </td>
            </tr>

            <!-- Empty state -->
            <tr v-if="filteredSurveys.length === 0">
              <td colspan="5" class="py-12 text-center text-slate-400">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-12 w-12 mx-auto text-slate-200 mb-3" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
                </svg>
                <div class="text-sm font-semibold">No surveys found</div>
                <div class="text-xs text-slate-300 mt-1">Try resetting filters or click "Create Survey" to add one.</div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { getSurveys, createSurvey } from '../../services/survey.service';
import { useSearchStore } from '../../stores/search';
import CustomDatePicker from '../../components/CustomDatePicker.vue';
import * as XLSX from 'xlsx';

const surveys = ref([]);
const searchQuery = ref('');
const statusFilter = ref('all');
const sortBy = ref('newest');
const searchStore = useSearchStore();
const copiedSurveyId = ref(null);

const copySurveyLink = (surveyId) => {
  const link = `${window.location.origin}/survey/identity?survey_id=${surveyId}`;
  if (navigator.clipboard && navigator.clipboard.writeText) {
    navigator.clipboard.writeText(link).then(() => {
      copiedSurveyId.value = surveyId;
      setTimeout(() => {
        copiedSurveyId.value = null;
      }, 2000);
    }).catch(err => {
      console.error('Failed to copy survey link via API:', err);
      fallbackCopyText(link, surveyId);
    });
  } else {
    fallbackCopyText(link, surveyId);
  }
};

const fallbackCopyText = (text, surveyId) => {
  const textArea = document.createElement("textarea");
  textArea.value = text;
  textArea.style.top = "0";
  textArea.style.left = "0";
  textArea.style.position = "fixed";
  document.body.appendChild(textArea);
  textArea.focus();
  textArea.select();
  try {
    const successful = document.execCommand('copy');
    if (successful) {
      copiedSurveyId.value = surveyId;
      setTimeout(() => {
        copiedSurveyId.value = null;
      }, 2000);
    }
  } catch (err) {
    console.error('Fallback copy failed:', err);
  }
  document.body.removeChild(textArea);
};

// Wizard Form State
const isWizardOpen = ref(false);
const currentStep = ref(1);
const stepTitles = ['General Information', 'Excel Import & Review Questions', 'Review & Publish'];

const wizardForm = ref({
  title: '',
  description: '',
  start_date: '',
  end_date: '',
  questions: []
});

const fetchInitialData = async () => {
  try {
    const res = await getSurveys();
    surveys.value = res.data;
  } catch (error) {
    console.error('Failed to fetch initial surveys:', error);
  }
};

const filteredSurveys = computed(() => {
  let list = [...surveys.value];

  const localQuery = searchQuery.value.trim();
  const globalQuery = searchStore.searchQuery.trim();
  const query = (localQuery || globalQuery).toLowerCase();

  if (query) {
    list = list.filter(s => s.title.toLowerCase().includes(query));
  }

  if (statusFilter.value !== 'all') {
    list = list.filter(s => s.status === statusFilter.value);
  }

  if (sortBy.value === 'newest') {
    list.sort((a, b) => new Date(b.start_date || b.created_at) - new Date(a.start_date || a.created_at));
  } else if (sortBy.value === 'oldest') {
    list.sort((a, b) => new Date(a.start_date || a.created_at) - new Date(b.start_date || b.created_at));
  }

  return list;
});

// Step validation
const isStepValid = computed(() => {
  if (currentStep.value === 1) {
    return wizardForm.value.title.trim() !== '' && wizardForm.value.start_date !== '' && wizardForm.value.end_date !== '';
  }
  if (currentStep.value === 2) {
    return wizardForm.value.questions.length > 0 && wizardForm.value.questions.every(q => q.text.trim() !== '');
  }
  return true;
});

const toggleCreateWizard = () => {
  if (isWizardOpen.value) {
    closeWizard();
  } else {
    openCreateWizard();
  }
};

const openCreateWizard = () => {
  wizardForm.value = {
    title: '',
    description: '',
    start_date: '',
    end_date: '',
    questions: []
  };
  currentStep.value = 1;
  isWizardOpen.value = true;
};

const closeWizard = () => {
  isWizardOpen.value = false;
};

const nextStep = () => {
  if (isStepValid.value && currentStep.value < 3) {
    currentStep.value++;
  }
};

const prevStep = () => {
  if (currentStep.value > 1) {
    currentStep.value--;
  }
};

// Dynamic Template Download
const downloadExcelTemplate = () => {
  const wsData = [
    ["Work-Life Balance"], // Category Header Row
    ["Pertanyaan", "Tipe", "", "ATURAN PENGISIAN TEMPLATE:"],
    ["Seberapa puas Anda dengan keseimbangan waktu kerja (Work-Life Balance) di perusahaan?", 1, "", "1. Tulis nama kategori pada baris tersendiri (tanpa mengisi kolom Tipe) sebagai pembatas kategori."],
    ["Berikan saran atau kritik Anda untuk meningkatkan kenyamanan bekerja di kantor.", 2, "", "2. Di bawah baris kategori, buat header tabel dengan kolom 'Pertanyaan' dan 'Tipe'."],
    ["", "", "", "3. Kolom 'Pertanyaan' diisi dengan teks pertanyaan kuesioner."],
    ["Manager Support"], // Category Header Row
    ["Pertanyaan", "Tipe", "", "4. Kolom 'Tipe' diisi angka: '1' (Rating Bintang) atau '2' (Essay)."],
    ["Bagaimana Anda menilai dukungan moral dan feedback dari atasan langsung Anda?", 1, "", "5. Pertanyaan di bawah judul kategori otomatis dikelompokkan ke dalam kategori tersebut."],
    ["", "", "", "6. Harap tidak mengubah nama kolom header 'Pertanyaan' dan 'Tipe'."]
  ];
  
  const wb = XLSX.utils.book_new();
  const ws = XLSX.utils.aoa_to_sheet(wsData);

  // Set column widths
  ws['!cols'] = [
    { wch: 65 }, // Pertanyaan
    { wch: 10 }, // Tipe
    { wch: 5 },  // Spasi kosong
    { wch: 55 }  // Rules/Petunjuk
  ];

  XLSX.utils.book_append_sheet(wb, ws, "Template Pertanyaan");
  XLSX.writeFile(wb, "survey_template.xlsx");
};

// Excel Upload Parser
const handleExcelUpload = (event) => {
  const file = event.target.files[0];
  if (!file) return;

  const reader = new FileReader();
  reader.onload = (e) => {
    const data = e.target.result;
    const workbook = XLSX.read(data, { type: 'binary' });
    const firstSheetName = workbook.SheetNames[0];
    const worksheet = workbook.Sheets[firstSheetName];
    
    // Parse to 2D array of rows
    const rawRows = XLSX.utils.sheet_to_json(worksheet, { header: 1 });
    
    let currentCategory = "";
    const parsedQuestions = [];

    for (let i = 0; i < rawRows.length; i++) {
      const row = rawRows[i];
      if (!row || row.length === 0) continue;

      const colA = String(row[0] || "").trim();
      const colB = String(row[1] || "").trim();

      // Check if it's empty row
      if (!colA && !colB) continue;

      // Check if it's a sub-header row like ["Pertanyaan", "Tipe"]
      const isHeaderRow = colA.toLowerCase().startsWith("pertanyaan") || colA.toLowerCase().startsWith("question") || colB.toLowerCase().startsWith("tipe") || colB.toLowerCase().startsWith("type");
      if (isHeaderRow) {
        continue;
      }

      // Check if this row is a Category Header
      if (colA && !colB) {
        // If it's a line from instruction block (usually rules are placed in column D/index 3)
        const colD = String(row[3] || "").trim();
        if (colA.includes("ATURAN") || colA.includes("RULES") || colD) {
          continue;
        }
        
        currentCategory = colA;
        continue;
      }

      // If it has both question text and type, it's a question row
      if (colA && colB) {
        let type = "star";
        if (colB === "2" || colB.toLowerCase().includes("essay") || colB.toLowerCase().includes("text") || colB.toLowerCase().includes("tulis")) {
          type = "essay";
        }

        const is_required = (type === "star");

        parsedQuestions.push({
          text: colA,
          type,
          category: currentCategory || (type === "star" ? "Rating Bintang" : "Essay"),
          is_required
        });
      }
    }

    wizardForm.value.questions = [...wizardForm.value.questions, ...parsedQuestions];
  };

  reader.readAsBinaryString(file);
  event.target.value = ""; // Reset input
};

// Interactive Question Grid actions
const addQuestionRow = () => {
  wizardForm.value.questions.push({
    text: '',
    type: 'star',
    category: 'Rating Bintang',
    is_required: true
  });
};

const removeQuestionRow = (idx) => {
  wizardForm.value.questions.splice(idx, 1);
};

const handleTypeChange = (q) => {
  // Enforce defaults on type change: Rating Bintang -> Required YES, Essay -> Required NO
  q.is_required = (q.type === 'star');
};

const submitSurvey = async () => {
  try {
    await createSurvey({
      title: wizardForm.value.title,
      description: wizardForm.value.description,
      start_date: wizardForm.value.start_date,
      end_date: wizardForm.value.end_date,
      questions: wizardForm.value.questions
    });
    isWizardOpen.value = false;
    await fetchInitialData();
  } catch (error) {
    console.error('Failed to publish survey:', error);
  }
};

const toggleStatus = (survey) => {
  survey.status = survey.status === 'active' ? 'closed' : 'active';
};

const deleteSurvey = (id) => {
  if (confirm('Are you sure you want to delete this survey?')) {
    surveys.value = surveys.value.filter(s => s.id !== id);
  }
};

const formatDate = (dateStr) => {
  if (!dateStr) return '-';
  const options = { year: 'numeric', month: 'short', day: 'numeric' };
  return new Date(dateStr).toLocaleDateString('en-US', options);
};

const statusBadgeClass = (status) => {
  if (status === 'active') return 'text-emerald-700 bg-emerald-50';
  if (status === 'draft') return 'text-slate-600 bg-slate-100';
  if (status === 'closed') return 'text-rose-700 bg-rose-50';
  return 'text-slate-400 bg-slate-50';
};

onMounted(() => {
  fetchInitialData();
});
</script>

<style scoped>
.animate-fade-in {
  animation: fadeIn 0.3s ease-out forwards;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(6px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
