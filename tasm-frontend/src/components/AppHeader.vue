<template>
  <header
    class="fixed top-0 left-0 w-full h-[60px] z-50 glass border-b border-border-default flex justify-between items-center px-8 font-body"
  >
    <!-- Left Section: Title -->
    <div class="flex items-center gap-4">
      <div class="md:hidden">
        <span class="material-symbols-outlined text-slate-600">menu</span>
      </div>
      <div class="flex items-baseline gap-2">
        <span class="text-sm font-semibold text-slate-400 uppercase tracking-widest">TASM</span>
        <span class="text-lg font-bold text-slate-900 tracking-tight">
          {{ currentTitle }}
        </span>
      </div>
    </div>

    <!-- Center Section: Command Palette / Search -->
    <div class="hidden md:flex flex-1 justify-center max-w-2xl px-12">
      <div class="relative w-full group">
        <span
          class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-slate-400 group-focus-within:text-indigo-500 transition-colors"
        >
          search
        </span>
        <input
          v-model="query"
          class="w-full bg-slate-100 border-none rounded-xl py-2 pl-10 pr-4 text-sm focus:bg-white focus:ring-2 focus:ring-indigo-500/20 transition-all outline-none placeholder:text-slate-400"
          placeholder="Quick search (Ctrl + K)"
          type="text"
        />
        <!-- Results Dropdown -->
        <div
          v-if="filteredRoutes.length > 0"
          class="absolute left-0 right-0 top-[calc(100%+8px)] rounded-2xl border border-border-default bg-white shadow-2xl shadow-slate-200/50 overflow-hidden z-50 animate-in fade-in slide-in-from-top-2 duration-200"
        >
          <div
            class="p-2 text-[10px] font-bold text-slate-400 uppercase tracking-widest bg-slate-50 px-4 py-1.5"
          >
            Suggested Pages
          </div>
          <RouterLink
            v-for="item in filteredRoutes"
            :key="item.name"
            :to="item.to"
            class="flex items-center gap-3 px-4 py-3 text-sm text-slate-700 hover:bg-indigo-50 hover:text-indigo-600 transition-colors"
            @click="query = ''"
          >
            <span class="material-symbols-outlined text-slate-400 group-hover:text-indigo-500">
              {{ item.icon ?? 'arrow_outward' }}
            </span>
            {{ item.title }}
          </RouterLink>
        </div>
      </div>
    </div>

    <!-- Right Section: Actions & Profile -->
    <div class="flex items-center gap-4">
      <button class="p-2 rounded-lg text-slate-500 hover:bg-slate-100 transition-colors relative">
        <span class="material-symbols-outlined">notifications</span>
        <span
          class="absolute top-2 right-2 w-2 h-2 bg-rose-500 rounded-full border-2 border-white"
        ></span>
      </button>

      <div class="h-6 w-[1px] bg-slate-200 mx-1"></div>

      <div class="flex items-center gap-3 pl-2">
        <div class="text-right hidden sm:block">
          <p class="text-xs font-bold text-slate-900 leading-none">Admin User</p>
          <p class="text-[10px] text-slate-500 font-medium mt-1">Super Administrator</p>
        </div>
        <div
          class="w-9 h-9 rounded-xl bg-indigo-100 border-2 border-white shadow-sm overflow-hidden flex-shrink-0 cursor-pointer hover:ring-2 hover:ring-indigo-500/20 transition-all"
        >
          <img
            alt="Administrator Profile"
            class="w-full h-full object-cover"
            src="https://lh3.googleusercontent.com/aida-public/AB6AXuA9rRpw9OWuaXJPeM8Xw6P4_pomgGYRQv2jHKtU4Ud3IrkSp4sa0Ph7JeIspRcJyGzQx8HAQYT_SmVvkcpSA72UDLp3XKXYaO4PECDU9PaqjKKz7O67QAgyvvQqbBSbQ_5hLGJsm9A7ybL3YvHziG_Z1mV0azn1_UV4WRwW285q5pyuFjfTOLZSaI2-Uq-3-mYSJ8UoUD2KXfH8G9_BDGEC9iehsCjyzfk9PB7dviNbRC_E2mf4cHNZvuttFsskTpue0zk_hsvv6sbN"
          />
        </div>
      </div>
    </div>
  </header>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue';
import { RouterLink, useRoute } from 'vue-router';
import { appRoutes } from '../router/routes';

const route = useRoute();
const query = ref('');

const currentTitle = computed(() => route.meta['title']?.toString() ?? 'Dashboard');
const routeIndex = appRoutes.map((item) => ({
  ...item,
  to: item.path === '' ? '/' : `/${item.path}`,
}));

const filteredRoutes = computed(() => {
  const normalizedQuery = query.value.trim().toLowerCase();
  if (normalizedQuery.length < 2) return [];
  return routeIndex
    .filter((item) => item.title.toLowerCase().includes(normalizedQuery))
    .slice(0, 6);
});
</script>
