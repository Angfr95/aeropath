import { useEffect, useState } from "react";
import { Link, useParams } from "react-router-dom";
import { questionsApi } from "@/api/questions";
import { useQuizStore } from "@/store/quizStore";
import { LICENSES, CATEGORIES } from "@/constants/licenses";
import type { Question } from "@/types";
import Navbar from "@/components/Navbar";

/** renderQuestionsByLicense + loadQuestionsByLicense */
export default function QuestionsByLicense() {
  const { license: licenseId } = useParams<{ license: string }>();
  const license = LICENSES.find((l) => l.id === licenseId) ?? {
    id: licenseId ?? "?",
    label: licenseId ?? "?",
    icon: "📋",
    desc: "",
  };
  const startByCategory = useQuizStore((s) => s.startByCategory);
  const [questions, setQuestions] = useState<Question[]>([]);
  const [loading, setLoading] = useState(true);
  const [catCounts, setCatCounts] = useState<Record<string, number>>({});

  useEffect(() => {
    if (!licenseId) return;
    questionsApi.byLicense(licenseId).then((res) => {
      const qs = res.data.questions ?? [];
      setQuestions(qs);
      const counts: Record<string, number> = {};
      CATEGORIES.forEach((c) => {
        counts[c.id] = qs.filter((q) => q.category === c.id).length;
      });
      setCatCounts(counts);
    }).finally(() => setLoading(false));
  }, [licenseId]);

  return (
    <div>
      <Navbar />
      <div className="max-w-6xl mx-auto p-4">
        <Link to="/questions" className="text-slate-400 hover:text-white mb-4 flex items-center gap-1">
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
              to={`/questions/${licenseId}/${c.id}`}
              className="bg-slate-800 hover:bg-slate-700 rounded-xl p-4 text-left transition border border-slate-700 hover:border-purple-500/50"
            >
              <div className="flex items-center gap-3">
                <span className="text-2xl">{c.icon}</span>
                <div>
                  <div className="text-white font-medium">{c.label}</div>
                  <div className="text-xs text-slate-500">
                    {loading ? "Chargement..." : `${catCounts[c.id] ?? 0} questions`}
                  </div>
                </div>
              </div>
            </Link>
          ))}
        </div>

        <h2 className="text-lg font-semibold text-slate-300 mb-3">Toutes les questions {license.label}</h2>
        {loading ? (
          <div className="space-y-3">{[1, 2, 3].map((i) => <div key={i} className="skeleton h-20 rounded-xl" />)}</div>
        ) : questions.length === 0 ? (
          <p className="text-slate-400 text-center py-8">Aucune question pour cette licence</p>
        ) : (
          <div className="space-y-3">
            {questions.map((q) => (
              <Link
                key={q.id}
                to={`/questions/${licenseId}/${q.category}/${q.id}`}
                className="block bg-slate-800 rounded-xl p-4 hover:bg-slate-750 transition"
              >
                <div className="flex justify-between items-start">
                  <div className="flex-1">
                    <p className="text-white font-medium">{q.question_fr || q.question_en}</p>
                    <div className="flex gap-2 mt-2 flex-wrap">
                      <span className="text-xs bg-purple-900 text-purple-300 px-2 py-0.5 rounded">{q.category}</span>
                      <span className="text-xs bg-amber-900 text-amber-300 px-2 py-0.5 rounded">Niv. {q.difficulty}</span>
                      <span className="text-xs bg-green-900 text-green-300 px-2 py-0.5 rounded">{q.theme}</span>
                    </div>
                  </div>
                  <span className="text-slate-500 ml-2">→</span>
                </div>
              </Link>
            ))}
          </div>
        )}
      </div>
    </div>
  );
}