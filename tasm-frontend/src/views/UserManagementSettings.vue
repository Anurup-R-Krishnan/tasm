<template>
  <main class="space-y-section-gap pb-24">
    <!-- Header Section -->
    <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
      <div>
        <h1 class="font-h1 text-h1 text-text-primary">User Management</h1>
        <p class="font-metadata text-metadata text-text-secondary mt-1">
          Manage system access, roles, and department assignments.
        </p>
      </div>
      <div class="flex items-center gap-3 w-full md:w-auto">
        <div class="relative w-full md:w-64">
          <span
            class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-text-secondary text-[20px]"
          >
            search
          </span>
          <input
            v-model="searchQuery"
            class="w-full bg-surface border border-border-default rounded-lg pl-10 pr-4 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-primary/20 focus:border-primary transition-all"
            placeholder="Search users..."
            type="text"
          />
        </div>
        <button
          class="bg-primary hover:bg-primary/90 text-on-primary py-2 px-4 rounded-lg font-h3 text-h3 flex items-center justify-center gap-2 transition-colors shadow-sm active:scale-[0.98] whitespace-nowrap"
        >
          <span class="material-symbols-outlined text-[18px]"> person_add </span>
          Add New User
        </button>
      </div>
    </div>
    <!-- Filters & Actions Bar -->
    <div
      class="bg-surface border border-border-default rounded-xl p-4 flex flex-wrap items-center gap-4 shadow-[0_4px_4px_-4px_rgba(0,0,0,0.05)]"
    >
      <div class="flex items-center gap-2">
        <span class="material-symbols-outlined text-text-secondary text-[20px]"> filter_list </span>
        <span class="font-h3 text-h3 text-text-primary"> Filters: </span>
      </div>
      <div class="flex gap-2">
        <select
          v-model="selectedRole"
          class="bg-surface-subtle border border-border-default rounded-lg px-3 py-1.5 text-sm font-metadata text-text-primary focus:outline-none focus:ring-2 focus:ring-primary/20"
        >
          <option value="">All Roles</option>
          <option value="Admin">Admin</option>
          <option value="Asset Manager">Asset Manager</option>
          <option value="Dept Head">Dept Head</option>
          <option value="Finance">Finance</option>
          <option value="Auditor">Auditor</option>
        </select>
        <select
          v-model="selectedDepartment"
          class="bg-surface-subtle border border-border-default rounded-lg px-3 py-1.5 text-sm font-metadata text-text-primary focus:outline-none focus:ring-2 focus:ring-primary/20"
        >
          <option value="">All Departments</option>
          <option value="IT Services">IT Services</option>
          <option value="Facilities">Facilities</option>
          <option value="Finance Dept">Finance Dept</option>
          <option value="Human Resources">Human Resources</option>
        </select>
        <select
          v-model="selectedStatus"
          class="bg-surface-subtle border border-border-default rounded-lg px-3 py-1.5 text-sm font-metadata text-text-primary focus:outline-none focus:ring-2 focus:ring-primary/20"
        >
          <option value="">Status: All</option>
          <option value="Active">Active</option>
          <option value="Inactive">Inactive</option>
        </select>
      </div>
      <div class="ml-auto flex items-center gap-2">
        <button
          @click="handleExport"
          class="text-text-secondary hover:text-text-primary p-1.5 rounded-md hover:bg-surface-subtle transition-colors"
          title="Export List"
        >
          <span class="material-symbols-outlined text-[20px]"> download </span>
        </button>
        <button
          class="text-text-secondary hover:text-text-primary p-1.5 rounded-md hover:bg-surface-subtle transition-colors"
          title="Column Settings"
        >
          <span class="material-symbols-outlined text-[20px]"> view_column </span>
        </button>
      </div>
    </div>
    <!-- Data Table Card -->
    <div
      class="bg-surface border border-border-default rounded-xl shadow-[0_4px_4px_-4px_rgba(0,0,0,0.05)] overflow-hidden p-4"
    >
      <DataTable
        :value="filteredUsers"
        :loading="loading"
        paginator
        :rows="10"
        tableStyle="min-width: 50rem"
        class="w-full text-left"
      >
        <Column field="name" header="User Details" sortable>
          <template #body="slotProps">
            <div class="flex items-center gap-3">
              <div
                class="w-10 h-10 rounded-full bg-surface-variant border border-border-default flex items-center justify-center text-on-surface-variant font-h3"
              >
                {{ slotProps.data.name.charAt(0) }}
              </div>
              <div>
                <div
                  class="font-h3 text-h3 text-text-primary group-hover:text-primary transition-colors"
                >
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
            <span
              class="inline-flex items-center gap-1.5"
              :class="
                slotProps.data.status === 'Active' ? 'text-status-in-stock' : 'text-text-secondary'
              "
            >
              <span
                class="w-2 h-2 rounded-full"
                :class="slotProps.data.status === 'Active' ? 'bg-status-in-stock' : 'bg-outline'"
              ></span>
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
          <template #body="slotProps">
            <div class="flex items-center justify-end gap-2">
              <button
                class="p-1.5 text-text-secondary hover:text-primary hover:bg-surface-subtle rounded-md transition-colors"
                title="Edit User"
                @click="openUserDetails(slotProps.data.id)"
              >
                <span class="material-symbols-outlined text-[18px]"> edit </span>
              </button>
              <button
                @click="deleteUser(slotProps.data.id)"
                class="p-1.5 text-text-secondary hover:text-status-critical hover:bg-error-container/50 rounded-md transition-colors"
                title="Delete User"
              >
                <span class="material-symbols-outlined text-[18px]"> delete </span>
              </button>
            </div>
          </template>
        </Column>
      </DataTable>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { getUsers, deleteUser as deleteUserApi, getUserById } from '../api/users';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import type { SystemUser } from '../types/models';

