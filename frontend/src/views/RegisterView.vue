<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
// Import Icon dari Lucide
import { User, Mail, Lock, Home, ArrowRight, Loader2, Bird, LogIn } from 'lucide-vue-next'
// Import Animasi AOS
import AOS from 'aos'
import 'aos/dist/aos.css'

const router = useRouter()

// Data Form
const form = ref({
  username: '',
  email: '',
  password: '',
  farm_name: ''
})

const loading = ref(false)

// Inisialisasi Animasi
onMounted(() => {
  AOS.init({ duration: 800, once: true })
})

const handleRegister = async () => {
  loading.value = true
  try {
    const response = await fetch('http://localhost:8080/api/register', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(form.value)
    })

    const result = await response.json()

    if (response.ok) {
      alert("✅ Pendaftaran Berhasil! Silakan Login.")
      router.push('/login')
    } else {
      alert("❌ Gagal: " + result.error)
    }

  } catch (error) {
    console.error(error)
    alert("Terjadi kesalahan sistem!")
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="min-h-[100dvh] flex items-center justify-center bg-[#0f172a] relative overflow-hidden font-sans selection:bg-blue-500 selection:text-white px-4 sm:px-6 py-6">
    
    <div class="absolute inset-0 overflow-hidden pointer-events-none">
      <div class="absolute top-0 right-1/4 w-64 h-64 md:w-96 md:h-96 bg-blue-600/20 rounded-full blur-[80px] md:blur-[100px] animate-blob"></div>
      <div class="absolute bottom-0 left-1/4 w-64 h-64 md:w-96 md:h-96 bg-purple-600/20 rounded-full blur-[80px] md:blur-[100px] animate-blob animation-delay-2000"></div>
      <div class="absolute inset-0 bg-[url('/pattern.svg')] opacity-5"></div>
    </div>

    <div 
      class="bg-slate-900/60 backdrop-blur-xl border border-slate-700/50 p-6 sm:p-8 md:p-10 rounded-2xl md:rounded-3xl shadow-2xl w-full max-w-md relative z-10"
      data-aos="fade-up"
    >
      
      <div class="text-center mb-6 md:mb-8">
        <div class="inline-flex items-center justify-center w-10 h-10 md:w-12 md:h-12 bg-gradient-to-br from-blue-600 to-indigo-600 rounded-xl shadow-lg shadow-blue-600/30 mb-3 md:mb-4 animate-bounce-slow">
          <Bird class="text-white w-5 h-5 md:w-6 md:h-6" />
        </div>
        <h2 class="text-xl md:text-3xl font-bold text-white tracking-tight">Buat Akun Farm</h2>
        <p class="text-slate-400 text-xs md:text-sm mt-1 md:mt-2">Mulai digitalisasi peternakan Anda.</p>
      </div>

      <form @submit.prevent="handleRegister" class="space-y-3 md:space-y-4">
        
        <div class="space-y-1">
          <label class="text-[10px] md:text-xs font-bold text-slate-300 uppercase tracking-wider ml-1">Username</label>
          <div class="relative group">
            <div class="absolute inset-y-0 left-0 pl-3 md:pl-4 flex items-center pointer-events-none">
              <User class="h-4 w-4 md:h-5 md:w-5 text-slate-500 group-focus-within:text-blue-500 transition-colors" />
            </div>
            <input 
              v-model="form.username" 
              type="text" 
              required 
              placeholder="Nama Pengguna"
              class="w-full pl-10 md:pl-11 pr-4 py-2.5 md:py-3 bg-slate-800/50 border border-slate-700 text-white text-sm md:text-base rounded-xl focus:ring-2 focus:ring-blue-500 focus:border-blue-500 block transition-all placeholder-slate-500"
            >
          </div>
        </div>

        <div class="space-y-1">
          <label class="text-[10px] md:text-xs font-bold text-slate-300 uppercase tracking-wider ml-1">Email</label>
          <div class="relative group">
            <div class="absolute inset-y-0 left-0 pl-3 md:pl-4 flex items-center pointer-events-none">
              <Mail class="h-4 w-4 md:h-5 md:w-5 text-slate-500 group-focus-within:text-blue-500 transition-colors" />
            </div>
            <input 
              v-model="form.email" 
              type="email" 
              required 
              placeholder="email@contoh.com"
              class="w-full pl-10 md:pl-11 pr-4 py-2.5 md:py-3 bg-slate-800/50 border border-slate-700 text-white text-sm md:text-base rounded-xl focus:ring-2 focus:ring-blue-500 focus:border-blue-500 block transition-all placeholder-slate-500"
            >
          </div>
        </div>

        <div class="space-y-1">
          <label class="text-[10px] md:text-xs font-bold text-slate-300 uppercase tracking-wider ml-1">Nama Farm</label>
          <div class="relative group">
            <div class="absolute inset-y-0 left-0 pl-3 md:pl-4 flex items-center pointer-events-none">
              <Home class="h-4 w-4 md:h-5 md:w-5 text-slate-500 group-focus-within:text-blue-500 transition-colors" />
            </div>
            <input 
              v-model="form.farm_name" 
              type="text" 
              required 
              placeholder="Contoh: Juara BF"
              class="w-full pl-10 md:pl-11 pr-4 py-2.5 md:py-3 bg-slate-800/50 border border-slate-700 text-white text-sm md:text-base rounded-xl focus:ring-2 focus:ring-blue-500 focus:border-blue-500 block transition-all placeholder-slate-500"
            >
          </div>
        </div>

        <div class="space-y-1">
          <label class="text-[10px] md:text-xs font-bold text-slate-300 uppercase tracking-wider ml-1">Password</label>
          <div class="relative group">
            <div class="absolute inset-y-0 left-0 pl-3 md:pl-4 flex items-center pointer-events-none">
              <Lock class="h-4 w-4 md:h-5 md:w-5 text-slate-500 group-focus-within:text-blue-500 transition-colors" />
            </div>
            <input 
              v-model="form.password" 
              type="password" 
              required 
              minlength="6"
              placeholder="••••••••"
              class="w-full pl-10 md:pl-11 pr-4 py-2.5 md:py-3 bg-slate-800/50 border border-slate-700 text-white text-sm md:text-base rounded-xl focus:ring-2 focus:ring-blue-500 focus:border-blue-500 block transition-all placeholder-slate-500"
            >
          </div>
          <p class="text-[10px] text-slate-500 ml-1 hidden md:block">*Minimal 6 karakter</p>
        </div>

        <button 
          type="submit" 
          :disabled="loading" 
          class="w-full flex items-center justify-center gap-2 bg-gradient-to-r from-blue-600 to-indigo-600 hover:from-blue-500 hover:to-indigo-500 text-white p-3 md:p-3.5 rounded-xl transition-all transform hover:-translate-y-0.5 shadow-lg shadow-blue-600/30 disabled:opacity-70 disabled:cursor-not-allowed mt-4 md:mt-6 font-bold text-sm md:text-base"
        >
          <Loader2 v-if="loading" class="animate-spin w-4 h-4 md:w-5 md:h-5" />
          <span v-else>Daftar Sekarang 🚀</span>
        </button>

      </form>

      <div class="mt-6 md:mt-8 pt-4 md:pt-6 border-t border-slate-700/50 text-center">
        <p class="text-slate-400 text-xs md:text-sm">
          Sudah punya akun? 
          <RouterLink to="/login" class="text-blue-400 font-bold hover:text-blue-300 transition inline-flex items-center gap-1 group p-2 md:p-0">
            <LogIn class="w-3 h-3 md:w-4 md:h-4" />
            Login di sini
          </RouterLink>
        </p>
      </div>

    </div>
  </div>
</template>

<style scoped>
/* Animasi Background Blob */
@keyframes blob {
  0% { transform: translate(0px, 0px) scale(1); }
  33% { transform: translate(30px, -50px) scale(1.1); }
  66% { transform: translate(-20px, 20px) scale(0.9); }
  100% { transform: translate(0px, 0px) scale(1); }
}
.animate-blob {
  animation: blob 7s infinite;
}
.animation-delay-2000 {
  animation-delay: 2s;
}

/* Animasi Bounce Halus untuk Logo */
@keyframes bounce-slow {
  0%, 100% { transform: translateY(0); }
  50% { transform: translateY(-5px); }
}
.animate-bounce-slow {
  animation: bounce-slow 3s infinite;
}
</style>