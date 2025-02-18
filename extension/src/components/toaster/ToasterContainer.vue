<template>
  <Teleport to="body">
    <div
      class="fixed z-50 flex flex-col gap-3"
      :class="[
        position.includes('top') ? 'top-4' : 'bottom-4',
        position.includes('right')
          ? 'right-4'
          : position.includes('left')
            ? 'left-4'
            : 'left-1/2 transform -translate-x-1/2'
      ]"
    >
      <TransitionGroup name="toast">
        <div
          v-for="toast in toasts"
          :key="toast.id"
          class="flex items-center bg-black w-72 px-4 py-3 rounded-lg shadow-lg border-1 border-gray-700 transition-all duration-300"
          :class="[
            toast.type === 'success' && 'border-l-4 border-l-green-500',
            toast.type === 'error' && 'border-l-4 border-l-red-500',
            toast.type === 'warning' && 'border-l-4 border-l-yellow-500',
            toast.type === 'info' && 'border-l-4 border-l-blue-500'
          ]"
        >
          <!-- Icon -->
          <div class="flex-shrink-0 mr-3">
            <svg
              v-if="toast.type === 'success'"
              class="w-5 h-5 text-green-500"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path
                fill-rule="evenodd"
                d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                clip-rule="evenodd"
              />
            </svg>
            <svg
              v-if="toast.type === 'error'"
              class="w-5 h-5 text-red-500"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path
                fill-rule="evenodd"
                d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z"
                clip-rule="evenodd"
              />
            </svg>
            <svg
              v-if="toast.type === 'warning'"
              class="w-5 h-5 text-yellow-500"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path
                fill-rule="evenodd"
                d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z"
                clip-rule="evenodd"
              />
            </svg>
            <svg
              v-if="toast.type === 'info'"
              class="w-5 h-5 text-blue-500"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path
                fill-rule="evenodd"
                d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z"
                clip-rule="evenodd"
              />
            </svg>
          </div>

          <!-- Content -->
          <div class="flex-grow">
            <p v-if="toast.title" class="font-medium text-sm text-white">
              {{ toast.title }}
            </p>
            <p class="text-sm text-gray-200">
              {{ toast.message }}
            </p>
          </div>

          <!-- Close button -->
          <button
            class="flex-shrink-0 ml-2 rounded-full p-1 text-gray-400 hover:text-gray-600 dark:text-gray-300 dark:hover:text-white focus:outline-none"
            @click="removeToast(toast.id)"
          >
            <svg
              class="w-4 h-4"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path
                fill-rule="evenodd"
                d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
                clip-rule="evenodd"
              />
            </svg>
          </button>
        </div>
      </TransitionGroup>
    </div>
  </Teleport>
</template>

<script setup>
import { ref, onMounted, computed, onUnmounted } from 'vue'

const props = defineProps({
  position: {
    type: String,
    default: 'top-right',
    validator: (value) => {
      return [
        'top-right',
        'top-left',
        'top-center',
        'bottom-right',
        'bottom-left',
        'bottom-center'
      ].includes(value)
    }
  },
  maxToasts: {
    type: Number,
    default: 5
  }
})

const toasts = ref([])
const toastTimeouts = ref({})
const progressIntervals = ref({})

// Toast system
const toast = {
  success: (message, options = {}) =>
    addToast({ message, type: 'success', ...options }),
  error: (message, options = {}) =>
    addToast({ message, type: 'error', ...options }),
  warning: (message, options = {}) =>
    addToast({ message, type: 'warning', ...options }),
  info: (message, options = {}) =>
    addToast({ message, type: 'info', ...options })
}

// Create toast
function addToast({
  message,
  type = 'info',
  title = '',
  duration = 5000,
  showProgress = true
}) {
  // Create a new toast
  const id = Date.now().toString()
  const newToast = {
    id,
    message,
    type,
    title,
    duration,
    showProgress,
    progress: 100
  }

  // Remove oldest toast if at max
  if (toasts.value.length >= props.maxToasts) {
    const oldestToast = toasts.value[0]
    removeToast(oldestToast.id)
  }

  // Add new toast
  toasts.value.push(newToast)

  // Set timeout to remove toast
  if (duration !== Infinity) {
    // Timeout to remove toast
    toastTimeouts.value[id] = setTimeout(() => {
      removeToast(id)
    }, duration)

    // Progress interval for animation
    if (showProgress) {
      let startTime = Date.now()
      let endTime = startTime + duration

      progressIntervals.value[id] = setInterval(() => {
        const now = Date.now()
        const remaining = endTime - now
        const percent = (remaining / duration) * 100

        // Find the toast and update its progress
        const toastIndex = toasts.value.findIndex((t) => t.id === id)
        if (toastIndex !== -1) {
          toasts.value[toastIndex].progress = Math.max(0, percent)
        }

        // Clear interval when progress reaches 0
        if (percent <= 0) {
          clearInterval(progressIntervals.value[id])
        }
      }, 10)
    }
  }

  return id
}

function removeToast(id) {
  const index = toasts.value.findIndex((t) => t.id === id)
  if (index !== -1) {
    // Clear timeout and interval
    clearTimeout(toastTimeouts.value[id])
    clearInterval(progressIntervals.value[id])
    delete toastTimeouts.value[id]
    delete progressIntervals.value[id]

    // Remove toast
    toasts.value.splice(index, 1)
  }
}

onUnmounted(() => {
  Object.values(toastTimeouts.value).forEach((timeout) => clearTimeout(timeout))
  Object.values(progressIntervals.value).forEach((interval) =>
    clearInterval(interval)
  )
})

// Expose methods and data for external use
defineExpose({ toast, removeToast, toasts })
</script>

<style scoped>
.toast-enter-active,
.toast-leave-active {
  transition: all 0.25s cubic-bezier(0.4, 0, 0.2, 1);
}

.toast-enter-from {
  opacity: 0;
  transform: translateY(8px);
}

.toast-leave-to {
  opacity: 0;
  transform: translateX(8px);
}
</style>
