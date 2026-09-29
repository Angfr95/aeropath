"""Attach lesson diagrams to quiz corrections and neutral exercise SVGs to prompts.

Run before scripts/seed/ppl_content.py. Existing question IDs stay unchanged.
"""
import json
import re
from ppl_illustrations import ROOT, ASSETS, SVG, CYAN, AMBER, MUTED, WHITE

EXERCISES = {
    ('PPL_MB_001', 11): ('moment', 10, 2),
    ('PPL_MB_001', 12): ('moment', 30, 1),
    ('PPL_FP_001', 11): ('endurance', 48, 24),
    ('PPL_FP_001', 12): ('endurance', 60, 20),
    ('PPL_INS_001', 13): ('route', 60, 120),
    ('PPL_INS_001', 14): ('route', 30, 60),
    ('PPL_EMG_001', 11): ('route', 20, 80),
    ('PPL_EMG_001', 12): ('route', 40, 120),
    ('PPL_NAV_003', 0): ('route', 30, 90),
    ('PPL_NAV_003', 3): ('route', 40, 120),
    ('PPL_NAV_006', 0): ('route', 45, 90),
    ('PPL_NAV_006', 1): ('consumption', 30, 24),
    ('PPL_NAV_006', 3): ('consumption', 36, 24),
    ('PPL_FP_001', 0): ('endurance', 72, 24),
    ('PPL_FP_001', 3): ('consumption', 45, 24),
    ('PPL_MB_001', 3): ('moment', 20, 1.5),
    ('PPL_MET_005', 2): ('cloudbase', 800, 2000),
    ('PPL_MET_005', 3): ('cloudbase', 600, 1500),
}

def draw(module, kind, a, b):
    titles = {'route': 'Calculer un temps de parcours', 'consumption': 'Calculer une consommation',
              'endurance': 'Calculer une autonomie théorique', 'moment': 'Calculer le moment du bagage',
              'cloudbase': 'Distinguer hauteur et altitude'}
    if kind == 'route':
        alt = f'Branche de {a} NM entre A et B, vitesse sol constante de {b} kt. Durée à déterminer.'
    elif kind == 'cloudbase':
        alt = f'Terrain à {a} ft au-dessus du niveau de la mer. Base des nuages à {b} ft au-dessus du terrain. Altitude de la base à déterminer.'
    elif kind == 'moment':
        alt = f'Bagage de {a} kg à un bras de {str(b).replace(".", ",")} m de la référence. Moment à déterminer.'
    elif kind == 'endurance':
        alt = f'{a} L utilisables et débit fictif constant de {b} L/h. Autonomie théorique à déterminer.'
    else:
        alt = f'Durée de {a} minutes à un débit fictif constant de {b} L/h. Consommation à déterminer.'
    s = SVG(module, titles[kind], 'Données de l’exercice — choisissez votre réponse ci-dessous.', alt)
    if kind == 'route':
        s.circle(160,415,10,CYAN,CYAN); s.arrow(175,415,735,415,CYAN,5)
        s.text(160,465,'A',26,CYAN,anchor='middle'); s.text(735,465,'B',26,CYAN,anchor='middle')
        s.plane(440,385,.45,90)
        s.text(445,290,f'{a} NM',40,WHITE,anchor='middle',weight=700)
        s.text(445,535,f'Vitesse sol : {b} kt',26,AMBER,anchor='middle')
        s.rail([('Donnée 1','Distance au sol.'),('Donnée 2','Vitesse sol constante.'),('À déterminer','Temps en minutes : ?')])
    elif kind == 'cloudbase':
        s.rect(110,530,650,35,'#236377'); s.text(120,592,'Niveau moyen de la mer',20,MUTED)
        s.rect(160,430,360,100,'#385344'); s.text(175,465,'Terrain',20,WHITE)
        s.cloud(335,250,1.4); s.line(170,290,600,290,WHITE,2,True)
        s.line(555,430,555,290,CYAN,3); s.text(575,365,f'{b} ft',24,CYAN)
        s.line(555,530,555,430,AMBER,3); s.text(575,495,f'{a} ft',24,AMBER)
        s.rail([('Référence','Terrain et niveau de la mer sont distincts.'),('Données','Les valeurs sont indiquées sur les segments.'),('À déterminer','Altitude de la base : ?')])
    elif kind == 'moment':
        s.line(175,280,175,540,MUTED,3,True); s.text(175,255,'Référence',22,MUTED,anchor='middle')
        s.rect(530,320,160,115,'#294665',CYAN,12); s.text(610,390,f'{a} kg',35,WHITE,anchor='middle')
        s.line(175,485,610,485,AMBER,3); s.line(610,450,610,520,MUTED,2,True)
        s.text(390,535,f'Bras : {str(b).replace(".", ",")} m',25,AMBER,anchor='middle')
        s.rail([('Masse','Un seul bagage dans cet exercice.'),('Bras','Distance depuis la référence.'),('À déterminer','Moment en kg·m : ?')])
    else:
        s.icon('fuel',235,335,CYAN); s.icon('clock',610,335,AMBER)
        left = f'{a} L utilisables' if kind == 'endurance' else f'{b} L/h'
        right = f'{b} L/h' if kind == 'endurance' else f'{a} minutes'
        s.text(235,440,left,29,CYAN,anchor='middle',weight=700)
        s.text(610,440,right,29,AMBER,anchor='middle',weight=700)
        s.text(425,555,'Débit fictif constant',23,MUTED,anchor='middle')
        target = 'Durée en heures : ?' if kind == 'endurance' else 'Quantité en litres : ?'
        s.rail([('Exercice','Utiliser les valeurs affichées.'),('Unités','Vérifier leur cohérence.'),('À déterminer',target)])
    return s.finish(), alt

def clean(value):
    return re.sub(r'!\[[^\]\n]*\]\(/illustrations/[^)]+\.svg\)\s*', '', value).strip()

def main():
    seen = set(); questions = 0
    for path in sorted((ROOT/'content/fr/ppl').rglob('*.json')):
        data = json.loads(path.read_text(encoding='utf-8'))
        meta = data['metadata']; module = meta['module_id']; illustration = meta['illustration']
        assert (ASSETS/illustration['src'].removeprefix('/illustrations/')).is_file()
        for index, q in enumerate(data['questions']):
            q['question'] = clean(q['question'])
            q['explanation'] = clean(q['explanation'])
            q['explanation'] += f"\n\n![{illustration['alt']}]({illustration['src']})"
            key = (module, index)
            if key in EXERCISES:
                kind, a, b = EXERCISES[key]
                svg, alt = draw(module, kind, a, b)
                filename = f'quiz-{module.lower().replace("_", "-")}-{index+1}.svg'
                (ASSETS/filename).write_text(svg+'\n',encoding='utf-8')
                q['question'] += f'\n\n![{alt}](/illustrations/{filename})'
                if q.get('question_en'):
                    q['question_en'] = clean(q['question_en'])
                    q['question_en'] += f'\n\n![Exercise data: {meta["chapter_title_en"]}](/illustrations/en/{filename})'
                seen.add(key)
            questions += 1
        path.write_text(json.dumps(data,ensure_ascii=False,indent=2)+'\n',encoding='utf-8')
    assert seen == set(EXERCISES), 'An exercise no longer matches its lesson'
    print(f'{questions} illustrated corrections; {len(seen)} exercise SVGs attached to questions.')

if __name__ == '__main__': main()
