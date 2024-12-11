export interface UserProfile {
  name: string
  email: string
  picture: string
}

export interface LoginResponse {
  success: boolean
  accessToken?: string
  error?: string
}

/**
 * Fetches the user profile from Google's API using the access token.
 * @param accessToken - The OAuth2 access token.
 * @returns The user profile if successful.
 */
export async function fetchUserProfile(
  accessToken: string
): Promise<UserProfile | undefined> {
  try {
    const response = await fetch(
      'https://www.googleapis.com/oauth2/v2/userinfo',
      {
        headers: {
          Authorization: `Bearer ${accessToken}`
        }
      }
    )

    if (!response.ok) {
      throw new Error('Failed to fetch user profile')
    }

    const data = await response.json()
    const profile: UserProfile = {
      name: data.name,
      email: data.email,
      picture: data.picture
    }

    await chrome.storage.local.set({ userProfile: profile })
    return profile
  } catch (err) {
    throw new Error('Failed to fetch user profile', { cause: err })
  }
}

/**
 * Retrieves the user profile from Chrome storage.
 * @returns The stored user profile, or undefined if not found.
 */
export async function getUserProfile(): Promise<UserProfile | undefined> {
  try {
    const { userProfile } = await chrome.storage.local.get('userProfile')
    if (!userProfile) {
      console.info('No user profile found in storage')
      return undefined
    }
    return userProfile as UserProfile
  } catch (err) {
    throw new Error('Failed to retrieve user profile', { cause: err })
  }
}

/**
 * Initiates the login flow using Chrome's message passing API.
 * @returns The access token if login succeeds.
 */
export async function handleLogin(): Promise<string | undefined> {
  try {
    const response = await _sendMessage<LoginResponse>({ action: 'login' })

    if (response.success && response.accessToken) {
      await fetchUserProfile(response.accessToken)
      await chrome.storage.local.set({ accessToken: response.accessToken })
      return response.accessToken
    } else {
      throw new Error('Login failed. No access token returned.')
    }
  } catch (err) {
    throw new Error('Login failed', { cause: err })
  }
}

/**
 * Logs the user out by clearing stored data and revoking the token.
 */
export async function handleLogout(): Promise<void> {
  try {
    const { accessToken } = await chrome.storage.local.get('accessToken')

    if (accessToken) {
      await revokeToken(accessToken) // Revoke the token from Google
    }

    await chrome.storage.local.remove(['accessToken', 'userProfile'])
    console.info('User logged out successfully')
  } catch (err) {
    throw new Error('Logout failed', { cause: err })
  }
}

/**
 * Revokes the access token by calling Google's revocation endpoint.
 * @param token - The OAuth2 access token to be revoked.
 */
async function revokeToken(token: string): Promise<void> {
  try {
    const response = await fetch(
      `https://accounts.google.com/o/oauth2/revoke?token=${token}`
    )
    if (!response.ok) {
      throw new Error('Failed to revoke token')
    }
    console.info('Token successfully revoked')
  } catch (err) {
    throw new Error('Failed to revoke token', { cause: err })
  }
}

/**
 * Sends a message using Chrome's runtime API with type safety.
 * @param message - The message to be sent.
 * @returns The response from the background script.
 */
function _sendMessage<T>(message: object): Promise<T> {
  return new Promise((resolve, reject) => {
    chrome.runtime.sendMessage(message, (response: T) => {
      if (chrome.runtime.lastError) {
        return reject(new Error(chrome.runtime.lastError.message))
      }
      resolve(response)
    })
  })
}
