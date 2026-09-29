"""ATPL teaching diagrams."""
import sys
sys.dont_write_bytecode = True
from licence_diagrams import render, table
from svg_common import CYAN, AMBER, WHITE, MUTED


def draw(s, t, en):
    s.text(600,239,'600 NM',30,CYAN,anchor='middle')
    s.line(170,380,1030,380,WHITE,3)
    for x,label in [(170,'A'),(1030,'B')]:
        s.circle(x,380,10,CYAN,CYAN)
        s.text(x,430,label,30,anchor='middle')
    s.arrow(950,295,250,295,CYAN,3)
    s.text(600,280,t('Vers A : 180 kt sol','Towards A: 180 kt GS'),24,CYAN,anchor='middle')
    s.arrow(250,470,950,470,AMBER,3)
    s.text(600,510,t('Vers B : 120 kt sol','Towards B: 120 kt GS'),24,AMBER,anchor='middle')
    s.text(600,575,t('Vitesses constantes · Tous moteurs en fonctionnement','Constant speeds · All engines operating'),22,MUTED,anchor='middle')


def air_data(s, t, en):
    table(s, [t('Grandeur','Quantity'),t('Valeur','Value'),t('Relation','Relation')],
          [('EAS','200 kt','TAS = EAS / √σ'),('σ','0.64',t('Rapport de densité','Density ratio')),
           (t('Vitesse du son','Sound speed'),'600 kt','M = TAS / a')],
          t('Les exercices précisent la TAS ou l’EAS utilisée. Ne pas confondre vitesse air et sol.',
            'Exercises specify the TAS or EAS used. Do not confuse airspeed and groundspeed.'))


def turbine(s, t, en):
    labels=[t('Compresseur','Compressor'),t('Combustion','Combustion'),t('Turbine','Turbine'),t('Échappement','Exhaust')]
    for i,label in enumerate(labels):
        x=70+i*280
        s.rect(x,280,230,110,'#192d45','#324963',12)
        s.text(x+115,343,label,24,CYAN,anchor='middle')
        if i<3:s.arrow(x+232,335,x+275,335,AMBER,3)
    s.arrow(710,450,185,450,AMBER,3)
    s.text(450,490,t('Travail mécanique vers le compresseur','Shaft work to the compressor'),24,AMBER,anchor='middle')
    s.text(85,575,t('Modèle de cœur moteur ; architecture et limites propres au moteur.',
                         'Engine-core model; architecture and limits are engine-specific.'),21,MUTED)


def coordinates(s, t, en):
    table(s,[t('Point','Point'),t('Latitude','Latitude'),t('Longitude','Longitude')],
          [('A','45°00′ N','010°00′ E'),('B','47°30′ N','010°00′ E')],
          t('Même méridien. Modèle sphérique : 1 minute de latitude ≈ 1 NM.',
            'Same meridian. Spherical model: 1 minute of latitude ≈ 1 NM.'))
    s.arrow(875,510,875,430,CYAN,3)
    s.text(920,475,'N',26,CYAN)


def pressure(s, t, en):
    s.rect(110,245,610,260,'#203a50',stroke=CYAN,radius=35)
    s.text(415,300,t('CABINE','CABIN'),28,CYAN,anchor='middle')
    s.text(415,365,'75 kPa',42,anchor='middle')
    s.text(925,300,t('EXTÉRIEUR','OUTSIDE'),26,CYAN,anchor='middle')
    s.text(925,365,'30 kPa',42,anchor='middle')
    s.arrow(670,435,920,435,AMBER,4)
    s.text(805,478,t('Sortie régulée','Regulated outflow'),22,anchor='middle')
    s.text(110,555,t('Δp = pression cabine − pression extérieure','Δp = cabin pressure − outside pressure'),26)
    s.text(110,610,t('Valeurs fictives • Pressions exprimées dans la même unité','Fictional values • Pressures expressed in the same unit'),20,MUTED)


def main():
    render('ATPL', 'ATPL_031_001', ('Bilan de masse : données de départ', 'Mass balance: starting inputs'), mass_stages)
    from communications_diagram import draw as radio
    render('ATPL', 'ATPL_090_001', ('Communications : fermer la boucle', 'Communications: close the loop'), radio)
    render('ATPL', 'ATPL_021_002', ('Pressurisation : données de pression', 'Pressurisation: pressure inputs'), pressure)
    render('ATPL', 'ATPL_033_001', ('Route A–B : données de calcul', 'A–B route: calculation data'), draw)
    render('ATPL', 'ATPL_022_001', ('Vitesses air : données', 'Airspeeds: inputs'), air_data)
    render('ATPL', 'ATPL_021_001', ('Turbine : flux et énergie', 'Turbine: flow and energy'), turbine)
    render('ATPL', 'ATPL_061_001', ('Navigation : coordonnées', 'Navigation: coordinates'), coordinates)


def mass_stages(s, t, en):
    table(s, [t('Donnée','Input'),t('Masse','Mass'),t('Unité','Unit')],
          [('DOW','30 000' if not en else '30,000','kg'),
           (t('Charge marchande','Traffic load'),'8 000' if not en else '8,000','kg'),
           (t('Carburant bloc','Block fuel'),'10 000' if not en else '10,000','kg'),
           (t('Roulage','Taxi fuel'),'200','kg'),
           (t('Trajet après décollage','Trip after take-off'),'6 000' if not en else '6,000','kg')],
          t('Calculer ZFM, masse au parking, au décollage et à l’atterrissage. Données fictives.',
            'Calculate ZFM, ramp, take-off and landing mass. Fictional inputs.'))


if __name__ == '__main__':
    main()
