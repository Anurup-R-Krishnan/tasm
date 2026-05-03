import { apiRequest } from './client';

export interface Asset {
  id: number;
  name: string;
  tagId: string;
  category: string;
  status: string;
  custodian: string;
  location: string;
  warrantyStatus: string;
  warrantyExpiry: string;
  value?: number;
  [key: string]: any;
}

export function getAssets(): Promise<Asset[]> {
  return apiRequest<Asset[]>('/assets');
}

export function getAssetById(id: string | number): Promise<Asset> {
  return apiRequest<Asset>(`/assets/${id}`);
}

export function createAsset(data: Partial<Asset>): Promise<Asset> {
  return apiRequest<Asset>('/assets', {
    method: 'POST',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function updateAsset(id: string | number, data: Partial<Asset>): Promise<Asset> {
  return apiRequest<Asset>(`/assets/${id}`, {
    method: 'PUT',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function deleteAsset(id: string | number): Promise<{ message: string }> {
  return apiRequest<{ message: string }>(`/assets/${id}`, {
    method: 'DELETE',
  });
}
