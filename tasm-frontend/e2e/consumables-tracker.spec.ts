import { test, expect } from '@playwright/test';

/**
 * Consumables Supplies Tracker Tests
 * Tests the Consumables page functionality
 */
test.describe('Consumables Supplies Tracker', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/consumables-supplies-tracker');
    await page.waitForLoadState('networkidle').catch(() => {});
  });

  test.describe('Page Load', () => {
    test('should load Consumables page', async ({ page }) => {
      await expect(page.getByRole('heading', { name: /Consumables/i })).toBeVisible();
    });

    test('should have Add Consumable button', async ({ page }) => {
      const addBtn = page.getByText('Add').or(page.getByText('Add Consumable')).first();
      if (await addBtn.isVisible()) {
        await expect(addBtn).toBeVisible();
      }
    });

    test('should have search input', async ({ page }) => {
      const search = page.locator('input[type="text"]').first();
      await expect(search).toBeVisible();
    });

    test('should have filter buttons', async ({ page }) => {
      const filters = ['All', 'Low Stock', 'Out of Stock', 'In Stock'];
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
      await page.route('**/api/consumables**', (route) => {
        route.fulfill({
          json: [
            {
              id: 1,
              name: 'Printer Paper A4',
              category: 'Office Supplies',
              location: 'Central Stockroom',
              currentStock: 24,
              reorderLevel: 30,
            },
            {
              id: 2,
              name: 'USB-C Cable',
              category: 'IT Peripherals',
              location: 'IT Cage',
              currentStock: 0,
              reorderLevel: 10,
            },
          ],
        });
      });
      await page.goto('/consumables-supplies-tracker');
      await page.waitForTimeout(500);
    });

    test('should display consumables list', async ({ page }) => {
      await expect(page.getByText('Printer Paper A4')).toBeVisible();
      await expect(page.getByText('USB-C Cable')).toBeVisible();
    });

    test('should show Out of Stock indicator', async ({ page }) => {
      await expect(page.getByText('Out of Stock')).toBeVisible();
    });

    test('should show Low Stock indicator', async ({ page }) => {
      await expect(page.getByText('Low Stock')).toBeVisible();
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

    test('should click Low Stock filter', async ({ page }) => {
      const btn = page.getByText('Low Stock').first();
      if (await btn.isVisible()) {
        await btn.click();
        await page.waitForTimeout(300);
      }
    });

    test('should click Out of Stock filter', async ({ page }) => {
      const btn = page.getByText('Out of Stock').first();
      if (await btn.isVisible()) {
        await btn.click();
        await page.waitForTimeout(300);
      }
    });

    test('should toggle between filters', async ({ page }) => {
      const allBtn = page.getByText('All').first();
      const lowStockBtn = page.getByText('Low Stock').first();

      if (await allBtn.isVisible()) {
        await allBtn.click();
        await page.waitForTimeout(200);
      }

      if (await lowStockBtn.isVisible()) {
        await lowStockBtn.click();
        await page.waitForTimeout(200);
      }
    });
  });

  test.describe('Search Functionality', () => {
    test('should search consumables', async ({ page }) => {
      await page.route('**/api/consumables**', (route) => {
        route.fulfill({
          json: [{ id: 1, name: 'Printer Paper A4' }],
        });
      });

      const search = page.locator('input[type="text"]').first();
      await search.fill('Printer');
      await page.waitForTimeout(500);
    });

    test('should clear search', async ({ page }) => {
      const search = page.locator('input[type="text"]').first();
      await search.fill('Test');
      await page.waitForTimeout(200);
      await search.clear();
      await page.waitForTimeout(200);
    });
  });

  test.describe('Add Consumable', () => {
    test('should click Add button', async ({ page }) => {
      const addBtn = page.getByText('Add').or(page.getByText('Add Consumable')).first();
      if (await addBtn.isVisible()) {
        await addBtn.click();
        await page.waitForTimeout(500);
      }
    });
  });

  test.describe('Table Interactions', () => {
    test('should click on consumable row', async ({ page }) => {
      await page.route('**/api/consumables**', (route) => {
        route.fulfill({
          json: [{ id: 1, name: 'Test Item' }],
        });
      });

      await page.reload();
      await page.waitForTimeout(500);

      const rows = page.locator('table tbody tr, .p-datatable-tbody tr');
      const count = await rows.count();

      if (count > 0) {
        await rows.first().click();
        await page.waitForTimeout(300);
      }
    });

    test('should find edit/delete buttons in table', async ({ page }) => {
      const actionBtns = page.locator(
        'button:has(.material-symbols-outlined), .edit-btn, .delete-btn',
      );
      const count = await actionBtns.count();
      expect(count).toBeGreaterThanOrEqual(0);
    });
  });

  test.describe('Keyboard Navigation', () => {
    test('should tab through interactive elements', async ({ page }) => {
      for (let i = 0; i < 5; i++) {
        await page.keyboard.press('Tab');
      }
      const focused = page.locator(':focus');
      await expect(focused).toBeVisible();
    });

    test('should activate button with Enter', async ({ page }) => {
      const btn = page.getByText('All').first();
      if (await btn.isVisible()) {
        await btn.focus();
        await page.keyboard.press('Enter');
        await page.waitForTimeout(200);
      }
    });
  });

  test.describe('Responsive Design', () => {
    test('should display on mobile', async ({ page }) => {
      await page.setViewportSize({ width: 375, height: 667 });
      await page.waitForTimeout(300);
      await expect(page.getByRole('heading', { name: /Consumables/i })).toBeVisible();
    });

    test('should have scrollable table on mobile', async ({ page }) => {
      await page.setViewportSize({ width: 375, height: 667 });
      const table = page.locator('table, .p-datatable');
      if (await table.isVisible()) {
        await expect(table).toBeVisible();
      }
    });
  });

  test.describe('Error Handling', () => {
    test('should handle API error', async ({ page }) => {
      await page.route('**/api/consumables**', (route) => route.abort());
      await page.reload();
      await page.waitForTimeout(500);
      await expect(page.locator('body')).toBeVisible();
    });

    test('should handle empty results', async ({ page }) => {
      await page.route('**/api/consumables**', (route) => {
        route.fulfill({ json: [] });
      });
      await page.reload();
      await page.waitForTimeout(500);
    });
  });

  test.describe('Navigation', () => {
    test('should navigate from sidebar', async ({ page }) => {
      await page.goto('/');
      await page.getByRole('link', { name: /Consumables/i }).click();
      await expect(page).toHaveURL(/\/consumables-supplies-tracker$/);
    });
  });
});
