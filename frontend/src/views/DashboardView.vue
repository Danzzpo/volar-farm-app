<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'
// Import Icon
import { 
  LayoutDashboard, Bird, HeartHandshake, Egg, Hourglass, 
  Dna, CircleDot, Pill, Wallet, Settings, LogOut, User, Menu, X 
} from 'lucide-vue-next'

const router = useRouter()
const route = useRoute()

// Ambil nama farm, kasih default jika kosong
const farmName = localStorage.getItem('farm_name') || 'My Farm'
const isSidebarOpen = ref(false)

// Cek Mode Gelap saat pertama load
onMounted(() => {
  const savedTheme = localStorage.getItem('theme')
  const systemDark = window.matchMedia('(prefers-color-scheme: dark)').matches
  if (savedTheme === 'dark' || (!savedTheme && systemDark)) {
    document.documentElement.classList.add('dark')
  } else {
    document.documentElement.classList.remove('dark')
  }
})

// Definisi Menu
const menus = [
  { name: 'Dashboard', icon: LayoutDashboard, path: '/dashboard' }, // Halaman Utama Stats
  { name: 'Birds (Stok)', icon: Bird, path: '/dashboard/animals' },
  { name: 'Pairs (Jodohan)', icon: HeartHandshake, path: '/dashboard/pairs' },
  { name: 'Incubating', icon: Egg, path: '/dashboard/incubating' },
  { name: 'Pending Hatch', icon: Hourglass, path: '/dashboard/pending' },
  { name: 'GenCal', icon: Dna, path: '/dashboard/gencal' },
  { name: 'Rings', icon: CircleDot, path: '/dashboard/rings' },
  { name: 'Treatments', icon: Pill, path: '/dashboard/treatments' },
  { name: 'Finance', icon: Wallet, path: '/dashboard/finance' }, 
  { name: 'Settings', icon: Settings, path: '/dashboard/settings' },
]

// Judul Halaman Dinamis
const currentPageTitle = computed(() => {
  // Logic khusus: Jika path persis '/dashboard', judulnya 'Overview'
  if (route.path === '/dashboard') return 'Dashboard Overview'
  
  const activeMenu = menus.find(m => m.path === route.path)
  return activeMenu ? activeMenu.name : 'Dashboard'
})

const logout = () => {
  if(confirm('Yakin ingin keluar?')) {
    localStorage.clear()
    router.push('/login')
  }
}

// Tutup sidebar otomatis saat pindah halaman (Mobile)
router.afterEach(() => {
  isSidebarOpen.value = false
})
</script>

<template>
  <div class="flex h-screen font-sans overflow-hidden bg-slate-50 dark:bg-[#0f172a] text-slate-800 dark:text-slate-200 transition-colors duration-300">
    
    <div 
      v-if="isSidebarOpen" 
      class="fixed inset-0 bg-black/60 z-30 md:hidden backdrop-blur-sm transition-opacity"
      @click="isSidebarOpen = false"
    ></div>

    <aside 
      class="fixed inset-y-0 left-0 z-40 w-64 bg-slate-900 dark:bg-[#111827] text-slate-400 flex flex-col shadow-2xl transition-transform duration-300 ease-in-out md:static md:translate-x-0 border-r border-slate-800"
      :class="isSidebarOpen ? 'translate-x-0' : '-translate-x-full'"
    >
      <div class="p-6 flex items-center gap-3 border-b border-slate-800">
        <div class="w-10 h-10 bg-blue-600 rounded-xl flex items-center justify-center text-white font-bold shadow-lg shadow-blue-600/20">
           <Bird :size="24" stroke-width="2.5" />
        </div>
        <div class="overflow-hidden">
           <h1 class="font-bold text-lg text-white tracking-wide">Volar Farm</h1>
           <p class="text-[10px] text-blue-400 font-bold uppercase truncate">{{ farmName }}</p>
        </div>
        <button @click="isSidebarOpen = false" class="md:hidden ml-auto text-slate-500 hover:text-white"><X /></button>
      </div>

      <nav class="flex-1 overflow-y-auto p-4 space-y-1 no-scrollbar">
        <RouterLink v-for="menu in menus" :key="menu.path" :to="menu.path"
          class="flex items-center gap-3 px-4 py-3 rounded-xl text-sm font-medium transition-all duration-200 group relative overflow-hidden"
          :class="route.path === menu.path 
            ? 'bg-blue-600 text-white shadow-lg shadow-blue-900/50' 
            : 'hover:bg-slate-800 hover:text-white'"
        >
          <span v-if="route.path === menu.path" class="absolute left-0 top-1/2 -translate-y-1/2 h-6 w-1 bg-blue-300 rounded-r-full"></span>
          
          <component :is="menu.icon" :size="20" stroke-width="2" class="transition-transform group-hover:scale-110" />
          <span>{{ menu.name }}</span>
        </RouterLink>
      </nav>

      <div class="p-4 border-t border-slate-800">
        <button @click="logout" class="w-full flex items-center justify-center gap-2 px-4 py-3 rounded-xl bg-slate-800 hover:bg-red-600 hover:text-white text-slate-300 transition-all duration-300 font-bold text-sm shadow-lg">
           <LogOut :size="18" /> <span>Keluar</span>
        </button>
      </div>
    </aside>

    <main class="flex-1 flex flex-col min-w-0 overflow-hidden relative">
      
      <header class="bg-white dark:bg-[#1e293b] shadow-sm p-4 md:px-6 flex justify-between items-center z-10 border-b border-slate-200 dark:border-slate-800">
        <div class="flex items-center gap-4">
           <button @click="isSidebarOpen = true" class="md:hidden p-2 hover:bg-slate-100 dark:hover:bg-slate-700 rounded-lg"><Menu /></button>
           <div>
             <h2 class="text-xl font-bold truncate text-slate-800 dark:text-white">{{ currentPageTitle }}</h2>
           </div>
        </div>
        
        <div class="flex items-center gap-3">
           <div class="text-right hidden sm:block">
             <p class="text-sm font-bold text-slate-700 dark:text-slate-200">Halo, Admin</p>
             <p class="text-xs text-slate-400">Peternak</p>
           </div>
           <div class="w-10 h-10 rounded-full bg-slate-100 dark:bg-slate-700 border border-slate-200 dark:border-slate-600 flex items-center justify-center text-slate-500 dark:text-slate-300">
             <User :size="20" />
           </div>
        </div>
      </header>

      <div class="flex-1 overflow-y-auto p-4 md:p-6 bg-slate-50 dark:bg-[#0f172a] scroll-smooth">
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
/* Scrollbar custom agar rapi */
.no-scrollbar::-webkit-scrollbar { display: none; }
.no-scrollbar { -ms-overflow-style: none; scrollbar-width: none; }

/* Transisi Halaman yang Halus */
.fade-enter-active, .fade-leave-active { transition: opacity 0.2s ease, transform 0.2s ease; }
.fade-enter-from { opacity: 0; transform: translateY(5px); }
.fade-leave-to { opacity: 0; transform: translateY(-5px); }
</style>