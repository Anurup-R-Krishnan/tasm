<template>
  <main class="space-y-section-gap pb-24">
    <!-- Left Column: KPIs and Table -->
    <div class="flex-1 flex flex-col gap-6">
      <!-- Header -->
      <div class="flex justify-between items-end">
        <div>
          <h1 class="font-h1 text-h1 text-text-primary mb-1">Software Licenses</h1>
          <p class="text-text-secondary font-body">Manage subscriptions, seats, and renewals</p>
        </div>
        <div class="flex gap-3">
          <button
            @click="handleExport"
            class="px-4 py-2 bg-surface text-text-primary border border-border-default rounded-lg font-medium text-sm hover:bg-surface-subtle transition-colors flex items-center gap-2 shadow-sm"
          >
            <span class="material-symbols-outlined text-[18px]"> download </span>
            Export CSV
          </button>
        </div>
      </div>
      <!-- KPI Cards Grid -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <!-- Expiring Licenses -->
        <div
          class="bg-surface border border-border-default rounded-xl p-card-padding shadow-sm hover:-translate-y-[2px] transition-transform cursor-pointer relative overflow-hidden group"
          @click="currentFilter = 'Expiring Soon'"
        >
          <div
            class="absolute inset-0 bg-gradient-to-br from-[#FFE4E6]/40 to-transparent z-0"
          ></div>
          <div class="relative z-10 flex justify-between items-start">
            <div>
              <p
                class="font-metadata text-metadata text-text-secondary uppercase tracking-wider mb-2"
              >
                Expiring Soon
              </p>
              <h2 class="font-kpi-number text-kpi-number text-text-primary mb-1">
                {{ expiringCount }}
              </h2>
              <p class="text-sm text-status-critical font-medium flex items-center gap-1">
                <span class="material-symbols-outlined text-[16px]"> warning </span>
                Action Required
              </p>
            </div>
            <div
              class="w-10 h-10 rounded-full bg-[#FFE4E6] flex items-center justify-center text-status-critical"
            >
              <span class="material-symbols-outlined"> event_busy </span>
            </div>
          </div>
        </div>
        <!-- Underutilized Seats -->
        <div
          class="bg-surface border border-border-default rounded-xl p-card-padding shadow-sm hover:-translate-y-[2px] transition-transform cursor-pointer relative overflow-hidden group"
        >
          <div
            class="absolute inset-0 bg-gradient-to-br from-metric-amber/50 to-transparent z-0"
          ></div>
          <div class="relative z-10 flex justify-between items-start">
            <div>
              <p
                class="font-metadata text-metadata text-text-secondary uppercase tracking-wider mb-2"
              >
                Underutilized Seats
              </p>
              <h2 class="font-kpi-number text-kpi-number text-text-primary mb-1">
                {{ underutilizedSeats }}
              </h2>
              <p class="text-sm text-surface-tint font-medium flex items-center gap-1">
                <span class="material-symbols-outlined text-[16px]"> trending_down </span>
                Potential savings
              </p>
            </div>
            <div
              class="w-10 h-10 rounded-full bg-metric-amber flex items-center justify-center text-surface-tint"
            >
              <span class="material-symbols-outlined"> group_remove </span>
            </div>
          </div>
        </div>
        <!-- Annual Spend -->
        <div
          class="bg-surface border border-border-default rounded-xl p-card-padding shadow-sm hover:-translate-y-[2px] transition-transform cursor-pointer relative overflow-hidden group"
        >
          <div
            class="absolute inset-0 bg-gradient-to-br from-metric-lavender/50 to-transparent z-0"
          ></div>
          <div class="relative z-10 flex justify-between items-start">
            <div>
              <p
                class="font-metadata text-metadata text-text-secondary uppercase tracking-wider mb-2"
              >
                Annual Software Spend
              </p>
              <h2 class="font-kpi-number text-kpi-number text-text-primary mb-1">
                ₹{{ annualSpend.toLocaleString() }}
              </h2>
              <p class="text-sm text-text-secondary flex items-center gap-1">Total</p>
            </div>
            <div
              class="w-10 h-10 rounded-full bg-metric-lavender flex items-center justify-center text-[#5B21B6]"
            >
              <span class="material-symbols-outlined"> payments </span>
            </div>
          </div>
        </div>
      </div>
      <!-- License Table -->
      <div
        class="bg-surface border border-border-default rounded-xl shadow-sm overflow-hidden flex-1 flex flex-col"
      >
        <div
          class="p-4 border-b border-border-default flex justify-between items-center bg-surface"
        >
          <div class="flex gap-4 items-center">
            <div class="flex gap-2">
              <button
                @click="currentFilter = 'All'"
                class="px-3 py-1.5 rounded-md text-sm font-medium transition-colors"
                :class="
                  currentFilter === 'All'
                    ? 'bg-surface-variant text-text-primary'
                    : 'text-text-secondary hover:bg-stone-50'
                "
              >
                All Licenses
              </button>
              <button
                @click="currentFilter = 'Active'"
                class="px-3 py-1.5 rounded-md text-sm font-medium transition-colors"
                :class="
                  currentFilter === 'Active'
                    ? 'bg-surface-variant text-text-primary'
                    : 'text-text-secondary hover:bg-stone-50'
                "
              >
                Active
              </button>
              <button
                @click="currentFilter = 'Expiring Soon'"
                class="px-3 py-1.5 rounded-md text-sm font-medium transition-colors"
                :class="
                  currentFilter === 'Expiring Soon'
                    ? 'bg-surface-variant text-text-primary'
                    : 'text-text-secondary hover:bg-stone-50'
                "
              >
                Expiring Soon
              </button>
            </div>
            <div class="h-6 w-[1px] bg-border-default mx-2"></div>
            <div class="relative w-64">
              <span
                class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-text-secondary text-[18px]"
                >search</span
              >
              <input
                v-model="searchQuery"
                type="text"
                placeholder="Search licenses..."
                class="w-full bg-surface-subtle border border-border-default rounded-lg pl-10 pr-4 py-1.5 text-sm focus:outline-none focus:border-primary transition-all"
              />
            </div>
          </div>
          <button class="text-text-secondary hover:text-text-primary">
            <span class="material-symbols-outlined"> filter_list </span>
          </button>
        </div>
        <div class="overflow-x-auto flex-1 p-4">
          <DataTable
            :value="filteredLicenses"
            :loading="loading"
            paginator
            :rows="10"
            tableStyle="min-width: 50rem"
            class="w-full text-left"
            @row-click="(e: any) => (selectedLicense = e.data)"
            rowHover
          >
            <Column field="softwareName" header="Software Asset" sortable>
              <template #body="slotProps">
                <div class="flex items-center gap-3">
                  <div
                    class="w-8 h-8 rounded bg-primary-container/10 flex items-center justify-center text-primary"
                  >
                    <span class="material-symbols-outlined text-[18px]">cloud</span>
                  </div>
                  <div>
                    <div class="font-medium text-text-primary">
                      {{ slotProps.data.softwareName }}
                    </div>
                    <div class="text-xs text-text-secondary">{{ slotProps.data.planName }}</div>
                  </div>
                </div>
              </template>
            </Column>
            <Column field="status" header="Status" sortable>
              <template #body="slotProps">
                <span
                  class="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium"
                  :class="{
                    'bg-status-in-stock/10 text-status-in-stock':
                      slotProps.data.status === 'Active',
                    'bg-metric-amber/20 text-surface-tint border border-metric-amber':
                      slotProps.data.status === 'Expiring Soon',
                    'bg-surface-variant text-text-secondary': slotProps.data.status === 'Expired',
                  }"
                >
                  {{ slotProps.data.status }}
                </span>
              </template>
            </Column>
            <Column field="usedSeats" header="Seat Utilization" sortable>
              <template #body="slotProps">
                <div class="flex items-center gap-2">
                  <div class="flex-1 h-1.5 bg-surface-variant rounded-full overflow-hidden">
                    <div
                      class="h-full rounded-full"
                      :class="
                        slotProps.data.usedSeats / slotProps.data.totalSeats > 0.9
                          ? 'bg-primary'
                          : 'bg-status-in-stock'
                      "
                      :style="{
                        width: `${(slotProps.data.usedSeats / slotProps.data.totalSeats) * 100}%`,
                      }"
                    ></div>
                  </div>
                  <span class="font-mono-data text-mono-data text-text-secondary w-12 text-right">
                    {{ slotProps.data.usedSeats }}/{{ slotProps.data.totalSeats }}
                  </span>
                </div>
              </template>
            </Column>
            <Column field="renewalDate" header="Renewal Date" sortable>
              <template #body="slotProps">
                <div
                  class="font-mono-data text-mono-data"
                  :class="{
                    'text-surface-tint font-medium': slotProps.data.status === 'Expiring Soon',
                  }"
                >
                  {{ new Date(slotProps.data.renewalDate).toLocaleDateString() }}
                </div>
              </template>
            </Column>
            <Column field="annualCost" header="Ann. Cost" alignFrozen="right" sortable>
              <template #body="slotProps">
                <div class="text-right font-mono-data text-mono-data">
                  ₹ {{ slotProps.data.annualCost?.toLocaleString() }}
                </div>
              </template>
            </Column>
          </DataTable>
        </div>
      </div>
    </div>
    <!-- Right Drawer (Level 2 Elevation) -->
    <div
      v-if="selectedLicense"
      class="w-[480px] bg-surface border-l border-border-default shadow-[0_0_24px_rgba(0,0,0,0.05)] rounded-l-2xl h-[calc(100vh-60px-64px)] sticky top-[32px] flex flex-col z-20 animate-fade-in"
    >
      <!-- Drawer Header -->
      <div class="p-6 border-b border-border-default flex justify-between items-start">
        <div class="flex gap-4">
          <div
            class="w-12 h-12 rounded-lg bg-primary/10 flex items-center justify-center text-primary shadow-sm border border-primary/20"
          >
            <span class="material-symbols-outlined text-[24px]"> cloud </span>
          </div>
          <div>
            <h2 class="font-h2 text-h2 text-text-primary">
              {{ selectedLicense.softwareName }}
            </h2>
            <p class="text-sm text-text-secondary mt-1">
              {{ selectedLicense.planName }}
            </p>
            <div class="mt-2 flex gap-2">
              <span
                class="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium bg-status-in-stock/10 text-status-in-stock"
              >
                {{ selectedLicense.status }}
              </span>
              <span
                class="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium bg-stone-100 text-stone-600 border border-stone-200"
              >
                ID: LIC-2023-{{ selectedLicense.id }}
              </span>
            </div>
          </div>
        </div>
        <button
          @click="selectedLicense = null"
          class="text-text-secondary hover:bg-stone-100 p-1 rounded-md transition-colors"
        >
          <span class="material-symbols-outlined"> close </span>
        </button>
      </div>
      <!-- Drawer Tabs -->
      <div class="px-6 flex gap-6 border-b border-border-default">
        <button
          @click="drawerActiveTab = 'overview'"
          class="py-3 text-sm font-medium border-b-2 transition-colors"
          :class="
            drawerActiveTab === 'overview'
              ? 'border-primary text-primary'
              : 'border-transparent text-text-secondary hover:text-text-primary'
          "
        >
          Overview
        </button>
        <button
          @click="drawerActiveTab = 'users'"
          class="py-3 text-sm font-medium border-b-2 transition-colors"
          :class="
            drawerActiveTab === 'users'
              ? 'border-primary text-primary'
              : 'border-transparent text-text-secondary hover:text-text-primary'
          "
        >
          Assigned Users ({{ assignedUsers.length }})
        </button>
        <button
          @click="drawerActiveTab = 'invoices'"
          class="py-3 text-sm font-medium border-b-2 transition-colors"
          :class="
            drawerActiveTab === 'invoices'
              ? 'border-primary text-primary'
              : 'border-transparent text-text-secondary hover:text-text-primary'
          "
        >
          Invoices
        </button>
      </div>
      <!-- Drawer Content -->
      <div class="p-6 flex-1 overflow-y-auto">
        <div v-if="drawerActiveTab === 'overview'" class="space-y-6 animate-fade-in">
          <div class="grid grid-cols-2 gap-4">
            <div class="p-3 bg-surface-subtle rounded-lg border border-border-default">
              <p class="text-[10px] font-bold text-text-secondary uppercase tracking-widest">
                Type
              </p>
              <p class="text-sm font-medium text-text-primary mt-1">SaaS Subscription</p>
            </div>
            <div class="p-3 bg-surface-subtle rounded-lg border border-border-default">
              <p class="text-[10px] font-bold text-text-secondary uppercase tracking-widest">
                Billing
              </p>
              <p class="text-sm font-medium text-text-primary mt-1">Annual Auto-renew</p>
            </div>
          </div>
          <div class="p-4 bg-primary/5 rounded-lg border border-primary/10">
            <h4 class="text-xs font-bold text-primary uppercase tracking-wider mb-2">
              Usage Summary
            </h4>
            <div class="flex justify-between items-center text-sm">
              <span class="text-text-secondary">Seats Assigned</span>
              <span class="font-bold text-text-primary"
                >{{ assignedUsers.length }} / {{ selectedLicense.totalSeats }}</span
              >
            </div>
            <div
              class="w-full h-2 bg-white rounded-full mt-2 overflow-hidden border border-primary/10"
            >
              <div
                class="h-full bg-primary"
                :style="{ width: (assignedUsers.length / selectedLicense.totalSeats) * 100 + '%' }"
              ></div>
            </div>
          </div>
        </div>

        <div v-if="drawerActiveTab === 'users'" class="space-y-4 animate-fade-in">
          <div class="flex justify-between items-center mb-4">
            <div class="relative w-64">
              <span
                class="material-symbols-outlined absolute left-2.5 top-1/2 -translate-y-1/2 text-stone-400 text-[18px]"
              >
                search
              </span>
              <input
                v-model="userSearchQuery"
                class="w-full bg-surface-subtle border border-border-default rounded-md pl-9 pr-3 py-1.5 text-sm focus:outline-none focus:border-primary"
                placeholder="Search users..."
                type="text"
              />
            </div>
            <button
              @click="handleAssignUser"
              class="text-primary hover:text-primary-hover text-sm font-medium flex items-center gap-1"
            >
              <span class="material-symbols-outlined text-[18px]"> person_add </span>
              Assign User
            </button>
          </div>
          <div class="space-y-3">
            <div
              v-for="user in filteredAssignedUsers"
              :key="user.email"
              class="flex items-center justify-between p-3 rounded-lg border border-border-default hover:bg-surface-subtle transition-colors group"
            >
              <div class="flex items-center gap-3">
                <div
                  class="w-8 h-8 rounded-full bg-surface-variant flex items-center justify-center font-bold text-primary text-[10px]"
                >
                  {{ user.initials }}
                </div>
                <div>
                  <div class="text-sm font-medium text-text-primary">{{ user.name }}</div>
                  <div class="text-xs text-text-secondary">
                    {{ user.dept }} • {{ user.location }}
                  </div>
                </div>
              </div>
              <button
                @click="handleRevokeUser(user.email)"
                class="text-stone-400 hover:text-error transition-colors"
              >
                <span class="material-symbols-outlined text-[18px]"> person_remove </span>
              </button>
            </div>
            <div
              v-if="filteredAssignedUsers.length === 0"
              class="text-center py-8 text-text-secondary italic text-sm"
            >
              No users found.
            </div>
          </div>
        </div>

        <div v-if="drawerActiveTab === 'invoices'" class="text-center py-12 animate-fade-in">
          <span class="material-symbols-outlined text-4xl text-text-secondary/30"
            >receipt_long</span
          >
          <p class="text-text-secondary text-sm mt-2">No invoices found for this license.</p>
        </div>
      </div>
      <!-- Drawer Footer Action -->
      <div class="p-6 border-t border-border-default bg-surface-subtle rounded-bl-2xl">
        <button
          @click="handleManageSubscription"
          class="w-full py-2.5 bg-primary text-on-primary rounded-lg font-medium hover:bg-primary-hover transition-colors"
        >
          Manage Subscription
        </button>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { getSoftwareLicenses } from '../api/financial';
