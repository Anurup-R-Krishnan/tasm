<template>
  <main class="flex-1 flex flex-col min-h-screen relative">
    <!-- Page Content Canvas -->
    <div class="p-page-margin max-w-[800px] w-full mx-auto space-y-section-gap pb-24">
      <!-- Page Header & Progress -->
      <div class="mb-section-gap">
        <h1 class="font-h1 text-h1 text-text-primary mb-2">Check-Out &amp; Assignment Wizard</h1>
        <p class="font-body text-body text-text-secondary mb-8">
          Assigning MacBook Pro M2 (Asset ID: MBP-2023-045)
        </p>
        <!-- Progress Bar -->
        <div class="relative">
          <div class="overflow-hidden h-2 mb-4 text-xs flex rounded-full bg-border-default">
            <div
              class="shadow-none flex flex-col text-center whitespace-nowrap text-white justify-center bg-primary"
              style="width: 50%"
            ></div>
          </div>
          <div class="flex justify-between font-metadata text-metadata text-text-secondary">
            <span class="text-primary font-medium">1. Asset Details</span>
            <span class="text-primary font-medium">2. Recipient Selection</span>
            <span>3. Terms &amp; Conditions</span>
            <span>4. Confirmation</span>
          </div>
        </div>
      </div>
      <!-- Wizard Card -->
      <div class="bg-surface rounded-xl border border-border-default shadow-sm p-card-padding">
        <h2 class="font-h2 text-h2 text-text-primary mb-6">Select Recipient</h2>
        <form>
          <!-- Search & Selection -->
          <div class="mb-8 space-y-6">
            <div>
              <label class="block font-h3 text-h3 text-text-primary mb-2" for="employee-search"
                >Search Employee</label
              >
              <div class="relative">
                <span
                  class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-text-secondary"
                  >search</span
                >
                <input
                  class="w-full pl-10 pr-4 py-2.5 rounded-lg border border-border-default focus:ring-2 focus:ring-primary focus:border-primary font-body text-body outline-none transition-shadow bg-surface-subtle"
                  id="employee-search"
                  placeholder="Name, Employee ID, or Email"
                  type="text"
                  value="Sarah Jenkins"
                />
                <button
                  class="absolute right-3 top-1/2 -translate-y-1/2 text-text-secondary hover:text-text-primary"
                  type="button"
                >
                  <span class="material-symbols-outlined text-[20px]">close</span>
                </button>
              </div>
            </div>
            <!-- Selected Employee Info (Bento-style layout) -->
            <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
              <!-- Profile Card -->
              <div
                class="col-span-1 md:col-span-2 bg-surface-subtle border border-border-default rounded-lg p-4 flex items-start gap-4"
              >
                <div
                  class="w-12 h-12 rounded-full bg-primary-container text-on-primary-container flex items-center justify-center font-h2 text-h2 shrink-0"
                >
                  SJ
                </div>
                <div>
                  <h3 class="font-h2 text-h2 text-text-primary">Sarah Jenkins</h3>
                  <p class="font-body text-body text-text-secondary mb-1">Senior UX Designer</p>
                  <div class="flex flex-wrap gap-2 mt-2">
                    <span
                      class="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium bg-surface-variant text-on-surface-variant border border-outline-variant"
                    >
                      <span class="material-symbols-outlined text-[14px] mr-1">badge</span>
                      EMP-8942
                    </span>
                    <span
                      class="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium bg-surface-variant text-on-surface-variant border border-outline-variant"
                    >
                      <span class="material-symbols-outlined text-[14px] mr-1">domain</span>
                      Product Design
                    </span>
                  </div>
                </div>
              </div>
              <!-- Current Assets Summary -->
              <div
                class="col-span-1 bg-surface-subtle border border-border-default rounded-lg p-4 flex flex-col justify-center items-center text-center"
              >
                <span
                  class="font-metadata text-metadata text-text-secondary uppercase tracking-wider mb-1"
                  >Current Assets Held</span
                >
                <span class="font-kpi-number text-kpi-number text-text-primary">3</span>
                <a class="font-metadata text-metadata text-primary hover:underline mt-2" href="#"
                  >View List</a
                >
              </div>
            </div>
          </div>
          <!-- Assignment Details -->
          <div class="space-y-6 border-t border-border-default pt-6">
            <h3 class="font-h3 text-h3 text-text-primary">Assignment Details</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label
                  class="block font-metadata text-metadata text-text-secondary mb-2 uppercase tracking-wider"
                  for="return-date"
                  >Expected Return Date</label
                >
                <div class="relative">
                  <span
                    class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-text-secondary"
                    >calendar_today</span
                  >
                  <input
                    class="w-full pl-10 pr-4 py-2 rounded-lg border border-border-default focus:ring-2 focus:ring-primary focus:border-primary font-body text-body outline-none bg-surface"
                    id="return-date"
                    type="date"
                  />
                </div>
                <p class="font-metadata text-metadata text-text-secondary mt-1">
                  Leave blank for indefinite assignment.
                </p>
              </div>
              <div>
                <label
                  class="block font-metadata text-metadata text-text-secondary mb-2 uppercase tracking-wider"
                  for="location"
                  >Primary Location</label
                >
                <div class="relative">
                  <span
                    class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-text-secondary"
                    >location_on</span
                  >
                  <select
                    class="w-full pl-10 pr-8 py-2 rounded-lg border border-border-default focus:ring-2 focus:ring-primary focus:border-primary font-body text-body outline-none bg-surface appearance-none"
                    id="location"
                  >
                    <option>Trivandrum Campus - Building 1</option>
                    <option>Kochi Campus</option>
                    <option>Remote / Work from Home</option>
                  </select>
                  <span
                    class="material-symbols-outlined absolute right-3 top-1/2 -translate-y-1/2 text-text-secondary pointer-events-none"
                    >expand_more</span
                  >
                </div>
              </div>
            </div>
            <div>
              <label
                class="block font-metadata text-metadata text-text-secondary mb-2 uppercase tracking-wider"
                for="notes"
                >Assignment Notes (Optional)</label
              >
              <textarea
                class="w-full p-3 rounded-lg border border-border-default focus:ring-2 focus:ring-primary focus:border-primary font-body text-body outline-none bg-surface resize-none"
                id="notes"
                placeholder="Add any relevant details regarding this assignment..."
                rows="3"
              ></textarea>
            </div>
          </div>
          <!-- Actions -->
          <div class="mt-8 flex items-center justify-between pt-6 border-t border-border-default">
            <button
              class="px-4 py-2 border border-border-default text-text-primary rounded-lg font-h3 text-h3 hover:bg-surface-subtle transition-colors flex items-center gap-2"
              type="button"
            >
              <span class="material-symbols-outlined text-[18px]">arrow_back</span>
              Previous
            </button>
            <button
              class="px-6 py-2 bg-[#1C1917] text-white rounded-lg font-h3 text-h3 hover:bg-stone-800 transition-colors shadow-sm flex items-center gap-2"
              type="button"
            >
              Review Assignment
              <span class="material-symbols-outlined text-[18px]">arrow_forward</span>
            </button>
          </div>
        </form>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
// Check-Out & Assignment Wizard - Step 2: Recipient Selection
</script>
