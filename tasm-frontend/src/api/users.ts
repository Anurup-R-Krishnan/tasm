import { apiRequest } from './client';

export function getUsers(): Promise<any[]> {
  return apiRequest<any[]>('/users');
}

export function getUserById(id: string | number): Promise<any> {
  return apiRequest<any>(`/users/${id}`);
}
