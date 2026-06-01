import { apiClient } from "./client";
import type { Lesson, Question } from "@/types";

// ---------------------------------------------------------------------------
// API Leçons
// ---------------------------------------------------------------------------

export const lessonsApi = {
  // GET  /api/lessons?limit=10
  list: (limit = 10) =>
    apiClient.get<{ data: Lesson[] } | Lesson[]>("/api/lessons", { params: { limit } }),

  // GET  /api/lessons/count
  count: () =>
    apiClient.get<{ count: number }>("/api/lessons/count"),

  // GET  /api/lessons/:id
  getById: (id: string) =>
    apiClient.get<Lesson>(`/api/lessons/${id}`),

  // GET  /api/lessons/license/:license
  byLicense: (license: string) =>
    apiClient.get<{ data: Lesson[] } | Lesson[]>(`/api/lessons/license/${license}`),

  // GET  /api/lessons/by-license/:license/category/:category
  byLicenseCategory: (license: string, category: string) =>
    apiClient.get<{ data: Lesson[] } | Lesson[]>(
      `/api/lessons/by-license/${license}/category/${category}`,
    ),

  // GET  /api/lessons/:id/quiz
  quiz: (lessonId: string) =>
    apiClient.get<{ questions: Question[] }>(`/api/lessons/${lessonId}/quiz`),
};