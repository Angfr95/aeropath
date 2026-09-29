import { useUIStore } from "@/store/uiStore";
import { validMeasurementSystem } from "@/utils/measurementSystem";

export default function MeasurementSelector() {
  const lang = useUIStore((state) => state.currentLang);
  const system = useUIStore((state) => state.measurementSystem);
  const setSystem = useUIStore((state) => state.setMeasurementSystem);
  const es = lang.startsWith("es"), en = lang.startsWith("en");
  const title = es ? "Sistema de unidades" : en ? "Measurement system" : "Système d’unités";
  return (
    <label className="flex flex-col gap-0.5 text-xs text-slate-400 min-w-0">
      <span>{title}</span>
      <select aria-label={title} value={system} onChange={(event) => setSystem(validMeasurementSystem(event.target.value))}
        className="bg-slate-800 text-slate-200 border border-slate-600 rounded-lg px-2 py-1.5 max-w-full">
        <option value="metric">{es ? "Métrico" : en ? "Metric" : "Métrique"} · kg, L, m</option>
        <option value="us">US · lb, US gal, ft</option>
        <option value="imperial">{es ? "Británico" : en ? "British" : "Britannique"} · lb, imp gal, ft</option>
      </select>
    </label>
  );
}
