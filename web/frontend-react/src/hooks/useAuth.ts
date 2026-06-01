import { useEffect } from "react";
import { useAuthStore } from "@/store/authStore";

/**
 * Hook d'initialisation de l'authentification.
 * Appelle loadUser() au montage si l'utilisateur n'est pas déjà chargé.
 */
export function useAuth() {
  const user = useAuthStore((s) => s.user);
  const isAuthenticated = useAuthStore((s) => s.isAuthenticated);
  const isLoading = useAuthStore((s) => s.isLoading);
  const loadUser = useAuthStore((s) => s.loadUser);
  const logout = useAuthStore((s) => s.logout);

  useEffect(() => {
    // Tente de restaurer la session au montage
    if (!isAuthenticated && !isLoading) {
      loadUser();
    }
  }, [isAuthenticated, isLoading, loadUser]);

  return { user, isAuthenticated, isLoading, logout };
}