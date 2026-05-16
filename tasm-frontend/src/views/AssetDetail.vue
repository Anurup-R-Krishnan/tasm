<template>
  <main class="space-y-section-gap pb-24">
    <div v-if="loading" class="flex flex-col items-center justify-center py-40 gap-4">
      <span class="material-symbols-outlined animate-spin text-4xl text-primary">sync</span>
      <p class="text-sm font-bold text-text-secondary uppercase tracking-widest">
        Synchronizing Asset Data...
      </p>
    </div>

    <div v-else-if="asset" class="space-y-8 animate-in fade-in slide-in-from-bottom-4 duration-700">
      <!-- Top Navigation & Actions -->
      <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-6">
        <div class="flex items-center gap-4">
          <button
            @click="router.back()"
            class="w-10 h-10 rounded-xl bg-white border border-border-default flex items-center justify-center text-text-secondary hover:text-primary hover:border-primary/20 transition-all shadow-sm"
          >
            <span class="material-symbols-outlined">arrow_back</span>
          </button>
          <div>
            <div class="flex items-center gap-3">
              <h1 class="text-text-primary mb-0 leading-none">{{ asset.name }}</h1>
              <span
                class="text-[10px] font-bold uppercase tracking-wider px-2 py-0.5 rounded-full"
                :class="getStatusClass(asset.status)"
              >
                {{ asset.status }}
              </span>
            </div>
            <p class="text-xs font-mono text-text-secondary mt-1.5 uppercase tracking-widest">
              {{ asset.tagId }}
            </p>
          </div>
        </div>
        <div class="flex flex-wrap gap-3 w-full md:w-auto">
          <button
            class="flex-1 md:flex-none bg-surface border border-border-default text-text-primary px-4 py-2.5 rounded-xl text-sm font-bold flex items-center justify-center gap-2 hover:bg-surface-subtle transition-all shadow-sm"
            @click="editAsset"
          >
            <span class="material-symbols-outlined text-[18px]">edit</span>
            Edit
          </button>

          <button
            v-if="asset.status === 'In Stock'"
            class="flex-1 md:flex-none btn-primary px-4 py-2.5 !text-sm"
            @click="startCheckout"
          >
            <span class="material-symbols-outlined">sync_alt</span>
            Checkout
          </button>

          <button
            v-if="asset.status === 'Checked Out'"
            class="flex-1 md:flex-none bg-metric-amber/10 border border-metric-amber/30 text-amber-800 px-4 py-2.5 rounded-xl text-sm font-bold flex items-center justify-center gap-2 hover:bg-metric-amber/20 transition-all shadow-sm"
            @click="handleCheckin"
          >
            <span class="material-symbols-outlined">assignment_return</span>
            Check-in
          </button>

          <button
            class="flex-1 md:flex-none bg-primary-container/10 border border-primary-container/30 text-primary px-4 py-2.5 rounded-xl text-sm font-bold flex items-center justify-center gap-2 hover:bg-primary-container/20 transition-all shadow-sm"
            @click="showTransferModal = true"
          >
            <span class="material-symbols-outlined">move_up</span>
            Transfer
          </button>

          <button
            v-if="asset.status !== 'End of Life' && asset.status !== 'Disposed'"
            class="flex-1 md:flex-none bg-metric-amber/10 border border-metric-amber/30 text-amber-800 px-4 py-2.5 rounded-xl text-sm font-bold flex items-center justify-center gap-2 hover:bg-metric-amber/20 transition-all shadow-sm"
            @click="handleRetire"
          >
            <span class="material-symbols-outlined">power_settings_new</span>
            Retire
          </button>

          <button
            v-if="asset.status === 'End of Life'"
            class="flex-1 md:flex-none bg-error-container/10 border border-error-container/30 text-status-critical px-4 py-2.5 rounded-xl text-sm font-bold flex items-center justify-center gap-2 hover:bg-error-container/20 transition-all shadow-sm"
            @click="showDisposeModal = true"
          >
            <span class="material-symbols-outlined">delete_forever</span>
            Dispose
          </button>
        </div>
      </div>

      <!-- Transfer Modal -->
      <div
        v-if="showTransferModal"
        class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/60 backdrop-blur-sm animate-in fade-in duration-300"
      >
        <div
          class="bg-surface w-full max-w-lg rounded-[32px] shadow-2xl overflow-hidden animate-in zoom-in-95 duration-300"
        >
          <div class="p-8 border-b border-border-default flex justify-between items-center">
            <div>
              <h2 class="text-2xl font-bold text-text-primary tracking-tight">Transfer Asset</h2>
              <p class="text-xs text-text-secondary mt-1 font-medium">
                Log a physical move or change in custody.
              </p>
            </div>
            <button
              @click="showTransferModal = false"
              class="text-text-secondary hover:text-text-primary transition-colors"
            >
              <span class="material-symbols-outlined">close</span>
            </button>
          </div>

          <div class="p-8 space-y-6">
            <div class="space-y-2">
              <label
                class="block text-[10px] font-black text-text-secondary uppercase tracking-[0.2em] ml-1"
                >Target Location</label
              >
              <select
                v-model="transferForm.newLocation"
                class="w-full h-14 px-5 bg-canvas border border-border-default rounded-2xl text-text-primary focus:outline-none focus:border-primary transition-all font-bold"
              >
                <option value="">Maintain Current ({{ asset.location }})</option>
                <option v-for="loc in locations" :key="loc.id" :value="loc.name">
                  {{ loc.name }}
                </option>
              </select>
            </div>

            <div class="space-y-2">
              <label
                class="block text-[10px] font-black text-text-secondary uppercase tracking-[0.2em] ml-1"
                >New Custodian (Optional)</label
              >
              <input
                v-model="transferForm.newCustodian"
                type="text"
                placeholder="e.g. Finance Dept / John Doe"
                class="w-full h-14 px-5 bg-canvas border border-border-default rounded-2xl text-text-primary focus:outline-none focus:border-primary transition-all font-bold"
              />
            </div>

            <div class="space-y-2">
              <label
                class="block text-[10px] font-black text-text-secondary uppercase tracking-[0.2em] ml-1"
                >Movement Notes</label
              >
              <textarea
                v-model="transferForm.notes"
                rows="3"
                placeholder="Reason for transfer..."
                class="w-full p-5 bg-canvas border border-border-default rounded-2xl text-text-primary focus:outline-none focus:border-primary transition-all font-medium resize-none"
              ></textarea>
            </div>
          </div>

          <div class="p-8 bg-surface-subtle flex gap-4">
            <button
              @click="showTransferModal = false"
              class="flex-1 h-14 rounded-2xl font-bold text-text-secondary hover:bg-white transition-all"
            >
              Cancel
            </button>
            <button
              @click="executeTransfer"
              :disabled="transferring"
              class="flex-[2] h-14 bg-primary text-on-primary rounded-2xl font-bold shadow-xl shadow-primary/20 hover:opacity-90 flex items-center justify-center gap-2"
            >
              <span v-if="transferring" class="material-symbols-outlined animate-spin">sync</span>
              <span>{{ transferring ? 'Transferring...' : 'Execute Transfer' }}</span>
            </button>
          </div>
        </div>
      </div>

      <!-- Dispose Modal -->
      <div
        v-if="showDisposeModal && asset"
        class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/60 backdrop-blur-sm animate-in fade-in duration-300"
      >
        <div
          class="bg-surface w-full max-w-lg rounded-[32px] shadow-2xl overflow-hidden animate-in zoom-in-95 duration-300"
        >
          <div class="p-8 border-b border-border-default flex justify-between items-center">
            <div>
              <h2 class="text-2xl font-bold text-text-primary tracking-tight">Dispose Asset</h2>
              <p class="text-xs text-text-secondary mt-1 font-medium">
                Select a disposal method for {{ asset.name }}.
              </p>
            </div>
            <button
              @click="showDisposeModal = false"
              class="text-text-secondary hover:text-text-primary transition-colors"
            >
              <span class="material-symbols-outlined">close</span>
            </button>
          </div>
          <div class="p-8 grid grid-cols-2 gap-4">
            <button
              v-for="method in ['Sell', 'Scrap', 'Recycle', 'Donate']"
              :key="method"
              @click="handleDispose(method)"
              class="h-20 rounded-2xl border border-border-default bg-canvas hover:bg-primary-container/10 hover:border-primary/20 transition-all flex flex-col items-center justify-center gap-2"
            >
              <span class="material-symbols-outlined text-primary">{{
                method === 'Sell'
                  ? 'sell'
                  : method === 'Scrap'
                    ? 'delete'
                    : method === 'Recycle'
                      ? 'recycling'
                      : 'volunteer_activism'
              }}</span>
              <span class="text-xs font-bold text-text-primary">{{ method }}</span>
            </button>
          </div>
          <div class="p-8 bg-surface-subtle">
            <button
              @click="showDisposeModal = false"
              class="w-full h-14 rounded-2xl font-bold text-text-secondary hover:bg-white transition-all"
            >
              Cancel
            </button>
          </div>
        </div>
      </div>

      <!-- Detail Grid -->
      <div class="grid grid-cols-1 lg:grid-cols-12 gap-8">
        <!-- Left Column: Specs & History -->
        <div class="lg:col-span-8 space-y-8">
          <!-- Primary Info Cards -->
          <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
            <div v-for="info in primarySpecs" :key="info.label" class="premium-card">
              <div
                class="w-10 h-10 rounded-xl flex items-center justify-center mb-4 shadow-sm"
                :class="info.bgClass"
              >
                <span class="material-symbols-outlined" :class="info.iconClass">{{
                  info.icon
                }}</span>
              </div>
              <p class="text-[10px] font-bold text-text-secondary uppercase tracking-widest">
                {{ info.label }}
              </p>
              <p class="text-sm font-bold text-text-primary mt-1">{{ info.value }}</p>
            </div>
          </div>

          <!-- Description / Technical Specs -->
          <div class="premium-card">
            <h3
              class="text-xs font-bold text-text-secondary uppercase tracking-widest mb-6 border-b border-border-default pb-4"
            >
              Asset Profile & Technicals
            </h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-x-12 gap-y-6">
              <div
                v-for="(val, key) in technicalSpecs"
                :key="key"
                class="flex justify-between items-center py-2 border-b border-border-default/50"
              >
                <span class="text-xs font-medium text-text-secondary capitalize">{{ key }}</span>
                <span class="text-xs font-bold text-text-primary">{{ val }}</span>
              </div>
            </div>
          </div>

          <!-- Activity History (Mocked list but real style) -->
          <div class="premium-card !p-0 overflow-hidden">
            <div
              class="p-6 border-b border-border-default bg-surface-subtle/30 flex justify-between items-center"
            >
              <h3 class="text-xs font-bold text-text-secondary uppercase tracking-widest">
                Lifecycle Log
              </h3>
              <button
                @click="scrollToHistory"
                class="text-[10px] font-bold text-primary hover:underline"
              >
                Full Audit Trail
              </button>
            </div>
            <div class="p-6 space-y-6 relative">
              <div class="absolute left-[35px] top-8 bottom-8 w-px bg-border-default"></div>
              <div v-for="log in historyLogs" :key="log.date" class="flex gap-6 relative">
                <div
                  class="w-10 h-10 rounded-full bg-white border-4 border-border-default flex items-center justify-center z-10 shadow-sm"
                >
                  <span class="material-symbols-outlined text-sm text-text-secondary">{{
                    log.icon
                  }}</span>
                </div>
                <div class="flex-1 pb-6 border-b border-border-default last:border-none">
                  <div class="flex justify-between items-start mb-1">
                    <h4 class="text-sm font-bold text-text-primary">{{ log.event }}</h4>
                    <span class="text-[10px] font-medium text-text-secondary">{{ log.date }}</span>
                  </div>
                  <p class="text-xs text-text-secondary leading-relaxed">{{ log.description }}</p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Right Column: Sidebar Stats -->
        <div class="lg:col-span-4 space-y-8">
          <!-- Valuation Card -->
          <div
            class="premium-card bg-gradient-to-br from-primary to-surface-secondary text-white border-none shadow-xl shadow-primary/10"
          >
            <div class="flex justify-between items-start mb-2">
              <p class="text-[10px] font-bold text-primary-container uppercase tracking-widest">
                Current Valuation
              </p>
              <RouterLink
                :to="`/depreciation/${asset.id}`"
                class="text-[10px] font-bold text-white uppercase tracking-widest hover:underline flex items-center gap-1"
              >
                View Ledger <span class="material-symbols-outlined text-xs">arrow_forward</span>
              </RouterLink>
            </div>
            <h2 class="text-4xl font-bold">₹{{ (asset.value || 0).toLocaleString() }}</h2>
            <div class="mt-6 p-4 rounded-xl bg-white/5 border border-white/10 backdrop-blur-md">
              <div class="flex justify-between items-center mb-2">
                <span class="text-[10px] font-medium text-primary-container"
                  >Depreciation Method</span
                >
                <span class="text-[10px] font-bold text-white uppercase tracking-wider">{{
                  asset.depreciationMethod || 'Straight Line'
                }}</span>
              </div>
              <div class="w-full h-1 bg-white/10 rounded-full overflow-hidden">
                <div
                  class="h-full bg-primary-container"
                  :style="{ width: depreciationPercent + '%' }"
                ></div>
              </div>
              <div class="flex justify-between mt-2">
                <span class="text-[10px] text-primary-container"
                  >Useful Life: {{ asset.usefulLifeYears || 5 }}y</span
                >
                <span class="text-[10px] text-primary-container"
                  >Residual: ₹{{ (asset.residualValue || 0).toLocaleString() }}</span
                >
              </div>
            </div>
          </div>

          <!-- Warranty Card -->
          <div class="premium-card space-y-4">
            <h3
              class="text-xs font-bold text-text-secondary uppercase tracking-widest border-b border-border-default pb-4"
            >
              Warranty
            </h3>
            <div class="flex items-center gap-3">
              <div
                class="w-10 h-10 rounded-xl flex items-center justify-center"
                :class="warrantyBadge.bgClass"
              >
                <span class="material-symbols-outlined" :class="warrantyBadge.iconClass">{{
                  warrantyBadge.icon
                }}</span>
              </div>
              <div>
                <span
                  class="text-[10px] font-bold uppercase tracking-wider px-2 py-0.5 rounded-full"
                  :class="warrantyBadge.badgeClass"
                  >{{ warrantyBadge.label }}</span
                >
                <p class="text-xs text-text-secondary mt-1">
                  Expires {{ formatDate(asset.warrantyExpiry || asset.warrantyEndDate) }}
                </p>
              </div>
            </div>
            <div v-if="asset.warrantyProvider" class="text-xs">
              <span class="text-text-secondary">Provider:</span>
              <span class="font-bold text-text-primary">{{ asset.warrantyProvider }}</span>
            </div>
            <div v-if="asset.warrantyTerms" class="text-xs">
              <span class="text-text-secondary">Terms:</span>
              <span class="font-medium text-text-primary">{{ asset.warrantyTerms }}</span>
            </div>
          </div>

          <!-- Lifecycle Status -->
          <div class="premium-card">
            <h3
              class="text-xs font-bold text-text-secondary uppercase tracking-widest border-b border-border-default pb-4 mb-4"
            >
              Lifecycle
            </h3>
            <span
              class="text-[10px] font-bold uppercase tracking-wider px-2.5 py-1 rounded-full"
              :class="getStatusClass(asset.lifecycleStatus || 'Deployed')"
              >{{ asset.lifecycleStatus || 'Deployed' }}</span
            >
            <p v-if="asset.serialNumber" class="text-xs text-text-secondary mt-3">
              S/N:
              <span class="font-mono font-bold text-text-primary">{{ asset.serialNumber }}</span>
            </p>
          </div>

          <!-- Custodian & QR -->
          <div class="premium-card space-y-6">
            <h3
              class="text-xs font-bold text-text-secondary uppercase tracking-widest border-b border-border-default pb-4"
            >
              Assigned Custodian
            </h3>
            <div class="flex items-center gap-4">
              <div
                class="w-12 h-12 rounded-2xl bg-primary-container/10 border border-primary-container/20 flex items-center justify-center text-primary font-bold text-lg uppercase shadow-sm"
              >
                {{ asset.custodian?.charAt(0) || '?' }}
              </div>
              <div>
                <p class="text-sm font-bold text-text-primary">
                  {{ asset.custodian || 'Unassigned' }}
                </p>
                <p
                  class="text-[10px] text-text-secondary font-medium uppercase tracking-wider mt-1"
                >
                  {{ asset.status === 'In Stock' ? 'Ready for Checkout' : 'Currently Holding' }}
                </p>
              </div>
            </div>
            <AssetQrCode v-if="asset.tagId" :value="asset.tagId" label="Asset Tag QR" />
            <div v-else class="p-6 bg-surface-subtle rounded-2xl border border-border-default">
              <p class="text-sm text-text-secondary">No asset tag is available for this record.</p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div v-else class="flex flex-col items-center justify-center py-32 gap-4">
      <span class="material-symbols-outlined text-4xl text-text-secondary">inventory_2</span>
      <p class="text-sm font-bold text-text-secondary uppercase tracking-widest">
        Select an asset to view details.
      </p>
      <button
        class="px-4 py-2 bg-primary text-on-primary rounded-lg text-xs font-bold"
        @click="router.push('/inventory')"
      >
        Go to Asset List
      </button>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue';
