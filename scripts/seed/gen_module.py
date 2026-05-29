#!/usr/bin/env python3
"""Generate the complete CRM training module JSON file."""
import json

OUT = "c:/Users/sarah/Documents/Codium/aeropath/scripts/seed/module_crm_decision.json"

d = {
    "module_title": "Gestion des Ressources de l'Equipage (CRM) et Prise de Decision Aeronautique",
    "target_level": "PPL/CPL",
    "training_goal": "Developper les competences non techniques essentielles du pilote : communication efficace, gestion des ressources, conscience situationnelle, prise de decision structuree et prevention des erreurs humaines. Ce module vise a reduire les risques operationnels en formant le pilote a reconnaitre les pieges cognitifs et a appliquer des processus decisionnels robustes.",
    "concepts": [],
    "operational_rules": [],
    "evaluation": [],
    "summary": "Ce module couvre les fondamentaux du Crew Resource Management (CRM) et de la prise de decision aeronautique. Les 12 concepts cles abordent la communication, la conscience situationnelle, le modele TEM, les biais cognitifs, le modele DECIDE, le gradient d'autorite, la gestion de la charge de travail, les pieges de l'automatisation, les briefings structures, les attitudes dangereuses et la gestion du stress. Les 8 regles operationnelles fournissent des procedures applicables immediatement en vol : boucle fermee, briefings obligatoires, regle des 5 secondes, cross-check, GO/NO-GO, sterile cockpit, consensus eclaire et IMSAFE. L'accent est mis sur la securite des vols, la conscience situationnelle et la prevention des erreurs humaines a travers 5 questions d'evaluation adaptatives."
}

