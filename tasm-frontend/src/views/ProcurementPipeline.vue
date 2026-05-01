<template>
  <main class="flex-1 flex flex-col md:ml-[248px] bg-canvas overflow-hidden">
    <!-- Header & Financial Summary -->
    <div
      class="flex-none p-page-margin pb-6 border-b border-border-default bg-canvas z-10 shadow-sm"
    >
      <div class="flex items-end justify-between mb-section-gap">
        <div>
          <h1 class="font-h1 text-h1 text-text-primary mb-2">Procurement Pipeline</h1>
          <p class="font-body text-body text-text-secondary">
            Track and manage active procurement workflows.
          </p>
        </div>
        <div class="flex gap-inline">
          <button
            class="px-4 py-2 bg-surface border border-outline-variant text-text-primary font-h3 text-h3 rounded-lg hover:bg-surface-subtle transition-colors shadow-sm"
          >
            Export Report
          </button>
          <button
            class="px-4 py-2 bg-text-primary text-on-primary font-h3 text-h3 rounded-lg hover:bg-stone-800 transition-colors shadow-sm flex items-center gap-2"
          >
            <span class="material-symbols-outlined text-[18px]"> add </span>
            New Request
          </button>
        </div>
      </div>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <!-- Metric Card 1 -->
        <div
          class="bg-surface rounded-xl p-card-padding border border-border-default shadow-[0_4px_12px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform cursor-pointer group"
        >
          <div class="flex justify-between items-start mb-4">
            <span class="font-table-header text-table-header text-text-secondary uppercase">
              Total Pipeline Value
            </span>
            <div
              class="w-8 h-8 rounded-full bg-secondary-container flex items-center justify-center text-text-primary group-hover:bg-primary-container group-hover:text-on-primary-container transition-colors"
            >
              <span class="material-symbols-outlined text-[18px]"> account_balance_wallet </span>
            </div>
          </div>
          <div class="font-kpi-number text-kpi-number text-text-primary">
            ₹ {{ totalPipelineValue.toLocaleString() }}
          </div>
          <div class="font-metadata text-metadata text-text-secondary mt-2 flex items-center gap-1">
            <span class="text-status-in-stock flex items-center">
              <span class="material-symbols-outlined text-[14px]"> trending_up </span>
              12%
            </span>
            vs last month
          </div>
        </div>
        <!-- Metric Card 2 -->
        <div
          class="bg-surface rounded-xl p-card-padding border border-border-default shadow-[0_4px_12px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform cursor-pointer group"
        >
          <div class="flex justify-between items-start mb-4">
            <span class="font-table-header text-table-header text-text-secondary uppercase">
              Approved This Month
            </span>
            <div
              class="w-8 h-8 rounded-full bg-metric-sage/50 flex items-center justify-center text-status-in-stock transition-colors"
            >
              <span class="material-symbols-outlined text-[18px]"> check_circle </span>
            </div>
          </div>
          <div class="font-kpi-number text-kpi-number text-text-primary">
            ₹ {{ approvedValue.toLocaleString() }}
          </div>
          <div class="font-metadata text-metadata text-text-secondary mt-2">
            Across
            {{
              procurements.filter((p) => p.status !== 'Draft' && p.status !== 'Pending Approval')
                .length
            }}
            purchase orders
          </div>
        </div>
        <!-- Metric Card 3 (Placeholder for balance visually) -->
        <div
          class="bg-surface rounded-xl p-card-padding border border-border-default shadow-[0_4px_12px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform cursor-pointer group"
        >
          <div class="flex justify-between items-start mb-4">
            <span class="font-table-header text-table-header text-text-secondary uppercase">
              Pending Approval Value
            </span>
            <div
              class="w-8 h-8 rounded-full bg-metric-amber/50 flex items-center justify-center text-primary transition-colors"
            >
              <span class="material-symbols-outlined text-[18px]"> pending_actions </span>
            </div>
          </div>
          <div class="font-kpi-number text-kpi-number text-text-primary">
            ₹ {{ pendingApprovalValue.toLocaleString() }}
          </div>
          <div class="font-metadata text-metadata text-text-secondary mt-2 text-primary">
            Requires management review
          </div>
        </div>
      </div>
    </div>
    <!-- Data Table Area -->
    <div class="flex-1 overflow-auto p-page-margin">
      <div class="bg-surface border border-border-default rounded-xl shadow-sm overflow-hidden">
        <DataTable
          :value="procurements"
          :loading="loadingProcurements"
          paginator
          :rows="10"
          tableStyle="min-width: 50rem"
          class="w-full text-left"
        >
          <Column field="title" header="Request Title" sortable>
            <template #body="slotProps">
              <div>
                <h3 class="font-h3 text-h3 text-text-primary mb-1">{{ slotProps.data.title }}</h3>
                <span class="font-metadata text-metadata text-text-secondary">{{
                  slotProps.data.poNumber ? slotProps.data.poNumber : 'No PO assigned'
                }}</span>
              </div>
            </template>
          </Column>
          <Column field="status" header="Status" sortable>
            <template #body="slotProps">
              <span
                class="inline-flex items-center px-2 py-1 rounded font-metadata text-metadata font-medium"
                :class="{
                  'bg-metric-amber/20 text-[#854d0e]': slotProps.data.status === 'Draft',
                  'bg-surface-variant text-on-surface-variant':
                    slotProps.data.status === 'Pending Approval',
                  'bg-[#ddd6fe] text-[#8b5cf6]': slotProps.data.status === 'PO Raised',
                  'bg-blue-100 text-blue-700': slotProps.data.status === 'Shipping',
                  'bg-metric-sage text-status-in-stock': slotProps.data.status === 'Received',
                }"
              >
                {{ slotProps.data.status }}
              </span>
            </template>
          </Column>
          <Column field="priority" header="Priority" sortable>
            <template #body="slotProps">
              <div
                class="px-2 py-1 rounded text-[11px] font-semibold tracking-wider uppercase inline-block"
                :class="{
                  'bg-[#fef08a] text-[#854d0e]': slotProps.data.priority === 'Low',
                  'bg-surface-variant text-on-surface-variant':
                    slotProps.data.priority === 'Medium',
                  'bg-error-container text-status-critical': slotProps.data.priority === 'High',
                }"
              >
                {{ slotProps.data.priority }}
              </div>
            </template>
          </Column>
          <Column field="estimatedValue" header="Est. Value" sortable>
            <template #body="slotProps">
              <span class="font-mono-data text-mono-data text-text-primary">
                ₹ {{ slotProps.data.estimatedValue?.toLocaleString() || '0' }}
              </span>
            </template>
          </Column>
          <Column field="department" header="Department" sortable>
            <template #body="slotProps">
              <div class="flex items-center gap-2">
                <div
                  class="w-6 h-6 rounded-full bg-surface-variant flex items-center justify-center text-[10px] font-bold text-text-primary border border-border-default"
                >
                  {{ slotProps.data.requestorInitials }}
                </div>
                <span class="font-metadata text-metadata text-text-secondary">
                  {{ slotProps.data.department }}
                </span>
              </div>
            </template>
          </Column>
          <Column header="Action" alignFrozen="right">
            <template #body>
              <button class="text-primary hover:text-amber-800 font-medium transition-colors">
                Review
              </button>
            </template>
          </Column>
        </DataTable>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';

