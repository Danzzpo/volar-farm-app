<script setup>
import { ref, onMounted, computed } from 'vue'
import { 
  Search, Plus, Bird, Calendar, Pencil, Trash, X, Save, 
  FileText
} from 'lucide-vue-next'

const animals = ref([])
const searchQuery = ref('')
const loading = ref(true)

// State Modal & Form
const isModalOpen = ref(false)
const isSubmitting = ref(false)
const editingId = ref(null) // Melacak ID yang sedang diedit

const userId = localStorage.getItem('user_id')

const form = ref({
  ring_number: '', visual: '', gender: 'U', hatch_date: '',
  sire_other: '', dam_other: '', status: 'Available', notes: ''
})

// --- 1. FETCH DATA (READ) ---
const fetchAnimals = async () => {
  loading.value = true
  try {
    const res = await fetch(`http://localhost:8080/api/animals?user_id=${userId}`)
    const json = await res.json()
    animals.value = json.data || []
  } catch (err) { console.error(err) } finally { loading.value = false }
}

// --- 2. SAVE DATA (CREATE / UPDATE) ---
const saveAnimal = async () => {
  if (!form.value.visual) return alert("Warna/Visual wajib diisi!")
  
  isSubmitting.value = true
  try {
    const payload = { ...form.value, user_id: parseInt(userId) }
    
    // Tentukan URL & Method (POST untuk Baru, PUT untuk Edit)
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

// --- 3. DELETE DATA ---
const deleteAnimal = async (id) => {
  if (!confirm("⚠️ Yakin ingin menghapus data ini? Data yang dihapus tidak bisa dikembalikan.")) return

  try {
    const res = await fetch(`http://localhost:8080/api/animals/${id}`, { method: 'DELETE' })
    if (res.ok) {
      fetchAnimals() // Refresh list
    } else {
      alert("Gagal menghapus data")
    }
  } catch (err) { console.error(err) }
}

// --- HELPERS (MODAL & FORMATTING) ---
const openModal = () => {
  // Mode Tambah Baru: Reset form & editingId
  editingId.value = null
  form.value = {
    ring_number: '', visual: '', gender: 'U', hatch_date: '',
    sire_other: '', dam_other: '', status: 'Available', notes: ''
  }
  isModalOpen.value = true
}

const openEditModal = (bird) => {
  // Mode Edit: Isi form dengan data yang dipilih
  editingId.value = bird.id
  form.value = { ...bird }
  
  // Format tanggal agar bisa masuk ke input type="date"
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

const filteredAnimals = computed(() => {
  if (!searchQuery.value) return animals.value
  const lower = searchQuery.value.toLowerCase()
  return animals.value.filter(a => 
    (a.visual && a.visual.toLowerCase().includes(lower)) || 
    (a.ring_number && a.ring_number.toLowerCase().includes(lower))
  )
})

onMounted(() => fetchAnimals())
</script>

<template>
  <div class="space-y-6 p-2">
    
    <div class="flex flex-col md:flex-row md:items-end justify-between gap-6">
      <div>
        <h2 class="text-3xl font-extrabold text-slate-800 dark:text-white flex items-center gap-3">
          <Bird class="w-8 h-8 text-blue-500" /> Stok Burung
        </h2>
        <p class="text-slate-500 mt-1 text-sm md:text-base">Kelola database burung, silsilah, dan status penjualan.</p>
      </div>
      
      <div class="flex flex-col sm:flex-row gap-3 w-full md:w-auto">
        <div class="relative flex-1">
          <Search class="absolute left-3 top-1/2 -translate-y-1/2 text-slate-400 w-4 h-4" />
          <input 
            v-model="searchQuery" type="text" placeholder="Cari Ring / Visual..." 
            class="w-full pl-10 pr-4 py-2.5 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl text-sm outline-none focus:ring-2 focus:ring-blue-500 dark:text-white shadow-sm transition"
          >
        </div>
        <button @click="openModal" class="bg-blue-600 hover:bg-blue-700 text-white px-5 py-2.5 rounded-xl font-bold flex items-center justify-center gap-2 shadow-lg shadow-blue-500/20 active:scale-95 transition">
          <Plus class="w-5 h-5" /> <span>Data Baru</span>
        </button>
      </div>
    </div>

    <div v-if="loading" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-3 gap-6">
       <div v-for="i in 6" :key="i" class="bg-white dark:bg-slate-800 h-56 rounded-2xl animate-pulse"></div>
    </div>
    
    <div v-else-if="filteredAnimals.length === 0" class="text-center py-20 bg-slate-50 dark:bg-slate-800/50 rounded-3xl border-2 border-dashed border-slate-200 dark:border-slate-700">
       <div class="w-16 h-16 bg-white dark:bg-slate-800 rounded-full flex items-center justify-center mx-auto mb-4 shadow-sm">
         <Bird class="w-8 h-8 text-slate-300" />
       </div>
       <h3 class="text-lg font-bold text-slate-600 dark:text-slate-300">Data Tidak Ditemukan</h3>
       <button @click="openModal" class="text-blue-500 font-bold hover:underline text-sm mt-2">Input Data Baru</button>
    </div>

    <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-3 gap-6">
      <div 
        v-for="bird in filteredAnimals" 
        :key="bird.id"
        class="bg-white dark:bg-slate-800 rounded-2xl border border-slate-100 dark:border-slate-700 shadow-sm hover:shadow-xl hover:-translate-y-1 transition-all duration-300 group overflow-hidden flex flex-col"
      >
        <div class="p-5 flex justify-between items-start relative">
           <div class="absolute top-4 right-4">
             <span class="px-2.5 py-1 rounded-md text-[10px] font-bold uppercase tracking-wider border"
               :class="{
                 'bg-green-50 text-green-700 border-green-200': bird.status === 'Available',
                 'bg-red-50 text-red-700 border-red-200': bird.status === 'Sold',
                 'bg-slate-100 text-slate-600 border-slate-200': bird.status === 'Dead' || bird.status === 'Keep'
               }">
               {{ bird.status }}
             </span>
           </div>
           <div class="pr-16"> 
             <span class="inline-block px-2 py-0.5 bg-slate-100 dark:bg-slate-700 text-slate-500 dark:text-slate-300 text-[10px] font-mono font-bold rounded mb-2">
               #{{ bird.ring_number || 'NO-RING' }}
             </span>
             <h3 class="text-xl font-bold text-slate-800 dark:text-white leading-tight break-words" :title="bird.visual">
               {{ bird.visual }}
             </h3>
           </div>
        </div>

        <div class="px-5 pb-5 flex-1">
          <div class="flex items-start gap-4">
            <div class="flex-shrink-0">
              <div class="w-12 h-12 rounded-2xl flex items-center justify-center text-white font-bold shadow-md text-lg"
                :class="{
                  'bg-gradient-to-br from-blue-500 to-blue-600': bird.gender === 'M',
                  'bg-gradient-to-br from-pink-500 to-pink-600': bird.gender === 'F',
                  'bg-slate-400': bird.gender === 'U'
                }">
                {{ bird.gender === 'M' ? '♂' : bird.gender === 'F' ? '♀' : '?' }}
              </div>
            </div>
            <div class="flex-1 bg-slate-50 dark:bg-slate-900/50 rounded-xl p-3 border border-slate-100 dark:border-slate-700/50 space-y-2">
              <div class="flex items-center gap-2 overflow-hidden">
                <span class="w-5 h-5 rounded-full bg-blue-100 text-blue-600 flex items-center justify-center text-[10px] font-bold flex-shrink-0">S</span>
                <span class="text-xs font-medium text-slate-600 dark:text-slate-300 truncate">{{ bird.sire_other || '-' }}</span>
              </div>
              <div class="flex items-center gap-2 overflow-hidden">
                <span class="w-5 h-5 rounded-full bg-pink-100 text-pink-600 flex items-center justify-center text-[10px] font-bold flex-shrink-0">D</span>
                <span class="text-xs font-medium text-slate-600 dark:text-slate-300 truncate">{{ bird.dam_other || '-' }}</span>
              </div>
            </div>
          </div>
        </div>

        <div class="bg-slate-50 dark:bg-slate-900/30 px-5 py-3 border-t border-slate-100 dark:border-slate-700 flex justify-between items-center mt-auto">
           <div class="flex items-center gap-1.5 text-xs text-slate-400 font-medium">
             <Calendar class="w-3.5 h-3.5" />
             <span>{{ formatDate(bird.hatch_date) }}</span>
           </div>
           
           <div class="flex gap-2">
             <button @click="openEditModal(bird)" class="p-1.5 text-slate-400 hover:text-blue-600 hover:bg-blue-50 rounded transition" title="Edit Data">
               <Pencil class="w-4 h-4"/>
             </button>
             <button @click="deleteAnimal(bird.id)" class="p-1.5 text-slate-400 hover:text-red-600 hover:bg-red-50 rounded transition" title="Hapus Data">
               <Trash class="w-4 h-4"/>
             </button>
           </div>
        </div>

      </div>
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
                <label class="form-label">No. Ring</label>
                <input v-model="form.ring_number" type="text" class="input-field font-mono" placeholder="001">
             </div>
             <div class="col-span-2">
                <label class="form-label">Visual / Warna <span class="text-red-500">*</span></label>
                <input v-model="form.visual" type="text" class="input-field" placeholder="Contoh: Biola Green" required>
             </div>
          </div>

          <div>
             <label class="form-label">Jenis Kelamin</label>
             <div class="flex gap-3">
               <button type="button" @click="form.gender = 'M'" class="gender-btn" :class="form.gender === 'M' ? 'bg-blue-600 text-white border-blue-600 shadow-md' : 'bg-white dark:bg-slate-800 border-slate-200 text-slate-500'">♂ Jantan</button>
               <button type="button" @click="form.gender = 'F'" class="gender-btn" :class="form.gender === 'F' ? 'bg-pink-600 text-white border-pink-600 shadow-md' : 'bg-white dark:bg-slate-800 border-slate-200 text-slate-500'">♀ Betina</button>
               <button type="button" @click="form.gender = 'U'" class="gender-btn" :class="form.gender === 'U' ? 'bg-slate-600 text-white border-slate-600 shadow-md' : 'bg-white dark:bg-slate-800 border-slate-200 text-slate-500'">? Unsex</button>
             </div>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
             <div><label class="form-label">Bapak (Sire)</label><input v-model="form.sire_other" type="text" class="input-field" placeholder="Nama / Ring Bapak"></div>
             <div><label class="form-label">Ibu (Dam)</label><input v-model="form.dam_other" type="text" class="input-field" placeholder="Nama / Ring Ibu"></div>
          </div>

          <div class="grid grid-cols-2 gap-4">
            <div><label class="form-label">Tgl Menetas</label><input v-model="form.hatch_date" type="date" class="input-field"></div>
            <div>
              <label class="form-label">Status</label>
              <select v-model="form.status" class="input-field">
                <option value="Available">Available</option>
                <option value="Sold">Sold</option>
                <option value="Dead">Dead</option>
                <option value="Keep">Keep</option>
              </select>
            </div>
          </div>
        </div>

        <div class="px-6 py-4 bg-slate-50 dark:bg-slate-900 border-t border-slate-100 dark:border-slate-800 flex gap-3">
          <button @click="closeModal" class="flex-1 py-3 rounded-xl border border-slate-300 font-bold text-sm text-slate-500 hover:bg-slate-100 transition">Batal</button>
          <button @click="saveAnimal" :disabled="isSubmitting" class="flex-1 py-3 rounded-xl bg-blue-600 text-white font-bold text-sm hover:bg-blue-700 transition shadow-lg shadow-blue-500/30 flex items-center justify-center gap-2">
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
.form-label { @apply block text-xs font-bold text-slate-500 uppercase tracking-wider mb-1.5; }
.input-field { @apply w-full p-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl text-sm outline-none focus:ring-2 focus:ring-blue-500 dark:text-white transition; }
.gender-btn { @apply flex-1 py-2.5 rounded-xl border font-bold text-sm transition shadow-sm hover:border-slate-300; }
</style>