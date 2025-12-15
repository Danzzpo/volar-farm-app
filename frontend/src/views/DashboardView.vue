<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'

const router = useRouter()
const route = useRoute()
const farmName = localStorage.getItem('farm_name')

// --- LOGIKA INISIALISASI TEMA ---
onMounted(() => {
  const savedTheme = localStorage.getItem('theme')
  const systemDark = window.matchMedia('(prefers-color-scheme: dark)').matches
  
  if (savedTheme === 'dark' || (!savedTheme && systemDark)) {
    document.documentElement.classList.add('dark')
  } else {
    document.documentElement.classList.remove('dark')
  }
})

const menus = [
  { name: 'Dashboard', icon: '🏠', path: '/dashboard' },
  { name: 'Birds (Stok)', icon: '🐦', path: '/dashboard/animals' },
  { name: 'Pairs (Jodohan)', icon: '💞', path: '/dashboard/pairs' },
  { name: 'Incubating', icon: '🥚', path: '/dashboard/incubating' },
  { name: 'Pending Hatch', icon: '⏳', path: '/dashboard/pending' },
  { name: 'GenCal', icon: '🧬', path: '/dashboard/gencal' },
  { name: 'Rings', icon: '💍', path: '/dashboard/rings' },
  { name: 'Treatments', icon: '💊', path: '/dashboard/treatments' },
  { name: 'Finance', icon: '💰', path: '/dashboard/finance' }, 
  { name: 'Settings', icon: '⚙️', path: '/dashboard/settings' },
]

const currentPageTitle = computed(() => {
  const activeMenu = menus.find(m => m.path === route.path)
  return activeMenu ? activeMenu.name : 'Dashboard'
})

const logout = () => {
  localStorage.clear()
  router.push('/login')
}
</script>

<template>
  <div class="flex h-screen font-sans overflow-hidden bg-white dark:bg-[#0f172a] transition-colors duration-300 text-slate-800 dark:text-slate-200">
    
    <aside class="w-64 bg-slate-100 dark:bg-[#111827] text-slate-600 dark:text-gray-400 flex flex-col shadow-xl z-20 transition-all duration-300 border-r border-slate-200 dark:border-gray-800">
      
      <div class="p-6 mb-2 flex items-center gap-3">
        <div class="w-10 h-10 bg-blue-600 rounded-xl flex items-center justify-center font-bold text-white text-xl shadow-lg shadow-blue-500/30">
          V
        </div>
        <div>
          <h1 class="font-bold text-lg tracking-wide text-slate-800 dark:text-white">Volar Farm</h1>
          <p class="text-[10px] text-blue-500 dark:text-blue-400 uppercase tracking-wider font-semibold">{{ farmName }}</p>
        </div>
      </div>

      <nav class="flex-1 overflow-y-auto px-4 space-y-2 no-scrollbar">
        <RouterLink 
          v-for="menu in menus" 
          :key="menu.path" 
          :to="menu.path"
          class="relative flex items-center gap-4 px-4 py-3 rounded-xl text-sm font-medium transition-all duration-200 group overflow-hidden"
          :class="route.path === menu.path 
            ? 'bg-blue-600 text-white shadow-lg shadow-blue-500/30' 
            : 'hover:bg-white dark:hover:bg-[#1f2937] hover:text-blue-600 dark:hover:text-white hover:shadow-sm'"
        >
          <span v-if="route.path === menu.path" class="absolute left-0 top-1/2 -translate-y-1/2 h-6 w-1 bg-blue-300 rounded-r-full"></span>
          <span class="text-xl relative z-10 transition-transform group-hover:scale-110">{{ menu.icon }}</span>
          <span class="relative z-10">{{ menu.name }}</span>
        </RouterLink>
      </nav>

      <div class="p-4 mt-auto border-t border-slate-200 dark:border-gray-800">
        <button @click="logout" class="w-full flex items-center justify-center gap-2 bg-white dark:bg-[#374151] hover:bg-red-600 text-slate-600 dark:text-gray-300 hover:text-white py-3 rounded-xl text-sm font-bold transition-all duration-300 shadow-sm border border-slate-200 dark:border-transparent group">
          <span>🚪</span> Keluar
        </button>
      </div>
    </aside>

    <main class="flex-1 flex flex-col min-w-0 overflow-hidden relative bg-white dark:bg-[#0f172a]">
      
      <header class="bg-white dark:bg-[#1e293b] shadow-sm dark:shadow-slate-900/50 p-5 flex justify-between items-center z-10 transition-colors duration-300 border-b border-slate-100 dark:border-gray-800">
        <div>
          <h2 class="text-xl font-bold text-slate-800 dark:text-white">{{ currentPageTitle }}</h2>
          <p class="text-xs text-slate-500 dark:text-slate-400">Overview & Statistik Farm</p>
        </div>
        
        <div class="flex items-center gap-3">
           <div class="text-right hidden md:block">
              <p class="text-sm font-bold text-slate-700 dark:text-slate-200">Halo, Peternak</p>
              <p class="text-xs text-slate-400">Admin</p>
           </div>
           <div class="w-10 h-10 rounded-full bg-slate-100 dark:bg-slate-700 flex items-center justify-center text-xl">👨‍🌾</div>
        </div>
      </header>

      <div class="flex-1 overflow-y-auto p-6 scroll-smooth no-scrollbar bg-white dark:bg-[#0f172a] transition-colors duration-300">
        <RouterView v-slot="{ Component }">
          <transition name="fade" mode="out-in">
            <component :is="Component" />
          </transition>
        </RouterView>
      </div>

    </main>
  </div>
</template>

<style scoped>
.no-scrollbar::-webkit-scrollbar { display: none; }
.no-scrollbar { -ms-overflow-style: none; scrollbar-width: none; }
.fade-enter-active, .fade-leave-active { transition: opacity 0.2s ease, transform 0.2s ease; }
.fade-enter-from { opacity: 0; transform: translateY(10px); }
.fade-leave-to { opacity: 0; transform: translateY(-10px); }
</style>