"""Catalogue-wide editorial triage. Metrics never certify teaching quality.

Writes one record per chapter, with question indices for actionable review.
Run after the licence generators. No database, content or application changes.
"""
import json
import re
from collections import Counter, defaultdict
from datetime import date
from pathlib import Path

ROOT = Path(__file__).resolve().parents[2]
IMAGE = re.compile(r'!\[[^\]]*\]\(([^)]+)\)')
PROGRAMS = ('PPL', 'LAPL', 'CPL', 'ATPL', 'IR')
PRIORITY_LANGUAGES = ('fr', 'en')
DEFERRED_LANGUAGES = ('es',)


def plain(text):
    return IMAGE.sub('', text or '').strip()


def words(text):
    return len(re.findall(r"\b[\w]+(?:[’'-][\w]+)*\b", plain(text)))


def inspect(path):
    d = json.loads(path.read_text(encoding='utf-8'))
    m, cs, qs = d['metadata'], d['concepts'], d['questions']
    problems = []
    def flag(code, items=None):
        problems.append({'code': code, 'questions': items or []})
    if not m.get('sources'): flag('missing_sources')
    if (m.get('learning_objective', '').startswith('Résoudre et vérifier les exercices de')
            or not m.get('learning_objective')): flag('objective_needs_specific_outcomes')
    if m.get('prerequisites') == ['Calcul des proportions et des durées ; bases PPL du sujet']:
        flag('generic_prerequisites')
    fr_words = sum(words(c['content']) for c in cs)
    en_words = sum(words(c.get('content_en', '')) for c in cs)
    # Triage threshold, never a quality score or a minimum to pad towards.
    if fr_words < 500: flag('brief_content_review_depth')
    if any(not c.get('content_en') or not c.get('title_en') for c in cs): flag('missing_english_concept')
    spanish_concepts_missing = any(not c.get('content_es') or not c.get('title_es') for c in cs)
    used = {q['concept_id'] for q in qs}
    unassessed = [c['concept_id'] for c in cs if c['concept_id'] not in used]
    if unassessed: flag('concept_without_linked_question')
    short = []; dependent = []; untranslated = []; spanish = []; options = []; repeated = []
    seen = {}; positions = Counter(); tf = Counter(); broken = []; illustrations = []
    for i, q in enumerate(qs, 1):
        opts = q['options']
        if q.get('correct_answer') not in opts or len(set(opts)) != len(opts): options.append(i)
        else:
            positions[len(opts), opts.index(q['correct_answer'])] += 1
            if len(opts) == 2: tf[opts.index(q['correct_answer'])] += 1
        if min(words(q.get('explanation', '')), words(q.get('explanation_en', ''))) < 12: short.append(i)
        if re.search(r'\b(même (?:obstruction|trajet|situation|exemple)|précédent|ci-dessus|same (?:blockage|trip|situation|example)|previous|above example)\b', plain(q['question']) + ' ' + plain(q.get('question_en', '')), re.I): dependent.append(i)
        if not q.get('question_en') or not q.get('explanation_en') or len(q.get('options_en', [])) != len(opts): untranslated.append(i)
        if not q.get('question_es') or not q.get('explanation_es') or len(q.get('options_es', [])) != len(opts): spanish.append(i)
        normalized = ' '.join(plain(q['question']).lower().split())
        if normalized in seen: repeated.append(i)
        seen[normalized] = i
        illustrations += [i] if IMAGE.search(q['question']) else []
    for c in cs:
        for field in ('content', 'content_en', 'content_es'):
            for src in IMAGE.findall(c.get(field, '')):
                if src.startswith('/') and not (ROOT / 'web/frontend-react/public' / src.lstrip('/')).exists(): broken.append(src)
    for q in qs:
        for field in ('question', 'question_en', 'question_es', 'explanation', 'explanation_en', 'explanation_es'):
            for src in IMAGE.findall(q.get(field, '')):
                if src.startswith('/') and not (ROOT / 'web/frontend-react/public' / src.lstrip('/')).exists(): broken.append(src)
    for code, items in [('short_explanations_review_reasoning', short), ('possible_context_dependency', dependent),
                        ('missing_english_question_fields', untranslated),
                        ('invalid_options_or_answer', options), ('duplicate_stem_in_chapter', repeated)]:
        if items: flag(code, items)
    if broken: flag('missing_local_image')
    if sum(tf.values()) >= 8 and max(tf.values()) / sum(tf.values()) > .8: flag('true_false_position_bias')
    return {'module': m['module_id'], 'program': m['program'], 'subject': m.get('exam_subject', path.parent.name),
            'path': path.relative_to(ROOT).as_posix(), 'questions': len(qs), 'fr_words': fr_words, 'en_words': en_words,
            'estimated_minutes': m.get('estimated_duration_minutes'), 'content_status': m.get('content_status'),
            'coverage_status': m.get('coverage_status'), 'unassessed_concepts': unassessed,
            'illustrated_question_indices': illustrations,
            'deferred_translation_gaps': {'es': {'concepts_incomplete': spanish_concepts_missing, 'question_indices': spanish}},
            'true_false_positions': dict(tf), 'missing_images': sorted(set(broken)),
            'flags': problems, 'manual_pedagogical_validation': 'not_established'}


