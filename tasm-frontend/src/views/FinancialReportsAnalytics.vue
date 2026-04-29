<template>
<main class="mt-[60px] p-page-margin max-w-[1400px]">
 <div class="mb-section-gap flex justify-between items-end">
  <div>
   <h1 class="font-h1 text-h1 text-text-primary mb-2">
    Financial Reports &amp; Analytics
   </h1>
   <p class="font-body text-body text-text-secondary">
    Comprehensive overview of asset valuation, depreciation, and total cost of ownership.
   </p>
  </div>
  <div class="flex gap-3">
   <button class="bg-surface border border-border-default px-4 py-2 rounded-lg text-sm font-medium hover:bg-surface-subtle transition-colors flex items-center gap-2 text-text-primary shadow-sm">
    <span class="material-symbols-outlined text-[18px]">
     download
    </span>
    Export PDF
   </button>
   <button class="bg-surface border border-border-default px-4 py-2 rounded-lg text-sm font-medium hover:bg-surface-subtle transition-colors flex items-center gap-2 text-text-primary shadow-sm">
    <span class="material-symbols-outlined text-[18px]">
     calendar_month
    </span>
    FY 2023-24
   </button>
  </div>
 </div>
 <!-- Metric Cards Grid -->
 <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 mb-section-gap">
  <!-- Total Purchase Value -->
  <div class="bg-metric-amber rounded-xl p-card-padding border border-border-default/50 shadow-sm hover:-translate-y-[2px] transition-transform cursor-pointer">
   <div class="flex justify-between items-start mb-4">
    <span class="font-h3 text-h3 text-text-primary">
     Total Purchase Value
    </span>
    <span class="material-symbols-outlined text-amber-700">
     account_balance_wallet
    </span>
   </div>
   <div class="font-kpi-number text-kpi-number text-text-primary mb-1">
    ₹45,20,500
   </div>
   <div class="font-metadata text-metadata text-text-secondary flex items-center gap-1">
    <span class="material-symbols-outlined text-[14px] text-green-700">
     trending_up
    </span>
    <span class="text-green-700 font-medium">
     +2.4%
    </span>
    vs last FY
   </div>
  </div>
  <!-- Accumulated Depreciation -->
  <div class="bg-metric-lavender rounded-xl p-card-padding border border-border-default/50 shadow-sm hover:-translate-y-[2px] transition-transform cursor-pointer">
   <div class="flex justify-between items-start mb-4">
    <span class="font-h3 text-h3 text-text-primary">
     Accumulated Depr.
    </span>
    <span class="material-symbols-outlined text-purple-700">
     trending_down
    </span>
   </div>
   <div class="font-kpi-number text-kpi-number text-text-primary mb-1">
    ₹18,45,200
   </div>
   <div class="font-metadata text-metadata text-text-secondary flex items-center gap-1">
    <span>
     Across 452 active assets
    </span>
   </div>
  </div>
  <!-- Net Book Value -->
  <div class="bg-metric-sage rounded-xl p-card-padding border border-border-default/50 shadow-sm hover:-translate-y-[2px] transition-transform cursor-pointer">
   <div class="flex justify-between items-start mb-4">
    <span class="font-h3 text-h3 text-text-primary">
     Net Book Value
    </span>
    <span class="material-symbols-outlined text-green-700">
     calculate
    </span>
   </div>
   <div class="font-kpi-number text-kpi-number text-text-primary mb-1">
    ₹26,75,300
   </div>
   <div class="font-metadata text-metadata text-text-secondary flex items-center gap-1">
    <span>
     Current valuation estimate
    </span>
   </div>
  </div>
  <!-- TCO Review Required -->
  <div class="bg-surface-container-high rounded-xl p-card-padding border border-border-default/50 shadow-sm hover:-translate-y-[2px] transition-transform cursor-pointer">
   <div class="flex justify-between items-start mb-4">
    <span class="font-h3 text-h3 text-text-primary">
     TCO Review Required
    </span>
    <span class="material-symbols-outlined text-orange-700">
     warning
    </span>
   </div>
   <div class="font-kpi-number text-kpi-number text-text-primary mb-1">
    12 assets
   </div>
   <div class="font-metadata text-metadata text-text-secondary flex items-center gap-1">
    <span class="text-status-critical font-medium">
     Action needed
    </span>
   </div>
  </div>
 </div>
 <!-- Charts Grid -->
 <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-section-gap">
  <!-- Bar Chart: Depreciation by Category -->
  <div class="bg-surface rounded-xl border border-border-default p-card-padding shadow-sm">
   <div class="flex justify-between items-center mb-6">
    <h2 class="font-h2 text-h2 text-text-primary">
     Depreciation by Category
    </h2>
    <button class="text-text-secondary hover:text-text-primary transition-colors">
     <span class="material-symbols-outlined">
      more_vert
     </span>
    </button>
   </div>
   <div class="flex items-center gap-4 mb-2 font-metadata text-metadata">
    <div class="flex items-center gap-1.5">
     <div class="w-3 h-3 rounded-sm bg-secondary-fixed-dim">
     </div>
     <span class="text-text-secondary">
      Depreciated Value
     </span>
    </div>
    <div class="flex items-center gap-1.5">
     <div class="w-3 h-3 rounded-sm bg-primary-fixed-dim">
     </div>
     <span class="text-text-secondary">
      Remaining Book Value
     </span>
    </div>
   </div>
   <div class="bar-chart-container border-b border-border-default pb-2">
    <!-- IT Equipment -->
    <div class="bar-group">
     <div class="bar h-full">
      <div class="bar-segment-top" style="height: 30%;" title="Remaining: ₹8,50,000">
      </div>
      <div class="bar-segment-bottom" style="height: 60%;" title="Depreciated: ₹12,20,000">
      </div>
     </div>
     <span class="bar-label">
      IT Equip.
     </span>
    </div>
    <!-- Furniture -->
    <div class="bar-group">
     <div class="bar h-full">
      <div class="bar-segment-top" style="height: 45%;" title="Remaining: ₹10,20,000">
      </div>
      <div class="bar-segment-bottom" style="height: 35%;" title="Depreciated: ₹5,40,000">
      </div>
     </div>
     <span class="bar-label">
      Furniture
     </span>
    </div>
    <!-- Vehicles -->
    <div class="bar-group">
     <div class="bar h-full">
      <div class="bar-segment-top" style="height: 20%;" title="Remaining: ₹4,50,000">
      </div>
      <div class="bar-segment-bottom" style="height: 70%;" title="Depreciated: ₹15,00,000">
      </div>
     </div>
     <span class="bar-label">
      Vehicles
     </span>
    </div>
    <!-- HVAC -->
    <div class="bar-group">
     <div class="bar h-full">
      <div class="bar-segment-top" style="height: 55%;" title="Remaining: ₹18,00,000">
      </div>
      <div class="bar-segment-bottom" style="height: 25%;" title="Depreciated: ₹6,00,000">
      </div>
     </div>
     <span class="bar-label">
      HVAC
     </span>
    </div>
    <!-- Security -->
    <div class="bar-group">
     <div class="bar h-full">
      <div class="bar-segment-top" style="height: 40%;" title="Remaining: ₹3,50,000">
      </div>
      <div class="bar-segment-bottom" style="height: 40%;" title="Depreciated: ₹3,20,000">
      </div>
     </div>
     <span class="bar-label">
      Security
     </span>
    </div>
   </div>
  </div>
  <!-- Line Chart: Monthly Depreciation Trend -->
  <div class="bg-surface rounded-xl border border-border-default p-card-padding shadow-sm">
   <div class="flex justify-between items-center mb-6">
    <div>
     <h2 class="font-h2 text-h2 text-text-primary">
      Monthly Depreciation Trend
     </h2>
     <span class="font-metadata text-metadata text-text-secondary">
      Financial Year Apr - Mar
     </span>
    </div>
    <button class="text-text-secondary hover:text-text-primary transition-colors">
     <span class="material-symbols-outlined">
      more_vert
     </span>
    </button>
   </div>
   <div class="line-chart-container">
    <svg class="line-chart-svg" preserveaspectratio="none" viewbox="0 0 500 160">
     <defs>
      <lineargradient id="line-gradient" x1="0" x2="0" y1="0" y2="1">
       <stop offset="0%" stop-color="#1E40AF" stop-opacity="0.2">
       </stop>
       <stop offset="100%" stop-color="#1E40AF" stop-opacity="0">
       </stop>
      </lineargradient>
     </defs>
     <!-- Grid Lines -->
     <line class="y-axis-line" x1="40" x2="480" y1="20" y2="20">
     </line>
     <line class="y-axis-line" x1="40" x2="480" y1="60" y2="60">
     </line>
     <line class="y-axis-line" x1="40" x2="480" y1="100" y2="100">
     </line>
     <line class="y-axis-line" x1="40" x2="480" y1="140" y2="140">
     </line>
     <!-- Y Axis Labels -->
     <text class="y-axis-label" x="35" y="24">
      2.0L
     </text>
     <text class="y-axis-label" x="35" y="64">
      1.5L
     </text>
     <text class="y-axis-label" x="35" y="104">
      1.0L
     </text>
     <text class="y-axis-label" x="35" y="144">
      0.5L
     </text>
     <!-- Area & Path -->
     <path class="line-area" d="M 40 140 L 40 120 L 80 115 L 120 110 L 160 105 L 200 90 L 240 85 L 280 80 L 320 60 L 360 55 L 400 45 L 440 35 L 480 30 L 480 140 Z">
     </path>
     <path class="line-path" d="M 40 120 L 80 115 L 120 110 L 160 105 L 200 90 L 240 85 L 280 80 L 320 60 L 360 55 L 400 45 L 440 35 L 480 30">
     </path>
     <!-- Points -->
     <circle class="line-point" cx="40" cy="120" r="4">
     </circle>
     <circle class="line-point" cx="80" cy="115" r="4">
     </circle>
     <circle class="line-point" cx="120" cy="110" r="4">
     </circle>
     <circle class="line-point" cx="160" cy="105" r="4">
     </circle>
     <circle class="line-point" cx="200" cy="90" r="4">
     </circle>
     <circle class="line-point" cx="240" cy="85" r="4">
     </circle>
     <circle class="line-point" cx="280" cy="80" r="4">
     </circle>
     <circle class="line-point" cx="320" cy="60" r="4">
     </circle>
     <circle class="line-point" cx="360" cy="55" r="4">
     </circle>
     <circle class="line-point" cx="400" cy="45" r="4">
     </circle>
     <circle class="line-point" cx="440" cy="35" r="4">
     </circle>
     <circle class="line-point" cx="480" cy="30" r="4">
     </circle>
     <!-- X Axis Labels -->
     <text class="x-axis-label" x="40" y="160">
      Apr
     </text>
     <text class="x-axis-label" x="80" y="160">
      May
     </text>
     <text class="x-axis-label" x="120" y="160">
      Jun
     </text>
     <text class="x-axis-label" x="160" y="160">
      Jul
     </text>
     <text class="x-axis-label" x="200" y="160">
      Aug
     </text>
     <text class="x-axis-label" x="240" y="160">
      Sep
     </text>
     <text class="x-axis-label" x="280" y="160">
      Oct
     </text>
     <text class="x-axis-label" x="320" y="160">
      Nov
     </text>
     <text class="x-axis-label" x="360" y="160">
      Dec
     </text>
     <text class="x-axis-label" x="400" y="160">
      Jan
     </text>
     <text class="x-axis-label" x="440" y="160">
      Feb
     </text>
     <text class="x-axis-label" x="480" y="160">
      Mar
     </text>
    </svg>
   </div>
  </div>
 </div>
 <!-- TCO Table Section -->
 <div class="bg-surface rounded-xl border border-border-default shadow-sm overflow-hidden mb-section-gap">
  <div class="p-card-padding border-b border-border-default flex justify-between items-center bg-surface-subtle">
   <h2 class="font-h2 text-h2 text-text-primary">
    Total Cost of Ownership (TCO) Review
   </h2>
   <div class="flex gap-2">
    <div class="relative">
     <span class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-text-secondary text-[16px]">
      filter_list
     </span>
     <select class="pl-9 pr-8 py-1.5 bg-surface border border-border-default rounded-md text-sm text-text-primary focus:outline-none focus:ring-1 focus:ring-amber-500 appearance-none cursor-pointer">
      <option>
       All Statuses
      </option>
      <option>
       Needs Review
      </option>
      <option>
       Retire
      </option>
     </select>
    </div>
   </div>
  </div>
  <div class="overflow-x-auto">
   <table class="w-full text-left border-collapse">
    <thead>
     <tr class="border-b border-border-default bg-surface">
      <th class="p-4 font-table-header text-table-header text-text-secondary w-1/4">
       ASSET DETAILS
      </th>
      <th class="p-4 font-table-header text-table-header text-text-secondary">
       TAG NO.
      </th>
      <th class="p-4 font-table-header text-table-header text-text-secondary text-right">
       PURCHASE COST
      </th>
      <th class="p-4 font-table-header text-table-header text-text-secondary text-right">
       MAINTENANCE
      </th>
      <th class="p-4 font-table-header text-table-header text-text-secondary text-right">
       TCO
      </th>
      <th class="p-4 font-table-header text-table-header text-text-secondary text-center">
       VERDICT
      </th>
     </tr>
    </thead>
    <tbody class="font-body text-body text-text-primary divide-y divide-border-default">
     <tr class="hover:bg-accent-subtle transition-colors group">
      <td class="p-4">
       <div class="font-medium">
        Dell PowerEdge R740
       </div>
       <div class="font-metadata text-metadata text-text-secondary">
        IT Infrastructure
       </div>
      </td>
      <td class="p-4 font-mono-data text-mono-data">
       TPK-SRV-0042
      </td>
      <td class="p-4 text-right font-mono-data text-mono-data">
       ₹4,50,000
      </td>
      <td class="p-4 text-right font-mono-data text-mono-data">
       ₹1,20,500
      </td>
      <td class="p-4 text-right font-mono-data text-mono-data font-medium">
       ₹5,70,500
      </td>
      <td class="p-4 text-center">
       <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-metric-sage text-status-in-stock">
        <span class="sr-only">
         Status:
        </span>
        Value
       </span>
      </td>
     </tr>
     <!-- Highlighted Row Example (Amber border left) -->
     <tr class="hover:bg-accent-subtle transition-colors group relative bg-surface-subtle border-l-4 border-amber-500">
      <td class="p-4 pl-3">
       <!-- Adjusted padding for border -->
       <div class="font-medium text-amber-900">
        HVAC Unit Block B
       </div>
       <div class="font-metadata text-metadata text-text-secondary">
        Facilities
       </div>
      </td>
      <td class="p-4 font-mono-data text-mono-data">
       TPK-FAC-0118
      </td>
      <td class="p-4 text-right font-mono-data text-mono-data">
       ₹12,00,000
      </td>
      <td class="p-4 text-right font-mono-data text-mono-data text-amber-700">
       ₹8,45,000
      </td>
      <td class="p-4 text-right font-mono-data text-mono-data font-medium">
       ₹20,45,000
      </td>
      <td class="p-4 text-center">
       <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-metric-amber text-amber-800 border border-amber-200">
        <span class="sr-only">
         Status:
        </span>
        Review
       </span>
      </td>
     </tr>
     <tr class="hover:bg-accent-subtle transition-colors group">
      <td class="p-4">
       <div class="font-medium">
        Cisco Catalyst 9300
       </div>
       <div class="font-metadata text-metadata text-text-secondary">
        Networking
       </div>
      </td>
      <td class="p-4 font-mono-data text-mono-data">
       TPK-NET-0089
      </td>
      <td class="p-4 text-right font-mono-data text-mono-data">
       ₹2,80,000
      </td>
      <td class="p-4 text-right font-mono-data text-mono-data">
       ₹45,000
      </td>
      <td class="p-4 text-right font-mono-data text-mono-data font-medium">
       ₹3,25,000
      </td>
      <td class="p-4 text-center">
       <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-metric-sage text-status-in-stock">
        <span class="sr-only">
         Status:
        </span>
        Value
       </span>
      </td>
     </tr>
     <tr class="hover:bg-accent-subtle transition-colors group bg-surface-subtle">
      <td class="p-4">
       <div class="font-medium">
        Mahindra Bolero Pickup
       </div>
       <div class="font-metadata text-metadata text-text-secondary">
        Fleet Vehicles
       </div>
      </td>
      <td class="p-4 font-mono-data text-mono-data">
       TPK-VEH-0012
      </td>
      <td class="p-4 text-right font-mono-data text-mono-data">
       ₹8,50,000
      </td>
      <td class="p-4 text-right font-mono-data text-mono-data text-red-600">
       ₹9,80,000
      </td>
      <td class="p-4 text-right font-mono-data text-mono-data font-medium">
       ₹18,30,000
      </td>
      <td class="p-4 text-center">
       <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-error-container text-status-critical">
        <span class="sr-only">
         Status:
        </span>
        Retire
       </span>
      </td>
     </tr>
     <tr class="hover:bg-accent-subtle transition-colors group">
      <td class="p-4">
       <div class="font-medium">
        Conference Table Oak
       </div>
       <div class="font-metadata text-metadata text-text-secondary">
        Furniture
       </div>
      </td>
      <td class="p-4 font-mono-data text-mono-data">
       TPK-FURN-0302
      </td>
      <td class="p-4 text-right font-mono-data text-mono-data">
       ₹1,20,000
      </td>
      <td class="p-4 text-right font-mono-data text-mono-data">
       ₹15,000
      </td>
      <td class="p-4 text-right font-mono-data text-mono-data font-medium">
       ₹1,35,000
      </td>
      <td class="p-4 text-center">
       <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-metric-sage text-status-in-stock">
        <span class="sr-only">
         Status:
        </span>
        Value
       </span>
      </td>
     </tr>
    </tbody>
   </table>
  </div>
  <div class="p-4 border-t border-border-default bg-surface flex items-center justify-between font-metadata text-metadata text-text-secondary">
   <span>
    Showing 1 to 5 of 452 entries
   </span>
   <div class="flex gap-1">
    <button class="px-2 py-1 rounded hover:bg-surface-subtle disabled:opacity-50 border border-transparent" disabled="">
     &lt; Prev
    </button>
    <button class="px-2 py-1 rounded bg-surface-subtle border border-border-default text-text-primary font-medium">
     1
    </button>
    <button class="px-2 py-1 rounded hover:bg-surface-subtle border border-transparent">
     2
    </button>
    <button class="px-2 py-1 rounded hover:bg-surface-subtle border border-transparent">
     3
    </button>
    <button class="px-2 py-1 rounded hover:bg-surface-subtle border border-transparent">
     Next &gt;
    </button>
   </div>
  </div>
 </div>
</main>

</template>

<script setup lang="ts">
// Autogenerated from financial_reports_analytics
</script>