import { getUsers } from '../api/users';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import type { SoftwareLicense, SystemUser } from '../types/models';

const licenses = ref<SoftwareLicense[]>([]);
const loading = ref(true);
const selectedLicense = ref<SoftwareLicense | null>(null);
const currentFilter = ref('All');
const searchQuery = ref('');
const drawerActiveTab = ref('overview');
const userSearchQuery = ref('');

const assignedUsers = ref<
  { name: string; email: string; dept: string; location: string; initials: string }[]
>([]);

const filteredAssignedUsers = computed(() => {
  if (!userSearchQuery.value) return assignedUsers.value;
  const q = userSearchQuery.value.toLowerCase();
  return assignedUsers.value.filter(
    (u) => u.name.toLowerCase().includes(q) || u.email.toLowerCase().includes(q),
  );
});

const expiringCount = computed(() => {
  return licenses.value.filter((l) => l.status === 'Expiring Soon').length;
});

const underutilizedSeats = computed(() => {
  return licenses.value.reduce((sum, l) => sum + (l.totalSeats - l.usedSeats), 0);
});

const annualSpend = computed(() => {
  return licenses.value.reduce((sum, l) => sum + l.annualCost, 0);
});

const filteredLicenses = computed(() => {
  let result = licenses.value;

  if (currentFilter.value !== 'All') {
    result = result.filter((l) => l.status === currentFilter.value);
  }

  if (searchQuery.value) {
    const q = searchQuery.value.toLowerCase();
    result = result.filter(
      (l) => l.softwareName.toLowerCase().includes(q) || l.planName.toLowerCase().includes(q),
    );
  }

  return result;
});

