import { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { lessonsApi } from "@/api/lessons";
import { useQuizStore } from "@/store/quizStore";
import type { Lesson } from "@/types";
import Navbar from "@/components/Navbar";

/** renderLessonDetail + loadLessonDetail + renderMarkdown */
function renderMarkdown(text: string): string {
  if (!text) return "";
  let html = text
    .replace(/&/g, "&")
    .replace(/</g, "<")
    .replace(/>/g, ">");
  html = html.replace(/^### (.+)$/gm, '<h3 class="text-lg font-bold text-white mt-4 mb-2">$1</h3>');
  html = html.replace(/^## (.+)$/gm, '<h2 class="text-xl font-bold text-white mt-5 mb-2">$1</h2>');
  html = html.replace(/^# (.+)$/gm, '<h1 class="text-2xl font-bold text-white mt-6 mb-3">$1</h1>');
  html = html.replace(/\*\*(.+?)\*\*/g, '<strong class="text-white font-bold">$1</strong>');
  html = html.replace(/\*(.+?)\*/g, '<em class="text-slate-200 italic">$1</em>');
  html = html.replace(/^- (.+)$/gm, '<li class="text-slate-300 ml-4 list-disc">$1</li>');
  html = html.replace(/^\d+\. (.+)$/gm, '<li class="text-slate-300 ml-4 list-decimal">$1</li>');
  html = html.replace(/\n\n/g, '</p><p class="text-slate-300 mb-3 leading-relaxed">');
  html = '<p class="text-slate-300 mb-3 leading-relaxed">' + html + '</p>';
  return html;
}

export default function LessonDetail() {
  const { lessonId } = useParams<{ license: string; category: string; lessonId: string }>();
  const navigate = useNavigate();
  const startByLesson = useQuizStore((s) => s.startByLesson);
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
          ← Retour
        </button>
        {loading ? (
          <div className="bg-slate-800 rounded-xl p-6 skeleton h-64" />
        ) : lesson ? (
          <div className="bg-slate-800 rounded-xl p-6">
            <div className="flex justify-between items-start mb-4">
              <h3 className="text-xl font-bold text-white">{lesson.title_fr || lesson.title_en}</h3>
            </div>
            <div
              className="mb-6"
              dangerouslySetInnerHTML={{ __html: renderMarkdown(lesson.content_fr || lesson.content_en || "") || '<p class="text-slate-400 italic">Contenu non disponible</p>' }}
            />
            <div className="flex gap-2 flex-wrap mb-6">
              <span className="text-xs bg-blue-900 text-blue-300 px-2 py-0.5 rounded">{lesson.license}</span>
              <span className="text-xs bg-purple-900 text-purple-300 px-2 py-0.5 rounded">{lesson.category}</span>
              <span className="text-xs bg-amber-900 text-amber-300 px-2 py-0.5 rounded">Niv. {lesson.difficulty}</span>
            </div>
            <button onClick={handleQuiz} className="w-full bg-green-600 hover:bg-green-700 text-white font-medium py-3 rounded-lg transition text-lg">
              📝 Quiz sur cette leçon
            </button>
          </div>
        ) : (
          <p className="text-red-400">Erreur de chargement</p>
        )}
      </div>
    </div>
  );
}