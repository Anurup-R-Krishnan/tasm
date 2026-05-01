<template>
  <main class="p-page-margin max-w-[1200px] mx-auto space-y-section-gap pb-24">
    <!-- Stepper Header -->
    <div class="flex items-center justify-center gap-4 mb-12">
      <div class="flex items-center gap-2 text-slate-400">
        <div
          class="w-8 h-8 rounded-full border-2 border-slate-200 flex items-center justify-center text-xs font-bold"
        >
          1
        </div>
        <span class="text-xs font-bold uppercase tracking-wider">Select Assets</span>
      </div>
      <div class="w-12 h-[2px] bg-slate-100"></div>
      <div class="flex items-center gap-2 text-indigo-600">
        <div
          class="w-8 h-8 rounded-full border-2 border-indigo-600 flex items-center justify-center text-xs font-bold"
        >
          2
        </div>
        <span class="text-xs font-bold uppercase tracking-wider">Assign & confirm</span>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-12">
      <!-- Left: Configuration Form -->
      <div class="lg:col-span-7 space-y-8">
        <div>
          <h1 class="text-text-primary mb-2">Checkout Assignment</h1>
          <p class="text-text-secondary">
            Assign the selected assets to an employee and set a return schedule.
          </p>
        </div>

        <div class="premium-card space-y-6">
          <div class="space-y-4">
            <label class="text-[11px] font-bold text-slate-400 uppercase tracking-widest px-1"
              >Recipient Employee</label
            >
            <div class="relative group">
              <span
                class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-slate-400 group-focus-within:text-indigo-500"
                >person_search</span
              >
              <select
                v-model="selectedUserId"
                class="w-full bg-slate-50 border border-slate-100 rounded-xl py-3 pl-10 pr-4 text-sm focus:bg-white focus:ring-2 focus:ring-indigo-500/20 transition-all outline-none appearance-none cursor-pointer"
              >
                <option value="" disabled>Search or select employee...</option>
                <option v-for="user in users" :key="user.id" :value="user.id">
                  {{ user.name }} ({{ user.employeeId }} - {{ user.department }})
                </option>
              </select>
              <span
                class="material-symbols-outlined absolute right-3 top-1/2 -translate-y-1/2 text-slate-400 pointer-events-none"
                >expand_more</span
              >
            </div>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div class="space-y-4">
              <label class="text-[11px] font-bold text-slate-400 uppercase tracking-widest px-1"
                >Checkout Date</label
              >
              <input
                v-model="checkoutDate"
                type="date"
                class="w-full bg-slate-50 border border-slate-100 rounded-xl py-3 px-4 text-sm focus:bg-white focus:ring-2 focus:ring-indigo-500/20 transition-all outline-none"
              />
            </div>
            <div class="space-y-4">
              <label class="text-[11px] font-bold text-slate-400 uppercase tracking-widest px-1"
                >Expected Return</label
              >
              <input
                v-model="returnDate"
                type="date"
                class="w-full bg-slate-50 border border-slate-100 rounded-xl py-3 px-4 text-sm focus:bg-white focus:ring-2 focus:ring-indigo-500/20 transition-all outline-none"
              />
            </div>
          </div>

          <div class="space-y-4">
            <label class="text-[11px] font-bold text-slate-400 uppercase tracking-widest px-1"
              >Handover Notes</label
            >
            <textarea
              v-model="notes"
              rows="3"
              class="w-full bg-slate-50 border border-slate-100 rounded-xl py-3 px-4 text-sm focus:bg-white focus:ring-2 focus:ring-indigo-500/20 transition-all outline-none resize-none"
              placeholder="Enter any specific instructions or physical condition notes..."
            ></textarea>
          </div>
        </div>

        <div class="flex items-center gap-4">
          <button
            class="flex-1 py-3 bg-white border border-slate-200 rounded-xl text-sm font-bold text-slate-600 hover:bg-slate-50 transition-colors shadow-sm"
            @click="router.back()"
          >
            Back to Selection
          </button>
          <button
            class="flex-[2] btn-primary py-3 !text-sm"
            :disabled="!isFormValid"
            @click="completeCheckout"
          >
            <span class="material-symbols-outlined">verified_user</span>
            Complete Assignment
          </button>
        </div>
      </div>

      <!-- Right: Summary Sidebar -->
      <div class="lg:col-span-5">
        <div class="premium-card sticky top-24 space-y-6">
          <h3
            class="text-xs font-bold text-slate-500 uppercase tracking-widest px-1 border-b border-slate-50 pb-4"
          >
            Checkout Summary
          </h3>

          <div class="space-y-4">
            <div
              v-for="i in 2"
              :key="i"
              class="flex gap-4 p-3 bg-slate-50/50 rounded-xl border border-slate-50"
            >
              <div
                class="w-12 h-12 rounded-lg bg-white flex items-center justify-center text-slate-300"
              >
                <span class="material-symbols-outlined">laptop_mac</span>
              </div>
              <div class="flex-1 min-w-0">
                <p class="text-sm font-bold text-slate-900 truncate">MacBook Pro 16" (M2 Max)</p>
                <p class="text-[10px] font-mono text-slate-400 mt-1 tracking-wider">AST-IT-4092</p>
              </div>
            </div>
          </div>

          <div class="space-y-3 pt-6 border-t border-slate-50">
            <div class="flex justify-between text-xs font-medium">
              <span class="text-slate-400">Total Items</span>
              <span class="text-slate-900">2 Units</span>
            </div>
            <div class="flex justify-between text-xs font-medium">
              <span class="text-slate-400">Recipient</span>
              <span class="text-indigo-600">{{ selectedUserName || 'Not selected' }}</span>
            </div>
            <div class="flex justify-between text-xs font-medium">
              <span class="text-slate-400">Estimated Duration</span>
              <span class="text-slate-900">6 Months</span>
            </div>
          </div>

          <div
            class="p-4 bg-indigo-50/50 rounded-xl border border-indigo-100 flex items-start gap-3 mt-8"
          >
            <span class="material-symbols-outlined text-indigo-500 text-lg">info</span>
            <p class="text-[10px] leading-relaxed text-indigo-700 font-medium italic">
              Digital receipt will be sent to the employee's official email address for confirmation
              upon completion.
            </p>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();
const users = ref<any[]>([]);
const selectedUserId = ref('');
const checkoutDate = ref(new Date().toISOString().split('T')[0]);
const returnDate = ref('');
const notes = ref('');

const fetchData = async () => {
  try {
    const res = await fetch('http://localhost:8080/api/users');
    if (res.ok) users.value = await res.json();
  } catch (err) {
    console.error('Failed to fetch users:', err);
  }
};

const selectedUserName = computed(() => {
  const user = users.value.find((u) => u.id === selectedUserId.value);
  return user ? user.name : '';
});

const isFormValid = computed(() => {
  return selectedUserId.value && checkoutDate.value && returnDate.value;
});

const completeCheckout = () => {
  // In a real app, this would be a POST to /api/checkout
  console.log('Completing checkout...', {
    userId: selectedUserId.value,
    checkoutDate: checkoutDate.value,
    returnDate: returnDate.value,
    notes: notes.value,
  });
  router.push('/asset-registry');
};

onMounted(fetchData);
</script>
