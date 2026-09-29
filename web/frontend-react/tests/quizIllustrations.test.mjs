import test from 'node:test';
import assert from 'node:assert/strict';
import { readdirSync, readFileSync, existsSync } from 'node:fs';
import { join } from 'node:path';
import { fileURLToPath } from 'node:url';
import vm from 'node:vm';
import ts from 'typescript';

const root = fileURLToPath(new URL('../../../', import.meta.url));
const source = readFileSync(new URL('../src/utils/renderMarkdown.ts', import.meta.url), 'utf8');
const context = { exports: {} };
vm.runInNewContext(ts.transpileModule(source, { compilerOptions: { module: ts.ModuleKind.CommonJS } }).outputText, context);
function* files(dir) {
  for (const item of readdirSync(dir, { withFileTypes: true })) {
    if (item.isDirectory()) yield* files(join(dir, item.name));
    else if (item.name.endsWith('.json')) yield join(dir, item.name);
  }
}
test('every licence correction renders its SVG and exercise illustrations resolve locally', () => {
  let exercises = 0;
  for (const file of files(join(root, 'content/fr'))) {
    const lesson = JSON.parse(readFileSync(file, 'utf8'));
    for (const q of lesson.questions) {
      for (const suffix of ['', '_en']) {
        const correction = context.exports.renderMarkdown(q['explanation' + suffix]);
        const illustration = q.explanation_illustration ?? lesson.metadata.illustration.src;
        const expected = suffix ? illustration.replace('/illustrations/', '/illustrations/en/') : illustration;
        assert.equal((correction.match(/<img /g) ?? []).length, 1, q.question);
        assert.ok(correction.includes(`src="${expected}"`));
        for (const text of [q['question' + suffix], q['explanation' + suffix]]) {
          const html = context.exports.renderMarkdown(text);
          for (const [, src] of html.matchAll(/<img src="([^"]+)"/g)) {
            assert.ok(existsSync(join(root, 'web/frontend-react/public', src)), src);
          }
        }
      }
      if (q.question.includes('/illustrations/quiz-')) {
        exercises++;
        assert.match(context.exports.renderMarkdown(q.question), /<img src="\/illustrations\/quiz-/);
      }
    }
  }
  assert.ok(exercises > 0);
});
