import { useEffect, useState } from "react";
import { useTranslation } from "react-i18next";
import { Search, ChevronLeft, ChevronRight, HelpCircle } from "lucide-react";
import { questionsApi, type Question } from "../lib/api";

export default function Questions() {
  const { t } = useTranslation();
  const [questions, setQuestions] = useState<Question[]>([]);
  const [total, setTotal] = useState(0);
  const [page, setPage] = useState(1);
  const [search, setSearch] = useState("");
  const [loading, setLoading] = useState(true);
  const [selectedQuestion, setSelectedQuestion] = useState<Question | null>(null);
  const pageSize = 20;

  useEffect(() => {
    setLoading(true);
    questionsApi
      .list({ page, page_size: pageSize })
      .then((res) => {
        setQuestions(res.data);
        setTotal(res.total);
      })
      .finally(() => setLoading(false));
  }, [page]);

  const handleSearch = async () => {
    if (!search.trim()) return;
    setLoading(true);
    try {
      const results = await questionsApi.search(search);
      setQuestions(results);
      setTotal(results.length);
    } finally {
      setLoading(false);
    }
  };

  const totalPages = Math.ceil(total / pageSize);

  return (
    <div className="space-y-6 animate-fade-in">
      <div>
        <h1 className="text-3xl font-bold text-white">{t("questions.title")}</h1>
        <p className="text-slate-400 mt-1">
          {t("questions.subtitle")} ({total} {t("questions.available")})
        </p>
      </div>

      {/* Search */}
      <div className="flex gap-3">
        <div className="relative flex-1">
          <Search className="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-500" />
          <input
            type="text"
            placeholder={t("questions.search")}
            className="input pl-10"
            value={search}
            onChange={(e) => setSearch(e.target.value)}
            onKeyDown={(e) => e.key === "Enter" && handleSearch()}
          />
        </div>
        <button onClick={handleSearch} className="btn-primary">
          <Search className="w-4 h-4" />
        </button>
      </div>

      {/* Questions List */}
      {loading ? (
        <div className="space-y-4 animate-pulse">
          {[...Array(5)].map((_, i) => (
            <div key={i} className="h-24 skeleton" />
          ))}
        </div>
      ) : (
        <>
          <div className="space-y-3">
            {questions.map((q) => (
              <button
                key={q.id}
                onClick={() => setSelectedQuestion(q)}
                className="w-full text-left card hover:border-blue-500/30 transition-all duration-200 cursor-pointer"
              >
                <div className="flex items-start justify-between gap-4">
                  <div className="flex-1 min-w-0">
                    <p className="text-sm text-slate-200 line-clamp-2">
                      {q.question_fr}
                    </p>
                    <div className="flex gap-2 mt-2 flex-wrap">
                      <span className="badge-blue">{q.license}</span>
                      <span className="badge-purple">{q.category}</span>
                      <span className="badge-yellow">{q.theme}</span>
                      <span className="badge-green">{t("questions.level")} {q.difficulty}</span>
                    </div>
                  </div>
                  <HelpCircle className="w-5 h-5 text-slate-600 shrink-0 mt-1" />
                </div>
              </button>
            ))}
          </div>

          {/* Pagination */}
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

      {/* Question Detail Modal */}
      {selectedQuestion && (
        <div
          className="fixed inset-0 bg-black/60 z-50 flex items-center justify-center p-4"
          onClick={() => setSelectedQuestion(null)}
        >
          <div
            className="card max-w-2xl w-full max-h-[80vh] overflow-y-auto"
            onClick={(e) => e.stopPropagation()}
          >
            <h3 className="text-lg font-semibold text-white mb-4">{t("questions.question")}</h3>
            <p className="text-slate-200 mb-4">{selectedQuestion.question_fr}</p>

            <div className="space-y-2 mb-4">
              {selectedQuestion.options.map((opt, i) => (
                <div
                  key={i}
                  className={`p-3 rounded-lg border ${
                    opt.startsWith(selectedQuestion.answer_key)
                      ? "border-green-500/50 bg-green-500/10"
                      : "border-slate-700 bg-slate-800/50"
                  }`}
                >
                  <span className="text-sm text-slate-300">{opt}</span>
                </div>
              ))}
            </div>

            <div className="p-4 bg-blue-600/10 border border-blue-500/20 rounded-lg mb-4">
              <p className="text-sm font-medium text-blue-400 mb-1">{t("questions.explanation")}</p>
              <p className="text-sm text-slate-300">{selectedQuestion.explanation_fr}</p>
            </div>

            <div className="flex gap-2 flex-wrap">
              <span className="badge-blue">{selectedQuestion.license}</span>
              <span className="badge-purple">{selectedQuestion.category}</span>
              <span className="badge-yellow">{selectedQuestion.theme}</span>
              <span className="badge-green">{t("questions.difficulty")}: {selectedQuestion.difficulty}/5</span>
            </div>

            <button
              onClick={() => setSelectedQuestion(null)}
              className="btn-secondary mt-4 w-full"
            >
              {t("questions.close")}
            </button>
          </div>
        </div>
      )}
    </div>
  );
}
