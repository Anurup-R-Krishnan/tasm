<template>
  <div class="min-h-screen bg-canvas flex items-center justify-center p-6 relative overflow-hidden">
    <!-- Background grid -->
    <div
      class="absolute inset-0 pointer-events-none"
      style="
        background-image: radial-gradient(var(--color-border-default) 1px, transparent 1px);
        background-size: 36px 36px;
        opacity: 0.4;
      "
    ></div>
    <div
      class="absolute -top-32 -left-32 w-96 h-96 bg-primary/5 rounded-full blur-[140px] pointer-events-none"
    ></div>
    <div
      class="absolute -bottom-32 -right-32 w-96 h-96 bg-accent/5 rounded-full blur-[140px] pointer-events-none"
    ></div>

    <div class="w-full max-w-lg relative z-10">
      <!-- Step indicator -->
      <div class="flex items-center justify-center gap-3 mb-8">
        <div v-for="s in 3" :key="s" class="flex items-center gap-3">
          <div
            class="w-8 h-8 rounded-full flex items-center justify-center text-sm font-bold transition-all"
            :class="
              step === s
                ? 'bg-primary text-on-primary scale-110 shadow-lg shadow-primary/30'
                : step > s
                  ? 'bg-primary/20 text-primary border border-primary'
                  : 'bg-surface border border-border-default text-text-secondary'
            "
          >
            <span v-if="step > s" class="material-symbols-outlined text-[16px]">check</span>
            <span v-else>{{ s }}</span>
          </div>
          <div
            v-if="s < 3"
            class="w-12 h-px"
            :class="step > s ? 'bg-primary' : 'bg-border-default'"
          ></div>
        </div>
      </div>

      <!-- Card -->
      <div class="bg-surface border border-border-default rounded-3xl shadow-xl p-8">
        <!-- Step 1: Welcome + Admin Account -->
        <div v-if="step === 1" class="space-y-6">
          <div class="text-center">
            <div
              class="w-16 h-16 rounded-2xl bg-primary-container text-primary flex items-center justify-center mx-auto mb-4"
            >
              <span class="material-symbols-outlined text-[32px]">inventory_2</span>
            </div>
            <h1 class="font-h1 text-h1 text-text-primary mb-1">Welcome to TASM</h1>
            <p class="text-text-secondary text-sm">
              Let's create your administrator account to get started.
            </p>
          </div>

          <div class="space-y-4">
            <div class="grid grid-cols-2 gap-4">
              <div>
                <label
                  class="block text-xs font-bold text-text-secondary uppercase tracking-wider mb-1.5"
                  >Full Name *</label
                >
                <input
                  v-model="form.name"
                  type="text"
                  placeholder="Jane Smith"
                  class="w-full h-11 px-4 bg-canvas border border-border-default rounded-xl text-text-primary focus:border-primary focus:ring-1 focus:ring-primary outline-none transition-all"
                  :class="{ 'border-status-critical': errors['name'] }"
                />
                <p v-if="errors['name']" class="text-status-critical text-xs mt-1">
                  {{ errors['name'] }}
                </p>
              </div>
              <div>
                <label
                  class="block text-xs font-bold text-text-secondary uppercase tracking-wider mb-1.5"
                  >Employee ID</label
                >
                <input
                  v-model="form.employeeId"
                  type="text"
                  placeholder="EMP-001"
                  class="w-full h-11 px-4 bg-canvas border border-border-default rounded-xl text-text-primary focus:border-primary focus:ring-1 focus:ring-primary outline-none transition-all"
                />
              </div>
            </div>

            <div>
              <label
                class="block text-xs font-bold text-text-secondary uppercase tracking-wider mb-1.5"
                >Email Address *</label
              >
              <input
                v-model="form.email"
                type="email"
                placeholder="admin@company.com"
                class="w-full h-11 px-4 bg-canvas border border-border-default rounded-xl text-text-primary focus:border-primary focus:ring-1 focus:ring-primary outline-none transition-all"
                :class="{ 'border-status-critical': errors['email'] }"
              />
              <p v-if="errors['email']" class="text-status-critical text-xs mt-1">
                {{ errors['email'] }}
              </p>
            </div>

            <div>
              <label
                class="block text-xs font-bold text-text-secondary uppercase tracking-wider mb-1.5"
                >Department</label
              >
              <input
                v-model="form.department"
                type="text"
                placeholder="IT Services"
                class="w-full h-11 px-4 bg-canvas border border-border-default rounded-xl text-text-primary focus:border-primary focus:ring-1 focus:ring-primary outline-none transition-all"
              />
            </div>

            <div>
              <label
                class="block text-xs font-bold text-text-secondary uppercase tracking-wider mb-1.5"
                >Password *</label
              >
              <div class="relative">
                <input
                  v-model="form.password"
                  :type="showPassword ? 'text' : 'password'"
                  placeholder="Min. 8 characters"
                  class="w-full h-11 px-4 pr-12 bg-canvas border border-border-default rounded-xl text-text-primary focus:border-primary focus:ring-1 focus:ring-primary outline-none transition-all"
                  :class="{ 'border-status-critical': errors['password'] }"
                />
                <button
                  type="button"
                  @click="showPassword = !showPassword"
                  class="absolute right-3 top-1/2 -translate-y-1/2 text-text-secondary hover:text-primary transition-colors"
                >
                  <span class="material-symbols-outlined text-[20px]">{{
                    showPassword ? 'visibility_off' : 'visibility'
                  }}</span>
                </button>
              </div>
              <p v-if="errors['password']" class="text-status-critical text-xs mt-1">
                {{ errors['password'] }}
              </p>
              <!-- Strength indicator -->
              <div v-if="form.password" class="flex gap-1 mt-2">
                <div
                  v-for="i in 4"
                  :key="i"
                  class="h-1 flex-1 rounded-full transition-colors"
                  :class="passwordStrength >= i ? strengthColor : 'bg-border-default'"
                ></div>
              </div>
              <p v-if="form.password" class="text-xs mt-1" :class="strengthTextColor">
                {{ strengthLabel }}
              </p>
            </div>

            <div>
              <label
                class="block text-xs font-bold text-text-secondary uppercase tracking-wider mb-1.5"
                >Confirm Password *</label
              >
              <input
                v-model="form.confirmPassword"
                :type="showPassword ? 'text' : 'password'"
                placeholder="Repeat password"
                class="w-full h-11 px-4 bg-canvas border border-border-default rounded-xl text-text-primary focus:border-primary focus:ring-1 focus:ring-primary outline-none transition-all"
                :class="{ 'border-status-critical': errors['confirmPassword'] }"
              />
              <p v-if="errors['confirmPassword']" class="text-status-critical text-xs mt-1">
                {{ errors['confirmPassword'] }}
              </p>
            </div>
          </div>
        </div>

        <!-- Step 2: Organization Details -->
        <div v-if="step === 2" class="space-y-6">
          <div class="text-center">
            <div
              class="w-16 h-16 rounded-2xl bg-primary-container text-primary flex items-center justify-center mx-auto mb-4"
            >
              <span class="material-symbols-outlined text-[32px]">business</span>
            </div>
            <h2 class="font-h1 text-h1 text-text-primary mb-1">Your Organization</h2>
            <p class="text-text-secondary text-sm">
              Tell us about your company so we can configure TASM for you.
            </p>
          </div>

          <div class="space-y-4">
            <div>
              <label
                class="block text-xs font-bold text-text-secondary uppercase tracking-wider mb-1.5"
                >Company Name *</label
              >
              <input
                v-model="form.companyName"
                type="text"
                placeholder="Acme Corporation"
                class="w-full h-11 px-4 bg-canvas border border-border-default rounded-xl text-text-primary focus:border-primary focus:ring-1 focus:ring-primary outline-none transition-all"
                :class="{ 'border-status-critical': errors['companyName'] }"
              />
              <p v-if="errors['companyName']" class="text-status-critical text-xs mt-1">
                {{ errors['companyName'] }}
              </p>
            </div>

            <div class="grid grid-cols-2 gap-4">
              <div>
                <label
                  class="block text-xs font-bold text-text-secondary uppercase tracking-wider mb-1.5"
                  >Industry</label
                >
                <select
                  v-model="form.industry"
                  class="w-full h-11 px-4 bg-canvas border border-border-default rounded-xl text-text-primary focus:border-primary outline-none transition-all cursor-pointer appearance-none"
                >
                  <option value="">Select industry</option>
                  <option v-for="ind in industries" :key="ind" :value="ind">{{ ind }}</option>
                </select>
              </div>
              <div>
                <label
                  class="block text-xs font-bold text-text-secondary uppercase tracking-wider mb-1.5"
                  >Company Size</label
                >
                <select
                  v-model="form.companySize"
                  class="w-full h-11 px-4 bg-canvas border border-border-default rounded-xl text-text-primary focus:border-primary outline-none transition-all cursor-pointer appearance-none"
                >
                  <option value="">Select size</option>
                  <option value="1-10">1–10 employees</option>
                  <option value="11-50">11–50 employees</option>
                  <option value="51-200">51–200 employees</option>
                  <option value="201-1000">201–1000 employees</option>
                  <option value="1000+">1000+ employees</option>
                </select>
              </div>
            </div>

            <div class="grid grid-cols-2 gap-4">
              <div>
                <label
                  class="block text-xs font-bold text-text-secondary uppercase tracking-wider mb-1.5"
                  >Base Currency *</label
                >
                <select
                  v-model="form.currency"
                  class="w-full h-11 px-4 bg-canvas border border-border-default rounded-xl text-text-primary focus:border-primary outline-none transition-all cursor-pointer appearance-none"
                >
                  <option v-for="c in currencies" :key="c.code" :value="c.code">
                    {{ c.symbol }} {{ c.name }}
                  </option>
                </select>
              </div>
              <div>
                <label
                  class="block text-xs font-bold text-text-secondary uppercase tracking-wider mb-1.5"
                  >Timezone</label
                >
                <select
                  v-model="form.timezone"
                  class="w-full h-11 px-4 bg-canvas border border-border-default rounded-xl text-text-primary focus:border-primary outline-none transition-all cursor-pointer appearance-none"
                >
                  <option v-for="tz in timezones" :key="tz.value" :value="tz.value">
                    {{ tz.label }}
                  </option>
                </select>
              </div>
            </div>

            <div>
              <label
                class="block text-xs font-bold text-text-secondary uppercase tracking-wider mb-1.5"
                >Primary Address</label
              >
              <textarea
                v-model="form.address"
                rows="2"
                placeholder="123 Business Park, City, Country"
                class="w-full px-4 py-3 bg-canvas border border-border-default rounded-xl text-text-primary focus:border-primary focus:ring-1 focus:ring-primary outline-none transition-all resize-none"
              ></textarea>
            </div>
          </div>
        </div>

        <!-- Step 3: Review & Launch -->
        <div v-if="step === 3" class="space-y-6">
          <div class="text-center">
            <div
              class="w-16 h-16 rounded-2xl bg-primary-container text-primary flex items-center justify-center mx-auto mb-4"
            >
              <span class="material-symbols-outlined text-[32px]">rocket_launch</span>
            </div>
            <h2 class="font-h1 text-h1 text-text-primary mb-1">Ready to Launch</h2>
            <p class="text-text-secondary text-sm">
              Review your details before creating your workspace.
            </p>
          </div>

          <div class="space-y-3">
            <div class="bg-canvas border border-border-default rounded-2xl p-4 space-y-3">
              <p class="text-xs font-bold text-text-secondary uppercase tracking-wider">
                Administrator Account
              </p>
              <div class="flex items-center gap-3">
                <div
                  class="w-10 h-10 rounded-full bg-primary-container text-primary font-bold flex items-center justify-center"
                >
                  {{ form.name?.charAt(0)?.toUpperCase() || '?' }}
                </div>
                <div>
                  <p class="font-bold text-text-primary text-sm">{{ form.name }}</p>
                  <p class="text-xs text-text-secondary">{{ form.email }}</p>
                </div>
                <span
                  class="ml-auto text-xs font-bold px-2 py-0.5 bg-primary/10 text-primary rounded-full"
                  >System Admin</span
                >
              </div>
            </div>

            <div class="bg-canvas border border-border-default rounded-2xl p-4 space-y-3">
              <p class="text-xs font-bold text-text-secondary uppercase tracking-wider">
                Organization
              </p>
              <div class="grid grid-cols-2 gap-y-2 text-sm">
                <span class="text-text-secondary">Company</span
                ><span class="font-medium text-text-primary">{{ form.companyName }}</span>
                <span class="text-text-secondary">Industry</span
                ><span class="font-medium text-text-primary">{{ form.industry || '—' }}</span>
                <span class="text-text-secondary">Size</span
                ><span class="font-medium text-text-primary">{{ form.companySize || '—' }}</span>
                <span class="text-text-secondary">Currency</span
                ><span class="font-medium text-text-primary">{{ form.currency }}</span>
                <span class="text-text-secondary">Timezone</span
                ><span class="font-medium text-text-primary">{{ form.timezone }}</span>
              </div>
            </div>

            <div class="bg-primary/5 border border-primary/20 rounded-2xl p-4 flex gap-3">
              <span class="material-symbols-outlined text-primary">info</span>
              <div class="text-sm text-text-secondary">
                <p class="font-bold text-text-primary mb-0.5">What happens next?</p>
                After this you'll be taken to the <strong>Setup Wizard</strong> to configure
                locations, asset categories, and feature modules.
              </div>
            </div>
          </div>

          <div
            v-if="error"
            class="bg-error-container border border-red-200 rounded-xl p-3 flex gap-2 text-status-critical text-sm"
          >
            <span class="material-symbols-outlined text-[18px]">error</span>
            {{ error }}
          </div>
        </div>

        <!-- Footer Actions -->
        <div class="flex items-center gap-3 mt-8 pt-6 border-t border-border-default">
          <button
            v-if="step > 1"
            @click="step--"
            class="flex items-center gap-2 text-text-secondary hover:text-text-primary font-bold transition-colors"
          >
            <span class="material-symbols-outlined">arrow_back</span> Back
          </button>
          <div class="flex-1"></div>
          <button
            v-if="step < 3"
            @click="nextStep"
            class="bg-primary text-on-primary px-8 h-11 rounded-xl font-bold shadow-lg shadow-primary/20 hover:opacity-90 transition-all active:scale-95 flex items-center gap-2 disabled:opacity-50 disabled:cursor-not-allowed"
            :disabled="!isStepValid"
          >
            Continue <span class="material-symbols-outlined">arrow_forward</span>
          </button>
          <button
            v-else
            @click="handleSubmit"
            :disabled="loading"
            class="bg-primary text-on-primary px-8 h-11 rounded-xl font-bold shadow-lg shadow-primary/20 hover:opacity-90 transition-all active:scale-95 flex items-center gap-2 disabled:opacity-50"
          >
            <span v-if="loading" class="material-symbols-outlined animate-spin">sync</span>
            {{ loading ? 'Creating workspace...' : 'Create Workspace' }}
            <span v-if="!loading" class="material-symbols-outlined">rocket_launch</span>
          </button>
        </div>
      </div>

      <p class="text-center text-xs text-text-secondary mt-4">
        Already have an account?
        <RouterLink to="/login" class="text-primary hover:underline font-bold">Sign in</RouterLink>
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';
import { useRouter } from 'vue-router';
import { RouterLink } from 'vue-router';
import { useAuth } from '../composables/useAuth';

