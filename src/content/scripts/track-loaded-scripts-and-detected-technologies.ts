import { sendMessage } from "webext-bridge";

// Track and log loaded scripts and detect technologies used on the page
export function trackLoadedScriptsAndDetectTechnologies() {
    const loadedScripts = [...document.scripts].map(
      (script) => script.src || 'inline script'
    )
  
    const technologies = {
      angular: !!window.angular,
      react: !!document.querySelector('[data-reactroot], [data-reactid]'),
      vue: !!document.querySelector('[data-vue-meta]') || !!window.Vue,
      jquery: !!window.jQuery,
      nextjs: !!window.__NEXT_DATA__,
      nuxt: !!window.__NUXT__,
      stripe: !!window.Stripe
    }
  
    sendMessage('loaded-scripts', { scripts: loadedScripts, technologies }, 'background');
  }