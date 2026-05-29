package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type QuestionData struct {
	subtopic   string
	difficulty int
	fr         string
	en         string
	options    string
	answer     string
	explainFr  string
	explainEn  string
}

var licenses = []string{"PPL", "LAPL", "ATPL", "CPL", "IR", "BPL"}
var categories = []string{
	"airlaw", "meteorology", "navigation", "performance",
	"aircraft_general", "flight_planning", "human_performance",
	"operational_procedures", "principles_of_flight",
	"communications", "mass_and_balance", "instrumentation",
}

// lessonsPerCategory defines how many lessons per category (realistic based on EASA PPL/ATPL syllabi)
// Each lesson represents a chapter in the official manuals
var lessonsPerCategory = map[string]int{
	"airlaw":                20,
	"meteorology":           25,
	"navigation":            25,
	"performance":           15,
	"aircraft_general":      20,
	"flight_planning":       20,
	"human_performance":     20,
	"operational_procedures": 25,
	"principles_of_flight":  20,
	"communications":        15,
	"mass_and_balance":      15,
	"instrumentation":       20,
}

func escapeSQL(s string) string {
	return strings.ReplaceAll(s, "'", "''")
}

func repeatQuestions(qs []QuestionData, count int) []QuestionData {

	if count >= len(qs) {
		return qs
	}
	perm := rand.Perm(len(qs))
	result := make([]QuestionData, count)
	for i := 0; i < count; i++ {
		result[i] = qs[perm[i]]
	}
	return result
}

func genAirLaw(count int) []QuestionData {
	return repeatQuestions([]QuestionData{
		{"regulations", 1, "Quel organisme regule l aviation civile en Europe ?", "Which body regulates civil aviation in Europe?", `["EASA","FAA","ICAO","IATA"]`, "EASA", "L EASA est l agence europeenne de la securite aerienne.", "EASA is the European aviation safety agency."},
		{"regulations", 1, "Quel document un pilote doit il presenter avant un vol ?", "Which document must a pilot show before a flight?", `["Licence de pilote","Passeport","Carte identite","Permis conduire"]`, "Licence de pilote", "La licence de pilote en cours de validite est obligatoire.", "A valid pilot license is mandatory."},
		{"regulations", 2, "Duree de validite certificat medical classe 1 avant 40 ans ?", "How long is Class 1 medical valid before age 40?", `["6 mois","12 mois","24 mois","5 ans"]`, "12 mois", "12 mois pour les pilotes de moins de 40 ans.", "12 months for pilots under 40."},
		{"regulations", 2, "Age minimum pour une licence PPL ?", "Minimum age for a PPL?", `["16 ans","17 ans","18 ans","21 ans"]`, "17 ans", "L age minimum pour la PPL est 17 ans.", "The minimum age for PPL is 17."},
		{"regulations", 3, "Que signifie espace aerien classe C ?", "What does Class C airspace mean?", `["Controle obligatoire","Libre acces","Zone militaire","Espace interdit"]`, "Controle obligatoire", "La classe C necessite une clairance du controle aerien.", "Class C requires ATC clearance."},
		{"regulations", 3, "Priorite planeur vs avion motorise ?", "Right-of-way glider vs powered aircraft?", `["Le planeur","L avion motorise","Le plus rapide","Le plus lent"]`, "Le planeur", "Les planeurs ont priorite sur les aeronefs motorises.", "Gliders have priority over powered aircraft."},
		{"regulations", 2, "A quoi servent les regles SERA ?", "What are SERA rules for?", `["Unifier regles air Europe","Certifier avions","Former pilotes","Controler trafic"]`, "Unifier regles air Europe", "Les SERA standardisent les regles de l air dans toute l Europe.", "SERA standardizes air rules across Europe."},
		{"regulations", 1, "Limite heures vol annuelle pilote prive ?", "Annual flight hour limit private pilot?", `["Non","100 heures","200 heures","500 heures"]`, "Non", "Pas de limite stricte pour les pilotes prives.", "No strict limit for private pilots."},
		{"regulations", 3, "Vol VFR de nuit autorise sans qualification ?", "Night VFR allowed without rating?", `["Non qualification requise","Oui toujours","Oui avec instructeur","Non jamais"]`, "Non qualification requise", "Une qualification VFR de nuit specifique est necessaire.", "A specific night VFR rating is required."},
		{"regulations", 2, "Panne radio en zone controlee ?", "Radio failure in controlled airspace?", `["Suivre procedure standard","Continuer","Se poser","Changer frequence"]`, "Suivre procedure standard", "Procedure panne radio squawk 7600 suivre plan de vol.", "Radio failure squawk 7600 follow flight plan."},
		{"regulations", 1, "Qu est ce que l OACI ?", "What is ICAO?", `["Organisation aviation civile ONU","Agence europeenne","Bureau federal","Compagnie aerienne"]`, "Organisation aviation civile ONU", "L OACI est l organisation internationale de l aviation civile.", "ICAO is the International Civil Aviation Organization."},
		{"regulations", 2, "Frequence de detresse en aviation ?", "Aviation distress frequency?", `["121.5 MHz","118.0 MHz","135.0 MHz","126.7 MHz"]`, "121.5 MHz", "121.5 MHz est la frequence de detresse internationale.", "121.5 MHz is the international distress frequency."},
		{"regulations", 3, "Code transpondeur 7500 ?", "Transponder code 7500?", `["Detournement","Panne radio","Urgence","Vol normal"]`, "Detournement", "7500 signale un detournement d aeronef.", "7500 signals aircraft hijacking."},
		{"regulations", 2, "Code transpondeur 7600 ?", "Transponder code 7600?", `["Panne radio","Detournement","Urgence medicale","Vol normal"]`, "Panne radio", "7600 indique une panne de communication radio.", "7600 indicates radio communication failure."},
		{"regulations", 1, "Code transpondeur 7700 ?", "Transponder code 7700?", `["Urgence generale","Panne radio","Detournement","Vol normal"]`, "Urgence generale", "7700 signale une situation d urgence a bord.", "7700 signals an emergency on board."},
		{"regulations", 2, "Qu est ce que la classe d espace aerien A ?", "What is Class A airspace?", `["Vol IFR seulement","Vol VFR seulement","Tous vols","Vol libre"]`, "Vol IFR seulement", "Classe A reservee aux vols IFR avec clairance.", "Class A reserved for IFR flights with clearance."},
		{"regulations", 2, "Qu est ce que la classe d espace aerien G ?", "What is Class G airspace?", `["Espace non controle","Espace controle","Zone militaire","Zone interdite"]`, "Espace non controle", "Classe G espace aerien non controle sans clairance requise.", "Class G uncontrolled airspace no clearance needed."},
		{"regulations", 3, "Altitude minimale survol agglomeration ?", "Minimum altitude over built-up areas?", `["1000 ft","500 ft","2000 ft","3000 ft"]`, "1000 ft", "Survol agglomeration a 1000 ft au-dessus obstacle le plus haut.", "Over built-up areas 1000 ft above highest obstacle."},
		{"regulations", 1, "Distance minimale d un nuage en VFR ?", "Minimum distance from cloud in VFR?", `["1500 m horizontal 300 m vertical","500 m","1000 m","2000 m"]`, "1500 m horizontal 300 m vertical", "VFR doit rester a 1500 m des nuages horizontalement 300 m verticalement.", "VFR must stay 1500 m from clouds horizontally 300 m vertically."},
		{"regulations", 2, "Validite d un plan de vol VFR ?", "Validity of a VFR flight plan?", `["30 min apres EOBT","1 heure","2 heures","15 min"]`, "30 min apres EOBT", "Plan de vol valide 30 minutes apres l heure estimee de depart.", "Flight plan valid 30 minutes after estimated departure time."},
		{"regulations", 3, "Que signifie le code transpondeur 7000 ?", "What does transponder code 7000 mean?", `["Vol VFR standard","Detresse","Panne radio","Vol IFR"]`, "Vol VFR standard", "7000 est le code VFR standard en Europe.", "7000 is the standard VFR code in Europe."},
		{"regulations", 1, "Qui est responsable de la separation en VFR ?", "Who is responsible for separation in VFR?", `["Le pilote lui-meme","Le controleur","Le copilote","La tour"]`, "Le pilote lui-meme", "En VFR le pilote est responsable de sa propre separation.", "In VFR the pilot is responsible for their own separation."},
		{"regulations", 2, "Conditions minimales pour vol VFR special ?", "Minimum conditions for special VFR?", `["Visibilite 1500 m hors nuages","Visibilite 5 km","Visibilite 8 km","Visibilite 500 m"]`, "Visibilite 1500 m hors nuages", "VFR special necessite visibilite 1500 m et rester hors nuages.", "Special VFR requires 1500 m visibility and clear of clouds."},
		{"regulations", 3, "Quand un plan de vol est il obligatoire ?", "When is a flight plan mandatory?", `["Franchissement frontieres vol IFR nuit","Toujours","Jamais","Vol local"]`, "Franchissement frontieres vol IFR nuit", "Plan de vol obligatoire pour franchir frontieres vol IFR et de nuit.", "Flight plan mandatory for borders IFR and night flights."},
		{"regulations", 2, "Duree de validite d une licence PPL ?", "Validity period of a PPL?", `["5 ans","1 an","10 ans","Illimitee"]`, "5 ans", "La licence PPL est valable 5 ans.", "The PPL is valid for 5 years."},
		{"regulations", 1, "Que signifie le terme Mayday ?", "What does Mayday mean?", `["Detresse mortelle","Urgence medicale","Panne moteur","Incendie"]`, "Detresse mortelle", "Mayday signale un danger grave et imminent.", "Mayday signals grave and imminent danger."},
		{"regulations", 2, "Que signifie le terme Pan-Pan ?", "What does Pan-Pan mean?", `["Urgence","Detresse","Panne","Information"]`, "Urgence", "Pan-Pan signale une situation d urgence sans danger immediat.", "Pan-Pan signals urgency without immediate danger."},
		{"regulations", 3, "Altitude minimale de survol en campagne ?", "Minimum altitude over countryside?", `["500 ft","1000 ft","2000 ft","300 ft"]`, "500 ft", "Altitude minimale 500 ft sauf decollage atterrissage.", "Minimum altitude 500 ft except takeoff landing."},
		{"regulations", 1, "Qu est ce qu une zone interdite P ?", "What is a prohibited area P?", `["Survol interdit en permanence","Survol reglemente","Zone dangereuse","Zone militaire"]`, "Survol interdit en permanence", "Zone P survol interdit en toutes circonstances.", "Zone P overflight prohibited under all circumstances."},
		{"regulations", 2, "Qu est ce qu une zone reglementee R ?", "What is a restricted area R?", `["Survol reglemente conditions specifiques","Survol interdit","Zone libre","Zone dangereuse"]`, "Survol reglemente conditions specifiques", "Zone R survol autorise sous certaines conditions.", "Zone R overflight allowed under certain conditions."},
		{"regulations", 3, "Qu est ce qu une zone dangereuse D ?", "What is a dangerous area D?", `["Activites dangereuses possibles","Survol interdit","Zone interdite","Zone militaire"]`, "Activites dangereuses possibles", "Zone D activites dangereuses peuvent s y derouler.", "Zone D dangerous activities may take place."},
		{"regulations", 1, "A quelle frequence appeler le controle ?", "On which frequency to call ATC?", `["Frequence indiquee sur carte","121.5 MHz","123.5 MHz","118.0 MHz"]`, "Frequence indiquee sur carte", "Utiliser la frequence indiquee sur la carte aeronautique.", "Use the frequency indicated on the aeronautical chart."},
		{"regulations", 2, "Que faire en cas d interception ?", "What to do in case of interception?", `["Suivre instructions intercepteur","Ignorer","Accelerer","Virer"]`, "Suivre instructions intercepteur", "Suivre les instructions de l aeronef intercepteur.", "Follow the intercepting aircraft instructions."},
		{"regulations", 3, "Signaux d interception jour ?", "Interception signals during day?", `["Battement ailerons jour allumage/extinction feux nuit","Tirer","Pomper","Klaxonner"]`, "Battement ailerons jour allumage/extinction feux nuit", "Signaux standardises pour interception.", "Standardized signals for interception."},
		{"regulations", 1, "Qu est ce que le droit de passage ?", "What is right of way?", `["Regle priorite entre aeronefs","Droit atterrir","Droit decoller","Droit survoler"]`, "Regle priorite entre aeronefs", "Le droit de passage determine quel aeronef passe en premier.", "Right of way determines which aircraft passes first."},
		{"regulations", 2, "Priorite entre montgolfiere et avion ?", "Right of way balloon vs aircraft?", `["La montgolfiere","L avion","Le plus rapide","Le plus lent"]`, "La montgolfiere", "Les montgolfieres ont priorite sur les avions.", "Balloons have priority over aircraft."},
		{"regulations", 3, "Priorite entre dirigeable et planeur ?", "Right of way airship vs glider?", `["Le planeur","Le dirigeable","Le plus haut","Le plus bas"]`, "Le planeur", "Les planeurs ont priorite sur les dirigeables.", "Gliders have priority over airships."},
		{"regulations", 2, "Que faire si deux avions en approche face a face ?", "What if two aircraft approaching head-on?", `["Virer a droite chacun","Virer a gauche","Monter","Descendre"]`, "Virer a droite chacun", "Face a face chaque pilote vire a droite.", "Head-on each pilot turns right."},
		{"regulations", 1, "Que faire si deux avions routes convergentes ?", "What if two aircraft on converging courses?", `["Celui de droite a priorite","Celui de gauche","Le plus rapide","Le plus lent"]`, "Celui de droite a priorite", "L aeronef venant de droite a priorite.", "Aircraft from the right has right of way."},
		{"regulations", 3, "Depassement entre aeronefs ?", "Overtaking between aircraft?", `["L aeronef depasse doit virer a droite","L aeronef depasse doit virer a gauche","L aeronef depasse doit monter","L aeronef depasse doit descendre"]`, "L aeronef depasse doit virer a droite", "L aeronef depasse vire a droite pour faciliter le depassement.", "The overtaken aircraft turns right to facilitate overtaking."},
		{"regulations", 2, "Altitude de croisiere VFR de jour ?", "VFR cruising altitude during day?", `["Route magnetique 0-179 impair+500 ft 180-359 pair+500 ft","Altitude libre","Altitude minimale","Altitude maximale"]`, "Route magnetique 0-179 impair+500 ft 180-359 pair+500 ft", "Altitudes semi-circulaires VFR selon route magnetique.", "VFR semi-circular altitudes based on magnetic track."},
		{"regulations", 1, "Qu est ce que le QNH ?", "What is QNH?", `["Pression ramenee au niveau de la mer","Pression aerodrome","Pression standard","Pression cabine"]`, "Pression ramenee au niveau de la mer", "QNH permet d afficher l altitude au-dessus du niveau de la mer.", "QNH displays altitude above mean sea level."},
		{"regulations", 2, "Qu est ce que le QFE ?", "What is QFE?", `["Pression aerodrome","Pression niveau mer","Pression standard","Pression cabine"]`, "Pression aerodrome", "QFE permet d afficher la hauteur au-dessus de l aerodrome.", "QFE displays height above aerodrome."},
		{"regulations", 3, "Transition niveau de vol altitude ?", "Transition altitude to flight level?", `["Alt transition 5000 ft","3000 ft","10000 ft","2000 ft"]`, "Alt transition 5000 ft", "Au-dessus de l altitude de transition on utilise les niveaux de vol.", "Above transition altitude flight levels are used."},
		{"regulations", 1, "Qu est ce que le niveau de vol FL ?", "What is Flight Level FL?", `["Altitude basee sur 1013 hPa","Altitude vraie","Hauteur","Altitude pression"]`, "Altitude basee sur 1013 hPa", "FL altitude basee sur le calage standard 1013 hPa.", "FL altitude based on standard setting 1013 hPa."},
		{"regulations", 2, "Validite certificat medical classe 2 ?", "Class 2 medical validity?", `["5 ans moins de 40 ans 2 ans apres","1 an","10 ans","3 ans"]`, "5 ans moins de 40 ans 2 ans apres", "Classe 2 valable 5 ans avant 40 ans 2 ans apres.", "Class 2 valid 5 years before 40 2 years after."},
		{"regulations", 3, "Age minimum licence LAPL ?", "Minimum age for LAPL?", `["16 ans","17 ans","18 ans","15 ans"]`, "16 ans", "L age minimum pour la LAPL est 16 ans.", "The minimum age for LAPL is 16."},
		{"regulations", 1, "Qu est ce que la licence LAPL ?", "What is LAPL license?", `["Licence avion leger","Licence transport","Licence planeur","Licence helicoptere"]`, "Licence avion leger", "LAPL Light Aircraft Pilot Licence pour avions legers.", "LAPL Light Aircraft Pilot Licence for light aircraft."},
		{"regulations", 2, "Masse maximale pour LAPL ?", "Maximum weight for LAPL?", `["2000 kg","1000 kg","3000 kg","5000 kg"]`, "2000 kg", "LAPL limitee aux avions de moins de 2000 kg.", "LAPL limited to aircraft under 2000 kg."},
		{"regulations", 3, "Nombre passagers maximum LAPL ?", "Maximum passengers LAPL?", `["3 passagers","1 passager","4 passagers","5 passagers"]`, "3 passagers", "LAPL autorise maximum 3 passagers.", "LAPL allows maximum 3 passengers."},
		{"regulations", 1, "Qu est ce que la licence ATPL ?", "What is ATPL license?", `["Licence pilote transport ligne","Licence privee","Licence planeur","Licence loisir"]`, "Licence pilote transport ligne", "ATPL Airline Transport Pilot Licence pour pilotes de ligne.", "ATPL Airline Transport Pilot Licence for airline pilots."},
		{"regulations", 2, "Heures de vol requises pour ATPL ?", "Flight hours required for ATPL?", `["1500 heures","200 heures","500 heures","1000 heures"]`, "1500 heures", "ATPL necessite 1500 heures de vol.", "ATPL requires 1500 flight hours."},
		{"regulations", 3, "Age minimum pour ATPL ?", "Minimum age for ATPL?", `["21 ans","18 ans","23 ans","25 ans"]`, "21 ans", "L age minimum pour l ATPL est 21 ans.", "The minimum age for ATPL is 21."},
		{"regulations", 1, "Qu est ce que la licence CPL ?", "What is CPL license?", `["Licence pilote professionnel","Licence privee","Licence loisir","Licence planeur"]`, "Licence pilote professionnel", "CPL Commercial Pilot Licence pour pilotes professionnels.", "CPL Commercial Pilot Licence for professional pilots."},
		{"regulations", 2, "Heures de vol requises pour CPL ?", "Flight hours required for CPL?", `["200 heures","1500 heures","500 heures","100 heures"]`, "200 heures", "CPL necessite 200 heures de vol.", "CPL requires 200 flight hours."},
		{"regulations", 3, "Age minimum pour CPL ?", "Minimum age for CPL?", `["18 ans","21 ans","16 ans","20 ans"]`, "18 ans", "L age minimum pour le CPL est 18 ans.", "The minimum age for CPL is 18."},
		{"regulations", 1, "Qu est ce que la qualification IR ?", "What is IR rating?", `["Vol aux instruments","Vol de nuit","Vol acrobatique","Vol planeur"]`, "Vol aux instruments", "IR Instrument Rating permet de voler aux instruments.", "IR Instrument Rating allows flying on instruments."},
		{"regulations", 2, "Heures de vol requises pour IR ?", "Flight hours required for IR?", `["50 heures vol aux instruments","200 heures","100 heures","500 heures"]`, "50 heures vol aux instruments", "IR necessite 50 heures de vol aux instruments.", "IR requires 50 hours of instrument flight."},
		{"regulations", 3, "Qu est ce que la licence BPL ?", "What is BPL license?", `["Licence pilote ballon","Licence planeur","Licence helicoptere","Licence avion"]`, "Licence pilote ballon", "BPL Balloon Pilot Licence pour pilotes de montgolfiere.", "BPL Balloon Pilot Licence for balloon pilots."},
		{"regulations", 1, "Age minimum pour BPL ?", "Minimum age for BPL?", `["16 ans","18 ans","17 ans","21 ans"]`, "16 ans", "L age minimum pour la BPL est 16 ans.", "The minimum age for BPL is 16."},
		{"regulations", 2, "Qu est ce que le certificat medical classe 1 ?", "What is Class 1 medical?", `["Pour ATPL CPL IR","Pour PPL","Pour LAPL","Pour planeur"]`, "Pour ATPL CPL IR", "Classe 1 requise pour ATPL CPL et IR.", "Class 1 required for ATPL CPL and IR."},
		{"regulations", 3, "Qu est ce que le certificat medical classe 3 ?", "What is Class 3 medical?", `["Pour controleurs aeriens","Pour pilotes prives","Pour mecaniciens","Pour hotesses"]`, "Pour controleurs aeriens", "Classe 3 pour les controleurs de la circulation aerienne.", "Class 3 for air traffic controllers."},
		{"regulations", 1, "Qu est ce que l AIP ?", "What is AIP?", `["Publication information aeronautique","Plan vol","Carte","Manuel vol"]`, "Publication information aeronautique", "AIP contient toutes les informations aeronautiques officielles.", "AIP contains all official aeronautical information."},
		{"regulations", 2, "Qu est ce qu un NOTAM ?", "What is a NOTAM?", `["Avis navigateurs aeriens","Prevision meteo","Plan vol","Carte"]`, "Avis navigateurs aeriens", "NOTAM avis urgent sur les changements temporaires.", "NOTAM urgent notice on temporary changes."},
		{"regulations", 3, "Qu est ce que le SUP AIP ?", "What is SUP AIP?", `["Supplement temporaire AIP","Nouvel AIP","Carte supplementaire","Manuel supplementaire"]`, "Supplement temporaire AIP", "SUP AIP supplement temporaire a l AIP.", "SUP AIP temporary supplement to AIP."},
		{"regulations", 1, "Qu est ce que le service AFIS ?", "What is AFIS service?", `["Service information vol aerodrome","Controle aerien","Radar","Tour de controle"]`, "Service information vol aerodrome", "AFIS fournit des informations sans controle.", "AFIS provides information without control."},
		{"regulations", 2, "Difference entre AFIS et controle ?", "Difference between AFIS and ATC?", `["AFIS informe ATC controle","AFIS controle ATC informe","Identique","AFIS radar ATC tour"]`, "AFIS informe ATC controle", "AFIS donne des informations ATC donne des instructions.", "AFIS gives information ATC gives instructions."},
		{"regulations", 3, "Qu est ce que le service FIS ?", "What is FIS service?", `["Service information vol","Controle aerien","Radar","Approche"]`, "Service information vol", "FIS Flight Information Service informations de vol.", "FIS Flight Information Service provides flight information."},
	}, count)
}

