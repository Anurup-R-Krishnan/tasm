<template>
  <main class="space-y-section-gap pb-24">
    <!-- Page Header & Tabs -->
    <div class="flex flex-col md:flex-row md:items-end justify-between gap-4 mb-section-gap">
      <div>
        <h2 class="font-h1 text-h1 text-text-primary mb-2">Reservations &amp; Bookings</h2>
        <p class="font-body text-body text-text-secondary">
          Manage asset availability and schedule requests.
        </p>
      </div>
      <div
        class="flex bg-surface-subtle p-1 rounded-lg border border-border-default self-start w-full md:w-auto overflow-x-auto"
      >
        <button
          @click="currentTab = 'Upcoming'"
          class="px-4 py-1.5 rounded-md font-h3 text-h3 transition-all whitespace-nowrap"
          :class="
            currentTab === 'Upcoming'
              ? 'bg-surface shadow-[0_1px_3px_rgba(0,0,0,0.05)] border border-border-default text-text-primary'
              : 'text-text-secondary hover:text-text-primary hover:bg-surface/50'
          "
        >
          Upcoming
        </button>
        <button
          @click="currentTab = 'Past'"
          class="px-4 py-1.5 rounded-md font-body text-body transition-all whitespace-nowrap"
          :class="
            currentTab === 'Past'
              ? 'bg-surface shadow-[0_1px_3px_rgba(0,0,0,0.05)] border border-border-default text-text-primary'
              : 'text-text-secondary hover:text-text-primary hover:bg-surface/50'
          "
        >
          Past
        </button>
        <button
          @click="currentTab = 'My Bookings'"
          class="px-4 py-1.5 rounded-md font-body text-body transition-all whitespace-nowrap"
          :class="
            currentTab === 'My Bookings'
              ? 'bg-surface shadow-[0_1px_3px_rgba(0,0,0,0.05)] border border-border-default text-text-primary'
              : 'text-text-secondary hover:text-text-primary hover:bg-surface/50'
          "
        >
          My Bookings
        </button>
      </div>
    </div>
    <!-- Bento Layout -->
    <div class="grid grid-cols-1 xl:grid-cols-12 gap-section-gap">
      <!-- Main Canvas (Calendar & List) -->
      <div class="xl:col-span-8 space-y-section-gap">
        <!-- Week View Calendar Card -->
        <div
          class="bg-surface rounded-xl border border-border-default shadow-[0_4px_24px_rgba(0,0,0,0.02)] overflow-hidden"
        >
          <div
            class="p-card-padding border-b border-border-default flex items-center justify-between bg-surface-subtle/50"
          >
            <div class="flex items-center gap-4">
              <h3 class="font-h2 text-h2 text-text-primary">{{ calendarRangeText }}</h3>
              <div class="flex gap-1">
                <button
                  @click="navigateCalendar('prev')"
                  class="p-1 text-text-secondary hover:text-text-primary hover:bg-surface rounded-md border border-transparent hover:border-border-default transition-all"
                >
                  <span class="material-symbols-outlined text-[18px]"> chevron_left </span>
                </button>
                <button
                  @click="navigateCalendar('next')"
                  class="p-1 text-text-secondary hover:text-text-primary hover:bg-surface rounded-md border border-transparent hover:border-border-default transition-all"
                >
                  <span class="material-symbols-outlined text-[18px]"> chevron_right </span>
                </button>
              </div>
            </div>
            <div class="flex gap-3 hidden sm:flex">
              <div class="flex items-center gap-1.5">
                <span class="w-2.5 h-2.5 rounded-full bg-tertiary-container"> </span>
                <span class="font-metadata text-metadata text-text-secondary"> Meeting Rooms </span>
              </div>
              <div class="flex items-center gap-1.5">
                <span class="w-2.5 h-2.5 rounded-full bg-surface-tint"> </span>
                <span class="font-metadata text-metadata text-text-secondary"> Vehicles </span>
              </div>
              <div class="flex items-center gap-1.5">
                <span class="w-2.5 h-2.5 rounded-full bg-status-in-stock"> </span>
                <span class="font-metadata text-metadata text-text-secondary"> Projectors </span>
              </div>
            </div>
          </div>
          <div class="p-4 overflow-x-auto">
            <div class="min-w-[700px]">
              <!-- Calendar Header -->
              <div class="grid grid-cols-7 gap-2 mb-4">
                <div v-for="day in calendarDays" :key="day.date" class="text-center">
                  <p
                    class="font-metadata text-metadata uppercase"
                    :class="day.isToday ? 'text-primary' : 'text-text-secondary'"
                  >
                    {{ day.weekday }}
                  </p>
                  <p
                    class="font-h3 text-h3 mt-0.5"
                    :class="
                      day.isToday
                        ? 'bg-primary text-on-primary w-8 h-8 rounded-full flex items-center justify-center mx-auto'
                        : 'text-text-primary'
                    "
                  >
                    {{ day.dayNumber }}
                  </p>
                </div>
              </div>
              <!-- Calendar Grid Data -->
              <div class="grid grid-cols-7 gap-2">
                <!-- Mon -->
                <div class="space-y-2">
                  <div
                    class="bg-tertiary-container/10 border border-tertiary-container/20 p-2 rounded-lg cursor-pointer hover:-translate-y-0.5 transition-transform"
                  >
                    <p
                      class="font-metadata text-metadata text-tertiary-container font-medium truncate"
                    >
                      Conf Room A
                    </p>
                    <p class="font-metadata text-[10px] text-tertiary-container/70">
                      09:00 - 11:00
                    </p>
                  </div>
                </div>
                <!-- Tue -->
                <div class="space-y-2">
                  <div
                    class="bg-surface-tint/10 border border-surface-tint/20 p-2 rounded-lg cursor-pointer hover:-translate-y-0.5 transition-transform"
                  >
                    <p class="font-metadata text-metadata text-surface-tint font-medium truncate">
                      Toyota Innova
                    </p>
                    <p class="font-metadata text-[10px] text-surface-tint/70">Full Day</p>
                  </div>
                </div>
                <!-- Wed (Today) -->
                <div class="space-y-2 relative">
                  <div
                    class="absolute -left-1 top-[40%] w-[calc(100%+8px)] h-0.5 bg-primary/20 z-0"
                  ></div>
                  <div
                    class="bg-status-in-stock/10 border border-status-in-stock/20 p-2 rounded-lg cursor-pointer hover:-translate-y-0.5 transition-transform relative z-10"
                  >
                    <p
                      class="font-metadata text-metadata text-status-in-stock font-medium truncate"
                    >
                      Epson PRO-9
                    </p>
                    <p class="font-metadata text-[10px] text-status-in-stock/70">13:00 - 15:00</p>
                  </div>
                  <div
                    class="bg-tertiary-container/10 border border-tertiary-container/20 p-2 rounded-lg cursor-pointer hover:-translate-y-0.5 transition-transform relative z-10"
                  >
                    <p
                      class="font-metadata text-metadata text-tertiary-container font-medium truncate"
                    >
                      Boardroom C
                    </p>
                    <p class="font-metadata text-[10px] text-tertiary-container/70">
                      15:30 - 17:00
                    </p>
                  </div>
                </div>
                <!-- Thu -->
                <div class="space-y-2">
                  <div
                    class="bg-surface-subtle border border-dashed border-border-default p-2 rounded-lg flex items-center justify-center h-12 hover:bg-surface-container-low cursor-pointer transition-colors"
                  >
                    <span class="material-symbols-outlined text-text-secondary text-[16px]">
                      add
                    </span>
                  </div>
                </div>
                <!-- Fri -->
                <div class="space-y-2">
                  <div
                    class="bg-surface-tint/10 border border-surface-tint/20 p-2 rounded-lg cursor-pointer hover:-translate-y-0.5 transition-transform"
                  >
                    <p class="font-metadata text-metadata text-surface-tint font-medium truncate">
                      Honda City
                    </p>
                    <p class="font-metadata text-[10px] text-surface-tint/70">08:00 - 18:00</p>
                  </div>
                </div>
                <!-- Sat -->
                <div class="space-y-2 bg-surface-subtle/50 rounded-lg min-h-[100px]"></div>
                <!-- Sun -->
                <div class="space-y-2 bg-surface-subtle/50 rounded-lg min-h-[100px]"></div>
              </div>
            </div>
          </div>
        </div>
        <!-- Active Reservations List -->
        <div
          class="bg-surface rounded-xl border border-border-default shadow-[0_4px_24px_rgba(0,0,0,0.02)] overflow-hidden"
        >
          <div
            class="p-card-padding border-b border-border-default flex items-center justify-between"
          >
            <h3 class="font-h2 text-h2 text-text-primary">Active Reservations</h3>
            <button class="text-surface-tint font-h3 text-h3 hover:underline">View All</button>
          </div>
          <div class="overflow-x-auto p-4">
            <DataTable
              :value="filteredReservations"
              :loading="loading"
              paginator
              :rows="5"
              class="w-full text-left"
              rowHover
            >
              <Column field="title" header="Asset / Resource" sortable>
                <template #body="slotProps">
                  <div class="flex items-center gap-3">
                    <div
                      class="w-8 h-8 rounded flex items-center justify-center"
                      :class="
                        slotProps.data.type === 'Meeting Room'
                          ? 'bg-tertiary-container/10 text-tertiary-container'
                          : slotProps.data.type === 'Vehicle'
                            ? 'bg-surface-tint/10 text-surface-tint'
                            : 'bg-status-in-stock/10 text-status-in-stock'
                      "
                    >
                      <span class="material-symbols-outlined text-[18px]">
                        {{
                          slotProps.data.type === 'Meeting Room'
                            ? 'meeting_room'
                            : slotProps.data.type === 'Vehicle'
                              ? 'directions_car'
                              : 'print'
                        }}
                      </span>
                    </div>
                    <div>
                      <p class="font-h3 text-h3 text-text-primary">
                        {{ slotProps.data.title }}
                      </p>
                      <p class="font-mono-data text-mono-data text-text-secondary mt-1">
                        {{ slotProps.data.location }}
                      </p>
                    </div>
                  </div>
                </template>
              </Column>
              <Column field="bookedBy" header="Reserved By" sortable></Column>
              <Column field="startTime" header="Date & Time" sortable>
                <template #body="slotProps">
                  <p class="font-body text-body text-text-primary">
                    {{ new Date(slotProps.data.startTime).toLocaleDateString() }}
                  </p>
                  <p class="font-metadata text-metadata text-text-secondary">
                    {{
                      new Date(slotProps.data.startTime).toLocaleTimeString([], {
                        hour: '2-digit',
                        minute: '2-digit',
                      })
                    }}
                    -
                    {{
                      new Date(slotProps.data.endTime).toLocaleTimeString([], {
                        hour: '2-digit',
                        minute: '2-digit',
                      })
                    }}
                  </p>
                </template>
              </Column>
              <Column field="status" header="Status" sortable>
                <template #body="slotProps">
                  <span
                    class="inline-flex items-center px-2 py-1 rounded-md font-metadata text-metadata border"
                    :class="
                      slotProps.data.status === 'Confirmed' || slotProps.data.status === 'Active'
                        ? 'bg-metric-sage/50 text-status-in-stock border-metric-sage'
                        : slotProps.data.status === 'Pending'
                          ? 'bg-metric-amber/50 text-surface-tint border-metric-amber'
                          : 'bg-error-container/20 text-status-critical border-error-container/30'
                    "
                  >
                    {{ slotProps.data.status }}
                  </span>
                </template>
              </Column>
              <Column header="Action" alignFrozen="right">
                <template #body="slotProps">
                  <button
                    @click="handleCancelBooking(slotProps.data.id)"
                    class="p-1.5 text-error hover:bg-error/10 rounded-md transition-colors"
                    title="Cancel Booking"
                  >
                    <span class="material-symbols-outlined text-[20px]"> delete </span>
                  </button>
                </template>
              </Column>
            </DataTable>
          </div>
        </div>
      </div>
      <!-- Right Sidebar (Quick Book Panel) -->
      <div class="xl:col-span-4">
        <div
          class="bg-surface rounded-xl border border-border-default shadow-[0_4px_24px_rgba(0,0,0,0.02)] p-card-padding sticky top-[92px]"
        >
          <div class="flex items-center gap-3 mb-6">
            <div
              class="w-10 h-10 rounded-lg bg-primary-container text-on-primary-container flex items-center justify-center"
            >
              <span class="material-symbols-outlined"> bolt </span>
            </div>
            <div>
              <h3 class="font-h2 text-h2 text-text-primary">Quick Book</h3>
              <p class="font-metadata text-metadata text-text-secondary">Reserve instantly</p>
            </div>
          </div>
          <form class="space-y-5" @submit.prevent="handleConfirmBooking">
            <!-- Asset Type -->
            <div class="space-y-2">
              <label class="font-h3 text-h3 text-text-primary block"> Asset Category </label>
              <div class="grid grid-cols-3 gap-2">
                <button
                  v-for="cat in ['Meeting Room', 'Vehicle', 'Equipment']"
                  :key="cat"
                  @click="currentCategory = cat"
                  class="border py-2 rounded-lg font-metadata text-metadata flex flex-col items-center gap-1 transition-all"
                  :class="
                    currentCategory === cat
                      ? 'border-surface-tint bg-surface-container-low text-surface-tint'
                      : 'border-border-default text-text-secondary hover:border-surface-tint hover:text-surface-tint bg-surface'
                  "
                  type="button"
                >
                  <span class="material-symbols-outlined text-[20px]">
                    {{
                      cat === 'Meeting Room'
                        ? 'meeting_room'
                        : cat === 'Vehicle'
                          ? 'directions_car'
                          : 'print'
                    }}
                  </span>
                  {{ cat === 'Meeting Room' ? 'Rooms' : cat === 'Vehicle' ? 'Vehicles' : 'Equip' }}
                </button>
              </div>
            </div>
            <!-- Search/Select -->
            <div class="space-y-2">
              <label class="font-h3 text-h3 text-text-primary block"> Select Resource </label>
              <div class="relative">
                <span
                  class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-text-secondary text-[18px]"
                >
                  search
                </span>
                <input
                  v-model="bookingForm.resourceId"
                  class="w-full pl-9 pr-3 py-2 border border-border-default rounded-lg font-body text-body focus:ring-2 focus:ring-surface-tint/20 focus:border-surface-tint outline-none bg-surface"
                  :placeholder="`Search ${currentCategory.toLowerCase()}s...`"
                  type="text"
                />
              </div>
            </div>
            <!-- Date Picker -->
            <div class="space-y-2">
              <label class="font-h3 text-h3 text-text-primary block"> Date </label>
              <div class="relative">
                <span
                  class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-text-secondary text-[18px]"
                >
                  calendar_today
                </span>
                <input
                  v-model="bookingForm.date"
                  class="w-full pl-9 pr-3 py-2 border border-border-default rounded-lg font-body text-body focus:ring-2 focus:ring-surface-tint/20 focus:border-surface-tint outline-none bg-surface text-text-primary"
                  type="date"
                />
              </div>
            </div>
            <!-- Time Split -->
            <div class="grid grid-cols-2 gap-3">
              <div class="space-y-2">
                <label class="font-h3 text-h3 text-text-primary block"> Start Time </label>
                <div class="relative">
                  <span
                    class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-text-secondary text-[18px]"
                  >
                    schedule
                  </span>
                  <input
                    v-model="bookingForm.startTime"
                    class="w-full pl-9 pr-3 py-2 border border-border-default rounded-lg font-body text-body focus:ring-2 focus:ring-surface-tint/20 focus:border-surface-tint outline-none bg-surface text-text-primary"
                    type="time"
                  />
                </div>
              </div>
              <div class="space-y-2">
                <label class="font-h3 text-h3 text-text-primary block"> End Time </label>
                <div class="relative">
                  <span
                    class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-text-secondary text-[18px]"
                  >
                    schedule
                  </span>
                  <input
                    v-model="bookingForm.endTime"
                    class="w-full pl-9 pr-3 py-2 border border-border-default rounded-lg font-body text-body focus:ring-2 focus:ring-surface-tint/20 focus:border-surface-tint outline-none bg-surface text-text-primary"
                    type="time"
                  />
                </div>
              </div>
            </div>
            <!-- Notes -->
            <div class="space-y-2">
              <label class="font-h3 text-h3 text-text-primary block">
                Purpose / Notes
                <span class="text-text-secondary font-normal"> (Optional) </span>
              </label>
              <textarea
                v-model="bookingForm.notes"
                class="w-full p-3 border border-border-default rounded-lg font-body text-body focus:ring-2 focus:ring-surface-tint/20 focus:border-surface-tint outline-none bg-surface resize-none"
                placeholder="Brief description..."
                rows="2"
              ></textarea>
            </div>
            <button
              class="w-full bg-[#1C1917] hover:bg-stone-800 text-white font-h3 text-h3 py-3 rounded-lg transition-colors flex items-center justify-center gap-2 mt-4"
              type="submit"
            >
              Confirm Booking
              <span class="material-symbols-outlined text-[18px]"> arrow_forward </span>
            </button>
          </form>
        </div>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { getReservations } from '../api/reservations';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import type { Reservation } from '../types/models';
