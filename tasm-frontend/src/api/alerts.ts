import { apiRequest } from './client';

export function getAlerts(): Promise<any[]> {
  return apiRequest<any[]>('/alerts');
}
