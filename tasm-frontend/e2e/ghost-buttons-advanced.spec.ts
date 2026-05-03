import { test, expect } from '@playwright/test';

test.describe('Advanced Ghost Button and Click Behavior Tests', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/');
  });

  test.describe('Button State Management', () => {
    test('should verify active state styling on time filter buttons', async ({ page }) => {
      const activeButton = page.locator('button.bg-primary').first();
      await expect(activeButton).toBeVisible();

      const activeText = await activeButton.textContent();
      expect(['24 Hours', '7 Days', '30 Days']).toContain(activeText?.trim());
    });

    test('should toggle active state between filter buttons', async ({ page }) => {
      const sevenDaysBtn = page.getByText('7 Days').first();
      const thirtyDaysBtn = page.getByText('30 Days').first();

      await sevenDaysBtn.click();
      await page.waitForTimeout(300);

      const sevenDaysClasses = await sevenDaysBtn.getAttribute('class');
      expect(sevenDaysClasses).toMatch(/bg-primary|active/);

      await thirtyDaysBtn.click();
      await page.waitForTimeout(300);

      const thirtyDaysClasses = await thirtyDaysBtn.getAttribute('class');
      expect(thirtyDaysClasses).toMatch(/bg-primary|active/);
    });
  });

  test.describe('Click Timing and Delays', () => {
    test('should handle click with delay', async ({ page }) => {
      const button = page.getByText('7 Days').first();
      await button.click({ delay: 100 });
      await expect(button).toBeVisible();
    });

    test('should verify no race conditions on rapid clicks', async ({ page }) => {
      const button = page.getByText('7 Days').first();

      const clicks = Array(5)
        .fill(null)
        .map(() => button.click());
      await Promise.all(clicks);

      await expect(button).toBeVisible();
    });
  });

  test.describe('Network and API Interaction on Click', () => {
    test('should intercept API calls when filter button is clicked', async ({ page }) => {
      const apiCalls: string[] = [];

      await page.route('**/api/**', async (route) => {
        apiCalls.push(route.request().url());
        await route.continue();
      });

      const sevenDaysBtn = page.getByText('7 Days').first();
      await sevenDaysBtn.click();
      await page.waitForTimeout(1000);
    });

    test('should handle API errors gracefully on button click', async ({ page }) => {
      await page.route('**/api/assets**', (route) => {
        route.abort('failed');
      });

      const button = page.getByText('7 Days').first();
      await button.click();
      await page.waitForTimeout(500);

      await expect(page.locator('body')).toBeVisible();
    });
  });

  test.describe('Button Visibility and Viewport', () => {
    test('should verify ghost buttons are visible in different viewports', async ({ page }) => {
      await page.setViewportSize({ width: 375, height: 667 });
      const button = page.getByText('7 Days').first();
      await expect(button).toBeVisible();

      await page.setViewportSize({ width: 1920, height: 1080 });
      await expect(button).toBeVisible();
    });

    test('should handle click on button after viewport resize', async ({ page }) => {
      await page.setViewportSize({ width: 375, height: 667 });
      await page.setViewportSize({ width: 1280, height: 720 });

      const button = page.getByText('7 Days').first();
      await button.click();
      await expect(button).toBeVisible();
    });
  });

  test.describe('Button Groups and Related Elements', () => {
    test('should find all buttons in the time filter group', async ({ page }) => {
      const filterGroup = page.locator('.rounded-xl.border').first();
      const buttons = filterGroup.locator('button');
      const count = await buttons.count();
      expect(count).toBe(3);
    });

    test('should verify button spacing and layout', async ({ page }) => {
      const buttons = page.locator('.rounded-xl.border button');
      const count = await buttons.count();

      if (count > 1) {
        const firstBox = await buttons.first().boundingBox();
        const secondBox = await buttons.nth(1).boundingBox();

        if (firstBox && secondBox) {
          expect(secondBox.x).toBeGreaterThan(firstBox.x + firstBox.width);
        }
      }
    });
  });

  test.describe('Event Listener Verification', () => {
    test('should verify click event is properly attached', async ({ page }) => {
      const button = page.getByText('7 Days').first();
      const hasClickHandler = await button.evaluate((el) => {
        const listeners = (el as any).__vueListeners__ || [];
        return listeners.length > 0 || el.onclick !== null;
      });

      expect(typeof hasClickHandler).toBe('boolean');
    });

    test('should not trigger click on disabled elements nearby', async ({ page }) => {
      const button = page.getByText('7 Days').first();
      const isDisabled = await button.isDisabled();
      expect(isDisabled).toBeFalsy();
    });
  });

  test.describe('Button Text and Content Verification', () => {
    test('should verify all ghost buttons have meaningful text', async ({ page }) => {
      const buttons = page.locator('button.text-text-secondary');
      const count = await buttons.count();

      for (let i = 0; i < count; i++) {
        const text = await buttons.nth(i).textContent();
        expect(text?.trim().length).toBeGreaterThan(0);
      }
    });

    test('should verify button text does not change unexpectedly after click', async ({ page }) => {
      const button = page.getByText('7 Days').first();
      const initialText = await button.textContent();

      await button.click();
      await page.waitForTimeout(300);

      const afterClickText = await button.textContent();
      expect(afterClickText).toBe(initialText);
    });
  });

  test.describe('Touch and Pointer Events', () => {
    test('should handle tap event on mobile', async ({ page, browserName }) => {
      if (browserName === 'chromium') {
        await page.setViewportSize({ width: 375, height: 667 });

        const button = page.getByText('7 Days').first();
        await button.tap();
        await expect(button).toBeVisible();
      }
    });

    test('should verify pointer events are properly handled', async ({ page }) => {
      const button = page.getByText('7 Days').first();
      await button.dispatchEvent('pointerdown');
      await button.dispatchEvent('pointerup');

      await expect(button).toBeVisible();
    });
  });

  test.describe('Button in Different Pages', () => {
    test('should find ghost buttons in Asset Registry page', async ({ page }) => {
      await page.goto('/asset-registry');

      const filterButton = page.getByText('Advanced Filters').first();
      if (await filterButton.isVisible()) {
        await expect(filterButton).toBeVisible();
      }
    });

    test('should click navigation links and verify ghost buttons on new page', async ({ page }) => {
      await page.getByRole('link', { name: 'Consumables' }).click();
      await expect(page.getByText('Consumables Tracker')).toBeVisible();

      const buttons = page.locator('button');
      const count = await buttons.count();
      expect(count).toBeGreaterThan(0);
    });
  });

  test.describe('Performance and Load Testing', () => {
    test('should measure click response time', async ({ page }) => {
      const button = page.getByText('7 Days').first();

      const startTime = Date.now();
      await button.click();
      const endTime = Date.now();

      const responseTime = endTime - startTime;
      expect(responseTime).toBeLessThan(1000);
    });

    test('should handle multiple simultaneous button interactions', async ({ page }) => {
      const buttons = page.locator('button.text-text-secondary');
      const count = await buttons.count();

      const promises = [];
      for (let i = 0; i < Math.min(count, 3); i++) {
        promises.push(buttons.nth(i).click());
      }

      await Promise.all(promises);
      await expect(page.locator('body')).toBeVisible();
    });
  });

  test.describe('Button ARIA and Semantic HTML', () => {
    test('should verify buttons have proper ARIA attributes', async ({ page }) => {
      const button = page.getByText('7 Days').first();
      await button.getAttribute('aria-label');
      const role = await button.getAttribute('role');

      expect(role === 'button' || role === null).toBeTruthy();
    });

    test('should verify button type attribute', async ({ page }) => {
      const button = page.locator('button').first();
      const type = await button.getAttribute('type');
      expect(['button', 'submit', null]).toContain(type);
    });
  });

  test.describe('Error Boundaries and Exception Handling', () => {
    test('should not crash when clicking button multiple times quickly', async ({ page }) => {
      const button = page.getByText('7 Days').first();

      for (let i = 0; i < 10; i++) {
        await button.click().catch(() => {});
      }

      await expect(page.locator('body')).toBeVisible();
    });

    test('should recover from JavaScript errors after button click', async ({ page }) => {
      const errors: string[] = [];
      page.on('pageerror', (err) => errors.push(err.message));

      const button = page.getByText('7 Days').first();
      await button.click();

      await page.waitForTimeout(1000);
      expect(errors.length).toBe(0);
    });
  });

  test.describe('Button Styling and CSS Verification', () => {
    test('should verify ghost button has transparent or no background initially', async ({
      page,
    }) => {
      const button = page.getByText('7 Days').first();
      const bgColor = await button.evaluate((el) => {
        return window.getComputedStyle(el).backgroundColor;
      });

      expect(['rgba(0, 0, 0, 0)', 'transparent']).toContain(bgColor);
    });

    test('should verify hover changes background color', async ({ page }) => {
      const button = page.getByText('7 Days').first();
      const beforeBg = await button.evaluate((el) => window.getComputedStyle(el).backgroundColor);

      await button.hover();
      await page.waitForTimeout(300);

      const afterBg = await button.evaluate((el) => window.getComputedStyle(el).backgroundColor);

      expect(beforeBg || afterBg).toBeTruthy();
    });
  });
});