# Build concepts
d["concepts"] = [
    {
        "concept_id": "CRM_001",
        "title": "Definition et Origine du CRM",
        "content": "Le Crew Resource Management (CRM) est l'utilisation efficace de toutes les ressources disponibles - equipage, materiel, informations et procedures - pour atteindre un vol sur et efficace. Ne dans les annees 1970 suite a l'analyse des accidents majeurs (Tenerife 1977, United 173), le CRM est passe d'une approche centree sur le commandant de bord a une gestion collaborative des ressources. Les piliers du CRM sont : la communication, le leadership, la prise de decision, la gestion des priorites et la conscience situationnelle.",
        "difficulty": "facile",
        "risk_level": "faible",
        "phase_of_flight": ["toutes phases"],
        "tags": ["CRM", "facteurs humains", "securite", "gestion d'equipage", "non-technical skills"],
        "common_errors": [
            "Confondre CRM avec simple politesse en cabine",
            "Croire que le CRM ne s'applique qu'aux equipages multi-pilotes",
            "Negliger le CRM en vol solo alors que les ressources ATC et systemes restent disponibles"
        ],
        "references": ["ICAO Doc 9683 - Human Factors Training Manual", "EASA Part-FCL - Human Performance and Limitations"]
    },
    {
        "concept_id": "CRM_002",
        "title": "Les Trois Niveaux de Communication en Cockpit",
        "content": "La communication en aviation opere a trois niveaux distincts : (1) Communication interpersonnelle entre pilotes, entre pilote et ATC, entre pilote et passagers ; (2) Communication intra-personnelle (dialogue interieur, auto-briefing) ; (3) Communication homme-machine via les interfaces du cockpit (PFD, ND, FMS, alarmes). Chaque niveau peut etre source d'erreurs : ambiguite phraseologique, mauvaise interpretation des alarmes, ou absence de briefings structures. La regle d'or est la boucle fermee (closed-loop) : emission, reception, readback, confirmation.",
        "difficulty": "moyen",
        "risk_level": "moyen",
        "phase_of_flight": ["toutes phases"],
        "tags": ["communication", "CRM", "boucle fermee", "phraseologie", "briefing"],
        "common_errors": [
            "Supposer que le message a ete compris sans confirmation",
            "Utiliser un langage non standard en situation stressante",
            "Interrompre un membre d'equipage en phase critique"
        ],
        "references": ["FAA AC 120-51E - Crew Resource Management Training", "NTSB Safety Report - Communication in Aviation"]
    },
    {
        "concept_id": "CRM_003",
        "title": "La Conscience Situationnelle (SA)",
        "content": "La conscience situationnelle est la perception exacte des elements de l'environnement dans le temps et l'espace, la comprehension de leur signification et la projection de leur statut futur. Selon le modele d'Endsley (1995), elle comporte trois niveaux : Niveau 1 (Perception) - capter les donnees brutes (altitude, vitesse, meteo) ; Niveau 2 (Comprehension) - integrer ces donnees en une image coherente ; Niveau 3 (Projection) - anticiper l'evolution. La perte de SA est la cause contributive majeure dans plus de 70% des accidents aeriens.",
        "difficulty": "moyen",
        "risk_level": "eleve",
        "phase_of_flight": ["toutes phases", "approche", "descente"],
        "tags": ["conscience situationnelle", "SA", "Endsley", "anticipation", "securite"],
        "common_errors": [
            "Fixation sur un seul instrument au detriment de la vue d'ensemble",
            "Confiance excessive dans l'automatisation sans recoupement",
            "Ne pas remettre en cause une image mentale erronee face a des donnees contradictoires"
        ],
        "references": ["Endsley, M.R. (1995) - Toward a Theory of Situation Awareness", "Skybrary - Situation Awareness"]
    },
    {
        "concept_id": "CRM_004",
        "title": "Le Modele TEM (Threat and Error Management)",
        "content": "Le Threat and Error Management (TEM) est un cadre conceptuel developpe par l'Universite du Texas et adopte par l'OACI pour comprendre comment les equipages gerent les menaces et les erreurs. Les menaces sont des evenements externes hors du controle de l'equipage (meteo, trafic, panne technique). Les erreurs sont des actions ou inactions de l'equipage qui s'ecartent des intentions ou des normes. Le TEM distingue trois types d'erreurs : erreurs de procedure, erreurs de communication et erreurs de pilotage. L'objectif n'est pas d'eviter toute erreur (impossible), mais de les detecter et les rattraper avant qu'elles ne deviennent des etats indesirables (Undesired Aircraft State) menant a un accident.",
        "difficulty": "moyen",
        "risk_level": "moyen",
        "phase_of_flight": ["toutes phases"],
        "tags": ["TEM", "gestion des menaces", "gestion des erreurs", "securite", "resilience"],
        "common_errors": [
            "Considerer les menaces comme des excuses plutot que des facteurs a anticiper",
            "Ne pas reconnaitre une erreur par orgueil professionnel",
            "Confondre erreur recuperee avec erreur sans consequence"
        ],
        "references": ["ICAO Doc 9859 - Safety Management Manual", "University of Texas - LOSA (Line Operations Safety Audit)"]
    },
    {
        "concept_id": "CRM_005",
        "title": "Les Biais Cognitifs en Prise de Decision",
        "content": "Les biais cognitifs sont des raccourcis mentaux (heuristiques) qui deforment le jugement du pilote. Les plus dangereux en aviation sont : (1) Biais de confirmation - chercher les informations qui confirment notre decision plutot que celles qui l'infirment ; (2) Biais d'ancrage - se fixer sur une premiere information et ne pas assez reviser son jugement ; (3) Biais de disponibilite - surestimer la probabilite d'evenements recents ou marquants ; (4) Biais d'optimisme - sous-estimer les risques ; (5) Pression de planification (plan continuation bias) - persister dans un plan initial malgre des signaux d'alarme. La reconnaissance de ces biais est la premiere etape pour les contrer.",
        "difficulty": "difficile",
        "risk_level": "eleve",
        "phase_of_flight": ["toutes phases", "approche", "decollage"],
        "tags": ["biais cognitifs", "prise de decision", "psychologie", "erreurs humaines", "ADM"],
        "common_errors": [
            "Continuer une approche instable par pression d'arriver a l'heure",
            "Ignorer les bulletins meteo defavorables parce qu'ils contredisent le plan de vol",
            "Rationaliser une decision risquee a posteriori"
        ],
        "references": ["Kahneman, D. (2011) - Thinking, Fast and Slow", "FAA - Aeronautical Decision Making (ADM) Handbook"]
    },
    {
        "concept_id": "CRM_006",
        "title": "Le Modele DECIDE de Prise de Decision",
        "content": "Le modele DECIDE est un processus structure en six etapes pour la prise de decision aeronautique : D (Detect) - Detecter qu'un changement ou une situation anormale s'est produit ; E (Estimate) - Estimer la signification et l'impact de ce changement sur la securite du vol ; C (Choose) - Choisir parmi les options disponibles la plus appropriee ; I (Identify) - Identifier les actions correctives necessaires ; D (Do) - Executer les actions choisies ; E (Evaluate) - Evaluer les resultats et ajuster si necessaire. Ce modele force le pilote a sortir d'une reaction instinctive pour adopter une approche analytique, particulierement utile en situation de stress ou d'incertitude.",
        "difficulty": "facile",
        "risk_level": "faible",
        "phase_of_flight": ["toutes phases"],
        "tags": ["DECIDE", "prise de decision", "ADM", "processus", "analyse"],
        "common_errors": [
            "Sauter l'etape d'estimation et passer directement a l'action",
            "Ne pas evaluer le resultat de la decision prise",
            "Appliquer DECIDE de maniere rigide sans tenir compte de l'urgence"
        ],
        "references": ["FAA-H-8083-25 - Pilot's Handbook of Aeronautical Knowledge", "Jeppesen - Private Pilot Manual"]
    },
    {
        "concept_id": "CRM_007",
        "title": "Le Gradient d'Autorite et l'Assertivite",
        "content": "Le gradient d'autorite est la difference hierarchique percue entre les membres d'un equipage. Un gradient trop raide (commandant autoritaire, copilote passif) inhibe la communication et empeche la remontee d'informations critiques. Un gradient trop plat peut mener a des conflits ou a une perte de leadership. Le concept d'assertivite - la capacite a exprimer son opinion de maniere professionnelle sans agressivite ni passivite - est essentiel. Le copilote doit oser dire \"Je ne suis pas d'accord\" ou \"Je pense que nous devrions verifier\" lorsque la securite est en jeu. La culture juste (just culture) encourage ce type de feedback sans crainte de represailles.",
        "difficulty": "moyen",
        "risk_level": "moyen",
        "phase_of_flight": ["toutes phases"],
        "tags": ["gradient d'autorite", "assertivite", "leadership", "just culture", "CRM"],
        "common_errors": [
            "Un copilote qui n'ose pas corriger une erreur du commandant de bord",
            "Un commandant qui rejette systematiquement les suggestions du copilote",
            "Confondre assertivite avec agressivite ou insubordination"
        ],
        "references": ["Helmreich, R.L. and Merritt, A.C. (1998) - Culture at Work in Aviation and Medicine", "EASA - Just Culture Principles"]
    },
    {
        "concept_id": "CRM_008",
        "title": "La Gestion de la Charge de Travail et des Priorites",
        "content": "La gestion de la charge de travail repose sur la capacite a prioriser les taches selon leur criticite et leur urgence. La matrice d'Eisenhower (urgent/important) s'applique au cockpit : priorite 1 - taches critiques pour la securite immediate (eviter un obstacle, gerer une panne) ; priorite 2 - taches importantes mais non urgentes (briefing approche, check-list) ; priorite 3 - taches urgentes mais non critiques (changer frequence radio) ; priorite 4 - taches ni urgentes ni critiques (reglage confort). La regle des \"mains sur le manche, tete dans le cockpit\" rappelle que la priorite absolue est le pilotage de l'aeronef. La delegation et la repartition des taches sont des competences CRM cles.",
        "difficulty": "moyen",
        "risk_level": "eleve",
        "phase_of_flight": ["approche", "decollage", "descente", "phases critiques"],
        "tags": ["charge de travail", "priorites", "gestion des taches", "delegation", "workload"],
        "common_errors": [
            "Se laisser submerger par des taches secondaires en phase critique",
            "Tout faire soi-meme sans deleguer",
            "Sous-estimer le temps necessaire pour accomplir une tache"
        ],
        "references": ["Reason, J. (1990) - Human Error", "ICAO - Human Factors Guidelines for Aircraft Maintenance"]
    },
    {
        "concept_id": "CRM_009",
        "title": "Les Pieges de l'Automatisation et la Complaisance",
        "content": "L'automatisation des cockpits modernes (FMS, pilote automatique, gestion des alarmes) apporte des benefices indeniables en reduisant la charge de travail, mais introduit des risques specifiques : (1) Complaisance - surestimer la fiabilite du systeme et baisser la vigilance ; (2) Perte de competences manuelles par manque de pratique ; (3) Confusion de mode (mode error) - croire que le systeme est dans un mode alors qu'il est dans un autre ; (4) Surprise par l'automatisation - le systeme fait quelque chose d'inattendu. Le pilote doit rester \"dans la boucle\" (in the loop), surveiller activement l'automatisation et etre pret a reprendre le controle manuellement a tout moment.",
        "difficulty": "difficile",
        "risk_level": "eleve",
        "phase_of_flight": ["croisiere", "approche", "montee"],
        "tags": ["automatisation", "complaisance", "mode error", "vigilance", "technologie"],
        "common_errors": [
            "Ne pas verifier manuellement les donnees entrees dans le FMS",
            "Laisser le pilote automatique gerer une situation anormale sans surveillance",
            "Perdre ses reflexes de pilotage manuel par exces d'automatisation"
        ],
        "references": ["Boeing - 737 NG FCTM (Flight Crew Training Manual)", "Parasuraman, R. and Riley, V. (1997) - Humans and Automation"]
    },
    {
        "concept_id": "CRM_010",
        "title": "Le Briefing et le Debriefing Structures",
        "content": "Le briefing est un outil CRM fondamental qui permet de creer une image mentale partagee (shared mental model) avant chaque phase de vol. Un briefing efficace suit la structure \"Qui, Quoi, Ou, Quand, Comment, Alternatives\" : qui fait quoi, quels sont les risques identifies, quelle est la strategie, que faire en cas d'imprevu. Le briefing n'est pas une lecture monotone de la check-list mais un echange actif. Le debriefing apres le vol est tout aussi important : il permet d'analyser ce qui s'est bien passe, ce qui peut etre ameliore, et de capitaliser sur l'experience. Un debriefing constructif se concentre sur les comportements, pas sur les personnes.",
        "difficulty": "facile",
        "risk_level": "faible",
        "phase_of_flight": ["toutes phases", "pre-vol", "post-vol"],
        "tags": ["briefing", "debriefing", "communication", "preparation", "retour d'experience"],
        "common_errors": [
            "Faire un briefing trop long et peu cible",
            "Lire la check-list sans reflexion ni anticipation",
            "Sauter le debriefing par manque de temps ou par fatigue"
        ],
        "references": ["FAA - Risk Management Handbook", "Airbus - Flight Operations Briefing Notes"]
    },
    {
        "concept_id": "CRM_011",
        "title": "Les 5 Attitudes Dangereuses du Pilote",
        "content": "La FAA identifie cinq attitudes dangereuses qui compromettent la securite des vols : (1) Anti-autorite - \"Ne me dites pas quoi faire\" (refus des regles et procedures) ; (2) Impulsivite - \"Faisons-le vite\" (decision sans reflexion) ; (3) Invulnerabilite - \"Ca n'arrive qu'aux autres\" (sous-estimation des risques) ; (4) Macho - \"Je peux le faire\" (besoin de prouver sa valeur par des risques inutiles) ; (5) Resignation - \"A quoi bon ?\" (passivite face aux difficultes). Chaque attitude dangereuse a une antidote specifique : par exemple, l'antidote a l'invulnerabilite est \"Ca pourrait m'arriver\". La reconnaissance de ces attitudes chez soi est une competence metacognitive essentielle.",
        "difficulty": "facile",
        "risk_level": "moyen",
        "phase_of_flight": ["toutes phases"],
        "tags": ["attitudes dangereuses", "psychologie", "ADM", "securite", "metacognition"],
        "common_errors": [
            "Ne pas reconnaitre sa propre attitude dangereuse",
            "Confondre confiance legitime et attitude macho",
            "Laisser la resignation empecher une action corrective necessaire"
        ],
        "references": ["FAA-H-8083-25 - Pilot's Handbook of Aeronautical Knowledge", "FAA - Aeronautical Decision Making"]
    },
    {
        "concept_id": "CRM_012",
        "title": "La Gestion du Stress et des Emotions en Vol",
        "content": "Le stress est une reponse physiologique et psychologique a une demande percue comme depassant les capacites disponibles. La loi de Yerkes-Dodson montre qu'un stress modere ameliore la performance (eustress), mais qu'un stress excessif la degrade (distress). Les symptomes du distress en vol incluent : vision tunnel, fixation sur une tache, oubli des procedures de base, irritabilite, decisions precipitees. Les techniques de gestion du stress comprennent : la respiration controlee (4-7-8), la priorisation des taches, la verbalisation des actions (think aloud), et l'utilisation systematique des check-lists. La gestion des emotions - peur, frustration, colere - est tout aussi cruciale : reconnaitre l'emotion, la nommer, et ne pas la laisser dicter les decisions.",
        "difficulty": "moyen",
        "risk_level": "eleve",
        "phase_of_flight": ["urgence", "approche", "decollage", "phases critiques"],
        "tags": ["stress", "emotions", "Yerkes-Dodson", "gestion du stress", "resilience"],
        "common_errors": [
            "Nier son stress et continuer comme si de rien n'etait",
            "Prendre une decision sous le coup de la frustration",
            "Sous-estimer l'impact du stress chronique (fatigue, problemes personnels)"
        ],
        "references": ["Yerkes, R.M. and Dodson, J.D. (1908) - The Relation of Strength of Stimulus to Rapidity of Habit-Formation", "ICAO - Human Performance and Limitations"]
    }
]

