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
    <ToastContainer ref="toastContainer" />
  </div>
</template>

<script setup lang="ts">
import { ref, provide } from 'vue';
import AppHeader from './AppHeader.vue';
import AppSidebar from './AppSidebar.vue';
import ToastContainer from './ToastContainer.vue';

const toastContainer = ref<InstanceType<typeof ToastContainer> | null>(null);

provide('toast', {
  success: (msg: string) => toastContainer.value?.toast.success(msg),
  error: (msg: string) => toastContainer.value?.toast.error(msg),
  warning: (msg: string) => toastContainer.value?.toast.warning(msg),
  info: (msg: string) => toastContainer.value?.toast.info(msg),
});
</script>
