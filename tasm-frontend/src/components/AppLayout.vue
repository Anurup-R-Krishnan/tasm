<template>
  <div class="min-h-screen bg-canvas">
    <AppHeader />
    <AppSidebar />
    <main class="ml-[248px] pt-[60px] p-[32px] max-w-[1400px] mx-auto space-y-[32px]">
      <router-view />
    </main>
  </div>
</template>

<script setup lang="ts">
import AppHeader from './AppHeader.vue'
import AppSidebar from './AppSidebar.vue'
</script>
