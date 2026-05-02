<template>
  <header
    class="fixed top-0 left-0 w-full h-[60px] z-50 bg-surface border-b border-border-default shadow-sm flex justify-between items-center px-8 font-body antialiased tracking-tight"
  >
    <div class="flex items-center gap-inline">
      <span class="text-lg font-black tracking-tighter text-text-primary"> Technopark AMS </span>
      <span class="text-sm text-text-secondary">
        {{ currentTitle }}
      </span>
    </div>

    <div class="flex-1 flex justify-start pl-12">
      <div class="relative w-96">
        <span
          class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-text-secondary"
        >
          search
        </span>
        <input
          v-model="query"
          class="w-full bg-surface-subtle border border-border-default rounded-full py-1.5 pl-10 pr-4 text-sm focus:outline-none focus:border-primary/30 transition-all"
          placeholder="Search routes..."
          type="text"
        />
        <div
          v-if="filteredRoutes.length > 0"
          class="absolute left-0 right-0 top-[calc(100%+8px)] rounded-xl border border-border-default bg-surface shadow-lg overflow-hidden"
        >
          <RouterLink
            v-for="item in filteredRoutes"
            :key="item.name"
            :to="item.to"
            class="flex items-center gap-3 px-4 py-3 text-sm text-text-primary hover:bg-primary-container/10"
            @click="query = ''"
          >
            <span class="material-symbols-outlined text-text-secondary">
              {{ item.icon ?? 'arrow_outward' }}
            </span>
            {{ item.title }}
          </RouterLink>
        </div>
      </div>
    </div>

    <div class="flex items-center gap-inline">
      <button class="p-2 text-text-secondary hover:text-primary transition-colors">
        <span class="material-symbols-outlined">notifications</span>
      </button>
      <button class="p-2 text-text-secondary hover:text-primary transition-colors">
        <span class="material-symbols-outlined">help</span>
      </button>
      <RouterLink
        to="/user-management-settings"
        class="p-2 text-text-secondary hover:text-primary transition-colors"
      >
        <span class="material-symbols-outlined">settings</span>
      </RouterLink>
      <div
        class="w-8 h-8 rounded-full bg-surface-variant overflow-hidden ml-2 border border-border-default"
      >
        <img
          alt="Administrator Profile"
          class="w-full h-full object-cover"
          src="https://lh3.googleusercontent.com/aida-public/AB6AXuA9rRpw9OWuaXJPeM8Xw6P4_pomgGYRQv2jHKtU4Ud3IrkSp4sa0Ph7JeIspRcJyGzQx8HAQYT_SmVvkcpSA72UDLp3XKXYaO4PECDU9PaqjKKz7O67QAgyvvQqbBSbQ_5hLGJsm9A7ybL3YvHziG_Z1mV0azn1_UV4WRwW285q5pyuFjfTOLZSaI2-Uq-3-mYSJ8UoUD2KXfH8G9_BDGEC9iehsCjyzfk9PB7dviNbRC_E2mf4cHNZvuttFsskTpue0zk_hsvv6sbN"
        />
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
