# Revue de qualité pédagogique — 27 septembre 2026

## Verdict

**La qualité attendue pour des cours complets de préparation n’est pas encore atteinte.**
Les contenus actuels constituent surtout des fiches de révision et des exercices
introductifs. Des tests techniques réussis et quinze questions par chapitre ne
suffisent pas à valider leur profondeur, leur difficulté ou leur représentativité.

## Périmètre réellement vérifié

- Inventaire automatique des 148 chapitres et 2 220 questions FR : taille des
  textes, formats de réponses et présence de SVG dans les énoncés.
- Relecture ciblée des quatre derniers chapitres : `IR_010_001`, `IR_040_001`,
  `CPL_022_001`, `ATPL_031_001` ; contrôle des calculs et de la formulation des questions.
- Échantillon PPL relu : `PPL_COM_001`, `PPL_MET_009`, `PPL_NAV_002`,
  `PPL_OP_003`, `PPL_POF_001`.
- Lecture du générateur partagé, qui fixe uniformément la difficulté et les
  prérequis des chapitres professionnels. Le LAPL reprend le programme PPL ;
  il n’a pas fait l’objet d’une seconde relecture indépendante dans cette passe.
- Les autres chapitres n’ont pas tous été relus ligne par ligne. Aucune validation
  intégrale des 2 220 réponses, aucune mesure de difficulté auprès d’élèves.

## Mesures de départ

Nombre de mots du texte français des concepts, images Markdown exclues ; notices
incluses. Ces mesures signalent la brièveté, sans imposer un seuil artificiel de mots.

| Parcours | Chapitres | Mots min./médiane/max. | Questions | Vrai/faux | Énoncés avec SVG |
|---|---:|---:|---:|---:|---:|
| PPL | 63 | 151 / 172 / 464 | 945 | 530 | 18 |
| LAPL | 63 | 190 / 211 / 503 | 945 | 530 | 18 |
| CPL | 7 | 328 / 365 / 393 | 105 | 0 | 4 |
| ATPL | 7 | 328 / 358 / 403 | 105 | 0 | 5 |
| IR | 8 | 328 / 364 / 416 | 120 | 0 | 4 |

Les textes LAPL sont notamment allongés par leur notice d’adaptation : cette
différence de longueur ne démontre pas une meilleure explication.

## Défauts constatés et actions

| Critère | Constat concret | État après cette passe |
|---|---|---|
| Explications | PPL_MET_009 présente les codes et groupes de changement en quelques paragraphes, sans exercice complet d’exploitation d’une évolution météo. | À approfondir. |
| Progression | Les auteurs professionnels imposent `advanced`, niveau 3, 35 minutes et le même prérequis de calcul, même en communication ou physiologie. | Métadonnées à individualiser ; leur valeur ne mesure pas la difficulté réelle. |
| Objectifs | « Résoudre et vérifier les exercices de… » ne décrit pas précisément une compétence observable ; références souvent limitées à la matière entière. | Rattachement aux objectifs précis à compléter. |
| Distracteurs | IR_040_001 proposait cap nord, pression cabine et vitesse nulle face à une illusion de cabré ; PPL_COM_001 oppose le collationnement à mesurer une vitesse ou calculer une masse. | Huit questions IR reformulées ; défaut encore présent dans d’autres chapitres. |
| Autonomie du QCM | Renvois « même hypothèse », « même panne », « cette impression » et paramètres MAC absents si une question est tirée seule. | Cinq énoncés rendus autonomes dans le dernier lot, dont un également renforcé sur ses distracteurs. |
| Corrections | Beaucoup de corrections répètent la réponse sans expliquer les erreurs typiques. | Raisonnement et erreurs expliqués sur les douze questions retouchées ; généralisation nécessaire. |
| Schémas | Un SVG de synthèse réapparaît dans toutes les corrections d’un chapitre, même quand il ne représente pas le raisonnement de la question. | Illustrations utilisables comme synthèse ; davantage d’annexes et de corrections visuelles spécifiques nécessaires. |
| Évaluation | Nombreuses questions de reconnaissance ou oui/non déguisées en quatre choix. | Ne pas considérer la forme QCM comme preuve d’un niveau examen. |
| Exactitude | Les calculs du dernier lot passent des vérifications indépendantes ; hypothèses SERA explicites dans les questions corrigées. | Constat limité aux points vérifiés, pas certification de tous les cours. |

