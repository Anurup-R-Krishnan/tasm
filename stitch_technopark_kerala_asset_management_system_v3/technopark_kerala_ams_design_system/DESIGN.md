---
name: Technopark Kerala AMS Design System
colors:
  surface: '#FFFFFF'
  surface-dim: '#e8d6cf'
  surface-bright: '#fff8f6'
  surface-container-lowest: '#ffffff'
  surface-container-low: '#fff1eb'
  surface-container: '#fdeae3'
  surface-container-high: '#f7e4dd'
  surface-container-highest: '#f1dfd8'
  on-surface: '#231a15'
  on-surface-variant: '#55433a'
  inverse-surface: '#392e2a'
  inverse-on-surface: '#ffede6'
  outline: '#887269'
  outline-variant: '#dcc1b6'
  surface-tint: '#9a4614'
  primary: '#712c00'
  on-primary: '#ffffff'
  primary-container: '#92400e'
  on-primary-container: '#ffc2a5'
  inverse-primary: '#ffb693'
  secondary: '#625d5b'
  on-secondary: '#ffffff'
  secondary-container: '#e9e1dd'
  on-secondary-container: '#686361'
  tertiary: '#00446d'
  on-tertiary: '#ffffff'
  tertiary-container: '#005c92'
  on-tertiary-container: '#a8d3ff'
  error: '#ba1a1a'
  on-error: '#ffffff'
  error-container: '#ffdad6'
  on-error-container: '#93000a'
  primary-fixed: '#ffdbcb'
  primary-fixed-dim: '#ffb693'
  on-primary-fixed: '#341000'
  on-primary-fixed-variant: '#7a3000'
  secondary-fixed: '#e9e1dd'
  secondary-fixed-dim: '#ccc5c2'
  on-secondary-fixed: '#1e1b19'
  on-secondary-fixed-variant: '#4a4643'
  tertiary-fixed: '#cee5ff'
  tertiary-fixed-dim: '#98cbff'
  on-tertiary-fixed: '#001d33'
  on-tertiary-fixed-variant: '#004a77'
  background: '#fff8f6'
  on-background: '#231a15'
  surface-variant: '#f1dfd8'
  canvas: '#FAF7F2'
  surface-subtle: '#FAFAF8'
  text-primary: '#1C1917'
  text-secondary: '#78716C'
  border-default: '#EEEBE4'
  accent-subtle: '#FEF3C7'
  metric-amber: '#FEF3C7'
  metric-lavender: '#EDE9FE'
  metric-sage: '#DCFCE7'
  status-in-stock: '#166534'
  status-critical: '#991B1B'
  status-checked-out: '#1E40AF'
typography:
  h1:
    fontFamily: Inter
    fontSize: 22px
    fontWeight: '600'
    lineHeight: '1.2'
  h2:
    fontFamily: Inter
    fontSize: 16px
    fontWeight: '600'
    lineHeight: '1.4'
  h3:
    fontFamily: Inter
    fontSize: 14px
    fontWeight: '600'
    lineHeight: '1.4'
  body:
    fontFamily: Inter
    fontSize: 14px
    fontWeight: '400'
    lineHeight: '1.5'
  metadata:
    fontFamily: Inter
    fontSize: 12px
    fontWeight: '400'
    lineHeight: '1.4'
  table-header:
    fontFamily: Inter
    fontSize: 12px
    fontWeight: '500'
    letterSpacing: 0.06em
  mono-data:
    fontFamily: JetBrains Mono
    fontSize: 13px
    fontWeight: '400'
    lineHeight: '1'
  kpi-number:
    fontFamily: Inter
    fontSize: 28px
    fontWeight: '600'
    lineHeight: '1.1'
rounded:
  sm: 0.25rem
  DEFAULT: 0.5rem
  md: 0.75rem
  lg: 1rem
  xl: 1.5rem
  full: 9999px
spacing:
  base: 4px
  inline: 12px
  stack: 16px
  card-padding: 24px
  section-gap: 32px
  page-margin: 32px
---

# Technopark Kerala — Asset Management System
## Design System · v1.0

> **Status:** Living document. Update when tokens change.  
> **Stack:** React + Tailwind CSS v3 + Inter + JetBrains Mono + Recharts  
> **Last updated:** April 2026

---

## Table of Contents

