import test from 'node:test';
import assert from 'node:assert/strict';
import { readFileSync } from 'node:fs';
import { createRequire } from 'node:module';
import vm from 'node:vm';
import ts from 'typescript';

const require = createRequire(import.meta.url);
const source = readFileSync(new URL('../src/store/quizStore.ts', import.meta.url), 'utf8');
const compiled = ts.transpileModule(source, { compilerOptions: { module: ts.ModuleKind.CommonJS } });
const question = { id: 'q1', options: ['Faux', 'Juste'], question_fr: 'Exemple', license: 'PPL', category: 'navigation' };
const result = { data: { correct: true, correct_answer: 'Juste', explanation_fr: 'Explication.', explanation_en: '' } };
function deferred() {
  let resolve, reject;
  const promise = new Promise((yes, no) => { resolve = yes; reject = no; });
  return { promise, resolve, reject };
}
function setup(overrides = {}) {
  const calls = { answer: 0, random: 0, lesson: 0 };
  const questionsApi = {
    random: async () => { calls.random++; return { data: { questions: [question] } }; },
    examByCategory: async () => ({ data: [question] }),
    answer: async (...args) => { calls.answer++; return overrides.answer ? overrides.answer(...args) : result; },
  };
  const lessonsApi = { quiz: async (...args) => {
    calls.lesson++;
    return overrides.quiz ? overrides.quiz(...args) : { data: { questions: [question] } };
  } };
  const context = { exports: {}, require: (name) => {
    if (name === 'zustand') return require('zustand');
    if (name === '@/api/questions') return { questionsApi };
    if (name === '@/api/lessons') return { lessonsApi };
    throw new Error(`Unexpected module: ${name}`);
  } };
  vm.runInNewContext(compiled.outputText, context);
  return { store: context.exports.useQuizStore, calls };
}

test('lesson quiz scores once, translates the answer text, and replays the same questions', async () => {
  const { store, calls } = setup();
  await store.getState().startByLesson('lesson1');
  store.getState().next();
  assert.equal(store.getState().phase, 'active');
  await store.getState().submitAnswer('q1', 'B');
  assert.equal(store.getState().feedback.correct_answer, 'B');
  assert.equal(store.getState().feedback.explanation_fr, 'Explication.');
  await store.getState().submitAnswer('q1', 'B');
  assert.equal(calls.answer, 1);
  assert.equal(store.getState().score, 1);
  store.getState().next();
  assert.equal(store.getState().phase, 'finished');
  store.getState().restart();
  assert.equal(store.getState().phase, 'active');
  assert.equal(store.getState().lessonId, 'lesson1');
  assert.equal(store.getState().questions[0].id, 'q1');
  assert.equal(store.getState().score, 0);
  assert.equal(store.getState().feedback, null);
  assert.equal(calls.lesson, 1);
  assert.equal(calls.random, 0);
});

test('a fifteen-question lesson finishes after the last correction and can be replayed', async () => {
  const questions = Array.from({ length: 15 }, (_, i) => ({ ...question, id: `q${i}` }));
  const { store, calls } = setup({ quiz: async () => ({ data: { questions } }) });
  await store.getState().startByLesson('lesson15');
  for (const q of questions) {
    assert.equal(store.getState().phase, 'active');
    await store.getState().submitAnswer(q.id, 'B');
    store.getState().next();
  }
  assert.equal(store.getState().phase, 'finished');
  assert.equal(store.getState().score, 15);
  assert.equal(calls.answer, 15);
  store.getState().restart();
  assert.equal(store.getState().phase, 'active');
  assert.equal(store.getState().questions.length, 15);
  assert.equal(store.getState().score, 0);
  assert.equal(calls.lesson, 1);
});

test('double click submits once while the correction is pending', async () => {
  const pending = deferred();
  const { store, calls } = setup({ answer: () => pending.promise });
  await store.getState().startByLesson('lesson1');
  const first = store.getState().submitAnswer('q1', 'B');
  await store.getState().submitAnswer('q1', 'A');
  assert.equal(calls.answer, 1);
  assert.equal(store.getState().submitting, true);
  pending.resolve(result);
  await first;
  assert.equal(store.getState().score, 1);
  assert.equal(store.getState().submitting, false);
});

test('an answer failure keeps the question available for retry', async () => {
  let fail = true;
  const { store } = setup({ answer: async () => {
    if (fail) throw new Error('offline');
    return result;
  } });
  await store.getState().startByLesson('lesson1');
  await store.getState().submitAnswer('q1', 'B');
  assert.equal(store.getState().phase, 'active');
  assert.equal(store.getState().score, 0);
  assert.equal(store.getState().submitting, false);
  assert.ok(store.getState().error);
  fail = false;
  await store.getState().submitAnswer('q1', 'B');
  assert.equal(store.getState().score, 1);
  assert.equal(store.getState().error, null);
});

test('empty and failed loads show an error rather than a blank active quiz', async () => {
  for (const quiz of [async () => ({ data: { questions: [] } }), async () => { throw new Error('offline'); }]) {
    const { store } = setup({ quiz });
    await store.getState().startByLesson('lesson1');
    assert.equal(store.getState().phase, 'idle');
    assert.equal(store.getState().loading, false);
    assert.ok(store.getState().error);
  }
});

test('a late correction cannot modify a different quiz', async () => {
  const pending = deferred();
  const { store } = setup({ answer: () => pending.promise });
  await store.getState().startByLesson('lesson1');
  const answer = store.getState().submitAnswer('q1', 'B');
  store.getState().reset();
  await store.getState().startByCategory('PPL', 'navigation');
  pending.resolve(result);
  await answer;
  assert.equal(store.getState().phase, 'active');
  assert.equal(store.getState().lessonId, null);
  assert.equal(store.getState().score, 0);
  assert.equal(store.getState().feedback, null);
});

test('a late lesson load cannot replace a newer quiz', async () => {
  const pending = deferred();
  const { store } = setup({ quiz: () => pending.promise });
  const load = store.getState().startByLesson('lesson1');
  await store.getState().startRandom();
  pending.resolve({ data: { questions: [{ ...question, id: 'old' }] } });
  await load;
  assert.equal(store.getState().questions[0].id, 'q1');
  assert.equal(store.getState().lessonId, null);
});

test('incorrect answers retain zero score and invalid submissions are ignored', async () => {
  const { store, calls } = setup({ answer: async () => ({ data: { ...result.data, correct: false } }) });
  await store.getState().startByLesson('lesson1');
  await store.getState().submitAnswer('other', 'A');
  await store.getState().submitAnswer('q1', 'Z');
  assert.equal(calls.answer, 0);
  await store.getState().submitAnswer('q1', 'A');
  assert.equal(store.getState().score, 0);
  assert.equal(store.getState().feedback.selected_answer, 'A');
  assert.equal(store.getState().feedback.correct_answer, 'B');
  store.getState().next();
  assert.equal(store.getState().phase, 'finished');
});
