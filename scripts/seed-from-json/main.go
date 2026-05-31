package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ─── Structs compatibles avec le format RÉEL des fichiers JSON ───

type ConceptQuestion struct {
	ConceptID       string   `json:"concept_id"`
	ConceptTitle    string   `json:"concept_title"`
	Level           int      `json:"level"`
	Question        string   `json:"question"`
	QuestionEn      string   `json:"question_en"`
	Options         []string `json:"options"`
	CorrectAnswer   string   `json:"correct_answer"`
	Explanation     string   `json:"explanation"`
	ExplanationEn   string   `json:"explanation_en"`
	FaaNote         string   `json:"faa_note,omitempty"`
	DifficultyScore float64  `json:"difficulty_score,omitempty"`
	Tags            []string `json:"tags,omitempty"`
}

type Concept struct {
	ConceptID       string   `json:"concept_id"`
	Title           string   `json:"title"`
	TitleEn         string   `json:"title_en"`
	Content         string   `json:"content"`
	ContentEn       string   `json:"content_en"`
	Difficulty      string   `json:"difficulty"`
	RiskLevel       string   `json:"risk_level"`
	PhaseOfFlight   []string `json:"phase_of_flight"`
	Tags            []string `json:"tags"`
	FaaNote         string   `json:"faa_note,omitempty"`
	CommonErrors    []string `json:"common_errors,omitempty"`
	RelatedConcepts []string `json:"related_concepts,omitempty"`
}

type Metadata struct {
	ModuleID                 string   `json:"module_id"`
	Program                  string   `json:"program"`
	Subject                  string   `json:"subject"`
	ChapterTitle             string   `json:"chapter_title"`
	ChapterTitleEn           string   `json:"chapter_title_en"`
	LearningObjective        string   `json:"learning_objective"`
	LearningObjectiveEn      string   `json:"learning_objective_en"`
	Difficulty               string   `json:"difficulty"`
	EstimatedDurationMinutes int      `json:"estimated_duration_minutes"`
	Prerequisites            []string `json:"prerequisites,omitempty"`
	Tags                     []string `json:"tags,omitempty"`
}

type ModuleFile struct {
	Metadata  Metadata          `json:"metadata"`
	Concepts  []Concept         `json:"concepts"`
	Questions []ConceptQuestion `json:"questions"`
}

// ─── Helpers ───

func escapeSQL(s string) string {
	s = strings.ReplaceAll(s, "'", "''")
	s = strings.ReplaceAll(s, "\\", "\\\\")
	return s
}

func jsonArrayStrings(arr []string) string {
	if len(arr) == 0 {
		return "[]"
	}
	escaped := make([]string, len(arr))
	for i, v := range arr {
		b, _ := json.Marshal(v)
		escaped[i] = string(b)
	}
	return "[" + strings.Join(escaped, ",") + "]"
}

func jsonOptions(opts []string) string {
	b, _ := json.Marshal(opts)
	return string(b)
}

func deterministicUUID(seed string) string {
	h := 0
	for _, c := range seed {
		h = (h*31 + int(c)) & 0x7FFFFFFF
	}
	return fmt.Sprintf("%08x-0000-4000-8000-%012x", h, h)
}

