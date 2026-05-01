<template>
  <nav
    class="fixed left-0 top-0 h-full w-[248px] z-40 bg-slate-900 text-slate-300 border-r border-slate-800 flex flex-col pt-[60px] pb-6 px-4 font-body"
  >
    <!-- Brand / Logo Area -->
    <div class="py-8 px-2 flex items-center gap-3">
      <div
        class="w-10 h-10 rounded-xl bg-indigo-500 flex items-center justify-center shadow-lg shadow-indigo-500/20"
      >
        <span class="material-symbols-outlined text-white text-2xl"> dashboard_customize </span>
      </div>
      <div class="overflow-hidden">
        <h2 class="text-lg font-bold text-white leading-tight truncate">Technopark</h2>
        <span class="text-slate-500 text-xs font-medium tracking-wider uppercase">
          Asset Manager
        </span>
      </div>
    </div>

    <!-- Action Button -->
    <RouterLink
      class="w-full mb-8 py-2.5 px-4 bg-indigo-600 text-white rounded-xl font-semibold hover:bg-indigo-500 transition-all duration-200 shadow-lg shadow-indigo-600/20 flex items-center justify-center gap-2 active:scale-95"
      to="/add-new-asset-form"
    >
      <span class="material-symbols-outlined text-[20px]"> add_circle </span>
      Add Asset
    </RouterLink>

    <!-- Navigation List -->
    <div class="flex-1 overflow-y-auto custom-scrollbar -mx-2 px-2">
      <ul class="space-y-1.5">
        <li v-for="item in primaryItems" :key="item.name">
          <RouterLink :class="navClass(item.name)" :to="item.to">
            <span
              class="material-symbols-outlined text-[22px]"
              :style="isActive(item.name) ? activeIconStyle : undefined"
            >
              {{ item.icon }}
            </span>
            <span class="font-medium">{{ item.title }}</span>
          </RouterLink>
        </li>
      </ul>
    </div>

    <!-- Footer Links -->
    <div class="mt-auto pt-6 border-t border-slate-800 space-y-1.5">
      <RouterLink v-if="settingsItem" :class="navClass(settingsItem.name)" :to="settingsItem.to">
        <span
          class="material-symbols-outlined text-[22px]"
          :style="isActive(settingsItem.name) ? activeIconStyle : undefined"
        >
          {{ settingsItem.icon }}
        </span>
        <span class="font-medium">{{ settingsItem.title }}</span>
      </RouterLink>
      <a
        class="flex items-center gap-3 px-4 py-2.5 rounded-xl text-slate-400 hover:text-white hover:bg-slate-800 transition-all duration-200 group"
        href="mailto:support@technopark.com"
      >
        <span class="material-symbols-outlined text-[22px]"> contact_support </span>
        <span class="font-medium">Support</span>
      </a>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { RouterLink, useRoute } from 'vue-router';
import { sidebarRoutes } from '../router/routes';

const route = useRoute();
const activeIconStyle = "font-variation-settings: 'FILL' 1;";

const navigationItems = sidebarRoutes.map((item) => ({
  ...item,
  to: item.path === '' ? '/' : `/${item.path}`,
}));

const settingsItem = navigationItems.find((item) => item.name === 'UserManagementSettings');
const primaryItems = navigationItems.filter((item) => item.name !== 'UserManagementSettings');

const activeName = computed(() => route.name?.toString() ?? '');

function isActive(name: string): boolean {
  return activeName.value === name;
}

function navClass(name: string): string {
  return isActive(name)
    ? 'flex items-center gap-3 px-4 py-2.5 rounded-xl bg-indigo-600/10 text-indigo-400 border-r-4 border-indigo-500 font-bold group transition-all duration-200'
    : 'flex items-center gap-3 px-4 py-2.5 rounded-xl text-slate-400 hover:text-white hover:bg-slate-800 transition-all duration-200 group';
}
</script>

<style scoped>
.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background: #334155;
  border-radius: 10px;
}
</style>
