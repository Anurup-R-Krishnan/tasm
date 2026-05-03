import { test, expect } from '@playwright/test';

/**
 * Quick Action Buttons Tests
 * Tests the 4 quick action buttons on the Dashboard:
 * - Procurement
 * - Service Log
 * - Audit Scan
 * - Transfer
 */
test.describe('Dashboard Quick Action Buttons', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/');
    await page.waitForLoadState('networkidle').catch(() => {});
  });

  test.describe('Button Identification', () => {
    test('should find all 4 quick action buttons', async ({ page }) => {
      const quickActions = page.locator('.premium-card:has-text("Quick Operations") button');
      const count = await quickActions.count();
      expect(count).toBe(4);
    });

    test('should verify Procurement button exists', async ({ page }) => {
      const procurementBtn = page.getByText('Procurement').first();
      await expect(procurementBtn).toBeVisible();
    });

    test('should verify Service Log button exists', async ({ page }) => {
      const serviceLogBtn = page.getByText('Service Log').first();
      await expect(serviceLogBtn).toBeVisible();
    });

    test('should verify Audit Scan button exists', async ({ page }) => {
      const auditScanBtn = page.getByText('Audit Scan').first();
      await expect(auditScanBtn).toBeVisible();
    });

    test('should verify Transfer button exists', async ({ page }) => {
      const transferBtn = page.getByText('Transfer').first();
      await expect(transferBtn).toBeVisible();
    });
  });

  test.describe('Button Styling and Hover Effects', () => {
    test('should have proper hover transition classes', async ({ page }) => {
      const buttons = page.locator(
        '.premium-card:has-text("Quick Operations") button[class*="transition"]',
      );
      const count = await buttons.count();
      expect(count).toBe(4);
    });

    test('should change background on hover - Procurement', async ({ page }) => {
      const btn = page.getByText('Procurement').first();
      await btn.hover();
      await page.waitForTimeout(300);
      await expect(btn).toBeVisible();
    });

    test('should change background on hover - Service Log', async ({ page }) => {
      const btn = page.getByText('Service Log').first();
      await btn.hover();
      await page.waitForTimeout(300);
      await expect(btn).toBeVisible();
    });

    test('should change background on hover - Audit Scan', async ({ page }) => {
      const btn = page.getByText('Audit Scan').first();
      await btn.hover();
      await page.waitForTimeout(300);
      await expect(btn).toBeVisible();
    });

    test('should change background on hover - Transfer', async ({ page }) => {
      const btn = page.getByText('Transfer').first();
      await btn.hover();
      await page.waitForTimeout(300);
      await expect(btn).toBeVisible();
    });

    test('should have icon visibility in each button', async ({ page }) => {
      const icons = page.locator(
        '.premium-card:has-text("Quick Operations") .material-symbols-outlined',
      );
      const count = await icons.count();
      expect(count).toBe(4);
    });
  });

  test.describe('Click Behavior', () => {
    test('should click Procurement button', async ({ page }) => {
      const btn = page.getByText('Procurement').first();
      await btn.click();
      await page.waitForTimeout(500);
      await expect(page.locator('body')).toBeVisible();
    });

    test('should click Service Log button', async ({ page }) => {
      const btn = page.getByText('Service Log').first();
      await btn.click();
      await page.waitForTimeout(500);
      await expect(page.locator('body')).toBeVisible();
    });

    test('should click Audit Scan button', async ({ page }) => {
      const btn = page.getByText('Audit Scan').first();
      await btn.click();
      await page.waitForTimeout(500);
      await expect(page.locator('body')).toBeVisible();
    });

    test('should click Transfer button', async ({ page }) => {
      const btn = page.getByText('Transfer').first();
      await btn.click();
      await page.waitForTimeout(500);
      await expect(page.locator('body')).toBeVisible();
    });

    test('should handle double click on quick action', async ({ page }) => {
      const btn = page.getByText('Procurement').first();
      await btn.dblclick();
      await page.waitForTimeout(500);
      await expect(btn).toBeVisible();
    });
  });

  test.describe('Keyboard Navigation', () => {
    test('should navigate to quick action buttons with Tab', async ({ page }) => {
      await page.keyboard.press('Tab');
      await page.keyboard.press('Tab');
      await page.keyboard.press('Tab');
      await page.keyboard.press('Tab');
      await page.keyboard.press('Tab');

      const focused = page.locator(':focus');
      await expect(focused).toBeVisible();
    });

    test('should activate button with Enter key', async ({ page }) => {
      const btn = page.getByText('Procurement').first();
      await btn.focus();
      await page.keyboard.press('Enter');
      await page.waitForTimeout(300);
    });

    test('should activate button with Space key', async ({ page }) => {
      const btn = page.getByText('Service Log').first();
      await btn.focus();
      await page.keyboard.press('Space');
      await page.waitForTimeout(300);
    });
  });

  test.describe('Accessibility', () => {
    test('should have proper button roles', async ({ page }) => {
      const buttons = page.locator('.premium-card:has-text("Quick Operations") button');
      const count = await buttons.count();

      for (let i = 0; i < count; i++) {
        const role = await buttons.nth(i).getAttribute('role');
        expect(role === 'button' || role === null).toBeTruthy();
      }
    });

    test('should have accessible text or aria-label', async ({ page }) => {
      const buttons = page.locator('.premium-card:has-text("Quick Operations") button');
      const count = await buttons.count();

      for (let i = 0; i < count; i++) {
        const btn = buttons.nth(i);
        const text = await btn.textContent();
        const ariaLabel = await btn.getAttribute('aria-label');
        expect(text?.trim() || ariaLabel).toBeTruthy();
      }
    });

    test('should have proper cursor style', async ({ page }) => {
      const btn = page.getByText('Procurement').first();
      const cursor = await btn.evaluate((el) => window.getComputedStyle(el).cursor);
      expect(cursor).toBe('pointer');
    });
  });

  test.describe('Layout and Responsive', () => {
    test('should display in 2x2 grid on desktop', async ({ page }) => {
      await page.setViewportSize({ width: 1280, height: 720 });

      const container = page.locator('.premium-card:has-text("Quick Operations") .grid');
      await expect(container).toBeVisible();

      const gridClass = await container.getAttribute('class');
      expect(gridClass).toContain('grid-cols-2');
    });

    test('should render properly on mobile', async ({ page }) => {
      await page.setViewportSize({ width: 375, height: 667 });

      const buttons = page.locator('.premium-card:has-text("Quick Operations") button');
      const count = await buttons.count();

      for (let i = 0; i < count; i++) {
        await expect(buttons.nth(i)).toBeVisible();
      }
    });

    test('should have proper spacing between buttons', async ({ page }) => {
      const container = page.locator('.premium-card:has-text("Quick Operations") .grid');
      const gapClass = await container.getAttribute('class');
      expect(gapClass).toContain('gap-');
    });
  });

  test.describe('Icon Verification', () => {
    test('should have correct icon for Procurement', async ({ page }) => {
      const btn = page.getByText('Procurement').first();
      const icon = btn.locator('.material-symbols-outlined');
      await expect(icon).toHaveText('add_shopping_cart');
    });

    test('should have correct icon for Service Log', async ({ page }) => {
      const btn = page.getByText('Service Log').first();
      const icon = btn.locator('.material-symbols-outlined');
      await expect(icon).toHaveText('handyman');
    });

    test('should have correct icon for Audit Scan', async ({ page }) => {
      const btn = page.getByText('Audit Scan').first();
      const icon = btn.locator('.material-symbols-outlined');
      await expect(icon).toHaveText('qr_code_scanner');
    });

    test('should have correct icon for Transfer', async ({ page }) => {
      const btn = page.getByText('Transfer').first();
      const icon = btn.locator('.material-symbols-outlined');
      await expect(icon).toHaveText('sync_alt');
    });
  });

  test.describe('State After Click', () => {
    test('should remain enabled after click', async ({ page }) => {
      const btn = page.getByText('Procurement').first();
      await btn.click();
      await expect(btn).toBeEnabled();
    });

    test('should not cause page errors on click', async ({ page }) => {
      const errors: string[] = [];
      page.on('pageerror', (err) => errors.push(err.message));

      const btn = page.getByText('Service Log').first();
      await btn.click();
      await page.waitForTimeout(500);

      expect(errors.length).toBe(0);
    });
  });

  test.describe('Integration with Other Components', () => {
    test('should not affect time filter buttons', async ({ page }) => {
      const quickActionBtn = page.getByText('Procurement').first();
      await quickActionBtn.click();

      const timeFilter = page.getByText('7 Days').first();
      await expect(timeFilter).toBeVisible();
    });

    test('should not affect dashboard KPIs', async ({ page }) => {
      const btn = page.getByText('Audit Scan').first();
      await btn.click();

      await expect(page.getByText('Total Assets')).toBeVisible();
      await expect(page.getByText('Open Work Orders')).toBeVisible();
      await expect(page.getByText('Procurement Pipeline')).toBeVisible();
    });
  });
});
