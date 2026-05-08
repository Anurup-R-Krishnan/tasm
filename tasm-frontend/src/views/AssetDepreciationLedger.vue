<template>
  <main class="space-y-section-gap pb-24 font-body">
    <!-- Page Header: Asset Identity -->
    <div class="flex flex-col md:flex-row md:items-end justify-between gap-4">
      <div v-if="loading" class="animate-pulse space-y-2">
        <div class="h-4 w-32 bg-border-default rounded"></div>
        <div class="h-8 w-64 bg-border-default rounded"></div>
      </div>
      <div v-else>
        <div class="flex items-center gap-2 mb-2">
          <RouterLink
            to="/financials"
            class="font-metadata text-metadata text-text-secondary hover:text-primary transition-colors"
          >
            Financials
          </RouterLink>
          <span class="material-symbols-outlined text-[14px] text-text-secondary">
            chevron_right
          </span>
          <RouterLink
            to="/depreciation"
            class="font-metadata text-metadata font-medium"
            :class="
              isDetailView ? 'text-text-secondary hover:text-primary' : 'text-primary-container'
            "
          >
            Depreciation Ledger
          </RouterLink>
          <template v-if="isDetailView && selectedSchedule">
            <span class="material-symbols-outlined text-[14px] text-text-secondary">
              chevron_right
            </span>
            <span class="font-metadata text-metadata text-primary-container font-medium">
              {{ selectedSchedule.assetName }}
            </span>
          </template>
        </div>

        <h2 class="font-h1 text-h1 text-text-primary mb-1">
          {{
            isDetailView && selectedSchedule
              ? selectedSchedule.assetName
              : 'Asset Depreciation Ledger'
          }}
          <span
            v-if="isDetailView && selectedSchedule"
            class="text-text-secondary font-mono font-normal ml-2"
            >#{{ selectedSchedule.assetId }}</span
          >
        </h2>

        <div class="flex items-center gap-3">
          <span
            class="inline-flex items-center px-2 py-0.5 rounded text-[10px] font-bold uppercase tracking-widest bg-surface-variant text-on-surface-variant border border-outline-variant"
          >
            {{ isDetailView && selectedSchedule ? 'Asset Focus' : 'System-wide Overview' }}
          </span>
          <span v-if="!isDetailView" class="text-xs font-medium text-text-secondary">
            Tracking {{ schedules.length }} active asset lifecycles
          </span>
          <span
            v-else-if="selectedSchedule"
            class="inline-flex items-center gap-1 text-xs font-medium text-status-in-stock"
          >
            <span class="w-2 h-2 rounded-full bg-status-in-stock"> </span>
            {{ selectedSchedule.method }} Method
          </span>
        </div>
      </div>

      <div class="flex items-center gap-3">
        <button
          @click="handlePrint"
          class="px-4 py-2 bg-surface text-text-primary border border-border-default rounded-xl font-bold text-xs hover:bg-surface-subtle transition-all flex items-center gap-2 shadow-sm"
        >
          <span class="material-symbols-outlined text-[18px]"> print </span>
          Print
        </button>
        <button
          @click="handleExport"
          class="px-4 py-2 bg-primary text-on-primary rounded-xl font-bold text-xs hover:opacity-90 transition-all flex items-center gap-2 shadow-lg shadow-primary/10"
        >
          <span class="material-symbols-outlined text-[18px]"> download </span>
          Export CSV
        </button>
      </div>
    </div>

    <!-- KPI Cards Grid -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
      <!-- KPI 1: Acquisition Cost -->
      <div class="premium-card group hover:border-primary/20 transition-all">
        <div class="flex items-center justify-between mb-4">
          <span
            class="font-metadata text-[10px] font-bold text-text-secondary uppercase tracking-[0.15em]"
          >
            {{ isDetailView ? 'Initial Cost' : 'Total Acquisition' }}
          </span>
          <div
            class="w-8 h-8 rounded-lg bg-primary-container/10 text-primary flex items-center justify-center"
          >
            <span class="material-symbols-outlined text-[20px]">payments</span>
          </div>
        </div>
        <div class="text-3xl font-bold text-text-primary">
          ₹{{ currentPurchaseValue.toLocaleString() }}
        </div>
        <div class="text-[10px] text-text-secondary mt-2 font-medium">
          {{ isDetailView ? 'Asset purchase price' : 'Combined value of all tracked assets' }}
        </div>
      </div>

      <!-- KPI 2: Depreciation -->
      <div class="premium-card border-rose-500/10">
        <div class="flex items-center justify-between mb-4">
          <span
            class="font-metadata text-[10px] font-bold text-text-secondary uppercase tracking-[0.15em]"
          >
            Accumulated Depreciation
          </span>
          <div
            class="w-8 h-8 rounded-lg bg-error-container/20 text-status-critical flex items-center justify-center"
          >
            <span class="material-symbols-outlined text-[20px]">trending_down</span>
          </div>
        </div>
        <div class="text-3xl font-bold text-status-critical">
          ₹{{ currentAccumulatedDepreciation.toLocaleString() }}
        </div>
        <div class="text-[10px] text-text-secondary mt-2 font-medium">
          Total value written off to date
        </div>
      </div>

      <!-- KPI 3: Net Book Value -->
      <div
        class="premium-card bg-metric-sage/30 border-status-in-stock/20 relative overflow-hidden"
      >
        <div class="absolute -right-2 -top-2 opacity-10 pointer-events-none transform rotate-12">
          <span
            class="material-symbols-outlined text-[100px]"
            style="font-variation-settings: 'FILL' 1"
            >account_balance</span
          >
        </div>
        <div class="flex items-center justify-between mb-4 relative z-10">
          <span
            class="font-metadata text-[10px] font-bold text-status-in-stock uppercase tracking-[0.15em]"
          >
            Net Book Value
          </span>
          <div
            class="w-8 h-8 rounded-lg bg-white/50 text-status-in-stock flex items-center justify-center shadow-sm"
          >
            <span class="material-symbols-outlined text-[20px]">account_balance_wallet</span>
          </div>
        </div>
        <div class="text-3xl font-bold text-status-in-stock relative z-10">
          ₹{{ currentNetBookValue.toLocaleString() }}
        </div>
        <div class="text-[10px] text-status-in-stock/70 mt-2 font-medium relative z-10">
          Current carrying value on balance sheet
        </div>
      </div>
    </div>

    <!-- Table Section (Full Width) -->
    <div class="premium-card !p-0 overflow-hidden">
      <div
        class="p-6 border-b border-border-default bg-surface-subtle/30 flex items-center justify-between"
      >
        <div>
          <h3 class="text-sm font-bold text-text-primary">Depreciation Schedules</h3>
          <p class="text-[10px] text-text-secondary font-medium uppercase tracking-widest mt-1">
            Historical & projected valuation metrics
          </p>
        </div>
        <div class="flex items-center gap-3">
          <div class="relative">
            <span
              class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-text-secondary text-[18px]"
              >search</span
            >
            <input
              v-model="searchQuery"
              placeholder="Search assets..."
              class="h-10 pl-10 pr-4 bg-white border border-border-default rounded-xl text-xs focus:ring-2 focus:ring-primary/20 outline-none transition-all w-64"
            />
          </div>
        </div>
      </div>

      <div class="overflow-x-auto">
        <table class="w-full text-left">
          <thead
            class="bg-surface-subtle/50 text-[10px] font-bold text-text-secondary uppercase tracking-widest border-b border-border-default"
          >
            <tr>
              <th class="px-6 py-4">Asset ID</th>
              <th class="px-6 py-4">Asset Name</th>
              <th class="px-6 py-4">Purchase Value</th>
              <th class="px-6 py-4">Net Book Value</th>
              <th class="px-6 py-4">Method</th>
              <th class="px-6 py-4 text-right">Actions</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-50">
            <tr
              v-for="s in filteredSchedules"
              :key="s.id"
              class="hover:bg-surface-subtle/50 transition-colors group cursor-pointer"
              :class="{ 'bg-primary/5': isDetailView && s.id === Number(route.params['id']) }"
              @click="router.push(`/depreciation/${s.id}`)"
            >
              <td class="px-6 py-4">
                <span class="font-mono text-[11px] text-surface-tint font-bold"
                  >#{{ s.assetId }}</span
                >
              </td>
              <td class="px-6 py-4">
                <span
                  class="text-sm font-bold text-text-primary group-hover:text-primary transition-colors"
                  >{{ s.assetName }}</span
                >
              </td>
              <td class="px-6 py-4">
                <span class="text-xs font-medium text-text-secondary"
                  >₹{{ s.purchaseValue.toLocaleString() }}</span
                >
              </td>
              <td class="px-6 py-4">
                <span class="text-xs font-bold text-primary-container"
                  >₹{{ s.currentValue.toLocaleString() }}</span
                >
              </td>
              <td class="px-6 py-4">
                <span
                  class="text-[10px] font-bold uppercase tracking-wider px-2 py-0.5 bg-surface-variant text-on-surface-variant rounded-full border border-border-default"
                >
                  {{ s.method }}
                </span>
              </td>
              <td class="px-6 py-4 text-right">
                <RouterLink
                  :to="`/asset/${s.assetId}`"
                  class="p-2 text-text-secondary hover:text-primary transition-colors"
                >
                  <span class="material-symbols-outlined text-[20px]">visibility</span>
                </RouterLink>
              </td>
            </tr>
            <tr v-if="filteredSchedules.length === 0">
              <td colspan="6" class="px-6 py-12 text-center text-text-secondary text-sm italic">
                No depreciation records found.
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue';
import { useRoute, useRouter, RouterLink } from 'vue-router';
import { getDepreciations, getDepreciationById } from '../api/financial';
import type { DepreciationSchedule } from '../types/models';

