<template>
  <main class="md:ml-[248px] min-h-screen p-page-margin max-w-[1400px]">
    <!-- Page Header -->
    <div class="flex justify-between items-center mb-section-gap">
      <div>
        <h2 class="text-h1 font-h1 text-text-primary mb-1">Maintenance Schedules & Work Orders</h2>
        <p class="text-body font-body text-text-secondary">
          Track and manage upcoming asset maintenance across facilities.
        </p>
      </div>
      <button
        class="bg-text-primary text-surface px-5 py-2.5 rounded-lg text-body font-body font-medium hover:bg-opacity-90 transition-colors shadow-sm flex items-center gap-2"
      >
        <span class="material-symbols-outlined text-[20px]"> calendar_add_on </span>
        Add Schedule
      </button>
    </div>

    <!-- Filters & Controls Bar -->
    <div
      class="bg-surface p-4 rounded-xl border border-border-default shadow-[0_4px_4px_rgba(0,0,0,0.02)] mb-stack flex flex-wrap gap-4 items-center justify-between"
    >
      <div class="flex gap-4">
        <div class="relative">
          <select
            v-model="selectedStatus"
            class="appearance-none bg-surface-subtle border border-border-default text-body font-body text-text-primary rounded-md pl-4 pr-10 py-2 focus:outline-none focus:border-primary-container"
          >
            <option value="">All Statuses</option>
            <option value="Open">Open</option>
            <option value="In Progress">In Progress</option>
            <option value="Closed">Closed</option>
          </select>
          <span
            class="material-symbols-outlined absolute right-3 top-1/2 -translate-y-1/2 text-text-secondary pointer-events-none text-[20px]"
          >
            expand_more
          </span>
        </div>
        <div class="relative w-64">
          <span
            class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-text-secondary pointer-events-none text-[18px]"
          >
            search
          </span>
          <input
            v-model="searchQuery"
            class="w-full bg-surface-subtle border border-border-default text-body font-body text-text-primary rounded-md pl-10 pr-4 py-2 focus:outline-none focus:border-primary-container"
            placeholder="Search by ID or Title..."
          />
        </div>
      </div>
      <div class="flex gap-2">
        <button
          class="p-2 text-text-secondary hover:text-text-primary border border-transparent hover:border-border-default rounded-md transition-all"
        >
          <span class="material-symbols-outlined">filter_list</span>
        </button>
        <button
          class="p-2 text-text-secondary hover:text-text-primary border border-transparent hover:border-border-default rounded-md transition-all"
        >
          <span class="material-symbols-outlined">download</span>
        </button>
      </div>
    </div>

    <!-- High-Density Table Card -->
    <div
      class="bg-surface rounded-xl border border-border-default shadow-[0_4px_4px_rgba(0,0,0,0.02)] overflow-hidden p-4"
    >
      <DataTable :value="filteredWorkOrders" :loading="loading" paginator :rows="10" class="w-full">
        <template #empty> No work orders or schedules found. </template>
        <Column field="workOrderId" header="Work Order ID" sortable>
          <template #body="slotProps">
            <span class="font-mono-data text-mono-data text-primary-container font-medium">{{
              slotProps.data.workOrderId
            }}</span>
          </template>
        </Column>
        <Column field="title" header="Maintenance Task" sortable>
          <template #body="slotProps">
            <div>
              <div class="font-medium text-text-primary">{{ slotProps.data.title }}</div>
              <div class="text-xs text-text-secondary truncate max-w-[200px]">
                {{ slotProps.data.issue }}
              </div>
            </div>
          </template>
        </Column>
        <Column field="assetLocation" header="Location" sortable></Column>
        <Column field="severity" header="Severity" sortable>
          <template #body="slotProps">
            <span
              class="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium"
              :class="{
                'bg-error-container text-on-error-container':
                  slotProps.data.severity === 'High' || slotProps.data.severity === 'Critical',
                'bg-metric-amber text-amber-900': slotProps.data.severity === 'Medium',
                'bg-surface-variant text-text-secondary': slotProps.data.severity === 'Low',
              }"
            >
              {{ slotProps.data.severity }}
            </span>
          </template>
        </Column>
        <Column field="targetDate" header="Target Date" sortable>
          <template #body="slotProps">
            <span
              class="font-medium"
              :class="
                isOverdue(slotProps.data.targetDate, slotProps.data.status)
                  ? 'text-error font-bold'
                  : 'text-text-primary'
              "
            >
              {{ formatDate(slotProps.data.targetDate) }}
            </span>
          </template>
        </Column>
        <Column field="status" header="Status" sortable>
          <template #body="slotProps">
            <span
              class="inline-flex items-center px-2 py-0.5 rounded-full text-[11px] font-medium border"
              :class="{
                'bg-tertiary-container text-on-tertiary border-tertiary':
                  slotProps.data.status === 'In Progress',
                'bg-metric-sage text-status-in-stock border-metric-sage/50':
                  slotProps.data.status === 'Closed',
                'bg-surface-variant text-text-secondary border-border-default':
                  slotProps.data.status === 'Open',
              }"
            >
              {{ slotProps.data.status }}
            </span>
          </template>
        </Column>
        <Column header="">
          <template #body>
            <button
              class="p-1 text-text-secondary hover:text-primary-container hover:bg-surface-subtle rounded transition-colors"
            >
              <span class="material-symbols-outlined text-[18px]">more_vert</span>
            </button>
          </template>
        </Column>
      </DataTable>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';

interface WorkOrder {
  id: number;
  workOrderId: string;
  title: string;
  assetLocation: string;
  issue: string;
  severity: string;
  targetDate: string;
  status: string;
}

const workOrders = ref<WorkOrder[]>([]);
const loading = ref(true);
const searchQuery = ref('');
const selectedStatus = ref('');

const filteredWorkOrders = computed(() => {
  return workOrders.value.filter((wo) => {
    const matchesSearch =
      !searchQuery.value ||
      wo.title.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      wo.workOrderId.toLowerCase().includes(searchQuery.value.toLowerCase());

    const matchesStatus = !selectedStatus.value || wo.status === selectedStatus.value;

    return matchesSearch && matchesStatus;
  });
});

const fetchWorkOrders = async () => {
  try {
    const res = await fetch('http://localhost:8080/api/work-orders');
    if (res.ok) {
      workOrders.value = await res.json();
    }
  } catch (error) {
    console.error('Failed to load work orders', error);
  } finally {
    loading.value = false;
  }
};

const formatDate = (dateStr: string) => {
  return new Date(dateStr).toLocaleDateString(undefined, {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
  });
};

const isOverdue = (dateStr: string, status: string) => {
  if (status === 'Closed') return false;
  const target = new Date(dateStr);
  const now = new Date();
  return target < now;
};

onMounted(() => {
  fetchWorkOrders();
});
</script>
