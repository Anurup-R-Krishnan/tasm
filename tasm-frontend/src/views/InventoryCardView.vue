<template>
  <main class="flex-1 flex flex-col min-h-screen relative">
    <div class="p-page-margin max-w-[1400px] w-full mx-auto space-y-section-gap pb-24">
      <!-- Page Header & Actions -->
      <div class="flex flex-col md:flex-row md:items-end justify-between gap-stack">
        <div>
          <h1 class="font-h1 text-h1 text-text-primary mb-1">Inventory Management</h1>
          <p class="font-body text-body text-text-secondary">
            Manage and track assets across all campus locations.
          </p>
        </div>
        <div class="flex items-center gap-inline">
          <button
            class="bg-surface border border-border-default text-text-primary px-4 py-2 rounded-lg font-body text-body flex items-center gap-2 hover:-translate-y-[2px] transition-transform shadow-sm"
          >
            <span class="material-symbols-outlined text-[18px]">filter_list</span>
            Filters
          </button>
          <button
            class="bg-surface border border-border-default text-text-primary px-4 py-2 rounded-lg font-body text-body flex items-center gap-2 hover:-translate-y-[2px] transition-transform shadow-sm"
          >
            <span class="material-symbols-outlined text-[18px]">download</span>
            Export
          </button>
        </div>
      </div>
      <!-- Status Summary Pills -->
      <div class="flex flex-wrap gap-inline">
        <button
          class="bg-metric-amber border border-amber-200 px-4 py-2 rounded-full flex items-center gap-2 text-sm font-medium text-amber-900 hover:bg-amber-200 transition-colors"
        >
          <span class="w-2 h-2 rounded-full bg-amber-600"></span>
          All Assets
          <span class="ml-1 px-1.5 py-0.5 bg-amber-900/10 rounded-full text-xs">1,248</span>
        </button>
        <button
          class="bg-surface border border-border-default px-4 py-2 rounded-full flex items-center gap-2 text-sm font-medium text-text-primary hover:bg-surface-subtle transition-colors shadow-sm"
        >
          <span class="w-2 h-2 rounded-full bg-status-in-stock"></span>
          In Stock
          <span class="ml-1 px-1.5 py-0.5 bg-stone-100 rounded-full text-xs">842</span>
        </button>
        <button
          class="bg-surface border border-border-default px-4 py-2 rounded-full flex items-center gap-2 text-sm font-medium text-text-primary hover:bg-surface-subtle transition-colors shadow-sm"
        >
          <span class="w-2 h-2 rounded-full bg-status-checked-out"></span>
          Checked Out
          <span class="ml-1 px-1.5 py-0.5 bg-stone-100 rounded-full text-xs">315</span>
        </button>
        <button
          class="bg-surface border border-border-default px-4 py-2 rounded-full flex items-center gap-2 text-sm font-medium text-text-primary hover:bg-surface-subtle transition-colors shadow-sm"
        >
          <span class="w-2 h-2 rounded-full bg-status-critical"></span>
          Under Repair
          <span class="ml-1 px-1.5 py-0.5 bg-stone-100 rounded-full text-xs">91</span>
        </button>
      </div>
      <!-- Hybrid Layout: Sidebar + Grid -->
      <div class="flex flex-col lg:flex-row gap-section-gap">
        <!-- Category Sidebar -->
        <aside class="w-full lg:w-[240px] shrink-0 space-y-6">
          <!-- Categories -->
          <div class="bg-surface border border-border-default rounded-[16px] p-4 shadow-sm">
            <h3 class="font-h3 text-h3 text-text-primary mb-3 px-2">Categories</h3>
            <ul class="space-y-1">
              <li>
                <button
                  class="w-full text-left px-2 py-1.5 rounded-md bg-amber-50 text-primary-container font-medium text-sm flex justify-between items-center"
                >
                  <span>IT Equipment</span>
                  <span class="text-xs bg-amber-100 px-1.5 rounded-full">540</span>
                </button>
              </li>
              <li>
                <button
                  class="w-full text-left px-2 py-1.5 rounded-md text-text-secondary hover:bg-surface-subtle hover:text-text-primary font-medium text-sm flex justify-between items-center transition-colors"
                >
                  <span>Furniture</span>
                  <span class="text-xs bg-stone-100 px-1.5 rounded-full">420</span>
                </button>
              </li>
              <li>
                <button
                  class="w-full text-left px-2 py-1.5 rounded-md text-text-secondary hover:bg-surface-subtle hover:text-text-primary font-medium text-sm flex justify-between items-center transition-colors"
                >
                  <span>Vehicles</span>
                  <span class="text-xs bg-stone-100 px-1.5 rounded-full">45</span>
                </button>
              </li>
              <li>
                <button
                  class="w-full text-left px-2 py-1.5 rounded-md text-text-secondary hover:bg-surface-subtle hover:text-text-primary font-medium text-sm flex justify-between items-center transition-colors"
                >
                  <span>HVAC Systems</span>
                  <span class="text-xs bg-stone-100 px-1.5 rounded-full">120</span>
                </button>
              </li>
              <li>
                <button
                  class="w-full text-left px-2 py-1.5 rounded-md text-text-secondary hover:bg-surface-subtle hover:text-text-primary font-medium text-sm flex justify-between items-center transition-colors"
                >
                  <span>Security Gear</span>
                  <span class="text-xs bg-stone-100 px-1.5 rounded-full">123</span>
                </button>
              </li>
            </ul>
          </div>
          <!-- Locations -->
          <div class="bg-surface border border-border-default rounded-[16px] p-4 shadow-sm">
            <h3 class="font-h3 text-h3 text-text-primary mb-3 px-2">Locations</h3>
            <ul class="space-y-1">
              <li>
                <label class="flex items-center gap-2 px-2 py-1.5 cursor-pointer">
                  <input
                    checked
                    class="rounded border-stone-300 text-primary-container focus:ring-primary-container"
                    type="checkbox"
                  />
                  <span class="text-sm text-text-primary">Bhavani Building</span>
                </label>
              </li>
              <li>
                <label class="flex items-center gap-2 px-2 py-1.5 cursor-pointer">
                  <input
                    class="rounded border-stone-300 text-primary-container focus:ring-primary-container"
                    type="checkbox"
                  />
                  <span class="text-sm text-text-secondary">Thejaswini</span>
                </label>
              </li>
              <li>
                <label class="flex items-center gap-2 px-2 py-1.5 cursor-pointer">
                  <input
                    class="rounded border-stone-300 text-primary-container focus:ring-primary-container"
                    type="checkbox"
                  />
                  <span class="text-sm text-text-secondary">Nila</span>
                </label>
              </li>
              <li>
                <label class="flex items-center gap-2 px-2 py-1.5 cursor-pointer">
                  <input
                    class="rounded border-stone-300 text-primary-container focus:ring-primary-container"
                    type="checkbox"
                  />
                  <span class="text-sm text-text-secondary">Gayathri</span>
                </label>
              </li>
            </ul>
          </div>
        </aside>
        <!-- Asset Grid -->
        <div class="flex-1">
          <div class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-stack">
            <!-- Asset Card 1 -->
            <div
              class="bg-surface border border-border-default rounded-[16px] overflow-hidden shadow-sm hover:shadow-md hover:-translate-y-[2px] transition-all cursor-pointer group flex flex-col"
            >
              <div class="h-40 bg-stone-100 relative overflow-hidden">
                <img
                  alt="MacBook Pro"
                  class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500"
                  src="https://lh3.googleusercontent.com/aida-public/AB6AXuBTNo2hWSWLbvgJaP8tQ-NqXWeugO_-AYHMMKjmQKoTlRXohk96QV2ZcEyY3Ydc3XnLM5g-lrUAyGl7K577xvxiBep-KQPVgheeEMpbMI5wiAGuLeNKnYvbyRvZHAf9mMV_uG8soLr4rsqswxRr-xGLt-nR0Eschbo1nHeHczVqqVUG9uhUmtIi63bLdfIruu5nVeFAZQ_ZsOJ7IeHBDBNAG1wLrujWfj2gE2mma18vJcDczJu98nEavsVS0ZCcboi3WsRfLbq0GXAr"
                />
                <div class="absolute top-3 right-3">
                  <span
                    class="bg-status-in-stock text-white text-xs font-medium px-2 py-1 rounded-md shadow-sm"
                    >In Stock</span
                  >
                </div>
              </div>
              <div class="p-4 flex flex-col flex-1">
                <div class="flex justify-between items-start mb-2">
                  <h3 class="font-h3 text-h3 text-text-primary line-clamp-1">
                    MacBook Pro 16" M2 Max
                  </h3>
                </div>
                <div class="flex items-center gap-2 mb-3">
                  <span
                    class="font-mono-data text-mono-data text-stone-500 bg-stone-100 px-2 py-1 rounded"
                    >AST-IT-4092</span
                  >
                </div>
                <div class="space-y-2 mt-auto pt-3 border-t border-border-default">
                  <div class="flex justify-between items-center">
                    <span
                      class="font-metadata text-metadata text-text-secondary flex items-center gap-1"
                    >
                      <span class="material-symbols-outlined text-[14px]">location_on</span>
                      Bhavani, Floor 3
                    </span>
                    <span class="font-h3 text-h3 text-text-primary">₹ 2,45,000</span>
                  </div>
                </div>
              </div>
            </div>
            <!-- Asset Card 2 -->
            <div
              class="bg-surface border border-border-default rounded-[16px] overflow-hidden shadow-sm hover:shadow-md hover:-translate-y-[2px] transition-all cursor-pointer group flex flex-col"
            >
              <div class="h-40 bg-stone-100 relative overflow-hidden">
                <img
                  alt="Ergonomic Chair"
                  class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500"
                  src="https://lh3.googleusercontent.com/aida-public/AB6AXuDe1NLE63SBTb2I-UTfZxv_w-6uYobEx-lr0h1jf8A2xuk7tvdaMiiEeDxy2AuIX_9R7n5Qv5qaen-5fYzpnI-c78Vaj3U_1NjabAj_9Le4h6rnZc-tDSSbv-hYaB-_cbvSspQobSnzeHCCVehaX5HRZt_JlOOafmeAf0DJi1tbXHnDEySJiLtu6_dXmBg28aXQ7seRpVWr9yqU83cUjB9hxuQr9nVX28q2EQxarqniH0iQUKne4cDADW6If5J_lNVS5TL04k-GMbZy"
                />
                <div class="absolute top-3 right-3">
                  <span
                    class="bg-status-checked-out text-white text-xs font-medium px-2 py-1 rounded-md shadow-sm"
                    >Checked Out</span
                  >
                </div>
              </div>
              <div class="p-4 flex flex-col flex-1">
                <div class="flex justify-between items-start mb-2">
                  <h3 class="font-h3 text-h3 text-text-primary line-clamp-1">
                    Herman Miller Aeron Chair
                  </h3>
                </div>
                <div class="flex items-center gap-2 mb-3">
                  <span
                    class="font-mono-data text-mono-data text-stone-500 bg-stone-100 px-2 py-1 rounded"
                    >AST-FN-1104</span
                  >
                  <span class="text-xs text-text-secondary flex items-center gap-1">
                    <span class="material-symbols-outlined text-[12px]">person</span>
                    Rahul M.
                  </span>
                </div>
                <div class="space-y-2 mt-auto pt-3 border-t border-border-default">
                  <div class="flex justify-between items-center">
                    <span
                      class="font-metadata text-metadata text-text-secondary flex items-center gap-1"
                    >
                      <span class="material-symbols-outlined text-[14px]">location_on</span>
                      Thejaswini, Floor 5
                    </span>
                    <span class="font-h3 text-h3 text-text-primary">₹ 1,12,500</span>
                  </div>
                </div>
              </div>
            </div>
            <!-- Asset Card 3 (Under Repair) -->
            <div
              class="bg-surface border border-border-default rounded-[16px] overflow-hidden shadow-sm hover:shadow-md hover:-translate-y-[2px] transition-all cursor-pointer group flex flex-col border-l-4 border-l-status-critical"
            >
              <div class="h-40 bg-stone-100 relative overflow-hidden">
                <div class="absolute inset-0 bg-stone-200 flex items-center justify-center">
                  <span class="material-symbols-outlined text-stone-400 text-4xl">print</span>
                </div>
                <div class="absolute top-3 right-3">
                  <span
                    class="bg-status-critical text-white text-xs font-medium px-2 py-1 rounded-md shadow-sm"
                    >Under Repair</span
                  >
                </div>
              </div>
              <div class="p-4 flex flex-col flex-1">
                <div class="flex justify-between items-start mb-2">
                  <h3 class="font-h3 text-h3 text-text-primary line-clamp-1">
                    HP Enterprise LaserJet Pro
                  </h3>
                </div>
                <div class="flex items-center gap-2 mb-3">
                  <span
                    class="font-mono-data text-mono-data text-stone-500 bg-stone-100 px-2 py-1 rounded"
                    >AST-IT-3392</span
                  >
                </div>
                <div class="space-y-2 mt-auto pt-3 border-t border-border-default">
                  <div class="flex justify-between items-center">
                    <span
                      class="font-metadata text-metadata text-text-secondary flex items-center gap-1 text-status-critical"
                    >
                      <span class="material-symbols-outlined text-[14px]">build</span>
                      IT Service Desk
                    </span>
                    <span class="font-h3 text-h3 text-text-primary">₹ 45,000</span>
                  </div>
                </div>
              </div>
            </div>
            <!-- Asset Card 4 -->
            <div
              class="bg-surface border border-border-default rounded-[16px] overflow-hidden shadow-sm hover:shadow-md hover:-translate-y-[2px] transition-all cursor-pointer group flex flex-col"
            >
              <div class="h-40 bg-stone-100 relative overflow-hidden">
                <img
                  alt="Dell Monitor"
                  class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500"
                  src="https://lh3.googleusercontent.com/aida-public/AB6AXuAcF3salIx4WbBhN-IUX5OHt4P3HcwxPorpA3vfCDRT1Q-vFUeVbzzfRkAzBhtLbGht_sTmJSUS3xAcWtKH-S-NrZ4GmaVad5iEEruILtx4Y_vCDpirLAOH88Q0RMWGNPzaSJfYX-E2mVNna8B2fRNagFwIgTTVXI5B69B6U5y7ha9dptjiArkVZWZemjsOsBbA-xQ8NizbIMTBwwUQazaDib7SVg4TMjpGh7Si4DBWM4tcfLGqUhmN9gL2OOLAQiVehX7u5MtCkp3w"
                />
                <div class="absolute top-3 right-3">
                  <span
                    class="bg-status-in-stock text-white text-xs font-medium px-2 py-1 rounded-md shadow-sm"
                    >In Stock</span
                  >
                </div>
              </div>
              <div class="p-4 flex flex-col flex-1">
                <div class="flex justify-between items-start mb-2">
                  <h3 class="font-h3 text-h3 text-text-primary line-clamp-1">
                    Dell UltraSharp 27" 4K Monitor
                  </h3>
                </div>
                <div class="flex items-center gap-2 mb-3">
                  <span
                    class="font-mono-data text-mono-data text-stone-500 bg-stone-100 px-2 py-1 rounded"
                    >AST-IT-5120</span
                  >
                </div>
                <div class="space-y-2 mt-auto pt-3 border-t border-border-default">
                  <div class="flex justify-between items-center">
                    <span
                      class="font-metadata text-metadata text-text-secondary flex items-center gap-1"
                    >
                      <span class="material-symbols-outlined text-[14px]">location_on</span>
                      Nila, Floor 2
                    </span>
                    <span class="font-h3 text-h3 text-text-primary">₹ 42,000</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <!-- Pagination -->
          <div class="mt-8 flex items-center justify-between border-t border-border-default pt-4">
            <span class="text-sm text-text-secondary">Showing 1 to 4 of 540 assets</span>
            <div class="flex gap-2">
              <button
                class="px-3 py-1 border border-border-default rounded bg-surface text-stone-400 cursor-not-allowed text-sm"
              >
                Previous
              </button>
              <button class="px-3 py-1 bg-primary-container text-white rounded text-sm font-medium">
                1
              </button>
              <button
                class="px-3 py-1 border border-border-default rounded bg-surface text-text-primary hover:bg-surface-subtle text-sm"
              >
                2
              </button>
              <button
                class="px-3 py-1 border border-border-default rounded bg-surface text-text-primary hover:bg-surface-subtle text-sm"
              >
                3
              </button>
              <span class="px-2 py-1 text-stone-400">...</span>
              <button
                class="px-3 py-1 border border-border-default rounded bg-surface text-text-primary hover:bg-surface-subtle text-sm"
              >
                Next
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
// Inventory Management - Card View
</script>
