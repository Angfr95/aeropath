"""Write a per-licence inventory from the actual authored chapters."""
import json
from collections import defaultdict
from .content import ROOT, CATEGORIES, stable_id


def write_catalogue(program):
    sections=defaultdict(list)
    for path in sorted((ROOT/'content/fr'/program.lower()).rglob('*.json')):
        data=json.loads(path.read_text(encoding='utf-8'))
        meta=data['metadata']
        assert meta['program']==program
        module=meta['module_id']
        sections[meta['exam_subject']].append({
            'module':module,'title':meta['chapter_title'],'category':CATEGORIES[path.parent.name],
            'lesson_id':stable_id('lesson-'+program+'-'+module),
            'questions':len(data['questions']),
            'exam_mcq':sum('exam-mcq' in q['tags'] for q in data['questions']),
            'objectives':[c['title'] for c in data['concepts']],
            'official_subject_format':meta.get('official_subject_format'),
            'sources':meta['sources'],
        })
    chapters=[c for cs in sections.values() for c in cs]
    total=sum(c['questions'] for c in chapters);mcq=sum(c['exam_mcq'] for c in chapters)
    report=dict(program=program,coverage='partial',lessons=len(chapters),questions=total,
                exam_mcq=mcq,revision_questions=total-mcq,sections=dict(sections))
    directory=ROOT/'docs/licences'
    (directory/(program.lower()+'-syllabus-index.json')).write_text(json.dumps(report,ensure_ascii=False,indent=2)+'\n',encoding='utf-8')
    text=f'# {program}(A) — cours disponibles\n\n'
    text+=f'**{len(chapters)} chapitres, {total} questions**, en français et en anglais, avec SVG et variantes d’unités.\n\n'
    text+=f'{mcq} QCM à quatre choix sont marqués pour le format examen ; {total-mcq} questions sont réservées à la révision. Chaque chapitre comporte 15 questions.\n\n'
    if program=='LAPL':
        text+='Le parcours reprend les 63 chapitres du programme théorique commun PPL/LAPL, avec des identifiants et un suivi LAPL distincts. Les passages qui comparent les licences conservent leurs désignations et exigences propres. Les vrai/faux restent des exercices de révision.\n\n'
    text+='La couverture reste partielle : présence d’une matière ne signifie pas que tous ses objectifs EASA sont couverts. Les quiz de chapitre ne constituent pas des épreuves complètes. Validation instructeur à obtenir.\n\n'
    for section, items in sorted(sections.items()):
        text+=f'## Matière {section}\n\n'
        for c in items:
            url=f'http://localhost:3000/lessons/{program}/{c["category"]}/{c["lesson_id"]}'
            text+=f'### [{c["title"]}]({url})\n\n'
            text+=f'`{c["module"]}` — {c["questions"]} questions, dont {c["exam_mcq"]} à quatre choix.\n\n'
            text+='Objectifs : '+', '.join(c['objectives'])+'.\n\n'
        fmt=items[0]['official_subject_format']
        if fmt:
            text+=f'Format de la matière dans le [tableau EASA cité]({fmt["source"]}) : {fmt["questions"]} questions / {fmt["minutes"]} minutes, seuil {fmt["pass_percent"]} %. Le tableau public consulté affiche l’édition août 2023 ; les conditions d’inscription sont à vérifier auprès de l’autorité.\n\n'
    text+='## Sources\n\n'
    sources={s['url']:s['title'] for c in chapters for s in c['sources']}
    for url,title in sources.items():text+=f'- [{title}]({url})\n'
    text+=f'\n## Régénération\n\n```text\npython scripts/seed/{program.lower()}_content.py\npython scripts/illustrations/{program.lower()}_illustrations.py\nnode scripts/illustrations/ppl_unit_variants.mjs\n```\n\n'
    text+=f'Les chapitres sont rangés par matière dans `content/fr/{program.lower()}` et leurs auteurs dans `scripts/seed/{program.lower()}`. Import transactionnel : `scripts/seed/{program.lower()}_content.sql`. Le changement d’unités ou de langue ne modifie pas le cadre France / EASA.\n'
    (directory/(program.lower()+'.md')).write_text(text,encoding='utf-8')
