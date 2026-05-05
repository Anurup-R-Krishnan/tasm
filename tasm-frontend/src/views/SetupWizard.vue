<template>
  <div
    class="setup-wizard-container min-h-screen bg-canvas flex items-center justify-center p-6 relative overflow-hidden"
  >
    <!-- Abstract background decorations -->
    <div class="absolute top-0 left-0 w-full h-full pointer-events-none overflow-hidden">
      <div class="absolute -top-24 -left-24 w-96 h-96 bg-primary/5 rounded-full blur-[120px]"></div>
      <div
        class="absolute -bottom-24 -right-24 w-96 h-96 bg-accent/5 rounded-full blur-[120px]"
      ></div>
    </div>

    <div
      class="w-full max-w-4xl bg-surface border border-border-default rounded-3xl shadow-xl p-10 relative z-10"
    >
      <!-- Header -->
      <div class="text-center mb-10">
        <div
          class="inline-flex items-center justify-center w-16 h-16 rounded-2xl bg-primary-container text-primary mb-4 shadow-sm"
        >
          <span class="material-symbols-outlined text-[32px]">inventory_2</span>
        </div>
        <h1 class="font-h1 text-h1 text-text-primary mb-2">Welcome to TASM</h1>
        <p class="font-body text-body text-text-secondary">
          Let's set up your asset management workspace in a few quick steps.
        </p>
      </div>

      <!-- Custom Stepper -->
      <div class="flex items-center justify-between mb-12 px-4 relative">
        <div class="absolute top-5 left-8 right-8 h-px bg-border-default z-0"></div>
        <div v-for="step in 5" :key="step" class="relative z-10 flex flex-col items-center gap-2">
          <div
            class="w-10 h-10 rounded-full flex items-center justify-center transition-all duration-300 border-2"
            :class="[
              activeStep === step
                ? 'bg-primary border-primary text-white scale-110 shadow-lg'
                : activeStep > step
                  ? 'bg-primary/20 border-primary text-primary'
                  : 'bg-surface border-border-default text-text-secondary',
            ]"
          >
            <span v-if="activeStep > step" class="material-symbols-outlined text-[20px]"
              >check</span
            >
            <span v-else class="font-bold">{{ step }}</span>
          </div>
          <span
            class="font-metadata text-[10px] uppercase tracking-widest font-bold transition-colors"
            :class="activeStep === step ? 'text-primary' : 'text-text-secondary'"
          >
            {{ stepNames[step - 1] }}
          </span>
        </div>
      </div>

      <!-- Step Content -->
      <div class="min-h-[400px]">
        <!-- Step 1: Organization -->
        <div v-if="activeStep === 1" class="animate-in fade-in slide-in-from-bottom-4 duration-500">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-8 py-4">
            <div class="space-y-2">
              <label
                class="block font-metadata text-metadata text-text-secondary uppercase tracking-wider"
                >Company Name</label
              >
              <input
                v-model="setupData.companyName"
                type="text"
                class="w-full h-12 px-4 bg-canvas border border-border-default rounded-xl font-body text-text-primary focus:border-primary focus:ring-1 focus:ring-primary outline-none transition-all"
                placeholder="Acme Corp"
              />
            </div>
            <div class="space-y-2">
              <label
                class="block font-metadata text-metadata text-text-secondary uppercase tracking-wider"
                >Base Currency</label
              >
              <select
                v-model="setupData.currency"
                class="w-full h-12 px-4 bg-canvas border border-border-default rounded-xl font-body text-text-primary focus:border-primary outline-none transition-all appearance-none cursor-pointer"
                style="
                  background-image: url('data:image/svg+xml;charset=US-ASCII,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20width%3D%22292.4%22%20height%3D%22292.4%22%3E%3Cpath%20fill%3D%22%23712c00%22%20d%3D%22M287%2069.4a17.6%2017.6%200%200%200-13-5.4H18.4c-5%200-9.3%201.8-12.9%205.4A17.6%2017.6%200%200%200%200%2082.2c0%205%201.8%209.3%205.4%2012.9l128%20127.9c3.6%203.6%207.8%205.4%2012.8%205.4s9.2-1.8%2012.8-5.4L287%2095c3.5-3.5%205.4-7.8%205.4-12.8%200-5-1.9-9.2-5.5-12.8z%22%2F%3E%3C%2Fsvg%3E');
                  background-repeat: no-repeat;
                  background-position: right 1rem top 50%;
                  background-size: 0.65rem auto;
                "
              >
                <option v-for="curr in currencies" :key="curr.code" :value="curr.code">
                  {{ curr.name }}
                </option>
              </select>
            </div>
          </div>
        </div>

        <!-- Step 2: Locations -->
        <div v-if="activeStep === 2" class="animate-in fade-in slide-in-from-bottom-4 duration-500">
          <div class="flex items-center justify-between mb-6">
            <h3 class="font-h3 text-text-primary">Physical Sites & Buildings</h3>
            <button
              @click="addLocation"
              class="text-primary font-bold text-sm flex items-center gap-1 hover:underline"
            >
              <span class="material-symbols-outlined text-[18px]">add</span> Add Another
            </button>
          </div>
          <div class="space-y-4 max-h-[350px] overflow-y-auto pr-2">
            <div
              v-for="(loc, index) in setupData.locations"
              :key="index"
              class="p-4 bg-canvas border border-border-default rounded-2xl flex items-center gap-4 group"
            >
              <div
                class="w-8 h-8 rounded-lg bg-surface border border-border-default flex items-center justify-center text-text-secondary font-mono-data text-xs"
              >
                {{ index + 1 }}
              </div>
              <div class="grid grid-cols-2 flex-1 gap-4">
                <input
                  v-model="loc.name"
                  placeholder="Site Name (e.g. HQ Office)"
                  class="bg-transparent border-b border-border-default py-2 focus:border-primary outline-none transition-all"
                />
                <select
                  v-model="loc.type"
                  class="bg-transparent border-b border-border-default py-2 focus:border-primary outline-none transition-all cursor-pointer"
                >
                  <option v-for="t in locationTypes" :key="t" :value="t">{{ t }}</option>
                </select>
              </div>
              <button
                @click="removeLocation(index)"
                class="text-text-tertiary hover:text-error opacity-0 group-hover:opacity-100 transition-all"
              >
                <span class="material-symbols-outlined text-[20px]">delete</span>
              </button>
            </div>
          </div>
        </div>

        <!-- Step 3: Categories -->
        <div v-if="activeStep === 3" class="animate-in fade-in slide-in-from-bottom-4 duration-500">
          <div class="grid grid-cols-2 sm:grid-cols-4 gap-4 py-4">
            <div
              v-for="cat in availableCategories"
              :key="cat.name"
              class="relative flex flex-col items-center gap-3 p-5 border rounded-2xl cursor-pointer transition-all hover:-translate-y-1"
              :class="
                setupData.categories.includes(cat.name)
                  ? 'border-primary bg-primary/5 ring-1 ring-primary'
                  : 'border-border-default bg-surface'
              "
              @click="toggleCategory(cat.name)"
            >
              <div
                class="w-12 h-12 rounded-xl flex items-center justify-center transition-colors"
                :class="
                  setupData.categories.includes(cat.name)
                    ? 'bg-primary text-white shadow-md'
                    : 'bg-canvas text-text-secondary'
                "
              >
                <span class="material-symbols-outlined text-[24px]">{{ cat.icon }}</span>
              </div>
              <span
                class="font-metadata text-[11px] font-bold text-center uppercase tracking-tight"
                >{{ cat.name }}</span
              >
              <div v-if="setupData.categories.includes(cat.name)" class="absolute top-2 right-2">
                <span class="material-symbols-outlined text-primary text-[16px]">check_circle</span>
              </div>
            </div>
          </div>
        </div>

        <!-- Step 4: Asset Fields -->
        <div v-if="activeStep === 4" class="animate-in fade-in slide-in-from-bottom-4 duration-500">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4 py-4">
            <div
              v-for="field in Object.keys(setupData.fields)"
              :key="field"
              class="flex items-center justify-between p-4 border border-border-default rounded-2xl bg-canvas hover:border-primary/30 transition-all cursor-pointer"
              @click="setupData.fields[field] = !setupData.fields[field]"
            >
              <div class="flex flex-col">
                <span class="font-h3 text-text-primary capitalize">{{
                  field.replace(/([A-Z])/g, ' $1')
                }}</span>
                <span class="font-metadata text-[11px] text-text-secondary"
                  >Capture during asset creation</span
                >
              </div>
              <div
                class="w-12 h-6 rounded-full relative transition-colors duration-200"
                :class="setupData.fields[field] ? 'bg-primary' : 'bg-border-default'"
              >
                <div
                  class="absolute top-1 left-1 w-4 h-4 bg-white rounded-full transition-transform duration-200"
                  :class="{ 'translate-x-6': setupData.fields[field] }"
                ></div>
              </div>
            </div>
          </div>
        </div>

        <!-- Step 5: Advanced Features -->
        <div v-if="activeStep === 5" class="animate-in fade-in slide-in-from-bottom-4 duration-500">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6 py-4">
            <div
              v-for="feature in Object.keys(setupData.features)"
              :key="feature"
              class="flex items-start gap-4 p-5 border border-border-default rounded-2xl transition-all"
              :class="setupData.features[feature] ? 'bg-primary/5 border-primary/30' : 'bg-surface'"
            >
              <div
                class="w-12 h-12 rounded-xl flex items-center justify-center transition-colors"
                :class="
                  setupData.features[feature]
                    ? 'bg-primary text-white shadow-md'
                    : 'bg-canvas text-text-secondary'
                "
              >
                <span class="material-symbols-outlined text-[24px]">{{
                  featureIcons[feature]
                }}</span>
              </div>
              <div class="flex-1">
                <div class="flex items-center justify-between mb-1">
                  <span class="font-h3 text-text-primary capitalize">{{
                    feature.replace(/([A-Z])/g, ' $1')
                  }}</span>
                  <div
                    class="w-10 h-5 rounded-full relative transition-colors duration-200 cursor-pointer"
                    :class="setupData.features[feature] ? 'bg-primary' : 'bg-border-default'"
                    @click="setupData.features[feature] = !setupData.features[feature]"
                  >
                    <div
                      class="absolute top-0.5 left-0.5 w-4 h-4 bg-white rounded-full transition-transform duration-200"
                      :class="{ 'translate-x-5': setupData.features[feature] }"
                    ></div>
                  </div>
                </div>
                <p class="font-metadata text-[11px] text-text-secondary leading-tight">
                  {{ featureDescriptions[feature] }}
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Action Footer -->
      <div class="flex items-center justify-between mt-12 pt-8 border-t border-border-default">
        <button
          v-if="activeStep > 1"
          @click="activeStep--"
          class="flex items-center gap-2 text-text-secondary hover:text-text-primary font-h3 transition-colors"
        >
          <span class="material-symbols-outlined">arrow_back</span> Back
        </button>
        <div v-else></div>

        <button
          v-if="activeStep < 5"
          @click="activeStep++"
          :disabled="!isCurrentStepValid"
          class="bg-primary text-on-primary px-8 h-12 rounded-xl font-h3 shadow-lg shadow-primary/20 hover:opacity-90 transition-all active:scale-95 disabled:opacity-50 disabled:cursor-not-allowed flex items-center gap-2"
        >
          Next Step <span class="material-symbols-outlined">arrow_forward</span>
        </button>
        <button
          v-else
          @click="finishSetup"
          :disabled="loading"
          class="bg-primary text-on-primary px-10 h-12 rounded-xl font-h3 shadow-lg shadow-primary/20 hover:opacity-90 transition-all active:scale-95 disabled:opacity-50 flex items-center gap-2"
        >
          <span v-if="loading" class="material-symbols-outlined animate-spin"
            >progress_activity</span
          >
          {{ loading ? 'Finalizing...' : 'Complete Setup' }}
          <span v-if="!loading" class="material-symbols-outlined">rocket_launch</span>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';

