import { apiRequest } from './client';
import type { Reservation } from '../types/models';

export function getReservations(): Promise<Reservation[]> {
  return apiRequest<Reservation[]>('/reservations');
}

export function getReservationById(id: string | number): Promise<Reservation> {
  return apiRequest<Reservation>(`/reservations/${id}`);
}

export function createReservation(data: Partial<Reservation>): Promise<Reservation> {
  return apiRequest<Reservation>('/reservations', {
    method: 'POST',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function updateReservation(
  id: string | number,
  data: Partial<Reservation>,
): Promise<Reservation> {
  return apiRequest<Reservation>(`/reservations/${id}`, {
    method: 'PUT',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function deleteReservation(id: string | number): Promise<{ message: string }> {
  return apiRequest<{ message: string }>(`/reservations/${id}`, {
    method: 'DELETE',
  });
}