# Build operational rules
d["operational_rules"] = [
    {
        "rule_id": "OPR_CRM_001",
        "title": "Regle de la Boucle Fermee en Communication",
        "procedure": "Toute communication critique (instruction ATC, modification de cap, altitude, frequence) doit suivre le cycle : 1) Emetteur formule le message clairement ; 2) Recepteur ecoute activement et repete (readback) ; 3) Emetteur confirme la repetition correcte (acknowledgment). En cas de doute, utiliser \"Say again\" ou \"Confirm\". Ne jamais supposer que le message a ete compris sans confirmation explicite.",
        "conditions": ["toute communication avec ATC", "briefing equipage", "modification de plan de vol", "instruction de securite"],
        "common_errors": [
            "Readback partiel sans les elements critiques",
            "Oublier de confirmer le readback",
            "Acquiescer par habitude sans ecouter reellement"
        ],
        "safety_notes": [
            "Les etudes NTSB montrent que 30% des erreurs de communication impliquent un readback incorrect",
            "En environnement international, verifier que l'anglais est compris par tous",
            "En situation stressante, ralentir le debit et articuler"
        ]
    },
    {
        "rule_id": "OPR_CRM_002",
        "title": "Regle des 3 Briefings Obligatoires",
        "procedure": "Trois briefings structures sont obligatoires a chaque vol : 1) Briefing pre-vol (avant mise en route) : meteo, NOTAM, carburant, route, masses, risques identifies, roles de chacun, decision GO/NO-GO ; 2) Briefing approche (avant la descente) : type d'approche, minima, configuration, procedure de degagement, gestion des pannes ; 3) Briefing depart (avant decollage) : piste, procedure de depart, altitude initiale, procedure en cas de panne apres V1. Chaque briefing doit etre interactif et permettre les questions.",
        "conditions": ["avant chaque vol", "avant chaque approche", "avant chaque decollage"],
        "common_errors": [
            "Briefing expeditif sans reelle anticipation",
            "Briefing unidirectionnel sans participation de l'equipage",
            "Omettre le briefing approche sur un terrain connu"
        ],
        "safety_notes": [
            "Un briefing de qualite prend 2 a 5 minutes selon la complexite",
            "Le briefing doit etre adapte aux conditions du jour, pas une routine",
            "Encourager le copilote ou l'eleve a poser des questions"
        ]
    },
    {
        "rule_id": "OPR_CRM_003",
        "title": "Regle des 5 Secondes pour la Prise de Decision",
        "procedure": "Face a une situation anormale ou une urgence, appliquer la regle des 5 secondes : 1) Arreter toute action en cours ; 2) Respirer profondement (4 secondes inspiration, 4 secondes expiration) ; 3) Evaluer la situation (quoi, ou, quand, gravite) ; 4) Decider de l'action appropriee ; 5) Agir. Cette regle brise le cycle de la panique et de l'impulsivite. En equipage, le pilote non-manoeuvrant (PNF) annonce \"Stop, evaluons la situation\" si le PF semble precipite.",
        "conditions": ["toute situation anormale", "alarme soudaine", "panne en vol", "conflit de trafic"],
        "common_errors": [
            "Agir immediatement sans evaluation (impulsivite)",
            "Rester fige sans decider (resignation)",
            "Discuter trop longtemps sans agir (paralysie par l'analyse)"
        ],
        "safety_notes": [
            "Dans la majorite des accidents, l'equipage disposait de plus de temps qu'il ne le pensait",
            "La regle \"Aviate, Navigate, Communicate\" prime : d'abord piloter l'avion",
            "En cas de doute, la decision la plus prudente est generalement la bonne"
        ]
    },
    {
        "rule_id": "OPR_CRM_004",
        "title": "Regle de la Verification Croisee (Cross-Check)",
        "procedure": "Toute action critique sur les systemes de bord doit etre verifiee par un second regard (cross-check). Le pilote qui actionne annonce son intention (\"Je selectionne 5000 pieds\"), le second pilote verifie visuellement et verbalement (\"Confirme, 5000 pieds selectionnes\"). En vol solo, le cross-check est remplace par la verification instrumentale : pointer l'instrument, lire la valeur, verifier la coherence avec les autres instruments. Ne jamais entrer une donnee dans le FMS sans verifier deux fois.",
        "conditions": ["entree de donnees FMS", "selection d'altitude", "changement de frequence", "modification de cap", "configuration aeronef"],
        "common_errors": [
            "Verification de routine sans reelle attention",
            "Confiance excessive dans sa propre saisie",
            "Omettre le cross-check en situation de forte charge de travail"
        ],
        "safety_notes": [
            "Les erreurs de saisie FMS sont une cause frequente de deviation d'altitude",
            "Le cross-check est particulierement critique lors des changements de niveau de vol",
            "En instruction, l'instructeur doit verifier toute action de l'eleve"
        ]
    },
    {
        "rule_id": "OPR_CRM_005",
        "title": "Regle de la Decision GO/NO-GO",
        "procedure": "Avant chaque vol, appliquer formellement la decision GO/NO-GO en evaluant les 5 facteurs PAVE : Pilote (fatigue, maladie, stress, medicaments, competence), Aeronef (documents, carburant, maintenance, equipement), Environnement (meteo, terrain, espace aerien, nuit), Valeur externe (pression des passagers, planning, cout), et Evaluation (synthese des risques). Si un seul facteur est rouge (non acceptable), la decision est NO-GO. La decision GO/NO-GO peut etre revisee a tout moment du vol (decision de deroutement, demi-tour).",
        "conditions": ["avant chaque vol", "avant chaque etape", "en vol si conditions changent"],
        "common_errors": [
            "Pression sociale pour decoller malgre des conditions limites",
            "Auto-evaluation biaisee de sa propre condition physique/mentale",
            "Confondre GO/NO-GO avec \"on verra en vol\""
        ],
        "safety_notes": [
            "La decision NO-GO n'est jamais un echec, c'est une marque de professionnalisme",
            "Les accidents lies a une pression de planification sont parmi les plus evitables",
            "Un pilote qui refuse de voler doit etre soutenu, pas critique"
        ]
    },
    {
        "rule_id": "OPR_CRM_006",
        "title": "Regle de la Gestion des Interruptions",
        "procedure": "Les interruptions en phase critique (decollage, approche, atterrissage) sont une cause majeure d'erreurs. Appliquer le principe \"Sterile Cockpit\" : en dessous de 10 000 pieds (ou en phase critique), seules les communications essentielles a la securite du vol sont autorisees. Si une interruption survient : 1) Accuser reception brievement (\"Compris, je reviens\") ; 2) Terminer la tache en cours ; 3) Revenir a l'interruption. Ne jamais interrompre un pilote qui execute une check-list ou une procedure critique.",
        "conditions": ["decollage et montee initiale", "descente et approche", "atterrissage", "procedure d'urgence"],
        "common_errors": [
            "Repondre a une question non essentielle en phase d'approche",
            "Interrompre la lecture d'une check-list",
            "Multiplier les taches simultanees en phase critique"
        ],
        "safety_notes": [
            "Le concept \"Sterile Cockpit\" est une reglementation FAA (14 CFR 121.542) et une recommandation EASA",
            "Les passagers doivent etre briefes sur les phases de silence cabine",
            "En instruction, l'instructeur doit eviter les commentaires non essentiels en phase critique"
        ]
    },
    {
        "rule_id": "OPR_CRM_007",
        "title": "Regle de la Prise de Decision par Consensus Eclaire",
        "procedure": "En equipage, la decision finale revient au commandant de bord, mais elle doit etre eclairee par l'avis de tous les membres. Processus : 1) Chaque membre expose son analyse et sa recommandation sans filtre hierarchique ; 2) Le commandant synthetise les informations ; 3) Le commandant annonce la decision et la justifie brievement ; 4) Si un membre exprime un desaccord fonde, le commandant reconsidere. En cas de desaccord persistant sur un point de securite, la regle du \"veto securite\" s'applique : tout membre d'equipage peut exiger une pause dans la procedure pour clarifier la situation.",
        "conditions": ["equipage multi-pilotes", "decision importante", "situation anormale", "desaccord potentiel"],
        "common_errors": [
            "Le commandant impose sa decision sans ecouter les autres",
            "Le copilote n'exprime pas son desaccord par peur de la hierarchie",
            "Confondre consensus (accord unanime) avec decision eclairee"
        ],
        "safety_notes": [
            "Les accidents les plus graves impliquent souvent un defaut de communication hierarchique",
            "Un bon commandant cree un climat ou le copilote ose parler",
            "La decision finale reste au commandant, mais eclairee par l'equipage"
        ]
    },
    {
        "rule_id": "OPR_CRM_008",
        "title": "Regle de l'Auto-Evaluation avant Vol (IMSAFE)",
        "procedure": "Avant chaque vol, le pilote doit s'auto-evaluer selon le modele IMSAFE : Illness (suis-je malade ?), Medication (ai-je pris des medicaments ?), Stress (suis-je stresse ?), Alcohol (ai-je consomme de l'alcool dans les 8-12 dernieres heures ?), Fatigue (suis-je fatigue ?), Eating (ai-je bien mange et bu ?). Si une reponse est positive, le pilote doit evaluer si cela affecte sa capacite a piloter en securite. En cas de doute, la decision est NO-GO. Cette auto-evaluation est une responsabilite ethique et reglementaire.",
        "conditions": ["avant chaque vol", "avant chaque etape de vol", "en cas de fatigue en vol"],
        "common_errors": [
            "Minimiser les symptomes par envie de voler",
            "Confondre fatigue normale et fatigue incapacitante",
            "Ne pas reevaluer IMSAFE en cours de journee (vols multiples)"
        ],
        "safety_notes": [
            "La fatigue est responsable de 15 a 20% des accidents aeriens",
            "Les medicaments en vente libre (antihistaminiques, decongestionnants) peuvent alterer les capacites",
            "Un pilote honnete dans son auto-evaluation est un pilote professionnel"
        ]
    }
]

