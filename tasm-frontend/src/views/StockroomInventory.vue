<template>
  <main class="p-page-margin max-w-[1400px] mx-auto space-y-section-gap">
    <!-- Header -->
    <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
      <div>
        <h1 class="text-text-primary">Stockroom & Storage</h1>
        <p class="text-text-secondary mt-1">
          Detailed breakdown of assets by location and facility status.
        </p>
      </div>
      <div class="flex gap-3">
        <button
          class="bg-surface border border-border-default text-text-primary px-4 py-2 rounded-lg text-sm flex items-center gap-2 hover:bg-surface-subtle transition-colors shadow-sm"
        >
          <span class="material-symbols-outlined text-[18px]">print</span>
          Reports
        </button>
        <button class="btn-primary">
          <span class="material-symbols-outlined">add_location</span>
          New Location
        </button>
      </div>
    </div>

    <!-- Summary Widgets -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
      <div
        v-for="stat in summaryStats"
        :key="stat.label"
        class="premium-card flex items-center gap-4"
      >
        <div class="w-12 h-12 rounded-2xl flex items-center justify-center" :class="stat.bgClass">
          <span class="material-symbols-outlined" :class="stat.iconClass">{{ stat.icon }}</span>
        </div>
        <div>
          <p class="text-[10px] font-bold text-text-secondary uppercase tracking-widest">
            {{ stat.label }}
          </p>
          <h2 class="text-2xl font-bold text-text-primary mt-0.5">{{ stat.value }}</h2>
        </div>
      </div>
    </div>

    <!-- Dual Panel Layout -->
    <div class="grid grid-cols-1 lg:grid-cols-12 gap-8">
      <!-- Left Panel: Location Tree/List -->
      <div class="lg:col-span-4 space-y-6">
        <div class="premium-card !p-0 overflow-hidden">
          <div
            class="p-4 border-b border-border-default flex justify-between items-center bg-surface-subtle/50"
          >
            <h3 class="text-xs font-bold text-text-secondary uppercase tracking-wider">
              Campus Facilities
            </h3>
            <span
              class="text-[10px] font-bold bg-white px-2 py-0.5 rounded-full border border-border-default text-text-secondary"
              >{{ locations.length }} Total</span
            >
          </div>
          <div class="divide-y divide-slate-50">
            <button
              v-for="loc in locations"
              :key="loc.id"
              class="w-full p-4 text-left hover:bg-surface-subtle transition-colors flex items-center gap-4 group"
              :class="
                selectedLocation?.id === loc.id
                  ? 'bg-primary-container/10 border-r-4 border-primary'
                  : ''
              "
              @click="selectedLocation = loc"
            >
              <div
                class="w-10 h-10 rounded-xl bg-white border border-border-default flex items-center justify-center group-hover:scale-110 transition-transform shadow-sm"
              >
                <span
                  class="material-symbols-outlined text-text-secondary group-hover:text-primary"
                  >{{ getLocIcon(loc.type) }}</span
                >
              </div>
              <div class="flex-1 min-w-0">
                <p class="text-sm font-bold text-text-primary truncate">{{ loc.name }}</p>
                <p class="text-[11px] text-text-secondary truncate mt-0.5">{{ loc.address }}</p>
              </div>
              <span
                class="material-symbols-outlined text-slate-300 group-hover:translate-x-1 transition-transform"
                >chevron_right</span
              >
            </button>
          </div>
        </div>
      </div>

      <!-- Right Panel: Detailed Inventory in Location -->
      <div class="lg:col-span-8">
        <div v-if="selectedLocation" class="space-y-6">
          <div
            class="premium-card !p-6 bg-gradient-to-br from-primary to-surface-secondary text-white border-none shadow-xl shadow-primary/10"
          >
            <div class="flex justify-between items-start">
              <div>
                <span
                  class="text-[10px] font-bold uppercase tracking-widest text-primary-container opacity-80"
                  >{{ selectedLocation.type }} Profile</span
                >
                <h2 class="text-3xl font-bold mt-1">{{ selectedLocation.name }}</h2>
                <div class="flex items-center gap-4 mt-4">
                  <div
                    class="flex items-center gap-1.5 bg-white/10 px-3 py-1.5 rounded-lg backdrop-blur-md"
                  >
                    <span class="material-symbols-outlined text-sm">groups</span>
                    <span class="text-xs font-bold">{{ selectedLocation.capacity }} Capacity</span>
                  </div>
                  <div
                    class="flex items-center gap-1.5 bg-white/10 px-3 py-1.5 rounded-lg backdrop-blur-md"
                  >
                    <span class="material-symbols-outlined text-sm">inventory</span>
                    <span class="text-xs font-bold">{{ locationAssets.length }} Assets</span>
                  </div>
                </div>
              </div>
              <div
                class="w-16 h-16 rounded-2xl bg-white/10 backdrop-blur-md flex items-center justify-center"
              >
                <span class="material-symbols-outlined text-4xl text-white">{{
                  getLocIcon(selectedLocation.type)
                }}</span>
              </div>
            </div>
          </div>

          <!-- Assets in this location Table -->
          <div class="premium-card !p-0 overflow-hidden">
            <div
              class="p-4 border-b border-border-default flex justify-between items-center bg-surface-subtle/50"
            >
              <h3 class="text-xs font-bold text-text-secondary uppercase tracking-wider">
                Inventory Breakdown
              </h3>
            </div>
            <div class="overflow-x-auto">
              <table class="w-full text-left">
                <thead
                  class="bg-surface-subtle/30 text-[10px] font-bold text-text-secondary uppercase tracking-widest border-b border-border-default"
                >
                  <tr>
                    <th class="px-6 py-4">Asset Detail</th>
                    <th class="px-6 py-4">Category</th>
                    <th class="px-6 py-4">Status</th>
                    <th class="px-6 py-4">Value</th>
                    <th class="px-6 py-4 text-right">Actions</th>
                  </tr>
                </thead>
                <tbody class="divide-y divide-slate-50">
                  <tr
                    v-for="asset in locationAssets"
                    :key="asset.id"
                    class="hover:bg-surface-subtle/50 transition-colors group"
                  >
                    <td class="px-6 py-4">
                      <div class="flex items-center gap-3">
                        <div
                          class="w-8 h-8 rounded-lg bg-slate-100 flex items-center justify-center text-text-secondary"
                        >
                          <span class="material-symbols-outlined text-lg">{{
                            getAssetIcon(asset.category)
                          }}</span>
                        </div>
                        <div>
                          <p class="text-sm font-bold text-text-primary leading-none">
                            {{ asset.name }}
                          </p>
                          <p
                            class="text-[10px] font-mono text-text-secondary mt-1.5 tracking-wider"
                          >
                            {{ asset.tagId }}
                          </p>
                        </div>
                      </div>
                    </td>
                    <td class="px-6 py-4">
                      <span class="text-xs font-medium text-slate-600">{{ asset.category }}</span>
                    </td>
                    <td class="px-6 py-4">
                      <span
                        class="text-[10px] font-bold uppercase tracking-wider px-2 py-0.5 rounded-full"
                        :class="getStatusClass(asset.status)"
                      >
                        {{ asset.status }}
                      </span>
                    </td>
                    <td class="px-6 py-4 text-sm font-bold text-text-primary">
                      ₹{{ (asset.value || 0).toLocaleString() }}
                    </td>
                    <td class="px-6 py-4 text-right">
                      <button
                        class="p-2 hover:bg-white rounded-lg text-text-secondary hover:text-primary transition-all shadow-sm"
                      >
                        <span class="material-symbols-outlined">arrow_forward</span>
                      </button>
                    </td>
                  </tr>
                  <tr v-if="locationAssets.length === 0">
                    <td
                      colspan="5"
                      class="px-6 py-12 text-center text-text-secondary text-sm italic"
                    >
                      No assets registered at this location.
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
        <div
          v-else
          class="h-[500px] premium-card flex flex-col items-center justify-center text-center"
        >
          <span class="material-symbols-outlined text-7xl text-surface-variant mb-4 animate-pulse"
            >location_away</span
          >
          <h3 class="text-text-secondary font-medium">
            Select a facility to view inventory details
          </h3>
        </div>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';

