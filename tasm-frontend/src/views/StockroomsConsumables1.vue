<template>
<main class="p-page-margin w-full max-w-[1400px] flex gap-section-gap">
 <!-- Left Panel: Stockroom List -->
 <div class="w-[280px] shrink-0 flex flex-col gap-stack">
  <h1 class="font-h1 text-h1 text-text-primary">
   Locations
  </h1>
  <div class="bg-surface border border-border-default rounded-xl shadow-sm overflow-hidden flex flex-col">
   <div class="p-4 border-b border-border-default bg-surface-subtle">
    <div class="relative text-text-secondary">
     <span class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-[18px]">
      search
     </span>
     <input class="bg-surface border border-border-default rounded-lg pl-9 pr-3 py-1.5 text-body font-body w-full focus:outline-none focus:ring-1 focus:ring-amber-500 text-sm" placeholder="Filter stockrooms..." type="text"/>
    </div>
   </div>
   <div class="flex-1 overflow-y-auto">
    <!-- Selected List Item -->
    <div class="p-4 border-l-4 border-amber-700 bg-amber-50 cursor-pointer border-b border-border-default last:border-0 hover:bg-amber-100 transition-colors">
     <div class="flex justify-between items-start mb-1">
      <span class="font-h3 text-h3 text-amber-900">
       Main Store
      </span>
      <span class="bg-white border border-amber-200 text-amber-800 text-[10px] font-bold px-1.5 py-0.5 rounded">
       HQ
      </span>
     </div>
     <div class="text-metadata font-metadata text-text-secondary flex items-center gap-1 mt-1">
      <span class="material-symbols-outlined text-[14px]">
       person
      </span>
      Rajesh K.
     </div>
     <div class="text-metadata font-metadata text-text-secondary flex items-center gap-1 mt-0.5">
      <span class="material-symbols-outlined text-[14px]">
       inventory
      </span>
      1,245 Assets
     </div>
    </div>
    <!-- Unselected List Item -->
    <div class="p-4 border-l-4 border-transparent bg-surface cursor-pointer border-b border-border-default last:border-0 hover:bg-surface-subtle transition-colors">
     <div class="flex justify-between items-start mb-1">
      <span class="font-h3 text-h3 text-text-primary">
       IT Lab
      </span>
     </div>
     <div class="text-metadata font-metadata text-text-secondary flex items-center gap-1 mt-1">
      <span class="material-symbols-outlined text-[14px]">
       person
      </span>
      Anjali M.
     </div>
     <div class="text-metadata font-metadata text-text-secondary flex items-center gap-1 mt-0.5">
      <span class="material-symbols-outlined text-[14px]">
       inventory
      </span>
      842 Assets
     </div>
    </div>
    <!-- Unselected List Item -->
    <div class="p-4 border-l-4 border-transparent bg-surface cursor-pointer border-b border-border-default last:border-0 hover:bg-surface-subtle transition-colors">
     <div class="flex justify-between items-start mb-1">
      <span class="font-h3 text-h3 text-text-primary">
       Admin Block B
      </span>
     </div>
     <div class="text-metadata font-metadata text-text-secondary flex items-center gap-1 mt-1">
      <span class="material-symbols-outlined text-[14px]">
       person
      </span>
      Suresh P.
     </div>
     <div class="text-metadata font-metadata text-text-secondary flex items-center gap-1 mt-0.5">
      <span class="material-symbols-outlined text-[14px]">
       inventory
      </span>
      310 Assets
     </div>
    </div>
   </div>
  </div>
 </div>
 <!-- Right Panel: Selected Stockroom Details -->
 <div class="flex-1 flex flex-col gap-section-gap min-w-0">
  <!-- Header Card (Amber) -->
  <div class="bg-metric-amber rounded-xl p-card-padding shadow-sm border border-amber-200 relative overflow-hidden flex flex-col justify-between">
   <!-- Decorative background shape -->
   <div class="absolute right-0 top-0 w-64 h-64 bg-amber-200/50 rounded-full blur-3xl -translate-y-1/2 translate-x-1/4 pointer-events-none">
   </div>
   <div class="relative z-10 flex justify-between items-start">
    <div>
     <div class="flex items-center gap-2 mb-2">
      <h2 class="font-h1 text-h1 text-amber-900">
       Main Store
      </h2>
      <span class="bg-white/60 border border-amber-300 text-amber-800 text-xs font-semibold px-2 py-0.5 rounded-full">
       Primary Facility
      </span>
     </div>
     <div class="flex items-center gap-6 text-amber-800 font-body text-body">
      <span class="flex items-center gap-1.5">
       <span class="material-symbols-outlined text-[18px]">
        person
       </span>
       Manager: Rajesh Kumar
      </span>
      <span class="flex items-center gap-1.5">
       <span class="material-symbols-outlined text-[18px]">
        location_on
       </span>
       Technopark Phase 1, Ground Floor
      </span>
     </div>
    </div>
    <button class="bg-white text-text-primary border border-amber-300 hover:bg-amber-50 px-4 py-2 rounded-lg font-body text-body font-medium shadow-sm transition-colors flex items-center gap-2">
     <span class="material-symbols-outlined text-[18px]">
      edit
     </span>
     Edit Details
    </button>
   </div>
   <div class="relative z-10 grid grid-cols-2 gap-6 mt-8">
    <div class="bg-white/60 p-4 rounded-lg border border-amber-200/50 backdrop-blur-sm">
     <p class="font-metadata text-metadata text-amber-800 mb-1 uppercase tracking-wider">
      Total Assets
     </p>
     <p class="font-kpi-number text-kpi-number text-amber-950">
      1,245
     </p>
    </div>
    <div class="bg-white/60 p-4 rounded-lg border border-amber-200/50 backdrop-blur-sm">
     <p class="font-metadata text-metadata text-amber-800 mb-1 uppercase tracking-wider">
      Total Value
     </p>
     <p class="font-kpi-number text-kpi-number text-amber-950">
      ₹45,20,500
     </p>
    </div>
   </div>
  </div>
  <!-- Content Tabs -->
  <div class="flex gap-6 border-b border-border-default mt-2">
   <button class="font-h3 text-h3 text-amber-700 border-b-2 border-amber-700 pb-2 px-1">
    Fixed Assets
   </button>
   <button class="font-h3 text-h3 text-text-secondary hover:text-text-primary border-b-2 border-transparent hover:border-border-default pb-2 px-1 transition-colors">
    Consumables
   </button>
   <button class="font-h3 text-h3 text-text-secondary hover:text-text-primary border-b-2 border-transparent hover:border-border-default pb-2 px-1 transition-colors">
    Audit Logs
   </button>
  </div>
  <!-- Fixed Assets Table Area -->
  <div class="bg-surface border border-border-default rounded-xl shadow-sm overflow-hidden flex flex-col hover:-translate-y-[2px] transition-transform duration-200">
   <div class="p-4 border-b border-border-default flex justify-between items-center bg-surface-subtle">
    <h3 class="font-h2 text-h2 text-text-primary">
     Inventory List
    </h3>
    <div class="flex gap-2">
     <button class="bg-white border border-border-default hover:bg-surface-subtle text-text-primary p-1.5 rounded-md transition-colors">
      <span class="material-symbols-outlined text-[18px]">
       filter_list
      </span>
     </button>
     <button class="bg-white border border-border-default hover:bg-surface-subtle text-text-primary p-1.5 rounded-md transition-colors">
      <span class="material-symbols-outlined text-[18px]">
       download
      </span>
     </button>
    </div>
   </div>
   <div class="overflow-x-auto">
    <table class="w-full text-left border-collapse">
     <thead>
      <tr class="bg-surface border-b border-border-default text-text-secondary font-table-header text-table-header uppercase">
       <th class="p-4 font-medium">
        Asset Tag ID
       </th>
       <th class="p-4 font-medium">
        Asset Name
       </th>
       <th class="p-4 font-medium">
        Category
       </th>
       <th class="p-4 font-medium">
        Condition
       </th>
       <th class="p-4 font-medium">
        Date Stocked
       </th>
       <th class="p-4 font-medium">
        Last Movement
       </th>
      </tr>
     </thead>
     <tbody class="font-body text-body text-text-primary divide-y divide-border-default">
      <tr class="hover:bg-accent-subtle transition-colors group">
       <td class="p-4 border-l-4 border-transparent group-hover:border-amber-500 font-mono-data text-mono-data text-text-secondary">
        TPK-LT-8821
       </td>
       <td class="p-4 font-medium">
        Dell Latitude 5420
       </td>
       <td class="p-4">
        Laptops
       </td>
       <td class="p-4">
        <span class="inline-flex items-center gap-1.5 bg-metric-sage/50 text-status-in-stock px-2.5 py-1 rounded-full text-xs font-semibold border border-metric-sage">
         <span class="w-1.5 h-1.5 rounded-full bg-status-in-stock">
         </span>
         Excellent
        </span>
       </td>
       <td class="p-4 text-text-secondary">
        12 Oct 2023
       </td>
       <td class="p-4 text-text-secondary">
        05 Jan 2024
       </td>
      </tr>
      <tr class="bg-surface-subtle hover:bg-accent-subtle transition-colors group">
       <td class="p-4 border-l-4 border-transparent group-hover:border-amber-500 font-mono-data text-mono-data text-text-secondary">
        TPK-MN-4410
       </td>
       <td class="p-4 font-medium">
        Dell 24" Monitor P2422H
       </td>
       <td class="p-4">
        Monitors
       </td>
       <td class="p-4">
        <span class="inline-flex items-center gap-1.5 bg-metric-sage/50 text-status-in-stock px-2.5 py-1 rounded-full text-xs font-semibold border border-metric-sage">
         <span class="w-1.5 h-1.5 rounded-full bg-status-in-stock">
         </span>
         Good
        </span>
       </td>
       <td class="p-4 text-text-secondary">
        15 Nov 2023
       </td>
       <td class="p-4 text-text-secondary">
        -
       </td>
      </tr>
      <tr class="hover:bg-accent-subtle transition-colors group">
       <td class="p-4 border-l-4 border-transparent group-hover:border-amber-500 font-mono-data text-mono-data text-text-secondary">
        TPK-DK-1102
       </td>
       <td class="p-4 font-medium">
        WD19S Docking Station
       </td>
       <td class="p-4">
        Peripherals
       </td>
       <td class="p-4">
        <span class="inline-flex items-center gap-1.5 bg-error-container/50 text-status-critical px-2.5 py-1 rounded-full text-xs font-semibold border border-error-container">
         <span class="w-1.5 h-1.5 rounded-full bg-status-critical">
         </span>
         Needs Repair
        </span>
       </td>
       <td class="p-4 text-text-secondary">
        02 Sep 2022
       </td>
       <td class="p-4 text-text-secondary">
        18 Feb 2024
       </td>
      </tr>
      <tr class="bg-surface-subtle hover:bg-accent-subtle transition-colors group">
       <td class="p-4 border-l-4 border-transparent group-hover:border-amber-500 font-mono-data text-mono-data text-text-secondary">
        TPK-PR-0091
       </td>
       <td class="p-4 font-medium">
        HP LaserJet Pro M404n
       </td>
       <td class="p-4">
        Printers
       </td>
       <td class="p-4">
        <span class="inline-flex items-center gap-1.5 bg-metric-sage/50 text-status-in-stock px-2.5 py-1 rounded-full text-xs font-semibold border border-metric-sage">
         <span class="w-1.5 h-1.5 rounded-full bg-status-in-stock">
         </span>
         Good
        </span>
       </td>
       <td class="p-4 text-text-secondary">
        22 Jan 2023
       </td>
       <td class="p-4 text-text-secondary">
        10 Mar 2024
       </td>
      </tr>
     </tbody>
    </table>
   </div>
   <div class="p-4 border-t border-border-default flex items-center justify-between bg-surface text-sm text-text-secondary">
    <span>
     Showing 1 to 4 of 1,245 entries
    </span>
    <div class="flex gap-1">
     <button class="p-1 border border-border-default rounded hover:bg-surface-subtle disabled:opacity-50" disabled>
      <span class="material-symbols-outlined text-[18px]">
       chevron_left
      </span>
     </button>
     <button class="px-3 py-1 bg-amber-50 text-amber-700 border border-amber-200 rounded font-medium">
      1
     </button>
     <button class="px-3 py-1 border border-border-default rounded hover:bg-surface-subtle">
      2
     </button>
     <button class="px-3 py-1 border border-border-default rounded hover:bg-surface-subtle">
      3
     </button>
     <span class="px-2 py-1">
      ...
     </span>
     <button class="p-1 border border-border-default rounded hover:bg-surface-subtle">
      <span class="material-symbols-outlined text-[18px]">
       chevron_right
      </span>
     </button>
    </div>
   </div>
  </div>
  <!-- Consumables Quick View (Bento Grid Style) -->
  <div>
   <h3 class="font-h2 text-h2 text-text-primary mb-4 flex items-center gap-2">
    <span class="material-symbols-outlined text-amber-600 text-[20px]">
     water_drop
    </span>
    Consumables Alert
   </h3>
   <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
    <!-- Alert Card 1 -->
    <div class="bg-surface border border-error-container/50 rounded-xl p-4 shadow-sm relative overflow-hidden group hover:-translate-y-[2px] transition-all">
     <div class="absolute top-0 left-0 w-1 h-full bg-status-critical">
     </div>
     <div class="flex justify-between items-start mb-2">
      <h4 class="font-h3 text-h3 text-text-primary">
       HP Toner 59A
      </h4>
      <span class="bg-error-container text-on-error-container text-[10px] font-bold px-2 py-0.5 rounded-full uppercase tracking-wide">
       Out of Stock
      </span>
     </div>
     <p class="text-metadata font-metadata text-text-secondary mb-3">
      Black LaserJet Toner Cartridge
     </p>
     <div class="flex justify-between items-end">
      <span class="font-kpi-number text-2xl text-status-critical">
       0
       <span class="text-sm font-normal text-text-secondary">
        units
       </span>
      </span>
      <button class="text-amber-700 text-sm font-medium hover:underline flex items-center gap-1">
       Reorder
       <span class="material-symbols-outlined text-[14px]">
        arrow_forward
       </span>
      </button>
     </div>
    </div>
    <!-- Alert Card 2 -->
    <div class="bg-surface border border-amber-200/50 rounded-xl p-4 shadow-sm relative overflow-hidden group hover:-translate-y-[2px] transition-all">
     <div class="absolute top-0 left-0 w-1 h-full bg-amber-500">
     </div>
     <div class="flex justify-between items-start mb-2">
      <h4 class="font-h3 text-h3 text-text-primary">
       CAT6 Ethernet Cable
      </h4>
      <span class="bg-amber-100 text-amber-800 text-[10px] font-bold px-2 py-0.5 rounded-full uppercase tracking-wide">
       Low Stock
      </span>
     </div>
     <p class="text-metadata font-metadata text-text-secondary mb-3">
      10m Blue Patch Cable
     </p>
     <div class="flex justify-between items-end">
      <span class="font-kpi-number text-2xl text-amber-600">
       12
       <span class="text-sm font-normal text-text-secondary">
        units
       </span>
      </span>
      <button class="text-amber-700 text-sm font-medium hover:underline flex items-center gap-1">
       Reorder
       <span class="material-symbols-outlined text-[14px]">
        arrow_forward
       </span>
      </button>
     </div>
    </div>
    <!-- Normal Status Card -->
    <div class="bg-surface border border-border-default rounded-xl p-4 shadow-sm relative overflow-hidden group hover:-translate-y-[2px] transition-all">
     <div class="absolute top-0 left-0 w-1 h-full bg-status-in-stock">
     </div>
     <div class="flex justify-between items-start mb-2">
      <h4 class="font-h3 text-h3 text-text-primary">
       AA Batteries
      </h4>
      <span class="bg-metric-sage text-status-in-stock text-[10px] font-bold px-2 py-0.5 rounded-full uppercase tracking-wide">
       In Stock
      </span>
     </div>
     <p class="text-metadata font-metadata text-text-secondary mb-3">
      Alkaline, Pack of 4
     </p>
     <div class="flex justify-between items-end">
      <span class="font-kpi-number text-2xl text-text-primary">
       145
       <span class="text-sm font-normal text-text-secondary">
        packs
       </span>
      </span>
      <button class="text-text-secondary text-sm font-medium hover:text-text-primary flex items-center gap-1">
       View
       <span class="material-symbols-outlined text-[14px]">
        arrow_forward
       </span>
      </button>
     </div>
    </div>
   </div>
  </div>
 </div>
</main>

</template>

<script setup lang="ts">
// Autogenerated from stockrooms_consumables_1
</script>
