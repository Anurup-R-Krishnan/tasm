import { ref, computed } from 'vue';
import { login as apiLogin, getMe, type LoginCredentials } from '../api/auth';
import { useRouter } from 'vue-router';
import type { SystemUser } from '../types/models';

const TOKEN_KEY = 'tasm_auth_token';
const isE2E = import.meta.env['VITE_E2E'] === 'true';

// Global shared state
const currentUser = ref<SystemUser | null>(null);
const token = ref<string | null>(localStorage.getItem(TOKEN_KEY));
const isInitializing = ref(true);

export function useAuth() {
  const router = useRouter();

  const isAuthenticated = computed(() => !!token.value);

  const setToken = (newToken: string | null) => {
    token.value = newToken;
    if (newToken) {
      localStorage.setItem(TOKEN_KEY, newToken);
    } else {
      localStorage.removeItem(TOKEN_KEY);
    }
  };

  const login = async (credentials: LoginCredentials) => {
    try {
      const response = await apiLogin(credentials);
      setToken(response.token);
      currentUser.value = response.user;
      return true;
    } catch (error) {
      console.error('Login failed:', error);
      throw error;
    }
  };

  const logout = () => {
    setToken(null);
    currentUser.value = null;
    router.push({ name: 'Login' });
  };

  const initAuth = async () => {
    isInitializing.value = true;
    if (token.value) {
      if (isE2E) {
        isInitializing.value = false;
        return;
      }
      try {
        const user = await getMe();
        currentUser.value = user;
      } catch (error) {
        console.error('Session restoration failed:', error);
        logout(); // Token invalid or expired
      }
    }
    isInitializing.value = false;
  };

  return {
    currentUser,
    token,
    isAuthenticated,
    isInitializing,
    login,
    logout,
    initAuth,
  };
}
