<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router' // 1. Import Router
import { useI18n } from 'vue-i18n' 
import { 
  Bird, 
  Wallet, 
  HeartHandshake, 
  Egg, 
  TrendingUp, 
  Activity, 
  ArrowUpRight 
} from 'lucide-vue-next'

const router = useRouter() // 2. Inisialisasi Router
const { t } = useI18n()
const userId = localStorage.getItem('user_id')
const loading = ref(true)

// Data Kartu Statistik (Sekarang punya jalur 'path' navigasi)
const statsCards = ref([
  { 
    key: 'total_birds', 
    label: 'Total Populasi', 
    value: 0, 
    unit: 'Ekor', 
    icon: Bird, 
    colorClass: 'text-blue-600 bg-blue-50 border-blue-100 dark:bg-blue-900/20 dark:border-blue-900/50 dark:text-blue-400',
    path: '/dashboard/animals' // <-- Arah ke Stok Burung
  },
  { 
    key: 'total_sold', 
    label: 'Total Penjualan', 
    value: 0, 
    unit: 'Ekor', 
    icon: Wallet, 
    colorClass: 'text-emerald-600 bg-emerald-50 border-emerald-100 dark:bg-emerald-900/20 dark:border-emerald-900/50 dark:text-emerald-400',
    path: '/dashboard/finance' // <-- Arah ke Keuangan
  },
  { 
    key: 'active_pairs', 
    label: 'Pasangan Produktif', 
    value: 0, 
    unit: 'Pasang', 
    icon: HeartHandshake, 
    colorClass: 'text-rose-600 bg-rose-50 border-rose-100 dark:bg-rose-900/20 dark:border-rose-900/50 dark:text-rose-400',
    path: '/dashboard/pairs' // <-- Arah ke Pasangan
  },
  { 
    key: 'incubating_eggs', 
    label: 'Dalam Inkubasi', 
    value: 0, 
    unit: 'Butir', 
    icon: Egg, 
    colorClass: 'text-amber-600 bg-amber-50 border-amber-100 dark:bg-amber-900/20 dark:border-amber-900/50 dark:text-amber-400',
    path: '/dashboard/incubating' // <-- Arah ke Pengeraman
  }
])

const fetchStats = async () => {
  try {
    const response = await fetch(`http://localhost:8080/api/stats?user_id=${userId}`)
    const json = await response.json()
    
    if (json.data) {
        statsCards.value[0].value = json.data.total_birds || 0
        statsCards.value[1].value = json.data.total_sold || 0
        statsCards.value[2].value = json.data.active_pairs || 0
        statsCards.value[3].value = json.data.incubating_eggs || 0
    }
  } catch (error) {
    console.error("Gagal ambil statistik:", error)
  } finally {
    loading.value = false
  }
}

// Fungsi Navigasi
const goToPage = (path) => {
  if (path) router.push(path)
}

onMounted(() => {
  fetchStats()
})
</script>

<template>
  <div class="animate-fade-in space-y-8">
    
    <div class="flex flex-col md:flex-row md:items-end justify-between gap-4 border-b border-slate-200 dark:border-slate-800 pb-6">
        <div>
            <h2 class="text-3xl font-extrabold text-slate-900 dark:text-white flex items-center gap-3">
                <Activity class="w-8 h-8 text-blue-600"/> 
                Dashboard Overview
            </h2>
            <p class="text-slate-500 dark:text-slate-400 mt-2 text-sm">
                Pantau performa dan statistik peternakan Anda secara real-time.
            </p>
        </div>
        <div class="text-right hidden md:block">
            <span class="inline-flex items-center gap-1.5 px-3 py-1 rounded-full bg-slate-100 dark:bg-slate-800 text-xs font-bold text-slate-500 dark:text-slate-400">
                <span class="w-2 h-2 rounded-full bg-green-500 animate-pulse"></span>
                System Online
            </span>
        </div>
    </div>
    
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
      
      <div 
        v-for="(card, index) in statsCards" 
        :key="index"
        @click="goToPage(card.path)"
        class="group relative bg-white dark:bg-[#1e293b] p-6 rounded-2xl border border-slate-200 dark:border-slate-700 hover:border-blue-300 dark:hover:border-blue-700 transition-all duration-300 shadow-sm hover:shadow-md hover:-translate-y-1 cursor-pointer"
      >
        <div class="flex justify-between items-start mb-4">
            <div :class="`p-3 rounded-xl border ${card.colorClass} transition-transform group-hover:scale-110`">
                <component :is="card.icon" :size="24" stroke-width="2" />
            </div>
            <ArrowUpRight class="w-4 h-4 text-slate-300 dark:text-slate-600 group-hover:text-blue-500 transition-colors" />
        </div>
        
        <div>
          <h3 class="text-slate-500 dark:text-slate-400 text-xs font-bold uppercase tracking-wider">
             {{ card.label }}
          </h3>
          <div class="flex items-baseline gap-1 mt-1">
            <span class="text-3xl font-extrabold text-slate-900 dark:text-white tracking-tight">
                {{ loading ? '-' : card.value }}
            </span>
            <span class="text-xs font-medium text-slate-400">{{ card.unit }}</span>
          </div>
        </div>
      </div>

    </div>

    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <div class="bg-white dark:bg-[#1e293b] rounded-2xl border border-slate-200 dark:border-slate-700 overflow-hidden">
            <div class="p-5 border-b border-slate-100 dark:border-slate-800 flex justify-between items-center bg-slate-50/50 dark:bg-slate-800/50">
                <h4 class="font-bold text-slate-700 dark:text-white text-sm">Aktivitas Breeding</h4>
                <TrendingUp class="w-4 h-4 text-slate-400"/>
            </div>
            <div class="p-10 flex flex-col items-center justify-center text-center h-64">
                <div class="w-16 h-16 bg-slate-50 dark:bg-slate-800 rounded-full flex items-center justify-center mb-3">
                    <TrendingUp class="w-8 h-8 text-slate-300"/>
                </div>
                <p class="text-slate-400 text-sm font-medium">Grafik belum tersedia</p>
                <p class="text-xs text-slate-400 mt-1">Lakukan transaksi untuk melihat data.</p>
            </div>
        </div>

        <div class="bg-white dark:bg-[#1e293b] rounded-2xl border border-slate-200 dark:border-slate-700 overflow-hidden">
             <div class="p-5 border-b border-slate-100 dark:border-slate-800 flex justify-between items-center bg-slate-50/50 dark:bg-slate-800/50">
                <h4 class="font-bold text-slate-700 dark:text-white text-sm">Ringkasan Keuangan</h4>
                <Wallet class="w-4 h-4 text-slate-400"/>
            </div>
            <div class="p-10 flex flex-col items-center justify-center text-center h-64">
                <div class="w-16 h-16 bg-slate-50 dark:bg-slate-800 rounded-full flex items-center justify-center mb-3">
                    <Wallet class="w-8 h-8 text-slate-300"/>
                </div>
                <p class="text-slate-400 text-sm font-medium">Belum ada transaksi</p>
                <p class="text-xs text-slate-400 mt-1">Data keuangan akan muncul di sini.</p>
            </div>
        </div>
    </div>

  </div>
</template>