const router = useRouter();
const { setToken, checkSetupStatus } = useAuth();

const step = ref(1);
const loading = ref(false);
const error = ref('');
const showPassword = ref(false);

const form = ref({
  name: '',
  email: '',
  employeeId: '',
  department: '',
  password: '',
  confirmPassword: '',
  companyName: '',
  industry: '',
  companySize: '',
  currency: 'USD',
  timezone: 'Asia/Kolkata',
  address: '',
});

const errors = ref<Record<string, string>>({});

const industries = [
  'Technology',
  'Healthcare',
  'Education',
  'Manufacturing',
  'Finance & Banking',
  'Retail',
  'Real Estate',
  'Government',
  'Non-Profit',
  'Hospitality',
  'Other',
];

const currencies = [
  { code: 'USD', name: 'US Dollar', symbol: '$' },
  { code: 'EUR', name: 'Euro', symbol: '€' },
  { code: 'INR', name: 'Indian Rupee', symbol: '₹' },
  { code: 'GBP', name: 'Pound Sterling', symbol: '£' },
  { code: 'AUD', name: 'Australian Dollar', symbol: 'A$' },
  { code: 'CAD', name: 'Canadian Dollar', symbol: 'C$' },
  { code: 'SGD', name: 'Singapore Dollar', symbol: 'S$' },
  { code: 'AED', name: 'UAE Dirham', symbol: 'AED' },
];

