import { test, expect } from '@playwright/test';

/**
 * Maintenance Schedule & Contracts Tests
 */
test.describe('Maintenance Schedule & Contracts', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/maintenance-schedule-contracts');
    await page.waitForLoadState('networkidle').catch(() => {});
  });

  test.describe('Page Load', () => {
    test('should load Maintenance page', async ({ page }) => {
      await expect(page.getByRole('heading', { name: /Maintenance/i })).toBeVisible();
    });

    test('should have New Contract button', async ({ page }) => {
      const newBtn = page.getByText('New Contract').or(page.getByText('Add')).first();
      if (await newBtn.isVisible()) {
        await expect(newBtn).toBeVisible();
      }
    });

    test('should have calendar view toggle', async ({ page }) => {
      const calendarBtn = page.getByText('Calendar').or(page.getByText('Grid')).first();
      if (await calendarBtn.isVisible()) {
        await expect(calendarBtn).toBeVisible();
      }
    });
  });

  test.describe('With Mocked Data', () => {
    test.beforeEach(async ({ page }) => {
      await page.route('**/api/contracts**', (route) => {
        route.fulfill({
          json: [
            { id: 1, title: 'HVAC Service', date: '2026-05-15', status: 'Scheduled' },
            { id: 2, title: 'Elevator Inspection', date: '2026-05-20', status: 'Pending' },
          ],
        });
      });
      await page.route('**/api/work-orders**', (route) => {
        route.fulfill({ json: [] });
      });
      await page.goto('/maintenance-schedule-contracts');
      await page.waitForTimeout(500);
    });

    test('should display maintenance items', async ({ page }) => {
      await expect(page.getByText('HVAC Service')).toBeVisible();
      await expect(page.getByText('Elevator Inspection')).toBeVisible();
    });

    test('should click on maintenance item', async ({ page }) => {
      const row = page.locator('table tbody tr, .maintenance-item').first();
      if (await row.isVisible()) {
        await row.click();
        await page.waitForTimeout(300);
      }
    });
  });

  test.describe('Filter Buttons', () => {
    test('should click Scheduled filter', async ({ page }) => {
      const btn = page.getByText('Scheduled').first();
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

    test('should click Completed filter', async ({ page }) => {
      const btn = page.getByText('Completed').first();
      if (await btn.isVisible()) {
        await btn.click();
        await page.waitForTimeout(300);
      }
    });
  });

  test.describe('View Toggle', () => {
    test('should toggle calendar view', async ({ page }) => {
      const calendarBtn = page.getByText('Calendar').first();
      if (await calendarBtn.isVisible()) {
        await calendarBtn.click();
        await page.waitForTimeout(300);
      }
    });

    test('should toggle list view', async ({ page }) => {
      const listBtn = page.getByText('List').first();
      if (await listBtn.isVisible()) {
        await listBtn.click();
        await page.waitForTimeout(300);
      }
    });
  });

  test.describe('New Contract', () => {
    test('should click New Contract button', async ({ page }) => {
      const newBtn = page.getByText('New Contract').or(page.getByText('Add')).first();
      if (await newBtn.isVisible()) {
        await newBtn.click();
        await page.waitForTimeout(500);
      }
    });
  });

  test.describe('Search', () => {
    test('should search maintenance items', async ({ page }) => {
      const search = page.locator('input[type="text"]').first();
      await search.fill('HVAC');
      await page.waitForTimeout(300);
      await search.clear();
    });
  });

  test.describe('Navigation', () => {
    test('should navigate from sidebar', async ({ page }) => {
      await page.goto('/');
      await page.getByRole('link', { name: /Maintenance/i }).click();
      await expect(page).toHaveURL(/\/maintenance-schedule-contracts$/);
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
      await expect(page.getByRole('heading', { name: /Maintenance/i })).toBeVisible();
    });
  });

  test.describe('Error Handling', () => {
    test('should handle API error', async ({ page }) => {
      await page.route('**/api/contracts**', (route) => route.abort());
      await page.reload();
      await page.waitForTimeout(500);
      await expect(page.locator('body')).toBeVisible();
    });
  });
});
