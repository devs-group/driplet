<template>
  <div class="fixed inset-0 bg-white dark:bg-gray-900 z-50">
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
                <div class="w-24 h-24 mx-auto mb-6 flex items-center justify-center">
                  <template v-if="slide.image">
                    <img
                      :src="slide.image"
                      :alt="slide.title"
                      class="w-24 h-24 mx-auto mb-6"
                    />
                  </template>
                  <template v-else-if="slide.icon">
                    <!-- Use SVG icons or components -->
                    <img :src="slide.icon" class="w-16 h-16 text-blue-500" />
                  </template>
                </div>
                <h2 class="text-2xl font-bold mb-4">{{ slide.title }}</h2>
                <p class="text-gray-600 dark:text-gray-400 mb-4">{{ slide.description }}</p>
                <div v-if="slide.isButtonVisible">
                  <button
                    @click="$emit('finish')"
                    class="bg-blue-500 text-white px-6 py-2 rounded-lg hover:bg-blue-600 transition-colors"
                  >
                    Login to Continue
                  </button>
                </div>
              </div>
            </div>
          </div>

          <!-- Dots -->
          <div class="flex justify-center mt-6 space-x-2">
            <button
              v-for="(slide, index) in slides"
              :key="index"
              @click="goToSlide(index)"
              class="w-2 h-2 rounded-full transition-colors focus:outline-none"
              :class="currentSlide === index ? 'bg-blue-500' : 'bg-gray-300 dark:bg-gray-700'"
              :aria-label="`Go to slide ${index + 1}`"
            ></button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onUnmounted } from 'vue';

interface Slide {
  title: string;
  description: string;
  image?: string;
  icon?: string;
  isButtonVisible: boolean;
}

const slides: Slide[] = [
  {
    title: 'Welcome to Dripplet',
    description: 'Your personal web companion that rewards you for browsing.',
    image: '../assets/icon.svg',
    isButtonVisible: false,
  },
  {
    title: 'Earn Rewards',
    description: 'Get rewarded with Driplets for your daily browsing activities.',
    icon: '../assets/reward.svg',
    isButtonVisible: false,
  },
  {
    title: "Let's Get Started",
    description:
      'Sign in to start earning rewards and customize your experience.',
    icon: '../assets/earn.svg',
    isButtonVisible: true,
  },
];

const currentSlide = ref<number>(0);
let autoAdvance: number;

const startAutoAdvance = () => {
  autoAdvance = setInterval(() => {
    if (currentSlide.value < slides.length - 1) {
      currentSlide.value++;
    } else {
      clearInterval(autoAdvance);
    }
  }, 5000);
};

const resetAutoAdvance = () => {
  clearInterval(autoAdvance);
  startAutoAdvance();
};

const goToSlide = (index: number) => {
  currentSlide.value = index;
  resetAutoAdvance();
};

startAutoAdvance();

onUnmounted(() => {
  clearInterval(autoAdvance);
});
</script>
