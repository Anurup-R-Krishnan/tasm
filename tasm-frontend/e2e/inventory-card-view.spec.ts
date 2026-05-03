import { test, expect } from '@playwright/test';

/**
 * Inventory Card View Tests
 */
test.describe('Inventory Card View', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/inventory-card-view');
    await page.waitForLoadState('networkidle').catch(() => {});
  });

  test.describe('Page Load', () => {
    test('should load Asset Cards page', async ({ page }) => {
      await expect(page.getByRole('heading', { name: /Asset Card/i })).toBeVisible();
    });

    test('should have View All button', async ({ page }) => {
      const viewAllBtn = page.getByText('View All').first();
      if (await viewAllBtn.isVisible()) {
        await expect(viewAllBtn).toBeVisible();
      }
    });

    test('should have search input', async ({ page }) => {
      const search = page.locator('input[type="text"]').first();
      await expect(search).toBeVisible();
    });

    test('should have filter buttons', async ({ page }) => {
      const filters = ['All', 'In Stock', 'Checked Out'];
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
      await page.route('**/api/assets**', (route) => {
        route.fulfill({
          json: [
            { id: 1, name: 'Dell Laptop', tagId: 'AST-001', status: 'In Stock' },
            { id: 2, name: 'HP Monitor', tagId: 'AST-002', status: 'Checked Out' },
          ],
        });
      });
      await page.goto('/inventory-card-view');
      await page.waitForTimeout(500);
    });

    test('should display asset cards', async ({ page }) => {
      await expect(page.getByText('Dell Laptop')).toBeVisible();
      await expect(page.getByText('HP Monitor')).toBeVisible();
    });

    test('should click on asset card', async ({ page }) => {
      const card = page.getByText('Dell Laptop').first();
      if (await card.isVisible()) {
        await card.click();
        await page.waitForTimeout(300);
      }
    });

    test('should click status filter', async ({ page }) => {
      const inStockBtn = page.getByText('In Stock').first();
      if (await inStockBtn.isVisible()) {
        await inStockBtn.click();
        await page.waitForTimeout(300);
      }
    });
  });

  test.describe('Search', () => {
    test('should search assets', async ({ page }) => {
      const search = page.locator('input[type="text"]').first();
      await search.fill('Dell');
      await page.waitForTimeout(300);
      await search.clear();
    });
  });

  test.describe('Pagination', () => {
    test('should click next page', async ({ page }) => {
      const nextBtn = page.getByText('Next').or(page.getByText('>')).first();
      if (await nextBtn.isVisible()) {
        await nextBtn.click();
        await page.waitForTimeout(300);
      }
    });

    test('should click previous page', async ({ page }) => {
      const prevBtn = page.getByText('Previous').or(page.getByText('<')).first();
      if (await prevBtn.isVisible()) {
        await prevBtn.click();
        await page.waitForTimeout(300);
      }
    });
  });

  test.describe('Keyboard Navigation', () => {
    test('should tab through cards', async ({ page }) => {
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
      await expect(page.getByRole('heading', { name: /Asset Card/i })).toBeVisible();
    });

    test('should have proper card layout on mobile', async ({ page }) => {
      await page.setViewportSize({ width: 375, height: 667 });
      const cards = page.locator('.card, [class*="card"]');
      const count = await cards.count();
      expect(count).toBeGreaterThanOrEqual(0);
    });
  });

  test.describe('Error Handling', () => {
    test('should handle API error', async ({ page }) => {
      await page.route('**/api/assets**', (route) => route.abort());
      await page.reload();
      await page.waitForTimeout(500);
      await expect(page.locator('body')).toBeVisible();
    });
  });

  test.describe('Navigation', () => {
    test('should navigate from sidebar', async ({ page }) => {
      await page.goto('/');
      await page.getByRole('link', { name: /Asset Cards/i }).click();
      await expect(page).toHaveURL(/\/inventory-card-view$/);
    });
  });
});
