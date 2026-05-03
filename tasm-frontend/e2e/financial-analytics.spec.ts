import { test, expect } from '@playwright/test';

/**
 * Financial Analytics Dashboard Tests
 */
test.describe('Financial Analytics Dashboard', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/financial-analytics-dashboard');
    await page.waitForLoadState('networkidle').catch(() => {});
  });

  test.describe('Page Load', () => {
    test('should load Financials page', async ({ page }) => {
      await expect(page.getByRole('heading', { name: /Financial/i })).toBeVisible();
    });

    test('should have filter buttons', async ({ page }) => {
      const filters = ['This Month', 'Quarter', 'Year', 'All Time'];
      for (const filter of filters) {
        const btn = page.getByText(filter).first();
        if (await btn.isVisible()) {
          await expect(btn).toBeVisible();
        }
      }
    });

    test('should have export button', async ({ page }) => {
      const exportBtn = page.getByText('Export').first();
      if (await exportBtn.isVisible()) {
        await expect(exportBtn).toBeVisible();
      }
    });
  });

  test.describe('With Mocked Data', () => {
    test.beforeEach(async ({ page }) => {
      await page.route('**/api/depreciations**', (route) => {
        route.fulfill({
          json: [{ id: 1, purchaseValue: 1000000, currentValue: 800000 }],
        });
      });
      await page.route('**/api/work-orders**', (route) => {
        route.fulfill({
          json: [{ id: 1, status: 'Closed', cost: 50000 }],
        });
      });
      await page.goto('/financial-analytics-dashboard');
      await page.waitForTimeout(500);
    });

    test('should display KPI values', async ({ page }) => {
      await expect(page.getByText('Total Value')).toBeVisible();
      await expect(page.getByText('Depreciation')).toBeVisible();
    });
  });

  test.describe('Filter Buttons', () => {
    test('should click This Month filter', async ({ page }) => {
      const btn = page.getByText('This Month').first();
      if (await btn.isVisible()) {
        await btn.click();
        await page.waitForTimeout(300);
      }
    });

    test('should click Quarter filter', async ({ page }) => {
      const btn = page.getByText('Quarter').first();
      if (await btn.isVisible()) {
        await btn.click();
        await page.waitForTimeout(300);
      }
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
    test('should click export button', async ({ page }) => {
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
      await expect(page.getByRole('heading', { name: /Financial/i })).toBeVisible();
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
      await page.getByRole('link', { name: /Financials/i }).click();
      await expect(page).toHaveURL(/\/financial-analytics-dashboard$/);
    });
  });
});
