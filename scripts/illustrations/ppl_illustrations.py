"""Create editable, self-contained SVG diagrams and attach them to PPL lessons.

Standard library only. Run from any directory with Python 3.
The original flight-controls drawing is retained. No browser or image service is used.
"""
import json
import math
import re
import textwrap
from html import escape
from pathlib import Path
import xml.etree.ElementTree as ET

from svg_common import SVG, ROOT, ASSETS, CYAN, AMBER, PURPLE, WHITE, MUTED, RED, PANEL


def cards(s, spec):
    for i,(icon,title,body) in enumerate(spec['cards']):
        x=60+i*365; color=[CYAN,AMBER,PURPLE][i]
        s.rect(x,208,350,368,'#192d45','#324963',14)
        s.icon(icon,x+58,274,color)
        s.text(x+305,241,f'0{i+1}',15,MUTED,anchor='end')
        y=s.wrap(x+24,357,title,300,26,color,34,700)
        s.wrap(x+24,y+20,body,296,21,line=32)
        if spec.get('flow') and i<2: s.arrow(x+353,397,x+363,397,color,2)


def airframe(s, spec):
    s.plane(430,388,1.7)
    for x1,y1,x2,y2,label in [(280,409,112,306,'AILE / LONGERON'),(430,310,650,269,'FUSELAGE'),(451,494,660,545,'EMPENNAGE'),(435,404,130,520,'TRAIN SOUS LE FUSELAGE')]:
        s.circle(x1,y1,5,CYAN,CYAN); s.line(x1,y1,x2,y2,CYAN,2)
        s.text(x2,y2-12,label,16,CYAN,anchor='middle')
    s.rail([('Aile','Produit l’essentiel de la portance.'),('Fuselage','Accueille occupants et équipements.'),('Empennage','Contribue à la stabilité et au contrôle.')])


def engine(s, spec):
    names=['ADMISSION','COMPRESSION','DÉTENTE','ÉCHAPPEMENT']
    positions=[435,345,435,345]
    for i,(name,piston) in enumerate(zip(names,positions)):
        x=85+i*280; color=[CYAN,AMBER,RED,PURPLE][i]
        s.rect(x,275,188,240,'#0b192b','#536e8d',12)
        s.rect(x+12,285,164,piston-285,color,radius=4)
        s.rect(x+12,piston,164,28,'#d0deec',radius=4)
        s.line(x+94,piston+28,x+94,532,'#d0deec',12)
        s.text(x+94,240,name,19,color,anchor='middle',weight=700)
        s.arrow(x+210,355 if i%2==0 else 465,x+210,465 if i%2==0 else 355,color,3)
        s.text(x+94,566,['Le mélange entre','Le mélange se comprime','La combustion pousse','Les gaz sortent'][i],17,MUTED,anchor='middle')


def circuits(s, spec):
    rows=[('CARBURANT','Réservoir','Sélecteur / pompe','Moteur',CYAN),('ÉLECTRICITÉ','Alternateur + batterie','Réseau de bord','Équipements',AMBER),('HYDRAULIQUE','Commande de frein','Pression du fluide','Freins',PURPLE)]
    for i,(name,a,b,c,col) in enumerate(rows):
        y=245+i*120; s.text(65,y,name,17,col,weight=700)
        for j,text in enumerate([a,b,c]):
            x=270+j*288; s.rect(x,y-32,244,69,'#20354f',col,10); s.wrap(x+15,y-5,text,215,20,WHITE,25)
            if j<2:s.arrow(x+250,y,x+280,y,col,3)


def pitot(s, spec):
    blocked=spec.get('blocked',False)
    s.rect(72,251,205,78,'#193d4b',CYAN,12);s.text(96,285,'PRESSION TOTALE',18,CYAN,weight=700);s.text(96,312,'Tube Pitot',18)
    s.rect(72,433,205,78,'#353147',RED if blocked else PURPLE,12);s.text(96,466,'PRESSION STATIQUE',17,PURPLE,weight=700);s.text(96,493,'Source '+('bloquée' if blocked else 'commune'),18)
    s.line(277,290,533,290,CYAN,4)
    s.line(277,472,400,472,PURPLE,4);s.line(400,317,400,550,PURPLE,4)
    for y,label in [(290,'ANÉMOMÈTRE'),(422,'ALTIMÈTRE'),(550,'VARIOMÈTRE')]:
        s.line(400,y+27 if y==290 else y,533,y+27 if y==290 else y,PURPLE,4)
        s.rect(533,y-30,245,65,'#1c314a','#52718c',12);s.text(655,y+10,label,19,WHITE,anchor='middle',weight=700)
    if blocked:
        s.line(298,450,335,494,RED,6);s.line(335,450,298,494,RED,6)
        s.rail([('Source commune','Plusieurs indications peuvent devenir douteuses ensemble.'),('Altimètre','Indication approximativement figée.'),('Variomètre','Retour vers zéro après le transitoire.')])
    else:s.rail([('Vitesse indiquée','Pression totale moins pression statique.'),('Altitude','Pression statique + calage choisi.'),('Vitesse verticale','Variation de la pression statique.')])


