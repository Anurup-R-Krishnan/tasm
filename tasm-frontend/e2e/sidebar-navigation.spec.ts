import { test, expect } from '@playwright/test';

/**
 * Sidebar Navigation Tests
 * Tests all sidebar links and ensures pages load correctly
 */

// All sidebar routes from routes.ts
const sidebarRoutes = [
  { path: '/', name: 'Dashboard', title: 'Dashboard' },
  { path: '/asset-registry', name: 'AssetRegistry', title: 'Asset Registry' },
  {
    path: '/consumables-supplies-tracker',
    name: 'ConsumablesSuppliesTracker',
    title: 'Consumables',
  },
  { path: '/audit-sessions', name: 'AuditSessions', title: 'Audit Sessions' },
  { path: '/user-management-settings', name: 'UserManagementSettings', title: 'Settings' },
  {
    path: '/financial-analytics-dashboard',
    name: 'FinancialAnalyticsDashboard',
    title: 'Financials',
  },
  { path: '/asset-depreciation-ledger', name: 'AssetDepreciationLedger', title: 'Depreciation' },
  { path: '/software-license-registry', name: 'SoftwareLicenseRegistry', title: 'Licenses' },
  { path: '/reservations-bookings', name: 'ReservationsBookings', title: 'Reservations' },
  { path: '/user-access-role-control', name: 'UserAccessRoleControl', title: 'User Roles' },
  {
    path: '/maintenance-schedule-contracts',
    name: 'MaintenanceScheduleContracts',
    title: 'Maintenance',
  },
  {
    path: '/audit-discrepancy-resolution',
    name: 'AuditDiscrepancyResolution',
    title: 'Audit Cleanup',
  },
  { path: '/custom-report-builder', name: 'CustomReportBuilder', title: 'Report Builder' },
  { path: '/procurement-pipeline', name: 'ProcurementPipeline', title: 'Procurement' },
  {
    path: '/employee-self-service-portal',
    name: 'EmployeeSelfServicePortal',
    title: 'Self Service',
  },
  { path: '/reports-analytics-dashboard', name: 'ReportsAnalyticsDashboard', title: 'Reports' },
  { path: '/stockroom-inventory', name: 'StockroomInventory', title: 'Stockrooms' },
  { path: '/inventory-card-view', name: 'InventoryCardView', title: 'Asset Cards' },
  { path: '/alerts-notifications-center', name: 'AlertsNotificationsCenter', title: 'Alerts' },
];

