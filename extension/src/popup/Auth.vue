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
import { handleLogin, UserProfile, getUserProfile } from '../logic'
import { useRouter } from 'vue-router'

const loading = ref(false)
const error = ref<string | null>(null)
const userProfile = ref<UserProfile | undefined>()

const router = useRouter()

const onLoginClick = async () => {
  loading.value = true
  error.value = null

  handleLogin()
    .then(async () => {
      userProfile.value = await getUserProfile()
      await router.push('/')
    })
    .catch((err) => {
      console.error(err)
      // TODO: show toaster here
      error.value =
        err.message || 'An unexpected error occurred. Please try again.'
    })
    .finally(() => (loading.value = false))
}

onMounted(async () => {
  userProfile.value = await getUserProfile()
})
</script>
