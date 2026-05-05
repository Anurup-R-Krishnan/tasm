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
          @focus="isFocused = true"
          @blur="isFocused = false"
        />
        <div
          v-if="filteredRoutes.length > 0 && isFocused && (query.length > 0 || isFocused)"
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
      <RouterLink
        to="/audit-scan"
        class="p-2 text-text-secondary hover:text-primary transition-colors flex items-center"
        title="Scan Asset QR"
      >
        <span class="material-symbols-outlined">qr_code_scanner</span>
      </RouterLink>
      <button class="p-2 text-text-secondary hover:text-primary transition-colors">
        <span class="material-symbols-outlined">notifications</span>
      </button>
      <button class="p-2 text-text-secondary hover:text-primary transition-colors">
        <span class="material-symbols-outlined">help</span>
      </button>
      <RouterLink
        to="/settings"
        class="p-2 text-text-secondary hover:text-primary transition-colors"
      >
        <span class="material-symbols-outlined">settings</span>
      </RouterLink>
      <div class="flex items-center gap-2 ml-2 pl-3 border-l border-border-default">
        <div
          class="w-8 h-8 rounded-full bg-primary-container flex items-center justify-center text-primary font-bold text-sm"
        >
          {{ userInitials }}
        </div>
        <div v-if="currentUser" class="hidden lg:block">
          <p class="text-xs font-semibold text-text-primary leading-tight">
            {{ currentUser.name }}
          </p>
          <p class="text-[10px] text-text-secondary leading-tight">{{ currentUser.role }}</p>
        </div>
      </div>
    </div>
  </header>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue';
import { RouterLink, useRoute } from 'vue-router';
import { appRoutes } from '../router/routes';
import { useAuth } from '../composables/useAuth';

const route = useRoute();
const query = ref('');
const isFocused = ref(false);
const { currentUser } = useAuth();

const currentTitle = computed(() => route.meta['title']?.toString() ?? 'Dashboard');
const routeIndex = appRoutes.map((item) => ({
  ...item,
  to: item.path === '' ? '/' : `/${item.path}`,
}));

const userInitials = computed(() => {
  if (!currentUser.value?.name) return '?';
  return currentUser.value.name
    .split(' ')
    .map((n: string) => n[0])
    .join('')
    .toUpperCase()
    .slice(0, 2);
});

const filteredRoutes = computed(() => {
  if (!isFocused.value && !query.value) return [];
  const normalizedQuery = query.value.trim().toLowerCase();
  if (normalizedQuery.length < 2 && !isFocused.value) return [];
  return routeIndex
    .filter((item) => item.title.toLowerCase().includes(normalizedQuery))
    .slice(0, 6);
});
</script>
