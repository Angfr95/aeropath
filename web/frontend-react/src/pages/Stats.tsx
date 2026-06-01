import { useEffect, useState } from "react";
import { statsApi } from "@/api/stats";
import type { UserStats, AdminStats } from "@/types";
import Navbar from "@/components/Navbar";

/** renderStats + loadStats */
export default function Stats() {
  const [userStats, setUserStats] = useState<UserStats | null>(null);
  const [adminStats, setAdminStats] = useState<AdminStats | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    (async () => {
      const [u, a] = await Promise.all([
        statsApi.user().catch(() => null),
        statsApi.admin().catch(() => null),
      ]);
      if (u) setUserStats(u.data);
      if (a) setAdminStats(a.data);
      setLoading(false);
    })();
  }, []);

  if (loading) {
    return (
      <div>
        <Navbar />
        <div className="max-w-4xl mx-auto p-4">
          <h1 className="text-2xl font-bold text-white mb-4">Statistiques</h1>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            {[1, 2].map((i) => <div key={i} className="skeleton h-48 rounded-xl" />)}
          </div>
        </div>
      </div>
    );
  }

  return (
    <div>
      <Navbar />
      <div className="max-w-4xl mx-auto p-4">
        <h1 className="text-2xl font-bold text-white mb-4">Statistiques</h1>
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div className="bg-slate-800 rounded-xl p-4">
            <h3 className="text-white font-medium mb-3">📊 Mes statistiques</h3>
            {userStats ? (
              <div className="space-y-2 text-sm">
                {Object.entries(userStats).map(([k, v]) => (
                  <div key={k} className="flex justify-between">
                    <span className="text-slate-400">{k}</span>
                    <span className="text-white">{typeof v === "number" ? v : String(v)}</span>
                  </div>
                ))}
              </div>
            ) : (
              <p className="text-slate-400">Non disponible</p>
            )}
          </div>
          <div className="bg-slate-800 rounded-xl p-4">
            <h3 className="text-white font-medium mb-3">🌍 Global</h3>
            {adminStats ? (
              <div className="space-y-2 text-sm">
                <div className="flex justify-between">
                  <span className="text-slate-400">Étudiants</span>
                  <span className="text-white">{adminStats.total_students ?? 0}</span>
                </div>
                <div className="flex justify-between">
                  <span className="text-slate-400">Questions</span>
                  <span className="text-white">{adminStats.total_questions ?? 0}</span>
                </div>
                <div className="flex justify-between">
                  <span className="text-slate-400">Réponses</span>
                  <span className="text-white">{adminStats.total_answers ?? 0}</span>
                </div>
              </div>
            ) : (
              <p className="text-slate-400">Non disponible</p>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}