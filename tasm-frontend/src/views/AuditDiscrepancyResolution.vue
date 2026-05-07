<template>
  <main class="space-y-section-gap pb-24">
    <div class="mb-section-gap flex items-end justify-between">
      <div>
        <h2 class="font-h1 text-h1 text-text-primary mb-2">Audit Discrepancy Resolution</h2>
        <p class="font-body text-body text-text-secondary">
          <span v-if="auditId">Session ID: #{{ auditId }} •</span>
          {{ stats.total }} total items scanned
        </p>
      </div>
      <div class="flex gap-3">
        <button
          class="px-4 py-2 bg-surface border border-border-default rounded-lg font-body text-body text-text-primary hover:-translate-y-0.5 transition-transform shadow-sm flex items-center gap-2"
          @click="fetchDiscrepancies"
        >
          <span class="material-symbols-outlined text-[18px]">refresh</span>
          Refresh
        </button>
      </div>
    </div>

    <!-- Stats Cards -->
    <div class="grid grid-cols-2 md:grid-cols-5 gap-inline mb-section-gap">
      <div
        class="bg-surface border border-border-default rounded-xl p-4 shadow-[0_4px_4px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform"
      >
        <div class="font-metadata text-metadata text-text-secondary mb-1">Total Discrepancies</div>
        <div class="font-kpi-number text-kpi-number text-text-primary">{{ stats.total }}</div>
      </div>
      <div
        class="bg-metric-sage border border-border-default rounded-xl p-4 shadow-[0_4px_4px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform"
      >
        <div class="font-metadata text-metadata text-text-secondary mb-1">Resolved</div>
        <div class="font-kpi-number text-kpi-number text-status-in-stock">{{ stats.resolved }}</div>
      </div>
      <div
        class="bg-error-container border border-error-container rounded-xl p-4 shadow-[0_4px_4px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform relative overflow-hidden"
      >
        <div class="absolute left-0 top-0 bottom-0 w-1 bg-status-critical"></div>
        <div class="font-metadata text-metadata text-status-critical mb-1 pl-2">Missing</div>
        <div class="font-kpi-number text-kpi-number text-status-critical pl-2">
          {{ stats.missing }}
        </div>
      </div>
      <div
        class="bg-metric-lavender border border-border-default rounded-xl p-4 shadow-[0_4px_4px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform"
      >
        <div class="font-metadata text-metadata text-text-secondary mb-1">Unregistered</div>
        <div class="font-kpi-number text-kpi-number text-text-primary">
          {{ stats.unregistered }}
        </div>
      </div>
      <div
        class="bg-metric-amber border border-border-default rounded-xl p-4 shadow-[0_4px_4px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform relative overflow-hidden"
      >
        <div class="absolute left-0 top-0 bottom-0 w-1 bg-amber-500"></div>
        <div class="font-metadata text-metadata text-text-secondary mb-1 pl-2">Mismatches</div>
        <div class="font-kpi-number text-kpi-number text-text-primary pl-2">
          {{ stats.mismatch }}
        </div>
      </div>
    </div>

    <!-- Discrepancy Table -->
    <div
      class="bg-surface border border-border-default rounded-xl shadow-[0_4px_4px_rgba(0,0,0,0.02)] overflow-hidden"
    >
      <div
        class="p-4 border-b border-border-default flex justify-between items-center bg-surface-subtle"
      >
        <div class="font-h2 text-h2 text-text-primary">Discrepancy Details</div>
        <div class="flex items-center gap-3">
          <select
            v-model="issueTypeFilter"
            class="bg-surface border border-border-default rounded-md font-body text-body text-text-primary py-1.5 pl-3 pr-8 focus:ring-1 focus:ring-primary text-sm"
          >
            <option value="">All Issue Types</option>
            <option value="Missing">Missing</option>
            <option value="Location Mismatch">Location Mismatch</option>
            <option value="Unregistered">Unregistered</option>
          </select>
          <select
            v-model="statusFilter"
            class="bg-surface border border-border-default rounded-md font-body text-body text-text-primary py-1.5 pl-3 pr-8 focus:ring-1 focus:ring-primary text-sm"
          >
            <option value="">All Statuses</option>
            <option value="Open">Open</option>
            <option value="Resolved">Resolved</option>
            <option value="Dismissed">Dismissed</option>
          </select>
        </div>
      </div>
      <div class="overflow-x-auto p-4">
        <div v-if="loading" class="flex justify-center py-12">
          <span class="material-symbols-outlined animate-spin text-primary text-4xl">sync</span>
        </div>
        <DataTable
          v-else
          :value="filteredDiscrepancies"
          paginator
          :rows="10"
          class="w-full text-left"
          rowHover
        >
          <Column field="assetTag" header="Asset Tag" sortable>
            <template #body="slotProps">
              <span class="font-mono-data text-mono-data text-text-primary">{{
                slotProps.data.assetTag
              }}</span>
            </template>
          </Column>
          <Column field="assetName" header="Asset Name" sortable>
            <template #body="slotProps">
              <span class="text-text-primary">{{ slotProps.data.assetName || '—' }}</span>
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
          <Column field="lastKnownLocation" header="Last Known" sortable>
            <template #body="slotProps">
              <span class="font-body text-body text-text-secondary">
                {{ slotProps.data.lastKnownLocation || '—' }}
              </span>
            </template>
          </Column>
          <Column field="scannedLocation" header="Scanned At" sortable>
            <template #body="slotProps">
              <span class="font-body text-body font-medium text-text-primary">
                {{ slotProps.data.scannedLocation || '—' }}
              </span>
            </template>
          </Column>
          <Column field="status" header="Status" sortable>
            <template #body="slotProps">
              <span
                class="inline-flex items-center px-2 py-0.5 rounded-full text-xs font-bold"
                :class="
                  slotProps.data.status === 'Resolved'
                    ? 'bg-metric-sage text-status-in-stock'
                    : slotProps.data.status === 'Dismissed'
                      ? 'bg-surface-variant text-text-secondary'
                      : 'bg-error-container text-status-critical'
                "
              >
                {{ slotProps.data.status || 'Open' }}
              </span>
            </template>
          </Column>
          <Column header="Actions" alignFrozen="right">
            <template #body="slotProps">
              <div
                v-if="slotProps.data.status === 'Open' || !slotProps.data.status"
                class="flex gap-2"
              >
                <select
                  v-model="selectedActions[slotProps.data.id]"
                  class="text-xs border border-border-default rounded px-2 py-1 bg-surface"
                >
                  <option disabled value="">Action...</option>
                  <option
                    v-for="action in actionsForDisc(slotProps.data)"
                    :key="action.value"
                    :value="action.value"
                  >
                    {{ action.label }}
                  </option>
                </select>
                <button
                  class="px-3 py-1 bg-primary text-on-primary text-xs rounded font-bold hover:opacity-90 disabled:opacity-50"
                  :disabled="
                    !selectedActions[slotProps.data.id] || resolvingId === slotProps.data.id
                  "
                  @click="handleResolve(slotProps.data)"
                >
                  <span
                    v-if="resolvingId === slotProps.data.id"
                    class="material-symbols-outlined animate-spin text-sm"
                    >sync</span
                  >
                  <span v-else>Resolve</span>
                </button>
              </div>
              <div v-else class="text-xs text-text-secondary italic">
                {{ slotProps.data.resolution || slotProps.data.status }}
              </div>
            </template>
          </Column>
        </DataTable>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useRoute } from 'vue-router';
