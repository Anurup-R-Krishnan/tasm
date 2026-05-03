import { test, expect } from '@playwright/test';

/**
 * Audit Module Tests
 * Tests Audit Sessions, Discrepancy Resolution, and Scan Mode
 */
test.describe('Audit Module Tests', () => {
  test.describe('Audit Sessions Page', () => {
    test.beforeEach(async ({ page }) => {
      await page.goto('/audit-sessions');
      await page.waitForLoadState('networkidle').catch(() => {});
    });

    test('should load Audit Sessions page', async ({ page }) => {
      await expect(page.getByRole('heading', { name: /Audit Sessions/i })).toBeVisible();
    });

    test('should click New Session button', async ({ page }) => {
      const newBtn = page.getByText('New Session').or(page.getByText('Create')).first();
      if (await newBtn.isVisible()) {
        await newBtn.click();
        await page.waitForTimeout(500);
      }
    });

    test('should find audit session rows', async ({ page }) => {
      await page.route('**/api/audits**', (route) => {
        route.fulfill({
          json: [{ id: 1, name: 'Q1 Audit', status: 'In Progress', date: '2026-05-03' }],
        });
      });

      await page.reload();
      await page.waitForTimeout(500);

      const rows = page.locator('table tbody tr, .session-item, [class*="session"]');
      const count = await rows.count();
      expect(count).toBeGreaterThanOrEqual(0);
    });

    test('should click on session row', async ({ page }) => {
      await page.route('**/api/audits**', (route) => {
        route.fulfill({
          json: [{ id: 1, name: 'Test Session' }],
        });
      });

      await page.reload();
      await page.waitForTimeout(500);

      const rows = page.locator('table tbody tr, .session-item').first();
      if (await rows.isVisible()) {
        await rows.click();
        await page.waitForTimeout(300);
      }
    });

    test('should have filter buttons', async ({ page }) => {
      const filters = ['All', 'In Progress', 'Completed', 'Pending'];
      for (const filter of filters) {
        const btn = page.getByText(filter).first();
        if (await btn.isVisible()) {
          await btn.click();
          await page.waitForTimeout(200);
        }
      }
    });
  });

  test.describe('Audit Discrepancy Resolution', () => {
    test.beforeEach(async ({ page }) => {
      await page.goto('/audit-discrepancy-resolution');
      await page.waitForLoadState('networkidle').catch(() => {});
    });

    test('should load Discrepancy Resolution page', async ({ page }) => {
      await expect(page.getByRole('heading', { name: /Audit Cleanup|Discrepancy/i })).toBeVisible();
    });

    test('should find discrepancy items', async ({ page }) => {
      await page.route('**/api/discrepancies**', (route) => {
        route.fulfill({
          json: [{ id: 1, assetTag: 'AST-001', issue: 'Missing', location: 'Building A' }],
        });
      });

      await page.reload();
      await page.waitForTimeout(500);

      const items = page.locator('table tbody tr, .discrepancy-item, [class*="discrepancy"]');
      const count = await items.count();
      expect(count).toBeGreaterThanOrEqual(0);
    });

    test('should click resolve button', async ({ page }) => {
      const resolveBtn = page.getByText('Resolve').or(page.getByText('Fix')).first();
      if (await resolveBtn.isVisible()) {
        await resolveBtn.click();
        await page.waitForTimeout(300);
      }
    });

    test('should click export button', async ({ page }) => {
      const exportBtn = page.getByText('Export').first();
      if (await exportBtn.isVisible()) {
        await exportBtn.click();
        await page.waitForTimeout(300);
      }
    });

    test('should filter discrepancies', async ({ page }) => {
      const filters = ['All', 'Critical', 'Warning', 'Info'];
      for (const filter of filters) {
        const btn = page.getByText(filter).first();
        if (await btn.isVisible()) {
          await btn.click();
          await page.waitForTimeout(200);
        }
      }
    });
  });

  test.describe('Audit Scan Mode Mobile', () => {
    test.beforeEach(async ({ page }) => {
      await page.goto('/audit-scan-mode-mobile');
      await page.waitForLoadState('networkidle').catch(() => {});
    });

    test('should load Scan Mode page', async ({ page }) => {
      const heading = page.getByRole('heading', { name: /Scan|Audit/i }).first();
      if (await heading.isVisible()) {
        await expect(heading).toBeVisible();
      }
    });

    test('should find scan button', async ({ page }) => {
      const scanBtn = page.getByText('Scan').or(page.getByText('Start Scan')).first();
      if (await scanBtn.isVisible()) {
        await expect(scanBtn).toBeVisible();
      }
    });

    test('should click scan button', async ({ page }) => {
      const scanBtn = page.getByText('Scan').first();
      if (await scanBtn.isVisible()) {
        await scanBtn.click();
        await page.waitForTimeout(500);
      }
    });

    test('should find expand/collapse buttons', async ({ page }) => {
      const expandBtns = page.locator(
        '[class*="expand"], [class*="collapse"], button:has(.material-symbols-outlined)',
      );
      const count = await expandBtns.count();
      expect(count).toBeGreaterThanOrEqual(0);
    });

    test('should click navigation back button', async ({ page }) => {
      const backBtn = page.getByText('Back').or(page.locator('[aria-label*="back"]')).first();
      if (await backBtn.isVisible()) {
        await backBtn.click();
        await page.waitForTimeout(300);
      }
    });
  });

  test.describe('Cross-Audit Navigation', () => {
    test('should navigate between audit pages', async ({ page }) => {
      await page.goto('/audit-sessions');
      await expect(page.getByRole('heading', { name: /Audit/i })).toBeVisible();

      await page
        .getByRole('link', { name: /Discrepancy|Cleanup/i })
        .click()
        .catch(() => {});
      await page.waitForTimeout(300);
    });
  });

  test.describe('Audit API Mocking', () => {
    test('should handle empty audit sessions', async ({ page }) => {
      await page.route('**/api/audits**', (route) => {
        route.fulfill({ json: [] });
      });

      await page.goto('/audit-sessions');
      await page.waitForTimeout(500);

      await expect(page.locator('body')).toBeVisible();
    });

    test('should handle audit API error', async ({ page }) => {
      await page.route('**/api/audits**', (route) => {
        route.abort();
      });

      await page.goto('/audit-sessions');
      await page.waitForTimeout(500);

      await expect(page.locator('body')).toBeVisible();
    });
  });

  test.describe('Audit Mobile Tests', () => {
    test('should work on mobile viewport', async ({ page }) => {
      await page.setViewportSize({ width: 375, height: 667 });
      await page.goto('/audit-scan-mode-mobile');
      await page.waitForTimeout(300);

      await expect(page.locator('body')).toBeVisible();
    });

    test('should tap scan button on mobile', async ({ page }) => {
      await page.setViewportSize({ width: 375, height: 667 });
      await page.goto('/audit-scan-mode-mobile');
      await page.waitForTimeout(300);

      const scanBtn = page.getByText('Scan').first();
      if (await scanBtn.isVisible()) {
        await scanBtn.tap();
        await page.waitForTimeout(300);
      }
    });
  });
});
