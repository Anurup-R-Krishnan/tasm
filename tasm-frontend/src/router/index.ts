import { createRouter, createWebHistory } from 'vue-router';
import AppLayout from '../components/AppLayout.vue';
import { appRoutes } from './routes';

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

router.beforeEach((to, _from, next) => {
  const token = localStorage.getItem('tasm_auth_token');
  const requiresAuth = to.matched.some((record) => record.meta['requiresAuth']);

  if (requiresAuth && !token) {
    next({ name: 'Login' });
  } else if (to.name === 'Login' && token) {
    next({ name: 'Dashboard' }); // Or wherever you want logged-in users to go
  } else {
    next();
  }
});

export default router;
