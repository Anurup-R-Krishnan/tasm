<template>
<main class="flex-1 overflow-y-auto custom-scrollbar bg-canvas p-page-margin pb-32">
 <div class="max-w-[1000px] mx-auto">
  <div class="mb-8">
   <div class="flex items-center gap-2 text-stone-500 font-body text-body mb-2">
    <a class="hover:text-stone-800" href="#">
     Asset Inventory
    </a>
    <span class="material-symbols-outlined text-[16px]">
     chevron_right
    </span>
    <span class="text-stone-800">
     Add New Asset
    </span>
   </div>
   <h2 class="font-h1 text-h1 text-text-primary">
    Add New Asset
   </h2>
   <p class="font-body text-body text-text-secondary mt-1">
    Enter details to register a new asset into the system.
   </p>
  </div>
  <div class="bg-surface rounded-xl border border-border-default shadow-[0_4px_4px_rgba(0,0,0,0.02)] p-card-padding">
   <form>
    <!-- 1. Identity -->
    <div class="mb-section-gap">
     <h3 class="font-h3 text-h3 text-text-primary border-b border-border-default pb-2 mb-6 flex items-center gap-2">
      <span class="material-symbols-outlined text-amber-600 text-[20px]">
       badge
      </span>
      Identity
     </h3>
     <div class="grid grid-cols-1 md:grid-cols-2 gap-x-6 gap-y-5">
      <div class="col-span-1 md:col-span-2">
       <label class="block font-table-header text-table-header text-text-secondary mb-1.5">
        Asset Name (Description)
        <span class="text-error">
         *
        </span>
       </label>
       <input class="w-full border border-border-default rounded-[6px] px-3 py-2 font-body text-body focus:outline-none focus:border-amber-500 focus:ring-1 focus:ring-amber-500" placeholder="e.g. Dell Latitude 5420 Laptop" type="text"/>
      </div>
      <div>
       <label class="block font-table-header text-table-header text-text-secondary mb-1.5 flex justify-between">
        <span>
         Tag ID
        </span>
        <span class="font-metadata text-metadata text-amber-600 cursor-pointer hover:underline">
         Auto-generate
        </span>
       </label>
       <input class="w-full border border-border-default rounded-[6px] px-3 py-2 font-mono-data text-mono-data bg-stone-50 focus:outline-none focus:border-amber-500 focus:ring-1 focus:ring-amber-500" placeholder="AST-2024-0001" type="text"/>
      </div>
      <div>
       <label class="block font-table-header text-table-header text-text-secondary mb-1.5">
        Category
        <span class="text-error">
         *
        </span>
       </label>
       <div class="relative">
        <select class="w-full border border-border-default rounded-[6px] px-3 py-2 font-body text-body appearance-none focus:outline-none focus:border-amber-500 focus:ring-1 focus:ring-amber-500 bg-white pr-8">
         <option disabled="" selected="" value="">
          Select Category
         </option>
         <option>
          IT Equipment
         </option>
         <option>
          Furniture
         </option>
         <option>
          Vehicles
         </option>
         <option>
          Machinery
         </option>
        </select>
        <span class="material-symbols-outlined absolute right-2 top-1/2 -translate-y-1/2 pointer-events-none text-stone-400">
         expand_more
        </span>
       </div>
      </div>
      <div>
       <label class="block font-table-header text-table-header text-text-secondary mb-1.5">
        Sub-category
       </label>
       <div class="relative">
        <select class="w-full border border-border-default rounded-[6px] px-3 py-2 font-body text-body appearance-none focus:outline-none focus:border-amber-500 focus:ring-1 focus:ring-amber-500 bg-white pr-8">
         <option disabled="" selected="" value="">
          Select Sub-category
         </option>
         <option>
          Laptops
         </option>
         <option>
          Desktops
         </option>
         <option>
          Monitors
         </option>
        </select>
        <span class="material-symbols-outlined absolute right-2 top-1/2 -translate-y-1/2 pointer-events-none text-stone-400">
         expand_more
        </span>
       </div>
      </div>
      <div>
       <label class="block font-table-header text-table-header text-text-secondary mb-1.5">
        Brand
       </label>
       <input class="w-full border border-border-default rounded-[6px] px-3 py-2 font-body text-body focus:outline-none focus:border-amber-500 focus:ring-1 focus:ring-amber-500" placeholder="e.g. Dell" type="text"/>
      </div>
      <div>
       <label class="block font-table-header text-table-header text-text-secondary mb-1.5">
        Model
       </label>
       <input class="w-full border border-border-default rounded-[6px] px-3 py-2 font-body text-body focus:outline-none focus:border-amber-500 focus:ring-1 focus:ring-amber-500" placeholder="e.g. Latitude 5420" type="text"/>
      </div>
      <div>
       <label class="block font-table-header text-table-header text-text-secondary mb-1.5">
        Serial Number
       </label>
       <input class="w-full border border-border-default rounded-[6px] px-3 py-2 font-mono-data text-mono-data focus:outline-none focus:border-amber-500 focus:ring-1 focus:ring-amber-500" placeholder="SN-XXXX-YYYY" type="text"/>
      </div>
      <div>
       <label class="block font-table-header text-table-header text-text-secondary mb-1.5">
        Condition
        <span class="text-error">
         *
        </span>
       </label>
       <div class="relative">
        <select class="w-full border border-border-default rounded-[6px] px-3 py-2 font-body text-body appearance-none focus:outline-none focus:border-amber-500 focus:ring-1 focus:ring-amber-500 bg-white pr-8">
         <option>
          New
         </option>
         <option>
          Good
         </option>
         <option>
          Fair
         </option>
         <option>
          Poor
         </option>
        </select>
        <span class="material-symbols-outlined absolute right-2 top-1/2 -translate-y-1/2 pointer-events-none text-stone-400">
         expand_more
        </span>
       </div>
      </div>
     </div>
    </div>
    <!-- 2. Purchase & Finance -->
    <div class="mb-section-gap">
     <h3 class="font-h3 text-h3 text-text-primary border-b border-border-default pb-2 mb-6 flex items-center gap-2">
      <span class="material-symbols-outlined text-amber-600 text-[20px]">
       payments
      </span>
      Purchase &amp; Finance
     </h3>
     <div class="grid grid-cols-1 md:grid-cols-2 gap-x-6 gap-y-5">
      <div>
       <label class="block font-table-header text-table-header text-text-secondary mb-1.5">
        Purchase Date
       </label>
       <div class="relative">
        <input class="w-full border border-border-default rounded-[6px] px-3 py-2 font-body text-body focus:outline-none focus:border-amber-500 focus:ring-1 focus:ring-amber-500" type="date"/>
       </div>
      </div>
      <div>
       <label class="block font-table-header text-table-header text-text-secondary mb-1.5">
        Purchase Cost
       </label>
       <div class="relative">
        <span class="absolute left-3 top-1/2 -translate-y-1/2 text-stone-500 font-body text-body">
         ₹
        </span>
        <input class="w-full border border-border-default rounded-[6px] pl-7 pr-3 py-2 font-mono-data text-mono-data focus:outline-none focus:border-amber-500 focus:ring-1 focus:ring-amber-500" placeholder="0.00" type="text"/>
       </div>
      </div>
      <div>
       <label class="block font-table-header text-table-header text-text-secondary mb-1.5">
        Vendor
       </label>
       <div class="relative">
        <select class="w-full border border-border-default rounded-[6px] px-3 py-2 font-body text-body appearance-none focus:outline-none focus:border-amber-500 focus:ring-1 focus:ring-amber-500 bg-white pr-8">
         <option disabled="" selected="" value="">
          Select Vendor
         </option>
         <option>
          TechCorp Solutions
         </option>
         <option>
          Global Furnishings
         </option>
         <option>
          AutoMax
         </option>
        </select>
        <span class="material-symbols-outlined absolute right-2 top-1/2 -translate-y-1/2 pointer-events-none text-stone-400">
         expand_more
        </span>
       </div>
      </div>
      <div>
       <label class="block font-table-header text-table-header text-text-secondary mb-1.5">
        PO Number
       </label>
       <input class="w-full border border-border-default rounded-[6px] px-3 py-2 font-body text-body focus:outline-none focus:border-amber-500 focus:ring-1 focus:ring-amber-500" placeholder="PO-2024-XXXX" type="text"/>
      </div>
      <div>
       <label class="block font-table-header text-table-header text-text-secondary mb-1.5">
        Funding Source
       </label>
       <div class="relative">
        <select class="w-full border border-border-default rounded-[6px] px-3 py-2 font-body text-body appearance-none focus:outline-none focus:border-amber-500 focus:ring-1 focus:ring-amber-500 bg-white pr-8">
         <option>
          Capital Budget
         </option>
         <option>
          Operational Budget
         </option>
         <option>
          Grant
         </option>
        </select>
        <span class="material-symbols-outlined absolute right-2 top-1/2 -translate-y-1/2 pointer-events-none text-stone-400">
         expand_more
        </span>
       </div>
      </div>
     </div>
    </div>
   </form>
  </div>
 </div>
</main>

</template>

<script setup lang="ts">
// Autogenerated from add_new_asset_form
</script>
