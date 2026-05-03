import { apiRequest } from './client';

export function getAudits(): Promise<any[]> {
  return apiRequest<any[]>('/audits');
}

export function getDiscrepancies(): Promise<any[]> {
  return apiRequest<any[]>('/discrepancies');
}
