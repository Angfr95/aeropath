# Reprise pédagogique — instrumentation CPL, 29 septembre 2026

## Ce que l’élève travaille

Le chapitre `CPL_022_001` conserve sa place dans le parcours et ses quinze
questions. Le cours FR/EN développe une méthode : écrire les pressions reçues,
les comparer aux pressions correctes, prévoir les indications et évaluer les
limites du diagnostic.

- Exemples résolus, calculs de contrôle et questions de rappel dans les quatre parties.
- Distinction entre entrée pitot bouchée avec drainage libre et pression totale piégée.
- Contre-exemple avec **deux pressions piégées** : indication constante malgré
  une variation de vitesse réelle dans le modèle idéal.
- Contre-exemple de **descente restant au-dessus de l’altitude d’obstruction** :
  une sous-indication peut persister ; le sens vertical seul ne détermine pas l’erreur.
- Statique de secours : raisonner avec la pression effectivement fournie,
  sans supposer un sens universel de correction sur tous les avions.
- Chaîne prise/conduite/calculateur/écran : distinguer indépendance électrique,
  informatique et pneumatique ; accord de deux écrans ne vaut pas preuve d’exactitude.
- IAS/TAS/GS : un écart numérique ne mesure pas directement une erreur instrumentale.

Les quinze QCM et leurs corrections sont remaniés en FR/EN. Les distracteurs
visent des erreurs identifiables : addition ou inversion des pressions, confusion
entre différence et variation, transfert abusif d’un modèle de panne à un autre,
vote majoritaire de sources communes et substitution de GS à IAS.

## Images et limites du modèle

Le schéma de synthèse reste dans la leçon. Deux nouveaux SVG FR/EN accompagnent
neuf corrections : `cpl-pitot-reasoning.svg` et `cpl-sensor-independence.svg`.
Ils expliquent le raisonnement après réponse. Les variantes US/britanniques sont générées.

Les calculs utilisent un modèle idéal sans fuite ni effet thermique, avec calage
fixe sauf indication contraire. Les nombres sont des exercices originaux ; ils
ne fournissent pas de procédure de vitesse non fiable propre à un avion.
La référence technique existante reste le chapitre
[Flight Instruments du FAA PHAK](https://www.faa.gov/sites/faa.gov/files/regulations_policies/handbooks_manuals/aviation/phak/10_phak_ch8.pdf).
Le cours reste dans le cadre du parcours CPL EASA.

## Contrôles

- **33 tests Python réussis**, dont cinq nouveaux contrôles sur les modèles de
  pression, les réponses FR/EN et les images de correction.
- Test de rendu des illustrations de tout le catalogue FR/EN réussi.
- Quatre SVG rendus et inspectés visuellement ; espacement du libellé commun A/B ajusté.
- Positions, concepts et identifiants des quinze questions conservés.
- Après relance des services, sauvegarde puis import local :
  `.local/content-before-instrument-pedagogy-20260929.sql`.
- **15 questions et 30 réponses invitées** contrôlées via API : énoncés, choix,
  réponses et corrections FR/EN concordants ; comptes, résultats et progression inchangés.
- Les douze chemins de SVG FR/EN × métrique/US/britannique répondent HTTP 200.
- Durée de 45 minutes estimée. La difficulté n’est pas encore calibrée auprès
  d’élèves ; ces vérifications démontrent la cohérence des cas, pas une validation
  pédagogique externe de toute l’instrumentation CPL.

[Ouvrir le chapitre local](http://localhost:3000/lessons/CPL/instrumentation/1ad7b73f-0000-4000-8000-00001ad7b73f).

## Reproduction

```text
python scripts/illustrations/cpl_instrument_pedagogy.py
node scripts/illustrations/ppl_unit_variants.mjs
python scripts/seed/cpl_content.py
python -m unittest discover -s scripts/seed -p 'test_*.py'
node --test web/frontend-react/tests/quizIllustrations.test.mjs
```
