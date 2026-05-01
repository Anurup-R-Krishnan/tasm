<template>
  <header
    class="fixed top-0 left-0 w-full h-[60px] z-50 bg-[#FAFAF8] dark:bg-stone-950 border-b border-[#EEEBE4] dark:border-stone-800 shadow-sm flex justify-between items-center px-8 w-full font-['Inter'] antialiased tracking-tight"
  >
    <div class="flex items-center gap-inline">
      <span class="text-lg font-black tracking-tighter text-stone-900 dark:text-stone-50">
        Technopark AMS
      </span>
    </div>
    <div class="flex-1 flex justify-start pl-12">
      <div class="relative w-96">
        <span
          class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-stone-400"
        >
          search
        </span>
        <input
          class="w-full bg-surface border border-border-default rounded-full py-1.5 pl-10 pr-4 text-body font-body focus:outline-none focus:border-amber-700 transition-colors"
          placeholder="Search consumables..."
          type="text"
        />
      </div>
    </div>
    <div class="flex items-center gap-inline">
      <button
        class="p-2 text-stone-500 dark:text-stone-400 font-medium hover:bg-amber-50 dark:hover:bg-stone-900 transition-all duration-200 rounded-full active:opacity-80 transition-opacity flex items-center justify-center"
      >
        <span class="material-symbols-outlined"> notifications </span>
      </button>
      <button
        class="p-2 text-stone-500 dark:text-stone-400 font-medium hover:bg-amber-50 dark:hover:bg-stone-900 transition-all duration-200 rounded-full active:opacity-80 transition-opacity flex items-center justify-center"
      >
        <span class="material-symbols-outlined"> help </span>
      </button>
      <button
        class="p-2 text-stone-500 dark:text-stone-400 font-medium hover:bg-amber-50 dark:hover:bg-stone-900 transition-all duration-200 rounded-full active:opacity-80 transition-opacity flex items-center justify-center"
      >
        <span class="material-symbols-outlined"> settings </span>
      </button>
      <div
        class="w-8 h-8 rounded-full bg-surface-variant overflow-hidden ml-2 border border-border-default"
      >
        <img
          alt="Administrator Profile"
          class="w-full h-full object-cover"
          data-alt="Professional headshot of an administrator"
          src="https://lh3.googleusercontent.com/aida-public/AB6AXuA9rRpw9OWuaXJPeM8Xw6P4_pomgGYRQv2jHKtU4Ud3IrkSp4sa0Ph7JeIspRcJyGzQx8HAQYT_SmVvkcpSA72UDLp3XKXYaO4PECDU9PaqjKKz7O67QAgyvvQqbBSbQ_5hLGJsm9A7ybL3YvHziG_Z1mV0azn1_UV4WRwW285q5pyuFjfTOLZSaI2-Uq-3-mYSJ8UoUD2KXfH8G9_BDGEC9iehsCjyzfk9PB7dviNbRC_E2mf4cHNZvuttFsskTpue0zk_hsvv6sbN"
        />
      </div>
    </div>
  </header>
</template>

<script setup lang="ts"></script>
