import { createRouter, createWebHistory } from 'vue-router'
import AppLayout from '../components/AppLayout.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      component: AppLayout,
      children: [
        {
          path: '',
          name: 'Dashboard',
          component: () => import('../views/Dashboard.vue')
        },
        { path: 'financial-analytics-dashboard', name: 'FinancialAnalyticsDashboard', component: () => import('../views/FinancialAnalyticsDashboard.vue') },
        { path: 'asset-depreciation-ledger', name: 'AssetDepreciationLedger', component: () => import('../views/AssetDepreciationLedger.vue') },
        { path: 'stockrooms-consumables-2', name: 'StockroomsConsumables2', component: () => import('../views/StockroomsConsumables2.vue') },
        { path: 'asset-registry-redesigned', name: 'AssetRegistryRedesigned', component: () => import('../views/AssetRegistryRedesigned.vue') },
        { path: 'inventory-management-card-view', name: 'InventoryManagementCardView', component: () => import('../views/InventoryManagementCardView.vue') },
        { path: 'software-license-registry', name: 'SoftwareLicenseRegistry', component: () => import('../views/SoftwareLicenseRegistry.vue') },
        { path: 'financial-reports-analytics', name: 'FinancialReportsAnalytics', component: () => import('../views/FinancialReportsAnalytics.vue') },
        { path: 'reservations-bookings', name: 'ReservationsBookings', component: () => import('../views/ReservationsBookings.vue') },
        { path: 'audit-session-management-1', name: 'AuditSessionManagement1', component: () => import('../views/AuditSessionManagement1.vue') },
        { path: 'maintenance-contracts-management', name: 'MaintenanceContractsManagement', component: () => import('../views/MaintenanceContractsManagement.vue') },
        { path: 'procurement-pipeline-2', name: 'ProcurementPipeline2', component: () => import('../views/ProcurementPipeline2.vue') },
        { path: 'employee-self-service-portal-1', name: 'EmployeeSelfServicePortal1', component: () => import('../views/EmployeeSelfServicePortal1.vue') },
        { path: 'consumables-supplies-tracker', name: 'ConsumablesSuppliesTracker', component: () => import('../views/ConsumablesSuppliesTracker.vue') },
        { path: 'asset-detail-macbook-pro-16-ast-2023-0892', name: 'AssetDetailMacbookPro16Ast20230892', component: () => import('../views/AssetDetailMacbookPro16Ast20230892.vue') },
        { path: 'employee-self-service-portal-2', name: 'EmployeeSelfServicePortal2', component: () => import('../views/EmployeeSelfServicePortal2.vue') },
        { path: 'maintenance-contracts-overview', name: 'MaintenanceContractsOverview', component: () => import('../views/MaintenanceContractsOverview.vue') },
        { path: 'procurement-pipeline-tracker', name: 'ProcurementPipelineTracker', component: () => import('../views/ProcurementPipelineTracker.vue') },
        { path: 'user-access-role-control', name: 'UserAccessRoleControl', component: () => import('../views/UserAccessRoleControl.vue') },
        { path: 'stockrooms-consumables-1', name: 'StockroomsConsumables1', component: () => import('../views/StockroomsConsumables1.vue') },
        { path: 'reports-financials-dashboard', name: 'ReportsFinancialsDashboard', component: () => import('../views/ReportsFinancialsDashboard.vue') },
        { path: 'asset-check-out-step-2-select-recipient', name: 'AssetCheckOutStep2SelectRecipient', component: () => import('../views/AssetCheckOutStep2SelectRecipient.vue') },
        { path: 'scheduled-maintenance-tracker', name: 'ScheduledMaintenanceTracker', component: () => import('../views/ScheduledMaintenanceTracker.vue') },
        { path: 'maintenance-service-history-log', name: 'MaintenanceServiceHistoryLog', component: () => import('../views/MaintenanceServiceHistoryLog.vue') },
        { path: 'audit-session-management-2', name: 'AuditSessionManagement2', component: () => import('../views/AuditSessionManagement2.vue') },
        { path: 'work-order-detail', name: 'WorkOrderDetail', component: () => import('../views/WorkOrderDetail.vue') },
        { path: 'maintenance-schedule-contracts', name: 'MaintenanceScheduleContracts', component: () => import('../views/MaintenanceScheduleContracts.vue') },
        { path: 'asset-registry-with-detail-drawer', name: 'AssetRegistryWithDetailDrawer', component: () => import('../views/AssetRegistryWithDetailDrawer.vue') },
        { path: 'audit-discrepancy-resolution', name: 'AuditDiscrepancyResolution', component: () => import('../views/AuditDiscrepancyResolution.vue') },
        { path: 'custom-report-builder', name: 'CustomReportBuilder', component: () => import('../views/CustomReportBuilder.vue') },
        { path: 'financial-ledger-detail', name: 'FinancialLedgerDetail', component: () => import('../views/FinancialLedgerDetail.vue') },
        { path: 'asset-check-out-flow', name: 'AssetCheckOutFlow', component: () => import('../views/AssetCheckOutFlow.vue') },
        { path: 'lease-agreement-detail', name: 'LeaseAgreementDetail', component: () => import('../views/LeaseAgreementDetail.vue') },
        { path: 'reports-analytics-dashboard', name: 'ReportsAnalyticsDashboard', component: () => import('../views/ReportsAnalyticsDashboard.vue') },
        { path: 'employee-self-service-portal-3', name: 'EmployeeSelfServicePortal3', component: () => import('../views/EmployeeSelfServicePortal3.vue') },
        { path: 'procurement-request-detail', name: 'ProcurementRequestDetail', component: () => import('../views/ProcurementRequestDetail.vue') },
        { path: 'stockroom-inventory-management', name: 'StockroomInventoryManagement', component: () => import('../views/StockroomInventoryManagement.vue') },
        { path: 'work-order-detail-view', name: 'WorkOrderDetailView', component: () => import('../views/WorkOrderDetailView.vue') },
        { path: 'asset-registry', name: 'AssetRegistry', component: () => import('../views/AssetRegistry.vue') },
        { path: 'procurement-pipeline-1', name: 'ProcurementPipeline1', component: () => import('../views/ProcurementPipeline1.vue') },
        { path: 'audit-scan-mode-mobile', name: 'AuditScanModeMobile', component: () => import('../views/AuditScanModeMobile.vue') },
        { path: 'user-management-settings', name: 'UserManagementSettings', component: () => import('../views/UserManagementSettings.vue') },
        { path: 'add-new-asset-form', name: 'AddNewAssetForm', component: () => import('../views/AddNewAssetForm.vue') },
        { path: 'employee-self-service-portal', name: 'EmployeeSelfServicePortal', component: () => import('../views/EmployeeSelfServicePortal.vue') },
        { path: 'alerts-notifications-center', name: 'AlertsNotificationsCenter', component: () => import('../views/AlertsNotificationsCenter.vue') },
        { path: 'setup-wizard-locations-configuration', name: 'SetupWizardLocationsConfiguration', component: () => import('../views/SetupWizardLocationsConfiguration.vue') },
        { path: 'procurement-pipeline-tracking', name: 'ProcurementPipelineTracking', component: () => import('../views/ProcurementPipelineTracking.vue') },
        { path: 'depreciation-ledger-view', name: 'DepreciationLedgerView', component: () => import('../views/DepreciationLedgerView.vue') }
      ]
    }
  ]
})

export default router
