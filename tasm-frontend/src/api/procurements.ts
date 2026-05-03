import { apiRequest } from './client';
import type { ProcurementRequest } from '../types/models';

export function getProcurements(): Promise<ProcurementRequest[]> {
  return apiRequest<ProcurementRequest[]>('/procurements');
}

export function getProcurementById(id: string | number): Promise<ProcurementRequest> {
  return apiRequest<ProcurementRequest>(`/procurements/${id}`);
}

export function createProcurement(data: Partial<ProcurementRequest>): Promise<ProcurementRequest> {
  return apiRequest<ProcurementRequest>('/procurements', {
    method: 'POST',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function updateProcurement(
  id: string | number,
  data: Partial<ProcurementRequest>,
): Promise<ProcurementRequest> {
  return apiRequest<ProcurementRequest>(`/procurements/${id}`, {
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
