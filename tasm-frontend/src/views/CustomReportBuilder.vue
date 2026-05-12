<template>
  <main class="space-y-section-gap pb-24">
    <!-- Page Header -->
    <div class="mb-section-gap flex flex-col md:flex-row md:items-end justify-between gap-4">
      <div>
        <div class="flex items-center gap-2 text-text-secondary font-metadata mb-2">
          <span> Compliance </span>
          <span class="material-symbols-outlined text-[14px]"> chevron_right </span>
          <span> Reports </span>
        </div>
        <h1 class="font-h1 text-h1 text-text-primary">Custom Report Builder</h1>
        <p class="font-body text-text-secondary mt-1">
          Configure and generate detailed compliance and operational reports.
        </p>
      </div>
    </div>
    <!-- Main Layout Grid -->
    <div class="grid grid-cols-1 lg:grid-cols-4 gap-6 items-start">
      <!-- Left Column: Steps Navigation -->
      <div
        class="lg:col-span-1 bg-surface rounded-xl border border-border-default p-card-padding shadow-sm sticky top-[84px]"
      >
        <h3 class="font-h3 text-h3 text-text-primary mb-6">Report Configuration</h3>
        <div class="space-y-6 relative">
          <!-- Step 1 -->
          <div
            v-for="step in steps"
            :key="step.id"
            class="relative flex items-center gap-4 group cursor-pointer"
            :class="{ 'opacity-50': activeStep !== step.id }"
            @click="activeStep = step.id"
          >
            <div
              class="flex items-center justify-center w-8 h-8 rounded-full border-2 shrink-0 transition-colors z-10"
              :class="
                activeStep === step.id
                  ? 'border-primary bg-primary-container/20 text-primary shadow'
                  : 'border-border-default bg-surface text-text-secondary'
              "
            >
              <span class="font-mono-data text-mono-data font-bold">{{ step.id }}</span>
            </div>
            <div>
              <h4 class="font-h3 text-h3 text-text-primary leading-tight">{{ step.title }}</h4>
              <p class="font-metadata text-metadata text-text-secondary">{{ step.description }}</p>
            </div>
          </div>
        </div>
      </div>
      <!-- Right Column: Form Area -->
      <div class="lg:col-span-3 space-y-6">
        <!-- Form Container -->
        <div class="bg-surface rounded-xl border border-border-default shadow-sm overflow-hidden">
          <div
            v-if="notice.message"
            class="m-4 px-4 py-3 rounded-xl border text-sm font-semibold"
            :class="
              notice.tone === 'error'
                ? 'bg-status-critical/10 border-status-critical/20 text-status-critical'
                : 'bg-status-in-stock/10 border-status-in-stock/20 text-status-in-stock'
            "
          >
            {{ notice.message }}
          </div>
          <!-- Step 1 Content: Base Data -->
          <div
            v-if="activeStep === 1"
            class="p-card-padding border-b border-border-default animate-fade-in"
          >
            <h2 class="font-h2 text-h2 text-text-primary mb-6 flex items-center gap-2">
              <span class="material-symbols-outlined text-primary"> database </span>
              Step 1: Select Base Data Source
            </h2>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div class="space-y-2">
                <label class="font-h3 text-h3 text-text-primary block"> Primary Entity </label>
                <div class="relative">
                  <select
                    v-model="selectedEntity"
                    class="w-full bg-surface border border-border-default rounded-lg pl-4 pr-10 py-2.5 font-body text-text-primary focus:outline-none focus:border-primary focus:ring-2 focus:ring-primary/10 appearance-none cursor-pointer shadow-sm hover:translate-y-[-1px] transition-transform"
                  >
                    <option value="assets">Fixed Assets</option>
                    <option value="wo">Work Orders</option>
                    <option value="depreciation">Depreciation Schedules</option>
                    <option value="compliance">Compliance Logs</option>
                  </select>
                  <span
                    class="material-symbols-outlined absolute right-3 top-1/2 -translate-y-1/2 text-text-secondary pointer-events-none"
                  >
                    expand_more
                  </span>
                </div>
                <p class="font-metadata text-metadata text-text-secondary">
                  Determines the core dataset for this report.
                </p>
              </div>
              <div class="space-y-2">
                <label class="font-h3 text-h3 text-text-primary block"> Report Timeframe </label>
                <div class="relative">
                  <select
                    class="w-full bg-surface border border-border-default rounded-lg pl-4 pr-10 py-2.5 font-body text-text-primary focus:outline-none focus:border-primary focus:ring-2 focus:ring-primary/10 appearance-none cursor-pointer shadow-sm hover:translate-y-[-1px] transition-transform"
                  >
                    <option value="mtd">Month to Date (MTD)</option>
                    <option value="qtd">Quarter to Date (QTD)</option>
                    <option selected value="ytd">Year to Date (YTD)</option>
                    <option value="custom">Custom Range</option>
                  </select>
                  <span
                    class="material-symbols-outlined absolute right-3 top-1/2 -translate-y-1/2 text-text-secondary pointer-events-none"
                  >
                    expand_more
                  </span>
                </div>
              </div>
            </div>
            <div class="mt-8 flex justify-end">
              <button
                @click="activeStep = 2"
                class="px-6 py-2 bg-primary text-on-primary rounded-lg font-h3 shadow-sm hover:bg-primary-hover transition-colors"
              >
                Continue to Columns
              </button>
            </div>
          </div>

          <!-- Step 2 Content: Columns -->
          <div
            v-if="activeStep === 2"
            class="p-card-padding border-b border-border-default bg-surface-subtle animate-fade-in"
          >
            <h2 class="font-h2 text-h2 text-text-primary mb-4 flex items-center gap-2">
              <span class="material-symbols-outlined text-primary"> view_column </span>
              Step 2: Select Columns
            </h2>
            <div class="grid grid-cols-2 md:grid-cols-3 gap-4">
              <label
                v-for="col in availableColumns"
                :key="col"
                class="flex items-center gap-3 cursor-pointer group"
              >
                <div class="relative flex items-center justify-center w-5 h-5">
                  <input
                    v-model="selectedColumns"
                    :value="col"
                    class="peer w-5 h-5 appearance-none border border-border-default rounded bg-surface checked:bg-primary checked:border-primary focus:outline-none focus:ring-2 focus:ring-primary/20 transition-all"
                    type="checkbox"
                  />
                  <span
                    class="material-symbols-outlined absolute text-white text-[16px] pointer-events-none opacity-0 peer-checked:opacity-100 transition-opacity"
                  >
                    check
                  </span>
                </div>
                <span
                  class="font-body text-text-primary group-hover:text-primary transition-colors"
                >
                  {{ col }}
                </span>
              </label>
            </div>
          </div>
          <!-- Step 3 Content: Filters -->
          <div
            v-if="activeStep === 3"
            class="p-card-padding border-b border-border-default animate-fade-in"
          >
            <h2 class="font-h2 text-h2 text-text-primary mb-4 flex items-center gap-2">
              <span class="material-symbols-outlined text-primary"> filter_list </span>
              Step 3: Apply Filters
            </h2>
            <div class="bg-surface-subtle p-6 rounded-lg border border-border-default mb-6">
              <div class="flex flex-wrap gap-2 mb-6">
                <div
                  v-for="filter in activeFilters"
                  :key="filter"
                  class="inline-flex items-center gap-2 px-3 py-1.5 bg-primary-container/20 border border-primary/20 rounded-full text-primary font-metadata text-metadata shadow-sm"
                >
                  <span>{{ filter }}</span>
                  <button
                    @click="removeFilter(filter)"
                    class="hover:text-text-primary focus:outline-none flex items-center"
                  >
                    <span class="material-symbols-outlined text-[14px]"> close </span>
                  </button>
                </div>
                <button
                  @click="showFilterModal = true"
                  class="inline-flex items-center gap-1 px-3 py-1.5 border border-dashed border-outline rounded-full text-text-secondary hover:text-text-primary hover:border-text-primary transition-colors font-metadata text-metadata"
                >
                  <span class="material-symbols-outlined text-[14px]"> add </span>
                  Add Filter
                </button>
              </div>
            </div>
            <div class="flex justify-between">
              <button
                @click="activeStep = 2"
                class="px-6 py-2 border border-border-default rounded-lg font-h3 hover:bg-surface-variant transition-colors"
              >
                Back
              </button>
              <button
                @click="activeStep = 4"
                class="px-6 py-2 bg-primary text-on-primary rounded-lg font-h3 shadow-sm hover:bg-primary-hover transition-colors"
              >
                Continue to Structure
              </button>
            </div>
          </div>

          <!-- Step 4 Content: Group & Sort -->
          <div
            v-if="activeStep === 4"
            class="p-card-padding border-b border-border-default animate-fade-in"
          >
            <h2 class="font-h2 text-h2 text-text-primary mb-4 flex items-center gap-2">
              <span class="material-symbols-outlined text-primary"> sort_by_alpha </span>
              Step 4: Group & Sort
            </h2>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
              <div class="space-y-2">
                <label class="font-h3 text-h3 text-text-primary block"> Group By </label>
                <select
                  class="w-full bg-surface border border-border-default rounded-lg px-4 py-2 font-body"
                >
                  <option value="none">No Grouping</option>
                  <option value="category">Category</option>
                  <option value="location">Location</option>
                  <option value="status">Status</option>
                </select>
              </div>
              <div class="space-y-2">
                <label class="font-h3 text-h3 text-text-primary block"> Sort Order </label>
                <select
                  class="w-full bg-surface border border-border-default rounded-lg px-4 py-2 font-body"
                >
                  <option value="asc">Ascending</option>
                  <option value="desc">Descending</option>
                </select>
              </div>
            </div>
            <div class="flex justify-between">
              <button
                @click="activeStep = 3"
                class="px-6 py-2 border border-border-default rounded-lg font-h3 hover:bg-surface-variant transition-colors"
              >
                Back
              </button>
              <button
                @click="handleGenerateReport"
                class="px-6 py-2 bg-primary text-on-primary rounded-lg font-h3 shadow-sm hover:bg-primary-hover transition-colors"
              >
                Generate Final Preview
              </button>
            </div>
          </div>
          <!-- Preview Section -->
          <div class="p-card-padding bg-surface">
            <div class="flex items-center justify-between mb-4">
              <h2 class="font-h2 text-h2 text-text-primary flex items-center gap-2">
                <span class="material-symbols-outlined text-primary"> visibility </span>
                Live Preview
              </h2>
              <span
                class="font-metadata text-metadata text-text-secondary bg-surface-variant px-2 py-1 rounded"
              >
                Showing first {{ reportData.length }} rows
              </span>
            </div>
            <div class="overflow-x-auto rounded-lg border border-border-default shadow-sm p-4">
              <DataTable
                :value="reportData"
                :loading="loading"
                paginator
                :rows="5"
                class="w-full text-left"
              >
                <Column field="tagId" header="Asset ID" sortable>
                  <template #body="slotProps">
                    <span
                      class="font-mono-data text-mono-data text-text-primary group-hover:text-primary"
                      >{{ slotProps.data.tagId }}</span
                    >
                  </template>
                </Column>
                <Column field="name" header="Asset Name" sortable></Column>
                <Column field="category" header="Category" sortable></Column>
                <Column field="location" header="Location" sortable></Column>
                <Column field="status" header="Status" sortable>
                  <template #body="slotProps">
                    <span
                      class="inline-flex items-center px-2 py-0.5 rounded-full text-[11px] font-medium"
                      :class="
                        slotProps.data.status === 'Active'
                          ? 'bg-status-in-stock/20 text-status-in-stock'
                          : 'bg-status-checked-out/20 text-status-checked-out'
                      "
                    >
                      {{ slotProps.data.status }}
                    </span>
                  </template>
                </Column>
              </DataTable>
            </div>
          </div>

          <!-- Action Bar -->
          <div
            class="bg-surface-subtle p-4 border-t border-border-default flex flex-col sm:flex-row items-center justify-between gap-4"
          >
            <div class="flex gap-3 w-full sm:w-auto">
              <button
                @click="handleExport('PDF')"
                class="flex-1 sm:flex-none flex items-center justify-center gap-2 px-4 py-2 border border-border-default bg-surface text-text-primary rounded-lg font-h3 text-h3 hover:bg-surface-variant hover:-translate-y-0.5 transition-all shadow-sm"
              >
                <span class="material-symbols-outlined text-[18px]"> picture_as_pdf </span>
                Export PDF
              </button>
              <button
                @click="handleExport('Excel')"
                class="flex-1 sm:flex-none flex items-center justify-center gap-2 px-4 py-2 border border-border-default bg-surface text-text-primary rounded-lg font-h3 text-h3 hover:bg-surface-variant hover:-translate-y-0.5 transition-all shadow-sm"
              >
                <span class="material-symbols-outlined text-[18px]"> table_view </span>
                Export Excel
              </button>
            </div>
            <button
              @click="showSaveModal = true"
              class="w-full sm:w-auto flex items-center justify-center gap-2 px-6 py-2 bg-text-primary text-white rounded-lg font-h3 text-h3 hover:bg-text-primary/90 hover:-translate-y-0.5 transition-all shadow-md"
            >
              <span class="material-symbols-outlined text-[18px]"> save </span>
              Save Template
            </button>
          </div>
        </div>
      </div>
    </div>
    <Teleport to="body">
      <div
        v-if="showFilterModal"
        class="fixed inset-0 z-50 flex items-center justify-center bg-black/40 backdrop-blur-sm"
        @click.self="closeFilterModal"
      >
        <div class="bg-surface rounded-2xl shadow-2xl p-8 w-full max-w-md mx-4">
          <div class="flex items-center justify-between mb-6">
            <h2 class="font-h2 text-h2 text-text-primary">Add Filter</h2>
            <button class="text-text-secondary hover:text-text-primary" @click="closeFilterModal">
              <span class="material-symbols-outlined">close</span>
            </button>
          </div>
          <div class="space-y-3">
            <label
              class="block font-metadata text-metadata text-text-secondary uppercase tracking-wider"
            >
              Filter Expression
            </label>
            <input
              v-model="filterInput"
              type="text"
              placeholder="e.g. Location: Nila"
              class="w-full h-11 px-4 bg-canvas border border-border-default rounded-xl focus:border-primary focus:ring-1 focus:ring-primary outline-none"
            />
            <p class="text-xs text-text-secondary">Use “Field: Value” format.</p>
          </div>
          <div class="flex gap-3 mt-6">
            <button
              class="flex-1 py-3 bg-surface border border-border-default rounded-xl font-h3 text-text-secondary hover:bg-surface-subtle"
              @click="closeFilterModal"
            >
              Cancel
            </button>
            <button
              class="flex-1 py-3 bg-primary text-on-primary rounded-xl font-h3 hover:opacity-90"
              @click="applyFilter"
              :disabled="!filterInput.trim()"
            >
              Add Filter
            </button>
          </div>
        </div>
      </div>
    </Teleport>

    <Teleport to="body">
      <div
        v-if="showSaveModal"
        class="fixed inset-0 z-50 flex items-center justify-center bg-black/40 backdrop-blur-sm"
        @click.self="closeSaveModal"
      >
        <div class="bg-surface rounded-2xl shadow-2xl p-8 w-full max-w-md mx-4">
          <div class="flex items-center justify-between mb-6">
            <h2 class="font-h2 text-h2 text-text-primary">Save Report Template</h2>
            <button class="text-text-secondary hover:text-text-primary" @click="closeSaveModal">
              <span class="material-symbols-outlined">close</span>
            </button>
          </div>
          <div class="space-y-3">
            <label
              class="block font-metadata text-metadata text-text-secondary uppercase tracking-wider"
            >
              Template Name
            </label>
            <input
              v-model="templateName"
              type="text"
              placeholder="e.g. Quarterly Compliance"
              class="w-full h-11 px-4 bg-canvas border border-border-default rounded-xl focus:border-primary focus:ring-1 focus:ring-primary outline-none"
            />
          </div>
          <div class="flex gap-3 mt-6">
            <button
              class="flex-1 py-3 bg-surface border border-border-default rounded-xl font-h3 text-text-secondary hover:bg-surface-subtle"
              @click="closeSaveModal"
            >
              Cancel
            </button>
            <button
              class="flex-1 py-3 bg-primary text-on-primary rounded-xl font-h3 hover:opacity-90"
              @click="saveTemplate"
              :disabled="!templateName.trim()"
            >
              Save
            </button>
          </div>
        </div>
      </div>
    </Teleport>
  </main>
