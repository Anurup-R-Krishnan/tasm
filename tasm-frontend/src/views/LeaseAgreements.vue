<template>
  <main class="space-y-section-gap pb-24">
    <!-- Header -->
    <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
      <div>
        <h1 class="text-text-primary mb-1">Lease Agreements</h1>
        <p class="text-text-secondary font-body">
          Manage active asset leases, terms, and payment schedules.
        </p>
      </div>
      <div class="flex gap-3">
        <button
          @click="handleExport"
          class="bg-surface border border-border-default text-text-primary px-4 py-2 rounded-lg text-sm flex items-center gap-2 hover:bg-surface-subtle transition-colors shadow-sm"
        >
          <span class="material-symbols-outlined text-[18px]">download</span>
          Export
        </button>
        <button @click="showNewLeaseModal = true" class="btn-primary">
          <span class="material-symbols-outlined">gavel</span>
          New Agreement
        </button>
      </div>
    </div>

    <!-- KPI Grid -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
      <div class="premium-card">
        <div class="flex justify-between items-start mb-4">
          <div
            class="w-10 h-10 rounded-xl bg-primary-container/10 text-primary flex items-center justify-center"
          >
            <span class="material-symbols-outlined">payments</span>
          </div>
          <span
            class="text-[10px] font-bold text-text-secondary bg-surface-subtle px-2 py-0.5 rounded-full border border-border-default"
            >Active</span
          >
        </div>
        <p class="text-[10px] font-bold text-text-secondary uppercase tracking-widest">
          Total Monthly Liability
        </p>
        <h2 class="text-2xl font-bold text-text-primary mt-1">
          ₹{{ totalMonthlyValue.toLocaleString() }}
        </h2>
        <p class="text-xs text-text-secondary mt-4">Across {{ leases.length }} active agreements</p>
      </div>

      <div class="premium-card">
        <div class="flex justify-between items-start mb-4">
          <div
            class="w-10 h-10 rounded-xl bg-metric-amber/10 text-surface-tint flex items-center justify-center"
          >
            <span class="material-symbols-outlined">hourglass_empty</span>
          </div>
          <span
            class="text-[10px] font-bold text-status-critical bg-status-critical/10 px-2 py-0.5 rounded-full border border-status-critical/20"
            >Critical</span
          >
        </div>
        <p class="text-[10px] font-bold text-text-secondary uppercase tracking-widest">
          Expiring Within 90 Days
        </p>
        <h2 class="text-2xl font-bold text-text-primary mt-1">{{ expiringSoonCount }}</h2>
        <p class="text-xs text-text-secondary mt-4">Require immediate renewal review</p>
      </div>

      <div class="premium-card">
        <div class="flex justify-between items-start mb-4">
          <div
            class="w-10 h-10 rounded-xl bg-metric-sage/10 text-status-in-stock flex items-center justify-center"
          >
            <span class="material-symbols-outlined">verified</span>
          </div>
        </div>
        <p class="text-[10px] font-bold text-text-secondary uppercase tracking-widest">
          Compliance Rate
        </p>
        <h2 class="text-2xl font-bold text-text-primary mt-1">98.5%</h2>
        <p class="text-xs text-text-secondary mt-4">Documented and verified leases</p>
      </div>
    </div>

    <!-- Table -->
    <div class="premium-card !p-0 overflow-hidden">
      <div
        class="p-4 border-b border-border-default flex justify-between items-center bg-surface-subtle/30"
      >
        <div class="relative w-64">
          <span
            class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-text-secondary text-[18px]"
            >search</span
          >
          <input
            v-model="searchQuery"
            class="w-full bg-white border border-border-default rounded-xl py-2 pl-10 pr-4 text-sm focus:ring-2 focus:ring-primary/20 outline-none transition-all"
            placeholder="Search leases..."
          />
        </div>
      </div>
      <div class="overflow-x-auto p-4">
        <DataTable
          :value="filteredLeases"
          :loading="loading"
          paginator
          :rows="10"
          tableStyle="min-width: 50rem"
          class="w-full text-left"
          @row-click="(e: any) => router.push(`/lease/${e.data.id}`)"
          rowHover
        >
          <Column field="leaseId" header="AGREEMENT #" sortable>
            <template #body="slotProps">
              <span class="font-mono-data text-primary font-bold">{{
                slotProps.data.leaseId
              }}</span>
            </template>
          </Column>
          <Column field="vendor" header="LESSOR / VENDOR" sortable></Column>
          <Column field="monthlyCost" header="MONTHLY VALUE" sortable>
            <template #body="slotProps">
              <span class="font-mono-data"
                >₹{{ slotProps.data.monthlyCost?.toLocaleString() }}</span
              >
            </template>
          </Column>
          <Column field="startDate" header="COMMENCED" sortable>
            <template #body="slotProps">
              {{ new Date(slotProps.data.startDate).toLocaleDateString() }}
            </template>
          </Column>
          <Column field="endDate" header="EXPIRY" sortable>
            <template #body="slotProps">
              <span
                :class="{
                  'text-status-critical font-bold': isExpiringSoon(slotProps.data.endDate),
                }"
              >
                {{ new Date(slotProps.data.endDate).toLocaleDateString() }}
              </span>
            </template>
          </Column>
          <Column field="status" header="STATUS" sortable>
            <template #body="slotProps">
              <span
                class="px-2 py-0.5 rounded-full text-[10px] font-bold uppercase"
                :class="getStatusClass(slotProps.data.status)"
              >
                {{ slotProps.data.status }}
              </span>
            </template>
          </Column>
        </DataTable>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { getLeases } from '../api/financial';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import type { LeaseAgreement } from '../types/models';

