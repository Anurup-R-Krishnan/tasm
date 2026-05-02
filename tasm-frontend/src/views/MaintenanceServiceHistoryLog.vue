<template>
  <main class="space-y-section-gap pb-24">
    <div class="mb-section-gap flex justify-between items-end">
      <div>
        <h1 class="font-h1 text-h1 text-text-primary mb-1">Maintenance History</h1>
        <p class="font-body text-body text-text-secondary">
          Track work orders, service logs, and maintenance costs.
        </p>
      </div>
      <div class="flex gap-3">
        <button
          class="px-4 py-2 bg-surface border border-border-default rounded-lg text-text-primary font-medium hover:-translate-y-0.5 transition-transform shadow-sm flex items-center gap-2"
        >
          <span class="material-symbols-outlined text-[18px]">download</span>
          Export CSV
        </button>
      </div>
    </div>

    <!-- KPI Cards -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-section-gap">
      <div
        class="bg-surface rounded-xl border border-border-default p-card-padding shadow-sm hover:-translate-y-0.5 transition-transform"
      >
        <div class="flex justify-between items-start mb-4">
          <p class="font-metadata text-metadata text-text-secondary uppercase tracking-wider">
            Total Maint. Spend YTD
          </p>
          <div class="p-2 bg-metric-amber rounded-lg text-primary-container">
            <span class="material-symbols-outlined">payments</span>
          </div>
        </div>
        <p class="font-kpi-number text-kpi-number text-text-primary font-mono-data mb-1">
          ₹{{ totalSpend.toLocaleString() }}
        </p>
        <p class="font-metadata text-metadata text-status-critical flex items-center gap-1">
          <span class="material-symbols-outlined text-[14px]">trending_up</span>
          +12.5% vs last year
        </p>
      </div>
      <div
        class="bg-surface rounded-xl border border-border-default p-card-padding shadow-sm hover:-translate-y-0.5 transition-transform"
      >
        <div class="flex justify-between items-start mb-4">
          <p class="font-metadata text-metadata text-text-secondary uppercase tracking-wider">
            Completed Work Orders
          </p>
          <div class="p-2 bg-metric-lavender rounded-lg text-tertiary-container">
            <span class="material-symbols-outlined">build_circle</span>
          </div>
        </div>
        <p class="font-kpi-number text-kpi-number text-text-primary">{{ logs.length }}</p>
        <p class="font-metadata text-metadata text-text-secondary flex items-center gap-1">
          <span class="material-symbols-outlined text-[14px]">schedule</span>
          Recent activity tracked
        </p>
      </div>
      <div
        class="bg-surface rounded-xl border border-border-default p-card-padding shadow-sm hover:-translate-y-0.5 transition-transform"
      >
        <div class="flex justify-between items-start mb-4">
          <p class="font-metadata text-metadata text-text-secondary uppercase tracking-wider">
            Avg Resolution Time
          </p>
          <div class="p-2 bg-metric-sage rounded-lg text-status-in-stock">
            <span class="material-symbols-outlined">timer</span>
          </div>
        </div>
        <p class="font-kpi-number text-kpi-number text-text-primary">3.2 Days</p>
        <p class="font-metadata text-metadata text-status-in-stock flex items-center gap-1">
          <span class="material-symbols-outlined text-[14px]">trending_down</span>
          -0.5 days vs last month
        </p>
      </div>
    </div>

    <!-- Content Grid: Table & Form Split -->
    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <!-- Table Section -->
      <div
        class="lg:col-span-2 bg-surface rounded-xl border border-border-default shadow-sm overflow-hidden flex flex-col h-full"
      >
        <div
          class="p-4 border-b border-border-default bg-surface-subtle flex justify-between items-center"
        >
          <h3 class="font-h3 text-h3 text-text-primary">Service Logs</h3>
          <div class="flex gap-2">
            <div class="relative">
              <span
                class="material-symbols-outlined absolute left-2.5 top-2 text-stone-400 text-[18px]"
                >filter_list</span
              >
              <select
                v-model="filterType"
                class="pl-9 pr-8 py-1.5 bg-white border border-border-default rounded-md text-sm text-text-primary appearance-none focus:ring-1 focus:ring-amber-600 focus:border-amber-600 outline-none cursor-pointer"
              >
                <option value="">All Types</option>
                <option value="Preventive">Preventive</option>
                <option value="Corrective">Corrective</option>
              </select>
            </div>
          </div>
        </div>
        <div class="overflow-x-auto flex-1 p-4">
          <DataTable :value="filteredLogs" :loading="loading" paginator :rows="5" class="w-full">
            <template #empty>No completed service logs found.</template>
            <Column field="workOrderId" header="WO ID" sortable>
              <template #body="slotProps">
                <span class="font-mono-data text-mono-data font-medium text-text-primary">{{
                  slotProps.data.workOrderId
                }}</span>
              </template>
            </Column>
            <Column field="targetDate" header="Date" sortable>
              <template #body="slotProps">
                <span class="text-text-secondary">{{ formatDate(slotProps.data.targetDate) }}</span>
              </template>
            </Column>
            <Column field="title" header="Asset & Type" sortable>
              <template #body="slotProps">
                <div class="text-text-primary font-medium">{{ slotProps.data.assetLocation }}</div>
                <div class="text-xs text-text-secondary mt-0.5">{{ slotProps.data.title }}</div>
              </template>
            </Column>
            <Column field="technician" header="Technician" sortable>
              <template #body="slotProps">
                <span class="text-text-secondary">{{
                  slotProps.data.technician || 'Internal Team'
                }}</span>
              </template>
            </Column>
            <Column field="cost" header="Cost (₹)" sortable>
              <template #body="slotProps">
                <span class="font-mono-data text-right text-text-primary">{{
                  slotProps.data.cost?.toLocaleString() || '0'
                }}</span>
              </template>
            </Column>
            <Column field="status" header="Status">
              <template #body="slotProps">
                <span
                  class="inline-flex items-center px-2 py-1 rounded-md text-xs font-medium"
                  :class="
                    slotProps.data.status === 'Closed'
                      ? 'bg-metric-sage text-status-in-stock'
                      : 'bg-surface-variant text-text-secondary'
                  "
                >
                  {{ slotProps.data.status === 'Closed' ? 'Completed' : slotProps.data.status }}
                </span>
              </template>
            </Column>
          </DataTable>
        </div>
      </div>

      <!-- Add Log Form Section -->
      <div
        class="lg:col-span-1 bg-surface rounded-xl border border-border-default shadow-sm h-full flex flex-col"
      >
        <div class="p-4 border-b border-border-default bg-surface-subtle">
          <h3 class="font-h3 text-h3 text-text-primary flex items-center gap-2">
            <span class="material-symbols-outlined text-[20px] text-amber-700">post_add</span>
            Quick Add Log
          </h3>
        </div>
        <form @submit.prevent="submitLog" class="p-6 flex-1 flex flex-col gap-5">
          <div class="space-y-1">
            <label class="font-table-header text-table-header text-text-secondary uppercase"
              >Asset ID / Name</label
            >
            <input
              v-model="form.assetLocation"
              required
              class="w-full px-3 py-2 border border-border-default rounded-lg focus:ring-1 focus:ring-amber-600 focus:border-amber-600 outline-none text-sm text-text-primary"
              placeholder="e.g. HVAC-B3"
              type="text"
            />
          </div>
          <div class="grid grid-cols-2 gap-4">
            <div class="space-y-1">
              <label class="font-table-header text-table-header text-text-secondary uppercase"
                >Service Date</label
              >
              <input
                v-model="form.targetDate"
                required
                class="w-full px-3 py-2 border border-border-default rounded-lg focus:ring-1 focus:ring-amber-600 focus:border-amber-600 outline-none text-sm text-text-primary"
                type="date"
              />
            </div>
            <div class="space-y-1">
              <label class="font-table-header text-table-header text-text-secondary uppercase"
                >Service Type</label
              >
              <select
                v-model="form.title"
                class="w-full px-3 py-2 border border-border-default rounded-lg focus:ring-1 focus:ring-amber-600 focus:border-amber-600 outline-none text-sm text-text-primary appearance-none bg-white"
              >
                <option value="Preventive Maintenance">Preventive</option>
                <option value="Corrective Maintenance">Corrective</option>
                <option value="Emergency Repair">Emergency</option>
              </select>
            </div>
          </div>
          <div class="space-y-1">
            <label class="font-table-header text-table-header text-text-secondary uppercase"
              >Technician / Vendor</label
            >
            <input
              v-model="form.technician"
              class="w-full px-3 py-2 border border-border-default rounded-lg focus:ring-1 focus:ring-amber-600 focus:border-amber-600 outline-none text-sm text-text-primary"
              placeholder="Name or Company"
              type="text"
            />
          </div>
          <div class="space-y-1">
            <label class="font-table-header text-table-header text-text-secondary uppercase"
              >Total Cost (₹)</label
            >
            <div class="relative">
              <span class="absolute left-3 top-2 text-text-secondary font-mono-data">₹</span>
              <input
                v-model.number="form.cost"
                class="w-full pl-8 pr-3 py-2 border border-border-default rounded-lg focus:ring-1 focus:ring-amber-600 focus:border-amber-600 outline-none text-sm font-mono-data text-text-primary"
                placeholder="0.00"
                type="number"
              />
            </div>
          </div>
          <div class="space-y-1">
            <label class="font-table-header text-table-header text-text-secondary uppercase"
              >Notes</label
            >
            <textarea
              v-model="form.issue"
              class="w-full px-3 py-2 border border-border-default rounded-lg focus:ring-1 focus:ring-amber-600 focus:border-amber-600 outline-none text-sm text-text-primary h-24 resize-none"
              placeholder="Details of work performed..."
            ></textarea>
          </div>
          <div class="mt-auto pt-4 flex gap-3">
            <button
              type="button"
              @click="resetForm"
              class="flex-1 px-4 py-2 bg-surface border border-border-default rounded-lg text-text-primary font-medium hover:bg-stone-50 transition-colors"
            >
              Clear
            </button>
            <button
              type="submit"
              :disabled="submitting"
              class="flex-1 px-4 py-2 bg-[#1C1917] rounded-lg text-white font-medium hover:bg-black transition-colors shadow-sm disabled:opacity-50"
            >
              Save Log
            </button>
          </div>
        </form>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';

