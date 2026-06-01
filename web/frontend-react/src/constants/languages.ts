import type { SupportedLanguage } from "@/types";

export const SUPPORTED_LANGUAGES: SupportedLanguage[] = [
  { code: "fr", label: "Français", flag: "🇫🇷" },
  { code: "en", label: "English",   flag: "🇬🇧" },
  { code: "de", label: "Deutsch",   flag: "🇩🇪" },
  { code: "es", label: "Español",   flag: "🇪🇸" },
  { code: "it", label: "Italiano",  flag: "🇮🇹" },
];

export const DEFAULT_LANGUAGE = "fr";