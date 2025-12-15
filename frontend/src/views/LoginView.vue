<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const form = ref({ email: '', password: '' })
const loading = ref(false)

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
      // SIMPAN DATA USER DI BROWSER (LocalStorage)
      // Supaya kalau di-refresh tidak log out
      localStorage.setItem('user_id', result.user_id)
      localStorage.setItem('username', result.username)
      localStorage.setItem('farm_name', result.farm_name)

      alert("✅ Login Sukses! Selamat datang " + result.username)
      
      // Arahkan ke Dashboard (Nanti kita buat)
      router.push('/dashboard') 
    } else {
      alert("❌ Gagal: " + result.error)
    }

  } catch (error) {
    console.error(error)
    alert("Server error!")
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-100">
    <div class="bg-white p-8 rounded-lg shadow-lg w-full max-w-md border-t-4 border-green-600">
      
      <h2 class="text-2xl font-bold text-center text-gray-800 mb-6">Masuk ke Kandang</h2>

      <form @submit.prevent="handleLogin" class="space-y-4">
        <div>
          <label class="block text-gray-700 font-bold mb-1">Email</label>
          <input v-model="form.email" type="email" required class="w-full p-3 border rounded focus:ring-2 focus:ring-green-500 outline-none">
        </div>

        <div>
          <label class="block text-gray-700 font-bold mb-1">Password</label>
          <input v-model="form.password" type="password" required class="w-full p-3 border rounded focus:ring-2 focus:ring-green-500 outline-none">
        </div>

        <button type="submit" :disabled="loading" class="w-full bg-green-600 text-white p-3 rounded hover:bg-green-700 transition font-bold text-lg">
          {{ loading ? 'Mengecek...' : 'MASUK' }}
        </button>
      </form>

      <p class="text-center mt-4 text-gray-600">
        Belum punya akun? 
        <RouterLink to="/register" class="text-green-600 font-bold hover:underline">Daftar dulu</RouterLink>
      </p>

    </div>
  </div>
</template>