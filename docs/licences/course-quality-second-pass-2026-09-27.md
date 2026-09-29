# Nouvelle passe de vérification — 27 septembre 2026

> État historique. Les défauts des sections 1 à 3 ont depuis été corrigés dans
> les sources et la base locale ; voir le [bilan actualisé](course-quality-improvements-2026-09-27.md).

## Périmètre

Lecture des 270 couples énoncé/réponse des 18 autres chapitres CPL/ATPL/IR,
après les quatre chapitres relus précédemment. Les trois chapitres de communication
partagent le même texte : cela représente 240 couples textuellement distincts.
Recoupement ciblé avec les concepts et traductions des chapitres signalés,
et lecture des textes des 13 SVG utilisés dans les énoncés professionnels/IR.
Contrôle de présence des traductions des 2 220 questions de tous les parcours.

Ce périmètre ne constitue pas une relecture intégrale des 2 220 questions ni
une comparaison avec l’ECQB. Cette passe documente les défauts ; elle ne modifie
pas les cours ou les réponses.

## Nouveaux constats précis

### 1. IR_022_001 — condition physique incomplète, priorité haute

Le quatrième concept et les questions 12–13 associent directement montée à
sous-indication et descente à surindication, avec statique piégée et pitot libre.
Il faut comparer l’altitude actuelle à **l’altitude de l’obstruction**.

Contre-exemple : obstruction à 5 000 ft, montée jusqu’à 7 000 ft puis descente
à 6 000 ft. L’avion descend mais reste au-dessus de l’altitude d’obstruction :
la pression statique piégée reste supérieure à la pression extérieure et
l’anémomètre sous-indique encore par rapport à une statique correcte.
La réponse « surindication en descente » ne vaut donc pas dans tous les cas.

Source : [FAA PHAK, chapitre 8, circuit statique obstrué](https://www.faa.gov/sites/faa.gov/files/regulations_policies/handbooks_manuals/aviation/phak/10_phak_ch8.pdf).
La formulation doit être corrigée dans les concepts, énoncés et corrections FR/EN.
La question 11 conserve également « même obstruction », sans préciser dans
son propre énoncé que la pression statique est entièrement piégée.

### 2. CPL_081_001 — grandeur demandée insuffisamment définie

Question 6 : « Vs à 1 g = 50 kt. À n = 1,414, vitesse approximative ? ».
La réponse 59,5 kt est celle de la **vitesse de décrochage**, à masse et
configuration constantes. Le facteur de charge seul ne détermine pas la vitesse
réelle de l’avion. Il faut nommer la grandeur et rappeler les hypothèses.
La version anglaise « approximate speed » présente la même omission.

Question 3 : le virage devrait être explicitement coordonné et stabilisé, comme
le cours le précise, pour appliquer directement `n = 1/cos φ`.

### 3. ATPL_033_001 — hypothèses laissées dans le cours

Question 8 : le modèle de point de retour suppose une autonomie dédiée à
l’aller-retour, un débit identique et constant et des vitesses sol constantes.
Le cours précise ces hypothèses, mais l’énoncé isolé ne les reprend pas toutes.
La réponse 288 NM est correcte sous ce modèle ; le risque est sa généralisation
à une situation opérationnelle différente. Il faut expliciter le modèle dans l’énoncé.

### 4. Aides dans les annexes — à distinguer d’une évaluation sans aide

Les SVG des énoncés ATPL_022_001 et CPL_081_001 fournissent les formules utilisées.
Ce sont donc des exercices d’application guidée. Aucune réponse numérique directe
au calcul demandé n’a été identifiée dans les textes des 13 SVG examinés.
Ce constat porte sur les textes des fichiers, pas sur un nouveau rendu visuel complet.
Le statut d’exercice guidé doit être explicite ; ces items n’évaluent pas seuls
la capacité à choisir ou retrouver la relation correcte.

## Points vérifiés sans anomalie nouvelle

- Résultats numériques examinés : interpolation 632,5 m puis 728 m avec facteur,
  ETP 360 NM, point de retour 288 NM, EAS/TAS 200/250 kt, géométrie DME 3–4–5,
  taux de descente 636 et 795 ft/min : calculs cohérents avec les données fournies.
- Les 2 220 questions possèdent les champs anglais d’énoncé, choix et correction.
  Présence des champs ne signifie pas validation sémantique de toutes les traductions.
- Le renvoi « même trajet » dans CPL_033_001 Q11 reste calculable isolément :
  les deux consommations 72 et 60 L figurent dans l’énoncé. Ce n’est donc pas
  automatiquement une dépendance bloquante.

## Conclusion de cette passe

Un défaut de généralisation physique, une grandeur ambiguë et des hypothèses
insuffisamment explicites sont confirmés dans des chapitres différents du dernier
lot. Les corrections précédentes ne permettent donc pas de déclarer l’ensemble
du catalogue validé. Traiter d’abord IR_022_001, puis les énoncés CPL_081_001 et
ATPL_033_001 avant de considérer ces questions comme évaluations autonomes.
