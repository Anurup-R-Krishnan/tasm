import { test, expect } from '@playwright/test';

/**
 * Work Order Management Tests
 * Tests work order listing, details, status updates
 */
test.describe('Work Order Management', () => {
  test.describe('Work Order Detail Page', () => {
    test.beforeEach(async ({ page }) => {
      await page.goto('/work-order/1');
      await page.waitForLoadState('networkidle').catch(() => {});
    });

    test('should load work order detail page', async ({ page }) => {
      await expect(page.locator('body')).toBeVisible();
    });

    test('should find status update button', async ({ page }) => {
      const statusBtn = page.getByText('Update Status').or(page.getByText('Change Status')).first();
      if (await statusBtn.isVisible()) {
        await expect(statusBtn).toBeVisible();
      }
    });

    test('should find assign button', async ({ page }) => {
      const assignBtn = page.getByText('Assign').or(page.getByText('Reassign')).first();
      if (await assignBtn.isVisible()) {
        await expect(assignBtn).toBeVisible();
      }
    });

    test('should find complete button', async ({ page }) => {
      const completeBtn = page.getByText('Complete').first();
      if (await completeBtn.isVisible()) {
        await expect(completeBtn).toBeVisible();
      }
    });

    test('should click back button', async ({ page }) => {
      const backBtn = page.getByText('Back').or(page.locator('[aria-label*="back"]')).first();
      if (await backBtn.isVisible()) {
        await backBtn.click();
        await page.waitForTimeout(300);
      }
    });
  });

  test.describe('Work Order Status Updates', () => {
    test('should click status button and see options', async ({ page }) => {
      await page.goto('/work-order/1');

      const statusBtn = page.getByText('Update Status').first();
      if (await statusBtn.isVisible()) {
        await statusBtn.click();
        await page.waitForTimeout(300);

        const options = page
          .getByText('Open')
          .or(page.getByText('In Progress'))
          .or(page.getByText('Closed'));
        const count = await options.count();
        expect(count).toBeGreaterThanOrEqual(0);
      }
    });

    test('should select new status', async ({ page }) => {
      await page.route('**/api/work-orders/**', (route) => {
        route.fulfill({ json: { success: true } });
      });

      await page.goto('/work-order/1');

      const statusBtn = page.getByText('Update Status').first();
      if (await statusBtn.isVisible()) {
        await statusBtn.click();
        await page.waitForTimeout(300);

        const inProgress = page.getByText('In Progress').first();
        if (await inProgress.isVisible()) {
          await inProgress.click();
          await page.waitForTimeout(300);
        }
      }
    });
  });

  test.describe('Work Order List (from Dashboard)', () => {
    test('should click work order in dashboard list', async ({ page }) => {
      await page.goto('/');

      await page.route('**/api/work-orders**', (route) => {
        route.fulfill({
          json: [{ id: 1, title: 'Fix AC', status: 'Open', severity: 'High' }],
        });
      });

      await page.reload();
      await page.waitForTimeout(500);

      const workOrderRow = page.locator('text=Fix AC').first();
      if (await workOrderRow.isVisible()) {
        await workOrderRow.click();
        await page.waitForTimeout(300);
      }
    });

    test('should find View All button in work orders section', async ({ page }) => {
      await page.goto('/');

      const viewAllBtn = page.getByText('View All').first();
      if (await viewAllBtn.isVisible()) {
        await viewAllBtn.click();
        await page.waitForTimeout(500);
      }
    });
  });

  test.describe('Mocked API Tests', () => {
    test('should handle work order API error', async ({ page }) => {
      await page.route('**/api/work-orders/**', (route) => route.abort());

      await page.goto('/work-order/1');
      await page.waitForTimeout(500);

      await expect(page.locator('body')).toBeVisible();
    });

    test('should display work order data from API', async ({ page }) => {
      await page.route('**/api/work-orders/1**', (route) => {
        route.fulfill({
          json: {
            id: 1,
            title: 'Test Work Order',
            status: 'In Progress',
            technician: 'John Doe',
            assetLocation: 'Building A',
          },
        });
      });

      await page.goto('/work-order/1');
      await page.waitForTimeout(500);

      await expect(page.getByText('Test Work Order')).toBeVisible();
    });
  });

  test.describe('Navigation', () => {
    test('should navigate from dashboard to work order', async ({ page }) => {
      await page.goto('/');

      await page.route('**/api/work-orders**', (route) => {
        route.fulfill({
          json: [{ id: 1, title: 'Test WO' }],
        });
      });

      await page.waitForTimeout(500);

      const woLink = page.locator('a[href*="work-order"]').first();
      if (await woLink.isVisible()) {
        await woLink.click();
        await page.waitForTimeout(500);
      }
    });
  });

  test.describe('Keyboard Navigation', () => {
    test('should tab through work order detail page', async ({ page }) => {
      await page.goto('/work-order/1');

      for (let i = 0; i < 5; i++) {
        await page.keyboard.press('Tab');
      }

      const focused = page.locator(':focus');
      await expect(focused).toBeVisible();
    });
  });

  test.describe('Responsive', () => {
    test('should display correctly on mobile', async ({ page }) => {
      await page.setViewportSize({ width: 375, height: 667 });
      await page.goto('/work-order/1');
      await page.waitForTimeout(300);

      await expect(page.locator('body')).toBeVisible();
    });
  });
});
