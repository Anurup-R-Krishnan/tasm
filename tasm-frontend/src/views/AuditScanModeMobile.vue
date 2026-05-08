<template>
  <main class="space-y-section-gap pb-24 px-page-margin">
    <!-- Scanner View (fullscreen overlay) -->
    <section
      v-if="isScanning"
      class="fixed inset-0 z-50 bg-black flex flex-col items-center justify-center"
    >
      <div class="absolute top-8 left-8 right-8 flex justify-between items-center text-white z-10">
        <div>
          <h2 class="font-h2 text-h2">Scan Asset QR</h2>
          <p class="text-white/60 text-sm mt-1" v-if="currentAudit">
            Session: {{ currentAudit.title }}
          </p>
        </div>
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

      <!-- Manual Entry -->
      <div class="mt-6 w-full max-w-sm px-6">
        <div class="flex gap-2">
          <input
            v-model="manualTagId"
            type="text"
            placeholder="Or type asset tag ID..."
            class="flex-1 bg-white/10 border border-white/20 rounded-xl px-4 py-2.5 text-white placeholder-white/40 focus:outline-none focus:border-white/50"
            @keyup.enter="submitManualTag"
          />
          <button
            @click="submitManualTag"
            class="bg-primary text-on-primary rounded-xl px-4 py-2.5 font-medium"
          >
            Go
          </button>
        </div>
      </div>
    </section>

    <!-- Scan Result Toast -->
    <Transition name="slide-up">
      <div
        v-if="lastScanResult"
        class="fixed bottom-6 left-1/2 -translate-x-1/2 z-40 w-full max-w-sm mx-auto px-4"
      >
        <div
          class="rounded-2xl p-4 shadow-2xl border flex items-start gap-3"
          :class="
            lastScanResult.result === 'found' && lastScanResult.locationMatch
              ? 'bg-metric-sage border-green-200 text-status-in-stock'
              : lastScanResult.result === 'found' && !lastScanResult.locationMatch
                ? 'bg-metric-amber border-amber-200 text-amber-800'
                : 'bg-error-container border-red-200 text-status-critical'
          "
        >
          <span class="material-symbols-outlined text-2xl">
            {{
              lastScanResult.result === 'found' && lastScanResult.locationMatch
                ? 'check_circle'
                : lastScanResult.result === 'found' && !lastScanResult.locationMatch
                  ? 'location_off'
                  : 'report'
            }}
          </span>
          <div class="flex-1">
            <p class="font-bold text-sm">
              {{
                lastScanResult.result === 'found' && lastScanResult.locationMatch
                  ? '✓ Asset Verified'
                  : lastScanResult.result === 'found' && !lastScanResult.locationMatch
                    ? '⚠ Location Mismatch'
                    : '✗ Unregistered Asset'
              }}
            </p>
            <p class="text-xs mt-0.5 opacity-80">
              {{
                lastScanResult.result === 'found'
                  ? lastScanResult.asset?.name + ' (' + lastScanResult.asset?.tagId + ')'
                  : 'Tag not found in system — logged as discrepancy'
              }}
            </p>
          </div>
          <button @click="lastScanResult = null" class="opacity-60 hover:opacity-100">
            <span class="material-symbols-outlined text-lg">close</span>
          </button>
        </div>
      </div>
    </Transition>

    <!-- Audit Session Info -->
    <section
      v-if="currentAudit"
      class="bg-surface border border-border-default rounded-xl p-4 shadow-sm"
    >
      <div class="flex items-center justify-between">
        <div>
          <p
            class="font-metadata text-metadata text-text-secondary uppercase tracking-wider mb-0.5"
          >
            Active Session
          </p>
          <h2 class="font-h3 text-h3 text-text-primary">{{ currentAudit.title }}</h2>
          <p class="text-xs text-text-secondary mt-0.5">Auditor: {{ currentAudit.auditorName }}</p>
        </div>
        <div class="text-right">
          <p class="font-kpi-number text-2xl text-primary">{{ currentAudit.progress }}%</p>
          <p class="text-xs text-text-secondary">
            {{ currentAudit.scannedAssets }}/{{ currentAudit.totalAssets }} scanned
          </p>
        </div>
      </div>
      <!-- Progress Bar -->
      <div class="mt-3 h-2 bg-surface-subtle rounded-full overflow-hidden">
        <div
          class="h-full bg-primary rounded-full transition-all duration-500"
          :style="`width: ${currentAudit.progress}%`"
        ></div>
      </div>
    </section>

    <!-- Hero: Massive Scan Button -->
    <section class="flex flex-col items-center justify-center py-section-gap">
      <button
        @click="isScanning = true"
        class="w-full max-w-sm h-[80px] bg-primary text-on-primary rounded-2xl shadow-xl flex items-center justify-center gap-inline active:scale-95 transition-all"
        :disabled="!auditId"
      >
        <span
          class="material-symbols-outlined text-[32px]"
          style="font-variation-settings: 'FILL' 1"
        >
          qr_code_scanner
        </span>
        <span class="font-h2 text-h2 uppercase tracking-widest">Start Scanning</span>
      </button>
      <p v-if="!auditId" class="mt-3 text-text-secondary text-sm text-center">
        No active audit session.
        <button class="text-primary underline" @click="$router.push('/audits')">Start one</button>.
      </p>
      <p
        v-else
        class="mt-stack font-metadata text-metadata text-text-secondary uppercase tracking-widest"
      >
        Session: {{ currentAudit?.title }}
      </p>
    </section>

    <!-- Scan Stats -->
    <div class="grid grid-cols-3 gap-3">
      <div class="bg-metric-sage border border-border-default rounded-xl p-3 text-center shadow-sm">
        <p class="font-kpi-number text-kpi-number text-status-in-stock">{{ verifiedCount }}</p>
        <p class="font-metadata text-metadata text-text-secondary">Verified</p>
      </div>
      <div
        class="bg-metric-amber border border-border-default rounded-xl p-3 text-center shadow-sm"
      >
        <p class="font-kpi-number text-kpi-number text-amber-700">{{ mismatchCount }}</p>
        <p class="font-metadata text-metadata text-text-secondary">Mismatch</p>
      </div>
      <div
        class="bg-error-container border border-border-default rounded-xl p-3 text-center shadow-sm"
      >
        <p class="font-kpi-number text-kpi-number text-status-critical">{{ unresolvedCount }}</p>
        <p class="font-metadata text-metadata text-text-secondary">Unresolved</p>
      </div>
    </div>

    <!-- Unresolved Discrepancies -->
    <section class="bg-surface rounded-xl border border-border-default overflow-hidden shadow-sm">
      <div
        @click="unresolvedExpanded = !unresolvedExpanded"
        class="p-stack flex items-center justify-between cursor-pointer hover:bg-surface-subtle transition-colors"
      >
        <div class="flex items-center gap-inline">
          <div class="w-8 h-8 rounded-full bg-metric-amber flex items-center justify-center">
            <span class="material-symbols-outlined text-primary-container text-sm">warning</span>
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
              <p class="font-body text-text-secondary">
                {{ item.assetName || 'Unknown' }}
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
            Last seen: {{ item.lastKnownLocation || '—' }}
          </p>
        </div>
        <div
          v-if="discrepancies.length === 0"
          class="p-stack text-center text-text-secondary italic"
        >
          No discrepancies yet.
        </div>
      </div>
      <div class="p-stack bg-surface-subtle border-t border-border-default text-center">
        <button
          class="font-h3 text-h3 text-primary hover:text-primary-container transition-colors"
          @click="$router.push({ name: 'AuditDiscrepancyResolution', query: { auditId } })"
        >
          View All Unresolved
        </button>
      </div>
    </section>
  </main>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useRoute } from 'vue-router';
