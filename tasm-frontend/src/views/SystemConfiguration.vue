<template>
  <main class="space-y-section-gap pb-24">
    <div class="flex flex-col md:flex-row md:items-end justify-between gap-4">
      <div>
        <h1 class="font-h1 text-h1 text-text-primary">System Configuration</h1>
        <p class="font-body text-text-secondary">
          Manage organization settings, facilities, and asset metadata from one place.
        </p>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-8">
      <div class="lg:col-span-6 space-y-8">
        <section class="premium-card space-y-6">
          <div>
            <h2 class="text-sm font-bold text-text-primary">Organisation</h2>
            <p class="text-[10px] text-text-secondary uppercase tracking-widest mt-1">
              Core system identity
            </p>
          </div>
          <div class="space-y-4">
            <div>
              <label
                class="block text-[10px] font-bold text-text-secondary uppercase tracking-widest"
                >Company Name</label
              >
              <input
                v-model="organisation.companyName"
                type="text"
                class="w-full h-11 px-4 bg-canvas border border-border-default rounded-xl focus:border-primary focus:ring-1 focus:ring-primary outline-none"
              />
            </div>
            <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
              <div>
                <label
                  class="block text-[10px] font-bold text-text-secondary uppercase tracking-widest"
                  >Currency</label
                >
                <input
                  v-model="organisation.currency"
                  type="text"
                  class="w-full h-11 px-4 bg-canvas border border-border-default rounded-xl focus:border-primary focus:ring-1 focus:ring-primary outline-none"
                />
              </div>
              <div>
                <label
                  class="block text-[10px] font-bold text-text-secondary uppercase tracking-widest"
                  >Timezone</label
                >
                <input
                  v-model="organisation.timezone"
                  type="text"
                  class="w-full h-11 px-4 bg-canvas border border-border-default rounded-xl focus:border-primary focus:ring-1 focus:ring-primary outline-none"
                />
              </div>
            </div>
          </div>
          <div class="flex justify-end">
            <button class="btn-primary" @click="saveOrganisation" :disabled="saving.organisation">
              {{ saving.organisation ? 'Saving...' : 'Save Organisation' }}
            </button>
          </div>
        </section>

        <section class="premium-card space-y-6">
          <div>
            <h2 class="text-sm font-bold text-text-primary">Buildings</h2>
            <p class="text-[10px] text-text-secondary uppercase tracking-widest mt-1">
              Campuses & blocks
            </p>
          </div>
          <div class="space-y-4">
            <div class="flex gap-3">
              <input
                v-model="newBuilding"
                type="text"
                placeholder="Add building name"
                class="flex-1 h-11 px-4 bg-canvas border border-border-default rounded-xl focus:border-primary focus:ring-1 focus:ring-primary outline-none"
              />
              <button class="btn-primary" @click="addBuilding" :disabled="!newBuilding.trim()">
                Add
              </button>
            </div>
            <div class="flex flex-wrap gap-2">
              <span
                v-for="building in buildings"
                :key="building"
                class="inline-flex items-center gap-2 px-3 py-1.5 bg-surface-subtle border border-border-default rounded-full text-xs font-semibold text-text-secondary"
              >
                {{ building }}
                <button
                  class="text-text-secondary hover:text-status-critical"
                  @click="removeBuilding(building)"
                >
                  <span class="material-symbols-outlined text-[14px]">close</span>
                </button>
              </span>
              <span v-if="buildings.length === 0" class="text-xs text-text-secondary">
                No buildings added yet.
              </span>
            </div>
          </div>
          <div class="flex justify-end">
            <button class="btn-primary" @click="saveBuildings" :disabled="saving.buildings">
              {{ saving.buildings ? 'Saving...' : 'Save Buildings' }}
            </button>
          </div>
        </section>

        <section class="premium-card space-y-6">
          <div>
            <h2 class="text-sm font-bold text-text-primary">Departments</h2>
            <p class="text-[10px] text-text-secondary uppercase tracking-widest mt-1">
              Business units & cost owners
            </p>
          </div>
          <div class="space-y-4">
            <div class="flex gap-3">
              <input
                v-model="newDepartment"
                type="text"
                placeholder="Add department name"
                class="flex-1 h-11 px-4 bg-canvas border border-border-default rounded-xl focus:border-primary focus:ring-1 focus:ring-primary outline-none"
              />
              <button class="btn-primary" @click="addDepartment" :disabled="!newDepartment.trim()">
                Add
              </button>
            </div>
            <div class="flex flex-wrap gap-2">
              <span
                v-for="dept in departments"
                :key="dept"
                class="inline-flex items-center gap-2 px-3 py-1.5 bg-surface-subtle border border-border-default rounded-full text-xs font-semibold text-text-secondary"
              >
                {{ dept }}
                <button
                  class="text-text-secondary hover:text-status-critical"
                  @click="removeDepartment(dept)"
                >
                  <span class="material-symbols-outlined text-[14px]">close</span>
                </button>
              </span>
              <span v-if="departments.length === 0" class="text-xs text-text-secondary">
                No departments added yet.
              </span>
            </div>
          </div>
          <div class="flex justify-end">
            <button class="btn-primary" @click="saveDepartments" :disabled="saving.departments">
              {{ saving.departments ? 'Saving...' : 'Save Departments' }}
            </button>
          </div>
        </section>
      </div>

      <div class="lg:col-span-6 space-y-8">
        <section class="premium-card space-y-6">
          <div>
            <h2 class="text-sm font-bold text-text-primary">Locations</h2>
            <p class="text-[10px] text-text-secondary uppercase tracking-widest mt-1">
              Rooms and operational areas
            </p>
          </div>
          <div class="space-y-3">
            <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
              <input
                v-model="newLocation.name"
                type="text"
                placeholder="Location name"
                class="h-11 px-4 bg-canvas border border-border-default rounded-xl focus:border-primary focus:ring-1 focus:ring-primary outline-none"
              />
              <input
                v-model="newLocation.type"
                type="text"
                placeholder="Type (Office, Warehouse)"
                class="h-11 px-4 bg-canvas border border-border-default rounded-xl focus:border-primary focus:ring-1 focus:ring-primary outline-none"
              />
            </div>
            <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
              <input
                v-model="newLocation.address"
                type="text"
                placeholder="Address / Description"
                class="h-11 px-4 bg-canvas border border-border-default rounded-xl focus:border-primary focus:ring-1 focus:ring-primary outline-none"
              />
              <input
                v-model.number="newLocation.capacity"
                type="number"
                min="0"
                placeholder="Capacity"
                class="h-11 px-4 bg-canvas border border-border-default rounded-xl focus:border-primary focus:ring-1 focus:ring-primary outline-none"
              />
            </div>
            <div class="flex justify-end">
              <button
                class="btn-primary"
                @click="addLocation"
                :disabled="saving.locations || !newLocation.name.trim()"
              >
                {{ saving.locations ? 'Saving...' : 'Add Location' }}
              </button>
            </div>
            <div class="space-y-2">
              <div
                v-for="loc in locations"
                :key="loc.id"
                class="flex items-center justify-between border border-border-default rounded-xl px-4 py-3"
              >
                <div>
                  <p class="text-sm font-bold text-text-primary">{{ loc.name }}</p>
                  <p class="text-[10px] text-text-secondary uppercase tracking-widest">
                    {{ loc.type }} • {{ loc.address || 'No description' }}
                  </p>
                </div>
                <button
                  class="text-text-secondary hover:text-status-critical"
                  @click="removeLocation(loc.id)"
                >
                  <span class="material-symbols-outlined text-[18px]">delete</span>
                </button>
              </div>
              <p v-if="locations.length === 0" class="text-xs text-text-secondary">
                No locations configured yet.
              </p>
            </div>
          </div>
        </section>

        <section class="premium-card space-y-6">
          <div>
            <h2 class="text-sm font-bold text-text-primary">Categories</h2>
            <p class="text-[10px] text-text-secondary uppercase tracking-widest mt-1">
              Asset classification
            </p>
          </div>
          <div class="space-y-4">
            <div class="flex gap-3">
              <input
                v-model="newCategory"
                type="text"
                placeholder="Add category"
                class="flex-1 h-11 px-4 bg-canvas border border-border-default rounded-xl focus:border-primary focus:ring-1 focus:ring-primary outline-none"
              />
              <button class="btn-primary" @click="addCategory" :disabled="!newCategory.trim()">
                Add
              </button>
            </div>
            <div class="flex flex-wrap gap-2">
              <span
                v-for="category in categories"
                :key="category"
                class="inline-flex items-center gap-2 px-3 py-1.5 bg-surface-subtle border border-border-default rounded-full text-xs font-semibold text-text-secondary"
              >
                {{ category }}
                <button
                  class="text-text-secondary hover:text-status-critical"
                  @click="removeCategory(category)"
                >
                  <span class="material-symbols-outlined text-[14px]">close</span>
                </button>
              </span>
              <span v-if="categories.length === 0" class="text-xs text-text-secondary">
                No categories added yet.
              </span>
            </div>
          </div>
          <div class="flex justify-end">
            <button class="btn-primary" @click="saveCategories" :disabled="saving.categories">
              {{ saving.categories ? 'Saving...' : 'Save Categories' }}
            </button>
          </div>
        </section>

        <section class="premium-card space-y-6">
          <div>
            <h2 class="text-sm font-bold text-text-primary">Fields</h2>
            <p class="text-[10px] text-text-secondary uppercase tracking-widest mt-1">
              Asset metadata toggles
            </p>
          </div>
          <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <label
              v-for="field in fieldOptions"
              :key="field.key"
              class="flex items-center gap-3 border border-border-default rounded-xl px-4 py-3"
            >
              <input v-model="fields[field.key]" type="checkbox" class="h-4 w-4" />
              <span class="text-xs font-semibold text-text-secondary">{{ field.label }}</span>
            </label>
          </div>
          <div class="flex justify-end">
            <button class="btn-primary" @click="saveFields" :disabled="saving.fields">
              {{ saving.fields ? 'Saving...' : 'Save Fields' }}
            </button>
          </div>
        </section>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { getLocations, createLocation, deleteLocation } from '../api/locations';
