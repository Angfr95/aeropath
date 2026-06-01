import { useEffect, useState } from "react";
import { Link, useParams } from "react-router-dom";
import { lessonsApi } from "@/api/lessons";
import { LICENSES, CATEGORIES } from "@/constants/licenses";
import type { Lesson } from "@/types";
import Navbar from "@/components/Navbar";

/** renderLessonsByLicense + loadLessonsByLicense */
export default function LessonsByLicense() {
  const { license: licenseId } = useParams<{ license: string }>();
  const license = LICENSES.find((l) => l.id === licenseId) ?? { id: licenseId ?? "?", label: licenseId ?? "?", icon: "📋", desc: "" };
  const [lessons, setLessons] = useState<Lesson[]>([]);
  const [loading, setLoading] = useState(true);
  const [catCounts, setCatCounts] = useState<Record<string, number>>({});

  useEffect(() => {
    if (!licenseId) return;
    lessonsApi.byLicense(licenseId).then((res) => {
      const data: Lesson[] = Array.isArray(res.data) ? res.data : res.data.data ?? [];
      setLessons(data);
      const counts: Record<string, number> = {};
      CATEGORIES.forEach((c) => {
        counts[c.id] = data.filter((l) => l.category === c.id).length;
      });
      setCatCounts(counts);
    }).finally(() => setLoading(false));
  }, [licenseId]);

  return (
    <div>
      <Navbar />
      <div className="max-w-6xl mx-auto p-4">
        <Link to="/lessons" className="text-slate-400 hover:text-white mb-4 flex items-center gap-1">
          ← Retour aux licences
        </Link>
        <div className="flex items-center gap-3 mb-6">
          <span className="text-4xl">{license.icon}</span>
          <div>
            <h1 className="text-2xl font-bold text-white">{license.label}</h1>
            <p className="text-slate-400">{license.desc}</p>
          </div>
        </div>
        <h2 className="text-lg font-semibold text-slate-300 mb-3">Choisis une catégorie</h2>
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-3 mb-8">
          {CATEGORIES.map((c) => (
            <Link
              key={c.id}
              to={`/lessons/${licenseId}/${c.id}`}
              className="bg-slate-800 hover:bg-slate-700 rounded-xl p-4 text-left transition border border-slate-700 hover:border-green-500/50"
            >
              <div className="flex items-center gap-3">
                <span className="text-2xl">{c.icon}</span>
                <div>
                  <div className="text-white font-medium">{c.label}</div>
                  <div className="text-xs text-slate-500">{loading ? "Chargement..." : `${catCounts[c.id] ?? 0} leçons`}</div>
                </div>
              </div>
            </Link>
          ))}
        </div>
        <h2 className="text-lg font-semibold text-slate-300 mb-3">Toutes les leçons {license.label}</h2>
        {loading ? (
          <div className="space-y-3">{[1, 2].map((i) => <div key={i} className="skeleton h-20 rounded-xl" />)}</div>
        ) : lessons.length === 0 ? (
          <p className="text-slate-400 text-center py-8">Aucune leçon pour cette licence</p>
        ) : (
          <div className="space-y-3">
            {lessons.map((l) => (
              <Link key={l.id} to={`/lessons/${l.license}/${l.category}/${l.id}`} className="block bg-slate-800 rounded-xl p-4 hover:bg-slate-750 transition">
                <h3 className="text-white font-medium">{l.title_fr || l.title_en}</h3>
                <div className="flex gap-2 mt-2">
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