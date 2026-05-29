import { Routes, Route, Link, useLocation } from "react-router-dom";
import { useTranslation } from "react-i18next";
import {
  LayoutDashboard,
  BookOpen,
  HelpCircle,
  BarChart3,
  History,
  GraduationCap,
  Sparkles,
} from "lucide-react";
import Dashboard from "./pages/Dashboard";
import Questions from "./pages/Questions";
import Quiz from "./pages/Quiz";
import Lessons from "./pages/Lessons";
import HistoryPage from "./pages/HistoryPage";
import Stats from "./pages/Stats";
import Recommendations from "./pages/Recommendations";
import LanguageSwitcher from "./components/LanguageSwitcher";

export default function App() {
  const { t } = useTranslation();
  const location = useLocation();

  const navItems = [
    { path: "/", label: t("nav.dashboard"), icon: LayoutDashboard },
    { path: "/questions", label: t("nav.questions"), icon: HelpCircle },
    { path: "/quiz", label: t("nav.quiz"), icon: GraduationCap },
    { path: "/lessons", label: t("nav.lessons"), icon: BookOpen },
    { path: "/history", label: t("nav.history"), icon: History },
    { path: "/stats", label: t("nav.stats"), icon: BarChart3 },
    { path: "/recommendations", label: t("nav.recommendations"), icon: Sparkles },
  ];

  return (
    <div className="min-h-screen bg-slate-950 flex">
      {/* Sidebar */}
      <aside className="w-64 bg-slate-900 border-r border-slate-800 hidden lg:flex flex-col">
        <div className="p-6 border-b border-slate-800">
          <div className="flex items-center gap-3">
            <div className="w-10 h-10 bg-gradient-to-br from-blue-500 to-purple-600 rounded-xl flex items-center justify-center">
              <GraduationCap className="w-6 h-6 text-white" />
            </div>
            <div>
              <h1 className="text-lg font-bold text-white">AeroForge</h1>
              <p className="text-xs text-slate-400">Apprentissage adaptatif</p>
            </div>
          </div>
        </div>

        <nav className="flex-1 p-4 space-y-1">
          {navItems.map((item) => {
            const isActive = location.pathname === item.path;
            const Icon = item.icon;
            return (
              <Link
                key={item.path}
                to={item.path}
                className={`flex items-center gap-3 px-4 py-3 rounded-lg text-sm font-medium transition-all duration-200 ${
                  isActive
                    ? "bg-blue-600/20 text-blue-400 border border-blue-500/30"
                    : "text-slate-400 hover:text-slate-200 hover:bg-slate-800"
                }`}
              >
                <Icon className="w-5 h-5" />
                {item.label}
              </Link>
            );
          })}
        </nav>

        <div className="p-4 border-t border-slate-800 space-y-2">
          <div className="flex items-center gap-3 px-4 py-3 rounded-lg bg-slate-800/50">
            <div className="w-8 h-8 bg-blue-600 rounded-full flex items-center justify-center text-xs font-bold">
              U
            </div>
            <div className="flex-1 min-w-0">
              <p className="text-sm font-medium text-slate-200 truncate">
                Utilisateur
              </p>
              <p className="text-xs text-slate-500">Connecté</p>
            </div>
          </div>
          <LanguageSwitcher />
        </div>
      </aside>

      {/* Mobile Nav */}
      <nav className="lg:hidden fixed bottom-0 left-0 right-0 bg-slate-900 border-t border-slate-800 z-50">
        <div className="flex overflow-x-auto">
          {navItems.map((item) => {
            const isActive = location.pathname === item.path;
            const Icon = item.icon;
            return (
              <Link
                key={item.path}
                to={item.path}
                className={`flex flex-col items-center gap-1 px-3 py-2 text-xs font-medium min-w-fit ${
                  isActive ? "text-blue-400" : "text-slate-500"
                }`}
              >
                <Icon className="w-5 h-5" />
                {item.label}
              </Link>
            );
          })}
        </div>
      </nav>

      {/* Main Content */}
      <main className="flex-1 overflow-auto pb-20 lg:pb-0">
        <div className="max-w-7xl mx-auto p-6">
          <Routes>
            <Route path="/" element={<Dashboard />} />
            <Route path="/questions" element={<Questions />} />
            <Route path="/quiz" element={<Quiz />} />
            <Route path="/lessons" element={<Lessons />} />
            <Route path="/history" element={<HistoryPage />} />
            <Route path="/stats" element={<Stats />} />
            <Route path="/recommendations" element={<Recommendations />} />
          </Routes>
        </div>
      </main>
    </div>
  );
}