func main() {
	contentDir := "content/fr/ppl/air_law"
	outputPath := "scripts/seed/seed_from_json.sql"

	entries, err := os.ReadDir(contentDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erreur lecture dossier %s: %v\n", contentDir, err)
		os.Exit(1)
	}

	var jsonFiles []string
	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".json") && !e.IsDir() {
			jsonFiles = append(jsonFiles, filepath.Join(contentDir, e.Name()))
		}
	}

	if len(jsonFiles) == 0 {
		fmt.Fprintf(os.Stderr, "Aucun fichier JSON trouvé dans %s\n", contentDir)
		os.Exit(1)
	}

	fmt.Printf("📂 Fichiers trouvés: %d\n", len(jsonFiles))

	var modules []ModuleFile
	for _, f := range jsonFiles {
		data, err := os.ReadFile(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erreur lecture %s: %v\n", f, err)
			continue
		}
		var mod ModuleFile
		if err := json.Unmarshal(data, &mod); err != nil {
			fmt.Fprintf(os.Stderr, "Erreur parse JSON %s: %v\n", f, err)
			continue
		}
		modules = append(modules, mod)
		fmt.Printf("  ✅ %s — %s (%d concepts, %d questions)\n",
			filepath.Base(f), mod.Metadata.ChapterTitle,
			len(mod.Concepts), len(mod.Questions))
	}

	// ─── Génération SQL ───
	var sb strings.Builder

	sb.WriteString("-- ============================================================================\n")
	sb.WriteString(fmt.Sprintf("-- SEED FROM JSON : %d modules chargés depuis content/fr/ppl/air_law/\n", len(modules)))
	sb.WriteString("-- Généré par scripts/seed-from-json/main.go\n")
	sb.WriteString("-- ============================================================================\n\n")

	sb.WriteString("-- Nettoyage des données existantes\n")
	sb.WriteString("DELETE FROM student_question_history;\n")
	sb.WriteString("DELETE FROM questions;\n")
	sb.WriteString("DELETE FROM lessons;\n")
	sb.WriteString("DELETE FROM user_gamification;\n")
	sb.WriteString("DELETE FROM user_progress;\n")
	sb.WriteString("DELETE FROM achievements;\n")
	sb.WriteString("DELETE FROM leaderboard_weekly;\n")
	sb.WriteString("DELETE FROM user_quests;\n")
	sb.WriteString("DELETE FROM spaced_repetition;\n")
	sb.WriteString("DELETE FROM app_config;\n\n")

	// ─── 1. APP CONFIG ───
	sb.WriteString("-- ============================================================================\n")
	sb.WriteString("-- 1. APP CONFIG\n")
	sb.WriteString("-- ============================================================================\n")
	sb.WriteString("INSERT INTO app_config (config_key, config_value) VALUES\n")
	sb.WriteString("('adaptive_learning', '{\n")
	sb.WriteString("    \"level_1\": {\"pass_threshold\": 0.80, \"hint_enabled\": true, \"time_limit_sec\": null, \"immediate_feedback\": true},\n")
	sb.WriteString("    \"level_2\": {\"pass_threshold\": 0.75, \"hint_enabled\": false, \"time_limit_sec\": null, \"immediate_feedback\": true},\n")
	sb.WriteString("    \"level_3\": {\"pass_threshold\": 0.75, \"hint_enabled\": false, \"time_limit_sec\": 45, \"immediate_feedback\": false}\n")
	sb.WriteString("}');\n\n")

	sb.WriteString("INSERT INTO app_config (config_key, config_value) VALUES\n")
	sb.WriteString("('gamification_rules', '{\n")
	sb.WriteString("    \"hearts_enabled\": true, \"hearts_total\": 5,\n")
	sb.WriteString("    \"heart_penalty_level2\": 1, \"heart_penalty_level3\": 2,\n")
	sb.WriteString("    \"xp_per_lesson\": 10, \"xp_per_qcm_level1\": 20, \"xp_per_qcm_level2\": 30, \"xp_per_qcm_level3\": 50,\n")
	sb.WriteString("    \"streak_enabled\": true, \"leaderboard_weekly\": true\n")
	sb.WriteString("}');\n\n")

	sb.WriteString("INSERT INTO app_config (config_key, config_value) VALUES\n")
	sb.WriteString("('exam_config', '{\n")
	sb.WriteString("    \"ppl.air_law\": {\"questions_count\": 16, \"time_minutes\": 30, \"passing_score_percent\": 75, \"randomized\": true, \"no_backward\": true}\n")
	sb.WriteString("}');\n\n")

	// ─── 2. LESSONS ───
	sb.WriteString("-- ============================================================================\n")
	sb.WriteString("-- 2. LESSONS\n")
	sb.WriteString("-- ============================================================================\n")

	for modIdx, mod := range modules {
		lessonUUID := deterministicUUID(fmt.Sprintf("lesson-%s-%s", mod.Metadata.Program, mod.Metadata.ModuleID))

		difficulty := 1
		switch mod.Metadata.Difficulty {
		case "facile", "beginner":
			difficulty = 1
		case "moyen", "intermediate":
			difficulty = 2
		case "difficile", "advanced":
			difficulty = 3
		}

		tagsJSON := jsonArrayStrings(mod.Metadata.Tags)

		// Contenu = concaténation de tous les concepts
		var contentFr, contentEn strings.Builder
		for _, c := range mod.Concepts {
			if c.Content != "" {
				contentFr.WriteString("## ")
				contentFr.WriteString(c.Title)
				contentFr.WriteString("\n\n")
				contentFr.WriteString(c.Content)
				contentFr.WriteString("\n\n")
			}
			if c.ContentEn != "" {
				contentEn.WriteString("## ")
				contentEn.WriteString(c.TitleEn)
				contentEn.WriteString("\n\n")
				contentEn.WriteString(c.ContentEn)
				contentEn.WriteString("\n\n")
			}
		}

		// Learning objectives (depuis le metadata)
		var loJSON = "[]"
		if mod.Metadata.LearningObjective != "" {
			loJSON = jsonArrayStrings([]string{mod.Metadata.LearningObjective})
		}

		sb.WriteString("INSERT INTO lessons (id, license, category, theme, title_fr, title_en, content_fr, content_en, difficulty, order_index, level, duration_minutes, tags, learning_objectives) VALUES\n")
		sb.WriteString(fmt.Sprintf("('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', %d, %d, %d, %d, '%s', '%s');\n\n",
			lessonUUID,
			mod.Metadata.Program,
			strings.ToLower(mod.Metadata.Subject),
			strings.ToLower(mod.Metadata.Subject),
			escapeSQL(mod.Metadata.ChapterTitle),
			escapeSQL(mod.Metadata.ChapterTitleEn),
			escapeSQL(contentFr.String()),
			escapeSQL(contentEn.String()),
			difficulty,
			modIdx+1,
			1,
			mod.Metadata.EstimatedDurationMinutes,
			escapeSQL(tagsJSON),
			escapeSQL(loJSON)))

		// ─── 3. QUESTIONS ───
		if len(mod.Questions) > 0 {
			sb.WriteString(fmt.Sprintf("-- Questions for lesson %s (%d questions)\n", lessonUUID, len(mod.Questions)))

			for qIdx, q := range mod.Questions {
				qid := deterministicUUID(fmt.Sprintf("q-%s-%s-%d", mod.Metadata.ModuleID, q.ConceptID, qIdx))
				subtopic := q.ConceptTitle
				if len(subtopic) > 100 {
					subtopic = subtopic[:100]
				}

				optsJSON := jsonOptions(q.Options)
				qTagsJSON := jsonArrayStrings(q.Tags)

				sb.WriteString("INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, answer_key, explanation_fr, explanation_en, faa_note_fr, faa_note_en, tags, difficulty_score) VALUES\n")
				sb.WriteString(fmt.Sprintf("('%s','%s','%s','%s','%s','%s',%d,%d,'%s','%s','%s','%s','%s','%s','%s','%s','%s',%.2f);\n",
					qid,
					lessonUUID,
					mod.Metadata.Program,
					strings.ToLower(mod.Metadata.Subject),
					strings.ToLower(mod.Metadata.Subject),
					escapeSQL(subtopic),
					q.Level,
					q.Level,
					escapeSQL(q.Question),
					escapeSQL(q.QuestionEn),
					escapeSQL(optsJSON),
					escapeSQL(q.CorrectAnswer),
					escapeSQL(q.Explanation),
					escapeSQL(q.ExplanationEn),
					"",
					escapeSQL(q.FaaNote),
					escapeSQL(qTagsJSON),
					q.DifficultyScore))
			}
			sb.WriteString("\n")
		}
	}

	// ─── 4. ÉTUDIANT DE TEST ───
	sb.WriteString("-- ============================================================================\n")
	sb.WriteString("-- 3. ÉTUDIANT DE TEST\n")
	sb.WriteString("-- ============================================================================\n")
	sb.WriteString("INSERT INTO students (id, email, password_hash, lang, preferred_license, hearts, xp, streak, user_level) VALUES\n")
	sb.WriteString("('00000000-0000-4000-8000-000000000001', 'test@aeropath.app', 'test123', 'fr', 'PPL', 5, 0, 0, 1);\n\n")

	sb.WriteString("INSERT INTO user_gamification (user_id, hearts, xp, streak, level, preferred_language, preferred_license) VALUES\n")
	sb.WriteString("('00000000-0000-4000-8000-000000000001', 5, 0, 0, 1, 'fr', 'PPL');\n\n")

	// ─── Stats ───
	totalQuestions := 0
	totalConcepts := 0
	for _, mod := range modules {
		totalConcepts += len(mod.Concepts)
		totalQuestions += len(mod.Questions)
	}

	sb.WriteString("-- ============================================================================\n")
	sb.WriteString(fmt.Sprintf("-- RÉSUMÉ: %d modules, %d concepts, %d questions\n", len(modules), totalConcepts, totalQuestions))
	sb.WriteString("-- ============================================================================\n")

	// Écriture du fichier
	if err := os.MkdirAll("scripts/seed", 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Erreur création dossier scripts/seed: %v\n", err)
		os.Exit(1)
	}
	if err := os.WriteFile(outputPath, []byte(sb.String()), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Erreur écriture %s: %v\n", outputPath, err)
		os.Exit(1)
	}

	fmt.Printf("\n✅ Seed généré: %s\n", outputPath)
	fmt.Printf("📊 %d modules, %d concepts, %d questions\n", len(modules), totalConcepts, totalQuestions)
	fmt.Println("\n🚀 Pour appliquer: make seed/from-json")
	fmt.Println("   Ou: psql \"postgres://aeropath:changer_moi_postgres@localhost:5432/aeropath\" < scripts/seed/seed_from_json.sql")
}
