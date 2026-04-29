<template>
<main class="p-page-margin max-w-[1400px] mx-auto space-y-section-gap">
 <!-- Header Section -->
 <div class="flex justify-between items-end">
  <div>
   <h1 class="font-h1 text-h1 text-text-primary mb-1">
    Consumables Tracker
   </h1>
   <p class="font-body text-body text-text-secondary">
    Manage and track daily supplies and consumable inventory.
   </p>
  </div>
  <div class="flex gap-inline">
   <button class="px-4 py-2 bg-surface text-text-primary border border-[#D6D3CE] rounded-lg font-h3 text-h3 hover:bg-surface-subtle transition-colors shadow-sm flex items-center gap-2">
    <span class="material-symbols-outlined text-[18px]">
     download
    </span>
    Receive Stock
   </button>
   <button class="px-4 py-2 bg-[#1C1917] text-white rounded-lg font-h3 text-h3 hover:bg-opacity-90 transition-colors shadow-sm flex items-center gap-2">
    <span class="material-symbols-outlined text-[18px]">
     output
    </span>
    Issue Consumable
   </button>
  </div>
 </div>
 <!-- KPI Cards -->
 <div class="grid grid-cols-4 gap-inline">
  <div class="bg-metric-amber rounded-xl p-card-padding shadow-[0_4px_16px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform cursor-pointer border border-[#EEEBE4]">
   <h3 class="font-metadata text-metadata text-on-surface-variant uppercase mb-2">
    Total Items
   </h3>
   <div class="font-kpi-number text-kpi-number text-text-primary">
    1,248
   </div>
   <div class="font-metadata text-metadata text-text-secondary mt-1">
    Across 12 categories
   </div>
  </div>
  <div class="bg-surface-container-high rounded-xl p-card-padding shadow-[0_4px_16px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform cursor-pointer border border-[#EEEBE4]">
   <h3 class="font-metadata text-metadata text-on-surface-variant uppercase mb-2">
    Low Stock Items
   </h3>
   <div class="font-kpi-number text-kpi-number text-primary-container">
    34
   </div>
   <div class="font-metadata text-metadata text-text-secondary mt-1 flex items-center gap-1">
    <span class="material-symbols-outlined text-[14px] text-primary-container">
     warning
    </span>
    Requires attention
   </div>
  </div>
  <div class="bg-metric-lavender rounded-xl p-card-padding shadow-[0_4px_16px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform cursor-pointer border border-[#EEEBE4]">
   <h3 class="font-metadata text-metadata text-on-surface-variant uppercase mb-2">
    Monthly Usage
   </h3>
   <div class="font-kpi-number text-kpi-number text-tertiary">
    4,892
   </div>
   <div class="font-metadata text-metadata text-text-secondary mt-1">
    Units consumed this month
   </div>
  </div>
  <div class="bg-metric-sage rounded-xl p-card-padding shadow-[0_4px_16px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform cursor-pointer border border-[#EEEBE4]">
   <h3 class="font-metadata text-metadata text-on-surface-variant uppercase mb-2">
    Total Value
   </h3>
   <div class="font-kpi-number text-kpi-number text-status-in-stock">
    ₹4.2L
   </div>
   <div class="font-metadata text-metadata text-text-secondary mt-1">
    Estimated current stock value
   </div>
  </div>
 </div>
 <!-- Filter Bar & Table Area -->
 <div class="bg-surface rounded-xl border border-border-default shadow-[0_4px_16px_rgba(0,0,0,0.02)] overflow-hidden">
  <!-- Filters -->
  <div class="p-4 border-b border-border-default flex justify-between items-center bg-surface-subtle">
   <div class="flex gap-inline">
    <div class="relative w-64">
     <span class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-stone-400 text-[18px]">
      search
     </span>
     <input class="w-full bg-surface border border-border-default rounded-lg py-1.5 pl-9 pr-3 text-body font-body focus:outline-none focus:border-amber-700 shadow-sm" placeholder="Filter by item name..." type="text"/>
    </div>
    <select class="bg-surface border border-border-default rounded-lg py-1.5 px-3 text-body font-body shadow-sm focus:outline-none focus:border-amber-700">
     <option>
      All Categories
     </option>
     <option>
      Office Supplies
     </option>
     <option>
      Pantry
     </option>
     <option>
      Cleaning
     </option>
     <option>
      IT Peripherals
     </option>
    </select>
   </div>
   <div class="flex gap-2">
    <button class="p-1.5 text-text-secondary hover:text-text-primary hover:bg-surface-variant rounded transition-colors">
     <span class="material-symbols-outlined">
      filter_list
     </span>
    </button>
    <button class="p-1.5 text-text-secondary hover:text-text-primary hover:bg-surface-variant rounded transition-colors">
     <span class="material-symbols-outlined">
      view_column
     </span>
    </button>
   </div>
  </div>
  <!-- Table -->
  <div class="overflow-x-auto">
   <table class="w-full text-left border-collapse">
    <thead>
     <tr class="bg-surface border-b border-border-default">
      <th class="font-table-header text-table-header text-text-secondary py-3 px-4 uppercase">
       Item Name
      </th>
      <th class="font-table-header text-table-header text-text-secondary py-3 px-4 uppercase">
       Category
      </th>
      <th class="font-table-header text-table-header text-text-secondary py-3 px-4 uppercase">
       Unit
      </th>
      <th class="font-table-header text-table-header text-text-secondary py-3 px-4 uppercase text-right">
       Current Stock
      </th>
      <th class="font-table-header text-table-header text-text-secondary py-3 px-4 uppercase text-right">
       Reorder Level
      </th>
      <th class="font-table-header text-table-header text-text-secondary py-3 px-4 uppercase">
       Location
      </th>
      <th class="font-table-header text-table-header text-text-secondary py-3 px-4 uppercase">
       Last Restocked
      </th>
      <th class="font-table-header text-table-header text-text-secondary py-3 px-4 uppercase text-center">
       Status
      </th>
     </tr>
    </thead>
    <tbody class="font-body text-body divide-y divide-border-default">
     <!-- Row 1 -->
     <tr class="bg-surface hover:bg-metric-amber/20 transition-colors cursor-pointer group">
      <td class="py-3 px-4 font-medium text-text-primary">
       A4 Copy Paper Reams
      </td>
      <td class="py-3 px-4 text-text-secondary">
       Office Supplies
      </td>
      <td class="py-3 px-4 text-text-secondary">
       Box (5 Reams)
      </td>
      <td class="py-3 px-4 font-mono-data text-mono-data text-right">
       142
      </td>
      <td class="py-3 px-4 font-mono-data text-mono-data text-right text-text-secondary">
       50
      </td>
      <td class="py-3 px-4 text-text-secondary">
       Main Store, Rack A1
      </td>
      <td class="py-3 px-4 text-text-secondary">
       12 Oct 2023
      </td>
      <td class="py-3 px-4 text-center">
       <span class="inline-flex items-center px-2 py-0.5 rounded-full text-[11px] font-medium bg-metric-sage text-status-in-stock whitespace-nowrap">
        Adequate
       </span>
      </td>
     </tr>
     <!-- Row 2 -->
     <tr class="bg-surface-subtle hover:bg-metric-amber/20 transition-colors cursor-pointer border-l-[3px] border-l-amber-500">
      <td class="py-3 px-4 font-medium text-text-primary">
       Printer Toner Cartridge (Black)
      </td>
      <td class="py-3 px-4 text-text-secondary">
       IT Peripherals
      </td>
      <td class="py-3 px-4 text-text-secondary">
       Piece
      </td>
      <td class="py-3 px-4 font-mono-data text-mono-data text-right text-primary-container font-bold">
       4
      </td>
      <td class="py-3 px-4 font-mono-data text-mono-data text-right text-text-secondary">
       10
      </td>
      <td class="py-3 px-4 text-text-secondary">
       IT Cupboard 2
      </td>
      <td class="py-3 px-4 text-text-secondary">
       01 Sep 2023
      </td>
      <td class="py-3 px-4 text-center">
       <span class="inline-flex items-center px-2 py-0.5 rounded-full text-[11px] font-medium bg-metric-amber text-primary-container whitespace-nowrap">
        Low Stock
       </span>
      </td>
     </tr>
     <!-- Row 3 -->
     <tr class="bg-surface hover:bg-metric-amber/20 transition-colors cursor-pointer">
      <td class="py-3 px-4 font-medium text-text-primary">
       Whiteboard Markers (Pack of 4)
      </td>
      <td class="py-3 px-4 text-text-secondary">
       Office Supplies
      </td>
      <td class="py-3 px-4 text-text-secondary">
       Pack
      </td>
      <td class="py-3 px-4 font-mono-data text-mono-data text-right">
       0
      </td>
      <td class="py-3 px-4 font-mono-data text-mono-data text-right text-text-secondary">
       15
      </td>
      <td class="py-3 px-4 text-text-secondary">
       Stationery Cabinet
      </td>
      <td class="py-3 px-4 text-text-secondary">
       15 Aug 2023
      </td>
      <td class="py-3 px-4 text-center">
       <span class="inline-flex items-center px-2 py-0.5 rounded-full text-[11px] font-medium bg-error-container text-on-error-container whitespace-nowrap">
        Out of Stock
       </span>
      </td>
     </tr>
     <!-- Row 4 -->
     <tr class="bg-surface-subtle hover:bg-metric-amber/20 transition-colors cursor-pointer">
      <td class="py-3 px-4 font-medium text-text-primary">
       Paper Coffee Cups
      </td>
      <td class="py-3 px-4 text-text-secondary">
       Pantry
      </td>
      <td class="py-3 px-4 text-text-secondary">
       Sleeve (50 pcs)
      </td>
      <td class="py-3 px-4 font-mono-data text-mono-data text-right">
       85
      </td>
      <td class="py-3 px-4 font-mono-data text-mono-data text-right text-text-secondary">
       20
      </td>
      <td class="py-3 px-4 text-text-secondary">
       Pantry Store
      </td>
      <td class="py-3 px-4 text-text-secondary">
       20 Oct 2023
      </td>
      <td class="py-3 px-4 text-center">
       <span class="inline-flex items-center px-2 py-0.5 rounded-full text-[11px] font-medium bg-metric-sage text-status-in-stock whitespace-nowrap">
        Adequate
       </span>
      </td>
     </tr>
     <!-- Row 5 -->
     <tr class="bg-surface hover:bg-metric-amber/20 transition-colors cursor-pointer">
      <td class="py-3 px-4 font-medium text-text-primary">
       Hand Sanitizer (500ml)
      </td>
      <td class="py-3 px-4 text-text-secondary">
       Cleaning
      </td>
      <td class="py-3 px-4 text-text-secondary">
       Bottle
      </td>
      <td class="py-3 px-4 font-mono-data text-mono-data text-right text-primary-container font-bold">
       12
      </td>
      <td class="py-3 px-4 font-mono-data text-mono-data text-right text-text-secondary">
       30
      </td>
      <td class="py-3 px-4 text-text-secondary">
       Janitorial Closet
      </td>
      <td class="py-3 px-4 text-text-secondary">
       05 Oct 2023
      </td>
      <td class="py-3 px-4 text-center">
       <span class="inline-flex items-center px-2 py-0.5 rounded-full text-[11px] font-medium bg-metric-amber text-primary-container whitespace-nowrap">
        Low Stock
       </span>
      </td>
     </tr>
    </tbody>
   </table>
  </div>
  <!-- Pagination (simplified) -->
  <div class="p-4 border-t border-border-default flex justify-between items-center bg-surface-subtle">
   <span class="font-metadata text-metadata text-text-secondary">
    Showing 1 to 5 of 1,248 entries
   </span>
   <div class="flex gap-1">
    <button class="px-2 py-1 rounded border border-border-default text-text-secondary hover:bg-surface disabled:opacity-50" disabled>
     Prev
    </button>
    <button class="px-2 py-1 rounded border border-border-default bg-surface text-text-primary font-medium">
     1
    </button>
    <button class="px-2 py-1 rounded border border-border-default text-text-secondary hover:bg-surface">
     2
    </button>
    <button class="px-2 py-1 rounded border border-border-default text-text-secondary hover:bg-surface">
     3
    </button>
    <span class="px-2 py-1 text-text-secondary">
     ...
    </span>
    <button class="px-2 py-1 rounded border border-border-default text-text-secondary hover:bg-surface">
     Next
    </button>
   </div>
  </div>
 </div>
</main>

</template>

<script setup lang="ts">
// Autogenerated from consumables_supplies_tracker
</script>
