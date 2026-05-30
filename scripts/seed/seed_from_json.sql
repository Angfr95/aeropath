-- ============================================================================
-- SEED FROM JSON : 3 modules chargés depuis content/fr/ppl/air_law/
-- Généré par scripts/seed_from_json.go
-- ============================================================================

-- Nettoyage des données existantes
DELETE FROM student_question_history;
DELETE FROM questions;
DELETE FROM lessons;
DELETE FROM user_gamification;
DELETE FROM user_progress;
DELETE FROM achievements;
DELETE FROM leaderboard_weekly;
DELETE FROM user_quests;
DELETE FROM spaced_repetition;
DELETE FROM app_config;

-- ============================================================================
-- 1. APP CONFIG
-- ============================================================================
INSERT INTO app_config (config_key, config_value) VALUES
('adaptive_learning', '{
    "level_1": {"pass_threshold": 0.80, "hint_enabled": true, "time_limit_sec": null, "immediate_feedback": true},
    "level_2": {"pass_threshold": 0.75, "hint_enabled": false, "time_limit_sec": null, "immediate_feedback": true},
    "level_3": {"pass_threshold": 0.75, "hint_enabled": false, "time_limit_sec": 45, "immediate_feedback": false}
}');

INSERT INTO app_config (config_key, config_value) VALUES
('gamification_rules', '{
    "hearts_enabled": true, "hearts_total": 5,
    "heart_penalty_level2": 1, "heart_penalty_level3": 2,
    "xp_per_lesson": 10, "xp_per_qcm_level1": 20, "xp_per_qcm_level2": 30, "xp_per_qcm_level3": 50,
    "streak_enabled": true, "leaderboard_weekly": true
}');

INSERT INTO app_config (config_key, config_value) VALUES
('exam_config', '{
    "ppl.air_law": {"questions_count": 16, "time_minutes": 30, "passing_score_percent": 75, "randomized": true, "no_backward": true}
}');

