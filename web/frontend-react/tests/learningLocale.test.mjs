import test from 'node:test';
import assert from 'node:assert/strict';
import { readFileSync, readdirSync, existsSync } from 'node:fs';
import vm from 'node:vm';
import ts from 'typescript';

const context = { exports: {} };
const source = readFileSync(new URL('../src/utils/learningLocale.ts', import.meta.url), 'utf8');
vm.runInNewContext(ts.transpileModule(source, { compilerOptions: { module: ts.ModuleKind.CommonJS } }).outputText, context);
const { localizedText, localizedOptions } = context.exports;

test('language switching preserves choice positions and canonical answers', () => {
  const question = { options: ['gauche', 'droite'], options_en: ['left', 'right'] };
  const before = JSON.stringify(question);
  assert.equal(localizedOptions(question, 'en')[1], 'right');
  assert.equal(localizedOptions(question, 'fr')[1], 'droite');
  assert.equal(localizedText('en', 'français', 'English'), 'English');
  assert.equal(JSON.stringify(question), before);
  assert.equal(localizedOptions({ ...question, options_en: ['left'] }, 'en')[1], 'droite');
  assert.equal(localizedText('en', 'fallback', ''), 'fallback');
});

test('all PPL lessons and quizzes have English content and English SVGs', () => {
  const root = new URL('../../../content/fr/ppl/', import.meta.url);
  let lessons = 0, questions = 0;
  function checkImages(text) {
    for (const match of text.matchAll(/!\[[^\]]*\]\(([^)]+)\)/g)) {
      assert.match(match[1], /^\/illustrations\/en\//);
      assert.ok(existsSync(new URL('../public' + match[1], import.meta.url)), match[1]);
    }
  }
  for (const file of readdirSync(root, { recursive: true }).filter((name) => name.endsWith('.json'))) {
    const data = JSON.parse(readFileSync(new URL(file.replaceAll('\\', '/'), root), 'utf8'));
    lessons++;
    assert.equal(data.questions.length, 15);
    assert.ok(data.metadata.chapter_title_en);
    for (const concept of data.concepts) {
      assert.ok(concept.title_en && concept.content_en);
      checkImages(concept.content_en);
    }
    assert.match(data.concepts[0].content_en, /\/illustrations\/en\//);
    for (const question of data.questions) {
      questions++;
      assert.ok(question.question_en && question.explanation_en);
      assert.equal(question.options_en.length, question.options.length);
      assert.ok(question.options_en.every((choice) => choice.trim()));
      assert.equal(new Set(question.options_en).size, question.options.length);
      assert.ok(question.options.includes(question.correct_answer));
      assert.match(question.explanation_en, /\/illustrations\/en\//);
      checkImages(question.question_en); checkImages(question.explanation_en);
    }
  }
  assert.equal(lessons, 63); assert.equal(questions, 945);
});
