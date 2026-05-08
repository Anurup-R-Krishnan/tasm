<template>
  <div class="min-h-screen flex flex-col md:flex-row bg-canvas font-body overflow-hidden">
    <!-- Left Section: Progress & Branding -->
    <div
      class="relative hidden md:flex md:w-1/3 lg:w-[35%] flex-col p-12 overflow-hidden bg-primary text-white"
    >
      <!-- Background Image with Overlay -->
      <img
        src="/home/anuruprkris/.gemini/antigravity/brain/cfaca06b-fed7-43bb-969b-9dcad87f6f35/tasm_login_branding_1778221047984.png"
        alt="Branding"
        class="absolute inset-0 w-full h-full object-cover opacity-20 mix-blend-overlay"
      />

      <!-- Abstract decorative elements -->
      <div
        class="absolute top-[-10%] left-[-10%] w-[50%] h-[50%] bg-metric-amber/20 rounded-full blur-[120px] pointer-events-none"
      ></div>

      <!-- Header -->
      <div class="relative z-10 flex items-center gap-4 mb-16">
        <div
          class="w-10 h-10 rounded-xl bg-white text-primary flex items-center justify-center shadow-lg"
        >
          <span class="material-symbols-outlined text-[24px]">inventory_2</span>
        </div>
        <h1 class="text-xl font-bold tracking-tight text-white uppercase">Technopark AMS</h1>
      </div>

      <!-- Stepper Content -->
      <div class="relative z-10 flex-1 flex flex-col justify-center max-w-xs mx-auto md:mx-0">
        <div class="space-y-12">
          <div v-for="(title, idx) in stepTitles" :key="idx" class="flex items-start gap-6 group">
            <div class="relative flex flex-col items-center">
              <div
                class="w-10 h-10 rounded-full flex items-center justify-center font-bold border-2 transition-all duration-500 z-10"
                :class="[
                  step === idx + 1
                    ? 'bg-white border-white text-primary scale-110 shadow-xl'
                    : step > idx + 1
                      ? 'bg-white/20 border-white/20 text-white'
                      : 'bg-transparent border-white/10 text-white/30',
                ]"
              >
                <span v-if="step > idx + 1" class="material-symbols-outlined text-[20px]"
                  >check</span
                >
                <span v-else>{{ idx + 1 }}</span>
              </div>
              <!-- Connector Line -->
              <div
                v-if="idx < stepTitles.length - 1"
                class="absolute top-10 w-[2px] h-12 bg-white/10"
              >
                <div
                  class="w-full bg-white transition-all duration-700 ease-out"
                  :style="{ height: step > idx + 1 ? '100%' : '0%' }"
                ></div>
              </div>
            </div>
            <div class="pt-1">
              <p
                class="text-xs font-black uppercase tracking-[0.2em] transition-colors duration-500"
                :class="step === idx + 1 ? 'text-white' : 'text-white/40'"
              >
                {{ title }}
              </p>
              <p
                class="text-[10px] font-bold text-white/30 mt-1 uppercase tracking-widest overflow-hidden transition-all duration-500"
                :class="step === idx + 1 ? 'max-h-10 opacity-100' : 'max-h-0 opacity-0'"
              >
                {{ stepSubtitles[idx] }}
              </p>
            </div>
          </div>
        </div>
      </div>

      <!-- Footer Info -->
      <div
        class="absolute bottom-12 left-12 z-10 flex items-center gap-3 text-white/30 font-metadata text-[10px] uppercase tracking-[0.3em]"
      >
        <span>Step {{ step }} of 5</span>
        <div class="w-1 h-1 rounded-full bg-white/20"></div>
        <span>Setup Wizard</span>
      </div>
    </div>

    <!-- Right Section: Step Content -->
    <div class="flex-1 flex flex-col bg-surface relative overflow-y-auto">
      <div class="flex-1 flex flex-col items-center justify-center p-8 md:p-16 lg:p-24 w-full">
        <div
          class="w-full max-w-3xl space-y-12 animate-in fade-in slide-in-from-right-8 duration-700"
        >
          <!-- Step 1: Identity -->
          <div v-if="step === 1" class="space-y-10">
            <div class="space-y-4">
              <h2 class="text-4xl font-bold text-text-primary tracking-tight">Master Account</h2>
              <p class="text-text-secondary text-lg font-medium">
                Define the primary administrator for the system. This account will have full control
                over all modules and users.
              </p>
            </div>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
              <div class="space-y-6">
                <div class="space-y-2">
                  <label
                    class="block text-[10px] font-black text-text-secondary uppercase tracking-[0.2em] ml-1"
                    >Full Name</label
                  >
                  <input
                    v-model="form.name"
                    type="text"
                    placeholder="e.g. Alex Rivera"
                    class="w-full h-14 px-5 bg-canvas border border-border-default rounded-2xl text-text-primary focus:outline-none focus:border-primary focus:ring-4 focus:ring-primary/5 transition-all font-medium"
                    :class="{ 'border-status-critical': errors['name'] }"
                  />
                </div>
                <div class="space-y-2">
                  <label
                    class="block text-[10px] font-black text-text-secondary uppercase tracking-[0.2em] ml-1"
                    >Email Address</label
                  >
                  <input
                    v-model="form.email"
                    type="email"
                    placeholder="alex@company.com"
                    class="w-full h-14 px-5 bg-canvas border border-border-default rounded-2xl text-text-primary focus:outline-none focus:border-primary focus:ring-4 focus:ring-primary/5 transition-all font-medium"
                  />
                </div>
              </div>

              <div class="space-y-6">
                <div class="space-y-2">
                  <label
                    class="block text-[10px] font-black text-text-secondary uppercase tracking-[0.2em] ml-1"
                    >Secure Password</label
                  >
                  <div class="relative group">
                    <input
                      v-model="form.password"
                      :type="showPass ? 'text' : 'password'"
                      placeholder="••••••••"
                      class="w-full h-14 px-5 bg-canvas border border-border-default rounded-2xl text-text-primary focus:outline-none focus:border-primary focus:ring-4 focus:ring-primary/5 transition-all font-medium pr-12"
                    />
                    <button
                      @click="showPass = !showPass"
                      class="absolute right-4 top-1/2 -translate-y-1/2 text-text-disabled hover:text-primary transition-colors"
                    >
                      <span class="material-symbols-outlined text-[20px]">{{
                        showPass ? 'visibility_off' : 'visibility'
                      }}</span>
                    </button>
                  </div>
                </div>
                <!-- Strength Checklist -->
                <div class="p-5 rounded-2xl bg-canvas border border-border-default space-y-4">
                  <p class="text-[9px] font-black text-text-secondary uppercase tracking-[0.2em]">
                    Security Compliance
                  </p>
                  <div class="grid grid-cols-2 gap-y-3">
                    <div
                      v-for="rule in passwordRules"
                      :key="rule.label"
                      class="flex items-center gap-2"
                    >
                      <span
                        class="material-symbols-outlined text-[16px]"
                        :class="rule.valid ? 'text-metric-emerald' : 'text-text-disabled'"
                      >
                        {{ rule.valid ? 'check_circle' : 'circle' }}
                      </span>
                      <span
                        class="text-[11px] font-bold"
                        :class="rule.valid ? 'text-text-primary' : 'text-text-secondary'"
                        >{{ rule.label }}</span
                      >
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <p class="text-center text-sm text-text-secondary font-medium pt-4">
              Already have an account?
              <span
                @click="router.push({ name: 'Login' })"
                class="text-primary font-black hover:underline cursor-pointer ml-1"
                >Sign In</span
              >
            </p>
          </div>

          <!-- Step 2: Organization -->
          <div v-if="step === 2" class="space-y-10">
            <div class="space-y-4">
              <h2 class="text-4xl font-bold text-text-primary tracking-tight">
                Organization Profile
              </h2>
              <p class="text-text-secondary text-lg font-medium">
                Establish your workspace parameters and branding.
              </p>
            </div>

            <div class="max-w-2xl space-y-8">
              <div class="space-y-2">
                <label
                  class="block text-[10px] font-black text-text-secondary uppercase tracking-[0.2em] ml-1"
                  >Legal Organization Name</label
                >
                <input
                  v-model="form.companyName"
                  type="text"
                  placeholder="e.g. Acme Corporation"
                  class="w-full px-5 py-10 bg-canvas border border-border-default rounded-[32px] text-text-primary focus:outline-none focus:border-primary focus:ring-4 focus:ring-primary/5 transition-all text-center text-4xl font-black"
                />
              </div>

              <div class="grid grid-cols-2 gap-8">
                <div class="space-y-2">
                  <label
                    class="block text-[10px] font-black text-text-secondary uppercase tracking-[0.2em] ml-1"
                    >Base Currency</label
                  >
                  <select
                    v-model="form.currency"
                    class="w-full h-14 px-5 bg-canvas border border-border-default rounded-2xl text-text-primary focus:outline-none focus:border-primary transition-all cursor-pointer font-bold"
                  >
                    <option v-for="c in currencies" :key="c.code" :value="c.code">
                      {{ c.symbol }} {{ c.name }} ({{ c.code }})
                    </option>
                  </select>
                </div>
                <div class="space-y-2">
                  <label
                    class="block text-[10px] font-black text-text-secondary uppercase tracking-[0.2em] ml-1"
                    >Primary Department</label
                  >
                  <input
                    v-model="form.department"
                    type="text"
                    placeholder="e.g. IT Services"
                    class="w-full h-14 px-5 bg-canvas border border-border-default rounded-2xl text-text-primary focus:outline-none focus:border-primary transition-all font-bold"
                  />
                </div>
              </div>
            </div>
          </div>

          <!-- Step 3: Sites -->
          <div v-if="step === 3" class="space-y-10">
            <div class="space-y-4">
              <h2 class="text-4xl font-bold text-text-primary tracking-tight">Infrastructure</h2>
              <p class="text-text-secondary text-lg font-medium">
                Define the primary physical locations where your assets will be stationed.
              </p>
            </div>

            <div class="space-y-4">
              <div
                v-for="(loc, idx) in form.locations"
                :key="idx"
                class="group flex items-center gap-6 p-6 rounded-3xl bg-canvas border border-border-default hover:border-primary/40 transition-all animate-in fade-in zoom-in-95 duration-300"
              >
                <div
                  class="w-12 h-12 rounded-2xl bg-white flex items-center justify-center text-primary font-black shadow-sm text-lg border border-border-default"
                >
                  {{ idx + 1 }}
                </div>
                <div class="flex-1 grid grid-cols-2 gap-8">
                  <div class="space-y-1">
                    <p
                      class="text-[9px] font-black text-text-disabled uppercase tracking-widest ml-1"
                    >
                      Site Name
                    </p>
                    <input
                      v-model="loc.name"
                      placeholder="Building A"
                      class="w-full bg-transparent border-b-2 border-border-default py-1 focus:border-primary outline-none transition-all font-black text-xl placeholder:text-text-disabled"
                    />
                  </div>
                  <div class="space-y-1">
                    <p
                      class="text-[9px] font-black text-text-disabled uppercase tracking-widest ml-1"
                    >
                      Type
                    </p>
                    <select
                      v-model="loc.type"
                      class="w-full bg-transparent border-b-2 border-border-default py-1 focus:border-primary outline-none transition-all cursor-pointer font-bold"
                    >
                      <option v-for="t in locationTypes" :key="t" :value="t">{{ t }}</option>
                    </select>
                  </div>
                </div>
                <button
                  @click="removeLocation(idx)"
                  v-if="form.locations.length > 1"
                  class="w-10 h-10 rounded-xl text-text-disabled hover:text-status-critical hover:bg-status-critical/10 transition-all opacity-0 group-hover:opacity-100"
                >
                  <span class="material-symbols-outlined">delete</span>
                </button>
              </div>

              <button
                @click="addLocation"
                class="w-full py-6 border-2 border-dashed border-border-default rounded-[32px] text-text-secondary hover:text-primary hover:border-primary/40 hover:bg-primary/5 transition-all flex items-center justify-center gap-3 font-black uppercase tracking-[0.2em] text-xs"
              >
                <span class="material-symbols-outlined">add_circle</span>
                Add Another Site
              </button>
            </div>
          </div>

          <!-- Step 4: Capabilities -->
          <div v-if="step === 4" class="space-y-10">
            <div class="space-y-4">
              <h2 class="text-4xl font-bold text-text-primary tracking-tight">
                System Capabilities
              </h2>
              <p class="text-text-secondary text-lg font-medium">
                Activate specialized modules and select your initial asset categories.
              </p>
            </div>

            <div class="grid grid-cols-1 lg:grid-cols-2 gap-12">
              <div class="space-y-6">
                <p
                  class="text-[10px] font-black text-text-secondary uppercase tracking-[0.2em] ml-1"
                >
                  Specialized Modules
                </p>
                <div class="space-y-3">
                  <div
                    v-for="(desc, feat) in featureDescriptions"
                    :key="feat"
                    @click="form.features[feat] = !form.features[feat]"
                    class="flex items-center gap-5 p-5 rounded-3xl border-2 transition-all cursor-pointer hover:shadow-lg"
                    :class="
                      form.features[feat]
                        ? 'border-primary bg-primary/5'
                        : 'border-border-default bg-surface'
                    "
                  >
                    <div
                      class="w-12 h-12 rounded-2xl flex items-center justify-center transition-all shadow-sm"
                      :class="
                        form.features[feat]
                          ? 'bg-primary text-white shadow-primary/20'
                          : 'bg-canvas text-text-disabled border border-border-default'
                      "
                    >
                      <span class="material-symbols-outlined text-[24px]">{{
                        featureIcons[feat]
                      }}</span>
                    </div>
                    <div class="flex-1">
                      <p
                        class="font-black text-sm uppercase tracking-tight"
                        :class="form.features[feat] ? 'text-primary' : 'text-text-primary'"
                      >
                        {{ feat }}
                      </p>
                      <p class="text-[10px] font-bold text-text-secondary leading-tight mt-0.5">
                        {{ desc }}
                      </p>
                    </div>
                    <div
                      class="w-6 h-6 rounded-full border-2 flex items-center justify-center transition-all"
                      :class="
                        form.features[feat] ? 'border-primary bg-primary' : 'border-border-default'
                      "
                    >
                      <span
                        v-if="form.features[feat]"
                        class="material-symbols-outlined text-white text-[16px] font-black"
                        >check</span
                      >
                    </div>
                  </div>
                </div>
              </div>

              <div class="space-y-6">
                <p
                  class="text-[10px] font-black text-text-secondary uppercase tracking-[0.2em] ml-1"
                >
                  Asset Categories
                </p>
                <div class="grid grid-cols-2 gap-4">
                  <div
                    v-for="cat in availableCategories"
                    :key="cat.name"
                    @click="toggleCategory(cat.name)"
                    class="flex flex-col items-center gap-4 p-6 rounded-[32px] border-2 transition-all cursor-pointer text-center group"
                    :class="
                      form.categories.includes(cat.name)
                        ? 'border-metric-amber bg-metric-amber/5'
                        : 'border-border-default bg-surface hover:border-border-default'
                    "
                  >
                    <div
                      class="w-14 h-14 rounded-2xl flex items-center justify-center transition-all group-hover:scale-110"
                      :class="
                        form.categories.includes(cat.name)
                          ? 'bg-metric-amber text-white shadow-xl shadow-metric-amber/20'
                          : 'bg-canvas text-text-disabled border border-border-default'
                      "
                    >
                      <span class="material-symbols-outlined text-[28px]">{{ cat.icon }}</span>
                    </div>
                    <span
                      class="text-xs font-black uppercase tracking-widest"
                      :class="
                        form.categories.includes(cat.name)
                          ? 'text-text-primary'
                          : 'text-text-disabled'
                      "
                      >{{ cat.name }}</span
                    >
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Step 5: Initializing -->
          <div
            v-if="step === 5"
            class="animate-in zoom-in-95 fade-in duration-700 flex flex-col items-center justify-center text-center space-y-8 py-12"
          >
            <div
              class="w-24 h-24 rounded-[40px] bg-primary text-white flex items-center justify-center shadow-2xl shadow-primary/20 animate-bounce"
            >
              <span class="material-symbols-outlined text-[48px]">rocket_launch</span>
            </div>
            <div class="space-y-3">
              <h1 class="text-5xl font-black text-text-primary tracking-tight">
                Initializing Ecosystem
              </h1>
              <p class="text-text-secondary text-xl font-medium max-w-md">
                We're configuring your high-performance workspace. This will only take a moment.
              </p>
            </div>

            <div
              class="w-full max-w-sm h-3 bg-canvas border border-border-default rounded-full overflow-hidden relative"
            >
              <div
                class="absolute inset-y-0 bg-primary animate-progress-indefinite rounded-full w-1/3"
              ></div>
            </div>

            <p
              class="text-[10px] font-black text-text-disabled uppercase tracking-[0.5em] pt-8 animate-pulse"
            >
              Establishing secure connection
            </p>
          </div>
        </div>

        <!-- Sticky Navigation Footer -->
        <div
          v-if="step < 5"
          class="mt-auto w-full max-w-3xl pt-12 border-t border-border-default flex justify-between items-center"
        >
          <button
            v-if="step > 1"
            @click="step--"
            class="group flex items-center gap-3 text-text-secondary hover:text-text-primary font-black uppercase tracking-widest text-xs transition-all"
          >
            <span class="material-symbols-outlined group-hover:-translate-x-1 transition-transform"
              >arrow_back</span
            >
            Back
          </button>
          <div v-else></div>

          <button
            @click="nextStep"
            :disabled="!isStepValid"
            class="h-16 px-12 bg-primary text-on-primary rounded-[24px] font-black uppercase tracking-[0.2em] shadow-2xl shadow-primary/20 hover:opacity-90 hover:-translate-y-1 transition-all active:translate-y-0 disabled:opacity-50 disabled:translate-y-0 flex items-center gap-3 text-xs"
          >
            {{ step === 4 ? 'Finalize Setup' : 'Continue' }}
            <span class="material-symbols-outlined font-black">{{
              step === 4 ? 'auto_awesome' : 'arrow_forward'
            }}</span>
          </button>
        </div>
      </div>

      <!-- Footer Credits -->
      <div class="p-8 text-center bg-canvas/30 backdrop-blur-sm border-t border-border-default">
        <p class="text-[10px] font-bold text-text-disabled uppercase tracking-[0.4em]">
          &copy; {{ new Date().getFullYear() }} Technopark Asset Management
        </p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { useAuth } from '../composables/useAuth';

