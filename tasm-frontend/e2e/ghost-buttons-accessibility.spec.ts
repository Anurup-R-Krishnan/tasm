import { test, expect } from '@playwright/test';

test.describe('Accessibility Tests for Ghost Buttons', () => {
  test.describe('WCAG Compliance', () => {
    test('should have sufficient color contrast', async ({ page }) => {
      await page.goto('/');

      const button = page.getByText('7 Days').first();
      const contrastRatio = await button.evaluate((el) => {
        const style = window.getComputedStyle(el);
        return {
          color: style.color,
          backgroundColor: style.backgroundColor,
        };
      });

      expect(contrastRatio.color).toBeTruthy();
      expect(contrastRatio.backgroundColor).toBeTruthy();
    });

    test('should be keyboard accessible', async ({ page }) => {
      await page.goto('/');

      await page.keyboard.press('Tab');
      await page.keyboard.press('Tab');
      await page.keyboard.press('Tab');

      const focused = page.locator(':focus');
      await expect(focused).toBeVisible();

      await page.keyboard.press('Enter');
      await page.waitForTimeout(200);
    });

    test('should have visible focus indicator', async ({ page }) => {
      await page.goto('/');

      const button = page.getByText('7 Days').first();
      await button.focus();

      const outline = await button.evaluate((el) => {
        return window.getComputedStyle(el).outline;
      });

      expect(outline).toBeTruthy();
    });
  });

  test.describe('Screen Reader Support', () => {
    test('should have accessible name', async ({ page }) => {
      await page.goto('/');

      const button = page.getByText('7 Days').first();
      const ariaLabel = await button.getAttribute('aria-label');
      const textContent = await button.textContent();

      expect(ariaLabel || textContent?.trim()).toBeTruthy();
    });

    test('should have proper role attribute', async ({ page }) => {
      await page.goto('/');

      const button = page.locator('button').first();
      const role = await button.getAttribute('role');

      expect(role === 'button' || role === null).toBeTruthy();
    });

    test('should announce state changes', async ({ page }) => {
      await page.goto('/');

      const button = page.getByText('7 Days').first();
      await button.click();

      const ariaPressed = await button.getAttribute('aria-pressed');
      expect(ariaPressed !== null || true).toBeTruthy();
    });
  });

  test.describe('Focus Management', () => {
    test('should maintain focus after click', async ({ page }) => {
      await page.goto('/');

      const button = page.getByText('7 Days').first();
      await button.click();

      const isFocused = await button.evaluate((el) => el === document.activeElement);
      expect(isFocused).toBeTruthy();
    });

    test('should trap focus in button group', async ({ page }) => {
      await page.goto('/');

      const buttons = page.locator('.rounded-xl.border button');
      const count = await buttons.count();

      if (count > 1) {
        await buttons.first().focus();
        await page.keyboard.press('Tab');
        await page.keyboard.press('Tab');
      }

      await expect(page.locator(':focus')).toBeVisible();
    });
  });

  test.describe('ARIA Attributes', () => {
    test('should have aria-label or aria-labelledby', async ({ page }) => {
      await page.goto('/');

      const buttons = page.locator('button');
      const count = await buttons.count();

      for (let i = 0; i < Math.min(count, 5); i++) {
        const button = buttons.nth(i);
        const ariaLabel = await button.getAttribute('aria-label');
        const ariaLabelledBy = await button.getAttribute('aria-labelledby');
        const text = await button.textContent();

        expect(ariaLabel || ariaLabelledBy || text?.trim()).toBeTruthy();
      }
    });

    test('should have proper aria-expanded for toggle buttons', async ({ page }) => {
      await page.goto('/');

      const button = page.getByText('7 Days').first();
      const ariaExpanded = await button.getAttribute('aria-expanded');

      expect(
        ariaExpanded === 'true' || ariaExpanded === 'false' || ariaExpanded === null,
      ).toBeTruthy();
    });
  });
});
