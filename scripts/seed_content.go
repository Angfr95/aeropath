package main

// lessonContentFr returns the French content for a given category and lesson number.
func lessonContentFr(category string, lessonNum int) string {
	contents := map[string]map[int]string{
		"airlaw": {
			1: `## Introduction au droit aérien et à l'OACI

Le droit aérien est l'ensemble des règles qui régissent l'utilisation de l'espace aérien et les activités aéronautiques.

### L'Organisation de l'Aviation Civile Internationale (OACI)

L'OACI est une agence spécialisée des Nations Unies créée en 1944 par la Convention de Chicago. Son rÃ´le principal est de:
- Ã‰tablir les normes et pratiques recommandées (SARPs)
- Assurer la sécurité et l'efficacité du transport aérien
- Harmoniser les réglementations entre les Ã‰tats membres

### Les Annexes à la Convention de Chicago

Les 19 annexes couvrent tous les aspects de l'aviation civile:
- **Annexe 1**: Licences du personnel
- **Annexe 2**: Règles de l'air
- **Annexe 3**: Météorologie
- **Annexe 6**: Exploitation technique des aéronefs
- **Annexe 8**: Navigabilité des aéronefs`,
			2: `## L'EASA et la réglementation européenne

L'Agence Européenne de la Sécurité Aérienne (EASA) est l'autorité compétente pour la réglementation aéronautique en Europe.

### Règlements européens clés

- **Règlement (UE) 2018/1139**: Règlement de base EASA
- **Part-FCL**: Licences du personnel navigant
- **Part-M**: Maintien de la navigabilité
- **Part-66**: Certificat de maintenance
- **Part-147**: Organismes de formation agréés

### Structure réglementaire

La réglementation européenne est organisée en:
1. **Règlements de base**: Cadre général
2. **Règlements d'application**: Détails techniques
3. **AMC (Acceptable Means of Compliance)**: Moyens acceptables
4. **CS (Certification Specifications)**: Spécifications de certification`,
			3: `## Les licences de pilote et certificats médicaux

### Types de licences

- **LAPL (Light Aircraft Pilot Licence)**: Pour les aéronefs légers
- **PPL (Private Pilot Licence)**: Licence de pilote privé
- **CPL (Commercial Pilot Licence)**: Licence de pilote professionnel
- **ATPL (Airline Transport Pilot Licence)**: Licence de pilote de ligne
- **IR (Instrument Rating)**: Qualification de vol aux instruments

### Certificat médical

Le certificat médical est obligatoire pour exercer les privilèges d'une licence:
- **Classe 1**: Pour ATPL et CPL (examen annuel)
- **Classe 2**: Pour PPL et LAPL (examen tous les 2 à 5 ans)
- **LAPL**: Examen médical allégé

### Conditions médicales importantes

- Acuité visuelle minimale 6/9 pour chaque Å“il
- Pas de daltonisme pour les feux de navigation
- Audiométrie normale
- Pas de troubles cardiovasculaires graves
- Pas d'épilepsie`,
			4: `## Les espaces aériens classés A à G

### Classification de l'espace aérien

L'espace aérien est divisé en 7 classes, de A à G:

- **Classe A**: Vol IFR uniquement, contrÃ´le permanent
- **Classe B**: IFR et VFR, séparation assurée
- **Classe C**: IFR et VFR, séparation IFR/IFR et IFR/VFR
- **Classe D**: IFR et VFR, séparation IFR/IFR uniquement
- **Classe E**: IFR contrÃ´lé, VFR avec information de trafic
- **Classe F**: Espace consultatif
- **Classe G**: Espace non contrÃ´lé

### Règles de priorité

1. Aéronef en détresse
2. Ballons
3. Planeurs
4. Dirigeables
5. Avions et hélicoptères
6. Aéronefs remorquant

### Services du contrÃ´le aérien

- **ContrÃ´le régional (ACC)**: En route
- **ContrÃ´le d'approche (APP)**: Arrivée et départ
- **Tour de contrÃ´le (TWR)**: Circulation au sol et piste`,
			5: `## Les règles de l'air SERA

Les règles de l'air standardisées européennes (SERA) sont basées sur l'Annexe 2 de l'OACI.

### Règles générales

- **Protection des personnes et des biens**: L'exploitant doit prendre toutes les mesures raisonnables
- **Prévention des collisions**: Règles de priorité et de croisement
- **Feux de navigation**: Vert à droite, rouge à gauche, blanc à l'arrière
- **Signaux de détresse**: Mayday (urgence immédiate), Pan-Pan (urgence)

### Règles VFR

- Visibilité minimale: 5 km (ou 1500 m selon les cas)
- Distance des nuages: 1500 m horizontal, 300 m (1000 ft) vertical
- Plafond minimal: Pas de plafond minimum en G

### Règles IFR

- Plan de vol obligatoire
- Communication radio permanente
- Respect des routes et niveaux de vol`,
		},
		"meteorology": {
			1: `## L'atmosphère terrestre : composition et structure

### Composition de l'atmosphère

L'atmosphère terrestre est composée de:
- **Azote (Nâ‚‚)**: 78%
- **Oxygène (Oâ‚‚)**: 21%
- **Argon (Ar)**: 0.93%
- **Gaz traces**: COâ‚‚, néon, hélium, méthane, vapeur d'eau

### Structure verticale

- **Troposphère**: 0-11 km, contient 80% de la masse, température diminue avec l'altitude
- **Tropopause**: Limite entre troposphère et stratosphère
- **Stratosphère**: 11-50 km, température stable puis augmente
- **Mésosphère**: 50-85 km
- **Thermosphère**: Au-delà de 85 km

### Atmosphère type OACI

- Pression au niveau de la mer: 1013.25 hPa
- Température: 15Â°C (288.15 K)
- Gradient thermique: -6.5Â°C par km
- Densité: 1.225 kg/mÂ³`,
			2: `## La température et les échanges thermiques

### Sources de chaleur

- **Rayonnement solaire**: Source principale de chaleur
- **Rayonnement terrestre**: La Terre réémet l'énergie reÃ§ue
- **Effet de serre**: Les gaz à effet de serre retiennent la chaleur

### Transferts thermiques

- **Conduction**: Transfert par contact direct
- **Convection**: Transfert par mouvement vertical de l'air
- **Advection**: Transfert par mouvement horizontal
- **Rayonnement**: Transfert par ondes électromagnétiques

### Gradient thermique

Le gradient thermique vertical est la variation de température avec l'altitude:
- **Gradient adiabatique sec**: -1Â°C/100m (air sec)
- **Gradient adiabatique humide**: -0.6Â°C/100m (air saturé)
- **Gradient réel**: Variable selon les conditions

### Inversions de température

Une inversion se produit quand la température augmente avec l'altitude:
- Inversion de rayonnement (nuit claire)
- Inversion de subsidence
- Inversion frontale`,
			3: `## La pression atmosphérique

### Définition et mesure

La pression atmosphérique est le poids de la colonne d'air au-dessus d'un point donné. Elle se mesure en:
- **Hectopascals (hPa)**: Unité standard
- **Millimètres de mercure (mmHg)**: Unité historique
- **Pouces de mercure (inHg)**: Unité américaine

### Variation avec l'altitude

La pression diminue avec l'altitude selon une loi exponentielle:
- Au niveau de la mer: 1013.25 hPa
- Ã€ 5000 ft: environ 850 hPa
- Ã€ 10000 ft: environ 700 hPa
- Ã€ 18000 ft: environ 500 hPa

### Systèmes de pression

- **Anticyclone (Haute pression)**: Air descendant, temps stable
- **Dépression (Basse pression)**: Air ascendant, temps perturbé
- **Dorsale**: Extension d'un anticyclone
- **Creux**: Extension d'une dépression

### Calage altimétrique

- **QNH**: Pression au niveau de la mer (altitude)
- **QFE**: Pression au niveau de l'aérodrome (hauteur)
- **Standard (1013)**: Niveaux de vol (FL)`,
			4: `## Les nuages et les précipitations

### Classification des nuages

- **Cirrus (Ci)**: Nuages élevés, filamentaires
- **Cumulus (Cu)**: Nuages à développement vertical
- **Stratus (St)**: Nuages bas en couche
- **Nimbus (Nb)**: Nuages porteurs de précipitations

### Types de précipitations

- **Pluie**: Gouttes d'eau liquide
- **Neige**: Cristaux de glace
- **GrÃªle**: GrÃªlons (cumulonimbus)
- **Bruine**: Gouttes très fines
- **Verglas**: Pluie surfondue gélant au contact

### Nuages dangereux

- **Cumulonimbus (Cb)**: Orages, turbulence, grÃªle, cisaillement
- **Stratocumulus**: Possibilité de givrage
- **Nimbostratus**: Pluie continue, visibilité réduite`,
			5: `## Les vents et les masses d'air

### Origine du vent

Le vent est causé par les différences de pression atmosphérique:
- L'air se déplace des hautes pressions vers les basses pressions
- La force de Coriolis dévie le vent vers la droite (hémisphère nord)
- Le vent est d'autant plus fort que le gradient de pression est important

### Types de vents

- **Vent géostrophique**: Ã‰quilibre entre gradient et Coriolis
- **Vent de surface**: Ralenti par le frottement
- **Brises thermiques**: Mer/terre, montagne/vallée
- **Jet stream**: Courant d'altitude rapide

### Masses d'air

- **Arctique (A)**: Très froid et sec
- **Polaire (P)**: Froid et sec
- **Tropical (T)**: Chaud et humide
- **Ã‰quatorial (E)**: Très chaud et très humide
- **Maritime (m)**: Humide
- **Continental (c)**: Sec`,
		},
		"navigation": {
			1: `## Principes de base de la navigation aérienne

### Types de navigation

- **Navigation à vue (VFR)**: Utilisation de repères visuels
- **Navigation aux instruments (IFR)**: Utilisation des instruments de bord
- **Navigation radio**: Utilisation des aides radio
- **Navigation GNSS**: Utilisation des satellites

### Grandeurs fondamentales

- **Cap**: Direction dans laquelle pointe l'avion
- **Route**: Trajectoire prévue au sol
- **Dérive**: Angle entre cap et route dÃ» au vent
- **Vitesse propre**: Vitesse de l'avion dans l'air
- **Vitesse sol**: Vitesse de l'avion au sol

### Calculs de navigation

- **Temps = Distance / Vitesse**
- **Consommation = Temps Ã— Débit**
- **Correction de dérive**: Angle de correction pour compenser le vent`,
			2: `## La carte aéronautique

### Types de cartes

- **Carte OACI 1:500 000**: Navigation à vue
- **Carte en-route**: Navigation IFR
- **Carte d'approche**: Procédures d'atterrissage
- **Carte d'aérodrome**: Plan de l'aérodrome

### Informations sur la carte

- Relief et altitudes
- Espaces aériens et leurs limites
- Aides à la navigation
- Aérodromes et leurs caractéristiques
- Zones dangereuses, réglementées et interdites

### Lecture de carte

- **Latitude**: Lignes horizontales (parallèles)
- **Longitude**: Lignes verticales (méridiens)
- **Ã‰chelle**: Rapport distance carte/réalité
- **Déclinaison magnétique**: Différence entre nord vrai et nord magnétique`,
			3: `## Le compas magnétique

### Principe de fonctionnement

Le compas magnétique utilise le champ magnétique terrestre pour indiquer le nord magnétique.

### Erreurs du compas

- **Déviation**: Erreur due aux masses métalliques de l'avion
- **Variation**: Différence entre nord vrai et nord magnétique
- **Erreur de virage**: En virage, le compas indique incorrectement
- **Erreur d'accélération**: En accélération/décélération

### Utilisation en vol

- Vérifier le compas avant chaque vol
- Compenser la déviation avec le compensateur
- Utiliser le cap magnétique pour la navigation`,
			4: `## Les aides radio à la navigation

### VOR (VHF Omnidirectional Range)

- Ã‰metteur au sol émettant sur 360 radiales
- Permet de connaÃ®tre le cap vers ou depuis la station
- Portée: ligne de vue (environ 200 NM à haute altitude)

### NDB (Non-Directional Beacon)

- Ã‰metteur au sol émettant dans toutes les directions
- L'ADF (Automatic Direction Finder) indique la direction
- Portée: jusqu'à 400 NM
- Sujet aux interférences atmosphériques

### DME (Distance Measuring Equipment)

- Mesure la distance oblique entre l'avion et la station
- Fonctionne en UHF
- Précision: Â±0.5 NM ou 3%

### ILS (Instrument Landing System)

- Localizer: Guidage latéral
- Glide Slope: Guidage vertical
- Marker Beacons: Points de référence à l'approche`,
			5: `## Le GPS et les systèmes de navigation modernes

### Principe du GPS

Le GPS (Global Positioning System) utilise une constellation de 24 satellites:
- Chaque satellite émet sa position et l'heure
- Le récepteur calcule sa position par trilatération
- Précision: 5-10 mètres

### Autres systèmes GNSS

- **GLONASS**: Système russe
- **Galileo**: Système européen
- **BeiDou**: Système chinois

### Limitations du GPS

- Perte de signal (relief, bÃ¢timents)
- Interférences volontaires ou involontaires
- Précision réduite en haute latitude
- Nécessite une alimentation électrique`,
		},
		"aircraft_general": {
			1: `## Structure de l'aéronef

### Composants principaux

- **Fuselage**: Structure principale contenant la cabine
- **Voilure (ailes)**: Génération de la portance
- **Empennage**: Stabilité et contrÃ´le (dérive + stabilisateur)
- **Train d'atterrissage**: Support au sol
- **Groupe motopropulseur**: Moteur et hélice

### Types de structures

- **Structure en treillis**: Tubes métalliques assemblés
- **Structure monocoque**: Peau travaillante
- **Structure semi-monocoque**: Longerons + lisses + peau

### Matériaux

- **Aluminium**: Léger et résistant
- **Acier**: Résistance élevée
- **Titane**: Haute température
- **Composites**: Fibre de carbone, kevlar`,
			2: `## Les systèmes de l'aéronef

### Système électrique

- Batterie (12V ou 24V)
- Alternateur ou génératrice
- Bus électrique principal et secondaire
- Disjoncteurs et fusibles

### Système hydraulique

- Fluide hydraulique (Skydrol)
- Pompe hydraulique
- Vérins et actionneurs
- Accumulateur

### Système de carburant

- Réservoirs (voilure, fuselage)
- Pompe de gavage
- Filtres et purgeurs
- Jauge de carburant

### Système de pressurisation

- Compresseur ou prélèvement moteur
- Vanne de régulation
- Indicateur de pression cabine`,
			3: `## Le moteur à piston

### Principe de fonctionnement

Le moteur à piston transforme l'énergie chimique du carburant en énergie mécanique:
1. Admission: Mélange air-carburant
2. Compression: Le piston comprime le mélange
3. Combustion: L'étincelle enflamme le mélange
4. Ã‰chappement: Les gaz brÃ»lés sont évacués

### Composants principaux

- Pistons et cylindres
- Vilebrequin
- Soupapes d'admission et d'échappement
- Magnétos (allumage)
- Carburateur ou injection

### Paramètres moteur

- **Régime (RPM)**: Tours par minute
- **Pression d'admission (MAP)**: Manifold Absolute Pressure
- **Température d'huile**: 80-100Â°C
- **Température des tÃªtes**: 200-250Â°C`,
			4: `## L'hélice

### Types d'hélices

- **Hélice à pas fixe**: Simple, légère
- **Hélice à pas variable**: Optimisation du rendement
- **Hélice à vitesse constante**: Régulation automatique

### Principes aérodynamiques

- **Pas géométrique**: Angle de la pale
- **Pas effectif**: Distance parcourue par tour
- **Angle d'attaque**: Angle entre la pale et le vent relatif
- **Rendement**: Rapport puissance fournie/puissance moteur

### Fonctionnement

- L'hélice crée une force de traction
- Le pas détermine la charge sur le moteur
- Le régime est contrÃ´lé par la manette de pas`,
			5: `## Les instruments de bord

### Instruments de vol (6 de base)

1. **Anémomètre (ASI)**: Vitesse indiquée
2. **Altimètre**: Altitude pression
3. **Variomètre (VSI)**: Vitesse verticale
4. **Horizon artificiel**: Attitude de l'avion
5. **Compas gyroscopique (DG)**: Cap
6. **Coordinateur de virage**: Taux de virage

### Instruments moteur

- Tachymètre (RPM)
- Manomètre d'huile
- Thermomètre d'huile
- Jauge de carburant
- Manomètre d'admission`,
		},
		"human_performance": {
			1: `## Les facteurs humains en aviation

### Introduction

Les facteurs humains étudient l'interaction entre les humains et leur environnement de travail. En aviation, ils sont essentiels pour la sécurité.

### Le modèle SHELL

- **S (Software)**: Procédures, checklists
- **H (Hardware)**: Ã‰quipements, instruments
- **E (Environment)**: Environnement de travail
- **L (Liveware)**: Les personnes
- **L (Liveware)**: Interactions entre personnes

### Erreur humaine

- **Erreur de performance**: Inattention, fatigue
- **Erreur de décision**: Mauvais jugement
- **Erreur de perception**: Illusions sensorielles
- **Violation**: Non-respect délibéré des règles`,
			2: `## La vision et les illusions visuelles

### Anatomie de l'Å“il

- **Rétine**: Cellules photoréceptrices (cÃ´nes et bÃ¢tonnets)
- **Fovéa**: Vision centrale (cÃ´nes)
- **Vision périphérique**: BÃ¢tonnets (sensibles à la lumière)

### Illusions visuelles en vol

- **Illusion de pente**: Terrain en pente donne l'impression d'Ãªtre incliné
- **Illusion de largeur de piste**: Piste étroite = plus haut, large = plus bas
- **Illusion d'approche**: Eau ou brouillard donne l'impression d'Ãªtre plus haut
- **Autokinésie**: Point lumineux fixe semble bouger

### Adaptation à l'obscurité

- 30 minutes pour une adaptation complète
- Utiliser une lumière rouge pour préserver la vision nocturne
- Ã‰viter les lumières vives avant le vol de nuit`,
			3: `## L'oreille et l'équilibre

### Anatomie de l'oreille

- **Oreille externe**: Pavillon et conduit auditif
- **Oreille moyenne**: Tympan et osselets
- **Oreille interne**: Cochlée (audition) et vestibule (équilibre)

### Le système vestibulaire

- **Canaux semi-circulaires**: Détection des rotations
- **Otolithes**: Détection des accélérations linéaires
- **Illusions vestibulaires**: Fausses sensations de mouvement

### Barotraumatisme

- Douleur due à la différence de pression
- Se produit surtout à la descente
- ManÅ“uvre de Valsalva pour équilibrer
- Ne pas voler avec un rhume ou une sinusite`,
			4: `## L'hypoxie et la respiration

### L'hypoxie

Manque d'oxygène dans les tissus:
- **Hypoxie hypobare**: Altitude (pression partielle réduite)
- **Hypoxie anémique**: Manque de globules rouges
- **Hypoxie stagnante**: Circulation réduite
- **Hypoxie histotoxique**: Cellules incapables d'utiliser l'Oâ‚‚

### SymptÃ´mes de l'hypoxie

- **Ã€ 10000 ft**: Diminution de la vision nocturne
- **Ã€ 15000 ft**: Euphorie, maux de tÃªte, fatigue
- **Ã€ 20000 ft**: Perte de conscience possible
- **Ã€ 25000 ft**: Perte de conscience en 3-5 minutes

### Temps d'utilité consciente

- 15000 ft: 30 minutes
- 20000 ft: 5-10 minutes
- 25000 ft: 3-5 minutes
- 30000 ft: 1-2 minutes`,
			5: `## La fatigue et le stress

### Types de fatigue

- **Fatigue aiguÃ«**: Après une longue période d'éveil
- **Fatigue chronique**: Accumulation sur plusieurs jours
- **Fatigue physique**: Due à l'effort
- **Fatigue mentale**: Due à la concentration

### Gestion de la fatigue

- Sommeil régulier (7-9 heures)
- Ã‰viter la caféine avant le coucher
- Faire des pauses régulières
- Bien s'hydrater

### Le stress

- **Eustress**: Stress positif qui améliore la performance
- **Distress**: Stress négatif qui diminue la performance
- **Courbe de Yerkes-Dodson**: Performance optimale à stress modéré

### Gestion du stress

- Respiration profonde
- Priorisation des tÃ¢ches
- Communication claire
- Connaissance de ses limites`,
		},
		"performance": {
			1: `## Introduction aux performances avion

### Définitions

Les performances d'un avion décrivent ses capacités dans différentes phases de vol:
- **Décollage**: Distance et pente
- **Montée**: Taux et pente de montée
- **Croisière**: Vitesse et consommation
- **Atterrissage**: Distance d'atterrissage

### Facteurs influenÃ§ant les performances

- **Masse**: Plus l'avion est lourd, plus les performances sont réduites
- **Altitude**: L'air moins dense réduit la portance et la puissance moteur
- **Température**: L'air chaud est moins dense
- **Vent**: Vent de face réduit les distances, vent arrière les augmente
- **Ã‰tat de la piste**: Piste mouillée ou en herbe augmente les distances`,
			2: `## Les masses et limitations structurales

### Masses maximales certifiées

- **MTOW (Maximum TakeOff Weight)**: Masse maximale au décollage
- **MLW (Maximum Landing Weight)**: Masse maximale à l'atterrissage
- **MZFW (Maximum Zero Fuel Weight)**: Masse maximale sans carburant utilisable
- **MRW (Maximum Ramp Weight)**: Masse maximale au parking

### Limitations structurales

- **Facteur de charge**: Limites positives et négatives (ex: +3.8g / -1.52g)
- **Vitesse maximale (Vne)**: Ne jamais dépasser
- **Vitesse de manÅ“uvre (Va)**: Ne pas braquer les gouvernes brusquement au-delà

### Importance du respect des limitations

Le dépassement des masses maximales peut entraÃ®ner:
- Allongement des distances de décollage et d'atterrissage
- Réduction des performances en montée
- Risque de dommages structuraux
- Non-conformité réglementaire`,
			3: `## La densité de l'air et performances

### Qu'est-ce que la densité-altitude ?

La densité-altitude est l'altitude à laquelle correspond la densité réelle de l'air dans l'atmosphère standard. Elle combine:
- **L'altitude pression**: Altitude basée sur la pression
- **La température**: Correction de la température réelle

### Effets de la densité-altitude élevée

- **Diminution de la portance**: L'aile génère moins de portance
- **Diminution de la puissance moteur**: Moins d'oxygène pour la combustion
- **Augmentation des distances de décollage**: Jusqu'à 50% ou plus
- **Réduction du taux de montée**: Moins de puissance disponible

### Calcul de la densité-altitude

- Altitude pression + 120 ft par degré au-dessus de la température standard
- Utiliser les graphiques du manuel de vol`,
			4: `## Les distances de décollage

### Composantes de la distance de décollage

- **Distance de roulage**: Du point de départ à la rotation
- **Distance de décollage**: Jusqu'à 50 ft (15 m) de hauteur
- **Distance d'accélération-arrÃªt**: Distance pour accélérer puis s'arrÃªter

### Facteurs augmentant la distance de décollage

- **Masse élevée**: Plus d'inertie à vaincre
- **Altitude élevée**: Air moins dense
- **Température élevée**: Air moins dense
- **Vent arrière**: Augmente la vitesse sol nécessaire
- **Piste en montée**: Composante de poids défavorable
- **Piste mouillée/herbe**: Frottement réduit

### Utilisation des graphiques

Le manuel de vol fournit des graphiques pour calculer:
- La distance de décollage en fonction de la masse, altitude, température et vent`,
			5: `## Les vitesses V1, Vr et V2

### V1 - Vitesse de décision

- Vitesse maximale à laquelle le décollage peut Ãªtre interrompu
- Au-delà de V1, le décollage doit Ãªtre poursuivi
- Dépend de la masse, de la longueur de piste, de l'état de la piste

### Vr - Vitesse de rotation

- Vitesse à laquelle le pilote tire sur le manche pour lever le nez
- Généralement calculée pour assurer V2 à 35 ou 50 ft
- Dépend de la masse et de la configuration

### V2 - Vitesse de sécurité au décollage

- Vitesse minimale après panne moteur au décollage
- Assure une pente de montée suffisante
- Généralement 1.2 x Vs (vitesse de décrochage)`,
		},
		"flight_planning": {
			1: `## Introduction au planning de vol

### Objectifs du planning de vol

Le planning de vol permet de préparer un vol en toute sécurité en déterminant:
- La route à suivre et les altitudes
- Le carburant nécessaire
- Les aérodromes de dégagement
- Les NOTAM et restrictions

### Ã‰tapes du planning

1. **Consultation météo**: Prévisions et conditions réelles
2. **NOTAM**: Restrictions et informations temporaires
3. **Choix de la route**: Voies aériennes, points de report
4. **Calcul du carburant**: Route, réserve, dégagement
5. **Calcul des masses et centrage**: Vérification des limitations
6. **Plan de vol**: DépÃ´t du formulaire`,
			2: `## Les NOTAM et leur interprétation

### Qu'est-ce qu'un NOTAM ?

NOTAM (Notice to Air Missions) est un avis contenant des informations essentielles pour la sécurité des vols:
- Fermeture de pistes ou d'aérodromes
- Balises ou aides radio hors service
- Exercices militaires ou zones dangereuses
- Changements temporaires importants

### Types de NOTAM

- **NOTAMN**: Nouveau NOTAM
- **NOTAMR**: NOTAM remplaÃ§ant un précédent
- **NOTAMC**: NOTAM annulant un précédent

### Consultation des NOTAM

- Avant chaque vol, consulter les NOTAM pour:
  - L'aérodrome de départ
  - La route empruntée
  - L'aérodrome de destination
  - Les aérodromes de dégagement`,
			3: `## Le choix de la route et altitudes

### Sélection de la route

- **Voies aériennes**: Routes IFR standardisées
- **Navigation à vue**: Suivi de repères visuels
- **Points de report**: Balises, intersections, waypoints

### Altitudes de croisière

- **Règle semi-circulaire VFR**:
  - Route magnétique 000Â°-179Â°: altitude impaire + 500 ft
  - Route magnétique 180Â°-359Â°: altitude paire + 500 ft
- **Niveaux de vol IFR**: Basés sur le calage standard 1013 hPa

### Contraintes d'espace aérien

- Ã‰viter les zones interdites (P), réglementées (R), dangereuses (D)
- Respecter les altitudes minimales de survol`,
			4: `## Le calcul du carburant nécessaire

### Composantes du carburant

- **Carburant de route**: Du départ à la destination
- **Carburant de dégagement**: Jusqu'à l'aérodrome de dégagement
- **Réserve**: 30 min VFR jour, 45 min VFR nuit
- **Carburant taxi**: Consommation au sol

### Calcul pratique

1. Déterminer la distance totale
2. Calculer le temps de vol estimé
3. Multiplier par la consommation horaire
4. Ajouter les réserves réglementaires

### Exemple

Pour un vol de 2h avec consommation 25 L/h:
- Route: 50 L
- Réserve 30 min: 12.5 L
- Dégagement: 15 L
- Taxi: 3 L
- **Total: 80.5 L**`,
			5: `## Les réserves de carburant réglementaires

### Réserves VFR

- **VFR de jour**: 30 minutes de carburant au-dessus de la destination
- **VFR de nuit**: 45 minutes de carburant au-dessus de la destination

### Réserves IFR

- Carburant pour rejoindre l'aérodrome de dégagement
- Plus 30 minutes de réserve au-dessus du dégagement

### Gestion du carburant en vol

- Surveiller régulièrement la consommation
- Comparer le carburant restant au temps restant
- En cas de doute, atterrir pour se ravitailler
- Ne jamais hésiter à déclarer une urgence carburant`,
		},
		"operational_procedures": {
			1: `## Introduction aux procédures opérationnelles

### Définition

Les procédures opérationnelles sont l'ensemble des actions standardisées à effectuer pour chaque phase de vol.

### Importance des procédures

- **Standardisation**: Tous les pilotes effectuent les mÃªmes actions
- **Sécurité**: Réduction des erreurs et oublis
- **Efficacité**: Optimisation du déroulement du vol
- **Communication**: Langage commun entre pilotes et contrÃ´leurs

### Les check-lists

- Outil essentiel pour ne rien oublier
- Ã€ utiliser impérativement à chaque phase clé
- Lire et exécuter, pas réciter de mémoire`,
			2: `## La préparation du vol

### Briefing météo

- Consulter METAR, TAF, SIGMET
- Analyser les cartes météo
- Ã‰valuer les conditions sur la route

### Inspection de l'aéronef

- **Extérieure**: Visite pré-vol complète
- **Intérieure**: Vérification des documents et instruments
- **Niveaux**: Huile, carburant, liquides

### Planification

- Calcul du carburant
- Calcul des masses et centrage
- Plan de vol
- NOTAM`,
			3: `## Les inspections pré-vol

### Inspection extérieure

- **Hélice**: Ã‰tat des pales, fixation
- **Moteur**: Niveaux d'huile, cÃ¢bles, durites
- **Train**: Pneus, amortisseurs, freins
- **Ailes**: Ã‰tat général, volets, ailerons
- **Empennage**: Dérive, profondeur, gouverne
- **Feux**: Navigation, atterrissage, anticollision
- **Réservoirs**: Niveau carburant, bouchon

### Inspection intérieure

- **Documents**: Licence, certificats, manuel de vol
- **Instruments**: Test et calibration
- **Commandes**: Liberté de mouvement
- **Sièges et ceintures**: Réglage et verrouillage`,
			4: `## La mise en route et démarrage

### Procédure de démarrage

1. **Frein de parking**: SERRÉ
2. **Batterie**: ON
3. **Instruments**: Vérification
4. **Mixture**: RICHE
5. **Starter (si froid)**: TIRÉ
6. **Contact**: ON (les deux magnetos)
7. **Démarreur**: ACTION
8. **Après démarrage**: Starter poussé, régime ralenti

### Vérifications après démarrage

- Pression d'huile: doit monter dans les 30 secondes
- Alternateur: charge
- Instruments: fonctionnement normal
- Radio: test`,
			5: `## Le roulage et les consignes

### Procédure de roulage

- **Freins**: Testés avant de commencer à rouler
- **Palonnier**: Direction avec le palonnier (roulette de nez)
- **Freins différentiels**: Pour les virages serrés
- **Vitesse**: Pas plus rapide qu'une marche rapide

### Communication au roulage

- Demander l'autorisation de roulage au contrôle
- Suivre les voies de circulation (taxiways)
- S'arrêter avant de pénétrer sur une piste
- Effectuer les essais moteur avant le décollage`,
		},
		"principles_of_flight": {
			1: `## Introduction à l'aérodynamique

### Définition

L'aérodynamique est l'étude des forces agissant sur un corps en mouvement dans l'air.

### Les quatre forces du vol

- **Portance (Lift)**: Force vers le haut générée par les ailes
- **Poids (Weight)**: Force vers le bas due à la gravité
- **Traction (Thrust)**: Force vers l'avant générée par le moteur
- **Traînée (Drag)**: Force vers l'arrière due à la résistance de l'air

### Équilibre des forces

En vol stabilisé:
- Portance = Poids
- Traction = Traînée`,
			2: `## Le profil d'aile et ses caractéristiques

### Définition du profil

Le profil d'aile est la forme de la coupe transversale de l'aile.

### Caractéristiques du profil

- **Extrados**: Surface supérieure (cambrée)
- **Intrados**: Surface inférieure (plus plate)
- **Bord d'attaque**: Avant du profil
- **Bord de fuite**: Arrière du profil
- **Corde**: Ligne droite du bord d'attaque au bord de fuite

### Types de profils

- **Profil symétrique**: Extrados et intrados identiques
- **Profil cambré**: Extrados plus courbé que l'intrados
- **Profil laminaire**: Résistance réduite`,
			3: `## La portance : génération et facteurs

### Principe de Bernoulli

- L'air qui circule plus vite a une pression plus faible
- L'extrados (courbé) accélère l'air → pression plus basse
- L'intrados (plat) a une pression plus élevée
- La différence de pression crée la portance

### Formule de la portance

**L = 1/2 × ρ × V² × S × Cl**

- ρ (rho): Densité de l'air
- V: Vitesse de l'avion
- S: Surface alaire
- Cl: Coefficient de portance (dépend de l'angle d'attaque)

### Facteurs augmentant la portance

- Augmentation de la vitesse
- Augmentation de l'angle d'attaque (jusqu'au décrochage)
- Volets sortis (augmente le coefficient de portance)`,
			4: `## La traînée : parasite et induite

### Traînée parasite

- **Traînée de frottement**: Due à la viscosité de l'air sur les surfaces
- **Traînée de forme**: Due à la forme de l'objet
- **Traînée d'interférence**: Due aux intersections entre surfaces

### Traînée induite

- Générée par la portance elle-même
- Plus importante à basse vitesse (décollage, atterrissage)
- Diminue avec la vitesse
- Liée à l'angle d'attaque

### Traînée totale

- Courbe en U: minimale à une vitesse spécifique
- Vitesse de meilleure finesse: traînée minimale
- Important pour la planification de la descente`,
			5: `## Le décrochage et sa récupération

### Qu'est-ce que le décrochage ?

- Perte de portance due à un angle d'attaque trop élevé
- L'écoulement de l'air se décolle de l'extrados
- Survient à l'angle d'attaque critique (généralement 15-20°)

### Signes annonciateurs

- Vibrations de l'avion (buffeting)
- Diminution de l'efficacité des gouvernes
- Alarme de décrochage (avertisseur sonore)
- Nez qui baisse

### Récupération

1. **Pousser le manche** (réduire l'angle d'attaque)
2. **Plein gaz** (augmenter la vitesse)
3. **Ailes à niveau**
4. **Reprendre l'assiette de vol normale**`,
		},
		"communications": {
			1: `## Introduction aux communications aéronautiques

### Importance des communications

Les communications radio sont essentielles pour:
- La sécurité des vols
- La coordination avec le contrôle aérien
- L'information des autres trafics
- Les procédures d'urgence

### Principes de base

- **Clarté**: Parler distinctement
- **Concision**: Messages courts et précis
- **Standardisation**: Utiliser la phraséologie officielle
- **Discipline**: Écouter avant de transmettre`,
			2: `## La phraséologie standard

### Structure d'un message

1. **Qui vous appelez**: Indicatif de la station
2. **Qui vous êtes**: Votre indicatif
3. **Où vous êtes**: Position
4. **Ce que vous voulez**: Intention ou demande

### Exemples

- "Paris Info, bonjour, F-GABC, Cessna 172, de Toussus-le-Noble vers Chartres, 2000 ft, information Tango"
- "F-GABC, Paris Info, bonjour, transmettez"

### Mots-clés standard

- **Roger**: Message reçu
- **Wilco**: Will comply (obéira)
- **Affirm**: Oui
- **Negative**: Non
- **Say again**: Répétez
- **Stand by**: Attendez`,
			3: `## L'alphabet international OACI

### Alphabet phonétique

- A: Alfa, B: Bravo, C: Charlie, D: Delta
- E: Echo, F: Foxtrot, G: Golf, H: Hotel
- I: India, J: Juliett, K: Kilo, L: Lima
- M: Mike, N: November, O: Oscar, P: Papa
- Q: Quebec, R: Romeo, S: Sierra, T: Tango
- U: Uniform, V: Victor, W: Whiskey, X: X-ray
- Y: Yankee, Z: Zulu

### Chiffres

- 0: Zero, 1: One, 2: Two, 3: Three
- 4: Four, 5: Five, 6: Six, 7: Seven
- 8: Eight, 9: Niner

### Utilisation

- Indicatifs d'aéronef: F-GABC = Foxtrot Golf Alfa Bravo Charlie
- Niveaux de vol: FL 120 = Flight Level One Two Zero`,
			4: `## Les fréquences aéronautiques

### Bande VHF (118-137 MHz)

- **118.000 - 121.400**: Contrôle aérien
- **121.500**: Fréquence de détresse
- **122.000 - 123.050**: Services aéronautiques
- **123.500**: Auto-information VFR
- **126.700**: Information de vol (FIS)

### Bande HF

- Utilisée pour les vols long-courriers et océaniques
- Portée mondiale

### Sélection des fréquences

- Consulter les cartes aéronautiques
- Utiliser la fréquence indiquée pour chaque secteur
- Noter les fréquences de dégagement`,
			5: `## Les procédures VFR en zone non contrôlée

### Auto-information

En espace non contrôlé (classe G), les pilotes s'informent mutuellement:
- **Fréquence**: 123.500 MHz
- **Annonces**: Position, altitude, intentions

### Annonces obligatoires

- **Départ**: "Trafic, F-GABC décolle piste 24, aérodrome de X"
- **Position**: "Trafic, F-GABC, 5 NM au sud de X, 2000 ft"
- **Approche**: "Trafic, F-GABC, en finale piste 24, X"
- **Atterrissage**: "Trafic, F-GABC, atterri piste 24, X"

### Vigilance

- Rester à l'écoute en permanence
- Annoncer clairement ses intentions
- Surveiller visuellement les autres trafics`,
		},
		"mass_and_balance": {
			1: `## Introduction aux masses et centrage

### Importance du calcul

Le calcul des masses et du centrage est essentiel pour:
- La sécurité du vol
- Les performances de l'avion
- La stabilité et le contrôle
- La conformité réglementaire

### Principes de base

- **Masse**: Quantité de matière (kg)
- **Centrage**: Position du centre de gravité (CG)
- **Bras de levier**: Distance entre la charge et le point de référence
- **Moment**: Masse × Bras de levier`,
			2: `## Les définitions des masses

### Masses de base

- **Masse à vide (Empty Weight)**: Masse de l'avion sans charge
- **Masse à vide de base**: Inclut l'huile et les liquides inutilisables
- **Charge utile (Payload)**: Passagers + bagages + fret

### Masses maximales

- **MTOW**: Maximum TakeOff Weight
- **MLW**: Maximum Landing Weight
- **MZFW**: Maximum Zero Fuel Weight
- **MRW**: Maximum Ramp Weight

### Terminologie

- **Carburant utilisable**: Carburant pouvant être consommé en vol
- **Carburant inutilisable**: Reste dans les réservoirs
- **Charge marchande**: Charge utile + carburant`,
			3: `## Le centrage et le bras de levier

### Point de référence (Datum)

- Point fixe choisi par le constructeur
- Toutes les distances sont mesurées depuis ce point
- Généralement situé au niveau du bord d'attaque de l'aile ou du firewall

### Bras de levier

- Distance entre la charge et le datum
- **Bras positif**: Charge située derrière le datum
- **Bras négatif**: Charge située devant le datum

### Calcul du moment

**Moment = Masse × Bras**

- Exemple: 100 kg à 2.5 m du datum = 250 kg.m`,
			4: `## La détermination du centrage

### Calcul du centrage total

1. Additionner tous les moments
2. Additionner toutes les masses
3. Diviser le moment total par la masse totale

**CG = Σ Moments / Σ Masses**

### Exemple de calcul

| Élément | Masse (kg) | Bras (m) | Moment (kg.m) |
|---------|-----------|---------|--------------|
| Avion vide | 600 | 0.8 | 480 |
| Pilote | 80 | 0.4 | 32 |
| Passager | 70 | 1.2 | 84 |
| Carburant | 100 | 1.0 | 100 |
| **Total** | **850** | | **696** |

CG = 696 / 850 = 0.819 m`,
			5: `## Les enveloppes de centrage

### Définition

L'enveloppe de centrage est le domaine à l'intérieur duquel le CG doit se trouver pour que le vol soit sûr.

### Limites

- **Limite avant**: Centrage trop avant → stabilité excessive, manœuvrabilité réduite
- **Limite arrière**: Centrage trop arrière → instabilité, risque de décrochage

### Vérification

- Utiliser le diagramme de centrage du manuel de vol
- Placer le point (masse, CG) sur le graphique
- Vérifier qu'il se trouve à l'intérieur de l'enveloppe
- Si en dehors, redistribuer la charge`,
		},
		"instrumentation": {
			1: `## Introduction aux instruments de bord

### Classification des instruments

- **Instruments de vol**: Pilote-statique, gyroscopiques, compas
- **Instruments moteur**: Régime, pressions, températures
- **Instruments de navigation**: VOR, ADF, GPS, ILS
- **Instruments de gestion**: Carburant, temps, systèmes

### Les six instruments de base (6-pack)

1. Anémomètre (ASI)
2. Altimètre
3. Variomètre (VSI)
4. Horizon artificiel
5. Coordinateur de virage
6. Compas gyroscopique (DG)`,
			2: `## Le système pitot-statique

### Composants

- **Tube Pitot**: Mesure la pression dynamique (face au vent)
- **Prise statique**: Mesure la pression statique (air ambiant)
- **Instruments connectés**: ASI, altimètre, VSI

### Principe

- **Pression totale** = Pression statique + Pression dynamique
- **Pression dynamique** = Pression totale - Pression statique
- La pression dynamique est proportionnelle au carré de la vitesse

### Pannes possibles

- **Pitot bouché**: L'ASI se comporte comme un altimètre
- **Statique bouchée**: L'altimètre et le VSI sont bloqués
- **Givrage**: Chauffage pitot obligatoire en conditions givrantes`,
			3: `## L'anémomètre (ASI)

### Fonctionnement

L'anémomètre mesure la différence entre la pression totale (pitot) et la pression statique pour indiquer la vitesse.

### Types de vitesses

- **IAS (Indicated Airspeed)**: Vitesse lue directement sur l'instrument
- **CAS (Calibrated Airspeed)**: IAS corrigée des erreurs instrumentales
- **TAS (True Airspeed)**: CAS corrigée de la densité de l'air
- **GS (Ground Speed)**: TAS corrigée du vent

### Codes couleur

- **Arc blanc**: Plage d'utilisation des volets
- **Arc vert**: Plage d'utilisation normale
- **Arc jaune**: Plage d'utilisation avec précautions
- **Ligne rouge**: Vne (Ne jamais dépasser)`,
			4: `## L'altimètre

### Principe

L'altimètre mesure la pression atmosphérique et l'affiche en altitude.

### Calage altimétrique

- **QNH**: Calage pour lire l'altitude au-dessus du niveau de la mer
- **QFE**: Calage pour lire la hauteur au-dessus de l'aérodrome
- **Standard 1013**: Calage pour les niveaux de vol (FL)

### Erreurs altimétriques

- **Erreur de température**: Air plus froid = altitude indiquée plus haute que réelle
- **Erreur de pression**: QNH incorrect
- **Erreur instrumentale**: Précision limitée`,
			5: `## Le variomètre (VSI)

### Fonctionnement

Le variomètre (Vertical Speed Indicator) mesure la vitesse de variation d'altitude.

### Principe

- Utilise une fuite calibrée entre la capsule statique et le boîtier
- Lorsque l'altitude change, la pression dans la capsule change plus vite que dans le boîtier
- La différence de pression indique la vitesse verticale

### Utilisation

- Indiqué en centaines de pieds par minute (ft/min)
- Utilisé pour maintenir un taux de montée/descente constant
- Essentiel pour les approches de précision
- Réponse légèrement retardée (2-3 secondes)`,
		},
	}

	if contents[category] != nil {
		if content, ok := contents[category][lessonNum]; ok {
			return content
		}
	}
	return "Contenu en cours de rédaction pour " + category + " leçon " + string(rune('0'+lessonNum))
}

