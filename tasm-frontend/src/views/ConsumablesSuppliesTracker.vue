<template>
<main class="p-page-margin max-w-[1400px] mx-auto space-y-section-gap">
 <!-- Header Section -->
 <div class="flex justify-between items-end">
  <div>
   <h1 class="font-h1 text-h1 text-text-primary mb-1">
    Consumables Tracker
   </h1>
   <p class="font-body text-body text-text-secondary">
    Manage and track daily supplies and consumable inventory.
   </p>
  </div>
  <div class="flex gap-inline">
   <button class="px-4 py-2 bg-surface text-text-primary border border-[#D6D3CE] rounded-lg font-h3 text-h3 hover:bg-surface-subtle transition-colors shadow-sm flex items-center gap-2">
    <span class="material-symbols-outlined text-[18px]">
     download
    </span>
    Receive Stock
   </button>
   <button class="px-4 py-2 bg-[#1C1917] text-white rounded-lg font-h3 text-h3 hover:bg-opacity-90 transition-colors shadow-sm flex items-center gap-2">
    <span class="material-symbols-outlined text-[18px]">
     output
    </span>
    Issue Consumable
   </button>
  </div>
 </div>
 <!-- KPI Cards -->
 <div class="grid grid-cols-4 gap-inline">
  <div class="bg-metric-amber rounded-xl p-card-padding shadow-[0_4px_16px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform cursor-pointer border border-[#EEEBE4]">
   <h3 class="font-metadata text-metadata text-on-surface-variant uppercase mb-2">
    Total Items
   </h3>
   <div class="font-kpi-number text-kpi-number text-text-primary">
    1,248
   </div>
   <div class="font-metadata text-metadata text-text-secondary mt-1">
    Across 12 categories
   </div>
  </div>
  <div class="bg-surface-container-high rounded-xl p-card-padding shadow-[0_4px_16px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform cursor-pointer border border-[#EEEBE4]">
   <h3 class="font-metadata text-metadata text-on-surface-variant uppercase mb-2">
    Low Stock Items
   </h3>
   <div class="font-kpi-number text-kpi-number text-primary-container">
    34
   </div>
   <div class="font-metadata text-metadata text-text-secondary mt-1 flex items-center gap-1">
    <span class="material-symbols-outlined text-[14px] text-primary-container">
     warning
    </span>
    Requires attention
   </div>
  </div>
  <div class="bg-metric-lavender rounded-xl p-card-padding shadow-[0_4px_16px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform cursor-pointer border border-[#EEEBE4]">
   <h3 class="font-metadata text-metadata text-on-surface-variant uppercase mb-2">
    Monthly Usage
   </h3>
   <div class="font-kpi-number text-kpi-number text-tertiary">
    4,892
   </div>
   <div class="font-metadata text-metadata text-text-secondary mt-1">
    Units consumed this month
   </div>
  </div>
  <div class="bg-metric-sage rounded-xl p-card-padding shadow-[0_4px_16px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform cursor-pointer border border-[#EEEBE4]">
   <h3 class="font-metadata text-metadata text-on-surface-variant uppercase mb-2">
    Total Value
   </h3>
   <div class="font-kpi-number text-kpi-number text-status-in-stock">
    ₹4.2L
   </div>
   <div class="font-metadata text-metadata text-text-secondary mt-1">
    Estimated current stock value
   </div>
  </div>
 </div>
 <!-- Filter Bar & Table Area -->
 <div class="bg-surface rounded-xl border border-border-default shadow-[0_4px_16px_rgba(0,0,0,0.02)] overflow-hidden">
  <!-- Filters -->
  <div class="p-4 border-b border-border-default flex justify-between items-center bg-surface-subtle">
   <div class="flex gap-inline">
    <div class="relative w-64">
     <span class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-stone-400 text-[18px]">
      search
     </span>
     <input class="w-full bg-surface border border-border-default rounded-lg py-1.5 pl-9 pr-3 text-body font-body focus:outline-none focus:border-amber-700 shadow-sm" placeholder="Filter by item name..." type="text"/>
    </div>
    <select class="bg-surface border border-border-default rounded-lg py-1.5 px-3 text-body font-body shadow-sm focus:outline-none focus:border-amber-700">
     <option>
      All Categories
     </option>
     <option>
      Office Supplies
     </option>
     <option>
      Pantry
     </option>
     <option>
      Cleaning
     </option>
     <option>
      IT Peripherals
     </option>
    </select>
   </div>
   <div class="flex gap-2">
    <button class="p-1.5 text-text-secondary hover:text-text-primary hover:bg-surface-variant rounded transition-colors">
     <span class="material-symbols-outlined">
      filter_list
     </span>
    </button>
    <button class="p-1.5 text-text-secondary hover:text-text-primary hover:bg-surface-variant rounded transition-colors">
     <span class="material-symbols-outlined">
      view_column
     </span>
    </button>
   </div>
  </div>
  <!-- Table -->
  <div class="bg-surface rounded-xl shadow-sm overflow-hidden border-t border-border-default">
    <DataTable :value="consumables" :loading="loading" paginator :rows="10" tableStyle="min-width: 50rem" class="w-full">
      <Column field="name" header="Item Name" sortable></Column>
      <Column field="category" header="Category" sortable></Column>
      <Column field="location" header="Location" sortable></Column>
      <Column field="currentStock" header="Current Stock" sortable>
        <template #body="slotProps">
          <span :class="{'text-error-container font-bold': slotProps.data.currentStock === 0, 'text-primary-container font-bold': slotProps.data.currentStock > 0 && slotProps.data.currentStock <= slotProps.data.reorderLevel}">
            {{ slotProps.data.currentStock }}
          </span>
        </template>
      </Column>
      <Column field="reorderLevel" header="Reorder Level" sortable></Column>
      <Column header="Status">
        <template #body="slotProps">
          <span v-if="slotProps.data.currentStock === 0" class="inline-flex items-center px-2 py-0.5 rounded-full text-[11px] font-medium bg-error-container text-on-error-container whitespace-nowrap">
            Out of Stock
          </span>
          <span v-else-if="slotProps.data.currentStock <= slotProps.data.reorderLevel" class="inline-flex items-center px-2 py-0.5 rounded-full text-[11px] font-medium bg-metric-amber text-primary-container whitespace-nowrap">
            Low Stock
          </span>
          <span v-else class="inline-flex items-center px-2 py-0.5 rounded-full text-[11px] font-medium bg-metric-sage text-status-in-stock whitespace-nowrap">
            Adequate
          </span>
        </template>
      </Column>
    </DataTable>
  </div>
 </div>
</main>

</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'

interface Consumable {
  id: number;
  name: string;
  category: string;
  location: string;
  currentStock: number;
  reorderLevel: number;
}

const consumables = ref<Consumable[]>([])
const loading = ref(true)

const fetchConsumables = async () => {
  try {
    const res = await fetch('http://localhost:8080/api/consumables')
    if (res.ok) {
      consumables.value = await res.json()
    } else {
      console.error('Failed to fetch consumables')
    }
  } catch (error) {
    console.error('API connection error', error)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchConsumables()
})
</script>
