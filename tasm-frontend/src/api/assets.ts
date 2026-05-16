import { apiRequest } from './client';
import type { Asset, AssetEvent, AssetTransfer, DepreciationYearRow } from '../types/models';

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
  data: {
    userId: number;
    dueDate?: string | undefined;
    notes?: string | undefined;
    custodian?: string | undefined;
  },
): Promise<Asset> {
  return apiRequest<Asset>(`/assets/${id}/checkout`, {
    method: 'POST',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function checkinAsset(
  id: string | number,
  data: { notes?: string | undefined; custodian?: string | undefined },
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

export function transferAsset(
  id: string | number,
  data: {
    newLocation?: string | undefined;
    newCustodian?: string | undefined;
    notes?: string | undefined;
  },
): Promise<Asset> {
  return apiRequest<Asset>(`/assets/${id}/transfer`, {
    method: 'POST',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function retireAsset(
  id: string | number,
  data: { reason: string; notes?: string },
): Promise<Asset> {
  return apiRequest<Asset>(`/assets/${id}/retire`, {
    method: 'POST',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function disposeAsset(
  id: string | number,
  data: {
    disposalMethod: string;
    residualValue: number;
    environmentalCompliance: boolean;
    dataWipingConfirmed: boolean;
    notes?: string;
  },
): Promise<Asset> {
  return apiRequest<Asset>(`/assets/${id}/dispose`, {
    method: 'POST',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function getAssetDepreciation(id: string | number): Promise<DepreciationYearRow[]> {
  return apiRequest<DepreciationYearRow[]>(`/assets/${id}/depreciation`);
}

export function getAssetTransfers(id: string | number): Promise<AssetTransfer[]> {
  return apiRequest<AssetTransfer[]>(`/assets/${id}/transfers`);
}

export function revertTransfer(
  assetId: string | number,
  transferId: string | number,
): Promise<AssetTransfer> {
  return apiRequest<AssetTransfer>(`/assets/${assetId}/transfers/${transferId}/revert`, {
    method: 'POST',
  });
}

export function uploadReceipt(id: string | number, file: File): Promise<{ path: string }> {
  const formData = new FormData();
  formData.append('receipt', file);
  return apiRequest<{ path: string }>(`/assets/${id}/receipt`, {
    method: 'POST',
    body: formData,
  });
}

export function getReceiptUrl(id: string | number): string {
  const base = import.meta.env['VITE_API_BASE_URL']?.replace(/\/$/, '') ?? '/api';
  return `${base}/assets/${id}/receipt`;
}

export function getAssetLifecycle(id: string | number): Promise<any> {
  return apiRequest<any>(`/assets/${id}/lifecycle`);
}
