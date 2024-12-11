import { useLocalStorage } from '@vueuse/core'

const solanaPublicKey = useLocalStorage('soalan-public-key', '', {
  listenToStorageChanges: true
})

const trackBrowsing = useLocalStorage('track-browsing', true, {
  listenToStorageChanges: true
})

const trackInteractions = useLocalStorage('track-interaction', true, {
  listenToStorageChanges: true
})

const crawlWebsites = useLocalStorage('crawl-websites', true, {
  listenToStorageChanges: true
})

const emailNotifications = useLocalStorage('email-notifications', true, {
  listenToStorageChanges: true
})

const browserNotifications = useLocalStorage('browser-notifications', true, {
  listenToStorageChanges: true
})

const accessToken = useLocalStorage<string | null>('access-token', null, {
  listenToStorageChanges: true
})

const isFirstTimeUser = useLocalStorage('first-time-user', true, {
  listenToStorageChanges: true
})

export const storage = {
  isFirstTimeUser,
  accessToken,
  browserNotifications,
  emailNotifications,
  crawlWebsites,
  trackBrowsing,
  trackInteractions,
  solanaPublicKey
}

export const clearAllStorage = () => {
  solanaPublicKey.value = ''
  trackBrowsing.value = true
  trackInteractions.value = true
  crawlWebsites.value = true
  accessToken.value = null
  isFirstTimeUser.value = true
  emailNotifications.value = true
  browserNotifications.value = true
  accessToken.value = null
}
