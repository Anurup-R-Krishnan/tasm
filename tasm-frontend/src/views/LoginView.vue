<template>
  <div class="min-h-screen flex flex-col md:flex-row bg-canvas font-body overflow-hidden">
    <!-- Left Section: Premium Branding -->
    <div
      class="relative hidden md:flex md:w-1/2 lg:w-[60%] flex-col items-center justify-center p-12 overflow-hidden bg-primary text-white"
    >
      <!-- Background Image with Overlay -->
      <img
        :src="heroImage"
        alt="Aura Design System - TASM"
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

    <!-- Right Section: Login Form -->
    <div
      class="flex-1 flex flex-col items-center justify-center p-8 md:p-16 lg:p-24 bg-surface relative"
    >
      <!-- Mobile header (visible only on small screens) -->
      <div class="md:hidden w-full max-w-md mb-12 text-center">
        <div
          class="inline-flex items-center justify-center w-16 h-16 rounded-2xl bg-primary-container text-primary mb-6 shadow-xl"
        >
          <span class="material-symbols-outlined text-[32px]">inventory_2</span>
        </div>
        <h1 class="text-3xl font-bold text-text-primary tracking-tight">TASM</h1>
        <p class="text-text-secondary mt-2">Technopark Asset Management System</p>
      </div>

      <div class="w-full max-w-md space-y-10 animate-in fade-in slide-in-from-right-8 duration-700">
        <div class="space-y-2">
          <h3 class="text-3xl font-bold text-text-primary tracking-tight">Welcome back</h3>
          <p class="text-text-secondary font-medium">
            Please enter your credentials to access your account.
          </p>
        </div>

        <form @submit.prevent="handleLogin" class="space-y-6">
          <div class="space-y-2">
            <label
              for="email"
              class="block text-[10px] font-black text-text-secondary uppercase tracking-[0.2em] ml-1"
              >Email Address</label
            >
            <div class="relative group">
              <span
                class="material-symbols-outlined absolute left-4 top-1/2 -translate-y-1/2 text-text-disabled group-focus-within:text-primary transition-colors"
                >mail</span
              >
              <input
                id="email"
                v-model="email"
                type="email"
                required
                class="w-full h-14 pl-12 pr-4 bg-canvas border border-border-default rounded-2xl font-medium text-text-primary placeholder:text-text-disabled focus:outline-none focus:border-primary focus:ring-4 focus:ring-primary/5 transition-all"
                placeholder="alex@technopark.com"
              />
            </div>
          </div>

          <div class="space-y-2">
            <div class="flex items-center justify-between ml-1">
              <label
                for="password"
                class="block text-[10px] font-black text-text-secondary uppercase tracking-[0.2em]"
                >Password</label
              >
              <button
                type="button"
                @click="handleForgotPassword"
                class="text-[10px] font-black text-primary hover:underline uppercase tracking-widest"
              >
                Forgot Password?
              </button>
            </div>
            <div class="relative group">
              <span
                class="material-symbols-outlined absolute left-4 top-1/2 -translate-y-1/2 text-text-disabled group-focus-within:text-primary transition-colors"
                >lock</span
              >
              <input
                id="password"
                v-model="password"
                :type="showPassword ? 'text' : 'password'"
                required
                class="w-full h-14 pl-12 pr-12 bg-canvas border border-border-default rounded-2xl font-medium text-text-primary placeholder:text-text-disabled focus:outline-none focus:border-primary focus:ring-4 focus:ring-primary/5 transition-all"
                placeholder="••••••••"
              />
              <button
                type="button"
                @click="showPassword = !showPassword"
                class="absolute right-4 top-1/2 -translate-y-1/2 text-text-disabled hover:text-text-primary transition-colors outline-none"
              >
                <span class="material-symbols-outlined text-[20px]">{{
                  showPassword ? 'visibility_off' : 'visibility'
                }}</span>
              </button>
            </div>
          </div>

          <div
            v-if="error"
            class="p-4 rounded-2xl bg-status-critical/10 border border-status-critical/20 flex items-start gap-3 animate-in fade-in zoom-in-95 duration-300"
          >
            <span class="material-symbols-outlined text-status-critical text-[20px]">error</span>
            <p class="text-xs font-bold text-status-critical leading-snug">{{ error }}</p>
          </div>

          <button
            type="submit"
            :disabled="isLoading"
            class="w-full h-14 bg-primary text-on-primary font-bold rounded-2xl shadow-xl shadow-primary/20 hover:opacity-90 hover:-translate-y-0.5 active:translate-y-0 transition-all disabled:opacity-50 disabled:cursor-not-allowed flex items-center justify-center gap-3 mt-4"
          >
            <span v-if="isLoading" class="material-symbols-outlined animate-spin"
              >progress_activity</span
            >
            <span>{{ isLoading ? 'Authenticating...' : 'Sign In' }}</span>
            <span v-if="!isLoading" class="material-symbols-outlined text-[20px]">login</span>
          </button>
        </form>

        <div class="pt-8 border-t border-border-default">
          <p class="text-center text-sm text-text-secondary font-medium">
            Don't have an account yet?
            <button
              type="button"
              @click="handleRequestAccess"
              class="text-primary font-black hover:underline cursor-pointer ml-1"
            >
              Create Account
            </button>
          </p>
        </div>
      </div>

      <!-- Footer Credits -->
      <div class="absolute bottom-8 left-0 right-0 text-center">
        <p class="text-[10px] font-bold text-text-disabled uppercase tracking-[0.4em]">
          &copy; {{ new Date().getFullYear() }} {{ companyName || 'Asset Management' }}
        </p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useAuth } from '../composables/useAuth';
import heroImage from '../assets/hero.png';

const router = useRouter();
const { login, isSetupCompleted, companyName } = useAuth();

const email = ref('');
const password = ref('');
const showPassword = ref(false);
const isLoading = ref(false);
const error = ref('');

const handleLogin = async () => {
  if (!email.value || !password.value) return;

  error.value = '';
  isLoading.value = true;

  try {
    await login({ email: email.value.trim(), password: password.value });
    if (!isSetupCompleted.value) {
      router.push({ name: 'FirstRun' });
    } else {
      router.push({ name: 'Dashboard' });
    }
  } catch (err: any) {
    error.value = err?.message || 'Invalid email or password. Please try again.';
  } finally {
    isLoading.value = false;
  }
};

const handleRequestAccess = () => {
  router.push({ name: 'Register' });
};

const handleForgotPassword = () => {
  alert('Password reset is not enabled yet. Please contact your administrator.');
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
