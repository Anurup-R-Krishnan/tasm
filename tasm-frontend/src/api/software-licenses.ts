import { apiRequest } from './client';

export function getSoftwareLicenses(): Promise<any[]> {
  return apiRequest<any[]>('/software-licenses');
}
