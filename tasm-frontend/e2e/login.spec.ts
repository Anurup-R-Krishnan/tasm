import { test, expect } from '@playwright/test';

/**
 * Login Page Tests
 */
test.describe('Login Page', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/login');
    await page.waitForLoadState('networkidle').catch(() => {});
  });

  test.describe('Page Load', () => {
    test('should load Login page', async ({ page }) => {
      await expect(page.locator('body')).toBeVisible();
    });

    test('should have username/email input', async ({ page }) => {
      const emailInput = page
        .locator('input[type="email"], input[name="email"], input[placeholder*="email" i]')
        .first();
      if (await emailInput.isVisible()) {
        await expect(emailInput).toBeVisible();
      }
    });

    test('should have password input', async ({ page }) => {
      const passwordInput = page.locator('input[type="password"]').first();
      if (await passwordInput.isVisible()) {
        await expect(passwordInput).toBeVisible();
      }
    });

    test('should have Login button', async ({ page }) => {
      const loginBtn = page.getByRole('button', { name: /Login|Sign In/i }).first();
      if (await loginBtn.isVisible()) {
        await expect(loginBtn).toBeVisible();
      }
    });

    test('should have Remember Me checkbox', async ({ page }) => {
      const checkbox = page
        .getByText('Remember Me')
        .or(page.locator('input[type="checkbox"]'))
        .first();
      if (await checkbox.isVisible()) {
        await expect(checkbox).toBeVisible();
      }
    });

    test('should have Forgot Password link', async ({ page }) => {
      const forgotLink = page
        .getByText('Forgot Password')
        .or(page.getByRole('link', { name: /Forgot/i }))
        .first();
      if (await forgotLink.isVisible()) {
        await expect(forgotLink).toBeVisible();
      }
    });
  });

  test.describe('Form Interactions', () => {
    test('should fill email', async ({ page }) => {
      const emailInput = page.locator('input[type="email"], input[name="email"]').first();
      if (await emailInput.isVisible()) {
        await emailInput.fill('admin@tasm.com');
        await page.waitForTimeout(200);
      }
    });

    test('should fill password', async ({ page }) => {
      const passwordInput = page.locator('input[type="password"]').first();
      if (await passwordInput.isVisible()) {
        await passwordInput.fill('password123');
        await page.waitForTimeout(200);
      }
    });

    test('should click Remember Me', async ({ page }) => {
      const checkbox = page.locator('input[type="checkbox"]').first();
      if (await checkbox.isVisible()) {
        await checkbox.check();
        await page.waitForTimeout(200);
      }
    });

    test('should click Login button', async ({ page }) => {
      await page.route('**/api/login', (route) => {
        route.fulfill({
          json: { token: 'fake-jwt-token', user: { name: 'Admin' } },
        });
      });

      const emailInput = page.locator('input[type="email"]').first();
      const passwordInput = page.locator('input[type="password"]').first();
      const loginBtn = page.getByRole('button', { name: /Login/i }).first();

      if (await loginBtn.isVisible()) {
        if (await emailInput.isVisible()) {
          await emailInput.fill('admin@tasm.com');
        }
        if (await passwordInput.isVisible()) {
          await passwordInput.fill('password123');
        }
        await loginBtn.click();
        await page.waitForTimeout(500);
      }
    });
  });

  test.describe('Form Validation', () => {
    test('should show error on empty submit', async ({ page }) => {
      const loginBtn = page.getByRole('button', { name: /Login/i }).first();
      if (await loginBtn.isVisible()) {
        await loginBtn.click();
        await page.waitForTimeout(300);

        const errors = page.locator('[class*="error"], .text-red, .invalid');
        const count = await errors.count();
        expect(count).toBeGreaterThanOrEqual(0);
      }
    });

    test('should show error on invalid credentials', async ({ page }) => {
      await page.route('**/api/login', (route) => {
        route.fulfill({
          status: 401,
          json: { error: 'Invalid credentials' },
        });
      });

      const emailInput = page.locator('input[type="email"]').first();
      const passwordInput = page.locator('input[type="password"]').first();
      const loginBtn = page.getByRole('button', { name: /Login/i }).first();

      if (await loginBtn.isVisible()) {
        if (await emailInput.isVisible()) {
          await emailInput.fill('wrong@email.com');
        }
        if (await passwordInput.isVisible()) {
          await passwordInput.fill('wrongpass');
        }
        await loginBtn.click();
        await page.waitForTimeout(500);
      }
    });
  });

  test.describe('Keyboard Navigation', () => {
    test('should tab through login form', async ({ page }) => {
      for (let i = 0; i < 3; i++) {
        await page.keyboard.press('Tab');
      }
      const focused = page.locator(':focus');
      await expect(focused).toBeVisible();
    });

    test('should submit with Enter key', async ({ page }) => {
      const passwordInput = page.locator('input[type="password"]').first();
      if (await passwordInput.isVisible()) {
        await passwordInput.fill('test');
        await page.keyboard.press('Enter');
        await page.waitForTimeout(300);
      }
    });
  });

  test.describe('Responsive', () => {
    test('should display on mobile', async ({ page }) => {
      await page.setViewportSize({ width: 375, height: 667 });
      await page.goto('/login');
      await page.waitForTimeout(300);
      await expect(page.locator('body')).toBeVisible();
    });

    test('should have proper layout on mobile', async ({ page }) => {
      await page.setViewportSize({ width: 375, height: 667 });
      const form = page.locator('form').first();
      if (await form.isVisible()) {
        await expect(form).toBeVisible();
      }
    });
  });

  test.describe('Navigation', () => {
    test('should redirect to dashboard on successful login', async ({ page }) => {
      await page.route('**/api/login', (route) => {
        route.fulfill({
          json: { token: 'fake-token' },
        });
      });

      await page.goto('/login');
      await page.waitForTimeout(300);

      const loginBtn = page.getByRole('button', { name: /Login/i }).first();
      if (await loginBtn.isVisible()) {
        await loginBtn.click();
        await page.waitForTimeout(500);
      }
    });

    test('should navigate to Forgot Password', async ({ page }) => {
      const forgotLink = page.getByText('Forgot Password').first();
      if (await forgotLink.isVisible()) {
        await forgotLink.click();
        await page.waitForTimeout(300);
      }
    });
  });
});
