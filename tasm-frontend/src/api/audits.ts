import { apiRequest } from './client';
import type { AuditSession, AuditDiscrepancy } from '../types/models';

export function getAudits(params?: { status?: string }): Promise<AuditSession[]> {
  const qs = params?.status ? `?status=${encodeURIComponent(params.status)}` : '';
  return apiRequest<AuditSession[]>(`/audits${qs}`);
}

export function getAuditById(id: number | string): Promise<AuditSession> {
  return apiRequest<AuditSession>(`/audits/${id}`);
}

export function createAudit(data: Partial<AuditSession>): Promise<AuditSession> {
  return apiRequest<AuditSession>('/audits', {
    method: 'POST',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function updateAudit(
  id: number | string,
  data: Partial<AuditSession>,
): Promise<AuditSession> {
  return apiRequest<AuditSession>(`/audits/${id}`, {
    method: 'PUT',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function deleteAudit(id: number | string): Promise<{ message: string }> {
  return apiRequest<{ message: string }>(`/audits/${id}`, { method: 'DELETE' });
}

export interface ScanResult {
  result: 'found' | 'unregistered';
  message?: string;
  asset?: import('../types/models').Asset;
  discrepancy?: AuditDiscrepancy;
  locationMatch?: boolean;
  progress?: number;
}

export function scanAssetInAudit(
  auditId: number | string,
  data: { tagId: string; scannedLocation?: string },
): Promise<ScanResult> {
  return apiRequest<ScanResult>(`/audits/${auditId}/scan`, {
    method: 'POST',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function getDiscrepancies(params?: {
  auditSessionId?: number | string;
  status?: string;
  issueType?: string;
}): Promise<AuditDiscrepancy[]> {
  const qs = new URLSearchParams();
  if (params?.auditSessionId) qs.set('auditSessionId', String(params.auditSessionId));
  if (params?.status) qs.set('status', params.status);
  if (params?.issueType) qs.set('issueType', params.issueType);
  const q = qs.toString();
  return apiRequest<AuditDiscrepancy[]>(`/discrepancies${q ? '?' + q : ''}`);
}

export function getDiscrepancyById(id: number | string): Promise<AuditDiscrepancy> {
  return apiRequest<AuditDiscrepancy>(`/discrepancies/${id}`);
}

export function createDiscrepancy(data: Partial<AuditDiscrepancy>): Promise<AuditDiscrepancy> {
  return apiRequest<AuditDiscrepancy>('/discrepancies', {
    method: 'POST',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function resolveDiscrepancy(
  id: number | string,
  data: {
    action: 'confirm_location' | 'mark_lost' | 'register' | 'update_location' | 'dismiss';
    newLocation?: string;
    notes?: string;
  },
): Promise<AuditDiscrepancy> {
  return apiRequest<AuditDiscrepancy>(`/discrepancies/${id}/resolve`, {
    method: 'POST',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}
