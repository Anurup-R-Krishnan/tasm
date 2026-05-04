<template>
  <main class="space-y-section-gap pb-24">
    <!-- Hero: Massive Scan Button -->
    <section class="flex flex-col items-center justify-center py-section-gap">
      <button
        @click="handleScan"
        class="w-full max-w-sm h-[72px] bg-text-primary text-surface rounded-xl shadow-[0_4px_12px_rgba(0,0,0,0.1)] flex items-center justify-center gap-inline active:translate-y-[2px] transition-transform"
      >
        <span
          class="material-symbols-outlined text-[32px]"
          style="font-variation-settings: 'FILL' 1"
        >
          qr_code_scanner
        </span>
        <span class="font-h2 text-h2 uppercase tracking-wide"> Scan Asset </span>
      </button>
      <button
        @click="handleManualEntry"
        class="mt-stack font-body text-body text-text-secondary underline decoration-border-default hover:text-primary transition-colors"
      >
        Enter asset ID manually
      </button>
    </section>
    <!-- Verified Section (Collapsible) -->
    <section class="bg-surface rounded-xl border border-border-default overflow-hidden shadow-sm">
      <div
        class="p-stack flex items-center justify-between cursor-pointer hover:bg-surface-subtle transition-colors border-b border-border-default"
      >
        <div class="flex items-center gap-inline">
          <div class="w-8 h-8 rounded-full bg-metric-sage flex items-center justify-center">
            <span class="material-symbols-outlined text-status-in-stock text-sm"> check </span>
          </div>
          <h2 class="font-h3 text-h3">Verified</h2>
          <span
            class="font-mono-data text-mono-data text-text-secondary bg-surface-subtle px-2 py-1 rounded-full border border-border-default"
          >
            42
          </span>
        </div>
        <span
          class="material-symbols-outlined text-text-secondary transition-transform duration-200"
        >
          expand_more
        </span>
      </div>
      <!-- Collapsed content hidden for brevity in this state, normally JS would toggle a class -->
    </section>
    <!-- Unresolved Section (Collapsible) -->
    <section class="bg-surface rounded-xl border border-border-default overflow-hidden shadow-sm">
      <div
        @click="unresolvedExpanded = !unresolvedExpanded"
        class="p-stack flex items-center justify-between cursor-pointer hover:bg-surface-subtle transition-colors"
      >
        <div class="flex items-center gap-inline">
          <div class="w-8 h-8 rounded-full bg-metric-amber flex items-center justify-center">
            <span class="material-symbols-outlined text-primary-container text-sm"> warning </span>
          </div>
          <h2 class="font-h3 text-h3">Unresolved</h2>
          <span
            class="font-mono-data text-mono-data text-text-secondary bg-surface-subtle px-2 py-1 rounded-full border border-border-default"
          >
            {{ discrepancies.length }}
          </span>
        </div>
        <span
          class="material-symbols-outlined text-text-secondary transition-transform duration-200"
          :class="unresolvedExpanded ? 'rotate-180' : ''"
        >
          expand_more
        </span>
      </div>
      <!-- Expanded Content List -->
      <div v-show="unresolvedExpanded" class="flex flex-col border-t border-border-default">
        <div
          v-for="item in discrepancies"
          :key="item.id"
          class="p-stack border-b border-border-default last:border-b-0 hover:bg-surface-subtle transition-colors"
        >
          <div class="flex justify-between items-start">
            <div>
              <p class="font-mono-data text-mono-data text-text-primary mb-1">
                {{ item.assetTag }}
              </p>
              <p class="font-body text-body text-text-secondary">
                {{ item.issueType }}
              </p>
            </div>
            <span
              class="font-metadata text-metadata px-2 py-1 rounded border"
              :class="
                item.issueType === 'Missing'
                  ? 'bg-error-container text-on-error-container border-red-200'
                  : 'bg-metric-amber text-primary-container border-amber-200'
              "
            >
              {{ item.issueType }}
            </span>
          </div>
          <p class="font-metadata text-metadata text-text-secondary mt-base">
            Last seen: {{ item.lastKnownLocation }}
          </p>
        </div>
      </div>
      <div class="p-stack bg-surface-subtle border-t border-border-default text-center">
        <button
          class="font-h3 text-h3 text-primary hover:text-primary-container transition-colors"
          @click="$router.push({ name: 'AuditDiscrepancyResolution' })"
        >
          View All Unresolved
        </button>
      </div>
    </section>
  </main>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { getDiscrepancies } from '../api/audits';
import type { AuditDiscrepancy } from '../types/models';

const unresolvedExpanded = ref(true);
const discrepancies = ref<AuditDiscrepancy[]>([]);

const fetchDiscrepancies = async () => {
  try {
    discrepancies.value = await getDiscrepancies();
  } catch (err) {
    console.error(err);
  }
};

onMounted(() => {
  fetchDiscrepancies();
});
</script>