const timezones = [
  { value: 'Asia/Kolkata', label: 'Asia/Kolkata (IST UTC+5:30)' },
  { value: 'America/New_York', label: 'America/New_York (EST UTC-5)' },
  { value: 'America/Los_Angeles', label: 'America/Los_Angeles (PST UTC-8)' },
  { value: 'Europe/London', label: 'Europe/London (GMT UTC+0)' },
  { value: 'Europe/Berlin', label: 'Europe/Berlin (CET UTC+1)' },
  { value: 'Asia/Singapore', label: 'Asia/Singapore (SGT UTC+8)' },
  { value: 'Asia/Dubai', label: 'Asia/Dubai (GST UTC+4)' },
  { value: 'Australia/Sydney', label: 'Australia/Sydney (AEST UTC+10)' },
];

// Password strength
const passwordStrength = computed(() => {
  const p = form.value.password;
  if (!p) return 0;
  let score = 0;
  if (p.length >= 8) score++;
  if (/[A-Z]/.test(p)) score++;
  if (/[0-9]/.test(p)) score++;
  if (/[^A-Za-z0-9]/.test(p)) score++;
  return score;
});

const strengthColor = computed(() => {
  const s = passwordStrength.value;
  if (s <= 1) return 'bg-status-critical';
  if (s === 2) return 'bg-amber-400';
  if (s === 3) return 'bg-yellow-400';
  return 'bg-status-in-stock';
});
const strengthTextColor = computed(() => {
  const s = passwordStrength.value;
  if (s <= 1) return 'text-status-critical';
  if (s <= 2) return 'text-amber-600';
  return 'text-status-in-stock';
});
const strengthLabel = computed(
  () => ['', 'Weak', 'Fair', 'Good', 'Strong'][passwordStrength.value],
);

