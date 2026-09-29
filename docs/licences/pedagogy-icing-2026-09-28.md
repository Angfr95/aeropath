# Reprise pédagogique — givrage IR, 28 septembre 2026

## Résultat livré

Le chapitre `IR_050_001` est réécrit en français et anglais et importé localement.
Ses quinze questions conservent leurs identifiants. Il enseigne un raisonnement
sur le givrage ; il ne représente pas à lui seul la météorologie IR complète.

| Étape | Travail demandé à l’élève |
| --- | --- |
| Comprendre | Relier eau liquide, température et exposition ; expliquer l’effet d’une inversion. |
| Interpréter | Distinguer aspect du dépôt, gravité et limites de l’observation. |
| Diagnostiquer | Séparer fonctionnement d’un système, zone protégée et domaine approuvé. |
| Décider | Éliminer une issue incompatible, calculer les temps, actualiser selon les observations. |

Chaque partie comprend une explication et un exemple guidé. Des questions de
rappel invitent à expliquer les limites d’une conclusion. Le cas résolu compare
6 et 18 NM à 120 kt ; le QCM de transfert utilise 8 et 24 NM. La bonne décision
dépend d’abord du minimum d’altitude fourni, puis du calcul du temps.

Les quinze QCM et corrections sont remaniés. Les choix portant sur la couleur
des sièges, le nombre de fenêtres ou l’indicatif radio sont supprimés. Les erreurs
visées concernent notamment la température seule, la généralisation d’une
observation locale, la confusion entre mesure d’IAS et performances, ou la
priorité indue donnée au temps de sortie le plus court.

## Illustrations

Le schéma de synthèse existant reste associé au cours et aux corrections adaptées.
Deux nouvelles illustrations bilingues accompagnent quatre corrections :

- `ir-icing-effects.svg` : profil contaminé, traînée, portance et limites de
  l’interprétation d’une IAS cohérente ou d’une altitude stable (Q9–Q11).
- `ir-icing-exit.svg` : comparaison chiffrée des deux issues du Q14, après réponse.

Les variantes d’unités US/britanniques sont générées. Le SVG contenant la solution
du Q14 reste dans la correction, sans dévoiler la réponse dans l’énoncé.

## Vérifications

- Relecture FR/EN des quatre parties, quinze énoncés, choix et corrections.
- Trois tests nouveaux : décision sous contraintes, données différentes entre
  exemple et transfert, concordance du SVG avec les calculs et absence de fuite.
- **28 tests Python réussis** ; test de rendu des illustrations de toutes les
  corrections du catalogue FR/EN réussi.
- Quatre SVG FR/EN rendus et inspectés ; douze chemins de fichiers (deux schémas,
  deux langues, trois systèmes d’unités) accessibles via HTTP 200.
  Le dernier rendu français, retardé par une indisponibilité de l’approbation
  automatique, a été contrôlé le 29 septembre après harmonisation du libellé
  sur la portance maximale avec l’anglais.
- API locale : **15 questions et 30 réponses invitées** vérifiées ; données
  personnelles inchangées. Sauvegarde préalable dans
  `.local/content-before-icing-pedagogy-20260928.sql`.
- Correspondance des positions et concepts contrôlée avant import, préservant
  les identifiants. Les anciennes réponses ne sont pas réévaluées rétroactivement.

Les références techniques sont enregistrées dans le chapitre, notamment les
mécanismes et effets décrits dans le [guide FAA sur le givrage](https://www.faa.gov/documentLibrary/media/Advisory_Circular/AC_91-74B.pdf).
Les décisions chiffrées sont des exercices originaux avec contraintes explicites.
La durée de 45 minutes reste une estimation, pas une mesure auprès d’élèves.
La difficulté et la capacité de transfert n’ont pas encore été mesurées sur un
groupe d’apprenants ; les tests automatisés contrôlent la cohérence, pas cette
efficacité pédagogique.

[Ouvrir le chapitre local](http://localhost:3000/lessons/IR/meteo/272998f6-0000-4000-8000-0000272998f6).

## Reproduction

```text
python scripts/illustrations/ir_icing_pedagogy.py
node scripts/illustrations/ppl_unit_variants.mjs
python scripts/seed/ir_content.py
python -m unittest discover -s scripts/seed -p 'test_*.py'
node --test web/frontend-react/tests/quizIllustrations.test.mjs
```
