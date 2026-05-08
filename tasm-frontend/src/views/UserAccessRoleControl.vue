<template>
  <main class="space-y-section-gap pb-24">
    <!-- Page Header -->
    <div class="flex items-center justify-between mb-section-gap">
      <div>
        <h2 class="font-h1 text-h1 text-text-primary mb-1">User &amp; Role Management</h2>
        <p class="font-body text-text-secondary">
          Manage system access, roles, and department assignments.
        </p>
      </div>
      <button
        @click="router.push('/settings')"
        class="bg-primary text-on-primary px-4 py-2.5 rounded-lg flex items-center gap-2 font-h3 text-h3 hover:bg-primary/90 transition-colors shadow-sm active:scale-95 duration-150"
      >
        <span class="material-symbols-outlined" style="font-size: 18px"> person_add </span>
        Add New User
      </button>
    </div>
    <!-- Role Metrics Grid -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-stack mb-section-gap">
      <div
        v-for="role in roleMetrics"
        :key="role.roleName"
        class="bg-surface rounded-xl border border-border-default p-card-padding shadow-[0_4px_16px_-4px_rgba(0,0,0,0.02)] hover:-translate-y-0.5 transition-transform duration-200 cursor-pointer relative overflow-hidden group"
        @click="showRoleDetails(role.id)"
      >
        <div
          class="absolute top-0 right-0 w-24 h-24 rounded-bl-full -mr-4 -mt-4 transition-transform group-hover:scale-110"
          :class="role.accent"
        ></div>
        <div class="flex items-center gap-3 mb-4 relative z-10">
          <div class="w-10 h-10 rounded-lg flex items-center justify-center" :class="role.iconBg">
            <span class="material-symbols-outlined" :class="role.iconColor"> {{ role.icon }} </span>
          </div>
          <span class="font-h3 text-h3 text-text-primary"> {{ role.roleName }} </span>
        </div>
        <div class="relative z-10">
          <div class="font-kpi-number text-kpi-number text-text-primary mb-1">
            {{ role.usersCount }}
          </div>
          <p class="font-metadata text-metadata text-text-secondary">
            {{ role.description }}
          </p>
        </div>
      </div>
    </div>
    <!-- Controls / Filters -->
    <div
      class="flex flex-col md:flex-row items-center justify-between gap-stack mb-stack bg-surface p-4 rounded-xl border border-border-default shadow-sm"
    >
      <div class="flex items-center gap-inline w-full md:w-auto">
        <div class="relative w-full md:w-80">
          <span
            class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-text-secondary"
            style="font-size: 20px"
          >
            search
          </span>
          <input
            v-model="searchQuery"
            class="w-full pl-10 pr-4 py-2 bg-canvas border border-border-default rounded-lg font-body text-text-primary outline-none focus:border-primary-container focus:ring-1 focus:ring-primary-container transition-all"
            placeholder="Search users by name or email..."
            type="text"
          />
        </div>
        <div class="h-6 w-px bg-border-default hidden md:block"></div>
        <select
          v-model="selectedDepartment"
          class="bg-canvas border border-border-default text-text-primary font-body rounded-lg py-2 pl-3 pr-8 outline-none focus:border-primary-container focus:ring-1 focus:ring-primary-container appearance-none hidden md:block"
          style="
            background-image: url('data:image/svg+xml;charset=US-ASCII,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20width%3D%22292.4%22%20height%3D%22292.4%22%3E%3Cpath%20fill%3D%22%231C1917%22%20d%3D%22M287%2069.4a17.6%2017.6%200%200%200-13-5.4H18.4c-5%200-9.3%201.8-12.9%205.4A17.6%2017.6%200%200%200%200%2082.2c0%205%201.8%209.3%205.4%2012.9l128%20127.9c3.6%203.6%207.8%205.4%2012.8%205.4s9.2-1.8%2012.8-5.4L287%2095c3.5-3.5%205.4-7.8%205.4-12.8%200-5-1.9-9.2-5.5-12.8z%22%2F%3E%3C%2Fsvg%3E');
            background-repeat: no-repeat;
            background-position: right 0.7rem top 50%;
            background-size: 0.65rem auto;
          "
        >
          <option value="">All Departments</option>
          <option value="Information Technology">Information Technology</option>
          <option value="Facilities Management">Facilities Management</option>
          <option value="Finance & Procurement">Finance & Procurement</option>
          <option value="Human Resources">Human Resources</option>
        </select>
      </div>
      <div class="flex items-center gap-inline w-full md:w-auto">
        <button
          class="flex items-center gap-2 px-3 py-2 border border-border-default rounded-lg font-h3 text-h3 text-text-primary hover:bg-canvas transition-colors"
        >
          <span class="material-symbols-outlined" style="font-size: 18px"> filter_list </span>
          More Filters
        </button>
        <button
          @click="handleExport"
          class="flex items-center justify-center p-2 border border-border-default rounded-lg text-text-secondary hover:bg-canvas transition-colors"
        >
          <span class="material-symbols-outlined" style="font-size: 20px"> download </span>
        </button>
      </div>
    </div>
    <!-- User Table -->
    <div
      class="bg-surface border border-border-default rounded-xl overflow-hidden shadow-[0_4px_16px_-4px_rgba(0,0,0,0.02)] p-4"
    >
      <DataTable
        :value="filteredUsers"
        :loading="loading"
        paginator
        :rows="10"
        tableStyle="min-width: 50rem"
        class="w-full text-left"
      >
        <Column field="name" header="User" sortable>
          <template #body="slotProps">
            <div class="flex items-center gap-3">
              <div
                class="w-9 h-9 rounded-full bg-surface-container flex items-center justify-center border border-border-default font-h3 text-h3 text-primary-container"
              >
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
            <span
              class="inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full font-metadata text-metadata font-medium"
              :class="
                slotProps.data.status === 'Active'
                  ? 'bg-status-in-stock/20 text-status-in-stock border border-status-in-stock/30'
                  : 'bg-status-critical/10 text-status-critical border border-status-critical/30'
              "
            >
              <span
                class="w-1.5 h-1.5 rounded-full"
                :class="
                  slotProps.data.status === 'Active' ? 'bg-status-in-stock' : 'bg-status-critical'
                "
              ></span>
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
            <button
              class="text-text-secondary hover:text-text-primary p-1 rounded transition-colors"
            >
              <span class="material-symbols-outlined"> more_vert </span>
            </button>
          </template>
        </Column>
      </DataTable>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useRouter } from 'vue-router';
