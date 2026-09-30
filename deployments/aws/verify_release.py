"""Check public course content and private endpoint boundaries without an account."""
import collections
import json
import re
import sys
import urllib.error
import urllib.request
from pathlib import Path

root, base = Path(sys.argv[1]), sys.argv[2].rstrip('/')
public = '--public' in sys.argv[3:]

def stable_id(seed):
    value = 0
    for char in seed:
        value = (value * 31 + ord(char)) & 0x7fffffff
    return f'{value:08x}-0000-4000-8000-{value:012x}'

def request(path, payload=None, expected=200):
    req = urllib.request.Request(base + path,
        data=json.dumps(payload).encode() if payload is not None else None,
        headers={'Content-Type': 'application/json'})
    try:
        with urllib.request.urlopen(req, timeout=25) as response:
            status, body = response.status, response.read()
    except urllib.error.HTTPError as error:
        status, body = error.code, error.read()
    assert status == expected, f'{path}: HTTP {status}, expected {expected}'
    return body

assert json.loads(request('/health'))['status'] == 'ok'
for path in ['/api/me', '/api/students', '/api/students/count', '/api/admin/stats',
             '/api/history', '/api/stats', '/api/recommendations']:
    request(path, expected=401)
for path in ['/api/lessons', '/api/questions']:
    request(path, payload={}, expected=401)

counts = collections.Counter()
images = set()
graded = 0
for source in sorted((root/'content/fr').glob('*/*/*.json')):
    data = json.loads(source.read_text(encoding='utf-8'))
    meta = data['metadata']
    program, module = meta['program'], meta['module_id']
    lesson_id = stable_id(f'lesson-{program}-{module}')
    lesson = json.loads(request('/api/lessons/' + lesson_id))
    for concept in data['concepts']:
        for source_key, api_key in [('content', 'content_fr'), ('content_en', 'content_en')]:
            assert concept[source_key].replace('\r\n', '\n') in lesson[api_key].replace('\r\n', '\n'), module
    quiz = json.loads(request('/api/lessons/' + lesson_id + '/quiz'))['questions']
    assert len(quiz) == 15 and len({q['id'] for q in quiz}) == 15, module
    expected_questions = {stable_id(f'q-{module}-{q["concept_id"]}-{i}'): q
                          for i, q in enumerate(data['questions'])}
    for question in quiz:
        expected = expected_questions[question['id']]
        assert question['options'] == expected['options'], module
        assert question['options_en'] == expected['options_en'], module
        assert not ({'answer_key', 'correct_answer', 'password_hash', 'email'} & question.keys()), module
    # Grade one lesson per licence plus the latest PPL altimetry chapter.
    if counts[program] == 0 or module == 'PPL_MET_002':
        qid, q = next(iter(expected_questions.items()))
        correct = q['options'].index(q['correct_answer'])
        for index, want in [(correct, True), ((correct + 1) % len(q['options']), False)]:
            result = json.loads(request('/api/questions/answer', {'question_id': qid, 'answer': chr(65 + index)}))
            assert result['correct'] is want, module
            assert result['explanation_en'].replace('\r\n', '\n') == q['explanation_en'].replace('\r\n', '\n'), module
            graded += 1
    for language in ['content_fr', 'content_en']:
        images.update(re.findall(r'\]\((/illustrations/[^)]+\.svg)\)', lesson[language]))
    counts[program] += 1

papers = json.loads(request('/api/ppl/exam-papers'))['papers']
assert len(papers) == 9
for paper in papers:
    result = json.loads(request('/api/ppl/exam-papers/' + paper['id']))
    assert len(result['questions']) == paper['count'], paper['id']

if public:
    page = request('/').decode()
    assets = re.findall(r'(?:src|href)="(/assets/[^\"]+)"', page)
    assert assets
    for asset in assets:
        request(asset)
    for path in sorted(images):
        assert b'<svg' in request(path), path
    for path in ['/.env', '/.git/config', '/backup.sql', '/ssl/key.pem', '/frontend-react/src/App.tsx']:
        request(path, expected=404)
    # Both measurement variants of the latest FR/EN altimetry figures must be served.
    for units in ['us', 'imperial']:
        for language in ['', 'en/']:
            request(f'/illustrations/units/{units}/{language}ppl-met-002.svg')

print(json.dumps({'lessons_verified': dict(counts), 'quizzes_verified': sum(counts.values()),
                  'guest_answers': graded, 'ppl_papers': len(papers),
                  'private_endpoints_protected': True, 'public_assets_checked': public}))
