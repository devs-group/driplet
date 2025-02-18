import { UserApiService } from '~/background/services/user'

export interface GoogleUserProfile {
  name: string
  email: string
  picture: string
}

export type UserProfile = GoogleUserProfile & { credits: number }

/**
 * Fetches the user profile from Google's API using the access token.
 * @param accessToken - The OAuth2 access token.
 * @returns The user profile if successful.
 */
export async function fetchGoogleUserProfile(
  accessToken: string
): Promise<GoogleUserProfile | undefined> {
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
    const profile: GoogleUserProfile = {
      name: data.name,
      email: data.email,
      picture: data.picture
    }

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
    return userProfile
  } catch (err) {
    throw new Error('Failed to retrieve user profile', { cause: err })
  }
}

function getClientId() {
  return chrome.runtime.getManifest().oauth2?.client_id
}

function getRedirectURL() {
  return chrome.identity.getRedirectURL()
}

function getAuthURL() {
  const clientId = getClientId()
  const redirectUri = getRedirectURL()
  const scopes = ['openid', 'email', 'profile'].join(' ')

  return (
    `https://accounts.google.com/o/oauth2/v2/auth?` +
    `client_id=${clientId}&` +
    `redirect_uri=${encodeURIComponent(redirectUri)}&` +
    `response_type=token&` +
    `scope=${encodeURIComponent(scopes)}`
  )
}

function getAccessTokenFromRedirectURL(redirectUrl: string) {
  if (chrome.runtime.lastError || !redirectUrl) {
    throw Error(chrome.runtime.lastError?.message)
  }
  const params = new URLSearchParams(new URL(redirectUrl).hash.substring(1))
  return params.get('access_token')
}

/**
 * Initiates the login flow using Chrome's message passing API.
 * @returns The access token if login succeeds.
 */
export async function handleLogin() {
  return new Promise<string>((resolve, reject) => {
    try {
      chrome.identity.launchWebAuthFlow(
        {
          url: getAuthURL(),
          interactive: true
        },
        async (redirectUrl) => {
          const accessToken = getAccessTokenFromRedirectURL(redirectUrl ?? '')
          if (accessToken) {
            await chrome.storage.local.set({ accessToken })
            const profile = await fetchGoogleUserProfile(accessToken)
            await chrome.storage.local.set({ userProfile: profile })
            await signUpUser()
            resolve(accessToken)
          } else {
            reject('Login failed. No access token returned.')
          }
        }
      )
    } catch (err) {
      reject(`login failed: ${err}`)
    }
  })
}

/**
 * Logs the user out by clearing stored data and revoking the token.
 */
export async function handleLogout(): Promise<void> {
  try {
    const { accessToken } = await chrome.storage.local.get('accessToken')

    if (accessToken) {
      await revokeTokenFromGoogle(accessToken)
    }

    await chrome.storage.local.remove(['accessToken', 'userProfile'])
    console.info('User logged out successfully')
  } catch (err) {
    throw new Error('Logout failed', { cause: err })
  }
}

/**
 * saves the current user on the backend
 * @returns The response
 */
export async function signUpUser() {
  const userApiService = new UserApiService()
  return userApiService.GET_User()
}

/**
 * Revokes the access token by calling Google's revocation endpoint.
 * @param token - The OAuth2 access token to be revoked.
 */
async function revokeTokenFromGoogle(token: string): Promise<void> {
  try {
    const response = await fetch(
      `https://accounts.google.com/o/oauth2/revoke?token=${token}`
    )
    if (!response.ok) {
      throw new Error('Failed to revoke token')
    }
  } catch (err) {
    throw new Error('Failed to revoke token', { cause: err })
  }
}
