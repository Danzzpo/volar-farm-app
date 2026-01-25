<script setup>
import { ref, onMounted, watch } from 'vue'
import { 
  Bird, Search, ChevronLeft, ChevronRight, Plus, Calendar, Pencil, Trash, X, Save, FileText, Filter
} from 'lucide-vue-next'

const animals = ref([])
const loading = ref(false)
const userId = localStorage.getItem('user_id')

// --- STATE PAGINATION, SEARCH, & FILTER ---
const currentPage = ref(1)
const totalPages = ref(1)
const totalData = ref(0)
const searchQuery = ref('')
const filterStatus = ref('') // <--- State Filter Status Baru
const limit = 9 

// --- STATE MODAL ---
const isModalOpen = ref(false)
const isSubmitting = ref(false)
const editingId = ref(null)
const form = ref({
  ring_number: '', visual: '', gender: 'U', hatch_date: '',
  sire_other: '', dam_other: '', status: 'AVAILABLE', notes: ''
})

// --- 1. FETCH DATA (READ) ---
const fetchAnimals = async () => {
  loading.value = true
  try {
    // Kirim parameter status ke backend
    const res = await fetch(
      `http://localhost:8080/api/animals?user_id=${userId}&page=${currentPage.value}&limit=${limit}&search=${searchQuery.value}&status=${filterStatus.value}`
    )
    const json = await res.json()
    
    animals.value = json.data || []
    totalPages.value = json.pages || 1
    totalData.value = json.total || 0
    
  } catch (err) {
    console.error(err)
  } finally {
    loading.value = false
  }
}

// Watcher: Reset halaman ke 1 jika filter berubah, lalu ambil data
watch([searchQuery, filterStatus], () => {
  currentPage.value = 1
  fetchAnimals()
})

// Watcher: Jika ganti halaman, ambil data tanpa reset halaman
watch(currentPage, () => {
  fetchAnimals()
})

// --- 2. FITUR CARI INDUKAN (KLIK PARENT) ---
const cariIndukan = (ringNumber) => {
  if (ringNumber) {
    searchQuery.value = ringNumber
    filterStatus.value = '' // Reset status agar pencarian global
  }
}

// --- 3. SAVE DATA (CREATE / UPDATE) ---
const saveAnimal = async () => {
  if (!form.value.visual) return alert("Warna/Visual wajib diisi!")
  
  isSubmitting.value = true
  try {
    const payload = { ...form.value, user_id: parseInt(userId) }
    
    let url = 'http://localhost:8080/api/animals'
    let method = 'POST'

    if (editingId.value) {
      url = `http://localhost:8080/api/animals/${editingId.value}`
      method = 'PUT'
    }

    const res = await fetch(url, {
      method: method,
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload)
    })

    const json = await res.json()
    
    if (res.ok) {
      alert(editingId.value ? "✅ Data berhasil diupdate!" : "✅ Data berhasil disimpan!")
      closeModal()
      fetchAnimals()
    } else {
      alert("❌ Gagal: " + (json.error || "Terjadi kesalahan"))
    }
  } catch (err) {
    alert("Error sistem: " + err.message)
  } finally {
    isSubmitting.value = false
  }
}

// --- 4. DELETE DATA ---
const deleteAnimal = async (id) => {
  if (!confirm("⚠️ Yakin ingin menghapus data ini?")) return

  try {
    const res = await fetch(`http://localhost:8080/api/animals/${id}`, { method: 'DELETE' })
    if (res.ok) {
      fetchAnimals()
    } else {
      alert("Gagal menghapus data")
    }
  } catch (err) { console.error(err) }
}

// --- HELPERS ---
const prevPage = () => { if (currentPage.value > 1) currentPage.value-- }
const nextPage = () => { if (currentPage.value < totalPages.value) currentPage.value++ }

const openModal = () => {
  editingId.value = null
  form.value = {
    ring_number: '', visual: '', gender: 'U', hatch_date: '',
    sire_other: '', dam_other: '', status: 'AVAILABLE', notes: ''
  }
  isModalOpen.value = true
}

const openEditModal = (bird) => {
  editingId.value = bird.id
  form.value = { ...bird }
  if (form.value.hatch_date) {
    form.value.hatch_date = new Date(form.value.hatch_date).toISOString().split('T')[0]
  }
  isModalOpen.value = true
}

const closeModal = () => isModalOpen.value = false