import { useRoute, useRouter, RouterLink } from 'vue-router';
import {
  getAssetById,
  getAssetHistory,
  transferAsset,
  checkinAsset,
  retireAsset,
  disposeAsset,
} from '../api/assets';
import { getLocations } from '../api/locations';
import type { Asset, AssetEvent, Location } from '../types/models';
import AssetQrCode from '../components/AssetQrCode.vue';

const route = useRoute();
const router = useRouter();
const asset = ref<Asset | null>(null);
const loading = ref(true);
const events = ref<AssetEvent[]>([]);
const locations = ref<Location[]>([]);
const showDisposeModal = ref(false);

// Transfer State
const showTransferModal = ref(false);
const transferring = ref(false);
const transferForm = ref({
  newLocation: '',
  newCustodian: '',
  notes: '',
});

const fetchAssetDetails = async () => {
  loading.value = true;
  try {
    const id = route.params['id'] as string;
    if (!id) {
      asset.value = null;
      events.value = [];
      return;
    }

    asset.value = await getAssetById(id);
    events.value = await getAssetHistory(id);
    locations.value = await getLocations();
  } catch (error) {
    console.error('Failed to fetch asset details:', error);
  } finally {
    loading.value = false;
  }
};

const handleCheckin = async () => {
  if (!asset.value) return;
  if (!confirm('Are you sure you want to check-in this asset? It will be marked as "In Stock".'))
    return;

  try {
    loading.value = true;
    await checkinAsset(asset.value.id, { notes: 'Asset checked in via detail view.' });
    await fetchAssetDetails();
    alert('Asset checked in successfully.');
  } catch (error) {
    console.error('Check-in failed:', error);
    alert('Failed to check-in asset.');
  } finally {
    loading.value = false;
  }
};

