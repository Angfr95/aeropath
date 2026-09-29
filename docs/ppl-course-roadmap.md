# Construction progressive du cours PPL(A)

## Cadre

Cours en français, pour le PPL avion dans le cadre européen / français.
Les leçons sont des supports de révision en cours de rédaction, à faire valider
par un instructeur. Le nombre de fichiers ne mesure pas la couverture complète
du programme : certains sujets sont seulement introduits et restent à approfondir.

## Lots rédigés

| Lot | Contenu | Leçons au total | Questions au total |
| --- | --- | ---: | ---: |
| 1 | Chapitres existants complétés et matières générales ajoutées | 50 | 186 |
| 2 | Performances, instruments, déroutement et clairances | 56 | 216 |
| Quiz | 94 questions d'application ajoutées, minimum cinq par leçon | 56 | 310 |

### Lot 2 — 16 septembre 2026

- `PPL_PERF_001` — Lire un tableau de performances.
- `PPL_PERF_002` — Taux de montée, pente et altitude-densité.
- `PPL_INS_001` — Vitesse indiquée, corrigée, vraie et vitesse sol.
- `PPL_INS_002` — Recouper les instruments et reconnaître une indication douteuse.
- `PPL_EMG_001` — Déroutement : décider avant de manquer d'options.
- `PPL_RAD_001` — Comprendre une clairance et vérifier son collationnement.

Chaque leçon de ce lot comporte un objectif explicite, des prérequis, quatre
sections, cinq questions avec explications et des références. Les identifiants
des anciennes leçons et questions sont conservés par l'import.

## Prochain lot proposé : principes du vol

- Stabilité statique et dynamique, rôle du centrage.
- Traînée induite et parasite, courbe de traînée.
- Finesse, vol plané et effet du vent.
- Effet de sol et dispositifs hypersustentateurs.

## Lots suivants à préparer

- **Masse et performances** : chargement complet, consommation et déplacement
  du centrage, lecture de plusieurs tableaux successifs.
- **Navigation** : triangle des vitesses, dérive, corrections d'estime et
  préparation d'une navigation avec décisions de repli.
- **Météorologie** : stabilité, diagrammes, lecture guidée de bulletins et
  scénarios d'évolution sur une route.
- **Facteurs humains** : vision, perception, charge de travail, stress et
  situations de décision.
- **Réglementation** : revue des trois anciennes leçons, puis cas d'application
  des espaces, minima VMC et conditions d'exercice des privilèges.
- **Procédures** : scénarios de préparation, circulation d'aérodrome et gestion
  des anomalies, toujours reliés aux procédures propres à l'appareil.

## Méthode pour chaque petit lot

1. Sélectionner quelques objectifs précis et vérifier les références pertinentes.
2. Rédiger les explications, exemples et questions dans `content/fr/ppl/`.
3. Vérifier calculs, unités, réponses et cohérence des prérequis. Ajouter un schéma
   adapté à chaque leçon dans `scripts/illustrations/ppl_diagrams.json`, puis
   lancer `python scripts/illustrations/ppl_illustrations.py` et vérifier le rendu.
4. Générer l'import avec `python scripts/seed/ppl_content.py`.
5. Sauvegarder le contenu local, importer sans effacer les résultats, puis
   vérifier les nouvelles leçons et leurs quiz via l'API publique.
6. Actualiser ce suivi sans assimiler validation technique et validation pédagogique.

## Anglais — 17 septembre 2026

Les 56 leçons, 310 questions (choix et corrections compris) et 66 SVG sont
disponibles en anglais via le sélecteur de langue. Import local et essais
des réponses validés ; identifiants et suivi des comptes conservés.

## Extension des quiz — 17 septembre 2026

- 56 leçons à 15 questions : 840 questions FR/EN, dont 530 nouvelles affirmations vrai/faux.
- Chaque correction reprend le SVG de la leçon ; 18 énoncés ont aussi leur propre schéma, dont 8 nouveaux.
- 74 SVG par langue et 296 variantes US/britanniques ; identifiants historiques conservés.

## Alignement pédagogique — 22 septembre 2026

61 leçons, 915 questions. Cinq leçons d’application supplémentaires et trois
leçons réglementaires corrigées. Mode examen blanc à neuf matières avec 385 QCM
éligibles, quotas français et chronomètre. [Bilan et limites](ppl-alignment-changes.md).


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
