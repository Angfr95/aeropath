"""Correction diagrams for IR icing: causal effects and constrained choices."""
from licence_diagrams import render
from svg_common import CYAN, AMBER, WHITE, MUTED, RED


def effects(s, t, en):
    s.text(65,215,t('Un dépôt local peut modifier tout le comportement',
                    'A local deposit can change overall behaviour'),24,CYAN,weight=700)
    s.path('M110 365 C130 280 280 265 520 365 C300 400 140 400 110 365Z',
           fill='#254760',stroke=WHITE,width=3)
    s.path('M112 353 L100 338 L121 336 L112 319 L138 322 L137 306 L158 311 L167 298',
           stroke=AMBER,width=7)
    s.text(65,264,t('Dépôt / rugosité','Deposit / roughness'),21,AMBER)
    s.line(174,274,136,316,AMBER,2)
    s.text(80,438,t('Profil représentatif, sans échelle','Representative profile, not to scale'),18,MUTED)
    s.arrow(565,340,650,340,CYAN)
    for y,fr,eng in [(285,'Écoulement modifié','Altered airflow'),
                     (332,'Traînée possiblement accrue','Drag may increase'),
                     (379,'Portance max. peut diminuer','Maximum lift may decrease'),
                     (426,'Marges à réévaluer','Reassess margins')]:
        s.text(690,y,t(fr,eng),23,WHITE)
    s.wrap(65,500,t('IAS cohérente ≠ aile propre. Altitude stable ≠ marges préservées.',
                     'Consistent IAS ≠ clean wing. Stable altitude ≠ preserved margins.'),1060,22,AMBER,line=30)
    s.wrap(65,565,t('Recouper les indices ; appliquer les procédures propres à l’avion.',
                     'Cross-check cues; apply aircraft-specific procedures.'),1060,21,MUTED)


def exits(s, t, en):
    s.text(65,216,t('Q14 — profils fixes ; autres contraintes vérifiées pour B',
                    'Q14 — fixed profiles; other constraints verified for B'),22,CYAN)
    s.text(65,251,t('GS : 120 kt ; minimum applicable aux deux trajets : 7 000 ft',
                    'GS: 120 kt; applicable minimum on both routes: 7,000 ft'),22,WHITE)
    for x,title,distance,altitude,time,color,note in [
        (65,'A','8 NM',t('6 500 ft','6,500 ft'),'4 min',RED,t('Sous le minimum → exclue','Below minimum → reject')),
        (620,'B','24 NM',t('8 000 ft','8,000 ft'),'12 min',CYAN,t('Compatible avec les données','Meets the supplied inputs'))]:
        s.rect(x,280,515,190,'#1c344b','#304961',12)
        s.text(x+18,316,title,28,color,weight=700)
        s.text(x+75,316,distance+' · '+altitude,23,WHITE)
        s.text(x+18,366,t('Temps : ','Time: ')+time,23,WHITE)
        s.text(x+18,423,note,23,color)
    s.text(65,514,'t (min) = 60 × d (NM) / GS (kt)',25,AMBER)
    s.wrap(65,558,t('Éliminer les options incompatibles avant de comparer les temps.',
                     'Reject incompatible options before comparing times.'),1060,22,MUTED)


if __name__ == '__main__':
    render('IR','IR_ICING_EFFECTS',('Givrage : relier les effets','Icing: connect the effects'),effects)
    render('IR','IR_ICING_EXIT',('Givrage : comparer deux issues','Icing: compare two exits'),exits)
