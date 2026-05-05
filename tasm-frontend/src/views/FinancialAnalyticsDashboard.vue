<template>
  <main class="space-y-section-gap pb-24">
    <!-- Header -->
    <div class="flex flex-col md:flex-row md:items-end justify-between gap-4 mb-section-gap">
      <div>
        <h2 class="font-h1 text-h1 text-text-primary mb-2">Financial Analytics</h2>
        <p class="font-body text-body text-text-secondary">
          Comprehensive overview of asset valuations and maintenance expenditures.
        </p>
      </div>
      <div class="flex items-center gap-3">
        <button
          @click="handleRefresh"
          class="p-2 bg-surface text-text-secondary border border-border-default rounded-lg hover:bg-surface-subtle transition-colors"
          title="Refresh Data"
          :disabled="loading"
        >
          <span class="material-symbols-outlined text-[20px]" :class="{ 'animate-spin': loading }">
            refresh
          </span>
        </button>
        <button
          @click="handleExport"
          class="px-4 py-2 bg-[#1C1917] text-white rounded-lg font-h3 text-h3 hover:bg-stone-800 transition-colors flex items-center gap-2 shadow-sm"
        >
          <span class="material-symbols-outlined text-[18px]"> download </span>
          Export Report
        </button>
      </div>
    </div>
    <!-- KPI Grid -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-section-gap">
      <div
        @click="navigateToLedger"
        class="bg-surface border border-border-default rounded-xl p-card-padding shadow-sm hover:-translate-y-[2px] transition-transform cursor-pointer group"
      >
        <div class="flex items-center justify-between mb-4">
          <span class="font-metadata text-metadata text-text-secondary uppercase tracking-wider">
            Net Asset Value
          </span>
          <span
            class="material-symbols-outlined text-primary-container bg-surface-subtle p-2 rounded-lg"
          >
            account_balance
          </span>
        </div>
        <div class="font-kpi-number text-kpi-number text-text-primary">
          ₹{{ netAssetValue.toLocaleString() }}
        </div>
        <div class="flex items-center gap-2 mt-2">
          <span class="text-status-in-stock text-xs font-medium flex items-center gap-0.5">
            <span class="material-symbols-outlined text-[14px]"> trending_up </span>
            +4.2%
          </span>
          <span class="font-metadata text-[11px] text-text-secondary"> vs last quarter </span>
        </div>
      </div>
      <div
        @click="navigateToLedger"
        class="bg-surface border border-border-default rounded-xl p-card-padding shadow-sm hover:-translate-y-[2px] transition-transform cursor-pointer group"
      >
        <div class="flex items-center justify-between mb-4">
          <span class="font-metadata text-metadata text-text-secondary uppercase tracking-wider">
            Depreciation YTD
          </span>
          <span class="material-symbols-outlined text-error bg-error/5 p-2 rounded-lg">
            trending_down
          </span>
        </div>
        <div class="font-kpi-number text-kpi-number text-text-primary">
          ₹{{ totalDepreciation.toLocaleString() }}
        </div>
        <div class="font-metadata text-[11px] text-text-secondary mt-2">
          Calculated using Straight-Line method
        </div>
      </div>
      <div
        class="bg-metric-sage border border-status-in-stock/10 rounded-xl p-card-padding shadow-sm hover:-translate-y-[2px] transition-transform"
      >
        <div class="flex items-center justify-between mb-4">
          <span
            class="font-metadata text-metadata text-status-in-stock uppercase tracking-wider font-semibold"
          >
            Maintenance Spend
          </span>
          <span class="material-symbols-outlined text-status-in-stock bg-white/50 p-2 rounded-lg">
            build
          </span>
        </div>
        <div class="font-kpi-number text-kpi-number text-text-primary">
          ₹{{ maintenanceSpend.toLocaleString() }}
        </div>
      </div>
      <!-- KPI 4 -->
      <div
        class="bg-metric-sage rounded-xl p-card-padding border border-border-default shadow-sm flex flex-col justify-between hover:-translate-y-0.5 transition-transform cursor-pointer"
      >
        <div class="flex justify-between items-start mb-4">
          <span class="font-metadata text-metadata text-on-surface-variant font-medium">
            Disposal Value
          </span>
          <div class="w-8 h-8 rounded-full bg-white/50 flex items-center justify-center">
            <span class="material-symbols-outlined text-primary text-[18px]"> recycling </span>
          </div>
        </div>
        <div>
          <div class="font-kpi-number text-kpi-number text-text-primary">₹1,18,000</div>
          <div class="flex items-center gap-1 mt-2 text-text-secondary font-metadata text-metadata">
            <span class="material-symbols-outlined text-[14px]"> horizontal_rule </span>
            <span> On track </span>
          </div>
        </div>
      </div>
    </div>
    <!-- Middle Row: Charts -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-inline">
      <!-- Line Chart: Depreciation Forecast -->
      <div
        class="bg-surface rounded-xl p-card-padding border border-border-default shadow-sm flex flex-col"
      >
        <div class="flex justify-between items-center mb-6">
          <h2 class="font-h2 text-h2 text-text-primary">Depreciation Forecast</h2>
          <span
            class="material-symbols-outlined text-text-secondary cursor-pointer hover:text-primary"
          >
            more_vert
          </span>
        </div>
        <div class="flex-1 min-h-[280px] relative w-full pt-4">
          <!-- Simulated Recharts LineChart using SVG -->
          <svg class="chart-svg" preserveaspectratio="none" viewbox="0 0 500 250">
            <!-- Grid Lines -->
            <line
              stroke="#EEEBE4"
              stroke-dasharray="4"
              stroke-width="1"
              x1="40"
              x2="480"
              y1="20"
              y2="20"
            ></line>
            <line
              stroke="#EEEBE4"
              stroke-dasharray="4"
              stroke-width="1"
              x1="40"
              x2="480"
              y1="80"
              y2="80"
            ></line>
            <line
              stroke="#EEEBE4"
              stroke-dasharray="4"
              stroke-width="1"
              x1="40"
              x2="480"
              y1="140"
              y2="140"
            ></line>
            <line stroke="#EEEBE4" stroke-width="1" x1="40" x2="480" y1="200" y2="200"></line>
            <!-- Y-Axis Labels -->
            <text fill="#78716C" font-family="Inter" font-size="10" text-anchor="end" x="30" y="25">
              80L
            </text>
            <text fill="#78716C" font-family="Inter" font-size="10" text-anchor="end" x="30" y="85">
              60L
            </text>
            <text
              fill="#78716C"
              font-family="Inter"
              font-size="10"
              text-anchor="end"
              x="30"
              y="145"
            >
              40L
            </text>
            <text
              fill="#78716C"
              font-family="Inter"
              font-size="10"
              text-anchor="end"
              x="30"
              y="205"
            >
              20L
            </text>
            <!-- X-Axis Labels -->
            <text
              fill="#78716C"
              font-family="Inter"
              font-size="10"
              text-anchor="middle"
              x="80"
              y="225"
            >
              2024
            </text>
            <text
              fill="#78716C"
              font-family="Inter"
              font-size="10"
              text-anchor="middle"
              x="180"
              y="225"
            >
              2025
            </text>
            <text
              fill="#78716C"
              font-family="Inter"
              font-size="10"
              text-anchor="middle"
              x="280"
              y="225"
            >
              2026
            </text>
            <text
              fill="#78716C"
              font-family="Inter"
              font-size="10"
              text-anchor="middle"
              x="380"
              y="225"
            >
              2027
            </text>
            <text
              fill="#78716C"
              font-family="Inter"
              font-size="10"
              text-anchor="middle"
              x="480"
              y="225"
            >
              2028
            </text>
            <!-- Data Line & Area -->
            <path
              d="M80,40 L180,90 L280,130 L380,160 L480,180"
              fill="none"
              stroke="#712c00"
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="3"
            ></path>
            <!-- Area Gradient Fill -->
            <defs>
              <lineargradient id="areaGradient" x1="0%" x2="0%" y1="0%" y2="100%">
                <stop offset="0%" stop-color="#712c00" stop-opacity="0.2"></stop>
                <stop offset="100%" stop-color="#712c00" stop-opacity="0"></stop>
              </lineargradient>
            </defs>
            <path
              d="M80,40 L180,90 L280,130 L380,160 L480,180 L480,200 L80,200 Z"
              fill="url(#areaGradient)"
            ></path>
            <!-- Data Points -->
            <circle cx="80" cy="40" fill="#FFFFFF" r="4" stroke="#712c00" stroke-width="2"></circle>
            <circle
              cx="180"
              cy="90"
              fill="#FFFFFF"
              r="4"
              stroke="#712c00"
              stroke-width="2"
            ></circle>
            <circle
              cx="280"
              cy="130"
              fill="#FFFFFF"
              r="4"
              stroke="#712c00"
              stroke-width="2"
            ></circle>
            <circle
              cx="380"
              cy="160"
              fill="#FFFFFF"
              r="4"
              stroke="#712c00"
              stroke-width="2"
            ></circle>
            <circle
              cx="480"
              cy="180"
              fill="#FFFFFF"
              r="4"
              stroke="#712c00"
              stroke-width="2"
            ></circle>
          </svg>
        </div>
      </div>
      <!-- Bar Chart: Asset Value by Category -->
      <div
        class="bg-surface rounded-xl p-card-padding border border-border-default shadow-sm flex flex-col"
      >
        <div class="flex justify-between items-center mb-6">
          <h2 class="font-h2 text-h2 text-text-primary">Asset Value by Category</h2>
          <span
            class="material-symbols-outlined text-text-secondary cursor-pointer hover:text-primary"
          >
            more_vert
          </span>
        </div>
        <div class="flex-1 min-h-[280px] flex items-end gap-4 pt-4 pb-8 relative w-full px-8">
          <!-- Y-Axis (Simulated visually) -->
          <div
            class="absolute left-0 top-0 bottom-8 w-8 flex flex-col justify-between text-[10px] text-text-secondary font-metadata font-medium text-right pr-2"
          >
            <span> 40L </span>
            <span> 30L </span>
            <span> 20L </span>
            <span> 10L </span>
            <span> 0 </span>
          </div>
          <!-- Grid Lines -->
          <div
            class="absolute left-8 right-8 top-0 bottom-8 border-l border-b border-border-default z-0"
          >
            <div class="w-full border-t border-border-default border-dashed absolute top-0"></div>
            <div class="w-full border-t border-border-default border-dashed absolute top-1/4"></div>
            <div class="w-full border-t border-border-default border-dashed absolute top-2/4"></div>
            <div class="w-full border-t border-border-default border-dashed absolute top-3/4"></div>
          </div>
          <!-- Bars -->
          <div
            class="flex-1 flex flex-col items-center justify-end h-full z-10 group cursor-pointer relative"
          >
            <div
              class="w-full max-w-[48px] bg-primary-fixed-dim rounded-t-sm h-[80%] transition-all group-hover:bg-primary-fixed relative"
            >
              <div
                class="opacity-0 group-hover:opacity-100 absolute -top-8 left-1/2 -translate-x-1/2 bg-text-primary text-white text-[10px] py-1 px-2 rounded whitespace-nowrap transition-opacity"
              >
                ₹32,00,000
              </div>
            </div>
            <span
              class="absolute -bottom-6 text-[11px] font-metadata text-text-secondary text-center w-full truncate"
            >
              IT Equip
            </span>
          </div>
          <div
            class="flex-1 flex flex-col items-center justify-end h-full z-10 group cursor-pointer relative"
          >
            <div
              class="w-full max-w-[48px] bg-[#FEF3C7] rounded-t-sm h-[60%] transition-all group-hover:brightness-95 relative border border-[#EEEBE4]"
            >
              <div
                class="opacity-0 group-hover:opacity-100 absolute -top-8 left-1/2 -translate-x-1/2 bg-text-primary text-white text-[10px] py-1 px-2 rounded whitespace-nowrap transition-opacity"
              >
                ₹24,00,000
              </div>
            </div>
            <span
              class="absolute -bottom-6 text-[11px] font-metadata text-text-secondary text-center w-full truncate"
            >
              Furniture
            </span>
          </div>
          <div
            class="flex-1 flex flex-col items-center justify-end h-full z-10 group cursor-pointer relative"
          >
            <div
              class="w-full max-w-[48px] bg-[#EDE9FE] rounded-t-sm h-[35%] transition-all group-hover:brightness-95 relative border border-[#EEEBE4]"
            >
              <div
                class="opacity-0 group-hover:opacity-100 absolute -top-8 left-1/2 -translate-x-1/2 bg-text-primary text-white text-[10px] py-1 px-2 rounded whitespace-nowrap transition-opacity"
              >
                ₹14,00,000
              </div>
            </div>
            <span
              class="absolute -bottom-6 text-[11px] font-metadata text-text-secondary text-center w-full truncate"
            >
              Vehicles
            </span>
          </div>
          <div
            class="flex-1 flex flex-col items-center justify-end h-full z-10 group cursor-pointer relative"
          >
            <div
              class="w-full max-w-[48px] bg-[#DCFCE7] rounded-t-sm h-[45%] transition-all group-hover:brightness-95 relative border border-[#EEEBE4]"
            >
              <div
                class="opacity-0 group-hover:opacity-100 absolute -top-8 left-1/2 -translate-x-1/2 bg-text-primary text-white text-[10px] py-1 px-2 rounded whitespace-nowrap transition-opacity"
              >
                ₹18,00,000
              </div>
            </div>
            <span
              class="absolute -bottom-6 text-[11px] font-metadata text-text-secondary text-center w-full truncate"
            >
              HVAC
            </span>
          </div>
        </div>
      </div>
    </div>
    <!-- Bottom Row: Cost of Ownership Ranking -->
    <div class="bg-surface rounded-xl p-card-padding border border-border-default shadow-sm mb-8">
      <div class="flex justify-between items-center mb-6">
        <div>
          <h2 class="font-h2 text-h2 text-text-primary">Cost of Ownership</h2>
          <p class="font-metadata text-metadata text-text-secondary mt-1">
            Highest accumulating costs across all categories (YTD)
          </p>
        </div>
        <button
          @click="router.push('/report-builder')"
          class="text-sm font-medium text-primary hover:underline"
        >
          View Full Report
        </button>
      </div>
      <div class="flex flex-col gap-4">
        <!-- Row 1 -->
        <div class="flex items-center gap-4 group">
          <div
            class="w-10 h-10 rounded-lg bg-surface-subtle border border-border-default flex items-center justify-center shrink-0"
          >
            <span class="material-symbols-outlined text-stone-600"> dns </span>
          </div>
          <div class="w-48 shrink-0">
            <div class="font-medium text-sm text-text-primary truncate">
              Main Data Center Server
            </div>
            <div class="font-metadata text-metadata text-text-secondary">IT Equipment</div>
          </div>
          <div class="flex-1 flex items-center gap-4">
            <div class="flex-1 h-2 bg-surface-container-high rounded-full overflow-hidden">
              <div class="h-full bg-primary rounded-full w-[85%]"></div>
            </div>
            <div
              class="w-24 text-right font-mono-data text-mono-data text-text-primary font-medium"
            >
              ₹8,45,000
            </div>
          </div>
        </div>
        <!-- Row 2 -->
        <div class="flex items-center gap-4 group">
          <div
            class="w-10 h-10 rounded-lg bg-surface-subtle border border-border-default flex items-center justify-center shrink-0"
          >
            <span class="material-symbols-outlined text-stone-600"> directions_car </span>
          </div>
          <div class="w-48 shrink-0">
            <div class="font-medium text-sm text-text-primary truncate">
              Security Patrol Vehicle A
            </div>
            <div class="font-metadata text-metadata text-text-secondary">Vehicles</div>
          </div>
          <div class="flex-1 flex items-center gap-4">
            <div class="flex-1 h-2 bg-surface-container-high rounded-full overflow-hidden">
              <div class="h-full bg-outline rounded-full w-[60%]"></div>
            </div>
            <div
              class="w-24 text-right font-mono-data text-mono-data text-text-primary font-medium"
            >
              ₹5,12,000
            </div>
          </div>
        </div>
        <!-- Row 3 -->
        <div class="flex items-center gap-4 group">
          <div
            class="w-10 h-10 rounded-lg bg-surface-subtle border border-border-default flex items-center justify-center shrink-0"
          >
            <span class="material-symbols-outlined text-stone-600"> ac_unit </span>
          </div>
          <div class="w-48 shrink-0">
            <div class="font-medium text-sm text-text-primary truncate">Central HVAC Unit 3</div>
            <div class="font-metadata text-metadata text-text-secondary">HVAC</div>
          </div>
          <div class="flex-1 flex items-center gap-4">
            <div class="flex-1 h-2 bg-surface-container-high rounded-full overflow-hidden">
              <div class="h-full bg-outline-variant rounded-full w-[45%]"></div>
            </div>
            <div
              class="w-24 text-right font-mono-data text-mono-data text-text-primary font-medium"
            >
              ₹3,90,000
            </div>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { getDepreciations } from '../api/financial';