import { useAuth } from '../composables/useAuth';

const { currentUser } = useAuth();
const reservations = ref<Reservation[]>([]);
const loading = ref(true);
const currentTab = ref('Upcoming');
const currentCategory = ref('Meeting Room');

// Form state for Quick Book
const bookingForm = ref({
  resourceId: '',
  date: new Date().toISOString().split('T')[0],
  startTime: '09:00',
  endTime: '11:00',
  notes: '',
});

const calendarBaseDate = ref(new Date());

const calendarRangeText = computed(() => {
  const start = new Date(calendarBaseDate.value);
  start.setDate(start.getDate() - start.getDay() + 1); // Monday
  const end = new Date(start);
  end.setDate(end.getDate() + 6); // Sunday

  const options: Intl.DateTimeFormatOptions = { month: 'long', day: 'numeric', year: 'numeric' };
  return `${start.toLocaleDateString(undefined, { month: 'long', day: 'numeric' })} - ${end.toLocaleDateString(undefined, options)}`;
});

const calendarDays = computed(() => {
  const start = new Date(calendarBaseDate.value);
  start.setDate(start.getDate() - start.getDay() + 1);
  const days = [];
  const today = new Date().toDateString();

  for (let i = 0; i < 7; i++) {
    const d = new Date(start);
    d.setDate(d.getDate() + i);
    days.push({
      date: d.toISOString().split('T')[0],
      dayNumber: d.getDate(),
      weekday: d.toLocaleDateString(undefined, { weekday: 'short' }),
      isToday: d.toDateString() === today,
    });
  }
  return days;
});

