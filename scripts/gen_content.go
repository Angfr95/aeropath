package main

import "fmt"

// generateLessonContentFr generates French content for a lesson based on category and lesson number
func generateLessonContentFr(category string, lessonNum int) string {
	// For lessons 1-5, use the hand-crafted content from seed_content.go
	if lessonNum <= 5 {
		return lessonContentFr(category, lessonNum)
	}
	fmt.Printf("DEBUG generateLessonContentFr called for %s lesson %d\n", category, lessonNum)

	// For lessons 6+, generate structured content automatically
	title := getLessonTitleFr(category, lessonNum)
	keyPoints := getKeyPointsFr(category, lessonNum)
	examples := getExamplesFr(category, lessonNum)
	summary := getSummaryFr(category, lessonNum)

	content := fmt.Sprintf(`## %s

### Objectifs d'apprentissage

Dans cette leçon, vous allez approfondir vos connaissances sur les aspects avancés de %s.

### Points clés

%s

### Exemples pratiques

%s

### Résumé

%s

### Points à retenir

- Cette leçon fait partie du programme de formation continue
- Les concepts abordés sont essentiels pour la progression vers des niveaux supérieurs
- Pratiquez régulièrement avec les exercices associés
- Consultez les références officielles pour plus de détails`,
		title, getCategoryNameFr(category), keyPoints, examples, summary)

	return content
}

// generateLessonContentEn generates English content for a lesson based on category and lesson number
func generateLessonContentEn(category string, lessonNum int) string {
	// For lessons 1-5, use the hand-crafted content from seed_content.go
	if lessonNum <= 5 {
		return lessonContentEn(category, lessonNum)
	}
	fmt.Printf("DEBUG generateLessonContentEn called for %s lesson %d\n", category, lessonNum)

	// For lessons 6+, generate structured content automatically
	title := getLessonTitleEn(category, lessonNum)
	keyPoints := getKeyPointsEn(category, lessonNum)
	examples := getExamplesEn(category, lessonNum)
	summary := getSummaryEn(category, lessonNum)

	content := fmt.Sprintf(`## %s

### Learning Objectives

In this lesson, you will deepen your knowledge of advanced aspects of %s.

### Key Points

%s

### Practical Examples

%s

### Summary

%s

### Key Takeaways

- This lesson is part of the continuing education program
- The concepts covered are essential for progression to higher levels
- Practice regularly with the associated exercises
- Consult official references for more details`,
		title, getCategoryNameEn(category), keyPoints, examples, summary)

	return content
}

func getCategoryNameFr(category string) string {
	names := map[string]string{
		"airlaw":                "le droit aérien",
		"meteorology":           "la météorologie",
		"navigation":            "la navigation",
		"performance":           "les performances",
		"aircraft_general":      "la connaissance générale de l'aéronef",
		"flight_planning":       "la planification de vol",
		"human_performance":     "les facteurs humains",
		"operational_procedures": "les procédures opérationnelles",
		"principles_of_flight":  "les principes du vol",
		"communications":        "les communications",
		"mass_and_balance":      "les masses et centrage",
		"instrumentation":       "l'instrumentation",
	}
	if name, ok := names[category]; ok {
		return name
	}
	return category
}

func getCategoryNameEn(category string) string {
	names := map[string]string{
		"airlaw":                "air law",
		"meteorology":           "meteorology",
		"navigation":            "navigation",
		"performance":           "performance",
		"aircraft_general":      "aircraft general knowledge",
		"flight_planning":       "flight planning",
		"human_performance":     "human performance",
		"operational_procedures": "operational procedures",
		"principles_of_flight":  "principles of flight",
		"communications":        "communications",
		"mass_and_balance":      "mass and balance",
		"instrumentation":       "instrumentation",
	}
	if name, ok := names[category]; ok {
		return name
	}
	return category
}

