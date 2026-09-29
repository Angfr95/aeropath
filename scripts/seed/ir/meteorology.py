"""Bilingual teaching content, revised after individual pedagogical review."""
from shared.authoring import save


def build():
    data = {'metadata': {'module_id': 'IR_050_001',
                  'program': 'IR',
                  'subject': 'meteo',
                  'chapter_title': 'Givrage : reconnaître et décider',
                  'chapter_title_en': 'Icing: recognition and decisions',
                  'reference_checked_on': '2026-09-29',
                  'learning_objective': 'Expliquer les mécanismes et limites de protection, distinguer '
                                        'observation et conclusion, puis comparer des sorties sous contraintes '
                                        'd’altitude, de carburant et de performances.',
                  'learning_objective_en': 'Explain mechanisms and protection limits, distinguish observation '
                                           'from inference, and compare exits under altitude, fuel and '
                                           'performance constraints.',
                  'estimated_duration_minutes': 50,
                  'difficulty': 'advanced',
                  'prerequisites': ['Température, nuages et profil vertical',
                                    'Portance, traînée et décrochage',
                                    'Lecture des limitations et calcul distance/vitesse/temps'],
                  'exam_subject': '050',
                  'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 050',
                  'official_subject_format': {'questions': 63,
                                              'minutes': 90,
                                              'pass_percent': 75,
                                              'source': 'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67'},
                  'sources': [{'title': 'EASA — programme Part-FCL',
                               'url': 'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8'},
                              {'title': 'EASA — format de la matière, AMC1 ARA.FCL.300(b)',
                               'url': 'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67'},
                              {'title': 'FAA AC 91-74B — mécanismes et effets du givrage, référence technique',
                               'url': 'https://www.faa.gov/documentLibrary/media/Advisory_Circular/AC_91-74B.pdf'},
                              {'title': 'FAA AIM 7 — dépôts et observations, référence technique',
                               'url': 'https://www.faa.gov/air_traffic/publications/atpubs/aim_html/chap7_section_1.html'},
                              {'title': 'EASA Air Operations — conditions givrantes',
                               'url': 'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-air-operations?erules-id=ERULES-1963177438-13362'}],
                  'content_status': 'draft',
                  'coverage_status': 'partial',
                  'scope': 'IR(A), France / EASA',
                  'quiz_format': 'chapter_practice',
                  'tags': ['IR', 'mcq-four-options'],
                  'illustration': {'src': '/illustrations/ir-050-001.svg',
                                   'alt': 'Givrage : reconnaître et décider'}},
     'concepts': [{'concept_id': 'IR_050_001_C01',
                   'title': 'Identifier les conditions',
                   'title_en': 'Identify conditions',
                   'content': '### Objectifs et méthode\n'
                              '\n'
                              'Expliquer le mécanisme avant de choisir une action : quelles données '
                              'observez-vous, que permettent-elles de conclure et que manque-t-il ? À la fin du '
                              'chapitre, vous devez pouvoir relier les conditions au dépôt, distinguer '
                              'protection et domaine approuvé, puis comparer des issues. Les exemples sont '
                              'fictifs ; les actions propres à un avion viennent de son manuel de vol.\n'
                              '\n'
                              '### Relier eau, température et exposition\n'
                              '\n'
                              'Une goutte en surfusion reste liquide sous 0 °C et peut geler à l’impact. Pour '
                              'raisonner sur l’accumulation, examiner ensemble eau liquide disponible, '
                              'température, taille des gouttes et durée d’exposition. Deux zones à −5 °C ne '
                              'présentent donc pas nécessairement le même risque. Cette comparaison concerne le '
                              'givrage de cellule par impact ; elle ne décrit pas tous les mécanismes de '
                              'givrage.\n'
                              '\n'
                              '**Exemple guidé :** à −6 °C, l’avion traverse de l’air sec, puis une couche '
                              'contenant des gouttes liquides. La température reste identique mais l’exposition '
                              'change. Comparer uniquement les températures ferait manquer l’apparition d’un '
                              'ingrédient du dépôt.\n'
                              '\n'
                              '### Lire un profil vertical\n'
                              '\n'
                              'Une couche positive peut faire fondre des précipitations qui entrent ensuite dans '
                              'une couche froide en restant liquides. Une température positive au sol ne décrit '
                              'pas toute la trajectoire. L’isotherme 0 °C est un repère : une inversion peut '
                              'produire plusieurs franchissements. Croiser le profil, les nuages, les '
                              'précipitations et la validité des observations.\n'
                              '\n'
                              '**Auto-évaluation :** « −8 °C à 7 000 ft » suffit-il à quantifier le dépôt ? Non '
                              ': il manque notamment l’eau liquide, les gouttes, l’exposition et les '
                              'caractéristiques de l’avion. Une température favorable au mécanisme n’est pas une '
                              'mesure de l’accumulation.',
                   'content_en': '### Objectives and method\n'
                                 '\n'
                                 'Explain the mechanism before choosing an action: what do you observe, what can '
                                 'you conclude and what is missing? By the end, relate conditions to '
                                 'accumulation, distinguish protection from the approved envelope, and compare '
                                 'escape options. Examples are fictional; aircraft-specific actions come from '
                                 'its flight manual.\n'
                                 '\n'
                                 '### Link water, temperature and exposure\n'
                                 '\n'
                                 'A supercooled droplet remains liquid below 0 °C and may freeze on impact. '
                                 'Consider available liquid water, temperature, droplet size and exposure time '
                                 'together. Two areas at −5 °C need not present the same risk. This comparison '
                                 'concerns impact icing on the airframe, not every icing mechanism.\n'
                                 '\n'
                                 '**Worked example:** at −6 °C, an aircraft crosses dry air, then a layer '
                                 'containing liquid droplets. Temperature remains unchanged but exposure '
                                 'changes. Comparing temperature alone misses the appearance of an ingredient '
                                 'for accumulation.\n'
                                 '\n'
                                 '### Read a vertical profile\n'
                                 '\n'
                                 'A warm layer can melt precipitation, which then enters a cold layer while '
                                 'remaining liquid. Positive surface temperature does not describe the entire '
                                 'flight path. The freezing level is a reference: an inversion may produce '
                                 'several crossings. Cross-check the profile, clouds, precipitation and '
                                 'observation validity.\n'
                                 '\n'
                                 '**Self-check:** is “−8 °C at 7,000 ft” enough to quantify a deposit? No: '
                                 'liquid water, droplets, exposure and aircraft characteristics are among the '
                                 'missing inputs. A temperature favourable to the mechanism is not a measurement '
                                 'of accumulation.',
                   'difficulty': 'advanced',
                   'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — 050',
                   'tags': ['IR']},
                  {'concept_id': 'IR_050_001_C02',
                   'title': 'Distinguer les dépôts',
                   'title_en': 'Distinguish deposits',
                   'content': '### Décrire avant d’interpréter\n'
                              '\n'
                              'Le givre blanc est généralement opaque et rugueux : une congélation rapide '
                              'emprisonne de l’air. Un dépôt plus transparent peut se former lorsque l’eau '
                              's’étale avant de geler : c’est le verglas. Le dépôt mixte combine des '
                              'caractéristiques des deux. Ces descriptions expliquent un mécanisme ; elles ne '
                              'constituent pas un classement universel de dangerosité.\n'
                              '\n'
                              'Noter le lieu du dépôt, son évolution et les changements de comportement. Blanc '
                              'ne signifie pas bénin ; peu visible ne signifie pas absent. Une observation '
                              'partielle ne permet pas de déduire l’état de toutes les surfaces.\n'
                              '\n'
                              '### Grosses gouttes et zones protégées\n'
                              '\n'
                              'Les grosses gouttes surfondues peuvent atteindre des zones derrière les surfaces '
                              'protégées. Pluie ou bruine verglaçante demandent donc une attention particulière '
                              'au domaine approuvé. Distinguer système installé, système fonctionnel et '
                              'conditions comprises dans les limitations : aucune de ces informations ne '
                              'remplace les deux autres.\n'
                              '\n'
                              '**Cas guidé :** le dégivrage nettoie sa zone, mais une bande de glace apparaît '
                              'plus en arrière. L’observation est une accumulation hors de la zone protégée. '
                              'Elle ne prouve pas, à elle seule, une panne électrique : le système peut '
                              'fonctionner sans maîtriser cette exposition. Appliquer les procédures associées '
                              'et organiser la sortie selon le manuel, sans attendre une identification parfaite '
                              'des gouttes.\n'
                              '\n'
                              '**Vérification personnelle :** pouvez-vous expliquer pourquoi « le bord d’attaque '
                              'est propre » ne prouve pas « l’avion est protégé partout » ? Le premier constat '
                              'ne porte que sur une zone ; la conclusion étend abusivement son périmètre.',
                   'content_en': '### Describe before interpreting\n'
                                 '\n'
                                 'Rime is generally opaque and rough: rapid freezing traps air. A more '
                                 'transparent deposit can form when water spreads before freezing: clear ice. '
                                 'Mixed ice combines characteristics of both. These descriptions explain a '
                                 'mechanism; they are not a universal ranking of danger.\n'
                                 '\n'
                                 'Note deposit location, development and changes in aircraft behaviour. White '
                                 'does not mean harmless; hard to see does not mean absent. Partial observation '
                                 'cannot establish the condition of every surface.\n'
                                 '\n'
                                 '### Large droplets and protected areas\n'
                                 '\n'
                                 'Supercooled large droplets can reach areas aft of protected surfaces. Freezing '
                                 'rain or drizzle therefore requires particular attention to the approved '
                                 'envelope. Distinguish installed equipment, functioning equipment and '
                                 'conditions within limitations: none substitutes for the other two.\n'
                                 '\n'
                                 '**Worked case:** de-icing clears its own area, but a band of ice appears '
                                 'farther aft. The observation is accumulation outside the protected area. '
                                 'Alone, it does not prove electrical failure: the system may function without '
                                 'controlling this exposure. Apply the relevant procedures and arrange an exit '
                                 'according to the manual without waiting for perfect droplet identification.\n'
                                 '\n'
                                 '**Self-check:** why does “the leading edge is clear” not prove “the aircraft '
                                 'is protected everywhere”? The first observation covers one area; the '
                                 'conclusion unjustifiably extends its scope.',
                   'difficulty': 'advanced',
                   'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — 050',
                   'tags': ['IR']},
                  {'concept_id': 'IR_050_001_C03',
                   'title': 'Comprendre les effets',
                   'title_en': 'Understand effects',
                   'content': '### Une dégradation de forme, pas seulement de masse\n'
                              '\n'
                              'Rugosité et déformation du profil perturbent l’écoulement. La traînée peut '
                              'augmenter et la portance maximale diminuer ; les caractéristiques de décrochage '
                              'peuvent changer. La masse de glace ne décrit donc pas à elle seule le problème. '
                              'Une faible épaisseur visible ne permet pas de calculer une marge de sécurité.\n'
                              '\n'
                              '**Cas guidé :** à altitude, configuration et puissance inchangées, la vitesse '
                              'diminue pendant une accumulation observée. Une augmentation de traînée est une '
                              'explication cohérente à vérifier. Cette tendance ne donne pas une nouvelle '
                              'vitesse de décrochage : il manque les données de l’avion. Recouper les '
                              'indications et appliquer les procédures, sans inventer une majoration universelle '
                              'en nœuds.\n'
                              '\n'
                              '### Un paramètre stable peut masquer la tendance\n'
                              '\n'
                              'Le pilote automatique peut maintenir l’altitude alors que les marges se '
                              'dégradent. L’altitude stable ne mesure ni l’état de surface ni la marge de '
                              'vitesse. Surveiller les paramètres et respecter les limitations d’utilisation du '
                              'pilote automatique propres à l’avion ; ce chapitre ne prescrit pas de procédure '
                              'universelle de déconnexion ou de récupération.\n'
                              '\n'
                              'Sondes, hélices, entrées d’air et empennages peuvent aussi être affectés. Un '
                              'pitot chauffé ne démontre pas la protection de l’aile. Une IAS cohérente peut '
                              'coexister avec une aérodynamique dégradée. Ne pas transposer indistinctement une '
                              'procédure de décrochage d’aile à une anomalie d’empennage : les actions viennent '
                              'de la procédure adaptée à l’avion et aux symptômes.',
                   'content_en': '### A change in shape, not just mass\n'
                                 '\n'
                                 'Roughness and altered aerofoil shape disturb airflow. Drag may increase and '
                                 'maximum lift decrease; stall characteristics may change. Ice mass alone '
                                 'therefore does not describe the problem. A thin visible deposit cannot '
                                 'quantify a safety margin.\n'
                                 '\n'
                                 '**Worked case:** at unchanged altitude, configuration and power, speed '
                                 'decreases during observed accumulation. Increased drag is a consistent '
                                 'explanation to check. The trend does not supply a new stall speed: aircraft '
                                 'data are missing. Cross-check indications and apply procedures rather than '
                                 'inventing a universal increment in knots.\n'
                                 '\n'
                                 '### One stable parameter can hide the trend\n'
                                 '\n'
                                 'Autopilot may maintain altitude while margins deteriorate. Stable altitude '
                                 'measures neither surface condition nor airspeed margin. Monitor parameters and '
                                 'respect aircraft-specific autopilot limitations; this chapter prescribes no '
                                 'universal disconnection or recovery procedure.\n'
                                 '\n'
                                 'Probes, propellers, intakes and tail surfaces may also be affected. A heated '
                                 'pitot does not establish wing protection. Consistent IAS may coexist with '
                                 'degraded aerodynamics. Do not indiscriminately transfer a wing-stall procedure '
                                 'to a tailplane anomaly: use procedures appropriate to the aircraft and '
                                 'symptoms.',
                   'difficulty': 'advanced',
                   'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — 050',
                   'tags': ['IR']},
                  {'concept_id': 'IR_050_001_C04',
                   'title': 'Préparer une sortie',
                   'title_en': 'Plan an exit',
                   'content': '### Comparer des options sous contraintes\n'
                              '\n'
                              'L’antigivrage vise à prévenir l’accumulation ; le dégivrage à retirer le dépôt. '
                              'Leur emploi dépend du système. Attendre arbitrairement une épaisseur ou '
                              'transposer le cycle d’un autre avion n’est pas une méthode générale.\n'
                              '\n'
                              'Commencer par éliminer les issues incompatibles avec relief, performances, '
                              'limitations, météo ou carburant. Comparer ensuite le temps d’exposition des '
                              'options restantes. Une clairance disponible ne démontre pas la faisabilité ; une '
                              'trajectoire calculée ne constitue pas une clairance.\n'
                              '\n'
                              '**Exemple guidé, profils fixes :** A conduit en air clair à 6 NM, à 6 500 ft, '
                              'alors que le minimum applicable est 7 000 ft. B conduit en air clair à 18 NM, à 8 '
                              '000 ft, avec le même minimum. GS 120 kt pour les deux ; toutes les autres '
                              'contraintes sont vérifiées pour B. A demanderait 60 × 6/120 = 3 minutes mais '
                              'viole le minimum. B demande 9 minutes et satisfait les données. Choisir d’abord '
                              'le temps le plus court ferait retenir une option incompatible. Ces nombres ne '
                              'sont pas une procédure opérationnelle.\n'
                              '\n'
                              '### Réévaluer avec les observations\n'
                              '\n'
                              'Si le dépôt contredit la prévision, actualiser la décision avec position, niveau, '
                              'heure, évolution et effets observés. Un compte rendu ancien, à un autre niveau ou '
                              'd’un autre avion ne garantit pas vos conditions. L’absence de compte rendu n’est '
                              'pas une observation d’absence de glace.\n'
                              '\n'
                              '**À restituer sans le cours :** séparer observation et interprétation ; expliquer '
                              'pourquoi un système fonctionnel peut ne pas suffire ; éliminer une issue '
                              'incompatible avant de comparer les temps ; préciser les données encore inconnues. '
                              'Une réponse justifiée nomme les informations utilisées et les limites de sa '
                              'conclusion.\n'
                              '\n'
                              '### Exercice de transfert : ne pas laisser le temps décider seul\n'
                              'Dans un dossier nouveau, construire une ligne par option : altitude prévue et '
                              "minimum, domaine de l'avion, météo jusqu'à la sortie, distance, GS, carburant "
                              'nécessaire et disponible. Éliminer une option incompatible avant de classer les '
                              'durées. Si une donnée manque, écrire « inconnue », pas « favorable ».\n'
                              '\n'
                              '**Contre-exemple guidé :** C demande 7 minutes, mais ses performances de montée '
                              'dans les conditions prévues ne sont pas établies ; D demande 11 minutes et ses '
                              'contraintes sont vérifiées. Les seules durées ne démontrent pas que C est '
                              'réalisable. La donnée à rechercher concerne précisément sa montée. Si les deux '
                              'trajets sont admissibles selon les hypothèses, le temps peut aider à comparer '
                              "l'exposition ; il ne quantifie pas à lui seul la gravité du dépôt. Les exercices "
                              'suivants changent distances, vitesses et contraintes ; ils ne prescrivent pas de '
                              'trajectoire réelle.\n'
                              '\n'
                              '**Format de travail :** 15 QCM de chapitre, quatre choix et une réponse correcte. '
                              'La matière 050 comporte 63 questions en 90 minutes dans le tableau EASA cité. Ce '
                              'chapitre ne couvre qu’une partie de la matière ; il ne constitue pas une épreuve '
                              'complète. Cadre EASA, même en anglais ou avec les unités US/britanniques.',
                   'content_en': '### Compare options against constraints\n'
                                 '\n'
                                 'Anti-icing aims to prevent accumulation; de-icing removes deposits. Operation '
                                 'depends on system design. Arbitrarily waiting for a thickness or copying '
                                 'another aircraft’s cycle is not a general method.\n'
                                 '\n'
                                 'First reject exits incompatible with terrain, performance, limitations, '
                                 'weather or fuel. Then compare exposure time for the remaining options. An '
                                 'available clearance does not establish feasibility; a calculated path is not a '
                                 'clearance.\n'
                                 '\n'
                                 '**Worked example, fixed profiles:** A reaches clear air 6 NM away at 6,500 ft, '
                                 'but the applicable minimum is 7,000 ft. B reaches clear air 18 NM away at '
                                 '8,000 ft with the same minimum. GS is 120 kt for both; all other constraints '
                                 'are verified for B. A would take 60 × 6/120 = 3 minutes but violates the '
                                 'minimum. B takes 9 minutes and meets the inputs. Selecting the shortest time '
                                 'first would favour an incompatible option. These numbers are not an '
                                 'operational procedure.\n'
                                 '\n'
                                 '### Reassess using observations\n'
                                 '\n'
                                 'If accumulation contradicts the forecast, update the decision using position, '
                                 'level, time, trend and observed effects. An old report, another level or '
                                 'another aircraft cannot guarantee your conditions. No report is not an '
                                 'observation of no ice.\n'
                                 '\n'
                                 '**Recall without the lesson:** separate observation and interpretation; '
                                 'explain why functioning equipment may not suffice; reject an incompatible exit '
                                 'before comparing times; identify missing information. A justified answer names '
                                 'the inputs used and the limits of its conclusion.\n'
                                 '\n'
                                 '### Transfer exercise: time alone must not decide\n'
                                 'For a new brief, create one row per option: planned altitude and minimum, '
                                 'aircraft approval, weather to exit, distance, GS, required and available fuel. '
                                 'Eliminate an incompatible option before ranking durations. Mark missing '
                                 'information “unknown”, not “favourable”.\n'
                                 '\n'
                                 '**Worked counterexample:** C takes seven minutes but climb performance in the '
                                 'expected conditions is unestablished; D takes eleven minutes with verified '
                                 "constraints. Durations alone do not establish C's feasibility. The missing "
                                 'evidence concerns its climb. If both routes are admissible under the '
                                 'assumptions, time can help compare exposure; it does not alone quantify '
                                 'deposit severity. The following exercises change distances, speeds and '
                                 'constraints; they do not prescribe real trajectories.\n'
                                 '\n'
                                 '**Practice format:** 15 chapter MCQs, four choices and one correct answer. '
                                 'Subject 050 has 63 questions in 90 minutes in the cited EASA table. This '
                                 'chapter covers only part of the subject and is not a full paper. EASA '
                                 'framework, including with English or US/British units.',
                   'difficulty': 'advanced',
                   'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — 050',
                   'tags': ['IR']}],
     'questions': [{'concept_id': 'IR_050_001_C01',
                    'concept_title': 'Identifier les conditions',
                    'level': 3,
                    'question': 'Deux segments sont à −6 °C : A en air sec sans précipitations, B dans un nuage '
                                'contenant des gouttes liquides. Pour le givrage de cellule par impact, quelle '
                                'comparaison est justifiée ?',
                    'question_en': 'Two legs are at −6 °C: A in dry air without precipitation, B in cloud '
                                   'containing liquid droplets. Which comparison is justified for impact icing '
                                   'on the airframe?',
                    'options': ['B fournit l’eau liquide nécessaire ; la température seule ne suffit pas',
                                'Le risque est identique puisque la température est identique',
                                'A présente le plus de risque car l’air sec refroidit davantage',
                                'B présente un risque seulement après congélation des gouttes en cristaux'],
                    'options_en': ['B supplies the required liquid water; temperature alone is insufficient',
                                   'Risk is identical because temperature is identical',
                                   'A presents greater risk because dry air cools more',
                                   'B poses a risk only after droplets freeze into crystals'],
                    'correct_answer': 'B fournit l’eau liquide nécessaire ; la température seule ne suffit pas',
                    'explanation': 'La différence utile est l’eau liquide en B. Les gouttes peuvent rester '
                                   'liquides sous 0 °C puis geler à l’impact. Une égalité de température ne '
                                   'donne pas une égalité d’exposition ; attendre des cristaux confond l’état '
                                   'avant et après congélation.',
                    'explanation_en': 'The useful difference is liquid water in B. Droplets may remain liquid '
                                      'below freezing, then freeze on impact. Equal temperature does not mean '
                                      'equal exposure; waiting for crystals confuses states before and after '
                                      'freezing.',
                    'use_exercise_svg': False,
                    'tags': ['IR', 'exam-mcq', 'exam-subject:050'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/documentLibrary/media/Advisory_Circular/AC_91-74B.pdf',
                                          'https://www.faa.gov/air_traffic/publications/atpubs/aim_html/chap7_section_1.html',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-air-operations?erules-id=ERULES-1963177438-13362'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 050'},
                   {'concept_id': 'IR_050_001_C01',
                    'concept_title': 'Identifier les conditions',
                    'level': 3,
                    'question': 'L’isotherme 0 °C est prévu à 5 000 ft, mais un profil local révèle une '
                                'inversion. Quelle information évalue le mieux une couche à 7 000 ft ?',
                    'question_en': 'The freezing level is forecast at 5,000 ft, but a local profile shows an '
                                   'inversion. Which information best assesses a layer at 7,000 ft?',
                    'options': ['Un gradient standard appliqué uniquement aux 2 000 ft d’écart',
                                'La température au sol extrapolée sans considérer l’inversion',
                                'La première altitude de 0 °C prise comme frontière unique',
                                'Le profil de température et l’eau liquide à ce niveau'],
                    'options_en': ['A standard lapse rate applied only to the 2,000 ft difference',
                                   'Surface temperature extrapolated without considering the inversion',
                                   'The first 0 °C altitude treated as the only boundary',
                                   'Temperature profile and liquid water at that level'],
                    'correct_answer': 'Le profil de température et l’eau liquide à ce niveau',
                    'explanation': 'L’inversion peut produire plusieurs passages à 0 °C. Il faut examiner la '
                                   'couche visée, avec son eau liquide. Les autres choix remplacent le profil '
                                   'réel par une simplification qui écarte précisément la donnée nouvelle.',
                    'explanation_en': 'An inversion may produce several crossings of 0 °C. Examine the target '
                                      'layer and its liquid water. The other choices replace the actual profile '
                                      'with a simplification that discards the new evidence.',
                    'use_exercise_svg': False,
                    'tags': ['IR', 'exam-mcq', 'exam-subject:050'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/documentLibrary/media/Advisory_Circular/AC_91-74B.pdf',
                                          'https://www.faa.gov/air_traffic/publications/atpubs/aim_html/chap7_section_1.html',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-air-operations?erules-id=ERULES-1963177438-13362'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 050'},
                   {'concept_id': 'IR_050_001_C01',
                    'concept_title': 'Identifier les conditions',
                    'level': 3,
                    'question': 'Deux observations à −5 °C : A, dix minutes dans un nuage ; B, cinq minutes dans '
                                'un autre. Quelle donnée supplémentaire est nécessaire pour comparer les dépôts, '
                                'à avion identique ?',
                    'question_en': 'Two observations at −5 °C: A, ten minutes in cloud; B, five minutes in '
                                   'another cloud. What additional data is needed to compare deposits on the '
                                   'same aircraft type?',
                    'options': ["Seulement l'égalité de température",
                                "Seulement l'altitude pression de chaque vol",
                                'Eau liquide et caractéristiques des gouttes dans chaque nuage',
                                'Seulement la durée totale déjà fournie'],
                    'options_en': ['Only the equal temperature',
                                   "Only each flight's pressure altitude",
                                   'Liquid water and droplet characteristics in each cloud',
                                   'Only the total duration already supplied'],
                    'correct_answer': 'Eau liquide et caractéristiques des gouttes dans chaque nuage',
                    'explanation': "L'exposition est plus longue en A mais les nuages ne sont pas caractérisés. "
                                   "Une accumulation dépend aussi de l'eau et des gouttes ; ni température égale "
                                   'ni altitude seule ne rend les taux de dépôt comparables.',
                    'explanation_en': 'Exposure is longer in A, but the clouds are uncharacterised. Accumulation '
                                      'also depends on water and droplets; equal temperature or altitude alone '
                                      'cannot establish comparable deposition rates.',
                    'use_exercise_svg': False,
                    'tags': ['IR', 'exam-mcq', 'exam-subject:050'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/documentLibrary/media/Advisory_Circular/AC_91-74B.pdf',
                                          'https://www.faa.gov/air_traffic/publications/atpubs/aim_html/chap7_section_1.html',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-air-operations?erules-id=ERULES-1963177438-13362'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 050'},
                   {'concept_id': 'IR_050_001_C01',
                    'concept_title': 'Identifier les conditions',
                    'level': 3,
                    'question': 'De la pluie traverse une couche positive puis une couche sous 0 °C en restant '
                                'liquide. Quelle interprétation décrit le danger d’impact ?',
                    'question_en': 'Rain crosses an above-freezing layer, then a subzero layer while remaining '
                                   'liquid. Which interpretation describes the impact hazard?',
                    'options': ['Le risque dépend uniquement de la température à l’aérodrome',
                                'Les gouttes surfondues peuvent geler sur les surfaces rencontrées',
                                'La couche chaude supprime le risque dans la couche froide suivante',
                                'La pluie devient nécessairement neige dès le passage sous 0 °C'],
                    'options_en': ['Risk depends only on temperature at the aerodrome',
                                   'Supercooled drops may freeze on surfaces they strike',
                                   'The warm layer removes risk in the following cold layer',
                                   'Rain necessarily becomes snow as soon as it crosses below 0 °C'],
                    'correct_answer': 'Les gouttes surfondues peuvent geler sur les surfaces rencontrées',
                    'explanation': 'L’eau reste liquide d’après l’énoncé : elle est surfondue dans la couche '
                                   'froide. La couche chaude précédente ne protège pas la suite du trajet. La '
                                   'température au sol ne décrit pas ce profil vertical.',
                    'explanation_en': 'The stem says the water remains liquid: it is supercooled in the cold '
                                      'layer. The preceding warm layer does not protect the rest of the route. '
                                      'Surface temperature does not describe this vertical profile.',
                    'use_exercise_svg': False,
                    'tags': ['IR', 'exam-mcq', 'exam-subject:050'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/documentLibrary/media/Advisory_Circular/AC_91-74B.pdf',
                                          'https://www.faa.gov/air_traffic/publications/atpubs/aim_html/chap7_section_1.html',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-air-operations?erules-id=ERULES-1963177438-13362'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 050'},
                   {'concept_id': 'IR_050_001_C02',
                    'concept_title': 'Distinguer les dépôts',
                    'level': 3,
                    'question': 'Un dépôt devient blanc, opaque et rugueux. Quel mécanisme est le plus '
                                'compatible avec cette observation ?',
                    'question_en': 'A deposit becomes white, opaque and rough. Which mechanism best matches this '
                                   'observation?',
                    'options': ['Une congélation rapide emprisonnant de l’air',
                                'Un étalement prolongé avant congélation donnant surtout une couche transparente',
                                'Une fonte complète suivie d’évaporation sans dépôt',
                                'Une condensation liquide restant au-dessus de 0 °C'],
                    'options_en': ['Rapid freezing trapping air',
                                   'Prolonged spreading before freezing, mainly producing a transparent layer',
                                   'Complete melting followed by evaporation without a deposit',
                                   'Liquid condensation remaining above 0 °C'],
                    'correct_answer': 'Une congélation rapide emprisonnant de l’air',
                    'explanation': 'L’aspect évoque le givre blanc. L’étalement avant congélation correspond '
                                   'plutôt au verglas. Identifier le mécanisme n’autorise pas à classer le dépôt '
                                   'comme bénin : position, forme et effets restent à examiner.',
                    'explanation_en': 'The appearance suggests rime. Spreading before freezing is more '
                                      'characteristic of clear ice. Identifying the mechanism does not make the '
                                      'deposit harmless: location, shape and effects still matter.',
                    'use_exercise_svg': False,
                    'tags': ['IR', 'exam-mcq', 'exam-subject:050'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/documentLibrary/media/Advisory_Circular/AC_91-74B.pdf',
                                          'https://www.faa.gov/air_traffic/publications/atpubs/aim_html/chap7_section_1.html',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-air-operations?erules-id=ERULES-1963177438-13362'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 050'},
                   {'concept_id': 'IR_050_001_C02',
                    'concept_title': 'Distinguer les dépôts',
                    'level': 3,
                    'question': 'Un dépôt transparent est difficile à distinguer sur une surface mal éclairée. '
                                'Quelle conclusion est fondée ?',
                    'question_en': 'A transparent deposit is difficult to distinguish on a poorly lit surface. '
                                   'Which conclusion is justified?',
                    'options': ['Une faible visibilité du dépôt indique une faible accumulation',
                                'Un dépôt transparent est moins dangereux à toute épaisseur',
                                'L’absence de rugosité visible prouve le maintien du profil initial',
                                'La visibilité du dépôt ne quantifie pas sa gravité'],
                    'options_en': ['Poor deposit visibility indicates little accumulation',
                                   'A transparent deposit is less hazardous at every thickness',
                                   'No visible roughness proves the original profile is maintained',
                                   'Deposit visibility does not quantify severity'],
                    'correct_answer': 'La visibilité du dépôt ne quantifie pas sa gravité',
                    'explanation': 'Éclairage et aspect limitent l’observation. Celle-ci ne mesure ni épaisseur, '
                                   'ni géométrie, ni effets. Les autres propositions transforment une difficulté '
                                   'de détection en preuve de sécurité ; il faut poursuivre la surveillance '
                                   'adaptée.',
                    'explanation_en': 'Lighting and appearance limit observation. It measures neither thickness, '
                                      'geometry nor effects. The other choices turn detection difficulty into '
                                      'evidence of safety; appropriate monitoring must continue.',
                    'use_exercise_svg': False,
                    'tags': ['IR', 'exam-mcq', 'exam-subject:050'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/documentLibrary/media/Advisory_Circular/AC_91-74B.pdf',
                                          'https://www.faa.gov/air_traffic/publications/atpubs/aim_html/chap7_section_1.html',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-air-operations?erules-id=ERULES-1963177438-13362'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 050'},
                   {'concept_id': 'IR_050_001_C02',
                    'concept_title': 'Distinguer les dépôts',
                    'level': 3,
                    'question': 'Le dégivrage nettoie sa zone, mais une bande de glace apparaît derrière la '
                                'surface protégée. Quelle interprétation doit guider la réévaluation ?',
                    'question_en': 'De-icing clears its own area, but a band of ice appears aft of the protected '
                                   'surface. Which interpretation should guide reassessment?',
                    'options': ['La bande prouve nécessairement une panne électrique du dégivrage',
                                'La protection des sondes compense cette accumulation',
                                'L’accumulation peut dépasser la zone que le système protège',
                                'Le nettoyage du bord d’attaque prouve la maîtrise du dépôt total'],
                    'options_en': ['The band necessarily proves an electrical de-icing failure',
                                   'Probe protection compensates for this accumulation',
                                   'Accumulation may extend beyond the area protected by the system',
                                   'A clear leading edge proves control of the entire deposit'],
                    'correct_answer': 'L’accumulation peut dépasser la zone que le système protège',
                    'explanation': 'Le système peut fonctionner localement sans maîtriser les zones arrière. '
                                   'Appliquer les procédures et réévaluer la sortie ; cette observation ne '
                                   'diagnostique pas à elle seule une panne électrique ni un type exact de '
                                   'gouttes.',
                    'explanation_en': 'The system may work locally without controlling aft areas. Apply '
                                      'procedures and reassess the exit; this observation alone diagnoses '
                                      'neither electrical failure nor an exact droplet type.',
                    'use_exercise_svg': False,
                    'tags': ['IR', 'exam-mcq', 'exam-subject:050'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/documentLibrary/media/Advisory_Circular/AC_91-74B.pdf',
                                          'https://www.faa.gov/air_traffic/publications/atpubs/aim_html/chap7_section_1.html',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-air-operations?erules-id=ERULES-1963177438-13362'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 050'},
                   {'concept_id': 'IR_050_001_C02',
                    'concept_title': 'Distinguer les dépôts',
                    'level': 3,
                    'question': 'Exercice de préparation : le manuel de l’avion exclut la pluie verglaçante. Le '
                                'trajet prévu en rencontre, mais tous les équipements fonctionnent. Quelle '
                                'décision respecte ces données ?',
                    'question_en': 'Planning exercise: the aircraft manual excludes freezing rain. The planned '
                                   'route encounters it, but all equipment works. Which decision respects these '
                                   'inputs?',
                    'options': ['Remplacer la restriction par un compte rendu favorable d’un autre type d’avion',
                                'Modifier le projet pour éviter les conditions exclues',
                                'Garder le trajet en activant tous les équipements plus tôt',
                                'Garder le trajet si l’exposition dure moins de cinq minutes'],
                    'options_en': ['Replace the restriction with a favourable report from another aircraft type',
                                   'Change the plan to avoid the excluded conditions',
                                   'Keep the route and activate all equipment earlier',
                                   'Keep the route if exposure lasts under five minutes'],
                    'correct_answer': 'Modifier le projet pour éviter les conditions exclues',
                    'explanation': 'La restriction est explicite : équipement fonctionnel ne signifie pas '
                                   'domaine illimité. Ni un délai inventé de cinq minutes ni l’expérience d’un '
                                   'autre avion ne modifient ce manuel. La décision se prend ici avant '
                                   'exposition.',
                    'explanation_en': 'The restriction is explicit: functioning equipment does not mean an '
                                      'unlimited envelope. Neither an invented five-minute allowance nor another '
                                      'aircraft’s experience changes the manual. This decision is made before '
                                      'exposure.',
                    'use_exercise_svg': False,
                    'tags': ['IR', 'exam-mcq', 'exam-subject:050'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/documentLibrary/media/Advisory_Circular/AC_91-74B.pdf',
                                          'https://www.faa.gov/air_traffic/publications/atpubs/aim_html/chap7_section_1.html',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-air-operations?erules-id=ERULES-1963177438-13362'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 050'},
                   {'concept_id': 'IR_050_001_C03',
                    'concept_title': 'Comprendre les effets',
                    'level': 3,
                    'question': 'En accumulation observée, altitude, configuration et puissance restent '
                                "constantes tandis que l'IAS baisse. Quel raisonnement distingue effet possible "
                                'et preuve ?',
                    'question_en': 'During observed accumulation, altitude, configuration and power stay '
                                   'constant while IAS falls. Which reasoning separates a possible effect from '
                                   'proof?',
                    'options': ['Une hausse de traînée est cohérente, à recouper avec les autres paramètres',
                                "La baisse d'IAS mesure directement la nouvelle vitesse de décrochage",
                                "La baisse d'IAS quantifie à elle seule la masse de glace",
                                "L'altitude constante prouve que la portance maximale est inchangée"],
                    'options_en': ['Increased drag is consistent and needs cross-checking',
                                   'Falling IAS directly measures the new stall speed',
                                   'Falling IAS alone quantifies ice mass',
                                   'Constant altitude proves maximum lift is unchanged'],
                    'correct_answer': 'Une hausse de traînée est cohérente, à recouper avec les autres '
                                      'paramètres',
                    'explanation': 'La tendance est compatible avec une augmentation de traînée mais ne fournit '
                                   'ni vitesse de décrochage ni masse déposée. Maintenir la portance nécessaire '
                                   'au palier ne mesure pas la portance maximale disponible.',
                    'explanation_en': 'The trend is consistent with increased drag but supplies neither stall '
                                      'speed nor deposited mass. Maintaining lift needed for level flight does '
                                      'not measure maximum available lift.',
                    'use_exercise_svg': False,
                    'explanation_illustration': '/illustrations/ir-icing-effects.svg',
                    'tags': ['IR', 'exam-mcq', 'exam-subject:050'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/documentLibrary/media/Advisory_Circular/AC_91-74B.pdf',
                                          'https://www.faa.gov/air_traffic/publications/atpubs/aim_html/chap7_section_1.html',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-air-operations?erules-id=ERULES-1963177438-13362'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 050'},
                   {'concept_id': 'IR_050_001_C03',
                    'concept_title': 'Comprendre les effets',
                    'level': 3,
                    'question': 'Sous pilote automatique, l’altitude reste stable mais l’IAS décroît pendant une '
                                'accumulation. Que démontre l’altitude stable ?',
                    'question_en': 'On autopilot, altitude remains stable but IAS decreases during accumulation. '
                                   'What does stable altitude establish?',
                    'options': ['Une marge au décrochage constante',
                                'L’absence d’effet aérodynamique puisque la trajectoire est tenue',
                                'L’efficacité de tous les dispositifs de protection',
                                'Seulement le maintien d’altitude observé, pas celui des marges'],
                    'options_en': ['A constant stall margin',
                                   'No aerodynamic effect because the flight path is maintained',
                                   'Effectiveness of all protection systems',
                                   'Only observed altitude holding, not preservation of margins'],
                    'correct_answer': 'Seulement le maintien d’altitude observé, pas celui des marges',
                    'explanation': 'Un paramètre maintenu peut masquer la tendance des autres. La baisse d’IAS '
                                   'doit être recoupée et traitée selon les procédures de l’avion. L’altitude ne '
                                   'mesure ni le dépôt ni la marge de vitesse.',
                    'explanation_en': 'One maintained parameter can hide trends in others. Decreasing IAS must '
                                      'be cross-checked and handled under aircraft procedures. Altitude measures '
                                      'neither accumulation nor airspeed margin.',
                    'use_exercise_svg': False,
                    'explanation_illustration': '/illustrations/ir-icing-effects.svg',
                    'tags': ['IR', 'exam-mcq', 'exam-subject:050'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/documentLibrary/media/Advisory_Circular/AC_91-74B.pdf',
                                          'https://www.faa.gov/air_traffic/publications/atpubs/aim_html/chap7_section_1.html',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-air-operations?erules-id=ERULES-1963177438-13362'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 050'},
                   {'concept_id': 'IR_050_001_C03',
                    'concept_title': 'Comprendre les effets',
                    'level': 3,
                    'question': 'Deux chaînes IAS indépendantes concordent et les sondes sont chauffées. Une '
                                'aile présente une contamination visible. Quelle question reste non résolue par '
                                'cet accord instrumental ?',
                    'question_en': 'Two independent IAS chains agree and probes are heated. A wing is visibly '
                                   'contaminated. Which question remains unresolved by instrument agreement?',
                    'options': ["La présence de contamination signalée dans l'énoncé",
                                'La valeur numérique affichée par chaque chaîne',
                                "La marge aérodynamique de l'aile contaminée",
                                'La concordance actuelle entre les deux IAS'],
                    'options_en': ['Presence of the stated visible contamination',
                                   'The numerical value displayed by each chain',
                                   'Aerodynamic margin of the contaminated wing',
                                   'Current agreement between both IAS values'],
                    'correct_answer': "La marge aérodynamique de l'aile contaminée",
                    'explanation': "L'accord renseigne sur la cohérence des mesures, pas sur la portance "
                                   'maximale du profil contaminé. Les autres propositions sont des observations '
                                   "déjà données, non l'inconnue à préserver dans le raisonnement.",
                    'explanation_en': 'Agreement informs measurement consistency, not the contaminated '
                                      "aerofoil's maximum lift. The other choices are already supplied "
                                      'observations, not the remaining unknown.',
                    'use_exercise_svg': False,
                    'explanation_illustration': '/illustrations/ir-icing-effects.svg',
                    'tags': ['IR', 'exam-mcq', 'exam-subject:050'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/documentLibrary/media/Advisory_Circular/AC_91-74B.pdf',
                                          'https://www.faa.gov/air_traffic/publications/atpubs/aim_html/chap7_section_1.html',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-air-operations?erules-id=ERULES-1963177438-13362'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 050'},
                   {'concept_id': 'IR_050_001_C04',
                    'concept_title': 'Préparer une sortie',
                    'level': 3,
                    'question': 'Un manuel prescrit l’activation d’un système avant l’entrée dans les conditions '
                                'concernées. Quel raisonnement respecte la consigne ?',
                    'question_en': 'A manual requires system activation before entering the relevant conditions. '
                                   'Which reasoning respects the instruction?',
                    'options': ['Retarder l’activation jusqu’à une baisse mesurable de vitesse',
                                'Prévenir l’accumulation en suivant le déclenchement prescrit',
                                'Attendre une épaisseur visible pour confirmer l’utilité du système',
                                'Transposer le cycle d’un autre avion équipé de boudins'],
                    'options_en': ['Delay activation until measurable speed loss',
                                   'Prevent accumulation using the prescribed activation trigger',
                                   'Wait for visible thickness to confirm the system is useful',
                                   'Copy the cycle from another aircraft equipped with boots'],
                    'correct_answer': 'Prévenir l’accumulation en suivant le déclenchement prescrit',
                    'explanation': 'Le déclenchement est fourni : ne pas le remplacer par un seuil personnel. '
                                   'Attendre glace ou perte de vitesse transforme la prévention en réaction '
                                   'tardive. Les systèmes n’ont pas tous les mêmes procédures.',
                    'explanation_en': 'The activation trigger is supplied: do not replace it with a personal '
                                      'threshold. Waiting for ice or speed loss turns prevention into a late '
                                      'reaction. Different systems need not use the same procedures.',
                    'use_exercise_svg': False,
                    'tags': ['IR', 'exam-mcq', 'exam-subject:050'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/documentLibrary/media/Advisory_Circular/AC_91-74B.pdf',
                                          'https://www.faa.gov/air_traffic/publications/atpubs/aim_html/chap7_section_1.html',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-air-operations?erules-id=ERULES-1963177438-13362'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 050'},
                   {'concept_id': 'IR_050_001_C04',
                    'concept_title': 'Préparer une sortie',
                    'level': 3,
                    'question': 'La prévision annonçait peu de risque, mais une accumulation nouvelle s’étend '
                                'hors des zones protégées. Quel élément impose de réexaminer le plan ?',
                    'question_en': 'The forecast suggested little risk, but new accumulation extends beyond '
                                   'protected areas. Which evidence calls for reassessment of the plan?',
                    'options': ['L’observation actuelle et son évolution sur cet avion',
                                'L’absence de signalement une heure auparavant',
                                'Le seul classement peu risqué retenu au briefing',
                                'Le fonctionnement apparent du système dans sa zone'],
                    'options_en': ['Current observation and its trend on this aircraft',
                                   'No report one hour earlier',
                                   'Only the low-risk classification used in the briefing',
                                   'The system apparently working in its own area'],
                    'correct_answer': 'L’observation actuelle et son évolution sur cet avion',
                    'explanation': 'L’observation remet en cause les hypothèses du plan. Appliquer les '
                                   'procédures et réévaluer la sortie sans attendre une nouvelle prévision. Un '
                                   'système actif et un ancien compte rendu ne neutralisent pas le dépôt '
                                   'constaté.',
                    'explanation_en': 'The observation challenges assumptions behind the plan. Apply procedures '
                                      'and reassess the exit without waiting for a new forecast. An active '
                                      'system and an old report do not negate the deposit observed.',
                    'use_exercise_svg': False,
                    'tags': ['IR', 'exam-mcq', 'exam-subject:050'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/documentLibrary/media/Advisory_Circular/AC_91-74B.pdf',
                                          'https://www.faa.gov/air_traffic/publications/atpubs/aim_html/chap7_section_1.html',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-air-operations?erules-id=ERULES-1963177438-13362'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 050'},
                   {'concept_id': 'IR_050_001_C04',
                    'concept_title': 'Préparer une sortie',
                    'level': 3,
                    'question': 'Sorties fictives, profils fixes : A 12 NM à GS 120 kt, altitude 6 500 '
                                'ft/minimum 7 000 ; B 21 NM à GS 140 kt, 8 000/7 000 ; C 16 NM à GS 120 kt, 8 '
                                '000/7 000, consommation 10 kg. Carburant utilisable pour ces segments après '
                                'protection des besoins ultérieurs : 8 kg ; B consomme 7 kg. Météo, performances '
                                'et autres contraintes vérifiées. Quelle option satisfait altitude et carburant '
                                '?',
                    'question_en': 'Fictional exits, fixed profiles: A 12 NM at GS 120 kt, altitude 6,500 '
                                   'ft/minimum 7,000; B 21 NM at GS 140 kt, 8,000/7,000; C 16 NM at GS 120 kt, '
                                   '8,000/7,000, consumption 10 kg. Fuel usable for these segments after '
                                   'protecting later requirements: 8 kg; B uses 7 kg. Weather, performance and '
                                   'other constraints checked. Which meets altitude and fuel limits?',
                    'options': ['A, 6 minutes', 'C, 8 minutes', 'B, 21 minutes', 'B, 9 minutes'],
                    'options_en': ['A, 6 minutes', 'C, 8 minutes', 'B, 21 minutes', 'B, 9 minutes'],
                    'correct_answer': 'B, 9 minutes',
                    'explanation': 'A viole le minimum ; C exige 10 kg pour 8 disponibles. B respecte les deux '
                                   'contraintes et demande 60 × 21/140 = 9 minutes. La durée la plus courte '
                                   "n'est pas la première sélection ; les besoins ultérieurs sont protégés par "
                                   'hypothèse, pas supposés nuls.',
                    'explanation_en': 'A violates the minimum; C needs 10 kg with 8 available. B meets both '
                                      'constraints and takes 60 × 21/140 = 9 minutes. Shortest time is not the '
                                      'first filter; later fuel requirements are protected by assumption, not '
                                      'assumed zero.',
                    'use_exercise_svg': True,
                    'explanation_illustration': '/illustrations/ir-050-001.svg',
                    'tags': ['IR', 'exam-mcq', 'exam-subject:050'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/documentLibrary/media/Advisory_Circular/AC_91-74B.pdf',
                                          'https://www.faa.gov/air_traffic/publications/atpubs/aim_html/chap7_section_1.html',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-air-operations?erules-id=ERULES-1963177438-13362'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 050'},
                   {'concept_id': 'IR_050_001_C04',
                    'concept_title': 'Préparer une sortie',
                    'level': 3,
                    'question': 'Route A sort du nuage en 5 minutes mais exige une montée dont la faisabilité en '
                                "conditions givrantes n'est pas établie. B sort en 10 minutes avec contraintes "
                                'vérifiées. Que manque-t-il pour comparer A comme option admissible ?',
                    'question_en': 'Route A exits cloud in five minutes but requires a climb whose feasibility '
                                   'in icing is unestablished. B exits in ten minutes with constraints checked. '
                                   'What is missing before treating A as admissible?',
                    'options': ["Un compte rendu favorable d'un autre avion à un autre niveau",
                                'La confirmation que cinq minutes sont inférieures à dix',
                                'La capacité de montée de cet avion dans ces conditions',
                                'Un calcul plus précis des cinq minutes seulement'],
                    'options_en': ['A favourable report from another aircraft at another level',
                                   'Confirmation that five minutes is less than ten',
                                   "This aircraft's climb capability in those conditions",
                                   'Only a more precise five-minute calculation'],
                    'correct_answer': 'La capacité de montée de cet avion dans ces conditions',
                    'explanation': 'La contrainte inconnue est la montée. Une durée avantageuse ou un compte '
                                   "rendu non représentatif ne l'établit pas. L'exercice ne demande pas "
                                   "d'inventer une performance ; conserver explicitement cette incertitude avant "
                                   'de comparer les options.',
                    'explanation_en': 'The unknown constraint is climb capability. A shorter duration or '
                                      'unrepresentative report does not establish it. Do not invent performance; '
                                      'retain the uncertainty before comparing options.',
                    'use_exercise_svg': False,
                    'tags': ['IR', 'exam-mcq', 'exam-subject:050'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/documentLibrary/media/Advisory_Circular/AC_91-74B.pdf',
                                          'https://www.faa.gov/air_traffic/publications/atpubs/aim_html/chap7_section_1.html',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-air-operations?erules-id=ERULES-1963177438-13362'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 050'}]}
    save(data)
