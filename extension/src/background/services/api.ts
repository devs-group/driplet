export class ApiService {
  private baseUrl: string

  public static readonly STATUS_UNAUTHORIZED = 'STATUS_UNAUTHORIZED'

  constructor() {
    this.baseUrl = 'http://localhost:9000' // TODO: replace with variable
  }

  private async getAccessToken(): Promise<string | null> {
    const key = 'accessToken'
    const token = await chrome.storage.local.get(key)
    if (!token) {
      console.error('No authentication token found')
      return null
    }
    return token[key]
  }

  public async post<T>({ endpoint, body }: { endpoint: string; body?: any }) {
    const token = await this.getAccessToken()
    if (!token) {
      throw new Error(`Token could not be found in local storage`)
    }

    const response = await fetch(`${this.baseUrl}${endpoint}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`
      },
      body: body ? JSON.stringify(body) : JSON.stringify({})
    })
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }

    return response.json() as T
  }

  public async put<T>({ endpoint, body }: { endpoint: string; body?: any }) {
    const token = await this.getAccessToken()
    if (!token) {
      throw new Error(`Token could not be found in local storage`)
    }

    const response = await fetch(`${this.baseUrl}${endpoint}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`
      },
      body: body ? JSON.stringify(body) : JSON.stringify({})
    })
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }

    return response.json() as T
  }

  public async get<T>({ endpoint }: { endpoint: string }) {
    const token = await this.getAccessToken()
    if (!token) {
      throw new Error(`Token could not be found in local storage`)
    }
    const response = await fetch(`${this.baseUrl}${endpoint}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`
      }
    })
    if (!response.ok) {
      console.error(response)
      if (response.status === 401) {
        throw new Error(ApiService.STATUS_UNAUTHORIZED)
      }
      throw new Error(`HTTP error! status: ${response.status}`)
    }

    return response.json() as T
  }
}
