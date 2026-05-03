<template>
  <main class="space-y-section-gap pb-24">
    <!-- Header -->
    <!-- Top Row: Pastel KPI Cards -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-inline">
      <!-- KPI 1 -->
      <div
        class="bg-metric-amber rounded-xl p-card-padding border border-border-default shadow-sm flex flex-col justify-between hover:-translate-y-0.5 transition-transform cursor-pointer"
      >
        <div class="flex justify-between items-start mb-4">
          <span class="font-metadata text-metadata text-on-surface-variant font-medium">
            Net Asset Value
          </span>
          <div class="w-8 h-8 rounded-full bg-white/50 flex items-center justify-center">
            <span class="material-symbols-outlined text-primary text-[18px]">
              account_balance
            </span>
          </div>
        </div>
        <div>
          <div class="font-kpi-number text-kpi-number text-text-primary">
            ₹{{ netAssetValue.toLocaleString() }}
          </div>
          <div
            class="flex items-center gap-1 mt-2 text-status-in-stock font-metadata text-metadata"
          >
            <span class="material-symbols-outlined text-[14px]"> trending_up </span>
            <span> +4.2% from last quarter </span>
          </div>
        </div>
      </div>
      <!-- KPI 2 -->
      <div
        class="bg-metric-lavender rounded-xl p-card-padding border border-border-default shadow-sm flex flex-col justify-between hover:-translate-y-0.5 transition-transform cursor-pointer"
      >
        <div class="flex justify-between items-start mb-4">
          <span class="font-metadata text-metadata text-on-surface-variant font-medium">
            Depreciation YTD
          </span>
          <div class="w-8 h-8 rounded-full bg-white/50 flex items-center justify-center">
            <span class="material-symbols-outlined text-primary text-[18px]"> trending_down </span>
          </div>
        </div>
        <div>
          <div class="font-kpi-number text-kpi-number text-text-primary">
            ₹{{ totalDepreciation.toLocaleString() }}
          </div>
          <div
            class="flex items-center gap-1 mt-2 text-status-critical font-metadata text-metadata"
          >
            <span class="material-symbols-outlined text-[14px]"> arrow_upward </span>
            <span> +1.1% vs projected </span>
          </div>
        </div>
      </div>
      <!-- KPI 3 -->
      <div
        class="bg-surface-container-high rounded-xl p-card-padding border border-border-default shadow-sm flex flex-col justify-between hover:-translate-y-0.5 transition-transform cursor-pointer"
      >
        <div class="flex justify-between items-start mb-4">
          <span class="font-metadata text-metadata text-on-surface-variant font-medium">
            Maintenance Spend
          </span>
          <div class="w-8 h-8 rounded-full bg-white/50 flex items-center justify-center">
            <span class="material-symbols-outlined text-primary text-[18px]"> build_circle </span>
          </div>
        </div>
        <div>
          <div class="font-kpi-number text-kpi-number text-text-primary">
            ₹{{ maintenanceSpend.toLocaleString() }}
          </div>
          <div
            class="flex items-center gap-1 mt-2 text-status-in-stock font-metadata text-metadata"
          >
            <span class="material-symbols-outlined text-[14px]"> arrow_downward </span>
            <span> -2.4% vs budget </span>
          </div>
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
        <button class="text-sm font-medium text-primary hover:underline">View Full Report</button>
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
import { getDepreciations } from '../api/financial';
import { getWorkOrders } from '../api/workOrders';

const netAssetValue = ref(0);
const totalDepreciation = ref(0);
const maintenanceSpend = ref(0);

const fetchMetrics = async () => {
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
  }
};

onMounted(() => {
  fetchMetrics();
});
</script>
