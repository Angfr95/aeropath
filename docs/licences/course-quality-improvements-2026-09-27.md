# Qualité des cours : corrections et état réel du catalogue

> État du premier lot de corrections. Le travail se poursuit en priorité FR/EN
> avec le [lot masse et centrage](course-quality-mass-and-balance-2026-09-27.md).
> Les autres langues ont été reportées à la demande de l'utilisateur.

## Conclusion

Le niveau de qualité demandé n'est **pas encore atteint par tout le catalogue**.
La vérification structurelle couvre 148 chapitres et 2 220 questions ; elle ne
constitue pas une validation technique et pédagogique exhaustive de leur texte.
Les corrections ci-dessous sont réalisées, générées et importées dans la base
locale. Les autres lacunes restent explicitement ouvertes.

## Corrections appliquées

### PPL et LAPL : supprimer l'indice donné par l'ordre des choix

Les 530 vrai/faux de chaque parcours utilisent maintenant l'ordre
`Vrai, Faux` / `True, False`, quelle que soit la réponse correcte. La permutation
est faite dans la source et propagée à l'import ; le correcteur conserve la même
identité de réponse entre langues. Les générateurs refusent désormais un ordre
non conforme pour ces questions.

La stratégie « toujours la première réponse » passe de 487/530 (91,89 %) à
266/530 (50,19 %) dans chaque parcours. Les réponses vraies et fausses sont
équilibrées : 266 et 264. Les vrai/faux restent séparés de l'examen blanc PPL.

### IR_022_001 : reprise du cours et des quinze questions FR/EN

- Objectifs et prérequis propres au sujet.
- Distinction des références QNH, QFE et standard avec un exemple chiffré.
- Méthode signée d'erreur de calage, exemple résolu, contrôle inverse et
  application autonome utilisant de nouvelles valeurs.
- Séparation entre erreur de calage, effet du froid et marge terrain.
- Explication de la statique piégée **par rapport à l'altitude de blocage**,
  avec le contre-exemple 5 000 → 7 000 → 6 000 ft.
- Énoncés autonomes et corrections expliquant les erreurs de raisonnement.

Le texte des concepts atteint 789 mots FR et 716 EN, contre environ 350 mots
auparavant. Cette longueur est une mesure descriptive, pas une note de qualité.
La durée de 45 minutes reste une estimation à confronter à des utilisateurs.