-- ============================================================================
-- 2. LESSONS
-- ============================================================================
INSERT INTO lessons (id, license, category, theme, title_fr, title_en, content_fr, content_en, difficulty, order_index, level, duration_minutes, tags, learning_objectives) VALUES
('621491b3-0000-4000-8000-0000621491b3', 'PPL', 'air law', 'air law', 'Organisation aviation internationale (ICAO, OACI)', 'International Aviation Organization (ICAO)', '## Qu''est-ce que l''ICAO ?

**ICAO** = International Civil Aviation Organization (OACI en français)

- **Créée en 1944** par la **Convention de Chicago**
- **193 pays membres** (presque tous les pays du monde)
- **Rôle principal** : fixer les **STANDARDS MINIMAUX internationaux** (Annexes)
- **Siège** : Montréal, Canada
- **Nature** : Agence spécialisée de l''ONU
- **Différence clé avec FAA** : ICAO = recommandations internationales, FAA = réglementation **obligatoire** aux USA

Les pays membres doivent implémenter les standards ICAO ou notifier les ''différences'' officiellement.

## Les 19 Annexes ICAO

Les **19 Annexes ICAO** couvrent tous les aspects de l''aviation civile :

1. **Personnel Licensing** (licences pilotes PPL/CPL/ATPL)
2. **Rules of the Air** (règles VFR/IFR)
3. **Meteorological Service** (météo)
4. **Charts and Navigation** (cartes)
5. **Airspace Classification**
6. **Operation of Aircraft**
7. **Aircraft Nationality** (immatriculation)
8. **Airworthiness** (navigabilité)
9. **Facilitation** (douane, immigration)
10. **Communication Procedures**
11. **Air Traffic Services** (ATC)
12. **Search and Rescue** (SAR)
13. **Accident Investigation**
14. **Aerodromes** (aéroports)
15. **Aeronautical Information Service** (AIS)
16. **Environmental Protection** (bruit, CO2)
17. **Security** (sûreté)
18. **Transport of Dangerous Goods**
19. **Safety Management** (SMS)

Chaque Annexe contient des **STANDARDS (S)** (obligatoires) et **RECOMMANDATIONS (P)** (optionnelles).

## Standards vs Recommandations + Notification de différence

**Différence cruciale** dans les Annexes ICAO :

- **STANDARD (S)** : Spécification **OBLIGATOIRE**. Les pays DOIVENT l''implémenter ou notifier une ''différence'' officielle.
  *Exemple : Annexe 1 - âge minimum PPL = 17 ans*

- **RECOMMANDATION (P)** : Spécification **OPTIONNELLE**. Les pays SONT ENCOURAGÉS à l''implémenter.
  *Exemple : Annexe 3 - couleur des balises météo*

**''Notification de différence''** : Si un pays ne peut pas suivre un Standard, il doit informer l''ICAO officiellement. Cette différence est publiée dans l''**AIP** du pays.

**Pour un pilote PPL** : Toujours vérifier l''AIP du pays où tu vols !

', '## What is ICAO?

**ICAO** = International Civil Aviation Organization

- **Created in 1944** by the **Chicago Convention**
- **193 member countries** (almost all countries worldwide)
- **Main role** : set **INTERNATIONAL MINIMUM STANDARDS** (Annexes)
- **Headquarters** : Montreal, Canada
- **Nature** : UN specialized agency
- **Key difference with FAA** : ICAO = international recommendations, FAA = **mandatory** regulations in USA

Member countries must implement ICAO standards or officially notify ''differences''.

## The 19 ICAO Annexes

The **19 ICAO Annexes** cover all aspects of civil aviation :

1. **Personnel Licensing** (PPL/CPL/ATPL pilot licenses)
2. **Rules of the Air** (VFR/IFR rules)
3. **Meteorological Service** (weather)
4. **Charts and Navigation**
5. **Airspace Classification**
6. **Operation of Aircraft**
7. **Aircraft Nationality** (registration)
8. **Airworthiness** (certification)
9. **Facilitation** (customs, immigration)
10. **Communication Procedures**
11. **Air Traffic Services** (ATC)
12. **Search and Rescue** (SAR)
13. **Accident Investigation**
14. **Aerodromes** (airports)
15. **Aeronautical Information Service** (AIS)
16. **Environmental Protection** (noise, CO2)
17. **Security** (safeguarding)
18. **Transport of Dangerous Goods**
19. **Safety Management** (SMS)

Each Annexe contains **STANDARDS (S)** (mandatory) and **RECOMMENDED PRACTICES (P)** (optional).

## Standards vs Recommended Practices + Difference notification

**Crucial difference** in ICAO Annexes :

- **STANDARD (S)** : **MANDATORY** specification. Countries MUST implement it or officially notify a ''difference''.
  *Example: Annexe 1 - minimum age for PPL = 17 years*

- **RECOMMENDED PRACTICE (P)** : **OPTIONAL** specification. Countries are ENCOURAGED to implement.
  *Example: Annexe 3 - color of weather beacons*

**''Difference notification''** : If a country cannot follow a Standard, it must inform ICAO officially. This difference is published in the country''s **AIP**.

**For PPL pilots** : Always check the AIP of the country where you fly!

', 1, 1, 1, 25, ["ICAO","OACI","international","organization","annexes","basics","PPL"], ["Comprendre l'ICAO, sa création, ses 19 Annexes, et son rôle dans l'aviation civile mondiale pour le PPL."]),

-- Questions for lesson 621491b3-0000-4000-8000-0000621491b3 (15 questions)
INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, answer_key, explanation_fr, explanation_en, faa_note_fr, faa_note_en, tags, difficulty_score) VALUES
('180737d9-0000-4000-8000-0000180737d9','621491b3-0000-4000-8000-0000621491b3','PPL','air law','air law','Qu''est-ce que l''ICAO ?',1,1,'Quand l''ICAO a-t-elle été créée ?','When was ICAO created?','["1939","1944","1950","1960"]','1944','Bonne réponse ! **1944** = Année de la Convention de Chicago qui a créé l''ICAO.','Correct! **1944** = Year of the Chicago Convention that created ICAO.','','USA has been an ICAO member since 1944 and follows ICAO standards, but FAA imposes additional mandatory rules.',["ICAO","history","Chicago Convention","basic"],0.25),
('180737da-0000-4000-8000-0000180737da','621491b3-0000-4000-8000-0000621491b3','PPL','air law','air law','Qu''est-ce que l''ICAO ?',1,1,'Combien de pays sont membres de l''ICAO ?','How many countries are ICAO members?','["150","175","193","200"]','193','Bonne réponse ! L''ICAO compte **193 pays membres**, couvrant presque tous les pays du monde.','Correct! ICAO has **193 member countries**, covering almost all countries worldwide.','','',["ICAO","membership","basic"],0.20),
('180737db-0000-4000-8000-0000180737db','621491b3-0000-4000-8000-0000621491b3','PPL','air law','air law','Qu''est-ce que l''ICAO ?',1,1,'Quel est le rôle principal de l''ICAO ?','What is ICAO''s main role?','["Imposer des règles obligatoires à tous les pays","Fixer des standards minimaux internationaux","Gérer le contrôle aérien mondial","Former les pilotes partout dans le monde"]','Fixer des standards minimaux internationaux','Bonne réponse ! L''ICAO fixe des **STANDARDS MINIMAUX internationaux**.','Correct! ICAO sets **INTERNATIONAL MINIMUM STANDARDS**.','','Unlike ICAO, FAA imposes MANDATORY rules on US territory.',["ICAO","role","standards","basic"],0.35),
('180737dc-0000-4000-8000-0000180737dc','621491b3-0000-4000-8000-0000621491b3','PPL','air law','air law','Qu''est-ce que l''ICAO ?',1,1,'Où se trouve le siège de l''ICAO ?','Where is ICAO headquarters located?','["Genève, Suisse","Paris, France","Montréal, Canada","Washington, USA"]','Montréal, Canada','Bonne réponse ! Le siège de l''ICAO est à **Montréal, Canada**.','Correct! ICAO headquarters is in **Montreal, Canada**.','','',["ICAO","headquarters","basic"],0.30),
('180737dd-0000-4000-8000-0000180737dd','621491b3-0000-4000-8000-0000621491b3','PPL','air law','air law','Qu''est-ce que l''ICAO ?',1,1,'Quelle est l''organisation ONU dont l''ICAO fait partie ?','Which UN organization is ICAO part of?','["OMS","UNESCO","Agence spécialisée de l''ONU","OTAN"]','Agence spécialisée de l''ONU','Bonne réponse ! L''ICAO est une **agence spécialisée de l''ONU** depuis sa création.','Correct! ICAO is a **UN specialized agency** since its creation.','','',["ICAO","UN","organization","basic"],0.40),
('18073b9f-0000-4000-8000-000018073b9f','621491b3-0000-4000-8000-0000621491b3','PPL','air law','air law','Les 19 Annexes ICAO',1,1,'Combien d''Annexes ICAO existent actuellement ?','How many ICAO Annexes exist currently?','["15","17","19","21"]','19','Bonne réponse ! Il y a **19 Annexes ICAO** couvrant tous les aspects de l''aviation civile.','Correct! There are **19 ICAO Annexes** covering all aspects of civil aviation.','','',["ICAO","annexes","basic"],0.25),
('18073ba0-0000-4000-8000-000018073ba0','621491b3-0000-4000-8000-0000621491b3','PPL','air law','air law','Les 19 Annexes ICAO',1,1,'Quelle Annexe ICAO concerne les licences pilotes (PPL, CPL, ATPL) ?','Which ICAO Annexe covers pilot licenses (PPL, CPL, ATPL)?','["Annexe 1","Annexe 5","Annexe 10","Annexe 14"]','Annexe 1','**Annexe 1 = Personnel Licensing**. Elle définit les standards pour toutes les licences pilotes.','**Annexe 1 = Personnel Licensing**. It sets standards for all pilot licenses.','','FAA uses Part 61 and Part 141 for licenses, based on ICAO Annexe 1.',["ICAO","annexes","licensing","PPL","basic"],0.30),
('18073ba1-0000-4000-8000-000018073ba1','621491b3-0000-4000-8000-0000621491b3','PPL','air law','air law','Les 19 Annexes ICAO',2,2,'Quelle Annexe ICAO concerne les règles de l''air (VFR/IFR) ?','Which ICAO Annexe covers rules of the air (VFR/IFR)?','["Annexe 1","Annexe 2","Annexe 6","Annexe 11"]','Annexe 2','**Annexe 2 = Rules of the Air**. Elle définit les règles VFR et IFR.','**Annexe 2 = Rules of the Air**. It defines VFR and IFR rules.','','',["ICAO","annexes","rules_of_air","intermediate"],0.40),
('18073ba2-0000-4000-8000-000018073ba2','621491b3-0000-4000-8000-0000621491b3','PPL','air law','air law','Les 19 Annexes ICAO',2,2,'Quelle Annexe ICAO traite de la navigabilité des aéronefs ?','Which ICAO Annexe deals with aircraft airworthiness?','["Annexe 6","Annexe 7","Annexe 8","Annexe 18"]','Annexe 8','**Annexe 8 = Airworthiness of Aircraft**. Elle définit les standards de navigabilité.','**Annexe 8 = Airworthiness of Aircraft**. It sets airworthiness standards.','','FAA uses ''Type Certificate'' and ''Airworthiness Certificate'' per FAR Part 21.',["ICAO","annexes","airworthiness","intermediate"],0.45),
('18073ba3-0000-4000-8000-000018073ba3','621491b3-0000-4000-8000-0000621491b3','PPL','air law','air law','Les 19 Annexes ICAO',2,2,'Quelle Annexe ICAO concerne le contrôle aérien (ATC) ?','Which ICAO Annexe covers air traffic control (ATC)?','["Annexe 10","Annexe 11","Annexe 14","Annexe 15"]','Annexe 11','**Annexe 11 = Air Traffic Services**. Elle définit les services ATC.','**Annexe 11 = Air Traffic Services**. It defines ATC services.','','',["ICAO","annexes","ATC","intermediate"],0.40),
('68e0ac54-0000-4000-8000-000068e0ac54','621491b3-0000-4000-8000-0000621491b3','PPL','air law','air law','Standards vs Recommandations',2,2,'Quelle est la différence entre un STANDARD (S) et une RECOMMANDATION (P) ICAO ?','What is the difference between a STANDARD (S) and a RECOMMENDED PRACTICE (P) in ICAO?','["S est optionnel, P est obligatoire","S est obligatoire, P est optionnel","S et P sont tous les deux obligatoires","S et P sont tous les deux optionnels"]','S est obligatoire, P est optionnel','**STANDARD (S)** = obligatoire. **RECOMMANDATION (P)** = optionnel (encouragé).','**STANDARD (S)** = mandatory. **RECOMMENDATION (P)** = optional (encouraged).','','FAA doesn''t use this system. All FARs are mandatory.',["ICAO","standards","recommendations","intermediate"],0.50),
('68e0ac55-0000-4000-8000-000068e0ac55','621491b3-0000-4000-8000-0000621491b3','PPL','air law','air law','Standards vs Recommandations',2,2,'Si un pays ne peut pas suivre un Standard ICAO, que doit-il faire ?','If a country cannot follow an ICAO Standard, what must it do?','["Rien, c''est optionnel","Notifier officiellement une ''différence'' à l''ICAO","Quitter l''ICAO","Changer le Standard lui-même"]','Notifier officiellement une ''différence'' à l''ICAO','Le pays doit **notifier officiellement une ''différence''** à l''ICAO. Publiée dans l''AIP.','The country must **officially notify a ''difference''** to ICAO. Published in the AIP.','','',["ICAO","differences","AIP","intermediate"],0.55),
('68e0ac56-0000-4000-8000-000068e0ac56','621491b3-0000-4000-8000-0000621491b3','PPL','air law','air law','Standards vs Recommandations',2,2,'Où un pilote PPL peut-il trouver les ''différences'' ICAO d''un pays ?','Where can a PPL pilot find a country''s ICAO ''differences''?','["Dans le manuel de vol","Dans l''AIP du pays","Sur le site FAA","Dans la carte VAC"]','Dans l''AIP du pays','Les différences sont publiées dans l''**AIP (Aeronautical Information Publication)** du pays.','Differences are published in the country''s **AIP**.','','',["ICAO","AIP","differences","intermediate"],0.50),
('68e0ac57-0000-4000-8000-000068e0ac57','621491b3-0000-4000-8000-0000621491b3','PPL','air law','air law','Standards vs Recommandations',1,1,'Quel est l''âge minimum pour obtenir un PPL selon l''Annexe 1 ICAO ?','What is the minimum age for a PPL according to ICAO Annexe 1?','["15 ans","17 ans","18 ans","21 ans"]','17 ans','**Annexe 1 : âge minimum PPL = 17 ans** (Standard ICAO).','**Annexe 1: minimum age for PPL = 17 years** (ICAO Standard).','','FAA Part 61: also 17 years for PPL in USA.',["ICAO","Annexe 1","PPL","age","basic"],0.30),
('68e0ac58-0000-4000-8000-000068e0ac58','621491b3-0000-4000-8000-0000621491b3','PPL','air law','air law','Standards vs Recommandations',2,2,'Quelle Annexe ICAO concerne la sécurité (Safety Management System - SMS) ?','Which ICAO Annexe covers Safety Management System (SMS)?','["Annexe 6","Annexe 17","Annexe 18","Annexe 19"]','Annexe 19','**Annexe 19 = Safety Management**. SMS obligatoire pour les airlines depuis 2018.','**Annexe 19 = Safety Management**. SMS is mandatory for airlines since 2018.','','',["ICAO","annexes","SMS","safety","intermediate"],0.60),

