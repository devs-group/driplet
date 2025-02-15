import { ApiService } from './api'

export interface GetUserCreditsResponse {
  credits: number
}

export class UserApiService extends ApiService {
  constructor() {
    super()
  }

  public async saveUser(): Promise<Response> {
    return this.post({
      endpoint: '/auth/user'
    })
  }

  public async savePublicKey(publicKey: string): Promise<Response> {
    return this.post({
      endpoint: '/auth/user',
      body: {
        publicKey
      }
    })
  }

  public async getUserCredits(): Promise<GetUserCreditsResponse> {
    return await this.get<GetUserCreditsResponse>({
      endpoint: '/auth/user/credits'
    })
  }
}
