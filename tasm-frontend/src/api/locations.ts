import { apiRequest } from './client';
import type { Location } from '../types/models';

export function getLocations(): Promise<Location[]> {
  return apiRequest<Location[]>('/locations');
}

export function getLocationById(id: string | number): Promise<Location> {
  return apiRequest<Location>(`/locations/${id}`);
}

export function createLocation(payload: Partial<Location>): Promise<Location> {
  return apiRequest<Location>('/locations', {
    method: 'POST',
    body: JSON.stringify(payload),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function updateLocation(id: string | number, data: Partial<Location>): Promise<Location> {
  return apiRequest<Location>(`/locations/${id}`, {
    method: 'PUT',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function deleteLocation(id: string | number): Promise<{ message: string }> {
  return apiRequest<{ message: string }>(`/locations/${id}`, {
    method: 'DELETE',
  });
}
