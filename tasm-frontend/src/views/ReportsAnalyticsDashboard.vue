<template>
  <main class="p-page-margin max-w-[1400px] mx-auto space-y-section-gap pb-24">
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
          class="bg-surface border border-border-default text-text-primary px-4 py-2 rounded-lg text-sm flex items-center gap-2 hover:bg-surface-subtle transition-colors shadow-sm"
        >
          <span class="material-symbols-outlined text-[18px]">calendar_month</span>
          FY 2025-26
        </button>
        <button class="btn-primary">
          <span class="material-symbols-outlined">download</span>
          Export PDF
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
            class="text-[10px] font-bold text-slate-400 bg-slate-50 px-2 py-0.5 rounded-full border border-slate-100"
            >Live</span
          >
        </div>
        <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">
          {{ kpi.label }}
        </p>
        <h2 class="text-2xl font-bold text-slate-900 mt-1">₹{{ kpi.value.toLocaleString() }}</h2>
        <div class="mt-4 h-1 w-full bg-slate-100 rounded-full overflow-hidden">
          <div
            class="h-full bg-indigo-500 rounded-full"
            :style="{ width: kpi.progress + '%' }"
          ></div>
        </div>
      </div>
    </div>

    <!-- Charts & Pipeline -->
    <div class="grid grid-cols-1 lg:grid-cols-12 gap-8">
      <!-- Procurement Pipeline -->
      <div class="lg:col-span-8">
        <div class="premium-card !p-0 overflow-hidden h-full flex flex-col">
          <div
            class="p-6 border-b border-slate-50 flex justify-between items-center bg-slate-50/30"
          >
            <div>
              <h3 class="text-sm font-bold text-slate-900">Procurement Pipeline</h3>
              <p class="text-[10px] text-slate-400 font-medium uppercase tracking-widest mt-1">
                Pending Approvals & Shipping
              </p>
            </div>
            <div class="flex gap-2">
              <button
                class="text-[10px] font-bold px-3 py-1.5 rounded-lg border border-slate-200 text-slate-600 hover:bg-white shadow-sm transition-all"
              >
                By Dept
              </button>
              <button
                class="text-[10px] font-bold px-3 py-1.5 rounded-lg bg-white border border-slate-200 text-slate-900 shadow-sm transition-all"
              >
                By Priority
              </button>
            </div>
          </div>
          <div class="overflow-x-auto flex-1">
            <table class="w-full text-left">
              <thead
                class="bg-slate-50/50 text-[10px] font-bold text-slate-400 uppercase tracking-widest border-b border-slate-50"
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
                  class="hover:bg-slate-50/50 transition-colors group"
                >
                  <td class="px-6 py-4">
                    <div class="flex items-center gap-3">
                      <div
                        class="w-8 h-8 rounded-lg bg-indigo-50 text-indigo-500 flex items-center justify-center font-bold text-[10px]"
                      >
                        {{ req.requestorInitials }}
                      </div>
                      <div>
                        <p
                          class="text-sm font-bold text-slate-900 leading-none group-hover:text-indigo-600 transition-colors"
                        >
                          {{ req.title }}
                        </p>
                        <p class="text-[10px] text-slate-400 mt-1.5 font-medium">
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
                    <span class="text-[10px] font-bold text-slate-600 flex items-center gap-1.5">
                      <span
                        class="w-1.5 h-1.5 rounded-full bg-slate-400"
                        :class="getStatusDotClass(req.status)"
                      ></span>
                      {{ req.status }}
                    </span>
                  </td>
                  <td class="px-6 py-4 text-sm font-bold text-slate-900">
                    ₹{{ (req.estimatedValue || 0).toLocaleString() }}
                  </td>
                </tr>
                <tr v-if="procurements.length === 0">
                  <td colspan="4" class="px-6 py-12 text-center text-slate-400 text-sm italic">
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
            class="p-6 border-b border-slate-50 flex justify-between items-center bg-slate-50/30"
          >
            <h3 class="text-sm font-bold text-slate-900">Recent Transactions</h3>
            <span class="text-[10px] font-bold bg-slate-100 text-slate-500 px-2 py-0.5 rounded-full"
              >Ledger Feed</span
            >
          </div>
          <div class="p-4 space-y-3 overflow-y-auto max-h-[500px] flex-1">
            <div
              v-for="entry in ledgers"
              :key="entry.id"
              class="p-4 bg-slate-50/50 border border-slate-100 rounded-xl hover:border-slate-200 transition-all group"
            >
              <div class="flex justify-between items-start mb-2">
                <span
                  class="text-[9px] font-mono font-bold text-slate-300 uppercase tracking-widest"
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
              <h4 class="text-xs font-bold text-slate-900 truncate">{{ entry.description }}</h4>
              <div class="flex justify-between items-center mt-3">
                <span class="text-[10px] font-medium text-slate-400">{{
                  formatDate(entry.date)
                }}</span>
                <span
                  class="text-[9px] font-bold uppercase tracking-wider text-slate-500 bg-white px-2 py-0.5 rounded border border-slate-100"
                  >{{ entry.category }}</span
                >
              </div>
            </div>
            <div
              v-if="ledgers.length === 0"
              class="py-12 text-center text-slate-400 text-sm italic"
            >
              No transactions recorded in ledger.
            </div>
          </div>
          <div class="p-4 bg-slate-50/50 border-t border-slate-50">
            <button
              class="w-full py-2 bg-white border border-slate-200 rounded-lg text-xs font-bold text-slate-600 hover:text-indigo-600 transition-colors shadow-sm"
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

const ledgers = ref<any[]>([]);
const procurements = ref<any[]>([]);
const loading = ref(true);

const fetchData = async () => {
  loading.value = true;
  try {
    const [ledgerRes, procRes] = await Promise.all([
      fetch('http://localhost:8080/api/ledgers'),
      fetch('http://localhost:8080/api/procurements'),
    ]);
    if (ledgerRes.ok) ledgers.value = await ledgerRes.json();
    if (procRes.ok) procurements.value = await procRes.json();
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
  const totalBudget = 5000000; // Mock budget for progress
  const pendingValue = procurements.value
    .filter((p) => p.status !== 'Received')
    .reduce((sum, p) => sum + (p.estimatedValue || 0), 0);

  return [
    {
      label: 'Total Expenditure',
      value: totalSpend,
      icon: 'payments',
      bgClass: 'bg-rose-50',
      iconClass: 'text-rose-500',
      progress: (totalSpend / totalBudget) * 100,
    },
    {
      label: 'Pending Procurement',
      value: pendingValue,
      icon: 'shopping_cart_checkout',
      bgClass: 'bg-indigo-50',
      iconClass: 'text-indigo-500',
      progress: 45,
    },
    {
      label: 'Opex Savings',
      value: 842000,
      icon: 'savings',
      bgClass: 'bg-emerald-50',
      iconClass: 'text-emerald-500',
      progress: 72,
    },
    {
      label: 'Asset Valuation',
      value: 12450000,
      icon: 'account_balance_wallet',
      bgClass: 'bg-slate-50',
      iconClass: 'text-slate-500',
      progress: 60,
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

onMounted(fetchData);
</script>
