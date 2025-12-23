<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
// Import Icon dari Lucide
import { Mail, Lock, LogIn, ArrowRight, Loader2, Bird } from 'lucide-vue-next'
// Import Animasi AOS
import AOS from 'aos'
import 'aos/dist/aos.css'

const router = useRouter()
const form = ref({ email: '', password: '' })
const loading = ref(false)

onMounted(() => {
  AOS.init({ duration: 800, once: true })
})

const handleLogin = async () => {
  loading.value = true
  try {
    const response = await fetch('http://localhost:8080/api/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(form.value)
    })

    const result = await response.json()

    if (response.ok) {
      localStorage.setItem('user_id', result.user_id)
      localStorage.setItem('username', result.username)
      localStorage.setItem('farm_name', result.farm_name)
      
      // Redirect ke Dashboard
      router.push('/dashboard') 
    } else {
      alert("❌ Gagal: " + result.error)
    }

  } catch (error) {
    console.error(error)
    alert("Server error! Pastikan backend jalan.")
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-[#0f172a] relative overflow-hidden font-sans selection:bg-blue-500 selection:text-white px-4">
    
    <div class="absolute inset-0 overflow-hidden pointer-events-none">
      <div class="absolute top-0 left-1/4 w-96 h-96 bg-blue-600/20 rounded-full blur-[100px] animate-blob"></div>
      <div class="absolute bottom-0 right-1/4 w-96 h-96 bg-purple-600/20 rounded-full blur-[100px] animate-blob animation-delay-2000"></div>
      <div class="absolute inset-0 bg-[url('/pattern.svg')] opacity-5"></div>
    </div>

    <div 
      class="bg-slate-900/50 backdrop-blur-xl border border-slate-700/50 p-8 md:p-10 rounded-3xl shadow-2xl w-full max-w-md relative z-10"
      data-aos="zoom-in"
    >
      
      <div class="text-center mb-8">
        <div class="inline-flex items-center justify-center w-12 h-12 bg-blue-600 rounded-xl shadow-lg shadow-blue-600/30 mb-4 animate-bounce-slow">
          <Bird class="text-white w-6 h-6" />
        </div>
        <h2 class="text-2xl md:text-3xl font-bold text-white tracking-tight">Selamat Datang</h2>
        <p class="text-slate-400 text-sm mt-2">Masuk untuk mengelola peternakan Anda.</p>
      </div>

      <form @submit.prevent="handleLogin" class="space-y-5">
        
        <div class="space-y-1.5">
          <label class="text-xs font-bold text-slate-300 uppercase tracking-wider ml-1">Email</label>
          <div class="relative group">
            <div class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none">
              <Mail class="h-5 w-5 text-slate-500 group-focus-within:text-blue-500 transition-colors" />
            </div>
            <input 
              v-model="form.email" 
              type="email" 
              required 
              placeholder="nama@email.com"
              class="w-full pl-11 pr-4 py-3.5 bg-slate-800/50 border border-slate-700 text-white rounded-xl focus:ring-2 focus:ring-blue-500 focus:border-blue-500 block transition-all placeholder-slate-500"
            >
          </div>
        </div>

        <div class="space-y-1.5">
          <div class="flex justify-between items-center ml-1">
            <label class="text-xs font-bold text-slate-300 uppercase tracking-wider">Password</label>
            <a href="#" class="text-xs text-blue-400 hover:text-blue-300 transition">Lupa password?</a>
          </div>
          <div class="relative group">
            <div class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none">
              <Lock class="h-5 w-5 text-slate-500 group-focus-within:text-blue-500 transition-colors" />
            </div>
            <input 
              v-model="form.password" 
              type="password" 
              required 
              placeholder="••••••••"
              class="w-full pl-11 pr-4 py-3.5 bg-slate-800/50 border border-slate-700 text-white rounded-xl focus:ring-2 focus:ring-blue-500 focus:border-blue-500 block transition-all placeholder-slate-500"
            >
          </div>
        </div>

        <button 
          type="submit" 
          :disabled="loading" 
          class="w-full flex items-center justify-center gap-2 bg-gradient-to-r from-blue-600 to-indigo-600 hover:from-blue-500 hover:to-indigo-500 text-white p-3.5 rounded-xl transition-all transform hover:-translate-y-0.5 shadow-lg shadow-blue-600/30 disabled:opacity-70 disabled:cursor-not-allowed mt-2 font-bold"
        >
          <Loader2 v-if="loading" class="animate-spin w-5 h-5" />
          <LogIn v-else class="w-5 h-5" />
          <span>{{ loading ? 'Memproses...' : 'Masuk Sekarang' }}</span>
        </button>

      </form>

      <div class="mt-8 pt-6 border-t border-slate-700/50 text-center">
        <p class="text-slate-400 text-sm">
          Belum punya akun? 
          <RouterLink to="/register" class="text-blue-400 font-bold hover:text-blue-300 transition inline-flex items-center gap-1 group">
            Daftar Gratis
            <ArrowRight class="w-3 h-3 group-hover:translate-x-1 transition-transform" />
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