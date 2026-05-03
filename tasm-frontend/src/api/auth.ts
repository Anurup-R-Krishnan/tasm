import { apiRequest } from './client';

export interface LoginCredentials {
  email: string;
  password?: string;
}

export interface AuthResponse {
  token: string;
  user: any;
}

export function login(credentials: LoginCredentials): Promise<AuthResponse> {
  // Simulating a POST to /auth/login for now
  return apiRequest<AuthResponse>('/auth/login', {
    method: 'POST',
    body: JSON.stringify(credentials),
  });
}

export function getMe(): Promise<any> {
  return apiRequest<any>('/auth/me');
}
