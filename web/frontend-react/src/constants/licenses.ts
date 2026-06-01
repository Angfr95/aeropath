import type { LicenseInfo, CategoryInfo } from "@/types";

export const LICENSES: LicenseInfo[] = [
  { id: "PPL",  label: "PPL",  icon: "✈️",  desc: "Pilote Privé" },
  { id: "LAPL", label: "LAPL", icon: "✈️",  desc: "Pilote Privé Léger" },
  { id: "CPL",  label: "CPL",  icon: "🛩️",  desc: "Pilote Professionnel" },
  { id: "ATPL", label: "ATPL", icon: "✈️",  desc: "Transport Aérien" },
  { id: "IR",   label: "IR",   icon: "✈️",  desc: "Vol aux Instruments" },
];

export const CATEGORIES: CategoryInfo[] = [
  { id: "meteorology",            label: "Météorologie",          icon: "🌦️" },
  { id: "navigation",             label: "Navigation",            icon: "🧭" },
  { id: "airlaw",                 label: "Réglementation",        icon: "📋" },
  { id: "aircraft_general",       label: "Connaissance Aéronef",  icon: "🔧" },
  { id: "performance",            label: "Performance",           icon: "📈" },
  { id: "human_performance",      label: "Facteurs Humains",      icon: "🧑" },
  { id: "operational_procedures", label: "Procédures",            icon: "📋" },
  { id: "communications",         label: "Communications",        icon: "📡" },
  { id: "principles_of_flight",   label: "Principes du Vol",      icon: "✈️" },
  { id: "flight_planning",        label: "Planification",         icon: "📋" },
  { id: "instrumentation",        label: "Instruments",            icon: "🖲️" },
  { id: "emergency",              label: "Urgences",              icon: "🆘" },
  { id: "mass_and_balance",       label: "Masse & Centrage",      icon: "📋" },
  { id: "radio_procedure",        label: "Procédures Radio",      icon: "📡" },
];