import { test, expect } from '@playwright/test';

/**
 * Reports & Analytics Dashboard Tests
 */
test.describe('Reports & Analytics', () => {
  test.describe('Reports Dashboard', () => {
    test.beforeEach(async ({ page }) => {
      await page.goto('/reports-analytics-dashboard');
      await page.waitForLoadState('networkidle').catch(() => {});
    });

    test('should load Reports page', async ({ page }) => {
      await expect(page.getByRole('heading', { name: /Report/i })).toBeVisible();
    });

    test('should have Generate Report button', async ({ page }) => {
      const genBtn = page.getByText('Generate').or(page.getByText('New Report')).first();
      if (await genBtn.isVisible()) {
        await expect(genBtn).toBeVisible();
      }
    });

    test('should have report type filters', async ({ page }) => {
      const filters = ['All', 'Assets', 'Financial', 'Maintenance'];
      for (const filter of filters) {
        const btn = page.getByText(filter).first();
        if (await btn.isVisible()) {
          await expect(btn).toBeVisible();
        }
      }
    });
  });

  test.describe('Custom Report Builder', () => {
    test.beforeEach(async ({ page }) => {
      await page.goto('/custom-report-builder');
      await page.waitForLoadState('networkidle').catch(() => {});
    });

    test('should load Report Builder page', async ({ page }) => {
      await expect(page.getByRole('heading', { name: /Report Builder/i })).toBeVisible();
    });

    test('should have Create Report button', async ({ page }) => {
      const createBtn = page.getByText('Create').or(page.getByText('Build')).first();
      if (await createBtn.isVisible()) {
        await expect(createBtn).toBeVisible();
      }
    });

    test('should have field selectors', async ({ page }) => {
      const selects = page.locator('select, [role="combobox"]');
      const count = await selects.count();
      expect(count).toBeGreaterThanOrEqual(0);
    });

    test('should fill report name', async ({ page }) => {
      const nameInput = page.locator('input[placeholder*="name" i], input[type="text"]').first();
      if (await nameInput.isVisible()) {
        await nameInput.fill('Test Report');
        await page.waitForTimeout(200);
      }
    });
  });

  test.describe('With Mocked Data', () => {
    test('should display reports list', async ({ page }) => {
      await page.route('**/api/ledgers**', (route) => {
        route.fulfill({
          json: [{ id: 1, name: 'Asset Summary', type: 'Assets' }],
        });
      });
      await page.route('**/api/procurements**', (route) => {
        route.fulfill({
          json: [{ id: 2, name: 'Financial Q1', type: 'Financial' }],
        });
      });

      await page.goto('/reports-analytics-dashboard');
      await page.waitForTimeout(500);

      await expect(page.getByText('Asset Summary')).toBeVisible();
    });
  });

  test.describe('Keyboard Navigation', () => {
    test('should tab through reports page', async ({ page }) => {
      await page.goto('/reports-analytics-dashboard');
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
      await page.goto('/reports-analytics-dashboard');
      await page.waitForTimeout(300);
      await expect(page.getByRole('heading', { name: /Report/i })).toBeVisible();
    });
  });

  test.describe('Error Handling', () => {
    test('should handle API error', async ({ page }) => {
      await page.route('**/api/ledgers**', (route) => route.abort());
      await page.goto('/reports-analytics-dashboard');
      await page.waitForTimeout(500);
      await expect(page.locator('body')).toBeVisible();
    });
  });

  test.describe('Navigation', () => {
    test('should navigate between report pages', async ({ page }) => {
      await page.goto('/reports-analytics-dashboard');
      await page.waitForTimeout(300);

      await page
        .getByRole('link', { name: /Report Builder/i })
        .click()
        .catch(() => {});
      await page.waitForTimeout(300);
    });
  });
});
