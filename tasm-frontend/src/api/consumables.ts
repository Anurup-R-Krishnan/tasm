import { apiRequest } from './client';
import type { Consumable } from '../types/models';

export function getConsumables(): Promise<Consumable[]> {
  return apiRequest<Consumable[]>('/consumables');
}

export function getConsumableById(id: string | number): Promise<Consumable> {
  return apiRequest<Consumable>(`/consumables/${id}`);
}

export function createConsumable(data: Partial<Consumable>): Promise<Consumable> {
  return apiRequest<Consumable>('/consumables', {
    method: 'POST',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function updateConsumable(
  id: string | number,
  data: Partial<Consumable>,
): Promise<Consumable> {
  return apiRequest<Consumable>(`/consumables/${id}`, {
    method: 'PUT',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function deleteConsumable(id: string | number): Promise<{ message: string }> {
  return apiRequest<{ message: string }>(`/consumables/${id}`, {
    method: 'DELETE',
  });
}
