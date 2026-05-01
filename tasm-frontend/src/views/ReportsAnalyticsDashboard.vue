<template>
  <main class="flex-1 flex flex-col min-h-screen relative">
    <div class="p-page-margin max-w-[1400px] w-full mx-auto flex flex-col gap-section-gap pb-24">
      <!-- Page Header -->
      <div class="flex items-end justify-between">
        <div>
          <h1 class="font-h1 text-h1 text-text-primary tracking-tight">
            Financial Reports &amp; Analytics
          </h1>
          <p class="font-body text-body text-text-secondary mt-1">Fiscal Year 2023-2024</p>
        </div>
        <div class="flex gap-inline">
          <button
            class="bg-surface border border-border-default text-text-primary hover:bg-surface-subtle font-body text-body py-2 px-4 rounded-lg shadow-sm flex items-center gap-2 transition-transform hover:-translate-y-0.5"
          >
            <span class="material-symbols-outlined text-sm">download</span>
            Export PDF
          </button>
          <button
            class="bg-[#1C1917] text-white hover:bg-stone-800 font-body text-body py-2 px-4 rounded-lg shadow-sm flex items-center gap-2 transition-transform hover:-translate-y-0.5"
          >
            <span class="material-symbols-outlined text-sm">filter_list</span>
            Filter Dates
          </button>
        </div>
      </div>
      <!-- KPI Cards Grid -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-stack">
        <!-- KPI 1 -->
        <div
          class="bg-surface border border-border-default rounded-xl p-card-padding shadow-[0_4px_16px_-4px_rgba(0,0,0,0.05)] hover:-translate-y-0.5 transition-transform cursor-pointer relative overflow-hidden group"
        >
          <div
            class="absolute inset-0 bg-metric-amber opacity-10 group-hover:opacity-20 transition-opacity"
          ></div>
          <div class="relative z-10">
            <div class="flex items-center justify-between mb-4">
              <span class="font-metadata text-metadata text-text-secondary uppercase tracking-wider"
                >Total Purchase Cost</span
              >
              <span class="material-symbols-outlined text-primary/60">account_balance_wallet</span>
            </div>
            <div class="font-kpi-number text-kpi-number text-text-primary mb-1">
              <span class="font-mono-data text-xl align-top mr-1">₹</span>1,24,50,000
            </div>
            <div class="font-metadata text-metadata text-status-in-stock flex items-center gap-1">
              <span class="material-symbols-outlined text-[14px]">trending_up</span>
              +12.5% from last year
            </div>
          </div>
        </div>
        <!-- KPI 2 -->
        <div
          class="bg-surface border border-border-default rounded-xl p-card-padding shadow-[0_4px_16px_-4px_rgba(0,0,0,0.05)] hover:-translate-y-0.5 transition-transform cursor-pointer relative overflow-hidden group"
        >
          <div
            class="absolute inset-0 bg-metric-sage opacity-10 group-hover:opacity-20 transition-opacity"
          ></div>
          <div class="relative z-10">
            <div class="flex items-center justify-between mb-4">
              <span class="font-metadata text-metadata text-text-secondary uppercase tracking-wider"
                >Net Book Value</span
              >
              <span class="material-symbols-outlined text-status-in-stock/60"
                >domain_verification</span
              >
            </div>
            <div class="font-kpi-number text-kpi-number text-text-primary mb-1">
              <span class="font-mono-data text-xl align-top mr-1">₹</span>82,30,000
            </div>
            <div class="font-metadata text-metadata text-text-secondary flex items-center gap-1">
              Current active assets valuation
            </div>
          </div>
        </div>
        <!-- KPI 3 -->
        <div
          class="bg-surface border border-border-default rounded-xl p-card-padding shadow-[0_4px_16px_-4px_rgba(0,0,0,0.05)] hover:-translate-y-0.5 transition-transform cursor-pointer relative overflow-hidden group"
        >
          <div
            class="absolute inset-0 bg-metric-lavender opacity-10 group-hover:opacity-20 transition-opacity"
          ></div>
          <div class="relative z-10">
            <div class="flex items-center justify-between mb-4">
              <span class="font-metadata text-metadata text-text-secondary uppercase tracking-wider"
                >Accumulated Depreciation</span
              >
              <span class="material-symbols-outlined text-status-critical/60">trending_down</span>
            </div>
            <div class="font-kpi-number text-kpi-number text-text-primary mb-1">
              <span class="font-mono-data text-xl align-top mr-1">₹</span>42,20,000
            </div>
            <div class="font-metadata text-metadata text-status-critical flex items-center gap-1">
              <span class="material-symbols-outlined text-[14px]">warning</span>
              Requires audit review
            </div>
          </div>
        </div>
      </div>
      <!-- Charts Section -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-stack">
        <!-- Bar Chart: Depreciation by Category -->
        <div
          class="bg-surface border border-border-default rounded-xl p-card-padding shadow-[0_4px_16px_-4px_rgba(0,0,0,0.05)] flex flex-col h-[360px]"
        >
          <div class="flex items-center justify-between mb-6">
            <h2 class="font-h2 text-h2 text-text-primary">Depreciation by Category</h2>
            <button class="text-text-secondary hover:text-text-primary">
              <span class="material-symbols-outlined">more_vert</span>
            </button>
          </div>
          <div
            class="flex-1 flex items-end justify-between gap-2 pt-8 pb-4 border-b border-border-default relative"
          >
            <div
              class="absolute left-0 top-0 bottom-0 flex flex-col justify-between text-xs text-text-secondary font-mono-data pointer-events-none w-8 border-r border-border-default border-dashed"
            >
              <span>15L</span>
              <span>10L</span>
              <span>5L</span>
              <span>0</span>
            </div>
            <div class="pl-10 flex-1 flex items-end justify-around h-full">
              <div class="w-12 bg-[#1C1917] rounded-t-sm h-[80%] relative group">
                <div
                  class="absolute -top-8 left-1/2 -translate-x-1/2 bg-inverse-surface text-inverse-on-surface text-xs py-1 px-2 rounded opacity-0 group-hover:opacity-100 transition-opacity font-mono-data whitespace-nowrap"
                >
                  ₹12.5L
                </div>
              </div>
              <div class="w-12 bg-primary/80 rounded-t-sm h-[60%] relative group">
                <div
                  class="absolute -top-8 left-1/2 -translate-x-1/2 bg-inverse-surface text-inverse-on-surface text-xs py-1 px-2 rounded opacity-0 group-hover:opacity-100 transition-opacity font-mono-data whitespace-nowrap"
                >
                  ₹9.2L
                </div>
              </div>
              <div class="w-12 bg-secondary/60 rounded-t-sm h-[40%] relative group">
                <div
                  class="absolute -top-8 left-1/2 -translate-x-1/2 bg-inverse-surface text-inverse-on-surface text-xs py-1 px-2 rounded opacity-0 group-hover:opacity-100 transition-opacity font-mono-data whitespace-nowrap"
                >
                  ₹6.0L
                </div>
              </div>
              <div class="w-12 bg-outline-variant rounded-t-sm h-[25%] relative group">
                <div
                  class="absolute -top-8 left-1/2 -translate-x-1/2 bg-inverse-surface text-inverse-on-surface text-xs py-1 px-2 rounded opacity-0 group-hover:opacity-100 transition-opacity font-mono-data whitespace-nowrap"
                >
                  ₹3.8L
                </div>
              </div>
            </div>
          </div>
          <div
            class="flex justify-around pl-10 pt-2 font-metadata text-metadata text-text-secondary"
          >
            <span class="w-12 text-center truncate">IT Eqpt</span>
            <span class="w-12 text-center truncate">Furniture</span>
            <span class="w-12 text-center truncate">Vehicles</span>
            <span class="w-12 text-center truncate">Plant</span>
          </div>
        </div>
        <!-- Line Chart: Monthly Spend Trend -->
        <div
          class="bg-surface border border-border-default rounded-xl p-card-padding shadow-[0_4px_16px_-4px_rgba(0,0,0,0.05)] flex flex-col h-[360px]"
        >
          <div class="flex items-center justify-between mb-6">
            <h2 class="font-h2 text-h2 text-text-primary">Monthly Spend Trend</h2>
            <div class="flex gap-2">
              <span class="flex items-center gap-1 font-metadata text-metadata text-text-secondary"
                ><span class="w-2 h-2 rounded-full bg-[#1C1917]"></span> Actual</span
              >
              <span class="flex items-center gap-1 font-metadata text-metadata text-text-secondary"
                ><span class="w-2 h-2 rounded-full bg-border-default"></span> Budget</span
              >
            </div>
          </div>
          <div class="flex-1 relative border-b border-l border-border-default pb-4 pl-4">
            <div
              class="absolute -left-8 top-0 bottom-4 flex flex-col justify-between text-xs text-text-secondary font-mono-data items-end w-6"
            >
              <span>4L</span>
              <span>2L</span>
              <span>0</span>
            </div>
            <div
              class="absolute inset-0 pl-4 pb-4 flex flex-col justify-between pointer-events-none"
            >
              <div class="w-full border-t border-border-default border-dashed opacity-50"></div>
              <div class="w-full border-t border-border-default border-dashed opacity-50"></div>
              <div class="w-full border-t border-border-default border-dashed opacity-50"></div>
            </div>
            <div class="absolute inset-0 pl-4 pb-4 overflow-hidden">
              <div
                class="absolute bottom-0 left-4 right-0 h-[60%] bg-surface-subtle border-t-2 border-border-default border-dashed"
              ></div>
              <svg class="w-full h-full" preserveAspectRatio="none" viewBox="0 0 100 100">
                <path
                  d="M0,80 Q20,70 40,50 T80,30 T100,20"
                  fill="none"
                  stroke="#1C1917"
                  stroke-width="2"
                  vector-effect="non-scaling-stroke"
                ></path>
                <circle cx="40" cy="50" fill="#1C1917" r="3"></circle>
                <circle cx="80" cy="30" fill="#1C1917" r="3"></circle>
              </svg>
            </div>
          </div>
          <div
            class="flex justify-between pl-4 pt-2 font-metadata text-metadata text-text-secondary px-2"
          >
            <span>Apr</span>
            <span>Jun</span>
            <span>Aug</span>
            <span>Oct</span>
            <span>Dec</span>
            <span>Feb</span>
          </div>
        </div>
      </div>
      <!-- Financial Ledger Table -->
      <div
        class="bg-surface border border-border-default rounded-xl shadow-[0_4px_16px_-4px_rgba(0,0,0,0.05)] overflow-hidden"
      >
        <div
          class="px-card-padding py-4 border-b border-border-default bg-surface-subtle flex items-center justify-between"
        >
          <h3 class="font-h3 text-h3 text-text-primary">Recent Ledger Entries</h3>
          <div class="relative">
            <span
              class="material-symbols-outlined absolute left-2.5 top-1/2 -translate-y-1/2 text-text-secondary text-[16px]"
              >search</span
            >
            <input
              class="bg-surface border border-border-default rounded py-1 pl-8 pr-3 font-body text-body text-text-primary focus:outline-none focus:border-primary text-sm w-48"
              placeholder="Search entries..."
              type="text"
            />
          </div>
        </div>
        <div class="overflow-x-auto">
          <table class="w-full text-left border-collapse whitespace-nowrap">
            <thead>
              <tr>
                <th
                  class="font-table-header text-table-header text-text-secondary px-card-padding py-3 border-b border-border-default w-16"
                >
                  ID
                </th>
                <th
                  class="font-table-header text-table-header text-text-secondary px-card-padding py-3 border-b border-border-default"
                >
                  Date
                </th>
                <th
                  class="font-table-header text-table-header text-text-secondary px-card-padding py-3 border-b border-border-default"
                >
                  Category
                </th>
                <th
                  class="font-table-header text-table-header text-text-secondary px-card-padding py-3 border-b border-border-default"
                >
                  Description
                </th>
                <th
                  class="font-table-header text-table-header text-text-secondary px-card-padding py-3 border-b border-border-default text-right"
                >
                  Debit (₹)
                </th>
                <th
                  class="font-table-header text-table-header text-text-secondary px-card-padding py-3 border-b border-border-default text-right"
                >
                  Credit (₹)
                </th>
                <th
                  class="font-table-header text-table-header text-text-secondary px-card-padding py-3 border-b border-border-default text-right"
                >
                  Balance (₹)
                </th>
              </tr>
            </thead>
            <tbody class="font-body text-body text-text-primary divide-y divide-border-default">
              <tr class="bg-surface hover:bg-accent-subtle/30 transition-colors group">
                <td class="px-card-padding py-3 font-mono-data text-mono-data text-text-secondary">
                  #LD-4092
                </td>
                <td class="px-card-padding py-3 font-mono-data text-mono-data">2023-10-24</td>
                <td class="px-card-padding py-3">
                  <span
                    class="inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium bg-secondary-container text-on-secondary-container"
                    >IT Equipment</span
                  >
                </td>
                <td class="px-card-padding py-3">Server Racks Procurement</td>
                <td class="px-card-padding py-3 font-mono-data text-mono-data text-right">
                  4,50,000
                </td>
                <td
                  class="px-card-padding py-3 font-mono-data text-mono-data text-right text-text-secondary"
                >
                  -
                </td>
                <td
                  class="px-card-padding py-3 font-mono-data text-mono-data text-right font-medium"
                >
                  82,30,000
                </td>
              </tr>
              <tr class="bg-surface-subtle hover:bg-accent-subtle/30 transition-colors group">
                <td
                  class="px-card-padding py-3 font-mono-data text-mono-data text-text-secondary border-l-[3px] border-amber-500"
                >
                  #LD-4091
                </td>
                <td class="px-card-padding py-3 font-mono-data text-mono-data">2023-10-21</td>
                <td class="px-card-padding py-3">
                  <span
                    class="inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium bg-metric-sage text-status-in-stock"
                    >Depreciation</span
                  >
                </td>
                <td class="px-card-padding py-3">Monthly Accrual - Vehicles</td>
                <td
                  class="px-card-padding py-3 font-mono-data text-mono-data text-right text-text-secondary"
                >
                  -
                </td>
                <td class="px-card-padding py-3 font-mono-data text-mono-data text-right">
                  1,20,000
                </td>
                <td
                  class="px-card-padding py-3 font-mono-data text-mono-data text-right font-medium"
                >
                  77,80,000
                </td>
              </tr>
              <tr class="bg-surface hover:bg-accent-subtle/30 transition-colors group">
                <td class="px-card-padding py-3 font-mono-data text-mono-data text-text-secondary">
                  #LD-4090
                </td>
                <td class="px-card-padding py-3 font-mono-data text-mono-data">2023-10-18</td>
                <td class="px-card-padding py-3">
                  <span
                    class="inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium bg-secondary-container text-on-secondary-container"
                    >Furniture</span
                  >
                </td>
                <td class="px-card-padding py-3">Ergonomic Chairs Batch B</td>
                <td class="px-card-padding py-3 font-mono-data text-mono-data text-right">
                  2,15,000
                </td>
                <td
                  class="px-card-padding py-3 font-mono-data text-mono-data text-right text-text-secondary"
                >
                  -
                </td>
                <td
                  class="px-card-padding py-3 font-mono-data text-mono-data text-right font-medium"
                >
                  79,00,000
                </td>
              </tr>
              <tr class="bg-surface-subtle hover:bg-accent-subtle/30 transition-colors group">
                <td class="px-card-padding py-3 font-mono-data text-mono-data text-text-secondary">
                  #LD-4089
                </td>
                <td class="px-card-padding py-3 font-mono-data text-mono-data">2023-10-15</td>
                <td class="px-card-padding py-3">
                  <span
                    class="inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium bg-metric-amber text-primary-container"
                    >Maintenance</span
                  >
                </td>
                <td class="px-card-padding py-3">HVAC Quarterly Service</td>
                <td class="px-card-padding py-3 font-mono-data text-mono-data text-right">
                  85,000
                </td>
                <td
                  class="px-card-padding py-3 font-mono-data text-mono-data text-right text-text-secondary"
                >
                  -
                </td>
                <td
                  class="px-card-padding py-3 font-mono-data text-mono-data text-right font-medium"
                >
                  76,85,000
                </td>
              </tr>
              <tr class="bg-surface hover:bg-accent-subtle/30 transition-colors group">
                <td class="px-card-padding py-3 font-mono-data text-mono-data text-text-secondary">
                  #LD-4088
                </td>
                <td class="px-card-padding py-3 font-mono-data text-mono-data">2023-10-01</td>
                <td class="px-card-padding py-3">
                  <span
                    class="inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium bg-metric-sage text-status-in-stock"
                    >Disposal</span
                  >
                </td>
                <td class="px-card-padding py-3">Old Desktops Salvage Value</td>
                <td
                  class="px-card-padding py-3 font-mono-data text-mono-data text-right text-text-secondary"
                >
                  -
                </td>
                <td class="px-card-padding py-3 font-mono-data text-mono-data text-right">
                  45,000
                </td>
                <td
                  class="px-card-padding py-3 font-mono-data text-mono-data text-right font-medium"
                >
                  76,00,000
                </td>
              </tr>
            </tbody>
          </table>
        </div>
        <div
          class="px-card-padding py-3 border-t border-border-default flex items-center justify-between bg-surface"
        >
          <span class="font-metadata text-metadata text-text-secondary"
            >Showing 1-5 of 142 entries</span
          >
          <div class="flex gap-1">
            <button
              class="px-3 py-1 border border-border-default rounded bg-surface hover:bg-surface-subtle font-body text-body text-text-secondary disabled:opacity-50"
            >
              Prev
            </button>
            <button
              class="px-3 py-1 border border-primary rounded bg-primary/5 text-primary font-body text-body"
            >
              1
            </button>
            <button
              class="px-3 py-1 border border-border-default rounded bg-surface hover:bg-surface-subtle font-body text-body text-text-secondary"
            >
              2
            </button>
            <button
              class="px-3 py-1 border border-border-default rounded bg-surface hover:bg-surface-subtle font-body text-body text-text-secondary"
            >
              Next
            </button>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
// Financial Reports & Analytics Dashboard — KPI cards, bar/line charts, and ledger table
</script>