const router = useRouter();
const { setToken, checkSetupStatus, isAuthenticated, currentUser } = useAuth();

const step = ref(1);
const loading = ref(false);
const showPass = ref(false);
const errors = ref<Record<string, string>>({});

onMounted(() => {
  if (isAuthenticated.value) {
    step.value = 2;
    if (currentUser.value) {
      form.value.name = currentUser.value.name;
      form.value.email = currentUser.value.email;
      form.value.department = currentUser.value.department;
    }
  }
});

const stepTitles = ['Identity', 'Organization', 'Infrastructure', 'Capabilities', 'Finalize'];
const stepSubtitles = [
  'Admin account creation',
  'Workspace parameters',
  'Physical site hierarchy',
  'Module & category selection',
  'Initializing ecosystem',
];

const form = ref({
  name: '',
  email: '',
  password: '',
  department: '',
  companyName: '',
  currency: 'USD',
  timezone: 'UTC',
  locations: [{ name: 'Main Campus', type: 'Office' }],
  categories: ['IT Equipment', 'Furniture'],
  features: {
    depreciation: true,
    maintenance: true,
    licenses: true,
    audits: true,
    procurement: true,
  } as Record<string, boolean>,
});

const passwordRules = computed(() => [
  { label: '8+ Characters', valid: form.value.password.length >= 8 },
  { label: 'Uppercase', valid: /[A-Z]/.test(form.value.password) },
  { label: 'Lowercase', valid: /[a-z]/.test(form.value.password) },
  { label: 'Number', valid: /[0-9]/.test(form.value.password) },
]);

