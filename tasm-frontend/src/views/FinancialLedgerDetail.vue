<template>
  <main class="flex-1 flex flex-col min-w-0">
    <!-- TopNavBar (Shared Component) -->
    <!-- Page Content -->
    <div class="p-page-margin w-full max-w-[1400px] mx-auto flex-1">
      <!-- Page Header -->
      <div class="flex flex-col md:flex-row md:items-end justify-between gap-4 mb-section-gap">
        <div>
          <h1 class="font-h1 text-h1 text-text-primary tracking-tight">Financial Ledger</h1>
          <p class="font-body text-body text-text-secondary mt-1">
            Comprehensive view of asset valuations and depreciation metrics.
          </p>
        </div>
        <div class="flex gap-3">
          <button
            class="px-4 py-2 rounded-lg bg-surface border border-border-default font-h3 text-h3 text-text-primary shadow-[0_2px_4px_rgba(0,0,0,0.02)] hover:bg-surface-subtle transition-colors flex items-center gap-2"
          >
            <span class="material-symbols-outlined text-[18px]"> download </span>
            Export CSV
          </button>
          <button
            class="px-4 py-2 rounded-lg bg-[#1C1917] text-white font-h3 text-h3 shadow-sm hover:opacity-90 transition-opacity flex items-center gap-2"
          >
            <span class="material-symbols-outlined text-[18px]"> filter_list </span>
            Filter View
          </button>
        </div>
      </div>
      <!-- Summary Cards (Bento Grid) -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-section-gap mb-section-gap">
        <!-- Total Purchase Cost -->
        <div
          class="bg-metric-amber rounded-xl p-card-padding border border-border-default shadow-[0_4px_4px_rgba(0,0,0,0.02)] hover:-translate-y-[2px] transition-transform cursor-pointer relative overflow-hidden group"
        >
          <div class="relative z-10">
            <div class="flex justify-between items-start mb-stack">
              <span class="font-h3 text-h3 text-on-surface-variant"> Total Purchase Cost </span>
              <div class="w-8 h-8 rounded-full bg-white/50 flex items-center justify-center">
                <span class="material-symbols-outlined text-[18px] text-surface-tint">
                  shopping_cart
                </span>
              </div>
            </div>
            <div class="font-kpi-number text-kpi-number text-on-surface">₹ 12,45,50,000</div>
            <div
              class="font-metadata text-metadata text-status-in-stock mt-2 flex items-center gap-1"
            >
              <span class="material-symbols-outlined text-[16px]"> trending_up </span>
              +4.2% YTD
            </div>
          </div>
          <div
            class="absolute -bottom-4 -right-4 text-surface-tint opacity-10 transform scale-150 group-hover:scale-110 transition-transform duration-500"
          >
            <span
              class="material-symbols-outlined text-[120px]"
              style="font-variation-settings: 'FILL' 1"
            >
              account_balance
            </span>
          </div>
        </div>
        <!-- Accumulated Depreciation -->
        <div
          class="bg-metric-lavender rounded-xl p-card-padding border border-border-default shadow-[0_4px_4px_rgba(0,0,0,0.02)] hover:-translate-y-[2px] transition-transform cursor-pointer relative overflow-hidden group"
        >
          <div class="relative z-10">
            <div class="flex justify-between items-start mb-stack">
              <span class="font-h3 text-h3 text-on-surface-variant">
                Accumulated Depreciation
              </span>
              <div class="w-8 h-8 rounded-full bg-white/50 flex items-center justify-center">
                <span class="material-symbols-outlined text-[18px] text-tertiary">
                  trending_down
                </span>
              </div>
            </div>
            <div class="font-kpi-number text-kpi-number text-on-surface">₹ 3,12,80,000</div>
            <div
              class="font-metadata text-metadata text-text-secondary mt-2 flex items-center gap-1"
            >
              Calculated via Straight-Line
            </div>
          </div>
          <div
            class="absolute -bottom-4 -right-4 text-tertiary opacity-10 transform scale-150 group-hover:scale-110 transition-transform duration-500"
          >
            <span
              class="material-symbols-outlined text-[120px]"
              style="font-variation-settings: 'FILL' 1"
            >
              show_chart
            </span>
          </div>
        </div>
        <!-- Net Book Value -->
        <div
          class="bg-metric-sage rounded-xl p-card-padding border border-border-default shadow-[0_4px_4px_rgba(0,0,0,0.02)] hover:-translate-y-[2px] transition-transform cursor-pointer relative overflow-hidden group"
        >
          <div class="relative z-10">
            <div class="flex justify-between items-start mb-stack">
              <span class="font-h3 text-h3 text-on-surface-variant"> Net Book Value </span>
              <div class="w-8 h-8 rounded-full bg-white/50 flex items-center justify-center">
                <span class="material-symbols-outlined text-[18px] text-status-in-stock">
                  verified
                </span>
              </div>
            </div>
            <div class="font-kpi-number text-kpi-number text-on-surface">₹ 9,32,70,000</div>
            <div
              class="font-metadata text-metadata text-status-in-stock mt-2 flex items-center gap-1"
            >
              <span class="material-symbols-outlined text-[16px]"> check_circle </span>
              Audited 12 Oct 2023
            </div>
          </div>
          <div
            class="absolute -bottom-4 -right-4 text-status-in-stock opacity-10 transform scale-150 group-hover:scale-110 transition-transform duration-500"
          >
            <span
              class="material-symbols-outlined text-[120px]"
              style="font-variation-settings: 'FILL' 1"
            >
              savings
            </span>
          </div>
        </div>
      </div>
      <!-- Chart Area -->
      <div
        class="bg-surface rounded-xl p-card-padding border border-border-default shadow-[0_4px_4px_rgba(0,0,0,0.02)] mb-section-gap flex flex-col h-[400px]"
      >
        <div class="flex justify-between items-center mb-6">
          <h2 class="font-h2 text-h2 text-text-primary">
            Asset Value vs Depreciation Trend (FY 2023-24)
          </h2>
          <div class="flex gap-4 items-center">
            <div class="flex items-center gap-2">
              <div class="w-3 h-3 rounded-full bg-surface-tint"></div>
              <span class="font-metadata text-metadata text-text-secondary"> Gross Value </span>
            </div>
            <div class="flex items-center gap-2">
              <div class="w-3 h-3 rounded-full bg-tertiary"></div>
              <span class="font-metadata text-metadata text-text-secondary"> Depreciation </span>
            </div>
          </div>
        </div>
        <!-- CSS Area Chart Visual Representation -->
        <div class="flex-1 relative mt-4 border-l border-b border-border-default flex items-end">
          <!-- Y Axis Labels -->
          <div
            class="absolute -left-12 top-0 h-full flex flex-col justify-between items-end py-2 font-mono-data text-mono-data text-text-secondary"
          >
            <span> 12 Cr </span>
            <span> 9 Cr </span>
            <span> 6 Cr </span>
            <span> 3 Cr </span>
            <span> 0 </span>
          </div>
          <!-- Grid Lines -->
          <div class="absolute inset-0 flex flex-col justify-between z-0">
            <div class="border-t border-border-default w-full opacity-50"></div>
            <div class="border-t border-border-default w-full opacity-50"></div>
            <div class="border-t border-border-default w-full opacity-50"></div>
            <div class="border-t border-border-default w-full opacity-50"></div>
            <div class="border-t border-border-default w-full opacity-50"></div>
          </div>
          <!-- Chart SVG Placeholder (using pure CSS/SVG shapes for visual fidelity) -->
          <div class="w-full h-full relative z-10 px-4 flex items-end">
            <svg
              class="w-full h-full overflow-visible"
              preserveaspectratio="none"
              viewbox="0 0 100 100"
            >
              <!-- Gross Value Area -->
              <path
                d="M0,10 L20,15 L40,12 L60,20 L80,18 L100,25 L100,100 L0,100 Z"
                fill="#fff1eb"
                opacity="0.6"
              ></path>
              <path
                d="M0,10 L20,15 L40,12 L60,20 L80,18 L100,25"
                fill="none"
                stroke="#9a4614"
                stroke-width="2"
              ></path>
              <!-- Depreciation Area -->
              <path
                d="M0,90 L20,85 L40,80 L60,70 L80,60 L100,50 L100,100 L0,100 Z"
                fill="#cee5ff"
                opacity="0.5"
              ></path>
              <path
                d="M0,90 L20,85 L40,80 L60,70 L80,60 L100,50"
                fill="none"
                stroke="#005c92"
                stroke-width="2"
              ></path>
            </svg>
          </div>
          <!-- X Axis Labels -->
          <div
            class="absolute -bottom-8 left-0 w-full flex justify-between px-4 font-metadata text-metadata text-text-secondary"
          >
            <span> Q1 </span>
            <span> Q2 </span>
            <span> Q3 </span>
            <span> Q4 </span>
          </div>
        </div>
      </div>
      <!-- High-Density Table Section -->
      <div
        class="bg-surface rounded-xl border border-border-default shadow-[0_4px_4px_rgba(0,0,0,0.02)] flex flex-col overflow-hidden"
      >
        <div
          class="px-card-padding py-4 border-b border-border-default flex justify-between items-center bg-surface"
        >
          <h2 class="font-h2 text-h2 text-text-primary">Recent Ledger Entries</h2>
          <div class="relative">
            <span
              class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-stone-400 text-[16px]"
            >
              search
            </span>
            <input
              class="pl-9 pr-4 py-1.5 rounded-lg border border-border-default bg-surface-subtle text-sm focus:outline-none focus:ring-1 focus:ring-surface-tint focus:border-surface-tint w-[200px]"
              placeholder="Search entries..."
              type="text"
            />
          </div>
        </div>
        <div class="overflow-x-auto p-4">
          <DataTable
            :value="ledgers"
            :loading="loading"
            paginator
            :rows="10"
            tableStyle="min-width: 50rem"
            class="w-full text-left"
          >
            <Column field="transactionId" header="ENTRY ID" sortable>
              <template #body="slotProps">
                <span class="font-mono-data text-mono-data text-surface-tint">
                  {{ slotProps.data.transactionId }}
                </span>
              </template>
            </Column>
            <Column field="date" header="DATE" sortable>
              <template #body="slotProps">
                <span class="whitespace-nowrap">{{
                  new Date(slotProps.data.date).toLocaleDateString()
                }}</span>
              </template>
            </Column>
            <Column field="category" header="CATEGORY" sortable>
              <template #body="slotProps">
                <span
                  class="inline-flex items-center px-2 py-0.5 rounded-md bg-surface-container text-primary-container font-metadata text-metadata border border-surface-dim"
                >
                  {{ slotProps.data.category }}
                </span>
              </template>
            </Column>
            <Column field="description" header="DESCRIPTION">
              <template #body="slotProps">
                <div class="max-w-[200px] truncate" :title="slotProps.data.description">
                  {{ slotProps.data.description }}
                </div>
              </template>
            </Column>
            <Column header="DEBIT (₹)" alignFrozen="right">
              <template #body="slotProps">
                <span v-if="slotProps.data.type === 'Debit'" class="font-mono-data text-mono-data">
                  {{ slotProps.data.amount?.toLocaleString() }}
                </span>
                <span v-else class="font-mono-data text-mono-data text-text-secondary">-</span>
              </template>
            </Column>
            <Column header="CREDIT (₹)" alignFrozen="right">
              <template #body="slotProps">
                <span v-if="slotProps.data.type === 'Credit'" class="font-mono-data text-mono-data">
                  {{ slotProps.data.amount?.toLocaleString() }}
                </span>
                <span v-else class="font-mono-data text-mono-data text-text-secondary">-</span>
              </template>
            </Column>
          </DataTable>
        </div>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';

interface LedgerEntry {
  id: number;
  transactionId: string;
  date: string;
  description: string;
  amount: number;
  type: string;
  category: string;
}

const ledgers = ref<LedgerEntry[]>([]);
const loading = ref(true);

const fetchLedgers = async () => {
  try {
    const res = await fetch('http://localhost:8080/api/ledgers');
    if (res.ok) {
      ledgers.value = await res.json();
    }
  } catch (error) {
    console.error('Failed to fetch ledgers:', error);
  } finally {
    loading.value = false;
  }
};

onMounted(() => {
  fetchLedgers();
});
</script>