const locations = ref<any[]>([]);
const assets = ref<any[]>([]);
const selectedLocation = ref<any>(null);
const loading = ref(true);

const fetchData = async () => {
  loading.value = true;
  try {
    const [locRes, assetRes] = await Promise.all([
      fetch('http://localhost:8080/api/locations'),
      fetch('http://localhost:8080/api/assets'),
    ]);
    if (locRes.ok) {
      locations.value = await locRes.json();
      if (locations.value.length > 0) selectedLocation.value = locations.value[0];
    }
    if (assetRes.ok) assets.value = await assetRes.json();
  } catch (err) {
    console.error('Failed to fetch stockroom data:', err);
  } finally {
    loading.value = false;
  }
};

const locationAssets = computed(() => {
  if (!selectedLocation.value) return [];
  return assets.value.filter((a) => a.location === selectedLocation.value.name);
});

const summaryStats = computed(() => [
  {
    label: 'Total Facilities',
    value: locations.value.length,
    icon: 'corporate_fare',
    bgClass: 'bg-primary-container/10',
    iconClass: 'text-primary',
  },
  {
    label: 'Asset Capacity',
    value: locations.value.reduce((sum, l) => sum + (l.capacity || 0), 0),
    icon: 'person_search',
    bgClass: 'bg-emerald-50',
    iconClass: 'text-emerald-500',
  },
  {
    label: 'Active Alerts',
    value: 0,
    icon: 'warning',
    bgClass: 'bg-rose-50',
    iconClass: 'text-rose-500',
  },
]);

const getLocIcon = (type: string) => {
  if (type.includes('Office')) return 'business';
  if (type.includes('Data')) return 'dns';
  if (type.includes('Warehouse')) return 'warehouse';
  return 'location_on';
};

const getAssetIcon = (cat: string) => {
  if (cat.includes('IT')) return 'laptop_mac';
  if (cat.includes('Furniture')) return 'chair';
  return 'inventory_2';
};

const getStatusClass = (status: string) => {
  switch (status) {
    case 'In Stock':
      return 'bg-emerald-100 text-emerald-700';
    case 'Checked Out':
      return 'bg-blue-100 text-blue-700';
    case 'Under Repair':
      return 'bg-rose-100 text-rose-700';
    default:
      return 'bg-slate-100 text-slate-700';
  }
};

onMounted(fetchData);
</script>
