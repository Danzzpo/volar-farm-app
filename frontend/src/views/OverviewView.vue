<script setup>
import { ref, onMounted } from 'vue'

const userId = localStorage.getItem('user_id')
const stats = ref({ total_birds: 0, total_sold: 0 })
const loading = ref(true)

const fetchStats = async () => {
  try {
    const response = await fetch(`http://localhost:8080/api/stats?user_id=${userId}`)
    const json = await response.json()
    stats.value = json
  } catch (error) {
    console.error("Gagal ambil statistik:", error)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchStats()
})
</script>

<template>
  <div class="animate-fade-in">
    <h2 class="text-2xl font-bold text-slate-800 dark:text-white mb-6">Ringkasan Farm</h2>
    
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
      
      <div class="bg-white dark:bg-[#1e293b] p-6 rounded-2xl shadow-sm border border-slate-100 dark:border-slate-700 flex items-center gap-4 transition-colors duration-300">
        <div class="w-12 h-12 rounded-full bg-blue-100 dark:bg-blue-900/30 text-blue-600 dark:text-blue-400 flex items-center justify-center text-2xl">🐦</div>
        <div>
          <p class="text-slate-400 dark:text-slate-400 text-sm">Stok Burung</p>
          <p class="text-2xl font-bold text-slate-700 dark:text-white">
            {{ loading ? '...' : stats.total_birds }} 
            <span class="text-xs text-green-500 font-normal">Ekor</span>
          </p>
        </div>
      </div>

      <div class="bg-white dark:bg-[#1e293b] p-6 rounded-2xl shadow-sm border border-slate-100 dark:border-slate-700 flex items-center gap-4 transition-colors duration-300">
        <div class="w-12 h-12 rounded-full bg-green-100 dark:bg-green-900/30 text-green-600 dark:text-green-400 flex items-center justify-center text-2xl">💰</div>
        <div>
          <p class="text-slate-400 dark:text-slate-400 text-sm">Total Terjual</p>
          <p class="text-2xl font-bold text-slate-700 dark:text-white">
            {{ loading ? '...' : stats.total_sold }}
            <span class="text-xs text-slate-500 font-normal">Ekor</span>
          </p>
        </div>
      </div>

      <div class="bg-white dark:bg-[#1e293b] p-6 rounded-2xl shadow-sm border border-slate-100 dark:border-slate-700 flex items-center gap-4 opacity-60 transition-colors duration-300">
        <div class="w-12 h-12 rounded-full bg-pink-100 dark:bg-pink-900/30 text-pink-600 dark:text-pink-400 flex items-center justify-center text-2xl">💞</div>
        <div>
          <p class="text-slate-400 dark:text-slate-400 text-sm">Pasangan</p>
          <p class="text-2xl font-bold text-slate-700 dark:text-white">0 <span class="text-xs text-slate-500 font-normal">Pasang</span></p>
        </div>
      </div>

      <div class="bg-white dark:bg-[#1e293b] p-6 rounded-2xl shadow-sm border border-slate-100 dark:border-slate-700 flex items-center gap-4 opacity-60 transition-colors duration-300">
        <div class="w-12 h-12 rounded-full bg-yellow-100 dark:bg-yellow-900/30 text-yellow-600 dark:text-yellow-400 flex items-center justify-center text-2xl">🥚</div>
        <div>
          <p class="text-slate-400 dark:text-slate-400 text-sm">Sedang Mengeram</p>
          <p class="text-2xl font-bold text-slate-700 dark:text-white">0 <span class="text-xs text-slate-500 font-normal">Sarang</span></p>
        </div>
      </div>

    </div>

    <div class="bg-white dark:bg-[#1e293b] p-8 rounded-xl shadow-sm border border-slate-100 dark:border-slate-700 text-center text-slate-400 h-64 flex flex-col items-center justify-center gap-2 transition-colors duration-300">
      <span class="text-4xl">📈</span>
      <p>Grafik aktivitas breeding akan muncul di sini.</p>
    </div>
  </div>
</template>