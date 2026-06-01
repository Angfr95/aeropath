import { apiClient } from "./client";
import type { HistoryEntry } from "@/types";

export const historyApi = {
  // GET /api/history?limit=100
  list: (limit = 100) =>
    apiClient.get<{ entries: HistoryEntry[] }>("/api/history", { params: { limit } }),
};