const fetchLicenses = async () => {
  loading.value = true;
  try {
    const data = await getSoftwareLicenses();
    licenses.value = data as SoftwareLicense[];
  } catch (error) {
    console.error('Failed to fetch licenses:', error);
  } finally {
    loading.value = false;
  }
};

const fetchAssignedUsers = async () => {
  try {
    const users = (await getUsers()) as SystemUser[];
    assignedUsers.value = users.map((u) => ({
      name: u.name,
      email: u.email,
      dept: u.department || 'General',
      location: u.department || 'General',
      initials: u.name
        .split(' ')
        .filter(Boolean)
        .map((n) => n[0])
        .join('')
        .toUpperCase()
        .slice(0, 2),
    }));
  } catch (error) {
    console.error('Failed to fetch users for license assignments:', error);
  }
};

const handleExport = () => {
  const headers = [
    'Software',
    'Plan',
    'Status',
    'Total Seats',
    'Used Seats',
    'Renewal Date',
    'Annual Cost',
  ];
  const rows = filteredLicenses.value.map((l) => [
    l.softwareName,
    l.planName,
    l.status,
    l.totalSeats,
    l.usedSeats,
    new Date(l.renewalDate).toLocaleDateString(),
    l.annualCost,
  ]);

  const csvContent = [
    headers.join(','),
    ...rows.map((row) => row.map((cell) => `"${cell}"`).join(',')),
  ].join('\n');

  const blob = new Blob([csvContent], { type: 'text/csv;charset=utf-8;' });
  const link = document.createElement('a');
  link.href = URL.createObjectURL(blob);
  link.download = `tasm_software_licenses_${new Date().toISOString().split('T')[0]}.csv`;
  link.click();
};

const handleManageSubscription = () => {
  if (!selectedLicense.value) return;
  alert(`Redirecting to vendor portal for ${selectedLicense.value.softwareName}...`);
  // In a real app, this might navigate to a specific subscription management page or open an external link
};

const handleRevokeUser = (email: string) => {
  if (confirm(`Are you sure you want to revoke license access for ${email}?`)) {
    assignedUsers.value = assignedUsers.value.filter((u) => u.email !== email);
  }
};

const handleAssignUser = () => {
  const name = prompt('Enter user name:');
  const email = prompt('Enter user email:');
  if (name && email) {
    assignedUsers.value.push({
      name,
      email,
      dept: 'Unassigned',
      location: 'Central',
      initials: name
        .split(' ')
        .map((n) => n[0])
        .join('')
        .toUpperCase(),
    });
  }
};

onMounted(() => {
  fetchLicenses();
  fetchAssignedUsers();
});
</script>
