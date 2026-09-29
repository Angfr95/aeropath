# Altimétrie PPL : approfondissement pédagogique — 30 septembre 2026

## Résultat

La leçon [PPL_MET_002](../../content/fr/ppl/meteo/PPL_MET_002_pressure.json) et ses 15 questions sont réécrites en français et en anglais, puis importées dans l'application locale. La déclinaison LAPL reçoit les mêmes améliorations. La langue espagnole n'a pas été retravaillée dans ce passage.

La [relecture initiale des 63 leçons](ppl-content-review-2026-09-30.md) avait constaté une explication introductive de l'altimétrie, sans exemple chiffré, et des distracteurs trop faciles à éliminer. Ce lot traite ce chapitre ; il ne clôt pas les autres lacunes PPL.

## Ce que l'élève apprend maintenant

- Distinguer altitude, hauteur au-dessus d'une référence et distance verticale au relief. Un QFE ne fournit pas une mesure permanente de la hauteur au-dessus du terrain survolé.
- Identifier QNH, QFE et pression standard ; comprendre pourquoi un niveau de vol ne garantit pas une altitude vraie identique au nombre affiché.
- Calculer le changement d'indication produit par un changement de calage, en séparant la rotation du bouton d'un déplacement physique de l'avion.
- Utiliser altitude de transition, niveau de transition et sens de traversée de la couche de transition, avec des valeurs fictives annoncées comme telles.
- Raisonner sur un calage trop élevé ou trop bas, puis distinguer cette erreur de l'effet d'une couche d'air froide.
- Comparer une indication au sol à une référence connue et à la tolérance applicable, sans conclure trop vite à une panne ou au bon fonctionnement complet de l'instrument.

Les calculs utilisent **27 ft/hPa comme approximation fournie de basse altitude**. Cette valeur n'est présentée ni comme exacte à toute altitude, ni comme une tolérance instrumentale, ni comme une correction de température.

## Évaluation et illustrations

Les 15 questions ont chacune quatre propositions et un corrigé FR/EN qui explique le raisonnement et les erreurs proposées. Dix anciens vrai/faux sont ainsi remplacés. Le corpus PPL conserve 945 questions : 455 à quatre choix et 490 vrai/faux après ce lot.

L'exercice de synthèse fournit une indication de 3 100 ft, un calage utilisé de 1 022 hPa, un QNH applicable de 1 010 hPa et un obstacle à 2 300 ft AMSL. Avec le modèle explicitement fourni, l'élève doit calculer 2 776 ft, puis une marge de 476 ft, et constater un manque de 24 ft par rapport à la contrainte fictive de 500 ft. Cette contrainte d'exercice n'est pas enseignée comme une règle générale de survol. Le modèle suppose explicitement que l'altitude corrigée au QNH peut être assimilée à l'altitude vraie et néglige les autres erreurs.

Le SVG du cours montre un autre exemple résolu : 1 500 ft au calage 1 005 deviennent 1 770 ft au calage 1 015, à pression statique constante. Le SVG joint à l'énoncé de synthèse ne contient que ses données, sans résultat. Les quatre images FR/EN ont été rendues et inspectées visuellement : texte lisible, sans débordement constaté. Les corrigés disposent aussi de l'illustration du cours.

Le générateur ciblé est [ppl_altimetry_pedagogy.py](../../scripts/illustrations/ppl_altimetry_pedagogy.py). Après une régénération générale des illustrations, exécuter cette surcouche puis `node scripts/illustrations/ppl_unit_variants.mjs` pour conserver les schémas approfondis et leurs variantes d'unités.

## Références utilisées et limites

- [DGAC — SERA France, version du 17 janvier 2026](https://www.ecologie.gouv.fr/sites/default/files/documents/sera_complet_version17012026.pdf), SERA.8015(eb) : références verticales et traversée de la couche de transition. Les valeurs numériques de transition proposées dans les exercices sont fictives.
- [CAA — Skyway Code](https://www.caa.co.uk/publication/download/16110) : références altimétriques et approximation de variation de pression. Les procédures nationales britanniques ne sont pas transposées au cadre France/EASA du cours.
- [FAA — AIM, chapitre 7, section 3](https://www.faa.gov/air_traffic/publications/atpubs/aim_html/chap7_section_3.html) : distinction entre altitude indiquée et altitude vraie en air froid. Les procédures IFR américaines de correction ne sont pas importées dans ce chapitre.

Les situations, données et distracteurs sont des exercices originaux. Ils ne proviennent pas d'une banque officielle confidentielle. Les calculs accompagnent l'explication des principes ; les procédures opérationnelles restent dépendantes des publications et documents applicables.

## Vérification après modification

- Suite Python : **40 tests réussis**, dont les vérifications indépendantes des signes de correction, des références de hauteur, du calcul de marge et des clés de réponse bilingues.
- Vérifications ciblées du rendu frontend : **9 tests réussis** concernant Markdown, illustrations des QCM et variantes d'unités. Pas de revendication d'une inspection visuelle de toutes les pages de l'application.
- Import local PPL et LAPL avec sauvegarde préalable. Les identifiants de questions existants et les données personnelles/progressions sont inchangés.
- Vérification par l'API après import : **2 leçons, 30 questions, 60 réponses invitées**. Textes, propositions et explications FR/EN comparés au contenu source ; 15 questions distinctes par quiz ; bonnes et mauvaises réponses vérifiées.

Ces contrôles établissent la cohérence du lot importé. L'appréciation pédagogique repose sur la relecture des explications et des exercices, pas sur le seul succès des tests. Les prochaines priorités identifiées restent le bilan carburant complet et le plan de vol appliqué.
