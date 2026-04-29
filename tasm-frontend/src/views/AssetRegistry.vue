<template>
  <div>
    <div class="flex justify-between items-end mb-6">
      <div>
        <h1 class="font-h1 text-h1 text-text-primary mb-1">Asset Registry</h1>
        <p class="font-body text-body text-text-secondary">Manage and track all physical and digital assets.</p>
      </div>
      <button class="px-4 py-2 bg-amber-700 text-white rounded-lg hover:bg-amber-800 transition-colors">
        + Add New Asset
      </button>
    </div>

    <div class="bg-surface rounded-xl border border-border-default shadow-sm p-4">
      <DataTable :value="assets" :loading="loading" paginator :rows="10">
        <Column field="name" header="Name"></Column>
        <Column field="category" header="Category"></Column>
        <Column field="location" header="Location"></Column>
        <Column field="status" header="Status"></Column>
        <Column field="currentStock" header="Stock"></Column>
      </DataTable>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'

interface Asset {
  id: number;
  name: string;
  category: string;
  location: string;
  status: string;
  currentStock: number;
}

const assets = ref<Asset[]>([])
const loading = ref(true)

const fetchAssets = async () => {
  try {
    const res = await fetch('http://localhost:8080/api/assets')
    if (res.ok) {
      assets.value = await res.json()
    } else {
      console.error('Failed to fetch assets')
    }
  } catch (error) {
    console.error('API connection error', error)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchAssets()
})
</script>
