<template>
<main class="md:ml-[248px] min-h-screen">
 <div class="max-w-[1400px] mx-auto p-page-margin space-y-section-gap">
  <!-- Header & Filter Bar -->
  <section class="flex flex-col xl:flex-row justify-between items-start xl:items-center gap-4">
   <div>
    <h1 class="font-h1 text-h1 text-text-primary mb-1">
     Financial Overview
    </h1>
    <p class="font-body text-body text-text-secondary">
     Asset valuation and depreciation metrics for current fiscal year.
    </p>
   </div>
   <div class="flex flex-wrap items-center gap-inline bg-surface border border-border-default p-2 rounded-lg shadow-[0_4px_4px_rgba(0,0,0,0.02)]">
    <div class="flex items-center gap-2 px-3 py-1.5 border-r border-border-default">
     <span class="material-symbols-outlined text-text-secondary text-sm">
      calendar_month
     </span>
     <select class="bg-transparent border-none text-body font-body text-text-primary focus:ring-0 p-0 pr-6 text-sm cursor-pointer">
      <option>
       FY 2023-24
      </option>
      <option>
       FY 2022-23
      </option>
     </select>
    </div>
    <div class="flex items-center gap-2 px-3 py-1.5 border-r border-border-default">
     <span class="material-symbols-outlined text-text-secondary text-sm">
      domain
     </span>
     <select class="bg-transparent border-none text-body font-body text-text-primary focus:ring-0 p-0 pr-6 text-sm cursor-pointer">
      <option>
       All Departments
      </option>
      <option>
       IT Infrastructure
      </option>
      <option>
       Facilities
      </option>
     </select>
    </div>
    <div class="flex items-center gap-2 px-3 py-1.5">
     <span class="material-symbols-outlined text-text-secondary text-sm">
      category
     </span>
     <select class="bg-transparent border-none text-body font-body text-text-primary focus:ring-0 p-0 pr-6 text-sm cursor-pointer">
      <option>
       All Categories
      </option>
      <option>
       Electronics
      </option>
      <option>
       Furniture
      </option>
     </select>
    </div>
   </div>
  </section>
  <!-- KPI Cards Bento Grid -->
  <section class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
   <!-- Total Assets (Amber) -->
   <div class="bg-metric-amber/30 border border-metric-amber/50 rounded-[16px] p-card-padding hover:-translate-y-[2px] transition-transform duration-200 cursor-pointer">
    <div class="flex justify-between items-start mb-4">
     <span class="font-metadata text-metadata text-text-secondary uppercase tracking-wider">
      Total Assets
     </span>
     <span class="material-symbols-outlined text-primary">
      inventory
     </span>
    </div>
    <div class="font-kpi-number text-kpi-number text-text-primary">
     12,458
    </div>
    <div class="mt-2 flex items-center gap-1 font-metadata text-metadata text-status-in-stock">
     <span class="material-symbols-outlined text-[14px]">
      trending_up
     </span>
     <span>
      +4.2% from last year
     </span>
    </div>
   </div>
   <!-- Gross Purchase Value (Lavender) -->
   <div class="bg-metric-lavender/40 border border-metric-lavender/60 rounded-[16px] p-card-padding hover:-translate-y-[2px] transition-transform duration-200 cursor-pointer">
    <div class="flex justify-between items-start mb-4">
     <span class="font-metadata text-metadata text-text-secondary uppercase tracking-wider">
      Gross Purchase Value
     </span>
     <span class="material-symbols-outlined text-tertiary">
      payments
     </span>
    </div>
    <div class="font-mono-data text-[24px] font-semibold text-text-primary">
     ₹ 45,20,50,000
    </div>
    <div class="mt-2 font-metadata text-metadata text-text-secondary">
     Initial capital expenditure
    </div>
   </div>
   <!-- Net Book Value (Sage) -->
   <div class="bg-metric-sage/40 border border-metric-sage/60 rounded-[16px] p-card-padding hover:-translate-y-[2px] transition-transform duration-200 cursor-pointer">
    <div class="flex justify-between items-start mb-4">
     <span class="font-metadata text-metadata text-text-secondary uppercase tracking-wider">
      Net Book Value
     </span>
     <span class="material-symbols-outlined text-status-in-stock">
      account_balance
     </span>
    </div>
    <div class="font-mono-data text-[24px] font-semibold text-text-primary">
     ₹ 28,15,40,000
    </div>
    <div class="mt-2 font-metadata text-metadata text-text-secondary">
     Current depreciated value
    </div>
   </div>
   <!-- Total Depreciation (Peach/Primary Fixed) -->
   <div class="bg-primary-fixed/30 border border-primary-fixed/50 rounded-[16px] p-card-padding hover:-translate-y-[2px] transition-transform duration-200 cursor-pointer">
    <div class="flex justify-between items-start mb-4">
     <span class="font-metadata text-metadata text-text-secondary uppercase tracking-wider">
      Total Depreciation
     </span>
     <span class="material-symbols-outlined text-primary-container">
      trending_down
     </span>
    </div>
    <div class="font-mono-data text-[24px] font-semibold text-text-primary">
     ₹ 17,05,10,000
    </div>
    <div class="mt-2 flex items-center gap-1 font-metadata text-metadata text-status-critical">
     <span class="material-symbols-outlined text-[14px]">
      warning
     </span>
     <span>
      Approaching end-of-life for 12%
     </span>
    </div>
   </div>
  </section>
  <!-- Charts Section -->
  <section class="grid grid-cols-1 lg:grid-cols-3 gap-6">
   <!-- Area Chart Placeholder -->
   <div class="lg:col-span-2 bg-surface border border-border-default rounded-[16px] shadow-[0_4px_4px_rgba(0,0,0,0.02)] p-card-padding flex flex-col">
    <div class="flex justify-between items-center mb-6">
     <div>
      <h3 class="font-h3 text-h3 text-text-primary">
       Asset Value vs. Depreciation
      </h3>
      <p class="font-metadata text-metadata text-text-secondary mt-1">
       5-Year Historical Trend
      </p>
     </div>
     <button class="p-2 border border-border-default rounded-md hover:bg-surface-subtle transition-colors">
      <span class="material-symbols-outlined text-text-secondary text-sm">
       more_horiz
      </span>
     </button>
    </div>
    <div class="flex-1 min-h-[300px] rounded-lg border border-border-default/50 bg-surface-subtle/50 flex items-center justify-center relative overflow-hidden">
     <!-- Simulated Chart Visual -->
     <div class="absolute inset-0 opacity-20" style="background: linear-gradient(180deg, #98cbff 0%, transparent 100%); clip-path: polygon(0 40%, 20% 30%, 40% 50%, 60% 20%, 80% 40%, 100% 10%, 100% 100%, 0 100%);">
     </div>
     <div class="absolute inset-0 opacity-30" style="background: linear-gradient(180deg, #ffb693 0%, transparent 100%); clip-path: polygon(0 60%, 20% 50%, 40% 70%, 60% 40%, 80% 60%, 100% 30%, 100% 100%, 0 100%);">
     </div>
     <!-- Grid lines -->
     <div class="absolute inset-0 flex flex-col justify-between py-8 px-12 pointer-events-none">
      <div class="w-full border-t border-border-default border-dashed opacity-50">
      </div>
      <div class="w-full border-t border-border-default border-dashed opacity-50">
      </div>
      <div class="w-full border-t border-border-default border-dashed opacity-50">
      </div>
      <div class="w-full border-t border-border-default border-dashed opacity-50">
      </div>
     </div>
     <div class="text-center z-10 bg-surface/80 backdrop-blur-sm p-4 rounded-lg border border-border-default">
      <span class="material-symbols-outlined text-text-secondary mb-2 block">
       monitoring
      </span>
      <span class="font-body text-body text-text-secondary">
       [Recharts AreaChart Placeholder]
      </span>
     </div>
    </div>
   </div>
   <!-- Bar Chart Placeholder -->
   <div class="bg-surface border border-border-default rounded-[16px] shadow-[0_4px_4px_rgba(0,0,0,0.02)] p-card-padding flex flex-col">
    <div class="mb-6">
     <h3 class="font-h3 text-h3 text-text-primary">
      Distribution by Category
     </h3>
     <p class="font-metadata text-metadata text-text-secondary mt-1">
      Value allocation across types
     </p>
    </div>
    <div class="flex-1 min-h-[300px] flex flex-col justify-end gap-4 pb-8 px-4 relative">
     <!-- Simulated Bar Chart -->
     <div class="w-full border-t border-border-default border-dashed opacity-50 absolute bottom-8 left-0">
     </div>
     <div class="flex items-end justify-around w-full h-[250px] z-10">
      <div class="w-12 bg-tertiary-fixed rounded-t-md h-[80%] hover:opacity-80 transition-opacity cursor-pointer group relative">
       <div class="absolute -top-8 left-1/2 -translate-x-1/2 opacity-0 group-hover:opacity-100 transition-opacity bg-text-primary text-white text-xs py-1 px-2 rounded whitespace-nowrap">
        ₹ 18Cr
       </div>
      </div>
      <div class="w-12 bg-metric-amber rounded-t-md h-[45%] hover:opacity-80 transition-opacity cursor-pointer group relative">
       <div class="absolute -top-8 left-1/2 -translate-x-1/2 opacity-0 group-hover:opacity-100 transition-opacity bg-text-primary text-white text-xs py-1 px-2 rounded whitespace-nowrap">
        ₹ 10Cr
       </div>
      </div>
      <div class="w-12 bg-metric-sage rounded-t-md h-[60%] hover:opacity-80 transition-opacity cursor-pointer group relative">
       <div class="absolute -top-8 left-1/2 -translate-x-1/2 opacity-0 group-hover:opacity-100 transition-opacity bg-text-primary text-white text-xs py-1 px-2 rounded whitespace-nowrap">
        ₹ 13Cr
       </div>
      </div>
      <div class="w-12 bg-primary-fixed rounded-t-md h-[20%] hover:opacity-80 transition-opacity cursor-pointer group relative">
       <div class="absolute -top-8 left-1/2 -translate-x-1/2 opacity-0 group-hover:opacity-100 transition-opacity bg-text-primary text-white text-xs py-1 px-2 rounded whitespace-nowrap">
        ₹ 4Cr
       </div>
      </div>
     </div>
     <div class="flex justify-around text-xs text-text-secondary font-metadata mt-2 absolute bottom-2 w-[calc(100%-2rem)]">
      <span>
       IT
      </span>
      <span>
       Furn.
      </span>
      <span>
       Facil.
      </span>
      <span>
       Veh.
      </span>
     </div>
    </div>
   </div>
  </section>
  <!-- Summary Table -->
  <section class="bg-surface border border-border-default rounded-[16px] shadow-[0_4px_4px_rgba(0,0,0,0.02)] overflow-hidden">
   <div class="p-card-padding border-b border-border-default flex justify-between items-center bg-white">
    <div>
     <h3 class="font-h3 text-h3 text-text-primary">
      Category Financial Summary
     </h3>
     <p class="font-metadata text-metadata text-text-secondary mt-1">
      Detailed breakdown of asset values and depreciation.
     </p>
    </div>
    <button class="flex items-center gap-2 px-3 py-1.5 border border-border-default rounded bg-surface text-text-primary font-body text-sm hover:bg-surface-subtle transition-colors">
     <span class="material-symbols-outlined text-[18px]">
      download
     </span>
     Export CSV
    </button>
   </div>
   <div class="overflow-x-auto">
    <table class="w-full text-left border-collapse">
     <thead>
      <tr class="bg-surface-subtle border-b border-border-default">
       <th class="font-table-header text-table-header text-text-secondary py-3 px-6">
        CATEGORY
       </th>
       <th class="font-table-header text-table-header text-text-secondary py-3 px-6 text-right">
        COUNT
       </th>
       <th class="font-table-header text-table-header text-text-secondary py-3 px-6 text-right">
        PURCHASE COST
       </th>
       <th class="font-table-header text-table-header text-text-secondary py-3 px-6 text-right">
        ACCUM. DEPRECIATION
       </th>
       <th class="font-table-header text-table-header text-text-secondary py-3 px-6 text-right">
        NET BOOK VALUE
       </th>
      </tr>
     </thead>
     <tbody class="font-body text-body text-text-primary">
      <tr class="border-b border-border-default hover:bg-metric-amber/10 transition-colors group">
       <td class="py-4 px-6 border-l-2 border-transparent group-hover:border-primary flex items-center gap-3 font-medium">
        <span class="material-symbols-outlined text-text-secondary text-[20px]">
         computer
        </span>
        IT Equipment
       </td>
       <td class="py-4 px-6 text-right font-mono-data text-mono-data">
        5,240
       </td>
       <td class="py-4 px-6 text-right font-mono-data text-mono-data">
        ₹ 18,50,00,000
       </td>
       <td class="py-4 px-6 text-right font-mono-data text-mono-data text-text-secondary">
        ₹ 9,25,00,000
       </td>
       <td class="py-4 px-6 text-right font-mono-data text-mono-data font-semibold">
        ₹ 9,25,00,000
       </td>
      </tr>
      <tr class="bg-surface-subtle/30 border-b border-border-default hover:bg-metric-amber/10 transition-colors group">
       <td class="py-4 px-6 border-l-2 border-transparent group-hover:border-primary flex items-center gap-3 font-medium">
        <span class="material-symbols-outlined text-text-secondary text-[20px]">
         chair
        </span>
        Furniture &amp; Fixtures
       </td>
       <td class="py-4 px-6 text-right font-mono-data text-mono-data">
        3,100
       </td>
       <td class="py-4 px-6 text-right font-mono-data text-mono-data">
        ₹ 10,20,00,000
       </td>
       <td class="py-4 px-6 text-right font-mono-data text-mono-data text-text-secondary">
        ₹ 4,10,00,000
       </td>
       <td class="py-4 px-6 text-right font-mono-data text-mono-data font-semibold">
        ₹ 6,10,00,000
       </td>
      </tr>
      <tr class="border-b border-border-default hover:bg-metric-amber/10 transition-colors group">
       <td class="py-4 px-6 border-l-2 border-transparent group-hover:border-primary flex items-center gap-3 font-medium">
        <span class="material-symbols-outlined text-text-secondary text-[20px]">
         hvac
        </span>
        Facilities &amp; MEP
       </td>
       <td class="py-4 px-6 text-right font-mono-data text-mono-data">
        850
       </td>
       <td class="py-4 px-6 text-right font-mono-data text-mono-data">
        ₹ 12,80,50,000
       </td>
       <td class="py-4 px-6 text-right font-mono-data text-mono-data text-text-secondary">
        ₹ 2,40,10,000
       </td>
       <td class="py-4 px-6 text-right font-mono-data text-mono-data font-semibold">
        ₹ 10,40,40,000
       </td>
      </tr>
      <tr class="bg-surface-subtle/30 hover:bg-metric-amber/10 transition-colors group">
       <td class="py-4 px-6 border-l-2 border-transparent group-hover:border-primary flex items-center gap-3 font-medium">
        <span class="material-symbols-outlined text-text-secondary text-[20px]">
         directions_car
        </span>
        Vehicles
       </td>
       <td class="py-4 px-6 text-right font-mono-data text-mono-data">
        28
       </td>
       <td class="py-4 px-6 text-right font-mono-data text-mono-data">
        ₹ 3,70,00,000
       </td>
       <td class="py-4 px-6 text-right font-mono-data text-mono-data text-text-secondary">
        ₹ 1,30,00,000
       </td>
       <td class="py-4 px-6 text-right font-mono-data text-mono-data font-semibold">
        ₹ 2,40,00,000
       </td>
      </tr>
     </tbody>
     <tfoot class="border-t-2 border-border-default bg-surface-subtle font-semibold">
      <tr>
       <td class="py-4 px-6 font-body text-body text-text-primary">
        Grand Total
       </td>
       <td class="py-4 px-6 text-right font-mono-data text-mono-data">
        9,218
       </td>
       <td class="py-4 px-6 text-right font-mono-data text-mono-data">
        ₹ 45,20,50,000
       </td>
       <td class="py-4 px-6 text-right font-mono-data text-mono-data text-text-secondary">
        ₹ 17,05,10,000
       </td>
       <td class="py-4 px-6 text-right font-mono-data text-mono-data text-primary">
        ₹ 28,15,40,000
       </td>
      </tr>
     </tfoot>
    </table>
   </div>
  </section>
 </div>
</main>

</template>

<script setup lang="ts">
// Autogenerated from reports_financials_dashboard
</script>
