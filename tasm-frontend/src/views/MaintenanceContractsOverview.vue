<template>
<main class="flex-1 p-page-margin w-full max-w-[1400px] mx-auto flex flex-col gap-section-gap">
 <!-- Page Title -->
 <div class="flex flex-col gap-1">
  <h1 class="font-h1 text-h1 text-text-primary">
   Operations Overview
  </h1>
  <p class="font-body text-body text-text-secondary">
   Manage maintenance queues and vendor contracts.
  </p>
 </div>
 <!-- Header KPIs (4 Cards) -->
 <section class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-inline">
  <!-- KPI 1: Amber -->
  <div class="bg-metric-amber rounded-xl p-card-padding flex flex-col gap-base border border-border-default shadow-sm hover:-translate-y-0.5 transition-transform duration-200 cursor-pointer">
   <div class="flex justify-between items-start mb-2">
    <span class="font-metadata text-metadata text-text-secondary uppercase tracking-wider">
     Active Work Orders
    </span>
    <span class="material-symbols-outlined text-outline">
     build
    </span>
   </div>
   <div class="font-kpi-number text-kpi-number text-text-primary">
    42
   </div>
   <div class="font-metadata text-metadata text-text-secondary mt-1 flex items-center gap-1">
    <span class="material-symbols-outlined text-[14px] text-status-critical">
     arrow_upward
    </span>
    <span>
     12 critical
    </span>
   </div>
  </div>
  <!-- KPI 2: Peach (Using error-container for semantic overdue feel) -->
  <div class="bg-error-container rounded-xl p-card-padding flex flex-col gap-base border border-border-default shadow-sm hover:-translate-y-0.5 transition-transform duration-200 cursor-pointer">
   <div class="flex justify-between items-start mb-2">
    <span class="font-metadata text-metadata text-text-secondary uppercase tracking-wider">
     Overdue Maintenance
    </span>
    <span class="material-symbols-outlined text-outline">
     warning
    </span>
   </div>
   <div class="font-kpi-number text-kpi-number text-text-primary">
    18
   </div>
   <div class="font-metadata text-metadata text-text-secondary mt-1 flex items-center gap-1">
    <span class="text-on-error-container font-medium">
     +3 since yesterday
    </span>
   </div>
  </div>
  <!-- KPI 3: Lavender -->
  <div class="bg-metric-lavender rounded-xl p-card-padding flex flex-col gap-base border border-border-default shadow-sm hover:-translate-y-0.5 transition-transform duration-200 cursor-pointer">
   <div class="flex justify-between items-start mb-2">
    <span class="font-metadata text-metadata text-text-secondary uppercase tracking-wider">
     Expiring Contracts
    </span>
    <span class="material-symbols-outlined text-outline">
     contract
    </span>
   </div>
   <div class="font-kpi-number text-kpi-number text-text-primary">
    7
   </div>
   <div class="font-metadata text-metadata text-text-secondary mt-1 flex items-center gap-1">
    <span>
     Next 30 days
    </span>
   </div>
  </div>
  <!-- KPI 4: Sage -->
  <div class="bg-metric-sage rounded-xl p-card-padding flex flex-col gap-base border border-border-default shadow-sm hover:-translate-y-0.5 transition-transform duration-200 cursor-pointer">
   <div class="flex justify-between items-start mb-2">
    <span class="font-metadata text-metadata text-text-secondary uppercase tracking-wider">
     Total Vendor Spend
    </span>
    <span class="material-symbols-outlined text-outline">
     payments
    </span>
   </div>
   <div class="font-kpi-number text-kpi-number text-text-primary">
    $1.2M
   </div>
   <div class="font-metadata text-metadata text-text-secondary mt-1 flex items-center gap-1">
    <span class="material-symbols-outlined text-[14px] text-status-in-stock">
     arrow_downward
    </span>
    <span>
     4% vs last quarter
    </span>
   </div>
  </div>
 </section>
 <!-- 60/40 Split Layout -->
 <section class="grid grid-cols-1 lg:grid-cols-12 gap-inline items-start">
  <!-- Left Column (60%) - Work Order Queue -->
  <div class="lg:col-span-8 bg-surface rounded-xl border border-border-default shadow-sm overflow-hidden flex flex-col">
   <div class="p-card-padding border-b border-border-default flex items-center justify-between bg-surface-subtle">
    <h2 class="font-h2 text-h2 text-text-primary">
     Work Order Queue
    </h2>
    <button class="bg-text-primary text-on-primary font-body text-body px-4 py-2 rounded-lg hover:bg-secondary transition-colors text-sm">
     New Order
    </button>
   </div>
   <div class="overflow-x-auto">
    <table class="w-full text-left border-collapse">
     <thead>
      <tr>
       <th class="font-table-header text-table-header text-text-secondary py-3 px-4 bg-surface-subtle border-b border-border-default uppercase">
        ID
       </th>
       <th class="font-table-header text-table-header text-text-secondary py-3 px-4 bg-surface-subtle border-b border-border-default uppercase">
        Asset / Location
       </th>
       <th class="font-table-header text-table-header text-text-secondary py-3 px-4 bg-surface-subtle border-b border-border-default uppercase">
        Severity
       </th>
       <th class="font-table-header text-table-header text-text-secondary py-3 px-4 bg-surface-subtle border-b border-border-default uppercase">
        Reported
       </th>
       <th class="font-table-header text-table-header text-text-secondary py-3 px-4 bg-surface-subtle border-b border-border-default uppercase text-right">
        Actions
       </th>
      </tr>
     </thead>
     <tbody class="font-body text-body text-text-primary">
      <!-- Row 1 (Critical) -->
      <tr class="border-b border-border-default hover:bg-metric-amber transition-colors bg-surface group cursor-pointer border-l-[3px] border-l-status-critical">
       <td class="py-4 px-4 font-mono-data text-mono-data text-text-secondary">
        WO-8922
       </td>
       <td class="py-4 px-4">
        <div class="font-medium text-text-primary">
         HVAC Unit B-East
        </div>
        <div class="font-metadata text-metadata text-text-secondary mt-0.5">
         Bldg 4, Roof
        </div>
       </td>
       <td class="py-4 px-4">
        <span class="inline-flex items-center px-2 py-0.5 rounded-full font-metadata text-metadata bg-status-critical text-on-error">
         <span class="sr-only">
          Severity:
         </span>
         Critical
        </span>
       </td>
       <td class="py-4 px-4 text-text-secondary">
        2 hrs ago
       </td>
       <td class="py-4 px-4 text-right">
        <button class="text-text-secondary hover:text-text-primary">
         <span class="material-symbols-outlined text-[20px]">
          more_vert
         </span>
        </button>
       </td>
      </tr>
      <!-- Row 2 (High) -->
      <tr class="border-b border-border-default hover:bg-metric-amber transition-colors bg-surface-subtle group cursor-pointer border-l-[3px] border-l-transparent hover:border-l-primary-container">
       <td class="py-4 px-4 font-mono-data text-mono-data text-text-secondary">
        WO-8921
       </td>
       <td class="py-4 px-4">
        <div class="font-medium text-text-primary">
         Elevator Bank 2
        </div>
        <div class="font-metadata text-metadata text-text-secondary mt-0.5">
         Bldg 1, Lobby
        </div>
       </td>
       <td class="py-4 px-4">
        <span class="inline-flex items-center px-2 py-0.5 rounded-full font-metadata text-metadata bg-primary-container text-on-primary-container">
         <span class="sr-only">
          Severity:
         </span>
         High
        </span>
       </td>
       <td class="py-4 px-4 text-text-secondary">
        5 hrs ago
       </td>
       <td class="py-4 px-4 text-right">
        <button class="text-text-secondary hover:text-text-primary">
         <span class="material-symbols-outlined text-[20px]">
          more_vert
         </span>
        </button>
       </td>
      </tr>
      <!-- Row 3 (Medium) -->
      <tr class="border-b border-border-default hover:bg-metric-amber transition-colors bg-surface group cursor-pointer border-l-[3px] border-l-transparent hover:border-l-primary-container">
       <td class="py-4 px-4 font-mono-data text-mono-data text-text-secondary">
        WO-8918
       </td>
       <td class="py-4 px-4">
        <div class="font-medium text-text-primary">
         Water Leak
        </div>
        <div class="font-metadata text-metadata text-text-secondary mt-0.5">
         Cafeteria Restroom
        </div>
       </td>
       <td class="py-4 px-4">
        <span class="inline-flex items-center px-2 py-0.5 rounded-full font-metadata text-metadata bg-surface-variant text-on-surface-variant">
         <span class="sr-only">
          Severity:
         </span>
         Medium
        </span>
       </td>
       <td class="py-4 px-4 text-text-secondary">
        1 day ago
       </td>
       <td class="py-4 px-4 text-right">
        <button class="text-text-secondary hover:text-text-primary">
         <span class="material-symbols-outlined text-[20px]">
          more_vert
         </span>
        </button>
       </td>
      </tr>
      <!-- Row 4 (Low) -->
      <tr class="hover:bg-metric-amber transition-colors bg-surface-subtle group cursor-pointer border-l-[3px] border-l-transparent hover:border-l-primary-container">
       <td class="py-4 px-4 font-mono-data text-mono-data text-text-secondary">
        WO-8915
       </td>
       <td class="py-4 px-4">
        <div class="font-medium text-text-primary">
         Flickering Lights
        </div>
        <div class="font-metadata text-metadata text-text-secondary mt-0.5">
         Office Suite 300
        </div>
       </td>
       <td class="py-4 px-4">
        <span class="inline-flex items-center px-2 py-0.5 rounded-full font-metadata text-metadata bg-secondary-container text-on-secondary-container border border-border-default">
         <span class="sr-only">
          Severity:
         </span>
         Low
        </span>
       </td>
       <td class="py-4 px-4 text-text-secondary">
        2 days ago
       </td>
       <td class="py-4 px-4 text-right">
        <button class="text-text-secondary hover:text-text-primary">
         <span class="material-symbols-outlined text-[20px]">
          more_vert
         </span>
        </button>
       </td>
      </tr>
     </tbody>
    </table>
   </div>
   <div class="p-4 border-t border-border-default bg-surface text-center">
    <a class="font-metadata text-metadata text-tertiary hover:underline" href="#">
     View All Orders
    </a>
   </div>
  </div>
  <!-- Right Column (40%) - Contract Expiry List -->
  <div class="lg:col-span-4 bg-surface rounded-xl border border-border-default shadow-sm flex flex-col">
   <div class="p-card-padding border-b border-border-default flex items-center justify-between">
    <h2 class="font-h2 text-h2 text-text-primary">
     Upcoming Renewals
    </h2>
    <button class="text-text-secondary hover:text-text-primary p-1 rounded-md hover:bg-surface-subtle transition-colors">
     <span class="material-symbols-outlined text-[20px]">
      filter_list
     </span>
    </button>
   </div>
   <div class="flex flex-col p-4 gap-2">
    <!-- List Item 1 -->
    <div class="flex items-start gap-4 p-3 rounded-lg hover:bg-surface-subtle transition-colors cursor-pointer border border-transparent hover:border-border-default">
     <div class="w-10 h-10 rounded-full bg-error-container flex items-center justify-center shrink-0">
      <span class="material-symbols-outlined text-on-error-container text-[20px]">
       corporate_fare
      </span>
     </div>
     <div class="flex-1 min-w-0">
      <h3 class="font-h3 text-h3 text-text-primary truncate">
       Apex Security Services
      </h3>
      <p class="font-metadata text-metadata text-text-secondary truncate mt-0.5">
       Facility guarding contract
      </p>
     </div>
     <div class="text-right shrink-0">
      <div class="font-h3 text-h3 text-status-critical">
       5 days
      </div>
      <div class="font-metadata text-metadata text-text-secondary">
       Oct 12
      </div>
     </div>
    </div>
    <!-- List Item 2 -->
    <div class="flex items-start gap-4 p-3 rounded-lg hover:bg-surface-subtle transition-colors cursor-pointer border border-transparent hover:border-border-default">
     <div class="w-10 h-10 rounded-full bg-metric-lavender flex items-center justify-center shrink-0">
      <span class="material-symbols-outlined text-tertiary text-[20px]">
       cleaning_services
      </span>
     </div>
     <div class="flex-1 min-w-0">
      <h3 class="font-h3 text-h3 text-text-primary truncate">
       ProClean Janitorial
      </h3>
      <p class="font-metadata text-metadata text-text-secondary truncate mt-0.5">
       Nightly deep clean SLA
      </p>
     </div>
     <div class="text-right shrink-0">
      <div class="font-h3 text-h3 text-text-primary">
       14 days
      </div>
      <div class="font-metadata text-metadata text-text-secondary">
       Oct 21
      </div>
     </div>
    </div>
    <!-- List Item 3 -->
    <div class="flex items-start gap-4 p-3 rounded-lg hover:bg-surface-subtle transition-colors cursor-pointer border border-transparent hover:border-border-default">
     <div class="w-10 h-10 rounded-full bg-surface-variant flex items-center justify-center shrink-0">
      <span class="material-symbols-outlined text-on-surface-variant text-[20px]">
       hvac
      </span>
     </div>
     <div class="flex-1 min-w-0">
      <h3 class="font-h3 text-h3 text-text-primary truncate">
       ClimateTech Solutions
      </h3>
      <p class="font-metadata text-metadata text-text-secondary truncate mt-0.5">
       Quarterly preventive maint.
      </p>
     </div>
     <div class="text-right shrink-0">
      <div class="font-h3 text-h3 text-text-primary">
       28 days
      </div>
      <div class="font-metadata text-metadata text-text-secondary">
       Nov 04
      </div>
     </div>
    </div>
    <!-- List Item 4 -->
    <div class="flex items-start gap-4 p-3 rounded-lg hover:bg-surface-subtle transition-colors cursor-pointer border border-transparent hover:border-border-default">
     <div class="w-10 h-10 rounded-full bg-surface-variant flex items-center justify-center shrink-0">
      <span class="material-symbols-outlined text-on-surface-variant text-[20px]">
       router
      </span>
     </div>
     <div class="flex-1 min-w-0">
      <h3 class="font-h3 text-h3 text-text-primary truncate">
       DataLink ISP
      </h3>
      <p class="font-metadata text-metadata text-text-secondary truncate mt-0.5">
       Campus broadband
      </p>
     </div>
     <div class="text-right shrink-0">
      <div class="font-h3 text-h3 text-text-primary">
       45 days
      </div>
      <div class="font-metadata text-metadata text-text-secondary">
       Nov 21
      </div>
     </div>
    </div>
   </div>
   <div class="p-4 mt-auto border-t border-border-default text-center">
    <button class="font-metadata text-metadata text-tertiary hover:underline font-medium">
     Manage Contracts
    </button>
   </div>
  </div>
 </section>
</main>

</template>

<script setup lang="ts">
// Autogenerated from maintenance_contracts_overview
</script>
