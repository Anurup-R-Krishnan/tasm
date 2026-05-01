<template>
  <main class="flex-1 flex min-h-screen relative">
    <!-- Left Panel: Stockroom List -->
    <aside
      class="w-[260px] bg-surface-subtle border-r border-border-default flex flex-col h-full flex-shrink-0"
    >
      <div class="p-4 border-b border-border-default flex items-center justify-between bg-surface">
        <h2 class="font-h3 text-text-primary">Locations</h2>
        <button class="text-text-secondary hover:text-primary transition-colors">
          <span class="material-symbols-outlined text-[20px]">filter_list</span>
        </button>
      </div>
      <div class="flex-1 overflow-y-auto">
        <!-- Selected Item -->
        <div class="p-4 border-b border-border-default bg-accent-subtle cursor-pointer relative">
          <div class="absolute left-0 top-0 bottom-0 w-[3px] bg-primary"></div>
          <div class="flex justify-between items-start mb-1">
            <span class="font-h2 text-text-primary">Main IT Store</span>
            <span
              class="bg-surface border border-border-default px-1.5 py-0.5 rounded text-[11px] font-medium text-text-secondary shadow-sm"
              >1,245</span
            >
          </div>
          <div class="font-metadata text-text-secondary flex items-center gap-1">
            <span class="material-symbols-outlined text-[14px]">business</span>
            Bhavani Building
          </div>
        </div>
        <!-- Unselected Items -->
        <div
          class="p-4 border-b border-border-default bg-surface hover:bg-surface-dim cursor-pointer transition-colors group"
        >
          <div class="flex justify-between items-start mb-1">
            <span class="font-h2 text-text-primary group-hover:text-primary transition-colors"
              >Furniture Annex</span
            >
            <span
              class="bg-surface-subtle border border-border-default px-1.5 py-0.5 rounded text-[11px] font-medium text-text-secondary"
              >482</span
            >
          </div>
          <div class="font-metadata text-text-secondary flex items-center gap-1">
            <span class="material-symbols-outlined text-[14px]">business</span>
            Thejaswini Building
          </div>
        </div>
        <div
          class="p-4 border-b border-border-default bg-surface hover:bg-surface-dim cursor-pointer transition-colors group"
        >
          <div class="flex justify-between items-start mb-1">
            <span class="font-h2 text-text-primary group-hover:text-primary transition-colors"
              >Network Closet A</span
            >
            <span
              class="bg-surface-subtle border border-border-default px-1.5 py-0.5 rounded text-[11px] font-medium text-text-secondary"
              >84</span
            >
          </div>
          <div class="font-metadata text-text-secondary flex items-center gap-1">
            <span class="material-symbols-outlined text-[14px]">business</span>
            Nila Building
          </div>
        </div>
        <div
          class="p-4 border-b border-border-default bg-surface hover:bg-surface-dim cursor-pointer transition-colors group"
        >
          <div class="flex justify-between items-start mb-1">
            <span class="font-h2 text-text-primary group-hover:text-primary transition-colors"
              >Executive Store</span
            >
            <span
              class="bg-surface-subtle border border-border-default px-1.5 py-0.5 rounded text-[11px] font-medium text-text-secondary"
              >156</span
            >
          </div>
          <div class="font-metadata text-text-secondary flex items-center gap-1">
            <span class="material-symbols-outlined text-[14px]">business</span>
            Park Centre
          </div>
        </div>
      </div>
    </aside>
    <!-- Right Panel: Selected Stockroom View -->
    <section class="flex-1 bg-canvas flex flex-col h-full overflow-hidden">
      <div class="flex-1 overflow-y-auto p-page-margin">
        <!-- View Header -->
        <div class="flex justify-between items-end mb-section-gap">
          <div>
            <div class="flex items-center gap-2 text-text-secondary font-metadata mb-1">
              <span>Locations</span>
              <span class="material-symbols-outlined text-[14px]">chevron_right</span>
              <span>Bhavani Building</span>
            </div>
            <h1 class="font-h1 text-text-primary mb-1">Main IT Store</h1>
            <p class="font-metadata text-text-secondary flex items-center gap-1">
              <span class="material-symbols-outlined text-[16px]">location_on</span>
              Ground Floor, East Wing, Bhavani Building
            </p>
          </div>
          <div class="flex gap-inline">
            <button
              class="bg-surface border border-border-default text-text-primary hover:bg-surface-dim px-4 py-2 rounded text-body font-medium transition-colors shadow-sm flex items-center gap-2"
            >
              <span class="material-symbols-outlined text-[18px]">qr_code_scanner</span>
              Scan Assets
            </button>
            <button
              class="bg-surface border border-border-default text-text-primary hover:bg-surface-dim px-4 py-2 rounded text-body font-medium transition-colors shadow-sm flex items-center gap-2"
            >
              <span class="material-symbols-outlined text-[18px]">download</span>
              Export List
            </button>
          </div>
        </div>
        <!-- Amber Metric Cards -->
        <div class="grid grid-cols-3 gap-stack mb-section-gap">
          <div
            class="bg-metric-amber border border-[#FDE68A] rounded-xl p-card-padding shadow-sm relative overflow-hidden group hover:-translate-y-0.5 transition-transform duration-200 cursor-pointer"
          >
            <div
              class="absolute right-0 top-0 p-4 opacity-20 group-hover:opacity-30 transition-opacity"
            >
              <span class="material-symbols-outlined text-[64px] text-primary-container"
                >devices</span
              >
            </div>
            <div class="relative z-10">
              <div
                class="font-metadata text-primary-container uppercase tracking-wider font-semibold mb-2"
              >
                Total Assets
              </div>
              <div class="font-kpi-number text-text-primary">1,245</div>
              <div class="font-metadata text-text-secondary mt-1 flex items-center gap-1">
                <span class="material-symbols-outlined text-[14px] text-status-in-stock"
                  >arrow_upward</span
                >
                <span>+12 this month</span>
              </div>
            </div>
          </div>
          <div
            class="bg-metric-amber border border-[#FDE68A] rounded-xl p-card-padding shadow-sm relative overflow-hidden group hover:-translate-y-0.5 transition-transform duration-200 cursor-pointer"
          >
            <div
              class="absolute right-0 top-0 p-4 opacity-20 group-hover:opacity-30 transition-opacity"
            >
              <span class="material-symbols-outlined text-[64px] text-primary-container"
                >account_balance_wallet</span
              >
            </div>
            <div class="relative z-10">
              <div
                class="font-metadata text-primary-container uppercase tracking-wider font-semibold mb-2"
              >
                Total Value (₹)
              </div>
              <div class="font-kpi-number text-text-primary">45,20,000</div>
              <div class="font-metadata text-text-secondary mt-1">
                Book value minus depreciation
              </div>
            </div>
          </div>
          <div
            class="bg-surface border border-border-default rounded-xl p-card-padding shadow-sm relative overflow-hidden group hover:-translate-y-0.5 transition-transform duration-200 cursor-pointer"
          >
            <div class="flex items-start justify-between">
              <div>
                <div
                  class="font-metadata text-text-secondary uppercase tracking-wider font-semibold mb-2"
                >
                  Responsible Person
                </div>
                <div class="font-h2 text-text-primary mb-1">Jane Doe</div>
                <div class="font-metadata text-text-secondary">Senior IT Manager</div>
              </div>
              <div class="h-12 w-12 rounded-full overflow-hidden border border-border-default">
                <img
                  alt="Jane Doe Profile"
                  class="w-full h-full object-cover"
                  src="https://lh3.googleusercontent.com/aida-public/AB6AXuDlpAhSdcZX_w4fNlqduHwnt03p2pgQmcN95nCdjx5GX73rjbOROniHGIjK9XYmEF379QStV4qTap70lbUf5k1BV1ypiAYqyaFS0PEzdyBB1qOubfoKCXfyScyAzpVAyu_oUQc935HtmN2vVrCSbv1cZEEHvzABDu10YMco_f625jo2Z5ITs97tREGGy4pd-x-l8nY2HvTijWm0XHVQFJeU10B57ZLeMUpfOzmkZqRHZh6LfGZFV2ptNJBX7GBWw1FvV_tCPWZc00u2"
                />
              </div>
            </div>
            <div
              class="mt-4 pt-4 border-t border-border-default flex items-center gap-2 text-primary font-medium text-sm"
            >
              <span class="material-symbols-outlined text-[16px]">mail</span>
              Contact Manager
            </div>
          </div>
        </div>
        <!-- High-Density Table Section -->
        <div class="bg-surface border border-border-default rounded-xl shadow-sm flex flex-col">
          <!-- Table Toolbar -->
          <div
            class="p-4 border-b border-border-default flex justify-between items-center bg-surface-subtle rounded-t-xl"
          >
            <div class="flex items-center gap-4">
              <div class="relative">
                <span
                  class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-text-secondary text-[18px]"
                  >search</span
                >
                <input
                  class="w-64 bg-surface border border-border-default rounded text-body font-body py-1.5 pl-9 pr-3 focus:ring-1 focus:ring-primary focus:border-primary outline-none text-sm"
                  placeholder="Filter this stockroom..."
                  type="text"
                />
              </div>
              <select
                class="bg-surface border border-border-default rounded text-body font-body py-1.5 px-3 text-sm focus:ring-1 focus:ring-primary outline-none text-text-primary"
              >
                <option>All Categories</option>
                <option>Laptops</option>
                <option>Monitors</option>
                <option>Peripherals</option>
              </select>
            </div>
            <div class="flex items-center gap-2">
              <span class="font-metadata text-text-secondary mr-2">Showing 1-6 of 1,245</span>
              <button
                class="p-1 border border-border-default rounded bg-surface hover:bg-surface-dim text-text-secondary disabled:opacity-50"
              >
                <span class="material-symbols-outlined text-[18px]">chevron_left</span>
              </button>
              <button
                class="p-1 border border-border-default rounded bg-surface hover:bg-surface-dim text-text-secondary"
              >
                <span class="material-symbols-outlined text-[18px]">chevron_right</span>
              </button>
            </div>
          </div>
          <!-- Table -->
          <div class="overflow-x-auto">
            <table class="w-full text-left border-collapse whitespace-nowrap">
              <thead
                class="bg-surface font-table-header text-text-secondary border-b border-border-default"
              >
                <tr>
                  <th class="py-2.5 px-4 font-medium border-r border-border-default/50">Tag ID</th>
                  <th class="py-2.5 px-4 font-medium border-r border-border-default/50">
                    Asset Name
                  </th>
                  <th class="py-2.5 px-4 font-medium border-r border-border-default/50">
                    Category
                  </th>
                  <th class="py-2.5 px-4 font-medium border-r border-border-default/50">
                    Condition
                  </th>
                  <th class="py-2.5 px-4 font-medium border-r border-border-default/50">
                    Date Stocked
                  </th>
                  <th class="py-2.5 px-4 w-12 text-center">Actions</th>
                </tr>
              </thead>
              <tbody
                class="font-body text-text-primary text-[13px] divide-y divide-border-default/50"
              >
                <tr
                  class="bg-surface hover:bg-accent-subtle transition-colors group cursor-pointer"
                >
                  <td
                    class="py-2 px-4 border-r border-border-default/50 font-mono-data text-primary"
                  >
                    TPK-IT-9021
                  </td>
                  <td class="py-2 px-4 border-r border-border-default/50 font-medium">
                    Dell Latitude 7420
                  </td>
                  <td class="py-2 px-4 border-r border-border-default/50 text-text-secondary">
                    Laptop
                  </td>
                  <td class="py-2 px-4 border-r border-border-default/50">
                    <span
                      class="inline-flex items-center px-2 py-0.5 rounded-full text-[11px] font-medium bg-status-in-stock text-white"
                      >Good</span
                    >
                  </td>
                  <td
                    class="py-2 px-4 border-r border-border-default/50 text-text-secondary font-mono-data"
                  >
                    2023-10-12
                  </td>
                  <td class="py-2 px-4 text-center">
                    <button
                      class="text-text-secondary hover:text-primary p-1 rounded hover:bg-surface"
                    >
                      <span class="material-symbols-outlined text-[18px]">more_vert</span>
                    </button>
                  </td>
                </tr>
                <tr
                  class="bg-surface-subtle hover:bg-accent-subtle transition-colors group cursor-pointer"
                >
                  <td
                    class="py-2 px-4 border-r border-border-default/50 font-mono-data text-primary"
                  >
                    TPK-IT-9022
                  </td>
                  <td class="py-2 px-4 border-r border-border-default/50 font-medium">
                    Dell Latitude 7420
                  </td>
                  <td class="py-2 px-4 border-r border-border-default/50 text-text-secondary">
                    Laptop
                  </td>
                  <td class="py-2 px-4 border-r border-border-default/50">
                    <span
                      class="inline-flex items-center px-2 py-0.5 rounded-full text-[11px] font-medium bg-status-in-stock text-white"
                      >Good</span
                    >
                  </td>
                  <td
                    class="py-2 px-4 border-r border-border-default/50 text-text-secondary font-mono-data"
                  >
                    2023-10-12
                  </td>
                  <td class="py-2 px-4 text-center">
                    <button
                      class="text-text-secondary hover:text-primary p-1 rounded hover:bg-surface"
                    >
                      <span class="material-symbols-outlined text-[18px]">more_vert</span>
                    </button>
                  </td>
                </tr>
                <tr
                  class="bg-surface hover:bg-accent-subtle transition-colors group cursor-pointer"
                >
                  <td
                    class="py-2 px-4 border-r border-border-default/50 font-mono-data text-primary"
                  >
                    TPK-MN-1104
                  </td>
                  <td
                    class="py-2 px-4 border-r border-border-default/50 font-medium text-text-primary"
                  >
                    LG 27" UltraSharp
                  </td>
                  <td class="py-2 px-4 border-r border-border-default/50 text-text-secondary">
                    Monitor
                  </td>
                  <td class="py-2 px-4 border-r border-border-default/50">
                    <span
                      class="inline-flex items-center px-2 py-0.5 rounded-full text-[11px] font-medium bg-[#B45309] text-white"
                      >Needs Repair</span
                    >
                  </td>
                  <td
                    class="py-2 px-4 border-r border-border-default/50 text-text-secondary font-mono-data"
                  >
                    2022-04-05
                  </td>
                  <td class="py-2 px-4 text-center">
                    <button
                      class="text-text-secondary hover:text-primary p-1 rounded hover:bg-surface"
                    >
                      <span class="material-symbols-outlined text-[18px]">more_vert</span>
                    </button>
                  </td>
                </tr>
                <tr
                  class="bg-surface-subtle hover:bg-accent-subtle transition-colors group cursor-pointer"
                >
                  <td
                    class="py-2 px-4 border-r border-border-default/50 font-mono-data text-primary"
                  >
                    TPK-AC-0089
                  </td>
                  <td class="py-2 px-4 border-r border-border-default/50 font-medium">
                    Logitech MX Master 3
                  </td>
                  <td class="py-2 px-4 border-r border-border-default/50 text-text-secondary">
                    Peripheral
                  </td>
                  <td class="py-2 px-4 border-r border-border-default/50">
                    <span
                      class="inline-flex items-center px-2 py-0.5 rounded-full text-[11px] font-medium bg-status-in-stock text-white"
                      >New</span
                    >
                  </td>
                  <td
                    class="py-2 px-4 border-r border-border-default/50 text-text-secondary font-mono-data"
                  >
                    2024-01-20
                  </td>
                  <td class="py-2 px-4 text-center">
                    <button
                      class="text-text-secondary hover:text-primary p-1 rounded hover:bg-surface"
                    >
                      <span class="material-symbols-outlined text-[18px]">more_vert</span>
                    </button>
                  </td>
                </tr>
                <tr
                  class="bg-surface hover:bg-accent-subtle transition-colors group cursor-pointer"
                >
                  <td
                    class="py-2 px-4 border-r border-border-default/50 font-mono-data text-primary"
                  >
                    TPK-IT-8840
                  </td>
                  <td
                    class="py-2 px-4 border-r border-border-default/50 font-medium text-status-critical"
                  >
                    MacBook Pro 16"
                  </td>
                  <td class="py-2 px-4 border-r border-border-default/50 text-text-secondary">
                    Laptop
                  </td>
                  <td class="py-2 px-4 border-r border-border-default/50">
                    <span
                      class="inline-flex items-center px-2 py-0.5 rounded-full text-[11px] font-medium bg-status-critical text-white"
                      >Damaged</span
                    >
                  </td>
                  <td
                    class="py-2 px-4 border-r border-border-default/50 text-text-secondary font-mono-data"
                  >
                    2021-11-15
                  </td>
                  <td class="py-2 px-4 text-center">
                    <button
                      class="text-text-secondary hover:text-primary p-1 rounded hover:bg-surface"
                    >
                      <span class="material-symbols-outlined text-[18px]">more_vert</span>
                    </button>
                  </td>
                </tr>
                <tr
                  class="bg-surface-subtle hover:bg-accent-subtle transition-colors group cursor-pointer"
                >
                  <td
                    class="py-2 px-4 border-r border-border-default/50 font-mono-data text-primary"
                  >
                    TPK-AC-0102
                  </td>
                  <td class="py-2 px-4 border-r border-border-default/50 font-medium">
                    Keychron K2 Keyboard
                  </td>
                  <td class="py-2 px-4 border-r border-border-default/50 text-text-secondary">
                    Peripheral
                  </td>
                  <td class="py-2 px-4 border-r border-border-default/50">
                    <span
                      class="inline-flex items-center px-2 py-0.5 rounded-full text-[11px] font-medium bg-status-in-stock text-white"
                      >Good</span
                    >
                  </td>
                  <td
                    class="py-2 px-4 border-r border-border-default/50 text-text-secondary font-mono-data"
                  >
                    2023-08-01
                  </td>
                  <td class="py-2 px-4 text-center">
                    <button
                      class="text-text-secondary hover:text-primary p-1 rounded hover:bg-surface"
                    >
                      <span class="material-symbols-outlined text-[18px]">more_vert</span>
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </section>
  </main>
</template>

<script setup lang="ts">
// Stockroom Inventory Management — Two-panel layout with location list and asset table
</script>
