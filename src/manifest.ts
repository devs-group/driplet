import type { Manifest } from 'webextension-polyfill'
import pkg from '../package.json'
import { IS_DEV, PORT } from '../scripts/utils'
interface ExtendedWebExtensionManifest extends Manifest.WebExtensionManifest {
  oauth2?: {
    client_id: string;
    scopes: string[];
  };
}

export async function getManifest(): Promise<ExtendedWebExtensionManifest> {
  // update this file to update this manifest.json
  // can also be conditional based on your need

  const manifest: ExtendedWebExtensionManifest= {
    manifest_version: 3,
    name: pkg.displayName || pkg.name,
    
    version: pkg.version,
    description: pkg.description,
    action: {
      default_icon: './assets/icon-512.png',
      default_popup: './popup/index.html'
    },
    options_ui: {
      page: './options/index.html',
      open_in_tab: true
    },
    background: {
      service_worker: 'background.js'
    },
    content_scripts: [
      {
        matches: ['<all_urls>'],
        js: ['./content/index.global.js']
      }
    ],
    icons: {
      16: './assets/icon-512.png',
      48: './assets/icon-512.png',
      128: './assets/icon-512.png'
    },
    permissions: [
      'scripting',
      'identity',
      'storage',
      'tabs',
      'contextMenus',
      'webRequest',
      'storage',
      'activeTab'
    ],
    host_permissions: [
      'http://*/*',
      'https://*/*',
      'http://localhost:1991/*'
    ],
    oauth2: {
      client_id: '34932676923-luth1cmkkhoqhq3kaoealqj2poosaln4.apps.googleusercontent.com',
      scopes: [
        'openid',
        'email',
        'profile',
        'https://www.googleapis.com/auth/calendar',
        'https://www.googleapis.com/auth/userinfo.email',
        'https://www.googleapis.com/auth/userinfo.profile'
      ]
    },
    web_accessible_resources: [
      {
        resources: ['core/inpage_scripts/sdk_page_script.js', './assets/*'],
        matches: ['<all_urls>']
      }
    ],
    content_security_policy: {}
  }

  if (IS_DEV) {
    // this is required on dev for Vite script to load
    manifest.content_security_policy = {
      extension_pages: `script-src 'self' http://localhost:${PORT}; object-src 'self'`
    }
  }

  return manifest
}
