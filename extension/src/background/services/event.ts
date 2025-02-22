import { ApiService } from './api'

export class EventApiService extends ApiService {
  constructor() {
    super()
  }

  public async sendEvent(data: any): Promise<Response> {
    return this.post({
      endpoint: '/api/v1/event',
      body: { data }
    })
  }
}
