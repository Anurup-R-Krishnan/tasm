<template>
<main class="mt-[60px] p-page-margin max-w-[1400px] mx-auto w-full flex flex-col gap-section-gap">
 <!-- Page Header -->
 <div class="flex items-center justify-between">
  <div>
   <h1 class="font-h1 text-h1 text-text-primary mb-1">
    Self-Service Portal
   </h1>
   <p class="font-body text-body text-text-secondary">
    Manage your assigned equipment and pending requests.
   </p>
  </div>
  <button class="bg-primary text-surface font-h3 text-h3 px-4 py-2 rounded-DEFAULT hover:bg-primary-container transition-colors shadow-sm flex items-center gap-2">
   <span class="material-symbols-outlined text-[18px]">
    add_circle
   </span>
   Request New Asset
  </button>
 </div>
 <!-- KPI Cards Grid (Top) -->
 <div class="grid grid-cols-1 md:grid-cols-3 gap-card-padding">
  <!-- Card 1: My Assets -->
  <div class="bg-metric-amber rounded-xl p-card-padding shadow-card hover:-translate-y-0.5 transition-transform cursor-pointer border border-border-default/50">
   <div class="flex items-start justify-between mb-2">
    <h3 class="font-h3 text-h3 text-text-primary">
     My Assets
    </h3>
    <div class="w-8 h-8 rounded-full bg-surface/50 flex items-center justify-center text-primary">
     <span class="material-symbols-outlined text-[18px]">
      devices
     </span>
    </div>
   </div>
   <div class="font-kpi-number text-kpi-number text-text-primary mb-1">
    3
   </div>
   <p class="font-metadata text-metadata text-text-secondary">
    All acknowledged
   </p>
  </div>
  <!-- Card 2: Reservations -->
  <div class="bg-metric-lavender rounded-xl p-card-padding shadow-card hover:-translate-y-0.5 transition-transform cursor-pointer border border-border-default/50">
   <div class="flex items-start justify-between mb-2">
    <h3 class="font-h3 text-h3 text-text-primary">
     Reservations
    </h3>
    <div class="w-8 h-8 rounded-full bg-surface/50 flex items-center justify-center text-status-checked-out">
     <span class="material-symbols-outlined text-[18px]">
      event_available
     </span>
    </div>
   </div>
   <div class="font-kpi-number text-kpi-number text-text-primary mb-1">
    1
   </div>
   <p class="font-metadata text-metadata text-text-secondary">
    Next: 15 May
   </p>
  </div>
  <!-- Card 3: Open Issues -->
  <div class="bg-surface-dim rounded-xl p-card-padding shadow-card hover:-translate-y-0.5 transition-transform cursor-pointer border border-border-default/50">
   <div class="flex items-start justify-between mb-2">
    <h3 class="font-h3 text-h3 text-text-primary">
     Open Issues
    </h3>
    <div class="w-8 h-8 rounded-full bg-surface/50 flex items-center justify-center text-status-critical">
     <span class="material-symbols-outlined text-[18px]">
      report_problem
     </span>
    </div>
   </div>
   <div class="font-kpi-number text-kpi-number text-text-primary mb-1">
    0
   </div>
   <p class="font-metadata text-metadata text-text-secondary">
    No pending reports
   </p>
  </div>
 </div>
 <!-- Two-Column Main Grid (60/40) -->
 <div class="grid grid-cols-1 lg:grid-cols-12 gap-section-gap">
  <!-- Left Column (60% approx -> 7 cols) -->
  <div class="lg:col-span-7 flex flex-col gap-stack">
   <div class="bg-surface border border-border-default rounded-xl shadow-card flex flex-col overflow-hidden">
    <!-- Card Header -->
    <div class="p-card-padding border-b border-border-default flex items-center justify-between bg-surface">
     <h2 class="font-h2 text-h2 text-text-primary">
      My Assigned Assets
     </h2>
     <button class="text-text-secondary hover:text-text-primary transition-colors flex items-center gap-1 font-body text-body text-sm">
      <span class="material-symbols-outlined text-[18px]">
       filter_list
      </span>
      Filter
     </button>
    </div>
    <!-- Table -->
    <div class="overflow-x-auto">
     <table class="w-full text-left border-collapse">
      <thead class="bg-surface-subtle border-b border-border-default">
       <tr>
        <th class="py-3 px-4 font-table-header text-table-header text-text-secondary">
         Asset Name
        </th>
        <th class="py-3 px-4 font-table-header text-table-header text-text-secondary">
         Tag ID
        </th>
        <th class="py-3 px-4 font-table-header text-table-header text-text-secondary">
         Checked Out
        </th>
        <th class="py-3 px-4 font-table-header text-table-header text-text-secondary">
         Status
        </th>
        <th class="py-3 px-4 font-table-header text-table-header text-text-secondary text-right">
         Actions
        </th>
       </tr>
      </thead>
      <tbody class="font-body text-body text-text-primary">
       <!-- Row 1 -->
       <tr class="border-b border-border-default hover:bg-accent-subtle transition-colors group">
        <td class="py-4 px-4">
         <div class="font-h3 text-h3">
          ThinkPad T14 Gen 2
         </div>
         <div class="font-metadata text-metadata text-text-secondary">
          Lenovo Laptop
         </div>
        </td>
        <td class="py-4 px-4 font-mono-data text-mono-data text-text-secondary">
         AST-10492
        </td>
        <td class="py-4 px-4">
         12/01/2023
        </td>
        <td class="py-4 px-4">
         <span class="inline-flex items-center px-2 py-0.5 rounded-DEFAULT bg-metric-sage text-status-in-stock font-table-header text-table-header">
          Active
         </span>
        </td>
        <td class="py-4 px-4 text-right">
         <div class="flex items-center justify-end gap-2 opacity-0 group-hover:opacity-100 transition-opacity">
          <button class="text-text-secondary hover:text-status-critical p-1 rounded-DEFAULT hover:bg-surface transition-colors" title="Report Issue">
           <span class="material-symbols-outlined text-[18px]">
            warning
           </span>
          </button>
          <button class="text-text-secondary hover:text-primary p-1 rounded-DEFAULT hover:bg-surface transition-colors" title="Return Asset">
           <span class="material-symbols-outlined text-[18px]">
            keyboard_return
           </span>
          </button>
         </div>
        </td>
       </tr>
       <!-- Row 2 -->
       <tr class="border-b border-border-default bg-surface-subtle hover:bg-accent-subtle transition-colors group">
        <td class="py-4 px-4">
         <div class="font-h3 text-h3">
          Dell UltraSharp 27"
         </div>
         <div class="font-metadata text-metadata text-text-secondary">
          Monitor
         </div>
        </td>
        <td class="py-4 px-4 font-mono-data text-mono-data text-text-secondary">
         AST-22810
        </td>
        <td class="py-4 px-4">
         15/06/2023
        </td>
        <td class="py-4 px-4">
         <span class="inline-flex items-center px-2 py-0.5 rounded-DEFAULT bg-metric-sage text-status-in-stock font-table-header text-table-header">
          Active
         </span>
        </td>
        <td class="py-4 px-4 text-right">
         <div class="flex items-center justify-end gap-2 opacity-0 group-hover:opacity-100 transition-opacity">
          <button class="text-text-secondary hover:text-status-critical p-1 rounded-DEFAULT hover:bg-surface transition-colors" title="Report Issue">
           <span class="material-symbols-outlined text-[18px]">
            warning
           </span>
          </button>
          <button class="text-text-secondary hover:text-primary p-1 rounded-DEFAULT hover:bg-surface transition-colors" title="Return Asset">
           <span class="material-symbols-outlined text-[18px]">
            keyboard_return
           </span>
          </button>
         </div>
        </td>
       </tr>
       <!-- Row 3 -->
       <tr class="hover:bg-accent-subtle transition-colors group">
        <td class="py-4 px-4">
         <div class="font-h3 text-h3">
          Ergonomic Chair Pro
         </div>
         <div class="font-metadata text-metadata text-text-secondary">
          Furniture
         </div>
        </td>
        <td class="py-4 px-4 font-mono-data text-mono-data text-text-secondary">
         FURN-093
        </td>
        <td class="py-4 px-4">
         02/11/2023
        </td>
        <td class="py-4 px-4">
         <span class="inline-flex items-center px-2 py-0.5 rounded-DEFAULT bg-metric-sage text-status-in-stock font-table-header text-table-header">
          Active
         </span>
        </td>
        <td class="py-4 px-4 text-right">
         <div class="flex items-center justify-end gap-2 opacity-0 group-hover:opacity-100 transition-opacity">
          <button class="text-text-secondary hover:text-status-critical p-1 rounded-DEFAULT hover:bg-surface transition-colors" title="Report Issue">
           <span class="material-symbols-outlined text-[18px]">
            warning
           </span>
          </button>
          <button class="text-text-secondary hover:text-primary p-1 rounded-DEFAULT hover:bg-surface transition-colors" title="Return Asset">
           <span class="material-symbols-outlined text-[18px]">
            keyboard_return
           </span>
          </button>
         </div>
        </td>
       </tr>
      </tbody>
     </table>
    </div>
   </div>
  </div>
  <!-- Right Column (40% approx -> 5 cols) -->
  <div class="lg:col-span-5 flex flex-col gap-stack">
   <div class="bg-surface border border-border-default rounded-xl shadow-card flex flex-col h-full">
    <!-- Card Header -->
    <div class="p-card-padding border-b border-border-default flex items-center justify-between">
     <h2 class="font-h2 text-h2 text-text-primary">
      Recent Requests
     </h2>
     <button class="text-primary hover:underline font-body text-body text-sm">
      View All
     </button>
    </div>
    <!-- List -->
    <div class="p-card-padding flex flex-col gap-4 flex-1">
     <!-- Item 1 -->
     <div class="flex items-start gap-4 p-3 rounded-DEFAULT hover:bg-surface-subtle transition-colors border border-transparent hover:border-border-default cursor-pointer">
      <div class="w-10 h-10 rounded-full bg-surface-dim flex items-center justify-center flex-shrink-0 text-text-secondary mt-1">
       <span class="material-symbols-outlined text-[20px]">
        smartphone
       </span>
      </div>
      <div class="flex-1">
       <div class="flex justify-between items-start mb-1">
        <h3 class="font-h3 text-h3 text-text-primary">
         Test Device iPhone 15
        </h3>
        <span class="font-metadata text-metadata text-text-secondary">
         Today
        </span>
       </div>
       <div class="flex items-center justify-between mt-2">
        <span class="font-mono-data text-mono-data text-text-secondary">
         REQ-8892
        </span>
        <span class="inline-flex items-center px-2 py-0.5 rounded-DEFAULT bg-metric-amber text-primary font-table-header text-table-header">
         Approval Pending
        </span>
       </div>
      </div>
     </div>
     <!-- Item 2 -->
     <div class="flex items-start gap-4 p-3 rounded-DEFAULT hover:bg-surface-subtle transition-colors border border-transparent hover:border-border-default cursor-pointer">
      <div class="w-10 h-10 rounded-full bg-surface-dim flex items-center justify-center flex-shrink-0 text-text-secondary mt-1">
       <span class="material-symbols-outlined text-[20px]">
        headphones
       </span>
      </div>
      <div class="flex-1">
       <div class="flex justify-between items-start mb-1">
        <h3 class="font-h3 text-h3 text-text-primary">
         Noise Cancelling Headset
        </h3>
        <span class="font-metadata text-metadata text-text-secondary">
         12 May
        </span>
       </div>
       <div class="flex items-center justify-between mt-2">
        <span class="font-mono-data text-mono-data text-text-secondary">
         REQ-8750
        </span>
        <span class="inline-flex items-center px-2 py-0.5 rounded-DEFAULT bg-surface-dim text-text-secondary font-table-header text-table-header">
         PO Raised
        </span>
       </div>
      </div>
     </div>
     <!-- Empty State for styling variety -->
     <div class="mt-auto pt-4 border-t border-border-default flex flex-col items-center justify-center py-6 text-center">
      <span class="material-symbols-outlined text-[32px] text-surface-dim mb-2">
       assignment_turned_in
      </span>
      <p class="font-body text-body text-text-secondary mb-3">
       Need software licenses or specific access?
      </p>
      <button class="text-primary font-h3 text-h3 hover:underline">
       Submit Access Request
      </button>
     </div>
    </div>
   </div>
  </div>
 </div>
</main>

</template>

<script setup lang="ts">
// Autogenerated from employee_self_service_portal_2
</script>
