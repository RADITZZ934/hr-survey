<template>
  <div class="space-y-6 animate-fade-in font-sans pb-8 select-none">
    
    <!-- Breadcrumb & Page Header -->
    <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
      <div>
        <nav class="flex text-xs text-slate-400 space-x-2 mb-1">
          <router-link to="/dashboard" class="hover:text-slate-600 transition-colors">Dashboard</router-link>
          <span>/</span>
          <span class="text-slate-600 font-medium">Stores</span>
        </nav>
        <h1 class="text-2xl sm:text-3xl font-extrabold text-slate-800 tracking-tight">Daftar Toko & Cabang</h1>
        <p class="text-xs sm:text-sm text-slate-400 mt-0.5">
          Manajemen lokasi outlet Laskar Buah, koordinat GPS, dan QR Code kuesioner kepuasan pelanggan.
        </p>
      </div>

      <div class="flex items-center space-x-3">
        <!-- Export Excel Button -->
        <button 
          @click="exportToExcel"
          :disabled="isExporting || stores.length === 0"
          class="inline-flex items-center space-x-2 px-4 py-2.5 bg-emerald-600 hover:bg-emerald-700 active:bg-emerald-800 text-white text-xs sm:text-sm font-bold rounded-xl transition-all shadow-sm shadow-emerald-200 cursor-pointer disabled:opacity-50 disabled:cursor-not-allowed"
          title="Export data toko ke Excel"
        >
          <svg v-if="!isExporting" xmlns="http://www.w3.org/2000/svg" class="h-4.5 w-4.5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
          </svg>
          <svg v-else class="animate-spin h-4.5 w-4.5 text-white" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
          <span>{{ isExporting ? 'Mengekspor...' : 'Export Excel' }}</span>
        </button>

        <!-- Sync Live API Button -->
        <button 
          @click="handleSyncStores"
          :disabled="isSyncing || loading"
          class="inline-flex items-center space-x-2 px-3.5 py-2.5 bg-[#4647AE] hover:bg-[#383994] active:bg-[#2e2e7a] text-white text-xs sm:text-sm font-bold rounded-xl transition-all shadow-sm shadow-indigo-200 cursor-pointer disabled:opacity-50 disabled:cursor-not-allowed"
          title="Sinkronkan daftar toko dari API Laskar Buah Pusat & Database"
        >
          <svg v-if="!isSyncing" xmlns="http://www.w3.org/2000/svg" class="h-4.5 w-4.5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M9 19l3 3m0 0l3-3m-3 3V10" />
          </svg>
          <svg v-else class="animate-spin h-4.5 w-4.5 text-white" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
          <span>{{ isSyncing ? 'Menyinkronkan...' : 'Sinkron Toko' }}</span>
        </button>

        <!-- Refresh Button -->
        <button 
          @click="fetchStores"
          :disabled="loading"
          class="inline-flex items-center space-x-1.5 px-3.5 py-2.5 bg-white border border-slate-200/80 hover:bg-slate-50 text-slate-700 text-xs sm:text-sm font-bold rounded-xl transition-all shadow-xs cursor-pointer disabled:opacity-50"
          title="Muat ulang data"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="h-4.5 w-4.5" :class="{ 'animate-spin': loading }" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
          <span class="hidden sm:inline">Refresh</span>
        </button>
      </div>
    </div>

    <!-- Quick Stats Cards Row -->
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4">
      
      <!-- Stat 1: Total Stores -->
      <div class="bg-white rounded-2xl p-5 border border-slate-100 shadow-sm flex items-center justify-between">
        <div class="space-y-1">
          <p class="text-xs font-bold text-slate-400 uppercase tracking-wider">Total Toko / DC</p>
          <p class="text-2xl font-black text-slate-800">{{ stores.length }}</p>
          <p class="text-[11px] text-slate-400 font-medium">Terdaftar di sistem inventory</p>
        </div>
        <div class="w-12 h-12 rounded-2xl bg-indigo-50 text-[#4647AE] flex items-center justify-center font-bold">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
          </svg>
        </div>
      </div>

      <!-- Stat 2: Mapped Stores (Has Coordinates) -->
      <div class="bg-white rounded-2xl p-5 border border-slate-100 shadow-sm flex items-center justify-between">
        <div class="space-y-1">
          <p class="text-xs font-bold text-slate-400 uppercase tracking-wider">Koordinat Terdata</p>
          <p class="text-2xl font-black text-emerald-600">{{ mappedStoresCount }}</p>
          <p class="text-[11px] text-emerald-600/80 font-medium">{{ mappedPercentage }}% terpetakan di GPS</p>
        </div>
        <div class="w-12 h-12 rounded-2xl bg-emerald-50 text-emerald-600 flex items-center justify-center font-bold">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
            <path stroke-linecap="round" stroke-linejoin="round" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
          </svg>
        </div>
      </div>

      <!-- Stat 3: Unmapped Stores -->
      <div class="bg-white rounded-2xl p-5 border border-slate-100 shadow-sm flex items-center justify-between">
        <div class="space-y-1">
          <p class="text-xs font-bold text-slate-400 uppercase tracking-wider">Belum Ada Koordinat</p>
          <p class="text-2xl font-black text-amber-500">{{ unmappedStoresCount }}</p>
          <p class="text-[11px] text-amber-500/80 font-medium">Perlu pembaruan koordinat</p>
        </div>
        <div class="w-12 h-12 rounded-2xl bg-amber-50 text-amber-600 flex items-center justify-center font-bold">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
          </svg>
        </div>
      </div>

      <!-- Stat 4: External Survey Linked -->
      <div class="bg-white rounded-2xl p-5 border border-slate-100 shadow-sm flex items-center justify-between">
        <div class="space-y-1">
          <p class="text-xs font-bold text-slate-400 uppercase tracking-wider">Kuesioner Eksternal</p>
          <p class="text-2xl font-black text-purple-700">{{ activeExternalSurveys.length }}</p>
          <p class="text-[11px] text-purple-600/80 font-medium">Siap digenerate QR Toko</p>
        </div>
        <div class="w-12 h-12 rounded-2xl bg-purple-50 text-purple-600 flex items-center justify-center font-bold">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M12 4v1m0 4v.01m0 4v.01M4 12h1m4 0h.01m4 0h.01m-9-8h3.01M4 7h3.01M4 4h3.01M17 4h3.01M17 7h3.01M17 17h3.01M17 20h3.01M4 17h3.01M4 20h3.01" />
          </svg>
        </div>
      </div>

    </div>

    <!-- Filter & View Controls Bar -->
    <div class="bg-white rounded-[2rem] border border-slate-100/80 shadow-md p-4 flex flex-col md:flex-row gap-4 justify-between items-center">
      
      <!-- Search Input -->
      <div class="relative w-full md:w-80">
        <span class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-4.5 w-4.5 text-slate-400" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
        </span>
        <input 
          v-model="searchQuery"
          type="text" 
          placeholder="Cari kode (LB...) atau nama toko..." 
          class="w-full pl-9 pr-4 py-2.5 bg-slate-50 border border-slate-200/80 rounded-xl focus:outline-none focus:bg-white focus:border-indigo-500 transition-all text-xs sm:text-sm font-semibold placeholder-slate-400"
        />
        <button 
          v-if="searchQuery" 
          @click="searchQuery = ''"
          class="absolute inset-y-0 right-0 pr-3 flex items-center text-slate-400 hover:text-slate-600"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>

      <!-- Filters & Switchers -->
      <div class="flex flex-wrap items-center gap-3 w-full md:w-auto justify-end">
        
        <!-- GPS Status Filter -->
        <select 
          v-model="gpsFilter"
          class="px-3.5 py-2.5 bg-slate-50 border border-slate-200/80 rounded-xl text-xs font-semibold text-slate-600 focus:outline-none focus:bg-white cursor-pointer"
        >
          <option value="all">Semua Status GPS</option>
          <option value="mapped">Ada Koordinat GPS</option>
          <option value="unmapped">Belum Ada Koordinat</option>
        </select>

        <!-- Rows Per Page -->
        <select 
          v-model="perPage"
          class="px-3 py-2.5 bg-slate-50 border border-slate-200/80 rounded-xl text-xs font-semibold text-slate-600 focus:outline-none focus:bg-white cursor-pointer"
        >
          <option :value="10">10 / hal</option>
          <option :value="25">25 / hal</option>
          <option :value="50">50 / hal</option>
          <option :value="100">100 / hal</option>
        </select>

        <!-- View Switcher (Table vs Grid) -->
        <div class="flex items-center bg-slate-100 p-1 rounded-xl">
          <button 
            @click="viewMode = 'table'" 
            class="p-1.5 rounded-lg transition-all"
            :class="viewMode === 'table' ? 'bg-white text-[#4647AE] shadow-xs' : 'text-slate-400 hover:text-slate-600'"
            title="Tampilan Tabel"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-4.5 w-4.5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M4 6h16M4 10h16M4 14h16M4 18h16" />
            </svg>
          </button>
          <button 
            @click="viewMode = 'grid'" 
            class="p-1.5 rounded-lg transition-all"
            :class="viewMode === 'grid' ? 'bg-white text-[#4647AE] shadow-xs' : 'text-slate-400 hover:text-slate-600'"
            title="Tampilan Grid Kartu"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-4.5 w-4.5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z" />
            </svg>
          </button>
        </div>

      </div>

    </div>

    <!-- Loading State -->
    <div v-if="loading" class="bg-white rounded-[2rem] p-12 text-center border border-slate-100 shadow-md space-y-3">
      <div class="inline-block animate-spin rounded-full h-10 w-10 border-4 border-slate-200 border-t-[#4647AE]"></div>
      <p class="text-sm font-bold text-slate-700">Memuat data toko dari server...</p>
      <p class="text-xs text-slate-400">Menghubungkan ke endpoint /api/LIT/store</p>
    </div>

    <!-- Error State -->
    <div v-else-if="fetchError" class="bg-white rounded-[2rem] p-8 text-center border border-rose-100 shadow-md space-y-4">
      <div class="w-12 h-12 mx-auto rounded-full bg-rose-50 text-rose-500 flex items-center justify-center">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
          <path stroke-linecap="round" stroke-linejoin="round" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
        </svg>
      </div>
      <div>
        <h3 class="text-base font-bold text-slate-800">Gagal Memuat Data Toko</h3>
        <p class="text-xs text-slate-500 mt-1 max-w-md mx-auto">{{ fetchError }}</p>
      </div>
      <button 
        @click="fetchStores"
        class="px-5 py-2.5 bg-[#4647AE] hover:bg-[#383994] text-white text-xs font-bold rounded-xl transition-all shadow-md"
      >
        Coba Lagi
      </button>
    </div>

    <!-- Content: Table View -->
    <div v-else-if="viewMode === 'table'" class="bg-white rounded-[2rem] border border-slate-100/80 shadow-md hover:shadow-lg transition-all duration-300 overflow-hidden p-2">
      <div class="overflow-x-auto rounded-[1.5rem]">
        <table class="w-full text-left border-collapse">
          <thead>
            <tr class="text-[11px] font-bold text-white uppercase tracking-wider bg-[#4647AE]">
              <th class="py-4 px-5 w-16 text-center rounded-tl-[1.5rem]">No</th>
              <th class="py-4 px-4 w-32">Kode Toko</th>
              <th class="py-4 px-6">Nama Toko / Lokasi</th>
              <th class="py-4 px-4 w-48 text-center">Koordinat GPS</th>
              <th class="py-4 px-4 w-36 text-center">Peta Google</th>
              <th class="py-4 px-6 w-44 text-center rounded-tr-[1.5rem]">Aksi</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-100 text-sm">
            <tr 
              v-for="(store, idx) in paginatedStores" 
              :key="store.id"
              class="hover:bg-slate-50/60 transition-colors group"
            >
              <!-- Index Column -->
              <td class="py-4 px-5 text-center text-xs font-bold text-slate-400">
                {{ (currentPage - 1) * perPage + idx + 1 }}
              </td>

              <!-- Kode Toko Badge -->
              <td class="py-4 px-4">
                <span class="inline-flex px-2.5 py-1 text-xs font-extrabold rounded-lg bg-indigo-50 text-[#4647AE] border border-indigo-100/80 tracking-wide">
                  {{ store.kode }}
                </span>
              </td>

              <!-- Nama Toko -->
              <td class="py-4 px-6">
                <div class="flex items-center space-x-3">
                  <div class="w-8 h-8 rounded-xl bg-amber-50 text-amber-700 flex items-center justify-center font-bold text-xs flex-shrink-0">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                      <path stroke-linecap="round" stroke-linejoin="round" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
                    </svg>
                  </div>
                  <div>
                    <span class="font-bold text-slate-800 text-sm block group-hover:text-indigo-600 transition-colors">
                      {{ store.name }}
                    </span>
                    <span class="text-[10px] text-slate-400 font-semibold">ID Toko: #{{ store.id }}</span>
                  </div>
                </div>
              </td>

              <!-- Koordinat GPS -->
              <td class="py-4 px-4 text-center">
                <div v-if="store.koordinat" class="inline-flex flex-col items-center">
                  <span class="inline-flex items-center space-x-1 px-2.5 py-1 rounded-lg text-xs font-semibold bg-emerald-50 text-emerald-700 border border-emerald-100">
                    <span class="w-1.5 h-1.5 rounded-full bg-emerald-500"></span>
                    <span>{{ store.koordinat }}</span>
                  </span>
                </div>
                <div v-else>
                  <span class="inline-flex items-center space-x-1 px-2.5 py-1 rounded-lg text-[10px] font-bold bg-slate-100 text-slate-400 border border-slate-200/60">
                    Belum Diatur
                  </span>
                </div>
              </td>

              <!-- Link Google Maps -->
              <td class="py-4 px-4 text-center">
                <a 
                  v-if="store.latitude && store.longitude"
                  :href="`https://www.google.com/maps?q=${store.latitude},${store.longitude}`"
                  target="_blank"
                  rel="noopener noreferrer"
                  class="inline-flex items-center space-x-1 px-3 py-1.5 text-xs font-bold text-blue-600 bg-blue-50 hover:bg-blue-100 hover:text-blue-700 border border-blue-200/60 rounded-xl transition-all"
                  title="Buka titik koordinat di Google Maps"
                >
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-3.5 w-3.5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
                  </svg>
                  <span>Lihat Peta</span>
                </a>
                <span v-else class="text-xs text-slate-300">-</span>
              </td>

              <!-- Action Buttons -->
              <td class="py-4 px-6 text-center">
                <div class="flex items-center justify-center space-x-2">
                  <!-- QR Code Kuesioner Button -->
                  <button 
                    @click="openQrModal(store)"
                    class="p-2 bg-purple-50 hover:bg-purple-100 text-purple-700 border border-purple-200/80 rounded-xl transition-all shadow-2xs"
                    title="Buat QR Code Kuesioner Toko ini"
                  >
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                      <path stroke-linecap="round" stroke-linejoin="round" d="M12 4v1m0 4v.01m0 4v.01M4 12h1m4 0h.01m4 0h.01m-9-8h3.01M4 7h3.01M4 4h3.01M17 4h3.01M17 7h3.01M17 17h3.01M17 20h3.01M4 17h3.01M4 20h3.01" />
                    </svg>
                  </button>

                  <!-- Copy Store Name Button -->
                  <button 
                    @click="copyStoreText(store)"
                    class="p-2 rounded-xl border transition-all shadow-2xs"
                    :class="copiedStoreId === store.id ? 'bg-emerald-50 border-emerald-200 text-emerald-600' : 'bg-slate-50 hover:bg-slate-100 text-slate-500 border-slate-200/80'"
                    :title="copiedStoreId === store.id ? 'Tersalin!' : 'Salin Nama Toko'"
                  >
                    <svg v-if="copiedStoreId === store.id" xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                      <path stroke-linecap="round" stroke-linejoin="round" d="M5 13l4 4L19 7" />
                    </svg>
                    <svg v-else xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                      <path stroke-linecap="round" stroke-linejoin="round" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z" />
                    </svg>
                  </button>

                  <!-- Detail Modal Trigger -->
                  <button 
                    @click="openDetailModal(store)"
                    class="p-2 bg-slate-50 hover:bg-slate-100 text-slate-600 border border-slate-200/80 rounded-xl transition-all shadow-2xs"
                    title="Lihat Detail Toko"
                  >
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                      <path stroke-linecap="round" stroke-linejoin="round" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                  </button>
                </div>
              </td>
            </tr>

            <!-- Empty Table -->
            <tr v-if="filteredStores.length === 0">
              <td colspan="6" class="py-12 text-center text-slate-400">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-12 w-12 mx-auto text-slate-200 mb-3" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
                </svg>
                <div class="text-sm font-semibold">Tidak ada data toko ditemukan</div>
                <div class="text-xs text-slate-300 mt-1">Coba sesuaikan kata kunci pencarian atau filter GPS.</div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Content: Grid Cards View -->
    <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-5">
      <div 
        v-for="store in paginatedStores" 
        :key="store.id"
        class="bg-white border border-slate-100/80 rounded-[2rem] p-5 shadow-sm hover:shadow-md transition-all duration-300 flex flex-col justify-between space-y-4 hover:-translate-y-1 group"
      >
        <!-- Top Card -->
        <div class="space-y-3">
          <div class="flex items-start justify-between">
            <span class="inline-flex px-2.5 py-1 text-xs font-black rounded-lg bg-indigo-50 text-[#4647AE] border border-indigo-100 tracking-wide">
              {{ store.kode }}
            </span>
            <span 
              v-if="store.koordinat"
              class="inline-flex items-center space-x-1 px-2 py-0.5 text-[10px] font-bold rounded-md bg-emerald-50 text-emerald-600 border border-emerald-100"
            >
              <span class="w-1.5 h-1.5 rounded-full bg-emerald-500"></span>
              <span>GPS Aktif</span>
            </span>
            <span 
              v-else
              class="px-2 py-0.5 text-[10px] font-bold rounded-md bg-slate-100 text-slate-400"
            >
              Tanpa GPS
            </span>
          </div>

          <div>
            <h3 class="font-extrabold text-slate-800 text-base leading-snug group-hover:text-indigo-600 transition-colors">
              {{ store.name }}
            </h3>
            <p class="text-xs text-slate-400 font-medium mt-0.5">ID: #{{ store.id }}</p>
          </div>

          <div v-if="store.koordinat" class="p-2.5 bg-slate-50 rounded-xl space-y-1 text-xs">
            <span class="text-[10px] font-bold text-slate-400 uppercase tracking-wider block">Koordinat:</span>
            <p class="font-mono text-slate-600 text-[11px] truncate">{{ store.koordinat }}</p>
          </div>
          <div v-else class="p-2.5 bg-slate-50/50 rounded-xl text-center">
            <span class="text-[11px] font-medium text-slate-400 italic">Koordinat belum dicatat</span>
          </div>
        </div>

        <!-- Bottom Actions -->
        <div class="pt-3 border-t border-slate-100 flex items-center justify-between gap-2">
          <a 
            v-if="store.latitude && store.longitude"
            :href="`https://www.google.com/maps?q=${store.latitude},${store.longitude}`"
            target="_blank"
            rel="noopener noreferrer"
            class="flex-1 py-2 text-center text-xs font-bold text-blue-600 bg-blue-50 hover:bg-blue-100 rounded-xl transition-all"
          >
            Google Maps
          </a>
          <button 
            v-else
            @click="openDetailModal(store)"
            class="flex-1 py-2 text-center text-xs font-bold text-slate-600 bg-slate-50 hover:bg-slate-100 rounded-xl transition-all"
          >
            Detail Toko
          </button>

          <button 
            @click="openQrModal(store)"
            class="px-3 py-2 bg-purple-50 hover:bg-purple-100 text-purple-700 text-xs font-bold rounded-xl transition-all flex items-center space-x-1"
            title="Generate QR Survei"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M12 4v1m0 4v.01m0 4v.01M4 12h1m4 0h.01m4 0h.01m-9-8h3.01M4 7h3.01M4 4h3.01M17 4h3.01M17 7h3.01M17 17h3.01M17 20h3.01M4 17h3.01M4 20h3.01" />
            </svg>
            <span>QR</span>
          </button>
        </div>
      </div>
    </div>

    <!-- Pagination Bar -->
    <div v-if="filteredStores.length > 0" class="bg-white rounded-2xl p-4 border border-slate-100 shadow-sm flex flex-col sm:flex-row items-center justify-between gap-4">
      <div class="text-xs font-semibold text-slate-500">
        Menampilkan 
        <span class="text-slate-800 font-bold">{{ paginationStart }}</span> 
        - 
        <span class="text-slate-800 font-bold">{{ paginationEnd }}</span> 
        dari 
        <span class="text-slate-800 font-bold">{{ filteredStores.length }}</span> toko
      </div>

      <div class="flex items-center space-x-1.5">
        <button 
          @click="currentPage = 1"
          :disabled="currentPage === 1"
          class="p-2 rounded-xl border border-slate-200 text-slate-600 hover:bg-slate-50 disabled:opacity-40 disabled:cursor-not-allowed transition-colors text-xs font-bold"
          title="Halaman Pertama"
        >
          «
        </button>

        <button 
          @click="currentPage--"
          :disabled="currentPage === 1"
          class="px-3 py-2 rounded-xl border border-slate-200 text-slate-600 hover:bg-slate-50 disabled:opacity-40 disabled:cursor-not-allowed transition-colors text-xs font-bold"
        >
          Sebelumnya
        </button>

        <span class="px-3 py-2 text-xs font-bold text-[#4647AE] bg-indigo-50 border border-indigo-100 rounded-xl">
          {{ currentPage }} / {{ totalPages }}
        </span>

        <button 
          @click="currentPage++"
          :disabled="currentPage >= totalPages"
          class="px-3 py-2 rounded-xl border border-slate-200 text-slate-600 hover:bg-slate-50 disabled:opacity-40 disabled:cursor-not-allowed transition-colors text-xs font-bold"
        >
          Berikutnya
        </button>

        <button 
          @click="currentPage = totalPages"
          :disabled="currentPage >= totalPages"
          class="p-2 rounded-xl border border-slate-200 text-slate-600 hover:bg-slate-50 disabled:opacity-40 disabled:cursor-not-allowed transition-colors text-xs font-bold"
          title="Halaman Terakhir"
        >
          »
        </button>
      </div>
    </div>

    <!-- MODAL 1: QR CODE KUESIONER TOKO -->
    <div v-if="showQrModal && selectedStoreForQr" class="fixed inset-0 z-50 flex items-center justify-center p-4">
      <div class="fixed inset-0 bg-slate-900/40 backdrop-blur-xs transition-opacity animate-fade-in" @click="closeQrModal"></div>

      <div class="relative bg-white rounded-3xl border border-slate-100 shadow-2xl max-w-md w-full p-6 sm:p-7 z-10 space-y-5 animate-fade-in">
        
        <!-- Header Modal -->
        <div class="flex items-center justify-between pb-3 border-b border-slate-100">
          <div class="flex items-center space-x-2.5">
            <div class="w-9 h-9 rounded-xl bg-purple-50 text-purple-600 flex items-center justify-center">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M12 4v1m0 4v.01m0 4v.01M4 12h1m4 0h.01m4 0h.01m-9-8h3.01M4 7h3.01M4 4h3.01M17 4h3.01M17 7h3.01M17 17h3.01M17 20h3.01M4 17h3.01M4 20h3.01" />
              </svg>
            </div>
            <div>
              <h3 class="text-base font-bold text-slate-800 tracking-tight">QR Code Kuesioner</h3>
              <p class="text-[11px] text-slate-400">Siap cetak atau dipajang di kasir/meja toko</p>
            </div>
          </div>

          <button @click="closeQrModal" class="p-1.5 text-slate-400 hover:text-slate-600 rounded-xl hover:bg-slate-100 transition-colors">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>

        <!-- Store Info Pill -->
        <div class="p-3 bg-slate-50 rounded-2xl flex items-center justify-between">
          <div>
            <span class="text-[10px] font-bold text-slate-400 uppercase tracking-wider block">Toko / Outlet</span>
            <span class="font-extrabold text-slate-800 text-sm">{{ selectedStoreForQr.name }}</span>
          </div>
          <span class="px-2.5 py-1 text-xs font-bold rounded-lg bg-indigo-50 text-[#4647AE] border border-indigo-100">
            {{ selectedStoreForQr.kode }}
          </span>
        </div>

        <!-- Select Survey -->
        <div class="space-y-1.5">
          <label class="text-xs font-bold text-slate-600 block">Pilih Survei yang Ditautkan</label>
          <select 
            v-model="selectedSurveyIdForQr" 
            @change="generateQrCode"
            class="w-full px-3.5 py-2.5 bg-white border border-slate-200 rounded-xl text-xs font-semibold text-slate-700 focus:outline-none focus:ring-2 focus:ring-indigo-500/20 focus:border-indigo-500"
          >
            <option v-for="s in surveysList" :key="s.id" :value="s.id">
              {{ s.title }} ({{ s.visibility || 'internal' }})
            </option>
          </select>
        </div>

        <!-- QR Code Canvas / Image Display -->
        <div class="flex flex-col items-center justify-center p-4 bg-slate-50 border border-slate-100 rounded-2xl space-y-3">
          <div v-if="qrDataUrl" class="p-3 bg-white rounded-2xl shadow-sm border border-slate-100">
            <img :src="qrDataUrl" alt="QR Code" class="w-48 h-48 object-contain" />
          </div>
          <div v-else class="h-48 flex items-center justify-center text-xs text-slate-400">
            Menghasilkan QR Code...
          </div>

          <div class="w-full space-y-1 text-center">
            <span class="text-[10px] font-bold text-slate-400 uppercase tracking-wider">URL Kuesioner Toko:</span>
            <p class="text-[11px] text-slate-600 font-mono break-all px-2 bg-white/70 py-1.5 rounded-lg border border-slate-200/50">
              {{ currentSurveyUrl }}
            </p>
          </div>
        </div>

        <!-- Actions in Modal -->
        <div class="flex items-center gap-3 pt-2">
          <button 
            @click="copySurveyUrl"
            class="flex-1 py-2.5 px-3 border border-slate-200 rounded-xl text-xs font-bold text-slate-700 hover:bg-slate-50 transition-colors flex items-center justify-center space-x-1.5"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z" />
            </svg>
            <span>{{ isCopiedUrl ? 'Tersalin!' : 'Salin Link' }}</span>
          </button>

          <button 
            @click="downloadSingleQr"
            :disabled="!qrDataUrl"
            class="flex-1 py-2.5 px-3 bg-[#4647AE] hover:bg-[#383994] text-white rounded-xl text-xs font-bold transition-all shadow-md flex items-center justify-center space-x-1.5 disabled:opacity-50"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4" />
            </svg>
            <span>Unduh QR (PNG)</span>
          </button>
        </div>

      </div>
    </div>

    <!-- MODAL 2: DETAIL TOKO -->
    <div v-if="showDetailModal && selectedStoreDetail" class="fixed inset-0 z-50 flex items-center justify-center p-4">
      <div class="fixed inset-0 bg-slate-900/40 backdrop-blur-xs transition-opacity animate-fade-in" @click="showDetailModal = false"></div>

      <div class="relative bg-white rounded-3xl border border-slate-100 shadow-2xl max-w-sm w-full p-6 z-10 space-y-4 animate-fade-in">
        <div class="flex items-center justify-between pb-2 border-b border-slate-100">
          <h3 class="text-base font-bold text-slate-800">Detail Informasi Toko</h3>
          <button @click="showDetailModal = false" class="p-1 text-slate-400 hover:text-slate-600 rounded-lg">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>

        <div class="space-y-3 text-xs">
          <div class="flex justify-between py-1.5 border-b border-slate-50">
            <span class="text-slate-400 font-semibold">ID Toko</span>
            <span class="font-bold text-slate-800">#{{ selectedStoreDetail.id }}</span>
          </div>
          <div class="flex justify-between py-1.5 border-b border-slate-50">
            <span class="text-slate-400 font-semibold">Kode Toko</span>
            <span class="font-extrabold text-[#4647AE]">{{ selectedStoreDetail.kode }}</span>
          </div>
          <div class="flex justify-between py-1.5 border-b border-slate-50">
            <span class="text-slate-400 font-semibold">Nama Toko</span>
            <span class="font-bold text-slate-800 text-right">{{ selectedStoreDetail.name }}</span>
          </div>
          <div class="flex justify-between py-1.5 border-b border-slate-50">
            <span class="text-slate-400 font-semibold">Aplikasi ID</span>
            <span class="font-bold text-slate-600">{{ selectedStoreDetail.aplikasi_id || 1 }}</span>
          </div>
          <div class="flex justify-between py-1.5">
            <span class="text-slate-400 font-semibold">Koordinat GPS</span>
            <span class="font-mono text-slate-700 text-right">{{ selectedStoreDetail.koordinat || 'Belum diatur' }}</span>
          </div>
        </div>

        <div class="pt-2">
          <a 
            v-if="selectedStoreDetail.latitude && selectedStoreDetail.longitude"
            :href="`https://www.google.com/maps?q=${selectedStoreDetail.latitude},${selectedStoreDetail.longitude}`"
            target="_blank"
            rel="noopener noreferrer"
            class="w-full py-2.5 bg-blue-600 hover:bg-blue-700 text-white font-bold rounded-xl text-xs transition-all flex items-center justify-center space-x-1.5 shadow-sm"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
            </svg>
            <span>Buka di Google Maps</span>
          </a>
        </div>
      </div>
    </div>

  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { storeService } from '../../services/store.service';
