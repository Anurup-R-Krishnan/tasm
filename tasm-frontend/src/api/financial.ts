import { apiRequest } from './client';
import type {
  LedgerEntry,
  LeaseAgreement,
  DepreciationSchedule,
  SoftwareLicense,
  MaintenanceContract,
} from '../types/models';

export function getLedgers(): Promise<LedgerEntry[]> {
  return apiRequest<LedgerEntry[]>('/financial/ledgers');
}

export function getLedgerById(id: string | number): Promise<LedgerEntry> {
  return apiRequest<LedgerEntry>(`/financial/ledgers/${id}`);
}

export function createLedger(data: Partial<LedgerEntry>): Promise<LedgerEntry> {
  return apiRequest<LedgerEntry>('/financial/ledgers', {
    method: 'POST',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function updateLedger(
  id: string | number,
  data: Partial<LedgerEntry>,
): Promise<LedgerEntry> {
  return apiRequest<LedgerEntry>(`/financial/ledgers/${id}`, {
    method: 'PUT',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function deleteLedger(id: string | number): Promise<{ message: string }> {
  return apiRequest<{ message: string }>(`/financial/ledgers/${id}`, {
    method: 'DELETE',
  });
}

export function getLeases(): Promise<LeaseAgreement[]> {
  return apiRequest<LeaseAgreement[]>('/financial/leases');
}

export function getLeaseById(id: string | number): Promise<LeaseAgreement> {
  return apiRequest<LeaseAgreement>(`/financial/leases/${id}`);
}

export function createLease(data: Partial<LeaseAgreement>): Promise<LeaseAgreement> {
  return apiRequest<LeaseAgreement>('/financial/leases', {
    method: 'POST',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function updateLease(
  id: string | number,
  data: Partial<LeaseAgreement>,
): Promise<LeaseAgreement> {
  return apiRequest<LeaseAgreement>(`/financial/leases/${id}`, {
    method: 'PUT',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function deleteLease(id: string | number): Promise<{ message: string }> {
  return apiRequest<{ message: string }>(`/financial/leases/${id}`, {
    method: 'DELETE',
  });
}

export function getDepreciations(): Promise<DepreciationSchedule[]> {
  return apiRequest<DepreciationSchedule[]>('/financial/depreciations');
}

export function getDepreciationById(id: string | number): Promise<DepreciationSchedule> {
  return apiRequest<DepreciationSchedule>(`/financial/depreciations/${id}`);
}

export function createDepreciation(
  data: Partial<DepreciationSchedule>,
): Promise<DepreciationSchedule> {
  return apiRequest<DepreciationSchedule>('/financial/depreciations', {
    method: 'POST',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function updateDepreciation(
  id: string | number,
  data: Partial<DepreciationSchedule>,
): Promise<DepreciationSchedule> {
  return apiRequest<DepreciationSchedule>(`/financial/depreciations/${id}`, {
    method: 'PUT',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function deleteDepreciation(id: string | number): Promise<{ message: string }> {
  return apiRequest<{ message: string }>(`/financial/depreciations/${id}`, {
    method: 'DELETE',
  });
}

export function getSoftwareLicenses(): Promise<SoftwareLicense[]> {
  return apiRequest<SoftwareLicense[]>('/financial/software-licenses');
}

export function getSoftwareLicenseById(id: string | number): Promise<SoftwareLicense> {
  return apiRequest<SoftwareLicense>(`/financial/software-licenses/${id}`);
}

export function createSoftwareLicense(data: Partial<SoftwareLicense>): Promise<SoftwareLicense> {
  return apiRequest<SoftwareLicense>('/financial/software-licenses', {
    method: 'POST',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function updateSoftwareLicense(
  id: string | number,
  data: Partial<SoftwareLicense>,
): Promise<SoftwareLicense> {
  return apiRequest<SoftwareLicense>(`/financial/software-licenses/${id}`, {
    method: 'PUT',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function deleteSoftwareLicense(id: string | number): Promise<{ message: string }> {
  return apiRequest<{ message: string }>(`/financial/software-licenses/${id}`, {
    method: 'DELETE',
  });
}

export function getContracts(): Promise<MaintenanceContract[]> {
  return apiRequest<MaintenanceContract[]>('/maintenance/contracts');
}
