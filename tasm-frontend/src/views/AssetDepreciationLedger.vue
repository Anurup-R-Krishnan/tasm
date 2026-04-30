<template>
<main class="flex-1 flex flex-col min-h-screen bg-canvas">
 <!-- TopNavBar -->
 <header class="bg-[#FAF9F6] dark:bg-stone-950 docked full-width top-0 h-[60px] border-b border-[#EEEBE4] dark:border-stone-800 shadow-sm flex justify-between items-center px-8 w-full sticky top-0 z-40">
  <!-- Left Side: Brand/Search -->
  <div class="flex items-center gap-6">
   <div class="text-lg font-bold text-stone-900 dark:text-stone-50 tracking-tight">
    Technopark AMS
   </div>
   <!-- Search Bar -->
   <div class="relative w-64">
    <span class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-stone-400 text-[18px]">
     search
    </span>
    <input class="w-full pl-9 pr-4 py-1.5 bg-surface border border-[#EEEBE4] rounded-md font-['Inter'] text-sm tracking-tight text-stone-900 focus:outline-none focus:border-[#92400E] focus:ring-1 focus:ring-[#92400E] transition-colors" placeholder="Search assets, records..." type="text"/>
   </div>
  </div>
  <!-- Right Side: Actions -->
  <div class="flex items-center gap-4">
   <button class="font-['Inter'] text-sm tracking-tight text-[#92400E] dark:text-amber-500 font-medium hover:bg-stone-100 dark:hover:bg-stone-900 transition-colors px-3 py-1.5 rounded-md cursor-pointer active:opacity-80 transition-opacity">
    Audit Logs
   </button>
   <div class="flex items-center gap-1 border-l border-[#EEEBE4] pl-4 ml-2">
    <button class="w-8 h-8 flex items-center justify-center text-stone-600 dark:text-stone-400 hover:bg-stone-100 dark:hover:bg-stone-900 transition-colors rounded-full cursor-pointer active:opacity-80 transition-opacity">
     <span class="material-symbols-outlined text-[20px]">
      notifications
     </span>
    </button>
    <button class="w-8 h-8 flex items-center justify-center text-stone-600 dark:text-stone-400 hover:bg-stone-100 dark:hover:bg-stone-900 transition-colors rounded-full cursor-pointer active:opacity-80 transition-opacity">
     <span class="material-symbols-outlined text-[20px]">
      help
     </span>
    </button>
    <button class="w-8 h-8 flex items-center justify-center text-stone-600 dark:text-stone-400 hover:bg-stone-100 dark:hover:bg-stone-900 transition-colors rounded-full cursor-pointer active:opacity-80 transition-opacity">
     <span class="material-symbols-outlined text-[20px]">
      settings
     </span>
    </button>
    <button class="ml-2 w-8 h-8 rounded-full bg-stone-200 border border-stone-300 overflow-hidden cursor-pointer active:opacity-80 transition-opacity">
     <img alt="Administrator profile picture" class="w-full h-full object-cover" data-alt="Default administrator avatar with initials" src="https://lh3.googleusercontent.com/aida-public/AB6AXuBvJxl17Wuzfzja-KV7WmvViZpvXWRD0RhGWxOK4p5uCV5AGrDKXTAmjOxZDNHLSbnEydjpgU7aRBai9lVr_tWIul5zgHSr_lqz3yRNawomCxBO1xE-ySDb3LB_pLT9kKhLQ-N2ozG4RZLY8XJuqDgcKH2os9BacffKIxCo1J9ZzbFJzDTSCCi3cq25oB80Z9SmHY2lnTuo6gRVvf9KJBbrblLY4V0Q_-THuOntdC9b-plLtRtDqNcT7FWJM2imVK7bpuqBUi8VpM0o"/>
    </button>
   </div>
  </div>
 </header>
 <!-- Scrollable Content Canvas -->
 <div class="flex-1 overflow-y-auto p-page-margin">
  <div class="max-w-[1400px] mx-auto space-y-section-gap">
   <!-- Page Header: Asset Identity -->
   <div class="flex flex-col md:flex-row md:items-end justify-between gap-4">
    <div>
     <div class="flex items-center gap-2 mb-2">
      <span class="font-metadata text-metadata text-text-secondary">
       Financial Ledgers
      </span>
      <span class="material-symbols-outlined text-[14px] text-text-secondary">
       chevron_right
      </span>
      <span class="font-metadata text-metadata text-primary-container font-medium">
       Asset Depreciation
      </span>
     </div>
     <h2 class="font-h1 text-h1 text-text-primary mb-1">
      Laptop Alpha #TAG-1001
     </h2>
     <div class="flex items-center gap-3">
      <span class="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium bg-surface-variant text-on-surface-variant border border-outline-variant">
       Category: IT Equipment
      </span>
      <span class="inline-flex items-center gap-1 text-xs font-medium text-status-in-stock">
       <span class="w-2 h-2 rounded-full bg-status-in-stock">
       </span>
       Active Asset
      </span>
     </div>
    </div>
    <div class="flex items-center gap-3">
     <button class="px-4 py-2 bg-surface text-text-primary border border-border-default rounded-lg font-h3 text-h3 hover:bg-surface-subtle transition-colors flex items-center gap-2 shadow-sm -translate-y-0 hover:-translate-y-[1px]">
      <span class="material-symbols-outlined text-[18px]">
       print
      </span>
      Print
     </button>
     <button class="px-4 py-2 bg-[#1C1917] text-white rounded-lg font-h3 text-h3 hover:bg-stone-800 transition-colors flex items-center gap-2 shadow-sm -translate-y-0 hover:-translate-y-[1px]">
      <span class="material-symbols-outlined text-[18px]">
       download
      </span>
      Export CSV
     </button>
    </div>
   </div>
   <!-- KPI Cards Grid -->
   <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
    <!-- KPI 1 -->
    <div class="bg-surface border border-border-default rounded-xl p-card-padding shadow-sm hover:-translate-y-[2px] transition-transform duration-200">
     <div class="flex items-center justify-between mb-4">
      <span class="font-metadata text-metadata text-text-secondary uppercase tracking-wider">
       Total Cost
      </span>
      <span class="material-symbols-outlined text-text-secondary bg-surface-subtle p-1.5 rounded-md">
       payments
      </span>
     </div>
     <div class="font-kpi-number text-kpi-number text-text-primary">
      ₹1,20,000
     </div>
     <div class="font-metadata text-metadata text-text-secondary mt-1">
      Initial acquisition value
     </div>
    </div>
    <!-- KPI 2 -->
    <div class="bg-surface border border-border-default rounded-xl p-card-padding shadow-sm hover:-translate-y-[2px] transition-transform duration-200">
     <div class="flex items-center justify-between mb-4">
      <span class="font-metadata text-metadata text-text-secondary uppercase tracking-wider">
       Accumulated Depreciation
      </span>
      <span class="material-symbols-outlined text-text-secondary bg-surface-subtle p-1.5 rounded-md">
       trending_down
      </span>
     </div>
     <div class="font-kpi-number text-kpi-number text-error">
      ₹45,000
     </div>
     <div class="font-metadata text-metadata text-text-secondary mt-1">
      Total written off to date
     </div>
    </div>
    <!-- KPI 3 (Sage Fill) -->
    <div class="bg-metric-sage border border-status-in-stock/20 rounded-xl p-card-padding shadow-sm hover:-translate-y-[2px] transition-transform duration-200 relative overflow-hidden">
     <div class="absolute right-0 top-0 opacity-10 pointer-events-none">
      <span class="material-symbols-outlined text-[120px] -mt-4 -mr-4" style="font-variation-settings: 'FILL' 1;">
       account_balance
      </span>
     </div>
     <div class="flex items-center justify-between mb-4 relative z-10">
      <span class="font-metadata text-metadata text-status-in-stock uppercase tracking-wider font-semibold">
       Net Book Value
      </span>
      <span class="material-symbols-outlined text-status-in-stock bg-white/50 p-1.5 rounded-md">
       account_balance_wallet
      </span>
     </div>
     <div class="font-kpi-number text-kpi-number text-status-in-stock relative z-10">
      ₹75,000
     </div>
     <div class="font-metadata text-metadata text-status-in-stock/80 mt-1 relative z-10">
      Current carrying value
     </div>
    </div>
   </div>
   <!-- Dense Table Section -->
   <div class="bg-surface border border-border-default rounded-xl shadow-sm overflow-hidden flex flex-col">
    <div class="px-card-padding py-4 border-b border-border-default bg-surface flex items-center justify-between">
     <h3 class="font-h2 text-h2 text-text-primary flex items-center gap-2">
      <span class="material-symbols-outlined text-primary-container">
       calendar_month
      </span>
      Depreciation Schedules
     </h3>
    </div>
    <div class="overflow-x-auto p-4">
     <DataTable :value="schedules" :loading="loading" paginator :rows="10" tableStyle="min-width: 50rem" class="w-full text-left">
      <Column field="assetId" header="Asset ID" sortable>
       <template #body="slotProps">
        <span class="font-mono-data text-mono-data text-surface-tint">
         {{ slotProps.data.assetId }}
        </span>
       </template>
      </Column>
      <Column field="assetName" header="Asset Name" sortable></Column>
      <Column field="purchaseValue" header="Purchase Value" sortable>
       <template #body="slotProps">
        <span class="font-mono-data text-mono-data">
         ₹ {{ slotProps.data.purchaseValue?.toLocaleString() }}
        </span>
       </template>
      </Column>
      <Column field="currentValue" header="Current Value" sortable>
       <template #body="slotProps">
        <span class="font-mono-data text-mono-data text-primary-container">
         ₹ {{ slotProps.data.currentValue?.toLocaleString() }}
        </span>
       </template>
      </Column>
      <Column field="method" header="Method" sortable>
       <template #body="slotProps">
        <span class="inline-flex items-center px-2 py-0.5 rounded text-[11px] font-medium bg-surface-variant text-on-surface-variant border border-outline-variant">
         {{ slotProps.data.method }}
        </span>
       </template>
      </Column>
     </DataTable>
    </div>
   </div>
  </div>
 </div>
</main>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'

interface DepreciationSchedule {
  id: number;
  assetId: string;
  assetName: string;
  purchaseValue: number;
  currentValue: number;
  method: string;
}

const schedules = ref<DepreciationSchedule[]>([])
const loading = ref(true)

const fetchSchedules = async () => {
  try {
    const res = await fetch('http://localhost:8080/api/depreciations')
    if (res.ok) {
      schedules.value = await res.json()
    }
  } catch (error) {
    console.error('Failed to fetch depreciation schedules:', error)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchSchedules()
})
</script>
