import { apiRequest } from './client';

export function getLocations(): Promise<any[]> {
  return apiRequest<any[]>('/locations');
}

export function createLocation(payload: any): Promise<any> {
  return apiRequest<any>('/locations', {
    method: 'POST',
    body: JSON.stringify(payload),
  });
}
