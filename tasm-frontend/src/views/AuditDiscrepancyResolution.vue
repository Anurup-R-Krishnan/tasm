<template>
  <main class="flex-1 p-page-margin max-w-[1400px] w-full mx-auto pb-32">
    <div class="mb-section-gap flex items-end justify-between">
      <div>
        <h2 class="font-h1 text-h1 text-text-primary mb-2">Audit Discrepancy Resolution</h2>
        <p class="font-body text-body text-text-secondary">
          Session ID: #AUD-2023-Q4-082 • Building C Wing
        </p>
      </div>
      <div class="flex gap-3">
        <button
          class="px-4 py-2 bg-surface border border-border-default rounded-lg font-body text-body text-text-primary hover:-translate-y-0.5 transition-transform shadow-sm flex items-center gap-2"
        >
          <span class="material-symbols-outlined text-[18px]" data-icon="download"> download </span>
          Export Report
        </button>
      </div>
    </div>
    <div class="grid grid-cols-5 gap-inline mb-section-gap">
      <div
        class="bg-surface border border-border-default rounded-xl p-4 shadow-[0_4px_4px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform cursor-pointer"
      >
        <div class="font-metadata text-metadata text-text-secondary mb-1">Total Scope</div>
        <div class="font-kpi-number text-kpi-number text-text-primary">1,452</div>
      </div>
      <div
        class="bg-metric-sage border border-border-default rounded-xl p-4 shadow-[0_4px_4px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform cursor-pointer"
      >
        <div class="font-metadata text-metadata text-text-secondary mb-1">Verified</div>
        <div class="font-kpi-number text-kpi-number text-text-primary">1,390</div>
      </div>
      <div
        class="bg-error-container border border-error-container rounded-xl p-4 shadow-[0_4px_4px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform cursor-pointer relative overflow-hidden"
      >
        <div class="absolute left-0 top-0 bottom-0 w-1 bg-status-critical"></div>
        <div class="font-metadata text-metadata text-status-critical mb-1 pl-2">Missing</div>
        <div class="font-kpi-number text-kpi-number text-status-critical pl-2">18</div>
      </div>
      <div
        class="bg-metric-lavender border border-border-default rounded-xl p-4 shadow-[0_4px_4px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform cursor-pointer"
      >
        <div class="font-metadata text-metadata text-text-secondary mb-1">Unregistered</div>
        <div class="font-kpi-number text-kpi-number text-text-primary">12</div>
      </div>
      <div
        class="bg-metric-amber border border-border-default rounded-xl p-4 shadow-[0_4px_4px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform cursor-pointer relative overflow-hidden"
      >
        <div class="absolute left-0 top-0 bottom-0 w-1 bg-amber-500"></div>
        <div class="font-metadata text-metadata text-text-secondary mb-1 pl-2">Mismatches</div>
        <div class="font-kpi-number text-kpi-number text-text-primary pl-2">32</div>
      </div>
    </div>
    <div
      class="bg-surface border border-border-default rounded-xl shadow-[0_4px_4px_rgba(0,0,0,0.02)] overflow-hidden"
    >
      <div
        class="p-4 border-b border-border-default flex justify-between items-center bg-surface-subtle"
      >
        <div class="font-h2 text-h2 text-text-primary">Discrepancy Details</div>
        <div class="flex items-center gap-3">
          <select
            class="bg-surface border border-border-default rounded-md font-body text-body text-text-primary py-1.5 pl-3 pr-8 focus:ring-1 focus:ring-amber-500 focus:border-amber-500 text-sm"
          >
            <option>All Issue Types</option>
            <option>Missing</option>
            <option>Location Mismatch</option>
            <option>Unregistered</option>
          </select>
        </div>
      </div>
      <div class="overflow-x-auto p-4">
        <DataTable
          :value="discrepancies"
          :loading="loading"
          paginator
          :rows="10"
          class="w-full text-left"
          rowHover
        >
          <Column selectionMode="multiple" headerStyle="width: 3rem"></Column>
          <Column field="assetTag" header="Asset Tag" sortable>
            <template #body="slotProps">
              <span class="font-mono-data text-mono-data text-text-primary">{{
                slotProps.data.assetTag
              }}</span>
            </template>
          </Column>
          <Column field="issueType" header="Issue Type" sortable>
            <template #body="slotProps">
              <span
                class="inline-flex items-center px-2 py-0.5 rounded-full font-metadata text-metadata border"
                :class="
                  slotProps.data.issueType === 'Location Mismatch'
                    ? 'bg-metric-amber text-on-primary-container border-amber-200'
                    : slotProps.data.issueType === 'Missing'
                      ? 'bg-error-container text-on-error-container border-red-200'
                      : 'bg-metric-lavender text-tertiary border-purple-200'
                "
              >
                {{ slotProps.data.issueType }}
              </span>
            </template>
          </Column>
          <Column field="lastKnownLocation" header="Last Known Location" sortable>
            <template #body="slotProps">
              <span
                class="font-body text-body"
                :class="
                  slotProps.data.lastKnownLocation === 'None'
                    ? 'text-text-secondary italic'
                    : 'text-text-secondary'
                "
              >
                {{ slotProps.data.lastKnownLocation }}
              </span>
            </template>
          </Column>
          <Column field="scannedLocation" header="Scanned Location" sortable>
            <template #body="slotProps">
              <span
                class="font-body text-body font-medium"
                :class="
                  slotProps.data.scannedLocation === 'Not Scanned'
                    ? 'text-text-secondary italic'
                    : 'text-text-primary'
                "
              >
                {{ slotProps.data.scannedLocation }}
              </span>
            </template>
          </Column>
          <Column field="recommendedAction" header="Recommended Action" sortable>
            <template #body="slotProps">
              <span
                class="font-body text-body"
                :class="
                  slotProps.data.issueType === 'Missing' &&
                  slotProps.data.recommendedAction.includes('High Priority')
                    ? 'text-status-critical font-medium'
                    : 'text-text-secondary'
                "
              >
                {{ slotProps.data.recommendedAction }}
              </span>
            </template>
          </Column>
          <Column header="Actions" alignFrozen="right">
            <template #body="slotProps">
              <div class="space-x-2 text-right">
                <button
                  v-if="slotProps.data.issueType === 'Unregistered'"
                  class="px-3 py-1.5 bg-[#1C1917] text-white rounded font-metadata text-metadata hover:opacity-90 transition-opacity"
                >
                  Register
                </button>
                <button
                  v-else-if="slotProps.data.issueType === 'Missing'"
                  class="px-3 py-1.5 bg-surface border border-border-default rounded font-metadata text-metadata text-text-primary hover:bg-stone-50 transition-colors"
                >
                  Mark Lost
                </button>
                <button
                  v-else
                  class="px-3 py-1.5 bg-surface border border-border-default rounded font-metadata text-metadata text-text-primary hover:bg-stone-50 transition-colors"
                >
                  Confirm New
                </button>
              </div>
            </template>
          </Column>
        </DataTable>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';

const discrepancies = ref<any[]>([]);
const loading = ref(true);

const fetchDiscrepancies = async () => {
  try {
    const res = await fetch('http://localhost:8080/api/discrepancies');
    if (res.ok) {
      discrepancies.value = await res.json();
    }
  } catch (error) {
    console.error('Failed to fetch discrepancies:', error);
  } finally {
    loading.value = false;
  }
};

onMounted(() => {
  fetchDiscrepancies();
});
</script>
