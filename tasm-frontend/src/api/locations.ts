import { apiRequest } from './client';

export function getLocations(): Promise<any[]> {
  return apiRequest<any[]>('/locations');
}