const users = ref<SystemUser[]>([]);
const loading = ref(true);
const searchQuery = ref('');
const selectedRole = ref('');
const selectedDepartment = ref('');
const selectedStatus = ref('');

const filteredUsers = computed(() => {
  return users.value.filter((user) => {
    const matchesSearch =
      user.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      user.email.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      user.employeeId.toLowerCase().includes(searchQuery.value.toLowerCase());

    const matchesRole = !selectedRole.value || user.role === selectedRole.value;
    const matchesDept = !selectedDepartment.value || user.department === selectedDepartment.value;
    const matchesStatus = !selectedStatus.value || user.status === selectedStatus.value;

    return matchesSearch && matchesRole && matchesDept && matchesStatus;
  });
});

const fetchUsers = async () => {
  loading.value = true;
  try {
    users.value = await getUsers();
  } catch (error) {
    console.error('Failed to fetch users:', error);
  } finally {
    loading.value = false;
  }
};

const deleteUser = async (id: number) => {
  if (!confirm('Are you sure you want to delete this user?')) return;
  try {
    await deleteUserApi(id);
    users.value = users.value.filter((u) => u.id !== id);
  } catch (error) {
    console.error('Error deleting user:', error);
    alert('Failed to delete user');
  }
};

const openUserDetails = async (id: number) => {
  try {
    const user = await getUserById(id);
    alert(`User: ${user.name}\nRole: ${user.role}\nDepartment: ${user.department}`);
  } catch (error) {
    console.error('Failed to fetch user', error);
    alert('Failed to load user details');
  }
};

const handleExport = () => {
  const headers = ['Employee ID', 'Name', 'Email', 'Role', 'Department', 'Status', 'Last Login'];
  const rows = filteredUsers.value.map((u) => [
    u.employeeId,
    u.name,
    u.email,
    u.role,
    u.department,
    u.status,
    u.lastLogin,
  ]);

  const csvContent = [headers, ...rows].map((e) => e.join(',')).join('\n');
  const blob = new Blob([csvContent], { type: 'text/csv;charset=utf-8;' });
  const link = document.createElement('a');
  link.href = URL.createObjectURL(blob);
  link.download = `users_export_${new Date().toISOString().split('T')[0]}.csv`;
  link.click();
};

onMounted(() => {
  fetchUsers();
});
</script>
