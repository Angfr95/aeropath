// ======================== TYPES PARTAGÉS AEROPATH ========================

// ---------- Auth ----------
export interface User {
  id: string;
  email: string;
  display_name: string;
  avatar_url: string | null;
  preferred_lang: string;
  license_targets: License[];
  created_at: string;
  updated_at: string;
}

export interface AuthResponse {
  message: string;
  user: User;
}

export interface LoginPayload {
  email: string;
  password: string;
}

export interface RegisterPayload {
  email: string;
  password: string;
}

// ---------- Licence & Catégorie ----------
export type License = "PPL" | "LAPL" | "CPL" | "ATPL" | "IR";

export type Category =
  | "meteorology"
  | "navigation"
  | "airlaw"
  | "aircraft_general"
  | "performance"
  | "human_performance"
  | "operational_procedures"
  | "communications"
  | "principles_of_flight"
  | "flight_planning"
  | "instrumentation"
  | "emergency"
  | "mass_and_balance"
  | "radio_procedure";

export interface LicenseInfo {
  id: License;
  label: string;
  icon: string;   // emoji UTF-8
  desc: string;
}

export interface CategoryInfo {
  id: Category;
  label: string;
  icon: string;   // emoji UTF-8
}

// ---------- Question ----------
export interface QuestionOption {
  letter: string; // A, B, C, D
  text: string;
}

export interface Question {
  id: string;
  question_fr: string;
  question_en: string;
  options: string[] | QuestionOption[];
  answer_key?: string;          // back-end only
  explanation_fr: string;
  explanation_en: string;
  license: License;
  category: Category;
  theme: string;
  subtopic: string;
  difficulty: number;
  reference: string;
}

export interface AnswerResult {
  correct: boolean;
  correct_answer: string;
  explanation_fr: string;
  explanation_en: string;
}

export interface QuizFeedback {
  correct: boolean;
  correct_answer: string;
  selected_answer: string;
  explanation_fr: string;
  explanation_en: string;
}

// ---------- Leçon ----------
export interface Lesson {
  id: string;
  title_fr: string;
  title_en: string;
  content_fr: string;
  content_en: string;
  license: License;
  category: Category;
  theme: string;
  difficulty: number;
  order_index: number;
}

// ---------- Historique ----------
export interface HistoryEntry {
  id: string;
  question_id: string;
  question_fr: string;
  question_en: string;
  answer: string;
  was_correct: boolean;
  license: License;
  category: Category;
  theme: string;
  difficulty: number;
  answered_at: string;
}

// ---------- Statistiques ----------
export interface UserStats {
  total_answers: number;
  correct_answers: number;
  accuracy: number;
  questions_by_license: Record<License, number>;
  questions_by_category: Record<string, number>;
  streak_days: number;
  last_activity: string;
}

export interface AdminStats {
  total_questions: number;
  total_lessons: number;
  total_students: number;
  total_answers: number;
  correct_answers: number;
  global_accuracy: number;
  questions_by_license: Record<string, number>;
  questions_by_category: Record<string, number>;
}

// ---------- Recommandations ----------
export interface WeakTopic {
  theme: string;
  score: number;      // 0–100
  priority: "high" | "medium" | "low";
}

export interface DueCard {
  question_id: string;
  question_fr: string;
  question_en: string;
  next_review: string;   // ISO 8601
  interval_days: number;
}

export interface MasteryByLicense {
  license: License;
  score: number;          // 0–100
}

export interface Recommendations {
  progression_percent: number;          // 0–100
  weak_topics: WeakTopic[];
  due_cards: DueCard[];
  next_milestone: string;
  mastery_by_license: MasteryByLicense[];
}

// ---------- État du quiz (machine d'état) ----------
export type QuizPhase = "idle" | "active" | "feedback" | "finished";

export interface QuizState {
  phase: QuizPhase;
  questions: Question[];
  currentIndex: number;
  score: number;
  feedback: QuizFeedback | null;
}

// ---------- API pagination ----------
export interface PaginatedResponse<T> {
  data: T[];
  total: number;
  page: number;
  page_size: number;
  total_pages: number;
}

// ---------- Toast ----------
export type ToastType = "info" | "success" | "error";

export interface Toast {
  id: string;
  message: string;
  type: ToastType;
  durationMs?: number;
}

// ---------- Langue ----------
export interface SupportedLanguage {
  code: string;
  label: string;
  flag: string;   // emoji UTF-8
}

// ---------- Offline ----------
export interface PendingAnswer {
  id?: number;
  question_id: string;
  answer: string;
  token: string;
  timestamp: number;
}