func genMeteo(count int) []QuestionData {
	return repeatQuestions([]QuestionData{
		{"clouds", 1, "Comment se forment les stratus ?", "How do stratus clouds form?", `["Refroidissement couche air humide","Convection thermique","Soulevement orographique","Evaporation"]`, "Refroidissement couche air humide", "Les stratus se forment quand une couche d air humide se refroidit.", "Stratus forms when moist air layer cools."},
		{"clouds", 1, "Nuage associe aux turbulences severes ?", "Cloud associated with severe turbulence?", `["Cumulonimbus","Stratus","Cirrus","Altocumulus"]`, "Cumulonimbus", "Les cumulonimbus ont de puissants courants ascendants et descendants.", "Cumulonimbus have powerful updrafts and downdrafts."},
		{"clouds", 2, "Altitude des cirrus ?", "Altitude of cirrus clouds?", `["Au-dessus 6000 m","2000-4000 m","Sous 2000 m","Pres sol"]`, "Au-dessus 6000 m", "Les cirrus sont des nuages eleves de cristaux de glace.", "Cirrus are high-level ice crystal clouds."},
		{"clouds", 1, "Qu est ce qu un cumulus ?", "What is a cumulus cloud?", `["Nuage beau temps base plate","Nuage pluie","Nuage glace","Brouillard"]`, "Nuage beau temps base plate", "Cumulus nuage de beau temps base plate sommet arrondi.", "Cumulus fair-weather cloud flat base rounded top."},
		{"clouds", 2, "Nuage associe a precipitations continues ?", "Cloud with continuous precipitation?", `["Nimbostratus","Cumulus","Cirrus","Altocumulus"]`, "Nimbostratus", "Nimbostratus epaisse couche nuageuse precipitations continues.", "Nimbostratus thick cloud layer continuous precipitation."},
		{"clouds", 1, "Qu est ce que l altocumulus ?", "What is altocumulus?", `["Nuage moyen blanc gris","Nuage bas","Nuage haut","Brouillard"]`, "Nuage moyen blanc gris", "Altocumulus nuage moyen compose de gouttelettes d eau.", "Altocumulus mid-level cloud of water droplets."},
		{"clouds", 2, "Qu est ce que l altostratus ?", "What is altostratus?", `["Voile nuageux moyen gris","Nuage bas","Nuage haut","Brouillard"]`, "Voile nuageux moyen gris", "Altostratus voile nuageux moyen laissant voir soleil flou.", "Altostratus mid-level veil blurring sun."},
		{"clouds", 3, "Qu est ce que le cirrostratus ?", "What is cirrostratus?", `["Voile nuageux haut cristaux glace","Nuage bas","Brouillard","Cumulus"]`, "Voile nuageux haut cristaux glace", "Cirrostratus voile haut de cristaux de glace halo autour soleil.", "Cirrostratus high veil of ice crystals halo around sun."},
		{"clouds", 2, "Qu est ce que le stratocumulus ?", "What is stratocumulus?", `["Couche nuageuse basse rouleaux","Nuage haut","Brouillard","Cumulonimbus"]`, "Couche nuageuse basse rouleaux", "Stratocumulus couche basse en rouleaux ou plaques.", "Stratocumulus low layer in rolls or patches."},
		{"clouds", 3, "Nuage donnant pluie faible continue ?", "Cloud giving light continuous rain?", `["Stratus","Cumulus","Cirrus","Altocumulus"]`, "Stratus", "Stratus peut donner bruine ou pluie faible continue.", "Stratus can give drizzle or light continuous rain."},
		{"clouds", 1, "Nuage donnant averses ?", "Cloud giving showers?", `["Cumulus congestus","Stratus","Cirrus","Altostratus"]`, "Cumulus congestus", "Cumulus congestus peut donner averses.", "Cumulus congestus can give showers."},
		{"clouds", 2, "Qu est ce que le cumulonimbus calvus ?", "What is cumulonimbus calvus?", `["Cb sommet arrondi sans enclume","Cb avec enclume","Cb sans pluie","Cb faible"]`, "Cb sommet arrondi sans enclume", "Calvus stade jeune cumulonimbus sans enclume glace.", "Calvus young cumulonimbus without ice anvil."},
		{"clouds", 3, "Qu est ce que le cumulonimbus incus ?", "What is cumulonimbus incus?", `["Cb avec enclume glace sommet","Cb jeune","Cb faible","Cb sans pluie"]`, "Cb avec enclume glace sommet", "Incus cumulonimbus mature avec enclume glace sommet.", "Incus mature cumulonimbus with ice anvil top."},
		{"clouds", 2, "Base typique cumulus beau temps ?", "Typical base fair-weather cumulus?", `["1000-2000 m","0-500 m","5000-6000 m","8000 m"]`, "1000-2000 m", "Base cumulus beau temps entre 1000 et 2000 m.", "Fair-weather cumulus base between 1000 and 2000 m."},
		{"clouds", 1, "Nuages bas altitude ?", "Low clouds altitude?", `["Sous 2000 m","2000-6000 m","Au-dessus 6000 m","Pres sol"]`, "Sous 2000 m", "Nuages bas stratus stratocumulus sous 2000 m.", "Low clouds stratus stratocumulus below 2000 m."},
		{"clouds", 2, "Nuages moyens altitude ?", "Medium clouds altitude?", `["2000-6000 m","Sous 2000 m","Au-dessus 6000 m","Pres sol"]`, "2000-6000 m", "Nuages moyens altocumulus altostratus 2000-6000 m.", "Medium clouds altocumulus altostratus 2000-6000 m."},
		{"clouds", 1, "Nuages hauts altitude ?", "High clouds altitude?", `["Au-dessus 6000 m","2000-6000 m","Sous 2000 m","Pres sol"]`, "Au-dessus 6000 m", "Nuages hauts cirrus cirrostratus cirrocumulus au-dessus 6000 m.", "High clouds cirrus cirrostratus cirrocumulus above 6000 m."},
		{"pressure", 1, "Baisse rapide pression atmospherique ?", "Rapid pressure drop indicates?", `["Approche depression","Arrivee beau temps","Vent faible","Temps stable"]`, "Approche depression", "Baisse rapide de pression annonce l arrivee d une depression.", "Rapid pressure drop signals arrival of a depression."},
		{"pressure", 2, "Difference entre QNH et QFE ?", "Difference between QNH and QFE?", `["QNH niveau mer QFE aerodrome","QNH standard QFE local","Identiques","QNH atterrissage QFE decollage"]`, "QNH niveau mer QFE aerodrome", "QNH pression ramenee niveau mer QFE pression aerodrome.", "QNH pressure reduced to sea level QFE aerodrome level."},
		{"pressure", 3, "Variation pression avec altitude ?", "Pressure variation with altitude?", `["Diminue exponentiellement","Augmente lineairement","Constante","Diminue lineairement"]`, "Diminue exponentiellement", "Pression divisee par 2 tous les 5500 m.", "Pressure halves every 5500 m."},
		{"pressure", 1, "Pression standard niveau mer ?", "Standard pressure at sea level?", `["1013.25 hPa","1000 hPa","1030 hPa","980 hPa"]`, "1013.25 hPa", "1013.25 hPa pression standard au niveau de la mer.", "1013.25 hPa standard pressure at sea level."},
		{"pressure", 2, "Unite mesure pression atmospherique ?", "Unit for atmospheric pressure?", `["hPa ou mb","mm Hg","psi","bar"]`, "hPa ou mb", "Pression mesuree en hectopascals ou millibars.", "Pressure measured in hectopascals or millibars."},
		{"pressure", 3, "Isobare definition ?", "Isobar definition?", `["Ligne egale pression","Ligne egale temperature","Ligne egale altitude","Ligne egale vent"]`, "Ligne egale pression", "Isobare relie points de meme pression atmospherique.", "Isobar connects points of equal pressure."},
		{"pressure", 2, "Pression typique centre depression ?", "Typical pressure low pressure center?", `["980-1005 hPa","1013 hPa","1020-1040 hPa","1050 hPa"]`, "980-1005 hPa", "Depression pression centrale entre 980 et 1005 hPa.", "Low pressure center between 980 and 1005 hPa."},
		{"pressure", 1, "Pression typique centre anticyclone ?", "Typical pressure high pressure center?", `["1020-1040 hPa","980-1005 hPa","1013 hPa","950 hPa"]`, "1020-1040 hPa", "Anticyclone pression centrale entre 1020 et 1040 hPa.", "High pressure center between 1020 and 1040 hPa."},
		{"pressure", 3, "Qu est ce que le gradient de pression ?", "What is pressure gradient?", `["Difference pression sur distance","Variation pression temps","Pression moyenne","Pression standard"]`, "Difference pression sur distance", "Gradient difference pression entre deux points.", "Gradient pressure difference between two points."},
		{"pressure", 2, "Isobares serrees signifient ?", "Tight isobars mean?", `["Vent fort","Vent faible","Beau temps","Pluie"]`, "Vent fort", "Isobares serrees gradient fort vent fort.", "Tight isobars strong gradient strong wind."},
		{"winds", 2, "Qu est ce que l effet de foehn ?", "What is the foehn effect?", `["Vent chaud sec descendant montagne","Vent froid mer","Brise vallee","Vent catabatique"]`, "Vent chaud sec descendant montagne", "Effet de foehn vent chaud sec versant sous le vent.", "Foehn effect warm dry wind leeward side."},
		{"winds", 1, "Vent mer vers terre le jour ?", "Wind sea to land during day?", `["Brise de mer","Brise de terre","Mousson","Alize"]`, "Brise de mer", "Brise de mer causee par le rechauffement diurne du sol.", "Sea breeze caused by daytime heating of land."},
		{"winds", 3, "Qu est ce que le cisaillement de vent ?", "What is wind shear?", `["Changement brusque direction ou vitesse vent","Vent constant altitude","Absence vent","Vent tourbillonnaire"]`, "Changement brusque direction ou vitesse vent", "Cisaillement changement rapide du vent sur courte distance.", "Wind shear rapid wind change over short distance."},
		{"winds", 2, "Qu est ce que le jet stream ?", "What is the jet stream?", `["Courant air rapide haute altitude","Type nuage","Instrument navigation","Vent surface"]`, "Courant air rapide haute altitude", "Jet stream courant tres rapide entre 9000 et 15000 m.", "Jet stream fast current between 9000 and 15000 m."},
		{"winds", 1, "Vent terre vers mer la nuit ?", "Wind land to sea at night?", `["Brise de terre","Brise de mer","Mousson","Alize"]`, "Brise de terre", "Brise de terre causee par refroidissement nocturne du sol.", "Land breeze caused by nighttime cooling of land."},
		{"winds", 2, "Qu est ce que le vent catabatique ?", "What is katabatic wind?", `["Vent descendant pente gravite","Vent ascendant","Vent chaud","Vent mer"]`, "Vent descendant pente gravite", "Vent catabatique air froid descendant pente.", "Katabatic wind cold air flowing down slope."},
		{"winds", 3, "Qu est ce que le vent anabatique ?", "What is anabatic wind?", `["Vent ascendant pente","Vent descendant","Vent mer","Vent froid"]`, "Vent ascendant pente", "Vent anabatique air chaud montant pente.", "Anabatic wind warm air rising up slope."},
		{"winds", 2, "Qu est ce que la brise de vallee ?", "What is valley breeze?", `["Vent montant vallee jour","Vent descendant vallee","Vent mer","Vent froid"]`, "Vent montant vallee jour", "Brise de vallee vent remontant vallee le jour.", "Valley breeze wind up valley during day."},
		{"winds", 1, "Qu est ce que la brise de montagne ?", "What is mountain breeze?", `["Vent descendant vallee nuit","Vent montant vallee","Vent mer","Vent chaud"]`, "Vent descendant vallee nuit", "Brise de montagne vent descendant vallee la nuit.", "Mountain breeze wind down valley at night."},
		{"winds", 3, "Qu est ce que le vent geostrophique ?", "What is geostrophic wind?", `["Vent parallele isobares haute altitude","Vent surface","Vent local","Vent thermique"]`, "Vent parallele isobares haute altitude", "Vent geostrophique equilibre gradient force Coriolis.", "Geostrophic wind balance gradient Coriolis force."},
		{"winds", 2, "Qu est ce que le vent de surface ?", "What is surface wind?", `["Vent mesure 10 m sol","Vent haute altitude","Vent geostrophique","Vent thermique"]`, "Vent mesure 10 m sol", "Vent surface vent pres du sol influence frottement.", "Surface wind near ground affected by friction."},
		{"winds", 1, "Direction vent est donnee par ?", "Wind direction is given by?", `["Direction d ou vient vent","Direction ou va vent","Nord","Sud"]`, "Direction d ou vient vent", "Direction vent d ou il vient pas ou il va.", "Wind direction from where it comes not where it goes."},
		{"winds", 2, "Qu est ce que la rotation du vent avec altitude ?", "What is wind rotation with altitude?", `["Rotation droite en montant","Rotation gauche","Pas rotation","Rotation aleatoire"]`, "Rotation droite en montant", "Vent tourne droite en montant cause frottement.", "Wind turns right with altitude due to friction."},
		{"visibility", 2, "Qu est ce que le brouillard de rayonnement ?", "What is radiation fog?", `["Forme par refroidissement nocturne","Brouillard advection","Brouillard pente","Brouillard precipitation"]`, "Forme par refroidissement nocturne", "Brouillard de rayonnement nuit par refroidissement du sol.", "Radiation fog forms at night through ground cooling."},
		{"visibility", 1, "Visibilite minimale VFR hors zone controlee ?", "Minimum visibility VFR outside controlled airspace?", `["1500 m","5 km","8 km","10 km"]`, "1500 m", "1500 m visibilite minimale VFR hors zone controlee.", "1500 m minimum VFR visibility outside controlled airspace."},
		{"visibility", 3, "Que signifie RVR ?", "What does RVR mean?", `["Portee visuelle piste","Visibilite vol","Distance freinage","Longueur piste"]`, "Portee visuelle piste", "RVR distance jusqu ou pilote voit marques piste.", "RVR distance pilot sees runway markings."},
		{"visibility", 2, "Qu est ce que le brouillard d advection ?", "What is advection fog?", `["Air chaud humide sur surface froide","Refroidissement nocturne","Evaporation","Convection"]`, "Air chaud humide sur surface froide", "Brouillard advection air chaud humide passant sur mer froide.", "Advection fog warm moist air over cold sea."},
		{"visibility", 3, "Qu est ce que le brouillard de pente ?", "What is hill fog?", `["Air humide monte pente refroidit","Refroidissement nocturne","Evaporation","Rayonnement"]`, "Air humide monte pente refroidit", "Brouillard pente air humide souleve orographiquement refroidit.", "Hill fog moist air lifted orographically cools."},
		{"visibility", 1, "Qu est ce que la brume ?", "What is mist?", `["Visibilite 1-5 km","Visibilite sous 1 km","Visibilite 5-10 km","Visibilite 10+ km"]`, "Visibilite 1-5 km", "Brume visibilite entre 1 et 5 km.", "Mist visibility between 1 and 5 km."},
		{"visibility", 2, "Qu est ce que le brouillard ?", "What is fog?", `["Visibilite sous 1 km","Visibilite 1-5 km","Visibilite 5-10 km","Visibilite 10+ km"]`, "Visibilite sous 1 km", "Brouillard visibilite inferieure a 1 km.", "Fog visibility less than 1 km."},
		{"visibility", 3, "Qu est ce que la brume seche haze ?", "What is haze?", `["Particules fines suspension","Gouttelettes eau","Cristaux glace","Pollen"]`, "Particules fines suspension", "Brume seche particules fines poussiere pollution.", "Haze fine particles dust pollution."},
		{"visibility", 2, "Conditions formation brouillard rayonnement ?", "Conditions for radiation fog?", `["Nuit claire vent faible humidite elevee","Jour ensoleille","Vent fort","Pluie"]`, "Nuit claire vent faible humidite elevee", "Nuit claire ciel degage vent faible humidite haute.", "Clear night light wind high humidity."},
		{"icing", 2, "Conditions givrage les plus probables ?", "Conditions most likely for icing?", `["Air humide pres 0C","Temps sec chaud","Haute altitude","Au-dessus 30C"]`, "Air humide pres 0C", "Givrage probable air humide temperature proche 0C.", "Icing likely humid air temperature near 0C."},
		{"icing", 3, "Type givrage le plus dangereux ?", "Most dangerous icing type?", `["Givre transparent verglas","Givre blanc","Givre carburateur","Givre sol"]`, "Givre transparent verglas", "Givre transparent adhere fortement modifie profil aerodynamique.", "Clear ice adheres strongly modifies aerodynamic profile."},
		{"icing", 1, "Qu est ce que le givre blanc ?", "What is rime ice?", `["Givre opaque blanc granuleux","Glace transparente","Givre carburateur","Givre sol"]`, "Givre opaque blanc granuleux", "Givre blanc givre opaque granuleux facile a enlever.", "Rime ice opaque granular easy to remove."},
		{"icing", 2, "Qu est ce que le givre carburateur ?", "What is carburetor icing?", `["Glace forme dans carburateur","Givre aile","Givre helice","Givre parebrise"]`, "Glace forme dans carburateur", "Givre carburateur cause par refroidissement evaporation carburant.", "Carburetor icing from fuel evaporation cooling."},
		{"icing", 3, "Conditions givre carburateur ?", "Carburetor icing conditions?", `["Humidite 50%+ temperature 0-20C","Temperature negative","Temperature positive","Humidite faible"]`, "Humidite 50%+ temperature 0-20C", "Givre carburateur probable humidite 50%+ temperature 0-20C.", "Carburetor icing likely humidity 50%+ temp 0-20C."},
		{"icing", 1, "Comment detecter givre carburateur ?", "How to detect carburetor icing?", `["Ralentissement moteur perte puissance","Bruit","Vibration","Fumee"]`, "Ralentissement moteur perte puissance", "Perte puissance moteur signe givre carburateur.", "Engine power loss sign of carburetor icing."},
		{"icing", 2, "Comment corriger givre carburateur ?", "How to correct carburetor icing?", `["Chauffer carburateur","Reduire puissance","Augmenter richesse","Refroidir"]`, "Chauffer carburateur", "Chauffer carburateur avec air chaud echappement.", "Apply carburetor heat with exhaust warm air."},
		{"icing", 3, "Qu est ce que le givre de descente ?", "What is descent icing?", `["Givre en descente air froid","Givre montee","Givre sol","Givre carburateur"]`, "Givre en descente air froid", "Givre descente cause par refroidissement en descente rapide.", "Descent icing from cooling during rapid descent."},
		{"thunderstorms", 2, "Trois phases vie d un orage ?", "Three life stages of a thunderstorm?", `["Cumulus mature dissipation","Naissance vie mort","Formation pluie fin","Convection precipitation evaporation"]`, "Cumulus mature dissipation", "Phase cumulus developpement mature precipitation dissipation affaiblissement.", "Cumulus development mature precipitation dissipation weakening."},
		{"thunderstorms", 1, "Dangers d un orage pour un avion ?", "Thunderstorm dangers to aircraft?", `["Turbulences grele cisaillement vent","Pluie","Aucun","Brouillard"]`, "Turbulences grele cisaillement vent", "Orages turbulences severes grele cisaillement vent foudre.", "Thunderstorms severe turbulence hail wind shear lightning."},
		{"thunderstorms", 2, "Distance minimale contournement orage ?", "Minimum distance to avoid thunderstorm?", `["20 km","5 km","50 km","100 km"]`, "20 km", "Contourner orage a au moins 20 km.", "Avoid thunderstorm at least 20 km."},
		{"thunderstorms", 3, "Qu est ce qu un orage supercellulaire ?", "What is a supercell thunderstorm?", `["Orage violent tournant","Orage faible","Orage sec","Orage neige"]`, "Orage violent tournant", "Supercellule orage violent avec courant ascendant tournant.", "Supercell violent storm with rotating updraft."},
		{"thunderstorms", 1, "Altitude sommet cumulonimbus ?", "Cumulonimbus top altitude?", `["Jusque 15000 m","5000 m","8000 m","2000 m"]`, "Jusque 15000 m", "Cumulonimbus peut atteindre 15000 m tropopause.", "Cumulonimbus can reach 15000 m tropopause."},
		{"thunderstorms", 2, "Qu est ce que la foudre ?", "What is lightning?", `["Decharge electrique orage","Eclair chaleur","Bruit","Lumiere"]`, "Decharge electrique orage", "Foudre decharge electrique entre nuages ou sol.", "Lightning electrical discharge between clouds or ground."},
		{"thunderstorms", 3, "Qu est ce que la grele ?", "What is hail?", `["Glace formee orage","Pluie gelee","Neige","Givre"]`, "Glace formee orage", "Grele grumeaux glace formes dans cumulonimbus.", "Hail ice lumps formed in cumulonimbus."},
		{"thunderstorms", 2, "Qu est ce qu un orage frontal ?", "What is frontal thunderstorm?", `["Orage le long front froid","Orage chaud","Orage orographique","Orage nocturne"]`, "Orage le long front froid", "Orage frontal se forme le long d un front froid.", "Frontal thunderstorm forms along cold front."},
		{"thunderstorms", 1, "Qu est ce qu un orage thermique ?", "What is thermal thunderstorm?", `["Orage convection diurne","Orage frontal","Orage orographique","Orage nocturne"]`, "Orage convection diurne", "Orage thermique convection air chaud jour.", "Thermal thunderstorm daytime hot air convection."},
		{"fronts", 2, "Qu est ce qu un front froid ?", "What is a cold front?", `["Air froid pousse air chaud","Air chaud pousse air froid","Zone stable","Zone haute pression"]`, "Air froid pousse air chaud", "Front froid masse air froid avance remplace air chaud.", "Cold front cold air advances replaces warm air."},
		{"fronts", 1, "Qu est ce qu un front chaud ?", "What is a warm front?", `["Air chaud pousse air froid","Air froid pousse air chaud","Zone stable","Zone haute pression"]`, "Air chaud pousse air froid", "Front chaud masse air chaud avance remplace air froid.", "Warm front warm air advances replaces cold air."},
		{"fronts", 3, "Qu est ce qu un front occlus ?", "What is an occluded front?", `["Front froid rattrape front chaud","Front double","Front stationnaire","Front faible"]`, "Front froid rattrape front chaud", "Occlusion front froid rattrape front chaud.", "Occlusion cold front catches up with warm front."},
		{"fronts", 2, "Temps typique front froid ?", "Typical weather cold front?", `["Averses orages puis eclaircie","Pluie continue","Beau temps","Brouillard"]`, "Averses orages puis eclaircie", "Front froid averses orages puis eclaircie rapide.", "Cold front showers storms then rapid clearing."},
		{"fronts", 1, "Temps typique front chaud ?", "Typical weather warm front?", `["Pluie continue puis eclaircie","Averses orages","Beau temps","Brouillard"]`, "Pluie continue puis eclaircie", "Front chaud pluie continue puis eclaircie.", "Warm front continuous rain then clearing."},
		{"fronts", 3, "Qu est ce qu un front stationnaire ?", "What is a stationary front?", `["Front ne bouge pas","Front double","Front froid","Front chaud"]`, "Front ne bouge pas", "Front stationnaire reste en place sans mouvement.", "Stationary front stays in place without movement."},
		{"fronts", 2, "Nuages associes front chaud ?", "Clouds associated with warm front?", `["Cirrus altostratus nimbostratus","Cumulonimbus","Cumulus","Stratus"]`, "Cirrus altostratus nimbostratus", "Front chaud nuages etages cirrus puis altostratus puis nimbostratus.", "Warm front layered clouds cirrus then altostratus then nimbostratus."},
		{"fronts", 1, "Nuages associes front froid ?", "Clouds associated with cold front?", `["Cumulonimbus cumulus","Stratus","Cirrus","Altostratus"]`, "Cumulonimbus cumulus", "Front froid cumulonimbus et cumulus.", "Cold front cumulonimbus and cumulus."},
		{"fronts", 2, "Pente front chaud vs front froid ?", "Warm front vs cold front slope?", `["Front chaud pente faible front froid pente raide","Identique","Front chaud raide","Front froid faible"]`, "Front chaud pente faible front froid pente raide", "Front chaud pente douce front froid pente raide.", "Warm front gentle slope cold front steep slope."},
		{"fronts", 3, "Vitesse deplacement front froid ?", "Cold front movement speed?", `["40-60 km/h","10-20 km/h","100-120 km/h","5-10 km/h"]`, "40-60 km/h", "Front froid se deplace a 40-60 km/h.", "Cold front moves at 40-60 km/h."},
		{"fronts", 1, "Qu est ce qu une masse d air ?", "What is an air mass?", `["Grande volume air caracteristiques homogenes","Petite zone air","Nuage","Vent"]`, "Grande volume air caracteristiques homogenes", "Masse d air grande etendue air temperature humidite uniformes.", "Air mass large area uniform temperature humidity."},
		{"fronts", 2, "Masse d air continentale arctique ?", "Continental arctic air mass?", `["Tres froide seche","Froide humide","Chaude seche","Chaude humide"]`, "Tres froide seche", "Air arctique continental tres froid et sec.", "Continental arctic very cold and dry."},
		{"fronts", 3, "Masse d air maritime tropical ?", "Maritime tropical air mass?", `["Chaude humide","Froide seche","Froide humide","Chaude seche"]`, "Chaude humide", "Air maritime tropical chaud et humide.", "Maritime tropical warm and humid."},
		{"fronts", 2, "Masse d air continentale polaire ?", "Continental polar air mass?", `["Froide seche","Froide humide","Chaude seche","Chaude humide"]`, "Froide seche", "Air continental polaire froid et sec.", "Continental polar cold and dry."},
		{"fronts", 1, "Masse d air maritime polaire ?", "Maritime polar air mass?", `["Froide humide","Froide seche","Chaude humide","Chaude seche"]`, "Froide humide", "Air maritime polaire froid et humide.", "Maritime polar cold and humid."},
		{"metar", 2, "Que signifie METAR ?", "What does METAR mean?", `["Rapport meteo aerodrome","Prevision meteo","Carte meteo","Message alerte"]`, "Rapport meteo aerodrome", "METAR rapport d observation meteorologique regulier.", "METAR routine meteorological report."},
		{"metar", 1, "Frequence emission METAR ?", "METAR emission frequency?", `["Toutes les 30 min","Toutes les heures","Toutes les 2h","Toutes les 15 min"]`, "Toutes les 30 min", "METAR emis toutes les 30 minutes.", "METAR issued every 30 minutes."},
		{"metar", 3, "Que signifie TAF ?", "What does TAF mean?", `["Prevision aerodrome","Rapport meteo","Carte meteo","Message alerte"]`, "Prevision aerodrome", "TAF Terminal Aerodrome Forecast prevision 24-30h.", "TAF Terminal Aerodrome Forecast 24-30h forecast."},
		{"metar", 2, "Que signifie SIGMET ?", "What does SIGMET mean?", `["Alerte meteo severe","Rapport meteo","Prevision","Message routine"]`, "Alerte meteo severe", "SIGMET alerte phenomenes meteorologiques dangereux.", "SIGMET alert dangerous weather phenomena."},
		{"metar", 1, "Que signifie AIRMET ?", "What does AIRMET mean?", `["Alerte meteo legere","Alerte severe","Rapport","Prevision"]`, "Alerte meteo legere", "AIRMET alerte conditions meteorologiques moins severes que SIGMET.", "AIRMET alert less severe than SIGMET."},
		{"temperature", 2, "Qu est ce que le gradient adiabatique ?", "What is adiabatic lapse rate?", `["Taux refroidissement air monte","Taux rechauffement","Taux pression","Taux humidite"]`, "Taux refroidissement air monte", "Gradient adiabatique sec 3C/1000 ft humide 1.5C/1000 ft.", "Dry adiabatic lapse rate 3C/1000 ft moist 1.5C/1000 ft."},
		{"temperature", 1, "Qu est ce que l inversion de temperature ?", "What is temperature inversion?", `["Temperature augmente avec altitude","Temperature diminue","Temperature constante","Temperature nulle"]`, "Temperature augmente avec altitude", "Inversion couche ou temperature augmente avec altitude.", "Inversion layer where temperature increases with altitude."},
		{"temperature", 3, "Effet inversion sur visibilite ?", "Inversion effect on visibility?", `["Emprisonne pollution brouillard","Ameliore visibilite","Pas effet","Vent fort"]`, "Emprisonne pollution brouillard", "Inversion emprisonne polluants et brouillard pres sol.", "Inversion traps pollutants and fog near ground."},
		{"temperature", 2, "Qu est ce que l isotherme ?", "What is an isotherm?", `["Ligne egale temperature","Ligne egale pression","Ligne egale altitude","Ligne egale vent"]`, "Ligne egale temperature", "Isotherme relie points de meme temperature.", "Isotherm connects points of equal temperature."},
		{"temperature", 1, "Temperature standard niveau mer ?", "Standard temperature sea level?", `["15C","20C","10C","25C"]`, "15C", "Temperature standard 15C au niveau de la mer.", "Standard temperature 15C at sea level."},
		{"temperature", 3, "Qu est ce que la tropopause ?", "What is the tropopause?", `["Limite entre troposphere et stratosphere","Couche nuageuse","Zone vent","Couche glace"]`, "Limite entre troposphere et stratosphere", "Tropopause temperature cesse de diminuer avec altitude.", "Tropopause temperature stops decreasing with altitude."},
		{"temperature", 2, "Altitude de la tropopause ?", "Tropopause altitude?", `["11000 m","5000 m","20000 m","3000 m"]`, "11000 m", "Tropopause environ 11000 m sous nos latitudes.", "Tropopause about 11000 m in our latitudes."},
		{"temperature", 1, "Couches de l atmosphere ?", "Atmosphere layers?", `["Troposphere stratosphere mesosphere thermosphere","Troposphere seulement","Stratosphere seulement","Mesosphere seulement"]`, "Troposphere stratosphere mesosphere thermosphere", "Atmosphere 4 couches principales.", "Atmosphere 4 main layers."},
		{"temperature", 2, "Epaisseur de la troposphere ?", "Troposphere thickness?", `["8-16 km selon latitude","5 km","20 km","30 km"]`, "8-16 km selon latitude", "Troposphere plus epaisse equator plus fine poles.", "Troposphere thicker at equator thinner at poles."},
	}, count)
}

