import { sendMessage } from "webext-bridge";
import { getPermissions } from "../util"

// Collect detailed data to mimic behavior similar to web3 app like Swash
export function collectUserData() {
    const userData = {
      browserInfo: {
        userAgent: navigator.userAgent,
        platform: navigator.platform,
        language: navigator.language
      },
      screenInfo: {
        width: window.screen.width,
        height: window.screen.height,
        colorDepth: window.screen.colorDepth
      },
      connectionInfo: {
        online: navigator.onLine,
        connectionType: navigator.connection
          ? navigator.connection.effectiveType
          : 'unknown'
      },
      permissions: navigator.permissions
        ? getPermissions()
        : 'Permissions API not supported'
    }
    sendMessage('user-data', userData, 'background');
  }