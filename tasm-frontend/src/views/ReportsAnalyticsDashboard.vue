<template>
  <main class="space-y-section-gap pb-24 font-body">
    <!-- Header -->
    <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
      <div>
        <h1 class="text-text-primary">Financial Reports & Analytics</h1>
        <p class="text-text-secondary mt-1">
          Detailed overview of asset lifecycle costs, procurement pipelines, and general ledger.
        </p>
      </div>
      <div class="flex gap-3">
        <button
          class="bg-surface border border-border-default text-text-primary px-4 py-2 rounded-lg text-sm flex items-center gap-2 hover:bg-surface-subtle transition-colors shadow-sm opacity-50 cursor-not-allowed"
        >
          <span class="material-symbols-outlined text-[18px]">calendar_month</span>
          Current Period
        </button>
        <button @click="handleExport" class="btn-primary">
          <span class="material-symbols-outlined">download</span>
          Export Report
        </button>
      </div>
    </div>

    <!-- Financial KPI Row -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div v-for="kpi in financialKPIs" :key="kpi.label" class="premium-card">
        <div class="flex justify-between items-start mb-4">
          <div
            class="w-10 h-10 rounded-xl flex items-center justify-center shadow-sm"
            :class="kpi.bgClass"
          >
            <span class="material-symbols-outlined" :class="kpi.iconClass">{{ kpi.icon }}</span>
          </div>
          <span
            class="text-[10px] font-bold text-text-secondary bg-surface-subtle px-2 py-0.5 rounded-full border border-border-default"
            >Live</span
          >
        </div>
        <p class="text-[10px] font-bold text-text-secondary uppercase tracking-widest">
          {{ kpi.label }}
        </p>
        <h2 class="text-2xl font-bold text-text-primary mt-1">₹{{ kpi.value.toLocaleString() }}</h2>
        <div class="mt-4 h-1 w-full bg-surface-variant rounded-full overflow-hidden">
          <div class="h-full bg-primary/20 rounded-full" style="width: 100%"></div>
        </div>
      </div>
    </div>

    <!-- Charts & Pipeline -->
    <div class="grid grid-cols-1 lg:grid-cols-12 gap-8">
      <!-- Procurement Pipeline -->
      <div class="lg:col-span-8">
        <div class="premium-card !p-0 overflow-hidden h-full flex flex-col">
          <div
            class="p-6 border-b border-border-default flex justify-between items-center bg-surface-subtle/30"
          >
            <div>
              <h3 class="text-sm font-bold text-text-primary">Procurement Pipeline</h3>
              <p class="text-[10px] text-text-secondary font-medium uppercase tracking-widest mt-1">
                Pending Approvals & Shipping
              </p>
            </div>
          </div>
          <div class="overflow-x-auto flex-1">
            <table class="w-full text-left">
              <thead
                class="bg-surface-subtle/50 text-[10px] font-bold text-text-secondary uppercase tracking-widest border-b border-border-default"
              >
                <tr>
                  <th class="px-6 py-4">Request Detail</th>
                  <th class="px-6 py-4">Priority</th>
                  <th class="px-6 py-4">Status</th>
                  <th class="px-6 py-4">Estimated Value</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-slate-50">
                <tr
                  v-for="req in procurements"
                  :key="req.id"
                  class="hover:bg-surface-subtle/50 transition-colors group"
                >
                  <td class="px-6 py-4">
                    <div class="flex items-center gap-3">
                      <div
                        class="w-8 h-8 rounded-lg bg-primary-container/20 text-primary flex items-center justify-center font-bold text-[10px]"
                      >
                        {{ req.requestorInitials || '??' }}
                      </div>
                      <div>
                        <p
                          class="text-sm font-bold text-text-primary leading-none group-hover:text-primary transition-colors"
                        >
                          {{ req.title }}
                        </p>
                        <p class="text-[10px] text-text-secondary mt-1.5 font-medium">
                          {{ req.department }} • PO: {{ req.poNumber || 'N/A' }}
                        </p>
                      </div>
                    </div>
                  </td>
                  <td class="px-6 py-4">
                    <span
                      class="text-[9px] font-bold uppercase px-2 py-0.5 rounded-full"
                      :class="getPriorityClass(req.priority)"
                      >{{ req.priority }}</span
                    >
                  </td>
                  <td class="px-6 py-4">
                    <span
                      class="text-[10px] font-bold text-text-secondary flex items-center gap-1.5"
                    >
                      <span
                        class="w-1.5 h-1.5 rounded-full"
                        :class="getStatusDotClass(req.status)"
                      ></span>
                      {{ req.status }}
                    </span>
                  </td>
                  <td class="px-6 py-4 text-sm font-bold text-text-primary">
                    ₹{{ (req.estimatedValue || 0).toLocaleString() }}
                  </td>
                </tr>
                <tr v-if="procurements.length === 0">
                  <td colspan="4" class="px-6 py-12 text-center text-text-secondary text-sm italic">
                    No active procurement requests.
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>

      <!-- General Ledger Feed -->
      <div class="lg:col-span-4">
        <div class="premium-card !p-0 overflow-hidden h-full flex flex-col">
          <div
            class="p-6 border-b border-border-default flex justify-between items-center bg-surface-subtle/30"
          >
            <h3 class="text-sm font-bold text-text-primary">Recent Transactions</h3>
            <span
              class="text-[10px] font-bold bg-surface-variant text-text-secondary px-2 py-0.5 rounded-full"
              >Ledger Feed</span
            >
          </div>
          <div class="p-4 space-y-3 overflow-y-auto max-h-[500px] flex-1">
            <div
              v-for="entry in ledgers"
              :key="entry.id"
              class="p-4 bg-surface-subtle/50 border border-border-default rounded-xl hover:border-slate-200 transition-all group"
            >
              <div class="flex justify-between items-start mb-2">
                <span
                  class="text-[9px] font-mono font-bold text-text-disabled uppercase tracking-widest"
                  >{{ entry.transactionId }}</span
                >
                <span
                  class="text-[10px] font-bold"
                  :class="entry.type === 'Credit' ? 'text-emerald-500' : 'text-rose-500'"
                >
                  {{ entry.type === 'Credit' ? '+' : '-' }} ₹{{
                    (entry.amount || 0).toLocaleString()
                  }}
                </span>
              </div>
              <h4 class="text-xs font-bold text-text-primary truncate">{{ entry.description }}</h4>
              <div class="flex justify-between items-center mt-3">
                <span class="text-[10px] font-medium text-text-secondary">{{
                  formatDate(entry.date)
                }}</span>
                <span
                  class="text-[9px] font-bold uppercase tracking-wider text-text-secondary bg-white px-2 py-0.5 rounded border border-border-default"
                  >{{ entry.category }}</span
                >
              </div>
            </div>
            <div
              v-if="ledgers.length === 0"
              class="py-12 text-center text-text-secondary text-sm italic"
            >
              No transactions recorded in ledger.
            </div>
          </div>
          <div class="p-4 bg-surface-subtle/50 border-t border-border-default">
            <button
              @click="router.push('/ledgers')"
              class="w-full py-2 bg-white border border-border-default rounded-lg text-xs font-bold text-text-secondary hover:text-primary transition-colors shadow-sm"
            >
              View Full Audit Log
            </button>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { getLedgers } from '../api/financial';
