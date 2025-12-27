<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n' // <--- Import i18n untuk terjemahan

// Import Icon
import { 
  LayoutDashboard, 
  Bird, 
  HeartHandshake, 
  Egg, 
  Dna, 
  Pill, 
  Wallet, 
  Settings, 
  LogOut,
  User,
  Menu, // Icon Menu untuk HP
  X     // Icon Close untuk HP
} from 'lucide-vue-next'

const router = useRouter()
const route = useRoute()
const { t } = useI18n() // <--- Gunakan fungsi translate

const farmName = localStorage.getItem('farm_name') || 'My Farm'

// State untuk Sidebar Mobile
const isSidebarOpen = ref(false)

onMounted(() => {
  const savedTheme = localStorage.getItem('theme')
  const systemDark = window.matchMedia('(prefers-color-scheme: dark)').matches
  if (savedTheme === 'dark' || (!savedTheme && systemDark)) {
    document.documentElement.classList.add('dark')
  } else {
    document.documentElement.classList.remove('dark')
  }
})

// Daftar Menu (Menggunakan computed agar reaktif saat ganti bahasa)
const menus = computed(() => [
  { name: t('menu.dashboard'), icon: LayoutDashboard, path: '/dashboard' },
  { name: t('menu.birds'), icon: Bird, path: '/dashboard/animals' },
  { name: t('menu.pairs'), icon: HeartHandshake, path: '/dashboard/pairs' },
  { name: t('menu.incubating'), icon: Egg, path: '/dashboard/incubating' },
  // Menu Pending Hatch & Rings sudah DIHAPUS
  { name: t('menu.gencal'), icon: Dna, path: '/dashboard/gencal' },
  { name: t('menu.treatments'), icon: Pill, path: '/dashboard/treatments' },
  { name: t('menu.finance'), icon: Wallet, path: '/dashboard/finance' }, 
  { name: t('menu.settings'), icon: Settings, path: '/dashboard/settings' },
])

const currentPageTitle = computed(() => {
  // Logic khusus: Jika path persis '/dashboard', judulnya 'Overview'
  if (route.path === '/dashboard') return 'Dashboard Overview'
  
  const activeMenu = menus.value.find(m => m.path === route.path)
  return activeMenu ? activeMenu.name : 'Dashboard'
})

const logout = () => {
  if(confirm(t('settings.logout') + '?')) { // Translate konfirmasi logout
    localStorage.clear()
    router.push('/login')
  }
}

// Tutup sidebar saat pindah halaman (khusus mobile)
router.afterEach(() => {
  isSidebarOpen.value = false
})
</script>

<template>
  <div class="flex h-screen font-sans overflow-hidden bg-slate-50 dark:bg-[#0f172a] transition-colors duration-300 text-slate-800 dark:text-slate-200">
    
    <div 
      v-if="isSidebarOpen" 
      class="fixed inset-0 bg-black/50 z-30 md:hidden backdrop-blur-sm transition-opacity"
      @click="isSidebarOpen = false"
    ></div>

    <aside 
      class="fixed inset-y-0 left-0 z-40 w-64 bg-slate-900 dark:bg-[#111827] text-slate-400 flex flex-col shadow-2xl transition-transform duration-300 ease-in-out md:static md:translate-x-0 border-r border-slate-800"
      :class="isSidebarOpen ? 'translate-x-0' : '-translate-x-full'"
    >
      <div class="p-6 mb-2 flex items-center justify-between border-b border-slate-800">
        <div class="flex items-center gap-3">
          <div class="w-10 h-10 bg-blue-600 rounded-xl flex items-center justify-center font-bold text-white shadow-lg shadow-blue-500/30">
            <Bird :size="24" stroke-width="2.5" /> 
          </div>
          <div class="overflow-hidden">
            <h1 class="font-bold text-lg tracking-wide text-white">Volar Farm</h1>
            <p class="text-[10px] text-blue-400 uppercase tracking-wider font-semibold truncate max-w-[100px]">{{ farmName }}</p>
          </div>
        </div>
        
        <button @click="isSidebarOpen = false" class="md:hidden text-slate-400 hover:text-white transition">
          <X :size="24" />
        </button>
      </div>

      <nav class="flex-1 overflow-y-auto px-4 space-y-2 py-4 no-scrollbar">
        <RouterLink 
          v-for="menu in menus" 
          :key="menu.path" 
          :to="menu.path"
          class="relative flex items-center gap-4 px-4 py-3 rounded-xl text-sm font-medium transition-all duration-200 group overflow-hidden"
          :class="route.path === menu.path 
            ? 'bg-blue-600 text-white shadow-lg shadow-blue-900/50' 
            : 'hover:bg-slate-800 hover:text-white'"
        >
          <span v-if="route.path === menu.path" class="absolute left-0 top-1/2 -translate-y-1/2 h-6 w-1 bg-blue-400 rounded-r-full"></span>
          
          <component :is="menu.icon" :size="20" stroke-width="2" class="relative z-10 transition-transform group-hover:scale-110" />
          <span class="relative z-10">{{ menu.name }}</span>
        </RouterLink>
      </nav>

      <div class="p-4 mt-auto border-t border-slate-800">
        <button @click="logout" class="w-full flex items-center justify-center gap-2 bg-slate-800 hover:bg-red-600 text-slate-300 hover:text-white py-3 rounded-xl text-sm font-bold transition-all duration-300 group shadow-lg">
          <LogOut :size="18" />
          <span>{{ t('settings.logout') }}</span>
        </button>
      </div>
    </aside>

    <main class="flex-1 flex flex-col min-w-0 overflow-hidden relative bg-slate-50 dark:bg-[#0f172a]">
      
      <header class="bg-white dark:bg-[#1e293b] shadow-sm dark:shadow-slate-900/50 p-4 md:p-5 flex justify-between items-center z-10 transition-colors duration-300 border-b border-slate-200 dark:border-gray-800 sticky top-0">
        
        <div class="flex items-center gap-3">
          <button 
            @click="isSidebarOpen = true" 
            class="md:hidden p-2 -ml-2 text-slate-600 dark:text-slate-300 hover:bg-slate-100 dark:hover:bg-slate-700 rounded-lg transition"
          >
            <Menu :size="24" />
          </button>

          <div>
            <h2 class="text-lg md:text-xl font-bold text-slate-800 dark:text-white truncate">{{ currentPageTitle }}</h2>
            <p class="hidden md:block text-xs text-slate-500 dark:text-slate-400">Overview & Statistik Farm</p>
          </div>
        </div>
        
        <div class="flex items-center gap-3">
           <div class="text-right hidden md:block">
             <p class="text-sm font-bold text-slate-700 dark:text-slate-200">Halo, Peternak</p>
             <p class="text-xs text-slate-400">Admin</p>
           </div>
           <div class="w-9 h-9 md:w-10 md:h-10 rounded-full bg-blue-100 dark:bg-slate-700 flex items-center justify-center text-blue-600 dark:text-blue-300 border border-blue-200 dark:border-transparent">
             <User :size="20" />
           </div>
        </div>
      </header>

      <div class="flex-1 overflow-y-auto p-4 md:p-6 scroll-smooth no-scrollbar bg-slate-50 dark:bg-[#0f172a] transition-colors duration-300">
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