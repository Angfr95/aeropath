import { useEffect, useState } from "react";
import { Link, useParams, useNavigate } from "react-router-dom";
import { questionsApi } from "@/api/questions";
import { useQuizStore } from "@/store/quizStore";
import { LICENSES, CATEGORIES } from "@/constants/licenses";
import type { Question } from "@/types";
import Navbar from "@/components/Navbar";

/** renderQuestionsByCategory + loadQuestionsByCategory */
export default function QuestionsByCategory() {
  const { license: licenseId, category: catId } = useParams<{ license: string; category: string }>();
  const navigate = useNavigate();
  const startByCategory = useQuizStore((s) => s.startByCategory);
  const [questions, setQuestions] = useState<Question[]>([]);
  const [loading, setLoading] = useState(true);

  const license = LICENSES.find((l) => l.id === licenseId) ?? { id: licenseId ?? "?", label: licenseId ?? "?", icon: "📋", desc: "" };
  const cat = CATEGORIES.find((c) => c.id === catId) ?? { id: catId ?? "?", label: catId ?? "?", icon: "📋" };

  useEffect(() => {
    if (!licenseId || !catId) return;
    questionsApi.byLicenseCategory(licenseId, catId).then((res) => {
      setQuestions(res.data.questions ?? []);
    }).finally(() => setLoading(false));
  }, [licenseId, catId]);

  async function handleQuiz() {
    await startByCategory(licenseId!, catId!);
    navigate("/quiz");
  }

  return (
    <div>
      <Navbar />
      <div className="max-w-6xl mx-auto p-4">
        <Link to={`/questions/${licenseId}`} className="text-slate-400 hover:text-white mb-4 flex items-center gap-1">
          ← Retour à {license.label}
        </Link>
        <div className="flex items-center gap-3 mb-6">
          <span className="text-4xl">{cat.icon}</span>
          <div>
            <h1 className="text-2xl font-bold text-white">{cat.label}</h1>
            <p className="text-slate-400">{license.icon} {license.label}</p>
          </div>
        </div>
        <button onClick={handleQuiz} className="bg-green-600 hover:bg-green-700 text-white px-4 py-2 rounded-lg text-sm mb-4">
          🧠 Quiz sur cette catégorie
        </button>
        {loading ? (
          <div className="space-y-3">{[1, 2, 3].map((i) => <div key={i} className="skeleton h-20 rounded-xl" />)}</div>
        ) : questions.length === 0 ? (
          <p className="text-slate-400 text-center py-8">Aucune question dans cette catégorie</p>
        ) : (
          <div className="space-y-3">
            {questions.map((q) => (
              <Link
                key={q.id}
                to={`/questions/${licenseId}/${catId}/${q.id}`}
                className="block bg-slate-800 rounded-xl p-4 hover:bg-slate-750 transition"
              >
                <div className="flex justify-between items-start">
                  <div className="flex-1">
                    <p className="text-white font-medium">{q.question_fr || q.question_en}</p>
                    <div className="flex gap-2 mt-2 flex-wrap">
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