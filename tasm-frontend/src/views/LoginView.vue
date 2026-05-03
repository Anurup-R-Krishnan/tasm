<template>
  <div
    class="min-h-screen bg-surface flex items-center justify-center p-stack overflow-hidden relative"
  >
    <!-- Abstract Background Elements -->
    <div
      class="absolute top-[-10%] left-[-10%] w-[40%] h-[40%] bg-primary/10 rounded-full blur-[120px] pointer-events-none"
    ></div>
    <div
      class="absolute bottom-[-10%] right-[-10%] w-[40%] h-[40%] bg-metric-amber/10 rounded-full blur-[120px] pointer-events-none"
    ></div>

    <div class="w-full max-w-md relative z-10">
      <!-- Logo / Header -->
      <div class="text-center mb-section-gap">
        <div
          class="inline-flex items-center justify-center w-16 h-16 rounded-2xl bg-primary-container text-primary mb-stack shadow-sm"
        >
          <span class="material-symbols-outlined text-[32px]">inventory_2</span>
        </div>
        <h1 class="font-h1 text-h1 text-text-primary mb-2">TASM</h1>
        <p class="font-body text-body text-text-secondary">Enterprise Asset Management</p>
      </div>

      <!-- Login Card -->
      <div
        class="bg-surface rounded-2xl border border-border-default shadow-[0_8px_30px_rgba(0,0,0,0.04)] p-8"
      >
        <h2 class="font-h2 text-h2 text-text-primary mb-stack">Welcome back</h2>

        <form @submit.prevent="handleLogin" class="space-y-stack">
          <!-- Email Field -->
          <div>
            <label
              for="email"
              class="block font-metadata text-metadata text-text-secondary mb-2 uppercase tracking-wider"
              >Email Address</label
            >
            <div class="relative">
              <span
                class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-text-secondary"
                >mail</span
              >
              <input
                id="email"
                v-model="email"
                type="email"
                required
                class="w-full h-11 pl-10 pr-4 bg-surface border border-border-default rounded-lg font-body text-body text-text-primary placeholder:text-text-tertiary focus:outline-none focus:border-primary focus:ring-1 focus:ring-primary transition-all"
                placeholder="you@company.com"
              />
            </div>
          </div>

          <!-- Password Field -->
          <div>
            <div class="flex items-center justify-between mb-2">
              <label
                for="password"
                class="block font-metadata text-metadata text-text-secondary uppercase tracking-wider"
                >Password</label
              >
              <a
                href="#"
                class="font-metadata text-metadata text-primary hover:text-primary-container transition-colors"
                >Forgot password?</a
              >
            </div>
            <div class="relative">
              <span
                class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-text-secondary"
                >lock</span
              >
              <input
                id="password"
                v-model="password"
                :type="showPassword ? 'text' : 'password'"
                required
                class="w-full h-11 pl-10 pr-10 bg-surface border border-border-default rounded-lg font-body text-body text-text-primary placeholder:text-text-tertiary focus:outline-none focus:border-primary focus:ring-1 focus:ring-primary transition-all"
                placeholder="••••••••"
              />
              <button
                type="button"
                @click="showPassword = !showPassword"
                class="absolute right-3 top-1/2 -translate-y-1/2 text-text-secondary hover:text-text-primary transition-colors focus:outline-none"
              >
                <span class="material-symbols-outlined">{{
                  showPassword ? 'visibility_off' : 'visibility'
                }}</span>
              </button>
            </div>
          </div>

          <!-- Error Message -->
          <div
            v-if="error"
            class="p-3 rounded-lg bg-error-container border border-red-200 flex items-start gap-2"
          >
            <span class="material-symbols-outlined text-on-error-container text-[20px]">error</span>
            <p class="font-metadata text-metadata text-on-error-container mt-[2px]">{{ error }}</p>
          </div>

          <!-- Submit Button -->
          <button
            type="submit"
            :disabled="isLoading"
            class="w-full h-11 bg-primary text-on-primary font-body text-body font-medium rounded-lg shadow-sm hover:bg-primary-container hover:-translate-y-[1px] active:translate-y-0 transition-all disabled:opacity-50 disabled:cursor-not-allowed disabled:hover:transform-none flex items-center justify-center gap-2 mt-section-gap"
          >
            <span v-if="isLoading" class="material-symbols-outlined animate-spin text-[20px]"
              >progress_activity</span
            >
            <span>{{ isLoading ? 'Signing in...' : 'Sign in' }}</span>
            <span v-if="!isLoading" class="material-symbols-outlined text-[20px]"
              >arrow_forward</span
            >
          </button>
        </form>
      </div>

      <!-- Footer -->
      <p class="text-center mt-section-gap font-metadata text-metadata text-text-tertiary">
        &copy; {{ new Date().getFullYear() }} TASM. All rights reserved.
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useAuth } from '../composables/useAuth';

const router = useRouter();
const { login } = useAuth();

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
    await login({ email: email.value, password: password.value });
    // Navigate to dashboard on success
    router.push({ name: 'Dashboard' });
  } catch (err: any) {
    error.value = err?.message || 'Invalid email or password. Please try again.';
  } finally {
    isLoading.value = false;
  }
};
</script>
