<template>
  <main class="p-page-margin max-w-[1400px] mx-auto space-y-section-gap">
    <div class="flex justify-between items-end">
      <div>
        <h1 class="font-h1 text-h1 text-text-primary mb-1">Consumables Tracker</h1>
        <p class="font-body text-body text-text-secondary">
          Manage and track daily supplies and consumable inventory.
        </p>
      </div>
      <div class="flex gap-inline">
        <Button icon="pi pi-download" label="Receive Stock" outlined severity="secondary" />
        <Button icon="pi pi-arrow-up-right" label="Issue Consumable" />
      </div>
    </div>

    <div class="grid grid-cols-4 gap-inline">
      <div
        class="bg-metric-amber rounded-xl p-card-padding shadow-[0_4px_16px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform cursor-pointer border border-[#EEEBE4]"
      >
        <h3 class="font-metadata text-metadata text-on-surface-variant uppercase mb-2">
          Total Items
        </h3>
        <div class="font-kpi-number text-kpi-number text-text-primary">
          {{ metrics.totalItems }}
        </div>
        <div class="font-metadata text-metadata text-text-secondary mt-1">
          Across {{ metrics.totalCategories }} categories
        </div>
      </div>
      <div
        class="bg-surface-container-high rounded-xl p-card-padding shadow-[0_4px_16px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform cursor-pointer border border-[#EEEBE4]"
      >
        <h3 class="font-metadata text-metadata text-on-surface-variant uppercase mb-2">
          Low Stock Items
        </h3>
        <div class="font-kpi-number text-kpi-number text-primary-container">
          {{ metrics.lowStockItems }}
        </div>
        <div class="font-metadata text-metadata text-text-secondary mt-1 flex items-center gap-1">
          <span class="material-symbols-outlined text-[14px] text-primary-container">
            warning
          </span>
          Requires attention
        </div>
      </div>
      <div
        class="bg-metric-lavender rounded-xl p-card-padding shadow-[0_4px_16px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform cursor-pointer border border-[#EEEBE4]"
      >
        <h3 class="font-metadata text-metadata text-on-surface-variant uppercase mb-2">
          Active Locations
        </h3>
        <div class="font-kpi-number text-kpi-number text-tertiary">
          {{ metrics.totalLocations }}
        </div>
        <div class="font-metadata text-metadata text-text-secondary mt-1">
          Locations stocking consumables
        </div>
      </div>
      <div
        class="bg-metric-sage rounded-xl p-card-padding shadow-[0_4px_16px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform cursor-pointer border border-[#EEEBE4]"
      >
        <h3 class="font-metadata text-metadata text-on-surface-variant uppercase mb-2">
          Adequate Stock
        </h3>
        <div class="font-kpi-number text-kpi-number text-status-in-stock">
          {{ metrics.healthyItems }}
        </div>
        <div class="font-metadata text-metadata text-text-secondary mt-1">
          Items above reorder level
        </div>
      </div>
    </div>

    <div
      class="bg-surface rounded-xl border border-border-default shadow-[0_4px_16px_rgba(0,0,0,0.02)] overflow-hidden"
    >
      <div
        class="p-4 border-b border-border-default flex justify-between items-center bg-surface-subtle"
      >
        <div class="flex gap-inline">
          <div class="relative w-64">
            <span
              class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-stone-400 text-[18px]"
            >
              search
            </span>
            <InputText
              v-model="searchQuery"
              class="w-full !bg-surface !border-border-default !rounded-lg !py-1.5 !pl-9 !pr-3 !text-body !font-body"
              placeholder="Filter by item name..."
            />
          </div>
          <Select
            v-model="selectedCategory"
            :options="categoryOptions"
            class="min-w-56"
            option-label="label"
            option-value="value"
            placeholder="All Categories"
          />
        </div>
        <div class="flex gap-2">
          <Button aria-label="Filter list" icon="pi pi-filter" rounded severity="secondary" text />
          <Button
            aria-label="Toggle columns"
            icon="pi pi-table"
            rounded
            severity="secondary"
            text
          />
        </div>
      </div>

      <div class="bg-surface rounded-xl shadow-sm overflow-hidden border-t border-border-default">
        <Message v-if="errorMessage" severity="error" class="m-4">
          {{ errorMessage }}
        </Message>
        <DataTable
          :value="filteredConsumables"
          :loading="loading"
          paginator
          :rows="10"
          tableStyle="min-width: 50rem"
          class="w-full"
          dataKey="id"
        >
          <template #empty> No consumables matched the current filters. </template>
          <Column field="name" header="Item Name" sortable />
          <Column field="category" header="Category" sortable />
          <Column field="location" header="Location" sortable />
          <Column field="currentStock" header="Current Stock" sortable>
            <template #body="slotProps">
              <span
                :class="{
                  'text-error-container font-bold': slotProps.data.currentStock === 0,
                  'text-primary-container font-bold':
                    slotProps.data.currentStock > 0 &&
                    slotProps.data.currentStock <= slotProps.data.reorderLevel,
                }"
              >
                {{ slotProps.data.currentStock }}
              </span>
            </template>
          </Column>
          <Column field="reorderLevel" header="Reorder Level" sortable />
          <Column header="Status">
            <template #body="slotProps">
              <Tag :severity="stockSeverity(slotProps.data)" :value="stockLabel(slotProps.data)" />
            </template>
          </Column>
          <Column header="Actions" alignFrozen="right">
            <template #body="slotProps">
              <button
                @click="deleteConsumable(slotProps.data.id)"
                class="p-1 text-text-secondary hover:text-status-critical rounded transition-colors"
                title="Delete Item"
              >
                <span class="material-symbols-outlined text-[20px]">delete</span>
              </button>
            </template>
          </Column>
        </DataTable>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import Button from 'primevue/button';
