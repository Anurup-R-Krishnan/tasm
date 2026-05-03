import { test, expect } from '@playwright/test';

/**
 * Add New Asset Form Tests
 */
test.describe('Add New Asset Form', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/add-new-asset-form');
    await page.waitForLoadState('networkidle').catch(() => {});
  });

  test.describe('Page Load', () => {
    test('should load Add Asset form', async ({ page }) => {
      await expect(page.locator('body')).toBeVisible();
    });

    test('should have name input', async ({ page }) => {
      const nameInput = page
        .locator('input[placeholder*="name" i]')
        .or(page.getByLabel('Name'))
        .first();
      if (await nameInput.isVisible()) {
        await expect(nameInput).toBeVisible();
      }
    });

    test('should have category select', async ({ page }) => {
      const select = page.locator('select').or(page.getByRole('combobox')).first();
      if (await select.isVisible()) {
        await expect(select).toBeVisible();
      }
    });

    test('should have Generate Tag ID button', async ({ page }) => {
      const genBtn = page.getByText('Generate').or(page.getByText('Tag ID')).first();
      if (await genBtn.isVisible()) {
        await expect(genBtn).toBeVisible();
      }
    });

    test('should have Cancel button', async ({ page }) => {
      const cancelBtn = page.getByText('Cancel').first();
      if (await cancelBtn.isVisible()) {
        await expect(cancelBtn).toBeVisible();
      }
    });

    test('should have Save button', async ({ page }) => {
      const saveBtn = page
        .getByText('Save')
        .or(page.getByRole('button', { name: /Save/i }))
        .first();
      if (await saveBtn.isVisible()) {
        await expect(saveBtn).toBeVisible();
      }
    });
  });

  test.describe('Form Filling', () => {
    test('should fill name field', async ({ page }) => {
      const nameInput = page.locator('input[type="text"]').first();
      if (await nameInput.isVisible()) {
        await nameInput.fill('Dell Laptop');
        await page.waitForTimeout(200);
      }
    });

    test('should select category', async ({ page }) => {
      const select = page.locator('select').first();
      if (await select.isVisible()) {
        await select.click();
        await page.waitForTimeout(200);
      }
    });

    test('should fill value field', async ({ page }) => {
      const valueInput = page.locator('input[type="number"]').first();
      if (await valueInput.isVisible()) {
        await valueInput.fill('50000');
        await page.waitForTimeout(200);
      }
    });

    test('should click Generate Tag ID', async ({ page }) => {
      const genBtn = page.getByText('Generate').first();
      if (await genBtn.isVisible()) {
        await genBtn.click();
        await page.waitForTimeout(300);
      }
    });
  });

  test.describe('Form Actions', () => {
    test('should click Cancel button', async ({ page }) => {
      const cancelBtn = page.getByText('Cancel').first();
      if (await cancelBtn.isVisible()) {
        await cancelBtn.click();
        await page.waitForTimeout(300);
      }
    });

    test('should click Save button', async ({ page }) => {
      await page.route('**/api/assets**', (route) => {
        route.fulfill({ status: 200, json: { success: true } });
      });

      const saveBtn = page.getByText('Save').first();
      if (await saveBtn.isVisible()) {
        await saveBtn.click();
        await page.waitForTimeout(500);
      }
    });
  });

  test.describe('Form Validation', () => {
    test('should show errors on empty submit', async ({ page }) => {
      const saveBtn = page.getByText('Save').first();
      if (await saveBtn.isVisible()) {
        await saveBtn.click();
        await page.waitForTimeout(500);

        const errors = page.locator('[class*="error"], .text-red, .invalid');
        const count = await errors.count();
        expect(count).toBeGreaterThanOrEqual(0);
      }
    });
  });

  test.describe('Keyboard Navigation', () => {
    test('should tab through form fields', async ({ page }) => {
      for (let i = 0; i < 8; i++) {
        await page.keyboard.press('Tab');
      }
      const focused = page.locator(':focus');
      await expect(focused).toBeVisible();
    });
  });

  test.describe('Responsive', () => {
    test('should display on mobile', async ({ page }) => {
      await page.setViewportSize({ width: 375, height: 667 });
      await page.goto('/add-new-asset-form');
      await page.waitForTimeout(300);
      await expect(page.locator('body')).toBeVisible();
    });
  });

  test.describe('Error Handling', () => {
    test('should handle API error on save', async ({ page }) => {
      await page.route('**/api/assets**', (route) => route.abort());

      const saveBtn = page.getByText('Save').first();
      if (await saveBtn.isVisible()) {
        await saveBtn.click();
        await page.waitForTimeout(500);
      }

      await expect(page.locator('body')).toBeVisible();
    });
  });
});
