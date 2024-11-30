<template>
  <main class="w-[300px] px-4 py-5">
    <div class="flex flex-col items-center">
      <div class="flex justify-between items-center w-full mb-4">
        <h1 class="text-xl font-bold">Dripplet</h1>
        <button 
          v-if="isAuthenticated"
          @click="openSettings"
          class="p-2 text-gray-600 hover:text-gray-900 dark:text-gray-400 dark:hover:text-white"
          title="Settings"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
          </svg>
        </button>
      </div>

      <!-- Login Component -->
      <Login class="w-full" />

      <!-- Show this section only when logged in -->
      <template v-if="isAuthenticated">
        <div class="mt-4 bg-gray-100 dark:bg-gray-800 p-3 rounded-lg w-full">
          <h2 class="text-lg font-bold">Rewards Earned</h2>
          <p class="text-2xl font-semibold text-green-500">14,213 Driplets</p>
        </div>

        <div class="mt-4 w-full">
          <h3 class="text-lg font-semibold mb-2">Tracking Options</h3>
          <div class="flex flex-col items-start w-full">
            <label class="inline-flex items-center mt-2">
              <input
                type="checkbox"
                v-model="trackBrowsing"
                class="form-checkbox h-5 w-5 text-blue-600"
                @change="updateTracking"
              />
              <span class="ml-2">Track Browsing</span>
            </label>
            <label class="inline-flex items-center mt-2">
              <input
                type="checkbox"
                v-model="trackInteractions"
                class="form-checkbox h-5 w-5 text-blue-600"
                @change="updateTracking"
              />
              <span class="ml-2">Track Interactions</span>
            </label>
            <label class="inline-flex items-center mt-2">
              <input
                type="checkbox"
                v-model="crawlWebsites"
                class="form-checkbox h-5 w-5 text-blue-600"
                @change="updateTracking"
              />
              <span class="ml-2">Crawl Websites</span>
            </label>
          </div>
        </div>

        <Footer />

        <div class="mt-2">
          <span class="opacity-50">{{ $t('popup.publicKey') }}:</span>
          {{ solanaPublicKeyStorage }}
        </div>
      </template>
    </div>
  </main>
</template>
<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import {
  solanaPublicKeyStorage,
  trackBrowsingStorage,
  trackInteractionsStorage,
  crawlWebsitesStorage,
  accessTokenStorage
} from '~/logic/storage'
import Login from '../components/Login.vue'

const trackBrowsing = ref(false)
const trackInteractions = ref(false)
const crawlWebsites = ref(false)

const isAuthenticated = computed(() => !!accessTokenStorage.value)
console.log('isAuthenticated', isAuthenticated)
onMounted(() => {
  trackBrowsing.value = trackBrowsingStorage.value
  trackInteractions.value = trackInteractionsStorage.value
  crawlWebsites.value = crawlWebsitesStorage.value
})

function openOptionsPage() {
  chrome.runtime.openOptionsPage()
}

function openSettings() {
  chrome.runtime.openOptionsPage()
}

function updateTracking() {
  // Here you would typically update the background script or service worker
  // to reflect the new tracking settings
  trackBrowsingStorage.value = trackBrowsing.value
  trackInteractionsStorage.value = trackInteractions.value
  crawlWebsitesStorage.value = crawlWebsites.value

  console.log('Tracking settings updated:', {
    trackBrowsing: trackBrowsing.value,
    trackInteractions: trackInteractions.value,
    crawlWebsites: crawlWebsites.value
  })
}
</script>
