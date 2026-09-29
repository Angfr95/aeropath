import { create } from "zustand";
import { apiClient } from "@/api/client";
import { questionsApi } from "@/api/questions";
import type { Question, AnswerResult } from "@/types";

export interface PPLPaper { id: string; title_fr: string; title_en: string; count: number; minutes: number }
type Result = AnswerResult & { selected: string };
interface ExamState {
  phase: "idle" | "active" | "grading" | "finished";
  paper: PPLPaper | null;
  questions: Question[];
  answers: Record<string, string>;
  results: Record<string, Result>;
  completed: Record<string, { correct: number; total: number; passed: boolean }>;
  deadline: number;
  busy: boolean;
  error: string | null;
  start: (id: string) => Promise<void>;
  choose: (id: string, letter: string) => void;
  finish: () => Promise<void>;
  close: () => void;
}
const key = "aeropath_ppl_exam_v1";
const blank = { phase: "idle" as const, paper: null, questions: [] as Question[], answers: {}, results: {}, deadline: 0 };
function restored() {
  try {
    const value = JSON.parse(sessionStorage.getItem(key) || "null");
    if (value && ["idle", "active", "grading", "finished"].includes(value.phase) &&
        Array.isArray(value.questions) && value.answers && value.results && value.completed &&
        (value.phase === "idle" || (value.paper?.count === value.questions.length && Number.isFinite(value.deadline)))) return value;
  } catch { /* Practice also works when browser storage is unavailable. */ }
  return { ...blank, completed: {} };
}
export const usePPLExamStore = create<ExamState>()((set, get) => ({
  ...restored(), busy: false, error: null,
  start: async (id) => {
    if (get().busy || get().phase === "active" || get().phase === "grading") return;
    set({ busy: true, error: null });
    try {
      const { data } = await apiClient.get<{ paper: PPLPaper; questions: Question[] }>(`/api/ppl/exam-papers/${id}`);
      if (data.questions.length !== data.paper.count || data.questions.some(q => q.options.length !== 4)) throw new Error("Invalid paper");
      set({ ...blank, ...data, phase: "active", deadline: Date.now() + data.paper.minutes * 60_000, busy: false });
    } catch { set({ busy: false, error: "Impossible de charger cette épreuve. Réessayez." }); }
  },
  choose: (id, letter) => {
    const s = get();
    if (s.phase !== "active" || Date.now() >= s.deadline || !/^[A-D]$/.test(letter) || !s.questions.some(q => q.id === id)) return;
    set({ answers: { ...s.answers, [id]: letter } });
  },
  finish: async () => {
    const s = get();
    if (!s.paper || s.busy || !["active", "grading"].includes(s.phase)) return;
    // Freeze choices before requesting any correction. Partial results survive retry/reload.
    set({ phase: "grading", busy: true, error: null });
    try {
      for (const q of s.questions) {
        if (get().results[q.id]) continue;
        const selected = s.answers[q.id] || "";
        // Blank answers receive no point. Request their explanation without inventing
        // a correct answer locally; the existing evaluator rejects this sentinel.
        const { data } = await questionsApi.answer(q.id, selected || "__unanswered__");
        set({ results: { ...get().results, [q.id]: { ...data, correct: selected !== "" && data.correct, selected } } });
      }
      const correct = s.questions.filter(q => get().results[q.id].correct).length;
      set({ phase: "finished", busy: false, completed: { ...get().completed,
        [s.paper.id]: { correct, total: s.paper.count, passed: correct * 100 >= s.paper.count * 75 } } });
    } catch { set({ busy: false, error: "Correction interrompue. Réessayez ; vos réponses sont conservées." }); }
  },
  close: () => { if (!get().busy && get().phase === "finished") set({ ...blank, error: null }); },
}));
usePPLExamStore.subscribe((state) => {
  try {
    const { phase, paper, questions, answers, results, completed, deadline } = state;
    sessionStorage.setItem(key, JSON.stringify({ phase, paper, questions, answers, results, completed, deadline }));
  } catch { /* Storage is optional; never interrupt an answer. */ }
});
