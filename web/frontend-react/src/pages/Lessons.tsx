import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { lessonsApi } from "@/api/lessons";
import { LICENSES } from "@/constants/licenses";
import type { Lesson } from "@/types";
import Navbar from "@/components/Navbar";

/** renderLessons + loadLessons */
export default function Lessons() {
  const [recent, setRecent] = useState<Lesson[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    lessonsApi.list(10).then((res) => {
      const data = Array.isArray(res.data) ? res.data : res.data.data ?? [];
      setRecent(data);
    }).finally(() => setLoading(false));
  }, []);

  return (
    <div>
      <Navbar />
      <div className="max-w-6xl mx-auto p-4">
        <h1 className="text-2xl font-bold text-white mb-6">Leçons</h1>
        <h2 className="text-lg font-semibold text-slate-300 mb-3">Choisis une licence</h2>
        <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-5 gap-3 mb-8">
          {LICENSES.map((l) => (
            <Link
              key={l.id}
              to={`/lessons/${l.id}`}
              className="bg-slate-800 hover:bg-slate-700 rounded-xl p-4 text-center transition border border-slate-700 hover:border-green-500/50"
            >
              <div className="text-3xl mb-2">{l.icon}</div>
              <div className="text-white font-bold">{l.label}</div>
              <div className="text-xs text-slate-400">{l.desc}</div>
            </Link>
          ))}
        </div>
        <h2 className="text-lg font-semibold text-slate-300 mb-3">Dernières leçons</h2>
        {loading ? (
          <div className="space-y-3">{[1, 2].map((i) => <div key={i} className="skeleton h-20 rounded-xl" />)}</div>
        ) : recent.length === 0 ? (
          <p className="text-slate-400 text-center py-8">Aucune leçon trouvée</p>
        ) : (
          <div className="space-y-3">
            {recent.map((l) => (
              <Link
                key={l.id}
                to={`/lessons/${l.license}/${l.category}/${l.id}`}
                className="block bg-slate-800 rounded-xl p-4 hover:bg-slate-750 transition"
              >
                <h3 className="text-white font-medium">{l.title_fr || l.title_en}</h3>
                <div className="flex gap-2 mt-2">
                  <span className="text-xs bg-blue-900 text-blue-300 px-2 py-0.5 rounded">{l.license}</span>
                  <span className="text-xs bg-purple-900 text-purple-300 px-2 py-0.5 rounded">{l.category}</span>
                  <span className="text-xs bg-amber-900 text-amber-300 px-2 py-0.5 rounded">Niv. {l.difficulty}</span>
                </div>
              </Link>
            ))}
          </div>
        )}
      </div>
    </div>
  );
}