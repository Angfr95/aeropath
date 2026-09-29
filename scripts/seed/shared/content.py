"""Validate licence content and generate non-destructive SQL imports."""
import json
from pathlib import Path

ROOT = Path(__file__).resolve().parents[3]
CATEGORIES = {
    'air_law': 'airlaw', 'aircraft_general': 'aircraft_general',
    'meteo': 'meteorology', 'navigation': 'navigation',
    'operational_procedures': 'operational_procedures',
    'principles_of_flight': 'principles_of_flight',
    'human_performance': 'human_performance', 'communications': 'communications',
    'flight_planning': 'flight_planning', 'mass_and_balance': 'mass_and_balance',
    'performance': 'performance', 'instrumentation': 'instrumentation',
    'emergency': 'emergency', 'radio_procedure': 'radio_procedure',
}

def stable_id(seed):
    # Preserve the identifiers produced by scripts/seed-from-json/main.go.
    value = 0
    for char in seed:
        value = (value * 31 + ord(char)) & 0x7fffffff
    return f'{value:08x}-0000-4000-8000-{value:012x}'

def sql(value):
    if isinstance(value, (list, dict)):
        value = json.dumps(value, ensure_ascii=False)
    if isinstance(value, int):
        return str(value)
    return "'" + str(value).replace("'", "''") + "'"

def upsert(table, fields):
    names = list(fields)
    return (f'INSERT INTO {table} ({", ".join(names)}) VALUES\n('
            + ', '.join(sql(fields[k]) for k in names)
            + ')\nON CONFLICT (id) DO UPDATE SET '
            + ', '.join(f'{k} = EXCLUDED.{k}' for k in names if k != 'id') + ';\n')

