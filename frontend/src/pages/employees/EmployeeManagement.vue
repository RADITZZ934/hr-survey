<template>
  <div class="space-y-6 animate-fade-in font-sans pb-6 select-none">
    
    <!-- Header Block -->
    <div class="flex flex-col sm:flex-row items-start sm:items-center justify-between gap-4">
      <div>
        <h1 class="text-2xl sm:text-3xl font-extrabold text-slate-800 tracking-tight">Manajemen Departemen & Karyawan</h1>
        <p class="text-xs sm:text-sm text-slate-400 mt-1">Daftar kepuasan per departemen beserta administrasi akun karyawan di dalamnya.</p>
      </div>
      <div class="flex items-center space-x-3 w-full sm:w-auto">
        <!-- Search bar -->
        <div class="relative w-full sm:w-64">
          <span class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-4.5 w-4.5 text-slate-400" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
          </span>
          <input 
            v-model="searchQuery"
            type="text" 
            placeholder="Cari karyawan..." 
            class="w-full pl-9 pr-4 py-2.5 bg-white border border-slate-200/80 rounded-xl text-xs font-semibold text-slate-700 placeholder-slate-400 focus:outline-none focus:ring-2 focus:ring-blue-500/20 focus:border-blue-500 transition-all shadow-xs"
          />
        </div>

        <button 
          @click="openAddModal('')"
          class="inline-flex items-center space-x-2 px-5 py-3 bg-[#4647AE] hover:bg-[#383994] active:bg-[#2e2e7a] text-white text-xs sm:text-sm font-bold rounded-xl shadow-md shadow-indigo-100 hover:-translate-y-0.5 active:scale-95 transition-all cursor-pointer whitespace-nowrap"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="h-4.5 w-4.5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2.5">
            <path stroke-linecap="round" stroke-linejoin="round" d="M12 4v16m8-8H4" />
          </svg>
          <span>Tambah Akun</span>
        </button>
      </div>
    </div>

    <!-- Grid Layout: 1 Department = 1 Grid Cell -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      
      <div 
        v-for="dept in groupedDepartments" 
        :key="dept.name"
        class="bg-white border border-slate-100/80 rounded-[2rem] shadow-md hover:shadow-lg transition-all duration-300 flex flex-col p-6 space-y-4 hover:-translate-y-1 relative"
      >
        <!-- Department Header Card -->
        <div class="flex items-start justify-between">
          <div class="flex items-center space-x-3">
            <div class="w-10 h-10 rounded-2xl bg-indigo-50 text-indigo-600 flex items-center justify-center font-bold text-sm">
              {{ (dept.displayName || dept.name).substring(0, 2).toUpperCase() }}
            </div>
            <div>
              <h3 class="font-extrabold text-slate-800 text-sm tracking-tight leading-none">{{ dept.displayName || dept.name }}</h3>
              <span class="text-[10px] font-bold text-slate-400 mt-1 block">
                <span v-if="dept.name === 'ANONYMOUS'">Khusus Respon Anonim</span>
                <span v-else>{{ dept.employees.length }} Karyawan</span>
              </span>
            </div>
          </div>

          <!-- Satisfaction Score Indicator Badge -->
          <div class="text-right">
            <div 
              v-if="dept.responseCount > 0" 
              class="px-2.5 py-1 text-[11px] font-extrabold rounded-lg inline-block"
              :class="getSatisfactionClass(dept.percentage)"
            >
              ★ {{ dept.avgScore.toFixed(1) }} / 5.0 ({{ Math.round(dept.percentage) }}%)
            </div>
            <div 
              v-else 
              class="px-2.5 py-1 text-[10px] font-bold rounded-lg inline-block bg-slate-50 text-slate-400 border border-slate-100"
            >
              Belum Ada Respon
            </div>
          </div>
        </div>

        <!-- Progress bar for satisfaction representation -->
        <div class="space-y-1">
          <div class="flex justify-between text-[9px] font-bold text-slate-400 uppercase tracking-wider">
            <span>Tingkat Kepuasan Kerja</span>
            <span v-if="dept.responseCount > 0">{{ Math.round(dept.percentage) }}%</span>
            <span v-else>-</span>
          </div>
          <div class="h-2 bg-slate-100 rounded-full overflow-hidden">
            <div 
              class="h-full rounded-full transition-all duration-500" 
              :class="getSatisfactionBarClass(dept.percentage)"
              :style="{ width: dept.responseCount > 0 ? dept.percentage + '%' : '0%' }"
            ></div>
          </div>
        </div>

        <!-- Employees List (Scrollable Container) -->
        <div class="flex-1 max-h-64 overflow-y-auto space-y-2 pr-1 scrollbar-thin">
          <h4 class="text-[10px] font-extrabold text-slate-400 uppercase tracking-wider pb-1 border-b border-slate-50">Daftar Anggota</h4>
          
          <div 
            v-for="emp in dept.employees" 
            :key="emp.id" 
            class="flex items-center justify-between p-2 hover:bg-slate-50 rounded-xl transition-all group/item"
          >
            <div class="flex items-center space-x-2.5 min-w-0">
               <div class="w-7 h-7 rounded-lg bg-slate-100 text-slate-600 flex items-center justify-center font-bold text-[10px] flex-shrink-0">
                {{ emp.username.substring(0, 2).toUpperCase() }}
              </div>
              <div class="min-w-0">
                <p class="font-bold text-xs text-slate-800 truncate">{{ emp.username }}</p>
                <p class="text-[10px] text-slate-400 truncate">{{ emp.email }}</p>
              </div>
            </div>

            <!-- Role pill + Action Buttons -->
            <div class="flex items-center space-x-1.5 flex-shrink-0">
              <!-- Satisfaction score of individual respondent -->
              <span 
                v-if="emp.percentage !== undefined && emp.percentage !== null"
                class="px-1.5 py-0.5 text-[9px] font-extrabold rounded-md bg-indigo-50 text-[#4647AE] border border-indigo-100"
                title="Skor Kepuasan Karyawan"
              >
                ★ {{ emp.avg_score.toFixed(1) }} ({{ emp.percentage }}%)
              </span>

              <!-- Role Badge -->
              <span 
                class="px-1.5 py-0.5 text-[9px] font-extrabold rounded-md group-hover/item:hidden"
                :class="getRoleClass(emp.role)"
              >
                {{ getRoleLabel(emp.role) }}
              </span>

              <!-- Hover Action Buttons -->
              <div class="hidden group-hover/item:flex items-center space-x-1">
                <button 
                  @click="openEditModal(emp)"
                  class="p-1 bg-blue-50 text-blue-600 rounded-md hover:bg-blue-100 transition-colors cursor-pointer"
                  title="Edit"
                >
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-3.5 w-3.5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" />
                  </svg>
                </button>
                <button 
                  @click="confirmDelete(emp)"
                  class="p-1 bg-rose-50 text-rose-600 rounded-md hover:bg-rose-100 transition-colors cursor-pointer"
                  title="Hapus"
                >
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-3.5 w-3.5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                  </svg>
                </button>
              </div>
            </div>

          </div>

          <div v-if="dept.employees.length === 0" class="py-4 text-center text-slate-400 font-medium text-[11px] italic">
            <span v-if="dept.name === 'ANONYMOUS'">Respon tidak dikaitkan dengan akun karyawan.</span>
            <span v-else>Belum ada anggota karyawan.</span>
          </div>
        </div>

        <!-- Add Employee inline to this department -->
        <button 
          v-if="dept.name !== 'ANONYMOUS'"
          @click="openAddModal(dept.name)"
          class="w-full py-2 bg-slate-50 hover:bg-slate-100 active:bg-slate-200 text-slate-600 text-xs font-bold rounded-xl transition-all flex items-center justify-center space-x-1 cursor-pointer border border-slate-100"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="h-3.5 w-3.5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2.5">
            <path stroke-linecap="round" stroke-linejoin="round" d="M12 4v16m8-8H4" />
          </svg>
          <span>Tambah Anggota</span>
        </button>

      </div>

    </div>

    <!-- Add/Edit Employee Dialog Modal -->
    <div 
      v-if="showModal" 
      class="fixed inset-0 bg-slate-900/40 backdrop-blur-xs flex items-center justify-center p-4 z-50 animate-fade-in"
      @click.self="closeModal"
    >
      <div class="bg-white rounded-[2rem] border border-slate-100 shadow-2xl max-w-md w-full overflow-hidden animate-scale-up">
        
        <!-- Header -->
        <div class="px-6 py-5 border-b border-slate-100 flex items-center justify-between bg-slate-50/50">
          <h3 class="font-bold text-slate-800 text-sm sm:text-base tracking-tight">
            {{ isEditMode ? 'Edit Karyawan' : 'Tambah Karyawan Baru' }}
          </h3>
          <button @click="closeModal" class="p-1 hover:bg-slate-100 rounded-lg text-slate-400 hover:text-slate-600 transition-colors">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>

        <!-- Form Body -->
        <form @submit.prevent="saveEmployee" class="p-6 space-y-4">
          <!-- Username input -->
          <div class="space-y-1">
            <label class="text-[11px] font-bold text-slate-500 uppercase tracking-wider block">Nama Pengguna (Username)</label>
            <input 
              v-model="form.username"
              type="text"
              required
              placeholder="Contoh: budi_siregar"
              class="w-full px-4 py-2.5 bg-slate-50 border border-slate-200/85 rounded-xl text-xs font-semibold text-slate-700 placeholder-slate-400 focus:outline-none focus:bg-white focus:ring-2 focus:ring-blue-500/20 focus:border-blue-500 transition-all"
            />
          </div>

          <!-- Email input -->
          <div class="space-y-1">
            <label class="text-[11px] font-bold text-slate-500 uppercase tracking-wider block">Email Resmi</label>
            <input 
              v-model="form.email"
              type="email"
              required
              placeholder="Contoh: budi@laskarbuah.com"
              class="w-full px-4 py-2.5 bg-slate-50 border border-slate-200/85 rounded-xl text-xs font-semibold text-slate-700 placeholder-slate-400 focus:outline-none focus:bg-white focus:ring-2 focus:ring-blue-500/20 focus:border-blue-500 transition-all"
            />
          </div>

          <!-- Password input -->
          <div class="space-y-1">
            <label class="text-[11px] font-bold text-slate-500 uppercase tracking-wider block">
              Kata Sandi {{ isEditMode ? '(Kosongkan jika tak diubah)' : '' }}
            </label>
            <input 
              v-model="form.password"
              type="password"
              :required="!isEditMode"
              placeholder="Masukkan sandi..."
              class="w-full px-4 py-2.5 bg-slate-50 border border-slate-200/85 rounded-xl text-xs font-semibold text-slate-700 placeholder-slate-400 focus:outline-none focus:bg-white focus:ring-2 focus:ring-blue-500/20 focus:border-blue-500 transition-all"
            />
          </div>

          <!-- Role input -->
          <div class="space-y-1">
            <label class="text-[11px] font-bold text-slate-500 uppercase tracking-wider block">Peran (Role)</label>
            <select 
              v-model="form.role"
              required
              class="w-full px-4 py-2.5 bg-slate-50 border border-slate-200/85 rounded-xl text-xs font-semibold text-slate-700 focus:outline-none focus:bg-white focus:ring-2 focus:ring-blue-500/20 focus:border-blue-500 transition-all cursor-pointer"
            >
              <option value="employee">Employee (Staf/Karyawan)</option>
              <option value="manager">Manager (Kepala Divisi)</option>
              <option value="hr">HR (Human Resources)</option>
              <option value="admin">Admin (Administrator)</option>
            </select>
          </div>

          <!-- Department input -->
          <div class="space-y-1">
            <label class="text-[11px] font-bold text-slate-500 uppercase tracking-wider block">Departemen</label>
            <select 
              v-model="form.department"
              required
              class="w-full px-4 py-2.5 bg-slate-50 border border-slate-200/85 rounded-xl text-xs font-semibold text-slate-700 focus:outline-none focus:bg-white focus:ring-2 focus:ring-blue-500/20 focus:border-blue-500 transition-all cursor-pointer"
            >
              <option value="" disabled>Pilih Departemen</option>
              <option v-for="dept in departments" :key="dept" :value="dept">{{ dept }}</option>
            </select>
          </div>

          <!-- Footer Actions -->
          <div class="pt-4 flex items-center justify-end space-x-3">
            <button 
              type="button"
              @click="closeModal"
              class="px-5 py-2.5 bg-slate-100 hover:bg-slate-200 text-slate-700 text-xs font-bold rounded-xl transition-colors cursor-pointer"
            >
              Batal
            </button>
            <button 
              type="submit"
              class="px-5 py-2.5 bg-[#4647AE] hover:bg-[#383994] active:bg-[#2e2e7a] text-white text-xs font-bold rounded-xl shadow-md shadow-indigo-100 hover:-translate-y-0.5 transition-all cursor-pointer"
            >
              {{ isEditMode ? 'Simpan Perubahan' : 'Daftarkan Karyawan' }}
            </button>
          </div>

        </form>

      </div>
    </div>

  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { getEmployees, createEmployee, updateEmployee, deleteEmployee, getDepartmentSatisfaction } from '../../services/employee.service';

