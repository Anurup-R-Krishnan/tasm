<template>
<main class="flex-1 mt-[60px] p-page-margin overflow-y-auto relative z-0 flex gap-6">
 <!-- Left Column: KPIs and Table -->
 <div class="flex-1 flex flex-col gap-6">
  <!-- Header -->
  <div class="flex justify-between items-end">
   <div>
    <h1 class="font-h1 text-h1 text-text-primary mb-1">
     Software Licenses
    </h1>
    <p class="text-text-secondary font-body text-body">
     Manage subscriptions, seats, and renewals
    </p>
   </div>
   <div class="flex gap-3">
    <button class="px-4 py-2 bg-surface text-text-primary border border-border-default rounded-lg font-medium text-sm hover:bg-surface-subtle transition-colors flex items-center gap-2 shadow-sm">
     <span class="material-symbols-outlined text-[18px]">
      download
     </span>
     Export CSV
    </button>
   </div>
  </div>
  <!-- KPI Cards Grid -->
  <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
   <!-- Expiring Licenses -->
   <div class="bg-surface border border-border-default rounded-xl p-card-padding shadow-sm hover:-translate-y-[2px] transition-transform cursor-pointer relative overflow-hidden group">
    <div class="absolute inset-0 bg-gradient-to-br from-[#FFE4E6]/40 to-transparent z-0">
    </div>
    <div class="relative z-10 flex justify-between items-start">
     <div>
      <p class="font-metadata text-metadata text-text-secondary uppercase tracking-wider mb-2">
       Expiring &lt; 30 Days
      </p>
      <h2 class="font-kpi-number text-kpi-number text-text-primary mb-1">
       12
      </h2>
      <p class="text-sm text-status-critical font-medium flex items-center gap-1">
       <span class="material-symbols-outlined text-[16px]">
        warning
       </span>
       Action Required
      </p>
     </div>
     <div class="w-10 h-10 rounded-full bg-[#FFE4E6] flex items-center justify-center text-status-critical">
      <span class="material-symbols-outlined">
       event_busy
      </span>
     </div>
    </div>
   </div>
   <!-- Underutilized Seats -->
   <div class="bg-surface border border-border-default rounded-xl p-card-padding shadow-sm hover:-translate-y-[2px] transition-transform cursor-pointer relative overflow-hidden group">
    <div class="absolute inset-0 bg-gradient-to-br from-metric-amber/50 to-transparent z-0">
    </div>
    <div class="relative z-10 flex justify-between items-start">
     <div>
      <p class="font-metadata text-metadata text-text-secondary uppercase tracking-wider mb-2">
       Underutilized Seats
      </p>
      <h2 class="font-kpi-number text-kpi-number text-text-primary mb-1">
       45
      </h2>
      <p class="text-sm text-surface-tint font-medium flex items-center gap-1">
       <span class="material-symbols-outlined text-[16px]">
        trending_down
       </span>
       Potential savings
      </p>
     </div>
     <div class="w-10 h-10 rounded-full bg-metric-amber flex items-center justify-center text-surface-tint">
      <span class="material-symbols-outlined">
       group_remove
      </span>
     </div>
    </div>
   </div>
   <!-- Annual Spend -->
   <div class="bg-surface border border-border-default rounded-xl p-card-padding shadow-sm hover:-translate-y-[2px] transition-transform cursor-pointer relative overflow-hidden group">
    <div class="absolute inset-0 bg-gradient-to-br from-metric-lavender/50 to-transparent z-0">
    </div>
    <div class="relative z-10 flex justify-between items-start">
     <div>
      <p class="font-metadata text-metadata text-text-secondary uppercase tracking-wider mb-2">
       Annual Software Spend
      </p>
      <h2 class="font-kpi-number text-kpi-number text-text-primary mb-1">
       $1.2M
      </h2>
      <p class="text-sm text-text-secondary flex items-center gap-1">
       YTD FY2024
      </p>
     </div>
     <div class="w-10 h-10 rounded-full bg-metric-lavender flex items-center justify-center text-[#5B21B6]">
      <span class="material-symbols-outlined">
       payments
      </span>
     </div>
    </div>
   </div>
  </div>
  <!-- License Table -->
  <div class="bg-surface border border-border-default rounded-xl shadow-sm overflow-hidden flex-1 flex flex-col">
   <div class="p-4 border-b border-border-default flex justify-between items-center bg-surface">
    <div class="flex gap-2">
     <button class="px-3 py-1.5 rounded-md bg-stone-100 text-sm font-medium text-stone-800">
      All Licenses
     </button>
     <button class="px-3 py-1.5 rounded-md hover:bg-stone-50 text-sm font-medium text-stone-600">
      Active
     </button>
     <button class="px-3 py-1.5 rounded-md hover:bg-stone-50 text-sm font-medium text-stone-600">
      Expiring Soon
     </button>
    </div>
    <button class="text-text-secondary hover:text-text-primary">
     <span class="material-symbols-outlined">
      filter_list
     </span>
    </button>
   </div>
   <div class="overflow-x-auto flex-1">
    <table class="w-full text-left border-collapse">
     <thead>
      <tr class="border-b border-border-default bg-surface-subtle font-table-header text-table-header text-text-secondary uppercase">
       <th class="py-3 px-4 w-1/3">
        Software Asset
       </th>
       <th class="py-3 px-4">
        Status
       </th>
       <th class="py-3 px-4 w-64">
        Seat Utilization
       </th>
       <th class="py-3 px-4">
        Renewal Date
       </th>
       <th class="py-3 px-4 text-right">
        Ann. Cost
       </th>
      </tr>
     </thead>
     <tbody class="font-body text-sm">
      <!-- Row 1: Selected -->
      <tr class="border-b border-border-default bg-metric-amber/20 hover:bg-metric-amber/40 transition-colors cursor-pointer relative">
       <td class="py-3 px-4 absolute left-0 top-0 bottom-0 w-[3px] bg-amber-600">
       </td>
       <td class="py-3 px-4">
        <div class="flex items-center gap-3">
         <div class="w-8 h-8 rounded bg-red-50 flex items-center justify-center text-red-600">
          <span class="material-symbols-outlined text-[18px]">
           brush
          </span>
         </div>
         <div>
          <div class="font-medium text-text-primary">
           Adobe Creative Cloud
          </div>
          <div class="text-xs text-text-secondary">
           Enterprise All Apps
          </div>
         </div>
        </div>
       </td>
       <td class="py-3 px-4">
        <span class="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium bg-status-in-stock/10 text-status-in-stock">
         <span class="sr-only">
          Status:
         </span>
         Active
        </span>
       </td>
       <td class="py-3 px-4">
        <div class="flex items-center gap-2">
         <div class="flex-1 h-1.5 bg-stone-100 rounded-full overflow-hidden">
          <div class="h-full bg-amber-500 rounded-full" style="width: 85%">
          </div>
         </div>
         <span class="font-mono-data text-mono-data text-text-secondary w-12 text-right">
          85/100
         </span>
        </div>
       </td>
       <td class="py-3 px-4">
        <div class="font-mono-data text-mono-data">
         2024-11-15
        </div>
       </td>
       <td class="py-3 px-4 text-right font-mono-data text-mono-data">
        $84,000
       </td>
      </tr>
      <!-- Row 2 -->
      <tr class="border-b border-border-default hover:bg-surface-subtle transition-colors cursor-pointer">
       <td class="py-3 px-4">
        <div class="flex items-center gap-3">
         <div class="w-8 h-8 rounded bg-blue-50 flex items-center justify-center text-blue-600">
          <span class="material-symbols-outlined text-[18px]">
           cloud
          </span>
         </div>
         <div>
          <div class="font-medium text-text-primary">
           Microsoft 365
          </div>
          <div class="text-xs text-text-secondary">
           E5 License
          </div>
         </div>
        </div>
       </td>
       <td class="py-3 px-4">
        <span class="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium bg-status-in-stock/10 text-status-in-stock">
         <span class="sr-only">
          Status:
         </span>
         Active
        </span>
       </td>
       <td class="py-3 px-4">
        <div class="flex items-center gap-2">
         <div class="flex-1 h-1.5 bg-stone-100 rounded-full overflow-hidden">
          <div class="h-full bg-status-in-stock rounded-full" style="width: 95%">
          </div>
         </div>
         <span class="font-mono-data text-mono-data text-text-secondary w-12 text-right">
          475/500
         </span>
        </div>
       </td>
       <td class="py-3 px-4">
        <div class="font-mono-data text-mono-data">
         2025-03-01
        </div>
       </td>
       <td class="py-3 px-4 text-right font-mono-data text-mono-data">
        $228,000
       </td>
      </tr>
      <!-- Row 3: Expiring Soon -->
      <tr class="border-b border-border-default hover:bg-surface-subtle transition-colors cursor-pointer">
       <td class="py-3 px-4">
        <div class="flex items-center gap-3">
         <div class="w-8 h-8 rounded bg-purple-50 flex items-center justify-center text-purple-600">
          <span class="material-symbols-outlined text-[18px]">
           design_services
          </span>
         </div>
         <div>
          <div class="font-medium text-text-primary">
           Figma
          </div>
          <div class="text-xs text-text-secondary">
           Organization Plan
          </div>
         </div>
        </div>
       </td>
       <td class="py-3 px-4">
        <span class="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium bg-amber-100 text-amber-800 border border-amber-200">
         <span class="sr-only">
          Status:
         </span>
         Expiring Soon
        </span>
       </td>
       <td class="py-3 px-4">
        <div class="flex items-center gap-2">
         <div class="flex-1 h-1.5 bg-stone-100 rounded-full overflow-hidden">
          <div class="h-full bg-status-in-stock rounded-full" style="width: 100%">
          </div>
         </div>
         <span class="font-mono-data text-mono-data text-text-secondary w-12 text-right">
          50/50
         </span>
        </div>
       </td>
       <td class="py-3 px-4">
        <div class="font-mono-data text-mono-data text-amber-700 font-medium">
         2024-05-10
        </div>
       </td>
       <td class="py-3 px-4 text-right font-mono-data text-mono-data">
        $27,000
       </td>
      </tr>
     </tbody>
    </table>
   </div>
  </div>
 </div>
 <!-- Right Drawer (Level 2 Elevation) -->
 <div class="w-[480px] bg-surface border-l border-border-default shadow-[0_0_24px_rgba(0,0,0,0.05)] rounded-l-2xl h-[calc(100vh-60px-64px)] sticky top-[32px] flex flex-col z-20">
  <!-- Drawer Header -->
  <div class="p-6 border-b border-border-default flex justify-between items-start">
   <div class="flex gap-4">
    <div class="w-12 h-12 rounded-lg bg-red-50 flex items-center justify-center text-red-600 shadow-sm border border-red-100">
     <span class="material-symbols-outlined text-[24px]">
      brush
     </span>
    </div>
    <div>
     <h2 class="font-h2 text-h2 text-text-primary">
      Adobe Creative Cloud
     </h2>
     <p class="text-sm text-text-secondary mt-1">
      Enterprise All Apps • Vendor: Adobe Inc.
     </p>
     <div class="mt-2 flex gap-2">
      <span class="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium bg-status-in-stock/10 text-status-in-stock">
       Active
      </span>
      <span class="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium bg-stone-100 text-stone-600 border border-stone-200">
       ID: LIC-2023-089
      </span>
     </div>
    </div>
   </div>
   <button class="text-text-secondary hover:bg-stone-100 p-1 rounded-md transition-colors">
    <span class="material-symbols-outlined">
     close
    </span>
   </button>
  </div>
  <!-- Drawer Tabs -->
  <div class="px-6 flex gap-6 border-b border-border-default">
   <button class="py-3 text-sm font-medium border-b-2 border-amber-600 text-amber-900">
    Overview
   </button>
   <button class="py-3 text-sm font-medium border-b-2 border-transparent text-text-secondary hover:text-text-primary">
    Assigned Users (85)
   </button>
   <button class="py-3 text-sm font-medium border-b-2 border-transparent text-text-secondary hover:text-text-primary">
    Invoices
   </button>
  </div>
  <!-- Drawer Content (Assigned Users Tab Selected Visually for Demo) -->
  <div class="p-6 flex-1 overflow-y-auto">
   <div class="flex justify-between items-center mb-4">
    <div class="relative w-64">
     <span class="material-symbols-outlined absolute left-2.5 top-1/2 -translate-y-1/2 text-stone-400 text-[18px]">
      search
     </span>
     <input class="w-full bg-stone-50 border border-border-default rounded-md pl-9 pr-3 py-1.5 text-sm focus:outline-none focus:border-amber-600" placeholder="Search users..." type="text"/>
    </div>
    <button class="text-amber-700 hover:text-amber-800 text-sm font-medium flex items-center gap-1">
     <span class="material-symbols-outlined text-[18px]">
      person_add
     </span>
     Assign User
    </button>
   </div>
   <div class="space-y-3">
    <!-- User Item -->
    <div class="flex items-center justify-between p-3 rounded-lg border border-border-default hover:bg-surface-subtle transition-colors">
     <div class="flex items-center gap-3">
      <img alt="Profile" class="w-8 h-8 rounded-full bg-stone-200" data-alt="Corporate headshot of a young professional woman smiling lightly, clean white background" src="https://lh3.googleusercontent.com/aida-public/AB6AXuBmhKDZK--iT2dE9l_2_UoPavDyYhL3qieLmskVM9BN4Cgge0yzhYlqs_mUSf4jg4IXsUz0jeNqvpYd-HKgL06JCNEXHo-OPIOku6d3U2KFBDaGxJFhxlNX7cvhqr5htoOA9__vMm783ZLsDncyZqSdl4aj-nokEs-u3Gdy7mKb9UNXYvoeRgH4ScXzC-zQyHoKE4T-gEldE7imFCeN1jEHhE11RM1mEkWV-5XxgI-e3_XDf1wPKhiLdJQhLNPv_G8vfcJEENeaP0ux"/>
      <div>
       <div class="text-sm font-medium text-text-primary">
        Sarah Jenkins
       </div>
       <div class="text-xs text-text-secondary">
        UI/UX Designer • Marketing Team
       </div>
      </div>
     </div>
     <div class="flex items-center gap-3">
      <span class="text-xs text-stone-500 font-mono-data">
       Last Active: 2h ago
      </span>
      <button class="text-stone-400 hover:text-error transition-colors" title="Revoke License">
       <span class="material-symbols-outlined text-[18px]">
        person_remove
       </span>
      </button>
     </div>
    </div>
    <!-- User Item -->
    <div class="flex items-center justify-between p-3 rounded-lg border border-border-default hover:bg-surface-subtle transition-colors">
     <div class="flex items-center gap-3">
      <img alt="Profile" class="w-8 h-8 rounded-full bg-stone-200" data-alt="Corporate headshot of a middle aged man with glasses, professional lighting" src="https://lh3.googleusercontent.com/aida-public/AB6AXuCqqaiXZQ-aYrBIl1F2NwbwmZxRS0YiGTmfIboAO-icsvTX5TcNy6Ao7vc5wn3BAPLp15rZ1CRSmQYPF4WOX91knqqzRqhsX-JTGN-nX9B0KPkLIUbuW6Q8p0vdoHwB4dkB4Op3oE5NMA42d3DDNhlBFKUqYoUDG6U2lPWLzaezWPu8A29PlM7l7Mq7hoPDN5-dgJTvgvzxfUZdfx7LJx3z4bjPczjUGAP_Z-tOuVwjfHoYJzjCETMnkTQwmeIJ0B_k3QLNvUumdxKN"/>
      <div>
       <div class="text-sm font-medium text-text-primary">
        David Chen
       </div>
       <div class="text-xs text-text-secondary">
        Creative Director • Design Studio
       </div>
      </div>
     </div>
     <div class="flex items-center gap-3">
      <span class="text-xs text-stone-500 font-mono-data">
       Last Active: 1d ago
      </span>
      <button class="text-stone-400 hover:text-error transition-colors" title="Revoke License">
       <span class="material-symbols-outlined text-[18px]">
        person_remove
       </span>
      </button>
     </div>
    </div>
   </div>
  </div>
  <!-- Drawer Footer Action -->
  <div class="p-6 border-t border-border-default bg-surface-subtle rounded-bl-2xl">
   <button class="w-full py-2.5 bg-[#1C1917] text-white rounded-lg font-medium hover:bg-stone-800 transition-colors">
    Manage Subscription
   </button>
  </div>
 </div>
</main>

</template>

<script setup lang="ts">
// Autogenerated from software_license_registry
</script>
