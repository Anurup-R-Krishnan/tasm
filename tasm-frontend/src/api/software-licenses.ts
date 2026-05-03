import { apiRequest } from './client';
import type { SoftwareLicense } from '../types/models';

export function getSoftwareLicenses(): Promise<SoftwareLicense[]> {
  return apiRequest<SoftwareLicense[]>('/software-licenses');
}
