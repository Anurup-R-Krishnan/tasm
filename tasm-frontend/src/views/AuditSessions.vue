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
      <div class="flex gap-3">
        <button
          class="bg-surface border border-border-default text-text-primary rounded font-body py-2 px-4 transition-colors shadow-sm flex items-center gap-2 hover:bg-surface-subtle"
          @click="showCreateModal = true"
        >
          <span class="material-symbols-outlined" style="font-size: 18px">add</span>
          New Session
        </button>
        <button
          class="bg-primary hover:bg-primary-hover text-on-primary rounded font-body py-2 px-5 transition-colors shadow-sm flex items-center gap-2"
          @click="handleStartAudit"
        >
          <span class="material-symbols-outlined" style="font-size: 18px">qr_code_scanner</span>
          Start Scanning
        </button>
      </div>
    </div>

    <!-- Summary Stats -->
    <div class="grid grid-cols-2 md:grid-cols-4 gap-4 mb-section-gap">
      <div class="bg-surface border border-border-default rounded-xl p-4 shadow-sm">
        <p class="font-metadata text-metadata text-text-secondary mb-1">Total Sessions</p>
        <p class="font-kpi-number text-kpi-number text-text-primary">{{ audits.length }}</p>
      </div>
      <div class="bg-metric-sage border border-border-default rounded-xl p-4 shadow-sm">
        <p class="font-metadata text-metadata text-text-secondary mb-1">Active</p>
        <p class="font-kpi-number text-kpi-number text-status-in-stock">{{ activeAuditCount }}</p>
      </div>
      <div class="bg-error-container border border-border-default rounded-xl p-4 shadow-sm">
        <p class="font-metadata text-metadata text-status-critical mb-1">Open Discrepancies</p>
        <p class="font-kpi-number text-kpi-number text-status-critical">
          {{ openDiscrepancyCount }}
        </p>
      </div>
      <div class="bg-metric-lavender border border-border-default rounded-xl p-4 shadow-sm">
        <p class="font-metadata text-metadata text-text-secondary mb-1">Resolved</p>
        <p class="font-kpi-number text-kpi-number text-text-primary">
          {{ resolvedDiscrepancyCount }}
        </p>
      </div>
    </div>

    <!-- Session Table -->
    <section
      class="bg-surface border border-border-default rounded-xl shadow-[0_4px_4px_rgba(0,0,0,0.02)] mb-section-gap overflow-hidden"
    >
      <div
        class="px-card-padding py-4 border-b border-border-default bg-surface-subtle flex justify-between items-center"
      >
        <h2 class="font-h2 text-h2 text-text-primary">Audit Sessions</h2>
        <div class="flex items-center gap-2">
          <select
            v-model="statusFilter"
            class="bg-surface border border-border-default rounded text-metadata font-metadata text-text-secondary py-1 px-2 focus:ring-1 focus:ring-primary outline-none cursor-pointer"
          >
            <option value="all">All Sessions</option>
            <option value="Active">Active</option>
            <option value="Completed">Completed</option>
          </select>
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
          <Column field="title" header="Session Title" sortable>
            <template #body="slotProps">
              <span
                class="font-mono-data text-mono-data text-text-primary border-l-[3px] pl-2"
                :class="
                  slotProps.data.status === 'Active' ? 'border-primary' : 'border-transparent'
                "
              >
                {{ slotProps.data.title }}
              </span>
            </template>
          </Column>
          <Column field="auditorName" header="Auditor" sortable>
            <template #body="slotProps">
              <span class="text-text-primary">{{ slotProps.data.auditorName }}</span>
            </template>
          </Column>
          <Column field="startDate" header="Date" sortable>
            <template #body="slotProps">
              <span class="text-text-secondary">{{
                new Date(slotProps.data.startDate).toLocaleDateString()
              }}</span>
            </template>
          </Column>
          <Column field="progress" header="Progress" sortable>
            <template #body="slotProps">
              <div class="flex items-center gap-2">
                <span
                  class="font-medium"
                  :class="slotProps.data.progress === 100 ? 'text-status-in-stock' : ''"
                >
                  {{ slotProps.data.scannedAssets }}/{{ slotProps.data.totalAssets }}
                </span>
                <div class="w-16 h-1.5 bg-surface-variant rounded-full overflow-hidden">
                  <div
                    class="h-full rounded-full"
                    :class="slotProps.data.progress === 100 ? 'bg-status-in-stock' : 'bg-primary'"
                    :style="'width: ' + slotProps.data.progress + '%'"
                  ></div>
                </div>
                <span class="text-xs text-text-secondary">{{ slotProps.data.progress }}%</span>
              </div>
            </template>
          </Column>
          <Column field="discrepancyCount" header="Discrepancies" sortable>
            <template #body="slotProps">
              <span
                :class="
                  slotProps.data.discrepancyCount > 0
                    ? 'text-status-critical font-medium'
                    : 'text-text-secondary'
                "
              >
                {{ slotProps.data.discrepancyCount }}
              </span>
            </template>
          </Column>
          <Column field="status" header="Status" sortable>
            <template #body="slotProps">
              <span
                v-if="slotProps.data.status === 'Active'"
                class="inline-flex items-center px-2 py-1 rounded bg-primary-container/20 text-primary font-metadata text-metadata font-medium"
              >
                Active
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
            <template #body="slotProps">
              <div class="flex gap-2">
                <button
                  v-if="slotProps.data.status === 'Active'"
                  class="text-primary hover:text-primary-hover font-medium transition-colors text-sm"
                  @click="handleScanAudit(slotProps.data)"
                >
                  Scan
                </button>
                <button
                  class="text-text-secondary hover:text-primary font-medium transition-colors text-sm"
                  @click="viewAuditDiscrepancies(slotProps.data)"
                >
                  View Report
                </button>
              </div>
            </template>
          </Column>
        </DataTable>
      </div>
    </section>

    <!-- Discrepancy Resolution Queue -->
    <h2 class="font-h2 text-h2 text-text-primary mb-4">
      Discrepancy Resolution Queue
      <span class="font-metadata text-metadata text-text-secondary ml-2"
        >({{ openDiscrepancies.length }} open)</span
      >
    </h2>
    <div v-if="loadingDiscrepancies" class="flex justify-center py-12">
      <span class="material-symbols-outlined animate-spin text-primary text-4xl">sync</span>
    </div>
    <div v-else class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-6">
      <div
        v-for="disc in openDiscrepancies"
        :key="disc.id"
        class="bg-surface border border-border-default rounded-xl p-card-padding shadow-[0_4px_4px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform duration-200 flex flex-col h-full"
      >
        <div class="flex justify-between items-start mb-4">
          <div class="flex items-center gap-2">
            <span
              class="material-symbols-outlined p-1.5 rounded-lg"
              :class="
                disc.issueType === 'Missing'
                  ? 'text-status-critical bg-error-container'
                  : disc.issueType === 'Location Mismatch'
                    ? 'text-primary-container bg-accent-subtle'
                    : 'text-tertiary bg-tertiary-fixed'
              "
            >
              {{
                disc.issueType === 'Missing'
                  ? 'warning'
                  : disc.issueType === 'Location Mismatch'
                    ? 'location_off'
                    : 'help_center'
              }}
            </span>
            <span
              class="font-metadata text-metadata font-medium uppercase tracking-wider"
              :class="
                disc.issueType === 'Missing'
                  ? 'text-status-critical'
                  : disc.issueType === 'Location Mismatch'
                    ? 'text-primary-container'
                    : 'text-tertiary'
              "
            >
              {{ disc.issueType }}
            </span>
          </div>
          <span class="font-metadata text-metadata text-text-secondary">
            {{ new Date(disc.createdAt).toLocaleDateString() }}
          </span>
        </div>
        <h3 class="font-h3 text-h3 text-text-primary mb-1">
          {{ disc.assetName || 'Unknown Asset' }}
        </h3>
        <p class="font-mono-data text-mono-data text-text-secondary mb-4">{{ disc.assetTag }}</p>
        <div
          class="bg-surface-subtle p-3 rounded border border-border-default mb-4 flex-1 space-y-1"
        >
          <div class="flex justify-between text-sm">
            <span class="text-text-secondary">Last Known:</span>
            <span class="text-text-primary font-medium">{{ disc.lastKnownLocation || '—' }}</span>
          </div>
          <div v-if="disc.scannedLocation" class="flex justify-between text-sm">
            <span class="text-text-secondary">Scanned At:</span>
            <span class="text-text-primary font-medium">{{ disc.scannedLocation }}</span>
          </div>
          <div v-if="disc.recommendedAction" class="flex justify-between text-sm">
            <span class="text-text-secondary">Recommended:</span>
            <span class="text-text-primary font-medium text-right">{{
              disc.recommendedAction
            }}</span>
          </div>
        </div>
        <div class="mt-auto pt-4 border-t border-border-default flex justify-end gap-2">
          <select
            v-model="selectedActions[disc.id]"
            class="form-select w-full bg-surface border border-[#D6D3CE] rounded font-body text-text-primary py-1.5 px-3 focus:ring-primary focus:border-primary"
          >
            <option disabled value="">Select Action...</option>
            <option
              v-for="action in actionsForDisc(disc)"
              :key="action.value"
              :value="action.value"
            >
              {{ action.label }}
            </option>
          </select>
          <button
            class="bg-primary text-on-primary hover:opacity-90 rounded p-1.5 transition-colors flex-shrink-0 disabled:opacity-50"
            :disabled="!selectedActions[disc.id] || resolvingId === disc.id"
            @click="handleResolveDiscrepancy(disc)"
          >
            <span
              class="material-symbols-outlined"
              :class="{ 'animate-spin': resolvingId === disc.id }"
            >
              {{ resolvingId === disc.id ? 'sync' : 'check' }}
            </span>
          </button>
        </div>
      </div>
      <div
        v-if="!loadingDiscrepancies && openDiscrepancies.length === 0"
        class="col-span-full py-12 text-center"
      >
        <span class="material-symbols-outlined text-5xl text-status-in-stock block mb-3"
          >check_circle</span
        >
        <p class="text-text-secondary italic">All discrepancies resolved.</p>
      </div>
    </div>

    <!-- Create Audit Modal -->
    <Teleport to="body">
      <div
        v-if="showCreateModal"
        class="fixed inset-0 z-50 flex items-center justify-center bg-black/50 backdrop-blur-sm"
        @click.self="showCreateModal = false"
      >
        <div class="bg-surface rounded-2xl shadow-2xl p-8 w-full max-w-md mx-4 animate-in fade-in">
          <h2 class="font-h2 text-h2 text-text-primary mb-6">New Audit Session</h2>
          <div class="space-y-4">
            <div>
              <label
                class="block font-metadata text-metadata text-text-secondary uppercase tracking-wider mb-1"
              >
                Session Title
              </label>
              <input
                v-model="newAudit.title"
                type="text"
                placeholder="e.g. Q2 2026 Full Inventory Audit"
                class="w-full h-11 px-4 bg-canvas border border-border-default rounded-xl focus:border-primary focus:ring-1 focus:ring-primary outline-none"
              />
            </div>
            <div>
              <label
                class="block font-metadata text-metadata text-text-secondary uppercase tracking-wider mb-1"
              >
                Auditor Name
              </label>
              <input
                v-model="newAudit.auditorName"
                type="text"
                placeholder="e.g. John Smith"
                class="w-full h-11 px-4 bg-canvas border border-border-default rounded-xl focus:border-primary focus:ring-1 focus:ring-primary outline-none"
              />
            </div>
            <div>
              <label
                class="block font-metadata text-metadata text-text-secondary uppercase tracking-wider mb-1"
              >
                Start Date
              </label>
              <input
                v-model="newAudit.startDate"
                type="date"
                class="w-full h-11 px-4 bg-canvas border border-border-default rounded-xl focus:border-primary focus:ring-1 focus:ring-primary outline-none"
              />
            </div>
          </div>
          <div class="flex gap-3 mt-6">
            <button
              class="flex-1 py-3 bg-surface border border-border-default rounded-xl font-h3 text-text-secondary hover:bg-surface-subtle"
              @click="showCreateModal = false"
            >
              Cancel
            </button>
            <button
              class="flex-1 py-3 bg-primary text-on-primary rounded-xl font-h3 hover:opacity-90 disabled:opacity-50 flex items-center justify-center gap-2"
              :disabled="
                !newAudit.title || !newAudit.auditorName || !newAudit.startDate || creatingAudit
              "
              @click="handleCreateAudit"
            >
              <span v-if="creatingAudit" class="material-symbols-outlined animate-spin">sync</span>
              {{ creatingAudit ? 'Creating...' : 'Create Session' }}
            </button>
          </div>
        </div>
      </div>
    </Teleport>
  </main>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useRouter } from 'vue-router';
