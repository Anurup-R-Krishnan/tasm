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

    <RouterLink class="w-full mb-6 py-2 px-4 bg-amber-100 text-amber-900 rounded-lg font-semibold hover:bg-[#FEF3C7] dark:hover:bg-stone-900 transition-colors duration-200 active:translate-x-1 flex items-center justify-center gap-2" to="/add-new-asset-form">
      <span class="material-symbols-outlined">
        add
      </span>
      Add New Asset
    </RouterLink>

    <ul class="flex-1 space-y-1">
      <li v-for="item in primaryItems" :key="item.name">
        <RouterLink :class="navClass(item.name)" :to="item.to">
          <span class="material-symbols-outlined" :style="isActive(item.name) ? activeIconStyle : undefined">
            {{ item.icon }}
          </span>
          {{ item.title }}
        </RouterLink>
      </li>
    </ul>

    <div class="mt-auto pt-4 border-t border-border-default space-y-1">
      <RouterLink v-if="settingsItem" :class="navClass(settingsItem.name)" :to="settingsItem.to">
        <span class="material-symbols-outlined" :style="isActive(settingsItem.name) ? activeIconStyle : undefined">
          {{ settingsItem.icon }}
        </span>
        {{ settingsItem.title }}
      </RouterLink>
      <a class="flex items-center gap-3 px-3 py-2 rounded-lg text-stone-600 dark:text-stone-400 hover:text-amber-700 hover:bg-[#FEF3C7] dark:hover:bg-stone-900 transition-colors duration-200 group" href="mailto:support@technopark.com">
        <span class="material-symbols-outlined group-hover:text-amber-700">
          contact_support
        </span>
        Support
      </a>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { RouterLink, useRoute } from 'vue-router'
import { sidebarRoutes } from '../router/routes'

const route = useRoute()
const activeIconStyle = "font-variation-settings: 'FILL' 1;"

const navigationItems = sidebarRoutes.map((item) => ({
  ...item,
  to: item.path === '' ? '/' : `/${item.path}`,
}))

const settingsItem = navigationItems.find((item) => item.name === 'UserManagementSettings')
const primaryItems = navigationItems.filter((item) => item.name !== 'UserManagementSettings')

const activeName = computed(() => route.name?.toString() ?? '')

function isActive(name: string): boolean {
  return activeName.value === name
}

function navClass(name: string): string {
  return isActive(name)
    ? 'flex items-center gap-3 px-3 py-2.5 rounded-lg bg-amber-50 dark:bg-amber-900/20 text-amber-900 dark:text-amber-400 border-l-[3px] border-amber-600 font-semibold group'
    : 'flex items-center gap-3 px-3 py-2.5 rounded-lg text-stone-600 dark:text-stone-400 hover:text-amber-700 hover:bg-[#FEF3C7] dark:hover:bg-stone-900 transition-colors duration-200 group'
}
</script>
