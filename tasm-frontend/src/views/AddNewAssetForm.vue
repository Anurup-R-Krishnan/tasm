<template>
  <main class="space-y-section-gap pb-24 max-w-[1000px]">
    <div class="mb-8">
      <div class="flex items-center gap-2 text-text-disabled font-body mb-2">
        <a
          class="hover:text-text-primary"
          @click.prevent="router.push('/inventory')"
          href="/inventory"
        >
          Asset Inventory
        </a>
        <span class="material-symbols-outlined text-[16px]"> chevron_right </span>
        <span class="text-text-primary"> Add New Asset </span>
      </div>
      <h2 class="font-h1 text-h1 text-text-primary">Add New Asset</h2>
      <p class="font-body text-text-secondary mt-1">
        Enter details to register a new asset into the system.
      </p>
    </div>
    <div
      class="bg-surface rounded-xl border border-border-default shadow-[0_4px_4px_rgba(0,0,0,0.02)] p-card-padding"
    >
      <form @submit.prevent="submitForm">
        <!-- 1. Identity -->
        <div class="mb-section-gap">
          <h3
            class="font-h3 text-h3 text-text-primary border-b border-border-default pb-2 mb-6 flex items-center gap-2"
          >
            <span class="material-symbols-outlined text-primary text-[20px]"> badge </span>
            Identity
          </h3>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-x-6 gap-y-5">
            <div class="col-span-1 md:col-span-2">
              <label class="block font-table-header text-table-header text-text-secondary mb-1.5">
                Asset Name (Description)
                <span class="text-error">*</span>
              </label>
              <input
                v-model="form.name"
                required
                class="w-full border border-border-default rounded-[6px] px-3 py-2 font-body focus:outline-none focus:border-primary focus:ring-1 focus:ring-primary"
                placeholder="e.g. Dell Latitude 5420 Laptop"
                type="text"
              />
            </div>
            <div>
              <label
                class="block font-table-header text-table-header text-text-secondary mb-1.5 flex justify-between"
              >
                <span>Tag ID</span>
                <span
                  @click="generateTagId"
                  class="font-metadata text-metadata text-primary cursor-pointer hover:underline"
                >
                  Auto-generate
                </span>
              </label>
              <input
                v-model="form.tagId"
                required
                class="w-full border border-border-default rounded-[6px] px-3 py-2 font-mono-data text-mono-data bg-surface-subtle focus:outline-none focus:border-primary focus:ring-1 focus:ring-primary"
                placeholder="AST-2024-0001"
                type="text"
              />
            </div>
            <div>
              <label class="block font-table-header text-table-header text-text-secondary mb-1.5">
                Category <span class="text-error">*</span>
              </label>
              <div class="relative">
                <select
                  v-model="form.category"
                  required
                  class="w-full border border-border-default rounded-[6px] px-3 py-2 font-body appearance-none focus:outline-none focus:border-primary focus:ring-1 focus:ring-primary bg-white pr-8"
                >
                  <option disabled value="">Select Category</option>
                  <option>IT Equipment</option>
                  <option>Furniture</option>
                  <option>Vehicles</option>
                  <option>Machinery</option>
                </select>
                <span
                  class="material-symbols-outlined absolute right-2 top-1/2 -translate-y-1/2 pointer-events-none text-text-disabled"
                >
                  expand_more
                </span>
              </div>
            </div>
            <div>
              <label class="block font-table-header text-table-header text-text-secondary mb-1.5">
                Condition <span class="text-error">*</span>
              </label>
              <div class="relative">
                <select
                  v-model="form.condition"
                  class="w-full border border-border-default rounded-[6px] px-3 py-2 font-body appearance-none focus:outline-none focus:border-primary focus:ring-1 focus:ring-primary bg-white pr-8"
                >
                  <option>New</option>
                  <option>Good</option>
                  <option>Fair</option>
                  <option>Poor</option>
                </select>
                <span
                  class="material-symbols-outlined absolute right-2 top-1/2 -translate-y-1/2 pointer-events-none text-text-disabled"
                >
                  expand_more
                </span>
              </div>
            </div>
          </div>
        </div>
        <!-- 2. Purchase & Finance -->
        <div class="mb-section-gap">
          <h3
            class="font-h3 text-h3 text-text-primary border-b border-border-default pb-2 mb-6 flex items-center gap-2"
          >
            <span class="material-symbols-outlined text-primary text-[20px]"> payments </span>
            Purchase &amp; Finance
          </h3>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-x-6 gap-y-5">
            <div>
              <label class="block font-table-header text-table-header text-text-secondary mb-1.5">
                Purchase Date
              </label>
              <div class="relative">
                <input
                  v-model="form.purchaseDate"
                  type="date"
                  class="w-full border border-border-default rounded-[6px] px-3 py-2 font-body focus:outline-none focus:border-primary focus:ring-1 focus:ring-primary"
                />
              </div>
            </div>
            <div>
              <label class="block font-table-header text-table-header text-text-secondary mb-1.5">
                Purchase Cost
              </label>
              <div class="relative">
                <span class="absolute left-3 top-1/2 -translate-y-1/2 text-stone-500 font-body"
                  >₹</span
                >
                <input
                  v-model.number="form.value"
                  type="number"
                  step="0.01"
                  class="w-full border border-border-default rounded-[6px] pl-7 pr-3 py-2 font-mono-data text-mono-data focus:outline-none focus:border-primary focus:ring-1 focus:ring-primary"
                  placeholder="0.00"
                />
              </div>
            </div>
            <div>
              <label class="block font-table-header text-table-header text-text-secondary mb-1.5">
                Location
              </label>
              <div class="relative">
                <input
                  v-model="form.location"
                  required
                  class="w-full border border-border-default rounded-[6px] px-3 py-2 font-body focus:outline-none focus:border-primary focus:ring-1 focus:ring-primary"
                  placeholder="e.g. Stockroom A"
                  type="text"
                />
              </div>
            </div>
          </div>
        </div>

        <!-- Actions -->
        <div class="flex justify-end gap-4 mt-8 border-t border-border-default pt-6">
          <button
            type="button"
            class="px-6 py-2 border border-border-default rounded-lg text-text-primary hover:bg-surface-subtle font-medium transition-colors"
            @click="router.back()"
          >
            Cancel
          </button>
          <button
            type="submit"
            :disabled="isSubmitting"
            class="px-6 py-2 bg-primary text-on-primary rounded-lg font-medium hover:bg-primary-hover transition-colors flex items-center gap-2"
          >
            <span v-if="isSubmitting" class="material-symbols-outlined animate-spin text-[18px]"
              >refresh</span
            >
            Save Asset
          </button>
        </div>
      </form>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { createAsset, getAssetById, updateAsset } from '../api/assets';

