<script setup>
import { ref, onMounted, nextTick } from 'vue'
import { 
  Circle, Clock, Plus, Minus, // Calendar dihapus dari import
  ArrowRight, Bird, X, Save, Egg, Search, ChevronDown, Check 
} from 'lucide-vue-next'

const incubations = ref([])
const activePairs = ref([])
const loading = ref(true)
const isModalOpen = ref(false)
const isSubmitting = ref(false)

const userId = localStorage.getItem('user_id')

const isDropdownOpen = ref(false)
const searchQuery = ref('')
const selectedPairLabel = ref('')
const dropdownRef = ref(null)

const form = ref({
  pair_id: '',
  start_date: '',
  egg_count: 1,
  notes: ''
})

const fetchIncubations = async () => {
  loading.value = true
  try {
    const res = await fetch(`http://localhost:8080/api/incubations?user_id=${userId}`)
    const json = await res.json()
    incubations.value = json.data || [] 
  } catch (err) { console.error(err) } finally { loading.value = false }
}

const fetchActivePairs = async () => {
  try {
    const res = await fetch(`http://localhost:8080/api/pairs/active?user_id=${userId}`)
    const json = await res.json()
    activePairs.value = (json.data || []).map(p => ({
      id: p.id, cage: p.cage, visual: p.visual 
    }))
  } catch (err) { console.error(err) }
}

const getHatchDate = (startDate) => {
  const date = new Date(startDate)
  date.setDate(date.getDate() + 21) 
  return date.toLocaleDateString('id-ID', { day: 'numeric', month: 'short', year: 'numeric' })
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
  const diffDays = Math.ceil((end - today) / (1000 * 60 * 60 * 24)) 
  return { percent: Math.round(percent), daysLeft: diffDays }
}

const openModal = () => {
  fetchActivePairs()
  form.value = { pair_id: '', start_date: new Date().toISOString().split('T')[0], egg_count: 1, notes: '' }
  selectedPairLabel.value = ''
  searchQuery.value = ''
  isModalOpen.value = true
}

const filteredPairs = () => {
  if (!searchQuery.value) return activePairs.value
  const lower = searchQuery.value.toLowerCase()
  return activePairs.value.filter(p => p.cage.toLowerCase().includes(lower) || p.visual.toLowerCase().includes(lower))
}

const selectPair = (pair) => {
  form.value.pair_id = pair.id
  selectedPairLabel.value = `${pair.cage} - ${pair.visual}`
  isDropdownOpen.value = false
  searchQuery.value = ''
}

const toggleDropdown = () => {
  isDropdownOpen.value = !isDropdownOpen.value
  if (isDropdownOpen.value) nextTick(() => dropdownRef.value?.focus())
}

const incrementEgg = () => form.value.egg_count++
const decrementEgg = () => { if(form.value.egg_count > 1) form.value.egg_count-- }

const saveIncubation = async () => {
  if (!form.value.pair_id) return alert("Pilih pasangan dulu!")
  isSubmitting.value = true
  try {
    const payload = {
      pair_id: parseInt(form.value.pair_id),
      start_date: form.value.start_date,
      egg_count: parseInt(form.value.egg_count),
      notes: form.value.notes,
      user_id: parseInt(userId)
    }
    const res = await fetch('http://localhost:8080/api/incubations', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload)
    })
    if (res.ok) {
        alert("✅ Data Disimpan!")
        isModalOpen.value = false
        fetchIncubations()
    } else {
        alert("❌ Gagal Simpan")
    }
  } catch (err) { console.error(err) } finally { isSubmitting.value = false }
}

onMounted(() => fetchIncubations())
</script>

