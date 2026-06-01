import { useState, useRef, useEffect } from "react";
import { useUIStore } from "@/store/uiStore";
import { useAuthStore } from "@/store/authStore";
import { SUPPORTED_LANGUAGES } from "@/constants/languages";
import { Check, ChevronDown } from "lucide-react";

export default function LanguageSelector() {
  const currentLang = useUIStore((s) => s.currentLang);
  const setLang = useUIStore((s) => s.setLang);
  const updateLang = useAuthStore((s) => s.updateLang);
  const [open, setOpen] = useState(false);
  const ref = useRef<HTMLDivElement>(null);

  const current = SUPPORTED_LANGUAGES.find((l) => l.code === currentLang) ?? SUPPORTED_LANGUAGES[0];

  // Ferme le dropdown au clic extérieur
  useEffect(() => {
    function handleClick(e: MouseEvent) {
      if (ref.current && !ref.current.contains(e.target as Node)) {
        setOpen(false);
      }
    }
    document.addEventListener("mousedown", handleClick);
    return () => document.removeEventListener("mousedown", handleClick);
  }, []);

  function handleSelect(code: string) {
    setLang(code);
    updateLang(code);
    setOpen(false);
  }

  return (
    <div ref={ref} className="relative">
      <button
        onClick={() => setOpen((o) => !o)}
        className="flex items-center gap-1.5 px-2.5 py-1.5 rounded-lg text-sm text-slate-300 hover:bg-slate-700 transition"
        title="Changer la langue"
      >
        <span>{current.flag}</span>
        <span className="hidden sm:inline">{current.label}</span>
        <ChevronDown className="w-3 h-3 ml-0.5" />
      </button>

      {open && (
        <div className="absolute right-0 mt-1 w-44 bg-slate-800 border border-slate-700 rounded-xl shadow-xl z-50">
          <div className="py-1">
            {SUPPORTED_LANGUAGES.map((l) => (
              <button
                key={l.code}
                onClick={() => handleSelect(l.code)}
                className={`w-full flex items-center gap-3 px-4 py-2.5 text-sm text-slate-300 hover:bg-slate-700 hover:text-white transition ${
                  l.code === currentLang ? "bg-slate-700/50 text-white" : ""
                }`}
              >
                <span className="text-lg">{l.flag}</span>
                <span>{l.label}</span>
                {l.code === currentLang && (
                  <Check className="ml-auto w-4 h-4 text-blue-400" />
                )}
              </button>
            ))}
          </div>
        </div>
      )}
    </div>
  );
}