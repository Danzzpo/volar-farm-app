<script setup>
import { ref, onMounted, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { 
  Pill, Syringe, Droplets, Calendar, Plus, Trash, Search, FileText, X, CheckCircle 
} from 'lucide-vue-next'

const { t } = useI18n()
const userId = localStorage.getItem('user_id')

// State
const treatments = ref([])
const loading = ref(true)
const isModalOpen = ref(false)
const isSubmitting = ref(false)
const searchQuery = ref('')

// Form Default
const form = ref({
  type: 'Vitamin',
  name: '',
  date: new Date().toISOString().split('T')[0],
  next_date: '',
  note: ''
})

// Pilihan Jenis Perawatan
const types = ['Vitamin', 'Obat', 'Vaksin', 'Disinfektan', 'Lainnya']

// --- API ACTIONS ---
const fetchData = async () => {
  loading.value = true
  try {
    const res = await fetch(`http://localhost:8080/api/treatments?user_id=${userId}`)
    const json = await res.json()
    treatments.value = json.data || []
  } catch (e) { console.error(e) } 
  finally { loading.value = false }
}

const saveTreatment = async () => {
  if (!form.value.name) return alert("Nama obat wajib diisi!")
  isSubmitting.value = true
  
  try {
    const payload = {
      user_id: parseInt(userId),
      type: form.value.type,
      name: form.value.name,
      date: new Date(form.value.date).toISOString(),
      // Kirim null jika next_date kosong
      next_date: form.value.next_date ? new Date(form.value.next_date).toISOString() : null,
      note: form.value.note
    }

    const res = await fetch('http://localhost:8080/api/treatments', {
      method: 'POST',
      headers: {'Content-Type': 'application/json'},
      body: JSON.stringify(payload)
    })

    if (res.ok) {
      closeModal()
      fetchData()
    } else {
      alert("Gagal menyimpan")
    }
  } catch (e) { console.error(e) } 
  finally { isSubmitting.value = false }
}

const deleteItem = async (id) => {
  if(!confirm("Hapus data ini?")) return
  try {
    await fetch(`http://localhost:8080/api/treatments/${id}`, { method: 'DELETE' })
    fetchData()
  } catch(e) { console.error(e) }
}

// --- HELPERS ---
const openModal = () => {
  form.value = { type: 'Vitamin', name: '', date: new Date().toISOString().split('T')[0], next_date: '', note: '' }
  isModalOpen.value = true
}
const closeModal = () => isModalOpen.value = false

const formatDate = (d) => {
  if (!d) return '-'
  return new Date(d).toLocaleDateString('id-ID', { day: 'numeric', month: 'short', year: 'numeric' })
}

// Filter Pencarian
const filteredList = computed(() => {
  if (!searchQuery.value) return treatments.value
  const lower = searchQuery.value.toLowerCase()
  return treatments.value.filter(item => 
    item.name.toLowerCase().includes(lower) || 
    item.note.toLowerCase().includes(lower)
  )
})

// Icon Helper
const getTypeIcon = (type) => {
  if (type === 'Vaksin') return Syringe
  if (type === 'Vitamin') return Droplets
  return Pill
}
// Color Helper
const getTypeColor = (type) => {
  if (type === 'Vaksin') return 'text-purple-600 bg-purple-100 dark:bg-purple-900/30'
  if (type === 'Vitamin') return 'text-orange-600 bg-orange-100 dark:bg-orange-900/30'
  return 'text-blue-600 bg-blue-100 dark:bg-blue-900/30'
}

onMounted(() => fetchData())
</script>

<template>
  <div class="space-y-6 animate-fade-in p-2 pb-10">
    
    <div class="flex flex-col md:flex-row md:items-end justify-between gap-6">
      <div>
        <h2 class="text-3xl font-extrabold text-slate-800 dark:text-white flex items-center gap-3">
            <Pill class="w-8 h-8 text-blue-600"/> {{ t('treatments.title') }}
        </h2>
        <p class="text-slate-500 mt-1">{{ t('treatments.subtitle') }}</p>
      </div>
      <button @click="openModal" class="bg-blue-600 hover:bg-blue-700 text-white px-5 py-2.5 rounded-xl font-bold flex items-center justify-center gap-2 shadow-lg shadow-blue-500/20 active:scale-95 transition">
        <Plus class="w-5 h-5" /> <span>{{ t('treatments.add_btn') }}</span>
      </button>
    </div>

    <div class="bg-white dark:bg-slate-800 rounded-2xl shadow-sm border border-slate-200 dark:border-slate-700 overflow-hidden">
      
      <div class="p-5 border-b border-slate-100 dark:border-slate-700 flex flex-col sm:flex-row justify-between items-center gap-4">
         <h3 class="font-bold text-lg text-slate-800 dark:text-white flex items-center gap-2">
            <FileText class="w-5 h-5 text-slate-400"/> Riwayat Medis
         </h3>
         <div class="relative w-full sm:w-64">
            <Search class="absolute left-3 top-1/2 -translate-y-1/2 text-slate-400 w-4 h-4" />
            <input v-model="searchQuery" type="text" placeholder="Cari nama obat..." class="w-full pl-10 pr-4 py-2 bg-slate-50 dark:bg-slate-900 border border-slate-200 dark:border-slate-700 rounded-xl text-sm outline-none focus:ring-2 focus:ring-blue-500 dark:text-white transition">
         </div>
      </div>

      <div v-if="loading" class="p-10 text-center text-slate-400">Memuat data...</div>
      
      <div v-else-if="filteredList.length === 0" class="p-16 text-center">
          <div class="bg-slate-100 dark:bg-slate-700 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-4">
              <Pill class="w-8 h-8 text-slate-400"/>
          </div>
          <p class="text-slate-500 font-bold">Belum ada riwayat perawatan</p>
      </div>

      <div v-else class="overflow-x-auto">
        <table class="w-full text-left text-sm">
          <thead class="bg-slate-50 dark:bg-slate-900/50 text-slate-500 uppercase font-bold text-xs border-b border-slate-100 dark:border-slate-700">
            <tr>
              <th class="px-6 py-4">{{ t('treatments.date') }}</th>
              <th class="px-6 py-4">{{ t('treatments.name') }}</th>
              <th class="px-6 py-4">{{ t('treatments.note') }}</th>
              <th class="px-6 py-4">{{ t('treatments.next_date') }}</th>
              <th class="px-6 py-4 text-center">Aksi</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-100 dark:divide-slate-700">
            <tr v-for="item in filteredList" :key="item.id" class="hover:bg-slate-50 dark:hover:bg-slate-700/50 transition">
              
              <td class="px-6 py-4">
                 <div class="flex items-center gap-3">
                    <div :class="`p-2 rounded-lg ${getTypeColor(item.type)}`">
                       <component :is="getTypeIcon(item.type)" class="w-5 h-5" />
                    </div>
                    <div>
                       <p class="font-bold text-slate-700 dark:text-white">{{ formatDate(item.date) }}</p>
                       <p class="text-xs text-slate-500">{{ item.type }}</p>
                    </div>
                 </div>
              </td>

              <td class="px-6 py-4 font-bold text-slate-700 dark:text-white text-base">
                 {{ item.name }}
              </td>

              <td class="px-6 py-4 text-slate-500 italic max-w-xs truncate">
                 {{ item.note || '-' }}
              </td>

              <td class="px-6 py-4">
                 <div v-if="item.next_date" class="inline-flex items-center gap-2 px-3 py-1 bg-blue-50 dark:bg-blue-900/20 text-blue-600 dark:text-blue-300 rounded-full text-xs font-bold border border-blue-100 dark:border-blue-800">
                    <Calendar class="w-3 h-3"/> {{ formatDate(item.next_date) }}
                 </div>
                 <span v-else class="text-slate-400 text-xs">-</span>
              </td>

              <td class="px-6 py-4 text-center">
                 <button @click="deleteItem(item.id)" class="text-slate-400 hover:text-red-600 p-2 rounded-lg hover:bg-red-50 transition">
                    <Trash class="w-4 h-4" />
                 </button>
              </td>

            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <div v-if="isModalOpen" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-slate-900/60 backdrop-blur-sm" @click.self="closeModal">
      <div class="bg-white dark:bg-slate-900 w-full max-w-md rounded-2xl shadow-2xl overflow-hidden animate-fade-in-up">
        
        <div class="px-6 py-4 border-b border-slate-100 dark:border-slate-800 flex justify-between items-center bg-slate-50 dark:bg-slate-900">
          <h3 class="font-bold text-lg text-slate-800 dark:text-white">Catat Perawatan</h3>
          <button @click="closeModal"><X class="w-5 h-5 text-slate-400 hover:text-slate-600"/></button>
        </div>

        <div class="p-6 space-y-4">
          
          <div>
             <label class="block text-xs font-bold text-slate-500 uppercase mb-1.5">Jenis</label>
             <div class="flex gap-2">
                <button v-for="t in types" :key="t" @click="form.type = t"
                  class="px-3 py-1.5 text-xs font-bold rounded-lg border transition"
                  :class="form.type === t ? 'bg-blue-600 text-white border-blue-600' : 'bg-slate-50 border-slate-200 text-slate-500 hover:bg-slate-100 dark:bg-slate-800 dark:border-slate-700'">
                  {{ t }}
                </button>
             </div>
          </div>

          <div>
             <label class="block text-xs font-bold text-slate-500 uppercase mb-1.5">{{ t('treatments.name') }}</label>
             <input v-model="form.name" type="text" class="w-full p-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl text-sm outline-none font-bold" placeholder="Contoh: Super N">
          </div>

          <div class="grid grid-cols-2 gap-4">
             <div>
               <label class="block text-xs font-bold text-slate-500 uppercase mb-1.5">{{ t('treatments.date') }}</label>
               <input v-model="form.date" type="date" class="w-full p-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl text-sm font-bold">
             </div>
             <div>
               <label class="block text-xs font-bold text-slate-500 uppercase mb-1.5">{{ t('treatments.next_date') }} (Opsional)</label>
               <input v-model="form.next_date" type="date" class="w-full p-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl text-sm font-bold">
             </div>
          </div>

          <div>
             <label class="block text-xs font-bold text-slate-500 uppercase mb-1.5">{{ t('treatments.note') }}</label>
             <textarea v-model="form.note" rows="2" class="w-full p-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl text-sm outline-none" placeholder="Dosis, alasan, dll..."></textarea>
          </div>

        </div>

        <div class="px-6 py-4 bg-slate-50 dark:bg-slate-900 border-t border-slate-100 dark:border-slate-700 flex gap-3">
          <button @click="closeModal" class="flex-1 py-3 rounded-xl border border-slate-300 font-bold text-sm text-slate-500 hover:bg-slate-100">Batal</button>
          <button @click="saveTreatment" :disabled="isSubmitting" class="flex-1 py-3 rounded-xl bg-blue-600 text-white font-bold text-sm hover:bg-blue-700 shadow-lg shadow-blue-500/30">
             {{ isSubmitting ? 'Menyimpan...' : 'Simpan' }}
          </button>
        </div>

      </div>
    </div>

  </div>
</template>