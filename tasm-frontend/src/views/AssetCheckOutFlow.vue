<template>
<main class="flex-1 overflow-y-auto px-page-margin py-page-margin">
 <div class="max-w-[1000px] mx-auto w-full">
  <!-- Page Header -->
  <div class="mb-section-gap text-center">
   <h1 class="font-h1 text-h1 text-text-primary mb-base">
    Asset Check-Out
   </h1>
   <p class="font-body text-body text-text-secondary">
    Assign equipment to a staff member.
   </p>
  </div>
  <!-- Stepper -->
  <div class="mb-section-gap">
   <div class="flex items-center justify-between w-full relative">
    <!-- Connecting Line -->
    <div class="absolute left-0 top-1/2 -translate-y-1/2 w-full h-0.5 bg-border-default -z-10">
    </div>
    <div class="absolute left-0 top-1/2 -translate-y-1/2 w-1/4 h-0.5 bg-primary-container -z-10">
    </div>
    <!-- Step 1: Select Asset (Active) -->
    <div class="flex flex-col items-center gap-base bg-canvas px-inline relative">
     <div class="w-8 h-8 rounded-full bg-primary-container text-on-primary flex items-center justify-center font-h3 text-h3 shadow-sm border-2 border-primary-container">
      1
     </div>
     <span class="font-h3 text-h3 text-primary-container">
      Select Asset
     </span>
    </div>
    <!-- Step 2: Select Recipient -->
    <div class="flex flex-col items-center gap-base bg-canvas px-inline relative">
     <div class="w-8 h-8 rounded-full bg-surface text-text-secondary flex items-center justify-center font-h3 text-h3 border-2 border-border-default">
      2
     </div>
     <span class="font-h3 text-h3 text-text-secondary">
      Select Recipient
     </span>
    </div>
    <!-- Step 3: Review -->
    <div class="flex flex-col items-center gap-base bg-canvas px-inline relative">
     <div class="w-8 h-8 rounded-full bg-surface text-text-secondary flex items-center justify-center font-h3 text-h3 border-2 border-border-default">
      3
     </div>
     <span class="font-h3 text-h3 text-text-secondary">
      Review
     </span>
    </div>
    <!-- Step 4: Confirm -->
    <div class="flex flex-col items-center gap-base bg-canvas px-inline relative">
     <div class="w-8 h-8 rounded-full bg-surface text-text-secondary flex items-center justify-center font-h3 text-h3 border-2 border-border-default">
      4
     </div>
     <span class="font-h3 text-h3 text-text-secondary">
      Confirm
     </span>
    </div>
   </div>
  </div>
  <!-- Step Content: Select Asset & Recipient (Hybrid Layout) -->
  <div class="grid grid-cols-1 lg:grid-cols-12 gap-section-gap">
   <!-- Left Column: Asset Selection -->
   <div class="lg:col-span-7 flex flex-col gap-stack">
    <!-- Search and Filter Bar -->
    <div class="flex gap-inline">
     <div class="relative flex-1">
      <span class="material-symbols-outlined absolute left-inline top-1/2 -translate-y-1/2 text-text-secondary" data-icon="search">
       search
      </span>
      <input class="w-full pl-10 pr-inline py-2 bg-surface border border-border-default rounded-DEFAULT font-body text-body text-text-primary focus:outline-none focus:ring-2 focus:ring-primary-container focus:border-transparent shadow-sm" placeholder="Search by asset tag, model, or type..." type="text"/>
     </div>
     <button class="px-inline py-2 bg-surface border border-border-default rounded-DEFAULT flex items-center gap-base text-text-secondary hover:bg-surface-subtle transition-colors shadow-sm">
      <span class="material-symbols-outlined text-[20px]" data-icon="filter_list">
       filter_list
      </span>
      <span class="font-h3 text-h3">
       Filters
      </span>
     </button>
    </div>
    <!-- Asset Grid -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-stack">
     <!-- Selected Asset Card -->
     <div class="bg-surface rounded-[16px] border border-primary-container shadow-md p-stack relative cursor-pointer hover:-translate-y-0.5 transition-transform">
      <div class="absolute top-stack right-stack">
       <span class="material-symbols-outlined text-primary-container" data-icon="check_circle" data-weight="fill" style="font-variation-settings: 'FILL' 1;">
        check_circle
       </span>
      </div>
      <div class="w-12 h-12 bg-surface-subtle rounded-lg flex items-center justify-center mb-stack border border-border-default">
       <span class="material-symbols-outlined text-text-primary text-[24px]" data-icon="laptop_mac">
        laptop_mac
       </span>
      </div>
      <div class="mb-base">
       <span class="inline-flex items-center px-2 py-0.5 rounded-full bg-metric-sage text-status-in-stock font-metadata text-metadata mb-2">
        In Stock
       </span>
      </div>
      <h3 class="font-h2 text-h2 text-text-primary mb-1">
       MacBook Pro 16"
      </h3>
      <p class="font-mono-data text-mono-data text-text-secondary mb-stack">
       TAG: AST-2023-089
      </p>
      <div class="flex gap-2">
       <span class="font-metadata text-metadata text-text-secondary bg-surface-subtle px-2 py-1 rounded border border-border-default">
        M2 Max
       </span>
       <span class="font-metadata text-metadata text-text-secondary bg-surface-subtle px-2 py-1 rounded border border-border-default">
        32GB RAM
       </span>
      </div>
     </div>
     <!-- Available Asset Card -->
     <div class="bg-surface rounded-[16px] border border-border-default shadow-[0_4px_4px_0_rgba(0,0,0,0.02)] p-stack relative cursor-pointer hover:-translate-y-0.5 transition-transform hover:shadow-md">
      <div class="absolute top-stack right-stack">
       <span class="material-symbols-outlined text-text-secondary opacity-0 group-hover:opacity-100 transition-opacity" data-icon="radio_button_unchecked">
        radio_button_unchecked
       </span>
      </div>
      <div class="w-12 h-12 bg-surface-subtle rounded-lg flex items-center justify-center mb-stack border border-border-default">
       <span class="material-symbols-outlined text-text-primary text-[24px]" data-icon="laptop_windows">
        laptop_windows
       </span>
      </div>
      <div class="mb-base">
       <span class="inline-flex items-center px-2 py-0.5 rounded-full bg-metric-sage text-status-in-stock font-metadata text-metadata mb-2">
        In Stock
       </span>
      </div>
      <h3 class="font-h2 text-h2 text-text-primary mb-1">
       Dell XPS 15
      </h3>
      <p class="font-mono-data text-mono-data text-text-secondary mb-stack">
       TAG: AST-2023-142
      </p>
      <div class="flex gap-2">
       <span class="font-metadata text-metadata text-text-secondary bg-surface-subtle px-2 py-1 rounded border border-border-default">
        i9 13th Gen
       </span>
       <span class="font-metadata text-metadata text-text-secondary bg-surface-subtle px-2 py-1 rounded border border-border-default">
        64GB RAM
       </span>
      </div>
     </div>
     <!-- Available Asset Card -->
     <div class="bg-surface rounded-[16px] border border-border-default shadow-[0_4px_4px_0_rgba(0,0,0,0.02)] p-stack relative cursor-pointer hover:-translate-y-0.5 transition-transform hover:shadow-md">
      <div class="absolute top-stack right-stack">
       <span class="material-symbols-outlined text-text-secondary opacity-0 group-hover:opacity-100 transition-opacity" data-icon="radio_button_unchecked">
        radio_button_unchecked
       </span>
      </div>
      <div class="w-12 h-12 bg-surface-subtle rounded-lg flex items-center justify-center mb-stack border border-border-default">
       <span class="material-symbols-outlined text-text-primary text-[24px]" data-icon="tablet_mac">
        tablet_mac
       </span>
      </div>
      <div class="mb-base">
       <span class="inline-flex items-center px-2 py-0.5 rounded-full bg-metric-sage text-status-in-stock font-metadata text-metadata mb-2">
        In Stock
       </span>
      </div>
      <h3 class="font-h2 text-h2 text-text-primary mb-1">
       iPad Pro 12.9"
      </h3>
      <p class="font-mono-data text-mono-data text-text-secondary mb-stack">
       TAG: AST-2023-205
      </p>
      <div class="flex gap-2">
       <span class="font-metadata text-metadata text-text-secondary bg-surface-subtle px-2 py-1 rounded border border-border-default">
        M2
       </span>
       <span class="font-metadata text-metadata text-text-secondary bg-surface-subtle px-2 py-1 rounded border border-border-default">
        512GB
       </span>
      </div>
     </div>
    </div>
   </div>
   <!-- Right Column: Contextual Details & Recipient -->
   <div class="lg:col-span-5 flex flex-col gap-stack">
    <!-- Linked Accessories (Appears when asset is selected) -->
    <div class="bg-surface rounded-[16px] border border-border-default shadow-[0_4px_4px_0_rgba(0,0,0,0.02)] p-card-padding">
     <h2 class="font-h2 text-h2 text-text-primary mb-stack flex items-center gap-2">
      <span class="material-symbols-outlined text-text-secondary text-[20px]" data-icon="cable">
       cable
      </span>
      Linked Accessories
     </h2>
     <p class="font-body text-body text-text-secondary mb-stack">
      Select accessories to include with AST-2023-089.
     </p>
     <div class="flex flex-col gap-2">
      <label class="flex items-center p-3 rounded-DEFAULT border border-primary-container bg-metric-amber/10 cursor-pointer">
       <input checked class="w-4 h-4 text-primary-container rounded border-primary-container focus:ring-primary-container" type="checkbox"/>
       <div class="ml-3 flex flex-col">
        <span class="font-h3 text-h3 text-text-primary">
         140W USB-C Power Adapter
        </span>
        <span class="font-mono-data text-mono-data text-text-secondary">
         ACC-PWR-042
        </span>
       </div>
      </label>
      <label class="flex items-center p-3 rounded-DEFAULT border border-border-default bg-surface hover:bg-surface-subtle cursor-pointer transition-colors">
       <input class="w-4 h-4 text-primary-container rounded border-outline focus:ring-primary-container" type="checkbox"/>
       <div class="ml-3 flex flex-col">
        <span class="font-h3 text-h3 text-text-primary">
         USB-C to MagSafe 3 Cable (2 m)
        </span>
        <span class="font-mono-data text-mono-data text-text-secondary">
         ACC-CBL-118
        </span>
       </div>
      </label>
      <label class="flex items-center p-3 rounded-DEFAULT border border-primary-container bg-metric-amber/10 cursor-pointer">
       <input checked class="w-4 h-4 text-primary-container rounded border-primary-container focus:ring-primary-container" type="checkbox"/>
       <div class="ml-3 flex flex-col">
        <span class="font-h3 text-h3 text-text-primary">
         Incase 16" Laptop Sleeve
        </span>
        <span class="font-mono-data text-mono-data text-text-secondary">
         ACC-BAG-092
        </span>
       </div>
      </label>
     </div>
    </div>
    <!-- Recipient Selection -->
    <div class="bg-surface rounded-[16px] border border-border-default shadow-[0_4px_4px_0_rgba(0,0,0,0.02)] p-card-padding">
     <h2 class="font-h2 text-h2 text-text-primary mb-stack flex items-center gap-2">
      <span class="material-symbols-outlined text-text-secondary text-[20px]" data-icon="person_add">
       person_add
      </span>
      Select Recipient
     </h2>
     <!-- Recipient Search -->
     <div class="relative mb-stack">
      <span class="material-symbols-outlined absolute left-inline top-1/2 -translate-y-1/2 text-text-secondary" data-icon="search">
       search
      </span>
      <input class="w-full pl-10 pr-inline py-2 bg-surface border border-primary-container rounded-DEFAULT font-body text-body text-text-primary focus:outline-none focus:ring-2 focus:ring-primary-container focus:border-transparent shadow-sm" type="text" value="Sarah J"/>
     </div>
     <!-- Selected Recipient Card -->
     <div class="flex items-center p-inline rounded-DEFAULT border border-primary-container bg-metric-amber/10">
      <img alt="Profile picture of Sarah Jenkins, Engineering Dept" class="w-12 h-12 rounded-full border border-border-default mr-inline object-cover" data-alt="professional headshot of a young woman with dark hair in a corporate setting, soft lighting" src="https://lh3.googleusercontent.com/aida-public/AB6AXuDwtakJ2rCE2OVQ-BAkajtwvc7Rno5ZInFFJH58qy4-AEKBCxRg3507BTIBGEfCpOSZibvnPjmSYRPdrhL273OL1hRZoQEUw80jeqJyTOTbbSlEMAD6N7ChR5WWoWZSnIdR4Rnt3kKHgHVKlm-bGirQha5ngzS9oeOZbOO1mDpAgyN-wsK9_jZuFoMe9M1v1AhdrkIiKtz9fv76JFk6SRhwvfN6Ko-Qq21peSXtXpqoAYiMV9Zr6pXTTSKAB0skATHNP16a4BVkvbbK"/>
      <div class="flex-1">
       <h3 class="font-h3 text-h3 text-text-primary">
        Sarah Jenkins
       </h3>
       <p class="font-metadata text-metadata text-text-secondary">
        Engineering • Frontend Developer
       </p>
      </div>
      <div class="flex flex-col items-end">
       <span class="inline-flex items-center gap-1 font-metadata text-metadata text-text-secondary">
        <span class="material-symbols-outlined text-[14px]" data-icon="inventory_2">
         inventory_2
        </span>
        Has 2 assets
       </span>
      </div>
     </div>
    </div>
   </div>
  </div>
  <!-- Footer Actions -->
  <div class="mt-section-gap flex justify-end gap-inline border-t border-border-default pt-stack">
   <button class="px-6 py-2 bg-surface border border-outline text-text-primary font-h3 text-h3 rounded-DEFAULT hover:bg-surface-subtle transition-colors shadow-sm">
    Back
   </button>
   <button class="px-6 py-2 bg-text-primary text-surface font-h3 text-h3 rounded-DEFAULT hover:bg-inverse-surface transition-colors shadow-sm flex items-center gap-2">
    Continue to Review
    <span class="material-symbols-outlined text-[18px]" data-icon="arrow_forward">
     arrow_forward
    </span>
   </button>
  </div>
 </div>
</main>

</template>

<script setup lang="ts">
// Autogenerated from asset_check_out_flow
</script>
