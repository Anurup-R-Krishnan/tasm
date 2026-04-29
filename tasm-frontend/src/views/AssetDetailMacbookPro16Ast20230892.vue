<template>
<main class="md:pl-[248px] min-h-screen">
 <div class="p-page-margin max-w-[1400px] mx-auto">
  <!-- Breadcrumbs & Header -->
  <div class="mb-section-gap">
   <div class="flex items-center gap-2 text-metadata text-text-secondary font-metadata mb-stack">
    <a class="hover:text-primary transition-colors" href="#">
     Asset Inventory
    </a>
    <span class="material-symbols-outlined text-[14px]">
     chevron_right
    </span>
    <a class="hover:text-primary transition-colors" href="#">
     IT Infrastructure
    </a>
    <span class="material-symbols-outlined text-[14px]">
     chevron_right
    </span>
    <span class="text-text-primary font-medium">
     SVR-2023-X99
    </span>
   </div>
   <div class="flex justify-between items-start">
    <div>
     <div class="flex items-center gap-inline mb-2">
      <h1 class="font-h1 text-h1 text-text-primary">
       Enterprise Database Server Rack
      </h1>
      <span class="inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium bg-metric-sage text-status-in-stock border border-[#BBF7D0]">
       <span class="sr-only">
        Status:
       </span>
       Active
      </span>
     </div>
     <div class="flex items-center gap-4 text-metadata text-text-secondary font-metadata">
      <div class="flex items-center gap-1">
       <span class="material-symbols-outlined text-[14px]">
        tag
       </span>
       <span class="font-mono-data text-mono-data text-text-primary bg-surface-subtle px-1.5 py-0.5 rounded border border-border-default">
        SVR-2023-X99
       </span>
      </div>
      <div class="flex items-center gap-1">
       <span class="material-symbols-outlined text-[14px]">
        location_on
       </span>
       <span>
        Data Center Alpha, Row C
       </span>
      </div>
     </div>
    </div>
    <div class="flex gap-inline">
     <button class="px-4 py-2 bg-surface text-text-primary border border-border-default rounded-lg font-h3 text-h3 shadow-[0_2px_4px_rgba(0,0,0,0.02)] hover:-translate-y-[1px] transition-transform flex items-center gap-2">
      <span class="material-symbols-outlined text-[18px]">
       edit
      </span>
      Edit
     </button>
     <button class="px-4 py-2 bg-[#1C1917] text-white rounded-lg font-h3 text-h3 shadow-sm hover:-translate-y-[1px] transition-transform flex items-center gap-2">
      <span class="material-symbols-outlined text-[18px]">
       handyman
      </span>
      Log Maintenance
     </button>
    </div>
   </div>
  </div>
  <!-- Top Section: Image + KPIs -->
  <div class="grid grid-cols-1 lg:grid-cols-12 gap-inline mb-section-gap">
   <!-- Image Gallery -->
   <div class="lg:col-span-5 bg-surface border border-border-default rounded-xl overflow-hidden shadow-[0_4px_12px_rgba(0,0,0,0.03)] flex flex-col">
    <div class="relative h-64 w-full bg-stone-100 border-b border-border-default">
     <img alt="Server rack main view" class="w-full h-full object-cover" data-alt="high end enterprise database server rack in modern data center with blue led lights" src="https://lh3.googleusercontent.com/aida-public/AB6AXuDHt9h6pblN8VURl3u8hTDEQxZRm42TeoqIPPr6nbk8uhk0pNJCe_lAx-dqhhXNrH6PcVnaQp6ZOIpvrpWi9pokJWUV7e7yYz_JripZODPn7iu-t9Rq1yxuBaHq0XMau9GVS6MqUDWrwBvephQG-ZDxSr90rzKxHd44nJJP42xnXhKUoZaduKkVJMb8UM8bpB8OCibVNopqrB3zxwc3oFntkqBMIJJpd_6aRDFGK3hJ0PPqIa69s5fe8h3_9hv531uhyrrqlqSAXmLW"/>
     <button class="absolute top-4 right-4 p-2 bg-white/80 backdrop-blur-sm rounded-lg border border-white/20 shadow-sm text-text-primary hover:bg-white transition-colors">
      <span class="material-symbols-outlined text-[20px]">
       fullscreen
      </span>
     </button>
    </div>
    <div class="flex gap-2 p-4 bg-surface-subtle overflow-x-auto">
     <div class="h-16 w-16 flex-shrink-0 rounded-lg border-2 border-primary overflow-hidden">
      <img alt="Thumbnail 1" class="w-full h-full object-cover" data-alt="server rack close up detail view" src="https://lh3.googleusercontent.com/aida-public/AB6AXuBgICiw28-FbgVwsWazMRoEjffoIFlMhLfhwNZj6h9p6yVweXDlRobX1KgZ_YKyZpRKLGpb0UUVPXg-uvuXk31EKOJwW_341Wgejd2SelDZGXacTvyGLGOIa8uEJJBHFtCvZmAf3ybpEExJ-XhvqD-ujQKzrCEuuWkZ7FTt7ZXXVzJCZpTz3RVboyHHTatUDqRqsoe1wClnFdefSLbLiaUHkndJiZvaWI9vWfrxOMGbPvihjN9PcvlMqM0HtbFG96fNw1I0YGGbd5q-"/>
     </div>
     <div class="h-16 w-16 flex-shrink-0 rounded-lg border border-border-default opacity-70 hover:opacity-100 transition-opacity overflow-hidden cursor-pointer">
      <img alt="Thumbnail 2" class="w-full h-full object-cover" data-alt="circuit board macro shot green" src="https://lh3.googleusercontent.com/aida-public/AB6AXuB7fX0rIiLXXmpiZBEQXPN__4D7n6r_Lv32gyATmcRFdF53oc1cScT6TDbuOIvj9Eb8c7IMunPt_ii2SqSGD46WJV5VvQPFZ4am9Gifr30-mGS4jZoNf_KgkjL_KotqA05onhLwAtLgz_5G9TOQOaid3fQCICuOuDkyQHdwbG_lZ8RspMBTuWvH_RINSqb5RlfyC3gvCl0XkRxCJuhx2JOlX9-vzerT0-YODWytgzrkJR1r_n579y7pMA7g1tI9GgJDsd_YZfGOy4aW"/>
     </div>
     <div class="h-16 w-16 flex-shrink-0 rounded-lg border border-border-default opacity-70 hover:opacity-100 transition-opacity flex items-center justify-center bg-stone-100 cursor-pointer">
      <span class="material-symbols-outlined text-stone-400">
       add_photo_alternate
      </span>
     </div>
    </div>
   </div>
   <!-- KPIs -->
   <div class="lg:col-span-7 grid grid-cols-2 md:grid-cols-3 gap-inline">
    <div class="bg-surface border border-border-default rounded-xl p-card-padding shadow-[0_4px_12px_rgba(0,0,0,0.03)] hover:-translate-y-[2px] transition-transform cursor-pointer group flex flex-col justify-between">
     <div class="flex justify-between items-start mb-4">
      <div class="w-10 h-10 rounded-lg bg-metric-amber flex items-center justify-center text-amber-700">
       <span class="material-symbols-outlined">
        payments
       </span>
      </div>
      <span class="material-symbols-outlined text-stone-300 group-hover:text-amber-500 transition-colors">
       open_in_new
      </span>
     </div>
     <div>
      <div class="text-metadata text-text-secondary font-metadata mb-1">
       Current Book Value
      </div>
      <div class="font-kpi-number text-kpi-number text-text-primary tracking-tight">
       ₹12,45,000
      </div>
      <div class="text-xs text-stone-500 mt-1 flex items-center gap-1">
       <span class="material-symbols-outlined text-[14px] text-red-500">
        trending_down
       </span>
       -15% YoY Depr.
      </div>
     </div>
    </div>
    <div class="bg-surface border border-border-default rounded-xl p-card-padding shadow-[0_4px_12px_rgba(0,0,0,0.03)] hover:-translate-y-[2px] transition-transform cursor-pointer group flex flex-col justify-between">
     <div class="flex justify-between items-start mb-4">
      <div class="w-10 h-10 rounded-lg bg-metric-lavender flex items-center justify-center text-purple-700">
       <span class="material-symbols-outlined">
        health_and_safety
       </span>
      </div>
      <span class="material-symbols-outlined text-stone-300 group-hover:text-purple-500 transition-colors">
       open_in_new
      </span>
     </div>
     <div>
      <div class="text-metadata text-text-secondary font-metadata mb-1">
       Condition Score
      </div>
      <div class="font-kpi-number text-kpi-number text-text-primary tracking-tight">
       92
       <span class="text-lg text-stone-400">
        /100
       </span>
      </div>
      <div class="text-xs text-stone-500 mt-1">
       Excellent (Next audit: Oct '24)
      </div>
     </div>
    </div>
    <div class="bg-surface border border-border-default rounded-xl p-card-padding shadow-[0_4px_12px_rgba(0,0,0,0.03)] hover:-translate-y-[2px] transition-transform cursor-pointer group flex flex-col justify-between col-span-2 md:col-span-1">
     <div class="flex justify-between items-start mb-4">
      <div class="w-10 h-10 rounded-lg bg-stone-100 flex items-center justify-center text-stone-700">
       <span class="material-symbols-outlined">
        assignment_ind
       </span>
      </div>
      <span class="material-symbols-outlined text-stone-300 group-hover:text-stone-500 transition-colors">
       open_in_new
      </span>
     </div>
     <div>
      <div class="text-metadata text-text-secondary font-metadata mb-1">
       Primary Custodian
      </div>
      <div class="font-h2 text-h2 text-text-primary truncate">
       IT Infrastructure Dept
      </div>
      <div class="text-xs text-stone-500 mt-1 truncate">
       Dept ID: IT-INF-01
      </div>
     </div>
    </div>
   </div>
  </div>
  <!-- Middle & Right Split -->
  <div class="grid grid-cols-1 lg:grid-cols-12 gap-inline">
   <!-- Detailed Tabs Section -->
   <div class="lg:col-span-8 bg-surface border border-border-default rounded-xl shadow-[0_4px_12px_rgba(0,0,0,0.03)] overflow-hidden flex flex-col">
    <!-- Tabs Header -->
    <div class="border-b border-border-default px-6 pt-4 flex gap-6 overflow-x-auto bg-surface-subtle">
     <button class="pb-3 border-b-2 border-primary font-h3 text-h3 text-primary whitespace-nowrap">
      Identity &amp; Specs
     </button>
     <button class="pb-3 border-b-2 border-transparent font-h3 text-h3 text-text-secondary hover:text-text-primary transition-colors whitespace-nowrap">
      Financials
     </button>
     <button class="pb-3 border-b-2 border-transparent font-h3 text-h3 text-text-secondary hover:text-text-primary transition-colors whitespace-nowrap">
      Maintenance logs
     </button>
     <button class="pb-3 border-b-2 border-transparent font-h3 text-h3 text-text-secondary hover:text-text-primary transition-colors whitespace-nowrap">
      Custody Chain
     </button>
    </div>
    <!-- Tab Content (Identity) -->
    <div class="p-6 flex-1">
     <div class="grid grid-cols-1 md:grid-cols-2 gap-x-8 gap-y-6">
      <div>
       <h3 class="font-table-header text-table-header text-text-secondary uppercase mb-4">
        Manufacturer Details
       </h3>
       <dl class="space-y-4">
        <div class="grid grid-cols-3 gap-4 border-b border-stone-100 pb-2">
         <dt class="text-sm text-text-secondary">
          Brand
         </dt>
         <dd class="text-sm font-medium text-text-primary col-span-2">
          Cisco Systems
         </dd>
        </div>
        <div class="grid grid-cols-3 gap-4 border-b border-stone-100 pb-2">
         <dt class="text-sm text-text-secondary">
          Model Number
         </dt>
         <dd class="text-sm font-medium text-text-primary col-span-2 font-mono-data">
          UCS-C220-M5SX
         </dd>
        </div>
        <div class="grid grid-cols-3 gap-4 border-b border-stone-100 pb-2">
         <dt class="text-sm text-text-secondary">
          Serial Number
         </dt>
         <dd class="text-sm font-medium text-text-primary col-span-2 font-mono-data">
          FCH2240V1A4
         </dd>
        </div>
        <div class="grid grid-cols-3 gap-4">
         <dt class="text-sm text-text-secondary">
          Year of Mfg
         </dt>
         <dd class="text-sm font-medium text-text-primary col-span-2">
          2021
         </dd>
        </div>
       </dl>
      </div>
      <div>
       <h3 class="font-table-header text-table-header text-text-secondary uppercase mb-4">
        Technical Specs
       </h3>
       <dl class="space-y-4">
        <div class="grid grid-cols-3 gap-4 border-b border-stone-100 pb-2">
         <dt class="text-sm text-text-secondary">
          Processor
         </dt>
         <dd class="text-sm font-medium text-text-primary col-span-2">
          Dual Intel Xeon Silver 4214
         </dd>
        </div>
        <div class="grid grid-cols-3 gap-4 border-b border-stone-100 pb-2">
         <dt class="text-sm text-text-secondary">
          Memory
         </dt>
         <dd class="text-sm font-medium text-text-primary col-span-2">
          256GB DDR4 ECC
         </dd>
        </div>
        <div class="grid grid-cols-3 gap-4 border-b border-stone-100 pb-2">
         <dt class="text-sm text-text-secondary">
          Storage
         </dt>
         <dd class="text-sm font-medium text-text-primary col-span-2">
          4x 1.2TB SAS 10K RPM
         </dd>
        </div>
        <div class="grid grid-cols-3 gap-4">
         <dt class="text-sm text-text-secondary">
          Power Supply
         </dt>
         <dd class="text-sm font-medium text-text-primary col-span-2">
          Dual Redundant 770W AC
         </dd>
        </div>
       </dl>
      </div>
     </div>
     <div class="mt-8">
      <h3 class="font-table-header text-table-header text-text-secondary uppercase mb-4">
       Attached Documents
      </h3>
      <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4">
       <div class="p-3 border border-border-default rounded-lg flex items-start gap-3 hover:bg-stone-50 cursor-pointer transition-colors">
        <span class="material-symbols-outlined text-red-500">
         picture_as_pdf
        </span>
        <div class="overflow-hidden">
         <div class="text-sm font-medium text-text-primary truncate">
          Purchase_Invoice_2021.pdf
         </div>
         <div class="text-xs text-text-secondary mt-0.5">
          1.2 MB
         </div>
        </div>
       </div>
       <div class="p-3 border border-border-default rounded-lg flex items-start gap-3 hover:bg-stone-50 cursor-pointer transition-colors">
        <span class="material-symbols-outlined text-blue-500">
         description
        </span>
        <div class="overflow-hidden">
         <div class="text-sm font-medium text-text-primary truncate">
          Warranty_Certificate.docx
         </div>
         <div class="text-xs text-text-secondary mt-0.5">
          450 KB
         </div>
        </div>
       </div>
      </div>
     </div>
    </div>
   </div>
   <!-- Right Side: Immutable Activity Feed -->
   <div class="lg:col-span-4 bg-surface border border-border-default rounded-xl shadow-[0_4px_12px_rgba(0,0,0,0.03)] flex flex-col h-[600px]">
    <div class="p-4 border-b border-border-default bg-surface-subtle rounded-t-xl flex justify-between items-center">
     <h2 class="font-h2 text-h2 text-text-primary">
      Audit Trail
     </h2>
     <button class="text-xs text-text-secondary hover:text-primary transition-colors flex items-center gap-1">
      <span class="material-symbols-outlined text-[14px]">
       filter_list
      </span>
      Filter
     </button>
    </div>
    <div class="p-6 flex-1 overflow-y-auto">
     <div class="relative border-l border-stone-200 ml-3 space-y-6">
      <!-- Feed Item -->
      <div class="relative pl-6">
       <div class="absolute -left-1.5 top-1 w-3 h-3 bg-metric-sage border-2 border-white rounded-full">
       </div>
       <div class="text-xs text-text-secondary font-metadata mb-1">
        Today, 10:45 AM
       </div>
       <div class="bg-stone-50 border border-stone-100 p-3 rounded-lg">
        <p class="text-sm text-text-primary font-medium mb-1">
         Routine Maintenance Completed
        </p>
        <p class="text-xs text-text-secondary">
         Firmware updated to v4.2.1. Dust cleaning performed. Logged by
         <span class="font-medium text-stone-700">
          Rahul M.
         </span>
        </p>
       </div>
      </div>
      <!-- Feed Item -->
      <div class="relative pl-6">
       <div class="absolute -left-1.5 top-1 w-3 h-3 bg-stone-300 border-2 border-white rounded-full">
       </div>
       <div class="text-xs text-text-secondary font-metadata mb-1">
        Oct 12, 2023, 02:15 PM
       </div>
       <div class="bg-stone-50 border border-stone-100 p-3 rounded-lg">
        <p class="text-sm text-text-primary font-medium mb-1">
         Asset Location Verified
        </p>
        <p class="text-xs text-text-secondary">
         Quarterly audit scan successful at Data Center Alpha. System automated log.
        </p>
       </div>
      </div>
      <!-- Feed Item -->
      <div class="relative pl-6">
       <div class="absolute -left-1.5 top-1 w-3 h-3 bg-metric-amber border-2 border-white rounded-full">
       </div>
       <div class="text-xs text-text-secondary font-metadata mb-1">
        Sep 05, 2023, 09:30 AM
       </div>
       <div class="bg-stone-50 border border-stone-100 p-3 rounded-lg">
        <p class="text-sm text-text-primary font-medium mb-1">
         Depreciation Calculated
        </p>
        <p class="text-xs text-text-secondary">
         Annual depreciation applied. Book value updated to ₹12,45,000.
        </p>
       </div>
      </div>
      <!-- Feed Item -->
      <div class="relative pl-6">
       <div class="absolute -left-1.5 top-1 w-3 h-3 bg-stone-800 border-2 border-white rounded-full">
       </div>
       <div class="text-xs text-text-secondary font-metadata mb-1">
        Jan 10, 2021, 11:00 AM
       </div>
       <div class="bg-stone-50 border border-stone-100 p-3 rounded-lg">
        <p class="text-sm text-text-primary font-medium mb-1">
         Asset Registered
        </p>
        <p class="text-xs text-text-secondary">
         Initial entry into AMS. Purchase value: ₹18,50,000. Invoice #INV-2021-089 attached.
        </p>
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
// Autogenerated from asset_detail_macbook_pro_16_ast_2023_0892
</script>
