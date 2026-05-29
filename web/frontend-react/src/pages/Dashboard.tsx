import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { useTranslation } from "react-i18next";
import {
  TrendingUp,
  Target,
  Clock,
  Award,
  ArrowRight,
  Brain,
  BookOpen,
  HelpCircle,
  GraduationCap,
} from "lucide-react";
import { recommendationsApi, statsApi, type Recommendation, type AdminStats } from "../lib/api";

export default function Dashboard() {
  const { t } = useTranslation();
  const [recommendation, setRecommendation] = useState<Recommendation | null>(null);
  const [stats, setStats] = useState<AdminStats | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    Promise.all([
      recommendationsApi.get().catch(() => null),
      statsApi.admin().catch(() => null),
    ]).then(([rec, st]) => {
      setRecommendation(rec);
      setStats(st);
      setLoading(false);
    });
  }, []);

  if (loading) {
    return (
      <div className="space-y-6 animate-pulse">
        <div className="h-8 w-64 skeleton" />
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
          {[...Array(4)].map((_, i) => (
            <div key={i} className="h-32 skeleton" />
          ))}
        </div>
        <div className="h-64 skeleton" />
      </div>
    );
  }

  const progression = recommendation?.progression_percent ?? 0;
  const weakTopics = recommendation?.weak_topics ?? [];
  const dueCards = recommendation?.due_cards ?? [];

  return (
    <div className="space-y-6 animate-fade-in">
      {/* Header */}
      <div>
        <h1 className="text-3xl font-bold text-white">{t("dashboard.title")}</h1>
        <p className="text-slate-400 mt-1">
          {t("dashboard.subtitle")}
        </p>
      </div>

      {/* Stats Cards */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
        <div className="card">
          <div className="flex items-center justify-between">
            <div>
              <p className="text-sm text-slate-400">{t("dashboard.progression")}</p>
              <p className="text-2xl font-bold text-white mt-1">
                {Math.round(progression)}%
              </p>
            </div>
            <div className="w-12 h-12 bg-blue-500/20 rounded-xl flex items-center justify-center">
              <TrendingUp className="w-6 h-6 text-blue-400" />
            </div>
          </div>
          <div className="mt-4 w-full bg-slate-800 rounded-full h-2">
            <div
              className="bg-blue-500 h-2 rounded-full transition-all duration-1000"
              style={{ width: `${progression}%` }}
            />
          </div>
        </div>

        <div className="card">
          <div className="flex items-center justify-between">
            <div>
              <p className="text-sm text-slate-400">{t("dashboard.questions")}</p>
              <p className="text-2xl font-bold text-white mt-1">
                {stats?.total_questions ?? 0}
              </p>
            </div>
            <div className="w-12 h-12 bg-purple-500/20 rounded-xl flex items-center justify-center">
              <HelpCircle className="w-6 h-6 text-purple-400" />
            </div>
          </div>
          <p className="mt-2 text-xs text-slate-500">
            {stats?.total_answers ?? 0} {t("dashboard.answers")}{" "}
            {Math.round((stats?.global_accuracy ?? 0) * 100)}% {t("dashboard.correct")}
          </p>
        </div>

        <div className="card">
          <div className="flex items-center justify-between">
            <div>
              <p className="text-sm text-slate-400">{t("dashboard.lessons")}</p>
              <p className="text-2xl font-bold text-white mt-1">
                {stats?.total_lessons ?? 0}
              </p>
            </div>
            <div className="w-12 h-12 bg-green-500/20 rounded-xl flex items-center justify-center">
              <BookOpen className="w-6 h-6 text-green-400" />
            </div>
          </div>
          <p className="mt-2 text-xs text-slate-500">{t("dashboard.available")}</p>
        </div>

        <div className="card">
          <div className="flex items-center justify-between">
            <div>
              <p className="text-sm text-slate-400">{t("dashboard.toReview")}</p>
              <p className="text-2xl font-bold text-white mt-1">
                {dueCards.length}
              </p>
            </div>
            <div className="w-12 h-12 bg-yellow-500/20 rounded-xl flex items-center justify-center">
              <Clock className="w-6 h-6 text-yellow-400" />
            </div>
          </div>
          <p className="mt-2 text-xs text-slate-500">{t("dashboard.cardsToReview")}</p>
        </div>
      </div>

      {/* Main Grid */}
      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
        {/* Weak Topics */}
        <div className="card">
          <div className="flex items-center justify-between mb-4">
            <h2 className="text-lg font-semibold text-white">
              <Brain className="w-5 h-5 inline mr-2 text-purple-400" />
              {t("dashboard.weakTopics")}
            </h2>
            <Link
              to="/recommendations"
              className="text-sm text-blue-400 hover:text-blue-300 flex items-center gap-1"
            >
              {t("dashboard.seeAll")} <ArrowRight className="w-4 h-4" />
            </Link>
          </div>
          {weakTopics.length === 0 ? (
            <p className="text-slate-500 text-sm py-8 text-center">
              {t("dashboard.noWeakTopics")}
            </p>
          ) : (
            <div className="space-y-3">
              {weakTopics.slice(0, 5).map((topic) => (
                <div
                  key={topic.theme}
                  className="flex items-center justify-between p-3 bg-slate-800/50 rounded-lg"
                >
                  <div className="flex-1">
                    <p className="text-sm font-medium text-slate-200">
                      {topic.theme}
                    </p>
                    <div className="mt-2 w-full bg-slate-700 rounded-full h-1.5">
                      <div
                        className={`h-1.5 rounded-full transition-all ${
                          topic.score < 40
                            ? "bg-red-500"
                            : topic.score < 70
                            ? "bg-yellow-500"
                            : "bg-green-500"
                        }`}
                        style={{ width: `${topic.score}%` }}
                      />
                    </div>
                  </div>
                  <span
                    className={`ml-3 badge ${
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
              ))}
            </div>
          )}
        </div>

        {/* Due Cards */}
        <div className="card">
          <div className="flex items-center justify-between mb-4">
            <h2 className="text-lg font-semibold text-white">
              <Target className="w-5 h-5 inline mr-2 text-blue-400" />
              {t("dashboard.scheduledReviews")}
            </h2>
            <Link
              to="/quiz"
              className="text-sm text-blue-400 hover:text-blue-300 flex items-center gap-1"
            >
              {t("dashboard.review")} <ArrowRight className="w-4 h-4" />
            </Link>
          </div>
          {dueCards.length === 0 ? (
            <p className="text-slate-500 text-sm py-8 text-center">
              {t("dashboard.noReviews")}
            </p>
          ) : (
            <div className="space-y-3">
              {dueCards.slice(0, 5).map((card) => (
                <div
                  key={card.question_id}
                  className="flex items-center justify-between p-3 bg-slate-800/50 rounded-lg"
                >
                  <div className="flex-1 min-w-0">
                    <p className="text-sm text-slate-200 truncate">
                      {card.question_fr}
                    </p>
                    <p className="text-xs text-slate-500 mt-1">
                      {t("dashboard.interval")}: {card.interval_days} {t("dashboard.days")}
                    </p>
                  </div>
                  <span className="badge-blue ml-3 whitespace-nowrap">
                    J+{card.interval_days}
                  </span>
                </div>
              ))}
            </div>
          )}
        </div>
      </div>

      {/* Quick Actions */}
      <div className="card">
        <h2 className="text-lg font-semibold text-white mb-4">
          <Award className="w-5 h-5 inline mr-2 text-yellow-400" />
          {t("dashboard.quickActions")}
        </h2>
        <div className="grid grid-cols-1 sm:grid-cols-3 gap-4">
          <Link
            to="/quiz"
            className="flex items-center gap-3 p-4 bg-blue-600/10 border border-blue-500/20 rounded-xl hover:bg-blue-600/20 transition-colors"
          >
            <GraduationCap className="w-8 h-8 text-blue-400" />
            <div>
              <p className="font-medium text-white">{t("dashboard.quickQuiz")}</p>
              <p className="text-xs text-slate-400">{t("dashboard.quickQuizDesc")}</p>
            </div>
          </Link>
          <Link
            to="/questions"
            className="flex items-center gap-3 p-4 bg-purple-600/10 border border-purple-500/20 rounded-xl hover:bg-purple-600/20 transition-colors"
          >
            <HelpCircle className="w-8 h-8 text-purple-400" />
            <div>
              <p className="font-medium text-white">{t("dashboard.explore")}</p>
              <p className="text-xs text-slate-400">{t("dashboard.exploreDesc")}</p>
            </div>
          </Link>
          <Link
            to="/lessons"
            className="flex items-center gap-3 p-4 bg-green-600/10 border border-green-500/20 rounded-xl hover:bg-green-600/20 transition-colors"
          >
            <BookOpen className="w-8 h-8 text-green-400" />
            <div>
              <p className="font-medium text-white">{t("dashboard.lessons")}</p>
              <p className="text-xs text-slate-400">{t("dashboard.lessonsDesc")}</p>
            </div>
          </Link>
        </div>
      </div>
    </div>
  );
}