const isStepValid = computed(() => {
  if (step.value === 1)
    return (
      form.value.name && form.value.email.includes('@') && passwordRules.value.every((r) => r.valid)
    );
  if (step.value === 2) return form.value.companyName.length >= 2;
  if (step.value === 3)
    return form.value.locations.length > 0 && form.value.locations.every((l) => l.name);
  if (step.value === 4) return form.value.categories.length > 0;
  return true;
});

const nextStep = () => {
  if (step.value === 4) {
    step.value = 5;
    handleSubmit();
  } else {
    step.value++;
  }
};

const addLocation = () => form.value.locations.push({ name: '', type: 'Office' });
const removeLocation = (idx: number) => form.value.locations.splice(idx, 1);
const toggleCategory = (cat: string) => {
  const i = form.value.categories.indexOf(cat);
  if (i > -1) form.value.categories.splice(i, 1);
  else form.value.categories.push(cat);
};

const handleSubmit = async () => {
  loading.value = true;
  try {
    let token = localStorage.getItem('tasm_auth_token');

    if (!isAuthenticated.value) {
      const adminRes = await fetch('/api/auth/create-admin', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          name: form.value.name,
          email: form.value.email,
          password: form.value.password,
          department: form.value.department || 'Administration',
          companyName: form.value.companyName,
        }),
      });

      const adminData = await adminRes.json();
      if (!adminRes.ok) throw new Error(adminData.error || 'Failed to create admin');

      // If server returns a token (legacy), use it; otherwise require explicit sign-in
      if (adminData.token) {
        token = adminData.token;
        setToken(token);
      } else {
        alert('Admin account created. Please sign in to continue.');
        window.location.href = '/login';
        return;
      }
    }

    const setupRes = await fetch('/api/setup/complete', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify({
        companyName: form.value.companyName,
        currency: form.value.currency,
        locations: form.value.locations,
        categories: form.value.categories,
        features: form.value.features,
        fields: {
          serialNumber: true,
          model: true,
          warranty: true,
          purchasePrice: true,
          department: true,
        },
      }),
    });

    if (!setupRes.ok) {
      const setupData = await setupRes.json();
      throw new Error(setupData.error || 'Failed to finalize setup');
    }

    await checkSetupStatus();
    setTimeout(() => {
      window.location.href = '/';
    }, 2000);
  } catch (err: any) {
    console.error(err);
    alert(err.message);
    step.value = 1;
  } finally {
    loading.value = false;
  }
};

