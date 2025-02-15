import { ApiService } from './api'

export class EventApiService extends ApiService {
  constructor() {
    super()
  }

  public async sendEvent(data: any): Promise<Response> {
    return this.post({
      endpoint: '/collect/client-event',
      body: { event: data }
    })
  }
}
