<template>
  <main class="p-page-margin max-w-[1400px] mx-auto space-y-section-gap pb-24">
    <!-- Header -->
    <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
      <div>
        <h1 class="text-text-primary mb-1">Operational Overview</h1>
        <p class="text-text-secondary">
          Real-time insights across campus assets and infrastructure.
        </p>
      </div>
      <div class="flex gap-3 bg-surface p-1 rounded-xl border border-border-default shadow-sm">
        <button
          class="px-4 py-1.5 rounded-lg text-xs font-bold transition-all bg-indigo-600 text-white shadow-lg shadow-indigo-200"
        >
          24 Hours
        </button>
        <button
          class="px-4 py-1.5 rounded-lg text-xs font-bold text-slate-500 hover:text-indigo-600 transition-all"
        >
          7 Days
        </button>
        <button
          class="px-4 py-1.5 rounded-lg text-xs font-bold text-slate-500 hover:text-indigo-600 transition-all"
        >
          30 Days
        </button>
      </div>
    </div>

    <!-- Main KPIs -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <div
        v-for="kpi in mainKPIs"
        :key="kpi.label"
        class="premium-card group hover:border-indigo-200 transition-all"
      >
        <div class="flex justify-between items-start mb-4">
          <div
            class="w-12 h-12 rounded-2xl flex items-center justify-center shadow-sm transition-transform group-hover:scale-110"
            :class="kpi.bgClass"
          >
            <span class="material-symbols-outlined" :class="kpi.iconClass">{{ kpi.icon }}</span>
          </div>
          <div
            class="flex items-center gap-1 text-[10px] font-bold px-2 py-0.5 rounded-full border"
            :class="kpi.trendClass"
          >
            <span class="material-symbols-outlined text-xs">{{ kpi.trendIcon }}</span>
            {{ kpi.trend }}
          </div>
        </div>
        <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">
          {{ kpi.label }}
        </p>
        <h2 class="text-3xl font-bold text-slate-900 mt-1">{{ kpi.value }}</h2>
        <div class="mt-4 flex items-center gap-2">
          <div class="flex-1 h-1 bg-slate-100 rounded-full overflow-hidden">
            <div
              class="h-full rounded-full transition-all duration-1000"
              :class="kpi.barClass"
              :style="{ width: kpi.progress + '%' }"
            ></div>
          </div>
          <span class="text-[10px] font-bold text-slate-400">{{ kpi.progress }}%</span>
        </div>
      </div>
    </div>

    <!-- Middle Section: Activity & Charts -->
    <div class="grid grid-cols-1 lg:grid-cols-12 gap-8">
      <!-- High Priority Work Orders -->
      <div class="lg:col-span-8 flex flex-col">
        <div class="premium-card !p-0 overflow-hidden flex-1">
          <div
            class="p-6 border-b border-slate-50 flex justify-between items-center bg-slate-50/30"
          >
            <div>
              <h3 class="text-sm font-bold text-slate-900">Critical Work Orders</h3>
              <p class="text-[10px] text-slate-400 font-medium uppercase tracking-widest mt-1">
                Requiring immediate attention
              </p>
            </div>
            <button class="text-xs font-bold text-indigo-600 hover:text-indigo-700">
              View All
            </button>
          </div>
          <div class="p-6 space-y-4">
            <div
              v-for="order in workOrders"
              :key="order.id"
              class="flex items-center gap-6 p-4 rounded-2xl bg-white border border-slate-50 hover:border-indigo-100 hover:shadow-md transition-all group"
            >
              <div
                class="w-12 h-12 rounded-xl flex items-center justify-center text-rose-500 bg-rose-50 font-bold text-xs uppercase"
              >
                {{ order.severity.charAt(0) }}
              </div>
              <div class="flex-1 min-w-0">
                <h4
                  class="text-sm font-bold text-slate-900 group-hover:text-indigo-600 transition-colors"
                >
                  {{ order.title }}
                </h4>
                <div class="flex items-center gap-4 mt-1.5">
                  <span class="text-[10px] font-medium text-slate-400 flex items-center gap-1">
                    <span class="material-symbols-outlined text-xs">location_on</span>
                    {{ order.assetLocation }}
                  </span>
                  <span class="text-[10px] font-medium text-slate-400 flex items-center gap-1">
                    <span class="material-symbols-outlined text-xs">person</span>
                    {{ order.technician }}
                  </span>
                </div>
              </div>
              <div class="text-right">
                <span
                  class="text-[10px] font-bold uppercase tracking-wider px-2 py-1 rounded-lg bg-amber-50 text-amber-600"
                  >{{ order.status }}</span
                >
                <p class="text-[10px] text-slate-400 mt-2 font-medium">
                  {{ formatDate(order.targetDate) }}
                </p>
              </div>
            </div>
            <div
              v-if="workOrders.length === 0"
              class="py-12 text-center text-slate-400 text-sm italic"
            >
              No critical work orders found.
            </div>
          </div>
        </div>
      </div>

      <!-- Quick Actions & Alerts -->
      <div class="lg:col-span-4 space-y-8">
        <div class="premium-card">
          <h3 class="text-xs font-bold text-slate-400 uppercase tracking-widest mb-6">
            Quick Operations
          </h3>
          <div class="grid grid-cols-2 gap-4">
            <button
              class="flex flex-col items-center gap-3 p-4 rounded-2xl bg-indigo-50/50 border border-indigo-100 hover:bg-indigo-600 hover:text-white transition-all group shadow-sm"
            >
              <span class="material-symbols-outlined text-indigo-600 group-hover:text-white"
                >add_shopping_cart</span
              >
              <span class="text-[10px] font-bold uppercase tracking-wider">Procurement</span>
            </button>
            <button
              class="flex flex-col items-center gap-3 p-4 rounded-2xl bg-emerald-50/50 border border-emerald-100 hover:bg-emerald-600 hover:text-white transition-all group shadow-sm"
            >
              <span class="material-symbols-outlined text-emerald-600 group-hover:text-white"
                >handyman</span
              >
              <span class="text-[10px] font-bold uppercase tracking-wider">Service Log</span>
            </button>
            <button
              class="flex flex-col items-center gap-3 p-4 rounded-2xl bg-rose-50/50 border border-rose-100 hover:bg-rose-600 hover:text-white transition-all group shadow-sm"
            >
              <span class="material-symbols-outlined text-rose-600 group-hover:text-white"
                >qr_code_scanner</span
              >
              <span class="text-[10px] font-bold uppercase tracking-wider">Audit Scan</span>
            </button>
            <button
              class="flex flex-col items-center gap-3 p-4 rounded-2xl bg-amber-50/50 border border-amber-100 hover:bg-amber-600 hover:text-white transition-all group shadow-sm"
            >
              <span class="material-symbols-outlined text-amber-600 group-hover:text-white"
                >sync_alt</span
              >
              <span class="text-[10px] font-bold uppercase tracking-wider">Transfer</span>
            </button>
          </div>
        </div>

        <!-- Recent Alerts -->
        <div class="premium-card">
          <div class="flex justify-between items-center mb-6">
            <h3 class="text-xs font-bold text-slate-400 uppercase tracking-widest">
              Recent Alerts
            </h3>
            <span class="w-2 h-2 rounded-full bg-rose-500 animate-ping"></span>
          </div>
          <div class="space-y-6">
            <div v-for="alert in alerts" :key="alert.id" class="flex gap-4 group cursor-pointer">
              <div
                class="w-1 h-8 rounded-full transition-all group-hover:h-12"
                :class="alert.type === 'Critical' ? 'bg-rose-500' : 'bg-amber-500'"
              ></div>
              <div class="flex-1 min-w-0">
                <h4
                  class="text-xs font-bold text-slate-900 group-hover:text-indigo-600 transition-colors"
                >
                  {{ alert.title }}
                </h4>
                <p class="text-[10px] text-slate-400 mt-1 line-clamp-1">{{ alert.message }}</p>
              </div>
            </div>
            <div v-if="alerts.length === 0" class="py-4 text-center text-slate-400 text-xs italic">
              No active system alerts.
            </div>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';