import { getDiscrepancies, scanAssetInAudit, getAuditById } from '../api/audits';
import type { AuditDiscrepancy, AuditSession, ScanResult } from '../types/models';
import QRScanner from '../components/QRScanner.vue';

const route = useRoute();
const unresolvedExpanded = ref(true);
const discrepancies = ref<AuditDiscrepancy[]>([]);
const isScanning = ref(false);
const lastScanResult = ref<ScanResult | null>(null);
const manualTagId = ref('');
const currentAudit = ref<AuditSession | null>(null);
const scanLog = ref<ScanResult[]>([]);

const auditId = computed(() => route.query['auditId'] as string | undefined);

const verifiedCount = computed(
  () => scanLog.value.filter((s: ScanResult) => s.result === 'found' && s.locationMatch).length,
);
const mismatchCount = computed(
  () => scanLog.value.filter((s: ScanResult) => s.result === 'found' && !s.locationMatch).length,
);
const unresolvedCount = computed(() => discrepancies.value.length);

const processTagId = async (tagId: string) => {
  if (!auditId.value) {
    alert('No audit session selected. Please start or select an audit session first.');
    return;
  }

  try {
    const result = await scanAssetInAudit(auditId.value, { tagId: tagId.trim() });
    lastScanResult.value = result;
    scanLog.value.push(result);

    // Update progress on current audit display
    if (result.progress !== undefined && currentAudit.value) {
      currentAudit.value.progress = result.progress;
    }

    // Refresh discrepancies if a new one was created
    if (result.result === 'unregistered' || !result.locationMatch) {
      await fetchDiscrepancies();
    }

    // Auto-dismiss result after 4 seconds
    setTimeout(() => {
      if (lastScanResult.value === result) {
        lastScanResult.value = null;
      }
    }, 4000);
  } catch (err) {
    console.error('Scan error:', err);
    lastScanResult.value = null;
  }
};

const onScanResult = async (decodedText: string) => {
  isScanning.value = false;
  await processTagId(decodedText);
};

const onScanError = (error: string) => {
  console.error('Scanner error:', error);
};

const submitManualTag = async () => {
  if (!manualTagId.value.trim()) return;
  isScanning.value = false;
  await processTagId(manualTagId.value);
  manualTagId.value = '';
};

const fetchDiscrepancies = async () => {
  try {
    discrepancies.value = await getDiscrepancies({
      auditSessionId: auditId.value || undefined,
      status: 'Open',
    });
  } catch (err) {
    console.error(err);
  }
};

const fetchAudit = async () => {
  if (!auditId.value) return;
  try {
    currentAudit.value = await getAuditById(auditId.value);
  } catch (err) {
    console.error('Failed to load audit session:', err);
  }
};

onMounted(() => {
  fetchAudit();
  fetchDiscrepancies();
});
</script>

<style scoped>
.slide-up-enter-active,
.slide-up-leave-active {
  transition: all 0.3s ease;
}
.slide-up-enter-from,
.slide-up-leave-to {
  opacity: 0;
  transform: translate(-50%, 1rem);
}
</style>
