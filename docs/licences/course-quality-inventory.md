# Inventaire de qualité du catalogue

Contrôle du 2026-09-29 : 148 chapitres, 2220 questions.

**Ce relevé automatise le triage, pas la validation pédagogique.** Les seuils de longueur signalent des textes à relire ; ajouter des mots ne suffit pas à les améliorer. Les renvois de contexte sont des candidats à vérifier, pas des erreurs certaines.

La relecture doit établir : exactitude et hypothèses ; objectifs précis ; progression et explications ; exemples résolus puis transfert autonome ; distracteurs plausibles et correction des erreurs ; pertinence et lisibilité des SVG ; équivalence des langues et unités ; couverture du programme public. Aucun score global ne remplace ces preuves.

Priorité de travail : français et anglais. Les autres langues sont reportées à la demande de l’utilisateur ; les lacunes espagnoles restent dans le JSON, hors des alertes actives. Les durées restent estimées et les références de matière ne démontrent pas une couverture objectif par objectif.

## Synthèse FR/EN

| Parcours | Chapitres | Questions | Textes FR < 500 mots | Avec correction courte |
| --- | ---: | ---: | ---: | ---: |
| PPL | 63 | 945 | 62 | 62 |
| LAPL | 63 | 945 | 62 | 62 |
| CPL | 7 | 105 | 4 | 4 |
| ATPL | 7 | 105 | 6 | 6 |
| IR | 8 | 120 | 6 | 6 |

## Liste de travail par chapitre

Les numéros de questions concernés et les concepts sans question liée figurent dans [le relevé JSON](course-quality-inventory.json). `—` signifie aucune alerte automatisée, jamais validation complète.

