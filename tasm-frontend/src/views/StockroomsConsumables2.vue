<template>
<main class="flex-1 mt-[60px] p-page-margin max-w-[1400px] w-full mx-auto">
 <!-- Page Header -->
 <div class="flex justify-between items-end mb-section-gap">
  <div>
   <h1 class="font-h1 text-h1 text-text-primary">
    Stockrooms &amp; Consumables
   </h1>
   <p class="font-body text-body text-text-secondary mt-1">
    Manage physical inventory locations and consumable supplies.
   </p>
  </div>
  <button class="bg-[#1C1917] text-white px-4 py-2 rounded-lg font-medium hover:bg-black transition-colors flex items-center space-x-2 text-sm shadow-sm">
   <span class="material-symbols-outlined text-[18px]">
    add
   </span>
   <span>
    Add New Stock
   </span>
  </button>
 </div>
 <!-- KPI Cards Row -->
 <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 mb-section-gap">
  <!-- Card 1 -->
  <div class="bg-metric-amber rounded-card p-card-padding card-lift border border-amber-200 cursor-pointer">
   <div class="flex justify-between items-start mb-2">
    <span class="font-metadata text-metadata text-amber-900 uppercase tracking-wider font-semibold">
     Total Stockrooms
    </span>
    <span class="material-symbols-outlined text-amber-700">
     warehouse
    </span>
   </div>
   <div class="font-kpi-number text-kpi-number text-amber-950 mb-1">
    12
   </div>
   <div class="font-metadata text-metadata text-amber-800">
    Across 8 buildings
   </div>
  </div>
  <!-- Card 2 -->
  <div class="bg-metric-lavender rounded-card p-card-padding card-lift border border-purple-200 cursor-pointer">
   <div class="flex justify-between items-start mb-2">
    <span class="font-metadata text-metadata text-purple-900 uppercase tracking-wider font-semibold">
     Total Assets in Stock
    </span>
    <span class="material-symbols-outlined text-purple-700">
     devices
    </span>
   </div>
   <div class="font-kpi-number text-kpi-number text-purple-950 mb-1">
    412
   </div>
   <div class="font-metadata text-metadata text-purple-800">
    Ready for checkout
   </div>
  </div>
  <!-- Card 3 -->
  <div class="bg-metric-pink rounded-card p-card-padding card-lift border border-pink-200 cursor-pointer">
   <div class="flex justify-between items-start mb-2">
    <span class="font-metadata text-metadata text-pink-900 uppercase tracking-wider font-semibold">
     Low Consumables
    </span>
    <span class="material-symbols-outlined text-pink-700">
     warning
    </span>
   </div>
   <div class="font-kpi-number text-kpi-number text-pink-950 mb-1">
    8
   </div>
   <div class="font-metadata text-metadata text-pink-800">
    Requires reorder
   </div>
  </div>
  <!-- Card 4 -->
  <div class="bg-metric-sage rounded-card p-card-padding card-lift border border-green-200 cursor-pointer">
   <div class="flex justify-between items-start mb-2">
    <span class="font-metadata text-metadata text-green-900 uppercase tracking-wider font-semibold">
     Inventory Value
    </span>
    <span class="material-symbols-outlined text-green-700">
     account_balance_wallet
    </span>
   </div>
   <div class="font-kpi-number text-kpi-number text-green-950 mb-1">
    ₹42,15,000
   </div>
   <div class="font-metadata text-metadata text-green-800">
    Fixed + Consumables
   </div>
  </div>
 </div>
 <!-- Hybrid Layout: Two Column -->
 <div class="flex flex-col lg:flex-row gap-6 mb-section-gap">
  <!-- Left Column: Locations List (240px) -->
  <div class="lg:w-[240px] flex-shrink-0">
   <div class="bg-surface border border-border-default rounded-card shadow-sm overflow-hidden h-full flex flex-col">
    <div class="p-4 border-b border-border-default bg-surface-subtle">
     <h2 class="font-h2 text-h2 text-text-primary">
      Locations
     </h2>
    </div>
    <div class="flex-1 overflow-y-auto p-2">
     <ul class="space-y-1">
      <li>
       <button class="w-full text-left px-3 py-2.5 rounded-lg bg-amber-50 border-l-4 border-amber-800 text-amber-900 flex justify-between items-center group transition-colors">
        <div class="flex flex-col">
         <span class="font-body text-body font-medium">
          Central IT Store
         </span>
         <span class="font-metadata text-metadata text-amber-700">
          Bhavani Bldg
         </span>
        </div>
        <span class="bg-amber-200 text-amber-900 text-xs px-2 py-0.5 rounded-full font-medium">
         156
        </span>
       </button>
      </li>
      <li>
       <button class="w-full text-left px-3 py-2.5 rounded-lg hover:bg-surface-subtle text-text-primary flex justify-between items-center group transition-colors border-l-4 border-transparent">
        <div class="flex flex-col">
         <span class="font-body text-body font-medium">
          Nila B-Wing Store
         </span>
         <span class="font-metadata text-metadata text-text-secondary">
          Nila Bldg
         </span>
        </div>
        <span class="bg-surface-variant text-text-secondary text-xs px-2 py-0.5 rounded-full font-medium">
         84
        </span>
       </button>
      </li>
      <li>
       <button class="w-full text-left px-3 py-2.5 rounded-lg hover:bg-surface-subtle text-text-primary flex justify-between items-center group transition-colors border-l-4 border-transparent">
        <div class="flex flex-col">
         <span class="font-body text-body font-medium">
          Facilities Supply
         </span>
         <span class="font-metadata text-metadata text-text-secondary">
          Thejaswini
         </span>
        </div>
        <span class="bg-surface-variant text-text-secondary text-xs px-2 py-0.5 rounded-full font-medium">
         112
        </span>
       </button>
      </li>
      <li>
       <button class="w-full text-left px-3 py-2.5 rounded-lg hover:bg-surface-subtle text-text-primary flex justify-between items-center group transition-colors border-l-4 border-transparent">
        <div class="flex flex-col">
         <span class="font-body text-body font-medium">
          Network Hub East
         </span>
         <span class="font-metadata text-metadata text-text-secondary">
          Gayathri
         </span>
        </div>
        <span class="bg-surface-variant text-text-secondary text-xs px-2 py-0.5 rounded-full font-medium">
         60
        </span>
       </button>
      </li>
     </ul>
    </div>
   </div>
  </div>
  <!-- Right Column: Asset Registry Table -->
  <div class="flex-1">
   <div class="bg-surface border border-border-default rounded-card shadow-sm overflow-hidden flex flex-col h-full">
    <!-- Table Toolbar -->
    <div class="p-4 border-b border-border-default flex justify-between items-center bg-surface">
     <div>
      <h2 class="font-h2 text-h2 text-text-primary">
       Central IT Store Registry
      </h2>
      <p class="font-metadata text-metadata text-text-secondary mt-0.5">
       Showing fixed assets currently in stock
      </p>
     </div>
     <div class="flex space-x-2">
      <button class="px-3 py-1.5 border border-border-default rounded-md text-sm font-medium text-text-primary hover:bg-surface-subtle transition-colors flex items-center space-x-1">
       <span class="material-symbols-outlined text-[16px]">
        filter_list
       </span>
       <span>
        Filter
       </span>
      </button>
      <button class="px-3 py-1.5 border border-border-default rounded-md text-sm font-medium text-text-primary hover:bg-surface-subtle transition-colors flex items-center space-x-1">
       <span class="material-symbols-outlined text-[16px]">
        download
       </span>
       <span>
        Export
       </span>
      </button>
     </div>
    </div>
    <!-- Table -->
    <div class="overflow-x-auto flex-1">
     <table class="w-full text-left border-collapse">
      <thead>
       <tr class="border-b border-border-default bg-surface-subtle font-table-header text-table-header text-text-secondary uppercase">
        <th class="py-3 px-4 w-12">
         <input class="rounded border-border-default text-primary-container focus:ring-primary-container" type="checkbox"/>
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
         Condition
        </th>
        <th class="py-3 px-4 text-right">
         Date Stocked
        </th>
       </tr>
      </thead>
      <tbody class="font-body text-body text-text-primary divide-y divide-border-default">
       <!-- Row 1 -->
       <tr class="hover:bg-accent-subtle transition-colors group">
        <td class="py-3 px-4 border-l-[3px] border-transparent group-hover:border-amber-500">
         <input class="rounded border-border-default text-primary-container focus:ring-primary-container" type="checkbox"/>
        </td>
        <td class="py-3 px-4 font-mono-data text-mono-data text-text-secondary">
         AST-2023-401
        </td>
        <td class="py-3 px-4 font-medium">
         Dell Latitude 5420
        </td>
        <td class="py-3 px-4 text-text-secondary">
         Laptops
        </td>
        <td class="py-3 px-4">
         <span class="inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800">
          New
         </span>
        </td>
        <td class="py-3 px-4 text-right text-text-secondary">
         12 Oct 2023
        </td>
       </tr>
       <!-- Row 2 -->
       <tr class="hover:bg-accent-subtle transition-colors group">
        <td class="py-3 px-4 border-l-[3px] border-transparent group-hover:border-amber-500">
         <input class="rounded border-border-default text-primary-container focus:ring-primary-container" type="checkbox"/>
        </td>
        <td class="py-3 px-4 font-mono-data text-mono-data text-text-secondary">
         AST-2023-405
        </td>
        <td class="py-3 px-4 font-medium">
         ThinkPad T14 Gen 2
        </td>
        <td class="py-3 px-4 text-text-secondary">
         Laptops
        </td>
        <td class="py-3 px-4">
         <span class="inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium bg-amber-100 text-amber-800">
          Refurbished
         </span>
        </td>
        <td class="py-3 px-4 text-right text-text-secondary">
         15 Oct 2023
        </td>
       </tr>
       <!-- Row 3 -->
       <tr class="hover:bg-accent-subtle transition-colors group">
        <td class="py-3 px-4 border-l-[3px] border-transparent group-hover:border-amber-500">
         <input class="rounded border-border-default text-primary-container focus:ring-primary-container" type="checkbox"/>
        </td>
        <td class="py-3 px-4 font-mono-data text-mono-data text-text-secondary">
         AST-2022-112
        </td>
        <td class="py-3 px-4 font-medium">
         Cisco Catalyst 9300
        </td>
        <td class="py-3 px-4 text-text-secondary">
         Networking
        </td>
        <td class="py-3 px-4">
         <span class="inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800">
          Good
         </span>
        </td>
        <td class="py-3 px-4 text-right text-text-secondary">
         01 Nov 2022
        </td>
       </tr>
       <!-- Row 4 -->
       <tr class="hover:bg-accent-subtle transition-colors group">
        <td class="py-3 px-4 border-l-[3px] border-transparent group-hover:border-amber-500">
         <input class="rounded border-border-default text-primary-container focus:ring-primary-container" type="checkbox"/>
        </td>
        <td class="py-3 px-4 font-mono-data text-mono-data text-text-secondary">
         AST-2023-550
        </td>
        <td class="py-3 px-4 font-medium">
         Samsung 27" Monitor
        </td>
        <td class="py-3 px-4 text-text-secondary">
         Peripherals
        </td>
        <td class="py-3 px-4">
         <span class="inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800">
          New
         </span>
        </td>
        <td class="py-3 px-4 text-right text-text-secondary">
         20 Oct 2023
        </td>
       </tr>
       <!-- Row 5 -->
       <tr class="hover:bg-accent-subtle transition-colors group">
        <td class="py-3 px-4 border-l-[3px] border-transparent group-hover:border-amber-500">
         <input class="rounded border-border-default text-primary-container focus:ring-primary-container" type="checkbox"/>
        </td>
        <td class="py-3 px-4 font-mono-data text-mono-data text-text-secondary">
         AST-2021-089
        </td>
        <td class="py-3 px-4 font-medium">
         Polycom RealPresence
        </td>
        <td class="py-3 px-4 text-text-secondary">
         A/V Equip
        </td>
        <td class="py-3 px-4">
         <span class="inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium bg-red-100 text-red-800">
          Damaged
         </span>
        </td>
        <td class="py-3 px-4 text-right text-text-secondary">
         15 Mar 2021
        </td>
       </tr>
      </tbody>
     </table>
    </div>
    <!-- Table Pagination -->
    <div class="p-3 border-t border-border-default bg-surface flex items-center justify-between">
     <span class="font-metadata text-metadata text-text-secondary">
      Showing 1 to 5 of 156 entries
     </span>
     <div class="flex space-x-1">
      <button class="p-1 rounded text-text-secondary hover:bg-surface-subtle disabled:opacity-50" disabled="">
       <span class="material-symbols-outlined text-[20px]">
        chevron_left
       </span>
      </button>
      <button class="w-7 h-7 rounded bg-primary-container text-white text-sm font-medium flex items-center justify-center">
       1
      </button>
      <button class="w-7 h-7 rounded text-text-secondary hover:bg-surface-subtle text-sm font-medium flex items-center justify-center">
       2
      </button>
      <button class="w-7 h-7 rounded text-text-secondary hover:bg-surface-subtle text-sm font-medium flex items-center justify-center">
       3
      </button>
      <button class="p-1 rounded text-text-secondary hover:bg-surface-subtle">
       <span class="material-symbols-outlined text-[20px]">
        chevron_right
       </span>
      </button>
     </div>
    </div>
   </div>
  </div>
 </div>
 <!-- Consumables Section (Bento Grid Style) -->
 <div>
  <div class="flex justify-between items-center mb-4">
   <div>
    <h2 class="font-h2 text-h2 text-text-primary">
     Consumables Inventory
    </h2>
    <p class="font-metadata text-metadata text-text-secondary mt-0.5">
     Track supplies that require frequent reordering
    </p>
   </div>
   <a class="text-primary-container hover:text-amber-900 font-medium text-sm flex items-center space-x-1" href="#">
    <span>
     View All Consumables
    </span>
    <span class="material-symbols-outlined text-[16px]">
     arrow_forward
    </span>
   </a>
  </div>
  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
   <!-- Consumable Item 1 -->
   <div class="bg-surface border border-border-default rounded-card p-5 card-lift relative overflow-hidden">
    <div class="flex justify-between items-start mb-4">
     <div class="flex items-center space-x-3">
      <div class="w-10 h-10 rounded-lg bg-blue-50 flex items-center justify-center text-blue-600">
       <span class="material-symbols-outlined">
        print
       </span>
      </div>
      <div>
       <h3 class="font-h3 text-h3 text-text-primary">
        HP 206X Toner (Black)
       </h3>
       <span class="font-metadata text-metadata text-text-secondary">
        SKU: CON-TNR-001
       </span>
      </div>
     </div>
    </div>
    <div class="mb-4">
     <div class="flex justify-between items-end mb-1">
      <span class="font-metadata text-metadata text-text-secondary">
       Stock Level
      </span>
      <span class="font-medium text-sm text-text-primary">
       4 / 20 Units
      </span>
     </div>
     <div class="w-full bg-surface-variant rounded-full h-2">
      <div class="bg-red-500 h-2 rounded-full" style="width: 20%">
      </div>
     </div>
     <p class="font-metadata text-metadata text-red-600 mt-1 flex items-center">
      <span class="material-symbols-outlined text-[14px] mr-1">
       warning
      </span>
      Below reorder point (5)
     </p>
    </div>
    <button class="w-full py-2 border border-border-default rounded-lg text-sm font-medium hover:bg-surface-subtle transition-colors flex justify-center items-center space-x-2">
     <span class="material-symbols-outlined text-[18px]">
      shopping_cart
     </span>
     <span>
      Create Reorder
     </span>
    </button>
   </div>
   <!-- Consumable Item 2 -->
   <div class="bg-surface border border-border-default rounded-card p-5 card-lift relative overflow-hidden">
    <div class="flex justify-between items-start mb-4">
     <div class="flex items-center space-x-3">
      <div class="w-10 h-10 rounded-lg bg-amber-50 flex items-center justify-center text-amber-600">
       <span class="material-symbols-outlined">
        battery_charging_full
       </span>
      </div>
      <div>
       <h3 class="font-h3 text-h3 text-text-primary">
        AA Alkaline Batteries
       </h3>
       <span class="font-metadata text-metadata text-text-secondary">
        SKU: CON-BAT-045
       </span>
      </div>
     </div>
    </div>
    <div class="mb-4">
     <div class="flex justify-between items-end mb-1">
      <span class="font-metadata text-metadata text-text-secondary">
       Stock Level
      </span>
      <span class="font-medium text-sm text-text-primary">
       145 / 200 Units
      </span>
     </div>
     <div class="w-full bg-surface-variant rounded-full h-2">
      <div class="bg-green-500 h-2 rounded-full" style="width: 72.5%">
      </div>
     </div>
     <p class="font-metadata text-metadata text-text-secondary mt-1">
      Healthy stock level
     </p>
    </div>
    <button class="w-full py-2 border border-border-default rounded-lg text-sm font-medium hover:bg-surface-subtle transition-colors flex justify-center items-center space-x-2">
     <span class="material-symbols-outlined text-[18px]">
      update
     </span>
     <span>
      Update Count
     </span>
    </button>
   </div>
   <!-- Consumable Item 3 -->
   <div class="bg-surface border border-border-default rounded-card p-5 card-lift relative overflow-hidden">
    <div class="flex justify-between items-start mb-4">
     <div class="flex items-center space-x-3">
      <div class="w-10 h-10 rounded-lg bg-purple-50 flex items-center justify-center text-purple-600">
       <span class="material-symbols-outlined">
        cable
       </span>
      </div>
      <div>
       <h3 class="font-h3 text-h3 text-text-primary">
        Cat6 Ethernet Cable (2m)
       </h3>
       <span class="font-metadata text-metadata text-text-secondary">
        SKU: CON-CBL-102
       </span>
      </div>
     </div>
    </div>
    <div class="mb-4">
     <div class="flex justify-between items-end mb-1">
      <span class="font-metadata text-metadata text-text-secondary">
       Stock Level
      </span>
      <span class="font-medium text-sm text-text-primary">
       12 / 50 Units
      </span>
     </div>
     <div class="w-full bg-surface-variant rounded-full h-2">
      <div class="bg-amber-500 h-2 rounded-full" style="width: 24%">
      </div>
     </div>
     <p class="font-metadata text-metadata text-amber-600 mt-1 flex items-center">
      <span class="material-symbols-outlined text-[14px] mr-1">
       info
      </span>
      Nearing reorder point (10)
     </p>
    </div>
    <button class="w-full py-2 border border-border-default rounded-lg text-sm font-medium hover:bg-surface-subtle transition-colors flex justify-center items-center space-x-2">
     <span class="material-symbols-outlined text-[18px]">
      shopping_cart
     </span>
     <span>
      Create Reorder
     </span>
    </button>
   </div>
  </div>
 </div>
</main>

</template>

<script setup lang="ts">
// Autogenerated from stockrooms_consumables_2
</script>
