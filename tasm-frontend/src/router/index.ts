import { createRouter, createWebHistory } from 'vue-router';
import AppLayout from '../components/AppLayout.vue';
import { appRoutes } from './routes';

import { useAuth } from '../composables/useAuth';

const childRoutes = appRoutes.map((route) => ({
  path: route.path,
  name: route.name,
  meta: {
    title: route.title,
  },
  component: route.component,
}));

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
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
  const { initAuth, isInitializing, isAuthenticated, isSetupCompleted } = useAuth();
  if (isInitializing.value) {
    await initAuth();
  }
  const requiresAuth = to.matched.some((record) => record.meta['requiresAuth']);

  if (requiresAuth && !isAuthenticated.value) {
    return { name: 'Login' };
  }

  if (isAuthenticated.value && !isSetupCompleted.value && to.name !== 'Setup') {
    return { name: 'Setup' };
  }

  if (isAuthenticated.value && isSetupCompleted.value && to.name === 'Setup') {
    return { name: 'Dashboard' };
  }

  if (to.name === 'Login' && isAuthenticated.value) {
    return isSetupCompleted.value ? { name: 'Dashboard' } : { name: 'Setup' };
  }

  return true;
});

export default router;
