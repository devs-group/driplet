// @ts-nocheck

/* eslint-disable no-console */
import { onMessage,sendMessage } from 'webext-bridge'
import { extractPageInfo } from './scripts/page-info'
import { trackLoadedScriptsAndDetectTechnologies } from './scripts/track-loaded-scripts-and-detected-technologies'
import { collectUserData } from './scripts/collect-user-data'
import { trackUserClicks } from './scripts/track-user-clicks'
import { trackBrowsingBehavior } from './scripts/track-browsing-behaviour'
import { trackUserInteractions } from './scripts/track-user-interaction'
import { trackPerformanceMetrics } from './scripts/track-performance-metrics'


console.info('[chrome-ext-mv3-starter...] Hello world from content script!!')

// Communication example: send previous tab title from background page
onMessage('tab-prev', ({ data }) => {
  console.log(`[chrome-ext-mv3-starter] Navigate from page "${data.title}"`)
})

// Track user interactions like mouse movement and keystrokes
trackUserInteractions()

// Track browsing behavior on page load and navigation
window.addEventListener('load', trackBrowsingBehavior())
window.addEventListener('popstate', trackBrowsingBehavior())
window.addEventListener('hashchange', trackBrowsingBehavior()) // For single page applications

window.addEventListener('load', collectUserData())
window.addEventListener('load', trackLoadedScriptsAndDetectTechnologies())
window.addEventListener('load', trackPerformanceMetrics(performance))
window.addEventListener('load', extractPageInfo())


document.addEventListener('click', trackUserClicks(event))