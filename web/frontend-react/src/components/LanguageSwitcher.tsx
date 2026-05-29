import { useTranslation } from "react-i18next";
import { Languages } from "lucide-react";

export default function LanguageSwitcher() {
  const { i18n } = useTranslation();

  const toggleLanguage = () => {
    const nextLang = i18n.language === "fr" ? "en" : "fr";
    i18n.changeLanguage(nextLang);
  };

  return (
    <button
      onClick={toggleLanguage}
      className="flex items-center gap-2 px-3 py-2 rounded-lg text-sm font-medium text-slate-400 hover:text-slate-200 hover:bg-slate-800 transition-all duration-200"
      title={i18n.language === "fr" ? "Switch to English" : "Passer en français"}
    >
      <Languages className="w-4 h-4" />
      <span className="uppercase font-bold">{i18n.language}</span>
    </button>
  );
}
