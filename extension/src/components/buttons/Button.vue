<template>
  <button
    class="relative flex items-center justify-center w-full p-3 bg-white border border-gray-300 rounded text-gray-700 text-sm cursor-pointer transition hover:bg-gray-100 hover:shadow disabled:opacity-70 disabled:cursor-not-allowed"
    :disabled="isLoading || disabled"
  >
    <span v-if="isLoading" class="spinner mr-2" />
    <span :class="{ 'opacity-0': isLoading }">
      <slot />
    </span>
  </button>
</template>

<script>
export default {
  props: {
    isLoading: {
      type: Boolean,
      default: false
    },
    disabled: {
      type: Boolean,
      default: false
    }
  }
}
</script>

<style scoped>
.spinner {
  display: inline-block;
  width: 1rem;
  height: 1rem;
  border: 2px solid rgba(0, 0, 0, 0.1);
  border-top-color: currentColor;
  border-radius: 50%;
  animation: spin 0.8s linear infinite;
  position: absolute;
  margin: 0;
}

@keyframes spin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}
</style>
