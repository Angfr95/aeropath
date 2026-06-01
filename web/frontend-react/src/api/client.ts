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
      // On nettoie le store auth → le ProtectedRoute redirigera
      // L'import dynamique évite la dépendance circulaire
      import("@/store/authStore").then(({ useAuthStore }) => {
        useAuthStore.getState().logout();
      });
    }
    return Promise.reject(error);
  },
);