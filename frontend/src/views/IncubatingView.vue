<script setup>
import { ref, onMounted } from 'vue'
// HANYA IMPORT ICON YANG PASTI AMAN
import { 
  Circle, Calendar, Clock, Activity, Plus, 
  ArrowRight, CheckCircle, Bird, X, Save 
} from 'lucide-vue-next'

const incubations = ref([])
const activePairs = ref([])
const loading = ref(true)
const isModalOpen = ref(false)
const isSubmitting = ref(false)

const userId = localStorage.getItem('user_id')

const form = ref({
  pair_id: '',
  start_date: '',
  egg_count: 0,
  notes: ''
})

// --- 1. FETCH DATA ASLI (DARI BACKEND GO) ---
const fetchIncubations = async () => {
  loading.value = true
  try {
    // Tembak API Go yang sudah kita perbaiki
    const res = await fetch(`http://localhost:8080/api/incubations?user_id=${userId}`)
    const json = await res.json()
    
    // Masukkan data asli dari database
    incubations.value = json.data || [] 
    
  } catch (err) {
    console.error("Gagal ambil data:", err)
  } finally {
    loading.value = false
  }
}

// --- 2. FETCH PASANGAN (DARI BACKEND GO) ---
const fetchActivePairs = async () => {
  try {
    const res = await fetch(`http://localhost:8080/api/pairs/active?user_id=${userId}`)
    const json = await res.json()
    
    // Format data agar sesuai dropdown
    activePairs.value = (json.data || []).map(p => ({
      id: p.id,
      cage: p.cage,
      visual: p.visual 
    }))
  } catch (err) {
    console.error(err)
  }
}

// --- LOGIKA HITUNG TANGGAL ---
const getHatchDate = (startDate) => {
  const date = new Date(startDate)
  date.setDate(date.getDate() + 21) 
  return date.toLocaleDateString('id-ID', { day: 'numeric', month: 'short' })
}

const getProgress = (startDate) => {
  const start = new Date(startDate)
  const end = new Date(startDate)
  end.setDate(end.getDate() + 21)
  
  const today = new Date()
  const totalDuration = 21 * 24 * 60 * 60 * 1000 
  const elapsed = today - start
  
  let percent = (elapsed / totalDuration) * 100
  if (percent > 100) percent = 100
  if (percent < 0) percent = 0
  
  const diffTime = end - today
  const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24)) 

  return { percent: Math.round(percent), daysLeft: diffDays }
}

// --- ACTIONS ---
const openModal = () => {
  fetchActivePairs()
  // Reset Form
  form.value = { 
    pair_id: '', 
    start_date: new Date().toISOString().split('T')[0], 
    egg_count: 1, 
    notes: '' 
  }
  isModalOpen.value = true
}

const saveIncubation = async () => {
  isSubmitting.value = true
  try {
    // Kirim data ke Backend
    const payload = {
      pair_id: parseInt(form.value.pair_id),
      start_date: form.value.start_date, // Format YYYY-MM-DD
      egg_count: form.value.egg_count,
      notes: form.value.notes,
      user_id: parseInt(userId)
    }

    const res = await fetch('http://localhost:8080/api/incubations', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload)
    })

    const json = await res.json()

    if (res.ok) {
        alert("✅ Data Pengeraman Disimpan!")
        isModalOpen.value = false
        fetchIncubations() // Refresh data otomatis
    } else {
        alert("❌ Gagal: " + (json.error || "Terjadi kesalahan"))
    }
  } catch (err) {
    alert("Error sistem: " + err.message)
  } finally {
    isSubmitting.value = false
  }
}

onMounted(() => {
  fetchIncubations()
})
</script>

