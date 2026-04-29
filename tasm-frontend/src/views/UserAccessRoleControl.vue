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
 <div class="bg-surface border border-border-default rounded-xl overflow-hidden shadow-[0_4px_16px_-4px_rgba(0,0,0,0.02)]">
  <div class="overflow-x-auto">
   <table class="w-full text-left border-collapse min-w-[800px]">
    <thead>
     <tr class="bg-surface-subtle border-b border-border-default">
      <th class="py-4 px-6 font-table-header text-table-header text-text-secondary uppercase w-1/3">
       User
      </th>
      <th class="py-4 px-6 font-table-header text-table-header text-text-secondary uppercase">
       Role
      </th>
      <th class="py-4 px-6 font-table-header text-table-header text-text-secondary uppercase">
       Department
      </th>
      <th class="py-4 px-6 font-table-header text-table-header text-text-secondary uppercase">
       Access Status
      </th>
      <th class="py-4 px-6 font-table-header text-table-header text-text-secondary uppercase">
       Last Login
      </th>
      <th class="py-4 px-6 font-table-header text-table-header text-text-secondary uppercase text-right w-16">
      </th>
     </tr>
    </thead>
    <tbody class="divide-y divide-border-default">
     <!-- Row 1 -->
     <tr class="hover:bg-accent-subtle/30 transition-colors group">
      <td class="py-4 px-6">
       <div class="flex items-center gap-3">
        <img alt="Profile" class="w-9 h-9 rounded-full border border-border-default object-cover" data-alt="Portrait of a smiling professional woman with dark hair in a bright office" src="https://lh3.googleusercontent.com/aida-public/AB6AXuDmmHr5rl5vdrwqm9w7iDUo8vaNtdekBvJttBtGpiPWLEVWWrZKUpppikebNfPyJXguRQWKvI29-iZ9keFbw9ggad5EGy9s0wfHS-8Atc9uxeg3t7Uo9ReY8_70-hPUhhvSrrFR_AhyKZtchLdEKNg5Nw4Di7hXiV3tNnZNYxbb4NiQv4wNxOSGNZZk-56mfJpCE-rDHcnBDwfNfqC8rguOlV-U0ulRatM0RL3PCQ4DgyRZSslMw--7m6oDr2UZsQNE5wI-AZHCPe5w"/>
        <div>
         <div class="font-h3 text-h3 text-text-primary">
          Sarah Jenkins
         </div>
         <div class="font-metadata text-metadata text-text-secondary">
          s.jenkins@technopark.gov.in
         </div>
        </div>
       </div>
      </td>
      <td class="py-4 px-6 font-body text-body text-text-primary">
       Administrator
      </td>
      <td class="py-4 px-6 font-body text-body text-text-primary">
       Executive Board
      </td>
      <td class="py-4 px-6">
       <span class="inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full bg-metric-sage text-status-in-stock font-metadata text-metadata font-medium border border-status-in-stock/20">
        <span class="sr-only">
         Status:
        </span>
        <span class="w-1.5 h-1.5 rounded-full bg-status-in-stock">
        </span>
        Active
       </span>
      </td>
      <td class="py-4 px-6 font-mono-data text-mono-data text-text-secondary">
       Today, 09:41 AM
      </td>
      <td class="py-4 px-6 text-right">
       <button class="text-text-secondary hover:text-text-primary p-1 rounded transition-colors opacity-0 group-hover:opacity-100 focus:opacity-100">
        <span class="material-symbols-outlined">
         more_vert
        </span>
       </button>
      </td>
     </tr>
     <!-- Row 2 -->
     <tr class="hover:bg-accent-subtle/30 transition-colors group">
      <td class="py-4 px-6">
       <div class="flex items-center gap-3">
        <div class="w-9 h-9 rounded-full bg-surface-container flex items-center justify-center border border-border-default font-h3 text-h3 text-primary-container">
         MR
        </div>
        <div>
         <div class="font-h3 text-h3 text-text-primary">
          Michael Rodriguez
         </div>
         <div class="font-metadata text-metadata text-text-secondary">
          m.rodriguez@technopark.gov.in
         </div>
        </div>
       </div>
      </td>
      <td class="py-4 px-6 font-body text-body text-text-primary">
       IT Technician
      </td>
      <td class="py-4 px-6 font-body text-body text-text-primary">
       Information Technology
      </td>
      <td class="py-4 px-6">
       <span class="inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full bg-metric-sage text-status-in-stock font-metadata text-metadata font-medium border border-status-in-stock/20">
        <span class="sr-only">
         Status:
        </span>
        <span class="w-1.5 h-1.5 rounded-full bg-status-in-stock">
        </span>
        Active
       </span>
      </td>
      <td class="py-4 px-6 font-mono-data text-mono-data text-text-secondary">
       Yesterday, 14:22
      </td>
      <td class="py-4 px-6 text-right">
       <button class="text-text-secondary hover:text-text-primary p-1 rounded transition-colors opacity-0 group-hover:opacity-100 focus:opacity-100">
        <span class="material-symbols-outlined">
         more_vert
        </span>
       </button>
      </td>
     </tr>
     <!-- Row 3 -->
     <tr class="hover:bg-accent-subtle/30 transition-colors group">
      <td class="py-4 px-6">
       <div class="flex items-center gap-3">
        <img alt="Profile" class="w-9 h-9 rounded-full border border-border-default object-cover" data-alt="Portrait of a smiling man with short hair wearing a light blue shirt" src="https://lh3.googleusercontent.com/aida-public/AB6AXuASUb36bbvLji1RjvsCfZhonojQh3Ydq62U-rzOIBhonWa5XRhokM2LTqM4_GkP4q11a498DfCdWUtdVf9lJ_MffRWoZ42OGxvc1tRIDKQC3AIKNcJUI1lyiwkrti6yR5F6OaE1YvanKVQQgL5xyCrUgJcjxm3Aulue7tHNVC2psuGg5Ej2gNbNzQFe4jPzzDl7NmyghzGynO2VZpQHppu5b6POATp2T6VIsFLWXcD1IjVZEHkeOZwsQk5sRi42NHth98OeIPjld6_F"/>
        <div>
         <div class="font-h3 text-h3 text-text-primary opacity-50">
          David Chen
         </div>
         <div class="font-metadata text-metadata text-text-secondary opacity-50">
          d.chen@technopark.gov.in
         </div>
        </div>
       </div>
      </td>
      <td class="py-4 px-6 font-body text-body text-text-primary opacity-50">
       Asset Manager
      </td>
      <td class="py-4 px-6 font-body text-body text-text-primary opacity-50">
       Facilities Management
      </td>
      <td class="py-4 px-6">
       <span class="inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full bg-error-container text-status-critical font-metadata text-metadata font-medium border border-status-critical/20">
        <span class="sr-only">
         Status:
        </span>
        <span class="w-1.5 h-1.5 rounded-full bg-status-critical">
        </span>
        Locked
       </span>
      </td>
      <td class="py-4 px-6 font-mono-data text-mono-data text-text-secondary">
       Oct 12, 2023
      </td>
      <td class="py-4 px-6 text-right">
       <button class="text-text-secondary hover:text-text-primary p-1 rounded transition-colors opacity-0 group-hover:opacity-100 focus:opacity-100">
        <span class="material-symbols-outlined">
         more_vert
        </span>
       </button>
      </td>
     </tr>
     <!-- Row 4 -->
     <tr class="hover:bg-accent-subtle/30 transition-colors group border-l-4 border-amber-500 bg-surface-subtle">
      <td class="py-4 px-6 pl-5">
       <div class="flex items-center gap-3">
        <div class="w-9 h-9 rounded-full bg-metric-lavender flex items-center justify-center border border-border-default font-h3 text-h3 text-secondary">
         AK
        </div>
        <div>
         <div class="font-h3 text-h3 text-text-primary">
          Anita Kumar
         </div>
         <div class="font-metadata text-metadata text-text-secondary">
          a.kumar@ext.technopark.gov.in
         </div>
        </div>
       </div>
      </td>
      <td class="py-4 px-6 font-body text-body text-text-primary">
       Auditor
      </td>
      <td class="py-4 px-6 font-body text-body text-text-primary">
       External Audit
      </td>
      <td class="py-4 px-6">
       <span class="inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full bg-metric-amber text-surface-tint font-metadata text-metadata font-medium border border-surface-tint/20">
        <span class="sr-only">
         Status:
        </span>
        <span class="w-1.5 h-1.5 rounded-full bg-surface-tint">
        </span>
        Pending Review
       </span>
      </td>
      <td class="py-4 px-6 font-mono-data text-mono-data text-text-secondary italic">
       Never logged in
      </td>
      <td class="py-4 px-6 text-right">
       <button class="text-text-primary p-1 rounded hover:bg-canvas transition-colors">
        <span class="material-symbols-outlined">
         more_vert
        </span>
       </button>
      </td>
     </tr>
     <!-- Row 5 -->
     <tr class="hover:bg-accent-subtle/30 transition-colors group">
      <td class="py-4 px-6">
       <div class="flex items-center gap-3">
        <img alt="Profile" class="w-9 h-9 rounded-full border border-border-default object-cover" data-alt="Portrait of a confident older professional woman in a grey blazer" src="https://lh3.googleusercontent.com/aida-public/AB6AXuDp8jjReme6UnFWyj7WKz6zROmH7sj64uRBG1KVFZ496INc1YU2mHO_Y0wdk-PSL4GM6Dbg2W02lieWcUMKUw_y2mtLw9QSTHBb8UHtypXxi_gLl4N9sCRBYF7Z3blVqQt-yxptaX61vXWLb9-eyHGlgYphe6p6I3pX6fGXcR38ZV9rvEGP9kFuRfcDwEaHUmAV84zGdxzU93VMysXWRMS2b3imSWYXRqxxZQrLhqMQ2X3j8h8q_cDIjMsoZspqDhxj1fJrYCVDkwFW"/>
        <div>
         <div class="font-h3 text-h3 text-text-primary">
          Elena Rostov
         </div>
         <div class="font-metadata text-metadata text-text-secondary">
          e.rostov@technopark.gov.in
         </div>
        </div>
       </div>
      </td>
      <td class="py-4 px-6 font-body text-body text-text-primary">
       Asset Manager
      </td>
      <td class="py-4 px-6 font-body text-body text-text-primary">
       Finance &amp; Procurement
      </td>
      <td class="py-4 px-6">
       <span class="inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full bg-metric-sage text-status-in-stock font-metadata text-metadata font-medium border border-status-in-stock/20">
        <span class="sr-only">
         Status:
        </span>
        <span class="w-1.5 h-1.5 rounded-full bg-status-in-stock">
        </span>
        Active
       </span>
      </td>
      <td class="py-4 px-6 font-mono-data text-mono-data text-text-secondary">
       Oct 24, 11:05
      </td>
      <td class="py-4 px-6 text-right">
       <button class="text-text-secondary hover:text-text-primary p-1 rounded transition-colors opacity-0 group-hover:opacity-100 focus:opacity-100">
        <span class="material-symbols-outlined">
         more_vert
        </span>
       </button>
      </td>
     </tr>
    </tbody>
   </table>
  </div>
  <!-- Pagination Footer -->
  <div class="bg-surface-subtle border-t border-border-default p-4 flex items-center justify-between">
   <p class="font-metadata text-metadata text-text-secondary">
    Showing
    <span class="font-medium text-text-primary">
     1
    </span>
    to
    <span class="font-medium text-text-primary">
     5
    </span>
    of
    <span class="font-medium text-text-primary">
     104
    </span>
    users
   </p>
   <div class="flex items-center gap-2">
    <button class="p-1.5 border border-border-default rounded bg-surface text-text-secondary hover:bg-canvas disabled:opacity-50" disabled="">
     <span class="material-symbols-outlined" style="font-size: 20px;">
      chevron_left
     </span>
    </button>
    <button class="p-1.5 border border-border-default rounded bg-surface text-text-primary hover:bg-canvas">
     <span class="material-symbols-outlined" style="font-size: 20px;">
      chevron_right
     </span>
    </button>
   </div>
  </div>
 </div>
</main>

</template>

<script setup lang="ts">
// Autogenerated from user_access_role_control
</script>
