import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { createI18n } from 'vue-i18n' // 1. Import i18n

import App from './App.vue'
import router from './router'

// 2. Import Kamus Bahasa (Pastikan file id.json & en.json sudah dibuat di folder locales)
import id from './locales/id.json'
import en from './locales/en.json'

const app = createApp(App)

// 3. Konfigurasi i18n
const i18n = createI18n({
  legacy: false, // Wajib false agar bisa dipakai di script setup
  locale: localStorage.getItem('language') || 'id', // Ambil bahasa dari memori
  fallbackLocale: 'id',
  globalInjection: true,
  messages: {
    id,
    en
  }
})

app.use(createPinia())
app.use(router)
app.use(i18n) // 4. Pasang i18n ke aplikasi (JANGAN LUPA INI)

app.mount('#app')