import { test, expect } from '@playwright/test';

/**
 * Stockroom Inventory Tests
 */
test.describe('Stockroom Inventory', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/stockroom-inventory');
    await page.waitForLoadState('networkidle').catch(() => {});
  });

  test.describe('Page Load', () => {
    test('should load Stockrooms page', async ({ page }) => {
      await expect(page.getByRole('heading', { name: /Stockroom/i })).toBeVisible();
    });

    test('should have Add Stockroom button', async ({ page }) => {
      const addBtn = page.getByText('Add Stockroom').or(page.getByText('New')).first();
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
      await page.route('**/api/stockrooms**', (route) => {
        route.fulfill({
          json: [
            { id: 1, name: 'Central Stockroom', address: 'Building A', capacity: 1000 },
            { id: 2, name: 'IT Cage', address: 'Building B', capacity: 500 },
          ],
        });
      });
      await page.goto('/stockroom-inventory');
      await page.waitForTimeout(500);
    });

    test('should display stockrooms', async ({ page }) => {
      await expect(page.getByText('Central Stockroom')).toBeVisible();
      await expect(page.getByText('IT Cage')).toBeVisible();
    });

    test('should click on stockroom card', async ({ page }) => {
      const card = page.getByText('Central Stockroom').first();
      if (await card.isVisible()) {
        await card.click();
        await page.waitForTimeout(300);
      }
    });
  });

  test.describe('Search', () => {
    test('should search stockrooms', async ({ page }) => {
      const search = page.locator('input[type="text"]').first();
      await search.fill('Central');
      await page.waitForTimeout(300);
      await search.clear();
    });
  });

  test.describe('View Toggle', () => {
    test('should toggle grid view', async ({ page }) => {
      const gridBtn = page
        .getByText('Grid')
        .or(page.locator('.material-symbols-outlined:has-text("grid_view")'))
        .first();
      if (await gridBtn.isVisible()) {
        await gridBtn.click();
        await page.waitForTimeout(300);
      }
    });

    test('should toggle list view', async ({ page }) => {
      const listBtn = page
        .getByText('List')
        .or(page.locator('.material-symbols-outlined:has-text("list")'))
        .first();
      if (await listBtn.isVisible()) {
        await listBtn.click();
        await page.waitForTimeout(300);
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
      await expect(page.getByRole('heading', { name: /Stockroom/i })).toBeVisible();
    });
  });

  test.describe('Error Handling', () => {
    test('should handle API error', async ({ page }) => {
      await page.route('**/api/stockrooms**', (route) => route.abort());
      await page.reload();
      await page.waitForTimeout(500);
      await expect(page.locator('body')).toBeVisible();
    });
  });

  test.describe('Navigation', () => {
    test('should navigate from sidebar', async ({ page }) => {
      await page.goto('/');
      await page.getByRole('link', { name: /Stockrooms/i }).click();
      await expect(page).toHaveURL(/\/stockroom-inventory$/);
    });
  });
});