func genNavigation(count int) []QuestionData {
	return repeatQuestions([]QuestionData{
		{"vor", 2, "Comment identifier une station VOR ?", "How to identify a VOR station?", `["Code Morse","Frequence","Position geographique","Nom"]`, "Code Morse", "Chaque station VOR emet un code Morse d identification.", "Each VOR transmits an identification Morse code."},
		{"vor", 1, "Plage de frequences VOR ?", "VOR frequency range?", `["108-118 MHz","118-137 MHz","30-88 MHz","960-1215 MHz"]`, "108-118 MHz", "Les VOR fonctionnent dans la bande VHF 108-118 MHz.", "VORs operate in VHF band 108-118 MHz."},
		{"vor", 3, "Qu est ce qu un radial VOR ?", "What is a VOR radial?", `["Direction magnetique depuis station","Distance station","Altitude","Frequence secondaire"]`, "Direction magnetique depuis station", "Radial ligne de position depuis station VOR direction magnetique.", "Radial line of position from VOR in magnetic direction."},
		{"vor", 2, "Precision typique d un VOR ?", "Typical VOR accuracy?", `["+-1 degre","+-5 degres","+-10 degres","+-0.1 degre"]`, "+-1 degre", "Le VOR offre une precision d environ +-1 degre.", "VOR accuracy approximately +-1 degree."},
		{"vor", 1, "Portee maximale VOR basse altitude ?", "Max range low altitude VOR?", `["40-130 NM selon altitude","200 NM","300 NM","500 NM"]`, "40-130 NM selon altitude", "Portee VOR limitee par ligne vue directe.", "VOR range limited by line of sight."},
		{"vor", 2, "Qu est ce que le VOT ?", "What is VOT?", `["Equipement test VOR sol","Type VOR","Frequence VOR","Radial VOR"]`, "Equipement test VOR sol", "VOT permet verifier precision recepteur VOR bord.", "VOT allows checking onboard VOR receiver accuracy."},
		{"vor", 3, "Que signifie TO/FROM sur VOR ?", "What does TO/FROM on VOR mean?", `["Indique direction station","Distance station","Altitude station","Frequence station"]`, "Indique direction station", "TO vers station FROM depuis station.", "TO towards station FROM away from station."},
		{"vor", 2, "Qu est ce que le CDI ?", "What is CDI?", `["Indicateur deviation radial","Compas","Altimetre","Variometre"]`, "Indicateur deviation radial", "CDI Course Deviation Indicator montre ecart radial.", "CDI shows deviation from selected radial."},
		{"vor", 1, "Que signifie OBS sur VOR ?", "What does OBS on VOR mean?", `["Selecteur radial","Bouton marche","Volume","Luminosite"]`, "Selecteur radial", "OBS Omni Bearing Selector choisit le radial.", "OBS selects the desired radial."},
		{"ndb", 1, "Difference principale NDB vs VOR ?", "Main difference NDB vs VOR?", `["NDB donne direction VOR donne radial","NDB plus precis","NDB utilise UHF","NDB plus recent"]`, "NDB donne direction VOR donne radial", "NDB indique direction station VOR donne radial precis.", "NDB indicates direction to station VOR gives precise radial."},
		{"ndb", 2, "Effet de nuit sur un NDB ?", "Night effect on NDB?", `["Deviation signal coucher soleil","Amplification signal","Disparition signal","Aucun effet"]`, "Deviation signal coucher soleil", "Effet de nuit deviation signal NDB par changements ionospheriques.", "Night effect NDB signal deviation from ionospheric changes."},
		{"ndb", 1, "Instrument qui recoit signaux NDB ?", "Instrument that receives NDB signals?", `["ADF","VOR","DME","ILS"]`, "ADF", "L ADF est l instrument qui recoit les signaux NDB.", "The ADF receives NDB station signals."},
		{"ndb", 3, "Plage frequences NDB ?", "NDB frequency range?", `["200-1750 kHz","108-118 MHz","960-1215 MHz","30-88 MHz"]`, "200-1750 kHz", "NDB fonctionne en ondes kilometriques LF/MF.", "NDB operates in LF/MF band."},
		{"ndb", 2, "Qu est ce que l effet de cote sur NDB ?", "What is coastal effect on NDB?", `["Deviation signal pres cote","Amplification","Disparition","Aucun"]`, "Deviation signal pres cote", "Effet cote deviation onde radio en passant terre-mer.", "Coastal effect radio wave deviation land-sea."},
		{"ndb", 1, "Avantage NDB sur VOR ?", "NDB advantage over VOR?", `["Portee plus longue","Plus precis","Moins interference","Plus fiable"]`, "Portee plus longue", "NDB portee plus longue que VOR surtout nuit.", "NDB longer range than VOR especially at night."},
		{"gps", 1, "Satellites necessaires pour position GPS 3D ?", "Satellites needed for 3D GPS?", `["4 satellites","3 satellites","2 satellites","6 satellites"]`, "4 satellites", "Recepteur GPS besoin d au moins 4 satellites pour position 3D.", "GPS receiver needs at least 4 satellites for 3D."},
		{"gps", 2, "Que signifie RAIM ?", "What does RAIM mean?", `["Systeme verification integrite","Mode navigation","Type antenne","Frequence GPS"]`, "Systeme verification integrite", "RAIM verifie la fiabilite et l integrite des signaux GPS.", "RAIM checks reliability and integrity of GPS signals."},
		{"gps", 3, "Principale limitation du GPS en aeronautique ?", "Main GPS limitation in aviation?", `["Vulnerabilite interferences radio","Manque precision","Couverture insuffisante","Consommation elevee"]`, "Vulnerabilite interferences radio", "GPS vulnerable aux interferences radio et brouillage.", "GPS vulnerable to radio interference and jamming."},
		{"gps", 1, "Precision du GPS civil ?", "Civilian GPS accuracy?", `["5-10 m","1 m","50 m","100 m"]`, "5-10 m", "GPS civil precision d environ 5-10 metres.", "Civilian GPS accuracy about 5-10 meters."},
		{"gps", 2, "Nombre total satellites GPS ?", "Total GPS satellites?", `["31 satellites","24 satellites","48 satellites","12 satellites"]`, "31 satellites", "Constellation GPS environ 31 satellites operationnels.", "GPS constellation about 31 operational satellites."},
		{"gps", 3, "Que signifie WAAS ?", "What does WAAS mean?", `["Systeme augmentation GPS precision","Type recepteur","Antenne GPS","Frequence GPS"]`, "Systeme augmentation GPS precision", "WAAS ameliore precision GPS pour approches.", "WAAS improves GPS accuracy for approaches."},
		{"gps", 1, "Frequence signal GPS civil ?", "Civilian GPS signal frequency?", `["1575.42 MHz L1","122.7 MHz","108 MHz","960 MHz"]`, "1575.42 MHz L1", "Signal GPS civil sur frequence L1 1575.42 MHz.", "Civilian GPS signal on L1 1575.42 MHz."},
		{"gps", 2, "Qu est ce que le GPS differentiel ?", "What is differential GPS?", `["Correction sol pour precision","GPS militaire","GPS double","GPS simple"]`, "Correction sol pour precision", "DGPS utilise station sol pour corriger erreurs GPS.", "DGPS uses ground station to correct GPS errors."},
		{"dme", 2, "Que mesure le DME ?", "What does DME measure?", `["Distance oblique station","Distance horizontale","Altitude","Vitesse"]`, "Distance oblique station", "DME mesure distance oblique entre avion et station.", "DME measures slant distance between aircraft and station."},
		{"dme", 1, "Bande frequences DME ?", "DME frequency band?", `["960-1215 MHz","108-118 MHz","118-137 MHz","200-1750 kHz"]`, "960-1215 MHz", "DME fonctionne en bande UHF 960-1215 MHz.", "DME operates in UHF band 960-1215 MHz."},
		{"dme", 3, "Precision DME ?", "DME accuracy?", `["+-0.1 NM ou 1%","+-1 NM","+-5 NM","+-10 NM"]`, "+-0.1 NM ou 1%", "DME precision +-0.1 NM ou 1% distance.", "DME accuracy +-0.1 NM or 1% distance."},
		{"dme", 2, "Qu est ce que le DME arc ?", "What is DME arc?", `["Arc cercle autour station DME","Type DME","Frequence DME","Antenne DME"]`, "Arc cercle autour station DME", "DME arc procedure navigation en arc autour station.", "DME arc navigation procedure around station."},
		{"flight_computing", 2, "Calcul temps de vol entre deux points ?", "Flight time calculation?", `["Distance / Vitesse sol","Distance / Vitesse air","Distance x Vitesse","Carburant / Consommation"]`, "Distance / Vitesse sol", "Temps de vol = distance / vitesse sol.", "Flight time = distance / ground speed."},
		{"flight_computing", 3, "Qu est ce que la derive drift ?", "What is drift?", `["Angle entre cap et route","Difference vitesse air sol","Variation altitude","Temps perdu"]`, "Angle entre cap et route", "Derive angle entre cap pointe et route reelle causee par vent.", "Drift angle between heading and track caused by wind."},
		{"flight_computing", 2, "Comment corriger la derive ?", "How to correct drift?", `["Modifiant cap","Augmentant vitesse","Changeant altitude","Reduisant puissance"]`, "Modifiant cap", "On corrige la derive en modifiant le cap pour compenser le vent.", "Drift corrected by adjusting heading to compensate for wind."},
		{"flight_computing", 1, "Qu est ce que la vitesse sol ?", "What is ground speed?", `["Vitesse avion par rapport sol","Vitesse air","Vitesse vent","Vitesse moteur"]`, "Vitesse avion par rapport sol", "Vitesse sol vitesse de l avion mesuree par rapport au sol.", "Ground speed aircraft speed measured relative to ground."},
		{"flight_computing", 2, "Qu est ce que la vitesse vraie TAS ?", "What is true airspeed TAS?", `["Vitesse air corrigee densite altitude","Vitesse indiquee","Vitesse sol","Vitesse vent"]`, "Vitesse air corrigee densite altitude", "TAS vitesse air reelle tenant compte densite altitude.", "TAS actual airspeed accounting for density altitude."},
		{"flight_computing", 1, "Qu est ce que la vitesse indiquee IAS ?", "What is indicated airspeed IAS?", `["Vitesse lue sur anemometre","Vitesse sol","Vitesse vraie","Vitesse vent"]`, "Vitesse lue sur anemometre", "IAS vitesse directement lue sur l anemometre.", "IAS speed directly read on airspeed indicator."},
		{"flight_computing", 3, "Qu est ce que le triangle des vitesses ?", "What is the wind triangle?", `["Representation vectorielle vent cap route","Calcul altitude","Calcul carburant","Calcul temps"]`, "Representation vectorielle vent cap route", "Triangle vitesses combine vent cap et route.", "Wind triangle combines wind heading and track."},
		{"flight_computing", 2, "Qu est ce que le cap vrai ?", "What is true heading?", `["Direction pointe avion nord vrai","Direction nord magnetique","Route","Cap compas"]`, "Direction pointe avion nord vrai", "Cap vrai angle entre axe avion et nord vrai.", "True heading angle between aircraft axis and true north."},
		{"flight_computing", 1, "Qu est ce que la route track ?", "What is track?", `["Trajectoire reelle sol avion","Direction pointe","Cap compas","Route prevue"]`, "Trajectoire reelle sol avion", "Route chemin reel parcouru par avion au sol.", "Track actual path of aircraft over ground."},
		{"maps", 2, "Qu est ce que la projection Lambert ?", "What is Lambert projection?", `["Projection conique conforme","Projection Mercator","Projection polaire","Projection cylindrique"]`, "Projection conique conforme", "Lambert projection conique conserve angles.", "Lambert conformal conic projection preserves angles."},
		{"maps", 1, "Echelle carte aeronautique ?", "Aeronautical chart scale?", `["1:500000","1:100000","1:10000","1:1000000"]`, "1:500000", "Cartes OACI echelle 1:500000.", "ICAO charts scale 1:500000."},
		{"maps", 3, "Qu est ce que la declinaison magnetique ?", "What is magnetic declination?", `["Angle nord vrai nord magnetique","Erreur compas","Variation annuelle","Deviation"]`, "Angle nord vrai nord magnetique", "Declinaison difference entre nord vrai et nord magnetique.", "Declination difference between true north and magnetic north."},
		{"maps", 2, "Qu est ce que la deviation compas ?", "What is compass deviation?", `["Erreur due masses metalliques","Difference nord vrai","Variation annuelle","Erreur lecture"]`, "Erreur due masses metalliques", "Deviation causee par masses metalliques dans avion.", "Deviation caused by metallic masses in aircraft."},
		{"maps", 1, "Qu est ce que le nord magnetique ?", "What is magnetic north?", `["Direction pointe aiguille compas","Nord geographique","Nord carte","Nord vrai"]`, "Direction pointe aiguille compas", "Nord magnetique direction indiquee par compas.", "Magnetic north direction indicated by compass."},
		{"maps", 2, "Qu est ce que le nord vrai ?", "What is true north?", `["Pole Nord geographique","Nord compas","Nord magnetique","Nord carte"]`, "Pole Nord geographique", "Nord vrai direction du pole Nord geographique.", "True north direction to geographic North Pole."},
	}, count)
}

func genPerformance(count int) []QuestionData {
	return repeatQuestions([]QuestionData{
		{"takeoff", 1, "Facteurs augmentant distance decollage ?", "Factors increasing takeoff distance?", `["Poids eleve temperature elevee altitude elevee","Poids faible temperature basse","Vent face","Piste mouillee"]`, "Poids eleve temperature elevee altitude elevee", "Ces facteurs reduisent densite air augmentent distance decollage.", "These factors reduce air density increase takeoff distance."},
		{"takeoff", 2, "Effet vent face sur distance decollage ?", "Headwind effect on takeoff distance?", `["La reduit","L augmente","Pas effet","Double"]`, "La reduit", "Vent face augmente portance a vitesse plus faible reduit distance.", "Headwind increases lift at lower speed reduces distance."},
		{"takeoff", 3, "Qu est ce que la densite altitude ?", "What is density altitude?", `["Altitude pression corrigee temperature","Altitude pression","Altitude vraie","Altitude indiquee"]`, "Altitude pression corrigee temperature", "Densite-altitude represente densite reelle air par rapport standard.", "Density altitude represents actual air density relative to standard."},
		{"landing", 1, "Effet vent arriere a l atterrissage ?", "Tailwind effect on landing?", `["Augmente distance atterrissage","Reduit distance","Pas effet","Ameliore confort"]`, "Augmente distance atterrissage", "Vent arriere augmente vitesse sol donc distance arret.", "Tailwind increases ground speed therefore stopping distance."},
		{"landing", 2, "Qu est ce que l effet de sol ?", "What is ground effect?", `["Augmentation portance pres sol","Perte portance","Effet vent","Bruit particulier"]`, "Augmentation portance pres sol", "Effet sol augmente portance reduit trainee quand aile proche sol.", "Ground effect increases lift reduces drag near ground."},
		{"climb", 2, "Effet temperature sur performances montee ?", "Temperature effect on climb performance?", `["Temperature elevee reduit taux montee","Temperature elevee ameliore","Pas effet","Negligeable"]`, "Temperature elevee reduit taux montee", "Air chaud moins dense reduit portance et puissance moteur.", "Hot less dense air reduces lift and engine power."},
		{"climb", 3, "Difference taux montee vs pente montee ?", "Rate of climb vs climb gradient?", `["Taux ft/min Pente % degres","Taux degres Pente ft/min","Identique","Taux vitesse Pente distance"]`, "Taux ft/min Pente % degres", "Taux vitesse verticale pente angle ou pourcentage.", "Rate vertical speed gradient angle or percentage."},
		{"takeoff", 2, "Que represente V1 ?", "What does V1 represent?", `["Vitesse decision decollage","Vitesse rotation","Vitesse decollage","Vitesse croisiere"]`, "Vitesse decision decollage", "V1 vitesse max pour interrompre le decollage.", "V1 max speed to abort takeoff."},
		{"takeoff", 2, "Que represente Vr ?", "What does Vr represent?", `["Vitesse rotation","Vitesse decision","Vitesse decollage","Vitesse approche"]`, "Vitesse rotation", "Vr vitesse ou on tire sur le manche pour decoller.", "Vr speed at which pilot rotates for takeoff."},
		{"takeoff", 2, "Que represente V2 ?", "What does V2 represent?", `["Vitesse securite decollage","Vitesse rotation","Vitesse decision","Vitesse atterrissage"]`, "Vitesse securite decollage", "V2 vitesse minimale securite apres panne moteur decollage.", "V2 minimum safe speed after engine failure on takeoff."},
		{"takeoff", 1, "Effet altitude aerodrome sur decollage ?", "Airport altitude effect on takeoff?", `["Augmente distance decollage","Reduit distance","Pas effet","Ameliore"]`, "Augmente distance decollage", "Altitude elevee air moins dense performances reduites.", "High altitude less dense air reduced performance."},
		{"takeoff", 2, "Effet piste mouillee sur decollage ?", "Wet runway effect on takeoff?", `["Augmente distance decollage","Reduit distance","Pas effet","Ameliore"]`, "Augmente distance decollage", "Piste mouillee reduit acceleration et freinage.", "Wet runway reduces acceleration and braking."},
		{"takeoff", 3, "Qu est ce que la pente de decollage ?", "What is takeoff climb gradient?", `["Pente montee apres decollage","Distance decollage","Vitesse decollage","Hauteur decollage"]`, "Pente montee apres decollage", "Pente decollage angle montee apres avoir quitte sol.", "Takeoff gradient climb angle after leaving ground."},
		{"takeoff", 2, "Qu est ce que le point de decision ?", "What is the decision point?", `["Point ou on decide decoller ou arreter","Point rotation","Point decollage","Point montee"]`, "Point ou on decide decoller ou arreter", "Point decision avant V1 on peut arreter apres V1 on decolle.", "Decision point before V1 abort after V1 continue."},
		{"landing", 2, "Facteurs affectant distance atterrissage ?", "Factors affecting landing distance?", `["Poids vent pente etat piste","Couleur avion","Heure jour","Type helice"]`, "Poids vent pente etat piste", "Poids vent pente etat piste affectent distance atterrissage.", "Weight wind slope runway condition affect landing distance."},
		{"landing", 3, "Qu est ce que la distance d arret ?", "What is stopping distance?", `["Distance entre toucher et arret complet","Distance approche","Distance decollage","Distance roulage"]`, "Distance entre toucher et arret complet", "Distance arret du toucher des roues a l arret complet.", "Stopping distance from touchdown to full stop."},
		{"landing", 1, "Effet volets sur atterrissage ?", "Flaps effect on landing?", `["Reduisent vitesse approche","Augmentent vitesse","Pas effet","Augmentent distance"]`, "Reduisent vitesse approche", "Volets augmentent portance reduisent vitesse approche.", "Flaps increase lift reduce approach speed."},
		{"landing", 2, "Qu est ce que l arrondi flare ?", "What is flare?", `["Redressement avion avant toucher","Virage final","Descente rapide","Montee"]`, "Redressement avion avant toucher", "Arrondi action redresser avion juste avant toucher piste.", "Flare action to level aircraft just before runway contact."},
		{"climb", 1, "Qu est ce que le taux de montee ?", "What is rate of climb?", `["Vitesse verticale ft/min","Vitesse horizontale","Pente","Distance"]`, "Vitesse verticale ft/min", "Taux montee vitesse verticale en pieds par minute.", "Rate of climb vertical speed in feet per minute."},
		{"climb", 2, "Qu est ce que la pente de montee ?", "What is climb gradient?", `["Angle montee en %","Vitesse montee","Distance montee","Temps montee"]`, "Angle montee en %", "Pente montee rapport hauteur sur distance horizontale.", "Climb gradient ratio height over horizontal distance."},
		{"climb", 3, "Qu est ce que le plafond pratique ?", "What is service ceiling?", `["Altitude taux montee 100 ft/min","Altitude max","Altitude croisiere","Altitude decollage"]`, "Altitude taux montee 100 ft/min", "Plafond pratique altitude ou taux montee tombe a 100 ft/min.", "Service ceiling altitude where climb rate drops to 100 ft/min."},
		{"climb", 1, "Qu est ce que le plafond absolu ?", "What is absolute ceiling?", `["Altitude max atteignable","Altitude croisiere","Altitude decollage","Altitude securite"]`, "Altitude max atteignable", "Plafond absolu altitude maximale que avion peut atteindre.", "Absolute ceiling maximum altitude aircraft can reach."},
		{"cruise", 2, "Qu est ce que la puissance de croisiere ?", "What is cruise power?", `["Puissance optimale economique","Puissance max","Puissance decollage","Puissance minimale"]`, "Puissance optimale economique", "Puissance croisiere equilibre vitesse et consommation.", "Cruise power balance between speed and fuel consumption."},
		{"cruise", 1, "Qu est ce que le regime de croisiere ?", "What is cruise regime?", `["Vol stabilise vitesse altitude constantes","Montee","Descente","Decollage"]`, "Vol stabilise vitesse altitude constantes", "Croisiere vol en palier a vitesse et altitude constantes.", "Cruise level flight at constant speed and altitude."},
		{"cruise", 3, "Qu est ce que la distance franchissable ?", "What is range?", `["Distance max avec plein carburant","Distance vol","Distance decollage","Distance atterrissage"]`, "Distance max avec plein carburant", "Distance franchissable distance maximale parcourable.", "Range maximum distance flyable."},
		{"cruise", 2, "Qu est ce que l autonomie ?", "What is endurance?", `["Temps max vol avec plein carburant","Distance max","Vitesse max","Altitude max"]`, "Temps max vol avec plein carburant", "Autonomie duree maximale de vol.", "Endurance maximum flight duration."},
		{"crosswind", 2, "Qu est ce que la composante vent traversier ?", "What is crosswind component?", `["Force vent perpendiculaire piste","Vent face","Vent arriere","Vent nul"]`, "Force vent perpendiculaire piste", "Composante traversiere force vent de cote par rapport piste.", "Crosswind component wind force perpendicular to runway."},
		{"crosswind", 1, "Limite vent traversier typique avion leger ?", "Typical crosswind limit light aircraft?", `["15-20 kt","30 kt","5 kt","40 kt"]`, "15-20 kt", "Limite vent traversier avion leger 15-20 kt.", "Crosswind limit light aircraft 15-20 kt."},
		{"crosswind", 3, "Comment atterrir avec vent traversier ?", "How to land with crosswind?", `["Technique crabe ou aile basse","Approche normale","Atterrir droit","Plein gaz"]`, "Technique crabe ou aile basse", "Crabe ou aile basse pour compenser vent traversier.", "Crab or wing-down technique to compensate crosswind."},
		{"weight", 2, "Effet poids sur performances ?", "Weight effect on performance?", `["Poids eleve reduit performances","Poids eleve ameliore","Pas effet","Negligeable"]`, "Poids eleve reduit performances", "Poids eleve augmente distance decollage reduit montee.", "Higher weight increases takeoff distance reduces climb."},
		{"weight", 1, "Qu est ce que la charge alaire ?", "What is wing loading?", `["Poids / Surface alaire","Poids total","Surface aile","Envergure"]`, "Poids / Surface alaire", "Charge alaire rapport poids total sur surface aile.", "Wing loading total weight divided by wing area."},
		{"weight", 3, "Effet charge alaire sur decrochage ?", "Wing loading effect on stall?", `["Charge elevee augmente vitesse decrochage","Charge elevee reduit","Pas effet","Negligeable"]`, "Charge elevee augmente vitesse decrochage", "Charge alaire elevee vitesse decrochage plus haute.", "High wing loading higher stall speed."},
	}, count)
}

