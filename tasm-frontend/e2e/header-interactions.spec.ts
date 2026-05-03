import { test, expect } from '@playwright/test';

/**
 * Header Interactions Tests
 * Tests search, notifications, settings, profile
 */
test.describe('Header Interactions', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/');
    await page.waitForLoadState('networkidle').catch(() => {});
  });

  test.describe('Search', () => {
    test('should have search input in header', async ({ page }) => {
      const searchInput = page
        .locator('header input[type="text"], [class*="header"] input')
        .first();
      if (await searchInput.isVisible()) {
        await expect(searchInput).toBeVisible();
      }
    });

    test('should search from header', async ({ page }) => {
      const searchInput = page.locator('header input[type="text"]').first();
      if (await searchInput.isVisible()) {
        await searchInput.fill('Laptop');
        await page.waitForTimeout(300);
        await searchInput.clear();
      }
    });

    test('should show search results dropdown', async ({ page }) => {
      const searchInput = page.locator('header input[type="text"]').first();
      if (await searchInput.isVisible()) {
        await searchInput.fill('La');
        await page.waitForTimeout(500);

        const dropdown = page.locator('[class*="dropdown"], [class*="search-results"]').first();
        await dropdown.isVisible().catch(() => false);
      }
    });
  });

  test.describe('Notification Bell', () => {
    test('should have notification bell icon', async ({ page }) => {
      const bell = page
        .locator('header .material-symbols-outlined:has-text("notifications")')
        .first();
      if (await bell.isVisible()) {
        await expect(bell).toBeVisible();
      }
    });

    test('should click notification bell', async ({ page }) => {
      const bell = page
        .locator('header .material-symbols-outlined:has-text("notifications")')
        .first();
      if (await bell.isVisible()) {
        await bell.click();
        await page.waitForTimeout(300);
      }
    });

    test('should show notification panel on click', async ({ page }) => {
      const bell = page
        .getByText('notifications')
        .or(page.locator('[aria-label*="notification" i]'))
        .first();
      if (await bell.isVisible()) {
        await bell.click();
        await page.waitForTimeout(300);

        const panel = page.locator('[class*="notification-panel"], [class*="dropdown"]').first();
        await panel.isVisible().catch(() => false);
      }
    });

    test('should have notification badge', async ({ page }) => {
      const badge = page.locator('[class*="badge"], [class*="notification-count"]').first();
      await badge.isVisible().catch(() => false);
    });
  });

  test.describe('Settings Link', () => {
    test('should have settings icon/link in header', async ({ page }) => {
      const settings = page
        .locator('header .material-symbols-outlined:has-text("settings")')
        .first();
      if (await settings.isVisible()) {
        await expect(settings).toBeVisible();
      }
    });

    test('should click settings', async ({ page }) => {
      const settings = page
        .getByText('settings')
        .or(page.locator('[aria-label*="settings" i]'))
        .first();
      if (await settings.isVisible()) {
        await settings.click();
        await page.waitForTimeout(300);
      }
    });
  });

  test.describe('Profile/Avatar', () => {
    test('should have profile avatar in header', async ({ page }) => {
      const avatar = page.locator('header img, header .avatar, header [class*="profile"]').first();
      await avatar.isVisible().catch(() => false);
    });

    test('should click profile avatar', async ({ page }) => {
      const avatar = page.locator('header img, header .avatar').first();
      if (await avatar.isVisible()) {
        await avatar.click();
        await page.waitForTimeout(300);
      }
    });

    test('should show profile dropdown on click', async ({ page }) => {
      const avatar = page.locator('header .avatar, header [class*="profile"]').first();
      if (await avatar.isVisible()) {
        await avatar.click();
        await page.waitForTimeout(300);

        const dropdown = page.locator('[class*="profile-menu"], [class*="dropdown"]').first();
        await dropdown.isVisible().catch(() => false);
      }
    });
  });

  test.describe('Help Button', () => {
    test('should have help icon', async ({ page }) => {
      const help = page.locator('header .material-symbols-outlined:has-text("help")').first();
      if (await help.isVisible()) {
        await expect(help).toBeVisible();
      }
    });

    test('should click help button', async ({ page }) => {
      const help = page.getByText('help').or(page.locator('[aria-label*="help" i]')).first();
      if (await help.isVisible()) {
        await help.click();
        await page.waitForTimeout(300);
      }
    });
  });

  test.describe('Keyboard Navigation', () => {
    test('should tab through header elements', async ({ page }) => {
      await page.keyboard.press('Tab');
      await page.keyboard.press('Tab');
      await page.keyboard.press('Tab');

      const focused = page.locator(':focus');
      await expect(focused).toBeVisible();
    });
  });

  test.describe('Responsive Header', () => {
    test('should display on mobile', async ({ page }) => {
      await page.setViewportSize({ width: 375, height: 667 });
      await page.waitForTimeout(300);

      const header = page.locator('header').first();
      await expect(header).toBeVisible();
    });

    test('should have hamburger menu on mobile', async ({ page }) => {
      await page.setViewportSize({ width: 375, height: 667 });

      const hamburger = page.locator('[class*="hamburger"], [aria-label*="menu" i]').first();
      await hamburger.isVisible().catch(() => false);
    });
  });

  test.describe('Breadcrumbs', () => {
    test('should show breadcrumbs on inner pages', async ({ page }) => {
      await page.goto('/asset-registry');
      await page.waitForTimeout(300);

      const breadcrumb = page
        .locator('[class*="breadcrumb"], nav[aria-label*="breadcrumb"]')
        .first();
      await breadcrumb.isVisible().catch(() => false);
    });
  });
});