import { useAuth } from '../composables/useAuth';
import { getAudits, createAudit, getDiscrepancies, resolveDiscrepancy } from '../api/audits';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import type { AuditSession, AuditDiscrepancy } from '../types/models';

const router = useRouter();
const { currentUser } = useAuth();

const audits = ref<AuditSession[]>([]);
const loadingAudits = ref(true);
const statusFilter = ref('all');
const allDiscrepancies = ref<AuditDiscrepancy[]>([]);
const loadingDiscrepancies = ref(true);
const selectedActions = ref<Record<number, string>>({});
const resolvingId = ref<number | null>(null);
const showCreateModal = ref(false);
const creatingAudit = ref(false);
const newAudit = ref({ title: '', auditorName: '', startDate: '' });

const filteredAudits = computed(() => {
  if (statusFilter.value === 'all') return audits.value;
  return audits.value.filter((a) => a.status === statusFilter.value);
});

const activeAuditCount = computed(() => audits.value.filter((a) => a.status === 'Active').length);
const openDiscrepancies = computed(() =>
  allDiscrepancies.value.filter((d) => d.status === 'Open' || !d.status),
);
const openDiscrepancyCount = computed(() => openDiscrepancies.value.length);
const resolvedDiscrepancyCount = computed(
  () => allDiscrepancies.value.filter((d) => d.status === 'Resolved').length,
);