func genAircraftGeneral(count int) []QuestionData {
	return repeatQuestions([]QuestionData{
		{"engine", 1, "Quatre temps d un moteur a piston ?", "Four strokes of a piston engine?", `["Admission compression combustion echappement","Injection compression explosion sortie","Aspiration pression travail rejet","Entree feu sortie arret"]`, "Admission compression combustion echappement", "Cycle 4 temps admission compression combustion echappement.", "4-stroke cycle intake compression power exhaust."},
		{"engine", 2, "Difference melange riche et pauvre ?", "Rich vs lean mixture?", `["Riche plus carburant Pauvre moins","Riche moins carburant","Automatique","Non reglable"]`, "Riche plus carburant Pauvre moins", "Melange riche refroidit moteur melange pauvre economise carburant.", "Rich mixture cools engine lean mixture saves fuel."},
		{"engine", 1, "A quoi sert le starter choke ?", "What is the choke for?", `["Enrichir melange demarrage froid","Appauvrir melange","Augmenter regime","Refroidir moteur"]`, "Enrichir melange demarrage froid", "Starter enrichit melange air-carburant pour demarrage froid.", "Choke enriches air-fuel mixture for cold starting."},
		{"engine", 3, "Qu est ce que le detonnement ?", "What is engine detonation?", `["Combustion anormale spontanee melange","Nettoyage moteur","Revision complete","Changement huile"]`, "Combustion anormale spontanee melange", "Detonnement combustion spontanee qui peut endommager le moteur.", "Detonation spontaneous combustion that can damage engine."},
		{"engine", 2, "Qu est ce que le preallumage ?", "What is pre-ignition?", `["Allumage avant etincelle normale","Allumage apres","Pas allumage","Double allumage"]`, "Allumage avant etincelle normale", "Preallumage melange s enflamme avant etincelle bougie.", "Pre-ignition mixture ignites before spark plug."},
		{"engine", 1, "Role des magnetos ?", "Role of magnetos?", `["Generer etincelle bougies","Charger batterie","Alimenter pompe","Refroidir moteur"]`, "Generer etincelle bougies", "Magnetos generent courant haute tension pour bougies.", "Magnetos generate high voltage for spark plugs."},
		{"engine", 3, "Pourquoi deux magnetos ?", "Why two magnetos?", `["Redondance securite meilleure combustion","Puissance double","Vitesse double","Refroidissement"]`, "Redondance securite meilleure combustion", "Deux magnetos securite et combustion plus complete.", "Two magnetos safety and more complete combustion."},
		{"engine", 2, "Essai magnetos avant decollage ?", "Magnetos check before takeoff?", `["Verifier chute regime chaque magneto","Verifier puissance","Verifier vitesse","Verifier temperature"]`, "Verifier chute regime chaque magneto", "Essai magnetos verifie chute regime acceptable chaque magneto.", "Magnetos check verifies acceptable RPM drop each magneto."},
		{"engine", 1, "Qu est ce que le regime ralenti ?", "What is idle speed?", `["Vitesse minimale moteur marche","Vitesse max","Vitesse decollage","Vitesse croisiere"]`, "Vitesse minimale moteur marche", "Ralenti regime minimal moteur sans caler.", "Idle minimum engine speed without stalling."},
		{"engine", 2, "Qu est ce que la puissance max continue ?", "What is max continuous power?", `["Puissance max utilisable indefiniment","Puissance decollage","Puissance minimale","Puissance croisiere"]`, "Puissance max utilisable indefiniment", "Puissance max continue sans endommager moteur.", "Max continuous power without engine damage."},
		{"electrical", 2, "Role de l alternateur dans un avion ?", "Role of alternator in aircraft?", `["Charger batterie alimenter systemes","Demarrer moteur","Allumer bougies","Refroidir moteur"]`, "Charger batterie alimenter systemes", "Alternateur fournit courant electrique pour batterie et systemes.", "Alternator provides electrical current for battery and systems."},
		{"electrical", 1, "Tension electrique typique avion leger ?", "Typical voltage light aircraft?", `["12V ou 24V","110V","220V","6V"]`, "12V ou 24V", "Avions legers utilisent generalement du 12V ou 24V.", "Light aircraft typically use 12V or 24V."},
		{"electrical", 3, "Panne alternateur que faire ?", "Alternator failure what to do?", `["Reduire charge electrique atterrir","Continuer vol","Augmenter charge","Ignorer"]`, "Reduire charge electrique atterrir", "Panne alternateur reduire consommation atterrir rapidement.", "Alternator failure reduce consumption land soon."},
		{"electrical", 2, "Role de la batterie avion ?", "Role of aircraft battery?", `["Demarrer moteur reserve secours","Alimenter toujours","Charger alternateur","Refroidir"]`, "Demarrer moteur reserve secours", "Batterie demarre moteur et sert de reserve.", "Battery starts engine and serves as backup."},
		{"electrical", 1, "Que signifie master switch ?", "What does master switch mean?", `["Interrupteur general electrique","Interrupteur moteur","Interrupteur pompe","Interrupteur radio"]`, "Interrupteur general electrique", "Master switch coupe ou met sous tension tout le systeme electrique.", "Master switch turns entire electrical system on/off."},
		{"fuel", 2, "Carburant moteurs a piston aviation ?", "Fuel for aviation piston engines?", `["AvGas essence aviation","Jet A1 kerosene","Kerosene","Diesel"]`, "AvGas essence aviation", "AvGas essence specialement formulee pour l aviation a piston.", "AvGas gasoline specially formulated for piston aviation."},
		{"fuel", 1, "Couleur de l AvGas 100LL ?", "Color of AvGas 100LL?", `["Bleue","Verte","Rouge","Incolore"]`, "Bleue", "AvGas 100LL est de couleur bleue pour eviter confusions.", "AvGas 100LL is blue to prevent confusion."},
		{"fuel", 3, "Qu est ce que le Jet A1 ?", "What is Jet A1?", `["Kerosene pour turbines","Essence pour pistons","Diesel","Biocarburant"]`, "Kerosene pour turbines", "Jet A1 kerosene utilise par moteurs a reaction.", "Jet A1 kerosene used by jet engines."},
		{"fuel", 2, "Couleur du Jet A1 ?", "Color of Jet A1?", `["Incolore ou jaune pale","Bleu","Vert","Rouge"]`, "Incolore ou jaune pale", "Jet A1 est incolore ou jaune tres pale.", "Jet A1 is colorless or very pale yellow."},
		{"fuel", 1, "Pourquoi vidanger reservoirs avant vol ?", "Why drain fuel tanks before flight?", `["Verifier absence eau impuretes","Vider reservoir","Remplir","Nettoyer"]`, "Verifier absence eau impuretes", "Vidanger verifie absence eau et contaminants dans carburant.", "Drain checks for water and contaminants in fuel."},
		{"fuel", 2, "Qu est ce que le carburant mogas ?", "What is mogas?", `["Essence automobile pour certains avions","Kerosene","Diesel","AvGas"]`, "Essence automobile pour certains avions", "Mogas essence auto utilisable sur certains avions certifies.", "Mogas car gasoline usable on some certified aircraft."},
		{"fuel", 3, "Risque utilisation mauvaise essence ?", "Risk of wrong fuel type?", `["Detonnement panne moteur","Rien","Meilleure perf","Moins consommation"]`, "Detonnement panne moteur", "Mauvaise essence peut causer detonnement et panne moteur.", "Wrong fuel can cause detonation and engine failure."},
		{"oil", 2, "Pourquoi verifier niveau huile avant chaque vol ?", "Why check oil before each flight?", `["Verifier niveau detecter fuites","Changer huile","Remplir complet","Nettoyer moteur"]`, "Verifier niveau detecter fuites", "Verifier niveau huile et absence fuites avant chaque vol.", "Check oil level and absence of leaks before each flight."},
		{"oil", 1, "Role de l huile moteur ?", "Role of engine oil?", `["Lubrifier refroidir nettoyer proteger","Chauffer","Refroidir seulement","Nettoyer seulement"]`, "Lubrifier refroidir nettoyer proteger", "Huile lubrifie refroidit nettoie et protege moteur.", "Oil lubricates cools cleans and protects engine."},
		{"oil", 3, "Pression huile basse que faire ?", "Low oil pressure what to do?", `["Atterrir rapidement","Continuer","Augmenter regime","Ignorer"]`, "Atterrir rapidement", "Pression huile basse atterrir rapidement pour eviter casse moteur.", "Low oil pressure land quickly to avoid engine damage."},
		{"oil", 2, "Temperature huile elevee que faire ?", "High oil temperature what to do?", `["Reduire puissance augmenter vitesse","Augmenter puissance","Ignorer","Atterrir"]`, "Reduire puissance augmenter vitesse", "Temperature huile haute reduire puissance augmenter refroidissement.", "High oil temp reduce power increase cooling."},
		{"propeller", 1, "Qu est ce que le pas d une helice ?", "What is propeller pitch?", `["Angle pales helice","Diametre helice","Vitesse rotation","Nombre pales"]`, "Angle pales helice", "Pas angle des pales qui determine l efficacite de l helice.", "Pitch blade angle that determines propeller efficiency."},
		{"propeller", 2, "Helice pas fixe vs pas variable ?", "Fixed vs variable pitch propeller?", `["Pas fixe angle constant pas variable reglable","Pas fixe reglable","Identique","Pas variable fixe"]`, "Pas fixe angle constant pas variable reglable", "Pas fixe angle constant pas variable angle ajustable en vol.", "Fixed pitch constant angle variable pitch adjustable in flight."},
		{"propeller", 3, "Qu est ce que le moulinet ?", "What is windmilling?", `["Helice tourne sous effet vent","Helice arretee","Helice vitesse max","Helice pas nul"]`, "Helice tourne sous effet vent", "Moulinet helice tourne entrainee par air moteur coupe.", "Windmilling propeller turns driven by air engine off."},
		{"propeller", 1, "Danger helice en rotation ?", "Danger of rotating propeller?", `["Tres dangereux respecter zone securite","Aucun","Faible","Mineur"]`, "Tres dangereux respecter zone securite", "Helice tournante extremement dangereuse zone de securite obligatoire.", "Rotating propeller extremely dangerous safety zone mandatory."},
		{"propeller", 2, "Qu est ce que l helice a vitesse constante ?", "What is constant speed propeller?", `["Helice maintient regime constant","Helice pas fixe","Helice arretee","Helice lente"]`, "Helice maintient regime constant", "Helice vitesse constante maintient RPM choisi par pilote.", "Constant speed propeller maintains selected RPM."},
		{"airframe", 2, "Materiaux construction aeronautique ?", "Aircraft construction materials?", `["Aluminium alliages composites acier titane","Bois seulement","Plastique","Fer"]`, "Aluminium alliages composites acier titane", "Avions modernes aluminium composites acier titane.", "Modern aircraft aluminum composites steel titanium."},
		{"airframe", 1, "Qu est ce que la cellule ?", "What is the airframe?", `["Structure complete avion sans moteur","Moteur","Helice","Train"]`, "Structure complete avion sans moteur", "Cellule fuselage ailes empennage train.", "Airframe fuselage wings empennage landing gear."},
		{"airframe", 3, "Qu est ce que la fatigue structurale ?", "What is structural fatigue?", `["Degradation progressive sous contraintes repetees","Rupture brutale","Corrosion","Usure normale"]`, "Degradation progressive sous contraintes repetees", "Fatigue structurale due aux cycles de contraintes repetes.", "Structural fatigue from repeated stress cycles."},
		{"airframe", 2, "Qu est ce que la corrosion sur avion ?", "What is aircraft corrosion?", `["Degradation metal par environnement","Rouille normale","Peinture","Salete"]`, "Degradation metal par environnement", "Corrosion attaque chimique des metaux par environnement.", "Corrosion chemical attack on metals from environment."},
		{"systems", 2, "Qu est ce que le systeme de chauffage cabine ?", "What is cabin heating system?", `["Chauffage par air echappement","Chauffage electrique","Chauffage gaz","Chauffage solaire"]`, "Chauffage par air echappement", "Chauffage cabine utilise air chaud echappement moteur.", "Cabin heating uses hot air from engine exhaust."},
		{"systems", 1, "Qu est ce que la ventilation cabine ?", "What is cabin ventilation?", `["Air frais exterieur pour cabine","Air conditionne","Chauffage","Climatisation"]`, "Air frais exterieur pour cabine", "Ventilation apporte air frais exterieur dans cabine.", "Ventilation brings fresh outside air into cabin."},
		{"systems", 3, "Qu est ce que le systeme d oxygene ?", "What is oxygen system?", `["Fournit oxygene haute altitude","Chauffage","Refroidissement","Pressurisation"]`, "Fournit oxygene haute altitude", "Systeme oxygene necessaire au-dessus 10000 ft.", "Oxygen system needed above 10000 ft."},
		{"systems", 2, "Altitude ou oxygene obligatoire ?", "Altitude where oxygen mandatory?", `["Au-dessus 10000 ft","5000 ft","15000 ft","20000 ft"]`, "Au-dessus 10000 ft", "Oxygene obligatoire au-dessus 10000 ft.", "Oxygen mandatory above 10000 ft."},
		{"maintenance", 2, "Frequence inspections obligatoires ?", "Mandatory inspection frequency?", `["Annuelle ou 100 heures","Mensuelle","Hebdomadaire","Tous les 5 ans"]`, "Annuelle ou 100 heures", "Inspection annuelle ou 100 heures de vol.", "Annual or 100 flight hours inspection."},
		{"maintenance", 1, "Qu est ce que la visite pre-vol ?", "What is pre-flight inspection?", `["Inspection visuelle avant chaque vol","Inspection annuelle","Revision moteur","Changement pneu"]`, "Inspection visuelle avant chaque vol", "Visite pre-vol inspection visuelle obligatoire avant chaque vol.", "Pre-flight visual inspection mandatory before each flight."},
		{"maintenance", 3, "Qu est ce que le carnet de route ?", "What is the aircraft logbook?", `["Document enregistre tous vols et maintenance","Plan de vol","Carte","Manuel"]`, "Document enregistre tous vols et maintenance", "Carnet de route historique complet de l aeronef.", "Logbook complete aircraft history."},
	}, count)
}

func genFlightPlanning(count int) []QuestionData {
	return repeatQuestions([]QuestionData{
		{"notam", 1, "Qu est ce qu un NOTAM ?", "What is a NOTAM?", `["Avis navigateurs aeriens","Prevision meteo","Plan vol","Carte aeronautique"]`, "Avis navigateurs aeriens", "NOTAM Notice to Air Missions avis important securite vols.", "NOTAM Notice to Air Missions important for flight safety."},
		{"notam", 2, "Ou consulter les NOTAM avant vol ?", "Where to consult NOTAMs before flight?", `["Services information aeronautique","Meteo France","Tour controle","Radio"]`, "Services information aeronautique", "NOTAM disponibles via services information aeronautique AIS.", "NOTAMs available via Aeronautical Information Services AIS."},
		{"notam", 3, "Types de NOTAM ?", "Types of NOTAM?", `["NOTAM N R et C","NOTAM A B C","NOTAM 1 2 3","NOTAM X Y Z"]`, "NOTAM N R et C", "NOTAM N nouveau R remplacant C annulation.", "NOTAM N new R replacing C cancellation."},
		{"notam", 2, "Validite d un NOTAM ?", "NOTAM validity?", `["Variable selon nature","24h","7 jours","1 mois"]`, "Variable selon nature", "Validite NOTAM depend de la nature de l information.", "NOTAM validity depends on information nature."},
		{"notam", 1, "Que signifie NOTAMN ?", "What does NOTAMN mean?", `["Nouveau NOTAM","NOTAM annule","NOTAM remplace","NOTAM urgent"]`, "Nouveau NOTAM", "NOTAMN nouveau NOTAM.", "NOTAMN new NOTAM."},
		{"notam", 2, "Que signifie NOTAMR ?", "What does NOTAMR mean?", `["NOTAM remplacant","Nouveau NOTAM","NOTAM annule","NOTAM urgent"]`, "NOTAM remplacant", "NOTAMR NOTAM qui en remplace un autre.", "NOTAMR NOTAM replacing another."},
		{"notam", 3, "Que signifie NOTAMC ?", "What does NOTAMC mean?", `["NOTAM annulation","Nouveau NOTAM","NOTAM remplace","NOTAM urgent"]`, "NOTAM annulation", "NOTAMC NOTAM annulant un NOTAM precedent.", "NOTAMC NOTAM canceling a previous NOTAM."},
		{"fuel", 2, "Calcul carburant necessaire pour vol ?", "Required fuel calculation?", `["Route + reserve + degagement","Route seulement","Reserve seulement","Hasard"]`, "Route + reserve + degagement", "Carburant total = route + reserve + degagement.", "Total fuel = route + reserve + alternate."},
		{"fuel", 1, "Reserve carburant minimale VFR jour ?", "Minimum fuel reserve day VFR?", `["30 minutes","15 minutes","1 heure","Pas reserve"]`, "30 minutes", "30 minutes reserve carburant pour VFR de jour.", "30 minutes fuel reserve for day VFR."},
		{"fuel", 2, "Reserve minimale VFR nuit ?", "Minimum reserve night VFR?", `["45 minutes","30 minutes","1 heure","20 minutes"]`, "45 minutes", "45 minutes reserve pour VFR de nuit.", "45 minutes reserve for night VFR."},
		{"fuel", 3, "Reserve carburant IFR ?", "IFR fuel reserve?", `["Carburant degagement + 30 min","45 min","1 heure","15 min"]`, "Carburant degagement + 30 min", "IFR reserve = carburant degagement + 30 minutes.", "IFR reserve = alternate fuel + 30 minutes."},
		{"fuel", 2, "Qu est ce que le carburant de degagement ?", "What is alternate fuel?", `["Carburant pour aller aerodrome degagement","Carburant route","Reserve","Carburant taxi"]`, "Carburant pour aller aerodrome degagement", "Carburant degagement pour rejoindre aerodrome alternative.", "Alternate fuel to reach alternate aerodrome."},
		{"fuel", 1, "Qu est ce que le carburant de route ?", "What is trip fuel?", `["Carburant de destination a destination","Carburant reserve","Carburant taxi","Carburant degagement"]`, "Carburant de destination a destination", "Carburant route du depart a la destination.", "Trip fuel from departure to destination."},
		{"fuel", 3, "Qu est ce que le carburant taxi ?", "What is taxi fuel?", `["Carburant consomme avant decollage","Carburant route","Reserve","Degagement"]`, "Carburant consomme avant decollage", "Carburant taxi consomme roulage et essais moteur.", "Taxi fuel consumed during taxi and run-up."},
		{"fuel", 2, "Qu est ce que le carburant de reserve final ?", "What is final reserve fuel?", `["Carburant minimum apres arrivée","Carburant route","Carburant taxi","Carburant degagement"]`, "Carburant minimum apres arrivée", "Reserve finale carburant minimum apres atterrissage.", "Final reserve minimum fuel after landing."},
		{"route", 2, "Point de report obligatoire ?", "Mandatory reporting point?", `["Point ou pilote contacte controle","Point navigation","Aeroport","Balise"]`, "Point ou pilote contacte controle", "Pilote doit signaler sa position a ces points.", "Pilot must report position at these points."},
		{"route", 1, "A quoi sert un plan de vol ?", "What is a flight plan for?", `["Informer services circulation aerienne","Naviguer","Calculer carburant","Choisir route"]`, "Informer services circulation aerienne", "Plan de vol informe les ATS de votre itineraire.", "Flight plan informs ATS of your route."},
		{"route", 3, "Quand deposer un plan de vol ?", "When to file a flight plan?", `["60 min avant depart","Decollage","Veille","Apres vol"]`, "60 min avant depart", "Plan de vol doit etre depose 60 min avant depart.", "Flight plan must be filed 60 min before departure."},
		{"route", 2, "Que faire si plan de vol non ferme ?", "What if flight plan not closed?", `["Declencher operations recherche","Rien","Attendre","Replanifier"]`, "Declencher operations recherche", "Plan vol non ferme declenche operations alerte recherche.", "Unclosed flight plan triggers alert search operations."},
		{"route", 1, "Comment fermer un plan de vol ?", "How to close a flight plan?", `["Contacter ATS apres atterrissage","Automatique","Radio","Email"]`, "Contacter ATS apres atterrissage", "Fermer plan de vol aupres ATS apres atterrissage.", "Close flight plan with ATS after landing."},
		{"route", 3, "Modification plan de vol en vol ?", "Flight plan modification in flight?", `["Contacter ATS pour modification","Impossible","Automatique","Radio"]`, "Contacter ATS pour modification", "Modifier plan de vol via ATS en vol.", "Modify flight plan via ATC in flight."},
		{"maps", 1, "Qu est ce qu une carte OACI ?", "What is an ICAO chart?", `["Carte aeronautique standardisee","Carte routiere","Carte meteo","Carte topographique"]`, "Carte aeronautique standardisee", "Cartes OACI cartes aeronautiques officielles.", "ICAO charts are official aeronautical charts."},
		{"maps", 2, "Infos sur carte aeronautique ?", "Info on aeronautical chart?", `["Espaces aeriens balises aerodromes","Routes","Villes","Relief"]`, "Espaces aeriens balises aerodromes", "Cartes montrent espaces aeriens balises et aerodromes.", "Charts show airspace beacons and aerodromes."},
		{"maps", 3, "Qu est ce qu une carte d approche ?", "What is an approach chart?", `["Carte detaillee procedure atterrissage","Carte en route","Carte aerodrome","Carte meteo"]`, "Carte detaillee procedure atterrissage", "Carte approche montre procedure atterrissage detaillee.", "Approach chart shows detailed landing procedure."},
		{"maps", 1, "Qu est ce qu une carte d aerodrome ?", "What is an aerodrome chart?", `["Plan detaille aerodrome pistes voies","Carte region","Carte meteo","Carte approche"]`, "Plan detaille aerodrome pistes voies", "Carte aerodrome plan detaille de l aeroport.", "Aerodrome chart detailed airport plan."},
		{"maps", 2, "Symboles sur carte aeronautique ?", "Symbols on aeronautical chart?", `["Standardises OACI","Variables","Optionnels","Locaux"]`, "Standardises OACI", "Symboles cartes aeronautiques standardises par OACI.", "Chart symbols standardized by ICAO."},
		{"weather_minima", 2, "Minima VFR en espace aerien controle ?", "VFR minima in controlled airspace?", `["Visibilite 5 km hors nuages","1500 m","8 km","10 km"]`, "Visibilite 5 km hors nuages", "VFR en espace controle visibilite 5 km hors nuages.", "VFR in controlled airspace 5 km visibility clear of clouds."},
		{"weather_minima", 1, "Minima VFR en espace non controle ?", "VFR minima in uncontrolled airspace?", `["Visibilite 1500 m hors nuages","5 km","8 km","10 km"]`, "Visibilite 1500 m hors nuages", "VFR hors espace controle visibilite 1500 m hors nuages.", "VFR outside controlled airspace 1500 m clear of clouds."},
		{"weather_minima", 3, "Minima VFR special ?", "Special VFR minima?", `["Visibilite 1500 m hors nuages","5 km","8 km","500 m"]`, "Visibilite 1500 m hors nuages", "VFR special visibilite 1500 m hors nuages.", "Special VFR 1500 m visibility clear of clouds."},
		{"alternate", 2, "Quand aerodrome degagement obligatoire ?", "When alternate aerodrome mandatory?", `["Si meteo destination incertaine","Toujours","Jamais","Vol local"]`, "Si meteo destination incertaine", "Aerodrome degagement obligatoire si meteo incertaine.", "Alternate mandatory if destination weather uncertain."},
		{"alternate", 1, "Distance maximale aerodrome degagement ?", "Max distance to alternate?", `["Pas limite specifique","100 NM","200 NM","50 NM"]`, "Pas limite specifique", "Pas de limite specifique mais doit etre atteignable.", "No specific limit but must be reachable."},
		{"alternate", 3, "Minima pour aerodrome degagement ?", "Minima for alternate aerodrome?", `["Au-dessus minima approche","Minima exacts","Sous minima","Pas minima"]`, "Au-dessus minima approche", "Aerodrome degagement doit avoir meteo au-dessus minima.", "Alternate must have weather above minima."},
		{"documentation", 2, "Documents obligatoires a bord ?", "Mandatory onboard documents?", `["Certificat immatriculation licence radio manuel vol","Passeport","Carte credit","Telephone"]`, "Certificat immatriculation licence radio manuel vol", "Certificat immatriculation licence station radio manuel vol.", "Registration certificate radio license flight manual."},
		{"documentation", 1, "Validite certificat immatriculation ?", "Registration certificate validity?", `["Tant que proprietaire","1 an","5 ans","10 ans"]`, "Tant que proprietaire", "Certificat immatriculation valide tant que proprietaire.", "Registration valid as long as owner."},
		{"documentation", 3, "Que faire si documents perdus ?", "What if documents lost?", `["Signaler autorites aviation","Rien","Reimprimer","Ignorer"]`, "Signaler autorites aviation", "Documents perdus signaler aux autorites aviation.", "Lost documents report to aviation authorities."},
	}, count)
}

