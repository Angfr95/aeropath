import { useEffect, useState } from "react";
import { recommendationsApi } from "@/api/recommendations";
import type { Recommendations as Recs } from "@/types";
import Navbar from "@/components/Navbar";

/** renderRecommendations + loadRecommendations */
export default function Recommendations() {
  const [data, setData] = useState<Recs | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    recommendationsApi.get().then((res) => setData(res.data)).finally(() => setLoading(false));
  }, []);

  if (loading) {
    return (
      <div>
        <Navbar />
        <div className="max-w-4xl mx-auto p-4">
          <h1 className="text-2xl font-bold text-white mb-4">Recommandations</h1>
          <div className="space-y-4">{[1, 2, 3].map((i) => <div key={i} className="skeleton h-32 rounded-xl" />)}</div>
        </div>
      </div>
    );
  }

  if (!data) {
    return (
      <div>
        <Navbar />
        <div className="max-w-4xl mx-auto p-4">
          <h1 className="text-2xl font-bold text-white mb-4">Recommandations</h1>
          <p className="text-slate-400">Non disponible hors-ligne</p>
        </div>
      </div>
    );
  }

  return (
    <div>
      <Navbar />
      <div className="max-w-4xl mx-auto p-4">
        <h1 className="text-2xl font-bold text-white mb-4">Recommandations</h1>
        <div className="space-y-4">
          {/* Progression */}
          <div className="bg-slate-800 rounded-xl p-4">
            <h3 className="text-white font-medium mb-3">📈 Progression</h3>
            <div className="flex justify-between text-sm mb-1">
              <span className="text-slate-400">Globale</span>
              <span className="text-white font-medium">{Math.round(data.progression_percent ?? 0)}%</span>
            </div>
            <div className="w-full bg-slate-700 rounded-full h-3">
              <div className="bg-blue-500 rounded-full h-3 transition-all" style={{ width: `${Math.round(data.progression_percent ?? 0)}%` }} />
            </div>
            <p className="text-sm text-slate-500 mt-2">Prochain palier : {data.next_milestone ?? "-"}</p>
          </div>

          {/* Sujets faibles */}
          <div className="bg-slate-800 rounded-xl p-4">
            <h3 className="text-white font-medium mb-3">🧠 Sujets à travailler</h3>
            {(data.weak_topics ?? []).length > 0 ? (
              <div className="space-y-2">
                {data.weak_topics!.map((t, i) => (
                  <div key={i} className="flex justify-between items-center">
                    <span className="text-slate-300">{t.theme}</span>
                    <span className="text-red-400 text-sm">{Math.round(t.score ?? 0)}%</span>
                  </div>
                ))}
              </div>
            ) : (
              <p className="text-slate-400">Aucun sujet faible détecté</p>
            )}
          </div>

          {/* Cartes à réviser */}
          <div className="bg-slate-800 rounded-xl p-4">
            <h3 className="text-white font-medium mb-3">📖 Cartes à réviser</h3>
            {(data.due_cards ?? []).length > 0 ? (
              <div className="space-y-2">
                {data.due_cards!.slice(0, 10).map((c) => (
                  <div key={c.question_id} className="bg-slate-700 rounded-lg p-3">
                    <p className="text-white text-sm">{c.question_fr || c.question_en || "Question"}</p>
                    <p className="text-xs text-slate-500 mt-1">Prochaine révision : {c.next_review ?? "bientôt"}</p>
                  </div>
                ))}
                {data.due_cards!.length > 10 && (
                  <p className="text-sm text-slate-500 mt-2">Et {data.due_cards!.length - 10} autres...</p>
                )}
              </div>
            ) : (
              <p className="text-slate-400">Tout est à jour ! 🎉</p>
            )}
          </div>

          {/* Maîtrise par licence */}
          <div className="bg-slate-800 rounded-xl p-4">
            <h3 className="text-white font-medium mb-3">🎯 Maîtrise par licence</h3>
            {(data.mastery_by_license ?? []).length > 0 ? (
              <div className="space-y-2">
                {data.mastery_by_license!.map((m) => (
                  <div key={m.license}>
                    <div className="flex justify-between text-sm mb-1">
                      <span className="text-slate-300">{m.license}</span>
                      <span className="text-white">{Math.round(m.score ?? 0)}%</span>
                    </div>
                    <div className="w-full bg-slate-700 rounded-full h-2">
                      <div className="bg-green-500 rounded-full h-2" style={{ width: `${Math.round(m.score ?? 0)}%` }} />
                    </div>
                  </div>
                ))}
              </div>
            ) : (
              <p className="text-slate-400">Pas encore de données</p>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}