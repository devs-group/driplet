import { createMemoryHistory, createRouter } from 'vue-router'
import AuthView from './Auth.vue'
import PopupView from './Popup.vue'
import { ref } from 'vue'

const isAuthenticated = ref(false)

// Initialize auth state
const checkAuth = async () => {
  const key = 'accessToken'
  const result = await chrome.storage.local.get(key)
  isAuthenticated.value = !!result[key]
}

checkAuth()

const routes = [
  { path: '/', component: PopupView },
  { path: '/auth', component: AuthView }
]

export const router = createRouter({
  history: createMemoryHistory(),
  routes
})

router.beforeEach(async (to, from, next) => {
  await checkAuth()
  if (!isAuthenticated.value && to.path !== '/auth') {
    next('/auth')
  } else {
    next()
  }
})