const route = useRoute();
const router = useRouter();

const schedules = ref<DepreciationSchedule[]>([]);
const selectedSchedule = ref<DepreciationSchedule | null>(null);
const loading = ref(true);
const searchQuery = ref('');

const isDetailView = computed(() => !!route.params['id']);

const currentPurchaseValue = computed(() => {
  if (isDetailView.value && selectedSchedule.value) {
    return selectedSchedule.value.purchaseValue;
  }
  return schedules.value.reduce((sum, item) => sum + (item.purchaseValue || 0), 0);
});

const currentAccumulatedDepreciation = computed(() => {
  if (isDetailView.value && selectedSchedule.value) {
    return selectedSchedule.value.purchaseValue - selectedSchedule.value.currentValue;
  }
  return schedules.value.reduce(
    (sum, item) => sum + ((item.purchaseValue || 0) - (item.currentValue || 0)),
    0,
  );
});

const currentNetBookValue = computed(() => {
  if (isDetailView.value && selectedSchedule.value) {
    return selectedSchedule.value.currentValue;
  }
  return schedules.value.reduce((sum, item) => sum + (item.currentValue || 0), 0);
});

const filteredSchedules = computed(() => {
  const query = searchQuery.value.toLowerCase();
  return schedules.value.filter(
    (s) => s.assetName.toLowerCase().includes(query) || s.assetId.toLowerCase().includes(query),
  );
});

const fetchData = async () => {
  loading.value = true;
  try {
    const all = await getDepreciations();
    schedules.value = all;

    if (route.params['id']) {
      selectedSchedule.value = await getDepreciationById(route.params['id'] as string);
    } else {
      selectedSchedule.value = null;
    }
  } catch (error) {
    console.error('Failed to fetch depreciation data:', error);
  } finally {
    loading.value = false;
  }
};

watch(() => route.params['id'], fetchData);

const handleExport = () => {
  const headers = ['Asset ID', 'Asset Name', 'Purchase Value', 'Current Value', 'Method'];
  const rows = schedules.value.map((s) => [
    s.assetId,
    s.assetName,
    s.purchaseValue,
    s.currentValue,
    s.method,
  ]);

  const csvContent = [headers, ...rows].map((e) => e.join(',')).join('\n');
  const blob = new Blob([csvContent], { type: 'text/csv;charset=utf-8;' });
  const link = document.createElement('a');
  link.href = URL.createObjectURL(blob);
  link.download = `depreciation_ledger_${new Date().toISOString().split('T')[0]}.csv`;
  link.click();
};

const handlePrint = () => {
  window.print();
};

onMounted(fetchData);
</script>
