<template>
<main class="flex-1 flex flex-col md:ml-[248px] min-h-screen max-w-[1400px]">
 <!-- TopNavBar (Mobile Only) -->
 <header class="md:hidden flex justify-between items-center px-6 w-full bg-stone-50 dark:bg-stone-950 text-amber-700 dark:text-amber-500 font-sans text-sm font-medium tracking-tight fixed top-0 h-[60px] z-40 border-b border-stone-200 dark:border-stone-800 shadow-sm dark:shadow-none">
  <div class="text-lg font-extrabold text-stone-900 dark:text-stone-50">
   Technopark Kerala
  </div>
  <div class="flex items-center gap-4">
   <button class="text-stone-600 dark:text-stone-400 hover:bg-stone-100 dark:hover:bg-stone-900 transition-colors p-2 rounded-full active:scale-95 transition-transform duration-150">
    <span class="material-symbols-outlined">
     notifications
    </span>
   </button>
   <button class="text-stone-600 dark:text-stone-400 hover:bg-stone-100 dark:hover:bg-stone-900 transition-colors p-2 rounded-full active:scale-95 transition-transform duration-150">
    <span class="material-symbols-outlined">
     help
    </span>
   </button>
   <button class="text-stone-600 dark:text-stone-400 hover:bg-stone-100 dark:hover:bg-stone-900 transition-colors p-2 rounded-full active:scale-95 transition-transform duration-150">
    <span class="material-symbols-outlined">
     settings
    </span>
   </button>
  </div>
 </header>
 <!-- Page Content -->
 <div class="flex-1 p-page-margin mt-[60px] md:mt-0 flex flex-col gap-section-gap">
  <!-- Header Section -->
  <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
   <div>
    <h1 class="font-h1 text-h1 text-text-primary">
     User Management
    </h1>
    <p class="font-metadata text-metadata text-text-secondary mt-1">
     Manage system access, roles, and department assignments.
    </p>
   </div>
   <div class="flex items-center gap-3 w-full md:w-auto">
    <div class="relative w-full md:w-64">
     <span class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-text-secondary text-[20px]">
      search
     </span>
     <input class="w-full bg-surface border border-border-default rounded-lg pl-10 pr-4 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-primary/20 focus:border-primary transition-all" placeholder="Search users..." type="text"/>
    </div>
    <button class="bg-primary hover:bg-primary/90 text-on-primary py-2 px-4 rounded-lg font-h3 text-h3 flex items-center justify-center gap-2 transition-colors shadow-sm active:scale-[0.98] whitespace-nowrap">
     <span class="material-symbols-outlined text-[18px]">
      person_add
     </span>
     Add New User
    </button>
   </div>
  </div>
  <!-- Filters & Actions Bar -->
  <div class="bg-surface border border-border-default rounded-xl p-4 flex flex-wrap items-center gap-4 shadow-[0_4px_4px_-4px_rgba(0,0,0,0.05)]">
   <div class="flex items-center gap-2">
    <span class="material-symbols-outlined text-text-secondary text-[20px]">
     filter_list
    </span>
    <span class="font-h3 text-h3 text-text-primary">
     Filters:
    </span>
   </div>
   <div class="flex gap-2">
    <select class="bg-surface-subtle border border-border-default rounded-lg px-3 py-1.5 text-sm font-metadata text-text-primary focus:outline-none focus:ring-2 focus:ring-primary/20">
     <option value="">
      All Roles
     </option>
     <option value="admin">
      Admin
     </option>
     <option value="asset_manager">
      Asset Manager
     </option>
     <option value="dept_head">
      Dept Head
     </option>
     <option value="finance">
      Finance
     </option>
     <option value="auditor">
      Auditor
     </option>
    </select>
    <select class="bg-surface-subtle border border-border-default rounded-lg px-3 py-1.5 text-sm font-metadata text-text-primary focus:outline-none focus:ring-2 focus:ring-primary/20">
     <option value="">
      All Departments
     </option>
     <option value="it">
      IT Services
     </option>
     <option value="facilities">
      Facilities
     </option>
     <option value="finance">
      Finance Dept
     </option>
     <option value="hr">
      Human Resources
     </option>
    </select>
    <select class="bg-surface-subtle border border-border-default rounded-lg px-3 py-1.5 text-sm font-metadata text-text-primary focus:outline-none focus:ring-2 focus:ring-primary/20">
     <option value="">
      Status: All
     </option>
     <option value="active">
      Active
     </option>
     <option value="inactive">
      Inactive
     </option>
    </select>
   </div>
   <div class="ml-auto flex items-center gap-2">
    <button class="text-text-secondary hover:text-text-primary p-1.5 rounded-md hover:bg-surface-subtle transition-colors" title="Export List">
     <span class="material-symbols-outlined text-[20px]">
      download
     </span>
    </button>
    <button class="text-text-secondary hover:text-text-primary p-1.5 rounded-md hover:bg-surface-subtle transition-colors" title="Column Settings">
     <span class="material-symbols-outlined text-[20px]">
      view_column
     </span>
    </button>
   </div>
  </div>
  <!-- Data Table Card -->
  <div class="bg-surface border border-border-default rounded-xl shadow-[0_4px_4px_-4px_rgba(0,0,0,0.05)] overflow-hidden p-4">
   <DataTable :value="users" :loading="loading" paginator :rows="10" tableStyle="min-width: 50rem" class="w-full text-left">
    <Column field="name" header="User Details" sortable>
     <template #body="slotProps">
      <div class="flex items-center gap-3">
       <div class="w-10 h-10 rounded-full bg-surface-variant border border-border-default flex items-center justify-center text-on-surface-variant font-h3">
        {{ slotProps.data.name.charAt(0) }}
       </div>
       <div>
        <div class="font-h3 text-h3 text-text-primary group-hover:text-primary transition-colors">
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
      <span class="inline-flex items-center gap-1.5" :class="slotProps.data.status === 'Active' ? 'text-status-in-stock' : 'text-text-secondary'">
       <span class="w-2 h-2 rounded-full" :class="slotProps.data.status === 'Active' ? 'bg-status-in-stock' : 'bg-outline'"></span>
       <span class="font-metadata text-metadata font-medium">
        {{ slotProps.data.status }}
       </span>
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
    <Column header="Actions" bodyStyle="text-align: right">
     <template #body>
      <div class="flex items-center justify-end gap-2">
       <button class="p-1.5 text-text-secondary hover:text-primary hover:bg-surface-subtle rounded-md transition-colors" title="Edit User">
        <span class="material-symbols-outlined text-[18px]">
         edit
        </span>
       </button>
       <button class="p-1.5 text-text-secondary hover:text-status-critical hover:bg-error-container/50 rounded-md transition-colors" title="Deactivate">
        <span class="material-symbols-outlined text-[18px]">
         block
        </span>
       </button>
      </div>
     </template>
    </Column>
   </DataTable>
  </div>
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