func getLessonTitleFr(category string, lessonNum int) string {
	titles := map[string]map[int]string{
		"airlaw": {
			6:  "Les accords bilatéraux et la reconnaissance des licences",
			7:  "La gestion du trafic aérien (ATM) et les services de la circulation aérienne (ATS)",
			8:  "Les enquêtes de sécurité aérienne et le reporting d'incidents",
			9:  "La réglementation environnementale et les nuisances sonores",
			10: "Les droits des passagers et la responsabilité du transporteur",
			11: "La certification des aéronefs et le maintien de la navigabilité",
			12: "Les organismes de formation agréés (ATO) et la formation des pilotes",
			13: "La gestion de la sécurité (SMS) dans les organisations aéronautiques",
			14: "Les règles d'exploitation commerciale (OPS) et la gestion des équipages",
			15: "La réglementation des drones et des aéronefs télépilotés",
			16: "Les accords internationaux de partage de codes et alliances",
			17: "La protection des données et la vie privée dans l'aviation",
			18: "Les normes techniques et la certification des équipements",
			19: "La gestion des crises et la continuité des opérations",
			20: "Les évolutions réglementaires récentes et perspectives futures",
		},
		"meteorology": {
			6:  "Les phénomènes météorologiques dangereux : orages et cisaillement",
			7:  "La prévision météorologique et les modèles numériques",
			8:  "Les cartes météorologiques : analyse et interprétation",
			9:  "La météorologie tropicale et les cyclones",
			10: "Les phénomènes optiques dans l'atmosphère",
			11: "La météorologie de montagne et les effets orographiques",
			12: "Les systèmes frontaux et leur évolution",
			13: "La turbulence en air clair (CAT) et les ondes de montagne",
			14: "Le givrage atmosphérique et la prévention",
			15: "Les radars météorologiques et leur utilisation en vol",
			16: "La climatologie et les variations saisonnières",
			17: "Les phénomènes électriques dans l'atmosphère",
			18: "La météorologie arctique et les vols polaires",
			19: "Les bulletins météorologiques SPECI et les tendances",
			20: "La météorologie spatiale et les éruptions solaires",
			21: "Les microclimats et les effets locaux",
			22: "La mesure de la visibilité et les RVR",
			23: "Les nuages à développement vertical et les cumulonimbus",
			24: "Les précipitations : formation et types",
			25: "La pression atmosphérique et les systèmes dépressionnaires",
		},
		"navigation": {
			6:  "La navigation inertielle (INS/IRS)",
			7:  "La navigation par satellite (GNSS) et les systèmes de correction",
			8:  "La navigation RNAV et les procédures PBN",
			9:  "La navigation d'urgence et les procédures de déroutement",
			10: "La navigation océanique et les routes MNPS",
			11: "Les systèmes de gestion de vol (FMS)",
			12: "La navigation verticale et les profils de descente",
			13: "Les procédures d'approche aux instruments (IAP)",
			14: "La navigation en espace aérien contrôlé",
			15: "Les systèmes d'atterrissage (ILS, MLS, GLS)",
			16: "La navigation à longue distance et les vols ETOPS",
			17: "Les cartes de navigation électroniques (EFB)",
			18: "La navigation en zone montagneuse",
			19: "Les points de report obligatoires et la gestion des waypoints",
			20: "La navigation de précision et les minima opérationnels",
			21: "Les systèmes de navigation hyperboliques (LORAN, Omega)",
			22: "La navigation astronomique et les méthodes traditionnelles",
			23: "Les routes aériennes et les espaces aériens organisés",
			24: "La navigation tactique et le vol à vue",
			25: "Les systèmes de navigation futurs et l'évolution technologique",
		},
		"performance": {
			6:  "Les performances en montée et les facteurs d'influence",
			7:  "Les performances en croisière et l'optimisation du rendement",
			8:  "Les performances à l'atterrissage et les distances d'arrêt",
			9:  "Les limitations de piste et les obstacles",
			10: "Les performances par temps chaud et en altitude",
			11: "Les graphiques de performance et leur utilisation",
			12: "Les performances avec panne moteur",
			13: "Les facteurs de charge et les limites structurelles",
			14: "Les performances des hélicoptères",
			15: "Les calculs de performance avancés",
		},
		"aircraft_general": {
			6:  "Les systèmes de commandes de vol électriques (Fly-by-Wire)",
			7:  "Les systèmes de protection et de sécurité",
			8:  "Les systèmes d'oxygène et de pressurisation",
			9:  "Les systèmes de dégivrage et d'antigivrage",
			10: "Les trains d'atterrissage et les systèmes de freinage",
			11: "Les systèmes de gestion du carburant avancés",
			12: "Les moteurs à turbine et les turboréacteurs",
			13: "Les systèmes électriques avancés et les bus",
			14: "Les systèmes hydrauliques et pneumatiques",
			15: "Les systèmes de climatisation et de contrôle environnemental",
			16: "Les matériaux composites et les nouvelles technologies",
			17: "Les systèmes de détection et d'évitement (TCAS)",
			18: "Les enregistreurs de vol (boîtes noires)",
			19: "Les systèmes de navigation intégrés",
			20: "La maintenance et la fiabilité des systèmes",
		},
		"flight_planning": {
			6:  "La planification des vols internationaux",
			7:  "Les procédures de dégagement et les aérodromes alternats",
			8:  "La gestion du carburant en vol et les décisions de déroutement",
			9:  "La planification des vols ETOPS",
			10: "Les restrictions de vol et les zones spéciales",
			11: "La planification des vols de nuit",
			12: "Les procédures de départ et d'arrivée standardisées (SID/STAR)",
			13: "La gestion des NOTAM et des bulletins d'information",
			14: "La planification des vols en montagne",
			15: "Les calculs de masse et centrage avancés",
			16: "La planification des vols long-courriers",
			17: "Les minimas opérationnels et les décisions de vol",
			18: "La gestion des équipages et les limitations de temps de vol",
			19: "Les procédures d'urgence et la planification de contingence",
			20: "L'utilisation des logiciels de planification de vol",
		},
		"human_performance": {
			6:  "La prise de décision aéronautique (ADM)",
			7:  "La gestion des ressources de l'équipage (CRM)",
			8:  "La communication non-verbale et les indicateurs de stress",
			9:  "Les biais cognitifs et les erreurs de jugement",
			10: "La gestion de la fatigue et les stratégies de sommeil",
			11: "Les facteurs nutritionnels et l'hydratation",
			12: "La gestion des menaces et des erreurs (TEM)",
			13: "Les différences culturelles et la communication multiculturelle",
			14: "L'automatisation et la conscience de la situation",
			15: "Les facteurs humains dans la maintenance aéronautique",
			16: "La résilience et la gestion du stress chronique",
			17: "Les facteurs humains dans les opérations de vol",
			18: "La formation aux facteurs humains et la simulation",
			19: "Les aspects juridiques des facteurs humains",
			20: "L'ergonomie et la conception des postes de pilotage",
		},
		"operational_procedures": {
			6:  "Les procédures d'urgence en vol",
			7:  "Les procédures de communication en situation d'urgence",
			8:  "Les procédures de détournement et de sécurité",
			9:  "Les procédures de vol aux instruments",
			10: "Les procédures de vol de nuit",
			11: "Les procédures de vol en formation",
			12: "Les procédures de vol à basse altitude",
			13: "Les procédures de vol en zone montagneuse",
			14: "Les procédures de vol par conditions givrantes",
			15: "Les procédures de vol par forte turbulence",
			16: "Les procédures de gestion du carburant",
			17: "Les procédures de gestion des passagers",
			18: "Les procédures de gestion des marchandises dangereuses",
			19: "Les procédures de gestion des opérations au sol",
			20: "Les procédures de gestion des opérations de maintenance",
			21: "Les procédures de gestion des opérations d'escale",
			22: "Les procédures de gestion des opérations de vol",
			23: "Les procédures de gestion des situations anormales",
			24: "Les procédures de gestion des situations d'urgence",
			25: "Les procédures de gestion des opérations spéciales",
		},
		"principles_of_flight": {
			6:  "Les effets de sol et l'effet de sol en vol",
			7:  "Les performances aérodynamiques en virage",
			8:  "Les caractéristiques de stabilité longitudinale",
			9:  "Les caractéristiques de stabilité latérale et directionnelle",
			10: "Les phénomènes de flutter et de vibration",
			11: "Les effets de la compressibilité et le vol à haute vitesse",
			12: "Les caractéristiques aérodynamiques des hélices",
			13: "Les interactions aérodynamiques entre voilure et empennage",
			14: "Les effets de la distribution de masse sur la stabilité",
			15: "Les principes aérodynamiques des décrochages avancés",
			16: "Les systèmes hypersustentateurs et leur fonctionnement",
			17: "Les phénomènes aérodynamiques transsoniques",
			18: "Les caractéristiques aérodynamiques des ailes en flèche",
			19: "Les principes de la mécanique du vol",
			20: "Les avancées en aérodynamique et les nouvelles configurations",
		},
		"communications": {
			6:  "Les communications en espace aérien contrôlé",
			7:  "Les communications d'urgence et de détresse",
			8:  "Les communications inter-pilotes et la coordination",
			9:  "Les communications avec les services d'information de vol (FIS)",
			10: "Les communications en environnement international",
			11: "Les communications par satellite et les liaisons de données",
			12: "Les communications en zone océanique",
			13: "Les communications en situation de panne radio",
			14: "Les communications avec les services de sauvetage",
			15: "Les communications en vol de nuit et par conditions difficiles",
		},
		"mass_and_balance": {
			6:  "Les calculs de centrage pour configurations spéciales",
			7:  "Les effets du centrage sur les performances",
			8:  "Les procédures de pesée des aéronefs",
			9:  "Les limitations de charge et les facteurs de sécurité",
			10: "Les calculs de masse et centrage pour les hélicoptères",
			11: "Les logiciels de calcul de masse et centrage",
			12: "Les procédures de chargement et d'arrimage",
			13: "Les effets du carburant sur le centrage en vol",
			14: "Les calculs de masse et centrage pour les vols cargo",
			15: "Les réglementations sur les masses et centrage",
		},
		"instrumentation": {
			6:  "Les systèmes gyroscopiques avancés",
			7:  "Les systèmes de référence de cap et d'attitude (AHRS)",
			8:  "Les systèmes d'affichage tête haute (HUD)",
			9:  "Les systèmes de gestion de vol (FMS) avancés",
			10: "Les systèmes d'alarme et de surveillance",
			11: "Les systèmes de détection de givrage",
			12: "Les systèmes de mesure de la température et de la pression",
			13: "Les systèmes de communication intégrés",
			14: "Les systèmes de navigation d'urgence",
			15: "Les systèmes d'enregistrement et de maintenance",
			16: "Les systèmes de visualisation synthétique (SVS)",
			17: "Les systèmes de vision améliorée (EVS)",
			18: "Les systèmes de détection de trafic et d'évitement",
			19: "Les systèmes de gestion des alarmes et des warnings",
			20: "Les évolutions technologiques des instruments de bord",
		},
	}

	if catTitles, ok := titles[category]; ok {
		if title, ok := catTitles[lessonNum]; ok {
			return title
		}
	}
	return fmt.Sprintf("Leçon avancée %d de %s", lessonNum, getCategoryNameFr(category))
}

