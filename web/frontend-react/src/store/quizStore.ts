import { create } from "zustand";
import type { Question, QuizFeedback, QuizPhase } from "@/types";
import { questionsApi } from "@/api/questions";
import { lessonsApi } from "@/api/lessons";

// ---------------------------------------------------------------------------
// Types internes
// ---------------------------------------------------------------------------
interface QuizState {
  phase: QuizPhase;
  questions: Question[];
  currentIndex: number;
  score: number;
  feedback: QuizFeedback | null;

  // actions
  startRandom: (limit?: number) => Promise<void>;
  startByCategory: (license: string, category: string) => Promise<void>;
  startByLesson: (lessonId: string) => Promise<void>;
  submitAnswer: (questionId: string, answer: string) => Promise<void>;
  next: () => void;
  reset: () => void;
}

// ---------------------------------------------------------------------------
// Store
// ---------------------------------------------------------------------------
export const useQuizStore = create<QuizState>()((set, get) => ({
  phase: "idle",
  questions: [],
  currentIndex: 0,
  score: 0,
  feedback: null,

  // ---- startRandom ---------------------------------------------------------
  startRandom: async (limit = 5) => {
    const res = await questionsApi.random(limit);
    const questions = Array.isArray(res.data) ? res.data : [res.data as unknown as Question];
    set({
      phase: "active",
      questions: questions.filter(Boolean),
      currentIndex: 0,
      score: 0,
      feedback: null,
    });
  },

  // ---- startByCategory -----------------------------------------------------
  startByCategory: async (license, category) => {
    // endpoint existant dans app.js : /api/exam/license/:license/category/:category
    const res = await questionsApi.examByCategory(license, category);
    const questions = Array.isArray(res.data) ? res.data : (res.data as any)?.questions ?? [];
    set({
      phase: "active",
      questions,
      currentIndex: 0,
      score: 0,
      feedback: null,
    });
  },

  // ---- startByLesson -------------------------------------------------------
  startByLesson: async (lessonId) => {
    const res = await lessonsApi.quiz(lessonId);
    const questions = Array.isArray(res.data) ? res.data : (res.data as any)?.questions ?? [];
    set({
      phase: "active",
      questions,
      currentIndex: 0,
      score: 0,
      feedback: null,
    });
  },

  // ---- submitAnswer --------------------------------------------------------
  submitAnswer: async (questionId, answer) => {
    const res = await questionsApi.answer(questionId, answer);
    const fb: QuizFeedback = {
      correct: res.data.correct,
      correct_answer: res.data.correct_answer,
      selected_answer: answer,
      explanation_fr: res.data.explanation_fr,
      explanation_en: res.data.explanation_en,
    };
    set((s) => ({
      phase: "feedback",
      score: fb.correct ? s.score + 1 : s.score,
      feedback: fb,
    }));
  },

  // ---- next ----------------------------------------------------------------
  next: () => {
    const { currentIndex, questions } = get();
    const nextIndex = currentIndex + 1;
    if (nextIndex >= questions.length) {
      set({ phase: "finished", feedback: null });
    } else {
      set({ phase: "active", currentIndex: nextIndex, feedback: null });
    }
  },

  // ---- reset ---------------------------------------------------------------
  reset: () => {
    set({
      phase: "idle",
      questions: [],
      currentIndex: 0,
      score: 0,
      feedback: null,
    });
  },
}));