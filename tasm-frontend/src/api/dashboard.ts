import { apiRequest } from './client';

export interface DashboardStats {
  totalAssets: number;
  checkedOutAssets: number;
  inStockAssets: number;
  retiredAssets: number;
  totalValue: number;
  activeAudits: number;
  openDiscrepancies: number;
  openWorkOrders: number;
  pendingProcurements: number;
  expiringWarranties: number;
  unreadAlerts: number;
  lowStockConsumables: number;
  activeReservations: number;
  totalUsers: number;
}

export function getDashboardStats(): Promise<DashboardStats> {
  return apiRequest<DashboardStats>('/dashboard/stats');
}
