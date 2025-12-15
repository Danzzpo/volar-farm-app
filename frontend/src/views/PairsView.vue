<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const pairs = ref([])
const animals = ref([]) // Untuk pilihan dropdown
const showForm = ref(false)
const loading = ref(false)

const userId = localStorage.getItem('user_id')

const form = ref({
  male_id: null,
  female_id: null,
  cage: '',
  pair_date: new Date().toISOString().split('T')[0], // Hari ini
  notes: ''
})

// --- FILTER BURUNG UNTUK DROPDOWN ---
// Hanya tampilkan yang gender sesuai DAN statusnya AVAILABLE/KEEPER (bukan SOLD)
const males = computed(() => animals.value.filter(a => a.gender === 'M' && a.status !== 'SOLD'))
const females = computed(() => animals.value.filter(a => a.gender === 'F' && a.status !== 'SOLD'))

onMounted(() => {
  if (!userId) { router.push('/login'); return }
  fetchAnimals()
  fetchPairs()
})

// 1. Ambil Stok Burung (Buat Dropdown)
const fetchAnimals = async () => {
  try {
    const res = await fetch(`http://localhost:8080/api/animals?user_id=${userId}`)
    const json = await res.json()
    animals.value = json.data || []
  } catch (err) { console.error(err) }
}

// 2. Ambil Daftar Pasangan (List Pairs)
const fetchPairs = async () => {
  try {
    const res = await fetch(`http://localhost:8080/api/pairs?user_id=${userId}`)
    const json = await res.json()
    pairs.value = json.data || []
  } catch (err) { console.error(err) }
}

// 3. Simpan Pasangan Baru
const submitPair = async () => {
  if(!form.value.male_id || !form.value.female_id || !form.value.cage) {
    alert("Jantan, Betina, dan Kandang wajib diisi!")
    return
  }
  
  loading.value = true
  const payload = {
    user_id: parseInt(userId),
    ...form.value,
    pair_date: new Date(form.value.pair_date).toISOString() // Format Go Time
  }

  try {
    const res = await fetch('http://localhost:8080/api/pairs', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload)
    })

    if (res.ok) {
      alert("✅ Pasangan Berhasil Dibuat!")
      showForm.value = false
      fetchPairs() // Refresh list
      // Reset form (kecuali tanggal biar enak kalau input banyak)
      form.value.male_id = null
      form.value.female_id = null
      form.value.cage = ''
      form.value.notes = ''
    } else {
      alert("Gagal menyimpan!")
    }
  } catch (err) { console.error(err) } finally { loading.value = false }
}

// 4. Ubah Status (Misal: Bubar Jalan / Cabut)
const updateStatus = async (pairId, newStatus) => {
  if(!confirm(`Ubah status menjadi ${newStatus}?`)) return
  try {
    await fetch(`http://localhost:8080/api/pairs/${pairId}`, {
      method: 'PATCH',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ status: newStatus })
    })
    fetchPairs()
  } catch(err) { console.error(err) }
}

const formatDate = (d) => new Date(d).toLocaleDateString('id-ID', {day: 'numeric', month: 'short', year: 'numeric'})
</script>