const executeTransfer = async () => {
  if (!asset.value) return;
  transferring.value = true;
  try {
    await transferAsset(asset.value.id, {
      newLocation: transferForm.value.newLocation || undefined,
      newCustodian: transferForm.value.newCustodian || undefined,
      notes: transferForm.value.notes,
    });
    showTransferModal.value = false;
    transferForm.value = { newLocation: '', newCustodian: '', notes: '' };
    await fetchAssetDetails();
    alert('Asset transferred successfully.');
  } catch (error) {
    console.error('Transfer failed:', error);
    alert('Failed to transfer asset.');
  } finally {
    transferring.value = false;
  }
};

const handleRetire = async () => {
  if (!asset.value) return;
  const reason = prompt('Reason for retirement:');
  if (!reason) return;
  try {
    loading.value = true;
    await retireAsset(asset.value.id, { reason, notes: '' });
    await fetchAssetDetails();
    alert('Asset retired successfully.');
  } catch (error) {
    console.error('Retire failed:', error);
    alert('Failed to retire asset.');
  } finally {
    loading.value = false;
  }
};

const handleDispose = async (method: string) => {
  if (!asset.value) return;
  try {
    loading.value = true;
    await disposeAsset(asset.value.id, {
      disposalMethod: method,
      residualValue: asset.value.residualValue || 0,
      environmentalCompliance: true,
      dataWipingConfirmed: true,
      notes: '',
    });
    showDisposeModal.value = false;
    await fetchAssetDetails();
    alert('Asset disposed successfully.');
  } catch (error) {
    console.error('Dispose failed:', error);
    alert('Failed to dispose asset.');
  } finally {
    loading.value = false;
  }
};

