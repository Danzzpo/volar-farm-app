<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
// Import Icon Modern
import { 
  HeartHandshake, Search, Plus, X, Save, 
  Bird, Calendar, AlertCircle, Trash2, Power, History
} from 'lucide-vue-next'

const router = useRouter()
const pairs = ref([])
const animals = ref([]) 
const showModal = ref(false)
const loading = ref(false)
const searchQuery = ref('')
const isSubmitting = ref(false)

const userId = localStorage.getItem('user_id')

const form = ref({
  male_id: "",
  female_id: "",
  cage: '',
  pair_date: new Date().toISOString().split('T')[0],
  notes: ''
})

// --- COMPUTED DATA ---

// 1. Filter Burung untuk Dropdown (Hanya yg Available & Sesuai Gender)
const males = computed(() => animals.value.filter(a => a.gender === 'M' && (a.status === 'Available' || a.status === 'Keep')))
const females = computed(() => animals.value.filter(a => a.gender === 'F' && (a.status === 'Available' || a.status === 'Keep')))

// 2. Filter List Pasangan (Search)
const filteredPairs = computed(() => {
  if (!searchQuery.value) return pairs.value
  const lower = searchQuery.value.toLowerCase()
  return pairs.value.filter(p => 
    p.cage.toLowerCase().includes(lower) ||
    (p.male?.visual && p.male.visual.toLowerCase().includes(lower)) ||
    (p.female?.visual && p.female.visual.toLowerCase().includes(lower))
  )
})

// --- FETCH DATA ---
onMounted(() => {
  if (!userId) { router.push('/login'); return }
  fetchAnimals()
  fetchPairs()
})

const fetchAnimals = async () => {
  try {
    const res = await fetch(`http://localhost:8080/api/animals?user_id=${userId}`)
    const json = await res.json()
    animals.value = json.data || []
  } catch (err) { console.error(err) }
}

const fetchPairs = async () => {
  loading.value = true
  try {
    const res = await fetch(`http://localhost:8080/api/pairs?user_id=${userId}`)
    const json = await res.json()
    pairs.value = json.data || []
  } catch (err) { console.error(err) } finally { loading.value = false }
}

// --- ACTIONS ---

const openModal = () => {
  // Reset form
  form.value = {
    male_id: "", female_id: "", cage: '',
    pair_date: new Date().toISOString().split('T')[0], notes: ''
  }
  showModal.value = true
}

const submitPair = async () => {
  if(!form.value.male_id || !form.value.female_id || !form.value.cage) {
    alert("Jantan, Betina, dan Kandang wajib diisi!")
    return
  }
  
  isSubmitting.value = true
  const payload = {
    user_id: parseInt(userId),
    male_id: parseInt(form.value.male_id),
    female_id: parseInt(form.value.female_id),
    cage: form.value.cage,
    notes: form.value.notes,
    pair_date: new Date(form.value.pair_date).toISOString()
  }

  try {
    const res = await fetch('http://localhost:8080/api/pairs', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload)
    })

    if (res.ok) {
      alert("✅ Pasangan Berhasil Dibuat!")
      showModal.value = false
      fetchPairs() 
    } else {
      const data = await res.json()
      alert("Gagal: " + (data.error || "Terjadi kesalahan"))
    }
  } catch (err) { console.error(err) } finally { isSubmitting.value = false }
}

const updateStatus = async (pairId, newStatus) => {
  const msg = newStatus === 'REST' ? 'Istirahatkan pasangan ini?' : 'Cabut/Bubarkan pasangan ini?'
  if(!confirm(msg)) return
  
  try {
    await fetch(`http://localhost:8080/api/pairs/${pairId}`, {
      method: 'PATCH',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ status: newStatus })
    })
    fetchPairs()
  } catch(err) { console.error(err) }
}

const formatDate = (d) => new Date(d).toLocaleDateString('id-ID', {day: 'numeric', month: 'short', year: '2-digit'})
</script>

