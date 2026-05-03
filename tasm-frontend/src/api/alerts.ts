import { apiRequest } from './client';
import type { SystemAlert } from '../types/models';

export function getAlerts(): Promise<SystemAlert[]> {
  return apiRequest<SystemAlert[]>('/alerts');
}
