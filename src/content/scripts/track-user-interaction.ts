import { sendMessage } from 'webext-bridge'

export function trackUserInteractions() {
    document.addEventListener('keydown', (event) => {
      const interactionData = {
        key: event.key,
        code: event.code
      }
      sendMessage('user-interaction', interactionData, 'background')
    })
  }