const formatDate = (d) => {
  if (!d || d.startsWith('0001')) return '-'
  return new Date(d).toLocaleDateString('id-ID', { day: 'numeric', month: 'short', year: 'numeric' })
}

onMounted(() => fetchAnimals())
</script>

<template>
  <div class="space-y-6 animate-fade-in p-2 pb-20">
    
    <div class="flex flex-col md:flex-row justify-between items-center gap-4">
       <div>
          <h2 class="text-2xl font-bold text-slate-800 dark:text-white flex items-center gap-2">
             <Bird class="text-blue-600"/> Stok Burung
          </h2>
          <p class="text-slate-500 text-sm">Total: {{ totalData }} Ekor</p>
       </div>

       <div class="flex flex-col sm:flex-row gap-3 w-full md:w-auto">
          
          <div class="relative flex-1 md:w-56">
             <Search class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400"/>
             <input 
               v-model.lazy="searchQuery" 
               type="text" 
               placeholder="Cari Ring / Jenis..." 
               class="w-full pl-10 pr-8 py-2 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl text-sm outline-none focus:ring-2 focus:ring-blue-500 transition"
             >
             <button v-if="searchQuery" @click="searchQuery = ''" class="absolute right-3 top-1/2 -translate-y-1/2 text-slate-400 hover:text-slate-600">
                <X class="w-3 h-3"/>
             </button>
          </div>

          <div class="relative w-full sm:w-40">
             <Filter class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400"/>
             <select 
               v-model="filterStatus" 
               class="w-full pl-10 pr-4 py-2 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl text-sm outline-none focus:ring-2 focus:ring-blue-500 appearance-none cursor-pointer"
             >
               <option value="">Semua Status</option>
               <option value="AVAILABLE">Tersedia</option>
               <option value="SOLD">Terjual</option>
               <option value="Keep">Simpanan</option>
               <option value="Dead">Mati</option>
             </select>
             <div class="absolute right-3 top-1/2 -translate-y-1/2 pointer-events-none text-slate-400">
               <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path></svg>
             </div>
          </div>
          
          <button @click="openModal" class="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-xl flex items-center justify-center gap-2 font-bold text-sm shadow-lg shadow-blue-500/30 transition">
             <Plus class="w-4 h-4"/> <span class="inline">Tambah</span>
          </button>
       </div>
    </div>

    <div v-if="loading" class="text-center py-20 text-slate-400">Memuat data...</div>
    <div v-else-if="animals.length === 0" class="text-center py-20 text-slate-400 border-2 border-dashed border-slate-200 rounded-2xl">Data tidak ditemukan.</div>

    <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div 
        v-for="bird in animals" 
        :key="bird.id"
        class="bg-white dark:bg-slate-800 rounded-2xl border border-slate-100 dark:border-slate-700 shadow-sm hover:shadow-lg hover:-translate-y-1 transition-all duration-300 flex flex-col"
      >
        <div class="p-5 flex justify-between items-start relative">
           <div class="absolute top-4 right-4">
             <span class="px-2.5 py-1 rounded-md text-[10px] font-bold uppercase tracking-wider border"
               :class="{
                 'bg-green-50 text-green-700 border-green-200': bird.status === 'AVAILABLE',
                 'bg-red-50 text-red-700 border-red-200': bird.status === 'SOLD',
                 'bg-yellow-50 text-yellow-700 border-yellow-200': bird.status === 'Keep',
                 'bg-slate-100 text-slate-600 border-slate-200': bird.status === 'Dead'
               }">
               {{ bird.status }}
             </span>
           </div>
           <div class="pr-16"> 
             <span class="inline-block px-2 py-0.5 bg-slate-100 dark:bg-slate-700 text-slate-500 dark:text-slate-300 text-[10px] font-mono font-bold rounded mb-2">
               #{{ bird.ring_number || 'NO-RING' }}
             </span>
             <h3 class="text-xl font-bold text-slate-800 dark:text-white leading-tight truncate" :title="bird.visual">
               {{ bird.visual }}
             </h3>
             <p class="text-xs text-slate-400 mt-1">{{ bird.species }}</p>
           </div>
        </div>

        <div class="px-5 pb-5 flex-1">
          <div class="flex items-start gap-4">
            <div class="flex-shrink-0">
              <div class="w-12 h-12 rounded-2xl flex items-center justify-center text-white font-bold shadow-md text-lg"
                :class="{
                  'bg-gradient-to-br from-blue-500 to-blue-600': bird.gender === 'Male' || bird.gender === 'M',
                  'bg-gradient-to-br from-pink-500 to-pink-600': bird.gender === 'Female' || bird.gender === 'F',
                  'bg-slate-400': !['Male','Female','M','F'].includes(bird.gender)
                }">
                {{ (bird.gender === 'Male' || bird.gender === 'M') ? '♂' : (bird.gender === 'Female' || bird.gender === 'F') ? '♀' : '?' }}
              </div>
            </div>

            <div class="flex-1 bg-slate-50 dark:bg-slate-900/50 rounded-xl p-3 border border-slate-100 dark:border-slate-700/50 space-y-2">
              <div class="flex items-center gap-2 overflow-hidden">
                <span class="w-5 h-5 rounded-full bg-blue-100 text-blue-600 flex items-center justify-center text-[10px] font-bold flex-shrink-0">S</span>
                <span v-if="bird.sire" @click.stop="cariIndukan(bird.sire.ring_number)" class="text-xs font-bold text-blue-600 hover:underline cursor-pointer truncate">
                   {{ bird.sire.ring_number }}
                </span>
                <span v-else class="text-xs font-medium text-slate-500 truncate">{{ bird.sire_other || '-' }}</span>
              </div>

              <div class="flex items-center gap-2 overflow-hidden">
                <span class="w-5 h-5 rounded-full bg-pink-100 text-pink-600 flex items-center justify-center text-[10px] font-bold flex-shrink-0">D</span>
                <span v-if="bird.dam" @click.stop="cariIndukan(bird.dam.ring_number)" class="text-xs font-bold text-pink-600 hover:underline cursor-pointer truncate">
                   {{ bird.dam.ring_number }}
                </span>
                <span v-else class="text-xs font-medium text-slate-500 truncate">{{ bird.dam_other || '-' }}</span>
              </div>
            </div>
          </div>
        </div>

        <div class="bg-slate-50 dark:bg-slate-900/30 px-5 py-3 border-t border-slate-100 dark:border-slate-700 flex justify-between items-center mt-auto rounded-b-2xl">
           <div class="flex items-center gap-1.5 text-xs text-slate-400 font-medium">
             <Calendar class="w-3.5 h-3.5" />
             <span>{{ formatDate(bird.created_at) }}</span>
           </div>
           
           <div class="flex gap-2">
             <button @click="openEditModal(bird)" class="p-1.5 text-slate-400 hover:text-blue-600 hover:bg-blue-50 rounded transition" title="Edit">
               <Pencil class="w-4 h-4"/>
             </button>
             <button @click="deleteAnimal(bird.id)" class="p-1.5 text-slate-400 hover:text-red-600 hover:bg-red-50 rounded transition" title="Hapus">
               <Trash class="w-4 h-4"/>
             </button>
           </div>
        </div>

      </div>
    </div>

    <div v-if="totalPages > 1" class="flex justify-center items-center gap-4 mt-8">
       <button @click="prevPage" :disabled="currentPage === 1" class="p-2 rounded-lg border bg-white hover:bg-slate-50 disabled:opacity-50 disabled:cursor-not-allowed transition">
          <ChevronLeft class="w-5 h-5"/>
       </button>
       <span class="text-sm font-bold text-slate-600">Halaman {{ currentPage }} / {{ totalPages }}</span>
       <button @click="nextPage" :disabled="currentPage === totalPages" class="p-2 rounded-lg border bg-white hover:bg-slate-50 disabled:opacity-50 disabled:cursor-not-allowed transition">
          <ChevronRight class="w-5 h-5"/>
       </button>
    </div>

    <div v-if="isModalOpen" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-slate-900/60 backdrop-blur-sm" @click.self="closeModal">
      <div class="bg-white dark:bg-slate-900 w-full max-w-lg rounded-2xl shadow-2xl overflow-hidden animate-fade-in-up">
        
        <div class="px-6 py-4 border-b border-slate-100 dark:border-slate-800 flex justify-between items-center bg-slate-50 dark:bg-slate-900">
          <h3 class="font-bold text-lg text-slate-800 dark:text-white flex items-center gap-2">
            <FileText class="w-5 h-5 text-blue-500" /> 
            {{ editingId ? 'Edit Data Burung' : 'Tambah Data Baru' }}
          </h3>
          <button @click="closeModal" class="p-1 rounded-full hover:bg-slate-200 dark:hover:bg-slate-800 text-slate-400 transition"><X class="w-5 h-5"/></button>
        </div>

        <div class="p-6 space-y-5 max-h-[75vh] overflow-y-auto">
          <div class="grid grid-cols-3 gap-4">
             <div class="col-span-1">
                <label class="block text-xs font-bold text-slate-500 uppercase mb-1.5">No. Ring</label>
                <input v-model="form.ring_number" type="text" class="w-full p-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl text-sm font-mono" placeholder="001">
             </div>
             <div class="col-span-2">
                <label class="block text-xs font-bold text-slate-500 uppercase mb-1.5">Visual / Warna <span class="text-red-500">*</span></label>
                <input v-model="form.visual" type="text" class="w-full p-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl text-sm" placeholder="Contoh: Biola Green" required>
             </div>
          </div>

          <div>
             <label class="block text-xs font-bold text-slate-500 uppercase mb-1.5">Jenis Kelamin</label>
             <div class="flex gap-3">
               <button type="button" @click="form.gender = 'Male'" class="flex-1 py-2.5 rounded-xl border font-bold text-sm transition" :class="form.gender === 'Male' ? 'bg-blue-600 text-white border-blue-600' : 'bg-white dark:bg-slate-800 border-slate-200 text-slate-500'">♂ Jantan</button>
               <button type="button" @click="form.gender = 'Female'" class="flex-1 py-2.5 rounded-xl border font-bold text-sm transition" :class="form.gender === 'Female' ? 'bg-pink-600 text-white border-pink-600' : 'bg-white dark:bg-slate-800 border-slate-200 text-slate-500'">♀ Betina</button>
               <button type="button" @click="form.gender = 'Unknown'" class="flex-1 py-2.5 rounded-xl border font-bold text-sm transition" :class="form.gender === 'Unknown' ? 'bg-slate-600 text-white border-slate-600' : 'bg-white dark:bg-slate-800 border-slate-200 text-slate-500'">? Unsex</button>
             </div>
          </div>

          <div class="grid grid-cols-2 gap-4">
            <div><label class="block text-xs font-bold text-slate-500 uppercase mb-1.5">Tgl Menetas</label><input v-model="form.hatch_date" type="date" class="w-full p-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl text-sm"></div>
            <div>
              <label class="block text-xs font-bold text-slate-500 uppercase mb-1.5">Status</label>
              <select v-model="form.status" class="w-full p-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl text-sm">
                <option value="AVAILABLE">Available</option>
                <option value="SOLD">Sold</option>
                <option value="Dead">Dead</option>
                <option value="Keep">Keep</option>
              </select>
            </div>
          </div>
          
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
             <div><label class="block text-xs font-bold text-slate-500 uppercase mb-1.5">Bapak (Manual)</label><input v-model="form.sire_other" type="text" class="w-full p-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl text-sm" placeholder="Isi jika tidak ada di DB"></div>
             <div><label class="block text-xs font-bold text-slate-500 uppercase mb-1.5">Ibu (Manual)</label><input v-model="form.dam_other" type="text" class="w-full p-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl text-sm" placeholder="Isi jika tidak ada di DB"></div>
          </div>
        </div>

        <div class="px-6 py-4 bg-slate-50 dark:bg-slate-900 border-t border-slate-100 dark:border-slate-800 flex gap-3">
          <button @click="closeModal" class="flex-1 py-3 rounded-xl border border-slate-300 font-bold text-sm text-slate-500 hover:bg-slate-100 transition">Batal</button>
          <button @click="saveAnimal" :disabled="isSubmitting" class="flex-1 py-3 rounded-xl bg-blue-600 text-white font-bold text-sm hover:bg-blue-700 transition flex items-center justify-center gap-2 shadow-lg shadow-blue-500/30">
            <Save v-if="!isSubmitting" class="w-5 h-5" />
            <span>{{ isSubmitting ? 'Menyimpan...' : (editingId ? 'Update Data' : 'Simpan Data') }}</span>
          </button>
        </div>

      </div>
    </div>

  </div>
</template>

<style scoped>
@keyframes fadeInUp { from { opacity: 0; transform: scale(0.95) translateY(10px); } to { opacity: 1; transform: scale(1) translateY(0); } }
.animate-fade-in-up { animation: fadeInUp 0.3s ease-out forwards; }
</style>