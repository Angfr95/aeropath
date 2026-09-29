# Vérification de l’audit du 27 septembre 2026

Les contrôles ci-dessous décrivent le premier lot de cette journée. Le complément
en fin de document décrit les quatre chapitres ajoutés ensuite.

Voir l’[audit de chaque matière](exam-alignment-audit.md) et son
[inventaire structuré](exam-alignment-audit.json).

## Contrôles effectués

- Audit de 51 couples parcours/matière : aucune discordance des quotas déclarés
  avec les tableaux DGAC/EASA cités. Les matières absentes sont comptabilisées.
- 10 tests Python, 23 tests frontend et tests Go des paquets `internal/learning`
  et `internal/transport/http` réussis.
- API : 81 quiz LAPL/CPL/ATPL/IR, chacun de 15 questions ; 1 215 questions
  vérifiées avec réponse correcte et incorrecte, soit 2 430 soumissions invitées.
  Textes FR/EN, choix et corrections concordent avec les fichiers sources.
- API PPL : neuf épreuves, 128 QCM au total, quotas et durées conformes au
  tableau DGAC, quatre propositions par question, aucun corrigé exposé dans
  les réponses de tirage vérifiées.
- Import des six derniers chapitres effectué après sauvegarde ; conservation
  des identifiants existants et empreintes des données personnelles vérifiée.
- 36 SVG CPL/ATPL/IR rendus en FR/EN sans chevauchement ou débordement détecté ;
  648 variantes d’unités contrôlées. Des rendus ont été inspectés visuellement.

## Ce que ces contrôles ne démontrent pas

Ils ne valident pas individuellement chaque affirmation ou distracteur et ne
prouvent pas une couverture complète des objectifs EASA. Les six nouveaux
chapitres sont des supports de travail partiels. L’ECQB officielle n’a pas été
consultée ; les questions sont originales. Aucun parcours visuel complet dans
le navigateur n’a été vérifié dans cette passe.

Le mode examen complet reste réservé au PPL. Le LAPL dispose du programme
commun et de ses quiz propres ; CPL, ATPL et IR ont chacun six chapitres,
avec plusieurs matières encore absentes. La prochaine revue pédagogique doit
traiter les priorités du rapport et documenter chaque objectif réellement validé.

## Complément — quatre matières développées

Ajout de `IR_010_001`, `IR_040_001`, `CPL_022_001` et `ATPL_031_001` :
60 nouvelles questions, 15 par chapitre, en FR/EN avec SVG. Import après sauvegarde,
sans perte d’identifiants ou de données personnelles. Vérification API complète
après import : 85 quiz, 1 275 questions et 2 550 soumissions invitées.

Les 10 tests Python et 23 tests frontend passent. Les nouvelles réponses
numériques sont contrôlées indépendamment : minima IFR sous hypothèses explicites,
différence de pressions, masses successives, charge maximale et centrage en % MAC.
44 SVG professionnels/IR ont été rendus et 664 variantes contrôlées.

Le catalogue atteint CPL 7 chapitres / 105 questions, ATPL 7 / 105 et IR 8 / 120.
L’audit des 51 couples parcours/matière reste sans discordance de quota ;
présence de contenu dans 7/13 matières CPL, 6/13 ATPL et 7/7 IR.
La couverture pédagogique complète n’est toujours pas établie.
