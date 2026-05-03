# Ghost Button Playwright Tests - Summary

## Test Files Created

### 1. `ghost-buttons.spec.ts` - Core Tests
Basic ghost button identification and click behavior:
- Ghost button identification by styling
- Basic click, double-click, keyboard clicks
- Hover behavior and cursor styles
- State verification after clicks
- Accessibility (tab navigation)
- Multiple button interactions
- Edge cases (out of viewport, modifiers)

### 2. `ghost-buttons-advanced.spec.ts` - Advanced Scenarios
Complex interaction patterns:
- Button state management and toggling
- Click timing, delays, race conditions
- Network/API interaction interception
- Viewport responsiveness (mobile/desktop)
- Button groups and layout verification
- Event listener verification
- Touch/pointer events
- Cross-page navigation tests
- Performance measurement
- ARIA/semantic HTML verification
- CSS/styling verification

### 3. `ghost-buttons-comprehensive.spec.ts` - Edge Cases
Thorough edge case coverage:
- Disabled/loading states
- Buttons with icons and mixed content
- Route changes on click
- Dynamic content handling
- Focus and blur events
- Nested elements
- Animation and transitions
- Keyboard shortcuts
- Modal/overlay context
- Data attributes
- Scroll behavior
- Right click and context menu
- State persistence
- Form submission context
- AJAX calls
- Memory leak prevention
- Boundary testing
- Drag and drop simulation
- LocalStorage/SessionStorage
- Console warnings/errors
- Network offline simulation
- CSS pseudo-classes
- Event propagation
- Snapshot testing

### 4. `ghost-buttons-mcp.spec.ts` - MCP Integration
Model Context Protocol specific tests:
- MCP click command execution
- Locator strategies
- Screenshot and visual tests
- JavaScript evaluation and injection
- Network interception
- Wait and timing commands
- Keyboard and input simulation
- Dialog and popup handling
- Multi-tab/window handling
- Accessibility tree
- Storage and state
- Geolocation and permissions
- Tracing and debugging
- PDF and print
- Video and media
- Drag and drop
- File download
- Service worker and cache

### 5. `ghost-buttons-mobile.spec.ts` - Mobile Specific
Mobile and touch interactions:
- Tap on ghost buttons
- Long press
- Swipe gestures
- Mobile viewport tests
- Orientation changes
- Mobile gesture events

### 6. `ghost-buttons-performance.spec.ts` - Performance
Performance benchmarking:
- Click response time
- Memory and resource usage
- Network performance
- Rendering performance
- Layout thrashing

### 7. `ghost-buttons-accessibility.spec.ts` - A11y
WCAG compliance and accessibility:
- Color contrast verification
- Keyboard accessibility
- Screen reader support
- Focus management
- ARIA attributes

### 8. `ghost-buttons-error-handling.spec.ts` - Error Cases
Robust error handling:
- JavaScript errors
- Network errors (API failures, timeouts)
- DOM mutations
- Browser compatibility
- Memory cleanup
- Concurrent operations
- Edge cases (zero dimensions, outside viewport)

### 9. `ghost-buttons-visual.spec.ts` - Visual Regression
Visual testing:
- Screenshot comparisons
- CSS property verification
- Layout verification
- Responsive visual tests
- State change visualization
- Color contrast verification

### 10. `ghost-buttons-integration.spec.ts` - Integration
Integration with other components:
- Sidebar navigation
- Search inputs
- Dropdowns
- Forms and validation
- Modals and dialogs
- Data tables
- Routing
- API data
- LocalStorage/SessionStorage
- Cookies
- Events and analytics
- Animations

## Running the Tests

```bash
cd tasm-frontend

# Run all ghost button tests
npx playwright test ghost-buttons

# Run specific test file
npx playwright test ghost-buttons.spec.ts
npx playwright test ghost-buttons-mobile.spec.ts

# Run with UI mode
npx playwright test --ui

# Run with debug mode
npx playwright test --debug ghost-buttons.spec.ts

# Run with specific browser
npx playwright test --project=chromium ghost-buttons.spec.ts
```

## Test Coverage

- **Button Identification**: 5+ patterns
- **Click Behaviors**: 15+ variations
- **Keyboard Interactions**: 8+ scenarios
- **Mobile/Touch**: 10+ gestures
- **Accessibility**: WCAG 2.1 AA compliance
- **Performance**: Response time, memory, network
- **Error Handling**: 20+ edge cases
- **Visual Regression**: Screenshot comparisons
- **Integration**: All major UI components

## Key Ghost Button Selectors Used

```typescript
// Primary selectors
page.getByText('7 Days')
page.getByText('30 Days')
page.locator('button.text-text-secondary')
page.locator('button[class*="hover:text-primary"]')
page.locator('.rounded-xl.border button')

// MCP-specific
page.getByRole('button', { name: /7 Days|30 Days/ })
```

## Notes

- Ghost buttons in this project use `text-text-secondary` class with hover effects
- Time filter buttons ("7 Days", "30 Days") are the primary ghost button examples
- Tests mock API calls where needed to avoid external dependencies
- Visual tests require baseline screenshots to be generated first
