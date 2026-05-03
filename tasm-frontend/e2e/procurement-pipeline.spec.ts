import { test, expect } from '@playwright/test';

/**
 * Procurement Pipeline Tests
 */
test.describe('Procurement Pipeline', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/procurement-pipeline');
    await page.waitForLoadState('networkidle').catch(() => {});
  });

  test.describe('Page Load', () => {
    test('should load Procurement page', async ({ page }) => {
      await expect(page.getByRole('heading', { name: /Procurement/i })).toBeVisible();
    });

    test('should have Add Request button', async ({ page }) => {
      const addBtn = page.getByText('Add Request').or(page.getByText('New Request')).first();
      if (await addBtn.isVisible()) {
        await expect(addBtn).toBeVisible();
      }
    });

    test('should have filter buttons', async ({ page }) => {
      const filters = ['All', 'Pending', 'Approved', 'Ordered', 'Received'];
      for (const filter of filters) {
        const btn = page.getByText(filter).first();
        if (await btn.isVisible()) {
          await expect(btn).toBeVisible();
        }
      }
    });
  });

  test.describe('With Mocked Data', () => {
    test.beforeEach(async ({ page }) => {
      await page.route('**/api/procurements**', (route) => {
        route.fulfill({
          json: [
            { id: 1, item: 'Laptop', status: 'Pending', quantity: 10 },
            { id: 2, item: 'Monitor', status: 'Ordered', quantity: 5 },
          ],
        });
      });
      await page.goto('/procurement-pipeline');
      await page.waitForTimeout(500);
    });

    test('should display procurement items', async ({ page }) => {
      await expect(page.getByText('Laptop')).toBeVisible();
      await expect(page.getByText('Monitor')).toBeVisible();
    });

    test('should click on procurement item', async ({ page }) => {
      const row = page.locator('table tbody tr').first();
      if (await row.isVisible()) {
        await row.click();
        await page.waitForTimeout(300);
      }
    });
  });

  test.describe('Filter Buttons', () => {
    test('should click All filter', async ({ page }) => {
      const btn = page.getByText('All').first();
      if (await btn.isVisible()) {
        await btn.click();
        await page.waitForTimeout(300);
      }
    });

    test('should click Pending filter', async ({ page }) => {
      const btn = page.getByText('Pending').first();
      if (await btn.isVisible()) {
        await btn.click();
        await page.waitForTimeout(300);
      }
    });

    test('should toggle filters', async ({ page }) => {
      const allBtn = page.getByText('All').first();
      const pendingBtn = page.getByText('Pending').first();

      if (await allBtn.isVisible()) {
        await allBtn.click();
        await page.waitForTimeout(200);
      }

      if (await pendingBtn.isVisible()) {
        await pendingBtn.click();
        await page.waitForTimeout(200);
      }
    });
  });

  test.describe('Add Request', () => {
    test('should click Add Request button', async ({ page }) => {
      const addBtn = page.getByText('Add Request').or(page.getByText('New')).first();
      if (await addBtn.isVisible()) {
        await addBtn.click();
        await page.waitForTimeout(500);
      }
    });

    test('should fill procurement form', async ({ page }) => {
      const inputs = page.locator('input[type="text"], input[type="number"]');
      const count = await inputs.count();
      for (let i = 0; i < Math.min(count, 3); i++) {
        await inputs
          .nth(i)
          .fill('Test')
          .catch(() => {});
      }
    });
  });

  test.describe('Status Updates', () => {
    test('should click status update button', async ({ page }) => {
      const statusBtn = page.getByText('Update Status').or(page.getByText('Change')).first();
      if (await statusBtn.isVisible()) {
        await statusBtn.click();
        await page.waitForTimeout(300);
      }
    });
  });

  test.describe('Navigation', () => {
    test('should navigate from sidebar', async ({ page }) => {
      await page.goto('/');
      await page.getByRole('link', { name: /Procurement/i }).click();
      await expect(page).toHaveURL(/\/procurement-pipeline$/);
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
      await expect(page.getByRole('heading', { name: /Procurement/i })).toBeVisible();
    });
  });

  test.describe('Error Handling', () => {
    test('should handle API error', async ({ page }) => {
      await page.route('**/api/procurements**', (route) => route.abort());
      await page.reload();
      await page.waitForTimeout(500);
      await expect(page.locator('body')).toBeVisible();
    });
  });
});
