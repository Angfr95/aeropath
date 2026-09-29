import { measurementText } from "./measurementSystem";
import { useUIStore } from "@/store/uiStore";
import { localizedText, localizedOptions, learningLabel } from "./learningLocale";
import type { Question } from "@/types";

export function useLearningLocale() {
  const lang = useUIStore((state) => state.currentLang);
  const system = useUIStore((state) => state.measurementSystem);
  const display = (value: string, sourceLang: string) => measurementText(value, system, lang, sourceLang);
  return {
    text: (fr?: string, en?: string, es?: string) => {
      const sourceLang = lang.startsWith("es") && es ? "es"
        : lang.startsWith("en") && en ? "en" : fr ? "fr" : "en";
      return display(localizedText(lang, fr, en, es), sourceLang);
    },
    label: (value: string) => learningLabel(lang, value),
    choices: (question: Pick<Question, "options" | "options_en" | "options_es">) => {
      const options = localizedOptions(question, lang);
      const sourceLang = options === question.options_es ? "es" : options === question.options_en ? "en" : "fr";
      return options.map((option) => display(option, sourceLang));
    },
  };
}