const employees = ref([]);
const searchQuery = ref('');
const deptSatisfaction = ref([]);

const showModal = ref(false);
const isEditMode = ref(false);
const selectedEmployeeId = ref(null);

const form = ref({
  username: '',
  email: '',
  password: '',
  role: 'employee',
  department: ''
});

// Seed static departments list
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

const loadEmployees = async () => {
  try {
    const res = await getEmployees();
    employees.value = res.data || [];
  } catch (error) {
    console.error('Failed to load employees:', error);
  }
};

const loadSatisfactionData = async () => {
  try {
    const res = await getDepartmentSatisfaction();
    deptSatisfaction.value = res.data || [];
  } catch (error) {
    console.error('Failed to load department satisfaction:', error);
  }
};

// Group employees by department and map satisfaction values
const groupedDepartments = computed(() => {
  const groups = {};
  
  // Initialize standard departments
  departments.forEach(dept => {
    groups[dept] = [];
  });

  // Initialize ANONYMOUS group
  groups['ANONYMOUS'] = [];

  // Filter employees first by search query
  let filtered = employees.value;
  if (searchQuery.value.trim() !== '') {
    const q = searchQuery.value.toLowerCase().trim();
    filtered = filtered.filter(emp => 
      emp.username.toLowerCase().includes(q) || 
      emp.email.toLowerCase().includes(q)
    );
  }

  // Populate groups
  filtered.forEach(emp => {
    const dept = emp.department;
    if (groups[dept] !== undefined) {
      groups[dept].push(emp);
    }
  });

  // Map to list with scores
  return Object.keys(groups).map(deptName => {
    const scoreData = deptSatisfaction.value.find(s => s.department === deptName);
    const avgScore = scoreData ? scoreData.avg_score : 0;
    const percentage = avgScore ? (avgScore / 5.0) * 100 : 0;
    const responseCount = scoreData ? scoreData.count : 0;

    return {
      name: deptName,
      displayName: deptName === 'ANONYMOUS' ? 'Anonim (Anonymous)' : deptName,
      employees: groups[deptName],
      avgScore,
      percentage,
      responseCount
    };
  });
});

