import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { useAuthStore } from "@/store/authStore";
import { useQuizStore } from "@/store/quizStore";
import { questionsApi } from "@/api/questions";
import { lessonsApi } from "@/api/lessons";
import { statsApi } from "@/api/stats";
import { recommendationsApi } from "@/api/recommendations";
import type { AdminStats, Recommendations } from "@/types";
import Navbar from "@/components/Navbar";

/** Dashboard authentifié — renderDashboard + loadDashboardData */
export default function Dashboard() {
  const user = useAuthStore((s) => s.user);
  const startRandom = useQuizStore((s) => s.startRandom);
  const [qCount, setQCount] = useState<number | null>(null);
  const [lCount, setLCount] = useState<number | null>(null);
  const [adminStats, setAdminStats] = useState<AdminStats | null>(null);
  const [recs, setRecs] = useState<Recommendations | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    (async () => {
      try {
        const [qC, lC, aS, r] = await Promise.all([
          questionsApi.count().catch(() => null),
          lessonsApi.count().catch(() => null),
          statsApi.admin().catch(() => null),
          recommendationsApi.get().catch(() => null),
        ]);
        if (qC) setQCount(qC.data.count);
        if (lC) setLCount(lC.data.count);
        if (aS) setAdminStats(aS.data);
        if (r) setRecs(r.data);
      } finally {
        setLoading(false);
      }
    })();
  }, []);

  async function handleRandomQuiz() {
    await startRandom(5);
  }

  return (
    <div className="min-h-screen bg-slate-900">
      <Navbar />
      <div className="max-w-6xl mx-auto p-4">
        <div className="mb-6">
          <h1 className="text-2xl font-bold text-white">
            Bonjour {user?.email ?? "pilote"} 👋
          </h1>
          <p className="text-slate-400">Prêt à réviser ?</p>
        </div>

        {loading ? (
          <div className="grid grid-cols-1 md:grid-cols-3 gap-4 mb-6">
            {[1, 2, 3].map((i) => (
              <div key={i} className="bg-slate-800 rounded-xl p-4 skeleton h-24" />
            ))}
          </div>
        ) : (
          <div className="grid grid-cols-1 md:grid-cols-3 gap-4 mb-6">
            <div className="bg-slate-800 rounded-xl p-4">
              <div className="text-3xl mb-2">📚</div>
              <div className="text-2xl font-bold text-white">{qCount ?? "-"}</div>
              <div className="text-sm text-slate-400">Questions</div>
            </div>
            <div className="bg-slate-800 rounded-xl p-4">
              <div className="text-3xl mb-2">📓</div>
              <div className="text-2xl font-bold text-white">{lCount ?? "-"}</div>
              <div className="text-sm text-slate-400">Leçons</div>
            </div>
            <div className="bg-slate-800 rounded-xl p-4">
              <div className="text-3xl mb-2">✅</div>
              <div className="text-2xl font-bold text-white">{adminStats?.total_answers ?? "-"}</div>
              <div className="text-sm text-slate-400">Réponses</div>
            </div>
          </div>
        )}

        {/* Actions rapides */}
        <div className="bg-slate-800 rounded-xl p-4 mb-4">
          <h2 className="text-lg font-bold text-white mb-3">Actions rapides</h2>
          <div className="grid grid-cols-2 md:grid-cols-4 gap-3">
            <Link
              to="/quiz"
              onClick={(e) => { e.preventDefault(); handleRandomQuiz().then(() => { window.location.href = "/quiz"; }); }}
              className="bg-blue-600 hover:bg-blue-700 text-white rounded-lg p-3 text-center transition"
            >
              <div className="text-2xl mb-1">🧠</div>
              <div className="text-sm">Question aléatoire</div>
            </Link>
            <Link to="/questions" className="bg-purple-600 hover:bg-purple-700 text-white rounded-lg p-3 text-center transition">
              <div className="text-2xl mb-1">📋</div>
              <div className="text-sm">Toutes les questions</div>
            </Link>
            <Link to="/lessons" className="bg-green-600 hover:bg-green-700 text-white rounded-lg p-3 text-center transition">
              <div className="text-2xl mb-1">📓</div>
              <div className="text-sm">Leçons</div>
            </Link>
            <Link to="/recommendations" className="bg-amber-600 hover:bg-amber-700 text-white rounded-lg p-3 text-center transition">
              <div className="text-2xl mb-1">🧠</div>
              <div className="text-sm">Recommandations</div>
            </Link>
          </div>
        </div>

        {/* Recommandations */}
        <div className="bg-slate-800 rounded-xl p-4">
          <h2 className="text-lg font-bold text-white mb-3">Recommandations</h2>
          {loading ? (
            <p className="text-slate-400">Chargement...</p>
          ) : recs ? (
            <div className="space-y-2">
              <div className="flex justify-between text-sm">
                <span className="text-slate-400">Progression</span>
                <span className="text-white font-medium">{Math.round(recs.progression_percent ?? 0)}%</span>
              </div>
              <div className="w-full bg-slate-700 rounded-full h-2">
                <div className="bg-blue-500 rounded-full h-2" style={{ width: `${Math.round(recs.progression_percent ?? 0)}%` }} />
              </div>
              <div className="flex justify-between text-sm mt-3">
                <span className="text-slate-400">Sujets faibles</span>
                <span className="text-red-400 font-medium">{recs.weak_topics?.length ?? 0}</span>
              </div>
              <div className="flex justify-between text-sm">
                <span className="text-slate-400">Cartes à réviser</span>
                <span className="text-amber-400 font-medium">{recs.due_cards?.length ?? 0}</span>
              </div>
              <div className="flex justify-between text-sm">
                <span className="text-slate-400">Prochain palier</span>
                <span className="text-green-400 font-medium">{recs.next_milestone ?? "-"}</span>
              </div>
            </div>
          ) : (
            <p className="text-slate-400">Non disponible</p>
          )}
        </div>
      </div>
    </div>
  );
}