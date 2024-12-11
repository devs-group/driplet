<template>
  <IntroSlider>
    <template #lastStep>
      <Button @click="onLoginClick">
        <img
          src="../assets/google.svg"
          alt="Google"
          class="google-icon w-4 h-4 mr-2"
        />
        {{ loading ? 'Signing in...' : 'Continue with Google' }}
      </Button>
      {{ error ?? '' }}
    </template>
  </IntroSlider>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import IntroSlider from '../components/IntroSlider.vue'
import Button from '../components/buttons/Button.vue'
import { storage } from '../logic/storage'
import { handleLogin, fetchUserProfile, UserProfile } from '../logic'
import { useRouter } from 'vue-router'

const loading = ref(false)
const error = ref<string | null>(null)
const userProfile = ref<UserProfile | undefined>()

const router = useRouter()

const onLoginClick = async () => {
  loading.value = true
  error.value = null

  try {
    const accessToken = await handleLogin()
    if (accessToken) {
      storage.accessToken.value = accessToken
      userProfile.value = await fetchUserProfile(accessToken)
      storage.isFirstTimeUser.value = false
      router.push('/')
    }
  } catch (err: any) {
    console.error(err)
    error.value =
      err.message || 'An unexpected error occurred. Please try again.'
  } finally {
    loading.value = false
  }
}

// Restore session on mount
onMounted(async () => {
  chrome.storage.local.get(['accessToken', 'userProfile'], (result) => {
    if (result.accessToken) {
      storage.accessToken.value = result.accessToken
      userProfile.value = result.userProfile || null
    }
  })
})
</script>
