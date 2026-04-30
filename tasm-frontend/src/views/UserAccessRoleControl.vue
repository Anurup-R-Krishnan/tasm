<template>
<main class="mt-[60px] p-page-margin max-w-[1400px]">
 <!-- Page Header -->
 <div class="flex items-center justify-between mb-section-gap">
  <div>
   <h2 class="font-h1 text-h1 text-text-primary mb-1">
    User &amp; Role Management
   </h2>
   <p class="font-body text-body text-text-secondary">
    Manage system access, roles, and department assignments.
   </p>
  </div>
  <button class="bg-primary-container text-on-primary px-4 py-2.5 rounded-lg flex items-center gap-2 font-h3 text-h3 hover:bg-surface-tint transition-colors shadow-sm active:scale-95 duration-150">
   <span class="material-symbols-outlined" style="font-size: 18px;">
    person_add
   </span>
   Add New User
  </button>
 </div>
 <!-- Role Metrics Grid -->
 <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-stack mb-section-gap">
  <!-- Card 1 -->
  <div class="bg-surface rounded-xl border border-border-default p-card-padding shadow-[0_4px_16px_-4px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform duration-200 cursor-pointer relative overflow-hidden group">
   <div class="absolute top-0 right-0 w-24 h-24 bg-metric-amber/30 rounded-bl-full -mr-4 -mt-4 transition-transform group-hover:scale-110">
   </div>
   <div class="flex items-center gap-3 mb-4 relative z-10">
    <div class="w-10 h-10 rounded-lg bg-metric-amber flex items-center justify-center text-surface-tint">
     <span class="material-symbols-outlined">
      shield_person
     </span>
    </div>
    <span class="font-h3 text-h3 text-text-primary">
     Administrators
    </span>
   </div>
   <div class="relative z-10">
    <div class="font-kpi-number text-kpi-number text-text-primary mb-1">
     12
    </div>
    <p class="font-metadata text-metadata text-text-secondary">
     Full system access &amp; settings
    </p>
   </div>
  </div>
  <!-- Card 2 -->
  <div class="bg-surface rounded-xl border border-border-default p-card-padding shadow-[0_4px_16px_-4px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform duration-200 cursor-pointer relative overflow-hidden group">
   <div class="absolute top-0 right-0 w-24 h-24 bg-metric-sage/30 rounded-bl-full -mr-4 -mt-4 transition-transform group-hover:scale-110">
   </div>
   <div class="flex items-center gap-3 mb-4 relative z-10">
    <div class="w-10 h-10 rounded-lg bg-metric-sage flex items-center justify-center text-status-in-stock">
     <span class="material-symbols-outlined">
      assignment_ind
     </span>
    </div>
    <span class="font-h3 text-h3 text-text-primary">
     Asset Managers
    </span>
   </div>
   <div class="relative z-10">
    <div class="font-kpi-number text-kpi-number text-text-primary mb-1">
     48
    </div>
    <p class="font-metadata text-metadata text-text-secondary">
     Write access for inventory
    </p>
   </div>
  </div>
  <!-- Card 3 -->
  <div class="bg-surface rounded-xl border border-border-default p-card-padding shadow-[0_4px_16px_-4px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform duration-200 cursor-pointer relative overflow-hidden group">
   <div class="absolute top-0 right-0 w-24 h-24 bg-tertiary-fixed/30 rounded-bl-full -mr-4 -mt-4 transition-transform group-hover:scale-110">
   </div>
   <div class="flex items-center gap-3 mb-4 relative z-10">
    <div class="w-10 h-10 rounded-lg bg-tertiary-fixed flex items-center justify-center text-tertiary">
     <span class="material-symbols-outlined">
      engineering
     </span>
    </div>
    <span class="font-h3 text-h3 text-text-primary">
     IT Technicians
    </span>
   </div>
   <div class="relative z-10">
    <div class="font-kpi-number text-kpi-number text-text-primary mb-1">
     36
    </div>
    <p class="font-metadata text-metadata text-text-secondary">
     Maintenance &amp; ticketing
    </p>
   </div>
  </div>
  <!-- Card 4 -->
  <div class="bg-surface rounded-xl border border-border-default p-card-padding shadow-[0_4px_16px_-4px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform duration-200 cursor-pointer relative overflow-hidden group">
   <div class="absolute top-0 right-0 w-24 h-24 bg-metric-lavender/30 rounded-bl-full -mr-4 -mt-4 transition-transform group-hover:scale-110">
   </div>
   <div class="flex items-center gap-3 mb-4 relative z-10">
    <div class="w-10 h-10 rounded-lg bg-metric-lavender flex items-center justify-center text-secondary">
     <span class="material-symbols-outlined">
      visibility
     </span>
    </div>
    <span class="font-h3 text-h3 text-text-primary">
     Auditors
    </span>
   </div>
   <div class="relative z-10">
    <div class="font-kpi-number text-kpi-number text-text-primary mb-1">
     8
    </div>
    <p class="font-metadata text-metadata text-text-secondary">
     Global read-only access
    </p>
   </div>
  </div>
 </div>
 <!-- Controls / Filters -->
 <div class="flex flex-col md:flex-row items-center justify-between gap-stack mb-stack bg-surface p-4 rounded-xl border border-border-default shadow-sm">
  <div class="flex items-center gap-inline w-full md:w-auto">
   <div class="relative w-full md:w-80">
    <span class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-text-secondary" style="font-size: 20px;">
     search
    </span>
    <input class="w-full pl-10 pr-4 py-2 bg-canvas border border-border-default rounded-lg font-body text-body text-text-primary outline-none focus:border-primary-container focus:ring-1 focus:ring-primary-container transition-all" placeholder="Search users by name or email..." type="text"/>
   </div>
   <div class="h-6 w-px bg-border-default hidden md:block">
   </div>
   <select class="bg-canvas border border-border-default text-text-primary font-body text-body rounded-lg py-2 pl-3 pr-8 outline-none focus:border-primary-container focus:ring-1 focus:ring-primary-container appearance-none hidden md:block" style="background-image: url('data:image/svg+xml;charset=US-ASCII,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20width%3D%22292.4%22%20height%3D%22292.4%22%3E%3Cpath%20fill%3D%22%231C1917%22%20d%3D%22M287%2069.4a17.6%2017.6%200%200%200-13-5.4H18.4c-5%200-9.3%201.8-12.9%205.4A17.6%2017.6%200%200%200%200%2082.2c0%205%201.8%209.3%205.4%2012.9l128%20127.9c3.6%203.6%207.8%205.4%2012.8%205.4s9.2-1.8%2012.8-5.4L287%2095c3.5-3.5%205.4-7.8%205.4-12.8%200-5-1.9-9.2-5.5-12.8z%22%2F%3E%3C%2Fsvg%3E'); background-repeat: no-repeat; background-position: right .7rem top 50%; background-size: .65rem auto;">
    <option value="">
     All Departments
    </option>
    <option value="it">
     Information Technology
    </option>
    <option value="facilities">
     Facilities Management
    </option>
    <option value="finance">
     Finance &amp; Procurement
    </option>
    <option value="hr">
     Human Resources
    </option>
   </select>
  </div>
  <div class="flex items-center gap-inline w-full md:w-auto">
   <button class="flex items-center gap-2 px-3 py-2 border border-border-default rounded-lg font-h3 text-h3 text-text-primary hover:bg-canvas transition-colors">
    <span class="material-symbols-outlined" style="font-size: 18px;">
     filter_list
    </span>
    More Filters
   </button>
   <button class="flex items-center justify-center p-2 border border-border-default rounded-lg text-text-secondary hover:bg-canvas transition-colors">
    <span class="material-symbols-outlined" style="font-size: 20px;">
     download
    </span>
   </button>
  </div>
 </div>
 <!-- User Table -->
 <div class="bg-surface border border-border-default rounded-xl overflow-hidden shadow-[0_4px_16px_-4px_rgba(0,0,0,0.02)] p-4">
  <DataTable :value="users" :loading="loading" paginator :rows="10" tableStyle="min-width: 50rem" class="w-full text-left">
   <Column field="name" header="User" sortable>
    <template #body="slotProps">
     <div class="flex items-center gap-3">
      <div class="w-9 h-9 rounded-full bg-surface-container flex items-center justify-center border border-border-default font-h3 text-h3 text-primary-container">
       {{ slotProps.data.name.charAt(0) }}
      </div>
      <div>
       <div class="font-h3 text-h3 text-text-primary">
        {{ slotProps.data.name }}
       </div>
       <div class="font-metadata text-metadata text-text-secondary">
        {{ slotProps.data.email }}
       </div>
      </div>
     </div>
    </template>
   </Column>
   <Column field="role" header="Role" sortable></Column>
   <Column field="department" header="Department" sortable></Column>
   <Column field="status" header="Status" sortable>
    <template #body="slotProps">
     <span class="inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full font-metadata text-metadata font-medium"
           :class="slotProps.data.status === 'Active' ? 'bg-metric-sage text-status-in-stock border border-status-in-stock/20' : 'bg-error-container text-status-critical border border-status-critical/20'">
      <span class="w-1.5 h-1.5 rounded-full" :class="slotProps.data.status === 'Active' ? 'bg-status-in-stock' : 'bg-status-critical'"></span>
      {{ slotProps.data.status }}
     </span>
    </template>
   </Column>
   <Column field="lastLogin" header="Last Login" sortable>
    <template #body="slotProps">
     <span class="font-mono-data text-mono-data text-text-secondary">
      {{ new Date(slotProps.data.lastLogin).toLocaleString() }}
     </span>
    </template>
   </Column>
   <Column header="">
    <template #body>
     <button class="text-text-secondary hover:text-text-primary p-1 rounded transition-colors">
      <span class="material-symbols-outlined">
       more_vert
      </span>
     </button>
    </template>
   </Column>
  </DataTable>
 </div>
</main>

</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'

interface SystemUser {
  id: number;
  employeeId: string;
  name: string;
  email: string;
  department: string;
  role: string;
  status: string;
  lastLogin: string;
}

const users = ref<SystemUser[]>([])
const loading = ref(true)

const fetchUsers = async () => {
  try {
    const res = await fetch('http://localhost:8080/api/users')
    if (res.ok) {
      users.value = await res.json()
    }
  } catch (error) {
    console.error('Failed to fetch users:', error)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchUsers()
})
</script>
