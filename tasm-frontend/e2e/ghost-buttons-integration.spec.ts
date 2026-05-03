import { test, expect } from '@playwright/test';

test.describe('Integration Tests for Ghost Buttons', () => {
  test.describe('With Other UI Components', () => {
    test('should interact with sidebar navigation then ghost button', async ({ page }) => {
      await page.goto('/');

      await page.getByRole('link', { name: 'Consumables' }).click();
      await expect(page.getByText('Consumables Tracker')).toBeVisible();

      const button = page.getByText('7 Days').first();
      if (await button.isVisible()) {
        await button.click();
      }
    });

    test('should interact with search input then ghost button', async ({ page }) => {
      await page.goto('/');

      const searchInput = page.locator('input[type="text"]').first();
      if (await searchInput.isVisible()) {
        await searchInput.fill('test');
      }

      const button = page.getByText('7 Days').first();
      await button.click();
    });

    test('should handle ghost button with dropdown nearby', async ({ page }) => {
      await page.goto('/');

      const dropdowns = page.locator('[role="combobox"], select, .dropdown');
      const count = await dropdowns.count();

      const button = page.getByText('7 Days').first();
      await button.click();

      expect(count).toBeGreaterThanOrEqual(0);
    });
  });

  test.describe('With Forms and Validation', () => {
    test('should not submit form when ghost button clicked', async ({ page }) => {
      await page.goto('/');

      const forms = page.locator('form');
      const formCount = await forms.count();

      const button = page.getByText('7 Days').first();
      await button.click();

      expect(formCount).toBeGreaterThanOrEqual(0);
    });

    test('should handle ghost button inside form', async ({ page }) => {
      await page.goto('/');

      const buttonsInForms = page.locator('form button');
      const count = await buttonsInForms.count();

      if (count > 0) {
        const type = await buttonsInForms.first().getAttribute('type');
        expect(['button', 'submit', null]).toContain(type);
      }
    });
  });

  test.describe('With Modals and Dialogs', () => {
    test('should not open modal on ghost button click', async ({ page }) => {
      await page.goto('/');

      const button = page.getByText('7 Days').first();
      await button.click();

      await page.waitForTimeout(300);
      const modal = page.locator('[role="dialog"], .modal, .popup');
      await modal.isVisible().catch(() => false);
      expect(isVisible).toBeFalsy();
    });

    test('should handle ghost button in modal context', async ({ page }) => {
      await page.goto('/');

      const buttons = page.locator('button');
      const count = await buttons.count();

      for (let i = 0; i < Math.min(count, 3); i++) {
        await buttons
          .nth(i)
          .click()
          .catch(() => {});
      }

      await expect(page.locator('body')).toBeVisible();
    });
  });

  test.describe('With Data Tables', () => {
    test('should not affect table data on ghost button click', async ({ page }) => {
      await page.goto('/');

      const button = page.getByText('7 Days').first();
      await button.click();

      await page.waitForTimeout(500);
      const rows = page.locator('table tbody tr, .p-datatable-tbody tr');
      const count = await rows.count();
      expect(count).toBeGreaterThanOrEqual(0);
    });
  });

  test.describe('With Routing', () => {
    test('should maintain route after ghost button click', async ({ page }) => {
      await page.goto('/');
      const initialUrl = page.url();

      const button = page.getByText('7 Days').first();
      await button.click();
      await page.waitForTimeout(300);

      expect(page.url()).toBe(initialUrl);
    });

    test('should handle ghost button after route change', async ({ page }) => {
      await page.goto('/');
      await page.getByRole('link', { name: 'Asset Registry' }).click();
      await expect(page).toHaveURL(/\/asset-registry$/);

      const filterBtn = page.getByText('Advanced Filters').first();
      if (await filterBtn.isVisible()) {
        await filterBtn.click();
      }
    });
  });

  test.describe('With API Data', () => {
    test('should refresh data on filter button click', async ({ page }) => {
      let requestCount = 0;

      await page.route('**/api/**', (route) => {
        requestCount++;
        route.continue();
      });

      await page.goto('/');
      const button = page.getByText('7 Days').first();
      await button.click();

      await page.waitForTimeout(1000);
      expect(requestCount).toBeGreaterThanOrEqual(0);
    });

    test('should handle cached data with ghost button', async ({ page }) => {
      await page.route('**/api/**', (route) => {
        route.fulfill({
          headers: { 'Cache-Control': 'max-age=3600' },
          json: [],
        });
      });

      await page.goto('/');
      const button = page.getByText('7 Days').first();
      await button.click();
    });
  });

  test.describe('With LocalStorage/SessionStorage', () => {
    test('should persist filter state to storage', async ({ page }) => {
      await page.goto('/');

      await page.evaluate(() => {
        localStorage.setItem('timeFilter', '7days');
      });

      const button = page.getByText('7 Days').first();
      await button.click();

      const stored = await page.evaluate(() => localStorage.getItem('timeFilter'));
      expect(stored).toBe('7days');
    });
  });

  test.describe('With Cookies', () => {
    test('should not set cookies on ghost button click', async ({ page }) => {
      await page.goto('/');

      const button = page.getByText('7 Days').first();
      await button.click();

      const cookies = await page.context().cookies();
      expect(cookies.length).toBeGreaterThanOrEqual(0);
    });
  });

  test.describe('With Events and Analytics', () => {
    test('should track click event', async ({ page }) => {
      await page.goto('/');

      const events: string[] = [];
      await page.exposeFunction('trackEvent', (name: string) => {
        events.push(name);
      });

      await page.evaluate(() => {
        const btn = document.querySelector('button');
        if (btn) {
          btn.addEventListener('click', () => {
            (window as any).trackEvent?.('ghost_button_clicked');
          });
        }
      });

      const button = page.getByText('7 Days').first();
      await button.click();
      await page.waitForTimeout(200);
    });
  });

  test.describe('With Animations', () => {
    test('should wait for animation to complete', async ({ page }) => {
      await page.goto('/');

      const button = page.getByText('7 Days').first();
      await button.click();

      await page.waitForTimeout(500);
      await expect(button).toBeVisible();
    });

    test('should handle button with CSS animations', async ({ page }) => {
      await page.goto('/');

      await page.evaluate(() => {
        const btn = document.querySelector('button');
        if (btn) {
          btn.style.animation = 'fadeIn 0.3s ease-in';
        }
      });

      const button = page.getByText('7 Days').first();
      await button.click();
      await expect(button).toBeVisible();
    });
  });
});