function actionsForDisc(disc: AuditDiscrepancy) {
  if (disc.issueType === 'Missing') {
    return [
      { label: 'Confirm Location (if found)', value: 'confirm_location' },
      { label: 'Mark as Lost', value: 'mark_lost' },
      { label: 'Dismiss', value: 'dismiss' },
    ];
  }
  if (disc.issueType === 'Location Mismatch') {
    return [
      { label: `Confirm New Location (${disc.scannedLocation})`, value: 'confirm_location' },
      { label: 'Revert to Original Location', value: 'update_location' },
      { label: 'Dismiss', value: 'dismiss' },
    ];
  }
  if (disc.issueType === 'Unregistered') {
    return [
      { label: 'Register as New Asset', value: 'register' },
      { label: 'Dismiss / Not an Asset', value: 'dismiss' },
    ];
  }
  return [{ label: 'Dismiss', value: 'dismiss' }];
}

const handleStartAudit = () => {
  const activeAudit = audits.value.find((a) => a.status === 'Active');
  if (activeAudit) {
    router.push({ name: 'AuditScanModeMobile', query: { auditId: String(activeAudit.id) } });
  } else {
    void startQuickScanSession();
  }
};

const startQuickScanSession = async () => {
  creatingAudit.value = true;
  try {
    const today = new Date().toISOString().split('T')[0] || '';
    const auditLabelDate = new Date().toLocaleDateString();
    const auditor = currentUser.value?.name || 'Field Auditor';

    const created = await createAudit({
      title: `Quick Scan Session - ${auditLabelDate}`,
      auditorName: auditor,
      startDate: today,
      status: 'Active',
    });

    audits.value.unshift(created as AuditSession);
    router.push({ name: 'AuditScanModeMobile', query: { auditId: String(created.id) } });
  } catch (err) {
    console.error('Failed to start quick scan session:', err);
    alert(err instanceof Error ? err.message : 'Failed to start scanning session.');
    showCreateModal.value = true;
  } finally {
    creatingAudit.value = false;
  }
};

