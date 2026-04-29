<template>
<main class="pl-[248px] w-full min-h-screen">
 <div class="max-w-[1400px] mx-auto p-section-gap space-y-section-gap">
  <!-- Page Header & Tabs -->
  <div class="flex flex-col md:flex-row md:items-end justify-between gap-4 border-b border-border-default pb-4">
   <div>
    <h1 class="font-h1 text-h1 text-text-primary mb-2">
     Maintenance Operations
    </h1>
    <p class="font-body text-body text-text-secondary">
     Manage work orders, view preventive schedules, and oversee active service contracts.
    </p>
   </div>
   <div class="flex items-center p-1 bg-surface-subtle border border-border-default rounded-lg">
    <button class="px-4 py-1.5 rounded-md bg-surface border border-border-default shadow-sm font-h3 text-h3 text-text-primary">
     Work Orders
    </button>
    <button class="px-4 py-1.5 rounded-md font-body text-body text-text-secondary hover:text-text-primary transition-colors">
     Schedules
    </button>
    <button class="px-4 py-1.5 rounded-md font-body text-body text-text-secondary hover:text-text-primary transition-colors">
     Contracts
    </button>
   </div>
  </div>
  <!-- Bento Grid: KPIs -->
  <div class="grid grid-cols-1 md:grid-cols-3 gap-inline">
   <!-- KPI 1: Overdue (Peach/Error context) -->
   <div class="bg-error-container rounded-xl p-card-padding flex flex-col justify-between hover:-translate-y-[2px] transition-transform cursor-pointer shadow-sm border border-error-container">
    <div class="flex items-center justify-between mb-4">
     <span class="font-h2 text-h2 text-on-error-container">
      Overdue Services
     </span>
     <div class="w-10 h-10 rounded-full bg-white/40 flex items-center justify-center text-on-error-container">
      <span class="material-symbols-outlined">
       warning
      </span>
     </div>
    </div>
    <div>
     <div class="font-kpi-number text-kpi-number text-on-error-container mb-1">
      12
     </div>
     <p class="font-metadata text-metadata text-on-error-container/80 flex items-center gap-1">
      <span class="material-symbols-outlined text-[14px]">
       arrow_upward
      </span>
      +3 since yesterday
     </p>
    </div>
   </div>
   <!-- KPI 2: Upcoming (Amber) -->
   <div class="bg-metric-amber rounded-xl p-card-padding flex flex-col justify-between hover:-translate-y-[2px] transition-transform cursor-pointer shadow-sm border border-border-default">
    <div class="flex items-center justify-between mb-4">
     <span class="font-h2 text-h2 text-on-background">
      Upcoming Tasks
     </span>
     <div class="w-10 h-10 rounded-full bg-white/50 flex items-center justify-center text-on-background">
      <span class="material-symbols-outlined">
       event_upcoming
      </span>
     </div>
    </div>
    <div>
     <div class="font-kpi-number text-kpi-number text-on-background mb-1">
      45
     </div>
     <p class="font-metadata text-metadata text-text-secondary flex items-center gap-1">
      Scheduled next 7 days
     </p>
    </div>
   </div>
   <!-- KPI 3: Active Contracts (Lavender) -->
   <div class="bg-metric-lavender rounded-xl p-card-padding flex flex-col justify-between hover:-translate-y-[2px] transition-transform cursor-pointer shadow-sm border border-border-default">
    <div class="flex items-center justify-between mb-4">
     <span class="font-h2 text-h2 text-on-background">
      Active Contracts
     </span>
     <div class="w-10 h-10 rounded-full bg-white/50 flex items-center justify-center text-on-background">
      <span class="material-symbols-outlined">
       assignment
      </span>
     </div>
    </div>
    <div>
     <div class="font-kpi-number text-kpi-number text-on-background mb-1">
      18
     </div>
     <p class="font-metadata text-metadata text-text-secondary flex items-center gap-1">
      2 expiring this month
     </p>
    </div>
   </div>
  </div>
  <!-- Main Layout Grid (2/3 Table, 1/3 Schedule) -->
  <div class="grid grid-cols-1 lg:grid-cols-3 gap-inline">
   <!-- Left: High Density Work Orders Table -->
   <div class="lg:col-span-2 bg-surface rounded-xl border border-border-default shadow-sm overflow-hidden flex flex-col">
    <div class="p-5 border-b border-border-default flex items-center justify-between bg-white">
     <h2 class="font-h2 text-h2 text-text-primary">
      Active Work Orders
     </h2>
     <button class="text-primary-container font-h3 text-h3 flex items-center gap-1 hover:underline">
      View All
      <span class="material-symbols-outlined text-[16px]">
       chevron_right
      </span>
     </button>
    </div>
    <div class="overflow-x-auto">
     <table class="w-full text-left border-collapse">
      <thead>
       <tr class="bg-surface-subtle border-b border-border-default font-table-header text-table-header text-text-secondary uppercase">
        <th class="p-3 pl-5 font-medium">
         Order ID
        </th>
        <th class="p-3 font-medium">
         Asset &amp; Location
        </th>
        <th class="p-3 font-medium">
         Issue
        </th>
        <th class="p-3 font-medium">
         Severity
        </th>
        <th class="p-3 pr-5 font-medium text-right">
         Target Date
        </th>
       </tr>
      </thead>
      <tbody class="font-body text-body text-text-primary divide-y divide-border-default">
       <!-- Row 1 -->
       <tr class="hover:bg-metric-amber/20 transition-colors group cursor-pointer border-l-4 border-l-transparent hover:border-l-primary-container">
        <td class="p-3 pl-5 font-mono-data text-mono-data text-text-secondary">
         WO-4092
        </td>
        <td class="p-3">
         <div class="font-medium text-text-primary">
          HVAC Unit B-East
         </div>
         <div class="font-metadata text-metadata text-text-secondary">
          Bhavani Building, Roof
         </div>
        </td>
        <td class="p-3 text-text-secondary">
         Compressor failure, temp drop
        </td>
        <td class="p-3">
         <span class="inline-flex items-center px-2 py-0.5 rounded text-[11px] font-bold tracking-wide uppercase bg-status-critical text-white">
          Critical
         </span>
        </td>
        <td class="p-3 pr-5 text-right font-mono-data text-mono-data text-status-critical">
         Today
        </td>
       </tr>
       <!-- Row 2 -->
       <tr class="hover:bg-metric-amber/20 transition-colors group cursor-pointer border-l-4 border-l-transparent hover:border-l-primary-container">
        <td class="p-3 pl-5 font-mono-data text-mono-data text-text-secondary">
         WO-4088
        </td>
        <td class="p-3">
         <div class="font-medium text-text-primary">
          Main Server Rack 3
         </div>
         <div class="font-metadata text-metadata text-text-secondary">
          Tejaswini Data Center
         </div>
        </td>
        <td class="p-3 text-text-secondary">
         UPS Battery replacement due
        </td>
        <td class="p-3">
         <span class="inline-flex items-center px-2 py-0.5 rounded text-[11px] font-bold tracking-wide uppercase bg-surface-tint text-white">
          High
         </span>
        </td>
        <td class="p-3 pr-5 text-right font-mono-data text-mono-data text-text-primary">
         Oct 24, 2023
        </td>
       </tr>
       <!-- Row 3 -->
       <tr class="bg-surface-subtle hover:bg-metric-amber/20 transition-colors group cursor-pointer border-l-4 border-l-transparent hover:border-l-primary-container">
        <td class="p-3 pl-5 font-mono-data text-mono-data text-text-secondary">
         WO-4085
        </td>
        <td class="p-3">
         <div class="font-medium text-text-primary">
          Elevator Bank A
         </div>
         <div class="font-metadata text-metadata text-text-secondary">
          Nila Building, Lobby
         </div>
        </td>
        <td class="p-3 text-text-secondary">
         Routine quarterly inspection
        </td>
        <td class="p-3">
         <span class="inline-flex items-center px-2 py-0.5 rounded text-[11px] font-bold tracking-wide uppercase bg-metric-amber text-on-background">
          Medium
         </span>
        </td>
        <td class="p-3 pr-5 text-right font-mono-data text-mono-data text-text-primary">
         Oct 26, 2023
        </td>
       </tr>
       <!-- Row 4 -->
       <tr class="hover:bg-metric-amber/20 transition-colors group cursor-pointer border-l-4 border-l-transparent hover:border-l-primary-container">
        <td class="p-3 pl-5 font-mono-data text-mono-data text-text-secondary">
         WO-4080
        </td>
        <td class="p-3">
         <div class="font-medium text-text-primary">
          Landscape Sprinklers
         </div>
         <div class="font-metadata text-metadata text-text-secondary">
          Campus North Lawn
         </div>
        </td>
        <td class="p-3 text-text-secondary">
         Fix broken head Zone 4
        </td>
        <td class="p-3">
         <span class="inline-flex items-center px-2 py-0.5 rounded text-[11px] font-bold tracking-wide uppercase bg-surface-variant text-on-surface-variant">
          Low
         </span>
        </td>
        <td class="p-3 pr-5 text-right font-mono-data text-mono-data text-text-secondary">
         Oct 28, 2023
        </td>
       </tr>
       <!-- Row 5 -->
       <tr class="bg-surface-subtle hover:bg-metric-amber/20 transition-colors group cursor-pointer border-l-4 border-l-transparent hover:border-l-primary-container">
        <td class="p-3 pl-5 font-mono-data text-mono-data text-text-secondary">
         WO-4075
        </td>
        <td class="p-3">
         <div class="font-medium text-text-primary">
          Parking Barrier Gate
         </div>
         <div class="font-metadata text-metadata text-text-secondary">
          South Entrance
         </div>
        </td>
        <td class="p-3 text-text-secondary">
         Sensor alignment calibration
        </td>
        <td class="p-3">
         <span class="inline-flex items-center px-2 py-0.5 rounded text-[11px] font-bold tracking-wide uppercase bg-metric-amber text-on-background">
          Medium
         </span>
        </td>
        <td class="p-3 pr-5 text-right font-mono-data text-mono-data text-text-primary">
         Oct 29, 2023
        </td>
       </tr>
      </tbody>
     </table>
    </div>
   </div>
   <!-- Right: Calendar Schedule View -->
   <div class="bg-surface rounded-xl border border-border-default shadow-sm p-5 flex flex-col h-full">
    <div class="flex items-center justify-between mb-4">
     <h2 class="font-h2 text-h2 text-text-primary">
      Schedule
     </h2>
     <div class="flex gap-1">
      <button class="p-1 hover:bg-surface-subtle rounded text-text-secondary">
       <span class="material-symbols-outlined text-[18px]">
        chevron_left
       </span>
      </button>
      <button class="p-1 hover:bg-surface-subtle rounded text-text-secondary">
       <span class="material-symbols-outlined text-[18px]">
        chevron_right
       </span>
      </button>
     </div>
    </div>
    <div class="text-center font-h3 text-h3 mb-4 text-text-primary">
     October 2023
    </div>
    <!-- Mini Calendar Grid -->
    <div class="grid grid-cols-7 gap-1 text-center font-metadata text-metadata mb-4">
     <div class="text-text-secondary pb-1">
      S
     </div>
     <div class="text-text-secondary pb-1">
      M
     </div>
     <div class="text-text-secondary pb-1">
      T
     </div>
     <div class="text-text-secondary pb-1">
      W
     </div>
     <div class="text-text-secondary pb-1">
      T
     </div>
     <div class="text-text-secondary pb-1">
      F
     </div>
     <div class="text-text-secondary pb-1">
      S
     </div>
     <!-- Dates (Mocked) -->
     <div class="p-1 text-text-secondary opacity-50">
      29
     </div>
     <div class="p-1 text-text-secondary opacity-50">
      30
     </div>
     <div class="p-1 text-text-secondary opacity-50">
      1
     </div>
     <div class="p-1">
      2
     </div>
     <div class="p-1">
      3
     </div>
     <div class="p-1">
      4
     </div>
     <div class="p-1">
      5
     </div>
     <div class="p-1">
      6
     </div>
     <div class="p-1">
      7
     </div>
     <div class="p-1">
      8
     </div>
     <div class="p-1">
      9
     </div>
     <div class="p-1 relative">
      10
      <span class="absolute bottom-0 left-1/2 -translate-x-1/2 w-1 h-1 bg-surface-tint rounded-full">
      </span>
     </div>
     <div class="p-1">
      11
     </div>
     <div class="p-1">
      12
     </div>
     <div class="p-1">
      13
     </div>
     <div class="p-1">
      14
     </div>
     <div class="p-1">
      15
     </div>
     <div class="p-1">
      16
     </div>
     <div class="p-1">
      17
     </div>
     <div class="p-1 relative">
      18
      <span class="absolute bottom-0 left-1/2 -translate-x-1/2 w-1 h-1 bg-surface-tint rounded-full">
      </span>
     </div>
     <div class="p-1">
      19
     </div>
     <div class="p-1">
      20
     </div>
     <div class="p-1">
      21
     </div>
     <div class="p-1">
      22
     </div>
     <div class="p-1 bg-primary-container text-white rounded font-bold shadow-sm">
      23
     </div>
     <div class="p-1 relative font-bold text-primary-container">
      24
      <span class="absolute bottom-0 left-1/2 -translate-x-1/2 w-1 h-1 bg-status-critical rounded-full">
      </span>
     </div>
     <div class="p-1">
      25
     </div>
     <div class="p-1">
      26
     </div>
    </div>
    <!-- Upcoming List -->
    <div class="mt-auto border-t border-border-default pt-4 space-y-3">
     <div class="text-metadata font-table-header uppercase tracking-wider text-text-secondary">
      Today's Focus
     </div>
     <div class="flex gap-3 items-start p-2 rounded-lg hover:bg-surface-subtle transition-colors cursor-pointer border border-transparent hover:border-border-default">
      <div class="w-2 h-2 rounded-full bg-status-critical mt-1.5 flex-shrink-0">
      </div>
      <div>
       <div class="font-h3 text-h3 text-text-primary">
        HVAC Unit B-East Check
       </div>
       <div class="font-metadata text-metadata text-text-secondary">
        10:00 AM • Maintenance Team
       </div>
      </div>
     </div>
     <div class="flex gap-3 items-start p-2 rounded-lg hover:bg-surface-subtle transition-colors cursor-pointer border border-transparent hover:border-border-default">
      <div class="w-2 h-2 rounded-full bg-surface-tint mt-1.5 flex-shrink-0">
      </div>
      <div>
       <div class="font-h3 text-h3 text-text-primary">
        Vendor Walkthrough
       </div>
       <div class="font-metadata text-metadata text-text-secondary">
        2:00 PM • Elevators Inc.
       </div>
      </div>
     </div>
    </div>
   </div>
  </div>
 </div>
</main>

</template>

<script setup lang="ts">
// Autogenerated from maintenance_contracts_management
</script>
