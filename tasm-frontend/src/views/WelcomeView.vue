<template>
  <div class="min-h-screen bg-surface flex flex-col overflow-hidden relative">
    <!-- Abstract Background Elements -->
    <div
      class="absolute top-[-10%] left-[-10%] w-[50%] h-[50%] bg-primary/10 rounded-full blur-[120px] pointer-events-none"
    ></div>
    <div
      class="absolute bottom-[-10%] right-[-10%] w-[40%] h-[40%] bg-metric-amber/10 rounded-full blur-[120px] pointer-events-none"
    ></div>

    <header class="w-full p-6 flex justify-between items-center z-10 relative max-w-7xl mx-auto">
      <div class="flex items-center gap-3">
        <div
          class="w-10 h-10 rounded-xl bg-primary-container text-primary flex items-center justify-center"
        >
          <span class="material-symbols-outlined text-[24px]">inventory_2</span>
        </div>
        <span class="font-h2 text-h2 font-bold text-text-primary tracking-tight"
          >Technopark AMS</span
        >
      </div>
    </header>

    <main class="flex-1 flex flex-col items-center justify-center p-6 z-10 relative">
      <div class="max-w-3xl text-center space-y-8 animate-fade-in-up">
        <h1 class="text-5xl md:text-7xl font-bold text-text-primary tracking-tight leading-tight">
          Track assets with <br class="hidden md:block" />
          <span class="text-transparent bg-clip-text bg-gradient-to-r from-primary to-metric-amber"
            >precision &amp; clarity.</span
          >
        </h1>

        <p class="text-lg md:text-xl text-text-secondary max-w-2xl mx-auto font-body">
          The unified platform for enterprise asset lifecycle management. From procurement to
          depreciation, keep your organization running smoothly.
        </p>

        <div class="flex flex-col sm:flex-row items-center justify-center gap-4 pt-8">
          <button
            @click="handleGetStarted"
            class="w-full sm:w-auto h-14 px-8 bg-primary text-on-primary rounded-xl font-bold shadow-lg shadow-primary/20 hover:opacity-90 hover:-translate-y-1 transition-all active:scale-95 flex items-center justify-center gap-2 text-lg"
          >
            Sign In
            <span class="material-symbols-outlined text-[20px]">arrow_forward</span>
          </button>

          <button
            @click="router.push({ name: 'Register' })"
            class="w-full sm:w-auto h-14 px-8 bg-surface border-2 border-border-default text-text-primary rounded-xl font-bold hover:border-primary hover:text-primary transition-all active:scale-95 flex items-center justify-center gap-2 text-lg"
          >
            Create Account
            <span class="material-symbols-outlined text-[20px]">person_add</span>
          </button>
        </div>
      </div>

      <!-- Feature highlights -->
      <div
        class="grid grid-cols-1 md:grid-cols-3 gap-6 mt-24 max-w-5xl mx-auto opacity-0 animate-fade-in-up"
        style="animation-delay: 0.2s; animation-fill-mode: forwards"
      >
        <div
          class="bg-surface border border-border-default rounded-2xl p-6 shadow-sm hover:shadow-md transition-shadow"
        >
          <div
            class="w-12 h-12 rounded-full bg-primary/10 text-primary flex items-center justify-center mb-4"
          >
            <span class="material-symbols-outlined">qr_code_scanner</span>
          </div>
          <h3 class="text-lg font-bold text-text-primary mb-2">Smart Audits</h3>
          <p class="text-sm text-text-secondary">
            Conduct rapid inventory checks with mobile-friendly barcode scanning and automated
            discrepancy reporting.
          </p>
        </div>
        <div
          class="bg-surface border border-border-default rounded-2xl p-6 shadow-sm hover:shadow-md transition-shadow"
        >
          <div
            class="w-12 h-12 rounded-full bg-metric-amber/10 text-metric-amber flex items-center justify-center mb-4"
          >
            <span class="material-symbols-outlined">build</span>
          </div>
          <h3 class="text-lg font-bold text-text-primary mb-2">Maintenance</h3>
          <p class="text-sm text-text-secondary">
            Track work orders, manage vendor contracts, and schedule preventative maintenance
            effortlessly.
          </p>
        </div>
        <div
          class="bg-surface border border-border-default rounded-2xl p-6 shadow-sm hover:shadow-md transition-shadow"
        >
          <div
            class="w-12 h-12 rounded-full bg-metric-emerald/10 text-metric-emerald flex items-center justify-center mb-4"
          >
            <span class="material-symbols-outlined">account_balance</span>
          </div>
          <h3 class="text-lg font-bold text-text-primary mb-2">Financials</h3>
          <p class="text-sm text-text-secondary">
            Automate depreciation schedules, track lease agreements, and integrate with your general
            ledger.
          </p>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';

const router = useRouter();

const handleGetStarted = () => {
  router.push({ name: 'Login' });
};
</script>

<style scoped>
@keyframes fade-in-up {
  0% {
    opacity: 0;
    transform: translateY(20px);
  }
  100% {
    opacity: 1;
    transform: translateY(0);
  }
}

.animate-fade-in-up {
  animation: fade-in-up 0.6s cubic-bezier(0.16, 1, 0.3, 1) forwards;
}
</style>
