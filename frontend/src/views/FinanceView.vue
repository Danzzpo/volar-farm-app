<script setup>
import { ref, onMounted, computed } from 'vue'
// Import Icon Lucide yang relevan untuk keuangan
import { 
  Wallet, TrendingUp, TrendingDown, Plus, Trash, 
  Search, Calendar, ArrowUpRight, ArrowDownRight, X, FileText, DollarSign
} from 'lucide-vue-next'

// --- STATE MANAGEMENT ---
const transactions = ref([])
const summary = ref({ total_income: 0, total_expense: 0, balance: 0 })
const loading = ref(true)
const isModalOpen = ref(false)
const isSubmitting = ref(false)
const searchQuery = ref('')
const userId = localStorage.getItem('user_id')

// Form Default State
const form = ref({
  type: 'Income', // Default Pemasukan
  category: '',
  amount: '',
  date: new Date().toISOString().split('T')[0], // Tanggal hari ini (YYYY-MM-DD)
  note: ''
})

// Opsi Kategori Dinamis (Berdasarkan Tipe)
const categories = computed(() => {
  return form.value.type === 'Income' 
    ? ['Penjualan Burung', 'Hadiah Lomba', 'Jual Pakan/Acc', 'Lainnya']
    : ['Beli Pakan', 'Vitamin & Obat', 'Beli Burung', 'Ring & Acc', 'Listrik & Air', 'Operasional Farm', 'Lainnya']
})

// --- API ACTIONS ---

// 1. Fetch Data (Summary & List)
const fetchData = async () => {
  loading.value = true
  try {
    // Ambil Summary & List secara paralel agar cepat
    const [resSum, resList] = await Promise.all([
      fetch(`http://localhost:8080/api/finance/summary?user_id=${userId}`),
      fetch(`http://localhost:8080/api/finance/transactions?user_id=${userId}`)
    ])

    const jsonSum = await resSum.json()
    const jsonList = await resList.json()

    summary.value = jsonSum.data || { total_income: 0, total_expense: 0, balance: 0 }
    transactions.value = jsonList.data || []

  } catch (err) { 
    console.error("Gagal ambil data keuangan:", err) 
  } finally { 
    loading.value = false 
  }
}

// 2. Simpan Transaksi Baru
const saveTransaction = async () => {
  if (!form.value.amount || !form.value.category) return alert("Jumlah dan Kategori wajib diisi!")
  if (form.value.amount <= 0) return alert("Nominal harus lebih dari 0")
  
  isSubmitting.value = true
  try {
    const payload = {
      user_id: parseInt(userId),
      type: form.value.type,
      category: form.value.category,
      amount: parseFloat(form.value.amount),
      date: new Date(form.value.date).toISOString(),
      note: form.value.note
    }

    const res = await fetch('http://localhost:8080/api/finance/transactions', {
      method: 'POST',
      headers: {'Content-Type': 'application/json'},
      body: JSON.stringify(payload)
    })

    if(res.ok) {
      closeModal()
      fetchData() // Refresh data otomatis tanpa reload
      // alert("✅ Transaksi berhasil disimpan!") // Opsional
    } else {
      const json = await res.json()
      alert("Gagal: " + json.error)
    }
  } catch(err) { 
    console.error(err)
    alert("Terjadi kesalahan sistem") 
  } finally { 
    isSubmitting.value = false 
  }
}

// 3. Hapus Transaksi
const deleteTransaction = async (id) => {
  if(!confirm("⚠️ Yakin ingin menghapus catatan ini? Saldo akan berubah.")) return
  
  try {
    const res = await fetch(`http://localhost:8080/api/finance/transactions/${id}`, { method: 'DELETE' })
    if(res.ok) {
        fetchData() // Refresh data
    }
  } catch(err) { console.error(err) }
}

// --- HELPERS & FORMATTERS ---
const openModal = () => {
  // Reset Form ke default saat dibuka
  form.value = { type: 'Income', category: '', amount: '', date: new Date().toISOString().split('T')[0], note: '' }
  isModalOpen.value = true
}
const closeModal = () => isModalOpen.value = false

// Format ke Rupiah (Rp 1.000.000)
const formatRupiah = (number) => {
  return new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR', minimumFractionDigits: 0, maximumFractionDigits: 0 }).format(number || 0)
}

