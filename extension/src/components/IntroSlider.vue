<template>
  <div class="fixed inset-0 z-50">
    <div class="h-full flex flex-col items-center justify-center px-4">
      <div class="w-full max-w-md">
        <!-- Slider -->
        <div class="relative">
          <div class="overflow-hidden">
            <div
              class="flex transition-transform duration-300"
              :style="{ transform: `translateX(-${currentSlide * 100}%)` }"
            >
              <!-- Slides -->
              <div
                v-for="(slide, index) in slides"
                :key="index"
                class="w-full flex-shrink-0 text-center px-4"
              >
                <div class="w-24 mx-auto mb-6 flex items-center justify-center">
                  <template v-if="slide.image">
                    <img
                      :src="slide.image"
                      :alt="slide.title"
                      class="w-24 mx-auto mb-6"
                    />
                  </template>
                  <template v-else-if="slide.icon">
                    <!-- Use SVG icons or components -->
                    <img :src="slide.icon" class="w-16 h-16 text-blue-500" />
                  </template>
                </div>
                <h2 class="text-2xl font-bold mb-4">{{ slide.title }}</h2>
                <p class="mb-4">
                  {{ slide.description }}
                </p>
                <div v-if="slide.isLastStep">
                  <slot name="lastStep" />
                </div>
              </div>
            </div>
          </div>

          <!-- Dots -->
          <div class="flex justify-center mt-6 space-x-2">
            <button
              v-for="(_, index) in slides"
              :key="index"
              @click="goToSlide(index)"
              class="w-2 h-2 rounded-full transition-colors focus:outline-none"
              :class="
                currentSlide === index
                  ? 'bg-blue-500'
                  : 'bg-gray-300 dark:bg-gray-700'
              "
              :aria-label="`Go to slide ${index + 1}`"
            ></button>
          </div>

          <div class="mt-5">
            <Button
              v-if="currentSlide < slides.length - 1"
              @click="goToSlide(currentSlide + 1)"
              >Next</Button
            >
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onUnmounted } from 'vue'
import Button from '../components/buttons/Button.vue'

interface Slide {
  title: string
  description: string
  image?: string
  icon?: string
  isLastStep: boolean
}

const slides: Slide[] = [
  {
    title: 'Welcome to Dripplet',
    description: 'Your personal web companion that rewards you for browsing.',
    image: '../assets/driplet-logo.png',
    isLastStep: false
  },
  {
    title: 'Earn Rewards',
    description:
      'Get rewarded with Driplets for your daily browsing activities.',
    icon: '../assets/reward.svg',
    isLastStep: false
  },
  {
    title: "Let's Get Started",
    description:
      'Sign in to start earning rewards and customize your experience.',
    icon: '../assets/earn.svg',
    isLastStep: true
  }
]

const currentSlide = ref<number>(0)
let autoAdvance: number

const startAutoAdvance = () => {
  // @ts-ignore set interval returtns Nodejs.Timer type
  autoAdvance = setInterval(() => {
    if (currentSlide.value < slides.length - 1) {
      currentSlide.value++
    } else {
      clearInterval(autoAdvance)
    }
  }, 5000)
}

const resetAutoAdvance = () => {
  clearInterval(autoAdvance)
  startAutoAdvance()
}

const goToSlide = (index: number) => {
  if (index > slides.length - 1) {
    return
  }
  currentSlide.value = index
  resetAutoAdvance()
}

startAutoAdvance()

onUnmounted(() => {
  clearInterval(autoAdvance)
})
</script>
