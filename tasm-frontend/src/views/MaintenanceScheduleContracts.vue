<template>
  <main class="space-y-section-gap pb-24">
    <!-- Header -->
    <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
      <div>
        <h1 class="text-text-primary">Maintenance & Contracts</h1>
        <p class="text-text-secondary mt-1">
          Monitor service agreements and active maintenance operations.
        </p>
      </div>
      <div class="flex gap-3">
        <button
          class="bg-surface border border-border-default text-text-primary px-4 py-2 rounded-lg text-sm flex items-center gap-2 hover:bg-surface-subtle transition-colors shadow-sm"
        >
          <span class="material-symbols-outlined text-[18px]">history</span>
          History
        </button>
        <button class="btn-primary">
          <span class="material-symbols-outlined">add_task</span>
          Schedule Service
        </button>
      </div>
    </div>

    <!-- Bento Grid Summary -->
    <div class="grid grid-cols-1 md:grid-cols-12 gap-6">
      <div class="md:col-span-4 premium-card flex flex-col justify-center">
        <p class="text-[10px] font-bold text-text-secondary uppercase tracking-widest">
          Active Contracts
        </p>
        <h2 class="text-4xl font-bold text-text-primary mt-2">{{ contracts.length }}</h2>
        <div
          class="flex items-center gap-2 mt-4 text-status-in-stock text-xs font-bold bg-emerald-50 w-fit px-3 py-1 rounded-full"
        >
          <span class="material-symbols-outlined text-sm">verified</span>
          All Compliance Verified
        </div>
      </div>
      <div class="md:col-span-4 premium-card flex flex-col justify-center">
        <p class="text-[10px] font-bold text-text-secondary uppercase tracking-widest">
          Pending Work Orders
        </p>
        <h2 class="text-4xl font-bold text-text-primary mt-2">{{ openWorkOrders.length }}</h2>
        <div
          class="flex items-center gap-2 mt-4 text-amber-600 text-xs font-bold bg-amber-50 w-fit px-3 py-1 rounded-full"
        >
          <span class="material-symbols-outlined text-sm">pending_actions</span>
          {{ highSeverityOrders.length }} High Priority
        </div>
      </div>
      <div
        class="md:col-span-4 premium-card flex flex-col justify-center bg-gradient-to-br from-primary to-surface-secondary text-white border-none shadow-lg shadow-primary/10"
      >
        <p class="text-[10px] font-bold text-primary-container uppercase tracking-widest">
          Monthly Maint. Cost
        </p>
        <h2 class="text-4xl font-bold mt-2">₹{{ totalMaintenanceCost.toLocaleString() }}</h2>
        <p class="text-primary-container text-[10px] mt-4 font-medium uppercase tracking-wider">
          Across 12 Vendors
        </p>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-8">
      <!-- Active Contracts Table -->
      <div class="lg:col-span-7">
        <div class="premium-card !p-0 overflow-hidden h-full flex flex-col">
          <div
            class="p-6 border-b border-border-default flex justify-between items-center bg-surface-subtle/30"
          >
            <div>
              <h3 class="text-sm font-bold text-text-primary">Service Agreements</h3>
              <p class="text-[10px] text-text-secondary font-medium uppercase tracking-widest mt-1">
                Vendor Contracts
              </p>
            </div>
            <button
              class="p-2 hover:bg-white rounded-lg text-text-secondary hover:text-primary transition-all shadow-sm"
            >
              <span class="material-symbols-outlined">filter_list</span>
            </button>
          </div>
          <div class="overflow-x-auto flex-1">
            <table class="w-full text-left">
              <thead
                class="bg-surface-subtle/50 text-[10px] font-bold text-text-secondary uppercase tracking-widest border-b border-border-default"
              >
                <tr>
                  <th class="px-6 py-4">Vendor</th>
                  <th class="px-6 py-4">Service Type</th>
                  <th class="px-6 py-4">Period</th>
                  <th class="px-6 py-4">Status</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-slate-50">
                <tr
                  v-for="contract in contracts"
                  :key="contract.id"
                  class="hover:bg-surface-subtle/50 transition-colors"
                >
                  <td class="px-6 py-4 font-bold text-text-primary text-sm">
                    {{ contract.vendorName }}
                  </td>
                  <td class="px-6 py-4">
                    <span class="text-xs text-slate-600 bg-slate-100 px-2 py-0.5 rounded-lg">{{
                      contract.serviceType
                    }}</span>
                  </td>
                  <td class="px-6 py-4">
                    <p class="text-[11px] font-bold text-text-secondary">
                      {{ formatDate(contract.startDate) }} - {{ formatDate(contract.endDate) }}
                    </p>
                  </td>
                  <td class="px-6 py-4">
                    <span
                      class="text-[10px] font-bold uppercase tracking-wider px-2 py-0.5 rounded-full"
                      :class="getContractStatusClass(contract.status)"
                    >
                      {{ contract.status }}
                    </span>
                  </td>
                </tr>
                <tr v-if="contracts.length === 0">
                  <td colspan="4" class="px-6 py-12 text-center text-text-secondary text-sm italic">
                    No active contracts found.
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>

      <!-- Active Work Orders / Schedule -->
      <div class="lg:col-span-5">
        <div class="premium-card !p-0 overflow-hidden h-full flex flex-col">
          <div
            class="p-6 border-b border-border-default flex justify-between items-center bg-surface-subtle/30"
          >
            <div>
              <h3 class="text-sm font-bold text-text-primary">Active Work Orders</h3>
              <p class="text-[10px] text-text-secondary font-medium uppercase tracking-widest mt-1">
                Operations Queue
              </p>
            </div>
            <span
              class="text-[10px] font-bold bg-primary-container/20 text-primary px-2 py-0.5 rounded-full"
              >{{ openWorkOrders.length }} Open</span
            >
          </div>
          <div class="p-4 space-y-3 overflow-y-auto max-h-[500px]">
            <div
              v-for="order in openWorkOrders"
              :key="order.id"
              class="p-4 bg-surface-subtle/50 border border-border-default rounded-xl hover:border-primary/20 transition-colors cursor-pointer group"
            >
              <div class="flex justify-between items-start mb-2">
                <span class="text-[10px] font-mono text-text-secondary font-bold tracking-widest">{{
                  order.workOrderId
                }}</span>
                <span
                  class="text-[9px] font-bold uppercase px-2 py-0.5 rounded-full"
                  :class="getSeverityClass(order.severity)"
                  >{{ order.severity }}</span
                >
              </div>
              <h4
                class="text-sm font-bold text-text-primary group-hover:text-primary transition-colors"
              >
                {{ order.title }}
              </h4>
              <p class="text-[11px] text-text-secondary mt-1 line-clamp-1">{{ order.issue }}</p>
              <div class="flex justify-between items-center mt-4">
                <div class="flex items-center gap-1.5">
                  <div
                    class="w-6 h-6 rounded-full bg-primary-container/20 border-2 border-white flex items-center justify-center text-[10px] font-bold text-primary uppercase"
                  >
                    {{ order.technician?.charAt(0) }}
                  </div>
                  <span class="text-[10px] font-bold text-text-secondary">{{
                    order.technician
                  }}</span>
                </div>
                <span class="text-[10px] font-bold text-text-secondary">{{
                  formatDate(order.targetDate)
                }}</span>
              </div>
            </div>
            <div
              v-if="openWorkOrders.length === 0"
              class="py-12 text-center text-text-secondary text-sm italic"
            >
              No pending work orders.
            </div>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { getContracts } from '../api/financial';