INSERT INTO lessons (id, license, category, theme, title_fr, title_en, content_fr, content_en, difficulty, order_index, level, duration_minutes, tags, learning_objectives) VALUES
('621491b4-0000-4000-8000-0000621491b4', 'PPL', 'air law', 'air law', 'Réglementation : EASA vs FAA (Europe vs USA)', 'Regulations: EASA vs FAA (Europe vs USA)', '## Qu''est-ce que l''EASA ?

**EASA** = European Union Aviation Safety Agency

- **Créée en 2002** (siège à **Cologne, Allemagne**)
- **Rôle** : fixer les normes de sécurité pour **tous les pays de l''UE** (27 membres + associés)
- **Règlements** : **Part-FCL** (licences), **Part-NCO** (opérations non-commerciales), **Part-M** (maintenance)
- **Avantages** : Licence EASA reconnue dans **toute l''UE** sans conversion
- **Différence clé avec FAA** : EASA = réglementation **obligatoire** pour l''UE, FAA = réglementation **obligatoire** pour les USA seulement

**Part-FCL** (Flight Crew Licensing) :
- PPL(A) : Private Pilot License (Avion)
- CPL(A) : Commercial Pilot License
- ATPL(A) : Airline Transport Pilot License
- IR(A) : Instrument Rating
- MEL : Multi-Engine Ratings

## Qu''est-ce que la FAA ?

**FAA** = Federal Aviation Administration

- **Créée en 1958** (siège à **Washington, DC, USA**)
- **Rôle** : réglementation **obligatoire** pour **tous les vols aux USA** (y compris vols étrangers)
- **Règlements** : **FAR** (Federal Aviation Regulations) divisés en **Parts**
  - **Part 61** : Licences (flexible, parcours individuel)
  - **Part 141** : Écoles de pilotage approuvées (structuré, plus rapide)
  - **Part 91** : Règles de l''air (général)
  - **Part 135** : Opérations aériennes commerciales (taxi, charter)
- **Avantages** : Formation souvent **moins chère** et **plus rapide** aux USA
- **Différence clé** : FAA n''est reconnue **QUE aux USA** (conversion nécessaire pour voler en Europe)

**Licences FAA** :
- Private Pilot Certificate (équivalent PPL)
- Commercial Pilot Certificate
- Airline Transport Pilot Certificate (ATP)
- Instrument Rating
- Multi-Engine Rating

## Comparatif EASA vs FAA + Conversion

**Tableau comparatif clé** :

| Critère | EASA (Europe) | FAA (USA) |
|---|---|---|
| **Âge minimum PPL** | 17 ans | 17 ans |
| **Heures vol minimum** | 45h (20h solo) | 40h (20h solo) |
| **Examen théorique** | 9 matières (Air Law, Meteo, etc.) | 60 questions (Private Pilot Airplane) |
| **Validité licence** | Illimitée (si medical valide + 12 mois recentrage) | Illimitée (si medical valide + 24 mois recentrage) |
| **Medical** | Classe 2 (PPL) | Classe 3 (Private) |
| **Reconnaissance** | Toute l''UE sans conversion | Seulement USA (conversion requise ailleurs) |

**Conversion EASA → FAA** (PPL) :
1. Avoir licence EASA valide + medical valide
2. Passer **examen théorique FAA Private Pilot** (60 questions)
3. Passer **examen pratique (checkride)** avec FAA DPE
4. Obtenir **Private Pilot Certificate FAA**
5. **Pas besoin de refaire 40h vol** (exigence de vol déjà satisfaite)

**Conversion FAA → EASA** (PPL) :
1. Avoir licence FAA valide + medical valide
2. Passer **examen théorique EASA PPL** (9 matieres)
3. Passer **examen pratique** avec EASA examiner
4. Possiblement **10h vol supplementaires** (reglementation europeenne plus stricte)
5. Obtenir **PPL EASA**

**Attention** : Les conversions prennent **2-6 mois** et coutent **500-2000€** selon le pays.

', '## What is EASA?

**EASA** = European Union Aviation Safety Agency

- **Created in 2002** (headquarters in **Cologne, Germany**)
- **Role** : set safety standards for **all EU countries** (27 members + associates)
- **Regulations** : **Part-FCL** (licenses), **Part-NCO** (non-commercial ops), **Part-M** (maintenance)
- **Advantage** : EASA license recognized **throughout EU** without conversion
- **Key difference with FAA** : EASA = mandatory for EU, FAA = mandatory for USA only

**Part-FCL** (Flight Crew Licensing) :
- PPL(A) : Private Pilot License (Airplane)
- CPL(A) : Commercial Pilot License
- ATPL(A) : Airline Transport Pilot License
- IR(A) : Instrument Rating
- MEL : Multi-Engine Ratings

## What is FAA?

**FAA** = Federal Aviation Administration

- **Created in 1958** (headquarters in **Washington, DC, USA**)
- **Role** : **mandatory** regulation for **all flights in USA** (including foreign flights)
- **Regulations** : **FAR** (Federal Aviation Regulations) divided into **Parts**
  - **Part 61** : Licenses (flexible, individual path)
  - **Part 141** : Approved flight schools (structured, faster)
  - **Part 91** : Flight rules (general)
  - **Part 135** : Commercial air operations (taxi, charter)
- **Advantage** : Training often **cheaper** and **faster** in USA
- **Key difference** : FAA recognized **ONLY in USA** (conversion needed for Europe)

**FAA Licenses** :
- Private Pilot Certificate (equivalent to PPL)
- Commercial Pilot Certificate
- Airline Transport Pilot Certificate (ATP)
- Instrument Rating
- Multi-Engine Rating

## EASA vs FAA Comparison + Conversion

**Key comparison table** :

| Criterion | EASA (Europe) | FAA (USA) |
|---|---|---|
| **Minimum age PPL** | 17 years | 17 years |
| **Minimum flight hours** | 45h (20h solo) | 40h (20h solo) |
| **Written exam** | 9 subjects (Air Law, Met, etc.) | 60 questions (Private Pilot Airplane) |
| **License validity** | Unlimited (if medical valid + 12 months recency) | Unlimited (if medical valid + 24 months recency) |
| **Medical** | Class 2 (PPL) | Class 3 (Private) |
| **Recognition** | All EU without conversion | Only USA (conversion required elsewhere) |

**EASA → FAA conversion** (PPL) :
1. Have valid EASA license + valid medical
2. Pass **FAA Private Pilot written exam** (60 questions)
3. Pass **checkride** with FAA DPE
4. Obtain **Private Pilot Certificate FAA**
5. **No need to redo 40h flight** (flight requirement already met)

**FAA → EASA conversion** (PPL) :
1. Have valid FAA license + valid medical
2. Pass **EASA PPL written exam** (9 subjects)
3. Pass **checkride** with EASA examiner
4. Possibly **10 additional flight hours** (European rules stricter)
5. Obtain **PPL EASA**

**Warning** : Conversions take **2-6 months** and cost **500-2000€** depending on country.

', 1, 2, 1, 30, ["EASA","FAA","regulation","Europe","USA","licensing","conversion","PPL"], ["Comprendre les différences clés entre EASA (Europe) et FAA (USA), les structures de licences, et les procédures de conversion pour un pilote PPL."]),

-- Questions for lesson 621491b4-0000-4000-8000-0000621491b4 (15 questions)
INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, answer_key, explanation_fr, explanation_en, faa_note_fr, faa_note_en, tags, difficulty_score) VALUES
('53732db9-0000-4000-8000-000053732db9','621491b4-0000-4000-8000-0000621491b4','PPL','air law','air law','Qu''est-ce que l''EASA ?',1,1,'Où se trouve le siège de l''EASA ?','Where is EASA headquarters located?','["Paris, France","Cologne, Allemagne","Montreal, Canada","Washington, USA"]','Cologne, Allemagne','Bonne reponse ! Le siege de l''EASA est a **Cologne, Allemagne** (depuis 2002).','Correct! EASA headquarters is in **Cologne, Germany** (since 2002).','','',["EASA","headquarters","Europe","basic"],0.25),
('53732dba-0000-4000-8000-000053732dba','621491b4-0000-4000-8000-0000621491b4','PPL','air law','air law','Qu''est-ce que l''EASA ?',1,1,'Quand l''EASA a-t-elle ete creee ?','When was EASA created?','["1944","1958","2002","2010"]','2002','Bonne reponse ! L''EASA a ete creee en **2002** par l''UE.','Correct! EASA was created in **2002** by the EU.','','',["EASA","history","basic"],0.30),
('53732dbb-0000-4000-8000-000053732dbb','621491b4-0000-4000-8000-0000621491b4','PPL','air law','air law','Qu''est-ce que l''EASA ?',1,1,'Quel reglement EASA regit les licences pilotes ?','Which EASA regulation governs pilot licenses?','["Part-M","Part-NCO","Part-FCL","Part-145"]','Part-FCL','Bonne reponse ! **Part-FCL** = Flight Crew Licensing (licences pilotes).','Correct! **Part-FCL** = Flight Crew Licensing (pilot licenses).','','FAA equivalent is Part 61 and Part 141.',["EASA","Part-FCL","licensing","basic"],0.35),
('53732dbc-0000-4000-8000-000053732dbc','621491b4-0000-4000-8000-0000621491b4','PPL','air law','air law','Qu''est-ce que l''EASA ?',1,1,'Dans combien de pays la licence EASA est-elle reconnue sans conversion ?','In how many countries is EASA license recognized without conversion?','["27 (UE)","193","5","1"]','27 (UE)','Bonne reponse ! Licence EASA valable dans **tous les 27 pays de l''UE** sans conversion.','Correct! EASA license valid in **all 27 EU countries** without conversion.','','',["EASA","EU","recognition","basic"],0.30),
('5373317e-0000-4000-8000-00005373317e','621491b4-0000-4000-8000-0000621491b4','PPL','air law','air law','Qu''est-ce que la FAA ?',1,1,'Où se trouve le siege de la FAA ?','Where is FAA headquarters located?','["Cologne, Allemagne","Montreal, Canada","Washington, DC, USA","Geneve, Suisse"]','Washington, DC, USA','Bonne reponse ! Siege FAA a **Washington, DC, USA**.','Correct! FAA headquarters in **Washington, DC, USA**.','','',["FAA","headquarters","USA","basic"],0.25),
('5373317f-0000-4000-8000-00005373317f','621491b4-0000-4000-8000-0000621491b4','PPL','air law','air law','Qu''est-ce que la FAA ?',1,1,'Quand la FAA a-t-elle ete creee ?','When was FAA created?','["1944","1958","2002","1970"]','1958','Bonne reponse ! FAA creee en **1958** par le Congres americain.','Correct! FAA created in **1958** by US Congress.','','',["FAA","history","basic"],0.30),
('53733180-0000-4000-8000-000053733180','621491b4-0000-4000-8000-0000621491b4','PPL','air law','air law','Qu''est-ce que la FAA ?',1,1,'Quel reglement FAA regit les licences pilotes (formation individuelle) ?','Which FAA regulation governs pilot licenses (individual training)?','["Part 91","Part 135","Part 61","Part 141"]','Part 61','Bonne reponse ! **Part 61** = formation flexible, individuelle.','Correct! **Part 61** = flexible, individual training.','','Part 141 is for approved flight schools (structured).',["FAA","Part 61","licensing","basic"],0.35),
('53733181-0000-4000-8000-000053733181','621491b4-0000-4000-8000-0000621491b4','PPL','air law','air law','Qu''est-ce que la FAA ?',1,1,'Quel reglement FAA pour les ecoles de pilotage approuvees (formation structuree) ?','Which FAA regulation for approved flight schools (structured training)?','["Part 61","Part 91","Part 141","Part 135"]','Part 141','Bonne reponse ! **Part 141** = ecoles approuvees, formation structuree (plus rapide).','Correct! **Part 141** = approved schools, structured training (faster).','','',["FAA","Part 141","flight schools","basic"],0.40),
('53733182-0000-4000-8000-000053733182','621491b4-0000-4000-8000-0000621491b4','PPL','air law','air law','Qu''est-ce que la FAA ?',1,1,'Quelle est l''heure de vol minimum pour un PPL FAA ?','What is the minimum flight hours for FAA PPL?','["35h","40h","45h","50h"]','40h','Bonne reponse ! **FAA Part 61 : 40h minimum** (dont 20h solo).','Correct! **FAA Part 61: 40h minimum** (including 20h solo).','','EASA requires 45h minimum.',["FAA","hours","PPL","basic"],0.30),
('53733544-0000-4000-8000-000053733544','621491b4-0000-4000-8000-0000621491b4','PPL','air law','air law','Comparatif EASA vs FAA + Conversion',2,2,'Quelle est la difference d''heures de vol minimum entre EASA et FAA pour PPL ?','What is the flight hour difference between EASA and FAA for PPL?','["EASA 40h, FAA 45h","EASA 45h, FAA 40h","EASA 35h, FAA 40h","EASA 50h, FAA 45h"]','EASA 45h, FAA 40h','Bonne reponse ! **EASA : 45h**, **FAA : 40h**. EASA exige 5h de plus.','Correct! **EASA: 45h**, **FAA: 40h**. EASA requires 5h more.','','EASA also requires more solo hours (20h EASA vs 10h FAA solo cross-country specific).',["EASA","FAA","hours","comparison","intermediate"],0.45),
('1af37274-0000-4000-8000-00001af37274','621491b4-0000-4000-8000-0000621491b4','PPL','air law','air law','Comparatif EASA vs FAA + Conversion',2,2,'Pour convertir un PPL EASA vers FAA, faut-il refaire les 40h de vol ?','To convert EASA PPL to FAA, must you redo 40 flight hours?','["Oui, obligatoirement","Non, pas besoin (deja satisfait)","Seulement 10h supplementaires","Seulement si vol aux USA"]','Non, pas besoin (deja satisfait)','Bonne reponse ! **Pas besoin de refaire 40h**. L''exigence de vol est deja satisfaite par la licence EASA.','Correct! **No need to redo 40h**. Flight requirement already met by EASA license.','','',["conversion","EASA","FAA","hours","intermediate"],0.50),
('1af37275-0000-4000-8000-00001af37275','621491b4-0000-4000-8000-0000621491b4','PPL','air law','air law','Comparatif EASA vs FAA + Conversion',2,2,'Combien de temps prend generalement une conversion EASA vers FAA ?','How long does EASA to FAA conversion typically take?','["1 semaine","2-6 mois","1 an","3 ans"]','2-6 mois','Bonne reponse ! Conversion prend **2-6 mois** (examen theorie + checkride + administration).','Correct! Conversion takes **2-6 months** (written exam + checkride + paperwork).','','',["conversion","EASA","FAA","time","intermediate"],0.40),
('1af37276-0000-4000-8000-00001af37276','621491b4-0000-4000-8000-0000621491b4','PPL','air law','air law','Comparatif EASA vs FAA + Conversion',1,1,'Quel medical est requis pour PPL EASA ?','Which medical is required for EASA PPL?','["Class 1","Class 2","Class 3","Basic Medical"]','Class 2','Bonne reponse ! **EASA PPL : Class 2 medical**.','Correct! **EASA PPL: Class 2 medical**.','','FAA PPL requires Class 3 medical.',["EASA","medical","PPL","basic"],0.35),
('1af37277-0000-4000-8000-00001af37277','621491b4-0000-4000-8000-0000621491b4','PPL','air law','air law','Comparatif EASA vs FAA + Conversion',1,1,'Quel medical est requis pour Private Pilot FAA ?','Which medical is required for FAA Private Pilot?','["Class 1","Class 2","Class 3","Basic Medical"]','Class 3','Bonne reponse ! **FAA Private Pilot : Class 3 medical**.','Correct! **FAA Private Pilot: Class 3 medical**.','','',["FAA","medical","private","basic"],0.35),
('1af37278-0000-4000-8000-00001af37278','621491b4-0000-4000-8000-0000621491b4','PPL','air law','air law','Comparatif EASA vs FAA + Conversion',2,2,'Une licence FAA est-elle reconnue directement en Europe sans conversion ?','Is FAA license recognized directly in Europe without conversion?','["Oui, dans toute l''UE","Oui, seulement en France","Non, conversion requise","Oui, seulement si vol VFR"]','Non, conversion requise','Bonne reponse ! **FAA reconnue UNIQUEMENT aux USA**. Conversion EASA requise pour voler en Europe.','Correct! **FAA recognized ONLY in USA**. EASA conversion required for Europe.','','',["FAA","recognition","Europe","conversion","intermediate"],0.40),