const router = useRouter();
const route = useRoute();

const form = ref({
  name: '',
  tagId: '',
  category: '',
  condition: 'New',
  purchaseDate: '',
  value: 0,
  location: '',
  status: 'In Stock',
  custodian: 'Unassigned',
  currentStock: 1,
  reorderLevel: 0,
  warrantyStatus: 'Active',
  warrantyExpiry: new Date().toISOString(),
});

const isSubmitting = ref(false);
const isEditMode = ref(false);

const generateTagId = () => {
  form.value.tagId = 'TAG-' + Math.floor(1000 + Math.random() * 9000);
};

const submitForm = async () => {
  isSubmitting.value = true;
  try {
    const payload = {
      ...form.value,
      purchaseDate: form.value.purchaseDate
        ? new Date(form.value.purchaseDate).toISOString()
        : new Date().toISOString(),
    };
    if (isEditMode.value) {
      const id = route.query['edit'];
      if (typeof id === 'string') {
        await updateAsset(id, payload as any);
        alert('Asset updated successfully!');
      }
    } else {
      await createAsset(payload as any);
      alert('Asset created successfully!');
    }
    router.push('/inventory');
  } catch (error: any) {
    console.error(error);
    alert(error.message || 'Failed to save asset');
  } finally {
    isSubmitting.value = false;
  }
};

const loadAsset = async (id: string) => {
  try {
    const asset = await getAssetById(id);
    form.value = {
      ...form.value,
      ...asset,
      purchaseDate: asset.purchaseDate ? (asset.purchaseDate.split('T')[0] ?? '') : '',
      warrantyExpiry: asset.warrantyExpiry ? asset.warrantyExpiry : new Date().toISOString(),
      condition: form.value.condition,
    };
    isEditMode.value = true;
  } catch (error) {
    console.error('Failed to load asset', error);
  }
};

onMounted(() => {
  const editId = route.query['edit'];
  if (typeof editId === 'string' && editId) {
    loadAsset(editId);
  }
});
</script>
