import { useLocalStorage } from '@vueuse/core'

export const solanaPublicKeyStorage = useLocalStorage(
  'soalan-public-key',
  '',
  {
    listenToStorageChanges: true
  }
)

export const trackBrowsingStorage = useLocalStorage('track-browsing', 'true', {
  listenToStorageChanges: true
})

export const trackInteractionsStorage = useLocalStorage(
  'track-interaction',
  'true',
  {
    listenToStorageChanges: true
  }
)

export const crawlWebsitesStorage = useLocalStorage(
  'crawl-websites',
  'true',
  {
    listenToStorageChanges: true
  }
)

export const emailNotificationsStorage = useLocalStorage(
  'email-notifications',
  'true',
  {
    listenToStorageChanges: true
  }
)

export const browserNotificationsStorage = useLocalStorage(
  'browser-notifications',
  'true',
  {
    listenToStorageChanges: true
  }
)

export const accessTokenStorage = useLocalStorage<string | null>('access-token', null, {
  listenToStorageChanges: true
})

// Function to clear all storage
export const clearAllStorage = () => {
  solanaPublicKeyStorage.value = ''
  trackBrowsingStorage.value = 'true'
  trackInteractionsStorage.value = 'true'
  crawlWebsitesStorage.value = 'true'
  emailNotificationsStorage.value = 'true'
  browserNotificationsStorage.value = 'true'
  accessTokenStorage.value = null
}
