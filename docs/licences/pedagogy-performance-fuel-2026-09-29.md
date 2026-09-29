# Révision pédagogique CPL — performances et préparation du vol

Date : 29 septembre 2026. Priorité français et anglais.

## Périmètre

Deux chapitres existants : `CPL_032_001` et `CPL_033_001`, soit 30 QCM relus. Six énoncés de performances et les quinze énoncés de préparation du vol ont été réécrits, avec leurs choix et leurs corrigés FR/EN. Les exemples guidés emploient des données différentes des exercices de synthèse pour faire travailler le transfert de méthode.

Ce lot approfondit l'interpolation et le raisonnement carburant. Il reste une couverture partielle des matières CPL 032 et 033 ; il ne constitue pas une validation complète du programme ou une banque officielle d'examen. Les politiques carburant et facteurs de correction fournis dans les exercices sont fictifs et explicitement présentés comme tels. Le choix de l'anglais ou des unités US/britanniques ne change pas le cadre EASA des cours.

## Performances : ce que l'élève doit apprendre

- Distinguer roulement et distance jusqu'à une hauteur écran, sans transformer une hauteur en majoration horizontale.
- Interpoler sur deux axes avec des fractions différentes, plutôt que mémoriser la moyenne de quatre cases.
- Identifier les corrections déjà incluses pour ne pas les appliquer deux fois.
- Respecter l'ordre d'une multiplication et d'une majoration additive.
- Arrondir selon la consigne, comparer des distances de même définition et reconnaître un déficit même faible.

Exemple guidé supplémentaire : les cases 400/480 m et 560/680 m donnent 485 m à 500 ft et 10 °C. La moyenne 530 m répond à un autre point.

Dossier autonome : cases 600/720 m à 0 ft et 800/960 m à 2 000 ft, aux températures 5/25 °C. À 1 500 ft et 10 °C, résultat brut 787,5 m. Le facteur 1,12 donne 882 m, soit 18 m de marge pour 900 m disponibles. Avec 1,08 supplémentaire, 952,56 m deviennent 953 m par excès : déficit de 3 m pour 950 m disponibles.

Les distracteurs reproduisent des erreurs calculées : moyenne des quatre cases, correction omise ou répétée, addition des pourcentages, arrondi vers le bas. Leur résultat a été recalculé, y compris les valeurs incorrectes proposées.

## Préparation du vol : ce que l'élève doit apprendre

- Décomposer le trajet en phases et calculer chaque durée avec la vitesse sol correspondante.
- Convertir minutes et heures sans confondre 1 h 24 avec 1,24 h.
- Distinguer carburant total, utilisable, engagé, prévu à l'arrivée et marge au-dessus de la quantité protégée.
- Comparer prévu et constaté au même point ; ne pas soustraire une deuxième fois un écart déjà reflété dans le carburant restant.
- Recalculer le besoin futur, puis comparer des options avec leurs contraintes de piste et de route.
- Convertir masse et volume avec une densité fournie, puis vérifier les contraintes de chargement.

Exemple guidé : deux segments donnent 65 minutes et 37 L. Le QCM indépendant utilise d'autres distances, vitesses et débits : 75 minutes et 38 L.

Synthèse : 58 L utilisables, protection fictive de 24 L. Les arrivées calculées sont A : 26 L, B : 33 L, C : 19 L. B est exclu par sa piste ; C ne satisfait pas la protection. A seul satisfait les contraintes données. Cette comparaison évite de réduire la décision au terrain le plus proche ou au plus grand solde carburant.

Les quantités protégées ne sont pas présentées comme des seuils d'annonce MINIMUM FUEL ou MAYDAY FUEL.

## Illustrations

Quatre supports en FR/EN : tableau autonome de performances, tableau du premier exemple, dossier des trois terrains et méthode de bilan carburant. Les tableaux joints aux énoncés contiennent les données sans révéler les résultats. Les corrigés de l'ancien tableau gardent leur propre illustration, afin de ne pas mélanger les deux jeux de données.

Les huit SVG FR/EN ont été rendus et inspectés. Des variantes d'unités US et britanniques sont générées avec les équivalences en pied de figure ; les valeurs sources restent celles du calcul. Deux variantes représentatives ont également été inspectées visuellement. Cette vérification porte sur les images rendues, pas sur une inspection visuelle de chaque page dans le navigateur.

## Références et limites de la vérification

Le raisonnement distance/vitesse sol/temps et le besoin carburant sont cohérents avec le [FAA Pilot’s Handbook of Aeronautical Knowledge, chapitre 16](https://www.faa.gov/sites/faa.gov/files/18_phak_ch16.pdf), consulté pour les principes de navigation et de consommation. Cette référence pédagogique ne remplace pas les règles EASA applicables à l'exploitation.

Les références EASA déjà rattachées aux chapitres restent présentes. Ce lot n'ajoute pas de nouvelle affirmation réglementaire sur les réserves ou les coefficients obligatoires. Les calculs des exercices fictifs ont été contrôlés indépendamment, avec des fractions exactes, et les réponses correspondantes vérifiées dans les deux langues.

## Vérification et régénération

- 37 contrôles de contenu réussis, dont quatre nouveaux contrôles des calculs, contraintes et données des SVG.
- 9 contrôles du rendu Markdown, des illustrations et des systèmes d'unités réussis.
- Sources pérennes : `scripts/seed/cpl/performance.py` et `scripts/seed/cpl/flight_planning.py` ; génération CPL sans suppression des leçons ou questions existantes.

Après les générateurs généraux d'illustrations, exécuter les retouches pédagogiques dans cet ordre :

```text
python scripts/illustrations/reviewed_chapter_diagrams.py
python scripts/illustrations/cpl_performance_fuel_pedagogy.py
node scripts/illustrations/ppl_unit_variants.mjs
```

Pour le contenu : `python scripts/seed/cpl_content.py`.

## Résultat dans l'application locale

Docker et l'application ont été redémarrés. Une sauvegarde PostgreSQL a précédé l'import CPL. Les identifiants des questions et les empreintes des tables de comptes, historique, progression et gamification sont inchangés après import.

La vérification par l'API a confirmé les deux cours, leurs textes FR/EN, les 30 questions et leurs choix dans les deux langues. Les 60 soumissions invitées (une bonne et une mauvaise réponse par question) renvoient le résultat et le corrigé attendus. Chaque quiz contient 15 questions distinctes ; les données personnelles restent inchangées après ces essais.

- [Ouvrir les performances](http://localhost:3000/lessons/CPL/performance/1c8c8fde-0000-4000-8000-00001c8c8fde)
- [Ouvrir la préparation du vol](http://localhost:3000/lessons/CPL/flight_planning/1c9aa75f-0000-4000-8000-00001c9aa75f)
