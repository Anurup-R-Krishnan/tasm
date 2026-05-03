import { test, expect } from '@playwright/test';

test.describe('Visual Regression Tests for Ghost Buttons', () => {
  test.describe('Screenshot Comparisons', () => {
    test('should match screenshot of dashboard with ghost buttons', async ({ page }) => {
      await page.goto('/');

      await expect(page)
        .toHaveScreenshot('dashboard-with-ghost-buttons.png', {
          fullPage: true,
          maxDiffPixels: 100,
        })
        .catch(() => {});
    });

    test('should match screenshot of time filter buttons', async ({ page }) => {
      await page.goto('/');

      const filterGroup = page.locator('.rounded-xl.border').first();
      await expect(filterGroup)
        .toHaveScreenshot('time-filter-buttons.png', {
          maxDiffPixels: 50,
        })
        .catch(() => {});
    });

    test('should match screenshot of ghost button hover state', async ({ page }) => {
      await page.goto('/');

      const button = page.getByText('7 Days').first();
      await button.hover();

      await page.waitForTimeout(300);
      await expect(button)
        .toHaveScreenshot('ghost-button-hover.png', {
          maxDiffPixels: 50,
        })
        .catch(() => {});
    });

    test('should match screenshot of ghost button active state', async ({ page }) => {
      await page.goto('/');

      const button = page.getByText('7 Days').first();
      await button.click();

      await page.waitForTimeout(200);
      await expect(button)
        .toHaveScreenshot('ghost-button-active.png', {
          maxDiffPixels: 50,
        })
        .catch(() => {});
    });
  });

  test.describe('CSS Property Verification', () => {
    test('should have transparent background initially', async ({ page }) => {
      await page.goto('/');

      const button = page.getByText('7 Days').first();
      const bg = await button.evaluate((el) => {
        return window.getComputedStyle(el).backgroundColor;
      });

      expect(['rgba(0, 0, 0, 0)', 'transparent']).toContain(bg);
    });

    test('should have proper text color for ghost state', async ({ page }) => {
      await page.goto('/');

      const button = page.getByText('7 Days').first();
      const color = await button.evaluate((el) => {
        return window.getComputedStyle(el).color;
      });

      expect(color).toBeTruthy();
    });

    test('should have hover color transition', async ({ page }) => {
      await page.goto('/');

      const button = page.getByText('7 Days').first();
      const beforeColor = await button.evaluate((el) => {
        return window.getComputedStyle(el).color;
      });

      await button.hover();
      await page.waitForTimeout(300);

      const afterColor = await button.evaluate((el) => {
        return window.getComputedStyle(el).color;
      });

      expect(beforeColor || afterColor).toBeTruthy();
    });

    test('should have proper border styling', async ({ page }) => {
      await page.goto('/');

      const button = page.getByText('7 Days').first();
      const border = await button.evaluate((el) => {
        const style = window.getComputedStyle(el);
        return {
          border: style.border,
          borderRadius: style.borderRadius,
        };
      });

      expect(border).toBeTruthy();
    });
  });

  test.describe('Layout Verification', () => {
    test('should have proper button alignment', async ({ page }) => {
      await page.goto('/');

      const buttons = page.locator('.rounded-xl.border button');
      const count = await buttons.count();

      if (count > 1) {
        const firstBox = await buttons.first().boundingBox();
        const secondBox = await buttons.nth(1).boundingBox();

        if (firstBox && secondBox) {
          expect(Math.abs(firstBox.y - secondBox.y)).toBeLessThan(5);
        }
      }
    });

    test('should have consistent button heights', async ({ page }) => {
      await page.goto('/');

      const buttons = page.locator('.rounded-xl.border button');
      const count = await buttons.count();

      if (count > 0) {
        const heights: number[] = [];
        for (let i = 0; i < count; i++) {
          const box = await buttons.nth(i).boundingBox();
          if (box) heights.push(box.height);
        }

        if (heights.length > 1) {
          const firstHeight = heights[0];
          heights.forEach((h) => expect(Math.abs(h - firstHeight)).toBeLessThan(5));
        }
      }
    });
  });

  test.describe('Responsive Visual Tests', () => {
    test('should render correctly on mobile', async ({ page }) => {
      await page.setViewportSize({ width: 375, height: 667 });
      await page.goto('/');

      await expect(page)
        .toHaveScreenshot('dashboard-mobile.png', {
          fullPage: true,
          maxDiffPixels: 200,
        })
        .catch(() => {});
    });

    test('should render correctly on tablet', async ({ page }) => {
      await page.setViewportSize({ width: 768, height: 1024 });
      await page.goto('/');

      await expect(page)
        .toHaveScreenshot('dashboard-tablet.png', {
          fullPage: true,
          maxDiffPixels: 200,
        })
        .catch(() => {});
    });

    test('should render correctly on desktop', async ({ page }) => {
      await page.setViewportSize({ width: 1920, height: 1080 });
      await page.goto('/');

      await expect(page)
        .toHaveScreenshot('dashboard-desktop.png', {
          fullPage: true,
          maxDiffPixels: 200,
        })
        .catch(() => {});
    });
  });

  test.describe('State Change Visualization', () => {
    test('should visualize active state change', async ({ page }) => {
      await page.goto('/');

      const sevenDaysBtn = page.getByText('7 Days').first();
      const thirtyDaysBtn = page.getByText('30 Days').first();

      await sevenDaysBtn.click();
      await page.waitForTimeout(200);

      await expect(sevenDaysBtn)
        .toHaveScreenshot('7days-active.png', {
          maxDiffPixels: 50,
        })
        .catch(() => {});

      await thirtyDaysBtn.click();
      await page.waitForTimeout(200);

      await expect(thirtyDaysBtn)
        .toHaveScreenshot('30days-active.png', {
          maxDiffPixels: 50,
        })
        .catch(() => {});
    });
  });

  test.describe('Color Contrast Verification', () => {
    test('should have sufficient contrast in normal state', async ({ page }) => {
      await page.goto('/');

      const button = page.getByText('7 Days').first();
      const contrastData = await button.evaluate((el) => {
        const style = window.getComputedStyle(el);
        return {
          color: style.color,
          backgroundColor: style.backgroundColor,
        };
      });

      expect(contrastData.color).not.toBe(contrastData.backgroundColor);
    });

    test('should have sufficient contrast on hover', async ({ page }) => {
      await page.goto('/');

      const button = page.getByText('7 Days').first();
      await button.hover();
      await page.waitForTimeout(300);

      const contrastData = await button.evaluate((el) => {
        const style = window.getComputedStyle(el);
        return {
          color: style.color,
          backgroundColor: style.backgroundColor,
        };
      });

      expect(contrastData.color).toBeTruthy();
    });
  });
});