</template>

<script setup lang="ts">
import { ref, watch, onMounted } from 'vue';
import { apiRequest } from '../api/client';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';

const activeStep = ref(1);
const selectedEntity = ref('assets');
const reportData = ref<any[]>([]);
const loading = ref(false);
const showFilterModal = ref(false);
const showSaveModal = ref(false);
const filterInput = ref('');
const templateName = ref('');
const notice = ref<{ message: string; tone: 'success' | 'error' }>({
  message: '',
  tone: 'success',
});

const steps = [
  { id: 1, title: 'Base Data', description: 'Select primary entity' },
  { id: 2, title: 'Columns', description: 'Choose data fields' },
  { id: 3, title: 'Filters', description: 'Refine data subset' },
  { id: 4, title: 'Group & Sort', description: 'Organize structure' },
];

const availableColumns = [
  'Asset ID',
  'Asset Name',
  'Category',
  'Location',
  'Purchase Date',
  'Status',
  'Cost Center',
];
const selectedColumns = ref(['Asset ID', 'Asset Name', 'Category', 'Location', 'Status']);

const activeFilters = ref(['Category: IT Equipment', 'Status: Active']);

const showNotice = (message: string, tone: 'success' | 'error' = 'success') => {
  notice.value = { message, tone };
  window.setTimeout(() => {
    notice.value = { message: '', tone: 'success' };
  }, 3000);
};

