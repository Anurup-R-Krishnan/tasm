import { apiRequest } from './client';
import type {
  LedgerEntry,
  LeaseAgreement,
  DepreciationSchedule,
  MaintenanceContract,
  SoftwareLicense,
} from '../types/models';

export function getLedgers(): Promise<LedgerEntry[]> {
  return apiRequest<LedgerEntry[]>('/ledgers');
}

export function getLeases(): Promise<LeaseAgreement[]> {
  return apiRequest<LeaseAgreement[]>('/leases');
}

export function getDepreciations(): Promise<DepreciationSchedule[]> {
  return apiRequest<DepreciationSchedule[]>('/depreciations');
}

export function getSoftwareLicenses(): Promise<SoftwareLicense[]> {
  return apiRequest<SoftwareLicense[]>('/software-licenses');
}

export function getContracts(): Promise<MaintenanceContract[]> {
  return apiRequest<MaintenanceContract[]>('/contracts');
}