const activeStep = ref(1);
const loading = ref(false);

const stepNames = ['Organization', 'Locations', 'Categories', 'Fields', 'Features'];

const currencies = [
  { name: 'US Dollar ($)', code: 'USD' },
  { name: 'Euro (€)', code: 'EUR' },
  { name: 'Indian Rupee (₹)', code: 'INR' },
  { name: 'Pound Sterling (£)', code: 'GBP' },
];

const locationTypes = ['Office', 'Warehouse', 'Data Center', 'Clinic', 'School'];

const availableCategories = [
  { name: 'IT Equipment', icon: 'laptop_mac' },
  { name: 'Furniture', icon: 'chair' },
  { name: 'Vehicles', icon: 'directions_car' },
  { name: 'Software', icon: 'terminal' },
  { name: 'Tools', icon: 'handyman' },
  { name: 'AV Gear', icon: 'videocam' },
  { name: 'Real Estate', icon: 'domain' },
  { name: 'Medical', icon: 'medical_services' },
];

const featureIcons: Record<string, string> = {
  depreciation: 'trending_down',
  maintenance: 'build_circle',
  licenses: 'verified_user',
  audits: 'fact_check',
  procurement: 'shopping_cart',
};

const featureDescriptions: Record<string, string> = {
  depreciation: 'Automatically calculate asset value over time using various methods.',
  maintenance: 'Schedule recurring service and track maintenance work orders.',
  licenses: 'Track software keys, renewal dates, and seat assignments.',
  audits: 'Perform physical inventory verification using mobile scanning.',
  procurement: 'Manage requests, approvals, and purchase orders for new assets.',
};

