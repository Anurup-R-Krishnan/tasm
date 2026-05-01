import { test, expect } from '@playwright/test';

test.describe('Dashboard and Main Application', () => {
  test('should load the dashboard and verify key metrics are present', async ({ page }) => {
    // Navigate to the app (assuming it runs on port 4173 for preview or 5173 for dev)
    await page.goto('/');

    // Check if the title is correct
    await expect(page).toHaveTitle(/TASM/i);

    // Check for KPI values
    await expect(page.locator('text=Total Assets')).toBeVisible();
    await expect(page.locator('text=Open Work Orders')).toBeVisible();
    await expect(page.locator('text=Procurement Pipeline')).toBeVisible();

    // Verify Dashboard modules
    await expect(page.locator('text=Global Operations Dashboard')).toBeVisible();
  });
});
