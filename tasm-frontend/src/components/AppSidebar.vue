<template>
  <nav
    class="fixed left-0 top-0 h-full w-[248px] z-40 bg-surface border-r border-border-default flex flex-col pt-[60px] pb-6 px-4 font-body text-sm font-medium"
  >
    <!-- Brand Info -->
    <div class="py-6 flex items-center gap-3">
      <div
        class="w-10 h-10 rounded-lg bg-white flex items-center justify-center shadow-sm border border-border-default"
      >
        <span class="material-symbols-outlined text-amber-700"> domain </span>
      </div>
      <div class="overflow-hidden">
        <h2 class="text-xl font-bold text-text-primary leading-tight truncate">Technopark</h2>
        <span class="text-text-secondary text-[10px] uppercase tracking-widest font-bold">
          Asset Management
        </span>
      </div>
    </div>

    <!-- Action Button -->
    <RouterLink
      class="w-full mb-6 py-2.5 px-4 bg-primary text-on-primary rounded-xl font-bold hover:opacity-90 transition-all duration-200 active:scale-95 flex items-center justify-center gap-2 shadow-lg shadow-primary/20"
      to="/add-asset"
    >
      <span class="material-symbols-outlined text-[20px]"> add </span>
      Add Asset
    </RouterLink>

    <!-- Navigation List with Groups -->
    <div class="flex-1 overflow-y-auto custom-scrollbar -mx-2 px-2 pb-12">
      <!-- Dashboard (Standalone) -->
      <RouterLink :class="navClass('Dashboard')" to="/">
        <span
          class="material-symbols-outlined"
          :style="isActive('Dashboard') ? activeIconStyle : undefined"
        >
          dashboard
        </span>
        Dashboard
      </RouterLink>

      <!-- Groups -->
      <div v-for="group in groupedNavigation" :key="group.name" class="mt-6">
        <h3 class="px-4 text-[10px] font-black text-text-disabled uppercase tracking-[0.2em] mb-2">
          {{ group.name }}
        </h3>
        <ul class="space-y-0.5">
          <li v-for="item in group.items" :key="item.name">
            <RouterLink :class="navClass(item.name)" :to="item.to">
              <span
                class="material-symbols-outlined text-[20px]"
                :style="isActive(item.name) ? activeIconStyle : undefined"
              >
                {{ item.icon }}
              </span>
              <span class="truncate">{{ item.title }}</span>
            </RouterLink>
          </li>
        </ul>
      </div>
    </div>

    <!-- Footer Settings -->
    <div class="mt-auto pt-4 border-t border-border-default space-y-1">
      <RouterLink v-if="settingsItem" :class="navClass(settingsItem.name)" :to="settingsItem.to">
        <span
          class="material-symbols-outlined"
          :style="isActive(settingsItem.name) ? activeIconStyle : undefined"
        >
          {{ settingsItem.icon }}
        </span>
        {{ settingsItem.title }}
      </RouterLink>
      <a
        class="flex items-center gap-3 px-3 py-2 rounded-lg text-text-secondary hover:text-primary hover:bg-primary-container/10 transition-colors duration-200 group"
        href="mailto:support@technopark.com"
      >
        <span class="material-symbols-outlined group-hover:text-primary"> contact_support </span>
        Support
      </a>
      <button
        class="w-full flex items-center gap-3 px-3 py-2 rounded-lg text-text-secondary hover:text-red-500 hover:bg-red-50 transition-colors duration-200 group"
        @click="handleLogout"
      >
        <span class="material-symbols-outlined group-hover:text-red-500">logout</span>
        Sign Out
      </button>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { RouterLink, useRoute } from 'vue-router';
import { appRoutes, sidebarGroups } from '../router/routes';
import { useAuth } from '../composables/useAuth';

const route = useRoute();
const activeIconStyle = "font-variation-settings: 'FILL' 1;";
const { logout } = useAuth();

const handleLogout = () => {
  logout();
};

const navigationItems = appRoutes.map((item) => ({
  ...item,
  to: item.path === '' ? '/' : `/${item.path}`,
}));

const settingsItem = navigationItems.find((item) => item.name === 'UserManagementSettings');

const groupedNavigation = computed(() => {
  return sidebarGroups.map((group) => ({
    name: group.name,
    items: group.items
      .map((name) => navigationItems.find((item) => item.name === name))
      .filter((item): item is typeof item & {} => Boolean(item)),
  }));
});

const activeName = computed(() => route.name?.toString() ?? '');

function isActive(name: string): boolean {
  return activeName.value === name;
}

function navClass(name: string): string {
  return isActive(name)
    ? 'flex items-center gap-3 px-3 py-2 rounded-xl bg-primary-container/10 text-primary border-l-[3px] border-primary font-bold group mb-0.5'
    : 'flex items-center gap-3 px-3 py-2 rounded-xl text-text-secondary hover:text-primary hover:bg-primary-container/10 transition-colors duration-200 group mb-0.5';
}
</script>

<style scoped>
.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background: var(--color-border-default);
  border-radius: 10px;
}
.custom-scrollbar::-webkit-scrollbar-track {
  background: transparent;
}
</style>