// Step validation
const validateStep1 = () => {
  const e: Record<string, string> = {};
  if (!form.value.name.trim()) e['name'] = 'Full name is required';
  if (!form.value.email.trim()) e['email'] = 'Email is required';
  else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(form.value.email))
    e['email'] = 'Enter a valid email address';
  if (!form.value.password) e['password'] = 'Password is required';
  else if (form.value.password.length < 8) e['password'] = 'Password must be at least 8 characters';
  if (form.value.password !== form.value.confirmPassword)
    e['confirmPassword'] = 'Passwords do not match';
  errors.value = e;
  return Object.keys(e).length === 0;
};

const validateStep2 = () => {
  const e: Record<string, string> = {};
  if (!form.value.companyName.trim()) e['companyName'] = 'Company name is required';
  errors.value = e;
  return Object.keys(e).length === 0;
};

const isStepValid = computed(() => {
  if (step.value === 1) {
    return (
      form.value.name &&
      form.value.email &&
      form.value.password &&
      form.value.password === form.value.confirmPassword &&
      form.value.password.length >= 8
    );
  }
  if (step.value === 2) return !!form.value.companyName.trim();
  return true;
});

const nextStep = () => {
  if (step.value === 1 && !validateStep1()) return;
  if (step.value === 2 && !validateStep2()) return;
  step.value++;
};

const handleSubmit = async () => {
  if (!validateStep1() || !validateStep2()) {
    step.value = !validateStep1() ? 1 : 2;
    return;
  }

  loading.value = true;
  error.value = '';
  try {
    const response = await fetch('/api/auth/create-admin', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        name: form.value.name.trim(),
        email: form.value.email.trim(),
        password: form.value.password,
        employeeId: form.value.employeeId.trim() || 'ADMIN-001',
        department: form.value.department.trim() || 'Administration',
        companyName: form.value.companyName.trim(),
      }),
    });

    const data = await response.json();

    if (!response.ok) {
      error.value = data.error || 'Failed to create account. Please try again.';
      return;
    }

    // Store the JWT and re-check setup status
    setToken(data.token);
    await checkSetupStatus();

    // Navigate to the setup wizard
    router.push({ name: 'Setup' });
  } catch (err) {
    console.error('First run setup failed:', err);
    error.value = 'Network error. Please check your connection and try again.';
  } finally {
    loading.value = false;
  }
};
</script>
