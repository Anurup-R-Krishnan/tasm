<template>
<main class="flex-1 p-page-margin max-w-[1400px] w-full mx-auto pb-32">
 <div class="mb-section-gap flex items-end justify-between">
  <div>
   <h2 class="font-h1 text-h1 text-text-primary mb-2">
    Audit Discrepancy Resolution
   </h2>
   <p class="font-body text-body text-text-secondary">
    Session ID: #AUD-2023-Q4-082 • Building C Wing
   </p>
  </div>
  <div class="flex gap-3">
   <button class="px-4 py-2 bg-surface border border-border-default rounded-lg font-body text-body text-text-primary hover:-translate-y-0.5 transition-transform shadow-sm flex items-center gap-2">
    <span class="material-symbols-outlined text-[18px]" data-icon="download">
     download
    </span>
    Export Report
   </button>
  </div>
 </div>
 <div class="grid grid-cols-5 gap-inline mb-section-gap">
  <div class="bg-surface border border-border-default rounded-xl p-4 shadow-[0_4px_4px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform cursor-pointer">
   <div class="font-metadata text-metadata text-text-secondary mb-1">
    Total Scope
   </div>
   <div class="font-kpi-number text-kpi-number text-text-primary">
    1,452
   </div>
  </div>
  <div class="bg-metric-sage border border-border-default rounded-xl p-4 shadow-[0_4px_4px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform cursor-pointer">
   <div class="font-metadata text-metadata text-text-secondary mb-1">
    Verified
   </div>
   <div class="font-kpi-number text-kpi-number text-text-primary">
    1,390
   </div>
  </div>
  <div class="bg-error-container border border-error-container rounded-xl p-4 shadow-[0_4px_4px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform cursor-pointer relative overflow-hidden">
   <div class="absolute left-0 top-0 bottom-0 w-1 bg-status-critical">
   </div>
   <div class="font-metadata text-metadata text-status-critical mb-1 pl-2">
    Missing
   </div>
   <div class="font-kpi-number text-kpi-number text-status-critical pl-2">
    18
   </div>
  </div>
  <div class="bg-metric-lavender border border-border-default rounded-xl p-4 shadow-[0_4px_4px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform cursor-pointer">
   <div class="font-metadata text-metadata text-text-secondary mb-1">
    Unregistered
   </div>
   <div class="font-kpi-number text-kpi-number text-text-primary">
    12
   </div>
  </div>
  <div class="bg-metric-amber border border-border-default rounded-xl p-4 shadow-[0_4px_4px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform cursor-pointer relative overflow-hidden">
   <div class="absolute left-0 top-0 bottom-0 w-1 bg-amber-500">
   </div>
   <div class="font-metadata text-metadata text-text-secondary mb-1 pl-2">
    Mismatches
   </div>
   <div class="font-kpi-number text-kpi-number text-text-primary pl-2">
    32
   </div>
  </div>
 </div>
 <div class="bg-surface border border-border-default rounded-xl shadow-[0_4px_4px_rgba(0,0,0,0.02)] overflow-hidden">
  <div class="p-4 border-b border-border-default flex justify-between items-center bg-surface-subtle">
   <div class="font-h2 text-h2 text-text-primary">
    Discrepancy Details
   </div>
   <div class="flex items-center gap-3">
    <select class="bg-surface border border-border-default rounded-md font-body text-body text-text-primary py-1.5 pl-3 pr-8 focus:ring-1 focus:ring-amber-500 focus:border-amber-500 text-sm">
     <option>
      All Issue Types
     </option>
     <option>
      Missing
     </option>
     <option>
      Location Mismatch
     </option>
     <option>
      Unregistered
     </option>
    </select>
   </div>
  </div>
  <div class="overflow-x-auto">
   <table class="w-full text-left border-collapse">
    <thead>
     <tr class="bg-surface border-b border-border-default">
      <th class="p-4 w-12">
       <input class="rounded border-outline-variant text-amber-600 focus:ring-amber-500" type="checkbox"/>
      </th>
      <th class="p-4 font-table-header text-table-header text-text-secondary uppercase">
       Asset Tag
      </th>
      <th class="p-4 font-table-header text-table-header text-text-secondary uppercase">
       Issue Type
      </th>
      <th class="p-4 font-table-header text-table-header text-text-secondary uppercase">
       Last Known Location
      </th>
      <th class="p-4 font-table-header text-table-header text-text-secondary uppercase">
       Scanned Location
      </th>
      <th class="p-4 font-table-header text-table-header text-text-secondary uppercase">
       Recommended Action
      </th>
      <th class="p-4 font-table-header text-table-header text-text-secondary uppercase text-right">
       Actions
      </th>
     </tr>
    </thead>
    <tbody class="divide-y divide-border-default">
     <tr class="bg-surface hover:bg-accent-subtle transition-colors group relative border-l-4 border-l-transparent hover:border-l-amber-500">
      <td class="p-4">
       <input class="rounded border-outline-variant text-amber-600 focus:ring-amber-500" type="checkbox"/>
      </td>
      <td class="p-4 font-mono-data text-mono-data text-text-primary">
       AST-992-LX
      </td>
      <td class="p-4">
       <span class="inline-flex items-center px-2 py-0.5 rounded-full font-metadata text-metadata bg-metric-amber text-on-primary-container border border-amber-200">
        Location Mismatch
       </span>
      </td>
      <td class="p-4 font-body text-body text-text-secondary">
       C-Wing, Server Rm 2
      </td>
      <td class="p-4 font-body text-body text-text-primary font-medium">
       B-Wing, IT Store
      </td>
      <td class="p-4 font-body text-body text-text-secondary">
       Update to B-Wing
      </td>
      <td class="p-4 text-right space-x-2">
       <button class="px-3 py-1.5 bg-surface border border-border-default rounded font-metadata text-metadata text-text-primary hover:bg-stone-50 transition-colors">
        Confirm New
       </button>
      </td>
     </tr>
     <tr class="bg-surface-subtle hover:bg-accent-subtle transition-colors group relative border-l-4 border-l-transparent hover:border-l-amber-500">
      <td class="p-4">
       <input class="rounded border-outline-variant text-amber-600 focus:ring-amber-500" type="checkbox"/>
      </td>
      <td class="p-4 font-mono-data text-mono-data text-text-primary">
       MAC-844-PR
      </td>
      <td class="p-4">
       <span class="inline-flex items-center px-2 py-0.5 rounded-full font-metadata text-metadata bg-error-container text-on-error-container border border-red-200">
        Missing
       </span>
      </td>
      <td class="p-4 font-body text-body text-text-primary font-medium">
       C-Wing, Desk 104
      </td>
      <td class="p-4 font-body text-body text-text-secondary italic">
       Not Scanned
      </td>
      <td class="p-4 font-body text-body text-text-secondary">
       Investigate / Mark Lost
      </td>
      <td class="p-4 text-right space-x-2">
       <button class="px-3 py-1.5 bg-surface border border-border-default rounded font-metadata text-metadata text-text-primary hover:bg-stone-50 transition-colors">
        Mark Lost
       </button>
      </td>
     </tr>
     <tr class="bg-surface hover:bg-accent-subtle transition-colors group relative border-l-4 border-l-transparent hover:border-l-amber-500">
      <td class="p-4">
       <input class="rounded border-outline-variant text-amber-600 focus:ring-amber-500" type="checkbox"/>
      </td>
      <td class="p-4 font-mono-data text-mono-data text-text-primary">
       MON-112-DK
      </td>
      <td class="p-4">
       <span class="inline-flex items-center px-2 py-0.5 rounded-full font-metadata text-metadata bg-metric-lavender text-tertiary border border-purple-200">
        Unregistered
       </span>
      </td>
      <td class="p-4 font-body text-body text-text-secondary italic">
       None
      </td>
      <td class="p-4 font-body text-body text-text-primary font-medium">
       C-Wing, Meeting Rm B
      </td>
      <td class="p-4 font-body text-body text-text-secondary">
       Register New Asset
      </td>
      <td class="p-4 text-right space-x-2">
       <button class="px-3 py-1.5 bg-[#1C1917] text-white rounded font-metadata text-metadata hover:opacity-90 transition-opacity">
        Register
       </button>
      </td>
     </tr>
     <tr class="bg-surface-subtle hover:bg-accent-subtle transition-colors group relative border-l-4 border-l-transparent hover:border-l-amber-500">
      <td class="p-4">
       <input class="rounded border-outline-variant text-amber-600 focus:ring-amber-500" type="checkbox"/>
      </td>
      <td class="p-4 font-mono-data text-mono-data text-text-primary">
       AST-945-LX
      </td>
      <td class="p-4">
       <span class="inline-flex items-center px-2 py-0.5 rounded-full font-metadata text-metadata bg-metric-amber text-on-primary-container border border-amber-200">
        Location Mismatch
       </span>
      </td>
      <td class="p-4 font-body text-body text-text-secondary">
       C-Wing, Desk 88
      </td>
      <td class="p-4 font-body text-body text-text-primary font-medium">
       C-Wing, Desk 89
      </td>
      <td class="p-4 font-body text-body text-text-secondary">
       Update to Desk 89
      </td>
      <td class="p-4 text-right space-x-2">
       <button class="px-3 py-1.5 bg-surface border border-border-default rounded font-metadata text-metadata text-text-primary hover:bg-stone-50 transition-colors">
        Confirm New
       </button>
      </td>
     </tr>
     <tr class="bg-surface hover:bg-accent-subtle transition-colors group relative border-l-4 border-l-transparent hover:border-l-amber-500">
      <td class="p-4">
       <input class="rounded border-outline-variant text-amber-600 focus:ring-amber-500" type="checkbox"/>
      </td>
      <td class="p-4 font-mono-data text-mono-data text-text-primary">
       SRV-004-PR
      </td>
      <td class="p-4">
       <span class="inline-flex items-center px-2 py-0.5 rounded-full font-metadata text-metadata bg-error-container text-on-error-container border border-red-200">
        Missing
       </span>
      </td>
      <td class="p-4 font-body text-body text-text-primary font-medium">
       Data Center Alpha
      </td>
      <td class="p-4 font-body text-body text-text-secondary italic">
       Not Scanned
      </td>
      <td class="p-4 font-body text-body text-text-secondary text-status-critical font-medium">
       High Priority Investigate
      </td>
      <td class="p-4 text-right space-x-2">
       <button class="px-3 py-1.5 bg-surface border border-border-default rounded font-metadata text-metadata text-text-primary hover:bg-stone-50 transition-colors">
        Mark Lost
       </button>
      </td>
     </tr>
    </tbody>
   </table>
  </div>
 </div>
</main>

</template>

<script setup lang="ts">
// Autogenerated from audit_discrepancy_resolution
</script>
