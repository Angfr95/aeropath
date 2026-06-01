import { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { questionsApi } from "@/api/questions";
import type { Question, AnswerResult } from "@/types";
import Navbar from "@/components/Navbar";

/** renderQuestionDetail + loadQuestionDetail + selectOption */
export default function QuestionDetail() {
  const { questionId } = useParams<{ license: string; category: string; questionId: string }>();
  const navigate = useNavigate();
  const [question, setQuestion] = useState<Question | null>(null);
  const [result, setResult] = useState<AnswerResult | null>(null);
  const [selected, setSelected] = useState<string | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    if (!questionId) return;
    questionsApi.getById(questionId).then((res) => {
      setQuestion(res.data);
    }).finally(() => setLoading(false));
  }, [questionId]);

  async function handleSelect(letter: string) {
    if (!questionId || selected) return;
    setSelected(letter);
    try {
      const res = await questionsApi.answer(questionId, letter);
      setResult(res.data);
    } catch {
      // ignore
    }
  }

  const options: string[] = Array.isArray(question?.options)
    ? question.options.map((o) => (typeof o === "string" ? o : o.text))
    : [];

  function optionClass(letter: string) {
    if (!result) return "bg-slate-700 hover:bg-slate-600 border-transparent";
    if (letter === result.correct_answer) return "bg-green-700 border-green-500";
    if (letter === selected && !result.correct) return "bg-red-700 border-red-500";
    return "bg-slate-700 border-transparent";
  }

  return (
    <div>
      <Navbar />
      <div className="max-w-3xl mx-auto p-4">
        <button onClick={() => navigate(-1)} className="text-slate-400 hover:text-white mb-4 flex items-center gap-1">
          ← Retour
        </button>
        {loading ? (
          <div className="bg-slate-800 rounded-xl p-6 skeleton h-64" />
        ) : question ? (
          <div className="bg-slate-800 rounded-xl p-6">
            <div className="flex justify-between items-start mb-4">
              <h3 className="text-lg font-bold text-white">Question</h3>
            </div>
            <p className="text-white text-lg mb-6">{question.question_fr || question.question_en}</p>
            <div className="space-y-3 mb-6">
              {options.map((opt, i) => {
                const letter = String.fromCharCode(65 + i);
                const disabled = !!selected;
                return (
                  <button
                    key={letter}
                    onClick={() => handleSelect(letter)}
                    disabled={disabled}
                    className={`w-full text-white text-left rounded-lg p-3 transition border ${optionClass(letter)} ${disabled ? "cursor-default" : ""}`}
                  >
                    {letter}. {opt}
                  </button>
                );
              })}
            </div>
            <div className="flex gap-2 flex-wrap mb-6">
              <span className="text-xs bg-blue-900 text-blue-300 px-2 py-0.5 rounded">{question.license}</span>
              <span className="text-xs bg-purple-900 text-purple-300 px-2 py-0.5 rounded">{question.category}</span>
              <span className="text-xs bg-amber-900 text-amber-300 px-2 py-0.5 rounded">Niv. {question.difficulty}</span>
              <span className="text-xs bg-green-900 text-green-300 px-2 py-0.5 rounded">{question.theme}</span>
            </div>
            {result && (
              <div className={`mt-4 p-4 rounded-lg ${result.correct ? "bg-green-900/30" : "bg-red-900/30"}`}>
                <p className={`font-bold text-lg mb-2 ${result.correct ? "text-green-400" : "text-red-400"}`}>
                  {result.correct ? "✅ Correct !" : "❌ Faux"}
                </p>
                <p className="text-green-400 font-medium mb-2">Réponse correcte : {result.correct_answer}</p>
                {result.explanation_fr && <p className="text-slate-300 mt-2">{result.explanation_fr}</p>}
                {result.explanation_en && !result.explanation_fr && <p className="text-slate-300 mt-2">{result.explanation_en}</p>}
              </div>
            )}
          </div>
        ) : (
          <p className="text-red-400">Erreur de chargement</p>
        )}
      </div>
    </div>
  );
}