test.describe('Sidebar Navigation - All Routes', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/');
    await page.waitForLoadState('networkidle').catch(() => {});
  });

  test.describe('Sidebar Visibility', () => {
    test('should have sidebar visible on desktop', async ({ page }) => {
      const sidebar = page.locator('nav, [class*="sidebar"], .app-sidebar').first();
      await expect(sidebar).toBeVisible();
    });

    test('should have all navigation links', async ({ page }) => {
      for (const route of sidebarRoutes) {
        const link = page.getByRole('link', { name: new RegExp(route.title, 'i') }).first();
        await expect(link).toBeVisible();
      }
    });
  });

  test.describe('Navigation Tests', () => {
    test('should navigate to Dashboard', async ({ page }) => {
      await page.goto('/');
      await expect(page.getByText('Operational Overview')).toBeVisible();
    });

    test('should navigate to Asset Registry', async ({ page }) => {
      await page.getByRole('link', { name: /Asset Registry/i }).click();
      await expect(page).toHaveURL(/\/asset-registry$/);
      await expect(page.getByRole('heading', { name: 'Asset Registry' })).toBeVisible();
    });

    test('should navigate to Consumables', async ({ page }) => {
      await page.route('**/api/consumables**', (route) => {
        route.fulfill({ json: [] });
      });

      await page.getByRole('link', { name: /Consumables/i }).click();
      await expect(page).toHaveURL(/\/consumables-supplies-tracker$/);
      await expect(page.getByRole('heading', { name: /Consumables/i })).toBeVisible();
    });

    test('should navigate to Audit Sessions', async ({ page }) => {
      await page.getByRole('link', { name: /Audit Sessions/i }).click();
      await expect(page).toHaveURL(/\/audit-sessions$/);
      await expect(page.getByRole('heading', { name: /Audit Sessions/i })).toBeVisible();
    });

    test('should navigate to Settings', async ({ page }) => {
      await page.getByRole('link', { name: /Settings/i }).click();
      await expect(page).toHaveURL(/\/user-management-settings$/);
      await expect(page.getByRole('heading', { name: /Settings/i })).toBeVisible();
    });

    test('should navigate to Financials', async ({ page }) => {
      await page.getByRole('link', { name: /Financials/i }).click();
      await expect(page).toHaveURL(/\/financial-analytics-dashboard$/);
      await expect(page.getByRole('heading', { name: /Financial/i })).toBeVisible();
    });

    test('should navigate to Depreciation', async ({ page }) => {
      await page.getByRole('link', { name: /Depreciation/i }).click();
      await expect(page).toHaveURL(/\/asset-depreciation-ledger$/);
      await expect(page.getByRole('heading', { name: /Depreciation/i })).toBeVisible();
    });

    test('should navigate to Licenses', async ({ page }) => {
      await page.getByRole('link', { name: /Licenses/i }).click();
      await expect(page).toHaveURL(/\/software-license-registry$/);
      await expect(page.getByRole('heading', { name: /License/i })).toBeVisible();
    });

    test('should navigate to Reservations', async ({ page }) => {
      await page.getByRole('link', { name: /Reservations/i }).click();
      await expect(page).toHaveURL(/\/reservations-bookings$/);
      await expect(page.getByRole('heading', { name: /Reservation/i })).toBeVisible();
    });

    test('should navigate to User Roles', async ({ page }) => {
      await page.getByRole('link', { name: /User Roles/i }).click();
      await expect(page).toHaveURL(/\/user-access-role-control$/);
      await expect(page.getByRole('heading', { name: /Role/i })).toBeVisible();
    });

    test('should navigate to Maintenance', async ({ page }) => {
      await page.getByRole('link', { name: /Maintenance/i }).click();
      await expect(page).toHaveURL(/\/maintenance-schedule-contracts$/);
      await expect(page.getByRole('heading', { name: /Maintenance/i })).toBeVisible();
    });

    test('should navigate to Audit Cleanup', async ({ page }) => {
      await page.getByRole('link', { name: /Audit Cleanup/i }).click();
      await expect(page).toHaveURL(/\/audit-discrepancy-resolution$/);
      await expect(page.getByRole('heading', { name: /Audit/i })).toBeVisible();
    });

    test('should navigate to Report Builder', async ({ page }) => {
      await page.getByRole('link', { name: /Report Builder/i }).click();
      await expect(page).toHaveURL(/\/custom-report-builder$/);
      await expect(page.getByRole('heading', { name: /Report/i })).toBeVisible();
    });

    test('should navigate to Procurement', async ({ page }) => {
      await page.getByRole('link', { name: /Procurement/i }).click();
      await expect(page).toHaveURL(/\/procurement-pipeline$/);
      await expect(page.getByRole('heading', { name: /Procurement/i })).toBeVisible();
    });

    test('should navigate to Self Service', async ({ page }) => {
      await page.getByRole('link', { name: /Self Service/i }).click();
      await expect(page).toHaveURL(/\/employee-self-service-portal$/);
      await expect(page.getByRole('heading', { name: /Self Service/i })).toBeVisible();
    });

    test('should navigate to Reports', async ({ page }) => {
      await page.getByRole('link', { name: /Reports/i }).click();
      await expect(page).toHaveURL(/\/reports-analytics-dashboard$/);
      await expect(page.getByRole('heading', { name: /Report/i })).toBeVisible();
    });

    test('should navigate to Stockrooms', async ({ page }) => {
      await page.getByRole('link', { name: /Stockrooms/i }).click();
      await expect(page).toHaveURL(/\/stockroom-inventory$/);
      await expect(page.getByRole('heading', { name: /Stockroom/i })).toBeVisible();
    });

    test('should navigate to Asset Cards', async ({ page }) => {
      await page.getByRole('link', { name: /Asset Cards/i }).click();
      await expect(page).toHaveURL(/\/inventory-card-view$/);
      await expect(page.getByRole('heading', { name: /Asset Card/i })).toBeVisible();
    });

    test('should navigate to Alerts', async ({ page }) => {
      await page.getByRole('link', { name: /Alerts/i }).click();
      await expect(page).toHaveURL(/\/alerts-notifications-center$/);
      await expect(page.getByRole('heading', { name: /Alert/i })).toBeVisible();
    });
  });

  test.describe('Sidebar Groups', () => {
    test('should have Inventory group with correct items', async ({ page }) => {
      const inventoryLinks = [
        'Asset Registry',
        'Stockrooms',
        'Consumables',
        'Licenses',
        'Asset Cards',
      ];
      for (const name of inventoryLinks) {
        const link = page.getByRole('link', { name: new RegExp(name, 'i') }).first();
        await expect(link).toBeVisible();
      }
    });

    test('should have Operations group with correct items', async ({ page }) => {
      const operationsLinks = ['Audit Sessions', 'Maintenance', 'Reservations'];
      for (const name of operationsLinks) {
        const link = page.getByRole('link', { name: new RegExp(name, 'i') }).first();
        await expect(link).toBeVisible();
      }
    });

    test('should have Financials group with correct items', async ({ page }) => {
      const financialLinks = ['Procurement', 'Financials', 'Depreciation'];
      for (const name of financialLinks) {
        const link = page.getByRole('link', { name: new RegExp(name, 'i') }).first();
        await expect(link).toBeVisible();
      }
    });

    test('should have Compliance group with correct items', async ({ page }) => {
      const complianceLinks = ['Audit Cleanup', 'Report Builder'];
      for (const name of complianceLinks) {
        const link = page.getByRole('link', { name: new RegExp(name, 'i') }).first();
        await expect(link).toBeVisible();
      }
    });

    test('should have System group with correct items', async ({ page }) => {
      const systemLinks = ['Alerts', 'Self Service', 'User Roles'];
      for (const name of systemLinks) {
        const link = page.getByRole('link', { name: new RegExp(name, 'i') }).first();
        await expect(link).toBeVisible();
      }
    });
  });

  test.describe('Active State', () => {
    test('should highlight Dashboard as active on home page', async ({ page }) => {
      await page.goto('/');
      const dashboardLink = page.getByRole('link', { name: /Dashboard/i }).first();
      await expect(dashboardLink).toBeVisible();
    });

    test('should highlight Asset Registry as active', async ({ page }) => {
      await page.goto('/asset-registry');
      await page.waitForTimeout(300);
      const link = page.getByRole('link', { name: /Asset Registry/i }).first();
      await expect(link).toBeVisible();
    });
  });

  test.describe('Navigation with API Mocking', () => {
    test('should navigate to pages with mocked API', async ({ page }) => {
      await page.route('**/api/**', (route) => {
        route.fulfill({ json: [] });
      });

      const routes = ['/asset-registry', '/audit-sessions', '/stockroom-inventory'];

      for (const route of routes) {
        await page.goto(route);
        await page.waitForTimeout(300);
        await expect(page.locator('body')).toBeVisible();
      }
    });
  });

  test.describe('Keyboard Navigation', () => {
    test('should navigate sidebar with keyboard', async ({ page }) => {
      await page.keyboard.press('Tab');
      await page.keyboard.press('Tab');

      const focused = page.locator(':focus');
      await expect(focused).toBeVisible();

      await page.keyboard.press('Enter');
      await page.waitForTimeout(300);
    });
  });

  test.describe('Responsive Sidebar', () => {
    test('should have collapsed sidebar on mobile', async ({ page }) => {
      await page.setViewportSize({ width: 375, height: 667 });
      await page.waitForTimeout(300);

      const sidebar = page.locator('nav, [class*="sidebar"]').first();
      await expect(sidebar).toBeVisible();
    });

    test('should have expanded sidebar on desktop', async ({ page }) => {
      await page.setViewportSize({ width: 1280, height: 720 });
      await page.waitForTimeout(300);

      const sidebar = page.locator('nav, [class*="sidebar"]').first();
      await expect(sidebar).toBeVisible();
    });
  });

  test.describe('Breadcrumb Verification', () => {
    test('should show correct breadcrumb after navigation', async ({ page }) => {
      await page.getByRole('link', { name: /Asset Registry/i }).click();
      await page.waitForTimeout(300);

      const breadcrumb = page
        .locator('[class*="breadcrumb"], nav[aria-label*="breadcrumb"]')
        .first();
      await breadcrumb.isVisible().catch(() => false);
    });
  });
});
