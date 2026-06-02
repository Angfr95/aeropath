import { useState, useEffect } from "react";
import { useNavigate, useLocation } from "react-router-dom";
import { useAuthStore } from "@/store/authStore";
import { useUIStore } from "@/store/uiStore";

/** Page de connexion / inscription (renderLogin) */
export default function LoginPage() {
  const navigate = useNavigate();
  const location = useLocation();
  const login = useAuthStore((s) => s.login);
  const register = useAuthStore((s) => s.register);
  const isAuthenticated = useAuthStore((s) => s.isAuthenticated);
  const addToast = useUIStore((s) => s.addToast);

  const [tab, setTab] = useState<"login" | "register">(
    location.hash === "#register" ? "register" : "login",
  );
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [loading, setLoading] = useState(false);

  // Redirige si déjà connecté
  useEffect(() => {
    if (isAuthenticated) {
      navigate("/", { replace: true });
    }
  }, [isAuthenticated, navigate]);

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    if (!email || !password) {
      addToast("Veuillez remplir tous les champs", "error");
      return;
    }
    setLoading(true);
    try {
      if (tab === "login") {
        await login({ email, password });
        addToast("Connecté !", "success");
      } else {
        await register({ email, password });
        addToast("Compte créé !", "success");
      }
      navigate("/", { replace: true });
    } catch {
      addToast(tab === "login" ? "Erreur de connexion" : "Erreur d'inscription", "error");
    } finally {
      setLoading(false);
    }
  }

  return (
    <div className="min-h-screen bg-slate-900 flex items-center justify-center p-4">
      <div className="w-full max-w-md">
        <div className="text-center mb-8">
          <div className="text-5xl mb-4">🛩️</div>
          <h1 className="text-3xl font-bold text-white">AeroPath</h1>
          <p className="text-slate-400 mt-2">Formation aéronautique</p>
        </div>

        <div className="bg-slate-800 rounded-xl p-6 shadow-xl">
          <div className="flex mb-6">
            <button
              onClick={() => setTab("login")}
              className={`flex-1 py-2 text-center font-medium rounded-l-lg transition ${
                tab === "login"
                  ? "bg-blue-600 text-white"
                  : "bg-slate-700 text-slate-300"
              }`}
            >
              Connexion
            </button>
            <button
              onClick={() => setTab("register")}
              className={`flex-1 py-2 text-center font-medium rounded-r-lg transition ${
                tab === "register"
                  ? "bg-blue-600 text-white"
                  : "bg-slate-700 text-slate-300"
              }`}
            >
              Inscription
            </button>
          </div>

          <form onSubmit={handleSubmit}>
            <div className="mb-4">
              <label className="block text-sm text-slate-400 mb-1">Email</label>
              <input
                type="email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                className="w-full bg-slate-700 text-white rounded-lg px-4 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
                placeholder="pilote@example.com"
                required
              />
            </div>
            <div className="mb-6">
              <label className="block text-sm text-slate-400 mb-1">Mot de passe</label>
              <input
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                className="w-full bg-slate-700 text-white rounded-lg px-4 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
                placeholder="••••••••"
                required
                minLength={8}
              />
            </div>
            <button
              type="submit"
              disabled={loading}
              className="w-full bg-blue-600 hover:bg-blue-700 disabled:opacity-50 text-white font-medium py-2 rounded-lg transition"
            >
              {loading
                ? "Chargement..."
                : tab === "login"
                ? "Se connecter"
                : "Créer un compte"}
            </button>
          </form>
        </div>

        <p className="text-center text-slate-500 text-sm mt-4">
          🛩️ Apprenez où que vous soyez, même sans connexion
        </p>
      </div>
    </div>
  );
}