interface SetupData {
  companyName: string;
  currency: string;
  locations: { name: string; type: string }[];
  categories: string[];
  fields: Record<string, boolean>;
  features: Record<string, boolean>;
}

const setupData = ref<SetupData>({
  companyName: '',
  currency: 'USD',
  locations: [{ name: '', type: 'Office' }],
  categories: ['IT Equipment', 'Furniture'],
  fields: {
    serialNumber: true,
    model: true,
    warranty: true,
    color: false,
    department: true,
    purchasePrice: true,
  },
  features: {
    depreciation: true,
    maintenance: true,
    licenses: true,
    audits: true,
    procurement: true,
  },
});

const addLocation = () => setupData.value.locations.push({ name: '', type: 'Office' });
const removeLocation = (index: number) => setupData.value.locations.splice(index, 1);

const toggleCategory = (cat: string) => {
  const index = setupData.value.categories.indexOf(cat);
  if (index > -1) {
    setupData.value.categories.splice(index, 1);
  } else {
    setupData.value.categories.push(cat);
  }
};

const isCurrentStepValid = computed(() => {
  if (activeStep.value === 1) return setupData.value.companyName.trim() !== '';
  if (activeStep.value === 2) return setupData.value.locations.every((l) => l.name.trim() !== '');
  if (activeStep.value === 3) return setupData.value.categories.length > 0;
  return true;
});

