import { apiRequest } from './client';
import type { SystemUser, UserRole } from '../types/models';

export function getUsers(): Promise<SystemUser[]> {
  return apiRequest<SystemUser[]>('/users');
}

export function getUserById(id: string | number): Promise<SystemUser> {
  return apiRequest<SystemUser>(`/users/${id}`);
}

export function createUser(data: Partial<SystemUser>): Promise<SystemUser> {
  return apiRequest<SystemUser>('/users', {
    method: 'POST',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function updateUser(id: string | number, data: Partial<SystemUser>): Promise<SystemUser> {
  return apiRequest<SystemUser>(`/users/${id}`, {
    method: 'PUT',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function deleteUser(id: string | number): Promise<void> {
  return apiRequest<void>(`/users/${id}`, {
    method: 'DELETE',
  });
}

export function getRoles(): Promise<UserRole[]> {
  return apiRequest<UserRole[]>('/users/roles');
}

export function getRoleById(id: string | number): Promise<UserRole> {
  return apiRequest<UserRole>(`/users/roles/${id}`);
}

export function createRole(data: Partial<UserRole>): Promise<UserRole> {
  return apiRequest<UserRole>('/users/roles', {
    method: 'POST',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function updateRole(id: string | number, data: Partial<UserRole>): Promise<UserRole> {
  return apiRequest<UserRole>(`/users/roles/${id}`, {
    method: 'PUT',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function deleteRole(id: string | number): Promise<{ message: string }> {
  return apiRequest<{ message: string }>(`/users/roles/${id}`, {
    method: 'DELETE',
  });
}