func getLessonTitleEn(category string, lessonNum int) string {
	titles := map[string]map[int]string{
		"airlaw": {
			6:  "Bilateral Agreements and License Recognition",
			7:  "Air Traffic Management (ATM) and Air Traffic Services (ATS)",
			8:  "Aviation Safety Investigations and Incident Reporting",
			9:  "Environmental Regulations and Noise Pollution",
			10: "Passenger Rights and Carrier Liability",
			11: "Aircraft Certification and Continuing Airworthiness",
			12: "Approved Training Organizations (ATO) and Pilot Training",
			13: "Safety Management Systems (SMS) in Aviation Organizations",
			14: "Commercial Operations Rules (OPS) and Crew Management",
			15: "Drone Regulations and Remotely Piloted Aircraft Systems",
			16: "International Code-Sharing Agreements and Alliances",
			17: "Data Protection and Privacy in Aviation",
			18: "Technical Standards and Equipment Certification",
			19: "Crisis Management and Business Continuity",
			20: "Recent Regulatory Developments and Future Perspectives",
		},
		"meteorology": {
			6:  "Dangerous Weather Phenomena: Thunderstorms and Wind Shear",
			7:  "Weather Forecasting and Numerical Models",
			8:  "Weather Charts: Analysis and Interpretation",
			9:  "Tropical Meteorology and Cyclones",
			10: "Optical Phenomena in the Atmosphere",
			11: "Mountain Meteorology and Orographic Effects",
			12: "Frontal Systems and Their Evolution",
			13: "Clear Air Turbulence (CAT) and Mountain Waves",
			14: "Atmospheric Icing and Prevention",
			15: "Weather Radars and Their Use in Flight",
			16: "Climatology and Seasonal Variations",
			17: "Electrical Phenomena in the Atmosphere",
			18: "Arctic Meteorology and Polar Flights",
			19: "SPECI Weather Bulletins and Trends",
			20: "Space Weather and Solar Flares",
			21: "Microclimates and Local Effects",
			22: "Visibility Measurement and RVR",
			23: "Vertical Development Clouds and Cumulonimbus",
			24: "Precipitation: Formation and Types",
			25: "Atmospheric Pressure and Depression Systems",
		},
		"navigation": {
			6:  "Inertial Navigation Systems (INS/IRS)",
			7:  "Satellite Navigation (GNSS) and Correction Systems",
			8:  "RNAV Navigation and PBN Procedures",
			9:  "Emergency Navigation and Diversion Procedures",
			10: "Oceanic Navigation and MNPS Routes",
			11: "Flight Management Systems (FMS)",
			12: "Vertical Navigation and Descent Profiles",
			13: "Instrument Approach Procedures (IAP)",
			14: "Navigation in Controlled Airspace",
			15: "Landing Systems (ILS, MLS, GLS)",
			16: "Long-Range Navigation and ETOPS Flights",
			17: "Electronic Flight Bags (EFB)",
			18: "Navigation in Mountainous Areas",
			19: "Compulsory Reporting Points and Waypoint Management",
			20: "Precision Navigation and Operational Minima",
			21: "Hyperbolic Navigation Systems (LORAN, Omega)",
			22: "Celestial Navigation and Traditional Methods",
			23: "Airways and Organized Airspace",
			24: "Tactical Navigation and Visual Flight",
			25: "Future Navigation Systems and Technological Evolution",
		},
		"performance": {
			6:  "Climb Performance and Influencing Factors",
			7:  "Cruise Performance and Efficiency Optimization",
			8:  "Landing Performance and Stopping Distances",
			9:  "Runway Limitations and Obstacles",
			10: "Hot Weather and High Altitude Performance",
			11: "Performance Charts and Their Use",
			12: "Engine Failure Performance",
			13: "Load Factors and Structural Limits",
			14: "Helicopter Performance",
			15: "Advanced Performance Calculations",
		},
		"aircraft_general": {
			6:  "Fly-by-Wire Flight Control Systems",
			7:  "Protection and Safety Systems",
			8:  "Oxygen and Pressurization Systems",
			9:  "De-icing and Anti-icing Systems",
			10: "Landing Gear and Braking Systems",
			11: "Advanced Fuel Management Systems",
			12: "Turbine Engines and Turbojets",
			13: "Advanced Electrical Systems and Buses",
			14: "Hydraulic and Pneumatic Systems",
			15: "Air Conditioning and Environmental Control Systems",
			16: "Composite Materials and New Technologies",
			17: "Traffic Alert and Collision Avoidance Systems (TCAS)",
			18: "Flight Data Recorders (Black Boxes)",
			19: "Integrated Navigation Systems",
			20: "System Maintenance and Reliability",
		},
		"flight_planning": {
			6:  "International Flight Planning",
			7:  "Diversion Procedures and Alternate Aerodromes",
			8:  "In-Flight Fuel Management and Diversion Decisions",
			9:  "ETOPS Flight Planning",
			10: "Flight Restrictions and Special Areas",
			11: "Night Flight Planning",
			12: "Standard Instrument Departures and Arrivals (SID/STAR)",
			13: "NOTAM Management and Information Bulletins",
			14: "Mountain Flight Planning",
			15: "Advanced Mass and Balance Calculations",
			16: "Long-Haul Flight Planning",
			17: "Operational Minima and Flight Decisions",
			18: "Crew Management and Flight Time Limitations",
			19: "Emergency Procedures and Contingency Planning",
			20: "Flight Planning Software Usage",
		},
		"human_performance": {
			6:  "Aeronautical Decision Making (ADM)",
			7:  "Crew Resource Management (CRM)",
			8:  "Non-Verbal Communication and Stress Indicators",
			9:  "Cognitive Biases and Judgment Errors",
			10: "Fatigue Management and Sleep Strategies",
			11: "Nutritional Factors and Hydration",
			12: "Threat and Error Management (TEM)",
			13: "Cultural Differences and Multicultural Communication",
			14: "Automation and Situation Awareness",
			15: "Human Factors in Aircraft Maintenance",
			16: "Resilience and Chronic Stress Management",
			17: "Human Factors in Flight Operations",
			18: "Human Factors Training and Simulation",
			19: "Legal Aspects of Human Factors",
			20: "Ergonomics and Cockpit Design",
		},
		"operational_procedures": {
			6:  "In-Flight Emergency Procedures",
			7:  "Emergency Communication Procedures",
			8:  "Hijacking and Security Procedures",
			9:  "Instrument Flight Procedures",
			10: "Night Flight Procedures",
			11: "Formation Flight Procedures",
			12: "Low Altitude Flight Procedures",
			13: "Mountain Flight Procedures",
			14: "Icing Condition Flight Procedures",
			15: "Severe Turbulence Flight Procedures",
			16: "Fuel Management Procedures",
			17: "Passenger Management Procedures",
			18: "Dangerous Goods Management Procedures",
			19: "Ground Operations Management Procedures",
			20: "Maintenance Operations Management Procedures",
			21: "Turnaround Operations Management Procedures",
			22: "Flight Operations Management Procedures",
			23: "Abnormal Situation Management Procedures",
			24: "Emergency Situation Management Procedures",
			25: "Special Operations Management Procedures",
		},
		"principles_of_flight": {
			6:  "Ground Effect and Ground Effect in Flight",
			7:  "Aerodynamic Performance in Turns",
			8:  "Longitudinal Stability Characteristics",
			9:  "Lateral and Directional Stability Characteristics",
			10: "Flutter and Vibration Phenomena",
			11: "Compressibility Effects and High-Speed Flight",
			12: "Aerodynamic Characteristics of Propellers",
			13: "Aerodynamic Interactions Between Wing and Tail",
			14: "Effects of Mass Distribution on Stability",
			15: "Advanced Stall Aerodynamics",
			16: "High-Lift Systems and Their Operation",
			17: "Transonic Aerodynamic Phenomena",
			18: "Aerodynamic Characteristics of Swept Wings",
			19: "Principles of Flight Mechanics",
			20: "Advances in Aerodynamics and New Configurations",
		},
		"communications": {
			6:  "Communications in Controlled Airspace",
			7:  "Emergency and Distress Communications",
			8:  "Inter-Pilot Communications and Coordination",
			9:  "Communications with Flight Information Services (FIS)",
			10: "Communications in International Environment",
			11: "Satellite Communications and Data Links",
			12: "Oceanic Area Communications",
			13: "Radio Failure Communications",
			14: "Communications with Rescue Services",
			15: "Night Flight and Adverse Condition Communications",
		},
		"mass_and_balance": {
			6:  "Center of Gravity Calculations for Special Configurations",
			7:  "Effects of Center of Gravity on Performance",
			8:  "Aircraft Weighing Procedures",
			9:  "Load Limitations and Safety Factors",
			10: "Helicopter Mass and Balance Calculations",
			11: "Mass and Balance Software",
			12: "Loading and Securing Procedures",
			13: "Effects of Fuel on In-Flight Center of Gravity",
			14: "Cargo Flight Mass and Balance Calculations",
			15: "Mass and Balance Regulations",
		},
		"instrumentation": {
			6:  "Advanced Gyroscopic Systems",
			7:  "Attitude and Heading Reference Systems (AHRS)",
			8:  "Head-Up Display Systems (HUD)",
			9:  "Advanced Flight Management Systems (FMS)",
			10: "Warning and Monitoring Systems",
			11: "Ice Detection Systems",
			12: "Temperature and Pressure Measurement Systems",
			13: "Integrated Communication Systems",
			14: "Emergency Navigation Systems",
			15: "Recording and Maintenance Systems",
			16: "Synthetic Vision Systems (SVS)",
			17: "Enhanced Vision Systems (EVS)",
			18: "Traffic Detection and Avoidance Systems",
			19: "Warning and Alert Management Systems",
			20: "Technological Evolution of Cockpit Instruments",
		},
	}

	if catTitles, ok := titles[category]; ok {
		if title, ok := catTitles[lessonNum]; ok {
			return title
		}
	}
	return fmt.Sprintf("Advanced Lesson %d of %s", lessonNum, getCategoryNameEn(category))
}

