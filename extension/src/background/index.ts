import { onMessage } from 'webext-bridge'
import { EventApiService } from './services/event'

const eventApiService = new EventApiService()

export const MESSAGE_TYPE_PAGE_INFO = 'page-info'
onMessage(MESSAGE_TYPE_PAGE_INFO, async ({ data }) => {
  await eventApiService.sendEvent(data)
})
