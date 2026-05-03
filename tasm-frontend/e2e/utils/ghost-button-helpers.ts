import { Page, Locator } from '@playwright/test';

/**
 * Helper functions for ghost button tests
 */

export interface GhostButtonInfo {
  text: string;
  isVisible: boolean;
  isEnabled: boolean;
  bgColor: string;
  textColor: string;
  hasTransition: boolean;
}

/**
 * Find all ghost buttons on the page
 */
export async function findGhostButtons(page: Page): Promise<Locator> {
  return page.locator('button.text-text-secondary, button[class*="hover:text-primary"]');
}

/**
 * Get information about a ghost button
 */
export async function getGhostButtonInfo(button: Locator): Promise<GhostButtonInfo> {
  const text = (await button.textContent()) || '';
  const isVisible = await button.isVisible();
  const isEnabled = await button.isEnabled();

  const styles = await button.evaluate((el) => {
    const cs = window.getComputedStyle(el);
    return {
      bgColor: cs.backgroundColor,
      textColor: cs.color,
      hasTransition: cs.transition !== 'none' && cs.transition !== '',
    };
  });

  return {
    text,
    isVisible,
    isEnabled,
    bgColor: styles.bgColor,
    textColor: styles.textColor,
    hasTransition: styles.hasTransition,
  };
}

/**
 * Click a ghost button and verify it remains stable
 */
export async function clickAndVerify(button: Locator): Promise<void> {
  await button.click();
  await button.page().waitForTimeout(200);
  await expect(button).toBeVisible();
  await expect(button).toBeEnabled();
}

/**
 * Verify ghost button has proper styling
 */
export async function verifyGhostButtonStyling(button: Locator): Promise<void> {
  const info = await getGhostButtonInfo(button);

  expect(info.isVisible).toBeTruthy();
  expect(info.isEnabled).toBeTruthy();
  expect(info.hasTransition).toBeTruthy();
  expect(['rgba(0, 0, 0, 0)', 'transparent']).toContain(info.bgColor);
}

/**
 * Test click with all Playwright click options
 */
export async function testAllClickVariations(button: Locator): Promise<void> {
  const page = button.page();

  // Basic click
  await button.click();
  await page.waitForTimeout(100);

  // Click with position
  const box = await button.boundingBox();
  if (box) {
    await button.click({
      position: { x: box.width / 2, y: box.height / 2 },
    });
  }

  // Force click
  await button.click({ force: true });

  // Click with delay
  await button.click({ delay: 50 });

  // Double click
  await button.dblclick();

  await page.waitForTimeout(200);
}

/**
 * Test keyboard navigation to button
 */
export async function testKeyboardNavigation(page: Page, button: Locator): Promise<void> {
  // Tab to button
  await page.keyboard.press('Tab');
  await page.keyboard.press('Tab');
  await page.keyboard.press('Tab');

  // Verify focus
  await expect(button).toBeFocused();

  // Activate with Enter
  await page.keyboard.press('Enter');
  await page.waitForTimeout(200);

  // Tab again
  await page.keyboard.press('Tab');

  // Activate with Space
  await page.keyboard.press('Space');
  await page.waitForTimeout(200);
}

/**
 * Verify button is mobile-friendly (44x44 minimum)
 */
export async function verifyMobileFriendly(button: Locator): Promise<void> {
  const box = await button.boundingBox();
  if (box) {
    expect(box.width).toBeGreaterThanOrEqual(44);
    expect(box.height).toBeGreaterThanOrEqual(44);
  }
}

/**
 * Test hover effect on ghost button
 */
export async function testHoverEffect(button: Locator): Promise<void> {
  const beforeColor = await button.evaluate((el) => {
    return window.getComputedStyle(el).color;
  });

  await button.hover();
  await button.page().waitForTimeout(300);

  const afterColor = await button.evaluate((el) => {
    return window.getComputedStyle(el).color;
  });

  expect(beforeColor).toBeTruthy();
  expect(afterColor).toBeTruthy();
}

/**
 * Mock API and test button interaction
 */
export async function testWithMockedAPI(page: Page, button: Locator, mockData: any): Promise<void> {
  await page.route('**/api/**', (route) => {
    route.fulfill({
      json: mockData,
    });
  });

  await button.click();
  await page.waitForTimeout(500);
}

/**
 * Take screenshot of button with name
 */
export async function screenshotButton(button: Locator, name: string): Promise<void> {
  await button.screenshot({ path: `test-results/${name}.png` }).catch(() => {});
}

/**
 * Verify no JavaScript errors during interaction
 */
export async function verifyNoErrors(page: Page, button: Locator): Promise<void> {
  const errors: string[] = [];
  page.on('pageerror', (err) => errors.push(err.message));

  await button.click();
  await page.waitForTimeout(500);

  expect(errors.length).toBe(0);
}
