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
        <div class="flex gap-3 w-full md:w-auto">
          <button
            class="flex-1 md:flex-none bg-surface border border-border-default text-text-primary px-6 py-2.5 rounded-xl text-sm font-bold flex items-center justify-center gap-2 hover:bg-surface-subtle transition-all shadow-sm"
          >
            <span class="material-symbols-outlined text-[18px]">edit</span>
            Edit Record
          </button>
          <button class="flex-1 md:flex-none btn-primary px-6 py-2.5 !text-sm">
            <span class="material-symbols-outlined">sync_alt</span>
            Checkout / Transfer
          </button>
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
            <p class="text-[10px] font-bold text-primary-container uppercase tracking-widest mb-2">
              Current Valuation
            </p>
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
            <div
              class="p-6 bg-surface-subtle rounded-2xl flex flex-col items-center gap-4 border border-border-default"
            >
              <div
                class="w-32 h-32 bg-white rounded-xl flex items-center justify-center border border-border-default p-2"
              >
                <span class="material-symbols-outlined text-6xl text-text-disabled">qr_code_2</span>
              </div>
              <button
                class="text-[10px] font-bold text-primary uppercase tracking-widest hover:underline"
              >
                Download Asset Tag
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';

const route = useRoute();
const router = useRouter();
const asset = ref<any>(null);
const loading = ref(true);

const fetchAsset = async () => {
  try {
    const res = await fetch(`http://localhost:8080/api/assets/${route.params['id']}`);
    if (res.ok) asset.value = await res.json();
  } catch (error) {
    console.error('Failed to fetch asset:', error);
  } finally {
    loading.value = false;
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
    'Asset Model': 'Enterprise Edition v2',
    Manufacturer: 'Dell Technopark Systems',
    'Network Status': 'Connected',
    'Last Audit': '12 Oct 2025',
  };
});

const historyLogs = [
  {
    event: 'Maintenance Completed',
    date: '24 Oct 2025',
    description:
      'Annual hardware calibration and component cleaning performed by onsite technician.',
    icon: 'handyman',
  },
  {
    event: 'Asset Checked Out',
    date: '15 Aug 2025',
    description: 'Assigned to Rajesh Kumar for Product Design project. Expected return: Feb 2026.',
    icon: 'sync_alt',
  },
  {
    event: 'Initial Procurement',
    date: '12 Jan 2025',
    description: 'Asset registered and tagged after PO-8821 verification.',
    icon: 'inventory',
  },
];

const getStatusClass = (status: string) => {
  switch (status) {
    case 'In Stock':
      return 'bg-emerald-100 text-emerald-700';
    case 'Checked Out':
      return 'bg-blue-100 text-blue-700';
    case 'Under Repair':
      return 'bg-rose-100 text-rose-700';
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

onMounted(fetchAsset);
</script>
