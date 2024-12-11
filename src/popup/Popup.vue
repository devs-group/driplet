<template>
  <main class="w-[300px] px-4 py-5 bg-black">
    <div class="flex flex-row justify-between items-center mb-8">
      <div class="text-gray-200">
        <div class="text-xl">Hello {{ userProfile?.name }}</div>
        <div class="text-md">Good Morning 👋</div>
      </div>
      <div class="w-12">
        <img :src="userProfile?.picture" class="rounded-full" />
      </div>
    </div>
    <div class="flex flex-col items-center">
      <div class="mt-4 bg-gray-100 dark:bg-gray-800 p-3 rounded-md w-full">
        <h2 class="text-black">Curernt earnings</h2>
        <p class="text-lg font-semibold text-black font-mono">
          1.234.030,32 $DRIPL
        </p>
      </div>

      <div class="mt-4">
        <LineChart />
      </div>

      <div class="mt-4 w-full">
        <h3 class="text-lg font-semibold mb-2">Tracking Options</h3>
        <div class="flex flex-col items-start w-full">
          <label class="inline-flex items-center mt-2">
            <input
              type="checkbox"
              v-model="trackBrowsing"
              @change="() => (storage.trackBrowsing.value = trackBrowsing)"
              class="form-checkbox h-5 w-5 text-blue-600"
            />
            <span class="ml-2">Track Browsing</span>
          </label>
          <label class="inline-flex items-center mt-2">
            <input
              type="checkbox"
              v-model="trackInteractions"
              @change="
                () => (storage.trackInteractions.value = trackInteractions)
              "
              class="form-checkbox h-5 w-5 text-blue-600"
            />
            <span class="ml-2">Track Interactions</span>
          </label>
          <label class="inline-flex items-center mt-2">
            <input
              type="checkbox"
              v-model="crawlWebsites"
              @change="() => (storage.crawlWebsites.value = crawlWebsites)"
              class="form-checkbox h-5 w-5 text-blue-600"
            />
            <span class="ml-2">Crawl Websites</span>
          </label>
        </div>
      </div>

      <div class="flex justify-between items-center w-full my-4">
        <button
          @click="openSettings"
          class="p-2 text-gray-600 hover:text-gray-400 dark:text-gray-400 dark:hover:text-white"
          title="Settings"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-6 w-6"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"
            />
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
            />
          </svg>
        </button>
        <div class="text-gray-400 cursor-pointer" @click="onClickLogout">
          Logout
        </div>
      </div>
    </div>
  </main>
</template>
<script setup lang="ts">
import { onMounted } from 'vue'
import { storage } from '~/logic/storage'
import { handleLogout, getUserProfile, UserProfile } from '../logic'
import { useRouter } from 'vue-router'
import LineChart from '../components/charts/LineChart.vue'

const userProfile = ref<UserProfile>()

const router = useRouter()

const trackBrowsing = ref(false)
const trackInteractions = ref(false)
const crawlWebsites = ref(false)

onMounted(async () => {
  userProfile.value = await getUserProfile()
  trackBrowsing.value = storage.trackBrowsing.value
  trackInteractions.value = storage.trackInteractions.value
  crawlWebsites.value = storage.crawlWebsites.value
})

function openSettings() {
  chrome.runtime.openOptionsPage()
}

async function onClickLogout() {
  try {
    await handleLogout()
    router.push('/auth')
  } catch (err) {
    // TODO: show error notification here.
  }
}
</script>
