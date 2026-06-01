import { useQuizStore } from "@/store/quizStore";
import { useNavigate } from "react-router-dom";
import Navbar from "@/components/Navbar";

/** renderQuiz + submitQuizAnswer + nextQuizQuestion + renderQuizResult */
export default function Quiz() {
  const phase = useQuizStore((s) => s.phase);
  const questions = useQuizStore((s) => s.questions);
  const currentIndex = useQuizStore((s) => s.currentIndex);
  const score = useQuizStore((s) => s.score);
  const feedback = useQuizStore((s) => s.feedback);
  const submitAnswer = useQuizStore((s) => s.submitAnswer);
  const next = useQuizStore((s) => s.next);
  const startRandom = useQuizStore((s) => s.startRandom);
  const reset = useQuizStore((s) => s.reset);
  const navigate = useNavigate();

  // Redirige si aucun quiz actif
  if (phase === "idle") {
    return (
      <div>
        <Navbar />
        <div className="max-w-2xl mx-auto p-4 text-center mt-20">
          <p className="text-slate-400 mb-4">Pas de quiz en cours</p>
          <button
            onClick={async () => { await startRandom(5); }}
            className="bg-blue-600 hover:bg-blue-700 text-white px-6 py-2 rounded-lg"
          >
            Lancer un quiz
          </button>
        </div>
      </div>
    );
  }

  // Quiz terminé
  if (phase === "finished") {
    const pct = questions.length > 0 ? Math.round((score / questions.length) * 100) : 0;
    return (
      <div>
        <Navbar />
        <div className="max-w-2xl mx-auto p-4 text-center">
          <div className="bg-slate-800 rounded-xl p-8">
            <div className="text-6xl mb-4">{pct >= 80 ? "🎉" : pct >= 50 ? "👍" : "📚"}</div>
            <h2 className="text-2xl font-bold text-white mb-2">Quiz terminé !</h2>
            <p className="text-4xl font-bold text-blue-400 mb-2">{pct}%</p>
            <p className="text-slate-400 mb-6">{score}/{questions.length} bonnes réponses</p>
            <div className="flex gap-3 justify-center">
              <button
                onClick={async () => { reset(); await startRandom(5); }}
                className="bg-blue-600 hover:bg-blue-700 text-white px-6 py-2 rounded-lg"
              >
                Rejouer
              </button>
              <button onClick={() => { reset(); navigate("/dashboard"); }} className="bg-slate-700 hover:bg-slate-600 text-white px-6 py-2 rounded-lg">
                Accueil
              </button>
            </div>
          </div>
        </div>
      </div>
    );
  }

  const q = questions[currentIndex];
  if (!q) return null;
  const options: string[] = Array.isArray(q.options)
    ? q.options.map((o) => (typeof o === "string" ? o : o.text))
    : [];

  return (
    <div>
      <Navbar />
      <div className="max-w-2xl mx-auto p-4">
        <div className="mb-4">
          <div className="flex justify-between text-sm text-slate-400 mb-2">
            <span>Question {currentIndex + 1}/{questions.length}</span>
            <span>Score: {score}/{currentIndex}</span>
          </div>
          <div className="w-full bg-slate-700 rounded-full h-2">
            <div className="bg-blue-500 rounded-full h-2" style={{ width: `${(currentIndex / questions.length) * 100}%` }} />
          </div>
        </div>

        <div className="bg-slate-800 rounded-xl p-6">
          <p className="text-white text-lg mb-6">{q.question_fr || q.question_en}</p>
          <div className="space-y-3 mb-6">
            {options.map((opt, i) => {
              const letter = String.fromCharCode(65 + i);
              let btnClass = feedback
                ? "bg-slate-700 text-white text-left rounded-lg p-3 cursor-default"
                : "bg-slate-700 hover:bg-slate-600 text-white text-left rounded-lg p-3 transition";
              if (feedback && letter === feedback.correct_answer) btnClass = "bg-green-700 text-white text-left rounded-lg p-3 border border-green-500";
              else if (feedback && feedback.selected_answer === letter && !feedback.correct) btnClass = "bg-red-700 text-white text-left rounded-lg p-3 border border-red-500";
              return (
                <button
                  key={letter}
                  onClick={() => feedback ? undefined : submitAnswer(q.id, letter)}
                  disabled={!!feedback}
                  className={btnClass}
                >
                  {letter}. {opt}
                </button>
              );
            })}
          </div>

          {feedback && (
            <>
              <div className={`p-4 rounded-lg ${feedback.correct ? "bg-green-900/30" : "bg-red-900/30"} mb-4`}>
                <p className={`font-bold text-lg mb-2 ${feedback.correct ? "text-green-400" : "text-red-400"}`}>
                  {feedback.correct ? "✅ Correct !" : "❌ Faux"}
                </p>
                <p className="text-green-400 font-medium mb-2">Réponse correcte : {feedback.correct_answer}</p>
                {feedback.explanation_fr && <p className="text-slate-300 mt-2">{feedback.explanation_fr}</p>}
                {feedback.explanation_en && !feedback.explanation_fr && <p className="text-slate-300 mt-2">{feedback.explanation_en}</p>}
              </div>
              <button onClick={next} className="w-full bg-blue-600 hover:bg-blue-700 text-white font-medium py-3 rounded-lg transition text-lg">
                {currentIndex + 1 >= questions.length ? "📊 Voir les résultats" : "➡️ Question suivante"}
              </button>
            </>
          )}
        </div>
      </div>
    </div>
  );
}