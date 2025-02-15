<template>
  <div class="min-h-screen bg-white dark:bg-gray-900">
    <header class="bg-white dark:bg-gray-800 shadow">
      <div class="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
        <h1 class="text-3xl font-bold text-gray-900 dark:text-white">Settings</h1>
      </div>
    </header>
    
    <main class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
      <div class="px-4 py-6 sm:px-0">
        <div class="bg-white dark:bg-gray-800 shadow rounded-lg p-6">
          <!-- Account Settings -->
          <section class="mb-8">
            <h2 class="text-xl font-semibold mb-4 text-gray-900 dark:text-white">Account Settings</h2>
            <div class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">Email</label>
                <p class="mt-1 text-sm text-gray-500">{{ email }}</p>
              </div>
              <button @click="logout" class="bg-red-600 text-white px-4 py-2 rounded-md hover:bg-red-700">
                Logout
              </button>
            </div>
          </section>

          <!-- Privacy Settings -->
          <section class="mb-8">
            <h2 class="text-xl font-semibold mb-4 text-gray-900 dark:text-white">Privacy Settings</h2>
            <div class="space-y-4">
              <label class="flex items-center space-x-3">
                <input
                  type="checkbox"
                  v-model="trackBrowsing"
                  class="form-checkbox h-5 w-5 text-blue-600"
                  @change="updateTracking"
                />
                <span class="text-gray-700 dark:text-gray-300">Track Browsing</span>
              </label>
              <label class="flex items-center space-x-3">
                <input
                  type="checkbox"
                  v-model="trackInteractions"
                  class="form-checkbox h-5 w-5 text-blue-600"
                  @change="updateTracking"
                />
                <span class="text-gray-700 dark:text-gray-300">Track Interactions</span>
              </label>
            </div>
          </section>

          <!-- Notification Settings -->
          <section>
            <h2 class="text-xl font-semibold mb-4 text-gray-900 dark:text-white">Notification Settings</h2>
            <div class="space-y-4">
              <label class="flex items-center space-x-3">
                <input
                  type="checkbox"
                  v-model="emailNotifications"
                  class="form-checkbox h-5 w-5 text-blue-600"
                  @change="updateNotifications"
                />
                <span class="text-gray-700 dark:text-gray-300">Email Notifications</span>
              </label>
              <label class="flex items-center space-x-3">
                <input
                  type="checkbox"
                  v-model="browserNotifications"
                  class="form-checkbox h-5 w-5 text-blue-600"
                  @change="updateNotifications"
                />
                <span class="text-gray-700 dark:text-gray-300">Browser Notifications</span>
              </label>
            </div>
          </section>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useStorage } from '@vueuse/core'

// User data
const email = ref('')
const trackBrowsing = useStorage('trackBrowsing', true)
const trackInteractions = useStorage('trackInteractions', true)
const emailNotifications = useStorage('emailNotifications', true)
const browserNotifications = useStorage('browserNotifications', true)

onMounted(async () => {
  // TODO: Fetch user data from your auth service
  email.value = 'user@example.com'
})

const updateTracking = async () => {
  // TODO: Update tracking settings in your backend
  console.log('Tracking settings updated:', {
    trackBrowsing: trackBrowsing.value,
    trackInteractions: trackInteractions.value,
  })
}

const updateNotifications = async () => {
  // TODO: Update notification settings in your backend
  console.log('Notification settings updated:', {
    emailNotifications: emailNotifications.value,
    browserNotifications: browserNotifications.value,
  })
}

const logout = async () => {
  // TODO: Implement logout logic
  console.log('Logging out...')
}
</script>
