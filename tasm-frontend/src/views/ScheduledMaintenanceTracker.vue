<template>
<main class="md:ml-[248px] min-h-screen p-page-margin max-w-[1400px]">
 <!-- Page Header -->
 <div class="flex justify-between items-center mb-section-gap">
  <div>
   <h2 class="text-h1 font-h1 text-text-primary mb-1">
    Maintenance Schedules
   </h2>
   <p class="text-body font-body text-text-secondary">
    Track and manage upcoming asset maintenance across facilities.
   </p>
  </div>
  <button class="bg-text-primary text-surface px-5 py-2.5 rounded-lg text-body font-body font-medium hover:bg-opacity-90 transition-colors shadow-sm flex items-center gap-2">
   <span class="material-symbols-outlined text-[20px]">
    calendar_add_on
   </span>
   Add Schedule
  </button>
 </div>
 <!-- Filters & Controls Bar -->
 <div class="bg-surface p-4 rounded-xl border border-border-default shadow-[0_4px_4px_rgba(0,0,0,0.02)] mb-stack flex flex-wrap gap-4 items-center justify-between">
  <div class="flex gap-4">
   <div class="relative">
    <select class="appearance-none bg-surface-subtle border border-border-default text-body font-body text-text-primary rounded-md pl-4 pr-10 py-2 focus:outline-none focus:border-primary-container">
     <option>
      All Frequencies
     </option>
     <option>
      Monthly
     </option>
     <option>
      Quarterly
     </option>
     <option>
      Annual
     </option>
    </select>
    <span class="material-symbols-outlined absolute right-3 top-1/2 -translate-y-1/2 text-text-secondary pointer-events-none text-[20px]">
     expand_more
    </span>
   </div>
   <div class="relative">
    <select class="appearance-none bg-surface-subtle border border-border-default text-body font-body text-text-primary rounded-md pl-4 pr-10 py-2 focus:outline-none focus:border-primary-container">
     <option>
      All Statuses
     </option>
     <option>
      Scheduled
     </option>
     <option>
      In Progress
     </option>
     <option>
      Overdue
     </option>
    </select>
    <span class="material-symbols-outlined absolute right-3 top-1/2 -translate-y-1/2 text-text-secondary pointer-events-none text-[20px]">
     expand_more
    </span>
   </div>
  </div>
  <div class="flex gap-2">
   <button class="p-2 text-text-secondary hover:text-text-primary border border-transparent hover:border-border-default rounded-md transition-all">
    <span class="material-symbols-outlined">
     filter_list
    </span>
   </button>
   <button class="p-2 text-text-secondary hover:text-text-primary border border-transparent hover:border-border-default rounded-md transition-all">
    <span class="material-symbols-outlined">
     download
    </span>
   </button>
  </div>
 </div>
 <!-- High-Density Table Card -->
 <div class="bg-surface rounded-xl border border-border-default shadow-[0_4px_4px_rgba(0,0,0,0.02)] overflow-hidden">
  <div class="overflow-x-auto">
   <table class="w-full text-left border-collapse">
    <thead>
     <tr class="bg-surface-subtle border-b border-border-default">
      <th class="py-3 px-4 font-table-header text-table-header text-text-secondary uppercase">
       Asset Name
      </th>
      <th class="py-3 px-4 font-table-header text-table-header text-text-secondary uppercase">
       Tag ID
      </th>
      <th class="py-3 px-4 font-table-header text-table-header text-text-secondary uppercase">
       Maintenance Type
      </th>
      <th class="py-3 px-4 font-table-header text-table-header text-text-secondary uppercase">
       Frequency
      </th>
      <th class="py-3 px-4 font-table-header text-table-header text-text-secondary uppercase">
       Last Done
      </th>
      <th class="py-3 px-4 font-table-header text-table-header text-text-secondary uppercase">
       Next Due
      </th>
      <th class="py-3 px-4 font-table-header text-table-header text-text-secondary uppercase">
       Vendor
      </th>
      <th class="py-3 px-4 font-table-header text-table-header text-text-secondary uppercase text-right">
       Status
      </th>
      <th class="py-3 px-4">
      </th>
     </tr>
    </thead>
    <tbody class="divide-y divide-border-default text-body font-body text-text-primary">
     <!-- Row 1: Overdue (< 7d / Overdue) -->
     <tr class="hover:bg-accent-subtle transition-colors group">
      <td class="py-3 px-4">
       <div class="font-h3 text-h3 text-text-primary">
        HVAC Chiller Unit A
       </div>
       <div class="text-metadata font-metadata text-text-secondary">
        Bhavani Building, Roof
       </div>
      </td>
      <td class="py-3 px-4 font-mono-data text-mono-data text-text-secondary">
       AST-HV-001
      </td>
      <td class="py-3 px-4">
       AMC Service
      </td>
      <td class="py-3 px-4">
       Quarterly
      </td>
      <td class="py-3 px-4 text-text-secondary">
       12 Aug 2023
      </td>
      <td class="py-3 px-4">
       <span class="flex items-center gap-1 text-status-critical font-medium">
        <span class="material-symbols-outlined text-[16px]">
         warning
        </span>
        12 Nov 2023
       </span>
      </td>
      <td class="py-3 px-4 text-text-secondary">
       BlueStar Services
      </td>
      <td class="py-3 px-4 text-right">
       <span class="inline-flex items-center px-2 py-0.5 rounded text-metadata font-metadata bg-error-container text-on-error-container">
        Overdue
       </span>
      </td>
      <td class="py-3 px-4 text-right">
       <button class="text-text-secondary hover:text-primary-container opacity-0 group-hover:opacity-100 transition-opacity">
        <span class="material-symbols-outlined">
         more_vert
        </span>
       </button>
      </td>
     </tr>
     <!-- Row 2: Upcoming (7-30d) -->
     <tr class="bg-surface-subtle hover:bg-accent-subtle transition-colors group">
      <td class="py-3 px-4">
       <div class="font-h3 text-h3 text-text-primary">
        Elevator Passenger 3
       </div>
       <div class="text-metadata font-metadata text-text-secondary">
        Nila Building, Lobby
       </div>
      </td>
      <td class="py-3 px-4 font-mono-data text-mono-data text-text-secondary">
       AST-EL-042
      </td>
      <td class="py-3 px-4">
       Inspection
      </td>
      <td class="py-3 px-4">
       Monthly
      </td>
      <td class="py-3 px-4 text-text-secondary">
       28 Oct 2023
      </td>
      <td class="py-3 px-4">
       <span class="flex items-center gap-1 text-primary-container font-medium">
        <span class="material-symbols-outlined text-[16px]">
         schedule
        </span>
        28 Nov 2023
       </span>
      </td>
      <td class="py-3 px-4 text-text-secondary">
       Otis Elevators
      </td>
      <td class="py-3 px-4 text-right">
       <span class="inline-flex items-center px-2 py-0.5 rounded text-metadata font-metadata bg-surface-container-high text-on-surface-variant">
        Scheduled
       </span>
      </td>
      <td class="py-3 px-4 text-right">
       <button class="text-text-secondary hover:text-primary-container opacity-0 group-hover:opacity-100 transition-opacity">
        <span class="material-symbols-outlined">
         more_vert
        </span>
       </button>
      </td>
     </tr>
     <!-- Row 3: Healthy (> 30d) -->
     <tr class="hover:bg-accent-subtle transition-colors group">
      <td class="py-3 px-4">
       <div class="font-h3 text-h3 text-text-primary">
        Server Rack Cooling Sys
       </div>
       <div class="text-metadata font-metadata text-text-secondary">
        Thejaswini, Datacenter
       </div>
      </td>
      <td class="py-3 px-4 font-mono-data text-mono-data text-text-secondary">
       AST-IT-105
      </td>
      <td class="py-3 px-4">
       Calibration
      </td>
      <td class="py-3 px-4">
       Annual
      </td>
      <td class="py-3 px-4 text-text-secondary">
       15 Jan 2023
      </td>
      <td class="py-3 px-4">
       <span class="flex items-center gap-1 text-status-in-stock">
        <span class="material-symbols-outlined text-[16px]">
         check_circle
        </span>
        15 Jan 2024
       </span>
      </td>
      <td class="py-3 px-4 text-text-secondary">
       Internal IT
      </td>
      <td class="py-3 px-4 text-right">
       <span class="inline-flex items-center px-2 py-0.5 rounded text-metadata font-metadata bg-surface-container-high text-on-surface-variant">
        Scheduled
       </span>
      </td>
      <td class="py-3 px-4 text-right">
       <button class="text-text-secondary hover:text-primary-container opacity-0 group-hover:opacity-100 transition-opacity">
        <span class="material-symbols-outlined">
         more_vert
        </span>
       </button>
      </td>
     </tr>
     <!-- Row 4: In Progress -->
     <tr class="bg-surface-subtle border-l-[3px] border-l-primary-container hover:bg-accent-subtle transition-colors group">
      <td class="py-3 px-4">
       <div class="font-h3 text-h3 text-text-primary">
        Diesel Generator 500kVA
       </div>
       <div class="text-metadata font-metadata text-text-secondary">
        Gayathri Building, Exterior
       </div>
      </td>
      <td class="py-3 px-4 font-mono-data text-mono-data text-text-secondary">
       AST-DG-008
      </td>
      <td class="py-3 px-4">
       AMC Service
      </td>
      <td class="py-3 px-4">
       Quarterly
      </td>
      <td class="py-3 px-4 text-text-secondary">
       20 Aug 2023
      </td>
      <td class="py-3 px-4">
       <span class="flex items-center gap-1 text-status-checked-out font-medium">
        <span class="material-symbols-outlined text-[16px]">
         sync
        </span>
        20 Nov 2023
       </span>
      </td>
      <td class="py-3 px-4 text-text-secondary">
       Cummins Power
      </td>
      <td class="py-3 px-4 text-right">
       <span class="inline-flex items-center px-2 py-0.5 rounded text-metadata font-metadata bg-tertiary-container text-on-tertiary">
        In Progress
       </span>
      </td>
      <td class="py-3 px-4 text-right">
       <button class="text-text-secondary hover:text-primary-container opacity-0 group-hover:opacity-100 transition-opacity">
        <span class="material-symbols-outlined">
         more_vert
        </span>
       </button>
      </td>
     </tr>
    </tbody>
   </table>
  </div>
  <!-- Pagination -->
  <div class="border-t border-border-default px-4 py-3 flex items-center justify-between bg-surface text-metadata font-metadata text-text-secondary">
   <div>
    Showing 1 to 4 of 48 entries
   </div>
   <div class="flex gap-1">
    <button class="px-2 py-1 rounded border border-border-default hover:bg-surface-subtle disabled:opacity-50" disabled>
     Prev
    </button>
    <button class="px-3 py-1 rounded border border-primary-container bg-primary-container text-on-primary">
     1
    </button>
    <button class="px-3 py-1 rounded border border-border-default hover:bg-surface-subtle">
     2
    </button>
    <button class="px-3 py-1 rounded border border-border-default hover:bg-surface-subtle">
     3
    </button>
    <button class="px-2 py-1 rounded border border-border-default hover:bg-surface-subtle">
     Next
    </button>
   </div>
  </div>
 </div>
</main>

</template>

<script setup lang="ts">
// Autogenerated from scheduled_maintenance_tracker
</script>