func lessonContentEn(category string, lessonNum int) string {

	contents := map[string]map[int]string{
		"airlaw": {
			1: `## Introduction to Air Law and ICAO

Air law is the set of rules governing the use of airspace and aviation activities.

### The International Civil Aviation Organization (ICAO)

ICAO is a specialized agency of the United Nations created in 1944 by the Chicago Convention. Its main role is to:
- Establish Standards and Recommended Practices (SARPs)
- Ensure safety and efficiency of air transport
- Harmonize regulations between member states

### Annexes to the Chicago Convention

The 19 annexes cover all aspects of civil aviation:
- **Annex 1**: Personnel Licensing
- **Annex 2**: Rules of the Air
- **Annex 3**: Meteorology
- **Annex 6**: Operation of Aircraft
- **Annex 8**: Airworthiness of Aircraft`,
			2: `## EASA and European Regulations

The European Union Aviation Safety Agency (EASA) is the competent authority for aviation regulation in Europe.

### Key European Regulations

- **Regulation (EU) 2018/1139**: EASA Basic Regulation
- **Part-FCL**: Flight Crew Licensing
- **Part-M**: Continuing Airworthiness
- **Part-66**: Maintenance Certification
- **Part-147**: Approved Maintenance Training Organizations

### Regulatory Structure

European regulation is organized as:
1. **Basic Regulations**: General framework
2. **Implementing Regulations**: Technical details
3. **AMC (Acceptable Means of Compliance)**
4. **CS (Certification Specifications)**`,
			3: `## Pilot Licenses and Medical Certificates

### Types of Licenses

- **LAPL (Light Aircraft Pilot Licence)**: For light aircraft
- **PPL (Private Pilot Licence)**: Private pilot
- **CPL (Commercial Pilot Licence)**: Professional pilot
- **ATPL (Airline Transport Pilot Licence)**: Airline pilot
- **IR (Instrument Rating)**: Instrument flight

### Medical Certificate

- **Class 1**: For ATPL and CPL (annual exam)
- **Class 2**: For PPL and LAPL (every 2-5 years)
- **LAPL**: Simplified medical exam`,
			4: `## Classified Airspace A to G

### Airspace Classification

- **Class A**: IFR only, permanent control
- **Class B**: IFR and VFR, separation provided
- **Class C**: IFR and VFR, IFR/IFR and IFR/VFR separation
- **Class D**: IFR and VFR, IFR/IFR separation only
- **Class E**: Controlled IFR, VFR with traffic information
- **Class F**: Advisory airspace
- **Class G**: Uncontrolled airspace`,
			5: `## Standardized European Rules of the Air (SERA)

### General Rules

- **Protection of persons and property**: All reasonable measures
- **Collision avoidance**: Right-of-way rules
- **Navigation lights**: Green right, red left, white rear
- **Distress signals**: Mayday (immediate danger), Pan-Pan (urgency)

### VFR Rules

- Minimum visibility: 5 km (or 1500 m depending on conditions)
- Distance from clouds: 1500 m horizontal, 300 m (1000 ft) vertical`,
		},
		"meteorology": {
			1: `## The Earth's Atmosphere: Composition and Structure

### Composition

- **Nitrogen (N₂)**: 78%
- **Oxygen (O₂)**: 21%
- **Argon (Ar)**: 0.93%
- **Trace gases**: CO₂, neon, helium, methane, water vapor

### Vertical Structure

- **Troposphere**: 0-11 km, 80% of mass, temperature decreases with altitude
- **Tropopause**: Boundary between troposphere and stratosphere
- **Stratosphere**: 11-50 km
- **Mesosphere**: 50-85 km
- **Thermosphere**: Beyond 85 km`,
			2: `## Temperature and Heat Exchange

### Heat Sources

- **Solar radiation**: Primary heat source
- **Terrestrial radiation**: Earth re-emits received energy
- **Greenhouse effect**: Greenhouse gases trap heat

### Heat Transfer

- **Conduction**: Direct contact transfer
- **Convection**: Vertical air movement
- **Advection**: Horizontal movement
- **Radiation**: Electromagnetic waves`,
			3: `## Atmospheric Pressure

### Definition

Atmospheric pressure is the weight of the air column above a given point.

### Units

- **Hectopascals (hPa)**: Standard unit
- **Millimeters of mercury (mmHg)**: Historical unit
- **Inches of mercury (inHg)**: US unit

### Pressure Systems

- **Anticyclone (High pressure)**: Descending air, stable weather
- **Depression (Low pressure)**: Ascending air, disturbed weather`,
			4: `## Clouds and Precipitation

### Cloud Classification

- **Cirrus (Ci)**: High clouds, filamentary
- **Cumulus (Cu)**: Vertical development clouds
- **Stratus (St)**: Low layer clouds
- **Nimbus (Nb)**: Precipitation clouds

### Dangerous Clouds

- **Cumulonimbus (Cb)**: Thunderstorms, turbulence, hail, wind shear
- **Nimbostratus**: Continuous rain, reduced visibility`,
			5: `## Winds and Air Masses

### Origin of Wind

Wind is caused by atmospheric pressure differences:
- Air moves from high to low pressure
- Coriolis force deflects wind to the right (northern hemisphere)

### Air Masses

- **Arctic (A)**: Very cold and dry
- **Polar (P)**: Cold and dry
- **Tropical (T)**: Warm and humid
- **Equatorial (E)**: Very hot and very humid`,
		},
		"navigation": {
			1: `## Basic Principles of Air Navigation

### Types of Navigation

- **Visual Navigation (VFR)**: Using visual references
- **Instrument Navigation (IFR)**: Using cockpit instruments
- **Radio Navigation**: Using radio aids
- **GNSS Navigation**: Using satellites

### Fundamental Quantities

- **Heading**: Direction the aircraft points
- **Track**: Planned path over ground
- **Drift**: Angle between heading and track due to wind
- **True Airspeed**: Aircraft speed through air
- **Ground Speed**: Aircraft speed over ground`,
			2: `## The Aeronautical Chart

### Chart Types

- **ICAO 1:500,000**: Visual navigation
- **En-route chart**: IFR navigation
- **Approach chart**: Landing procedures
- **Aerodrome chart**: Airport plan`,
			3: `## The Magnetic Compass

### Operating Principle

The magnetic compass uses the Earth's magnetic field to indicate magnetic north.

### Compass Errors

- **Deviation**: Error from aircraft metal masses
- **Variation**: Difference between true and magnetic north
- **Turning error**: Incorrect indication during turns
- **Acceleration error**: During acceleration/deceleration`,
			4: `## Radio Navigation Aids

### VOR (VHF Omnidirectional Range)

- Ground transmitter emitting on 360 radials
- Provides bearing to or from the station
- Range: Line of sight (about 200 NM at high altitude)

### NDB (Non-Directional Beacon)

- Ground transmitter emitting in all directions
- ADF (Automatic Direction Finder) indicates direction
- Range: Up to 400 NM`,
			5: `## GPS and Modern Navigation Systems

### GPS Principle

GPS uses a constellation of 24 satellites:
- Each satellite transmits its position and time
- Receiver calculates position by trilateration
- Accuracy: 5-10 meters

### Other GNSS Systems

- **GLONASS**: Russian system
- **Galileo**: European system
- **BeiDou**: Chinese system`,
		},
		"aircraft_general": {
			1: `## Aircraft Structure

### Main Components

- **Fuselage**: Main structure containing the cabin
- **Wings**: Lift generation
- **Empennage**: Stability and control (fin + stabilizer)
- **Landing Gear**: Ground support
- **Powerplant**: Engine and propeller`,
			2: `## Aircraft Systems

### Electrical System

- Battery (12V or 24V)
- Alternator or generator
- Main and secondary electrical bus
- Circuit breakers and fuses

### Fuel System

- Tanks (wing, fuselage)
- Boost pump
- Filters and sumps
- Fuel gauge`,
			3: `## The Piston Engine

### Operating Principle

1. Intake: Air-fuel mixture
2. Compression: Piston compresses mixture
3. Power: Spark ignites mixture
4. Exhaust: Burned gases expelled`,
			4: `## The Propeller

### Propeller Types

- **Fixed pitch**: Simple, light
- **Variable pitch**: Performance optimization
- **Constant speed**: Automatic regulation`,
			5: `## Flight Instruments

### Basic Flight Instruments (6-pack)

1. **Airspeed Indicator (ASI)**
2. **Altimeter**
3. **Vertical Speed Indicator (VSI)**
4. **Attitude Indicator**
5. **Directional Gyro (DG)**
6. **Turn Coordinator**`,
		},
		"human_performance": {
			1: `## Human Factors in Aviation

### Introduction

Human factors study the interaction between humans and their work environment. In aviation, they are essential for safety.

### The SHELL Model

- **S (Software)**: Procedures, checklists
- **H (Hardware)**: Equipment, instruments
- **E (Environment)**: Work environment
- **L (Liveware)**: People
- **L (Liveware)**: People interactions`,
			2: `## Vision and Visual Illusions

### Eye Anatomy

- **Retina**: Photoreceptor cells (cones and rods)
- **Fovea**: Central vision (cones)
- **Peripheral vision**: Rods (light sensitive)

### Visual Illusions in Flight

- **Sloping terrain illusion**: Sloping terrain gives false attitude
- **Runway width illusion**: Narrow runway = higher, wide = lower`,
			3: `## The Ear and Balance

### Ear Anatomy

- **Outer ear**: Pinna and ear canal
- **Middle ear**: Eardrum and ossicles
- **Inner ear**: Cochlea (hearing) and vestibule (balance)

### Barotrauma

- Pain from pressure difference
- Occurs mainly during descent
- Valsalva maneuver to equalize`,
			4: `## Hypoxia and Respiration

### Hypoxia

Lack of oxygen in tissues:
- **Hypobaric hypoxia**: Altitude (reduced partial pressure)
- **Anemic hypoxia**: Lack of red blood cells
- **Stagnant hypoxia**: Reduced circulation
- **Histotoxic hypoxia**: Cells unable to use O₂`,
			5: `## Fatigue and Stress

### Fatigue Types

- **Acute fatigue**: After long wake period
- **Chronic fatigue**: Accumulation over days
- **Physical fatigue**: From effort
- **Mental fatigue**: From concentration

### Stress Management

- Deep breathing
- Task prioritization
- Clear communication
- Know your limits`,
		},
		"performance": {
			1: `## Introduction to Aircraft Performance

### Definitions

Aircraft performance describes the capabilities of an aircraft in different flight phases:
- **Takeoff**: Distance and climb gradient
- **Climb**: Rate and gradient of climb
- **Cruise**: Speed and fuel consumption
- **Landing**: Landing distance

### Factors Affecting Performance

- **Weight**: Heavier aircraft have reduced performance
- **Altitude**: Less dense air reduces lift and engine power
- **Temperature**: Hot air is less dense
- **Wind**: Headwind reduces distances, tailwind increases them
- **Runway condition**: Wet or grass runway increases distances`,
			2: `## Weights and Structural Limitations

### Certified Maximum Weights

- **MTOW (Maximum TakeOff Weight)**
- **MLW (Maximum Landing Weight)**
- **MZFW (Maximum Zero Fuel Weight)**
- **MRW (Maximum Ramp Weight)**

### Structural Limitations

- **Load factor**: Positive and negative limits (e.g. +3.8g / -1.52g)
- **Vne (Never Exceed Speed)**
- **Va (Maneuvering Speed)**

### Importance of Limitations

Exceeding maximum weights can cause:
- Increased takeoff and landing distances
- Reduced climb performance
- Risk of structural damage
- Regulatory non-compliance`,
			3: `## Air Density and Performance

### What is Density Altitude?

Density altitude is the altitude corresponding to the actual air density in the standard atmosphere.

### Effects of High Density Altitude

- **Decreased lift**: Wing generates less lift
- **Decreased engine power**: Less oxygen for combustion
- **Increased takeoff distances**: Up to 50% or more
- **Reduced climb rate**: Less power available`,
			4: `## Takeoff Distances

### Components of Takeoff Distance

- **Ground roll**: From start point to rotation
- **Takeoff distance**: Up to 50 ft (15 m) height
- **Accelerate-stop distance**: Distance to accelerate then stop

### Factors Increasing Takeoff Distance

- **High weight**: More inertia to overcome
- **High altitude**: Less dense air
- **High temperature**: Less dense air
- **Tailwind**: Increases required ground speed
- **Uphill runway**: Unfavorable weight component
- **Wet/grass runway**: Reduced friction`,
			5: `## V1, Vr and V2 Speeds

### V1 - Decision Speed

- Maximum speed at which takeoff can be rejected
- Beyond V1, takeoff must be continued

### Vr - Rotation Speed

- Speed at which pilot pulls back to raise the nose
- Calculated to ensure V2 at 35 or 50 ft

### V2 - Takeoff Safety Speed

- Minimum speed after engine failure at takeoff
- Ensures sufficient climb gradient
- Generally 1.2 x Vs (stall speed)`,
		},
		"flight_planning": {
			1: `## Introduction to Flight Planning

### Objectives

Flight planning prepares a safe flight by determining:
- Route and altitudes
- Required fuel
- Alternate aerodromes
- NOTAM and restrictions

### Planning Steps

1. **Weather briefing**: Forecasts and actual conditions
2. **NOTAM**: Restrictions and temporary information
3. **Route selection**: Airways, reporting points
4. **Fuel calculation**: Route, reserve, alternate
5. **Weight and balance**: Check limitations
6. **Flight plan**: Submit form`,
			2: `## NOTAM and Their Interpretation

### What is a NOTAM?

NOTAM (Notice to Air Missions) contains essential safety information:
- Runway or aerodrome closures
- Radio aids out of service
- Military exercises or danger areas
- Important temporary changes

### Types of NOTAM

- **NOTAMN**: New NOTAM
- **NOTAMR**: NOTAM replacing a previous one
- **NOTAMC**: NOTAM canceling a previous one`,
			3: `## Route Selection and Altitudes

### Route Selection

- **Airways**: Standardized IFR routes
- **Visual navigation**: Following visual references
- **Reporting points**: Beacons, intersections, waypoints

### Cruise Altitudes

- **VFR semicircular rule**:
  - Magnetic track 000°-179°: odd + 500 ft
  - Magnetic track 180°-359°: even + 500 ft
- **IFR flight levels**: Based on standard 1013 hPa setting`,
			4: `## Fuel Calculation

### Fuel Components

- **Route fuel**: From departure to destination
- **Alternate fuel**: To alternate aerodrome
- **Reserve**: 30 min VFR day, 45 min VFR night
- **Taxi fuel**: Ground consumption

### Practical Calculation

1. Determine total distance
2. Calculate estimated flight time
3. Multiply by hourly consumption
4. Add regulatory reserves`,
			5: `## Regulatory Fuel Reserves

### VFR Reserves

- **VFR day**: 30 minutes above destination
- **VFR night**: 45 minutes above destination

### IFR Reserves

- Fuel to reach alternate aerodrome
- Plus 30 minutes reserve above alternate

### In-flight Fuel Management

- Monitor consumption regularly
- Compare remaining fuel to remaining time
- When in doubt, land to refuel`,
		},
		"operational_procedures": {
			1: `## Introduction to Operational Procedures

### Definition

Operational procedures are standardized actions to be performed for each flight phase.

### Importance of Procedures

- **Standardization**: All pilots perform same actions
- **Safety**: Reduced errors and omissions
- **Efficiency**: Optimized flight conduct
- **Communication**: Common language between pilots and controllers

### Checklists

- Essential tool to avoid forgetting items
- Must be used at each key phase
- Read and do, not memorize`,
			2: `## Flight Preparation

### Weather Briefing

- Check METAR, TAF, SIGMET
- Analyze weather charts
- Evaluate conditions along route

### Aircraft Inspection

- **Exterior**: Complete pre-flight inspection
- **Interior**: Check documents and instruments
- **Levels**: Oil, fuel, fluids`,
			3: `## Pre-flight Inspections

### Exterior Inspection

- **Propeller**: Blade condition, attachment
- **Engine**: Oil levels, cables, hoses
- **Landing gear**: Tires, shock absorbers, brakes
- **Wings**: General condition, flaps, ailerons
- **Empennage**: Fin, elevator, rudder
- **Lights**: Navigation, landing, anti-collision
- **Fuel tanks**: Fuel level, cap`,
			4: `## Engine Start Procedure

### Starting Procedure

1. **Parking brake**: SET
2. **Battery**: ON
3. **Instruments**: Check
4. **Mixture**: RICH
5. **Starter**: ENGAGE
6. **After start**: Oil pressure check, alternator check`,
			5: `## Taxiing and Runway Procedures

### Taxi Procedure

- **Brakes**: Tested before moving
- **Rudder**: Directional control with rudder pedals
- **Differential brakes**: For tight turns
- **Speed**: No faster than a brisk walk

### Runway Incursion Prevention

- Read back all runway crossing instructions
- Stop at holding point before entering runway
- Visually check for traffic before entering
- Use landing lights when on runway

### After Landing

- Clear runway completely before stopping
- Contact ground control when clear
- Follow taxi instructions to parking`,
		},
	}

	if contents[category] != nil {
		if content, ok := contents[category][lessonNum]; ok {
			return content
		}
	}
	return "Content under development for " + category + " lesson " + string(rune('0'+lessonNum))
}