const openAddModal = (defaultDept) => {
  isEditMode.value = false;
  selectedEmployeeId.value = null;
  form.value = {
    username: '',
    email: '',
    password: '',
    role: 'employee',
    department: defaultDept === 'Umum' ? '' : defaultDept
  };
  showModal.value = true;
};

const openEditModal = (emp) => {
  isEditMode.value = true;
  selectedEmployeeId.value = emp.id;
  form.value = {
    username: emp.username,
    email: emp.email,
    password: '', 
    role: emp.role,
    department: emp.department || ''
  };
  showModal.value = true;
};

const closeModal = () => {
  showModal.value = false;
};

const saveEmployee = async () => {
  try {
    if (isEditMode.value) {
      await updateEmployee(selectedEmployeeId.value, form.value);
    } else {
      await createEmployee(form.value);
    }
    showModal.value = false;
    loadEmployees();
  } catch (error) {
    console.error('Failed to save employee:', error);
    alert(error.response?.data?.error || 'Gagal menyimpan data karyawan.');
  }
};

const confirmDelete = async (emp) => {
  if (confirm(`Apakah Anda yakin ingin menghapus akun karyawan "${emp.username}"?`)) {
    try {
      await deleteEmployee(emp.id);
      loadEmployees();
    } catch (error) {
      console.error('Failed to delete employee:', error);
      alert('Gagal menghapus karyawan.');
    }
  }
};