const currencies = [
  { code: 'USD', name: 'US Dollar', symbol: '$' },
  { code: 'EUR', name: 'Euro', symbol: '€' },
  { code: 'INR', name: 'Indian Rupee', symbol: '₹' },
  { code: 'GBP', name: 'Pound Sterling', symbol: '£' },
];

const locationTypes = ['Office', 'Warehouse', 'Data Center', 'Clinic', 'School'];

const availableCategories = [
  { name: 'IT Equipment', icon: 'laptop_mac' },
  { name: 'Furniture', icon: 'chair' },
  { name: 'Vehicles', icon: 'directions_car' },
  { name: 'Software', icon: 'terminal' },
  { name: 'Tools', icon: 'handyman' },
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
  depreciation: 'Auto-calculate asset value over time.',
  maintenance: 'Schedule recurring service work orders.',
  licenses: 'Track software keys and seat assignments.',
  audits: 'Verify physical inventory via mobile scan.',
  procurement: 'Manage requests and purchase orders.',
};
</script>

<style scoped>
.animate-in {
  animation-duration: 0.7s;
  animation-fill-mode: both;
}

@keyframes slide-in-right {
  from {
    opacity: 0;
    transform: translateX(20px);
  }
  to {
    opacity: 1;
    transform: translateX(0);
  }
}

.slide-in-from-right-8 {
  animation-name: slide-in-right;
}

@keyframes progress-indefinite {
  0% {
    left: -40%;
  }
  100% {
    left: 100%;
  }
}
.animate-progress-indefinite {
  animation: progress-indefinite 2s infinite linear;
}
</style>