Les questions corrigées gardent leur position, leur concept et leurs identifiants.
Les résultats historiques ne constituent pas une réévaluation de la nouvelle version.

## Sources utilisées pour la revue ciblée

- [FAA AIM, physiologie et illusions](https://www.faa.gov/air_traffic/publications/atpubs/aim_html/chap8_section_1.html) : mécanismes techniques, sans transposer les règles FAA au cadre EASA.
- [SERA.5015, minima et changement de régime](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-standardised-european?erules-id=ERULES-1963177438-9807).
- [FAA, instruments de vol](https://www.faa.gov/sites/faa.gov/files/regulations_policies/handbooks_manuals/aviation/phak/10_phak_ch8.pdf).
- [FAA, masse et centrage](https://www.faa.gov/sites/faa.gov/files/2023-09/Weight_Balance_Handbook.pdf).
- [EASA, ECQB et objectifs](https://www.easa.europa.eu/en/domains/aircrew-and-medical/european-central-question-bank-ecqb) : la banque officielle est confidentielle. Les questions de l’application sont originales.

## Niveau de qualité à atteindre avant de déclarer un chapitre abouti

1. Objectifs observables, prérequis adaptés et périmètre explicite.
2. Explication du mécanisme, exemple guidé, exercice de transfert et synthèse.
3. Questions autonomes, réponse unique, distracteurs fondés sur des erreurs plausibles.
4. Correction expliquant le raisonnement et les confusions importantes.
5. Schéma ou annexe utile à la tâche, lisible en FR/EN et sans réponse divulguée.
6. Références précises et datées, revue de traduction et validation pédagogique.

La prochaine priorité est d’enrichir et de relire les cours existants selon cette
grille. Ajouter des chapitres courts ne résoudrait pas les défauts constatés.

## Vérification des corrections

Douze questions retouchées dans les quatre chapitres du dernier lot, en FR/EN.
Les dix tests Python et les neuf tests frontend ciblés passent. Après sauvegarde
et import, les 60 questions de ces chapitres ont été vérifiées via API avec une
réponse correcte et une incorrecte, soit 120 soumissions invitées. Identifiants
existants et données personnelles préservés. Cette validation technique confirme
la bonne livraison des corrections ; elle ne remplace pas la revue pédagogique.

## Seconde vérification à la demande de l’utilisateur

Relecture des 60 énoncés et choix du dernier lot dans les JSON effectivement
générés, puis comparaison des énoncés FR/EN avec les imports SQL : concordance.
Les cinq dépendances de contexte précédemment signalées sont corrigées.

La qualité globale reste insuffisante malgré ces corrections :

- `CPL_022_001`, question 13 : couleur et emplacement de l’écran restent des
  distracteurs trop faciles pour évaluer l’indépendance des sources.
- `IR_040_001`, question 13 : « indépendance des capteurs » ne représente pas
  une confusion crédible avec la fixation attentionnelle.
- `ATPL_031_001` : de nombreux items évaluent une seule opération arithmétique ;
  ils servent de préparation, mais ne suffisent pas à établir un niveau avancé.
- Les 22 chapitres CPL/ATPL/IR portent tous `advanced` et 35 minutes. Cette
  uniformité vient du générateur, pas d’une calibration pédagogique individuelle.

Conclusion maintenue : corrections locales confirmées, cours complets et niveau
d’examen non validés. La priorité reste l’approfondissement des contenus et des
évaluations, plutôt qu’une nouvelle augmentation du nombre de questions.
