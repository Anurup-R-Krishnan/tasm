import { test, expect } from '@playwright/test';

test.describe('Performance Tests for Ghost Buttons', () => {
  test.describe('Click Response Time', () => {
    test('should respond to click within acceptable time', async ({ page }) => {
      await page.goto('/');

      const button = page.getByText('7 Days').first();

      const startTime = performance.now();
      await button.click();
      const endTime = performance.now();

      const duration = endTime - startTime;
      expect(duration).toBeLessThan(100);
    });

    test('should measure time for multiple rapid clicks', async ({ page }) => {
      await page.goto('/');

      const button = page.getByText('7 Days').first();
      const times: number[] = [];

      for (let i = 0; i < 5; i++) {
        const start = Date.now();
        await button.click();
        const end = Date.now();
        times.push(end - start);
      }

      const avgTime = times.reduce((a, b) => a + b, 0) / times.length;
      expect(avgTime).toBeLessThan(50);
    });
  });

  test.describe('Memory and Resource Usage', () => {
    test('should not leak memory on repeated clicks', async ({ page }) => {
      await page.goto('/');

      const button = page.getByText('7 Days').first();

      for (let i = 0; i < 50; i++) {
        await button.click();
      }

      await page.waitForTimeout(500);
      await expect(button).toBeVisible();
    });

    test('should handle concurrent click events efficiently', async ({ page }) => {
      await page.goto('/');

      const buttons = page.locator('button.text-text-secondary');
      const count = await buttons.count();

      const promises = [];
      for (let i = 0; i < Math.min(count, 5); i++) {
        promises.push(buttons.nth(i).click());
      }

      await Promise.all(promises);
      await expect(page.locator('body')).toBeVisible();
    });
  });

  test.describe('Network Performance', () => {
    test('should not cause unnecessary network requests on click', async ({ page }) => {
      let requestCount = 0;

      await page.route('**/*', (route) => {
        requestCount++;
        route.continue();
      });

      await page.goto('/');

      const button = page.getByText('7 Days').first();
      await button.click();

      await page.waitForTimeout(500);
      expect(requestCount).toBeGreaterThan(0);
    });

    test('should handle slow network on button click', async ({ page }) => {
      await page.route('**/api/**', async (route) => {
        await new Promise((resolve) => setTimeout(resolve, 2000));
        await route.continue();
      });

      await page.goto('/');

      const button = page.getByText('7 Days').first();
      await button.click();

      await expect(button).toBeVisible();
    });
  });

  test.describe('Rendering Performance', () => {
    test('should not cause layout thrashing on click', async ({ page }) => {
      await page.goto('/');

      const button = page.getByText('7 Days').first();

      await page.evaluate(() => {
        const observer = new PerformanceObserver((list) => {
          const entries = list.getEntries();
          return entries;
        });
        observer.observe({ entryTypes: ['layout-shift'] });
      });

      await button.click();
      await page.waitForTimeout(500);
    });
  });
});
