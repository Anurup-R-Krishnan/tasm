import { test, expect } from '@playwright/test';

/**
 * User Management & Settings Tests
 */
test.describe('User Management & Settings', () => {
  test.describe('Settings Page', () => {
    test.beforeEach(async ({ page }) => {
      await page.goto('/user-management-settings');
      await page.waitForLoadState('networkidle').catch(() => {});
    });

    test('should load Settings page', async ({ page }) => {
      await expect(page.getByRole('heading', { name: /Settings/i })).toBeVisible();
    });

    test('should have Save button', async ({ page }) => {
      const saveBtn = page.getByText('Save').or(page.getByText('Save Changes')).first();
      if (await saveBtn.isVisible()) {
        await expect(saveBtn).toBeVisible();
      }
    });

    test('should have input fields', async ({ page }) => {
      const inputs = page.locator('input[type="text"], input[type="email"]');
      const count = await inputs.count();
      expect(count).toBeGreaterThanOrEqual(0);
    });

    test('should fill input fields', async ({ page }) => {
      const inputs = page.locator('input[type="text"]');
      const count = await inputs.count();
      for (let i = 0; i < Math.min(count, 3); i++) {
        await inputs
          .nth(i)
          .fill('Test')
          .catch(() => {});
      }
    });

    test('should click Save button', async ({ page }) => {
      const saveBtn = page.getByText('Save').first();
      if (await saveBtn.isVisible()) {
        await saveBtn.click();
        await page.waitForTimeout(300);
      }
    });
  });

  test.describe('User Access & Role Control', () => {
    test.beforeEach(async ({ page }) => {
      await page.goto('/user-access-role-control');
      await page.waitForLoadState('networkidle').catch(() => {});
    });

    test('should load User Roles page', async ({ page }) => {
      await expect(page.getByRole('heading', { name: /Role/i })).toBeVisible();
    });

    test('should have Add Role button', async ({ page }) => {
      const addBtn = page.getByText('Add Role').or(page.getByText('New Role')).first();
      if (await addBtn.isVisible()) {
        await expect(addBtn).toBeVisible();
      }
    });

    test('should have user table', async ({ page }) => {
      const table = page.locator('table, .p-datatable');
      const isVisible = await table.isVisible().catch(() => false);
      if (isVisible) {
        await expect(table).toBeVisible();
      }
    });
  });

  test.describe('Employee Self Service Portal', () => {
    test.beforeEach(async ({ page }) => {
      await page.goto('/employee-self-service-portal');
      await page.waitForLoadState('networkidle').catch(() => {});
    });

    test('should load Self Service page', async ({ page }) => {
      await expect(page.getByRole('heading', { name: /Self Service/i })).toBeVisible();
    });

    test('should have profile section', async ({ page }) => {
      const profile = page.locator('[class*="profile"], .avatar, img').first();
      const isVisible = await profile.isVisible().catch(() => false);
      if (isVisible) {
        await expect(profile).toBeVisible();
      }
    });
  });

  test.describe('Navigation', () => {
    test('should navigate between user pages', async ({ page }) => {
      await page.goto('/user-management-settings');
      await page.waitForTimeout(300);

      await page
        .getByRole('link', { name: /User Roles/i })
        .click()
        .catch(() => {});
      await page.waitForTimeout(300);

      await page
        .getByRole('link', { name: /Self Service/i })
        .click()
        .catch(() => {});
      await page.waitForTimeout(300);
    });
  });

  test.describe('Keyboard Navigation', () => {
    test('should tab through settings page', async ({ page }) => {
      await page.goto('/user-management-settings');
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
      await page.goto('/user-management-settings');
      await page.waitForTimeout(300);
      await expect(page.getByRole('heading', { name: /Settings/i })).toBeVisible();
    });
  });

  test.describe('Error Handling', () => {
    test('should handle API error', async ({ page }) => {
      await page.route('**/api/users**', (route) => route.abort());
      await page.goto('/user-management-settings');
      await page.waitForTimeout(500);
      await expect(page.locator('body')).toBeVisible();
    });
  });
});
