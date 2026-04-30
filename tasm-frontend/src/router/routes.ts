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
  { path: 'financial-analytics-dashboard', name: 'FinancialAnalyticsDashboard', title: 'Financial Analytics Dashboard', component: () => import('../views/FinancialAnalyticsDashboard.vue') },
  { path: 'asset-depreciation-ledger', name: 'AssetDepreciationLedger', title: 'Asset Depreciation Ledger', component: () => import('../views/AssetDepreciationLedger.vue') },
  { path: 'software-license-registry', name: 'SoftwareLicenseRegistry', title: 'Software License Registry', component: () => import('../views/SoftwareLicenseRegistry.vue') },
  { path: 'reservations-bookings', name: 'ReservationsBookings', title: 'Reservations Bookings', component: () => import('../views/ReservationsBookings.vue') },
  { path: 'asset-detail/:id', name: 'AssetDetail', title: 'Asset Detail', component: () => import('../views/AssetDetail.vue') },
  { path: 'maintenance-contracts-overview', name: 'MaintenanceContractsOverview', title: 'Maintenance Contracts Overview', component: () => import('../views/MaintenanceContractsOverview.vue') },
  { path: 'user-access-role-control', name: 'UserAccessRoleControl', title: 'User Access Role Control', component: () => import('../views/UserAccessRoleControl.vue') },
  { path: 'asset-check-out-step-2-select-recipient', name: 'AssetCheckOutStep2SelectRecipient', title: 'Asset Check Out Recipient', component: () => import('../views/AssetCheckOutStep2SelectRecipient.vue') },
  { path: 'scheduled-maintenance-tracker', name: 'ScheduledMaintenanceTracker', title: 'Scheduled Maintenance Tracker', component: () => import('../views/ScheduledMaintenanceTracker.vue') },
  { path: 'maintenance-service-history-log', name: 'MaintenanceServiceHistoryLog', title: 'Maintenance Service History Log', component: () => import('../views/MaintenanceServiceHistoryLog.vue') },
  { path: 'audit-discrepancy-resolution', name: 'AuditDiscrepancyResolution', title: 'Audit Discrepancy Resolution', component: () => import('../views/AuditDiscrepancyResolution.vue') },
  { path: 'custom-report-builder', name: 'CustomReportBuilder', title: 'Custom Report Builder', component: () => import('../views/CustomReportBuilder.vue') },
  { path: 'financial-ledger-detail', name: 'FinancialLedgerDetail', title: 'Financial Ledger Detail', component: () => import('../views/FinancialLedgerDetail.vue') },
  { path: 'asset-check-out-flow', name: 'AssetCheckOutFlow', title: 'Asset Check Out Flow', component: () => import('../views/AssetCheckOutFlow.vue') },
  { path: 'lease-agreement-detail', name: 'LeaseAgreementDetail', title: 'Lease Agreement Detail', component: () => import('../views/LeaseAgreementDetail.vue') },
  { path: 'procurement-request-detail', name: 'ProcurementRequestDetail', title: 'Procurement Request Detail', component: () => import('../views/ProcurementRequestDetail.vue') },
  { path: 'procurement-pipeline', name: 'ProcurementPipeline', title: 'Procurement Pipeline', component: () => import('../views/ProcurementPipeline.vue') },
  { path: 'audit-scan-mode-mobile', name: 'AuditScanModeMobile', title: 'Audit Scan Mode Mobile', component: () => import('../views/AuditScanModeMobile.vue') },
  { path: 'add-new-asset-form', name: 'AddNewAssetForm', title: 'Add New Asset', component: () => import('../views/AddNewAssetForm.vue') },
  { path: 'employee-self-service-portal', name: 'EmployeeSelfServicePortal', title: 'Employee Self Service Portal', component: () => import('../views/EmployeeSelfServicePortal.vue') },
  { path: 'alerts-notifications-center', name: 'AlertsNotificationsCenter', title: 'Alerts Notifications Center', component: () => import('../views/AlertsNotificationsCenter.vue') },
  { path: 'setup-wizard-locations-configuration', name: 'SetupWizardLocationsConfiguration', title: 'Setup Wizard Locations Configuration', component: () => import('../views/SetupWizardLocationsConfiguration.vue') },
]

export const sidebarRoutes = appRoutes.filter((route) => route.showInSidebar)
