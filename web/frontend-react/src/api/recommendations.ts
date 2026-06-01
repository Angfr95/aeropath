import { apiClient } from "./client";
import type { Recommendations } from "@/types";

export const recommendationsApi = {
  // GET /api/recommendations
  get: () => apiClient.get<Recommendations>("/api/recommendations"),
};