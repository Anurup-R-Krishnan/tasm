<template>
<main class="flex-1 flex flex-col min-h-screen bg-canvas">
 <!-- TopNavBar -->
 <header class="bg-[#FAF9F6] dark:bg-stone-950 docked full-width top-0 h-[60px] border-b border-[#EEEBE4] dark:border-stone-800 shadow-sm flex justify-between items-center px-8 w-full sticky top-0 z-40">
  <!-- Left Side: Brand/Search -->
  <div class="flex items-center gap-6">
   <div class="text-lg font-bold text-stone-900 dark:text-stone-50 tracking-tight">
    Technopark AMS
   </div>
   <!-- Search Bar -->
   <div class="relative w-64">
    <span class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-stone-400 text-[18px]">
     search
    </span>
    <input class="w-full pl-9 pr-4 py-1.5 bg-surface border border-[#EEEBE4] rounded-md font-['Inter'] text-sm tracking-tight text-stone-900 focus:outline-none focus:border-[#92400E] focus:ring-1 focus:ring-[#92400E] transition-colors" placeholder="Search assets, records..." type="text"/>
   </div>
  </div>
  <!-- Right Side: Actions -->
  <div class="flex items-center gap-4">
   <button class="font-['Inter'] text-sm tracking-tight text-[#92400E] dark:text-amber-500 font-medium hover:bg-stone-100 dark:hover:bg-stone-900 transition-colors px-3 py-1.5 rounded-md cursor-pointer active:opacity-80 transition-opacity">
    Audit Logs
   </button>
   <div class="flex items-center gap-1 border-l border-[#EEEBE4] pl-4 ml-2">
    <button class="w-8 h-8 flex items-center justify-center text-stone-600 dark:text-stone-400 hover:bg-stone-100 dark:hover:bg-stone-900 transition-colors rounded-full cursor-pointer active:opacity-80 transition-opacity">
     <span class="material-symbols-outlined text-[20px]">
      notifications
     </span>
    </button>
    <button class="w-8 h-8 flex items-center justify-center text-stone-600 dark:text-stone-400 hover:bg-stone-100 dark:hover:bg-stone-900 transition-colors rounded-full cursor-pointer active:opacity-80 transition-opacity">
     <span class="material-symbols-outlined text-[20px]">
      help
     </span>
    </button>
    <button class="w-8 h-8 flex items-center justify-center text-stone-600 dark:text-stone-400 hover:bg-stone-100 dark:hover:bg-stone-900 transition-colors rounded-full cursor-pointer active:opacity-80 transition-opacity">
     <span class="material-symbols-outlined text-[20px]">
      settings
     </span>
    </button>
    <button class="ml-2 w-8 h-8 rounded-full bg-stone-200 border border-stone-300 overflow-hidden cursor-pointer active:opacity-80 transition-opacity">
     <img alt="Administrator profile picture" class="w-full h-full object-cover" data-alt="Default administrator avatar with initials" src="https://lh3.googleusercontent.com/aida-public/AB6AXuBvJxl17Wuzfzja-KV7WmvViZpvXWRD0RhGWxOK4p5uCV5AGrDKXTAmjOxZDNHLSbnEydjpgU7aRBai9lVr_tWIul5zgHSr_lqz3yRNawomCxBO1xE-ySDb3LB_pLT9kKhLQ-N2ozG4RZLY8XJuqDgcKH2os9BacffKIxCo1J9ZzbFJzDTSCCi3cq25oB80Z9SmHY2lnTuo6gRVvf9KJBbrblLY4V0Q_-THuOntdC9b-plLtRtDqNcT7FWJM2imVK7bpuqBUi8VpM0o"/>
    </button>
   </div>
  </div>
 </header>
 <!-- Scrollable Content Canvas -->
 <div class="flex-1 overflow-y-auto p-page-margin">
  <div class="max-w-[1400px] mx-auto space-y-section-gap">
   <!-- Page Header: Asset Identity -->
   <div class="flex flex-col md:flex-row md:items-end justify-between gap-4">
    <div>
     <div class="flex items-center gap-2 mb-2">
      <span class="font-metadata text-metadata text-text-secondary">
       Financial Ledgers
      </span>
      <span class="material-symbols-outlined text-[14px] text-text-secondary">
       chevron_right
      </span>
      <span class="font-metadata text-metadata text-primary-container font-medium">
       Asset Depreciation
      </span>
     </div>
     <h2 class="font-h1 text-h1 text-text-primary mb-1">
      Laptop Alpha #TAG-1001
     </h2>
     <div class="flex items-center gap-3">
      <span class="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium bg-surface-variant text-on-surface-variant border border-outline-variant">
       Category: IT Equipment
      </span>
      <span class="inline-flex items-center gap-1 text-xs font-medium text-status-in-stock">
       <span class="w-2 h-2 rounded-full bg-status-in-stock">
       </span>
       Active Asset
      </span>
     </div>
    </div>
    <div class="flex items-center gap-3">
     <button class="px-4 py-2 bg-surface text-text-primary border border-border-default rounded-lg font-h3 text-h3 hover:bg-surface-subtle transition-colors flex items-center gap-2 shadow-sm -translate-y-0 hover:-translate-y-[1px]">
      <span class="material-symbols-outlined text-[18px]">
       print
      </span>
      Print
     </button>
     <button class="px-4 py-2 bg-[#1C1917] text-white rounded-lg font-h3 text-h3 hover:bg-stone-800 transition-colors flex items-center gap-2 shadow-sm -translate-y-0 hover:-translate-y-[1px]">
      <span class="material-symbols-outlined text-[18px]">
       download
      </span>
      Export CSV
     </button>
    </div>
   </div>
   <!-- KPI Cards Grid -->
   <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
    <!-- KPI 1 -->
    <div class="bg-surface border border-border-default rounded-xl p-card-padding shadow-sm hover:-translate-y-[2px] transition-transform duration-200">
     <div class="flex items-center justify-between mb-4">
      <span class="font-metadata text-metadata text-text-secondary uppercase tracking-wider">
       Total Cost
      </span>
      <span class="material-symbols-outlined text-text-secondary bg-surface-subtle p-1.5 rounded-md">
       payments
      </span>
     </div>
     <div class="font-kpi-number text-kpi-number text-text-primary">
      ₹1,20,000
     </div>
     <div class="font-metadata text-metadata text-text-secondary mt-1">
      Initial acquisition value
     </div>
    </div>
    <!-- KPI 2 -->
    <div class="bg-surface border border-border-default rounded-xl p-card-padding shadow-sm hover:-translate-y-[2px] transition-transform duration-200">
     <div class="flex items-center justify-between mb-4">
      <span class="font-metadata text-metadata text-text-secondary uppercase tracking-wider">
       Accumulated Depreciation
      </span>
      <span class="material-symbols-outlined text-text-secondary bg-surface-subtle p-1.5 rounded-md">
       trending_down
      </span>
     </div>
     <div class="font-kpi-number text-kpi-number text-error">
      ₹45,000
     </div>
     <div class="font-metadata text-metadata text-text-secondary mt-1">
      Total written off to date
     </div>
    </div>
    <!-- KPI 3 (Sage Fill) -->
    <div class="bg-metric-sage border border-status-in-stock/20 rounded-xl p-card-padding shadow-sm hover:-translate-y-[2px] transition-transform duration-200 relative overflow-hidden">
     <div class="absolute right-0 top-0 opacity-10 pointer-events-none">
      <span class="material-symbols-outlined text-[120px] -mt-4 -mr-4" style="font-variation-settings: 'FILL' 1;">
       account_balance
      </span>
     </div>
     <div class="flex items-center justify-between mb-4 relative z-10">
      <span class="font-metadata text-metadata text-status-in-stock uppercase tracking-wider font-semibold">
       Net Book Value
      </span>
      <span class="material-symbols-outlined text-status-in-stock bg-white/50 p-1.5 rounded-md">
       account_balance_wallet
      </span>
     </div>
     <div class="font-kpi-number text-kpi-number text-status-in-stock relative z-10">
      ₹75,000
     </div>
     <div class="font-metadata text-metadata text-status-in-stock/80 mt-1 relative z-10">
      Current carrying value
     </div>
    </div>
   </div>
   <!-- Dense Table Section -->
   <div class="bg-surface border border-border-default rounded-xl shadow-sm overflow-hidden flex flex-col">
    <div class="px-card-padding py-4 border-b border-border-default bg-surface flex items-center justify-between">
     <h3 class="font-h2 text-h2 text-text-primary flex items-center gap-2">
      <span class="material-symbols-outlined text-primary-container">
       calendar_month
      </span>
      Depreciation Schedule
     </h3>
     <div class="flex items-center gap-2">
      <span class="font-metadata text-metadata text-text-secondary">
       Method:
      </span>
      <span class="font-metadata text-metadata font-medium text-text-primary bg-surface-subtle px-2 py-1 rounded border border-border-default">
       Straight Line (SLM)
      </span>
     </div>
    </div>
    <div class="overflow-x-auto">
     <table class="w-full text-left border-collapse">
      <thead class="bg-surface-subtle border-b border-border-default">
       <tr>
        <th class="font-table-header text-table-header text-text-secondary uppercase px-6 py-3 sticky left-0 bg-surface-subtle z-10">
         Year
        </th>
        <th class="font-table-header text-table-header text-text-secondary uppercase px-6 py-3 text-right">
         Beginning Value
        </th>
        <th class="font-table-header text-table-header text-text-secondary uppercase px-6 py-3 text-right">
         Depreciation Rate (%)
        </th>
        <th class="font-table-header text-table-header text-text-secondary uppercase px-6 py-3 text-right">
         Depreciation Expense
        </th>
        <th class="font-table-header text-table-header text-text-secondary uppercase px-6 py-3 text-right">
         Ending Value
        </th>
       </tr>
      </thead>
      <tbody class="divide-y divide-border-default">
       <!-- Row 1 -->
       <tr class="bg-surface hover:bg-metric-amber transition-colors group">
        <td class="font-body text-body text-text-primary px-6 py-4 sticky left-0 bg-surface group-hover:bg-metric-amber transition-colors font-medium">
         2021-22
        </td>
        <td class="font-mono-data text-mono-data text-text-primary px-6 py-4 text-right">
         ₹1,20,000
        </td>
        <td class="font-mono-data text-mono-data text-text-secondary px-6 py-4 text-right">
         12.50%
        </td>
        <td class="font-mono-data text-mono-data text-error px-6 py-4 text-right">
         -₹15,000
        </td>
        <td class="font-mono-data text-mono-data text-text-primary px-6 py-4 text-right font-medium">
         ₹1,05,000
        </td>
       </tr>
       <!-- Row 2 -->
       <tr class="bg-surface-subtle hover:bg-metric-amber transition-colors group">
        <td class="font-body text-body text-text-primary px-6 py-4 sticky left-0 bg-surface-subtle group-hover:bg-metric-amber transition-colors font-medium">
         2022-23
        </td>
        <td class="font-mono-data text-mono-data text-text-primary px-6 py-4 text-right">
         ₹1,05,000
        </td>
        <td class="font-mono-data text-mono-data text-text-secondary px-6 py-4 text-right">
         12.50%
        </td>
        <td class="font-mono-data text-mono-data text-error px-6 py-4 text-right">
         -₹15,000
        </td>
        <td class="font-mono-data text-mono-data text-text-primary px-6 py-4 text-right font-medium">
         ₹90,000
        </td>
       </tr>
       <!-- Row 3 -->
       <tr class="bg-surface hover:bg-metric-amber transition-colors group">
        <td class="font-body text-body text-text-primary px-6 py-4 sticky left-0 bg-surface group-hover:bg-metric-amber transition-colors font-medium">
         2023-24
        </td>
        <td class="font-mono-data text-mono-data text-text-primary px-6 py-4 text-right">
         ₹90,000
        </td>
        <td class="font-mono-data text-mono-data text-text-secondary px-6 py-4 text-right">
         12.50%
        </td>
        <td class="font-mono-data text-mono-data text-error px-6 py-4 text-right">
         -₹15,000
        </td>
        <td class="font-mono-data text-mono-data text-text-primary px-6 py-4 text-right font-medium">
         ₹75,000
        </td>
       </tr>
       <!-- Row 4 (Current) -->
       <tr class="bg-metric-amber/30 hover:bg-metric-amber transition-colors group border-l-4 border-l-primary-container">
        <td class="font-body text-body text-primary-container px-6 py-4 sticky left-0 bg-metric-amber/30 group-hover:bg-metric-amber transition-colors font-semibold flex items-center gap-2">
         2024-25
         <span class="inline-flex items-center px-1.5 py-0.5 rounded text-[10px] font-bold bg-primary-container text-white uppercase tracking-wider">
          Current
         </span>
        </td>
        <td class="font-mono-data text-mono-data text-text-primary px-6 py-4 text-right">
         ₹75,000
        </td>
        <td class="font-mono-data text-mono-data text-text-secondary px-6 py-4 text-right">
         12.50%
        </td>
        <td class="font-mono-data text-mono-data text-text-secondary px-6 py-4 text-right italic">
         Projected
        </td>
        <td class="font-mono-data text-mono-data text-text-secondary px-6 py-4 text-right italic">
         ₹60,000
        </td>
       </tr>
       <!-- Row 5 (Future) -->
       <tr class="bg-surface-subtle hover:bg-metric-amber transition-colors group opacity-70">
        <td class="font-body text-body text-text-secondary px-6 py-4 sticky left-0 bg-surface-subtle group-hover:bg-metric-amber transition-colors font-medium">
         2025-26
        </td>
        <td class="font-mono-data text-mono-data text-text-secondary px-6 py-4 text-right">
         ₹60,000
        </td>
        <td class="font-mono-data text-mono-data text-text-secondary px-6 py-4 text-right">
         12.50%
        </td>
        <td class="font-mono-data text-mono-data text-text-secondary px-6 py-4 text-right italic">
         Projected
        </td>
        <td class="font-mono-data text-mono-data text-text-secondary px-6 py-4 text-right italic">
         ₹45,000
        </td>
       </tr>
      </tbody>
     </table>
    </div>
    <!-- Table Footer Summary -->
    <div class="px-6 py-3 border-t border-border-default bg-surface flex justify-end items-center gap-6">
     <div class="flex flex-col items-end">
      <span class="font-metadata text-metadata text-text-secondary">
       YTD Depreciation
      </span>
      <span class="font-mono-data text-mono-data text-error font-medium">
       -₹45,000
      </span>
     </div>
    </div>
   </div>
  </div>
 </div>
</main>

</template>

<script setup lang="ts">
// Autogenerated from asset_depreciation_ledger
</script>
