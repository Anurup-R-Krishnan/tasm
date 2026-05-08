<template>
  <main class="space-y-section-gap pb-24">
    <!-- Header -->
    <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
      <div>
        <h1 class="text-text-primary mb-1">Asset Registry</h1>
        <p class="text-text-secondary">
          Centralized database of all campus hardware and infrastructure.
        </p>
      </div>
      <div class="flex gap-3">
        <button
          @click="exportToCSV"
          class="bg-surface border border-border-default text-text-primary px-4 py-2 rounded-lg text-sm flex items-center gap-2 hover:bg-surface-subtle transition-colors shadow-sm"
        >
          <span class="material-symbols-outlined text-[18px]">file_download</span>
          Export
        </button>
        <button @click="router.push('/add-asset')" class="btn-primary">
          <span class="material-symbols-outlined">add_circle</span>
          Add Asset
        </button>
      </div>
    </div>

    <!-- Quick Filters -->
    <div class="flex flex-wrap gap-3">
      <button
        v-for="stat in filterStats"
        :key="stat.label"
        class="bg-surface border border-border-default px-4 py-2 rounded-full flex items-center gap-3 text-sm font-medium transition-all shadow-sm hover:border-primary/20"
        :class="
          selectedFilter === stat.label
            ? 'border-primary ring-2 ring-primary/10'
            : 'text-text-secondary'
        "
        @click="selectedFilter = stat.label"
      >
        <span class="w-2 h-2 rounded-full" :class="stat.dotClass"></span>
        {{ stat.label }}
        <span
          class="px-2 py-0.5 bg-slate-100 rounded-full text-[10px] text-text-secondary font-bold"
          >{{ stat.count }}</span
        >
      </button>
    </div>

    <!-- Search & Active Filters -->
    <div class="space-y-4">
      <div class="premium-card !p-4 flex flex-col md:flex-row justify-between items-center gap-4">
        <div class="relative flex-1 w-full max-w-xl">
          <span
            class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-text-secondary"
            >search</span
          >
          <input
            v-model="searchQuery"
            class="w-full bg-surface-subtle border border-border-default rounded-xl py-2 pl-10 pr-4 text-sm focus:bg-white focus:ring-2 focus:ring-primary/20 outline-none transition-all"
            placeholder="Search by name, tag, location, or custodian..."
          />
        </div>
        <div class="flex items-center gap-4 w-full md:w-auto">
          <button
            @click="toggleAdvancedFilters"
            class="flex items-center gap-2 px-4 py-2 text-xs font-bold transition-colors"
            :class="showAdvancedFilters ? 'text-primary' : 'text-text-secondary hover:text-primary'"
          >
            <span class="material-symbols-outlined text-sm">filter_list</span>
            Advanced Filters
          </button>
          <div class="h-6 w-px bg-border-default hidden md:block"></div>
          <div
            class="text-[10px] font-bold text-text-secondary uppercase tracking-widest hidden md:block"
          >
            {{ filteredAssets.length }} assets matching
          </div>
        </div>
      </div>

      <!-- Advanced Filter Panel -->
      <transition
        enter-active-class="transition duration-200 ease-out"
        enter-from-class="transform -translate-y-2 opacity-0"
        enter-to-class="transform translate-y-0 opacity-100"
        leave-active-class="transition duration-150 ease-in"
        leave-from-class="transform translate-y-0 opacity-100"
        leave-to-class="transform -translate-y-2 opacity-0"
      >
        <div
          v-if="showAdvancedFilters"
          class="premium-card !p-6 bg-surface-variant/30 border-primary/10"
        >
          <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
            <div class="space-y-2">
              <label class="text-[10px] font-bold text-text-secondary uppercase tracking-wider"
                >Category</label
              >
              <select
                v-model="selectedCategory"
                class="w-full bg-surface border border-border-default rounded-lg px-3 py-1.5 text-sm outline-none focus:ring-2 focus:ring-primary/20"
              >
                <option value="">All Categories</option>
                <option v-for="cat in uniqueCategories" :key="cat" :value="cat">{{ cat }}</option>
              </select>
            </div>
            <div class="space-y-2">
              <label class="text-[10px] font-bold text-text-secondary uppercase tracking-wider"
                >Location</label
              >
              <select
                v-model="selectedLocationFilter"
                class="w-full bg-surface border border-border-default rounded-lg px-3 py-1.5 text-sm outline-none focus:ring-2 focus:ring-primary/20"
              >
                <option value="">All Locations</option>
                <option v-for="loc in uniqueLocations" :key="loc" :value="loc">{{ loc }}</option>
              </select>
            </div>
            <div class="md:col-span-2 flex items-end justify-end gap-2">
              <button
                @click="showAdvancedFilters = false"
                class="px-4 py-1.5 text-xs font-bold text-text-secondary hover:bg-surface-variant rounded-lg transition-colors"
              >
                Clear Filters
              </button>
              <button class="btn-primary !px-6 !py-1.5 !text-xs">Apply Filters</button>
            </div>
          </div>
        </div>
      </transition>
    </div>

    <!-- Table Container -->
    <div class="premium-card !p-0 overflow-hidden">
      <div class="overflow-x-auto">
        <table class="w-full text-left">
          <thead
            class="bg-surface-subtle/50 text-[10px] font-bold text-text-secondary uppercase tracking-widest border-b border-border-default"
          >
            <tr>
              <th class="px-6 py-4">Asset Detail</th>
              <th class="px-6 py-4">Status</th>
              <th class="px-6 py-4">Custodian</th>
              <th class="px-6 py-4">Location</th>
              <th class="px-6 py-4">Warranty</th>
              <th class="px-6 py-4 text-right">Actions</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-border-default">
            <tr
              v-for="asset in filteredAssets"
              :key="asset.id"
              class="hover:bg-surface-subtle/50 transition-colors group"
            >
              <td class="px-6 py-4">
                <div class="flex items-center gap-3">
                  <div
                    class="w-10 h-10 rounded-xl bg-surface-subtle flex items-center justify-center text-text-disabled"
                  >
                    <span class="material-symbols-outlined">{{ getIcon(asset.category) }}</span>
                  </div>
                  <div>
                    <router-link
                      :to="`/asset/${asset.id}`"
                      class="text-sm font-bold text-text-primary leading-tight group-hover:text-primary transition-colors block"
                    >
                      {{ asset.name }}
                    </router-link>
                    <span
                      class="text-[10px] font-mono text-text-secondary mt-1 block tracking-wider uppercase"
                      >{{ asset.tagId }}</span
                    >
                  </div>
                </div>
              </td>
              <td class="px-6 py-4">
                <span
                  class="text-[10px] font-bold uppercase tracking-wider px-2 py-0.5 rounded-full"
                  :class="getStatusClass(asset.status)"
                >
                  {{ asset.status }}
                </span>
              </td>
              <td class="px-6 py-4">
                <div v-if="asset.custodian" class="flex items-center gap-2">
                  <div
                    class="w-6 h-6 rounded-full bg-primary-container/20 border border-primary-container/30 flex items-center justify-center text-[9px] font-bold text-primary"
                  >
                    {{ asset.custodian.charAt(0) }}
                  </div>
                  <span class="text-xs text-text-secondary font-medium">{{ asset.custodian }}</span>
                </div>
                <span v-else class="text-xs text-text-secondary italic">Unassigned</span>
              </td>
              <td class="px-6 py-4">
                <span class="text-xs text-text-secondary">{{ asset.location }}</span>
              </td>
              <td class="px-6 py-4">
                <span
                  class="text-[10px] font-bold"
                  :class="
                    asset.warrantyStatus === 'Active' ? 'text-status-in-stock' : 'text-rose-500'
                  "
                >
                  {{ asset.warrantyStatus }}
                </span>
                <p class="text-[9px] text-text-secondary mt-1">
                  {{ formatDate(asset.warrantyExpiry) }}
                </p>
              </td>
              <td class="px-6 py-4 text-right">
                <div
                  class="flex justify-end gap-1 opacity-0 group-hover:opacity-100 transition-opacity"
                >
                  <button
                    @click="router.push(`/asset/${asset.id}`)"
                    class="p-2 text-text-secondary hover:text-primary hover:bg-white rounded-lg transition-all shadow-sm"
                  >
                    <span class="material-symbols-outlined text-[18px]">visibility</span>
                  </button>
                  <button
                    @click="router.push(`/add-asset?edit=${asset.id}`)"
                    class="p-2 text-text-secondary hover:text-primary hover:bg-white rounded-lg transition-all shadow-sm"
                  >
                    <span class="material-symbols-outlined text-[18px]">edit</span>
                  </button>
                  <button
                    @click="deleteAsset(asset.id)"
                    class="p-2 text-text-secondary hover:text-rose-600 hover:bg-white rounded-lg transition-all shadow-sm"
                  >
                    <span class="material-symbols-outlined text-[18px]">delete</span>
                  </button>
                </div>
              </td>
            </tr>
            <tr v-if="filteredAssets.length === 0">
              <td colspan="6" class="px-6 py-12 text-center text-text-secondary text-sm italic">
                No assets found matching your criteria.
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { getAssets, deleteAsset as apiDeleteAsset } from '../api/assets';
import type { Asset } from '../types/models';