import { getWorkOrders } from '../api/workOrders';

const router = useRouter();
const netAssetValue = ref(0);
const totalDepreciation = ref(0);
const maintenanceSpend = ref(0);
const loading = ref(false);

const fetchMetrics = async () => {
  loading.value = true;
  try {
    const depreciations = (await getDepreciations()) as any[];
    netAssetValue.value = depreciations.reduce(
      (sum: number, item: any) => sum + (item.currentValue || 0),
      0,
    );
    totalDepreciation.value = depreciations.reduce(
      (sum: number, item: any) => sum + ((item.purchaseValue || 0) - (item.currentValue || 0)),
      0,
    );

    const workOrders = (await getWorkOrders()) as any[];
    maintenanceSpend.value = workOrders
      .filter((wo: any) => wo.status === 'Closed')
      .reduce((sum: number, wo: any) => sum + (wo.cost || 0), 0);
  } catch (error) {
    console.error('Failed to fetch analytics metrics:', error);
  } finally {
    loading.value = false;
  }
};

const handleExport = () => {
  const data = [
    ['Metric', 'Value'],
    ['Net Asset Value', `₹${netAssetValue.value.toLocaleString()}`],
    ['Depreciation YTD', `₹${totalDepreciation.value.toLocaleString()}`],
    ['Maintenance Spend', `₹${maintenanceSpend.value.toLocaleString()}`],
  ];

  const csvContent = data.map((row) => row.join(',')).join('\n');
  const blob = new Blob([csvContent], { type: 'text/csv;charset=utf-8;' });
  const link = document.createElement('a');
  link.href = URL.createObjectURL(blob);
  link.download = `tasm_financial_report_${new Date().toISOString().split('T')[0]}.csv`;
  link.click();
};

const handleRefresh = () => {
  fetchMetrics();
};

const navigateToLedger = () => {
  router.push({ name: 'AssetDepreciationLedger' });
};

onMounted(() => {
  fetchMetrics();
});
</script>