func genHumanPerformance(count int) []QuestionData {
	return repeatQuestions([]QuestionData{
		{"vision", 1, "Temps adaptation obscurite ?", "Dark adaptation time?", `["30 minutes","5 minutes","1 heure","10 minutes"]`, "30 minutes", "Adaptation a l obscurite prend environ 30 minutes.", "Dark adaptation takes about 30 minutes."},
		{"vision", 2, "Qu est ce que la vision tunnel ?", "What is tunnel vision?", `["Reduction champ visuel sous stress","Vision nocturne","Vision peripherique","Cecite"]`, "Reduction champ visuel sous stress", "Stress reduit le champ visuel peripherique.", "Stress reduces peripheral visual field."},
		{"vision", 3, "Qu est ce que la vision peripherique ?", "What is peripheral vision?", `["Vision cote champ visuel","Vision centrale","Vision couleur","Vision nuit"]`, "Vision cote champ visuel", "Vision peripherique detecte mouvements cote.", "Peripheral vision detects side movements."},
		{"vision", 1, "Qu est ce que la vision nocturne ?", "What is night vision?", `["Vision utilisant batonnets retine","Vision cones","Vision peripherique","Vision centrale"]`, "Vision utilisant batonnets retine", "Vision nocturne utilise batonnets sensibles lumiere faible.", "Night vision uses rods sensitive to low light."},
		{"vision", 2, "Comment proteger vision nocturne ?", "How to protect night vision?", `["Lumiere rouge cabine","Lumiere blanche","Lumiere bleue","Pas lumiere"]`, "Lumiere rouge cabine", "Lumiere rouge preserve adaptation obscurite.", "Red light preserves dark adaptation."},
		{"vision", 1, "Qu est ce que la vision centrale ?", "What is central vision?", `["Vision utilisant cones retine","Vision peripherique","Vision nocturne","Vision tunnel"]`, "Vision utilisant cones retine", "Vision centrale utilise cones pour details et couleurs.", "Central vision uses cones for details and colors."},
		{"vision", 3, "Qu est ce que la myopie du vide ?", "What is empty field myopia?", `["Myopie temporaire sans point fixation","Myopie permanente","Presbytie","Astigmatisme"]`, "Myopie temporaire sans point fixation", "Myopie du vide sans point fixation dans ciel.", "Empty field myopia without fixation point in sky."},
		{"hypoxia", 2, "Symptomes de l hypoxie ?", "Hypoxia symptoms?", `["Euphorie troubles jugement fatigue","Maux tete","Fatigue","Aucun"]`, "Euphorie troubles jugement fatigue", "Hypoxie cause euphorie troubles jugement maux tete fatigue.", "Hypoxia causes euphoria judgment issues headache fatigue."},
		{"hypoxia", 1, "Altitude ou hypoxie peut survenir ?", "Altitude where hypoxia can occur?", `["3000 m","1500 m","5000 m","8000 m"]`, "3000 m", "Hypoxie peut survenir des 3000 m d altitude.", "Hypoxia can occur from 3000 m."},
		{"hypoxia", 3, "Types d hypoxie ?", "Types of hypoxia?", `["Hypoxique anemique stagnante histotoxique","Hypoxique seulement","Anemique seulement","Stagnante seulement"]`, "Hypoxique anemique stagnante histotoxique", "4 types hypoxie hypoxique anemique stagnante histotoxique.", "4 types hypoxic anemic stagnant histotoxic."},
		{"hypoxia", 2, "Qu est ce que l hypoxie hypoxique ?", "What is hypoxic hypoxia?", `["Manque oxygene air inspire","Manque globules rouges","Circulation insuffisante","Cellules bloquent oxygene"]`, "Manque oxygene air inspire", "Hypoxie hypoxique faible pression oxygene haute altitude.", "Hypoxic hypoxia low oxygen pressure high altitude."},
		{"hypoxia", 1, "Qu est ce que l hypoxie anemique ?", "What is anemic hypoxia?", `["Manque globules rouges oxygene","Manque oxygene air","Circulation insuffisante","Cellules bloquent"]`, "Manque globules rouges oxygene", "Hypoxie anemique sang ne peut transporter assez oxygene.", "Anemic hypoxia blood cannot carry enough oxygen."},
		{"hypoxia", 3, "Qu est ce que l hypoxie stagnante ?", "What is stagnant hypoxia?", `["Circulation sanguine insuffisante","Manque oxygene air","Manque globules rouges","Cellules bloquent"]`, "Circulation sanguine insuffisante", "Hypoxie stagnante sang circule trop lentement.", "Stagnant hypoxia blood circulates too slowly."},
		{"hypoxia", 2, "Qu est ce que l hypoxie histotoxique ?", "What is histotoxic hypoxia?", `["Cellules ne peuvent utiliser oxygene","Manque oxygene air","Manque globules","Circulation insuffisante"]`, "Cellules ne peuvent utiliser oxygene", "Hypoxie histotoxique cellules empechees utiliser oxygene.", "Histotoxic hypoxia cells prevented from using oxygen."},
		{"hyperventilation", 2, "Symptomes hyperventilation ?", "Hyperventilation symptoms?", `["Vertiges fourmillements crampes","Euphorie","Fatigue","Maux tete"]`, "Vertiges fourmillements crampes", "Hyperventilation vertiges fourmillements doigts crampes.", "Hyperventilation dizziness tingling fingers cramps."},
		{"hyperventilation", 1, "Cause hyperventilation ?", "Hyperventilation cause?", `["Respiration trop rapide profonde","Respiration lente","Apnee","Toux"]`, "Respiration trop rapide profonde", "Hyperventilation respiration excessive elimine trop CO2.", "Hyperventilation excessive breathing removes too much CO2."},
		{"hyperventilation", 3, "Correction hyperventilation ?", "Hyperventilation correction?", `["Ralentir respiration sac papier","Augmenter respiration","Boire","Manger"]`, "Ralentir respiration sac papier", "Ralentir respiration ou respirer dans sac.", "Slow breathing or breathe into bag."},
		{"barotrauma", 2, "Qu est ce que le barotraumatisme ?", "What is barotrauma?", `["Lesion due difference pression","Traumatisme choc","Brulure","Fracture"]`, "Lesion due difference pression", "Barotraumatisme douleur oreilles sinus variation pression.", "Barotrauma ear sinus pain from pressure change."},
		{"barotrauma", 1, "Comment eviter barotraumatisme oreilles ?", "How to avoid ear barotrauma?", `["Deglutition manoeuvre Valsalva","Baillement","Toux","Eternuement"]`, "Deglutition manoeuvre Valsalva", "Deglutir ou Valsalva equilibre pression oreilles.", "Swallow or Valsalva equalize ear pressure."},
		{"barotrauma", 3, "Quand barotraumatisme plus probable ?", "When barotrauma most likely?", `["Descente rapide","Montee rapide","Vol stabilise","Sol"]`, "Descente rapide", "Descente rapide pression exterieure augmente vite.", "Rapid descent outside pressure increases quickly."},
		{"fatigue", 2, "Qu est ce que la fatigue aeronautique ?", "What is aviation fatigue?", `["Baisse performances manque sommeil","Fatigue musculaire","Faim","Soif"]`, "Baisse performances manque sommeil", "Fatigue reduit performances cognitives et physiques.", "Fatigue reduces cognitive and physical performance."},
		{"fatigue", 1, "Comment prevenir la fatigue en vol ?", "How to prevent fatigue in flight?", `["Reposer avant vol hydrater pauses","Cafe","Musique","Exercice"]`, "Reposer avant vol hydrater pauses", "Bon repos avant vol et hydratation reguliere aident.", "Good rest before flight and regular hydration help."},
		{"fatigue", 3, "Qu est ce que la fatigue chronique ?", "What is chronic fatigue?", `["Fatigue persistante repos insuffisant","Fatigue passagere","Fatigue musculaire","Fatigue visuelle"]`, "Fatigue persistante repos insuffisant", "Fatigue chronique accumulation manque sommeil sur jours.", "Chronic fatigue accumulated sleep debt over days."},
		{"fatigue", 2, "Effets fatigue sur performances vol ?", "Fatigue effects on flight performance?", `["Erreurs jugement reflexes lents","Amelioration","Aucun","Concentration accrue"]`, "Erreurs jugement reflexes lents", "Fatigue cause erreurs jugement reflexes lents.", "Fatigue causes judgment errors slow reflexes."},
		{"stress", 2, "Qu est ce que l eustress ?", "What is eustress?", `["Stress qui ameliore performance","Stress negatif","Absence stress","Stress chronique"]`, "Stress qui ameliore performance", "Eustress stress modere qui ameliore les performances.", "Eustress moderate stress that improves performance."},
		{"stress", 3, "Effet stress sur prise decision ?", "Stress effect on decision-making?", `["Reduit options considerees","Ameliore decisions","Aucun effet","Ralentit reflexes"]`, "Reduit options considerees", "Stress reduit capacite a envisager toutes les options.", "Stress reduces ability to consider all options."},
		{"stress", 1, "Qu est ce que la distress ?", "What is distress?", `["Stress negatif excessif nuit performance","Stress positif","Absence stress","Stress modere"]`, "Stress negatif excessif nuit performance", "Distress stress excessif qui degrade performances.", "Distress excessive stress that degrades performance."},
		{"stress", 2, "Sources stress en aviation ?", "Sources of stress in aviation?", `["Meteo panne fatigue pression temps","Confort","Musique","Nourriture"]`, "Meteo panne fatigue pression temps", "Stress sources meteo pannes fatigue pression temps.", "Stress sources weather failures fatigue time pressure."},
		{"stress", 3, "Gestion du stress en vol ?", "Stress management in flight?", `["Prioriser communiquer deleguer","Ignorer","Accelerer","Forcer"]`, "Prioriser communiquer deleguer", "Gerer stress prioriser taches communiquer deleguer.", "Manage stress prioritize tasks communicate delegate."},
		{"illusions", 2, "Illusion somatogravique ?", "Somatogravic illusion?", `["Fausse sensation acceleration","Illusion visuelle","Illusion auditive","Hallucination"]`, "Fausse sensation acceleration", "Illusion due aux forces d acceleration qui trompent oreille interne.", "Illusion due to acceleration forces tricking inner ear."},
		{"illusions", 3, "Illusion de Coriolis ?", "Coriolis illusion?", `["Fausse sensation rotation","Illusion visuelle","Perte orientation","Vertige"]`, "Fausse sensation rotation", "Illusion causee par mouvement liquides oreille interne lors virages.", "Illusion caused by fluid movement in inner ear during turns."},
		{"illusions", 1, "Illusion d inclinaison ?", "Lean illusion?", `["Fausse sensation inclinaison","Fausse sensation vitesse","Fausse sensation altitude","Fausse sensation temperature"]`, "Fausse sensation inclinaison", "Illusion inclinaison due a forces centrifuges.", "Lean illusion from centrifugal forces."},
		{"illusions", 2, "Illusion de fausse piste ?", "False horizon illusion?", `["Confusion avec lumières sol","Illusion mouvement","Illusion altitude","Illusion vitesse"]`, "Confusion avec lumières sol", "Fausse piste confondre lumières sol avec horizon.", "False horizon confusing ground lights with horizon."},
		{"illusions", 3, "Illusion d approche en pente ?", "Sloping terrain illusion?", `["Piste monte ou descend donne fausse impression altitude","Fausse vitesse","Fausse direction","Fausse temperature"]`, "Piste monte ou descend donne fausse impression altitude", "Terrain en pente donne fausse impression hauteur approche.", "Sloping terrain gives false height impression on approach."},
		{"illusions", 1, "Qu est ce que la desorientation spatiale ?", "What is spatial disorientation?", `["Perte repere spatial causee conflit sensoriel","Perte altitude","Perte vitesse","Perte cap"]`, "Perte repere spatial causee conflit sensoriel", "Desorientation spatiale conflit entre sens et instruments.", "Spatial disorientation conflict between senses and instruments."},
		{"alcohol", 2, "Regle alcool avant vol ?", "Alcohol rule before flight?", `["8 heures bouteille au bouchon","12 heures","4 heures","24 heures"]`, "8 heures bouteille au bouchon", "8 heures entre dernier verre et vol regle FAA.", "8 hours bottle to throttle FAA rule."},
		{"alcohol", 1, "Taux alcool autorise pilote ?", "Allowed blood alcohol pilot?", `["0.2 g/L","0.5 g/L","0.8 g/L","0.0 g/L"]`, "0.2 g/L", "Taux alcool max 0.2 g/L pour pilotes.", "Max blood alcohol 0.2 g/L for pilots."},
		{"alcohol", 3, "Effets alcool sur pilotage ?", "Alcohol effects on flying?", `["Jugement reflexes coordination reduits","Amelioration","Aucun","Concentration accrue"]`, "Jugement reflexes coordination reduits", "Alcool reduit jugement reflexes et coordination.", "Alcohol reduces judgment reflexes and coordination."},
		{"medications", 2, "Medicaments autorises en vol ?", "Medications allowed in flight?", `["Seulement sans ordonnance autorises aviation","Tous","Aucun","Uniquement homeopathie"]`, "Seulement sans ordonnance autorises aviation", "Certains medicaments sans ordonnance autorises.", "Some over-the-counter medications allowed."},
		{"medications", 1, "Que faire avant prendre medicament ?", "What to do before taking medication?", `["Consulter medecin aeronautique","Prendre normal","Ignorer","Doubler dose"]`, "Consulter medecin aeronautique", "Toujours consulter medecin aeronautique avant medicament.", "Always consult aviation doctor before medication."},
		{"circadian", 2, "Qu est ce que le rythme circadien ?", "What is circadian rhythm?", `["Cycle biologique 24h","Cycle sommeil","Cycle alimentation","Cycle travail"]`, "Cycle biologique 24h", "Rythme circadien cycle eveil sommeil 24h.", "Circadian rhythm 24h wake-sleep cycle."},
		{"circadian", 3, "Effets decalage horaire jet lag ?", "Jet lag effects?", `["Fatigue troubles sommeil digestion","Euphorie","Force","Appetit"]`, "Fatigue troubles sommeil digestion", "Jet lag fatigue troubles sommeil et digestion.", "Jet lag fatigue sleep and digestion issues."},
		{"circadian", 1, "Comment minimiser jet lag ?", "How to minimize jet lag?", `["S adapter fuseau horaire avant vol","Dormir plus","Manger plus","Boire cafe"]`, "S adapter fuseau horaire avant vol", "S adapter progressivement au fuseau horaire destination.", "Gradually adapt to destination time zone."},
		{"decision_making", 2, "Modele DECIDE ?", "DECIDE model?", `["Detect Estimate Choose Identify Do Evaluate","Decider Executer Controler","Decouvrir Essayer Choisir","Dire Ecouter Choisir"]`, "Detect Estimate Choose Identify Do Evaluate", "DECIDE Detect Estimate Choose Identify Do Evaluate.", "DECIDE Detect Estimate Choose Identify Do Evaluate."},
		{"decision_making", 1, "Qu est ce que le FORDEC ?", "What is FORDEC?", `["Facts Options Risks Decision Execute Check","Forcer Decider","Former Decider","Faire Decider"]`, "Facts Options Risks Decision Execute Check", "FORDEC Facts Options Risks Decision Execute Check.", "FORDEC Facts Options Risks Decision Execute Check."},
		{"decision_making", 3, "Qu est ce que le biais de confirmation ?", "What is confirmation bias?", `["Chercher infos confirmant ses croyances","Ignorer infos","Tout accepter","Tout refuser"]`, "Chercher infos confirmant ses croyances", "Biais confirmation chercher infos qui confirment nos idees.", "Confirmation bias seeking info confirming our beliefs."},
		{"crm", 2, "Qu est ce que le CRM ?", "What is CRM?", `["Gestion ressources equipe","Controle ressources","Communication ressources","Coordination ressources"]`, "Gestion ressources equipe", "CRM Crew Resource Management gestion ressources equipe.", "CRM Crew Resource Management team resource management."},
		{"crm", 1, "Principes CRM ?", "CRM principles?", `["Communication leadership prise decision","Silence obeissance","Autorite seule","Independance"]`, "Communication leadership prise decision", "CRM communication leadership prise de decision.", "CRM communication leadership decision making."},
		{"crm", 3, "Qu est ce que le gradient d autorite ?", "What is authority gradient?", `["Difference hierarchique entre membres equipe","Autorite absolue","Pas hierarchie","Egalite"]`, "Difference hierarchique entre membres equipe", "Gradient autorite difference hierarchie dans cockpit.", "Authority gradient hierarchy difference in cockpit."},
		{"hazards", 2, "Attitudes dangereuses pilote ?", "Hazardous pilot attitudes?", `["Invulnerabilite macho impulsivite resignation","Prudence","Peur","Hesitation"]`, "Invulnerabilite macho impulsivite resignation", "5 attitudes dangereuses anti-autorite impulsivite invulnerabilite macho resignation.", "5 hazardous attitudes anti-authority impulsivity invulnerability macho resignation."},
		{"hazards", 1, "Qu est ce que l attitude anti-autorite ?", "What is anti-authority attitude?", `["N aime pas qu on lui dise quoi faire","Trop obeissant","Indifferent","Passif"]`, "N aime pas qu on lui dise quoi faire", "Anti-autorite n aime pas instructions.", "Anti-authority dislikes being told what to do."},
		{"hazards", 3, "Qu est ce que l attitude macho ?", "What is macho attitude?", `["Veut prouver capacite prendre risques","Tres prudent","Tres timide","Tres passif"]`, "Veut prouver capacite prendre risques", "Macho veut prouver prend risques inutiles.", "Macho wants to prove takes unnecessary risks."},
	}, count)
}

