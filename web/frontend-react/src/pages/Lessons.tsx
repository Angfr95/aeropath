import { useEffect, useState } from "react";
import { useTranslation } from "react-i18next";
import { BookOpen, ChevronLeft, ChevronRight } from "lucide-react";
import { lessonsApi, type Lesson } from "../lib/api";

export default function Lessons() {
  const { t } = useTranslation();
  const [lessons, setLessons] = useState<Lesson[]>([]);
  const [total, setTotal] = useState(0);
  const [page, setPage] = useState(1);
  const [loading, setLoading] = useState(true);
  const [selectedLesson, setSelectedLesson] = useState<Lesson | null>(null);
  const pageSize = 10;

  useEffect(() => {
    setLoading(true);
    lessonsApi
      .list({ page, page_size: pageSize })
      .then((res) => {
        setLessons(res.data);
        setTotal(res.total);
      })
      .finally(() => setLoading(false));
  }, [page]);

  const totalPages = Math.ceil(total / pageSize);

  return (
    <div className="space-y-6 animate-fade-in">
      <div>
        <h1 className="text-3xl font-bold text-white">{t("lessons.title")}</h1>
        <p className="text-slate-400 mt-1">
          {t("lessons.subtitle")} ({total} {t("lessons.lessons")})
        </p>
      </div>

      {loading ? (
        <div className="space-y-4 animate-pulse">
          {[...Array(5)].map((_, i) => (
            <div key={i} className="h-20 skeleton" />
          ))}
        </div>
      ) : (
        <>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            {lessons.map((lesson) => (
              <button
                key={lesson.id}
                onClick={() => setSelectedLesson(lesson)}
                className="card text-left hover:border-blue-500/30 transition-all cursor-pointer"
              >
                <div className="flex items-start gap-4">
                  <div className="w-10 h-10 bg-purple-500/20 rounded-lg flex items-center justify-center shrink-0">
                    <BookOpen className="w-5 h-5 text-purple-400" />
                  </div>
                  <div className="flex-1 min-w-0">
                    <h3 className="font-medium text-white truncate">
                      {lesson.title_fr}
                    </h3>
                    <p className="text-xs text-slate-400 mt-1 line-clamp-2">
                      {lesson.content_fr?.substring(0, 100)}...
                    </p>
                    <div className="flex gap-2 mt-2">
                      <span className="badge-blue">{lesson.license}</span>
                      <span className="badge-purple">{lesson.category}</span>
                    </div>
                  </div>
                </div>
              </button>
            ))}
          </div>

          {totalPages > 1 && (
            <div className="flex items-center justify-center gap-4">
              <button
                onClick={() => setPage((p) => Math.max(1, p - 1))}
                disabled={page === 1}
                className="btn-secondary disabled:opacity-50"
              >
                <ChevronLeft className="w-4 h-4" />
              </button>
              <span className="text-sm text-slate-400">
                {t("questions.page")} {page} / {totalPages}
              </span>
              <button
                onClick={() => setPage((p) => Math.min(totalPages, p + 1))}
                disabled={page === totalPages}
                className="btn-secondary disabled:opacity-50"
              >
                <ChevronRight className="w-4 h-4" />
              </button>
            </div>
          )}
        </>
      )}

      {/* Lesson Detail Modal */}
      {selectedLesson && (
        <div
          className="fixed inset-0 bg-black/60 z-50 flex items-center justify-center p-4"
          onClick={() => setSelectedLesson(null)}
        >
          <div
            className="card max-w-3xl w-full max-h-[80vh] overflow-y-auto"
            onClick={(e) => e.stopPropagation()}
          >
            <div className="flex items-center gap-3 mb-4">
              <div className="w-10 h-10 bg-purple-500/20 rounded-lg flex items-center justify-center">
                <BookOpen className="w-5 h-5 text-purple-400" />
              </div>
              <div>
                <h2 className="text-xl font-bold text-white">
                  {selectedLesson.title_fr}
                </h2>
                <div className="flex gap-2 mt-1">
                  <span className="badge-blue">{selectedLesson.license}</span>
                  <span className="badge-purple">{selectedLesson.category}</span>
                </div>
              </div>
            </div>

            <div className="prose prose-invert max-w-none">
              <p className="text-slate-300 whitespace-pre-wrap">
                {selectedLesson.content_fr}
              </p>
            </div>

            <button
              onClick={() => setSelectedLesson(null)}
              className="btn-secondary mt-6 w-full"
            >
              {t("lessons.close")}
            </button>
          </div>
        </div>
      )}
    </div>
  );
}
