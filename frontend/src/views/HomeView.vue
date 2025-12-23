<script setup>
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
// Import Icon Lucide yang dibutuhkan untuk mockup dan fitur
import { 
  LayoutDashboard, 
  Bird, 
  HeartHandshake, 
  Dna, 
  Wallet, 
  ShieldCheck,
  Facebook,
  Instagram,
  Mail,

  Menu, 
  X,
  TrendingUp // Tambahan icon untuk mockup grafik
} from 'lucide-vue-next'

// Import Library Animasi AOS
import AOS from 'aos'
import 'aos/dist/aos.css'

const router = useRouter()

// Inisialisasi Animasi
onMounted(() => {
  AOS.init({
    duration: 800,
    once: true,
    offset: 50,
    easing: 'ease-out-cubic',
  })
})
</script>

<template>
  <div class="min-h-screen bg-slate-50 font-sans text-slate-800 overflow-x-hidden selection:bg-blue-200 selection:text-blue-900">
    
    <nav class="fixed w-full top-0 z-50 bg-white/80 backdrop-blur-lg border-b border-slate-200/50 transition-all duration-300">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center h-16">
          
          <div class="flex items-center gap-2 cursor-pointer" @click="router.push('/')">
            <div class="bg-gradient-to-br from-blue-600 to-blue-700 text-white p-2 rounded-xl shadow-lg shadow-blue-500/20 transform active:scale-95 transition">
              <Bird :size="20" stroke-width="2.5" />
            </div>
            <h1 class="text-lg md:text-xl font-bold text-slate-800 tracking-tight">Volar Farm</h1>
          </div>

          <div class="hidden md:flex items-center gap-6 font-medium text-slate-500 text-sm">
            <a href="#fitur" class="hover:text-blue-600 transition hover:-translate-y-0.5 inline-block">Fitur</a>
            <RouterLink to="/login" class="px-5 py-2 text-slate-600 font-bold hover:text-blue-600 transition">Masuk</RouterLink>
            <RouterLink to="/register" class="px-5 py-2.5 bg-blue-600 text-white font-bold rounded-full hover:bg-blue-700 shadow-md shadow-blue-500/30 transition transform hover:scale-105 active:scale-95">
              Daftar Gratis
            </RouterLink>
          </div>

          <div class="flex md:hidden gap-2">
            <RouterLink to="/login" class="px-4 py-2 text-sm font-bold text-slate-600 hover:text-blue-600">
              Masuk
            </RouterLink>
            <RouterLink to="/register" class="px-4 py-2 bg-blue-600 text-white text-sm font-bold rounded-full shadow-sm active:scale-95 transition">
              Daftar
            </RouterLink>
          </div>
        </div>
      </div>
    </nav>

