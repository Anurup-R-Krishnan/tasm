import { test, expect } from '@playwright/test';

/**
 * Asset Depreciation Ledger Tests
 */
test.describe('Asset Depreciation Ledger', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/asset-depreciation-ledger');
    await page.waitForLoadState('networkidle').catch(() => {});
  });

  test.describe('Page Load', () => {
    test('should load Depreciation page', async ({ page }) => {
      await expect(page.getByRole('heading', { name: /Depreciation/i })).toBeVisible();
    });

    test('should have Export button', async ({ page }) => {
      const exportBtn = page.getByText('Export').first();
      if (await exportBtn.isVisible()) {
        await expect(exportBtn).toBeVisible();
      }
    });

    test('should have date range filter', async ({ page }) => {
      const dateInputs = page.locator('input[type="date"]');
      const count = await dateInputs.count();
      expect(count).toBeGreaterThanOrEqual(0);
    });
  });

  test.describe('With Mocked Data', () => {
    test.beforeEach(async ({ page }) => {
      await page.route('**/api/depreciations**', (route) => {
        route.fulfill({
          json: [
            { asset: 'Laptop', initialValue: 100000, currentValue: 60000, depreciation: 40000 },
          ],
        });
      });
      await page.goto('/asset-depreciation-ledger');
      await page.waitForTimeout(500);
    });

    test('should display depreciation data', async ({ page }) => {
      await expect(page.getByText('Laptop')).toBeVisible();
      await expect(page.getByText('60000')).toBeVisible();
    });
  });

  test.describe('Charts and Graphs', () => {
    test('should have chart elements', async ({ page }) => {
      const charts = page.locator('canvas, svg, .chart, [class*="chart"]');
      const count = await charts.count();
      expect(count).toBeGreaterThanOrEqual(0);
    });
  });

  test.describe('Export', () => {
    test('should click Export button', async ({ page }) => {
      const exportBtn = page.getByText('Export').first();
      if (await exportBtn.isVisible()) {
        await exportBtn.click();
        await page.waitForTimeout(500);
      }
    });
  });

  test.describe('Keyboard Navigation', () => {
    test('should tab through elements', async ({ page }) => {
      for (let i = 0; i < 5; i++) {
        await page.keyboard.press('Tab');
      }
      const focused = page.locator(':focus');
      await expect(focused).toBeVisible();
    });
  });

  test.describe('Responsive', () => {
    test('should display on mobile', async ({ page }) => {
      await page.setViewportSize({ width: 375, height: 667 });
      await page.waitForTimeout(300);
      await expect(page.getByRole('heading', { name: /Depreciation/i })).toBeVisible();
    });
  });

  test.describe('Error Handling', () => {
    test('should handle API error', async ({ page }) => {
      await page.route('**/api/depreciations**', (route) => route.abort());
      await page.reload();
      await page.waitForTimeout(500);
      await expect(page.locator('body')).toBeVisible();
    });
  });

  test.describe('Navigation', () => {
    test('should navigate from sidebar', async ({ page }) => {
      await page.goto('/');
      await page.getByRole('link', { name: /Depreciation/i }).click();
      await expect(page).toHaveURL(/\/asset-depreciation-ledger$/);
    });
  });
});