const router = useRouter();
const leases = ref<LeaseAgreement[]>([]);
const loading = ref(true);
const searchQuery = ref('');
const showNewLeaseModal = ref(false);

const fetchLeases = async () => {
  try {
    leases.value = await getLeases();
  } catch (error) {
    console.error('Failed to fetch leases:', error);
  } finally {
    loading.value = false;
  }
};

const totalMonthlyValue = computed(() => {
  return leases.value
    .filter((l) => l.status === 'Active')
    .reduce((sum, l) => sum + (l.monthlyCost || 0), 0);
});

const expiringSoonCount = computed(() => {
  return leases.value.filter((l) => isExpiringSoon(l.endDate)).length;
});

const isExpiringSoon = (date: string) => {
  const expiry = new Date(date);
  const now = new Date();
  const diff = expiry.getTime() - now.getTime();
  const days = diff / (1000 * 60 * 60 * 24);
  return days > 0 && days < 90;
};

const filteredLeases = computed(() => {
  const q = searchQuery.value.toLowerCase();
  return leases.value.filter(
    (l) => l.leaseId.toLowerCase().includes(q) || l.vendor.toLowerCase().includes(q),
  );
});

const getStatusClass = (status: string) => {
  switch (status) {
    case 'Active':
      return 'bg-status-in-stock/10 text-status-in-stock';
    case 'Terminated':
      return 'bg-status-critical/10 text-status-critical';
    case 'Expired':
      return 'bg-surface-variant text-text-secondary';
    default:
      return 'bg-primary-container/10 text-primary';
  }
};

const handleExport = () => {
  // Simple CSV export
  const headers = ['Lease #', 'Lessor', 'Monthly Value', 'Start Date', 'End Date', 'Status'];
  const rows = filteredLeases.value.map((l) => [
    l.leaseId,
    l.vendor,
    l.monthlyCost,
    l.startDate,
    l.endDate,
    l.status,
  ]);
  const csv = [headers, ...rows].map((r) => r.join(',')).join('\n');
  const blob = new Blob([csv], { type: 'text/csv' });
  const url = URL.createObjectURL(blob);
  const link = document.createElement('a');
  link.href = url;
  link.download = 'lease_agreements.csv';
  link.click();
};

onMounted(fetchLeases);
</script>