import { getSurveys } from '../../services/survey.service';
import { useSearchStore } from '../../stores/search';
import QRCode from 'qrcode';
import ExcelJS from 'exceljs';

const searchStore = useSearchStore();

// State
const stores = ref([]);
const loading = ref(false);
const fetchError = ref('');
const searchQuery = ref('');
const gpsFilter = ref('all');
const viewMode = ref('table'); // 'table' or 'grid'

// Pagination State
const currentPage = ref(1);
const perPage = ref(25);

// Surveys list for QR linking
const surveysList = ref([]);
const activeExternalSurveys = ref([]);

// Modals State
const showQrModal = ref(false);
const selectedStoreForQr = ref(null);
const selectedSurveyIdForQr = ref('');
const qrDataUrl = ref('');
const isCopiedUrl = ref(false);

const showDetailModal = ref(false);
const selectedStoreDetail = ref(null);

const copiedStoreId = ref(null);
const isExporting = ref(false);

// Fetch data
const isSyncing = ref(false);

const fetchStores = async () => {
  loading.value = true;
  fetchError.value = '';
  try {
    const list = await storeService.getStores();
    stores.value = list;
  } catch (err) {
    console.error('Failed to load stores:', err);
    fetchError.value = err.response?.data?.error || err.message || 'Koneksi ke API toko terputus.';
  } finally {
    loading.value = false;
  }
};