def generate(program='PPL'):
    assert program in {'PPL', 'LAPL', 'CPL', 'ATPL', 'IR'}, program
    statements = [f'-- {program} content only. Preserves accounts, progress and answer history.\nBEGIN;\nSET LOCAL standard_conforming_strings = on;\n']
    ids, modules, question_count = set(), [], 0
    for path in sorted((ROOT / 'content/fr' / program.lower()).rglob('*.json')):
        data = json.loads(path.read_text(encoding='utf-8'))
        meta, concepts, questions = data['metadata'], data['concepts'], data['questions']
        module = meta['module_id']
        assert meta['program'] == program, path
        assert concepts and len(questions) >= 15, f'At least fifteen questions required: {path}'
        assert len({q['question'].strip().casefold() for q in questions}) == len(questions), f'Duplicate question text: {path}'
        category = CATEGORIES[path.parent.name]
        lesson_id = stable_id('lesson-' + program + '-' + module)
        assert lesson_id not in ids, f'Duplicate lesson: {module}'
        ids.add(lesson_id)
        concept_ids = {c['concept_id'] for c in concepts}
        assert len(concept_ids) == len(concepts), module
        assert all(c.get('content', '').strip() for c in concepts), module
        content = '\n\n'.join('## ' + c['title'] + '\n\n' + c['content'] for c in concepts)
        if meta.get('learning_objective'):
            content = '**Objectif :** ' + meta['learning_objective'] + '\n\n' + content
        sources = meta.get('sources', [])
        if sources:
            content += '\n\n## Références\n\n' + '\n'.join('- [' + s['title'] + '](' + s['url'] + ')' for s in sources)
        if meta.get('content_status') == 'draft':
            content += '\n\n**Support de révision PPL(A).** À faire valider par votre instructeur. Pour un vol réel, utiliser le manuel de l’avion et les publications en vigueur. Les exemples chiffrés sont pédagogiques.'
        content_en = '\n\n'.join('## ' + c.get('title_en', '') + '\n\n' + c['content_en'] for c in concepts if c.get('content_en'))
        assert all(c.get('title_en') and c.get('content_en') for c in concepts), (module, 'missing English concept')
        if meta.get('learning_objective_en'):
            content_en = '**Objective:** ' + meta['learning_objective_en'] + '\n\n' + content_en
        if sources:
            content_en += '\n\n## References\n\n' + '\n'.join('- [Reference ' + str(i+1) + '](' + s['url'] + ')' for i,s in enumerate(sources))
        if meta.get('content_status') == 'draft':
            content_en += '\n\n**PPL(A) revision material.** To be reviewed by your instructor. For an actual flight, use the aircraft manual and current publications. Numerical examples are for training.'
        difficulty = {'beginner': 1, 'intermediate': 2, 'advanced': 3}.get(meta['difficulty'], 1)
        if program != 'PPL':
            content = content.replace('**Support de révision PPL(A).**', '**Support de révision ' + program + '(A).**')
            content_en = content_en.replace('**PPL(A) revision material.**', '**' + program + '(A) revision material.**')
        statements.append(upsert('lessons', dict(id=lesson_id, license=program, category=category, theme=category, title_fr=meta['chapter_title'], title_en=meta['chapter_title_en'], content_fr=content, content_en=content_en, difficulty=difficulty, order_index=int(module.rsplit('_', 1)[1]), level=1, duration_minutes=meta['estimated_duration_minutes'], tags=meta.get('tags', []), learning_objectives=[meta['learning_objective']] if meta.get('learning_objective') else [])))
        for index, q in enumerate(questions):
            assert q['concept_id'] in concept_ids, module
            assert 2 <= len(q['options']) <= 4 and len(set(q['options'])) == len(q['options']), (module, index)
            if program != 'PPL':
                revision = program == 'LAPL' and q.get('practice_format') == 'true_false_revision' and len(q['options']) == 2 and 'exam-mcq' not in q.get('tags', [])
                assert (len(q['options']) == 4 or revision) and q.get('syllabus_reference') and q.get('source_references'), (module, index, 'exam format and references required')
            assert q['correct_answer'] in q['options'], (module, index, 'answer missing from options')
            if q.get('practice_format') == 'true_false_revision':
                assert q['options'] == ['Vrai', 'Faux'], (module, index, 'truth choices must not reveal the answer')
                assert q['options_en'] == ['True', 'False'], (module, index, 'truth values must align across languages')
            assert q['question'].strip() and q['explanation'].strip(), (module, index)
            assert q.get('question_en') and q.get('explanation_en'), (module, index, 'missing English text')
            assert len(q.get('options_en', [])) == len(q['options']), (module, index, 'English option alignment')
            assert len(set(q['options_en'])) == len(q['options_en']), (module, index, 'duplicate English options')
            question_id = stable_id(f'q-{module}-{q["concept_id"]}-{index}')
            assert question_id not in ids, f'Duplicate question: {module}/{index}'
            ids.add(question_id)
            statements.append(upsert('questions', dict(id=question_id, lesson_id=lesson_id, license=program, category=category, theme=category, subtopic=q.get('concept_title', '')[:100], difficulty=q['level'], level=q['level'], question_fr=q['question'], question_en=q.get('question_en', ''), options=q['options'], options_en=q['options_en'], answer_key=q['correct_answer'], explanation_fr=q['explanation'], explanation_en=q.get('explanation_en', ''), tags=q.get('tags', []))))
            question_count += 1
        modules.append((module, lesson_id, category, len(questions)))
    statements.append('COMMIT;\n')
    assert modules, f'No authored lessons for {program}'
    (ROOT / 'scripts/seed' / (program.lower() + '_content.sql')).write_text('\n'.join(statements), encoding='utf-8')
    print(f'Validated {len(modules)} lessons, {question_count} questions; SQL generated without DELETE or TRUNCATE.')
    for module, lesson_id, category, count in modules:
        if module == 'PPL_AC_007':
            print(f'Flight controls: http://localhost:3000/lessons/PPL/{category}/{lesson_id}')
