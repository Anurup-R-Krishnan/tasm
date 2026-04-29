<template>
<main class="flex-1 md:ml-[248px] p-page-margin w-full max-w-[1400px] mx-auto">
 <!-- Page Header -->
 <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4 mb-section-gap">
  <div>
   <h1 class="font-h1 text-h1 text-text-primary">
    System Alerts
   </h1>
   <p class="font-metadata text-metadata text-text-secondary mt-1">
    Manage and resolve system notifications and warnings.
   </p>
  </div>
  <div class="flex items-center gap-3">
   <button class="bg-surface border border-[#D6D3CE] text-text-primary px-4 py-2 rounded-lg font-h3 text-h3 hover:bg-surface-subtle hover:-translate-y-[2px] transition-all duration-200 shadow-[0_4px_4px_rgba(0,0,0,0.02)]">
    Mark All Read
   </button>
   <button class="bg-[#1C1917] text-white px-4 py-2 rounded-lg font-h3 text-h3 hover:bg-stone-800 hover:-translate-y-[2px] transition-all duration-200 flex items-center gap-2 shadow-[0_4px_4px_rgba(0,0,0,0.02)]">
    <span class="material-symbols-outlined text-[18px]">
     settings
    </span>
    Alert Rules
   </button>
  </div>
 </div>
 <!-- Filter Bar -->
 <div class="bg-surface rounded-xl border border-border-default shadow-[0_4px_4px_rgba(0,0,0,0.02)] p-2 mb-stack flex overflow-x-auto hide-scrollbar">
  <button class="px-4 py-1.5 rounded-lg bg-surface-subtle text-text-primary font-h3 text-h3 whitespace-nowrap">
   All Alerts (12)
  </button>
  <button class="px-4 py-1.5 rounded-lg text-text-secondary hover:text-text-primary hover:bg-surface-subtle font-h3 text-h3 whitespace-nowrap transition-colors flex items-center gap-2">
   <span class="w-2 h-2 rounded-full bg-status-critical">
   </span>
   Critical (3)
  </button>
  <button class="px-4 py-1.5 rounded-lg text-text-secondary hover:text-text-primary hover:bg-surface-subtle font-h3 text-h3 whitespace-nowrap transition-colors flex items-center gap-2">
   <span class="w-2 h-2 rounded-full bg-primary-container">
   </span>
   Warning (5)
  </button>
  <button class="px-4 py-1.5 rounded-lg text-text-secondary hover:text-text-primary hover:bg-surface-subtle font-h3 text-h3 whitespace-nowrap transition-colors flex items-center gap-2">
   <span class="w-2 h-2 rounded-full bg-status-checked-out">
   </span>
   Info (4)
  </button>
  <div class="ml-auto flex items-center gap-2 pl-4 border-l border-border-default">
   <button class="p-1.5 text-text-secondary hover:text-text-primary rounded-lg hover:bg-surface-subtle transition-colors">
    <span class="material-symbols-outlined text-[20px]">
     filter_list
    </span>
   </button>
  </div>
 </div>
 <!-- Alerts List -->
 <div class="space-y-inline">
  <!-- Critical Alert Card -->
  <div class="bg-surface rounded-xl border border-border-default shadow-[0_4px_4px_rgba(0,0,0,0.02)] hover:-translate-y-[2px] transition-transform duration-200 overflow-hidden flex flex-col sm:flex-row relative">
   <div class="absolute left-0 top-0 bottom-0 w-[3px] bg-status-critical">
   </div>
   <div class="p-card-padding flex-1 flex flex-col sm:flex-row gap-4 sm:items-center pl-6">
    <div class="flex-shrink-0 w-10 h-10 rounded-full bg-error-container text-status-critical flex items-center justify-center">
     <span class="material-symbols-outlined text-[20px]" style="font-variation-settings: 'FILL' 1;">
      error
     </span>
    </div>
    <div class="flex-1">
     <div class="flex items-center gap-2 mb-1">
      <span class="px-2 py-0.5 rounded text-[11px] font-semibold tracking-wide bg-error-container text-status-critical uppercase">
       Warranty Expired
      </span>
      <span class="font-metadata text-metadata text-text-secondary">
       2 hours ago
      </span>
     </div>
     <p class="font-body text-body text-text-primary">
      Main Generator B (
      <a class="font-h3 text-h3 hover:underline" href="#">
       GEN-B-002
      </a>
      ) warranty has expired. Immediate renewal required to maintain service SLA.
     </p>
    </div>
    <div class="flex-shrink-0 sm:ml-4 mt-4 sm:mt-0 flex gap-2">
     <button class="bg-surface border border-[#D6D3CE] text-text-primary px-3 py-1.5 rounded-lg font-h3 text-[13px] hover:bg-surface-subtle transition-colors">
      Dismiss
     </button>
     <button class="bg-primary-container text-white px-3 py-1.5 rounded-lg font-h3 text-[13px] hover:bg-primary transition-colors">
      Renew Warranty
     </button>
    </div>
   </div>
  </div>
  <!-- Warning Alert Card -->
  <div class="bg-surface rounded-xl border border-border-default shadow-[0_4px_4px_rgba(0,0,0,0.02)] hover:-translate-y-[2px] transition-transform duration-200 overflow-hidden flex flex-col sm:flex-row relative">
   <div class="absolute left-0 top-0 bottom-0 w-[3px] bg-primary-container">
   </div>
   <div class="p-card-padding flex-1 flex flex-col sm:flex-row gap-4 sm:items-center pl-6">
    <div class="flex-shrink-0 w-10 h-10 rounded-full bg-metric-amber text-primary-container flex items-center justify-center">
     <span class="material-symbols-outlined text-[20px]" style="font-variation-settings: 'FILL' 1;">
      warning
     </span>
    </div>
    <div class="flex-1">
     <div class="flex items-center gap-2 mb-1">
      <span class="px-2 py-0.5 rounded text-[11px] font-semibold tracking-wide bg-metric-amber text-primary-container uppercase">
       Maintenance Due
      </span>
      <span class="font-metadata text-metadata text-text-secondary">
       5 hours ago
      </span>
     </div>
     <p class="font-body text-body text-text-primary">
      HVAC Unit 4 (
      <a class="font-h3 text-h3 hover:underline" href="#">
       HVAC-004
      </a>
      ) scheduled maintenance is due in 3 days.
     </p>
    </div>
    <div class="flex-shrink-0 sm:ml-4 mt-4 sm:mt-0 flex gap-2">
     <button class="bg-surface border border-[#D6D3CE] text-text-primary px-3 py-1.5 rounded-lg font-h3 text-[13px] hover:bg-surface-subtle transition-colors">
      Dismiss
     </button>
     <button class="bg-surface border border-border-default text-text-primary px-3 py-1.5 rounded-lg font-h3 text-[13px] hover:bg-surface-subtle transition-colors">
      Schedule
     </button>
    </div>
   </div>
  </div>
  <!-- Info Alert Card -->
  <div class="bg-surface rounded-xl border border-border-default shadow-[0_4px_4px_rgba(0,0,0,0.02)] hover:-translate-y-[2px] transition-transform duration-200 overflow-hidden flex flex-col sm:flex-row relative opacity-75 hover:opacity-100">
   <div class="absolute left-0 top-0 bottom-0 w-[3px] bg-status-checked-out">
   </div>
   <div class="p-card-padding flex-1 flex flex-col sm:flex-row gap-4 sm:items-center pl-6">
    <div class="flex-shrink-0 w-10 h-10 rounded-full bg-tertiary-fixed text-status-checked-out flex items-center justify-center">
     <span class="material-symbols-outlined text-[20px]">
      info
     </span>
    </div>
    <div class="flex-1">
     <div class="flex items-center gap-2 mb-1">
      <span class="px-2 py-0.5 rounded text-[11px] font-semibold tracking-wide bg-tertiary-fixed text-status-checked-out uppercase">
       Stock Low
      </span>
      <span class="font-metadata text-metadata text-text-secondary">
       Yesterday
      </span>
     </div>
     <p class="font-body text-body text-text-primary">
      Server Rack Screws inventory dropping below threshold (current: 45, min: 50).
     </p>
    </div>
    <div class="flex-shrink-0 sm:ml-4 mt-4 sm:mt-0 flex gap-2">
     <button class="bg-surface border border-[#D6D3CE] text-text-primary px-3 py-1.5 rounded-lg font-h3 text-[13px] hover:bg-surface-subtle transition-colors">
      Dismiss
     </button>
     <button class="bg-surface border border-border-default text-text-primary px-3 py-1.5 rounded-lg font-h3 text-[13px] hover:bg-surface-subtle transition-colors">
      View Stock
     </button>
    </div>
   </div>
  </div>
 </div>
 <!-- Empty State (Hidden by default in this mockup) -->
 <div class="hidden bg-surface rounded-xl border border-border-default p-12 flex flex-col items-center justify-center text-center mt-stack">
  <div class="w-16 h-16 rounded-full bg-surface-subtle flex items-center justify-center mb-4">
   <span class="material-symbols-outlined text-[32px] text-text-secondary">
    check_circle
   </span>
  </div>
  <h3 class="font-h2 text-h2 text-text-primary mb-2">
   No alerts found
  </h3>
  <p class="font-body text-body text-text-secondary max-w-md">
   You're all caught up! There are currently no critical system alerts or pending notifications.
  </p>
 </div>
</main>

</template>

<script setup lang="ts">
// Autogenerated from alerts_notifications_center
</script>
