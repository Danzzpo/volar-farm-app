<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

// Data Form
const form = ref({
  username: '',
  email: '',
  password: '',
  farm_name: '' // Harus sama dengan JSON di Golang
})

const loading = ref(false)

// Fungsi Saat Tombol Daftar Diklik
const handleRegister = async () => {
  loading.value = true
  
  try {
    // Tembak ke Backend GO
    const response = await fetch('http://localhost:8080/api/register', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(form.value)
    })

    const result = await response.json()

    if (response.ok) {
      alert("✅ Berhasil Daftar! Silakan Login.")
      router.push('/login') // Nanti kita buat halaman login
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
  <div class="min-h-screen flex items-center justify-center bg-gray-100">
    <div class="bg-white p-8 rounded-lg shadow-lg w-full max-w-md">
      
      <h2 class="text-2xl font-bold text-center text-green-600 mb-6">Daftar Volar Farm</h2>

      <form @submit.prevent="handleRegister" class="space-y-4">
        
        <div>
          <label class="block text-gray-700">Username</label>
          <input v-model="form.username" type="text" required class="w-full p-2 border rounded focus:ring-2 focus:ring-green-500 outline-none">
        </div>

        <div>
          <label class="block text-gray-700">Email</label>
          <input v-model="form.email" type="email" required class="w-full p-2 border rounded focus:ring-2 focus:ring-green-500 outline-none">
        </div>

        <div>
          <label class="block text-gray-700">Nama Farm</label>
          <input v-model="form.farm_name" type="text" placeholder="Contoh: Danzz Bird Farm" required class="w-full p-2 border rounded focus:ring-2 focus:ring-green-500 outline-none">
        </div>

        <div>
          <label class="block text-gray-700">Password</label>
          <input v-model="form.password" type="password" required minlength="6" class="w-full p-2 border rounded focus:ring-2 focus:ring-green-500 outline-none">
        </div>

        <button type="submit" :disabled="loading" class="w-full bg-green-600 text-white p-2 rounded hover:bg-green-700 transition font-bold">
          {{ loading ? 'Memproses...' : 'Daftar Sekarang' }}
        </button>

      </form>

    </div>
  </div>
</template>