def runway(s, spec):
    s.rect(86,464,667,44,'#35465b',radius=5)
    for x in range(100,740,65):s.rect(x,483,35,5,'#bdcbdc')
    s.path('M112 458H422Q510 444 715 291',stroke=CYAN,width=5)
    s.plane(593,385,.42,56)
    s.rect(703,404,20,60,'#af6673');s.path('M688 404L713 356 737 404Z',RED,RED)
    s.arrow(115,550,421,550,AMBER);s.text(268,578,'Roulement',20,AMBER,anchor='middle')
    s.arrow(115,228,718,228,CYAN);s.text(414,211,'Distance jusqu’à la hauteur de référence',19,CYAN,anchor='middle')
    s.rail([('Conditions','Masse, température, altitude-pression, vent.'),('Piste','Surface, pente et distance disponible.'),('Deux distances','Décoller du sol ≠ franchir un obstacle.')])


def atmosphere(s, spec):
    for i,(y,col,label) in enumerate([(220,'#1c2c58','STRATOSPHÈRE'),(333,'#203e61','TROPOPAUSE'),(377,'#235472','TROPOSPHÈRE')]):
        h=[113,44,185][i];s.rect(95,y,690,h,col);s.text(117,y+30,label,18,WHITE,weight=700)
    for row,count in [(270,5),(415,12),(465,17),(525,23)]:
        for j in range(count):s.circle(335+j*18,row,2,'#80cddd','none')
    s.cloud(455,443,.75);s.cloud(645,495,.5)
    s.arrow(70,558,70,228,CYAN,3)
    s.text(105,590,'Densité généralement plus forte vers le sol',18,MUTED)
    s.rail([('Référence ISA','Au niveau de la mer : 15 °C et 1 013,25 hPa.'),('Troposphère','La plupart des phénomènes météo du pilote léger.'),('Modèle ≠ prévision','Les limites et les températures réelles varient.')])


def pressure(s, spec):
    s.rect(85,550,710,35,'#1c5267');s.text(100,576,'NIVEAU MOYEN DE LA MER',16,CYAN)
    s.path('M90 548L195 428H394L485 548Z','#384d53','#57716d',2)
    s.rect(219,419,110,10,'#cad7e5');s.plane(576,308,.4,90)
    s.line(330,306,756,306,MUTED,1,True)
    s.arrow(475,546,475,312,CYAN);s.text(450,412,'QNH',24,CYAN,anchor='end',weight=700)
    s.arrow(674,426,674,312,AMBER);s.text(695,380,'QFE',24,AMBER,weight=700)
    s.line(338,426,755,426,MUTED,1,True)
    s.text(440,228,'Une indication dépend de son calage',24,WHITE,anchor='middle')
    s.rail([('QNH','Altitude référencée au niveau moyen de la mer.'),('QFE','Hauteur par rapport à la référence du terrain.'),('1 013,25 hPa','Référence des niveaux de vol ; transition publiée.')])


def humidity(s, spec):
    s.text(435,240,'TEMPÉRATURE ET POINT DE ROSÉE',23,WHITE,anchor='middle',weight=700)
    s.line(170,449,720,449,'#566c87',5)
    for val,x in [(8,170),(9,280),(10,390),(11,500),(12,610),(13,720)]:
        s.line(x,435,x,462,MUTED,2);s.text(x,492,f'{val} °C',20,MUTED,anchor='middle')
    for x,col,label in [(500,PURPLE,'POINT DE ROSÉE'),(610,AMBER,'AIR')]:
        s.line(x,321,x,436,col,4);s.circle(x,443,10,col,col);s.text(x,303,label,18,col,anchor='middle',weight=700)
    s.line(500,527,610,527,CYAN,4);s.text(555,566,'Écart : 1 °C',24,CYAN,anchor='middle',weight=700)
    s.rail([('Observation','Air à 12 °C ; point de rosée à 11 °C.'),('Interprétation','La saturation est proche.'),('À recouper','Un faible écart ne prouve pas la présence de brouillard.')])


def wind(s, spec):
    s.line(155,441,750,441,'#46637f',30)
    for x in range(170,740,60):s.line(x,441,x+30,441,WHITE,3)
    s.arrow(162,520,748,520,MUTED,2);s.text(450,553,'AXE 090° — références identiques',19,MUTED,anchor='middle')
    s.arrow(688,278,429,428,CYAN,6)
    s.arrow(688,278,429,278,AMBER,3);s.arrow(429,278,429,428,PURPLE,3)
    s.text(539,242,'Vent du 060° / 20 kt',23,CYAN,anchor='middle',weight=700)
    s.text(555,307,'≈ 17 kt de face',20,AMBER,anchor='middle')
    s.text(244,367,'10 kt traversiers',20,PURPLE)
    s.rail([('Direction','Le vent est nommé par sa provenance.'),('Écart à l’axe : 30°','20 × sin 30° = 10 kt traversiers.'),('En pratique','Comparer les références vrai / magnétique et les limites.')])


