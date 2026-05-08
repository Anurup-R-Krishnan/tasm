<template>
  <main class="space-y-section-gap pb-24">
    <!-- Page Title -->
    <div class="flex flex-col gap-1">
      <h1 class="font-h1 text-h1 text-text-primary">Operations Overview</h1>
      <p class="font-body text-text-secondary">Manage maintenance queues and vendor contracts.</p>
    </div>
    <!-- Header KPIs (4 Cards) -->
    <section class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-inline">
      <!-- KPI 1: Amber -->
      <div
        class="bg-metric-amber rounded-xl p-card-padding flex flex-col gap-base border border-border-default shadow-sm hover:-translate-y-0.5 transition-transform duration-200 cursor-pointer"
      >
        <div class="flex justify-between items-start mb-2">
          <span class="font-metadata text-metadata text-text-secondary uppercase tracking-wider">
            Active Work Orders
          </span>
          <span class="material-symbols-outlined text-outline"> build </span>
        </div>
        <div class="font-kpi-number text-kpi-number text-text-primary">42</div>
        <div class="font-metadata text-metadata text-text-secondary mt-1 flex items-center gap-1">
          <span class="material-symbols-outlined text-[14px] text-status-critical">
            arrow_upward
          </span>
          <span> 12 critical </span>
        </div>
      </div>
      <!-- KPI 2: Peach (Using error-container for semantic overdue feel) -->
      <div
        class="bg-error-container rounded-xl p-card-padding flex flex-col gap-base border border-border-default shadow-sm hover:-translate-y-0.5 transition-transform duration-200 cursor-pointer"
      >
        <div class="flex justify-between items-start mb-2">
          <span class="font-metadata text-metadata text-text-secondary uppercase tracking-wider">
            Overdue Maintenance
          </span>
          <span class="material-symbols-outlined text-outline"> warning </span>
        </div>
        <div class="font-kpi-number text-kpi-number text-text-primary">18</div>
        <div class="font-metadata text-metadata text-text-secondary mt-1 flex items-center gap-1">
          <span class="text-on-error-container font-medium"> +3 since yesterday </span>
        </div>
      </div>
      <!-- KPI 3: Lavender -->
      <div
        class="bg-metric-lavender rounded-xl p-card-padding flex flex-col gap-base border border-border-default shadow-sm hover:-translate-y-0.5 transition-transform duration-200 cursor-pointer"
      >
        <div class="flex justify-between items-start mb-2">
          <span class="font-metadata text-metadata text-text-secondary uppercase tracking-wider">
            Expiring Contracts
          </span>
          <span class="material-symbols-outlined text-outline"> contract </span>
        </div>
        <div class="font-kpi-number text-kpi-number text-text-primary">7</div>
        <div class="font-metadata text-metadata text-text-secondary mt-1 flex items-center gap-1">
          <span> Next 30 days </span>
        </div>
      </div>
      <!-- KPI 4: Sage -->
      <div
        class="bg-metric-sage rounded-xl p-card-padding flex flex-col gap-base border border-border-default shadow-sm hover:-translate-y-0.5 transition-transform duration-200 cursor-pointer"
      >
        <div class="flex justify-between items-start mb-2">
          <span class="font-metadata text-metadata text-text-secondary uppercase tracking-wider">
            Total Vendor Spend
          </span>
          <span class="material-symbols-outlined text-outline"> payments </span>
        </div>
        <div class="font-kpi-number text-kpi-number text-text-primary">$1.2M</div>
        <div class="font-metadata text-metadata text-text-secondary mt-1 flex items-center gap-1">
          <span class="material-symbols-outlined text-[14px] text-status-in-stock">
            arrow_downward
          </span>
          <span> 4% vs last quarter </span>
        </div>
      </div>
    </section>
    <!-- 60/40 Split Layout -->
    <section class="grid grid-cols-1 lg:grid-cols-12 gap-inline items-start">
      <!-- Left Column (60%) - Work Order Queue -->
      <div
        class="lg:col-span-8 bg-surface rounded-xl border border-border-default shadow-sm overflow-hidden flex flex-col"
      >
        <div
          class="p-card-padding border-b border-border-default flex items-center justify-between bg-surface-subtle"
        >
          <h2 class="font-h2 text-h2 text-text-primary">Work Order Queue</h2>
          <button
            @click="router.push('/maintenance')"
            class="bg-text-primary text-on-primary font-body px-4 py-2 rounded-lg hover:bg-secondary transition-colors text-sm"
          >
            New Order
          </button>
        </div>
        <div class="overflow-x-auto border-t border-border-default">
          <DataTable
            :value="workOrders"
            :loading="loadingWorkOrders"
            paginator
            :rows="5"
            tableStyle="min-width: 40rem"
            class="w-full"
          >
            <Column field="workOrderId" header="ID" sortable>
              <template #body="slotProps">
                <span class="font-mono-data text-mono-data text-text-secondary">{{
                  slotProps.data.workOrderId
                }}</span>
              </template>
            </Column>
            <Column header="Asset / Location" sortable field="title">
              <template #body="slotProps">
                <div class="font-medium text-text-primary">{{ slotProps.data.title }}</div>
                <div class="font-metadata text-metadata text-text-secondary mt-0.5">
                  {{ slotProps.data.assetLocation }}
                </div>
              </template>
            </Column>
            <Column field="severity" header="Severity" sortable>
              <template #body="slotProps">
                <span
                  v-if="slotProps.data.severity === 'Critical'"
                  class="inline-flex items-center px-2 py-0.5 rounded-full font-metadata text-metadata bg-status-critical text-on-error"
                  >Critical</span
                >
                <span
                  v-else-if="slotProps.data.severity === 'High'"
                  class="inline-flex items-center px-2 py-0.5 rounded-full font-metadata text-metadata bg-primary-container text-on-primary-container"
                  >High</span
                >
                <span
                  v-else-if="slotProps.data.severity === 'Medium'"
                  class="inline-flex items-center px-2 py-0.5 rounded-full font-metadata text-metadata bg-surface-variant text-on-surface-variant"
                  >Medium</span
                >
                <span
                  v-else
                  class="inline-flex items-center px-2 py-0.5 rounded-full font-metadata text-metadata bg-secondary-container text-on-secondary-container border border-border-default"
                  >Low</span
                >
              </template>
            </Column>
            <Column field="createdAt" header="Reported" sortable>
              <template #body="slotProps">
                <span class="text-text-secondary">{{
                  new Date(slotProps.data.createdAt).toLocaleDateString()
                }}</span>
              </template>
            </Column>
            <Column header="Actions" alignFrozen="right">
              <template #body="slotProps">
                <button
                  @click="deleteWorkOrder(slotProps.data.id)"
                  class="p-1 text-text-secondary hover:text-status-critical rounded transition-colors"
                >
                  <span class="material-symbols-outlined text-[20px]">delete</span>
                </button>
              </template>
            </Column>
          </DataTable>
        </div>
        <div class="p-4 border-t border-border-default bg-surface text-center">
          <a
            class="font-metadata text-metadata text-tertiary hover:underline"
            href="/maintenance-tracker"
            @click.prevent="router.push('/maintenance-tracker')"
          >
            View All Orders
          </a>
        </div>
      </div>
      <!-- Right Column (40%) - Contract Expiry List -->
      <div
        class="lg:col-span-4 bg-surface rounded-xl border border-border-default shadow-sm flex flex-col"
      >
        <div
          class="p-card-padding border-b border-border-default flex items-center justify-between"
        >
          <h2 class="font-h2 text-h2 text-text-primary">Upcoming Renewals</h2>
          <button
            class="text-text-secondary hover:text-text-primary p-1 rounded-md hover:bg-surface-subtle transition-colors"
          >
            <span class="material-symbols-outlined text-[20px]"> filter_list </span>
          </button>
        </div>
        <div class="flex flex-col p-4 gap-2">
          <!-- Dynamic List Item -->
          <div
            v-for="contract in contracts"
            :key="contract.id"
            class="flex items-start gap-4 p-3 rounded-lg hover:bg-surface-subtle transition-colors cursor-pointer border border-transparent hover:border-border-default"
          >
            <div
              class="w-10 h-10 rounded-full bg-surface-variant flex items-center justify-center shrink-0"
            >
              <span class="material-symbols-outlined text-on-surface-variant text-[20px]">
                assignment
              </span>
            </div>
            <div class="flex-1 min-w-0">
              <h3 class="font-h3 text-h3 text-text-primary truncate">
                {{ contract.vendorName }}
              </h3>
              <p class="font-metadata text-metadata text-text-secondary truncate mt-0.5">
                {{ contract.serviceType }}
              </p>
            </div>
            <div class="text-right shrink-0">
              <div
                class="font-h3 text-h3"
                :class="
                  getDaysRemaining(contract.endDate) <= 7
                    ? 'text-status-critical'
                    : 'text-text-primary'
                "
              >
                {{ getDaysRemaining(contract.endDate) }} days
              </div>
              <div class="font-metadata text-metadata text-text-secondary">
                {{ new Date(contract.endDate).toLocaleDateString() }}
              </div>
            </div>
          </div>

          <div
            v-if="contracts.length === 0 && !loadingContracts"
            class="text-center py-4 text-text-secondary text-sm"
          >
            No upcoming renewals.
          </div>
        </div>
        <div class="p-4 mt-auto border-t border-border-default text-center">
          <button
            @click="router.push('/maintenance')"
            class="font-metadata text-metadata text-tertiary hover:underline font-medium"
          >
            Manage Contracts
          </button>
        </div>
      </div>
    </section>
  </main>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { getWorkOrders } from '../api/workOrders';
