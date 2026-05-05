<template>
  <main class="space-y-section-gap pb-24 max-w-[1000px]">
    <!-- Page Header -->
    <div class="mb-section-gap text-center">
      <h1 class="font-h1 text-h1 text-text-primary mb-base">Asset Check-Out</h1>
      <p class="font-body text-body text-text-secondary">Assign equipment to a staff member.</p>
    </div>
    <!-- Stepper -->
    <div class="mb-section-gap">
      <div class="flex items-center justify-between w-full relative">
        <div
          class="absolute left-0 top-1/2 -translate-y-1/2 w-full h-0.5 bg-border-default -z-10"
        ></div>
        <div
          class="absolute left-0 top-1/2 -translate-y-1/2 w-1/4 h-0.5 bg-primary-container -z-10"
        ></div>
        <div class="flex flex-col items-center gap-base bg-canvas px-inline relative">
          <div
            class="w-8 h-8 rounded-full bg-primary-container text-on-primary flex items-center justify-center font-h3 text-h3 shadow-sm border-2 border-primary-container"
          >
            1
          </div>
          <span class="font-h3 text-h3 text-primary-container">Select Asset</span>
        </div>
        <div class="flex flex-col items-center gap-base bg-canvas px-inline relative">
          <div
            class="w-8 h-8 rounded-full bg-surface text-text-secondary flex items-center justify-center font-h3 text-h3 border-2 border-border-default"
          >
            2
          </div>
          <span class="font-h3 text-h3 text-text-secondary">Select Recipient</span>
        </div>
        <div class="flex flex-col items-center gap-base bg-canvas px-inline relative">
          <div
            class="w-8 h-8 rounded-full bg-surface text-text-secondary flex items-center justify-center font-h3 text-h3 border-2 border-border-default"
          >
            3
          </div>
          <span class="font-h3 text-h3 text-text-secondary">Review</span>
        </div>
        <div class="flex flex-col items-center gap-base bg-canvas px-inline relative">
          <div
            class="w-8 h-8 rounded-full bg-surface text-text-secondary flex items-center justify-center font-h3 text-h3 border-2 border-border-default"
          >
            4
          </div>
          <span class="font-h3 text-h3 text-text-secondary">Confirm</span>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-section-gap">
      <!-- Left Column: Asset Selection -->
      <div class="lg:col-span-7 flex flex-col gap-stack">
        <!-- Search and Filter Bar -->
        <div class="flex gap-inline">
          <div class="relative flex-1">
            <span
              class="material-symbols-outlined absolute left-inline top-1/2 -translate-y-1/2 text-text-secondary"
              >search</span
            >
            <input
              v-model="searchQuery"
              class="w-full pl-10 pr-inline py-2 bg-surface border border-border-default rounded-DEFAULT font-body text-body text-text-primary focus:outline-none focus:ring-2 focus:ring-primary-container shadow-sm"
              placeholder="Search by asset tag, model, or type..."
              type="text"
            />
          </div>
        </div>

        <!-- Asset Grid -->
        <div v-if="loading" class="flex justify-center p-8">
          <span class="material-symbols-outlined animate-spin">refresh</span>
        </div>
        <div v-else class="grid grid-cols-1 md:grid-cols-2 gap-stack">
          <div
            v-for="asset in filteredAssets"
            :key="asset.id"
            @click="selectAsset(asset)"
            class="bg-surface rounded-[16px] shadow-sm p-stack relative cursor-pointer hover:-translate-y-0.5 transition-transform"
            :class="
              selectedAsset?.id === asset.id
                ? 'border-2 border-primary-container'
                : 'border border-border-default'
            "
          >
            <div class="absolute top-stack right-stack">
              <span
                v-if="selectedAsset?.id === asset.id"
                class="material-symbols-outlined text-primary-container"
                style="font-variation-settings: 'FILL' 1"
                >check_circle</span
              >
              <span
                v-else
                class="material-symbols-outlined text-text-secondary opacity-0 group-hover:opacity-100 transition-opacity"
                >radio_button_unchecked</span
              >
            </div>

            <div
              class="w-12 h-12 bg-surface-subtle rounded-lg flex items-center justify-center mb-stack border border-border-default"
            >
              <span class="material-symbols-outlined text-text-primary text-[24px]">devices</span>
            </div>

            <div class="mb-base">
              <span
                class="inline-flex items-center px-2 py-0.5 rounded-full font-metadata text-metadata mb-2"
                :class="{
                  'bg-metric-sage text-status-in-stock': asset.status === 'In Stock',
                  'bg-surface-variant text-text-secondary': asset.status !== 'In Stock',
                }"
              >
                {{ asset.status }}
              </span>
            </div>

            <h3 class="font-h2 text-h2 text-text-primary mb-1">{{ asset.name }}</h3>
            <p class="font-mono-data text-mono-data text-text-secondary mb-stack">
              TAG: {{ asset.tagId }}
            </p>

            <div class="flex gap-2 flex-wrap">
              <span
                class="font-metadata text-metadata text-text-secondary bg-surface-subtle px-2 py-1 rounded border border-border-default"
              >
                {{ asset.category }}
              </span>
              <span
                class="font-metadata text-metadata text-text-secondary bg-surface-subtle px-2 py-1 rounded border border-border-default"
              >
                {{ asset.location || 'Unknown location' }}
              </span>
            </div>
          </div>
        </div>
      </div>

      <!-- Right Column: Contextual Details & Recipient -->
      <div class="lg:col-span-5 flex flex-col gap-stack">
        <!-- Recipient Selection -->
        <div
          class="bg-surface rounded-[16px] border border-border-default shadow-[0_4px_4px_0_rgba(0,0,0,0.02)] p-card-padding"
          :class="{ 'opacity-50 pointer-events-none': !selectedAsset }"
        >
          <h2 class="font-h2 text-h2 text-text-primary mb-stack flex items-center gap-2">
            <span class="material-symbols-outlined text-text-secondary text-[20px]"
              >person_add</span
            >
            Select Recipient
          </h2>
          <!-- Recipient Search -->
          <div class="relative mb-stack">
            <span
              class="material-symbols-outlined absolute left-inline top-1/2 -translate-y-1/2 text-text-secondary"
              >search</span
            >
            <input
              v-model="userSearchQuery"
              class="w-full pl-10 pr-inline py-2 bg-surface border border-primary-container rounded-DEFAULT font-body text-body text-text-primary focus:outline-none focus:ring-2 focus:ring-primary-container focus:border-transparent shadow-sm"
              type="text"
              placeholder="Search employees..."
            />
          </div>
          <div v-if="loadingUsers" class="text-sm text-text-secondary italic p-4 text-center">
            Loading employees...
          </div>
          <div v-else class="space-y-2">
            <div
              v-for="user in filteredUsers"
              :key="user.id"
              class="flex items-center justify-between p-3 border border-border-default rounded-xl bg-surface-subtle/50"
            >
              <div class="flex items-center gap-3">
                <div
                  class="w-9 h-9 rounded-full bg-primary-container/20 border border-primary-container/30 flex items-center justify-center text-xs font-bold text-primary"
                >
                  {{ user.name?.charAt(0) || '?' }}
                </div>
                <div>
                  <p class="text-sm font-medium text-text-primary">{{ user.name }}</p>
                  <p class="text-[11px] text-text-secondary">
                    {{ user.employeeId }} • {{ user.department }}
                  </p>
                </div>
              </div>
              <button
                class="px-3 py-1.5 text-xs font-semibold rounded-lg border transition-colors"
                :class="
                  selectedUser?.id === user.id
                    ? 'bg-primary-container/20 text-primary border-primary-container'
                    : 'text-text-secondary border-border-default hover:text-primary hover:border-primary'
                "
                @click="selectUser(user)"
              >
                {{ selectedUser?.id === user.id ? 'Selected' : 'Assign' }}
              </button>
            </div>
            <div v-if="filteredUsers.length === 0" class="text-sm text-text-secondary italic">
              No users match this search.
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Footer Actions -->
    <div class="mt-section-gap flex justify-end gap-inline border-t border-border-default pt-stack">
      <button
        class="px-6 py-2 bg-surface border border-outline text-text-primary font-h3 text-h3 rounded-DEFAULT hover:bg-surface-subtle transition-colors shadow-sm"
        @click="router.push('/inventory')"
      >
        Back
      </button>
      <button
        :disabled="!selectedAsset || !selectedUser"
        class="px-6 py-2 bg-text-primary text-surface font-h3 text-h3 rounded-DEFAULT hover:bg-inverse-surface transition-colors shadow-sm flex items-center gap-2 disabled:opacity-50"
        @click="goToAssignment"
      >
        Continue to Review
        <span class="material-symbols-outlined text-[18px]">arrow_forward</span>
      </button>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { getAssets } from '../api/assets';
