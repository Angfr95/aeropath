import { apiClient } from "./client";
import type { UserStats, AdminStats } from "@/types";

export const statsApi = {
  // GET /api/stats
  user: () => apiClient.get<UserStats>("/api/stats"),

  // GET /api/admin/stats
  admin: () => apiClient.get<AdminStats>("/api/admin/stats"),
};