# Audit d’adéquation aux examens PPL(A)

Date : 18 septembre 2026. Périmètre : PPL(A) France / cadre EASA.

> Mise à jour du 22 septembre : voir le [lot de corrections appliquées](ppl-alignment-changes.md).
> Les constats ci-dessous décrivent l’état initial audité, avant ces corrections.

## Conclusion

**Adéquation partielle au programme ; préparation complète et simulation de
l’examen officiel non établies.** Les quiz actuels sont des exercices de révision.
La réussite des tests informatiques ne valide ni les connaissances enseignées ni
la représentativité des questions par rapport à l’examen.

Contrôle effectué : inventaire des 56 fichiers et de la structure des 840 questions,
lecture du moteur de quiz, comparaison avec les modalités françaises publiées,
revue ciblée des anciennes leçons de réglementation et de plusieurs matières.
Ce document ne constitue pas une validation individuelle des 840 réponses ni une
matrice exhaustive de tous les objectifs du programme. Aucune comparaison avec
une banque confidentielle de questions officielles n’a été effectuée.

## 1. Format français à prendre comme référence

La [page DGAC actualisée le 22 janvier 2026](https://www.ecologie.gouv.fr/politiques-publiques/examens-theoriques-devenir-personnel-navigant)
confirme 9 épreuves, 128 questions et 75 % requis dans chaque épreuve.
Les quotas et durées proviennent de la
[note DSAC sur le format applicable depuis le 4 décembre 2023](https://www.ecologie.gouv.fr/sites/default/files/documents/nouveau_format_examens_theoriques_PPL_LAPL_BPL_SPL.pdf).

| Matière | Questions | Minutes | Bonnes réponses minimales, calculées à 75 % |
| --- | ---: | ---: | ---: |
| Droit aérien | 20 | 30 | 15 |
| Performance humaine | 12 | 18 | 9 |
| Météorologie | 20 | 30 | 15 |
| Communications | 8 | 12 | 6 |
| Navigation | 20 | 50 | 15 |
| Principes du vol | 12 | 18 | 9 |
| Procédures opérationnelles | 12 | 18 | 9 |
| Performances et préparation du vol | 12 | 18 | 9 |
| Connaissance générale de l’aéronef | 12 | 18 | 9 |
| **Total** | **128** | **212** | **Chaque matière doit être réussie** |

Les quatre propositions avec une seule réponse correcte sont décrites dans le
[guide candidat DGAC](https://www.ecologie.gouv.fr/sites/default/files/Guide_candidat_examens_ordinateurs.pdf).
Le [tutoriel FFA EOLE](https://faq.ffa-aero.fr/smileexamens/wp-content/uploads/sites/15/2024/06/Tuto_utilisation-interface-examens.pdf)
est également indexé avec quatre propositions ; son téléchargement direct a
échoué pendant cette vérification. Le guide ancien ne sert pas de référence aux
quotas actuels : la note DSAC et la page DGAC de 2026 prévalent pour ces quotas.

## 2. Écarts vérifiés dans l’application

| Point | État constaté | Action nécessaire |
| --- | --- | --- |
| Choix de réponses | 310 questions à 4 choix ; 530 vrai/faux à 2 choix (63,1 %) | Réécrire les questions destinées à l’examen blanc avec quatre choix plausibles et une seule réponse exacte. |
| Nombre de questions | `internal/learning/engine.go` fixe quiz et examens par thème à 15 | Garder les révisions de 15 ; prévoir les quotas propres aux neuf épreuves pour les examens blancs. |
| Temps | Aucun chronomètre dans `Quiz.tsx` / `quizStore.ts` | Appliquer une durée par épreuve en mode examen. |
| Correction | Réponse et explication immédiatement après chaque choix | Réserver les corrections à la fin en mode examen. |
| Résultats | Score de session ; pictogramme favorable à partir de 80 % | Une décoration à 80 % n’est pas un seuil réglementaire. Afficher une réussite calculée à 75 % par matière pour l’examen blanc. |
| Organisation | 14 dossiers de contenu, sans correspondance exhaustive aux objectifs officiels | Rattacher chaque notion et question à une matière et à un objectif précis. |
| SVG | 840 corrections illustrées, mais 18 énoncés avec schéma propre | Les schémas de correction ne remplacent pas des exercices de lecture de cartes, tableaux et graphiques. |

Les préférences de langue et d’unités modifient la présentation. Elles ne
transforment pas cette préparation EASA/France en préparation FAA ou UK CAA.

## 3. Erreurs et affirmations à reprendre en priorité

### Erreurs confirmées

- `PPL_AIRLAW_001`, question 1 : « Quand l’ICAO a-t-elle été créée ? » donne
  **1944**. Il faut distinguer la Convention de Chicago de 1944 et l’OACI entrée
  en existence le **4 avril 1947**, avec la phase provisoire intermédiaire.
  [Historique officiel OACI](https://www.icao.int/history-icao-and-chicago-convention).
- `PPL_AIRLAW_002`, concept 3 et corrections : les tableaux indiquent **20 h solo**
  comme minimum EASA et FAA. Le parcours PPL(A) standard de FCL.210.A prévoit
  **10 h solo supervisées**, au sein des 45 h de formation, avec conditions
  supplémentaires. Le parcours FAA avion monomoteur Part 61 prévoit également
  **10 h solo**, et 20 h d’instruction, au sein des 40 h.
  [FCL.210.A, texte EUR-Lex](https://eur-lex.europa.eu/legal-content/EN/TXT/?uri=CELLAR%3A6a7f4549-5f88-11ed-92ed-01aa75ed71a1),
  [14 CFR 61.109(a), eCFR](https://www.ecfr.gov/current/title-14/chapter-I/subchapter-D/part-61/subpart-E/section-61.109).
- `PPL_AIRLAW_002`, question 4 : la réponse « 27 (UE) » assimile l’ensemble
  EASA à l’UE. L’EASA liste **31 États membres**. Reformuler précisément le
  périmètre et les conditions de reconnaissance au lieu d’une réponse générale
  sur tous les usages possibles d’une licence.
  [Liste officielle EASA](https://www.easa.europa.eu/en/taxonomy/term/1255/all).

### Affirmations non validées, à suspendre avant utilisation en examen blanc

- `PPL_AIRLAW_003`, concept 3 et question 10 : validité de l’immatriculation
  « 3 ans (France/Europe) ». La généralisation n’est pas étayée par une référence
  applicable. Vérifier le régime national et le document exact ; ne pas confondre
  immatriculation, navigabilité et examen de navigabilité.
  [Procédures françaises d’immatriculation](https://www.ecologie.gouv.fr/politiques-publiques/immatriculation-aeronefs).
- Même leçon : licence de station « 5 ans en Europe », certificat affiché dans
  le cockpit et assurance automatiquement non valable. Assertions trop générales
  pour être retenues telles quelles sans conditions et références précises.
- `PPL_AIRLAW_002` : conversions, délais, coûts et « validité licence / recentrage »
  mélangent plusieurs régimes. Distinguer licence, qualification, expérience
  récente, médical, validation et conversion avant de réécrire ces passages.

Les erreurs du français sont aussi présentes dans les traductions anglaises
examinées. Les corrections devront porter sur les cours, réponses et explications
des deux langues, en conservant les identifiants existants.

## 4. Couverture pédagogique insuffisante

Le [programme EASA AMC1 FCL.210 / FCL.215](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=LAPL&page=5)
demande notamment la stabilité statique/dynamique, la physiologie respiratoire,
les effets du monoxyde de carbone, l’hyperventilation et la vision nocturne.
La version HTML consultable porte la révision août 2023 ; cette consultation ne
constitue pas à elle seule une vérification de toutes les modifications ultérieures.

Constats sur nos contenus, issus de la lecture ciblée et d’une recherche dans
tous les textes de concepts :

- **Principes du vol** : seulement deux leçons introductives. Pas de développement
  identifié sur stabilité statique/dynamique ou effet de sol. La traînée est
  expliquée qualitativement, sans parcours complet de lecture de courbes.
- **Performance humaine** : deux leçons sur aptitude et attention. Hypoxie
  brièvement définie ; pas de développement identifié sur monoxyde de carbone,
  hyperventilation ou vision nocturne.
- **Droit aérien** : `PPL_AIRLAW_005` renvoie aux minima SERA sans les enseigner
  numériquement ni entraîner leur application selon les cas.
- **Navigation** : `PPL_NAV_003` introduit la combinaison vectorielle air/vent,
  mais son exercice travaille essentiellement distance, vitesse et temps.
  Développer la résolution de dérive et vitesse sol et les exercices sur carte.
- **Masse et performances** : `PPL_MB_001` calcule un centrage simplifié mais
  demande de consulter l’enveloppe sans fournir un exercice complet de lecture
  de cette enveloppe. Approfondir les tableaux et scénarios à plusieurs étapes.

Une recherche de mots ne démontre pas à elle seule une absence ; ces constats
doivent guider une matrice détaillée de couverture, pas un pourcentage de
« conformité » inventé. La présence des neuf grandes matières ne suffit pas.

## 5. Ordre de correction

1. Corriger les erreurs confirmées et vérifier les anciennes leçons réglementaires.
2. Associer chaque concept et question aux objectifs du programme en vigueur,
   en distinguant couvert, partiel et absent.
3. Compléter les connaissances manquantes et rédiger des QCM d’application à
   quatre choix ; la quantité seule ne valide pas la difficulté ni la pertinence.
4. Ajouter un mode examen blanc avec quotas, temps et notation par matière.
5. Revoir les traductions, les SVG et les conversions d’unités sur les exercices
   corrigés, puis faire relire pédagogiquement les contenus.

Statut à la clôture de cet audit : écarts documentés ; contenu et moteur de
l’application non corrigés dans cette passe de vérification.