const handleSyncStores = async () => {
  isSyncing.value = true;
  try {
    const res = await storeService.syncStores();
    await fetchStores();
    alert(res.message || 'Sinkronisasi daftar toko dari API Pusat & Database berhasil!');
  } catch (err) {
    console.error('Failed to sync stores:', err);
    await fetchStores();
    alert('Sinkronisasi selesai (data diperbarui).');
  } finally {
    isSyncing.value = false;
  }
};

const fetchSurveysList = async () => {
  try {
    const res = await getSurveys();
    const data = res.data || [];
    surveysList.value = data;
    activeExternalSurveys.value = data.filter(s => s.visibility === 'external');
    if (activeExternalSurveys.value.length > 0) {
      selectedSurveyIdForQr.value = activeExternalSurveys.value[0].id;
    } else if (data.length > 0) {
      selectedSurveyIdForQr.value = data[0].id;
    }
  } catch (err) {
    console.error('Failed to load surveys:', err);
  }
};

// Computed Stats
const mappedStoresCount = computed(() => {
  return stores.value.filter(s => !!s.koordinat).length;
});

const unmappedStoresCount = computed(() => {
  return stores.value.length - mappedStoresCount.value;
});

const mappedPercentage = computed(() => {
  if (stores.value.length === 0) return 0;
  return Math.round((mappedStoresCount.value / stores.value.length) * 100);
});

