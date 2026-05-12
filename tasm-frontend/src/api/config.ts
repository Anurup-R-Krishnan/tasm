import { apiRequest } from './client';

export interface SystemConfigPayload {
  [key: string]: string;
}

export function getSystemConfig(): Promise<SystemConfigPayload> {
  return apiRequest<SystemConfigPayload>('/config');
}

export function updateSystemConfig(configs: SystemConfigPayload): Promise<SystemConfigPayload> {
  return apiRequest<SystemConfigPayload>('/config', {
    method: 'PUT',
    body: { configs },
  });
}
