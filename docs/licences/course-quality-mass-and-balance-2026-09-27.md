# Lot FR/EN — masse et centrage

## Périmètre réalisé

Priorité confirmée : **français et anglais**. Les autres langues sont reportées.
Les alertes actives de l'inventaire suivent ce périmètre ; les lacunes des autres
langues restent conservées séparément dans le JSON.

Quatre chapitres sont enrichis et importés dans la base locale :

| Chapitre | Travail principal | Quiz |
| --- | --- | ---: |
| PPL_MB_002 | Bilan pondéré, lecture d'enveloppe, ajout/retrait/déplacement et cas autonome | 15 |
| LAPL_MB_002 | Reprise du même contenu théorique avec ses identifiants LAPL | 15 |
| CPL_031_001 | Datum, contrôle inverse, déplacement de charge inverse, retrait et %MAC | 15 |
| ATPL_031_001 | Étapes de masse, plafond limitant, sensibilité au carburant, consommation et %MAC | 15 |

Cela représente trois textes sources et 45 questions distinctes, déclinés dans
quatre chapitres et 60 questions publiées. Chaque correction est développée en
FR/EN ; plusieurs distracteurs et énoncés sont également remaniés.

## Améliorations pédagogiques concrètes

### PPL/LAPL

Le cours explique pourquoi le CG est une moyenne pondérée, contrôle les unités
et situe le résultat entre les bras extrêmes. L'exemple autonome 80 kg à 1 m
et 20 kg à 3 m donne 1,4 m, pas la moyenne simple de 2 m.

La lecture de l'enveloppe identifie explicitement les axes et la limite à la
masse étudiée. L'énoncé Q8 donne maintenant le domaine de masse 700–850 kg,
sans dépendre du cours pour conclure que le point est dans l'enveloppe.
Les limites d'un arrondi près d'une frontière sont expliquées.

Trois opérations sont distinguées, avec états initiaux identifiés :

- Consommer 30 kg : retirer masse et moment ; CG 1,08125 m.
- Déplacer les bagages dans l'état initial : masse inchangée ; CG 1,0723 m.
- Ajouter séparément 10 kg au bras 2 m : 840 kg, 930 kg·m, CG 1,1071 m ;
  comparaison avec la limite avant à cette nouvelle masse.

### CPL

Le cours vérifie le calcul par les unités et la pondération, explique comment
un changement de datum modifie les nombres sans déplacer le chargement et
retrouve la formule d'un déplacement à partir des moments retiré et ajouté.

Le problème inverse demande quelle masse déplacer de 2 m pour avancer le CG
de 0,05 m à 2 300 kg : **57,5 kg à déplacer**, pas à débarquer.
Un autre cas ajoute 50 kg à 4 m et donne un CG de 2,213 m.
La conversion %MAC est contrôlée dans les deux sens.

Les questions finales demandent les vérifications réelles du chargement ; les
anciens distracteurs portant sur la couleur des bagages ou l'immatriculation
sont remplacés par des confusions de méthode plausibles.

### ATPL

Le cours distingue le contenu de la DOW, la ZFM et les étapes parking,
décollage et arrivée. L'atterrissage est vérifié par deux méthodes :
soustractions successives et ZFM + carburant restant.

Le plafond limitant est présenté comme trois inégalités simultanées. Un cas
nouveau à 7 500 kg de charge donne **700 kg** de capacité additionnelle selon
ces seuls plafonds. Un ajout séparé de 400 kg de carburant démontre qu'un
décollage sous son plafond peut néanmoins conduire à dépasser celui d'arrivée.

Le contre-exemple du réservoir avant est intégré au cours et à Q13 : retirer
1 000 kg au bras 4 m d'un avion de 10 000 kg centré à 5 m donne **5,111 m**, donc
un déplacement vers l'arrière. Le cas du réservoir arrière produit l'effet
opposé. La Q9 compare une limite de performance de 46 500 kg avec le maximum
structural de 49 000 kg et une masse prévue de 47 800 kg.

## Sources et limites

