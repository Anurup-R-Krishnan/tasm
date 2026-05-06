import { apiRequest } from './client';
import type { SystemUser } from '../types/models';

export interface CreateUserPayload {
  employeeId: string;
  name: string;
  email: string;
  password: string;
  department: string;
  role: string;
  status?: string;
}

export interface UpdateUserPayload {
  employeeId?: string;
  name?: string;
  email?: string;
  department?: string;
  role?: string;
  status?: string;
  password?: string;
}

export function getUsers(): Promise<SystemUser[]> {
  return apiRequest<SystemUser[]>('/users');
}

export function getUserById(id: string | number): Promise<SystemUser> {
  return apiRequest<SystemUser>(`/users/${id}`);
}

export function createUser(payload: CreateUserPayload): Promise<SystemUser> {
  return apiRequest<SystemUser>('/users', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload),
  });
}

export function updateUser(id: string | number, payload: UpdateUserPayload): Promise<SystemUser> {
  return apiRequest<SystemUser>(`/users/${id}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload),
  });
}

export function deleteUser(id: string | number): Promise<void> {
  return apiRequest<void>(`/users/${id}`, { method: 'DELETE' });
}
