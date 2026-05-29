import { useState, useCallback } from "react";
import { useTranslation } from "react-i18next";
import { questionsApi, type Question } from "../lib/api";
import {
  GraduationCap,
  CheckCircle2,
  XCircle,
  RefreshCw,
  ArrowRight,
  Trophy,
} from "lucide-react";

type QuizState = "idle" | "playing" | "answered" | "finished";

export default function Quiz() {
  const { t } = useTranslation();
  const [state, setState] = useState<QuizState>("idle");
  const [questions, setQuestions] = useState<Question[]>([]);
  const [currentIndex, setCurrentIndex] = useState(0);
  const [selectedAnswer, setSelectedAnswer] = useState("");
  const [result, setResult] = useState<{
    correct: boolean;
    correct_answer: string;
    explanation_fr: string;
  } | null>(null);
  const [score, setScore] = useState(0);
  const [total, setTotal] = useState(0);
  const [loading, setLoading] = useState(false);

  const startQuiz = useCallback(async () => {
    setLoading(true);
    try {
      const qs: Question[] = [];
      for (let i = 0; i < 10; i++) {
        const q = await questionsApi.random();
        qs.push(q);
      }
      setQuestions(qs);
      setCurrentIndex(0);
      setScore(0);
      setTotal(0);
      setState("playing");
      setSelectedAnswer("");
      setResult(null);
    } finally {
      setLoading(false);
    }
  }, []);

  const handleAnswer = async (answer: string) => {
    if (state !== "playing") return;
    setSelectedAnswer(answer);
    setState("answered");

    try {
      const res = await questionsApi.answer(questions[currentIndex].id, answer);
      setResult(res);
      if (res.correct) setScore((s) => s + 1);
      setTotal((t) => t + 1);
    } catch {
      // fallback local
      const q = questions[currentIndex];
      const correct = answer === q.answer_key;
      setResult({
        correct,
        correct_answer: q.answer_key,
        explanation_fr: q.explanation_fr,
      });
      if (correct) setScore((s) => s + 1);
      setTotal((t) => t + 1);
    }
  };

  const nextQuestion = () => {
    if (currentIndex < questions.length - 1) {
      setCurrentIndex((i) => i + 1);
      setState("playing");
      setSelectedAnswer("");
      setResult(null);
    } else {
      setState("finished");
    }
  };

  if (state === "idle") {
    return (
      <div className="flex flex-col items-center justify-center py-20 animate-fade-in">
        <div className="w-20 h-20 bg-gradient-to-br from-blue-500 to-purple-600 rounded-2xl flex items-center justify-center mb-6">
          <GraduationCap className="w-10 h-10 text-white" />
        </div>
        <h1 className="text-3xl font-bold text-white mb-2">{t("quiz.title")}</h1>
        <p className="text-slate-400 mb-8 text-center max-w-md">
          {t("quiz.subtitle")}
        </p>
        <button
          onClick={startQuiz}
          disabled={loading}
          className="btn-primary text-lg px-8 py-3"
        >
          {loading ? t("quiz.loading") : t("quiz.start")}
        </button>
      </div>
    );
  }

  if (state === "finished") {
    const percent = Math.round((score / total) * 100);
    return (
      <div className="flex flex-col items-center justify-center py-20 animate-fade-in">
        <div className="w-20 h-20 bg-gradient-to-br from-yellow-500 to-orange-600 rounded-2xl flex items-center justify-center mb-6">
          <Trophy className="w-10 h-10 text-white" />
        </div>
        <h1 className="text-3xl font-bold text-white mb-2">{t("quiz.finished")}</h1>
        <p className="text-5xl font-bold text-blue-400 my-6">{percent}%</p>
        <p className="text-slate-400 mb-2">
          {score} / {total} {t("quiz.correctAnswers")}
        </p>
        <div className="w-64 bg-slate-800 rounded-full h-3 mb-8">
          <div
            className={`h-3 rounded-full transition-all ${
              percent >= 80
                ? "bg-green-500"
                : percent >= 50
                ? "bg-yellow-500"
                : "bg-red-500"
            }`}
            style={{ width: `${percent}%` }}
          />
        </div>
        <button onClick={startQuiz} className="btn-primary flex items-center gap-2">
          <RefreshCw className="w-4 h-4" /> {t("quiz.retry")}
        </button>
      </div>
    );
  }

  const question = questions[currentIndex];

  return (
    <div className="max-w-2xl mx-auto animate-fade-in">
      {/* Progress */}
      <div className="flex items-center justify-between mb-6">
        <span className="text-sm text-slate-400">
          {t("quiz.question")} {currentIndex + 1} / {questions.length}
        </span>
        <span className="text-sm text-slate-400">
          {t("quiz.score")}: {score} / {total}
        </span>
      </div>
      <div className="w-full bg-slate-800 rounded-full h-2 mb-8">
        <div
          className="bg-blue-500 h-2 rounded-full transition-all"
          style={{
            width: `${((currentIndex + 1) / questions.length) * 100}%`,
          }}
        />
      </div>

      {/* Question */}
      <div className="card mb-6">
        <h2 className="text-lg text-white font-medium mb-4">
          {question.question_fr}
        </h2>

        <div className="space-y-3">
          {question.options.map((opt, i) => {
            const isSelected = selectedAnswer === opt[0];
            const isCorrectAnswer =
              state === "answered" && opt[0] === result?.correct_answer;
            const isWrongSelection =
              state === "answered" && isSelected && !result?.correct;

            let borderClass = "border-slate-700 hover:border-blue-500/50";
            if (isCorrectAnswer) borderClass = "border-green-500 bg-green-500/10";
            else if (isWrongSelection)
              borderClass = "border-red-500 bg-red-500/10";
            else if (isSelected)
              borderClass = "border-blue-500 bg-blue-500/10";

            return (
              <button
                key={i}
                onClick={() => handleAnswer(opt[0])}
                disabled={state === "answered"}
                className={`w-full text-left p-4 rounded-lg border transition-all ${borderClass} ${
                  state === "answered" ? "cursor-default" : "cursor-pointer"
                }`}
              >
                <span className="text-sm text-slate-200">{opt}</span>
              </button>
            );
          })}
        </div>
      </div>

      {/* Feedback */}
      {state === "answered" && result && (
        <div className="animate-slide-up">
          <div
            className={`p-4 rounded-lg mb-4 flex items-start gap-3 ${
              result.correct
                ? "bg-green-500/10 border border-green-500/30"
                : "bg-red-500/10 border border-red-500/30"
            }`}
          >
            {result.correct ? (
              <CheckCircle2 className="w-5 h-5 text-green-400 mt-0.5 shrink-0" />
            ) : (
              <XCircle className="w-5 h-5 text-red-400 mt-0.5 shrink-0" />
            )}
            <div>
              <p className="font-medium text-white mb-1">
                {result.correct ? t("quiz.correct") : t("quiz.incorrect")}
              </p>
              <p className="text-sm text-slate-300">{result.explanation_fr}</p>
            </div>
          </div>

          <button onClick={nextQuestion} className="btn-primary w-full flex items-center justify-center gap-2">
            {currentIndex < questions.length - 1 ? (
              <>{t("quiz.next")} <ArrowRight className="w-4 h-4" /></>
            ) : (
              <>{t("quiz.seeResults")} <Trophy className="w-4 h-4" /></>
            )}
          </button>
        </div>
      )}
    </div>
  );
}
