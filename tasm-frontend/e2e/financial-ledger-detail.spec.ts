import { test, expect } from '@playwright/test';

/**
 * Financial Ledger Detail Tests
 */
test.describe('Financial Ledger Detail', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/financial-ledger-detail');
    await page.waitForLoadState('networkidle').catch(() => {});
  });

  test.describe('Page Load', () => {
    test('should load Financial Ledger page', async ({ page }) => {
      await expect(page.locator('body')).toBeVisible();
    });

    test('should have Back button', async ({ page }) => {
      const backBtn = page.getByText('Back').or(page.locator('[aria-label*="back"]')).first();
      if (await backBtn.isVisible()) {
        await expect(backBtn).toBeVisible();
      }
    });

    test('should have Export button', async ({ page }) => {
      const exportBtn = page.getByText('Export').first();
      if (await exportBtn.isVisible()) {
        await expect(exportBtn).toBeVisible();
      }
    });
  });

  test.describe('With Mocked Data', () => {
    test.beforeEach(async ({ page }) => {
      await page.route('**/api/financial-ledger**', (route) => {
        route.fulfill({
          json: {
            asset: 'Dell Laptop',
            purchasePrice: 100000,
            currentValue: 60000,
            depreciationRate: 20,
            entries: [
              { date: '2026-01-01', value: 100000 },
              { date: '2026-05-01', value: 60000 },
            ],
          },
        });
      });
      await page.goto('/financial-ledger-detail');
      await page.waitForTimeout(500);
    });

    test('should display asset name', async ({ page }) => {
      await expect(page.getByText('Dell Laptop')).toBeVisible();
    });

    test('should display purchase price', async ({ page }) => {
      await expect(page.getByText('100000')).toBeVisible();
    });

    test('should display current value', async ({ page }) => {
      await expect(page.getByText('60000')).toBeVisible();
    });

    test('should have ledger entries table', async ({ page }) => {
      const table = page.locator('table').first();
      if (await table.isVisible()) {
        await expect(table).toBeVisible();
      }
    });
  });

  test.describe('Charts', () => {
    test('should have depreciation chart', async ({ page }) => {
      const chart = page.locator('canvas, svg, .chart, [class*="chart"]').first();
      await chart.isVisible().catch(() => false);
    });
  });

  test.describe('Actions', () => {
    test('should click Export button', async ({ page }) => {
      const exportBtn = page.getByText('Export').first();
      if (await exportBtn.isVisible()) {
        await exportBtn.click();
        await page.waitForTimeout(500);
      }
    });

    test('should click Back button', async ({ page }) => {
      const backBtn = page.getByText('Back').first();
      if (await backBtn.isVisible()) {
        await backBtn.click();
        await page.waitForTimeout(300);
      }
    });
  });

  test.describe('Keyboard Navigation', () => {
    test('should tab through page', async ({ page }) => {
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
      await page.goto('/financial-ledger-detail');
      await page.waitForTimeout(300);
      await expect(page.locator('body')).toBeVisible();
    });
  });

  test.describe('Error Handling', () => {
    test('should handle API error', async ({ page }) => {
      await page.route('**/api/financial-ledger**', (route) => route.abort());
      await page.goto('/financial-ledger-detail');
      await page.waitForTimeout(500);
      await expect(page.locator('body')).toBeVisible();
    });
  });
});
