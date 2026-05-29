import { useEffect, useState } from "react";
import { useTranslation } from "react-i18next";
import { CheckCircle2, XCircle, ChevronLeft, ChevronRight } from "lucide-react";
import { historyApi, type AnswerHistory } from "../lib/api";

export default function HistoryPage() {
  const { t } = useTranslation();
  const [history, setHistory] = useState<AnswerHistory[]>([]);
  const [total, setTotal] = useState(0);
  const [page, setPage] = useState(1);
  const [loading, setLoading] = useState(true);
  const pageSize = 20;

  useEffect(() => {
    setLoading(true);
    historyApi
      .list({ page, page_size: pageSize })
      .then((res) => {
        setHistory(res.data);
        setTotal(res.total);
      })
      .finally(() => setLoading(false));
  }, [page]);

  const totalPages = Math.ceil(total / pageSize);

  return (
    <div className="space-y-6 animate-fade-in">
      <div>
        <h1 className="text-3xl font-bold text-white">{t("history.title")}</h1>
        <p className="text-slate-400 mt-1">
          {t("history.subtitle")} ({total} {t("history.entries")})
        </p>
      </div>

      {loading ? (
        <div className="space-y-4 animate-pulse">
          {[...Array(5)].map((_, i) => (
            <div key={i} className="h-16 skeleton" />
          ))}
        </div>
      ) : (
        <>
          <div className="space-y-3">
            {history.map((entry) => (
              <div
                key={entry.id}
                className="card flex items-start gap-4"
              >
                <div className="mt-1">
                  {entry.correct ? (
                    <CheckCircle2 className="w-5 h-5 text-green-400" />
                  ) : (
                    <XCircle className="w-5 h-5 text-red-400" />
                  )}
                </div>
                <div className="flex-1 min-w-0">
                  <p className="text-sm text-slate-200 line-clamp-2">
                    {entry.question_fr}
                  </p>
                  <div className="flex gap-2 mt-2 flex-wrap items-center">
                    <span className="text-xs text-slate-500">
                      {t("history.answer")}: {entry.answer}
                    </span>
                    <span className="badge-blue">{entry.license}</span>
                    <span className="badge-purple">{entry.category}</span>
                    <span className="text-xs text-slate-500">
                      {new Date(entry.answered_at).toLocaleDateString("fr-FR", {
                        day: "numeric",
                        month: "short",
                        hour: "2-digit",
                        minute: "2-digit",
                      })}
                    </span>
                  </div>
                </div>
              </div>
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
    </div>
  );
}
