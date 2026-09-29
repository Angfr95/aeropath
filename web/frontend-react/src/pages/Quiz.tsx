import { useLearningLocale } from "@/utils/useLearningLocale";
import { useAuthStore } from "@/store/authStore";
import { useQuizStore } from "@/store/quizStore";
import { useNavigate } from "react-router-dom";
import Navbar from "@/components/Navbar";
import { renderMarkdown } from "@/utils/renderMarkdown";

/** renderQuiz + submitQuizAnswer + nextQuizQuestion + renderQuizResult */
export default function Quiz() {
  const { text, label, choices } = useLearningLocale();
  const isAuthenticated = useAuthStore((s) => s.isAuthenticated);
  const phase = useQuizStore((s) => s.phase);
  const questions = useQuizStore((s) => s.questions);
  const currentIndex = useQuizStore((s) => s.currentIndex);
  const score = useQuizStore((s) => s.score);
  const feedback = useQuizStore((s) => s.feedback);
  const submitAnswer = useQuizStore((s) => s.submitAnswer);
  const next = useQuizStore((s) => s.next);
  const startRandom = useQuizStore((s) => s.startRandom);
  const reset = useQuizStore((s) => s.reset);
  const restart = useQuizStore((s) => s.restart);
  const lessonId = useQuizStore((s) => s.lessonId);
  const loading = useQuizStore((s) => s.loading);
  const submitting = useQuizStore((s) => s.submitting);
  const error = useQuizStore((s) => s.error);
  const navigate = useNavigate();
  const origin = questions[0];
  const lessonPath = lessonId && origin
    ? `/lessons/${origin.license}/${origin.category}/${lessonId}` : null;

  // Redirige si aucun quiz actif
  if (phase === "idle") {
    return (
      <div>
        <Navbar />
        <div className="max-w-2xl mx-auto p-4 text-center mt-20">
          <p className="text-slate-400 mb-4" role="status">{loading ? label("Chargement du quiz…") : label("Pas de quiz en cours")}</p>
          {error && <p role="alert" className="text-red-400 mb-4">{label(error)}</p>}
          <button
            disabled={loading}
            onClick={async () => { await startRandom(15); }}
            className="bg-blue-600 hover:bg-blue-700 text-white px-6 py-2 rounded-lg"
          >
            {label("Lancer un quiz")}
          </button>
          <button onClick={() => navigate("/lessons")} className="block mx-auto mt-4 text-blue-400 underline">{label("Choisir une leçon")}</button>
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
            <h2 className="text-2xl font-bold text-white mb-2">{label("Quiz terminé !")}</h2>
            <p className="text-4xl font-bold text-blue-400 mb-2">{pct}%</p>
            <p className="text-slate-400 mb-6">{score}/{questions.length} {label("bonnes réponses")}</p>
            {!isAuthenticated && (
              <p className="text-slate-400 mb-6">{label("Résultat non enregistré.")} <a href="/login" className="text-blue-400 underline">{label("Connectez-vous")}</a> {label("pour suivre vos prochains quiz.")}</p>
            )}
            <div className="flex flex-wrap gap-3 justify-center">
              <button
                onClick={restart}
                className="bg-blue-600 hover:bg-blue-700 text-white px-6 py-2 rounded-lg"
              >
                {label("Refaire ce quiz")}
              </button>
              {lessonPath && <button onClick={() => { reset(); navigate(lessonPath); }} className="bg-slate-700 hover:bg-slate-600 text-white px-6 py-2 rounded-lg">{label("Revoir la leçon")}</button>}
              <button onClick={() => { reset(); navigate(isAuthenticated ? "/dashboard" : "/lessons"); }} className="bg-slate-700 hover:bg-slate-600 text-white px-6 py-2 rounded-lg">
                {label("Accueil")}
              </button>
            </div>
          </div>
        </div>
      </div>
    );
  }

  const q = questions[currentIndex];
  if (!q) return null;
  const options = choices(q);

  return (
    <div>
      <Navbar />
      <div className="max-w-2xl mx-auto p-4">
        <div className="mb-4">
          <div className="flex justify-between text-sm text-slate-400 mb-2">
            <span>Question {currentIndex + 1}/{questions.length}</span>
            <span>Score : {score}/{currentIndex + (feedback ? 1 : 0)}</span>
          </div>
          <div className="w-full bg-slate-700 rounded-full h-2">
            <div className="bg-blue-500 rounded-full h-2" style={{ width: `${((currentIndex + (feedback ? 1 : 0)) / questions.length) * 100}%` }} />
          </div>
        </div>

        <div className="bg-slate-800 rounded-xl p-6">
          {lessonPath && <button onClick={() => navigate(lessonPath)} className="text-blue-400 underline mb-4">{label("Revoir la leçon")}</button>}
          <div className="text-white text-lg mb-6" dangerouslySetInnerHTML={{ __html: renderMarkdown(text(q.question_fr, q.question_en, q.question_es)) }} />
          {error && <p role="alert" className="text-red-400 mb-4">{label(error)}</p>}
          {submitting && <p role="status" className="text-slate-400 mb-4">{label("Vérification de la réponse…")}</p>}
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
                  disabled={!!feedback || submitting}
                  className={`w-full ${btnClass}`}
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
                  {feedback.correct ? label("✅ Correct !") : label("❌ Faux")}
                </p>
                <p className="text-green-400 font-medium mb-2">{label("Réponse correcte :")} {feedback.correct_answer}{options[feedback.correct_answer.charCodeAt(0) - 65] ? `. ${options[feedback.correct_answer.charCodeAt(0) - 65]}` : ""}</p>
                <div className="text-slate-300 mt-2" dangerouslySetInnerHTML={{ __html: renderMarkdown(text(feedback.explanation_fr, feedback.explanation_en, feedback.explanation_es) || "") }} />
              </div>
              <button onClick={next} className="w-full bg-blue-600 hover:bg-blue-700 text-white font-medium py-3 rounded-lg transition text-lg">
                {currentIndex + 1 >= questions.length ? label("📊 Voir les résultats") : label("➡️ Question suivante")}
              </button>
            </>
          )}
        </div>
      </div>
    </div>
  );
}
