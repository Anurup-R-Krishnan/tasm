<template>
  <main class="space-y-section-gap pb-24">
    <!-- Left Column: KPIs and Table -->
    <div class="flex-1 flex flex-col gap-6">
      <!-- Header -->
      <div class="flex justify-between items-end">
        <div>
          <h1 class="font-h1 text-h1 text-text-primary mb-1">Software Licenses</h1>
          <p class="text-text-secondary font-body text-body">
            Manage subscriptions, seats, and renewals
          </p>
        </div>
        <div class="flex gap-3">
          <button
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
          <div class="flex gap-2">
            <button
              class="px-3 py-1.5 rounded-md bg-surface-variant text-sm font-medium text-text-primary"
            >
              All Licenses
            </button>
            <button
              class="px-3 py-1.5 rounded-md hover:bg-stone-50 text-sm font-medium text-stone-600"
            >
              Active
            </button>
            <button
              class="px-3 py-1.5 rounded-md hover:bg-stone-50 text-sm font-medium text-stone-600"
            >
              Expiring Soon
            </button>
          </div>
          <button class="text-text-secondary hover:text-text-primary">
            <span class="material-symbols-outlined"> filter_list </span>
          </button>
        </div>
        <div class="overflow-x-auto flex-1 p-4">
          <DataTable
            :value="licenses"
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
      class="w-[480px] bg-surface border-l border-border-default shadow-[0_0_24px_rgba(0,0,0,0.05)] rounded-l-2xl h-[calc(100vh-60px-64px)] sticky top-[32px] flex flex-col z-20"
    >
      <!-- Drawer Header -->
      <div class="p-6 border-b border-border-default flex justify-between items-start">
        <div class="flex gap-4">
          <div
            class="w-12 h-12 rounded-lg bg-error-container/10 flex items-center justify-center text-error shadow-sm border border-error-container/20"
          >
            <span class="material-symbols-outlined text-[24px]"> brush </span>
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
        <button class="py-3 text-sm font-medium border-b-2 border-primary text-primary">
          Overview
        </button>
        <button
          class="py-3 text-sm font-medium border-b-2 border-transparent text-text-secondary hover:text-text-primary"
        >
          Assigned Users ({{ selectedLicense.usedSeats }})
        </button>
        <button
          class="py-3 text-sm font-medium border-b-2 border-transparent text-text-secondary hover:text-text-primary"
        >
          Invoices
        </button>
      </div>
      <!-- Drawer Content (Assigned Users Tab Selected Visually for Demo) -->
      <div class="p-6 flex-1 overflow-y-auto">
        <div class="flex justify-between items-center mb-4">
          <div class="relative w-64">
            <span
              class="material-symbols-outlined absolute left-2.5 top-1/2 -translate-y-1/2 text-stone-400 text-[18px]"
            >
              search
            </span>
            <input
              class="w-full bg-surface-subtle border border-border-default rounded-md pl-9 pr-3 py-1.5 text-sm focus:outline-none focus:border-primary"
              placeholder="Search users..."
              type="text"
            />
          </div>
          <button
            class="text-primary hover:text-primary-hover text-sm font-medium flex items-center gap-1"
          >
            <span class="material-symbols-outlined text-[18px]"> person_add </span>
            Assign User
          </button>
        </div>
        <div class="space-y-3">
          <!-- User Item -->
          <div
            class="flex items-center justify-between p-3 rounded-lg border border-border-default hover:bg-surface-subtle transition-colors"
          >
            <div class="flex items-center gap-3">
              <img
                alt="Profile"
                class="w-8 h-8 rounded-full bg-stone-200"
                data-alt="Corporate headshot of a young professional woman smiling lightly, clean white background"
                src="https://lh3.googleusercontent.com/aida-public/AB6AXuBmhKDZK--iT2dE9l_2_UoPavDyYhL3qieLmskVM9BN4Cgge0yzhYlqs_mUSf4jg4IXsUz0jeNqvpYd-HKgL06JCNEXHo-OPIOku6d3U2KFBDaGxJFhxlNX7cvhqr5htoOA9__vMm783ZLsDncyZqSdl4aj-nokEs-u3Gdy7mKb9UNXYvoeRgH4ScXzC-zQyHoKE4T-gEldE7imFCeN1jEHhE11RM1mEkWV-5XxgI-e3_XDf1wPKhiLdJQhLNPv_G8vfcJEENeaP0ux"
              />
              <div>
                <div class="text-sm font-medium text-text-primary">Sarah Jenkins</div>
                <div class="text-xs text-text-secondary">UI/UX Designer • Marketing Team</div>
              </div>
            </div>
            <div class="flex items-center gap-3">
              <span class="text-xs text-stone-500 font-mono-data"> Last Active: 2h ago </span>
              <button
                class="text-stone-400 hover:text-error transition-colors"
                title="Revoke License"
              >
                <span class="material-symbols-outlined text-[18px]"> person_remove </span>
              </button>
            </div>
          </div>
          <!-- User Item -->
          <div
            class="flex items-center justify-between p-3 rounded-lg border border-border-default hover:bg-surface-subtle transition-colors"
          >
            <div class="flex items-center gap-3">
              <img
                alt="Profile"
                class="w-8 h-8 rounded-full bg-stone-200"
                data-alt="Corporate headshot of a middle aged man with glasses, professional lighting"
                src="https://lh3.googleusercontent.com/aida-public/AB6AXuCqqaiXZQ-aYrBIl1F2NwbwmZxRS0YiGTmfIboAO-icsvTX5TcNy6Ao7vc5wn3BAPLp15rZ1CRSmQYPF4WOX91knqqzRqhsX-JTGN-nX9B0KPkLIUbuW6Q8p0vdoHwB4dkB4Op3oE5NMA42d3DDNhlBFKUqYoUDG6U2lPWLzaezWPu8A29PlM7l7Mq7hoPDN5-dgJTvgvzxfUZdfx7LJx3z4bjPczjUGAP_Z-tOuVwjfHoYJzjCETMnkTQwmeIJ0B_k3QLNvUumdxKN"
              />
              <div>
                <div class="text-sm font-medium text-text-primary">David Chen</div>
                <div class="text-xs text-text-secondary">Creative Director • Design Studio</div>
              </div>
            </div>
            <div class="flex items-center gap-3">
              <span class="text-xs text-stone-500 font-mono-data"> Last Active: 1d ago </span>
              <button
                class="text-stone-400 hover:text-error transition-colors"
                title="Revoke License"
              >
                <span class="material-symbols-outlined text-[18px]"> person_remove </span>
              </button>
            </div>
          </div>
        </div>
      </div>
      <!-- Drawer Footer Action -->
      <div class="p-6 border-t border-border-default bg-surface-subtle rounded-bl-2xl">
        <button
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
import { getSoftwareLicenses } from '../api/software-licenses';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';

interface SoftwareLicense {
  id: number;
  softwareName: string;
  planName: string;
  status: string;
  totalSeats: number;
  usedSeats: number;
  renewalDate: string;
  annualCost: number;
}

const licenses = ref<SoftwareLicense[]>([]);
const loading = ref(true);
const selectedLicense = ref<SoftwareLicense | null>(null);

const expiringCount = computed(() => {
  return licenses.value.filter((l) => l.status === 'Expiring Soon').length;
});

const underutilizedSeats = computed(() => {
  return licenses.value.reduce((sum, l) => sum + (l.totalSeats - l.usedSeats), 0);
});

const annualSpend = computed(() => {
  return licenses.value.reduce((sum, l) => sum + l.annualCost, 0);
});

const fetchLicenses = async () => {
  try {
    licenses.value = (await getSoftwareLicenses()) as any[];
  } catch (error) {
    console.error('Failed to fetch licenses:', error);
  } finally {
    loading.value = false;
  }
};

onMounted(() => {
  fetchLicenses();
});
</script>
