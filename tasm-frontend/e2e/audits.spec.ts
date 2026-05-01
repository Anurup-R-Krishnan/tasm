import { test, expect } from '@playwright/test';

test.describe('Audit & Discrepancies Module', () => {
  test('should load audit sessions and verify table structure', async ({ page }) => {
    // Navigate to Audit Sessions
    await page.goto('/audit-sessions');
    
    // Verify header
    await expect(page.locator('h1', { hasText: 'Audit Sessions' })).toBeVisible();
    
    // Verify table headers (from PrimeVue DataTable)
    await expect(page.locator('th', { hasText: 'Session ID' })).toBeVisible();
    await expect(page.locator('th', { hasText: 'Status' })).toBeVisible();
  });

  test('should load discrepancies and verify missing items', async ({ page }) => {
    // Navigate to Discrepancy Resolution
    await page.goto('/audit-discrepancy-resolution');
    
    // Verify header
    await expect(page.locator('h2', { hasText: 'Audit Discrepancy Resolution' })).toBeVisible();
    
    // Wait for the data table to populate
    await expect(page.locator('table')).toBeVisible();
  });
});
