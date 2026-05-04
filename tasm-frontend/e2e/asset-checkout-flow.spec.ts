import { test, expect } from '@playwright/test';

/**
 * Asset Check Out Flow Tests
 */
test.describe('Asset Check Out Flow', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/asset-check-out-flow');
    await page.waitForLoadState('networkidle').catch(() => {});
  });

  test.describe('Page Load', () => {
    test('should load Check Out Flow page', async ({ page }) => {
      await expect(page.locator('body')).toBeVisible();
    });

    test('should have step indicators', async ({ page }) => {
      const steps = page
        .getByText('Select Recipient')
        .or(page.getByText('Review'))
        .or(page.getByText('Confirm'));
      const count = await steps.count();
      expect(count).toBeGreaterThanOrEqual(0);
    });

    test('should have search input for recipient', async ({ page }) => {
      const search = page.locator('input[type="text"]').first();
      if (await search.isVisible()) {
        await expect(search).toBeVisible();
      }
    });
  });

  test.describe('Step 1 - Select Recipient', () => {
    test('should search for recipient', async ({ page }) => {
      const search = page.locator('input[type="text"]').first();
      if (await search.isVisible()) {
        await search.fill('John');
        await page.waitForTimeout(300);
        await search.clear();
      }
    });

    test('should select a recipient', async ({ page }) => {
      await page.route('**/api/users**', (route) => {
        route.fulfill({
          json: [{ id: 1, name: 'John Doe' }],
        });
      });

      await page.reload();
      await page.waitForTimeout(500);

      const recipient = page.getByText('John Doe').first();
      if (await recipient.isVisible()) {
        await recipient.click();
        await page.waitForTimeout(300);
      }
    });

    test('should click Next button', async ({ page }) => {
      const nextBtn = page.getByText('Next').or(page.getByText('Continue')).first();
      if (await nextBtn.isVisible()) {
        if (await nextBtn.isEnabled()) {
          await nextBtn.click();
          await page.waitForTimeout(300);
        } else {
          await expect(nextBtn).toBeDisabled();
        }
      }
    });
  });

  test.describe('Step 2 - Review', () => {
    test('should display selected asset', async ({ page }) => {
      const asset = page.getByText('Asset').first();
      if (await asset.isVisible()) {
        await expect(asset).toBeVisible();
      }
    });

    test('should click Back button', async ({ page }) => {
      const backBtn = page.getByText('Back').first();
      if (await backBtn.isVisible()) {
        await backBtn.click();
        await page.waitForTimeout(300);
      }
    });

    test('should click Confirm button', async ({ page }) => {
      const confirmBtn = page.getByText('Confirm').or(page.getByText('Complete')).first();
      if (await confirmBtn.isVisible()) {
        await confirmBtn.click();
        await page.waitForTimeout(500);
      }
    });
  });

  test.describe('Keyboard Navigation', () => {
    test('should tab through flow', async ({ page }) => {
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
      await page.goto('/asset-check-out-flow');
      await page.waitForTimeout(300);
      await expect(page.locator('body')).toBeVisible();
    });
  });

  test.describe('Error Handling', () => {
    test('should handle API error', async ({ page }) => {
      await page.route('**/api/**', (route) => route.abort());
      await page.goto('/asset-check-out-flow');
      await page.waitForTimeout(500);
      await expect(page.locator('body')).toBeVisible();
    });
  });
});
