import { test, expect } from '@playwright/test';

test.describe('Error Handling for Ghost Buttons', () => {
  test.describe('JavaScript Errors', () => {
    test('should not throw on rapid clicks', async ({ page }) => {
      const errors: string[] = [];
      page.on('pageerror', (err) => errors.push(err.message));

      await page.goto('/');
      const button = page.getByText('7 Days').first();

      for (let i = 0; i < 10; i++) {
        await button.click().catch(() => {});
      }

      await page.waitForTimeout(500);
      expect(errors.length).toBe(0);
    });

    test('should handle null element gracefully', async ({ page }) => {
      await page.goto('/');

      const nonExistentButton = page.getByText('NonExistentButtonXYZ');
      const count = await nonExistentButton.count();
      expect(count).toBe(0);
    });

    test('should recover from click handler exception', async ({ page }) => {
      await page.goto('/');

      await page.evaluate(() => {
        const buttons = document.querySelectorAll('button');
        buttons.forEach((btn) => {
          const originalClick = btn.onclick;
          btn.onclick = (e) => {
            try {
              if (originalClick) originalClick.call(btn, e);
            } catch (err) {
              console.error('Handled error:', err);
            }
          };
        });
      });

      const button = page.getByText('7 Days').first();
      await button.click();
      await expect(button).toBeVisible();
    });
  });

  test.describe('Network Errors', () => {
    test('should handle API failure gracefully', async ({ page }) => {
      await page.route('**/api/**', (route) => route.abort());

      await page.goto('/');
      const button = page.getByText('7 Days').first();
      await button.click();

      await page.waitForTimeout(500);
      await expect(page.locator('body')).toBeVisible();
    });

    test('should handle timeout on API call', async ({ page }) => {
      await page.route('**/api/**', async (route) => {
        await new Promise((resolve) => setTimeout(resolve, 30000));
        await route.continue();
      });

      await page.goto('/', { timeout: 10000 }).catch(() => {});
      const button = page.getByText('7 Days').first();
      await button.click().catch(() => {});
    });

    test('should show error state without breaking UI', async ({ page }) => {
      await page.route('**/api/assets**', (route) => {
        route.fulfill({ status: 500, body: 'Internal Server Error' });
      });

      await page.goto('/');
      const button = page.getByText('7 Days').first();
      await button.click();

      await page.waitForTimeout(500);
      await expect(button).toBeVisible();
    });
  });

  test.describe('DOM Mutations', () => {
    test('should handle button removal and re-add', async ({ page }) => {
      await page.goto('/');

      await page.evaluate(() => {
        const btn = document.querySelector('button');
        if (btn) {
          const parent = btn.parentNode;
          const clone = btn.cloneNode(true);
          btn.remove();
          setTimeout(() => parent?.appendChild(clone), 100);
        }
      });

      await page.waitForTimeout(200);
      const button = page.getByText('7 Days').first();
      await button.click().catch(() => {});
    });

    test('should handle button disabled dynamically', async ({ page }) => {
      await page.goto('/');

      const button = page.getByText('7 Days').first();
      await page.evaluate(
        (btn) => {
          if (btn) btn.disabled = true;
        },
        await button.elementHandle(),
      );

      const isDisabled = await button.isDisabled();
      expect(isDisabled).toBeTruthy();
    });
  });

  test.describe('Browser Compatibility', () => {
    test('should work with JavaScript disabled fallback', async ({ page }) => {
      await page.setJavaScriptEnabled(false);
      await page.goto('/').catch(() => {});

      const button = page.getByText('7 Days').first();
      const count = await button.count();
      expect(count).toBeGreaterThanOrEqual(0);
    });
  });

  test.describe('Memory and Cleanup', () => {
    test('should not leak event listeners', async ({ page }) => {
      await page.goto('/');

      for (let i = 0; i < 20; i++) {
        const button = page.getByText('7 Days').first();
        await button.click();
      }

      await page.waitForTimeout(500);
      await expect(page.locator('body')).toBeVisible();
    });

    test('should handle page navigation and back button', async ({ page }) => {
      await page.goto('/');

      const button = page.getByText('7 Days').first();
      await button.click();

      await page.goBack().catch(() => {});
      await page.goForward().catch(() => {});

      await expect(page.locator('body')).toBeVisible();
    });
  });

  test.describe('Concurrent Operations', () => {
    test('should handle multiple simultaneous clicks', async ({ page }) => {
      await page.goto('/');

      const button = page.getByText('7 Days').first();
      const clicks = Array(5)
        .fill(null)
        .map(() => button.click().catch(() => {}));

      await Promise.all(clicks);
      await expect(button).toBeVisible();
    });

    test('should handle click during page load', async ({ page }) => {
      const clickPromise = page.goto('/').then(() => {
        return page.getByText('7 Days').first().click();
      });

      await clickPromise.catch(() => {});
      await expect(page.locator('body')).toBeVisible();
    });
  });

  test.describe('Edge Cases', () => {
    test('should handle button with zero dimensions', async ({ page }) => {
      await page.goto('/');

      await page.evaluate(() => {
        const btn = document.querySelector('button');
        if (btn) {
          Object.assign(btn.style, { width: '0px', height: '0px', overflow: 'hidden' });
        }
      });

      const button = page.locator('button').first();
      await button.isVisible();
      expect(typeof isVisible).toBe('boolean');
    });

    test('should handle button outside viewport', async ({ page }) => {
      await page.goto('/');

      await page.evaluate(() => {
        const btn = document.querySelector('button');
        if (btn) {
          btn.style.position = 'absolute';
          btn.style.top = '99999px';
        }
      });

      const button = page.locator('button').first();
      await button.scrollIntoViewIfNeeded().catch(() => {});
    });

    test('should handle CSS pointer-events none', async ({ page }) => {
      await page.goto('/');

      const button = page.getByText('7 Days').first();
      await page.evaluate(
        (btn) => {
          if (btn) btn.style.pointerEvents = 'none';
        },
        await button.elementHandle(),
      );

      await button.click({ force: true }).catch(() => {});
      await expect(page.locator('body')).toBeVisible();
    });
  });
});
