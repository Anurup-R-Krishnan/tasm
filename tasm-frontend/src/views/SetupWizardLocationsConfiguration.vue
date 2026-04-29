<template>
<main class="flex-1 p-page-margin max-w-[1400px] w-full mx-auto flex flex-col gap-section-gap">
 <!-- Header Section -->
 <div class="flex flex-col gap-base">
  <h1 class="font-h1 text-h1 text-text-primary">
   System Setup Wizard
  </h1>
  <p class="font-body text-body text-text-secondary">
   Configure your basic organisational hierarchy to get started.
  </p>
 </div>
 <!-- Step Indicator -->
 <div class="w-full bg-surface border border-border-default rounded-xl p-6 shadow-sm flex items-center justify-between">
  <!-- Step 1 -->
  <div class="flex flex-col items-center gap-2 relative z-10 flex-1">
   <div class="w-8 h-8 rounded-full bg-status-in-stock text-white flex items-center justify-center font-medium text-sm">
    <span class="material-symbols-outlined text-[16px]">
     check
    </span>
   </div>
   <span class="font-metadata text-metadata text-text-secondary whitespace-nowrap">
    Organisation
   </span>
  </div>
  <!-- Connector -->
  <div class="h-px bg-status-in-stock flex-1 -mx-8">
  </div>
  <!-- Step 2 -->
  <div class="flex flex-col items-center gap-2 relative z-10 flex-1">
   <div class="w-8 h-8 rounded-full bg-status-in-stock text-white flex items-center justify-center font-medium text-sm">
    <span class="material-symbols-outlined text-[16px]">
     check
    </span>
   </div>
   <span class="font-metadata text-metadata text-text-secondary whitespace-nowrap">
    Buildings
   </span>
  </div>
  <!-- Connector -->
  <div class="h-px bg-amber-600 flex-1 -mx-8">
  </div>
  <!-- Step 3 (Active) -->
  <div class="flex flex-col items-center gap-2 relative z-10 flex-1">
   <div class="w-8 h-8 rounded-full bg-amber-600 text-white flex items-center justify-center font-medium text-sm ring-4 ring-amber-100">
    3
   </div>
   <span class="font-metadata text-metadata text-text-primary font-medium whitespace-nowrap">
    Locations
   </span>
  </div>
  <!-- Connector -->
  <div class="h-px bg-stone-200 flex-1 -mx-8">
  </div>
  <!-- Step 4 -->
  <div class="flex flex-col items-center gap-2 relative z-10 flex-1 opacity-50">
   <div class="w-8 h-8 rounded-full bg-stone-100 border border-stone-300 text-stone-500 flex items-center justify-center font-medium text-sm">
    4
   </div>
   <span class="font-metadata text-metadata text-text-secondary whitespace-nowrap">
    Departments
   </span>
  </div>
  <!-- Connector -->
  <div class="h-px bg-stone-200 flex-1 -mx-8">
  </div>
  <!-- Step 5 -->
  <div class="flex flex-col items-center gap-2 relative z-10 flex-1 opacity-50">
   <div class="w-8 h-8 rounded-full bg-stone-100 border border-stone-300 text-stone-500 flex items-center justify-center font-medium text-sm">
    5
   </div>
   <span class="font-metadata text-metadata text-text-secondary whitespace-nowrap">
    Categories
   </span>
  </div>
  <!-- Connector -->
  <div class="h-px bg-stone-200 flex-1 -mx-8">
  </div>
  <!-- Step 6 -->
  <div class="flex flex-col items-center gap-2 relative z-10 flex-1 opacity-50">
   <div class="w-8 h-8 rounded-full bg-stone-100 border border-stone-300 text-stone-500 flex items-center justify-center font-medium text-sm">
    6
   </div>
   <span class="font-metadata text-metadata text-text-secondary whitespace-nowrap">
    Fields
   </span>
  </div>
 </div>
 <!-- Main Content Card -->
 <div class="bg-surface border border-border-default rounded-xl shadow-sm overflow-hidden flex flex-col">
  <div class="p-card-padding border-b border-border-default flex justify-between items-center bg-white">
   <div>
    <h2 class="font-h2 text-h2 text-text-primary">
     Configure Locations (Rooms)
    </h2>
    <p class="font-metadata text-metadata text-text-secondary mt-1">
     Define the specific rooms or areas within your buildings where assets will be located.
    </p>
   </div>
   <div class="flex items-center gap-3 text-sm text-text-secondary bg-surface-subtle px-3 py-1.5 rounded-md border border-border-default">
    <span class="material-symbols-outlined text-[18px]">
     info
    </span>
    <span>
     Total Locations Configured:
     <span class="font-mono-data text-mono-data font-medium text-text-primary">
      1,45,000
     </span>
    </span>
   </div>
  </div>
  <div class="w-full overflow-x-auto">
   <table class="w-full text-left border-collapse">
    <thead>
     <tr class="bg-surface-subtle border-b border-border-default">
      <th class="py-3 px-4 font-table-header text-table-header text-text-secondary uppercase w-[250px]">
       Building
      </th>
      <th class="py-3 px-4 font-table-header text-table-header text-text-secondary uppercase w-[250px]">
       Room Name/Number
      </th>
      <th class="py-3 px-4 font-table-header text-table-header text-text-secondary uppercase w-[150px]">
       Floor Number
      </th>
      <th class="py-3 px-4 font-table-header text-table-header text-text-secondary uppercase">
       Description
      </th>
      <th class="py-3 px-4 font-table-header text-table-header text-text-secondary uppercase w-[80px]">
       Actions
      </th>
     </tr>
    </thead>
    <tbody class="divide-y divide-border-default">
     <!-- Existing Row 1 -->
     <tr class="bg-white hover:bg-accent-subtle transition-colors group">
      <td class="py-3 px-4">
       <div class="flex items-center gap-2">
        <span class="w-1.5 h-1.5 rounded-full bg-status-in-stock">
        </span>
        <span class="font-body text-body text-text-primary">
         Tejaswini (Block A)
        </span>
       </div>
      </td>
      <td class="py-3 px-4 font-body text-body text-text-primary">
       Server Room Alpha
      </td>
      <td class="py-3 px-4 font-mono-data text-mono-data text-text-primary">
       02
      </td>
      <td class="py-3 px-4 font-body text-body text-text-secondary truncate max-w-[300px]">
       Primary data center facility with raised flooring.
      </td>
      <td class="py-3 px-4">
       <button class="text-stone-400 hover:text-error transition-colors p-1 rounded">
        <span class="material-symbols-outlined text-[20px]">
         delete
        </span>
       </button>
      </td>
     </tr>
     <!-- Existing Row 2 -->
     <tr class="bg-surface-subtle hover:bg-accent-subtle transition-colors group">
      <td class="py-3 px-4">
       <div class="flex items-center gap-2">
        <span class="w-1.5 h-1.5 rounded-full bg-status-in-stock">
        </span>
        <span class="font-body text-body text-text-primary">
         Nila (Block B)
        </span>
       </div>
      </td>
      <td class="py-3 px-4 font-body text-body text-text-primary">
       Conference Room 104
      </td>
      <td class="py-3 px-4 font-mono-data text-mono-data text-text-primary">
       01
      </td>
      <td class="py-3 px-4 font-body text-body text-text-secondary truncate max-w-[300px]">
       12-seater conference room with video conferencing.
      </td>
      <td class="py-3 px-4">
       <button class="text-stone-400 hover:text-error transition-colors p-1 rounded">
        <span class="material-symbols-outlined text-[20px]">
         delete
        </span>
       </button>
      </td>
     </tr>
     <!-- Existing Row 3 -->
     <tr class="bg-white hover:bg-accent-subtle transition-colors group border-l-[3px] border-amber-600">
      <td class="py-3 px-4 pl-[13px]">
       <div class="flex items-center gap-2">
        <span class="w-1.5 h-1.5 rounded-full bg-status-in-stock">
        </span>
        <span class="font-body text-body text-text-primary">
         Bhavani (Block C)
        </span>
       </div>
      </td>
      <td class="py-3 px-4 font-body text-body text-text-primary">
       Executive Boardroom
      </td>
      <td class="py-3 px-4 font-mono-data text-mono-data text-text-primary">
       05
      </td>
      <td class="py-3 px-4 font-body text-body text-text-secondary truncate max-w-[300px]">
       Main boardroom for senior management meetings.
      </td>
      <td class="py-3 px-4">
       <button class="text-stone-400 hover:text-error transition-colors p-1 rounded">
        <span class="material-symbols-outlined text-[20px]">
         delete
        </span>
       </button>
      </td>
     </tr>
     <!-- Add New Row Appender -->
     <tr class="bg-amber-50/30 border-t-2 border-amber-100">
      <td class="py-3 px-4">
       <select class="w-full bg-white border border-stone-300 rounded-md py-1.5 px-3 text-sm focus:outline-none focus:border-amber-600 focus:ring-1 focus:ring-amber-600">
        <option disabled selected value="">
         Select Building...
        </option>
        <option value="tejaswini">
         Tejaswini (Block A)
        </option>
        <option value="nila">
         Nila (Block B)
        </option>
        <option value="bhavani">
         Bhavani (Block C)
        </option>
       </select>
      </td>
      <td class="py-3 px-4">
       <input class="w-full bg-white border border-stone-300 rounded-md py-1.5 px-3 text-sm focus:outline-none focus:border-amber-600 focus:ring-1 focus:ring-amber-600" placeholder="e.g. IT Store Room" type="text"/>
      </td>
      <td class="py-3 px-4">
       <input class="w-full bg-white border border-stone-300 rounded-md py-1.5 px-3 text-sm focus:outline-none focus:border-amber-600 focus:ring-1 focus:ring-amber-600 font-mono-data" placeholder="e.g. 03" type="text"/>
      </td>
      <td class="py-3 px-4">
       <input class="w-full bg-white border border-stone-300 rounded-md py-1.5 px-3 text-sm focus:outline-none focus:border-amber-600 focus:ring-1 focus:ring-amber-600" placeholder="Brief description of location use..." type="text"/>
      </td>
      <td class="py-3 px-4">
       <button class="text-amber-700 hover:bg-amber-100 transition-colors p-1.5 rounded-md flex items-center justify-center w-full border border-amber-200 bg-white">
        <span class="material-symbols-outlined text-[20px]">
         add
        </span>
       </button>
      </td>
     </tr>
    </tbody>
   </table>
  </div>
 </div>
 <!-- Action Bar -->
 <div class="flex justify-between items-center mt-auto pt-6 border-t border-border-default">
  <button class="px-6 py-2 bg-white border border-border-default text-text-primary rounded-md font-body text-body font-medium hover:bg-stone-50 transition-colors flex items-center gap-2 shadow-sm">
   <span class="material-symbols-outlined text-[18px]">
    arrow_back
   </span>
   Back to Buildings
  </button>
  <div class="flex gap-4">
   <button class="px-6 py-2 bg-white border border-border-default text-text-secondary rounded-md font-body text-body font-medium hover:bg-stone-50 transition-colors shadow-sm">
    Save as Draft
   </button>
   <button class="px-6 py-2 bg-text-primary text-white rounded-md font-body text-body font-medium hover:bg-stone-800 transition-colors flex items-center gap-2 shadow-sm">
    Continue to Departments
    <span class="material-symbols-outlined text-[18px]">
     arrow_forward
    </span>
   </button>
  </div>
 </div>
</main>

</template>

<script setup lang="ts">
// Autogenerated from setup_wizard_locations_configuration
</script>
