<script setup>
import { ref, onMounted, nextTick } from 'vue'
import { Search, ChevronDown, Check, Loader2 } from 'lucide-vue-next'

const props = defineProps({
  modelValue: [String, Number], 
  userId: String,
  gender: String, 
  label: String
})

const emit = defineEmits(['update:modelValue'])

const isOpen = ref(false)
const isLoading = ref(false)
const options = ref([])
const searchQuery = ref('')
const selectedLabel = ref('')
const searchInputRef = ref(null)

let debounceTimeout = null

const fetchBirds = async (search = '') => {
  isLoading.value = true
  try {
    let url = `http://localhost:8080/api/animals?user_id=${props.userId}&gender=${props.gender}&limit=50&search=${search}`
    const res = await fetch(url)
    const json = await res.json()
    
    const rawData = json.data || []
    options.value = rawData.filter(a => 
      ['AVAILABLE', 'KEEP', 'Available', 'Keep'].includes(a.status)
    )
  } catch (e) {
    console.error(e)
  } finally {
    isLoading.value = false
  }
}

const onSearch = () => {
  if (debounceTimeout) clearTimeout(debounceTimeout)
  debounceTimeout = setTimeout(() => {
    fetchBirds(searchQuery.value)
  }, 500)
}

const selectItem = (bird) => {
  selectedLabel.value = `${bird.ring_number} - ${bird.visual}`
  emit('update:modelValue', bird.id)
  isOpen.value = false
  searchQuery.value = '' 
}

const toggleDropdown = () => {
  isOpen.value = !isOpen.value
  if (isOpen.value) {
    if (options.value.length === 0) fetchBirds()
    nextTick(() => searchInputRef.value?.focus())
  }
}

onMounted(() => {
  if (props.modelValue) {
    fetchBirds().then(() => {
      const found = options.value.find(o => o.id === props.modelValue)
      if (found) selectedLabel.value = `${found.ring_number} - ${found.visual}`
    })
  }
})
</script>

<template>
  <div class="relative w-full">
    <label class="block text-sm font-bold text-slate-700 dark:text-slate-300 uppercase tracking-wide mb-2">{{ label }}</label>
    
    <button 
      type="button"
      @click="toggleDropdown"
      class="w-full p-3.5 bg-white dark:bg-slate-800 border border-slate-300 dark:border-slate-600 rounded-xl text-left text-base flex justify-between items-center focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition shadow-sm"
      :class="{'border-blue-500 ring-2 ring-blue-100': isOpen}"
    >
      <span class="truncate pr-2 font-semibold text-slate-800 dark:text-white">
        {{ selectedLabel || `-- Pilih ${label} --` }}
      </span>
      <ChevronDown class="w-5 h-5 text-slate-500 flex-shrink-0 transition-transform duration-300" :class="{'rotate-180': isOpen}"/>
    </button>

    <div v-if="isOpen" class="absolute z-50 mt-2 w-full bg-white dark:bg-slate-800 rounded-xl shadow-2xl border border-slate-200 dark:border-slate-600 overflow-hidden animate-in fade-in zoom-in-95 duration-200 origin-top-left">
      
      <div class="p-3 border-b border-slate-100 dark:border-slate-700 bg-slate-50 dark:bg-slate-900/50">
        <div class="relative">
          <div class="absolute left-3 top-1/2 -translate-y-1/2 pointer-events-none">
             <Search class="w-4 h-4 text-slate-400"/>
          </div>
          <input 
            v-model="searchQuery"
            @input="onSearch"
            ref="searchInputRef"
            type="text" 
            placeholder="Cari Ring / Warna..." 
            class="w-full pl-10 pr-9 py-2.5 bg-white dark:bg-slate-700 border border-slate-200 dark:border-slate-600 rounded-lg text-base outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 transition placeholder:text-slate-400 text-slate-700 dark:text-white"
          >
          <div v-if="isLoading" class="absolute right-3 top-1/2 -translate-y-1/2">
            <Loader2 class="w-4 h-4 text-blue-500 animate-spin"/>
          </div>
        </div>
      </div>

      <ul class="max-h-60 overflow-y-auto custom-scrollbar p-1">
        
        <li v-if="options.length === 0 && !isLoading" class="p-4 text-center text-base text-slate-500 italic">
          Data tidak ditemukan.
        </li>

        <li 
          v-for="bird in options" 
          :key="bird.id"
          @click="selectItem(bird)"
          class="flex justify-between items-center px-4 py-3 hover:bg-blue-50 dark:hover:bg-slate-700/50 rounded-lg cursor-pointer group transition border-b border-slate-100 dark:border-slate-700 last:border-0"
          :class="{'bg-blue-50 dark:bg-slate-700/30': modelValue === bird.id}"
        >
          <div class="truncate">
            <div class="font-bold text-slate-800 dark:text-white text-base group-hover:text-blue-600 truncate transition-colors">
              {{ bird.visual }}
            </div>
            <div class="text-sm text-slate-500 font-mono mt-0.5">
              #{{ bird.ring_number }}
            </div>
          </div>
          <Check v-if="modelValue === bird.id" class="w-5 h-5 text-blue-600 flex-shrink-0"/>
        </li>
      </ul>
    </div>

    <div v-if="isOpen" @click="toggleDropdown" class="fixed inset-0 z-40 bg-transparent"></div>
  </div>
</template>

<style scoped>
/* Scrollbar Bersih & Rapi */
.custom-scrollbar::-webkit-scrollbar { width: 6px; }
.custom-scrollbar::-webkit-scrollbar-track { background: transparent; }
.custom-scrollbar::-webkit-scrollbar-thumb { background-color: #cbd5e1; border-radius: 10px; }
.custom-scrollbar::-webkit-scrollbar-thumb:hover { background-color: #94a3b8; }
</style>