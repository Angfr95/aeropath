import { useEffect, useState } from "react";
import { Link, useParams } from "react-router-dom";
import { lessonsApi } from "@/api/lessons";
import { LICENSES, CATEGORIES } from "@/constants/licenses";
import type { Lesson } from "@/types";
import Navbar from "@/components/Navbar";

/** renderLessonsByCategory + loadLessonsByCategory */
export default function LessonsByCategory() {
  const { license: licenseId, category: catId } = useParams<{ license: string; category: string }>();
  const license = LICENSES.find((l) => l.id === licenseId) ?? { id: licenseId ?? "?", label: licenseId ?? "?", icon: "📋", desc: "" };
  const cat = CATEGORIES.find((c) => c.id === catId) ?? { id: catId ?? "?", label: catId ?? "?", icon: "📋" };
  const [lessons, setLessons] = useState<Lesson[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    if (!licenseId || !catId) return;
    lessonsApi.byLicenseCategory(licenseId, catId).then((res) => {
      const data: Lesson[] = Array.isArray(res.data) ? res.data : res.data.data ?? [];
      setLessons(data);
    }).finally(() => setLoading(false));
  }, [licenseId, catId]);

  return (
    <div>
      <Navbar />
      <div className="max-w-6xl mx-auto p-4">
        <Link to={`/lessons/${licenseId}`} className="text-slate-400 hover:text-white mb-4 flex items-center gap-1">
          ← Retour à {license.label}
        </Link>
        <div className="flex items-center gap-3 mb-6">
          <span className="text-4xl">{cat.icon}</span>
          <div>
            <h1 className="text-2xl font-bold text-white">{cat.label}</h1>
            <p className="text-slate-400">{license.icon} {license.label}</p>
          </div>
        </div>
        {loading ? (
          <div className="space-y-3">{[1, 2].map((i) => <div key={i} className="skeleton h-20 rounded-xl" />)}</div>
        ) : lessons.length === 0 ? (
          <p className="text-slate-400 text-center py-8">Aucune leçon dans cette catégorie</p>
        ) : (
          <div className="space-y-3">
            {lessons.map((l) => (
              <Link key={l.id} to={`/lessons/${l.license}/${l.category}/${l.id}`} className="block bg-slate-800 rounded-xl p-4 hover:bg-slate-750 transition">
                <h3 className="text-white font-medium">{l.title_fr || l.title_en}</h3>
                <div className="flex gap-2 mt-2">
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