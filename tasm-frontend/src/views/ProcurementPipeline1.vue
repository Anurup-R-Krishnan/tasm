<template>
<main class="flex-1 flex flex-col md:ml-[248px] bg-canvas overflow-hidden">
 <!-- Header & Financial Summary -->
 <div class="flex-none p-page-margin pb-6 border-b border-border-default bg-canvas z-10 shadow-sm">
  <div class="flex items-end justify-between mb-section-gap">
   <div>
    <h1 class="font-h1 text-h1 text-text-primary mb-2">
     Procurement Pipeline
    </h1>
    <p class="font-body text-body text-text-secondary">
     Track and manage active procurement workflows.
    </p>
   </div>
   <div class="flex gap-inline">
    <button class="px-4 py-2 bg-surface border border-outline-variant text-text-primary font-h3 text-h3 rounded-lg hover:bg-surface-subtle transition-colors shadow-sm">
     Export Report
    </button>
    <button class="px-4 py-2 bg-text-primary text-on-primary font-h3 text-h3 rounded-lg hover:bg-stone-800 transition-colors shadow-sm flex items-center gap-2">
     <span class="material-symbols-outlined text-[18px]">
      add
     </span>
     New Request
    </button>
   </div>
  </div>
  <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
   <!-- Metric Card 1 -->
   <div class="bg-surface rounded-xl p-card-padding border border-border-default shadow-[0_4px_12px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform cursor-pointer group">
    <div class="flex justify-between items-start mb-4">
     <span class="font-table-header text-table-header text-text-secondary uppercase">
      Total Pipeline Value
     </span>
     <div class="w-8 h-8 rounded-full bg-secondary-container flex items-center justify-center text-text-primary group-hover:bg-primary-container group-hover:text-on-primary-container transition-colors">
      <span class="material-symbols-outlined text-[18px]">
       account_balance_wallet
      </span>
     </div>
    </div>
    <div class="font-kpi-number text-kpi-number text-text-primary">
     ₹ 42,85,000
    </div>
    <div class="font-metadata text-metadata text-text-secondary mt-2 flex items-center gap-1">
     <span class="text-status-in-stock flex items-center">
      <span class="material-symbols-outlined text-[14px]">
       trending_up
      </span>
      12%
     </span>
     vs last month
    </div>
   </div>
   <!-- Metric Card 2 -->
   <div class="bg-surface rounded-xl p-card-padding border border-border-default shadow-[0_4px_12px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform cursor-pointer group">
    <div class="flex justify-between items-start mb-4">
     <span class="font-table-header text-table-header text-text-secondary uppercase">
      Approved This Month
     </span>
     <div class="w-8 h-8 rounded-full bg-metric-sage/50 flex items-center justify-center text-status-in-stock transition-colors">
      <span class="material-symbols-outlined text-[18px]">
       check_circle
      </span>
     </div>
    </div>
    <div class="font-kpi-number text-kpi-number text-text-primary">
     ₹ 18,20,000
    </div>
    <div class="font-metadata text-metadata text-text-secondary mt-2">
     Across 14 purchase orders
    </div>
   </div>
   <!-- Metric Card 3 (Placeholder for balance visually) -->
   <div class="bg-surface rounded-xl p-card-padding border border-border-default shadow-[0_4px_12px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform cursor-pointer group">
    <div class="flex justify-between items-start mb-4">
     <span class="font-table-header text-table-header text-text-secondary uppercase">
      Pending Approval Value
     </span>
     <div class="w-8 h-8 rounded-full bg-metric-amber/50 flex items-center justify-center text-primary transition-colors">
      <span class="material-symbols-outlined text-[18px]">
       pending_actions
      </span>
     </div>
    </div>
    <div class="font-kpi-number text-kpi-number text-text-primary">
     ₹ 8,45,000
    </div>
    <div class="font-metadata text-metadata text-text-secondary mt-2 text-primary">
     Requires management review
    </div>
   </div>
  </div>
 </div>
 <!-- Kanban Board Area -->
 <div class="flex-1 overflow-x-auto overflow-y-hidden kanban-scroll p-page-margin">
  <div class="flex gap-6 h-full items-start min-w-max">
   <!-- Column 1: Draft (Amber Palette) -->
   <div class="w-[340px] flex flex-col max-h-full">
    <div class="flex items-center justify-between mb-4 px-1">
     <h2 class="font-h2 text-h2 text-text-primary flex items-center gap-2">
      <span class="w-2 h-2 rounded-full bg-[#f59e0b]">
      </span>
      Draft
     </h2>
     <span class="bg-surface border border-border-default px-2 py-0.5 rounded-full font-metadata text-metadata text-text-secondary">
      3
     </span>
    </div>
    <div class="flex flex-col gap-4 overflow-y-auto pr-2 kanban-scroll pb-4">
     <!-- Card: Amber Style -->
     <div class="bg-metric-amber rounded-lg p-4 border border-[#fde68a] shadow-sm hover:-translate-y-[2px] transition-transform cursor-pointer flex flex-col gap-3 relative overflow-hidden">
      <div class="absolute top-0 left-0 w-1 h-full bg-[#f59e0b]">
      </div>
      <div class="flex justify-between items-start pl-1">
       <div class="bg-[#fef08a] text-[#854d0e] px-2 py-1 rounded text-[11px] font-semibold tracking-wider uppercase inline-block">
        Low Priority
       </div>
       <span class="material-symbols-outlined text-text-secondary text-[18px] hover:text-text-primary">
        more_horiz
       </span>
      </div>
      <div class="pl-1">
       <h3 class="font-h3 text-h3 text-text-primary mb-1 line-clamp-2">
        Developer Workstations (x5)
       </h3>
       <div class="font-mono-data text-mono-data text-text-secondary">
        Est: ₹ 6,50,000
       </div>
      </div>
      <div class="flex items-center justify-between mt-2 pt-3 border-t border-[#fde68a]/50 pl-1">
       <div class="flex items-center gap-2">
        <div class="w-6 h-6 rounded-full bg-surface flex items-center justify-center text-[10px] font-bold text-text-primary border border-border-default">
         RJ
        </div>
        <span class="font-metadata text-metadata text-text-secondary">
         Rahul J.
        </span>
       </div>
       <span class="material-symbols-outlined text-text-secondary text-[16px]">
        edit_document
       </span>
      </div>
     </div>
     <!-- Card 2 -->
     <div class="bg-metric-amber rounded-lg p-4 border border-[#fde68a] shadow-sm hover:-translate-y-[2px] transition-transform cursor-pointer flex flex-col gap-3 relative overflow-hidden">
      <div class="absolute top-0 left-0 w-1 h-full bg-[#f59e0b]">
      </div>
      <div class="flex justify-between items-start pl-1">
       <div class="bg-error-container text-status-critical px-2 py-1 rounded text-[11px] font-semibold tracking-wider uppercase inline-block">
        High Priority
       </div>
      </div>
      <div class="pl-1">
       <h3 class="font-h3 text-h3 text-text-primary mb-1 line-clamp-2">
        Emergency Core Switch Replacement
       </h3>
       <div class="font-mono-data text-mono-data text-text-secondary">
        Est: ₹ 1,20,000
       </div>
      </div>
      <div class="flex items-center justify-between mt-2 pt-3 border-t border-[#fde68a]/50 pl-1">
       <div class="flex items-center gap-2">
        <div class="w-6 h-6 rounded-full bg-surface flex items-center justify-center text-[10px] font-bold text-text-primary border border-border-default">
         SN
        </div>
        <span class="font-metadata text-metadata text-text-secondary">
         SysAdmin Net
        </span>
       </div>
      </div>
     </div>
    </div>
   </div>
   <!-- Column 2: Pending Approval -->
   <div class="w-[340px] flex flex-col max-h-full">
    <div class="flex items-center justify-between mb-4 px-1">
     <h2 class="font-h2 text-h2 text-text-primary flex items-center gap-2">
      <span class="w-2 h-2 rounded-full bg-[#64748b]">
      </span>
      Pending Approval
     </h2>
     <span class="bg-surface border border-border-default px-2 py-0.5 rounded-full font-metadata text-metadata text-text-secondary">
      2
     </span>
    </div>
    <div class="flex flex-col gap-4 overflow-y-auto pr-2 kanban-scroll pb-4">
     <!-- Standard Surface Card -->
     <div class="bg-surface rounded-lg p-4 border border-border-default shadow-sm hover:-translate-y-[2px] hover:shadow-md transition-all cursor-pointer flex flex-col gap-3">
      <div class="flex justify-between items-start">
       <div class="bg-surface-variant text-on-surface-variant px-2 py-1 rounded text-[11px] font-semibold tracking-wider uppercase inline-block">
        Medium Priority
       </div>
       <span class="material-symbols-outlined text-text-secondary text-[18px] hover:text-text-primary">
        more_horiz
       </span>
      </div>
      <div>
       <h3 class="font-h3 text-h3 text-text-primary mb-1 line-clamp-2">
        Ergonomic Office Chairs Phase 2
       </h3>
       <div class="font-mono-data text-mono-data text-text-primary font-medium">
        Est: ₹ 4,80,000
       </div>
      </div>
      <div class="flex items-center justify-between mt-2 pt-3 border-t border-border-default">
       <div class="flex items-center gap-2">
        <div class="w-6 h-6 rounded-full bg-secondary-container flex items-center justify-center text-[10px] font-bold text-text-primary border border-border-default">
         HR
        </div>
        <span class="font-metadata text-metadata text-text-secondary">
         Facilities
        </span>
       </div>
       <div class="flex items-center gap-1 text-primary">
        <span class="material-symbols-outlined text-[14px]">
         schedule
        </span>
        <span class="font-metadata text-metadata">
         3 days
        </span>
       </div>
      </div>
     </div>
    </div>
   </div>
   <!-- Column 3: PO Raised (Lavender Palette) -->
   <div class="w-[340px] flex flex-col max-h-full">
    <div class="flex items-center justify-between mb-4 px-1">
     <h2 class="font-h2 text-h2 text-text-primary flex items-center gap-2">
      <span class="w-2 h-2 rounded-full bg-[#8b5cf6]">
      </span>
      PO Raised
     </h2>
     <span class="bg-surface border border-border-default px-2 py-0.5 rounded-full font-metadata text-metadata text-text-secondary">
      1
     </span>
    </div>
    <div class="flex flex-col gap-4 overflow-y-auto pr-2 kanban-scroll pb-4">
     <!-- Card: Lavender Style -->
     <div class="bg-metric-lavender rounded-lg p-4 border border-[#ddd6fe] shadow-sm hover:-translate-y-[2px] transition-transform cursor-pointer flex flex-col gap-3 relative overflow-hidden">
      <div class="absolute top-0 left-0 w-1 h-full bg-[#8b5cf6]">
      </div>
      <div class="flex justify-between items-start pl-1">
       <div class="bg-error-container text-status-critical px-2 py-1 rounded text-[11px] font-semibold tracking-wider uppercase inline-block">
        High Priority
       </div>
       <span class="material-symbols-outlined text-text-secondary text-[18px] hover:text-text-primary">
        more_horiz
       </span>
      </div>
      <div class="pl-1">
       <h3 class="font-h3 text-h3 text-text-primary mb-1 line-clamp-2">
        Enterprise Storage Array Renewal
       </h3>
       <div class="font-mono-data text-mono-data text-text-primary font-medium">
        Est: ₹ 14,50,000
       </div>
       <div class="font-metadata text-metadata text-tertiary mt-1">
        PO-2023-089
       </div>
      </div>
      <div class="flex items-center justify-between mt-2 pt-3 border-t border-[#ddd6fe]/50 pl-1">
       <div class="flex items-center gap-2">
        <div class="w-6 h-6 rounded-full bg-surface flex items-center justify-center text-[10px] font-bold text-text-primary border border-border-default">
         IT
        </div>
        <span class="font-metadata text-metadata text-text-secondary">
         IT Dept.
        </span>
       </div>
       <span class="material-symbols-outlined text-text-secondary text-[16px]">
        receipt_long
       </span>
      </div>
     </div>
    </div>
   </div>
   <!-- Column 4: Shipping -->
   <div class="w-[340px] flex flex-col max-h-full">
    <div class="flex items-center justify-between mb-4 px-1">
     <h2 class="font-h2 text-h2 text-text-primary flex items-center gap-2">
      <span class="w-2 h-2 rounded-full bg-[#3b82f6]">
      </span>
      Shipping
     </h2>
     <span class="bg-surface border border-border-default px-2 py-0.5 rounded-full font-metadata text-metadata text-text-secondary">
      0
     </span>
    </div>
    <div class="flex flex-col gap-4 overflow-y-auto pr-2 kanban-scroll pb-4">
     <!-- Empty State -->
     <div class="border-2 border-dashed border-border-default rounded-lg p-6 flex flex-col items-center justify-center text-center gap-2 h-[120px]">
      <span class="material-symbols-outlined text-outline-variant text-[24px]">
       local_shipping
      </span>
      <span class="font-metadata text-metadata text-text-secondary">
       No items in transit
      </span>
     </div>
    </div>
   </div>
   <!-- Column 5: Received (Sage Palette) -->
   <div class="w-[340px] flex flex-col max-h-full">
    <div class="flex items-center justify-between mb-4 px-1">
     <h2 class="font-h2 text-h2 text-text-primary flex items-center gap-2">
      <span class="w-2 h-2 rounded-full bg-[#22c55e]">
      </span>
      Received
     </h2>
     <span class="bg-surface border border-border-default px-2 py-0.5 rounded-full font-metadata text-metadata text-text-secondary">
      1
     </span>
    </div>
    <div class="flex flex-col gap-4 overflow-y-auto pr-2 kanban-scroll pb-4">
     <!-- Card: Sage Style -->
     <div class="bg-metric-sage rounded-lg p-4 border border-[#bbf7d0] shadow-sm hover:-translate-y-[2px] transition-transform cursor-pointer flex flex-col gap-3 relative overflow-hidden">
      <div class="absolute top-0 left-0 w-1 h-full bg-[#22c55e]">
      </div>
      <div class="flex justify-between items-start pl-1">
       <div class="bg-surface-variant text-on-surface-variant px-2 py-1 rounded text-[11px] font-semibold tracking-wider uppercase inline-block">
        Medium Priority
       </div>
       <span class="material-symbols-outlined text-text-secondary text-[18px] hover:text-text-primary">
        more_horiz
       </span>
      </div>
      <div class="pl-1">
       <h3 class="font-h3 text-h3 text-text-primary mb-1 line-clamp-2">
        Conference Room Displays (x3)
       </h3>
       <div class="font-mono-data text-mono-data text-text-primary font-medium">
        Actual: ₹ 3,10,000
       </div>
       <div class="font-metadata text-metadata text-status-in-stock mt-1 flex items-center gap-1">
        <span class="material-symbols-outlined text-[14px]">
         inventory_2
        </span>
        Awaiting Registry
       </div>
      </div>
      <div class="flex items-center justify-between mt-2 pt-3 border-t border-[#bbf7d0]/50 pl-1">
       <div class="flex items-center gap-2">
        <div class="w-6 h-6 rounded-full bg-surface flex items-center justify-center text-[10px] font-bold text-text-primary border border-border-default">
         HR
        </div>
        <span class="font-metadata text-metadata text-text-secondary">
         Facilities
        </span>
       </div>
       <button class="text-xs font-semibold text-status-in-stock hover:underline">
        Add to Registry
       </button>
      </div>
     </div>
    </div>
   </div>
  </div>
 </div>
</main>

</template>

<script setup lang="ts">
// Autogenerated from procurement_pipeline_1
</script>
