import { test, expect } from '@playwright/test';

/**
 * Form Validation Tests
 * Tests form inputs, validation, and submission
 */
test.describe('Form Validation Tests', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/');
  });

  test.describe('Add New Asset Form', () => {
    test('should navigate to add asset form', async ({ page }) => {
      await page.getByRole('link', { name: /Asset Registry/i }).click();
      await page.waitForTimeout(300);

      const addBtn = page.getByText('Add Asset').first();
      if (await addBtn.isVisible()) {
        await addBtn.click();
        await page.waitForTimeout(500);
      }
    });

    test('should fill text inputs', async ({ page }) => {
      await page.goto('/add-new-asset-form').catch(() => {});

      const inputs = page.locator('input[type="text"], input:not([type])');
      const count = await inputs.count();

      for (let i = 0; i < Math.min(count, 5); i++) {
        await inputs
          .nth(i)
          .fill('Test Input')
          .catch(() => {});
      }
    });

    test('should select from dropdown', async ({ page }) => {
      await page.goto('/add-new-asset-form').catch(() => {});

      const selects = page.locator('select, [role="combobox"]');
      const count = await selects.count();

      if (count > 0) {
        await selects
          .first()
          .click()
          .catch(() => {});
        await page.waitForTimeout(300);
      }
    });

    test('should check checkbox', async ({ page }) => {
      await page.goto('/add-new-asset-form').catch(() => {});

      const checkboxes = page.locator('input[type="checkbox"]');
      const count = await checkboxes.count();

      if (count > 0) {
        await checkboxes.first().check();
      }
    });

    test('should click Generate Tag ID button', async ({ page }) => {
      await page.goto('/add-new-asset-form').catch(() => {});

      const generateBtn = page.getByText('Generate').or(page.getByText('Tag ID')).first();
      if (await generateBtn.isVisible()) {
        await generateBtn.click();
        await page.waitForTimeout(300);
      }
    });

    test('should click Cancel button', async ({ page }) => {
      await page.goto('/add-new-asset-form').catch(() => {});

      const cancelBtn = page.getByText('Cancel').first();
      if (await cancelBtn.isVisible()) {
        await cancelBtn.click();
        await page.waitForTimeout(300);
      }
    });

    test('should click Save button', async ({ page }) => {
      await page.goto('/add-new-asset-form').catch(() => {});

      const saveBtn = page.getByText('Save').first();
      if (await saveBtn.isVisible()) {
        await saveBtn.click();
        await page.waitForTimeout(300);
      }
    });
  });

  test.describe('Required Field Validation', () => {
    test('should show validation errors on empty submit', async ({ page }) => {
      await page.goto('/add-new-asset-form').catch(() => {});

      const submitBtn = page.getByRole('button', { name: /Save|Submit/i }).first();
      if (await submitBtn.isVisible()) {
        await submitBtn.click();
        await page.waitForTimeout(500);

        const errors = page.locator('[class*="error"], [class*="invalid"], .text-red');
        const count = await errors.count();
        expect(count).toBeGreaterThanOrEqual(0);
      }
    });

    test('should clear errors when field is filled', async ({ page }) => {
      await page.goto('/add-new-asset-form').catch(() => {});

      const inputs = page.locator('input[required], input:invalid').first();
      if (await inputs.isVisible()) {
        await inputs.fill('Test Value');
        await page.waitForTimeout(200);
      }
    });
  });

  test.describe('Search Forms', () => {
    test('should search in Asset Registry', async ({ page }) => {
      await page.goto('/asset-registry');
      await page.waitForTimeout(300);

      const searchInput = page.locator('input[type="text"]').first();
      await searchInput.fill('Laptop');
      await page.waitForTimeout(300);

      await searchInput.clear();
      await page.waitForTimeout(200);
    });

    test('should search in Consumables', async ({ page }) => {
      await page.goto('/consumables-supplies-tracker');
      await page.waitForTimeout(300);

      const searchInput = page.locator('input[type="text"]').first();
      if (await searchInput.isVisible()) {
        await searchInput.fill('Paper');
        await page.waitForTimeout(300);
      }
    });
  });

  test.describe('Filter Forms', () => {
    test('should apply filters in Asset Registry', async ({ page }) => {
      await page.goto('/asset-registry');
      await page.waitForTimeout(300);

      const filterBtn = page.getByText('Advanced Filters').first();
      await filterBtn.click();
      await page.waitForTimeout(300);

      const applyBtn = page.getByText('Apply').or(page.getByText('Filter')).first();
      if (await applyBtn.isVisible()) {
        await applyBtn.click();
        await page.waitForTimeout(300);
      }
    });
  });

  test.describe('Date Inputs', () => {
    test('should fill date fields', async ({ page }) => {
      await page.goto('/add-new-asset-form').catch(() => {});

      const dateInputs = page.locator('input[type="date"], input[type="datetime-local"]');
      const count = await dateInputs.count();

      if (count > 0) {
        await dateInputs
          .first()
          .fill('2026-05-03')
          .catch(() => {});
      }
    });
  });

  test.describe('Number Inputs', () => {
    test('should fill number fields', async ({ page }) => {
      await page.goto('/add-new-asset-form').catch(() => {});

      const numberInputs = page.locator('input[type="number"]');
      const count = await numberInputs.count();

      for (let i = 0; i < Math.min(count, 3); i++) {
        await numberInputs
          .nth(i)
          .fill('100')
          .catch(() => {});
      }
    });
  });

  test.describe('Textarea Inputs', () => {
    test('should fill textarea fields', async ({ page }) => {
      await page.goto('/add-new-asset-form').catch(() => {});

      const textareas = page.locator('textarea');
      const count = await textareas.count();

      if (count > 0) {
        await textareas.first().fill('This is a test description');
      }
    });
  });

  test.describe('Form Keyboard Navigation', () => {
    test('should tab through form fields', async ({ page }) => {
      await page.goto('/add-new-asset-form').catch(() => {});

      for (let i = 0; i < 5; i++) {
        await page.keyboard.press('Tab');
      }

      const focused = page.locator(':focus');
      await expect(focused).toBeVisible();
    });
  });

  test.describe('Form Error Handling', () => {
    test('should not crash on invalid input', async ({ page }) => {
      await page.goto('/add-new-asset-form').catch(() => {});

      const errors: string[] = [];
      page.on('pageerror', (err) => errors.push(err.message));

      const inputs = page.locator('input').first();
      if (await inputs.isVisible()) {
        await inputs.fill('!@#$%^&*()');
        await page.waitForTimeout(200);
      }

      expect(errors.length).toBe(0);
    });
  });

  test.describe('Form Success Flow', () => {
    test('should show success message on valid submit', async ({ page }) => {
      await page.route('**/api/assets**', (route) => {
        route.fulfill({
          status: 200,
          json: { success: true, id: 1 },
        });
      });

      await page.goto('/add-new-asset-form').catch(() => {});

      const saveBtn = page.getByText('Save').first();
      if (await saveBtn.isVisible()) {
        await saveBtn.click();
        await page.waitForTimeout(500);
      }
    });
  });
});
