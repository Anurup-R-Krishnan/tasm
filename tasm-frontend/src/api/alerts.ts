import { apiRequest } from './client';
import type { SystemAlert } from '../types/models';

export function getAlerts(): Promise<SystemAlert[]> {
  return apiRequest<SystemAlert[]>('/alerts');
}

export function updateAlert(id: string | number, data: Partial<SystemAlert>): Promise<SystemAlert> {
  return apiRequest<SystemAlert>(`/alerts/${id}`, {
    method: 'PUT',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function deleteAlert(id: string | number): Promise<{ message: string }> {
  return apiRequest<{ message: string }>(`/alerts/${id}`, {
    method: 'DELETE',
  });
}