| Chapitre | Mots FR / EN | Alertes à examiner |
| --- | ---: | --- |
| [PPL_AIRLAW_001](../../content/fr/ppl/air_law/PPL_AIRLAW_001_icao.json) | 229 / 192 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_AIRLAW_002](../../content/fr/ppl/air_law/PPL_AIRLAW_002_easa_faa.json) | 325 / 278 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_AIRLAW_003](../../content/fr/ppl/air_law/PPL_AIRLAW_003_certificats.json) | 267 / 230 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_AIRLAW_004](../../content/fr/ppl/air_law/PPL_AIRLAW_004_airspace.json) | 174 / 154 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_AIRLAW_005](../../content/fr/ppl/air_law/PPL_AIRLAW_005_flight_rules.json) | 168 / 147 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_AIRLAW_006](../../content/fr/ppl/air_law/PPL_AIRLAW_006_sera.json) | 159 / 139 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_AIRLAW_007](../../content/fr/ppl/air_law/PPL_AIRLAW_007_flight_planning.json) | 164 / 135 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_AIRLAW_008](../../content/fr/ppl/air_law/PPL_AIRLAW_008_ats.json) | 169 / 144 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_AIRLAW_009](../../content/fr/ppl/air_law/PPL_AIRLAW_009_emergencies.json) | 170 / 147 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_AIRLAW_010](../../content/fr/ppl/air_law/PPL_AIRLAW_010_documentation.json) | 159 / 134 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_AIRLAW_011](../../content/fr/ppl/air_law/PPL_AIRLAW_011_application.json) | 327 / 272 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_AIRLAW_012](../../content/fr/ppl/air_law/PPL_AIRLAW_012_application.json) | 400 / 383 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_AIRLAW_013](../../content/fr/ppl/air_law/PPL_AIRLAW_013_application.json) | 456 / 407 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_AC_001](../../content/fr/ppl/aircraft_general/PPL_AC_001_aircraft_structure.json) | 186 / 165 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_AC_002](../../content/fr/ppl/aircraft_general/PPL_AC_002_powerplant.json) | 186 / 164 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_AC_003](../../content/fr/ppl/aircraft_general/PPL_AC_003_systems.json) | 169 / 151 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_AC_004](../../content/fr/ppl/aircraft_general/PPL_AC_004_instruments.json) | 160 / 149 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_AC_005](../../content/fr/ppl/aircraft_general/PPL_AC_005_electronics.json) | 165 / 137 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_AC_006](../../content/fr/ppl/aircraft_general/PPL_AC_006_performance.json) | 170 / 134 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_AC_007](../../content/fr/ppl/aircraft_general/PPL_AC_007_flight_controls.json) | 177 / 157 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_AC_008](../../content/fr/ppl/aircraft_general/PPL_AC_008_limitations.json) | 167 / 127 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_COM_001](../../content/fr/ppl/communications/PPL_COM_001_course.json) | 158 / 134 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_COM_002](../../content/fr/ppl/communications/PPL_COM_002_course.json) | 163 / 136 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_EMG_001](../../content/fr/ppl/emergency/PPL_EMG_001_course.json) | 289 / 247 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_FP_001](../../content/fr/ppl/flight_planning/PPL_FP_001_course.json) | 168 / 141 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_HP_001](../../content/fr/ppl/human_performance/PPL_HP_001_course.json) | 172 / 143 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_HP_002](../../content/fr/ppl/human_performance/PPL_HP_002_course.json) | 166 / 149 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_HP_003](../../content/fr/ppl/human_performance/PPL_HP_003_application.json) | 336 / 279 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_INS_001](../../content/fr/ppl/instrumentation/PPL_INS_001_course.json) | 298 / 238 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_INS_002](../../content/fr/ppl/instrumentation/PPL_INS_002_course.json) | 269 / 253 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_MB_001](../../content/fr/ppl/mass_and_balance/PPL_MB_001_course.json) | 172 / 151 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_MB_002](../../content/fr/ppl/mass_and_balance/PPL_MB_002_application.json) | 890 / 816 | — |
| [PPL_MET_001](../../content/fr/ppl/meteo/PPL_MET_001_atmosphere.json) | 174 / 154 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_MET_002](../../content/fr/ppl/meteo/PPL_MET_002_pressure.json) | 175 / 151 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_MET_003](../../content/fr/ppl/meteo/PPL_MET_003_temperature.json) | 171 / 146 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_MET_004](../../content/fr/ppl/meteo/PPL_MET_004_wind.json) | 176 / 148 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_MET_005](../../content/fr/ppl/meteo/PPL_MET_005_clouds.json) | 180 / 142 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_MET_006](../../content/fr/ppl/meteo/PPL_MET_006_visibility.json) | 160 / 128 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_MET_007](../../content/fr/ppl/meteo/PPL_MET_007_air_masses.json) | 161 / 133 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_MET_008](../../content/fr/ppl/meteo/PPL_MET_008_hazards.json) | 166 / 151 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_MET_009](../../content/fr/ppl/meteo/PPL_MET_009_metar_taf.json) | 174 / 159 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_MET_010](../../content/fr/ppl/meteo/PPL_MET_010_flight_weather.json) | 168 / 139 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_NAV_001](../../content/fr/ppl/navigation/PPL_NAV_001_basics.json) | 169 / 158 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_NAV_002](../../content/fr/ppl/navigation/PPL_NAV_002_charts.json) | 170 / 157 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_NAV_003](../../content/fr/ppl/navigation/PPL_NAV_003_dead_reckoning.json) | 174 / 142 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_NAV_004](../../content/fr/ppl/navigation/PPL_NAV_004_radio_nav.json) | 169 / 138 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_NAV_005](../../content/fr/ppl/navigation/PPL_NAV_005_gps.json) | 162 / 143 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_NAV_006](../../content/fr/ppl/navigation/PPL_NAV_006_flight_log.json) | 169 / 146 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_NAV_007](../../content/fr/ppl/navigation/PPL_NAV_007_visual_fixes.json) | 168 / 145 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_NAV_008](../../content/fr/ppl/navigation/PPL_NAV_008_flight_planning_nav.json) | 172 / 144 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_NAV_009](../../content/fr/ppl/navigation/PPL_NAV_009_application.json) | 262 / 206 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_OP_001](../../content/fr/ppl/operational_procedures/PPL_OP_001_preflight.json) | 155 / 141 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_OP_002](../../content/fr/ppl/operational_procedures/PPL_OP_002_taxi_takeoff.json) | 154 / 139 | brief_content_review_depth, short_explanations_review_reasoning, possible_context_dependency |
| [PPL_OP_003](../../content/fr/ppl/operational_procedures/PPL_OP_003_landing.json) | 151 / 132 | brief_content_review_depth, short_explanations_review_reasoning, possible_context_dependency |
| [PPL_OP_004](../../content/fr/ppl/operational_procedures/PPL_OP_004_communication.json) | 156 / 132 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_OP_005](../../content/fr/ppl/operational_procedures/PPL_OP_005_emergency_proc.json) | 169 / 142 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_OP_006](../../content/fr/ppl/operational_procedures/PPL_OP_006_human_factors.json) | 166 / 145 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_PERF_001](../../content/fr/ppl/performance/PPL_PERF_001_course.json) | 288 / 240 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_PERF_002](../../content/fr/ppl/performance/PPL_PERF_002_course.json) | 303 / 263 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_POF_001](../../content/fr/ppl/principles_of_flight/PPL_POF_001_course.json) | 163 / 151 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_POF_002](../../content/fr/ppl/principles_of_flight/PPL_POF_002_course.json) | 167 / 157 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_POF_003](../../content/fr/ppl/principles_of_flight/PPL_POF_003_application.json) | 297 / 254 | brief_content_review_depth, short_explanations_review_reasoning |
| [PPL_RAD_001](../../content/fr/ppl/radio_procedure/PPL_RAD_001_course.json) | 272 / 225 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_AIRLAW_001](../../content/fr/lapl/air_law/LAPL_AIRLAW_001_application.json) | 270 / 224 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_AIRLAW_002](../../content/fr/lapl/air_law/LAPL_AIRLAW_002_application.json) | 366 / 310 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_AIRLAW_003](../../content/fr/lapl/air_law/LAPL_AIRLAW_003_application.json) | 308 / 262 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_AIRLAW_004](../../content/fr/lapl/air_law/LAPL_AIRLAW_004_application.json) | 215 / 186 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_AIRLAW_005](../../content/fr/lapl/air_law/LAPL_AIRLAW_005_application.json) | 209 / 179 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_AIRLAW_006](../../content/fr/lapl/air_law/LAPL_AIRLAW_006_application.json) | 200 / 171 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_AIRLAW_007](../../content/fr/lapl/air_law/LAPL_AIRLAW_007_application.json) | 205 / 167 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_AIRLAW_008](../../content/fr/lapl/air_law/LAPL_AIRLAW_008_application.json) | 210 / 176 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_AIRLAW_009](../../content/fr/lapl/air_law/LAPL_AIRLAW_009_application.json) | 211 / 179 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_AIRLAW_010](../../content/fr/lapl/air_law/LAPL_AIRLAW_010_application.json) | 200 / 166 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_AIRLAW_011](../../content/fr/lapl/air_law/LAPL_AIRLAW_011_application.json) | 368 / 304 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_AIRLAW_012](../../content/fr/lapl/air_law/LAPL_AIRLAW_012_application.json) | 441 / 415 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_AIRLAW_013](../../content/fr/lapl/air_law/LAPL_AIRLAW_013_application.json) | 497 / 439 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_AC_001](../../content/fr/lapl/aircraft_general/LAPL_AC_001_application.json) | 227 / 197 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_AC_002](../../content/fr/lapl/aircraft_general/LAPL_AC_002_application.json) | 227 / 196 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_AC_003](../../content/fr/lapl/aircraft_general/LAPL_AC_003_application.json) | 210 / 183 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_AC_004](../../content/fr/lapl/aircraft_general/LAPL_AC_004_application.json) | 201 / 181 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_AC_005](../../content/fr/lapl/aircraft_general/LAPL_AC_005_application.json) | 206 / 169 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_AC_006](../../content/fr/lapl/aircraft_general/LAPL_AC_006_application.json) | 211 / 166 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_AC_007](../../content/fr/lapl/aircraft_general/LAPL_AC_007_application.json) | 218 / 189 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_AC_008](../../content/fr/lapl/aircraft_general/LAPL_AC_008_application.json) | 208 / 159 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_COM_001](../../content/fr/lapl/communications/LAPL_COM_001_application.json) | 199 / 166 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_COM_002](../../content/fr/lapl/communications/LAPL_COM_002_application.json) | 204 / 168 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_EMG_001](../../content/fr/lapl/emergency/LAPL_EMG_001_application.json) | 330 / 279 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_FP_001](../../content/fr/lapl/flight_planning/LAPL_FP_001_application.json) | 209 / 173 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_HP_001](../../content/fr/lapl/human_performance/LAPL_HP_001_application.json) | 213 / 175 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_HP_002](../../content/fr/lapl/human_performance/LAPL_HP_002_application.json) | 207 / 181 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_HP_003](../../content/fr/lapl/human_performance/LAPL_HP_003_application.json) | 377 / 311 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_INS_001](../../content/fr/lapl/instrumentation/LAPL_INS_001_application.json) | 339 / 270 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_INS_002](../../content/fr/lapl/instrumentation/LAPL_INS_002_application.json) | 310 / 285 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_MB_001](../../content/fr/lapl/mass_and_balance/LAPL_MB_001_application.json) | 213 / 183 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_MB_002](../../content/fr/lapl/mass_and_balance/LAPL_MB_002_application.json) | 931 / 848 | — |
| [LAPL_MET_001](../../content/fr/lapl/meteo/LAPL_MET_001_application.json) | 215 / 186 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_MET_002](../../content/fr/lapl/meteo/LAPL_MET_002_application.json) | 216 / 183 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_MET_003](../../content/fr/lapl/meteo/LAPL_MET_003_application.json) | 212 / 178 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_MET_004](../../content/fr/lapl/meteo/LAPL_MET_004_application.json) | 217 / 180 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_MET_005](../../content/fr/lapl/meteo/LAPL_MET_005_application.json) | 221 / 174 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_MET_006](../../content/fr/lapl/meteo/LAPL_MET_006_application.json) | 201 / 160 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_MET_007](../../content/fr/lapl/meteo/LAPL_MET_007_application.json) | 202 / 165 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_MET_008](../../content/fr/lapl/meteo/LAPL_MET_008_application.json) | 207 / 183 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_MET_009](../../content/fr/lapl/meteo/LAPL_MET_009_application.json) | 215 / 191 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_MET_010](../../content/fr/lapl/meteo/LAPL_MET_010_application.json) | 209 / 171 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_NAV_001](../../content/fr/lapl/navigation/LAPL_NAV_001_application.json) | 310 / 248 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_NAV_002](../../content/fr/lapl/navigation/LAPL_NAV_002_application.json) | 210 / 190 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_NAV_003](../../content/fr/lapl/navigation/LAPL_NAV_003_application.json) | 211 / 189 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_NAV_004](../../content/fr/lapl/navigation/LAPL_NAV_004_application.json) | 215 / 174 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_NAV_005](../../content/fr/lapl/navigation/LAPL_NAV_005_application.json) | 210 / 170 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_NAV_006](../../content/fr/lapl/navigation/LAPL_NAV_006_application.json) | 203 / 175 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_NAV_007](../../content/fr/lapl/navigation/LAPL_NAV_007_application.json) | 210 / 178 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_NAV_008](../../content/fr/lapl/navigation/LAPL_NAV_008_application.json) | 209 / 177 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_NAV_009](../../content/fr/lapl/navigation/LAPL_NAV_009_application.json) | 213 / 176 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_OP_001](../../content/fr/lapl/operational_procedures/LAPL_OP_001_application.json) | 196 / 173 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_OP_002](../../content/fr/lapl/operational_procedures/LAPL_OP_002_application.json) | 195 / 171 | brief_content_review_depth, short_explanations_review_reasoning, possible_context_dependency |
| [LAPL_OP_003](../../content/fr/lapl/operational_procedures/LAPL_OP_003_application.json) | 192 / 164 | brief_content_review_depth, short_explanations_review_reasoning, possible_context_dependency |
| [LAPL_OP_004](../../content/fr/lapl/operational_procedures/LAPL_OP_004_application.json) | 197 / 164 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_OP_005](../../content/fr/lapl/operational_procedures/LAPL_OP_005_application.json) | 210 / 174 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_OP_006](../../content/fr/lapl/operational_procedures/LAPL_OP_006_application.json) | 207 / 177 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_PERF_001](../../content/fr/lapl/performance/LAPL_PERF_001_application.json) | 329 / 272 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_PERF_002](../../content/fr/lapl/performance/LAPL_PERF_002_application.json) | 344 / 295 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_POF_001](../../content/fr/lapl/principles_of_flight/LAPL_POF_001_application.json) | 204 / 183 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_POF_002](../../content/fr/lapl/principles_of_flight/LAPL_POF_002_application.json) | 208 / 189 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_POF_003](../../content/fr/lapl/principles_of_flight/LAPL_POF_003_application.json) | 338 / 286 | brief_content_review_depth, short_explanations_review_reasoning |
| [LAPL_RAD_001](../../content/fr/lapl/radio_procedure/LAPL_RAD_001_application.json) | 313 / 257 | brief_content_review_depth, short_explanations_review_reasoning |
| [CPL_090_001](../../content/fr/cpl/communications/CPL_090_001_application.json) | 323 / 287 | objective_needs_specific_outcomes, generic_prerequisites, brief_content_review_depth, short_explanations_review_reasoning |
| [CPL_033_001](../../content/fr/cpl/flight_planning/CPL_033_001_application.json) | 361 / 308 | objective_needs_specific_outcomes, generic_prerequisites, brief_content_review_depth, short_explanations_review_reasoning, possible_context_dependency |
| [CPL_040_001](../../content/fr/cpl/human_performance/CPL_040_001_application.json) | 358 / 300 | objective_needs_specific_outcomes, generic_prerequisites, brief_content_review_depth, short_explanations_review_reasoning |
| [CPL_022_001](../../content/fr/cpl/instrumentation/CPL_022_001_application.json) | 1036 / 953 | — |
| [CPL_031_001](../../content/fr/cpl/mass_and_balance/CPL_031_001_application.json) | 1126 / 981 | — |
| [CPL_032_001](../../content/fr/cpl/performance/CPL_032_001_application.json) | 910 / 843 | — |
| [CPL_081_001](../../content/fr/cpl/principles_of_flight/CPL_081_001_application.json) | 341 / 288 | objective_needs_specific_outcomes, generic_prerequisites, brief_content_review_depth, short_explanations_review_reasoning |
| [ATPL_021_001](../../content/fr/atpl/aircraft_general/ATPL_021_001_application.json) | 349 / 284 | objective_needs_specific_outcomes, generic_prerequisites, brief_content_review_depth, short_explanations_review_reasoning |
| [ATPL_021_002](../../content/fr/atpl/aircraft_general/ATPL_021_002_application.json) | 356 / 289 | objective_needs_specific_outcomes, generic_prerequisites, brief_content_review_depth, short_explanations_review_reasoning |
| [ATPL_090_001](../../content/fr/atpl/communications/ATPL_090_001_application.json) | 323 / 287 | objective_needs_specific_outcomes, generic_prerequisites, brief_content_review_depth, short_explanations_review_reasoning |
| [ATPL_033_001](../../content/fr/atpl/flight_planning/ATPL_033_001_application.json) | 391 / 350 | objective_needs_specific_outcomes, generic_prerequisites, brief_content_review_depth, short_explanations_review_reasoning |
| [ATPL_022_001](../../content/fr/atpl/instrumentation/ATPL_022_001_application.json) | 331 / 269 | objective_needs_specific_outcomes, generic_prerequisites, brief_content_review_depth, short_explanations_review_reasoning |
| [ATPL_031_001](../../content/fr/atpl/mass_and_balance/ATPL_031_001_application.json) | 1295 / 1159 | — |
| [ATPL_061_001](../../content/fr/atpl/navigation/ATPL_061_001_application.json) | 328 / 287 | objective_needs_specific_outcomes, generic_prerequisites, brief_content_review_depth, short_explanations_review_reasoning |
| [IR_010_001](../../content/fr/ir/air_law/IR_010_001_application.json) | 355 / 308 | objective_needs_specific_outcomes, generic_prerequisites, brief_content_review_depth, short_explanations_review_reasoning |
| [IR_090_001](../../content/fr/ir/communications/IR_090_001_application.json) | 323 / 287 | objective_needs_specific_outcomes, generic_prerequisites, brief_content_review_depth, short_explanations_review_reasoning |
| [IR_033_001](../../content/fr/ir/flight_planning/IR_033_001_application.json) | 399 / 347 | objective_needs_specific_outcomes, generic_prerequisites, brief_content_review_depth, short_explanations_review_reasoning |
| [IR_040_001](../../content/fr/ir/human_performance/IR_040_001_application.json) | 353 / 296 | objective_needs_specific_outcomes, generic_prerequisites, brief_content_review_depth, short_explanations_review_reasoning |
| [IR_022_001](../../content/fr/ir/instrumentation/IR_022_001_application.json) | 789 / 716 | — |
| [IR_050_001](../../content/fr/ir/meteo/IR_050_001_application.json) | 1027 / 872 | — |
| [IR_062_001](../../content/fr/ir/navigation/IR_062_001_application.json) | 402 / 332 | objective_needs_specific_outcomes, generic_prerequisites, brief_content_review_depth, short_explanations_review_reasoning |
| [IR_062_002](../../content/fr/ir/navigation/IR_062_002_application.json) | 368 / 305 | objective_needs_specific_outcomes, generic_prerequisites, brief_content_review_depth, short_explanations_review_reasoning |
