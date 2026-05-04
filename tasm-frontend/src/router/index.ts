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
      path: '/',
      component: AppLayout,
      meta: { requiresAuth: true },
      children: childRoutes,
    },
  ],
});

router.beforeEach(async (to) => {
  const { initAuth, isInitializing } = useAuth();

  // If initializing, wait for it (though initAuth should be called once at app level)
  // For safety, we can ensure it's called or check state.
  if (isInitializing.value) {
    await initAuth();
  }

  const { isAuthenticated } = useAuth();
  const requiresAuth = to.matched.some((record) => record.meta['requiresAuth']);

  if (requiresAuth && !isAuthenticated.value) {
    return { name: 'Login' };
  }

  if (to.name === 'Login' && isAuthenticated) {
    return { name: 'Dashboard' };
  }

  return true;
});

export default router;
