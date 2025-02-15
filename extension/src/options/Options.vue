<template>
  <div class="min-h-screen bg-white dark:bg-gray-900">
    <header class="bg-white dark:bg-gray-800 shadow">
      <div
        class="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8 flex items-center justify-between"
      >
        <h1 class="text-3xl font-bold text-gray-900">Settings</h1>
        <button
          @click="handleLogout"
          class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2"
        >
          Logout
        </button>
      </div>
    </header>

    <main class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
      <div class="px-4 py-6 sm:px-0">
        <div class="bg-white dark:bg-gray-800 shadow rounded-lg p-6">
          <!-- Solana Settings -->
          <section class="mb-8">
            <h2 class="text-xl font-semibold mb-4 text-gray-900">
              Solana Wallet
            </h2>
            <div class="space-y-4">
              <div>
                <label
                  for="solanaKey"
                  class="block text-sm font-medium text-black"
                >
                  Solana Public Key
                </label>
                <div class="mt-1">
                  <div class="flex gap-2">
                    <input
                      id="solanaKey"
                      v-model="solanaPublicKey"
                      placeholder="Enter your Solana Public Key"
                      class="text-black mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                    />
                    <button
                      @click="savePublicKey"
                      class="inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
                    >
                      Save
                    </button>
                  </div>
                  <p
                    v-if="solanaPublicKey"
                    class="mt-2 text-sm text-gray-600"
                  >
                    Current Public Key: {{ solanaPublicKey }}
                    <button
                      @click="editPublicKey"
                      class="ml-2 text-blue-600 hover:text-blue-800"
                    >
                      Edit
                    </button>
                  </p>
                </div>
              </div>
            </div>
          </section>

          <!-- Privacy Settings -->
          <section class="mb-8">
            <h2 class="text-xl font-semibold mb-4 text-gray-900">
              Privacy Settings
            </h2>
            <div class="space-y-4">
              <label class="flex items-center space-x-3">
                <input
                  type="checkbox"
                  v-model="trackBrowsing"
                  class="form-checkbox h-5 w-5 text-blue-600"
                  @change="updateTracking"
                />
                <span class="text-gray-700 dark:text-gray-300"
                  >Track Browsing</span
                >
              </label>
              <label class="flex items-center space-x-3">
                <input
                  type="checkbox"
                  v-model="trackInteractions"
                  class="form-checkbox h-5 w-5 text-blue-600"
                  @change="updateTracking"
                />
                <span class="text-gray-700 dark:text-gray-300"
                  >Track Interactions</span
                >
              </label>
            </div>
          </section>

          <!-- Notification Settings -->
          <section>
            <h2 class="text-xl font-semibold mb-4 text-gray-900">
              Notification Settings
            </h2>
            <div class="space-y-4">
              <label class="flex items-center space-x-3">
                <input
                  type="checkbox"
                  v-model="storage.emailNotifications"
                  class="form-checkbox h-5 w-5 text-blue-600"
                  @change="updateNotifications"
                />
                <span class="text-gray-700 dark:text-gray-300"
                  >Email Notifications</span
                >
              </label>
              <label class="flex items-center space-x-3">
                <input
                  type="checkbox"
                  v-model="storage.browserNotifications"
                  class="form-checkbox h-5 w-5 text-blue-600"
                  @change="updateNotifications"
                />
                <span class="text-gray-700 dark:text-gray-300"
                  >Browser Notifications</span
                >
              </label>
            </div>
          </section>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { storage, clearAllStorage } from '~/logic/storage'

const solanaPublicKey = ref('')
const trackBrowsing = ref(false)
const trackInteractions = ref(false)

onMounted(async () => {
  solanaPublicKey.value = storage.solanaPublicKey.value
  trackBrowsing.value = storage.trackBrowsing.value
  trackInteractions.value = storage.trackInteractions.value
})

const handleLogout = () => {
  clearAllStorage()
  // Close the options page after logout
  window.close()
}

const savePublicKey = () => {
  if (solanaPublicKey.value.trim()) {
    storage.solanaPublicKey.value = solanaPublicKey.value.trim()
    solanaPublicKey.value = ''
  }
}

const editPublicKey = () => {
  solanaPublicKey.value = storage.solanaPublicKey.value
  storage.solanaPublicKey.value = ''
}

const updateTracking = async () => {
  // Send message to background script to update tracking settings
  chrome.runtime.sendMessage({
    type: 'UPDATE_TRACKING',
    data: {
      trackBrowsing: storage.trackBrowsing.value === true,
      trackInteractions: storage.trackInteractions.value === true,
    }
  })
}

const updateNotifications = async () => {
  // Send message to background script to update notification settings
  chrome.runtime.sendMessage({
    type: 'UPDATE_NOTIFICATIONS',
    data: {
      emailNotifications: storage.emailNotifications.value === true,
      browserNotifications: storage.browserNotifications.value === true
    }
  })
}
</script>
