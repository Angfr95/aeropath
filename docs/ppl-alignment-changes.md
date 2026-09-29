# Alignement pédagogique PPL(A) — modifications du 18 au 22 septembre 2026

## Livré

- **61 leçons et 915 questions FR/EN**, avec 15 questions par leçon.
- Trois anciennes leçons réglementaires reprises : distinction Chicago/OACI,
  cadre EASA, formation, maintien des privilèges, documents et immatriculation.
  Les affirmations non étayées sur les durées de conversion, les tarifs et des
  durées européennes universelles de certificats ont été retirées.
- Cinq nouvelles leçons, 75 QCM : stabilité/effet de sol, physiologie/vision,
  minima VMC de jour, triangle du vent et enveloppe de centrage.
- 79 SVG par langue et 316 variantes d’unités. Les nouveaux schémas comprennent
  des oscillations amorties/divergentes, un triangle vectoriel et une enveloppe
  de centrage chiffrée.
- **385 QCM à quatre choix** alimentent les examens blancs ; les **530 vrai/faux**
  sont conservés comme exercices de révision. Un vrai/faux ne peut pas entrer dans
  une épreuve même si ses étiquettes sont incorrectes.
- Les énoncés repérés comme dépendants de la question précédente ont été rendus
  autonomes. Les calculs sont explicités dans les corrections.

## Examen blanc

Accès public : [examen blanc local](http://localhost:3000/ppl-exam), également
accessible depuis la page Questions. Les neuf matières suivent les quotas et
durées de la [note DSAC](https://www.ecologie.gouv.fr/sites/default/files/documents/nouveau_format_examens_theoriques_PPL_LAPL_BPL_SPL.pdf) :
128 questions, 212 minutes cumulées, 75 % exigés séparément dans chaque matière.
Les matières se passent séparément, avec un récapitulatif de leurs derniers résultats.

Le tirage répartit les questions entre les leçons disponibles. Le chronomètre
utilise une échéance absolue conservée dans l’onglet : quitter ou recharger la page
ne redonne pas du temps. À expiration, les choix sont verrouillés. Les corrections
ne sont demandées qu’après remise de l’épreuve ; les questions laissées vides
valent zéro. Un échec réseau pendant la correction permet de reprendre les
corrections non reçues sans changer les réponses.

L’accès reste libre. L’API existante enregistre les réponses si un compte est
connecté ; en mode invité les essais ne modifient pas les progressions. Le
résumé des neuf matières est conservé dans la session de l’onglet. Ce mode est un
entraînement : son chronomètre côté navigateur n’est pas un dispositif de
surveillance d’examen officiel. Langue et unités ne changent pas la juridiction.

## Traçabilité et limites

[L’index des contenus](ppl-syllabus-index.json) rattache les 61 leçons et leurs
concepts aux neuf matières ; les JSON des questions contiennent leurs références
de source, leur matière et leur format de révision. Cette correspondance par
matière n’est pas encore une matrice exhaustive de chaque sous-objectif EASA.

Les écarts ciblés de l’[audit initial](ppl-official-alignment-audit.md) ont été
traités, mais **la couverture intégrale du programme et la validation pédagogique
par un instructeur ne sont pas acquises**. Ne pas présenter le nombre de questions,
le format conforme des épreuves ou les tests logiciels comme une certification
de la formation. La banque demeure une base de révision originale.

À approfondir dans les lots suivants : règles de priorité et signaux, détails
des espaces et services, météo et cartes de prévision, radio-navigation,
performances à partir de documents variés, procédures et connaissances systèmes.
Les nouveaux cours détaillent les lacunes nommées dans l’audit ; ils ne remplacent
pas à eux seuls tous les objectifs de formation d’un ATO/DTO.

## Régénération et validation

Les JSON restent éditables. Les scripts `review_ppl_airlaw.py` et
`ppl_alignment_lessons.py` sont des lots éditoriaux datés : les rejouer écrase
les chapitres concernés. Après une modification manuelle, préférer régénérer
uniquement les artefacts nécessaires.

Ordre si l’on rejoue l’ensemble de ce lot :

```powershell
python scripts/seed/review_ppl_airlaw.py
python scripts/seed/ppl_alignment_lessons.py
python scripts/seed/ppl_standalone_questions.py
python scripts/seed/ppl_exam_index.py
python scripts/illustrations/ppl_alignment_diagrams.py
python scripts/illustrations/ppl_illustrations.py
python scripts/illustrations/ppl_illustrations_en.py
python scripts/illustrations/ppl_quiz_illustrations.py
node scripts/illustrations/ppl_unit_variants.mjs
python scripts/seed/test_ppl_alignment.py
python scripts/seed/ppl_content.py
```

Importer ensuite le SQL transactionnel après sauvegarde. Il conserve les
identifiants et ne supprime pas les comptes ni les résultats antérieurs. Un
résultat historique conserve sa valeur même si l’énoncé correspondant a été
corrigé : il ne doit pas être interprété comme une nouvelle évaluation du texte.

Vérifications effectuées : compilation frontend, 23 tests frontend, tests Go de
sélection des matières et des routes existantes ; 61 quiz, 915 questions et 1 830
réponses invitées contrôlés via API. Les neuf nouvelles épreuves ont les quotas,
durées et quatre choix attendus, sans correction exposée dans le tirage. Les
empreintes des données personnelles et les identifiants existants sont conservés.
Rendu de 158 SVG et contrôle de 316 variantes sans débordement détecté.
Le navigateur de test n’étant pas connecté, aucun parcours visuel complet dans
l’interface n’est déclaré validé.


## Lot du 22 septembre 2026 — priorités et espaces

Deux leçons FR/EN supplémentaires : `PPL_AIRLAW_012` (prévention des collisions,
priorités, dépassement, finale et barres d’arrêt) et `PPL_AIRLAW_013` (classes,
services ATS, clairances, RMZ et TMZ). Chacune contient quatre sections,
15 QCM originaux à quatre choix, des explications bilingues et un SVG de synthèse
repris dans les corrections. Les SVG existent aussi avec les profils US et britannique.

Les références précises SERA.3201, SERA.3210, SERA.6001 et SERA.6005 sont
attachées aux concepts et questions ; consultation EASA du 22 septembre 2026.
La banque atteint 63 leçons et 945 questions, dont 415 QCM utilisables en examen
blanc et 530 vrai/faux de révision. Le choix de langue ou d’unités ne change pas
le cadre réglementaire européen. Les signaux, l’ensemble des procédures locales
et les autres lacunes de l’audit restent à compléter ; validation instructeur à obtenir.

Validation de ce lot : 5 tests pédagogiques Python et 23 tests frontend réussis ;
81 SVG FR, 81 SVG EN et 324 variantes contrôlés sans débordement détecté.
Import local après sauvegarde, conservation des identifiants de questions et des
données personnelles vérifiée. Les 30 questions ajoutées ont été contrôlées via
l’API avec une réponse correcte et une incorrecte chacune en mode invité.
Les neuf épreuves renvoient les 128 questions attendues sans exposer les réponses.
