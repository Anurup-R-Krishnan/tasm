<template>
<main class="flex-1 mt-[60px] p-page-margin max-w-[1400px] mx-auto w-full">
 <!-- Page Header -->
 <div class="flex flex-col md:flex-row md:items-center justify-between gap-4 mb-section-gap">
  <div>
   <div class="flex items-center gap-3 mb-1">
    <a class="text-text-secondary hover:text-primary-container transition-colors" href="#">
     <span class="material-symbols-outlined text-sm">
      arrow_back
     </span>
    </a>
    <span class="font-metadata text-metadata text-text-secondary">
     Financials / Leases
    </span>
   </div>
   <div class="flex items-center gap-3">
    <h1 class="font-h1 text-h1 text-text-primary tracking-tight">
     Lease Agreement: LS-4055 - Tata Elxsi
    </h1>
    <span class="inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full bg-metric-sage text-status-in-stock font-metadata text-metadata border border-[#86efac]/50">
     <span class="w-1.5 h-1.5 rounded-full bg-status-in-stock">
     </span>
     Active
    </span>
   </div>
  </div>
  <div class="flex items-center gap-3">
   <button class="px-4 py-2 bg-surface border border-border-default rounded-lg font-h3 text-h3 text-text-primary hover:bg-surface-subtle transition-colors shadow-sm active:translate-y-0.5 flex items-center gap-2">
    <span class="material-symbols-outlined text-sm">
     edit
    </span>
    Edit
   </button>
   <button class="px-4 py-2 bg-surface border border-border-default rounded-lg font-h3 text-h3 text-text-primary hover:bg-surface-subtle transition-colors shadow-sm active:translate-y-0.5 flex items-center gap-2">
    <span class="material-symbols-outlined text-sm">
     download
    </span>
    Export PDF
   </button>
  </div>
 </div>
 <!-- Top Row: Metrics -->
 <div class="grid grid-cols-1 md:grid-cols-3 gap-stack mb-section-gap">
  <!-- Metric 1 -->
  <div class="bg-metric-amber rounded-[16px] p-card-padding border border-[#fde68a] shadow-sm hover:-translate-y-0.5 transition-transform cursor-pointer">
   <div class="flex justify-between items-start mb-2">
    <span class="font-metadata text-metadata text-amber-900 uppercase tracking-wider">
     Monthly Value
    </span>
    <span class="material-symbols-outlined text-amber-700">
     payments
    </span>
   </div>
   <div class="font-kpi-number text-kpi-number text-amber-950 mb-1">
    ₹4,50,000
   </div>
   <div class="font-metadata text-metadata text-amber-800">
    Due 5th of every month
   </div>
  </div>
  <!-- Metric 2 -->
  <div class="bg-metric-lavender rounded-[16px] p-card-padding border border-[#e9d5ff] shadow-sm hover:-translate-y-0.5 transition-transform cursor-pointer">
   <div class="flex justify-between items-start mb-2">
    <span class="font-metadata text-metadata text-violet-900 uppercase tracking-wider">
     Lease Duration
    </span>
    <span class="material-symbols-outlined text-violet-700">
     hourglass_empty
    </span>
   </div>
   <div class="font-kpi-number text-kpi-number text-violet-950 mb-1">
    36 Months
   </div>
   <div class="font-metadata text-metadata text-violet-800">
    Commenced: 01-Jan-2023
   </div>
  </div>
  <!-- Metric 3 -->
  <div class="bg-surface rounded-[16px] p-card-padding border border-border-default shadow-sm hover:-translate-y-0.5 transition-transform cursor-pointer">
   <div class="flex justify-between items-start mb-2">
    <span class="font-metadata text-metadata text-text-secondary uppercase tracking-wider">
     Expiry Date
    </span>
    <span class="material-symbols-outlined text-text-secondary">
     calendar_today
    </span>
   </div>
   <div class="font-kpi-number text-kpi-number text-text-primary mb-1">
    31-Dec-2025
   </div>
   <div class="font-metadata text-metadata text-text-secondary">
    712 days remaining
   </div>
  </div>
 </div>
  <!-- Leased Assets (now Lease Agreements List) -->
  <div class="lg:col-span-12 bg-surface rounded-[16px] border border-border-default shadow-sm overflow-hidden flex flex-col">
   <div class="p-4 border-b border-border-default bg-surface-subtle flex justify-between items-center">
    <h2 class="font-h2 text-h2 text-text-primary flex items-center gap-2">
     <span class="material-symbols-outlined text-lg">
      gavel
     </span>
     Lease Agreements
    </h2>
   </div>
   <div class="overflow-x-auto p-4">
    <DataTable :value="leases" :loading="loading" paginator :rows="10" tableStyle="min-width: 50rem" class="w-full text-left">
     <Column field="leaseId" header="Lease ID" sortable>
      <template #body="slotProps">
       <span class="font-mono-data text-mono-data text-surface-tint">
        {{ slotProps.data.leaseId }}
       </span>
      </template>
     </Column>
     <Column field="vendor" header="Vendor" sortable></Column>
     <Column field="startDate" header="Start Date" sortable>
      <template #body="slotProps">
       <span>{{ new Date(slotProps.data.startDate).toLocaleDateString() }}</span>
      </template>
     </Column>
     <Column field="endDate" header="End Date" sortable>
      <template #body="slotProps">
       <span>{{ new Date(slotProps.data.endDate).toLocaleDateString() }}</span>
      </template>
     </Column>
     <Column field="monthlyCost" header="Monthly Cost" sortable>
      <template #body="slotProps">
       <span class="font-mono-data text-mono-data text-text-primary">
        ₹ {{ slotProps.data.monthlyCost?.toLocaleString() }}
       </span>
      </template>
     </Column>
     <Column field="status" header="Status" sortable>
      <template #body="slotProps">
       <span class="inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full font-metadata text-metadata"
             :class="slotProps.data.status === 'Active' ? 'bg-metric-sage text-status-in-stock border border-[#86efac]/50' : 'bg-surface-variant text-text-secondary'">
        {{ slotProps.data.status }}
       </span>
      </template>
     </Column>
    </DataTable>
   </div>
  </div>
 </div>
</main>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'

interface LeaseAgreement {
  id: number;
  leaseId: string;
  vendor: string;
  startDate: string;
  endDate: string;
  monthlyCost: number;
  status: string;
}

const leases = ref<LeaseAgreement[]>([])
const loading = ref(true)

const fetchLeases = async () => {
  try {
    const res = await fetch('http://localhost:8080/api/leases')
    if (res.ok) {
      leases.value = await res.json()
    }
  } catch (error) {
    console.error('Failed to fetch leases:', error)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchLeases()
})
</script>
