<template>
  <div class="relative w-full inline-block font-sans">
    <!-- Trigger Input Field -->
    <div 
      @click="togglePicker"
      class="w-full bg-white border border-slate-200 hover:border-slate-300 rounded-xl px-4 py-2.5 text-sm flex items-center justify-between cursor-pointer transition-all shadow-xs"
      :class="isOpen ? 'ring-2 ring-[#4647AE]/20 border-[#4647AE]' : ''"
    >
      <span class="font-medium" :class="modelValue ? 'text-slate-800' : 'text-slate-400'">
        {{ formattedDisplayDate }}
      </span>
      <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 text-slate-400" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
        <path stroke-linecap="round" stroke-linejoin="round" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
      </svg>
    </div>

    <!-- Backdrop to close on outside click -->
    <div v-if="isOpen" @click="isOpen = false" class="fixed inset-0 z-40"></div>

    <!-- Dark Neumorphic DatePicker Card Popup (100% Matching Reference Design) -->
    <div 
      v-if="isOpen"
      class="absolute z-50 bg-[#2b2c30] text-white rounded-3xl p-5 shadow-2xl border border-slate-700/60 w-72 sm:w-80 animate-fade-in select-none"
      :class="[
        position === 'bottom' ? 'top-full mt-2' : 'bottom-full mb-2',
        align === 'right' ? 'right-0' : 'left-0'
      ]"
    >
      <!-- Top Header Navigation Bar -->
      <div class="flex items-center justify-between gap-3 mb-5">
        <!-- Month Selector Pill: ←  April •  → -->
        <div class="flex-1 bg-[#3a3b40] rounded-2xl px-3 py-2 flex items-center justify-between shadow-inner border border-white/5">
          <button 
            type="button" 
            @click.stop="prevMonth"
            class="text-slate-400 hover:text-white transition-colors text-sm px-1.5 py-0.5 rounded-lg hover:bg-white/10"
          >
            &larr;
          </button>

          <div class="flex items-center space-x-1.5 font-bold text-xs sm:text-sm tracking-wide">
            <span>{{ monthNames[currentMonth] }}</span>
            <span class="w-2 h-2 rounded-full bg-[#ff4d6d]"></span>
          </div>

          <button 
            type="button" 
            @click.stop="nextMonth"
            class="text-slate-400 hover:text-white transition-colors text-sm px-1.5 py-0.5 rounded-lg hover:bg-white/10"
          >
            &rarr;
          </button>
        </div>

        <!-- Year Selector Pill: 2026 -->
        <div class="bg-[#3a3b40] rounded-2xl px-4 py-2 font-bold text-xs sm:text-sm text-slate-200 shadow-inner border border-white/5 flex items-center justify-center">
          <span>{{ currentYear }}</span>
        </div>
      </div>

      <!-- Calendar Days Grid -->
      <div class="grid grid-cols-7 gap-1 text-center text-xs font-semibold">
        <div 
          v-for="cell in calendarCells" 
          :key="cell.key"
          @click.stop="selectCell(cell)"
          class="h-9 w-9 mx-auto rounded-full flex items-center justify-center transition-all cursor-pointer"
          :class="getCellClass(cell)"
        >
          {{ cell.day }}
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue';

const props = defineProps({
  modelValue: {
    type: String,
    default: ''
  },
  placeholder: {
    type: String,
    default: 'Select date'
  },
  position: {
    type: String,
    default: 'top' // 'top' or 'bottom'
  },
  align: {
    type: String,
    default: 'left' // 'left' or 'right'
  }
});

const emit = defineEmits(['update:modelValue']);

const isOpen = ref(false);
const monthNames = [
  'January', 'February', 'March', 'April', 'May', 'June',
  'July', 'August', 'September', 'October', 'November', 'December'
];

const now = new Date();
const currentMonth = ref(now.getMonth());
const currentYear = ref(now.getFullYear());

watch(() => props.modelValue, (newVal) => {
  if (newVal) {
    const d = new Date(newVal);
    if (!isNaN(d.getTime())) {
      currentMonth.value = d.getMonth();
      currentYear.value = d.getFullYear();
    }
  }
}, { immediate: true });

const togglePicker = () => {
  isOpen.value = !isOpen.value;
};

const prevMonth = () => {
  if (currentMonth.value === 0) {
    currentMonth.value = 11;
    currentYear.value--;
  } else {
    currentMonth.value--;
  }
};

const nextMonth = () => {
  if (currentMonth.value === 11) {
    currentMonth.value = 0;
    currentYear.value++;
  } else {
    currentMonth.value++;
  }
};

const formattedDisplayDate = computed(() => {
  if (!props.modelValue) return props.placeholder;
  const d = new Date(props.modelValue);
  if (isNaN(d.getTime())) return props.placeholder;
  const day = String(d.getDate()).padStart(2, '0');
  const month = monthNames[d.getMonth()].substring(0, 3);
  const year = d.getFullYear();
  return `${day} ${month} ${year}`;
});

// Generate calendar grid cells (42 cells max: 6 rows x 7 days)
const calendarCells = computed(() => {
  const year = currentYear.value;
  const month = currentMonth.value;

  const firstDayIndex = new Date(year, month, 1).getDay(); // 0 = Sun, 1 = Mon ...
  const daysInMonth = new Date(year, month + 1, 0).getDate();
  const daysInPrevMonth = new Date(year, month, 0).getDate();

  const cells = [];

  // Previous month padding days
  for (let i = firstDayIndex - 1; i >= 0; i--) {
    const day = daysInPrevMonth - i;
    const prevM = month === 0 ? 11 : month - 1;
    const prevY = month === 0 ? year - 1 : year;
    cells.push({
      key: `prev-${day}`,
      day,
      month: prevM,
      year: prevY,
      isCurrentMonth: false
    });
  }

  // Current month days
  for (let day = 1; day <= daysInMonth; day++) {
    cells.push({
      key: `curr-${day}`,
      day,
      month,
      year,
      isCurrentMonth: true
    });
  }

  // Next month padding days to complete grid
  const remaining = (7 - (cells.length % 7)) % 7;
  for (let day = 1; day <= remaining; day++) {
    const nextM = month === 11 ? 0 : month + 1;
    const nextY = month === 11 ? year + 1 : year;
    cells.push({
      key: `next-${day}`,
      day,
      month: nextM,
      year: nextY,
      isCurrentMonth: false
    });
  }

  return cells;
});

const selectCell = (cell) => {
  const m = String(cell.month + 1).padStart(2, '0');
  const d = String(cell.day).padStart(2, '0');
  const dateStr = `${cell.year}-${m}-${d}`;
  
  emit('update:modelValue', dateStr);
  isOpen.value = false;
};

const getCellClass = (cell) => {
  const m = String(cell.month + 1).padStart(2, '0');
  const d = String(cell.day).padStart(2, '0');
  const cellDateStr = `${cell.year}-${m}-${d}`;

  const isSelected = props.modelValue === cellDateStr;

  if (isSelected) {
    return 'bg-[#ff4d6d] text-white font-extrabold shadow-lg shadow-[#ff4d6d]/40 scale-110';
  }
  if (!cell.isCurrentMonth) {
    return 'text-slate-600 hover:text-slate-400 hover:bg-[#3a3b40]/50';
  }
  return 'text-slate-200 hover:bg-[#3a3b40] hover:text-white font-medium';
};
</script>

<style scoped>
.animate-fade-in {
  animation: fadeIn 0.25s cubic-bezier(0.16, 1, 0.3, 1) forwards;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(6px) scale(0.97);
  }
  to {
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}
</style>