1. [Design Philosophy](#1-design-philosophy)
2. [Color System](#2-color-system)
3. [Typography](#3-typography)
4. [Spacing & Layout](#4-spacing--layout)
5. [Elevation & Borders](#5-elevation--borders)
6. [Components](#6-components)
7. [Status Badges](#7-status-badges)
8. [Navigation Shell](#8-navigation-shell)
9. [Forms](#9-forms)
10. [Tables](#10-tables)
11. [Cards & Metric Cards](#11-cards--metric-cards)
12. [Drawers & Modals](#12-drawers--modals)
13. [Alert System](#13-alert-system)
14. [Empty & Error States](#14-empty--error-states)
15. [Motion & Transitions](#15-motion--transitions)
16. [Mobile (Audit Mode)](#16-mobile-audit-mode)
17. [Accessibility](#17-accessibility)
18. [Design Tokens (CSS Variables)](#18-design-tokens-css-variables)

---

## 1. Design Philosophy

**Warm. Precise. Trustworthy.**

This is a data management tool for a government-affiliated technology park. It is used every day by asset managers, IT teams, and auditors. It must be:

- **Immediately readable** — dense information without visual fatigue
- **Warm, not cold** — cream canvas and amber accents instead of clinical blue-white
- **Unambiguous** — every status, every action, every number must mean exactly one thing
- **Nothing decorative** — every pixel earns its place by helping someone do their job faster

**Reference aesthetic:** The jewellery dashboard reference image — warm cream background, colorful pastel metric cards (amber, lavender, pink, sage), rounded 16px card corners, clean Inter typography, right-side detail drawers.

**Not:** A legacy government portal. Not a startup SaaS. Not dark mode. Not glassmorphism.

---

## 2. Color System

### Canvas & Surface

| Token | Hex | Usage |
|---|---|---|
| `--color-canvas` | `#FAF7F2` | Page background (warm cream) |
| `--color-surface` | `#FFFFFF` | Cards, drawers, modals, table rows |
| `--color-surface-subtle` | `#FAFAF8` | Alternating table rows, input backgrounds |
| `--color-sidebar` | `#FFFFFF` | Sidebar background |

### Text

| Token | Hex | Usage |
|---|---|---|
| `--color-text-primary` | `#1C1917` | Headings, table cell content, labels |
| `--color-text-secondary` | `#78716C` | Metadata, timestamps, table headers, helper text |
| `--color-text-disabled` | `#A8A29E` | Disabled inputs, inactive nav items |
| `--color-text-inverse` | `#FFFFFF` | Text on dark backgrounds |
| `--color-text-link` | `#92400E` | Links, tag IDs on hover |

### Borders

| Token | Hex | Usage |
|---|---|---|
| `--color-border-default` | `#EEEBE4` | Cards, sidebar, dividers, inputs |
| `--color-border-strong` | `#D6D3CE` | Secondary buttons, focused input |
| `--color-border-accent` | `#92400E` | Active sidebar item, selected row, focus ring |

### Brand Accent (Amber / Warm)

| Token | Hex | Usage |
|---|---|---|
| `--color-accent` | `#92400E` | Primary interactive: active nav, links, focus rings, selected states |
| `--color-accent-hover` | `#78350F` | Hover state for amber interactive elements |
| `--color-accent-subtle` | `#FEF3C7` | Active sidebar bg, selected row bg, amber card fill |
| `--color-accent-border` | `#FCD34D` | Amber badge border |

### Metric Card Fills

Each fill has a matching text colour. Never mix fills and text from different ramps.

| Card Type | Background | Text | Icon | Usage |
|---|---|---|---|---|
| Amber | `#FEF3C7` | `#92400E` | `#D97706` | Total assets, primary counts |
| Lavender | `#EDE9FE` | `#5B21B6` | `#7C3AED` | Financial values, work orders |
| Peach / Red | `#FEE2E2` | `#991B1B` | `#DC2626` | Critical counts, overdue items |
| Soft Pink | `#FCE7F3` | `#9D174D` | `#DB2777` | Alerts, attention items |
| Sage / Green | `#DCFCE7` | `#166534` | `#16A34A` | Net value, healthy metrics |

### Status Semantic Colours

| Status | Background | Text | Border |
|---|---|---|---|
| In Stock | `#DCFCE7` | `#166534` | `#86EFAC` |
| Checked Out | `#DBEAFE` | `#1E40AF` | `#93C5FD` |
| Reserved | `#EDE9FE` | `#5B21B6` | `#C4B5F8` |
| Under Repair | `#FEF3C7` | `#92400E` | `#FCD34D` |
| Leased Out | `#FCE7F3` | `#9D174D` | `#F9A8D4` |
| Lost | `#FEE2E2` | `#991B1B` | `#FCA5A5` |
| Disposed | `#F3F4F6` | `#6B7280` | `#D1D5DB` |
| Donated | `#ECFDF5` | `#065F46` | `#6EE7B7` |
| Sold | `#FFF7ED` | `#9A3412` | `#FDBA74` |
| Pending Ack. | `#FFFBEB` | `#92400E` | `#FDE68A` |

### Severity / Alert Colours

| Level | Background | Left Border | Text |
|---|---|---|---|
| Critical | `#FEF2F2` | `#DC2626` | `#991B1B` |
| Warning | `#FFFBEB` | `#D97706` | `#92400E` |
| Info | `#EFF6FF` | `#3B82F6` | `#1E40AF` |
| Success | `#F0FDF4` | `#16A34A` | `#166534` |

### Interactive Buttons

| Variant | Background | Text | Border | Hover BG |
|---|---|---|---|---|
| Primary | `#1C1917` | `#FFFFFF` | — | `#292524` |
| Secondary | `#FFFFFF` | `#1C1917` | `#D6D3CE` | `#FAF7F2` |
| Ghost | `transparent` | `#78716C` | — | `#FAF7F2` |
| Danger | `#FFFFFF` | `#DC2626` | `#FECACA` | `#FEF2F2` |
| Accent | `#FEF3C7` | `#92400E` | `#FCD34D` | `#FDE68A` |

> **Danger actions always require typed confirmation.** Never a simple single click.

---

## 3. Typography

### Font Stack

```css
--font-body:    'Inter', -apple-system, BlinkMacSystemFont, sans-serif;
--font-mono:    'JetBrains Mono', 'Fira Code', monospace;
```

Inter is used for everything except asset tag IDs, serial numbers, and depreciation formulas — those use JetBrains Mono for scannability.

### Type Scale

| Role | Size | Weight | Color | Notes |
|---|---|---|---|---|
| Page Title (H1) | 22px | 600 | `--color-text-primary` | One per page |
| Section Title (H2) | 16px | 600 | `--color-text-primary` | Module headers |
| Sub-section (H3) | 14px | 600 | `--color-text-primary` | Form sections, drawer sections |
| Table Header | 12px | 500 | `--color-text-secondary` | UPPERCASE · 0.06em tracking |
| Body / Table Cell | 14px | 400 | `--color-text-primary` | Default readable text |
| Small / Metadata | 12px | 400 | `--color-text-secondary` | Timestamps, helper text |
| Badge / Label | 11px | 500 | varies | Status pills |
| Tag ID / Serial | 13px | 400 | `--color-text-primary` | JetBrains Mono |
| KPI Number | 28px | 600 | varies per card | Dashboard metric cards |
| KPI Label | 12px | 400 | muted per card ramp | Above the number |
| KPI Sub-label | 12px | 400 | muted per card ramp | Below the number |
| Button | 14px | 500 | varies | All button variants |

### Currency Formatting

All monetary values use the **Indian numbering system** (lakhs and crores), not the Western million/billion system.

```
✓  ₹12,34,567   (twelve lakh thirty-four thousand)
✗  ₹1,234,567   (do not use)
```

Use `Intl.NumberFormat('en-IN', { style: 'currency', currency: 'INR' })` in all components.

---

## 4. Spacing & Layout

### Base Unit

`4px`. All spacing is a multiple of 4.

### Spacing Scale

| Token | Value | Usage |
|---|---|---|
| `--space-1` | `4px` | Icon gap, tight inline |
| `--space-2` | `8px` | Badge padding, tight list items |
| `--space-3` | `12px` | Default inline gap, table cell padding |
| `--space-4` | `16px` | Card internal padding (compact), form field gap |
| `--space-5` | `20px` | Form section spacing |
| `--space-6` | `24px` | Card padding (standard) |
| `--space-8` | `32px` | Section spacing |
| `--space-10` | `40px` | Page section break |
| `--space-12` | `48px` | Page-level vertical rhythm |

### Page Layout

```
┌─────────────────────────────────────────────────────────┐
│ Top Bar (60px, fixed)                                   │
├──────────────┬──────────────────────────────────────────┤
│              │                                          │
│  Sidebar     │  Main Content Area                       │
│  (248px,     │  (flex: 1, padding: 32px)               │
│   fixed)     │                                          │
│              │                                          │
└──────────────┴──────────────────────────────────────────┘
```

- **Top bar:** 60px height, fixed, z-index 50
- **Sidebar:** 248px width, fixed, full height below top bar
- **Sidebar collapsed:** 60px (icon only mode)
- **Main content:** `calc(100vw - 248px)`, padding `32px`
- **Max content width:** `1400px` (centred on very wide screens)

### Grid System

Use CSS Grid with `repeat(auto-fit, minmax(Xpx, 1fr))` for responsive layouts.

| Context | Column setup |
|---|---|
| Metric cards | `repeat(4, 1fr)` · gap `16px` |
| Dashboard 2-col | `60fr 40fr` · gap `24px` |
| Form 2-col | `1fr 1fr` · gap `16px` |
| Settings grid | `240px 1fr` · gap `32px` |

---

## 5. Elevation & Borders

### Border Radius

| Token | Value | Usage |
|---|---|---|
| `--radius-sm` | `6px` | Badges, chips, small buttons, inputs |
| `--radius-md` | `8px` | Modals, dropdowns, tooltips |
| `--radius-lg` | `16px` | Cards, metric cards, drawers, quick action cards |
| `--radius-full` | `9999px` | Status badge pills, avatar circles |

> **Rule:** Table rows always have `border-radius: 0`. Only cards and modals get the large radius.

### Shadows

| Token | Value | Usage |
|---|---|---|
| `--shadow-card` | `0 1px 4px rgba(28,25,23,0.07)` | Cards, metric cards |
| `--shadow-drawer` | `-4px 0 24px rgba(28,25,23,0.10)` | Right-side detail drawers |
| `--shadow-modal` | `0 8px 32px rgba(28,25,23,0.14)` | Centered modals |
| `--shadow-dropdown` | `0 4px 12px rgba(28,25,23,0.10)` | Dropdowns, popovers |

### Border Widths

| Context | Width |
|---|---|
| Default card, input, table | `1px` |
| Active sidebar item (left accent) | `3px` |
| Selected table row (left accent) | `3px` |
| Alert card (left accent) | `3px` |
| Focus ring | `2px offset-2 #92400E` |
| Dividers | `1px` |

---

## 6. Components

### Buttons

```
Height: 36px (default) · 32px (compact) · 40px (large)
Padding: 0 16px (default) · 0 12px (compact) · 0 20px (large)
Border radius: 8px
Font: Inter 500 14px
Transition: background 120ms ease
```

**Grouping:** Related actions use a button group with shared border. Primary action always leftmost.

**Icon buttons:** 36×36px square, ghost variant, icon 16px. Tooltip on hover (title attribute).

**Loading state:** Replace label with a 16px spinner. Disable pointer events.

### Inputs & Selects

```
Height: 36px
Padding: 0 12px
Border: 1px solid #EEEBE4
Border-radius: 6px
Background: #FFFFFF
Font: Inter 400 14px #1C1917
Placeholder: Inter 400 14px #A8A29E

Focus: border-color #92400E · outline: 2px solid #92400E · outline-offset: 2px
Error: border-color #DC2626 · error message below in 12px #DC2626
Disabled: background #FAFAF8 · text #A8A29E · cursor not-allowed
```

### Checkboxes & Radio Buttons

```
Size: 16×16px
Checked: background #1C1917 · checkmark white
Indeterminate: background #1C1917 · dash white (bulk select state)
Focus: 2px ring #92400E
```

### Tabs

```
Height: 40px
Border-bottom: 1px solid #EEEBE4 (on the tab bar, not individual tabs)
Active tab: border-bottom 2px solid #92400E · text #92400E · weight 500
Inactive: text #78716C · weight 400
Hover: text #1C1917
Gap between tabs: 4px
```

### Tooltips

```
Background: #1C1917
Text: #FFFFFF · 12px · Inter 400
Padding: 6px 10px
Border-radius: 6px
Max-width: 240px
Delay: 400ms
Position: prefer top-centre, flip if viewport-clipped
```

### Dropdown Menus

```
Background: #FFFFFF
Border: 1px solid #EEEBE4
Border-radius: 8px
Shadow: --shadow-dropdown
Min-width: 180px
Item height: 36px
Item padding: 0 12px
Item hover: background #FAF7F2
Divider: 1px solid #EEEBE4
Destructive item: text #DC2626
```

### Pagination

```
Height: 32px
Button width: 32px
Border: 1px solid #EEEBE4
Border-radius: 6px
Active page: background #FEF3C7 · text #92400E · border #FCD34D
Prev / Next: icon only
Show: "Showing X–Y of Z records" text left-aligned
```

---

## 7. Status Badges

Status badges are the single most important visual element in this system. They must be consistent across every table and drawer.

### Structure

```
display: inline-flex
align-items: center
padding: 2px 10px
border-radius: 9999px
font-size: 11px
font-weight: 500
white-space: nowrap
```

### Full Badge Reference

| Status | bg | text | When to use |
|---|---|---|---|
| In Stock | `#DCFCE7` | `#166534` | Asset is in a stockroom, unassigned |
| Checked Out | `#DBEAFE` | `#1E40AF` | Asset has an active custodian |
| Pending Ack. | `#FFFBEB` | `#92400E` | Awaiting employee sign-off |
| Reserved | `#EDE9FE` | `#5B21B6` | Future booking confirmed |
| Under Repair | `#FEF3C7` | `#92400E` | Active work order |
| Leased Out | `#FCE7F3` | `#9D174D` | Formal lease active |
| Lost | `#FEE2E2` | `#991B1B` | Reported missing, unresolved |
| Disposed | `#F3F4F6` | `#6B7280` | Written off |
| Donated | `#ECFDF5` | `#065F46` | Transferred as donation |
| Sold | `#FFF7ED` | `#9A3412` | Asset sold |

### Work Order Severity Badges

| Severity | bg | text |
|---|---|---|
| Critical | `#FEE2E2` | `#991B1B` |
| High | `#FEF3C7` | `#92400E` |
| Medium | `#FFF7ED` | `#9A3412` |
| Low | `#F3F4F6` | `#6B7280` |

### Work Order Status Badges

| Status | bg | text |
|---|---|---|
| Open | `#EDE9FE` | `#5B21B6` |
| In Progress | `#DBEAFE` | `#1E40AF` |
| Resolved | `#DCFCE7` | `#166534` |
| Closed | `#F3F4F6` | `#6B7280` |

---

## 8. Navigation Shell

### Top Bar

```
Height:     60px
Position:   fixed, top 0, full width, z-index 50
Background: #FFFFFF
Border:     1px solid #EEEBE4 (bottom only)
Padding:    0 24px
```

**Left slot:** Logo mark (24px) + "Asset Management System" text (Inter 600 15px)  
**Right slot:** Notification bell → User avatar circle → Name → Role chip → Dropdown

**Notification bell:**
- Icon: 20px, `#78716C`
- Badge: 16px circle, `#DC2626` bg, white text, 10px font, top-right of icon
- Click: opens notification overlay panel (right-aligned, 360px wide)

**User avatar:**
- 32px circle
- Initials if no photo (2 chars, Inter 500 13px, amber bg `#FEF3C7`, text `#92400E`)
- Photo if available

**Role chip:**
```
bg: #FAF7F2
border: 1px solid #EEEBE4
border-radius: 9999px
padding: 2px 8px
font-size: 11px · weight 500 · color #78716C
```

### Sidebar

```
Width:      248px (expanded) · 60px (collapsed)
Position:   fixed, left 0, top 60px, full height
Background: #FFFFFF
Border:     1px solid #EEEBE4 (right only)
Transition: width 200ms ease
```

**Nav group labels:**
```
font-size:    10px
font-weight:  500
color:        #A8A29E
text-transform: uppercase
letter-spacing: 0.08em
padding:      16px 16px 6px
not clickable
```

**Nav items:**
```
Height:       40px
Padding:      0 16px
Border-radius: 8px (inside the sidebar)
Icon:         16px, left-aligned, gap 10px from label
Font:         Inter 400 14px #78716C

Active:
  background:  #FEF3C7
  border-left: 3px solid #92400E
  color:       #92400E
  font-weight: 500

Hover (inactive):
  background: #FAF7F2
  color:      #1C1917
```

**Badge on Alerts item:** same red badge as notification bell, shown inline right of label.

**Collapsed state (60px):** icon only, centered. Tooltip shows label on hover. Section labels hidden. Logo collapses to mark only.

### Sidebar Navigation Structure

```
[Logo / App Name]

MAIN
  Dashboard
  Alerts          [badge]

ASSETS
  Asset Registry
  Check-In / Out & Reservations
  Stockrooms & Consumables

OPERATIONS
  Maintenance & Contracts

COMPLIANCE
  Audit
  Reports & Financials

─────────────────────────────
  Settings        (admin only)
  Help & Support
```

---

## 9. Forms

### Layout Principles

- Form sections use H3 headings with a `1px #EEEBE4` bottom border as a divider
- Two-column layout for fields that are paired (e.g., Purchase Date + Purchase Cost)
- Single column for long fields (Description, Notes, Address)
- Required fields: `*` asterisk in `#DC2626` after the label
- Helper text: 12px `#78716C` below the input

### Field Sizing

| Field type | Width |
|---|---|
| Short (IDs, codes, years) | `160px` fixed |
| Medium (names, amounts) | `50%` of form width |
| Long (descriptions, addresses) | `100%` |
| Date picker | `180px` fixed |
| Currency input | `200px` fixed with ₹ prefix |

### Currency Inputs

```
Prefix: ₹ (shown inside the input, left-aligned)
Alignment: right-aligned numeric value
Placeholder: "0.00"
Formatted on blur: applies Indian numbering (commas)
```

### File Upload

```
Border: 2px dashed #EEEBE4
Border-radius: 8px
Background: #FAFAF8
Padding: 24px
Text: "Drag & drop photos here, or click to browse"
      "JPG, PNG, GIF · Max 5MB each · Up to 5 files"
Accepted: accept="image/jpeg,image/png,image/gif"
Thumbnail grid: 80×80px per thumbnail, 8px gap, remove × on hover
```

### Depreciation Section Fields (critical — follow AssetTiger pattern)

```
Depreciable Asset:  Yes / No radio (default Yes)
Depreciable Cost:   ₹ amount — "Include GST, freight, installation" helper text
Salvage Value:      ₹ amount — "Residual value at end of useful life"
Asset Life:         number input + "months" suffix label
Depreciation Method: dropdown (Straight Line | Written Down Value)
Date Acquired:      date picker — "Depreciation begin date" helper text
                    defaults to Purchase Date but is independently editable
```

Auto-calculated preview (shown below the section, updates live):
```
Annual Depreciation = (Depreciable Cost − Salvage Value) ÷ Asset Life × 12
                    = (₹X,XX,XXX − ₹X,XXX) ÷ XX months × 12
                    = ₹X,XXX / year
```

### Form Actions Bar

Always sticky at the bottom of full-page forms:

```
Background: #FFFFFF
Border-top: 1px solid #EEEBE4
Padding: 16px 32px
Position: sticky bottom 0

Left side: "Save" (primary) · "Save & Clone" (secondary) · "Save & Check Out" (secondary)
Right side: "Cancel" (ghost)
```

---

## 10. Tables

Tables are the primary UI surface. Every decision must serve scannability.

### Structure

```
Border-collapse: collapse
Width: 100%
Font: Inter 400 14px

Header row:
  Background:    #FAF7F2
  Height:        40px
  Font:          Inter 500 12px #78716C UPPERCASE 0.06em tracking
  Padding:       0 16px
  Border-bottom: 1px solid #EEEBE4
  Position:      sticky top (below page title / filter bar)
  Cursor:        pointer (sortable columns)
  Sort indicator: ↑ ↓ arrow after label, #92400E when active

Data rows:
  Height:        46px
  Padding:       0 16px
  Odd rows:      background #FFFFFF
  Even rows:     background #FAFAF8
  Hover:         background #FEF3C7 (warm amber, NOT blue)
  Border-bottom: 1px solid #EEEBE4

Selected row (checkbox checked):
  Background:    #FFFBEB
  Border-left:   3px solid #92400E
```

### Checkbox Column

```
Width: 40px
Checkbox: 16×16px, centered
Header checkbox: selects/deselects all visible rows (indeterminate if partial)
```

### Column Guidelines

| Column | Width | Notes |
|---|---|---|
| Checkbox | 40px | Fixed |
| Tag ID | 120px | JetBrains Mono 13px, amber on hover |
| Asset Name | flexible (min 200px) | Truncate with ellipsis if too long, tooltip on hover |
| Category | 140px | Plain text |
| Status | 140px | Badge pill |
| Custodian | 160px | Employee name or stockroom name (italic, muted) |
| Location | 160px | Building · Floor · Room |
| Date columns | 110px | DD/MM/YYYY |
| Currency columns | 120px | Right-aligned, ₹ Indian format |
| Actions | 100px | Icon buttons, right-aligned |

### Actions Column

```
Display: flex, gap 4px, justify-content flex-end
Visible: always (do not show on hover only — accessibility)
Icon buttons: 28×28px · ghost variant · icon 14px
"⋯ More": dropdown with secondary actions
```

### Bulk Action Bar

Appears above the table when any rows are selected. Replaces the filter bar.

```
Background: #FEF3C7
Border: 1px solid #FCD34D
Border-radius: 8px
Padding: 10px 16px
Animation: slide-down 150ms ease

Left: "[X] assets selected" (Inter 500 14px #92400E)
Right: action buttons (secondary variant, compact size)
       Always include: Export Selected · Print Labels · Deselect All
       Contextual: Update Location · Dispose · ⋯ More
```

### Sorting

- Clickable column headers indicated by cursor: pointer
- Sort indicator: `↑` (ascending) `↓` (descending) in `#92400E` when active
- Default sort: newest first (created_at DESC)
- Sort state: URL-persisted (`?sort=purchase_date&dir=desc`)

### Pagination

Shown below the table:

```
Left:  "Showing 1–25 of 347 records"
Right: [<] [1] [2] [3] ... [14] [>]
       Page size selector: 25 | 50 | 100
```

---

## 11. Cards & Metric Cards

### Standard Card

```
Background:    #FFFFFF
Border:        1px solid #EEEBE4
Border-radius: 16px
Padding:       24px
Shadow:        --shadow-card
```

### Metric Card (Dashboard KPIs)

```
Background:    Pastel fill from card type (see Color System)
Border:        none
Border-radius: 16px
Padding:       20px 24px
Shadow:        --shadow-card
Min-height:    120px
Cursor:        pointer (links to filtered view)
```

Structure (top to bottom):
```
[Category label]        12px · 400 · muted ramp color
[Big Number]            28px · 600 · primary ramp color
[Sub-label]             12px · 400 · muted ramp color
```

Hover: shadow increases to `--shadow-modal`, translate-y -2px, transition 150ms.

### Quick Action Card

```
Background:    #FAF7F2
Border:        1px solid #EEEBE4
Border-radius: 16px
Padding:       20px
Cursor:        pointer
Transition:    box-shadow 150ms ease, transform 150ms ease
```

Structure:
```
[Icon]         40px, centered, muted color
[Label]        Inter 600 15px, center, primary text
[Helper]       Inter 400 12px, center, secondary text
```

Hover: shadow lifts, translate-y -2px.
First card (most-used action): Amber accent variant — bg `#FEF3C7`, border `#FCD34D`.

---

## 12. Drawers & Modals

### Right-Side Detail Drawer

```
Width:      520px
Position:   fixed, right 0, top 60px, height calc(100vh - 60px)
Background: #FFFFFF
Border-left: 1px solid #EEEBE4
Shadow:     --shadow-drawer
Z-index:    40
Animation:  slide-in from right, 200ms ease
```

**Drawer header:**
```
Height:       64px
Padding:      0 24px
Border-bottom: 1px solid #EEEBE4
Background:   #FFFFFF
Position:     sticky top 0

Left:  Asset Name (Inter 600 18px) · below: Tag ID (JetBrains Mono 13px muted)
Right: Status badge · Primary action button · Edit button · ⋯ More button · Close ×
```

**Drawer tabs:**
```
Border-bottom: 1px solid #EEEBE4 (on the tab bar)
Tab height: 40px
Same styling as global Tabs component
Padding: 0 24px (the tab bar itself)
```

**Drawer body:**
```
Padding:    24px
Overflow-y: auto
```

**Overlay (behind drawer when open):**
```
Background: rgba(28, 25, 23, 0.25)
Z-index:    39
Click to close
```

### Centered Modal

```
Max-width:     600px
Width:         100%
Background:    #FFFFFF
Border-radius: 16px
Shadow:        --shadow-modal
Padding:       0 (header and body padded separately)
Z-index:       50
Animation:     scale 0.95→1.0 + fade, 180ms ease
```

**Modal header:**
```
Padding:       20px 24px
Border-bottom: 1px solid #EEEBE4
H2:            Inter 600 18px #1C1917
Close ×:       top-right, ghost icon button
```

**Modal body:**
```
Padding: 24px
```

**Modal footer:**
```
Padding:    16px 24px
Border-top: 1px solid #EEEBE4
Flex:       row, gap 8px, justify-content flex-end
```

**Destructive confirmation modal:**

Before any destructive action (Dispose, Write Off, Delete, Bulk Delete), show this modal:

```
Header:  "Confirm [Action]"
Body:    Description of what will happen. Cannot be undone.
         "Type [ASSET NAME / DELETE] to confirm:"
         Text input (must match exactly, case-insensitive)
Footer:  [Cancel (secondary)] [Confirm (danger — disabled until text matches)]
```

---

## 13. Alert System

### Alert Cards (Dashboard Feed)

```
Background:    #FFFFFF
Border:        1px solid #EEEBE4
Border-left:   3px solid [severity color]
Border-radius: 8px
Padding:       12px 16px
Margin-bottom: 8px
```

Structure (horizontal):
```
[Icon 16px] · [Message text + asset link (bold)] · [CTA button] · [Dismiss ×]
```

Severity left-border colours:
- Critical: `#DC2626`
- Warning: `#D97706`
- Info: `#3B82F6`
- Success: `#16A34A`

### Alert Types (in dashboard priority order)

| Priority | Alert | CTA |
|---|---|---|
| 1 🔴 | Work order open > 7 days, no update | View WO |
| 2 🔴 | Asset overdue for return | View Checkout |
| 3 🔴 | Warranty expired, asset still in use | View Asset |
| 4 🔴 | Asset marked Lost, unresolved | Resolve |
| 5 🟡 | Warranty expiring within 30 days | View Asset |
| 6 🟡 | AMC contract expiring within 30 days | View Contract |
| 7 🟡 | Scheduled maintenance due within 7 days | View Schedule |
| 8 🟡 | Procurement request stalled > 3 days | View Request |
| 9 🟡 | Leased asset approaching end date | View Lease |
| 10 ⚪ | Consumable below reorder level | View Stock |

### Full-Page Alert Banner

Only shown when critical alerts exist. Shown at the top of the main content area (below page title).

```
Background:  #FEF2F2
Border:      1px solid #FECACA
Border-radius: 8px
Padding:     12px 16px
Color:       #991B1B
Font:        Inter 500 14px

Left:  warning icon · "X critical alerts require immediate attention"
Right: "View all" link · Dismiss ×
```

---

## 14. Empty & Error States

### Empty State

Every table and list must have an empty state. No blank white boxes.

```
Container: centered, padding 48px 0
Icon:      48px, color #D6D3CE (very muted)
Heading:   Inter 600 16px #1C1917 — specific to context
Helper:    Inter 400 14px #78716C — what the user should do
CTA:       primary button (optional)
```

Examples:
- Asset Registry empty: "No assets registered yet" · "Add your first asset to get started" · [Register Asset]
- Work Orders empty: "No work orders found" · "All assets are running smoothly" · no CTA
- Audit session empty: "No audit sessions yet" · "Start a session to verify your assets" · [Start Audit]
- Search no results: "No assets match '[search term]'" · "Try adjusting your filters" · [Clear Filters]

### Error State

For network failures, sync errors, or load failures:

```
Icon:   48px warning icon, #D97706
Title:  "Couldn't load [data]"
Body:   "There was a problem loading your data. Your changes have been saved."
CTA:    [Try Again] (primary)
```

Form field errors: inline, 12px `#DC2626`, below the field. Never use alert modals for field validation.

---

## 15. Motion & Transitions

Keep transitions short and purposeful. No decorative animations.

| Element | Property | Duration | Easing |
|---|---|---|---|
| Button hover | background | 120ms | ease |
| Card hover | shadow, transform | 150ms | ease |
| Sidebar expand/collapse | width | 200ms | ease |
| Drawer open/close | transform (translateX) | 200ms | ease |
| Modal open | scale, opacity | 180ms | ease |
| Dropdown open | opacity | 120ms | ease |
| Tab switch | border-bottom color | 100ms | ease |
| Row hover | background | 80ms | ease |
| Alert dismiss | opacity, height | 200ms | ease |
| Audit scan flash | background | 300ms total (150ms in, 150ms out) | ease |

No CSS animations longer than 300ms. No animated SVG illustrations. No loading spinners longer than 3 seconds without a fallback message.

**Loading skeleton:**
```
background: linear-gradient(90deg, #F3F4F6 25%, #E5E7EB 50%, #F3F4F6 75%)
background-size: 400% 100%
animation: shimmer 1.5s infinite
border-radius: 4px
```

---

## 16. Mobile (Audit Mode)

The audit scanning screen is the only screen designed primarily for mobile. All other screens are desktop-only.

### Mobile Constraints

- Viewport: 375–428px wide
- No sidebar (hidden)
- No top bar (minimal header only)
- Thumb-friendly touch targets: minimum 44×44px
- Single column layout

### Audit Mobile Layout

```
Background: #FAF7F2 (warm canvas)

[Session Header — 56px]
  Session name (Inter 600 15px) · Scope label (12px muted)

[Progress Bar — 48px]
  Label: "X / Y assets verified" (12px muted)
  Bar: full-width, 8px height, border-radius 4px
       Filled: #16A34A · Track: #EEEBE4
  Percentage: 28px 600 #1C1917, right-aligned

[SCAN ASSET Button — 72px height]
  Background: #1C1917
  Text: "SCAN ASSET" (Inter 600 16px #FFFFFF)
  Border-radius: 16px
  Full width

[Manual Entry Link — 32px]
  "Enter Tag ID manually" — amber link, centered

[Scan Result Flash — overlays button area momentarily]
  Verified:     full-width, #DCFCE7 bg, ✓ icon + asset name, auto-dismiss 1.5s
  Mismatch:     #FEF3C7 bg, ⚠ icon + location discrepancy info
  Unregistered: #FEE2E2 bg, ? icon + "Not in system" message
  Already done: #F3F4F6 bg, already-scanned message

[Collapsible Lists — below scan area]
  ✅ Verified (count) [↓]     — expand to show names
  ⚠ Unresolved (count) [↓]   — expand to show names

[END SESSION sticky bottom button — 60px]
  Background: #DC2626
  Text: "End Session & Generate Report" (Inter 600 15px #FFFFFF)
  Position: sticky bottom 0
  Always visible (no scroll-away)
```

---

## 17. Accessibility

### Keyboard Navigation

- All interactive elements reachable by Tab key in logical order
- Focus ring: `outline: 2px solid #92400E; outline-offset: 2px` — never `outline: none`
- Drawer and modal trap focus within when open
- Escape key closes drawer, modal, and dropdown
- Table sorting accessible via Enter/Space on column headers
- Skip-to-content link (visually hidden, shown on focus) at top of page

### Colour Contrast

All text meets WCAG 2.1 AA minimum:
- Body text on canvas: #1C1917 on #FAF7F2 — ratio 15.9:1 ✓
- Secondary text on canvas: #78716C on #FAF7F2 — ratio 5.2:1 ✓
- Badge text on badge bg: all combinations verified above 4.5:1
- Primary button: #FFFFFF on #1C1917 — ratio 17.6:1 ✓
- Links (#92400E) on canvas (#FAF7F2) — ratio 5.8:1 ✓

### Screen Readers

- All icon-only buttons have `aria-label`
- Status badges include visually-hidden text (not just color)
- Tables have `<caption>` and `scope` on headers
- Loading states announce via `aria-live="polite"`
- Error messages linked to inputs via `aria-describedby`
- Modals use `role="dialog"` with `aria-labelledby` and `aria-modal="true"`

### Touch Targets (Audit Mobile)

- All tappable elements minimum 44×44px
- SCAN ASSET button: 72px height (exceeds minimum — intentional, it's the primary action)
- Collapsible list headers: 44px height
- End Session button: 60px height, sticky (no mis-tap risk)

---

## 18. Design Tokens (CSS Variables)

```css
:root {

  /* Canvas & Surface */
  --color-canvas:          #FAF7F2;
  --color-surface:         #FFFFFF;
  --color-surface-subtle:  #FAFAF8;

  /* Text */
  --color-text-primary:    #1C1917;
  --color-text-secondary:  #78716C;
  --color-text-disabled:   #A8A29E;
  --color-text-inverse:    #FFFFFF;
  --color-text-link:       #92400E;

  /* Borders */
  --color-border-default:  #EEEBE4;
  --color-border-strong:   #D6D3CE;
  --color-border-accent:   #92400E;

  /* Brand Accent (Amber) */
  --color-accent:          #92400E;
  --color-accent-hover:    #78350F;
  --color-accent-subtle:   #FEF3C7;
  --color-accent-border:   #FCD34D;

  /* Status Colours */
  --status-in-stock-bg:    #DCFCE7;
  --status-in-stock-text:  #166534;
  --status-checked-out-bg: #DBEAFE;
  --status-checked-out-text:#1E40AF;
  --status-reserved-bg:    #EDE9FE;
  --status-reserved-text:  #5B21B6;
  --status-repair-bg:      #FEF3C7;
  --status-repair-text:    #92400E;
  --status-leased-bg:      #FCE7F3;
  --status-leased-text:    #9D174D;
  --status-lost-bg:        #FEE2E2;
  --status-lost-text:      #991B1B;
  --status-disposed-bg:    #F3F4F6;
  --status-disposed-text:  #6B7280;

  /* Metric Card Fills */
  --card-amber-bg:         #FEF3C7;
  --card-amber-text:       #92400E;
  --card-lavender-bg:      #EDE9FE;
  --card-lavender-text:    #5B21B6;
  --card-peach-bg:         #FEE2E2;
  --card-peach-text:       #991B1B;
  --card-pink-bg:          #FCE7F3;
  --card-pink-text:        #9D174D;
  --card-sage-bg:          #DCFCE7;
  --card-sage-text:        #166534;

  /* Alert Severity */
  --alert-critical-border: #DC2626;
  --alert-critical-bg:     #FEF2F2;
  --alert-warning-border:  #D97706;
  --alert-warning-bg:      #FFFBEB;
  --alert-info-border:     #3B82F6;
  --alert-info-bg:         #EFF6FF;

  /* Typography */
  --font-body:             'Inter', -apple-system, sans-serif;
  --font-mono:             'JetBrains Mono', monospace;

  /* Border Radius */
  --radius-sm:             6px;
  --radius-md:             8px;
  --radius-lg:             16px;
  --radius-full:           9999px;

  /* Shadows */
  --shadow-card:           0 1px 4px rgba(28, 25, 23, 0.07);
  --shadow-drawer:        -4px 0 24px rgba(28, 25, 23, 0.10);
  --shadow-modal:          0 8px 32px rgba(28, 25, 23, 0.14);
  --shadow-dropdown:       0 4px 12px rgba(28, 25, 23, 0.10);

  /* Spacing */
  --space-1:  4px;
  --space-2:  8px;
  --space-3:  12px;
  --space-4:  16px;
  --space-5:  20px;
  --space-6:  24px;
  --space-8:  32px;
  --space-10: 40px;
  --space-12: 48px;

  /* Layout */
  --sidebar-width:         248px;
  --sidebar-collapsed:     60px;
  --topbar-height:         60px;
  --drawer-width:          520px;
  --content-max-width:     1400px;

  /* Transitions */
  --transition-fast:       120ms ease;
  --transition-base:       200ms ease;
  --transition-slow:       300ms ease;

}
```

---

## Quick Reference Cheatsheet

| Thing | Value |
|---|---|
| Page canvas | `#FAF7F2` |
| Card surface | `#FFFFFF` |
| Primary accent | `#92400E` (amber) |
| Primary button | `#1C1917` bg |
| Table row hover | `#FEF3C7` (amber, NOT blue) |
| Selected row | `#FFFBEB` + left border `#92400E` |
| Active nav item | `#FEF3C7` + left border `#92400E` |
| Border (default) | `#EEEBE4` |
| Card radius | `16px` |
| Input radius | `6px` |
| Focus ring | `2px solid #92400E, offset 2px` |
| Currency format | ₹X,XX,XXX (Indian system) |
| Date format | DD/MM/YYYY |
| Font | Inter |
| Mono font | JetBrains Mono (tag IDs, serials, formulas) |
| Depreciation formula | (Cost − Salvage) ÷ Life (months) × 12 |

---

*Technopark Kerala AMS Design System · Version 1.0 · April 2026*  
*Maintain this file alongside the codebase. When a token changes in code, update this document.*