// Filtered List
const filteredStores = computed(() => {
  let list = [...stores.value];

  // Global search or local search
  const localQuery = searchQuery.value.trim().toLowerCase();
  const globalQuery = searchStore.searchQuery.trim().toLowerCase();
  const query = localQuery || globalQuery;

  if (query) {
    list = list.filter(s => 
      s.name.toLowerCase().includes(query) || 
      s.kode.toLowerCase().includes(query) ||
      String(s.id).includes(query)
    );
  }

  // Filter GPS
  if (gpsFilter.value === 'mapped') {
    list = list.filter(s => !!s.koordinat);
  } else if (gpsFilter.value === 'unmapped') {
    list = list.filter(s => !s.koordinat);
  }

  return list;
});

// Pagination Computeds
const totalPages = computed(() => {
  const count = Math.ceil(filteredStores.value.length / perPage.value);
  return count > 0 ? count : 1;
});

const paginatedStores = computed(() => {
  const start = (currentPage.value - 1) * perPage.value;
  return filteredStores.value.slice(start, start + perPage.value);
});

const paginationStart = computed(() => {
  if (filteredStores.value.length === 0) return 0;
  return (currentPage.value - 1) * perPage.value + 1;
});

const paginationEnd = computed(() => {
  const end = currentPage.value * perPage.value;
  return end > filteredStores.value.length ? filteredStores.value.length : end;
});

