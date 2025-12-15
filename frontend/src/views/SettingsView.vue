<script setup>
import { ref, onMounted } from 'vue'

const isDarkMode = ref(false)

onMounted(() => {
  // Ambil status asli dari browser saat halaman dibuka
  isDarkMode.value = document.documentElement.classList.contains('dark')
})

const toggleTheme = () => {
  isDarkMode.value = !isDarkMode.value
  
  if (isDarkMode.value) {
    document.documentElement.classList.add('dark') // Nyalakan Gelap (Global)
    localStorage.setItem('theme', 'dark')
  } else {
    document.documentElement.classList.remove('dark') // Nyalakan Terang (Global)
    localStorage.setItem('theme', 'light')
  }
}
</script>

<template>
  <div class="max-w-2xl mx-auto animate-fade-in">
    <div class="bg-white dark:bg-[#1e293b] rounded-2xl shadow-sm border border-gray-200 dark:border-gray-700 overflow-hidden transition-colors duration-300">
      
      <div class="p-6 border-b border-gray-100 dark:border-gray-700">
        <h3 class="text-lg font-bold text-slate-800 dark:text-white">Pengaturan Tampilan</h3>
        <p class="text-sm text-slate-500 dark:text-slate-400">Sesuaikan kenyamanan mata Anda.</p>
      </div>

      <div class="p-6 space-y-6">
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-4">
            <div class="w-12 h-12 rounded-full flex items-center justify-center text-2xl transition-colors duration-300"
              :class="isDarkMode ? 'bg-indigo-900 text-indigo-300' : 'bg-orange-100 text-orange-500'">
              {{ isDarkMode ? '🌙' : '☀️' }}
            </div>
            <div>
              <p class="font-bold text-slate-700 dark:text-white">Mode Gelap (Dark Mode)</p>
              <p class="text-sm text-slate-500 dark:text-slate-400">
                {{ isDarkMode ? 'Aktif. Tampilan nyaman di malam hari.' : 'Non-aktif. Tampilan cerah.' }}
              </p>
            </div>
          </div>

          <button 
            @click="toggleTheme"
            class="w-14 h-8 rounded-full p-1 transition-colors duration-300 focus:outline-none shadow-inner relative"
            :class="isDarkMode ? 'bg-blue-600' : 'bg-slate-300'"
          >
            <div 
              class="w-6 h-6 bg-white rounded-full shadow-md transform transition-transform duration-300"
              :class="isDarkMode ? 'translate-x-6' : 'translate-x-0'"
            ></div>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>