def main():
    records = [inspect(p) for program in PROGRAMS for p in sorted((ROOT/'content/fr'/program.lower()).rglob('*.json'))]
    groups = defaultdict(list)
    for r in records: groups[r['program']].append(r)
    counters = Counter(f['code'] for r in records for f in r['flags'])
    report = {'checked_on': str(date.today()), 'purpose': 'editorial_triage_not_certification',
              'priority_languages': PRIORITY_LANGUAGES, 'deferred_languages': DEFERRED_LANGUAGES,
              'thresholds': {'brief_content_words': 500, 'short_explanation_words': 12, 'position_bias_fraction': .8},
              'chapters': len(records), 'questions': sum(r['questions'] for r in records),
              'chapters_by_flag': dict(counters), 'records': records}
    dest = ROOT/'docs/licences'
    (dest/'course-quality-inventory.json').write_text(json.dumps(report, ensure_ascii=False, indent=2)+'\n', encoding='utf-8')
    lines = ['# Inventaire de qualité du catalogue', '', f'Contrôle du {report["checked_on"]} : {len(records)} chapitres, {report["questions"]} questions.', '',
             '**Ce relevé automatise le triage, pas la validation pédagogique.** Les seuils de longueur signalent des textes à relire ; ajouter des mots ne suffit pas à les améliorer. Les renvois de contexte sont des candidats à vérifier, pas des erreurs certaines.', '',
             'La relecture doit établir : exactitude et hypothèses ; objectifs précis ; progression et explications ; exemples résolus puis transfert autonome ; distracteurs plausibles et correction des erreurs ; pertinence et lisibilité des SVG ; équivalence des langues et unités ; couverture du programme public. Aucun score global ne remplace ces preuves.', '',
             'Priorité de travail : français et anglais. Les autres langues sont reportées à la demande de l’utilisateur ; les lacunes espagnoles restent dans le JSON, hors des alertes actives. Les durées restent estimées et les références de matière ne démontrent pas une couverture objectif par objectif.', '',
             '## Synthèse FR/EN', '', '| Parcours | Chapitres | Questions | Textes FR < 500 mots | Avec correction courte |', '| --- | ---: | ---: | ---: | ---: |']
    for program in PROGRAMS:
        rs = groups[program]
        count = lambda code: sum(any(f['code'] == code for f in r['flags']) for r in rs)
        lines.append(f'| {program} | {len(rs)} | {sum(r["questions"] for r in rs)} | {count("brief_content_review_depth")} | {count("short_explanations_review_reasoning")} |')
    lines += ['', '## Liste de travail par chapitre', '', 'Les numéros de questions concernés et les concepts sans question liée figurent dans [le relevé JSON](course-quality-inventory.json). `—` signifie aucune alerte automatisée, jamais validation complète.', '',
              '| Chapitre | Mots FR / EN | Alertes à examiner |', '| --- | ---: | --- |']
    for r in records:
        flags = ', '.join(f['code'] for f in r['flags']) or '—'
        lines.append(f'| [{r["module"]}](../../{r["path"]}) | {r["fr_words"]} / {r["en_words"]} | {flags} |')
    (dest/'course-quality-inventory.md').write_text('\n'.join(lines)+'\n', encoding='utf-8')
    print(json.dumps({k:v for k,v in report.items() if k != 'records'}, ensure_ascii=False))


if __name__ == '__main__': main()