INSERT INTO lessons (id, license, category, theme, title_fr, title_en, content_fr, content_en, difficulty, order_index, level, duration_minutes, tags, learning_objectives) VALUES
('621491b5-0000-4000-8000-0000621491b5', 'PPL', 'air law', 'air law', 'Certificats d''aéronefs (Type, Navigabilité, Immatriculation)', 'Aircraft Certificates (Type, Airworthiness, Registration)', '## Certificat de Type (Type Certificate)

**Certificat de Type** = Document qui certifie que la **conception** d''un aéronef (avion, moteur, hélice) respecte les **normes de sécurité** de l''autorité aéronautique.

- **Délivré par** : EASA (Europe) ou FAA (USA) ou autorité nationale (DGAC France, AESA Espagne)
- **Couvre** : Conception du modèle (ex: Cessna 172S, Diamond DA40)
- **Validité** : **Illimitée** (tant que le modèle reste conforme)
- **Qui le détient** : Le **constructeur** (ex: Cessna, Piper, Diamond)
- **Lien avec navigabilité** : Un avion ne peut obtenir un certificat de navigabilité sans certificat de type valide

**Parties du certificat de type** :
1. **Data Sheet** : Spécifications techniques du modèle
2. **Condition de limitation** : Restrictions d''exploitation
3. **Manuels requis** : Manuel de vol, manuel de maintenance

**Pour un pilote PPL** : Tu n''as pas besoin de vérifier le certificat de type (c''est au constructeur), mais tu dois savoir qu''il existe et qu''il est la base de la navigabilité.

## Certificat de Navigabilité (Airworthiness Certificate)

**Certificat de Navigabilité** = Document qui certifie que **cet aéronef spécifique** (numéro de série) est **en bon état** et conforme au certificat de type.

- **Délivré par** : Autorité nationale (DGAC France, AESA Espagne) après inspection
- **Couvre** : L''aéronef individuel (ex: F-HXYZ, numéro de série 12345)
- **Validité** : **Illimitée** SOUS CONDITIONS (maintenance régulière, modifications approuvées)
- **Qui le détient** : Le **propriétaire/exploitant** (pas le constructeur)
- **Où il est affiché** : Dans l''avion, généralement dans le cockpit (près du siège pilote) ou dans le manuel de vol

**Conditions de validité** :
1. Maintenance conforme au **programme de maintenance** (annuel, 100h, etc.)
2. Pas de modifications non approuvées
3. Toutes les **Airworthiness Directives (AD)** appliquées
4. Inspections périodiques effectuées

**Documents requis à bord** (pour un vol PPL) :
1. **Certificat de navigabilité original** (ou copie certifiée)
2. **Certificat d''immatriculation original**
3. **Licence de station d''aéronef** (si radio embarquée)
4. **Manuel de vol** (Approuvé)
5. **Journal de bord / Logbook** de l''aéronef

**Attention** : Si le certificat de navigabilité est périmé ou suspendu, **l''avion ne doit pas voler**.

## Certificat d''Immatriculation (Registration Certificate)

**Certificat d''Immatriculation** = Document qui atteste de l''**enregistrement national** de l''aéronef et de son **propriétaire**.

- **Délivré par** : Autorité nationale (DGAC France, AESA Espagne, CAA UK)
- **Couvre** : Propriété de l''aéronef (qui le possède)
- **Immatriculation** : Préfixe national + suffixe unique (ex: **F-HXYZ** pour France, **N12345** pour USA, **G-ABCD** pour UK)
- **Validité** : **3 ans** (France/Europe) → renouvellement obligatoire
- **Qui le détient** : Le **propriétaire** (personne physique ou juridique)

**Préfixes nationaux courants** :
| Pays | Préfixe | Exemple |
|---|---|---|
| France | **F-** | F-HXYZ |
| USA | **N-** | N12345 |
| Royaume-Uni | **G-** | G-ABCD |
| Allemagne | **D-** | D-EFGH |
| Espagne | **EC-** | EC-IJKL |
| Canada | **C-** | C-MNOP |

**Documents requis à bord** :
1. **Certificat d''immatriculation original** (valide)
2. Si l''avion a une radio : **Licence de station d''aéronef** (Valide 5 ans en Europe)

**Ce qui se passe si l''immatriculation expire** :
- **L''avion ne peut pas voler** jusqu''au renouvellement
- **Amende possible** si vol avec immatriculation périmée
- **Assurance non valable** en cas d''accident

**Pour un pilote PPL** : Toujours vérifier l''immatriculation avant de voler (dans le cockpit ou sur le certificat).

', '## Type Certificate

**Type Certificate** = Document certifying that the **design** of an aircraft (airplane, engine, propeller) meets **safety standards** of the aviation authority.

- **Issued by** : EASA (Europe) or FAA (USA) or national authority (DGAC France, AESA Spain)
- **Covers** : Design of the model (e.g., Cessna 172S, Diamond DA40)
- **Validity** : **Unlimited** (as long as the model remains compliant)
- **Holder** : The **manufacturer** (e.g., Cessna, Piper, Diamond)
- **Link to airworthiness** : An aircraft cannot obtain an airworthiness certificate without a valid type certificate

**Parts of type certificate** :
1. **Data Sheet** : Technical specifications of the model
2. **Limitation conditions** : Operational restrictions
3. **Required manuals** : Flight manual, maintenance manual

**For PPL pilots** : You don''t need to check the type certificate (it''s the manufacturer''s responsibility), but you must know it exists and is the basis of airworthiness.

## Airworthiness Certificate

**Airworthiness Certificate** = Document certifying that **this specific aircraft** (serial number) is **in good condition** and compliant with the type certificate.

- **Issued by** : National authority (DGAC France, AESA Spain) after inspection
- **Covers** : Individual aircraft (e.g., F-HXYZ, serial number 12345)
- **Validity** : **Unlimited** UNDER CONDITIONS (regular maintenance, approved modifications)
- **Holder** : The **owner/operator** (not the manufacturer)
- **Where displayed** : In the aircraft, usually in the cockpit (near pilot seat) or in the flight manual

**Validity conditions** :
1. Maintenance compliant with **maintenance program** (annual, 100h, etc.)
2. No unapproved modifications
3. All **Airworthiness Directives (AD)** applied
4. Periodic inspections completed

**Required onboard documents** (for a PPL flight) :
1. **Original Airworthiness Certificate** (or certified copy)
2. **Original Registration Certificate**
3. **Aircraft Station License** (if radio onboard)
4. **Approved Flight Manual**
5. **Aircraft Logbook**

**Warning** : If the airworthiness certificate is expired or suspended, **the aircraft must not fly**.

## Registration Certificate

**Registration Certificate** = Document attesting to the **national registration** of the aircraft and its **owner**.

- **Issued by** : National authority (DGAC France, AESA Spain, CAA UK)
- **Covers** : Ownership of the aircraft (who owns it)
- **Registration** : National prefix + unique suffix (e.g., **F-HXYZ** for France, **N12345** for USA, **G-ABCD** for UK)
- **Validity** : **3 years** (France/Europe) → mandatory renewal
- **Holder** : The **owner** (individual or legal entity)

**Common national prefixes** :
| Country | Prefix | Example |
|---|---|---|
| France | **F-** | F-HXYZ |
| USA | **N-** | N12345 |
| UK | **G-** | G-ABCD |
| Germany | **D-** | D-EFGH |
| Spain | **EC-** | EC-IJKL |
| Canada | **C-** | C-MNOP |

**Required onboard documents** :
1. **Original Registration Certificate** (valid)
2. If aircraft has radio: **Aircraft Station License** (Valid 5 years in Europe)

**What happens if registration expires** :
- **Aircraft cannot fly** until renewal
- **Possible fine** if flying with expired registration
- **Insurance invalid** in case of accident

**For PPL pilots** : Always check registration before flying (in cockpit or on certificate).

', 1, 3, 1, 25, ["certificats","navigabilité","immatriculation","type certificate","airworthiness","registration","PPL","documents"], ["Comprendre les 3 certificats essentiels d'un aéronef (Type, Navigabilité, Immatriculation), leur validité, et les documents obligatoires à bord pour un vol PPL."]);

-- Questions for lesson 621491b5-0000-4000-8000-0000621491b5 (15 questions)
INSERT INTO questions (id, lesson_id, license, category, theme, subtopic, difficulty, level, question_fr, question_en, options, answer_key, explanation_fr, explanation_en, faa_note_fr, faa_note_en, tags, difficulty_score) VALUES
('0edf2399-0000-4000-8000-00000edf2399','621491b5-0000-4000-8000-0000621491b5','PPL','air law','air law','Certificat de Type',1,1,'Qui détient le Certificat de Type d''un avion ?','Who holds the Type Certificate of an airplane?','["Le pilote","Le propriétaire","Le constructeur","L''autorité civile"]','Le constructeur','Bonne réponse ! Le **constructeur** (ex: Cessna, Piper) détient le certificat de type.','Correct! The **manufacturer** (e.g., Cessna, Piper) holds the type certificate.','','FAA Type Certificate is issued to the manufacturer under 14 CFR Part 21.',["type certificate","manufacturer","basic"],0.25),
('0edf239a-0000-4000-8000-00000edf239a','621491b5-0000-4000-8000-0000621491b5','PPL','air law','air law','Certificat de Type',1,1,'Quel est le rôle principal du Certificat de Type ?','What is the main purpose of the Type Certificate?','["Certifier que l''avion individuel est en bon état","Certifier que la CONCEPTION du modèle respecte les normes de sécurité","Identifier le propriétaire de l''avion","Permettre la vente de l''avion"]','Certifier que la CONCEPTION du modèle respecte les normes de sécurité','Le certificat de type certifie la **conception du modèle** (pas l''avion individuel).','Type certificate certifies the **model design** (not the individual aircraft).','','',["type certificate","design","safety","basic"],0.30),
('0edf239b-0000-4000-8000-00000edf239b','621491b5-0000-4000-8000-0000621491b5','PPL','air law','air law','Certificat de Type',1,1,'Quelle est la validité du Certificat de Type ?','What is the validity of the Type Certificate?','["1 an","5 ans","10 ans","Illimitée"]','Illimitée','Le certificat de type est **valable illimitément** tant que le modèle reste conforme.','Type certificate is **valid indefinitely** as long as the model remains compliant.','','',["type certificate","validity","basic"],0.25),
('0edf239c-0000-4000-8000-00000edf239c','621491b5-0000-4000-8000-0000621491b5','PPL','air law','air law','Certificat de Type',1,1,'Lequel n''est PAS une partie du Certificat de Type ?','Which is NOT part of the Type Certificate?','["Data Sheet","Conditions de limitation","Manuel de vol","Journal de bord de l''avion individuel"]','Journal de bord de l''avion individuel','Le **journal de bord** est pour l''avion individuel, pas pour le certificat de type (modèle).','**Logbook** is for individual aircraft, not for type certificate (model).','','',["type certificate","documentation","basic"],0.35),
('0edf275e-0000-4000-8000-00000edf275e','621491b5-0000-4000-8000-0000621491b5','PPL','air law','air law','Certificat de Navigabilité',1,1,'Quel certificat certifie qu''un avion INDIVIDUEL est en bon état ?','Which certificate certifies that an INDIVIDUAL aircraft is in good condition?','["Certificat de Type","Certificat de Navigabilité","Certificat d''Immatriculation","Licence de pilote"]','Certificat de Navigabilité','Le **certificat de navigabilité** certifie l''état de l''avion INDIVIDUEL.','**Airworthiness Certificate** certifies the condition of the INDIVIDUAL aircraft.','','',["airworthiness","individual aircraft","basic"],0.25),
('0edf275f-0000-4000-8000-00000edf275f','621491b5-0000-4000-8000-0000621491b5','PPL','air law','air law','Certificat de Navigabilité',1,1,'Qui détient le Certificat de Navigabilité ?','Who holds the Airworthiness Certificate?','["Le constructeur","Le propriétaire/exploitant","Le pilote","L''autorité civile"]','Le propriétaire/exploitant','Le **propriétaire/exploitant** détient le certificat de navigabilité.','The **owner/operator** holds the airworthiness certificate.','','FAA Standard Airworthiness Certificate is issued to the aircraft owner.',["airworthiness","owner","basic"],0.30),
('0edf2760-0000-4000-8000-00000edf2760','621491b5-0000-4000-8000-0000621491b5','PPL','air law','air law','Certificat de Navigabilité',2,2,'Quelle condition est NÉCESSAIRE pour que le certificat de navigabilité reste valide ?','Which condition is NECESSARY for the airworthiness certificate to remain valid?','["Renouvellement tous les 3 ans","Maintenance régulière conforme au programme","Changement de propriétaire","Modification de l''avion"]','Maintenance régulière conforme au programme','La **maintenance régulière** est la condition clé pour la validité du C of A.','**Regular maintenance** is the key condition for C of A validity.','','FAA requires compliance with Airworthiness Directives (ADs) and periodic inspections.',["airworthiness","maintenance","validity","intermediate"],0.40),
('0edf2761-0000-4000-8000-00000edf2761','621491b5-0000-4000-8000-0000621491b5','PPL','air law','air law','Certificat de Navigabilité',1,1,'Où se trouve généralement le certificat de navigabilité dans l''avion ?','Where is the airworthiness certificate usually located in the aircraft?','["Dans le coffre","Dans le cockpit (visible)","Dans le manuel de vol uniquement","Sur l''hélice"]','Dans le cockpit (visible)','Il est **affiché dans le cockpit** (visible pour les passagers).','It is **displayed in the cockpit** (visible to passengers).','','FAA requires it to be displayed where it is visible to passengers.',["airworthiness","location","cockpit","basic"],0.25),
('0edf2762-0000-4000-8000-00000edf2762','621491b5-0000-4000-8000-0000621491b5','PPL','air law','air law','Certificat de Navigabilité',1,1,'Quel document N''EST PAS obligatoire à bord pour un vol PPL ?','Which document is NOT mandatory onboard for a PPL flight?','["Certificat de navigabilité","Certificat d''immatriculation","Certificat de type","Manuel de vol approuvé"]','Certificat de type','Le **certificat de type** n''est pas à bord (le constructeur le détient).','**Type certificate** is not onboard (manufacturer holds it).','','',["documents","onboard","type certificate","basic"],0.35),
('0edf2b24-0000-4000-8000-00000edf2b24','621491b5-0000-4000-8000-0000621491b5','PPL','air law','air law','Certificat d''Immatriculation',1,1,'Quelle est la validité du certificat d''immatriculation en Europe ?','What is the validity of the registration certificate in Europe?','["1 an","3 ans","5 ans","Illimitée"]','3 ans','En Europe, l''immatriculation expire après **3 ans** (renouvellement obligatoire).','In Europe, registration expires after **3 years** (mandatory renewal).','','FAA registration (USA) is valid indefinitely, unlike Europe.',["registration","validity","Europe","basic"],0.30),
('4d063894-0000-4000-8000-00004d063894','621491b5-0000-4000-8000-0000621491b5','PPL','air law','air law','Certificat d''Immatriculation',1,1,'Quel est le préfixe d''immatriculation pour la France ?','What is the registration prefix for France?','["N-","G-","F-","D-"]','F-','La France utilise le préfixe **F-** (ex: F-HXYZ).','France uses prefix **F-** (e.g., F-HXYZ).','','USA uses N-, UK uses G-, Germany uses D-.',["registration","prefix","France","basic"],0.20),
('4d063895-0000-4000-8000-00004d063895','621491b5-0000-4000-8000-0000621491b5','PPL','air law','air law','Certificat d''Immatriculation',1,1,'Quel est le préfixe d''immatriculation pour les USA ?','What is the registration prefix for USA?','["F-","N-","G-","C-"]','N-','Les USA utilisent le préfixe **N-** (ex: N12345).','USA uses prefix **N-** (e.g., N12345).','','',["registration","prefix","USA","basic"],0.20),
('4d063896-0000-4000-8000-00004d063896','621491b5-0000-4000-8000-0000621491b5','PPL','air law','air law','Certificat d''Immatriculation',1,1,'Que se passe-t-il si l''immatriculation expire ?','What happens if registration expires?','["L''avion peut toujours voler","L''avion ne peut pas voler jusqu''au renouvellement","Seulement une amende, vol autorisé","L''assurance reste valable"]','L''avion ne peut pas voler jusqu''au renouvellement','**Vol interdit** jusqu''au renouvellement de l''immatriculation.','**Flight prohibited** until registration is renewed.','','',["registration","expiration","flight prohibition","basic"],0.35),
('4d063897-0000-4000-8000-00004d063897','621491b5-0000-4000-8000-0000621491b5','PPL','air law','air law','Certificat d''Immatriculation',2,2,'Si l''avion a une radio, quel document supplémentaire est requis à bord ?','If the aircraft has a radio, which additional document is required onboard?','["Certificat de type","Licence de station d''aéronef","Journal de bord","Manuel du pilote"]','Licence de station d''aéronef','**Licence de station d''aéronef** requise si radio embarquée (valide 5 ans Europe).','**Aircraft Station License** required if radio onboard (valid 5 years Europe).','','FAA requires FCC radio license for international flights.',["radio","station license","documents","intermediate"],0.40),
('4d063898-0000-4000-8000-00004d063898','621491b5-0000-4000-8000-0000621491b5','PPL','air law','air law','Certificat d''Immatriculation',1,1,'Quel certificat atteste de la propriété de l''avion ?','Which certificate attests to the ownership of the aircraft?','["Certificat de Type","Certificat de Navigabilité","Certificat d''Immatriculation","Manuel de vol"]','Certificat d''Immatriculation','Le **certificat d''immatriculation** atteste de la propriété de l''aéronef.','**Registration Certificate** attests to aircraft ownership.','','',["registration","ownership","basic"],0.30);

-- ============================================================================
-- 3. ÉTUDIANT DE TEST
-- ============================================================================
INSERT INTO students (id, email, password_hash, lang, preferred_license, hearts, xp, streak, user_level) VALUES
('00000000-0000-4000-8000-000000000001', 'test@aeropath.app', 'test123', 'fr', 'PPL', 5, 0, 0, 1);

INSERT INTO user_gamification (user_id, hearts, xp, streak, level, preferred_language, preferred_license) VALUES
('00000000-0000-4000-8000-000000000001', 5, 0, 0, 1, 'fr', 'PPL');

-- ============================================================================
-- RÉSUMÉ: 3 modules, 9 concepts, 45 questions
-- ============================================================================