func getKeyPointsFr(category string, lessonNum int) string {
	points := []string{
		"Les concepts fondamentaux sont expliqués avec des exemples concrets",
		"Les procédures standardisées sont détaillées étape par étape",
		"Les facteurs de sécurité sont mis en évidence",
		"Les références réglementaires applicables sont citées",
		"Les bonnes pratiques professionnelles sont présentées",
	}
	result := ""
	for i, p := range points {
		result += fmt.Sprintf("%d. **%s**\n", i+1, p)
	}
	return result
}

func getKeyPointsEn(category string, lessonNum int) string {
	points := []string{
		"Fundamental concepts are explained with concrete examples",
		"Standardized procedures are detailed step by step",
		"Safety factors are highlighted",
		"Applicable regulatory references are cited",
		"Professional best practices are presented",
	}
	result := ""
	for i, p := range points {
		result += fmt.Sprintf("%d. **%s**\n", i+1, p)
	}
	return result
}

func getExamplesFr(category string, lessonNum int) string {
	examples := []string{
		"Étude de cas : application des concepts dans un scénario réel",
		"Exercice pratique : mise en situation avec analyse des résultats",
		"Simulation : déroulement pas à pas d'une procédure type",
		"Analyse comparative : avantages et inconvénients des différentes approches",
	}
	result := ""
	for i, ex := range examples {
		result += fmt.Sprintf("- **Exemple %d**: %s\n", i+1, ex)
	}
	return result
}

