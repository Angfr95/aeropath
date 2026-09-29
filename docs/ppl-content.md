# Contenu PPL(A)

Les premiers chapitres LAPL, CPL, ATPL et IR sont documentés dans
[Organisation des parcours PPL, LAPL, CPL, ATPL et IR](licences/README.md).

Le catalogue contient **63 leçons et 945 questions en français et en anglais** :

- 3 leçons de réglementation préexistantes, conservées avec leurs 45 questions ;
- 39 chapitres complétés, dont le fichier auparavant vide sur les espaces aériens ;
- 8 nouvelles leçons : principes du vol, facteurs humains, communications, carburant et centrage.
- Un deuxième lot de 6 leçons : performances (2), instruments (2), déroutement (1)
  et clairances radio (1), avec 30 nouvelles questions.

Les 47 leçons du premier lot comprenaient chacune trois sections et cinq questions,
après l'ajout de 94 questions d'application le 16 septembre 2026.
Les 6 leçons du deuxième lot comprenaient quatre sections et cinq questions,
avec exercices corrigés et prérequis. Ce sont des leçons de synthèse, avec exemples
pédagogiques ; elles ne constituent pas une couverture exhaustive de tous les
objectifs d'examen du PPL(A). Les exemples de navigation et de performances ne
sont pas des données à utiliser pour préparer un vol réel.

## Statut éditorial et références

L’[audit d’adéquation aux examens du 18 septembre 2026](ppl-official-alignment-audit.md)
a conduit à un [lot de corrections et à un mode examen blanc](ppl-alignment-changes.md).
Les erreurs ciblées et le format des épreuves ont été repris ; la couverture
intégrale du programme et la validation pédagogique externe restent à établir.

Les nouveaux fichiers portent `content_status: draft`, un périmètre
« PPL(A), France / cadre européen » et leurs références. Ils nécessitent une
validation par un instructeur avant d'être présentés comme un cours approuvé.
Les trois leçons anciennes ont été reprises dans le lot d’alignement de septembre 2026.
Les cours, énoncés, choix et corrections sont disponibles en français et en anglais.
La traduction conserve le cadre France / européen et le statut éditorial initial.

Références de départ consultées le 15 septembre 2026 :

