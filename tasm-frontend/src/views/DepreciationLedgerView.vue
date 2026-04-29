<template>
<main class="flex-1 w-full max-w-[1400px] mx-auto p-page-margin space-y-section-gap">
 <!-- Header Section -->
 <div class="flex flex-col gap-base">
  <h1 class="font-h1 text-h1 text-text-primary">
   Asset Depreciation Ledger
  </h1>
  <div class="flex items-center gap-2 text-text-secondary font-body text-body">
   <span class="material-symbols-outlined text-[18px]">
    domain
   </span>
   <span>
    Asset ID: BLD-2019-0045
   </span>
   <span class="text-border-default mx-1">
    •
   </span>
   <span>
    Generator Set 500kVA
   </span>
   <span class="text-border-default mx-1">
    •
   </span>
   <span class="bg-surface-variant text-on-surface-variant px-2 py-0.5 rounded-full text-[11px] font-medium tracking-wide">
    STRAIGHT LINE
   </span>
  </div>
 </div>
 <!-- Bento Grid: Metrics & Chart -->
 <div class="grid grid-cols-1 lg:grid-cols-3 gap-stack">
  <!-- Purchase Price Card -->
  <div class="bg-metric-amber rounded-xl p-card-padding flex flex-col justify-between border border-[#fce9ab] shadow-sm transform transition-transform hover:-translate-y-[2px]">
   <div class="flex items-start justify-between">
    <span class="font-table-header text-table-header text-primary-container uppercase">
     Original Purchase Price
    </span>
    <span class="material-symbols-outlined text-primary-container opacity-50">
     receipt_long
    </span>
   </div>
   <div class="mt-stack font-kpi-number text-kpi-number text-on-background tracking-tight">
    ₹ 25,00,000
   </div>
   <div class="mt-stack flex items-center gap-2 border-t border-[#fce9ab] pt-inline">
    <span class="material-symbols-outlined text-[16px] text-primary-container">
     calendar_month
    </span>
    <span class="font-metadata text-metadata text-primary-container">
     Acquired: 01 Apr 2020
    </span>
   </div>
  </div>
  <!-- Current Book Value Card -->
  <div class="bg-surface-container rounded-xl p-card-padding flex flex-col justify-between border border-surface-container-high shadow-sm transform transition-transform hover:-translate-y-[2px]">
   <div class="flex items-start justify-between">
    <span class="font-table-header text-table-header text-on-surface-variant uppercase">
     Current Book Value
    </span>
    <span class="material-symbols-outlined text-on-surface-variant opacity-50">
     account_balance_wallet
    </span>
   </div>
   <div class="mt-stack flex items-baseline gap-2">
    <span class="font-kpi-number text-kpi-number text-on-background tracking-tight">
     ₹ 10,00,000
    </span>
    <span class="font-metadata text-metadata text-error bg-error-container px-2 py-0.5 rounded-full">
     -60%
    </span>
   </div>
   <div class="mt-stack flex items-center gap-2 border-t border-surface-container-high pt-inline">
    <span class="material-symbols-outlined text-[16px] text-on-surface-variant">
     schedule
    </span>
    <span class="font-metadata text-metadata text-on-surface-variant">
     Year 4 of 5 Useful Life
    </span>
   </div>
  </div>
  <!-- Chart Card -->
  <div class="bg-surface border border-border-default rounded-xl p-card-padding shadow-[0_4px_4px_rgba(0,0,0,0.02)] flex flex-col">
   <span class="font-h3 text-h3 text-text-primary mb-stack flex items-center gap-2">
    <span class="material-symbols-outlined text-[20px] text-primary">
     trending_down
    </span>
    Value Decline Projection
   </span>
   <div class="flex-1 relative w-full h-[120px]">
    <!-- CSS/SVG Representation of Line Chart -->
    <svg class="w-full h-full overflow-visible" preserveaspectratio="none" viewbox="0 0 100 50">
     <!-- X Axis / Grid -->
     <line stroke="#EEEBE4" stroke-width="1" x1="0" x2="100" y1="50" y2="50">
     </line>
     <line stroke="#EEEBE4" stroke-dasharray="2 2" stroke-width="0.5" x1="0" x2="100" y1="25" y2="25">
     </line>
     <!-- Area Fill -->
     <path d="M 0 50 L 0 5 L 25 16.25 L 50 27.5 L 75 38.75 L 100 50 Z" fill="#FEF3C7" opacity="0.3">
     </path>
     <!-- Data Line (Straight Line Depreciation) -->
     <path d="M 0 5 L 25 16.25 L 50 27.5 L 75 38.75 L 100 50" fill="none" stroke="#92400E" stroke-linecap="round" stroke-linejoin="round" stroke-width="2">
     </path>
     <!-- Data Points -->
     <circle cx="0" cy="5" fill="#FFFFFF" r="2.5" stroke="#92400E" stroke-width="1.5">
     </circle>
     <circle cx="25" cy="16.25" fill="#FFFFFF" r="2.5" stroke="#92400E" stroke-width="1.5">
     </circle>
     <circle cx="50" cy="27.5" fill="#FFFFFF" r="2.5" stroke="#92400E" stroke-width="1.5">
     </circle>
     <!-- Current Year Highlight -->
     <circle cx="75" cy="38.75" fill="#92400E" r="3.5" stroke="#FFFFFF" stroke-width="1.5">
     </circle>
     <circle cx="75" cy="38.75" fill="none" opacity="0.5" r="6" stroke="#92400E" stroke-width="1">
     </circle>
     <circle cx="100" cy="50" fill="#FFFFFF" r="2.5" stroke="#92400E" stroke-width="1.5">
     </circle>
    </svg>
   </div>
  </div>
 </div>
 <!-- Ledger Table Section -->
 <div class="bg-surface border border-border-default rounded-xl shadow-[0_4px_4px_rgba(0,0,0,0.02)] overflow-hidden">
  <div class="px-card-padding py-stack border-b border-border-default flex justify-between items-center bg-surface-subtle">
   <h2 class="font-h2 text-h2 text-text-primary">
    5-Year Depreciation Schedule
   </h2>
   <button class="bg-surface border border-outline-variant rounded px-inline py-1.5 font-body text-body text-text-primary hover:bg-surface-container flex items-center gap-2 transition-colors text-sm font-medium shadow-sm">
    <span class="material-symbols-outlined text-[18px]">
     download
    </span>
    Export CSV
   </button>
  </div>
  <div class="overflow-x-auto">
   <table class="w-full text-left border-collapse">
    <thead>
     <tr class="bg-surface-subtle border-b border-border-default font-table-header text-table-header text-text-secondary">
      <th class="px-card-padding py-inline font-medium text-left w-1/5">
       Year
      </th>
      <th class="px-card-padding py-inline font-medium text-right w-1/5">
       Beginning Value (₹)
      </th>
      <th class="px-card-padding py-inline font-medium text-right w-1/5">
       Depreciation Expense (₹)
      </th>
      <th class="px-card-padding py-inline font-medium text-right w-1/5">
       Accumulated Dep. (₹)
      </th>
      <th class="px-card-padding py-inline font-medium text-right w-1/5">
       Ending Book Value (₹)
      </th>
     </tr>
    </thead>
    <tbody class="font-mono-data text-mono-data text-text-primary divide-y divide-border-default">
     <!-- Year 1 -->
     <tr class="hover:bg-metric-amber/30 transition-colors group">
      <td class="px-card-padding py-stack text-left border-l-[3px] border-transparent group-hover:border-primary-container">
       <div class="flex items-center gap-2 font-body text-body">
        2020-21
        <span class="text-text-secondary text-xs">
         (Y1)
        </span>
       </div>
      </td>
      <td class="px-card-padding py-stack text-right">
       25,00,000
      </td>
      <td class="px-card-padding py-stack text-right text-error">
       5,00,000
      </td>
      <td class="px-card-padding py-stack text-right">
       5,00,000
      </td>
      <td class="px-card-padding py-stack text-right font-medium">
       20,00,000
      </td>
     </tr>
     <!-- Year 2 -->
     <tr class="hover:bg-metric-amber/30 transition-colors group">
      <td class="px-card-padding py-stack text-left border-l-[3px] border-transparent group-hover:border-primary-container">
       <div class="flex items-center gap-2 font-body text-body">
        2021-22
        <span class="text-text-secondary text-xs">
         (Y2)
        </span>
       </div>
      </td>
      <td class="px-card-padding py-stack text-right">
       20,00,000
      </td>
      <td class="px-card-padding py-stack text-right text-error">
       5,00,000
      </td>
      <td class="px-card-padding py-stack text-right">
       10,00,000
      </td>
      <td class="px-card-padding py-stack text-right font-medium">
       15,00,000
      </td>
     </tr>
     <!-- Year 3 -->
     <tr class="hover:bg-metric-amber/30 transition-colors group">
      <td class="px-card-padding py-stack text-left border-l-[3px] border-transparent group-hover:border-primary-container">
       <div class="flex items-center gap-2 font-body text-body">
        2022-23
        <span class="text-text-secondary text-xs">
         (Y3)
        </span>
       </div>
      </td>
      <td class="px-card-padding py-stack text-right">
       15,00,000
      </td>
      <td class="px-card-padding py-stack text-right text-error">
       5,00,000
      </td>
      <td class="px-card-padding py-stack text-right">
       15,00,000
      </td>
      <td class="px-card-padding py-stack text-right font-medium">
       10,00,000
      </td>
     </tr>
     <!-- Year 4 (Current) -->
     <tr class="bg-surface-container-lowest hover:bg-metric-amber/30 transition-colors group relative z-10 shadow-[0_0_8px_rgba(146,64,14,0.05)] border-y border-primary-container/20">
      <td class="px-card-padding py-stack text-left border-l-[3px] border-primary-container bg-metric-amber/10">
       <div class="flex items-center gap-2 font-body text-body font-semibold text-primary-container">
        2023-24
        <span class="text-primary-container/70 text-xs">
         (Y4 - Current)
        </span>
       </div>
      </td>
      <td class="px-card-padding py-stack text-right font-semibold bg-metric-amber/10">
       10,00,000
      </td>
      <td class="px-card-padding py-stack text-right text-error font-semibold bg-metric-amber/10">
       5,00,000
      </td>
      <td class="px-card-padding py-stack text-right font-semibold bg-metric-amber/10">
       20,00,000
      </td>
      <td class="px-card-padding py-stack text-right font-bold text-primary-container bg-metric-amber/10 text-[15px]">
       5,00,000
      </td>
     </tr>
     <!-- Year 5 (Future) -->
     <tr class="hover:bg-metric-amber/30 transition-colors group opacity-70 bg-surface-subtle">
      <td class="px-card-padding py-stack text-left border-l-[3px] border-transparent group-hover:border-primary-container">
       <div class="flex items-center gap-2 font-body text-body">
        2024-25
        <span class="text-text-secondary text-xs">
         (Y5)
        </span>
       </div>
      </td>
      <td class="px-card-padding py-stack text-right">
       5,00,000
      </td>
      <td class="px-card-padding py-stack text-right text-error/70">
       5,00,000
      </td>
      <td class="px-card-padding py-stack text-right">
       25,00,000
      </td>
      <td class="px-card-padding py-stack text-right font-medium">
       0
      </td>
     </tr>
    </tbody>
    <tfoot class="bg-surface-variant/30 border-t-2 border-border-default">
     <tr>
      <td class="px-card-padding py-stack text-right font-body text-body font-semibold text-text-secondary" colspan="2">
       Totals
      </td>
      <td class="px-card-padding py-stack text-right font-mono-data text-mono-data font-bold text-error">
       25,00,000
      </td>
      <td class="px-card-padding py-stack text-right font-mono-data text-mono-data font-bold text-text-primary">
       25,00,000
      </td>
      <td class="px-card-padding py-stack text-right font-mono-data text-mono-data font-bold text-text-primary">
       0
      </td>
     </tr>
    </tfoot>
   </table>
  </div>
 </div>
</main>

</template>

<script setup lang="ts">
// Autogenerated from depreciation_ledger_view
</script>
