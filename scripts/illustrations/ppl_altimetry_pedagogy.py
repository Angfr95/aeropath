"""Altimetry worked model and independent question inputs, FR/EN.

Run after general diagram generators; values are fictional training inputs.
"""
from licence_diagrams import render
from svg_common import CYAN,AMBER,WHITE,MUTED


def settings(s,t,en):
    s.text(65,213,t('Même avion · Même position · Même pression statique',
                    'Same aircraft · Same position · Same static pressure'),24,CYAN)
    for x,head,pressure,reading in [(65,t('Avant','Before'),'1005 hPa',t('1 500 ft','1,500 ft')),(650,t('Après','After'),'1015 hPa',t('1 770 ft','1,770 ft'))]:
        s.rect(x,246,480,202,'#1c344b',radius=12)
        s.text(x+28,284,head,25,AMBER)
        s.text(x+28,346,reading,38,WHITE,weight=700)
        s.text(x+28,405,t('Calage : ','Setting: ')+pressure,25,CYAN)
    s.arrow(560,345,630,345,AMBER,4)
    s.text(65,495,t('Modèle fourni : +10 hPa × 27 ft/hPa = +270 ft indiqués',
                    'Supplied model: +10 hPa × 27 ft/hPa = +270 ft indicated'),24,AMBER)
    s.wrap(65,546,t('L’avion ne monte pas parce que le nombre augmente. Cette approximation de basses couches ne corrige pas l’effet du froid.',
                    'The aircraft does not climb because the number rises. This low-level approximation does not correct cold-temperature error.'),1020,23,MUTED,line=30)


def clearance(s,t,en):
    s.text(65,213,t('Dossier autonome · Valeurs à exploiter',
                    'Independent case · Input values'),24,CYAN)
    rows=[(t('Indication lue','Observed indication'),t('3 100 ft','3,100 ft')),
          (t('Calage utilisé / QNH actuel','Setting used / current QNH'),'1022 hPa / 1010 hPa'),
          (t('Altitude de l’obstacle','Obstacle altitude'),t('2 300 ft AMSL','2,300 ft AMSL')),
          (t('Marge exigée par cet exercice','Clearance required by this exercise'),'500 ft')]
    for i,(a,b) in enumerate(rows):
        y=269+61*i;s.text(65,y,a,23,MUTED);s.text(755,y,b,24,WHITE)
        s.line(65,y+19,1130,y+19,'#304961',1)
    s.text(65,522,t('Approximation fournie : 27 ft/hPa','Supplied approximation: 27 ft/hPa'),24,AMBER)
    s.wrap(65,566,t('Modèle : altitude au QNH assimilée à l’altitude vraie ; autres erreurs négligées. Aucun résultat affiché.',
                    'Model: QNH altitude treated as true altitude; other errors ignored. No result displayed.'),1030,21,MUTED,line=28)


if __name__=='__main__':
    for program in ['PPL','LAPL']:
        render(program,program+'_MET_002',('Recalage : comprendre le sens','Setting changes: understand the direction'),settings)
    render('PPL','QUIZ_PPL_MET_002_CLEARANCE',('Altimétrie : vérifier une marge','Altimetry: check clearance'),clearance)