func genOperationalProcedures(count int) []QuestionData {
	return repeatQuestions([]QuestionData{
		{"departure", 1, "Que verifier avant demarrage moteur ?", "What to check before engine start?", `["Check-list complete","Rien","Juste carburant","Juste huile"]`, "Check-list complete", "Toujours utiliser la check-list avant demarrage.", "Always use the checklist before starting."},
		{"departure", 2, "Procedure demarrage moteur standard ?", "Standard engine start procedure?", `["Batterie ON mixture riche starter demarreur","Demarreur direct","Batterie seule","Mixture seule"]`, "Batterie ON mixture riche starter demarreur", "Batterie ON mixture riche starter si froid demarreur.", "Battery ON rich mixture choke if cold starter."},
		{"taxi", 1, "Comment diriger avion au sol ?", "How to steer aircraft on ground?", `["Palonnier gouvernail direction","Volant","Manche","Freins"]`, "Palonnier gouvernail direction", "Palonnier controle roulette nez ou gouvernail au sol.", "Rudder pedals control nose wheel or rudder on ground."},
		{"taxi", 2, "Vitesse max recommandee roulage ?", "Max recommended taxi speed?", `["Marche rapide 10-15 kt","30 kt","50 kt","100 kt"]`, "Marche rapide 10-15 kt", "Vitesse roulage = marche rapide.", "Taxi speed = fast walking pace."},
		{"before_takeoff", 2, "Que faire avant decollage ?", "What to do before takeoff?", `["Check-list decollage essai moteur","Decoller","Verifier meteo","Attendre"]`, "Check-list decollage essai moteur", "Check-list decollage + essai moteur pleine puissance.", "Takeoff checklist + full power run-up."},
		{"before_takeoff", 1, "A quoi sert essai moteur ?", "What is engine run-up for?", `["Verifier magnetos pas helice","Chauffer","Vidanger","Remplir carburant"]`, "Verifier magnetos pas helice", "Essai moteur verifie magnetos et helice.", "Run-up checks magnetos and propeller."},
		{"approach", 2, "Vitesse d approche ?", "Approach speed?", `["1.3 x Vs","Vs","2 x Vs","Max"]`, "1.3 x Vs", "Vitesse approche = 1.3 fois vitesse decrochage.", "Approach speed = 1.3 times stall speed."},
		{"approach", 1, "Quand sortir volets ?", "When to extend flaps?", `["En approche","Decollage","Croisiere","Jamais"]`, "En approche", "Volets augmentent portance basse vitesse approche.", "Flaps increase lift low speed approach."},
		{"landing_proc", 2, "Approche interrompue go-around ?", "Go-around procedure?", `["Pleine puissance reduire volets remonter","Freiner","Virer","Reduire gaz"]`, "Pleine puissance reduire volets remonter", "Pleine puissance reduire volets progressivement remonter.", "Full power reduce flaps gradually climb."},
		{"landing_proc", 1, "Quand sortir train fixe ?", "When to extend fixed gear?", `["Toujours sorti","En approche","Decollage","Croisiere"]`, "Toujours sorti", "Train fixe toujours sorti aucune manuvre.", "Fixed gear always down no operation."},
		{"emergency", 3, "Panne moteur au decollage ?", "Engine failure on takeoff?", `["Atterrir droit devant","Demi-tour","Monter","Appeler controle"]`, "Atterrir droit devant", "Panne decollage atterrir droit devant si possible.", "Takeoff failure land straight ahead if possible."},
		{"emergency", 2, "Feu moteur que faire ?", "Engine fire what to do?", `["Couper carburant plein gaz atterrir","Extincteur","Continuer","Ouvrir capot"]`, "Couper carburant plein gaz atterrir", "Couper carburant plein gaz evacuer atterrir.", "Cut fuel full throttle clear land."},
		{"emergency", 1, "Code transpondeur panne radio ?", "Transponder code radio failure?", `["7600","7500","7700","7000"]`, "7600", "7600 = panne communication radio.", "7600 = radio communication failure."},
		{"emergency", 1, "Code transpondeur detresse ?", "Transponder code distress?", `["7700","7600","7500","7000"]`, "7700", "7700 = situation urgence bord.", "7700 = emergency on board."},
		{"parking", 1, "Que faire apres atterrissage parking ?", "What to do after landing parking?", `["Frein parking couper moteur check-list","Sortir","Manger","Carburant"]`, "Frein parking couper moteur check-list", "Frein parking magnetos OFF mixture pauvre check-list arret.", "Parking brake magnetos OFF lean mixture shutdown checklist."},
	}, count)
}

func genPrinciplesOfFlight(count int) []QuestionData {
	return repeatQuestions([]QuestionData{
		{"lift", 1, "Qu est ce que la portance ?", "What is lift?", `["Force aerodynamique vers haut","Force avant","Force frein","Poids"]`, "Force aerodynamique vers haut", "Portance perpendiculaire vent relatif vers haut.", "Lift perpendicular to relative wind upward."},
		{"lift", 2, "Facteurs influencant portance ?", "Factors affecting lift?", `["Vitesse densite surface coefficient","Poids","Moteur","Carburant"]`, "Vitesse densite surface coefficient", "Portance = 1/2 x rho x V2 x S x Cl.", "Lift = 1/2 x rho x V2 x S x Cl."},
		{"lift", 1, "Qu est ce que le decrochage stall ?", "What is a stall?", `["Perte portance angle attaque trop eleve","Panne moteur","Arret vol","Descente rapide"]`, "Perte portance angle attaque trop eleve", "Decrochage quand angle attaque depasse angle critique.", "Stall when AoA exceeds critical angle."},
		{"lift", 2, "Comment recuperer decrochage ?", "How to recover from stall?", `["Pousser manche plein gaz","Tirer manche","Reduire gaz","Virer"]`, "Pousser manche plein gaz", "Pousser manche reduit AoA plein gaz augmente vitesse.", "Push stick reduces AoA full power increases speed."},
		{"drag", 2, "Deux types de trainee ?", "Two types of drag?", `["Parasite et induite","Avant et arriere","Haute et basse","Lisse et rugueuse"]`, "Parasite et induite", "Trainee parasite frottement + induite generee par portance.", "Parasite drag friction + induced drag from lift."},
		{"drag", 3, "Quand trainee induite max ?", "When induced drag max?", `["Basse vitesse fort angle attaque","Haute vitesse","Pique","Sol"]`, "Basse vitesse fort angle attaque", "Trainee induite max basse vitesse decollage atterrissage.", "Induced drag max low speed takeoff landing."},
		{"drag", 1, "Qu est ce que la trainee parasite ?", "What is parasite drag?", `["Frottement air surface avion","Portance","Poids","Poussee"]`, "Frottement air surface avion", "Trainee parasite due au frottement de l air sur les surfaces.", "Parasite drag from air friction on surfaces."},
		{"drag", 2, "Qu est ce que la trainee de frottement ?", "What is skin friction drag?", `["Frottement air surface aile","Portance","Poids","Poussee"]`, "Frottement air surface aile", "Trainee frottement due a la viscosite de l air.", "Skin friction drag from air viscosity."},
		{"drag", 3, "Qu est ce que la trainee de forme ?", "What is form drag?", `["Resistance due forme objet","Frottement surface","Portance","Poids"]`, "Resistance due forme objet", "Trainee forme causee par la forme de l objet.", "Form drag caused by object shape."},
		{"wing", 1, "Qu est ce qu un profil d aile ?", "What is an airfoil?", `["Forme coupe transversale aile","Longueur aile","Envergure","Surface"]`, "Forme coupe transversale aile", "Profil d aile concu pour generer portance.", "Airfoil designed to generate lift."},
		{"wing", 2, "Qu est ce que l angle d attaque ?", "What is angle of attack?", `["Angle corde profil vent relatif","Angle avion horizon","Inclinaison virage","Pente"]`, "Angle corde profil vent relatif", "Angle attaque angle entre profil et vent relatif.", "AoA angle between airfoil and relative wind."},
		{"wing", 3, "Qu est ce que l angle de calage ?", "What is angle of incidence?", `["Angle fixe aile fuselage","Angle attaque","Angle virage","Pente"]`, "Angle fixe aile fuselage", "Angle calage angle fixe entre aile et fuselage.", "Incidence fixed angle between wing and fuselage."},
		{"wing", 1, "Qu est ce que l envergure ?", "What is wingspan?", `["Distance entre extremites ailes","Longueur fuselage","Surface aile","Corde"]`, "Distance entre extremites ailes", "Envergure distance d un bout d aile a l autre.", "Wingspan distance from wingtip to wingtip."},
		{"wing", 2, "Qu est ce que la corde d une aile ?", "What is wing chord?", `["Distance bord attaque bord fuite","Envergure","Epaisseur","Longueur"]`, "Distance bord attaque bord fuite", "Corde distance entre bord d attaque et bord de fuite.", "Chord distance between leading and trailing edge."},
		{"wing", 3, "Qu est ce que l allongement ?", "What is aspect ratio?", `["Envergure^2 / Surface","Surface / Envergure","Envergure x Surface","Corde x Envergure"]`, "Envergure^2 / Surface", "Allongement rapport envergure au carre sur surface.", "Aspect ratio wingspan squared over area."},
		{"stability", 2, "Stabilite longitudinale ?", "Longitudinal stability?", `["Stabilite axe transversal tangage","Stabilite roulis","Stabilite lacet","Stabilite vitesse"]`, "Stabilite axe transversal tangage", "Stabilite longitudinale concerne mouvement tangage.", "Longitudinal stability concerns pitch movement."},
		{"stability", 1, "Element stabilite en lacet ?", "Element for yaw stability?", `["Derive verticale","Stabilisateur horizontal","Profondeur","Aileron"]`, "Derive verticale", "Derive verticale stabilise autour axe lacet.", "Vertical fin stabilizes around yaw axis."},
		{"stability", 2, "Element stabilite en roulis ?", "Element for roll stability?", `["Dihedre aile","Derive","Profondeur","Aileron"]`, "Dihedre aile", "Dihedre angle aile vers haut stabilise roulis.", "Dihedral wing upward angle stabilizes roll."},
		{"stability", 3, "Qu est ce que le dihedre ?", "What is dihedral?", `["Angle aile vers haut","Angle aile vers bas","Angle fleche","Angle calage"]`, "Angle aile vers haut", "Dihedre angle aile par rapport horizontale vers haut.", "Dihedral wing angle upward from horizontal."},
		{"stability", 1, "Qu est ce que la stabilite statique ?", "What is static stability?", `["Tendance revenir position initiale","Tendance continuer mouvement","Absence mouvement","Mouvement aleatoire"]`, "Tendance revenir position initiale", "Stabilite statique tendance a revenir a l equilibre.", "Static stability tendency to return to equilibrium."},
		{"stability", 2, "Qu est ce que la stabilite dynamique ?", "What is dynamic stability?", `["Comportement temps apres perturbation","Stabilite initiale","Stabilite laterale","Stabilite longitudinale"]`, "Comportement temps apres perturbation", "Stabilite dynamique evolution mouvement dans le temps.", "Dynamic stability motion evolution over time."},
		{"controls", 1, "A quoi servent ailerons ?", "What are ailerons for?", `["Controler roulis inclinaison","Controler tangage","Controler lacet","Freiner"]`, "Controler roulis inclinaison", "Ailerons controlent mouvement roulis.", "Ailerons control roll movement."},
		{"controls", 1, "A quoi sert gouverne profondeur ?", "What is elevator for?", `["Controler tangage monter descendre","Controler roulis","Controler lacet","Diriger sol"]`, "Controler tangage monter descendre", "Profondeur controle assiette longitudinale.", "Elevator controls pitch attitude."},
		{"controls", 1, "A quoi sert gouvernail direction ?", "What is rudder for?", `["Controler lacet coordonner virages","Controler roulis","Monter","Descendre"]`, "Controler lacet coordonner virages", "Gouvernail controle lacet et coordonne virages.", "Rudder controls yaw and coordinates turns."},
		{"controls", 2, "Qu est ce que le trim ?", "What is trim?", `["Compenser forces gouvernes vol stable","Corriger cap","Augmenter vitesse","Reduire puissance"]`, "Compenser forces gouvernes vol stable", "Trim permet vol stable sans force sur commandes.", "Trim allows stable flight without force on controls."},
		{"controls", 3, "Qu est ce que le compensateur ?", "What is trim tab?", `["Petite gouverne sur gouverne principale","Type aileron","Type volet","Type frein"]`, "Petite gouverne sur gouverne principale", "Compensateur petite surface mobile sur gouverne.", "Trim tab small movable surface on control surface."},
		{"controls", 2, "Qu est ce que les volets flaps ?", "What are flaps?", `["Dispositifs hypersustentateurs bord fuite","Ailerons","Gouvernes","Freins"]`, "Dispositifs hypersustentateurs bord fuite", "Volets augmentent portance et trainee basse vitesse.", "Flaps increase lift and drag at low speed."},
		{"controls", 1, "Qu est ce que les becs de bord attaque ?", "What are slats?", `["Dispositifs bord attaque augmentent portance","Volets","Ailerons","Freins"]`, "Dispositifs bord attaque augmentent portance", "Becs bord attaque retardent decrochage.", "Slats delay stall onset."},
		{"controls", 3, "Qu est ce que les spoilers ?", "What are spoilers?", `["Panneaux reduisent portance augmentent trainee","Volets","Ailerons","Gouvernes"]`, "Panneaux reduisent portance augmentent trainee", "Spoilers perturbent ecoulement air reduisent portance.", "Spoilers disrupt airflow reduce lift."},
		{"forces", 2, "Quatre forces sur avion en vol ?", "Four forces on aircraft in flight?", `["Portance poids trainee poussee","Vitesse altitude direction temps","Air vent nuages pluie","Droite gauche haut bas"]`, "Portance poids trainee poussee", "4 forces portance haut poids bas trainee arriere poussee avant.", "4 forces lift up weight down drag back thrust forward."},
		{"forces", 3, "Quand 4 forces en equilibre ?", "When 4 forces in balance?", `["Vol stabilise vitesse constante","Decollage","Atterrissage","Virage"]`, "Vol stabilise vitesse constante", "Vol stabilise portance=poids poussee=trainee.", "Steady flight lift=weight thrust=drag."},
		{"forces", 1, "Qu est ce que le facteur de charge ?", "What is load factor?", `["Rapport portance/poids","Rapport poussee/trainee","Vitesse","Altitude"]`, "Rapport portance/poids", "Facteur charge n = portance / poids.", "Load factor n = lift / weight."},
		{"forces", 2, "Facteur charge en virage 60 degres ?", "Load factor in 60 degree bank?", `["2 G","1 G","3 G","4 G"]`, "2 G", "Virage 60 degres facteur charge 2 G.", "60 degree bank load factor 2 G."},
		{"forces", 3, "Facteur charge en virage 45 degres ?", "Load factor in 45 degree bank?", `["1.414 G","2 G","1 G","3 G"]`, "1.414 G", "Virage 45 degres facteur charge 1.414 G.", "45 degree bank load factor 1.414 G."},
		{"stall", 2, "Vitesse decrochage augmente avec ?", "Stall speed increases with?", `["Facteur charge eleve","Altitude basse","Temperature basse","Poids faible"]`, "Facteur charge eleve", "Vs augmente avec racine carree facteur charge.", "Vs increases with square root of load factor."},
		{"stall", 1, "Qu est ce que le decrochage dynamique ?", "What is accelerated stall?", `["Decrochage a vitesse plus haute en virage","Decrochage normal","Decrochage basse vitesse","Decrochage haute altitude"]`, "Decrochage a vitesse plus haute en virage", "Decrochage dynamique survient a vitesse plus haute en virage.", "Accelerated stall occurs at higher speed in turn."},
		{"stall", 3, "Qu est ce que le decrochage de la gouverne ?", "What is control surface stall?", `["Perte efficacite gouverne angle attaque eleve","Decrochage aile","Panne commandes","Blocage gouvernes"]`, "Perte efficacite gouverne angle attaque eleve", "Decrochage gouverne perte efficacite a grand angle.", "Control stall loss of effectiveness at high angle."},
		{"aerodynamics", 2, "Qu est ce que l effet Bernoulli ?", "What is Bernoulli effect?", `["Pression diminue quand vitesse augmente","Pression augmente avec vitesse","Temperature constante","Densite constante"]`, "Pression diminue quand vitesse augmente", "Bernoulli pression diminue quand vitesse augmente.", "Bernoulli pressure decreases when speed increases."},
		{"aerodynamics", 1, "Qu est ce que le nombre de Mach ?", "What is Mach number?", `["Rapport vitesse avion/vitesse son","Vitesse avion","Vitesse son","Altitude"]`, "Rapport vitesse avion/vitesse son", "Mach = vitesse avion / vitesse du son.", "Mach = aircraft speed / speed of sound."},
		{"aerodynamics", 3, "Qu est ce que le mur du son ?", "What is sound barrier?", `["Phenomenes pres Mach 1","Vitesse max","Altitude max","Portance max"]`, "Phenomenes pres Mach 1", "Mur du son ondes choc et trainee pres Mach 1.", "Sound barrier shock waves and drag near Mach 1."},
		{"aerodynamics", 2, "Qu est ce que l onde de choc ?", "What is shock wave?", `["Variation brutale pression vitesse","Onde sonore","Turbulence","Vent"]`, "Variation brutale pression vitesse", "Onde choc variation brutale pression et vitesse.", "Shock wave abrupt pressure and speed change."},
	}, count)
}

func genCommunications(count int) []QuestionData {
	return repeatQuestions([]QuestionData{
		{"phraseology", 1, "Frequence auto-information VFR ?", "VFR self-announce frequency?", `["123.500 MHz","118.000 MHz","121.500 MHz","126.700 MHz"]`, "123.500 MHz", "123.500 MHz frequence auto-information VFR.", "123.500 MHz VFR self-announce frequency."},
		{"phraseology", 2, "Que signifie Mayday ?", "What does Mayday mean?", `["Detresse","Urgence","Panne","Information"]`, "Detresse", "Mayday signal detresse international aviation.", "Mayday international distress signal aviation."},
		{"phraseology", 1, "Que signifie Pan-Pan ?", "What does Pan-Pan mean?", `["Urgence","Detresse","Information","Avertissement"]`, "Urgence", "Pan-Pan urgence moins grave que Mayday.", "Pan-Pan urgency less critical than Mayday."},
		{"phraseology", 2, "Annoncer intention penetrer zone ?", "Announce intention to enter zone?", `["Trafic callsign position intention","Callsign","Position","Rien"]`, "Trafic callsign position intention", "Structure qui callsign ou position quoi intention.", "Structure who callsign where position what intention."},
		{"phraseology", 1, "Que signifie Wilco ?", "What does Wilco mean?", `["Will comply jobeis","Will contact","Will continue","Will confirm"]`, "Will comply jobeis", "Wilco pilote va obeir instruction.", "Wilco pilot will comply with instruction."},
		{"phraseology", 1, "Que signifie Roger ?", "What does Roger mean?", `["Message recu","Message compris","Message envoye","Message confirme"]`, "Message recu", "Roger message recu pas accord.", "Roger message received not agreement."},
		{"phraseology", 2, "Affirm et Negative ?", "Affirm and Negative?", `["Oui et Non","Peut-etre","Vrai et Faux","Bon et Mauvais"]`, "Oui et Non", "Affirm oui Negative non phraselogie radio.", "Affirm yes Negative no radio phraseology."},
		{"phraseology", 3, "Comment lire un cap a la radio ?", "How to read heading on radio?", `["Chiffre par chiffre","Dizaines","Centaines","Normalement"]`, "Chiffre par chiffre", "Cap se lit chiffre par chiffre 270 = two seven zero.", "Heading digit by digit 270 = two seven zero."},
		{"phraseology", 2, "Comment lire frequence a la radio ?", "How to read frequency on radio?", `["Chiffre par chiffre","MHz","Normalement","kHz"]`, "Chiffre par chiffre", "Frequence se lit chiffre par chiffre.", "Frequency digit by digit."},
		{"phraseology", 1, "Que signifie Over ?", "What does Over mean?", `["A vous parler","Fin transmission","Compris","Recu"]`, "A vous parler", "Over j ai fini de parler a vous.", "Over I have finished speaking go ahead."},
		{"phraseology", 2, "Que signifie Out ?", "What does Out mean?", `["Fin transmission sans reponse","A vous","Compris","Recu"]`, "Fin transmission sans reponse", "Out fin de transmission pas de reponse attendue.", "Out end of transmission no reply expected."},
		{"phraseology", 3, "Que signifie Read back ?", "What does Read back mean?", `["Repeter instruction pour confirmer","Lire","Ecouter","Transmettre"]`, "Repeter instruction pour confirmer", "Read back repeter instruction pour confirmer reception.", "Read back repeat instruction to confirm receipt."},
		{"phraseology", 2, "Que signifie Stand by ?", "What does Stand by mean?", `["Attendre","Pret","Debout","Pret a transmettre"]`, "Attendre", "Stand by attendez je reviens.", "Stand by wait I will come back."},
		{"phraseology", 1, "Que signifie Go ahead ?", "What does Go ahead mean?", `["Transmettez","Allez","Partez","Continuez"]`, "Transmettez", "Go ahead vous pouvez transmettre.", "Go ahead you may transmit."},
		{"phraseology", 2, "Que signifie Words twice ?", "What does Words twice mean?", `["Repeter chaque mot deux fois","Deux mots","Double mot","Mot long"]`, "Repeter chaque mot deux fois", "Words twice repeter chaque mot deux fois cause mauvaise reception.", "Words twice repeat each word twice due to poor reception."},
		{"phraseology", 3, "Que signifie Correction ?", "What does Correction mean?", `["Erreur dans dernier message","Corriger","Modifier","Changer"]`, "Erreur dans dernier message", "Correction j ai fait une erreur voici le bon message.", "Correction I made an error here is the correct message."},
		{"phraseology", 1, "Que signifie I say again ?", "What does I say again mean?", `["Je repete","Je dis","J ecoute","Je transmets"]`, "Je repete", "I say again je vais repeter le message.", "I say again I will repeat the message."},
		{"phraseology", 2, "Que signifie How do you read ?", "What does How do you read mean?", `["Comment recevez vous","Comment lire","Comment transmettre","Comment ecouter"]`, "Comment recevez vous", "How do you read demande qualite reception.", "How do you read asks reception quality."},
		{"phraseology", 1, "Echelle lecture qualite reception ?", "Readability scale?", `["1-5","1-3","1-10","A-E"]`, "1-5", "Echelle 1 illisible 2 lisible par moments 3 lisible avec difficulte 4 lisible 5 parfaitement lisible.", "Scale 1 unreadable 2 readable now and then 3 readable with difficulty 4 readable 5 perfectly readable."},
		{"radio_procedure", 1, "Quand contacter controle aerien ?", "When to contact ATC?", `["Avant penetrer espace controle","En vol","Decollage","Atterrissage"]`, "Avant penetrer espace controle", "Contacter controle avant entrer espace controle.", "Contact ATC before entering controlled airspace."},
		{"radio_procedure", 2, "Si on ne comprend pas instruction ?", "If you don't understand instruction?", `["Demander Say again","Deviner","Ignorer","Repeter hasard"]`, "Demander Say again", "Demander Say again pour faire repeter.", "Ask Say again to have repeated."},
		{"radio_procedure", 1, "Info donner au premier appel ?", "Info on first call?", `["Callsign complet","Position","Altitude","Intention"]`, "Callsign complet", "Commencer par indicatif callsign complet.", "Start with full callsign."},
		{"radio_procedure", 2, "Structure message radio standard ?", "Standard radio message structure?", `["Qui ou quoi","Qui ou quoi","Qui ou quoi","Qui ou quoi"]`, "Qui ou quoi", "Structure qui vous appelez qui vous etes ou vous etes quoi vous faites.", "Structure who you call who you are where you are what you do."},
		{"radio_procedure", 3, "Que faire si reception mauvaise ?", "What if reception is poor?", `["Changer frequence ou monter altitude","Crier","Parler vite","Arreter radio"]`, "Changer frequence ou monter altitude", "Changer frequence ou monter pour meilleure reception.", "Change frequency or climb for better reception."},
		{"radio_procedure", 2, "Procedure changement frequence ?", "Frequency change procedure?", `["Accuser reception nouvelle frequence","Changer directement","Attendre","Couper radio"]`, "Accuser reception nouvelle frequence", "Confirmer nouvelle frequence avant de changer.", "Confirm new frequency before changing."},
		{"radio_procedure", 1, "Quand arreter transmission ?", "When to stop transmitting?", `["Quand controle parle","Jamais","Toujours","Apres 30 sec"]`, "Quand controle parle", "Ne pas transmettre quand controle parle.", "Do not transmit when ATC is speaking."},
		{"radio_procedure", 3, "Procedure panne radio en zone controlee ?", "Radio failure procedure in controlled airspace?", `["Squawk 7600 suivre plan vol","Squawk 7700","Squawk 7500","Squawk 7000"]`, "Squawk 7600 suivre plan vol", "Panne radio squawk 7600 continuer selon plan vol.", "Radio failure squawk 7600 continue according to flight plan."},
		{"radio_procedure", 2, "Que faire si oublie de fermer plan vol ?", "What if forgot to close flight plan?", `["Contacter ATS rapidement","Rien","Attendre","Replanifier"]`, "Contacter ATS rapidement", "Contacter ATS pour fermer plan vol evite operations recherche.", "Contact ATS to close flight plan avoids search operations."},
		{"radio_procedure", 1, "Frequence de detresse secondaire ?", "Secondary distress frequency?", `["243.0 MHz","121.5 MHz","123.5 MHz","126.7 MHz"]`, "243.0 MHz", "243.0 MHz frequence detresse militaire secondaire.", "243.0 MHz military secondary distress frequency."},
		{"radio_procedure", 2, "Quand utiliser Pan-Pan ?", "When to use Pan-Pan?", `["Situation urgence sans danger immediat","Detresse","Information","Routine"]`, "Situation urgence sans danger immediat", "Pan-Pan pour urgence pas detresse immediate.", "Pan-Pan for urgency not immediate distress."},
		{"radio_procedure", 3, "Quand utiliser Mayday ?", "When to use Mayday?", `["Danger grave imminent","Urgence mineure","Information","Routine"]`, "Danger grave imminent", "Mayday pour danger grave et imminent.", "Mayday for grave and imminent danger."},
		{"radio_procedure", 2, "Que faire apres emission Mayday ?", "What to do after Mayday transmission?", `["Ecouter reponse suivre instructions","Repeter Mayday","Changer frequence","Couper radio"]`, "Ecouter reponse suivre instructions", "Apres Mayday ecouter reponse et suivre instructions.", "After Mayday listen for response and follow instructions."},
		{"radio_procedure", 1, "Qui peut annuler Mayday ?", "Who can cancel Mayday?", `["Le pilote qui l a emis","Controle","Tout le monde","Personne"]`, "Le pilote qui l a emis", "Seul le pilote qui a emis Mayday peut l annuler.", "Only the pilot who issued Mayday can cancel it."},
		{"radio_procedure", 2, "Que signifie Silence Mayday ?", "What does Silence Mayday mean?", `["Toutes stations arreter transmissions detresse","Arreter Mayday","Silence radio","Fin urgence"]`, "Toutes stations arreter transmissions detresse", "Silence Mayday ordre de silence radio pour detresse.", "Silence Mayday order for radio silence for distress."},
		{"radio_procedure", 3, "Que signifie Silence fini ?", "What does Silence finished mean?", `["Fin silence radio reprendre transmissions","Fin detresse","Fin urgence","Fin Mayday"]`, "Fin silence radio reprendre transmissions", "Silence fini les stations peuvent reprendre transmissions.", "Silence finished stations may resume transmissions."},
		{"radio_procedure", 1, "Frequence auto-information aerodrome ?", "Aerodrome self-announce frequency?", `["Frequence AFIS ou 123.500","121.500","118.000","126.700"]`, "Frequence AFIS ou 123.500", "Utiliser frequence AFIS ou auto-info aerodrome.", "Use AFIS frequency or aerodrome self-announce."},
		{"radio_procedure", 2, "Annoncer position au roulage ?", "Announce position during taxi?", `["Trafic position intention","Callsign seulement","Position seulement","Rien"]`, "Trafic position intention", "Annoncer qui vous etes ou vous etes et votre intention.", "Announce who you are where you are and your intention."},
		{"radio_procedure", 3, "Annoncer depart piste ?", "Announce departure runway?", `["Trafic depart piste nom","Piste seulement","Nom seulement","Rien"]`, "Trafic depart piste nom", "Annoncer depart piste nom et direction.", "Announce departure runway name and direction."},
		{"radio_procedure", 1, "Annoncer approche finale ?", "Announce final approach?", `["Trafic finale piste nom","Piste seulement","Nom seulement","Rien"]`, "Trafic finale piste nom", "Annoncer finale piste nom et position.", "Announce final runway name and position."},
		{"radio_procedure", 2, "Annoncer atterrissage ?", "Announce landing?", `["Trafic atterri piste nom","Piste seulement","Nom seulement","Rien"]`, "Trafic atterri piste nom", "Annoncer atterrissage piste et nom.", "Announce landing runway and name."},
		{"radio_procedure", 3, "Annoncer roulage apres atterrissage ?", "Announce after landing taxi?", `["Trafic degage piste nom","Piste seulement","Nom seulement","Rien"]`, "Trafic degage piste nom", "Annoncer avoir degage piste et nom.", "Announce runway vacated and name."},
	}, count)
}

