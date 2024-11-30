
interface ApiResponse {
  success: boolean
  error?: string
}

class ApiService {
  private baseUrl: string

  constructor() {
    this.baseUrl = 'http://localhost:1991' // replace this later
  }



  public async makeRequest(data: any): Promise<ApiResponse> {

    const token = await new Promise<string | null>((resolve) => {
      chrome.storage.local.get(['accessToken'], (result) => {
        resolve(result.accessToken || null);
      });
    });

    if (!token) {
      console.error('No authentication token found')
      return { success: false, error: 'No authentication token found' }
    }

    try {
      const response = await fetch(`${this.baseUrl}/collect/client-event`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`
        },
       body: JSON.stringify({event: data})
      })
      console.log('response', response)
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`)
      }

      return { success: true }
    } catch (error) {
      console.error('API request failed:', error)
      return { success: false, error: error.message }
    }
  }

  
}

export const apiService = new ApiService()
