export type MeasurementSystem = "metric" | "us" | "imperial";

export function validMeasurementSystem(value: string | null): MeasurementSystem {
  return value === "us" || value === "imperial" ? value : "metric";
}

// Exact definitions. Original values remain visible so rounding never changes
// an exercise, a published value or its canonical answer.
export function measurementText(text: string, system: MeasurementSystem, lang: string, sourceLang = lang): string {
  if (system === "metric") return text;
  const litrePerGallon = system === "us" ? 3.785411784 : 4.54609;
  const gallon = system === "us" ? "US gal" : "imp gal";
  const locale = lang.startsWith("fr") ? "fr-FR" : lang.startsWith("es") ? "es-ES" : "en-GB";
  const format = (value: number) => new Intl.NumberFormat(locale, { maximumFractionDigits: 2 }).format(value);
  // Protect Markdown link targets and code; convert only explicit quantities.
  return text.split(/(\]\([^\n)]*\)|`[^`]*`)/g).map((part) => {
    if (part.startsWith("](") || part.startsWith("`")) return part;
    return part.replace(/(?<![\w.,])(-?\d+(?:[ \u00a0\u202f]\d{3})*(?:[.,]\d+)*?)\s*(kg[·⋅]m|kg|L\/h|L|cm|m)(?![\w/²³])/g,
      (original, raw: string, unit: string) => {
        const compact = raw.replace(/[ \u00a0\u202f]/g, "");
        const number = sourceLang.startsWith("en")
          ? Number(compact.replace(/,(?=\d{3}(?:\D|$))/g, ""))
          : Number(compact.replace(/\.(?=\d{3}(?:\D|$))/g, "").replace(",", "."));
        if (!Number.isFinite(number)) return original;
        const [factor, target] = unit === "kg" ? [1 / 0.45359237, "lb"]
          : unit === "kg·m" || unit === "kg⋅m" ? [1 / (0.45359237 * 0.3048), "lb·ft"]
          : unit === "L" ? [1 / litrePerGallon, gallon]
          : unit === "L/h" ? [1 / litrePerGallon, `${gallon}/h`]
          : unit === "cm" ? [1 / 2.54, "in"] : [1 / 0.3048, "ft"];
        return `${original} (≈ ${format(number * Number(factor))} ${target})`;
      });
  }).join("").replace(/(!\[[^\]\n]*\]\()\/illustrations\/((?:(?:en|es)\/)?[a-zA-Z0-9_-]+\.svg\))/g, `$1/illustrations/units/${system}/$2`);
}
