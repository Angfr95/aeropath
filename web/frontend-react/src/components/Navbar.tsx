import MeasurementSelector from "./MeasurementSelector";
import { useLearningLocale } from "@/utils/useLearningLocale";
import { Link, useLocation, useNavigate } from "react-router-dom";
import { useAuthStore } from "@/store/authStore";
import { useUIStore } from "@/store/uiStore";
import LanguageSelector from "./LanguageSelector";
import { Plane, LogOut, Menu, X } from "lucide-react";
import { useState } from "react";

const NAV_ITEMS = [
  { path: "/dashboard", label: "Accueil" },
  { path: "/questions", label: "Questions" },
  { path: "/lessons", label: "Leçons" },
  { path: "/quiz", label: "Quiz" },
  { path: "/history", label: "Historique" },
  { path: "/stats", label: "Stats" },
  { path: "/recommendations", label: "Recommandations" },
];

export default function Navbar() {
  const { label } = useLearningLocale();
  const isAuthenticated = useAuthStore((s) => s.isAuthenticated);
  const navItems = isAuthenticated ? NAV_ITEMS : [
    { path: "/", label: "Accueil" },
    ...NAV_ITEMS.filter((item) => ["/questions", "/lessons", "/quiz"].includes(item.path)),
  ];
  const location = useLocation();
  const navigate = useNavigate();
  const logout = useAuthStore((s) => s.logout);
  const addToast = useUIStore((s) => s.addToast);
  const [mobileOpen, setMobileOpen] = useState(false);

  async function handleLogout() {
    await logout();
    addToast(label("Déconnecté"), "info");
    navigate("/");
  }

  return (
    <nav className="bg-slate-800 border-b border-slate-700 sticky top-0 z-40">
      <div className="max-w-6xl mx-auto px-4">
        <div className="flex items-center justify-between h-14">
          {/* Logo */}
          <Link to={isAuthenticated ? "/dashboard" : "/"} className="flex items-center gap-2">
            <span className="text-xl">✈️</span>
            <span className="font-bold text-white">AeroPath</span>
          </Link>

          {/* Desktop links */}
          <div className="hidden xl:flex items-center gap-1">
            {navItems.map((item) => {
              // Comparaison simplifiée : active si le pathname commence par item.path
              const isActive =
                (item.path === "/dashboard" || item.path === "/")
                  ? location.pathname === item.path
                  : location.pathname.startsWith(item.path);
              return (
                <Link
                  key={item.path}
                  to={item.path}
                  className={`px-3 py-1.5 rounded-lg text-sm transition ${
                    isActive
                      ? "bg-blue-600 text-white"
                      : "text-slate-300 hover:bg-slate-700"
                  }`}
                >
                  {label(item.label)}
                </Link>
              );
            })}
            <LanguageSelector />
            <MeasurementSelector />
            {isAuthenticated ? <button
              onClick={handleLogout}
              className="ml-2 px-3 py-1.5 rounded-lg text-sm text-red-400 hover:bg-red-900/30 transition"
            >
              <span className="hidden lg:inline">{label("Déconnexion")}</span>
              <LogOut className="w-4 h-4 lg:hidden" />
            </button> : <Link to="/login" className="ml-2 px-3 py-1.5 rounded-lg text-sm text-blue-300 hover:bg-slate-700">{label("Connexion")}</Link>}
          </div>

          {/* Mobile hamburger */}
          <button
            className="xl:hidden p-2 text-slate-300"
            onClick={() => setMobileOpen((o) => !o)}
            aria-label="Menu"
          >
            {mobileOpen ? <X className="w-5 h-5" /> : <Menu className="w-5 h-5" />}
          </button>
        </div>

        {/* Mobile menu */}
        {mobileOpen && (
          <div className="xl:hidden pb-3 space-y-1 animate-fade-in">
            {navItems.map((item) => {
              const isActive =
                (item.path === "/dashboard" || item.path === "/")
                  ? location.pathname === item.path
                  : location.pathname.startsWith(item.path);
              return (
                <Link
                  key={item.path}
                  to={item.path}
                  onClick={() => setMobileOpen(false)}
                  className={`block px-3 py-2 rounded-lg text-sm transition ${
                    isActive
                      ? "bg-blue-600 text-white"
                      : "text-slate-300 hover:bg-slate-700"
                  }`}
                >
                  {label(item.label)}
                </Link>
              );
            })}
            <div className="flex flex-wrap items-center gap-2 pt-2">
              <LanguageSelector />
            <MeasurementSelector />
              {isAuthenticated ? <button
                onClick={handleLogout}
                className="px-3 py-1.5 rounded-lg text-sm text-red-400 hover:bg-red-900/30 transition"
              >
                {label("Déconnexion")}
              </button> : <Link to="/login" className="px-3 py-1.5 text-sm text-blue-300">{label("Connexion")}</Link>}
            </div>
          </div>
        )}
      </div>
    </nav>
  );
}