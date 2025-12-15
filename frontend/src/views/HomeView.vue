<script setup>
import { ref, onMounted } from 'vue'

const animals = ref([])
const loading = ref(true)

// Ambil data burung yang AVAILABLE saja (Publik)
const fetchPublicAnimals = async () => {
  try {
    const res = await fetch('http://localhost:8080/api/public/animals')
    const json = await res.json()
    animals.value = json.data || []
  } catch (err) {
    console.error(err)
  } finally {
    loading.value = false
  }
}

// Format Tanggal
const formatDate = (d) => {
  if (!d || d.startsWith('0001')) return '-'
  return new Date(d).toLocaleDateString('id-ID', { day: 'numeric', month: 'short', year: 'numeric' })
}

// Link WhatsApp (Otomatis)
const contactSeller = (bird) => {
  const text = `Halo, saya tertarik dengan burung ${bird.visual} (Ring: ${bird.ring_number}). Apakah masih ada?`
  window.open(`https://wa.me/?text=${encodeURIComponent(text)}`, '_blank')
}

onMounted(() => {
  fetchPublicAnimals()
})
</script>

<template>
  <div class="min-h-screen bg-slate-50 font-sans">
    
    <nav class="bg-white/90 backdrop-blur-md shadow-sm sticky top-0 z-50 transition-all duration-300">
      <div class="max-w-6xl mx-auto px-4 py-4 flex justify-between items-center">
        <div class="flex items-center gap-2 cursor-pointer" @click="$router.push('/')">
          <div class="bg-green-600 text-white p-2 rounded-lg font-bold text-xl">VF</div>
          <h1 class="text-xl font-bold text-slate-800 tracking-tight">Volar Farm</h1>
        </div>

        <div class="hidden md:flex items-center gap-6 font-medium text-slate-600">
          <a href="#katalog" class="hover:text-green-600 transition">Katalog</a>
          <a href="#tentang" class="hover:text-green-600 transition">Tentang</a>
          <RouterLink to="/login" class="px-5 py-2 text-green-600 border border-green-600 rounded-full hover:bg-green-50 transition">Masuk</RouterLink>
          <RouterLink to="/register" class="px-5 py-2 bg-green-600 text-white rounded-full hover:bg-green-700 shadow-md transition transform hover:scale-105">Daftar</RouterLink>
        </div>
      </div>
    </nav>

    <header class="relative bg-green-700 text-white py-24 px-4 overflow-hidden">
      <div class="absolute top-0 left-0 w-full h-full opacity-10 bg-[url('https://www.transparenttextures.com/patterns/cubes.png')]"></div>
      
      <div class="max-w-4xl mx-auto text-center relative z-10">
        <h2 class="text-4xl md:text-6xl font-extrabold mb-6 leading-tight">
          Kelola Ternak Burung <br> Lebih Profesional
        </h2>
        <p class="text-lg md:text-xl text-green-100 mb-8 max-w-2xl mx-auto">
          Sistem manajemen breeding modern untuk mencatat silsilah, stok, dan memantau perkembangan farm Anda dengan mudah.
        </p>
        <div class="flex flex-col md:flex-row gap-4 justify-center">
          <a href="#katalog" class="px-8 py-3 bg-white text-green-700 font-bold rounded-full shadow-lg hover:bg-gray-100 transition">
            Lihat Stok Burung
          </a>
          <RouterLink to="/register" class="px-8 py-3 bg-green-800 text-white font-bold rounded-full shadow-lg hover:bg-green-900 transition border border-green-600">
            Mulai Ternak Sekarang
          </RouterLink>
        </div>
      </div>
    </header>

    <section id="katalog" class="py-20 px-4 max-w-6xl mx-auto">
      <div class="text-center mb-12">
        <h3 class="text-3xl font-bold text-slate-800">Burung Siap Adopsi</h3>
        <p class="text-slate-500 mt-2">Daftar burung berkualitas yang tersedia untuk Anda pinang.</p>
      </div>

      <div v-if="loading" class="text-center py-10 text-gray-400">
        <p>Memuat data...</p>
      </div>

      <div v-else-if="animals.length === 0" class="text-center py-10 bg-white rounded-xl shadow-sm border border-dashed border-gray-300">
        <p class="text-gray-500 text-lg">Belum ada burung yang Available saat ini.</p>
      </div>

      <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
        <div v-for="bird in animals" :key="bird.id" class="group bg-white rounded-2xl shadow-sm hover:shadow-xl transition duration-300 overflow-hidden border border-slate-100 flex flex-col">
          
          <div class="h-3 bg-green-500"></div>
          
          <div class="p-6">
            <div class="flex justify-between items-start mb-4">
              <div>
                <h4 class="text-xl font-bold text-slate-800 group-hover:text-green-600 transition">{{ bird.visual }}</h4>
                <span class="text-xs bg-slate-100 text-slate-500 px-2 py-1 rounded mt-1 inline-block font-mono">
                  {{ bird.ring_number }}
                </span>
              </div>
              <span class="text-2xl" :title="bird.gender === 'M' ? 'Jantan' : bird.gender === 'F' ? 'Betina' : 'Unknown'">
                {{ bird.gender === 'M' ? '💙' : bird.gender === 'F' ? '💗' : '❓' }}
              </span>
            </div>

            <div class="space-y-2 text-sm text-slate-600 bg-slate-50 p-3 rounded-lg mb-4">
              <div class="flex items-center gap-2">
                <span class="font-bold w-4 text-blue-500">♂</span>
                <span class="truncate">{{ bird.sire ? bird.sire.visual : (bird.sire_other || '-') }}</span>
              </div>
              <div class="flex items-center gap-2">
                <span class="font-bold w-4 text-pink-500">♀</span>
                <span class="truncate">{{ bird.dam ? bird.dam.visual : (bird.dam_other || '-') }}</span>
              </div>
            </div>

            <div class="flex items-center justify-between text-xs text-slate-400 mb-4">
              <span>Hatch: {{ formatDate(bird.hatch_date) }}</span>
              <span class="text-green-600 font-bold bg-green-50 px-2 py-1 rounded-full">AVAILABLE</span>
            </div>

            <button @click="contactSeller(bird)" class="w-full py-2 bg-green-600 text-white rounded-lg font-bold hover:bg-green-700 transition flex items-center justify-center gap-2">
              <span>💬</span> Hubungi Penjual
            </button>
          </div>
        </div>
      </div>
    </section>

    <footer class="bg-slate-900 text-slate-400 py-8 text-center text-sm">
      <p>&copy; 2025 Volar Farm. All rights reserved.</p>
    </footer>

  </div>
</template>