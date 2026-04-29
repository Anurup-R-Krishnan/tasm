<template>
<nav class="fixed left-0 top-0 h-full w-[248px] z-40 bg-[#FAFAF8] dark:bg-stone-950 border-r border-[#EEEBE4] dark:border-stone-800 flex flex-col pt-[60px] pb-6 px-4 font-['Inter'] text-sm font-medium">
 <div class="py-6 flex items-center gap-3">
  <div class="w-10 h-10 rounded-lg bg-surface flex items-center justify-center shadow-sm border border-border-default">
   <span class="material-symbols-outlined text-amber-700">
    domain
   </span>
  </div>
  <div>
   <h2 class="text-xl font-bold text-stone-900 dark:text-stone-50 leading-tight">
    Technopark Kerala
   </h2>
   <span class="text-stone-500 text-xs">
    Asset Management
   </span>
  </div>
 </div>
 <button class="w-full mb-6 py-2 px-4 bg-amber-100 text-amber-900 rounded-lg font-semibold hover:bg-[#FEF3C7] dark:hover:bg-stone-900 transition-colors duration-200 active:translate-x-1 duration-200 flex items-center justify-center gap-2">
  <span class="material-symbols-outlined">
   add
  </span>
  Add New Asset
 </button>
 <ul class="flex-1 space-y-1">
  <li>
   <a class="flex items-center gap-3 px-3 py-2.5 rounded-lg text-stone-600 dark:text-stone-400 hover:text-amber-700 hover:bg-[#FEF3C7] dark:hover:bg-stone-900 transition-colors duration-200 group" href="#">
    <span class="material-symbols-outlined group-hover:text-amber-700">
     dashboard
    </span>
    Dashboard
   </a>
  </li>
  <li>
   <a class="flex items-center gap-3 px-3 py-2.5 rounded-lg text-stone-600 dark:text-stone-400 hover:text-amber-700 hover:bg-[#FEF3C7] dark:hover:bg-stone-900 transition-colors duration-200 group" href="#">
    <span class="material-symbols-outlined group-hover:text-amber-700">
     inventory_2
    </span>
    Stockroom
   </a>
  </li>
  <li>
   <a class="flex items-center gap-3 px-3 py-2.5 rounded-lg bg-amber-50 dark:bg-amber-900/20 text-amber-900 dark:text-amber-400 border-l-[3px] border-amber-600 font-semibold group" href="#">
    <span class="material-symbols-outlined" style="font-variation-settings: 'FILL' 1;">
     inventory
    </span>
    Consumables
   </a>
  </li>
  <li>
   <a class="flex items-center gap-3 px-3 py-2.5 rounded-lg text-stone-600 dark:text-stone-400 hover:text-amber-700 hover:bg-[#FEF3C7] dark:hover:bg-stone-900 transition-colors duration-200 group" href="#">
    <span class="material-symbols-outlined group-hover:text-amber-700">
     fact_check
    </span>
    Audit Sessions
   </a>
  </li>
 </ul>
 <div class="mt-auto pt-4 border-t border-border-default space-y-1">
  <a class="flex items-center gap-3 px-3 py-2 rounded-lg text-stone-600 dark:text-stone-400 hover:text-amber-700 hover:bg-[#FEF3C7] dark:hover:bg-stone-900 transition-colors duration-200 group" href="#">
   <span class="material-symbols-outlined group-hover:text-amber-700">
    settings
   </span>
   Settings
  </a>
  <a class="flex items-center gap-3 px-3 py-2 rounded-lg text-stone-600 dark:text-stone-400 hover:text-amber-700 hover:bg-[#FEF3C7] dark:hover:bg-stone-900 transition-colors duration-200 group" href="#">
   <span class="material-symbols-outlined group-hover:text-amber-700">
    contact_support
   </span>
   Support
  </a>
 </div>
</nav>

</template>

<script setup lang="ts">
</script>