// Styling helper for Satisfaction badge
const getSatisfactionClass = (percentage) => {
  if (percentage >= 80) return 'bg-emerald-50 text-emerald-700 border border-emerald-100';
  if (percentage >= 60) return 'bg-amber-50 text-amber-700 border border-amber-100';
  return 'bg-rose-50 text-rose-700 border border-rose-100';
};

const getSatisfactionBarClass = (percentage) => {
  if (percentage >= 80) return 'bg-emerald-500';
  if (percentage >= 60) return 'bg-amber-500';
  return 'bg-rose-500';
};

// Role badge colors & text labels
const getRoleClass = (role) => {
  if (role === 'admin') return 'bg-rose-50 text-rose-700 border border-rose-100';
  if (role === 'hr') return 'bg-emerald-50 text-emerald-700 border border-emerald-100';
  if (role === 'manager') return 'bg-amber-50 text-amber-700 border border-amber-100';
  return 'bg-slate-50 text-slate-700 border border-slate-200/80';
};

const getRoleLabel = (role) => {
  if (role === 'admin') return 'Admin';
  if (role === 'hr') return 'HR';
  if (role === 'manager') return 'Manager';
  return 'Staf';
};

onMounted(() => {
  loadEmployees();
  loadSatisfactionData();
});
</script>

<style scoped>
.animate-fade-in {
  animation: fadeIn 0.4s ease-out forwards;
}
.animate-scale-up {
  animation: scaleUp 0.3s cubic-bezier(0.16, 1, 0.3, 1) forwards;
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(8px); }
  to { opacity: 1; transform: translateY(0); }
}

@keyframes scaleUp {
  from { opacity: 0; transform: scale(0.95); }
  to { opacity: 1; transform: scale(1); }
}

/* Custom scrollbar styling */
.scrollbar-thin::-webkit-scrollbar {
  width: 4px;
}
.scrollbar-thin::-webkit-scrollbar-track {
  background: transparent;
}
.scrollbar-thin::-webkit-scrollbar-thumb {
  background: #cbd5e1;
  border-radius: 9999px;
}
.scrollbar-thin::-webkit-scrollbar-thumb:hover {
  background: #94a3b8;
}
</style>
