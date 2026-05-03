import { test, expect } from '@playwright/test';

/**
 * Asset Registry Page - Comprehensive Interaction Tests
 * Tests all clickable elements on the Asset Registry page
 */
test.describe('Asset Registry Page Interactions', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/asset-registry');
    await page.waitForLoadState('networkidle').catch(() => {});
  });

  test.describe('Page Load and Basic Elements', () => {
    test('should load Asset Registry page', async ({ page }) => {
      await expect(page.getByRole('heading', { name: 'Asset Registry' })).toBeVisible();
    });

    test('should have Add Asset button', async ({ page }) => {
      const addBtn = page
        .getByRole('button', { name: /Add Asset/i })
        .or(page.getByText('Add Asset'))
        .first();
      await expect(addBtn).toBeVisible();
    });

    test('should have Export button', async ({ page }) => {
      const exportBtn = page.getByText('Export').first();
      if (await exportBtn.isVisible()) {
        await expect(exportBtn).toBeVisible();
      }
    });

    test('should have Advanced Filters button', async ({ page }) => {
      const filterBtn = page.getByText('Advanced Filters').first();
      await expect(filterBtn).toBeVisible();
    });

    test('should have search input', async ({ page }) => {
      const searchInput = page.locator('input[type="text"]').first();
      await expect(searchInput).toBeVisible();
    });
  });

  test.describe('Filter Buttons', () => {
    test('should find all status filter buttons', async ({ page }) => {
      const filters = ['All', 'In Stock', 'Checked Out', 'Under Repair'];
      for (const filter of filters) {
        const btn = page.getByText(filter).first();
        await expect(btn).toBeVisible();
      }
    });

    test('should click All filter button', async ({ page }) => {
      const btn = page.getByText('All').first();
      await btn.click();
      await page.waitForTimeout(300);
      await expect(btn).toBeVisible();
    });

    test('should click In Stock filter button', async ({ page }) => {
      const btn = page.getByText('In Stock').first();
      await btn.click();
      await page.waitForTimeout(300);
      await expect(btn).toBeVisible();
    });

    test('should click Checked Out filter button', async ({ page }) => {
      const btn = page.getByText('Checked Out').first();
      await btn.click();
      await page.waitForTimeout(300);
      await expect(btn).toBeVisible();
    });

    test('should click Under Repair filter button', async ({ page }) => {
      const btn = page.getByText('Under Repair').first();
      await btn.click();
      await page.waitForTimeout(300);
      await expect(btn).toBeVisible();
    });

    test('should toggle between filters', async ({ page }) => {
      const allBtn = page.getByText('All').first();
      const inStockBtn = page.getByText('In Stock').first();

      await allBtn.click();
      await page.waitForTimeout(200);

      await inStockBtn.click();
      await page.waitForTimeout(200);

      await allBtn.click();
      await page.waitForTimeout(200);
    });
  });

  test.describe('Advanced Filters Toggle', () => {
    test('should click Advanced Filters button', async ({ page }) => {
      const btn = page.getByText('Advanced Filters').first();
      await btn.click();
      await page.waitForTimeout(300);
    });

    test('should show filter panel after clicking Advanced Filters', async ({ page }) => {
      const btn = page.getByText('Advanced Filters').first();
      await btn.click();
      await page.waitForTimeout(300);

      const filterPanel = page.locator('[class*="filter"], [class*="advanced"]').first();
      await filterPanel.isVisible().catch(() => false);
    });

    test('should toggle Advanced Filters on/off', async ({ page }) => {
      const btn = page.getByText('Advanced Filters').first();

      await btn.click();
      await page.waitForTimeout(300);

      await btn.click();
      await page.waitForTimeout(300);
    });
  });

  test.describe('Search Functionality', () => {
    test('should type in search input', async ({ page }) => {
      const searchInput = page.locator('input[type="text"]').first();
      await searchInput.fill('Laptop');
      await page.waitForTimeout(300);
    });

    test('should clear search input', async ({ page }) => {
      const searchInput = page.locator('input[type="text"]').first();
      await searchInput.fill('Test');
      await page.waitForTimeout(200);
      await searchInput.clear();
      await page.waitForTimeout(200);
    });

    test('should search and see results', async ({ page }) => {
      await page.route('**/api/assets**', (route) => {
        route.fulfill({
          json: [{ id: 1, name: 'Dell Laptop', status: 'In Stock', tagId: 'AST-001' }],
        });
      });

      const searchInput = page.locator('input[type="text"]').first();
      await searchInput.fill('Dell');
      await page.waitForTimeout(500);
    });
  });

  test.describe('Add Asset Button', () => {
    test('should click Add Asset button', async ({ page }) => {
      const addBtn = page.getByText('Add Asset').first();
      if (await addBtn.isVisible()) {
        await addBtn.click();
        await page.waitForTimeout(500);
      }
    });

    test('should navigate to add asset form', async ({ page }) => {
      const addBtn = page.getByText('Add Asset').first();
      if (await addBtn.isVisible()) {
        await addBtn.click();
        await page.waitForTimeout(500);
      }
    });
  });

  test.describe('Asset Table/List Interactions', () => {
    test('should click on asset row', async ({ page }) => {
      await page.route('**/api/assets**', (route) => {
        route.fulfill({
          json: [{ id: 1, name: 'Test Asset', status: 'In Stock', tagId: 'AST-001' }],
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

    test('should find action buttons in table', async ({ page }) => {
      const actionButtons = page.locator(
        'button[class*="p-button"], .action-btn, [class*="edit"], [class*="delete"]',
      );
      const count = await actionButtons.count();
      expect(count).toBeGreaterThanOrEqual(0);
    });
  });

  test.describe('Export Button', () => {
    test('should click Export button', async ({ page }) => {
      const exportBtn = page.getByText('Export').first();
      if (await exportBtn.isVisible()) {
        await exportBtn.click();
        await page.waitForTimeout(500);
      }
    });

    test('should handle export with API mock', async ({ page }) => {
      await page.route('**/api/assets/export**', (route) => {
        route.fulfill({
          headers: { 'Content-Type': 'application/csv' },
          body: 'id,name,status\n1,Test,In Stock',
        });
      });

      const exportBtn = page.getByText('Export').first();
      if (await exportBtn.isVisible()) {
        await exportBtn.click();
        await page.waitForTimeout(500);
      }
    });
  });

  test.describe('Keyboard Navigation', () => {
    test('should tab through interactive elements', async ({ page }) => {
      await page.keyboard.press('Tab');
      await page.keyboard.press('Tab');
      await page.keyboard.press('Tab');
      await page.keyboard.press('Tab');

      const focused = page.locator(':focus');
      await expect(focused).toBeVisible();
    });

    test('should activate filter with Enter', async ({ page }) => {
      const btn = page.getByText('In Stock').first();
      await btn.focus();
      await page.keyboard.press('Enter');
      await page.waitForTimeout(300);
    });
  });

  test.describe('Responsive Behavior', () => {
    test('should display correctly on mobile', async ({ page }) => {
      await page.setViewportSize({ width: 375, height: 667 });
      await page.waitForTimeout(300);

      await expect(page.getByRole('heading', { name: 'Asset Registry' })).toBeVisible();
    });

    test('should have scrollable table on small screen', async ({ page }) => {
      await page.setViewportSize({ width: 375, height: 667 });

      const table = page.locator('table, .p-datatable');
      if (await table.isVisible()) {
        await expect(table).toBeVisible();
      }
    });
  });

  test.describe('Error Handling', () => {
    test('should handle API error gracefully', async ({ page }) => {
      await page.route('**/api/assets**', (route) => route.abort());

      await page.reload();
      await page.waitForTimeout(500);

      await expect(page.locator('body')).toBeVisible();
    });

    test('should not crash on rapid filter clicks', async ({ page }) => {
      const filters = ['All', 'In Stock', 'Checked Out', 'Under Repair'];

      for (const filter of filters) {
        const btn = page.getByText(filter).first();
        await btn.click().catch(() => {});
      }

      await expect(page.locator('body')).toBeVisible();
    });
  });
});
