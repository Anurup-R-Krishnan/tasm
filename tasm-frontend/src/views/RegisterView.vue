<template>
  <div class="min-h-screen flex flex-col md:flex-row bg-canvas font-body overflow-hidden">
    <!-- Left Section: Premium Branding (Shared with Login) -->
    <div
      class="relative hidden md:flex md:w-1/2 lg:w-[60%] flex-col items-center justify-center p-12 overflow-hidden bg-primary text-white"
    >
      <!-- Background Image with Overlay -->
      <img
        :src="heroImage"
        alt="Branding"
        class="absolute inset-0 w-full h-full object-cover opacity-40 mix-blend-overlay"
      />

      <!-- Abstract decorative elements -->
      <div
        class="absolute top-[-10%] left-[-10%] w-[50%] h-[50%] bg-metric-amber/20 rounded-full blur-[120px] pointer-events-none"
      ></div>
      <div
        class="absolute bottom-[-10%] right-[-10%] w-[40%] h-[40%] bg-white/10 rounded-full blur-[120px] pointer-events-none"
      ></div>

      <!-- Branding Content -->
      <div
        class="relative z-10 max-w-lg text-center md:text-left space-y-8 animate-in fade-in slide-in-from-left-8 duration-1000"
      >
        <div
          class="inline-flex items-center gap-4 px-6 py-3 rounded-2xl bg-white/10 border border-white/20 backdrop-blur-md shadow-2xl"
        >
          <div
            class="w-12 h-12 rounded-xl bg-white text-primary flex items-center justify-center shadow-lg"
          >
            <span class="material-symbols-outlined text-[32px]">inventory_2</span>
          </div>
          <div>
            <h1 class="text-2xl font-bold tracking-tight text-white mb-0 uppercase">
              Technopark Assets
            </h1>
            <p class="text-[10px] font-bold text-white/70 uppercase tracking-[0.2em] mt-0.5">
              Secure Operations & Management
            </p>
          </div>
        </div>

        <div class="space-y-6">
          <h2 class="text-5xl lg:text-7xl font-bold leading-tight tracking-tight">
            Welcome to the <span class="text-metric-amber">future</span> of asset management.
          </h2>
          <p class="text-xl text-white/80 font-medium max-w-md">
            The next generation platform for Technopark infrastructure.
          </p>
        </div>
      </div>

      <!-- Bottom Credits -->
      <div
        class="absolute bottom-12 left-12 z-10 flex items-center gap-3 text-white/40 font-metadata text-[10px] uppercase tracking-[0.3em]"
      >
        <span>Precision</span>
        <div class="w-1 h-1 rounded-full bg-white/20"></div>
        <span>Clarity</span>
        <div class="w-1 h-1 rounded-full bg-white/20"></div>
        <span>Control</span>
      </div>
    </div>

    <!-- Right Section: Registration Form -->
    <div
      class="flex-1 flex flex-col items-center justify-center p-8 md:p-16 lg:p-24 bg-surface relative overflow-y-auto"
    >
      <!-- Mobile header -->
      <div class="md:hidden w-full max-w-md mb-12 text-center">
        <div
          class="inline-flex items-center justify-center w-16 h-16 rounded-2xl bg-primary-container text-primary mb-6 shadow-xl"
        >
          <span class="material-symbols-outlined text-[32px]">inventory_2</span>
        </div>
        <h1 class="text-3xl font-bold text-text-primary tracking-tight">
          {{ companyName || 'Asset Management' }}
        </h1>
      </div>

      <div class="w-full max-w-md space-y-10 animate-in fade-in slide-in-from-right-8 duration-700">
        <div class="space-y-2">
          <h3 class="text-3xl font-bold text-text-primary tracking-tight">Create an account</h3>
          <p class="text-text-secondary font-medium">
            Get started with your enterprise asset management journey.
          </p>
        </div>

        <form @submit.prevent="handleSubmit" class="space-y-6">
          <div
            v-if="error"
            class="p-4 rounded-2xl bg-status-critical/10 border border-status-critical/20 flex items-start gap-3"
          >
            <span class="material-symbols-outlined text-status-critical text-[20px]">error</span>
            <p class="text-xs font-bold text-status-critical">{{ error }}</p>
          </div>
          <div
            v-if="success"
            class="p-4 rounded-2xl bg-status-in-stock/10 border border-status-in-stock/20 flex items-start gap-3"
          >
            <span class="material-symbols-outlined text-status-in-stock text-[20px]"
              >check_circle</span
            >
            <p class="text-xs font-bold text-status-in-stock">{{ success }}</p>
          </div>

          <div class="space-y-4">
            <div class="space-y-1.5">
              <label
                class="block text-[10px] font-black text-text-secondary uppercase tracking-widest ml-1"
                >Full Name *</label
              >
              <input
                v-model="form.name"
                type="text"
                placeholder="Jane Smith"
                class="w-full h-12 px-4 bg-canvas border border-border-default rounded-xl font-medium text-text-primary focus:outline-none focus:border-primary transition-all"
                :class="{ 'border-status-critical': errors['name'] }"
              />
              <p
                v-if="errors['name']"
                class="text-status-critical text-[10px] font-bold uppercase ml-1"
              >
                {{ errors['name'] }}
              </p>
            </div>

            <div class="space-y-1.5">
              <label
                class="block text-[10px] font-black text-text-secondary uppercase tracking-widest ml-1"
                >Email Address *</label
              >
              <input
                v-model="form.email"
                type="email"
                placeholder="jane@company.com"
                class="w-full h-12 px-4 bg-canvas border border-border-default rounded-xl font-medium text-text-primary focus:outline-none focus:border-primary transition-all"
                :class="{ 'border-status-critical': errors['email'] }"
              />
              <p
                v-if="errors['email']"
                class="text-status-critical text-[10px] font-bold uppercase ml-1"
              >
                {{ errors['email'] }}
              </p>
            </div>

            <div class="grid grid-cols-2 gap-4">
              <div class="space-y-1.5">
                <label
                  class="block text-[10px] font-black text-text-secondary uppercase tracking-widest ml-1"
                  >Employee ID</label
                >
                <input
                  v-model="form.employeeId"
                  type="text"
                  placeholder="EMP-001"
                  class="w-full h-12 px-4 bg-canvas border border-border-default rounded-xl font-medium text-text-primary focus:outline-none focus:border-primary transition-all"
                />
              </div>
              <div class="space-y-1.5">
                <label
                  class="block text-[10px] font-black text-text-secondary uppercase tracking-widest ml-1"
                  >Department</label
                >
                <input
                  v-model="form.department"
                  type="text"
                  placeholder="IT Services"
                  class="w-full h-12 px-4 bg-canvas border border-border-default rounded-xl font-medium text-text-primary focus:outline-none focus:border-primary transition-all"
                />
              </div>
            </div>

            <div class="space-y-1.5">
              <label
                class="block text-[10px] font-black text-text-secondary uppercase tracking-widest ml-1"
                >Password *</label
              >
              <div class="relative group">
                <input
                  v-model="form.password"
                  :type="showPassword ? 'text' : 'password'"
                  placeholder="Min. 8 characters"
                  class="w-full h-12 px-4 pr-12 bg-canvas border border-border-default rounded-xl font-medium text-text-primary focus:outline-none focus:border-primary transition-all"
                  :class="{ 'border-status-critical': errors['password'] }"
                />
                <button
                  type="button"
                  @click="showPassword = !showPassword"
                  class="absolute right-4 top-1/2 -translate-y-1/2 text-text-disabled hover:text-primary transition-colors"
                >
                  <span class="material-symbols-outlined text-[20px]">{{
                    showPassword ? 'visibility_off' : 'visibility'
                  }}</span>
                </button>
              </div>

              <!-- Strength indicator -->
              <div v-if="form.password" class="flex gap-1 mt-2 px-1">
                <div
                  v-for="i in 4"
                  :key="i"
                  class="h-1 flex-1 rounded-full transition-colors"
                  :class="passwordStrength >= i ? strengthColor : 'bg-border-default'"
                ></div>
              </div>
              <p
                v-if="errors['password']"
                class="text-status-critical text-[10px] font-bold uppercase ml-1 mt-1"
              >
                {{ errors['password'] }}
              </p>
            </div>

            <div class="space-y-1.5">
              <label
                class="block text-[10px] font-black text-text-secondary uppercase tracking-widest ml-1"
                >Confirm Password *</label
              >
              <input
                v-model="form.confirmPassword"
                :type="showPassword ? 'text' : 'password'"
                placeholder="Repeat password"
                class="w-full h-12 px-4 bg-canvas border border-border-default rounded-xl font-medium text-text-primary focus:outline-none focus:border-primary transition-all"
                :class="{ 'border-status-critical': errors['confirmPassword'] }"
              />
              <p
                v-if="errors['confirmPassword']"
                class="text-status-critical text-[10px] font-bold uppercase ml-1"
              >
                {{ errors['confirmPassword'] }}
              </p>
            </div>
          </div>

          <button
            type="submit"
            :disabled="loading || !isValid"
            class="w-full h-14 bg-primary text-on-primary rounded-xl font-bold shadow-xl shadow-primary/20 hover:opacity-90 hover:-translate-y-0.5 active:translate-y-0 transition-all disabled:opacity-50 disabled:cursor-not-allowed flex items-center justify-center gap-3 mt-4"
          >
            <span v-if="loading" class="material-symbols-outlined animate-spin"
              >progress_activity</span
            >
            <span>{{ loading ? 'Creating account...' : 'Create Account' }}</span>
            <span v-if="!loading" class="material-symbols-outlined text-[20px]">person_add</span>
          </button>
        </form>

        <div class="pt-8 border-t border-border-default text-center">
          <p class="text-sm text-text-secondary font-medium">
            Already have an account?
            <span
              @click="router.push({ name: 'Login' })"
              class="text-primary font-black hover:underline cursor-pointer ml-1"
              >Sign In</span
            >
          </p>
        </div>
      </div>

      <!-- Footer Credits -->
      <div class="mt-12 text-center">
        <p class="text-[10px] font-bold text-text-disabled uppercase tracking-[0.4em]">
          &copy; {{ new Date().getFullYear() }} {{ companyName || 'Asset Management' }}
        </p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';
