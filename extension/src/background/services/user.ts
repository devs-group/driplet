import { ApiService } from './api'

export interface GetUserResponse {
  credits: number
  email: string
  public_key: string
}

export class UserApiService extends ApiService {
  constructor() {
    super()
  }

  public async PUT_UpdateUsersPublicKey(publicKey: string): Promise<Response> {
    return this.put({
      endpoint: '/api/v1/user/public-key',
      body: {
        public_key: publicKey
      }
    })
  }

  public async GET_User(): Promise<GetUserResponse> {
    return await this.get<GetUserResponse>({
      endpoint: '/api/v1/user'
    })
  }
}