const router = useRouter();

const assets = ref<Asset[]>([]);
const loading = ref(true);
const searchQuery = ref('');
const selectedFilter = ref('All');
const selectedCategory = ref('');
const selectedLocationFilter = ref('');
const showAdvancedFilters = ref(false);

const fetchAssets = async () => {
  loading.value = true;
  try {
    assets.value = await getAssets();
  } catch (error) {
    console.error('Failed to fetch assets:', error);
  } finally {
    loading.value = false;
  }
};

const filteredAssets = computed(() => {
  let filtered = assets.value;

  if (selectedFilter.value !== 'All') {
    filtered = filtered.filter((a) => a.status === selectedFilter.value);
  }

  if (selectedCategory.value) {
    filtered = filtered.filter((a) => a.category === selectedCategory.value);
  }

  if (selectedLocationFilter.value) {
    filtered = filtered.filter((a) => a.location === selectedLocationFilter.value);
  }

  const query = searchQuery.value.toLowerCase();
  if (query) {
    filtered = filtered.filter(
      (a) =>
        a.name.toLowerCase().includes(query) ||
        a.tagId.toLowerCase().includes(query) ||
        (a.custodian && a.custodian.toLowerCase().includes(query)) ||
        a.location.toLowerCase().includes(query),
    );
  }

  return filtered;
});

