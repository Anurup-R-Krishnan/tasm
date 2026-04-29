<template>
<main class="flex-1 md:ml-[248px] p-page-margin w-full max-w-[1400px] mx-auto">
 <div class="mb-section-gap">
  <h1 class="font-h1 text-h1 text-text-primary mb-2">
   Welcome back, Sarah
  </h1>
  <p class="text-text-secondary font-body text-body">
   Here is an overview of your assigned assets and pending tasks.
  </p>
 </div>
 <div class="grid grid-cols-1 lg:grid-cols-3 gap-section-gap">
  <!-- Left Column: Pending & Quick Requests -->
  <div class="lg:col-span-1 flex flex-col gap-section-gap">
   <!-- Pending Acknowledgements -->
   <section>
    <h2 class="font-h2 text-h2 text-text-primary mb-stack flex items-center gap-2">
     <span class="material-symbols-outlined text-amber-700" data-icon="pending_actions">
      pending_actions
     </span>
     Pending Acknowledgements
    </h2>
    <div class="bg-surface border border-border-default rounded-xl shadow-sm overflow-hidden p-card-padding">
     <div class="flex flex-col gap-4">
      <div class="border border-amber-200 bg-amber-50 rounded-lg p-4 relative overflow-hidden">
       <div class="absolute top-0 left-0 w-1 h-full bg-amber-500">
       </div>
       <div class="flex justify-between items-start mb-2">
        <h3 class="font-h3 text-h3 text-text-primary">
         MacBook Pro 16"
        </h3>
        <span class="bg-metric-amber text-amber-900 font-metadata text-metadata px-2 py-1 rounded-full flex items-center gap-1">
         <span class="sr-only">
          Status:
         </span>
         Awaiting Sign-off
        </span>
       </div>
       <p class="font-mono-data text-mono-data text-text-secondary mb-4">
        AST-2023-0892
       </p>
       <p class="font-body text-body text-text-primary mb-4 text-sm">
        Assigned on Oct 12, 2023. Please acknowledge receipt of this device.
       </p>
       <button class="w-full bg-primary text-on-primary font-medium py-2 rounded-lg hover:translate-y-[-2px] transition-transform shadow-sm">
        Sign-off Document
       </button>
      </div>
     </div>
    </div>
   </section>
   <!-- Quick Requests -->
   <section>
    <h2 class="font-h2 text-h2 text-text-primary mb-stack flex items-center gap-2">
     <span class="material-symbols-outlined text-stone-600" data-icon="bolt">
      bolt
     </span>
     Quick Requests
    </h2>
    <div class="bg-surface border border-border-default rounded-xl shadow-sm p-card-padding">
     <div class="grid grid-cols-2 gap-inline">
      <button class="flex flex-col items-center justify-center gap-2 p-4 rounded-lg border border-border-default hover:bg-stone-50 hover:border-amber-300 transition-colors group">
       <div class="w-10 h-10 rounded-full bg-amber-50 text-amber-700 flex items-center justify-center group-hover:bg-amber-100 transition-colors">
        <span class="material-symbols-outlined" data-icon="laptop_mac">
         laptop_mac
        </span>
       </div>
       <span class="font-body text-body text-sm font-medium text-center">
        New Laptop
       </span>
      </button>
      <button class="flex flex-col items-center justify-center gap-2 p-4 rounded-lg border border-border-default hover:bg-stone-50 hover:border-amber-300 transition-colors group">
       <div class="w-10 h-10 rounded-full bg-stone-100 text-stone-700 flex items-center justify-center group-hover:bg-stone-200 transition-colors">
        <span class="material-symbols-outlined" data-icon="mouse">
         mouse
        </span>
       </div>
       <span class="font-body text-body text-sm font-medium text-center">
        Accessories
       </span>
      </button>
      <button class="flex flex-col items-center justify-center gap-2 p-4 rounded-lg border border-border-default hover:bg-stone-50 hover:border-amber-300 transition-colors group col-span-2">
       <div class="w-10 h-10 rounded-full bg-red-50 text-red-700 flex items-center justify-center group-hover:bg-red-100 transition-colors">
        <span class="material-symbols-outlined" data-icon="build">
         build
        </span>
       </div>
       <span class="font-body text-body text-sm font-medium text-center">
        Report Issue / Repair
       </span>
      </button>
     </div>
    </div>
   </section>
  </div>
  <!-- Right Column: Active Assets -->
  <div class="lg:col-span-2 flex flex-col gap-section-gap">
   <section class="h-full flex flex-col">
    <div class="flex items-center justify-between mb-stack">
     <h2 class="font-h2 text-h2 text-text-primary flex items-center gap-2">
      <span class="material-symbols-outlined text-stone-600" data-icon="devices">
       devices
      </span>
      My Active Assets
     </h2>
     <button class="text-amber-700 font-medium text-sm hover:underline">
      View All
     </button>
    </div>
    <div class="flex-1 bg-surface border border-border-default rounded-xl shadow-sm overflow-hidden">
     <div class="p-card-padding grid grid-cols-1 md:grid-cols-2 gap-stack">
      <!-- Asset Card 1 -->
      <div class="border border-border-default rounded-lg p-4 hover:translate-y-[-2px] transition-transform bg-surface-subtle shadow-sm flex flex-col h-full relative group">
       <div class="flex justify-between items-start mb-4">
        <div class="flex items-center gap-3">
         <div class="w-12 h-12 rounded-lg bg-stone-100 flex items-center justify-center text-stone-600 border border-stone-200">
          <span class="material-symbols-outlined text-[24px]" data-icon="monitor">
           monitor
          </span>
         </div>
         <div>
          <h3 class="font-h3 text-h3 text-text-primary">
           Dell UltraSharp 27"
          </h3>
          <p class="font-mono-data text-mono-data text-text-secondary mt-1">
           DSP-2021-0045
          </p>
         </div>
        </div>
       </div>
       <div class="mb-4">
        <span class="bg-metric-sage text-green-800 font-metadata text-metadata px-2 py-1 rounded-full inline-flex items-center gap-1">
         <span class="w-1.5 h-1.5 rounded-full bg-green-600">
         </span>
         <span class="sr-only">
          Status:
         </span>
         Active
        </span>
       </div>
       <div class="mt-auto pt-4 border-t border-border-default flex gap-2">
        <button class="flex-1 bg-white border border-border-default text-text-primary font-medium py-1.5 px-3 rounded text-sm hover:bg-stone-50 transition-colors">
         Details
        </button>
        <button class="flex-1 bg-white border border-border-default text-text-primary font-medium py-1.5 px-3 rounded text-sm hover:bg-red-50 hover:text-red-700 hover:border-red-200 transition-colors flex items-center justify-center gap-1">
         Report Issue
        </button>
       </div>
      </div>
      <!-- Asset Card 2 -->
      <div class="border border-border-default rounded-lg p-4 hover:translate-y-[-2px] transition-transform bg-surface-subtle shadow-sm flex flex-col h-full relative group">
       <div class="flex justify-between items-start mb-4">
        <div class="flex items-center gap-3">
         <div class="w-12 h-12 rounded-lg bg-stone-100 flex items-center justify-center text-stone-600 border border-stone-200">
          <span class="material-symbols-outlined text-[24px]" data-icon="keyboard">
           keyboard
          </span>
         </div>
         <div>
          <h3 class="font-h3 text-h3 text-text-primary">
           Logitech MX Keys
          </h3>
          <p class="font-mono-data text-mono-data text-text-secondary mt-1">
           ACC-2022-1102
          </p>
         </div>
        </div>
       </div>
       <div class="mb-4">
        <span class="bg-metric-sage text-green-800 font-metadata text-metadata px-2 py-1 rounded-full inline-flex items-center gap-1">
         <span class="w-1.5 h-1.5 rounded-full bg-green-600">
         </span>
         <span class="sr-only">
          Status:
         </span>
         Active
        </span>
       </div>
       <div class="mt-auto pt-4 border-t border-border-default flex gap-2">
        <button class="flex-1 bg-white border border-border-default text-text-primary font-medium py-1.5 px-3 rounded text-sm hover:bg-stone-50 transition-colors">
         Details
        </button>
        <button class="flex-1 bg-white border border-border-default text-text-primary font-medium py-1.5 px-3 rounded text-sm hover:bg-red-50 hover:text-red-700 hover:border-red-200 transition-colors flex items-center justify-center gap-1">
         Report Issue
        </button>
       </div>
      </div>
      <!-- Asset Card 3 -->
      <div class="border border-border-default rounded-lg p-4 hover:translate-y-[-2px] transition-transform bg-surface-subtle shadow-sm flex flex-col h-full relative group">
       <div class="flex justify-between items-start mb-4">
        <div class="flex items-center gap-3">
         <div class="w-12 h-12 rounded-lg bg-stone-100 flex items-center justify-center text-stone-600 border border-stone-200">
          <span class="material-symbols-outlined text-[24px]" data-icon="phone_iphone">
           phone_iphone
          </span>
         </div>
         <div>
          <h3 class="font-h3 text-h3 text-text-primary">
           iPhone 13 Pro
          </h3>
          <p class="font-mono-data text-mono-data text-text-secondary mt-1">
           MOB-2022-0341
          </p>
         </div>
        </div>
       </div>
       <div class="mb-4">
        <span class="bg-red-50 text-red-800 font-metadata text-metadata px-2 py-1 rounded-full inline-flex items-center gap-1 border border-red-100">
         <span class="material-symbols-outlined text-[14px]" data-icon="build">
          build
         </span>
         <span class="sr-only">
          Status:
         </span>
         Repair Pending
        </span>
       </div>
       <div class="mt-auto pt-4 border-t border-border-default flex gap-2">
        <button class="w-full bg-white border border-border-default text-text-primary font-medium py-1.5 px-3 rounded text-sm hover:bg-stone-50 transition-colors">
         View Ticket Details
        </button>
       </div>
      </div>
     </div>
    </div>
   </section>
  </div>
 </div>
</main>

</template>

<script setup lang="ts">
// Autogenerated from employee_self_service_portal_3
</script>