interface WorkOrder {
  id?: number;
  workOrderId: string;
  title: string;
  assetLocation: string;
  issue: string;
  severity: string;
  targetDate: string;
  status: string;
  technician?: string;
  cost?: number;
}

const logs = ref<WorkOrder[]>([]);
const loading = ref(true);
const submitting = ref(false);
const filterType = ref('');

const form = ref({
  assetLocation: '',
  targetDate: new Date().toISOString().split('T')[0] || '',
  title: 'Preventive Maintenance',
  technician: '',
  cost: 0,
  issue: '',
});

const filteredLogs = computed(() => {
  if (!filterType.value) return logs.value;
  return logs.value.filter((log) => log.title.includes(filterType.value));
});

const totalSpend = computed(() => {
  return logs.value.reduce((acc, log) => acc + (log.cost || 0), 0);
});

const fetchLogs = async () => {
  try {
    const res = await fetch('http://localhost:8080/api/work-orders');
    if (res.ok) {
      const data: WorkOrder[] = await res.json();
      // Only show closed work orders in the history log
      logs.value = data.filter((wo) => wo.status === 'Closed');
    }
  } catch (error) {
    console.error('Failed to load logs', error);
  } finally {
    loading.value = false;
  }
};

const submitLog = async () => {
  submitting.value = true;
  const payload: WorkOrder = {
    workOrderId: `WO-${new Date().getFullYear()}-${Math.floor(Math.random() * 10000)}`,
    title: form.value.title,
    assetLocation: form.value.assetLocation,
    issue: form.value.issue,
    severity: 'Medium',
    targetDate: new Date(form.value.targetDate).toISOString(),
    status: 'Closed',
    technician: form.value.technician,
    cost: form.value.cost,
  };

  try {
    const res = await fetch('http://localhost:8080/api/work-orders', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload),
    });

    if (res.ok) {
      resetForm();
      await fetchLogs();
    }
  } catch (err) {
    console.error('Failed to submit log', err);
  } finally {
    submitting.value = false;
  }
};

const resetForm = () => {
  form.value = {
    assetLocation: '',
    targetDate: new Date().toISOString().split('T')[0] || '',
    title: 'Preventive Maintenance',
    technician: '',
    cost: 0,
    issue: '',
  };
};

const formatDate = (dateStr: string) => {
  return new Date(dateStr).toLocaleDateString(undefined, {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
  });
};

onMounted(() => {
  fetchLogs();
});
</script>
