"""CPL teaching diagrams."""
import sys
sys.dont_write_bytecode = True
from licence_diagrams import render, table
from svg_common import CYAN, AMBER, WHITE, MUTED


def draw(s, t, en):
    s.text(95,238,t('Élément','Item'),23,CYAN)
    s.text(505,238,t('Masse','Mass'),23,CYAN)
    s.text(845,238,t('Bras depuis le repère','Arm from datum'),23,CYAN)
    for i,(fr,eng,mass,arm) in enumerate([
        ('Avion initial','Initial aircraft','1 800 kg','2,00 m'),
        ('Charge','Load','300 kg','3,00 m'),
        ('Carburant','Fuel','200 kg','2,50 m')]):
        y=310+i*75
        s.line(85,y+25,1105,y+25,'#304961',1)
        s.text(95,y,t(fr,eng),24)
        s.text(505,y,mass.replace('1 800','1,800') if en else mass,24)
        s.text(845,y,arm.replace(',','.') if en else arm,24,AMBER)
    s.text(95,560,t('Chaque bras utilise le même plan de référence.','Every arm uses the same datum.'),22,MUTED)


def performance(s, t, en):
    table(s, [t('Altitude-pression','Pressure altitude'), '10 °C', '30 °C'],
          [('0 ft', '500 m', '600 m'), ('2 000 ft' if not en else '2,000 ft', '650 m', '780 m')],
          t('Distance à 50 ft. Piste sèche, horizontale, sans vent. Interpolation linéaire autorisée.',
            'Distance to 50 ft. Dry, level runway, no wind. Linear interpolation allowed.'))
    s.text(85, 485, t('Utiliser l’altitude et la température de l’énoncé.', 'Use the altitude and temperature given in the question.'), 24, CYAN)


def load_factor(s, t, en):
    table(s, [t('Inclinaison','Bank'), 'cos φ', t('Hypothèse','Assumption')],
          [('0°','1',t('Palier','Level flight')),('45°','0.707',t('Coordonné','Coordinated')),('60°','0.5',t('Stabilisé','Steady'))],
          t('Modèle : n = 1 / cos φ ; Vs,n = Vs,1 × √n. Masse et configuration fixées.',
            'Model: n = 1 / cos φ; Vs,n = Vs,1 × √n. Fixed mass and configuration.'))
    s.text(85, 510, 'Vs,1 = 50 kt', 28, CYAN)


def fuel_plan(s, t, en):
    table(s, [t('Donnée','Input'),t('Valeur','Value'),t('Unité','Unit')],
          [(t('Distance','Distance'),'180','NM'),(t('Vitesse sol','Groundspeed'),'120','kt'),
           (t('Débit constant','Constant flow'),'40','L/h')],
          t('Calculer le temps et le carburant du segment. Autres allocations à traiter séparément.',
            'Calculate segment time and fuel. Account for other allocations separately.'))


def human_factors(s, t, en):
    table(s, [t('Situation','Situation'),t('Risque','Risk'),t('Barrière','Defence')],
          [(t('Météo dégradée','Poor weather'),t('Moins de marge','Reduced margins'),t('Options prévues','Planned options')),
           (t('Interruption','Interruption'),t('Omission','Omission'),t('Reprise vérifiée','Verified resumption')),
           (t('Plan initial','Initial plan'),t('Biais de poursuite','Continuation bias'),t('Seuils de décision','Decision thresholds')),
           (t('Automatisation','Automation'),t('Mode inattendu','Unexpected mode'),t('Surveiller l’effet','Monitor the result'))],
          t('Anticiper • Détecter • Corriger • Réévaluer les marges','Anticipate • Detect • Correct • Reassess margins'))


def main():
    render('CPL', 'CPL_022_001', ('Pitot-statique : sources et instruments', 'Pitot-static: sources and instruments'), pitot_static)
    from communications_diagram import draw as radio
    render('CPL', 'CPL_090_001', ('Communications : fermer la boucle', 'Communications: close the loop'), radio)
    render('CPL', 'CPL_040_001', ('Facteurs humains : barrières', 'Human factors: defences'), human_factors)
    render('CPL', 'CPL_031_001', ('Chargement : données de calcul', 'Loading: calculation data'), draw)
    render('CPL', 'CPL_032_001', ('Performances : interpolation', 'Performance: interpolation'), performance)
    render('CPL', 'CPL_081_001', ('Virage en palier : modèle', 'Level turn: model'), load_factor)
    render('CPL', 'CPL_033_001', ('Temps et carburant : données', 'Time and fuel: inputs'), fuel_plan)


def pitot_static(s, t, en):
    table(s, [t('Instrument','Instrument'),t('Source','Source'),t('Grandeur','Quantity')],
          [(t('Anémomètre','Airspeed indicator'),'Pt − Ps',t('Différence','Difference')),
           (t('Altimètre','Altimeter'),'Ps',t('Pression','Pressure')),
           (t('Variomètre','Vertical speed'),'Ps(t)',t('Variation','Change'))],
          t('Une source commune peut entraîner plusieurs indications erronées.',
            'A shared source may cause multiple incorrect indications.'))
    s.text(85,505,t('Pt : pression totale • Ps : pression statique','Pt: total pressure • Ps: static pressure'),25,CYAN)


if __name__ == '__main__':
    main()
