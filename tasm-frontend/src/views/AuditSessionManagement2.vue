<template>
<main class="md:pl-[248px] min-h-screen p-page-margin">
 <div class="max-w-[1400px] mx-auto space-y-section-gap">
  <!-- Page Header -->
  <div class="flex flex-col md:flex-row md:items-center justify-between gap-4">
   <div>
    <h1 class="font-h1 text-h1 text-text-primary">
     Audit Session Management
    </h1>
    <p class="font-body text-body text-text-secondary mt-1">
     Monitor active audits and resolve asset discrepancies across all facilities.
    </p>
   </div>
   <div class="flex items-center gap-3">
    <button class="bg-surface border border-[#D6D3CE] text-text-primary px-4 py-2 rounded-lg font-medium text-sm hover:bg-surface-subtle transition-colors flex items-center gap-2 shadow-sm">
     <span class="material-symbols-outlined text-[18px]">
      download
     </span>
     Export Report
    </button>
   </div>
  </div>
  <!-- KPI Metric Cards Grid -->
  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-inline">
   <!-- Card 1: Active Sessions -->
   <div class="bg-metric-amber rounded-xl p-card-padding shadow-[0_4px_12px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform duration-200 cursor-pointer border border-transparent hover:border-amber-200">
    <div class="flex justify-between items-start mb-4">
     <span class="font-metadata text-metadata text-amber-900 uppercase tracking-wider font-semibold">
      Active Sessions
     </span>
     <div class="w-8 h-8 rounded-full bg-amber-100 flex items-center justify-center text-amber-800">
      <span class="material-symbols-outlined text-[18px]">
       radar
      </span>
     </div>
    </div>
    <div class="font-kpi-number text-kpi-number text-amber-950 mb-1">
     2
    </div>
    <div class="font-metadata text-metadata text-amber-800/80">
     Thejaswini, Nila
    </div>
   </div>
   <!-- Card 2: Verified Assets -->
   <div class="bg-metric-lavender rounded-xl p-card-padding shadow-[0_4px_12px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform duration-200 cursor-pointer border border-transparent hover:border-violet-200">
    <div class="flex justify-between items-start mb-4">
     <span class="font-metadata text-metadata text-violet-900 uppercase tracking-wider font-semibold">
      Verified Assets
     </span>
     <div class="w-8 h-8 rounded-full bg-violet-100 flex items-center justify-center text-violet-800">
      <span class="material-symbols-outlined text-[18px]">
       verified
      </span>
     </div>
    </div>
    <div class="font-kpi-number text-kpi-number text-violet-950 mb-1">
     84%
    </div>
    <div class="font-metadata text-metadata text-violet-800/80">
     Current month progress
    </div>
   </div>
   <!-- Card 3: Discrepancies -->
   <div class="bg-primary-fixed rounded-xl p-card-padding shadow-[0_4px_12px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform duration-200 cursor-pointer border border-transparent hover:border-orange-200">
    <div class="flex justify-between items-start mb-4">
     <span class="font-metadata text-metadata text-orange-900 uppercase tracking-wider font-semibold">
      Discrepancies
     </span>
     <div class="w-8 h-8 rounded-full bg-orange-100 flex items-center justify-center text-orange-800">
      <span class="material-symbols-outlined text-[18px]">
       warning
      </span>
     </div>
    </div>
    <div class="font-kpi-number text-kpi-number text-orange-950 mb-1">
     14
    </div>
    <div class="font-metadata text-metadata text-orange-800/80">
     Unregistered or Missing
    </div>
   </div>
   <!-- Card 4: Last Full Audit -->
   <div class="bg-metric-sage rounded-xl p-card-padding shadow-[0_4px_12px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform duration-200 cursor-pointer border border-transparent hover:border-green-200">
    <div class="flex justify-between items-start mb-4">
     <span class="font-metadata text-metadata text-green-900 uppercase tracking-wider font-semibold">
      Last Full Audit
     </span>
     <div class="w-8 h-8 rounded-full bg-green-100 flex items-center justify-center text-green-800">
      <span class="material-symbols-outlined text-[18px]">
       history
      </span>
     </div>
    </div>
    <div class="font-kpi-number text-kpi-number text-green-950 mb-1">
     15 Mar
    </div>
    <div class="font-metadata text-metadata text-green-800/80">
     3472 assets verified
    </div>
   </div>
  </div>
  <!-- Bento Grid Layout for Main Content -->
  <div class="grid grid-cols-1 xl:grid-cols-12 gap-stack items-start">
   <!-- Left Column: Active Audit Sessions Table (Span 7) -->
   <div class="xl:col-span-7 bg-surface rounded-xl border border-border-default shadow-[0_4px_12px_rgba(0,0,0,0.03)] overflow-hidden flex flex-col h-full">
    <div class="p-5 border-b border-border-default flex justify-between items-center bg-surface-subtle">
     <h2 class="font-h2 text-h2 text-text-primary">
      Ongoing Audit Sessions
     </h2>
     <button class="text-primary-container hover:text-primary font-medium text-sm transition-colors flex items-center gap-1">
      View All
      <span class="material-symbols-outlined text-[16px]">
       arrow_forward
      </span>
     </button>
    </div>
    <div class="overflow-x-auto">
     <table class="w-full text-left border-collapse">
      <thead>
       <tr class="border-b border-border-default bg-surface">
        <th class="py-3 px-5 font-table-header text-table-header text-text-secondary uppercase">
         ID
        </th>
        <th class="py-3 px-5 font-table-header text-table-header text-text-secondary uppercase">
         Scope
        </th>
        <th class="py-3 px-5 font-table-header text-table-header text-text-secondary uppercase">
         Auditor
        </th>
        <th class="py-3 px-5 font-table-header text-table-header text-text-secondary uppercase">
         Progress
        </th>
        <th class="py-3 px-5 font-table-header text-table-header text-text-secondary uppercase text-right">
         Status
        </th>
       </tr>
      </thead>
      <tbody class="font-body text-body text-text-primary divide-y divide-border-default">
       <tr class="hover:bg-accent-subtle transition-colors group relative border-l-4 border-primary-container bg-surface-subtle">
        <td class="py-4 px-5 font-mono-data text-mono-data text-text-secondary">
         #AUD-2401
        </td>
        <td class="py-4 px-5">
         <div class="font-medium text-text-primary">
          Thejaswini Building
         </div>
         <div class="text-xs text-text-secondary mt-0.5">
          Floor 4 - IT ODC
         </div>
        </td>
        <td class="py-4 px-5">
         <div class="flex items-center gap-2">
          <div class="w-6 h-6 rounded-full bg-secondary-container text-xs flex items-center justify-center font-medium">
           RK
          </div>
          <span>
           Rahul K.
          </span>
         </div>
        </td>
        <td class="py-4 px-5">
         <div class="flex items-center gap-2 w-full max-w-[120px]">
          <div class="w-full bg-surface-variant rounded-full h-1.5 overflow-hidden">
           <div class="bg-primary-container h-1.5 rounded-full" style="width: 65%">
           </div>
          </div>
          <span class="text-xs font-mono-data text-text-secondary">
           65%
          </span>
         </div>
        </td>
        <td class="py-4 px-5 text-right">
         <span class="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium bg-blue-50 text-blue-700 border border-blue-200">
          <span class="w-1.5 h-1.5 rounded-full bg-blue-500 mr-1.5 animate-pulse">
          </span>
          Active
         </span>
        </td>
       </tr>
       <tr class="hover:bg-accent-subtle transition-colors group bg-surface">
        <td class="py-4 px-5 font-mono-data text-mono-data text-text-secondary pl-[21px]">
         #AUD-2402
        </td>
        <td class="py-4 px-5">
         <div class="font-medium text-text-primary">
          Nila Building
         </div>
         <div class="text-xs text-text-secondary mt-0.5">
          Ground Floor - Logistics
         </div>
        </td>
        <td class="py-4 px-5">
         <div class="flex items-center gap-2">
          <div class="w-6 h-6 rounded-full bg-secondary-container text-xs flex items-center justify-center font-medium">
           AM
          </div>
          <span>
           Anjali M.
          </span>
         </div>
        </td>
        <td class="py-4 px-5">
         <div class="flex items-center gap-2 w-full max-w-[120px]">
          <div class="w-full bg-surface-variant rounded-full h-1.5 overflow-hidden">
           <div class="bg-primary-container h-1.5 rounded-full" style="width: 22%">
           </div>
          </div>
          <span class="text-xs font-mono-data text-text-secondary">
           22%
          </span>
         </div>
        </td>
        <td class="py-4 px-5 text-right">
         <span class="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium bg-blue-50 text-blue-700 border border-blue-200">
          <span class="w-1.5 h-1.5 rounded-full bg-blue-500 mr-1.5 animate-pulse">
          </span>
          Active
         </span>
        </td>
       </tr>
       <tr class="hover:bg-accent-subtle transition-colors group bg-surface">
        <td class="py-4 px-5 font-mono-data text-mono-data text-text-secondary pl-[21px]">
         #AUD-2398
        </td>
        <td class="py-4 px-5">
         <div class="font-medium text-text-primary">
          Bhavani Building
         </div>
         <div class="text-xs text-text-secondary mt-0.5">
          Server Room A
         </div>
        </td>
        <td class="py-4 px-5">
         <div class="flex items-center gap-2">
          <div class="w-6 h-6 rounded-full bg-secondary-container text-xs flex items-center justify-center font-medium">
           SP
          </div>
          <span>
           Suresh P.
          </span>
         </div>
        </td>
        <td class="py-4 px-5">
         <div class="flex items-center gap-2 w-full max-w-[120px]">
          <div class="w-full bg-surface-variant rounded-full h-1.5 overflow-hidden">
           <div class="bg-status-in-stock h-1.5 rounded-full" style="width: 100%">
           </div>
          </div>
          <span class="text-xs font-mono-data text-text-secondary">
           100%
          </span>
         </div>
        </td>
        <td class="py-4 px-5 text-right">
         <span class="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium bg-green-50 text-green-700 border border-green-200">
          Completed
         </span>
        </td>
       </tr>
      </tbody>
     </table>
    </div>
   </div>
   <!-- Right Column: Discrepancy Resolution (Span 5) -->
   <div class="xl:col-span-5 flex flex-col gap-stack h-full">
    <div class="bg-surface rounded-xl border border-border-default shadow-[0_4px_12px_rgba(0,0,0,0.03)] flex flex-col overflow-hidden">
     <div class="p-5 border-b border-border-default flex justify-between items-center bg-error-container/20">
      <div class="flex items-center gap-2">
       <span class="material-symbols-outlined text-status-critical">
        assignment_late
       </span>
       <h2 class="font-h2 text-h2 text-text-primary">
        Action Required
       </h2>
      </div>
      <span class="bg-status-critical text-white text-xs font-bold px-2 py-1 rounded-full">
       14 Pending
      </span>
     </div>
     <div class="p-2 flex-1 overflow-y-auto max-h-[400px] space-y-2 bg-surface-subtle">
      <!-- Discrepancy Item 1 -->
      <div class="bg-surface p-4 rounded-lg border border-border-default shadow-sm hover:border-outline-variant transition-colors">
       <div class="flex justify-between items-start mb-2">
        <div>
         <div class="flex items-center gap-2 mb-1">
          <span class="font-mono-data text-mono-data text-text-secondary">
           INV-2023-8942
          </span>
          <span class="px-1.5 py-0.5 rounded text-[10px] font-bold uppercase tracking-wider bg-red-50 text-red-700 border border-red-200">
           Missing
          </span>
         </div>
         <h3 class="font-h3 text-h3 text-text-primary">
          Dell Latitude 5420 Laptop
         </h3>
         <p class="font-metadata text-metadata text-text-secondary mt-1 flex items-center gap-1">
          <span class="material-symbols-outlined text-[14px]">
           location_on
          </span>
          Expected: Thejaswini, Floor 4
         </p>
        </div>
       </div>
       <div class="mt-4 flex gap-2">
        <button class="flex-1 bg-surface border border-[#D6D3CE] text-text-primary py-1.5 rounded text-sm font-medium hover:bg-surface-subtle transition-colors">
         Locate
        </button>
        <button class="flex-1 bg-surface border border-error-container text-status-critical py-1.5 rounded text-sm font-medium hover:bg-error-container/30 transition-colors">
         Mark Lost
        </button>
       </div>
      </div>
      <!-- Discrepancy Item 2 -->
      <div class="bg-surface p-4 rounded-lg border border-border-default shadow-sm hover:border-outline-variant transition-colors">
       <div class="flex justify-between items-start mb-2">
        <div>
         <div class="flex items-center gap-2 mb-1">
          <span class="font-mono-data text-mono-data text-text-secondary">
           Unknown Asset
          </span>
          <span class="px-1.5 py-0.5 rounded text-[10px] font-bold uppercase tracking-wider bg-orange-50 text-orange-700 border border-orange-200">
           Unregistered
          </span>
         </div>
         <h3 class="font-h3 text-h3 text-text-primary">
          Cisco Meraki MR46 AP
         </h3>
         <p class="font-metadata text-metadata text-text-secondary mt-1 flex items-center gap-1">
          <span class="material-symbols-outlined text-[14px]">
           my_location
          </span>
          Found: Nila, Logistics Area
         </p>
        </div>
       </div>
       <div class="mt-4 flex gap-2">
        <button class="w-full bg-[#1C1917] hover:bg-black text-white py-1.5 rounded text-sm font-medium transition-colors">
         Register New Asset
        </button>
       </div>
      </div>
      <!-- Discrepancy Item 3 -->
      <div class="bg-surface p-4 rounded-lg border border-border-default shadow-sm hover:border-outline-variant transition-colors">
       <div class="flex justify-between items-start mb-2">
        <div>
         <div class="flex items-center gap-2 mb-1">
          <span class="font-mono-data text-mono-data text-text-secondary">
           INV-2022-4110
          </span>
          <span class="px-1.5 py-0.5 rounded text-[10px] font-bold uppercase tracking-wider bg-yellow-50 text-yellow-700 border border-yellow-200">
           Location Mismatch
          </span>
         </div>
         <h3 class="font-h3 text-h3 text-text-primary">
          Herman Miller Aeron Chair
         </h3>
         <p class="font-metadata text-metadata text-text-secondary mt-1 flex items-center gap-1">
          <span class="material-symbols-outlined text-[14px]">
           swap_horiz
          </span>
          Expected: Bhavani | Found: Thejaswini
         </p>
        </div>
       </div>
       <div class="mt-4 flex gap-2">
        <button class="flex-1 bg-primary-container text-white py-1.5 rounded text-sm font-medium hover:bg-primary transition-colors shadow-sm">
         Confirm New Location
        </button>
       </div>
      </div>
     </div>
    </div>
   </div>
  </div>
 </div>
</main>

</template>

<script setup lang="ts">
// Autogenerated from audit_session_management_2
</script>