<template>
  <div class="animate-fade-in">
    <div class="flex justify-between items-center mb-6">
      <div>
        <h2 class="text-2xl font-bold text-slate-800">Pasangan Indukan (Pairs)</h2>
        <p class="text-slate-500">Kelola pasangan aktif dalam kandang ternak.</p>
      </div>
      <button @click="showForm = !showForm" class="bg-pink-600 text-white px-6 py-2 rounded-full font-bold hover:bg-pink-700 shadow-lg flex items-center gap-2">
        <span>💞</span> Tambah Pasangan
      </button>
    </div>

    <div v-if="showForm" class="bg-white p-6 rounded-xl shadow-lg border border-pink-100 mb-8 relative z-20">
      <h3 class="font-bold text-slate-700 border-b pb-2 mb-4">Form Penjodohan</h3>
      <form @submit.prevent="submitPair" class="grid grid-cols-1 md:grid-cols-2 gap-6">
        
        <div class="bg-blue-50 p-4 rounded-lg">
          <label class="block text-xs font-bold text-blue-800 mb-2">Pilih Pejantan (Sire)</label>
          <select v-model="form.male_id" class="w-full p-2 border rounded" required>
            <option :value="null">-- Pilih Jantan Available --</option>
            <option v-for="b in males" :key="b.id" :value="b.id">
              {{ b.ring_number }} - {{ b.visual }}
            </option>
          </select>
        </div>

        <div class="bg-pink-50 p-4 rounded-lg">
          <label class="block text-xs font-bold text-pink-800 mb-2">Pilih Betina (Dam)</label>
          <select v-model="form.female_id" class="w-full p-2 border rounded" required>
            <option :value="null">-- Pilih Betina Available --</option>
            <option v-for="b in females" :key="b.id" :value="b.id">
              {{ b.ring_number }} - {{ b.visual }}
            </option>
          </select>
        </div>

        <div>
          <label class="block text-xs font-bold text-slate-500 mb-1">Nomor/Nama Kandang</label>
          <input v-model="form.cage" placeholder="Contoh: Glodok A1" class="w-full p-3 border rounded focus:ring-2 focus:ring-pink-500" required>
        </div>

        <div>
          <label class="block text-xs font-bold text-slate-500 mb-1">Tanggal Masuk</label>
          <input v-model="form.pair_date" type="date" class="w-full p-3 border rounded" required>
        </div>

        <div class="md:col-span-2">
          <label class="block text-xs font-bold text-slate-500 mb-1">Catatan (Opsional)</label>
          <textarea v-model="form.notes" placeholder="Catatan kondisi penjodohan..." class="w-full p-3 border rounded h-20"></textarea>
        </div>

        <button type="submit" class="md:col-span-2 bg-pink-600 text-white py-3 rounded-lg font-bold hover:bg-pink-700 shadow-md">
          {{ loading ? 'Menyimpan...' : 'SIMPAN PASANGAN' }}
        </button>
      </form>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div v-for="pair in pairs" :key="pair.id" class="bg-white rounded-xl shadow-sm border border-slate-100 overflow-hidden hover:shadow-lg transition">
        
        <div class="p-3 flex justify-between items-center" 
             :class="pair.status === 'ACTIVE' ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-500'">
          <div class="flex items-center gap-2 font-bold">
            <span class="text-xl">🏠</span> {{ pair.cage }}
          </div>
          <span class="text-xs font-bold px-2 py-1 rounded bg-white/50 border border-white/20">
            {{ pair.status }}
          </span>
        </div>

        <div class="p-4 space-y-3">
          <div class="flex items-center gap-3 p-2 bg-blue-50 rounded-lg border border-blue-100">
            <div class="w-8 h-8 rounded-full bg-blue-200 flex items-center justify-center text-blue-700 font-bold text-xs">♂</div>
            <div class="overflow-hidden">
              <p class="font-bold text-sm text-slate-700 truncate">{{ pair.male?.visual || '?' }}</p>
              <p class="text-xs text-slate-500">{{ pair.male?.ring_number }}</p>
            </div>
          </div>

          <div class="flex items-center gap-3 p-2 bg-pink-50 rounded-lg border border-pink-100">
            <div class="w-8 h-8 rounded-full bg-pink-200 flex items-center justify-center text-pink-700 font-bold text-xs">♀</div>
            <div class="overflow-hidden">
              <p class="font-bold text-sm text-slate-700 truncate">{{ pair.female?.visual || '?' }}</p>
              <p class="text-xs text-slate-500">{{ pair.female?.ring_number }}</p>
            </div>
          </div>
          
          <div class="text-xs text-slate-400 pt-2 border-t border-dashed border-slate-200 flex justify-between">
            <span>📅 {{ formatDate(pair.pair_date) }}</span>
            <span v-if="pair.notes" :title="pair.notes">📝 Info</span>
          </div>

          <div v-if="pair.status === 'ACTIVE'" class="pt-2 flex gap-2">
            <button @click="updateStatus(pair.id, 'REST')" class="flex-1 py-1 text-xs font-bold border border-orange-300 text-orange-600 rounded hover:bg-orange-50">Istirahatkan</button>
            <button @click="updateStatus(pair.id, 'HISTORY')" class="flex-1 py-1 text-xs font-bold border border-gray-300 text-gray-600 rounded hover:bg-gray-50">Cabut Jodoh</button>
          </div>
        </div>

      </div>
    </div>
  </div>
</template>