<template>
  <div class="space-y-6 p-2">
    <div class="flex flex-col md:flex-row md:items-end justify-between gap-4">
      <div>
        <h2 class="text-3xl font-extrabold text-slate-800 dark:text-white flex items-center gap-3">
          <Egg class="w-8 h-8 text-yellow-500" /> Inkubasi
        </h2>
        <p class="text-slate-500 mt-1 text-sm md:text-base">Pantau proses pengeraman telur aktif.</p>
      </div>
      <div class="flex gap-2">
        <button @click="openModal" class="bg-gradient-to-r from-yellow-500 to-amber-500 hover:from-yellow-600 hover:to-amber-600 text-white px-5 py-2.5 rounded-xl font-bold flex items-center gap-2 shadow-lg shadow-yellow-500/20 active:scale-95 transition">
          <Plus class="w-5 h-5" /> <span>Input Telur</span>
        </button>
      </div>
    </div>

    <div v-if="loading" class="text-center py-12 text-slate-400">Memuat data...</div>
    <div v-else-if="incubations.length === 0" class="text-center py-20 border-2 border-dashed border-slate-300 dark:border-slate-700 rounded-3xl bg-slate-50 dark:bg-slate-800/50">
       <Egg class="w-16 h-16 text-slate-300 mx-auto mb-4" />
       <h3 class="text-lg font-bold text-slate-600 dark:text-slate-300">Belum ada telur</h3>
       <button @click="openModal" class="text-yellow-600 font-bold hover:underline">Mulai Mengeram</button>
    </div>

    <div v-else class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-6">
      <div v-for="item in incubations" :key="item.id" class="bg-white dark:bg-slate-800 rounded-2xl border border-slate-100 dark:border-slate-700 shadow-sm hover:shadow-xl transition-all duration-300 group overflow-hidden flex flex-col">
        <div class="p-5 border-b border-slate-100 dark:border-slate-700 flex justify-between items-start bg-slate-50/50 dark:bg-slate-900/30">
          <div class="pr-2">
            <span class="inline-flex items-center gap-1.5 px-2 py-1 rounded-md bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-600 text-[11px] font-bold uppercase tracking-wider text-slate-500 mb-2">
              <Bird class="w-3 h-3" /> {{ item.cage }}
            </span>
            <h3 class="font-bold text-slate-800 dark:text-white text-lg leading-tight line-clamp-2">{{ item.pair_visual }}</h3>
          </div>
          <div class="flex-shrink-0 bg-yellow-100 dark:bg-yellow-900/30 text-yellow-700 dark:text-yellow-400 px-3 py-1.5 rounded-xl text-xs font-bold border border-yellow-200 dark:border-yellow-800 flex items-center gap-1">
             <Egg class="w-3.5 h-3.5" /> {{ item.egg_count }}
          </div>
        </div>

        <div class="p-5 flex-1">
          <div class="mb-5">
            <div class="flex justify-between text-xs mb-2">
              <span class="text-slate-500 font-bold uppercase tracking-wide">Hari Ke-{{ 21 - getProgress(item.start_date).daysLeft }}</span>
              <span class="font-bold text-blue-500">{{ getProgress(item.start_date).percent }}%</span>
            </div>
            <div class="w-full bg-slate-100 dark:bg-slate-700 rounded-full h-3 overflow-hidden">
              <div class="h-full rounded-full bg-gradient-to-r from-blue-400 to-blue-500" :style="{ width: getProgress(item.start_date).percent + '%' }"></div>
            </div>
          </div>
          <div class="grid grid-cols-2 gap-4">
            <div class="bg-slate-50 dark:bg-slate-700/30 p-3 rounded-xl border border-slate-100 dark:border-slate-700">
              <p class="text-[10px] text-slate-400 uppercase font-bold mb-1">Mulai</p>
              <div class="flex items-center gap-2 text-slate-700 dark:text-slate-200 font-bold text-sm"><Calendar class="w-4 h-4 text-slate-400" /> {{ new Date(item.start_date).toLocaleDateString('id-ID', {day: 'numeric', month: 'short'}) }}</div>
            </div>
            <div class="bg-slate-50 dark:bg-slate-700/30 p-3 rounded-xl border border-slate-100 dark:border-slate-700">
              <p class="text-[10px] text-slate-400 uppercase font-bold mb-1">Estimasi</p>
              <div class="flex items-center gap-2 text-slate-700 dark:text-slate-200 font-bold text-sm"><Clock class="w-4 h-4 text-slate-400" /> {{ getHatchDate(item.start_date) }}</div>
            </div>
          </div>
        </div>
        
        <div class="p-4 bg-slate-50 dark:bg-slate-900/50 border-t border-slate-100 dark:border-slate-700 grid grid-cols-2 gap-3">
           <button class="py-2.5 text-xs font-bold text-slate-600 bg-white dark:bg-slate-800 hover:bg-slate-100 border border-slate-200 rounded-lg shadow-sm">Edit Data</button>
           <button class="py-2.5 text-xs font-bold text-blue-600 bg-blue-50 hover:bg-blue-100 border border-blue-100 rounded-lg shadow-sm">Senter Telur</button>
        </div>
      </div>
    </div>

    <div v-if="isModalOpen" class="fixed inset-0 z-50 h-screen w-screen overflow-hidden bg-slate-900/70 backdrop-blur-sm flex items-center justify-center p-4" @click.self="isModalOpen = false">
      
      <div class="bg-white dark:bg-slate-900 w-full max-w-lg rounded-2xl shadow-2xl animate-fade-in-up flex flex-col relative">
        
        <div class="px-6 py-5 border-b border-slate-100 dark:border-slate-800 flex justify-between items-center bg-slate-50 dark:bg-slate-900 shrink-0 rounded-t-2xl">
          <h3 class="font-bold text-xl text-slate-800 dark:text-white flex items-center gap-2"><Plus class="w-6 h-6 text-yellow-500"/> Input Pengeraman</h3>
          <button @click="isModalOpen = false" class="p-2 rounded-full hover:bg-slate-200 dark:hover:bg-slate-800 text-slate-500 transition"><X class="w-5 h-5"/></button>
        </div>

        <div class="p-7 space-y-6 bg-white dark:bg-slate-900">
          
          <div class="relative w-full">
            <label class="form-label">Pilih Pasangan</label>
            <button @click="toggleDropdown" type="button" class="w-full p-3.5 bg-white dark:bg-slate-800 border border-slate-300 dark:border-slate-600 rounded-xl text-left text-base flex justify-between items-center focus:ring-2 focus:ring-yellow-500 transition shadow-sm">
              <span class="truncate pr-2 font-bold text-slate-700 dark:text-white">{{ selectedPairLabel || '-- Pilih Pasangan Aktif --' }}</span>
              <ChevronDown class="w-5 h-5 text-slate-500 transition-transform" :class="{'rotate-180': isDropdownOpen}"/>
            </button>

            <div v-if="isDropdownOpen" class="absolute z-50 mt-2 w-full bg-white dark:bg-slate-800 rounded-xl shadow-2xl border border-slate-200 dark:border-slate-600 overflow-hidden animate-in fade-in zoom-in-95 origin-top-left">
              <div class="p-3 border-b border-slate-100 bg-slate-50 dark:bg-slate-900">
                <div class="relative">
                  <Search class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400"/>
                  <input v-model="searchQuery" ref="dropdownRef" type="text" placeholder="Cari Kandang..." class="w-full pl-10 pr-4 py-2 bg-white dark:bg-slate-700 border border-slate-200 dark:border-slate-600 rounded-lg text-sm focus:ring-2 focus:ring-yellow-500 outline-none dark:text-white">
                </div>
              </div>
              <ul class="max-h-48 overflow-y-auto custom-scrollbar p-1">
                <li v-if="filteredPairs().length === 0" class="p-3 text-center text-sm text-slate-500">Tidak ditemukan.</li>
                <li v-for="pair in filteredPairs()" :key="pair.id" @click="selectPair(pair)" class="flex justify-between items-center px-4 py-3 hover:bg-yellow-50 dark:hover:bg-slate-700 rounded-lg cursor-pointer group">
                  <div class="font-bold text-slate-700 dark:text-white">{{ pair.cage }} <span class="font-normal text-slate-500 ml-1 text-xs">({{ pair.visual }})</span></div>
                  <Check v-if="form.pair_id === pair.id" class="w-4 h-4 text-yellow-500"/>
                </li>
              </ul>
            </div>
          </div>

          <div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
            <div>
              <label class="form-label">Tgl Mulai</label>
              <input 
                v-model="form.start_date" 
                type="date" 
                class="input-field cursor-pointer"
              >
            </div>
            
            <div>
              <label class="form-label">Jumlah Telur</label>
              <div class="flex items-center h-[52px]">
                 <button 
                   @click="decrementEgg"
                   type="button" 
                   class="h-full w-14 bg-slate-100 dark:bg-slate-800 border border-slate-300 dark:border-slate-600 rounded-l-xl flex items-center justify-center text-slate-500 hover:bg-slate-200 dark:hover:bg-slate-700 active:scale-95 transition"
                 >
                   <Minus class="w-5 h-5"/>
                 </button>
                 
                 <div class="h-full flex-1 border-y border-slate-300 dark:border-slate-600 bg-white dark:bg-slate-800 flex items-center justify-center relative">
                    <Egg class="w-4 h-4 text-slate-400 absolute left-4"/>
                    <input 
                      v-model="form.egg_count" 
                      type="number" 
                      min="1" 
                      class="w-full text-center font-bold text-lg text-slate-800 dark:text-white outline-none bg-transparent pl-4"
                    >
                 </div>

                 <button 
                   @click="incrementEgg"
                   type="button" 
                   class="h-full w-14 bg-yellow-500 border border-yellow-600 rounded-r-xl flex items-center justify-center text-white hover:bg-yellow-600 active:scale-95 transition shadow-sm"
                 >
                   <Plus class="w-5 h-5"/>
                 </button>
              </div>
            </div>
          </div>

          <div>
            <label class="form-label">Catatan</label>
            <textarea 
              v-model="form.notes" 
              rows="3" 
              placeholder="Kondisi telur, dll..."
              class="input-field resize-none custom-scrollbar">
            </textarea>
          </div>

        </div>

        <div class="px-6 py-5 bg-slate-50 dark:bg-slate-900 border-t border-slate-100 dark:border-slate-800 flex gap-4 shrink-0 rounded-b-2xl">
          <button @click="isModalOpen = false" class="flex-1 py-3.5 rounded-xl border border-slate-300 font-bold text-base text-slate-600 hover:bg-slate-100">Batal</button>
          <button @click="saveIncubation" :disabled="isSubmitting" class="flex-1 py-3.5 rounded-xl bg-gradient-to-r from-yellow-500 to-amber-500 text-white font-bold text-base hover:from-yellow-600 hover:to-amber-600 flex items-center justify-center gap-2 shadow-lg">
            <Save v-if="!isSubmitting" class="w-5 h-5" /> <span>Simpan</span>
          </button>
        </div>

      </div>
    </div>
  </div>
</template>

<style scoped>
.form-label { @apply block text-sm font-bold text-slate-600 dark:text-slate-300 uppercase tracking-wide mb-2; }
.input-field { @apply w-full p-3.5 bg-white dark:bg-slate-800 border border-slate-300 dark:border-slate-600 rounded-xl text-base outline-none focus:ring-2 focus:ring-yellow-500 dark:text-white transition shadow-sm placeholder:text-slate-400 text-slate-700; }
@keyframes fadeInUp { from { opacity: 0; transform: scale(0.95) translateY(10px); } to { opacity: 1; transform: scale(1) translateY(0); } }
.animate-fade-in-up { animation: fadeInUp 0.3s ease-out forwards; }

.custom-scrollbar::-webkit-scrollbar { width: 6px; }
.custom-scrollbar::-webkit-scrollbar-track { background: transparent; }
.custom-scrollbar::-webkit-scrollbar-thumb { background-color: #cbd5e1; border-radius: 10px; }
.custom-scrollbar::-webkit-scrollbar-thumb:hover { background-color: #94a3b8; }

input[type=number]::-webkit-inner-spin-button, 
input[type=number]::-webkit-outer-spin-button { 
  -webkit-appearance: none; 
  margin: 0; 
}
</style>