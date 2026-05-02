<template>
  <main class="space-y-section-gap pb-24">
    <!-- Page Header -->
    <div class="flex flex-col md:flex-row md:items-end justify-between gap-4">
      <div>
        <div class="flex items-center gap-2 text-text-secondary font-metadata text-metadata mb-2">
          <button
            @click="router.push('/work-order-detail')"
            class="hover:text-primary transition-colors flex items-center"
          >
            Work Orders
          </button>
          <span class="material-symbols-outlined text-[14px]"> chevron_right </span>
          <span v-if="order" class="font-mono-data text-mono-data text-text-primary">
            {{ order.id }}
          </span>
        </div>
        <div class="flex items-center gap-4">
          <h1 v-if="order" class="font-h1 text-h1 text-text-primary">{{ order.title }}</h1>
          <div v-if="order" class="flex items-center gap-2">
            <span
              class="bg-status-critical/20 text-status-critical px-2.5 py-1 rounded-full font-metadata text-metadata flex items-center gap-1 border border-status-critical/30"
            >
              <span class="material-symbols-outlined text-[12px] icon-fill"> error </span>
              {{ order.severity }}
            </span>
            <span
              class="bg-status-checked-out/20 text-status-checked-out px-2.5 py-1 rounded-full font-metadata text-metadata flex items-center gap-1 border border-status-checked-out/30"
            >
              <span class="material-symbols-outlined text-[12px]"> cycle </span>
              {{ order.status }}
            </span>
          </div>
        </div>
      </div>
      <div class="flex items-center gap-inline">
        <button
          class="px-4 py-2 bg-surface text-text-primary border border-border-default rounded-lg font-body text-body hover:bg-surface-subtle transition-colors shadow-sm flex items-center gap-2 hover:-translate-y-[1px] duration-200"
        >
          <span class="material-symbols-outlined text-[18px]"> print </span>
          Print
        </button>
        <button
          class="px-4 py-2 bg-primary text-on-primary rounded-lg font-body text-body hover:bg-primary-hover transition-colors shadow-sm flex items-center gap-2 hover:-translate-y-[1px] duration-200"
        >
          <span class="material-symbols-outlined text-[18px]"> edit </span>
          Edit Order
        </button>
      </div>
    </div>
    <!-- Bento Grid Layout -->
    <div class="grid grid-cols-1 lg:grid-cols-12 gap-stack items-start">
      <!-- Left Column (Main Content - 8 cols) -->
      <div class="lg:col-span-8 space-y-stack">
        <!-- Asset Header Card (Glassmorphism inspired, structural) -->
        <div
          class="bg-surface border border-border-default rounded-xl shadow-sm overflow-hidden hover:-translate-y-[2px] transition-transform duration-300"
        >
          <div class="p-card-padding flex flex-col md:flex-row gap-6 items-start">
            <div
              class="w-24 h-24 rounded-lg bg-surface-variant flex-shrink-0 border border-border-default overflow-hidden relative"
            >
              <img
                class="w-full h-full object-cover"
                data-alt="Industrial HVAC compressor unit in a commercial building mechanical room, well lit, metallic textures"
                src="https://lh3.googleusercontent.com/aida-public/AB6AXuClVqu29Efx7pXUteE-rtjp4U81zsD8LIA1CkRMeLzzxna94mh3ZMp6XlBnhPZEmHX8bVYSy1bdYK71QSntly6Ux3aTSYu0pEj_3jtdLloyRlicSLkN-NzZWffj5UzkSumcqf7joSmINXzgPAAl34sKZAWsqz53NM_e-izNd7dKmPxBs_LDyiwzkuWeK7BypFSyexHe14_aNqSW5EObdBvQbLhiGkNYpBAZ53pfeLwBYDkKj_A_RuS_dnlis0MOXcmCb2HjJiPJay7-"
              />
            </div>
            <div class="flex-1">
              <h3
                class="font-metadata text-metadata text-text-secondary uppercase tracking-wider mb-1"
              >
                Target Asset
              </h3>
              <div class="flex items-center gap-2 mb-2">
                <h2 v-if="order" class="font-h2 text-h2 text-text-primary">
                  {{ order.assetName }}
                </h2>
                <span
                  v-if="order"
                  class="font-mono-data text-mono-data bg-surface-subtle px-2 py-0.5 rounded border border-border-default text-text-secondary"
                >
                  {{ order.assetTag }}
                </span>
              </div>
              <p v-if="order" class="font-body text-body text-text-secondary max-w-2xl">
                {{ order.description }}
              </p>
              <div class="mt-4 flex gap-6 border-t border-border-default pt-4">
                <div>
                  <p class="font-metadata text-metadata text-text-secondary">Location</p>
                  <p v-if="order" class="font-body text-body font-medium">
                    {{ order.assetLocation }}
                  </p>
                </div>
                <div>
                  <p class="font-metadata text-metadata text-text-secondary">Reported By</p>
                  <p v-if="order" class="font-body text-body font-medium">{{ order.reportedBy }}</p>
                </div>
                <div>
                  <p class="font-metadata text-metadata text-text-secondary">Date Created</p>
                  <p v-if="order" class="font-body text-body font-medium">
                    {{ formatDate(order.createdAt) }}
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
        <!-- Work Log Thread -->
        <div
          class="bg-surface border border-border-default rounded-xl shadow-sm p-card-padding hover:-translate-y-[2px] transition-transform duration-300"
        >
          <h3
            class="font-h3 text-h3 text-text-primary border-b border-border-default pb-3 mb-6 flex items-center gap-2"
          >
            <span class="material-symbols-outlined text-[20px] text-text-secondary"> history </span>
            Activity Log
          </h3>
          <div class="relative pl-4 border-l-2 border-surface-variant space-y-8">
            <!-- Entry 1 -->
            <div class="relative">
              <div
                class="absolute -left-[23px] top-1 w-4 h-4 rounded-full bg-canvas border-2 border-primary-container z-10"
              ></div>
              <div class="flex items-center justify-between mb-1">
                <p class="font-h3 text-h3 text-text-primary">Parts Arrived on Site</p>
                <span class="font-metadata text-metadata text-text-secondary">
                  Today, 09:15 AM
                </span>
              </div>
              <p class="font-body text-body text-text-secondary">
                Replacement compressor and refrigerant tanks have been delivered to the loading
                dock. Moving to roof deck.
              </p>
              <div class="mt-2 flex items-center gap-2">
                <div
                  class="w-6 h-6 rounded-full bg-secondary-container flex items-center justify-center text-xs font-bold text-on-secondary-container"
                >
                  MR
                </div>
                <span class="font-metadata text-metadata text-text-primary">
                  Michael Reed (Technician)
                </span>
              </div>
            </div>
            <!-- Entry 2 -->
            <div class="relative">
              <div
                class="absolute -left-[23px] top-1 w-4 h-4 rounded-full bg-surface-variant border-2 border-outline-variant z-10"
              ></div>
              <div class="flex items-center justify-between mb-1">
                <p class="font-h3 text-h3 text-text-primary">Vendor Purchase Order Approved</p>
                <span class="font-metadata text-metadata text-text-secondary">
                  Yesterday, 14:30 PM
                </span>
              </div>
              <p class="font-body text-body text-text-secondary">
                PO-4921 generated and sent to Trane Commercial Parts. Expedited shipping requested
                due to critical severity.
              </p>
              <div class="mt-2 flex items-center gap-2">
                <div
                  class="w-6 h-6 rounded-full bg-primary-container flex items-center justify-center text-xs font-bold text-on-primary"
                >
                  AK
                </div>
                <span class="font-metadata text-metadata text-text-primary"> Admin (System) </span>
              </div>
            </div>
            <!-- Entry 3 -->
            <div class="relative">
              <div
                class="absolute -left-[23px] top-1 w-4 h-4 rounded-full bg-surface-variant border-2 border-outline-variant z-10"
              ></div>
              <div class="flex items-center justify-between mb-1">
                <p class="font-h3 text-h3 text-text-primary">Initial Inspection Completed</p>
                <span class="font-metadata text-metadata text-text-secondary">
                  Oct 12, 11:00 AM
                </span>
              </div>
              <p class="font-body text-body text-text-secondary">
                Confirmed bearing failure. The shaft is seized. Requires full compressor
                replacement. Tagged out unit at main breaker.
              </p>
              <!-- Photo thumbnails -->
              <div class="flex gap-2 mt-3">
                <img
                  class="w-16 h-16 rounded border border-border-default object-cover cursor-pointer hover:opacity-80"
                  data-alt="Close up of a burnt out mechanical bearing covered in dark grease and metal shavings"
                  src="https://lh3.googleusercontent.com/aida-public/AB6AXuBcwwHAudoxguJh9FRyxeDl6lOhL-tMY-QjBQ75PbKaCrFd-wllUmDM9Qpjdye9w9wmtQssKKqXi96PtHTueEd8n8Bw9jne5pcnbMv8Ck_eCx_H_nc9DGZRXmMWC1blzl25O3ZKlt2T7LdCaUkNVLIlJdc47EGLaGmLtMjuxNG3qcNFrnVQsNv-A92HFRmpT47r2D6RS-ghEos8_X7qgDK9S1p44Ain3_ObQr2GIaTTOTu4INXup8lL-3j6IW4A1Y8pdqsvXCIhtSIN"
                />
                <img
                  class="w-16 h-16 rounded border border-border-default object-cover cursor-pointer hover:opacity-80"
                  data-alt="Electrical lockout tag on an industrial breaker box"
                  src="https://lh3.googleusercontent.com/aida-public/AB6AXuD2iph5IY55sqFsF-82YrThEJQOFckfEbSb1ZSY9F8t75VLDypjJlO169jQqXmAcQizSMpoZxBowIuYiPsvQUAyQi_DImCcVTUaeI9MeyRcf-jHfD2M_aZxYue0jFYzpIKNxVzOSIu_6Ao7zoRbpzPPVFHDLAxtttMfb4OMzMjhb-3x3j5q8qt4QoM1hk3tl1hDwSDBXZGieZkp6VTaChtanfAeNnUom0T6RrrYYvMXprxyuPGgUl_QshuRxcrkBxdaHZ7vbS4jI_6a"
                />
              </div>
            </div>
          </div>
          <div class="mt-8 relative">
            <textarea
              class="w-full bg-surface-subtle border border-border-default rounded-lg p-3 font-body text-body focus:ring-primary-container focus:border-primary-container resize-none"
              placeholder="Add an update to the log..."
              rows="3"
            ></textarea>
            <div class="absolute bottom-3 right-3 flex gap-2">
              <button
                class="p-1.5 text-text-secondary hover:text-primary transition-colors rounded"
              >
                <span class="material-symbols-outlined text-[20px]"> attach_file </span>
              </button>
              <button
                class="px-3 py-1.5 bg-primary text-on-primary rounded font-body text-body text-sm hover:bg-primary/90 transition-colors"
              >
                Post
              </button>
            </div>
          </div>
        </div>
      </div>
      <!-- Right Column (Sidebar metrics/actions - 4 cols) -->
      <div class="lg:col-span-4 space-y-stack">
        <!-- Assignment Card -->
        <div
          class="bg-surface border border-border-default rounded-xl shadow-sm p-card-padding hover:-translate-y-[2px] transition-transform duration-300"
        >
          <h3 class="font-h3 text-h3 text-text-primary border-b border-border-default pb-3 mb-4">
            Assignment
          </h3>
          <div class="space-y-4">
            <div class="flex items-center gap-3">
              <div
                class="w-10 h-10 rounded-full bg-secondary-container flex items-center justify-center font-h3 text-on-secondary-container"
              >
                MR
              </div>
              <div>
                <p v-if="order" class="font-h3 text-h3 text-text-primary">{{ order.technician }}</p>
                <p class="font-metadata text-metadata text-text-secondary">Assigned Specialist</p>
              </div>
            </div>
            <div class="bg-surface-subtle p-3 rounded-lg border border-border-default">
              <p class="font-metadata text-metadata text-text-secondary mb-1">Assigned Vendor</p>
              <div class="flex items-center gap-2">
                <span class="material-symbols-outlined text-[18px] text-text-secondary">
                  store
                </span>
                <p class="font-body text-body font-medium text-text-primary">
                  Trane Commercial Service
                </p>
              </div>
            </div>
          </div>
        </div>
        <!-- Cost Breakdown Card -->
        <div
          class="bg-surface border border-border-default rounded-xl shadow-sm p-card-padding hover:-translate-y-[2px] transition-transform duration-300"
        >
          <div class="flex items-center justify-between border-b border-border-default pb-3 mb-4">
            <h3 class="font-h3 text-h3 text-text-primary">Financials</h3>
            <button class="text-primary-container text-sm font-medium hover:underline">
              Edit Items
            </button>
          </div>
          <!-- Mini Table -->
          <div class="space-y-2 mb-4">
            <div class="flex justify-between items-center py-1">
              <span class="font-body text-body text-text-secondary"> Compressor Unit </span>
              <span class="font-mono-data text-mono-data text-text-primary"> ₹2,45,000 </span>
            </div>
            <div class="flex justify-between items-center py-1">
              <span class="font-body text-body text-text-secondary"> Refrigerant (20kg) </span>
              <span class="font-mono-data text-mono-data text-text-primary"> ₹18,500 </span>
            </div>
            <div class="flex justify-between items-center py-1">
              <span class="font-body text-body text-text-secondary"> Labor (Est. 12h) </span>
              <span class="font-mono-data text-mono-data text-text-primary"> ₹14,400 </span>
            </div>
          </div>
          <div
            class="bg-metric-amber/20 p-4 rounded-lg border border-metric-amber/40 flex flex-col items-center justify-center text-center"
          >
            <span
              class="font-metadata text-metadata text-surface-tint uppercase tracking-wider mb-1"
            >
              Estimated Total Cost
            </span>
            <span
              v-if="order"
              class="font-kpi-number text-kpi-number text-text-primary tracking-tight"
            >
              ₹{{ (order.estimatedCost || 0).toLocaleString() }}
            </span>
          </div>
        </div>
        <!-- Actions Card -->
        <div class="bg-surface border border-border-default rounded-xl shadow-sm p-card-padding">
          <h3 class="font-h3 text-h3 text-text-primary border-b border-border-default pb-3 mb-4">
            Update Status
          </h3>
          <div class="space-y-3">
            <button
              class="w-full py-2.5 bg-canvas border border-outline rounded-lg font-h3 text-h3 text-text-primary hover:bg-surface-dim transition-colors flex items-center justify-center gap-2"
            >
              <span class="material-symbols-outlined text-[18px]"> pause_circle </span>
              Put on Hold
            </button>
            <button
              class="w-full py-2.5 bg-status-in-stock/20 border border-status-in-stock/40 rounded-lg font-h3 text-h3 text-status-in-stock hover:bg-status-in-stock/30 transition-colors flex items-center justify-center gap-2"
            >
              <span class="material-symbols-outlined text-[18px]"> check_circle </span>
              Mark as Resolved
            </button>
            <div class="pt-2">
              <button
                class="w-full py-2.5 bg-surface-variant text-text-disabled rounded-lg font-h3 text-h3 transition-colors flex items-center justify-center gap-2 opacity-50 cursor-not-allowed"
                disabled
              >
                <span class="material-symbols-outlined text-[18px]"> lock </span>
                Close Work Order
              </button>
              <p class="text-xs text-text-secondary text-center mt-2">
                Order must be resolved before closing.
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';

const route = useRoute();
const router = useRouter();

const order = ref<any>(null);
const loading = ref(true);

const fetchOrder = async () => {
  try {
    const res = await fetch(`http://localhost:8080/api/work-orders/${route.params['id']}`);
    if (res.ok) {
      order.value = await res.json();
    }
  } catch (error) {
    console.error('Error fetching work order:', error);
  } finally {
    loading.value = false;
  }
};

const formatDate = (dateString: string) => {
  if (!dateString) return 'N/A';
  return new Date(dateString).toLocaleDateString('en-GB', {
    day: '2-digit',
    month: 'short',
    year: 'numeric',
  });
};

onMounted(fetchOrder);
</script>
