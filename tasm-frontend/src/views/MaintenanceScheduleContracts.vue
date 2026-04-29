<template>
<main class="md:ml-[248px] p-page-margin max-w-[1400px] mx-auto flex flex-col gap-section-gap">
 <!-- Header Section -->
 <div class="flex justify-between items-end">
  <div>
   <h1 class="font-h1 text-h1 text-text-primary">
    Maintenance Schedule &amp; Contracts
   </h1>
   <p class="font-body text-body text-text-secondary mt-1">
    Manage upcoming service appointments and active AMC agreements.
   </p>
  </div>
  <div class="flex gap-inline">
   <button class="bg-surface text-text-primary border border-[#D6D3CE] font-h3 text-h3 px-4 py-2 rounded-md hover:-translate-y-[2px] transition-transform shadow-sm flex items-center gap-2">
    <span class="material-symbols-outlined text-[18px]">
     download
    </span>
    Export Report
   </button>
   <button class="bg-primary-container text-white font-h3 text-h3 px-4 py-2 rounded-md hover:-translate-y-[2px] transition-transform shadow-sm flex items-center gap-2">
    <span class="material-symbols-outlined text-[18px]">
     add
    </span>
    New Contract
   </button>
  </div>
 </div>
 <!-- KPI Cards Grid -->
 <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
  <!-- Card 1 (Amber) -->
  <div class="bg-metric-amber rounded-[16px] p-card-padding flex flex-col gap-2 hover:-translate-y-[2px] transition-transform cursor-pointer border border-[#EEEBE4]/50 shadow-[0_4px_4px_-2px_rgba(0,0,0,0.02)]">
   <div class="flex justify-between items-start">
    <span class="font-h3 text-h3 text-stone-800">
     Scheduled This Week
    </span>
    <span class="material-symbols-outlined text-stone-700">
     calendar_month
    </span>
   </div>
   <div class="font-kpi-number text-kpi-number text-stone-900">
    5
   </div>
   <div class="font-metadata text-metadata text-stone-600 mt-auto">
    3 AMC, 2 Cleaning
   </div>
  </div>
  <!-- Card 2 (Lavender) -->
  <div class="bg-metric-lavender rounded-[16px] p-card-padding flex flex-col gap-2 hover:-translate-y-[2px] transition-transform cursor-pointer border border-[#EEEBE4]/50 shadow-[0_4px_4px_-2px_rgba(0,0,0,0.02)]">
   <div class="flex justify-between items-start">
    <span class="font-h3 text-h3 text-stone-800">
     Active AMC Contracts
    </span>
    <span class="material-symbols-outlined text-stone-700">
     contract
    </span>
   </div>
   <div class="font-kpi-number text-kpi-number text-stone-900">
    28
   </div>
   <div class="font-metadata text-metadata text-stone-600 mt-auto">
    Total value ₹12,50,000
   </div>
  </div>
  <!-- Card 3 (Soft Pink) -->
  <div class="bg-metric-pink rounded-[16px] p-card-padding flex flex-col gap-2 hover:-translate-y-[2px] transition-transform cursor-pointer border border-[#EEEBE4]/50 shadow-[0_4px_4px_-2px_rgba(0,0,0,0.02)]">
   <div class="flex justify-between items-start">
    <span class="font-h3 text-h3 text-stone-800">
     Expiring Contracts
    </span>
    <span class="material-symbols-outlined text-stone-700">
     warning
    </span>
   </div>
   <div class="font-kpi-number text-kpi-number text-stone-900">
    3
   </div>
   <div class="font-metadata text-metadata text-stone-600 mt-auto">
    Within 30 days
   </div>
  </div>
  <!-- Card 4 (Sage) -->
  <div class="bg-metric-sage rounded-[16px] p-card-padding flex flex-col gap-2 hover:-translate-y-[2px] transition-transform cursor-pointer border border-[#EEEBE4]/50 shadow-[0_4px_4px_-2px_rgba(0,0,0,0.02)]">
   <div class="flex justify-between items-start">
    <span class="font-h3 text-h3 text-stone-800">
     Service Compliance
    </span>
    <span class="material-symbols-outlined text-stone-700">
     verified
    </span>
   </div>
   <div class="font-kpi-number text-kpi-number text-stone-900">
    98%
   </div>
   <div class="font-metadata text-metadata text-stone-600 mt-auto">
    YTD on-time completion
   </div>
  </div>
 </div>
 <!-- Tables Section -->
 <div class="grid grid-cols-1 lg:grid-cols-12 gap-section-gap">
  <!-- Upcoming Maintenance Table (Spans 7 cols) -->
  <div class="lg:col-span-7 bg-surface rounded-[16px] border border-border-default shadow-[0_4px_4px_-2px_rgba(0,0,0,0.02)] overflow-hidden flex flex-col">
   <div class="p-card-padding border-b border-border-default bg-surface flex justify-between items-center">
    <h2 class="font-h2 text-h2 text-text-primary">
     Upcoming Maintenance
    </h2>
    <button class="text-primary-container font-h3 text-h3 hover:underline">
     View All
    </button>
   </div>
   <div class="overflow-x-auto">
    <table class="w-full text-left border-collapse">
     <thead>
      <tr class="bg-surface-subtle border-b border-border-default">
       <th class="font-table-header text-table-header text-text-secondary py-3 px-4">
        Asset
       </th>
       <th class="font-table-header text-table-header text-text-secondary py-3 px-4">
        Type
       </th>
       <th class="font-table-header text-table-header text-text-secondary py-3 px-4">
        Next Due
       </th>
       <th class="font-table-header text-table-header text-text-secondary py-3 px-4">
        Assigned To
       </th>
       <th class="font-table-header text-table-header text-text-secondary py-3 px-4 text-right">
        Actions
       </th>
      </tr>
     </thead>
     <tbody class="font-body text-body">
      <tr class="border-b border-border-default hover:bg-accent-subtle transition-colors group cursor-pointer border-l-[3px] border-l-primary-container bg-surface">
       <td class="py-3 px-4 font-medium text-text-primary">
        HVAC Unit B-04
       </td>
       <td class="py-3 px-4 text-text-secondary">
        Service
       </td>
       <td class="py-3 px-4">
        <div class="flex flex-col">
         <span class="text-text-primary">
          Oct 24, 2023
         </span>
         <span class="font-metadata text-metadata text-[#991B1B]">
          In 3 days
         </span>
        </div>
       </td>
       <td class="py-3 px-4 text-text-secondary">
        CoolingTech Inc.
       </td>
       <td class="py-3 px-4 text-right">
        <button class="text-stone-400 hover:text-primary-container">
         <span class="material-symbols-outlined text-[20px]">
          more_vert
         </span>
        </button>
       </td>
      </tr>
      <tr class="border-b border-border-default hover:bg-accent-subtle transition-colors group cursor-pointer border-l-[3px] border-l-transparent bg-surface">
       <td class="py-3 px-4 font-medium text-text-primary">
        Server Rack 12
       </td>
       <td class="py-3 px-4 text-text-secondary">
        Inspection
       </td>
       <td class="py-3 px-4">
        <div class="flex flex-col">
         <span class="text-text-primary">
          Oct 26, 2023
         </span>
         <span class="font-metadata text-metadata text-stone-500">
          In 5 days
         </span>
        </div>
       </td>
       <td class="py-3 px-4 text-text-secondary">
        In-house IT
       </td>
       <td class="py-3 px-4 text-right">
        <button class="text-stone-400 hover:text-primary-container">
         <span class="material-symbols-outlined text-[20px]">
          more_vert
         </span>
        </button>
       </td>
      </tr>
      <tr class="border-b border-border-default hover:bg-accent-subtle transition-colors group cursor-pointer border-l-[3px] border-l-transparent bg-surface-subtle">
       <td class="py-3 px-4 font-medium text-text-primary">
        Main Elevators
       </td>
       <td class="py-3 px-4 text-text-secondary">
        Service
       </td>
       <td class="py-3 px-4">
        <div class="flex flex-col">
         <span class="text-text-primary">
          Nov 02, 2023
         </span>
         <span class="font-metadata text-metadata text-stone-500">
          In 12 days
         </span>
        </div>
       </td>
       <td class="py-3 px-4 text-text-secondary">
        Otis Elevator Co.
       </td>
       <td class="py-3 px-4 text-right">
        <button class="text-stone-400 hover:text-primary-container">
         <span class="material-symbols-outlined text-[20px]">
          more_vert
         </span>
        </button>
       </td>
      </tr>
     </tbody>
    </table>
   </div>
  </div>
  <!-- AMC Contracts Table (Spans 5 cols) -->
  <div class="lg:col-span-5 bg-surface rounded-[16px] border border-border-default shadow-[0_4px_4px_-2px_rgba(0,0,0,0.02)] overflow-hidden flex flex-col">
   <div class="p-card-padding border-b border-border-default bg-surface flex justify-between items-center">
    <h2 class="font-h2 text-h2 text-text-primary">
     AMC Contracts
    </h2>
    <button class="p-1 text-stone-400 hover:text-primary-container rounded">
     <span class="material-symbols-outlined">
      filter_list
     </span>
    </button>
   </div>
   <div class="overflow-x-auto">
    <table class="w-full text-left border-collapse">
     <thead>
      <tr class="bg-surface-subtle border-b border-border-default">
       <th class="font-table-header text-table-header text-text-secondary py-3 px-4">
        Vendor &amp; Scope
       </th>
       <th class="font-table-header text-table-header text-text-secondary py-3 px-4 text-right">
        Value (₹)
       </th>
       <th class="font-table-header text-table-header text-text-secondary py-3 px-4 text-right">
        Status
       </th>
      </tr>
     </thead>
     <tbody class="font-body text-body">
      <tr class="border-b border-border-default hover:bg-accent-subtle transition-colors group cursor-pointer border-l-[3px] border-l-transparent bg-surface">
       <td class="py-3 px-4">
        <div class="font-medium text-text-primary">
         CoolingTech Inc.
        </div>
        <div class="font-metadata text-metadata text-text-secondary">
         HVAC Systems
        </div>
        <div class="font-metadata text-metadata text-stone-500 mt-1">
         Exp: Dec 31, 2023
        </div>
       </td>
       <td class="py-3 px-4 font-mono-data text-mono-data text-text-primary text-right align-top">
        2,50,000
       </td>
       <td class="py-3 px-4 text-right align-top">
        <span class="inline-flex items-center px-2 py-1 rounded-full text-[11px] font-semibold bg-metric-sage text-status-in-stock border border-status-in-stock/20">
         <span class="sr-only">
          Status:
         </span>
         Active
        </span>
       </td>
      </tr>
      <tr class="border-b border-border-default hover:bg-accent-subtle transition-colors group cursor-pointer border-l-[3px] border-l-transparent bg-surface-subtle">
       <td class="py-3 px-4">
        <div class="font-medium text-text-primary">
         SecureIT Solutions
        </div>
        <div class="font-metadata text-metadata text-text-secondary">
         CCTV Network
        </div>
        <div class="font-metadata text-metadata text-[#991B1B] mt-1">
         Exp: Nov 15, 2023
        </div>
       </td>
       <td class="py-3 px-4 font-mono-data text-mono-data text-text-primary text-right align-top">
        1,80,000
       </td>
       <td class="py-3 px-4 text-right align-top">
        <span class="inline-flex items-center px-2 py-1 rounded-full text-[11px] font-semibold bg-[#FEF2F2] text-[#991B1B] border border-[#991B1B]/20">
         <span class="sr-only">
          Status:
         </span>
         Expiring
        </span>
       </td>
      </tr>
      <tr class="border-b border-border-default hover:bg-accent-subtle transition-colors group cursor-pointer border-l-[3px] border-l-transparent bg-surface">
       <td class="py-3 px-4">
        <div class="font-medium text-text-primary">
         CleanSweep Services
        </div>
        <div class="font-metadata text-metadata text-text-secondary">
         Janitorial
        </div>
        <div class="font-metadata text-metadata text-stone-500 mt-1">
         Exp: Mar 31, 2024
        </div>
       </td>
       <td class="py-3 px-4 font-mono-data text-mono-data text-text-primary text-right align-top">
        8,20,000
       </td>
       <td class="py-3 px-4 text-right align-top">
        <span class="inline-flex items-center px-2 py-1 rounded-full text-[11px] font-semibold bg-metric-sage text-status-in-stock border border-status-in-stock/20">
         <span class="sr-only">
          Status:
         </span>
         Active
        </span>
       </td>
      </tr>
     </tbody>
    </table>
   </div>
  </div>
 </div>
</main>

</template>

<script setup lang="ts">
// Autogenerated from maintenance_schedule_contracts
</script>
