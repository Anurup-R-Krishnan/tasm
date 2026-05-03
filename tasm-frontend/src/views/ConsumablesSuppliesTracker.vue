<template>
  <main class="space-y-section-gap pb-24">
    <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
      <div>
        <h1 class="text-text-primary mb-1">Consumables Tracker</h1>
        <p class="text-text-secondary">Manage and track daily supplies and consumable inventory.</p>
      </div>
      <div class="flex gap-3">
        <button
          class="bg-surface border border-border-default text-text-primary px-4 py-2 rounded-lg text-sm flex items-center gap-2 hover:bg-surface-subtle transition-colors shadow-sm"
        >
          <span class="material-symbols-outlined text-[18px]">download</span>
          Receive Stock
        </button>
        <button class="btn-primary">
          <span class="material-symbols-outlined">add_circle</span>
          Issue Consumable
        </button>
      </div>
    </div>

    <!-- Metrics Grid -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div v-for="kpi in kpiMetrics" :key="kpi.label" class="premium-card">
        <p class="text-[10px] font-bold text-text-secondary uppercase tracking-widest mb-2">
          {{ kpi.label }}
        </p>
        <div class="text-3xl font-bold" :class="kpi.colorClass">
          {{ kpi.value }}
        </div>
        <p class="text-xs text-text-secondary mt-2 flex items-center gap-1">
          <span v-if="kpi.icon" class="material-symbols-outlined text-sm">{{ kpi.icon }}</span>
          {{ kpi.subtext }}
        </p>
      </div>
    </div>

    <div class="premium-card !p-0 overflow-hidden">
      <div
        class="p-4 border-b border-border-default flex flex-col md:flex-row justify-between items-center gap-4 bg-surface-subtle/30"
      >
        <div class="flex flex-1 gap-4 w-full md:w-auto">
          <div class="relative flex-1 max-w-sm">
            <span
              class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-text-secondary text-[18px]"
              >search</span
            >
            <input
              v-model="searchQuery"
              class="w-full bg-white border border-border-default rounded-xl py-2 pl-10 pr-4 text-sm focus:ring-2 focus:ring-primary/20 outline-none transition-all"
              placeholder="Search items..."
            />
          </div>
          <select
            v-model="selectedCategory"
            class="bg-white border border-border-default rounded-xl py-2 px-4 text-sm outline-none focus:ring-2 focus:ring-primary/20"
          >
            <option :value="null">All Categories</option>
            <option v-for="opt in categoryOptions" :key="opt.value" :value="opt.value">
              {{ opt.label }}
            </option>
          </select>
        </div>
      </div>

      <div class="overflow-x-auto">
        <table class="w-full text-left">
          <thead
            class="bg-surface-subtle/50 text-[10px] font-bold text-text-secondary uppercase tracking-widest border-b border-border-default"
          >
            <tr>
              <th class="px-6 py-4">Item Name</th>
              <th class="px-6 py-4">Category</th>
              <th class="px-6 py-4">Location</th>
              <th class="px-6 py-4">Stock Level</th>
              <th class="px-6 py-4">Status</th>
              <th class="px-6 py-4 text-right">Actions</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-border-default/20">
            <tr
              v-for="item in filteredConsumables"
              :key="item.id"
              class="hover:bg-surface-subtle/50 transition-colors"
            >
              <td class="px-6 py-4 font-bold text-text-primary text-sm">{{ item.name }}</td>
              <td class="px-6 py-4">
                <span
                  class="text-xs text-text-secondary bg-surface-variant px-2 py-0.5 rounded-lg"
                  >{{ item.category }}</span
                >
              </td>
              <td class="px-6 py-4 text-xs text-text-secondary">{{ item.location }}</td>
              <td class="px-6 py-4">
                <div class="flex items-center gap-2">
                  <span
                    class="text-sm font-bold"
                    :class="
                      item.currentStock <= item.reorderLevel
                        ? 'text-status-critical'
                        : 'text-text-primary'
                    "
                  >
                    {{ item.currentStock }}
                  </span>
                  <span class="text-[10px] text-text-secondary">/ {{ item.reorderLevel }} min</span>
                </div>
              </td>
              <td class="px-6 py-4">
                <span
                  class="text-[10px] font-bold uppercase tracking-wider px-2 py-0.5 rounded-full"
                  :class="stockSeverityClass(item)"
                >
                  {{ stockLabel(item) }}
                </span>
              </td>
              <td class="px-6 py-4 text-right">
                <button
                  @click="deleteConsumable(item.id)"
                  class="p-2 text-text-secondary hover:text-status-critical hover:bg-error-container/50 rounded-lg transition-all"
                >
                  <span class="material-symbols-outlined text-[20px]">delete</span>
                </button>
              </td>
            </tr>
            <tr v-if="filteredConsumables.length === 0">
              <td colspan="6" class="px-6 py-12 text-center text-text-secondary text-sm italic">
                No items found matching filters.
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import {
  getConsumables,
  deleteConsumable as apiDeleteConsumable,
  type Consumable,
} from '../api/consumables';

const consumables = ref<Consumable[]>([]);
const loading = ref(true);
const errorMessage = ref('');
const searchQuery = ref('');
const selectedCategory = ref<string | null>(null);

const categoryOptions = computed(() => [
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

const kpiMetrics = computed(() => {
  const lowStockItems = consumables.value.filter(
    (item) => item.currentStock <= item.reorderLevel,
  ).length;
  const totalLocations = new Set(consumables.value.map((item) => item.location)).size;

  return [
    {
      label: 'Total Inventory',
      value: consumables.value.length,
      colorClass: 'text-text-primary',
      subtext: 'Active SKUs tracked',
    },
    {
      label: 'Low Stock Alerts',
      value: lowStockItems,
      colorClass: 'text-status-critical',
      icon: 'warning',
      subtext: 'Requires restocking',
    },
    {
      label: 'Storage Points',
      value: totalLocations,
      colorClass: 'text-primary',
      subtext: 'Across campus',
    },
    {
      label: 'Stock Health',
      value:
        Math.round(
          ((consumables.value.length - lowStockItems) / (consumables.value.length || 1)) * 100,
        ) + '%',
      colorClass: 'text-status-in-stock',
      subtext: 'Items above min level',
    },
  ];
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
    await apiDeleteConsumable(id);
    consumables.value = consumables.value.filter((c) => c.id !== id);
  } catch (error) {
    console.error('Error deleting item:', error);
    alert('Failed to delete item');
  }
}

function stockLabel(consumable: Consumable): string {
  if (consumable.currentStock === 0) return 'Out of Stock';
  if (consumable.currentStock <= consumable.reorderLevel) return 'Low Stock';
  return 'In Stock';
}

function stockSeverityClass(consumable: Consumable): string {
  if (consumable.currentStock === 0) return 'bg-status-critical/20 text-status-critical';
  if (consumable.currentStock <= consumable.reorderLevel)
    return 'bg-metric-amber/20 text-surface-tint';
  return 'bg-status-in-stock/20 text-status-in-stock';
}

onMounted(() => {
  void fetchConsumables();
});
</script>
