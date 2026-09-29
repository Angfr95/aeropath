import type { Question } from "@/types";

export function localizedText(lang: string, fr?: string, en?: string, es?: string): string {
  return (lang.startsWith("es") ? es || fr || en : lang.startsWith("en") ? en || fr : fr || en) || "";
}

// The position is the answer identity: translated choices must never be reordered.
export function localizedOptions(question: Pick<Question, "options" | "options_en" | "options_es">, lang: string): string[] {
  const canonical = question.options.map((option) => typeof option === "string" ? option : option.text);
  const translated = lang.startsWith("es") ? question.options_es : lang.startsWith("en") ? question.options_en : undefined;
  return translated?.length === canonical.length && translated.every((text) => text.trim())
    ? translated : canonical;
}

export function questionPreview(text: string): string {
  return text.replace(/!\[[^\]]*\]\([^)]*\)/g, "").trim();
}

const labels: Record<string, string> = {
  "Leçons": "Lessons",
  "Choisis une licence": "Choose a licence",
  "Dernières leçons": "Latest lessons",
  "Aucune leçon trouvée": "No lessons found",
  "← Retour aux licences": "← Back to licences",
  "Choisis une catégorie": "Choose a category",
  "Toutes les leçons": "All lessons",
  "Aucune leçon pour cette licence": "No lessons for this licence",
  "← Retour à": "← Back to",
  "Aucune leçon dans cette catégorie": "No lessons in this category",
  "Niv.": "Level",
  "← Retour": "← Back",
  "Erreur de chargement": "Unable to load content",
  "Chargement du quiz…": "Loading quiz…",
  "📝 Quiz sur cette leçon": "📝 Take this lesson's quiz",
  "Pas de quiz en cours": "No active quiz",
  "Lancer un quiz": "Start a quiz",
  "Choisir une leçon": "Choose a lesson",
  "Quiz terminé !": "Quiz complete!",
  "bonnes réponses": "correct answers",
  "Résultat non enregistré.": "Result not saved.",
  "Connectez-vous": "Sign in",
  "pour suivre vos prochains quiz.": "to track your next quizzes.",
  "Refaire ce quiz": "Retake this quiz",
  "Revoir la leçon": "Review the lesson",
  "Accueil": "Home",
  "Vérification de la réponse…": "Checking your answer…",
  "✅ Correct !": "✅ Correct!",
  "❌ Faux": "❌ Incorrect",
  "Réponse correcte :": "Correct answer:",
  "📊 Voir les résultats": "📊 View results",
  "➡️ Question suivante": "➡️ Next question",
  "🧠 Question aléatoire": "🧠 Random quiz",
  "Dernières questions": "Latest questions",
  "Aucune question trouvée": "No questions found",
  "Toutes les questions": "All questions",
  "Aucune question pour cette licence": "No questions for this licence",
  "🧠 Quiz sur cette catégorie": "🧠 Take a quiz in this category",
  "Aucune question dans cette catégorie": "No questions in this category",
  "Chargement...": "Loading...",
  "Historique": "History",
  "Recommandations": "Recommendations",
  "Déconnecté": "Signed out",
  "Déconnexion": "Sign out",
  "Connexion": "Sign in",
  "Pilote Privé": "Private Pilot",
  "Pilote Privé Léger": "Light Aircraft Pilot",
  "Pilote Professionnel": "Commercial Pilot",
  "Transport Aérien": "Airline Transport Pilot",
  "Vol aux Instruments": "Instrument Rating",
  "Météorologie": "Meteorology",
  "Réglementation": "Air law",
  "Connaissance Aéronef": "Aircraft knowledge",
  "Facteurs Humains": "Human performance",
  "Procédures": "Operational procedures",
  "Principes du Vol": "Principles of flight",
  "Planification": "Flight planning",
  "Urgences": "Emergencies",
  "Masse & Centrage": "Mass and balance",
  "Procédures Radio": "Radio procedures",
  "Aucune question disponible pour ce quiz.": "No questions available for this quiz.",
  "Impossible de charger le quiz. Réessayez depuis la leçon ou la liste des cours.": "Unable to load the quiz. Try again from the lesson or course list.",
  "La correction n’a pas pu être chargée. Réessayez votre réponse.": "Unable to load the correction. Try submitting your answer again."
};
export function learningLabel(lang: string, text: string): string {
  return lang.startsWith("en") ? labels[text] || text : text;
}
