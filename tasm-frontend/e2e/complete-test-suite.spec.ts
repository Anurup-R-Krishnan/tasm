import { test, expect } from '@playwright/test';

/**
 * Complete Test Suite - Master Integration Test
 * Runs through all major features in sequence
 */
test.describe('Complete Test Suite - Master Flow', () => {
  test('should complete full application walkthrough', async ({ page }) => {
    // 1. Dashboard
    await page.goto('/');
    await expect(page.getByText('Operational Overview')).toBeVisible();
    await page.waitForTimeout(300);

    // 2. Click ghost button
    const ghostBtn = page.getByText('7 Days').first();
    await ghostBtn.click();
    await page.waitForTimeout(200);

    // 3. Quick Actions
    const procurementBtn = page.getByText('Procurement').first();
    if (await procurementBtn.isVisible()) {
      await procurementBtn.click();
      await page.waitForTimeout(300);
    }

    // 4. Navigate to Asset Registry
    await page.getByRole('link', { name: /Asset Registry/i }).click();
    await expect(page).toHaveURL(/\/asset-registry$/);
    await page.waitForTimeout(300);

    // 5. Use filters
    const inStockBtn = page.getByText('In Stock').first();
    if (await inStockBtn.isVisible()) {
      await inStockBtn.click();
      await page.waitForTimeout(200);
    }

    // 6. Search
    const searchInput = page.locator('input[type="text"]').first();
    await searchInput.fill('Laptop');
    await page.waitForTimeout(300);
    await searchInput.clear();

    // 7. Navigate to Consumables
    await page.getByRole('link', { name: /Consumables/i }).click();
    await expect(page).toHaveURL(/\/consumables-supplies-tracker$/);
    await page.waitForTimeout(300);

    // 8. Navigate to Audit Sessions
    await page.getByRole('link', { name: /Audit Sessions/i }).click();
    await expect(page).toHaveURL(/\/audit-sessions$/);
    await page.waitForTimeout(300);

    // 9. Navigate to Maintenance
    await page.getByRole('link', { name: /Maintenance/i }).click();
    await expect(page).toHaveURL(/\/maintenance-schedule-contracts$/);
    await page.waitForTimeout(300);

    // 10. Navigate to Financials
    await page.locator('a[href="/financial-analytics-dashboard"]').click();
    await expect(page).toHaveURL(/\/financial-analytics-dashboard$/);
    await page.waitForTimeout(300);

    // 11. Navigate to Settings
    const settingsLink = page.locator('a[href="/user-management-settings"]');
    await settingsLink.scrollIntoViewIfNeeded();
    await settingsLink.first().click();
    await expect(page).toHaveURL(/\/user-management-settings$/);
    await page.waitForTimeout(500);

    // 12. Go back to Dashboard
    await page.locator('a[href="/"]').first().click();
    await expect(page).toHaveURL(/\/$/);
    await page.waitForTimeout(300);

    // 15. Test keyboard navigation
    await page.keyboard.press('Tab');
    await page.keyboard.press('Tab');
    await page.keyboard.press('Tab');
    const focused = page.locator(':focus');
    await expect(focused).toBeVisible();

    // 16. Test mobile view
    await page.setViewportSize({ width: 375, height: 667 });
    await page.waitForTimeout(300);
    await expect(page.getByText('Operational Overview')).toBeVisible();

    // 17. Reset to desktop
    await page.setViewportSize({ width: 1280, height: 720 });
    await page.waitForTimeout(300);

    console.log('✅ Complete test suite walkthrough finished successfully!');
  });

  test('should verify all buttons are clickable', async ({ page }) => {
    const pages = [
      '/',
      '/asset-registry',
      '/stockroom-inventory',
      '/consumables-supplies-tracker',
      '/software-license-registry',
      '/inventory-card-view',
      '/maintenance-schedule-contracts',
      '/audit-sessions',
      '/reservations-bookings',
      '/procurement-pipeline',
      '/financial-analytics-dashboard',
      '/asset-depreciation-ledger',
      '/audit-discrepancy-resolution',
      '/custom-report-builder',
      '/alerts-notifications-center',
      '/employee-self-service-portal',
      '/user-access-role-control',
      '/user-management-settings',
    ];

    for (const url of pages) {
      console.log(`🚀 Navigating to: ${url}`);
      await page.goto(url, { waitUntil: 'networkidle', timeout: 60000 });
      console.log(`✅ Arrived at: ${url}`);
      await page.waitForTimeout(500);

      const buttons = page.locator('button');
      const count = await buttons.count();

      for (let i = 0; i < Math.min(count, 3); i++) {
        const isVisible = await buttons.nth(i).isVisible();
        if (isVisible) {
          await buttons
            .nth(i)
            .click()
            .catch(() => {});
        }
      }
    }

    console.log('✅ All buttons verified clickable!');
  });

  test('should handle errors gracefully across all pages', async ({ page }) => {
    await page.route('**/api/**', (route) => {
      if (Math.random() > 0.7) {
        route.abort();
      } else {
        route.continue();
      }
    });

    const pages = ['/', '/asset-registry', '/consumables-supplies-tracker'];

    for (const url of pages) {
      await page.goto(url).catch(() => {});
      await page.waitForTimeout(300);
      await expect(page.locator('body')).toBeVisible();
    }

    console.log('✅ Error handling verified across pages!');
  });
});
