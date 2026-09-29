# Audit par matière des examens théoriques

Inventaire exécuté le 2026-09-28 — avions, France / EASA.

**Conclusion : format PPL vérifié ; conformité pédagogique complète non établie pour aucun parcours.**

Cet audit contrôle tous les fichiers de contenu, les catégories d’examen, les choix, traductions, références et quotas déclarés. Il inventorie les chapitres et les priorités par matière ; il ne prétend pas valider individuellement toutes les affirmations ni tous les objectifs EASA. Une quantité suffisante de QCM ne constitue pas une preuve de représentativité.

La [contre-vérification réglementaire et pédagogique du 28 septembre 2026](regulatory-level-review-2026-09-28.md) précise les sources actuelles, les lacunes confirmées et les limites de cette comparaison. La date d’exécution de ce script ne vaut pas nouvelle validation des textes réglementaires.

## Références et limites

- [DGAC : quotas PPL/LAPL](https://www.ecologie.gouv.fr/sites/default/files/documents/nouveau_format_examens_theoriques_PPL_LAPL_BPL_SPL.pdf) : 9 matières, 128 questions, 212 minutes, 75 % par matière.
- [EASA : quotas par matière](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67) : tableau public édition août 2023 ; comparaison ATPL(A), CPL(A), IR(A) modulaire complet.
- [Amendement 2025/002/R](https://www.easa.europa.eu/sites/default/files/dfu/amc_gm_to_part-ara_-_issue_1_amendment_13.pdf) : correction de durée du 070 CPL(H), sans modification des colonnes avion contrôlées ici.
- [ECQB 2026 et objectifs](https://www.easa.europa.eu/en/domains/aircrew-and-medical/european-central-question-bank-ecqb) : banque officielle non publique ; aucune comparaison question par question avec cette banque. Les objectifs publiés et les annexes de référence sont la base de la revue pédagogique.
- IR désigne ici la qualification complète, pas le CB-IR ni le BIR. Le choix de langue/unités ne transforme pas le cours en préparation FAA ou UK CAA.

## Mode examen et couverture

Le code `internal/learning/ppl_exam.go` propose uniquement les neuf épreuves PPL. Les quotas et durées sont comparés ci-dessous à la note DGAC. Le store PPL applique 75 % par matière, zéro aux blancs, aucun point négatif, un délai conservé au rechargement et des corrections après remise. LAPL, CPL, ATPL et IR disposent de quiz de chapitre, mais pas de simulateur complet propre à leur parcours. Les limites de tentatives, sessions et validité administrative ne sont pas simulées.

Les 530 vrai/faux de chaque parcours PPL/LAPL sont des révisions et sont exclus de la sélection PPL. Le marquage `exam-mcq` indique un type de question éligible ; il ne constitue pas une validation par une autorité. La sélection PPL répartit les questions entre chapitres, sans matrice exhaustive d’objectifs. Les références génériques à une matière doivent être complétées par des objectifs précis et une revue des distracteurs. Les SVG de correction ne remplacent pas des exercices sur annexes dans les énoncés.

## PPL(A)

| Matière | Officiel Q / min | Chapitres | QCM / révision | Déficit numérique QCM | Format déclaré |
|---|---:|---:|---:|---:|---|
| airlaw — Droit aérien | 20 / 30 | 13 | 125 / 70 | 0 | Moteur PPL concordant |
| human_performance — Performance humaine | 12 / 18 | 4 | 30 / 30 | 0 | Moteur PPL concordant |
| meteorology — Météorologie | 20 / 30 | 10 | 50 / 100 | 0 | Moteur PPL concordant |
| communications — Communications | 8 / 12 | 4 | 20 / 40 | 0 | Moteur PPL concordant |
| navigation — Navigation | 20 / 50 | 9 | 55 / 80 | 0 | Moteur PPL concordant |
| principles_of_flight — Principes du vol | 12 / 18 | 3 | 25 / 20 | 0 | Moteur PPL concordant |
| operational_procedures — Procédures opérationnelles | 12 / 18 | 5 | 25 / 50 | 0 | Moteur PPL concordant |
| flight_performance_planning — Performances et préparation du vol | 12 / 18 | 5 | 35 / 40 | 0 | Moteur PPL concordant |
| aircraft_general — Connaissance générale de l’aéronef | 12 / 18 | 10 | 50 / 100 | 0 | Moteur PPL concordant |

Le déficit est seulement le nombre de QCM manquants pour un tirage sans répétition ; zéro ne prouve ni la couverture ni la difficulté attendue.

### Droit aérien

Chapitres : `PPL_AIRLAW_001` — Organisation aviation internationale (ICAO, OACI); `PPL_AIRLAW_002` — Licences : formation, privilèges et cadres réglementaires; `PPL_AIRLAW_003` — Certificats d'aéronefs (Type, Navigabilité, Immatriculation); `PPL_AIRLAW_004` — Espaces aériens et conditions d’accès; `PPL_AIRLAW_005` — Règles de vol (VFR/IFR); `PPL_AIRLAW_006` — SERA — Règles de l'air standardisées; `PPL_AIRLAW_007` — Planification de vol et dépôt de plan de vol; `PPL_AIRLAW_008` — Services de la circulation aérienne (ATC); `PPL_AIRLAW_009` — Procédures d'urgence et de détresse; `PPL_AIRLAW_010` — Documents de bord et licences; `PPL_AIRLAW_011` — Appliquer les minima VMC de jour; `PPL_AIRLAW_012` — Priorités et prévention des collisions; `PPL_AIRLAW_013` — Classes d’espace et services ATS.

**Revue à mener / approfondir :** Vérifier les signaux, les cas locaux et la traçabilité de chaque règle ; conserver les corrections SERA déjà réalisées.

### Performance humaine

Chapitres : `PPL_HP_001` — Aptitude du pilote et limites physiologiques; `PPL_HP_002` — Attention, biais et prise de décision; `PPL_HP_003` — Respiration, vision et illusions; `PPL_OP_006` — Facteurs humains et gestion des ressources.

**Revue à mener / approfondir :** Évaluer les scénarios physiologiques, illusions et décisions ; les chapitres ajoutés ne prouvent pas une couverture exhaustive.

### Météorologie

Chapitres : `PPL_MET_001` — L'atmosphère terrestre; `PPL_MET_002` — Pression atmosphérique et altimétrie; `PPL_MET_003` — Température et humidité; `PPL_MET_004` — Vent et circulation atmosphérique; `PPL_MET_005` — Nuages et précipitations; `PPL_MET_006` — Visibilité et phénomènes réduisant la visibilité; `PPL_MET_007` — Masses d'air et fronts; `PPL_MET_008` — Phénomènes dangereux (orage, givre, cisaillement); `PPL_MET_009` — METAR, TAF et SIGMET; `PPL_MET_010` — Météorologie pour le vol à vue (VFR).

**Revue à mener / approfondir :** Contrôler le décodage complet et des décisions à partir de cartes et bulletins, avec distracteurs plausibles.

### Communications

Chapitres : `PPL_COM_001` — Radiotéléphonie et collationnement; `PPL_COM_002` — Panne radio et communications d’urgence; `PPL_OP_004` — Procédures de communication radio; `PPL_RAD_001` — Comprendre une clairance et vérifier son collationnement.

**Revue à mener / approfondir :** Compléter la pratique contextualisée de phraséologie, urgence et panne ; un QCM ne valide pas la compétence orale.

### Navigation

Chapitres : `PPL_NAV_001` — Principes de base de la navigation aérienne; `PPL_NAV_002` — Cartes aéronautiques (ICAO 1:500 000); `PPL_NAV_003` — Navigation à l'estime (Dead Reckoning); `PPL_NAV_004` — Radio-navigation (VOR, NDB, DME); `PPL_NAV_005` — Navigation par satellite (GPS/GNSS); `PPL_NAV_006` — Log de navigation et calculs de vol; `PPL_NAV_007` — Repérage visuel et points de report; `PPL_NAV_008` — Planification de vol approfondie; `PPL_NAV_009` — Vent, cap et correction de route.

**Revue à mener / approfondir :** Contrôler les exercices de carte, dérive, radionavigation et limites GNSS ; vérifier les annexes utilisables sans réponse visible.

### Principes du vol

Chapitres : `PPL_POF_001` — Portance, traînée et équilibre du vol; `PPL_POF_002` — Décrochage, virage et facteur de charge; `PPL_POF_003` — Stabilité, finesse et effet de sol.

**Revue à mener / approfondir :** Vérifier les courbes, commandes, stabilité et limites sur des situations variées, au-delà des exercices introductifs.

### Procédures opérationnelles

Chapitres : `PPL_EMG_001` — Déroutement : décider avant de manquer d'options; `PPL_OP_001` — Procédures pré-vol; `PPL_OP_002` — Roulage et décollage; `PPL_OP_003` — Approche et atterrissage; `PPL_OP_005` — Procédures d'urgence en vol.

**Revue à mener / approfondir :** Contrôler turbulence de sillage, risques au sol, cisaillement et procédures anormales avec leurs conditions d’application.

### Performances et préparation du vol

Chapitres : `PPL_FP_001` — Bilan carburant, autonomie et réserves; `PPL_MB_001` — Calcul de masse et de centrage; `PPL_MB_002` — Chargement et enveloppe de centrage; `PPL_PERF_001` — Lire un tableau de performances; `PPL_PERF_002` — Taux de montée, pente et altitude-densité.

**Revue à mener / approfondir :** Diversifier tableaux AFM fictifs, enveloppes et bilans ; distinguer hypothèses d’exercice et exigences de carburant.

### Connaissance générale de l’aéronef

Chapitres : `PPL_AC_001` — Structure de l'aéronef; `PPL_AC_002` — Groupe motopropulseur (moteur, hélice); `PPL_AC_003` — Systèmes de bord (électrique, carburant, hydraulique); `PPL_AC_004` — Instruments de vol; `PPL_AC_005` — Équipements électroniques et radio; `PPL_AC_006` — Performances (décollage, atterrissage, masse et centrage); `PPL_AC_007` — Commandes de vol; `PPL_AC_008` — Limitations d'utilisation; `PPL_INS_001` — Vitesse indiquée, corrigée, vraie et vitesse sol; `PPL_INS_002` — Recouper les instruments et reconnaître une indication douteuse.

**Revue à mener / approfondir :** Vérifier systèmes, limitations, hélices et diagnostic instrumental avec des annexes représentatives.

## LAPL(A)

| Matière | Officiel Q / min | Chapitres | QCM / révision | Déficit numérique QCM | Format déclaré |
|---|---:|---:|---:|---:|---|
| airlaw — Droit aérien | 20 / 30 | 13 | 125 / 70 | 0 | Non déclaré |
| human_performance — Performance humaine | 12 / 18 | 4 | 30 / 30 | 0 | Non déclaré |
| meteorology — Météorologie | 20 / 30 | 10 | 50 / 100 | 0 | Non déclaré |
| communications — Communications | 8 / 12 | 4 | 20 / 40 | 0 | Non déclaré |
| navigation — Navigation | 20 / 50 | 9 | 55 / 80 | 0 | Non déclaré |
| principles_of_flight — Principes du vol | 12 / 18 | 3 | 25 / 20 | 0 | Non déclaré |
| operational_procedures — Procédures opérationnelles | 12 / 18 | 5 | 25 / 50 | 0 | Non déclaré |
| flight_performance_planning — Performances et préparation du vol | 12 / 18 | 5 | 35 / 40 | 0 | Non déclaré |
| aircraft_general — Connaissance générale de l’aéronef | 12 / 18 | 10 | 50 / 100 | 0 | Non déclaré |

Le déficit est seulement le nombre de QCM manquants pour un tirage sans répétition ; zéro ne prouve ni la couverture ni la difficulté attendue.

### Droit aérien

Chapitres : `LAPL_AIRLAW_001` — Organisation aviation internationale (ICAO, OACI); `LAPL_AIRLAW_002` — Licences : formation, privilèges et cadres réglementaires; `LAPL_AIRLAW_003` — Certificats d'aéronefs (Type, Navigabilité, Immatriculation); `LAPL_AIRLAW_004` — Espaces aériens et conditions d’accès; `LAPL_AIRLAW_005` — Règles de vol (VFR/IFR); `LAPL_AIRLAW_006` — SERA — Règles de l'air standardisées; `LAPL_AIRLAW_007` — Planification de vol et dépôt de plan de vol; `LAPL_AIRLAW_008` — Services de la circulation aérienne (ATC); `LAPL_AIRLAW_009` — Procédures d'urgence et de détresse; `LAPL_AIRLAW_010` — Documents de bord et licences; `LAPL_AIRLAW_011` — Appliquer les minima VMC de jour; `LAPL_AIRLAW_012` — Priorités et prévention des collisions; `LAPL_AIRLAW_013` — Classes d’espace et services ATS.

**Revue à mener / approfondir :** Vérifier les signaux, les cas locaux et la traçabilité de chaque règle ; conserver les corrections SERA déjà réalisées.

### Performance humaine

Chapitres : `LAPL_HP_001` — Aptitude du pilote et limites physiologiques; `LAPL_HP_002` — Attention, biais et prise de décision; `LAPL_HP_003` — Respiration, vision et illusions; `LAPL_OP_006` — Facteurs humains et gestion des ressources.

**Revue à mener / approfondir :** Évaluer les scénarios physiologiques, illusions et décisions ; les chapitres ajoutés ne prouvent pas une couverture exhaustive.

### Météorologie

Chapitres : `LAPL_MET_001` — L'atmosphère terrestre; `LAPL_MET_002` — Pression atmosphérique et altimétrie; `LAPL_MET_003` — Température et humidité; `LAPL_MET_004` — Vent et circulation atmosphérique; `LAPL_MET_005` — Nuages et précipitations; `LAPL_MET_006` — Visibilité et phénomènes réduisant la visibilité; `LAPL_MET_007` — Masses d'air et fronts; `LAPL_MET_008` — Phénomènes dangereux (orage, givre, cisaillement); `LAPL_MET_009` — METAR, TAF et SIGMET; `LAPL_MET_010` — Météorologie pour le vol à vue (VFR).

**Revue à mener / approfondir :** Contrôler le décodage complet et des décisions à partir de cartes et bulletins, avec distracteurs plausibles.

### Communications

Chapitres : `LAPL_COM_001` — Radiotéléphonie et collationnement; `LAPL_COM_002` — Panne radio et communications d’urgence; `LAPL_OP_004` — Procédures de communication radio; `LAPL_RAD_001` — Comprendre une clairance et vérifier son collationnement.

**Revue à mener / approfondir :** Compléter la pratique contextualisée de phraséologie, urgence et panne ; un QCM ne valide pas la compétence orale.

### Navigation

Chapitres : `LAPL_NAV_001` — Navigation LAPL : vent et route; `LAPL_NAV_002` — Principes de base de la navigation aérienne; `LAPL_NAV_003` — Cartes aéronautiques (ICAO 1:500 000); `LAPL_NAV_004` — Navigation à l'estime (Dead Reckoning); `LAPL_NAV_005` — Radio-navigation (VOR, NDB, DME); `LAPL_NAV_006` — Navigation par satellite (GPS/GNSS); `LAPL_NAV_007` — Log de navigation et calculs de vol; `LAPL_NAV_008` — Repérage visuel et points de report; `LAPL_NAV_009` — Planification de vol approfondie.

**Revue à mener / approfondir :** Contrôler les exercices de carte, dérive, radionavigation et limites GNSS ; vérifier les annexes utilisables sans réponse visible.

### Principes du vol

Chapitres : `LAPL_POF_001` — Portance, traînée et équilibre du vol; `LAPL_POF_002` — Décrochage, virage et facteur de charge; `LAPL_POF_003` — Stabilité, finesse et effet de sol.

**Revue à mener / approfondir :** Vérifier les courbes, commandes, stabilité et limites sur des situations variées, au-delà des exercices introductifs.

### Procédures opérationnelles

Chapitres : `LAPL_EMG_001` — Déroutement : décider avant de manquer d'options; `LAPL_OP_001` — Procédures pré-vol; `LAPL_OP_002` — Roulage et décollage; `LAPL_OP_003` — Approche et atterrissage; `LAPL_OP_005` — Procédures d'urgence en vol.

**Revue à mener / approfondir :** Contrôler turbulence de sillage, risques au sol, cisaillement et procédures anormales avec leurs conditions d’application.

### Performances et préparation du vol

Chapitres : `LAPL_FP_001` — Bilan carburant, autonomie et réserves; `LAPL_MB_001` — Calcul de masse et de centrage; `LAPL_MB_002` — Chargement et enveloppe de centrage; `LAPL_PERF_001` — Lire un tableau de performances; `LAPL_PERF_002` — Taux de montée, pente et altitude-densité.

**Revue à mener / approfondir :** Diversifier tableaux AFM fictifs, enveloppes et bilans ; distinguer hypothèses d’exercice et exigences de carburant.

### Connaissance générale de l’aéronef

Chapitres : `LAPL_AC_001` — Structure de l'aéronef; `LAPL_AC_002` — Groupe motopropulseur (moteur, hélice); `LAPL_AC_003` — Systèmes de bord (électrique, carburant, hydraulique); `LAPL_AC_004` — Instruments de vol; `LAPL_AC_005` — Équipements électroniques et radio; `LAPL_AC_006` — Performances (décollage, atterrissage, masse et centrage); `LAPL_AC_007` — Commandes de vol; `LAPL_AC_008` — Limitations d'utilisation; `LAPL_INS_001` — Vitesse indiquée, corrigée, vraie et vitesse sol; `LAPL_INS_002` — Recouper les instruments et reconnaître une indication douteuse.

**Revue à mener / approfondir :** Vérifier systèmes, limitations, hélices et diagnostic instrumental avec des annexes représentatives.

## CPL(A)

| Matière | Officiel Q / min | Chapitres | QCM / révision | Déficit numérique QCM | Format déclaré |
|---|---:|---:|---:|---:|---|
| 010 — Droit aérien | 33 / 45 | 0 | 0 / 0 | 33 | Matière absente |
| 021 — Cellule, systèmes et moteurs | 60 / 90 | 0 | 0 / 0 | 60 | Matière absente |
| 022 — Instrumentation | 40 / 60 | 1 | 15 / 0 | 25 | Concordant |
| 031 — Masse et centrage | 21 / 60 | 1 | 15 / 0 | 6 | Concordant |
| 032 — Performances avion | 28 / 75 | 1 | 15 / 0 | 13 | Concordant |
| 033 — Préparation et suivi du vol | 33 / 90 | 1 | 15 / 0 | 18 | Concordant |
| 040 — Performance humaine | 35 / 60 | 1 | 15 / 0 | 20 | Concordant |
| 050 — Météorologie | 63 / 90 | 0 | 0 / 0 | 63 | Matière absente |
| 061 — Navigation générale | 45 / 120 | 0 | 0 / 0 | 45 | Matière absente |
| 062 — Radionavigation | 22 / 30 | 0 | 0 / 0 | 22 | Matière absente |
| 070 — Procédures opérationnelles | 30 / 60 | 0 | 0 / 0 | 30 | Matière absente |
| 081 — Principes du vol avion | 37 / 75 | 1 | 15 / 0 | 22 | Concordant |
| 090 — Communications | 34 / 60 | 1 | 15 / 0 | 19 | Concordant |

Le déficit est seulement le nombre de QCM manquants pour un tirage sans répétition ; zéro ne prouve ni la couverture ni la difficulté attendue.

### Droit aérien

Chapitres : **aucun**.

**Revue à mener / approfondir :** Développer le parcours réglementaire propre à la licence/qualification : Part-FCL, SERA, ATS et procédures applicables. Les seuls minima IFR et changements de régime ne couvrent pas toute la matière.

### Cellule, systèmes et moteurs

Chapitres : **aucun**.

**Revue à mener / approfondir :** Développer cellule, électricité, hydraulique, carburant, moteurs et protections ; turbine et pressurisation ne suffisent pas.

### Instrumentation

Chapitres : `CPL_022_001` — Circuit pitot-statique : interpréter une panne.

**Revue à mener / approfondir :** Compléter instruments, capteurs, gyroscopes, automatismes et alertes selon les objectifs applicables à chaque parcours.

### Masse et centrage

Chapitres : `CPL_031_001` — Chargement et déplacement du centrage.

**Revue à mener / approfondir :** Développer limitations, distribution des charges et documents de chargement, pas seulement le barycentre.

### Performances avion

Chapitres : `CPL_032_001` — Performances : tableaux et marges.

**Revue à mener / approfondir :** Développer les domaines et classes de performances applicables, montée, croisière et atterrissage avec annexes.

### Préparation et suivi du vol

Chapitres : `CPL_033_001` — Bilan carburant et suivi du vol.

**Revue à mener / approfondir :** Compléter plans de vol, routes/cartes, carburant réglementaire et suivi ; des calculs fictifs seuls ne couvrent pas la matière.

### Performance humaine

Chapitres : `CPL_040_001` — Menaces, erreurs et décisions.

**Revue à mener / approfondir :** Compléter physiologie, perception et psychologie ; un chapitre ciblé sur le TEM ou la désorientation ne couvre pas toutes les sous-matières.

### Météorologie

Chapitres : **aucun**.

**Revue à mener / approfondir :** Développer atmosphère, fronts, vents, nuages, dangers et exploitation des produits météo ; le givrage seul ne suffit pas.

### Navigation générale

Chapitres : **aucun**.

**Revue à mener / approfondir :** Développer magnétisme, projections, navigation à l’estime et navigation en vol, avec lecture de cartes.

### Radionavigation

Chapitres : **aucun**.

**Revue à mener / approfondir :** Compléter principes radio, radar, GNSS et navigation de surface/PBN selon la colonne du parcours ; VOR/DME/ILS sont partiels.

### Procédures opérationnelles

Chapitres : **aucun**.

**Revue à mener / approfondir :** Créer les procédures normales/spéciales, risques et contraintes opérationnelles du parcours.

### Principes du vol avion

Chapitres : `CPL_081_001` — Virage, facteur de charge et décrochage.

**Revue à mener / approfondir :** Développer aérodynamique, stabilité, commandes, hélices et limites ; adapter la compressibilité aux objectifs de la licence.

### Communications

Chapitres : `CPL_090_001` — Clairances et boucle de collationnement.

**Revue à mener / approfondir :** Compléter les sept rubriques du tableau 090 ; le seul collationnement ne représente pas toute l’épreuve.

## ATPL(A)

| Matière | Officiel Q / min | Chapitres | QCM / révision | Déficit numérique QCM | Format déclaré |
|---|---:|---:|---:|---:|---|
| 010 — Droit aérien | 44 / 60 | 0 | 0 / 0 | 44 | Matière absente |
| 021 — Cellule, systèmes et moteurs | 80 / 120 | 2 | 30 / 0 | 50 | Concordant |
| 022 — Instrumentation | 60 / 90 | 1 | 15 / 0 | 45 | Concordant |
| 031 — Masse et centrage | 25 / 75 | 1 | 15 / 0 | 10 | Concordant |
| 032 — Performances avion | 45 / 120 | 0 | 0 / 0 | 45 | Matière absente |
| 033 — Préparation et suivi du vol | 42 / 120 | 1 | 15 / 0 | 27 | Concordant |
| 040 — Performance humaine | 48 / 90 | 0 | 0 / 0 | 48 | Matière absente |
| 050 — Météorologie | 84 / 120 | 0 | 0 / 0 | 84 | Matière absente |
| 061 — Navigation générale | 55 / 135 | 1 | 15 / 0 | 40 | Concordant |
| 062 — Radionavigation | 66 / 90 | 0 | 0 / 0 | 66 | Matière absente |
| 070 — Procédures opérationnelles | 42 / 75 | 0 | 0 / 0 | 42 | Matière absente |
| 081 — Principes du vol avion | 46 / 90 | 0 | 0 / 0 | 46 | Matière absente |
| 090 — Communications | 34 / 60 | 1 | 15 / 0 | 19 | Concordant |

Le déficit est seulement le nombre de QCM manquants pour un tirage sans répétition ; zéro ne prouve ni la couverture ni la difficulté attendue.

### Droit aérien

Chapitres : **aucun**.

**Revue à mener / approfondir :** Développer le parcours réglementaire propre à la licence/qualification : Part-FCL, SERA, ATS et procédures applicables. Les seuls minima IFR et changements de régime ne couvrent pas toute la matière.

### Cellule, systèmes et moteurs

Chapitres : `ATPL_021_001` — Turbines : cycle et surveillance; `ATPL_021_002` — Pressurisation : pression, altitude cabine et surveillance.

**Revue à mener / approfondir :** Développer cellule, électricité, hydraulique, carburant, moteurs et protections ; turbine et pressurisation ne suffisent pas.

### Instrumentation

Chapitres : `ATPL_022_001` — Vitesses air, Mach et données aérodynamiques.

**Revue à mener / approfondir :** Compléter instruments, capteurs, gyroscopes, automatismes et alertes selon les objectifs applicables à chaque parcours.

### Masse et centrage

Chapitres : `ATPL_031_001` — Masses opérationnelles et évolution du centrage.

**Revue à mener / approfondir :** Développer limitations, distribution des charges et documents de chargement, pas seulement le barycentre.

### Performances avion

Chapitres : **aucun**.

**Revue à mener / approfondir :** Développer les domaines et classes de performances applicables, montée, croisière et atterrissage avec annexes.

### Préparation et suivi du vol

Chapitres : `ATPL_033_001` — Point équitemps et point de retour.

**Revue à mener / approfondir :** Compléter plans de vol, routes/cartes, carburant réglementaire et suivi ; des calculs fictifs seuls ne couvrent pas la matière.

### Performance humaine

Chapitres : **aucun**.

**Revue à mener / approfondir :** Compléter physiologie, perception et psychologie ; un chapitre ciblé sur le TEM ou la désorientation ne couvre pas toutes les sous-matières.

### Météorologie

Chapitres : **aucun**.

**Revue à mener / approfondir :** Développer atmosphère, fronts, vents, nuages, dangers et exploitation des produits météo ; le givrage seul ne suffit pas.

### Navigation générale

Chapitres : `ATPL_061_001` — Coordonnées, routes et temps UTC.

**Revue à mener / approfondir :** Développer magnétisme, projections, navigation à l’estime et navigation en vol, avec lecture de cartes.

### Radionavigation

Chapitres : **aucun**.

**Revue à mener / approfondir :** Compléter principes radio, radar, GNSS et navigation de surface/PBN selon la colonne du parcours ; VOR/DME/ILS sont partiels.

### Procédures opérationnelles

Chapitres : **aucun**.

**Revue à mener / approfondir :** Créer les procédures normales/spéciales, risques et contraintes opérationnelles du parcours.

### Principes du vol avion

Chapitres : **aucun**.

**Revue à mener / approfondir :** Développer aérodynamique, stabilité, commandes, hélices et limites ; adapter la compressibilité aux objectifs de la licence.

### Communications

Chapitres : `ATPL_090_001` — Clairances et boucle de collationnement.

**Revue à mener / approfondir :** Compléter les sept rubriques du tableau 090 ; le seul collationnement ne représente pas toute l’épreuve.

## IR(A)

| Matière | Officiel Q / min | Chapitres | QCM / révision | Déficit numérique QCM | Format déclaré |
|---|---:|---:|---:|---:|---|
| 010 — Droit aérien | 30 / 45 | 1 | 15 / 0 | 15 | Concordant |
| 022 — Instrumentation | 20 / 30 | 1 | 15 / 0 | 5 | Concordant |
| 033 — Préparation et suivi du vol | 27 / 75 | 1 | 15 / 0 | 12 | Concordant |
| 040 — Performance humaine | 35 / 60 | 1 | 15 / 0 | 20 | Concordant |
| 050 — Météorologie | 63 / 90 | 1 | 15 / 0 | 48 | Concordant |
| 062 — Radionavigation | 44 / 60 | 2 | 30 / 0 | 14 | Concordant |
| 090 — Communications | 34 / 60 | 1 | 15 / 0 | 19 | Concordant |

Le déficit est seulement le nombre de QCM manquants pour un tirage sans répétition ; zéro ne prouve ni la couverture ni la difficulté attendue.

### Droit aérien

Chapitres : `IR_010_001` — Règles IFR : minima, clairance et changement de régime.

**Revue à mener / approfondir :** Développer le parcours réglementaire propre à la licence/qualification : Part-FCL, SERA, ATS et procédures applicables. Les seuls minima IFR et changements de régime ne couvrent pas toute la matière.

### Instrumentation

Chapitres : `IR_022_001` — Altimétrie et circuit statique.

**Revue à mener / approfondir :** Compléter instruments, capteurs, gyroscopes, automatismes et alertes selon les objectifs applicables à chaque parcours.

### Préparation et suivi du vol

Chapitres : `IR_033_001` — Préparer et actualiser un bilan IFR.

**Revue à mener / approfondir :** Compléter plans de vol, routes/cartes, carburant réglementaire et suivi ; des calculs fictifs seuls ne couvrent pas la matière.

### Performance humaine

Chapitres : `IR_040_001` — Perception, désorientation et capacités en IMC.

**Revue à mener / approfondir :** Compléter physiologie, perception et psychologie ; un chapitre ciblé sur le TEM ou la désorientation ne couvre pas toutes les sous-matières.

### Météorologie

Chapitres : `IR_050_001` — Givrage : reconnaître et décider.

**Revue à mener / approfondir :** Développer atmosphère, fronts, vents, nuages, dangers et exploitation des produits météo ; le givrage seul ne suffit pas.

### Radionavigation

Chapitres : `IR_062_001` — VOR et distance oblique DME; `IR_062_002` — ILS : guidage et profil de descente.

**Revue à mener / approfondir :** Compléter principes radio, radar, GNSS et navigation de surface/PBN selon la colonne du parcours ; VOR/DME/ILS sont partiels.

### Communications

Chapitres : `IR_090_001` — Clairances et boucle de collationnement.

**Revue à mener / approfondir :** Compléter les sept rubriques du tableau 090 ; le seul collationnement ne représente pas toute l’épreuve.

## Contrôles reproductibles

```text
python scripts/seed/audit_exam_alignment.py
python scripts/seed/test_licence_content.py
python scripts/seed/test_ppl_alignment.py
node --test web/frontend-react/tests/*.test.mjs
```

51 couples parcours/matière examinés. 0 incohérence(s) structurelle(s) ou de quota détectée(s).

## Suite nécessaire

1. Rattacher concepts et questions aux objectifs précis du programme applicable.
2. Relire les faits, calculs, distracteurs et traductions ; traiter les priorités de chaque matière.
3. Compléter les matières absentes et les exercices sur annexes avant les examens blancs professionnels.
4. Valider la répartition par sous-matière, puis étendre le simulateur aux parcours disposant d’une banque suffisante.