import { useRouter } from 'vue-router';
import { useAuth } from '../composables/useAuth';
import heroImage from '../assets/hero.png';

const router = useRouter();
const { register, isSetupCompleted, companyName } = useAuth();

const loading = ref(false);
const error = ref('');
const success = ref('');
const showPassword = ref(false);

const form = ref({
  name: '',
  email: '',
  employeeId: '',
  department: '',
  password: '',
  confirmPassword: '',
});

const errors = ref<Record<string, string>>({});

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

const validateForm = () => {
  const e: Record<string, string> = {};
  if (!form.value.name.trim()) e['name'] = 'Full name is required';
  if (!form.value.email.trim()) e['email'] = 'Email is required';
  else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(form.value.email))
    e['email'] = 'Enter a valid email address';

  if (!form.value.password) e['password'] = 'Password is required';
  else {
    const p = form.value.password;
    if (p.length < 8) e['password'] = 'Min. 8 characters';
    else if (!/[A-Z]/.test(p)) e['password'] = 'Must contain uppercase';
    else if (!/[a-z]/.test(p)) e['password'] = 'Must contain lowercase';
    else if (!/[0-9]/.test(p)) e['password'] = 'Must contain a number';
  }

  if (form.value.password !== form.value.confirmPassword) {
    e['confirmPassword'] = 'Passwords do not match';
  }

  errors.value = e;
  return Object.keys(e).length === 0;
};

