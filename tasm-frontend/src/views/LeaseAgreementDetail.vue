<template>
<main class="flex-1 mt-[60px] p-page-margin max-w-[1400px] mx-auto w-full">
 <!-- Page Header -->
 <div class="flex flex-col md:flex-row md:items-center justify-between gap-4 mb-section-gap">
  <div>
   <div class="flex items-center gap-3 mb-1">
    <a class="text-text-secondary hover:text-primary-container transition-colors" href="#">
     <span class="material-symbols-outlined text-sm">
      arrow_back
     </span>
    </a>
    <span class="font-metadata text-metadata text-text-secondary">
     Financials / Leases
    </span>
   </div>
   <div class="flex items-center gap-3">
    <h1 class="font-h1 text-h1 text-text-primary tracking-tight">
     Lease Agreement: LS-4055 - Tata Elxsi
    </h1>
    <span class="inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full bg-metric-sage text-status-in-stock font-metadata text-metadata border border-[#86efac]/50">
     <span class="w-1.5 h-1.5 rounded-full bg-status-in-stock">
     </span>
     Active
    </span>
   </div>
  </div>
  <div class="flex items-center gap-3">
   <button class="px-4 py-2 bg-surface border border-border-default rounded-lg font-h3 text-h3 text-text-primary hover:bg-surface-subtle transition-colors shadow-sm active:translate-y-0.5 flex items-center gap-2">
    <span class="material-symbols-outlined text-sm">
     edit
    </span>
    Edit
   </button>
   <button class="px-4 py-2 bg-surface border border-border-default rounded-lg font-h3 text-h3 text-text-primary hover:bg-surface-subtle transition-colors shadow-sm active:translate-y-0.5 flex items-center gap-2">
    <span class="material-symbols-outlined text-sm">
     download
    </span>
    Export PDF
   </button>
  </div>
 </div>
 <!-- Top Row: Metrics -->
 <div class="grid grid-cols-1 md:grid-cols-3 gap-stack mb-section-gap">
  <!-- Metric 1 -->
  <div class="bg-metric-amber rounded-[16px] p-card-padding border border-[#fde68a] shadow-sm hover:-translate-y-0.5 transition-transform cursor-pointer">
   <div class="flex justify-between items-start mb-2">
    <span class="font-metadata text-metadata text-amber-900 uppercase tracking-wider">
     Monthly Value
    </span>
    <span class="material-symbols-outlined text-amber-700">
     payments
    </span>
   </div>
   <div class="font-kpi-number text-kpi-number text-amber-950 mb-1">
    ₹4,50,000
   </div>
   <div class="font-metadata text-metadata text-amber-800">
    Due 5th of every month
   </div>
  </div>
  <!-- Metric 2 -->
  <div class="bg-metric-lavender rounded-[16px] p-card-padding border border-[#e9d5ff] shadow-sm hover:-translate-y-0.5 transition-transform cursor-pointer">
   <div class="flex justify-between items-start mb-2">
    <span class="font-metadata text-metadata text-violet-900 uppercase tracking-wider">
     Lease Duration
    </span>
    <span class="material-symbols-outlined text-violet-700">
     hourglass_empty
    </span>
   </div>
   <div class="font-kpi-number text-kpi-number text-violet-950 mb-1">
    36 Months
   </div>
   <div class="font-metadata text-metadata text-violet-800">
    Commenced: 01-Jan-2023
   </div>
  </div>
  <!-- Metric 3 -->
  <div class="bg-surface rounded-[16px] p-card-padding border border-border-default shadow-sm hover:-translate-y-0.5 transition-transform cursor-pointer">
   <div class="flex justify-between items-start mb-2">
    <span class="font-metadata text-metadata text-text-secondary uppercase tracking-wider">
     Expiry Date
    </span>
    <span class="material-symbols-outlined text-text-secondary">
     calendar_today
    </span>
   </div>
   <div class="font-kpi-number text-kpi-number text-text-primary mb-1">
    31-Dec-2025
   </div>
   <div class="font-metadata text-metadata text-text-secondary">
    712 days remaining
   </div>
  </div>
 </div>
 <!-- Middle Section: Terms and Assets -->
 <div class="grid grid-cols-1 lg:grid-cols-12 gap-stack mb-section-gap">
  <!-- Contract Terms -->
  <div class="lg:col-span-4 bg-surface rounded-[16px] border border-border-default shadow-sm overflow-hidden flex flex-col">
   <div class="p-4 border-b border-border-default bg-surface-subtle">
    <h2 class="font-h2 text-h2 text-text-primary flex items-center gap-2">
     <span class="material-symbols-outlined text-lg">
      gavel
     </span>
     Contract Terms
    </h2>
   </div>
   <div class="p-card-padding flex-1 overflow-y-auto">
    <div class="space-y-6">
     <div>
      <h3 class="font-h3 text-h3 text-text-primary mb-1">
       Premises
      </h3>
      <p class="font-body text-body text-text-secondary">
       Unit 4B, 4th Floor, Tejaswini Building, Technopark Campus. Total super built-up area: 12,500 Sq.Ft.
      </p>
     </div>
     <div>
      <h3 class="font-h3 text-h3 text-text-primary mb-1">
       Escalation Clause
      </h3>
      <p class="font-body text-body text-text-secondary">
       5% increase on base rent annually, effective from the 13th month of the lease term.
      </p>
     </div>
     <div>
      <h3 class="font-h3 text-h3 text-text-primary mb-1">
       Termination Notice
      </h3>
      <p class="font-body text-body text-text-secondary">
       90 days written notice required from either party. Lock-in period of 12 months applies.
      </p>
     </div>
     <div>
      <h3 class="font-h3 text-h3 text-text-primary mb-1">
       Maintenance Charges
      </h3>
      <p class="font-body text-body text-text-secondary">
       ₹15 per Sq.Ft per month, payable in advance alongside monthly rent.
      </p>
     </div>
    </div>
   </div>
  </div>
  <!-- Leased Assets -->
  <div class="lg:col-span-8 bg-surface rounded-[16px] border border-border-default shadow-sm overflow-hidden flex flex-col">
   <div class="p-4 border-b border-border-default bg-surface-subtle flex justify-between items-center">
    <h2 class="font-h2 text-h2 text-text-primary flex items-center gap-2">
     <span class="material-symbols-outlined text-lg">
      inventory_2
     </span>
     Leased Assets (Fixtures)
    </h2>
    <button class="text-primary-container font-h3 text-h3 hover:underline">
     View All
    </button>
   </div>
   <div class="overflow-x-auto">
    <table class="w-full text-left border-collapse">
     <thead>
      <tr class="bg-surface-subtle border-b border-border-default">
       <th class="py-3 px-4 font-table-header text-table-header text-text-secondary uppercase">
        Tag ID
       </th>
       <th class="py-3 px-4 font-table-header text-table-header text-text-secondary uppercase">
        Asset Name
       </th>
       <th class="py-3 px-4 font-table-header text-table-header text-text-secondary uppercase">
        Category
       </th>
       <th class="py-3 px-4 font-table-header text-table-header text-text-secondary uppercase">
        Condition
       </th>
      </tr>
     </thead>
     <tbody class="font-body text-body text-text-primary divide-y divide-border-default">
      <tr class="hover:bg-accent-subtle transition-colors group">
       <td class="py-3 px-4 font-mono-data text-mono-data">
        AST-HVAC-001
       </td>
       <td class="py-3 px-4 font-medium">
        Carrier 5T AHU
       </td>
       <td class="py-3 px-4 text-text-secondary">
        HVAC
       </td>
       <td class="py-3 px-4">
        <span class="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium bg-surface-container text-text-primary border border-border-default">
         Good
        </span>
       </td>
      </tr>
      <tr class="hover:bg-accent-subtle transition-colors group">
       <td class="py-3 px-4 font-mono-data text-mono-data">
        AST-ELEC-045
       </td>
       <td class="py-3 px-4 font-medium">
        Main Distribution Board
       </td>
       <td class="py-3 px-4 text-text-secondary">
        Electrical
       </td>
       <td class="py-3 px-4">
        <span class="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium bg-surface-container text-text-primary border border-border-default">
         Excellent
        </span>
       </td>
      </tr>
      <tr class="hover:bg-accent-subtle transition-colors group border-l-[3px] border-l-transparent hover:border-l-primary-container">
       <td class="py-3 px-4 font-mono-data text-mono-data">
        AST-FURN-112
       </td>
       <td class="py-3 px-4 font-medium">
        Modular Workstations (x50)
       </td>
       <td class="py-3 px-4 text-text-secondary">
        Furniture
       </td>
       <td class="py-3 px-4">
        <span class="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium bg-surface-container text-text-primary border border-border-default">
         Fair
        </span>
       </td>
      </tr>
      <tr class="hover:bg-accent-subtle transition-colors group">
       <td class="py-3 px-4 font-mono-data text-mono-data">
        AST-FIRE-008
       </td>
       <td class="py-3 px-4 font-medium">
        Sprinkler System Grid B
       </td>
       <td class="py-3 px-4 text-text-secondary">
        Safety
       </td>
       <td class="py-3 px-4">
        <span class="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium bg-surface-container text-text-primary border border-border-default">
         Good
        </span>
       </td>
      </tr>
      <tr class="hover:bg-accent-subtle transition-colors group">
       <td class="py-3 px-4 font-mono-data text-mono-data">
        AST-IT-099
       </td>
       <td class="py-3 px-4 font-medium">
        Server Rack Cabinet
       </td>
       <td class="py-3 px-4 text-text-secondary">
        IT Infra
       </td>
       <td class="py-3 px-4">
        <span class="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium bg-surface-container text-text-primary border border-border-default">
         Good
        </span>
       </td>
      </tr>
     </tbody>
    </table>
   </div>
  </div>
 </div>
 <!-- Bottom Section: Payment Schedule & Docs -->
 <div class="grid grid-cols-1 lg:grid-cols-2 gap-stack">
  <!-- Payment Schedule -->
  <div class="bg-surface rounded-[16px] border border-border-default shadow-sm overflow-hidden flex flex-col">
   <div class="p-4 border-b border-border-default bg-surface-subtle">
    <h2 class="font-h2 text-h2 text-text-primary flex items-center gap-2">
     <span class="material-symbols-outlined text-lg">
      receipt_long
     </span>
     Payment Schedule (Upcoming)
    </h2>
   </div>
   <div class="overflow-x-auto">
    <table class="w-full text-left border-collapse">
     <thead>
      <tr class="bg-surface border-b border-border-default">
       <th class="py-3 px-4 font-table-header text-table-header text-text-secondary uppercase">
        Due Date
       </th>
       <th class="py-3 px-4 font-table-header text-table-header text-text-secondary uppercase">
        Type
       </th>
       <th class="py-3 px-4 font-table-header text-table-header text-text-secondary uppercase text-right">
        Amount
       </th>
       <th class="py-3 px-4 font-table-header text-table-header text-text-secondary uppercase text-center">
        Status
       </th>
      </tr>
     </thead>
     <tbody class="font-body text-body text-text-primary divide-y divide-border-default">
      <tr class="hover:bg-accent-subtle transition-colors">
       <td class="py-3 px-4">
        05-Nov-2023
       </td>
       <td class="py-3 px-4 text-text-secondary">
        Monthly Rent
       </td>
       <td class="py-3 px-4 font-mono-data text-mono-data text-right">
        ₹4,50,000
       </td>
       <td class="py-3 px-4 text-center">
        <span class="inline-flex items-center gap-1 px-2 py-0.5 rounded-full text-xs font-medium bg-metric-sage text-status-in-stock border border-[#86efac]/50">
         <span class="material-symbols-outlined text-[12px]">
          check_circle
         </span>
         Paid
        </span>
       </td>
      </tr>
      <tr class="hover:bg-accent-subtle transition-colors bg-[#FAFAF8]">
       <td class="py-3 px-4 font-medium">
        05-Dec-2023
       </td>
       <td class="py-3 px-4 text-text-secondary">
        Monthly Rent + Maint.
       </td>
       <td class="py-3 px-4 font-mono-data text-mono-data text-right font-medium">
        ₹6,37,500
       </td>
       <td class="py-3 px-4 text-center">
        <span class="inline-flex items-center gap-1 px-2 py-0.5 rounded-full text-xs font-medium bg-[#fef08a]/50 text-amber-800 border border-[#fef08a]">
         <span class="material-symbols-outlined text-[12px]">
          schedule
         </span>
         Pending
        </span>
       </td>
      </tr>
      <tr class="hover:bg-accent-subtle transition-colors">
       <td class="py-3 px-4">
        05-Jan-2024
       </td>
       <td class="py-3 px-4 text-text-secondary">
        Monthly Rent
       </td>
       <td class="py-3 px-4 font-mono-data text-mono-data text-right">
        ₹4,72,500
       </td>
       <td class="py-3 px-4 text-center">
        <span class="inline-flex items-center gap-1 px-2 py-0.5 rounded-full text-xs font-medium bg-surface-container text-text-secondary border border-border-default">
         Scheduled
        </span>
       </td>
      </tr>
     </tbody>
    </table>
   </div>
  </div>
  <!-- Supporting Documents -->
  <div class="bg-surface rounded-[16px] border border-border-default shadow-sm overflow-hidden flex flex-col">
   <div class="p-4 border-b border-border-default bg-surface-subtle flex justify-between items-center">
    <h2 class="font-h2 text-h2 text-text-primary flex items-center gap-2">
     <span class="material-symbols-outlined text-lg">
      folder_open
     </span>
     Supporting Documents
    </h2>
    <button class="text-primary-container hover:bg-surface-container p-1 rounded transition-colors">
     <span class="material-symbols-outlined">
      add
     </span>
    </button>
   </div>
   <div class="p-4 space-y-3">
    <div class="flex items-center justify-between p-3 border border-border-default rounded-lg hover:border-primary-container transition-colors group cursor-pointer bg-surface">
     <div class="flex items-center gap-3">
      <div class="w-10 h-10 rounded bg-red-50 flex items-center justify-center text-red-600">
       <span class="material-symbols-outlined">
        picture_as_pdf
       </span>
      </div>
      <div>
       <p class="font-h3 text-h3 text-text-primary group-hover:text-primary-container transition-colors">
        Signed_Lease_Agreement_TataElxsi.pdf
       </p>
       <p class="font-metadata text-metadata text-text-secondary">
        2.4 MB • Uploaded 15-Dec-2022
       </p>
      </div>
     </div>
     <button class="text-text-secondary hover:text-text-primary opacity-0 group-hover:opacity-100 transition-opacity">
      <span class="material-symbols-outlined">
       download
      </span>
     </button>
    </div>
    <div class="flex items-center justify-between p-3 border border-border-default rounded-lg hover:border-primary-container transition-colors group cursor-pointer bg-surface">
     <div class="flex items-center gap-3">
      <div class="w-10 h-10 rounded bg-blue-50 flex items-center justify-center text-blue-600">
       <span class="material-symbols-outlined">
        description
       </span>
      </div>
      <div>
       <p class="font-h3 text-h3 text-text-primary group-hover:text-primary-container transition-colors">
        Asset_Handover_Checklist_Rev2.docx
       </p>
       <p class="font-metadata text-metadata text-text-secondary">
        856 KB • Uploaded 02-Jan-2023
       </p>
      </div>
     </div>
     <button class="text-text-secondary hover:text-text-primary opacity-0 group-hover:opacity-100 transition-opacity">
      <span class="material-symbols-outlined">
       download
      </span>
     </button>
    </div>
    <div class="flex items-center justify-between p-3 border border-border-default rounded-lg hover:border-primary-container transition-colors group cursor-pointer bg-surface">
     <div class="flex items-center gap-3">
      <div class="w-10 h-10 rounded bg-red-50 flex items-center justify-center text-red-600">
       <span class="material-symbols-outlined">
        picture_as_pdf
       </span>
      </div>
      <div>
       <p class="font-h3 text-h3 text-text-primary group-hover:text-primary-container transition-colors">
        Floor_Plan_Unit4B_Approved.pdf
       </p>
       <p class="font-metadata text-metadata text-text-secondary">
        5.1 MB • Uploaded 10-Dec-2022
       </p>
      </div>
     </div>
     <button class="text-text-secondary hover:text-text-primary opacity-0 group-hover:opacity-100 transition-opacity">
      <span class="material-symbols-outlined">
       download
      </span>
     </button>
    </div>
   </div>
  </div>
 </div>
</main>

</template>

<script setup lang="ts">
// Autogenerated from lease_agreement_detail
</script>
