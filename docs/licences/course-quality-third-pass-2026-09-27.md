# Troisième passe de qualité — 27 septembre 2026

> État historique de la vérification. Les deux défauts ci-dessous sont maintenant
> corrigés et importés ; voir le [bilan actualisé](course-quality-improvements-2026-09-27.md).

## Périmètre et méthode

Relecture ciblée de neuf chapitres PPL (135 questions et leurs concepts) :
PPL_AC_004, PPL_FP_001, PPL_HP_001, PPL_HP_002, PPL_HP_003,
PPL_MET_002, PPL_NAV_004, PPL_PERF_001 et PPL_PERF_002.
Aucune nouvelle erreur arithmétique relevée dans cet échantillon ; cela ne
valide pas intégralement les cours.

Comptage exhaustif de la position de la bonne réponse dans les questions à
deux choix des 63 chapitres PPL et des 63 chapitres LAPL. Inspection du chemin
des choix depuis le moteur de quiz jusqu'à leur affichage. Recoupement du
premier énoncé CPL_032_001 avec son SVG FR et EN.

Les constats ci-dessous sont **confirmés dans les fichiers et le code, mais
non corrigés par cette passe**. Aucun cours, import SQL ou code applicatif
n'a été modifié. Les requêtes de contrôle vers localhost:3000 ont échoué
(connexion refusée) : aucune validation de l'application en fonctionnement
ni de l'état actuel de sa base n'est revendiquée.

## 1. PPL/LAPL : la position révèle presque toujours la réponse des vrai/faux

| Parcours | Questions à deux choix | Bonne réponse en premier | En second |
| --- | ---: | ---: | ---: |
| PPL | 530 | 487 | 43 |
| LAPL | 530 | 487 | 43 |

Choisir systématiquement la première proposition donne donc **91,89 %**
de bonnes réponses sur ces questions. Il s'agit du sous-ensemble vrai/faux,
pas du score de l'ensemble des quiz ni de l'examen blanc PPL.

Le défaut est exploitable dans le chemin de code actuel :

- `internal/learning/engine.go`, `LessonQuiz` : mélange des questions,
  sans permutation des options.
- `internal/transport/http/learning.go`, `LessonQuizHandler` : transmet
  directement `q.Options`, `q.OptionsEn` et `q.OptionsEs`.
- `web/frontend-react/src/utils/learningLocale.ts`, `localizedOptions` :
  conserve les positions entre langues.
- `web/frontend-react/src/pages/Quiz.tsx` : affiche les choix dans cet ordre.

L'examen blanc PPL exige quatre options et le tag `exam-mcq` dans
`internal/learning/ppl_exam.go` ; ces vrai/faux en sont exclus.

**Correction proposée :** normaliser l'ordre des vrai/faux indépendamment
de la réponse correcte, en conservant le même sens à chaque position dans
toutes les langues. Vérifier ensuite la distribution vrai/faux elle-même
et le correcteur. Un mélange uniquement à l'affichage désynchroniserait les
lettres envoyées au serveur et ne constitue pas une correction suffisante.

## 2. CPL_032_001 : contradiction entre le premier énoncé et son SVG

La question 1 demande l'interpolation à **0 ft / 20 °C** :
500 m à 10 °C et 600 m à 30 °C, soit **550 m**.

L'image jointe porte pourtant « Point à calculer : **1 000 ft / 20 °C** ».
Avec les données du tableau, cette autre cible vaut **632,5 m** :
la moyenne de 550 m à 0 ft et de 715 m à 2 000 ft.
Cette cible correspond à la question 3, pas à la question 1 illustrée.

Le conflit existe en français et en anglais dans :

- `content/fr/cpl/performance/CPL_032_001_application.json`, question 1 ;
- `scripts/illustrations/cpl_illustrations.py`, fonction `performance` ;
- `web/frontend-react/public/illustrations/cpl-032-001.svg` ;
- `web/frontend-react/public/illustrations/en/cpl-032-001.svg`.

**Correction proposée :** conserver un tableau commun sans cible imposée
dans le SVG partagé, ou fournir une image propre à chaque cible. Régénérer
les variantes d'unités après modification de la source. La réponse 550 m
du premier énoncé est correcte ; c'est l'instruction de l'image qui diverge.

## Conclusion

Deux nouveaux défauts pédagogiques concrets sont confirmés. Le premier
affaiblit l'évaluation des révisions PPL/LAPL ; le second rend un exercice
CPL contradictoire. Ils s'ajoutent aux conditions manquantes signalées dans
la [deuxième passe](course-quality-second-pass-2026-09-27.md).
Ces vérifications ne permettent toujours pas de qualifier l'ensemble du
catalogue de cours complets ni de préparation entièrement validée aux examens.
