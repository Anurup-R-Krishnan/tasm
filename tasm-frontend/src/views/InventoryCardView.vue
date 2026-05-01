<template>
  <main class="p-page-margin max-w-[1400px] mx-auto space-y-section-gap pb-24">
    <!-- Page Header & Actions -->
    <div class="flex flex-col md:flex-row md:items-end justify-between gap-4">
      <div>
        <h1 class="text-text-primary mb-1">Inventory Management</h1>
        <p class="text-text-secondary">Manage and track assets across all campus locations.</p>
      </div>
      <div class="flex items-center gap-3">
        <button
          class="bg-surface border border-border-default text-text-primary px-4 py-2 rounded-lg text-sm flex items-center gap-2 hover:bg-surface-subtle transition-colors shadow-sm"
        >
          <span class="material-symbols-outlined text-[18px]">filter_list</span>
          Filters
        </button>
        <button class="btn-primary">
          <span class="material-symbols-outlined">add_circle</span>
          Add Asset
        </button>
      </div>
    </div>

    <!-- Status Summary Pills -->
    <div class="flex flex-wrap gap-3">
      <button
        v-for="stat in statusStats"
        :key="stat.label"
        class="bg-surface border border-border-default px-4 py-2 rounded-full flex items-center gap-2 text-sm font-medium text-text-primary hover:bg-surface-subtle transition-colors shadow-sm"
      >
        <span class="w-2 h-2 rounded-full" :class="stat.dotClass"></span>
        {{ stat.label }}
        <span class="ml-1 px-1.5 py-0.5 bg-slate-100 rounded-full text-xs text-slate-600">{{
          stat.count
        }}</span>
      </button>
    </div>

    <!-- Hybrid Layout: Sidebar + Grid -->
    <div class="flex flex-col lg:flex-row gap-8">
      <!-- Category Sidebar -->
      <aside class="w-full lg:w-[260px] shrink-0 space-y-6">
        <div class="premium-card !p-4">
          <h3 class="text-sm font-bold text-slate-400 uppercase tracking-wider mb-4 px-2">
            Categories
          </h3>
          <ul class="space-y-1">
            <li v-for="(count, category) in categories" :key="category">
              <button
                class="w-full text-left px-3 py-2 rounded-xl text-sm font-medium flex justify-between items-center transition-colors"
                :class="
                  selectedCategory === category
                    ? 'bg-indigo-50 text-indigo-600'
                    : 'text-slate-600 hover:bg-slate-50'
                "
                @click="selectedCategory = category"
              >
                <span>{{ category }}</span>
                <span class="text-[10px] bg-slate-100 px-2 py-0.5 rounded-full text-slate-500">{{
                  count
                }}</span>
              </button>
            </li>
          </ul>
        </div>

        <div class="premium-card !p-4">
          <h3 class="text-sm font-bold text-slate-400 uppercase tracking-wider mb-4 px-2">
            Locations
          </h3>
          <ul class="space-y-1">
            <li v-for="loc in locations" :key="loc.id">
              <label
                class="flex items-center gap-3 px-3 py-2 cursor-pointer hover:bg-slate-50 rounded-xl transition-colors"
              >
                <input
                  v-model="selectedLocations"
                  :value="loc.name"
                  type="checkbox"
                  class="rounded border-slate-300 text-indigo-600 focus:ring-indigo-500/20"
                />
                <span class="text-sm text-slate-600">{{ loc.name }}</span>
              </label>
            </li>
          </ul>
        </div>
      </aside>

      <!-- Asset Grid -->
      <div class="flex-1">
        <div v-if="loading" class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-6">
          <div v-for="i in 6" :key="i" class="premium-card h-64 animate-pulse bg-slate-50"></div>
        </div>
        <div
          v-else-if="filteredAssets.length === 0"
          class="premium-card flex flex-col items-center justify-center py-20 text-center"
        >
          <span class="material-symbols-outlined text-6xl text-slate-200 mb-4">inventory_2</span>
          <h3 class="text-slate-900">No assets found</h3>
          <p class="text-slate-500">Try adjusting your filters or search query.</p>
        </div>
        <div v-else class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-6">
          <div
            v-for="asset in filteredAssets"
            :key="asset.id"
            class="premium-card !p-0 overflow-hidden group cursor-pointer"
            @click="router.push(`/asset/${asset.id}`)"
          >
            <div
              class="h-44 bg-slate-100 relative overflow-hidden flex items-center justify-center"
            >
              <span
                class="material-symbols-outlined text-slate-300 text-6xl group-hover:scale-110 transition-transform duration-500"
              >
                {{ getIcon(asset.category) }}
              </span>
              <div class="absolute top-3 right-3">
                <span
                  class="text-[10px] font-bold uppercase tracking-wider px-2 py-1 rounded-lg shadow-sm"
                  :class="getStatusClass(asset.status)"
                >
                  {{ asset.status }}
                </span>
              </div>
            </div>
            <div class="p-5 flex flex-col">
              <div class="mb-3">
                <h3
                  class="text-slate-900 font-bold leading-tight group-hover:text-indigo-600 transition-colors line-clamp-1"
                >
                  {{ asset.name }}
                </h3>
                <span class="text-[11px] font-mono text-slate-400 mt-1 block tracking-wider">{{
                  asset.tagId
                }}</span>
              </div>

              <div class="mt-auto pt-4 border-t border-slate-50 flex justify-between items-center">
                <div class="flex items-center gap-1.5 text-slate-500 text-xs">
                  <span class="material-symbols-outlined text-sm">location_on</span>
                  {{ asset.location }}
                </div>
                <div class="text-sm font-bold text-slate-900">
                  ₹ {{ (asset.value || 0).toLocaleString() }}
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Pagination (Mocked but functional state) -->
        <div
          v-if="!loading && filteredAssets.length > 0"
          class="mt-8 flex items-center justify-between border-t border-border-default pt-6"
        >
          <span class="text-xs font-medium text-slate-500"
            >Showing {{ filteredAssets.length }} assets</span
          >
          <div class="flex gap-2">
            <button
              class="px-4 py-1.5 rounded-lg border border-slate-200 text-xs font-bold text-slate-400 cursor-not-allowed"
            >
              Previous
            </button>
            <button
              class="px-4 py-1.5 rounded-lg bg-indigo-600 text-white text-xs font-bold shadow-lg shadow-indigo-600/20"
            >
              1
            </button>
            <button
              class="px-4 py-1.5 rounded-lg border border-slate-200 text-xs font-bold text-slate-600 hover:bg-slate-50"
            >
              Next
            </button>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();
