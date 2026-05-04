<template>
  <main class="space-y-section-gap pb-24">
    <!-- Header -->
    <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
      <div>
        <h1 class="text-text-primary mb-1">Procurement Pipeline</h1>
        <p class="text-text-secondary">
          Track and manage active procurement workflows across all departments.
        </p>
      </div>
      <div class="flex gap-3">
        <button
          @click="exportToCSV"
          class="bg-surface border border-border-default text-text-primary px-4 py-2 rounded-lg text-sm flex items-center gap-2 hover:bg-surface-subtle transition-colors shadow-sm"
        >
          <span class="material-symbols-outlined text-[18px]">file_download</span>
          Export Report
        </button>
        <button @click="router.push('/add-new-asset-form')" class="btn-primary">
          <span class="material-symbols-outlined">add_shopping_cart</span>
          New Request
        </button>
      </div>
    </div>

    <!-- Financial Summary Cards -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
      <div v-for="kpi in kpis" :key="kpi.label" class="premium-card">
        <div class="flex justify-between items-start mb-4">
          <div
            class="w-10 h-10 rounded-xl flex items-center justify-center shadow-sm"
            :class="kpi.bgClass"
          >
            <span class="material-symbols-outlined" :class="kpi.iconClass">{{ kpi.icon }}</span>
          </div>
          <div
            class="flex items-center gap-1 text-[10px] font-bold text-status-in-stock bg-status-in-stock/10 px-2 py-0.5 rounded-full border border-status-in-stock/20"
          >
            <span class="material-symbols-outlined text-xs">trending_up</span>
            12%
          </div>
        </div>
        <p class="text-[10px] font-bold text-text-secondary uppercase tracking-widest">
          {{ kpi.label }}
        </p>
        <h2 class="text-2xl font-bold text-text-primary mt-1">₹{{ kpi.value.toLocaleString() }}</h2>
        <p class="text-xs text-text-secondary mt-4">{{ kpi.subtext }}</p>
      </div>
    </div>

    <!-- Data Table -->
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
              placeholder="Search requests..."
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
              <th class="px-6 py-4">Request Detail</th>
              <th class="px-6 py-4">Status</th>
              <th class="px-6 py-4">Priority</th>
              <th class="px-6 py-4">Estimated Value</th>
              <th class="px-6 py-4">Department</th>
              <th class="px-6 py-4 text-right">Actions</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-50">
            <tr
              v-for="req in filteredProcurements"
              :key="req.id"
              class="hover:bg-surface-subtle/50 transition-colors group"
            >
              <td class="px-6 py-4">
                <div class="flex flex-col">
                  <span
                    class="text-sm font-bold text-text-primary leading-tight group-hover:text-primary transition-colors"
                    >{{ req.title }}</span
                  >
                  <span class="text-[10px] font-mono text-text-secondary mt-1 tracking-wider">{{
                    req.poNumber || 'No PO Assigned'
                  }}</span>
                </div>
              </td>
              <td class="px-6 py-4">
                <span
                  class="text-[10px] font-bold uppercase tracking-wider px-2 py-0.5 rounded-full"
                  :class="getStatusClass(req.status)"
                >
                  {{ req.status }}
                </span>
              </td>
              <td class="px-6 py-4">
                <span
                  class="text-[10px] font-bold uppercase tracking-wider px-2 py-0.5 rounded-full"
                  :class="getPriorityClass(req.priority)"
                >
                  {{ req.priority }}
                </span>
              </td>
              <td class="px-6 py-4">
                <span class="text-sm font-bold text-text-primary"
                  >₹{{ (req.estimatedValue || 0).toLocaleString() }}</span
                >
              </td>
              <td class="px-6 py-4">
                <div class="flex items-center gap-2">
                  <div
                    class="w-6 h-6 rounded-full bg-surface-variant flex items-center justify-center text-[9px] font-bold text-text-secondary uppercase border border-border-default"
                  >
                    {{ req.requestorInitials }}
                  </div>
                  <span class="text-xs text-text-secondary font-medium">{{ req.department }}</span>
                </div>
              </td>
              <td class="px-6 py-4 text-right">
                <button
                  @click="router.push(`/procurement/${req.id}`)"
                  class="p-2 text-text-secondary hover:text-primary hover:bg-white rounded-lg transition-all shadow-sm border border-transparent hover:border-border-default"
                >
                  <span class="material-symbols-outlined text-[20px]">visibility</span>
                </button>
              </td>
            </tr>
            <tr v-if="filteredProcurements.length === 0">
              <td colspan="6" class="px-6 py-12 text-center text-text-secondary text-sm italic">
                No procurement requests found.
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { getProcurements } from '../api/procurements';

const router = useRouter();

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

const requests = ref<ProcurementRequest[]>([]);
const loading = ref(true);
const searchQuery = ref('');

const fetchRequests = async () => {
  try {
    requests.value = (await getProcurements()) as any[];
  } catch (error) {
    console.error('Failed to fetch procurements:', error);
  } finally {
    loading.value = false;
  }
};

const filteredProcurements = computed(() => {
  const query = searchQuery.value.toLowerCase();
  return requests.value.filter(
    (p: any) =>
      p.title.toLowerCase().includes(query) ||
      p.department.toLowerCase().includes(query) ||
      p.poNumber.toLowerCase().includes(query),
  );
});

const kpis = computed(() => [
  {
    label: 'Total Pipeline Value',
    value: requests.value.reduce(
      (s: number, p: ProcurementRequest) => s + (p.estimatedValue || 0),
      0,
    ),
    icon: 'account_balance_wallet',
    bgClass: 'bg-primary-container/10',
    iconClass: 'text-primary',
    subtext: 'Cumulative estimated value',
  },
  {
    label: 'Approved This Month',
    value: requests.value
      .filter((p: ProcurementRequest) => p.status === 'Received')
      .reduce((s: number, p: ProcurementRequest) => s + (p.estimatedValue || 0), 0),
    icon: 'verified',
    bgClass: 'bg-status-in-stock/10',
    iconClass: 'text-status-in-stock',
    subtext: 'Finalized purchase orders',
  },
  {
    label: 'Pending Approval',
    value: requests.value
      .filter((p: ProcurementRequest) => p.status === 'Pending Approval')
      .reduce((s: number, p: ProcurementRequest) => s + (p.estimatedValue || 0), 0),
    icon: 'pending_actions',
    bgClass: 'bg-metric-amber/20',
    iconClass: 'text-surface-tint',
    subtext: 'Requires management review',
  },
]);

const getStatusClass = (status: string) => {
  switch (status) {
    case 'Received':
      return 'bg-status-in-stock/20 text-status-in-stock';
    case 'Shipping':
      return 'bg-status-checked-out/20 text-status-checked-out';
    case 'Pending Approval':
      return 'bg-metric-amber/20 text-surface-tint';
    case 'Draft':
      return 'bg-surface-variant text-text-secondary';
    default:
      return 'bg-primary-container/20 text-primary';
  }
};

const getPriorityClass = (priority: string) => {
  switch (priority) {
    case 'High':
      return 'bg-status-critical/20 text-status-critical';
    case 'Medium':
      return 'bg-metric-amber/20 text-surface-tint border border-metric-amber/30';
    default:
      return 'bg-surface-variant text-text-secondary';
  }
};

onMounted(fetchRequests);
</script>