<header class="relative bg-[#0f172a] text-white overflow-hidden px-4 sm:px-6 lg:px-8">
      
      <div class="absolute top-0 left-0 w-full h-full opacity-30 pointer-events-none overflow-hidden z-0">
         <div class="absolute right-0 top-0 w-[300px] h-[300px] md:w-[500px] md:h-[500px] bg-blue-600 rounded-full blur-[100px] mix-blend-multiply animate-blob"></div>
         <div class="absolute -left-20 bottom-0 w-[250px] h-[250px] md:w-[400px] md:h-[400px] bg-purple-600 rounded-full blur-[100px] mix-blend-multiply animate-blob animation-delay-2000"></div>
      </div>

      <div class="max-w-7xl mx-auto pt-32 pb-20 md:pt-40 md:pb-32 relative z-10">
        <div class="flex flex-col lg:flex-row items-center justify-between gap-12 lg:gap-24">
          
          <div class="flex-1 text-center lg:text-left space-y-8 order-2 lg:order-1 relative z-20" data-aos="fade-right" data-aos-duration="1000">
            <span class="inline-flex items-center gap-1.5 py-1.5 px-3 rounded-full bg-blue-900/40 border border-blue-500/30 text-blue-300 text-[10px] md:text-xs font-bold tracking-wide backdrop-blur-sm">
              <span class="relative flex h-2 w-2">
                <span class="animate-ping absolute inline-flex h-full w-full rounded-full bg-blue-400 opacity-75"></span>
                <span class="relative inline-flex rounded-full h-2 w-2 bg-blue-500"></span>
              </span>
              MANAJEMEN TERNAK MODERN
            </span>
            
            <h2 class="text-4xl sm:text-5xl md:text-6xl font-extrabold leading-[1.1] tracking-tight">
              Kelola Farm <br>
              <span class="text-transparent bg-clip-text bg-gradient-to-r from-blue-400 via-blue-200 to-purple-400">Lebih Cerdas.</span>
            </h2>
            
            <p class="text-base sm:text-lg text-slate-300 max-w-xl mx-auto lg:mx-0 leading-relaxed">
              Tinggalkan pencatatan manual. Platform all-in-one untuk silsilah, manajemen stok, dan analisis keuangan peternakan Anda.
            </p>
            
            <div class="flex flex-col sm:flex-row gap-4 justify-center lg:justify-start relative z-30">
              <RouterLink to="/register" class="px-8 py-4 bg-blue-600 text-white font-bold rounded-full shadow-xl shadow-blue-500/30 hover:bg-blue-500 hover:shadow-blue-500/50 transition transform hover:-translate-y-1 text-center cursor-pointer select-none">
                Mulai Gratis Sekarang 🚀
              </RouterLink>
              <a href="#fitur" class="px-8 py-4 bg-slate-800/50 text-slate-300 font-bold rounded-full border border-slate-700 hover:bg-slate-800 hover:text-white transition backdrop-blur-sm text-center cursor-pointer select-none">
                Pelajari Fitur
              </a>
            </div>
          </div>

          <div class="flex-1 w-full max-w-lg lg:max-w-none relative order-1 lg:order-2 z-10" data-aos="fade-left" data-aos-delay="200" data-aos-duration="1000">
            <div class="relative perspective-1000 group">
              <div class="relative transform lg:rotate-y-[-12deg] lg:rotate-x-[6deg] group-hover:rotate-y-0 group-hover:rotate-x-0 transition-all duration-700 ease-out-cubic shadow-2xl shadow-blue-900/30 rounded-2xl border-4 border-slate-800 bg-[#0f172a] overflow-hidden z-20">
                
                <div class="bg-slate-900 px-4 py-3 flex items-center gap-2 border-b border-slate-800/50">
                  <div class="flex gap-1.5">
                    <div class="w-2.5 h-2.5 rounded-full bg-red-500/80"></div>
                    <div class="w-2.5 h-2.5 rounded-full bg-yellow-500/80"></div>
                    <div class="w-2.5 h-2.5 rounded-full bg-green-500/80"></div>
                  </div>
                  <div class="flex-1 text-center text-[10px] text-slate-500 font-mono bg-slate-800/50 rounded-full py-0.5 mx-4 truncate">volarfarm.com/dashboard</div>
                </div>
                
                <div class="p-4 md:p-5 bg-gradient-to-br from-[#0f172a] to-[#1e293b] relative">
                  <div class="absolute inset-0 bg-gradient-to-tr from-white/5 to-transparent pointer-events-none"></div>
                  
                  <div class="grid grid-cols-3 gap-3 mb-5">
                    <div class="bg-blue-900/20 p-3 rounded-xl border border-blue-500/10 backdrop-blur-sm">
                      <Bird class="w-4 h-4 text-blue-400 mb-1.5" />
                      <div class="text-lg font-bold">145</div>
                      <div class="text-[9px] text-slate-400">Total Ekor</div>
                    </div>
                    <div class="bg-green-900/20 p-3 rounded-xl border border-green-500/10 backdrop-blur-sm">
                      <Wallet class="w-4 h-4 text-green-400 mb-1.5" />
                      <div class="text-lg font-bold">8.5Jt</div>
                      <div class="text-[9px] text-slate-400">Profit</div>
                    </div>
                    <div class="bg-pink-900/20 p-3 rounded-xl border border-pink-500/10 backdrop-blur-sm">
                      <HeartHandshake class="w-4 h-4 text-pink-400 mb-1.5" />
                      <div class="text-lg font-bold">12 Psng</div>
                      <div class="text-[9px] text-slate-400">Produksi</div>
                    </div>
                  </div>

                  <div class="bg-slate-800/40 p-4 rounded-xl border border-slate-700/30 h-36 flex flex-col justify-between relative overflow-hidden backdrop-blur-sm">
                    <div class="flex justify-between items-center mb-2">
                      <span class="text-[10px] text-slate-400 font-medium flex items-center gap-1">
                        <TrendingUp class="w-3 h-3" /> Grafik
                      </span>
                      <span class="text-[9px] text-green-400 bg-green-900/30 px-1.5 py-0.5 rounded-full">+12%</span>
                    </div>
                    <div class="relative flex-1 w-full h-full bg-slate-700/20 rounded-lg animate-pulse"></div>
                  </div>
                </div>
              </div>

              <div class="absolute -z-10 top-1/2 right-1/2 translate-x-1/2 -translate-y-1/2 w-[140%] h-[140%] bg-gradient-to-br from-blue-600/10 to-purple-600/10 blur-3xl rounded-full animate-pulse-slow opacity-70 pointer-events-none"></div>
            </div>
          </div>

        </div>
      </div>
    </header>

    <section id="fitur" class="py-20 bg-white relative z-10">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        
        <div class="text-center mb-16" data-aos="fade-up">
          <h3 class="text-2xl md:text-3xl font-bold text-slate-900 mb-3">Fitur Lengkap & Canggih</h3>
          <p class="text-slate-500 max-w-xl mx-auto text-sm md:text-base">Didesain khusus untuk memudahkan alur kerja breeder profesional.</p>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 md:gap-8">
          
          <div class="group p-6 bg-white rounded-2xl border border-slate-100 shadow-[0_4px_20px_-10px_rgba(0,0,0,0.05)] hover:shadow-[0_20px_40px_-15px_rgba(0,0,0,0.1)] hover:border-blue-100 transition duration-300 h-full" data-aos="fade-up" data-aos-delay="0">
            <div class="w-12 h-12 bg-blue-50 text-blue-600 rounded-xl flex items-center justify-center mb-4 group-hover:scale-110 group-hover:bg-blue-600 group-hover:text-white transition duration-300">
              <Bird :size="24" stroke-width="2" />
            </div>
            <h4 class="text-lg font-bold text-slate-800 mb-2">Manajemen Stok</h4>
            <p class="text-slate-500 text-sm leading-relaxed">Data visual, ring, dan gender tercatat rapi. Filter pencarian cepat dalam hitungan detik.</p>
          </div>

          <div class="group p-6 bg-white rounded-2xl border border-slate-100 shadow-[0_4px_20px_-10px_rgba(0,0,0,0.05)] hover:shadow-[0_20px_40px_-15px_rgba(0,0,0,0.1)] hover:border-pink-100 transition duration-300 h-full" data-aos="fade-up" data-aos-delay="100">
            <div class="w-12 h-12 bg-pink-50 text-pink-600 rounded-xl flex items-center justify-center mb-4 group-hover:scale-110 group-hover:bg-pink-600 group-hover:text-white transition duration-300">
              <HeartHandshake :size="24" stroke-width="2" />
            </div>
            <h4 class="text-lg font-bold text-slate-800 mb-2">Penjodohan</h4>
            <p class="text-slate-500 text-sm leading-relaxed">Kelola pasangan di kandang ternak. Pantau tanggal masuk dan riwayat produksi mereka.</p>
          </div>

          <div class="group p-6 bg-white rounded-2xl border border-slate-100 shadow-[0_4px_20px_-10px_rgba(0,0,0,0.05)] hover:shadow-[0_20px_40px_-15px_rgba(0,0,0,0.1)] hover:border-purple-100 transition duration-300 h-full" data-aos="fade-up" data-aos-delay="200">
            <div class="w-12 h-12 bg-purple-50 text-purple-600 rounded-xl flex items-center justify-center mb-4 group-hover:scale-110 group-hover:bg-purple-600 group-hover:text-white transition duration-300">
              <Dna :size="24" stroke-width="2" />
            </div>
            <h4 class="text-lg font-bold text-slate-800 mb-2">Silsilah Otomatis</h4>
            <p class="text-slate-500 text-sm leading-relaxed">Sistem otomatis membuat pohon keluarga (pedigree) dari data induk. Cetak sertifikat lebih mudah.</p>
          </div>

          <div class="group p-6 bg-white rounded-2xl border border-slate-100 shadow-[0_4px_20px_-10px_rgba(0,0,0,0.05)] hover:shadow-[0_20px_40px_-15px_rgba(0,0,0,0.1)] hover:border-green-100 transition duration-300 h-full" data-aos="fade-up" data-aos-delay="300">
            <div class="w-12 h-12 bg-green-50 text-green-600 rounded-xl flex items-center justify-center mb-4 group-hover:scale-110 group-hover:bg-green-600 group-hover:text-white transition duration-300">
              <Wallet :size="24" stroke-width="2" />
            </div>
            <h4 class="text-lg font-bold text-slate-800 mb-2">Keuangan</h4>
            <p class="text-slate-500 text-sm leading-relaxed">Pantau arus kas masuk (penjualan) dan keluar (pakan/obat). Analisa profit bulanan Anda.</p>
          </div>

          <div class="group p-6 bg-white rounded-2xl border border-slate-100 shadow-[0_4px_20px_-10px_rgba(0,0,0,0.05)] hover:shadow-[0_20px_40px_-15px_rgba(0,0,0,0.1)] hover:border-orange-100 transition duration-300 h-full" data-aos="fade-up" data-aos-delay="400">
            <div class="w-12 h-12 bg-orange-50 text-orange-600 rounded-xl flex items-center justify-center mb-4 group-hover:scale-110 group-hover:bg-orange-600 group-hover:text-white transition duration-300">
              <LayoutDashboard :size="24" stroke-width="2" />
            </div>
            <h4 class="text-lg font-bold text-slate-800 mb-2">Dashboard Pintar</h4>
            <p class="text-slate-500 text-sm leading-relaxed">Ringkasan total burung, pasangan aktif, dan grafik performa farm dalam satu layar.</p>
          </div>

          <div class="group p-6 bg-white rounded-2xl border border-slate-100 shadow-[0_4px_20px_-10px_rgba(0,0,0,0.05)] hover:shadow-[0_20px_40px_-15px_rgba(0,0,0,0.1)] hover:border-slate-300 transition duration-300 h-full" data-aos="fade-up" data-aos-delay="500">
            <div class="w-12 h-12 bg-slate-100 text-slate-600 rounded-xl flex items-center justify-center mb-4 group-hover:scale-110 group-hover:bg-slate-800 group-hover:text-white transition duration-300">
              <ShieldCheck :size="24" stroke-width="2" />
            </div>
            <h4 class="text-lg font-bold text-slate-800 mb-2">Keamanan Cloud</h4>
            <p class="text-slate-500 text-sm leading-relaxed">Data tersimpan di server cloud yang aman. Tidak perlu takut data hilang jika HP rusak.</p>
          </div>

        </div>
      </div>
    </section>

    <footer class="bg-slate-950 text-slate-400 pt-20 pb-10 border-t border-slate-900 font-sans">
      <div class="max-w-7xl mx-auto px-6">
        
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-12 mb-16">
          
          <div class="space-y-6">
            <div class="flex items-center gap-2 text-white">
              <div class="bg-blue-600 p-1.5 rounded-lg">
                <Bird :size="24" class="text-white" />
              </div>
              <span class="text-xl font-bold tracking-tight">Volar Farm</span>
            </div>
            <p class="text-sm leading-relaxed text-slate-500">
              Platform manajemen peternakan burung #1 di Indonesia. Membantu breeder mencatat silsilah, stok, dan profit dengan mudah.
            </p>
            <div class="flex gap-4">
              <a href="#" class="w-10 h-10 rounded-full bg-slate-900 flex items-center justify-center hover:bg-blue-600 hover:text-white transition duration-300"><Facebook :size="18" /></a>
              <a href="#" class="w-10 h-10 rounded-full bg-slate-900 flex items-center justify-center hover:bg-blue-400 hover:text-white transition duration-300"><Twitter :size="18" /></a>
              <a href="#" class="w-10 h-10 rounded-full bg-slate-900 flex items-center justify-center hover:bg-pink-600 hover:text-white transition duration-300"><Instagram :size="18" /></a>
              <a href="#" class="w-10 h-10 rounded-full bg-slate-900 flex items-center justify-center hover:bg-blue-700 hover:text-white transition duration-300"><Linkedin :size="18" /></a>
            </div>
          </div>

          <div>
            <h4 class="text-white font-bold mb-6">Produk</h4>
            <ul class="space-y-4 text-sm">
              <li><a href="#" class="hover:text-blue-500 transition">Fitur Utama</a></li>
              <li><a href="#" class="hover:text-blue-500 transition">Harga & Paket</a></li>
              <li><a href="#" class="hover:text-blue-500 transition">Roadmap Update</a></li>
              <li><a href="#" class="hover:text-blue-500 transition">Showcase Farm</a></li>
            </ul>
          </div>

          <div>
            <h4 class="text-white font-bold mb-6">Dukungan</h4>
            <ul class="space-y-4 text-sm">
              <li><a href="#" class="hover:text-blue-500 transition">Pusat Bantuan</a></li>
              <li><a href="#" class="hover:text-blue-500 transition">Dokumentasi API</a></li>
              <li><a href="#" class="hover:text-blue-500 transition">Syarat & Ketentuan</a></li>
              <li><a href="#" class="hover:text-blue-500 transition">Kebijakan Privasi</a></li>
            </ul>
          </div>

          <div>
            <h4 class="text-white font-bold mb-6">Berlangganan Info</h4>
            <p class="text-sm text-slate-500 mb-4">Dapatkan tips breeding dan info update terbaru.</p>
            <form class="flex flex-col gap-3">
              <div class="relative">
                <Mail :size="16" class="absolute left-3 top-3 text-slate-500" />
                <input type="email" placeholder="Email Anda" class="w-full bg-slate-900 border border-slate-800 rounded-lg py-2.5 pl-10 pr-4 text-sm focus:outline-none focus:border-blue-600 focus:ring-1 focus:ring-blue-600 transition">
              </div>
              <button type="button" class="w-full bg-blue-600 text-white font-bold py-2.5 rounded-lg hover:bg-blue-500 transition text-sm">
                Langganan
              </button>
            </form>
          </div>

        </div>

        <div class="border-t border-slate-900 pt-8 flex flex-col md:flex-row justify-between items-center gap-4 text-xs text-slate-600">
          <p>&copy; 2025 Volar Farm Indonesia. Hak Cipta Dilindungi.</p>
          <div class="flex gap-6">
            <a href="#" class="hover:text-slate-400">Privacy</a>
            <a href="#" class="hover:text-slate-400">Terms</a>
            <a href="#" class="hover:text-slate-400">Sitemap</a>
          </div>
        </div>

      </div>
    </footer>

  </div>
</template>

<style scoped>
/* Custom Animation untuk Background Blob */
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

/* CSS untuk Efek 3D Perspektif pada Mockup */
.perspective-1000 {
  perspective: 1000px;
}
/* Rotasi hanya pada layar besar (lg) */
@media (min-width: 1024px) {
  .lg\:rotate-y-\[-12deg\] {
    transform: rotateY(-12deg);
  }
  .lg\:rotate-x-\[6deg\] {
    transform: rotateX(6deg);
  }
}
/* Reset rotasi saat hover */
.group:hover .group-hover\:rotate-y-0 {
  transform: rotateY(0deg);
}
.group:hover .group-hover\:rotate-x-0 {
  transform: rotateX(0deg);
}

.ease-out-cubic {
  transition-timing-function: cubic-bezier(0.33, 1, 0.68, 1);
}

/* Animasi Pulse Lambat untuk Cahaya di Belakang Mockup */
.animate-pulse-slow {
  animation: pulse-slow 4s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}
@keyframes pulse-slow {
  0%, 100% { opacity: 0.5; transform: translate(-50%, -50%) scale(1); }
  50% { opacity: 0.8; transform: translate(-50%, -50%) scale(1.05); }
}
</style>