import { apiRequest } from './client';
import type { SystemUser } from '../types/models';

export interface LoginCredentials {
  email: string;
  password?: string;
}

export interface AuthResponse {
  token: string;
  user: SystemUser;
}

export function login(credentials: LoginCredentials): Promise<AuthResponse> {
  return apiRequest<AuthResponse>('/auth/login', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(credentials),
  });
}

export function getMe(): Promise<SystemUser> {
  return apiRequest<SystemUser>('/auth/me');
}

export interface RegisterCredentials {
  name: string;
  email: string;
  password?: string;
  employeeId?: string;
  department?: string;
}

export function register(credentials: RegisterCredentials): Promise<AuthResponse> {
  return apiRequest<AuthResponse>('/auth/register', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(credentials),
  });
}