const finishSetup = async () => {
  loading.value = true;
  try {
    const token = localStorage.getItem('tasm_auth_token');
    const response = await fetch('/api/setup/complete', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(setupData.value),
    });

    if (response.ok) {
      window.location.href = '/';
    } else {
      const err = await response.json();
      alert('Setup failed: ' + err.error);
    }
  } catch (error) {
    console.error('Setup error:', error);
    alert('An unexpected error occurred.');
  } finally {
    loading.value = false;
  }
};
</script>

<style scoped>
@keyframes slide-in-from-bottom-4 {
  from {
    transform: translateY(1rem);
    opacity: 0;
  }
  to {
    transform: translateY(0);
    opacity: 1;
  }
}
.animate-in {
  animation-duration: 0.5s;
  animation-fill-mode: both;
}

.setup-wizard-container::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-image: radial-gradient(var(--color-border-default) 1px, transparent 1px);
  background-size: 40px 40px;
  opacity: 0.4;
  pointer-events: none;
}

/* Custom Scrollbar */
::-webkit-scrollbar {
  width: 6px;
}
::-webkit-scrollbar-thumb {
  background-color: var(--color-border-default);
  border-radius: 9999px;
}
::-webkit-scrollbar-thumb:hover {
  background-color: var(--color-text-secondary);
}
</style>
