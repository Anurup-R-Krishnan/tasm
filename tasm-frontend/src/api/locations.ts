import { apiRequest } from './client';
import type { Location } from '../types/models';

export function getLocations(): Promise<Location[]> {
  return apiRequest<Location[]>('/locations');
}

export function createLocation(payload: Partial<Location>): Promise<Location> {
  return apiRequest<Location>('/locations', {
    method: 'POST',
    body: JSON.stringify(payload),
    headers: { 'Content-Type': 'application/json' },
  });
}