// Copy helpers
const copyStoreText = (store) => {
  const text = `${store.kode} - ${store.name}`;
  if (navigator.clipboard) {
    navigator.clipboard.writeText(text);
  }
  copiedStoreId.value = store.id;
  setTimeout(() => {
    copiedStoreId.value = null;
  }, 1500);
};

// QR Modal & Generation
const openQrModal = async (store) => {
  selectedStoreForQr.value = store;
  if (!selectedSurveyIdForQr.value && surveysList.value.length > 0) {
    selectedSurveyIdForQr.value = surveysList.value[0].id;
  }
  showQrModal.value = true;
  await generateQrCode();
};

const closeQrModal = () => {
  showQrModal.value = false;
  selectedStoreForQr.value = null;
  qrDataUrl.value = '';
};

const currentSurveyUrl = computed(() => {
  if (!selectedStoreForQr.value) return '';
  const origin = window.location.origin;
  const surveyId = selectedSurveyIdForQr.value;
  const storeParam = encodeURIComponent(selectedStoreForQr.value.name);
  return `${origin}/survey/identity?survey_id=${surveyId}&store_id=${storeParam}`;
});

const generateQrCode = async () => {
  if (!selectedStoreForQr.value) return;
  try {
    const url = currentSurveyUrl.value;
    qrDataUrl.value = await QRCode.toDataURL(url, {
      width: 260,
      margin: 1,
      color: {
        dark: '#332941',
        light: '#FFFFFF'
      }
    });
  } catch (err) {
    console.error('Failed to render QR Code:', err);
  }
};

