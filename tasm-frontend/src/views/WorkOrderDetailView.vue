<template>
<main class="min-h-screen">
 <div class="max-w-[1400px] mx-auto p-page-margin">
  <!-- Page Header -->
  <div class="flex items-start justify-between mb-section-gap">
   <div>
    <div class="flex items-center gap-inline mb-2">
     <button class="text-text-secondary hover:text-text-primary flex items-center justify-center p-1 rounded hover:bg-surface-subtle transition-colors">
      <span class="material-symbols-outlined text-[20px]">
       arrow_back
      </span>
     </button>
     <span class="font-metadata text-metadata text-text-secondary">
      Operations / Work Orders
     </span>
    </div>
    <div class="flex items-center gap-4">
     <h1 class="font-h1 text-h1 text-text-primary">
      WO-2026-0842: AC Service - Nila Building
     </h1>
     <div class="flex gap-2">
      <span class="inline-flex items-center px-2.5 py-1 rounded-full bg-tertiary-fixed text-on-tertiary-fixed font-metadata text-metadata border border-tertiary-fixed-dim shadow-sm">
       <span class="w-1.5 h-1.5 rounded-full bg-status-checked-out mr-1.5">
       </span>
       In Progress
      </span>
      <span class="inline-flex items-center px-2.5 py-1 rounded-full bg-metric-amber text-primary-container font-metadata text-metadata border border-accent-subtle shadow-sm">
       <span class="material-symbols-outlined text-[14px] mr-1">
        warning
       </span>
       High Severity
      </span>
     </div>
    </div>
   </div>
   <div class="flex items-center gap-inline">
    <button class="px-4 py-2 bg-surface border border-border-default rounded-lg font-h3 text-h3 text-text-primary hover:bg-surface-subtle hover:-translate-y-[1px] transition-all shadow-sm flex items-center gap-2">
     <span class="material-symbols-outlined text-[18px]">
      print
     </span>
     Print
    </button>
    <button class="px-4 py-2 bg-primary-container text-on-primary rounded-lg font-h3 text-h3 hover:bg-primary hover:-translate-y-[1px] transition-all shadow-sm flex items-center gap-2">
     <span class="material-symbols-outlined text-[18px] filled">
      edit_document
     </span>
     Update Status
    </button>
   </div>
  </div>
  <!-- Summary Cards -->
  <div class="grid grid-cols-1 md:grid-cols-3 gap-inline mb-section-gap">
   <div class="bg-surface rounded-xl p-card-padding border border-border-default shadow-[0_4px_4px_rgba(0,0,0,0.02)] hover:-translate-y-[2px] transition-transform">
    <div class="flex items-center justify-between mb-4">
     <span class="font-metadata text-metadata text-text-secondary uppercase tracking-wider">
      Estimated Cost
     </span>
     <div class="w-8 h-8 rounded-full bg-metric-sage flex items-center justify-center text-status-in-stock">
      <span class="material-symbols-outlined text-[18px]">
       payments
      </span>
     </div>
    </div>
    <div class="font-kpi-number text-kpi-number text-text-primary">
     ₹ 14,500
    </div>
    <div class="font-metadata text-metadata text-text-secondary mt-2 flex items-center gap-1">
     <span class="material-symbols-outlined text-[14px] text-status-critical">
      trending_up
     </span>
     +12% vs initial estimate
    </div>
   </div>
   <div class="bg-surface rounded-xl p-card-padding border border-border-default shadow-[0_4px_4px_rgba(0,0,0,0.02)] hover:-translate-y-[2px] transition-transform">
    <div class="flex items-center justify-between mb-4">
     <span class="font-metadata text-metadata text-text-secondary uppercase tracking-wider">
      Parts Required
     </span>
     <div class="w-8 h-8 rounded-full bg-tertiary-fixed flex items-center justify-center text-tertiary">
      <span class="material-symbols-outlined text-[18px]">
       build_circle
      </span>
     </div>
    </div>
    <div class="font-kpi-number text-kpi-number text-text-primary">
     4 Items
    </div>
    <div class="font-metadata text-metadata text-text-secondary mt-2 flex items-center gap-1">
     <span class="w-1.5 h-1.5 rounded-full bg-status-in-stock mr-1">
     </span>
     All parts in inventory
    </div>
   </div>
   <div class="bg-surface rounded-xl p-card-padding border border-border-default shadow-[0_4px_4px_rgba(0,0,0,0.02)] hover:-translate-y-[2px] transition-transform">
    <div class="flex items-center justify-between mb-4">
     <span class="font-metadata text-metadata text-text-secondary uppercase tracking-wider">
      Time Open
     </span>
     <div class="w-8 h-8 rounded-full bg-metric-amber flex items-center justify-center text-primary-container">
      <span class="material-symbols-outlined text-[18px]">
       schedule
      </span>
     </div>
    </div>
    <div class="font-kpi-number text-kpi-number text-text-primary">
     28h 45m
    </div>
    <div class="font-metadata text-metadata text-text-secondary mt-2 flex items-center gap-1">
     <span class="material-symbols-outlined text-[14px]">
      flag
     </span>
     SLA Target: 48h
    </div>
   </div>
  </div>
  <!-- Main Layout Split -->
  <div class="grid grid-cols-1 lg:grid-cols-12 gap-section-gap">
   <!-- Left Column (60%) -->
   <div class="lg:col-span-7 flex flex-col gap-section-gap">
    <!-- Issue Details -->
    <section class="bg-surface rounded-xl border border-border-default shadow-[0_4px_4px_rgba(0,0,0,0.02)] overflow-hidden">
     <div class="p-4 border-b border-border-default bg-surface-subtle flex items-center justify-between">
      <h2 class="font-h2 text-h2 text-text-primary flex items-center gap-2">
       <span class="material-symbols-outlined text-[20px] text-text-secondary">
        description
       </span>
       Issue Details
      </h2>
      <button class="text-text-secondary hover:text-text-primary">
       <span class="material-symbols-outlined text-[20px]">
        edit
       </span>
      </button>
     </div>
     <div class="p-card-padding">
      <div class="grid grid-cols-2 gap-y-6 gap-x-4 mb-6">
       <div>
        <p class="font-metadata text-metadata text-text-secondary mb-1">
         Asset Tag
        </p>
        <p class="font-mono-data text-mono-data text-text-primary bg-surface-subtle px-2 py-1 rounded inline-block border border-border-default">
         AST-HVAC-092
        </p>
       </div>
       <div>
        <p class="font-metadata text-metadata text-text-secondary mb-1">
         Location
        </p>
        <p class="font-body text-body text-text-primary">
         Nila Building, Floor 4, Server Room B
        </p>
       </div>
       <div>
        <p class="font-metadata text-metadata text-text-secondary mb-1">
         Reported By
        </p>
        <p class="font-body text-body text-text-primary">
         Arjun Nair (Facilities)
        </p>
       </div>
       <div>
        <p class="font-metadata text-metadata text-text-secondary mb-1">
         Created Date
        </p>
        <p class="font-body text-body text-text-primary">
         Oct 24, 2023 - 09:15 AM
        </p>
       </div>
      </div>
      <div class="bg-canvas p-4 rounded-lg border border-border-default">
       <p class="font-metadata text-metadata text-text-secondary mb-2">
        Description
       </p>
       <p class="font-body text-body text-text-primary leading-relaxed">
        Main AC unit in Server Room B is making a loud grinding noise and failing to maintain target temperature of 22°C. Current room temp is 26°C. Suspected compressor or fan motor issue. Requires urgent attention to prevent server overheating.
       </p>
      </div>
     </div>
    </section>
    <!-- Work Log -->
    <section class="bg-surface rounded-xl border border-border-default shadow-[0_4px_4px_rgba(0,0,0,0.02)] overflow-hidden flex-1">
     <div class="p-4 border-b border-border-default bg-surface-subtle flex items-center justify-between">
      <h2 class="font-h2 text-h2 text-text-primary flex items-center gap-2">
       <span class="material-symbols-outlined text-[20px] text-text-secondary">
        history
       </span>
       Work Log
      </h2>
      <button class="font-h3 text-h3 text-primary-container hover:text-primary transition-colors flex items-center gap-1">
       <span class="material-symbols-outlined text-[18px]">
        add
       </span>
       Add Update
      </button>
     </div>
     <div class="p-card-padding">
      <div class="relative border-l border-border-default ml-3 space-y-8 pb-4">
       <!-- Log Item 1 -->
       <div class="relative pl-6">
        <div class="absolute -left-3 top-0 w-6 h-6 rounded-full bg-surface border-2 border-primary-container flex items-center justify-center">
         <div class="w-2 h-2 rounded-full bg-primary-container">
         </div>
        </div>
        <div class="flex justify-between items-start mb-1">
         <div class="font-h3 text-h3 text-text-primary">
          Diagnostics Completed
         </div>
         <div class="font-mono-data text-mono-data text-text-secondary">
          Today, 10:30 AM
         </div>
        </div>
        <div class="font-metadata text-metadata text-text-secondary mb-2">
         by Ramesh Kumar (HVAC Tech)
        </div>
        <div class="bg-surface-subtle p-3 rounded-lg border border-border-default font-body text-body text-text-primary">
         Inspected the unit. The condenser fan motor bearings are worn out causing the grinding noise. Compressor is functioning normally but working overtime due to poor airflow. Ordered replacement fan motor from stores.
        </div>
       </div>
       <!-- Log Item 2 -->
       <div class="relative pl-6">
        <div class="absolute -left-3 top-0 w-6 h-6 rounded-full bg-surface border-2 border-border-default flex items-center justify-center">
         <div class="w-2 h-2 rounded-full bg-border-default">
         </div>
        </div>
        <div class="flex justify-between items-start mb-1">
         <div class="font-h3 text-h3 text-text-primary">
          Work Order Assigned
         </div>
         <div class="font-mono-data text-mono-data text-text-secondary">
          Oct 24, 11:15 AM
         </div>
        </div>
        <div class="font-metadata text-metadata text-text-secondary mb-2">
         by System Admin
        </div>
        <div class="font-body text-body text-text-primary">
         Assigned to Ramesh Kumar. Priority escalated to High due to critical server room location.
        </div>
       </div>
      </div>
     </div>
    </section>
   </div>
   <!-- Right Column (40%) -->
   <div class="lg:col-span-5 flex flex-col gap-section-gap">
    <!-- Technician Assignment -->
    <section class="bg-surface rounded-xl border border-border-default shadow-[0_4px_4px_rgba(0,0,0,0.02)] overflow-hidden">
     <div class="p-4 border-b border-border-default bg-surface-subtle">
      <h2 class="font-h2 text-h2 text-text-primary flex items-center gap-2">
       <span class="material-symbols-outlined text-[20px] text-text-secondary">
        person
       </span>
       Technician Assignment
      </h2>
     </div>
     <div class="p-card-padding flex items-center justify-between">
      <div class="flex items-center gap-4">
       <div class="w-12 h-12 rounded-full bg-surface-container-high border border-border-default overflow-hidden">
        <img alt="Ramesh Kumar profile" class="w-full h-full object-cover" src="https://lh3.googleusercontent.com/aida-public/AB6AXuBVGnPGV3ajYJge28SIYlKHiA9gf81YMkJB65JA1pdH9e-faBx8QmQN8ne_Vps_YRXm3XjwdlVW7Cfq8fjfC3q7tqMTjTKnDFuZ4HPrJtqA1_JVKg7yTH1L87z4rhDYCdG-yx5IG_u4oMI4IkiUEZJTr2DEVdB_KUhL3j1BvgGlpC6ppKq-qNcslnmdkg3m4iElaflOs-aYHTKNMhy_Hqk91aShuKvNIUYI_TvyiqL4HENEhHguCAGHFygbOfcm0asLNiBtIxyDa0BU"/>
       </div>
       <div>
        <p class="font-h3 text-h3 text-text-primary">
         Ramesh Kumar
        </p>
        <p class="font-metadata text-metadata text-text-secondary">
         Senior HVAC Technician
        </p>
       </div>
      </div>
      <button aria-label="Contact Technician" class="p-2 text-text-secondary hover:text-primary-container hover:bg-metric-amber rounded-full transition-colors">
       <span class="material-symbols-outlined">
        call
       </span>
      </button>
     </div>
    </section>
    <!-- Parts & Labour Table -->
    <section class="bg-surface rounded-xl border border-border-default shadow-[0_4px_4px_rgba(0,0,0,0.02)] overflow-hidden">
     <div class="p-4 border-b border-border-default bg-surface-subtle flex items-center justify-between">
      <h2 class="font-h2 text-h2 text-text-primary flex items-center gap-2">
       <span class="material-symbols-outlined text-[20px] text-text-secondary">
        inventory
       </span>
       Parts &amp; Labour
      </h2>
     </div>
     <div class="overflow-x-auto">
      <table class="w-full text-left border-collapse">
       <thead>
        <tr class="bg-surface-subtle border-b border-border-default font-table-header text-table-header text-text-secondary">
         <th class="px-4 py-3 font-medium">
          Item
         </th>
         <th class="px-4 py-3 font-medium text-right">
          Qty
         </th>
         <th class="px-4 py-3 font-medium text-right">
          Cost
         </th>
        </tr>
       </thead>
       <tbody class="font-body text-body text-text-primary divide-y divide-border-default">
        <tr class="hover:bg-metric-amber/20 transition-colors">
         <td class="px-4 py-3">
          <div class="font-medium">
           Condenser Fan Motor
          </div>
          <div class="font-metadata text-metadata text-text-secondary">
           PRT-FN-882
          </div>
         </td>
         <td class="px-4 py-3 text-right font-mono-data text-mono-data">
          1
         </td>
         <td class="px-4 py-3 text-right font-mono-data text-mono-data">
          ₹ 8,500
         </td>
        </tr>
        <tr class="hover:bg-metric-amber/20 transition-colors">
         <td class="px-4 py-3">
          <div class="font-medium">
           Refrigerant R-410A
          </div>
          <div class="font-metadata text-metadata text-text-secondary">
           CONS-REF-01
          </div>
         </td>
         <td class="px-4 py-3 text-right font-mono-data text-mono-data">
          2 kg
         </td>
         <td class="px-4 py-3 text-right font-mono-data text-mono-data">
          ₹ 2,400
         </td>
        </tr>
        <tr class="bg-surface-subtle border-t-2 border-border-default">
         <td class="px-4 py-3 font-h3 text-h3">
          Subtotal Parts
         </td>
         <td class="px-4 py-3 text-right">
         </td>
         <td class="px-4 py-3 text-right font-mono-data text-mono-data font-medium">
          ₹ 10,900
         </td>
        </tr>
        <tr class="hover:bg-metric-amber/20 transition-colors">
         <td class="px-4 py-3">
          <div class="font-medium">
           Labour (Senior Tech)
          </div>
          <div class="font-metadata text-metadata text-text-secondary">
           Est. 4 hours
          </div>
         </td>
         <td class="px-4 py-3 text-right font-mono-data text-mono-data">
          4h
         </td>
         <td class="px-4 py-3 text-right font-mono-data text-mono-data">
          ₹ 3,600
         </td>
        </tr>
       </tbody>
      </table>
     </div>
    </section>
    <!-- Cost Summary -->
    <section class="bg-surface rounded-xl border border-border-default shadow-[0_4px_4px_rgba(0,0,0,0.02)] overflow-hidden">
     <div class="p-card-padding">
      <div class="flex justify-between items-center mb-3">
       <span class="font-body text-body text-text-secondary">
        Total Parts
       </span>
       <span class="font-mono-data text-mono-data text-text-primary">
        ₹ 10,900
       </span>
      </div>
      <div class="flex justify-between items-center mb-4">
       <span class="font-body text-body text-text-secondary">
        Total Labour
       </span>
       <span class="font-mono-data text-mono-data text-text-primary">
        ₹ 3,600
       </span>
      </div>
      <div class="border-t border-border-default pt-4 flex justify-between items-center">
       <span class="font-h2 text-h2 text-text-primary">
        Estimated Total
       </span>
       <span class="font-h1 text-h1 text-primary-container">
        ₹ 14,500
       </span>
      </div>
     </div>
    </section>
   </div>
  </div>
 </div>
</main>

</template>

<script setup lang="ts">
// Autogenerated from work_order_detail_view
</script>
