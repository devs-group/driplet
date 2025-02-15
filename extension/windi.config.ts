import { defineConfig } from 'windicss/helpers'

export default defineConfig({
  darkMode: 'class',
  // https://windicss.org/posts/v30.html#attributify-mode
  attributify: true,
  extract: {
    include: ['**/*.{vue,html}']
  },
  theme: {
    extend: {
      fontFamily: {
        sans: ['Inter', 'sans-serif'], // Default text font
        mono: ['Fira Code', 'monospace'] // Code font
      }
    }
  }
})
