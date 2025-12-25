<script setup>
import { ref, onMounted } from 'vue'
import { 
  User, Lock, Moon, Store, Bell, Shield, LogOut, 
  Camera, MapPin, Smartphone, Mail, Save, ChevronRight
} from 'lucide-vue-next'
import { useRouter } from 'vue-router'

const router = useRouter()
const activeTab = ref('profile') // Default tab: profile, farm, security, preference
const isLoading = ref(false)
const isDarkMode = ref(false)

// Data State
const form = ref({
  // Profile
  name: localStorage.getItem('username') || '',
  email: 'admin@volarfarm.com',
  phone: '0812-3456-7890',
  avatar: null,
  
  // Farm
  farm_name: localStorage.getItem('farm_name') || '',
  address: 'Jl. Peternakan No. 1, Malang',
  est: '2023',
  
  // Security
  current_password: '',
  new_password: '',
  
  // Preferences
  notif_wa: true,
  notif_email: false
})

// --- LIFECYCLE ---
onMounted(() => {
  isDarkMode.value = document.documentElement.classList.contains('dark')
})

// --- ACTIONS ---
const toggleTheme = () => {
  isDarkMode.value = !isDarkMode.value
  if (isDarkMode.value) {
    document.documentElement.classList.add('dark')
    localStorage.setItem('theme', 'dark')
  } else {
    document.documentElement.classList.remove('dark')
    localStorage.setItem('theme', 'light')
  }
}

const saveSettings = () => {
  isLoading.value = true
  setTimeout(() => {
    // Simpan ke LocalStorage
    localStorage.setItem('username', form.value.name)
    localStorage.setItem('farm_name', form.value.farm_name)
    
    isLoading.value = false
    alert(`✅ Pengaturan ${activeTab.value} berhasil disimpan!`)
    
    // Refresh halaman jika nama berubah agar sidebar update
    if (activeTab.value === 'profile' || activeTab.value === 'farm') {
        window.location.reload()
    }
  }, 800)
}

const handleLogout = () => {
  if (confirm('Keluar dari aplikasi?')) {
    localStorage.clear()
    router.push('/login')
  }
}

// Menu Navigasi Samping
const tabs = [
  { id: 'profile', label: 'Profil Akun', icon: User, desc: 'Identitas & Kontak' },
  { id: 'farm', label: 'Data Farm', icon: Store, desc: 'Informasi Peternakan' },
  { id: 'security', label: 'Keamanan', icon: Shield, desc: 'Password & Akses' },
  { id: 'preference', label: 'Tampilan', icon: Moon, desc: 'Tema & Notifikasi' },
]
</script>

