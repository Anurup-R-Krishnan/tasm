import { apiRequest } from './client';
import type { Asset, AssetEvent } from '../types/models';

export function getAssets(): Promise<Asset[]> {
  return apiRequest<Asset[]>('/assets');
}

export function getAssetById(id: string | number): Promise<Asset> {
  return apiRequest<Asset>(`/assets/${id}`);
}

export function getAssetByTag(tagId: string): Promise<Asset> {
  return apiRequest<Asset>(`/assets/by-tag/${tagId}`);
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

export function checkoutAsset(
  id: string | number,
  data: { userId: number; dueDate?: string; notes?: string; custodian?: string },
): Promise<Asset> {
  return apiRequest<Asset>(`/assets/${id}/checkout`, {
    method: 'POST',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function checkinAsset(
  id: string | number,
  data: { notes?: string; custodian?: string },
): Promise<Asset> {
  return apiRequest<Asset>(`/assets/${id}/checkin`, {
    method: 'POST',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function getAssetHistory(id: string | number): Promise<AssetEvent[]> {
  return apiRequest<AssetEvent[]>(`/assets/${id}/history`);
}
