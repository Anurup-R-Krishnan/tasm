<template>
  <main class="space-y-section-gap pb-24">
    <!-- Page Header & Actions -->
    <div class="flex justify-between items-end mb-section-gap">
      <div>
        <h1 class="font-h1 text-h1 text-text-primary mb-1">Audit Sessions</h1>
        <p class="font-metadata text-metadata text-text-secondary">
          Manage and track physical asset verifications across facilities.
        </p>
      </div>
      <button
        class="bg-primary hover:bg-primary-hover text-on-primary rounded font-body text-body py-2 px-5 transition-colors shadow-sm flex items-center gap-2"
        @click="handleStartAudit"
      >
        <span class="material-symbols-outlined" style="font-size: 18px"> qr_code_scanner </span>
        Start New Audit Session
      </button>
    </div>
    <!-- Section 1: Active/Past Sessions Table -->
    <section
      class="bg-surface border border-border-default rounded-xl shadow-[0_4px_4px_rgba(0,0,0,0.02)] mb-section-gap overflow-hidden"
    >
      <div
        class="px-card-padding py-4 border-b border-border-default bg-surface-subtle flex justify-between items-center"
      >
        <h2 class="font-h2 text-h2 text-text-primary">Recent Audit Sessions</h2>
        <div class="flex items-center gap-2">
          <select
            v-model="statusFilter"
            class="bg-surface border border-border-default rounded text-metadata font-metadata text-text-secondary py-1 px-2 focus:ring-1 focus:ring-primary outline-none cursor-pointer"
          >
            <option value="all">All Sessions</option>
            <option value="In Progress">In Progress</option>
            <option value="Completed">Completed</option>
          </select>
          <button class="text-text-secondary hover:text-primary transition-colors">
            <span class="material-symbols-outlined"> filter_list </span>
          </button>
        </div>
      </div>
      <div class="overflow-x-auto border-t border-border-default">
        <DataTable
          :value="filteredAudits"
          :loading="loadingAudits"
          paginator
          :rows="5"
          tableStyle="min-width: 50rem"
          class="w-full"
        >
          <Column field="sessionId" header="Session ID" sortable>
            <template #body="slotProps">
              <span
                class="font-mono-data text-mono-data text-text-primary border-l-[3px]"
                :class="
                  slotProps.data.status === 'In Progress'
                    ? 'border-primary pl-2'
                    : 'border-transparent pl-2'
                "
              >
                {{ slotProps.data.sessionId }}
              </span>
            </template>
          </Column>
          <Column field="location" header="Scope" sortable>
            <template #body="slotProps">
              <span class="text-text-primary">{{ slotProps.data.location }}</span>
            </template>
          </Column>
          <Column field="createdAt" header="Date" sortable>
            <template #body="slotProps">
              <span class="text-text-secondary">{{
                new Date(slotProps.data.createdAt).toLocaleDateString()
              }}</span>
            </template>
          </Column>
          <Column field="verifiedPercent" header="Verified %" sortable>
            <template #body="slotProps">
              <div class="flex items-center gap-2">
                <span
                  class="font-medium"
                  :class="slotProps.data.verifiedPercent === 100 ? 'text-status-in-stock' : ''"
                >
                  {{ slotProps.data.verifiedPercent }}%
                </span>
                <div class="w-16 h-1.5 bg-surface-variant rounded-full overflow-hidden">
                  <div
                    class="h-full rounded-full"
                    :class="
                      slotProps.data.verifiedPercent === 100 ? 'bg-status-in-stock' : 'bg-primary'
                    "
                    :style="'width: ' + slotProps.data.verifiedPercent + '%'"
                  ></div>
                </div>
              </div>
            </template>
          </Column>
          <Column field="missingCount" header="Missing" sortable>
            <template #body="slotProps">
              <span
                :class="
                  slotProps.data.missingCount > 0
                    ? 'text-status-critical font-medium'
                    : 'text-text-secondary'
                "
              >
                {{ slotProps.data.missingCount }}
              </span>
            </template>
          </Column>
          <Column field="unregisteredCount" header="Unregistered" sortable>
            <template #body="slotProps">
              <span
                :class="
                  slotProps.data.unregisteredCount > 0
                    ? 'text-primary-container'
                    : 'text-text-secondary'
                "
              >
                {{ slotProps.data.unregisteredCount }}
              </span>
            </template>
          </Column>
          <Column field="status" header="Status" sortable>
            <template #body="slotProps">
              <span
                v-if="slotProps.data.status === 'In Progress'"
                class="inline-flex items-center px-2 py-1 rounded bg-primary-container/20 text-primary font-metadata text-metadata font-medium"
              >
                In Progress
              </span>
              <span
                v-else
                class="inline-flex items-center px-2 py-1 rounded bg-metric-sage text-status-in-stock font-metadata text-metadata font-medium"
              >
                Completed
              </span>
            </template>
          </Column>
          <Column header="Action" alignFrozen="right">
            <template #body>
              <button
                class="text-primary hover:text-primary-hover font-medium transition-colors"
                @click="router.push('/audit-cleanup')"
              >
                View Report
              </button>
            </template>
          </Column>
        </DataTable>
      </div>
      <div class="px-card-padding py-3 border-t border-border-default bg-surface flex justify-end">
        <button
          class="text-text-secondary hover:text-text-primary font-metadata text-metadata transition-colors flex items-center gap-1"
          @click="statusFilter = 'all'"
        >
          View All Sessions
          <span class="material-symbols-outlined" style="font-size: 16px"> arrow_forward </span>
        </button>
      </div>
    </section>
    <!-- Section 2: Discrepancy Resolution Queue (Bento/Card layout) -->
    <h2 class="font-h2 text-h2 text-text-primary mb-4">Discrepancy Resolution Queue</h2>
    <div class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-6">
      <div
        v-for="discrepancy in discrepancies"
        :key="discrepancy.id"
        class="bg-surface border border-border-default rounded-xl p-card-padding shadow-[0_4px_4px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform duration-200 flex flex-col h-full"
      >
        <div class="flex justify-between items-start mb-4">
          <div class="flex items-center gap-2">
            <span
              class="material-symbols-outlined p-1.5 rounded-lg"
              :class="
                discrepancy.type === 'Missing'
                  ? 'text-status-critical bg-error-container'
                  : discrepancy.type === 'Location Mismatch'
                    ? 'text-primary-container bg-accent-subtle'
                    : 'text-tertiary bg-tertiary-fixed'
              "
            >
              {{ discrepancy.icon }}
            </span>
            <span
              class="font-metadata text-metadata font-medium uppercase tracking-wider"
              :class="
                discrepancy.type === 'Missing'
                  ? 'text-status-critical'
                  : discrepancy.type === 'Location Mismatch'
                    ? 'text-primary-container'
                    : 'text-tertiary'
              "
            >
              {{ discrepancy.type }}
            </span>
          </div>
          <span
            class="font-metadata text-metadata font-bold px-2 py-1 rounded"
            :class="
              discrepancy.type === 'Missing'
                ? 'text-status-critical bg-error-container/20 border border-error-container/30'
                : discrepancy.type === 'Location Mismatch'
                  ? 'text-primary-container bg-accent-subtle'
                  : 'text-tertiary bg-tertiary-fixed'
            "
          >
            {{ discrepancy.duration }}
          </span>
        </div>
        <h3 class="font-h3 text-h3 text-text-primary mb-1">{{ discrepancy.assetName }}</h3>
        <p class="font-mono-data text-mono-data text-text-secondary mb-4">
          {{ discrepancy.assetTag }}
        </p>
        <div class="bg-surface-subtle p-3 rounded border border-border-default mb-4 flex-1">
          <p class="font-metadata text-metadata text-text-secondary mb-1">Flagged In:</p>
          <p class="font-body text-body text-text-primary">{{ discrepancy.flaggedIn }}</p>
          <p v-if="discrepancy.notes" class="font-metadata text-metadata text-text-secondary mt-2">
            {{ discrepancy.notes }}
          </p>
        </div>
        <div class="mt-auto pt-4 border-t border-border-default flex justify-end gap-2">
          <select
            v-model="discrepancy.selectedAction"
            class="form-select w-full bg-surface border border-[#D6D3CE] rounded text-body font-body text-text-primary py-1.5 px-3 focus:ring-primary focus:border-primary"
          >
            <option disabled value="">Select Action...</option>
            <option v-for="action in discrepancy.actions" :key="action.value" :value="action.value">
              {{ action.label }}
            </option>
          </select>
          <button
            class="bg-surface border border-[#D6D3CE] hover:bg-surface-subtle text-text-primary rounded p-1.5 transition-colors flex-shrink-0"
            @click="handleResolveDiscrepancy(discrepancy.id)"
          >
            <span class="material-symbols-outlined"> check </span>
          </button>
        </div>
      </div>
      <div
        v-if="discrepancies.length === 0"
        class="col-span-full py-12 text-center text-text-secondary italic"
      >
        All discrepancies resolved.
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useRouter } from 'vue-router';
import { getAudits } from '../api/audits';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import type { AuditSession as AuditSessionModel } from '../types/models';

