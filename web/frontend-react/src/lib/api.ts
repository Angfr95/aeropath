const API_BASE = "/api";

async function request<T>(
  endpoint: string,
  options?: RequestInit
): Promise<T> {
  const res = await fetch(`${API_BASE}${endpoint}`, {
    headers: {
      "Content-Type": "application/json",
      ...options?.headers,
    },
    ...options,
  });

  if (!res.ok) {
    const error = await res.json().catch(() => ({ error: res.statusText }));
    throw new Error(error.error || `HTTP ${res.status}`);
  }

  return res.json();
}

// === Questions ===
export interface Question {
  id: string;
  question_fr: string;
  question_en: string;
  options: string[];
  answer_key: string;
  explanation_fr: string;
  explanation_en: string;
  license: string;
  category: string;
  theme: string;
  subtopic: string;
  difficulty: number;
  reference: string;
}

export interface PaginatedResponse<T> {
  data: T[];
  total: number;
  page: number;
  page_size: number;
  total_pages: number;
}

export const questionsApi = {
  list: (params?: {
    page?: number;
    page_size?: number;
    license?: string;
    category?: string;
    theme?: string;
  }) => {
    const searchParams = new URLSearchParams();
    if (params?.page) searchParams.set("page", String(params.page));
    if (params?.page_size) searchParams.set("page_size", String(params.page_size));
    if (params?.license) searchParams.set("license", params.license);
    if (params?.category) searchParams.set("category", params.category);
    if (params?.theme) searchParams.set("theme", params.theme);
    return request<PaginatedResponse<Question>>(
      `/questions?${searchParams.toString()}`
    );
  },

  getById: (id: string) => request<Question>(`/questions/${id}`),

  search: (q: string) =>
    request<Question[]>(`/questions/search?q=${encodeURIComponent(q)}`),

  random: (license?: string) => {
    const params = license ? `?license=${license}` : "";
    return request<Question>(`/questions/random${params}`);
  },

  answer: (questionId: string, answer: string) =>
    request<{
      correct: boolean;
      correct_answer: string;
      explanation_fr: string;
      explanation_en: string;
    }>(`/questions/answer`, {
      method: "POST",
      body: JSON.stringify({ question_id: questionId, answer }),
    }),

  byLicense: (license: string) =>
    request<Question[]>(`/questions/by-license/${license}`),

  byCategory: (category: string) =>
    request<Question[]>(`/questions/by-category/${category}`),

  byTheme: (theme: string) =>
    request<Question[]>(`/questions/by-theme/${theme}`),

  byDifficulty: (level: number) =>
    request<Question[]>(`/questions/by-difficulty/${level}`),

  bySubtopic: (subtopic: string) =>
    request<Question[]>(`/questions/by-subtopic/${encodeURIComponent(subtopic)}`),
};

// === Lessons ===
export interface Lesson {
  id: string;
  title_fr: string;
  title_en: string;
  content_fr: string;
  content_en: string;
  license: string;
  category: string;
  theme: string;
  difficulty: number;
  order_index: number;
}

export const lessonsApi = {
  list: (params?: { page?: number; page_size?: number }) => {
    const searchParams = new URLSearchParams();
    if (params?.page) searchParams.set("page", String(params.page));
    if (params?.page_size) searchParams.set("page_size", String(params.page_size));
    return request<PaginatedResponse<Lesson>>(
      `/lessons?${searchParams.toString()}`
    );
  },

  getById: (id: string) => request<Lesson>(`/lessons/${id}`),

  byLicense: (license: string) =>
    request<Lesson[]>(`/lessons/by-license/${license}`),

  byCategory: (category: string) =>
    request<Lesson[]>(`/lessons/by-category/${category}`),
};

// === History ===
export interface AnswerHistory {
  id: string;
  question_id: string;
  question_fr: string;
  answer: string;
  correct: boolean;
  license: string;
  category: string;
  theme: string;
  answered_at: string;
}

export const historyApi = {
  list: (params?: { page?: number; page_size?: number }) => {
    const searchParams = new URLSearchParams();
    if (params?.page) searchParams.set("page", String(params.page));
    if (params?.page_size) searchParams.set("page_size", String(params.page_size));
    return request<PaginatedResponse<AnswerHistory>>(
      `/history?${searchParams.toString()}`
    );
  },
};

// === Stats ===
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

export const statsApi = {
  admin: () => request<AdminStats>("/admin/stats"),
};

// === Recommendations ===
export interface Recommendation {
  progression_percent: number;
  weak_topics: { theme: string; score: number; priority: string }[];
  due_cards: {
    question_id: string;
    question_fr: string;
    next_review: string;
    interval_days: number;
  }[];
  next_milestone: string;
  mastery_by_license: { license: string; score: number }[];
}

export const recommendationsApi = {
  get: () => request<Recommendation>("/recommendations"),
};

// === WebSocket ===
export function createWebSocket(): WebSocket {
  const protocol = window.location.protocol === "https:" ? "wss:" : "ws:";
  const wsUrl = `${protocol}//${window.location.host}/api/ws`;
  return new WebSocket(wsUrl);
}
