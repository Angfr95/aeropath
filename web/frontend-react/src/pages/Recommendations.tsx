import { useEffect, useState } from "react";
import { useTranslation } from "react-i18next";
import { Brain, Target, Clock, TrendingUp, GraduationCap } from "lucide-react";
import { recommendationsApi, type Recommendation } from "../lib/api";

export default function Recommendations() {
  const { t } = useTranslation();
  const [rec, setRec] = useState<Recommendation | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    recommendationsApi
      .get()
      .then(setRec)
      .finally(() => setLoading(false));
  }, []);

  if (loading) {
    return (
      <div className="space-y-6 animate-pulse">
        <div className="h-8 w-64 skeleton" />
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <div className="h-64 skeleton" />
          <div className="h-64 skeleton" />
        </div>
      </div>
    );
  }

  if (!rec) {
    return (
      <div className="text-center py-20">
        <p className="text-slate-400">
          {t("recommendations.error")}
        </p>
      </div>
    );
  }

  return (
    <div className="space-y-6 animate-fade-in">
      <div>
        <h1 className="text-3xl font-bold text-white">{t("recommendations.title")}</h1>
        <p className="text-slate-400 mt-1">
          {t("recommendations.subtitle")}
        </p>
      </div>

      {/* Progression */}
      <div className="card">
        <div className="flex items-center justify-between mb-4">
          <h2 className="text-lg font-semibold text-white">
            <TrendingUp className="w-5 h-5 inline mr-2 text-blue-400" />
            {t("recommendations.globalProgression")}
          </h2>
          <span className="text-2xl font-bold text-blue-400">
            {Math.round(rec.progression_percent)}%
          </span>
        </div>
        <div className="w-full bg-slate-800 rounded-full h-4">
          <div
            className="bg-gradient-to-r from-blue-500 to-purple-600 h-4 rounded-full transition-all duration-1000"
            style={{ width: `${rec.progression_percent}%` }}
          />
        </div>
        {rec.next_milestone && (
          <p className="text-sm text-slate-400 mt-3">
            {t("recommendations.nextMilestone")} : {rec.next_milestone}
          </p>
        )}
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
        {/* Weak Topics */}
        <div className="card">
          <h2 className="text-lg font-semibold text-white mb-4">
            <Brain className="w-5 h-5 inline mr-2 text-purple-400" />
            {t("recommendations.weakTopics")}
          </h2>
          {rec.weak_topics.length === 0 ? (
            <p className="text-slate-500 text-sm py-8 text-center">
              {t("recommendations.noWeakTopics")}
            </p>
          ) : (
            <div className="space-y-4">
              {rec.weak_topics.map((topic) => (
                <div key={topic.theme}>
                  <div className="flex items-center justify-between mb-1">
                    <span className="text-sm font-medium text-slate-200">
                      {topic.theme}
                    </span>
                    <span
                      className={`badge ${
                        topic.priority === "high"
                          ? "badge-red"
                          : topic.priority === "medium"
                          ? "badge-yellow"
                          : "badge-blue"
                      }`}
                    >
                      {topic.priority}
                    </span>
                  </div>
                  <div className="w-full bg-slate-800 rounded-full h-2">
                    <div
                      className={`h-2 rounded-full ${
                        topic.score < 40
                          ? "bg-red-500"
                          : topic.score < 70
                          ? "bg-yellow-500"
                          : "bg-green-500"
                      }`}
                      style={{ width: `${topic.score}%` }}
                    />
                  </div>
                  <p className="text-xs text-slate-500 mt-1">
                    {t("recommendations.score")}: {Math.round(topic.score)}%
                  </p>
                </div>
              ))}
            </div>
          )}
        </div>

        {/* Due Cards */}
        <div className="card">
          <h2 className="text-lg font-semibold text-white mb-4">
            <Clock className="w-5 h-5 inline mr-2 text-yellow-400" />
            {t("recommendations.reviews")}
          </h2>
          {rec.due_cards.length === 0 ? (
            <p className="text-slate-500 text-sm py-8 text-center">
              {t("recommendations.noReviews")}
            </p>
          ) : (
            <div className="space-y-3">
              {rec.due_cards.map((card) => (
                <div
                  key={card.question_id}
                  className="p-3 bg-slate-800/50 rounded-lg"
                >
                  <p className="text-sm text-slate-200 line-clamp-2">
                    {card.question_fr}
                  </p>
                  <div className="flex items-center justify-between mt-2">
                    <span className="text-xs text-slate-500">
                      {t("recommendations.nextReview")} :{" "}
                      {new Date(card.next_review).toLocaleDateString("fr-FR")}
                    </span>
                    <span className="badge-yellow">
                      J+{card.interval_days}
                    </span>
                  </div>
                </div>
              ))}
            </div>
          )}
        </div>
      </div>

      {/* Mastery by License */}
      <div className="card">
        <h2 className="text-lg font-semibold text-white mb-4">
          <Target className="w-5 h-5 inline mr-2 text-green-400" />
          {t("recommendations.masteryByLicense")}
        </h2>
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4">
          {rec.mastery_by_license.map((ml) => (
            <div key={ml.license} className="p-4 bg-slate-800/50 rounded-lg text-center">
              <p className="text-sm font-medium text-slate-200 mb-2">
                {ml.license}
              </p>
              <p className="text-3xl font-bold text-blue-400">
                {Math.round(ml.score)}%
              </p>
              <div className="mt-2 w-full bg-slate-700 rounded-full h-1.5">
                <div
                  className="bg-blue-500 h-1.5 rounded-full"
                  style={{ width: `${ml.score}%` }}
                />
              </div>
            </div>
          ))}
        </div>
      </div>

      {/* Action */}
      <div className="card bg-gradient-to-r from-blue-600/10 to-purple-600/10 border-blue-500/20">
        <div className="flex items-center gap-4">
          <div className="w-12 h-12 bg-gradient-to-br from-blue-500 to-purple-600 rounded-xl flex items-center justify-center shrink-0">
            <GraduationCap className="w-6 h-6 text-white" />
          </div>
          <div className="flex-1">
            <h3 className="font-semibold text-white">
              {t("recommendations.ready")}
            </h3>
            <p className="text-sm text-slate-400">
              {t("recommendations.readyDesc")}
            </p>
          </div>
          <a href="/quiz" className="btn-primary shrink-0">
            {t("recommendations.quiz")}
          </a>
        </div>
      </div>
    </div>
  );
}
