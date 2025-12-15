/** @type {import('tailwindcss').Config} */
export default {
  darkMode: 'class', // <--- PASTIIN BARIS INI ADA! JANGAN SAMPAI HILANG
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}