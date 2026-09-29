import { useLearningLocale } from "@/utils/useLearningLocale";
import { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { lessonsApi } from "@/api/lessons";
import { useQuizStore } from "@/store/quizStore";
import type { Lesson } from "@/types";
import Navbar from "@/components/Navbar";

/** renderLessonDetail + loadLessonDetail + renderMarkdown */
import { renderMarkdown } from "@/utils/renderMarkdown";

export default function LessonDetail() {
  const { text, label } = useLearningLocale();
  const { lessonId } = useParams<{ license: string; category: string; lessonId: string }>();
  const navigate = useNavigate();
  const startByLesson = useQuizStore((s) => s.startByLesson);
  const quizLoading = useQuizStore((s) => s.loading);
  const [lesson, setLesson] = useState<Lesson | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    if (!lessonId) return;
    lessonsApi.getById(lessonId).then((res) => setLesson(res.data)).finally(() => setLoading(false));
  }, [lessonId]);

  async function handleQuiz() {
    if (!lessonId) return;
    await startByLesson(lessonId);
    navigate("/quiz");
  }

  return (
    <div>
      <Navbar />
      <div className="max-w-4xl mx-auto p-4">
        <button onClick={() => navigate(-1)} className="text-slate-400 hover:text-white mb-4 flex items-center gap-1">
          {label("← Retour")}
        </button>
        {loading ? (
          <div className="bg-slate-800 rounded-xl p-6 skeleton h-64" />
        ) : lesson ? (
          <div className="bg-slate-800 rounded-xl p-6">
            <div className="flex justify-between items-start mb-4">
              <h3 className="text-xl font-bold text-white">{text(lesson.title_fr, lesson.title_en, lesson.title_es)}</h3>
            </div>
            <div
              className="mb-6"
              dangerouslySetInnerHTML={{ __html: renderMarkdown(text(lesson.content_fr, lesson.content_en, lesson.content_es) || "") || `<p class="text-slate-400 italic">${text('Contenu non disponible', 'Content unavailable')}</p>` }}
            />
            <div className="flex gap-2 flex-wrap mb-6">
              <span className="text-xs bg-blue-900 text-blue-300 px-2 py-0.5 rounded">{lesson.license}</span>
              <span className="text-xs bg-purple-900 text-purple-300 px-2 py-0.5 rounded">{lesson.category}</span>
              <span className="text-xs bg-amber-900 text-amber-300 px-2 py-0.5 rounded">{label("Niv.")} {lesson.difficulty}</span>
            </div>
            <button onClick={handleQuiz} disabled={quizLoading} className="w-full bg-green-600 hover:bg-green-700 disabled:opacity-60 text-white font-medium py-3 rounded-lg transition text-lg">
              {quizLoading ? label("Chargement du quiz…") : label("📝 Quiz sur cette leçon")}
            </button>
          </div>
        ) : (
          <p className="text-red-400">{label("Erreur de chargement")}</p>
        )}
      </div>
    </div>
  );
}
