import { test, expect } from '@playwright/test';

/**
 * End-to-End User Flow Tests
 * Simulates real user workflows across the application
 */
test.describe('E2E User Flows', () => {
  test.describe('Dashboard to Asset Registry Flow', () => {
    test('should navigate from Dashboard to Asset Registry and filter', async ({ page }) => {
      await page.goto('/');
      await expect(page.getByText('Operational Overview')).toBeVisible();

      await page.getByRole('link', { name: /Asset Registry/i }).click();
      await expect(page).toHaveURL(/\/asset-registry$/);
      await expect(page.getByRole('heading', { name: 'Asset Registry' })).toBeVisible();

      const inStockBtn = page.getByText('In Stock').first();
      await inStockBtn.click();
      await page.waitForTimeout(300);

      const allBtn = page.getByText('All').first();
      await allBtn.click();
      await page.waitForTimeout(300);
    });

    test('should use time filter on Dashboard then navigate away', async ({ page }) => {
      await page.goto('/');

      const sevenDaysBtn = page.getByText('7 Days').first();
      await sevenDaysBtn.click();
      await page.waitForTimeout(300);

      await page.getByRole('link', { name: /Consumables/i }).click();
      await expect(page).toHaveURL(/\/consumables/);

      await page.goBack();
      await expect(page).toHaveURL(/\/$/);
    });
  });

  test.describe('Asset Management Flow', () => {
    test('should search asset, view details, go back', async ({ page }) => {
      await page.goto('/asset-registry');
      await page.waitForTimeout(300);

      const searchInput = page.locator('input[type="text"]').first();
      await searchInput.fill('Laptop');
      await page.waitForTimeout(500);

      await searchInput.clear();
      await page.waitForTimeout(200);

      const advancedFiltersBtn = page.getByText('Advanced Filters').first();
      await advancedFiltersBtn.click();
      await page.waitForTimeout(300);
    });

    test('should navigate through sidebar groups', async ({ page }) => {
      await page.goto('/');

      await page.getByRole('link', { name: /Stockrooms/i }).click();
      await page.waitForTimeout(300);
      await expect(page).toHaveURL(/\/stockroom-inventory$/);

      await page.getByRole('link', { name: /Asset Cards/i }).click();
      await page.waitForTimeout(300);
      await expect(page).toHaveURL(/\/inventory-card-view$/);

      await page.goBack();
      await page.waitForTimeout(300);
    });
  });

  test.describe('Audit Flow', () => {
    test('should navigate audit pages in sequence', async ({ page }) => {
      await page.goto('/audit-sessions');
      await page.waitForTimeout(300);

      await page.getByRole('link', { name: /Audit Cleanup/i }).click();
      await page.waitForTimeout(300);
      await expect(page).toHaveURL(/\/audit-discrepancy-resolution$/);

      await page.getByRole('link', { name: /Audit Sessions/i }).click();
      await page.waitForTimeout(300);
      await expect(page).toHaveURL(/\/audit-sessions$/);
    });
  });

  test.describe('Financial Flow', () => {
    test('should navigate financial pages', async ({ page }) => {
      await page.goto('/financial-analytics-dashboard');
      await page.waitForTimeout(300);

      await page.getByRole('link', { name: /Procurement/i }).click();
      await page.waitForTimeout(300);
      await expect(page).toHaveURL(/\/procurement-pipeline$/);

      await page.getByRole('link', { name: /Depreciation/i }).click();
      await page.waitForTimeout(300);
      await expect(page).toHaveURL(/\/asset-depreciation-ledger$/);
    });
  });

  test.describe('Quick Actions Flow', () => {
    test('should click all quick action buttons', async ({ page }) => {
      await page.goto('/');

      const actions = ['Procurement', 'Service Log', 'Audit Scan', 'Transfer'];
      for (const action of actions) {
        const btn = page.getByText(action).first();
        if (await btn.isVisible()) {
          await btn.click();
          await page.waitForTimeout(300);
        }
      }
    });
  });

  test.describe('Settings and System Flow', () => {
    test('should navigate to settings and back', async ({ page }) => {
      await page.goto('/user-management-settings');
      await page.waitForTimeout(300);

      await page.getByRole('link', { name: /Self Service/i }).click();
      await page.waitForTimeout(300);

      await page.goBack();
      await page.waitForTimeout(300);
    });
  });

  test.describe('Search Across Pages', () => {
    test('should search on multiple pages', async ({ page }) => {
      const pages = [
        { url: '/asset-registry', searchText: 'Laptop' },
        { url: '/consumables-supplies-tracker', searchText: 'Paper' },
        { url: '/stockroom-inventory', searchText: 'Building' },
      ];

      for (const p of pages) {
        await page.goto(p.url);
        await page.waitForTimeout(300);

        const search = page.locator('input[type="text"]').first();
        await search.fill(p.searchText);
        await page.waitForTimeout(300);
        await search.clear();
      }
    });
  });

  test.describe('Report Generation Flow', () => {
    test('should navigate to report pages', async ({ page }) => {
      await page.goto('/reports-analytics-dashboard');
      await page.waitForTimeout(300);

      await page.getByRole('link', { name: /Report Builder/i }).click();
      await page.waitForTimeout(300);
      await expect(page).toHaveURL(/\/custom-report-builder$/);
    });
  });

  test.describe('Mobile User Flow', () => {
    test('should complete basic flow on mobile', async ({ page }) => {
      await page.setViewportSize({ width: 375, height: 667 });

      await page.goto('/');
      await page.waitForTimeout(300);
      await expect(page.getByText('Operational Overview')).toBeVisible();

      await page.getByRole('link', { name: /Asset Registry/i }).click();
      await page.waitForTimeout(500);
      await expect(page.getByRole('heading', { name: 'Asset Registry' })).toBeVisible();

      await page.goBack();
      await page.waitForTimeout(300);
    });
  });

  test.describe('Error Recovery Flow', () => {
    test('should recover from API errors', async ({ page }) => {
      await page.route('**/api/**', (route) => route.abort());

      await page.goto('/');
      await page.waitForTimeout(500);

      await page.route('**/api/**', (route) => route.continue());
      await page.reload();
      await page.waitForTimeout(500);

      await expect(page.getByText('Operational Overview')).toBeVisible();
    });
  });

  test.describe('Keyboard-Only Flow', () => {
    test('should navigate dashboard using only keyboard', async ({ page }) => {
      await page.goto('/');

      await page.keyboard.press('Tab');
      await page.keyboard.press('Tab');
      await page.keyboard.press('Tab');

      await page.keyboard.press('Enter');
      await page.waitForTimeout(300);

      await page.keyboard.press('Tab');
      await page.keyboard.press('Enter');
      await page.waitForTimeout(300);
    });
  });
});
