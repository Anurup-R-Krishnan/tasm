<template>
<main class="mt-[60px] p-page-margin w-[calc(100%-248px)] max-w-[1400px] flex flex-col gap-section-gap">
 <!-- Page Header -->
 <div class="flex items-center justify-between">
  <h1 class="font-h1 text-h1 text-text-primary">
   Procurement Pipeline
  </h1>
  <button class="bg-[#1C1917] text-white px-4 py-2 rounded-lg font-h3 text-h3 flex items-center gap-2 hover:bg-stone-800 transition-colors shadow-sm">
   <span class="material-symbols-outlined" style="font-variation-settings: 'FILL' 1;">
    add
   </span>
   New Request
  </button>
 </div>
 <!-- Progress Pipeline Viz -->
 <div class="bg-surface border border-border-default rounded-2xl p-card-padding shadow-[0_4px_12px_-4px_rgba(0,0,0,0.05)]">
  <h3 class="font-h3 text-h3 text-text-primary mb-6">
   Active Requisitions Flow
  </h3>
  <div class="flex items-center justify-between relative">
   <!-- Connecting Line Base -->
   <div class="absolute left-8 right-8 top-5 h-[2px] bg-border-default -z-10">
   </div>
   <!-- Connecting Line Progress -->
   <div class="absolute left-8 top-5 h-[2px] bg-primary w-[60%] -z-10">
   </div>
   <!-- Stage: Draft -->
   <div class="flex flex-col items-center gap-2 w-32">
    <div class="w-10 h-10 rounded-full bg-primary text-white flex items-center justify-center font-h3 text-h3 shadow-md">
     12
    </div>
    <span class="font-metadata text-metadata text-text-primary font-medium">
     Draft
    </span>
   </div>
   <!-- Stage: Dept Approval -->
   <div class="flex flex-col items-center gap-2 w-32">
    <div class="w-10 h-10 rounded-full bg-primary text-white flex items-center justify-center font-h3 text-h3 shadow-md">
     8
    </div>
    <span class="font-metadata text-metadata text-text-primary font-medium">
     Dept Approval
    </span>
   </div>
   <!-- Stage: Finance Approval -->
   <div class="flex flex-col items-center gap-2 w-32">
    <div class="w-10 h-10 rounded-full bg-primary text-white flex items-center justify-center font-h3 text-h3 shadow-md">
     4
    </div>
    <span class="font-metadata text-metadata text-text-primary font-medium">
     Finance Approval
    </span>
   </div>
   <!-- Stage: PO Raised -->
   <div class="flex flex-col items-center gap-2 w-32">
    <div class="w-10 h-10 rounded-full bg-surface border-2 border-primary text-primary flex items-center justify-center font-h3 text-h3 bg-white">
     15
    </div>
    <span class="font-metadata text-metadata text-text-primary font-medium">
     PO Raised
    </span>
   </div>
   <!-- Stage: Received -->
   <div class="flex flex-col items-center gap-2 w-32">
    <div class="w-10 h-10 rounded-full bg-surface border-2 border-border-default text-text-secondary flex items-center justify-center font-h3 text-h3 bg-white">
     42
    </div>
    <span class="font-metadata text-metadata text-text-secondary">
     Received (MTD)
    </span>
   </div>
  </div>
 </div>
 <!-- Main Content Area: Filter & Table -->
 <div class="flex flex-col gap-4">
  <!-- Filter Bar -->
  <div class="flex items-center gap-3">
   <div class="relative flex-1 max-w-md">
    <span class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-stone-400 text-[18px]">
     search
    </span>
    <input class="w-full pl-9 pr-4 py-2 bg-white border border-border-default rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-amber-500/20 focus:border-amber-500 placeholder:text-stone-400 shadow-sm" placeholder="Search PR ID, Item, or Vendor..." type="text"/>
   </div>
   <select class="py-2 pl-3 pr-8 bg-white border border-border-default rounded-lg text-sm text-text-primary focus:outline-none focus:ring-2 focus:ring-amber-500/20 shadow-sm">
    <option>
     All Statuses
    </option>
    <option>
     Draft
    </option>
    <option>
     Pending Dept
    </option>
    <option>
     Pending Finance
    </option>
    <option>
     PO Raised
    </option>
   </select>
   <select class="py-2 pl-3 pr-8 bg-white border border-border-default rounded-lg text-sm text-text-primary focus:outline-none focus:ring-2 focus:ring-amber-500/20 shadow-sm">
    <option>
     All Departments
    </option>
    <option>
     IT Infrastructure
    </option>
    <option>
     Facilities
    </option>
    <option>
     Security
    </option>
   </select>
   <div class="flex items-center gap-2 ml-auto">
    <button class="bg-white border border-border-default text-text-primary px-3 py-2 rounded-lg text-sm flex items-center gap-2 hover:bg-stone-50 transition-colors shadow-sm">
     <span class="material-symbols-outlined text-[18px]">
      calendar_today
     </span>
     Last 30 Days
    </button>
    <button class="bg-white border border-border-default text-text-primary px-3 py-2 rounded-lg text-sm flex items-center gap-2 hover:bg-stone-50 transition-colors shadow-sm">
     <span class="material-symbols-outlined text-[18px]">
      filter_list
     </span>
     More Filters
    </button>
   </div>
  </div>
  <!-- Data Table Card -->
  <div class="bg-surface border border-border-default rounded-2xl shadow-[0_4px_4px_rgba(0,0,0,0.02)] overflow-hidden">
   <div class="overflow-x-auto w-full">
    <table class="w-full text-left border-collapse">
     <thead>
      <tr class="border-b border-border-default bg-[#FAFAF8]">
       <th class="font-table-header text-table-header text-text-secondary py-4 px-6 font-medium whitespace-nowrap">
        PR ID
       </th>
       <th class="font-table-header text-table-header text-text-secondary py-4 px-6 font-medium">
        ITEM / CATEGORY
       </th>
       <th class="font-table-header text-table-header text-text-secondary py-4 px-6 font-medium text-right">
        QTY
       </th>
       <th class="font-table-header text-table-header text-text-secondary py-4 px-6 font-medium text-right">
        EST. COST
       </th>
       <th class="font-table-header text-table-header text-text-secondary py-4 px-6 font-medium">
        REQUESTOR
       </th>
       <th class="font-table-header text-table-header text-text-secondary py-4 px-6 font-medium">
        STATUS
       </th>
       <th class="font-table-header text-table-header text-text-secondary py-4 px-6 font-medium">
        VENDOR
       </th>
       <th class="font-table-header text-table-header text-text-secondary py-4 px-6 font-medium">
        EXP. DELIVERY
       </th>
       <th class="font-table-header text-table-header text-text-secondary py-4 px-6 font-medium text-right">
        ACTIONS
       </th>
      </tr>
     </thead>
     <tbody class="divide-y divide-border-default">
      <!-- Row 1 -->
      <tr class="hover:bg-accent-subtle transition-colors group">
       <td class="py-3 px-6">
        <span class="font-mono-data text-mono-data text-primary font-medium">
         PR-2023-1042
        </span>
       </td>
       <td class="py-3 px-6">
        <div class="font-body text-body font-medium text-text-primary">
         Dell Latitude Laptops
        </div>
        <div class="font-metadata text-metadata text-text-secondary">
         IT Hardware
        </div>
       </td>
       <td class="py-3 px-6 text-right font-mono-data text-mono-data text-text-primary">
        25
       </td>
       <td class="py-3 px-6 text-right font-mono-data text-mono-data text-text-primary">
        ₹ 18.5 L
       </td>
       <td class="py-3 px-6">
        <div class="font-body text-body text-text-primary">
         Sarah Jenkins
        </div>
        <div class="font-metadata text-metadata text-text-secondary">
         IT Dept
        </div>
       </td>
       <td class="py-3 px-6">
        <span class="inline-flex items-center px-2 py-1 rounded-md text-xs font-medium bg-metric-amber text-primary-container">
         Dept Approval
        </span>
       </td>
       <td class="py-3 px-6 font-body text-body text-text-primary">
        -
       </td>
       <td class="py-3 px-6 font-body text-body text-text-secondary">
        TBD
       </td>
       <td class="py-3 px-6 text-right">
        <button class="text-text-secondary hover:text-primary transition-colors p-1">
         <span class="material-symbols-outlined text-[20px]">
          more_vert
         </span>
        </button>
       </td>
      </tr>
      <!-- Row 2 -->
      <tr class="hover:bg-accent-subtle transition-colors group">
       <td class="py-3 px-6">
        <span class="font-mono-data text-mono-data text-primary font-medium">
         PR-2023-1040
        </span>
       </td>
       <td class="py-3 px-6">
        <div class="font-body text-body font-medium text-text-primary">
         Cisco Access Points
        </div>
        <div class="font-metadata text-metadata text-text-secondary">
         Networking
        </div>
       </td>
       <td class="py-3 px-6 text-right font-mono-data text-mono-data text-text-primary">
        12
       </td>
       <td class="py-3 px-6 text-right font-mono-data text-mono-data text-text-primary">
        ₹ 4.2 L
       </td>
       <td class="py-3 px-6">
        <div class="font-body text-body text-text-primary">
         Rahul Sharma
        </div>
        <div class="font-metadata text-metadata text-text-secondary">
         Infrastructure
        </div>
       </td>
       <td class="py-3 px-6">
        <span class="inline-flex items-center px-2 py-1 rounded-md text-xs font-medium bg-tertiary-fixed text-on-tertiary-fixed-variant">
         Finance Approval
        </span>
       </td>
       <td class="py-3 px-6 font-body text-body text-text-primary">
        -
       </td>
       <td class="py-3 px-6 font-body text-body text-text-secondary">
        TBD
       </td>
       <td class="py-3 px-6 text-right">
        <button class="text-text-secondary hover:text-primary transition-colors p-1">
         <span class="material-symbols-outlined text-[20px]">
          more_vert
         </span>
        </button>
       </td>
      </tr>
      <!-- Row 3 -->
      <tr class="hover:bg-accent-subtle transition-colors group border-l-4 border-l-transparent hover:border-l-primary">
       <td class="py-3 px-6">
        <span class="font-mono-data text-mono-data text-primary font-medium">
         PR-2023-1038
        </span>
       </td>
       <td class="py-3 px-6">
        <div class="font-body text-body font-medium text-text-primary">
         Herman Miller Chairs
        </div>
        <div class="font-metadata text-metadata text-text-secondary">
         Office Furniture
        </div>
       </td>
       <td class="py-3 px-6 text-right font-mono-data text-mono-data text-text-primary">
        50
       </td>
       <td class="py-3 px-6 text-right font-mono-data text-mono-data text-text-primary">
        ₹ 12.0 L
       </td>
       <td class="py-3 px-6">
        <div class="font-body text-body text-text-primary">
         Anita Desai
        </div>
        <div class="font-metadata text-metadata text-text-secondary">
         Facilities
        </div>
       </td>
       <td class="py-3 px-6">
        <span class="inline-flex items-center px-2 py-1 rounded-md text-xs font-medium bg-surface-variant text-on-surface-variant border border-outline-variant">
         Draft
        </span>
       </td>
       <td class="py-3 px-6 font-body text-body text-text-primary">
        -
       </td>
       <td class="py-3 px-6 font-body text-body text-text-secondary">
        TBD
       </td>
       <td class="py-3 px-6 text-right">
        <button class="text-text-secondary hover:text-primary transition-colors p-1">
         <span class="material-symbols-outlined text-[20px]">
          more_vert
         </span>
        </button>
       </td>
      </tr>
      <!-- Row 4 -->
      <tr class="hover:bg-accent-subtle transition-colors group">
       <td class="py-3 px-6">
        <span class="font-mono-data text-mono-data text-primary font-medium">
         PR-2023-1035
        </span>
       </td>
       <td class="py-3 px-6">
        <div class="font-body text-body font-medium text-text-primary">
         Server Rack Enclosures
        </div>
        <div class="font-metadata text-metadata text-text-secondary">
         Datacenter
        </div>
       </td>
       <td class="py-3 px-6 text-right font-mono-data text-mono-data text-text-primary">
        4
       </td>
       <td class="py-3 px-6 text-right font-mono-data text-mono-data text-text-primary">
        ₹ 2.8 L
       </td>
       <td class="py-3 px-6">
        <div class="font-body text-body text-text-primary">
         David Chen
        </div>
        <div class="font-metadata text-metadata text-text-secondary">
         IT Dept
        </div>
       </td>
       <td class="py-3 px-6">
        <span class="inline-flex items-center px-2 py-1 rounded-md text-xs font-medium bg-tertiary-container text-on-tertiary-container">
         PO Raised
        </span>
       </td>
       <td class="py-3 px-6 font-body text-body text-text-primary">
        TechCorp India
       </td>
       <td class="py-3 px-6 font-body text-body text-text-primary">
        15 Nov 2023
       </td>
       <td class="py-3 px-6 text-right">
        <button class="text-text-secondary hover:text-primary transition-colors p-1">
         <span class="material-symbols-outlined text-[20px]">
          more_vert
         </span>
        </button>
       </td>
      </tr>
      <!-- Row 5 -->
      <tr class="hover:bg-accent-subtle transition-colors group">
       <td class="py-3 px-6">
        <span class="font-mono-data text-mono-data text-primary font-medium">
         PR-2023-1020
        </span>
       </td>
       <td class="py-3 px-6">
        <div class="font-body text-body font-medium text-text-primary">
         Security Cameras PTZ
        </div>
        <div class="font-metadata text-metadata text-text-secondary">
         Security
        </div>
       </td>
       <td class="py-3 px-6 text-right font-mono-data text-mono-data text-text-primary">
        8
       </td>
       <td class="py-3 px-6 text-right font-mono-data text-mono-data text-text-primary">
        ₹ 6.5 L
       </td>
       <td class="py-3 px-6">
        <div class="font-body text-body text-text-primary">
         Rajesh Kumar
        </div>
        <div class="font-metadata text-metadata text-text-secondary">
         Security Ops
        </div>
       </td>
       <td class="py-3 px-6">
        <span class="inline-flex items-center px-2 py-1 rounded-md text-xs font-medium bg-metric-sage text-status-in-stock">
         Received
        </span>
       </td>
       <td class="py-3 px-6 font-body text-body text-text-primary">
        SafeGuard Solutions
       </td>
       <td class="py-3 px-6 font-body text-body text-text-secondary">
        Delivered
       </td>
       <td class="py-3 px-6 text-right">
        <button class="text-text-secondary hover:text-primary transition-colors p-1">
         <span class="material-symbols-outlined text-[20px]">
          more_vert
         </span>
        </button>
       </td>
      </tr>
     </tbody>
    </table>
   </div>
   <!-- Table Pagination/Footer -->
   <div class="bg-white border-t border-border-default px-6 py-3 flex items-center justify-between">
    <span class="font-metadata text-metadata text-text-secondary">
     Showing 1-5 of 124 requests
    </span>
    <div class="flex items-center gap-1">
     <button class="p-1 rounded text-stone-400 hover:text-stone-900 hover:bg-stone-100 disabled:opacity-50">
      <span class="material-symbols-outlined text-[20px]">
       chevron_left
      </span>
     </button>
     <button class="w-8 h-8 rounded-md bg-stone-100 text-stone-900 font-medium text-sm">
      1
     </button>
     <button class="w-8 h-8 rounded-md text-stone-600 hover:bg-stone-50 font-medium text-sm">
      2
     </button>
     <button class="w-8 h-8 rounded-md text-stone-600 hover:bg-stone-50 font-medium text-sm">
      3
     </button>
     <button class="p-1 rounded text-stone-600 hover:text-stone-900 hover:bg-stone-100">
      <span class="material-symbols-outlined text-[20px]">
       chevron_right
      </span>
     </button>
    </div>
   </div>
  </div>
 </div>
</main>

</template>

<script setup lang="ts">
// Autogenerated from procurement_pipeline_2
</script>
