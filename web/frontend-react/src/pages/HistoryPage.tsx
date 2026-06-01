import { useEffect, useState } from "react";
import { historyApi } from "@/api/history";
import type { HistoryEntry } from "@/types";
import Navbar from "@/components/Navbar";

/** renderHistory + loadHistory */
export default function HistoryPage() {
  const [entries, setEntries] = useState<HistoryEntry[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    historyApi.list(100).then((res) => {
      setEntries(res.data.entries ?? []);
    }).finally(() => setLoading(false));
  }, []);

  return (
    <div>
      <Navbar />
      <div className="max-w-4xl mx-auto p-4">
        <h1 className="text-2xl font-bold text-white mb-4">Historique</h1>
        {loading ? (
          <div className="space-y-2">{[1, 2, 3].map((i) => <div key={i} className="skeleton h-14 rounded-lg" />)}</div>
        ) : entries.length === 0 ? (
          <p className="text-slate-400 text-center py-8">Aucun historique</p>
        ) : (
          <div className="space-y-2">
            {entries.map((h) => (
              <div key={h.id} className="bg-slate-800 rounded-lg p-3 flex justify-between items-center">
                <div className="flex-1 min-w-0">
                  <p className="text-white text-sm truncate">{h.question_fr || h.question_en}</p>
                  <div className="flex gap-2 mt-1">
                    <span className="text-xs text-slate-500">{h.theme}</span>
                    <span className="text-xs text-slate-500">Niv. {h.difficulty}</span>
                  </div>
                </div>
                <span className={`ml-3 font-medium text-sm ${h.was_correct ? "text-green-400" : "text-red-400"}`}>
                  {h.was_correct ? "✅" : "❌"}
                </span>
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
}