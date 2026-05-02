<template>
  <main class="flex-1 md:ml-[248px] p-page-margin w-full max-w-[1400px] mx-auto">
    <!-- Page Header -->
    <div
      class="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4 mb-section-gap"
    >
      <div>
        <h1 class="font-h1 text-h1 text-text-primary">System Alerts</h1>
        <p class="font-metadata text-metadata text-text-secondary mt-1">
          Manage and resolve system notifications and warnings.
        </p>
      </div>
      <div class="flex items-center gap-3">
        <button
          class="bg-surface border border-border-default text-text-primary px-4 py-2 rounded-lg font-h3 text-h3 hover:bg-surface-subtle hover:-translate-y-[2px] transition-all duration-200 shadow-[0_4px_4px_rgba(0,0,0,0.02)]"
        >
          Mark All Read
        </button>
        <button
          class="bg-primary text-on-primary px-4 py-2 rounded-lg font-h3 text-h3 hover:bg-primary-hover hover:-translate-y-[2px] transition-all duration-200 flex items-center gap-2 shadow-[0_4px_4px_rgba(0,0,0,0.02)]"
        >
          <span class="material-symbols-outlined text-[18px]"> settings </span>
          Alert Rules
        </button>
      </div>
    </div>
    <!-- Filter Bar -->
    <div
      class="bg-surface rounded-xl border border-border-default shadow-[0_4px_4px_rgba(0,0,0,0.02)] p-2 mb-stack flex overflow-x-auto hide-scrollbar"
    >
      <button
        @click="selectedFilter = 'All'"
        :class="
          selectedFilter === 'All'
            ? 'bg-surface-subtle text-text-primary'
            : 'text-text-secondary hover:text-text-primary hover:bg-surface-subtle'
        "
        class="px-4 py-1.5 rounded-lg font-h3 text-h3 whitespace-nowrap transition-colors"
      >
        All Alerts ({{ counts.All }})
      </button>
      <button
        @click="selectedFilter = 'Critical'"
        :class="
          selectedFilter === 'Critical'
            ? 'bg-surface-subtle text-text-primary'
            : 'text-text-secondary hover:text-text-primary hover:bg-surface-subtle'
        "
        class="px-4 py-1.5 rounded-lg font-h3 text-h3 whitespace-nowrap transition-colors flex items-center gap-2"
      >
        <span class="w-2 h-2 rounded-full bg-status-critical"></span>
        Critical ({{ counts.Critical }})
      </button>
      <button
        @click="selectedFilter = 'Warning'"
        :class="
          selectedFilter === 'Warning'
            ? 'bg-surface-subtle text-text-primary'
            : 'text-text-secondary hover:text-text-primary hover:bg-surface-subtle'
        "
        class="px-4 py-1.5 rounded-lg font-h3 text-h3 whitespace-nowrap transition-colors flex items-center gap-2"
      >
        <span class="w-2 h-2 rounded-full bg-primary-container"></span>
        Warning ({{ counts.Warning }})
      </button>
      <button
        @click="selectedFilter = 'Info'"
        :class="
          selectedFilter === 'Info'
            ? 'bg-surface-subtle text-text-primary'
            : 'text-text-secondary hover:text-text-primary hover:bg-surface-subtle'
        "
        class="px-4 py-1.5 rounded-lg font-h3 text-h3 whitespace-nowrap transition-colors flex items-center gap-2"
      >
        <span class="w-2 h-2 rounded-full bg-status-checked-out"></span>
        Info ({{ counts.Info }})
      </button>
      <div class="ml-auto flex items-center gap-2 pl-4 border-l border-border-default">
        <button
          class="p-1.5 text-text-secondary hover:text-text-primary rounded-lg hover:bg-surface-subtle transition-colors"
        >
          <span class="material-symbols-outlined text-[20px]"> filter_list </span>
        </button>
      </div>
    </div>
    <!-- Alerts List -->
    <div class="space-y-inline" v-if="filteredAlerts.length > 0">
      <!-- Dynamic Alert Card -->
      <div
        v-for="alert in filteredAlerts"
        :key="alert.id"
        class="bg-surface rounded-xl border border-border-default shadow-[0_4px_4px_rgba(0,0,0,0.02)] hover:-translate-y-[2px] transition-transform duration-200 overflow-hidden flex flex-col sm:flex-row relative"
        :class="alert.isRead ? 'opacity-75 hover:opacity-100' : ''"
      >
        <div
          class="absolute left-0 top-0 bottom-0 w-[3px]"
          :class="
            alert.type === 'Critical'
              ? 'bg-status-critical'
              : alert.type === 'Warning'
                ? 'bg-primary-container'
                : 'bg-status-checked-out'
          "
        ></div>
        <div class="p-card-padding flex-1 flex flex-col sm:flex-row gap-4 sm:items-center pl-6">
          <div
            class="flex-shrink-0 w-10 h-10 rounded-full flex items-center justify-center"
            :class="
              alert.type === 'Critical'
                ? 'bg-error-container/20 text-status-critical border border-error-container/30'
                : alert.type === 'Warning'
                  ? 'bg-metric-amber/20 text-surface-tint border border-metric-amber/30'
                  : 'bg-primary-container/10 text-primary border border-primary-container/20'
            "
          >
            <span
              class="material-symbols-outlined text-[20px]"
              style="font-variation-settings: 'FILL' 1"
            >
              {{
                alert.type === 'Critical' ? 'error' : alert.type === 'Warning' ? 'warning' : 'info'
              }}
            </span>
          </div>
          <div class="flex-1">
            <div class="flex items-center gap-2 mb-1">
              <span
                class="px-2 py-0.5 rounded text-[11px] font-semibold tracking-wide uppercase"
                :class="
                  alert.type === 'Critical'
                    ? 'bg-error-container text-status-critical'
                    : alert.type === 'Warning'
                      ? 'bg-metric-amber text-primary-container'
                      : 'bg-tertiary-fixed text-status-checked-out'
                "
              >
                {{ alert.title }}
              </span>
              <span class="font-metadata text-metadata text-text-secondary">
                {{ new Date(alert.createdAt).toLocaleString() }}
              </span>
            </div>
            <p class="font-body text-body text-text-primary">
              {{ alert.message }}
            </p>
          </div>
          <div class="flex-shrink-0 sm:ml-4 mt-4 sm:mt-0 flex gap-2">
            <button
              class="bg-surface border border-border-default text-text-primary px-3 py-1.5 rounded-lg font-h3 text-[13px] hover:bg-surface-subtle transition-colors"
            >
              Dismiss
            </button>
            <button
              v-if="alert.type === 'Critical' || alert.type === 'Warning'"
              :class="
                alert.type === 'Critical'
                  ? 'bg-primary-container text-white hover:bg-primary'
                  : 'bg-surface border border-border-default text-text-primary hover:bg-surface-subtle'
              "
              class="px-3 py-1.5 rounded-lg font-h3 text-[13px] transition-colors"
            >
              {{ alert.type === 'Critical' ? 'Take Action' : 'Schedule' }}
            </button>
          </div>
        </div>
      </div>
    </div>
    <!-- Empty State (Hidden by default in this mockup) -->
    <div
      class="hidden bg-surface rounded-xl border border-border-default p-12 flex flex-col items-center justify-center text-center mt-stack"
    >
      <div class="w-16 h-16 rounded-full bg-surface-subtle flex items-center justify-center mb-4">
        <span class="material-symbols-outlined text-[32px] text-text-secondary">
          check_circle
        </span>
      </div>
      <h3 class="font-h2 text-h2 text-text-primary mb-2">No alerts found</h3>
      <p class="font-body text-body text-text-secondary max-w-md">
        You're all caught up! There are currently no critical system alerts or pending
        notifications.
      </p>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';

interface SystemAlert {
  id: number;
  title: string;
  message: string;
  type: string;
  source: string;
  isRead: boolean;
  createdAt: string;
}

const alerts = ref<SystemAlert[]>([]);
const loading = ref(true);
const selectedFilter = ref('All');

const filteredAlerts = computed(() => {
  if (selectedFilter.value === 'All') return alerts.value;
  return alerts.value.filter((a) => a.type === selectedFilter.value);
});

const counts = computed(() => {
  return {
    All: alerts.value.length,
    Critical: alerts.value.filter((a) => a.type === 'Critical').length,
    Warning: alerts.value.filter((a) => a.type === 'Warning').length,
    Info: alerts.value.filter((a) => a.type === 'Info').length,
  };
});

const fetchAlerts = async () => {
  try {
    const res = await fetch('http://localhost:8080/api/alerts');
    if (res.ok) {
      alerts.value = await res.json();
    }
  } catch (error) {
    console.error('Failed to fetch alerts:', error);
  } finally {
    loading.value = false;
  }
};

onMounted(() => {
  fetchAlerts();
});
</script>