import { getUsers } from '../api/users';
import type { Asset, SystemUser, BaseEntity } from '../types/models';

const router = useRouter();
const route = useRoute();
const assets = ref<Asset[]>([]);
const loading = ref(true);
const searchQuery = ref('');
const selectedAsset = ref<Asset | null>(null);
const users = ref<SystemUser[]>([]);
const loadingUsers = ref(true);
const userSearchQuery = ref('');
const selectedUser = ref<SystemUser | null>(null);

const filteredAssets = computed(() => {
  if (!searchQuery.value) return assets.value;
  const q = searchQuery.value.toLowerCase();
  return assets.value.filter(
    (a) => a.name.toLowerCase().includes(q) || a.tagId.toLowerCase().includes(q),
  );
});

const filteredUsers = computed(() => {
  if (!userSearchQuery.value) return users.value;
  const q = userSearchQuery.value.toLowerCase();
  return users.value.filter(
    (u) =>
      u.name.toLowerCase().includes(q) ||
      u.employeeId.toLowerCase().includes(q) ||
      u.department.toLowerCase().includes(q),
  );
});

const selectAsset = (asset: Asset) => {
  // Only allow selecting available assets
  if (asset.status === 'In Stock') {
    selectedAsset.value = asset;
  } else {
    alert('This asset is not available for check-out.');
  }
};

const selectUser = (user: SystemUser) => {
  selectedUser.value = user;
};

const goToAssignment = () => {
  if (!selectedAsset.value || !selectedUser.value) return;
  router.push({
    name: 'AssetCheckOutStep2',
    query: {
      assetId: String(selectedAsset.value.id),
      userId: String(selectedUser.value.id),
    },
  });
};

const mapAssetIds = (items: BaseEntity[]) => items.map((item) => item.id);

const fetchAssets = async () => {
  try {
    assets.value = (await getAssets()) as any[];
    mapAssetIds(assets.value);
    const preselectedId = route.query['assetId'];
    if (typeof preselectedId === 'string' && preselectedId) {
      const match = assets.value.find((asset) => String(asset.id) === preselectedId);
      if (match) {
        selectedAsset.value = match;
      } else {
        alert('Selected asset not found in inventory list.');
      }
    }
  } catch (error) {
    console.error(error);
  } finally {
    loading.value = false;
  }
};

const fetchUsers = async () => {
  try {
    users.value = await getUsers();
  } catch (error) {
    console.error('Failed to load users', error);
  } finally {
    loadingUsers.value = false;
  }
};

onMounted(() => {
  fetchAssets();
  fetchUsers();
});
</script>
