<template>
  <main class="p-page-margin">
    <div class="max-w-[1400px] mx-auto space-y-section-gap">
      <!-- Top Row: Welcome Banner & Global Search -->
      <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
        <div>
          <h1 class="font-h1 text-h1 text-text-primary">Global Operations Dashboard</h1>
          <p class="font-metadata text-metadata text-text-secondary mt-1">
            Overview of system assets, financials, and pending tasks.
          </p>
        </div>
        <div class="flex gap-3">
          <button class="btn-primary">
            <span class="material-symbols-outlined">add_circle</span>
            New Entry
          </button>
        </div>
      </div>

      <!-- KPI Cards -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div class="premium-card flex items-start justify-between">
          <div>
            <p class="font-metadata text-metadata text-text-secondary uppercase">Total Assets</p>
            <h2 class="font-kpi-number text-kpi-number text-text-primary mt-2">
              {{ assets.length }}
            </h2>
            <p
              class="text-sm mt-1"
              :class="activeAssets > 0 ? 'text-status-in-stock' : 'text-text-secondary'"
            >
              {{ activeAssets }} Active
            </p>
          </div>
          <div
            class="w-10 h-10 rounded-full bg-metric-sage/20 flex items-center justify-center text-status-in-stock"
          >
            <span class="material-symbols-outlined">devices</span>
          </div>
        </div>

        <div class="premium-card flex items-start justify-between">
          <div>
            <p class="font-metadata text-metadata text-text-secondary uppercase">
              Open Work Orders
            </p>
            <h2 class="font-kpi-number text-kpi-number text-text-primary mt-2">
              {{ openWorkOrders }}
            </h2>
            <p class="text-sm mt-1 text-metric-amber font-medium">Action Required</p>
          </div>
          <div
            class="w-10 h-10 rounded-full bg-metric-amber/20 flex items-center justify-center text-amber-700"
          >
            <span class="material-symbols-outlined">build</span>
          </div>
        </div>

        <div
          class="bg-surface rounded-xl p-card-padding border border-border-default shadow-sm flex items-start justify-between hover:-translate-y-0.5 transition-transform"
        >
          <div>
            <p class="font-metadata text-metadata text-text-secondary uppercase">
              Procurement Pipeline
            </p>
            <h2 class="font-kpi-number text-kpi-number text-text-primary mt-2">
              ₹{{ totalProcurement.toLocaleString() }}
            </h2>
            <p class="text-sm mt-1 text-text-secondary">Pending approvals</p>
          </div>
          <div
            class="w-10 h-10 rounded-full bg-metric-lavender/20 flex items-center justify-center text-purple-700"
          >
            <span class="material-symbols-outlined">inventory</span>
          </div>
        </div>

        <div
          class="bg-surface rounded-xl p-card-padding border border-border-default shadow-sm flex items-start justify-between hover:-translate-y-0.5 transition-transform"
        >
          <div>
            <p class="font-metadata text-metadata text-text-secondary uppercase">
              Total Expenses YTD
            </p>
            <h2 class="font-kpi-number text-kpi-number text-text-primary mt-2">
              ₹{{ totalExpenses.toLocaleString() }}
            </h2>
            <p class="text-sm mt-1 text-status-critical">From ledger entries</p>
          </div>
          <div
            class="w-10 h-10 rounded-full bg-error-container/50 flex items-center justify-center text-status-critical"
          >
            <span class="material-symbols-outlined">payments</span>
          </div>
        </div>
      </div>

      <!-- Recent Activity / Asset Preview -->
      <div
        class="bg-surface rounded-xl border border-border-default shadow-sm overflow-hidden flex flex-col"
      >
        <div
          class="p-4 border-b border-border-default bg-surface flex justify-between items-center"
        >
          <h2 class="font-h2 text-h2 text-text-primary flex items-center gap-2">
            <span class="material-symbols-outlined text-lg">history</span>
            Recent Assets Registered
          </h2>
        </div>
        <div class="overflow-x-auto p-4">
          <DataTable
            :value="recentAssets"
            :loading="loading"
            paginator
            :rows="5"
            class="w-full text-left"
          >
            <Column field="name" header="Asset Name" sortable></Column>
            <Column field="tagId" header="Tag ID" sortable>
              <template #body="slotProps">
                <span class="font-mono-data text-mono-data text-surface-tint">{{
                  slotProps.data.tagId
                }}</span>
              </template>
            </Column>
            <Column field="status" header="Status" sortable>
              <template #body="slotProps">
                <span
                  class="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium bg-metric-sage/20 text-status-in-stock"
                >
                  {{ slotProps.data.status }}
                </span>
              </template>
            </Column>
            <Column field="purchaseDate" header="Registered Date" sortable>
              <template #body="slotProps">
                <span class="font-mono-data text-mono-data">{{
                  new Date(slotProps.data.purchaseDate).toLocaleDateString()
                }}</span>
              </template>
            </Column>
          </DataTable>
        </div>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';

const assets = ref<any[]>([]);
const workOrders = ref<any[]>([]);
const procurements = ref<any[]>([]);
const ledgers = ref<any[]>([]);
const loading = ref(true);

const activeAssets = computed(
  () => assets.value.filter((a) => a.status === 'Active' || a.status === 'In Use').length,
);
const openWorkOrders = computed(
  () => workOrders.value.filter((w) => w.status !== 'Completed').length,
);
const totalProcurement = computed(() =>
  procurements.value.reduce((sum, p) => sum + (p.estimatedValue || 0), 0),
);
const totalExpenses = computed(() =>
  ledgers.value.filter((l) => l.type === 'Debit').reduce((sum, l) => sum + (l.amount || 0), 0),
);

const recentAssets = computed(() => {
  return [...assets.value]
    .sort((a, b) => new Date(b.purchaseDate).getTime() - new Date(a.purchaseDate).getTime())
    .slice(0, 5);
});

const fetchDashboardData = async () => {
  loading.value = true;
  try {
    const [assetsRes, workOrdersRes, procurementsRes, ledgersRes] = await Promise.all([
      fetch('http://localhost:8080/api/assets'),
      fetch('http://localhost:8080/api/work-orders'),
      fetch('http://localhost:8080/api/procurements'),
      fetch('http://localhost:8080/api/ledgers'),
    ]);

    if (assetsRes.ok) assets.value = await assetsRes.json();
    if (workOrdersRes.ok) workOrders.value = await workOrdersRes.json();
    if (procurementsRes.ok) procurements.value = await procurementsRes.json();
    if (ledgersRes.ok) ledgers.value = await ledgersRes.json();
  } catch (error) {
    console.error('Failed to fetch dashboard data:', error);
  } finally {
    loading.value = false;
  }
};

onMounted(() => {
  fetchDashboardData();
});
</script>