import { getDiscrepancies, resolveDiscrepancy } from '../api/audits';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import type { AuditDiscrepancy } from '../types/models';

const route = useRoute();
const auditId = computed(() => route.query['auditId'] as string | undefined);

const discrepancies = ref<AuditDiscrepancy[]>([]);
const loading = ref(true);
const issueTypeFilter = ref('');
const statusFilter = ref('');
const selectedActions = ref<Record<number, string>>({});
const resolvingId = ref<number | null>(null);

const filteredDiscrepancies = computed(() => {
  return discrepancies.value.filter((d) => {
    if (issueTypeFilter.value && d.issueType !== issueTypeFilter.value) return false;
    if (statusFilter.value) {
      const s = d.status || 'Open';
      if (s !== statusFilter.value) return false;
    }
    return true;
  });
});

const stats = computed(() => ({
  total: discrepancies.value.length,
  resolved: discrepancies.value.filter((d) => d.status === 'Resolved').length,
  missing: discrepancies.value.filter((d) => d.issueType === 'Missing').length,
  unregistered: discrepancies.value.filter((d) => d.issueType === 'Unregistered').length,
  mismatch: discrepancies.value.filter((d) => d.issueType === 'Location Mismatch').length,
}));

function actionsForDisc(disc: AuditDiscrepancy) {
  if (disc.issueType === 'Missing') {
    return [
      { label: 'Confirm Found (update location)', value: 'confirm_location' },
      { label: 'Mark as Lost', value: 'mark_lost' },
      { label: 'Dismiss', value: 'dismiss' },
    ];
  }
  if (disc.issueType === 'Location Mismatch') {
    return [
      { label: `Set Location to "${disc.scannedLocation}"`, value: 'confirm_location' },
      { label: 'Revert to Last Known', value: 'update_location' },
      { label: 'Dismiss', value: 'dismiss' },
    ];
  }
  if (disc.issueType === 'Unregistered') {
    return [
      { label: 'Flag for Registration', value: 'register' },
      { label: 'Dismiss / Not an Asset', value: 'dismiss' },
    ];
  }
  return [{ label: 'Dismiss', value: 'dismiss' }];
}

const handleResolve = async (disc: AuditDiscrepancy) => {
  const action = selectedActions.value[disc.id] as any;
  if (!action) return;
  resolvingId.value = disc.id;
  try {
    const resolved = await resolveDiscrepancy(disc.id, {
      action,
      newLocation: action === 'update_location' ? disc.lastKnownLocation : undefined,
    });
    const idx = discrepancies.value.findIndex((d) => d.id === disc.id);
    if (idx !== -1) discrepancies.value[idx] = resolved;
    delete selectedActions.value[disc.id];
  } catch (err) {
    console.error('Failed to resolve:', err);
    alert('Failed to resolve discrepancy.');
  } finally {
    resolvingId.value = null;
  }
};

const fetchDiscrepancies = async () => {
  loading.value = true;
  try {
    discrepancies.value = await getDiscrepancies({
      auditSessionId: auditId.value,
    });
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