import Column from 'primevue/column';
import DataTable from 'primevue/datatable';
import InputText from 'primevue/inputtext';
import Message from 'primevue/message';
import Select from 'primevue/select';
import Tag from 'primevue/tag';
import { getConsumables, type Consumable } from '../api/consumables';

const consumables = ref<Consumable[]>([]);
const loading = ref(true);
const errorMessage = ref('');
const searchQuery = ref('');
const selectedCategory = ref<string | null>(null);

const categoryOptions = computed(() => [
  { label: 'All Categories', value: null },
  ...Array.from(new Set(consumables.value.map((item) => item.category)))
    .sort()
    .map((category) => ({
      label: category,
      value: category,
    })),
]);

const filteredConsumables = computed(() => {
  const normalizedQuery = searchQuery.value.trim().toLowerCase();

  return consumables.value.filter((item) => {
    const matchesQuery =
      normalizedQuery.length === 0 ||
      item.name.toLowerCase().includes(normalizedQuery) ||
      item.location.toLowerCase().includes(normalizedQuery);

    const matchesCategory =
      selectedCategory.value === null || item.category === selectedCategory.value;

    return matchesQuery && matchesCategory;
  });
});

const metrics = computed(() => {
  const lowStockItems = consumables.value.filter(
    (item) => item.currentStock <= item.reorderLevel,
  ).length;
  const totalLocations = new Set(consumables.value.map((item) => item.location)).size;
  const totalCategories = new Set(consumables.value.map((item) => item.category)).size;

  return {
    totalItems: consumables.value.length,
    lowStockItems,
    totalLocations,
    totalCategories,
    healthyItems: Math.max(consumables.value.length - lowStockItems, 0),
  };
});

async function fetchConsumables(): Promise<void> {
  try {
    consumables.value = await getConsumables();
  } catch (error) {
    console.error('API connection error', error);
    errorMessage.value = error instanceof Error ? error.message : 'Failed to fetch consumables.';
  } finally {
    loading.value = false;
  }
}

async function deleteConsumable(id: number): Promise<void> {
  if (!confirm('Are you sure you want to delete this item?')) return;
  try {
    const res = await fetch(`http://localhost:8080/api/consumables/${id}`, {
      method: 'DELETE',
    });
    if (res.ok) {
      consumables.value = consumables.value.filter((c) => c.id !== id);
    } else {
      alert('Failed to delete item');
    }
  } catch (error) {
    console.error('Error deleting item:', error);
  }
}

function stockLabel(consumable: Consumable): string {
  if (consumable.currentStock === 0) {
    return 'Out of Stock';
  }

  if (consumable.currentStock <= consumable.reorderLevel) {
    return 'Low Stock';
  }

  return 'Adequate';
}

function stockSeverity(consumable: Consumable): 'danger' | 'warn' | 'success' {
  if (consumable.currentStock === 0) {
    return 'danger';
  }

  if (consumable.currentStock <= consumable.reorderLevel) {
    return 'warn';
  }

  return 'success';
}

onMounted(() => {
  void fetchConsumables();
});
</script>
