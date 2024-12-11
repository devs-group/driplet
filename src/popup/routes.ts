import { createMemoryHistory, createRouter } from 'vue-router'
import { storage } from '~/logic/storage'
import AuthView from './Auth.vue'
import PopupView from './Popup.vue'
import { computed } from 'vue'

const isAuthenticated = computed(() => !!storage.accessToken.value)

const routes = [
  { path: '/', component: PopupView },
  { path: '/auth', component: AuthView }
]

export const router = createRouter({
  history: createMemoryHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  if (!isAuthenticated.value && to.path !== '/auth') {
    next('/auth')
  } else {
    next()
  }
})