Les méthodes de moments et déplacements sont recoupées avec le
[FAA Weight and Balance Handbook](https://www.faa.gov/sites/faa.gov/files/2023-09/Weight_Balance_Handbook.pdf).
Les définitions et vérifications de masse sont replacées dans le cadre
[EASA Air Operations](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-air-operations?erules-id=ERULES-1963177438-12516).
Les nouveaux nombres et cas sont des exercices fictifs calculés indépendamment,
pas des données d'avion réel ni des questions issues de l'ECQB.

Les durées de 40 minutes PPL/LAPL, 45 minutes CPL et 50 minutes ATPL sont des
estimations de travail, pas des durées mesurées auprès d'élèves. Le contenu
reste partiel à l'échelle de la matière et attend une validation pédagogique
externe ; l'absence d'alerte automatique ne remplace pas cette validation.

## Vérification et intégration

- Relecture des textes FR/EN et des réponses attendues des trois sources ;
  propagation au LAPL à partir du PPL.
- Cinq nouveaux tests indépendants : ajout de charge/enveloppe, changement de
  datum et déplacement inverse, étapes de masse et plafond limitant, sens de
  déplacement selon le réservoir, conversion MAC inverse.
- **19 tests Python réussis** en incluant les régressions existantes.
- Audit de format sur les **51 couples parcours/matière** : aucune erreur.
- **4 quiz de 15 questions**, réponses et traductions comparées aux fichiers
  par l'API locale ; **120 réponses invitées** correctes/incorrectes vérifiées.
- Sauvegarde avant import : `.local/content-before-mass-quality-20260927.sql`.
- Identifiants des leçons/questions conservés ; comptes, progression et
  résultats personnels inchangés après import et essais.
- Les SVG existants restent associés aux leçons et aux corrections. Leurs
  données de base n'ont pas changé ; pas de nouvelle inspection du rendu
  complet dans un navigateur pendant ce lot.

Les compteurs de l'[inventaire actif FR/EN](course-quality-inventory.md) passent
de 146 à 142 chapitres signalés pour brièveté ou corrections très courtes.
Ce changement mesure uniquement ces indicateurs ; il n'attribue pas une
certification aux six chapitres approfondis à ce stade.

## Accès local

- [PPL — chargement et enveloppe](http://localhost:3000/lessons/PPL/mass_and_balance/555df131-0000-4000-8000-0000555df131)
- [LAPL — chargement et enveloppe](http://localhost:3000/lessons/LAPL/mass_and_balance/7d6134af-0000-4000-8000-00007d6134af)
- [CPL — chargement et déplacement du centrage](http://localhost:3000/lessons/CPL/mass_and_balance/1c7e785d-0000-4000-8000-00001c7e785d)
- [ATPL — masses opérationnelles et centrage](http://localhost:3000/lessons/ATPL/mass_and_balance/13833479-0000-4000-8000-000013833479)

## Contre-vérification achevée le 28 septembre 2026

Cette passe vérifie le lot déjà publié ; aucun contenu de cours n'a été modifié.

- Relecture des 45 énoncés distincts, choix et corrections FR/EN ; le chapitre
  LAPL reprend les 15 questions du PPL, soit 60 questions publiées.
- Recalcul indépendant des **40 réponses numériques publiées** (31 distinctes,
  plus neuf reprises LAPL). Comparaison avec les réponses françaises et leurs
  choix anglais, en tenant compte des arrondis : aucun écart trouvé.
- **19 tests Python réussis**, dont les contre-exemples sur le déplacement de
  centrage, les calculs de charge et les conversions de MAC.
- **Huit SVG FR/EN rendus et comparés visuellement** : enveloppe PPL/LAPL,
  tableau de chargement CPL et bilan de masses ATPL. Données et libellés
  concordants avec les exercices. Cette inspection porte sur les images,
  pas sur la présentation complète des pages dans un navigateur.
- Après relance des services locaux, **quatre quiz, 60 questions et
  120 réponses invitées correctes/incorrectes** revérifiés par l'API.
  Énoncés, choix et corrections FR/EN conformes aux fichiers ; données
  personnelles inchangées.

**Conclusion :** aucune nouvelle erreur de calcul ou divergence FR/EN repérée
dans ce périmètre. L'objectif de qualité pédagogique élevée demande encore
des améliorations : les distracteurs CPL_031_001 Q7 et Q11 restent faciles à
éliminer, et le quiz ATPL comporte beaucoup de calculs à une opération. Des
scénarios combinant choix de l'état initial, plusieurs modifications et contrôle
de limites permettraient de mieux évaluer le transfert autonome. Ce constat
ne remet pas en cause l'exactitude des réponses vérifiées et ne constitue pas
une validation exhaustive des autres chapitres du catalogue.

## Améliorations après contre-vérification — 28 septembre 2026

Les deux faiblesses ciblées ci-dessus ont été traitées dans les sources CPL et
ATPL, régénérées puis importées dans l’application locale, en français et anglais.
Chaque chapitre conserve ses 15 questions et ses identifiants.

### Dix questions remaniées et cours associés

- **CPL Q7 :** calcul inverse de la masse à déplacer. À 2 400 kg, déplacer le
  CG de 0,10 m en déplaçant une charge de 3 m impose **80 kg**. Les distracteurs
  distinguent distance, masse et variation du moment.
- **CPL Q11 :** consommation de 100 kg à 2,50 m puis ajout de 50 kg à 1 m.
  Depuis 2 300 kg et 5 000 kg·m, le bilan devient 2 250 kg et 4 800 kg·m,
  soit **2,133 m**. Les trois mauvaises réponses correspondent à l’oubli de
  l’ajout, à la conservation de la masse initiale et au mauvais signe du moment
  consommé. Chacun de ces calculs erronés est aussi contrôlé par un test.
- **ATPL Q2/Q3 :** reconstruire les étapes parking/décollage depuis la DOW,
  la charge marchande et le carburant, en traitant correctement le roulage.
- **ATPL Q5/Q6/Q7 :** contrôler simultanément plusieurs plafonds, intégrer une
  limite de performance et calculer le déchargement minimal après ajout de
  carburant. Résultats des trois cas : dépassement de **100 kg** à l’arrivée,
  charge supplémentaire maximale **300 kg**, déchargement minimal **200 kg**.
- **ATPL Q10/Q11 :** combiner déplacement interne et consommation (**4,800 m**),
  puis étudier deux réservoirs dont le carburant consommé a son bras moyen au
  CG initial (**5,000 m**).
- **ATPL Q15 :** calculer d’abord le bras depuis masse et moment, puis convertir
  en pourcentage de MAC : **20 %**.

Les cours expliquent ces méthodes avec des exemples complémentaires et les
hypothèses de consommation inchangée. Cette hypothèse appartient aux exercices ;
elle ne devient pas une règle de planification opérationnelle.

### SVG de méthode

Un nouveau schéma FR/EN explique les variations de masse et de moment, les
déplacements internes et la reconstruction du CG. Il accompagne **13 corrections**
de centrage CPL/ATPL. Il apparaît après la réponse ; les illustrations contenant
les données des énoncés restent distinctes. Les corrections sur les étapes de
masse conservent leur schéma de bilan adapté.

Reproduction des fichiers, y compris les variantes d’unités :

```text
python scripts/illustrations/mass_balance_method.py
node scripts/illustrations/ppl_unit_variants.mjs
python scripts/seed/cpl_content.py
python scripts/seed/atpl_content.py
```

### Validation de ce complément

- **25 tests Python réussis**, dont six nouveaux tests sur les calculs combinés,
  les distracteurs, les choix bilingues et l’association du SVG aux corrections.
- Test de rendu Markdown des illustrations réussi sur tout le catalogue,
  désormais contrôlé en FR et EN avec les schémas propres à chaque correction.
- Deux nouveaux SVG rendus et inspectés visuellement ; leurs six chemins
  FR/EN × métrique/US/britannique répondent HTTP 200 dans l’application locale.
- **Deux quiz, 30 questions, 60 réponses invitées** vérifiés par l’API : énoncés,
  choix, bonnes réponses et corrections FR/EN concordent avec les sources.
- Sauvegarde préalable : `.local/content-before-mass-scenarios-20260928.sql`.
  Identifiants conservés ; comptes, historique, progression et gamification
  inchangés après import et essais.
- Audit de format : 51 couples parcours/matière, aucune erreur signalée.

Cette validation porte sur ce complément. Le rendu complet des pages dans un
navigateur n’a pas été inspecté pendant cette passe. Le catalogue reste partiel
et demande la poursuite du travail pédagogique matière par matière ; ces tests
ne constituent pas une certification de conformité à une banque officielle.
