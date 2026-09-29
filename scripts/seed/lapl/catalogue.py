"""LAPL chapters from the explicitly shared PPL(A) theoretical syllabus.

Keep factual mentions of PPL privileges intact: the licence-comparison chapter
must not turn PPL requirements into LAPL requirements through text replacement.
"""
import json
from shared.content import ROOT
from .navigation import build as build_navigation

SOURCE = 'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=LAPL&page=2'


def target_module(source):
    # Preserve the already published LAPL_NAV_001 and its question IDs.
    if source.startswith('PPL_NAV_'):
        number=int(source.rsplit('_',1)[1])
        return f'LAPL_NAV_{1 if number==9 else number+1:03}'
    return source.replace('PPL_', 'LAPL_', 1)


def build():
    build_navigation()
    for path in sorted((ROOT/'content/fr/ppl').rglob('*.json')):
        data=json.loads(path.read_text(encoding='utf-8'))
        meta=data['metadata']; original=meta['module_id']; module=target_module(original)
        if original=='PPL_NAV_009':continue
        old_src=meta['illustration']['src'];new_src='/illustrations/'+module.lower().replace('_','-')+'.svg'
        # Only structured IDs and tags change, never factual lesson text.
        meta.update(program='LAPL',module_id=module,derived_from=original,
                    scope='LAPL(A), France / EASA',reference_checked_on='2026-09-26',
                    quiz_format='chapter_practice',coverage_status='partial',
                    syllabus_reference='AMC1 FCL.115; FCL.120 — '+meta['exam_subject'])
        meta['tags']=['LAPL',path.parent.name]
        meta['sources'].append(dict(title='EASA — programme théorique commun LAPL/PPL',url=SOURCE))
        meta['illustration']['src']=new_src
        for concept in data['concepts']:
            concept['concept_id']=concept['concept_id'].replace(original,module)
            concept['tags']=['LAPL']
            concept['syllabus_reference']=meta['syllabus_reference']
        for q in data['questions']:
            q['concept_id']=q['concept_id'].replace(original,module)
            q['tags']=['LAPL']+[t for t in q['tags'] if t not in {'PPL','LAPL'}]
            q['syllabus_reference']=meta['syllabus_reference']
            q['source_references']=[s['url'] for s in meta['sources']]
        def images(text):
            text=text.replace(old_src,new_src).replace(old_src.replace('/illustrations/','/illustrations/en/'),new_src.replace('/illustrations/','/illustrations/en/'))
            return text
        for concept in data['concepts']:
            for field in ['content','content_en']:concept[field]=images(concept[field])
        for q in data['questions']:
            for field in ['question','question_en','explanation','explanation_en']:q[field]=images(q[field])
        data['concepts'][-1]['content']+='\n\n**Parcours LAPL(A).** Chapitre du programme théorique commun PPL/LAPL. Les comparaisons de licences gardent leurs intitulés propres. Les questions vrai/faux sont des exercices de révision ; seules les questions à quatre choix sont marquées pour un entraînement au format examen.'
        data['concepts'][-1]['content_en']+='\n\n**LAPL(A) pathway.** Chapter from the shared PPL/LAPL theoretical syllabus. Licence comparisons retain their own licence names. True/false questions are revision exercises; only four-choice questions are tagged for exam-format practice.'
        dest=ROOT/'content/fr/lapl'/path.parent.name/(module+'_application.json')
        dest.parent.mkdir(parents=True,exist_ok=True)
        dest.write_text(json.dumps(data,ensure_ascii=False,indent=2)+'\n',encoding='utf-8')
