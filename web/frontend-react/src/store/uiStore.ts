import { create } from "zustand";
import type { Toast, ToastType } from "@/types";
import { DEFAULT_LANGUAGE } from "@/constants/languages";

// ---------------------------------------------------------------------------
// Types internes
// ---------------------------------------------------------------------------
interface UIState {
  // langue
  currentLang: string;

  // toasts
  toasts: Toast[];

  // actions
  setLang: (lang: string) => void;
  addToast: (message: string, type?: ToastType, durationMs?: number) => void;
  dismissToast: (id: string) => void;
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------
let toastCounter = 0;

function generateToastId(): string {
  toastCounter++;
  return `toast-${Date.now()}-${toastCounter}`;
}

// ---------------------------------------------------------------------------
// Store
// ---------------------------------------------------------------------------
export const useUIStore = create<UIState>()((set, get) => ({
  currentLang: localStorage.getItem("aeropath_lang") || DEFAULT_LANGUAGE,
  toasts: [],

  // ---- setLang -------------------------------------------------------------
  setLang: (lang) => {
    localStorage.setItem("aeropath_lang", lang);
    set({ currentLang: lang });
  },

  // ---- addToast ------------------------------------------------------------
  addToast: (message, type = "info", durationMs = 3_000) => {
    const id = generateToastId();
    const toast: Toast = { id, message, type, durationMs };
    set((s) => ({ toasts: [...s.toasts, toast] }));

    // auto-dismiss
    setTimeout(() => {
      set((s) => ({ toasts: s.toasts.filter((t) => t.id !== id) }));
    }, durationMs);
  },

  // ---- dismissToast --------------------------------------------------------
  dismissToast: (id) => {
    set((s) => ({ toasts: s.toasts.filter((t) => t.id !== id) }));
  },
}));