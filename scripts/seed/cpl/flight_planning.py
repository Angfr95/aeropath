"""Authored FR/EN teaching content: worked method and independent applications."""
from shared.authoring import save

def build():
    data = {'metadata': {'module_id': 'CPL_033_001',
                  'program': 'CPL',
                  'subject': 'flight_planning',
                  'chapter_title': 'Bilan carburant et suivi du vol',
                  'chapter_title_en': 'Fuel accounting and flight monitoring',
                  'reference_checked_on': '2026-09-29',
                  'learning_objective': 'Construire un bilan par phase, suivre les écarts sans double comptage et '
                                        'comparer des options d’arrivée avec protection et aptitude des terrains.',
                  'learning_objective_en': 'Build a phase-based fuel balance, monitor deviations without double '
                                           'counting and compare arrivals against protection and aerodrome '
                                           'suitability.',
                  'estimated_duration_minutes': 50,
                  'difficulty': 'advanced',
                  'prerequisites': ['Temps décimal et conversions minutes/heures',
                                    'Vitesse air et vitesse sol',
                                    'Volume, masse et densité'],
                  'exam_subject': '033',
                  'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 033',
                  'official_subject_format': {'questions': 33,
                                              'minutes': 90,
                                              'pass_percent': 75,
                                              'source': 'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67'},
                  'sources': [{'title': 'EASA — programme Part-FCL',
                               'url': 'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8'},
                              {'title': 'EASA — format de la matière, AMC1 ARA.FCL.300(b)',
                               'url': 'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67'},
                              {'title': 'FAA PHAK — navigation, calculs de temps et carburant',
                               'url': 'https://www.faa.gov/regulationspolicies/handbooksmanuals/aviation/phak/chapter-16-navigation'}],
                  'content_status': 'draft',
                  'coverage_status': 'partial',
                  'scope': 'CPL(A), France / EASA',
                  'quiz_format': 'chapter_practice',
                  'tags': ['CPL', 'mcq-four-options'],
                  'illustration': {'src': '/illustrations/cpl-033-001.svg', 'alt': 'Bilan carburant et suivi du vol'}},
     'concepts': [{'concept_id': 'CPL_033_001_C01',
                   'title': 'Relier distance, temps et consommation',
                   'title_en': 'Relate distance, time and consumption',
                   'content': 'Utiliser la vitesse sol pour le temps de parcours et le débit approprié à la phase. '
                              'Temps en heures = distance NM / vitesse sol kt. Carburant = temps × débit. Pour 180 NM '
                              'à 120 kt et 40 L/h constants, temps = 1,5 h et consommation = 60 L. 1 h 30 min vaut '
                              '1,5 h, pas 1,30 h. Le modèle exclut montée et descente, dont les allocations doivent '
                              'être ajoutées si elles ne sont pas déjà incluses.\n'
                              '\n'
                              '### Construire le calcul par phase\n'
                              'Avant de multiplier, écrire distance, GS, durée et débit pour chaque phase. Le débit '
                              'dépend de la puissance, de la configuration et des conditions ; ne pas appliquer le '
                              'débit de croisière à toute une montée sans données appropriées. La vitesse sol décrit '
                              "le déplacement sur la carte, alors que la vitesse air décrit le mouvement dans l'air. "
                              'Une distance en NM divisée par des kt donne des heures ; multiplier par 60 pour les '
                              'minutes.\n'
                              '\n'
                              '**Exemple guidé distinct du quiz :** 84 NM à 112 kt donnent 0,75 h, soit 45 min. À 36 '
                              "L/h, consommation 27 L. Puis 30 NM à 90 kt donnent un tiers d'heure, soit 20 min ; à "
                              '30 L/h, 10 L. Pour ces deux segments : 65 min et 37 L. Additionner les durées et les '
                              'consommations ; la moyenne arithmétique des deux GS ne permet pas, en général, de '
                              'calculer le temps total.\n'
                              '\n'
                              '**Erreur expliquée :** 1 h 15 min vaut 1 + 15/60 = 1,25 h. Écrire 1,15 h retire six '
                              "minutes au calcul. Garder les fractions jusqu'au résultat final ; transformer "
                              "prématurément un tiers d'heure en 0,3 sous-estime déjà le carburant du second "
                              'segment.\n'
                              '\n'
                              "Un tableau peut fournir directement une quantité pour la montée ou l'approche : la "
                              'traiter comme une allocation de phase. Ne pas la multiplier à nouveau par une durée. '
                              'Définir précisément début et fin du segment pour ne pas compter deux fois une phase.\n'
                              '\n'
                              'Le schéma en tête présente les seules données du cas autonome de comparaison de '
                              'terrains. Les exemples guidés ci-dessous utilisent des données différentes.',
                   'content_en': 'Use groundspeed for travel time and fuel flow appropriate to the phase. Time in '
                                 'hours = distance NM / groundspeed kt. Fuel = time × flow. For 180 NM at 120 kt and '
                                 'constant 40 L/h, time = 1.5 h and burn = 60 L. 1 h 30 min is 1.5 h, not 1.30 h. '
                                 'This model excludes climb and descent, whose allocations must be added if not '
                                 'already included.\n'
                                 '\n'
                                 '### Build the calculation by phase\n'
                                 'Before multiplying, record distance, GS, duration and flow for each phase. Flow '
                                 'depends on power, configuration and conditions; do not apply cruise flow to an '
                                 'entire climb without suitable data. Groundspeed describes displacement on a chart; '
                                 'airspeed describes motion through air. NM divided by knots gives hours; multiply by '
                                 '60 for minutes.\n'
                                 '\n'
                                 '**Worked example distinct from the quiz:** 84 NM at 112 kt gives 0.75 h, or 45 '
                                 'minutes. At 36 L/h, burn is 27 L. Then 30 NM at 90 kt gives one third of an hour, '
                                 'or 20 minutes; at 30 L/h, 10 L. Totals: 65 minutes and 37 L. Add times and burns; '
                                 'the arithmetic average of both groundspeeds generally does not give total travel '
                                 'time.\n'
                                 '\n'
                                 '**Explained error:** 1 h 15 min is 1 + 15/60 = 1.25 h. Writing 1.15 h removes six '
                                 'minutes. Retain fractions until the final result; prematurely rounding one third to '
                                 '0.3 already underestimates second-segment fuel.\n'
                                 '\n'
                                 'A table may directly supply a climb or approach quantity: treat it as a phase '
                                 'allocation, not a flow to multiply by time again. Define segment boundaries '
                                 'precisely to avoid counting any phase twice.\n'
                                 '\n'
                                 'The opening diagram supplies only the inputs for the independent aerodrome '
                                 'comparison case. Worked examples below use different inputs.',
                   'difficulty': 'advanced',
                   'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — 033',
                   'tags': ['CPL']},
                  {'concept_id': 'CPL_033_001_C02',
                   'title': 'Construire une allocation explicite',
                   'title_en': 'Build an explicit allocation',
                   'content': 'Exercice fictif : roulage 5 L, trajet 60 L, complément de planification 6 L, '
                              'dégagement 12 L et quantité protégée 20 L. Total = 103 L. Ces allocations sont des '
                              'données pédagogiques ; elles ne définissent pas une politique réglementaire '
                              'universelle. Distinguer carburant total, utilisable et inutilisable. Avec 110 L '
                              'utilisables au départ, l’excédent sur ce plan est 7 L. Une réserve protégée ne devient '
                              'pas automatiquement du carburant de trajet disponible.\n'
                              '\n'
                              '### Séparer le contenu du réservoir et les engagements du plan\n'
                              "Écrire d'abord le carburant utilisable. Si une donnée représente déjà l'utilisable, ne "
                              "pas retrancher une deuxième fois l'inutilisable. Ensuite nommer chaque allocation et "
                              'préciser les inclusions : roulage, phases du trajet, autres postes imposés et quantité '
                              "à protéger. Les sommes de ce chapitre sont des hypothèses d'exercice ; la licence CPL "
                              'ne définit pas à elle seule la politique carburant de toute exploitation.\n'
                              '\n'
                              '**Bilan guidé :** 100 L au total dont 4 inutilisables donnent 96 L utilisables. '
                              'Hypothèses : roulage 4 L, trajet 37 L, autre allocation 8 L, protection 25 L. '
                              "Engagement total 74 L ; excédent au départ 22 L. Si l'on dispose au contraire de « 96 "
                              "L utilisables » dans l'énoncé, le même excédent vaut 22 L, pas 18 L.\n"
                              '\n'
                              "L'excédent n'est pas le carburant prévu à l'arrivée. Après le seul roulage et le "
                              'trajet, il reste 96 − 4 − 37 = 55 L. Sur ces 55 L, le plan engage encore 8 + 25 = 33 L '
                              '; le solde libre reste 22 L. Cette double lecture contrôle les inclusions. Une somme '
                              'correcte avec une phase comptée deux fois reste un mauvais bilan.\n'
                              '\n'
                              "**À faire sans la solution :** tracer une flèche du parking à l'arrivée ; inscrire ce "
                              "qui est consommé et ce qui doit rester. Refaire le bilan après changement d'une "
                              'allocation, sans supprimer une protection pour rendre le total favorable.',
                   'content_en': 'Fictional exercise: taxi 5 L, trip 60 L, planning supplement 6 L, alternate 12 L '
                                 'and protected quantity 20 L. Total = 103 L. These are training inputs, not a '
                                 'universal regulatory fuel policy. Distinguish total, usable and unusable fuel. With '
                                 '110 L usable at departure, surplus over this plan is 7 L. A protected reserve does '
                                 'not automatically become available trip fuel.\n'
                                 '\n'
                                 '### Separate tank contents from plan commitments\n'
                                 'First establish usable fuel. If an input is already usable fuel, do not subtract '
                                 'unusable fuel again. Name each allocation and its inclusions: taxi, trip phases, '
                                 'other required items and protected quantity. These chapter totals are exercise '
                                 "assumptions; a CPL licence alone does not define every operation's fuel policy.\n"
                                 '\n'
                                 '**Worked balance:** 100 L total including 4 unusable gives 96 L usable. '
                                 'Assumptions: taxi 4 L, trip 37 L, other allocation 8 L, protected 25 L. Total '
                                 'commitment 74 L; departure surplus 22 L. If the input instead states “96 L usable”, '
                                 'surplus remains 22 L, not 18 L.\n'
                                 '\n'
                                 'Surplus is not predicted arrival fuel. After taxi and trip only, 96 − 4 − 37 = 55 L '
                                 'remains. Of this, the plan still commits 8 + 25 = 33 L; free balance remains 22 L. '
                                 'This second reading checks inclusions. Correct addition with a phase counted twice '
                                 'still gives a wrong plan.\n'
                                 '\n'
                                 '**Work before seeing answers:** draw a timeline from parking to arrival; label '
                                 'consumption and what must remain. Rebuild the balance after one allocation changes '
                                 'without deleting protection to make the total favourable.',
                   'difficulty': 'advanced',
                   'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — 033',
                   'tags': ['CPL']},
                  {'concept_id': 'CPL_033_001_C03',
                   'title': 'Mesurer les écarts',
                   'title_en': 'Measure deviations',
                   'content': 'Comparer régulièrement carburant calculé, indications disponibles et temps écoulé, en '
                              'tenant compte de la précision des instruments. Une consommation prévue de 40 L et '
                              'constatée de 48 L donne +8 L et +20 %. Rechercher les causes et recalculer l’arrivée. '
                              'Avec 70 L utilisables restants, une protection de 20 L et un débit futur de 50 L/h, le '
                              'temps disponible avant cette protection est 1 h. Ne pas confondre cette limite de '
                              'planification avec un droit de poursuivre.\n'
                              '\n'
                              "### Transformer un écart en prévision d'arrivée\n"
                              'Comparer les valeurs au même point et à la même heure. Si le départ utilisable était '
                              '96 L et que le suivi indique 61 L restants, la consommation déduite est 35 L. Une '
                              'prévision de 30 L à ce point donne un excès de 5 L, soit 5/30 = 16,7 % du prévu. '
                              'Diviser par 35 répondrait à une autre question. Cette arithmétique ne démontre pas, '
                              'seule, une fuite : vérifier données de départ, indications, durée, régime et '
                              'conditions, sans négliger un indice anormal.\n'
                              '\n'
                              "**Recalcul guidé :** il reste 50 min à 42 L/h, plus 5 L d'approche non inclus. Besoin "
                              'futur 35 + 5 = 40 L ; arrivée prévue 61 − 40 = 21 L. Avec protection 25 L, déficit de '
                              "4 L. L'écart passé sert à revoir les hypothèses futures, mais ne doit pas être retiré "
                              'une deuxième fois du carburant restant constaté.\n'
                              '\n'
                              "Une moyenne historique n'est pas forcément le débit futur : un changement de puissance "
                              'ou de régime demande une donnée adaptée. Quand des indications divergent, ne pas '
                              'retenir automatiquement la plus favorable. Rechercher la cohérence avec les moyens '
                              'prévus, traiter les anomalies selon la procédure et conserver explicitement '
                              "l'incertitude dans les options.",
                   'content_en': 'Regularly compare calculated fuel, available indications and elapsed time, '
                                 'considering instrument accuracy. Expected burn 40 L versus observed 48 L gives +8 L '
                                 'and +20%. Investigate causes and recalculate arrival fuel. With 70 L usable '
                                 'remaining, 20 L protected and future flow 50 L/h, time before reaching that '
                                 'protection is 1 h. This planning limit is not permission to continue.\n'
                                 '\n'
                                 '### Turn a deviation into an arrival prediction\n'
                                 'Compare values at the same point and time. With 96 L usable at departure and 61 L '
                                 'now indicated remaining, inferred burn is 35 L. Planned burn of 30 L to this point '
                                 'gives 5 L excess, or 5/30 = 16.7% of plan. Dividing by 35 answers a different '
                                 'question. Arithmetic alone does not prove a leak: check departure inputs, '
                                 'indications, time, power and conditions without dismissing an abnormal cue.\n'
                                 '\n'
                                 '**Worked replan:** 50 minutes remain at 42 L/h plus 5 L approach not included. '
                                 'Future need is 35 + 5 = 40 L; predicted arrival is 61 − 40 = 21 L. Against 25 L '
                                 'protection, deficit is 4 L. Past deviation informs future assumptions but must not '
                                 'be subtracted again from observed remaining fuel.\n'
                                 '\n'
                                 'Historical average flow is not necessarily future flow: power or operating changes '
                                 'require appropriate data. If indications disagree, do not automatically select the '
                                 'favourable one. Cross-check through prescribed means, handle anomalies under the '
                                 'procedure and retain uncertainty explicitly when comparing options.',
                   'difficulty': 'advanced',
                   'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — 033',
                   'tags': ['CPL']},
                  {'concept_id': 'CPL_033_001_C04',
                   'title': 'Réviser une décision',
                   'title_en': 'Revise a decision',
                   'content': 'Un vent défavorable réduit la vitesse sol : 180 NM à 100 kt demandent 1,8 h et 72 L à '
                              '40 L/h, soit 12 L de plus qu’à 120 kt. Un trajet plus court, un terrain approprié ou '
                              'un avitaillement intermédiaire peuvent être nécessaires. Pour convertir un volume en '
                              'masse, utiliser la densité fournie : 100 L à 0,72 kg/L représentent 72 kg. La densité '
                              'dépend du carburant et de sa température ; ne pas supposer que 1 L vaut 1 kg.\n'
                              '\n'
                              '### Comparer des options avec leurs contraintes\n'
                              "Pour chaque terrain : faisabilité de la route et de l'atterrissage, carburant jusqu'à "
                              "l'arrivée, quantité restante et protection à maintenir. Une route courte ne suffit pas "
                              'si piste, météo ou performances la rendent inappropriée. Une autonomie calculée ne '
                              "garantit pas qu'un aérodrome adéquat se trouve à portée.\n"
                              '\n'
                              "**Cas guidé :** 61 L utilisables, protection 25 L. A demande 40 L jusqu'à l'arrivée, B "
                              '29 L ; les deux terrains et routes sont appropriés par hypothèse. A laisserait 21 L, B '
                              '32 L. B respecte le critère carburant donné avec 7 L de marge ; A présente un déficit. '
                              "Si l'aptitude de B était inconnue, il faudrait l'établir plutôt que conclure sur le "
                              'seul solde.\n'
                              '\n'
                              '**Conversion et décision :** masse = volume × densité. Pour ajouter 18 kg à une '
                              'densité fournie de 0,75 kg/L, il faut 24 L. Vérifier ensuite capacités du réservoir, '
                              'masse et centrage. Ni une conversion exacte ni un réservoir de capacité suffisante ne '
                              'prouvent toutes les autres limites respectées.\n'
                              '\n'
                              'Le suivi est continu : actualiser le point de décision avant que les options '
                              'disparaissent. Les quantités protégées de ces exercices ne définissent pas les seuils '
                              "d'annonce MINIMUM FUEL ou MAYDAY FUEL ; ces procédures et la politique réglementaire "
                              'applicable constituent un apprentissage distinct.\n'
                              '\n'
                              '**Format de travail :** 15 QCM de chapitre, quatre choix et une réponse correcte. La '
                              'matière 033 comporte 33 questions en 90 minutes dans le tableau EASA cité. Ce chapitre '
                              'ne couvre qu’une partie de la matière ; il ne constitue pas une épreuve complète. '
                              'Cadre EASA, même en anglais ou avec les unités US/britanniques.',
                   'content_en': 'Adverse wind reduces groundspeed: 180 NM at 100 kt requires 1.8 h and 72 L at 40 '
                                 'L/h, 12 L more than at 120 kt. A shorter route, suitable airfield or intermediate '
                                 'refuelling may be needed. Convert volume to mass using the supplied density: 100 L '
                                 'at 0.72 kg/L is 72 kg. Density depends on fuel and temperature; do not assume 1 L '
                                 'equals 1 kg.\n'
                                 '\n'
                                 '### Compare options with their constraints\n'
                                 'For each aerodrome consider route and landing feasibility, fuel to arrival, '
                                 'remaining quantity and required protection. A short route is insufficient if '
                                 'runway, weather or performance makes the airfield unsuitable. Calculated endurance '
                                 'does not ensure a suitable aerodrome is reachable.\n'
                                 '\n'
                                 '**Worked case:** 61 L usable, 25 L protected. A needs 40 L to arrival, B 29 L; both '
                                 'routes and aerodromes are suitable by assumption. A leaves 21 L, B 32 L. B meets '
                                 "the stated fuel criterion with 7 L margin; A has a deficit. If B's suitability were "
                                 'unknown, establish it rather than concluding from fuel alone.\n'
                                 '\n'
                                 '**Conversion and decision:** mass = volume × density. Adding 18 kg at a supplied '
                                 'density of 0.75 kg/L requires 24 L. Then check tank capacity, mass and CG. Neither '
                                 'accurate conversion nor sufficient tank capacity proves all other limits met.\n'
                                 '\n'
                                 'Monitoring is continuous: update the decision point before options disappear. '
                                 'Protected quantities in these exercises do not define MINIMUM FUEL or MAYDAY FUEL '
                                 'declaration thresholds; those procedures and applicable regulatory policy are a '
                                 'separate learning task.\n'
                                 '\n'
                                 '**Practice format:** 15 chapter MCQs, four choices and one correct answer. Subject '
                                 '033 has 33 questions in 90 minutes in the cited EASA table. This chapter covers '
                                 'only part of the subject and is not a full paper. EASA framework, including with '
                                 'English or US/British units.',
                   'difficulty': 'advanced',
                   'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — 033',
                   'tags': ['CPL']}],
     'questions': [{'concept_id': 'CPL_033_001_C01',
                    'concept_title': 'Relier distance, temps et consommation',
                    'level': 3,
                    'question': 'Segment A : 90 NM à 120 kt sol, puis B : 40 NM à 80 kt. Durée totale ?',
                    'question_en': 'Segment A: 90 NM at 120 kt GS, then B: 40 NM at 80 kt. Total time?',
                    'options': ['1 h 15 min', '1 h 18 min', '1 h 25 min', '1 h 30 min'],
                    'options_en': ['1 h 15 min', '1 h 18 min', '1 h 25 min', '1 h 30 min'],
                    'correct_answer': '1 h 15 min',
                    'explanation': 'A prend 45 min, B 30 min : 75 min. 78 min proviennent de 130 NM divisés par la '
                                   'moyenne simple 100 kt ; chaque vitesse doit être appliquée à son propre segment.',
                    'explanation_en': 'A takes 45 minutes, B 30: total 75. 78 minutes comes from dividing 130 NM by '
                                      'the simple 100 kt average; apply each speed to its own segment.',
                    'use_exercise_svg': False,
                    'tags': ['CPL', 'exam-mcq', 'exam-subject:033'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/regulationspolicies/handbooksmanuals/aviation/phak/chapter-16-navigation'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 033',
                    'explanation_illustration': '/illustrations/cpl-fuel-method.svg'},
                   {'concept_id': 'CPL_033_001_C01',
                    'concept_title': 'Relier distance, temps et consommation',
                    'level': 3,
                    'question': 'Segment A : 90 NM à 120 kt et 32 L/h ; B : 40 NM à 80 kt et 28 L/h. Carburant des '
                                'deux segments ?',
                    'question_en': 'Segment A: 90 NM at 120 kt and 32 L/h; B: 40 NM at 80 kt and 28 L/h. Combined '
                                   'fuel?',
                    'options': ['40 L', '35 L', '37,5 L', '38 L'],
                    'options_en': ['40 L', '35 L', '37.5 L', '38 L'],
                    'correct_answer': '38 L',
                    'explanation': 'A : 0,75 × 32 = 24 L ; B : 0,5 × 28 = 14 L ; total 38. 40 applique le premier '
                                   'débit partout, 35 le second ; 37,5 utilise leur moyenne simple sans pondérer les '
                                   'durées.',
                    'explanation_en': 'A: 0.75 × 32 = 24 L; B: 0.5 × 28 = 14 L; total 38. 40 applies the first flow '
                                      'throughout, 35 the second; 37.5 uses their unweighted average.',
                    'use_exercise_svg': False,
                    'tags': ['CPL', 'exam-mcq', 'exam-subject:033'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/regulationspolicies/handbooksmanuals/aviation/phak/chapter-16-navigation'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 033',
                    'explanation_illustration': '/illustrations/cpl-fuel-method.svg'},
                   {'concept_id': 'CPL_033_001_C01',
                    'concept_title': 'Relier distance, temps et consommation',
                    'level': 3,
                    'question': '1 h 24 min à 35 L/h. Quelle consommation ?',
                    'question_en': '1 h 24 min at 35 L/h. What burn?',
                    'options': ['14 L', '59 L', '49 L', '43,4 L'],
                    'options_en': ['14 L', '59 L', '49 L', '43.4 L'],
                    'correct_answer': '49 L',
                    'explanation': '24/60 = 0,4 h ; 1,4 × 35 = 49 L. 43,4 traite 1 h 24 comme 1,24 h ; 14 omet '
                                   "l'heure entière. Additionner 35 et 24 pour obtenir 59 mélange un débit et une "
                                   'durée.',
                    'explanation_en': '24/60 = 0.4 h; 1.4 × 35 = 49 L. 43.4 treats 1 h 24 as 1.24 h; 14 omits the '
                                      'full hour. Adding 35 and 24 to get 59 mixes a flow and a duration.',
                    'use_exercise_svg': False,
                    'tags': ['CPL', 'exam-mcq', 'exam-subject:033'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/regulationspolicies/handbooksmanuals/aviation/phak/chapter-16-navigation'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 033',
                    'explanation_illustration': '/illustrations/cpl-fuel-method.svg'},
                   {'concept_id': 'CPL_033_001_C02',
                    'concept_title': 'Construire une allocation explicite',
                    'level': 3,
                    'question': 'Plan fictif : roulage 4 L, montée 8 L, croisière 38 L, descente/approche 6 L, autre '
                                "allocation 9 L, protection 24 L. Aucun poste n'est inclus dans un autre. Total ?",
                    'question_en': 'Fictional plan: taxi 4 L, climb 8 L, cruise 38 L, descent/approach 6 L, other '
                                   'allocation 9 L, protected 24 L. No overlap. Total?',
                    'options': ['97 L', '89 L', '85 L', '65 L'],
                    'options_en': ['97 L', '89 L', '85 L', '65 L'],
                    'correct_answer': '89 L',
                    'explanation': '4+8+38+6+9+24 = 89 L. 85 omet le roulage, 65 la protection ; 97 compte deux fois '
                                   'la montée. Les inclusions sont explicitement disjointes.',
                    'explanation_en': '4+8+38+6+9+24 = 89 L. 85 omits taxi, 65 protection; 97 counts climb twice. '
                                      'Inputs explicitly state no overlap.',
                    'use_exercise_svg': False,
                    'tags': ['CPL', 'exam-mcq', 'exam-subject:033'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/regulationspolicies/handbooksmanuals/aviation/phak/chapter-16-navigation'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 033',
                    'explanation_illustration': '/illustrations/cpl-fuel-method.svg'},
                   {'concept_id': 'CPL_033_001_C02',
                    'concept_title': 'Construire une allocation explicite',
                    'level': 3,
                    'question': 'Réservoirs : 104 L au total dont 4 inutilisables. Plan exigeant 89 L utilisables. '
                                'Excédent au départ ?',
                    'question_en': 'Tanks contain 104 L total including 4 unusable. Plan requires 89 L usable. '
                                   'Departure surplus?',
                    'options': ['11 L', '15 L', '7 L', '24 L'],
                    'options_en': ['11 L', '15 L', '7 L', '24 L'],
                    'correct_answer': '11 L',
                    'explanation': "Utilisable = 104−4 = 100 ; excédent 100−89 = 11. 15 compte l'inutilisable, 7 le "
                                   "retire deux fois ; 24 est une allocation du plan, pas l'excédent.",
                    'explanation_en': 'Usable = 104−4 = 100; surplus 100−89 = 11. 15 counts unusable fuel, 7 '
                                      'subtracts it twice; 24 is a plan allocation, not surplus.',
                    'use_exercise_svg': False,
                    'tags': ['CPL', 'exam-mcq', 'exam-subject:033'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/regulationspolicies/handbooksmanuals/aviation/phak/chapter-16-navigation'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 033',
                    'explanation_illustration': '/illustrations/cpl-fuel-method.svg'},
                   {'concept_id': 'CPL_033_001_C02',
                    'concept_title': 'Construire une allocation explicite',
                    'level': 3,
                    'question': 'La fiche indique déjà 100 L utilisables ; capacité inutilisable 4 L. Pour un besoin '
                                'de 89 L utilisables, quelle marge ?',
                    'question_en': 'The sheet already states 100 L usable; unusable capacity is 4 L. For an 89 L '
                                   'usable requirement, what margin?',
                    'options': ['7 L', '15 L', '4 L', '11 L'],
                    'options_en': ['7 L', '15 L', '4 L', '11 L'],
                    'correct_answer': '11 L',
                    'explanation': "La donnée utilisable exclut déjà l'inutilisable : 100−89 = 11. Retrancher encore "
                                   '4 change la définition de la donnée ; ajouter 4 rendrait disponible un carburant '
                                   "qui ne l'est pas.",
                    'explanation_en': 'The usable input already excludes unusable fuel: 100−89 = 11. Subtracting 4 '
                                      'again changes the input definition; adding it would treat unavailable fuel as '
                                      'usable.',
                    'use_exercise_svg': False,
                    'tags': ['CPL', 'exam-mcq', 'exam-subject:033'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/regulationspolicies/handbooksmanuals/aviation/phak/chapter-16-navigation'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 033',
                    'explanation_illustration': '/illustrations/cpl-fuel-method.svg'},
                   {'concept_id': 'CPL_033_001_C03',
                    'concept_title': 'Mesurer les écarts',
                    'level': 3,
                    'question': 'Au même point, prévu consommé 30 L ; départ utilisable 100 L, restant constaté 64 L. '
                                'Surconsommation absolue ?',
                    'question_en': 'At the same point, planned burn 30 L; starting usable fuel 100 L, observed '
                                   'remaining 64 L. Absolute excess burn?',
                    'options': ['34 L', '20 L', '6 L', '36 L'],
                    'options_en': ['34 L', '20 L', '6 L', '36 L'],
                    'correct_answer': '6 L',
                    'explanation': 'Consommation déduite 100−64 = 36 L ; excès 36−30 = 6. 36 est le consommé total ; '
                                   '34 compare à tort restant et consommé prévu. Vérifier aussi la fiabilité des '
                                   'données mesurées.',
                    'explanation_en': 'Inferred burn 100−64 = 36 L; excess 36−30 = 6. 36 is total burn; 34 wrongly '
                                      'compares remaining fuel with planned burn. Also check reliability of measured '
                                      'inputs.',
                    'use_exercise_svg': False,
                    'tags': ['CPL', 'exam-mcq', 'exam-subject:033'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/regulationspolicies/handbooksmanuals/aviation/phak/chapter-16-navigation'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 033',
                    'explanation_illustration': '/illustrations/cpl-fuel-method.svg'},
                   {'concept_id': 'CPL_033_001_C03',
                    'concept_title': 'Mesurer les écarts',
                    'level': 3,
                    'question': 'Prévu 30 L, consommé déduit 36 L. Écart relatif au prévu ?',
                    'question_en': 'Planned 30 L, inferred burn 36 L. Deviation relative to plan?',
                    'options': ['120 %', '20 %', '16,7 %', '6 %'],
                    'options_en': ['120%', '20%', '16.7%', '6%'],
                    'correct_answer': '20 %',
                    'explanation': "L'écart 6 est divisé par la base demandée 30 : 20 %. 16,7 % utilise 36 ; 120 % "
                                   'est le ratio total 36/30, non le seul excès ; 6 reste une quantité en litres.',
                    'explanation_en': 'Divide the 6 L difference by the requested 30 L reference: 20%. 16.7% uses 36; '
                                      '120% is total ratio 36/30, not excess alone; 6 is a quantity in litres.',
                    'use_exercise_svg': False,
                    'tags': ['CPL', 'exam-mcq', 'exam-subject:033'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/regulationspolicies/handbooksmanuals/aviation/phak/chapter-16-navigation'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 033',
                    'explanation_illustration': '/illustrations/cpl-fuel-method.svg'},
                   {'concept_id': 'CPL_033_001_C03',
                    'concept_title': 'Mesurer les écarts',
                    'level': 3,
                    'question': 'Restant utilisable 64 L ; 45 min à 40 L/h puis approche 5 L non incluse. Protection '
                                "à l'arrivée 24 L. Marge prévue au-dessus de cette protection ?",
                    'question_en': 'Usable remaining 64 L; 45 minutes at 40 L/h then 5 L approach not included. '
                                   'Arrival protection 24 L. Predicted margin above protection?',
                    'options': ['5 L', '11 L', '10 L', '29 L'],
                    'options_en': ['5 L', '11 L', '10 L', '29 L'],
                    'correct_answer': '5 L',
                    'explanation': 'Trajet futur 30 L, approche 5 : arrivée 29 L ; marge 29−24 = 5. 10 omet '
                                   "l'approche ; 29 est l'arrivée et non la marge. Un écart passé déjà reflété dans "
                                   '64 L ne se soustrait pas à nouveau.',
                    'explanation_en': 'Future trip 30 L, approach 5: arrival 29 L; margin 29−24 = 5. 10 omits '
                                      'approach; 29 is arrival fuel, not margin. Past excess already reflected in 64 '
                                      'L is not subtracted again.',
                    'use_exercise_svg': False,
                    'tags': ['CPL', 'exam-mcq', 'exam-subject:033'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/regulationspolicies/handbooksmanuals/aviation/phak/chapter-16-navigation'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 033',
                    'explanation_illustration': '/illustrations/cpl-fuel-method.svg'},
                   {'concept_id': 'CPL_033_001_C04',
                    'concept_title': 'Réviser une décision',
                    'level': 3,
                    'question': '180 NM à 100 kt sol, débit futur donné 40 L/h ; aucune autre phase incluse. '
                                'Consommation de ce segment ?',
                    'question_en': '180 NM at 100 kt GS, given future flow 40 L/h; no other phase included. Segment '
                                   'burn?',
                    'options': ['60 L', '64 L', '80 L', '72 L'],
                    'options_en': ['60 L', '64 L', '80 L', '72 L'],
                    'correct_answer': '72 L',
                    'explanation': "180/100 = 1,8 h ; ×40 = 72 L. 60 conserve l'ancien calcul à 120 kt ; 80 arrondit "
                                   "prématurément la durée à deux heures. Ce résultat n'inclut pas les autres "
                                   'allocations.',
                    'explanation_en': '180/100 = 1.8 h; ×40 = 72 L. 60 retains the previous 120 kt calculation; 80 '
                                      'rounds time prematurely to two hours. This excludes other allocations.',
                    'use_exercise_svg': False,
                    'tags': ['CPL', 'exam-mcq', 'exam-subject:033'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/regulationspolicies/handbooksmanuals/aviation/phak/chapter-16-navigation'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 033',
                    'explanation_illustration': '/illustrations/cpl-fuel-method.svg'},
                   {'concept_id': 'CPL_033_001_C04',
                    'concept_title': 'Réviser une décision',
                    'level': 3,
                    'question': 'Même débit : ancien trajet 180 NM à 120 kt ; nouveau 180 NM à 100 kt, 40 L/h. Avec '
                                'excédent initial de 7 L, quel solde après cette seule modification ?',
                    'question_en': 'Same flow: previous 180 NM at 120 kt; now 180 NM at 100 kt, 40 L/h. With initial '
                                   'surplus 7 L, what balance after this change alone?',
                    'options': ['Déficit de 12 L', 'Excédent de 19 L', 'Déficit de 5 L', 'Excédent de 5 L'],
                    'options_en': ['12 L deficit', '19 L surplus', '5 L deficit', '5 L surplus'],
                    'correct_answer': 'Déficit de 5 L',
                    'explanation': 'Le supplément vaut 72−60 = 12 L. Nouveau solde 7−12 = −5 L. Ne pas confondre '
                                   'supplément et déficit final, ni ajouter une consommation à une marge.',
                    'explanation_en': 'Extra burn is 72−60 = 12 L. New balance 7−12 = −5 L. Do not confuse additional '
                                      'consumption with final deficit or add burn to a margin.',
                    'use_exercise_svg': False,
                    'tags': ['CPL', 'exam-mcq', 'exam-subject:033'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/regulationspolicies/handbooksmanuals/aviation/phak/chapter-16-navigation'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 033',
                    'explanation_illustration': '/illustrations/cpl-fuel-method.svg'},
                   {'concept_id': 'CPL_033_001_C04',
                    'concept_title': 'Réviser une décision',
                    'level': 3,
                    'question': 'Il faut ajouter 21,6 kg de carburant, densité donnée 0,72 kg/L. Quel volume ajouter '
                                '?',
                    'question_en': 'Add 21.6 kg fuel at supplied density 0.72 kg/L. What volume is needed?',
                    'options': ['72 L', '30 L', '15,552 L', '21,6 L'],
                    'options_en': ['72 L', '30 L', '15.552 L', '21.6 L'],
                    'correct_answer': '30 L',
                    'explanation': 'Volume = masse/densité = 21,6/0,72 = 30 L. Multiplier donne 15,552 avec une '
                                   'opération inadaptée ; kg et L ne sont pas interchangeables. Contrôler ensuite '
                                   'capacité et chargement.',
                    'explanation_en': 'Volume = mass/density = 21.6/0.72 = 30 L. Multiplying gives 15.552 using the '
                                      'wrong operation; kg and L are not interchangeable. Then check capacity and '
                                      'loading.',
                    'use_exercise_svg': False,
                    'tags': ['CPL', 'exam-mcq', 'exam-subject:033'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/regulationspolicies/handbooksmanuals/aviation/phak/chapter-16-navigation'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 033',
                    'explanation_illustration': '/illustrations/cpl-fuel-method.svg'},
                   {'concept_id': 'CPL_033_001_C03',
                    'concept_title': 'Mesurer les écarts',
                    'level': 3,
                    'question': 'Restant 58 L, protection 24 L. A demande 40 min à 42 L/h + 4 L approche ; B 30 min à '
                                '42 L/h + 4 L mais piste inappropriée. C demande 50 min à 42 L/h + 4 L. A et C sont '
                                'appropriés. Quelle option satisfait les contraintes données ?',
                    'question_en': 'Remaining 58 L, protected 24 L. A needs 40 min at 42 L/h + 4 L approach; B 30 min '
                                   'at 42 L/h + 4 L but runway unsuitable. C needs 50 min at 42 L/h + 4 L. A and C '
                                   'are suitable. Which meets the stated constraints?',
                    'options': ['A, arrivée 26 L', 'B, arrivée 33 L', 'C, arrivée 19 L', 'A, arrivée 30 L'],
                    'options_en': ['A, arrival 26 L', 'B, arrival 33 L', 'C, arrival 19 L', 'A, arrival 30 L'],
                    'correct_answer': 'A, arrivée 26 L',
                    'explanation': 'A consomme 28+4 = 32 L et laisse 26, au-dessus de 24. B laisse plus de carburant '
                                   "mais sa piste l'exclut ; C laisse 19, sous la protection. 30 pour A omet "
                                   "l'approche.",
                    'explanation_en': 'A burns 28+4 = 32 L and leaves 26, above 24. B leaves more fuel but its runway '
                                      "excludes it; C leaves 19, below protection. A's 30 L answer omits approach.",
                    'use_exercise_svg': True,
                    'tags': ['CPL', 'exam-mcq', 'exam-subject:033'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/regulationspolicies/handbooksmanuals/aviation/phak/chapter-16-navigation'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 033'},
                   {'concept_id': 'CPL_033_001_C02',
                    'concept_title': 'Construire une allocation explicite',
                    'level': 3,
                    'question': 'Deux documents de trajet incluent chacun la même montée de 8 L. Comment éviter une '
                                'addition erronée ?',
                    'question_en': 'Two trip documents each include the same 8 L climb. How should erroneous addition '
                                   'be avoided?',
                    'options': ['Additionner les deux par principe de prudence',
                                'Supprimer toute allocation de montée',
                                'Retrancher 8 L de la quantité protégée',
                                "Identifier les limites des phases et ne compter la montée qu'une fois"],
                    'options_en': ['Always add both as a precaution',
                                   'Remove all climb allocation',
                                   'Subtract 8 L from protected quantity',
                                   'Identify phase boundaries and count climb once'],
                    'correct_answer': "Identifier les limites des phases et ne compter la montée qu'une fois",
                    'explanation': "La même consommation ne se produit qu'une fois dans ce scénario. Réconcilier les "
                                   'inclusions garde un bilan traçable ; une marge supplémentaire décidée séparément '
                                   'ne doit pas être confondue avec ce double comptage.',
                    'explanation_en': 'The same burn occurs once in this scenario. Reconciling inclusions keeps the '
                                      'balance traceable; a separately chosen additional margin must not be confused '
                                      'with double counting.',
                    'use_exercise_svg': False,
                    'tags': ['CPL', 'exam-mcq', 'exam-subject:033'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/regulationspolicies/handbooksmanuals/aviation/phak/chapter-16-navigation'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 033',
                    'explanation_illustration': '/illustrations/cpl-fuel-method.svg'},
                   {'concept_id': 'CPL_033_001_C01',
                    'concept_title': 'Relier distance, temps et consommation',
                    'level': 3,
                    'question': 'Navigation : TAS 125 kt, vent de face axial 25 kt supposé constant, distance 150 NM. '
                                'Quelle durée de segment ?',
                    'question_en': 'Navigation: TAS 125 kt, axial headwind 25 kt assumed constant, distance 150 NM. '
                                   'Segment time?',
                    'options': ['1 h', '6 h', '1 h 30 min', '1 h 12 min'],
                    'options_en': ['1 h', '6 h', '1 h 30 min', '1 h 12 min'],
                    'correct_answer': '1 h 30 min',
                    'explanation': 'GS = 125−25 = 100 kt ; 150/100 = 1,5 h. 1 h 12 emploie TAS, 1 h ajoute le vent de '
                                   "face comme vent arrière ; 6 h utilise la vitesse du vent seule. L'hypothèse est "
                                   'un vent axial, pas un triangle général.',
                    'explanation_en': 'GS = 125−25 = 100 kt; 150/100 = 1.5 h. 1 h 12 uses TAS, 1 h adds headwind as '
                                      'tailwind; 6 h uses wind speed alone. This assumes axial wind, not a general '
                                      'wind triangle.',
                    'use_exercise_svg': False,
                    'tags': ['CPL', 'exam-mcq', 'exam-subject:033'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/regulationspolicies/handbooksmanuals/aviation/phak/chapter-16-navigation'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 033',
                    'explanation_illustration': '/illustrations/cpl-fuel-method.svg'}]}
    save(data)
