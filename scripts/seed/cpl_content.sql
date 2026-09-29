-- CPL content only. Preserves accounts, progress and answer history.
BEGIN;
SET LOCAL standard_conforming_strings = on;

INSERT INTO lessons (id, license, category, theme, title_fr, title_en, content_fr, content_en, difficulty, order_index, level, duration_minutes, tags, learning_objectives) VALUES
('26ad7496-0000-4000-8000-000026ad7496', 'CPL', 'communications', 'communications', 'Clairances et boucle de collationnement', 'Clearances and the readback loop', '**Objectif :** Reconstruire une clairance après messages successifs, détecter omissions et substitutions et préserver conditions et limites au collationnement.

## Identifier le destinataire

![CPL_090_001](/illustrations/cpl-090-001.svg)

Avant une émission, écouter pour éviter de couper une transmission. Vérifier que la clairance concerne son indicatif, particulièrement lorsque deux indicatifs se ressemblent. Une transmission partiellement entendue ne doit pas être reconstruite à partir de ce que l’on attendait : demander la répétition du passage manquant. Employer la phraséologie publiée et une parole claire ; ce chapitre travaille la compréhension, sans remplacer un entraînement oral.

### Tenir un état des instructions
Préparer une grille mentale ou écrite : destinataire, route/cap, niveau, vitesse, calage, code et canal. Une instruction ultérieure remplace les éléments qu''elle modifie ; ne pas recopier tous les éléments d''un message adressé à un autre avion. Une donnée masquée reste inconnue même si elle paraît prévisible.

Exemple fictif : AER321 a reçu cap 250, 4 000 ft, 170 kt. AER312 reçoit ensuite cap 280. Le cap d''AER321 reste 250 : la dernière valeur entendue sur la fréquence n''est pas nécessairement sa dernière instruction. Si seul le destinataire est masqué, faire lever le doute avant d''attribuer la nouvelle instruction. L''abréviation d''un indicatif d''immatriculation peut être initiée par le sol dans les conditions SERA ; un indicatif exploitant + numéro de vol comme AER321 ne s''abrège pas.

## Collationner les éléments de sécurité

SERA.8015(e) impose le collationnement des clairances de route et des autorisations ou instructions relatives aux pistes, notamment entrée, traversée et maintien avant piste. Sont aussi concernés piste en service, calage altimétrique, code SSR, nouveau canal de communication, niveau, cap, vitesse et niveau de transition. Un simple accusé de réception ne remplace pas ces éléments. Inclure son indicatif selon la procédure de radiotéléphonie ; le contrôleur écoute et corrige les écarts.

### Détecter omission et substitution
Exemple écrit : « AER321, cap 250, maintenez 4 000 ft, QNH 1014, code 4251 ». Réponse analysée : « Cap 250, maintenons 4 000 ft, QNH 1041, AER321 ». Deux défauts distincts : QNH substitué et code omis. Un contrôle qui ne vérifie que le nombre d''éléments peut manquer la substitution ; un contrôle qui ne compare que les valeurs présentes peut manquer l''omission. La boucle doit permettre de corriger les deux.

ROGER accuse réception sans répéter les valeurs ; WILCO n''est pas un substitut au collationnement requis. STANDBY n''accorde aucune demande en attente. Les nombres écrits dans les exercices ne dispensent pas de la prononciation réglementaire en entraînement oral.

## Fermer la boucle

Exemple pédagogique, hors phraséologie complète : AER123 reçoit cap 270, maintien 3 000 ft, QNH 1016. Son collationnement indique par erreur 2 000 ft. Le contrôleur doit corriger cet écart ; le pilote doit comprendre et collationner la correction. Une absence de correction entendue ne justifie pas de conserver un doute. En cas de clairance impossible à respecter, informer l’ATC et demander une solution adaptée, sans accepter silencieusement.

### Corriger sans perdre les éléments inchangés
Reprendre le cas AER321 : cap 250, 4 000 ft, 170 kt. Puis « cap 260, correction, cap 270 ». Le cap corrigé est 270 ; niveau et vitesse ne sont pas modifiés par cette seule correction. Un collationnement « cap 260 » conserve une version explicitement corrigée. Une correction du contrôleur doit être comprise et répétée ; s''il manque la valeur corrigée, demander sa répétition.

**Travail sans les réponses :** noter les trois paramètres, barrer uniquement ceux qui changent, puis relire l''état obtenu. Ensuite analyser le dialogue dans les QCM avec de nouveaux indicatifs et chiffres. Ne pas confondre l''état autorisé avec l''état réel de l''avion : le premier décrit l''instruction, le second exige surveillance et exécution.

## Distinguer les autorisations

Une autorisation de roulage vers un point d’attente ne constitue pas une autorisation de traverser la piste. Une clairance d’approche ne constitue pas une autorisation d’atterrir. Préparer le changement de fréquence, vérifier le canal assigné et établir le contact requis. Si le contact échoue, appliquer les procédures de communication et de panne pertinentes ; ne pas inventer une nouvelle clairance. Le pilote reste responsable de la conduite sûre du vol.

### Identifier une limite, une condition et une impossibilité
« Roulez point d''attente piste 18 » ne décrit pas une traversée ; « alignez-vous et attendez » n''autorise pas le décollage. Lorsqu''une clairance conditionnelle est utilisée selon les règles applicables, sa condition fait partie de son sens : identifier le trafic concerné et conserver cette condition au collationnement. Une condition non comprise ne doit pas devenir une autorisation inconditionnelle.

**Cas de contrainte :** une vitesse demandée est sous la vitesse minimale applicable à la configuration donnée. Informer l''ATC avec UNABLE, la raison et une valeur réalisable proposée ; ne pas accepter puis improviser une autre vitesse en silence. Si un changement de canal échoue, rechercher le contact selon les procédures prévues, en conservant la distinction entre communication et clairance. La charge de travail ne justifie pas de transformer une supposition en donnée confirmée.

**Format de travail :** 15 QCM de chapitre, quatre choix et une réponse correcte. La matière 090 comporte 34 questions en 60 minutes dans le tableau EASA cité. Ce chapitre ne couvre qu’une partie de la matière ; il ne constitue pas une épreuve complète. Cadre EASA, même en anglais ou avec les unités US/britanniques.

## Références

- [EASA — programme Part-FCL](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8)
- [EASA — format de la matière, AMC1 ARA.FCL.300(b)](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67)
- [EASA SERA.8015 — clairances et collationnement](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-standardised-european?erules-id=ERULES-1963177438-9888)
- [Référence pédagogique vérifiée](https://www.ecologie.gouv.fr/sites/default/files/documents/sera_complet_version17012026.pdf)

**Support de révision CPL(A).** À faire valider par votre instructeur. Pour un vol réel, utiliser le manuel de l’avion et les publications en vigueur. Les exemples chiffrés sont pédagogiques.', '**Objective:** Reconstruct clearance after successive messages, detect omissions and substitutions and preserve conditions and limits in readback.

## Identify the addressee

![CPL_090_001](/illustrations/en/cpl-090-001.svg)

Listen before transmitting to avoid interrupting another transmission. Check that the clearance addresses your callsign, particularly when similar callsigns share the frequency. Do not reconstruct a partly heard transmission from what you expected: request repetition of the missing part. Use published phraseology and clear speech; this chapter practises comprehension and does not replace oral training.

### Maintain an instruction state
Use a mental or written grid: addressee, route/heading, level, speed, setting, code and channel. A later instruction replaces the items it amends; do not copy items from another aircraft''s message. A blocked item remains unknown even when predictable.

Fictional example: AER321 has heading 250, 4,000 ft and 170 kt. AER312 then receives heading 280. AER321 keeps 250: the latest number heard is not necessarily its latest instruction. If only the addressee is blocked, resolve that uncertainty before assigning the amendment. Ground may initiate a registration abbreviation under SERA conditions; an operator-plus-flight-number callsign such as AER321 is not abbreviated.

## Read back safety-related elements

SERA.8015(e) requires readback of route clearances and runway clearances or instructions, including entry, crossing and holding short. The runway in use, altimeter setting, SSR code, newly assigned communication channel, level, heading, speed and transition level are also included. A simple acknowledgement does not replace these elements. Include your callsign in accordance with radiotelephony procedure; the controller listens and corrects discrepancies.

### Detect omission and substitution
Written example: “AER321, heading 250, maintain 4,000 ft, QNH 1014, squawk 4251”. Analysed reply: “Heading 250, maintaining 4,000 ft, QNH 1041, AER321”. Two defects: substituted QNH and omitted code. Counting items alone may miss substitution; comparing only stated values may miss omission. The loop must repair both.

ROGER acknowledges receipt without repeating values; WILCO does not substitute for required readback. STANDBY grants no pending request. Printed numbers do not replace prescribed pronunciation in oral practice.

## Close the loop

Training example, not complete phraseology: AER123 receives heading 270, maintain 3,000 ft, QNH 1016. Its readback incorrectly states 2,000 ft. The controller must correct the discrepancy; the pilot must understand and read back the correction. Not hearing a correction does not resolve an existing doubt. If unable to comply with a clearance, inform ATC and request a suitable alternative rather than silently accepting.

### Correct without losing unchanged items
Return to AER321: heading 250, 4,000 ft, 170 kt. Then “heading 260, correction, heading 270”. Corrected heading is 270; this correction alone does not change level or speed. Reading back 260 retains a version explicitly corrected. An ATC correction must be understood and read back; request repetition if its value is missing.

**Work before seeing answers:** record the three parameters, cross out only amended ones and read the resulting state. Then analyse MCQ exchanges with new callsigns and numbers. Distinguish the cleared state from actual aircraft state: the first describes instructions, the second needs monitoring and execution.

## Distinguish authorisations

Taxi clearance to a holding point is not clearance to cross the runway. Approach clearance is not landing clearance. Prepare a frequency change, verify the assigned channel and establish the required contact. If contact fails, apply the relevant communication and failure procedures; do not invent a new clearance. The pilot remains responsible for safe conduct of the flight.

### Identify a limit, condition and inability
“Taxi holding point runway 18” does not clear a crossing; “line up and wait” does not clear take-off. Where conditional clearance is used under applicable rules, the condition is part of its meaning: identify the relevant traffic and retain the condition in readback. An unclear condition must not become unconditional permission.

**Constraint case:** an assigned speed is below the applicable minimum for the stated configuration. Report UNABLE, the reason and a proposed achievable value rather than accepting and silently using another speed. If a channel change fails, seek contact under the prescribed procedures while keeping communication separate from clearance. Workload does not turn an assumption into confirmed information.

**Practice format:** 15 chapter MCQs, four choices and one correct answer. Subject 090 has 34 questions in 60 minutes in the cited EASA table. This chapter covers only part of the subject and is not a full paper. EASA framework, including with English or US/British units.

## References

- [Reference 1](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8)
- [Reference 2](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67)
- [Reference 3](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-standardised-european?erules-id=ERULES-1963177438-9888)
- [Reference 4](https://www.ecologie.gouv.fr/sites/default/files/documents/sera_complet_version17012026.pdf)

**CPL(A) revision material.** To be reviewed by your instructor. For an actual flight, use the aircraft manual and current publications. Numerical examples are for training.', 3, 1, 1, 45, '["CPL", "mcq-four-options"]', '["Reconstruire une clairance après messages successifs, détecter omissions et substitutions et préserver conditions et limites au collationnement."]')
ON CONFLICT (id) DO UPDATE SET license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, title_fr = EXCLUDED.title_fr, title_en = EXCLUDED.title_en, content_fr = EXCLUDED.content_fr, content_en = EXCLUDED.content_en, difficulty = EXCLUDED.difficulty, order_index = EXCLUDED.order_index, level = EXCLUDED.level, duration_minutes = EXCLUDED.duration_minutes, tags = EXCLUDED.tags, learning_objectives = EXCLUDED.learning_objectives;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('35676a31-0000-4000-8000-000035676a31', '26ad7496-0000-4000-8000-000026ad7496', 'CPL', 'communications', 'communications', 'Identifier le destinataire', 3, 3, 'Vous êtes AER452, cap 240. L''ATC dit « AER425, cap 270 ». Quel état retenir pour AER452 ?', 'You are AER452, heading 240. ATC says “AER425, heading 270”. Which state applies to AER452?', '["Cap 240, instruction adressée à un autre avion", "Cap 270, dernière instruction sur la fréquence", "Cap 270 après collationnement avec AER452", "Cap 255, moyenne en attendant confirmation"]', '["Heading 240, instruction addressed to another aircraft", "Heading 270, latest instruction on frequency", "Heading 270 after reading back as AER452", "Heading 255, average pending confirmation"]', 'Cap 240, instruction adressée à un autre avion', 'Les chiffres de l''indicatif sont transposés : le destinataire est différent. Un collationnement fait au nom d''AER452 ne transfère pas la clairance ; aucune moyenne n''est autorisée.

![CPL_090_001](/illustrations/cpl-090-001.svg)', 'The callsign digits are transposed: this is another addressee. Reading back as AER452 does not transfer clearance; no averaged heading was assigned.

![CPL_090_001](/illustrations/en/cpl-090-001.svg)', '["CPL", "exam-mcq", "exam-subject:090"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('35676a32-0000-4000-8000-000035676a32', '26ad7496-0000-4000-8000-000026ad7496', 'CPL', 'communications', 'communications', 'Identifier le destinataire', 3, 3, 'Message : AER452, cap 260, descendez [masqué], QNH 1006. Quelle information doit être demandée avant d''utiliser le nouveau niveau ?', 'Message: AER452, heading 260, descend [blocked], QNH 1006. Which information must be requested before using the new level?', '["Confirmation de 3 000 ft parce que prévu au briefing", "Confirmation du QNH déjà reçu uniquement", "Le niveau du trafic précédent comme substitut", "La valeur du niveau de descente"]', '["Confirmation of 3,000 ft because it was briefed", "Only confirmation of the already received QNH", "The preceding aircraft''s level as a substitute", "The descent-level value"]', 'La valeur du niveau de descente', 'Le niveau n''a pas été reçu. Le QNH et le cap ne permettent pas de le déduire ; une attente de briefing ne fournit pas de clairance. Demander la répétition de l''élément manquant.

![CPL_090_001](/illustrations/cpl-090-001.svg)', 'The level was not received. QNH and heading cannot determine it; a briefing expectation supplies no clearance. Request repetition of the missing item.

![CPL_090_001](/illustrations/en/cpl-090-001.svg)', '["CPL", "exam-mcq", "exam-subject:090"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('35676a33-0000-4000-8000-000035676a33', '26ad7496-0000-4000-8000-000026ad7496', 'CPL', 'communications', 'communications', 'Identifier le destinataire', 3, 3, 'AER452 demande une montée ; réponse STANDBY. Quelle mise à jour du niveau autorisé inscrire ?', 'AER452 requests a climb; reply is STANDBY. What update belongs in the cleared-level record?', '["Inscrire le niveau demandé après trente secondes", "Considérer toute future montée comme refusée", "Aucune, conserver le niveau précédemment autorisé", "Remplacer par le niveau demandé"]', '["Enter the requested level after thirty seconds", "Treat any future climb as refused", "None, retain the previously cleared level", "Replace it with the requested level"]', 'Aucune, conserver le niveau précédemment autorisé', 'STANDBY laisse la demande en attente, sans nouveau niveau ni délai automatique d''acceptation. Ce n''est pas non plus un refus définitif.

![CPL_090_001](/illustrations/cpl-090-001.svg)', 'STANDBY leaves the request pending, without a new level or automatic acceptance delay. Nor is it a final denial.

![CPL_090_001](/illustrations/en/cpl-090-001.svg)', '["CPL", "exam-mcq", "exam-subject:090"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('35676df5-0000-4000-8000-000035676df5', '26ad7496-0000-4000-8000-000026ad7496', 'CPL', 'communications', 'communications', 'Collationner les éléments de sécurité', 3, 3, 'Instruction : cap 230, maintenez 5 000 ft, QNH 1008, AER452. Quel collationnement est complet et fidèle ?', 'Instruction: heading 230, maintain 5,000 ft, QNH 1008, AER452. Which readback is complete and accurate?', '["Cap 230, QNH 1008, AER452", "Cap 230, maintenons 5 000 ft, QNH 1008, AER452", "Cap 230, maintenons 5 000 ft, AER452", "Cap 230, maintenons 5 000 ft, QNH 1018, AER452"]', '["Heading 230, QNH 1008, AER452", "Heading 230, maintaining 5,000 ft, QNH 1008, AER452", "Heading 230, maintaining 5,000 ft, AER452", "Heading 230, maintaining 5,000 ft, QNH 1018, AER452"]', 'Cap 230, maintenons 5 000 ft, QNH 1008, AER452', 'Comparer présence et valeur de chaque élément. Deux réponses omettent calage ou niveau ; une autre remplace 1008 par 1018. La bonne identification seule ne rend pas ces réponses complètes.

![CPL_090_001](/illustrations/cpl-090-001.svg)', 'Check both presence and value of every item. Two replies omit setting or level; another substitutes 1018 for 1008. Correct identification alone does not make them complete.

![CPL_090_001](/illustrations/en/cpl-090-001.svg)', '["CPL", "exam-mcq", "exam-subject:090"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('35676df6-0000-4000-8000-000035676df6', '26ad7496-0000-4000-8000-000026ad7496', 'CPL', 'communications', 'communications', 'Collationner les éléments de sécurité', 3, 3, 'Code assigné 4621. Le pilote sélectionne 4621 mais annonce 4612. Quel écart reste à traiter ?', 'Assigned squawk 4621. The pilot selects 4621 but reads back 4612. What discrepancy remains?', '["Le code collationné doit être corrigé à 4621", "Aucun puisque le code sélectionné est correct", "Changer la sélection à 4612 pour l''accorder au message", "Changer de canal pour vérifier le code"]', '["The read-back code must be corrected to 4621", "None because the selected code is correct", "Select 4612 to match the reply", "Change channel to check the code"]', 'Le code collationné doit être corrigé à 4621', 'L''action correcte dans le cockpit ne rend pas fidèle le collationnement entendu par l''ATC. Corriger la parole selon l''instruction, plutôt que modifier une sélection juste pour suivre l''erreur verbale.

![CPL_090_001](/illustrations/cpl-090-001.svg)', 'Correct cockpit selection does not make the readback heard by ATC accurate. Correct the spoken reply to the instruction rather than changing a correct selection to match the verbal error.

![CPL_090_001](/illustrations/en/cpl-090-001.svg)', '["CPL", "exam-mcq", "exam-subject:090"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('35676df7-0000-4000-8000-000035676df7', '26ad7496-0000-4000-8000-000026ad7496', 'CPL', 'communications', 'communications', 'Collationner les éléments de sécurité', 3, 3, 'Nouveau canal 128.305, collationnement 128.350. Que faut-il résoudre avant transfert ?', 'New channel 128.305, readback 128.350. What must be resolved before transfer?', '["Seulement le nom de la station", "Aucun écart puisque les trois premiers chiffres coïncident", "Le canal précédent doit remplacer les deux valeurs", "La transposition 305/350"]', '["Only the station name", "No discrepancy because the first three digits match", "The previous channel must replace both values", "The 305/350 transposition"]', 'La transposition 305/350', 'Comparer l''intégralité du canal : une terminaison différente mène à une autre sélection. Le nom de station ou les chiffres initiaux ne valident pas la valeur reçue.

![CPL_090_001](/illustrations/cpl-090-001.svg)', 'Compare the complete channel: a different ending changes selection. Station name or matching initial digits cannot validate the received value.

![CPL_090_001](/illustrations/en/cpl-090-001.svg)', '["CPL", "exam-mcq", "exam-subject:090"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('35676df8-0000-4000-8000-000035676df8', '26ad7496-0000-4000-8000-000026ad7496', 'CPL', 'communications', 'communications', 'Collationner les éléments de sécurité', 3, 3, 'Instruction : roulez via A, maintenez avant piste 18. Réponse : via A, traversons 18. Quelle erreur modifie la portée ?', 'Instruction: taxi via A, hold short runway 18. Reply: via A, crossing 18. Which error changes scope?', '["La piste 18 est devenue piste 36", "Le seul défaut est l''absence de vitesse", "La traversée remplace le maintien avant piste", "La voie A est devenue inconnue"]', '["Runway 18 became runway 36", "The only defect is missing speed", "Crossing replaces holding short", "Taxiway A has become unknown"]', 'La traversée remplace le maintien avant piste', 'La route de roulage et le numéro sont conservés, mais le verbe autorise à tort une autre action. La correction doit rétablir maintien avant piste, pas simplement répéter son numéro.

![CPL_090_001](/illustrations/cpl-090-001.svg)', 'Taxi route and runway number match, but the verb wrongly substitutes another action. Correction must restore hold short, not merely repeat the runway number.

![CPL_090_001](/illustrations/en/cpl-090-001.svg)', '["CPL", "exam-mcq", "exam-subject:090"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('356771ba-0000-4000-8000-0000356771ba', '26ad7496-0000-4000-8000-000026ad7496', 'CPL', 'communications', 'communications', 'Fermer la boucle', 3, 3, 'AER452 : cap 240, 5 000 ft, 180 kt. Puis « cap 260, correction, cap 250 ». Aucun autre changement. État autorisé ?', 'AER452: heading 240, 5,000 ft, 180 kt. Then “heading 260, correction, heading 250”. No other change. Cleared state?', '["240 / 5 000 ft / 180 kt", "250 / 5 000 ft / 180 kt", "260 / 5 000 ft / 180 kt", "250 / 2 500 ft / 180 kt"]', '["240 / 5,000 ft / 180 kt", "250 / 5,000 ft / 180 kt", "260 / 5,000 ft / 180 kt", "250 / 2,500 ft / 180 kt"]', '250 / 5 000 ft / 180 kt', 'La correction remplace 260 par 250 ; elle ne transforme pas les chiffres du cap en niveau. 240 est l''ancien cap, désormais modifié. Niveau et vitesse restent inchangés.

![CPL_090_001](/illustrations/cpl-090-001.svg)', 'Correction replaces 260 with 250; heading digits do not become a level. 240 is the previous, amended heading. Level and speed remain unchanged.

![CPL_090_001](/illustrations/en/cpl-090-001.svg)', '["CPL", "exam-mcq", "exam-subject:090"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('356771bb-0000-4000-8000-0000356771bb', '26ad7496-0000-4000-8000-000026ad7496', 'CPL', 'communications', 'communications', 'Fermer la boucle', 3, 3, 'Instruction : cap 210, QNH 1013, code 5274. Réponse : cap 210, QNH 1031, AER452. Quels défauts détecter ?', 'Instruction: heading 210, QNH 1013, squawk 5274. Reply: heading 210, QNH 1031, AER452. Which defects are present?', '["QNH substitué et code omis", "Cap erroné et QNH correct", "Code erroné mais présent", "QNH seul, puisque le code peut être omis"]', '["Substituted QNH and omitted code", "Wrong heading and correct QNH", "Wrong but present code", "Only QNH, because code can be omitted"]', 'QNH substitué et code omis', '1013 est devenu 1031 et 5274 n''a pas été répété. Détecter une erreur ne doit pas arrêter la recherche d''une seconde ; omission et substitution demandent toutes deux une correction.

![CPL_090_001](/illustrations/cpl-090-001.svg)', '1013 became 1031 and 5274 was not repeated. Finding one error must not stop the search for another; omission and substitution both need repair.

![CPL_090_001](/illustrations/en/cpl-090-001.svg)', '["CPL", "exam-mcq", "exam-subject:090"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('356771bc-0000-4000-8000-0000356771bc', '26ad7496-0000-4000-8000-000026ad7496', 'CPL', 'communications', 'communications', 'Fermer la boucle', 3, 3, 'Exercice : vitesse minimale de configuration 160 kt, ATC demande 150 kt. Quel message expose correctement la contrainte ?', 'Exercise: configuration minimum speed 160 kt, ATC requests 150 kt. Which message correctly exposes the constraint?', '["WILCO 150, puis maintenir 160 sans annonce", "ROGER 150, puis attendre une détection ATC", "Collationner 160 comme si ce chiffre avait été assigné", "UNABLE 150, minimum 160 dans cette configuration, demandons 160"]', '["WILCO 150, then maintain 160 silently", "ROGER 150, then wait for ATC detection", "Read back 160 as if it had been assigned", "UNABLE 150, minimum 160 in this configuration, request 160"]', 'UNABLE 150, minimum 160 dans cette configuration, demandons 160', 'Annoncer impossibilité, raison et proposition distingue une demande d''une clairance. Les autres réponses masquent la non-exécution ou transforment un chiffre proposé en instruction reçue.

![CPL_090_001](/illustrations/cpl-090-001.svg)', 'Stating inability, reason and proposal distinguishes a request from clearance. Other replies conceal non-compliance or turn a proposed number into a received instruction.

![CPL_090_001](/illustrations/en/cpl-090-001.svg)', '["CPL", "exam-mcq", "exam-subject:090"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('7786c4fc-0000-4000-8000-00007786c4fc', '26ad7496-0000-4000-8000-000026ad7496', 'CPL', 'communications', 'communications', 'Fermer la boucle', 3, 3, 'Votre collationnement était 4 000 ft ; l''ATC répond « négatif, répétez [valeur masquée] ». Quelle information pouvez-vous considérer comme résolue ?', 'Your readback was 4,000 ft; ATC replies “negative, say again [blocked value]”. What can be considered resolved?', '["Le niveau prévu sur la fiche de briefing", "La différence exacte de 1 000 ft", "La présence d''un écart, mais pas le niveau corrigé", "Le niveau 4 000 ft, faute d''autre valeur entendue"]', '["The briefing-sheet level", "An exact 1,000 ft difference", "There is a discrepancy, but the corrected level is unresolved", "4,000 ft because no other value was heard"]', 'La présence d''un écart, mais pas le niveau corrigé', 'Le rejet du collationnement est reçu, sa valeur de remplacement ne l''est pas. Redemander la correction ; ni la mémoire ni une différence supposée ne comblent le message manquant.

![CPL_090_001](/illustrations/cpl-090-001.svg)', 'Rejection of the readback was heard, its replacement value was not. Request the correction again; memory and an assumed difference cannot fill the missing message.

![CPL_090_001](/illustrations/en/cpl-090-001.svg)', '["CPL", "exam-mcq", "exam-subject:090"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('7787395c-0000-4000-8000-00007787395c', '26ad7496-0000-4000-8000-000026ad7496', 'CPL', 'communications', 'communications', 'Distinguer les autorisations', 3, 3, 'Au point d''attente, autorisation « alignez-vous piste 32 et attendez ». Quelle limite doit rester dans la réponse ?', 'At the holding point, clearance “line up runway 32 and wait”. Which limit must remain in the reply?', '["Traverser 32 vers la voie opposée", "Attendre après alignement ; pas de décollage autorisé", "Maintenir avant la piste malgré l''alignement demandé", "Décoller après le précédent trafic sans autre instruction"]', '["Cross 32 to the opposite taxiway", "Wait after line-up; no take-off clearance", "Hold short despite the line-up instruction", "Take off after preceding traffic without further instruction"]', 'Attendre après alignement ; pas de décollage autorisé', 'Alignement, traversée et décollage sont différents. La limite « attendez » conserve le point d''arrêt de cette autorisation ; elle ne définit aucune condition implicite de départ.

![CPL_090_001](/illustrations/cpl-090-001.svg)', 'Line-up, crossing and take-off differ. “Wait” preserves the stopping point of this clearance; it creates no implicit departure condition.

![CPL_090_001](/illustrations/en/cpl-090-001.svg)', '["CPL", "exam-mcq", "exam-subject:090"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('7787395d-0000-4000-8000-00007787395d', '26ad7496-0000-4000-8000-000026ad7496', 'CPL', 'communications', 'communications', 'Distinguer les autorisations', 3, 3, 'Clairance d''approche reçue ; plus tard « continuez l''approche ». Aucune autorisation d''atterrir. Quel état inscrire ?', 'Approach clearance received; later “continue approach”. No landing clearance. Which state applies?', '["Approche autorisée, atterrissage encore non autorisé", "Atterrissage autorisé par la deuxième transmission", "Approche annulée en l''absence de clairance d''atterrir", "Atterrissage autorisé si la piste est visuellement libre"]', '["Approach cleared, landing not yet cleared", "Landing cleared by the second transmission", "Approach cancelled without landing clearance", "Landing cleared if runway visually clear"]', 'Approche autorisée, atterrissage encore non autorisé', '« Continuez » ne contient pas l''autorisation distincte d''atterrir. L''exercice demande la portée des messages, pas une décision de poursuite à une hauteur donnée ; cette dernière dépend de la situation réelle.

![CPL_090_001](/illustrations/cpl-090-001.svg)', '“Continue” does not contain separate landing clearance. This asks message scope, not whether to continue at a particular height; that decision depends on the actual situation.

![CPL_090_001](/illustrations/en/cpl-090-001.svg)', '["CPL", "exam-mcq", "exam-subject:090"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('778650a0-0000-4000-8000-0000778650a0', '26ad7496-0000-4000-8000-000026ad7496', 'CPL', 'communications', 'communications', 'Collationner les éléments de sécurité', 3, 3, 'Clairance conditionnelle correctement émise : derrière le trafic identifié en courte finale, alignez-vous 24 derrière. Quel collationnement préserve la condition ?', 'Properly issued conditional clearance: behind identified traffic on short final, line up 24 behind. Which readback preserves the condition?', '["Alignons 24, AER452", "Après notre contrôle visuel, décollons 24, AER452", "Maintenons avant 24 puis traversons derrière, AER452", "Derrière ce trafic, alignons 24 derrière, AER452"]', '["Lining up 24, AER452", "After our visual check, taking off 24, AER452", "Holding short 24 then crossing behind, AER452", "Behind that traffic, lining up 24 behind, AER452"]', 'Derrière ce trafic, alignons 24 derrière, AER452', 'La condition est une partie de la clairance, pas un commentaire. L''omettre transforme le sens ; décollage et traversée remplacent l''action autorisée. Le trafic doit être identifié sans ambiguïté.

![CPL_090_001](/illustrations/cpl-090-001.svg)', 'The condition is part of clearance, not commentary. Omitting it changes meaning; take-off and crossing replace the cleared action. Traffic must be unambiguously identified.

![CPL_090_001](/illustrations/en/cpl-090-001.svg)', '["CPL", "exam-mcq", "exam-subject:090"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('7787395f-0000-4000-8000-00007787395f', '26ad7496-0000-4000-8000-000026ad7496', 'CPL', 'communications', 'communications', 'Distinguer les autorisations', 3, 3, 'Dernier état AER452 : cap 220, 4 000 ft, 170 kt. Message suivant pour AER425 : 160 kt. Puis AER452 : cap 230, maintenez 3 000 ft. Quel état final ?

![CPL_090_001](/illustrations/cpl-090-001.svg)', 'Last AER452 state: heading 220, 4,000 ft, 170 kt. Next, AER425 receives 160 kt. Then AER452 receives heading 230, maintain 3,000 ft. Final state?

![CPL_090_001](/illustrations/en/cpl-090-001.svg)', '["220 / 3 000 ft / 170 kt", "230 / 4 000 ft / 170 kt", "230 / 3 000 ft / 170 kt", "230 / 3 000 ft / 160 kt"]', '["220 / 3,000 ft / 170 kt", "230 / 4,000 ft / 170 kt", "230 / 3,000 ft / 170 kt", "230 / 3,000 ft / 160 kt"]', '230 / 3 000 ft / 170 kt', 'Mettre à jour les deux éléments adressés à AER452 : cap et niveau. La vitesse 160 appartient à AER425 ; conserver 170. Les deux autres distracteurs oublient une modification correctement reçue.

![CPL_090_001](/illustrations/cpl-090-001.svg)', 'Update AER452''s two amended items: heading and level. 160 belongs to AER425; retain 170. The other two distractors each miss a correctly received amendment.

![CPL_090_001](/illustrations/en/cpl-090-001.svg)', '["CPL", "exam-mcq", "exam-subject:090"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO lessons (id, license, category, theme, title_fr, title_en, content_fr, content_en, difficulty, order_index, level, duration_minutes, tags, learning_objectives) VALUES
('1c9aa75f-0000-4000-8000-00001c9aa75f', 'CPL', 'flight_planning', 'flight_planning', 'Bilan carburant et suivi du vol', 'Fuel accounting and flight monitoring', '**Objectif :** Construire un bilan par phase, suivre les écarts sans double comptage et comparer des options d’arrivée avec protection et aptitude des terrains.

## Relier distance, temps et consommation

![CPL_033_001](/illustrations/cpl-033-001.svg)

Utiliser la vitesse sol pour le temps de parcours et le débit approprié à la phase. Temps en heures = distance NM / vitesse sol kt. Carburant = temps × débit. Pour 180 NM à 120 kt et 40 L/h constants, temps = 1,5 h et consommation = 60 L. 1 h 30 min vaut 1,5 h, pas 1,30 h. Le modèle exclut montée et descente, dont les allocations doivent être ajoutées si elles ne sont pas déjà incluses.

### Construire le calcul par phase
Avant de multiplier, écrire distance, GS, durée et débit pour chaque phase. Le débit dépend de la puissance, de la configuration et des conditions ; ne pas appliquer le débit de croisière à toute une montée sans données appropriées. La vitesse sol décrit le déplacement sur la carte, alors que la vitesse air décrit le mouvement dans l''air. Une distance en NM divisée par des kt donne des heures ; multiplier par 60 pour les minutes.

**Exemple guidé distinct du quiz :** 84 NM à 112 kt donnent 0,75 h, soit 45 min. À 36 L/h, consommation 27 L. Puis 30 NM à 90 kt donnent un tiers d''heure, soit 20 min ; à 30 L/h, 10 L. Pour ces deux segments : 65 min et 37 L. Additionner les durées et les consommations ; la moyenne arithmétique des deux GS ne permet pas, en général, de calculer le temps total.

**Erreur expliquée :** 1 h 15 min vaut 1 + 15/60 = 1,25 h. Écrire 1,15 h retire six minutes au calcul. Garder les fractions jusqu''au résultat final ; transformer prématurément un tiers d''heure en 0,3 sous-estime déjà le carburant du second segment.

Un tableau peut fournir directement une quantité pour la montée ou l''approche : la traiter comme une allocation de phase. Ne pas la multiplier à nouveau par une durée. Définir précisément début et fin du segment pour ne pas compter deux fois une phase.

Le schéma en tête présente les seules données du cas autonome de comparaison de terrains. Les exemples guidés ci-dessous utilisent des données différentes.

## Construire une allocation explicite

Exercice fictif : roulage 5 L, trajet 60 L, complément de planification 6 L, dégagement 12 L et quantité protégée 20 L. Total = 103 L. Ces allocations sont des données pédagogiques ; elles ne définissent pas une politique réglementaire universelle. Distinguer carburant total, utilisable et inutilisable. Avec 110 L utilisables au départ, l’excédent sur ce plan est 7 L. Une réserve protégée ne devient pas automatiquement du carburant de trajet disponible.

### Séparer le contenu du réservoir et les engagements du plan
Écrire d''abord le carburant utilisable. Si une donnée représente déjà l''utilisable, ne pas retrancher une deuxième fois l''inutilisable. Ensuite nommer chaque allocation et préciser les inclusions : roulage, phases du trajet, autres postes imposés et quantité à protéger. Les sommes de ce chapitre sont des hypothèses d''exercice ; la licence CPL ne définit pas à elle seule la politique carburant de toute exploitation.

**Bilan guidé :** 100 L au total dont 4 inutilisables donnent 96 L utilisables. Hypothèses : roulage 4 L, trajet 37 L, autre allocation 8 L, protection 25 L. Engagement total 74 L ; excédent au départ 22 L. Si l''on dispose au contraire de « 96 L utilisables » dans l''énoncé, le même excédent vaut 22 L, pas 18 L.

L''excédent n''est pas le carburant prévu à l''arrivée. Après le seul roulage et le trajet, il reste 96 − 4 − 37 = 55 L. Sur ces 55 L, le plan engage encore 8 + 25 = 33 L ; le solde libre reste 22 L. Cette double lecture contrôle les inclusions. Une somme correcte avec une phase comptée deux fois reste un mauvais bilan.

**À faire sans la solution :** tracer une flèche du parking à l''arrivée ; inscrire ce qui est consommé et ce qui doit rester. Refaire le bilan après changement d''une allocation, sans supprimer une protection pour rendre le total favorable.

## Mesurer les écarts

Comparer régulièrement carburant calculé, indications disponibles et temps écoulé, en tenant compte de la précision des instruments. Une consommation prévue de 40 L et constatée de 48 L donne +8 L et +20 %. Rechercher les causes et recalculer l’arrivée. Avec 70 L utilisables restants, une protection de 20 L et un débit futur de 50 L/h, le temps disponible avant cette protection est 1 h. Ne pas confondre cette limite de planification avec un droit de poursuivre.

### Transformer un écart en prévision d''arrivée
Comparer les valeurs au même point et à la même heure. Si le départ utilisable était 96 L et que le suivi indique 61 L restants, la consommation déduite est 35 L. Une prévision de 30 L à ce point donne un excès de 5 L, soit 5/30 = 16,7 % du prévu. Diviser par 35 répondrait à une autre question. Cette arithmétique ne démontre pas, seule, une fuite : vérifier données de départ, indications, durée, régime et conditions, sans négliger un indice anormal.

**Recalcul guidé :** il reste 50 min à 42 L/h, plus 5 L d''approche non inclus. Besoin futur 35 + 5 = 40 L ; arrivée prévue 61 − 40 = 21 L. Avec protection 25 L, déficit de 4 L. L''écart passé sert à revoir les hypothèses futures, mais ne doit pas être retiré une deuxième fois du carburant restant constaté.

Une moyenne historique n''est pas forcément le débit futur : un changement de puissance ou de régime demande une donnée adaptée. Quand des indications divergent, ne pas retenir automatiquement la plus favorable. Rechercher la cohérence avec les moyens prévus, traiter les anomalies selon la procédure et conserver explicitement l''incertitude dans les options.

## Réviser une décision

Un vent défavorable réduit la vitesse sol : 180 NM à 100 kt demandent 1,8 h et 72 L à 40 L/h, soit 12 L de plus qu’à 120 kt. Un trajet plus court, un terrain approprié ou un avitaillement intermédiaire peuvent être nécessaires. Pour convertir un volume en masse, utiliser la densité fournie : 100 L à 0,72 kg/L représentent 72 kg. La densité dépend du carburant et de sa température ; ne pas supposer que 1 L vaut 1 kg.

### Comparer des options avec leurs contraintes
Pour chaque terrain : faisabilité de la route et de l''atterrissage, carburant jusqu''à l''arrivée, quantité restante et protection à maintenir. Une route courte ne suffit pas si piste, météo ou performances la rendent inappropriée. Une autonomie calculée ne garantit pas qu''un aérodrome adéquat se trouve à portée.

**Cas guidé :** 61 L utilisables, protection 25 L. A demande 40 L jusqu''à l''arrivée, B 29 L ; les deux terrains et routes sont appropriés par hypothèse. A laisserait 21 L, B 32 L. B respecte le critère carburant donné avec 7 L de marge ; A présente un déficit. Si l''aptitude de B était inconnue, il faudrait l''établir plutôt que conclure sur le seul solde.

**Conversion et décision :** masse = volume × densité. Pour ajouter 18 kg à une densité fournie de 0,75 kg/L, il faut 24 L. Vérifier ensuite capacités du réservoir, masse et centrage. Ni une conversion exacte ni un réservoir de capacité suffisante ne prouvent toutes les autres limites respectées.

Le suivi est continu : actualiser le point de décision avant que les options disparaissent. Les quantités protégées de ces exercices ne définissent pas les seuils d''annonce MINIMUM FUEL ou MAYDAY FUEL ; ces procédures et la politique réglementaire applicable constituent un apprentissage distinct.

**Format de travail :** 15 QCM de chapitre, quatre choix et une réponse correcte. La matière 033 comporte 33 questions en 90 minutes dans le tableau EASA cité. Ce chapitre ne couvre qu’une partie de la matière ; il ne constitue pas une épreuve complète. Cadre EASA, même en anglais ou avec les unités US/britanniques.

## Références

- [EASA — programme Part-FCL](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8)
- [EASA — format de la matière, AMC1 ARA.FCL.300(b)](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67)
- [FAA PHAK — navigation, calculs de temps et carburant](https://www.faa.gov/regulationspolicies/handbooksmanuals/aviation/phak/chapter-16-navigation)

**Support de révision CPL(A).** À faire valider par votre instructeur. Pour un vol réel, utiliser le manuel de l’avion et les publications en vigueur. Les exemples chiffrés sont pédagogiques.', '**Objective:** Build a phase-based fuel balance, monitor deviations without double counting and compare arrivals against protection and aerodrome suitability.

## Relate distance, time and consumption

![CPL_033_001](/illustrations/en/cpl-033-001.svg)

Use groundspeed for travel time and fuel flow appropriate to the phase. Time in hours = distance NM / groundspeed kt. Fuel = time × flow. For 180 NM at 120 kt and constant 40 L/h, time = 1.5 h and burn = 60 L. 1 h 30 min is 1.5 h, not 1.30 h. This model excludes climb and descent, whose allocations must be added if not already included.

### Build the calculation by phase
Before multiplying, record distance, GS, duration and flow for each phase. Flow depends on power, configuration and conditions; do not apply cruise flow to an entire climb without suitable data. Groundspeed describes displacement on a chart; airspeed describes motion through air. NM divided by knots gives hours; multiply by 60 for minutes.

**Worked example distinct from the quiz:** 84 NM at 112 kt gives 0.75 h, or 45 minutes. At 36 L/h, burn is 27 L. Then 30 NM at 90 kt gives one third of an hour, or 20 minutes; at 30 L/h, 10 L. Totals: 65 minutes and 37 L. Add times and burns; the arithmetic average of both groundspeeds generally does not give total travel time.

**Explained error:** 1 h 15 min is 1 + 15/60 = 1.25 h. Writing 1.15 h removes six minutes. Retain fractions until the final result; prematurely rounding one third to 0.3 already underestimates second-segment fuel.

A table may directly supply a climb or approach quantity: treat it as a phase allocation, not a flow to multiply by time again. Define segment boundaries precisely to avoid counting any phase twice.

The opening diagram supplies only the inputs for the independent aerodrome comparison case. Worked examples below use different inputs.

## Build an explicit allocation

Fictional exercise: taxi 5 L, trip 60 L, planning supplement 6 L, alternate 12 L and protected quantity 20 L. Total = 103 L. These are training inputs, not a universal regulatory fuel policy. Distinguish total, usable and unusable fuel. With 110 L usable at departure, surplus over this plan is 7 L. A protected reserve does not automatically become available trip fuel.

### Separate tank contents from plan commitments
First establish usable fuel. If an input is already usable fuel, do not subtract unusable fuel again. Name each allocation and its inclusions: taxi, trip phases, other required items and protected quantity. These chapter totals are exercise assumptions; a CPL licence alone does not define every operation''s fuel policy.

**Worked balance:** 100 L total including 4 unusable gives 96 L usable. Assumptions: taxi 4 L, trip 37 L, other allocation 8 L, protected 25 L. Total commitment 74 L; departure surplus 22 L. If the input instead states “96 L usable”, surplus remains 22 L, not 18 L.

Surplus is not predicted arrival fuel. After taxi and trip only, 96 − 4 − 37 = 55 L remains. Of this, the plan still commits 8 + 25 = 33 L; free balance remains 22 L. This second reading checks inclusions. Correct addition with a phase counted twice still gives a wrong plan.

**Work before seeing answers:** draw a timeline from parking to arrival; label consumption and what must remain. Rebuild the balance after one allocation changes without deleting protection to make the total favourable.

## Measure deviations

Regularly compare calculated fuel, available indications and elapsed time, considering instrument accuracy. Expected burn 40 L versus observed 48 L gives +8 L and +20%. Investigate causes and recalculate arrival fuel. With 70 L usable remaining, 20 L protected and future flow 50 L/h, time before reaching that protection is 1 h. This planning limit is not permission to continue.

### Turn a deviation into an arrival prediction
Compare values at the same point and time. With 96 L usable at departure and 61 L now indicated remaining, inferred burn is 35 L. Planned burn of 30 L to this point gives 5 L excess, or 5/30 = 16.7% of plan. Dividing by 35 answers a different question. Arithmetic alone does not prove a leak: check departure inputs, indications, time, power and conditions without dismissing an abnormal cue.

**Worked replan:** 50 minutes remain at 42 L/h plus 5 L approach not included. Future need is 35 + 5 = 40 L; predicted arrival is 61 − 40 = 21 L. Against 25 L protection, deficit is 4 L. Past deviation informs future assumptions but must not be subtracted again from observed remaining fuel.

Historical average flow is not necessarily future flow: power or operating changes require appropriate data. If indications disagree, do not automatically select the favourable one. Cross-check through prescribed means, handle anomalies under the procedure and retain uncertainty explicitly when comparing options.

## Revise a decision

Adverse wind reduces groundspeed: 180 NM at 100 kt requires 1.8 h and 72 L at 40 L/h, 12 L more than at 120 kt. A shorter route, suitable airfield or intermediate refuelling may be needed. Convert volume to mass using the supplied density: 100 L at 0.72 kg/L is 72 kg. Density depends on fuel and temperature; do not assume 1 L equals 1 kg.

### Compare options with their constraints
For each aerodrome consider route and landing feasibility, fuel to arrival, remaining quantity and required protection. A short route is insufficient if runway, weather or performance makes the airfield unsuitable. Calculated endurance does not ensure a suitable aerodrome is reachable.

**Worked case:** 61 L usable, 25 L protected. A needs 40 L to arrival, B 29 L; both routes and aerodromes are suitable by assumption. A leaves 21 L, B 32 L. B meets the stated fuel criterion with 7 L margin; A has a deficit. If B''s suitability were unknown, establish it rather than concluding from fuel alone.

**Conversion and decision:** mass = volume × density. Adding 18 kg at a supplied density of 0.75 kg/L requires 24 L. Then check tank capacity, mass and CG. Neither accurate conversion nor sufficient tank capacity proves all other limits met.

Monitoring is continuous: update the decision point before options disappear. Protected quantities in these exercises do not define MINIMUM FUEL or MAYDAY FUEL declaration thresholds; those procedures and applicable regulatory policy are a separate learning task.

**Practice format:** 15 chapter MCQs, four choices and one correct answer. Subject 033 has 33 questions in 90 minutes in the cited EASA table. This chapter covers only part of the subject and is not a full paper. EASA framework, including with English or US/British units.

## References

- [Reference 1](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8)
- [Reference 2](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67)
- [Reference 3](https://www.faa.gov/regulationspolicies/handbooksmanuals/aviation/phak/chapter-16-navigation)

**CPL(A) revision material.** To be reviewed by your instructor. For an actual flight, use the aircraft manual and current publications. Numerical examples are for training.', 3, 1, 1, 50, '["CPL", "mcq-four-options"]', '["Construire un bilan par phase, suivre les écarts sans double comptage et comparer des options d’arrivée avec protection et aptitude des terrains."]')
ON CONFLICT (id) DO UPDATE SET license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, title_fr = EXCLUDED.title_fr, title_en = EXCLUDED.title_en, content_fr = EXCLUDED.content_fr, content_en = EXCLUDED.content_en, difficulty = EXCLUDED.difficulty, order_index = EXCLUDED.order_index, level = EXCLUDED.level, duration_minutes = EXCLUDED.duration_minutes, tags = EXCLUDED.tags, learning_objectives = EXCLUDED.learning_objectives;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('3c7714c3-0000-4000-8000-00003c7714c3', '1c9aa75f-0000-4000-8000-00001c9aa75f', 'CPL', 'flight_planning', 'flight_planning', 'Relier distance, temps et consommation', 3, 3, 'Segment A : 90 NM à 120 kt sol, puis B : 40 NM à 80 kt. Durée totale ?', 'Segment A: 90 NM at 120 kt GS, then B: 40 NM at 80 kt. Total time?', '["1 h 15 min", "1 h 18 min", "1 h 25 min", "1 h 30 min"]', '["1 h 15 min", "1 h 18 min", "1 h 25 min", "1 h 30 min"]', '1 h 15 min', 'A prend 45 min, B 30 min : 75 min. 78 min proviennent de 130 NM divisés par la moyenne simple 100 kt ; chaque vitesse doit être appliquée à son propre segment.

![CPL_033_001](/illustrations/cpl-fuel-method.svg)', 'A takes 45 minutes, B 30: total 75. 78 minutes comes from dividing 130 NM by the simple 100 kt average; apply each speed to its own segment.

![CPL_033_001](/illustrations/en/cpl-fuel-method.svg)', '["CPL", "exam-mcq", "exam-subject:033"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('3c7714c4-0000-4000-8000-00003c7714c4', '1c9aa75f-0000-4000-8000-00001c9aa75f', 'CPL', 'flight_planning', 'flight_planning', 'Relier distance, temps et consommation', 3, 3, 'Segment A : 90 NM à 120 kt et 32 L/h ; B : 40 NM à 80 kt et 28 L/h. Carburant des deux segments ?', 'Segment A: 90 NM at 120 kt and 32 L/h; B: 40 NM at 80 kt and 28 L/h. Combined fuel?', '["40 L", "35 L", "37,5 L", "38 L"]', '["40 L", "35 L", "37.5 L", "38 L"]', '38 L', 'A : 0,75 × 32 = 24 L ; B : 0,5 × 28 = 14 L ; total 38. 40 applique le premier débit partout, 35 le second ; 37,5 utilise leur moyenne simple sans pondérer les durées.

![CPL_033_001](/illustrations/cpl-fuel-method.svg)', 'A: 0.75 × 32 = 24 L; B: 0.5 × 28 = 14 L; total 38. 40 applies the first flow throughout, 35 the second; 37.5 uses their unweighted average.

![CPL_033_001](/illustrations/en/cpl-fuel-method.svg)', '["CPL", "exam-mcq", "exam-subject:033"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('3c7714c5-0000-4000-8000-00003c7714c5', '1c9aa75f-0000-4000-8000-00001c9aa75f', 'CPL', 'flight_planning', 'flight_planning', 'Relier distance, temps et consommation', 3, 3, '1 h 24 min à 35 L/h. Quelle consommation ?', '1 h 24 min at 35 L/h. What burn?', '["14 L", "59 L", "49 L", "43,4 L"]', '["14 L", "59 L", "49 L", "43.4 L"]', '49 L', '24/60 = 0,4 h ; 1,4 × 35 = 49 L. 43,4 traite 1 h 24 comme 1,24 h ; 14 omet l''heure entière. Additionner 35 et 24 pour obtenir 59 mélange un débit et une durée.

![CPL_033_001](/illustrations/cpl-fuel-method.svg)', '24/60 = 0.4 h; 1.4 × 35 = 49 L. 43.4 treats 1 h 24 as 1.24 h; 14 omits the full hour. Adding 35 and 24 to get 59 mixes a flow and a duration.

![CPL_033_001](/illustrations/en/cpl-fuel-method.svg)', '["CPL", "exam-mcq", "exam-subject:033"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('3c771887-0000-4000-8000-00003c771887', '1c9aa75f-0000-4000-8000-00001c9aa75f', 'CPL', 'flight_planning', 'flight_planning', 'Construire une allocation explicite', 3, 3, 'Plan fictif : roulage 4 L, montée 8 L, croisière 38 L, descente/approche 6 L, autre allocation 9 L, protection 24 L. Aucun poste n''est inclus dans un autre. Total ?', 'Fictional plan: taxi 4 L, climb 8 L, cruise 38 L, descent/approach 6 L, other allocation 9 L, protected 24 L. No overlap. Total?', '["97 L", "89 L", "85 L", "65 L"]', '["97 L", "89 L", "85 L", "65 L"]', '89 L', '4+8+38+6+9+24 = 89 L. 85 omet le roulage, 65 la protection ; 97 compte deux fois la montée. Les inclusions sont explicitement disjointes.

![CPL_033_001](/illustrations/cpl-fuel-method.svg)', '4+8+38+6+9+24 = 89 L. 85 omits taxi, 65 protection; 97 counts climb twice. Inputs explicitly state no overlap.

![CPL_033_001](/illustrations/en/cpl-fuel-method.svg)', '["CPL", "exam-mcq", "exam-subject:033"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('3c771888-0000-4000-8000-00003c771888', '1c9aa75f-0000-4000-8000-00001c9aa75f', 'CPL', 'flight_planning', 'flight_planning', 'Construire une allocation explicite', 3, 3, 'Réservoirs : 104 L au total dont 4 inutilisables. Plan exigeant 89 L utilisables. Excédent au départ ?', 'Tanks contain 104 L total including 4 unusable. Plan requires 89 L usable. Departure surplus?', '["11 L", "15 L", "7 L", "24 L"]', '["11 L", "15 L", "7 L", "24 L"]', '11 L', 'Utilisable = 104−4 = 100 ; excédent 100−89 = 11. 15 compte l''inutilisable, 7 le retire deux fois ; 24 est une allocation du plan, pas l''excédent.

![CPL_033_001](/illustrations/cpl-fuel-method.svg)', 'Usable = 104−4 = 100; surplus 100−89 = 11. 15 counts unusable fuel, 7 subtracts it twice; 24 is a plan allocation, not surplus.

![CPL_033_001](/illustrations/en/cpl-fuel-method.svg)', '["CPL", "exam-mcq", "exam-subject:033"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('3c771889-0000-4000-8000-00003c771889', '1c9aa75f-0000-4000-8000-00001c9aa75f', 'CPL', 'flight_planning', 'flight_planning', 'Construire une allocation explicite', 3, 3, 'La fiche indique déjà 100 L utilisables ; capacité inutilisable 4 L. Pour un besoin de 89 L utilisables, quelle marge ?', 'The sheet already states 100 L usable; unusable capacity is 4 L. For an 89 L usable requirement, what margin?', '["7 L", "15 L", "4 L", "11 L"]', '["7 L", "15 L", "4 L", "11 L"]', '11 L', 'La donnée utilisable exclut déjà l''inutilisable : 100−89 = 11. Retrancher encore 4 change la définition de la donnée ; ajouter 4 rendrait disponible un carburant qui ne l''est pas.

![CPL_033_001](/illustrations/cpl-fuel-method.svg)', 'The usable input already excludes unusable fuel: 100−89 = 11. Subtracting 4 again changes the input definition; adding it would treat unavailable fuel as usable.

![CPL_033_001](/illustrations/en/cpl-fuel-method.svg)', '["CPL", "exam-mcq", "exam-subject:033"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('3c771c4b-0000-4000-8000-00003c771c4b', '1c9aa75f-0000-4000-8000-00001c9aa75f', 'CPL', 'flight_planning', 'flight_planning', 'Mesurer les écarts', 3, 3, 'Au même point, prévu consommé 30 L ; départ utilisable 100 L, restant constaté 64 L. Surconsommation absolue ?', 'At the same point, planned burn 30 L; starting usable fuel 100 L, observed remaining 64 L. Absolute excess burn?', '["34 L", "20 L", "6 L", "36 L"]', '["34 L", "20 L", "6 L", "36 L"]', '6 L', 'Consommation déduite 100−64 = 36 L ; excès 36−30 = 6. 36 est le consommé total ; 34 compare à tort restant et consommé prévu. Vérifier aussi la fiabilité des données mesurées.

![CPL_033_001](/illustrations/cpl-fuel-method.svg)', 'Inferred burn 100−64 = 36 L; excess 36−30 = 6. 36 is total burn; 34 wrongly compares remaining fuel with planned burn. Also check reliability of measured inputs.

![CPL_033_001](/illustrations/en/cpl-fuel-method.svg)', '["CPL", "exam-mcq", "exam-subject:033"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('3c771c4c-0000-4000-8000-00003c771c4c', '1c9aa75f-0000-4000-8000-00001c9aa75f', 'CPL', 'flight_planning', 'flight_planning', 'Mesurer les écarts', 3, 3, 'Prévu 30 L, consommé déduit 36 L. Écart relatif au prévu ?', 'Planned 30 L, inferred burn 36 L. Deviation relative to plan?', '["120 %", "20 %", "16,7 %", "6 %"]', '["120%", "20%", "16.7%", "6%"]', '20 %', 'L''écart 6 est divisé par la base demandée 30 : 20 %. 16,7 % utilise 36 ; 120 % est le ratio total 36/30, non le seul excès ; 6 reste une quantité en litres.

![CPL_033_001](/illustrations/cpl-fuel-method.svg)', 'Divide the 6 L difference by the requested 30 L reference: 20%. 16.7% uses 36; 120% is total ratio 36/30, not excess alone; 6 is a quantity in litres.

![CPL_033_001](/illustrations/en/cpl-fuel-method.svg)', '["CPL", "exam-mcq", "exam-subject:033"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('3c771c4d-0000-4000-8000-00003c771c4d', '1c9aa75f-0000-4000-8000-00001c9aa75f', 'CPL', 'flight_planning', 'flight_planning', 'Mesurer les écarts', 3, 3, 'Restant utilisable 64 L ; 45 min à 40 L/h puis approche 5 L non incluse. Protection à l''arrivée 24 L. Marge prévue au-dessus de cette protection ?', 'Usable remaining 64 L; 45 minutes at 40 L/h then 5 L approach not included. Arrival protection 24 L. Predicted margin above protection?', '["5 L", "11 L", "10 L", "29 L"]', '["5 L", "11 L", "10 L", "29 L"]', '5 L', 'Trajet futur 30 L, approche 5 : arrivée 29 L ; marge 29−24 = 5. 10 omet l''approche ; 29 est l''arrivée et non la marge. Un écart passé déjà reflété dans 64 L ne se soustrait pas à nouveau.

![CPL_033_001](/illustrations/cpl-fuel-method.svg)', 'Future trip 30 L, approach 5: arrival 29 L; margin 29−24 = 5. 10 omits approach; 29 is arrival fuel, not margin. Past excess already reflected in 64 L is not subtracted again.

![CPL_033_001](/illustrations/en/cpl-fuel-method.svg)', '["CPL", "exam-mcq", "exam-subject:033"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('3c77200f-0000-4000-8000-00003c77200f', '1c9aa75f-0000-4000-8000-00001c9aa75f', 'CPL', 'flight_planning', 'flight_planning', 'Réviser une décision', 3, 3, '180 NM à 100 kt sol, débit futur donné 40 L/h ; aucune autre phase incluse. Consommation de ce segment ?', '180 NM at 100 kt GS, given future flow 40 L/h; no other phase included. Segment burn?', '["60 L", "64 L", "80 L", "72 L"]', '["60 L", "64 L", "80 L", "72 L"]', '72 L', '180/100 = 1,8 h ; ×40 = 72 L. 60 conserve l''ancien calcul à 120 kt ; 80 arrondit prématurément la durée à deux heures. Ce résultat n''inclut pas les autres allocations.

![CPL_033_001](/illustrations/cpl-fuel-method.svg)', '180/100 = 1.8 h; ×40 = 72 L. 60 retains the previous 120 kt calculation; 80 rounds time prematurely to two hours. This excludes other allocations.

![CPL_033_001](/illustrations/en/cpl-fuel-method.svg)', '["CPL", "exam-mcq", "exam-subject:033"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('526ce109-0000-4000-8000-0000526ce109', '1c9aa75f-0000-4000-8000-00001c9aa75f', 'CPL', 'flight_planning', 'flight_planning', 'Réviser une décision', 3, 3, 'Même débit : ancien trajet 180 NM à 120 kt ; nouveau 180 NM à 100 kt, 40 L/h. Avec excédent initial de 7 L, quel solde après cette seule modification ?', 'Same flow: previous 180 NM at 120 kt; now 180 NM at 100 kt, 40 L/h. With initial surplus 7 L, what balance after this change alone?', '["Déficit de 12 L", "Excédent de 19 L", "Déficit de 5 L", "Excédent de 5 L"]', '["12 L deficit", "19 L surplus", "5 L deficit", "5 L surplus"]', 'Déficit de 5 L', 'Le supplément vaut 72−60 = 12 L. Nouveau solde 7−12 = −5 L. Ne pas confondre supplément et déficit final, ni ajouter une consommation à une marge.

![CPL_033_001](/illustrations/cpl-fuel-method.svg)', 'Extra burn is 72−60 = 12 L. New balance 7−12 = −5 L. Do not confuse additional consumption with final deficit or add burn to a margin.

![CPL_033_001](/illustrations/en/cpl-fuel-method.svg)', '["CPL", "exam-mcq", "exam-subject:033"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('526ce10a-0000-4000-8000-0000526ce10a', '1c9aa75f-0000-4000-8000-00001c9aa75f', 'CPL', 'flight_planning', 'flight_planning', 'Réviser une décision', 3, 3, 'Il faut ajouter 21,6 kg de carburant, densité donnée 0,72 kg/L. Quel volume ajouter ?', 'Add 21.6 kg fuel at supplied density 0.72 kg/L. What volume is needed?', '["72 L", "30 L", "15,552 L", "21,6 L"]', '["72 L", "30 L", "15.552 L", "21.6 L"]', '30 L', 'Volume = masse/densité = 21,6/0,72 = 30 L. Multiplier donne 15,552 avec une opération inadaptée ; kg et L ne sont pas interchangeables. Contrôler ensuite capacité et chargement.

![CPL_033_001](/illustrations/cpl-fuel-method.svg)', 'Volume = mass/density = 21.6/0.72 = 30 L. Multiplying gives 15.552 using the wrong operation; kg and L are not interchangeable. Then check capacity and loading.

![CPL_033_001](/illustrations/en/cpl-fuel-method.svg)', '["CPL", "exam-mcq", "exam-subject:033"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('526c6cac-0000-4000-8000-0000526c6cac', '1c9aa75f-0000-4000-8000-00001c9aa75f', 'CPL', 'flight_planning', 'flight_planning', 'Mesurer les écarts', 3, 3, 'Restant 58 L, protection 24 L. A demande 40 min à 42 L/h + 4 L approche ; B 30 min à 42 L/h + 4 L mais piste inappropriée. C demande 50 min à 42 L/h + 4 L. A et C sont appropriés. Quelle option satisfait les contraintes données ?

![CPL_033_001](/illustrations/cpl-033-001.svg)', 'Remaining 58 L, protected 24 L. A needs 40 min at 42 L/h + 4 L approach; B 30 min at 42 L/h + 4 L but runway unsuitable. C needs 50 min at 42 L/h + 4 L. A and C are suitable. Which meets the stated constraints?

![CPL_033_001](/illustrations/en/cpl-033-001.svg)', '["A, arrivée 26 L", "B, arrivée 33 L", "C, arrivée 19 L", "A, arrivée 30 L"]', '["A, arrival 26 L", "B, arrival 33 L", "C, arrival 19 L", "A, arrival 30 L"]', 'A, arrivée 26 L', 'A consomme 28+4 = 32 L et laisse 26, au-dessus de 24. B laisse plus de carburant mais sa piste l''exclut ; C laisse 19, sous la protection. 30 pour A omet l''approche.

![CPL_033_001](/illustrations/cpl-033-001.svg)', 'A burns 28+4 = 32 L and leaves 26, above 24. B leaves more fuel but its runway excludes it; C leaves 19, below protection. A''s 30 L answer omits approach.

![CPL_033_001](/illustrations/en/cpl-033-001.svg)', '["CPL", "exam-mcq", "exam-subject:033"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('526bf84e-0000-4000-8000-0000526bf84e', '1c9aa75f-0000-4000-8000-00001c9aa75f', 'CPL', 'flight_planning', 'flight_planning', 'Construire une allocation explicite', 3, 3, 'Deux documents de trajet incluent chacun la même montée de 8 L. Comment éviter une addition erronée ?', 'Two trip documents each include the same 8 L climb. How should erroneous addition be avoided?', '["Additionner les deux par principe de prudence", "Supprimer toute allocation de montée", "Retrancher 8 L de la quantité protégée", "Identifier les limites des phases et ne compter la montée qu''une fois"]', '["Always add both as a precaution", "Remove all climb allocation", "Subtract 8 L from protected quantity", "Identify phase boundaries and count climb once"]', 'Identifier les limites des phases et ne compter la montée qu''une fois', 'La même consommation ne se produit qu''une fois dans ce scénario. Réconcilier les inclusions garde un bilan traçable ; une marge supplémentaire décidée séparément ne doit pas être confondue avec ce double comptage.

![CPL_033_001](/illustrations/cpl-fuel-method.svg)', 'The same burn occurs once in this scenario. Reconciling inclusions keeps the balance traceable; a separately chosen additional margin must not be confused with double counting.

![CPL_033_001](/illustrations/en/cpl-fuel-method.svg)', '["CPL", "exam-mcq", "exam-subject:033"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('526b83f0-0000-4000-8000-0000526b83f0', '1c9aa75f-0000-4000-8000-00001c9aa75f', 'CPL', 'flight_planning', 'flight_planning', 'Relier distance, temps et consommation', 3, 3, 'Navigation : TAS 125 kt, vent de face axial 25 kt supposé constant, distance 150 NM. Quelle durée de segment ?', 'Navigation: TAS 125 kt, axial headwind 25 kt assumed constant, distance 150 NM. Segment time?', '["1 h", "6 h", "1 h 30 min", "1 h 12 min"]', '["1 h", "6 h", "1 h 30 min", "1 h 12 min"]', '1 h 30 min', 'GS = 125−25 = 100 kt ; 150/100 = 1,5 h. 1 h 12 emploie TAS, 1 h ajoute le vent de face comme vent arrière ; 6 h utilise la vitesse du vent seule. L''hypothèse est un vent axial, pas un triangle général.

![CPL_033_001](/illustrations/cpl-fuel-method.svg)', 'GS = 125−25 = 100 kt; 150/100 = 1.5 h. 1 h 12 uses TAS, 1 h adds headwind as tailwind; 6 h uses wind speed alone. This assumes axial wind, not a general wind triangle.

![CPL_033_001](/illustrations/en/cpl-fuel-method.svg)', '["CPL", "exam-mcq", "exam-subject:033"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO lessons (id, license, category, theme, title_fr, title_en, content_fr, content_en, difficulty, order_index, level, duration_minutes, tags, learning_objectives) VALUES
('1e25397b-0000-4000-8000-00001e25397b', 'CPL', 'human_performance', 'human_performance', 'Menaces, erreurs et décisions', 'Threats, errors and decisions', '**Objectif :** Résoudre et vérifier les exercices de menaces, erreurs et décisions.

## Distinguer menace et erreur

![CPL_040_001](/illustrations/cpl-040-001.svg)

Une menace est une circonstance extérieure à l’action de l’équipage qui augmente la complexité : météo dégradée, piste contaminée ou clairance tardive. Une erreur est une action ou omission de l’équipage qui s’écarte de l’intention ou du résultat attendu : mauvais calage, oubli ou confusion de piste. Un état indésirable de l’avion peut en résulter, par exemple une approche non stabilisée. L’objectif du TEM est d’anticiper les menaces, détecter et corriger les erreurs, puis récupérer les marges lorsque nécessaire.

## Installer des barrières

Le briefing associe chaque menace pertinente à une réponse concrète : limites, terrain de dégagement, répartition des tâches et critères d’interruption. Une checklist vérifie les actions selon la procédure ; réciter de mémoire sans contrôle de l’état réel peut laisser passer une erreur. Une vérification indépendante apporte davantage qu’une répétition du même raisonnement avec les mêmes données erronées. Après une interruption, reprendre au point prévu par la procédure, en vérifiant les éléments dont l’état est incertain.

## Reconnaître les biais

Le biais de confirmation pousse à retenir les indices favorables à l’hypothèse initiale et à minimiser les contradictions. Le biais de poursuite du plan conduit à continuer malgré des conditions qui justifieraient une modification. Une décision déjà coûteuse ne rend pas la poursuite plus sûre. Définir des seuils avant le vol et rechercher volontairement des indices contraires aide à réévaluer la situation. L’automatisation doit être surveillée : vérifier le mode actif et le résultat réel.

## Gérer les ressources disponibles

La fatigue peut altérer vigilance, mémoire de travail et jugement, même si le pilote se sent motivé. La caféine ne remplace pas un repos approprié. Sous forte charge, prioriser la maîtrise de la trajectoire, différer les tâches secondaires et demander de l’aide ou du temps à l’ATC. Une communication assertive décrit le problème et propose une action ; elle n’exige pas d’attendre l’accord silencieux de tous. Reconsidérer l’aptitude au vol avant le départ et réévaluer les risques pendant le vol.

**Format de travail :** 15 QCM de chapitre, quatre choix et une réponse correcte. La matière 040 comporte 35 questions en 60 minutes dans le tableau EASA cité. Ce chapitre ne couvre qu’une partie de la matière ; il ne constitue pas une épreuve complète. Cadre EASA, même en anglais ou avec les unités US/britanniques.

## Références

- [EASA — programme Part-FCL](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8)
- [EASA — format de la matière, AMC1 ARA.FCL.300(b)](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67)
- [FAA Risk Management Handbook](https://www.faa.gov/sites/faa.gov/files/2022-06/risk_management_handbook_2A.pdf)

**Support de révision CPL(A).** À faire valider par votre instructeur. Pour un vol réel, utiliser le manuel de l’avion et les publications en vigueur. Les exemples chiffrés sont pédagogiques.', '**Objective:** Solve and cross-check exercises on threats, errors and decisions.

## Distinguish threat from error

![CPL_040_001](/illustrations/en/cpl-040-001.svg)

A threat is a circumstance outside crew action that increases complexity: deteriorating weather, a contaminated runway or a late clearance. An error is a crew action or omission that deviates from intention or the expected outcome: an incorrect setting, omission or runway confusion. An undesired aircraft state, such as an unstable approach, may result. TEM aims to anticipate threats, detect and correct errors, and recover margins when necessary.

## Build defences

The briefing links each relevant threat to a concrete response: limits, alternate aerodrome, task allocation and discontinuation criteria. A checklist verifies actions according to procedure; reciting from memory without checking actual state can miss an error. Independent verification adds more than repeating the same reasoning with the same wrong inputs. After interruption, resume at the point specified by procedure and verify items whose state is uncertain.

## Recognise biases

Confirmation bias favours evidence supporting the initial hypothesis and discounts contradictions. Plan continuation bias encourages continuing despite conditions that warrant a change. A decision that has already cost time or money does not make continuation safer. Setting thresholds before flight and deliberately seeking contrary evidence helps reassessment. Monitor automation: verify the active mode and actual result.

## Manage available resources

Fatigue can impair vigilance, working memory and judgement even when the pilot feels motivated. Caffeine does not replace appropriate rest. Under high workload, prioritise aircraft control, defer secondary tasks and request assistance or time from ATC. Assertive communication states the problem and proposes an action; it does not require waiting for everyone’s silent agreement. Reconsider fitness before departure and reassess risks during flight.

**Practice format:** 15 chapter MCQs, four choices and one correct answer. Subject 040 has 35 questions in 60 minutes in the cited EASA table. This chapter covers only part of the subject and is not a full paper. EASA framework, including with English or US/British units.

## References

- [Reference 1](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8)
- [Reference 2](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67)
- [Reference 3](https://www.faa.gov/sites/faa.gov/files/2022-06/risk_management_handbook_2A.pdf)

**CPL(A) revision material.** To be reviewed by your instructor. For an actual flight, use the aircraft manual and current publications. Numerical examples are for training.', 3, 1, 1, 35, '["CPL", "mcq-four-options"]', '["Résoudre et vérifier les exercices de menaces, erreurs et décisions."]')
ON CONFLICT (id) DO UPDATE SET license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, title_fr = EXCLUDED.title_fr, title_en = EXCLUDED.title_en, content_fr = EXCLUDED.content_fr, content_en = EXCLUDED.content_en, difficulty = EXCLUDED.difficulty, order_index = EXCLUDED.order_index, level = EXCLUDED.level, duration_minutes = EXCLUDED.duration_minutes, tags = EXCLUDED.tags, learning_objectives = EXCLUDED.learning_objectives;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('22e764fb-0000-4000-8000-000022e764fb', '1e25397b-0000-4000-8000-00001e25397b', 'CPL', 'human_performance', 'human_performance', 'Distinguer menace et erreur', 3, 3, 'Une piste contaminée annoncée avant le vol est principalement…', 'A contaminated runway reported before flight is primarily…', '["Une menace à anticiper", "Une erreur de collationnement", "Une preuve de compétence", "Un état de fatigue"]', '["A threat to anticipate", "A readback error", "Proof of competence", "A fatigue state"]', 'Une menace à anticiper', 'La contamination est une condition extérieure qui augmente la complexité.

![CPL_040_001](/illustrations/cpl-040-001.svg)', 'Contamination is an external condition increasing complexity.

![CPL_040_001](/illustrations/en/cpl-040-001.svg)', '["CPL", "exam-mcq", "exam-subject:040"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('22e764fc-0000-4000-8000-000022e764fc', '1e25397b-0000-4000-8000-00001e25397b', 'CPL', 'human_performance', 'human_performance', 'Distinguer menace et erreur', 3, 3, 'Le pilote sélectionne le mauvais QNH malgré une transmission correcte. Il s’agit…', 'The pilot selects the wrong QNH despite a correct transmission. This is…', '["D’une menace météo uniquement", "D’une panne radio certaine", "D’une limitation de piste", "D’une erreur de l’équipage"]', '["Only a weather threat", "A definite radio failure", "A runway limitation", "A crew error"]', 'D’une erreur de l’équipage', 'La sélection s’écarte de l’action attendue.

![CPL_040_001](/illustrations/cpl-040-001.svg)', 'The selection deviates from the expected action.

![CPL_040_001](/illustrations/en/cpl-040-001.svg)', '["CPL", "exam-mcq", "exam-subject:040"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('22e764fd-0000-4000-8000-000022e764fd', '1e25397b-0000-4000-8000-00001e25397b', 'CPL', 'human_performance', 'human_performance', 'Distinguer menace et erreur', 3, 3, 'Après plusieurs erreurs, l’avion arrive trop haut et trop vite en finale. Cela illustre…', 'After several errors, the aircraft reaches final too high and too fast. This illustrates…', '["Une amélioration de marge", "Une menace extérieure sans effet", "Un état indésirable de l’avion", "Un collationnement correct"]', '["An improved margin", "An external threat with no effect", "An undesired aircraft state", "A correct readback"]', 'Un état indésirable de l’avion', 'La situation de trajectoire exige une réponse adaptée aux critères établis.

![CPL_040_001](/illustrations/cpl-040-001.svg)', 'The flight-path state requires a response appropriate to established criteria.

![CPL_040_001](/illustrations/en/cpl-040-001.svg)', '["CPL", "exam-mcq", "exam-subject:040"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('22e768bf-0000-4000-8000-000022e768bf', '1e25397b-0000-4000-8000-00001e25397b', 'CPL', 'human_performance', 'human_performance', 'Installer des barrières', 3, 3, 'Un briefing utile associe la menace météo à…', 'A useful briefing links the weather threat to…', '["L’absence de terrain alternatif", "Des options et critères de décision", "Une simple lecture du titre météo", "Une promesse de poursuivre"]', '["The absence of an alternate", "Options and decision criteria", "Merely reading the weather heading", "A promise to continue"]', 'Des options et critères de décision', 'Une barrière efficace prévoit des actions réalisables.

![CPL_040_001](/illustrations/cpl-040-001.svg)', 'An effective defence prepares feasible actions.

![CPL_040_001](/illustrations/en/cpl-040-001.svg)', '["CPL", "exam-mcq", "exam-subject:040"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('22e768c0-0000-4000-8000-000022e768c0', '1e25397b-0000-4000-8000-00001e25397b', 'CPL', 'human_performance', 'human_performance', 'Installer des barrières', 3, 3, 'Deux calculs reprennent la même masse saisie par erreur. Le second calcul…', 'Two calculations reuse the same incorrectly entered mass. The second calculation…', '["Peut reproduire la même erreur", "Garantit une masse correcte", "Rend les données indépendantes", "Supprime tout besoin de contrôle"]', '["May reproduce the same error", "Guarantees a correct mass", "Makes the inputs independent", "Removes any need for verification"]', 'Peut reproduire la même erreur', 'La répétition ne corrige pas une donnée source commune erronée.

![CPL_040_001](/illustrations/cpl-040-001.svg)', 'Repetition does not correct a shared wrong input.

![CPL_040_001](/illustrations/en/cpl-040-001.svg)', '["CPL", "exam-mcq", "exam-subject:040"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('22e768c1-0000-4000-8000-000022e768c1', '1e25397b-0000-4000-8000-00001e25397b', 'CPL', 'human_performance', 'human_performance', 'Installer des barrières', 3, 3, 'Après une interruption de checklist, la meilleure réponse est…', 'After a checklist interruption, the best response is…', '["Supposer tous les items terminés", "Sauter directement la checklist suivante", "Réciter plus vite sans regarder", "Reprendre selon la procédure et vérifier les points incertains"]', '["Assume every item is complete", "Skip directly to the next checklist", "Recite faster without looking", "Resume according to procedure and check uncertain items"]', 'Reprendre selon la procédure et vérifier les points incertains', 'La reprise doit éviter omissions et fausses confirmations.

![CPL_040_001](/illustrations/cpl-040-001.svg)', 'Resumption should avoid omissions and false confirmations.

![CPL_040_001](/illustrations/en/cpl-040-001.svg)', '["CPL", "exam-mcq", "exam-subject:040"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('22e768c2-0000-4000-8000-000022e768c2', '1e25397b-0000-4000-8000-00001e25397b', 'CPL', 'human_performance', 'human_performance', 'Installer des barrières', 3, 3, 'Une checklist récitée sans regarder les états réels risque…', 'A checklist recited without checking actual states risks…', '["De remplacer toute surveillance", "D’éliminer les erreurs de mémoire", "De confirmer à tort une action non réalisée", "De garantir une meilleure indépendance"]', '["Replacing all monitoring", "Eliminating memory errors", "Falsely confirming an action not performed", "Guaranteeing better independence"]', 'De confirmer à tort une action non réalisée', 'Le contrôle doit porter sur l’état attendu, pas seulement sur les mots.

![CPL_040_001](/illustrations/cpl-040-001.svg)', 'Verification must address expected state, not just words.

![CPL_040_001](/illustrations/en/cpl-040-001.svg)', '["CPL", "exam-mcq", "exam-subject:040"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('22e76c84-0000-4000-8000-000022e76c84', '1e25397b-0000-4000-8000-00001e25397b', 'CPL', 'human_performance', 'human_performance', 'Reconnaître les biais', 3, 3, 'Retenir les seuls bulletins favorables au départ illustre…', 'Selecting only reports favourable to departure illustrates…', '["Une limite structurale", "Le biais de confirmation", "Une vérification indépendante", "Une perte de pression"]', '["A structural limit", "Confirmation bias", "Independent verification", "A pressure loss"]', 'Le biais de confirmation', 'Les indices contraires sont écartés au profit de l’hypothèse initiale.

![CPL_040_001](/illustrations/cpl-040-001.svg)', 'Contrary evidence is discarded in favour of the initial hypothesis.

![CPL_040_001](/illustrations/en/cpl-040-001.svg)', '["CPL", "exam-mcq", "exam-subject:040"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('22e76c85-0000-4000-8000-000022e76c85', '1e25397b-0000-4000-8000-00001e25397b', 'CPL', 'human_performance', 'human_performance', 'Reconnaître les biais', 3, 3, 'Continuer malgré un seuil d’interruption atteint parce que le voyage est presque terminé illustre…', 'Continuing after a discontinuation threshold is reached because the trip is almost over illustrates…', '["Le biais de poursuite du plan", "Une application correcte du seuil", "Un nouveau minimum réglementaire", "Une réduction certaine du risque"]', '["Plan continuation bias", "Correct application of the threshold", "A new regulatory minimum", "A certain risk reduction"]', 'Le biais de poursuite du plan', 'La proximité du but ne rétablit pas les marges perdues.

![CPL_040_001](/illustrations/cpl-040-001.svg)', 'Proximity to the goal does not restore lost margins.

![CPL_040_001](/illustrations/en/cpl-040-001.svg)', '["CPL", "exam-mcq", "exam-subject:040"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('22e76c86-0000-4000-8000-000022e76c86', '1e25397b-0000-4000-8000-00001e25397b', 'CPL', 'human_performance', 'human_performance', 'Reconnaître les biais', 3, 3, 'Quelle action aide à combattre le biais de confirmation ?', 'Which action helps counter confirmation bias?', '["Écarter les avis divergents", "Réutiliser seulement les premiers chiffres", "Éviter toute réévaluation", "Rechercher des éléments contredisant l’hypothèse"]', '["Discard conflicting views", "Reuse only the first figures", "Avoid all reassessment", "Seek evidence contradicting the hypothesis"]', 'Rechercher des éléments contredisant l’hypothèse', 'Une recherche contradictoire teste réellement l’hypothèse.

![CPL_040_001](/illustrations/cpl-040-001.svg)', 'Seeking contrary evidence genuinely tests the hypothesis.

![CPL_040_001](/illustrations/en/cpl-040-001.svg)', '["CPL", "exam-mcq", "exam-subject:040"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('3a062372-0000-4000-8000-00003a062372', '1e25397b-0000-4000-8000-00001e25397b', 'CPL', 'human_performance', 'human_performance', 'Reconnaître les biais', 3, 3, 'Après une sélection au pilote automatique, il faut vérifier…', 'After an autopilot selection, verify…', '["Seulement la destination", "Uniquement le temps de vol passé", "Le mode actif et la trajectoire obtenue", "Uniquement le bouton pressé"]', '["Only the destination", "Only elapsed flight time", "The active mode and resulting flight path", "Only the button pressed"]', 'Le mode actif et la trajectoire obtenue', 'L’intention et le comportement réel peuvent différer.

![CPL_040_001](/illustrations/cpl-040-001.svg)', 'Intended and actual behaviour may differ.

![CPL_040_001](/illustrations/en/cpl-040-001.svg)', '["CPL", "exam-mcq", "exam-subject:040"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('3a0697d2-0000-4000-8000-00003a0697d2', '1e25397b-0000-4000-8000-00001e25397b', 'CPL', 'human_performance', 'human_performance', 'Gérer les ressources disponibles', 3, 3, 'Un pilote fatigué mais motivé peut…', 'A fatigued but motivated pilot may…', '["Se dispenser de réévaluation", "Présenter un jugement et une vigilance dégradés", "Être nécessairement apte", "Annuler son manque de sommeil par volonté"]', '["Skip reassessment", "Have impaired judgement and vigilance", "Necessarily be fit", "Cancel sleep loss through willpower"]', 'Présenter un jugement et une vigilance dégradés', 'La motivation ne supprime pas les effets de la fatigue.

![CPL_040_001](/illustrations/cpl-040-001.svg)', 'Motivation does not remove fatigue effects.

![CPL_040_001](/illustrations/en/cpl-040-001.svg)', '["CPL", "exam-mcq", "exam-subject:040"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('3a0697d3-0000-4000-8000-00003a0697d3', '1e25397b-0000-4000-8000-00001e25397b', 'CPL', 'human_performance', 'human_performance', 'Gérer les ressources disponibles', 3, 3, 'La caféine remplace-t-elle un repos approprié ?', 'Does caffeine replace appropriate rest?', '["Non", "Oui, pour toute durée", "Oui, après deux boissons", "Oui, uniquement en IFR"]', '["No", "Yes, for any duration", "Yes, after two drinks", "Yes, in IFR only"]', 'Non', 'Elle ne restaure pas à elle seule les fonctions affectées par le manque de repos.

![CPL_040_001](/illustrations/cpl-040-001.svg)', 'It does not by itself restore functions impaired by lack of rest.

![CPL_040_001](/illustrations/en/cpl-040-001.svg)', '["CPL", "exam-mcq", "exam-subject:040"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('3a0697d4-0000-4000-8000-00003a0697d4', '1e25397b-0000-4000-8000-00001e25397b', 'CPL', 'human_performance', 'human_performance', 'Gérer les ressources disponibles', 3, 3, 'La charge de travail dépasse vos capacités immédiates. Action adaptée ?', 'Workload exceeds your immediate capacity. Suitable action?', '["Finir d’abord les tâches administratives", "Augmenter toutes les tâches simultanément", "Cesser de surveiller la trajectoire", "Prioriser le pilotage et différer les tâches secondaires"]', '["Finish administrative tasks first", "Increase every simultaneous task", "Stop monitoring the flight path", "Prioritise aircraft control and defer secondary tasks"]', 'Prioriser le pilotage et différer les tâches secondaires', 'Réduire la charge permet de préserver les fonctions essentielles.

![CPL_040_001](/illustrations/cpl-040-001.svg)', 'Reducing workload preserves essential functions.

![CPL_040_001](/illustrations/en/cpl-040-001.svg)', '["CPL", "exam-mcq", "exam-subject:040"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('3a0697d5-0000-4000-8000-00003a0697d5', '1e25397b-0000-4000-8000-00001e25397b', 'CPL', 'human_performance', 'human_performance', 'Gérer les ressources disponibles', 3, 3, 'Une intervention assertive efficace consiste à…', 'An effective assertive intervention consists of…', '["Attaquer la personne plutôt que le problème", "Attendre que le problème disparaisse", "Décrire le problème et proposer une action claire", "Rester silencieux malgré le danger"]', '["Attack the person rather than the problem", "Wait for the problem to disappear", "State the problem and propose a clear action", "Remain silent despite danger"]', 'Décrire le problème et proposer une action claire', 'Une formulation concrète facilite une réponse opérationnelle.

![CPL_040_001](/illustrations/cpl-040-001.svg)', 'Concrete wording supports an operational response.

![CPL_040_001](/illustrations/en/cpl-040-001.svg)', '["CPL", "exam-mcq", "exam-subject:040"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO lessons (id, license, category, theme, title_fr, title_en, content_fr, content_en, difficulty, order_index, level, duration_minutes, tags, learning_objectives) VALUES
('1ad7b73f-0000-4000-8000-00001ad7b73f', 'CPL', 'instrumentation', 'instrumentation', 'Circuit pitot-statique : interpréter une panne', 'Pitot-static system: interpret a failure', '**Objectif :** Comparer pressions reçues et attendues, expliquer le signe d’une erreur après inversion du mouvement et identifier les dépendances d’une architecture de capteurs nouvelle.

## Identifier les pressions

![CPL_022_001](/illustrations/cpl-022-001.svg)

Dans le modèle conventionnel, l’anémomètre compare pression totale Pt et pression statique Ps. La différence Pt − Ps correspond à la pression dynamique dans l’approximation incompressible. L’altimètre utilise Ps ; le variomètre compare son évolution dans le temps par un dispositif calibré. Un défaut de prise statique peut donc affecter les trois indications. Le schéma fonctionnel explique les dépendances, sans représenter un montage particulier.

### Construire le raisonnement

Avant de mémoriser une flèche « monte » ou « descend », écrire les deux pressions reçues par l’instrument. Distinguer la pression extérieure actuelle de celle qui arrive effectivement par la conduite. Dans nos modèles idéaux, il n’y a ni fuite, ni variation thermique des volumes piégés ; le calage altimétrique reste fixe sauf indication contraire.

**Exemple résolu :** Pt = 920 hPa et Ps = 890 hPa donnent 30 hPa de différence. Si seule Ps reçue diminue à 885 hPa, la différence devient 35 hPa. L’anémomètre indique davantage, mais ce changement ne démontre aucune accélération. On a modifié sa référence. Les hPa expriment une pression, pas une vitesse : ne pas transformer une différence de 5 hPa en 5 kt.

L’altimètre associe une pression à une indication selon son calage. Tourner le bouton de calage peut donc changer cette indication même si la pression reçue reste piégée. Cette action ne rétablit pas la mesure de la pression extérieure.

## Analyser une obstruction pitot

Si l’entrée pitot est obstruée mais son drainage reste libre, la pression totale ne se maintient plus et l’indication de vitesse tend vers zéro dans le modèle idéal. Si entrée et drainage sont obstrués, la pression est piégée. Avec une prise statique libre, monter fait baisser Ps : l’anémomètre tend alors à indiquer davantage ; descendre produit la tendance inverse. Cela ne prouve pas une variation de vitesse réelle. Les effets exacts et alertes dépendent de l’installation.

### Deux obstructions, deux modèles différents

Le drainage est décisif. Entrée bouchée et drainage libre : la pression dans la conduite tend vers la pression ambiante dans le modèle simplifié ; la différence utile tend vers zéro. Entrée et drainage bouchés : Pt reste piégée, mais Ps peut encore évoluer. « Pitot bouché » sans état du drainage ne suffit donc pas à choisir le modèle.

**Exemple résolu :** Pt piégée à 900 hPa, Ps initiale 870 hPa : différence 30 hPa. En montant, Ps tombe à 855 hPa : différence 45 hPa. Lors d’une descente à Ps = 880 hPa, elle vaut 20 hPa. La même pression totale piégée produit donc plusieurs indications.

**Contre-exemple à retenir :** si Pt ET Ps sont toutes deux piégées, leur différence reste constante dans ce modèle. L’IAS peut paraître stable alors que la vitesse réelle varie. Une indication immobile peut être une panne, pas une preuve de vol stabilisé.

## Analyser une obstruction statique

Lorsque la pression statique est piégée, l’altimètre reste près de l’altitude d’obstruction et le variomètre revient vers zéro après le transitoire. Avec pitot libre, l’anémomètre sous-indique en montant au-dessus de cette altitude et surindique en descendant au-dessous, par rapport à son indication avec une statique correcte dans les mêmes conditions. Une source statique de secours peut introduire des erreurs : appliquer les corrections du manuel.

### Comparer deux indications au même instant

Avec pitot libre, noter Pt actuelle, Ps extérieure et Ps piégée. La bonne différence serait Pt − Ps extérieure ; l’instrument défectueux reçoit Pt − Ps piégée. Comparer ces deux résultats, sans supposer que montée signifie toujours hausse de l’indication ou descente toujours surindication.

**Exemple résolu :** statique piégée à 840 hPa. Après une montée puis une descente partielle, Ps extérieure vaut encore 820 hPa ; Pt vaut 860 hPa. Différence correcte 40 hPa, différence reçue 20 hPa : l’instrument sous-indique encore. L’avion descend, mais il reste au-dessus de l’altitude d’obstruction. La position par rapport à cette altitude explique le signe de l’erreur.

### Lire les indices ensemble

À calage constant, altimètre figé et variomètre revenu à zéro après transitoire sont compatibles avec une statique piégée. Ce couple d’indices ne démontre pas un palier réel. Une source de secours change la référence disponible ; elle n’est pas nécessairement sans erreur. Employer les corrections du manuel et distinguer le transitoire de la nouvelle indication stabilisée.

**Auto-évaluation :** si une source de secours délivre 3 hPa de moins que la statique correcte, avec Pt inchangée, la différence mesurée augmente-t-elle ou diminue-t-elle ? Elle augmente de 3 hPa ; le signe se retrouve directement dans la soustraction.

## Recouper avant d’agir

Deux affichages utilisant le même capteur ne sont pas deux mesures indépendantes. Recouper attitude, puissance, sources réellement indépendantes et trajectoire, puis appliquer la procédure de vitesse non fiable propre à l’avion. La vitesse sol GNSS ne remplace pas la vitesse air, notamment avec du vent. Le chauffage pitot limite certains risques de givrage mais ne corrige ni toutes les obstructions ni toutes les pannes. Éviter de poursuivre une indication incohérente par des actions brusques.

### Vérifier l’indépendance en remontant la chaîne

Deux calculateurs différents peuvent recevoir la même prise statique. Leur accord exclut mal une panne commune en amont. Dessiner la chaîne prise → conduite → capteur/calculateur → écran. Une alimentation électrique indépendante ne prouve pas l’indépendance pneumatique ; il faut vérifier les dépendances réellement pertinentes pour la panne envisagée.

**Cas guidé :** deux affichages s’accordent, un troisième diverge. Si les deux premiers partagent leur source, compter les votes ne permet pas d’identifier le bon. Rechercher une référence dont la chaîne utile est indépendante et recouper attitude, puissance et trajectoire selon la procédure de l’avion. Une source indépendante fournit un contrôle supplémentaire ; elle n’est pas déclarée infaillible par principe.

Le GNSS fournit une vitesse par rapport au sol. L’écart avec l’IAS contient notamment les différences entre IAS et TAS ainsi que le vent. Une GS de 110 kt et une IAS de 90 kt ne prouvent donc pas une erreur instrumentale de 20 kt. Le recoupement cherche une cohérence d’ensemble, sans substituer arbitrairement une grandeur à une autre.

**À expliquer sans le cours :** quelle pression est reçue, laquelle devrait l’être, quels instruments partagent la source et quelle conclusion reste incertaine ? En vol, préserver la conduite de l’avion et suivre la procédure appropriée prime sur la résolution complète de l’exercice.

### Transférer la méthode à une architecture nouvelle
Une installation fictive peut partager la pression totale sans partager la statique, ou l''inverse. Exemple guidé distinct du schéma de QCM : D reçoit Pt1/Ps1, E reçoit Pt2/Ps1 et F reçoit Pt2/Ps2. Dessiner séparément les liens Pt et Ps : D/E partagent Ps1, E/F partagent Pt2. Une panne de Ps1 concerne donc D/E dans ce modèle. Le QCM propose une autre architecture A/B/C : reprendre le tracé, sans transférer les couples du cas guidé.

**Méthode de diagnostic limitée :** lister les affichages qu''une panne unique pourrait affecter ; comparer cette liste avec les observations. Un affichage stable peut être peu sensible à une panne dans les conditions du moment : il faut aussi considérer les pressions réellement en évolution. Un groupe de deux affichages concordants n''identifie pas, à lui seul, la chaîne correcte. Les exercices précisent les hypothèses de panne unique et de conduite intacte pour rendre la conclusion déterminée ; une installation réelle exige son propre schéma et sa procédure.

**Format de travail :** 15 QCM de chapitre, quatre choix et une réponse correcte. La matière 022 comporte 40 questions en 60 minutes dans le tableau EASA cité. Ce chapitre ne couvre qu’une partie de la matière ; il ne constitue pas une épreuve complète. Cadre EASA, même en anglais ou avec les unités US/britanniques.

## Références

- [EASA — programme Part-FCL](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8)
- [EASA — format de la matière, AMC1 ARA.FCL.300(b)](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67)
- [FAA PHAK — flight instruments, référence technique](https://www.faa.gov/sites/faa.gov/files/regulations_policies/handbooks_manuals/aviation/phak/10_phak_ch8.pdf)

**Support de révision CPL(A).** À faire valider par votre instructeur. Pour un vol réel, utiliser le manuel de l’avion et les publications en vigueur. Les exemples chiffrés sont pédagogiques.', '**Objective:** Compare received and expected pressures, explain error sign after reversing vertical movement and identify dependencies in a new sensor architecture.

## Identify the pressures

![CPL_022_001](/illustrations/en/cpl-022-001.svg)

In the conventional model, the airspeed indicator compares total pressure Pt with static pressure Ps. Pt − Ps corresponds to dynamic pressure in the incompressible approximation. The altimeter uses Ps; the vertical-speed indicator senses its change over time through a calibrated arrangement. A static-source fault may therefore affect all three indications. The functional diagram explains dependencies, not a particular installation.

### Build the reasoning

Before memorising an “up” or “down” arrow, write down both pressures received by the instrument. Separate current external pressure from pressure actually delivered through the line. Our ideal models assume no leaks or thermal changes in trapped volumes; the altimeter setting remains fixed unless specified otherwise.

**Worked example:** Pt = 920 hPa and Ps = 890 hPa give a 30 hPa difference. If only received Ps falls to 885 hPa, the difference becomes 35 hPa. The airspeed indication increases, but this establishes no acceleration: its reference changed. hPa measures pressure, not speed; a 5 hPa difference is not 5 kt.

An altimeter relates pressure to an indication according to its setting. Turning the setting knob may therefore change the indication even with trapped pressure. That action does not restore measurement of external pressure.

## Analyse a pitot blockage

If the pitot inlet is blocked but its drain remains open, total pressure is no longer maintained and indicated airspeed tends towards zero in the ideal model. If both inlet and drain are blocked, pressure is trapped. With a clear static source, climbing reduces Ps and the airspeed indication tends to increase; descending produces the opposite tendency. This does not demonstrate a real speed change. Exact effects and alerts depend on the installation.

### Two blockages, two different models

The drain is decisive. Blocked inlet with an open drain: line pressure tends towards ambient pressure in the simplified model, so the useful difference tends towards zero. Blocked inlet and drain: Pt remains trapped, but Ps can still change. “Blocked pitot” without drain status is therefore insufficient to choose the model.

**Worked example:** Pt trapped at 900 hPa, initial Ps 870 hPa: difference 30 hPa. During a climb, Ps falls to 855 hPa: difference 45 hPa. During descent to Ps = 880 hPa, it is 20 hPa. The same trapped total pressure therefore produces several indications.

**Counterexample to remember:** if BOTH Pt and Ps are trapped, their difference remains constant in this model. IAS may appear stable while actual speed changes. A stationary indication can be a fault, not evidence of steady flight.

## Analyse a static blockage

When static pressure is trapped, the altimeter remains near the blockage altitude and the vertical-speed indicator returns towards zero after the transient. With a clear pitot, airspeed underreads when climbing above that altitude and overreads when descending below it, compared with the indication using correct static pressure under the same conditions. An alternate static source may introduce errors: apply manual corrections.

### Compare two indications at the same instant

With an open pitot, write current Pt, external Ps and trapped Ps. The correct difference would be Pt minus external Ps; the faulty instrument receives Pt minus trapped Ps. Compare these results without assuming that climb always means a rising indication or descent always means overreading.

**Worked example:** static pressure trapped at 840 hPa. After climbing then partially descending, external Ps is still 820 hPa; Pt is 860 hPa. Correct difference 40 hPa, received difference 20 hPa: the instrument still underreads. The aircraft is descending but remains above the blockage altitude. Its position relative to that altitude explains the sign of the error.

### Read the cues together

At a constant setting, a frozen altimeter and a VSI returning to zero after the transient are consistent with trapped static pressure. Together, they do not establish actual level flight. An alternate source changes the available reference; it need not be error-free. Use manual corrections and distinguish the transient from the new steady indication.

**Self-check:** if an alternate source supplies 3 hPa less than correct static pressure with unchanged Pt, does the measured difference rise or fall? It rises by 3 hPa; the sign follows directly from subtraction.

## Cross-check before acting

Two displays using the same sensor are not two independent measurements. Cross-check attitude, power, genuinely independent sources and flight path, then follow the aircraft-specific unreliable-airspeed procedure. GNSS groundspeed does not replace airspeed, particularly in wind. Pitot heat reduces some icing risks but does not correct every blockage or failure. Avoid chasing inconsistent indications with abrupt control inputs.

### Trace the chain to establish independence

Two different computers may receive the same static source. Their agreement does little to exclude an upstream common fault. Draw the chain: port → line → sensor/computer → display. Independent electrical power does not establish pneumatic independence; check the dependencies relevant to the suspected failure.

**Worked case:** two displays agree and a third differs. If the first two share a source, counting votes cannot identify the correct one. Seek a reference with an independent relevant chain and cross-check attitude, power and flight path under aircraft procedures. An independent source provides another check; independence does not make it infallible.

GNSS supplies speed relative to the ground. Its difference from IAS includes IAS-to-TAS differences and wind. GS 110 kt and IAS 90 kt therefore do not establish a 20 kt instrument error. Cross-checking seeks overall consistency without arbitrarily substituting one quantity for another.

**Recall without the lesson:** which pressure is received, which should be received, which instruments share the source, and what remains uncertain? In flight, maintaining aircraft control and following the appropriate procedure takes priority over fully solving an exercise.

### Transfer the method to a new architecture
A fictional installation may share total pressure without sharing static pressure, or the reverse. Worked example distinct from the quiz diagram: D receives Pt1/Ps1, E receives Pt2/Ps1 and F receives Pt2/Ps2. Trace Pt and Ps separately: D/E share Ps1; E/F share Pt2. A Ps1 fault therefore affects D/E in this model. The quiz uses a different A/B/C architecture: trace it again instead of transferring the worked example''s pairs.

**Limited diagnostic method:** list displays a single fault could affect, then compare with observations. A stable display may be insensitive to a fault in current conditions: consider which pressures are actually changing. Two agreeing displays alone do not identify the correct chain. Exercises specify single-fault and intact-line assumptions to make the answer determined; real installations require their own diagram and procedure.

**Practice format:** 15 chapter MCQs, four choices and one correct answer. Subject 022 has 40 questions in 60 minutes in the cited EASA table. This chapter covers only part of the subject and is not a full paper. EASA framework, including with English or US/British units.

## References

- [Reference 1](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8)
- [Reference 2](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67)
- [Reference 3](https://www.faa.gov/sites/faa.gov/files/regulations_policies/handbooks_manuals/aviation/phak/10_phak_ch8.pdf)

**CPL(A) revision material.** To be reviewed by your instructor. For an actual flight, use the aircraft manual and current publications. Numerical examples are for training.', 3, 1, 1, 50, '["CPL", "mcq-four-options"]', '["Comparer pressions reçues et attendues, expliquer le signe d’une erreur après inversion du mouvement et identifier les dépendances d’une architecture de capteurs nouvelle."]')
ON CONFLICT (id) DO UPDATE SET license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, title_fr = EXCLUDED.title_fr, title_en = EXCLUDED.title_en, content_fr = EXCLUDED.content_fr, content_en = EXCLUDED.content_en, difficulty = EXCLUDED.difficulty, order_index = EXCLUDED.order_index, level = EXCLUDED.level, duration_minutes = EXCLUDED.duration_minutes, tags = EXCLUDED.tags, learning_objectives = EXCLUDED.learning_objectives;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('59ad9483-0000-4000-8000-000059ad9483', '1ad7b73f-0000-4000-8000-00001ad7b73f', 'CPL', 'instrumentation', 'instrumentation', 'Identifier les pressions', 3, 3, 'Dans le modèle conventionnel, une prise statique commune est obstruée. Quel groupe peut être affecté directement par cette seule panne ?', 'In the conventional model, a shared static port is blocked. Which group can be directly affected by this single fault?', '["Anémomètre, altimètre et variomètre", "Altimètre et variomètre seulement, car le pitot alimente l’anémomètre", "Anémomètre seulement, car les autres instruments ne comparent pas Pt et Ps", "Altimètre seulement, car le variomètre mesure directement le mouvement vertical"]', '["Airspeed indicator, altimeter and VSI", "Altimeter and VSI only, because pitot supplies the airspeed indicator", "Airspeed indicator only, because the others do not compare Pt and Ps", "Altimeter only, because the VSI directly measures vertical movement"]', 'Anémomètre, altimètre et variomètre', 'Les trois utilisent Ps, même s’ils la traitent différemment : différence avec Pt, conversion en altitude, variation dans le temps. Le pitot libre ne suffit pas à fiabiliser l’anémomètre ; le variomètre ne mesure pas directement une distance parcourue verticalement.

![CPL_022_001](/illustrations/cpl-022-001.svg)', 'All three use Ps differently: difference from Pt, conversion to altitude, and change over time. An open pitot is insufficient to ensure reliable airspeed; a VSI does not directly measure distance travelled vertically.

![CPL_022_001](/illustrations/en/cpl-022-001.svg)', '["CPL", "exam-mcq", "exam-subject:022"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('59ad9484-0000-4000-8000-000059ad9484', '1ad7b73f-0000-4000-8000-00001ad7b73f', 'CPL', 'instrumentation', 'instrumentation', 'Identifier les pressions', 3, 3, 'Pt = 850 hPa, Ps = 830 hPa. Quelle différence de pression reçoit l’anémomètre conventionnel ?', 'Pt = 850 hPa, Ps = 830 hPa. What pressure difference does the conventional airspeed indicator receive?', '["1 680 hPa", "830 hPa", "−20 hPa", "20 hPa"]', '["1,680 hPa", "830 hPa", "−20 hPa", "20 hPa"]', '20 hPa', 'Soustraire Ps de Pt : 850 − 830 = 20 hPa. Additionner donne 1 680 ; inverser l’ordre donne −20 ; 830 est la référence statique seule. La réponse est une pression et ne fournit pas directement une vitesse en kt.

![CPL_022_001](/illustrations/cpl-022-001.svg)', 'Subtract Ps from Pt: 850 − 830 = 20 hPa. Adding gives 1,680; reversing the order gives −20; 830 is static pressure alone. The result is pressure and does not directly give a speed in kt.

![CPL_022_001](/illustrations/en/cpl-022-001.svg)', '["CPL", "exam-mcq", "exam-subject:022"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('59ad9485-0000-4000-8000-000059ad9485', '1ad7b73f-0000-4000-8000-00001ad7b73f', 'CPL', 'instrumentation', 'instrumentation', 'Identifier les pressions', 3, 3, 'La statique d’un altimètre est piégée ; le pilote change son calage et l’indication bouge. Que démontre ce mouvement ?', 'An altimeter’s static pressure is trapped; the pilot changes its setting and the indication moves. What does that movement establish?', '["L’avion a changé d’altitude de la valeur affichée", "La correction de température a été automatiquement appliquée", "Le mécanisme de calage agit, sans prouver que la prise est redevenue libre", "La pression extérieure atteint de nouveau l’instrument"]', '["The aircraft changed altitude by the displayed amount", "Temperature correction has been applied automatically", "The setting mechanism responds, without proving the port has cleared", "External pressure is reaching the instrument again"]', 'Le mécanisme de calage agit, sans prouver que la prise est redevenue libre', 'L’indication dépend de la pression reçue ET du calage. Modifier le second peut déplacer l’aiguille sans modifier la première. Ce test ne prouve ni débouchage, ni déplacement réel, ni correction des erreurs de température.

![CPL_022_001](/illustrations/cpl-022-001.svg)', 'The indication depends on both received pressure and setting. Changing the latter can move the pointer without changing the former. This proves neither an unblocked port, actual movement nor temperature-error correction.

![CPL_022_001](/illustrations/en/cpl-022-001.svg)', '["CPL", "exam-mcq", "exam-subject:022"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('59ad9847-0000-4000-8000-000059ad9847', '1ad7b73f-0000-4000-8000-00001ad7b73f', 'CPL', 'instrumentation', 'instrumentation', 'Analyser une obstruction pitot', 3, 3, 'Entrée pitot bouchée, drainage et statique libres ; modèle idéal stabilisé. Quelle tendance d’IAS correspond à ces seules conditions ?', 'Blocked pitot inlet, open drain and static source; ideal steady model. Which IAS tendency follows from these conditions alone?', '["Valeur correcte, car la prise statique reste libre", "Vers zéro, car la différence utile se résorbe", "Valeur figée, car toute obstruction piège nécessairement Pt", "Hausse en montée, car Pt reste forcément piégée"]', '["Correct value because the static source remains open", "Towards zero as the useful difference dissipates", "Frozen value because every blockage necessarily traps Pt", "Rise in a climb because Pt must remain trapped"]', 'Vers zéro, car la différence utile se résorbe', 'Le drainage libre empêche ici le maintien de Pt piégée. La valeur figée ou la réaction d’un pitot fermé supposent une autre configuration. Une statique correcte ne rétablit pas une pression totale devenue indisponible.

![CPL_022_001](/illustrations/cpl-pitot-reasoning.svg)', 'The open drain prevents trapped Pt being maintained here. A frozen value or closed-pitot response assumes a different configuration. Correct static pressure does not restore unavailable total pressure.

![CPL_022_001](/illustrations/en/cpl-pitot-reasoning.svg)', '["CPL", "exam-mcq", "exam-subject:022"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('59ad9848-0000-4000-8000-000059ad9848', '1ad7b73f-0000-4000-8000-00001ad7b73f', 'CPL', 'instrumentation', 'instrumentation', 'Analyser une obstruction pitot', 3, 3, 'Pt est piégée à 800 hPa, entrée et drainage pitot bouchés ; Ps libre passe de 780 à 760 hPa en montée. Quelle variation de différence reçoit l’anémomètre ?', 'Pt is trapped at 800 hPa, pitot inlet and drain blocked; open Ps falls from 780 to 760 hPa in a climb. How does the airspeed indicator’s pressure difference change?', '["De 20 à 40 hPa", "De 20 à 0 hPa", "Elle reste à 20 hPa", "De 780 à 760 hPa"]', '["From 20 to 40 hPa", "From 20 to 0 hPa", "It remains at 20 hPa", "From 780 to 760 hPa"]', 'De 20 à 40 hPa', 'Au départ, 800 − 780 = 20 ; ensuite, 800 − 760 = 40 hPa. L’IAS tend à augmenter car Ps diminue, malgré Pt constante. Les valeurs 780 et 760 sont des pressions statiques, pas les différences mesurées.

![CPL_022_001](/illustrations/cpl-pitot-reasoning.svg)', 'Initially 800 − 780 = 20; then 800 − 760 = 40 hPa. IAS tends to increase as Ps falls despite constant Pt. The values 780 and 760 are static pressures, not measured differences.

![CPL_022_001](/illustrations/en/cpl-pitot-reasoning.svg)', '["CPL", "exam-mcq", "exam-subject:022"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('59ad9849-0000-4000-8000-000059ad9849', '1ad7b73f-0000-4000-8000-00001ad7b73f', 'CPL', 'instrumentation', 'instrumentation', 'Analyser une obstruction pitot', 3, 3, 'Pt piégée à 800 hPa, statique libre : Ps passe de 760 à 790 hPa pendant une descente. Quelle différence finale et quelle tendance d’IAS ?', 'Pt trapped at 800 hPa, open static source: Ps rises from 760 to 790 hPa during descent. What final difference and IAS trend result?', '["10 hPa ; indication en hausse", "40 hPa ; indication constante", "30 hPa ; indication en baisse", "10 hPa ; indication en baisse"]', '["10 hPa; indication increases", "40 hPa; indication remains constant", "30 hPa; indication decreases", "10 hPa; indication decreases"]', '10 hPa ; indication en baisse', 'La différence passe de 800 − 760 = 40 à 800 − 790 = 10 hPa. Les 30 hPa représentent sa diminution, pas sa valeur finale. Cette variation instrumentale ne prouve pas un ralentissement réel.

![CPL_022_001](/illustrations/cpl-pitot-reasoning.svg)', 'The difference changes from 800 − 760 = 40 to 800 − 790 = 10 hPa. The 30 hPa is its decrease, not its final value. This instrument change does not establish actual deceleration.

![CPL_022_001](/illustrations/en/cpl-pitot-reasoning.svg)', '["CPL", "exam-mcq", "exam-subject:022"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('59ad984a-0000-4000-8000-000059ad984a', '1ad7b73f-0000-4000-8000-00001ad7b73f', 'CPL', 'instrumentation', 'instrumentation', 'Analyser une obstruction pitot', 3, 3, 'Pt et Ps sont toutes deux piégées à 800 et 780 hPa. Modèle idéal sans fuite ni effet thermique : si la vitesse réelle change, que devient la différence reçue ?', 'Both Pt and Ps are trapped at 800 and 780 hPa. Ideal model without leaks or thermal effects: if actual speed changes, what happens to the received difference?', '["Elle devient nulle dès que l’altitude change", "Elle varie uniquement avec la vitesse sol GNSS", "Elle reste à 20 hPa", "Elle suit la nouvelle pression dynamique extérieure"]', '["It becomes zero as soon as altitude changes", "It varies only with GNSS groundspeed", "It remains at 20 hPa", "It follows the new external dynamic pressure"]', 'Elle reste à 20 hPa', 'Les deux valeurs reçues sont constantes : 800 − 780 = 20 hPa, même si les conditions extérieures évoluent. Ne pas réutiliser le modèle Pt piégée/Ps libre. Une IAS stable n’est donc pas une preuve de vitesse réelle constante.

![CPL_022_001](/illustrations/cpl-pitot-reasoning.svg)', 'Both received values are constant: 800 − 780 = 20 hPa even as outside conditions change. Do not reuse the trapped-Pt/open-Ps model. Stable IAS therefore does not prove constant actual speed.

![CPL_022_001](/illustrations/en/cpl-pitot-reasoning.svg)', '["CPL", "exam-mcq", "exam-subject:022"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('59ad9c0c-0000-4000-8000-000059ad9c0c', '1ad7b73f-0000-4000-8000-00001ad7b73f', 'CPL', 'instrumentation', 'instrumentation', 'Analyser une obstruction statique', 3, 3, 'Une prise statique se bouche à 5 000 ft indiqués. L’avion monte ensuite réellement ; calage inchangé, modèle idéal. Quelle indication altimétrique est attendue ?', 'A static port blocks at 5,000 ft indicated. The aircraft subsequently climbs; setting unchanged, ideal model. What altimeter indication is expected?', '["Un retour progressif vers zéro après transitoire", "Environ 5 000 ft", "Une augmentation suivant la montée réelle", "Une augmentation proportionnelle à Pt"]', '["A gradual return to zero after the transient", "Approximately 5,000 ft", "An increase following the actual climb", "An increase proportional to Pt"]', 'Environ 5 000 ft', 'L’altimètre conserve la pression statique correspondant à l’état d’obstruction, avec le même calage. Pt ne le pilote pas. Le retour vers zéro après transitoire décrit le variomètre dans ce modèle, pas l’altimètre.

![CPL_022_001](/illustrations/cpl-022-001.svg)', 'The altimeter retains static pressure from the blockage condition with the same setting. Pt does not drive it. Returning to zero after the transient describes the VSI in this model, not the altimeter.

![CPL_022_001](/illustrations/en/cpl-022-001.svg)', '["CPL", "exam-mcq", "exam-subject:022"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('59ad9c0d-0000-4000-8000-000059ad9c0d', '1ad7b73f-0000-4000-8000-00001ad7b73f', 'CPL', 'instrumentation', 'instrumentation', 'Analyser une obstruction statique', 3, 3, 'Statique piégée et montée réelle poursuivie : après stabilisation du transitoire interne, quelle indication du variomètre conventionnel est cohérente ?', 'Trapped static pressure while actual climb continues: after the internal transient settles, which conventional VSI indication is consistent?', '["Zéro malgré la montée réelle", "Le dernier taux de montée conservé indéfiniment", "Le taux de montée réel, car Pt continue d’évoluer", "Un taux de descente proportionnel à l’IAS"]', '["Zero despite the actual climb", "The last climb rate retained indefinitely", "Actual climb rate because Pt keeps changing", "A descent rate proportional to IAS"]', 'Zéro malgré la montée réelle', 'La pression piégée ne varie plus à l’entrée du système ; les pressions internes finissent par s’équilibrer. Le variomètre revient vers zéro sans démontrer un palier. Ni Pt ni une autre indication d’IAS ne remplacent l’entrée statique absente.

![CPL_022_001](/illustrations/cpl-022-001.svg)', 'Trapped pressure no longer changes at the system inlet; internal pressures eventually equalise. The VSI returns towards zero without establishing level flight. Neither Pt nor an IAS indication replaces the missing changing static input.

![CPL_022_001](/illustrations/en/cpl-022-001.svg)', '["CPL", "exam-mcq", "exam-subject:022"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('59ad9c0e-0000-4000-8000-000059ad9c0e', '1ad7b73f-0000-4000-8000-00001ad7b73f', 'CPL', 'instrumentation', 'instrumentation', 'Analyser une obstruction statique', 3, 3, 'Pitot libre : Pt actuelle 850 hPa, Ps extérieure 810 hPa, Ps piégée reçue 840 hPa. Quelle comparaison est correcte ?', 'Open pitot: current Pt 850 hPa, external Ps 810 hPa, received trapped Ps 840 hPa. Which comparison is correct?', '["40 hPa reçus contre 10 hPa corrects : surindication", "30 hPa reçus contre 40 hPa corrects : sous-indication", "10 hPa reçus contre 40 hPa corrects : surindication", "10 hPa reçus contre 40 hPa corrects : sous-indication"]', '["40 hPa received versus 10 hPa correct: overreading", "30 hPa received versus 40 hPa correct: underreading", "10 hPa received versus 40 hPa correct: overreading", "10 hPa received versus 40 hPa correct: underreading"]', '10 hPa reçus contre 40 hPa corrects : sous-indication', 'L’instrument reçoit 850 − 840 = 10 hPa ; une statique correcte donnerait 850 − 810 = 40 hPa. Soustraire une Ps trop forte diminue la différence. Les 30 hPa sont l’écart entre les deux résultats, pas la pression reçue.

![CPL_022_001](/illustrations/cpl-pitot-reasoning.svg)', 'The instrument receives 850 − 840 = 10 hPa; correct static pressure would give 850 − 810 = 40 hPa. Subtracting excessive Ps reduces the difference. The 30 hPa is the error between results, not the received difference.

![CPL_022_001](/illustrations/en/cpl-pitot-reasoning.svg)', '["CPL", "exam-mcq", "exam-subject:022"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('5c05e4ea-0000-4000-8000-00005c05e4ea', '1ad7b73f-0000-4000-8000-00001ad7b73f', 'CPL', 'instrumentation', 'instrumentation', 'Analyser une obstruction statique', 3, 3, 'Statique piégée à 4 000 ft. Après descente à 2 000 ft, l''avion remonte à 3 000 ft ; pitot libre, Ps décroît avec l''altitude. Quelle erreur IAS subsiste ?', 'Static trapped at 4,000 ft. After descending to 2,000 ft, the aircraft climbs to 3,000 ft; pitot open, Ps decreases with altitude. Which IAS error remains?', '["Aucune dès que la montée commence", "Indication figée car la statique est piégée", "Surindication, car l''avion reste sous l''altitude de blocage", "Sous-indication, car l''avion monte"]', '["None once the climb begins", "Frozen indication because static is trapped", "Overreading, because the aircraft remains below blockage altitude", "Underreading, because the aircraft is climbing"]', 'Surindication, car l''avion reste sous l''altitude de blocage', 'À 3 000 ft, Ps extérieure est supérieure à Ps piégée à 4 000 ft. Soustraire la valeur piégée plus faible augmente la différence reçue. Le sens de déplacement ne change pas ce signe tant que l''avion reste sous l''altitude de blocage.

![CPL_022_001](/illustrations/cpl-pitot-reasoning.svg)', 'At 3,000 ft, external Ps exceeds Ps trapped at 4,000 ft. Subtracting the smaller trapped value increases received difference. Direction of movement does not change this sign while still below blockage altitude.

![CPL_022_001](/illustrations/en/cpl-pitot-reasoning.svg)', '["CPL", "exam-mcq", "exam-subject:022"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('5c05e4eb-0000-4000-8000-00005c05e4eb', '1ad7b73f-0000-4000-8000-00001ad7b73f', 'CPL', 'instrumentation', 'instrumentation', 'Analyser une obstruction statique', 3, 3, 'Exercice : Pt = 900 hPa ; statique correcte = 870 hPa. Une source de secours fournit 865 hPa, Pt inchangée. Quelle nouvelle différence reçoit l’anémomètre ?', 'Exercise: Pt = 900 hPa; correct static = 870 hPa. An alternate source supplies 865 hPa with unchanged Pt. What new difference reaches the airspeed indicator?', '["5 hPa", "35 hPa", "25 hPa", "30 hPa"]', '["5 hPa", "35 hPa", "25 hPa", "30 hPa"]', '35 hPa', '900 − 865 = 35 hPa contre 30 auparavant. Une référence statique plus basse accroît ici la différence. Les 5 hPa sont l’erreur, pas la nouvelle différence. Ce scénario fixe le signe ; une installation réelle demande les corrections de son manuel.

![CPL_022_001](/illustrations/cpl-pitot-reasoning.svg)', '900 − 865 = 35 hPa versus 30 previously. A lower static reference increases the difference here. The 5 hPa is the error, not the new difference. This case specifies the sign; a real installation requires its manual corrections.

![CPL_022_001](/illustrations/en/cpl-pitot-reasoning.svg)', '["CPL", "exam-mcq", "exam-subject:022"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('5c06594b-0000-4000-8000-00005c06594b', '1ad7b73f-0000-4000-8000-00001ad7b73f', 'CPL', 'instrumentation', 'instrumentation', 'Recouper avant d’agir', 3, 3, 'Deux afficheurs possèdent des calculateurs et alimentations séparés mais partagent la même prise statique. Ils s’accordent après une obstruction de cette prise. Que vaut cet accord ?', 'Two displays have separate computers and power supplies but share one static port. They agree after that port becomes blocked. What does this agreement establish?', '["Il n’exclut pas une erreur commune venant de la prise", "Il valide Ps puisque les calculateurs sont indépendants", "Il valide Ps puisque les alimentations sont séparées", "Il permet de rejeter toute troisième indication divergente"]', '["It does not exclude a common error originating at the port", "It validates Ps because the computers are independent", "It validates Ps because power supplies are separate", "It allows any disagreeing third indication to be rejected"]', 'Il n’exclut pas une erreur commune venant de la prise', 'L’indépendance électrique ou de calcul ne supprime pas la dépendance pneumatique commune. Une erreur en amont peut être reproduite fidèlement par les deux chaînes. Le nombre d’écrans concordants ne suffit pas : examiner les sources.

![CPL_022_001](/illustrations/cpl-sensor-independence.svg)', 'Electrical or computing independence does not remove shared pneumatic dependence. Both chains may faithfully reproduce an upstream error. Counting agreeing displays is insufficient: examine their sources.

![CPL_022_001](/illustrations/en/cpl-sensor-independence.svg)', '["CPL", "exam-mcq", "exam-subject:022"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('5c06594c-0000-4000-8000-00005c06594c', '1ad7b73f-0000-4000-8000-00001ad7b73f', 'CPL', 'instrumentation', 'instrumentation', 'Recouper avant d’agir', 3, 3, 'En vol, GS 145 kt et IAS 120 kt. Un pilote attribue les 25 kt d''écart à une erreur instrumentale. Quelle étape manque à ce raisonnement ?', 'In flight, GS 145 kt and IAS 120 kt. A pilot attributes the 25 kt difference to instrument error. What step is missing?', '["Soustraire seulement la composante de vent traversier à l''IAS", "Considérer GS comme référence air car le GNSS est précis", "Convertir les 25 kt directement en hPa", "Relier IAS à TAS et intégrer le vent avant la comparaison"]', '["Subtract only crosswind component from IAS", "Treat GS as an air reference because GNSS is accurate", "Convert 25 kt directly into hPa", "Relate IAS to TAS and account for wind before comparison"]', 'Relier IAS à TAS et intégrer le vent avant la comparaison', 'Les deux valeurs n''ont pas la même référence. La précision GNSS ne change pas la nature de GS ; le vent seul ne résout pas la différence IAS/TAS. L''écart brut ne fournit pas une erreur ni une correction de pression.

![CPL_022_001](/illustrations/cpl-022-001.svg)', 'The values have different references. GNSS precision does not change the nature of GS; wind alone does not resolve IAS/TAS differences. The raw gap gives neither an error nor a pressure correction.

![CPL_022_001](/illustrations/en/cpl-022-001.svg)', '["CPL", "exam-mcq", "exam-subject:022"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('5c06594d-0000-4000-8000-00005c06594d', '1ad7b73f-0000-4000-8000-00001ad7b73f', 'CPL', 'instrumentation', 'instrumentation', 'Recouper avant d’agir', 3, 3, 'Schéma fictif : A reçoit Pt1/Ps1, B reçoit Pt1/Ps2, C reçoit Pt2/Ps2. Seule Ps2 se bloque, les autres sources et conduites restent intactes. En changement d''altitude, quel groupe peut partager une erreur de statique ?

![CPL_022_001](/illustrations/cpl-022-001.svg)', 'Fictional diagram: A receives Pt1/Ps1, B Pt1/Ps2, C Pt2/Ps2. Only Ps2 blocks; other sources and lines remain intact. During altitude change, which group can share a static error?

![CPL_022_001](/illustrations/en/cpl-022-001.svg)', '["A et C, parce que leurs calculateurs sont séparés", "Les trois, parce qu''ils sont à la même altitude réelle", "B et C, malgré leurs pitots distincts", "A et B, parce que leur pitot est commun"]', '["A and C because their computers are separate", "All three because actual altitude is the same", "B and C despite separate pitots", "A and B because they share a pitot"]', 'B et C, malgré leurs pitots distincts', 'Suivre le lien Ps2 mène à B et C. Pt1 partagé par A/B n''est pas la source défaillante de ce scénario. L''altitude réelle commune ne crée pas une connexion pneumatique ; A conserve ici Ps1 intacte.

![CPL_022_001](/illustrations/cpl-022-001.svg)', 'Tracing Ps2 leads to B and C. A/B''s shared Pt1 is not the failed source in this case. Common actual altitude creates no pneumatic connection; A retains intact Ps1 here.

![CPL_022_001](/illustrations/en/cpl-022-001.svg)', '["CPL", "exam-mcq", "exam-subject:022"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO lessons (id, license, category, theme, title_fr, title_en, content_fr, content_en, difficulty, order_index, level, duration_minutes, tags, learning_objectives) VALUES
('1c7e785d-0000-4000-8000-00001c7e785d', 'CPL', 'mass_and_balance', 'mass_and_balance', 'Chargement et déplacement du centrage', 'Loading and CG movement', '**Objectif :** Construire et contrôler un bilan de moments, distinguer déplacement et retrait, résoudre un déplacement inverse et convertir le centrage en pourcentage MAC.

## Établir la masse et le moment

![CPL_031_001](/illustrations/cpl-031-001.svg)

Utiliser un même plan de référence et des bras signés cohérents. Le moment vaut masse × bras ; le centrage vaut somme des moments / masse totale. Exemple fictif : avion 1 800 kg au bras 2,00 m, charge 300 kg au bras 3,00 m, carburant 200 kg au bras 2,50 m. Masse totale = 2 300 kg ; moment = 3 600 + 900 + 500 = 5 000 kg·m ; centrage = 2,174 m. Ne pas moyenner directement les bras. Les données du schéma sont celles de cet exemple, pas d’un modèle réel.

### Construire un bilan contrôlable

Pour cet exercice, les bras augmentent vers l’arrière. Une ligne de chargement doit toujours associer une masse, son bras depuis le datum commun et son moment. Ne pas mélanger bras depuis le datum et distances depuis le CG. La division (kg·m)/kg restitue un bras en mètres.

**Contrôle indépendant :** le CG 2,174 m est entre les bras extrêmes 2,00 et 3,00 m et proche du bras de l’avion initial, qui porte la plus grande masse. La moyenne simple des bras donnerait 2,50 m et surpondérerait les petites charges.

Si le datum est déplacé et que chaque bras augmente de 1 m, le moment total augmente de 2 300 kg·m et le CG exprimé dans la nouvelle référence augmente de 1 m. Le chargement physique n’a pourtant pas bougé : il faut également exprimer les limites dans la même référence. Une valeur numérique de bras n’a de sens qu’avec son datum.

## Déplacer une masse déjà embarquée

Sans ajout ni retrait de masse, variation du centrage = masse déplacée × déplacement du bras / masse totale. Déplacer 100 kg de 3,00 m à 1,00 m dans un avion de 2 300 kg change le moment de −200 kg·m et avance le centrage de 0,087 m. Le nouveau centrage de l’exemple est 4 800 / 2 300 = 2,087 m. Le signe indique le sens ; la masse totale reste inchangée. Vérifier aussi limites des compartiments, arrimage et chargement admissible.

### Retrouver la formule et la contrôler

Un déplacement interne retire le moment m × ancien bras puis ajoute m × nouveau bras. La différence vaut m × (nouveau − ancien). Puisque la masse totale M ne change pas, ΔCG = Δmoment/M. Ce raisonnement explique le signe ; il évite de mémoriser une formule sans savoir quand l’utiliser.

**Contrôle par deux méthodes :** dans l’exemple, nouveau moment 5 000 − 100 × 3 + 100 × 1 = 4 800 kg·m, donc CG 2,086956… m. Avec les valeurs non arrondies, ancien CG 5 000/2 300 diminué de 200/2 300 donne exactement le même résultat.

**Problème inverse :** pour avancer le CG de 0,05 m dans un avion de 2 300 kg, en déplaçant une charge de 2 m vers l’avant, il faut déplacer 2 300 × 0,05/2 = **57,5 kg**. Ce résultat est une masse à déplacer, pas une masse à débarquer. Vérifier ensuite que le compartiment de destination et l’arrimage le permettent.

## Retirer ou consommer une masse

Lors d’un retrait, soustraire la masse ET son moment. À partir de l’état initial 2 300 kg / 5 000 kg·m, consommer 100 kg de carburant au bras 2,50 m donne 2 200 kg / 4 750 kg·m, soit 2,159 m. Le centrage avance ici parce que le carburant retiré était derrière le centrage initial. Cette direction dépend du réservoir : elle n’est pas universelle. À masse variable, la formule du simple déplacement interne ne convient pas.

### Distinguer retrait et déplacement

La différence de CG après retrait d’une masse m au bras a est m × (CG initial − a)/(M − m). Le dénominateur est la masse restante, contrairement au déplacement interne. Cette expression confirme le sens : retirer une masse derrière le CG le fait avancer ; la retirer devant le fait reculer.

**Cas autonome :** repartir de 2 300 kg et 5 000 kg·m, puis ajouter 50 kg au bras 4 m. Les nouveaux totaux sont 2 350 kg et 5 200 kg·m ; CG = **2,213 m** arrondi. Le déplacement vers l’arrière est cohérent avec la position ajoutée. Ne pas confondre ce cas avec la consommation précédente.

Si le bras du carburant varie avec la quantité, le modèle à bras constant ne suffit plus : employer les données de réservoir et la méthode applicables. Toujours nommer l’état de référence avant de soustraire un carburant déjà déduit dans un calcul antérieur.

### Enchaîner plusieurs changements sans mélanger les états

Créer deux colonnes de variations : masse et moment. Un déplacement interne inscrit zéro en masse ; un retrait inscrit une masse négative et son moment négatif ; un ajout inscrit les deux positifs. Calculer ensuite les deux totaux finaux, puis leur quotient. Ne pas réutiliser un CG arrondi comme point de départ si les totaux exacts sont disponibles.

**Exemple guidé distinct :** état initial 2 300 kg et 5 000 kg·m ; retirer 50 kg au bras 3 m, puis ajouter 20 kg au bras 1 m. Masse finale 2 270 kg ; moment final 4 870 kg·m ; CG = **2,145 m** arrondi. Retrait arrière et ajout avant font tous deux avancer le CG, ce qui contrôle le sens. Ce résultat arithmétique ne vérifie pas à lui seul l’enveloppe ou les compartiments.

## MAC et limites simultanées

Pour exprimer le centrage en pourcentage de corde aérodynamique moyenne : %MAC = 100 × (bras CG − bras du bord d’attaque de la MAC) / longueur MAC. Si CG = 2,20 m, bord d’attaque = 1,80 m et MAC = 1,60 m, le résultat est 25 %. Les limites sont celles du manuel. Une masse sous la limite ne prouve ni le respect du centrage, ni celui des compartiments. Contrôler les états pertinents, notamment départ et arrivée, et toute modification de dernière minute.

### Changer de représentation sans changer de référence

La MAC est une longueur ; son bord d’attaque possède un bras depuis le datum. Soustraire d’abord ce bras au bras du CG pour obtenir une distance depuis le bord d’attaque, puis diviser par la longueur MAC. Le résultat est sans unité avant multiplication par 100.

**Contrôle inverse du 25 % :** 1,80 + 0,25 × 1,60 = 2,20 m. Utiliser directement 2,20/1,60 donnerait 137,5 % en oubliant le bord d’attaque. Une conversion correcte doit revenir au même bras.

**Décision :** si une limite arrière inclusive vaut 30 %MAC, 31 % est hors limite, quelle que soit la marge de masse. Replacer le point dans l’enveloppe à la masse considérée et vérifier départ, états en vol pertinents et arrivée. Les exemples sont fictifs ; les contraintes de compartiment et les changements de dernière minute font partie du bilan à contrôler.

**Format de travail :** 15 QCM de chapitre, quatre choix et une réponse correcte. La matière 031 comporte 21 questions en 60 minutes dans le tableau EASA cité. Ce chapitre ne couvre qu’une partie de la matière ; il ne constitue pas une épreuve complète. Cadre EASA, même en anglais ou avec les unités US/britanniques.

## Références

- [EASA — programme Part-FCL](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8)
- [EASA — format de la matière, AMC1 ARA.FCL.300(b)](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67)
- [FAA Weight and Balance Handbook — méthodes de calcul et déplacements de charge](https://www.faa.gov/sites/faa.gov/files/2023-09/Weight_Balance_Handbook.pdf)

**Support de révision CPL(A).** À faire valider par votre instructeur. Pour un vol réel, utiliser le manuel de l’avion et les publications en vigueur. Les exemples chiffrés sont pédagogiques.', '**Objective:** Build and cross-check a moment balance, distinguish shifting from removal, solve an inverse shift and convert CG to percentage MAC.

## Establish mass and moment

![CPL_031_001](/illustrations/en/cpl-031-001.svg)

Use one datum and consistent signed arms. Moment equals mass × arm; CG equals total moment / total mass. Fictional example: aircraft 1,800 kg at 2.00 m, load 300 kg at 3.00 m, fuel 200 kg at 2.50 m. Total mass = 2,300 kg; moment = 3,600 + 900 + 500 = 5,000 kg·m; CG = 2.174 m. Do not average the arms directly. Diagram data belong to this exercise, not a real aircraft.

### Build a balance that can be checked

In this exercise, arms increase aft. Each loading entry must associate mass, arm from the common datum and moment. Do not mix arms from datum with distances from CG. Dividing (kg·m)/kg returns an arm in metres.

**Independent check:** CG 2.174 m lies between extreme arms 2.00 and 3.00 m and near the initial aircraft’s arm, which carries most mass. Simply averaging arms gives 2.50 m and overweights the small loads.

If the datum is changed so every arm increases by 1 m, total moment increases by 2,300 kg·m and CG in the new reference increases by 1 m. Physical loading has not moved: limits must also use the same reference. A numerical arm only has meaning with its datum.

## Move a mass already on board

With no mass added or removed, CG change = moved mass × arm change / total mass. Moving 100 kg from 3.00 m to 1.00 m in a 2,300 kg aircraft changes moment by −200 kg·m and moves CG forward by 0.087 m. The example’s new CG is 4,800 / 2,300 = 2.087 m. The sign indicates direction; total mass is unchanged. Also check compartment limits, restraint and permitted loading.

### Derive and cross-check the formula

An internal movement removes moment m × old arm and adds m × new arm. The difference is m × (new − old). Since total mass M is unchanged, ΔCG = Δmoment/M. This reasoning explains the sign and avoids memorising a formula without its conditions.

**Two-method check:** new moment is 5,000 − 100 × 3 + 100 × 1 = 4,800 kg·m, so CG is 2.086956… m. Using unrounded values, original CG 5,000/2,300 minus 200/2,300 gives exactly the same result.

**Inverse problem:** moving CG forward by 0.05 m in a 2,300 kg aircraft by moving a load 2 m forward requires moving 2,300 × 0.05/2 = **57.5 kg**. This is mass to reposition, not mass to unload. Then verify destination compartment limits and restraint.

## Remove or consume mass

When removing mass, subtract both mass AND its moment. From the initial 2,300 kg / 5,000 kg·m condition, burning 100 kg of fuel at 2.50 m leaves 2,200 kg / 4,750 kg·m, giving 2.159 m. CG moves forward here because removed fuel was aft of the original CG. The direction depends on tank location; it is not universal. The simple internal-shift formula is inappropriate when total mass changes.

### Distinguish removal from movement

CG change after removing mass m at arm a is m × (initial CG − a)/(M − m). The denominator is remaining mass, unlike an internal shift. This confirms direction: removing mass aft of CG moves it forward; removing mass forward moves it aft.

**Independent case:** start from 2,300 kg and 5,000 kg·m, then add 50 kg at arm 4 m. New totals are 2,350 kg and 5,200 kg·m; CG = **2.213 m** rounded. Aft movement is consistent with the added position. Do not combine this case with the preceding fuel burn.

If fuel arm changes with quantity, the constant-arm model is insufficient: use applicable tank data and methods. Always name the reference condition before subtracting fuel already deducted in an earlier calculation.

### Combine changes without mixing states

Make two columns of changes: mass and moment. An internal shift records zero mass change; removal records negative mass and its negative moment; addition records both as positive. Then calculate both final totals and their quotient. Do not reuse a rounded CG as the starting point when exact totals are available.

**Separate guided example:** initial state 2,300 kg and 5,000 kg·m; remove 50 kg at arm 3 m, then add 20 kg at arm 1 m. Final mass 2,270 kg; final moment 4,870 kg·m; CG = **2.145 m** rounded. Aft removal and forward addition both move CG forward, providing a direction check. This arithmetic result alone does not verify the envelope or compartments.

## MAC and simultaneous limits

Express CG as percentage mean aerodynamic chord using %MAC = 100 × (CG arm − MAC leading-edge arm) / MAC length. With CG = 2.20 m, leading edge = 1.80 m and MAC = 1.60 m, the result is 25%. Limits come from the manual. Being below a mass limit does not prove CG or compartment compliance. Check relevant conditions, including departure and arrival, and every last-minute loading change.

### Change representation without changing reference

MAC is a length; its leading edge has an arm from datum. First subtract that arm from CG arm to obtain distance from the leading edge, then divide by MAC length. The result is dimensionless before multiplication by 100.

**Reverse check of 25%:** 1.80 + 0.25 × 1.60 = 2.20 m. Using 2.20/1.60 directly gives 137.5% by omitting the leading edge. A correct conversion must return the same arm.

**Decision:** with an inclusive aft limit of 30% MAC, 31% is outside, regardless of mass margin. Plot the point at the relevant mass and check departure, relevant in-flight conditions and arrival. Examples are fictional; compartment constraints and last-minute changes form part of the balance to verify.

**Practice format:** 15 chapter MCQs, four choices and one correct answer. Subject 031 has 21 questions in 60 minutes in the cited EASA table. This chapter covers only part of the subject and is not a full paper. EASA framework, including with English or US/British units.

## References

- [Reference 1](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8)
- [Reference 2](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67)
- [Reference 3](https://www.faa.gov/sites/faa.gov/files/2023-09/Weight_Balance_Handbook.pdf)

**CPL(A) revision material.** To be reviewed by your instructor. For an actual flight, use the aircraft manual and current publications. Numerical examples are for training.', 3, 1, 1, 45, '["CPL", "mcq-four-options"]', '["Construire et contrôler un bilan de moments, distinguer déplacement et retrait, résoudre un déplacement inverse et convertir le centrage en pourcentage MAC."]')
ON CONFLICT (id) DO UPDATE SET license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, title_fr = EXCLUDED.title_fr, title_en = EXCLUDED.title_en, content_fr = EXCLUDED.content_fr, content_en = EXCLUDED.content_en, difficulty = EXCLUDED.difficulty, order_index = EXCLUDED.order_index, level = EXCLUDED.level, duration_minutes = EXCLUDED.duration_minutes, tags = EXCLUDED.tags, learning_objectives = EXCLUDED.learning_objectives;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('7e4a7cbf-0000-4000-8000-00007e4a7cbf', '1c7e785d-0000-4000-8000-00001c7e785d', 'CPL', 'mass_and_balance', 'mass_and_balance', 'Établir la masse et le moment', 3, 3, 'Avion 1 800 kg au bras 2,00 m, charge 300 kg à 3,00 m, carburant 200 kg à 2,50 m : masse totale ?

![CPL_031_001](/illustrations/cpl-031-001.svg)', 'Aircraft 1,800 kg at 2.00 m, load 300 kg at 3.00 m, fuel 200 kg at 2.50 m: total mass?

![CPL_031_001](/illustrations/en/cpl-031-001.svg)', '["2 300 kg", "2 100 kg", "2 000 kg", "2 500 kg"]', '["2,300 kg", "2,100 kg", "2,000 kg", "2,500 kg"]', '2 300 kg', '1 800 + 300 + 200 = 2 300 kg. 2 100 kg omet le carburant ; additionner un moment à ces masses mélangerait des unités différentes.

![CPL_031_001](/illustrations/cpl-031-001.svg)', '1,800 + 300 + 200 = 2,300 kg. 2,100 kg omits fuel; adding a moment to these masses would mix different units.

![CPL_031_001](/illustrations/en/cpl-031-001.svg)', '["CPL", "exam-mcq", "exam-subject:031"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('7e4a7cc0-0000-4000-8000-00007e4a7cc0', '1c7e785d-0000-4000-8000-00001c7e785d', 'CPL', 'mass_and_balance', 'mass_and_balance', 'Établir la masse et le moment', 3, 3, 'Une charge de 300 kg au bras 3,00 m produit quel moment ?', 'What moment does a 300 kg load at 3.00 m produce?', '["100 kg·m", "303 kg·m", "3 000 kg·m", "900 kg·m"]', '["100 kg·m", "303 kg·m", "3,000 kg·m", "900 kg·m"]', '900 kg·m', '300 × 3 = 900 kg·m. 100 résulte d’une division ; 303 additionne masse et bras, ce qui ne produit pas un moment.

![CPL_031_001](/illustrations/cpl-031-001.svg)', '300 × 3 = 900 kg·m. 100 comes from division; 303 adds mass and arm, which does not produce a moment.

![CPL_031_001](/illustrations/en/cpl-031-001.svg)', '["CPL", "exam-mcq", "exam-subject:031"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('7e4a7cc1-0000-4000-8000-00007e4a7cc1', '1c7e785d-0000-4000-8000-00001c7e785d', 'CPL', 'mass_and_balance', 'mass_and_balance', 'Établir la masse et le moment', 3, 3, 'Masse totale 2 300 kg, moment total 5 000 kg·m : centrage arrondi à 0,001 m ?', 'Total mass 2,300 kg, total moment 5,000 kg·m: CG rounded to 0.001 m?', '["2,500 m", "2,000 m", "2,174 m", "0,460 m"]', '["2.500 m", "2.000 m", "2.174 m", "0.460 m"]', '2,174 m', '5 000/2 300 = 2,173913… m, soit 2,174 m. 0,460 inverse le rapport ; 2,500 ne pondère pas correctement les bras de l’exemple.

![CPL_031_001](/illustrations/cpl-031-001.svg)', '5,000/2,300 = 2.173913… m, rounded to 2.174 m. 0.460 reverses the ratio; 2.500 does not correctly weight the example’s arms.

![CPL_031_001](/illustrations/en/cpl-031-001.svg)', '["CPL", "exam-mcq", "exam-subject:031"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('7e4a7cc2-0000-4000-8000-00007e4a7cc2', '1c7e785d-0000-4000-8000-00001c7e785d', 'CPL', 'mass_and_balance', 'mass_and_balance', 'Établir la masse et le moment', 3, 3, 'Deux charges de masses différentes occupent deux bras différents. Quelle méthode donne leur centrage combiné ?', 'Two unequal masses have different arms. Which method gives their combined CG?', '["Différence des masses divisée par somme des bras", "Somme des moments divisée par somme des masses", "Moyenne simple des bras", "Somme des bras"]', '["Mass difference divided by sum of arms", "Total moment divided by total mass", "Simple average of arms", "Sum of arms"]', 'Somme des moments divisée par somme des masses', 'Chaque bras doit être pondéré par la masse qu’il porte. Diviser le total des moments par le total des masses donne cette moyenne pondérée et conserve l’unité de longueur.

![CPL_031_001](/illustrations/cpl-031-001.svg)', 'Each arm must be weighted by the mass at that arm. Total moment divided by total mass gives this weighted average and retains the length unit.

![CPL_031_001](/illustrations/en/cpl-031-001.svg)', '["CPL", "exam-mcq", "exam-subject:031"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('7e4a8084-0000-4000-8000-00007e4a8084', '1c7e785d-0000-4000-8000-00001c7e785d', 'CPL', 'mass_and_balance', 'mass_and_balance', 'Déplacer une masse déjà embarquée', 3, 3, 'On déplace 100 kg du bras 3,00 m au bras 1,00 m. Variation du moment total ?', '100 kg is moved from arm 3.00 m to arm 1.00 m. Change in total moment?', '["−200 kg·m", "+200 kg·m", "−100 kg·m", "+400 kg·m"]', '["−200 kg·m", "+200 kg·m", "−100 kg·m", "+400 kg·m"]', '−200 kg·m', '100 × (1 − 3) = −200 kg·m. Le signe négatif correspond au déplacement vers les bras plus petits ; +200 inverse le sens du déplacement.

![CPL_031_001](/illustrations/mass-balance-method.svg)', '100 × (1 − 3) = −200 kg·m. The negative sign corresponds to movement towards smaller arms; +200 reverses movement direction.

![CPL_031_001](/illustrations/en/mass-balance-method.svg)', '["CPL", "exam-mcq", "exam-subject:031"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('7e4a8085-0000-4000-8000-00007e4a8085', '1c7e785d-0000-4000-8000-00001c7e785d', 'CPL', 'mass_and_balance', 'mass_and_balance', 'Déplacer une masse déjà embarquée', 3, 3, 'Dans un avion de 2 300 kg, déplacer 100 kg de 2,00 m vers l’avant fait varier le centrage de…', 'In a 2,300 kg aircraft, moving 100 kg forward by 2.00 m changes CG by…', '["0,087 m vers l’arrière", "0,870 m vers l’avant", "2,000 m vers l’avant", "0,087 m vers l’avant"]', '["0.087 m aft", "0.870 m forward", "2.000 m forward", "0.087 m forward"]', '0,087 m vers l’avant', '100 × 2/2 300 = 0,086956… m vers l’avant. Le CG se déplace moins que la charge : celle-ci ne représente qu’une fraction de la masse totale.

![CPL_031_001](/illustrations/mass-balance-method.svg)', '100 × 2/2,300 = 0.086956… m forward. CG moves less than the load because the load is only a fraction of total mass.

![CPL_031_001](/illustrations/en/mass-balance-method.svg)', '["CPL", "exam-mcq", "exam-subject:031"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('7e4a8086-0000-4000-8000-00007e4a8086', '1c7e785d-0000-4000-8000-00001c7e785d', 'CPL', 'mass_and_balance', 'mass_and_balance', 'Déplacer une masse déjà embarquée', 3, 3, 'Avion de 2 400 kg, CG initial à 2,20 m. On veut obtenir 2,10 m en déplaçant une charge de 3,00 m vers l’avant, sans ajout ni retrait. Quelle masse déplacer ?', 'Aircraft mass 2,400 kg, initial CG 2.20 m. Target CG is 2.10 m by moving a load 3.00 m forward, without addition or removal. What mass must be moved?', '["240 kg", "8 kg", "80 kg", "120 kg"]', '["240 kg", "8 kg", "80 kg", "120 kg"]', '80 kg', 'Le déplacement souhaité vaut 0,10 m. La variation de moment requise est 2 400 × 0,10 = 240 kg·m ; divisée par 3 m, elle impose 80 kg. 240 est un moment, pas la masse à déplacer ; 120 utilise une distance de 2 m au lieu de 3 m.

![CPL_031_001](/illustrations/mass-balance-method.svg)', 'Required CG movement is 0.10 m. Required moment change is 2,400 × 0.10 = 240 kg·m; dividing by 3 m requires 80 kg. 240 is a moment, not the mass to move; 120 uses 2 m rather than 3 m.

![CPL_031_001](/illustrations/en/mass-balance-method.svg)', '["CPL", "exam-mcq", "exam-subject:031"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('7e4a8087-0000-4000-8000-00007e4a8087', '1c7e785d-0000-4000-8000-00001c7e785d', 'CPL', 'mass_and_balance', 'mass_and_balance', 'Déplacer une masse déjà embarquée', 3, 3, 'Masse 2 300 kg et moment 5 000 kg·m. Un déplacement interne réduit le moment de 200 kg·m. Nouveau centrage ?', 'Mass 2,300 kg and moment 5,000 kg·m. An internal shift reduces moment by 200 kg·m. New CG?', '["1,913 m", "2,087 m", "2,364 m", "2,174 m"]', '["1.913 m", "2.087 m", "2.364 m", "2.174 m"]', '2,087 m', 'Nouveau moment 4 800 kg·m ; masse toujours 2 300 kg. Le quotient vaut 2,087 m arrondi. Ajouter 200 au moment inverserait l’effet ; garder 5 000 ignorerait le déplacement.

![CPL_031_001](/illustrations/mass-balance-method.svg)', 'New moment is 4,800 kg·m; mass remains 2,300 kg. The quotient rounds to 2.087 m. Adding 200 reverses the effect; retaining 5,000 ignores the shift.

![CPL_031_001](/illustrations/en/mass-balance-method.svg)', '["CPL", "exam-mcq", "exam-subject:031"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('7e4a8449-0000-4000-8000-00007e4a8449', '1c7e785d-0000-4000-8000-00001c7e785d', 'CPL', 'mass_and_balance', 'mass_and_balance', 'Retirer ou consommer une masse', 3, 3, 'État initial 2 300 kg / 5 000 kg·m. On consomme 100 kg au bras 2,50 m. Nouveau moment ?', 'Initial condition 2,300 kg / 5,000 kg·m. 100 kg is burned at arm 2.50 m. New moment?', '["4 750 kg·m", "4 900 kg·m", "5 250 kg·m", "4 500 kg·m"]', '["4,750 kg·m", "4,900 kg·m", "5,250 kg·m", "4,500 kg·m"]', '4 750 kg·m', 'Le moment du carburant retiré est 100 × 2,50 = 250 kg·m. Il faut le soustraire de 5 000 ; retirer seulement 100 confond masse et moment.

![CPL_031_001](/illustrations/mass-balance-method.svg)', 'Removed fuel moment is 100 × 2.50 = 250 kg·m. Subtract it from 5,000; subtracting only 100 confuses mass and moment.

![CPL_031_001](/illustrations/en/mass-balance-method.svg)', '["CPL", "exam-mcq", "exam-subject:031"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('7e4a844a-0000-4000-8000-00007e4a844a', '1c7e785d-0000-4000-8000-00001c7e785d', 'CPL', 'mass_and_balance', 'mass_and_balance', 'Retirer ou consommer une masse', 3, 3, 'Après consommation, masse 2 200 kg et moment 4 750 kg·m. Nouveau centrage ?', 'After fuel burn, mass is 2,200 kg and moment 4,750 kg·m. New CG?', '["2,065 m", "2,273 m", "2,500 m", "2,159 m"]', '["2.065 m", "2.273 m", "2.500 m", "2.159 m"]', '2,159 m', '4 750/2 200 = 2,15909… m. 2,065 m utilise l’ancienne masse 2 300 kg ; 2,273 m oublie de retirer le moment du carburant.

![CPL_031_001](/illustrations/mass-balance-method.svg)', '4,750/2,200 = 2.15909… m. 2.065 m uses the old mass of 2,300 kg; 2.273 m fails to remove fuel moment.

![CPL_031_001](/illustrations/en/mass-balance-method.svg)', '["CPL", "exam-mcq", "exam-subject:031"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('4b06042e-0000-4000-8000-00004b06042e', '1c7e785d-0000-4000-8000-00001c7e785d', 'CPL', 'mass_and_balance', 'mass_and_balance', 'Retirer ou consommer une masse', 3, 3, 'État initial : 2 300 kg et 5 000 kg·m. On consomme 100 kg au bras 2,50 m, puis on ajoute 50 kg au bras 1,00 m, sans autre changement. Bras depuis le même datum : CG final arrondi à 0,001 m ?', 'Initial state: 2,300 kg and 5,000 kg·m. Burn 100 kg at arm 2.50 m, then add 50 kg at arm 1.00 m, with no other change. All arms use the same datum: final CG rounded to 0.001 m?', '["2,087 m", "2,356 m", "2,133 m", "2,159 m"]', '["2.087 m", "2.356 m", "2.133 m", "2.159 m"]', '2,133 m', 'Masse finale = 2 300 − 100 + 50 = 2 250 kg. Moment final = 5 000 − 250 + 50 = 4 800 kg·m. CG = 4 800/2 250 = 2,133 m. 2,159 oublie l’ajout ; 2,087 conserve la masse initiale ; 2,356 ajoute le moment du carburant au lieu de le retirer.

![CPL_031_001](/illustrations/mass-balance-method.svg)', 'Final mass = 2,300 − 100 + 50 = 2,250 kg. Final moment = 5,000 − 250 + 50 = 4,800 kg·m. CG = 4,800/2,250 = 2.133 m. 2.159 omits the addition; 2.087 retains initial mass; 2.356 adds fuel moment instead of removing it.

![CPL_031_001](/illustrations/en/mass-balance-method.svg)', '["CPL", "exam-mcq", "exam-subject:031"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('4b06788e-0000-4000-8000-00004b06788e', '1c7e785d-0000-4000-8000-00001c7e785d', 'CPL', 'mass_and_balance', 'mass_and_balance', 'MAC et limites simultanées', 3, 3, 'CG 2,20 m ; bord d’attaque MAC 1,80 m ; longueur MAC 1,60 m. Centrage en %MAC ?', 'CG 2.20 m; MAC leading edge 1.80 m; MAC length 1.60 m. CG in %MAC?', '["112,5 %", "25 %", "40 %", "137,5 %"]', '["112.5%", "25%", "40%", "137.5%"]', '25 %', 'Distance depuis le bord d’attaque : 2,20 − 1,80 = 0,40 m. Diviser par 1,60 puis multiplier par 100 donne 25 %. 137,5 % oublie la soustraction.

![CPL_031_001](/illustrations/cpl-031-001.svg)', 'Distance from leading edge: 2.20 − 1.80 = 0.40 m. Divide by 1.60 and multiply by 100 to obtain 25%. 137.5% omits the subtraction.

![CPL_031_001](/illustrations/en/cpl-031-001.svg)', '["CPL", "exam-mcq", "exam-subject:031"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('4b06788f-0000-4000-8000-00004b06788f', '1c7e785d-0000-4000-8000-00001c7e785d', 'CPL', 'mass_and_balance', 'mass_and_balance', 'MAC et limites simultanées', 3, 3, 'Limites fictives 20–30 %MAC incluses. Un centrage à 31 %MAC est…', 'Fictional limits are 20–30% MAC inclusive. A CG at 31% MAC is…', '["Hors limite arrière", "Acceptable si la masse est faible", "Hors limite avant", "Acceptable avec davantage de carburant dans tous les cas"]', '["Beyond the aft limit", "Acceptable if mass is low", "Beyond the forward limit", "Always acceptable with more fuel"]', 'Hors limite arrière', '31 % dépasse la limite arrière de 30 %. Une masse faible ne donne pas une dérogation à cette limite ; toute modification de carburant demande un nouveau calcul.

![CPL_031_001](/illustrations/cpl-031-001.svg)', '31% exceeds the 30% aft limit. Low mass does not waive this limit; any fuel change requires a new calculation.

![CPL_031_001](/illustrations/en/cpl-031-001.svg)', '["CPL", "exam-mcq", "exam-subject:031"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('4b067890-0000-4000-8000-00004b067890', '1c7e785d-0000-4000-8000-00001c7e785d', 'CPL', 'mass_and_balance', 'mass_and_balance', 'MAC et limites simultanées', 3, 3, 'La masse de départ respecte son maximum. Quelle conclusion peut-on tirer sur le chargement ?', 'Departure mass meets its maximum. What can be concluded about loading?', '["Le centrage est automatiquement acceptable", "Les limites de compartiment ne s’appliquent plus", "Le seul respect du départ garantit l’état à l’arrivée", "Centrage et autres limites de chargement restent à vérifier aux états pertinents"]', '["CG is automatically acceptable", "Compartment limits no longer apply", "Departure compliance alone guarantees arrival compliance", "CG and other loading limits still need checking at relevant conditions"]', 'Centrage et autres limites de chargement restent à vérifier aux états pertinents', 'Le respect d’une masse maximale ne prouve pas le respect du CG, des compartiments ou des autres états de vol. Il faut vérifier ces contraintes simultanément avec les données applicables.

![CPL_031_001](/illustrations/cpl-031-001.svg)', 'Meeting a maximum mass does not establish CG, compartment or other flight-condition compliance. Check these constraints together using applicable data.

![CPL_031_001](/illustrations/en/cpl-031-001.svg)', '["CPL", "exam-mcq", "exam-subject:031"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('4b067891-0000-4000-8000-00004b067891', '1c7e785d-0000-4000-8000-00001c7e785d', 'CPL', 'mass_and_balance', 'mass_and_balance', 'MAC et limites simultanées', 3, 3, 'Une charge est déplacée après le devis sans ajout ni retrait. Quelle action ?', 'A load is moved after the calculation without addition or removal. What action is needed?', '["Ajouter la masse déplacée au total", "Vérifier uniquement la masse maximale", "Recalculer le moment et vérifier les limites", "Conserver le moment car la masse est identique"]', '["Add the moved mass to the total", "Check maximum mass only", "Recalculate moment and check limits", "Keep moment because mass is unchanged"]', 'Recalculer le moment et vérifier les limites', 'Un déplacement interne change le moment même à masse constante. Recalculer son effet puis vérifier l’enveloppe et les limites locales ; ne pas ajouter une seconde fois la charge déjà embarquée.

![CPL_031_001](/illustrations/mass-balance-method.svg)', 'An internal movement changes moment even at constant mass. Recalculate its effect and check the envelope and local limits; do not add the already loaded mass again.

![CPL_031_001](/illustrations/en/mass-balance-method.svg)', '["CPL", "exam-mcq", "exam-subject:031"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO lessons (id, license, category, theme, title_fr, title_en, content_fr, content_en, difficulty, order_index, level, duration_minutes, tags, learning_objectives) VALUES
('1c8c8fde-0000-4000-8000-00001c8c8fde', 'CPL', 'performance', 'performance', 'Performances : tableaux et marges', 'Performance: tables and margins', '**Objectif :** Interpoler avec des fractions différentes, suivre les corrections déjà incluses et décider sur une marge calculée sans masquer un dépassement.

## Lire les conditions

![CPL_032_001](/illustrations/cpl-032-001.svg)

Un résultat de performances n’est valable que pour les conditions du tableau : masse, configuration, altitude-pression, température, vent et état de piste. Distinguer roulement au sol et distance jusqu’à une hauteur écran. Dans ce chapitre, le tableau fictif donne une distance jusqu’à 50 ft, sur piste sèche, horizontale, sans vent, avec configuration fixée. Les quatre cases sont : 0 ft/10 °C : 500 m ; 0 ft/30 °C : 600 m ; 2 000 ft/10 °C : 650 m ; 2 000 ft/30 °C : 780 m. Elles ne décrivent aucun avion réel.

### Lecture guidée des données

Pour 0 ft et 20 °C, conserver la ligne 0 ft et encadrer 20 °C par les colonnes 10 et 30 °C. Les cases utiles sont 500 et 600 m. Prendre la ligne 2 000 ft donnerait un résultat calculable mais répondrait à un autre problème. Le SVG fournit les données communes ; chaque énoncé définit sa propre cible. La masse reste fixée pour l’ensemble du tableau fictif.

Avant tout calcul, noter : type de distance, masse/configuration supposées, altitude-pression, température et corrections déjà incluses. Le roulement s’arrête à l’envol ; la distance jusqu’à 50 ft comprend également une phase en vol. Deux valeurs en mètres ne décrivent donc pas nécessairement la même grandeur. Une donnée manquante ne doit pas être remplacée tacitement par une condition standard.

Le schéma en tête est le tableau du dossier autonome des QCM de synthèse. Le premier exemple guidé ci-dessous utilise un autre tableau, repris dans ses corrigés. Dans chaque cas, conserver uniquement les données de son propre énoncé.

## Interpoler dans deux directions

L’interpolation linéaire est ici autorisée par l’énoncé. À 20 °C, la valeur à 0 ft est 550 m ; à 2 000 ft elle est 715 m. À 1 000 ft et 20 °C, prendre leur milieu : 632,5 m. Une interpolation conserve les unités et reste dans le domaine du tableau. Hors domaine, ne pas extrapoler sans méthode autorisée. Pour une donnée réelle, suivre la méthode du manuel même si elle diffère de cette interpolation pédagogique.

### Méthode détaillée et contrôle indépendant

Entre y0 et y1 associés à x0 et x1, calculer f = (x − x0)/(x1 − x0), puis y = y0 + f × (y1 − y0). La fraction f n’a pas d’unité et reste entre 0 et 1 pour une interpolation.

Pour 1 000 ft / 20 °C :

1. Fraction de température : (20 − 10)/(30 − 10) = 0,5.
2. À 0 ft : 500 + 0,5 × 100 = 550 m.
3. À 2 000 ft : 650 + 0,5 × 130 = 715 m.
4. Fraction d’altitude : 1 000/2 000 = 0,5.
5. Distance : 550 + 0,5 × (715 − 550) = 632,5 m.

Pour contrôler, interpoler d’abord l’altitude : à 1 000 ft, 575 m à 10 °C et 690 m à 30 °C. Leur milieu vaut aussi 632,5 m. Le résultat doit rester entre les valeurs qui l’encadrent.

**Application autonome :** déterminer la distance à 500 ft / 15 °C avant de lire la solution. Les deux fractions valent 0,25 : 525 m à 0 ft et 682,5 m à 2 000 ft donnent **564,375 m**. Une moyenne simple serait fausse : elle n’est valable qu’à mi-intervalle.

### Passer à un autre tableau
Ne pas mémoriser 632,5 m comme réponse générale. Avec quatre cases différentes, calculer les deux fractions et appliquer chacune à son axe. Un quart d''intervalle sur la température et trois quarts sur l''altitude ne donnent pas la moyenne des quatre cases.

**Contrôle guidé distinct :** un tableau donne 400/480 m à 0 ft et 560/680 m à 2 000 ft, pour 0/20 °C. À 500 ft et 10 °C, les valeurs intermédiaires sont 440 et 620 m ; la fraction d''altitude 0,25 donne 485 m. La moyenne des quatre cases vaut 530 m et répond au milieu des deux axes, pas à la cible demandée. Dans les QCM de synthèse, le tableau de données change à nouveau.

## Appliquer les corrections

Appliquer seulement les corrections prescrites, dans leur ordre et avec leur base de calcul. Si l’exercice impose ensuite un facteur 1,15 à 632,5 m, la distance devient 727,375 m, soit 728 m en arrondissant par excès au mètre. Ce facteur est fictif, pas une marge réglementaire universelle. Deux facteurs successifs se multiplient : 1,10 × 1,20 = 1,32 ; ils ne donnent pas 1,30. Ne pas compter deux fois une correction déjà incorporée.

### Identifier l’origine d’une erreur

Conserver les décimales intermédiaires. Arrondir 727,375 au plus proche donnerait 727 m ; l’instruction « par excès » exige 728 m. Une méthode correcte comprend aussi la bonne règle d’arrondi.

Sur une distance de 600 m, appliquer 1,10 puis 1,20 donne 660 m puis 792 m. Additionner les pourcentages donnerait 780 m : cette erreur oublie que la seconde correction agit sur une base déjà augmentée. Vérifier également que le tableau n’inclut pas déjà le facteur considéré. Une donnée de piste sèche ne devient pas une donnée de piste mouillée avec un facteur inventé.

### Lire la portée du facteur
Une valeur peut déjà être corrigée d''un effet, tandis qu''un autre facteur reste à appliquer. Exemple distinct : 610 m incluent déjà un facteur de surface 1,10 ; l''exercice exige encore 1,05 pour une autre correction. Résultat 640,5 m, soit 641 m par excès. Multiplier encore par 1,10 traiterait deux fois le même effet. Conserver une trace « brut → corrections incluses → correction restante → arrondi ».

Un facteur et une majoration additive ne se traitent pas de la même manière : « multiplier par 1,10 puis ajouter 30 m » sur 500 m donne 580 m ; ajouter avant multiplier donnerait 583 m. L''ordre doit venir de la méthode fournie.

## Comparer et décider

Comparer des grandeurs compatibles : une distance d’atterrissage requise avec la LDA applicable, une distance de décollage avec la distance déclarée pertinente et les autres contraintes requises. Dans l’exercice, 728 m requis pour 800 m disponibles laisse 72 m de marge arithmétique ; cela ne vérifie pas les obstacles ni toutes les limites opérationnelles. Une altitude-densité élevée tend à dégrader les performances ; un vent de face peut améliorer la distance sol, mais seul le manuel chiffre son effet. Recalculer après un changement de masse, vent ou piste.

### Formuler une conclusion contrôlable

Une marge de 72 m sur 800 m disponibles représente 9 % de la distance disponible, mais environ 9,89 % des 728 m requis. Tout pourcentage doit préciser sa base. Avec 820 m requis, le critère de longueur échoue de 20 m ; arrondir à la centaine masquerait ce déficit.

Une autre piste peut changer longueur, pente, vent et obstacles. Revoir les données pertinentes même si la longueur annoncée paraît identique. Avant le quiz, expliquer sans réciter : pourquoi 0 ft/20 °C donne 550 m et non 632,5 m ; pourquoi les facteurs se multiplient ; pourquoi une marge positive sur un critère ne suffit pas à autoriser le décollage. Les exercices du chapitre restent des modèles, pas des calculs pour un avion réel.

**Format de travail :** 15 QCM de chapitre, quatre choix et une réponse correcte. La matière 032 comporte 28 questions en 75 minutes dans le tableau EASA cité. Ce chapitre ne couvre qu’une partie de la matière ; il ne constitue pas une épreuve complète. Cadre EASA, même en anglais ou avec les unités US/britanniques.

## Références

- [EASA — programme Part-FCL](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8)
- [EASA — format de la matière, AMC1 ARA.FCL.300(b)](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67)
- [FAA PHAK — performances, principes techniques](https://www.faa.gov/regulationspolicies/handbooksmanuals/aviation/phak/chapter-11-aircraft-performance)

**Support de révision CPL(A).** À faire valider par votre instructeur. Pour un vol réel, utiliser le manuel de l’avion et les publications en vigueur. Les exemples chiffrés sont pédagogiques.', '**Objective:** Interpolate using different fractions, track corrections already included and evaluate a calculated margin without hiding an exceedance.

## Read the conditions

![CPL_032_001](/illustrations/en/cpl-032-001.svg)

A performance result applies only to the table conditions: mass, configuration, pressure altitude, temperature, wind and runway condition. Distinguish ground roll from distance to a screen height. This fictional table gives distance to 50 ft on a dry, level runway, no wind and fixed configuration. Its four cells are: 0 ft/10 °C: 500 m; 0 ft/30 °C: 600 m; 2,000 ft/10 °C: 650 m; 2,000 ft/30 °C: 780 m. They do not describe a real aircraft.

### Guided reading of the inputs

For 0 ft and 20 °C, retain the 0 ft row and bracket 20 °C with the 10 and 30 °C columns. Relevant cells are 500 and 600 m. Using the 2,000 ft row would produce a calculable result answering a different problem. The SVG supplies shared inputs; each question defines its own target. Mass remains fixed throughout the fictional table.

Before calculating, record distance type, assumed mass/configuration, pressure altitude, temperature and corrections already included. Ground roll ends at lift-off; distance to 50 ft also includes an airborne segment. Two values in metres therefore need not describe the same quantity. Do not silently replace a missing input with a standard condition.

The opening diagram is the independent synthesis MCQ table. The first worked example below uses a different table, reproduced in its feedback. In each case, use only that question’s own data.

## Interpolate in two directions

Linear interpolation is authorised for this exercise. At 20 °C, the value at 0 ft is 550 m and at 2,000 ft it is 715 m. At 1,000 ft and 20 °C, take their midpoint: 632.5 m. Interpolation preserves units and remains within the table domain. Do not extrapolate outside it without an authorised method. For real data, follow the manual even if its method differs from this training interpolation.

### Detailed method and independent cross-check

Between y0 and y1 associated with x0 and x1, calculate f = (x − x0)/(x1 − x0), then y = y0 + f × (y1 − y0). Fraction f is dimensionless and stays between 0 and 1 for interpolation.

For 1,000 ft / 20 °C:

1. Temperature fraction: (20 − 10)/(30 − 10) = 0.5.
2. At 0 ft: 500 + 0.5 × 100 = 550 m.
3. At 2,000 ft: 650 + 0.5 × 130 = 715 m.
4. Altitude fraction: 1,000/2,000 = 0.5.
5. Distance: 550 + 0.5 × (715 − 550) = 632.5 m.

Cross-check by interpolating altitude first: at 1,000 ft, 575 m at 10 °C and 690 m at 30 °C. Their midpoint also gives 632.5 m. The result must remain between its bracketing values.

**Independent application:** find distance at 500 ft / 15 °C before reading the solution. Both fractions are 0.25: 525 m at 0 ft and 682.5 m at 2,000 ft give **564.375 m**. A simple average would be wrong: it only works at halfway points.

### Transfer to another table
Do not memorise 632.5 m as a universal result. For different cells, calculate both fractions and apply each on its own axis. One quarter of the temperature interval and three quarters of the altitude interval do not give the four-cell average.

**Separate worked check:** a table gives 400/480 m at 0 ft and 560/680 m at 2,000 ft, for 0/20 °C. At 500 ft and 10 °C, intermediate values are 440 and 620 m; altitude fraction 0.25 gives 485 m. Averaging all cells gives 530 m, the midpoint of both axes rather than the requested target. Synthesis MCQs use another table again.

## Apply corrections

Apply only prescribed corrections, in their specified order and to the correct base. If the exercise then requires a factor of 1.15 on 632.5 m, distance becomes 727.375 m, or 728 m rounded up to a whole metre. This fictional factor is not a universal regulatory margin. Two successive factors multiply: 1.10 × 1.20 = 1.32, not 1.30. Do not apply a correction twice if already included.

### Identify the source of an error

Retain intermediate decimals. Rounding 727.375 to nearest gives 727 m; the instruction to round up requires 728 m. A correct method includes the right rounding rule.

For a 600 m distance, applying 1.10 then 1.20 gives 660 m then 792 m. Adding the percentages gives 780 m: this error ignores that the second correction acts on an already increased base. Also check whether the table already includes the factor. Dry-runway data does not become wet-runway data through an invented factor.

### Read the scope of each correction
A value may already include one effect while another factor remains. Separate example: 610 m already includes a 1.10 surface factor; the exercise still requires 1.05 for another correction. Result 640.5 m, rounded up to 641 m. Multiplying by 1.10 again counts the same effect twice. Record raw value → included corrections → remaining correction → rounding.

A factor and an additive increment differ: “multiply by 1.10 then add 30 m” applied to 500 m gives 580 m; adding before multiplying gives 583 m. Follow the supplied order.

## Compare and decide

Compare compatible quantities: required landing distance with applicable LDA, take-off distance with the relevant declared distance and other required constraints. In the exercise, 728 m required against 800 m available leaves a 72 m arithmetic margin; this does not check obstacles or every operational limit. High density altitude tends to degrade performance; headwind may improve ground distance, but the manual quantifies its effect. Recalculate after mass, wind or runway changes.

### State a conclusion that can be checked

A 72 m margin against 800 m available is 9% of available distance but about 9.89% of the 728 m required. Every percentage needs a stated base. With 820 m required, the length criterion fails by 20 m; rounding to hundreds would hide that shortfall.

A different runway may change length, slope, wind and obstacles. Review relevant inputs even if the stated length appears identical. Before the quiz, explain without rote recall: why 0 ft/20 °C gives 550 m rather than 632.5 m; why factors multiply; why a positive margin on one criterion does not authorise take-off. These exercises remain models, not calculations for a real aircraft.

**Practice format:** 15 chapter MCQs, four choices and one correct answer. Subject 032 has 28 questions in 75 minutes in the cited EASA table. This chapter covers only part of the subject and is not a full paper. EASA framework, including with English or US/British units.

## References

- [Reference 1](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8)
- [Reference 2](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67)
- [Reference 3](https://www.faa.gov/regulationspolicies/handbooksmanuals/aviation/phak/chapter-11-aircraft-performance)

**CPL(A) revision material.** To be reviewed by your instructor. For an actual flight, use the aircraft manual and current publications. Numerical examples are for training.', 3, 1, 1, 50, '["CPL", "mcq-four-options"]', '["Interpoler avec des fractions différentes, suivre les corrections déjà incluses et décider sur une marge calculée sans masquer un dépassement."]')
ON CONFLICT (id) DO UPDATE SET license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, title_fr = EXCLUDED.title_fr, title_en = EXCLUDED.title_en, content_fr = EXCLUDED.content_fr, content_en = EXCLUDED.content_en, difficulty = EXCLUDED.difficulty, order_index = EXCLUDED.order_index, level = EXCLUDED.level, duration_minutes = EXCLUDED.duration_minutes, tags = EXCLUDED.tags, learning_objectives = EXCLUDED.learning_objectives;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('5d60c8c1-0000-4000-8000-00005d60c8c1', '1c8c8fde-0000-4000-8000-00001c8c8fde', 'CPL', 'performance', 'performance', 'Lire les conditions', 3, 3, 'Le tableau fictif donne 500 m à 0 ft/10 °C et 600 m à 0 ft/30 °C. Par interpolation linéaire, valeur à 0 ft/20 °C ?', 'The fictional table gives 500 m at 0 ft/10 °C and 600 m at 0 ft/30 °C. Linear interpolation at 0 ft/20 °C?', '["550 m", "500 m", "600 m", "650 m"]', '["550 m", "500 m", "600 m", "650 m"]', '550 m', '20 °C est à mi-chemin de 10 et 30 °C : 500 + 0,5 × (600 − 500) = 550 m. 500 et 600 sont les bornes ; 650 appartient à une autre ligne d’altitude.

![CPL_032_001](/illustrations/cpl-performance-worked.svg)', '20 °C is halfway between 10 and 30 °C: 500 + 0.5 × (600 − 500) = 550 m. 500 and 600 are endpoints; 650 belongs to another altitude row.

![CPL_032_001](/illustrations/en/cpl-performance-worked.svg)', '["CPL", "exam-mcq", "exam-subject:032"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('5d60cc83-0000-4000-8000-00005d60cc83', '1c8c8fde-0000-4000-8000-00001c8c8fde', 'CPL', 'performance', 'performance', 'Interpoler dans deux directions', 3, 3, 'À 2 000 ft : 650 m à 10 °C et 780 m à 30 °C. Valeur à 20 °C ?', 'At 2,000 ft: 650 m at 10 °C and 780 m at 30 °C. Value at 20 °C?', '["650 m", "680 m", "780 m", "715 m"]', '["650 m", "680 m", "780 m", "715 m"]', '715 m', 'La fraction vaut 0,5 : 650 + 0,5 × (780 − 650) = 715 m. Utiliser l’écart de 100 m de la ligne 0 ft serait incorrect : cette ligne-ci présente un écart de 130 m.

![CPL_032_001](/illustrations/cpl-performance-worked.svg)', 'The fraction is 0.5: 650 + 0.5 × (780 − 650) = 715 m. Using the 0 ft row’s 100 m increment would be wrong: this row’s increment is 130 m.

![CPL_032_001](/illustrations/en/cpl-performance-worked.svg)', '["CPL", "exam-mcq", "exam-subject:032"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('5d60cc84-0000-4000-8000-00005d60cc84', '1c8c8fde-0000-4000-8000-00001c8c8fde', 'CPL', 'performance', 'performance', 'Interpoler dans deux directions', 3, 3, 'À 20 °C : 550 m à 0 ft et 715 m à 2 000 ft. Valeur interpolée à 1 000 ft ?', 'At 20 °C: 550 m at 0 ft and 715 m at 2,000 ft. Interpolated value at 1,000 ft?', '["715 m", "1 265 m", "632,5 m", "582,5 m"]', '["715 m", "1,265 m", "632.5 m", "582.5 m"]', '632,5 m', '1 000 ft est à mi-chemin de 0 et 2 000 ft : 550 + 0,5 × (715 − 550) = 632,5 m. 1 265 m additionne les bornes sans diviser par deux ; 715 m conserve la borne haute.

![CPL_032_001](/illustrations/cpl-performance-worked.svg)', '1,000 ft is halfway between 0 and 2,000 ft: 550 + 0.5 × (715 − 550) = 632.5 m. 1,265 m adds endpoints without dividing by two; 715 m retains the upper endpoint.

![CPL_032_001](/illustrations/en/cpl-performance-worked.svg)', '["CPL", "exam-mcq", "exam-subject:032"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('5d60cc85-0000-4000-8000-00005d60cc85', '1c8c8fde-0000-4000-8000-00001c8c8fde', 'CPL', 'performance', 'performance', 'Interpoler dans deux directions', 3, 3, 'Un tableau s’arrête à 30 °C sans extrapolation autorisée. À 40 °C, quelle conclusion ?', 'A table ends at 30 °C with no authorised extrapolation. At 40 °C, which conclusion follows ?', '["Appliquer le même écart que de 20 à 30 °C", "Les données ne permettent pas ce calcul", "Prolonger la dernière pente du tableau", "Utiliser 30 °C comme valeur conservatrice"]', '["Apply the same increment as from 20 to 30 °C", "These data do not support the calculation", "Extend the final table gradient", "Use 30 °C as a conservative value"]', 'Les données ne permettent pas ce calcul', 'Les trois autres méthodes prolongent ou détournent le domaine du tableau sans autorisation. Il faut une méthode valable pour les conditions demandées.

![CPL_032_001](/illustrations/cpl-performance-worked.svg)', 'The other methods extend or misuse the table domain without authorisation. A valid method for the requested conditions is needed.

![CPL_032_001](/illustrations/en/cpl-performance-worked.svg)', '["CPL", "exam-mcq", "exam-subject:032"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('5d60c8c5-0000-4000-8000-00005d60c8c5', '1c8c8fde-0000-4000-8000-00001c8c8fde', 'CPL', 'performance', 'performance', 'Lire les conditions', 3, 3, 'Une fiche donne 620 m de roulement. Le contrôle demandé porte sur la distance jusqu''à 50 ft. Quelle information manque ?', 'A sheet gives a 620 m ground roll. The required check concerns distance to 50 ft. What information is missing?', '["La distance incluant la phase en vol jusqu''à 50 ft", "Une conversion des 620 m en pieds seulement", "Une marge fixe de 50 m à ajouter", "La longueur disponible uniquement, qui rendrait les grandeurs équivalentes"]', '["Distance including the airborne phase to 50 ft", "Only conversion of 620 m into feet", "A fixed 50 m increment", "Only available length, which would make the quantities equivalent"]', 'La distance incluant la phase en vol jusqu''à 50 ft', 'Convertir l''unité ne change pas le périmètre. La hauteur écran n''est pas une distance horizontale à ajouter ; il faut la donnée correspondante dans la documentation du cas.

![CPL_032_001](/illustrations/cpl-performance-worked.svg)', 'Unit conversion does not change scope. Screen height is not a horizontal increment; obtain the corresponding case data.

![CPL_032_001](/illustrations/en/cpl-performance-worked.svg)', '["CPL", "exam-mcq", "exam-subject:032"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('5d60d048-0000-4000-8000-00005d60d048', '1c8c8fde-0000-4000-8000-00001c8c8fde', 'CPL', 'performance', 'performance', 'Appliquer les corrections', 3, 3, 'Distance 632,5 m, facteur prescrit 1,15. Résultat arrondi par excès au mètre ?', 'Distance 632.5 m, prescribed factor 1.15. Result rounded up to a whole metre?', '["727 m", "695 m", "550 m", "728 m"]', '["727 m", "695 m", "550 m", "728 m"]', '728 m', '632,5 × 1,15 = 727,375 m. L’arrondi par excès donne 728 m ; 727 m serait l’arrondi au plus proche. Garder la précision intermédiaire et appliquer la règle demandée à la fin.

![CPL_032_001](/illustrations/cpl-performance-worked.svg)', '632.5 × 1.15 = 727.375 m. Rounding up gives 728 m; 727 m is rounding to nearest. Keep intermediate precision and apply the requested rounding rule at the end.

![CPL_032_001](/illustrations/en/cpl-performance-worked.svg)', '["CPL", "exam-mcq", "exam-subject:032"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('5d60d049-0000-4000-8000-00005d60d049', '1c8c8fde-0000-4000-8000-00001c8c8fde', 'CPL', 'performance', 'performance', 'Appliquer les corrections', 3, 3, 'Corrections successives imposées : ×1,08 puis ×1,12. Facteur global ?', 'Required successive corrections: ×1.08 then ×1.12. Overall factor?', '["1,04", "2,20", "1,2096", "1,20"]', '["1.04", "2.20", "1.2096", "1.20"]', '1,2096', 'Multiplier : 1,08×1,12 = 1,2096. Additionner les majorations donne 1,20 et manque leur effet composé ; soustraire ou additionner les facteurs ne décrit pas deux corrections successives.

![CPL_032_001](/illustrations/cpl-performance-worked.svg)', 'Multiply: 1.08×1.12 = 1.2096. Adding increments gives 1.20 and misses compounding; subtracting or adding factors does not describe successive corrections.

![CPL_032_001](/illustrations/en/cpl-performance-worked.svg)', '["CPL", "exam-mcq", "exam-subject:032"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('5d60d04a-0000-4000-8000-00005d60d04a', '1c8c8fde-0000-4000-8000-00001c8c8fde', 'CPL', 'performance', 'performance', 'Appliquer les corrections', 3, 3, 'Une distance de 810 m inclut déjà la correction de surface ×1,12. Il reste seulement ×1,08 à appliquer ; arrondir par excès. Résultat ?', '810 m already includes the ×1.12 surface correction. Only ×1.08 remains; round up. Result?', '["874 m", "875 m", "980 m", "908 m"]', '["874 m", "875 m", "980 m", "908 m"]', '875 m', '810×1,08 = 874,8, soit 875 m. 980 applique encore ×1,12 ; 908 applique seulement la mauvaise correction ; 874 arrondit vers le bas. L''origine de chaque facteur détermine son emploi.

![CPL_032_001](/illustrations/cpl-performance-worked.svg)', '810×1.08 = 874.8, rounded up to 875 m. 980 reapplies ×1.12; 908 applies only the wrong correction; 874 rounds down. Each factor''s origin determines its use.

![CPL_032_001](/illustrations/en/cpl-performance-worked.svg)', '["CPL", "exam-mcq", "exam-subject:032"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('5d60d40c-0000-4000-8000-00005d60d40c', '1c8c8fde-0000-4000-8000-00001c8c8fde', 'CPL', 'performance', 'performance', 'Comparer et décider', 3, 3, 'Tableau fictif joint : à 0 ft, 600/720 m pour 5/25 °C ; à 2 000 ft, 800/960 m. Interpolation linéaire à 1 500 ft et 10 °C, puis ×1,12, arrondi par excès. Disponible pour ce même critère : 900 m. Marge ?

![CPL_032_001](/illustrations/cpl-032-001.svg)', 'Attached fictional table: at 0 ft, 600/720 m for 5/25 °C; at 2,000 ft, 800/960 m. Linear interpolation at 1,500 ft and 10 °C, then ×1.12, rounded up. Available for this same criterion: 900 m. Margin?

![CPL_032_001](/illustrations/en/cpl-032-001.svg)', '["18 m", "37 m", "112 m", "−18 m"]', '["18 m", "37 m", "112 m", "−18 m"]', '18 m', 'Fraction température 0,25 : 630 et 840 m. Fraction altitude 0,75 : 787,5 m ; ×1,12 = 882 m ; marge 900−882 = 18. 37 utilise la moyenne des quatre cases ; 112 omet le facteur ; −18 inverse le signe.

![CPL_032_001](/illustrations/cpl-032-001.svg)', 'Temperature fraction 0.25: 630 and 840 m. Altitude fraction 0.75: 787.5 m; ×1.12 = 882 m; margin 900−882 = 18. 37 uses the four-cell mean; 112 omits the factor; −18 reverses the sign.

![CPL_032_001](/illustrations/en/cpl-032-001.svg)', '["CPL", "exam-mcq", "exam-subject:032"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('5d60d40d-0000-4000-8000-00005d60d40d', '1c8c8fde-0000-4000-8000-00001c8c8fde', 'CPL', 'performance', 'performance', 'Comparer et décider', 3, 3, 'Même tableau et cible 1 500 ft/10 °C : distance brute 787,5 m. Appliquer ×1,12 puis ×1,08, arrondir par excès. Disponible 950 m. Conclusion sur ce seul critère ?

![CPL_032_001](/illustrations/cpl-032-001.svg)', 'Same table and 1,500 ft/10 °C target: raw distance 787.5 m. Apply ×1.12 then ×1.08, round up. Available 950 m. Conclusion on this criterion alone?

![CPL_032_001](/illustrations/en/cpl-032-001.svg)', '["Marge de 5 m", "Déficit de 2 m", "Marge de 68 m", "Déficit de 3 m"]', '["5 m margin", "2 m shortfall", "68 m margin", "3 m shortfall"]', 'Déficit de 3 m', '787,5×1,12×1,08 = 952,56, arrondi 953 ; déficit 3. 5 m vient de ×1,20 au lieu du produit ; 2 m arrondit vers le bas ; 68 omet ×1,08. Un faible déficit reste un échec du critère.

![CPL_032_001](/illustrations/cpl-032-001.svg)', '787.5×1.12×1.08 = 952.56, rounded to 953; shortfall 3. A 5 m margin comes from ×1.20 instead of the product; 2 rounds down; 68 omits ×1.08. A small deficit still fails the criterion.

![CPL_032_001](/illustrations/en/cpl-032-001.svg)', '["CPL", "exam-mcq", "exam-subject:032"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('4eb9accb-0000-4000-8000-00004eb9accb', '1c8c8fde-0000-4000-8000-00001c8c8fde', 'CPL', 'performance', 'performance', 'Comparer et décider', 3, 3, 'Distance jusqu’à 50 ft calculée inférieure à la distance disponible. Que reste-t-il à vérifier pour les obstacles ?', 'Calculated distance to 50 ft is below available distance. What remains to be checked for obstacles ?', '["Seulement que la marge de longueur soit positive", "Rien si une correction de vent a été appliquée", "Les contraintes et calculs de franchissement applicables", "Rien, les 50 ft garantissent tout le trajet après décollage"]', '["Only that the length margin is positive", "Nothing if a wind correction was applied", "Applicable obstacle-clearance constraints and calculations", "Nothing, 50 ft guarantees the entire departure path"]', 'Les contraintes et calculs de franchissement applicables', 'Atteindre la hauteur écran sur la distance considérée ne décrit pas toute la trajectoire au-delà. Une correction de vent ne constitue pas une étude d’obstacles.

![CPL_032_001](/illustrations/cpl-performance-worked.svg)', 'Reaching screen height within the distance considered does not describe the entire path beyond it. A wind correction is not an obstacle assessment.

![CPL_032_001](/illustrations/en/cpl-performance-worked.svg)', '["CPL", "exam-mcq", "exam-subject:032"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('4eb84faf-0000-4000-8000-00004eb84faf', '1c8c8fde-0000-4000-8000-00001c8c8fde', 'CPL', 'performance', 'performance', 'Lire les conditions', 3, 3, 'Pour choisir le tableau de décollage approprié, quel ensemble de données est complet parmi les propositions ?', 'Which proposed set of inputs is complete for selecting the appropriate take-off table ?', '["Masse à vide, QNH et température uniquement", "Masse, configuration, altitude-pression, température, vent et piste", "Masse et température, les autres conditions étant toujours standard", "Altitude AMSL et longueur de piste uniquement"]', '["Empty mass, QNH and temperature only", "Mass, configuration, pressure altitude, temperature, wind and runway", "Mass and temperature, with all other conditions always standard", "AMSL altitude and runway length only"]', 'Masse, configuration, altitude-pression, température, vent et piste', 'Les autres ensembles omettent des conditions ou substituent une grandeur différente. La masse à vide n’est pas la masse au décollage ; l’altitude AMSL n’est pas systématiquement l’altitude-pression.

![CPL_032_001](/illustrations/cpl-performance-worked.svg)', 'The other sets omit conditions or substitute another quantity. Empty mass is not take-off mass; AMSL altitude is not always pressure altitude.

![CPL_032_001](/illustrations/en/cpl-performance-worked.svg)', '["CPL", "exam-mcq", "exam-subject:032"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('4eb9accd-0000-4000-8000-00004eb9accd', '1c8c8fde-0000-4000-8000-00001c8c8fde', 'CPL', 'performance', 'performance', 'Comparer et décider', 3, 3, 'Une hausse d’altitude-densité tend généralement à…', 'An increase in density altitude generally tends to…', '["Dégrader les performances de décollage", "Réduire toujours le roulement", "Augmenter automatiquement la puissance disponible", "Supprimer l’effet de la masse"]', '["Degrade take-off performance", "Always shorten ground roll", "Automatically increase available power", "Eliminate mass effects"]', 'Dégrader les performances de décollage', 'À masse et configuration comparables, une faible densité implique davantage de vitesse vraie pour une même vitesse indiquée. Les effets sur le moteur et la propulsion peuvent aussi dégrader le décollage ; la documentation de l’avion chiffre le résultat.

![CPL_032_001](/illustrations/cpl-performance-worked.svg)', 'At comparable mass and configuration, lower density requires greater true airspeed for a given indicated airspeed. Engine and propulsion effects may also degrade take-off; aircraft documentation quantifies the result.

![CPL_032_001](/illustrations/en/cpl-performance-worked.svg)', '["CPL", "exam-mcq", "exam-subject:032"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('4eb9386f-0000-4000-8000-00004eb9386f', '1c8c8fde-0000-4000-8000-00001c8c8fde', 'CPL', 'performance', 'performance', 'Appliquer les corrections', 3, 3, 'Méthode fictive : distance brute 700 m, multiplier par 1,10 puis ajouter 25 m. Valeur avant comparaison ?', 'Fictional method: raw distance 700 m, multiply by 1.10 then add 25 m. Value before comparison?', '["797,5 m", "770 m", "725 m", "795 m"]', '["797.5 m", "770 m", "725 m", "795 m"]', '795 m', '700×1,10+25 = 795. 797,5 ajoute avant de multiplier ; 770 omet l''ajout, 725 omet le facteur. La méthode impose l''ordre, pas seulement une liste de nombres.

![CPL_032_001](/illustrations/cpl-performance-worked.svg)', '700×1.10+25 = 795. 797.5 adds before multiplying; 770 omits the increment; 725 omits the factor. The method specifies order, not just numbers.

![CPL_032_001](/illustrations/en/cpl-performance-worked.svg)', '["CPL", "exam-mcq", "exam-subject:032"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('4eb9accf-0000-4000-8000-00004eb9accf', '1c8c8fde-0000-4000-8000-00001c8c8fde', 'CPL', 'performance', 'performance', 'Comparer et décider', 3, 3, 'La piste attribuée change mais sa longueur est identique. Quelle action ?', 'Assigned runway changes but length is identical. What action follows ?', '["Modifier uniquement l’intitulé de piste du devis", "Conserver le résultat dès que la masse est inchangée", "Revoir vent, pente, obstacles et données de performance concernées", "Conserver le calcul car la longueur n’a pas changé"]', '["Change only the runway name on the worksheet", "Keep the result whenever mass is unchanged", "Review wind, slope, obstacles and relevant performance inputs", "Keep the calculation because length has not changed"]', 'Revoir vent, pente, obstacles et données de performance concernées', 'La longueur n’est qu’une condition parmi plusieurs. Même longueur et même masse ne garantissent pas les mêmes performances ni les mêmes contraintes.

![CPL_032_001](/illustrations/cpl-performance-worked.svg)', 'Length is only one of several conditions. Equal length and mass do not guarantee equal performance or constraints.

![CPL_032_001](/illustrations/en/cpl-performance-worked.svg)', '["CPL", "exam-mcq", "exam-subject:032"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO lessons (id, license, category, theme, title_fr, title_en, content_fr, content_en, difficulty, order_index, level, duration_minutes, tags, learning_objectives) VALUES
('2506b378-0000-4000-8000-00002506b378', 'CPL', 'principles_of_flight', 'principles_of_flight', 'Virage, facteur de charge et décrochage', 'Turning, load factor and stall', '**Objectif :** Résoudre et vérifier les exercices de virage, facteur de charge et décrochage.

## Équilibrer un virage en palier

![CPL_081_001](/illustrations/cpl-081-001.svg)

Dans un virage coordonné stabilisé en palier, la composante verticale de la portance équilibre le poids. La portance totale dépasse donc le poids : n = L/W = 1/cos φ. À 60° d’inclinaison, n = 2 ; à 45°, n ≈ 1,414. Ces relations supposent un palier maintenu ; l’inclinaison seule ne suffit pas à déterminer n dans n’importe quelle manœuvre. La composante horizontale de portance fournit l’accélération centripète.

## Relier charge et décrochage

À masse et configuration fixées, la vitesse de décrochage sous facteur de charge positif vaut approximativement Vs,n = Vs,1 × √n. Si Vs,1 = 50 kt, elle devient 70,7 kt à n = 2 et 59,5 kt à n = 1,414. Le décrochage provient de l’atteinte de l’incidence critique, pas d’une vitesse unique valable pour toutes les manœuvres. Une ressource peut atteindre cette incidence au-dessus de la vitesse de décrochage en palier à 1 g.

## Masse et limites

À configuration et densité identiques, Vs varie comme la racine carrée de la masse. Passer de 1 000 kg à 1 210 kg multiplie Vs par √1,21 = 1,10. Ne pas confondre limite aérodynamique et limite structurale : le domaine de vol doit respecter les deux. La vitesse de manœuvre publiée varie avec la masse ; elle ne protège pas contre toutes les combinaisons d’actions, inversions rapides ou rafales. Utiliser les valeurs et restrictions du manuel.

## Gérer la marge en virage

Augmenter l’inclinaison tout en maintenant le palier demande plus de portance et augmente la traînée induite. Sans adaptation, vitesse ou altitude peuvent diminuer. Près du sol, éviter de resserrer un virage avec une forte action de lacet : une dissymétrie au décrochage accroît le risque de départ en autorotation. Préserver les marges et appliquer les procédures enseignées ; réduire l’incidence est fondamental pour sortir du décrochage. Les calculs ne remplacent pas l’entraînement pratique.

**Format de travail :** 15 QCM de chapitre, quatre choix et une réponse correcte. La matière 081 comporte 37 questions en 75 minutes dans le tableau EASA cité. Ce chapitre ne couvre qu’une partie de la matière ; il ne constitue pas une épreuve complète. Cadre EASA, même en anglais ou avec les unités US/britanniques.

## Références

- [EASA — programme Part-FCL](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8)
- [EASA — format de la matière, AMC1 ARA.FCL.300(b)](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67)
- [FAA PHAK — aérodynamique](https://www.faa.gov/regulationspolicies/handbooksmanuals/aviation/phak/chapter-5-aerodynamics-flight)

**Support de révision CPL(A).** À faire valider par votre instructeur. Pour un vol réel, utiliser le manuel de l’avion et les publications en vigueur. Les exemples chiffrés sont pédagogiques.', '**Objective:** Solve and cross-check exercises on turning, load factor and stall.

## Balance a level turn

![CPL_081_001](/illustrations/en/cpl-081-001.svg)

In a steady coordinated level turn, the vertical lift component balances weight. Total lift therefore exceeds weight: n = L/W = 1/cos φ. At 60° bank, n = 2; at 45°, n ≈ 1.414. These relations assume maintained level flight; bank alone does not determine n in every manoeuvre. Horizontal lift provides centripetal acceleration.

## Relate load to stall

At fixed mass and configuration, stall speed under positive load is approximately Vs,n = Vs,1 × √n. If Vs,1 = 50 kt, it becomes 70.7 kt at n = 2 and 59.5 kt at n = 1.414. Stall results from reaching critical angle of attack, not one speed valid for every manoeuvre. A pull-up may reach that angle above the 1 g level-flight stall speed.

## Mass and limits

At equal configuration and density, Vs varies with the square root of mass. Increasing mass from 1,000 kg to 1,210 kg multiplies Vs by √1.21 = 1.10. Distinguish aerodynamic and structural limits: the flight envelope must respect both. Published manoeuvring speed varies with mass; it does not protect against every combined input, rapid reversal or gust. Use manual values and restrictions.

## Manage turning margins

Increasing bank while maintaining level flight requires more lift and increases induced drag. Without adjustment, speed or altitude may decrease. Near the ground, avoid tightening a turn using strong yaw input: asymmetry at stall increases spin-entry risk. Preserve margins and follow trained procedures; reducing angle of attack is fundamental to stall recovery. Calculations do not replace practical training.

**Practice format:** 15 chapter MCQs, four choices and one correct answer. Subject 081 has 37 questions in 75 minutes in the cited EASA table. This chapter covers only part of the subject and is not a full paper. EASA framework, including with English or US/British units.

## References

- [Reference 1](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8)
- [Reference 2](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67)
- [Reference 3](https://www.faa.gov/regulationspolicies/handbooksmanuals/aviation/phak/chapter-5-aerodynamics-flight)

**CPL(A) revision material.** To be reviewed by your instructor. For an actual flight, use the aircraft manual and current publications. Numerical examples are for training.', 3, 1, 1, 35, '["CPL", "mcq-four-options"]', '["Résoudre et vérifier les exercices de virage, facteur de charge et décrochage."]')
ON CONFLICT (id) DO UPDATE SET license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, title_fr = EXCLUDED.title_fr, title_en = EXCLUDED.title_en, content_fr = EXCLUDED.content_fr, content_en = EXCLUDED.content_en, difficulty = EXCLUDED.difficulty, order_index = EXCLUDED.order_index, level = EXCLUDED.level, duration_minutes = EXCLUDED.duration_minutes, tags = EXCLUDED.tags, learning_objectives = EXCLUDED.learning_objectives;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('10ca81f5-0000-4000-8000-000010ca81f5', '2506b378-0000-4000-8000-00002506b378', 'CPL', 'principles_of_flight', 'principles_of_flight', 'Équilibrer un virage en palier', 3, 3, 'Virage coordonné stabilisé en palier à 60°, cos 60° = 0,5. Facteur de charge ?

![CPL_081_001](/illustrations/cpl-081-001.svg)', 'Steady coordinated level turn at 60°, cos 60° = 0.5. Load factor?

![CPL_081_001](/illustrations/en/cpl-081-001.svg)', '["2", "0,5", "1", "4"]', '["2", "0.5", "1", "4"]', '2', 'n = 1/0,5 = 2.

![CPL_081_001](/illustrations/cpl-081-001.svg)', 'n = 1/0.5 = 2.

![CPL_081_001](/illustrations/en/cpl-081-001.svg)', '["CPL", "exam-mcq", "exam-subject:081"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('10ca81f6-0000-4000-8000-000010ca81f6', '2506b378-0000-4000-8000-00002506b378', 'CPL', 'principles_of_flight', 'principles_of_flight', 'Équilibrer un virage en palier', 3, 3, 'En virage en palier, quelle composante équilibre le poids ?', 'In a level turn, which component balances weight?', '["La composante horizontale seule", "La traînée seule", "La force centrifuge sans portance", "La composante verticale de portance"]', '["Horizontal component alone", "Drag alone", "Centrifugal force without lift", "Vertical lift component"]', 'La composante verticale de portance', 'L cos φ = W.

![CPL_081_001](/illustrations/cpl-081-001.svg)', 'L cos φ = W.

![CPL_081_001](/illustrations/en/cpl-081-001.svg)', '["CPL", "exam-mcq", "exam-subject:081"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('10ca81f7-0000-4000-8000-000010ca81f7', '2506b378-0000-4000-8000-00002506b378', 'CPL', 'principles_of_flight', 'principles_of_flight', 'Équilibrer un virage en palier', 3, 3, 'Virage coordonné stabilisé en palier à 45°, cos 45° ≈ 0,707. Facteur de charge n ?', 'Steady coordinated level turn at 45°, cos 45° ≈ 0.707. Load factor n?', '["2,000", "1,000", "1,414", "0,707"]', '["2.000", "1.000", "1.414", "0.707"]', '1,414', '1/0,707 ≈ 1,414.

![CPL_081_001](/illustrations/cpl-081-001.svg)', '1/0.707 ≈ 1.414.

![CPL_081_001](/illustrations/en/cpl-081-001.svg)', '["CPL", "exam-mcq", "exam-subject:081"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('10ca81f8-0000-4000-8000-000010ca81f8', '2506b378-0000-4000-8000-00002506b378', 'CPL', 'principles_of_flight', 'principles_of_flight', 'Équilibrer un virage en palier', 3, 3, 'La formule n = 1/cos φ s’applique directement à…', 'The formula n = 1/cos φ directly applies to…', '["Un avion au parking", "Un virage coordonné stabilisé en palier", "Toute inclinaison, même en chute libre", "Toute manœuvre sur le dos"]', '["A parked aircraft", "A steady coordinated level turn", "Any bank, even in free fall", "Every inverted manoeuvre"]', 'Un virage coordonné stabilisé en palier', 'Le maintien du palier est une hypothèse essentielle.

![CPL_081_001](/illustrations/cpl-081-001.svg)', 'Maintained level flight is an essential assumption.

![CPL_081_001](/illustrations/en/cpl-081-001.svg)', '["CPL", "exam-mcq", "exam-subject:081"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('10ca85ba-0000-4000-8000-000010ca85ba', '2506b378-0000-4000-8000-00002506b378', 'CPL', 'principles_of_flight', 'principles_of_flight', 'Relier charge et décrochage', 3, 3, 'Vs à 1 g = 50 kt. À n = 2, vitesse de décrochage approximative ?', '1 g Vs = 50 kt. At n = 2, approximate stall speed?', '["70,7 kt", "100 kt", "25 kt", "50 kt"]', '["70.7 kt", "100 kt", "25 kt", "50 kt"]', '70,7 kt', '50 × √2 = 70,7 kt.

![CPL_081_001](/illustrations/cpl-081-001.svg)', '50 × √2 = 70.7 kt.

![CPL_081_001](/illustrations/en/cpl-081-001.svg)', '["CPL", "exam-mcq", "exam-subject:081"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('10ca85bb-0000-4000-8000-000010ca85bb', '2506b378-0000-4000-8000-00002506b378', 'CPL', 'principles_of_flight', 'principles_of_flight', 'Relier charge et décrochage', 3, 3, 'À masse et configuration identiques, Vs à 1 g = 50 kt. À n = 1,414, vitesse de décrochage approximative ?', 'At unchanged mass and configuration, 1 g Vs = 50 kt. At n = 1.414, approximate stall speed?', '["70,7 kt", "35,4 kt", "100 kt", "59,5 kt"]', '["70.7 kt", "35.4 kt", "100 kt", "59.5 kt"]', '59,5 kt', '50 × √1,414 ≈ 59,5 kt. Multiplier directement par n donnerait 70,7 kt : la relation utilise sa racine. Elle ne donne pas la vitesse réelle de l’avion.

![CPL_081_001](/illustrations/cpl-081-001.svg)', '50 × √1.414 ≈ 59.5 kt. Multiplying directly by n gives 70.7 kt: the relation uses its square root. It does not determine actual aircraft speed.

![CPL_081_001](/illustrations/en/cpl-081-001.svg)', '["CPL", "exam-mcq", "exam-subject:081"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('10ca85bc-0000-4000-8000-000010ca85bc', '2506b378-0000-4000-8000-00002506b378', 'CPL', 'principles_of_flight', 'principles_of_flight', 'Relier charge et décrochage', 3, 3, 'Le déclencheur aérodynamique du décrochage est…', 'The aerodynamic trigger for stall is…', '["Une vitesse sol fixe", "Un cap magnétique précis", "L’incidence critique atteinte", "Une altitude fixe"]', '["A fixed groundspeed", "A specific magnetic heading", "Reaching critical angle of attack", "A fixed altitude"]', 'L’incidence critique atteinte', 'L’incidence critique peut être atteinte à plusieurs vitesses selon la charge.

![CPL_081_001](/illustrations/cpl-081-001.svg)', 'Critical angle can be reached at different speeds depending on load.

![CPL_081_001](/illustrations/en/cpl-081-001.svg)', '["CPL", "exam-mcq", "exam-subject:081"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('10ca85bd-0000-4000-8000-000010ca85bd', '2506b378-0000-4000-8000-00002506b378', 'CPL', 'principles_of_flight', 'principles_of_flight', 'Relier charge et décrochage', 3, 3, 'Une ressource augmente le facteur de charge. Le décrochage peut-il survenir au-dessus de Vs à 1 g ?', 'A pull-up increases load factor. Can stall occur above 1 g Vs?', '["Seulement à masse nulle", "Oui", "Non, jamais", "Seulement moteur arrêté"]', '["Only at zero mass", "Yes", "No, never", "Only with the engine stopped"]', 'Oui', 'Vs augmente avec √n, aux hypothèses données.

![CPL_081_001](/illustrations/cpl-081-001.svg)', 'Vs increases with √n under the stated assumptions.

![CPL_081_001](/illustrations/en/cpl-081-001.svg)', '["CPL", "exam-mcq", "exam-subject:081"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('10ca897f-0000-4000-8000-000010ca897f', '2506b378-0000-4000-8000-00002506b378', 'CPL', 'principles_of_flight', 'principles_of_flight', 'Masse et limites', 3, 3, 'Vs = 50 kt à 1 000 kg. À 1 210 kg, configuration identique, Vs approximative ?', 'Vs = 50 kt at 1,000 kg. At 1,210 kg, same configuration, approximate Vs?', '["55 kt", "60,5 kt", "45,5 kt", "50 kt"]', '["55 kt", "60.5 kt", "45.5 kt", "50 kt"]', '55 kt', '50 × √(1 210/1 000) = 55 kt.

![CPL_081_001](/illustrations/cpl-081-001.svg)', '50 × √(1,210/1,000) = 55 kt.

![CPL_081_001](/illustrations/en/cpl-081-001.svg)', '["CPL", "exam-mcq", "exam-subject:081"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('10ca8980-0000-4000-8000-000010ca8980', '2506b378-0000-4000-8000-00002506b378', 'CPL', 'principles_of_flight', 'principles_of_flight', 'Masse et limites', 3, 3, 'Une vitesse inférieure à la vitesse de manœuvre garantit-elle la protection contre toute sollicitation ?', 'Does speed below manoeuvring speed guarantee protection against every load input?', '["Oui, toutes combinaisons comprises", "Oui, même au-delà des limites de masse", "Oui, en turbulence sans restriction", "Non"]', '["Yes, including every combination", "Yes, even beyond mass limits", "Yes, without restriction in turbulence", "No"]', 'Non', 'Les protections ont des hypothèses et limites ; consulter le manuel.

![CPL_081_001](/illustrations/cpl-081-001.svg)', 'Protection has assumptions and limits; consult the manual.

![CPL_081_001](/illustrations/en/cpl-081-001.svg)', '["CPL", "exam-mcq", "exam-subject:081"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('0886a5b8-0000-4000-8000-00000886a5b8', '2506b378-0000-4000-8000-00002506b378', 'CPL', 'principles_of_flight', 'principles_of_flight', 'Masse et limites', 3, 3, 'À configuration donnée, réduire la masse tend à faire évoluer Vs comment ?', 'At a given configuration, how does reducing mass tend to affect Vs?', '["Elle devient nulle", "Elle reste nécessairement constante", "Elle diminue", "Elle augmente toujours"]', '["It becomes zero", "It must remain constant", "It decreases", "It always increases"]', 'Elle diminue', 'La vitesse de décrochage varie avec √masse.

![CPL_081_001](/illustrations/cpl-081-001.svg)', 'Stall speed varies with √mass.

![CPL_081_001](/illustrations/en/cpl-081-001.svg)', '["CPL", "exam-mcq", "exam-subject:081"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('08871a18-0000-4000-8000-000008871a18', '2506b378-0000-4000-8000-00002506b378', 'CPL', 'principles_of_flight', 'principles_of_flight', 'Gérer la marge en virage', 3, 3, 'Maintenir le palier en augmentant l’inclinaison exige généralement…', 'Maintaining level flight while increasing bank generally requires…', '["Une disparition de la traînée induite", "Davantage de portance totale", "Moins de portance totale", "Une portance nulle"]', '["Elimination of induced drag", "More total lift", "Less total lift", "Zero lift"]', 'Davantage de portance totale', 'La composante verticale doit rester égale au poids.

![CPL_081_001](/illustrations/cpl-081-001.svg)', 'The vertical component must remain equal to weight.

![CPL_081_001](/illustrations/en/cpl-081-001.svg)', '["CPL", "exam-mcq", "exam-subject:081"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('08871a19-0000-4000-8000-000008871a19', '2506b378-0000-4000-8000-00002506b378', 'CPL', 'principles_of_flight', 'principles_of_flight', 'Gérer la marge en virage', 3, 3, 'Un fort lacet dissymétrique près du décrochage augmente notamment le risque de…', 'Strong asymmetric yaw near stall particularly increases the risk of…', '["Départ en autorotation", "Suppression du décrochage", "Hausse garantie de la finesse", "Diminution certaine de l’incidence"]', '["Spin entry", "Elimination of stall", "Guaranteed higher glide ratio", "Certain reduction in angle of attack"]', 'Départ en autorotation', 'La dissymétrie au décrochage peut provoquer une autorotation.

![CPL_081_001](/illustrations/cpl-081-001.svg)', 'Asymmetry at stall may cause autorotation.

![CPL_081_001](/illustrations/en/cpl-081-001.svg)', '["CPL", "exam-mcq", "exam-subject:081"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('08871a1a-0000-4000-8000-000008871a1a', '2506b378-0000-4000-8000-00002506b378', 'CPL', 'principles_of_flight', 'principles_of_flight', 'Gérer la marge en virage', 3, 3, 'Quelle action aérodynamique est fondamentale pour sortir du décrochage ?', 'Which aerodynamic action is fundamental to stall recovery?', '["Augmenter systématiquement l’incidence", "Maintenir toute traction", "Attendre une vitesse sol donnée", "Réduire l’incidence"]', '["Always increase angle of attack", "Maintain all back pressure", "Wait for a specified groundspeed", "Reduce angle of attack"]', 'Réduire l’incidence', 'Il faut repasser sous l’incidence critique selon la procédure enseignée.

![CPL_081_001](/illustrations/cpl-081-001.svg)', 'Return below critical angle of attack using the trained procedure.

![CPL_081_001](/illustrations/en/cpl-081-001.svg)', '["CPL", "exam-mcq", "exam-subject:081"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('0886a5bc-0000-4000-8000-00000886a5bc', '2506b378-0000-4000-8000-00002506b378', 'CPL', 'principles_of_flight', 'principles_of_flight', 'Masse et limites', 3, 3, 'L’enveloppe de vol doit respecter…', 'The flight envelope must respect…', '["Uniquement l’incidence en croisière", "Seulement les limites de carburant", "Les limites aérodynamiques et structurales", "Seulement les vitesses sol"]', '["Only cruise angle of attack", "Only fuel limits", "Aerodynamic and structural limits", "Only groundspeeds"]', 'Les limites aérodynamiques et structurales', 'Les deux familles de limites sont distinctes et cumulatives.

![CPL_081_001](/illustrations/cpl-081-001.svg)', 'Both sets of limits are distinct and apply together.

![CPL_081_001](/illustrations/en/cpl-081-001.svg)', '["CPL", "exam-mcq", "exam-subject:081"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

COMMIT;