const scrollToHistory = () => {
  document.querySelector('[data-section="lifecycle-log"]')?.scrollIntoView({ behavior: 'smooth' });
};

const depreciationPercent = computed(() => {
  if (!asset.value || !asset.value.usefulLifeYears) return 50;
  const purchaseDate = new Date(asset.value.purchaseDate);
  const yearsElapsed = (Date.now() - purchaseDate.getTime()) / (365.25 * 24 * 60 * 60 * 1000);
  return Math.min(100, Math.max(0, (yearsElapsed / asset.value.usefulLifeYears) * 100));
});

const warrantyBadge = computed(() => {
  if (!asset.value)
    return {
      label: 'N/A',
      icon: 'verified',
      bgClass: 'bg-surface-variant',
      iconClass: 'text-text-secondary',
      badgeClass: 'bg-surface-variant text-text-secondary',
    };
  const expiry = new Date(asset.value.warrantyExpiry || asset.value.warrantyEndDate);
  const now = new Date();
  const daysLeft = (expiry.getTime() - now.getTime()) / (1000 * 60 * 60 * 24);
  if (daysLeft < 0)
    return {
      label: 'Expired',
      icon: 'gpp_bad',
      bgClass: 'bg-error-container/20',
      iconClass: 'text-status-critical',
      badgeClass: 'bg-status-critical/20 text-status-critical border border-status-critical/30',
    };
  if (daysLeft < 90)
    return {
      label: 'Expiring Soon',
      icon: 'warning',
      bgClass: 'bg-metric-amber',
      iconClass: 'text-amber-800',
      badgeClass: 'bg-metric-amber text-amber-800 border border-amber-300',
    };
  return {
    label: 'Active',
    icon: 'verified',
    bgClass: 'bg-metric-sage',
    iconClass: 'text-status-in-stock',
    badgeClass: 'bg-status-in-stock/20 text-status-in-stock border border-status-in-stock/30',
  };
});

