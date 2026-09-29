"""Bilingual diagrams for the individually revised teaching chapters."""
from licence_diagrams import render
from svg_common import CYAN, AMBER, WHITE, MUTED, PURPLE


def readback(s,t,en):
    s.text(65,210,t('Un échange fictif, trois contrôles','A fictional exchange, three checks'),23,CYAN)
    rows=[(t('Instruction','Instruction'),'F-GABC · 2 500 ft · QNH 1018'),
          (t('Réponse erronée','Incorrect reply'),'F-GABC · 3 500 ft · QNH 1018'),
          (t('Correction','Correction'),t('Niveau 2 500 ft ; QNH inchangé','Level 2,500 ft; QNH unchanged'))]
    for i,(a,b) in enumerate(rows):
        y=245+i*88;s.rect(65,y,1070,72,'#1c344b',radius=10)
        s.text(85,y+43,a,22,AMBER);s.text(395,y+43,b,23,WHITE)
    s.wrap(65,554,t('Vérifier : bon destinataire → chaque valeur → action et limite.',
                    'Check: correct addressee → each value → action and limit.'),1030,23,CYAN)


def failure(s,t,en):
    s.text(65,212,t('Signaux : choisir la bonne situation','Signals: choose the correct situation'),23,CYAN)
    for x,txt in [(65,t('Signal','Signal')),(430,t('En vol','In flight')),(820,t('Au sol','On ground'))]:s.text(x,256,txt,22,AMBER)
    rows=[(t('Vert continu','Steady green'),t('Atterrissage autorisé','Cleared to land'),t('Décollage autorisé','Cleared for take-off')),
          (t('Éclats verts','Green flashes'),t('Revenir pour atterrir*','Return for landing*'),t('Roulage autorisé','Cleared to taxi')),
          (t('Rouge continu','Steady red'),t('Céder, rester en circuit','Give way, keep circling'),t('Arrêter','Stop'))]
    for i,row in enumerate(rows):
        y=301+i*53
        for x,txt in zip([65,430,820],row):s.text(x,y,txt,21,WHITE)
        s.line(65,y+16,1125,y+16,'#304961',1)
    s.text(65,463,t('* Autorisation d’atterrir à suivre.','* Landing clearance to follow.'),20,MUTED)
    s.wrap(65,515,t('MAYDAY : danger et assistance immédiate. PAN PAN : sécurité concernée sans ce besoin immédiat.',
                    'MAYDAY: danger requiring immediate assistance. PAN PAN: safety concern without that immediate need.'),1020,22,CYAN,line=30)
    s.text(65,594,t('Signal → identité → problème → intentions → position / niveau / cap',
                    'Signal → identity → problem → intentions → position / level / heading'),20,AMBER)


def taf(s,t,en):
    s.text(65,209,'TAF LFFF 211000Z 2112/2121 24010KT 9999 SCT030',21,CYAN)
    s.text(65,243,'BECMG 2114/2116 5000 RA BKN012',21,AMBER)
    s.text(65,276,'TEMPO 2116/2120 2500 SHRA BKN006=',21,PURPLE)
    x=lambda h:90+(h-12)*112
    for h in [12,14,16,18,20,21]:
        s.line(x(h),314,x(h),500,'#304961',1,True);s.text(x(h),305,f'{h:02}:00',18,MUTED,anchor='middle')
    s.rect(x(12),332,x(14)-x(12),48,'#284354',CYAN,4)
    s.text(x(12)+12,362,'9999 SCT030',18,WHITE)
    s.rect(x(14),332,x(16)-x(14),48,'#4d4130',AMBER,4)
    s.text(x(14)+12,362,'BECMG',18,AMBER)
    s.rect(x(16),332,x(21)-x(16),48,'#284354',CYAN,4)
    s.text(x(16)+15,362,'5000 RA BKN012',21,WHITE)
    s.rect(x(16),416,x(20)-x(16),48,'#342c52',PURPLE,4)
    s.text(x(16)+15,446,'TEMPO 2500 SHRA BKN006',20,WHITE)
    s.wrap(65,532,t('Entre 14 et 16 h : transition sans instant exact. Entre 16 et 20 h : fluctuations, pas quatre heures continues.',
                    '1400–1600: transition without an exact change time. 1600–2000: fluctuations, not four continuous hours.'),1040,22,MUTED,line=30)
    s.text(65,601,t('Heures UTC · Vent 24010KT non modifié · Bulletin fictif',
                    'UTC times · Wind 24010KT unchanged · Fictional bulletin'),20,CYAN)


def professional_com(s,t,en):
    s.text(65,215,t('Dossier de dialogue : conserver les éléments non modifiés','Dialogue brief: retain unchanged items'),23,CYAN)
    rows=[('1','AER452',t('Cap 220 · 4 000 ft · 170 kt','Heading 220 · 4,000 ft · 170 kt')),
          ('2','AER425',t('Vitesse 160 kt','Speed 160 kt')),
          ('3','AER452',t('Cap 230 · maintenez 3 000 ft','Heading 230 · maintain 3,000 ft'))]
    for i,(n,who,msg) in enumerate(rows):
        y=255+90*i;s.rect(65,y,1070,72,'#1c344b',radius=10)
        s.text(85,y+43,n,23,AMBER);s.text(150,y+43,who,23,CYAN);s.text(385,y+43,msg,23,WHITE)
    s.wrap(65,555,t('Pour chaque ligne : à qui ? Quel paramètre change ? Que reste-t-il à appliquer ?',
                    'For each line: who is addressed? Which item changes? What remains in force?'),1020,23,AMBER)


