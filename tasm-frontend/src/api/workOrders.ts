import { apiRequest } from './client';
import type { WorkOrder, MaintenanceContract } from '../types/models';

export function getWorkOrders(): Promise<WorkOrder[]> {
  return apiRequest<WorkOrder[]>('/maintenance/work-orders');
}

export function getWorkOrderById(id: string | number): Promise<WorkOrder> {
  return apiRequest<WorkOrder>(`/maintenance/work-orders/${id}`);
}

export function createWorkOrder(data: Partial<WorkOrder>): Promise<WorkOrder> {
  return apiRequest<WorkOrder>('/maintenance/work-orders', {
    method: 'POST',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function updateWorkOrder(id: string | number, data: Partial<WorkOrder>): Promise<WorkOrder> {
  return apiRequest<WorkOrder>(`/maintenance/work-orders/${id}`, {
    method: 'PUT',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function deleteWorkOrder(id: string | number): Promise<{ message: string }> {
  return apiRequest<{ message: string }>(`/maintenance/work-orders/${id}`, {
    method: 'DELETE',
  });
}

export function getContracts(): Promise<MaintenanceContract[]> {
  return apiRequest<MaintenanceContract[]>('/maintenance/contracts');
}

export function getContractById(id: string | number): Promise<MaintenanceContract> {
  return apiRequest<MaintenanceContract>(`/maintenance/contracts/${id}`);
}

export function createContract(data: Partial<MaintenanceContract>): Promise<MaintenanceContract> {
  return apiRequest<MaintenanceContract>('/maintenance/contracts', {
    method: 'POST',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function updateContract(
  id: string | number,
  data: Partial<MaintenanceContract>,
): Promise<MaintenanceContract> {
  return apiRequest<MaintenanceContract>(`/maintenance/contracts/${id}`, {
    method: 'PUT',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function deleteContract(id: string | number): Promise<{ message: string }> {
  return apiRequest<{ message: string }>(`/maintenance/contracts/${id}`, {
    method: 'DELETE',
  });
}
