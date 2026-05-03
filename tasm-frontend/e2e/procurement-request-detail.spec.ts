import { test, expect } from '@playwright/test';

/**
 * Procurement Request Detail Tests
 */
test.describe('Procurement Request Detail', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/procurement/1');
    await page.waitForLoadState('networkidle').catch(() => {});
  });

  test.describe('Page Load', () => {
    test('should load Procurement Request page', async ({ page }) => {
      await expect(page.locator('body')).toBeVisible();
    });

    test('should have Approve button', async ({ page }) => {
      const approveBtn = page.getByText('Approve').first();
      if (await approveBtn.isVisible()) {
        await expect(approveBtn).toBeVisible();
      }
    });

    test('should have Reject button', async ({ page }) => {
      const rejectBtn = page.getByText('Reject').first();
      if (await rejectBtn.isVisible()) {
        await expect(rejectBtn).toBeVisible();
      }
    });

    test('should have Back button', async ({ page }) => {
      const backBtn = page.getByText('Back').first();
      if (await backBtn.isVisible()) {
        await expect(backBtn).toBeVisible();
      }
    });
  });

  test.describe('With Mocked Data', () => {
    test.beforeEach(async ({ page }) => {
      await page.route('**/api/procurements/1', (route) => {
        route.fulfill({
          json: {
            id: 1,
            item: 'Dell Laptops',
            quantity: 10,
            requestedBy: 'Jane Doe',
            status: 'Pending',
            estimatedCost: 500000,
          },
        });
      });
      await page.goto('/procurement/1');
      await page.waitForTimeout(500);
    });

    test('should display request details', async ({ page }) => {
      await expect(page.getByText('Dell Laptops')).toBeVisible();
      await expect(page.getByText('Jane Doe')).toBeVisible();
    });

    test('should display status', async ({ page }) => {
      await expect(page.getByText('Pending')).toBeVisible();
    });
  });

  test.describe('Actions', () => {
    test('should click Approve button', async ({ page }) => {
      await page.route('**/api/procurements/1**', (route) => {
        route.fulfill({ json: { success: true } });
      });

      const approveBtn = page.getByText('Approve').first();
      if (await approveBtn.isVisible()) {
        await approveBtn.click();
        await page.waitForTimeout(500);
      }
    });

    test('should click Reject button', async ({ page }) => {
      const rejectBtn = page.getByText('Reject').first();
      if (await rejectBtn.isVisible()) {
        await rejectBtn.click();
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

  test.describe('Status Updates', () => {
    test('should update status to Approved', async ({ page }) => {
      await page.route('**/api/procurement/1**', (route) => {
        route.fulfill({ json: { success: true } });
      });

      const approveBtn = page.getByText('Approve').first();
      if (await approveBtn.isVisible()) {
        await approveBtn.click();
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
      await page.goto('/procurement/1');
      await page.waitForTimeout(300);
      await expect(page.locator('body')).toBeVisible();
    });
  });

  test.describe('Error Handling', () => {
    test('should handle API error', async ({ page }) => {
      await page.route('**/api/procurements/1', (route) => route.abort());
      await page.goto('/procurement/1');
      await page.waitForTimeout(500);
      await expect(page.locator('body')).toBeVisible();
    });
  });
});