import { getWorkOrders } from '../api/workOrders';

const contracts = ref<any[]>([]);
const workOrders = ref<any[]>([]);
const loading = ref(true);

const fetchData = async () => {
  loading.value = true;
  try {
    const [contractData, orderData] = await Promise.all([getContracts(), getWorkOrders()]);
    contracts.value = contractData as any[];
    workOrders.value = orderData as any[];
  } catch (err) {
    console.error('Failed to fetch maintenance data:', err);
  } finally {
    loading.value = false;
  }
};

const openWorkOrders = computed(() => {
  return workOrders.value.filter((o) => o.status !== 'Closed');
});

const highSeverityOrders = computed(() => {
  return openWorkOrders.value.filter((o) => o.severity === 'Critical' || o.severity === 'High');
});

const totalMaintenanceCost = computed(() => {
  return workOrders.value.reduce((sum, o) => sum + (o.cost || 0), 0);
});

const formatDate = (date: string) => {
  if (!date) return 'N/A';
  return new Date(date).toLocaleDateString('en-GB', { day: '2-digit', month: 'short' });
};

const getContractStatusClass = (status: string) => {
  return status === 'Active' ? 'bg-emerald-100 text-emerald-700' : 'bg-slate-100 text-slate-700';
};

const getSeverityClass = (sev: string) => {
  switch (sev) {
    case 'Critical':
      return 'bg-rose-100 text-rose-700';
    case 'High':
      return 'bg-amber-100 text-amber-700';
    case 'Medium':
      return 'bg-blue-100 text-blue-700';
    default:
      return 'bg-slate-100 text-slate-700';
  }
};

onMounted(fetchData);
</script>