const assets = ref<any[]>([]);
const locations = ref<any[]>([]);
const loading = ref(true);
const selectedCategory = ref('All');
const selectedLocations = ref<string[]>([]);

const fetchInventoryData = async () => {
  loading.value = true;
  try {
    const [assetsRes, locationsRes] = await Promise.all([
      fetch('http://localhost:8080/api/assets'),
      fetch('http://localhost:8080/api/locations'),
    ]);
    if (assetsRes.ok) assets.value = await assetsRes.json();
    if (locationsRes.ok) locations.value = await locationsRes.json();
  } catch (err) {
    console.error('Failed to fetch inventory:', err);
  } finally {
    loading.value = false;
  }
};

const categories = computed(() => {
  const counts: Record<string, number> = { All: assets.value.length };
  assets.value.forEach((a) => {
    counts[a.category] = (counts[a.category] || 0) + 1;
  });
  return counts;
});

const statusStats = computed(() => [
  { label: 'All Assets', count: assets.value.length, dotClass: 'bg-indigo-500' },
  {
    label: 'In Stock',
    count: assets.value.filter((a) => a.status === 'In Stock').length,
    dotClass: 'bg-emerald-500',
  },
  {
    label: 'Checked Out',
    count: assets.value.filter((a) => a.status === 'Checked Out').length,
    dotClass: 'bg-blue-500',
  },
  {
    label: 'Under Repair',
    count: assets.value.filter((a) => a.status === 'Under Repair').length,
    dotClass: 'bg-rose-500',
  },
]);

const filteredAssets = computed(() => {
  return assets.value.filter((a) => {
    const catMatch = selectedCategory.value === 'All' || a.category === selectedCategory.value;
    const locMatch =
      selectedLocations.value.length === 0 || selectedLocations.value.includes(a.location);
    return catMatch && locMatch;
  });
});

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

const getIcon = (category: string) => {
  if (category.includes('IT')) return 'laptop_mac';
  if (category.includes('Furniture')) return 'chair';
  if (category.includes('Vehicle')) return 'directions_car';
  if (category.includes('HVAC')) return 'ac_unit';
  return 'inventory_2';
};

onMounted(fetchInventoryData);
</script>
