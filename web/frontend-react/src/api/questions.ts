import { apiClient } from "./client";
import type { Question, AnswerResult, PaginatedResponse } from "@/types";

// ---------------------------------------------------------------------------
// API Questions (typée, basée sur les endpoints de app.js)
// ---------------------------------------------------------------------------

export const questionsApi = {
  // GET  /api/questions?limit=10
  list: (limit = 10) =>
    apiClient.get<{ questions: Question[] }>("/api/questions", { params: { limit } }),

  // GET  /api/questions/count
  count: () =>
    apiClient.get<{ count: number }>("/api/questions/count"),

  // GET  /api/questions/:id
  getById: (id: string) =>
    apiClient.get<Question>(`/api/questions/${id}`),

  // GET  /api/questions/random?limit=N
  random: (limit = 5) =>
    apiClient.get<{ questions: Question[] }>("/api/questions/random", { params: { limit } }),

  // GET  /api/questions/by-license/:license
  byLicense: (license: string) =>
    apiClient.get<{ questions: Question[] }>(`/api/questions/by-license/${license}`),

  // GET  /api/questions/by-license/:license/category/:category
  byLicenseCategory: (license: string, category: string) =>
    apiClient.get<{ questions: Question[] }>(
      `/api/questions/by-license/${license}/category/${category}`,
    ),

  // GET  /api/exam/license/:license/category/:category  (quiz existant)
  examByCategory: (license: string, category: string) =>
    apiClient.get<{ questions: Question[] }>(
      `/api/exam/license/${license}/category/${category}`,
    ),

  // POST /api/questions/answer
  answer: (questionId: string, answer: string) =>
    apiClient.post<AnswerResult>("/api/questions/answer", {
      question_id: questionId,
      answer,
    }),
};