Références techniques : [FAA PHAK, Flight Instruments](https://www.faa.gov/sites/faa.gov/files/regulations_policies/handbooks_manuals/aviation/phak/10_phak_ch8.pdf),
[FAA AIM, erreurs altimétriques](https://www.faa.gov/air_traffic/publications/aim_html/chap7_section_2.html).
Les procédures américaines ne remplacent pas les procédures EASA du parcours.

### CPL_032_001 : méthode complète et corrections exploitables

- Lecture guidée des conditions et choix de la bonne ligne du tableau.
- Interpolation en deux étapes avec contrôle dans l'ordre inverse.
- Application autonome à 500 ft / 15 °C : 564,375 m, et non une moyenne systématique.
- Distinction entre interpolation, facteurs successifs et règle d'arrondi.
- Raisonnement sur les marges : 72 m, 9 % de la distance disponible et environ
  9,89 % de la distance requise ne désignent pas la même base de pourcentage.
- Distracteurs revus dans sept questions et explications développées dans les
  quinze corrections.
- Objectifs et prérequis spécifiques ; durée estimée de 45 minutes.

Le SVG commun n'impose plus une cible à 1 000 ft lorsque l'énoncé demande 0 ft.
Il renvoie aux conditions de l'énoncé. Les versions FR, EN, US et impériales
sont régénérées : six fichiers pour cette illustration.
Le texte des concepts atteint 910 mots FR et 843 EN.

### Autres ambiguïtés corrigées

- CPL_081_001 Q3 : virage explicitement coordonné, stabilisé et en palier.
- CPL_081_001 Q6 : vitesse **de décrochage**, masse et configuration fixées ;
  explication de l'erreur consistant à multiplier par n au lieu de √n.
- ATPL_033_001 Q8 : autonomie hors réserves, débit identique et constant,
  vitesses sol constantes, exclusion des temps de virage/montée/descente ;
  contrôle du résultat 288 NM par les temps 2,4 h + 1,6 h.

Au total, 33 questions ont reçu une révision d'énoncé, de choix ou de correction
sur ces quatre chapitres, en plus de la normalisation des vrai/faux.

## Vérifications réalisées

- Génération et validation des cinq parcours ; audit des 51 couples
  parcours/matière sans erreur de format détectée.
- 14 tests Python : contenu interlicences, PPL et nouvelles régressions.
- 9 tests frontend sur quiz et illustrations.
- 44 SVG professionnels FR/EN rendus, sans dépassement ou chevauchement détecté.
  Inspection visuelle du tableau CPL corrigé ; les six versions de ce tableau
  sont vérifiées par le test de régression. 664 variantes d'unités régénérées.
- Sauvegarde de la base avant import dans
  `.local/content-before-quality-strengthening-20260927.sql`.
- Identifiants des 148 leçons et des 2 220 questions conservés.
- 148 quiz de quinze questions comparés aux sources FR/EN par l'API locale :
  4 440 réponses correctes/incorrectes vérifiées en mode invité.
- Empreintes des tables de comptes, historiques, progression et gamification
  identiques avant/après import et après les essais invités.

L'application locale a été démarrée. Aucun navigateur n'était disponible via
Browser : la présentation complète des leçons dans le navigateur n'a pas été
revérifiée lors de cette passe. Un rendu SVG et un test d'API ne prouvent pas
à eux seuls la qualité de lecture de toute la page.

## Lacunes à traiter sur tout le catalogue

L'[inventaire par chapitre](course-quality-inventory.md) et son
[détail JSON](course-quality-inventory.json) sont produits par
`python scripts/seed/audit_course_quality.py`.

| Signal observé | Chapitres concernés | Interprétation |
| --- | ---: | --- |
| Texte FR inférieur à 500 mots | 146 | Relire la profondeur et la progression ; ne pas ajouter du remplissage |
| Au moins une explication de moins de 12 mots FR ou EN | 146 | Vérifier que l'élève comprend sa faute ; certaines corrections de calcul simples peuvent rester courtes |
| Objectif professionnel générique | 20 | Définir des acquis précis et observables |
| Prérequis professionnels génériques | 20 | Relier le chapitre aux connaissances réellement nécessaires |
| Contenu espagnol incomplet dans les sources | 148 | Traduction et validation sémantique encore nécessaires |

La base locale confirme zéro question avec énoncé ou choix espagnols renseignés
dans chacun des cinq parcours. La présence d'un sélecteur de langue ou d'un
repli français n'est pas une traduction.

Les cinq alertes de dépendance au contexte ont été relues : « avion précédent »
en PPL/LAPL OP_002/003 désigne un autre avion et non une question précédente ;
CPL_033_001 Q11 redonne les deux valeurs utiles. Aucune nouvelle dépendance
bloquante n'est établie dans ces cinq candidats. Le relevé conserve les alertes
brutes pour rendre les limites de la détection automatique visibles.

La couverture reste partielle : CPL présente 7 matières sur 13, ATPL 6 sur 13,
IR 7 sur 7 mais avec seulement huit chapitres. PPL/LAPL présentent leurs neuf
matières ; cela ne prouve pas la couverture de tous les objectifs du programme.

## Critères pour considérer un chapitre réellement abouti

1. **Exactitude :** source précise, hypothèses et limites explicites, calculs
   reproduits indépendamment, distinction entre règle et modèle pédagogique.
2. **Progression :** prérequis utiles, explication du mécanisme, exemple résolu,
   exercice guidé puis cas nouveau à résoudre sans recopier les nombres.
3. **Évaluation :** une seule réponse défendable sous les hypothèses énoncées,
   distracteurs issus d'erreurs plausibles et correction qui explique pourquoi.
4. **Illustration :** données cohérentes avec la question, légendes et unités
   lisibles, schéma adapté à la notion, absence de réponse involontairement révélée.
5. **Langues et unités :** même sens et mêmes conditions dans chaque version,
   choix alignés et conversions sans ambiguïté sur le régime réglementaire.
6. **Couverture et usage :** correspondance détaillée aux objectifs publics,
   relecture technique compétente, mise à l'essai avec des élèves et révision
   des questions ambiguës ou trop faciles après observation des réponses.

Les deux chapitres remaniés progressent sur plusieurs de ces critères, sans
être déclarés intégralement validés. Les autres chapitres doivent être repris
sujet par sujet. Aucun contrôle de présence de champs ni seuil de mots ne
remplace cette revue.

L'EASA décrit elle-même une validation en plusieurs étapes avec des experts
pour l'ECQB ; sa banque n'est pas publique. L'alignement d'Aeropath doit donc
être établi à partir du programme et des objectifs publics, sans prétendre
reproduire ou valider la banque officielle.
[Source EASA](https://www.easa.europa.eu/en/domains/aircrew-and-medical/european-central-question-bank-ecqb).