const primarySpecs = computed(() => {
  if (!asset.value) return [];
  return [
    {
      label: 'Category',
      value: asset.value.category,
      icon: 'category',
      bgClass: 'bg-primary-container/20',
      iconClass: 'text-primary',
    },
    {
      label: 'Location',
      value: asset.value.location,
      icon: 'location_on',
      bgClass: 'bg-metric-sage',
      iconClass: 'text-status-in-stock',
    },
    {
      label: 'Book Value',
      value: '₹' + (asset.value.value || 0).toLocaleString(),
      icon: 'account_balance_wallet',
      bgClass: 'bg-primary-container/10',
      iconClass: 'text-accent',
    },
  ];
});

const technicalSpecs = computed(() => {
  if (!asset.value) return {};
  return {
    'Purchase Date': formatDate(asset.value.purchaseDate),
    'Purchase Price': '₹' + (asset.value.purchasePrice || asset.value.value || 0).toLocaleString(),
    'Warranty Expiry': formatDate(asset.value.warrantyExpiry),
    Category: asset.value.category,
    Location: asset.value.location,
    Custodian: asset.value.custodian || 'Unassigned',
    Status: asset.value.status,
    'Serial Number': asset.value.serialNumber || 'N/A',
    'Depreciation Method': asset.value.depreciationMethod || 'Straight Line',
  };
});

