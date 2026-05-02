<template>
  <main class="space-y-section-gap pb-24">
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
          class="px-4 py-1.5 rounded-lg text-xs font-bold transition-all bg-primary text-on-primary shadow-lg shadow-primary/10"
        >
          24 Hours
        </button>
        <button
          class="px-4 py-1.5 rounded-lg text-xs font-bold text-text-secondary hover:text-primary transition-all"
        >
          7 Days
        </button>
        <button
          class="px-4 py-1.5 rounded-lg text-xs font-bold text-text-secondary hover:text-primary transition-all"
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
        class="premium-card group hover:border-primary/20 transition-all"
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
        <p class="text-[10px] font-bold text-text-secondary uppercase tracking-widest">
          {{ kpi.label }}
        </p>
        <h2 class="text-3xl font-bold text-text-primary mt-1">{{ kpi.value }}</h2>
        <div class="mt-4 flex items-center gap-2">
          <div class="flex-1 h-1 bg-surface-variant rounded-full overflow-hidden">
            <div
              class="h-full rounded-full transition-all duration-1000"
              :class="kpi.barClass"
              :style="{ width: kpi.progress + '%' }"
            ></div>
          </div>
          <span class="text-[10px] font-bold text-text-secondary">{{ kpi.progress }}%</span>
        </div>
      </div>
    </div>

    <!-- Middle Section: Activity & Charts -->
    <div class="grid grid-cols-1 lg:grid-cols-12 gap-8">
      <!-- High Priority Work Orders -->
      <div class="lg:col-span-8 flex flex-col">
        <div class="premium-card !p-0 overflow-hidden flex-1">
          <div
            class="p-6 border-b border-border-default flex justify-between items-center bg-surface-subtle/50"
          >
            <div>
              <h3 class="text-sm font-bold text-text-primary">Critical Work Orders</h3>
              <p class="text-[10px] text-text-secondary font-medium uppercase tracking-widest mt-1">
                Requiring immediate attention
              </p>
            </div>
            <button class="text-xs font-bold text-primary hover:text-primary/80">View All</button>
          </div>
          <div class="p-6 space-y-4">
            <div
              v-for="order in workOrders"
              :key="order.id"
              class="flex items-center gap-6 p-4 rounded-2xl bg-white border border-border-default hover:border-indigo-100 hover:shadow-md transition-all group"
            >
              <div
                class="w-12 h-12 rounded-xl flex items-center justify-center text-status-critical bg-error-container/20 font-bold text-xs uppercase"
              >
                {{ order.severity.charAt(0) }}
              </div>
              <div class="flex-1 min-w-0">
                <h4
                  class="text-sm font-bold text-text-primary group-hover:text-primary transition-colors"
                >
                  {{ order.title }}
                </h4>
                <div class="flex items-center gap-4 mt-1.5">
                  <span class="text-[10px] font-medium text-text-secondary flex items-center gap-1">
                    <span class="material-symbols-outlined text-xs">location_on</span>
                    {{ order.assetLocation }}
                  </span>
                  <span class="text-[10px] font-medium text-text-secondary flex items-center gap-1">
                    <span class="material-symbols-outlined text-xs">person</span>
                    {{ order.technician }}
                  </span>
                </div>
              </div>
              <div class="text-right">
                <span
                  class="text-[10px] font-bold uppercase tracking-wider px-2 py-1 rounded-lg bg-metric-amber/20 text-surface-tint"
                  >{{ order.status }}</span
                >
                <p class="text-[10px] text-text-secondary mt-2 font-medium">
                  {{ formatDate(order.targetDate) }}
                </p>
              </div>
            </div>
            <div
              v-if="workOrders.length === 0"
              class="py-12 text-center text-text-secondary text-sm italic"
            >
              No critical work orders found.
            </div>
          </div>
        </div>
      </div>

      <!-- Quick Actions & Alerts -->
      <div class="lg:col-span-4 space-y-8">
        <div class="premium-card">
          <h3 class="text-xs font-bold text-text-secondary uppercase tracking-widest mb-6">
            Quick Operations
          </h3>
          <div class="grid grid-cols-2 gap-4">
            <button
              class="flex flex-col items-center gap-3 p-4 rounded-2xl bg-primary-container/20 border border-primary-container/30 hover:bg-primary hover:text-on-primary transition-all group shadow-sm"
            >
              <span class="material-symbols-outlined text-primary group-hover:text-on-primary"
                >add_shopping_cart</span
              >
              <span class="text-[10px] font-bold uppercase tracking-wider">Procurement</span>
            </button>
            <button
              class="flex flex-col items-center gap-3 p-4 rounded-2xl bg-status-in-stock/10 border border-status-in-stock/20 hover:bg-status-in-stock hover:text-white transition-all group shadow-sm"
            >
              <span class="material-symbols-outlined text-status-in-stock group-hover:text-white"
                >handyman</span
              >
              <span class="text-[10px] font-bold uppercase tracking-wider">Service Log</span>
            </button>
            <button
              class="flex flex-col items-center gap-3 p-4 rounded-2xl bg-error-container/20 border border-status-critical/20 hover:bg-status-critical hover:text-white transition-all group shadow-sm"
            >
              <span class="material-symbols-outlined text-status-critical group-hover:text-white"
                >qr_code_scanner</span
              >
              <span class="text-[10px] font-bold uppercase tracking-wider">Audit Scan</span>
            </button>
            <button
              class="flex flex-col items-center gap-3 p-4 rounded-2xl bg-metric-amber/10 border border-metric-amber/30 hover:bg-surface-tint hover:text-white transition-all group shadow-sm"
            >
              <span class="material-symbols-outlined text-surface-tint group-hover:text-white"
                >sync_alt</span
              >
              <span class="text-[10px] font-bold uppercase tracking-wider">Transfer</span>
            </button>
          </div>
        </div>

        <!-- Recent Alerts -->
        <div class="premium-card">
          <div class="flex justify-between items-center mb-6">
            <h3 class="text-xs font-bold text-text-secondary uppercase tracking-widest">
              Recent Alerts
            </h3>
            <span class="w-2 h-2 rounded-full bg-status-critical animate-ping"></span>
          </div>
          <div class="space-y-6">
            <div v-for="alert in alerts" :key="alert.id" class="flex gap-4 group cursor-pointer">
              <div
                class="w-1 h-8 rounded-full transition-all group-hover:h-12"
                :class="alert.type === 'Critical' ? 'bg-status-critical' : 'bg-surface-tint'"
              ></div>
              <div class="flex-1 min-w-0">
                <h4
                  class="text-xs font-bold text-text-primary group-hover:text-primary transition-colors"
                >
                  {{ alert.title }}
                </h4>
                <p class="text-[10px] text-text-secondary mt-1 line-clamp-1">{{ alert.message }}</p>
              </div>
            </div>
            <div
              v-if="alerts.length === 0"
              class="py-4 text-center text-text-secondary text-xs italic"
            >
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
    bgClass: 'bg-primary-container/20',
    iconClass: 'text-primary',
    trend: '4.2%',
    trendIcon: 'arrow_upward',
    trendClass: 'text-status-in-stock bg-metric-sage border-status-in-stock/10',
    progress: 78,
    barClass: 'bg-primary',
  },
  {
    label: 'Operational Health',
    value:
      Math.round(
        (assets.value.filter((a) => a.status === 'In Stock').length / (assets.value.length || 1)) *
          100,
      ) + '%',
    icon: 'verified_user',
    bgClass: 'bg-metric-sage',
    iconClass: 'text-status-in-stock',
    trend: '99.9%',
    trendIcon: 'check_circle',
    trendClass: 'text-primary bg-primary-container/20 border-primary/10',
    progress: 92,
    barClass: 'bg-status-in-stock',
  },
  {
    label: 'Maintenance Queue',
    value: workOrders.value.length,
    icon: 'handyman',
    bgClass: 'bg-error-container/20',
    iconClass: 'text-status-critical',
    trend: '12%',
    trendIcon: 'arrow_downward',
    trendClass: 'text-status-critical bg-error-container border-status-critical/10',
    progress: 45,
    barClass: 'bg-status-critical',
  },
  {
    label: 'Est. Valuation',
    value: '₹' + (assets.value.reduce((s, a) => s + (a.value || 0), 0) / 1000000).toFixed(1) + 'M',
    icon: 'account_balance_wallet',
    bgClass: 'bg-secondary-container',
    iconClass: 'text-secondary',
    trend: '2.5%',
    trendIcon: 'arrow_upward',
    trendClass: 'text-secondary bg-secondary-container/50 border-secondary/10',
    progress: 60,
    barClass: 'bg-secondary',
  },
]);

const formatDate = (date: string) => {
  if (!date) return 'N/A';
  return new Date(date).toLocaleDateString('en-GB', { day: '2-digit', month: 'short' });
};

onMounted(fetchData);
</script>
