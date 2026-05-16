import { apiRequest } from './client';
import type { SystemUser } from '../types/models';

export function getUsers(): Promise<SystemUser[]> {
  return apiRequest<SystemUser[]>('/users');
}

export function getUserById(id: string | number): Promise<SystemUser> {
  return apiRequest<SystemUser>(`/users/${id}`);
}

export function deleteUser(id: string | number): Promise<void> {
  return apiRequest<void>(`/users/${id}`, {
    method: 'DELETE',
  });
}