const historyLogs = computed(() => {
  return events.value.map((e) => ({
    event: e.eventType.replace(/_/g, ' ').replace(/\b\w/g, (l) => l.toUpperCase()),
    date: new Date(e.createdAt).toLocaleDateString('en-GB', {
      day: '2-digit',
      month: 'short',
      year: 'numeric',
    }),
    description: e.notes || e.description || `${e.previousStatus || ''} → ${e.newStatus || ''}`,
    actor: e.actorName || 'System',
    icon: ['CHECKOUT', 'CHECKIN', 'TRANSFER'].includes(e.eventType)
      ? 'sync_alt'
      : ['RETIRED', 'DISPOSED'].includes(e.eventType)
        ? 'power_settings_new'
        : 'history',
  }));
});

const getStatusClass = (status: string) => {
  switch (status) {
    case 'In Stock':
    case 'Deployed':
      return 'bg-status-in-stock/20 text-status-in-stock border border-status-in-stock/30';
    case 'Checked Out':
      return 'bg-status-checked-out/20 text-status-checked-out border border-status-checked-out/30';
    case 'End of Life':
    case 'Under Repair':
    case 'Under Maintenance':
      return 'bg-status-critical/10 text-status-critical border border-status-critical/30';
    case 'Disposed':
      return 'bg-surface-variant text-text-secondary border border-border-default';
    default:
      return 'bg-surface-variant text-text-secondary';
  }
};

const formatDate = (date: string) => {
  if (!date) return 'N/A';
  return new Date(date).toLocaleDateString('en-GB', {
    day: '2-digit',
    month: 'short',
    year: 'numeric',
  });
};

const editAsset = () => {
  if (!asset.value) return;
  router.push(`/add-asset?edit=${asset.value.id}`);
};

const startCheckout = () => {
  if (!asset.value) return;
  router.push({ name: 'AssetCheckOutFlow', query: { assetId: String(asset.value.id) } });
};

onMounted(fetchAssetDetails);
watch(() => route.params['id'], fetchAssetDetails);
</script>