const assets = ref<any[]>([]);
const workOrders = ref<any[]>([]);
const alerts = ref<any[]>([]);
const loading = ref(true);

const fetchData = async () => {
  loading.value = true;
  try {
    const [assetRes, orderRes, alertRes] = await Promise.all([
      fetch('http://localhost:8080/api/assets'),
      fetch('http://localhost:8080/api/work-orders'),
      fetch('http://localhost:8080/api/alerts'),
    ]);
    if (assetRes.ok) assets.value = await assetRes.json();
    if (orderRes.ok)
      workOrders.value = (await orderRes.json())
        .filter((o: any) => o.severity === 'Critical' || o.severity === 'High')
        .slice(0, 4);
    if (alertRes.ok) alerts.value = (await alertRes.json()).slice(0, 5);
  } catch (err) {
    console.error('Failed to fetch dashboard data:', err);
  } finally {
    loading.value = false;
  }
};

const mainKPIs = computed(() => [
  {
    label: 'Total Campus Assets',
    value: assets.value.length,
    icon: 'inventory_2',
    bgClass: 'bg-indigo-50',
    iconClass: 'text-indigo-500',
    trend: '4.2%',
    trendIcon: 'arrow_upward',
    trendClass: 'text-emerald-600 bg-emerald-50 border-emerald-100',
    progress: 78,
    barClass: 'bg-indigo-500',
  },
  {
    label: 'Operational Health',
    value:
      Math.round(
        (assets.value.filter((a) => a.status === 'In Stock').length / (assets.value.length || 1)) *
          100,
      ) + '%',
    icon: 'verified_user',
    bgClass: 'bg-emerald-50',
    iconClass: 'text-emerald-500',
    trend: '99.9%',
    trendIcon: 'check_circle',
    trendClass: 'text-indigo-600 bg-indigo-50 border-indigo-100',
    progress: 92,
    barClass: 'bg-emerald-500',
  },
  {
    label: 'Maintenance Queue',
    value: workOrders.value.length,
    icon: 'handyman',
    bgClass: 'bg-rose-50',
    iconClass: 'text-rose-500',
    trend: '12%',
    trendIcon: 'arrow_downward',
    trendClass: 'text-rose-600 bg-rose-50 border-rose-100',
    progress: 45,
    barClass: 'bg-rose-500',
  },
  {
    label: 'Est. Valuation',
    value: '₹' + (assets.value.reduce((s, a) => s + (a.value || 0), 0) / 1000000).toFixed(1) + 'M',
    icon: 'account_balance_wallet',
    bgClass: 'bg-slate-100',
    iconClass: 'text-slate-600',
    trend: '2.5%',
    trendIcon: 'arrow_upward',
    trendClass: 'text-slate-600 bg-slate-50 border-slate-100',
    progress: 60,
    barClass: 'bg-slate-400',
  },
]);

const formatDate = (date: string) => {
  if (!date) return 'N/A';
  return new Date(date).toLocaleDateString('en-GB', { day: '2-digit', month: 'short' });
};

onMounted(fetchData);
</script>
