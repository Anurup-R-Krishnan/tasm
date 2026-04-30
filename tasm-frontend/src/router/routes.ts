import type { RouteRecordRaw } from 'vue-router'

export interface AppRouteDefinition {
  path: string
  name: string
  title: string
  icon?: string
  showInSidebar?: boolean
  component: NonNullable<RouteRecordRaw['component']>
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
    path: 'audit-session-management-1',
    name: 'AuditSessionManagement1',
    title: 'Audit Sessions',
    icon: 'fact_check',
    showInSidebar: true,
    component: () => import('../views/AuditSessionManagement1.vue'),
  },
  {
    path: 'user-management-settings',
    name: 'UserManagementSettings',
    title: 'Settings',
    icon: 'settings',
    showInSidebar: true,
    component: () => import('../views/UserManagementSettings.vue'),
  },
  { path: 'financial-analytics-dashboard', name: 'FinancialAnalyticsDashboard', title: 'Financial Analytics Dashboard', component: () => import('../views/FinancialAnalyticsDashboard.vue') },
  { path: 'asset-depreciation-ledger', name: 'AssetDepreciationLedger', title: 'Asset Depreciation Ledger', component: () => import('../views/AssetDepreciationLedger.vue') },
  { path: 'stockrooms-consumables-2', name: 'StockroomsConsumables2', title: 'Stockrooms Consumables 2', component: () => import('../views/StockroomsConsumables2.vue') },
  { path: 'asset-registry-redesigned', name: 'AssetRegistryRedesigned', title: 'Asset Registry Redesigned', component: () => import('../views/AssetRegistryRedesigned.vue') },
  { path: 'inventory-management-card-view', name: 'InventoryManagementCardView', title: 'Inventory Management Card View', component: () => import('../views/InventoryManagementCardView.vue') },
  { path: 'software-license-registry', name: 'SoftwareLicenseRegistry', title: 'Software License Registry', component: () => import('../views/SoftwareLicenseRegistry.vue') },
  { path: 'financial-reports-analytics', name: 'FinancialReportsAnalytics', title: 'Financial Reports Analytics', component: () => import('../views/FinancialReportsAnalytics.vue') },
  { path: 'reservations-bookings', name: 'ReservationsBookings', title: 'Reservations Bookings', component: () => import('../views/ReservationsBookings.vue') },
  { path: 'maintenance-contracts-management', name: 'MaintenanceContractsManagement', title: 'Maintenance Contracts Management', component: () => import('../views/MaintenanceContractsManagement.vue') },
  { path: 'procurement-pipeline-2', name: 'ProcurementPipeline2', title: 'Procurement Pipeline 2', component: () => import('../views/ProcurementPipeline2.vue') },
  { path: 'employee-self-service-portal-1', name: 'EmployeeSelfServicePortal1', title: 'Employee Self Service Portal 1', component: () => import('../views/EmployeeSelfServicePortal1.vue') },
  { path: 'asset-detail-macbook-pro-16-ast-2023-0892', name: 'AssetDetailMacbookPro16Ast20230892', title: 'Asset Detail', component: () => import('../views/AssetDetailMacbookPro16Ast20230892.vue') },
  { path: 'employee-self-service-portal-2', name: 'EmployeeSelfServicePortal2', title: 'Employee Self Service Portal 2', component: () => import('../views/EmployeeSelfServicePortal2.vue') },
  { path: 'maintenance-contracts-overview', name: 'MaintenanceContractsOverview', title: 'Maintenance Contracts Overview', component: () => import('../views/MaintenanceContractsOverview.vue') },
  { path: 'procurement-pipeline-tracker', name: 'ProcurementPipelineTracker', title: 'Procurement Pipeline Tracker', component: () => import('../views/ProcurementPipelineTracker.vue') },
  { path: 'user-access-role-control', name: 'UserAccessRoleControl', title: 'User Access Role Control', component: () => import('../views/UserAccessRoleControl.vue') },
  { path: 'stockrooms-consumables-1', name: 'StockroomsConsumables1', title: 'Stockrooms Consumables 1', component: () => import('../views/StockroomsConsumables1.vue') },
  { path: 'reports-financials-dashboard', name: 'ReportsFinancialsDashboard', title: 'Reports Financials Dashboard', component: () => import('../views/ReportsFinancialsDashboard.vue') },
  { path: 'asset-check-out-step-2-select-recipient', name: 'AssetCheckOutStep2SelectRecipient', title: 'Asset Check Out Recipient', component: () => import('../views/AssetCheckOutStep2SelectRecipient.vue') },
  { path: 'scheduled-maintenance-tracker', name: 'ScheduledMaintenanceTracker', title: 'Scheduled Maintenance Tracker', component: () => import('../views/ScheduledMaintenanceTracker.vue') },
  { path: 'maintenance-service-history-log', name: 'MaintenanceServiceHistoryLog', title: 'Maintenance Service History Log', component: () => import('../views/MaintenanceServiceHistoryLog.vue') },
  { path: 'audit-session-management-2', name: 'AuditSessionManagement2', title: 'Audit Session Management 2', component: () => import('../views/AuditSessionManagement2.vue') },
  { path: 'maintenance-schedule-contracts', name: 'MaintenanceScheduleContracts', title: 'Maintenance Schedule Contracts', component: () => import('../views/MaintenanceScheduleContracts.vue') },
  { path: 'asset-registry-with-detail-drawer', name: 'AssetRegistryWithDetailDrawer', title: 'Asset Registry With Detail Drawer', component: () => import('../views/AssetRegistryWithDetailDrawer.vue') },
  { path: 'audit-discrepancy-resolution', name: 'AuditDiscrepancyResolution', title: 'Audit Discrepancy Resolution', component: () => import('../views/AuditDiscrepancyResolution.vue') },
  { path: 'custom-report-builder', name: 'CustomReportBuilder', title: 'Custom Report Builder', component: () => import('../views/CustomReportBuilder.vue') },
  { path: 'financial-ledger-detail', name: 'FinancialLedgerDetail', title: 'Financial Ledger Detail', component: () => import('../views/FinancialLedgerDetail.vue') },
  { path: 'asset-check-out-flow', name: 'AssetCheckOutFlow', title: 'Asset Check Out Flow', component: () => import('../views/AssetCheckOutFlow.vue') },
  { path: 'lease-agreement-detail', name: 'LeaseAgreementDetail', title: 'Lease Agreement Detail', component: () => import('../views/LeaseAgreementDetail.vue') },
  { path: 'reports-analytics-dashboard', name: 'ReportsAnalyticsDashboard', title: 'Reports Analytics Dashboard', component: () => import('../views/ReportsAnalyticsDashboard.vue') },
  { path: 'employee-self-service-portal-3', name: 'EmployeeSelfServicePortal3', title: 'Employee Self Service Portal 3', component: () => import('../views/EmployeeSelfServicePortal3.vue') },
  { path: 'procurement-request-detail', name: 'ProcurementRequestDetail', title: 'Procurement Request Detail', component: () => import('../views/ProcurementRequestDetail.vue') },
  { path: 'stockroom-inventory-management', name: 'StockroomInventoryManagement', title: 'Stockroom Inventory Management', component: () => import('../views/StockroomInventoryManagement.vue') },
  { path: 'work-order-detail-view', name: 'WorkOrderDetailView', title: 'Work Order Detail View', component: () => import('../views/WorkOrderDetailView.vue') },
  { path: 'procurement-pipeline-1', name: 'ProcurementPipeline1', title: 'Procurement Pipeline 1', component: () => import('../views/ProcurementPipeline1.vue') },
  { path: 'audit-scan-mode-mobile', name: 'AuditScanModeMobile', title: 'Audit Scan Mode Mobile', component: () => import('../views/AuditScanModeMobile.vue') },
  { path: 'add-new-asset-form', name: 'AddNewAssetForm', title: 'Add New Asset', component: () => import('../views/AddNewAssetForm.vue') },
  { path: 'employee-self-service-portal', name: 'EmployeeSelfServicePortal', title: 'Employee Self Service Portal', component: () => import('../views/EmployeeSelfServicePortal.vue') },
  { path: 'alerts-notifications-center', name: 'AlertsNotificationsCenter', title: 'Alerts Notifications Center', component: () => import('../views/AlertsNotificationsCenter.vue') },
  { path: 'setup-wizard-locations-configuration', name: 'SetupWizardLocationsConfiguration', title: 'Setup Wizard Locations Configuration', component: () => import('../views/SetupWizardLocationsConfiguration.vue') },
  { path: 'procurement-pipeline-tracking', name: 'ProcurementPipelineTracking', title: 'Procurement Pipeline Tracking', component: () => import('../views/ProcurementPipelineTracking.vue') },
  { path: 'depreciation-ledger-view', name: 'DepreciationLedgerView', title: 'Depreciation Ledger View', component: () => import('../views/DepreciationLedgerView.vue') },
]

export const sidebarRoutes = appRoutes.filter((route) => route.showInSidebar)
