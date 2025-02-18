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
          {{ formattedCredits }} $DRIPL
        </p>
      </div>

      <div class="mt-4">
        <LineChart />
      </div>

      <div class="mt-6 w-full">
        <h3 class="text-lg font-semibold mb-2">Your Public Key</h3>
        <div class="flex flex-col items-start w-full">
          <p class="pb-2 text-gray-400">
            Please enter here your public key to which $DRIPL tokens will be
            transfered to.
          </p>
          <div class="display flex flex-row items-center space-x-2">
            <input
              type="text"
              v-model="publicKey"
              class="text-black text-lg focus:outline-none rounded-md px-2 h-11"
            />
            <Button @click="savePublicKey" :isLoading="isLoading.savePublicKey">
              Save
            </Button>
          </div>
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

    <!-- Toast Container -->
    <ToastContainer position="bottom-center" :maxToasts="3" ref="toastRef" />
  </main>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { handleLogout, getUserProfile, UserProfile } from '../logic'
import { UserApiService } from '~/background/services/user'
import { useRouter } from 'vue-router'
import LineChart from '../components/charts/LineChart.vue'
import Button from '../components/buttons/Button.vue'
import ToastContainer from '../components/toaster/ToasterContainer.vue'

const router = useRouter()
const toastRef = ref<any>(null)

const userProfile = ref<UserProfile>()
const credits = ref(0)
const publicKey = ref('')
const isLoading = ref({
  savePublicKey: false
})

const formattedCredits = computed(() => formatCreditsNumber(credits.value))

onMounted(() => {
  setCurrentUserProfile()
  setCurrentUserCredits()
})

async function setCurrentUserProfile() {
  userProfile.value = await getUserProfile()
}

async function setCurrentUserCredits() {
  const userApiService = new UserApiService()
  try {
    const response = await userApiService.GET_User()
    credits.value = response?.credits || 0
    publicKey.value = response?.public_key || ''
  } catch (err) {
    router.push('/auth')
    console.error(err)
  }
}

function openSettings() {
  chrome.runtime.openOptionsPage()
}

async function onClickLogout() {
  try {
    await handleLogout()
    router.push('/auth')
  } catch (err) {
    console.error(err)
    router.push('/auth')
  }
}

function formatCreditsNumber(credits: number) {
  return new Intl.NumberFormat('de-DE', {
    style: 'decimal',
    minimumFractionDigits: 2,
    maximumFractionDigits: 2
  }).format(credits)
}

async function savePublicKey() {
  isLoading.value.savePublicKey = true
  const userApiService = new UserApiService()
  try {
    await userApiService.PUT_UpdateUsersPublicKey(publicKey.value)
    toastRef.value?.toast.success(
      'Your public key has been saved successfully',
      {
        duration: 4000
      }
    )
  } catch (err) {
    console.error(err)
    toastRef.value?.toast.error('Failed to save public key', {
      title: 'Error',
      duration: 4000
    })
  } finally {
    isLoading.value.savePublicKey = false
  }
}
</script>