import { getSystemConfig, updateSystemConfig } from '../api/config';
import type { Location } from '../types/models';

const organisation = ref({
  companyName: '',
  currency: '',
  timezone: '',
});

const buildings = ref<string[]>([]);
const newBuilding = ref('');
const departments = ref<string[]>([]);
const newDepartment = ref('');
const categories = ref<string[]>([]);
const newCategory = ref('');
const fields = ref<Record<string, boolean>>({});
const fieldOptions = [
  { key: 'serialNumber', label: 'Serial Number' },
  { key: 'model', label: 'Model' },
  { key: 'warranty', label: 'Warranty' },
  { key: 'purchasePrice', label: 'Purchase Price' },
  { key: 'department', label: 'Department' },
];

const locations = ref<Location[]>([]);
const newLocation = ref({
  name: '',
  type: 'Office',
  address: '',
  capacity: 0,
  status: 'Active',
});

const saving = ref({
  organisation: false,
  buildings: false,
  departments: false,
  categories: false,
  fields: false,
  locations: false,
});

const parseList = (value?: string) =>
  (value || '')
    .split(',')
    .map((item) => item.trim())
    .filter(Boolean);

const saveConfig = async (
  payload: Record<string, string>,
  savingKey: keyof typeof saving.value,
) => {
  saving.value[savingKey] = true;
  try {
    await updateSystemConfig(payload);
  } catch (error) {
    console.error('Failed to save configuration:', error);
  } finally {
    saving.value[savingKey] = false;
  }
};

