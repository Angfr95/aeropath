import { create } from "zustand";
import type { User, LoginPayload, RegisterPayload, License } from "@/types";
import { apiClient } from "@/api/client";

// ---------------------------------------------------------------------------
// Types internes au store
// ---------------------------------------------------------------------------
interface AuthState {
  // données
  user: User | null;
  isAuthenticated: boolean;
  isLoading: boolean;

  // actions
  login: (payload: LoginPayload) => Promise<void>;
  register: (payload: RegisterPayload) => Promise<void>;
  logout: () => Promise<void>;
  loadUser: () => Promise<void>;
  updateLang: (lang: string) => Promise<void>;
}

// ---------------------------------------------------------------------------
// Store
// ---------------------------------------------------------------------------
export const useAuthStore = create<AuthState>()((set, get) => ({
  user: null,
  isAuthenticated: false,
  isLoading: false,

  // ---- login ---------------------------------------------------------------
  login: async (payload) => {
    set({ isLoading: true });
    try {
      await apiClient.post("/auth/login", payload);
      // Le cookie httpOnly est posé, on charge les données utilisateur
      const res = await apiClient.get<{ user: User }>("/api/me");
      set({ user: res.data.user, isAuthenticated: true, isLoading: false });
    } catch {
      set({ isLoading: false, user: null, isAuthenticated: false });
      throw new Error("Échec de connexion");
    }
  },

  // ---- register ------------------------------------------------------------
  register: async (payload) => {
    set({ isLoading: true });
    try {
      await apiClient.post("/auth/register", payload);
      // Le cookie httpOnly est posé, on charge les données utilisateur
      const res = await apiClient.get<{ user: User }>("/api/me");
      set({ user: res.data.user, isAuthenticated: true, isLoading: false });
    } catch {
      set({ isLoading: false, user: null, isAuthenticated: false });
      throw new Error("Échec d'inscription");
    }
  },

  // ---- logout --------------------------------------------------------------
  logout: async () => {
    try {
      await apiClient.post("/auth/logout");
    } finally {
      set({ user: null, isAuthenticated: false, isLoading: false });
    }
  },

  // ---- loadUser (appelée au montage) ---------------------------------------
  loadUser: async () => {
    // évite les appels multiples pendant le chargement
    if (get().isLoading) return;

    set({ isLoading: true });
    try {
      const res = await apiClient.get<{ user: User }>("/api/me");
      set({ user: res.data.user, isAuthenticated: true, isLoading: false });
    } catch {
      // non authentifié – on nettoie
      set({ user: null, isAuthenticated: false, isLoading: false });
    }
  },

  // ---- updateLang ----------------------------------------------------------
  updateLang: async (lang) => {
    try {
      await apiClient.patch("/api/me/lang", { lang });
    } catch {
      // silencieux (préférence mineure)
    }
  },
}));