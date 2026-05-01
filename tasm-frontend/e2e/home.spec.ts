import { expect, test } from '@playwright/test';

test('navigates through the shell and renders mocked consumables data', async ({ page }) => {
  await page.route('http://localhost:8080/api/consumables', async (route) => {
    await route.fulfill({
      json: [
        {
          id: 1,
          name: 'Printer Paper A4',
          category: 'Office Supplies',
          location: 'Central Stockroom',
          currentStock: 24,
          reorderLevel: 30,
        },
        {
          id: 2,
          name: 'USB-C Cable',
          category: 'IT Peripherals',
          location: 'IT Cage',
          currentStock: 0,
          reorderLevel: 10,
        },
      ],
    });
  });

  await page.goto('/');

  await expect(page).toHaveTitle(/tasm-frontend/);
  await expect(page.getByText('Technopark AMS')).toBeVisible();

  await page.getByRole('link', { name: 'Consumables' }).click();
  await expect(page.getByRole('heading', { name: 'Consumables Tracker' })).toBeVisible();
  await expect(page.getByText('Printer Paper A4')).toBeVisible();
  await expect(page.getByText('Out of Stock')).toBeVisible();

  await page.getByRole('link', { name: 'Asset Registry' }).click();
  await expect(page).toHaveURL(/\/asset-registry$/);
  await expect(page.getByRole('heading', { name: 'Asset Registry' })).toBeVisible();
});
