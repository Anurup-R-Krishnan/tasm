import { apiRequest } from './client';
import type { AuditSession, AuditDiscrepancy } from '../types/models';

export function getAudits(): Promise<AuditSession[]> {
  return apiRequest<AuditSession[]>('/audits');
}

export function getDiscrepancies(): Promise<AuditDiscrepancy[]> {
  return apiRequest<AuditDiscrepancy[]>('/discrepancies');
}