interface ProcurementRequest {
  id: number;
  title: string;
  status: string;
  priority: string;
  estimatedValue: number;
  actualValue: number;
  requestorName: string;
  requestorInitials: string;
  department: string;
  poNumber: string;
  createdAt: string;
}

const procurements = ref<ProcurementRequest[]>([]);
const loadingProcurements = ref(true);

const totalPipelineValue = computed(() => {
  return procurements.value.reduce((sum, item) => sum + (item.estimatedValue || 0), 0);
});

const pendingApprovalValue = computed(() => {
  return procurements.value
    .filter((p) => p.status === 'Pending Approval')
    .reduce((sum, item) => sum + (item.estimatedValue || 0), 0);
});

const approvedValue = computed(() => {
  return procurements.value
    .filter((p) => p.status !== 'Draft' && p.status !== 'Pending Approval')
    .reduce((sum, item) => sum + (item.estimatedValue || 0), 0);
});

const fetchProcurements = async () => {
  try {
    const res = await fetch('http://localhost:8080/api/procurements');
    if (res.ok) {
      procurements.value = await res.json();
    }
  } catch (error) {
    console.error('Failed to fetch procurements:', error);
  } finally {
    loadingProcurements.value = false;
  }
};

onMounted(() => {
  fetchProcurements();
});
</script>
