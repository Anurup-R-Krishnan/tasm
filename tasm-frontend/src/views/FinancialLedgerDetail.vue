<template>
<main class="flex-1 flex flex-col min-w-0">
 <!-- TopNavBar (Shared Component) -->
 <header class="h-[60px] w-full border-b sticky top-0 z-50 border-stone-200 dark:border-stone-800 shadow-sm bg-[#FFFFFF] dark:bg-stone-900 flex items-center justify-between px-6 font-['Inter'] antialiased">
  <div class="flex items-center gap-4">
   <!-- Search Bar Placeholder based on "search_bar": "on_left" -->
   <div class="relative hidden md:block">
    <span class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-stone-400 text-[18px]">
     search
    </span>
    <input class="pl-10 pr-4 py-1.5 rounded-full border border-border-default bg-surface-subtle focus:outline-none focus:ring-2 focus:ring-surface-tint focus:border-transparent font-body text-body text-text-primary w-[240px]" placeholder="Search ledger..." type="text"/>
   </div>
  </div>
  <div class="flex items-center gap-4">
   <!-- Trailing Icons -->
   <button class="w-8 h-8 rounded-full flex items-center justify-center text-stone-600 dark:text-stone-400 hover:bg-stone-50 dark:hover:bg-stone-800 transition-colors cursor-pointer active:opacity-80">
    <span class="material-symbols-outlined text-[20px]">
     notifications
    </span>
   </button>
   <button class="w-8 h-8 rounded-full flex items-center justify-center text-stone-600 dark:text-stone-400 hover:bg-stone-50 dark:hover:bg-stone-800 transition-colors cursor-pointer active:opacity-80">
    <span class="material-symbols-outlined text-[20px]">
     help
    </span>
   </button>
   <button class="w-8 h-8 rounded-full flex items-center justify-center text-stone-600 dark:text-stone-400 hover:bg-stone-50 dark:hover:bg-stone-800 transition-colors cursor-pointer active:opacity-80">
    <span class="material-symbols-outlined text-[20px]">
     settings
    </span>
   </button>
   <div class="w-px h-6 bg-border-default mx-1">
   </div>
   <!-- Profile Avatar -->
   <div class="w-8 h-8 rounded-full bg-tertiary-container overflow-hidden cursor-pointer active:opacity-80 border border-border-default">
    <img alt="User Profile Avatar" class="w-full h-full object-cover" data-alt="Professional headshot of a corporate admin user on a neutral background" src="https://lh3.googleusercontent.com/aida-public/AB6AXuDJj4xaxAuI5MwjnYN4bZfvrd7KmTpOfz0qOzYJCJZWGidD0nvdchDT3wowYcsPLzhXJC6lDvCdr3Kvctc0AkQE9Fg2CMIrQa2psh1KXtMOO5hTl--JYbrs3sW8_nKBJ3CDszLfmkKEzmYla6Kny4iROF6mTVReLBeBSAxQiqqNTpFJIMYYim7KMlRtadQqPkJxpYocdSZHgioe_zHaETWNW3QwBu_sTEmRKsq7Vsdj8S7jUefdtzXKXUiGpjIJ1miqJM-QRU0oe3w6"/>
   </div>
  </div>
 </header>
 <!-- Page Content -->
 <div class="p-page-margin w-full max-w-[1400px] mx-auto flex-1">
  <!-- Page Header -->
  <div class="flex flex-col md:flex-row md:items-end justify-between gap-4 mb-section-gap">
   <div>
    <h1 class="font-h1 text-h1 text-text-primary tracking-tight">
     Financial Ledger
    </h1>
    <p class="font-body text-body text-text-secondary mt-1">
     Comprehensive view of asset valuations and depreciation metrics.
    </p>
   </div>
   <div class="flex gap-3">
    <button class="px-4 py-2 rounded-lg bg-surface border border-border-default font-h3 text-h3 text-text-primary shadow-[0_2px_4px_rgba(0,0,0,0.02)] hover:bg-surface-subtle transition-colors flex items-center gap-2">
     <span class="material-symbols-outlined text-[18px]">
      download
     </span>
     Export CSV
    </button>
    <button class="px-4 py-2 rounded-lg bg-[#1C1917] text-white font-h3 text-h3 shadow-sm hover:opacity-90 transition-opacity flex items-center gap-2">
     <span class="material-symbols-outlined text-[18px]">
      filter_list
     </span>
     Filter View
    </button>
   </div>
  </div>
  <!-- Summary Cards (Bento Grid) -->
  <div class="grid grid-cols-1 md:grid-cols-3 gap-section-gap mb-section-gap">
   <!-- Total Purchase Cost -->
   <div class="bg-metric-amber rounded-xl p-card-padding border border-border-default shadow-[0_4px_4px_rgba(0,0,0,0.02)] hover:-translate-y-[2px] transition-transform cursor-pointer relative overflow-hidden group">
    <div class="relative z-10">
     <div class="flex justify-between items-start mb-stack">
      <span class="font-h3 text-h3 text-on-surface-variant">
       Total Purchase Cost
      </span>
      <div class="w-8 h-8 rounded-full bg-white/50 flex items-center justify-center">
       <span class="material-symbols-outlined text-[18px] text-surface-tint">
        shopping_cart
       </span>
      </div>
     </div>
     <div class="font-kpi-number text-kpi-number text-on-surface">
      ₹ 12,45,50,000
     </div>
     <div class="font-metadata text-metadata text-status-in-stock mt-2 flex items-center gap-1">
      <span class="material-symbols-outlined text-[16px]">
       trending_up
      </span>
      +4.2% YTD
     </div>
    </div>
    <div class="absolute -bottom-4 -right-4 text-surface-tint opacity-10 transform scale-150 group-hover:scale-110 transition-transform duration-500">
     <span class="material-symbols-outlined text-[120px]" style="font-variation-settings: 'FILL' 1;">
      account_balance
     </span>
    </div>
   </div>
   <!-- Accumulated Depreciation -->
   <div class="bg-metric-lavender rounded-xl p-card-padding border border-border-default shadow-[0_4px_4px_rgba(0,0,0,0.02)] hover:-translate-y-[2px] transition-transform cursor-pointer relative overflow-hidden group">
    <div class="relative z-10">
     <div class="flex justify-between items-start mb-stack">
      <span class="font-h3 text-h3 text-on-surface-variant">
       Accumulated Depreciation
      </span>
      <div class="w-8 h-8 rounded-full bg-white/50 flex items-center justify-center">
       <span class="material-symbols-outlined text-[18px] text-tertiary">
        trending_down
       </span>
      </div>
     </div>
     <div class="font-kpi-number text-kpi-number text-on-surface">
      ₹ 3,12,80,000
     </div>
     <div class="font-metadata text-metadata text-text-secondary mt-2 flex items-center gap-1">
      Calculated via Straight-Line
     </div>
    </div>
    <div class="absolute -bottom-4 -right-4 text-tertiary opacity-10 transform scale-150 group-hover:scale-110 transition-transform duration-500">
     <span class="material-symbols-outlined text-[120px]" style="font-variation-settings: 'FILL' 1;">
      show_chart
     </span>
    </div>
   </div>
   <!-- Net Book Value -->
   <div class="bg-metric-sage rounded-xl p-card-padding border border-border-default shadow-[0_4px_4px_rgba(0,0,0,0.02)] hover:-translate-y-[2px] transition-transform cursor-pointer relative overflow-hidden group">
    <div class="relative z-10">
     <div class="flex justify-between items-start mb-stack">
      <span class="font-h3 text-h3 text-on-surface-variant">
       Net Book Value
      </span>
      <div class="w-8 h-8 rounded-full bg-white/50 flex items-center justify-center">
       <span class="material-symbols-outlined text-[18px] text-status-in-stock">
        verified
       </span>
      </div>
     </div>
     <div class="font-kpi-number text-kpi-number text-on-surface">
      ₹ 9,32,70,000
     </div>
     <div class="font-metadata text-metadata text-status-in-stock mt-2 flex items-center gap-1">
      <span class="material-symbols-outlined text-[16px]">
       check_circle
      </span>
      Audited 12 Oct 2023
     </div>
    </div>
    <div class="absolute -bottom-4 -right-4 text-status-in-stock opacity-10 transform scale-150 group-hover:scale-110 transition-transform duration-500">
     <span class="material-symbols-outlined text-[120px]" style="font-variation-settings: 'FILL' 1;">
      savings
     </span>
    </div>
   </div>
  </div>
  <!-- Chart Area -->
  <div class="bg-surface rounded-xl p-card-padding border border-border-default shadow-[0_4px_4px_rgba(0,0,0,0.02)] mb-section-gap flex flex-col h-[400px]">
   <div class="flex justify-between items-center mb-6">
    <h2 class="font-h2 text-h2 text-text-primary">
     Asset Value vs Depreciation Trend (FY 2023-24)
    </h2>
    <div class="flex gap-4 items-center">
     <div class="flex items-center gap-2">
      <div class="w-3 h-3 rounded-full bg-surface-tint">
      </div>
      <span class="font-metadata text-metadata text-text-secondary">
       Gross Value
      </span>
     </div>
     <div class="flex items-center gap-2">
      <div class="w-3 h-3 rounded-full bg-tertiary">
      </div>
      <span class="font-metadata text-metadata text-text-secondary">
       Depreciation
      </span>
     </div>
    </div>
   </div>
   <!-- CSS Area Chart Visual Representation -->
   <div class="flex-1 relative mt-4 border-l border-b border-border-default flex items-end">
    <!-- Y Axis Labels -->
    <div class="absolute -left-12 top-0 h-full flex flex-col justify-between items-end py-2 font-mono-data text-mono-data text-text-secondary">
     <span>
      12 Cr
     </span>
     <span>
      9 Cr
     </span>
     <span>
      6 Cr
     </span>
     <span>
      3 Cr
     </span>
     <span>
      0
     </span>
    </div>
    <!-- Grid Lines -->
    <div class="absolute inset-0 flex flex-col justify-between z-0">
     <div class="border-t border-border-default w-full opacity-50">
     </div>
     <div class="border-t border-border-default w-full opacity-50">
     </div>
     <div class="border-t border-border-default w-full opacity-50">
     </div>
     <div class="border-t border-border-default w-full opacity-50">
     </div>
     <div class="border-t border-border-default w-full opacity-50">
     </div>
    </div>
    <!-- Chart SVG Placeholder (using pure CSS/SVG shapes for visual fidelity) -->
    <div class="w-full h-full relative z-10 px-4 flex items-end">
     <svg class="w-full h-full overflow-visible" preserveaspectratio="none" viewbox="0 0 100 100">
      <!-- Gross Value Area -->
      <path d="M0,10 L20,15 L40,12 L60,20 L80,18 L100,25 L100,100 L0,100 Z" fill="#fff1eb" opacity="0.6">
      </path>
      <path d="M0,10 L20,15 L40,12 L60,20 L80,18 L100,25" fill="none" stroke="#9a4614" stroke-width="2">
      </path>
      <!-- Depreciation Area -->
      <path d="M0,90 L20,85 L40,80 L60,70 L80,60 L100,50 L100,100 L0,100 Z" fill="#cee5ff" opacity="0.5">
      </path>
      <path d="M0,90 L20,85 L40,80 L60,70 L80,60 L100,50" fill="none" stroke="#005c92" stroke-width="2">
      </path>
     </svg>
    </div>
    <!-- X Axis Labels -->
    <div class="absolute -bottom-8 left-0 w-full flex justify-between px-4 font-metadata text-metadata text-text-secondary">
     <span>
      Q1
     </span>
     <span>
      Q2
     </span>
     <span>
      Q3
     </span>
     <span>
      Q4
     </span>
    </div>
   </div>
  </div>
  <!-- High-Density Table Section -->
  <div class="bg-surface rounded-xl border border-border-default shadow-[0_4px_4px_rgba(0,0,0,0.02)] flex flex-col overflow-hidden">
   <div class="px-card-padding py-4 border-b border-border-default flex justify-between items-center bg-surface">
    <h2 class="font-h2 text-h2 text-text-primary">
     Recent Ledger Entries
    </h2>
    <div class="relative">
     <span class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-stone-400 text-[16px]">
      search
     </span>
     <input class="pl-9 pr-4 py-1.5 rounded-lg border border-border-default bg-surface-subtle text-sm focus:outline-none focus:ring-1 focus:ring-surface-tint focus:border-surface-tint w-[200px]" placeholder="Search entries..." type="text"/>
    </div>
   </div>
   <div class="overflow-x-auto custom-scrollbar">
    <table class="w-full text-left border-collapse min-w-[800px]">
     <thead>
      <tr class="bg-surface-subtle border-b border-border-default">
       <th class="px-4 py-3 font-table-header text-table-header text-text-secondary whitespace-nowrap">
        ENTRY ID
       </th>
       <th class="px-4 py-3 font-table-header text-table-header text-text-secondary whitespace-nowrap">
        DATE
       </th>
       <th class="px-4 py-3 font-table-header text-table-header text-text-secondary whitespace-nowrap">
        CATEGORY
       </th>
       <th class="px-4 py-3 font-table-header text-table-header text-text-secondary">
        DESCRIPTION
       </th>
       <th class="px-4 py-3 font-table-header text-table-header text-text-secondary text-right whitespace-nowrap">
        DEBIT (₹)
       </th>
       <th class="px-4 py-3 font-table-header text-table-header text-text-secondary text-right whitespace-nowrap">
        CREDIT (₹)
       </th>
       <th class="px-4 py-3 font-table-header text-table-header text-text-secondary text-right whitespace-nowrap">
        BALANCE (₹)
       </th>
      </tr>
     </thead>
     <tbody class="font-body text-body text-text-primary divide-y divide-border-default">
      <!-- Row 1 -->
      <tr class="hover:bg-metric-amber/20 transition-colors bg-surface group cursor-pointer">
       <td class="px-4 py-3 font-mono-data text-mono-data text-surface-tint">
        TRX-9921
       </td>
       <td class="px-4 py-3 whitespace-nowrap">
        24 Oct 2023
       </td>
       <td class="px-4 py-3">
        <span class="inline-flex items-center px-2 py-0.5 rounded-md bg-surface-container text-primary-container font-metadata text-metadata border border-surface-dim">
         IT Hardware
        </span>
       </td>
       <td class="px-4 py-3 max-w-[200px] truncate" title="Procurement of 50 Dell Latitude Laptops for new joiners">
        Procurement of 50 Dell Latitude Laptops for new...
       </td>
       <td class="px-4 py-3 text-right font-mono-data text-mono-data">
        35,50,000.00
       </td>
       <td class="px-4 py-3 text-right font-mono-data text-mono-data text-text-secondary">
        -
       </td>
       <td class="px-4 py-3 text-right font-mono-data text-mono-data font-medium">
        12,45,50,000.00
       </td>
      </tr>
      <!-- Row 2 -->
      <tr class="hover:bg-metric-amber/20 transition-colors bg-surface-subtle group cursor-pointer">
       <td class="px-4 py-3 font-mono-data text-mono-data text-surface-tint">
        DEP-8432
       </td>
       <td class="px-4 py-3 whitespace-nowrap">
        30 Sep 2023
       </td>
       <td class="px-4 py-3">
        <span class="inline-flex items-center px-2 py-0.5 rounded-md bg-tertiary-fixed text-on-tertiary-fixed-variant font-metadata text-metadata border border-tertiary-fixed-dim">
         Depreciation
        </span>
       </td>
       <td class="px-4 py-3 max-w-[200px] truncate" title="Q3 Straight-Line Depreciation across all vehicle assets">
        Q3 Straight-Line Depreciation across all vehicle...
       </td>
       <td class="px-4 py-3 text-right font-mono-data text-mono-data text-text-secondary">
        -
       </td>
       <td class="px-4 py-3 text-right font-mono-data text-mono-data">
        4,25,000.00
       </td>
       <td class="px-4 py-3 text-right font-mono-data text-mono-data font-medium">
        12,10,00,000.00
       </td>
      </tr>
      <!-- Row 3 -->
      <tr class="hover:bg-metric-amber/20 transition-colors bg-surface group cursor-pointer border-l-[3px] border-l-surface-tint">
       <td class="px-4 py-3 font-mono-data text-mono-data text-surface-tint pl-3">
        TRX-9905
       </td>
       <td class="px-4 py-3 whitespace-nowrap">
        15 Sep 2023
       </td>
       <td class="px-4 py-3">
        <span class="inline-flex items-center px-2 py-0.5 rounded-md bg-surface-container text-primary-container font-metadata text-metadata border border-surface-dim">
         Infrastructure
        </span>
       </td>
       <td class="px-4 py-3 max-w-[200px] truncate" title="HVAC System upgrade for Phase 3 Building">
        HVAC System upgrade for Phase 3 Building
       </td>
       <td class="px-4 py-3 text-right font-mono-data text-mono-data">
        1,12,00,000.00
       </td>
       <td class="px-4 py-3 text-right font-mono-data text-mono-data text-text-secondary">
        -
       </td>
       <td class="px-4 py-3 text-right font-mono-data text-mono-data font-medium">
        12,14,25,000.00
       </td>
      </tr>
      <!-- Row 4 -->
      <tr class="hover:bg-metric-amber/20 transition-colors bg-surface-subtle group cursor-pointer">
       <td class="px-4 py-3 font-mono-data text-mono-data text-surface-tint">
        DIS-1120
       </td>
       <td class="px-4 py-3 whitespace-nowrap">
        05 Sep 2023
       </td>
       <td class="px-4 py-3">
        <span class="inline-flex items-center px-2 py-0.5 rounded-md bg-error-container text-on-error-container font-metadata text-metadata border border-surface-dim">
         Disposal
        </span>
       </td>
       <td class="px-4 py-3 max-w-[200px] truncate" title="Sale of obsolete server racks (End of Life)">
        Sale of obsolete server racks (End of Life)
       </td>
       <td class="px-4 py-3 text-right font-mono-data text-mono-data text-text-secondary">
        -
       </td>
       <td class="px-4 py-3 text-right font-mono-data text-mono-data">
        8,50,000.00
       </td>
       <td class="px-4 py-3 text-right font-mono-data text-mono-data font-medium">
        11,02,25,000.00
       </td>
      </tr>
      <!-- Row 5 -->
      <tr class="hover:bg-metric-amber/20 transition-colors bg-surface group cursor-pointer">
       <td class="px-4 py-3 font-mono-data text-mono-data text-surface-tint">
        TRX-9884
       </td>
       <td class="px-4 py-3 whitespace-nowrap">
        22 Aug 2023
       </td>
       <td class="px-4 py-3">
        <span class="inline-flex items-center px-2 py-0.5 rounded-md bg-surface-container text-primary-container font-metadata text-metadata border border-surface-dim">
         Furniture
        </span>
       </td>
       <td class="px-4 py-3 max-w-[200px] truncate" title="Ergonomic chairs for new incubation center">
        Ergonomic chairs for new incubation center
       </td>
       <td class="px-4 py-3 text-right font-mono-data text-mono-data">
        15,75,000.00
       </td>
       <td class="px-4 py-3 text-right font-mono-data text-mono-data text-text-secondary">
        -
       </td>
       <td class="px-4 py-3 text-right font-mono-data text-mono-data font-medium">
        11,10,75,000.00
       </td>
      </tr>
     </tbody>
    </table>
   </div>
   <!-- Table Pagination -->
   <div class="px-card-padding py-3 border-t border-border-default flex justify-between items-center bg-surface-subtle">
    <span class="font-metadata text-metadata text-text-secondary">
     Showing 1 to 5 of 248 entries
    </span>
    <div class="flex gap-1">
     <button class="w-8 h-8 rounded border border-border-default flex items-center justify-center text-text-secondary hover:bg-surface disabled:opacity-50" disabled="">
      <span class="material-symbols-outlined text-[18px]">
       chevron_left
      </span>
     </button>
     <button class="w-8 h-8 rounded border border-surface-tint bg-surface-tint text-white font-metadata text-metadata flex items-center justify-center">
      1
     </button>
     <button class="w-8 h-8 rounded border border-border-default bg-surface hover:bg-surface-subtle text-text-primary font-metadata text-metadata flex items-center justify-center">
      2
     </button>
     <button class="w-8 h-8 rounded border border-border-default bg-surface hover:bg-surface-subtle text-text-primary font-metadata text-metadata flex items-center justify-center">
      3
     </button>
     <button class="w-8 h-8 rounded border border-border-default flex items-center justify-center text-text-secondary hover:bg-surface">
      <span class="material-symbols-outlined text-[18px]">
       chevron_right
      </span>
     </button>
    </div>
   </div>
  </div>
 </div>
</main>

</template>

<script setup lang="ts">
// Autogenerated from financial_ledger_detail
</script>
