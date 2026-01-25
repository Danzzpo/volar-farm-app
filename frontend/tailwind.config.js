/** @type {import('tailwindcss').Config} */
import scrollbar from 'tailwind-scrollbar'; // <--- Import plugin

export default {
  darkMode: 'class',
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {},
  },
  plugins: [
    scrollbar, // <--- Masukkan ke sini
  ],
}