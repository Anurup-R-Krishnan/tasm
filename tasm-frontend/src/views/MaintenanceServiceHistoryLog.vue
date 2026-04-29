<template>
<main class="flex-1 md:ml-[248px] p-page-margin max-w-[1400px] w-full mx-auto">
 <div class="mb-section-gap flex justify-between items-end">
  <div>
   <h1 class="font-h1 text-h1 text-text-primary mb-1">
    Maintenance History
   </h1>
   <p class="font-body text-body text-text-secondary">
    Track work orders, service logs, and maintenance costs.
   </p>
  </div>
  <div class="flex gap-3">
   <button class="px-4 py-2 bg-surface border border-border-default rounded-lg text-text-primary font-medium hover:-translate-y-0.5 transition-transform shadow-sm flex items-center gap-2">
    <span class="material-symbols-outlined text-[18px]">
     download
    </span>
    Export CSV
   </button>
  </div>
 </div>
 <!-- KPI Cards -->
 <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-section-gap">
  <div class="bg-surface rounded-xl border border-border-default p-card-padding shadow-sm hover:-translate-y-0.5 transition-transform">
   <div class="flex justify-between items-start mb-4">
    <p class="font-metadata text-metadata text-text-secondary uppercase tracking-wider">
     Total Maint. Spend YTD
    </p>
    <div class="p-2 bg-metric-amber rounded-lg text-primary-container">
     <span class="material-symbols-outlined">
      payments
     </span>
    </div>
   </div>
   <p class="font-kpi-number text-kpi-number text-text-primary font-mono-data mb-1">
    ₹14,52,400
   </p>
   <p class="font-metadata text-metadata text-status-critical flex items-center gap-1">
    <span class="material-symbols-outlined text-[14px]">
     trending_up
    </span>
    +12.5% vs last year
   </p>
  </div>
  <div class="bg-surface rounded-xl border border-border-default p-card-padding shadow-sm hover:-translate-y-0.5 transition-transform">
   <div class="flex justify-between items-start mb-4">
    <p class="font-metadata text-metadata text-text-secondary uppercase tracking-wider">
     Open Work Orders
    </p>
    <div class="p-2 bg-metric-lavender rounded-lg text-tertiary-container">
     <span class="material-symbols-outlined">
      build_circle
     </span>
    </div>
   </div>
   <p class="font-kpi-number text-kpi-number text-text-primary">
    24
   </p>
   <p class="font-metadata text-metadata text-text-secondary flex items-center gap-1">
    <span class="material-symbols-outlined text-[14px]">
     schedule
    </span>
    5 critical priority
   </p>
  </div>
  <div class="bg-surface rounded-xl border border-border-default p-card-padding shadow-sm hover:-translate-y-0.5 transition-transform">
   <div class="flex justify-between items-start mb-4">
    <p class="font-metadata text-metadata text-text-secondary uppercase tracking-wider">
     Avg Resolution Time
    </p>
    <div class="p-2 bg-metric-sage rounded-lg text-status-in-stock">
     <span class="material-symbols-outlined">
      timer
     </span>
    </div>
   </div>
   <p class="font-kpi-number text-kpi-number text-text-primary">
    3.2 Days
   </p>
   <p class="font-metadata text-metadata text-status-in-stock flex items-center gap-1">
    <span class="material-symbols-outlined text-[14px]">
     trending_down
    </span>
    -0.5 days vs last month
   </p>
  </div>
 </div>
 <!-- Content Grid: Table & Form Split -->
 <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
  <!-- Table Section (Takes 2/3 width on large screens) -->
  <div class="lg:col-span-2 bg-surface rounded-xl border border-border-default shadow-sm overflow-hidden flex flex-col h-full">
   <div class="p-4 border-b border-border-default bg-surface-subtle flex justify-between items-center">
    <h3 class="font-h3 text-h3 text-text-primary">
     Service Logs
    </h3>
    <div class="flex gap-2">
     <div class="relative">
      <span class="material-symbols-outlined absolute left-2.5 top-2 text-stone-400 text-[18px]">
       filter_list
      </span>
      <select class="pl-9 pr-8 py-1.5 bg-white border border-border-default rounded-md text-sm text-text-primary appearance-none focus:ring-1 focus:ring-amber-600 focus:border-amber-600 outline-none cursor-pointer">
       <option>
        All Types
       </option>
       <option>
        Preventive
       </option>
       <option>
        Corrective
       </option>
      </select>
     </div>
    </div>
   </div>
   <div class="overflow-x-auto flex-1">
    <table class="w-full text-left border-collapse">
     <thead>
      <tr class="border-b border-border-default bg-surface">
       <th class="p-4 font-table-header text-table-header text-text-secondary uppercase">
        WO ID
       </th>
       <th class="p-4 font-table-header text-table-header text-text-secondary uppercase">
        Date
       </th>
       <th class="p-4 font-table-header text-table-header text-text-secondary uppercase">
        Asset &amp; Type
       </th>
       <th class="p-4 font-table-header text-table-header text-text-secondary uppercase">
        Technician
       </th>
       <th class="p-4 font-table-header text-table-header text-text-secondary uppercase text-right">
        Cost (₹)
       </th>
       <th class="p-4 font-table-header text-table-header text-text-secondary uppercase text-center">
        Status
       </th>
      </tr>
     </thead>
     <tbody class="font-body text-body">
      <tr class="border-b border-border-default hover:bg-metric-amber/20 transition-colors cursor-pointer group">
       <td class="p-4 font-mono-data text-mono-data text-text-primary group-hover:text-amber-800">
        WO-2023-1042
       </td>
       <td class="p-4 text-text-secondary">
        24 Oct 2023
       </td>
       <td class="p-4">
        <div class="text-text-primary font-medium">
         HVAC Unit B-3
        </div>
        <div class="text-xs text-text-secondary mt-0.5">
         Preventive Maintenance
        </div>
       </td>
       <td class="p-4 text-text-secondary">
        Ramesh K.
       </td>
       <td class="p-4 font-mono-data text-mono-data text-right text-text-primary">
        12,500.00
       </td>
       <td class="p-4 text-center">
        <span class="inline-flex items-center px-2 py-1 rounded-md bg-metric-sage text-status-in-stock text-xs font-medium">
         Completed
        </span>
       </td>
      </tr>
      <tr class="border-b border-border-default hover:bg-metric-amber/20 transition-colors cursor-pointer group">
       <td class="p-4 font-mono-data text-mono-data text-text-primary group-hover:text-amber-800">
        WO-2023-1041
       </td>
       <td class="p-4 text-text-secondary">
        22 Oct 2023
       </td>
       <td class="p-4">
        <div class="text-text-primary font-medium">
         Server Rack 04-A
        </div>
        <div class="text-xs text-text-secondary mt-0.5">
         Hardware Replacement
        </div>
       </td>
       <td class="p-4 text-text-secondary">
        TechSys Vendor
       </td>
       <td class="p-4 font-mono-data text-mono-data text-right text-text-primary">
        45,000.00
       </td>
       <td class="p-4 text-center">
        <span class="inline-flex items-center px-2 py-1 rounded-md bg-metric-amber text-primary-container text-xs font-medium">
         Pending Parts
        </span>
       </td>
      </tr>
      <tr class="border-b border-border-default hover:bg-metric-amber/20 transition-colors cursor-pointer group">
       <td class="p-4 font-mono-data text-mono-data text-text-primary group-hover:text-amber-800">
        WO-2023-1038
       </td>
       <td class="p-4 text-text-secondary">
        18 Oct 2023
       </td>
       <td class="p-4">
        <div class="text-text-primary font-medium">
         Generator 2 (Pamba)
        </div>
        <div class="text-xs text-text-secondary mt-0.5">
         Emergency Repair
        </div>
       </td>
       <td class="p-4 text-text-secondary">
        Suresh M.
       </td>
       <td class="p-4 font-mono-data text-mono-data text-right text-text-primary">
        1,12,000.00
       </td>
       <td class="p-4 text-center">
        <span class="inline-flex items-center px-2 py-1 rounded-md bg-metric-sage text-status-in-stock text-xs font-medium">
         Completed
        </span>
       </td>
      </tr>
      <tr class="border-b border-border-default hover:bg-metric-amber/20 transition-colors cursor-pointer group">
       <td class="p-4 font-mono-data text-mono-data text-text-primary group-hover:text-amber-800">
        WO-2023-1035
       </td>
       <td class="p-4 text-text-secondary">
        15 Oct 2023
       </td>
       <td class="p-4">
        <div class="text-text-primary font-medium">
         Elevator L4 (Nila)
        </div>
        <div class="text-xs text-text-secondary mt-0.5">
         Routine Inspection
        </div>
       </td>
       <td class="p-4 text-text-secondary">
        Otis Vendor
       </td>
       <td class="p-4 font-mono-data text-mono-data text-right text-text-primary">
        8,500.00
       </td>
       <td class="p-4 text-center">
        <span class="inline-flex items-center px-2 py-1 rounded-md bg-metric-sage text-status-in-stock text-xs font-medium">
         Completed
        </span>
       </td>
      </tr>
      <tr class="border-b border-border-default hover:bg-metric-amber/20 transition-colors cursor-pointer group">
       <td class="p-4 font-mono-data text-mono-data text-text-primary group-hover:text-amber-800">
        WO-2023-1031
       </td>
       <td class="p-4 text-text-secondary">
        10 Oct 2023
       </td>
       <td class="p-4">
        <div class="text-text-primary font-medium">
         CCTV Camera Ext-09
        </div>
        <div class="text-xs text-text-secondary mt-0.5">
         Corrective - No Signal
        </div>
       </td>
       <td class="p-4 text-text-secondary">
        Arun J.
       </td>
       <td class="p-4 font-mono-data text-mono-data text-right text-text-primary">
        3,200.00
       </td>
       <td class="p-4 text-center">
        <span class="inline-flex items-center px-2 py-1 rounded-md bg-metric-sage text-status-in-stock text-xs font-medium">
         Completed
        </span>
       </td>
      </tr>
     </tbody>
    </table>
   </div>
   <div class="p-4 border-t border-border-default bg-surface-subtle flex justify-between items-center text-sm text-text-secondary">
    <span>
     Showing 1-5 of 142 logs
    </span>
    <div class="flex gap-1">
     <button class="p-1 rounded hover:bg-stone-200 disabled:opacity-50" disabled="">
      <span class="material-symbols-outlined text-[18px]">
       chevron_left
      </span>
     </button>
     <button class="p-1 rounded hover:bg-stone-200">
      <span class="material-symbols-outlined text-[18px]">
       chevron_right
      </span>
     </button>
    </div>
   </div>
  </div>
  <!-- Add Log Form Section (Takes 1/3 width on large screens) -->
  <div class="lg:col-span-1 bg-surface rounded-xl border border-border-default shadow-sm h-full flex flex-col">
   <div class="p-4 border-b border-border-default bg-surface-subtle">
    <h3 class="font-h3 text-h3 text-text-primary flex items-center gap-2">
     <span class="material-symbols-outlined text-[20px] text-amber-700">
      post_add
     </span>
     Quick Add Log
    </h3>
   </div>
   <form class="p-6 flex-1 flex flex-col gap-5">
    <div class="space-y-1">
     <label class="font-table-header text-table-header text-text-secondary uppercase">
      Asset ID / Name
     </label>
     <div class="relative">
      <input class="w-full px-3 py-2 border border-border-default rounded-lg focus:ring-1 focus:ring-amber-600 focus:border-amber-600 outline-none text-sm text-text-primary" placeholder="e.g. HVAC-B3" type="text"/>
      <span class="material-symbols-outlined absolute right-3 top-2 text-stone-400 text-[20px]">
       search
      </span>
     </div>
    </div>
    <div class="grid grid-cols-2 gap-4">
     <div class="space-y-1">
      <label class="font-table-header text-table-header text-text-secondary uppercase">
       Service Date
      </label>
      <input class="w-full px-3 py-2 border border-border-default rounded-lg focus:ring-1 focus:ring-amber-600 focus:border-amber-600 outline-none text-sm text-text-primary" type="date"/>
     </div>
     <div class="space-y-1">
      <label class="font-table-header text-table-header text-text-secondary uppercase">
       Service Type
      </label>
      <select class="w-full px-3 py-2 border border-border-default rounded-lg focus:ring-1 focus:ring-amber-600 focus:border-amber-600 outline-none text-sm text-text-primary appearance-none bg-white">
       <option>
        Preventive
       </option>
       <option>
        Corrective
       </option>
       <option>
        Emergency
       </option>
      </select>
     </div>
    </div>
    <div class="space-y-1">
     <label class="font-table-header text-table-header text-text-secondary uppercase">
      Technician / Vendor
     </label>
     <input class="w-full px-3 py-2 border border-border-default rounded-lg focus:ring-1 focus:ring-amber-600 focus:border-amber-600 outline-none text-sm text-text-primary" placeholder="Name or Company" type="text"/>
    </div>
    <div class="space-y-1">
     <label class="font-table-header text-table-header text-text-secondary uppercase">
      Total Cost (₹)
     </label>
     <div class="relative">
      <span class="absolute left-3 top-2 text-text-secondary font-mono-data">
       ₹
      </span>
      <input class="w-full pl-8 pr-3 py-2 border border-border-default rounded-lg focus:ring-1 focus:ring-amber-600 focus:border-amber-600 outline-none text-sm font-mono-data text-text-primary" placeholder="0.00" type="number"/>
     </div>
    </div>
    <div class="space-y-1">
     <label class="font-table-header text-table-header text-text-secondary uppercase">
      Notes
     </label>
     <textarea class="w-full px-3 py-2 border border-border-default rounded-lg focus:ring-1 focus:ring-amber-600 focus:border-amber-600 outline-none text-sm text-text-primary h-24 resize-none" placeholder="Details of work performed..."></textarea>
    </div>
    <div class="mt-auto pt-4 flex gap-3">
     <button class="flex-1 px-4 py-2 bg-surface border border-border-default rounded-lg text-text-primary font-medium hover:bg-stone-50 transition-colors" type="button">
      Clear
     </button>
     <button class="flex-1 px-4 py-2 bg-[#1C1917] rounded-lg text-white font-medium hover:bg-black transition-colors shadow-sm" type="submit">
      Save Log
     </button>
    </div>
   </form>
  </div>
 </div>
</main>

</template>

<script setup lang="ts">
// Autogenerated from maintenance_service_history_log
</script>
