import { sendMessage } from 'webext-bridge'
import { MESSAGE_TYPE_PAGE_INFO } from '../background/index'

function extractPageInfo(eventType = 'load') {
  const pageInfo = {
    event: eventType,
    website: window.location.hostname,
    path: window.location.pathname,
    timestamp: new Date().toISOString(),
    title: document.title,
    url: window.location.href,
    referrer: document.referrer,
    cookies: document.cookie,
    links: Array.from(document.querySelectorAll('a')).map((a) => a.href),
    images: Array.from(document.querySelectorAll('img')).map((img) => ({
      src: img.src,
      alt: img.alt
    })),
    forms: Array.from(document.forms).map((form) => ({
      id: form.id,
      action: form.action,
      method: form.method,
      inputs: Array.from(form.elements).map((input) => ({
        name: (input as HTMLInputElement).name,
        type: (input as HTMLInputElement).type,
        value: input
      }))
    })),
    timeSpentSeconds: Math.floor((Date.now() - pageLoadTime) / 1000) // Added time spent
  }
  // @ts-ignore
  sendMessage(MESSAGE_TYPE_PAGE_INFO, pageInfo, 'background')
}

const pageLoadTime = Date.now()

// Track page load
window.addEventListener('load', () => extractPageInfo('load'))
window.addEventListener('hashchange', () => extractPageInfo('navigation'))

// Track page exit
window.addEventListener('beforeunload', () => extractPageInfo('exit'))
window.addEventListener('visibilitychange', () => {
  if (document.visibilityState === 'hidden') {
    extractPageInfo('visibility_hidden')
  }
})
