"""Bilingual teaching content, revised after individual pedagogical review."""
from shared.authoring import save


def build():
    data = {'metadata': {'module_id': 'ATPL_031_001',
                  'program': 'ATPL',
                  'subject': 'mass_and_balance',
                  'chapter_title': 'Masses opérationnelles et évolution du centrage',
                  'chapter_title_en': 'Operating masses and CG evolution',
                  'reference_checked_on': '2026-09-29',
                  'learning_objective': 'Établir un bilan sur données nouvelles, identifier la contrainte '
                                        'limitante, lire une enveloppe interpolée et détecter un dépassement '
                                        'intermédiaire de centrage.',
                  'learning_objective_en': 'Build a balance from new inputs, identify the limiting constraint, '
                                           'read an interpolated envelope and detect intermediate CG exceedance.',
                  'estimated_duration_minutes': 55,
                  'difficulty': 'advanced',
                  'prerequisites': ['Masse, moments et centrage',
                                    'DOW, charge marchande et carburant utilisable',
                                    'Proportions et pourcentage MAC'],
                  'exam_subject': '031',
                  'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 031',
                  'official_subject_format': {'questions': 25,
                                              'minutes': 75,
                                              'pass_percent': 75,
                                              'source': 'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67'},
                  'sources': [{'title': 'EASA — programme Part-FCL',
                               'url': 'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8'},
                              {'title': 'EASA — format de la matière, AMC1 ARA.FCL.300(b)',
                               'url': 'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67'},
                              {'title': 'FAA Weight and Balance Handbook — principes techniques',
                               'url': 'https://www.faa.gov/sites/faa.gov/files/2023-09/Weight_Balance_Handbook.pdf'}],
                  'content_status': 'draft',
                  'coverage_status': 'partial',
                  'scope': 'ATPL(A), France / EASA',
                  'quiz_format': 'chapter_practice',
                  'tags': ['ATPL', 'mcq-four-options'],
                  'illustration': {'src': '/illustrations/atpl-031-001.svg',
                                   'alt': 'Masses opérationnelles et évolution du centrage'}},
     'concepts': [{'concept_id': 'ATPL_031_001_C01',
                   'title': 'Distinguer les étapes du bilan',
                   'title_en': 'Distinguish mass stages',
                   'content': 'Exercice fictif : masse opérationnelle sèche DOW = 30 000 kg, charge marchande = '
                              '8 000 kg. La masse sans carburant utilisable ZFM vaut 38 000 kg. Ajouter 10 000 '
                              'kg de carburant bloc donne une masse au parking de 48 000 kg. Après 200 kg de '
                              'roulage, la masse au décollage vaut 47 800 kg. Avec 6 000 kg de carburant trajet '
                              'consommés après décollage, la masse prévue à l’atterrissage vaut 41 800 kg. Ne '
                              'pas soustraire le roulage deux fois.\n'
                              '\n'
                              '### Identifier le contenu de chaque masse\n'
                              '\n'
                              'La DOW représente ici l’avion prêt pour le type d’exploitation considéré, hors '
                              'carburant utilisable et charge marchande. Ce n’est pas simplement la masse de la '
                              'structure vide. Vérifier ce que le document d’exploitation inclut déjà, pour ne '
                              'pas ajouter deux fois équipage ou équipements.\n'
                              '\n'
                              'Suivre une chronologie écrite : DOW + charge marchande = ZFM ; ZFM + carburant '
                              'bloc = masse au parking ; parking − roulage = décollage ; décollage − trajet = '
                              'atterrissage prévu. Dans cet exercice, aucune autre masse n’est ajoutée ou '
                              'retirée entre ces étapes, et les consommations données sont des masses de '
                              'carburant.\n'
                              '\n'
                              '**Contrôle par le carburant restant :** 10 000 − 200 − 6 000 = 3 800 kg. ZFM 38 '
                              '000 + restant 3 800 = **41 800 kg**, identique au calcul par soustractions '
                              'successives. Déduire encore 200 kg au stade atterrissage compterait le roulage '
                              'deux fois.\n'
                              '\n'
                              'Les 3 800 kg restants ne démontrent pas à eux seuls le respect des réserves '
                              'réglementaires : ce chapitre vérifie un bilan de masses. Les allocations '
                              'carburant et les hypothèses opérationnelles demandent leur propre contrôle.',
                   'content_en': 'Fictional exercise: dry operating mass DOW = 30,000 kg, traffic load = 8,000 '
                                 'kg. Zero-fuel mass ZFM is 38,000 kg. Adding 10,000 kg block fuel gives 48,000 '
                                 'kg ramp mass. After 200 kg taxi fuel, take-off mass is 47,800 kg. With 6,000 '
                                 'kg trip fuel consumed after take-off, planned landing mass is 41,800 kg. Do '
                                 'not subtract taxi fuel twice.\n'
                                 '\n'
                                 '### Identify what each mass includes\n'
                                 '\n'
                                 'DOW here represents the aircraft ready for the specified operation, excluding '
                                 'usable fuel and traffic load. It is not simply the mass of an empty structure. '
                                 'Check what the operating document already includes to avoid adding crew or '
                                 'equipment twice.\n'
                                 '\n'
                                 'Write the chronology: DOW + traffic load = ZFM; ZFM + block fuel = ramp mass; '
                                 'ramp minus taxi fuel = take-off mass; take-off minus trip fuel = planned '
                                 'landing mass. In this exercise, no other mass is added or removed between '
                                 'stages, and supplied consumptions are fuel masses.\n'
                                 '\n'
                                 '**Cross-check using remaining fuel:** 10,000 − 200 − 6,000 = 3,800 kg. ZFM '
                                 '38,000 + remaining 3,800 = **41,800 kg**, matching sequential subtraction. '
                                 'Subtracting another 200 kg at landing would count taxi fuel twice.\n'
                                 '\n'
                                 'The remaining 3,800 kg alone does not establish regulatory reserve compliance: '
                                 'this chapter checks a mass balance. Fuel allocations and operational '
                                 'assumptions require their own checks.',
                   'difficulty': 'advanced',
                   'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — 031',
                   'tags': ['ATPL']},
                  {'concept_id': 'ATPL_031_001_C02',
                   'title': 'Vérifier chaque limitation',
                   'title_en': 'Check every limit',
                   'content': 'Limites fictives : maximum ZFM 39 000 kg, décollage 49 000 kg, atterrissage 42 '
                              '000 kg. Le bilan initial respecte ces trois plafonds de masse, sans prouver la '
                              'conformité des performances ou du centrage. À carburant inchangé, ajouter une '
                              'charge ne peut dépasser la plus petite marge restante : 1 000, 1 200 et 200 kg '
                              'respectivement. La limite d’atterrissage restreint donc l’ajout à 200 kg dans cet '
                              'exercice. Des limites opérationnelles peuvent être plus restrictives que les '
                              'maxima structuraux.\n'
                              '\n'
                              '### Trouver la contrainte qui commande la décision\n'
                              '\n'
                              'Écrire les trois inégalités plutôt que d’additionner les marges. Une charge '
                              'supplémentaire p, à carburant bloc, roulage et trajet inchangés, impose p ≤ 1 000 '
                              'kg, p ≤ 1 200 kg et p ≤ 200 kg. Il faut satisfaire les trois simultanément, donc '
                              '**p ≤ 200 kg**. Une moyenne de ces marges n’a pas de sens opérationnel.\n'
                              '\n'
                              '**Sensibilité au carburant :** si l’on ajoute séparément 400 kg de carburant '
                              'bloc, avec mêmes consommations, la ZFM reste 38 000 kg, le décollage passe à 48 '
                              '200 kg et l’atterrissage à 42 200 kg. Le décollage respecte 49 000 kg, mais '
                              'l’arrivée dépasse 42 000 kg de 200 kg. Ces seules données ne permettent pas de '
                              'choisir quelle allocation carburant réduire : il faut conserver les exigences de '
                              'carburant et recalculer le scénario.\n'
                              '\n'
                              '**Cas autonome :** DOW 30 000 kg, charge 7 500 kg, bloc 10 000 kg, roulage 200 '
                              'kg, trajet 6 000 kg, mêmes plafonds. ZFM 37 500, décollage 47 300, arrivée 41 300 '
                              'kg ; marges 1 500, 1 700 et 700 kg. Le plafond limitant pour un ajout de charge '
                              'est donc **700 kg**, sous les hypothèses fixées et les autres contraintes.\n'
                              '\n'
                              '### Ajouter une limite de performance au bilan\n'
                              '\n'
                              'Reprendre le cas autonome à 7 500 kg de charge, dont la masse au décollage est 47 '
                              '300 kg. Si une limite de performance fictive impose 47 600 kg, la marge '
                              'supplémentaire vaut **300 kg**. Elle est plus faible que les 700 kg permis par '
                              'les trois seuls plafonds structuraux. Ajouter la nouvelle contrainte à la liste '
                              'et reprendre le minimum ; ne pas remplacer arbitrairement tous les autres '
                              'contrôles par celui du décollage.\n'
                              '\n'
                              'Si l’exercice impose des consommations inchangées, les utiliser comme telles. Une '
                              'exploitation réelle doit réévaluer les consommations et performances lorsque le '
                              'chargement change : la constance du carburant trajet est ici une hypothèse de '
                              'calcul, pas une règle générale.',
                   'content_en': 'Fictional limits: maximum ZFM 39,000 kg, take-off 49,000 kg, landing 42,000 '
                                 'kg. The initial balance satisfies these three mass ceilings without '
                                 'establishing performance or CG compliance. With unchanged fuel quantities, '
                                 'added payload cannot exceed the smallest remaining margin: 1,000, 1,200 and '
                                 '200 kg respectively. Landing mass therefore limits the addition to 200 kg in '
                                 'this exercise. Operational limits may be more restrictive than structural '
                                 'maxima.\n'
                                 '\n'
                                 '### Find the constraint that controls the decision\n'
                                 '\n'
                                 'Write the three inequalities instead of adding margins. Additional traffic '
                                 'load p, with unchanged block, taxi and trip fuel, requires p ≤ 1,000 kg, p ≤ '
                                 '1,200 kg and p ≤ 200 kg. All must hold together, so **p ≤ 200 kg**. Averaging '
                                 'the margins has no operational meaning.\n'
                                 '\n'
                                 '**Fuel sensitivity:** separately adding 400 kg of block fuel with unchanged '
                                 'consumption keeps ZFM at 38,000 kg, raises take-off to 48,200 kg and landing '
                                 'to 42,200 kg. Take-off meets 49,000 kg, but arrival exceeds 42,000 kg by 200 '
                                 'kg. These inputs alone cannot determine which fuel allocation to reduce: fuel '
                                 'requirements must remain satisfied and the scenario must be recalculated.\n'
                                 '\n'
                                 '**Independent case:** DOW 30,000 kg, traffic load 7,500 kg, block 10,000 kg, '
                                 'taxi 200 kg, trip 6,000 kg, same ceilings. ZFM 37,500, take-off 47,300, '
                                 'arrival 41,300 kg; margins 1,500, 1,700 and 700 kg. The limiting ceiling for '
                                 'added traffic load is therefore **700 kg**, subject to fixed assumptions and '
                                 'other constraints.\n'
                                 '\n'
                                 '### Add a performance limit to the balance\n'
                                 '\n'
                                 'Return to the independent case with 7,500 kg traffic load and take-off mass '
                                 '47,300 kg. If a fictional performance limit imposes 47,600 kg, the additional '
                                 'margin is **300 kg**. It is smaller than the 700 kg allowed by the three '
                                 'structural ceilings alone. Add the new constraint and take the minimum again; '
                                 'do not arbitrarily replace all other checks with the take-off check.\n'
                                 '\n'
                                 'If the exercise specifies unchanged consumption, use that assumption. Real '
                                 'operations must reassess consumption and performance after loading changes: '
                                 'constant trip fuel is a calculation assumption here, not a general rule.',
                   'difficulty': 'advanced',
                   'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — 031',
                   'tags': ['ATPL']},
                  {'concept_id': 'ATPL_031_001_C03',
                   'title': 'Recalculer le centrage après consommation',
                   'title_en': 'Recalculate CG after fuel burn',
                   'content': 'Exercice de centrage indépendant : masse 10 000 kg, centrage à 5,00 m du datum ; '
                              'moment total 50 000 kg·m. Consommer 1 000 kg de carburant au bras constant 6,00 m '
                              'retire 6 000 kg·m. Le nouveau centrage est 44 000 / 9 000 = 4,889 m : il avance, '
                              'car la masse retirée était derrière le centrage initial. Ne pas déduire une '
                              'direction générale du seul mot « consommation » : elle dépend du bras et de la '
                              'séquence des réservoirs.\n'
                              '\n'
                              '### Prévoir le sens avant de calculer\n'
                              '\n'
                              'Dans cet exercice indépendant, les bras augmentent vers l’arrière. Au départ, 10 '
                              '000 × 5 = 50 000 kg·m. Retirer 1 000 kg à 6 m laisse 44 000 kg·m et 9 000 kg. Le '
                              'nouveau CG 4,8889… m est inférieur à 5 m, ce qui confirme le déplacement vers '
                              'l’avant attendu.\n'
                              '\n'
                              '**Contre-exemple :** repartir du même état initial mais retirer les 1 000 kg au '
                              'bras 4 m. Le nouveau CG est (50 000 − 4 000)/9 000 = **5,1111 m** : il recule. La '
                              'consommation ne fait donc pas toujours avancer le CG. Il faut connaître le '
                              'réservoir, son bras et la séquence de consommation.\n'
                              '\n'
                              'Ne pas réutiliser le bilan de 47 800 kg pour ce cas de 10 000 kg : les deux '
                              'exercices sont indépendants. Dans un calcul réel, suivre la méthode de réservoir '
                              'applicable, notamment si le bras varie avec la quantité. Vérifier les états '
                              'intermédiaires pertinents ; des extrémités conformes ne prouvent pas '
                              'automatiquement que tout le chemin de consommation l’est.\n'
                              '\n'
                              '### Combiner chargement et consommation\n'
                              '\n'
                              '**Exemple distinct :** état initial 12 000 kg au CG 5 m, soit 60 000 kg·m. '
                              'Déplacer 300 kg du bras 7 m au bras 4 m enlève 900 kg·m sans changer la masse. '
                              'Consommer ensuite 1 000 kg au bras constant 6 m enlève 1 000 kg et 6 000 kg·m. '
                              'Totaux finaux : 11 000 kg et 53 100 kg·m ; CG = **4,827 m** arrondi.\n'
                              '\n'
                              'Une ligne par modification évite de soustraire la charge déplacée de la masse '
                              'totale ou d’oublier son ancien moment. Pour plusieurs réservoirs, sommer '
                              'séparément les masses et les moments consommés ; le bras moyen du carburant '
                              'retiré peut être en avant, en arrière ou exactement au CG initial. Le résultat '
                              'final n’exonère pas du contrôle des états intermédiaires pertinents.',
                   'content_en': 'Independent CG exercise: mass 10,000 kg, CG 5.00 m from datum; total moment '
                                 '50,000 kg·m. Burning 1,000 kg fuel at constant arm 6.00 m removes 6,000 kg·m. '
                                 'New CG is 44,000 / 9,000 = 4.889 m: it moves forward because the removed mass '
                                 'was aft of the initial CG. Fuel burn alone does not define a universal '
                                 'direction: it depends on arm and tank sequence.\n'
                                 '\n'
                                 '### Predict direction before calculating\n'
                                 '\n'
                                 'In this independent exercise, arms increase aft. Initially, 10,000 × 5 = '
                                 '50,000 kg·m. Removing 1,000 kg at 6 m leaves 44,000 kg·m and 9,000 kg. New CG '
                                 '4.8889… m is below 5 m, confirming the expected forward movement.\n'
                                 '\n'
                                 '**Counterexample:** start from the same initial condition but remove 1,000 kg '
                                 'at arm 4 m. New CG is (50,000 − 4,000)/9,000 = **5.1111 m**: it moves aft. '
                                 'Fuel burn therefore does not always move CG forward. Tank, arm and consumption '
                                 'sequence must be known.\n'
                                 '\n'
                                 'Do not reuse the 47,800 kg balance for this 10,000 kg case: the exercises are '
                                 'independent. For a real calculation, use the applicable tank method, '
                                 'especially if arm varies with quantity. Check relevant intermediate '
                                 'conditions; compliant endpoints do not automatically prove the entire '
                                 'consumption path is compliant.\n'
                                 '\n'
                                 '### Combine loading and consumption\n'
                                 '\n'
                                 '**Separate example:** initial state 12,000 kg at CG 5 m, giving 60,000 kg·m. '
                                 'Moving 300 kg from arm 7 m to 4 m removes 900 kg·m without changing mass. Then '
                                 'burning 1,000 kg at constant arm 6 m removes 1,000 kg and 6,000 kg·m. Final '
                                 'totals: 11,000 kg and 53,100 kg·m; CG = **4.827 m** rounded.\n'
                                 '\n'
                                 'One entry per change prevents subtracting moved load from total mass or '
                                 'forgetting its old moment. For multiple tanks, add consumed masses and moments '
                                 'separately; average arm of removed fuel may be forward of, aft of or exactly '
                                 'at initial CG. The final result does not remove the need to check relevant '
                                 'intermediate conditions.',
                   'difficulty': 'advanced',
                   'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — 031',
                   'tags': ['ATPL']},
                  {'concept_id': 'ATPL_031_001_C04',
                   'title': 'Exprimer le centrage en pourcentage de corde',
                   'title_en': 'Express CG as percentage MAC',
                   'content': 'La position en pourcentage de corde aérodynamique moyenne vaut 100 × (bras CG − '
                              'bras du bord d’attaque de la MAC) / longueur MAC. Exemple : CG 5,00 m, bord '
                              'd’attaque 4,00 m, MAC 4,00 m : 25 %. Un centrage à 4,80 m donne 20 %. Comparer '
                              'masse et centrage à l’enveloppe applicable à chaque phase, pas uniquement au '
                              'départ. La concordance d’un pourcentage ne dispense pas de vérifier la référence, '
                              'la configuration et les limites propres à l’avion.\n'
                              '\n'
                              '### Vérifier une conversion dans les deux sens\n'
                              '\n'
                              'À 25 %MAC, retrouver le bras : 4,00 + 0,25 × 4,00 = **5,00 m**. Le pourcentage '
                              'décrit une position à l’intérieur d’une corde de référence ; il ne représente pas '
                              'une fraction de la longueur totale de l’avion. Une valeur de 20 signifie 20 %, '
                              'soit 0,20 dans la formule inverse.\n'
                              '\n'
                              '**Application nouvelle :** bord d’attaque à 3,20 m, MAC de 2 m, CG à 3,70 m. Le '
                              'résultat vaut 100 × 0,50/2 = **25 %**, mais le bras n’est pas 5 m. Un même '
                              'pourcentage sur deux avions ou deux références ne prouve donc pas une même '
                              'position en mètres.\n'
                              '\n'
                              'Avant de conclure, contrôler les unités et le datum, les définitions de masse, la '
                              'masse correspondant à l’enveloppe et les limitations du cas. Séparer '
                              'explicitement le résultat arithmétique, le respect des plafonds présentés et les '
                              'vérifications opérationnelles qui restent nécessaires.\n'
                              '\n'
                              '### Lire une enveloppe et vérifier la séquence\n'
                              'Enveloppe fictive de cet exercice : masse autorisée de 8 000 à 12 000 kg inclus ; '
                              'limite arrière 35 %MAC ; limite avant linéaire de 15 % à 8 000 kg à 25 % à 12 000 '
                              'kg. Les limites sont incluses. Cette enveloppe sert uniquement à apprendre la '
                              "lecture ; elle n'appartient à aucun avion réel.\n"
                              '\n'
                              'Pour une masse intermédiaire M, limite avant = 15 + (M − 8 000)/4 000 × 10, '
                              'exprimée en points de pourcentage. À 9 500 kg, elle vaut 18,75 %MAC. Un CG à 22 % '
                              "est à l'intérieur horizontalement et la masse est admise ; un CG à 18 % est trop "
                              'avant. Lire la limite à la masse du point, pas à la masse initiale ni à la seule '
                              'valeur minimale 15 %.\n'
                              '\n'
                              "Tracer les états successifs dans le même repère. Si une séquence consomme d'abord "
                              'du carburant situé en avant du CG puis du carburant arrière, un déplacement '
                              "temporaire vers l'arrière peut être suivi d'un retour. Deux extrémités à "
                              "l'intérieur ne prouvent donc pas l'acceptabilité de l'état intermédiaire. Pour "
                              'chaque ligne, recalculer masse, moment, bras, %MAC puis limites. Les bras '
                              'constants des exercices sont une hypothèse ; utiliser les données réelles du '
                              'réservoir si son bras varie.\n'
                              '\n'
                              '**Auto-évaluation :** à 11 000 kg, quelle est la limite avant ? 22,5 %MAC. Un '
                              'point à 21 % échoue malgré une masse inférieure au maximum. Dans les questions, '
                              'la masse, le centrage et les séquences changent ; tracer avant de conclure et ne '
                              'pas arrondir assez tôt pour faire entrer un point hors limite.\n'
                              '\n'
                              '**Format de travail :** 15 QCM de chapitre, quatre choix et une réponse correcte. '
                              'La matière 031 comporte 25 questions en 75 minutes dans le tableau EASA cité. Ce '
                              'chapitre ne couvre qu’une partie de la matière ; il ne constitue pas une épreuve '
                              'complète. Cadre EASA, même en anglais ou avec les unités US/britanniques.',
                   'content_en': 'CG as a percentage of mean aerodynamic chord is 100 × (CG arm − MAC '
                                 'leading-edge arm) / MAC length. Example: CG 5.00 m, leading edge 4.00 m, MAC '
                                 '4.00 m: 25%. A CG at 4.80 m gives 20%. Check mass and CG against the envelope '
                                 'applicable to every phase, not just departure. A matching percentage does not '
                                 'remove the need to verify datum, configuration and aircraft-specific limits.\n'
                                 '\n'
                                 '### Check a conversion in both directions\n'
                                 '\n'
                                 'At 25% MAC, recover the arm: 4.00 + 0.25 × 4.00 = **5.00 m**. The percentage '
                                 'describes position along a reference chord; it is not a fraction of total '
                                 'aircraft length. A value of 20 means 20%, or 0.20 in the inverse formula.\n'
                                 '\n'
                                 '**New application:** leading edge 3.20 m, MAC length 2 m, CG 3.70 m. The '
                                 'result is 100 × 0.50/2 = **25%**, but the arm is not 5 m. The same percentage '
                                 'on different aircraft or references therefore does not establish the same '
                                 'position in metres.\n'
                                 '\n'
                                 'Before concluding, check units and datum, mass definitions, the mass relevant '
                                 'to the envelope and case limitations. Explicitly separate the arithmetic '
                                 'result, compliance with the presented ceilings and remaining operational '
                                 'checks.\n'
                                 '\n'
                                 '### Read an envelope and check the sequence\n'
                                 'Fictional exercise envelope: permitted mass from 8,000 to 12,000 kg inclusive; '
                                 'aft limit 35% MAC; forward limit linear from 15% at 8,000 kg to 25% at 12,000 '
                                 'kg. Boundaries are included. This teaching envelope belongs to no real '
                                 'aircraft.\n'
                                 '\n'
                                 'At intermediate mass M, forward limit = 15 + (M − 8,000)/4,000 × 10 in '
                                 'percentage points. At 9,500 kg it is 18.75% MAC. A CG of 22% is horizontally '
                                 'inside and the mass is permitted; 18% is too far forward. Read the limit at '
                                 "the point's actual mass, not the initial mass or only the minimum value of "
                                 '15%.\n'
                                 '\n'
                                 'Plot successive states on the same axes. Burning fuel forward of CG and then '
                                 'aft fuel can produce a temporary aft movement followed by return. Two '
                                 'endpoints inside therefore do not prove an acceptable intermediate state. For '
                                 'every row, recalculate mass, moment, arm, %MAC and limits. Constant tank arms '
                                 'are an exercise assumption; use actual tank data when the arm varies.\n'
                                 '\n'
                                 '**Self-check:** at 11,000 kg, what is the forward limit? 22.5% MAC. A point at '
                                 '21% fails despite mass below the maximum. Questions change mass, CG and '
                                 'sequences; plot before concluding and do not round early enough to move an '
                                 'outside point inside.\n'
                                 '\n'
                                 '**Practice format:** 15 chapter MCQs, four choices and one correct answer. '
                                 'Subject 031 has 25 questions in 75 minutes in the cited EASA table. This '
                                 'chapter covers only part of the subject and is not a full paper. EASA '
                                 'framework, including with English or US/British units.',
                   'difficulty': 'advanced',
                   'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — 031',
                   'tags': ['ATPL']}],
     'questions': [{'concept_id': 'ATPL_031_001_C01',
                    'concept_title': 'Distinguer les étapes du bilan',
                    'level': 3,
                    'question': 'DOW 28 000 kg, charge marchande 6 500 kg, carburant utilisable bloc 9 000 kg. '
                                'ZFM ?',
                    'question_en': 'DOW 28,000 kg, traffic load 6,500 kg, usable block fuel 9,000 kg. ZFM?',
                    'options': ['34 500 kg', '43 500 kg', '28 000 kg', '37 000 kg'],
                    'options_en': ['34,500 kg', '43,500 kg', '28,000 kg', '37,000 kg'],
                    'correct_answer': '34 500 kg',
                    'explanation': 'ZFM = 28 000 + 6 500 = 34 500 kg. 43 500 ajoute le bloc ; 28 000 oublie la '
                                   'charge ; 37 000 ajoute le carburant à la place de la charge. Identifier le '
                                   "contenu avant l'addition.",
                    'explanation_en': 'ZFM = 28,000 + 6,500 = 34,500 kg. 43,500 adds block fuel; 28,000 omits '
                                      'load; 37,000 adds fuel instead of load. Identify components before '
                                      'adding.',
                    'use_exercise_svg': False,
                    'tags': ['ATPL', 'exam-mcq', 'exam-subject:031'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/sites/faa.gov/files/2023-09/Weight_Balance_Handbook.pdf'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 031'},
                   {'concept_id': 'ATPL_031_001_C01',
                    'concept_title': 'Distinguer les étapes du bilan',
                    'level': 3,
                    'question': 'DOW 28 000 kg, charge 6 500 kg, bloc 9 000 kg, roulage 150 kg. Masse au parking '
                                'avant roulage ?',
                    'question_en': 'DOW 28,000 kg, load 6,500 kg, block fuel 9,000 kg, taxi 150 kg. Ramp mass '
                                   'before taxi?',
                    'options': ['43 350 kg', '43 650 kg', '34 500 kg', '43 500 kg'],
                    'options_en': ['43,350 kg', '43,650 kg', '34,500 kg', '43,500 kg'],
                    'correct_answer': '43 500 kg',
                    'explanation': '28 000 + 6 500 + 9 000 = 43 500. Le bloc inclut le carburant de roulage : ni '
                                   "le soustraire avant cette phase, ni l'ajouter une seconde fois. 34 500 est "
                                   'la ZFM.',
                    'explanation_en': '28,000 + 6,500 + 9,000 = 43,500. Block fuel includes taxi fuel: neither '
                                      'subtract it before taxi nor add it twice. 34,500 is ZFM.',
                    'use_exercise_svg': False,
                    'tags': ['ATPL', 'exam-mcq', 'exam-subject:031'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/sites/faa.gov/files/2023-09/Weight_Balance_Handbook.pdf'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 031'},
                   {'concept_id': 'ATPL_031_001_C01',
                    'concept_title': 'Distinguer les étapes du bilan',
                    'level': 3,
                    'question': 'DOW 28 000 kg, charge 6 500 kg, bloc 9 000 kg dont 150 consommés au roulage. '
                                'Masse au décollage ?',
                    'question_en': 'DOW 28,000 kg, load 6,500 kg, block fuel 9,000 kg including 150 consumed in '
                                   'taxi. Take-off mass?',
                    'options': ['43 200 kg', '43 650 kg', '43 350 kg', '43 500 kg'],
                    'options_en': ['43,200 kg', '43,650 kg', '43,350 kg', '43,500 kg'],
                    'correct_answer': '43 350 kg',
                    'explanation': 'Parking 43 500 moins 150 = 43 350 kg. 43 200 retire le roulage deux fois ; '
                                   "43 650 l'ajoute au lieu de le consommer. Conserver la chronologie des états.",
                    'explanation_en': 'Ramp 43,500 minus 150 = 43,350 kg. 43,200 subtracts taxi twice; 43,650 '
                                      'adds rather than burns it. Preserve state chronology.',
                    'use_exercise_svg': False,
                    'tags': ['ATPL', 'exam-mcq', 'exam-subject:031'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/sites/faa.gov/files/2023-09/Weight_Balance_Handbook.pdf'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 031'},
                   {'concept_id': 'ATPL_031_001_C01',
                    'concept_title': 'Distinguer les étapes du bilan',
                    'level': 3,
                    'question': 'Décollage 43 350 kg ; carburant trajet après décollage 5 500 kg ; roulage '
                                "antérieur 150 kg. Masse prévue à l'atterrissage ?",
                    'question_en': 'Take-off mass 43,350 kg; trip fuel after take-off 5,500 kg; previous taxi '
                                   'fuel 150 kg. Planned landing mass?',
                    'options': ['43 200 kg', '37 850 kg', '37 700 kg', '38 000 kg'],
                    'options_en': ['43,200 kg', '37,850 kg', '37,700 kg', '38,000 kg'],
                    'correct_answer': '37 850 kg',
                    'explanation': "43 350 − 5 500 = 37 850. Le roulage est déjà compris dans l'état de départ "
                                   'au décollage ; le retirer ou le rajouter change à tort la référence. 43 200 '
                                   'ne retire que le roulage, pas le trajet.',
                    'explanation_en': '43,350 − 5,500 = 37,850. Taxi is already accounted for in take-off mass; '
                                      'subtracting or adding it changes the reference incorrectly. 43,200 '
                                      'removes only taxi, not trip fuel.',
                    'use_exercise_svg': False,
                    'tags': ['ATPL', 'exam-mcq', 'exam-subject:031'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/sites/faa.gov/files/2023-09/Weight_Balance_Handbook.pdf'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 031'},
                   {'concept_id': 'ATPL_031_001_C02',
                    'concept_title': 'Vérifier chaque limitation',
                    'level': 3,
                    'question': 'États initiaux ZFM/décollage/arrivée : 34 500/43 350/37 850 kg. Plafonds : 36 '
                                '000/45 000/39 000 kg. Ajout de charge 1 400 kg, carburants inchangés. '
                                'Dépassement ?',
                    'question_en': 'Initial ZFM/take-off/landing: 34,500/43,350/37,850 kg. Ceilings: '
                                   '36,000/45,000/39,000 kg. Add 1,400 kg load, fuel unchanged. Exceedance?',
                    'options': ['Atterrissage de 250 kg',
                                'ZFM de 100 kg',
                                'Décollage de 250 kg',
                                'Aucun dépassement'],
                    'options_en': ['Landing by 250 kg', 'ZFM by 100 kg', 'Take-off by 250 kg', 'No exceedance'],
                    'correct_answer': 'Atterrissage de 250 kg',
                    'explanation': "Nouveaux états 35 900/44 750/39 250 : seule l'arrivée dépasse, de 250. Une "
                                   "marge restante au décollage ne compense pas l'arrivée ; chaque plafond "
                                   "s'applique à son état.",
                    'explanation_en': 'New states 35,900/44,750/39,250: only landing exceeds, by 250. Remaining '
                                      'take-off margin cannot compensate for landing exceedance; each ceiling '
                                      'applies to its state.',
                    'use_exercise_svg': False,
                    'tags': ['ATPL', 'exam-mcq', 'exam-subject:031'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/sites/faa.gov/files/2023-09/Weight_Balance_Handbook.pdf'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 031'},
                   {'concept_id': 'ATPL_031_001_C02',
                    'concept_title': 'Vérifier chaque limitation',
                    'level': 3,
                    'question': 'DOW 28 000, charge 6 500, bloc 9 000, roulage 150, trajet 5 500 kg. Maxima '
                                'ZFM/décollage/arrivée 36 000/45 000/39 000 ; limite de performance au décollage '
                                '43 800 kg. Ajout maximal de charge, carburants inchangés ?',
                    'question_en': 'DOW 28,000, load 6,500, block 9,000, taxi 150, trip 5,500 kg. '
                                   'ZFM/take-off/landing maxima 36,000/45,000/39,000; take-off performance limit '
                                   '43,800 kg. Maximum added load with fuel unchanged?',
                    'options': ['1 150 kg', '1 500 kg', '1 650 kg', '450 kg'],
                    'options_en': ['1,150 kg', '1,500 kg', '1,650 kg', '450 kg'],
                    'correct_answer': '450 kg',
                    'explanation': 'États 34 500/43 350/37 850. Marges structurales 1 500/1 650/1 150 ; marge de '
                                   'performance 450. Le minimum des quatre vaut 450 ; prendre 1 150 ignorerait '
                                   'la performance.',
                    'explanation_en': 'States 34,500/43,350/37,850. Structural margins 1,500/1,650/1,150; '
                                      'performance margin 450. The minimum is 450; choosing 1,150 ignores '
                                      'performance.',
                    'use_exercise_svg': False,
                    'tags': ['ATPL', 'exam-mcq', 'exam-subject:031'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/sites/faa.gov/files/2023-09/Weight_Balance_Handbook.pdf'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 031'},
                   {'concept_id': 'ATPL_031_001_C02',
                    'concept_title': 'Vérifier chaque limitation',
                    'level': 3,
                    'question': 'ZFM 34 500, bloc 9 000, roulage 150, trajet 5 500 kg. Ajouter 1 600 kg de bloc, '
                                'consommations inchangées. Plafonds ZFM/décollage/arrivée 36 000/45 000/39 000. '
                                'Retrait minimal de charge selon ces trois plafonds ?',
                    'question_en': 'ZFM 34,500, block 9,000, taxi 150, trip 5,500 kg. Add 1,600 kg block fuel '
                                   'with unchanged consumption. ZFM/take-off/landing ceilings '
                                   '36,000/45,000/39,000. Minimum load removal under these three ceilings?',
                    'options': ['1 150 kg', '1 600 kg', '450 kg', '50 kg'],
                    'options_en': ['1,150 kg', '1,600 kg', '450 kg', '50 kg'],
                    'correct_answer': '450 kg',
                    'explanation': 'Avec ajout : décollage 44 950, arrivée 39 450, ZFM inchangée. Il faut '
                                   "enlever 450 kg de charge pour l'arrivée ; 50 est la marge de décollage, pas "
                                   'le retrait requis. Enlever tout le carburant ajouté comme charge serait '
                                   'excessif.',
                    'explanation_en': 'After addition: take-off 44,950, landing 39,450, ZFM unchanged. Remove '
                                      '450 kg load for landing; 50 is take-off margin, not required removal. '
                                      'Removing load equal to all added fuel would be excessive.',
                    'use_exercise_svg': False,
                    'tags': ['ATPL', 'exam-mcq', 'exam-subject:031'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/sites/faa.gov/files/2023-09/Weight_Balance_Handbook.pdf'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 031'},
                   {'concept_id': 'ATPL_031_001_C02',
                    'concept_title': 'Vérifier chaque limitation',
                    'level': 3,
                    'question': 'Marges ZFM/décollage/arrivée : 1 500/1 650/1 150 kg. Carburants inchangés, '
                                'aucun autre plafond fourni : ajout maximal de charge selon ces marges ?',
                    'question_en': 'ZFM/take-off/landing margins: 1,500/1,650/1,150 kg. Fuel unchanged, no other '
                                   'ceiling provided: maximum added load under these margins?',
                    'options': ['4 300 kg', '1 150 kg', '1 500 kg', '1 650 kg'],
                    'options_en': ['4,300 kg', '1,150 kg', '1,500 kg', '1,650 kg'],
                    'correct_answer': '1 150 kg',
                    'explanation': 'Chaque kilogramme ajouté doit satisfaire les trois plafonds ; retenir 1 150. '
                                   'La somme 4 300 cumule des marges qui ne sont pas interchangeables ; les deux '
                                   "autres choix excèdent la marge d'arrivée.",
                    'explanation_en': 'Every added kilogram must meet all three ceilings; choose 1,150. Summing '
                                      'to 4,300 treats non-interchangeable margins as additive; the other values '
                                      'exceed landing margin.',
                    'use_exercise_svg': False,
                    'tags': ['ATPL', 'exam-mcq', 'exam-subject:031'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/sites/faa.gov/files/2023-09/Weight_Balance_Handbook.pdf'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 031'},
                   {'concept_id': 'ATPL_031_001_C02',
                    'concept_title': 'Vérifier chaque limitation',
                    'level': 3,
                    'question': 'Masse au décollage 47 800 kg, maximum structural 49 000 kg, limite de '
                                'performance du cas 46 500 kg. Conclusion ?',
                    'question_en': 'Take-off mass 47,800 kg, structural maximum 49,000 kg, case performance '
                                   'limit 46,500 kg. Conclusion?',
                    'options': ['Maximum structural respecté, limite de performance dépassée',
                                'Les deux limites sont respectées',
                                'Maximum structural dépassé, limite de performance respectée',
                                'La moyenne des deux limites autorise le départ'],
                    'options_en': ['Structural maximum met, performance limit exceeded',
                                   'Both limits are met',
                                   'Structural maximum exceeded, performance limit met',
                                   'The average of both limits permits departure'],
                    'correct_answer': 'Maximum structural respecté, limite de performance dépassée',
                    'explanation': 'La limite de performance 46 500 kg est plus restrictive que le maximum '
                                   'structural 49 000 kg. 47 800 kg respecte ce dernier mais dépasse la limite '
                                   'de performance de 1 300 kg.',
                    'explanation_en': 'The 46,500 kg performance limit is more restrictive than the 49,000 kg '
                                      'structural maximum. 47,800 kg meets the latter but exceeds the '
                                      'performance limit by 1,300 kg.',
                    'use_exercise_svg': False,
                    'tags': ['ATPL', 'exam-mcq', 'exam-subject:031'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/sites/faa.gov/files/2023-09/Weight_Balance_Handbook.pdf'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 031'},
                   {'concept_id': 'ATPL_031_001_C03',
                    'concept_title': 'Recalculer le centrage après consommation',
                    'level': 3,
                    'question': 'État initial 10 000 kg au CG 5,00 m. On déplace 200 kg du bras 7,00 m au bras '
                                '3,00 m, puis consomme 1 000 kg au bras constant 6,00 m. Même datum, aucun autre '
                                'changement : CG final arrondi à 0,001 m ?',
                    'question_en': 'Initial state 10,000 kg at CG 5.00 m. Move 200 kg from arm 7.00 m to 3.00 m, '
                                   'then burn 1,000 kg at constant arm 6.00 m. Common datum, no other change: '
                                   'final CG rounded to 0.001 m?',
                    'options': ['4,889 m', '4,320 m', '4,978 m', '4,800 m'],
                    'options_en': ['4.889 m', '4.320 m', '4.978 m', '4.800 m'],
                    'correct_answer': '4,800 m',
                    'explanation': 'Moment initial 50 000 ; déplacement 200 × (3 − 7) = −800 ; consommation −6 '
                                   '000 kg·m. Moment final 43 200, masse 9 000 : CG 4,800 m. 4,889 oublie le '
                                   'déplacement, 4,320 conserve la masse initiale, 4,978 inverse le signe du '
                                   'déplacement.',
                    'explanation_en': 'Initial moment 50,000; shift 200 × (3 − 7) = −800; fuel burn −6,000 kg·m. '
                                      'Final moment 43,200, mass 9,000: CG 4.800 m. 4.889 ignores the shift, '
                                      '4.320 retains initial mass, 4.978 reverses the shift sign.',
                    'use_exercise_svg': False,
                    'explanation_illustration': '/illustrations/mass-balance-method.svg',
                    'tags': ['ATPL', 'exam-mcq', 'exam-subject:031'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/sites/faa.gov/files/2023-09/Weight_Balance_Handbook.pdf'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 031'},
                   {'concept_id': 'ATPL_031_001_C03',
                    'concept_title': 'Recalculer le centrage après consommation',
                    'level': 3,
                    'question': 'État initial 10 000 kg au CG 5,00 m. On consomme 500 kg d’un réservoir au bras '
                                'constant 6,00 m et 500 kg d’un autre au bras constant 4,00 m. Même datum, aucun '
                                'autre changement : CG final ?',
                    'question_en': 'Initial state 10,000 kg at CG 5.00 m. Burn 500 kg from a tank at constant '
                                   'arm 6.00 m and 500 kg from another at constant arm 4.00 m. Common datum, no '
                                   'other change: final CG?',
                    'options': ['5,111 m', '4,500 m', '5,000 m', '4,889 m'],
                    'options_en': ['5.111 m', '4.500 m', '5.000 m', '4.889 m'],
                    'correct_answer': '5,000 m',
                    'explanation': 'Moment retiré : 500 × 6 + 500 × 4 = 5 000 kg·m. CG = (50 000 − 5 000)/(10 '
                                   '000 − 1 000) = 5,000 m. Le bras moyen du carburant consommé vaut le CG '
                                   'initial. Les valeurs 4,889 et 5,111 supposent à tort toute la consommation '
                                   'dans un seul réservoir.',
                    'explanation_en': 'Removed moment: 500 × 6 + 500 × 4 = 5,000 kg·m. CG = (50,000 − '
                                      '5,000)/(10,000 − 1,000) = 5.000 m. Average arm of burned fuel equals '
                                      'initial CG. 4.889 and 5.111 incorrectly put all fuel burn in one tank.',
                    'use_exercise_svg': False,
                    'explanation_illustration': '/illustrations/mass-balance-method.svg',
                    'tags': ['ATPL', 'exam-mcq', 'exam-subject:031'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/sites/faa.gov/files/2023-09/Weight_Balance_Handbook.pdf'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 031'},
                   {'concept_id': 'ATPL_031_001_C03',
                    'concept_title': 'Recalculer le centrage après consommation',
                    'level': 3,
                    'question': 'Enveloppe fictive jointe : masse 8 000–12 000 kg, limite avant linéaire de 15 à '
                                '25 %MAC, arrière 35 %. À 10 000 kg et 19 %MAC, quel diagnostic ?',
                    'question_en': 'Attached fictional envelope: mass 8,000–12,000 kg, forward limit linear from '
                                   '15 to 25% MAC, aft limit 35%. At 10,000 kg and 19% MAC, what follows?',
                    'options': ['Masse excessive de 2 000 kg',
                                'Masse admise, CG un point de pourcentage trop avant',
                                'Conforme parce que 19 est entre 15 et 35',
                                'CG un point de pourcentage trop arrière'],
                    'options_en': ['Mass excessive by 2,000 kg',
                                   'Mass permitted, CG one percentage point too far forward',
                                   'Compliant because 19 is between 15 and 35',
                                   'CG one percentage point too far aft'],
                    'correct_answer': 'Masse admise, CG un point de pourcentage trop avant',
                    'explanation': 'À 10 000 kg, fraction (10 000−8 000)/4 000 = 0,5 ; limite avant 15+0,5×10 = '
                                   '20 %. Le CG 19 % est un point trop avant. Utiliser 15 % lirait la limite '
                                   "d'une autre masse.",
                    'explanation_en': 'At 10,000 kg, fraction (10,000−8,000)/4,000 = 0.5; forward limit '
                                      '15+0.5×10 = 20%. CG 19% is one point too far forward. Using 15% reads the '
                                      'limit for a different mass.',
                    'use_exercise_svg': True,
                    'explanation_illustration': '/illustrations/atpl-031-001.svg',
                    'tags': ['ATPL', 'exam-mcq', 'exam-subject:031'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/sites/faa.gov/files/2023-09/Weight_Balance_Handbook.pdf'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 031'},
                   {'concept_id': 'ATPL_031_001_C03',
                    'concept_title': 'Recalculer le centrage après consommation',
                    'level': 3,
                    'question': 'Même enveloppe fictive. Départ 10 000 kg, CG 4,60 m ; LEMAC 4 m, MAC 2 m. '
                                "Consommer d'abord 1 000 kg au bras 2 m, puis 1 000 kg au bras 7 m ; bras "
                                'constants. Quel état est hors enveloppe parmi départ, après premier réservoir, '
                                'fin ?',
                    'question_en': 'Same fictional envelope. Start 10,000 kg, CG 4.60 m; LEMAC 4 m, MAC 2 m. '
                                   'Burn 1,000 kg at arm 2 m, then 1,000 kg at arm 7 m; constant arms. Which of '
                                   'start, after first tank and final states is outside?',
                    'options': ['Seulement après le premier réservoir',
                                'Seulement le départ',
                                "Seulement l'état final",
                                'Aucun, puisque départ et fin sont conformes'],
                    'options_en': ['Only after the first tank',
                                   'Only the start',
                                   'Only the final state',
                                   'None, because start and finish comply'],
                    'correct_answer': 'Seulement après le premier réservoir',
                    'explanation': 'Départ : 30 % à 10 000 kg, admis. Après premier retrait : 44 000/9 000 = '
                                   '4,889 m, soit 44,44 %, au-delà de 35. Fin : 37 000/8 000 = 4,625 m, soit '
                                   '31,25 %, admis. Des extrémités conformes masquent ici un dépassement '
                                   'intermédiaire.',
                    'explanation_en': 'Start: 30% at 10,000 kg, inside. After first burn: 44,000/9,000 = 4.889 '
                                      'm, or 44.44%, beyond 35. Final: 37,000/8,000 = 4.625 m, or 31.25%, '
                                      'inside. Compliant endpoints conceal an intermediate exceedance.',
                    'use_exercise_svg': True,
                    'explanation_illustration': '/illustrations/atpl-031-001.svg',
                    'tags': ['ATPL', 'exam-mcq', 'exam-subject:031'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/sites/faa.gov/files/2023-09/Weight_Balance_Handbook.pdf'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 031'},
                   {'concept_id': 'ATPL_031_001_C04',
                    'concept_title': 'Exprimer le centrage en pourcentage de corde',
                    'level': 3,
                    'question': 'CG 3,65 m, LEMAC 3,00 m, longueur MAC 2,60 m. Position en %MAC ?',
                    'question_en': 'CG 3.65 m, LEMAC 3.00 m, MAC length 2.60 m. Position in %MAC?',
                    'options': ['140,4 %', '65 %', '0,25 %', '25 %'],
                    'options_en': ['140.4%', '65%', '0.25%', '25%'],
                    'correct_answer': '25 %',
                    'explanation': '100 × (3,65−3,00)/2,60 = 25 %. 140,4 omet LEMAC ; 65 omet la division par la '
                                   'longueur ; 0,25 omet la conversion du rapport en pourcentage.',
                    'explanation_en': '100 × (3.65−3.00)/2.60 = 25%. 140.4 omits LEMAC; 65 omits division by '
                                      'length; 0.25 omits conversion of ratio to percentage.',
                    'use_exercise_svg': False,
                    'tags': ['ATPL', 'exam-mcq', 'exam-subject:031'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/sites/faa.gov/files/2023-09/Weight_Balance_Handbook.pdf'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 031'},
                   {'concept_id': 'ATPL_031_001_C04',
                    'concept_title': 'Exprimer le centrage en pourcentage de corde',
                    'level': 3,
                    'question': 'Enveloppe fictive jointe. Masse 9 000 kg, moment 41 400 kg·m, LEMAC 4 m, MAC 2 '
                                'm. Quelle conclusion correspond à ce seul contrôle masse/centrage ?',
                    'question_en': 'Attached fictional envelope. Mass 9,000 kg, moment 41,400 kg·m, LEMAC 4 m, '
                                   'MAC 2 m. What follows from this mass/CG check alone?',
                    'options': ['CG 30 %MAC, trop avant car limite avant 35 %',
                                'CG 20 %MAC, conforme à toutes les exigences du vol',
                                'CG 30 %MAC, dans les limites 17,5–35 % de cette masse',
                                'CG 230 %MAC, trop arrière'],
                    'options_en': ['CG 30% MAC, too far forward because forward limit is 35%',
                                   'CG 20% MAC, compliant with every flight requirement',
                                   "CG 30% MAC, within this mass's 17.5–35% limits",
                                   'CG 230% MAC, too far aft'],
                    'correct_answer': 'CG 30 %MAC, dans les limites 17,5–35 % de cette masse',
                    'explanation': 'Bras = 41 400/9 000 = 4,60 m ; %MAC = 100×0,60/2 = 30. Limite avant à 9 000 '
                                   ': 15+1 000/4 000×10 = 17,5 %. Masse et CG sont admis dans cette enveloppe ; '
                                   'performances, carburant et états ultérieurs restent des contrôles distincts.',
                    'explanation_en': 'Arm = 41,400/9,000 = 4.60 m; %MAC = 100×0.60/2 = 30. Forward limit at '
                                      '9,000: 15+1,000/4,000×10 = 17.5%. Mass and CG meet this envelope; '
                                      'performance, fuel and later states remain separate checks.',
                    'use_exercise_svg': True,
                    'explanation_illustration': '/illustrations/atpl-031-001.svg',
                    'tags': ['ATPL', 'exam-mcq', 'exam-subject:031'],
                    'practice_format': 'mcq',
                    'source_references': ['https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?page=8',
                                          'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67',
                                          'https://www.faa.gov/sites/faa.gov/files/2023-09/Weight_Balance_Handbook.pdf'],
                    'syllabus_reference': 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b) — Subject 031'}]}
    save(data)