// Format Tanggal Indo (12 Jan 2024)
const formatDate = (d) => new Date(d).toLocaleDateString('id-ID', {day: 'numeric', month: 'short', year: 'numeric'})

// Filter Pencarian di Tabel
const filteredList = computed(() => {
  if (!searchQuery.value) return transactions.value
  const lower = searchQuery.value.toLowerCase()
  return transactions.value.filter(t => 
    t.category.toLowerCase().includes(lower) || 
    (t.note && t.note.toLowerCase().includes(lower))
  )
})

onMounted(() => fetchData())
</script>

<template>
  <div class="space-y-6 p-2 pb-10 animate-fade-in">
    
    <div class="flex flex-col md:flex-row md:items-end justify-between gap-6">
      <div>
        <h2 class="text-3xl font-extrabold text-slate-800 dark:text-white flex items-center gap-3">
            <Wallet class="w-8 h-8 text-blue-600"/> Keuangan Farm
        </h2>
        <p class="text-slate-500 mt-1">Catat arus kas masuk dan keluar agar transparan.</p>
      </div>
      <button @click="openModal" class="btn-primary flex items-center justify-center gap-2">
        <Plus class="w-5 h-5" /> <span>Catat Transaksi</span>
      </button>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
      <div class="bg-gradient-to-br from-blue-600 to-indigo-700 rounded-2xl p-6 text-white shadow-lg relative overflow-hidden group hover:-translate-y-1 transition duration-300">
         <div class="absolute right-0 top-0 w-32 h-32 bg-white/10 rounded-full -mr-10 -mt-10 blur-2xl group-hover:scale-110 transition duration-700"></div>
         
         <div class="relative z-10">
             <p class="text-blue-100 text-sm font-bold uppercase tracking-wider mb-2 flex items-center gap-2">
                <Wallet class="w-4 h-4"/> Total Saldo Saat Ini
             </p>
             <h3 class="text-3xl font-extrabold tracking-tight truncate" :title="formatRupiah(summary.balance)">
                {{ formatRupiah(summary.balance) }}
             </h3>
         </div>
      </div>

      <div class="summary-card group">
         <div class="flex justify-between items-start">
            <div>
               <p class="text-slate-500 text-xs uppercase font-bold tracking-wider">Total Pemasukan</p>
               <h3 class="text-2xl font-bold text-green-600 mt-1 truncate">{{ formatRupiah(summary.total_income) }}</h3>
            </div>
            <div class="icon-wrapper bg-green-100 dark:bg-green-900/30 text-green-600 group-hover:bg-green-600 group-hover:text-white">
               <TrendingUp class="w-6 h-6" />
            </div>
         </div>
      </div>

      <div class="summary-card group">
         <div class="flex justify-between items-start">
            <div>
               <p class="text-slate-500 text-xs uppercase font-bold tracking-wider">Total Pengeluaran</p>
               <h3 class="text-2xl font-bold text-red-600 mt-1 truncate">{{ formatRupiah(summary.total_expense) }}</h3>
            </div>
            <div class="icon-wrapper bg-red-100 dark:bg-red-900/30 text-red-600 group-hover:bg-red-600 group-hover:text-white">
               <TrendingDown class="w-6 h-6" />
            </div>
         </div>
      </div>
    </div>

    <div class="bg-white dark:bg-slate-800 rounded-2xl shadow-sm border border-slate-200 dark:border-slate-700 overflow-hidden">
      <div class="p-5 border-b border-slate-100 dark:border-slate-700 flex flex-col sm:flex-row justify-between items-center gap-4">
         <h3 class="font-bold text-lg text-slate-800 dark:text-white flex items-center gap-2">
            <FileText class="w-5 h-5 text-slate-400"/> Riwayat Transaksi
         </h3>
         <div class="relative w-full sm:w-64">
            <Search class="absolute left-3 top-1/2 -translate-y-1/2 text-slate-400 w-4 h-4" />
            <input v-model="searchQuery" type="text" placeholder="Cari Kategori / Catatan..." class="w-full pl-10 pr-4 py-2.5 bg-slate-50 dark:bg-slate-900 border border-slate-200 dark:border-slate-700 rounded-xl text-sm outline-none focus:ring-2 focus:ring-blue-500 dark:text-white transition shadow-sm">
         </div>
      </div>

      <div v-if="loading" class="p-10 text-center flex flex-col items-center justify-center gap-3 text-slate-400 animate-pulse">
          <DollarSign class="w-10 h-10 opacity-50"/>
          <span>Sedang memuat data keuangan...</span>
      </div>

      <div v-else-if="filteredList.length === 0" class="p-16 text-center">
          <div class="bg-slate-100 dark:bg-slate-700/50 w-20 h-20 rounded-full flex items-center justify-center mx-auto mb-4 border-2 border-dashed border-slate-200 dark:border-slate-600">
              <FileText class="w-10 h-10 text-slate-300 dark:text-slate-500"/>
          </div>
          <p class="text-slate-600 dark:text-slate-300 font-bold text-lg">Belum ada data transaksi</p>
          <p class="text-sm text-slate-400 mt-1">Klik tombol "Catat Transaksi" di atas untuk memulai.</p>
      </div>

      <div v-else class="overflow-x-auto">
        <table class="w-full text-left text-sm">
          <thead class="bg-slate-50 dark:bg-slate-900/50 text-slate-500 uppercase font-bold text-xs border-b border-slate-100 dark:border-slate-700">
            <tr>
              <th class="px-6 py-4">Kategori & Catatan</th>
              <th class="px-6 py-4">Tanggal</th>
              <th class="px-6 py-4 text-right">Nominal</th>
              <th class="px-6 py-4 text-center w-20">Aksi</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-100 dark:divide-slate-700">
            <tr v-for="t in filteredList" :key="t.id" class="hover:bg-slate-50 dark:hover:bg-slate-700/50 transition group">
              <td class="px-6 py-4">
                 <div class="flex items-center gap-4">
                    <div class="p-3 rounded-xl border shadow-sm transition-colors" 
                        :class="t.type === 'Income' 
                            ? 'bg-green-50 border-green-100 text-green-600 group-hover:bg-green-100 dark:bg-green-900/20 dark:border-green-900/30' 
                            : 'bg-red-50 border-red-100 text-red-600 group-hover:bg-red-100 dark:bg-red-900/20 dark:border-red-900/30'">
                       <ArrowUpRight v-if="t.type === 'Income'" class="w-5 h-5" stroke-width="2.5" />
                       <ArrowDownRight v-else class="w-5 h-5" stroke-width="2.5" />
                    </div>
                    <div>
                       <p class="font-bold text-slate-800 dark:text-white text-base">{{ t.category }}</p>
                       <p class="text-xs text-slate-500 dark:text-slate-400 line-clamp-1" :title="t.note">{{ t.note || '-' }}</p>
                    </div>
                 </div>
              </td>
              <td class="px-6 py-4 text-slate-600 dark:text-slate-300 font-medium whitespace-nowrap">
                 <div class="flex items-center gap-2">
                    <Calendar class="w-4 h-4 text-slate-400" /> {{ formatDate(t.date) }}
                 </div>
              </td>
              <td class="px-6 py-4 text-right font-bold text-base whitespace-nowrap" :class="t.type === 'Income' ? 'text-green-600 dark:text-green-400' : 'text-red-600 dark:text-red-400'">
                 {{ t.type === 'Income' ? '+' : '-' }} {{ formatRupiah(t.amount) }}
              </td>
              <td class="px-6 py-4 text-center">
                 <button @click="deleteTransaction(t.id)" class="text-slate-400 hover:text-red-600 transition p-2 rounded-lg hover:bg-red-50 dark:hover:bg-red-900/20 opacity-0 group-hover:opacity-100" title="Hapus Transaksi">
                    <Trash class="w-5 h-5" />
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
          <h3 class="font-bold text-lg text-slate-800 dark:text-white flex items-center gap-2">
            <Plus class="w-5 h-5 text-blue-600"/> Transaksi Baru
          </h3>
          <button @click="closeModal" class="p-1 rounded-full hover:bg-slate-200 dark:hover:bg-slate-800 transition"><X class="w-5 h-5 text-slate-400 hover:text-slate-600"/></button>
        </div>

        <div class="p-6 space-y-5">
          
          <div class="flex bg-slate-100 dark:bg-slate-800 p-1.5 rounded-xl">
             <button @click="form.type = 'Income'" class="type-switch" :class="form.type === 'Income' ? 'bg-white dark:bg-slate-700 text-green-600 shadow-sm' : 'text-slate-500 dark:text-slate-400 hover:text-slate-700'">
               <ArrowUpRight class="w-5 h-5"/> Pemasukan
             </button>
             <button @click="form.type = 'Expense'" class="type-switch" :class="form.type === 'Expense' ? 'bg-white dark:bg-slate-700 text-red-600 shadow-sm' : 'text-slate-500 dark:text-slate-400 hover:text-slate-700'">
               <ArrowDownRight class="w-5 h-5"/> Pengeluaran
             </button>
          </div>

          <div>
             <label class="form-label">Kategori</label>
             <div class="relative">
                <select v-model="form.category" class="input-field font-bold appearance-none">
                    <option value="" disabled>-- Pilih Kategori --</option>
                    <option v-for="cat in categories" :key="cat" :value="cat">{{ cat }}</option>
                </select>
                <div class="absolute inset-y-0 right-0 flex items-center pr-4 pointer-events-none text-slate-400">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor"><path fill-rule="evenodd" d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z" clip-rule="evenodd" /></svg>
                </div>
             </div>
          </div>

          <div>
             <label class="form-label">Nominal (Rp)</label>
             <div class="relative">
                <span class="absolute left-4 top-1/2 -translate-y-1/2 text-slate-400 font-bold text-lg">Rp</span>
                <input v-model="form.amount" type="number" min="0" class="input-field pl-12 font-bold text-xl" placeholder="0">
             </div>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
             <div>
               <label class="form-label">Tanggal</label>
               <input v-model="form.date" type="date" class="input-field font-bold">
             </div>
             <div>
               <label class="form-label">Catatan (Opsional)</label>
               <input v-model="form.note" type="text" class="input-field" placeholder="Contoh: Jual Ring 001">
             </div>
          </div>

        </div>

        <div class="px-6 py-4 bg-slate-50 dark:bg-slate-900 border-t border-slate-100 dark:border-slate-800 flex gap-3">
          <button @click="closeModal" class="btn-secondary flex-1">Batal</button>
          <button @click="saveTransaction" :disabled="isSubmitting" class="btn-primary flex-1 flex items-center justify-center gap-2">
             <Plus v-if="!isSubmitting" class="w-5 h-5"/>
             <span>{{ isSubmitting ? 'Menyimpan...' : 'Simpan Data' }}</span>
          </button>
        </div>

      </div>
    </div>

  </div>
