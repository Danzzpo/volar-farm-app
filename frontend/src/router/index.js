import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import RegisterView from '../views/RegisterView.vue'
import DashboardView from '../views/DashboardView.vue'
import AnimalsView from '../views/AnimalsView.vue'
import OverviewView from '../views/OverviewView.vue'
import PairsView from '../views/PairsView.vue'
import SettingsView from '../views/SettingsView.vue'
import IncubatingView from '../views/IncubatingView.vue' 
import FinanceView from '../views/FinanceView.vue'
import TreatmentsView from '../views/TreatmentsView.vue'// <--- 1. Import Ini

// Placeholder Page
const ComingSoon = { template: '<div class="p-10 text-center text-slate-400 text-xl border-2 border-dashed border-slate-300 rounded-xl">🚧 Fitur ini sedang dibangun!</div>' }

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/', name: 'home', component: HomeView },
    { path: '/login', name: 'login', component: LoginView },
    { path: '/register', name: 'register', component: RegisterView },
    
    // --- DASHBOARD ---
    { 
      path: '/dashboard', 
      component: DashboardView, 
      children: [
        { path: '', name: 'dashboard-home', component: OverviewView },
        { path: 'animals', name: 'animals', component: AnimalsView },
        { path: 'pairs', name: 'pairs', component: PairsView },
        { path: 'settings', name: 'settings', component: SettingsView }, // <--- 2. Tambah Ini
        { path: 'incubating', name: 'incubating', component: IncubatingView }, 
        { path: 'gencal', component: ComingSoon },
        { path: 'treatments', component: TreatmentsView },
        { path: 'finance', component: FinanceView },
      ]
    },
  ]
})

export default router