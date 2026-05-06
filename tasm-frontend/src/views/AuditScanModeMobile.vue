<template>
  <main class="space-y-section-gap pb-24 px-page-margin">
    <!-- Scanner View -->
    <section
      v-if="isScanning"
      class="fixed inset-0 z-50 bg-black flex flex-col items-center justify-center"
    >
      <div class="absolute top-8 left-8 right-8 flex justify-between items-center text-white z-10">
        <h2 class="font-h2 text-h2">Scan Asset QR</h2>
        <button
          @click="isScanning = false"
          class="w-10 h-10 rounded-full bg-white/10 flex items-center justify-center"
        >
          <span class="material-symbols-outlined">close</span>
        </button>
      </div>

      <div class="w-full max-w-sm px-6">
        <QRScanner @result="onScanResult" @error="onScanError" />
        <p class="text-white/60 text-center mt-6 font-body text-sm px-10">
          Position the asset QR code within the frame to scan automatically.
        </p>
      </div>
    </section>

    <!-- Hero: Massive Scan Button -->
    <section class="flex flex-col items-center justify-center py-section-gap">
      <button
        @click="isScanning = true"
        class="w-full max-w-sm h-[80px] bg-primary text-on-primary rounded-2xl shadow-xl flex items-center justify-center gap-inline active:scale-95 transition-all"
      >
        <span
          class="material-symbols-outlined text-[32px]"
          style="font-variation-settings: 'FILL' 1"
        >
          qr_code_scanner
        </span>
        <span class="font-h2 text-h2 uppercase tracking-widest"> Start Scanning </span>
      </button>
      <button
        @click="handleManualEntry"
        class="mt-stack font-metadata text-metadata text-text-secondary uppercase tracking-widest hover:text-primary transition-colors flex items-center gap-1"
      >
        <span class="material-symbols-outlined text-sm">keyboard</span>
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
import { useRouter } from 'vue-router';
import { getDiscrepancies } from '../api/audits';
import type { AuditDiscrepancy } from '../types/models';
import QRScanner from '../components/QRScanner.vue';

const router = useRouter();
const unresolvedExpanded = ref(true);
const discrepancies = ref<AuditDiscrepancy[]>([]);
const isScanning = ref(false);

const onScanResult = (decodedText: string) => {
  isScanning.value = false;
  // Assuming the QR code contains the asset ID or tag
  // Redirect to asset detail or check-out flow
  router.push({ name: 'AssetDetail', params: { id: decodedText } });
};

const onScanError = (error: string) => {
  console.error('Scanner error:', error);
};

const handleManualEntry = () => {
  router.push('/inventory');
};

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
