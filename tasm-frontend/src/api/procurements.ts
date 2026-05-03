import { apiRequest } from './client';

export interface Procurement {
  id: number;
  title: string;
  status: string;
  [key: string]: any;
}

export function getProcurements(): Promise<Procurement[]> {
  return apiRequest<Procurement[]>('/procurements');
}

export function getProcurementById(id: string | number): Promise<Procurement> {
  return apiRequest<Procurement>(`/procurements/${id}`);
}

export function createProcurement(data: Partial<Procurement>): Promise<Procurement> {
  return apiRequest<Procurement>('/procurements', {
    method: 'POST',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function updateProcurement(id: string | number, data: Partial<Procurement>): Promise<Procurement> {
  return apiRequest<Procurement>(`/procurements/${id}`, {
    method: 'PUT',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function deleteProcurement(id: string | number): Promise<{ message: string }> {
  return apiRequest<{ message: string }>(`/procurements/${id}`, {
    method: 'DELETE',
  });
}