import { getUsers } from '../api/users';
import { getRoles, getRoleById } from '../api/roles';
import type { UserRole } from '../types/models';
import type { SystemUser } from '../types/models';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';

const router = useRouter();

const users = ref<SystemUser[]>([]);
const roles = ref<UserRole[]>([]);
const loading = ref(true);
const searchQuery = ref('');
const selectedDepartment = ref('');

const filteredUsers = computed(() => {
  let result = users.value;

  if (selectedDepartment.value) {
    result = result.filter(
      (u) => u.department.toLowerCase() === selectedDepartment.value.toLowerCase(),
    );
  }

  if (searchQuery.value) {
    const q = searchQuery.value.toLowerCase();
    result = result.filter(
      (u) => u.name.toLowerCase().includes(q) || u.email.toLowerCase().includes(q),
    );
  }

  return result;
});

const fetchUsers = async () => {
  try {
    users.value = (await getUsers()) as any[];
  } catch (error) {
    console.error('Failed to fetch users:', error);
  } finally {
    loading.value = false;
  }
};

const fetchRoles = async () => {
  try {
    roles.value = await getRoles();
  } catch (error) {
    console.error('Failed to fetch roles:', error);
  }
};

const showRoleDetails = async (roleId: number) => {
  try {
    const role = await getRoleById(roleId);
    alert(`${role.roleName}\n${role.description}\nUsers: ${role.usersCount}`);
  } catch (error) {
    console.error('Failed to load role details:', error);
    alert('Failed to load role details');
  }
};

const roleMetrics = computed(() => {
  const palette = [
    {
      accent: 'bg-metric-amber/30',
      iconBg: 'bg-metric-amber',
      iconColor: 'text-surface-tint',
      icon: 'shield_person',
    },
    {
      accent: 'bg-metric-sage/30',
      iconBg: 'bg-metric-sage',
      iconColor: 'text-status-in-stock',
      icon: 'assignment_ind',
    },
    {
      accent: 'bg-tertiary-fixed/30',
      iconBg: 'bg-tertiary-fixed',
      iconColor: 'text-tertiary',
      icon: 'engineering',
    },
    {
      accent: 'bg-metric-lavender/30',
      iconBg: 'bg-metric-lavender',
      iconColor: 'text-secondary',
      icon: 'visibility',
    },
  ];

  return roles.value.map((role, index) => ({
    ...role,
    ...palette[index % palette.length],
  }));
});

const handleExport = () => {
  const headers = ['Name', 'Email', 'Role', 'Department', 'Status', 'Last Login'];
  const rows = filteredUsers.value.map((u) => [
    u.name,
    u.email,
    u.role,
    u.department,
    u.status,
    new Date(u.lastLogin).toLocaleString(),
  ]);

  const csvContent = [headers, ...rows].map((e) => e.join(',')).join('\n');
  const blob = new Blob([csvContent], { type: 'text/csv;charset=utf-8;' });
  const link = document.createElement('a');
  link.href = URL.createObjectURL(blob);
  link.download = `tasm_users_export_${new Date().toISOString().split('T')[0]}.csv`;
  link.click();
};

onMounted(() => {
  fetchUsers();
  fetchRoles();
});
</script>
