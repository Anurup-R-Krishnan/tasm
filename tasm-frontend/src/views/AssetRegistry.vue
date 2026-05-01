<template>
  <main class="flex-1 flex flex-col min-w-0">
    <!-- TopNavBar Component -->
    <header
      class="h-[60px] w-full border-b sticky top-0 z-30 border-stone-200 dark:border-stone-800 shadow-sm bg-stone-50 dark:bg-stone-900 font-sans antialiased text-stone-900 dark:text-stone-100"
    >
      <div class="flex items-center justify-between px-6 w-full max-w-[1400px] mx-auto h-full">
        <!-- Brand (Optional in Hybrid, but part of TopNavBar JSON) -->
        <div
          class="text-lg font-bold tracking-tight text-stone-900 dark:text-stone-50 flex items-center gap-4"
        >
          <div class="relative w-64">
            <span
              class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-text-secondary text-[20px]"
              data-icon="search"
            >
              search
            </span>
            <input
              class="w-full pl-10 pr-4 py-1.5 bg-surface border border-border-default rounded-md text-sm focus:outline-none focus:border-primary-container"
              placeholder="Search..."
              type="text"
            />
          </div>
        </div>
        <!-- Trailing Actions -->
        <div class="flex items-center gap-2">
          <button
            class="p-2 rounded-full text-stone-500 dark:text-stone-400 hover:bg-stone-100 dark:hover:bg-stone-800 transition-colors cursor-pointer active:scale-95 duration-150"
          >
            <span class="material-symbols-outlined" data-icon="notifications"> notifications </span>
          </button>
          <button
            class="p-2 rounded-full text-stone-500 dark:text-stone-400 hover:bg-stone-100 dark:hover:bg-stone-800 transition-colors cursor-pointer active:scale-95 duration-150"
          >
            <span class="material-symbols-outlined" data-icon="settings"> settings </span>
          </button>
          <button
            class="p-2 rounded-full text-stone-500 dark:text-stone-400 hover:bg-stone-100 dark:hover:bg-stone-800 transition-colors cursor-pointer active:scale-95 duration-150"
          >
            <span class="material-symbols-outlined" data-icon="help"> help </span>
          </button>
          <div
            class="ml-2 w-8 h-8 rounded-full bg-border-default overflow-hidden border border-border-default"
          >
            <img
              alt="User profile"
              class="w-full h-full object-cover"
              data-alt="Close up portrait of a professional smiling slightly in natural lighting"
              src="https://lh3.googleusercontent.com/aida-public/AB6AXuAgoTg1INEPcxLsEzOeYnNxgTpLngLZZlBHcjv5aUMYnl6DakeZPzs2JyBmrTV_ZqDDAZahGRt8MV4vZ3Z2W_AiNmkUoseDdOU7c2jMe_9OahcIx-DFeX9SwWYDu9IVS0RL5dE4Ayl--_AmjpV0UW7hlbLB5wrQk-nfsJo2It5n2wpB4hrj2iFeI5T0f7oAbZVq8nPP-DW_sye8TTOfNNf8qZkaOt7MmTOK9eqieEUVs2CrfaKQqBQgh007sS6AfTPasBSBKTKPwWXi"
            />
          </div>
        </div>
      </div>
    </header>
    <!-- Canvas Content -->
    <div class="p-page-margin max-w-[1400px] mx-auto w-full">
      <!-- Page Header -->
      <div class="flex justify-between items-center mb-stack">
        <h1 class="font-h1 text-h1 text-text-primary">Asset Registry</h1>
        <div class="flex items-center gap-inline">
          <button
            class="px-4 py-2 bg-surface border border-border-default rounded-lg text-text-primary font-h3 text-h3 hover:bg-surface-subtle transition-colors flex items-center gap-2 shadow-sm"
          >
            <span class="material-symbols-outlined text-[18px]" data-icon="download">
              download
            </span>
            Export
          </button>
          <button
            class="px-4 py-2 bg-text-primary text-on-primary rounded-lg font-h3 text-h3 hover:bg-black transition-colors flex items-center gap-2 shadow-sm"
          >
            <span class="material-symbols-outlined text-[18px]" data-icon="add"> add </span>
            Add New Asset
          </button>
        </div>
      </div>
      <!-- Status Summary Pills -->
      <div class="flex gap-inline overflow-x-auto pb-2 mb-stack hide-scrollbar">
        <button
          class="px-4 py-1.5 rounded-full bg-primary-container text-on-primary font-body text-body whitespace-nowrap flex items-center gap-2 border border-primary-container"
        >
          All
          <span class="bg-white/20 px-2 py-0.5 rounded-full text-[12px]"> 1,248 </span>
        </button>
        <button
          class="px-4 py-1.5 rounded-full bg-surface border border-border-default text-text-secondary hover:border-text-secondary transition-colors font-body text-body whitespace-nowrap flex items-center gap-2"
        >
          In Stock
          <span class="bg-surface-subtle px-2 py-0.5 rounded-full text-[12px] text-text-primary">
            432
          </span>
        </button>
        <button
          class="px-4 py-1.5 rounded-full bg-surface border border-border-default text-text-secondary hover:border-text-secondary transition-colors font-body text-body whitespace-nowrap flex items-center gap-2"
        >
          Checked Out
          <span class="bg-surface-subtle px-2 py-0.5 rounded-full text-[12px] text-text-primary">
            782
          </span>
        </button>
        <button
          class="px-4 py-1.5 rounded-full bg-surface border border-border-default text-text-secondary hover:border-text-secondary transition-colors font-body text-body whitespace-nowrap flex items-center gap-2"
        >
          Reserved
          <span class="bg-surface-subtle px-2 py-0.5 rounded-full text-[12px] text-text-primary">
            12
          </span>
        </button>
        <button
          class="px-4 py-1.5 rounded-full bg-surface border border-border-default text-text-secondary hover:border-text-secondary transition-colors font-body text-body whitespace-nowrap flex items-center gap-2"
        >
          Under Repair
          <span class="bg-surface-subtle px-2 py-0.5 rounded-full text-[12px] text-text-primary">
            18
          </span>
        </button>
        <button
          class="px-4 py-1.5 rounded-full bg-surface border border-border-default text-text-secondary hover:border-text-secondary transition-colors font-body text-body whitespace-nowrap flex items-center gap-2"
        >
          Lost
          <span class="bg-surface-subtle px-2 py-0.5 rounded-full text-[12px] text-text-primary">
            4
          </span>
        </button>
      </div>
      <!-- Filter Bar -->
      <div
        class="flex justify-between items-center bg-surface p-inline rounded-xl border border-border-default mb-stack shadow-sm"
      >
        <div class="flex items-center gap-inline flex-1">
          <div class="relative w-80">
            <span
              class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-text-secondary text-[20px]"
              data-icon="search"
            >
              search
            </span>
            <input
              class="w-full pl-10 pr-4 py-2 bg-canvas border border-border-default rounded-lg text-sm focus:outline-none focus:border-primary-container focus:ring-1 focus:ring-primary-container transition-all"
              placeholder="Search by asset name, tag, or serial..."
              type="text"
            />
          </div>
          <div class="h-6 w-px bg-border-default mx-2"></div>
          <button
            class="px-3 py-1.5 bg-surface-subtle border border-border-default rounded-md text-text-primary font-body text-body hover:bg-border-default transition-colors flex items-center gap-2"
          >
            <span class="material-symbols-outlined text-[18px]" data-icon="filter_list">
              filter_list
            </span>
            Filters
          </button>
          <!-- Active Chips -->
          <div class="flex items-center gap-2">
            <span
              class="px-3 py-1 bg-surface-variant text-on-surface-variant rounded-md text-[13px] flex items-center gap-1 border border-outline-variant"
            >
              Category: IT
              <span
                class="material-symbols-outlined text-[14px] cursor-pointer hover:text-text-primary"
                data-icon="close"
              >
                close
              </span>
            </span>
            <span
              class="px-3 py-1 bg-surface-variant text-on-surface-variant rounded-md text-[13px] flex items-center gap-1 border border-outline-variant"
            >
              Dept: HR
              <span
                class="material-symbols-outlined text-[14px] cursor-pointer hover:text-text-primary"
                data-icon="close"
              >
                close
              </span>
            </span>
          </div>
        </div>
        <!-- Bulk Actions (Hidden by default, shown for context) -->
        <div class="flex items-center gap-2 hidden">
          <span class="text-sm text-text-secondary mr-2"> 2 selected </span>
          <button
            class="px-3 py-1.5 bg-surface border border-border-default rounded-md text-text-primary hover:bg-surface-subtle transition-colors text-sm"
          >
            Update Status
          </button>
          <button
            class="px-3 py-1.5 bg-surface border border-border-default rounded-md text-error hover:bg-error-container hover:border-error-container transition-colors text-sm"
          >
            Delete
          </button>
        </div>
      </div>
      <!-- Main Table Container -->
      <div
        class="bg-surface rounded-xl border border-border-default overflow-hidden shadow-[0_4px_12px_rgba(0,0,0,0.02)] p-4"
      >
        <DataTable
          :value="assets"
          :loading="loading"
          paginator
          :rows="10"
          tableStyle="min-width: 50rem"
          class="w-full text-left"
        >
          <Column header="">
            <template #body>
              <input
                class="rounded border-outline-variant text-primary-container focus:ring-primary-container"
                type="checkbox"
              />
            </template>
          </Column>
          <Column field="tagId" header="Tag ID" sortable>
            <template #body="slotProps">
              <router-link
                :to="`/asset-detail/${slotProps.data.id}`"
                class="font-mono-data text-mono-data text-primary-container hover:underline underline-offset-2 cursor-pointer"
              >
                {{ slotProps.data.tagId }}
              </router-link>
            </template>
          </Column>
          <Column field="name" header="Asset Name" sortable>
            <template #body="slotProps">
              <span class="font-medium">{{ slotProps.data.name }}</span>
            </template>
          </Column>
          <Column field="category" header="Category" sortable>
            <template #body="slotProps">
              <span class="text-text-secondary">{{ slotProps.data.category }}</span>
            </template>
          </Column>
          <Column field="status" header="Status" sortable>
            <template #body="slotProps">
              <span
                class="inline-flex items-center px-2 py-0.5 rounded-full text-[11px] font-medium border"
                :class="{
                  'bg-metric-sage text-status-in-stock border-metric-sage/50':
                    slotProps.data.status === 'In Stock',
                  'bg-tertiary-fixed text-status-checked-out border-tertiary-fixed-dim/50':
                    slotProps.data.status === 'Checked Out',
                  'bg-surface-variant text-text-secondary border-border-default':
                    slotProps.data.status === 'Reserved',
                }"
              >
                {{ slotProps.data.status }}
              </span>
            </template>
          </Column>
          <Column field="custodian" header="Custodian/Stockroom" sortable></Column>
          <Column field="location" header="Location" sortable>
            <template #body="slotProps">
              <span class="text-text-secondary">{{ slotProps.data.location }}</span>
            </template>
          </Column>
          <Column field="purchaseDate" header="Purchase Date" sortable>
            <template #body="slotProps">
              <span class="text-text-secondary">
                {{ new Date(slotProps.data.purchaseDate).toLocaleDateString() }}
              </span>
            </template>
          </Column>
          <Column field="warrantyStatus" header="Warranty" sortable>
            <template #body="slotProps">
              <span
                :class="
                  slotProps.data.warrantyStatus === 'Active'
                    ? 'text-status-in-stock font-medium'
                    : 'text-status-critical font-medium'
                "
              >
                {{ slotProps.data.warrantyStatus }}
              </span>
            </template>
          </Column>
          <Column header="Actions" bodyStyle="text-align: right">
            <template #body>
              <div class="flex items-center justify-end space-x-1">
                <button
                  class="p-1 text-text-secondary hover:text-primary-container hover:bg-surface-subtle rounded transition-colors"
                >
                  <span class="material-symbols-outlined text-[16px]">visibility</span>
                </button>
                <button
                  class="p-1 text-text-secondary hover:text-primary-container hover:bg-surface-subtle rounded transition-colors"
                >
                  <span class="material-symbols-outlined text-[16px]">edit</span>
                </button>
                <button
                  class="p-1 text-text-secondary hover:text-text-primary hover:bg-surface-subtle rounded transition-colors"
                >
                  <span class="material-symbols-outlined text-[16px]">more_vert</span>
                </button>
              </div>
            </template>
          </Column>
        </DataTable>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';

interface Asset {
  id: number;
  tagId: string;
  name: string;
  category: string;
  status: string;
  custodian: string;
  location: string;
  purchaseDate: string;
  warrantyStatus: string;
}

const assets = ref<Asset[]>([]);
const loading = ref(true);

const fetchAssets = async () => {
  try {
    const res = await fetch('http://localhost:8080/api/assets');
    if (res.ok) {
      assets.value = await res.json();
    }
  } catch (error) {
    console.error('Failed to fetch assets:', error);
  } finally {
    loading.value = false;
  }
};

onMounted(() => {
  fetchAssets();
});
</script>
