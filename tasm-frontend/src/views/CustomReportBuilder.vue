<template>
<main class="flex-1 md:ml-[248px] p-page-margin overflow-y-auto bg-canvas">
 <!-- Page Header -->
 <div class="mb-section-gap flex flex-col md:flex-row md:items-end justify-between gap-4">
  <div>
   <div class="flex items-center gap-2 text-text-secondary font-metadata mb-2">
    <span>
     Compliance
    </span>
    <span class="material-symbols-outlined text-[14px]">
     chevron_right
    </span>
    <span>
     Reports
    </span>
   </div>
   <h1 class="font-h1 text-h1 text-text-primary">
    Custom Report Builder
   </h1>
   <p class="font-body text-body text-text-secondary mt-1">
    Configure and generate detailed compliance and operational reports.
   </p>
  </div>
 </div>
 <!-- Main Layout Grid -->
 <div class="grid grid-cols-1 lg:grid-cols-4 gap-6 items-start">
  <!-- Left Column: Steps Navigation -->
  <div class="lg:col-span-1 bg-surface rounded-xl border border-border-default p-card-padding shadow-sm sticky top-[84px]">
   <h3 class="font-h3 text-h3 text-text-primary mb-6">
    Report Configuration
   </h3>
   <div class="space-y-6 relative before:absolute before:inset-0 before:ml-4 before:-translate-x-px md:before:mx-auto md:before:translate-x-0 before:h-full before:w-0.5 before:bg-gradient-to-b before:from-transparent before:via-border-default before:to-transparent">
    <!-- Step 1 (Active) -->
    <div class="relative flex items-center justify-between md:justify-normal md:odd:flex-row-reverse group is-active">
     <div class="flex items-center justify-center w-8 h-8 rounded-full border-2 border-primary bg-accent-subtle text-primary shadow shrink-0 md:order-1 md:group-odd:-translate-x-1/2 md:group-even:translate-x-1/2 relative z-10">
      <span class="font-mono-data text-mono-data font-bold">
       1
      </span>
     </div>
     <div class="w-[calc(100%-3rem)] md:w-[calc(50%-2rem)] md:group-odd:text-right md:group-odd:pr-4 md:group-even:pl-4">
      <h4 class="font-h3 text-h3 text-text-primary">
       Base Data
      </h4>
      <p class="font-metadata text-metadata text-text-secondary">
       Select primary entity
      </p>
     </div>
    </div>
    <!-- Step 2 -->
    <div class="relative flex items-center justify-between md:justify-normal md:odd:flex-row-reverse group">
     <div class="flex items-center justify-center w-8 h-8 rounded-full border-2 border-border-default bg-surface text-text-secondary shrink-0 md:order-1 md:group-odd:-translate-x-1/2 md:group-even:translate-x-1/2 relative z-10">
      <span class="font-mono-data text-mono-data">
       2
      </span>
     </div>
     <div class="w-[calc(100%-3rem)] md:w-[calc(50%-2rem)] md:group-odd:text-right md:group-odd:pr-4 md:group-even:pl-4 opacity-70">
      <h4 class="font-h3 text-h3 text-text-primary">
       Columns
      </h4>
      <p class="font-metadata text-metadata text-text-secondary">
       Choose data fields
      </p>
     </div>
    </div>
    <!-- Step 3 -->
    <div class="relative flex items-center justify-between md:justify-normal md:odd:flex-row-reverse group">
     <div class="flex items-center justify-center w-8 h-8 rounded-full border-2 border-border-default bg-surface text-text-secondary shrink-0 md:order-1 md:group-odd:-translate-x-1/2 md:group-even:translate-x-1/2 relative z-10">
      <span class="font-mono-data text-mono-data">
       3
      </span>
     </div>
     <div class="w-[calc(100%-3rem)] md:w-[calc(50%-2rem)] md:group-odd:text-right md:group-odd:pr-4 md:group-even:pl-4 opacity-70">
      <h4 class="font-h3 text-h3 text-text-primary">
       Filters
      </h4>
      <p class="font-metadata text-metadata text-text-secondary">
       Refine data subset
      </p>
     </div>
    </div>
    <!-- Step 4 -->
    <div class="relative flex items-center justify-between md:justify-normal md:odd:flex-row-reverse group">
     <div class="flex items-center justify-center w-8 h-8 rounded-full border-2 border-border-default bg-surface text-text-secondary shrink-0 md:order-1 md:group-odd:-translate-x-1/2 md:group-even:translate-x-1/2 relative z-10">
      <span class="font-mono-data text-mono-data">
       4
      </span>
     </div>
     <div class="w-[calc(100%-3rem)] md:w-[calc(50%-2rem)] md:group-odd:text-right md:group-odd:pr-4 md:group-even:pl-4 opacity-70">
      <h4 class="font-h3 text-h3 text-text-primary">
       Group &amp; Sort
      </h4>
      <p class="font-metadata text-metadata text-text-secondary">
       Organize structure
      </p>
     </div>
    </div>
   </div>
  </div>
  <!-- Right Column: Form Area -->
  <div class="lg:col-span-3 space-y-6">
   <!-- Form Container -->
   <div class="bg-surface rounded-xl border border-border-default shadow-sm overflow-hidden">
    <!-- Step 1 Content -->
    <div class="p-card-padding border-b border-border-default">
     <h2 class="font-h2 text-h2 text-text-primary mb-6 flex items-center gap-2">
      <span class="material-symbols-outlined text-primary">
       database
      </span>
      Step 1: Select Base Data Source
     </h2>
     <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
      <div class="space-y-2">
       <label class="font-h3 text-h3 text-text-primary block">
        Primary Entity
       </label>
       <div class="relative">
        <select class="w-full bg-surface border border-border-default rounded-lg pl-4 pr-10 py-2.5 font-body text-body text-text-primary focus:outline-none focus:border-primary focus:ring-1 focus:ring-primary appearance-none cursor-pointer shadow-sm hover:translate-y-[-1px] transition-transform">
         <option selected value="assets">
          Fixed Assets
         </option>
         <option value="wo">
          Work Orders
         </option>
         <option value="depreciation">
          Depreciation Schedules
         </option>
         <option value="compliance">
          Compliance Logs
         </option>
        </select>
        <span class="material-symbols-outlined absolute right-3 top-1/2 -translate-y-1/2 text-text-secondary pointer-events-none">
         expand_more
        </span>
       </div>
       <p class="font-metadata text-metadata text-text-secondary">
        Determines the core dataset for this report.
       </p>
      </div>
      <div class="space-y-2">
       <label class="font-h3 text-h3 text-text-primary block">
        Report Timeframe
       </label>
       <div class="relative">
        <select class="w-full bg-surface border border-border-default rounded-lg pl-4 pr-10 py-2.5 font-body text-body text-text-primary focus:outline-none focus:border-primary focus:ring-1 focus:ring-primary appearance-none cursor-pointer shadow-sm hover:translate-y-[-1px] transition-transform">
         <option value="mtd">
          Month to Date (MTD)
         </option>
         <option value="qtd">
          Quarter to Date (QTD)
         </option>
         <option selected value="ytd">
          Year to Date (YTD)
         </option>
         <option value="custom">
          Custom Range
         </option>
        </select>
        <span class="material-symbols-outlined absolute right-3 top-1/2 -translate-y-1/2 text-text-secondary pointer-events-none">
         expand_more
        </span>
       </div>
      </div>
     </div>
    </div>
    <!-- Step 2 Content -->
    <div class="p-card-padding border-b border-border-default bg-surface-subtle">
     <h2 class="font-h2 text-h2 text-text-primary mb-4 flex items-center gap-2">
      <span class="material-symbols-outlined text-primary">
       view_column
      </span>
      Step 2: Select Columns
     </h2>
     <p class="font-body text-body text-text-secondary mb-4">
      Choose which data fields to include in the report output.
     </p>
     <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-y-4 gap-x-6 bg-surface p-4 rounded-lg border border-border-default shadow-sm">
      <!-- Checkbox Item -->
      <label class="flex items-center gap-3 cursor-pointer group">
       <div class="relative flex items-center justify-center w-5 h-5">
        <input checked class="peer w-5 h-5 appearance-none border border-outline rounded bg-surface checked:bg-primary checked:border-primary focus:outline-none focus:ring-2 focus:ring-primary/20 transition-all" type="checkbox"/>
        <span class="material-symbols-outlined absolute text-white text-[16px] pointer-events-none opacity-0 peer-checked:opacity-100 transition-opacity">
         check
        </span>
       </div>
       <span class="font-body text-body text-text-primary group-hover:text-primary transition-colors">
        Asset ID
       </span>
      </label>
      <label class="flex items-center gap-3 cursor-pointer group">
       <div class="relative flex items-center justify-center w-5 h-5">
        <input checked class="peer w-5 h-5 appearance-none border border-outline rounded bg-surface checked:bg-primary checked:border-primary focus:outline-none focus:ring-2 focus:ring-primary/20 transition-all" type="checkbox"/>
        <span class="material-symbols-outlined absolute text-white text-[16px] pointer-events-none opacity-0 peer-checked:opacity-100 transition-opacity">
         check
        </span>
       </div>
       <span class="font-body text-body text-text-primary group-hover:text-primary transition-colors">
        Asset Name
       </span>
      </label>
      <label class="flex items-center gap-3 cursor-pointer group">
       <div class="relative flex items-center justify-center w-5 h-5">
        <input checked class="peer w-5 h-5 appearance-none border border-outline rounded bg-surface checked:bg-primary checked:border-primary focus:outline-none focus:ring-2 focus:ring-primary/20 transition-all" type="checkbox"/>
        <span class="material-symbols-outlined absolute text-white text-[16px] pointer-events-none opacity-0 peer-checked:opacity-100 transition-opacity">
         check
        </span>
       </div>
       <span class="font-body text-body text-text-primary group-hover:text-primary transition-colors">
        Category
       </span>
      </label>
      <label class="flex items-center gap-3 cursor-pointer group">
       <div class="relative flex items-center justify-center w-5 h-5">
        <input checked class="peer w-5 h-5 appearance-none border border-outline rounded bg-surface checked:bg-primary checked:border-primary focus:outline-none focus:ring-2 focus:ring-primary/20 transition-all" type="checkbox"/>
        <span class="material-symbols-outlined absolute text-white text-[16px] pointer-events-none opacity-0 peer-checked:opacity-100 transition-opacity">
         check
        </span>
       </div>
       <span class="font-body text-body text-text-primary group-hover:text-primary transition-colors">
        Location
       </span>
      </label>
      <label class="flex items-center gap-3 cursor-pointer group">
       <div class="relative flex items-center justify-center w-5 h-5">
        <input class="peer w-5 h-5 appearance-none border border-outline rounded bg-surface checked:bg-primary checked:border-primary focus:outline-none focus:ring-2 focus:ring-primary/20 transition-all" type="checkbox"/>
        <span class="material-symbols-outlined absolute text-white text-[16px] pointer-events-none opacity-0 peer-checked:opacity-100 transition-opacity">
         check
        </span>
       </div>
       <span class="font-body text-body text-text-primary group-hover:text-primary transition-colors">
        Purchase Date
       </span>
      </label>
      <label class="flex items-center gap-3 cursor-pointer group">
       <div class="relative flex items-center justify-center w-5 h-5">
        <input checked class="peer w-5 h-5 appearance-none border border-outline rounded bg-surface checked:bg-primary checked:border-primary focus:outline-none focus:ring-2 focus:ring-primary/20 transition-all" type="checkbox"/>
        <span class="material-symbols-outlined absolute text-white text-[16px] pointer-events-none opacity-0 peer-checked:opacity-100 transition-opacity">
         check
        </span>
       </div>
       <span class="font-body text-body text-text-primary group-hover:text-primary transition-colors">
        Status
       </span>
      </label>
      <label class="flex items-center gap-3 cursor-pointer group">
       <div class="relative flex items-center justify-center w-5 h-5">
        <input class="peer w-5 h-5 appearance-none border border-outline rounded bg-surface checked:bg-primary checked:border-primary focus:outline-none focus:ring-2 focus:ring-primary/20 transition-all" type="checkbox"/>
        <span class="material-symbols-outlined absolute text-white text-[16px] pointer-events-none opacity-0 peer-checked:opacity-100 transition-opacity">
         check
        </span>
       </div>
       <span class="font-body text-body text-text-primary group-hover:text-primary transition-colors">
        Custom Field: Dep. Rate
       </span>
      </label>
     </div>
    </div>
    <!-- Step 3 Content -->
    <div class="p-card-padding border-b border-border-default">
     <h2 class="font-h2 text-h2 text-text-primary mb-4 flex items-center gap-2">
      <span class="material-symbols-outlined text-primary">
       filter_list
      </span>
      Step 3: Apply Filters
     </h2>
     <div class="flex flex-wrap gap-2 mb-4">
      <!-- Filter Chip -->
      <div class="inline-flex items-center gap-2 px-3 py-1.5 bg-accent-subtle border border-primary/20 rounded-full text-primary font-metadata text-metadata">
       <span>
        Category: IT Equipment
       </span>
       <button class="hover:text-stone-900 focus:outline-none flex items-center">
        <span class="material-symbols-outlined text-[14px]">
         close
        </span>
       </button>
      </div>
      <div class="inline-flex items-center gap-2 px-3 py-1.5 bg-accent-subtle border border-primary/20 rounded-full text-primary font-metadata text-metadata">
       <span>
        Status: Active
       </span>
       <button class="hover:text-stone-900 focus:outline-none flex items-center">
        <span class="material-symbols-outlined text-[14px]">
         close
        </span>
       </button>
      </div>
      <button class="inline-flex items-center gap-1 px-3 py-1.5 border border-dashed border-outline rounded-full text-text-secondary hover:text-text-primary hover:border-text-primary transition-colors font-metadata text-metadata">
       <span class="material-symbols-outlined text-[14px]">
        add
       </span>
       Add Filter
      </button>
     </div>
    </div>
    <!-- Preview Section -->
    <div class="p-card-padding bg-surface">
     <div class="flex items-center justify-between mb-4">
      <h2 class="font-h2 text-h2 text-text-primary flex items-center gap-2">
       <span class="material-symbols-outlined text-primary">
        visibility
       </span>
       Live Preview
      </h2>
      <span class="font-metadata text-metadata text-text-secondary bg-surface-variant px-2 py-1 rounded">
       Showing first 5 rows
      </span>
     </div>
     <div class="overflow-x-auto rounded-lg border border-border-default shadow-sm">
      <table class="w-full text-left border-collapse">
       <thead>
        <tr class="bg-surface-subtle border-b border-border-default font-table-header text-table-header text-text-secondary">
         <th class="py-3 px-4 uppercase tracking-wider">
          Asset ID
         </th>
         <th class="py-3 px-4 uppercase tracking-wider">
          Asset Name
         </th>
         <th class="py-3 px-4 uppercase tracking-wider">
          Category
         </th>
         <th class="py-3 px-4 uppercase tracking-wider">
          Location
         </th>
         <th class="py-3 px-4 uppercase tracking-wider">
          Status
         </th>
        </tr>
       </thead>
       <tbody class="font-mono-data text-mono-data text-text-primary divide-y divide-border-default">
        <tr class="hover:bg-accent-subtle transition-colors group">
         <td class="py-3 px-4 group-hover:text-primary">
          AST-1042
         </td>
         <td class="py-3 px-4 font-body text-body">
          Dell Precision 5550
         </td>
         <td class="py-3 px-4 font-body text-body">
          IT Equipment
         </td>
         <td class="py-3 px-4 font-body text-body">
          Bldg 3, Floor 2
         </td>
         <td class="py-3 px-4">
          <span class="inline-flex items-center px-2 py-0.5 rounded-full text-[11px] font-medium bg-metric-sage text-status-in-stock">
           Active
          </span>
         </td>
        </tr>
        <tr class="hover:bg-accent-subtle transition-colors group">
         <td class="py-3 px-4 group-hover:text-primary">
          AST-1043
         </td>
         <td class="py-3 px-4 font-body text-body">
          Cisco Meraki MR46
         </td>
         <td class="py-3 px-4 font-body text-body">
          IT Equipment
         </td>
         <td class="py-3 px-4 font-body text-body">
          Bldg 3, Ceiling Hub
         </td>
         <td class="py-3 px-4">
          <span class="inline-flex items-center px-2 py-0.5 rounded-full text-[11px] font-medium bg-metric-sage text-status-in-stock">
           Active
          </span>
         </td>
        </tr>
        <tr class="hover:bg-accent-subtle transition-colors group">
         <td class="py-3 px-4 group-hover:text-primary">
          AST-1050
         </td>
         <td class="py-3 px-4 font-body text-body">
          Herman Miller Aeron
         </td>
         <td class="py-3 px-4 font-body text-body">
          Furniture
         </td>
         <td class="py-3 px-4 font-body text-body">
          Bldg 1, Exec Suite
         </td>
         <td class="py-3 px-4">
          <span class="inline-flex items-center px-2 py-0.5 rounded-full text-[11px] font-medium bg-metric-sage text-status-in-stock">
           Active
          </span>
         </td>
        </tr>
        <tr class="hover:bg-accent-subtle transition-colors group">
         <td class="py-3 px-4 group-hover:text-primary">
          AST-1055
         </td>
         <td class="py-3 px-4 font-body text-body">
          MacBook Pro 16"
         </td>
         <td class="py-3 px-4 font-body text-body">
          IT Equipment
         </td>
         <td class="py-3 px-4 font-body text-body">
          Remote / Assigned
         </td>
         <td class="py-3 px-4">
          <span class="inline-flex items-center px-2 py-0.5 rounded-full text-[11px] font-medium bg-metric-lavender text-status-checked-out">
           Assigned
          </span>
         </td>
        </tr>
        <tr class="hover:bg-accent-subtle transition-colors group">
         <td class="py-3 px-4 group-hover:text-primary">
          AST-1061
         </td>
         <td class="py-3 px-4 font-body text-body">
          Epson Projector ProL
         </td>
         <td class="py-3 px-4 font-body text-body">
          AV Equipment
         </td>
         <td class="py-3 px-4 font-body text-body">
          Conf Room A
         </td>
         <td class="py-3 px-4">
          <span class="inline-flex items-center px-2 py-0.5 rounded-full text-[11px] font-medium bg-error-container text-status-critical">
           Maintenance
          </span>
         </td>
        </tr>
       </tbody>
      </table>
     </div>
    </div>
    <!-- Action Bar -->
    <div class="bg-surface-subtle p-4 border-t border-border-default flex flex-col sm:flex-row items-center justify-between gap-4">
     <div class="flex gap-3 w-full sm:w-auto">
      <button class="flex-1 sm:flex-none flex items-center justify-center gap-2 px-4 py-2 border border-border-default bg-surface text-text-primary rounded-lg font-h3 text-h3 hover:bg-surface-variant hover:-translate-y-0.5 transition-all shadow-sm">
       <span class="material-symbols-outlined text-[18px]">
        picture_as_pdf
       </span>
       Export PDF
      </button>
      <button class="flex-1 sm:flex-none flex items-center justify-center gap-2 px-4 py-2 border border-border-default bg-surface text-text-primary rounded-lg font-h3 text-h3 hover:bg-surface-variant hover:-translate-y-0.5 transition-all shadow-sm">
       <span class="material-symbols-outlined text-[18px]">
        table_view
       </span>
       Export Excel
      </button>
     </div>
     <button class="w-full sm:w-auto flex items-center justify-center gap-2 px-6 py-2 bg-text-primary text-white rounded-lg font-h3 text-h3 hover:bg-black hover:-translate-y-0.5 transition-all shadow-sm">
      <span class="material-symbols-outlined text-[18px]">
       save
      </span>
      Save Report
     </button>
    </div>
   </div>
  </div>
 </div>
</main>

</template>

<script setup lang="ts">
// Autogenerated from custom_report_builder
</script>