def clouds(s, spec):
    for i,(label,body) in enumerate([('COUCHES','Stratiforme'),('BOURGEONS','Cumuliforme'),('DÉVELOPPEMENT','Cumulonimbus')]):
        x=245+i*360;s.text(x,242,label,18,[CYAN,AMBER,PURPLE][i],anchor='middle',weight=700)
        if i==0:
            s.cloud(x-62,373,.8);s.cloud(x+52,374,.8)
        elif i==1:s.cloud(x,375,1.3)
        else:
            s.cloud(x,405,1.1);s.cloud(x-10,343,.9);s.cloud(x,288,.7)
            s.path(f'M{x+12} 438l-30 45h25l-20 44',stroke=AMBER,width=5)
        s.text(x,538,body,25,WHITE,anchor='middle',weight=700)
    s.text(600,592,'La forme renseigne sur les mouvements de l’air ; elle ne suffit pas à décider du passage.',18,MUTED,anchor='middle')


def visibility(s, spec):
    for i,(title,opacity) in enumerate([('BON CONTRASTE',1),('BRUME',.45),('FAIBLE CONTRASTE',.14)]):
        x=62+i*365;s.rect(x,260,344,260,'#496575',radius=10)
        s.path(f'M{x} 493L{x+70} 363L{x+145} 448L{x+220} 318L{x+344} 494Z','#152b40','#152b40',1,opacity)
        s.rect(x+149,454,35,40,f'rgba(20,40,60,{opacity})')
        s.text(x+172,235,title,18,[CYAN,AMBER,PURPLE][i],anchor='middle',weight=700)
    s.text(600,572,'Les mêmes repères deviennent plus difficiles à reconnaître.',24,WHITE,anchor='middle')


def front(s, spec):
    s.rect(92,224,694,331,'#553d4a',radius=10)
    s.path('M92 555V375Q300 378 660 242L786 224V555Z','#39596e','#698da4',2)
    s.arrow(217,329,618,224,AMBER,5);s.cloud(459,298,.8);s.cloud(651,251,.6)
    s.text(242,477,'AIR FROID',28,CYAN,weight=700)
    s.text(180,275,'AIR CHAUD',28,AMBER,weight=700)
    s.text(454,587,'Coupe verticale — soulèvement au-dessus de l’air froid',17,MUTED,anchor='middle')
    s.rail([('Rencontre','Des masses d’air de propriétés différentes.'),('Soulèvement','L’air chaud monte et peut se condenser.'),('Évolution','Nuages, précipitations et chronologie à surveiller.')])


def hazards(s, spec):
    s.cloud(249,318,1.05);s.path('M263 365l-40 63h33l-30 60',stroke=AMBER,width=7)
    s.text(249,552,'CONVECTION',22,AMBER,anchor='middle',weight=700)
    s.path('M461 418Q477 333 705 417Q615 436 461 418Z','#d3e4ee','#a5c6d8',3)
    for x,y in [(465,394),(478,373),(498,367)]:s.circle(x,y,10,CYAN,CYAN)
    s.text(590,552,'GIVRAGE',22,CYAN,anchor='middle',weight=700)
    s.arrow(858,316,1070,316,PURPLE,7);s.arrow(997,393,861,393,RED,7);s.plane(1030,445,.33,90)
    s.text(958,552,'CISAILLEMENT',22,PURPLE,anchor='middle',weight=700)
    s.text(600,592,'Identifier et éviter l’exposition avant de perdre ses marges.',18,MUTED,anchor='middle')


