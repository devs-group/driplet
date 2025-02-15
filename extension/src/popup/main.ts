import 'vue-global-api'
import { createApp } from 'vue'
import App from './App.vue'
import '../styles'
import { router } from './routes'

import { Buffer } from 'buffer'
window.Buffer = Buffer

const app = createApp(App)
app.use(router).mount('#app')