const filteredReservations = computed(() => {
  const now = new Date();
  let base = reservations.value;

  if (currentTab.value === 'Upcoming') {
    return base.filter((r) => new Date(r.startTime) >= now);
  } else if (currentTab.value === 'Past') {
    return base.filter((r) => new Date(r.startTime) < now);
  } else if (currentTab.value === 'My Bookings') {
    return base.filter(
      (r) => r.userId === currentUser.value?.id || r.bookedBy === currentUser.value?.name,
    );
  }
  return base;
});

const fetchReservations = async () => {
  loading.value = true;
  try {
    const data = await getReservations();
    reservations.value = data;
  } catch (error) {
    console.error('Failed to fetch reservations:', error);
  } finally {
    loading.value = false;
  }
};

const handleConfirmBooking = () => {
  if (!bookingForm.value.date || !bookingForm.value.resourceId) {
    alert('Please provide resource name and date.');
    return;
  }

  const newBooking = {
    id: Date.now(),
    title: bookingForm.value.resourceId,
    type: currentCategory.value,
    bookedBy: currentUser.value?.name || 'Current User',
    startTime: `${bookingForm.value.date}T${bookingForm.value.startTime}:00`,
    endTime: `${bookingForm.value.date}T${bookingForm.value.endTime}:00`,
    status: 'Confirmed',
    location: 'Main Campus',
  };

  // Simulate API call
  reservations.value = [newBooking as any, ...reservations.value];
  alert(`Booking confirmed for ${newBooking.title} on ${bookingForm.value.date}`);

  // Reset form
  bookingForm.value.resourceId = '';
  bookingForm.value.notes = '';
};

const handleCancelBooking = (id: number) => {
  if (confirm('Are you sure you want to cancel this reservation?')) {
    reservations.value = reservations.value.filter((r) => r.id !== id);
    alert('Reservation cancelled successfully.');
  }
};

const navigateCalendar = (direction: 'prev' | 'next') => {
  const newDate = new Date(calendarBaseDate.value);
  if (direction === 'prev') {
    newDate.setDate(newDate.getDate() - 7);
  } else {
    newDate.setDate(newDate.getDate() + 7);
  }
  calendarBaseDate.value = newDate;
};

onMounted(() => {
  fetchReservations();
});
</script>
