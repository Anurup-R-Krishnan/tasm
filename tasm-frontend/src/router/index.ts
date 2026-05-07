import { createRouter, createWebHistory } from 'vue-router';
import AppLayout from '../components/AppLayout.vue';
import { appRoutes } from './routes';
import { useAuth } from '../composables/useAuth';

const childRoutes = appRoutes.map((route) => ({
  path: route.path,
  name: route.name,
  meta: { title: route.title },
  component: route.component,
}));

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/first-run',
      name: 'FirstRun',
      component: () => import('../views/FirstRunSetup.vue'),
      meta: { requiresAuth: false },
    },
    {
      path: '/login',
      name: 'Login',
      component: () => import('../views/LoginView.vue'),
      meta: { requiresAuth: false },
    },
    {
      path: '/setup',
      name: 'Setup',
      component: () => import('../views/SetupWizard.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/',
      component: AppLayout,
      meta: { requiresAuth: true },
      children: childRoutes,
    },
  ],
});

router.beforeEach(async (to) => {
  const { initAuth, isInitializing, isAuthenticated, isSetupCompleted, isFirstRun } = useAuth();

  if (isInitializing.value) {
    await initAuth();
  }

  // 1. Fresh install — no users exist at all → go to first-run setup
  if (isFirstRun.value && to.name !== 'FirstRun') {
    return { name: 'FirstRun' };
  }

  // 2. First-run is done but stay on first-run is invalid
  if (!isFirstRun.value && to.name === 'FirstRun') {
    return isAuthenticated.value ? { name: 'Setup' } : { name: 'Login' };
  }

  const requiresAuth = to.matched.some((record) => record.meta['requiresAuth']);

  // 3. Protected route but not authenticated → login
  if (requiresAuth && !isAuthenticated.value) {
    return { name: 'Login' };
  }

  // 4. Authenticated but org setup not done → setup wizard
  if (isAuthenticated.value && !isSetupCompleted.value && to.name !== 'Setup') {
    return { name: 'Setup' };
  }

  // 5. Org setup done, don't show wizard again
  if (isAuthenticated.value && isSetupCompleted.value && to.name === 'Setup') {
    return { name: 'Dashboard' };
  }

  // 6. Already logged in, at login → dashboard or setup
  if (to.name === 'Login' && isAuthenticated.value) {
    return isSetupCompleted.value ? { name: 'Dashboard' } : { name: 'Setup' };
  }

  return true;
});

export default router;