const isValid = computed(() => {
  return (
    form.value.name.trim() &&
    form.value.email.trim() &&
    form.value.password &&
    form.value.password === form.value.confirmPassword &&
    form.value.password.length >= 8 &&
    /[A-Z]/.test(form.value.password) &&
    /[a-z]/.test(form.value.password) &&
    /[0-9]/.test(form.value.password)
  );
});

const handleSubmit = async () => {
  if (!validateForm()) return;

  loading.value = true;
  error.value = '';
  success.value = '';

  try {
    await register({
      name: form.value.name.trim(),
      email: form.value.email.trim(),
      employeeId: form.value.employeeId.trim(),
      department: form.value.department.trim(),
      password: form.value.password,
    });
    success.value = 'Account created successfully! Redirecting...';
    setTimeout(() => {
      if (!isSetupCompleted.value) {
        router.push({ name: 'FirstRun' });
      } else {
        router.push({ name: 'Dashboard' });
      }
    }, 1000);
  } catch (err: any) {
    error.value = err.message || 'Registration failed. Please try again.';
  } finally {
    loading.value = false;
  }
};
</script>

<style scoped>
.animate-in {
  animation-duration: 0.7s;
  animation-fill-mode: both;
}

@keyframes slide-in-left {
  from {
    opacity: 0;
    transform: translateX(-20px);
  }
  to {
    opacity: 1;
    transform: translateX(0);
  }
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

.slide-in-from-left-8 {
  animation-name: slide-in-left;
}

.slide-in-from-right-8 {
  animation-name: slide-in-right;
}
</style>
