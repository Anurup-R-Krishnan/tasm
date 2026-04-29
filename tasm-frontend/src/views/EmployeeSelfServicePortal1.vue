<template>
<main class="flex-1 mt-[60px] h-[calc(100vh-60px)] overflow-y-auto p-page-margin relative z-0">
 <div class="max-w-[1400px] mx-auto flex flex-col gap-section-gap">
  <!-- Page Header & Primary Action -->
  <header class="flex flex-col md:flex-row md:items-end justify-between gap-stack border-b border-border-default pb-stack">
   <div>
    <h1 class="font-h1 text-h1 text-text-primary mb-1">
     Employee Self-Service
    </h1>
    <p class="font-body text-body text-text-secondary">
     Manage your assigned assets and submit new requests.
    </p>
   </div>
   <button class="bg-surface-tint text-on-primary hover:bg-primary transition-colors px-6 py-3 rounded-lg shadow-[0_4px_12px_rgba(154,70,20,0.2)] flex items-center gap-2 active:scale-95 duration-150 group">
    <span class="material-symbols-outlined group-hover:rotate-90 transition-transform duration-300" data-icon="add_circle" data-weight="fill">
     add_circle
    </span>
    <span class="font-h3 text-h3">
     Request New Asset
    </span>
   </button>
  </header>
  <!-- My Assets Section (Bento Grid Style) -->
  <section>
   <h2 class="font-h2 text-h2 text-text-primary mb-stack flex items-center gap-2">
    <span class="material-symbols-outlined text-text-secondary" data-icon="devices">
     devices
    </span>
    My Assigned Assets
   </h2>
   <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-stack">
    <!-- Asset Card 1: Laptop -->
    <div class="bg-surface rounded-xl border border-border-default p-card-padding shadow-[0_2px_8px_-2px_rgba(0,0,0,0.05)] hover:-translate-y-1 transition-transform duration-200 flex flex-col h-full relative overflow-hidden group">
     <div class="absolute top-0 right-0 w-24 h-24 bg-metric-lavender/30 rounded-bl-[100px] -z-10 group-hover:scale-110 transition-transform">
     </div>
     <div class="flex justify-between items-start mb-6">
      <div class="w-12 h-12 bg-metric-lavender rounded-lg flex items-center justify-center text-tertiary-container shadow-sm border border-white/50">
       <span class="material-symbols-outlined" data-icon="laptop_mac">
        laptop_mac
       </span>
      </div>
      <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-metadata font-metadata bg-surface-container text-status-checked-out border border-[#dbeafe]">
       Checked Out
      </span>
     </div>
     <div class="mb-6 flex-1">
      <h3 class="font-h3 text-h3 text-text-primary mb-1">
       MacBook Pro 14"
      </h3>
      <p class="font-metadata text-metadata text-text-secondary mb-3">
       M2 Pro, 16GB RAM, 512GB SSD
      </p>
      <div class="space-y-2 bg-surface-subtle p-3 rounded-lg border border-border-default/50">
       <div class="flex justify-between items-center">
        <span class="font-metadata text-metadata text-text-secondary">
         Asset Tag
        </span>
        <span class="font-mono-data text-mono-data text-text-primary">
         AST-23-0892
        </span>
       </div>
       <div class="flex justify-between items-center">
        <span class="font-metadata text-metadata text-text-secondary">
         Assigned
        </span>
        <span class="font-metadata text-metadata text-text-primary">
         Jan 12, 2024
        </span>
       </div>
      </div>
     </div>
     <button class="w-full py-2.5 bg-surface border border-outline-variant rounded-lg font-h3 text-h3 text-text-primary hover:bg-surface-subtle hover:border-outline transition-colors flex items-center justify-center gap-2">
      <span class="material-symbols-outlined text-text-secondary" data-icon="report" style="font-size: 18px;">
       report
      </span>
      Report Issue
     </button>
    </div>
    <!-- Asset Card 2: Phone -->
    <div class="bg-surface rounded-xl border border-border-default p-card-padding shadow-[0_2px_8px_-2px_rgba(0,0,0,0.05)] hover:-translate-y-1 transition-transform duration-200 flex flex-col h-full relative overflow-hidden group">
     <div class="absolute top-0 right-0 w-24 h-24 bg-metric-amber/30 rounded-bl-[100px] -z-10 group-hover:scale-110 transition-transform">
     </div>
     <div class="flex justify-between items-start mb-6">
      <div class="w-12 h-12 bg-metric-amber rounded-lg flex items-center justify-center text-primary-container shadow-sm border border-white/50">
       <span class="material-symbols-outlined" data-icon="smartphone">
        smartphone
       </span>
      </div>
      <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-metadata font-metadata bg-surface-container text-status-checked-out border border-[#dbeafe]">
       Checked Out
      </span>
     </div>
     <div class="mb-6 flex-1">
      <h3 class="font-h3 text-h3 text-text-primary mb-1">
       iPhone 14 Pro
      </h3>
      <p class="font-metadata text-metadata text-text-secondary mb-3">
       Space Black, 256GB - Corporate Plan
      </p>
      <div class="space-y-2 bg-surface-subtle p-3 rounded-lg border border-border-default/50">
       <div class="flex justify-between items-center">
        <span class="font-metadata text-metadata text-text-secondary">
         Asset Tag
        </span>
        <span class="font-mono-data text-mono-data text-text-primary">
         MOB-22-1104
        </span>
       </div>
       <div class="flex justify-between items-center">
        <span class="font-metadata text-metadata text-text-secondary">
         Assigned
        </span>
        <span class="font-metadata text-metadata text-text-primary">
         Nov 05, 2023
        </span>
       </div>
      </div>
     </div>
     <button class="w-full py-2.5 bg-surface border border-outline-variant rounded-lg font-h3 text-h3 text-text-primary hover:bg-surface-subtle hover:border-outline transition-colors flex items-center justify-center gap-2">
      <span class="material-symbols-outlined text-text-secondary" data-icon="report" style="font-size: 18px;">
       report
      </span>
      Report Issue
     </button>
    </div>
    <!-- Asset Card 3: ID Card -->
    <div class="bg-surface rounded-xl border border-border-default p-card-padding shadow-[0_2px_8px_-2px_rgba(0,0,0,0.05)] hover:-translate-y-1 transition-transform duration-200 flex flex-col h-full relative overflow-hidden group">
     <div class="absolute top-0 right-0 w-24 h-24 bg-metric-sage/30 rounded-bl-[100px] -z-10 group-hover:scale-110 transition-transform">
     </div>
     <div class="flex justify-between items-start mb-6">
      <div class="w-12 h-12 bg-metric-sage rounded-lg flex items-center justify-center text-status-in-stock shadow-sm border border-white/50">
       <span class="material-symbols-outlined" data-icon="badge">
        badge
       </span>
      </div>
      <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-metadata font-metadata bg-surface-container text-status-checked-out border border-[#dbeafe]">
       Active
      </span>
     </div>
     <div class="mb-6 flex-1">
      <h3 class="font-h3 text-h3 text-text-primary mb-1">
       Corporate Access Card
      </h3>
      <p class="font-metadata text-metadata text-text-secondary mb-3">
       Building A &amp; B, Data Center Level 1
      </p>
      <div class="space-y-2 bg-surface-subtle p-3 rounded-lg border border-border-default/50">
       <div class="flex justify-between items-center">
        <span class="font-metadata text-metadata text-text-secondary">
         Card ID
        </span>
        <span class="font-mono-data text-mono-data text-text-primary">
         ACC-99-402
        </span>
       </div>
       <div class="flex justify-between items-center">
        <span class="font-metadata text-metadata text-text-secondary">
         Expires
        </span>
        <span class="font-metadata text-metadata text-text-primary">
         Dec 31, 2025
        </span>
       </div>
      </div>
     </div>
     <button class="w-full py-2.5 bg-surface border border-outline-variant rounded-lg font-h3 text-h3 text-text-primary hover:bg-surface-subtle hover:border-outline transition-colors flex items-center justify-center gap-2">
      <span class="material-symbols-outlined text-text-secondary" data-icon="report" style="font-size: 18px;">
       report
      </span>
      Report Lost
     </button>
    </div>
   </div>
  </section>
  <!-- My Requests Timeline Section -->
  <section class="mb-section-gap">
   <h2 class="font-h2 text-h2 text-text-primary mb-stack flex items-center gap-2">
    <span class="material-symbols-outlined text-text-secondary" data-icon="history">
     history
    </span>
    Recent Requests
   </h2>
   <div class="bg-surface border border-border-default rounded-xl p-card-padding shadow-[0_4px_16px_-4px_rgba(0,0,0,0.03)]">
    <div class="relative border-l-2 border-border-default ml-4 space-y-8 pb-4">
     <!-- Timeline Item 1 -->
     <div class="relative pl-8">
      <div class="absolute -left-[17px] top-1 w-8 h-8 bg-surface border-2 border-primary-container rounded-full flex items-center justify-center text-primary-container shadow-sm">
       <span class="material-symbols-outlined" data-icon="hourglass_empty" style="font-size: 16px;">
        hourglass_empty
       </span>
      </div>
      <div class="bg-surface-subtle border border-border-default p-4 rounded-lg">
       <div class="flex justify-between items-start mb-2">
        <h4 class="font-h3 text-h3 text-text-primary">
         Ergonomic Keyboard &amp; Mouse
        </h4>
        <span class="font-metadata text-metadata text-text-secondary">
         Today, 09:41 AM
        </span>
       </div>
       <p class="font-body text-body text-text-secondary mb-2">
        REQ-2024-0512 • Peripheral Hardware
       </p>
       <div class="flex items-center gap-2">
        <span class="inline-flex w-2 h-2 rounded-full bg-amber-500">
        </span>
        <span class="font-metadata text-metadata text-text-primary font-medium">
         Pending Approval
        </span>
        <span class="font-metadata text-metadata text-text-secondary ml-2 border-l border-border-default pl-2">
         Awaiting Manager Sign-off
        </span>
       </div>
      </div>
     </div>
     <!-- Timeline Item 2 -->
     <div class="relative pl-8">
      <div class="absolute -left-[17px] top-1 w-8 h-8 bg-surface border-2 border-tertiary rounded-full flex items-center justify-center text-tertiary shadow-sm">
       <span class="material-symbols-outlined" data-icon="sync" style="font-size: 16px;">
        sync
       </span>
      </div>
      <div class="bg-surface-subtle border border-border-default p-4 rounded-lg">
       <div class="flex justify-between items-start mb-2">
        <h4 class="font-h3 text-h3 text-text-primary">
         Dual Monitor Stand
        </h4>
        <span class="font-metadata text-metadata text-text-secondary">
         Oct 15, 2024
        </span>
       </div>
       <p class="font-body text-body text-text-secondary mb-2">
        REQ-2024-0498 • Office Equipment
       </p>
       <div class="flex items-center gap-2">
        <span class="inline-flex w-2 h-2 rounded-full bg-tertiary">
        </span>
        <span class="font-metadata text-metadata text-text-primary font-medium">
         Processing
        </span>
        <span class="font-metadata text-metadata text-text-secondary ml-2 border-l border-border-default pl-2">
         IT Procurement Team
        </span>
       </div>
      </div>
     </div>
     <!-- Timeline Item 3 -->
     <div class="relative pl-8">
      <div class="absolute -left-[17px] top-1 w-8 h-8 bg-surface border-2 border-status-in-stock rounded-full flex items-center justify-center text-status-in-stock shadow-sm">
       <span class="material-symbols-outlined" data-icon="check" style="font-size: 16px;">
        check
       </span>
      </div>
      <div class="bg-metric-sage/20 border border-[#bbf7d0] p-4 rounded-lg">
       <div class="flex justify-between items-start mb-2">
        <h4 class="font-h3 text-h3 text-text-primary">
         Privacy Screen Filter 14"
        </h4>
        <span class="font-metadata text-metadata text-text-secondary">
         Oct 12, 2024
        </span>
       </div>
       <p class="font-body text-body text-text-secondary mb-2">
        REQ-2024-0455 • Accessories
       </p>
       <div class="flex items-center gap-2">
        <span class="inline-flex w-2 h-2 rounded-full bg-status-in-stock">
        </span>
        <span class="font-metadata text-metadata text-status-in-stock font-medium">
         Ready for Pickup
        </span>
        <span class="font-metadata text-metadata text-text-secondary ml-2 border-l border-border-default pl-2">
         IT Helpdesk, Ground Floor
        </span>
       </div>
      </div>
     </div>
    </div>
    <div class="mt-4 text-center">
     <button class="font-h3 text-h3 text-primary hover:text-surface-tint hover:underline transition-all">
      View All Request History
     </button>
    </div>
   </div>
  </section>
 </div>
</main>

</template>

<script setup lang="ts">
// Autogenerated from employee_self_service_portal_1
</script>
