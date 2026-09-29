"""Reviewed CPL exercise inputs and fuel reasoning; run after general generators."""
from licence_diagrams import render
from cpl_illustrations import performance as worked_performance
from svg_common import CYAN, AMBER, WHITE, MUTED


def performance(s,t,en):
    s.text(65,214,t('Dossier autonome · Distances jusqu’à 50 ft (m)',
                    'Independent case · Distances to 50 ft (m)'),24,CYAN)
    for x,value in zip([85,560,910],[t('Altitude-pression','Pressure altitude'),'5 °C','25 °C']):
        s.text(x,276,value,24,AMBER)
    for y,row in [(340,('0 ft','600 m','720 m')),(416,(t('2 000 ft','2,000 ft'),'800 m','960 m'))]:
        for x,value in zip([85,560,910],row):s.text(x,y,value,26,WHITE)
        s.line(85,y+20,1105,y+20,'#304961',1)
    s.text(65,491,t('Cible et corrections : utiliser les données de l’énoncé.',
                    'Target and corrections: use the question inputs.'),23,CYAN)
    s.wrap(65,537,t('Interpolation linéaire autorisée. Masse et configuration fixées. Valeurs brutes : piste sèche, horizontale, sans vent.',
                    'Linear interpolation allowed. Fixed mass and configuration. Raw values: dry, level runway, no wind.'),1030,22,MUTED,line=30)


def fuel_options(s,t,en):
    s.text(65,216,t('Restant utilisable : 58 L · Protection à l’arrivée : 24 L',
                    'Usable remaining: 58 L · Arrival protection: 24 L'),25,CYAN)
    xs=[65,205,405,640,885]
    for x,value in zip(xs,[t('Option','Option'),t('Durée','Time'),t('Débit','Flow'),t('Approche*','Approach*'),t('Aptitude','Suitability')]):
        s.text(x,280,value,22,AMBER)
    for i,(name,time,fr,eng) in enumerate([('A','40 min','Approprié','Suitable'),('B','30 min','Piste inadaptée','Unsuitable runway'),('C','50 min','Approprié','Suitable')]):
        y=344+70*i
        for x,value in zip(xs,[name,time,'42 L/h','4 L',t(fr,eng)]):s.text(x,y,value,22,WHITE)
        s.line(65,y+19,1125,y+19,'#304961',1)
    s.text(65,544,t('* À ajouter : non incluse dans la durée indiquée.',
                    '* Add separately: excluded from the stated time.'),22,CYAN)
    s.text(65,587,t('Hypothèses fictives · Routes A et C également appropriées.',
                    'Fictional assumptions · Routes A and C also suitable.'),21,MUTED)


def fuel_method(s,t,en):
    rows=[(t('1 · Chaque phase','1 · Each phase'),t('NM ÷ kt = heures ; heures × L/h = litres','NM ÷ kt = hours; hours × L/h = litres')),
          (t('2 · À l’arrivée','2 · At arrival'),t('Restant utilisable − consommation future','Usable remaining − future consumption')),
          (t('3 · Marge','3 · Margin'),t('Carburant d’arrivée − quantité protégée','Arrival fuel − protected quantity')),
          (t('4 · Décision','4 · Decision'),t('Vérifier aussi route, météo, piste et performances','Also check route, weather, runway and performance'))]
    for i,(label,value) in enumerate(rows):
        y=203+i*87;s.rect(65,y,1070,73,'#1c344b',radius=9)
        s.text(84,y+29,label,21,AMBER);s.text(84,y+58,value,23,WHITE)
    s.text(65,594,t('Même instant de référence · Aucune phase comptée deux fois',
                    'Same reference instant · No phase counted twice'),21,CYAN)


if __name__=='__main__':
    for module,titles,draw in [
        ('CPL_032_001',('Performances : dossier de synthèse','Performance: synthesis exercise'),performance),
        ('CPL_PERFORMANCE_WORKED',('Performances : premier exemple','Performance: first worked example'),worked_performance),
        ('CPL_033_001',('Carburant : comparer trois options','Fuel: compare three options'),fuel_options),
        ('CPL_FUEL_METHOD',('Carburant : construire le raisonnement','Fuel: build the reasoning'),fuel_method),
    ]:render('CPL',module,titles,draw)
