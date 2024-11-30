import { sendMessage, onMessage } from 'webext-bridge'
import { apiService } from './services/api'

chrome.runtime.onInstalled.addListener((): void => {
  // eslint-disable-next-line no-console
  console.log('Extension installed')
})

chrome.runtime.onMessage.addListener((request, sender, sendResponse) => {
  if (request.action === 'login') {
    const clientId = chrome.runtime.getManifest().oauth2?.client_id
    const redirectUri = chrome.identity.getRedirectURL()
    const scopes = ['openid', 'email', 'profile'].join(' ')

    const authUrl =
      `https://accounts.google.com/o/oauth2/v2/auth?` +
      `client_id=${clientId}&` +
      `redirect_uri=${encodeURIComponent(redirectUri)}&` +
      `response_type=token&` +
      `scope=${encodeURIComponent(scopes)}`

    chrome.identity.launchWebAuthFlow(
      {
        url: authUrl,
        interactive: true
      },
      (redirectUrl) => {
        if (chrome.runtime.lastError || !redirectUrl) {
          sendResponse({
            success: false,
            error: chrome.runtime.lastError?.message || 'Authentication failed'
          })
          return
        }

        // Extract the access token from the redirect URL
        const params = new URLSearchParams(
          new URL(redirectUrl).hash.substring(1)
        )
        const accessToken = params.get('access_token')

        if (!accessToken) {
          sendResponse({ success: false, error: 'No access token received' })
          return
        }

        // Store the access token securely
        chrome.storage.local.set({ accessToken }, () => {
          sendResponse({ success: true, accessToken })
        })
      }
    )

    // Keep the message channel open for sendResponse
    return true
  }
})

let previousTabId = 0

// communication example: send previous tab title from background page
// see shim.d.ts for type decleration
chrome.tabs.onActivated.addListener(async ({ tabId }) => {
  if (!previousTabId) {
    previousTabId = tabId
    return
  }
  const tab = await chrome.tabs.get(previousTabId)
  previousTabId = tabId
  if (!tab) return

  // eslint-disable-next-line no-console
  console.log('previous tab', tab)
  sendMessage(
    'tab-prev',
    { title: tab.title },
    { context: 'content-script', tabId }
  )
})

onMessage('get-current-tab', async () => {
  try {
    const tab = await chrome.tabs.get(previousTabId)
    return {
      title: tab?.id
    }
  } catch {
    return {
      title: undefined
    }
  }
})

onMessage('page-info', async ({ data }) => {
  console.log('page-info', data)
  await apiService.makeRequest(data)
})

onMessage('performance-metrics', async ({ data }) => {
  console.log('performance-metrics', data)
  await apiService.makeRequest(data)
})

onMessage('browsing-behaviour', async ({ data }) => {
  console.log('browsing-behaviour', data)
  await apiService.makeRequest(data)
})

onMessage('loaded-scripts', async ({ data }) => {
  console.log('loaded-scripts', data)
  await apiService.makeRequest(data)
})

onMessage('user-clicks', async ({ data }) => {
  console.log('user-clicks', data)
  await apiService.makeRequest(data)
})

onMessage('user-interaction', async ({ data }) => {
  console.log('user-interaction', data)
  await apiService.makeRequest(data)
})

onMessage('user-data', async ({ data }) => {
  console.log('user-data', data)
  await apiService.makeRequest(data)
})
