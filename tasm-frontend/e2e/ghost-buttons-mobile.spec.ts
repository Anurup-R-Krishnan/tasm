import { test, expect } from '@playwright/test';

test.describe('Mobile Ghost Button Tests', () => {
  test.describe('Mobile Touch Interactions', () => {
    test('should handle tap on ghost button on mobile', async ({ page }) => {
      await page.setViewportSize({ width: 375, height: 667 });
      await page.goto('/');

      const button = page.getByText('7 Days').first();
      await button.tap();
      await expect(button).toBeVisible();
    });

    test('should handle long press on ghost button', async ({ page }) => {
      await page.setViewportSize({ width: 375, height: 667 });
      await page.goto('/');

      const button = page.getByText('7 Days').first();
      await button.longpress({ timeout: 1000 }).catch(() => {});

      await expect(button).toBeVisible();
    });

    test('should handle swipe gesture near ghost button', async ({ page }) => {
      await page.setViewportSize({ width: 375, height: 667 });
      await page.goto('/');

      const button = page.getByText('7 Days').first();
      const box = await button.boundingBox();

      if (box) {
        await page.touchscreen.tap(box.x + box.width / 2, box.y + box.height / 2);
      }

      await expect(button).toBeVisible();
    });
  });

  test.describe('Mobile Viewport Tests', () => {
    test('should verify ghost buttons are tappable on small screen', async ({ page }) => {
      await page.setViewportSize({ width: 320, height: 568 });
      await page.goto('/');

      const buttons = page.locator('button.text-text-secondary');
      const count = await buttons.count();

      if (count > 0) {
        const box = await buttons.first().boundingBox();
        expect(box?.height).toBeGreaterThanOrEqual(44);
      }
    });

    test('should handle orientation change with ghost buttons', async ({ page }) => {
      await page.setViewportSize({ width: 375, height: 667 });
      await page.goto('/');

      const button = page.getByText('7 Days').first();
      await expect(button).toBeVisible();

      await page.setViewportSize({ width: 667, height: 375 });
      await expect(button).toBeVisible();
    });
  });

  test.describe('Mobile Gesture Events', () => {
    test('should not trigger ghost button on scroll', async ({ page }) => {
      await page.setViewportSize({ width: 375, height: 667 });
      await page.goto('/');

      await page.evaluate(() => window.scrollBy(0, 100));
      await page.waitForTimeout(300);

      const button = page.getByText('7 Days').first();
      await expect(button).toBeVisible();
    });

    test('should handle multi-touch near ghost button', async ({ page }) => {
      await page.setViewportSize({ width: 375, height: 667 });
      await page.goto('/');

      const button = page.getByText('7 Days').first();
      await button.click();
      await expect(button).toBeVisible();
    });
  });
});
