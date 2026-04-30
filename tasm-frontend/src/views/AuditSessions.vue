<template>
<main class="flex-1 mt-[60px] p-page-margin max-w-[1400px]">
 <!-- Page Header & Actions -->
 <div class="flex justify-between items-end mb-section-gap">
  <div>
   <h1 class="font-h1 text-h1 text-text-primary mb-1">
    Audit Sessions
   </h1>
   <p class="font-metadata text-metadata text-text-secondary">
    Manage and track physical asset verifications across facilities.
   </p>
  </div>
  <button class="bg-[#1C1917] hover:bg-stone-800 text-white rounded font-body text-body py-2 px-5 transition-colors shadow-sm flex items-center gap-2">
   <span class="material-symbols-outlined" style="font-size: 18px;">
    qr_code_scanner
   </span>
   Start New Audit Session
  </button>
 </div>
 <!-- Section 1: Active/Past Sessions Table -->
 <section class="bg-surface border border-border-default rounded-xl shadow-[0_4px_4px_rgba(0,0,0,0.02)] mb-section-gap overflow-hidden">
  <div class="px-card-padding py-4 border-b border-border-default bg-surface-subtle flex justify-between items-center">
   <h2 class="font-h2 text-h2 text-text-primary">
    Recent Audit Sessions
   </h2>
   <button class="text-text-secondary hover:text-primary transition-colors">
    <span class="material-symbols-outlined">
     filter_list
    </span>
   </button>
  </div>
  <div class="overflow-x-auto border-t border-border-default">
   <DataTable :value="audits" :loading="loadingAudits" paginator :rows="5" tableStyle="min-width: 50rem" class="w-full">
    <Column field="sessionId" header="Session ID" sortable>
     <template #body="slotProps">
      <span class="font-mono-data text-mono-data text-text-primary border-l-[3px]" :class="slotProps.data.status === 'In Progress' ? 'border-amber-600 pl-2' : 'border-transparent pl-2'">
       {{ slotProps.data.sessionId }}
      </span>
     </template>
    </Column>
    <Column field="location" header="Scope" sortable>
     <template #body="slotProps">
      <span class="text-text-primary">{{ slotProps.data.location }}</span>
     </template>
    </Column>
    <Column field="createdAt" header="Date" sortable>
     <template #body="slotProps">
      <span class="text-text-secondary">{{ new Date(slotProps.data.createdAt).toLocaleDateString() }}</span>
     </template>
    </Column>
    <Column field="verifiedPercent" header="Verified %" sortable>
     <template #body="slotProps">
      <div class="flex items-center gap-2">
       <span class="font-medium" :class="slotProps.data.verifiedPercent === 100 ? 'text-status-in-stock' : ''">
        {{ slotProps.data.verifiedPercent }}%
       </span>
       <div class="w-16 h-1.5 bg-surface-variant rounded-full overflow-hidden">
        <div class="h-full rounded-full" :class="slotProps.data.verifiedPercent === 100 ? 'bg-status-in-stock' : 'bg-amber-500'" :style="'width: ' + slotProps.data.verifiedPercent + '%'">
        </div>
       </div>
      </div>
     </template>
    </Column>
    <Column field="missingCount" header="Missing" sortable>
     <template #body="slotProps">
      <span :class="slotProps.data.missingCount > 0 ? 'text-status-critical font-medium' : 'text-text-secondary'">
       {{ slotProps.data.missingCount }}
      </span>
     </template>
    </Column>
    <Column field="unregisteredCount" header="Unregistered" sortable>
     <template #body="slotProps">
      <span :class="slotProps.data.unregisteredCount > 0 ? 'text-primary-container' : 'text-text-secondary'">
       {{ slotProps.data.unregisteredCount }}
      </span>
     </template>
    </Column>
    <Column field="status" header="Status" sortable>
     <template #body="slotProps">
      <span v-if="slotProps.data.status === 'In Progress'" class="inline-flex items-center px-2 py-1 rounded bg-metric-lavender text-tertiary-container font-metadata text-metadata font-medium">
       In Progress
      </span>
      <span v-else class="inline-flex items-center px-2 py-1 rounded bg-metric-sage text-status-in-stock font-metadata text-metadata font-medium">
       Completed
      </span>
     </template>
    </Column>
    <Column header="Action" alignFrozen="right">
     <template #body>
      <button class="text-primary hover:text-amber-800 font-medium transition-colors">
       View Report
      </button>
     </template>
    </Column>
   </DataTable>
  </div>
  <div class="px-card-padding py-3 border-t border-border-default bg-surface flex justify-end">
   <button class="text-text-secondary hover:text-text-primary font-metadata text-metadata transition-colors flex items-center gap-1">
    View All Sessions
    <span class="material-symbols-outlined" style="font-size: 16px;">
     arrow_forward
    </span>
   </button>
  </div>
 </section>
 <!-- Section 2: Discrepancy Resolution Queue (Bento/Card layout) -->
 <h2 class="font-h2 text-h2 text-text-primary mb-4">
  Discrepancy Resolution Queue
 </h2>
 <div class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-6">
  <!-- Discrepancy Card 1 -->
  <div class="bg-surface border border-border-default rounded-xl p-card-padding shadow-[0_4px_4px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform duration-200 flex flex-col h-full">
   <div class="flex justify-between items-start mb-4">
    <div class="flex items-center gap-2">
     <span class="material-symbols-outlined text-status-critical bg-error-container p-1.5 rounded-lg">
      warning
     </span>
     <span class="font-metadata text-metadata text-status-critical font-medium uppercase tracking-wider">
      Missing
     </span>
    </div>
    <span class="text-status-critical font-metadata text-metadata font-bold bg-error-container px-2 py-1 rounded">
     4 Days
    </span>
   </div>
   <h3 class="font-h3 text-h3 text-text-primary mb-1">
    Dell Latitude 5420
   </h3>
   <p class="font-mono-data text-mono-data text-text-secondary mb-4">
    AST-LPT-8892
   </p>
   <div class="bg-surface-subtle p-3 rounded border border-border-default mb-4 flex-1">
    <p class="font-metadata text-metadata text-text-secondary mb-1">
     Flagged In:
    </p>
    <p class="font-body text-body text-text-primary">
     AUD-2023-11-A (Gayathri - Floor 2)
    </p>
   </div>
   <div class="mt-auto pt-4 border-t border-border-default flex justify-end gap-2">
    <select class="form-select w-full bg-surface border border-[#D6D3CE] rounded text-body font-body text-text-primary py-1.5 px-3 focus:ring-primary focus:border-primary">
     <option disabled selected value="">
      Select Action...
     </option>
     <option value="confirm">
      Confirm Location
     </option>
     <option value="lost">
      Mark as Lost
     </option>
     <option value="reassign">
      Reassign Asset
     </option>
    </select>
    <button class="bg-surface border border-[#D6D3CE] hover:bg-surface-subtle text-text-primary rounded p-1.5 transition-colors flex-shrink-0">
     <span class="material-symbols-outlined">
      check
     </span>
    </button>
   </div>
  </div>
  <!-- Discrepancy Card 2 -->
  <div class="bg-surface border border-border-default rounded-xl p-card-padding shadow-[0_4px_4px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform duration-200 flex flex-col h-full">
   <div class="flex justify-between items-start mb-4">
    <div class="flex items-center gap-2">
     <span class="material-symbols-outlined text-primary-container bg-accent-subtle p-1.5 rounded-lg">
      location_off
     </span>
     <span class="font-metadata text-metadata text-primary-container font-medium uppercase tracking-wider">
      Location Mismatch
     </span>
    </div>
    <span class="text-primary-container font-metadata text-metadata font-bold bg-accent-subtle px-2 py-1 rounded">
     2 Days
    </span>
   </div>
   <h3 class="font-h3 text-h3 text-text-primary mb-1">
    Herman Miller Aeron
   </h3>
   <p class="font-mono-data text-mono-data text-text-secondary mb-4">
    AST-FUR-1104
   </p>
   <div class="bg-surface-subtle p-3 rounded border border-border-default mb-4 flex-1">
    <p class="font-metadata text-metadata text-text-secondary mb-1">
     Flagged In:
    </p>
    <p class="font-body text-body text-text-primary">
     AUD-2023-11-A (Gayathri - Floor 2)
    </p>
    <p class="font-metadata text-metadata text-text-secondary mt-2">
     Expected: Nila - Floor 4
    </p>
   </div>
   <div class="mt-auto pt-4 border-t border-border-default flex justify-end gap-2">
    <select class="form-select w-full bg-surface border border-[#D6D3CE] rounded text-body font-body text-text-primary py-1.5 px-3 focus:ring-primary focus:border-primary">
     <option disabled selected value="">
      Select Action...
     </option>
     <option value="confirm">
      Update Location to Gayathri
     </option>
     <option value="lost">
      Mark for Return to Nila
     </option>
    </select>
    <button class="bg-surface border border-[#D6D3CE] hover:bg-surface-subtle text-text-primary rounded p-1.5 transition-colors flex-shrink-0">
     <span class="material-symbols-outlined">
      check
     </span>
    </button>
   </div>
  </div>
  <!-- Discrepancy Card 3 (Unregistered) -->
  <div class="bg-surface border border-border-default rounded-xl p-card-padding shadow-[0_4px_4px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform duration-200 flex flex-col h-full">
   <div class="flex justify-between items-start mb-4">
    <div class="flex items-center gap-2">
     <span class="material-symbols-outlined text-tertiary bg-tertiary-fixed p-1.5 rounded-lg">
      help_center
     </span>
     <span class="font-metadata text-metadata text-tertiary font-medium uppercase tracking-wider">
      Unregistered
     </span>
    </div>
    <span class="text-tertiary font-metadata text-metadata font-bold bg-tertiary-fixed px-2 py-1 rounded">
     1 Day
    </span>
   </div>
   <h3 class="font-h3 text-h3 text-text-primary mb-1">
    Unknown Cisco Switch
   </h3>
   <p class="font-mono-data text-mono-data text-text-secondary mb-4">
    No Asset Tag
   </p>
   <div class="bg-surface-subtle p-3 rounded border border-border-default mb-4 flex-1">
    <p class="font-metadata text-metadata text-text-secondary mb-1">
     Flagged In:
    </p>
    <p class="font-body text-body text-text-primary">
     AUD-2023-11-A (Gayathri - Floor 2)
    </p>
    <p class="font-metadata text-metadata text-text-secondary mt-2">
     Notes: Found in network closet G2-B.
    </p>
   </div>
   <div class="mt-auto pt-4 border-t border-border-default flex justify-end gap-2">
    <select class="form-select w-full bg-surface border border-[#D6D3CE] rounded text-body font-body text-text-primary py-1.5 px-3 focus:ring-primary focus:border-primary">
     <option disabled selected value="">
      Select Action...
     </option>
     <option value="register">
      Register New Asset
     </option>
     <option value="ignore">
      Ignore/Not Asset
     </option>
    </select>
    <button class="bg-surface border border-[#D6D3CE] hover:bg-surface-subtle text-text-primary rounded p-1.5 transition-colors flex-shrink-0">
     <span class="material-symbols-outlined">
      check
     </span>
    </button>
   </div>
  </div>
 </div>
</main>

</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'

interface AuditSession {
  id: number;
  sessionId: string;
  title: string;
  status: string;
  auditor: string;
  location: string;
  verifiedPercent: number;
  missingCount: number;
  unregisteredCount: number;
  createdAt: string;
}

const audits = ref<AuditSession[]>([])
const loadingAudits = ref(true)

const fetchAudits = async () => {
  try {
    const res = await fetch('http://localhost:8080/api/audits')
    if (res.ok) {
      audits.value = await res.json()
    }
  } catch (error) {
    console.error('Failed to fetch audits:', error)
  } finally {
    loadingAudits.value = false
  }
}

onMounted(() => {
  fetchAudits()
})
</script>
