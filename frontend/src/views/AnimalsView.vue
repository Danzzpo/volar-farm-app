<script setup>
import { ref, onMounted, computed, watch } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const animals = ref([])
const showForm = ref(false)
const loading = ref(false)

const userId = localStorage.getItem('user_id')

// --- STATE PENCARIAN & INPUT ---
const searchQuery = ref('') 
const sireQuery = ref('')   
const damQuery = ref('')    
const showSireList = ref(false) 
const showDamList = ref(false)  

// FORM DATA
const form = ref({
  ring_number: '', species: 'Lovebird', gender: 'U', visual: '',
  sire_id: null, sire_other: '', 
  dam_id: null, dam_other: '',  
  status: 'AVAILABLE'
})

// Reset pencarian saat form dibuka/tutup
watch(showForm, (val) => {
  if (!val) { sireQuery.value = ''; damQuery.value = '' }
})

// --- FILTER LOGIC ---
const filteredAnimals = computed(() => {
  const q = searchQuery.value.toLowerCase()
  return animals.value.filter(a => a.ring_number.toLowerCase().includes(q) || a.visual.toLowerCase().includes(q))
})

const filteredSires = computed(() => {
  if (!sireQuery.value) return []
  const q = sireQuery.value.toLowerCase()
  return animals.value.filter(a => a.gender === 'M' && (a.ring_number.toLowerCase().includes(q) || a.visual.toLowerCase().includes(q)))
})

const filteredDams = computed(() => {
  if (!damQuery.value) return []
  const q = damQuery.value.toLowerCase()
  return animals.value.filter(a => a.gender === 'F' && (a.ring_number.toLowerCase().includes(q) || a.visual.toLowerCase().includes(q)))
})

// --- METHODS ---
onMounted(() => {
  if (!userId) { router.push('/login') } else { fetchAnimals() }
})

const fetchAnimals = async () => {
  try {
    const res = await fetch(`http://localhost:8080/api/animals?user_id=${userId}`)
    const json = await res.json()
    animals.value = json.data || []
  } catch (err) { console.error(err) }
}

const onSireInput = () => {
  showSireList.value = true
  form.value.sire_other = sireQuery.value
  form.value.sire_id = null 
}

const selectSire = (bird) => {
  form.value.sire_id = bird.id
  sireQuery.value = `${bird.ring_number} - ${bird.visual}`
  form.value.sire_other = ''
  showSireList.value = false
}

const onDamInput = () => {
  showDamList.value = true
  form.value.dam_other = damQuery.value
  form.value.dam_id = null
}

const selectDam = (bird) => {
  form.value.dam_id = bird.id
  damQuery.value = `${bird.ring_number} - ${bird.visual}`
  form.value.dam_other = ''
  showDamList.value = false
}

const submitAnimal = async () => {
  loading.value = true
  if (!form.value.sire_id) form.value.sire_other = sireQuery.value
  if (!form.value.dam_id) form.value.dam_other = damQuery.value

  const payload = { user_id: parseInt(userId), ...form.value }

  try {
    const res = await fetch('http://localhost:8080/api/animals', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload)
    })

    if (res.ok) {
      alert("✅ Data Berhasil Disimpan!")
      showForm.value = false
      fetchAnimals()
      form.value = { ring_number: '', species: 'Lovebird', gender: 'U', visual: '', sire_id: null, sire_other: '', dam_id: null, dam_other: '', status: 'AVAILABLE' }
      sireQuery.value = ''
      damQuery.value = ''
    } else {
      alert("Gagal menyimpan!")
    }
  } catch (err) { console.error(err) } finally { loading.value = false }
}
</script>

