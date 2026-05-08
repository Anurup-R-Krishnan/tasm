import { apiRequest } from './client';
import type { SystemAlert } from '../types/models';

export function getAlerts(params?: {
  unread?: boolean | undefined;
  type?: string | undefined;
}): Promise<SystemAlert[]> {
  const qs = new URLSearchParams();
  if (params?.unread) qs.set('unread', 'true');
  if (params?.type) qs.set('type', params.type);
  const q = qs.toString();
  return apiRequest<SystemAlert[]>(`/alerts${q ? '?' + q : ''}`);
}

export function updateAlert(id: string | number, data: Partial<SystemAlert>): Promise<SystemAlert> {
  return apiRequest<SystemAlert>(`/alerts/${id}`, {
    method: 'PUT',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function deleteAlert(id: string | number): Promise<{ message: string }> {
  return apiRequest<{ message: string }>(`/alerts/${id}`, { method: 'DELETE' });
}

export function markAllAlertsRead(): Promise<{ message: string }> {
  return apiRequest<{ message: string }>('/alerts/mark-all-read', {
    method: 'POST',
  });
}

export function generateAlerts(): Promise<{ message: string }> {
  return apiRequest<{ message: string }>('/alerts/generate', {
    method: 'POST',
  });
}