def sensors(s,t,en):
    s.text(65,214,t('Installation fictive : deux réseaux distincts','Fictional installation: two distinct networks'),23,CYAN)
    # Tables plus explicit lines avoid ambiguous crossing junctions.
    for x,label in [(130,t('Affichage','Display')),(470,t('Pression totale','Total pressure')),(820,t('Pression statique','Static pressure'))]:s.text(x,266,label,22,AMBER)
    for i,(a,pt,ps) in enumerate([('A','Pt1','Ps1'),('B','Pt1','Ps2'),('C','Pt2','Ps2')]):
        y=318+i*72
        s.rect(100,y-30,180,52,'#1c344b',CYAN,8);s.text(190,y+5,a,25,WHITE,anchor='middle')
        s.text(550,y+5,pt,25,CYAN,anchor='middle');s.text(900,y+5,ps,25,PURPLE,anchor='middle')
        s.arrow(490,y-4,290,y-4,CYAN,2);s.line(730,y+19,300,y+19,PURPLE,2);s.line(730,y+19,730,y-5,PURPLE,2);s.arrow(850,y-5,730,y-5,PURPLE,2)
    s.wrap(65,551,t('Suivre séparément Pt et Ps. Une alimentation électrique séparée ne démontre pas une source de pression indépendante.',
                    'Trace Pt and Ps separately. Separate electrical power does not establish an independent pressure source.'),1050,22,MUTED,line=30)


def icing(s,t,en):
    s.text(65,214,t('Comparer les données avant de classer les durées','Compare inputs before ranking times'),23,CYAN)
    headers=[t('Option','Option'),'NM','GS (kt)',t('Altitude / minimum (ft)','Altitude / minimum (ft)'),t('Carburant (kg)','Fuel (kg)')]
    xs=[65,215,365,550,930]
    for x,txt in zip(xs,headers):s.text(x,267,txt,20,AMBER)
    for i,row in enumerate([['A','12','120','6 500 / 7 000','—'],['B','21','140','8 000 / 7 000','7'],['C','16','120','8 000 / 7 000','10']]):
        y=325+i*65
        for x,txt in zip(xs,row):s.text(x,y,txt,24,WHITE)
        s.line(65,y+20,1120,y+20,'#304961',1)
    s.wrap(65,515,t('8 kg disponibles pour ces segments ; besoins ultérieurs protégés. Profils fixes. Autres contraintes vérifiées.',
                    '8 kg available for these segments; later requirements protected. Fixed profiles. Other constraints checked.'),1040,22,CYAN,line=30)
    s.text(65,598,t('Calculer : temps (min) = 60 × distance (NM) / GS (kt)',
                    'Calculate: time (min) = 60 × distance (NM) / GS (kt)'),21,AMBER)


def envelope(s,t,en):
    # Numerically scaled axes. No quiz answer points appear in the prompt diagram.
    s.parts=[part.replace('Exercice fictif · Schéma non à l’échelle','Exercice fictif · Axes gradués').replace('Fictional exercise · Diagram not to scale','Fictional exercise · Scaled axes') for part in s.parts]
    x=lambda pct:130+(pct-10)*20
    y=lambda mass:540-(mass-8000)*.065
    s.text(65,211,t('Enveloppe fictive · Limites incluses','Fictional envelope · Boundaries included'),23,CYAN)
    for mass in [8000,9000,10000,11000,12000]:
        s.line(130,y(mass),830,y(mass),'#304961',1)
        s.text(116,y(mass)+6,f'{mass:,}'.replace(',',' '),17,MUTED,anchor='end')
    for pct in range(10,46,5):
        s.line(x(pct),270,x(pct),540,'#304961',1);s.text(x(pct),566,str(pct),18,MUTED,anchor='middle')
    s.path(f'M{x(15)} {y(8000)}L{x(25)} {y(12000)}L{x(35)} {y(12000)}L{x(35)} {y(8000)}Z','#234553',CYAN,3)
    s.text(485,413,t('Zone admise','Permitted region'),22,WHITE,anchor='middle')
    s.text(65,247,t('Masse (kg)','Mass (kg)'),18,AMBER)
    s.text(500,603,t('Centrage (%MAC)','CG (%MAC)'),20,AMBER,anchor='middle')
    s.wrap(865,294,t('Avant : 15 % à 8 000 kg ; 25 % à 12 000 kg. Interpolation linéaire.',
                    'Forward: 15% at 8,000 kg; 25% at 12,000 kg. Linear interpolation.'),260,20,WHITE,line=29)
    s.wrap(865,450,t('Arrière : 35 %. Vérifier chaque état de la séquence.',
                    'Aft: 35%. Check every state in the sequence.'),260,20,AMBER,line=29)


if __name__=='__main__':
    for program in ['PPL','LAPL']:
        for suffix,titles,draw in [
            ('COM_001',('Collationnement : repérer et corriger','Readback: detect and correct'),readback),
            ('COM_002',('Panne radio : signaux et messages','Radio failure: signals and messages'),failure),
            ('MET_009',('TAF : lire la chronologie','TAF: read the timeline'),taf)]:
            render(program,program+'_'+suffix,titles,draw)
    for program in ['CPL','ATPL','IR']:
        render(program,program+'_090_001',('Reconstruire la clairance','Reconstruct the clearance'),professional_com)
    render('CPL','CPL_022_001',('Capteurs : suivre chaque source','Sensors: trace each source'),sensors)
    render('IR','IR_050_001',('Givrage : trois issues à comparer','Icing: compare three exits'),icing)
    render('ATPL','ATPL_031_001',('Masse et centrage : lire l’enveloppe','Mass and balance: read the envelope'),envelope)
