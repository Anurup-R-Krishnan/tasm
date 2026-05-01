<template>
  <main class="flex-1 md:pl-[248px] p-page-margin">
    <div class="max-w-[1400px] mx-auto space-y-section-gap">
      <!-- Top Bento Row: Welcome & Quick Action -->
      <div class="grid grid-cols-1 md:grid-cols-12 gap-section-gap">
        <!-- Welcome Banner (Span 8) -->
        <div
          class="md:col-span-8 bg-surface rounded-xl border border-border-default shadow-[0_4px_12px_rgba(0,0,0,0.02)] p-card-padding flex flex-col justify-center relative overflow-hidden"
        >
          <div
            class="absolute right-0 top-0 w-64 h-full bg-gradient-to-l from-metric-amber/30 to-transparent pointer-events-none"
          ></div>
          <h1 class="font-h1 text-h1 text-text-primary mb-2 z-10">Hello, Divya.</h1>
          <p class="font-body text-body text-text-secondary z-10 mb-6 max-w-lg">
            Welcome back to the Employee Self-Service Portal. You have critical actions pending
            regarding your assigned equipment.
          </p>
          <div
            class="bg-surface-subtle border border-amber-200 rounded-lg p-4 flex items-start gap-4 inline-flex w-fit z-10 shadow-sm"
          >
            <div
              class="w-8 h-8 rounded-full bg-metric-amber flex items-center justify-center shrink-0 mt-0.5"
            >
              <span class="material-symbols-outlined text-amber-800 text-[18px]"> warning </span>
            </div>
            <div>
              <h3 class="font-h3 text-h3 text-text-primary mb-1">Action Required</h3>
              <p class="font-body text-body text-text-secondary">
                You have
                <span class="font-semibold text-text-primary"> 2 assets </span>
                awaiting acknowledgement.
              </p>
            </div>
          </div>
        </div>
        <!-- Hero CTA (Span 4) -->
        <div
          class="md:col-span-4 bg-primary text-on-primary rounded-xl p-card-padding flex flex-col justify-between shadow-md relative overflow-hidden group cursor-pointer hover:-translate-y-1 transition-transform duration-300"
        >
          <div
            class="absolute -right-12 -top-12 w-40 h-40 bg-white/10 rounded-full blur-2xl group-hover:scale-110 transition-transform duration-500"
          ></div>
          <div class="z-10">
            <div class="w-12 h-12 rounded-full bg-white/20 flex items-center justify-center mb-6">
              <span class="material-symbols-outlined text-white text-[24px]"> devices </span>
            </div>
            <h2 class="font-h2 text-h2 text-white mb-2">Need equipment?</h2>
            <p class="font-body text-body text-white/80 mb-6">
              Submit a request for new laptops, peripherals, or specialized software licenses.
            </p>
          </div>
          <button
            class="w-full bg-metric-amber text-amber-900 rounded-lg py-3 font-semibold hover:bg-amber-200 transition-colors z-10 flex items-center justify-center gap-2"
          >
            Request New Asset
            <span class="material-symbols-outlined text-[18px]"> arrow_forward </span>
          </button>
        </div>
      </div>
      <!-- My Assets Table -->
      <section>
        <div class="flex items-center justify-between mb-4">
          <h2 class="font-h2 text-h2 text-text-primary">My Assigned Assets</h2>
          <button class="text-primary font-medium text-sm hover:underline">View All History</button>
        </div>
        <div
          class="bg-surface border border-border-default rounded-xl shadow-[0_4px_12px_rgba(0,0,0,0.02)] overflow-hidden p-4"
        >
          <DataTable :value="myAssets" :loading="loading" class="w-full text-left">
            <Column field="name" header="Asset Name" sortable>
              <template #body="slotProps">
                <div class="flex items-center gap-3">
                  <div
                    class="w-10 h-10 rounded bg-stone-100 flex items-center justify-center text-stone-500 shrink-0"
                  >
                    <span class="material-symbols-outlined">
                      {{
                        slotProps.data.category === 'Laptops'
                          ? 'laptop_mac'
                          : slotProps.data.category === 'Mobile Devices'
                            ? 'smartphone'
                            : 'devices'
                      }}
                    </span>
                  </div>
                  <div class="font-medium">
                    {{ slotProps.data.name }}
                  </div>
                </div>
              </template>
            </Column>
            <Column field="tagId" header="Tag ID" sortable>
              <template #body="slotProps">
                <span class="font-mono-data text-mono-data text-text-secondary">{{
                  slotProps.data.tagId
                }}</span>
              </template>
            </Column>
            <Column field="purchaseDate" header="Checked Out" sortable>
              <template #body="slotProps">
                <span>{{ new Date(slotProps.data.purchaseDate).toLocaleDateString() }}</span>
              </template>
            </Column>
            <Column field="status" header="Status" sortable>
              <template #body="slotProps">
                <span
                  class="inline-flex items-center px-2.5 py-1 rounded-full text-xs font-medium"
                  :class="
                    slotProps.data.status === 'Active'
                      ? 'bg-metric-sage text-green-900'
                      : 'bg-status-checked-out text-white'
                  "
                >
                  {{ slotProps.data.status }}
                </span>
              </template>
            </Column>
            <Column header="Action" alignFrozen="right">
              <template #body>
                <button
                  class="text-sm font-medium text-primary hover:bg-primary/10 px-3 py-1.5 rounded transition-colors"
                >
                  Report Issue
                </button>
              </template>
            </Column>
          </DataTable>
        </div>
      </section>
      <!-- Bottom Row: Reservations & Requests -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-section-gap">
        <!-- My Reservations -->
        <div
          class="bg-surface border border-border-default rounded-xl shadow-sm p-card-padding flex flex-col h-full"
        >
          <div class="flex items-center justify-between mb-6">
            <h2 class="font-h2 text-h2 text-text-primary flex items-center gap-2">
              <span class="material-symbols-outlined text-stone-400"> calendar_month </span>
              Upcoming Reservations
            </h2>
            <button class="text-stone-400 hover:text-stone-600">
              <span class="material-symbols-outlined"> more_horiz </span>
            </button>
          </div>
          <div class="space-y-4 flex-1">
            <!-- Reservation Item -->
            <div
              class="p-4 rounded-lg border border-border-default bg-surface-subtle flex items-start gap-4"
            >
              <div
                class="w-12 h-12 rounded bg-white border border-stone-200 flex flex-col items-center justify-center shrink-0"
              >
                <span class="text-xs font-semibold text-text-secondary uppercase"> Oct </span>
                <span class="text-lg font-bold text-text-primary leading-none"> 18 </span>
              </div>
              <div class="flex-1">
                <h3 class="font-h3 text-h3 text-text-primary">Conference Room Projector</h3>
                <p class="font-body text-sm text-text-secondary mt-0.5">09:00 AM - 12:00 PM</p>
              </div>
              <span
                class="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium bg-metric-sage text-green-800"
              >
                Confirmed
              </span>
            </div>
            <!-- Reservation Item -->
            <div
              class="p-4 rounded-lg border border-border-default bg-surface flex items-start gap-4 opacity-70"
            >
              <div
                class="w-12 h-12 rounded bg-stone-50 border border-stone-200 flex flex-col items-center justify-center shrink-0"
              >
                <span class="text-xs font-semibold text-text-secondary uppercase"> Oct </span>
                <span class="text-lg font-bold text-text-primary leading-none"> 22 </span>
              </div>
              <div class="flex-1">
                <h3 class="font-h3 text-h3 text-text-primary">DSLR Camera Kit</h3>
                <p class="font-body text-sm text-text-secondary mt-0.5">Full Day • Offsite Event</p>
              </div>
              <span
                class="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium bg-stone-100 text-stone-600"
              >
                Pending Approval
              </span>
            </div>
          </div>
          <button
            class="mt-4 w-full py-2 border border-border-default rounded-lg text-sm font-medium text-text-secondary hover:bg-stone-50 transition-colors"
          >
            Book Equipment
          </button>
        </div>
        <!-- My Requests -->
        <div
          class="bg-surface border border-border-default rounded-xl shadow-sm p-card-padding flex flex-col h-full"
        >
          <div class="flex items-center justify-between mb-6">
            <h2 class="font-h2 text-h2 text-text-primary flex items-center gap-2">
              <span class="material-symbols-outlined text-stone-400"> receipt_long </span>
              Open Requests
            </h2>
            <button class="text-stone-400 hover:text-stone-600">
              <span class="material-symbols-outlined"> filter_list </span>
            </button>
          </div>
          <div class="space-y-0 divide-y divide-border-default flex-1">
            <!-- Request Item -->
            <div class="py-4 first:pt-0 flex items-center justify-between group">
              <div class="flex items-center gap-4">
                <div
                  class="w-10 h-10 rounded-full bg-blue-50 text-blue-600 flex items-center justify-center shrink-0"
                >
                  <span class="material-symbols-outlined text-[20px]"> build </span>
                </div>
                <div>
                  <div
                    class="font-medium text-text-primary group-hover:text-primary transition-colors"
                  >
                    Keyboard Replacement
                  </div>
                  <div class="flex items-center gap-2 mt-0.5">
                    <span class="font-mono-data text-xs text-text-secondary"> REQ-1092 </span>
                    <span class="w-1 h-1 rounded-full bg-stone-300"> </span>
                    <span class="text-xs text-text-secondary"> Filed Oct 10 </span>
                  </div>
                </div>
              </div>
              <span
                class="inline-flex items-center px-2.5 py-1 rounded-full text-xs font-medium bg-blue-100 text-blue-800"
              >
                In Progress
              </span>
            </div>
            <!-- Request Item -->
            <div class="py-4 flex items-center justify-between group">
              <div class="flex items-center gap-4">
                <div
                  class="w-10 h-10 rounded-full bg-purple-50 text-purple-600 flex items-center justify-center shrink-0"
                >
                  <span class="material-symbols-outlined text-[20px]"> shopping_bag </span>
                </div>
                <div>
                  <div
                    class="font-medium text-text-primary group-hover:text-primary transition-colors"
                  >
                    Adobe CC License
                  </div>
                  <div class="flex items-center gap-2 mt-0.5">
                    <span class="font-mono-data text-xs text-text-secondary"> REQ-1088 </span>
                    <span class="w-1 h-1 rounded-full bg-stone-300"> </span>
                    <span class="text-xs text-text-secondary"> Cost: ₹4,500/mo </span>
                  </div>
                </div>
              </div>
              <span
                class="inline-flex items-center px-2.5 py-1 rounded-full text-xs font-medium bg-metric-amber text-amber-900"
              >
                Manager Approval
              </span>
            </div>
            <!-- Request Item -->
            <div class="py-4 pb-0 flex items-center justify-between group opacity-60">
              <div class="flex items-center gap-4">
                <div
                  class="w-10 h-10 rounded-full bg-green-50 text-green-600 flex items-center justify-center shrink-0"
                >
                  <span class="material-symbols-outlined text-[20px]"> check_circle </span>
                </div>
                <div>
                  <div class="font-medium text-text-primary">Ergonomic Chair</div>
                  <div class="flex items-center gap-2 mt-0.5">
                    <span class="font-mono-data text-xs text-text-secondary"> REQ-0941 </span>
                    <span class="w-1 h-1 rounded-full bg-stone-300"> </span>
                    <span class="text-xs text-text-secondary"> Closed Oct 02 </span>
                  </div>
                </div>
              </div>
              <span class="text-xs font-medium text-text-secondary"> Completed </span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';

const myAssets = ref<any[]>([]);
const loading = ref(true);

const fetchMyAssets = async () => {
  try {
    const res = await fetch('http://localhost:8080/api/assets');
    if (res.ok) {
      // In a real app we would filter by logged in user, mock for now
      const allAssets = await res.json();
      myAssets.value = allAssets.slice(0, 3);
    }
  } catch (error) {
    console.error('Failed to fetch assets:', error);
  } finally {
    loading.value = false;
  }
};

onMounted(() => {
  fetchMyAssets();
});
</script>
