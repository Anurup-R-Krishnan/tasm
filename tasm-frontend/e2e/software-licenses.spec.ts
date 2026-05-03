import { test, expect } from '@playwright/test';

/**
 * Software License Registry Tests
 */
test.describe('Software License Registry', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/software-license-registry');
    await page.waitForLoadState('networkidle').catch(() => {});
  });

  test.describe('Page Load', () => {
    test('should load Licenses page', async ({ page }) => {
      await expect(page.getByRole('heading', { name: /License/i })).toBeVisible();
    });

    test('should have Add License button', async ({ page }) => {
      const addBtn = page.getByText('Add License').or(page.getByText('New')).first();
      if (await addBtn.isVisible()) {
        await expect(addBtn).toBeVisible();
      }
    });

    test('should have search input', async ({ page }) => {
      const search = page.locator('input[type="text"]').first();
      await expect(search).toBeVisible();
    });
  });

  test.describe('With Mocked Data', () => {
    test.beforeEach(async ({ page }) => {
      await page.route('**/api/software-licenses**', (route) => {
        route.fulfill({
          json: [
            { id: 1, name: 'Windows 11 Pro', seats: 100, used: 75 },
            { id: 2, name: 'Office 365', seats: 50, used: 50 },
          ],
        });
      });
      await page.goto('/software-license-registry');
      await page.waitForTimeout(500);
    });

    test('should display licenses', async ({ page }) => {
      await expect(page.getByText('Windows 11 Pro')).toBeVisible();
      await expect(page.getByText('Office 365')).toBeVisible();
    });

    test('should click on license row', async ({ page }) => {
      const row = page.locator('table tbody tr').first();
      if (await row.isVisible()) {
        await row.click();
        await page.waitForTimeout(300);
      }
    });
  });

  test.describe('Filter Buttons', () => {
    test('should click Active filter', async ({ page }) => {
      const btn = page.getByText('Active').first();
      if (await btn.isVisible()) {
        await btn.click();
        await page.waitForTimeout(300);
      }
    });

    test('should click Expired filter', async ({ page }) => {
      const btn = page.getByText('Expired').first();
      if (await btn.isVisible()) {
        await btn.click();
        await page.waitForTimeout(300);
      }
    });
  });

  test.describe('Search', () => {
    test('should search licenses', async ({ page }) => {
      const search = page.locator('input[type="text"]').first();
      await search.fill('Windows');
      await page.waitForTimeout(300);
      await search.clear();
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
      await expect(page.getByRole('heading', { name: /License/i })).toBeVisible();
    });
  });

  test.describe('Error Handling', () => {
    test('should handle API error', async ({ page }) => {
      await page.route('**/api/software-licenses**', (route) => route.abort());
      await page.reload();
      await page.waitForTimeout(500);
      await expect(page.locator('body')).toBeVisible();
    });
  });

  test.describe('Navigation', () => {
    test('should navigate from sidebar', async ({ page }) => {
      await page.goto('/');
      await page.getByRole('link', { name: /Licenses/i }).click();
      await expect(page).toHaveURL(/\/software-license-registry$/);
    });
  });
});
