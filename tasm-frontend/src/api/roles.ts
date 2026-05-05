import { apiRequest } from './client';
import type { UserRole } from '../types/models';

export function getRoles(): Promise<UserRole[]> {
  return apiRequest<UserRole[]>('/roles');
}

export function getRoleById(id: string | number): Promise<UserRole> {
  return apiRequest<UserRole>(`/roles/${id}`);
}
