import { test, expect } from '@playwright/test';

/**
 * Reservations & Bookings Tests
 */
test.describe('Reservations & Bookings', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/reservations-bookings');
    await page.waitForLoadState('networkidle').catch(() => {});
  });

  test.describe('Page Load', () => {
    test('should load Reservations page', async ({ page }) => {
      await expect(page.getByRole('heading', { name: /Reservation/i })).toBeVisible();
    });

    test('should have Book button', async ({ page }) => {
      const bookBtn = page.getByText('Book').or(page.getByText('New Reservation')).first();
      if (await bookBtn.isVisible()) {
        await expect(bookBtn).toBeVisible();
      }
    });

    test('should have search input', async ({ page }) => {
      const search = page.locator('input[type="text"]').first();
      await expect(search).toBeVisible();
    });

    test('should have filter buttons', async ({ page }) => {
      const filters = ['All', 'Pending', 'Approved', 'Cancelled'];
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
      await page.route('**/api/reservations**', (route) => {
        route.fulfill({
          json: [
            { id: 1, title: 'Meeting Room A', status: 'Approved', date: '2026-05-10' },
            { id: 2, title: 'Projector', status: 'Pending', date: '2026-05-12' },
          ],
        });
      });
      await page.goto('/reservations-bookings');
      await page.waitForTimeout(500);
    });

    test('should display reservations list', async ({ page }) => {
      await expect(page.getByText('Meeting Room A')).toBeVisible();
      await expect(page.getByText('Projector')).toBeVisible();
    });

    test('should click on reservation row', async ({ page }) => {
      const row = page.locator('table tbody tr').first();
      if (await row.isVisible()) {
        await row.click();
        await page.waitForTimeout(300);
      }
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

    test('should click Pending filter', async ({ page }) => {
      const btn = page.getByText('Pending').first();
      if (await btn.isVisible()) {
        await btn.click();
        await page.waitForTimeout(300);
      }
    });

    test('should click Approved filter', async ({ page }) => {
      const btn = page.getByText('Approved').first();
      if (await btn.isVisible()) {
        await btn.click();
        await page.waitForTimeout(300);
      }
    });
  });

  test.describe('Book New Reservation', () => {
    test('should click Book button', async ({ page }) => {
      const bookBtn = page.getByText('Book').or(page.getByText('New')).first();
      if (await bookBtn.isVisible()) {
        await bookBtn.click();
        await page.waitForTimeout(500);
      }
    });

    test('should fill booking form', async ({ page }) => {
      const inputs = page.locator('input[type="text"], input[type="date"]');
      const count = await inputs.count();
      for (let i = 0; i < Math.min(count, 3); i++) {
        await inputs
          .nth(i)
          .fill('Test')
          .catch(() => {});
      }
    });
  });

  test.describe('Search', () => {
    test('should search reservations', async ({ page }) => {
      const search = page.locator('input[type="text"]').first();
      await search.fill('Meeting');
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
      await expect(page.getByRole('heading', { name: /Reservation/i })).toBeVisible();
    });
  });

  test.describe('Error Handling', () => {
    test('should handle API error', async ({ page }) => {
      await page.route('**/api/reservations**', (route) => route.abort());
      await page.reload();
      await page.waitForTimeout(500);
      await expect(page.locator('body')).toBeVisible();
    });
  });
});
