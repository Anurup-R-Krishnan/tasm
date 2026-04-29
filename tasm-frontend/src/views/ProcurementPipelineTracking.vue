<template>
<main class="md:ml-[248px] p-page-margin min-h-screen">
 <div class="max-w-[1400px] mx-auto space-y-section-gap">
  <!-- Header Area -->
  <div class="flex justify-between items-end">
   <div>
    <h1 class="font-h1 text-h1 text-text-primary tracking-tight">
     Procurement Pipeline
    </h1>
    <p class="font-body text-body text-text-secondary mt-1">
     Track and manage active purchase requests through delivery.
    </p>
   </div>
   <div class="flex gap-3">
    <button class="bg-surface border border-border-default text-text-primary px-4 py-2 rounded font-medium hover:-translate-y-[2px] hover:shadow-sm transition-all flex items-center gap-2">
     <span class="material-symbols-outlined" style="font-size: 18px;">
      filter_list
     </span>
     Filter
    </button>
    <button class="bg-text-primary text-surface px-4 py-2 rounded font-medium hover:-translate-y-[2px] hover:shadow-sm transition-all">
     Export Report
    </button>
   </div>
  </div>
  <!-- Visual Pipeline Aggregation -->
  <div class="grid grid-cols-5 gap-4">
   <div class="bg-surface rounded-xl border border-border-default p-card-padding shadow-[0_4px_4px_0_rgba(0,0,0,0.02)] hover:-translate-y-[2px] transition-transform relative overflow-hidden group">
    <div class="absolute left-0 top-0 bottom-0 w-1 bg-surface-dim">
    </div>
    <div class="font-metadata text-metadata text-text-secondary uppercase tracking-wider mb-2">
     Draft Phase
    </div>
    <div class="font-kpi-number text-kpi-number text-text-primary">
     12
    </div>
    <div class="font-metadata text-metadata text-text-secondary mt-1">
     ₹4.2L est.
    </div>
   </div>
   <div class="bg-surface rounded-xl border border-border-default p-card-padding shadow-[0_4px_4px_0_rgba(0,0,0,0.02)] hover:-translate-y-[2px] transition-transform relative overflow-hidden group">
    <div class="absolute left-0 top-0 bottom-0 w-1 bg-metric-amber">
    </div>
    <div class="font-metadata text-metadata text-text-secondary uppercase tracking-wider mb-2">
     Pending Approval
    </div>
    <div class="font-kpi-number text-kpi-number text-text-primary">
     8
    </div>
    <div class="font-metadata text-metadata text-text-secondary mt-1">
     ₹12.5L est.
    </div>
   </div>
   <div class="bg-surface rounded-xl border border-border-default p-card-padding shadow-[0_4px_4px_0_rgba(0,0,0,0.02)] hover:-translate-y-[2px] transition-transform relative overflow-hidden group">
    <div class="absolute left-0 top-0 bottom-0 w-1 bg-metric-lavender">
    </div>
    <div class="font-metadata text-metadata text-text-secondary uppercase tracking-wider mb-2">
     PO Issued
    </div>
    <div class="font-kpi-number text-kpi-number text-text-primary">
     15
    </div>
    <div class="font-metadata text-metadata text-text-secondary mt-1">
     ₹34.8L committed
    </div>
   </div>
   <div class="bg-surface rounded-xl border border-border-default p-card-padding shadow-[0_4px_4px_0_rgba(0,0,0,0.02)] hover:-translate-y-[2px] transition-transform relative overflow-hidden group">
    <div class="absolute left-0 top-0 bottom-0 w-1 bg-metric-sage">
    </div>
    <div class="font-metadata text-metadata text-text-secondary uppercase tracking-wider mb-2">
     Awaiting Receipt
    </div>
    <div class="font-kpi-number text-kpi-number text-text-primary">
     4
    </div>
    <div class="font-metadata text-metadata text-text-secondary mt-1">
     Next delivery in 2d
    </div>
   </div>
   <div class="bg-surface rounded-xl border border-border-default p-card-padding shadow-[0_4px_4px_0_rgba(0,0,0,0.02)] hover:-translate-y-[2px] transition-transform relative overflow-hidden group">
    <div class="absolute left-0 top-0 bottom-0 w-1 bg-status-in-stock">
    </div>
    <div class="font-metadata text-metadata text-text-secondary uppercase tracking-wider mb-2">
     Ready to Register
    </div>
    <div class="font-kpi-number text-kpi-number text-text-primary">
     7
    </div>
    <div class="font-metadata text-metadata text-text-secondary mt-1">
     Items in holding
    </div>
   </div>
  </div>
  <!-- High Density Data Table -->
  <div class="bg-surface rounded-xl border border-border-default shadow-[0_4px_4px_0_rgba(0,0,0,0.02)] overflow-hidden">
   <div class="px-6 py-4 border-b border-border-default bg-surface-subtle flex justify-between items-center">
    <h2 class="font-h2 text-h2 text-text-primary">
     Active Procurement Queue
    </h2>
    <div class="text-sm text-text-secondary">
     Showing 1-10 of 46 entries
    </div>
   </div>
   <div class="overflow-x-auto">
    <div class="absolute left-0 top-0 bottom-0 w-[3px] bg-metric-lavender opacity-0 group-hover:opacity-100 transition-opacity">
    </div>
    <table class="w-full text-left border-collapse">
     <thead>
      <tr class="bg-surface-subtle border-b border-border-default">
       <th class="font-table-header text-table-header text-text-secondary py-3 px-6 w-24">
        PR ID
       </th>
       <th class="font-table-header text-table-header text-text-secondary py-3 px-6">
        CATEGORY
       </th>
       <th class="font-table-header text-table-header text-text-secondary py-3 px-6 w-20 text-right">
        QTY
       </th>
       <th class="font-table-header text-table-header text-text-secondary py-3 px-6 w-32 text-right">
        EST. COST
       </th>
       <th class="font-table-header text-table-header text-text-secondary py-3 px-6 w-40">
        EXPECTED
       </th>
       <th class="font-table-header text-table-header text-text-secondary py-3 px-6 w-56">
        PIPELINE STAGE
       </th>
      </tr>
     </thead>
     <tbody class="font-body text-body">
      <!-- Row 1: Draft -->
      <tr class="border-b border-border-default bg-surface hover:bg-accent-subtle transition-colors group">
       <td class="py-3 px-6 font-mono-data text-mono-data text-text-primary">
        PR-8092
       </td>
       <td class="py-3 px-6">
        Networking Gear
       </td>
       <td class="py-3 px-6 text-right font-mono-data text-mono-data">
        12
       </td>
       <td class="py-3 px-6 text-right font-mono-data text-mono-data">
        ₹ 4,50,000
       </td>
       <td class="py-3 px-6 text-text-secondary">
        TBD
       </td>
       <td class="py-3 px-6">
        <div class="flex items-center gap-2">
         <div class="w-2.5 h-2.5 rounded-full bg-surface-dim border border-border-default">
         </div>
         <span class="text-sm font-medium text-text-secondary">
          Draft
         </span>
        </div>
        <div class="w-full bg-surface-subtle h-1 rounded-full mt-1.5 overflow-hidden flex">
         <div class="bg-surface-dim w-1/5 h-full">
         </div>
        </div>
       </td>
      </tr>
      <!-- Row 2: Approval -->
      <tr class="border-b border-border-default bg-surface-subtle hover:bg-accent-subtle transition-colors group">
       <td class="py-3 px-6 font-mono-data text-mono-data text-text-primary">
        PR-8088
       </td>
       <td class="py-3 px-6">
        Office Furniture
       </td>
       <td class="py-3 px-6 text-right font-mono-data text-mono-data">
        45
       </td>
       <td class="py-3 px-6 text-right font-mono-data text-mono-data">
        ₹ 8,25,000
       </td>
       <td class="py-3 px-6 text-text-secondary">
        TBD
       </td>
       <td class="py-3 px-6">
        <div class="flex items-center gap-2">
         <div class="w-2.5 h-2.5 rounded-full bg-metric-amber border border-amber-200">
         </div>
         <span class="text-sm font-medium text-amber-800">
          Pending Approval
         </span>
        </div>
        <div class="w-full bg-surface-dim h-1 rounded-full mt-1.5 overflow-hidden flex">
         <div class="bg-metric-amber w-2/5 h-full">
         </div>
        </div>
       </td>
      </tr>
      <!-- Row 3: PO Issued (Highlighted) -->
      <tr class="border-b border-border-default bg-surface hover:bg-accent-subtle transition-colors group relative">
       <td class="py-3 px-6 font-mono-data text-mono-data text-text-primary">
        PR-8075
       </td>
       <td class="py-3 px-6 font-medium">
        Dell Latitude Laptops
       </td>
       <td class="py-3 px-6 text-right font-mono-data text-mono-data">
        20
       </td>
       <td class="py-3 px-6 text-right font-mono-data text-mono-data">
        ₹ 14,00,000
       </td>
       <td class="py-3 px-6">
        24 Oct 2023
       </td>
       <td class="py-3 px-6">
        <div class="flex items-center gap-2">
         <div class="w-2.5 h-2.5 rounded-full bg-metric-lavender border border-purple-200">
         </div>
         <span class="text-sm font-medium text-purple-800">
          PO Issued
         </span>
        </div>
        <div class="w-full bg-surface-dim h-1 rounded-full mt-1.5 overflow-hidden flex">
         <div class="bg-metric-lavender w-3/5 h-full">
         </div>
        </div>
       </td>
       <!-- Highlight Indicator -->
      </tr>
      <!-- Row 4: Receipt -->
      <tr class="border-b border-border-default bg-surface-subtle hover:bg-accent-subtle transition-colors group">
       <td class="py-3 px-6 font-mono-data text-mono-data text-text-primary">
        PR-8062
       </td>
       <td class="py-3 px-6">
        A4 Printer Paper
       </td>
       <td class="py-3 px-6 text-right font-mono-data text-mono-data">
        100
       </td>
       <td class="py-3 px-6 text-right font-mono-data text-mono-data">
        ₹ 25,000
       </td>
       <td class="py-3 px-6 text-status-critical font-medium">
        18 Oct 2023
       </td>
       <td class="py-3 px-6">
        <div class="flex items-center gap-2">
         <div class="w-2.5 h-2.5 rounded-full bg-metric-sage border border-green-200">
         </div>
         <span class="text-sm font-medium text-green-800">
          Awaiting Receipt
         </span>
        </div>
        <div class="w-full bg-surface-dim h-1 rounded-full mt-1.5 overflow-hidden flex">
         <div class="bg-metric-sage w-4/5 h-full">
         </div>
        </div>
       </td>
      </tr>
      <!-- Row 5: Register -->
      <tr class="border-b border-border-default bg-surface hover:bg-accent-subtle transition-colors group">
       <td class="py-3 px-6 font-mono-data text-mono-data text-text-primary">
        PR-8055
       </td>
       <td class="py-3 px-6">
        Conference Room Displays
       </td>
       <td class="py-3 px-6 text-right font-mono-data text-mono-data">
        4
       </td>
       <td class="py-3 px-6 text-right font-mono-data text-mono-data">
        ₹ 3,20,000
       </td>
       <td class="py-3 px-6 text-text-secondary">
        Arrived
       </td>
       <td class="py-3 px-6">
        <div class="flex items-center gap-2">
         <div class="w-2.5 h-2.5 rounded-full bg-status-in-stock">
         </div>
         <span class="text-sm font-medium text-status-in-stock">
          Ready to Register
         </span>
        </div>
        <div class="w-full bg-surface-dim h-1 rounded-full mt-1.5 overflow-hidden flex">
         <div class="bg-status-in-stock w-full h-full">
         </div>
        </div>
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
// Autogenerated from procurement_pipeline_tracking
</script>
