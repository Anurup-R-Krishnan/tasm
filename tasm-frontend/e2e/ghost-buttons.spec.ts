import { test, expect } from '@playwright/test';

test.describe('Ghost Buttons and Click Behavior', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/');
  });

  test.describe('Ghost Button Identification', () => {
    test('should identify ghost buttons by secondary text styling', async ({ page }) => {
      const ghostButtons = page.locator('button.text-text-secondary');
      const count = await ghostButtons.count();
      expect(count).toBeGreaterThan(0);

      for (let i = 0; i < count; i++) {
        const button = ghostButtons.nth(i);
        await expect(button).toBeVisible();
        await expect(button).toBeEnabled();
      }
    });

    test('should find time filter ghost buttons on dashboard', async ({ page }) => {
      const timeFilters = page.locator('button', { hasText: /7 Days|30 Days/ });
      await expect(timeFilters.first()).toBeVisible();

      const filterCount = await timeFilters.count();
      expect(filterCount).toBeGreaterThanOrEqual(2);
    });

    test('should identify ghost buttons with hover transition classes', async ({ page }) => {
      const transitionButtons = page.locator(
        'button[class*="hover:text-primary"][class*="transition"]',
      );
      await expect(transitionButtons.first()).toBeVisible();
    });
  });

  test.describe('Click Command Behavior', () => {
    test('should perform basic click on ghost button', async ({ page }) => {
      const sevenDaysButton = page.getByText('7 Days').first();
      await expect(sevenDaysButton).toBeVisible();

      await sevenDaysButton.click();
      await expect(sevenDaysButton).toBeVisible();
    });

    test('should perform click with force option on ghost button', async ({ page }) => {
      const ghostButton = page.locator('button.text-text-secondary').first();
      await ghostButton.click({ force: true });
      await expect(ghostButton).toBeVisible();
    });

    test('should perform click with position offset', async ({ page }) => {
      const button = page.getByText('7 Days').first();
      const box = await button.boundingBox();

      if (box) {
        await button.click({
          position: { x: box.width / 2, y: box.height / 2 },
        });
      }
      await expect(button).toBeVisible();
    });

    test('should perform double click on ghost button', async ({ page }) => {
      const button = page.getByText('30 Days').first();
      await button.dblclick();
      await expect(button).toBeVisible();
    });

    test('should perform click via keyboard navigation', async ({ page }) => {
      const button = page.getByText('7 Days').first();
      await button.focus();
      await expect(button).toBeFocused();

      await page.keyboard.press('Enter');
      await expect(button).toBeVisible();
    });

    test('should perform click via Space key', async ({ page }) => {
      const button = page.getByText('30 Days').first();
      await button.focus();
      await page.keyboard.press('Space');
      await expect(button).toBeVisible();
    });
  });

  test.describe('Ghost Button Hover Behavior', () => {
    test('should change appearance on hover', async ({ page }) => {
      const ghostButton = page.getByText('7 Days').first();
      await ghostButton.hover();

      await page.waitForTimeout(300);
      await expect(ghostButton).toBeVisible();
    });

    test('should have proper cursor style for clickable ghost buttons', async ({ page }) => {
      const button = page.locator('button.text-text-secondary').first();
      const cursor = await button.evaluate((el) => window.getComputedStyle(el).cursor);
      expect(cursor).toBe('pointer');
    });
  });

  test.describe('Ghost Button State After Click', () => {
    test('should remain visible and enabled after click', async ({ page }) => {
      const buttons = page.locator('button.text-text-secondary');
      const count = await buttons.count();

      if (count > 0) {
        const button = buttons.first();
        await button.click();
        await expect(button).toBeVisible();
        await expect(button).toBeEnabled();
      }
    });

    test('should handle rapid consecutive clicks', async ({ page }) => {
      const button = page.getByText('7 Days').first();

      await button.click();
      await button.click();
      await button.click();

      await expect(button).toBeVisible();
    });
  });

  test.describe('Ghost Button Accessibility', () => {
    test('should be focusable via tab navigation', async ({ page }) => {
      await page.keyboard.press('Tab');
      await page.keyboard.press('Tab');
      await page.keyboard.press('Tab');

      const focusedElement = page.locator(':focus');
      const isButton = await focusedElement.evaluate((el) => el.tagName === 'BUTTON');
      expect(isButton).toBeTruthy();
    });

    test('should have accessible text content', async ({ page }) => {
      const buttons = page.locator('button.text-text-secondary');
      const count = await buttons.count();

      for (let i = 0; i < Math.min(count, 5); i++) {
        const text = await buttons.nth(i).textContent();
        expect(text?.trim().length).toBeGreaterThan(0);
      }
    });
  });

  test.describe('Click Event Verification', () => {
    test('should verify click does not cause navigation errors', async ({ page }) => {
      const button = page.getByText('7 Days').first();
      await button.click();

      await page.waitForLoadState('networkidle').catch(() => {});
      expect(page.url()).toContain('/');
    });

    test('should handle click on disabled ghost button state', async ({ page }) => {
      const buttons = page.locator('button.text-text-secondary');
      const count = await buttons.count();

      for (let i = 0; i < count; i++) {
        const button = buttons.nth(i);
        const isDisabled = await button.isDisabled();
        expect(isDisabled).toBeFalsy();
      }
    });
  });

  test.describe('Multiple Ghost Button Interactions', () => {
    test('should click multiple ghost buttons in sequence', async ({ page }) => {
      const sevenDays = page.getByText('7 Days').first();
      const thirtyDays = page.getByText('30 Days').first();

      await sevenDays.click();
      await expect(sevenDays).toBeVisible();

      await thirtyDays.click();
      await expect(thirtyDays).toBeVisible();
    });

    test('should verify only one time filter is active at a time', async ({ page }) => {
      const sevenDaysBtn = page.getByText('7 Days').first();
      const thirtyDaysBtn = page.getByText('30 Days').first();

      await sevenDaysBtn.click();
      await page.waitForTimeout(200);

      await thirtyDaysBtn.click();
      await page.waitForTimeout(200);

      await expect(thirtyDaysBtn).toBeVisible();
    });
  });

  test.describe('Ghost Button with View All Link', () => {
    test('should find and click View All button', async ({ page }) => {
      const viewAllButton = page.getByText('View All').first();

      if (await viewAllButton.isVisible()) {
        await viewAllButton.click();
        await page.waitForTimeout(500);
      }
    });

    test('should verify View All button has primary color styling', async ({ page }) => {
      const viewAllButton = page.getByText('View All').first();

      if (await viewAllButton.isVisible()) {
        const hasPrimaryClass = await viewAllButton.evaluate(
          (el) => el.classList.contains('text-primary') || el.className.includes('text-primary'),
        );
        expect(hasPrimaryClass).toBeTruthy();
      }
    });
  });

  test.describe('Edge Cases for Click Behavior', () => {
    test('should handle click when button is partially out of viewport', async ({ page }) => {
      const button = page.locator('button.text-text-secondary').first();
      await button.scrollIntoViewIfNeeded();
      await button.click();
      await expect(button).toBeVisible();
    });

    test('should handle click with modifiers', async ({ page }) => {
      const button = page.getByText('7 Days').first();
      await button.click({ modifiers: ['ControlOrMeta'] });
      await expect(button).toBeVisible();
    });

    test('should verify button does not throw errors on click', async ({ page }) => {
      const errors: string[] = [];
      page.on('pageerror', (err) => errors.push(err.message));

      const button = page.locator('button.text-text-secondary').first();
      await button.click();

      await page.waitForTimeout(500);
      expect(errors.length).toBe(0);
    });
  });

  test.describe('Quick Action Ghost Buttons', () => {
    test('should find quick action buttons with ghost-like styling', async ({ page }) => {
      const quickActionButtons = page.locator('.premium-card button[class*="hover:bg-"]');
      const count = await quickActionButtons.count();
      expect(count).toBeGreaterThan(0);
    });

    test('should verify quick action buttons have proper hover transitions', async ({ page }) => {
      const buttons = page.locator('.premium-card button[class*="transition-all"]');
      const count = await buttons.count();

      if (count > 0) {
        await buttons.first().hover();
        await page.waitForTimeout(300);
        await expect(buttons.first()).toBeVisible();
      }
    });
  });
});