import { getProcurements } from '../api/procurements';
import { getAssets } from '../api/assets';
import type { LedgerEntry, ProcurementRequest, Asset } from '../types/models';

const router = useRouter();
const ledgers = ref<LedgerEntry[]>([]);
const procurements = ref<ProcurementRequest[]>([]);
const assets = ref<Asset[]>([]);
const loading = ref(true);

const fetchData = async () => {
  loading.value = true;
  try {
    const [ledgerData, procData, assetData] = await Promise.all([
      getLedgers(),
      getProcurements(),
      getAssets(),
    ]);
    ledgers.value = ledgerData;
    procurements.value = procData;
    assets.value = assetData;
  } catch (err) {
    console.error('Failed to fetch financial data:', err);
  } finally {
    loading.value = false;
  }
};

const financialKPIs = computed(() => {
  const totalSpend = ledgers.value
    .filter((l) => l.type === 'Debit')
    .reduce((sum, l) => sum + (l.amount || 0), 0);

  const pendingValue = procurements.value
    .filter((p) => p.status !== 'Received')
    .reduce((sum, p) => sum + (p.estimatedValue || 0), 0);

  const assetValue = assets.value.reduce((sum, a) => sum + (a.value || 0), 0);

  return [
    {
      label: 'Total Expenditure',
      value: totalSpend,
      icon: 'payments',
      bgClass: 'bg-rose-50',
      iconClass: 'text-rose-500',
    },
    {
      label: 'Pending Pipeline',
      value: pendingValue,
      icon: 'shopping_cart_checkout',
      bgClass: 'bg-primary-container/10',
      iconClass: 'text-primary',
    },
    {
      label: 'Asset Valuation',
      value: assetValue,
      icon: 'account_balance_wallet',
      bgClass: 'bg-surface-subtle',
      iconClass: 'text-text-secondary',
    },
    {
      label: 'Ledger Entries',
      value: ledgers.value.length,
      icon: 'receipt_long',
      bgClass: 'bg-amber-50',
      iconClass: 'text-amber-500',
    },
  ];
});

const formatDate = (date: string) => {
  if (!date) return 'N/A';
  return new Date(date).toLocaleDateString('en-GB', { day: '2-digit', month: 'short' });
};

const getPriorityClass = (priority: string) => {
  switch (priority) {
    case 'High':
      return 'bg-rose-100 text-rose-700';
    case 'Medium':
      return 'bg-amber-100 text-amber-700';
    default:
      return 'bg-slate-100 text-slate-700';
  }
};

const getStatusDotClass = (status: string) => {
  if (status.includes('Received')) return 'bg-emerald-500';
  if (status.includes('Pending') || status.includes('PO')) return 'bg-amber-500';
  if (status.includes('Shipping')) return 'bg-blue-500';
  return 'bg-slate-400';
};

const handleExport = () => {
  window.print();
};

onMounted(fetchData);
</script>
