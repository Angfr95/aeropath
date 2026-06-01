import { Routes, Route } from "react-router-dom";
import ProtectedRoute from "./ProtectedRoute";

// Pages publiques
import HomePage from "@/pages/HomePage";
import LoginPage from "@/pages/LoginPage";

// Pages authentifiées
import Dashboard from "@/pages/Dashboard";
import Questions from "@/pages/Questions";
import QuestionsByLicense from "@/pages/QuestionsByLicense";
import QuestionsByCategory from "@/pages/QuestionsByCategory";
import QuestionDetail from "@/pages/QuestionDetail";
import Quiz from "@/pages/Quiz";
import Lessons from "@/pages/Lessons";
import LessonsByLicense from "@/pages/LessonsByLicense";
import LessonsByCategory from "@/pages/LessonsByCategory";
import LessonDetail from "@/pages/LessonDetail";
import HistoryPage from "@/pages/HistoryPage";
import Stats from "@/pages/Stats";
import Recommendations from "@/pages/Recommendations";

/**
 * Configuration des routes React Router v6.
 *
 * Correspondance avec les vues de app.js :
 *   home             → /
 *   login-form       → /login
 *   dashboard        → /dashboard
 *   questions        → /questions
 *   questions-license→ /questions/:license
 *   questions-category→/questions/:license/:category
 *   question-detail  → /questions/:license/:category/:questionId
 *   quiz             → /quiz
 *   lessons          → /lessons
 *   lessons-license  → /lessons/:license
 *   lessons-category → /lessons/:license/:category
 *   lesson-detail    → /lessons/:license/:category/:lessonId
 *   history          → /history
 *   stats            → /stats
 *   recommendations  → /recommendations
 */
export default function AppRouter() {
  return (
    <Routes>
      {/* Public */}
      <Route path="/" element={<HomePage />} />
      <Route path="/login" element={<LoginPage />} />

      {/* Authentifié */}
      <Route element={<ProtectedRoute />}>
        <Route path="/dashboard" element={<Dashboard />} />
        <Route path="/questions" element={<Questions />} />
        <Route path="/questions/:license" element={<QuestionsByLicense />} />
        <Route path="/questions/:license/:category" element={<QuestionsByCategory />} />
        <Route path="/questions/:license/:category/:questionId" element={<QuestionDetail />} />
        <Route path="/quiz" element={<Quiz />} />
        <Route path="/lessons" element={<Lessons />} />
        <Route path="/lessons/:license" element={<LessonsByLicense />} />
        <Route path="/lessons/:license/:category" element={<LessonsByCategory />} />
        <Route path="/lessons/:license/:category/:lessonId" element={<LessonDetail />} />
        <Route path="/history" element={<HistoryPage />} />
        <Route path="/stats" element={<Stats />} />
        <Route path="/recommendations" element={<Recommendations />} />
      </Route>
    </Routes>
  );
}