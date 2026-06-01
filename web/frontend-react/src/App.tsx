import { useEffect } from "react";
import { useAuthStore } from "@/store/authStore";
import AppRouter from "@/router";
import ToastContainer from "@/components/ToastContainer";

/**
 * Composant racine.
 * - Restaure la session JWT (httpOnly cookie) via loadUser()
 * - Monte le routeur (public + protégé)
 * - Affiche les toasts globaux
 */
export default function App() {
  const loadUser = useAuthStore((s) => s.loadUser);
  const isAuthenticated = useAuthStore((s) => s.isAuthenticated);
  const isLoading = useAuthStore((s) => s.isLoading);

  // Restauration de session au montage
  useEffect(() => {
    if (!isAuthenticated && !isLoading) {
      loadUser();
    }
  }, []); // eslint-disable-line react-hooks/exhaustive-deps

  return (
    <div className="min-h-screen bg-slate-900 text-slate-100">
      <AppRouter />
      <ToastContainer />
    </div>
  );
}