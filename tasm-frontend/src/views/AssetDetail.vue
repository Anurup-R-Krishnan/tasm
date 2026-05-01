<template>
  <main class="md:pl-[248px] min-h-screen">
    <div v-if="loading" class="flex justify-center items-center py-20 text-stone-500">
      <span class="material-symbols-outlined animate-spin mr-2">refresh</span> Loading Asset...
    </div>
    <div v-else-if="asset" class="p-page-margin max-w-[1400px] mx-auto">
      <!-- Breadcrumbs & Header -->
      <div class="mb-section-gap">
        <div
          class="flex items-center gap-2 text-metadata text-text-secondary font-metadata mb-stack"
        >
          <router-link to="/asset-registry" class="hover:text-primary transition-colors">
            Asset Inventory
          </router-link>
          <span class="material-symbols-outlined text-[14px]"> chevron_right </span>
          <a class="hover:text-primary transition-colors" href="#">
            {{ asset.category }}
          </a>
          <span class="material-symbols-outlined text-[14px]"> chevron_right </span>
          <span class="text-text-primary font-medium">
            {{ asset.tagId }}
          </span>
        </div>
        <div class="flex justify-between items-start">
          <div>
            <div class="flex items-center gap-inline mb-2">
              <h1 class="font-h1 text-h1 text-text-primary">
                {{ asset.name }}
              </h1>
              <span
                class="inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium border"
                :class="{
                  'bg-metric-sage text-status-in-stock border-[#BBF7D0]':
                    asset.status === 'In Stock',
                  'bg-tertiary-fixed text-status-checked-out border-tertiary-fixed-dim/50':
                    asset.status === 'Checked Out',
                  'bg-surface-variant text-text-secondary border-border-default':
                    asset.status === 'Reserved',
                }"
              >
                <span class="sr-only">Status:</span>
                {{ asset.status }}
              </span>
            </div>
            <div class="flex items-center gap-4 text-metadata text-text-secondary font-metadata">
              <div class="flex items-center gap-1">
                <span class="material-symbols-outlined text-[14px]">tag</span>
                <span
                  class="font-mono-data text-mono-data text-text-primary bg-surface-subtle px-1.5 py-0.5 rounded border border-border-default"
                >
                  {{ asset.tagId }}
                </span>
              </div>
              <div class="flex items-center gap-1">
                <span class="material-symbols-outlined text-[14px]">location_on</span>
                <span>{{ asset.location }}</span>
              </div>
            </div>
          </div>
          <div class="flex gap-inline">
            <button
              class="px-4 py-2 bg-surface text-text-primary border border-border-default rounded-lg font-h3 text-h3 shadow-[0_2px_4px_rgba(0,0,0,0.02)] hover:-translate-y-[1px] transition-transform flex items-center gap-2"
            >
              <span class="material-symbols-outlined text-[18px]">edit</span>
              Edit
            </button>
            <button
              class="px-4 py-2 bg-[#1C1917] text-white rounded-lg font-h3 text-h3 shadow-sm hover:-translate-y-[1px] transition-transform flex items-center gap-2"
            >
              <span class="material-symbols-outlined text-[18px]">handyman</span>
              Log Maintenance
            </button>
          </div>
        </div>
      </div>
      <!-- Top Section: Image + KPIs -->
      <div class="grid grid-cols-1 lg:grid-cols-12 gap-inline mb-section-gap">
        <!-- Image Gallery -->
        <div
          class="lg:col-span-5 bg-surface border border-border-default rounded-xl overflow-hidden shadow-[0_4px_12px_rgba(0,0,0,0.03)] flex flex-col"
        >
          <div
            class="relative h-64 w-full bg-stone-100 border-b border-border-default flex items-center justify-center"
          >
            <span class="material-symbols-outlined text-[80px] text-stone-300">devices</span>
          </div>
          <div class="flex gap-2 p-4 bg-surface-subtle overflow-x-auto">
            <div
              class="h-16 w-16 flex-shrink-0 rounded-lg border border-border-default opacity-70 hover:opacity-100 transition-opacity flex items-center justify-center bg-stone-100 cursor-pointer"
            >
              <span class="material-symbols-outlined text-stone-400">add_photo_alternate</span>
            </div>
          </div>
        </div>
        <!-- KPIs -->
        <div class="lg:col-span-7 grid grid-cols-2 md:grid-cols-3 gap-inline">
          <div
            class="bg-surface border border-border-default rounded-xl p-card-padding shadow-[0_4px_12px_rgba(0,0,0,0.03)] hover:-translate-y-[2px] transition-transform cursor-pointer group flex flex-col justify-between"
          >
            <div class="flex justify-between items-start mb-4">
              <div
                class="w-10 h-10 rounded-lg bg-metric-amber flex items-center justify-center text-amber-700"
              >
                <span class="material-symbols-outlined">payments</span>
              </div>
            </div>
            <div>
              <div class="text-metadata text-text-secondary font-metadata mb-1">Book Value</div>
              <div class="font-kpi-number text-kpi-number text-text-primary tracking-tight">
                ₹{{ asset.value.toLocaleString() }}
              </div>
            </div>
          </div>
          <div
            class="bg-surface border border-border-default rounded-xl p-card-padding shadow-[0_4px_12px_rgba(0,0,0,0.03)] hover:-translate-y-[2px] transition-transform cursor-pointer group flex flex-col justify-between"
          >
            <div class="flex justify-between items-start mb-4">
              <div
                class="w-10 h-10 rounded-lg bg-metric-lavender flex items-center justify-center text-purple-700"
              >
                <span class="material-symbols-outlined">health_and_safety</span>
              </div>
            </div>
            <div>
              <div class="text-metadata text-text-secondary font-metadata mb-1">
                Warranty Status
              </div>
              <div class="font-kpi-number text-kpi-number text-text-primary tracking-tight text-lg">
                {{ asset.warrantyStatus }}
              </div>
            </div>
          </div>
          <div
            class="bg-surface border border-border-default rounded-xl p-card-padding shadow-[0_4px_12px_rgba(0,0,0,0.03)] hover:-translate-y-[2px] transition-transform cursor-pointer group flex flex-col justify-between col-span-2 md:col-span-1"
          >
            <div class="flex justify-between items-start mb-4">
              <div
                class="w-10 h-10 rounded-lg bg-stone-100 flex items-center justify-center text-stone-700"
              >
                <span class="material-symbols-outlined">assignment_ind</span>
              </div>
            </div>
            <div>
              <div class="text-metadata text-text-secondary font-metadata mb-1">Custodian</div>
              <div class="font-h2 text-h2 text-text-primary truncate">
                {{ asset.custodian || 'Unassigned' }}
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute();

interface Asset {
  id: number;
  tagId: string;
  name: string;
  category: string;
  status: string;
  custodian: string;
  location: string;
  purchaseDate: string;
  warrantyStatus: string;
  value: number;
}

const asset = ref<Asset | null>(null);
const loading = ref(true);

const fetchAsset = async () => {
  try {
    const res = await fetch(`http://localhost:8080/api/assets/${route.params['id']}`);
    if (res.ok) {
      asset.value = await res.json();
    }
  } catch (error) {
    console.error('Failed to fetch asset:', error);
  } finally {
    loading.value = false;
  }
};

onMounted(() => {
  fetchAsset();
});
</script>
