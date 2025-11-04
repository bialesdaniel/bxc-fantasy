import { createClient } from '@connectrpc/connect'
import { createConnectTransport } from '@connectrpc/connect-web'
import { Health } from '../gen/health/v1/health_pb'

const transport = createConnectTransport({
  baseUrl: 'http://localhost:8080', // Will be changed for production
  useBinaryFormat: true,
})

export const apiClient = createClient(Health, transport)
