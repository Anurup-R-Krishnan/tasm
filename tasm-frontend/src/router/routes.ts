import type { RouteRecordRaw } from 'vue-router';

interface AppRouteDefinition {
  path: string;
  name: string;
  title: string;
  icon?: string;
  showInSidebar?: boolean;
  component: NonNullable<RouteRecordRaw['component']>;
}

export const appRoutes: AppRouteDefinition[] = [
  {
    path: '',
    name: 'Dashboard',
    title: 'Dashboard',
    icon: 'dashboard',
    showInSidebar: true,
    component: () => import('../views/Dashboard.vue'),
  },
  {
    path: 'asset-registry',
    name: 'AssetRegistry',
    title: 'Asset Registry',
    icon: 'inventory_2',
    showInSidebar: true,
    component: () => import('../views/AssetRegistry.vue'),
  },
  {
    path: 'consumables-supplies-tracker',
    name: 'ConsumablesSuppliesTracker',
    title: 'Consumables',
    icon: 'inventory',
    showInSidebar: true,
    component: () => import('../views/ConsumablesSuppliesTracker.vue'),
  },
  {
    path: 'work-order-detail',
    name: 'WorkOrderDetail',
    title: 'Work Orders',
    icon: 'assignment',
    showInSidebar: true,
    component: () => import('../views/WorkOrderDetail.vue'),
  },
  {
    path: 'audit-sessions',
    name: 'AuditSessions',
    title: 'Audit Sessions',
    icon: 'fact_check',
    showInSidebar: true,
    component: () => import('../views/AuditSessions.vue'),
  },
  {
    path: 'user-management-settings',
    name: 'UserManagementSettings',
    title: 'Settings',
    icon: 'settings',
    showInSidebar: true,
    component: () => import('../views/UserManagementSettings.vue'),
  },
  {
    path: 'financial-analytics-dashboard',
    name: 'FinancialAnalyticsDashboard',
    title: 'Financials',
    icon: 'payments',
    showInSidebar: true,
    component: () => import('../views/FinancialAnalyticsDashboard.vue'),
  },
  {
    path: 'asset-depreciation-ledger',
    name: 'AssetDepreciationLedger',
    title: 'Depreciation',
    icon: 'trending_down',
    showInSidebar: true,
    component: () => import('../views/AssetDepreciationLedger.vue'),
  },
  {
    path: 'software-license-registry',
    name: 'SoftwareLicenseRegistry',
    title: 'Licenses',
    icon: 'verified_user',
    showInSidebar: true,
    component: () => import('../views/SoftwareLicenseRegistry.vue'),
  },
  {
    path: 'reservations-bookings',
    name: 'ReservationsBookings',
    title: 'Reservations',
    icon: 'event_available',
    showInSidebar: true,
    component: () => import('../views/ReservationsBookings.vue'),
  },
  {
    path: 'asset-detail/:id',
    name: 'AssetDetail',
    title: 'Asset Detail',
    component: () => import('../views/AssetDetail.vue'),
  },
  {
    path: 'maintenance-contracts-overview',
    name: 'MaintenanceContractsOverview',
    title: 'Maintenance Contracts Overview',
    component: () => import('../views/MaintenanceContractsOverview.vue'),
  },
  {
    path: 'user-access-role-control',
    name: 'UserAccessRoleControl',
    title: 'User Roles',
    icon: 'admin_panel_settings',
    showInSidebar: true,
    component: () => import('../views/UserAccessRoleControl.vue'),
  },
  {
    path: 'maintenance-schedule-contracts',
    name: 'MaintenanceScheduleContracts',
    title: 'Maintenance',
    icon: 'build_circle',
    showInSidebar: true,
    component: () => import('../views/MaintenanceScheduleContracts.vue'),
  },
  {
    path: 'maintenance-service-history-log',
    name: 'MaintenanceServiceHistoryLog',
    title: 'Maintenance Service History Log',
    component: () => import('../views/MaintenanceServiceHistoryLog.vue'),
  },
  {
    path: 'audit-discrepancy-resolution',
    name: 'AuditDiscrepancyResolution',
    title: 'Audit Cleanup',
    icon: 'rule_folder',
    showInSidebar: true,
    component: () => import('../views/AuditDiscrepancyResolution.vue'),
  },
  {
    path: 'custom-report-builder',
    name: 'CustomReportBuilder',
    title: 'Report Builder',
    icon: 'edit_note',
    showInSidebar: true,
    component: () => import('../views/CustomReportBuilder.vue'),
  },
  {
    path: 'financial-ledger-detail',
    name: 'FinancialLedgerDetail',
    title: 'Financial Ledger Detail',
    component: () => import('../views/FinancialLedgerDetail.vue'),
  },
  {
    path: 'asset-check-out-flow',
    name: 'AssetCheckOutFlow',
    title: 'Asset Check Out Flow',
    component: () => import('../views/AssetCheckOutFlow.vue'),
  },
  {
    path: 'lease-agreement-detail',
    name: 'LeaseAgreementDetail',
    title: 'Lease Agreement Detail',
    component: () => import('../views/LeaseAgreementDetail.vue'),
  },
  {
    path: 'procurement-request-detail',
    name: 'ProcurementRequestDetail',
    title: 'Procurement Request Detail',
    component: () => import('../views/ProcurementRequestDetail.vue'),
  },
  {
    path: 'procurement-pipeline',
    name: 'ProcurementPipeline',
    title: 'Procurement',
    icon: 'shopping_cart',
    showInSidebar: true,
    component: () => import('../views/ProcurementPipeline.vue'),
  },
  {
    path: 'audit-scan-mode-mobile',
    name: 'AuditScanModeMobile',
    title: 'Audit Scan Mode Mobile',
    component: () => import('../views/AuditScanModeMobile.vue'),
  },
  {
    path: 'add-new-asset-form',
    name: 'AddNewAssetForm',
    title: 'Add New Asset',
    component: () => import('../views/AddNewAssetForm.vue'),
  },
  {
    path: 'employee-self-service-portal',
    name: 'EmployeeSelfServicePortal',
    title: 'Self Service',
    icon: 'person_pin',
    showInSidebar: true,
    component: () => import('../views/EmployeeSelfServicePortal.vue'),
  },
  {
    path: 'reports-analytics-dashboard',
    name: 'ReportsAnalyticsDashboard',
    title: 'Reports',
    icon: 'analytics',
    showInSidebar: true,
    component: () => import('../views/ReportsAnalyticsDashboard.vue'),
  },
  {
    path: 'stockroom-inventory',
    name: 'StockroomInventory',
    title: 'Stockrooms',
    icon: 'warehouse',
    showInSidebar: true,
    component: () => import('../views/StockroomInventory.vue'),
  },
  {
    path: 'inventory-card-view',
    name: 'InventoryCardView',
    title: 'Asset Cards',
    icon: 'grid_view',
    showInSidebar: true,
    component: () => import('../views/InventoryCardView.vue'),
  },
  {
    path: 'alerts-notifications-center',
    name: 'AlertsNotificationsCenter',
    title: 'Alerts',
    icon: 'notifications',
    showInSidebar: true,
    component: () => import('../views/AlertsNotificationsCenter.vue'),
  },
];

export const sidebarGroups = [
  {
    name: 'Inventory',
    items: [
      'AssetRegistry',
      'StockroomInventory',
      'ConsumablesSuppliesTracker',
      'SoftwareLicenseRegistry',
      'InventoryCardView',
    ],
  },
  {
    name: 'Operations',
    items: [
      'WorkOrderDetail',
      'MaintenanceScheduleContracts',
      'AuditSessions',
      'ReservationsBookings',
    ],
  },
  {
    name: 'Financials',
    items: ['ProcurementPipeline', 'FinancialAnalyticsDashboard', 'AssetDepreciationLedger'],
  },
  {
    name: 'Compliance',
    items: ['AuditDiscrepancyResolution', 'CustomReportBuilder'],
  },
  {
    name: 'System',
    items: ['AlertsNotificationsCenter', 'EmployeeSelfServicePortal', 'UserAccessRoleControl'],
  },
];

export const sidebarRoutes = appRoutes.filter((route) => route.showInSidebar);
