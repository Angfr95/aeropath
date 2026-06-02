import axios from "axios";

// ---------------------------------------------------------------------------
// Instance Axios unique pour toute l'application
// ---------------------------------------------------------------------------

export const apiClient = axios.create({
  baseURL: "",          // proxy Vite gère /api et /auth
  withCredentials: true, // httpOnly cookie JWT
  headers: {
    "Content-Type": "application/json",
  },
});

// ---------------------------------------------------------------------------
// Intercepteur 401 : redirige vers /login en cas d'expiration du token
// ---------------------------------------------------------------------------
apiClient.interceptors.response.use(
  (response) => response,
  (error) => {
    if (axios.isAxiosError(error) && error.response?.status === 401) {
      // Nettoie le store LOCALEMENT (pas d'appel API pour éviter les boucles)
      import("@/store/authStore").then(({ useAuthStore }) => {
        // set direct au lieu d'appeler logout() qui ferait un POST /auth/logout
        useAuthStore.setState({ user: null, isAuthenticated: false, isLoading: false });
      });
    }
    return Promise.reject(error);
  },
);