const saveOrganisation = async () => {
  await saveConfig(
    {
      company_name: organisation.value.companyName,
      currency: organisation.value.currency,
      timezone: organisation.value.timezone,
    },
    'organisation',
  );
};

const addBuilding = () => {
  if (!newBuilding.value.trim()) return;
  buildings.value.push(newBuilding.value.trim());
  newBuilding.value = '';
};

const removeBuilding = (building: string) => {
  buildings.value = buildings.value.filter((b) => b !== building);
};

const saveBuildings = async () => {
  await saveConfig({ buildings: buildings.value.join(',') }, 'buildings');
};

const addDepartment = () => {
  if (!newDepartment.value.trim()) return;
  departments.value.push(newDepartment.value.trim());
  newDepartment.value = '';
};

const removeDepartment = (dept: string) => {
  departments.value = departments.value.filter((d) => d !== dept);
};

const saveDepartments = async () => {
  await saveConfig({ departments: departments.value.join(',') }, 'departments');
};

const addCategory = () => {
  if (!newCategory.value.trim()) return;
  categories.value.push(newCategory.value.trim());
  newCategory.value = '';
};

const removeCategory = (category: string) => {
  categories.value = categories.value.filter((c) => c !== category);
};

const saveCategories = async () => {
  await saveConfig({ asset_categories: categories.value.join(',') }, 'categories');
};

const saveFields = async () => {
  const payload: Record<string, string> = {};
  Object.entries(fields.value).forEach(([key, value]) => {
    payload[`field_${key}`] = value ? 'true' : 'false';
  });
  await saveConfig(payload, 'fields');
};

const addLocation = async () => {
  if (!newLocation.value.name.trim()) return;
  saving.value.locations = true;
  try {
    const created = await createLocation(newLocation.value);
    locations.value = [...locations.value, created];
    newLocation.value = {
      name: '',
      type: 'Office',
      address: '',
      capacity: 0,
      status: 'Active',
    };
  } catch (error) {
    console.error('Failed to add location:', error);
  } finally {
    saving.value.locations = false;
  }
};

const removeLocation = async (id: number) => {
  try {
    await deleteLocation(id);
    locations.value = locations.value.filter((loc) => loc.id !== id);
  } catch (error) {
    console.error('Failed to delete location:', error);
  }
};

const loadConfig = async () => {
  try {
    const data = await getSystemConfig();
    organisation.value.companyName = data.company_name || '';
    organisation.value.currency = data.currency || '';
    organisation.value.timezone = data.timezone || '';
    buildings.value = parseList(data.buildings);
    departments.value = parseList(data.departments);
    categories.value = parseList(data.asset_categories);
    fieldOptions.forEach((field) => {
      const raw = data[`field_${field.key}`];
      fields.value[field.key] = raw === 'true';
    });
  } catch (error) {
    console.error('Failed to load system config:', error);
  }
};

const loadLocations = async () => {
  try {
    locations.value = await getLocations();
  } catch (error) {
    console.error('Failed to load locations:', error);
  }
};

onMounted(() => {
  loadConfig();
  loadLocations();
});
</script>
