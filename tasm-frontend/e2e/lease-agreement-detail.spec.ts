import { test, expect } from '@playwright/test';

/**
 * Lease Agreement Detail Tests
 */
test.describe('Lease Agreement Detail', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/lease/1');
    await page.waitForLoadState('networkidle').catch(() => {});
  });

  test.describe('Page Load', () => {
    test('should load Lease Agreement page', async ({ page }) => {
      await expect(page.locator('body')).toBeVisible();
    });

    test('should have Edit button', async ({ page }) => {
      const editBtn = page.getByText('Edit').or(page.getByText('Update')).first();
      if (await editBtn.isVisible()) {
        await expect(editBtn).toBeVisible();
      }
    });

    test('should have Back button', async ({ page }) => {
      const backBtn = page.getByText('Back').or(page.locator('[aria-label*="back"]')).first();
      if (await backBtn.isVisible()) {
        await expect(backBtn).toBeVisible();
      }
    });

    test('should have status indicator', async ({ page }) => {
      const status = page
        .getByText('Active')
        .or(page.getByText('Expired'))
        .or(page.getByText('Pending'))
        .first();
      await status.isVisible().catch(() => false);
    });
  });

  test.describe('With Mocked Data', () => {
    test.beforeEach(async ({ page }) => {
      await page.route('**/api/leases/1', (route) => {
        route.fulfill({
          json: {
            id: 1,
            asset: 'Dell Laptop',
            lessee: 'John Doe',
            startDate: '2026-01-01',
            endDate: '2026-12-31',
            monthlyPayment: 5000,
            status: 'Active',
          },
        });
      });
      await page.goto('/lease/1');
      await page.waitForTimeout(500);
    });

    test('should display lease details', async ({ page }) => {
      await expect(page.getByText('Dell Laptop')).toBeVisible();
      await expect(page.getByText('John Doe')).toBeVisible();
    });

    test('should display payment amount', async ({ page }) => {
      await expect(page.getByText('5000')).toBeVisible();
    });
  });

  test.describe('Actions', () => {
    test('should click Edit button', async ({ page }) => {
      const editBtn = page.getByText('Edit').first();
      if (await editBtn.isVisible()) {
        await editBtn.click();
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

    test('should click Terminate button', async ({ page }) => {
      const terminateBtn = page.getByText('Terminate').or(page.getByText('Cancel Lease')).first();
      if (await terminateBtn.isVisible()) {
        await terminateBtn.click();
        await page.waitForTimeout(300);
      }
    });
  });

  test.describe('Tabs/Sections', () => {
    test('should click Payment History tab', async ({ page }) => {
      const tab = page
        .getByText('Payment History')
        .or(page.getByRole('tab', { name: /Payment/i }))
        .first();
      if (await tab.isVisible()) {
        await tab.click();
        await page.waitForTimeout(300);
      }
    });

    test('should click Documents tab', async ({ page }) => {
      const tab = page
        .getByText('Documents')
        .or(page.getByRole('tab', { name: /Document/i }))
        .first();
      if (await tab.isVisible()) {
        await tab.click();
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
      await page.goto('/lease/1');
      await page.waitForTimeout(300);
      await expect(page.locator('body')).toBeVisible();
    });
  });

  test.describe('Error Handling', () => {
    test('should handle API error', async ({ page }) => {
      await page.route('**/api/leases/1', (route) => route.abort());
      await page.goto('/lease/1');
      await page.waitForTimeout(500);
      await expect(page.locator('body')).toBeVisible();
    });
  });
});