const handleScanAudit = (audit: AuditSession) => {
  router.push({ name: 'AuditScanModeMobile', query: { auditId: String(audit.id) } });
};

const viewAuditDiscrepancies = (audit: AuditSession) => {
  router.push({ name: 'AuditDiscrepancyResolution', query: { auditId: String(audit.id) } });
};

const handleCreateAudit = async () => {
  if (!newAudit.value.title || !newAudit.value.auditorName || !newAudit.value.startDate) return;
  creatingAudit.value = true;
  try {
    const created = await createAudit({
      title: newAudit.value.title,
      auditorName: newAudit.value.auditorName,
      startDate: newAudit.value.startDate,
      status: 'Active',
    });
    audits.value.unshift(created as AuditSession);
    showCreateModal.value = false;
    newAudit.value = { title: '', auditorName: '', startDate: '' };
    // Navigate to scan view for the newly created audit
    router.push({ name: 'AuditScanModeMobile', query: { auditId: String(created.id) } });
  } catch (err) {
    console.error('Failed to create audit:', err);
    alert(err instanceof Error ? err.message : 'Failed to create audit session. Please try again.');
  } finally {
    creatingAudit.value = false;
  }
};

const handleResolveDiscrepancy = async (disc: AuditDiscrepancy) => {
  const action = selectedActions.value[disc.id] as any;
  if (!action) {
    alert('Please select an action first.');
    return;
  }

  resolvingId.value = disc.id;
  try {
    const resolved = await resolveDiscrepancy(disc.id, { action });
    // Update in local state
    const idx = allDiscrepancies.value.findIndex((d) => d.id === disc.id);
    if (idx !== -1) allDiscrepancies.value[idx] = resolved;
    delete selectedActions.value[disc.id];

    // If "register" navigate to add asset
    if (action === 'register') {
      router.push({ name: 'AddNewAssetForm' });
    }
  } catch (err) {
    console.error('Failed to resolve:', err);
    alert('Failed to resolve discrepancy. Please try again.');
  } finally {
    resolvingId.value = null;
  }
};

const fetchAudits = async () => {
  try {
    audits.value = await getAudits();
  } catch (error) {
    console.error('Failed to fetch audits:', error);
  } finally {
    loadingAudits.value = false;
  }
};

const fetchDiscrepancies = async () => {
  try {
    allDiscrepancies.value = await getDiscrepancies();
  } catch (error) {
    console.error('Failed to fetch discrepancies:', error);
  } finally {
    loadingDiscrepancies.value = false;
  }
};

onMounted(() => {
  fetchAudits();
  fetchDiscrepancies();
});
</script>
