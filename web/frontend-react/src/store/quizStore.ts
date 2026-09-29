import { create } from "zustand";
import type { Question, QuizFeedback, QuizPhase } from "@/types";
import { questionsApi } from "@/api/questions";
import { lessonsApi } from "@/api/lessons";

interface QuizState {
  phase: QuizPhase;
  questions: Question[];
  currentIndex: number;
  score: number;
  feedback: QuizFeedback | null;
  lessonId: string | null;
  loading: boolean;
  submitting: boolean;
  error: string | null;
  startRandom: (limit?: number) => Promise<void>;
  startByCategory: (license: string, category: string) => Promise<void>;
  startByLesson: (lessonId: string) => Promise<void>;
  submitAnswer: (questionId: string, answer: string) => Promise<void>;
  next: () => void;
  restart: () => void;
  reset: () => void;
}

export const useQuizStore = create<QuizState>()((set, get) => {
  // Ignore responses from a quiz that has since been replaced or closed.
  let session = 0;
  const empty = {
    phase: "idle" as QuizPhase, questions: [] as Question[], currentIndex: 0,
    score: 0, feedback: null, lessonId: null, loading: false,
    submitting: false, error: null,
  };
  async function load(request: () => Promise<{ data: Question[] | { questions: Question[] } }>, lessonId: string | null = null) {
    const token = ++session;
    set({ ...empty, loading: true, lessonId });
    try {
      const { data } = await request();
      if (token !== session) return;
      const questions = (Array.isArray(data) ? data : data.questions ?? []).filter(Boolean);
      set({ questions, phase: questions.length ? "active" : "idle", loading: false,
        error: questions.length ? null : "Aucune question disponible pour ce quiz." });
    } catch {
      if (token === session) set({ loading: false, error: "Impossible de charger le quiz. Réessayez depuis la leçon ou la liste des cours." });
    }
  }
  return {
    ...empty,
    startRandom: (limit = 15) => load(() => questionsApi.random(limit)),
    startByCategory: (license, category) => load(() => questionsApi.examByCategory(license, category)),
    startByLesson: (lessonId) => load(() => lessonsApi.quiz(lessonId), lessonId),
    submitAnswer: async (questionId, answer) => {
      const state = get();
      const question = state.questions[state.currentIndex];
      if (state.phase !== "active" || state.submitting || question?.id !== questionId) return;
      if (!question.options.some((_, i) => answer === String.fromCharCode(65 + i))) return;
      const token = session;
      set({ submitting: true, error: null });
      try {
        const { data } = await questionsApi.answer(questionId, answer);
        if (token !== session) return;
        const correctIndex = question.options.findIndex((option) =>
          (typeof option === "string" ? option : option.text) === data.correct_answer);
        const feedback: QuizFeedback = {
          correct: data.correct,
          correct_answer: correctIndex >= 0 ? String.fromCharCode(65 + correctIndex) : data.correct_answer,
          selected_answer: answer,
          explanation_fr: data.explanation_fr,
          explanation_en: data.explanation_en,
          explanation_es: data.explanation_es,
        };
        set((s) => ({ phase: "feedback", submitting: false, feedback,
          score: s.score + (feedback.correct ? 1 : 0) }));
      } catch {
        if (token === session) set({ submitting: false, error: "La correction n’a pas pu être chargée. Réessayez votre réponse." });
      }
    },
    next: () => {
      const { phase, currentIndex, questions } = get();
      if (phase !== "feedback") return;
      if (currentIndex + 1 >= questions.length) set({ phase: "finished", feedback: null });
      else set({ phase: "active", currentIndex: currentIndex + 1, feedback: null, error: null });
    },
    restart: () => {
      if (!get().questions.length) return;
      ++session;
      set({ phase: "active", currentIndex: 0, score: 0, feedback: null,
        loading: false, submitting: false, error: null });
    },
    reset: () => { ++session; set(empty); },
  };
});
