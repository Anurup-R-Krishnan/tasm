<template>
  <div class="min-h-screen bg-canvas">
    <AppHeader />
    <AppSidebar />
    <!-- Single source of layout truth -->
    <div class="md:pl-[248px] pt-[60px] min-h-[calc(100vh-60px)]">
      <div class="p-page-margin max-w-[1600px] w-full">
        <router-view />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import AppHeader from './AppHeader.vue';
import AppSidebar from './AppSidebar.vue';
</script>
