<template>
<main class="md:ml-[248px] min-h-screen flex flex-col">
 <!-- Metrics Header -->
 <div class="p-page-margin border-b border-border-default bg-surface/50">
  <div class="flex justify-between items-end mb-stack">
   <div>
    <h1 class="font-h1 text-h1 text-text-primary mb-1">
     Asset Registry
    </h1>
    <p class="font-body text-body text-text-secondary">
     Manage and track organization-wide assets across all Technopark phases.
    </p>
   </div>
   <div class="flex gap-inline">
    <button class="bg-surface border border-[#D6D3CE] text-text-primary px-4 py-2 rounded-lg font-body text-body font-semibold hover:bg-surface-subtle transition-colors flex items-center gap-2 shadow-sm">
     <span class="material-symbols-outlined text-[18px]">
      download
     </span>
     Export
    </button>
    <button class="bg-surface border border-[#D6D3CE] text-text-primary px-4 py-2 rounded-lg font-body text-body font-semibold hover:bg-surface-subtle transition-colors flex items-center gap-2 shadow-sm">
     <span class="material-symbols-outlined text-[18px]">
      filter_list
     </span>
     Filter
    </button>
   </div>
  </div>
  <div class="grid grid-cols-1 md:grid-cols-3 gap-stack">
   <!-- Metric 1 -->
   <div class="bg-metric-amber rounded-[16px] p-card-padding cursor-pointer hover:-translate-y-[2px] transition-transform shadow-sm border border-border-default/50">
    <div class="flex justify-between items-start mb-2">
     <p class="font-table-header text-table-header text-on-surface-variant uppercase">
      Total Asset Value
     </p>
     <span class="material-symbols-outlined text-primary-container">
      account_balance_wallet
     </span>
    </div>
    <p class="font-kpi-number text-kpi-number text-text-primary mb-1">
     ₹ 24.5 Cr
    </p>
    <p class="font-metadata text-metadata text-text-secondary flex items-center gap-1">
     <span class="material-symbols-outlined text-[14px] text-status-in-stock">
      trending_up
     </span>
     +2.4% vs last quarter
    </p>
   </div>
   <!-- Metric 2 -->
   <div class="bg-metric-lavender rounded-[16px] p-card-padding cursor-pointer hover:-translate-y-[2px] transition-transform shadow-sm border border-border-default/50">
    <div class="flex justify-between items-start mb-2">
     <p class="font-table-header text-table-header text-on-surface-variant uppercase">
      Active Assets
     </p>
     <span class="material-symbols-outlined text-tertiary-container">
      devices
     </span>
    </div>
    <p class="font-kpi-number text-kpi-number text-text-primary mb-1">
     1,248
    </p>
    <p class="font-metadata text-metadata text-text-secondary flex items-center gap-1">
     Across 3 campuses
    </p>
   </div>
   <!-- Metric 3 -->
   <div class="bg-metric-sage rounded-[16px] p-card-padding cursor-pointer hover:-translate-y-[2px] transition-transform shadow-sm border border-border-default/50">
    <div class="flex justify-between items-start mb-2">
     <p class="font-table-header text-table-header text-on-surface-variant uppercase">
      Pending Maintenance
     </p>
     <span class="material-symbols-outlined text-status-in-stock">
      build_circle
     </span>
    </div>
    <p class="font-kpi-number text-kpi-number text-text-primary mb-1">
     32
    </p>
    <p class="font-metadata text-metadata text-text-secondary flex items-center gap-1">
     <span class="material-symbols-outlined text-[14px] text-status-critical">
      warning
     </span>
     12 critical items
    </p>
   </div>
  </div>
 </div>
 <!-- Hybrid Content Area -->
 <div class="flex flex-1 overflow-hidden">
  <!-- Context Sidebar (Categories/Locations) -->
  <aside class="w-[240px] border-r border-border-default bg-surface/30 overflow-y-auto hidden lg:block">
   <div class="p-4">
    <h3 class="font-h3 text-h3 text-text-primary mb-4">
     Categories
    </h3>
    <ul class="space-y-1">
     <li>
      <button class="w-full flex items-center justify-between text-left px-3 py-2 rounded-md bg-accent-subtle text-primary-container font-medium font-body text-body">
       <div class="flex items-center gap-2">
        <span class="material-symbols-outlined text-[18px]">
         computer
        </span>
        IT Equipment
       </div>
       <span class="bg-surface text-primary-container px-2 py-0.5 rounded-full text-[10px] font-bold">
        842
       </span>
      </button>
     </li>
     <li>
      <button class="w-full flex items-center justify-between text-left px-3 py-2 rounded-md hover:bg-surface-subtle text-text-secondary font-body text-body transition-colors">
       <div class="flex items-center gap-2">
        <span class="material-symbols-outlined text-[18px]">
         chair
        </span>
        Furniture
       </div>
       <span class="px-2 py-0.5 rounded-full text-[10px] font-medium bg-surface-subtle">
        315
       </span>
      </button>
     </li>
     <li>
      <button class="w-full flex items-center justify-between text-left px-3 py-2 rounded-md hover:bg-surface-subtle text-text-secondary font-body text-body transition-colors">
       <div class="flex items-center gap-2">
        <span class="material-symbols-outlined text-[18px]">
         directions_car
        </span>
        Vehicles
       </div>
       <span class="px-2 py-0.5 rounded-full text-[10px] font-medium bg-surface-subtle">
        42
       </span>
      </button>
     </li>
     <li>
      <button class="w-full flex items-center justify-between text-left px-3 py-2 rounded-md hover:bg-surface-subtle text-text-secondary font-body text-body transition-colors">
       <div class="flex items-center gap-2">
        <span class="material-symbols-outlined text-[18px]">
         ac_unit
        </span>
        HVAC &amp; Plant
       </div>
       <span class="px-2 py-0.5 rounded-full text-[10px] font-medium bg-surface-subtle">
        49
       </span>
      </button>
     </li>
    </ul>
    <hr class="my-6 border-border-default"/>
    <h3 class="font-h3 text-h3 text-text-primary mb-4">
     Locations
    </h3>
    <ul class="space-y-1">
     <li>
      <button class="w-full flex items-center justify-between text-left px-3 py-2 rounded-md hover:bg-surface-subtle text-text-secondary font-body text-body transition-colors">
       <div class="flex items-center gap-2">
        <span class="material-symbols-outlined text-[18px]">
         location_on
        </span>
        Phase I (Main)
       </div>
      </button>
     </li>
     <li>
      <button class="w-full flex items-center justify-between text-left px-3 py-2 rounded-md hover:bg-surface-subtle text-text-secondary font-body text-body transition-colors">
       <div class="flex items-center gap-2">
        <span class="material-symbols-outlined text-[18px]">
         location_on
        </span>
        Phase II (Bhavani)
       </div>
      </button>
     </li>
     <li>
      <button class="w-full flex items-center justify-between text-left px-3 py-2 rounded-md hover:bg-surface-subtle text-text-secondary font-body text-body transition-colors">
       <div class="flex items-center gap-2">
        <span class="material-symbols-outlined text-[18px]">
         location_on
        </span>
        Phase III (Ganga)
       </div>
      </button>
     </li>
    </ul>
   </div>
  </aside>
  <!-- Main Grid Canvas -->
  <div class="flex-1 p-page-margin overflow-y-auto">
   <div class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-stack">
    <!-- Asset Card 1 -->
    <div class="bg-surface rounded-[16px] border border-border-default shadow-[0_4px_12px_rgba(0,0,0,0.03)] overflow-hidden hover:-translate-y-[2px] transition-all duration-200 group flex flex-col">
     <div class="h-32 bg-surface-subtle relative">
      <img alt="MacBook Pro" class="w-full h-full object-cover opacity-90 group-hover:opacity-100 transition-opacity" data-alt="Sleek silver laptop resting on a clean wooden desk with soft natural lighting from a window" src="https://lh3.googleusercontent.com/aida-public/AB6AXuB_LVanr8UN6wYE_Cvw3x0hT9OTLHKVpRBmnIRTXjiRAMJgZuhYYszrWQY6rggcAMV1L7CRGfIKelLKSEYyf9gMY2B1wMJpqIDxKnqNfIseowrJUd4BGkobu-MYuTlN9Ja3WMefg3Jw_GbDl4zjbSHSjEsovtLrXIOWMKHOFvxMCjs3gSNcW24aG_gBDyAVQaqTKuOjM6Pin8NzYcuHZC3mrWIqgYzqjgjcOeWI8cJd9LzrnwGHBSll5k5c2WeVJs3PVa_iIWCyszjE"/>
      <div class="absolute top-3 left-3 bg-white/90 backdrop-blur-sm px-2 py-1 rounded shadow-sm border border-border-default/50 flex items-center gap-1">
       <span class="w-2 h-2 rounded-full bg-status-in-stock">
       </span>
       <span class="font-table-header text-table-header text-text-primary">
        In Use
       </span>
      </div>
     </div>
     <div class="p-5 flex-1 flex flex-col">
      <div class="flex justify-between items-start mb-2">
       <span class="font-mono-data text-mono-data text-primary-container bg-accent-subtle px-2 py-1 rounded">
        TPK-IT-2023-084
       </span>
      </div>
      <h3 class="font-h2 text-h2 text-text-primary mb-1">
       MacBook Pro 16" M2 Max
      </h3>
      <p class="font-metadata text-metadata text-text-secondary mb-4 line-clamp-1">
       Assigned to: Rahul N. (Engineering)
      </p>
      <div class="mt-auto grid grid-cols-2 gap-2 mb-4 pt-4 border-t border-surface-subtle">
       <div>
        <p class="font-table-header text-table-header text-on-surface-variant uppercase mb-0.5">
         Location
        </p>
        <p class="font-body text-body text-text-primary truncate">
         Phase III, Fl 4
        </p>
       </div>
       <div>
        <p class="font-table-header text-table-header text-on-surface-variant uppercase mb-0.5">
         Value
        </p>
        <p class="font-body text-body text-text-primary">
         ₹ 2,85,000
        </p>
       </div>
      </div>
      <button class="w-full py-2 bg-surface border border-border-default rounded-lg font-h3 text-h3 text-text-primary hover:bg-surface-subtle hover:border-[#D6D3CE] transition-colors">
       View Details
      </button>
     </div>
    </div>
    <!-- Asset Card 2 -->
    <div class="bg-surface rounded-[16px] border border-border-default shadow-[0_4px_12px_rgba(0,0,0,0.03)] overflow-hidden hover:-translate-y-[2px] transition-all duration-200 group flex flex-col">
     <div class="h-32 bg-surface-subtle relative">
      <img alt="Headphones" class="w-full h-full object-cover opacity-90 group-hover:opacity-100 transition-opacity" data-alt="Premium over-ear wireless headphones resting on a dark slate surface with dramatic side lighting" src="https://lh3.googleusercontent.com/aida-public/AB6AXuAvSExjn8cuQeJPFKLVDln3lZ-YvQ3oOvTQIN5n2YJKg28hzLweg9wcT9vLLhhPhq0gE5FHO4NTKmcnwKHyNYmd_mHfvYYg4apuReyBo88GUX8oxlqphYyMWnwDBckNTwSXkxITjzBvyc3lHP1E-nfkzhgd9S-47EiQ9suu-DQz1KOXafxT_KTUIWydNkZ-Q260esOKclWK4F4hMKz1SzZWvDf6x_a1fJlPctSSU0q5gRBUGl4rMqDqPa65IN_n08SdOSNcl1LYHlTS"/>
      <div class="absolute top-3 left-3 bg-white/90 backdrop-blur-sm px-2 py-1 rounded shadow-sm border border-border-default/50 flex items-center gap-1">
       <span class="w-2 h-2 rounded-full bg-status-checked-out">
       </span>
       <span class="font-table-header text-table-header text-text-primary">
        Checked Out
       </span>
      </div>
     </div>
     <div class="p-5 flex-1 flex flex-col">
      <div class="flex justify-between items-start mb-2">
       <span class="font-mono-data text-mono-data text-primary-container bg-accent-subtle px-2 py-1 rounded">
        TPK-IT-2023-112
       </span>
      </div>
      <h3 class="font-h2 text-h2 text-text-primary mb-1">
       Sony WH-1000XM5
      </h3>
      <p class="font-metadata text-metadata text-text-secondary mb-4 line-clamp-1">
       Checked out by: Priya S.
      </p>
      <div class="mt-auto grid grid-cols-2 gap-2 mb-4 pt-4 border-t border-surface-subtle">
       <div>
        <p class="font-table-header text-table-header text-on-surface-variant uppercase mb-0.5">
         Return Date
        </p>
        <p class="font-body text-body text-text-primary truncate">
         15 Nov 2023
        </p>
       </div>
       <div>
        <p class="font-table-header text-table-header text-on-surface-variant uppercase mb-0.5">
         Value
        </p>
        <p class="font-body text-body text-text-primary">
         ₹ 29,990
        </p>
       </div>
      </div>
      <button class="w-full py-2 bg-surface border border-border-default rounded-lg font-h3 text-h3 text-text-primary hover:bg-surface-subtle hover:border-[#D6D3CE] transition-colors">
       View Details
      </button>
     </div>
    </div>
    <!-- Asset Card 3 -->
    <div class="bg-surface rounded-[16px] border border-border-default shadow-[0_4px_12px_rgba(0,0,0,0.03)] overflow-hidden hover:-translate-y-[2px] transition-all duration-200 group flex flex-col">
     <div class="h-32 bg-surface-subtle relative flex items-center justify-center bg-gradient-to-br from-surface-subtle to-[#E5E7EB]">
      <span class="material-symbols-outlined text-[48px] text-text-secondary/50">
       print
      </span>
      <div class="absolute top-3 left-3 bg-white/90 backdrop-blur-sm px-2 py-1 rounded shadow-sm border border-border-default/50 flex items-center gap-1">
       <span class="w-2 h-2 rounded-full bg-status-critical">
       </span>
       <span class="font-table-header text-table-header text-text-primary">
        Needs Repair
       </span>
      </div>
     </div>
     <div class="p-5 flex-1 flex flex-col">
      <div class="flex justify-between items-start mb-2">
       <span class="font-mono-data text-mono-data text-primary-container bg-accent-subtle px-2 py-1 rounded">
        TPK-IT-2022-045
       </span>
      </div>
      <h3 class="font-h2 text-h2 text-text-primary mb-1">
       HP LaserJet Pro MFP
      </h3>
      <p class="font-metadata text-metadata text-text-secondary mb-4 line-clamp-1">
       Issue: Fuser error, paper jam
      </p>
      <div class="mt-auto grid grid-cols-2 gap-2 mb-4 pt-4 border-t border-surface-subtle">
       <div>
        <p class="font-table-header text-table-header text-on-surface-variant uppercase mb-0.5">
         Location
        </p>
        <p class="font-body text-body text-text-primary truncate">
         Phase I, Block B
        </p>
       </div>
       <div>
        <p class="font-table-header text-table-header text-on-surface-variant uppercase mb-0.5">
         Value
        </p>
        <p class="font-body text-body text-text-primary">
         ₹ 45,500
        </p>
       </div>
      </div>
      <button class="w-full py-2 bg-primary-container text-white rounded-lg font-h3 text-h3 hover:bg-primary transition-colors">
       Log Maintenance
      </button>
     </div>
    </div>
    <!-- Asset Card 4 -->
    <div class="bg-surface rounded-[16px] border border-border-default shadow-[0_4px_12px_rgba(0,0,0,0.03)] overflow-hidden hover:-translate-y-[2px] transition-all duration-200 group flex flex-col">
     <div class="h-32 bg-surface-subtle relative">
      <img alt="Dell Monitor" class="w-full h-full object-cover opacity-90 group-hover:opacity-100 transition-opacity" data-alt="A sleek black Dell monitor standing on a desk with a blurred office background" src="https://lh3.googleusercontent.com/aida-public/AB6AXuAPZo3BqjcsttE5AQxs5EVA9LaH95g9WsG6xXRiY5rjnh3ECCkMqrzuBKU8t0Fukgu1SrfLXlZC4ffFy9bxQKzEQQD03oSJRnzY4iSIniAD2qPYbGtppFY3w_j6bfVHX50FufQBnANc2GfWuCICbTKVWuCrBbadfYY7gVOHD-TJHFxc-CyxwEsbrOPNk0QPkAQSVGA-kZiqusF0P7BgKvTEB5FFMCo_1P2-C8gvB7qBHTmh1FQPhjKusIWEMjO5RjpDFTQ8gQIfClWy"/>
      <div class="absolute top-3 left-3 bg-white/90 backdrop-blur-sm px-2 py-1 rounded shadow-sm border border-border-default/50 flex items-center gap-1">
       <span class="w-2 h-2 rounded-full bg-status-in-stock">
       </span>
       <span class="font-table-header text-table-header text-text-primary">
        In Use
       </span>
      </div>
     </div>
     <div class="p-5 flex-1 flex flex-col">
      <div class="flex justify-between items-start mb-2">
       <span class="font-mono-data text-mono-data text-primary-container bg-accent-subtle px-2 py-1 rounded">
        TPK-IT-2023-201
       </span>
      </div>
      <h3 class="font-h2 text-h2 text-text-primary mb-1">
       Dell UltraSharp 27" 4K
      </h3>
      <p class="font-metadata text-metadata text-text-secondary mb-4 line-clamp-1">
       Assigned to: Design Team
      </p>
      <div class="mt-auto grid grid-cols-2 gap-2 mb-4 pt-4 border-t border-surface-subtle">
       <div>
        <p class="font-table-header text-table-header text-on-surface-variant uppercase mb-0.5">
         Location
        </p>
        <p class="font-body text-body text-text-primary truncate">
         Phase III, Fl 2
        </p>
       </div>
       <div>
        <p class="font-table-header text-table-header text-on-surface-variant uppercase mb-0.5">
         Value
        </p>
        <p class="font-body text-body text-text-primary">
         ₹ 52,000
        </p>
       </div>
      </div>
      <button class="w-full py-2 bg-surface border border-border-default rounded-lg font-h3 text-h3 text-text-primary hover:bg-surface-subtle hover:border-[#D6D3CE] transition-colors">
       View Details
      </button>
     </div>
    </div>
   </div>
  </div>
 </div>
</main>

</template>

<script setup lang="ts">
// Autogenerated from asset_registry_redesigned
</script>
