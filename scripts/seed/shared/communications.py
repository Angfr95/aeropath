"""Bilingual teaching content, revised after individual pedagogical review."""
from .authoring import save


def build(program):
    data = {'metadata': {'module_id': 'ATPL_090_001',
                  'program': 'ATPL',
                  'subject': 'communications',
                  'chapter_title': 'Clairances et boucle de collationnement',
                  'chapter_title_en': 'Clearances and the readback loop',
                  'reference_checked_on': '2026-09-29',
                  'learning_objective': 'Reconstruire une clairance après messages successifs, détecter '
                                        'omissions et substitutions et préserver conditions et limites au '
                                        'collationnement.',
                  'learning_objective_en': 'Reconstruct clearance after successive messages, detect omissions '
                                           'and substitutions and preserve conditions and limits in readback.',
                  'estimated_duration_minutes': 45,
                  'difficulty': 'advanced',
                  'prerequisites': ['Calcul des proportions et des durées ; bases PPL du sujet'],
                  'exam_subject': '090',
                  'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 090',
                  'official_subject_format': {'questions': 34,
                                              'minutes': 60,
                                              'pass_percent': 75,
                                              'source': 'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67'},
                  'sources': [{'title': 'EASA — programme Part-FCL',
                               'url': 'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8'},
                              {'title': 'EASA — format de la matière, AMC1 ARA.FCL.300(b)',
                               'url': 'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67'},
                              {'title': 'EASA SERA.8015 — clairances et collationnement',
                               'url': 'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-standardised-european?erules-id=ERULES-1963177438-9888'},
                              {'title': 'Référence pédagogique vérifiée',
                               'url': 'https://www.ecologie.gouv.fr/sites/default/files/documents/sera_complet_version17012026.pdf'}],
                  'content_status': 'draft',
                  'coverage_status': 'partial',
                  'scope': 'ATPL(A), France / EASA',
                  'quiz_format': 'chapter_practice',
                  'tags': ['ATPL', 'mcq-four-options'],
                  'illustration': {'src': '/illustrations/atpl-090-001.svg',
                                   'alt': 'Clairances et boucle de collationnement'}},
     'concepts': [{'concept_id': 'ATPL_090_001_C01',
                   'title': 'Identifier le destinataire',
                   'title_en': 'Identify the addressee',
                   'content': 'Avant une émission, écouter pour éviter de couper une transmission. Vérifier que '
                              'la clairance concerne son indicatif, particulièrement lorsque deux indicatifs se '
                              'ressemblent. Une transmission partiellement entendue ne doit pas être '
                              'reconstruite à partir de ce que l’on attendait : demander la répétition du '
                              'passage manquant. Employer la phraséologie publiée et une parole claire ; ce '
                              'chapitre travaille la compréhension, sans remplacer un entraînement oral.\n'
                              '\n'
                              '### Tenir un état des instructions\n'
                              'Préparer une grille mentale ou écrite : destinataire, route/cap, niveau, vitesse, '
                              "calage, code et canal. Une instruction ultérieure remplace les éléments qu'elle "
                              "modifie ; ne pas recopier tous les éléments d'un message adressé à un autre "
                              'avion. Une donnée masquée reste inconnue même si elle paraît prévisible.\n'
                              '\n'
                              'Exemple fictif : AER321 a reçu cap 250, 4 000 ft, 170 kt. AER312 reçoit ensuite '
                              "cap 280. Le cap d'AER321 reste 250 : la dernière valeur entendue sur la fréquence "
                              "n'est pas nécessairement sa dernière instruction. Si seul le destinataire est "
                              "masqué, faire lever le doute avant d'attribuer la nouvelle instruction. "
                              "L'abréviation d'un indicatif d'immatriculation peut être initiée par le sol dans "
                              'les conditions SERA ; un indicatif exploitant + numéro de vol comme AER321 ne '
                              "s'abrège pas.",
                   'content_en': 'Listen before transmitting to avoid interrupting another transmission. Check '
                                 'that the clearance addresses your callsign, particularly when similar '
                                 'callsigns share the frequency. Do not reconstruct a partly heard transmission '
                                 'from what you expected: request repetition of the missing part. Use published '
                                 'phraseology and clear speech; this chapter practises comprehension and does '
                                 'not replace oral training.\n'
                                 '\n'
                                 '### Maintain an instruction state\n'
                                 'Use a mental or written grid: addressee, route/heading, level, speed, setting, '
                                 'code and channel. A later instruction replaces the items it amends; do not '
                                 "copy items from another aircraft's message. A blocked item remains unknown "
                                 'even when predictable.\n'
                                 '\n'
                                 'Fictional example: AER321 has heading 250, 4,000 ft and 170 kt. AER312 then '
                                 'receives heading 280. AER321 keeps 250: the latest number heard is not '
                                 'necessarily its latest instruction. If only the addressee is blocked, resolve '
                                 'that uncertainty before assigning the amendment. Ground may initiate a '
                                 'registration abbreviation under SERA conditions; an '
                                 'operator-plus-flight-number callsign such as AER321 is not abbreviated.',
                   'difficulty': 'advanced',
                   'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — 090',
                   'tags': ['ATPL']},
                  {'concept_id': 'ATPL_090_001_C02',
                   'title': 'Collationner les éléments de sécurité',
                   'title_en': 'Read back safety-related elements',
                   'content': 'SERA.8015(e) impose le collationnement des clairances de route et des '
                              'autorisations ou instructions relatives aux pistes, notamment entrée, traversée '
                              'et maintien avant piste. Sont aussi concernés piste en service, calage '
                              'altimétrique, code SSR, nouveau canal de communication, niveau, cap, vitesse et '
                              'niveau de transition. Un simple accusé de réception ne remplace pas ces éléments. '
                              'Inclure son indicatif selon la procédure de radiotéléphonie ; le contrôleur '
                              'écoute et corrige les écarts.\n'
                              '\n'
                              '### Détecter omission et substitution\n'
                              'Exemple écrit : « AER321, cap 250, maintenez 4 000 ft, QNH 1014, code 4251 ». '
                              'Réponse analysée : « Cap 250, maintenons 4 000 ft, QNH 1041, AER321 ». Deux '
                              'défauts distincts : QNH substitué et code omis. Un contrôle qui ne vérifie que le '
                              "nombre d'éléments peut manquer la substitution ; un contrôle qui ne compare que "
                              "les valeurs présentes peut manquer l'omission. La boucle doit permettre de "
                              'corriger les deux.\n'
                              '\n'
                              "ROGER accuse réception sans répéter les valeurs ; WILCO n'est pas un substitut au "
                              "collationnement requis. STANDBY n'accorde aucune demande en attente. Les nombres "
                              'écrits dans les exercices ne dispensent pas de la prononciation réglementaire en '
                              'entraînement oral.',
                   'content_en': 'SERA.8015(e) requires readback of route clearances and runway clearances or '
                                 'instructions, including entry, crossing and holding short. The runway in use, '
                                 'altimeter setting, SSR code, newly assigned communication channel, level, '
                                 'heading, speed and transition level are also included. A simple '
                                 'acknowledgement does not replace these elements. Include your callsign in '
                                 'accordance with radiotelephony procedure; the controller listens and corrects '
                                 'discrepancies.\n'
                                 '\n'
                                 '### Detect omission and substitution\n'
                                 'Written example: “AER321, heading 250, maintain 4,000 ft, QNH 1014, squawk '
                                 '4251”. Analysed reply: “Heading 250, maintaining 4,000 ft, QNH 1041, AER321”. '
                                 'Two defects: substituted QNH and omitted code. Counting items alone may miss '
                                 'substitution; comparing only stated values may miss omission. The loop must '
                                 'repair both.\n'
                                 '\n'
                                 'ROGER acknowledges receipt without repeating values; WILCO does not substitute '
                                 'for required readback. STANDBY grants no pending request. Printed numbers do '
                                 'not replace prescribed pronunciation in oral practice.',
                   'difficulty': 'advanced',
                   'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — 090',
                   'tags': ['ATPL']},
                  {'concept_id': 'ATPL_090_001_C03',
                   'title': 'Fermer la boucle',
                   'title_en': 'Close the loop',
                   'content': 'Exemple pédagogique, hors phraséologie complète : AER123 reçoit cap 270, maintien '
                              '3 000 ft, QNH 1016. Son collationnement indique par erreur 2 000 ft. Le '
                              'contrôleur doit corriger cet écart ; le pilote doit comprendre et collationner la '
                              'correction. Une absence de correction entendue ne justifie pas de conserver un '
                              'doute. En cas de clairance impossible à respecter, informer l’ATC et demander une '
                              'solution adaptée, sans accepter silencieusement.\n'
                              '\n'
                              '### Corriger sans perdre les éléments inchangés\n'
                              'Reprendre le cas AER321 : cap 250, 4 000 ft, 170 kt. Puis « cap 260, correction, '
                              'cap 270 ». Le cap corrigé est 270 ; niveau et vitesse ne sont pas modifiés par '
                              'cette seule correction. Un collationnement « cap 260 » conserve une version '
                              'explicitement corrigée. Une correction du contrôleur doit être comprise et '
                              "répétée ; s'il manque la valeur corrigée, demander sa répétition.\n"
                              '\n'
                              '**Travail sans les réponses :** noter les trois paramètres, barrer uniquement '
                              "ceux qui changent, puis relire l'état obtenu. Ensuite analyser le dialogue dans "
                              "les QCM avec de nouveaux indicatifs et chiffres. Ne pas confondre l'état autorisé "
                              "avec l'état réel de l'avion : le premier décrit l'instruction, le second exige "
                              'surveillance et exécution.',
                   'content_en': 'Training example, not complete phraseology: AER123 receives heading 270, '
                                 'maintain 3,000 ft, QNH 1016. Its readback incorrectly states 2,000 ft. The '
                                 'controller must correct the discrepancy; the pilot must understand and read '
                                 'back the correction. Not hearing a correction does not resolve an existing '
                                 'doubt. If unable to comply with a clearance, inform ATC and request a suitable '
                                 'alternative rather than silently accepting.\n'
                                 '\n'
                                 '### Correct without losing unchanged items\n'
                                 'Return to AER321: heading 250, 4,000 ft, 170 kt. Then “heading 260, '
                                 'correction, heading 270”. Corrected heading is 270; this correction alone does '
                                 'not change level or speed. Reading back 260 retains a version explicitly '
                                 'corrected. An ATC correction must be understood and read back; request '
                                 'repetition if its value is missing.\n'
                                 '\n'
                                 '**Work before seeing answers:** record the three parameters, cross out only '
                                 'amended ones and read the resulting state. Then analyse MCQ exchanges with new '
                                 'callsigns and numbers. Distinguish the cleared state from actual aircraft '
                                 'state: the first describes instructions, the second needs monitoring and '
                                 'execution.',
                   'difficulty': 'advanced',
                   'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — 090',
                   'tags': ['ATPL']},
                  {'concept_id': 'ATPL_090_001_C04',
                   'title': 'Distinguer les autorisations',
                   'title_en': 'Distinguish authorisations',
                   'content': 'Une autorisation de roulage vers un point d’attente ne constitue pas une '
                              'autorisation de traverser la piste. Une clairance d’approche ne constitue pas une '
                              'autorisation d’atterrir. Préparer le changement de fréquence, vérifier le canal '
                              'assigné et établir le contact requis. Si le contact échoue, appliquer les '
                              'procédures de communication et de panne pertinentes ; ne pas inventer une '
                              'nouvelle clairance. Le pilote reste responsable de la conduite sûre du vol.\n'
                              '\n'
                              '### Identifier une limite, une condition et une impossibilité\n'
                              "« Roulez point d'attente piste 18 » ne décrit pas une traversée ; « alignez-vous "
                              "et attendez » n'autorise pas le décollage. Lorsqu'une clairance conditionnelle "
                              'est utilisée selon les règles applicables, sa condition fait partie de son sens : '
                              'identifier le trafic concerné et conserver cette condition au collationnement. '
                              'Une condition non comprise ne doit pas devenir une autorisation '
                              'inconditionnelle.\n'
                              '\n'
                              '**Cas de contrainte :** une vitesse demandée est sous la vitesse minimale '
                              "applicable à la configuration donnée. Informer l'ATC avec UNABLE, la raison et "
                              'une valeur réalisable proposée ; ne pas accepter puis improviser une autre '
                              'vitesse en silence. Si un changement de canal échoue, rechercher le contact selon '
                              'les procédures prévues, en conservant la distinction entre communication et '
                              'clairance. La charge de travail ne justifie pas de transformer une supposition en '
                              'donnée confirmée.\n'
                              '\n'
                              '**Format de travail :** 15 QCM de chapitre, quatre choix et une réponse correcte. '
                              'La matière 090 comporte 34 questions en 60 minutes dans le tableau EASA cité. Ce '
                              'chapitre ne couvre qu’une partie de la matière ; il ne constitue pas une épreuve '
                              'complète. Cadre EASA, même en anglais ou avec les unités US/britanniques.',
                   'content_en': 'Taxi clearance to a holding point is not clearance to cross the runway. '
                                 'Approach clearance is not landing clearance. Prepare a frequency change, '
                                 'verify the assigned channel and establish the required contact. If contact '
                                 'fails, apply the relevant communication and failure procedures; do not invent '
                                 'a new clearance. The pilot remains responsible for safe conduct of the '
                                 'flight.\n'
                                 '\n'
                                 '### Identify a limit, condition and inability\n'
                                 '“Taxi holding point runway 18” does not clear a crossing; “line up and wait” '
                                 'does not clear take-off. Where conditional clearance is used under applicable '
                                 'rules, the condition is part of its meaning: identify the relevant traffic and '
                                 'retain the condition in readback. An unclear condition must not become '
                                 'unconditional permission.\n'
                                 '\n'
                                 '**Constraint case:** an assigned speed is below the applicable minimum for the '
                                 'stated configuration. Report UNABLE, the reason and a proposed achievable '
                                 'value rather than accepting and silently using another speed. If a channel '
                                 'change fails, seek contact under the prescribed procedures while keeping '
                                 'communication separate from clearance. Workload does not turn an assumption '
                                 'into confirmed information.\n'
                                 '\n'
                                 '**Practice format:** 15 chapter MCQs, four choices and one correct answer. '
                                 'Subject 090 has 34 questions in 60 minutes in the cited EASA table. This '
                                 'chapter covers only part of the subject and is not a full paper. EASA '
                                 'framework, including with English or US/British units.',
                   'difficulty': 'advanced',
                   'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — 090',
                   'tags': ['ATPL']}],
     'questions': [{'concept_id': 'ATPL_090_001_C01',
                    'concept_title': 'Identifier le destinataire',
                    'level': 3,
                    'question': "Vous êtes AER452, cap 240. L'ATC dit « AER425, cap 270 ». Quel état retenir "
                                'pour AER452 ?',
                    'question_en': 'You are AER452, heading 240. ATC says “AER425, heading 270”. Which state '
                                   'applies to AER452?',
                    'options': ['Cap 240, instruction adressée à un autre avion',
                                'Cap 270, dernière instruction sur la fréquence',
                                'Cap 270 après collationnement avec AER452',
                                'Cap 255, moyenne en attendant confirmation'],
                    'options_en': ['Heading 240, instruction addressed to another aircraft',
                                   'Heading 270, latest instruction on frequency',
                                   'Heading 270 after reading back as AER452',
                                   'Heading 255, average pending confirmation'],
                    'correct_answer': 'Cap 240, instruction adressée à un autre avion',
                    'explanation': "Les chiffres de l'indicatif sont transposés : le destinataire est différent. "
                                   "Un collationnement fait au nom d'AER452 ne transfère pas la clairance ; "
                                   "aucune moyenne n'est autorisée.",
                    'explanation_en': 'The callsign digits are transposed: this is another addressee. Reading '
                                      'back as AER452 does not transfer clearance; no averaged heading was '
                                      'assigned.',
                    'use_exercise_svg': False,
                    'tags': ['ATPL', 'exam-mcq', 'exam-subject:090'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-standardised-european?erules-id=ERULES-1963177438-9888',
                                          'https://www.ecologie.gouv.fr/sites/default/files/documents/sera_complet_version17012026.pdf'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 090'},
                   {'concept_id': 'ATPL_090_001_C01',
                    'concept_title': 'Identifier le destinataire',
                    'level': 3,
                    'question': 'Message : AER452, cap 260, descendez [masqué], QNH 1006. Quelle information '
                                "doit être demandée avant d'utiliser le nouveau niveau ?",
                    'question_en': 'Message: AER452, heading 260, descend [blocked], QNH 1006. Which information '
                                   'must be requested before using the new level?',
                    'options': ['Confirmation de 3 000 ft parce que prévu au briefing',
                                'Confirmation du QNH déjà reçu uniquement',
                                'Le niveau du trafic précédent comme substitut',
                                'La valeur du niveau de descente'],
                    'options_en': ['Confirmation of 3,000 ft because it was briefed',
                                   'Only confirmation of the already received QNH',
                                   "The preceding aircraft's level as a substitute",
                                   'The descent-level value'],
                    'correct_answer': 'La valeur du niveau de descente',
                    'explanation': "Le niveau n'a pas été reçu. Le QNH et le cap ne permettent pas de le déduire "
                                   '; une attente de briefing ne fournit pas de clairance. Demander la '
                                   "répétition de l'élément manquant.",
                    'explanation_en': 'The level was not received. QNH and heading cannot determine it; a '
                                      'briefing expectation supplies no clearance. Request repetition of the '
                                      'missing item.',
                    'use_exercise_svg': False,
                    'tags': ['ATPL', 'exam-mcq', 'exam-subject:090'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-standardised-european?erules-id=ERULES-1963177438-9888',
                                          'https://www.ecologie.gouv.fr/sites/default/files/documents/sera_complet_version17012026.pdf'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 090'},
                   {'concept_id': 'ATPL_090_001_C01',
                    'concept_title': 'Identifier le destinataire',
                    'level': 3,
                    'question': 'AER452 demande une montée ; réponse STANDBY. Quelle mise à jour du niveau '
                                'autorisé inscrire ?',
                    'question_en': 'AER452 requests a climb; reply is STANDBY. What update belongs in the '
                                   'cleared-level record?',
                    'options': ['Inscrire le niveau demandé après trente secondes',
                                'Considérer toute future montée comme refusée',
                                'Aucune, conserver le niveau précédemment autorisé',
                                'Remplacer par le niveau demandé'],
                    'options_en': ['Enter the requested level after thirty seconds',
                                   'Treat any future climb as refused',
                                   'None, retain the previously cleared level',
                                   'Replace it with the requested level'],
                    'correct_answer': 'Aucune, conserver le niveau précédemment autorisé',
                    'explanation': 'STANDBY laisse la demande en attente, sans nouveau niveau ni délai '
                                   "automatique d'acceptation. Ce n'est pas non plus un refus définitif.",
                    'explanation_en': 'STANDBY leaves the request pending, without a new level or automatic '
                                      'acceptance delay. Nor is it a final denial.',
                    'use_exercise_svg': False,
                    'tags': ['ATPL', 'exam-mcq', 'exam-subject:090'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-standardised-european?erules-id=ERULES-1963177438-9888',
                                          'https://www.ecologie.gouv.fr/sites/default/files/documents/sera_complet_version17012026.pdf'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 090'},
                   {'concept_id': 'ATPL_090_001_C02',
                    'concept_title': 'Collationner les éléments de sécurité',
                    'level': 3,
                    'question': 'Instruction : cap 230, maintenez 5 000 ft, QNH 1008, AER452. Quel '
                                'collationnement est complet et fidèle ?',
                    'question_en': 'Instruction: heading 230, maintain 5,000 ft, QNH 1008, AER452. Which '
                                   'readback is complete and accurate?',
                    'options': ['Cap 230, QNH 1008, AER452',
                                'Cap 230, maintenons 5 000 ft, QNH 1008, AER452',
                                'Cap 230, maintenons 5 000 ft, AER452',
                                'Cap 230, maintenons 5 000 ft, QNH 1018, AER452'],
                    'options_en': ['Heading 230, QNH 1008, AER452',
                                   'Heading 230, maintaining 5,000 ft, QNH 1008, AER452',
                                   'Heading 230, maintaining 5,000 ft, AER452',
                                   'Heading 230, maintaining 5,000 ft, QNH 1018, AER452'],
                    'correct_answer': 'Cap 230, maintenons 5 000 ft, QNH 1008, AER452',
                    'explanation': 'Comparer présence et valeur de chaque élément. Deux réponses omettent calage '
                                   'ou niveau ; une autre remplace 1008 par 1018. La bonne identification seule '
                                   'ne rend pas ces réponses complètes.',
                    'explanation_en': 'Check both presence and value of every item. Two replies omit setting or '
                                      'level; another substitutes 1018 for 1008. Correct identification alone '
                                      'does not make them complete.',
                    'use_exercise_svg': False,
                    'tags': ['ATPL', 'exam-mcq', 'exam-subject:090'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-standardised-european?erules-id=ERULES-1963177438-9888',
                                          'https://www.ecologie.gouv.fr/sites/default/files/documents/sera_complet_version17012026.pdf'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 090'},
                   {'concept_id': 'ATPL_090_001_C02',
                    'concept_title': 'Collationner les éléments de sécurité',
                    'level': 3,
                    'question': 'Code assigné 4621. Le pilote sélectionne 4621 mais annonce 4612. Quel écart '
                                'reste à traiter ?',
                    'question_en': 'Assigned squawk 4621. The pilot selects 4621 but reads back 4612. What '
                                   'discrepancy remains?',
                    'options': ['Le code collationné doit être corrigé à 4621',
                                'Aucun puisque le code sélectionné est correct',
                                "Changer la sélection à 4612 pour l'accorder au message",
                                'Changer de canal pour vérifier le code'],
                    'options_en': ['The read-back code must be corrected to 4621',
                                   'None because the selected code is correct',
                                   'Select 4612 to match the reply',
                                   'Change channel to check the code'],
                    'correct_answer': 'Le code collationné doit être corrigé à 4621',
                    'explanation': "L'action correcte dans le cockpit ne rend pas fidèle le collationnement "
                                   "entendu par l'ATC. Corriger la parole selon l'instruction, plutôt que "
                                   "modifier une sélection juste pour suivre l'erreur verbale.",
                    'explanation_en': 'Correct cockpit selection does not make the readback heard by ATC '
                                      'accurate. Correct the spoken reply to the instruction rather than '
                                      'changing a correct selection to match the verbal error.',
                    'use_exercise_svg': False,
                    'tags': ['ATPL', 'exam-mcq', 'exam-subject:090'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-standardised-european?erules-id=ERULES-1963177438-9888',
                                          'https://www.ecologie.gouv.fr/sites/default/files/documents/sera_complet_version17012026.pdf'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 090'},
                   {'concept_id': 'ATPL_090_001_C02',
                    'concept_title': 'Collationner les éléments de sécurité',
                    'level': 3,
                    'question': 'Nouveau canal 128.305, collationnement 128.350. Que faut-il résoudre avant '
                                'transfert ?',
                    'question_en': 'New channel 128.305, readback 128.350. What must be resolved before '
                                   'transfer?',
                    'options': ['Seulement le nom de la station',
                                'Aucun écart puisque les trois premiers chiffres coïncident',
                                'Le canal précédent doit remplacer les deux valeurs',
                                'La transposition 305/350'],
                    'options_en': ['Only the station name',
                                   'No discrepancy because the first three digits match',
                                   'The previous channel must replace both values',
                                   'The 305/350 transposition'],
                    'correct_answer': 'La transposition 305/350',
                    'explanation': "Comparer l'intégralité du canal : une terminaison différente mène à une "
                                   'autre sélection. Le nom de station ou les chiffres initiaux ne valident pas '
                                   'la valeur reçue.',
                    'explanation_en': 'Compare the complete channel: a different ending changes selection. '
                                      'Station name or matching initial digits cannot validate the received '
                                      'value.',
                    'use_exercise_svg': False,
                    'tags': ['ATPL', 'exam-mcq', 'exam-subject:090'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-standardised-european?erules-id=ERULES-1963177438-9888',
                                          'https://www.ecologie.gouv.fr/sites/default/files/documents/sera_complet_version17012026.pdf'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 090'},
                   {'concept_id': 'ATPL_090_001_C02',
                    'concept_title': 'Collationner les éléments de sécurité',
                    'level': 3,
                    'question': 'Instruction : roulez via A, maintenez avant piste 18. Réponse : via A, '
                                'traversons 18. Quelle erreur modifie la portée ?',
                    'question_en': 'Instruction: taxi via A, hold short runway 18. Reply: via A, crossing 18. '
                                   'Which error changes scope?',
                    'options': ['La piste 18 est devenue piste 36',
                                "Le seul défaut est l'absence de vitesse",
                                'La traversée remplace le maintien avant piste',
                                'La voie A est devenue inconnue'],
                    'options_en': ['Runway 18 became runway 36',
                                   'The only defect is missing speed',
                                   'Crossing replaces holding short',
                                   'Taxiway A has become unknown'],
                    'correct_answer': 'La traversée remplace le maintien avant piste',
                    'explanation': 'La route de roulage et le numéro sont conservés, mais le verbe autorise à '
                                   'tort une autre action. La correction doit rétablir maintien avant piste, pas '
                                   'simplement répéter son numéro.',
                    'explanation_en': 'Taxi route and runway number match, but the verb wrongly substitutes '
                                      'another action. Correction must restore hold short, not merely repeat the '
                                      'runway number.',
                    'use_exercise_svg': False,
                    'tags': ['ATPL', 'exam-mcq', 'exam-subject:090'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-standardised-european?erules-id=ERULES-1963177438-9888',
                                          'https://www.ecologie.gouv.fr/sites/default/files/documents/sera_complet_version17012026.pdf'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 090'},
                   {'concept_id': 'ATPL_090_001_C03',
                    'concept_title': 'Fermer la boucle',
                    'level': 3,
                    'question': 'AER452 : cap 240, 5 000 ft, 180 kt. Puis « cap 260, correction, cap 250 ». '
                                'Aucun autre changement. État autorisé ?',
                    'question_en': 'AER452: heading 240, 5,000 ft, 180 kt. Then “heading 260, correction, '
                                   'heading 250”. No other change. Cleared state?',
                    'options': ['240 / 5 000 ft / 180 kt',
                                '250 / 5 000 ft / 180 kt',
                                '260 / 5 000 ft / 180 kt',
                                '250 / 2 500 ft / 180 kt'],
                    'options_en': ['240 / 5,000 ft / 180 kt',
                                   '250 / 5,000 ft / 180 kt',
                                   '260 / 5,000 ft / 180 kt',
                                   '250 / 2,500 ft / 180 kt'],
                    'correct_answer': '250 / 5 000 ft / 180 kt',
                    'explanation': 'La correction remplace 260 par 250 ; elle ne transforme pas les chiffres du '
                                   "cap en niveau. 240 est l'ancien cap, désormais modifié. Niveau et vitesse "
                                   'restent inchangés.',
                    'explanation_en': 'Correction replaces 260 with 250; heading digits do not become a level. '
                                      '240 is the previous, amended heading. Level and speed remain unchanged.',
                    'use_exercise_svg': False,
                    'tags': ['ATPL', 'exam-mcq', 'exam-subject:090'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-standardised-european?erules-id=ERULES-1963177438-9888',
                                          'https://www.ecologie.gouv.fr/sites/default/files/documents/sera_complet_version17012026.pdf'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 090'},
                   {'concept_id': 'ATPL_090_001_C03',
                    'concept_title': 'Fermer la boucle',
                    'level': 3,
                    'question': 'Instruction : cap 210, QNH 1013, code 5274. Réponse : cap 210, QNH 1031, '
                                'AER452. Quels défauts détecter ?',
                    'question_en': 'Instruction: heading 210, QNH 1013, squawk 5274. Reply: heading 210, QNH '
                                   '1031, AER452. Which defects are present?',
                    'options': ['QNH substitué et code omis',
                                'Cap erroné et QNH correct',
                                'Code erroné mais présent',
                                'QNH seul, puisque le code peut être omis'],
                    'options_en': ['Substituted QNH and omitted code',
                                   'Wrong heading and correct QNH',
                                   'Wrong but present code',
                                   'Only QNH, because code can be omitted'],
                    'correct_answer': 'QNH substitué et code omis',
                    'explanation': "1013 est devenu 1031 et 5274 n'a pas été répété. Détecter une erreur ne doit "
                                   "pas arrêter la recherche d'une seconde ; omission et substitution demandent "
                                   'toutes deux une correction.',
                    'explanation_en': '1013 became 1031 and 5274 was not repeated. Finding one error must not '
                                      'stop the search for another; omission and substitution both need repair.',
                    'use_exercise_svg': False,
                    'tags': ['ATPL', 'exam-mcq', 'exam-subject:090'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-standardised-european?erules-id=ERULES-1963177438-9888',
                                          'https://www.ecologie.gouv.fr/sites/default/files/documents/sera_complet_version17012026.pdf'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 090'},
                   {'concept_id': 'ATPL_090_001_C03',
                    'concept_title': 'Fermer la boucle',
                    'level': 3,
                    'question': 'Exercice : vitesse minimale de configuration 160 kt, ATC demande 150 kt. Quel '
                                'message expose correctement la contrainte ?',
                    'question_en': 'Exercise: configuration minimum speed 160 kt, ATC requests 150 kt. Which '
                                   'message correctly exposes the constraint?',
                    'options': ['WILCO 150, puis maintenir 160 sans annonce',
                                'ROGER 150, puis attendre une détection ATC',
                                'Collationner 160 comme si ce chiffre avait été assigné',
                                'UNABLE 150, minimum 160 dans cette configuration, demandons 160'],
                    'options_en': ['WILCO 150, then maintain 160 silently',
                                   'ROGER 150, then wait for ATC detection',
                                   'Read back 160 as if it had been assigned',
                                   'UNABLE 150, minimum 160 in this configuration, request 160'],
                    'correct_answer': 'UNABLE 150, minimum 160 dans cette configuration, demandons 160',
                    'explanation': "Annoncer impossibilité, raison et proposition distingue une demande d'une "
                                   'clairance. Les autres réponses masquent la non-exécution ou transforment un '
                                   'chiffre proposé en instruction reçue.',
                    'explanation_en': 'Stating inability, reason and proposal distinguishes a request from '
                                      'clearance. Other replies conceal non-compliance or turn a proposed number '
                                      'into a received instruction.',
                    'use_exercise_svg': False,
                    'tags': ['ATPL', 'exam-mcq', 'exam-subject:090'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-standardised-european?erules-id=ERULES-1963177438-9888',
                                          'https://www.ecologie.gouv.fr/sites/default/files/documents/sera_complet_version17012026.pdf'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 090'},
                   {'concept_id': 'ATPL_090_001_C03',
                    'concept_title': 'Fermer la boucle',
                    'level': 3,
                    'question': "Votre collationnement était 4 000 ft ; l'ATC répond « négatif, répétez [valeur "
                                'masquée] ». Quelle information pouvez-vous considérer comme résolue ?',
                    'question_en': 'Your readback was 4,000 ft; ATC replies “negative, say again [blocked '
                                   'value]”. What can be considered resolved?',
                    'options': ['Le niveau prévu sur la fiche de briefing',
                                'La différence exacte de 1 000 ft',
                                "La présence d'un écart, mais pas le niveau corrigé",
                                "Le niveau 4 000 ft, faute d'autre valeur entendue"],
                    'options_en': ['The briefing-sheet level',
                                   'An exact 1,000 ft difference',
                                   'There is a discrepancy, but the corrected level is unresolved',
                                   '4,000 ft because no other value was heard'],
                    'correct_answer': "La présence d'un écart, mais pas le niveau corrigé",
                    'explanation': "Le rejet du collationnement est reçu, sa valeur de remplacement ne l'est "
                                   'pas. Redemander la correction ; ni la mémoire ni une différence supposée ne '
                                   'comblent le message manquant.',
                    'explanation_en': 'Rejection of the readback was heard, its replacement value was not. '
                                      'Request the correction again; memory and an assumed difference cannot '
                                      'fill the missing message.',
                    'use_exercise_svg': False,
                    'tags': ['ATPL', 'exam-mcq', 'exam-subject:090'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-standardised-european?erules-id=ERULES-1963177438-9888',
                                          'https://www.ecologie.gouv.fr/sites/default/files/documents/sera_complet_version17012026.pdf'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 090'},
                   {'concept_id': 'ATPL_090_001_C04',
                    'concept_title': 'Distinguer les autorisations',
                    'level': 3,
                    'question': "Au point d'attente, autorisation « alignez-vous piste 32 et attendez ». Quelle "
                                'limite doit rester dans la réponse ?',
                    'question_en': 'At the holding point, clearance “line up runway 32 and wait”. Which limit '
                                   'must remain in the reply?',
                    'options': ['Traverser 32 vers la voie opposée',
                                'Attendre après alignement ; pas de décollage autorisé',
                                "Maintenir avant la piste malgré l'alignement demandé",
                                'Décoller après le précédent trafic sans autre instruction'],
                    'options_en': ['Cross 32 to the opposite taxiway',
                                   'Wait after line-up; no take-off clearance',
                                   'Hold short despite the line-up instruction',
                                   'Take off after preceding traffic without further instruction'],
                    'correct_answer': 'Attendre après alignement ; pas de décollage autorisé',
                    'explanation': 'Alignement, traversée et décollage sont différents. La limite « attendez » '
                                   "conserve le point d'arrêt de cette autorisation ; elle ne définit aucune "
                                   'condition implicite de départ.',
                    'explanation_en': 'Line-up, crossing and take-off differ. “Wait” preserves the stopping '
                                      'point of this clearance; it creates no implicit departure condition.',
                    'use_exercise_svg': False,
                    'tags': ['ATPL', 'exam-mcq', 'exam-subject:090'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-standardised-european?erules-id=ERULES-1963177438-9888',
                                          'https://www.ecologie.gouv.fr/sites/default/files/documents/sera_complet_version17012026.pdf'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 090'},
                   {'concept_id': 'ATPL_090_001_C04',
                    'concept_title': 'Distinguer les autorisations',
                    'level': 3,
                    'question': "Clairance d'approche reçue ; plus tard « continuez l'approche ». Aucune "
                                "autorisation d'atterrir. Quel état inscrire ?",
                    'question_en': 'Approach clearance received; later “continue approach”. No landing '
                                   'clearance. Which state applies?',
                    'options': ['Approche autorisée, atterrissage encore non autorisé',
                                'Atterrissage autorisé par la deuxième transmission',
                                "Approche annulée en l'absence de clairance d'atterrir",
                                'Atterrissage autorisé si la piste est visuellement libre'],
                    'options_en': ['Approach cleared, landing not yet cleared',
                                   'Landing cleared by the second transmission',
                                   'Approach cancelled without landing clearance',
                                   'Landing cleared if runway visually clear'],
                    'correct_answer': 'Approche autorisée, atterrissage encore non autorisé',
                    'explanation': "« Continuez » ne contient pas l'autorisation distincte d'atterrir. "
                                   "L'exercice demande la portée des messages, pas une décision de poursuite à "
                                   'une hauteur donnée ; cette dernière dépend de la situation réelle.',
                    'explanation_en': '“Continue” does not contain separate landing clearance. This asks message '
                                      'scope, not whether to continue at a particular height; that decision '
                                      'depends on the actual situation.',
                    'use_exercise_svg': False,
                    'tags': ['ATPL', 'exam-mcq', 'exam-subject:090'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-standardised-european?erules-id=ERULES-1963177438-9888',
                                          'https://www.ecologie.gouv.fr/sites/default/files/documents/sera_complet_version17012026.pdf'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 090'},
                   {'concept_id': 'ATPL_090_001_C02',
                    'concept_title': 'Collationner les éléments de sécurité',
                    'level': 3,
                    'question': 'Clairance conditionnelle correctement émise : derrière le trafic identifié en '
                                'courte finale, alignez-vous 24 derrière. Quel collationnement préserve la '
                                'condition ?',
                    'question_en': 'Properly issued conditional clearance: behind identified traffic on short '
                                   'final, line up 24 behind. Which readback preserves the condition?',
                    'options': ['Alignons 24, AER452',
                                'Après notre contrôle visuel, décollons 24, AER452',
                                'Maintenons avant 24 puis traversons derrière, AER452',
                                'Derrière ce trafic, alignons 24 derrière, AER452'],
                    'options_en': ['Lining up 24, AER452',
                                   'After our visual check, taking off 24, AER452',
                                   'Holding short 24 then crossing behind, AER452',
                                   'Behind that traffic, lining up 24 behind, AER452'],
                    'correct_answer': 'Derrière ce trafic, alignons 24 derrière, AER452',
                    'explanation': "La condition est une partie de la clairance, pas un commentaire. L'omettre "
                                   "transforme le sens ; décollage et traversée remplacent l'action autorisée. "
                                   'Le trafic doit être identifié sans ambiguïté.',
                    'explanation_en': 'The condition is part of clearance, not commentary. Omitting it changes '
                                      'meaning; take-off and crossing replace the cleared action. Traffic must '
                                      'be unambiguously identified.',
                    'use_exercise_svg': False,
                    'tags': ['ATPL', 'exam-mcq', 'exam-subject:090'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-standardised-european?erules-id=ERULES-1963177438-9888',
                                          'https://www.ecologie.gouv.fr/sites/default/files/documents/sera_complet_version17012026.pdf'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 090'},
                   {'concept_id': 'ATPL_090_001_C04',
                    'concept_title': 'Distinguer les autorisations',
                    'level': 3,
                    'question': 'Dernier état AER452 : cap 220, 4 000 ft, 170 kt. Message suivant pour AER425 : '
                                '160 kt. Puis AER452 : cap 230, maintenez 3 000 ft. Quel état final ?',
                    'question_en': 'Last AER452 state: heading 220, 4,000 ft, 170 kt. Next, AER425 receives 160 '
                                   'kt. Then AER452 receives heading 230, maintain 3,000 ft. Final state?',
                    'options': ['220 / 3 000 ft / 170 kt',
                                '230 / 4 000 ft / 170 kt',
                                '230 / 3 000 ft / 170 kt',
                                '230 / 3 000 ft / 160 kt'],
                    'options_en': ['220 / 3,000 ft / 170 kt',
                                   '230 / 4,000 ft / 170 kt',
                                   '230 / 3,000 ft / 170 kt',
                                   '230 / 3,000 ft / 160 kt'],
                    'correct_answer': '230 / 3 000 ft / 170 kt',
                    'explanation': 'Mettre à jour les deux éléments adressés à AER452 : cap et niveau. La '
                                   'vitesse 160 appartient à AER425 ; conserver 170. Les deux autres '
                                   'distracteurs oublient une modification correctement reçue.',
                    'explanation_en': "Update AER452's two amended items: heading and level. 160 belongs to "
                                      'AER425; retain 170. The other two distractors each miss a correctly '
                                      'received amendment.',
                    'use_exercise_svg': True,
                    'tags': ['ATPL', 'exam-mcq', 'exam-subject:090'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-standardised-european?erules-id=ERULES-1963177438-9888',
                                          'https://www.ecologie.gouv.fr/sites/default/files/documents/sera_complet_version17012026.pdf'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 090'}]}
    original = data["metadata"]["module_id"]
    module = program + "_090_001"
    data["metadata"].update(program=program, module_id=module)
    for c in data["concepts"]:
        c["concept_id"] = c["concept_id"].replace(original, module)
    for q in data["questions"]:
        q["concept_id"] = q["concept_id"].replace(original, module)

    save(data)
