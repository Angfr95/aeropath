"""Index the current PPL revision bank by examination subject.

This mapping is not a claim that every syllabus objective is covered. Only the
explicit four-choice bank enters practice papers; true/false stays in revision.
"""
import json
from collections import Counter
from pathlib import Path

ROOT=Path(__file__).resolve().parents[2]
SUBJECTS={
    'air_law':'airlaw','aircraft_general':'aircraft_general','instrumentation':'aircraft_general',
    'meteo':'meteorology','navigation':'navigation','human_performance':'human_performance',
    'communications':'communications','radio_procedure':'communications',
    'principles_of_flight':'principles_of_flight','operational_procedures':'operational_procedures',
    'emergency':'operational_procedures','performance':'flight_performance_planning',
    'flight_planning':'flight_performance_planning','mass_and_balance':'flight_performance_planning',
}
QUOTAS={'airlaw':20,'human_performance':12,'meteorology':20,'communications':8,'navigation':20,
        'principles_of_flight':12,'operational_procedures':12,'flight_performance_planning':12,'aircraft_general':12}

def main():
    counts=Counter(); rows=[]
    for path in sorted((ROOT/'content/fr/ppl').rglob('*.json')):
        d=json.loads(path.read_text(encoding='utf-8'));meta=d['metadata'];module=meta['module_id']
        subject=SUBJECTS[path.parent.name]
        # These two legacy chapters teach cross-cutting subjects within operations.
        if module=='PPL_OP_004':subject='communications'
        if module=='PPL_OP_006':subject='human_performance'
        meta['exam_subject']=subject
        meta['coverage_status']='partial'  # No automatic certification from a count.
        for c in d['concepts']:
            c['syllabus_reference']='AMC1 FCL.210; FCL.215 — '+subject
            c['learning_objective']=c['title']
        eligible=0
        for q in d['questions']:
            # The order must express truth value, never reveal correctness.
            # Apply the same permutation to every existing translation.
            if len(q['options']) == 2 and set(q['options']) == {'Vrai', 'Faux'}:
                order = [q['options'].index(label) for label in ['Vrai', 'Faux']]
                for field in ['options', 'options_en', 'options_es']:
                    if q.get(field):
                        assert len(q[field]) == 2, (module, field)
                        q[field] = [q[field][index] for index in order]
            tags=[t for t in q.get('tags',[]) if t!='exam-mcq' and not t.startswith('exam-subject:')]
            q['exam_subject']=subject
            q['syllabus_reference']='AMC1 FCL.210; FCL.215 — '+subject
            q['practice_format']='mcq' if len(q['options'])==4 else 'true_false_revision'
            q['source_references']=[s['url'] for s in meta.get('sources',[])]
            if len(q['options'])==4:
                tags+=['exam-mcq','exam-subject:'+subject];eligible+=1;counts[subject]+=1
            q['tags']=tags
        rows.append(dict(module=module,subject=subject,questions=len(d['questions']),mcq=eligible,
                         concepts=[dict(id=c['concept_id'],objective=c['title']) for c in d['concepts']]))
        path.write_text(json.dumps(d,ensure_ascii=False,indent=2)+'\n',encoding='utf-8')
    for subject,quota in QUOTAS.items():assert counts[subject]>=quota,(subject,counts[subject],quota)
    report={'scope':'PPL(A) France','status':'partial syllabus coverage; instructor review pending',
            'exam_bank':dict(counts),'papers':QUOTAS,'lessons':rows}
    (ROOT/'docs/ppl-syllabus-index.json').write_text(json.dumps(report,ensure_ascii=False,indent=2)+'\n',encoding='utf-8')
    print(f'{len(rows)} lessons, {sum(counts.values())} four-choice questions indexed; nine paper quotas available.')

if __name__=='__main__':main()