- [EASA, Aircrew et programme Part-FCL](https://www.easa.europa.eu/en/downloads/115485/en)
- [EASA, SERA](https://www.easa.europa.eu/en/document-library/easy-access-rules/easy-access-rules-standardised-european-rules-air)
- [EASA, Air Operations et Part-NCO](https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-air-operations?erules-id=ERULES-1963177438-13596)
- [Météo-France, Guide aviation](https://aviation.meteo.fr/documentation/guide_aviation.pdf)
- [FAA, Pilot’s Handbook of Aeronautical Knowledge](https://www.faa.gov/regulations_policies/handbooks_manuals/aviation/phak)
- [FAA, Airplane Flying Handbook](https://www.faa.gov/regulations_policies/handbooks_manuals/aviation/airplane_handbook)
- [SIA, publications aéronautiques françaises](https://www.sia.aviation-civile.gouv.fr/)

Les références FAA servent aux notions techniques générales ; les règles
américaines ne sont pas présentées comme des règles européennes. Les liens vers
les manuels techniques constituent aussi des pistes d'approfondissement.

## Sources des contenus et import local

La source éditable est `content/fr/ppl/**/*.json`, au format déjà utilisé par le
projet. Modifier ces fichiers puis régénérer le SQL :

```powershell
python scripts/seed/ppl_content.py
docker cp scripts/seed/ppl_content.sql aeropath-local-postgres-1:/tmp/ppl_content.sql
docker compose -p aeropath-local -f compose.local.yml exec -T postgres psql -U aeropath -d aeropath -v ON_ERROR_STOP=1 -f /tmp/ppl_content.sql
```

Le nouvel import est transactionnel et utilise `ON CONFLICT (id) DO UPDATE`.
Il conserve la méthode de calcul des identifiants existants. Il ne supprime
aucune question, leçon, donnée personnelle, progression ou configuration.
Les catégories sont normalisées selon les identifiants déjà utilisés par
l'interface, par exemple `airlaw`, `aircraft_general` et `meteorology`.
Retirer un fichier source ne supprime donc pas automatiquement sa ligne en base.

Ne pas substituer à cette commande l'ancien seed global : celui-ci contient
des suppressions de données et répond à un autre usage.

## Illustration et affichage

Les quiz affichent aussi des SVG : les 945 corrections reprennent le schéma
de leur leçon, après la réponse. Dix-huit énoncés de calcul ont un schéma d'exercice
distinct, avec les données et l'inconnue, sans résultat affiché. Ils portent sur
le temps de parcours, le carburant, le moment d'un bagage et la base des nuages.
L'énoncé textuel conserve toutes les données pour rester utilisable sans image.

Régénérer ces liens et les dix-huit SVG avec
`python scripts/illustrations/ppl_quiz_illustrations.py`, puis lancer
`python scripts/seed/ppl_content.py` et importer le SQL. Le rendu Markdown
sécurisé existant est utilisé dans les énoncés et les corrections ; aucune
modification du schéma de base de données n'est nécessaire.

Les 63 leçons contiennent chacune une illustration SVG locale, intégrée au début
du premier concept. Les fichiers sont dans
`web/frontend-react/public/illustrations/`. La leçon `PPL_AC_007` conserve
`commandes-de-vol.svg` ; les autres utilisent leur identifiant, par exemple
`ppl-met-007.svg`.

Les schémas sont définis dans `scripts/illustrations/ppl_diagrams.json` et générés
avec `python scripts/illustrations/ppl_illustrations.py`, avant de régénérer
l'import SQL. Le générateur vérifie que chaque leçon dispose d'un schéma et
remplace les références existantes sans multiplier les images. Toute nouvelle
leçon doit recevoir une définition adaptée à son sujet dans ce fichier.
Le rendu autorise les SVG du répertoire public `illustrations` et les liens HTTPS
des références, tout en continuant à échapper le HTML brut.

[Ouvrir Commandes de vol en local](http://localhost:3000/lessons/PPL/aircraft_general/40f1e143-0000-4000-8000-000040f1e143)

## Validation

Les quiz de leçon présentent les quinze questions de leur banque, dans un ordre aléatoire.
Les trois anciens chapitres conservent chacun leurs quinze questions ; les
53 autres leçons en comportent également quinze. Chaque question possède une correction
expliquée. L'import exige au moins quinze questions et refuse les énoncés dupliqués
dans une même leçon. Les nouvelles questions sont ajoutées en fin de liste pour
conserver les identifiants existants.

Le bouton « Refaire ce quiz » reprend les mêmes questions avec un score remis
à zéro. Le quiz permet de revenir à sa leçon. Les corrections affichent le texte
de la bonne réponse ; les envois multiples pendant une correction sont bloqués.
Les erreurs de chargement et de correction restent visibles et permettent de
réessayer. L'accès sans compte est conservé ; seul un compte connecté alimente
l'historique des réponses.

Tests du parcours : `node --test tests/quizStore.test.mjs tests/renderMarkdown.test.mjs`
depuis `web/frontend-react`, et `go test ./internal/learning ./internal/transport/http
./internal/transport/middleware` depuis la racine.

- Le générateur vérifie JSON, identifiants uniques, concepts et réponses attendues.
- Les tests du rendu Markdown couvrent images locales et refus d'URL exécutables.
- La compilation du frontend passe.
- Le lot quiz du 16 septembre 2026 a été importé : 56 quiz publics de cinq
  questions contrôlés, 310 questions vérifiées via 620 réponses invitées
  (une juste et une fausse par question). Les corrections correspondent aux
  sources JSON, les identifiants existants sont conservés et les empreintes
  des comptes, historiques et progressions sont inchangées après ces essais.
- Le premier import a été répété sans doublons : 50 leçons, 186 questions,
  2 comptes préexistants conservés et historique inchangé.
- Les endpoints publics de détail et de quiz du premier lot ont été vérifiés.
- Le deuxième lot a été importé le 16 septembre 2026. Les empreintes des
  comptes, historiques, progressions et données de gamification sont identiques
  avant et après l'import. Les calculs d'exemple et les prérequis ont été vérifiés.

La progression des lots et les prochaines priorités sont consignées dans
[le suivi du programme](ppl-course-roadmap.md).

Cette validation technique ne vaut pas validation pédagogique ou réglementaire.

## Anglais

Le menu de langue sélectionne les titres, le contenu, les choix et les corrections.
Les 81 illustrations anglaises sont dans `web/frontend-react/public/illustrations/en`.
Les positions A–D restent identiques entre langues ; `answer_key` reste canonique.
Une traduction manquante revient au français. Le changement de langue ne remet
pas le quiz à zéro.

Appliquer `scripts/migrations/006_question_options_en.sql` avant l'import SQL.
Pour régénérer les SVG anglais :
`python scripts/illustrations/ppl_illustrations_en.py`.
Le dictionnaire est `scripts/illustrations/ppl_svg_en.json`.
Régénérer ensuite l'import avec `python scripts/seed/ppl_content.py`.

Les tests `tests/learningLocale.test.mjs` couvrent la conservation des positions,
le repli et la présence des traductions et SVG pour les 63 leçons et 945 questions.

Validation du 17 septembre 2026 : import bilingue effectué, 56 quiz et 310
questions contrôlés via API, 620 réponses invitées justes/fausses validées.
Identifiants, comptes et progressions inchangés. Les 66 SVG anglais ont été
rendus et contrôlés sans chevauchement de texte détecté. Compilation et tests
frontend/Go réussis.

## Langue et unités indépendantes

La barre de navigation contient deux sélecteurs distincts, également accessibles
via le menu sur petit écran. Les préférences sont conservées localement dans
`aeropath_lang` et `aeropath_units`, y compris sans compte.

Profils : métrique, américain (US gal) et britannique (imp gal). Les profils
US et britannique ajoutent les équivalences arrondies de kg, kg·m, L, L/h, m
et cm à côté des valeurs originales. Les expressions de calcul, les positions
des réponses et les unités aéronautiques usuelles restent explicites. Les
équivalences ne changent pas la réglementation du cours.

Les variantes SVG comprennent un tableau d’équivalences sous le schéma.
Régénération : `node scripts/illustrations/ppl_unit_variants.mjs`, après toute
modification des schémas ou du convertisseur. Les schémas FR/EN disposent de
324 variantes (81 × 2 langues × 2 profils supplémentaires).

Tests : `node --test tests/measurementSystem.test.mjs` depuis le frontend.
La traduction espagnole est en cours : le choix de l’espagnol utilise le repli
français quand le contenu traduit manque. Ce sélecteur ne constitue pas une
validation de complétude de chaque langue proposée.

## Extension à quinze questions par leçon — 17 septembre 2026

Les 56 banques contiennent chacune 15 questions, soit 840 au total. Les trois
anciennes banques de réglementation conservent leurs 15 questions. Dans les
53 autres leçons, dix affirmations vrai/faux bilingues complètent les cinq QCM
existants, avec une justification issue du cours et une correction illustrée.
Huit nouveaux schémas de calcul portent à 18 les SVG spécifiques aux énoncés.
Les identifiants des questions précédentes sont conservés ; les quiz de leçon,
les sessions aléatoires et les sessions par thème proposent 15 questions.

Après modification des SVG, régénérer dans cet ordre : illustrations de quiz,
illustrations anglaises, variantes d’unités, puis SQL de contenu.

Validation locale : 56 quiz de 15 questions, 840 questions et 1 680 réponses
invitées vérifiés via API. Comptes, historiques et progressions inchangés.
Les 148 SVG FR/EN et 296 variantes d’unités passent le contrôle de débordement.
Tests Go et frontend et compilation de production réussis.

## Lot d’alignement — 22 septembre 2026

Cinq leçons et 75 QCM ajoutés ; trois leçons réglementaires corrigées.
[Modifications, vérifications et limites](ppl-alignment-changes.md).
[Ouvrir les neuf épreuves d’entraînement](http://localhost:3000/ppl-exam).


## Lot du 22 septembre 2026 — priorités et espaces

Deux leçons FR/EN supplémentaires : `PPL_AIRLAW_012` (prévention des collisions,
priorités, dépassement, finale et barres d’arrêt) et `PPL_AIRLAW_013` (classes,
services ATS, clairances, RMZ et TMZ). Chacune contient quatre sections,
15 QCM originaux à quatre choix, des explications bilingues et un SVG de synthèse
repris dans les corrections. Les SVG existent aussi avec les profils US et britannique.

Les références précises SERA.3201, SERA.3210, SERA.6001 et SERA.6005 sont
attachées aux concepts et questions ; consultation EASA du 22 septembre 2026.
La banque atteint 63 leçons et 945 questions, dont 415 QCM utilisables en examen
blanc et 530 vrai/faux de révision. Le choix de langue ou d’unités ne change pas
le cadre réglementaire européen. Les signaux, l’ensemble des procédures locales
et les autres lacunes de l’audit restent à compléter ; validation instructeur à obtenir.