func genMassAndBalance(count int) []QuestionData {
	return repeatQuestions([]QuestionData{
		{"basics", 1, "Masse a vide empty weight ?", "What is empty weight?", `["Masse sans charge ni carburant","Masse max","Masse avec carburant","Masse decollage"]`, "Masse sans charge ni carburant", "Masse a vide masse de base avion.", "Empty weight basic aircraft mass."},
		{"basics", 2, "Qu est ce que MTOW ?", "What is MTOW?", `["Masse max certifiee decollage","Masse max vol","Masse a vide","Masse atterrissage"]`, "Masse max certifiee decollage", "MTOW Maximum TakeOff Weight masse max decollage.", "MTOW Maximum TakeOff Weight."},
		{"basics", 1, "Qu est ce que centrage CG ?", "What is center of gravity CG?", `["Point equilibre avion","Centre avion","Milieu fuselage","Centre poussee"]`, "Point equilibre avion", "Centrage point ou avion en equilibre.", "CG point where aircraft balances."},
		{"basics", 2, "Pourquoi centrage important ?", "Why CG important?", `["Affecte stabilite performances","Confort","Esthetique","Bruit"]`, "Affecte stabilite performances", "Centrage incorrect peut rendre avion instable.", "Incorrect CG can make aircraft unstable."},
		{"basics", 3, "Centrage trop avant ?", "CG too far forward?", `["Stable mais manuvrabilite reduite","Instable","Decrochage","Incontrolable"]`, "Stable mais manuvrabilite reduite", "Centrage avant stable mais difficile arrondir.", "Forward CG stable but hard to flare."},
		{"basics", 3, "Centrage trop arriere ?", "CG too far aft?", `["Instable risque decrochage","Meilleure perf","Plus portance","Moins trainee"]`, "Instable risque decrochage", "Centrage arriere instable risque decrochage virage.", "Aft CG unstable stall risk turns."},
		{"basics", 2, "Qu est ce que la masse max atterrissage MLW ?", "What is MLW?", `["Masse max certifiee atterrissage","Masse max decollage","Masse a vide","Masse carburant"]`, "Masse max certifiee atterrissage", "MLW Maximum Landing Weight masse max atterrissage.", "MLW Maximum Landing Weight."},
		{"basics", 1, "Qu est ce que la masse max zero carburant MZFW ?", "What is MZFW?", `["Masse max sans carburant utilisable","Masse max decollage","Masse a vide","Masse atterrissage"]`, "Masse max sans carburant utilisable", "MZFW masse max sans carburant dans reservoirs.", "MZFW max weight without usable fuel."},
		{"basics", 3, "Qu est ce que la masse max ramp MRW ?", "What is MRW?", `["Masse max au parking","Masse max decollage","Masse max atterrissage","Masse a vide"]`, "Masse max au parking", "MRW Maximum Ramp Weight masse max au parking.", "MRW Maximum Ramp Weight at parking."},
		{"basics", 2, "Qu est ce que la charge utile payload ?", "What is payload?", `["Masse passagers bagages fret","Masse carburant","Masse avion vide","Masse totale"]`, "Masse passagers bagages fret", "Charge utile masse transportee generant revenu.", "Payload mass carried generating revenue."},
		{"basics", 1, "Qu est ce que la charge marchande ?", "What is commercial load?", `["Charge utile + carburant","Charge utile seulement","Carburant seulement","Avion vide"]`, "Charge utile + carburant", "Charge marchande tout ce qui rapporte de l argent.", "Commercial load everything generating revenue."},
		{"basics", 3, "Qu est ce que le carburant utilisable ?", "What is usable fuel?", `["Carburant pouvant etre consomme en vol","Carburant total reservoirs","Carburant reserve","Carburant taxi"]`, "Carburant pouvant etre consomme en vol", "Carburant utilisable accessible par moteur en vol.", "Usable fuel accessible by engine in flight."},
		{"basics", 2, "Qu est ce que le carburant inutilisable ?", "What is unusable fuel?", `["Carburant restant reservoirs non accessible","Carburant reserve","Carburant taxi","Carburant route"]`, "Carburant restant reservoirs non accessible", "Carburant inutilisable reste dans reservoirs.", "Unusable fuel remains in tanks."},
		{"calculation", 2, "Comment calculer moment ?", "How to calculate moment?", `["Masse x Bras","Masse + Bras","Masse / Bras","Bras - Masse"]`, "Masse x Bras", "Moment = masse x distance point reference.", "Moment = weight x distance from datum."},
		{"calculation", 1, "Unite du centrage ?", "CG unit?", `["% corde ou mm","kg","m/s","degres"]`, "% corde ou mm", "Centrage en % corde ou mm.", "CG in % MAC or mm."},
		{"calculation", 3, "Qu est ce que le bras de levier ?", "What is lever arm?", `["Distance entre charge et point reference","Longueur avion","Envergure","Corde"]`, "Distance entre charge et point reference", "Bras distance entre charge et point de reference.", "Arm distance between load and datum."},
		{"calculation", 2, "Qu est ce que le point de reference datum ?", "What is datum?", `["Point reference fixe pour calculs centrage","Centre avion","Bord attaque aile","Bord fuite aile"]`, "Point reference fixe pour calculs centrage", "Datum point fixe pour tous calculs moments.", "Datum fixed point for all moment calculations."},
		{"calculation", 1, "Qu est ce que la corde moyenne aerodynamique MAC ?", "What is MAC?", `["Corde moyenne aerodynamique aile","Corde aile","Envergure","Surface"]`, "Corde moyenne aerodynamique aile", "MAC Mean Aerodynamic Chord corde moyenne aile.", "MAC Mean Aerodynamic Chord."},
		{"calculation", 3, "Centrage en % MAC ?", "CG in % MAC?", `["Position CG en % de la MAC","Position CG en mm","Position CG en m","Position CG en degres"]`, "Position CG en % de la MAC", "% MAC position centrage par rapport corde.", "% MAC CG position relative to chord."},
		{"calculation", 2, "Comment calculer centrage total ?", "How to calculate total CG?", `["Moment total / Masse totale","Masse totale / Moment total","Moment total x Masse","Masse totale - Moment"]`, "Moment total / Masse totale", "Centrage = somme moments / somme masses.", "CG = total moment / total weight."},
		{"loading", 2, "Ou placer charges lourdes ?", "Where to place heavy loads?", `["Pres centrage ideal","Avant","Arriere","Ailes"]`, "Pres centrage ideal", "Charges lourdes pres centrage ideal.", "Heavy loads near ideal CG."},
		{"loading", 1, "Peut on depasser masse max ?", "Can max weight be exceeded?", `["Jamais","Beau temps","5%","Urgence"]`, "Jamais", "Masse max jamais depassee.", "Max weight never exceeded."},
		{"loading", 3, "Ordre chargement avion ?", "Aircraft loading order?", `["Planifier calculer charger verifier","Charger puis calculer","Hasard","Charger tout"]`, "Planifier calculer charger verifier", "Planifier calculs puis charger puis verifier.", "Plan calculate then load then verify."},
		{"loading", 2, "Qu est ce que la soute a bagages ?", "What is baggage compartment?", `["Espace reserve bagages limites masse","Espace libre","Reservoir","Cabine"]`, "Espace reserve bagages limites masse", "Soute bagages avec limite de masse specifique.", "Baggage compartment with specific weight limit."},
		{"loading", 1, "Limite masse soute bagages ?", "Baggage compartment weight limit?", `["Specifiee dans manuel vol","Pas limite","100 kg","50 kg"]`, "Specifiee dans manuel vol", "Limite masse soute dans manuel de vol.", "Compartment limit in flight manual."},
		{"loading", 3, "Qu est ce que le chargement asymetrique ?", "What is asymmetric loading?", `["Charge repartie inegalement","Charge symetrique","Charge nulle","Charge max"]`, "Charge repartie inegalement", "Chargement asymetrique peut affecter controle.", "Asymmetric loading can affect control."},
		{"loading", 2, "Effet chargement asymetrique ?", "Asymmetric loading effect?", `["Roulis permanent ou compensation","Aucun","Ameliore","Reduit trainee"]`, "Roulis permanent ou compensation", "Chargement asymetrique cause roulis permanent.", "Asymmetric loading causes permanent roll."},
		{"envelope", 2, "Qu est ce que l enveloppe de centrage ?", "What is CG envelope?", `["Limites centrage autorisees","Limites vitesse","Limites altitude","Limites poids"]`, "Limites centrage autorisees", "Enveloppe centrage limites avant et arriere.", "CG envelope forward and aft limits."},
		{"envelope", 1, "Que se passe si centrage hors enveloppe ?", "What if CG outside envelope?", `["Vol interdit","Vol possible","Meilleure perf","Moins consommation"]`, "Vol interdit", "Centrage hors enveloppe vol interdit.", "CG outside envelope flight prohibited."},
		{"envelope", 3, "Facteurs affectant enveloppe centrage ?", "Factors affecting CG envelope?", `["Masse altitude temperature","Vitesse","Vent","Meteo"]`, "Masse altitude temperature", "Enveloppe centrage varie avec masse altitude temperature.", "CG envelope varies with weight altitude temperature."},
		{"envelope", 2, "Qu est ce que le diagramme de centrage ?", "What is CG diagram?", `["Graphique masse vs centrage","Graphique vitesse","Graphique altitude","Graphique temps"]`, "Graphique masse vs centrage", "Diagramme centrage montre enveloppe autorisee.", "CG diagram shows authorized envelope."},
		{"envelope", 1, "Qu est ce que la table de chargement ?", "What is loading table?", `["Tableau calcul rapide centrage","Tableau vitesse","Tableau altitude","Tableau temps"]`, "Tableau calcul rapide centrage", "Table chargement permet calcul rapide centrage.", "Loading table allows quick CG calculation."},
		{"effects", 2, "Effet centrage avant sur performances ?", "Forward CG effect on performance?", `["Augmente trainee reduit vitesse","Reduit trainee","Augmente portance","Reduit poids"]`, "Augmente trainee reduit vitesse", "Centrage avant augmente trainee reduit vitesse.", "Forward CG increases drag reduces speed."},
		{"effects", 3, "Effet centrage arriere sur consommation ?", "Aft CG effect on fuel consumption?", `["Reduit consommation moins trainee","Augmente consommation","Pas effet","Double consommation"]`, "Reduit consommation moins trainee", "Centrage arriere reduit trainee donc consommation.", "Aft CG reduces drag therefore consumption."},
		{"effects", 1, "Effet centrage avant sur stabilite ?", "Forward CG effect on stability?", `["Augmente stabilite","Reduit stabilite","Pas effet","Instable"]`, "Augmente stabilite", "Centrage avant augmente stabilite longitudinale.", "Forward CG increases longitudinal stability."},
		{"effects", 2, "Effet centrage arriere sur stabilite ?", "Aft CG effect on stability?", `["Reduit stabilite","Augmente stabilite","Pas effet","Stable"]`, "Reduit stabilite", "Centrage arriere reduit stabilite longitudinale.", "Aft CG reduces longitudinal stability."},
		{"effects", 3, "Effet centrage sur vitesse decrochage ?", "CG effect on stall speed?", `["Centrage arriere reduit Vs","Centrage avant reduit Vs","Pas effet","Double Vs"]`, "Centrage arriere reduit Vs", "Centrage arriere reduit vitesse decrochage.", "Aft CG reduces stall speed."},
		{"effects", 2, "Effet centrage sur distance decollage ?", "CG effect on takeoff distance?", `["Centrage arriere reduit distance","Centrage avant reduit","Pas effet","Double"]`, "Centrage arriere reduit distance", "Centrage arriere reduit distance decollage.", "Aft CG reduces takeoff distance."},
	}, count)
}

func genInstrumentation(count int) []QuestionData {
	return repeatQuestions([]QuestionData{
		{"basic_six", 1, "Six instruments de base ?", "Basic six instruments?", `["ASI altimetre VSI horizon TC compas","Vitesse altitude temps","Moteur carburant huile","Radio GPS transpondeur"]`, "ASI altimetre VSI horizon TC compas", "6 instruments anemometre altimetre variometre horizon coordinateur compas.", "Basic 6 ASI altimeter VSI attitude indicator turn coordinator compass."},
		{"asic", 1, "Que mesure ASI ?", "What does ASI measure?", `["Vitesse indiquee IAS","Vitesse sol","Vitesse vraie","Vitesse vent"]`, "Vitesse indiquee IAS", "ASI mesure vitesse indiquee Indicated Air Speed.", "ASI measures Indicated Air Speed IAS."},
		{"altimeter", 1, "Que mesure altimetre ?", "What does altimeter measure?", `["Altitude pression","Altitude vraie","Altitude radio","Hauteur"]`, "Altitude pression", "Altimetre mesure altitude basee sur pression.", "Altimeter measures pressure altitude."},
		{"altimeter", 2, "Comment caler altimetre ?", "How to set altimeter?", `["Regler QNH ou QFE","Tourner molette hasard","Bouton","Auto"]`, "Regler QNH ou QFE", "Regler QNH altitude ou QFE hauteur fenetre.", "Set QNH altitude or QFE height window."},
		{"vsi", 1, "Que mesure VSI ?", "What does VSI measure?", `["Vitesse verticale ft/min","Vitesse horizontale","Altitude","Pression"]`, "Vitesse verticale ft/min", "VSI indique vitesse montee ou descente.", "VSI shows rate of climb or descent."},
		{"attitude", 1, "Que montre horizon artificiel ?", "What does attitude indicator show?", `["Assiette avion rapport horizon","Cap","Altitude","Vitesse"]`, "Assiette avion rapport horizon", "Horizon artificiel montre inclinaison et assiette.", "Attitude indicator shows pitch and bank."},
		{"compass", 2, "Deviation du compas ?", "Compass deviation?", `["Erreur due masses metalliques","Erreur vent","Erreur temperature","Erreur normale"]`, "Erreur due masses metalliques", "Deviation erreur compas causee masses metalliques.", "Deviation compass error from metallic masses."},
		{"compass", 2, "Variation magnetique ?", "Magnetic variation?", `["Difference nord vrai nord magnetique","Erreur compas","Changement cap","Variation saisonniere"]`, "Difference nord vrai nord magnetique", "Variation depend position geographique.", "Variation depends on geographic location."},
		{"turn_coordinator", 1, "A quoi sert TC ?", "What is turn coordinator for?", `["Indiquer taux virage coordination","Indiquer vitesse","Indiquer altitude","Indiquer cap"]`, "Indiquer taux virage coordination", "TC montre taux virage et coordination.", "TC shows turn rate and coordination."},
		{"gps_instr", 2, "Qu est ce qu un ILS ?", "What is an ILS?", `["Systeme atterrissage instruments","GPS","VOR","NDB"]`, "Systeme atterrissage instruments", "ILS Instrument Landing System guide avion atterrissage.", "ILS Instrument Landing System guides aircraft landing."},
		{"gps_instr", 2, "Deux composants ILS ?", "Two ILS components?", `["Localizer LOC et Glideslope GS","VOR et DME","NDB et ADF","GPS et RAIM"]`, "Localizer LOC et Glideslope GS", "Localizer alignement horizontal Glideslope pente verticale.", "Localizer lateral guidance Glideslope vertical guidance."},
		{"gps_instr", 1, "A quoi sert DME ?", "What is DME for?", `["Mesurer distance station","Mesurer altitude","Mesurer vitesse","Mesurer cap"]`, "Mesurer distance station", "DME Distance Measuring Equipment mesure distance oblique.", "DME Distance Measuring Equipment measures slant distance."},
		{"transponder", 1, "A quoi sert transpondeur ?", "What is transponder for?", `["Identifier avion radar","Communiquer","Naviguer","Mesurer altitude"]`, "Identifier avion radar", "Transpondeur permet controle vous identifier.", "Transponder allows ATC to identify you."},
		{"transponder", 2, "Que signifie mode C ?", "What does Mode C mean?", `["Transmet altitude","Transmet cap","Transmet vitesse","Transmet identite"]`, "Transmet altitude", "Mode C transmet altitude pression avion.", "Mode C transmits aircraft pressure altitude."},
	}, count)
}

func genFinalExamQuestions(lic string, count int) []QuestionData {
	// Gather all questions from all categories
	var allQs []QuestionData
	allQs = append(allQs, genAirLaw(100)...)
	allQs = append(allQs, genMeteo(100)...)
	allQs = append(allQs, genNavigation(100)...)
	allQs = append(allQs, genPerformance(100)...)
	allQs = append(allQs, genAircraftGeneral(100)...)
	allQs = append(allQs, genFlightPlanning(100)...)
	allQs = append(allQs, genHumanPerformance(100)...)
	allQs = append(allQs, genOperationalProcedures(100)...)
	allQs = append(allQs, genPrinciplesOfFlight(100)...)
	allQs = append(allQs, genCommunications(100)...)
	allQs = append(allQs, genMassAndBalance(100)...)
	allQs = append(allQs, genInstrumentation(100)...)

	// Filter by difficulty based on license
	var filtered []QuestionData
	for _, q := range allQs {
		if lic == "PPL" || lic == "LAPL" || lic == "BPL" {
			if q.difficulty <= 2 {
				filtered = append(filtered, q)
			}
		} else {
			// ATPL, CPL, IR - all difficulties
			filtered = append(filtered, q)
		}
	}

	// Shuffle and pick
	perm := rand.Perm(len(filtered))
	if count > len(filtered) {
		count = len(filtered)
	}
	result := make([]QuestionData, count)
	for i := 0; i < count; i++ {
		result[i] = filtered[perm[i]]
	}
	return result
}

