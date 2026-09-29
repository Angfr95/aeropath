# Organisation des parcours

Chaque parcours possède sa section, ses matières et ses chapitres :

- [PPL](../ppl-content.md)
- [LAPL](lapl.md)
- [CPL](cpl.md)
- [ATPL](atpl.md)
- [IR](ir.md)

## Structure du code

```text
content/fr/
  ppl/<matière>/<chapitre>.json
  lapl/<matière>/<chapitre>.json
  cpl/<matière>/<chapitre>.json
  atpl/<matière>/<chapitre>.json
  ir/<matière>/<chapitre>.json

scripts/seed/
  shared/                 # validation, import SQL et construction des chapitres
  lapl/navigation.py
  cpl/mass_and_balance.py
  atpl/flight_planning.py
  ir/navigation.py
  ppl_content.py          # entrée PPL existante
  lapl_content.py
  cpl_content.py
  atpl_content.py
  ir_content.py

scripts/illustrations/
  svg_common.py           # primitives graphiques
  licence_diagrams.py     # export bilingue commun
  ppl_illustrations.py
  lapl_illustrations.py
  cpl_illustrations.py
  atpl_illustrations.py
  ir_illustrations.py
```

Ajouter une matière au dossier du parcours concerné, puis appeler sa construction
depuis le point d’entrée de ce parcours. Les dossiers de contenu restent la
source utilisée par le validateur et l’import SQL. Les règles pédagogiques,
les sources et les données spécifiques à une matière restent dans son module.
Le code partagé ne choisit pas les sujets à enseigner.

## Vérifications communes

**Priorité éditoriale : français et anglais.** Les autres langues sont reportées.
La [reprise pédagogique de l’instrumentation CPL](pedagogy-instruments-2026-09-29.md)
ajoute des raisonnements par pressions, des contre-exemples, quinze QCM remaniés
et deux schémas bilingues sur les pannes et l’indépendance des sources.
La [reprise pédagogique du givrage IR](pedagogy-icing-2026-09-28.md) développe les
mécanismes, exemples résolus et décisions sous contraintes, avec quinze QCM
remaniés et deux nouveaux schémas bilingues.
La [contre-vérification réglementaire du 28 septembre](regulatory-level-review-2026-09-28.md)
confirme le format PPL mais relève des omissions de programme, dont SERA.14083,
et des QCM encore trop faciles. Le niveau complet des examens n’est pas validé.
Le [lot masse et centrage FR/EN](course-quality-mass-and-balance-2026-09-27.md)
reprend quatre chapitres PPL/LAPL/CPL/ATPL, leurs 60 corrections et leurs exemples.
Le complément du 28 septembre remanie dix questions CPL/ATPL avec des calculs
combinés et associe un nouveau SVG de méthode à treize corrections, en FR/EN.

La [revue de qualité pédagogique](course-quality-review-2026-09-27.md) conclut que
les contenus restent principalement des fiches de révision : profondeur,
distracteurs, objectifs et exercices sur annexes demandent encore du travail.

Le [bilan des corrections et de la qualité](course-quality-improvements-2026-09-27.md)
décrit les corrections importées, les chapitres remaniés et les lacunes restantes.
L'[inventaire des 148 chapitres](course-quality-inventory.md) sert au suivi de
la reprise pédagogique ; ses indicateurs automatiques ne valent pas validation.
Les défauts de position des vrai/faux et de SVG CPL de la
[troisième passe](course-quality-third-pass-2026-09-27.md) sont désormais corrigés.

L’[audit par matière](exam-alignment-audit.md) compare les 51 couples parcours/matière
aux formats DGAC/EASA, y compris les matières encore absentes. Il distingue les
quotas vérifiés de la couverture pédagogique, qui reste à valider objectif par objectif.

```text
python scripts/seed/audit_exam_alignment.py
python scripts/seed/test_licence_content.py
python scripts/seed/test_ppl_alignment.py
python scripts/seed/test_course_quality.py
python scripts/seed/test_mass_and_balance_quality.py
python scripts/seed/audit_course_quality.py
node --test web/frontend-react/tests/*.test.mjs
node scripts/illustrations/ppl_unit_variants.mjs
```

Le script historique de variantes d’unités parcourt les SVG de tous les parcours.
Les imports utilisent des transactions et des upserts de contenu ; conserver
une sauvegarde avant import. Les identifiants des leçons et questions sont stables.

## Contenus disponibles

| Parcours | Chapitres | Questions |
|---|---:|---:|
| PPL | 63 | 945 |
| LAPL | 63 | 945 |
| CPL | 7 | 105 |
| ATPL | 7 | 105 |
| IR | 8 | 120 |

Le LAPL reprend le programme théorique commun PPL/LAPL : 415 QCM à quatre choix
et 530 vrai/faux de révision, distingués dans les métadonnées. Les séries CPL,
ATPL et IR comportent uniquement des QCM à quatre choix. Les chapitres restent
des supports de formation partiels ; le remplissage des programmes professionnels
et IR doit continuer. Les pages et index de chaque parcours sont générés depuis
les contenus présents, pour éviter de présenter des matières vides comme terminées.

## Vérification du lot

10 tests Python et 23 tests frontend ; import après sauvegarde, identifiants et
données personnelles préservés. Les 85 quiz LAPL/CPL/ATPL/IR ont été vérifiés via
l’API le 27 septembre 2026 : 1 275 questions, 2 550 réponses correctes/incorrectes en mode invité.
Les neuf épreuves PPL renvoient 128 QCM avec les quotas et durées DGAC attendus.
Les tests Go de sélection et de transport HTTP passent également.
44 SVG CPL/ATPL/IR rendus sans débordement ou chevauchement de texte détecté ;
664 variantes d’unités contrôlées. Les schémas LAPL reprennent les dessins du
programme commun, avec identification LAPL sur les schémas de leçon.

### Complément du 27 septembre — matières manquantes

Quatre chapitres ajoutés : droit aérien IFR et facteurs humains IR,
circuit pitot-statique CPL, masses opérationnelles et centrage ATPL.
60 QCM supplémentaires, cours et corrections FR/EN avec SVG.
Les calculs d’obstacles, de pressions, de masses et de centrage disposent de
vérifications indépendantes dans les tests pédagogiques.

L’IR dispose maintenant de contenu dans ses sept matières, le CPL dans sept
sur treize, l’ATPL dans six sur treize. Ces nombres indiquent la présence de
contenu, pas une couverture complète des objectifs d’examen.
