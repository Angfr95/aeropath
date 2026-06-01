import { useUIStore } from "@/store/uiStore";
import { X } from "lucide-react";

const colorMap = {
  info: "bg-blue-600",
  success: "bg-green-600",
  error: "bg-red-600",
};

export default function ToastContainer() {
  const toasts = useUIStore((s) => s.toasts);
  const dismiss = useUIStore((s) => s.dismissToast);

  if (toasts.length === 0) return null;

  return (
    <div className="fixed bottom-4 right-4 z-50 flex flex-col gap-2 pointer-events-none">
      {toasts.map((t) => (
        <div
          key={t.id}
          className={`${colorMap[t.type]} text-white px-4 py-2 rounded-lg shadow-lg pointer-events-auto flex items-center gap-3 animate-slide-up`}
        >
          <span className="text-sm font-medium flex-1">{t.message}</span>
          <button
            onClick={() => dismiss(t.id)}
            className="text-white/70 hover:text-white transition"
            aria-label="Fermer"
          >
            <X className="w-4 h-4" />
          </button>
        </div>
      ))}
    </div>
  );
}