<template>
  <div class="pb-20 max-w-6xl mx-auto">
    
    <div class="flex flex-col md:flex-row justify-between items-center mb-8 gap-4">
      <div>
        <h2 class="text-3xl font-bold text-slate-800 dark:text-white">Daftar Burung</h2>
        <p class="text-slate-500 dark:text-slate-400">Total: {{ animals.length }} Ekor</p>
      </div>
      <div class="flex gap-3 w-full md:w-auto">
        <input 
          v-model="searchQuery" 
          placeholder="🔍 Cari Ring / Visual..." 
          class="w-full md:w-64 pl-4 p-3 rounded-full border border-slate-300 dark:border-slate-600 bg-white dark:bg-slate-800 text-slate-900 dark:text-white focus:ring-2 focus:ring-green-500 shadow-sm transition-colors"
        >
        <button @click="showForm = !showForm" class="bg-blue-600 text-white px-6 py-3 rounded-full hover:bg-blue-700 font-bold shadow-lg flex items-center gap-2 whitespace-nowrap transition-transform hover:scale-105">
          <span>+</span> Input Baru
        </button>
      </div>
    </div>

    <div v-if="showForm" class="bg-white dark:bg-[#1e293b] p-8 rounded-2xl shadow-xl mb-10 border border-slate-200 dark:border-slate-700 animate-fade-in relative z-40 transition-colors duration-300">
      <h3 class="text-xl font-bold mb-6 text-slate-700 dark:text-slate-200 border-b dark:border-slate-700 pb-2">📝 Biodata Burung</h3>
      <form @submit.prevent="submitAnimal" class="grid grid-cols-1 md:grid-cols-2 gap-8">
        
        <div class="space-y-4">
          <h4 class="font-bold text-sm text-blue-600 dark:text-blue-400 uppercase border-b dark:border-slate-700 pb-1">1. Identitas</h4>
          
          <div>
            <label class="text-xs font-bold text-slate-500 dark:text-slate-400">Nomor Ring</label>
            <input v-model="form.ring_number" placeholder="Contoh: VF-001" class="w-full p-3 bg-slate-50 dark:bg-slate-900 border border-slate-200 dark:border-slate-700 rounded-lg text-slate-900 dark:text-white focus:ring-2 focus:ring-blue-500" required>
          </div>
          
          <div>
            <label class="text-xs font-bold text-slate-500 dark:text-slate-400">Warna / Visual</label>
            <input v-model="form.visual" placeholder="Contoh: Biola Green" class="w-full p-3 bg-slate-50 dark:bg-slate-900 border border-slate-200 dark:border-slate-700 rounded-lg text-slate-900 dark:text-white" required>
          </div>
          
          <div class="flex gap-4">
            <div class="w-1/2">
              <label class="text-xs font-bold text-slate-500 dark:text-slate-400">Jenis Kelamin</label>
              <select v-model="form.gender" class="w-full p-3 bg-slate-50 dark:bg-slate-900 border border-slate-200 dark:border-slate-700 rounded-lg text-slate-900 dark:text-white">
                <option value="U">❓ Unknown</option>
                <option value="M">♂️ Jantan</option>
                <option value="F">♀️ Betina</option>
              </select>
            </div>
            <div class="w-1/2">
              <label class="text-xs font-bold text-slate-500 dark:text-slate-400">Status</label>
              <select v-model="form.status" class="w-full p-3 bg-slate-50 dark:bg-slate-900 border border-slate-200 dark:border-slate-700 rounded-lg font-bold text-slate-900 dark:text-white">
                <option value="AVAILABLE">✅ Available</option>
                <option value="SOLD">❌ Sold</option>
              </select>
            </div>
          </div>
        </div>

        <div class="space-y-5 relative">
          <h4 class="font-bold text-sm text-purple-600 dark:text-purple-400 uppercase border-b dark:border-slate-700 pb-1">2. Data Indukan</h4>
          
          <div class="bg-blue-50 dark:bg-slate-800 p-4 rounded-lg border border-blue-100 dark:border-slate-700 relative">
            <label class="block text-xs font-bold text-blue-800 dark:text-blue-300 mb-2">Bapak (Sire)</label>
            <input 
              v-model="sireQuery" 
              @input="onSireInput"
              @focus="showSireList = true"
              placeholder="Ketik Ring Bapak..." 
              class="w-full p-3 border border-blue-200 dark:border-slate-600 rounded-lg text-sm focus:ring-2 focus:ring-blue-400 bg-white dark:bg-slate-900 text-slate-900 dark:text-white"
            >
            <ul v-if="showSireList && filteredSires.length > 0" class="absolute left-0 right-0 bg-white dark:bg-slate-800 border dark:border-slate-600 shadow-xl max-h-48 overflow-y-auto z-50 rounded-lg mx-4 mt-1">
              <li v-for="bird in filteredSires" :key="bird.id" @click="selectSire(bird)" class="p-3 hover:bg-blue-100 dark:hover:bg-slate-700 cursor-pointer text-sm border-b dark:border-slate-700 flex justify-between text-slate-700 dark:text-slate-200">
                <span class="font-bold text-blue-800 dark:text-blue-400">{{ bird.ring_number }}</span> 
                <span class="text-gray-600 dark:text-gray-400">{{ bird.visual }}</span>
              </li>
            </ul>
            <p class="text-[10px] text-blue-400 dark:text-slate-500 mt-1 italic">*Pilih dari list atau ketik manual jika dari luar.</p>
          </div>

          <div class="bg-pink-50 dark:bg-slate-800 p-4 rounded-lg border border-pink-100 dark:border-slate-700 relative">
            <label class="block text-xs font-bold text-pink-800 dark:text-pink-300 mb-2">Ibu (Dam)</label>
            <input 
              v-model="damQuery" 
              @input="onDamInput"
              @focus="showDamList = true"
              placeholder="Ketik Ring Ibu..." 
              class="w-full p-3 border border-pink-200 dark:border-slate-600 rounded-lg text-sm focus:ring-2 focus:ring-pink-400 bg-white dark:bg-slate-900 text-slate-900 dark:text-white"
            >
            <ul v-if="showDamList && filteredDams.length > 0" class="absolute left-0 right-0 bg-white dark:bg-slate-800 border dark:border-slate-600 shadow-xl max-h-48 overflow-y-auto z-50 rounded-lg mx-4 mt-1">
              <li v-for="bird in filteredDams" :key="bird.id" @click="selectDam(bird)" class="p-3 hover:bg-pink-100 dark:hover:bg-slate-700 cursor-pointer text-sm border-b dark:border-slate-700 flex justify-between text-slate-700 dark:text-slate-200">
                <span class="font-bold text-pink-800 dark:text-pink-400">{{ bird.ring_number }}</span> 
                <span class="text-gray-600 dark:text-gray-400">{{ bird.visual }}</span>
              </li>
            </ul>
            <p class="text-[10px] text-pink-400 dark:text-slate-500 mt-1 italic">*Pilih dari list atau ketik manual jika dari luar.</p>
          </div>
        </div>

        <button type="submit" class="md:col-span-2 bg-green-600 hover:bg-green-700 text-white py-3 rounded-lg font-bold shadow-lg mt-4 transition-colors">
          {{ loading ? 'Menyimpan...' : 'SIMPAN DATA' }}
        </button>
      </form>
    </div>

    <div v-if="filteredAnimals.length === 0" class="text-center py-20 text-gray-400">
      <p v-if="searchQuery">Tidak ada burung dengan kata kunci "<b>{{ searchQuery }}</b>"</p>
      <p v-else>Belum ada data burung.</p>
    </div>

    <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div v-for="bird in filteredAnimals" :key="bird.id" class="bg-white dark:bg-[#1e293b] rounded-2xl shadow-sm dark:shadow-lg hover:shadow-xl transition duration-300 overflow-hidden border border-slate-100 dark:border-slate-700 flex flex-col relative group">
        
        <div v-if="bird.status === 'SOLD'" class="absolute top-0 right-0 bg-red-600 text-white text-xs font-bold px-3 py-1 rounded-bl-lg z-10 shadow-sm">SOLD OUT</div>

        <div class="p-4 flex items-start justify-between bg-gradient-to-r from-gray-50 to-white dark:from-slate-800 dark:to-slate-700 pt-6 border-b dark:border-slate-600">
          <div>
            <h3 class="font-bold text-lg text-slate-800 dark:text-white transition-colors" :class="{'line-through text-slate-400 dark:text-slate-500': bird.status === 'SOLD'}">
              {{ bird.visual }}
            </h3>
            <p class="text-xs font-mono text-slate-500 dark:text-slate-300 bg-slate-200 dark:bg-slate-600 inline-block px-2 py-0.5 rounded mt-1">
              {{ bird.ring_number }}
            </p>
          </div>
          <span class="text-xl font-bold px-3 py-1 rounded-full shadow-sm" 
            :class="bird.gender === 'M' ? 'bg-blue-100 text-blue-600 dark:bg-blue-900/50 dark:text-blue-200' : bird.gender === 'F' ? 'bg-pink-100 text-pink-600 dark:bg-pink-900/50 dark:text-pink-200' : 'bg-gray-100 text-gray-500 dark:bg-slate-600 dark:text-gray-300'">
            {{ bird.gender === 'M' ? '♂' : bird.gender === 'F' ? '♀' : '?' }}
          </span>
        </div>

        <div class="px-4 py-3 border-t border-b border-slate-100 dark:border-slate-700 bg-slate-50 dark:bg-[#1e293b] text-sm space-y-2">
          <div class="flex items-center gap-2">
            <span class="text-blue-500 font-bold w-4">♂</span>
            <span v-if="bird.sire" class="text-slate-800 dark:text-slate-300 font-medium truncate">{{ bird.sire.visual }} ({{ bird.sire.ring_number }})</span>
            <span v-else-if="bird.sire_other" class="text-slate-600 dark:text-slate-400 italic truncate">{{ bird.sire_other }}</span>
            <span v-else class="text-slate-300 dark:text-slate-600 text-xs">-</span>
          </div>
          <div class="flex items-center gap-2">
            <span class="text-pink-500 font-bold w-4">♀</span>
            <span v-if="bird.dam" class="text-slate-800 dark:text-slate-300 font-medium truncate">{{ bird.dam.visual }} ({{ bird.dam.ring_number }})</span>
            <span v-else-if="bird.dam_other" class="text-slate-600 dark:text-slate-400 italic truncate">{{ bird.dam_other }}</span>
            <span v-else class="text-slate-300 dark:text-slate-600 text-xs">-</span>
          </div>
        </div>

        <div class="p-4 bg-gray-50 dark:bg-slate-800/50 flex justify-between items-center mt-auto text-xs text-gray-500 dark:text-gray-400">
           <span class="font-bold text-slate-400 dark:text-slate-500">DATA STOK</span>
           <span v-if="bird.status === 'AVAILABLE'" class="text-green-600 dark:text-green-400 font-bold bg-green-100 dark:bg-green-900/30 px-2 py-0.5 rounded">AVAILABLE</span>
        </div>
      </div>
    </div>
  </div>
</template>