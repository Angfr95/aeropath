"""Shared bilingual chapter authoring; licence-specific content lives in its section."""
import json
import re
import sys
sys.dont_write_bytecode = True
from .content import ROOT

AIRCREW = 'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no'
FORMAT = AIRCREW + '?kw=ppl&page=67'
REFERENCE = 'Appendix to AMC1 FCL.310; FCL.515(b); FCL.615(b)'


def save(data):
    meta = data['metadata']; module = meta['module_id']; program = meta['program']
    meta.update(content_status='draft', coverage_status='partial', reference_checked_on=meta.get('reference_checked_on','2026-09-26'),
                scope=program+'(A), France / EASA', quiz_format='chapter_practice', tags=[program, 'mcq-four-options'])
    src = '/illustrations/' + module.lower().replace('_', '-') + '.svg'
    meta['illustration'] = {'src': src, 'alt': meta['chapter_title']}
    for i, c in enumerate(data['concepts']):
        c['tags'] = [program]
        for field in ['content', 'content_en']:
            c[field] = re.sub(r'!\[[^\]]*\]\([^)]*\)\s*', '', c[field])
            if i == 0:
                link = src if field == 'content' else src.replace('/illustrations/', '/illustrations/en/')
                c[field] = f'![{module}]({link})\n\n' + c[field]
    for q in data['questions']:
        q['tags'] = [program, 'exam-mcq', 'exam-subject:'+meta['exam_subject']]
        q['practice_format'] = 'mcq'
        q['source_references'] = [s['url'] for s in meta['sources']]
        q['syllabus_reference'] = meta['syllabus_reference']
        for field in ['question', 'question_en', 'explanation', 'explanation_en']:
            q[field] = re.sub(r'!\[[^\]]*\]\([^)]*\)\s*', '', q[field]).strip()
            if field.startswith('explanation'):
                correction_src = q.get('explanation_illustration', src)
                link = correction_src if field == 'explanation' else correction_src.replace('/illustrations/', '/illustrations/en/')
                q[field] += f'\n\n![{module}]({link})'
    # The exercise drawing contains input data only; it does not reveal answers.
    for q in data['questions']:
        if q.get('use_exercise_svg'):
            for field in ['question','question_en']:
                link = src if field == 'question' else src.replace('/illustrations/', '/illustrations/en/')
                q[field] += f'\n\n![{module}]({link})'
    path = ROOT / 'content/fr' / program.lower() / meta['subject'] / (module+'_application.json')
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(json.dumps(data, ensure_ascii=False, indent=2)+'\n', encoding='utf-8')


def author(program, subject, code, title, title_en, concepts, rows, paper_count, paper_minutes, extra_sources=(), chapter=1, illustrated_question=True, reference_checked_on='2026-09-26', learning_objective=None, learning_objective_en=None, prerequisites=None, duration_minutes=35, explanation_illustrations=None):
    module = f'{program}_{code}_{chapter:03}'
    cs = [dict(concept_id=f'{module}_C{i+1:02}', title=frt, title_en=ent, content=fr, content_en=en,
               difficulty='advanced', syllabus_reference=REFERENCE+' — '+code) for i,(frt,ent,fr,en) in enumerate(concepts)]
    qs = []
    for i,line in enumerate(rows.strip().splitlines()):
        ci, fr, en, opts, opts_en, why, why_en = line.split('|')
        choices, english = opts.split('~'), opts_en.split('~')
        assert len(choices)==len(english)==4
        c=cs[int(ci)-1]; shift=i%4
        qs.append(dict(concept_id=c['concept_id'], concept_title=c['title'], level=3,
                       question=fr, question_en=en, options=choices[shift:]+choices[:shift],
                       options_en=english[shift:]+english[:shift], correct_answer=choices[0],
                       explanation=why, explanation_en=why_en, use_exercise_svg=(i==0 and illustrated_question)))
    assert len(qs)==15
    for number, illustration in (explanation_illustrations or {}).items():
        assert 1 <= number <= len(qs), (module, number)
        assert re.fullmatch(r'/illustrations/[a-z0-9-]+\.svg', illustration), illustration
        qs[number-1]['explanation_illustration'] = illustration
    meta=dict(module_id=module, program=program, subject=subject, chapter_title=title, chapter_title_en=title_en, reference_checked_on=reference_checked_on,
              learning_objective=learning_objective or 'Résoudre et vérifier les exercices de '+title.lower()+'.',
              learning_objective_en=learning_objective_en or 'Solve and cross-check exercises on '+title_en.lower()+'.',
              estimated_duration_minutes=duration_minutes, difficulty='advanced', prerequisites=prerequisites if prerequisites is not None else ['Calcul des proportions et des durées ; bases PPL du sujet'],
              exam_subject=code, syllabus_reference=REFERENCE+' — Subject '+code,
              official_subject_format=dict(questions=paper_count, minutes=paper_minutes, pass_percent=75, source=FORMAT),
              sources=[dict(title='EASA — programme Part-FCL',url=AIRCREW+'?page=8'),dict(title='EASA — format de la matière, AMC1 ARA.FCL.300(b)',url=FORMAT)])
    meta['sources'].extend(extra_sources)
    notice_fr=f'\n\n**Format de travail :** 15 QCM de chapitre, quatre choix et une réponse correcte. La matière {code} comporte {paper_count} questions en {paper_minutes} minutes dans le tableau EASA cité. Ce chapitre ne couvre qu’une partie de la matière ; il ne constitue pas une épreuve complète. Cadre EASA, même en anglais ou avec les unités US/britanniques.'
    notice_en=f'\n\n**Practice format:** 15 chapter MCQs, four choices and one correct answer. Subject {code} has {paper_count} questions in {paper_minutes} minutes in the cited EASA table. This chapter covers only part of the subject and is not a full paper. EASA framework, including with English or US/British units.'
    cs[-1]['content']+=notice_fr;cs[-1]['content_en']+=notice_en
    save(dict(metadata=meta,concepts=cs,questions=qs))
