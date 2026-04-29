<template>
<main class="flex-1 md:ml-[248px] p-page-margin max-w-[1400px] w-full mx-auto">
 <!-- Page Header -->
 <div class="flex justify-between items-end mb-section-gap">
  <div>
   <h1 class="font-h1 text-h1 text-text-primary">
    Procurement Tracker
   </h1>
   <p class="font-body text-body text-text-secondary mt-1">
    Manage and track high-value asset acquisitions.
   </p>
  </div>
  <button class="bg-[#1C1917] text-white px-4 py-2 rounded-DEFAULT font-h3 text-h3 flex items-center gap-2 hover:bg-black transition-colors shadow-sm transform hover:-translate-y-0.5">
   <span class="material-symbols-outlined text-[18px]">
    add_circle
   </span>
   New Request
  </button>
 </div>
 <!-- KPI Cards -->
 <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-section-gap">
  <div class="bg-metric-amber rounded-[16px] p-card-padding shadow-sm border border-border-default cursor-pointer transform hover:-translate-y-0.5 transition-transform">
   <div class="flex justify-between items-start mb-2">
    <span class="font-h3 text-h3 text-amber-900">
     Active Requests
    </span>
    <span class="material-symbols-outlined text-amber-700" style="font-variation-settings: 'FILL' 1;">
     pending_actions
    </span>
   </div>
   <div class="font-kpi-number text-kpi-number text-amber-900 mb-1">
    42
   </div>
   <div class="font-metadata text-metadata text-amber-800 flex items-center gap-1">
    <span class="material-symbols-outlined text-[14px]">
     arrow_upward
    </span>
    12% from last month
   </div>
  </div>
  <div class="bg-metric-lavender rounded-[16px] p-card-padding shadow-sm border border-border-default cursor-pointer transform hover:-translate-y-0.5 transition-transform">
   <div class="flex justify-between items-start mb-2">
    <span class="font-h3 text-h3 text-indigo-900">
     Awaiting Approval
    </span>
    <span class="material-symbols-outlined text-indigo-700" style="font-variation-settings: 'FILL' 1;">
     fact_check
    </span>
   </div>
   <div class="font-kpi-number text-kpi-number text-indigo-900 mb-1">
    18
   </div>
   <div class="font-metadata text-metadata text-indigo-800 flex items-center gap-1">
    <span class="material-symbols-outlined text-[14px]">
     priority_high
    </span>
    Requires immediate attention
   </div>
  </div>
  <div class="bg-metric-sage rounded-[16px] p-card-padding shadow-sm border border-border-default cursor-pointer transform hover:-translate-y-0.5 transition-transform">
   <div class="flex justify-between items-start mb-2">
    <span class="font-h3 text-h3 text-green-900">
     Total Received YTD
    </span>
    <span class="material-symbols-outlined text-green-700" style="font-variation-settings: 'FILL' 1;">
     inventory
    </span>
   </div>
   <div class="font-kpi-number text-kpi-number text-green-900 mb-1">
    ₹4.2M
   </div>
   <div class="font-metadata text-metadata text-green-800 flex items-center gap-1">
    <span class="material-symbols-outlined text-[14px]">
     check_circle
    </span>
    Across 156 separate POs
   </div>
  </div>
 </div>
 <!-- Table Section -->
 <div class="bg-surface rounded-lg shadow-sm border border-border-default overflow-hidden">
  <div class="px-6 py-4 border-b border-border-default flex justify-between items-center bg-surface-subtle">
   <h2 class="font-h2 text-h2 text-text-primary">
    Pipeline View
   </h2>
   <div class="flex gap-2">
    <button class="p-1.5 border border-border-default rounded bg-surface hover:bg-surface-subtle transition-colors text-text-secondary">
     <span class="material-symbols-outlined text-[20px]">
      filter_list
     </span>
    </button>
    <button class="p-1.5 border border-border-default rounded bg-surface hover:bg-surface-subtle transition-colors text-text-secondary">
     <span class="material-symbols-outlined text-[20px]">
      download
     </span>
    </button>
   </div>
  </div>
  <div class="overflow-x-auto">
   <table class="w-full text-left border-collapse">
    <thead>
     <tr class="bg-surface-subtle border-b border-border-default">
      <th class="px-6 py-3 font-table-header text-table-header text-text-secondary uppercase">
       PR ID
      </th>
      <th class="px-6 py-3 font-table-header text-table-header text-text-secondary uppercase">
       Category
      </th>
      <th class="px-6 py-3 font-table-header text-table-header text-text-secondary uppercase">
       Qty
      </th>
      <th class="px-6 py-3 font-table-header text-table-header text-text-secondary uppercase">
       Est. Cost (₹)
      </th>
      <th class="px-6 py-3 font-table-header text-table-header text-text-secondary uppercase">
       Raised By
      </th>
      <th class="px-6 py-3 font-table-header text-table-header text-text-secondary uppercase w-[300px]">
       Pipeline Status
      </th>
      <th class="px-6 py-3 font-table-header text-table-header text-text-secondary uppercase">
       PO No.
      </th>
      <th class="px-6 py-3 font-table-header text-table-header text-text-secondary uppercase text-right">
       Actions
      </th>
     </tr>
    </thead>
    <tbody class="divide-y divide-border-default">
     <!-- Row 1 -->
     <tr class="bg-surface hover:bg-metric-amber/30 transition-colors border-l-4 border-transparent hover:border-amber-500 group">
      <td class="px-6 py-4 font-mono-data text-mono-data text-text-primary">
       PR-2023-089
      </td>
      <td class="px-6 py-4 font-body text-body text-text-primary">
       IT Hardware
      </td>
      <td class="px-6 py-4 font-body text-body text-text-secondary">
       25
      </td>
      <td class="px-6 py-4 font-body text-body text-text-primary">
       1,250,000
      </td>
      <td class="px-6 py-4 font-body text-body text-text-secondary">
       Arun K.
      </td>
      <td class="px-6 py-4">
       <div class="flex items-center gap-1 text-[11px] font-medium">
        <span class="w-2 h-2 rounded-full bg-green-500 mr-1">
        </span>
        <span class="text-green-700 bg-green-50 px-1.5 py-0.5 rounded">
         Draft
        </span>
        <span class="material-symbols-outlined text-[14px] text-stone-300">
         chevron_right
        </span>
        <span class="w-2 h-2 rounded-full bg-green-500 mr-1">
        </span>
        <span class="text-green-700 bg-green-50 px-1.5 py-0.5 rounded">
         Dept
        </span>
        <span class="material-symbols-outlined text-[14px] text-stone-300">
         chevron_right
        </span>
        <span class="w-2 h-2 rounded-full bg-amber-500 mr-1">
        </span>
        <span class="text-amber-800 bg-amber-100 px-1.5 py-0.5 rounded font-bold border border-amber-200 shadow-sm">
         Finance
        </span>
        <span class="material-symbols-outlined text-[14px] text-stone-300">
         chevron_right
        </span>
        <span class="w-1.5 h-1.5 rounded-full bg-stone-300 mr-1">
        </span>
        <span class="text-stone-400">
         PO
        </span>
       </div>
      </td>
      <td class="px-6 py-4 font-mono-data text-mono-data text-text-secondary">
       -
      </td>
      <td class="px-6 py-4 text-right">
       <button class="text-text-secondary hover:text-amber-700 transition-colors opacity-0 group-hover:opacity-100">
        <span class="material-symbols-outlined text-[20px]">
         more_vert
        </span>
       </button>
      </td>
     </tr>
     <!-- Row 2 -->
     <tr class="bg-surface-subtle hover:bg-metric-amber/30 transition-colors border-l-4 border-transparent hover:border-amber-500 group">
      <td class="px-6 py-4 font-mono-data text-mono-data text-text-primary">
       PR-2023-085
      </td>
      <td class="px-6 py-4 font-body text-body text-text-primary">
       Office Furniture
      </td>
      <td class="px-6 py-4 font-body text-body text-text-secondary">
       10
      </td>
      <td class="px-6 py-4 font-body text-body text-text-primary">
       450,000
      </td>
      <td class="px-6 py-4 font-body text-body text-text-secondary">
       Priya M.
      </td>
      <td class="px-6 py-4">
       <div class="flex items-center gap-1 text-[11px] font-medium">
        <span class="w-2 h-2 rounded-full bg-green-500 mr-1">
        </span>
        <span class="text-green-700 bg-green-50 px-1.5 py-0.5 rounded">
         Draft
        </span>
        <span class="material-symbols-outlined text-[14px] text-stone-300">
         chevron_right
        </span>
        <span class="w-2 h-2 rounded-full bg-amber-500 mr-1">
        </span>
        <span class="text-amber-800 bg-amber-100 px-1.5 py-0.5 rounded font-bold border border-amber-200 shadow-sm">
         Dept
        </span>
        <span class="material-symbols-outlined text-[14px] text-stone-300">
         chevron_right
        </span>
        <span class="w-1.5 h-1.5 rounded-full bg-stone-300 mr-1">
        </span>
        <span class="text-stone-400">
         Finance
        </span>
       </div>
      </td>
      <td class="px-6 py-4 font-mono-data text-mono-data text-text-secondary">
       -
      </td>
      <td class="px-6 py-4 text-right">
       <button class="text-text-secondary hover:text-amber-700 transition-colors opacity-0 group-hover:opacity-100">
        <span class="material-symbols-outlined text-[20px]">
         more_vert
        </span>
       </button>
      </td>
     </tr>
     <!-- Row 3 (Selected State Example) -->
     <tr class="bg-metric-amber/10 border-l-4 border-amber-500 group">
      <td class="px-6 py-4 font-mono-data text-mono-data text-text-primary">
       PR-2023-072
      </td>
      <td class="px-6 py-4 font-body text-body text-text-primary">
       Network Switches
      </td>
      <td class="px-6 py-4 font-body text-body text-text-secondary">
       5
      </td>
      <td class="px-6 py-4 font-body text-body text-text-primary">
       850,000
      </td>
      <td class="px-6 py-4 font-body text-body text-text-secondary">
       Arun K.
      </td>
      <td class="px-6 py-4">
       <div class="flex items-center gap-1 text-[11px] font-medium">
        <span class="w-2 h-2 rounded-full bg-green-500 mr-1">
        </span>
        <span class="text-green-700">
         Draft
        </span>
        <span class="material-symbols-outlined text-[14px] text-stone-300">
         chevron_right
        </span>
        <span class="w-2 h-2 rounded-full bg-green-500 mr-1">
        </span>
        <span class="text-green-700">
         Dept
        </span>
        <span class="material-symbols-outlined text-[14px] text-stone-300">
         chevron_right
        </span>
        <span class="w-2 h-2 rounded-full bg-green-500 mr-1">
        </span>
        <span class="text-green-700">
         Finance
        </span>
        <span class="material-symbols-outlined text-[14px] text-stone-300">
         chevron_right
        </span>
        <span class="w-2 h-2 rounded-full bg-indigo-500 mr-1">
        </span>
        <span class="text-indigo-800 bg-indigo-100 px-1.5 py-0.5 rounded font-bold border border-indigo-200 shadow-sm">
         PO Raised
        </span>
       </div>
      </td>
      <td class="px-6 py-4 font-mono-data text-mono-data text-text-primary font-medium">
       PO-9942
      </td>
      <td class="px-6 py-4 text-right">
       <button class="text-amber-700 hover:text-amber-900 transition-colors">
        <span class="material-symbols-outlined text-[20px]">
         more_vert
        </span>
       </button>
      </td>
     </tr>
    </tbody>
   </table>
  </div>
  <div class="px-6 py-3 border-t border-border-default flex justify-between items-center bg-surface-subtle font-metadata text-metadata text-text-secondary">
   <span>
    Showing 1 to 3 of 42 entries
   </span>
   <div class="flex gap-1">
    <button class="px-2 py-1 border border-border-default rounded bg-surface disabled:opacity-50" disabled="">
     Prev
    </button>
    <button class="px-2 py-1 border border-amber-500 rounded bg-amber-50 text-amber-900 font-medium">
     1
    </button>
    <button class="px-2 py-1 border border-border-default rounded bg-surface hover:bg-surface-subtle">
     2
    </button>
    <button class="px-2 py-1 border border-border-default rounded bg-surface hover:bg-surface-subtle">
     3
    </button>
    <button class="px-2 py-1 border border-border-default rounded bg-surface hover:bg-surface-subtle">
     Next
    </button>
   </div>
  </div>
 </div>
</main>

</template>

<script setup lang="ts">
// Autogenerated from procurement_pipeline_tracker
</script>