<template>
  <div class="space-y-6 p-2">
    
    <div class="flex flex-col md:flex-row md:items-end justify-between gap-6">
      <div>
        <h2 class="text-3xl font-extrabold text-slate-800 dark:text-white flex items-center gap-3">
          <HeartHandshake class="w-8 h-8 text-pink-500" /> Pairs (Jodohan)
        </h2>
        <p class="text-slate-500 mt-1 text-sm md:text-base">Kelola pasangan aktif dalam kandang ternak.</p>
      </div>

      <div class="flex flex-col sm:flex-row gap-3 w-full md:w-auto">
        <div class="relative flex-1">
          <Search class="absolute left-3 top-1/2 -translate-y-1/2 text-slate-400 w-4 h-4" />
          <input 
            v-model="searchQuery" 
            type="text" 
            placeholder="Cari Kandang / Visual..." 
            class="w-full pl-10 pr-4 py-2.5 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl text-sm focus:ring-2 focus:ring-pink-500 outline-none dark:text-white shadow-sm transition"
          >
        </div>
        <button @click="openModal" class="bg-gradient-to-r from-pink-500 to-rose-500 hover:from-pink-600 hover:to-rose-600 text-white px-5 py-2.5 rounded-xl font-bold flex items-center justify-center gap-2 shadow-lg shadow-pink-500/20 active:scale-95 transition">
          <Plus class="w-5 h-5" /> <span>Pasangan Baru</span>
        </button>
      </div>
    </div>

    <div v-if="loading" class="text-center py-12 text-slate-400">Memuat data pasangan...</div>

    <div v-else-if="filteredPairs.length === 0" class="text-center py-20 border-2 border-dashed border-slate-300 dark:border-slate-700 rounded-3xl bg-slate-50 dark:bg-slate-800/50">
       <HeartHandshake class="w-16 h-16 text-slate-300 mx-auto mb-4" />
       <h3 class="text-lg font-bold text-slate-600 dark:text-slate-300">Belum ada pasangan</h3>
       <p class="text-slate-400 text-sm mb-4">Mulai menjodohkan burung untuk produksi.</p>
       <button @click="openModal" class="text-pink-500 font-bold hover:underline">Buat Jodohan Sekarang</button>
    </div>

    <div v-else class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-6">
      
      <div 
        v-for="pair in filteredPairs" 
        :key="pair.id" 
        class="bg-white dark:bg-slate-800 rounded-2xl shadow-sm hover:shadow-xl transition duration-300 border border-slate-100 dark:border-slate-700 group overflow-hidden flex flex-col"
      >
        <div class="px-5 py-3 border-b border-slate-100 dark:border-slate-700 flex justify-between items-center bg-slate-50/50 dark:bg-slate-900/30">
          <div class="flex items-center gap-2">
            <div class="w-8 h-8 rounded-lg bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-600 flex items-center justify-center shadow-sm">
              <span class="text-lg">🏠</span>
            </div>
            <span class="font-bold text-slate-700 dark:text-slate-200 text-lg">{{ pair.cage }}</span>
          </div>
          <span class="text-[10px] font-bold px-2.5 py-1 rounded-full border"
            :class="pair.status === 'Active' 
              ? 'bg-green-100 text-green-700 border-green-200' 
              : pair.status === 'Rest' ? 'bg-orange-100 text-orange-700 border-orange-200' : 'bg-slate-100 text-slate-500 border-slate-200'">
            {{ pair.status.toUpperCase() }}
          </span>
        </div>

        <div class="p-5 flex-1 relative">
          <div class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-full h-px border-t-2 border-dashed border-slate-200 dark:border-slate-700 -z-0"></div>
          <div class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 bg-white dark:bg-slate-800 p-1.5 rounded-full border border-slate-100 dark:border-slate-700 z-10 shadow-sm">
             <HeartHandshake class="w-4 h-4 text-pink-400" />
          </div>

          <div class="flex justify-between items-start gap-4 z-10 relative">
            
            <div class="flex-1 flex flex-col items-center text-center">
              <div class="w-12 h-12 rounded-full bg-blue-100 dark:bg-blue-900/30 text-blue-600 border-2 border-white dark:border-slate-700 shadow-md flex items-center justify-center mb-2">
                <Bird class="w-6 h-6" />
              </div>
              <h4 class="font-bold text-slate-800 dark:text-white text-sm leading-tight line-clamp-2" :title="pair.male?.visual">{{ pair.male?.visual || 'Unknown' }}</h4>
              <span class="text-[10px] text-slate-500 bg-slate-100 dark:bg-slate-700 px-1.5 py-0.5 rounded mt-1 font-mono">{{ pair.male?.ring_number }}</span>
            </div>

            <div class="flex-1 flex flex-col items-center text-center">
              <div class="w-12 h-12 rounded-full bg-pink-100 dark:bg-pink-900/30 text-pink-600 border-2 border-white dark:border-slate-700 shadow-md flex items-center justify-center mb-2">
                <Bird class="w-6 h-6" />
              </div>
              <h4 class="font-bold text-slate-800 dark:text-white text-sm leading-tight line-clamp-2" :title="pair.female?.visual">{{ pair.female?.visual || 'Unknown' }}</h4>
              <span class="text-[10px] text-slate-500 bg-slate-100 dark:bg-slate-700 px-1.5 py-0.5 rounded mt-1 font-mono">{{ pair.female?.ring_number }}</span>
            </div>

          </div>
        </div>

        <div class="bg-slate-50 dark:bg-slate-900/50 p-3 border-t border-slate-100 dark:border-slate-700">
           <div class="flex justify-between items-center text-xs text-slate-400 mb-3">
              <div class="flex items-center gap-1">
                <Calendar class="w-3 h-3" /> {{ formatDate(pair.pair_date || pair.created_at) }}
              </div>
              <div v-if="pair.notes" class="flex items-center gap-1 cursor-help" :title="pair.notes">
                <AlertCircle class="w-3 h-3" /> Info
              </div>
           </div>

           <div v-if="pair.status !== 'History'" class="grid grid-cols-2 gap-2">
             <button v-if="pair.status === 'Active'" @click="updateStatus(pair.id, 'Rest')" 
                class="flex items-center justify-center gap-1 py-1.5 rounded-lg border border-orange-200 text-orange-600 bg-orange-50 hover:bg-orange-100 text-xs font-bold transition">
                <Power class="w-3 h-3" /> Istirahat
             </button>
             <button v-if="pair.status === 'Rest'" @click="updateStatus(pair.id, 'Active')" 
                class="flex items-center justify-center gap-1 py-1.5 rounded-lg border border-green-200 text-green-600 bg-green-50 hover:bg-green-100 text-xs font-bold transition">
                <Power class="w-3 h-3" /> Aktifkan
             </button>
             
             <button @click="updateStatus(pair.id, 'History')" 
                class="flex items-center justify-center gap-1 py-1.5 rounded-lg border border-slate-200 text-slate-500 bg-white hover:bg-slate-100 hover:text-red-600 text-xs font-bold transition">
                <History class="w-3 h-3" /> Cabut/Bubar
             </button>
           </div>
           <div v-else class="text-center py-1 text-xs font-bold text-slate-400">
             Pasangan Non-Aktif (History)
           </div>
        </div>

      </div>
    </div>

    <div v-if="showModal" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-slate-900/60 backdrop-blur-sm" @click.self="showModal = false">
      <div class="bg-white dark:bg-slate-900 w-full max-w-2xl rounded-2xl shadow-2xl overflow-hidden animate-fade-in-up">
        
        <div class="px-6 py-4 border-b border-slate-100 dark:border-slate-800 flex justify-between items-center bg-slate-50 dark:bg-slate-900">
          <div>
            <h3 class="font-bold text-lg text-slate-800 dark:text-white flex items-center gap-2">
              <HeartHandshake class="w-5 h-5 text-pink-500"/> Jodohkan Pasangan Baru
            </h3>
          </div>
          <button @click="showModal = false" class="p-1 rounded-full hover:bg-slate-200 dark:hover:bg-slate-800 text-slate-500 transition"><X class="w-5 h-5" /></button>
        </div>

        <div class="p-6 space-y-5 max-h-[75vh] overflow-y-auto">
          
          <div class="grid grid-cols-1 md:grid-cols-2 gap-5">
            <div class="space-y-1">
              <label class="text-xs font-bold text-blue-600 uppercase tracking-wider">Pejantan (Sire)</label>
              <div class="relative">
                <select v-model="form.male_id" class="w-full p-3 bg-blue-50/50 dark:bg-slate-800 border border-blue-100 dark:border-slate-700 rounded-xl text-sm outline-none focus:ring-2 focus:ring-blue-500 dark:text-white appearance-none">
                  <option value="" disabled>-- Pilih Jantan Available --</option>
                  <option v-for="b in males" :key="b.id" :value="b.id">
                    {{ b.ring_number }} - {{ b.visual }}
                  </option>
                </select>
                <div class="absolute right-3 top-3 pointer-events-none text-blue-400">▼</div>
              </div>
            </div>

            <div class="space-y-1">
              <label class="text-xs font-bold text-pink-600 uppercase tracking-wider">Betina (Dam)</label>
              <div class="relative">
                <select v-model="form.female_id" class="w-full p-3 bg-pink-50/50 dark:bg-slate-800 border border-pink-100 dark:border-slate-700 rounded-xl text-sm outline-none focus:ring-2 focus:ring-pink-500 dark:text-white appearance-none">
                  <option value="" disabled>-- Pilih Betina Available --</option>
                  <option v-for="b in females" :key="b.id" :value="b.id">
                    {{ b.ring_number }} - {{ b.visual }}
                  </option>
                </select>
                <div class="absolute right-3 top-3 pointer-events-none text-pink-400">▼</div>
              </div>
            </div>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-5">
            <div class="space-y-1">
              <label class="text-xs font-bold text-slate-500 uppercase tracking-wider">Nomor / Nama Kandang</label>
              <input v-model="form.cage" type="text" placeholder="Contoh: Glodok A1" class="w-full p-3 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl text-sm outline-none focus:ring-2 focus:ring-slate-400 dark:text-white">
            </div>
            <div class="space-y-1">
              <label class="text-xs font-bold text-slate-500 uppercase tracking-wider">Tanggal Masuk</label>
              <input v-model="form.pair_date" type="date" class="w-full p-3 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl text-sm outline-none focus:ring-2 focus:ring-slate-400 dark:text-white">
            </div>
          </div>

          <div class="space-y-1">
            <label class="text-xs font-bold text-slate-500 uppercase tracking-wider">Catatan Tambahan</label>
            <textarea v-model="form.notes" rows="3" placeholder="Kondisi burung, program ternak, dll..." class="w-full p-3 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl text-sm outline-none focus:ring-2 focus:ring-slate-400 dark:text-white resize-none"></textarea>
          </div>

        </div>

        <div class="px-6 py-4 bg-slate-50 dark:bg-slate-900 border-t border-slate-100 dark:border-slate-800 flex gap-3">
          <button @click="showModal = false" class="flex-1 py-3 rounded-xl border border-slate-300 font-bold text-sm text-slate-500 hover:bg-slate-100 transition">Batal</button>
          <button @click="submitPair" :disabled="isSubmitting" class="flex-1 py-3 rounded-xl bg-gradient-to-r from-pink-600 to-rose-600 text-white font-bold text-sm hover:from-pink-700 hover:to-rose-700 transition shadow-lg shadow-pink-500/30 flex items-center justify-center gap-2">
            <Save v-if="!isSubmitting" class="w-5 h-5" />
            <span>{{ isSubmitting ? 'Menyimpan...' : 'Simpan Pasangan' }}</span>
          </button>
        </div>

      </div>
    </div>

  </div>
</template>

<style scoped>
@keyframes fadeInUp {
  from { opacity: 0; transform: scale(0.95) translateY(10px); }
  to { opacity: 1; transform: scale(1) translateY(0); }
}
.animate-fade-in-up {
  animation: fadeInUp 0.3s ease-out forwards;
}
</style>