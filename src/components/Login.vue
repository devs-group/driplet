<template>
  <div class="login-container w-full p-5 text-center">
    <div v-if="!isAuthenticated" class="login-section">
      <div class="login-header mb-6">
        <img src="../assets/icon.svg" alt="Dripplet Logo" class="logo w-16 h-16 mb-4" />
        <h2 class="text-xl font-semibold text-gray-900 mb-2">Welcome to Dripplet</h2>
        <p class="text-gray-600 text-sm">Sign in to start earning rewards</p>
      </div>
      <button @click="handleLogin" class="login-button flex items-center justify-center w-full p-3 bg-white border border-gray-300 rounded text-gray-700 text-sm cursor-pointer transition hover:bg-gray-100 hover:shadow disabled:opacity-70 disabled:cursor-not-allowed" :disabled="loading">
        <img src="../assets/google.svg" alt="Google" class="google-icon w-4 h-4 mr-2" v-if="!loading" />
        <span v-if="loading" class="loading-spinner w-4 h-4 mr-2 border-2 border-gray-200 border-t-blue-500 rounded-full animate-spin"></span>
        {{ loading ? 'Signing in...' : 'Continue with Google' }}
      </button>
      <p v-if="error" class="error-message text-red-600 mt-3 text-xs">{{ error }}</p>
    </div>
    <div v-else-if="isAuthenticated && userProfile" class="logged-in-section text-left">
      <div class="user-info flex items-center mb-4 p-3 bg-gray-100 rounded">
        <img :src="userProfile.picture" alt="Profile" class="profile-picture w-12 h-12 rounded-full mr-3" />
        <div class="user-details">
          <h3 class="text-base text-gray-900 mb-1">{{ userProfile.name }}</h3>
          <p class="text-xs text-gray-600">{{ userProfile.email }}</p>
        </div>
      </div>
      <button @click="handleLogout" class="logout-button w-full p-2 bg-red-600 text-white rounded text-sm cursor-pointer transition hover:bg-red-500">
        Sign Out
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { accessTokenStorage, clearAllStorage } from '~/logic/storage'

interface UserProfile {
  name: string
  email: string
  picture: string
}

const isAuthenticated = computed(() => !!accessTokenStorage.value)
const loading = ref(false)
const error = ref<string | null>(null)
const userProfile = ref<UserProfile | null>(null)

const handleLogin = async () => {
  loading.value = true
  error.value = null

  try {
    const response = await new Promise((resolve) => {
      chrome.runtime.sendMessage({ action: 'login' }, resolve)
    })

    if (response && response.success) {
      accessTokenStorage.value = response.accessToken
      await fetchUserProfile(response.accessToken)
    } else {
      error.value = response.error || 'Login failed. Please try again.'
    }
  } catch (err) {
    error.value = 'An unexpected error occurred. Please try again.'
  } finally {
    loading.value = false
  }
}

const handleLogout = async () => {
  clearAllStorage()
  accessTokenStorage.value = null
  userProfile.value = null
}

const fetchUserProfile = async (accessToken: string) => {
  try {
    const response = await fetch('https://www.googleapis.com/oauth2/v2/userinfo', {
      headers: {
        Authorization: `Bearer ${accessToken}`
      }
    })
    
    if (!response.ok) {
      throw new Error('Failed to fetch user profile')
    }

    const data = await response.json()
    userProfile.value = {
      name: data.name,
      email: data.email,
      picture: data.picture
    }

    // Store user profile
    chrome.storage.local.set({ userProfile: userProfile.value })
  } catch (err) {
    console.error('Error fetching user profile:', err)
  }
}

onMounted(async () => {
  chrome.storage.local.get(['accessToken', 'userProfile'], (result) => {
    if (result.accessToken) {
      accessTokenStorage.value = result.accessToken
      userProfile.value = result.userProfile || null
    }
  })
})
</script>