<template>
  <div class="space-y-6 p-2">
    
    <div class="flex flex-col md:flex-row md:items-end justify-between gap-4">
      <div>
        <h2 class="text-3xl font-extrabold text-slate-800 dark:text-white flex items-center gap-3">
          <Circle class="w-8 h-8 text-yellow-500" /> Incubating
        </h2>
        <p class="text-slate-500 mt-1">Pantau proses pengeraman telur.</p>
      </div>
      
      <div class="flex gap-2">
        <button @click="openModal" class="bg-yellow-500 hover:bg-yellow-600 text-white px-5 py-2.5 rounded-xl font-bold flex items-center gap-2 shadow-lg shadow-yellow-500/20 active:scale-95 transition">
          <Plus class="w-5 h-5" /> <span>Input Telur</span>
        </button>
      </div>
    </div>

    <div v-if="loading" class="text-center py-10 text-slate-400">Loading data...</div>

    <div v-else-if="incubations.length === 0" class="text-center py-20 border-2 border-dashed border-slate-300 rounded-2xl">
       <Circle class="w-12 h-12 text-slate-300 mx-auto mb-2" />
       <p class="text-slate-500">Belum ada telur yang dierami.</p>
       <button @click="openModal" class="text-yellow-600 font-bold hover:underline mt-2">Mulai Mengeram</button>
    </div>

    <div v-else class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-6">
      
      <div 
        v-for="item in incubations" 
        :key="item.id"
        class="bg-white dark:bg-slate-800 rounded-2xl border border-slate-200 dark:border-slate-700 shadow-sm hover:shadow-md transition group overflow-hidden"
      >
        <div class="p-5 border-b border-slate-100 dark:border-slate-700 flex justify-between items-start bg-slate-50 dark:bg-slate-900/50">
          <div>
            <span class="text-[10px] font-bold uppercase tracking-wider text-slate-500 flex items-center gap-1 mb-1">
              <Bird class="w-3 h-3" /> {{ item.cage }}
            </span>
            <h3 class="font-bold text-slate-800 dark:text-white text-lg">{{ item.pair_visual }}</h3>
          </div>
          <div class="bg-yellow-100 dark:bg-yellow-900/30 text-yellow-700 dark:text-yellow-400 px-3 py-1 rounded-lg text-xs font-bold border border-yellow-200 dark:border-yellow-800">
             {{ item.egg_count }} Telur
          </div>
        </div>

        <div class="p-5">
          <div class="mb-4">
            <div class="flex justify-between text-xs mb-1.5">
              <span class="text-slate-500 font-medium">Proses Inkubasi</span>
              <span class="font-bold" :class="getProgress(item.start_date).percent >= 90 ? 'text-red-500' : 'text-blue-500'">
                {{ getProgress(item.start_date).percent }}%
              </span>
            </div>
            <div class="w-full bg-slate-100 dark:bg-slate-700 rounded-full h-2.5 overflow-hidden">
              <div 
                class="h-2.5 rounded-full transition-all duration-1000"
                :class="getProgress(item.start_date).percent >= 90 ? 'bg-red-500' : 'bg-blue-500'"
                :style="{ width: getProgress(item.start_date).percent + '%' }"
              ></div>
            </div>
          </div>

          <div class="grid grid-cols-2 gap-4 mt-4">
            <div class="bg-slate-50 dark:bg-slate-700/30 p-3 rounded-xl">
              <p class="text-[10px] text-slate-400 uppercase font-bold mb-1">Mulai</p>
              <div class="flex items-center gap-2 text-slate-700 dark:text-slate-200 font-bold text-sm">
                <Calendar class="w-4 h-4 text-slate-400" />
                {{ new Date(item.start_date).toLocaleDateString('id-ID', {day: 'numeric', month: 'short'}) }}
              </div>
            </div>
            <div class="bg-slate-50 dark:bg-slate-700/30 p-3 rounded-xl border border-slate-100 dark:border-transparent" 
                 :class="getProgress(item.start_date).daysLeft <= 3 ? 'bg-red-50 dark:bg-red-900/20 border-red-100' : ''">
              <p class="text-[10px] uppercase font-bold mb-1" :class="getProgress(item.start_date).daysLeft <= 3 ? 'text-red-500' : 'text-slate-400'">
                Estimasi
              </p>
              <div class="flex items-center gap-2 font-bold text-sm" :class="getProgress(item.start_date).daysLeft <= 3 ? 'text-red-600' : 'text-slate-700 dark:text-slate-200'">
                <Clock class="w-4 h-4" />
                <span>{{ getHatchDate(item.start_date) }}</span>
              </div>
            </div>
          </div>

          <div v-if="getProgress(item.start_date).daysLeft <= 0" class="mt-4 p-2 bg-green-100 text-green-700 text-xs font-bold text-center rounded-lg">
            🐣 Waktunya Menetas!
          </div>
          <div v-else-if="getProgress(item.start_date).daysLeft <= 3" class="mt-4 p-2 bg-red-100 text-red-700 text-xs font-bold text-center rounded-lg">
            ⚠️ Siaga! Menetas dalam {{ getProgress(item.start_date).daysLeft }} hari.
          </div>
        </div>

        <div class="p-3 border-t border-slate-100 dark:border-slate-700 grid grid-cols-2 gap-2">
           <button class="py-2 text-xs font-bold text-slate-500 hover:bg-slate-50 dark:hover:bg-slate-700 rounded-lg transition">Edit</button>
           <button class="py-2 text-xs font-bold text-blue-600 hover:bg-blue-50 dark:hover:bg-blue-900/30 rounded-lg transition">Cek Senter</button>
        </div>

      </div>
    </div>

    <div v-if="isModalOpen" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/60 backdrop-blur-sm" @click.self="isModalOpen = false">
      <div class="bg-white dark:bg-slate-900 w-full max-w-md rounded-2xl shadow-2xl overflow-hidden">
        
        <div class="px-6 py-4 border-b border-slate-100 dark:border-slate-800 flex justify-between items-center bg-slate-50 dark:bg-slate-900">
          <h3 class="font-bold text-lg text-slate-800 dark:text-white flex items-center gap-2">
            <Plus class="w-5 h-5 text-yellow-500"/> Input Pengeraman
          </h3>
          <button @click="isModalOpen = false"><X class="w-5 h-5 text-slate-400"/></button>
        </div>

        <div class="p-6 space-y-4">
          <div>
            <label class="block text-xs font-bold text-slate-500 mb-1.5 uppercase">Pilih Pasangan</label>
            <div class="relative">
              <select v-model="form.pair_id" class="w-full p-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl text-sm outline-none dark:text-white font-bold">
                <option value="" disabled>-- Pilih --</option>
                <option v-for="pair in activePairs" :key="pair.id" :value="pair.id">
                  {{ pair.cage }} - {{ pair.visual }}
                </option>
              </select>
              <ArrowRight class="absolute right-3 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400 pointer-events-none" />
            </div>
            <p v-if="activePairs.length === 0" class="text-[10px] text-red-500 mt-1">
               * Belum ada Pasangan Aktif. Silakan input di menu "Pairs" dulu.
            </p>
          </div>

          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="block text-xs font-bold text-slate-500 mb-1.5 uppercase">Tgl Mulai</label>
              <input v-model="form.start_date" type="date" class="w-full p-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl text-sm outline-none dark:text-white">
            </div>
            <div>
              <label class="block text-xs font-bold text-slate-500 mb-1.5 uppercase">Jml Telur</label>
              <input v-model="form.egg_count" type="number" class="w-full p-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl text-sm outline-none dark:text-white font-bold text-center">
            </div>
          </div>
        </div>

        <div class="px-6 py-4 bg-slate-50 dark:bg-slate-900 border-t border-slate-100 dark:border-slate-800 flex gap-3">
          <button @click="isModalOpen = false" class="flex-1 py-3 rounded-xl border border-slate-300 font-bold text-sm text-slate-500">Batal</button>
          <button @click="saveIncubation" :disabled="isSubmitting" class="flex-1 py-3 rounded-xl bg-yellow-500 text-white font-bold text-sm hover:bg-yellow-600 transition shadow-lg shadow-yellow-500/30 flex items-center justify-center gap-2">
            <Save v-if="!isSubmitting" class="w-5 h-5" />
            <span>{{ isSubmitting ? 'Simpan...' : 'Simpan' }}</span>
          </button>
        </div>

      </div>
    </div>

  </div>
</template>