import { apiRequest } from './client';

export interface Consumable {
  id: number;
  name: string;
  category: string;
  location: string;
  currentStock: number;
  reorderLevel: number;
}

export function getConsumables(): Promise<Consumable[]> {
  return apiRequest<Consumable[]>('/consumables');
}
