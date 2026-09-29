"""LAPL navigation, adapted from the common PPL theoretical syllabus."""
import json
from shared.content import ROOT
from shared.authoring import AIRCREW, save

LAPL = AIRCREW + "?kw=LAPL&page=2"


def build():
    source=next((ROOT/'content/fr/ppl/navigation').glob('PPL_NAV_009*.json'))
    d=json.loads(source.read_text(encoding='utf-8'))
    d=json.loads(json.dumps(d,ensure_ascii=False).replace('PPL_NAV_009','LAPL_NAV_001'))
    m=d['metadata'];m.update(program='LAPL',module_id='LAPL_NAV_001',chapter_title='Navigation LAPL : vent et route',chapter_title_en='LAPL navigation: wind and track',
        exam_subject='navigation',syllabus_reference='AMC1 FCL.115; FCL.120 — PPL(A) navigation syllabus',
        derived_from='PPL_NAV_009',prerequisites=['Caps, routes et vitesse air'])
    m['sources'].append(dict(title='EASA — programme théorique commun LAPL/PPL',url=LAPL))
    m['sources'].append(dict(title='DGAC — examens théoriques LAPL/PPL',url='https://www.ecologie.gouv.fr/politiques-publiques/examens-theoriques-devenir-personnel-navigant'))
    d['concepts'][-1]['content']+='\n\nLe LAPL(A) utilise le programme théorique PPL(A). Ce chapitre commun est adapté au parcours LAPL. Les examens français comportent neuf épreuves, 128 questions au total et un seuil de 75 % par épreuve. Cette série de 15 questions est un quiz de chapitre, pas un examen complet.'
    d['concepts'][-1]['content_en']+='\n\nLAPL(A) uses the PPL(A) theoretical syllabus. This shared chapter is adapted to the LAPL pathway. French examinations comprise nine papers, 128 questions overall and a 75% pass mark per paper. This 15-question set is chapter practice, not a complete examination.'
    save(d)