const router = useRouter();

interface AuditSessionRow extends AuditSessionModel {
  sessionId: string;
  location: string;
  verifiedPercent: number;
  missingCount: number;
  unregisteredCount: number;
}

const audits = ref<AuditSessionRow[]>([]);
const loadingAudits = ref(true);
const statusFilter = ref('all');

const discrepancies = ref([
  {
    id: 1,
    type: 'Missing',
    icon: 'warning',
    duration: '4 Days',
    assetName: 'Dell Latitude 5420',
    assetTag: 'AST-LPT-8892',
    flaggedIn: 'AUD-2023-11-A (Gayathri - Floor 2)',
    selectedAction: '',
    actions: [
      { label: 'Confirm Location', value: 'confirm' },
      { label: 'Mark as Lost', value: 'lost' },
      { label: 'Reassign Asset', value: 'reassign' },
    ],
  },
  {
    id: 2,
    type: 'Location Mismatch',
    icon: 'location_off',
    duration: '2 Days',
    assetName: 'Herman Miller Aeron',
    assetTag: 'AST-FUR-1104',
    flaggedIn: 'AUD-2023-11-A (Gayathri - Floor 2)',
    notes: 'Expected: Nila - Floor 4',
    selectedAction: '',
    actions: [
      { label: 'Update Location to Gayathri', value: 'confirm' },
      { label: 'Mark for Return to Nila', value: 'lost' },
    ],
  },
  {
    id: 3,
    type: 'Unregistered',
    icon: 'help_center',
    duration: '1 Day',
    assetName: 'Unknown Cisco Switch',
    assetTag: 'No Asset Tag',
    flaggedIn: 'AUD-2023-11-A (Gayathri - Floor 2)',
    notes: 'Found in network closet G2-B.',
    selectedAction: '',
    actions: [
      { label: 'Register New Asset', value: 'register' },
      { label: 'Ignore/Not Asset', value: 'ignore' },
    ],
  },
]);

const filteredAudits = computed(() => {
  if (statusFilter.value === 'all') return audits.value;
  return audits.value.filter((audit) => audit.status === statusFilter.value);
});

const handleStartAudit = () => {
  // Navigate to mobile scan mode as a demonstration of starting a field audit
  router.push('/audit-scan');
};

const handleResolveDiscrepancy = (id: number) => {
  const discrepancy = discrepancies.value.find((d) => d.id === id);
  if (!discrepancy?.selectedAction) {
    alert('Please select an action first.');
    return;
  }

  if (confirm(`Resolve ${discrepancy.assetTag} with action: ${discrepancy.selectedAction}?`)) {
    discrepancies.value = discrepancies.value.filter((d) => d.id !== id);
  }
};

const fetchAudits = async () => {
  try {
    const data = await getAudits();
    audits.value = data as unknown as AuditSessionRow[];
  } catch (error) {
    console.error('Failed to fetch audits:', error);
  } finally {
    loadingAudits.value = false;
  }
};

onMounted(() => {
  fetchAudits();
});
</script>