def metar(s, spec):
    values=[('151200Z','15 du mois, 12 h UTC'),('24012KT','Vent du 240° vrai, 12 kt'),('9999','Visibilité ≥ 10 km'),('SCT025','Nuages épars, base 2 500 ft'),('18/10','Température / rosée : °C'),('Q1015','QNH : 1 015 hPa')]
    for i,(a,b) in enumerate(values):
        x=61+(i%3)*365;y=218+(i//3)*181;col=[CYAN,AMBER,PURPLE][i%3]
        s.rect(x,y,350,158,'#1b314b','#35516e',12);s.text(x+22,y+51,a,31,col,weight=700);s.wrap(x+22,y+97,b,305,19,line=26)
    s.text(600,598,'Fragment pédagogique : la base nuageuse est référencée à l’aérodrome.',17,MUTED,anchor='middle')


def map_base(s):
    for x in range(90,801,80):s.line(x,210,x,570,'#253e54',1)
    for y in range(220,581,60):s.line(86,y,799,y,'#253e54',1)
    s.path('M84 550Q155 500 190 419T397 366T625 274T799 247',stroke='#3f91ad',width=12)
    s.path('M84 299L300 475 550 450 798 523',stroke='#6f7480',width=6)
    s.arrow(756,280,756,223,MUTED,2);s.text(756,210,'N',17,MUTED,anchor='middle')


def navigation(s,spec):
    mode=spec.get('mode','route')
    map_base(s)
    if mode=='scale':
        s.rect(286,309,191,128,'#384359',PURPLE,12);s.text(382,366,'ESPACE',21,PURPLE,anchor='middle');s.text(382,397,'3 dimensions',18,WHITE,anchor='middle')
        s.line(134,555,294,555,WHITE,5);s.line(134,545,134,565,WHITE,3);s.line(294,545,294,565,WHITE,3);s.text(214,589,'1 cm → 5 km',20,CYAN,anchor='middle')
        s.rail([('Échelle 1:500 000','Sur le document original : 1 cm représente 5 km.'),('Volume','Lire aussi le plancher et le plafond.'),('Mise à jour','Compléter avec les informations aéronautiques.')])
    elif mode=='fix':
        s.path('M95 385Q146 395 190 419',stroke='#3589b5',width=8)
        s.path('M550 365L550 540',stroke='#6f7480',width=6)
        for x,y,label in [(190,419,'CONFLUENCE'),(550,450,'CROISEMENT'),(625,274,'COUDE DE RIVIÈRE')]:
            s.circle(x,y,27,'none',AMBER,3);s.text(x,y-39,label,15,AMBER,anchor='middle')
        s.arrow(166,515,668,359,CYAN,3);s.plane(399,444,.34,73)
        s.rail([('Forme','Reconnaître un ensemble de détails.'),('Position relative','Comparer plusieurs repères.'),('Temps','Vérifier l’heure de passage attendue.')])
    elif mode=='detour':
        s.circle(425,382,94,'#493041',RED,2);s.text(425,373,'ZONE À',20,RED,anchor='middle');s.text(425,403,'ÉVITER',20,RED,anchor='middle')
        s.path('M158 501L315 532 617 478 720 300',stroke=CYAN,width=5)
        s.line(315,532,479,572,AMBER,3,True);s.circle(479,572,9,AMBER,AMBER)
        s.plane(203,510,.34,100);s.circle(720,300,9,CYAN,CYAN)
        s.rail([('Route principale','Relief, météo et espaces compatibles.'),('Décision anticipée','Évaluer la suite avant la zone difficile.'),('Solution de repli','Un terrain accessible et utilisable.')])
    else:
        s.arrow(143,498,718,287,CYAN,5);s.plane(396,412,.43,70)
        s.circle(143,498,8,AMBER,AMBER);s.circle(718,287,8,AMBER,AMBER)
        s.rect(279,505,359,77,'#1c334b','#3a5974',10);s.text(459,538,'30 NM ÷ 90 kt × 60',24,CYAN,anchor='middle');s.text(459,568,'20 minutes',22,WHITE,anchor='middle',weight=700)
        s.rail([('Position connue','Le point de départ de l’estime.'),('Direction + vitesse','Prévoir le déplacement au sol.'),('Contrôle','Confirmer les repères et actualiser le temps.')])


def triangle(s,spec):
    s.arrow(176,535,449,266,CYAN,5);s.arrow(449,266,704,321,AMBER,5);s.arrow(176,535,704,321,PURPLE,5)
    s.plane(317,395,.42,45)
    s.text(210,297,'VITESSE DANS L’AIR',18,CYAN,weight=700)
    s.text(586,235,'VENT',22,AMBER,weight=700)
    s.text(508,487,'VITESSE SOL',21,PURPLE,weight=700)
    s.rail([('Cap','Orientation de l’axe de l’avion.'),('Route','Direction du déplacement au sol.'),('Addition vectorielle','Vitesse dans l’air + vent = vitesse sol.')])


def radio_nav(s,spec):
    s.circle(246,391,111,'none','#3e5c78',2)
    s.circle(246,391,9,CYAN,CYAN);s.arrow(257,391,374,391,CYAN,4)
    s.text(246,237,'VOR',25,CYAN,anchor='middle',weight=700);s.text(246,559,'Radial 090 : à l’est',18,MUTED,anchor='middle')
    s.text(600,237,'NDB / ADF',25,AMBER,anchor='middle',weight=700)
    s.circle(653,329,9,AMBER,AMBER);s.plane(522,466,.43);s.arrow(540,441,644,346,AMBER,4);s.text(600,559,'Direction vers la balise',18,MUTED,anchor='middle')
    s.text(965,237,'DME',25,PURPLE,anchor='middle',weight=700)
    s.circle(850,487,8,PURPLE,PURPLE);s.plane(1060,321,.36,70);s.line(856,480,1043,330,PURPLE,4);s.line(850,491,1050,491,MUTED,2,True);s.line(1050,491,1050,337,MUTED,2,True)
    s.text(963,559,'Distance oblique',18,MUTED,anchor='middle')


def gps(s,spec):
    for x,y in [(164,262),(430,222),(697,262)]:
        s.rect(x-14,y-15,28,30,AMBER,radius=3);s.rect(x-47,y-10,27,20,'#405c88');s.rect(x+20,y-10,27,20,'#405c88');s.line(x,y+26,433,454,'#548098',2,True)
    s.plane(430,457,.72)
    s.rect(97,549,672,40,'#244258',radius=8);s.text(433,576,'POSITION ≠ BASE CARTOGRAPHIQUE À JOUR',18,CYAN,anchor='middle',weight=700)
    s.rail([('Signaux','Position et déplacement estimés.'),('Données de carte','Espaces, terrains et dates de validité.'),('Recoupement','Repères visuels, carte et estime.')])


def log(s,spec):
    headers=['DISTANCE','VITESSE SOL','TEMPS','DÉBIT','CONSOMMATION']
    vals=['45 NM','90 kt','30 min','24 L/h','12 L']
    for i,(h,v) in enumerate(zip(headers,vals)):
        x=65+i*216;s.rect(x,281,206,136,'#20364e','#3e5875',10);s.text(x+103,311,h,14,MUTED,anchor='middle',weight=700);s.text(x+103,378,v,31,[CYAN,AMBER,PURPLE,CYAN,AMBER][i],anchor='middle',weight=700)
    s.text(600,232,'UNE BRANCHE DE NAVIGATION',22,WHITE,anchor='middle',weight=700)
    s.text(335,492,'Temps = 60 × 45 / 90 = 30 min',23,CYAN,anchor='middle')
    s.text(880,492,'Carburant = 24 × 0,5 = 12 L',23,AMBER,anchor='middle')
    s.text(600,568,'Convertir les minutes en heures avant de multiplier par le débit horaire.',20,MUTED,anchor='middle')


def balance(s,spec):
    s.line(130,405,780,405,'#9ab1c8',8)
    for x,lab in [(130,'0 m'),(430,'1 m'),(730,'2 m')]:s.line(x,390,x,426,MUTED,2);s.text(x,458,lab,19,MUTED,anchor='middle')
    s.rect(345,268,170,128,'#235363',CYAN,12);s.text(430,331,'600 kg',30,CYAN,anchor='middle',weight=700)
    s.rect(680,328,100,68,'#574c38',AMBER,12);s.text(730,368,'100 kg',23,AMBER,anchor='middle',weight=700)
    s.path('M473 413L449 490H497Z',PURPLE,PURPLE,2);s.text(473,521,'CG ≈ 1,143 m',24,PURPLE,anchor='middle',weight=700)
    s.text(453,574,'(600 × 1 + 100 × 2) / 700 = 1,143 m',23,WHITE,anchor='middle')
    s.rail([('Moment','Masse × bras : kg·m.'),('Centrage','Somme des moments / masse totale.'),('Validation','Comparer le résultat à l’enveloppe du manuel.')])


def fuel(s,spec):
    s.rect(115,249,201,295,'#102033','#5e7c99',20);s.rect(129,270,173,260,'#236377',radius=12)
    s.text(215,375,'72 L',49,CYAN,anchor='middle',weight=700);s.text(215,411,'utilisables',21,WHITE,anchor='middle')
    s.text(385,387,'÷',54,MUTED,anchor='middle');s.text(536,374,'24 L/h',40,AMBER,anchor='middle',weight=700)
    s.text(536,418,'débit fictif constant',19,MUTED,anchor='middle');s.text(713,387,'=',45,MUTED,anchor='middle')
    s.rail([('3 heures','Autonomie théorique, pas durée de trajet autorisée.'),('À comptabiliser','Phases du vol, aléas et dégagement.'),('À préserver','Les réserves applicables au vol.')])


def performance(s,spec):
    s.arrow(141,537,766,537,MUTED,2);s.arrow(141,537,141,226,MUTED,2)
    s.text(158,215,'Distance (m)',18,MUTED);s.text(748,582,'Température (°C)',18,MUTED,anchor='end')
    for x,y,t,d,col in [(222,476,'10','480',CYAN),(442,376,'15','520',AMBER),(662,276,'20','560',CYAN)]:
        s.line(x,537,x,y,col,1,True);s.circle(x,y,8,col,col);s.text(x,564,t,18,col,anchor='middle');s.text(x,y-23,d+' m',22,col,anchor='middle',weight=700)
    s.line(222,476,662,276,CYAN,4)
    s.rail([('À mi-intervalle','À 15 °C : 480 + (560 − 480) / 2 = 520 m.'),('Correction fictive','520 × 1,15 = 598 m.'),('Attention au modèle','Interpolation dans la plage et selon la méthode du manuel.')])


def climb(s,spec):
    s.line(111,544,777,544,MUTED,2);s.line(111,544,111,229,MUTED,2)
    s.line(111,258,745,258,MUTED,1,True)
    s.arrow(111,544,534,258,CYAN,5);s.arrow(111,544,675,258,AMBER,5)
    s.line(534,258,534,544,CYAN,1,True);s.line(675,258,675,544,AMBER,1,True)
    s.text(109,225,'450 ft gagnés en 1 min',22,WHITE)
    s.text(442,348,'90 kt sol',23,CYAN,weight=700);s.text(542,419,'120 kt sol',23,AMBER,weight=700)
    s.text(534,578,'1,5 NM',19,CYAN,anchor='middle');s.text(697,578,'2 NM',19,AMBER,anchor='middle')
    s.rail([('Même taux','450 ft/min dans cet exercice.'),('À 90 kt sol','450 / 1,5 = 300 ft/NM.'),('À 120 kt sol','450 / 2 = 225 ft/NM : pente plus faible.')])


def airspeeds(s,spec):
    for i,(title,body,col) in enumerate([('IAS','Vitesse indiquée',CYAN),('CAS','Erreurs corrigées',AMBER),('TAS','Vitesse dans l’air',PURPLE),('SOL','Déplacement au sol',CYAN)]):
        x=60+i*280;s.circle(x+120,312,69,'#1c324c',col,3);s.text(x+120,322,title,35,col,anchor='middle',weight=700);s.text(x+120,425,body,20,WHITE,anchor='middle')
        if i<3:s.arrow(x+196,312,x+247,312,MUTED,2)
    for x,label in [(320,'Instrument / position'),(600,'Atmosphère*'),(880,'Vent')]:s.text(x,234,label,17,MUTED,anchor='middle')
    s.text(600,518,'110 kt dans l’air − 20 kt de face = 90 kt sol',28,CYAN,anchor='middle',weight=700)
    s.text(600,575,'* Chaîne simplifiée aux vitesses usuelles d’un avion école ; voir aussi EAS dans la leçon.',16,MUTED,anchor='middle')


def forces(s,spec):
    # Side silhouette, nose to the right; force directions are explicit.
    s.path('M289 406L366 394 416 365 451 365 484 397 604 406 618 420 288 420 258 377 277 377Z','#d8e6f0','#aac2d4',2)
    s.path('M405 408L340 446H370L459 409Z','#698ca8','#aac2d4',2)
    s.arrow(441,359,441,223,CYAN,5);s.text(468,244,'PORTANCE',21,CYAN,weight=700)
    s.arrow(441,436,441,577,AMBER,5);s.text(467,568,'POIDS',21,AMBER,weight=700)
    s.arrow(620,410,787,410,PURPLE,5);s.text(703,383,'TRACTION',18,PURPLE,anchor='middle',weight=700)
    s.arrow(281,410,105,410,RED,5);s.text(176,382,'TRAÎNÉE',18,RED,anchor='middle',weight=700)
    s.rail([('Cas représenté','Vol rectiligne horizontal stabilisé.'),('Équilibre vertical','Portance ≈ poids.'),('Équilibre horizontal','Traction ≈ traînée ; pas d’accélération.')])


def stall(s,spec):
    s.arrow(133,535,770,535,MUTED,2);s.arrow(133,535,133,221,MUTED,2)
    s.text(153,223,'Vitesse de décrochage',20,MUTED)
    points=[]
    for i in range(41):
        n=1+i/40;points.append((250+410*(n-1),485-7*(50*math.sqrt(n)-50)))
    s.path('M'+'L'.join(f'{x:.1f} {y:.1f}' for x,y in points),stroke=CYAN,width=5)
    for x,y,label,axis in [(250,485,'50 kt','1 g'),(660,340,'≈ 71 kt','2 g')]:
        s.circle(x,y,8,AMBER,AMBER);s.text(x,y-26,label,25,AMBER,anchor='middle',weight=700);s.text(x,570,axis,23,MUTED,anchor='middle')
    s.rail([('Même configuration','Masse et configuration fixées pour la comparaison.'),('Relation du modèle','Vs(n) = Vs(1 g) × √n.'),('Cause du décrochage','Dépassement de l’incidence critique, pas une vitesse universelle.')])


def diversion(s,spec):
    s.plane(195,401,.65,90)
    s.arrow(280,380,700,279,CYAN,4);s.arrow(280,433,742,535,AMBER,4)
    s.circle(702,279,13,CYAN,CYAN);s.circle(744,535,13,AMBER,AMBER)
    s.text(430,281,'A : 24 NM / 80 kt',23,CYAN,anchor='middle',weight=700)
    s.text(432,316,'18 minutes',23,WHITE,anchor='middle')
    s.text(515,429,'B : 30 NM / 120 kt',23,AMBER,anchor='middle',weight=700)
    s.text(515,462,'15 minutes',23,WHITE,anchor='middle')
    s.rail([('Distance ≠ temps','Le terrain plus loin peut être plus rapide à rejoindre.'),('Comparer aussi','Météo, relief, piste, vent et horaires.'),('Décider tôt','Conserver du temps et du carburant pour les options.')])


def feedback(s,spec):
    a,b,c=spec['labels']
    s.icon('radio',207,296,CYAN);s.icon('plane',991,296,AMBER)
    s.text(208,397,'SOL',25,CYAN,anchor='middle',weight=700);s.text(991,397,'PILOTE',25,AMBER,anchor='middle',weight=700)
    s.arrow(306,273,879,273,CYAN,4);s.text(594,239,a,21,CYAN,anchor='middle')
    s.arrow(880,405,307,405,AMBER,4);s.text(594,377,b,21,AMBER,anchor='middle')
    s.rect(304,486,579,74,'#26384d',PURPLE,12);s.text(594,532,c,21,PURPLE,anchor='middle',weight=700)


def attention(s,spec):
    for x,y,lab,col in [(183,249,'MÉTÉO',CYAN),(183,397,'NAVIGATION',AMBER),(183,545,'RADIO',PURPLE)]:
        s.rect(x-100,y-35,206,65,'#213950',col,10);s.text(x,y+7,lab,18,col,anchor='middle',weight=700);s.arrow(x+114,y,423,395,col,3)
    s.circle(504,395,78,'#25354e',WHITE,2);s.text(504,387,'ATTENTION',19,WHITE,anchor='middle',weight=700);s.text(504,416,'limitée',20,MUTED,anchor='middle')
    s.arrow(590,395,773,395,CYAN,4)
    s.rail([('Préparer','Réduire les tâches à improviser.'),('Prioriser','Préserver le contrôle de l’avion.'),('Recouper','Chercher aussi les indices qui contredisent son idée.')])


def airspace(s,spec):
    s.path('M168 327L360 222H742L550 327Z','#354963',PURPLE,2)
    s.path('M168 327H550V541H168Z','#263b55',CYAN,2)
    s.path('M550 327L742 222V436L550 541Z','#303e59',PURPLE,2)
    s.line(170,327,550,327,CYAN,4);s.line(170,541,550,541,AMBER,4)
    s.text(337,308,'PLAFOND',22,CYAN,anchor='middle',weight=700)
    s.text(337,580,'PLANCHER',22,AMBER,anchor='middle',weight=700)
    s.text(348,420,'VOLUME',30,WHITE,anchor='middle',weight=700)
    s.text(348,454,'limites latérales',21,MUTED,anchor='middle')
    s.rail([('Référence verticale','Altitude, hauteur ou niveau de vol : lire la légende.'),('Classe et statut','Déterminer les conditions d’accès.'),('Activité','Vérifier horaires et informations en vigueur.')])


def taxi(s,spec):
    s.rect(115,237,640,84,'#3a4d64',radius=5)
    for x in range(140,750,73):s.line(x,279,x+38,279,WHITE,4)
    s.rect(396,321,91,240,'#36485d')
    s.line(442,543,442,393,AMBER,3)
    for y in [359,368]:s.line(396,y,487,y,AMBER,3)
    for y in [338,347]:
        for x in range(398,486,20):s.line(x,y,x+10,y,AMBER,3)
    s.plane(442,478,.45)
    s.text(180,394,'POINT D’ATTENTE',18,AMBER,weight=700)
    s.text(442,597,'Rester du côté des lignes continues avant autorisation.',17,MUTED,anchor='middle')
    s.rail([('Roulage','Suivre le cheminement et la limite reçue.'),('Avant la piste','Vérifier l’autorisation nécessaire.'),('En cas de doute','Rester en attente et demander clarification.')])


def landing(s,spec):
    s.rect(458,508,333,38,'#354b63',radius=5)
    for x in range(477,785,50):s.line(x,526,x+28,526,WHITE,3)
    s.path('M110 249L517 503',stroke=CYAN,width=5)
    s.path('M341 393Q465 464 574 366L755 225',stroke=AMBER,width=5)
    s.plane(252,337,.38,122)
    s.circle(341,393,9,WHITE,WHITE)
    s.text(168,219,'APPROCHE',22,CYAN,weight=700)
    s.text(532,218,'REMISE DE GAZ',22,AMBER,weight=700)
    s.text(410,587,'Trajectoires de principe : aucune pente ni altitude prescrite.',17,MUTED,anchor='middle')
    s.rail([('Stabiliser','Vitesse, configuration, axe et énergie maîtrisés.'),('Décider','Appliquer les critères préparés avec l’instructeur.'),('Interrompre si besoin','La remise de gaz est une option normale.')])


def cg_envelope(s,spec):
    # Axis mapping: x = 100 + (arm - 0.95)*2400; y = 550 - (mass-700)*2.
    s.line(100,550,775,550);s.line(100,550,100,210)
    s.path('M220 550L700 550L700 250L340 250Z',fill='#254b51',stroke=CYAN,width=3)
    for x,label in [(220,'1,00'),(340,'1,05'),(700,'1,20')]:
        s.line(x,550,x,560);s.text(x,585,label,18,anchor='middle')
    for y,label in [(550,'700'),(250,'850')]:s.text(85,y+5,label,18,anchor='end')
    s.text(425,610,'Bras (m)',18,anchor='middle');s.text(110,210,'Masse (kg)',18)
    s.circle(450,290,7,AMBER,AMBER);s.text(465,288,'830 kg ; 1,096 m',18,AMBER)
    s.circle(415,350,7,PURPLE,PURPLE);s.text(430,378,'800 kg ; 1,081 m',18,PURPLE)
    s.rail([('Enveloppe fictive','Interpolation linéaire de la limite avant.'),('Départ','830 kg et 910 kg·m.'),('Après consommation','800 kg et 865 kg·m.')])


def wind_vectors(s,spec):
    s.arrow(520,550,460,256,CYAN,4)
    s.arrow(460,256,520,256,AMBER,4)
    s.arrow(520,550,520,256,PURPLE,4)
    s.text(140,345,'Air : 100 kt',25,CYAN);s.text(140,385,'Cap : 348,5° vrai',23,CYAN)
    s.text(490,220,'Vent du 270° : 20 kt',23,AMBER,anchor='middle')
    s.text(550,410,'Sol : ≈ 98 kt',23,PURPLE);s.text(550,448,'Route : 000° vrai',21,PURPLE)
    s.text(420,594,'Vecteurs à la même échelle ; vent vers l’est.',18,MUTED,anchor='middle')
    s.rail([('Référence','Toutes les directions sont vraies.'),('Correction','Composante ouest de 20 kt dans la vitesse air.'),('Résultat','La somme des vecteurs donne la vitesse sol.')])


def stability_curves(s,spec):
    for row,decay,label,color in [(0,True,'Oscillations amorties',CYAN),(1,False,'Oscillations divergentes',AMBER)]:
        y=320+row*190;s.line(105,y,775,y,MUTED,2,True)
        points=[]
        for i in range(301):
            t=i/300;amplitude=65*math.exp(-2.7*t) if decay else 15*math.exp(1.4*t)
            points.append(f'{110+650*t:.2f} {y-amplitude*math.cos(5*math.pi*t):.2f}')
        s.path('M'+'L'.join(points),stroke=color,width=3)
        s.text(110,y-90,label,23,color)
    s.text(440,600,'Temps →',20,anchor='middle')
    s.rail([('Statique','La tendance initiale après perturbation.'),('Dynamique','L’évolution de l’amplitude au fil du temps.'),('Interprétation','Un retour initial ne garantit pas l’amortissement.')])


RENDERERS = {fn.__name__:fn for fn in [cards,airframe,engine,circuits,pitot,runway,atmosphere,pressure,humidity,wind,clouds,visibility,front,hazards,metar,navigation,triangle,radio_nav,gps,log,balance,fuel,performance,climb,airspeeds,forces,stall,diversion,feedback,attention,airspace,taxi,landing,cg_envelope,wind_vectors,stability_curves]}

def main():
    specs=json.loads((Path(__file__).with_name('ppl_diagrams.json')).read_text(encoding='utf-8'))
    files=sorted((ROOT/'content/fr/ppl').rglob('*.json'))
    modules={json.loads(p.read_text(encoding='utf-8'))['metadata']['module_id'] for p in files}
    assert set(specs)==modules-{'PPL_AC_007'}, 'Every lesson needs an explicit illustration design'
    ASSETS.mkdir(parents=True,exist_ok=True)
    for path in files:
        lesson=json.loads(path.read_text(encoding='utf-8')); module=lesson['metadata']['module_id']
        if module=='PPL_AC_007':
            name='commandes-de-vol.svg';alt='Roulis, tangage et lacet : les axes de rotation et les gouvernes associées.'
        else:
            spec=specs[module];name=module.lower().replace('_','-')+'.svg'
            alt=spec['description']
            drawing=SVG(module,spec['title'],spec['subtitle'],alt)
            RENDERERS[spec['kind']](drawing,spec)
            result=drawing.finish();ET.fromstring(result)
            (ASSETS/name).write_text(result+'\n',encoding='utf-8')
        assert (ASSETS/name).exists()
        for concept in lesson['concepts']:
            concept['content']=re.sub(r'!\[[^\]\n]*\]\(/illustrations/[^)]+\.svg\)\s*','',concept['content']).strip()
        src='/illustrations/'+name
        lesson['concepts'][0]['content']=f'![{alt}]({src})\n\n'+lesson['concepts'][0]['content']
        lesson['metadata']['illustration']={'src':src,'alt':alt}
        path.write_text(json.dumps(lesson,ensure_ascii=False,indent=2)+'\n',encoding='utf-8')
    print(f'{len(files)} lessons illustrated; {len(specs)} SVGs generated, original flight-controls SVG retained.')

if __name__=='__main__':main()