func getExamplesEn(category string, lessonNum int) string {
	examples := []string{
		"Case study: application of concepts in a real scenario",
		"Practical exercise: situational analysis with results",
		"Simulation: step-by-step procedure walkthrough",
		"Comparative analysis: advantages and disadvantages of different approaches",
	}
	result := ""
	for i, ex := range examples {
		result += fmt.Sprintf("- **Example %d**: %s\n", i+1, ex)
	}
	return result
}

func getSummaryFr(category string, lessonNum int) string {
	summaries := []string{
		"Cette leçon vous a présenté les concepts avancés essentiels pour votre progression.",
		"La maîtrise de ces notions est indispensable pour réussir les examens théoriques.",
		"Les connaissances acquises dans cette leçon s'appliquent directement en situation de vol.",
		"N'hésitez pas à revoir les leçons précédentes si certains concepts ne sont pas clairs.",
	}
	return summaries[(lessonNum-1)%len(summaries)]
}

func getSummaryEn(category string, lessonNum int) string {
	summaries := []string{
		"This lesson has presented the essential advanced concepts for your progression.",
		"Mastery of these concepts is essential for passing theoretical exams.",
		"The knowledge gained in this lesson applies directly to flight situations.",
		"Feel free to review previous lessons if some concepts are unclear.",
	}
	return summaries[(lessonNum-1)%len(summaries)]
}
