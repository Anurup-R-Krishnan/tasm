import { apiRequest } from './client';

export function getUsers(): Promise<any[]> {
  return apiRequest<any[]>('/users');
}

export function getUserById(id: string | number): Promise<any> {
  return apiRequest<any>(`/users/${id}`);
}

export function deleteUser(id: string | number): Promise<void> {
  return apiRequest<void>(`/users/${id}`, { method: 'DELETE' });
}
