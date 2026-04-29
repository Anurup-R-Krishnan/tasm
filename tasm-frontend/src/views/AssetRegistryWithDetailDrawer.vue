<template>
<main class="flex-1 flex flex-col min-w-0">
 <!-- TopNavBar Component -->
 <header class="h-[60px] w-full border-b sticky top-0 z-30 border-stone-200 dark:border-stone-800 shadow-sm bg-stone-50 dark:bg-stone-900 font-sans antialiased text-stone-900 dark:text-stone-100">
  <div class="flex items-center justify-between px-6 w-full max-w-[1400px] mx-auto h-full">
   <!-- Brand (Optional in Hybrid, but part of TopNavBar JSON) -->
   <div class="text-lg font-bold tracking-tight text-stone-900 dark:text-stone-50 flex items-center gap-4">
    <div class="relative w-64">
     <span class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-text-secondary text-[20px]" data-icon="search">
      search
     </span>
     <input class="w-full pl-10 pr-4 py-1.5 bg-surface border border-border-default rounded-md text-sm focus:outline-none focus:border-primary-container" placeholder="Search..." type="text"/>
    </div>
   </div>
   <!-- Trailing Actions -->
   <div class="flex items-center gap-2">
    <button class="p-2 rounded-full text-stone-500 dark:text-stone-400 hover:bg-stone-100 dark:hover:bg-stone-800 transition-colors cursor-pointer active:scale-95 duration-150">
     <span class="material-symbols-outlined" data-icon="notifications">
      notifications
     </span>
    </button>
    <button class="p-2 rounded-full text-stone-500 dark:text-stone-400 hover:bg-stone-100 dark:hover:bg-stone-800 transition-colors cursor-pointer active:scale-95 duration-150">
     <span class="material-symbols-outlined" data-icon="settings">
      settings
     </span>
    </button>
    <button class="p-2 rounded-full text-stone-500 dark:text-stone-400 hover:bg-stone-100 dark:hover:bg-stone-800 transition-colors cursor-pointer active:scale-95 duration-150">
     <span class="material-symbols-outlined" data-icon="help">
      help
     </span>
    </button>
    <div class="ml-2 w-8 h-8 rounded-full bg-border-default overflow-hidden border border-border-default">
     <img alt="User profile" class="w-full h-full object-cover" data-alt="Close up portrait of a professional smiling slightly in natural lighting" src="https://lh3.googleusercontent.com/aida-public/AB6AXuAgoTg1INEPcxLsEzOeYnNxgTpLngLZZlBHcjv5aUMYnl6DakeZPzs2JyBmrTV_ZqDDAZahGRt8MV4vZ3Z2W_AiNmkUoseDdOU7c2jMe_9OahcIx-DFeX9SwWYDu9IVS0RL5dE4Ayl--_AmjpV0UW7hlbLB5wrQk-nfsJo2It5n2wpB4hrj2iFeI5T0f7oAbZVq8nPP-DW_sye8TTOfNNf8qZkaOt7MmTOK9eqieEUVs2CrfaKQqBQgh007sS6AfTPasBSBKTKPwWXi"/>
    </div>
   </div>
  </div>
 </header>
 <!-- Canvas Content -->
 <div class="p-page-margin max-w-[1400px] mx-auto w-full">
  <!-- Page Header -->
  <div class="flex justify-between items-center mb-stack">
   <h1 class="font-h1 text-h1 text-text-primary">
    Asset Registry
   </h1>
   <div class="flex items-center gap-inline">
    <button class="px-4 py-2 bg-surface border border-border-default rounded-lg text-text-primary font-h3 text-h3 hover:bg-surface-subtle transition-colors flex items-center gap-2 shadow-sm">
     <span class="material-symbols-outlined text-[18px]" data-icon="download">
      download
     </span>
     Export
    </button>
    <button class="px-4 py-2 bg-text-primary text-on-primary rounded-lg font-h3 text-h3 hover:bg-black transition-colors flex items-center gap-2 shadow-sm">
     <span class="material-symbols-outlined text-[18px]" data-icon="add">
      add
     </span>
     Add New Asset
    </button>
   </div>
  </div>
  <!-- Status Summary Pills -->
  <div class="flex gap-inline overflow-x-auto pb-2 mb-stack hide-scrollbar">
   <button class="px-4 py-1.5 rounded-full bg-primary-container text-on-primary font-body text-body whitespace-nowrap flex items-center gap-2 border border-primary-container">
    All
    <span class="bg-white/20 px-2 py-0.5 rounded-full text-[12px]">
     1,248
    </span>
   </button>
   <button class="px-4 py-1.5 rounded-full bg-surface border border-border-default text-text-secondary hover:border-text-secondary transition-colors font-body text-body whitespace-nowrap flex items-center gap-2">
    In Stock
    <span class="bg-surface-subtle px-2 py-0.5 rounded-full text-[12px] text-text-primary">
     432
    </span>
   </button>
   <button class="px-4 py-1.5 rounded-full bg-surface border border-border-default text-text-secondary hover:border-text-secondary transition-colors font-body text-body whitespace-nowrap flex items-center gap-2">
    Checked Out
    <span class="bg-surface-subtle px-2 py-0.5 rounded-full text-[12px] text-text-primary">
     782
    </span>
   </button>
   <button class="px-4 py-1.5 rounded-full bg-surface border border-border-default text-text-secondary hover:border-text-secondary transition-colors font-body text-body whitespace-nowrap flex items-center gap-2">
    Reserved
    <span class="bg-surface-subtle px-2 py-0.5 rounded-full text-[12px] text-text-primary">
     12
    </span>
   </button>
   <button class="px-4 py-1.5 rounded-full bg-surface border border-border-default text-text-secondary hover:border-text-secondary transition-colors font-body text-body whitespace-nowrap flex items-center gap-2">
    Under Repair
    <span class="bg-surface-subtle px-2 py-0.5 rounded-full text-[12px] text-text-primary">
     18
    </span>
   </button>
   <button class="px-4 py-1.5 rounded-full bg-surface border border-border-default text-text-secondary hover:border-text-secondary transition-colors font-body text-body whitespace-nowrap flex items-center gap-2">
    Lost
    <span class="bg-surface-subtle px-2 py-0.5 rounded-full text-[12px] text-text-primary">
     4
    </span>
   </button>
  </div>
  <!-- Filter Bar -->
  <div class="flex justify-between items-center bg-surface p-inline rounded-xl border border-border-default mb-stack shadow-sm">
   <div class="flex items-center gap-inline flex-1">
    <div class="relative w-80">
     <span class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-text-secondary text-[20px]" data-icon="search">
      search
     </span>
     <input class="w-full pl-10 pr-4 py-2 bg-canvas border border-border-default rounded-lg text-sm focus:outline-none focus:border-primary-container focus:ring-1 focus:ring-primary-container transition-all" placeholder="Search by asset name, tag, or serial..." type="text"/>
    </div>
    <div class="h-6 w-px bg-border-default mx-2">
    </div>
    <button class="px-3 py-1.5 bg-surface-subtle border border-border-default rounded-md text-text-primary font-body text-body hover:bg-border-default transition-colors flex items-center gap-2">
     <span class="material-symbols-outlined text-[18px]" data-icon="filter_list">
      filter_list
     </span>
     Filters
    </button>
    <!-- Active Chips -->
    <div class="flex items-center gap-2">
     <span class="px-3 py-1 bg-surface-variant text-on-surface-variant rounded-md text-[13px] flex items-center gap-1 border border-outline-variant">
      Category: IT
      <span class="material-symbols-outlined text-[14px] cursor-pointer hover:text-text-primary" data-icon="close">
       close
      </span>
     </span>
     <span class="px-3 py-1 bg-surface-variant text-on-surface-variant rounded-md text-[13px] flex items-center gap-1 border border-outline-variant">
      Dept: HR
      <span class="material-symbols-outlined text-[14px] cursor-pointer hover:text-text-primary" data-icon="close">
       close
      </span>
     </span>
    </div>
   </div>
   <!-- Bulk Actions (Hidden by default, shown for context) -->
   <div class="flex items-center gap-2 hidden">
    <span class="text-sm text-text-secondary mr-2">
     2 selected
    </span>
    <button class="px-3 py-1.5 bg-surface border border-border-default rounded-md text-text-primary hover:bg-surface-subtle transition-colors text-sm">
     Update Status
    </button>
    <button class="px-3 py-1.5 bg-surface border border-border-default rounded-md text-error hover:bg-error-container hover:border-error-container transition-colors text-sm">
     Delete
    </button>
   </div>
  </div>
  <!-- Main Table Container -->
  <div class="bg-surface rounded-xl border border-border-default overflow-hidden shadow-[0_4px_12px_rgba(0,0,0,0.02)]">
   <table class="w-full text-left border-collapse">
    <thead>
     <tr class="bg-surface-subtle border-b border-border-default text-table-header font-table-header text-text-secondary uppercase">
      <th class="py-3 px-4 w-12 text-center">
       <input class="rounded border-outline-variant text-primary-container focus:ring-primary-container" type="checkbox"/>
      </th>
      <th class="py-3 px-4">
       Tag ID
      </th>
      <th class="py-3 px-4">
       Asset Name
      </th>
      <th class="py-3 px-4">
       Category
      </th>
      <th class="py-3 px-4">
       Status
      </th>
      <th class="py-3 px-4">
       Custodian
      </th>
      <th class="py-3 px-4">
       Location
      </th>
      <th class="py-3 px-4 text-right">
       Actions
      </th>
     </tr>
    </thead>
    <tbody class="text-body font-body text-text-primary divide-y divide-border-default">
     <!-- Row 1 (Simulating selection/active detail view) -->
     <tr class="bg-accent-subtle/30 hover:bg-accent-subtle transition-colors group cursor-pointer border-l-4 border-l-primary-container">
      <td class="py-3 px-4 text-center">
       <input class="rounded border-outline-variant text-primary-container focus:ring-primary-container" type="checkbox"/>
      </td>
      <td class="py-3 px-4 font-mono-data text-mono-data text-primary-container underline decoration-primary-container/30 underline-offset-2">
       TAG-1082
      </td>
      <td class="py-3 px-4 font-medium">
       MacBook Pro 16"
      </td>
      <td class="py-3 px-4 text-text-secondary">
       IT Equipment
      </td>
      <td class="py-3 px-4">
       <span class="inline-flex items-center px-2 py-0.5 rounded-full text-[12px] font-medium bg-tertiary-fixed text-status-checked-out border border-tertiary-fixed-dim/50">
        Checked Out
       </span>
      </td>
      <td class="py-3 px-4">
       Sarah Jenkins
      </td>
      <td class="py-3 px-4 text-text-secondary">
       Bldg 3, Fl 4
      </td>
      <td class="py-3 px-4 text-right space-x-1">
       <button class="p-1.5 text-text-secondary hover:text-primary-container hover:bg-surface rounded transition-colors">
        <span class="material-symbols-outlined text-[18px]" data-icon="visibility">
         visibility
        </span>
       </button>
       <button class="p-1.5 text-text-secondary hover:text-primary-container hover:bg-surface rounded transition-colors">
        <span class="material-symbols-outlined text-[18px]" data-icon="edit">
         edit
        </span>
       </button>
       <button class="p-1.5 text-text-secondary hover:text-text-primary hover:bg-surface rounded transition-colors">
        <span class="material-symbols-outlined text-[18px]" data-icon="more_vert">
         more_vert
        </span>
       </button>
      </td>
     </tr>
     <!-- Row 2 -->
     <tr class="bg-surface hover:bg-accent-subtle transition-colors group cursor-pointer border-l-4 border-l-transparent">
      <td class="py-3 px-4 text-center">
       <input class="rounded border-outline-variant text-primary-container focus:ring-primary-container" type="checkbox"/>
      </td>
      <td class="py-3 px-4 font-mono-data text-mono-data text-primary-container hover:underline underline-offset-2">
       TAG-2094
      </td>
      <td class="py-3 px-4 font-medium">
       Dell Ultrasharp 27"
      </td>
      <td class="py-3 px-4 text-text-secondary">
       IT Peripherals
      </td>
      <td class="py-3 px-4">
       <span class="inline-flex items-center px-2 py-0.5 rounded-full text-[12px] font-medium bg-metric-sage text-status-in-stock border border-metric-sage/50">
        In Stock
       </span>
      </td>
      <td class="py-3 px-4 text-text-secondary">
       --
      </td>
      <td class="py-3 px-4 text-text-secondary">
       IT Storage A
      </td>
      <td class="py-3 px-4 text-right space-x-1 opacity-0 group-hover:opacity-100 transition-opacity">
       <button class="p-1.5 text-text-secondary hover:text-primary-container hover:bg-surface rounded transition-colors">
        <span class="material-symbols-outlined text-[18px]" data-icon="visibility">
         visibility
        </span>
       </button>
       <button class="p-1.5 text-text-secondary hover:text-primary-container hover:bg-surface rounded transition-colors">
        <span class="material-symbols-outlined text-[18px]" data-icon="edit">
         edit
        </span>
       </button>
       <button class="p-1.5 text-text-secondary hover:text-text-primary hover:bg-surface rounded transition-colors">
        <span class="material-symbols-outlined text-[18px]" data-icon="more_vert">
         more_vert
        </span>
       </button>
      </td>
     </tr>
     <!-- Row 3 -->
     <tr class="bg-surface-subtle hover:bg-accent-subtle transition-colors group cursor-pointer border-l-4 border-l-transparent">
      <td class="py-3 px-4 text-center">
       <input class="rounded border-outline-variant text-primary-container focus:ring-primary-container" type="checkbox"/>
      </td>
      <td class="py-3 px-4 font-mono-data text-mono-data text-primary-container hover:underline underline-offset-2">
       TAG-0921
      </td>
      <td class="py-3 px-4 font-medium">
       Herman Miller Aeron
      </td>
      <td class="py-3 px-4 text-text-secondary">
       Furniture
      </td>
      <td class="py-3 px-4">
       <span class="inline-flex items-center px-2 py-0.5 rounded-full text-[12px] font-medium bg-tertiary-fixed text-status-checked-out border border-tertiary-fixed-dim/50">
        Checked Out
       </span>
      </td>
      <td class="py-3 px-4">
       David Chen
      </td>
      <td class="py-3 px-4 text-text-secondary">
       Bldg 1, Fl 2
      </td>
      <td class="py-3 px-4 text-right space-x-1 opacity-0 group-hover:opacity-100 transition-opacity">
       <button class="p-1.5 text-text-secondary hover:text-primary-container hover:bg-surface rounded transition-colors">
        <span class="material-symbols-outlined text-[18px]" data-icon="visibility">
         visibility
        </span>
       </button>
       <button class="p-1.5 text-text-secondary hover:text-primary-container hover:bg-surface rounded transition-colors">
        <span class="material-symbols-outlined text-[18px]" data-icon="edit">
         edit
        </span>
       </button>
       <button class="p-1.5 text-text-secondary hover:text-text-primary hover:bg-surface rounded transition-colors">
        <span class="material-symbols-outlined text-[18px]" data-icon="more_vert">
         more_vert
        </span>
       </button>
      </td>
     </tr>
     <!-- Row 4 -->
     <tr class="bg-surface hover:bg-accent-subtle transition-colors group cursor-pointer border-l-4 border-l-transparent">
      <td class="py-3 px-4 text-center">
       <input class="rounded border-outline-variant text-primary-container focus:ring-primary-container" type="checkbox"/>
      </td>
      <td class="py-3 px-4 font-mono-data text-mono-data text-primary-container hover:underline underline-offset-2">
       TAG-3312
      </td>
      <td class="py-3 px-4 font-medium">
       Conference Table Oak
      </td>
      <td class="py-3 px-4 text-text-secondary">
       Furniture
      </td>
      <td class="py-3 px-4">
       <span class="inline-flex items-center px-2 py-0.5 rounded-full text-[12px] font-medium bg-error-container text-status-critical border border-error-container/50">
        Under Repair
       </span>
      </td>
      <td class="py-3 px-4 text-text-secondary">
       Facilities
      </td>
      <td class="py-3 px-4 text-text-secondary">
       Maint Workshop
      </td>
      <td class="py-3 px-4 text-right space-x-1 opacity-0 group-hover:opacity-100 transition-opacity">
       <button class="p-1.5 text-text-secondary hover:text-primary-container hover:bg-surface rounded transition-colors">
        <span class="material-symbols-outlined text-[18px]" data-icon="visibility">
         visibility
        </span>
       </button>
       <button class="p-1.5 text-text-secondary hover:text-primary-container hover:bg-surface rounded transition-colors">
        <span class="material-symbols-outlined text-[18px]" data-icon="edit">
         edit
        </span>
       </button>
       <button class="p-1.5 text-text-secondary hover:text-text-primary hover:bg-surface rounded transition-colors">
        <span class="material-symbols-outlined text-[18px]" data-icon="more_vert">
         more_vert
        </span>
       </button>
      </td>
     </tr>
    </tbody>
   </table>
   <!-- Pagination Footer -->
   <div class="bg-surface px-4 py-3 border-t border-border-default flex items-center justify-between sm:px-6">
    <div class="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
     <div>
      <p class="text-sm text-text-secondary">
       Showing
       <span class="font-medium text-text-primary">
        1
       </span>
       to
       <span class="font-medium text-text-primary">
        10
       </span>
       of
       <span class="font-medium text-text-primary">
        1,248
       </span>
       results
      </p>
     </div>
     <div>
      <nav aria-label="Pagination" class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px">
       <button class="relative inline-flex items-center px-2 py-2 rounded-l-md border border-border-default bg-surface text-sm font-medium text-text-secondary hover:bg-surface-subtle">
        <span class="sr-only">
         Previous
        </span>
        <span class="material-symbols-outlined text-[18px]" data-icon="chevron_left">
         chevron_left
        </span>
       </button>
       <button class="relative inline-flex items-center px-4 py-2 border border-border-default bg-surface-subtle text-sm font-medium text-primary-container">
        1
       </button>
       <button class="relative inline-flex items-center px-4 py-2 border border-border-default bg-surface text-sm font-medium text-text-secondary hover:bg-surface-subtle">
        2
       </button>
       <button class="relative inline-flex items-center px-4 py-2 border border-border-default bg-surface text-sm font-medium text-text-secondary hover:bg-surface-subtle">
        3
       </button>
       <span class="relative inline-flex items-center px-4 py-2 border border-border-default bg-surface text-sm font-medium text-text-secondary">
        ...
       </span>
       <button class="relative inline-flex items-center px-2 py-2 rounded-r-md border border-border-default bg-surface text-sm font-medium text-text-secondary hover:bg-surface-subtle">
        <span class="sr-only">
         Next
        </span>
        <span class="material-symbols-outlined text-[18px]" data-icon="chevron_right">
         chevron_right
        </span>
       </button>
      </nav>
     </div>
    </div>
   </div>
  </div>
 </div>
</main>

</template>

<script setup lang="ts">
// Autogenerated from asset_registry_with_detail_drawer
</script>
