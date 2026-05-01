<template>
<main class="pl-[248px] min-h-screen">
 <div class="max-w-[1400px] mx-auto p-page-margin">
  <!-- Header Section -->
  <div class="flex flex-col md:flex-row md:items-start justify-between gap-stack mb-section-gap">
   <div>
    <div class="flex items-center gap-3 mb-2">
     <button class="text-text-secondary hover:text-text-primary transition-colors flex items-center justify-center rounded-full p-1 -ml-1">
      <span class="material-symbols-outlined">
       arrow_back
      </span>
     </button>
     <h1 v-if="request" class="font-h1 text-h1 text-text-primary flex items-center gap-3">
      <span class="font-mono-data text-mono-data bg-surface-variant px-2 py-1 rounded text-text-primary">
       {{ request.poNumber || 'PR-' + request.id }}
      </span>
      {{ request.title }}
     </h1>
     <span v-if="request" class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-metric-amber text-on-surface-variant border border-amber-200">
      <span class="w-1.5 h-1.5 rounded-full bg-amber-500 mr-1.5">
      </span>
      {{ request.status }}
     </span>
    </div>
    <div class="flex items-center gap-4 text-text-secondary ml-10">
     <div class="flex items-center gap-2">
      <img alt="Sarah Jenkins" class="w-6 h-6 rounded-full border border-border-default" data-alt="professional headshot of a woman with dark hair smiling gently" src="https://lh3.googleusercontent.com/aida-public/AB6AXuAMl2rAFUYQ96VmKc44gP_8GSOGpxIfOjgB4gNwbB-wLVrcSKIZvEfAMpAywsvt9FvBWKqGrHkcfQ-VqZnxZfYc5opM4f7EDVpbtOfG6-F9KOuz-8qE6uK_PC3mrZo7csekgodU5T5yEX3Qf0luQGL5qqblVJ3Y5Qtw2bA3g2R4n_QvmJvv7BM6bh1qxPvaBWzAWTihA1sGYCAekuXRALCphOxaWCMKBN9QtahK2eIOQLgrlvTsom1WMw-XJYkZOzXZ0z3yrYFfxJCU"/>
      <span v-if="request" class="font-metadata text-metadata">
       Raised by
       <strong class="font-medium text-text-primary">
        {{ request.requestorName }}
       </strong>
       ({{ request.department }})
      </span>
     </div>
     <span class="text-border-default">
      |
     </span>
     <span class="font-metadata text-metadata flex items-center gap-1">
      <span class="material-symbols-outlined" style="font-size: 14px;">
       calendar_today
      </span>
      Submitted Oct 24, 2026
     </span>
    </div>
   </div>
   <div class="flex items-center gap-inline">
    <button class="px-4 py-2 bg-surface text-text-primary border border-border-default rounded-lg font-body text-body font-medium hover:bg-surface-subtle transition-colors flex items-center gap-2 shadow-sm">
     <span class="material-symbols-outlined" style="font-size: 18px;">
      info
     </span>
     Request Info
    </button>
    <button class="px-4 py-2 bg-surface text-status-critical border border-status-critical/30 rounded-lg font-body text-body font-medium hover:bg-error-container hover:border-status-critical transition-colors shadow-sm">
     Reject
    </button>
    <button class="px-5 py-2 bg-text-primary text-white rounded-lg font-body text-body font-medium hover:bg-on-surface-variant transition-colors shadow-sm flex items-center gap-2">
     <span class="material-symbols-outlined" style="font-size: 18px;">
      check
     </span>
     Approve
    </button>
   </div>
  </div>
  <!-- Bento Grid Layout -->
  <div class="grid grid-cols-1 lg:grid-cols-12 gap-stack">
   <!-- Left Column (Wider for main details) -->
   <div class="lg:col-span-8 flex flex-col gap-stack">
    <!-- Request Details Card -->
    <div class="bg-surface rounded-2xl border border-border-default p-card-padding shadow-[0_4px_12px_rgba(0,0,0,0.02)] hover:-translate-y-[2px] transition-transform duration-200">
     <h2 class="font-h2 text-h2 text-text-primary mb-6 flex items-center gap-2">
      <span class="material-symbols-outlined text-primary">
       analytics
      </span>
      Request Details
     </h2>
     <div class="grid grid-cols-2 md:grid-cols-4 gap-4 mb-8">
      <div class="bg-canvas rounded-xl p-4 border border-border-default/50">
       <span class="font-metadata text-metadata text-text-secondary block mb-1">
        Category
       </span>
       <span v-if="request" class="font-h3 text-h3 text-text-primary flex items-center gap-2">
        <span class="material-symbols-outlined text-stone-400" style="font-size: 16px;">
         laptop_mac
        </span>
        {{ request.department }}
       </span>
      </div>
      <div class="bg-metric-sage/30 rounded-xl p-4 border border-metric-sage">
       <span class="font-metadata text-metadata text-status-in-stock block mb-1">
        Priority
       </span>
       <span v-if="request" class="font-kpi-number text-kpi-number text-status-in-stock">
        {{ request.priority }}
       </span>
      </div>
      <div class="bg-canvas rounded-xl p-4 border border-border-default/50">
       <span class="font-metadata text-metadata text-text-secondary block mb-1">
        Est. Unit Cost
       </span>
       <span class="font-mono-data text-mono-data text-text-primary text-base">
        -
       </span>
      </div>
      <div class="bg-metric-lavender/40 rounded-xl p-4 border border-metric-lavender">
       <span class="font-metadata text-metadata text-tertiary block mb-1">
        Total Est. Value
       </span>
       <span v-if="request" class="font-mono-data text-mono-data text-tertiary text-lg font-semibold">
        ₹{{ request.estimatedValue.toLocaleString() }}
       </span>
      </div>
     </div>
     <div>
      <span class="font-table-header text-table-header text-text-secondary block mb-2 uppercase">
       Reason / Justification
      </span>
      <div class="bg-surface-subtle p-4 rounded-lg border border-border-default/50">
       <p class="font-body text-body text-text-primary">
        Required for the upcoming Q1 intake of the Engineering Excellence program. Current inventory of high-performance machines (16GB+ RAM, dedicated GPU) is depleted. These specific models are required to run the local containerized environments specified by the DevSecOps team.
       </p>
      </div>
     </div>
    </div>
    <!-- Vendor Quotes Card -->
    <div class="bg-surface rounded-2xl border border-border-default p-card-padding shadow-[0_4px_12px_rgba(0,0,0,0.02)] hover:-translate-y-[2px] transition-transform duration-200">
     <div class="flex items-center justify-between mb-6">
      <h2 class="font-h2 text-h2 text-text-primary flex items-center gap-2">
       <span class="material-symbols-outlined text-primary">
        request_quote
       </span>
       Vendor Quotes
      </h2>
      <span class="font-metadata text-metadata text-text-secondary bg-surface-subtle px-2 py-1 rounded border border-border-default">
       3 Attached
      </span>
     </div>
     <div class="overflow-x-auto border border-border-default rounded-xl">
      <table class="w-full text-left border-collapse">
       <thead>
        <tr class="bg-surface-subtle border-b border-border-default">
         <th class="font-table-header text-table-header text-text-secondary py-3 px-4">
          Vendor
         </th>
         <th class="font-table-header text-table-header text-text-secondary py-3 px-4">
          Model
         </th>
         <th class="font-table-header text-table-header text-text-secondary py-3 px-4">
          Unit Price
         </th>
         <th class="font-table-header text-table-header text-text-secondary py-3 px-4">
          Lead Time
         </th>
         <th class="font-table-header text-table-header text-text-secondary py-3 px-4 text-right">
          Action
         </th>
        </tr>
       </thead>
       <tbody>
        <tr class="border-b border-border-default hover:bg-metric-amber/20 transition-colors border-l-[3px] border-l-amber-500 bg-amber-50/30">
         <td class="py-3 px-4">
          <div class="flex items-center gap-2">
           <div class="w-6 h-6 rounded bg-stone-200 flex items-center justify-center font-bold text-xs text-stone-600">
            T
           </div>
           <span class="font-body text-body font-medium text-text-primary">
            TechCorp India
           </span>
           <span class="ml-2 inline-flex items-center px-2 py-0.5 rounded text-[10px] font-medium bg-metric-sage text-status-in-stock">
            Recommended
           </span>
          </div>
         </td>
         <td class="py-3 px-4 font-body text-body text-text-secondary">
          ProBook G9 16GB
         </td>
         <td class="py-3 px-4 font-mono-data text-mono-data text-text-primary">
          ₹64,500
         </td>
         <td class="py-3 px-4 font-body text-body text-text-secondary">
          14 Days
         </td>
         <td class="py-3 px-4 text-right">
          <button class="text-tertiary hover:underline font-body text-body font-medium flex items-center justify-end gap-1 w-full">
           <span class="material-symbols-outlined" style="font-size: 16px;">
            download
           </span>
           PDF
          </button>
         </td>
        </tr>
        <tr class="border-b border-border-default hover:bg-surface-subtle transition-colors">
         <td class="py-3 px-4 border-l-[3px] border-l-transparent">
          <div class="flex items-center gap-2">
           <div class="w-6 h-6 rounded bg-stone-200 flex items-center justify-center font-bold text-xs text-stone-600">
            E
           </div>
           <span class="font-body text-body font-medium text-text-primary">
            Enterprise IT Solutions
           </span>
          </div>
         </td>
         <td class="py-3 px-4 font-body text-body text-text-secondary">
          Latitude 5530 16GB
         </td>
         <td class="py-3 px-4 font-mono-data text-mono-data text-text-primary">
          ₹66,200
         </td>
         <td class="py-3 px-4 font-body text-body text-text-secondary">
          21 Days
         </td>
         <td class="py-3 px-4 text-right">
          <button class="text-tertiary hover:underline font-body text-body font-medium flex items-center justify-end gap-1 w-full">
           <span class="material-symbols-outlined" style="font-size: 16px;">
            download
           </span>
           PDF
          </button>
         </td>
        </tr>
        <tr class="hover:bg-surface-subtle transition-colors">
         <td class="py-3 px-4 border-l-[3px] border-l-transparent">
          <div class="flex items-center gap-2">
           <div class="w-6 h-6 rounded bg-stone-200 flex items-center justify-center font-bold text-xs text-stone-600">
            G
           </div>
           <span class="font-body text-body font-medium text-text-primary">
            Global Hardware Pvt
           </span>
          </div>
         </td>
         <td class="py-3 px-4 font-body text-body text-text-secondary">
          ThinkPad L15 16GB
         </td>
         <td class="py-3 px-4 font-mono-data text-mono-data text-text-primary">
          ₹68,000
         </td>
         <td class="py-3 px-4 font-body text-body text-text-secondary">
          10 Days
         </td>
         <td class="py-3 px-4 text-right">
          <button class="text-tertiary hover:underline font-body text-body font-medium flex items-center justify-end gap-1 w-full">
           <span class="material-symbols-outlined" style="font-size: 16px;">
            download
           </span>
           PDF
          </button>
         </td>
        </tr>
       </tbody>
      </table>
     </div>
    </div>
   </div>
   <!-- Right Column (Narrower for sidebar-style info) -->
   <div class="lg:col-span-4 flex flex-col gap-stack">
    <!-- Approval Pipeline Card -->
    <div class="bg-surface rounded-2xl border border-border-default p-card-padding shadow-[0_4px_12px_rgba(0,0,0,0.02)] hover:-translate-y-[2px] transition-transform duration-200">
     <h2 class="font-h2 text-h2 text-text-primary mb-6 flex items-center gap-2">
      <span class="material-symbols-outlined text-primary">
       account_tree
      </span>
      Approval Pipeline
     </h2>
     <div class="relative pl-6 border-l-2 border-stone-100 dark:border-stone-800 ml-3 space-y-6">
      <!-- Step 1: Completed -->
      <div class="relative">
       <div class="absolute -left-[35px] top-0 w-6 h-6 rounded-full bg-status-in-stock text-white flex items-center justify-center border-4 border-surface z-10">
        <span class="material-symbols-outlined" style="font-size: 12px; font-weight: bold;">
         check
        </span>
       </div>
       <h3 class="font-h3 text-h3 text-text-primary">
        Draft Submitted
       </h3>
       <p class="font-metadata text-metadata text-text-secondary mt-1">
        Sarah Jenkins • Oct 24, 09:15 AM
       </p>
      </div>
      <!-- Step 2: Completed -->
      <div class="relative">
       <div class="absolute -left-[35px] top-0 w-6 h-6 rounded-full bg-status-in-stock text-white flex items-center justify-center border-4 border-surface z-10">
        <span class="material-symbols-outlined" style="font-size: 12px; font-weight: bold;">
         check
        </span>
       </div>
       <h3 class="font-h3 text-h3 text-text-primary">
        Dept Head Approved
       </h3>
       <p class="font-metadata text-metadata text-text-secondary mt-1">
        Michael Chang • Oct 25, 14:30 PM
       </p>
       <div class="mt-2 p-2 bg-surface-subtle rounded border border-border-default/50 text-sm italic text-text-secondary flex gap-2">
        <span class="material-symbols-outlined text-stone-400" style="font-size: 16px;">
         format_quote
        </span>
        <span>
         Approved. Essential for Q1 intake.
        </span>
       </div>
      </div>
      <!-- Step 3: Current -->
      <div class="relative">
       <div class="absolute -left-[35px] top-0 w-6 h-6 rounded-full bg-metric-amber border-2 border-amber-500 flex items-center justify-center border-4 border-surface z-10">
        <div class="w-2 h-2 rounded-full bg-amber-500">
        </div>
       </div>
       <h3 class="font-h3 text-h3 text-text-primary font-bold text-amber-700">
        Pending Finance
       </h3>
       <p class="font-metadata text-metadata text-text-secondary mt-1">
        Awaiting your action
       </p>
      </div>
      <!-- Step 4: Future -->
      <div class="relative">
       <div class="absolute -left-[35px] top-0 w-6 h-6 rounded-full bg-surface border-2 border-border-default flex items-center justify-center border-4 border-surface z-10">
       </div>
       <h3 class="font-h3 text-h3 text-text-secondary">
        Procurement Execution
       </h3>
       <p class="font-metadata text-metadata text-border-default mt-1">
        Not started
       </p>
      </div>
     </div>
    </div>
    <!-- Comments Card -->
    <div class="bg-surface rounded-2xl border border-border-default p-card-padding shadow-[0_4px_12px_rgba(0,0,0,0.02)] flex-1 hover:-translate-y-[2px] transition-transform duration-200">
     <h2 class="font-h2 text-h2 text-text-primary mb-4 flex items-center gap-2">
      <span class="material-symbols-outlined text-primary">
       forum
      </span>
      Comments
     </h2>
     <div class="space-y-4 mb-6">
      <div class="flex gap-3">
       <img alt="Sarah Jenkins" class="w-8 h-8 rounded-full border border-border-default mt-1" data-alt="professional headshot of a woman with dark hair smiling gently" src="https://lh3.googleusercontent.com/aida-public/AB6AXuCvnAJhDv7ibhMM8BhSkMmGtGTXMGAbVJFI-7RgBwMxNA4IUOmyMqkYliunGvzQwdomGOXs1ogRjfyCME0w8RsIMH6yk7DSGEAhTJezWLHidLNTeQL_n_8ztJjkRukdqQqni_e7-yumU1qQH1oKK00hitiR6_h1hHSJwGbnuptUT2fy5BBvjU8ZSEGTm-y8l7o9RetKzK_yT8CCRzh9yvZKH69dM9K4ApACRPQCgnlBIrUDVugJ3Q43n_uA8yiI15rgcYsRwHMdLXsL"/>
       <div class="bg-surface-subtle p-3 rounded-xl rounded-tl-none border border-border-default/50">
        <div class="flex justify-between items-baseline mb-1 gap-4">
         <span class="font-h3 text-h3 text-text-primary">
          Sarah Jenkins
         </span>
         <span class="font-metadata text-metadata text-text-secondary text-[10px]">
          Oct 24, 09:20 AM
         </span>
        </div>
        <p class="font-body text-body text-text-secondary text-sm">
         I've attached three quotes as requested in the new policy. TechCorp is slightly cheaper but Enterprise IT has better warranty terms. Open to guidance here.
        </p>
       </div>
      </div>
      <div class="flex gap-3">
       <div class="w-8 h-8 rounded-full bg-tertiary text-white flex items-center justify-center font-bold text-sm mt-1 shrink-0">
        MC
       </div>
       <div class="bg-surface-subtle p-3 rounded-xl rounded-tl-none border border-border-default/50">
        <div class="flex justify-between items-baseline mb-1 gap-4">
         <span class="font-h3 text-h3 text-text-primary">
          Michael Chang
         </span>
         <span class="font-metadata text-metadata text-text-secondary text-[10px]">
          Oct 25, 14:25 PM
         </span>
        </div>
        <p class="font-body text-body text-text-secondary text-sm">
         TechCorp's lead time is better for our Q1 start date. Let's proceed with them despite the standard warranty.
        </p>
       </div>
      </div>
     </div>
     <div class="mt-auto">
      <div class="relative">
       <textarea class="w-full bg-canvas border border-border-default rounded-xl p-3 pl-3 pr-10 font-body text-body focus:ring-2 focus:ring-amber-500 focus:border-amber-500 resize-none h-[80px]" placeholder="Add a note or request more info..."></textarea>
       <button class="absolute right-2 bottom-2 p-1.5 bg-text-primary text-white rounded-lg hover:bg-on-surface-variant transition-colors flex items-center justify-center">
        <span class="material-symbols-outlined" style="font-size: 16px;">
         send
        </span>
       </button>
      </div>
     </div>
    </div>
   </div>
  </div>
 </div>
</main>

</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'

interface ProcurementRequest {
  id: number;
  title: string;
  status: string;
  priority: string;
  estimatedValue: number;
  actualValue: number;
  requestorName: string;
  requestorInitials: string;
  department: string;
  poNumber: string;
  createdAt: string;
}

const request = ref<ProcurementRequest | null>(null)
const loading = ref(true)

const fetchRequest = async () => {
  try {
    const res = await fetch('http://localhost:8080/api/procurements')
    if (res.ok) {
      const allRequests: ProcurementRequest[] = await res.json()
      // Fallback to first if no ID route param provided in this demo
      request.value = (allRequests.length > 0 ? allRequests[0] : null) ?? null
    }
  } catch (error) {
    console.error('Failed to fetch procurements:', error)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchRequest()
})
</script>