# Build evaluation questions
d["evaluation"] = [
    {
        "question_id": "Q_CRM_001",
        "competency": "CRM - Communication",
        "difficulty": "facile",
        "learning_objective": "Comprendre le principe de la boucle fermee en communication aeronautique",
        "question": 'Un controleur vous donne l instruction : "ABC123, montez au niveau 130, QNH 1015". Quelle est la procedure de communication correcte ?',
        "options": [
            'Repondre "Compris" et monter au FL130',
            'Repondre "ABC123, monte au niveau 130, QNH 1015" et attendre la confirmation du controleur',
            "Monter immediatement au FL130 sans reponse pour liberer la frequence",
            'Repondre "Roger" et regler l altimetre'
        ],
        "correct_answer": 'Repondre "ABC123, monte au niveau 130, QNH 1015" et attendre la confirmation du controleur',
        "explanation": 'La boucle fermee exige un readback complet de l instruction (qui, quoi, valeur) suivi d une confirmation par le controleur. "Compris" ou "Roger" ne sont pas des readbacks valides car ils ne permettent pas au controleur de verifier que l instruction a ete correctement recue.',
        "common_misconception": 'Beaucoup de pilotes pensent que "Roger" ou "Compris" suffisent, mais ces reponses ne confirment pas la bonne reception des donnees chiffrees (altitude, QNH).'
    },
    {
        "question_id": "Q_CRM_002",
        "competency": "CRM - Prise de Decision",
        "difficulty": "moyen",
        "learning_objective": "Appliquer le modele DECIDE face a une situation anormale",
        "question": "En croisiere a 4500 ft, vous remarquez une baisse progressive de la pression d huile de 60 a 35 PSI, avec une temperature d huile qui augmente de 85C a 100C. Selon le modele DECIDE, quelle est la premiere etape a appliquer ?",
        "options": [
            "Choisir immediatement un terrain pour un atterrissage d urgence (etape Choose)",
            "Detecter que la situation a change et estimer la gravite (etapes Detect et Estimate)",
            "Identifier l action corrective en consultant le manuel de vol (etape Identify)",
            "Evaluer le resultat en atterrissant (etape Evaluate)"
        ],
        "correct_answer": "Detecter que la situation a change et estimer la gravite (etapes Detect et Estimate)",
        "explanation": "Le modele DECIDE commence par D