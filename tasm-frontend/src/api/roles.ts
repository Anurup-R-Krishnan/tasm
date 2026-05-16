import { apiRequest } from './client';
import type { UserRole } from '../types/models';

export function getRoles(): Promise<UserRole[]> {
  return apiRequest<UserRole[]>('/roles');
}

export function getRoleById(id: string | number): Promise<UserRole> {
  return apiRequest<UserRole>(`/roles/${id}`);
}

export function createRole(data: Partial<UserRole>): Promise<UserRole> {
  return apiRequest<UserRole>('/roles', {
    method: 'POST',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function updateRole(id: string | number, data: Partial<UserRole>): Promise<UserRole> {
  return apiRequest<UserRole>(`/roles/${id}`, {
    method: 'PUT',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function deleteRole(id: string | number): Promise<{ message: string }> {
  return apiRequest<{ message: string }>(`/roles/${id}`, {
    method: 'DELETE',
  });
}
