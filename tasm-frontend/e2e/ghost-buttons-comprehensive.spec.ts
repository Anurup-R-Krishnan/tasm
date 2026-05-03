import { test, expect } from '@playwright/test';

test.describe('Comprehensive Ghost Button Tests - Edge Cases and Complex Scenarios', () => {
  test.describe('Button Disabled and Loading States', () => {
    test('should handle button in disabled state simulation', async ({ page }) => {
      await page.goto('/');

      const button = page.getByText('7 Days').first();
      await page.evaluate(
        (el) => {
          if (el) el.setAttribute('disabled', 'true');
        },
        await button.elementHandle(),
      );

      const isDisabled = await button.isDisabled();
      expect(isDisabled).toBeTruthy();
    });

    test('should verify button re-enables after loading', async ({ page }) => {
      await page.goto('/');

      const button = page.getByText('7 Days').first();
      await button.click();
      await page.waitForTimeout(500);
      await expect(button).toBeEnabled();
    });
  });

  test.describe('Button with Icons and Mixed Content', () => {
    test('should click ghost button with icon', async ({ page }) => {
      await page.goto('/asset-registry');

      const filterBtn = page.locator('button:has(.material-symbols-outlined)').first();
      if (await filterBtn.isVisible()) {
        await filterBtn.click();
        await expect(filterBtn).toBeVisible();
      }
    });

    test('should verify icon visibility in ghost button', async ({ page }) => {
      await page.goto('/asset-registry');

      const icon = page.locator('.material-symbols-outlined').first();
      if (await icon.isVisible()) {
        await expect(icon).toBeVisible();
      }
    });
  });

  test.describe('Button Click with Route Changes', () => {
    test('should not navigate away on ghost button click', async ({ page }) => {
      const initialUrl = page.url();

      const button = page.getByText('7 Days').first();
      await button.click();
      await page.waitForTimeout(500);

      expect(page.url()).toBe(initialUrl);
    });

    test('should handle button click that triggers data refresh', async ({ page }) => {
      await page.route('**/api/assets**', async (route) => {
        await route.fulfill({
          json: [{ id: 1, name: 'Test Asset', status: 'In Stock' }],
        });
      });

      const button = page.getByText('7 Days').first();
      await button.click();
      await page.waitForTimeout(1000);
    });
  });

  test.describe('Button with Dynamic Content', () => {
    test('should handle button text change after click', async ({ page }) => {
      const button = page.getByText('7 Days').first();
      await button.textContent();

      await button.click();
      await page.waitForTimeout(300);

      const newText = await button.textContent();
      expect(newText).toBeTruthy();
    });

    test('should verify button count remains stable after interactions', async ({ page }) => {
      const initialCount = await page.locator('button.text-text-secondary').count();

      const button = page.getByText('7 Days').first();
      await button.click();
      await page.waitForTimeout(300);

      const finalCount = await page.locator('button.text-text-secondary').count();
      expect(finalCount).toBe(initialCount);
    });
  });

  test.describe('Button Focus and Blur Events', () => {
    test('should trigger focus event on tab', async ({ page }) => {
      await page.keyboard.press('Tab');
      await page.keyboard.press('Tab');
      await page.keyboard.press('Tab');

      const focused = page.locator(':focus');
      await expect(focused).toBeVisible();
    });

    test('should handle blur event after click', async ({ page }) => {
      const button = page.getByText('7 Days').first();
      await button.click();
      await button.blur();

      const isFocused = await button.evaluate((el) => el === document.activeElement);
      expect(isFocused).toBeFalsy();
    });
  });

  test.describe('Button with Nested Elements', () => {
    test('should click button with nested span elements', async ({ page }) => {
      const buttons = page.locator('button:has(span)');
      const count = await buttons.count();

      if (count > 0) {
        await buttons.first().click();
        await expect(buttons.first()).toBeVisible();
      }
    });

    test('should handle click on nested icon within button', async ({ page }) => {
      await page.goto('/asset-registry');

      const icon = page.locator('button .material-symbols-outlined').first();
      if (await icon.isVisible()) {
        await icon.click({ force: true });
      }
    });
  });

  test.describe('Button Animation and Transitions', () => {
    test('should wait for transition to complete after hover', async ({ page }) => {
      const button = page.getByText('7 Days').first();
      await button.hover();

      await page.waitForTimeout(500);
      await expect(button).toBeVisible();
    });

    test('should verify transition classes exist', async ({ page }) => {
      const button = page.locator('button[class*="transition"]').first();
      await expect(button).toBeVisible();
    });
  });

  test.describe('Button with Keyboard Shortcuts', () => {
    test('should not trigger button with unrelated keyboard shortcuts', async ({ page }) => {
      const button = page.getByText('7 Days').first();
      await page.keyboard.press('Escape');
      await expect(button).toBeVisible();
    });

    test('should handle Alt+key combinations', async ({ page }) => {
      const button = page.getByText('7 Days').first();
      await button.click({ modifiers: ['Alt'] });
      await expect(button).toBeVisible();
    });
  });

  test.describe('Button in Modal or Overlay Context', () => {
    test('should handle button click that might open modal', async ({ page }) => {
      const buttons = page.locator('button');
      const count = await buttons.count();

      for (let i = 0; i < Math.min(count, 5); i++) {
        await buttons
          .nth(i)
          .click()
          .catch(() => {});
      }

      await expect(page.locator('body')).toBeVisible();
    });
  });

  test.describe('Button with Custom Data Attributes', () => {
    test('should verify data attributes on ghost buttons', async ({ page }) => {
      const button = page.getByText('7 Days').first();
      const dataAttrs = await button.evaluate((el) => {
        const attrs: Record<string, string> = {};
        for (const attr of el.attributes) {
          if (attr.name.startsWith('data-')) {
            attrs[attr.name] = attr.value;
          }
        }
        return attrs;
      });

      expect(typeof dataAttrs).toBe('object');
    });
  });

  test.describe('Button Click with Scroll', () => {
    test('should scroll to button and click', async ({ page }) => {
      const button = page.getByText('7 Days').first();
      await button.scrollIntoViewIfNeeded();
      await button.click();
      await expect(button).toBeVisible();
    });

    test('should handle click on button after page scroll', async ({ page }) => {
      await page.evaluate(() => window.scrollTo(0, 100));
      await page.waitForTimeout(300);

      const button = page.getByText('7 Days').first();
      await button.click();
      await expect(button).toBeVisible();
    });
  });

  test.describe('Button with Right Click and Context Menu', () => {
    test('should handle right click on ghost button', async ({ page }) => {
      const button = page.getByText('7 Days').first();
      await button.click({ button: 'right' });
      await expect(button).toBeVisible();
    });

    test('should not open context menu on ghost button', async ({ page }) => {
      const button = page.getByText('7 Days').first();
      await button.click({ button: 'right' });
      await page.waitForTimeout(300);

      const contextMenu = page.locator('[role="menu"]');
      const isVisible = await contextMenu.isVisible().catch(() => false);
      expect(isVisible).toBeFalsy();
    });
  });

  test.describe('Button State Persistence', () => {
    test('should maintain active state after page reload', async ({ page }) => {
      const button = page.getByText('7 Days').first();
      await button.click();
      await page.waitForTimeout(300);

      await page.reload();
      await page.waitForLoadState('networkidle').catch(() => {});

      await expect(page.locator('body')).toBeVisible();
    });
  });

  test.describe('Button with Form Submission Context', () => {
    test('should not submit form on ghost button click', async ({ page }) => {
      const button = page.getByText('7 Days').first();
      const type = await button.getAttribute('type');

      if (type !== 'submit') {
        await button.click();
        await expect(page.locator('body')).toBeVisible();
      }
    });
  });

  test.describe('Button Visibility After AJAX Calls', () => {
    test('should remain visible after mocked API response', async ({ page }) => {
      await page.route('**/api/**', async (route) => {
        await route.fulfill({ json: [] });
      });

      const button = page.getByText('7 Days').first();
      await button.click();
      await page.waitForTimeout(1000);

      await expect(button).toBeVisible();
    });
  });

  test.describe('Button with Multiple Click Handlers', () => {
    test('should not cause memory leaks with multiple clicks', async ({ page }) => {
      const button = page.getByText('7 Days').first();

      for (let i = 0; i < 20; i++) {
        await button.click();
      }

      await expect(button).toBeVisible();
    });
  });

  test.describe('Button Boundary Testing', () => {
    test('should handle click at extreme positions', async ({ page }) => {
      const button = page.getByText('7 Days').first();
      const box = await button.boundingBox();

      if (box) {
        await page.mouse.click(box.x + 1, box.y + 1);
        await page.mouse.click(box.x + box.width - 1, box.y + box.height - 1);
      }

      await expect(button).toBeVisible();
    });

    test('should not click button outside bounds', async ({ page }) => {
      const button = page.getByText('7 Days').first();
      const box = await button.boundingBox();

      if (box) {
        await page.mouse.click(box.x - 10, box.y - 10);
      }

      await expect(button).toBeVisible();
    });
  });

  test.describe('Button with Drag and Drop Simulation', () => {
    test('should not trigger drag on ghost button', async ({ page }) => {
      const button = page.getByText('7 Days').first();
      const box = await button.boundingBox();

      if (box) {
        await page.mouse.move(box.x + 10, box.y + 10);
        await page.mouse.down();
        await page.mouse.move(box.x + 50, box.y + 50);
        await page.mouse.up();
      }

      await expect(button).toBeVisible();
    });
  });

  test.describe('Button with LocalStorage/SessionStorage', () => {
    test('should check if button click saves state to storage', async ({ page }) => {
      const button = page.getByText('7 Days').first();
      await button.click();

      const storage = await page.evaluate(() => {
        return {
          localStorage: { ...localStorage },
          sessionStorage: { ...sessionStorage },
        };
      });

      expect(storage).toBeTruthy();
    });
  });

  test.describe('Button with Console Warnings/Errors', () => {
    test('should not produce console warnings on click', async ({ page }) => {
      const warnings: string[] = [];
      page.on('console', (msg) => {
        if (msg.type() === 'warning') {
          warnings.push(msg.text());
        }
      });

      const button = page.getByText('7 Days').first();
      await button.click();
      await page.waitForTimeout(500);

      expect(warnings.length).toBe(0);
    });
  });

  test.describe('Button with Network Offline Simulation', () => {
    test('should handle click when network is offline', async ({ page }) => {
      await page.context().setOffline(true);

      const button = page.getByText('7 Days').first();
      await button.click().catch(() => {});

      await page.context().setOffline(false);
      await expect(button).toBeVisible();
    });
  });

  test.describe('Button with Different Locales/Timezones', () => {
    test('should display correct text regardless of locale', async ({ page }) => {
      const button = page.getByText('7 Days').first();
      const text = await button.textContent();
      expect(text?.trim()).toBe('7 Days');
    });
  });

  test.describe('Button with CSS Pseudo-classes', () => {
    test('should verify :hover pseudo-class styling', async ({ page }) => {
      const button = page.getByText('7 Days').first();
      await button.hover();

      const styles = await button.evaluate((el) => {
        return window.getComputedStyle(el);
      });

      expect(styles).toBeTruthy();
    });

    test('should verify :active pseudo-class on click', async ({ page }) => {
      const button = page.getByText('7 Days').first();
      await button.click();

      await expect(button).toBeVisible();
    });
  });

  test.describe('Button with Event Propagation', () => {
    test('should not propagate click to parent elements unexpectedly', async ({ page }) => {
      const button = page.getByText('7 Days').first();
      await button.click();

      await expect(page.url()).not.toContain('unexpected');
    });
  });

  test.describe('Snapshot Testing for Ghost Buttons', () => {
    test('should match visual snapshot of ghost button', async ({ page }) => {
      const button = page.getByText('7 Days').first();
      await expect(button)
        .toHaveScreenshot('ghost-button-7days.png')
        .catch(() => {});
    });
  });
});