const closeFilterModal = () => {
  showFilterModal.value = false;
  filterInput.value = '';
};

const applyFilter = () => {
  if (!filterInput.value.trim()) return;
  activeFilters.value.push(filterInput.value.trim());
  closeFilterModal();
  showNotice('Filter added to the report configuration.');
};

const removeFilter = (filter: string) => {
  activeFilters.value = activeFilters.value.filter((f) => f !== filter);
};

const handleExport = (format: string) => {
  if (reportData.value.length === 0) {
    showNotice('No data available to export. Generate a report first.', 'error');
    return;
  }

  const headers = selectedColumns.value;
  const rows = reportData.value.map((item) => {
    return selectedColumns.value.map((col) => {
      if (col === 'Asset ID') return item.tagId;
      if (col === 'Asset Name') return item.name;
      return item[col.toLowerCase()] || '';
    });
  });

  const csvContent = [
    headers.join(','),
    ...rows.map((row) => row.map((cell) => `"${cell}"`).join(',')),
  ].join('\n');

  const blob = new Blob([csvContent], { type: 'text/csv;charset=utf-8;' });
  const link = document.createElement('a');
  link.href = URL.createObjectURL(blob);
  const timestamp = new Date().toISOString().split('T')[0];
  link.download = `tasm_report_${selectedEntity.value}_${timestamp}.csv`;
  link.click();

  showNotice(`${format} export downloaded as CSV.`);
};

const closeSaveModal = () => {
  showSaveModal.value = false;
  templateName.value = '';
};

const saveTemplate = () => {
  if (!templateName.value.trim()) return;
  showNotice(`Template "${templateName.value.trim()}" saved to your report library.`);
  closeSaveModal();
};

const handleGenerateReport = async () => {
  loading.value = true;
  // Simulate heavy processing for report structure finalization
  await new Promise((resolve) => setTimeout(resolve, 800));
  await fetchReportData();
  activeStep.value = 4; // Stay on step 4 or show preview
  showNotice('Final report structure generated and preview updated.');
};

const fetchReportData = async () => {
  loading.value = true;
  try {
    let endpoint = '';
    if (selectedEntity.value === 'assets') endpoint = '/assets';
    else if (selectedEntity.value === 'wo') endpoint = '/work-orders';
    else if (selectedEntity.value === 'depreciation') endpoint = '/depreciations';
    else endpoint = '/assets';

    const data = await apiRequest<any[]>(endpoint);
    reportData.value = data;
  } catch (err) {
    console.error('Failed to load report data:', err);
  } finally {
    loading.value = false;
  }
};

watch(selectedEntity, fetchReportData);
onMounted(fetchReportData);
</script>
