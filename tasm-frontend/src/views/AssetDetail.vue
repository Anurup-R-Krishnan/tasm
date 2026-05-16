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
              <button class="text-[10px] font-bold text-primary hover:underline">
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
                View Ledger
                <span class="material-symbols-outlined text-xs">arrow_forward</span>
              </RouterLink>
            </div>
            <h2 class="text-4xl font-bold">₹{{ (asset.value || 0).toLocaleString() }}</h2>
            <div class="mt-6 p-4 rounded-xl bg-white/5 border border-white/10 backdrop-blur-md">
              <div class="flex justify-between items-center mb-2">
                <span class="text-[10px] font-medium text-primary-container"
                  >Depreciation Method</span
                >
                <span class="text-[10px] font-bold text-white uppercase tracking-wider"
                  >Straight Line</span
                >
              </div>
              <div class="w-full h-1 bg-white/10 rounded-full overflow-hidden">
                <div class="h-full bg-primary-container" style="width: 65%"></div>
              </div>
            </div>
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
import { getAssetById, getAssetHistory, transferAsset, checkinAsset } from '../api/assets';
import { getLocations } from '../api/locations';
import type { Asset, AssetEvent, Location } from '../types/models';
import AssetQrCode from '../components/AssetQrCode.vue';

const route = useRoute();
const router = useRouter();
const asset = ref<Asset | null>(null);
const loading = ref(true);
const events = ref<AssetEvent[]>([]);
const locations = ref<Location[]>([]);

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
      label: 'Warranty',
      value: asset.value.warrantyStatus,
      icon: 'verified',
      bgClass: 'bg-primary-container/10',
      iconClass: 'text-accent',
    },
  ];
});

const technicalSpecs = computed(() => {
  if (!asset.value) return {};
  return {
    'Purchase Date': formatDate(asset.value.purchaseDate),
    'Warranty Expiry': formatDate(asset.value.warrantyExpiry),
    Category: asset.value.category,
    Location: asset.value.location,
    Custodian: asset.value.custodian || 'Unassigned',
    Status: asset.value.status,
  };
});

const historyLogs = computed(() => {
  return events.value.map((e) => ({
    event: e.eventType.replace('_', ' ').replace(/\b\w/g, (l) => l.toUpperCase()),
    date: new Date(e.createdAt).toLocaleDateString('en-GB', {
      day: '2-digit',
      month: 'short',
      year: 'numeric',
    }),
    description: e.notes || `${e.previousStatus} → ${e.newStatus}`,
    icon:
      e.eventType === 'checkout' || e.eventType === 'checkin' || e.eventType === 'transfer'
        ? 'sync_alt'
        : 'history',
  }));
});

const getStatusClass = (status: string) => {
  switch (status) {
    case 'In Stock':
      return 'bg-status-in-stock/20 text-status-in-stock border border-status-in-stock/30';
    case 'Checked Out':
      return 'bg-status-checked-out/20 text-status-checked-out border border-status-checked-out/30';
    case 'Under Repair':
      return 'bg-status-critical/10 text-status-critical border border-status-critical/30';
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
