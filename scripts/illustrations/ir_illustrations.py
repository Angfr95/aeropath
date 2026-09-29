"""IR teaching diagrams."""
import sys
sys.dont_write_bytecode = True
from licence_diagrams import render, table
from svg_common import CYAN, AMBER, WHITE, MUTED


def draw(s, t, en):
    s.line(210,520,900,520,WHITE,3)
    s.line(900,520,900,250,AMBER,3)
    s.line(210,520,900,250,CYAN,3)
    s.rect(875,495,25,25,'none',WHITE)
    s.plane(900,230,.28,90,WHITE)
    s.circle(210,520,9,CYAN,CYAN)
    s.text(150,566,'DME',23,CYAN)
    s.text(555,565,'d = 4 NM',26,WHITE,anchor='middle')
    s.text(928,395,'h = 3 NM',23,AMBER)
    s.text(490,350,'s = ?',30,CYAN)
    s.text(85,220,t('s : oblique   d : horizontale   h : hauteur relative','s: slant   d: horizontal   h: relative height'),21,MUTED)


def ils(s, t, en):
    s.line(170,510,1040,510,WHITE,3)
    s.line(170,510,1010,255,CYAN,3)
    s.line(1010,255,1010,510,AMBER,2)
    s.text(180,555,t('Point de référence','Reference point'),23)
    s.text(540,555,'5 NM',26,CYAN)
    s.text(1030,400,'h = ?',22,AMBER)
    s.text(540,300,'3°',28,CYAN)
    s.text(85,220,t('Modèle : 318 ft / NM horizontal','Model: 318 ft / horizontal NM'),24,CYAN)
    s.text(85,598,t('La hauteur relative calculée ne remplace pas les altitudes publiées.',
                         'Calculated relative height does not replace published altitudes.'),20,MUTED)


def altimetry(s, t, en):
    table(s,[t('Donnée','Input'),t('Calage','Setting'),t('Condition','Condition')],
          [(t('Avant','Before'),'1000 hPa',t('Pression statique fixe','Fixed static pressure')),
           (t('Après','After'),'1010 hPa',t('Même position','Same position'))],
          t('Modèle d’exercice près du niveau de la mer : 27 ft par hPa. Calculer la variation.',
            'Near-sea-level exercise model: 27 ft per hPa. Calculate the change.'))


def icing(s, t, en):
    table(s,[t('Situation','Situation'),t('Surveillance','Monitor'),t('Décision','Decision')],
          [(t('Eau surfondue','Supercooled water'),t('Dépôts / sondes','Deposits / probes'),t('Procédure avion','Aircraft procedure')),
           (t('Grosses gouttes','Large droplets'),t('Zones non protégées','Unprotected areas'),t('Domaine approuvé','Approved envelope')),
           (t('Accumulation','Accumulation'),t('Performances','Performance'),t('Issue compatible','Viable escape'))],
          t('Protection installée ≠ autorisation de voler dans toutes les conditions givrantes.',
            'Installed protection ≠ approval for every icing condition.'))


def flight_log(s, t, en):
    table(s, [t('Segment','Leg'),t('Distance','Distance'),t('Vitesse sol','Groundspeed')],
          [('A → B','120 NM','120 kt'),('B → C','90 NM','180 kt')],
          t('Calculer les temps avec la vitesse sol, puis le carburant avec le débit fourni.',
            'Calculate times using groundspeed, then fuel using the supplied flow.'))
    s.text(85,465,t('Débit constant : 40 L/h','Constant flow: 40 L/h'),26,CYAN)
    s.text(85,515,t('Départ A : 10:20 UTC','Departure A: 10:20 UTC'),26)


def main():
    render('IR', 'IR_010_001', ('Minima IFR : choisir la règle applicable', 'IFR minima: choose the applicable rule'), ifr_rules)
    render('IR', 'IR_040_001', ('IMC : sensations et références', 'IMC: sensations and references'), perception)
    from communications_diagram import draw as radio
    render('IR', 'IR_090_001', ('Communications : fermer la boucle', 'Communications: close the loop'), radio)
    render('IR', 'IR_033_001', ('Préparation IFR : journal de navigation', 'IFR planning: navigation log'), flight_log)
    render('IR', 'IR_062_001', ('DME : géométrie de la distance', 'DME: range geometry'), draw)
    render('IR', 'IR_062_002', ('ILS : géométrie du plan', 'ILS: glide-path geometry'), ils)
    render('IR', 'IR_022_001', ('Altimétrie : variation de calage', 'Altimetry: setting change'), altimetry)
    render('IR', 'IR_050_001', ('Givrage : observation et décision', 'Icing: observation and decision'), icing)


def ifr_rules(s, t, en):
    s.wrap(85,235,t('Hors besoin du décollage/atterrissage et autorisation spécifique de l’autorité.',
                       'Except as necessary for take-off/landing or specifically authorised by the authority.'),1020,23,line=30)
    for y,fr,eng in [(340,'Minimum établi par l’État → appliquer ce minimum','State-established minimum → apply that minimum'),
                     (410,'Sinon : montagne/relief élevé → obstacle + 2 000 ft','Otherwise: mountain/high terrain → obstacle + 2,000 ft'),
                     (480,'Sinon, ailleurs → obstacle + 1 000 ft','Otherwise, elsewhere → obstacle + 1,000 ft')]:
        s.text(85,y,t(fr,eng),24,CYAN)
    s.wrap(85,550,t('Règle subsidiaire : plus haut obstacle dans 8 km de la position estimée. Tenir compte de la précision de navigation.',
                    'Fallback rule: highest obstacle within 8 km of estimated position. Account for navigation accuracy.'),1020,20,line=27)


def perception(s, t, en):
    table(s,[t('Sensation','Sensation'),t('Piège','Trap'),t('Vérification','Cross-check')],
          [(t('Accélération','Acceleration'),t('Cabré apparent','Apparent pitch-up'),t('Attitude fiable','Reliable attitude')),
           (t('Virage prolongé','Prolonged turn'),t('Rotation non perçue','Unnoticed rotation'),t('Balayage instrumental','Instrument scan')),
           (t('Indications divergentes','Conflicting indications'),t('Source commune','Shared source'),t('Source indépendante','Independent source'))],
          t('Préserver la trajectoire ; recouper puis appliquer les procédures de l’avion.',
            'Preserve the flight path; cross-check and apply aircraft procedures.'))


if __name__ == '__main__':
    main()
