<template>
  <div class="min-h-screen bg-gradient-to-tr from-slate-100 via-slate-50 to-blue-50/50 flex flex-col relative">
    <!-- Ambient Background Blur Orbs Container -->
    <div class="absolute inset-0 overflow-hidden pointer-events-none z-0">
      <div class="absolute top-[-10%] left-[-10%] w-[600px] h-[600px] rounded-full bg-blue-400/15 blur-[120px]"></div>
      <div class="absolute bottom-[-10%] right-[-10%] w-[600px] h-[600px] rounded-full bg-indigo-400/15 blur-[120px]"></div>
      <div class="absolute top-[35%] right-[-15%] w-[500px] h-[500px] rounded-full bg-amber-300/15 blur-[100px]"></div>
    </div>

    <header class="h-16 bg-white/80 backdrop-blur-md border-b border-slate-200/60 flex items-center px-6 relative z-10 space-x-3">
      <img src="/laskar-corps.png" class="h-8 w-auto object-contain" alt="HR Survey Logo" />
      <h1 class="text-lg font-bold text-slate-800">Employee Satisfaction Survey</h1>
    </header>
    <main class="flex-1 max-w-3xl w-full mx-auto p-6 relative z-10">
      <router-view></router-view>
    </main>
  </div>
</template>
