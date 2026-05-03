import { apiRequest } from './client';
import type { Reservation } from '../types/models';

export function getReservations(): Promise<Reservation[]> {
  return apiRequest<Reservation[]>('/reservations');
}
