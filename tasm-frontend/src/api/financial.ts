import { apiRequest } from './client';

export function getLedgers(): Promise<any[]> {
  return apiRequest<any[]>('/ledgers');
}

export function getLeases(): Promise<any[]> {
  return apiRequest<any[]>('/leases');
}

export function getDepreciations(): Promise<any[]> {
  return apiRequest<any[]>('/depreciations');
}

export function getSoftwareLicenses(): Promise<any[]> {
  return apiRequest<any[]>('/software-licenses');
}

export function getContracts(): Promise<any[]> {
  return apiRequest<any[]>('/contracts');
}