import { getContracts } from '../api/financial';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';

const router = useRouter();

interface WorkOrder {
  id: number;
  workOrderId: string;
  title: string;
  assetLocation: string;
  status: string;
  severity: string;
  createdAt: string;
}

interface MaintenanceContract {
  id: number;
  vendorName: string;
  serviceType: string;
  endDate: string;
}

const workOrders = ref<WorkOrder[]>([]);
const contracts = ref<MaintenanceContract[]>([]);
const loadingWorkOrders = ref(true);
const loadingContracts = ref(true);

const fetchWorkOrders = async () => {
  try {
    workOrders.value = (await getWorkOrders()) as any[];
  } catch (err) {
    console.error('Failed to fetch work orders:', err);
  } finally {
    loadingWorkOrders.value = false;
  }
};

const deleteWorkOrder = async (id: number) => {
  if (!confirm('Are you sure you want to delete this work order?')) return;
  try {
    const { deleteWorkOrder: apiDeleteWorkOrder } = await import('../api/workOrders');
    await apiDeleteWorkOrder(id);
    workOrders.value = workOrders.value.filter((w) => w.id !== id);
  } catch (error) {
    console.error('Error deleting work order:', error);
  }
};

const fetchContracts = async () => {
  try {
    contracts.value = await getContracts();
  } catch (err) {
    console.error('Failed to fetch contracts:', err);
  } finally {
    loadingContracts.value = false;
  }
};

const getDaysRemaining = (endDateStr: string) => {
  const endDate = new Date(endDateStr);
  const today = new Date();
  const diffTime = endDate.getTime() - today.getTime();
  return Math.ceil(diffTime / (1000 * 60 * 60 * 24));
};

onMounted(() => {
  fetchWorkOrders();
  fetchContracts();
});
</script>
