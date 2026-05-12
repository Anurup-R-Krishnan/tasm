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
      path: '/welcome',
      name: 'Welcome',
      component: () => import('../views/WelcomeView.vue'),
      meta: { requiresAuth: false },
    },
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
      path: '/register',
      name: 'Register',
      component: () => import('../views/RegisterView.vue'),
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
  const { initAuth, isInitializing, isAuthenticated, isSetupCompleted } = useAuth();

  if (isInitializing.value) {
    await initAuth();
  }

  const isAuthPath = to.name === 'Login' || to.path.startsWith('/login');
  const isRegisterPath = to.name === 'Register' || to.path.startsWith('/register');
  const isWelcomePath = to.name === 'Welcome' || to.path.startsWith('/welcome');
  const isFirstRunPath = to.name === 'FirstRun' || to.path.startsWith('/first-run');

  // /first-run: requires authentication and is used when setup is incomplete.
  if (isFirstRunPath) {
    if (!isAuthenticated.value) return { name: 'Register' };
    return true;
  }

  // Unauthenticated user hitting a protected route → Welcome
  const requiresAuth = to.matched.some((record) => record.meta['requiresAuth']);
  if (requiresAuth && !isAuthenticated.value) return { name: 'Welcome' };

  // Authenticated user with incomplete setup should always go to the setup wizard
  if (isAuthenticated.value && !isSetupCompleted.value && !isFirstRunPath) {
    return { name: 'FirstRun' };
  }

  // Authenticated + setup done: bounce away from auth/welcome pages to the app
  if (
    isAuthenticated.value &&
    isSetupCompleted.value &&
    (isAuthPath || isRegisterPath || isWelcomePath)
  ) {
    return { name: 'Dashboard' };
  }

  return true;
});

export default router;
