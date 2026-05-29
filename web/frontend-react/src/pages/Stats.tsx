import { useEffect, useState } from "react";
import { useTranslation } from "react-i18next";
import { BarChart3, TrendingUp, Target, Users } from "lucide-react";
import { statsApi, type AdminStats } from "../lib/api";

export default function Stats() {
  const { t } = useTranslation();
  const [stats, setStats] = useState<AdminStats | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    statsApi
      .admin()
      .then(setStats)
      .finally(() => setLoading(false));
  }, []);

  if (loading) {
    return (
      <div className="space-y-6 animate-pulse">
        <div className="h-8 w-48 skeleton" />
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
          {[...Array(4)].map((_, i) => (
            <div key={i} className="h-28 skeleton" />
          ))}
        </div>
        <div className="h-64 skeleton" />
      </div>
    );
  }

  if (!stats) {
    return (
      <div className="text-center py-20">
        <p className="text-slate-400">{t("stats.error")}</p>
      </div>
    );
  }

  const accuracyPercent = Math.round(stats.global_accuracy * 100);

  return (
    <div className="space-y-6 animate-fade-in">
      <div>
        <h1 className="text-3xl font-bold text-white">{t("stats.title")}</h1>
        <p className="text-slate-400 mt-1">
          {t("stats.subtitle")}
        </p>
      </div>

      {/* Stats Grid */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
        <div className="card">
          <div className="flex items-center justify-between">
            <div>
              <p className="text-sm text-slate-400">{t("stats.questions")}</p>
              <p className="text-3xl font-bold text-white mt-1">
                {stats.total_questions}
              </p>
            </div>
            <div className="w-12 h-12 bg-blue-500/20 rounded-xl flex items-center justify-center">
              <Target className="w-6 h-6 text-blue-400" />
            </div>
          </div>
        </div>

        <div className="card">
          <div className="flex items-center justify-between">
            <div>
              <p className="text-sm text-slate-400">{t("stats.lessons")}</p>
              <p className="text-3xl font-bold text-white mt-1">
                {stats.total_lessons}
              </p>
            </div>
            <div className="w-12 h-12 bg-purple-500/20 rounded-xl flex items-center justify-center">
              <TrendingUp className="w-6 h-6 text-purple-400" />
            </div>
          </div>
        </div>

        <div className="card">
          <div className="flex items-center justify-between">
            <div>
              <p className="text-sm text-slate-400">{t("stats.students")}</p>
              <p className="text-3xl font-bold text-white mt-1">
                {stats.total_students}
              </p>
            </div>
            <div className="w-12 h-12 bg-green-500/20 rounded-xl flex items-center justify-center">
              <Users className="w-6 h-6 text-green-400" />
            </div>
          </div>
        </div>

        <div className="card">
          <div className="flex items-center justify-between">
            <div>
              <p className="text-sm text-slate-400">{t("stats.accuracy")}</p>
              <p className="text-3xl font-bold text-white mt-1">
                {accuracyPercent}%
              </p>
            </div>
            <div className="w-12 h-12 bg-yellow-500/20 rounded-xl flex items-center justify-center">
              <BarChart3 className="w-6 h-6 text-yellow-400" />
            </div>
          </div>
          <div className="mt-3 w-full bg-slate-800 rounded-full h-2">
            <div
              className={`h-2 rounded-full ${
                accuracyPercent >= 80
                  ? "bg-green-500"
                  : accuracyPercent >= 50
                  ? "bg-yellow-500"
                  : "bg-red-500"
              }`}
              style={{ width: `${accuracyPercent}%` }}
            />
          </div>
        </div>
      </div>

      {/* Questions by License */}
      <div className="card">
        <h2 className="text-lg font-semibold text-white mb-4">
          {t("stats.byLicense")}
        </h2>
        <div className="space-y-3">
          {Object.entries(stats.questions_by_license).map(
            ([license, count]) => {
              const maxCount = Math.max(
                ...Object.values(stats.questions_by_license)
              );
              const percent = (count / maxCount) * 100;
              return (
                <div key={license} className="flex items-center gap-4">
                  <span className="text-sm text-slate-300 w-20 font-medium">
                    {license}
                  </span>
                  <div className="flex-1 bg-slate-800 rounded-full h-6">
                    <div
                      className="bg-blue-500 h-6 rounded-full flex items-center justify-end px-3 transition-all"
                      style={{ width: `${percent}%` }}
                    >
                      <span className="text-xs text-white font-medium">
                        {count}
                      </span>
                    </div>
                  </div>
                </div>
              );
            }
          )}
        </div>
      </div>

      {/* Questions by Category */}
      <div className="card">
        <h2 className="text-lg font-semibold text-white mb-4">
          {t("stats.byCategory")}
        </h2>
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-3">
          {Object.entries(stats.questions_by_category).map(
            ([category, count]) => (
              <div
                key={category}
                className="flex items-center justify-between p-3 bg-slate-800/50 rounded-lg"
              >
                <span className="text-sm text-slate-300">{category}</span>
                <span className="badge-blue">{count}</span>
              </div>
            )
          )}
        </div>
      </div>

      {/* Summary */}
      <div className="card">
        <h2 className="text-lg font-semibold text-white mb-4">{t("stats.summary")}</h2>
        <div className="grid grid-cols-1 sm:grid-cols-3 gap-4 text-center">
          <div className="p-4 bg-slate-800/50 rounded-lg">
            <p className="text-2xl font-bold text-blue-400">
              {stats.total_answers}
            </p>
            <p className="text-xs text-slate-400 mt-1">{t("stats.totalAnswers")}</p>
          </div>
          <div className="p-4 bg-slate-800/50 rounded-lg">
            <p className="text-2xl font-bold text-green-400">
              {stats.correct_answers}
            </p>
            <p className="text-xs text-slate-400 mt-1">{t("stats.correctAnswers")}</p>
          </div>
          <div className="p-4 bg-slate-800/50 rounded-lg">
            <p className="text-2xl font-bold text-yellow-400">
              {stats.total_answers - stats.correct_answers}
            </p>
            <p className="text-xs text-slate-400 mt-1">{t("stats.incorrectAnswers")}</p>
          </div>
        </div>
      </div>
    </div>
  );
}
