import { apiRequest } from './client';

export function getReservations(): Promise<any[]> {
  return apiRequest<any[]>('/reservations');
}
