<template>
  <Teleport to="body">
    <div class="toast-container fixed bottom-6 right-6 z-[9999] flex flex-col gap-3">
      <TransitionGroup name="toast">
        <div
          v-for="toast in toasts"
          :key="toast.id"
          class="toast-item px-4 py-3 rounded-lg shadow-lg flex items-center gap-3 min-w-[280px] max-w-md"
          :class="toastClasses[toast.type]"
        >
          <span class="material-symbols-outlined text-lg">{{ toastIcons[toast.type] }}</span>
          <div class="flex-1">
            <p class="text-sm font-medium">{{ toast.message }}</p>
          </div>
          <button class="p-1 rounded hover:bg-black/10" @click="removeToast(toast.id)">
            <span class="material-symbols-outlined text-lg">close</span>
          </button>
        </div>
      </TransitionGroup>
    </div>
  </Teleport>
</template>

<script setup lang="ts">
import { ref } from 'vue';

type ToastType = 'success' | 'error' | 'warning' | 'info';

interface Toast {
  id: number;
  message: string;
  type: ToastType;
}

const toasts = ref<Toast[]>([]);

const toastClasses: Record<ToastType, string> = {
  success: 'bg-status-in-stock text-white',
  error: 'bg-status-critical text-white',
  warning: 'bg-metric-amber text-white',
  info: 'bg-primary text-white',
};

const toastIcons: Record<ToastType, string> = {
  success: 'check_circle',
  error: 'error',
  warning: 'warning',
  info: 'info',
};

let toastId = 0;

const showToast = (message: string, type: ToastType = 'info', duration = 4000) => {
  const id = ++toastId;
  toasts.value.push({ id, message, type });

  if (duration > 0) {
    setTimeout(() => {
      removeToast(id);
    }, duration);
  }
};

const removeToast = (id: number) => {
  toasts.value = toasts.value.filter((t) => t.id !== id);
};

const toast = {
  success: (message: string) => showToast(message, 'success'),
  error: (message: string) => showToast(message, 'error'),
  warning: (message: string) => showToast(message, 'warning'),
  info: (message: string) => showToast(message, 'info'),
};

defineExpose({ toast });
</script>

<style scoped>
.toast-container {
  pointer-events: none;
}
.toast-item {
  pointer-events: auto;
}
.toast-enter-active,
.toast-leave-active {
  transition: all 0.3s ease;
}
.toast-enter-from {
  opacity: 0;
  transform: translateX(100%);
}
.toast-leave-to {
  opacity: 0;
  transform: translateX(100%);
}
</style>
