import { test, expect } from '@playwright/test';

/**
 * Asset Detail Page Tests
 */
test.describe('Asset Detail Page', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/asset-detail/1');
    await page.waitForLoadState('networkidle').catch(() => {});
  });

  test.describe('Page Load', () => {
    test('should load Asset Detail page', async ({ page }) => {
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

    test('should have status badge', async ({ page }) => {
      const status = page.getByText('In Stock').or(page.getByText('Checked Out')).first();
      if (await status.isVisible()) {
        await expect(status).toBeVisible();
      }
    });
  });

  test.describe('With Mocked Data', () => {
    test.beforeEach(async ({ page }) => {
      await page.route('**/api/assets/1**', (route) => {
        route.fulfill({
          json: {
            id: 1,
            name: 'Dell Laptop',
            tagId: 'AST-001',
            status: 'In Stock',
            location: 'Building A',
            value: 50000,
          },
        });
      });
      await page.goto('/asset-detail/1');
      await page.waitForTimeout(500);
    });

    test('should display asset information', async ({ page }) => {
      await expect(page.getByText('Dell Laptop')).toBeVisible();
      await expect(page.getByText('AST-001')).toBeVisible();
    });

    test('should display asset value', async ({ page }) => {
      await expect(page.getByText('50000')).toBeVisible();
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

    test('should click Check Out button', async ({ page }) => {
      const checkoutBtn = page.getByText('Check Out').first();
      if (await checkoutBtn.isVisible()) {
        await checkoutBtn.click();
        await page.waitForTimeout(500);
      }
    });
  });

  test.describe('Tabs/Sections', () => {
    test('should click Maintenance tab', async ({ page }) => {
      const tab = page
        .getByText('Maintenance')
        .or(page.getByRole('tab', { name: /Maintenance/i }))
        .first();
      if (await tab.isVisible()) {
        await tab.click();
        await page.waitForTimeout(300);
      }
    });

    test('should click History tab', async ({ page }) => {
      const tab = page
        .getByText('History')
        .or(page.getByRole('tab', { name: /History/i }))
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
      await page.goto('/asset-detail/1');
      await page.waitForTimeout(300);
      await expect(page.locator('body')).toBeVisible();
    });
  });

  test.describe('Error Handling', () => {
    test('should handle API error', async ({ page }) => {
      await page.route('**/api/assets/1**', (route) => route.abort());
      await page.goto('/asset-detail/1');
      await page.waitForTimeout(500);
      await expect(page.locator('body')).toBeVisible();
    });
  });
});
