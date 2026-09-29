# Corrections pédagogiques — 29 septembre 2026

Les sept chapitres de la [relecture](content-reading-review-2026-09-29.md) ont été corrigés en français et en anglais. Les variantes qui partagent ces enseignements ont reçu les mêmes corrections : trois chapitres LAPL et les communications CPL/IR. Cela représente douze chapitres publiés, avec quinze questions chacun.

## Ce que l'élève apprend désormais

| Chapitre relu | Lacune corrigée | Application proposée |
|---|---|---|
| PPL_COM_001 | Appel initial, indicatif abrégé, éléments à collationner, expressions et portée des autorisations enseignés explicitement | Comparer des collationnements plausibles, repérer une omission ou une valeur substituée, traiter un destinataire ambigu |
| PPL_COM_002 | Diagnostic par indices, branche VFR de panne radio, différence 7600/7601, signaux visuels et messages MAYDAY/PAN PAN | Distinguer émission/réception, sélectionner un terrain approprié, interpréter couleur/cadence/situation, compléter un message et réévaluer sa catégorie |
| PPL_MET_009 | METAR complet expliqué, critères CAVOK, bases des nuages et références, validité, BECMG/TEMPO/PROB, lecture d'un avertissement en route | Décoder un autre bulletin, croiser heure d'arrivée et évolution, comparer des seuils donnés, examiner temps/zone/niveau d'un SIGMET |
| ATPL_090_001 | Suivi de messages successifs, omissions et substitutions, corrections partielles et conditions | Reconstruire cap/niveau/vitesse après trois messages dont un pour un autre avion ; ne pas transférer sa vitesse au bon destinataire |
| CPL_022_001 | Application différente de l'exemple de panne statique et nouvelle architecture de sources | Expliquer une surindication alors que l'avion remonte mais reste sous l'altitude de blocage ; tracer les dépendances Pt/Ps d'une autre installation |
| IR_050_001 | Comparaison sous plusieurs contraintes et traitement d'une performance inconnue | Écarter une sortie sous le minimum et une autre dépassant le carburant disponible, puis calculer le temps de l'option admissible |
| ATPL_031_001 | Dossier chiffré distinct du cours, lecture d'enveloppe et contrôle d'états intermédiaires | Interpoler une limite avant, contrôler une séquence de consommation qui sort puis revient dans l'enveloppe, distinguer conformité du centrage et autres exigences |

## Évaluation et corrections

- **80 énoncés réécrits** parmi les 105 questions des sept chapitres originaux. Les autres conservent un rôle de contrôle des bases ou d'application pertinent.
- Les **30 vrai/faux** des trois chapitres PPL ont été remplacés par des QCM à quatre choix. Leur intérêt vient des décisions et erreurs proposées, pas du seul nombre de choix.
- Les corrections nouvelles expliquent les confusions visées : omission contre substitution, heure d'observation contre réception, provenance du vent, maintien du régime après TEMPO, référence de pression, masse contre moment, marge limitante et contrôle intermédiaire.
- Les valeurs nouvelles des exercices ne sont pas de simples corrections du résultat affiché : les cas de centrage et de givrage ajoutent des contraintes à traiter ; l'instrumentation inverse la situation verticale et change les dépendances des sources.
- Les exemples et enveloppes restent explicitement fictifs. Les seuils météorologiques donnés dans les exercices ne sont pas présentés comme des minima réglementaires universels.

### Exemples de résultats recalculés

- Bilan ATPL : parking 43 500 kg, décollage 43 350 kg, arrivée 37 850 kg ; avec les quatre plafonds de Q6, ajout maximal de charge **450 kg**.
- Enveloppe fictive : à 10 000 kg, limite avant **20 %MAC**, donc 19 % est trop avant. La séquence de Q13 donne **30 % → 44,44 % → 31,25 %** : seul l'état intermédiaire sort de l'enveloppe.
- Sortie givrage : A viole le minimum ; C consomme 10 kg pour 8 kg disponibles ; B satisfait les contraintes données en **9 minutes**.
- Architecture instrumentale du QCM : A reçoit Pt1/Ps1, B Pt1/Ps2, C Pt2/Ps2 ; une obstruction de Ps2 peut affecter **B et C**, malgré leurs pitots différents. L'exemple guidé du cours utilise D/E/F et d'autres associations.

## Illustrations

Sept dessins de référence ont été repris en FR/EN, puis déclinés dans les sections partageant le contenu : échange radio corrigé, signaux lumineux, frise TAF, suivi d'instructions, dépendances des capteurs, tableau de sorties et enveloppe graduée. Les dessins joints aux énoncés professionnels donnent les données, sans afficher les réponses calculées.

Les quatorze rendus FR/EN de référence ont été inspectés. La légende de l'enveloppe précise que les axes sont gradués. Les variantes US et britanniques conservent les valeurs sources des calculs et ajoutent les équivalences disponibles, comme les autres supports de l'application.

## Références utilisées pour la reprise

- [SERA/France, version du 17 janvier 2026](https://www.ecologie.gouv.fr/sites/default/files/documents/sera_complet_version17012026.pdf) : collationnement, indicatifs, panne radio, transmission en l'air, urgence/détresse et signaux.
- [Met Office — GetMet](https://www.metoffice.gov.uk/binaries/content/assets/metofficegovuk/pdf/services/transport/aviation/training-resources/02413-getmet-update-v6.pdf) : lecture METAR/TAF et critères CAVOK.
- [Météo-France — guide aviation](https://meteofrance.fr/sites/meteofrance.fr/files/files/editorial/guide_aviation_2021.pdf) : sens et durée des fluctuations TEMPO.
- [Met Office — SIGMET](https://www.metoffice.gov.uk/services/transport/aviation/regulated/training-resources-for-aviation/sigmets) : avertissements de phénomènes significatifs.
- [FAA — Weight and Balance Handbook](https://www.faa.gov/sites/faa.gov/files/2023-09/Weight_Balance_Handbook.pdf) : principes techniques de masse, moments et enveloppes. Les références techniques des chapitres instrumentation et givrage ont été conservées.

La relecture juge l'explication, la progression et les erreurs évaluées. Elle ne fournit pas de mesure de difficulté auprès d'élèves réels ni de validation intégrale des programmes de licences. Ces chapitres restent des parties de matières ; l'entraînement oral radio reste nécessaire en complément du support écrit.

## Livraison et conservation du contenu

Contenus importés dans la base locale après sauvegarde. Les douze leçons et leurs 180 questions ont été vérifiées via l'application, avec réponses correctes et incorrectes et corrections FR/EN. Les identifiants existants, comptes, résultats et progressions sont conservés. Les vérifications de calcul et d'affichage complètent la relecture ; elles ne sont pas utilisées comme preuve du niveau pédagogique.

Sources éditoriales : les trois JSON PPL sont les contenus de référence ; les contenus professionnels restent dans leurs sections existantes sous `scripts/seed`. Pour régénérer les dessins après les générateurs généraux, terminer par `python scripts/illustrations/reviewed_chapter_diagrams.py`, puis `node scripts/illustrations/ppl_unit_variants.mjs`. Cet ordre conserve les nouvelles illustrations.