// lessonTitlesFr and lessonTitlesEn provide specific titles for each lesson in each category
var lessonTitlesFr = map[string][]string{
	"airlaw": {
		"Introduction au droit aerien et a l OACI",
		"L EASA et la reglementation europeenne",
		"Les licences de pilote et certificats medicaux",
		"Les espaces aeriens classes A a G",
		"Les regles de l air SERA",
		"Les priorites de passage et regles anti-abordage",
		"Les documents de bord obligatoires",
		"Les plans de vol et procedures",
		"Les transpondeurs et codes SSR",
		"Les procedures de communication radio",
		"Les vols VFR et IFR regles et conditions",
		"Les vols de nuit et conditions speciales",
		"Les procedures de detresse et urgence",
		"La securite et la surete aerienne",
		"Les assurances et responsabilites du pilote",
		"Les sanctions et infractions aeronautiques",
		"La protection de l environnement et nuisances sonores",
		"Les accords internationaux et bilateralux",
		"Les services de la circulation aerienne ATS",
		"Les publications aeronautiques AIP et NOTAM",
	},
	"meteorology": {
		"L atmosphere terrestre composition et structure",
		"La temperature et les echanges thermiques",
		"La pression atmospherique et ses variations",
		"Le vent general et local",
		"Les masses d air et les fronts",
		"Les nuages classification et formation",
		"Les precipitations et la visibilite",
		"Le brouillard types et formation",
		"Le givrage conditions et dangers",
		"Les orages formation et evolution",
		"Les phenomenes meteorologiques dangereux",
		"Les cartes meteorologiques et symboles",
		"Les messages METAR et TAF",
		"Les messages SIGMET et AIRMET",
		"Les vents en altitude et jet stream",
		"La turbulence et le cisaillement de vent",
		"Les conditions meteorologiques pour le vol VFR",
		"Les conditions meteorologiques pour le vol IFR",
		"La meteorologie tropicale et equatoriale",
		"Les outils de prevision meteorologique",
		"L analyse des cartes isobariques",
		"Les radiosondages et diagrammes aerologiques",
		"Les phenomenes optiques dans l atmosphere",
		"La climatologie et statistiques meteorologiques",
		"La prise de decisions meteorologiques",
	},
	"navigation": {
		"Introduction a la navigation aerienne",
		"La sphere terrestre et les coordonnees geographiques",
		"Les cartes aeronautiques projections et echelles",
		"Le nord vrai magnetique et la declinaison",
		"La navigation au compas deviation et variation",
		"Le calcul du cap et de la route",
		"La vitesse air et vitesse sol",
		"Le vent et la derive navigation",
		"Le triangle des vitesses",
		"La navigation a l estime dead reckoning",
		"Le VOR fonctionnement et utilisation",
		"Le NDB et l ADF fonctionnement",
		"Le DME mesure de distance",
		"Le GPS principe et utilisation",
		"Le RNAV et la navigation de surface",
		"L ILS atterrissage aux instruments",
		"Les procedures de depart et arrivee",
		"Les points de report et balises",
		"La navigation en espace aerien controle",
		"La navigation oceanique et polaire",
		"Les systemes de navigation inertielle",
		"Les erreurs de navigation et corrections",
		"La planification d une navigation longue distance",
		"Les reglementations de navigation",
		"La navigation d urgence et alternatives",
	},
	"performance": {
		"Introduction aux performances avion",
		"Les masses et limitations structurales",
		"La densite de l air et performances",
		"Les distances de decollage",
		"Les facteurs affectant le decollage",
		"Les vitesses V1 Vr et V2",
		"Les performances en montee",
		"Les performances en croisiere",
		"Les distances d atterrissage",
		"Les facteurs affectant l atterrissage",
		"L effet de sol et ses consequences",
		"Les graphiques de performances",
		"Les limitations de vent de travers",
		"Les performances avec panne moteur",
		"La reglementation des performances",
	},
	"aircraft_general": {
		"Introduction a la cellule et structure",
		"Le fuselage construction et materiaux",
		"La voilure structure et profil",
		"L empennage et les gouvernes",
		"Le train d atterrissage types",
		"Le moteur a piston 4 temps",
		"Le systeme d alimentation carburant",
		"Le systeme d allumage et magnetos",
		"Le systeme de refroidissement",
		"Le systeme de lubrification",
		"L helice pas fixe et variable",
		"Le systeme electrique alternateur batterie",
		"Les instruments du moteur",
		"Le systeme de chauffage et ventilation",
		"Les systemes hydrauliques",
		"Le systeme de pressurisation",
		"Les systemes antigivrage et degivrage",
		"L entretien et les inspections",
		"Les limitations d utilisation",
		"Les systemes d oxygene",
	},
	"flight_planning": {
		"Introduction au planning de vol",
		"Les NOTAM et leur interpretation",
		"L AIP et les publications aeronautiques",
		"Le choix de la route et altitudes",
		"Le calcul du carburant necessaire",
		"Les reserves de carburant reglementaires",
		"Le plan de vol formulaire et depot",
		"Les cartes aeronautiques en route",
		"Les cartes d aerodrome",
		"Les espaces aeriens et leur traversee",
		"Les points de report obligatoires",
		"Les frequences radio et balises",
		"Les minima meteorologiques",
		"Les aerodromes de degagement",
		"Le calcul des masses et centrage",
		"La fiche de navigation",
		"Les procedures avant depart",
		"Les redevances et taxes aeronautiques",
		"La documentation de bord",
		"Les procedures d urgence planification",
	},
	"human_performance": {
		"Introduction aux facteurs humains",
		"La vision et ses limitations",
		"L audition et les communications",
		"L oreille interne et l equilibre",
		"L hypoxie causes et symptomes",
		"L hyperventilation",
		"Les barotraumatismes",
		"La fatigue et la vigilance",
		"Le stress et ses effets",
		"Les illusions sensorielles en vol",
		"Les effets de l alcool et medicaments",
		"La nutrition et l hydratation",
		"Le rythme circadien et decalage horaire",
		"La prise de decision aeronautique",
		"Le CRM Crew Resource Management",
		"La gestion des erreurs",
		"La communication en equipe",
		"Les facteurs humains dans les accidents",
		"La charge de travail et priorisation",
		"Les attitudes dangereuses",
	},
	"operational_procedures": {
		"Introduction aux procedures operationnelles",
		"La preparation du vol",
		"Les inspections pre-vol exterieures",
		"Les inspections pre-vol interieures",
		"La mise en route et demarrage moteur",
		"Le roulage et les consignes",
		"Les essais moteur avant decollage",
		"La procedure de decollage normal",
		"La montee initiale et depart",
		"La navigation en route",
		"Les changements de niveau",
		"La descente et l approche",
		"La procedure d atterrissage normal",
		"Le roulage apres atterrissage",
		"L arret moteur et parking",
		"Les procedures par temps de pluie",
		"Les procedures par vent fort",
		"Les procedures de nuit",
		"Les pannes moteur en vol",
		"Les pannes moteur au decollage",
		"Les pannes electriques",
		"Les pannes d instruments",
		"Les procedures de feu en vol",
		"Les atterrissages forces et d urgence",
		"Les operations sur terrains non prepares",
	},
	"principles_of_flight": {
		"Introduction a l aerodynamique",
		"Le profil d aile et ses caracteristiques",
		"La portance generation et facteurs",
		"La trainee parasite et induite",
		"Le coefficient de portance et trainee",
		"Le foyer aerodynamique",
		"Le centre de gravite et stabilite",
		"La stabilite longitudinale",
		"La stabilite laterale et directionnelle",
		"Les gouvernes de vol primaires",
		"Les gouvernes de vol secondaires",
		"Le decrochage et sa recuperation",
		"Le vrille et la recuperation",
		"Les forces en vol stabilise",
		"Les virages et facteur de charge",
		"Les limites de vol et envelope",
		"Les effets du vent traversier",
		"L aerodynamique a haute vitesse",
		"Les dispositifs hypersustentateurs",
		"Les phenomenes aerodynamiques particuliers",
	},
	"communications": {
		"Introduction aux communications aeronautiques",
		"La phraselogie standard",
		"L alphabet international OACI",
		"Les frequences aeronautiques",
		"Les procedures VFR en zone non controlee",
		"Les procedures VFR en zone controlee",
		"Les procedures IFR",
		"Les communications de detresse Mayday",
		"Les communications d urgence Pan-Pan",
		"Les communications en vol",
		"Les communications au sol",
		"Les communications avec les services d information",
		"Les communications en anglais",
		"Les procedures de panne radio",
		"Les communications speciales",
	},
	"mass_and_balance": {
		"Introduction aux masses et centrage",
		"Les definitions masses a vide et charge",
		"Le MTOW MLW et autres limitations",
		"Le centrage et le bras de levier",
		"Le calcul du moment",
		"La determination du centrage",
		"Les enveloppes de centrage",
		"Les graphiques de masse et centrage",
		"Le chargement de l aeronef",
		"Les effets du centrage sur les performances",
		"Les effets du centrage sur la stabilite",
		"Les calculs de masse et centrage pratiques",
		"La reglementation masse et centrage",
		"Les cas particuliers de chargement",
		"Les outils de calcul masse et centrage",
	},
	"instrumentation": {
		"Introduction aux instruments de bord",
		"Le systeme pitot statique",
		"L anemometre ASI",
		"L altimetre principe et calage",
		"Le variometre VSI",
		"L horizon artificiel",
		"Le coordinateur de virage",
		"Le compas magnetique",
		"Le gyrocompas directionnel",
		"Les instruments gyroscopiques",
		"Le radiocompas ADF",
		"Le VOR indicateur",
		"L ILS localizer et glide",
		"Le transpondeur modes A C S",
		"Le GPS et les systemes de navigation",
		"Les instruments du moteur",
		"Les systemes d alarme et avertissement",
		"Les instruments de gestion carburant",
		"Les systemes automatiques pilote automatique",
		"Les pannes d instruments et procedures",
	},
}

var lessonTitlesEn = map[string][]string{
	"airlaw": {
		"Introduction to Air Law and ICAO",
		"EASA and European Regulations",
		"Pilot Licenses and Medical Certificates",
		"Classified Airspace A to G",
		"SERA Rules of the Air",
		"Right of Way and Collision Avoidance",
		"Mandatory Onboard Documents",
		"Flight Plans and Procedures",
		"Transponders and SSR Codes",
		"Radio Communication Procedures",
		"VFR and IFR Flight Rules and Conditions",
		"Night Flights and Special Conditions",
		"Distress and Emergency Procedures",
		"Aviation Safety and Security",
		"Insurance and Pilot Responsibilities",
		"Penalties and Aviation Offenses",
		"Environmental Protection and Noise",
		"International and Bilateral Agreements",
		"Air Traffic Services ATS",
		"Aeronautical Publications AIP and NOTAM",
	},
	"meteorology": {
		"The Atmosphere Composition and Structure",
		"Temperature and Heat Exchange",
		"Atmospheric Pressure and Variations",
		"General and Local Wind",
		"Air Masses and Fronts",
		"Clouds Classification and Formation",
		"Precipitation and Visibility",
		"Fog Types and Formation",
		"Icing Conditions and Dangers",
		"Thunderstorms Formation and Evolution",
		"Dangerous Weather Phenomena",
		"Weather Charts and Symbols",
		"METAR and TAF Messages",
		"SIGMET and AIRMET Messages",
		"Upper Winds and Jet Stream",
		"Turbulence and Wind Shear",
		"VFR Weather Conditions",
		"IFR Weather Conditions",
		"Tropical and Equatorial Meteorology",
		"Weather Forecasting Tools",
		"Isobaric Chart Analysis",
		"Radiosondes and Aerological Diagrams",
		"Optical Phenomena in the Atmosphere",
		"Climatology and Weather Statistics",
		"Meteorological Decision Making",
	},
	"navigation": {
		"Introduction to Air Navigation",
		"The Earth and Geographic Coordinates",
		"Aeronautical Charts Projections and Scales",
		"True North Magnetic North and Declination",
		"Compass Navigation Deviation and Variation",
		"Heading and Track Calculation",
		"Air Speed and Ground Speed",
		"Wind and Drift Navigation",
		"The Wind Triangle",
		"Dead Reckoning Navigation",
		"VOR Operation and Use",
		"NDB and ADF Operation",
		"DME Distance Measurement",
		"GPS Principles and Use",
		"RNAV and Area Navigation",
		"ILS Instrument Landing System",
		"Departure and Arrival Procedures",
		"Reporting Points and Beacons",
		"Navigation in Controlled Airspace",
		"Oceanic and Polar Navigation",
		"Inertial Navigation Systems",
		"Navigation Errors and Corrections",
		"Long Distance Navigation Planning",
		"Navigation Regulations",
		"Emergency Navigation and Alternatives",
	},
	"performance": {
		"Introduction to Aircraft Performance",
		"Weights and Structural Limitations",
		"Air Density and Performance",
		"Takeoff Distances",
		"Factors Affecting Takeoff",
		"V1 Vr and V2 Speeds",
		"Climb Performance",
		"Cruise Performance",
		"Landing Distances",
		"Factors Affecting Landing",
		"Ground Effect and Consequences",
		"Performance Charts",
		"Crosswind Limitations",
		"Engine Failure Performance",
		"Performance Regulations",
	},
	"aircraft_general": {
		"Introduction to Airframe and Structure",
		"Fuselage Construction and Materials",
		"Wing Structure and Airfoil",
		"Empennage and Control Surfaces",
		"Landing Gear Types",
		"Four Stroke Piston Engine",
		"Fuel System",
		"Ignition System and Magnetos",
		"Cooling System",
		"Lubrication System",
		"Fixed and Variable Pitch Propeller",
		"Electrical System Alternator Battery",
		"Engine Instruments",
		"Heating and Ventilation System",
		"Hydraulic Systems",
		"Pressurization System",
		"Anti-icing and De-icing Systems",
		"Maintenance and Inspections",
		"Operating Limitations",
		"Oxygen Systems",
	},
	"flight_planning": {
		"Introduction to Flight Planning",
		"NOTAMs and Their Interpretation",
		"AIP and Aeronautical Publications",
		"Route Selection and Altitudes",
		"Required Fuel Calculation",
		"Regulatory Fuel Reserves",
		"Flight Plan Form and Filing",
		"Enroute Aeronautical Charts",
		"Aerodrome Charts",
		"Airspace Crossing",
		"Mandatory Reporting Points",
		"Radio Frequencies and Beacons",
		"Weather Minima",
		"Alternate Aerodromes",
		"Weight and Balance Calculation",
		"Navigation Log",
		"Pre-departure Procedures",
		"Aeronautical Fees and Charges",
		"Onboard Documentation",
		"Emergency Planning Procedures",
	},
	"human_performance": {
		"Introduction to Human Factors",
		"Vision and Its Limitations",
		"Hearing and Communications",
		"Inner Ear and Balance",
		"Hypoxia Causes and Symptoms",
		"Hyperventilation",
		"Barotrauma",
		"Fatigue and Alertness",
		"Stress and Its Effects",
		"Sensory Illusions in Flight",
		"Effects of Alcohol and Medications",
		"Nutrition and Hydration",
		"Circadian Rhythm and Jet Lag",
		"Aeronautical Decision Making",
		"Crew Resource Management CRM",
		"Error Management",
		"Team Communication",
		"Human Factors in Accidents",
		"Workload and Prioritization",
		"Hazardous Attitudes",
	},
	"operational_procedures": {
		"Introduction to Operational Procedures",
		"Flight Preparation",
		"Exterior Pre-flight Inspections",
		"Interior Pre-flight Inspections",
		"Engine Start and Startup",
		"Taxi and Ground Procedures",
		"Pre-takeoff Engine Run-up",
		"Normal Takeoff Procedure",
		"Initial Climb and Departure",
		"Enroute Navigation",
		"Level Changes",
		"Descent and Approach",
		"Normal Landing Procedure",
		"After Landing Taxi",
		"Engine Shutdown and Parking",
		"Rain Weather Procedures",
		"Strong Wind Procedures",
		"Night Procedures",
		"Engine Failure in Flight",
		"Engine Failure on Takeoff",
		"Electrical Failures",
		"Instrument Failures",
		"In-flight Fire Procedures",
		"Forced and Emergency Landings",
		"Operations on Unprepared Surfaces",
	},
	"principles_of_flight": {
		"Introduction to Aerodynamics",
		"Airfoil and Its Characteristics",
		"Lift Generation and Factors",
		"Parasite and Induced Drag",
		"Lift and Drag Coefficients",
		"Aerodynamic Center",
		"Center of Gravity and Stability",
		"Longitudinal Stability",
		"Lateral and Directional Stability",
		"Primary Flight Controls",
		"Secondary Flight Controls",
		"Stall and Recovery",
		"Spin and Recovery",
		"Forces in Steady Flight",
		"Turns and Load Factor",
		"Flight Envelope and Limits",
		"Crosswind Effects",
		"High Speed Aerodynamics",
		"High Lift Devices",
		"Special Aerodynamic Phenomena",
	},
	"communications": {
		"Introduction to Aeronautical Communications",
		"Standard Phraseology",
		"ICAO International Alphabet",
		"Aeronautical Frequencies",
		"VFR Procedures in Uncontrolled Airspace",
		"VFR Procedures in Controlled Airspace",
		"IFR Procedures",
		"Mayday Distress Communications",
		"Pan-Pan Urgency Communications",
		"In-flight Communications",
		"Ground Communications",
		"Information Service Communications",
		"English Language Communications",
		"Radio Failure Procedures",
		"Special Communications",
	},
	"mass_and_balance": {
		"Introduction to Mass and Balance",
		"Weight Definitions Empty and Load",
		"MTOW MLW and Other Limitations",
		"Center of Gravity and Lever Arm",
		"Moment Calculation",
		"Center of Gravity Determination",
		"Center of Gravity Envelopes",
		"Mass and Balance Graphs",
		"Aircraft Loading",
		"CG Effects on Performance",
		"CG Effects on Stability",
		"Practical Mass and Balance Calculations",
		"Mass and Balance Regulations",
		"Special Loading Cases",
		"Mass and Balance Calculation Tools",
	},
	"instrumentation": {
		"Introduction to Flight Instruments",
		"Pitot Static System",
		"Airspeed Indicator ASI",
		"Altimeter Principle and Setting",
		"Vertical Speed Indicator VSI",
		"Attitude Indicator",
		"Turn Coordinator",
		"Magnetic Compass",
		"Directional Gyro",
		"Gyroscopic Instruments",
		"ADF Radio Compass",
		"VOR Indicator",
		"ILS Localizer and Glide",
		"Transponder Modes A C S",
		"GPS and Navigation Systems",
		"Engine Instruments",
		"Warning and Alert Systems",
		"Fuel Management Instruments",
		"Automatic Systems Autopilot",
		"Instrument Failures and Procedures",
	},
}

func main() {
	qPerLesson := 10
	if len(os.Args) > 1 {
		fmt.Sscanf(os.Args[1], "%d", &qPerLesson)
	}
	outputPath := "seed/seed_complete.sql"
	if len(os.Args) > 2 {
		outputPath = os.Args[2]
	}

	var sb strings.Builder
	sb.WriteString("DELETE FROM student_question_history;\n")
	sb.WriteString("DELETE FROM questions;\n")
	sb.WriteString("DELETE FROM lessons;\n\n")

	lessonID := 1
	questionID := 1

	for _, lic := range licenses {
		for _, cat := range categories {
			numLessons := lessonsPerCategory[cat]
			titlesFr := lessonTitlesFr[cat]
			titlesEn := lessonTitlesEn[cat]
			for lessonNum := 1; lessonNum <= numLessons; lessonNum++ {
				lid := fmt.Sprintf("a%07x-0000-4000-8000-%012x", lessonID, lessonID)
				diff := lessonNum % 5
				if diff == 0 {
					diff = 5
				}
				if lic == "ATPL" || lic == "CPL" || lic == "IR" {
					diff = diff + 2
					if diff > 5 {
						diff = 5
					}
				}
				titleFr := titlesFr[lessonNum-1]
				titleEn := titlesEn[lessonNum-1]
				sb.WriteString(fmt.Sprintf("INSERT INTO lessons (id,license,category,theme,title_fr,title_en,content_fr,content_en,difficulty,order_index) VALUES\n"))
				contentFr := generateLessonContentFr(cat, lessonNum)
				contentEn := generateLessonContentEn(cat, lessonNum)
				// Échapper les apostrophes pour SQL
				contentFr = strings.ReplaceAll(contentFr, "'", "''")
				contentEn = strings.ReplaceAll(contentEn, "'", "''")
				sb.WriteString(fmt.Sprintf("('%s','%s','%s','%s','%s','%s','%s','%s',%d,%d);\n\n",
					lid, lic, cat, cat, titleFr, titleEn, contentFr, contentEn, diff, lessonNum))

				var qs []QuestionData
				switch cat {
				case "airlaw":
					qs = genAirLaw(qPerLesson)
				case "meteorology":
					qs = genMeteo(qPerLesson)
				case "navigation":
					qs = genNavigation(qPerLesson)
				case "performance":
					qs = genPerformance(qPerLesson)
				case "aircraft_general":
					qs = genAircraftGeneral(qPerLesson)
				case "flight_planning":
					qs = genFlightPlanning(qPerLesson)
				case "human_performance":
					qs = genHumanPerformance(qPerLesson)
				case "operational_procedures":
					qs = genOperationalProcedures(qPerLesson)
				case "principles_of_flight":
					qs = genPrinciplesOfFlight(qPerLesson)
				case "communications":
					qs = genCommunications(qPerLesson)
				case "mass_and_balance":
					qs = genMassAndBalance(qPerLesson)
				case "instrumentation":
					qs = genInstrumentation(qPerLesson)
				}

				if len(qs) > 0 {
					sb.WriteString(fmt.Sprintf("-- Questions for lesson %s\n", lid))
					sb.WriteString("INSERT INTO questions (id,lesson_id,license,category,theme,subtopic,difficulty,question_fr,question_en,options,answer_key,explanation_fr,explanation_en) VALUES\n")
					for i, q := range qs {
				qid := fmt.Sprintf("b%07x-0000-4000-8000-%012x", questionID, questionID)
						comma := ";"
						if i < len(qs)-1 {
							comma = ","
						}
						sb.WriteString(fmt.Sprintf("('%s','%s','%s','%s','%s','%s',%d,'%s','%s','%s','%s','%s','%s')%s\n",
							qid, lid, lic, cat, cat, q.subtopic, q.difficulty, escapeSQL(q.fr), escapeSQL(q.en), escapeSQL(q.options), escapeSQL(q.answer), escapeSQL(q.explainFr), escapeSQL(q.explainEn), comma))

						questionID++
					}
					sb.WriteString("\n")
				}
				lessonID++
			}

			// Add category exam (after the 2 lessons of this category)
			lidCat := fmt.Sprintf("a%07x-0000-4000-8000-%012x", lessonID, lessonID)
			sb.WriteString(fmt.Sprintf("INSERT INTO lessons (id,license,category,theme,title_fr,title_en,content_fr,content_en,difficulty,order_index) VALUES\n"))
			sb.WriteString(fmt.Sprintf("('%s','%s','%s','exam_cat','Examen %s (%s)','%s exam (%s)','Examen de fin de chapitre %s pour %s.','End of chapter exam %s for %s.',%d,98);\n\n",
				lidCat, lic, cat, cat, lic, cat, lic, cat, lic, cat, lic, 5))

			// Generate category exam questions (mix of all subtopics from this category)
			var catQs []QuestionData
			switch cat {
			case "airlaw":
				catQs = genAirLaw(20)
			case "meteorology":
				catQs = genMeteo(20)
			case "navigation":
				catQs = genNavigation(20)
			case "performance":
				catQs = genPerformance(20)
			case "aircraft_general":
				catQs = genAircraftGeneral(20)
			case "flight_planning":
				catQs = genFlightPlanning(20)
			case "human_performance":
				catQs = genHumanPerformance(20)
			case "operational_procedures":
				catQs = genOperationalProcedures(20)
			case "principles_of_flight":
				catQs = genPrinciplesOfFlight(20)
			case "communications":
				catQs = genCommunications(20)
			case "mass_and_balance":
				catQs = genMassAndBalance(20)
			case "instrumentation":
				catQs = genInstrumentation(20)
			}

			if len(catQs) > 0 {
				sb.WriteString(fmt.Sprintf("-- Category exam questions for %s - %s\n", lic, cat))
				sb.WriteString("INSERT INTO questions (id,lesson_id,license,category,theme,subtopic,difficulty,question_fr,question_en,options,answer_key,explanation_fr,explanation_en) VALUES\n")
				for i, q := range catQs {
				qid := fmt.Sprintf("b%07x-0000-4000-8000-%012x", questionID, questionID)
					comma := ";"
					if i < len(catQs)-1 {
						comma = ","
					}
					sb.WriteString(fmt.Sprintf("('%s','%s','%s','%s','%s','%s',%d,'%s','%s','%s','%s','%s','%s')%s\n",
						qid, lidCat, lic, cat, "exam_cat", q.subtopic, q.difficulty, escapeSQL(q.fr), escapeSQL(q.en), escapeSQL(q.options), escapeSQL(q.answer), escapeSQL(q.explainFr), escapeSQL(q.explainEn), comma))

					questionID++
				}
				sb.WriteString("\n")
			}
			lessonID++
		}

		// Add final exam for this license
		lid := fmt.Sprintf("a%07x-0000-4000-8000-%012x", lessonID, lessonID)
		sb.WriteString(fmt.Sprintf("INSERT INTO lessons (id,license,category,theme,title_fr,title_en,content_fr,content_en,difficulty,order_index) VALUES\n"))
		sb.WriteString(fmt.Sprintf("('%s','%s','exam','exam','Examen final %s','Final exam %s','Examen final couvrant toutes les matieres %s.','Final exam covering all subjects %s.',%d,99);\n\n",
			lid, lic, lic, lic, lic, lic, 5))

		examQs := genFinalExamQuestions(lic, 40)
		if len(examQs) > 0 {
			sb.WriteString(fmt.Sprintf("-- Final exam questions for %s\n", lic))
			sb.WriteString("INSERT INTO questions (id,lesson_id,license,category,theme,subtopic,difficulty,question_fr,question_en,options,answer_key,explanation_fr,explanation_en) VALUES\n")
			for i, q := range examQs {
			qid := fmt.Sprintf("b%07x-0000-4000-8000-%012x", questionID, questionID)
				comma := ";"
				if i < len(examQs)-1 {
					comma = ","
				}
				sb.WriteString(fmt.Sprintf("('%s','%s','%s','%s','%s','%s',%d,'%s','%s','%s','%s','%s','%s')%s\n",
					qid, lid, lic, q.subtopic, "exam", q.subtopic, q.difficulty, escapeSQL(q.fr), escapeSQL(q.en), escapeSQL(q.options), escapeSQL(q.answer), escapeSQL(q.explainFr), escapeSQL(q.explainEn), comma))

				questionID++
			}
			sb.WriteString("\n")
		}
		lessonID++
	}

	os.MkdirAll("scripts/seed", 0755)
	f, err := os.Create(outputPath)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()
	f.WriteString(sb.String())
	fmt.Printf("Seed generated: %s\n", outputPath)
	fmt.Printf("Questions per lesson: %d\n", qPerLesson)
	fmt.Printf("Total lessons: %d\n", lessonID-1)
	fmt.Printf("Total questions: %d\n", questionID-1)
}