</template>

<style scoped>
/* --- UTILITY CLASSES (Agar kode HTML lebih bersih) --- */

/* Card Ringkasan Kecil (Hijau/Merah) */
.summary-card {
  @apply bg-white dark:bg-slate-800 rounded-2xl p-6 border border-slate-200 dark:border-slate-700 shadow-sm flex flex-col justify-between hover:shadow-md transition duration-300 hover:-translate-y-1;
}

/* Wrapper Icon di Card Kecil */
.icon-wrapper {
  @apply p-3 rounded-xl transition-colors duration-300;
}

/* Label Form di Modal */
.form-label {
  @apply block text-xs font-bold text-slate-500 uppercase tracking-wider mb-1.5;
}

/* Input Field Standar */
.input-field {
  @apply w-full p-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl text-sm outline-none focus:ring-2 focus:ring-blue-500 dark:text-white transition shadow-sm dark:focus:ring-blue-600;
}

/* Tombol Switch Pemasukan/Pengeluaran */
.type-switch {
  @apply flex-1 py-3 text-sm font-bold rounded-xl transition flex items-center justify-center gap-2;
}

/* Tombol Utama (Biru) */
.btn-primary {
  @apply bg-blue-600 hover:bg-blue-700 text-white px-5 py-2.5 rounded-xl font-bold shadow-lg shadow-blue-500/20 transition active:scale-95 disabled:opacity-50 disabled:cursor-not-allowed;
}

/* Tombol Sekunder (Putih/Abu) */
.btn-secondary {
  @apply py-3 rounded-xl border border-slate-300 dark:border-slate-700 font-bold text-sm text-slate-500 dark:text-slate-400 hover:bg-slate-100 dark:hover:bg-slate-800 transition;
}

/* --- ANIMATIONS --- */
@keyframes fadeInUp { 
    from { opacity: 0; transform: translateY(10px) scale(0.98); } 
    to { opacity: 1; transform: translateY(0) scale(1); } 
}
.animate-fade-in-up { 
    animation: fadeInUp 0.3s cubic-bezier(0.16, 1, 0.3, 1) forwards; 
}
</style>