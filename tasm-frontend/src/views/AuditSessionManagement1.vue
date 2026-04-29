<template>
<main class="flex-1 mt-[60px] p-page-margin max-w-[1400px]">
 <!-- Page Header & Actions -->
 <div class="flex justify-between items-end mb-section-gap">
  <div>
   <h1 class="font-h1 text-h1 text-text-primary mb-1">
    Audit Sessions
   </h1>
   <p class="font-metadata text-metadata text-text-secondary">
    Manage and track physical asset verifications across facilities.
   </p>
  </div>
  <button class="bg-[#1C1917] hover:bg-stone-800 text-white rounded font-body text-body py-2 px-5 transition-colors shadow-sm flex items-center gap-2">
   <span class="material-symbols-outlined" style="font-size: 18px;">
    qr_code_scanner
   </span>
   Start New Audit Session
  </button>
 </div>
 <!-- Section 1: Active/Past Sessions Table -->
 <section class="bg-surface border border-border-default rounded-xl shadow-[0_4px_4px_rgba(0,0,0,0.02)] mb-section-gap overflow-hidden">
  <div class="px-card-padding py-4 border-b border-border-default bg-surface-subtle flex justify-between items-center">
   <h2 class="font-h2 text-h2 text-text-primary">
    Recent Audit Sessions
   </h2>
   <button class="text-text-secondary hover:text-primary transition-colors">
    <span class="material-symbols-outlined">
     filter_list
    </span>
   </button>
  </div>
  <div class="overflow-x-auto">
   <table class="w-full text-left border-collapse">
    <thead>
     <tr class="border-b border-border-default bg-surface">
      <th class="font-table-header text-table-header text-text-secondary py-3 px-6 uppercase">
       Session ID
      </th>
      <th class="font-table-header text-table-header text-text-secondary py-3 px-6 uppercase">
       Scope
      </th>
      <th class="font-table-header text-table-header text-text-secondary py-3 px-6 uppercase">
       Date
      </th>
      <th class="font-table-header text-table-header text-text-secondary py-3 px-6 uppercase">
       Verified %
      </th>
      <th class="font-table-header text-table-header text-text-secondary py-3 px-6 uppercase">
       Missing
      </th>
      <th class="font-table-header text-table-header text-text-secondary py-3 px-6 uppercase">
       Unregistered
      </th>
      <th class="font-table-header text-table-header text-text-secondary py-3 px-6 uppercase">
       Status
      </th>
      <th class="font-table-header text-table-header text-text-secondary py-3 px-6 uppercase text-right">
       Action
      </th>
     </tr>
    </thead>
    <tbody class="font-body text-body">
     <tr class="border-b border-border-default hover:bg-metric-amber transition-colors bg-surface group">
      <td class="py-4 px-6 font-mono-data text-mono-data text-text-primary border-l-[3px] border-amber-600">
       AUD-2023-11-A
      </td>
      <td class="py-4 px-6 text-text-primary">
       Gayathri - Floor 2
      </td>
      <td class="py-4 px-6 text-text-secondary">
       Nov 14, 2023
      </td>
      <td class="py-4 px-6">
       <div class="flex items-center gap-2">
        <span class="font-medium">
         85%
        </span>
        <div class="w-16 h-1.5 bg-surface-variant rounded-full overflow-hidden">
         <div class="h-full bg-amber-500 w-[85%] rounded-full">
         </div>
        </div>
       </div>
      </td>
      <td class="py-4 px-6 text-status-critical font-medium">
       12
      </td>
      <td class="py-4 px-6 text-primary-container">
       3
      </td>
      <td class="py-4 px-6">
       <span class="inline-flex items-center px-2 py-1 rounded bg-metric-lavender text-tertiary-container font-metadata text-metadata font-medium">
        In Progress
       </span>
      </td>
      <td class="py-4 px-6 text-right">
       <button class="text-primary hover:text-amber-800 font-medium transition-colors">
        View Report
       </button>
      </td>
     </tr>
     <tr class="border-b border-border-default hover:bg-metric-amber transition-colors bg-surface-subtle">
      <td class="py-4 px-6 font-mono-data text-mono-data text-text-primary border-l-[3px] border-transparent">
       AUD-2023-10-F
      </td>
      <td class="py-4 px-6 text-text-primary">
       Bhavani - IT Server Room
      </td>
      <td class="py-4 px-6 text-text-secondary">
       Oct 28, 2023
      </td>
      <td class="py-4 px-6">
       <div class="flex items-center gap-2">
        <span class="font-medium text-status-in-stock">
         100%
        </span>
        <div class="w-16 h-1.5 bg-surface-variant rounded-full overflow-hidden">
         <div class="h-full bg-status-in-stock w-[100%] rounded-full">
         </div>
        </div>
       </div>
      </td>
      <td class="py-4 px-6 text-text-secondary">
       0
      </td>
      <td class="py-4 px-6 text-text-secondary">
       0
      </td>
      <td class="py-4 px-6">
       <span class="inline-flex items-center px-2 py-1 rounded bg-metric-sage text-status-in-stock font-metadata text-metadata font-medium">
        Completed
       </span>
      </td>
      <td class="py-4 px-6 text-right">
       <button class="text-primary hover:text-amber-800 font-medium transition-colors">
        View Report
       </button>
      </td>
     </tr>
     <tr class="border-b border-border-default hover:bg-metric-amber transition-colors bg-surface">
      <td class="py-4 px-6 font-mono-data text-mono-data text-text-primary border-l-[3px] border-transparent">
       AUD-2023-10-E
      </td>
      <td class="py-4 px-6 text-text-primary">
       Thejaswini - Ground Floor
      </td>
      <td class="py-4 px-6 text-text-secondary">
       Oct 15, 2023
      </td>
      <td class="py-4 px-6">
       <div class="flex items-center gap-2">
        <span class="font-medium">
         98%
        </span>
        <div class="w-16 h-1.5 bg-surface-variant rounded-full overflow-hidden">
         <div class="h-full bg-status-in-stock w-[98%] rounded-full">
         </div>
        </div>
       </div>
      </td>
      <td class="py-4 px-6 text-status-critical font-medium">
       2
      </td>
      <td class="py-4 px-6 text-text-secondary">
       0
      </td>
      <td class="py-4 px-6">
       <span class="inline-flex items-center px-2 py-1 rounded bg-metric-sage text-status-in-stock font-metadata text-metadata font-medium">
        Completed
       </span>
      </td>
      <td class="py-4 px-6 text-right">
       <button class="text-primary hover:text-amber-800 font-medium transition-colors">
        View Report
       </button>
      </td>
     </tr>
    </tbody>
   </table>
  </div>
  <div class="px-card-padding py-3 border-t border-border-default bg-surface flex justify-end">
   <button class="text-text-secondary hover:text-text-primary font-metadata text-metadata transition-colors flex items-center gap-1">
    View All Sessions
    <span class="material-symbols-outlined" style="font-size: 16px;">
     arrow_forward
    </span>
   </button>
  </div>
 </section>
 <!-- Section 2: Discrepancy Resolution Queue (Bento/Card layout) -->
 <h2 class="font-h2 text-h2 text-text-primary mb-4">
  Discrepancy Resolution Queue
 </h2>
 <div class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-6">
  <!-- Discrepancy Card 1 -->
  <div class="bg-surface border border-border-default rounded-xl p-card-padding shadow-[0_4px_4px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform duration-200 flex flex-col h-full">
   <div class="flex justify-between items-start mb-4">
    <div class="flex items-center gap-2">
     <span class="material-symbols-outlined text-status-critical bg-error-container p-1.5 rounded-lg">
      warning
     </span>
     <span class="font-metadata text-metadata text-status-critical font-medium uppercase tracking-wider">
      Missing
     </span>
    </div>
    <span class="text-status-critical font-metadata text-metadata font-bold bg-error-container px-2 py-1 rounded">
     4 Days
    </span>
   </div>
   <h3 class="font-h3 text-h3 text-text-primary mb-1">
    Dell Latitude 5420
   </h3>
   <p class="font-mono-data text-mono-data text-text-secondary mb-4">
    AST-LPT-8892
   </p>
   <div class="bg-surface-subtle p-3 rounded border border-border-default mb-4 flex-1">
    <p class="font-metadata text-metadata text-text-secondary mb-1">
     Flagged In:
    </p>
    <p class="font-body text-body text-text-primary">
     AUD-2023-11-A (Gayathri - Floor 2)
    </p>
   </div>
   <div class="mt-auto pt-4 border-t border-border-default flex justify-end gap-2">
    <select class="form-select w-full bg-surface border border-[#D6D3CE] rounded text-body font-body text-text-primary py-1.5 px-3 focus:ring-primary focus:border-primary">
     <option disabled="" selected="" value="">
      Select Action...
     </option>
     <option value="confirm">
      Confirm Location
     </option>
     <option value="lost">
      Mark as Lost
     </option>
     <option value="reassign">
      Reassign Asset
     </option>
    </select>
    <button class="bg-surface border border-[#D6D3CE] hover:bg-surface-subtle text-text-primary rounded p-1.5 transition-colors flex-shrink-0">
     <span class="material-symbols-outlined">
      check
     </span>
    </button>
   </div>
  </div>
  <!-- Discrepancy Card 2 -->
  <div class="bg-surface border border-border-default rounded-xl p-card-padding shadow-[0_4px_4px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform duration-200 flex flex-col h-full">
   <div class="flex justify-between items-start mb-4">
    <div class="flex items-center gap-2">
     <span class="material-symbols-outlined text-primary-container bg-accent-subtle p-1.5 rounded-lg">
      location_off
     </span>
     <span class="font-metadata text-metadata text-primary-container font-medium uppercase tracking-wider">
      Location Mismatch
     </span>
    </div>
    <span class="text-primary-container font-metadata text-metadata font-bold bg-accent-subtle px-2 py-1 rounded">
     2 Days
    </span>
   </div>
   <h3 class="font-h3 text-h3 text-text-primary mb-1">
    Herman Miller Aeron
   </h3>
   <p class="font-mono-data text-mono-data text-text-secondary mb-4">
    AST-FUR-1104
   </p>
   <div class="bg-surface-subtle p-3 rounded border border-border-default mb-4 flex-1">
    <p class="font-metadata text-metadata text-text-secondary mb-1">
     Flagged In:
    </p>
    <p class="font-body text-body text-text-primary">
     AUD-2023-11-A (Gayathri - Floor 2)
    </p>
    <p class="font-metadata text-metadata text-text-secondary mt-2">
     Expected: Nila - Floor 4
    </p>
   </div>
   <div class="mt-auto pt-4 border-t border-border-default flex justify-end gap-2">
    <select class="form-select w-full bg-surface border border-[#D6D3CE] rounded text-body font-body text-text-primary py-1.5 px-3 focus:ring-primary focus:border-primary">
     <option disabled="" selected="" value="">
      Select Action...
     </option>
     <option value="confirm">
      Update Location to Gayathri
     </option>
     <option value="lost">
      Mark for Return to Nila
     </option>
    </select>
    <button class="bg-surface border border-[#D6D3CE] hover:bg-surface-subtle text-text-primary rounded p-1.5 transition-colors flex-shrink-0">
     <span class="material-symbols-outlined">
      check
     </span>
    </button>
   </div>
  </div>
  <!-- Discrepancy Card 3 (Unregistered) -->
  <div class="bg-surface border border-border-default rounded-xl p-card-padding shadow-[0_4px_4px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform duration-200 flex flex-col h-full">
   <div class="flex justify-between items-start mb-4">
    <div class="flex items-center gap-2">
     <span class="material-symbols-outlined text-tertiary bg-tertiary-fixed p-1.5 rounded-lg">
      help_center
     </span>
     <span class="font-metadata text-metadata text-tertiary font-medium uppercase tracking-wider">
      Unregistered
     </span>
    </div>
    <span class="text-tertiary font-metadata text-metadata font-bold bg-tertiary-fixed px-2 py-1 rounded">
     1 Day
    </span>
   </div>
   <h3 class="font-h3 text-h3 text-text-primary mb-1">
    Unknown Cisco Switch
   </h3>
   <p class="font-mono-data text-mono-data text-text-secondary mb-4">
    No Asset Tag
   </p>
   <div class="bg-surface-subtle p-3 rounded border border-border-default mb-4 flex-1">
    <p class="font-metadata text-metadata text-text-secondary mb-1">
     Flagged In:
    </p>
    <p class="font-body text-body text-text-primary">
     AUD-2023-11-A (Gayathri - Floor 2)
    </p>
    <p class="font-metadata text-metadata text-text-secondary mt-2">
     Notes: Found in network closet G2-B.
    </p>
   </div>
   <div class="mt-auto pt-4 border-t border-border-default flex justify-end gap-2">
    <select class="form-select w-full bg-surface border border-[#D6D3CE] rounded text-body font-body text-text-primary py-1.5 px-3 focus:ring-primary focus:border-primary">
     <option disabled="" selected="" value="">
      Select Action...
     </option>
     <option value="register">
      Register New Asset
     </option>
     <option value="ignore">
      Ignore/Not Asset
     </option>
    </select>
    <button class="bg-surface border border-[#D6D3CE] hover:bg-surface-subtle text-text-primary rounded p-1.5 transition-colors flex-shrink-0">
     <span class="material-symbols-outlined">
      check
     </span>
    </button>
   </div>
  </div>
 </div>
</main>

</template>

<script setup lang="ts">
// Autogenerated from audit_session_management_1
</script>