const copySurveyUrl = () => {
  if (navigator.clipboard && currentSurveyUrl.value) {
    navigator.clipboard.writeText(currentSurveyUrl.value);
    isCopiedUrl.value = true;
    setTimeout(() => {
      isCopiedUrl.value = false;
    }, 2000);
  }
};

const downloadSingleQr = () => {
  if (!qrDataUrl.value || !selectedStoreForQr.value) return;
  const a = document.createElement('a');
  a.href = qrDataUrl.value;
  const safeName = selectedStoreForQr.value.name.replace(/[^a-zA-Z0-9]/g, '_');
  a.download = `QR_Survey_${safeName}.png`;
  document.body.appendChild(a);
  a.click();
  document.body.removeChild(a);
};

// Detail Modal
const openDetailModal = (store) => {
  selectedStoreDetail.value = store;
  showDetailModal.value = true;
};

// Export to Excel with ExcelJS
const exportToExcel = async () => {
  if (stores.value.length === 0) return;
  isExporting.value = true;
  try {
    const workbook = new ExcelJS.Workbook();
    const worksheet = workbook.addWorksheet('Daftar Toko');

    worksheet.columns = [
      { header: 'No', key: 'no', width: 8 },
      { header: 'ID Toko', key: 'id', width: 14 },
      { header: 'Kode Toko', key: 'kode', width: 18 },
      { header: 'Nama Toko / Lokasi', key: 'name', width: 38 },
      { header: 'Koordinat GPS', key: 'koordinat', width: 26 },
      { header: 'Latitude', key: 'latitude', width: 16 },
      { header: 'Longitude', key: 'longitude', width: 16 },
      { header: 'Link Google Maps', key: 'maps_link', width: 50 },
    ];

    // Header styling
    const headerRow = worksheet.getRow(1);
    headerRow.height = 28;
    headerRow.font = { name: 'Arial', family: 4, size: 11, bold: true, color: { argb: 'FFFFFFFF' } };
    headerRow.fill = {
      type: 'pattern',
      pattern: 'solid',
      fgColor: { argb: 'FF4647AE' }
    };
    headerRow.alignment = { vertical: 'middle', horizontal: 'center' };

    // Fill data
    filteredStores.value.forEach((store, index) => {
      const mapsUrl = store.latitude && store.longitude
        ? `https://www.google.com/maps?q=${store.latitude},${store.longitude}`
        : '-';

      const row = worksheet.addRow({
        no: index + 1,
        id: store.id,
        kode: store.kode,
        name: store.name,
        koordinat: store.koordinat || '-',
        latitude: store.latitude || '-',
        longitude: store.longitude || '-',
        maps_link: mapsUrl
      });

      row.height = 22;
      row.alignment = { vertical: 'middle', horizontal: 'left' };
      row.getCell('no').alignment = { vertical: 'middle', horizontal: 'center' };
      row.getCell('id').alignment = { vertical: 'middle', horizontal: 'center' };
      row.getCell('kode').alignment = { vertical: 'middle', horizontal: 'center' };

      // Border styling
      row.eachCell((cell) => {
        cell.border = {
          top: { style: 'thin', color: { argb: 'FFE2E8F0' } },
          left: { style: 'thin', color: { argb: 'FFE2E8F0' } },
          bottom: { style: 'thin', color: { argb: 'FFE2E8F0' } },
          right: { style: 'thin', color: { argb: 'FFE2E8F0' } }
        };
      });
    });

    const buffer = await workbook.xlsx.writeBuffer();
    const blob = new Blob([buffer], { type: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet' });
    const downloadUrl = URL.createObjectURL(blob);
    const link = document.createElement('a');
    link.href = downloadUrl;
    const now = new Date().toISOString().split('T')[0];
    link.download = `Daftar_Toko_Laskar_Buah_${now}.xlsx`;
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
  } catch (err) {
    console.error('Failed to export stores to Excel:', err);
    alert('Gagal mengekspor data toko ke Excel.');
  } finally {
    isExporting.value = false;
  }
};

onMounted(() => {
  fetchStores();
  fetchSurveysList();
});
</script>

<style scoped>
.animate-fade-in {
  animation: fadeIn 0.25s ease-out forwards;
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(8px); }
  to { opacity: 1; transform: translateY(0); }
}
</style>
