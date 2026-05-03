import { apiRequest } from './client';

export interface WorkOrder {
  id: number;
  title: string;
  severity: string;
  assetLocation: string;
  technician: string;
  status: string;
  targetDate: string;
  [key: string]: any;
}

export function getWorkOrders(): Promise<WorkOrder[]> {
  return apiRequest<WorkOrder[]>('/work-orders');
}

export function getWorkOrderById(id: string | number): Promise<WorkOrder> {
  return apiRequest<WorkOrder>(`/work-orders/${id}`);
}

export function createWorkOrder(data: Partial<WorkOrder>): Promise<WorkOrder> {
  return apiRequest<WorkOrder>('/work-orders', {
    method: 'POST',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function updateWorkOrder(id: string | number, data: Partial<WorkOrder>): Promise<WorkOrder> {
  return apiRequest<WorkOrder>(`/work-orders/${id}`, {
    method: 'PUT',
    body: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' },
  });
}

export function deleteWorkOrder(id: string | number): Promise<{ message: string }> {
  return apiRequest<{ message: string }>(`/work-orders/${id}`, {
    method: 'DELETE',
  });
}