const uniqueCategories = computed(() => {
  const cats = new Set(assets.value.map((a) => a.category));
  return Array.from(cats).filter(Boolean).sort();
});

const uniqueLocations = computed(() => {
  const locs = new Set(assets.value.map((a) => a.location));
  return Array.from(locs).filter(Boolean).sort();
});

const filterStats = computed(() => [
  { label: 'All', count: assets.value.length, dotClass: 'bg-text-secondary' },
  {
    label: 'In Stock',
    count: assets.value.filter((a) => a.status === 'In Stock').length,
    dotClass: 'bg-status-in-stock',
  },
  {
    label: 'Checked Out',
    count: assets.value.filter((a) => a.status === 'Checked Out').length,
    dotClass: 'bg-primary',
  },
  {
    label: 'Under Repair',
    count: assets.value.filter((a) => a.status === 'Under Repair').length,
    dotClass: 'bg-status-critical',
  },
]);

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

const getIcon = (cat: string) => {
  if (cat.includes('IT')) return 'laptop_mac';
  if (cat.includes('Furniture')) return 'chair';
  return 'inventory_2';
};

const formatDate = (date: string) => {
  if (!date) return 'N/A';
  return new Date(date).toLocaleDateString('en-GB', {
    day: '2-digit',
    month: 'short',
    year: 'numeric',
  });
};

const deleteAsset = async (id: number) => {
  if (!confirm('Are you sure you want to delete this asset? This action is permanent.')) return;
  try {
    await apiDeleteAsset(id);
    assets.value = assets.value.filter((a) => a.id !== id);
    alert('Asset deleted successfully.');
  } catch (error) {
    console.error('Error deleting asset:', error);
    alert('Failed to delete asset. Please check permissions.');
  }
};

const toggleAdvancedFilters = () => {
  showAdvancedFilters.value = !showAdvancedFilters.value;
};

const exportToCSV = () => {
  const headers = [
    'Name',
    'Tag ID',
    'Category',
    'Status',
    'Custodian',
    'Location',
    'Value',
    'Warranty',
  ];
  const rows = filteredAssets.value.map((a) => [
    a.name,
    a.tagId,
    a.category,
    a.status,
    a.custodian || 'Unassigned',
    a.location,
    a.value,
    a.warrantyStatus,
  ]);

  const csvContent = [
    headers.join(','),
    ...rows.map((row) => row.map((cell) => `"${cell}"`).join(',')),
  ].join('\n');

  const blob = new Blob([csvContent], { type: 'text/csv;charset=utf-8;' });
  const link = document.createElement('a');
  const url = URL.createObjectURL(blob);
  link.setAttribute('href', url);
  link.setAttribute(
    'download',
    `tasm_asset_registry_${new Date().toISOString().split('T')[0]}.csv`,
  );
  link.style.visibility = 'hidden';
  document.body.appendChild(link);
  link.click();
  document.body.removeChild(link);
};

onMounted(fetchAssets);
</script>
