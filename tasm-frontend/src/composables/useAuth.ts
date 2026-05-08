import { ref, computed } from 'vue';
import {
  login as apiLogin,
  register as apiRegister,
  getMe,
  type LoginCredentials,
  type RegisterCredentials,
  type AuthResponse,
} from '../api/auth';
import type { SystemUser } from '../types/models';

const TOKEN_KEY = 'tasm_auth_token';
const isE2E = import.meta.env['VITE_E2E'] === 'true';

// Global shared state
const currentUser = ref<SystemUser | null>(null);
const token = ref<string | null>(localStorage.getItem(TOKEN_KEY));
const isInitializing = ref(true);
const isSetupCompleted = ref(false);
const isFirstRun = ref(false); // True when no users exist at all (fresh install)
const companyName = ref('');
let authPromise: Promise<void> | null = null;

export function useAuth() {
  const isAuthenticated = computed(() => !!token.value);

  const checkSetupStatus = async () => {
    try {
      const response = await fetch('/api/setup/status');
      const data = await response.json();
      isFirstRun.value = !!data.isFirstRun;
      isSetupCompleted.value = !!data.isSetupCompleted;
      companyName.value = data.companyName || '';
    } catch (error) {
      console.error('Failed to check setup status:', error);
    }
  };

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
      const response: AuthResponse = await apiLogin(credentials);
      setToken(response.token);
      currentUser.value = response.user;
      await checkSetupStatus();
      return true;
    } catch (error) {
      console.error('Login failed:', error);
      throw error;
    }
  };

  const register = async (credentials: RegisterCredentials) => {
    try {
      const response: AuthResponse = await apiRegister(credentials);
      setToken(response.token);
      currentUser.value = response.user;
      await checkSetupStatus();
      return true;
    } catch (error) {
      console.error('Registration failed:', error);
      throw error;
    }
  };

  const logout = () => {
    setToken(null);
    currentUser.value = null;
    isFirstRun.value = false;
    isSetupCompleted.value = false;
    if (typeof window !== 'undefined') {
      window.location.href = '/login';
    }
  };

  const initAuth = async () => {
    if (authPromise) return authPromise;

    authPromise = (async () => {
      isInitializing.value = true;
      if (isE2E) {
        if (!token.value) {
          setToken('e2e-dummy-token');
        }
        isInitializing.value = false;
        return;
      }

      // Always check setup status first — this tells us if it's a first run
      await checkSetupStatus();

      if (token.value) {
        try {
          const user = await getMe();
          currentUser.value = user;
        } catch (error) {
          console.error('Session restoration failed:', error);
          setToken(null);
          currentUser.value = null;
        }
      }
      isInitializing.value = false;
    })();

    return authPromise;
  };

  return {
    currentUser,
    token,
    isAuthenticated,
    isInitializing,
    isSetupCompleted,
    isFirstRun,
    companyName,
    login,
    register,
    logout,
    initAuth,
    checkSetupStatus,
    setToken,
  };
}