<template>
  <div class="max-w-6xl mx-auto p-2 pb-10 animate-fade-in">

    <div class="mb-8">
      <h2 class="text-3xl font-extrabold text-slate-800 dark:text-white">Pengaturan</h2>
      <p class="text-slate-500 dark:text-slate-400 mt-1">Kelola akun, profil farm, dan preferensi aplikasi Anda.</p>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-4 gap-8">
      
      <div class="lg:col-span-1 space-y-2">
        <button 
          v-for="tab in tabs" 
          :key="tab.id"
          @click="activeTab = tab.id"
          class="w-full flex items-center justify-between p-4 rounded-xl transition-all duration-200 border text-left group"
          :class="activeTab === tab.id 
            ? 'bg-blue-600 text-white border-blue-600 shadow-lg shadow-blue-500/30' 
            : 'bg-white dark:bg-slate-800 border-transparent hover:bg-slate-50 dark:hover:bg-slate-700 text-slate-600 dark:text-slate-300'"
        >
          <div class="flex items-center gap-3">
             <component :is="tab.icon" class="w-5 h-5" :class="activeTab === tab.id ? 'text-white' : 'text-slate-400 group-hover:text-blue-500'" />
             <div>
               <p class="font-bold text-sm">{{ tab.label }}</p>
               <p class="text-[10px] opacity-80" :class="activeTab === tab.id ? 'text-blue-100' : 'text-slate-400'">{{ tab.desc }}</p>
             </div>
          </div>
          <ChevronRight v-if="activeTab === tab.id" class="w-4 h-4 text-white" />
        </button>

        <button @click="handleLogout" class="w-full flex items-center gap-3 p-4 rounded-xl text-red-500 hover:bg-red-50 dark:hover:bg-red-900/20 transition mt-6 font-bold text-sm">
           <LogOut class="w-5 h-5" /> Keluar
        </button>
      </div>

      <div class="lg:col-span-3">
        <div class="bg-white dark:bg-slate-800 rounded-2xl shadow-sm border border-slate-200 dark:border-slate-700 p-6 md:p-8 min-h-[400px]">
          
          <div v-if="activeTab === 'profile'" class="space-y-6 animate-fade-in-up">
             <div class="flex items-center gap-4 mb-6">
                <div class="w-20 h-20 rounded-full bg-slate-100 dark:bg-slate-700 flex items-center justify-center text-3xl overflow-hidden relative group cursor-pointer border-2 border-slate-200 dark:border-slate-600">
                   <span v-if="!form.avatar">👨‍🌾</span>
                   <img v-else :src="form.avatar" class="w-full h-full object-cover">
                   <div class="absolute inset-0 bg-black/40 flex items-center justify-center opacity-0 group-hover:opacity-100 transition">
                      <Camera class="w-6 h-6 text-white" />
                   </div>
                </div>
                <div>
                  <h3 class="text-lg font-bold text-slate-800 dark:text-white">Foto Profil</h3>
                  <p class="text-xs text-slate-500 max-w-xs">Klik foto untuk mengubah. Format JPG, PNG maks 2MB.</p>
                </div>
             </div>

             <div class="grid grid-cols-1 md:grid-cols-2 gap-5">
               <div class="space-y-1.5">
                  <label class="form-label">Nama Lengkap</label>
                  <div class="input-wrapper">
                    <User class="input-icon" />
                    <input v-model="form.name" type="text" class="input-field pl-10">
                  </div>
               </div>
               <div class="space-y-1.5">
                  <label class="form-label">Email</label>
                  <div class="input-wrapper">
                    <Mail class="input-icon" />
                    <input v-model="form.email" type="email" class="input-field pl-10 bg-slate-50 dark:bg-slate-900/50 cursor-not-allowed" disabled>
                  </div>
               </div>
               <div class="space-y-1.5 md:col-span-2">
                  <label class="form-label">No. WhatsApp</label>
                  <div class="input-wrapper">
                    <Smartphone class="input-icon" />
                    <input v-model="form.phone" type="text" class="input-field pl-10">
                  </div>
               </div>
             </div>
          </div>

          <div v-if="activeTab === 'farm'" class="space-y-6 animate-fade-in-up">
             <div class="border-l-4 border-blue-500 pl-4 py-1 mb-6 bg-blue-50 dark:bg-blue-900/20 rounded-r-lg">
                <h3 class="font-bold text-slate-800 dark:text-white">Identitas Peternakan</h3>
                <p class="text-sm text-slate-500 dark:text-slate-400">Data ini akan muncul di laporan dan sertifikat silsilah (pedigree).</p>
             </div>

             <div class="grid grid-cols-1 gap-5">
               <div class="space-y-1.5">
                  <label class="form-label">Nama Farm</label>
                  <div class="input-wrapper">
                    <Store class="input-icon" />
                    <input v-model="form.farm_name" type="text" class="input-field pl-10" placeholder="Contoh: Juara Bird Farm">
                  </div>
               </div>
               <div class="space-y-1.5">
                  <label class="form-label">Alamat Farm</label>
                  <div class="input-wrapper">
                    <MapPin class="input-icon" />
                    <textarea v-model="form.address" rows="3" class="input-field pl-10 py-2.5 resize-none"></textarea>
                  </div>
               </div>
               <div class="space-y-1.5">
                  <label class="form-label">Tahun Berdiri (Est)</label>
                  <input v-model="form.est" type="text" class="input-field w-full">
               </div>
             </div>
          </div>

          <div v-if="activeTab === 'security'" class="space-y-6 animate-fade-in-up">
             <div class="space-y-1.5">
                <label class="form-label">Password Saat Ini</label>
                <div class="input-wrapper">
                  <Lock class="input-icon" />
                  <input v-model="form.current_password" type="password" class="input-field pl-10" placeholder="••••••">
                </div>
             </div>
             
             <hr class="border-slate-100 dark:border-slate-700 my-4">

             <div class="grid grid-cols-1 md:grid-cols-2 gap-5">
               <div class="space-y-1.5">
                  <label class="form-label">Password Baru</label>
                  <div class="input-wrapper">
                    <Lock class="input-icon" />
                    <input v-model="form.new_password" type="password" class="input-field pl-10" placeholder="Ketik password baru">
                  </div>
               </div>
               <div class="space-y-1.5">
                  <label class="form-label">Konfirmasi Password</label>
                  <div class="input-wrapper">
                    <Lock class="input-icon" />
                    <input type="password" class="input-field pl-10" placeholder="Ulangi password baru">
                  </div>
               </div>
             </div>
          </div>

          <div v-if="activeTab === 'preference'" class="space-y-6 animate-fade-in-up">
             
             <div class="flex items-center justify-between p-4 border border-slate-200 dark:border-slate-700 rounded-xl bg-slate-50 dark:bg-slate-900/30">
                <div class="flex items-center gap-3">
                   <div class="p-2 rounded-full" :class="isDarkMode ? 'bg-slate-700 text-yellow-400' : 'bg-orange-100 text-orange-500'">
                      <component :is="isDarkMode ? Moon : 'Sun'" class="w-6 h-6" />
                   </div>
                   <div>
                      <p class="font-bold text-slate-800 dark:text-white">Mode Gelap</p>
                      <p class="text-xs text-slate-500">Sesuaikan tema aplikasi.</p>
                   </div>
                </div>
                <button @click="toggleTheme" class="relative w-12 h-6 rounded-full transition-colors duration-300" :class="isDarkMode ? 'bg-blue-600' : 'bg-slate-300'">
                   <div class="absolute top-1 left-1 bg-white w-4 h-4 rounded-full transition-transform duration-300" :class="isDarkMode ? 'translate-x-6' : ''"></div>
                </button>
             </div>

             <div class="space-y-4 pt-4">
               <h4 class="font-bold text-sm text-slate-800 dark:text-white uppercase tracking-wider">Notifikasi</h4>
               
               <div class="flex items-center justify-between">
                  <div class="flex items-center gap-3">
                     <div class="p-2 bg-green-100 dark:bg-green-900/30 text-green-600 rounded-lg"><Smartphone class="w-5 h-5"/></div>
                     <span class="text-sm font-medium text-slate-700 dark:text-slate-300">Notifikasi WhatsApp</span>
                  </div>
                  <input type="checkbox" v-model="form.notif_wa" class="w-5 h-5 rounded text-blue-600 focus:ring-blue-500">
               </div>

               <div class="flex items-center justify-between">
                  <div class="flex items-center gap-3">
                     <div class="p-2 bg-purple-100 dark:bg-purple-900/30 text-purple-600 rounded-lg"><Bell class="w-5 h-5"/></div>
                     <span class="text-sm font-medium text-slate-700 dark:text-slate-300">Notifikasi Email</span>
                  </div>
                  <input type="checkbox" v-model="form.notif_email" class="w-5 h-5 rounded text-blue-600 focus:ring-blue-500">
               </div>
             </div>
          </div>

          <div v-if="activeTab !== 'preference'" class="mt-8 pt-6 border-t border-slate-100 dark:border-slate-700 flex justify-end">
             <button @click="saveSettings" :disabled="isLoading" class="btn-primary">
                <Save class="w-4 h-4" />
                <span>{{ isLoading ? 'Menyimpan...' : 'Simpan Perubahan' }}</span>
             </button>
          </div>

        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
/* Utility CSS untuk kebersihan kode */
.form-label {
  @apply block text-xs font-bold text-slate-500 uppercase tracking-wider mb-1.5;
}

.input-wrapper {
  @apply relative;
}

/* Teknik Absolute Center Icon yang Paling Aman */
.input-icon {
  @apply absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-400 pointer-events-none;
}

.input-field {
  @apply w-full py-3 pr-4 bg-slate-50 dark:bg-slate-900 border border-slate-200 dark:border-slate-700 rounded-xl text-sm outline-none focus:ring-2 focus:ring-blue-500 dark:text-white transition shadow-sm;
}

.btn-primary {
  @apply inline-flex items-center gap-2 bg-blue-600 hover:bg-blue-700 text-white px-8 py-3 rounded-xl font-bold text-sm shadow-lg shadow-blue-500/20 transition active:scale-95 disabled:opacity-50 disabled:cursor-not-allowed;
}

@keyframes fadeInUp {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}
.animate-fade-in-up {
  animation: fadeInUp 0.3s ease-out forwards;
}
</style>