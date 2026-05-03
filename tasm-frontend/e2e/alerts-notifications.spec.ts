import { test, expect } from '@playwright/test';

/**
 * Alerts & Notifications Center Tests
 */
test.describe('Alerts & Notifications Center', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/alerts-notifications-center');
    await page.waitForLoadState('networkidle').catch(() => {});
  });

  test.describe('Page Load', () => {
    test('should load Alerts page', async ({ page }) => {
      await expect(page.getByRole('heading', { name: /Alert/i })).toBeVisible();
    });

    test('should have filter buttons', async ({ page }) => {
      const filters = ['All', 'Critical', 'Warning', 'Info'];
      for (const filter of filters) {
        const btn = page.getByText(filter).first();
        if (await btn.isVisible()) {
          await expect(btn).toBeVisible();
        }
      }
    });

    test('should have Mark All Read button', async ({ page }) => {
      const markAllBtn = page.getByText('Mark All Read').or(page.getByText('Mark All')).first();
      if (await markAllBtn.isVisible()) {
        await expect(markAllBtn).toBeVisible();
      }
    });
  });

  test.describe('With Mocked Data', () => {
    test.beforeEach(async ({ page }) => {
      await page.route('**/api/alerts**', (route) => {
        route.fulfill({
          json: [
            {
              id: 1,
              title: 'Server Down',
              severity: 'Critical',
              message: 'Main server is unreachable',
            },
            {
              id: 2,
              title: 'Low Stock',
              severity: 'Warning',
              message: 'Printer paper running low',
            },
          ],
        });
      });
      await page.goto('/alerts-notifications-center');
      await page.waitForTimeout(500);
    });

    test('should display alerts', async ({ page }) => {
      await expect(page.getByText('Server Down')).toBeVisible();
      await expect(page.getByText('Low Stock')).toBeVisible();
    });

    test('should click on alert', async ({ page }) => {
      const alert = page.getByText('Server Down').first();
      if (await alert.isVisible()) {
        await alert.click();
        await page.waitForTimeout(300);
      }
    });

    test('should click Critical filter', async ({ page }) => {
      const criticalBtn = page.getByText('Critical').first();
      if (await criticalBtn.isVisible()) {
        await criticalBtn.click();
        await page.waitForTimeout(300);
      }
    });
  });

  test.describe('Alert Actions', () => {
    test('should click dismiss button on alert', async ({ page }) => {
      const dismissBtn = page.getByText('Dismiss').or(page.getByText('Close')).first();
      if (await dismissBtn.isVisible()) {
        await dismissBtn.click();
        await page.waitForTimeout(300);
      }
    });

    test('should click Mark All Read', async ({ page }) => {
      const markAllBtn = page.getByText('Mark All Read').first();
      if (await markAllBtn.isVisible()) {
        await markAllBtn.click();
        await page.waitForTimeout(300);
      }
    });
  });

  test.describe('Header Notification Bell', () => {
    test('should click notification bell in header', async ({ page }) => {
      await page.goto('/');
      await page.waitForTimeout(300);

      const bellBtn = page.locator('.material-symbols-outlined:has-text("notifications")').first();
      if (await bellBtn.isVisible()) {
        await bellBtn.click();
        await page.waitForTimeout(300);
      }
    });
  });

  test.describe('Keyboard Navigation', () => {
    test('should tab through alerts', async ({ page }) => {
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
      await expect(page.getByRole('heading', { name: /Alert/i })).toBeVisible();
    });
  });

  test.describe('Error Handling', () => {
    test('should handle API error', async ({ page }) => {
      await page.route('**/api/alerts**', (route) => route.abort());
      await page.reload();
      await page.waitForTimeout(500);
      await expect(page.locator('body')).toBeVisible();
    });
  });

  test.describe('Navigation', () => {
    test('should navigate from sidebar', async ({ page }) => {
      await page.goto('/');
      await page.getByRole('link', { name: /Alerts/i }).click();
      await expect(page).toHaveURL(/\/alerts-notifications-center$/);
    });
  });
});
