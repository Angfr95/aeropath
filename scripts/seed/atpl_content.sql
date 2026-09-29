-- ATPL content only. Preserves accounts, progress and answer history.
BEGIN;
SET LOCAL standard_conforming_strings = on;

INSERT INTO lessons (id, license, category, theme, title_fr, title_en, content_fr, content_en, difficulty, order_index, level, duration_minutes, tags, learning_objectives) VALUES
('11ce5bda-0000-4000-8000-000011ce5bda', 'ATPL', 'aircraft_general', 'aircraft_general', 'Turbines : cycle et surveillance', 'Turbines: cycle and monitoring', '**Objectif :** Résoudre et vérifier les exercices de turbines : cycle et surveillance.

## Suivre le flux gazeux

![ATPL_021_001](/illustrations/atpl-021-001.svg)

Dans le cœur d’une turbine, l’air passe par compression, combustion, détente dans la turbine puis échappement. Le compresseur élève la pression et demande de la puissance mécanique ; la turbine extrait de l’énergie des gaz pour entraîner le compresseur et, selon le moteur, une soufflante ou une hélice. Le modèle idéal de Brayton comporte un apport de chaleur à pression constante ; un moteur réel présente des pertes, notamment dans la chambre de combustion.

## Identifier les architectures

Un turboréacteur utilise principalement le jet pour la poussée. Un turbofan fait passer une partie du débit autour du cœur ; le taux de dilution est débit secondaire / débit du cœur. Exemple : 800 kg/s autour du cœur pour 100 kg/s dans le cœur donnent 8. Un turbopropulseur fournit l’essentiel de sa propulsion par l’hélice. Les arbres basse et haute pression peuvent tourner à des vitesses différentes ; N1 et N2 désignent des ensembles propres à l’architecture du moteur.

## Reconnaître un démarrage anormal

Surveiller la séquence et les limites du manuel : rotation, allumage, température, accélération et pression d’huile. Un hot start correspond à une température de démarrage excessive ; un hung start à une accélération insuffisante, le régime restant sous le ralenti autonome attendu. Une absence d’allumage après admission de carburant peut accumuler du carburant non brûlé. Les seuils, délais et actions d’interruption sont propres au moteur ; ce chapitre n’en invente pas.

## Distinguer les anomalies en fonctionnement

Un décrochage compresseur est une perturbation aérodynamique de l’écoulement sur les aubages ; un pompage peut concerner l’ensemble du système de compression, avec oscillations et parfois inversion de débit. Bruits, variations de régime ou de température et perte de poussée sont des indices possibles, pas un diagnostic exclusif. Une extinction de flamme est distincte. Les corps étrangers peuvent endommager les aubes. Un FADEC peut assurer des protections, mais ne supprime pas les limitations ni la surveillance du pilote.

**Format de travail :** 15 QCM de chapitre, quatre choix et une réponse correcte. La matière 021 comporte 80 questions en 120 minutes dans le tableau EASA cité. Ce chapitre ne couvre qu’une partie de la matière ; il ne constitue pas une épreuve complète. Cadre EASA, même en anglais ou avec les unités US/britanniques.

## Références

- [EASA — programme Part-FCL](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8)
- [EASA — format de la matière, AMC1 ARA.FCL.300(b)](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67)
- [FAA PHAK — turbines et systèmes](https://www.faa.gov/regulationspolicies/handbooksmanuals/aviation/phak/chapter-7-aircraft-systems)

**Support de révision ATPL(A).** À faire valider par votre instructeur. Pour un vol réel, utiliser le manuel de l’avion et les publications en vigueur. Les exemples chiffrés sont pédagogiques.', '**Objective:** Solve and cross-check exercises on turbines: cycle and monitoring.

## Follow the gas path

![ATPL_021_001](/illustrations/en/atpl-021-001.svg)

In a turbine core, air undergoes compression, combustion, expansion through the turbine and exhaust. The compressor raises pressure and requires shaft power; the turbine extracts gas energy to drive the compressor and, depending on engine type, a fan or propeller. The ideal Brayton model adds heat at constant pressure; real engines have losses, including in the combustor.

## Identify architectures

A turbojet mainly uses its jet for thrust. A turbofan passes part of the airflow around the core; bypass ratio is bypass flow / core flow. Example: 800 kg/s around the core and 100 kg/s through it give 8. A turboprop obtains most propulsion from its propeller. Low- and high-pressure spools may rotate at different speeds; N1 and N2 refer to assemblies specific to engine architecture.

## Recognise an abnormal start

Monitor the manual’s sequence and limits: rotation, ignition, temperature, acceleration and oil pressure. A hot start involves excessive starting temperature; a hung start involves insufficient acceleration, with speed remaining below expected self-sustaining idle. Failure to light after fuel admission may accumulate unburned fuel. Limits, timings and abort actions are engine-specific; this chapter does not invent them.

## Distinguish running abnormalities

Compressor stall is aerodynamic flow breakdown over compressor blades; surge may involve the entire compression system, with oscillations and sometimes flow reversal. Noise, speed or temperature fluctuations and thrust loss are possible indications, not an exclusive diagnosis. Flameout is different. Foreign objects can damage blades. FADEC may provide protections but does not remove limitations or pilot monitoring.

**Practice format:** 15 chapter MCQs, four choices and one correct answer. Subject 021 has 80 questions in 120 minutes in the cited EASA table. This chapter covers only part of the subject and is not a full paper. EASA framework, including with English or US/British units.

## References

- [Reference 1](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8)
- [Reference 2](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67)
- [Reference 3](https://www.faa.gov/regulationspolicies/handbooksmanuals/aviation/phak/chapter-7-aircraft-systems)

**ATPL(A) revision material.** To be reviewed by your instructor. For an actual flight, use the aircraft manual and current publications. Numerical examples are for training.', 3, 1, 1, 35, '["ATPL", "mcq-four-options"]', '["Résoudre et vérifier les exercices de turbines : cycle et surveillance."]')
ON CONFLICT (id) DO UPDATE SET license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, title_fr = EXCLUDED.title_fr, title_en = EXCLUDED.title_en, content_fr = EXCLUDED.content_fr, content_en = EXCLUDED.content_en, difficulty = EXCLUDED.difficulty, order_index = EXCLUDED.order_index, level = EXCLUDED.level, duration_minutes = EXCLUDED.duration_minutes, tags = EXCLUDED.tags, learning_objectives = EXCLUDED.learning_objectives;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('7c3e9a59-0000-4000-8000-00007c3e9a59', '11ce5bda-0000-4000-8000-000011ce5bda', 'ATPL', 'aircraft_general', 'aircraft_general', 'Suivre le flux gazeux', 3, 3, 'Quel ordre décrit le flux dans le cœur d’une turbine ?', 'Which sequence describes flow through a turbine core?', '["Compression, combustion, turbine, échappement", "Combustion, échappement, compression, turbine", "Turbine, compression, échappement, combustion", "Échappement, combustion, turbine, compression"]', '["Compression, combustion, turbine, exhaust", "Combustion, exhaust, compression, turbine", "Turbine, compression, exhaust, combustion", "Exhaust, combustion, turbine, compression"]', 'Compression, combustion, turbine, échappement', 'La compression précède l’apport de chaleur, puis la détente.

![ATPL_021_001](/illustrations/atpl-021-001.svg)', 'Compression precedes heat addition, then expansion.

![ATPL_021_001](/illustrations/en/atpl-021-001.svg)', '["ATPL", "exam-mcq", "exam-subject:021"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('7c3e9a5a-0000-4000-8000-00007c3e9a5a', '11ce5bda-0000-4000-8000-000011ce5bda', 'ATPL', 'aircraft_general', 'aircraft_general', 'Suivre le flux gazeux', 3, 3, 'Quel organe extrait l’énergie des gaz pour entraîner le compresseur ?', 'Which component extracts gas energy to drive the compressor?', '["Le réservoir", "Le tube Pitot", "La prise statique", "La turbine"]', '["The tank", "The pitot tube", "The static port", "The turbine"]', 'La turbine', 'La turbine fournit du travail mécanique sur l’arbre.

![ATPL_021_001](/illustrations/atpl-021-001.svg)', 'The turbine supplies shaft work.

![ATPL_021_001](/illustrations/en/atpl-021-001.svg)', '["ATPL", "exam-mcq", "exam-subject:021"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('7c3e9a5b-0000-4000-8000-00007c3e9a5b', '11ce5bda-0000-4000-8000-000011ce5bda', 'ATPL', 'aircraft_general', 'aircraft_general', 'Suivre le flux gazeux', 3, 3, 'Le compresseur…', 'The compressor…', '["Est toujours entraîné seulement par le vent", "Remplace la combustion", "Consomme de la puissance pour augmenter la pression", "Crée du carburant"]', '["Is always driven only by wind", "Replaces combustion", "Consumes power to increase pressure", "Creates fuel"]', 'Consomme de la puissance pour augmenter la pression', 'L’augmentation de pression demande de l’énergie.

![ATPL_021_001](/illustrations/atpl-021-001.svg)', 'Raising pressure requires energy.

![ATPL_021_001](/illustrations/en/atpl-021-001.svg)', '["ATPL", "exam-mcq", "exam-subject:021"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('7c3e9a5c-0000-4000-8000-00007c3e9a5c', '11ce5bda-0000-4000-8000-000011ce5bda', 'ATPL', 'aircraft_general', 'aircraft_general', 'Suivre le flux gazeux', 3, 3, 'Dans le cycle idéal de Brayton, l’apport de chaleur est modélisé à…', 'In the ideal Brayton cycle, heat addition is modelled at…', '["Pression nécessairement nulle", "Pression constante", "Volume nul", "Température constante sans énergie"]', '["Necessarily zero pressure", "Constant pressure", "Zero volume", "Constant temperature without energy"]', 'Pression constante', 'C’est une idéalisation ; le moteur réel comporte des pertes.

![ATPL_021_001](/illustrations/atpl-021-001.svg)', 'This is an idealisation; real engines have losses.

![ATPL_021_001](/illustrations/en/atpl-021-001.svg)', '["ATPL", "exam-mcq", "exam-subject:021"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('7c3e9e1e-0000-4000-8000-00007c3e9e1e', '11ce5bda-0000-4000-8000-000011ce5bda', 'ATPL', 'aircraft_general', 'aircraft_general', 'Identifier les architectures', 3, 3, 'Débit secondaire 800 kg/s, débit du cœur 100 kg/s. Taux de dilution ?', 'Bypass flow 800 kg/s, core flow 100 kg/s. Bypass ratio?', '["8", "9", "0,125", "700"]', '["8", "9", "0.125", "700"]', '8', '800/100 = 8 ; ne pas utiliser le débit total au numérateur.

![ATPL_021_001](/illustrations/atpl-021-001.svg)', '800/100 = 8; do not use total flow as numerator.

![ATPL_021_001](/illustrations/en/atpl-021-001.svg)', '["ATPL", "exam-mcq", "exam-subject:021"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('7c3e9e1f-0000-4000-8000-00007c3e9e1f', '11ce5bda-0000-4000-8000-000011ce5bda', 'ATPL', 'aircraft_general', 'aircraft_general', 'Identifier les architectures', 3, 3, 'Quelle architecture fournit principalement sa propulsion par l’hélice ?', 'Which architecture mainly propels the aircraft with a propeller?', '["Turboréacteur pur", "Ramjet sans hélice", "Moteur-fusée", "Turbopropulseur"]', '["Pure turbojet", "Propeller-free ramjet", "Rocket engine", "Turboprop"]', 'Turbopropulseur', 'La puissance à l’arbre entraîne l’hélice.

![ATPL_021_001](/illustrations/atpl-021-001.svg)', 'Shaft power drives the propeller.

![ATPL_021_001](/illustrations/en/atpl-021-001.svg)', '["ATPL", "exam-mcq", "exam-subject:021"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('7c3e9e20-0000-4000-8000-00007c3e9e20', '11ce5bda-0000-4000-8000-000011ce5bda', 'ATPL', 'aircraft_general', 'aircraft_general', 'Identifier les architectures', 3, 3, 'Deux arbres indépendants d’un moteur à double corps doivent-ils tourner à la même vitesse ?', 'Must two independent spools of a two-spool engine rotate at equal speed?', '["Oui, dès l’allumage", "Oui, parce que N1 et N2 sont identiques", "Non", "Oui, toujours"]', '["Yes, from ignition", "Yes, because N1 and N2 are identical", "No", "Yes, always"]', 'Non', 'Les ensembles peuvent avoir des vitesses différentes.

![ATPL_021_001](/illustrations/atpl-021-001.svg)', 'The assemblies can have different speeds.

![ATPL_021_001](/illustrations/en/atpl-021-001.svg)', '["ATPL", "exam-mcq", "exam-subject:021"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('7c3ea1e2-0000-4000-8000-00007c3ea1e2', '11ce5bda-0000-4000-8000-000011ce5bda', 'ATPL', 'aircraft_general', 'aircraft_general', 'Reconnaître un démarrage anormal', 3, 3, 'Une température de démarrage dépasse sa limite. Quel terme décrit cette anomalie ?', 'Starting temperature exceeds its limit. Which term describes this abnormality?', '["Extinction nécessairement confirmée", "Hot start", "Hung start par définition", "Démarrage normal"]', '["Necessarily confirmed flameout", "Hot start", "Hung start by definition", "Normal start"]', 'Hot start', 'Le critère est ici la température excessive.

![ATPL_021_001](/illustrations/atpl-021-001.svg)', 'The criterion here is excessive temperature.

![ATPL_021_001](/illustrations/en/atpl-021-001.svg)', '["ATPL", "exam-mcq", "exam-subject:021"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('7c3ea1e3-0000-4000-8000-00007c3ea1e3', '11ce5bda-0000-4000-8000-000011ce5bda', 'ATPL', 'aircraft_general', 'aircraft_general', 'Reconnaître un démarrage anormal', 3, 3, 'Le moteur s’allume mais son régime stagne sous le ralenti autonome attendu. Quelle anomalie évoquer ?', 'The engine lights but speed stagnates below expected self-sustaining idle. Which abnormality is suggested?', '["Hung start", "Démarrage normal achevé", "Survitesse confirmée", "Simple variation de QNH"]', '["Hung start", "Completed normal start", "Confirmed overspeed", "Simple QNH change"]', 'Hung start', 'L’accélération est insuffisante pour achever le démarrage.

![ATPL_021_001](/illustrations/atpl-021-001.svg)', 'Acceleration is insufficient to complete the start.

![ATPL_021_001](/illustrations/en/atpl-021-001.svg)', '["ATPL", "exam-mcq", "exam-subject:021"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('7c3ea1e4-0000-4000-8000-00007c3ea1e4', '11ce5bda-0000-4000-8000-000011ce5bda', 'ATPL', 'aircraft_general', 'aircraft_general', 'Reconnaître un démarrage anormal', 3, 3, 'Les délais d’allumage et températures limites viennent…', 'Ignition time limits and temperature limits come from…', '["D’une valeur universelle pour toutes les turbines", "Du nombre de passagers", "De la langue du cockpit", "Du manuel propre au moteur et à l’avion"]', '["One universal value for all turbines", "Passenger count", "Cockpit language", "The engine- and aircraft-specific manual"]', 'Du manuel propre au moteur et à l’avion', 'Les séquences et protections varient selon l’installation.

![ATPL_021_001](/illustrations/atpl-021-001.svg)', 'Sequences and protections vary with installation.

![ATPL_021_001](/illustrations/en/atpl-021-001.svg)', '["ATPL", "exam-mcq", "exam-subject:021"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('0b9599d4-0000-4000-8000-00000b9599d4', '11ce5bda-0000-4000-8000-000011ce5bda', 'ATPL', 'aircraft_general', 'aircraft_general', 'Reconnaître un démarrage anormal', 3, 3, 'L’absence d’allumage après admission de carburant peut entraîner…', 'No light-up after fuel admission may cause…', '["Une température extérieure plus basse", "Un compresseur sans besoin d’énergie", "Une accumulation de carburant non brûlé", "Une combustion nécessairement normale"]', '["Lower outside temperature", "A compressor requiring no energy", "Accumulation of unburned fuel", "Necessarily normal combustion"]', 'Une accumulation de carburant non brûlé', 'Appliquer la procédure propre au moteur ; ne pas prolonger arbitrairement.

![ATPL_021_001](/illustrations/atpl-021-001.svg)', 'Follow the engine-specific procedure; do not arbitrarily extend the start.

![ATPL_021_001](/illustrations/en/atpl-021-001.svg)', '["ATPL", "exam-mcq", "exam-subject:021"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('0b960e34-0000-4000-8000-00000b960e34', '11ce5bda-0000-4000-8000-000011ce5bda', 'ATPL', 'aircraft_general', 'aircraft_general', 'Distinguer les anomalies en fonctionnement', 3, 3, 'Un décrochage compresseur concerne d’abord…', 'Compressor stall primarily concerns…', '["Le calage de l’altimètre", "L’écoulement aérodynamique sur les aubages", "Le décrochage de l’aile uniquement", "La pression des pneus"]', '["Altimeter setting", "Aerodynamic flow over blades", "Only wing stall", "Tyre pressure"]', 'L’écoulement aérodynamique sur les aubages', 'Le terme désigne une instabilité du compresseur.

![ATPL_021_001](/illustrations/atpl-021-001.svg)', 'The term describes compressor instability.

![ATPL_021_001](/illustrations/en/atpl-021-001.svg)', '["ATPL", "exam-mcq", "exam-subject:021"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('0b960e35-0000-4000-8000-00000b960e35', '11ce5bda-0000-4000-8000-000011ce5bda', 'ATPL', 'aircraft_general', 'aircraft_general', 'Distinguer les anomalies en fonctionnement', 3, 3, 'Une extinction de flamme est-elle synonyme de tout décrochage compresseur ?', 'Is flameout synonymous with every compressor stall?', '["Non", "Oui", "Oui, pour N1 seulement", "Oui, au ralenti seulement"]', '["No", "Yes", "Yes, for N1 only", "Yes, only at idle"]', 'Non', 'Ce sont des phénomènes distincts même s’ils peuvent être liés.

![ATPL_021_001](/illustrations/atpl-021-001.svg)', 'They are distinct phenomena even if they may be related.

![ATPL_021_001](/illustrations/en/atpl-021-001.svg)', '["ATPL", "exam-mcq", "exam-subject:021"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('0b960e36-0000-4000-8000-00000b960e36', '11ce5bda-0000-4000-8000-000011ce5bda', 'ATPL', 'aircraft_general', 'aircraft_general', 'Distinguer les anomalies en fonctionnement', 3, 3, 'Pourquoi surveiller l’entrée d’air et les corps étrangers ?', 'Why monitor the intake and foreign objects?', '["Ils améliorent le taux de compression", "Ils remplacent le filtre à carburant", "Ils annulent le risque de pompage", "Ils peuvent endommager les aubes"]', '["They improve compression ratio", "They replace the fuel filter", "They eliminate surge risk", "They can damage blades"]', 'Ils peuvent endommager les aubes', 'L’ingestion peut affecter l’intégrité et les performances.

![ATPL_021_001](/illustrations/atpl-021-001.svg)', 'Ingestion can affect integrity and performance.

![ATPL_021_001](/illustrations/en/atpl-021-001.svg)', '["ATPL", "exam-mcq", "exam-subject:021"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('0b960e37-0000-4000-8000-00000b960e37', '11ce5bda-0000-4000-8000-000011ce5bda', 'ATPL', 'aircraft_general', 'aircraft_general', 'Distinguer les anomalies en fonctionnement', 3, 3, 'La présence d’un FADEC dispense-t-elle de respecter les limitations et surveiller le moteur ?', 'Does FADEC remove the need to observe limits and monitor the engine?', '["Oui, au démarrage seulement", "Oui, si N1 est affiché", "Non", "Oui, pour toutes les phases"]', '["Yes, only during start", "Yes, if N1 is displayed", "No", "Yes, in all phases"]', 'Non', 'Les protections automatiques ne suppriment pas les obligations de surveillance.

![ATPL_021_001](/illustrations/atpl-021-001.svg)', 'Automatic protections do not remove monitoring duties.

![ATPL_021_001](/illustrations/en/atpl-021-001.svg)', '["ATPL", "exam-mcq", "exam-subject:021"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO lessons (id, license, category, theme, title_fr, title_en, content_fr, content_en, difficulty, order_index, level, duration_minutes, tags, learning_objectives) VALUES
('11ce5bdb-0000-4000-8000-000011ce5bdb', 'ATPL', 'aircraft_general', 'aircraft_general', 'Pressurisation : pression, altitude cabine et surveillance', 'Pressurisation: pressure, cabin altitude and monitoring', '**Objectif :** Résoudre et vérifier les exercices de pressurisation : pression, altitude cabine et surveillance.

## Trois indications différentes

![ATPL_021_002](/illustrations/atpl-021-002.svg)

L’altitude cabine est l’altitude de l’atmosphère standard correspondant à la pression dans la cabine. Ce n’est ni l’altitude réelle de l’avion ni une mesure de sa hauteur au-dessus du terrain. La pression différentielle est Δp = p cabine − p extérieure, dans la même unité. Le variomètre cabine indique la vitesse de variation de l’altitude cabine. Exemple fictif : 75 kPa à l’intérieur et 30 kPa à l’extérieur donnent 45 kPa de différentiel.

## Réguler le débit de sortie

Sur un système classique alimenté en air, le contrôleur agit sur la vanne de décharge (outflow valve) pour régler la pression cabine. À alimentation et conditions extérieures comparables, ouvrir davantage la sortie tend à réduire la pression cabine et à augmenter son altitude ; fermer la sortie produit la tendance inverse. Ce raisonnement suppose un débit entrant disponible et ne remplace pas le schéma propre à l’avion. Ventilation, température et pressurisation sont liées mais ne mesurent pas la même fonction.

## Respecter les limites de structure

En montée de l’avion, une pression cabine constante et une pression extérieure décroissante augmentent Δp. Une fois la limite différentielle atteinte, le système ne peut pas conserver indéfiniment la même altitude cabine : celle-ci augmente si l’avion continue sa montée. La soupape de surpression limite un différentiel positif excessif ; une protection de pression négative évite un différentiel extérieur supérieur à l’intérieur au-delà des limites prévues. Les valeurs limites sont propres au type.

## Surveiller et réagir

Une perte de pressurisation peut accroître rapidement l’altitude cabine et provoquer une hypoxie. La réponse suit les procédures de l’avion : utilisation rapide de l’oxygène si requise, maîtrise de la trajectoire et descente d’urgence lorsque nécessaire, en tenant compte du relief et des contraintes pertinentes. Il n’existe pas de seuil universel inventé pour ce chapitre. Une pression cabine proche de la pression extérieure ne prouve pas que l’air est respirable à haute altitude. Avant ouverture des portes, vérifier la dépressurisation selon la procédure.

**Format de travail :** 15 QCM de chapitre, quatre choix et une réponse correcte. La matière 021 comporte 80 questions en 120 minutes dans le tableau EASA cité. Ce chapitre ne couvre qu’une partie de la matière ; il ne constitue pas une épreuve complète. Cadre EASA, même en anglais ou avec les unités US/britanniques.

## Références

- [EASA — programme Part-FCL](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8)
- [EASA — format de la matière, AMC1 ARA.FCL.300(b)](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67)
- [FAA PHAK — aircraft systems](https://www.faa.gov/sites/faa.gov/files/09_phak_ch7.pdf)

**Support de révision ATPL(A).** À faire valider par votre instructeur. Pour un vol réel, utiliser le manuel de l’avion et les publications en vigueur. Les exemples chiffrés sont pédagogiques.', '**Objective:** Solve and cross-check exercises on pressurisation: pressure, cabin altitude and monitoring.

## Three different indications

![ATPL_021_002](/illustrations/en/atpl-021-002.svg)

Cabin altitude is the standard-atmosphere altitude corresponding to cabin pressure. It is neither the aircraft’s true altitude nor its height above terrain. Differential pressure is Δp = cabin pressure − outside pressure, using the same unit. Cabin vertical speed indicates the rate of change of cabin altitude. Fictional example: 75 kPa inside and 30 kPa outside give a 45 kPa differential.

## Regulate outflow

In a conventional air-supplied system, the controller operates the outflow valve to regulate cabin pressure. With comparable supply and external conditions, opening the outlet further tends to reduce cabin pressure and increase cabin altitude; closing it produces the opposite tendency. This reasoning assumes available inflow and does not replace the aircraft-specific schematic. Ventilation, temperature and pressurisation are related but are not the same function.

## Respect structural limits

As the aircraft climbs, constant cabin pressure and decreasing outside pressure increase Δp. Once the differential limit is reached, the system cannot indefinitely maintain the same cabin altitude: cabin altitude increases if the aircraft continues climbing. A positive-pressure relief valve limits excessive positive differential; negative-pressure protection prevents outside pressure exceeding inside pressure beyond the intended limits. Limit values are type-specific.

## Monitor and respond

Loss of pressurisation can rapidly increase cabin altitude and cause hypoxia. Follow aircraft procedures: prompt oxygen use when required, flight-path control and emergency descent when necessary, considering terrain and relevant constraints. This chapter invents no universal threshold. Cabin pressure close to outside pressure does not establish breathable conditions at high altitude. Before opening doors, verify depressurisation according to procedure.

**Practice format:** 15 chapter MCQs, four choices and one correct answer. Subject 021 has 80 questions in 120 minutes in the cited EASA table. This chapter covers only part of the subject and is not a full paper. EASA framework, including with English or US/British units.

## References

- [Reference 1](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8)
- [Reference 2](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67)
- [Reference 3](https://www.faa.gov/sites/faa.gov/files/09_phak_ch7.pdf)

**ATPL(A) revision material.** To be reviewed by your instructor. For an actual flight, use the aircraft manual and current publications. Numerical examples are for training.', 3, 2, 1, 35, '["ATPL", "mcq-four-options"]', '["Résoudre et vérifier les exercices de pressurisation : pression, altitude cabine et surveillance."]')
ON CONFLICT (id) DO UPDATE SET license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, title_fr = EXCLUDED.title_fr, title_en = EXCLUDED.title_en, content_fr = EXCLUDED.content_fr, content_en = EXCLUDED.content_en, difficulty = EXCLUDED.difficulty, order_index = EXCLUDED.order_index, level = EXCLUDED.level, duration_minutes = EXCLUDED.duration_minutes, tags = EXCLUDED.tags, learning_objectives = EXCLUDED.learning_objectives;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('18c6abf9-0000-4000-8000-000018c6abf9', '11ce5bdb-0000-4000-8000-000011ce5bdb', 'ATPL', 'aircraft_general', 'aircraft_general', 'Trois indications différentes', 3, 3, 'Pression cabine 75 kPa, extérieure 30 kPa. Quel différentiel ?

![ATPL_021_002](/illustrations/atpl-021-002.svg)', 'Cabin pressure 75 kPa, outside 30 kPa. What differential?

![ATPL_021_002](/illustrations/en/atpl-021-002.svg)', '["45 kPa", "105 kPa", "30 kPa", "2,5 kPa"]', '["45 kPa", "105 kPa", "30 kPa", "2.5 kPa"]', '45 kPa', 'Δp = 75 − 30 = 45 kPa.

![ATPL_021_002](/illustrations/atpl-021-002.svg)', 'Δp = 75 − 30 = 45 kPa.

![ATPL_021_002](/illustrations/en/atpl-021-002.svg)', '["ATPL", "exam-mcq", "exam-subject:021"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('18c6abfa-0000-4000-8000-000018c6abfa', '11ce5bdb-0000-4000-8000-000011ce5bdb', 'ATPL', 'aircraft_general', 'aircraft_general', 'Trois indications différentes', 3, 3, 'L’altitude cabine représente…', 'Cabin altitude represents…', '["La hauteur radio de l’avion", "L’altitude vraie de l’avion", "La distance au terrain de destination", "Une altitude standard équivalente à la pression cabine"]', '["Aircraft radio height", "Aircraft true altitude", "Distance to destination", "A standard altitude equivalent to cabin pressure"]', 'Une altitude standard équivalente à la pression cabine', 'Elle est déduite de la pression cabine selon l’atmosphère standard.

![ATPL_021_002](/illustrations/atpl-021-002.svg)', 'It is derived from cabin pressure using the standard atmosphere.

![ATPL_021_002](/illustrations/en/atpl-021-002.svg)', '["ATPL", "exam-mcq", "exam-subject:021"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('18c6abfb-0000-4000-8000-000018c6abfb', '11ce5bdb-0000-4000-8000-000011ce5bdb', 'ATPL', 'aircraft_general', 'aircraft_general', 'Trois indications différentes', 3, 3, 'Une pression cabine qui diminue correspond à une altitude cabine qui…', 'Decreasing cabin pressure corresponds to cabin altitude that…', '["Reste toujours nulle", "Devient une vitesse", "Augmente", "Diminue"]', '["Always stays zero", "Becomes a speed", "Increases", "Decreases"]', 'Augmente', 'Dans l’atmosphère standard, une pression plus faible correspond à une altitude plus élevée.

![ATPL_021_002](/illustrations/atpl-021-002.svg)', 'In the standard atmosphere, lower pressure corresponds to higher altitude.

![ATPL_021_002](/illustrations/en/atpl-021-002.svg)', '["ATPL", "exam-mcq", "exam-subject:021"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('18c6abfc-0000-4000-8000-000018c6abfc', '11ce5bdb-0000-4000-8000-000011ce5bdb', 'ATPL', 'aircraft_general', 'aircraft_general', 'Trois indications différentes', 3, 3, 'Le variomètre cabine indique…', 'The cabin vertical-speed indicator shows…', '["Le débit carburant", "La variation d’altitude cabine par unité de temps", "La vitesse sol", "Le différentiel en kPa"]', '["Fuel flow", "Cabin-altitude change per unit time", "Groundspeed", "Differential pressure in kPa"]', 'La variation d’altitude cabine par unité de temps', 'Il mesure une évolution de l’altitude cabine.

![ATPL_021_002](/illustrations/atpl-021-002.svg)', 'It measures change in cabin altitude.

![ATPL_021_002](/illustrations/en/atpl-021-002.svg)', '["ATPL", "exam-mcq", "exam-subject:021"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('18c6afbe-0000-4000-8000-000018c6afbe', '11ce5bdb-0000-4000-8000-000011ce5bdb', 'ATPL', 'aircraft_general', 'aircraft_general', 'Réguler le débit de sortie', 3, 3, 'À alimentation comparable, ouvrir davantage l’outflow valve tend à…', 'With comparable supply, opening the outflow valve further tends to…', '["Réduire la pression cabine", "Augmenter la pression cabine", "Bloquer toute sortie d’air", "Augmenter le débit carburant directement"]', '["Reduce cabin pressure", "Increase cabin pressure", "Block all air outflow", "Directly increase fuel flow"]', 'Réduire la pression cabine', 'Une sortie accrue tend à diminuer la pression intérieure.

![ATPL_021_002](/illustrations/atpl-021-002.svg)', 'Increased outflow tends to decrease internal pressure.

![ATPL_021_002](/illustrations/en/atpl-021-002.svg)', '["ATPL", "exam-mcq", "exam-subject:021"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('18c6afbf-0000-4000-8000-000018c6afbf', '11ce5bdb-0000-4000-8000-000011ce5bdb', 'ATPL', 'aircraft_general', 'aircraft_general', 'Réguler le débit de sortie', 3, 3, 'La fermeture de la sortie, avec un débit entrant disponible, tend à…', 'Closing the outlet, with inflow available, tends to…', '["Augmenter l’altitude cabine", "Supprimer le débit entrant par définition", "Égaliser immédiatement les pressions", "Augmenter la pression cabine"]', '["Increase cabin altitude", "Remove inflow by definition", "Immediately equalise pressures", "Increase cabin pressure"]', 'Augmenter la pression cabine', 'La rétention d’air tend à accroître la pression et réduire l’altitude cabine.

![ATPL_021_002](/illustrations/atpl-021-002.svg)', 'Retaining air tends to increase pressure and reduce cabin altitude.

![ATPL_021_002](/illustrations/en/atpl-021-002.svg)', '["ATPL", "exam-mcq", "exam-subject:021"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('18c6afc0-0000-4000-8000-000018c6afc0', '11ce5bdb-0000-4000-8000-000011ce5bdb', 'ATPL', 'aircraft_general', 'aircraft_general', 'Réguler le débit de sortie', 3, 3, 'Une température cabine confortable prouve-t-elle une pressurisation correcte ?', 'Does comfortable cabin temperature prove correct pressurisation?', '["Oui, au-dessus de tout relief", "Oui, avec ventilation audible", "Non", "Oui, toujours"]', '["Yes, above all terrain", "Yes, with audible ventilation", "No", "Yes, always"]', 'Non', 'Température et pression sont des grandeurs distinctes.

![ATPL_021_002](/illustrations/atpl-021-002.svg)', 'Temperature and pressure are distinct quantities.

![ATPL_021_002](/illustrations/en/atpl-021-002.svg)', '["ATPL", "exam-mcq", "exam-subject:021"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('18c6b382-0000-4000-8000-000018c6b382', '11ce5bdb-0000-4000-8000-000011ce5bdb', 'ATPL', 'aircraft_general', 'aircraft_general', 'Respecter les limites de structure', 3, 3, 'Cabine maintenue à 75 kPa, extérieur passant de 30 à 25 kPa. Nouveau Δp ?', 'Cabin held at 75 kPa, outside falling from 30 to 25 kPa. New Δp?', '["25 kPa", "50 kPa", "45 kPa", "100 kPa"]', '["25 kPa", "50 kPa", "45 kPa", "100 kPa"]', '50 kPa', '75 − 25 = 50 kPa : le différentiel augmente.

![ATPL_021_002](/illustrations/atpl-021-002.svg)', '75 − 25 = 50 kPa: differential increases.

![ATPL_021_002](/illustrations/en/atpl-021-002.svg)', '["ATPL", "exam-mcq", "exam-subject:021"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('18c6b383-0000-4000-8000-000018c6b383', '11ce5bdb-0000-4000-8000-000011ce5bdb', 'ATPL', 'aircraft_general', 'aircraft_general', 'Respecter les limites de structure', 3, 3, 'À différentiel maximal, l’avion continue de monter. L’altitude cabine peut devoir…', 'At maximum differential, the aircraft keeps climbing. Cabin altitude may need to…', '["Augmenter", "Rester toujours constante", "Devenir négative nécessairement", "Suivre la vitesse indiquée"]', '["Increase", "Always remain constant", "Necessarily become negative", "Follow indicated airspeed"]', 'Augmenter', 'La pression cabine doit diminuer pour respecter Δp lorsque la pression extérieure baisse.

![ATPL_021_002](/illustrations/atpl-021-002.svg)', 'Cabin pressure must decrease to respect Δp as outside pressure falls.

![ATPL_021_002](/illustrations/en/atpl-021-002.svg)', '["ATPL", "exam-mcq", "exam-subject:021"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('18c6b384-0000-4000-8000-000018c6b384', '11ce5bdb-0000-4000-8000-000011ce5bdb', 'ATPL', 'aircraft_general', 'aircraft_general', 'Respecter les limites de structure', 3, 3, 'La soupape de surpression positive protège contre…', 'The positive-pressure relief valve protects against…', '["Toute hypoxie sans oxygène", "Une vitesse sol excessive", "Un défaut de navigation", "Un différentiel positif excessif"]', '["All hypoxia without oxygen", "Excessive groundspeed", "A navigation error", "Excessive positive differential"]', 'Un différentiel positif excessif', 'Sa fonction est de limiter la charge de pression prévue par la conception.

![ATPL_021_002](/illustrations/atpl-021-002.svg)', 'Its function is to limit the pressure load specified by design.

![ATPL_021_002](/illustrations/en/atpl-021-002.svg)', '["ATPL", "exam-mcq", "exam-subject:021"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('000fbc34-0000-4000-8000-0000000fbc34', '11ce5bdb-0000-4000-8000-000011ce5bdb', 'ATPL', 'aircraft_general', 'aircraft_general', 'Respecter les limites de structure', 3, 3, 'Cabine 60 kPa, extérieur 65 kPa : Δp vaut…', 'Cabin 60 kPa, outside 65 kPa: Δp is…', '["125 kPa", "65 kPa", "−5 kPa", "5 kPa"]', '["125 kPa", "65 kPa", "−5 kPa", "5 kPa"]', '−5 kPa', '60 − 65 = −5 kPa : différentiel négatif.

![ATPL_021_002](/illustrations/atpl-021-002.svg)', '60 − 65 = −5 kPa: negative differential.

![ATPL_021_002](/illustrations/en/atpl-021-002.svg)', '["ATPL", "exam-mcq", "exam-subject:021"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('00103094-0000-4000-8000-000000103094', '11ce5bdb-0000-4000-8000-000011ce5bdb', 'ATPL', 'aircraft_general', 'aircraft_general', 'Surveiller et réagir', 3, 3, 'Une hausse rapide d’altitude cabine en croisière peut signaler…', 'A rapid cabin-altitude increase in cruise may indicate…', '["Un fonctionnement toujours normal", "Une perte de pressurisation", "Une meilleure oxygénation garantie", "Une baisse certaine d’altitude avion"]', '["Always normal operation", "Loss of pressurisation", "Guaranteed better oxygenation", "A definite decrease in aircraft altitude"]', 'Une perte de pressurisation', 'L’altitude cabine augmente lorsque la pression cabine diminue.

![ATPL_021_002](/illustrations/atpl-021-002.svg)', 'Cabin altitude rises as cabin pressure falls.

![ATPL_021_002](/illustrations/en/atpl-021-002.svg)', '["ATPL", "exam-mcq", "exam-subject:021"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('00103095-0000-4000-8000-000000103095', '11ce5bdb-0000-4000-8000-000011ce5bdb', 'ATPL', 'aircraft_general', 'aircraft_general', 'Surveiller et réagir', 3, 3, 'En dépressurisation, quelle référence guide les actions immédiates ?', 'During depressurisation, what guides immediate actions?', '["Les procédures du type d’avion", "Un seuil unique valable pour tous les avions", "Le confort thermique seul", "La vitesse sol seule"]', '["Aircraft-type procedures", "One threshold valid for every aircraft", "Thermal comfort alone", "Groundspeed alone"]', 'Les procédures du type d’avion', 'Les systèmes, alertes et procédures dépendent de l’avion.

![ATPL_021_002](/illustrations/atpl-021-002.svg)', 'Systems, warnings and procedures depend on the aircraft.

![ATPL_021_002](/illustrations/en/atpl-021-002.svg)', '["ATPL", "exam-mcq", "exam-subject:021"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('00103096-0000-4000-8000-000000103096', '11ce5bdb-0000-4000-8000-000011ce5bdb', 'ATPL', 'aircraft_general', 'aircraft_general', 'Surveiller et réagir', 3, 3, 'À haute altitude, Δp proche de zéro garantit-il un air suffisamment pressurisé pour les occupants ?', 'At high altitude, does Δp near zero guarantee adequate cabin pressure for occupants?', '["Oui, toujours", "Oui, si les portes sont fermées", "Oui, si le chauffage fonctionne", "Non"]', '["Yes, always", "Yes, if doors are closed", "Yes, if heating works", "No"]', 'Non', 'Un différentiel nul peut signifier une cabine à la faible pression extérieure.

![ATPL_021_002](/illustrations/atpl-021-002.svg)', 'Zero differential can mean the cabin is at low outside pressure.

![ATPL_021_002](/illustrations/en/atpl-021-002.svg)', '["ATPL", "exam-mcq", "exam-subject:021"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('00103097-0000-4000-8000-000000103097', '11ce5bdb-0000-4000-8000-000011ce5bdb', 'ATPL', 'aircraft_general', 'aircraft_general', 'Surveiller et réagir', 3, 3, 'Avant l’ouverture d’une porte au sol, il faut notamment vérifier…', 'Before opening a door on the ground, verify in particular…', '["Seulement le cap parking", "Uniquement le niveau carburant", "La dépressurisation selon la procédure", "Uniquement la température extérieure"]', '["Only parking heading", "Only fuel level", "Depressurisation according to procedure", "Only outside temperature"]', 'La dépressurisation selon la procédure', 'Une pression résiduelle peut présenter un danger lors de l’ouverture.

![ATPL_021_002](/illustrations/atpl-021-002.svg)', 'Residual pressure may create a hazard when opening.

![ATPL_021_002](/illustrations/en/atpl-021-002.svg)', '["ATPL", "exam-mcq", "exam-subject:021"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO lessons (id, license, category, theme, title_fr, title_en, content_fr, content_en, difficulty, order_index, level, duration_minutes, tags, learning_objectives) VALUES
('1db230b2-0000-4000-8000-00001db230b2', 'ATPL', 'communications', 'communications', 'Clairances et boucle de collationnement', 'Clearances and the readback loop', '**Objectif :** Reconstruire une clairance après messages successifs, détecter omissions et substitutions et préserver conditions et limites au collationnement.

## Identifier le destinataire

![ATPL_090_001](/illustrations/atpl-090-001.svg)

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

**Support de révision ATPL(A).** À faire valider par votre instructeur. Pour un vol réel, utiliser le manuel de l’avion et les publications en vigueur. Les exemples chiffrés sont pédagogiques.', '**Objective:** Reconstruct clearance after successive messages, detect omissions and substitutions and preserve conditions and limits in readback.

## Identify the addressee

![ATPL_090_001](/illustrations/en/atpl-090-001.svg)

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

**ATPL(A) revision material.** To be reviewed by your instructor. For an actual flight, use the aircraft manual and current publications. Numerical examples are for training.', 3, 1, 1, 45, '["ATPL", "mcq-four-options"]', '["Reconstruire une clairance après messages successifs, détecter omissions et substitutions et préserver conditions et limites au collationnement."]')
ON CONFLICT (id) DO UPDATE SET license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, title_fr = EXCLUDED.title_fr, title_en = EXCLUDED.title_en, content_fr = EXCLUDED.content_fr, content_en = EXCLUDED.content_en, difficulty = EXCLUDED.difficulty, order_index = EXCLUDED.order_index, level = EXCLUDED.level, duration_minutes = EXCLUDED.duration_minutes, tags = EXCLUDED.tags, learning_objectives = EXCLUDED.learning_objectives;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('3e85f959-0000-4000-8000-00003e85f959', '1db230b2-0000-4000-8000-00001db230b2', 'ATPL', 'communications', 'communications', 'Identifier le destinataire', 3, 3, 'Vous êtes AER452, cap 240. L''ATC dit « AER425, cap 270 ». Quel état retenir pour AER452 ?', 'You are AER452, heading 240. ATC says “AER425, heading 270”. Which state applies to AER452?', '["Cap 240, instruction adressée à un autre avion", "Cap 270, dernière instruction sur la fréquence", "Cap 270 après collationnement avec AER452", "Cap 255, moyenne en attendant confirmation"]', '["Heading 240, instruction addressed to another aircraft", "Heading 270, latest instruction on frequency", "Heading 270 after reading back as AER452", "Heading 255, average pending confirmation"]', 'Cap 240, instruction adressée à un autre avion', 'Les chiffres de l''indicatif sont transposés : le destinataire est différent. Un collationnement fait au nom d''AER452 ne transfère pas la clairance ; aucune moyenne n''est autorisée.

![ATPL_090_001](/illustrations/atpl-090-001.svg)', 'The callsign digits are transposed: this is another addressee. Reading back as AER452 does not transfer clearance; no averaged heading was assigned.

![ATPL_090_001](/illustrations/en/atpl-090-001.svg)', '["ATPL", "exam-mcq", "exam-subject:090"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('3e85f95a-0000-4000-8000-00003e85f95a', '1db230b2-0000-4000-8000-00001db230b2', 'ATPL', 'communications', 'communications', 'Identifier le destinataire', 3, 3, 'Message : AER452, cap 260, descendez [masqué], QNH 1006. Quelle information doit être demandée avant d''utiliser le nouveau niveau ?', 'Message: AER452, heading 260, descend [blocked], QNH 1006. Which information must be requested before using the new level?', '["Confirmation de 3 000 ft parce que prévu au briefing", "Confirmation du QNH déjà reçu uniquement", "Le niveau du trafic précédent comme substitut", "La valeur du niveau de descente"]', '["Confirmation of 3,000 ft because it was briefed", "Only confirmation of the already received QNH", "The preceding aircraft''s level as a substitute", "The descent-level value"]', 'La valeur du niveau de descente', 'Le niveau n''a pas été reçu. Le QNH et le cap ne permettent pas de le déduire ; une attente de briefing ne fournit pas de clairance. Demander la répétition de l''élément manquant.

![ATPL_090_001](/illustrations/atpl-090-001.svg)', 'The level was not received. QNH and heading cannot determine it; a briefing expectation supplies no clearance. Request repetition of the missing item.

![ATPL_090_001](/illustrations/en/atpl-090-001.svg)', '["ATPL", "exam-mcq", "exam-subject:090"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('3e85f95b-0000-4000-8000-00003e85f95b', '1db230b2-0000-4000-8000-00001db230b2', 'ATPL', 'communications', 'communications', 'Identifier le destinataire', 3, 3, 'AER452 demande une montée ; réponse STANDBY. Quelle mise à jour du niveau autorisé inscrire ?', 'AER452 requests a climb; reply is STANDBY. What update belongs in the cleared-level record?', '["Inscrire le niveau demandé après trente secondes", "Considérer toute future montée comme refusée", "Aucune, conserver le niveau précédemment autorisé", "Remplacer par le niveau demandé"]', '["Enter the requested level after thirty seconds", "Treat any future climb as refused", "None, retain the previously cleared level", "Replace it with the requested level"]', 'Aucune, conserver le niveau précédemment autorisé', 'STANDBY laisse la demande en attente, sans nouveau niveau ni délai automatique d''acceptation. Ce n''est pas non plus un refus définitif.

![ATPL_090_001](/illustrations/atpl-090-001.svg)', 'STANDBY leaves the request pending, without a new level or automatic acceptance delay. Nor is it a final denial.

![ATPL_090_001](/illustrations/en/atpl-090-001.svg)', '["ATPL", "exam-mcq", "exam-subject:090"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('3e85fd1d-0000-4000-8000-00003e85fd1d', '1db230b2-0000-4000-8000-00001db230b2', 'ATPL', 'communications', 'communications', 'Collationner les éléments de sécurité', 3, 3, 'Instruction : cap 230, maintenez 5 000 ft, QNH 1008, AER452. Quel collationnement est complet et fidèle ?', 'Instruction: heading 230, maintain 5,000 ft, QNH 1008, AER452. Which readback is complete and accurate?', '["Cap 230, QNH 1008, AER452", "Cap 230, maintenons 5 000 ft, QNH 1008, AER452", "Cap 230, maintenons 5 000 ft, AER452", "Cap 230, maintenons 5 000 ft, QNH 1018, AER452"]', '["Heading 230, QNH 1008, AER452", "Heading 230, maintaining 5,000 ft, QNH 1008, AER452", "Heading 230, maintaining 5,000 ft, AER452", "Heading 230, maintaining 5,000 ft, QNH 1018, AER452"]', 'Cap 230, maintenons 5 000 ft, QNH 1008, AER452', 'Comparer présence et valeur de chaque élément. Deux réponses omettent calage ou niveau ; une autre remplace 1008 par 1018. La bonne identification seule ne rend pas ces réponses complètes.

![ATPL_090_001](/illustrations/atpl-090-001.svg)', 'Check both presence and value of every item. Two replies omit setting or level; another substitutes 1018 for 1008. Correct identification alone does not make them complete.

![ATPL_090_001](/illustrations/en/atpl-090-001.svg)', '["ATPL", "exam-mcq", "exam-subject:090"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('3e85fd1e-0000-4000-8000-00003e85fd1e', '1db230b2-0000-4000-8000-00001db230b2', 'ATPL', 'communications', 'communications', 'Collationner les éléments de sécurité', 3, 3, 'Code assigné 4621. Le pilote sélectionne 4621 mais annonce 4612. Quel écart reste à traiter ?', 'Assigned squawk 4621. The pilot selects 4621 but reads back 4612. What discrepancy remains?', '["Le code collationné doit être corrigé à 4621", "Aucun puisque le code sélectionné est correct", "Changer la sélection à 4612 pour l''accorder au message", "Changer de canal pour vérifier le code"]', '["The read-back code must be corrected to 4621", "None because the selected code is correct", "Select 4612 to match the reply", "Change channel to check the code"]', 'Le code collationné doit être corrigé à 4621', 'L''action correcte dans le cockpit ne rend pas fidèle le collationnement entendu par l''ATC. Corriger la parole selon l''instruction, plutôt que modifier une sélection juste pour suivre l''erreur verbale.

![ATPL_090_001](/illustrations/atpl-090-001.svg)', 'Correct cockpit selection does not make the readback heard by ATC accurate. Correct the spoken reply to the instruction rather than changing a correct selection to match the verbal error.

![ATPL_090_001](/illustrations/en/atpl-090-001.svg)', '["ATPL", "exam-mcq", "exam-subject:090"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('3e85fd1f-0000-4000-8000-00003e85fd1f', '1db230b2-0000-4000-8000-00001db230b2', 'ATPL', 'communications', 'communications', 'Collationner les éléments de sécurité', 3, 3, 'Nouveau canal 128.305, collationnement 128.350. Que faut-il résoudre avant transfert ?', 'New channel 128.305, readback 128.350. What must be resolved before transfer?', '["Seulement le nom de la station", "Aucun écart puisque les trois premiers chiffres coïncident", "Le canal précédent doit remplacer les deux valeurs", "La transposition 305/350"]', '["Only the station name", "No discrepancy because the first three digits match", "The previous channel must replace both values", "The 305/350 transposition"]', 'La transposition 305/350', 'Comparer l''intégralité du canal : une terminaison différente mène à une autre sélection. Le nom de station ou les chiffres initiaux ne valident pas la valeur reçue.

![ATPL_090_001](/illustrations/atpl-090-001.svg)', 'Compare the complete channel: a different ending changes selection. Station name or matching initial digits cannot validate the received value.

![ATPL_090_001](/illustrations/en/atpl-090-001.svg)', '["ATPL", "exam-mcq", "exam-subject:090"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('3e85fd20-0000-4000-8000-00003e85fd20', '1db230b2-0000-4000-8000-00001db230b2', 'ATPL', 'communications', 'communications', 'Collationner les éléments de sécurité', 3, 3, 'Instruction : roulez via A, maintenez avant piste 18. Réponse : via A, traversons 18. Quelle erreur modifie la portée ?', 'Instruction: taxi via A, hold short runway 18. Reply: via A, crossing 18. Which error changes scope?', '["La piste 18 est devenue piste 36", "Le seul défaut est l''absence de vitesse", "La traversée remplace le maintien avant piste", "La voie A est devenue inconnue"]', '["Runway 18 became runway 36", "The only defect is missing speed", "Crossing replaces holding short", "Taxiway A has become unknown"]', 'La traversée remplace le maintien avant piste', 'La route de roulage et le numéro sont conservés, mais le verbe autorise à tort une autre action. La correction doit rétablir maintien avant piste, pas simplement répéter son numéro.

![ATPL_090_001](/illustrations/atpl-090-001.svg)', 'Taxi route and runway number match, but the verb wrongly substitutes another action. Correction must restore hold short, not merely repeat the runway number.

![ATPL_090_001](/illustrations/en/atpl-090-001.svg)', '["ATPL", "exam-mcq", "exam-subject:090"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('3e8600e2-0000-4000-8000-00003e8600e2', '1db230b2-0000-4000-8000-00001db230b2', 'ATPL', 'communications', 'communications', 'Fermer la boucle', 3, 3, 'AER452 : cap 240, 5 000 ft, 180 kt. Puis « cap 260, correction, cap 250 ». Aucun autre changement. État autorisé ?', 'AER452: heading 240, 5,000 ft, 180 kt. Then “heading 260, correction, heading 250”. No other change. Cleared state?', '["240 / 5 000 ft / 180 kt", "250 / 5 000 ft / 180 kt", "260 / 5 000 ft / 180 kt", "250 / 2 500 ft / 180 kt"]', '["240 / 5,000 ft / 180 kt", "250 / 5,000 ft / 180 kt", "260 / 5,000 ft / 180 kt", "250 / 2,500 ft / 180 kt"]', '250 / 5 000 ft / 180 kt', 'La correction remplace 260 par 250 ; elle ne transforme pas les chiffres du cap en niveau. 240 est l''ancien cap, désormais modifié. Niveau et vitesse restent inchangés.

![ATPL_090_001](/illustrations/atpl-090-001.svg)', 'Correction replaces 260 with 250; heading digits do not become a level. 240 is the previous, amended heading. Level and speed remain unchanged.

![ATPL_090_001](/illustrations/en/atpl-090-001.svg)', '["ATPL", "exam-mcq", "exam-subject:090"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('3e8600e3-0000-4000-8000-00003e8600e3', '1db230b2-0000-4000-8000-00001db230b2', 'ATPL', 'communications', 'communications', 'Fermer la boucle', 3, 3, 'Instruction : cap 210, QNH 1013, code 5274. Réponse : cap 210, QNH 1031, AER452. Quels défauts détecter ?', 'Instruction: heading 210, QNH 1013, squawk 5274. Reply: heading 210, QNH 1031, AER452. Which defects are present?', '["QNH substitué et code omis", "Cap erroné et QNH correct", "Code erroné mais présent", "QNH seul, puisque le code peut être omis"]', '["Substituted QNH and omitted code", "Wrong heading and correct QNH", "Wrong but present code", "Only QNH, because code can be omitted"]', 'QNH substitué et code omis', '1013 est devenu 1031 et 5274 n''a pas été répété. Détecter une erreur ne doit pas arrêter la recherche d''une seconde ; omission et substitution demandent toutes deux une correction.

![ATPL_090_001](/illustrations/atpl-090-001.svg)', '1013 became 1031 and 5274 was not repeated. Finding one error must not stop the search for another; omission and substitution both need repair.

![ATPL_090_001](/illustrations/en/atpl-090-001.svg)', '["ATPL", "exam-mcq", "exam-subject:090"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('3e8600e4-0000-4000-8000-00003e8600e4', '1db230b2-0000-4000-8000-00001db230b2', 'ATPL', 'communications', 'communications', 'Fermer la boucle', 3, 3, 'Exercice : vitesse minimale de configuration 160 kt, ATC demande 150 kt. Quel message expose correctement la contrainte ?', 'Exercise: configuration minimum speed 160 kt, ATC requests 150 kt. Which message correctly exposes the constraint?', '["WILCO 150, puis maintenir 160 sans annonce", "ROGER 150, puis attendre une détection ATC", "Collationner 160 comme si ce chiffre avait été assigné", "UNABLE 150, minimum 160 dans cette configuration, demandons 160"]', '["WILCO 150, then maintain 160 silently", "ROGER 150, then wait for ATC detection", "Read back 160 as if it had been assigned", "UNABLE 150, minimum 160 in this configuration, request 160"]', 'UNABLE 150, minimum 160 dans cette configuration, demandons 160', 'Annoncer impossibilité, raison et proposition distingue une demande d''une clairance. Les autres réponses masquent la non-exécution ou transforment un chiffre proposé en instruction reçue.

![ATPL_090_001](/illustrations/atpl-090-001.svg)', 'Stating inability, reason and proposal distinguishes a request from clearance. Other replies conceal non-compliance or turn a proposed number into a received instruction.

![ATPL_090_001](/illustrations/en/atpl-090-001.svg)', '["ATPL", "exam-mcq", "exam-subject:090"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('123a1ad4-0000-4000-8000-0000123a1ad4', '1db230b2-0000-4000-8000-00001db230b2', 'ATPL', 'communications', 'communications', 'Fermer la boucle', 3, 3, 'Votre collationnement était 4 000 ft ; l''ATC répond « négatif, répétez [valeur masquée] ». Quelle information pouvez-vous considérer comme résolue ?', 'Your readback was 4,000 ft; ATC replies “negative, say again [blocked value]”. What can be considered resolved?', '["Le niveau prévu sur la fiche de briefing", "La différence exacte de 1 000 ft", "La présence d''un écart, mais pas le niveau corrigé", "Le niveau 4 000 ft, faute d''autre valeur entendue"]', '["The briefing-sheet level", "An exact 1,000 ft difference", "There is a discrepancy, but the corrected level is unresolved", "4,000 ft because no other value was heard"]', 'La présence d''un écart, mais pas le niveau corrigé', 'Le rejet du collationnement est reçu, sa valeur de remplacement ne l''est pas. Redemander la correction ; ni la mémoire ni une différence supposée ne comblent le message manquant.

![ATPL_090_001](/illustrations/atpl-090-001.svg)', 'Rejection of the readback was heard, its replacement value was not. Request the correction again; memory and an assumed difference cannot fill the missing message.

![ATPL_090_001](/illustrations/en/atpl-090-001.svg)', '["ATPL", "exam-mcq", "exam-subject:090"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('123a8f34-0000-4000-8000-0000123a8f34', '1db230b2-0000-4000-8000-00001db230b2', 'ATPL', 'communications', 'communications', 'Distinguer les autorisations', 3, 3, 'Au point d''attente, autorisation « alignez-vous piste 32 et attendez ». Quelle limite doit rester dans la réponse ?', 'At the holding point, clearance “line up runway 32 and wait”. Which limit must remain in the reply?', '["Traverser 32 vers la voie opposée", "Attendre après alignement ; pas de décollage autorisé", "Maintenir avant la piste malgré l''alignement demandé", "Décoller après le précédent trafic sans autre instruction"]', '["Cross 32 to the opposite taxiway", "Wait after line-up; no take-off clearance", "Hold short despite the line-up instruction", "Take off after preceding traffic without further instruction"]', 'Attendre après alignement ; pas de décollage autorisé', 'Alignement, traversée et décollage sont différents. La limite « attendez » conserve le point d''arrêt de cette autorisation ; elle ne définit aucune condition implicite de départ.

![ATPL_090_001](/illustrations/atpl-090-001.svg)', 'Line-up, crossing and take-off differ. “Wait” preserves the stopping point of this clearance; it creates no implicit departure condition.

![ATPL_090_001](/illustrations/en/atpl-090-001.svg)', '["ATPL", "exam-mcq", "exam-subject:090"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('123a8f35-0000-4000-8000-0000123a8f35', '1db230b2-0000-4000-8000-00001db230b2', 'ATPL', 'communications', 'communications', 'Distinguer les autorisations', 3, 3, 'Clairance d''approche reçue ; plus tard « continuez l''approche ». Aucune autorisation d''atterrir. Quel état inscrire ?', 'Approach clearance received; later “continue approach”. No landing clearance. Which state applies?', '["Approche autorisée, atterrissage encore non autorisé", "Atterrissage autorisé par la deuxième transmission", "Approche annulée en l''absence de clairance d''atterrir", "Atterrissage autorisé si la piste est visuellement libre"]', '["Approach cleared, landing not yet cleared", "Landing cleared by the second transmission", "Approach cancelled without landing clearance", "Landing cleared if runway visually clear"]', 'Approche autorisée, atterrissage encore non autorisé', '« Continuez » ne contient pas l''autorisation distincte d''atterrir. L''exercice demande la portée des messages, pas une décision de poursuite à une hauteur donnée ; cette dernière dépend de la situation réelle.

![ATPL_090_001](/illustrations/atpl-090-001.svg)', '“Continue” does not contain separate landing clearance. This asks message scope, not whether to continue at a particular height; that decision depends on the actual situation.

![ATPL_090_001](/illustrations/en/atpl-090-001.svg)', '["ATPL", "exam-mcq", "exam-subject:090"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('1239a678-0000-4000-8000-00001239a678', '1db230b2-0000-4000-8000-00001db230b2', 'ATPL', 'communications', 'communications', 'Collationner les éléments de sécurité', 3, 3, 'Clairance conditionnelle correctement émise : derrière le trafic identifié en courte finale, alignez-vous 24 derrière. Quel collationnement préserve la condition ?', 'Properly issued conditional clearance: behind identified traffic on short final, line up 24 behind. Which readback preserves the condition?', '["Alignons 24, AER452", "Après notre contrôle visuel, décollons 24, AER452", "Maintenons avant 24 puis traversons derrière, AER452", "Derrière ce trafic, alignons 24 derrière, AER452"]', '["Lining up 24, AER452", "After our visual check, taking off 24, AER452", "Holding short 24 then crossing behind, AER452", "Behind that traffic, lining up 24 behind, AER452"]', 'Derrière ce trafic, alignons 24 derrière, AER452', 'La condition est une partie de la clairance, pas un commentaire. L''omettre transforme le sens ; décollage et traversée remplacent l''action autorisée. Le trafic doit être identifié sans ambiguïté.

![ATPL_090_001](/illustrations/atpl-090-001.svg)', 'The condition is part of clearance, not commentary. Omitting it changes meaning; take-off and crossing replace the cleared action. Traffic must be unambiguously identified.

![ATPL_090_001](/illustrations/en/atpl-090-001.svg)', '["ATPL", "exam-mcq", "exam-subject:090"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('123a8f37-0000-4000-8000-0000123a8f37', '1db230b2-0000-4000-8000-00001db230b2', 'ATPL', 'communications', 'communications', 'Distinguer les autorisations', 3, 3, 'Dernier état AER452 : cap 220, 4 000 ft, 170 kt. Message suivant pour AER425 : 160 kt. Puis AER452 : cap 230, maintenez 3 000 ft. Quel état final ?

![ATPL_090_001](/illustrations/atpl-090-001.svg)', 'Last AER452 state: heading 220, 4,000 ft, 170 kt. Next, AER425 receives 160 kt. Then AER452 receives heading 230, maintain 3,000 ft. Final state?

![ATPL_090_001](/illustrations/en/atpl-090-001.svg)', '["220 / 3 000 ft / 170 kt", "230 / 4 000 ft / 170 kt", "230 / 3 000 ft / 170 kt", "230 / 3 000 ft / 160 kt"]', '["220 / 3,000 ft / 170 kt", "230 / 4,000 ft / 170 kt", "230 / 3,000 ft / 170 kt", "230 / 3,000 ft / 160 kt"]', '230 / 3 000 ft / 170 kt', 'Mettre à jour les deux éléments adressés à AER452 : cap et niveau. La vitesse 160 appartient à AER425 ; conserver 170. Les deux autres distracteurs oublient une modification correctement reçue.

![ATPL_090_001](/illustrations/atpl-090-001.svg)', 'Update AER452''s two amended items: heading and level. 160 belongs to AER425; retain 170. The other two distractors each miss a correctly received amendment.

![ATPL_090_001](/illustrations/en/atpl-090-001.svg)', '["ATPL", "exam-mcq", "exam-subject:090"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO lessons (id, license, category, theme, title_fr, title_en, content_fr, content_en, difficulty, order_index, level, duration_minutes, tags, learning_objectives) VALUES
('139f637b-0000-4000-8000-0000139f637b', 'ATPL', 'flight_planning', 'flight_planning', 'Point équitemps et point de retour', 'Equal-time point and safe return', '**Objectif :** Résoudre et vérifier les exercices de point équitemps et point de retour.

## Définir les hypothèses

![ATPL_033_001](/illustrations/atpl-033-001.svg)

On étudie une route rectiligne A–B, avec vitesses sol constantes, tous moteurs en fonctionnement et sans temps de virage, montée ou descente. Le point équitemps (ETP) est celui où revenir à A et poursuivre vers B demandent le même temps. Il ne garantit ni carburant suffisant ni accessibilité météo des terrains. Noter D la distance A–B, H la vitesse sol vers A et O la vitesse sol vers B. Toutes les vitesses doivent correspondre au scénario étudié.

## Calculer et contrôler le point équitemps

À la distance x de A : temps retour = x/H ; temps poursuite = (D−x)/O. Leur égalité donne x = D×H/(H+O). Exemple : D = 600 NM, H = 180 kt, O = 120 kt ; x = 360 NM de A. Contrôle : 360/180 = 2 h et 240/120 = 2 h. On mesure x depuis A, et non depuis B. L’ETP n’est au milieu que si H = O. Une hausse de H, les autres données constantes, éloigne l’ETP de A.

## Calculer le point de retour

Le point de retour dépend de l’autonomie utilisable pour l’aller-retour, après exclusion des réserves et autres allocations imposées au scénario. Avec débit constant identique dans les deux sens, notons E cette autonomie en heures et G la vitesse sol à l’aller. À une distance x : x/G + x/H = E, donc x = E×G×H/(G+H). Exemple fictif : E = 4 h, G = 120 kt, H = 180 kt ; x = 288 NM. Temps aller = 2,4 h, retour = 1,6 h. Ce modèle simplifié ne remplace pas un calcul opérationnel avec les consommations réelles.

## Décider et actualiser

Dans l’exemple, le point de retour à 288 NM est atteint avant l’ETP à 360 NM. Les deux points répondent à des questions différentes. Des débits différents imposent de raisonner en carburant par branche ; une panne impose un nouveau scénario de vitesses, altitude et consommation. Réévaluer avec vent réel, carburant mesuré, météo, NOTAM et disponibilité des terrains. Pour 900 kg dédiés à l’aller-retour et un débit constant de 225 kg/h, E = 4 h. Les 900 kg sont une allocation d’exercice : ils n’incluent pas une réserve légale universelle.

**Format de travail :** 15 QCM de chapitre, quatre choix et une réponse correcte. La matière 033 comporte 42 questions en 120 minutes dans le tableau EASA cité. Ce chapitre ne couvre qu’une partie de la matière ; il ne constitue pas une épreuve complète. Cadre EASA, même en anglais ou avec les unités US/britanniques.

## Références

- [EASA — programme Part-FCL](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8)
- [EASA — format de la matière, AMC1 ARA.FCL.300(b)](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67)

**Support de révision ATPL(A).** À faire valider par votre instructeur. Pour un vol réel, utiliser le manuel de l’avion et les publications en vigueur. Les exemples chiffrés sont pédagogiques.', '**Objective:** Solve and cross-check exercises on equal-time point and safe return.

## Define the assumptions

![ATPL_033_001](/illustrations/en/atpl-033-001.svg)

Consider a straight A–B route with constant groundspeeds, all engines operating and no turning, climb or descent time. The equal-time point (ETP) is where returning to A and continuing to B take the same time. It does not guarantee sufficient fuel or suitable weather at either airfield. Let D be A–B distance, H groundspeed towards A and O groundspeed towards B. All speeds must match the scenario.

## Calculate and cross-check the ETP

At distance x from A: return time = x/H; continuation time = (D−x)/O. Equality gives x = D×H/(H+O). Example: D = 600 NM, H = 180 kt, O = 120 kt; x = 360 NM from A. Cross-check: 360/180 = 2 h and 240/120 = 2 h. Measure x from A, not B. ETP is halfway only if H = O. Increasing H with other inputs fixed moves the ETP farther from A.

## Calculate the safe-return point

The safe-return point depends on usable out-and-back endurance after excluding reserves and other scenario allocations. With equal constant fuel flow in both directions, let E be that endurance in hours and G outbound groundspeed. At distance x: x/G + x/H = E, so x = E×G×H/(G+H). Fictional example: E = 4 h, G = 120 kt, H = 180 kt; x = 288 NM. Outbound time = 2.4 h; return = 1.6 h. This simplified model does not replace operational calculations using actual consumption.

## Decide and update

In this example, the safe-return point at 288 NM is reached before the ETP at 360 NM. The points answer different questions. Different fuel flows require fuel calculations by leg; a failure requires a new speed, altitude and consumption scenario. Reassess using actual wind, measured fuel, weather, NOTAMs and airfield availability. With 900 kg allocated to out-and-back flight and constant flow of 225 kg/h, E = 4 h. The 900 kg is an exercise allocation, not a universal legal reserve.

**Practice format:** 15 chapter MCQs, four choices and one correct answer. Subject 033 has 42 questions in 120 minutes in the cited EASA table. This chapter covers only part of the subject and is not a full paper. EASA framework, including with English or US/British units.

## References

- [Reference 1](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8)
- [Reference 2](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67)

**ATPL(A) revision material.** To be reviewed by your instructor. For an actual flight, use the aircraft manual and current publications. Numerical examples are for training.', 3, 1, 1, 35, '["ATPL", "mcq-four-options"]', '["Résoudre et vérifier les exercices de point équitemps et point de retour."]')
ON CONFLICT (id) DO UPDATE SET license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, title_fr = EXCLUDED.title_fr, title_en = EXCLUDED.title_en, content_fr = EXCLUDED.content_fr, content_en = EXCLUDED.content_en, difficulty = EXCLUDED.difficulty, order_index = EXCLUDED.order_index, level = EXCLUDED.level, duration_minutes = EXCLUDED.duration_minutes, tags = EXCLUDED.tags, learning_objectives = EXCLUDED.learning_objectives;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('29ed13ba-0000-4000-8000-000029ed13ba', '139f637b-0000-4000-8000-0000139f637b', 'ATPL', 'flight_planning', 'flight_planning', 'Calculer et contrôler le point équitemps', 3, 3, 'A–B = 600 NM ; vitesse sol vers A = 180 kt, vers B = 120 kt. ETP mesuré depuis A ?

![ATPL_033_001](/illustrations/atpl-033-001.svg)', 'A–B = 600 NM; groundspeed towards A = 180 kt, towards B = 120 kt. ETP measured from A?

![ATPL_033_001](/illustrations/en/atpl-033-001.svg)', '["360 NM", "240 NM", "300 NM", "400 NM"]', '["360 NM", "240 NM", "300 NM", "400 NM"]', '360 NM', '600 × 180 / (180 + 120) = 360 NM.

![ATPL_033_001](/illustrations/atpl-033-001.svg)', '600 × 180 / (180 + 120) = 360 NM.

![ATPL_033_001](/illustrations/en/atpl-033-001.svg)', '["ATPL", "exam-mcq", "exam-subject:033"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('29ed13bb-0000-4000-8000-000029ed13bb', '139f637b-0000-4000-8000-0000139f637b', 'ATPL', 'flight_planning', 'flight_planning', 'Calculer et contrôler le point équitemps', 3, 3, 'À 360 NM de A, avec vitesse sol retour 180 kt, temps jusqu’à A ?', 'At 360 NM from A, with return groundspeed 180 kt, time to A?', '["3 h", "1 h 20 min", "30 min", "2 h"]', '["3 h", "1 h 20 min", "30 min", "2 h"]', '2 h', '360 / 180 = 2 h.

![ATPL_033_001](/illustrations/atpl-033-001.svg)', '360 / 180 = 2 h.

![ATPL_033_001](/illustrations/en/atpl-033-001.svg)', '["ATPL", "exam-mcq", "exam-subject:033"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('29ed13bc-0000-4000-8000-000029ed13bc', '139f637b-0000-4000-8000-0000139f637b', 'ATPL', 'flight_planning', 'flight_planning', 'Calculer et contrôler le point équitemps', 3, 3, 'Il reste 240 NM vers B à 120 kt sol. Temps de poursuite ?', '240 NM remains towards B at 120 kt groundspeed. Continuation time?', '["3 h", "4 h", "2 h", "1 h"]', '["3 h", "4 h", "2 h", "1 h"]', '2 h', '240 / 120 = 2 h, ce qui confirme l’égalité des temps.

![ATPL_033_001](/illustrations/atpl-033-001.svg)', '240 / 120 = 2 h, confirming equal times.

![ATPL_033_001](/illustrations/en/atpl-033-001.svg)', '["ATPL", "exam-mcq", "exam-subject:033"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('29ed0ffc-0000-4000-8000-000029ed0ffc', '139f637b-0000-4000-8000-0000139f637b', 'ATPL', 'flight_planning', 'flight_planning', 'Définir les hypothèses', 3, 3, 'Quelle grandeur définit un point équitemps entre A et B ?', 'What defines an equal-time point between A and B?', '["Égalité des altitudes des terrains", "Égalité des temps nécessaires vers A et B", "Égalité obligatoire des distances", "Épuisement du carburant"]', '["Equal airfield elevations", "Equal times needed to reach A and B", "Necessarily equal distances", "Fuel exhaustion"]', 'Égalité des temps nécessaires vers A et B', 'Le critère est temporel, avec les vitesses du scénario.

![ATPL_033_001](/illustrations/atpl-033-001.svg)', 'The criterion is time, using the scenario’s speeds.

![ATPL_033_001](/illustrations/en/atpl-033-001.svg)', '["ATPL", "exam-mcq", "exam-subject:033"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('29ed13be-0000-4000-8000-000029ed13be', '139f637b-0000-4000-8000-0000139f637b', 'ATPL', 'flight_planning', 'flight_planning', 'Calculer et contrôler le point équitemps', 3, 3, 'A–B = 600 NM, même vitesse sol dans les deux sens. ETP depuis A ?', 'A–B = 600 NM, equal groundspeed in both directions. ETP from A?', '["300 NM", "600 NM", "200 NM", "400 NM"]', '["300 NM", "600 NM", "200 NM", "400 NM"]', '300 NM', 'Avec H = O, la formule donne D/2.

![ATPL_033_001](/illustrations/atpl-033-001.svg)', 'With H = O, the formula gives D/2.

![ATPL_033_001](/illustrations/en/atpl-033-001.svg)', '["ATPL", "exam-mcq", "exam-subject:033"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('29ed13bf-0000-4000-8000-000029ed13bf', '139f637b-0000-4000-8000-0000139f637b', 'ATPL', 'flight_planning', 'flight_planning', 'Calculer et contrôler le point équitemps', 3, 3, 'Distance totale et vitesse vers B fixes. La vitesse sol vers A augmente. L’ETP mesuré depuis A…', 'Total distance and speed towards B are fixed. Groundspeed towards A increases. ETP distance from A…', '["Diminue", "Reste forcément identique", "Devient négative", "Augmente"]', '["Decreases", "Must stay unchanged", "Becomes negative", "Increases"]', 'Augmente', 'Un retour plus rapide permet de s’éloigner davantage de A pour conserver l’égalité des temps.

![ATPL_033_001](/illustrations/atpl-033-001.svg)', 'A faster return allows greater distance from A while retaining equal times.

![ATPL_033_001](/illustrations/en/atpl-033-001.svg)', '["ATPL", "exam-mcq", "exam-subject:033"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('29ed0fff-0000-4000-8000-000029ed0fff', '139f637b-0000-4000-8000-0000139f637b', 'ATPL', 'flight_planning', 'flight_planning', 'Définir les hypothèses', 3, 3, 'Atteindre l’ETP garantit-il que les deux terrains sont accessibles avec le carburant restant ?', 'Does reaching the ETP guarantee both airfields can be reached with remaining fuel?', '["Oui, uniquement de jour", "Oui, si les distances sont égales", "Non", "Oui, par définition"]', '["Yes, only by day", "Yes, if distances are equal", "No", "Yes, by definition"]', 'Non', 'L’égalité des temps ne contrôle pas les réserves ni l’accessibilité des terrains.

![ATPL_033_001](/illustrations/atpl-033-001.svg)', 'Equal times do not check reserves or airfield suitability.

![ATPL_033_001](/illustrations/en/atpl-033-001.svg)', '["ATPL", "exam-mcq", "exam-subject:033"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('29ed1782-0000-4000-8000-000029ed1782', '139f637b-0000-4000-8000-0000139f637b', 'ATPL', 'flight_planning', 'flight_planning', 'Calculer le point de retour', 3, 3, 'Modèle rectiligne sans temps de virage, montée ou descente. Autonomie aller-retour après exclusion des réserves : 4 h ; débit constant identique dans les deux sens ; vitesses sol constantes 120 kt à l’aller, 180 kt au retour. Distance du point de retour depuis A ?', 'Straight-route model with no turning, climb or descent time. Out-and-back endurance after excluding reserves: 4 h; equal constant fuel flow both ways; constant groundspeeds 120 kt outbound, 180 kt inbound. Safe-return distance from A?', '["720 NM", "288 NM", "360 NM", "480 NM"]', '["720 NM", "288 NM", "360 NM", "480 NM"]', '288 NM', 'x/120 + x/180 = 4 donne 288 NM. Contrôle : 2,4 h aller + 1,6 h retour = 4 h. 480 et 720 NM consacreraient toute l’autonomie à une seule branche.

![ATPL_033_001](/illustrations/atpl-033-001.svg)', 'x/120 + x/180 = 4 gives 288 NM. Check: 2.4 h outbound + 1.6 h inbound = 4 h. 480 and 720 NM would allocate all endurance to only one leg.

![ATPL_033_001](/illustrations/en/atpl-033-001.svg)', '["ATPL", "exam-mcq", "exam-subject:033"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('29ed1783-0000-4000-8000-000029ed1783', '139f637b-0000-4000-8000-0000139f637b', 'ATPL', 'flight_planning', 'flight_planning', 'Calculer le point de retour', 3, 3, 'Point de retour à 288 NM ; vitesse aller 120 kt. Temps aller ?', 'Safe-return point at 288 NM; outbound groundspeed 120 kt. Outbound time?', '["2 h 24 min", "2 h 40 min", "1 h 36 min", "3 h"]', '["2 h 24 min", "2 h 40 min", "1 h 36 min", "3 h"]', '2 h 24 min', '288 / 120 = 2,4 h ; 0,4 h = 24 min.

![ATPL_033_001](/illustrations/atpl-033-001.svg)', '288 / 120 = 2.4 h; 0.4 h = 24 min.

![ATPL_033_001](/illustrations/en/atpl-033-001.svg)', '["ATPL", "exam-mcq", "exam-subject:033"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('29ed1784-0000-4000-8000-000029ed1784', '139f637b-0000-4000-8000-0000139f637b', 'ATPL', 'flight_planning', 'flight_planning', 'Calculer le point de retour', 3, 3, 'Retour de 288 NM à 180 kt. Temps retour ?', 'Return distance 288 NM at 180 kt. Return time?', '["1 h 60 min", "2 h 24 min", "48 min", "1 h 36 min"]', '["1 h 60 min", "2 h 24 min", "48 min", "1 h 36 min"]', '1 h 36 min', '288 / 180 = 1,6 h ; 0,6 h = 36 min.

![ATPL_033_001](/illustrations/atpl-033-001.svg)', '288 / 180 = 1.6 h; 0.6 h = 36 min.

![ATPL_033_001](/illustrations/en/atpl-033-001.svg)', '["ATPL", "exam-mcq", "exam-subject:033"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('13b5d834-0000-4000-8000-000013b5d834', '139f637b-0000-4000-8000-0000139f637b', 'ATPL', 'flight_planning', 'flight_planning', 'Calculer le point de retour', 3, 3, 'Les 4 h d’autonomie de l’exercice doivent représenter…', 'The exercise’s 4 h endurance must represent…', '["La durée totale de maintenance", "Le temps de vol déjà effectué", "Le temps utilisable pour l’aller-retour après les allocations exclues", "L’endurance jusqu’à réservoirs secs, réserves incluses"]', '["Total maintenance time", "Flight time already elapsed", "Usable out-and-back time after excluded allocations", "Endurance to dry tanks including reserves"]', 'Le temps utilisable pour l’aller-retour après les allocations exclues', 'Définir clairement l’allocation utilisable avant de calculer le point de retour.

![ATPL_033_001](/illustrations/atpl-033-001.svg)', 'Define the usable allocation before calculating the safe-return point.

![ATPL_033_001](/illustrations/en/atpl-033-001.svg)', '["ATPL", "exam-mcq", "exam-subject:033"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('13b64c94-0000-4000-8000-000013b64c94', '139f637b-0000-4000-8000-0000139f637b', 'ATPL', 'flight_planning', 'flight_planning', 'Décider et actualiser', 3, 3, 'ETP à 360 NM, point de retour à 288 NM de A. Lequel est rencontré en premier au départ de A ?', 'ETP at 360 NM, safe-return point at 288 NM from A. Which is reached first departing A?', '["Impossible malgré les distances données", "Le point de retour", "L’ETP", "Ils sont nécessairement confondus"]', '["Impossible despite the stated distances", "The safe-return point", "The ETP", "They necessarily coincide"]', 'Le point de retour', '288 NM est inférieur à 360 NM ; les critères des deux points sont différents.

![ATPL_033_001](/illustrations/atpl-033-001.svg)', '288 NM is less than 360 NM; the points use different criteria.

![ATPL_033_001](/illustrations/en/atpl-033-001.svg)', '["ATPL", "exam-mcq", "exam-subject:033"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('13b64c95-0000-4000-8000-000013b64c95', '139f637b-0000-4000-8000-0000139f637b', 'ATPL', 'flight_planning', 'flight_planning', 'Décider et actualiser', 3, 3, '900 kg sont alloués à l’aller-retour, débit constant 225 kg/h. Autonomie utilisable ?', '900 kg is allocated to out-and-back flight, constant flow 225 kg/h. Usable endurance?', '["4 h", "3 h", "5 h", "6 h"]', '["4 h", "3 h", "5 h", "6 h"]', '4 h', '900 / 225 = 4 h.

![ATPL_033_001](/illustrations/atpl-033-001.svg)', '900 / 225 = 4 h.

![ATPL_033_001](/illustrations/en/atpl-033-001.svg)', '["ATPL", "exam-mcq", "exam-subject:033"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('13b64c96-0000-4000-8000-000013b64c96', '139f637b-0000-4000-8000-0000139f637b', 'ATPL', 'flight_planning', 'flight_planning', 'Décider et actualiser', 3, 3, 'À vitesses inchangées, l’autonomie utilisable passe de 4 h à 3 h. Le point de retour précédemment à 288 NM devient…', 'At unchanged speeds, usable endurance falls from 4 h to 3 h. The safe-return point previously at 288 NM becomes…', '["384 NM", "288 NM", "144 NM", "216 NM"]', '["384 NM", "288 NM", "144 NM", "216 NM"]', '216 NM', 'La distance est proportionnelle à E : 288 × 3/4 = 216 NM.

![ATPL_033_001](/illustrations/atpl-033-001.svg)', 'Distance is proportional to E: 288 × 3/4 = 216 NM.

![ATPL_033_001](/illustrations/en/atpl-033-001.svg)', '["ATPL", "exam-mcq", "exam-subject:033"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('13b64c97-0000-4000-8000-000013b64c97', '139f637b-0000-4000-8000-0000139f637b', 'ATPL', 'flight_planning', 'flight_planning', 'Décider et actualiser', 3, 3, 'Une panne change vitesse et consommation. Peut-on conserver sans recalcul le point obtenu tous moteurs en fonctionnement ?', 'A failure changes speed and fuel flow. Can the all-engines calculation be retained without recalculation?', '["Oui, car la route reste droite", "Oui, si l’ETP est déjà franchi", "Non, actualiser le scénario", "Oui, car la distance A–B est inchangée"]', '["Yes, because the route remains straight", "Yes, if already past ETP", "No, update the scenario", "Yes, because A–B distance is unchanged"]', 'Non, actualiser le scénario', 'Les hypothèses ont changé ; réévaluer temps, carburant et terrains utilisables.

![ATPL_033_001](/illustrations/atpl-033-001.svg)', 'Assumptions changed; reassess time, fuel and suitable airfields.

![ATPL_033_001](/illustrations/en/atpl-033-001.svg)', '["ATPL", "exam-mcq", "exam-subject:033"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO lessons (id, license, category, theme, title_fr, title_en, content_fr, content_en, difficulty, order_index, level, duration_minutes, tags, learning_objectives) VALUES
('11dc735b-0000-4000-8000-000011dc735b', 'ATPL', 'instrumentation', 'instrumentation', 'Vitesses air, Mach et données aérodynamiques', 'Airspeeds, Mach and air data', '**Objectif :** Résoudre et vérifier les exercices de vitesses air, mach et données aérodynamiques.

## Séparer les vitesses

![ATPL_022_001](/illustrations/atpl-022-001.svg)

IAS est la vitesse indiquée. CAS corrige les erreurs instrumentales et de position ; EAS corrige ensuite les effets de compressibilité ; TAS est la vitesse réelle par rapport à la masse d’air. Le vent transforme le vecteur vitesse air en vecteur vitesse sol. Dans les exercices à faible Mach, on peut assimiler CAS et EAS seulement lorsque l’énoncé le permet. Une valeur de vitesse n’a de sens qu’avec son type et son unité.

## Densité et vitesse vraie

Avec σ = densité locale / densité standard au niveau de la mer, TAS = EAS/√σ. Exemple : EAS 200 kt et σ = 0,64 donnent TAS = 250 kt. À EAS constante, une densité plus faible correspond à une TAS plus élevée. Ne pas utiliser cette relation pour déduire directement la vitesse sol sans vent. Les calculs supposent des données corrigées et un écoulement approprié au modèle.

## Mach et température

Mach = TAS / vitesse locale du son. La vitesse du son dans l’air dépend de la température absolue : a ∝ √T. Pour l’exercice, si a = 600 kt et TAS = 450 kt, Mach = 0,75. À Mach constant, une baisse de T réduit la TAS. Utiliser des kelvins : 15 °C ≈ 288 K, −45 °C ≈ 228 K. La température totale mesurée en vol rapide inclut un réchauffement dynamique ; elle n’est pas directement la température statique de l’air.

## Détecter les incohérences

Les données aérodynamiques dépendent des pressions totale et statique et des capteurs associés. Une obstruction peut affecter plusieurs instruments ; plusieurs affichages partageant la même source ne constituent pas des vérifications indépendantes. Comparer les sources réellement indépendantes et appliquer les procédures de vitesse non fiable propres à l’avion. VMO et MMO sont des limites distinctes : respecter celle qui devient contraignante, sans supposer qu’une IAS acceptable garantit un Mach acceptable.

**Format de travail :** 15 QCM de chapitre, quatre choix et une réponse correcte. La matière 022 comporte 60 questions en 90 minutes dans le tableau EASA cité. Ce chapitre ne couvre qu’une partie de la matière ; il ne constitue pas une épreuve complète. Cadre EASA, même en anglais ou avec les unités US/britanniques.

## Références

- [EASA — programme Part-FCL](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8)
- [EASA — format de la matière, AMC1 ARA.FCL.300(b)](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67)
- [FAA PHAK — instruments](https://www.faa.gov/regulationspolicies/handbooksmanuals/aviation/phak/chapter-8-flight-instruments)

**Support de révision ATPL(A).** À faire valider par votre instructeur. Pour un vol réel, utiliser le manuel de l’avion et les publications en vigueur. Les exemples chiffrés sont pédagogiques.', '**Objective:** Solve and cross-check exercises on airspeeds, mach and air data.

## Separate the airspeeds

![ATPL_022_001](/illustrations/en/atpl-022-001.svg)

IAS is indicated airspeed. CAS corrects instrument and position errors; EAS then corrects compressibility effects; TAS is true speed relative to the air mass. Wind converts the air-velocity vector into the ground-velocity vector. In low-Mach exercises, CAS and EAS may be treated alike only when allowed by the problem. A speed value requires both its type and unit.

## Density and true airspeed

With σ = local density / standard sea-level density, TAS = EAS/√σ. Example: EAS 200 kt and σ = 0.64 give TAS = 250 kt. At constant EAS, lower density corresponds to higher TAS. This does not directly determine groundspeed without wind information. Calculations assume corrected inputs and flow appropriate to the model.

## Mach and temperature

Mach = TAS / local speed of sound. Sound speed in air depends on absolute temperature: a ∝ √T. In the exercise, a = 600 kt and TAS = 450 kt give Mach = 0.75. At constant Mach, lower T reduces TAS. Use kelvins: 15 °C ≈ 288 K, −45 °C ≈ 228 K. Total temperature measured in fast flight includes ram heating; it is not directly static air temperature.

## Detect inconsistencies

Air data depend on total and static pressure and associated sensors. A blockage can affect several instruments; displays sharing one source are not independent cross-checks. Compare genuinely independent sources and follow aircraft-specific unreliable-airspeed procedures. VMO and MMO are separate limits: respect whichever is limiting, without assuming acceptable IAS guarantees acceptable Mach.

**Practice format:** 15 chapter MCQs, four choices and one correct answer. Subject 022 has 60 questions in 90 minutes in the cited EASA table. This chapter covers only part of the subject and is not a full paper. EASA framework, including with English or US/British units.

## References

- [Reference 1](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8)
- [Reference 2](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67)
- [Reference 3](https://www.faa.gov/regulationspolicies/handbooksmanuals/aviation/phak/chapter-8-flight-instruments)

**ATPL(A) revision material.** To be reviewed by your instructor. For an actual flight, use the aircraft manual and current publications. Numerical examples are for training.', 3, 1, 1, 35, '["ATPL", "mcq-four-options"]', '["Résoudre et vérifier les exercices de vitesses air, mach et données aérodynamiques."]')
ON CONFLICT (id) DO UPDATE SET license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, title_fr = EXCLUDED.title_fr, title_en = EXCLUDED.title_en, content_fr = EXCLUDED.content_fr, content_en = EXCLUDED.content_en, difficulty = EXCLUDED.difficulty, order_index = EXCLUDED.order_index, level = EXCLUDED.level, duration_minutes = EXCLUDED.duration_minutes, tags = EXCLUDED.tags, learning_objectives = EXCLUDED.learning_objectives;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('0d24dfba-0000-4000-8000-00000d24dfba', '11dc735b-0000-4000-8000-000011dc735b', 'ATPL', 'instrumentation', 'instrumentation', 'Densité et vitesse vraie', 3, 3, 'EAS 200 kt, rapport de densité σ = 0,64. TAS ?

![ATPL_022_001](/illustrations/atpl-022-001.svg)', 'EAS 200 kt, density ratio σ = 0.64. TAS?

![ATPL_022_001](/illustrations/en/atpl-022-001.svg)', '["250 kt", "160 kt", "312,5 kt", "200 kt"]', '["250 kt", "160 kt", "312.5 kt", "200 kt"]', '250 kt', '200/√0,64 = 250 kt.

![ATPL_022_001](/illustrations/atpl-022-001.svg)', '200/√0.64 = 250 kt.

![ATPL_022_001](/illustrations/en/atpl-022-001.svg)', '["ATPL", "exam-mcq", "exam-subject:022"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('0d24dbfa-0000-4000-8000-00000d24dbfa', '11dc735b-0000-4000-8000-000011dc735b', 'ATPL', 'instrumentation', 'instrumentation', 'Séparer les vitesses', 3, 3, 'Quelle vitesse corrige l’IAS des erreurs de position et instrumentales ?', 'Which airspeed corrects IAS for position and instrument errors?', '["GS", "Vitesse du son", "Vitesse verticale", "CAS"]', '["GS", "Speed of sound", "Vertical speed", "CAS"]', 'CAS', 'La CAS est la vitesse indiquée ainsi corrigée.

![ATPL_022_001](/illustrations/atpl-022-001.svg)', 'CAS is indicated airspeed corrected for these errors.

![ATPL_022_001](/illustrations/en/atpl-022-001.svg)', '["ATPL", "exam-mcq", "exam-subject:022"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('0d24dbfb-0000-4000-8000-00000d24dbfb', '11dc735b-0000-4000-8000-000011dc735b', 'ATPL', 'instrumentation', 'instrumentation', 'Séparer les vitesses', 3, 3, 'Quelle correction distingue EAS de CAS ?', 'Which correction distinguishes EAS from CAS?', '["Déclinaison magnétique", "Pente de piste", "Compressibilité", "Vent uniquement"]', '["Magnetic variation", "Runway slope", "Compressibility", "Wind only"]', 'Compressibilité', 'EAS tient compte des effets de compressibilité.

![ATPL_022_001](/illustrations/atpl-022-001.svg)', 'EAS accounts for compressibility effects.

![ATPL_022_001](/illustrations/en/atpl-022-001.svg)', '["ATPL", "exam-mcq", "exam-subject:022"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('0d24dbfc-0000-4000-8000-00000d24dbfc', '11dc735b-0000-4000-8000-000011dc735b', 'ATPL', 'instrumentation', 'instrumentation', 'Séparer les vitesses', 3, 3, 'TAS désigne la vitesse relative…', 'TAS is speed relative to…', '["À la piste en permanence", "À la masse d’air", "Au sol uniquement", "Au nord magnétique"]', '["The runway at all times", "The air mass", "The ground only", "Magnetic north"]', 'À la masse d’air', 'La vitesse sol combine ensuite air et vent.

![ATPL_022_001](/illustrations/atpl-022-001.svg)', 'Groundspeed then combines air velocity and wind.

![ATPL_022_001](/illustrations/en/atpl-022-001.svg)', '["ATPL", "exam-mcq", "exam-subject:022"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('0d24dfbe-0000-4000-8000-00000d24dfbe', '11dc735b-0000-4000-8000-000011dc735b', 'ATPL', 'instrumentation', 'instrumentation', 'Densité et vitesse vraie', 3, 3, 'À EAS constante, σ passe de 1 à 0,64. La TAS…', 'At constant EAS, σ changes from 1 to 0.64. TAS…', '["Est multipliée par 1,25", "Est multipliée par 0,64", "Reste identique", "Est divisée par 1,25"]', '["Is multiplied by 1.25", "Is multiplied by 0.64", "Stays unchanged", "Is divided by 1.25"]', 'Est multipliée par 1,25', '1/√0,64 = 1,25.

![ATPL_022_001](/illustrations/atpl-022-001.svg)', '1/√0.64 = 1.25.

![ATPL_022_001](/illustrations/en/atpl-022-001.svg)', '["ATPL", "exam-mcq", "exam-subject:022"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('0d24dfbf-0000-4000-8000-00000d24dfbf', '11dc735b-0000-4000-8000-000011dc735b', 'ATPL', 'instrumentation', 'instrumentation', 'Densité et vitesse vraie', 3, 3, 'TAS 250 kt, vent de face aligné 30 kt. Vitesse sol ?', 'TAS 250 kt, aligned headwind 30 kt. Groundspeed?', '["280 kt", "250 kt", "30 kt", "220 kt"]', '["280 kt", "250 kt", "30 kt", "220 kt"]', '220 kt', '250 − 30 = 220 kt.

![ATPL_022_001](/illustrations/atpl-022-001.svg)', '250 − 30 = 220 kt.

![ATPL_022_001](/illustrations/en/atpl-022-001.svg)', '["ATPL", "exam-mcq", "exam-subject:022"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('0d24e381-0000-4000-8000-00000d24e381', '11dc735b-0000-4000-8000-000011dc735b', 'ATPL', 'instrumentation', 'instrumentation', 'Mach et température', 3, 3, 'TAS 450 kt et vitesse locale du son 600 kt. Mach ?', 'TAS 450 kt and local sound speed 600 kt. Mach?', '["0,60", "0,45", "0,75", "1,33"]', '["0.60", "0.45", "0.75", "1.33"]', '0,75', '450/600 = 0,75.

![ATPL_022_001](/illustrations/atpl-022-001.svg)', '450/600 = 0.75.

![ATPL_022_001](/illustrations/en/atpl-022-001.svg)', '["ATPL", "exam-mcq", "exam-subject:022"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('0d24e382-0000-4000-8000-00000d24e382', '11dc735b-0000-4000-8000-000011dc735b', 'ATPL', 'instrumentation', 'instrumentation', 'Mach et température', 3, 3, 'Mach 0,80 et vitesse locale du son 575 kt. TAS ?', 'Mach 0.80 and local sound speed 575 kt. TAS?', '["495 kt", "460 kt", "719 kt", "575 kt"]', '["495 kt", "460 kt", "719 kt", "575 kt"]', '460 kt', '0,80 × 575 = 460 kt.

![ATPL_022_001](/illustrations/atpl-022-001.svg)', '0.80 × 575 = 460 kt.

![ATPL_022_001](/illustrations/en/atpl-022-001.svg)', '["ATPL", "exam-mcq", "exam-subject:022"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('0d24e383-0000-4000-8000-00000d24e383', '11dc735b-0000-4000-8000-000011dc735b', 'ATPL', 'instrumentation', 'instrumentation', 'Mach et température', 3, 3, 'À Mach constant, une baisse de température statique absolue fait…', 'At constant Mach, falling absolute static temperature makes…', '["Diminuer la TAS", "Augmenter toujours la TAS", "Rester la TAS constante", "Annuler le vent"]', '["TAS decrease", "TAS always increase", "TAS stay constant", "Wind vanish"]', 'Diminuer la TAS', 'La vitesse du son varie avec √T.

![ATPL_022_001](/illustrations/atpl-022-001.svg)', 'Sound speed varies with √T.

![ATPL_022_001](/illustrations/en/atpl-022-001.svg)', '["ATPL", "exam-mcq", "exam-subject:022"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('0d24e384-0000-4000-8000-00000d24e384', '11dc735b-0000-4000-8000-000011dc735b', 'ATPL', 'instrumentation', 'instrumentation', 'Mach et température', 3, 3, 'Dans a ∝ √T, quelle échelle utiliser ?', 'In a ∝ √T, which scale is required?', '["Celsius directement", "Fahrenheit directement", "QNH en hPa", "Kelvin"]', '["Celsius directly", "Fahrenheit directly", "QNH in hPa", "Kelvin"]', 'Kelvin', 'Il faut une température absolue.

![ATPL_022_001](/illustrations/atpl-022-001.svg)', 'An absolute temperature is required.

![ATPL_022_001](/illustrations/en/atpl-022-001.svg)', '["ATPL", "exam-mcq", "exam-subject:022"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('17778c34-0000-4000-8000-000017778c34', '11dc735b-0000-4000-8000-000011dc735b', 'ATPL', 'instrumentation', 'instrumentation', 'Mach et température', 3, 3, '−45 °C correspond approximativement à…', '−45 °C corresponds approximately to…', '["45 K", "−45 K", "228 K", "318 K"]', '["45 K", "−45 K", "228 K", "318 K"]', '228 K', '−45 + 273 ≈ 228 K.

![ATPL_022_001](/illustrations/atpl-022-001.svg)', '−45 + 273 ≈ 228 K.

![ATPL_022_001](/illustrations/en/atpl-022-001.svg)', '["ATPL", "exam-mcq", "exam-subject:022"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('17778c35-0000-4000-8000-000017778c35', '11dc735b-0000-4000-8000-000011dc735b', 'ATPL', 'instrumentation', 'instrumentation', 'Mach et température', 3, 3, 'En vol rapide, la température totale mesurée comprend…', 'In fast flight, measured total temperature includes…', '["La déclinaison locale", "Un réchauffement dynamique", "Seulement la température du carburant", "Uniquement l’altitude-pression"]', '["Local variation", "Ram heating", "Only fuel temperature", "Only pressure altitude"]', 'Un réchauffement dynamique', 'Elle diffère de la température statique.

![ATPL_022_001](/illustrations/atpl-022-001.svg)', 'It differs from static temperature.

![ATPL_022_001](/illustrations/en/atpl-022-001.svg)', '["ATPL", "exam-mcq", "exam-subject:022"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('17780095-0000-4000-8000-000017780095', '11dc735b-0000-4000-8000-000011dc735b', 'ATPL', 'instrumentation', 'instrumentation', 'Détecter les incohérences', 3, 3, 'Deux écrans alimentés par la même source aérodynamique constituent-ils deux références indépendantes ?', 'Are two displays fed by one air-data source independent references?', '["Non", "Oui, par leur nombre", "Oui, si leurs couleurs diffèrent", "Oui, avec deux pilotes"]', '["No", "Yes, because there are two", "Yes, if colours differ", "Yes, with two pilots"]', 'Non', 'Une panne commune peut affecter les deux.

![ATPL_022_001](/illustrations/atpl-022-001.svg)', 'A common failure can affect both.

![ATPL_022_001](/illustrations/en/atpl-022-001.svg)', '["ATPL", "exam-mcq", "exam-subject:022"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('17780096-0000-4000-8000-000017780096', '11dc735b-0000-4000-8000-000011dc735b', 'ATPL', 'instrumentation', 'instrumentation', 'Détecter les incohérences', 3, 3, 'IAS sous VMO mais Mach au-dessus de MMO. La situation respecte-t-elle les limites ?', 'IAS below VMO but Mach above MMO. Are limits respected?', '["Oui, VMO suffit", "Oui, en palier", "Oui, si la vitesse sol est faible", "Non"]', '["Yes, VMO is sufficient", "Yes, in level flight", "Yes, with low groundspeed", "No"]', 'Non', 'Les deux limites doivent être respectées.

![ATPL_022_001](/illustrations/atpl-022-001.svg)', 'Both limits must be respected.

![ATPL_022_001](/illustrations/en/atpl-022-001.svg)', '["ATPL", "exam-mcq", "exam-subject:022"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('17780097-0000-4000-8000-000017780097', '11dc735b-0000-4000-8000-000011dc735b', 'ATPL', 'instrumentation', 'instrumentation', 'Détecter les incohérences', 3, 3, 'Face à des vitesses incohérentes, quelle référence de procédure utiliser ?', 'With inconsistent airspeeds, which procedural reference should be used?', '["La seule vitesse GPS comme IAS", "L’avis d’un passager", "La procédure de vitesse non fiable propre à l’avion", "Une valeur internet générique"]', '["GPS speed alone as IAS", "A passenger’s opinion", "Aircraft-specific unreliable-airspeed procedure", "A generic internet value"]', 'La procédure de vitesse non fiable propre à l’avion', 'Les sources et actions dépendent du système installé.

![ATPL_022_001](/illustrations/atpl-022-001.svg)', 'Sources and actions depend on the installed system.

![ATPL_022_001](/illustrations/en/atpl-022-001.svg)', '["ATPL", "exam-mcq", "exam-subject:022"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO lessons (id, license, category, theme, title_fr, title_en, content_fr, content_en, difficulty, order_index, level, duration_minutes, tags, learning_objectives) VALUES
('13833479-0000-4000-8000-000013833479', 'ATPL', 'mass_and_balance', 'mass_and_balance', 'Masses opérationnelles et évolution du centrage', 'Operating masses and CG evolution', '**Objectif :** Établir un bilan sur données nouvelles, identifier la contrainte limitante, lire une enveloppe interpolée et détecter un dépassement intermédiaire de centrage.

## Distinguer les étapes du bilan

![ATPL_031_001](/illustrations/atpl-031-001.svg)

Exercice fictif : masse opérationnelle sèche DOW = 30 000 kg, charge marchande = 8 000 kg. La masse sans carburant utilisable ZFM vaut 38 000 kg. Ajouter 10 000 kg de carburant bloc donne une masse au parking de 48 000 kg. Après 200 kg de roulage, la masse au décollage vaut 47 800 kg. Avec 6 000 kg de carburant trajet consommés après décollage, la masse prévue à l’atterrissage vaut 41 800 kg. Ne pas soustraire le roulage deux fois.

### Identifier le contenu de chaque masse

La DOW représente ici l’avion prêt pour le type d’exploitation considéré, hors carburant utilisable et charge marchande. Ce n’est pas simplement la masse de la structure vide. Vérifier ce que le document d’exploitation inclut déjà, pour ne pas ajouter deux fois équipage ou équipements.

Suivre une chronologie écrite : DOW + charge marchande = ZFM ; ZFM + carburant bloc = masse au parking ; parking − roulage = décollage ; décollage − trajet = atterrissage prévu. Dans cet exercice, aucune autre masse n’est ajoutée ou retirée entre ces étapes, et les consommations données sont des masses de carburant.

**Contrôle par le carburant restant :** 10 000 − 200 − 6 000 = 3 800 kg. ZFM 38 000 + restant 3 800 = **41 800 kg**, identique au calcul par soustractions successives. Déduire encore 200 kg au stade atterrissage compterait le roulage deux fois.

Les 3 800 kg restants ne démontrent pas à eux seuls le respect des réserves réglementaires : ce chapitre vérifie un bilan de masses. Les allocations carburant et les hypothèses opérationnelles demandent leur propre contrôle.

## Vérifier chaque limitation

Limites fictives : maximum ZFM 39 000 kg, décollage 49 000 kg, atterrissage 42 000 kg. Le bilan initial respecte ces trois plafonds de masse, sans prouver la conformité des performances ou du centrage. À carburant inchangé, ajouter une charge ne peut dépasser la plus petite marge restante : 1 000, 1 200 et 200 kg respectivement. La limite d’atterrissage restreint donc l’ajout à 200 kg dans cet exercice. Des limites opérationnelles peuvent être plus restrictives que les maxima structuraux.

### Trouver la contrainte qui commande la décision

Écrire les trois inégalités plutôt que d’additionner les marges. Une charge supplémentaire p, à carburant bloc, roulage et trajet inchangés, impose p ≤ 1 000 kg, p ≤ 1 200 kg et p ≤ 200 kg. Il faut satisfaire les trois simultanément, donc **p ≤ 200 kg**. Une moyenne de ces marges n’a pas de sens opérationnel.

**Sensibilité au carburant :** si l’on ajoute séparément 400 kg de carburant bloc, avec mêmes consommations, la ZFM reste 38 000 kg, le décollage passe à 48 200 kg et l’atterrissage à 42 200 kg. Le décollage respecte 49 000 kg, mais l’arrivée dépasse 42 000 kg de 200 kg. Ces seules données ne permettent pas de choisir quelle allocation carburant réduire : il faut conserver les exigences de carburant et recalculer le scénario.

**Cas autonome :** DOW 30 000 kg, charge 7 500 kg, bloc 10 000 kg, roulage 200 kg, trajet 6 000 kg, mêmes plafonds. ZFM 37 500, décollage 47 300, arrivée 41 300 kg ; marges 1 500, 1 700 et 700 kg. Le plafond limitant pour un ajout de charge est donc **700 kg**, sous les hypothèses fixées et les autres contraintes.

### Ajouter une limite de performance au bilan

Reprendre le cas autonome à 7 500 kg de charge, dont la masse au décollage est 47 300 kg. Si une limite de performance fictive impose 47 600 kg, la marge supplémentaire vaut **300 kg**. Elle est plus faible que les 700 kg permis par les trois seuls plafonds structuraux. Ajouter la nouvelle contrainte à la liste et reprendre le minimum ; ne pas remplacer arbitrairement tous les autres contrôles par celui du décollage.

Si l’exercice impose des consommations inchangées, les utiliser comme telles. Une exploitation réelle doit réévaluer les consommations et performances lorsque le chargement change : la constance du carburant trajet est ici une hypothèse de calcul, pas une règle générale.

## Recalculer le centrage après consommation

Exercice de centrage indépendant : masse 10 000 kg, centrage à 5,00 m du datum ; moment total 50 000 kg·m. Consommer 1 000 kg de carburant au bras constant 6,00 m retire 6 000 kg·m. Le nouveau centrage est 44 000 / 9 000 = 4,889 m : il avance, car la masse retirée était derrière le centrage initial. Ne pas déduire une direction générale du seul mot « consommation » : elle dépend du bras et de la séquence des réservoirs.

### Prévoir le sens avant de calculer

Dans cet exercice indépendant, les bras augmentent vers l’arrière. Au départ, 10 000 × 5 = 50 000 kg·m. Retirer 1 000 kg à 6 m laisse 44 000 kg·m et 9 000 kg. Le nouveau CG 4,8889… m est inférieur à 5 m, ce qui confirme le déplacement vers l’avant attendu.

**Contre-exemple :** repartir du même état initial mais retirer les 1 000 kg au bras 4 m. Le nouveau CG est (50 000 − 4 000)/9 000 = **5,1111 m** : il recule. La consommation ne fait donc pas toujours avancer le CG. Il faut connaître le réservoir, son bras et la séquence de consommation.

Ne pas réutiliser le bilan de 47 800 kg pour ce cas de 10 000 kg : les deux exercices sont indépendants. Dans un calcul réel, suivre la méthode de réservoir applicable, notamment si le bras varie avec la quantité. Vérifier les états intermédiaires pertinents ; des extrémités conformes ne prouvent pas automatiquement que tout le chemin de consommation l’est.

### Combiner chargement et consommation

**Exemple distinct :** état initial 12 000 kg au CG 5 m, soit 60 000 kg·m. Déplacer 300 kg du bras 7 m au bras 4 m enlève 900 kg·m sans changer la masse. Consommer ensuite 1 000 kg au bras constant 6 m enlève 1 000 kg et 6 000 kg·m. Totaux finaux : 11 000 kg et 53 100 kg·m ; CG = **4,827 m** arrondi.

Une ligne par modification évite de soustraire la charge déplacée de la masse totale ou d’oublier son ancien moment. Pour plusieurs réservoirs, sommer séparément les masses et les moments consommés ; le bras moyen du carburant retiré peut être en avant, en arrière ou exactement au CG initial. Le résultat final n’exonère pas du contrôle des états intermédiaires pertinents.

## Exprimer le centrage en pourcentage de corde

La position en pourcentage de corde aérodynamique moyenne vaut 100 × (bras CG − bras du bord d’attaque de la MAC) / longueur MAC. Exemple : CG 5,00 m, bord d’attaque 4,00 m, MAC 4,00 m : 25 %. Un centrage à 4,80 m donne 20 %. Comparer masse et centrage à l’enveloppe applicable à chaque phase, pas uniquement au départ. La concordance d’un pourcentage ne dispense pas de vérifier la référence, la configuration et les limites propres à l’avion.

### Vérifier une conversion dans les deux sens

À 25 %MAC, retrouver le bras : 4,00 + 0,25 × 4,00 = **5,00 m**. Le pourcentage décrit une position à l’intérieur d’une corde de référence ; il ne représente pas une fraction de la longueur totale de l’avion. Une valeur de 20 signifie 20 %, soit 0,20 dans la formule inverse.

**Application nouvelle :** bord d’attaque à 3,20 m, MAC de 2 m, CG à 3,70 m. Le résultat vaut 100 × 0,50/2 = **25 %**, mais le bras n’est pas 5 m. Un même pourcentage sur deux avions ou deux références ne prouve donc pas une même position en mètres.

Avant de conclure, contrôler les unités et le datum, les définitions de masse, la masse correspondant à l’enveloppe et les limitations du cas. Séparer explicitement le résultat arithmétique, le respect des plafonds présentés et les vérifications opérationnelles qui restent nécessaires.

### Lire une enveloppe et vérifier la séquence
Enveloppe fictive de cet exercice : masse autorisée de 8 000 à 12 000 kg inclus ; limite arrière 35 %MAC ; limite avant linéaire de 15 % à 8 000 kg à 25 % à 12 000 kg. Les limites sont incluses. Cette enveloppe sert uniquement à apprendre la lecture ; elle n''appartient à aucun avion réel.

Pour une masse intermédiaire M, limite avant = 15 + (M − 8 000)/4 000 × 10, exprimée en points de pourcentage. À 9 500 kg, elle vaut 18,75 %MAC. Un CG à 22 % est à l''intérieur horizontalement et la masse est admise ; un CG à 18 % est trop avant. Lire la limite à la masse du point, pas à la masse initiale ni à la seule valeur minimale 15 %.

Tracer les états successifs dans le même repère. Si une séquence consomme d''abord du carburant situé en avant du CG puis du carburant arrière, un déplacement temporaire vers l''arrière peut être suivi d''un retour. Deux extrémités à l''intérieur ne prouvent donc pas l''acceptabilité de l''état intermédiaire. Pour chaque ligne, recalculer masse, moment, bras, %MAC puis limites. Les bras constants des exercices sont une hypothèse ; utiliser les données réelles du réservoir si son bras varie.

**Auto-évaluation :** à 11 000 kg, quelle est la limite avant ? 22,5 %MAC. Un point à 21 % échoue malgré une masse inférieure au maximum. Dans les questions, la masse, le centrage et les séquences changent ; tracer avant de conclure et ne pas arrondir assez tôt pour faire entrer un point hors limite.

**Format de travail :** 15 QCM de chapitre, quatre choix et une réponse correcte. La matière 031 comporte 25 questions en 75 minutes dans le tableau EASA cité. Ce chapitre ne couvre qu’une partie de la matière ; il ne constitue pas une épreuve complète. Cadre EASA, même en anglais ou avec les unités US/britanniques.

## Références

- [EASA — programme Part-FCL](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8)
- [EASA — format de la matière, AMC1 ARA.FCL.300(b)](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67)
- [FAA Weight and Balance Handbook — principes techniques](https://www.faa.gov/sites/faa.gov/files/2023-09/Weight_Balance_Handbook.pdf)

**Support de révision ATPL(A).** À faire valider par votre instructeur. Pour un vol réel, utiliser le manuel de l’avion et les publications en vigueur. Les exemples chiffrés sont pédagogiques.', '**Objective:** Build a balance from new inputs, identify the limiting constraint, read an interpolated envelope and detect intermediate CG exceedance.

## Distinguish mass stages

![ATPL_031_001](/illustrations/en/atpl-031-001.svg)

Fictional exercise: dry operating mass DOW = 30,000 kg, traffic load = 8,000 kg. Zero-fuel mass ZFM is 38,000 kg. Adding 10,000 kg block fuel gives 48,000 kg ramp mass. After 200 kg taxi fuel, take-off mass is 47,800 kg. With 6,000 kg trip fuel consumed after take-off, planned landing mass is 41,800 kg. Do not subtract taxi fuel twice.

### Identify what each mass includes

DOW here represents the aircraft ready for the specified operation, excluding usable fuel and traffic load. It is not simply the mass of an empty structure. Check what the operating document already includes to avoid adding crew or equipment twice.

Write the chronology: DOW + traffic load = ZFM; ZFM + block fuel = ramp mass; ramp minus taxi fuel = take-off mass; take-off minus trip fuel = planned landing mass. In this exercise, no other mass is added or removed between stages, and supplied consumptions are fuel masses.

**Cross-check using remaining fuel:** 10,000 − 200 − 6,000 = 3,800 kg. ZFM 38,000 + remaining 3,800 = **41,800 kg**, matching sequential subtraction. Subtracting another 200 kg at landing would count taxi fuel twice.

The remaining 3,800 kg alone does not establish regulatory reserve compliance: this chapter checks a mass balance. Fuel allocations and operational assumptions require their own checks.

## Check every limit

Fictional limits: maximum ZFM 39,000 kg, take-off 49,000 kg, landing 42,000 kg. The initial balance satisfies these three mass ceilings without establishing performance or CG compliance. With unchanged fuel quantities, added payload cannot exceed the smallest remaining margin: 1,000, 1,200 and 200 kg respectively. Landing mass therefore limits the addition to 200 kg in this exercise. Operational limits may be more restrictive than structural maxima.

### Find the constraint that controls the decision

Write the three inequalities instead of adding margins. Additional traffic load p, with unchanged block, taxi and trip fuel, requires p ≤ 1,000 kg, p ≤ 1,200 kg and p ≤ 200 kg. All must hold together, so **p ≤ 200 kg**. Averaging the margins has no operational meaning.

**Fuel sensitivity:** separately adding 400 kg of block fuel with unchanged consumption keeps ZFM at 38,000 kg, raises take-off to 48,200 kg and landing to 42,200 kg. Take-off meets 49,000 kg, but arrival exceeds 42,000 kg by 200 kg. These inputs alone cannot determine which fuel allocation to reduce: fuel requirements must remain satisfied and the scenario must be recalculated.

**Independent case:** DOW 30,000 kg, traffic load 7,500 kg, block 10,000 kg, taxi 200 kg, trip 6,000 kg, same ceilings. ZFM 37,500, take-off 47,300, arrival 41,300 kg; margins 1,500, 1,700 and 700 kg. The limiting ceiling for added traffic load is therefore **700 kg**, subject to fixed assumptions and other constraints.

### Add a performance limit to the balance

Return to the independent case with 7,500 kg traffic load and take-off mass 47,300 kg. If a fictional performance limit imposes 47,600 kg, the additional margin is **300 kg**. It is smaller than the 700 kg allowed by the three structural ceilings alone. Add the new constraint and take the minimum again; do not arbitrarily replace all other checks with the take-off check.

If the exercise specifies unchanged consumption, use that assumption. Real operations must reassess consumption and performance after loading changes: constant trip fuel is a calculation assumption here, not a general rule.

## Recalculate CG after fuel burn

Independent CG exercise: mass 10,000 kg, CG 5.00 m from datum; total moment 50,000 kg·m. Burning 1,000 kg fuel at constant arm 6.00 m removes 6,000 kg·m. New CG is 44,000 / 9,000 = 4.889 m: it moves forward because the removed mass was aft of the initial CG. Fuel burn alone does not define a universal direction: it depends on arm and tank sequence.

### Predict direction before calculating

In this independent exercise, arms increase aft. Initially, 10,000 × 5 = 50,000 kg·m. Removing 1,000 kg at 6 m leaves 44,000 kg·m and 9,000 kg. New CG 4.8889… m is below 5 m, confirming the expected forward movement.

**Counterexample:** start from the same initial condition but remove 1,000 kg at arm 4 m. New CG is (50,000 − 4,000)/9,000 = **5.1111 m**: it moves aft. Fuel burn therefore does not always move CG forward. Tank, arm and consumption sequence must be known.

Do not reuse the 47,800 kg balance for this 10,000 kg case: the exercises are independent. For a real calculation, use the applicable tank method, especially if arm varies with quantity. Check relevant intermediate conditions; compliant endpoints do not automatically prove the entire consumption path is compliant.

### Combine loading and consumption

**Separate example:** initial state 12,000 kg at CG 5 m, giving 60,000 kg·m. Moving 300 kg from arm 7 m to 4 m removes 900 kg·m without changing mass. Then burning 1,000 kg at constant arm 6 m removes 1,000 kg and 6,000 kg·m. Final totals: 11,000 kg and 53,100 kg·m; CG = **4.827 m** rounded.

One entry per change prevents subtracting moved load from total mass or forgetting its old moment. For multiple tanks, add consumed masses and moments separately; average arm of removed fuel may be forward of, aft of or exactly at initial CG. The final result does not remove the need to check relevant intermediate conditions.

## Express CG as percentage MAC

CG as a percentage of mean aerodynamic chord is 100 × (CG arm − MAC leading-edge arm) / MAC length. Example: CG 5.00 m, leading edge 4.00 m, MAC 4.00 m: 25%. A CG at 4.80 m gives 20%. Check mass and CG against the envelope applicable to every phase, not just departure. A matching percentage does not remove the need to verify datum, configuration and aircraft-specific limits.

### Check a conversion in both directions

At 25% MAC, recover the arm: 4.00 + 0.25 × 4.00 = **5.00 m**. The percentage describes position along a reference chord; it is not a fraction of total aircraft length. A value of 20 means 20%, or 0.20 in the inverse formula.

**New application:** leading edge 3.20 m, MAC length 2 m, CG 3.70 m. The result is 100 × 0.50/2 = **25%**, but the arm is not 5 m. The same percentage on different aircraft or references therefore does not establish the same position in metres.

Before concluding, check units and datum, mass definitions, the mass relevant to the envelope and case limitations. Explicitly separate the arithmetic result, compliance with the presented ceilings and remaining operational checks.

### Read an envelope and check the sequence
Fictional exercise envelope: permitted mass from 8,000 to 12,000 kg inclusive; aft limit 35% MAC; forward limit linear from 15% at 8,000 kg to 25% at 12,000 kg. Boundaries are included. This teaching envelope belongs to no real aircraft.

At intermediate mass M, forward limit = 15 + (M − 8,000)/4,000 × 10 in percentage points. At 9,500 kg it is 18.75% MAC. A CG of 22% is horizontally inside and the mass is permitted; 18% is too far forward. Read the limit at the point''s actual mass, not the initial mass or only the minimum value of 15%.

Plot successive states on the same axes. Burning fuel forward of CG and then aft fuel can produce a temporary aft movement followed by return. Two endpoints inside therefore do not prove an acceptable intermediate state. For every row, recalculate mass, moment, arm, %MAC and limits. Constant tank arms are an exercise assumption; use actual tank data when the arm varies.

**Self-check:** at 11,000 kg, what is the forward limit? 22.5% MAC. A point at 21% fails despite mass below the maximum. Questions change mass, CG and sequences; plot before concluding and do not round early enough to move an outside point inside.

**Practice format:** 15 chapter MCQs, four choices and one correct answer. Subject 031 has 25 questions in 75 minutes in the cited EASA table. This chapter covers only part of the subject and is not a full paper. EASA framework, including with English or US/British units.

## References

- [Reference 1](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8)
- [Reference 2](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67)
- [Reference 3](https://www.faa.gov/sites/faa.gov/files/2023-09/Weight_Balance_Handbook.pdf)

**ATPL(A) revision material.** To be reviewed by your instructor. For an actual flight, use the aircraft manual and current publications. Numerical examples are for training.', 3, 1, 1, 55, '["ATPL", "mcq-four-options"]', '["Établir un bilan sur données nouvelles, identifier la contrainte limitante, lire une enveloppe interpolée et détecter un dépassement intermédiaire de centrage."]')
ON CONFLICT (id) DO UPDATE SET license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, title_fr = EXCLUDED.title_fr, title_en = EXCLUDED.title_en, content_fr = EXCLUDED.content_fr, content_en = EXCLUDED.content_en, difficulty = EXCLUDED.difficulty, order_index = EXCLUDED.order_index, level = EXCLUDED.level, duration_minutes = EXCLUDED.duration_minutes, tags = EXCLUDED.tags, learning_objectives = EXCLUDED.learning_objectives;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('08208cb9-0000-4000-8000-000008208cb9', '13833479-0000-4000-8000-000013833479', 'ATPL', 'mass_and_balance', 'mass_and_balance', 'Distinguer les étapes du bilan', 3, 3, 'DOW 28 000 kg, charge marchande 6 500 kg, carburant utilisable bloc 9 000 kg. ZFM ?', 'DOW 28,000 kg, traffic load 6,500 kg, usable block fuel 9,000 kg. ZFM?', '["34 500 kg", "43 500 kg", "28 000 kg", "37 000 kg"]', '["34,500 kg", "43,500 kg", "28,000 kg", "37,000 kg"]', '34 500 kg', 'ZFM = 28 000 + 6 500 = 34 500 kg. 43 500 ajoute le bloc ; 28 000 oublie la charge ; 37 000 ajoute le carburant à la place de la charge. Identifier le contenu avant l''addition.

![ATPL_031_001](/illustrations/atpl-031-001.svg)', 'ZFM = 28,000 + 6,500 = 34,500 kg. 43,500 adds block fuel; 28,000 omits load; 37,000 adds fuel instead of load. Identify components before adding.

![ATPL_031_001](/illustrations/en/atpl-031-001.svg)', '["ATPL", "exam-mcq", "exam-subject:031"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('08208cba-0000-4000-8000-000008208cba', '13833479-0000-4000-8000-000013833479', 'ATPL', 'mass_and_balance', 'mass_and_balance', 'Distinguer les étapes du bilan', 3, 3, 'DOW 28 000 kg, charge 6 500 kg, bloc 9 000 kg, roulage 150 kg. Masse au parking avant roulage ?', 'DOW 28,000 kg, load 6,500 kg, block fuel 9,000 kg, taxi 150 kg. Ramp mass before taxi?', '["43 350 kg", "43 650 kg", "34 500 kg", "43 500 kg"]', '["43,350 kg", "43,650 kg", "34,500 kg", "43,500 kg"]', '43 500 kg', '28 000 + 6 500 + 9 000 = 43 500. Le bloc inclut le carburant de roulage : ni le soustraire avant cette phase, ni l''ajouter une seconde fois. 34 500 est la ZFM.

![ATPL_031_001](/illustrations/atpl-031-001.svg)', '28,000 + 6,500 + 9,000 = 43,500. Block fuel includes taxi fuel: neither subtract it before taxi nor add it twice. 34,500 is ZFM.

![ATPL_031_001](/illustrations/en/atpl-031-001.svg)', '["ATPL", "exam-mcq", "exam-subject:031"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('08208cbb-0000-4000-8000-000008208cbb', '13833479-0000-4000-8000-000013833479', 'ATPL', 'mass_and_balance', 'mass_and_balance', 'Distinguer les étapes du bilan', 3, 3, 'DOW 28 000 kg, charge 6 500 kg, bloc 9 000 kg dont 150 consommés au roulage. Masse au décollage ?', 'DOW 28,000 kg, load 6,500 kg, block fuel 9,000 kg including 150 consumed in taxi. Take-off mass?', '["43 200 kg", "43 650 kg", "43 350 kg", "43 500 kg"]', '["43,200 kg", "43,650 kg", "43,350 kg", "43,500 kg"]', '43 350 kg', 'Parking 43 500 moins 150 = 43 350 kg. 43 200 retire le roulage deux fois ; 43 650 l''ajoute au lieu de le consommer. Conserver la chronologie des états.

![ATPL_031_001](/illustrations/atpl-031-001.svg)', 'Ramp 43,500 minus 150 = 43,350 kg. 43,200 subtracts taxi twice; 43,650 adds rather than burns it. Preserve state chronology.

![ATPL_031_001](/illustrations/en/atpl-031-001.svg)', '["ATPL", "exam-mcq", "exam-subject:031"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('08208cbc-0000-4000-8000-000008208cbc', '13833479-0000-4000-8000-000013833479', 'ATPL', 'mass_and_balance', 'mass_and_balance', 'Distinguer les étapes du bilan', 3, 3, 'Décollage 43 350 kg ; carburant trajet après décollage 5 500 kg ; roulage antérieur 150 kg. Masse prévue à l''atterrissage ?', 'Take-off mass 43,350 kg; trip fuel after take-off 5,500 kg; previous taxi fuel 150 kg. Planned landing mass?', '["43 200 kg", "37 850 kg", "37 700 kg", "38 000 kg"]', '["43,200 kg", "37,850 kg", "37,700 kg", "38,000 kg"]', '37 850 kg', '43 350 − 5 500 = 37 850. Le roulage est déjà compris dans l''état de départ au décollage ; le retirer ou le rajouter change à tort la référence. 43 200 ne retire que le roulage, pas le trajet.

![ATPL_031_001](/illustrations/atpl-031-001.svg)', '43,350 − 5,500 = 37,850. Taxi is already accounted for in take-off mass; subtracting or adding it changes the reference incorrectly. 43,200 removes only taxi, not trip fuel.

![ATPL_031_001](/illustrations/en/atpl-031-001.svg)', '["ATPL", "exam-mcq", "exam-subject:031"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('0820907e-0000-4000-8000-00000820907e', '13833479-0000-4000-8000-000013833479', 'ATPL', 'mass_and_balance', 'mass_and_balance', 'Vérifier chaque limitation', 3, 3, 'États initiaux ZFM/décollage/arrivée : 34 500/43 350/37 850 kg. Plafonds : 36 000/45 000/39 000 kg. Ajout de charge 1 400 kg, carburants inchangés. Dépassement ?', 'Initial ZFM/take-off/landing: 34,500/43,350/37,850 kg. Ceilings: 36,000/45,000/39,000 kg. Add 1,400 kg load, fuel unchanged. Exceedance?', '["Atterrissage de 250 kg", "ZFM de 100 kg", "Décollage de 250 kg", "Aucun dépassement"]', '["Landing by 250 kg", "ZFM by 100 kg", "Take-off by 250 kg", "No exceedance"]', 'Atterrissage de 250 kg', 'Nouveaux états 35 900/44 750/39 250 : seule l''arrivée dépasse, de 250. Une marge restante au décollage ne compense pas l''arrivée ; chaque plafond s''applique à son état.

![ATPL_031_001](/illustrations/atpl-031-001.svg)', 'New states 35,900/44,750/39,250: only landing exceeds, by 250. Remaining take-off margin cannot compensate for landing exceedance; each ceiling applies to its state.

![ATPL_031_001](/illustrations/en/atpl-031-001.svg)', '["ATPL", "exam-mcq", "exam-subject:031"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('0820907f-0000-4000-8000-00000820907f', '13833479-0000-4000-8000-000013833479', 'ATPL', 'mass_and_balance', 'mass_and_balance', 'Vérifier chaque limitation', 3, 3, 'DOW 28 000, charge 6 500, bloc 9 000, roulage 150, trajet 5 500 kg. Maxima ZFM/décollage/arrivée 36 000/45 000/39 000 ; limite de performance au décollage 43 800 kg. Ajout maximal de charge, carburants inchangés ?', 'DOW 28,000, load 6,500, block 9,000, taxi 150, trip 5,500 kg. ZFM/take-off/landing maxima 36,000/45,000/39,000; take-off performance limit 43,800 kg. Maximum added load with fuel unchanged?', '["1 150 kg", "1 500 kg", "1 650 kg", "450 kg"]', '["1,150 kg", "1,500 kg", "1,650 kg", "450 kg"]', '450 kg', 'États 34 500/43 350/37 850. Marges structurales 1 500/1 650/1 150 ; marge de performance 450. Le minimum des quatre vaut 450 ; prendre 1 150 ignorerait la performance.

![ATPL_031_001](/illustrations/atpl-031-001.svg)', 'States 34,500/43,350/37,850. Structural margins 1,500/1,650/1,150; performance margin 450. The minimum is 450; choosing 1,150 ignores performance.

![ATPL_031_001](/illustrations/en/atpl-031-001.svg)', '["ATPL", "exam-mcq", "exam-subject:031"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('08209080-0000-4000-8000-000008209080', '13833479-0000-4000-8000-000013833479', 'ATPL', 'mass_and_balance', 'mass_and_balance', 'Vérifier chaque limitation', 3, 3, 'ZFM 34 500, bloc 9 000, roulage 150, trajet 5 500 kg. Ajouter 1 600 kg de bloc, consommations inchangées. Plafonds ZFM/décollage/arrivée 36 000/45 000/39 000. Retrait minimal de charge selon ces trois plafonds ?', 'ZFM 34,500, block 9,000, taxi 150, trip 5,500 kg. Add 1,600 kg block fuel with unchanged consumption. ZFM/take-off/landing ceilings 36,000/45,000/39,000. Minimum load removal under these three ceilings?', '["1 150 kg", "1 600 kg", "450 kg", "50 kg"]', '["1,150 kg", "1,600 kg", "450 kg", "50 kg"]', '450 kg', 'Avec ajout : décollage 44 950, arrivée 39 450, ZFM inchangée. Il faut enlever 450 kg de charge pour l''arrivée ; 50 est la marge de décollage, pas le retrait requis. Enlever tout le carburant ajouté comme charge serait excessif.

![ATPL_031_001](/illustrations/atpl-031-001.svg)', 'After addition: take-off 44,950, landing 39,450, ZFM unchanged. Remove 450 kg load for landing; 50 is take-off margin, not required removal. Removing load equal to all added fuel would be excessive.

![ATPL_031_001](/illustrations/en/atpl-031-001.svg)', '["ATPL", "exam-mcq", "exam-subject:031"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('08209081-0000-4000-8000-000008209081', '13833479-0000-4000-8000-000013833479', 'ATPL', 'mass_and_balance', 'mass_and_balance', 'Vérifier chaque limitation', 3, 3, 'Marges ZFM/décollage/arrivée : 1 500/1 650/1 150 kg. Carburants inchangés, aucun autre plafond fourni : ajout maximal de charge selon ces marges ?', 'ZFM/take-off/landing margins: 1,500/1,650/1,150 kg. Fuel unchanged, no other ceiling provided: maximum added load under these margins?', '["4 300 kg", "1 150 kg", "1 500 kg", "1 650 kg"]', '["4,300 kg", "1,150 kg", "1,500 kg", "1,650 kg"]', '1 150 kg', 'Chaque kilogramme ajouté doit satisfaire les trois plafonds ; retenir 1 150. La somme 4 300 cumule des marges qui ne sont pas interchangeables ; les deux autres choix excèdent la marge d''arrivée.

![ATPL_031_001](/illustrations/atpl-031-001.svg)', 'Every added kilogram must meet all three ceilings; choose 1,150. Summing to 4,300 treats non-interchangeable margins as additive; the other values exceed landing margin.

![ATPL_031_001](/illustrations/en/atpl-031-001.svg)', '["ATPL", "exam-mcq", "exam-subject:031"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('08209082-0000-4000-8000-000008209082', '13833479-0000-4000-8000-000013833479', 'ATPL', 'mass_and_balance', 'mass_and_balance', 'Vérifier chaque limitation', 3, 3, 'Masse au décollage 47 800 kg, maximum structural 49 000 kg, limite de performance du cas 46 500 kg. Conclusion ?', 'Take-off mass 47,800 kg, structural maximum 49,000 kg, case performance limit 46,500 kg. Conclusion?', '["Maximum structural respecté, limite de performance dépassée", "Les deux limites sont respectées", "Maximum structural dépassé, limite de performance respectée", "La moyenne des deux limites autorise le départ"]', '["Structural maximum met, performance limit exceeded", "Both limits are met", "Structural maximum exceeded, performance limit met", "The average of both limits permits departure"]', 'Maximum structural respecté, limite de performance dépassée', 'La limite de performance 46 500 kg est plus restrictive que le maximum structural 49 000 kg. 47 800 kg respecte ce dernier mais dépasse la limite de performance de 1 300 kg.

![ATPL_031_001](/illustrations/atpl-031-001.svg)', 'The 46,500 kg performance limit is more restrictive than the 49,000 kg structural maximum. 47,800 kg meets the latter but exceeds the performance limit by 1,300 kg.

![ATPL_031_001](/illustrations/en/atpl-031-001.svg)', '["ATPL", "exam-mcq", "exam-subject:031"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('08209444-0000-4000-8000-000008209444', '13833479-0000-4000-8000-000013833479', 'ATPL', 'mass_and_balance', 'mass_and_balance', 'Recalculer le centrage après consommation', 3, 3, 'État initial 10 000 kg au CG 5,00 m. On déplace 200 kg du bras 7,00 m au bras 3,00 m, puis consomme 1 000 kg au bras constant 6,00 m. Même datum, aucun autre changement : CG final arrondi à 0,001 m ?', 'Initial state 10,000 kg at CG 5.00 m. Move 200 kg from arm 7.00 m to 3.00 m, then burn 1,000 kg at constant arm 6.00 m. Common datum, no other change: final CG rounded to 0.001 m?', '["4,889 m", "4,320 m", "4,978 m", "4,800 m"]', '["4.889 m", "4.320 m", "4.978 m", "4.800 m"]', '4,800 m', 'Moment initial 50 000 ; déplacement 200 × (3 − 7) = −800 ; consommation −6 000 kg·m. Moment final 43 200, masse 9 000 : CG 4,800 m. 4,889 oublie le déplacement, 4,320 conserve la masse initiale, 4,978 inverse le signe du déplacement.

![ATPL_031_001](/illustrations/mass-balance-method.svg)', 'Initial moment 50,000; shift 200 × (3 − 7) = −800; fuel burn −6,000 kg·m. Final moment 43,200, mass 9,000: CG 4.800 m. 4.889 ignores the shift, 4.320 retains initial mass, 4.978 reverses the shift sign.

![ATPL_031_001](/illustrations/en/mass-balance-method.svg)', '["ATPL", "exam-mcq", "exam-subject:031"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('7bf1f374-0000-4000-8000-00007bf1f374', '13833479-0000-4000-8000-000013833479', 'ATPL', 'mass_and_balance', 'mass_and_balance', 'Recalculer le centrage après consommation', 3, 3, 'État initial 10 000 kg au CG 5,00 m. On consomme 500 kg d’un réservoir au bras constant 6,00 m et 500 kg d’un autre au bras constant 4,00 m. Même datum, aucun autre changement : CG final ?', 'Initial state 10,000 kg at CG 5.00 m. Burn 500 kg from a tank at constant arm 6.00 m and 500 kg from another at constant arm 4.00 m. Common datum, no other change: final CG?', '["5,111 m", "4,500 m", "5,000 m", "4,889 m"]', '["5.111 m", "4.500 m", "5.000 m", "4.889 m"]', '5,000 m', 'Moment retiré : 500 × 6 + 500 × 4 = 5 000 kg·m. CG = (50 000 − 5 000)/(10 000 − 1 000) = 5,000 m. Le bras moyen du carburant consommé vaut le CG initial. Les valeurs 4,889 et 5,111 supposent à tort toute la consommation dans un seul réservoir.

![ATPL_031_001](/illustrations/mass-balance-method.svg)', 'Removed moment: 500 × 6 + 500 × 4 = 5,000 kg·m. CG = (50,000 − 5,000)/(10,000 − 1,000) = 5.000 m. Average arm of burned fuel equals initial CG. 4.889 and 5.111 incorrectly put all fuel burn in one tank.

![ATPL_031_001](/illustrations/en/mass-balance-method.svg)', '["ATPL", "exam-mcq", "exam-subject:031"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('7bf1f375-0000-4000-8000-00007bf1f375', '13833479-0000-4000-8000-000013833479', 'ATPL', 'mass_and_balance', 'mass_and_balance', 'Recalculer le centrage après consommation', 3, 3, 'Enveloppe fictive jointe : masse 8 000–12 000 kg, limite avant linéaire de 15 à 25 %MAC, arrière 35 %. À 10 000 kg et 19 %MAC, quel diagnostic ?

![ATPL_031_001](/illustrations/atpl-031-001.svg)', 'Attached fictional envelope: mass 8,000–12,000 kg, forward limit linear from 15 to 25% MAC, aft limit 35%. At 10,000 kg and 19% MAC, what follows?

![ATPL_031_001](/illustrations/en/atpl-031-001.svg)', '["Masse excessive de 2 000 kg", "Masse admise, CG un point de pourcentage trop avant", "Conforme parce que 19 est entre 15 et 35", "CG un point de pourcentage trop arrière"]', '["Mass excessive by 2,000 kg", "Mass permitted, CG one percentage point too far forward", "Compliant because 19 is between 15 and 35", "CG one percentage point too far aft"]', 'Masse admise, CG un point de pourcentage trop avant', 'À 10 000 kg, fraction (10 000−8 000)/4 000 = 0,5 ; limite avant 15+0,5×10 = 20 %. Le CG 19 % est un point trop avant. Utiliser 15 % lirait la limite d''une autre masse.

![ATPL_031_001](/illustrations/atpl-031-001.svg)', 'At 10,000 kg, fraction (10,000−8,000)/4,000 = 0.5; forward limit 15+0.5×10 = 20%. CG 19% is one point too far forward. Using 15% reads the limit for a different mass.

![ATPL_031_001](/illustrations/en/atpl-031-001.svg)', '["ATPL", "exam-mcq", "exam-subject:031"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('7bf1f376-0000-4000-8000-00007bf1f376', '13833479-0000-4000-8000-000013833479', 'ATPL', 'mass_and_balance', 'mass_and_balance', 'Recalculer le centrage après consommation', 3, 3, 'Même enveloppe fictive. Départ 10 000 kg, CG 4,60 m ; LEMAC 4 m, MAC 2 m. Consommer d''abord 1 000 kg au bras 2 m, puis 1 000 kg au bras 7 m ; bras constants. Quel état est hors enveloppe parmi départ, après premier réservoir, fin ?

![ATPL_031_001](/illustrations/atpl-031-001.svg)', 'Same fictional envelope. Start 10,000 kg, CG 4.60 m; LEMAC 4 m, MAC 2 m. Burn 1,000 kg at arm 2 m, then 1,000 kg at arm 7 m; constant arms. Which of start, after first tank and final states is outside?

![ATPL_031_001](/illustrations/en/atpl-031-001.svg)', '["Seulement après le premier réservoir", "Seulement le départ", "Seulement l''état final", "Aucun, puisque départ et fin sont conformes"]', '["Only after the first tank", "Only the start", "Only the final state", "None, because start and finish comply"]', 'Seulement après le premier réservoir', 'Départ : 30 % à 10 000 kg, admis. Après premier retrait : 44 000/9 000 = 4,889 m, soit 44,44 %, au-delà de 35. Fin : 37 000/8 000 = 4,625 m, soit 31,25 %, admis. Des extrémités conformes masquent ici un dépassement intermédiaire.

![ATPL_031_001](/illustrations/atpl-031-001.svg)', 'Start: 30% at 10,000 kg, inside. After first burn: 44,000/9,000 = 4.889 m, or 44.44%, beyond 35. Final: 37,000/8,000 = 4.625 m, or 31.25%, inside. Compliant endpoints conceal an intermediate exceedance.

![ATPL_031_001](/illustrations/en/atpl-031-001.svg)', '["ATPL", "exam-mcq", "exam-subject:031"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('7bf267d6-0000-4000-8000-00007bf267d6', '13833479-0000-4000-8000-000013833479', 'ATPL', 'mass_and_balance', 'mass_and_balance', 'Exprimer le centrage en pourcentage de corde', 3, 3, 'CG 3,65 m, LEMAC 3,00 m, longueur MAC 2,60 m. Position en %MAC ?', 'CG 3.65 m, LEMAC 3.00 m, MAC length 2.60 m. Position in %MAC?', '["140,4 %", "65 %", "0,25 %", "25 %"]', '["140.4%", "65%", "0.25%", "25%"]', '25 %', '100 × (3,65−3,00)/2,60 = 25 %. 140,4 omet LEMAC ; 65 omet la division par la longueur ; 0,25 omet la conversion du rapport en pourcentage.

![ATPL_031_001](/illustrations/atpl-031-001.svg)', '100 × (3.65−3.00)/2.60 = 25%. 140.4 omits LEMAC; 65 omits division by length; 0.25 omits conversion of ratio to percentage.

![ATPL_031_001](/illustrations/en/atpl-031-001.svg)', '["ATPL", "exam-mcq", "exam-subject:031"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('7bf267d7-0000-4000-8000-00007bf267d7', '13833479-0000-4000-8000-000013833479', 'ATPL', 'mass_and_balance', 'mass_and_balance', 'Exprimer le centrage en pourcentage de corde', 3, 3, 'Enveloppe fictive jointe. Masse 9 000 kg, moment 41 400 kg·m, LEMAC 4 m, MAC 2 m. Quelle conclusion correspond à ce seul contrôle masse/centrage ?

![ATPL_031_001](/illustrations/atpl-031-001.svg)', 'Attached fictional envelope. Mass 9,000 kg, moment 41,400 kg·m, LEMAC 4 m, MAC 2 m. What follows from this mass/CG check alone?

![ATPL_031_001](/illustrations/en/atpl-031-001.svg)', '["CG 30 %MAC, trop avant car limite avant 35 %", "CG 20 %MAC, conforme à toutes les exigences du vol", "CG 30 %MAC, dans les limites 17,5–35 % de cette masse", "CG 230 %MAC, trop arrière"]', '["CG 30% MAC, too far forward because forward limit is 35%", "CG 20% MAC, compliant with every flight requirement", "CG 30% MAC, within this mass''s 17.5–35% limits", "CG 230% MAC, too far aft"]', 'CG 30 %MAC, dans les limites 17,5–35 % de cette masse', 'Bras = 41 400/9 000 = 4,60 m ; %MAC = 100×0,60/2 = 30. Limite avant à 9 000 : 15+1 000/4 000×10 = 17,5 %. Masse et CG sont admis dans cette enveloppe ; performances, carburant et états ultérieurs restent des contrôles distincts.

![ATPL_031_001](/illustrations/atpl-031-001.svg)', 'Arm = 41,400/9,000 = 4.60 m; %MAC = 100×0.60/2 = 30. Forward limit at 9,000: 15+1,000/4,000×10 = 17.5%. Mass and CG meet this envelope; performance, fuel and later states remain separate checks.

![ATPL_031_001](/illustrations/en/atpl-031-001.svg)', '["ATPL", "exam-mcq", "exam-subject:031"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO lessons (id, license, category, theme, title_fr, title_en, content_fr, content_en, difficulty, order_index, level, duration_minutes, tags, learning_objectives) VALUES
('18a1be56-0000-4000-8000-000018a1be56', 'ATPL', 'navigation', 'navigation', 'Coordonnées, routes et temps UTC', 'Coordinates, tracks and UTC', '**Objectif :** Résoudre et vérifier les exercices de coordonnées, routes et temps utc.

## Lire les coordonnées

![ATPL_061_001](/illustrations/atpl-061-001.svg)

La latitude mesure l’angle au nord ou au sud de l’équateur ; la longitude l’angle à l’est ou à l’ouest du méridien origine. Dans le modèle sphérique d’exercice, 1 minute de latitude vaut environ 1 NM ; 1 degré de latitude vaut 60 NM. Deux points du même méridien séparés de 2°30′ de latitude sont distants d’environ 150 NM. Ne pas appliquer directement cette règle à une différence de longitude hors de l’équateur.

## Distance sur un parallèle

Sur un parallèle de latitude φ, distance ≈ 60 × différence de longitude en degrés × cos φ, en NM. À 60° de latitude, cos φ = 0,5 : 4° de longitude représentent 120 NM le long du parallèle. Cette distance n’est généralement pas celle de l’orthodromie entre les points. Vérifier le passage du méridien 180° : entre 179°E et 179°W, le petit écart est 2°, pas 358°.

## Orthodromie et loxodromie

Sur une sphère, l’arc mineur de grand cercle donne la plus courte distance entre deux points non antipodaux. Une loxodromie coupe les méridiens sous un angle constant : sa route vraie est constante. Sur une projection de Mercator, elle apparaît droite, alors qu’une orthodromie est généralement courbe. Les méridiens et l’équateur sont des cas particuliers de grands cercles. Pour une route orthodromique, la route vraie change généralement en cours de trajet.

## Conventions de cap et temps

Dans les exercices, route magnétique = route vraie − déclinaison Est, ou + déclinaison Ouest ; normaliser entre 000° et 359°. Le cap nécessite ensuite une correction de vent. Pour UTC, conserver la date lors du passage de minuit : 23:40 + 1 h 35 = 01:15 le lendemain. La rotation terrestre de 15° par heure donne une différence de temps solaire, pas automatiquement le fuseau civil, qui suit des choix administratifs.

**Format de travail :** 15 QCM de chapitre, quatre choix et une réponse correcte. La matière 061 comporte 55 questions en 135 minutes dans le tableau EASA cité. Ce chapitre ne couvre qu’une partie de la matière ; il ne constitue pas une épreuve complète. Cadre EASA, même en anglais ou avec les unités US/britanniques.

## Références

- [EASA — programme Part-FCL](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8)
- [EASA — format de la matière, AMC1 ARA.FCL.300(b)](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67)
- [FAA PHAK — navigation](https://www.faa.gov/regulationspolicies/handbooksmanuals/aviation/phak/chapter-16-navigation)

**Support de révision ATPL(A).** À faire valider par votre instructeur. Pour un vol réel, utiliser le manuel de l’avion et les publications en vigueur. Les exemples chiffrés sont pédagogiques.', '**Objective:** Solve and cross-check exercises on coordinates, tracks and utc.

## Read coordinates

![ATPL_061_001](/illustrations/en/atpl-061-001.svg)

Latitude measures angle north or south of the equator; longitude measures angle east or west of the prime meridian. In the exercise’s spherical model, one minute of latitude is approximately 1 NM; one degree is 60 NM. Two points on the same meridian separated by 2°30′ latitude are approximately 150 NM apart. Do not directly apply this rule to longitude differences away from the equator.

## Distance along a parallel

Along a parallel at latitude φ, distance ≈ 60 × longitude difference in degrees × cos φ, in NM. At latitude 60°, cos φ = 0.5: 4° longitude represents 120 NM along the parallel. This generally differs from great-circle distance between the points. Check the 180° meridian crossing: between 179°E and 179°W, the smaller difference is 2°, not 358°.

## Great circle and rhumb line

On a sphere, the minor great-circle arc gives the shortest distance between two non-antipodal points. A rhumb line crosses meridians at a constant angle: its true course is constant. On a Mercator chart it appears straight, while a great circle generally curves. Meridians and the equator are special great-circle cases. True course generally changes along a great-circle route.

## Heading and time conventions

In these exercises, magnetic course = true course − east variation, or + west variation; normalise between 000° and 359°. Heading then requires wind correction. For UTC, retain the date across midnight: 23:40 + 1 h 35 = 01:15 next day. Earth’s rotation of 15° per hour gives a solar-time difference, not automatically the civil time zone, which follows administrative choices.

**Practice format:** 15 chapter MCQs, four choices and one correct answer. Subject 061 has 55 questions in 135 minutes in the cited EASA table. This chapter covers only part of the subject and is not a full paper. EASA framework, including with English or US/British units.

## References

- [Reference 1](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8)
- [Reference 2](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67)
- [Reference 3](https://www.faa.gov/regulationspolicies/handbooksmanuals/aviation/phak/chapter-16-navigation)

**ATPL(A) revision material.** To be reviewed by your instructor. For an actual flight, use the aircraft manual and current publications. Numerical examples are for training.', 3, 1, 1, 35, '["ATPL", "mcq-four-options"]', '["Résoudre et vérifier les exercices de coordonnées, routes et temps utc."]')
ON CONFLICT (id) DO UPDATE SET license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, title_fr = EXCLUDED.title_fr, title_en = EXCLUDED.title_en, content_fr = EXCLUDED.content_fr, content_en = EXCLUDED.content_en, difficulty = EXCLUDED.difficulty, order_index = EXCLUDED.order_index, level = EXCLUDED.level, duration_minutes = EXCLUDED.duration_minutes, tags = EXCLUDED.tags, learning_objectives = EXCLUDED.learning_objectives;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('2bc663d9-0000-4000-8000-00002bc663d9', '18a1be56-0000-4000-8000-000018a1be56', 'ATPL', 'navigation', 'navigation', 'Lire les coordonnées', 3, 3, 'Sur un même méridien, écart de latitude 2°30′. Distance sphérique approximative ?

![ATPL_061_001](/illustrations/atpl-061-001.svg)', 'On one meridian, latitude difference 2°30′. Approximate spherical distance?

![ATPL_061_001](/illustrations/en/atpl-061-001.svg)', '["150 NM", "230 NM", "120 NM", "30 NM"]', '["150 NM", "230 NM", "120 NM", "30 NM"]', '150 NM', '2,5 × 60 = 150 NM.

![ATPL_061_001](/illustrations/atpl-061-001.svg)', '2.5 × 60 = 150 NM.

![ATPL_061_001](/illustrations/en/atpl-061-001.svg)', '["ATPL", "exam-mcq", "exam-subject:061"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('2bc663da-0000-4000-8000-00002bc663da', '18a1be56-0000-4000-8000-000018a1be56', 'ATPL', 'navigation', 'navigation', 'Lire les coordonnées', 3, 3, 'Une minute de latitude correspond approximativement à…', 'One minute of latitude corresponds approximately to…', '["60 NM", "1 degré", "1 km exactement", "1 NM"]', '["60 NM", "1 degree", "Exactly 1 km", "1 NM"]', '1 NM', 'Approximation du modèle sphérique de navigation.

![ATPL_061_001](/illustrations/atpl-061-001.svg)', 'Approximation of the spherical navigation model.

![ATPL_061_001](/illustrations/en/atpl-061-001.svg)', '["ATPL", "exam-mcq", "exam-subject:061"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('2bc6679c-0000-4000-8000-00002bc6679c', '18a1be56-0000-4000-8000-000018a1be56', 'ATPL', 'navigation', 'navigation', 'Distance sur un parallèle', 3, 3, 'Latitude 60°, écart de longitude 4°, cos 60° = 0,5. Distance le long du parallèle ?', 'Latitude 60°, longitude difference 4°, cos 60° = 0.5. Distance along the parallel?', '["480 NM", "60 NM", "120 NM", "240 NM"]', '["480 NM", "60 NM", "120 NM", "240 NM"]', '120 NM', '60 × 4 × 0,5 = 120 NM.

![ATPL_061_001](/illustrations/atpl-061-001.svg)', '60 × 4 × 0.5 = 120 NM.

![ATPL_061_001](/illustrations/en/atpl-061-001.svg)', '["ATPL", "exam-mcq", "exam-subject:061"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('2bc6679d-0000-4000-8000-00002bc6679d', '18a1be56-0000-4000-8000-000018a1be56', 'ATPL', 'navigation', 'navigation', 'Distance sur un parallèle', 3, 3, 'À l’équateur, 4° de longitude représentent approximativement…', 'At the equator, 4° longitude represents approximately…', '["960 NM", "240 NM", "120 NM", "4 NM"]', '["960 NM", "240 NM", "120 NM", "4 NM"]', '240 NM', 'cos 0° = 1 : 60 × 4 = 240 NM.

![ATPL_061_001](/illustrations/atpl-061-001.svg)', 'cos 0° = 1: 60 × 4 = 240 NM.

![ATPL_061_001](/illustrations/en/atpl-061-001.svg)', '["ATPL", "exam-mcq", "exam-subject:061"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('2bc6679e-0000-4000-8000-00002bc6679e', '18a1be56-0000-4000-8000-000018a1be56', 'ATPL', 'navigation', 'navigation', 'Distance sur un parallèle', 3, 3, 'Entre 179°E et 179°W, le plus petit écart de longitude est…', 'Between 179°E and 179°W, the smaller longitude difference is…', '["2°", "358°", "180°", "0°"]', '["2°", "358°", "180°", "0°"]', '2°', 'On traverse le méridien 180° sur 2°.

![ATPL_061_001](/illustrations/atpl-061-001.svg)', 'The short crossing of the 180° meridian spans 2°.

![ATPL_061_001](/illustrations/en/atpl-061-001.svg)', '["ATPL", "exam-mcq", "exam-subject:061"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('2bc6679f-0000-4000-8000-00002bc6679f', '18a1be56-0000-4000-8000-000018a1be56', 'ATPL', 'navigation', 'navigation', 'Distance sur un parallèle', 3, 3, 'La formule 60 × Δlongitude × cos latitude donne ici…', 'Here, 60 × Δlongitude × cos latitude gives…', '["Toujours la plus courte distance", "La vitesse sol", "Une distance verticale", "La distance le long du parallèle"]', '["Always the shortest distance", "Groundspeed", "Vertical distance", "Distance along the parallel"]', 'La distance le long du parallèle', 'Ne pas confondre parallèle et grand cercle.

![ATPL_061_001](/illustrations/atpl-061-001.svg)', 'Do not confuse a parallel with a great circle.

![ATPL_061_001](/illustrations/en/atpl-061-001.svg)', '["ATPL", "exam-mcq", "exam-subject:061"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('2bc66b61-0000-4000-8000-00002bc66b61', '18a1be56-0000-4000-8000-000018a1be56', 'ATPL', 'navigation', 'navigation', 'Orthodromie et loxodromie', 3, 3, 'La plus courte route sur une sphère entre deux points non antipodaux suit…', 'The shortest spherical route between two non-antipodal points follows…', '["Toute loxodromie", "Une route magnétique toujours constante", "L’arc mineur de grand cercle", "Tout parallèle"]', '["Any rhumb line", "An always-constant magnetic course", "The minor great-circle arc", "Any parallel"]', 'L’arc mineur de grand cercle', 'C’est la propriété géométrique de l’orthodromie.

![ATPL_061_001](/illustrations/atpl-061-001.svg)', 'This is the geometric property of a great circle.

![ATPL_061_001](/illustrations/en/atpl-061-001.svg)', '["ATPL", "exam-mcq", "exam-subject:061"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('2bc66b62-0000-4000-8000-00002bc66b62', '18a1be56-0000-4000-8000-000018a1be56', 'ATPL', 'navigation', 'navigation', 'Orthodromie et loxodromie', 3, 3, 'Une loxodromie conserve…', 'A rhumb line maintains…', '["Une vitesse constante", "Un angle constant avec les méridiens", "La plus courte distance dans tous les cas", "Une latitude constante dans tous les cas"]', '["Constant speed", "A constant angle to meridians", "The shortest distance in every case", "Constant latitude in every case"]', 'Un angle constant avec les méridiens', 'Sa route vraie est constante.

![ATPL_061_001](/illustrations/atpl-061-001.svg)', 'Its true course is constant.

![ATPL_061_001](/illustrations/en/atpl-061-001.svg)', '["ATPL", "exam-mcq", "exam-subject:061"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('2bc66b63-0000-4000-8000-00002bc66b63', '18a1be56-0000-4000-8000-000018a1be56', 'ATPL', 'navigation', 'navigation', 'Orthodromie et loxodromie', 3, 3, 'Sur une carte Mercator, une loxodromie apparaît…', 'On a Mercator chart, a rhumb line appears…', '["Droite", "Toujours circulaire", "Toujours sinusoïdale", "Invisible"]', '["Straight", "Always circular", "Always sinusoidal", "Invisible"]', 'Droite', 'C’est une propriété de cette projection.

![ATPL_061_001](/illustrations/atpl-061-001.svg)', 'This is a property of the projection.

![ATPL_061_001](/illustrations/en/atpl-061-001.svg)', '["ATPL", "exam-mcq", "exam-subject:061"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('2bc66b64-0000-4000-8000-00002bc66b64', '18a1be56-0000-4000-8000-000018a1be56', 'ATPL', 'navigation', 'navigation', 'Orthodromie et loxodromie', 3, 3, 'En suivant une orthodromie générale, la route vraie…', 'Along a general great-circle route, true course…', '["Est toujours constante", "Est toujours 090°", "Est identique au vent", "Change généralement"]', '["Is always constant", "Is always 090°", "Equals the wind direction", "Generally changes"]', 'Change généralement', 'La convergence des méridiens intervient.

![ATPL_061_001](/illustrations/atpl-061-001.svg)', 'Meridian convergence matters.

![ATPL_061_001](/illustrations/en/atpl-061-001.svg)', '["ATPL", "exam-mcq", "exam-subject:061"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('4d0774b3-0000-4000-8000-00004d0774b3', '18a1be56-0000-4000-8000-000018a1be56', 'ATPL', 'navigation', 'navigation', 'Conventions de cap et temps', 3, 3, 'Route vraie 100°, déclinaison 8° Est. Route magnétique selon la convention donnée ?', 'True course 100°, variation 8° east. Magnetic course using the stated convention?', '["082°", "100°", "092°", "108°"]', '["082°", "100°", "092°", "108°"]', '092°', 'Est : soustraire 8°.

![ATPL_061_001](/illustrations/atpl-061-001.svg)', 'East: subtract 8°.

![ATPL_061_001](/illustrations/en/atpl-061-001.svg)', '["ATPL", "exam-mcq", "exam-subject:061"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('4d0774b4-0000-4000-8000-00004d0774b4', '18a1be56-0000-4000-8000-000018a1be56', 'ATPL', 'navigation', 'navigation', 'Conventions de cap et temps', 3, 3, 'Route vraie 005°, déclinaison 10° Est. Route magnétique normalisée ?', 'True course 005°, variation 10° east. Normalised magnetic course?', '["350°", "355°", "015°", "005°"]', '["350°", "355°", "015°", "005°"]', '355°', '5 − 10 = −5°, soit 355°.

![ATPL_061_001](/illustrations/atpl-061-001.svg)', '5 − 10 = −5°, equivalent to 355°.

![ATPL_061_001](/illustrations/en/atpl-061-001.svg)', '["ATPL", "exam-mcq", "exam-subject:061"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('4d0774b5-0000-4000-8000-00004d0774b5', '18a1be56-0000-4000-8000-000018a1be56', 'ATPL', 'navigation', 'navigation', 'Conventions de cap et temps', 3, 3, 'Départ 23:40 UTC, durée 1 h 35. Arrivée UTC ?', 'Departure 23:40 UTC, duration 1 h 35. Arrival UTC?', '["01:15 le lendemain", "01:15 le même jour", "00:75 le lendemain", "22:05 le même jour"]', '["01:15 next day", "01:15 same day", "00:75 next day", "22:05 same day"]', '01:15 le lendemain', 'Additionner la durée et conserver le changement de date.

![ATPL_061_001](/illustrations/atpl-061-001.svg)', 'Add duration and retain the date change.

![ATPL_061_001](/illustrations/en/atpl-061-001.svg)', '["ATPL", "exam-mcq", "exam-subject:061"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('4d0774b6-0000-4000-8000-00004d0774b6', '18a1be56-0000-4000-8000-000018a1be56', 'ATPL', 'navigation', 'navigation', 'Conventions de cap et temps', 3, 3, '30° de longitude correspondent à quelle différence de temps solaire ?', '30° longitude corresponds to which solar-time difference?', '["30 min", "4 h", "15 h", "2 h"]', '["30 min", "4 h", "15 h", "2 h"]', '2 h', '30/15 = 2 h.

![ATPL_061_001](/illustrations/atpl-061-001.svg)', '30/15 = 2 h.

![ATPL_061_001](/illustrations/en/atpl-061-001.svg)', '["ATPL", "exam-mcq", "exam-subject:061"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, options_en, answer_key, explanation_fr, explanation_en, tags) VALUES
('4d0774b7-0000-4000-8000-00004d0774b7', '18a1be56-0000-4000-8000-000018a1be56', 'ATPL', 'navigation', 'navigation', 'Conventions de cap et temps', 3, 3, 'Peut-on déduire automatiquement l’heure civile d’un lieu de sa longitude seule ?', 'Can civil time automatically be determined from longitude alone?', '["Oui, en été uniquement", "Oui, si la latitude est connue", "Non", "Oui, sans exception"]', '["Yes, only in summer", "Yes, if latitude is known", "No", "Yes, without exception"]', 'Non', 'Fuseaux et heure saisonnière suivent des règles administratives.

![ATPL_061_001](/illustrations/atpl-061-001.svg)', 'Time zones and seasonal time follow administrative rules.

![ATPL_061_001](/illustrations/en/atpl-061-001.svg)', '["ATPL", "exam-mcq", "exam-subject:061"]')
ON CONFLICT (id) DO UPDATE SET lesson_id = EXCLUDED.lesson_id, license = EXCLUDED.license, category = EXCLUDED.category, theme = EXCLUDED.theme, subtopic = EXCLUDED.subtopic, difficulty = EXCLUDED.difficulty, level = EXCLUDED.level, question_fr = EXCLUDED.question_fr, question_en = EXCLUDED.question_en, options = EXCLUDED.options, options_en = EXCLUDED.options_en, answer_key = EXCLUDED.answer_key, explanation_fr = EXCLUDED.explanation_fr, explanation_en = EXCLUDED.explanation_en, tags = EXCLUDED.tags;

COMMIT;
