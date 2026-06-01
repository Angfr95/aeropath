import { Navigate, Outlet } from "react-router-dom";
import { useAuthStore } from "@/store/authStore";

/**
 * Wraps authenticated routes.
 * - If user is still loading → show spinner
 * - If not authenticated → redirect to /login
 * - Otherwise → render child routes via <Outlet />
 */
export default function ProtectedRoute() {
  const isAuthenticated = useAuthStore((s) => s.isAuthenticated);
  const isLoading = useAuthStore((s) => s.isLoading);

  // Premier montage : on attend que loadUser() ait fini
  if (isLoading) {
    return (
      <div className="min-h-screen flex items-center justify-center bg-slate-900">
        <div className="text-center animate-pulse">
          <div className="text-5xl mb-4">✈️</div>
          <div className="h-6 w-48 skeleton mx-auto mb-3" />
          <div className="h-4 w-32 skeleton mx-auto" />
        </div>
      </div>
    );
  }

  if (!isAuthenticated) {
    return <Navigate to="/login" replace />;
  }

  return <Outlet />;
}