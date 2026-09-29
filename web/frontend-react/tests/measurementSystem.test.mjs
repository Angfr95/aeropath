import test from 'node:test';
import assert from 'node:assert/strict';
import { readFileSync, existsSync, readdirSync } from 'node:fs';
import { createRequire } from 'node:module';
import vm from 'node:vm';
import ts from 'typescript';
const require = createRequire(import.meta.url);
function load(file, globals = {}) {
  const context = { exports: {}, Intl, ...globals };
  const code = ts.transpileModule(readFileSync(new URL('../src/' + file, import.meta.url), 'utf8'),
    { compilerOptions: { module: ts.ModuleKind.CommonJS, target: ts.ScriptTarget.ES2020 } }).outputText;
  vm.runInNewContext(code, context);
  return context.exports;
}
const units = load('utils/measurementSystem.ts');
const { measurementText } = units;

test('US and imperial gallons differ; source equations and aviation units remain explicit', () => {
  assert.equal(measurementText('24 L/h', 'us', 'en'), '24 L/h (≈ 6.34 US gal/h)');
  assert.equal(measurementText('24 L/h', 'imperial', 'en'), '24 L/h (≈ 5.28 imp gal/h)');
  assert.equal(measurementText('20 kg', 'us', 'en'), '20 kg (≈ 44.09 lb)');
  assert.equal(measurementText('1,5 m', 'imperial', 'fr'), '1,5 m (≈ 4,92 ft)');
  assert.equal(measurementText('1.5 m', 'us', 'en'), '1.5 m (≈ 4.92 ft)');
  assert.equal(measurementText('30 kg·m', 'us', 'en'), '30 kg·m (≈ 216.99 lb·ft)');
  assert.equal(measurementText('1,000 kg', 'us', 'en'), '1,000 kg (≈ 2,204.62 lb)');
  assert.equal(measurementText('1.000 kg', 'us', 'es'), '1.000 kg (≈ 2204,62 lb)');
  assert.equal(measurementText('1,000.5 kg', 'us', 'en'), '1,000.5 kg (≈ 2,205.72 lb)');
  assert.equal(measurementText('1,143 m', 'us', 'en', 'fr'), '1,143 m (≈ 3.75 ft)');
  const aviation = '90 kt · 45 NM · 2,000 ft · 1013 hPa · 15 °C';
  assert.equal(measurementText(aviation, 'us', 'en'), aviation);
  assert.equal(measurementText('24 × 0.5 = 12 L', 'us', 'en'), '24 × 0.5 = 12 L (≈ 3.17 US gal)');
  assert.equal(measurementText('24 L/h', 'metric', 'fr'), '24 L/h');
});

test('link targets and code are preserved; only local illustration variants are selected', () => {
  const input = '[24 L](https://example.com/24L) `24 L`';
  assert.equal(measurementText(input, 'us', 'en'), '[24 L (≈ 6.34 US gal)](https://example.com/24L) `24 L`');
  const image = '![Fuel](/illustrations/en/ppl-fp-001.svg)';
  assert.equal(measurementText(image, 'imperial', 'en'), '![Fuel](/illustrations/units/imperial/en/ppl-fp-001.svg)');
});

test('language and units persist independently, including reload and invalid stored units', () => {
  const saved = new Map();
  const localStorage = { getItem: (key) => saved.get(key) ?? null, setItem: (key, value) => saved.set(key, value) };
  const build = () => load('store/uiStore.ts', { localStorage, setTimeout, require: (name) => {
    if (name === 'zustand') return require('zustand');
    if (name === '@/utils/measurementSystem') return units;
    if (name === '@/constants/languages') return { DEFAULT_LANGUAGE: 'fr' };
    throw new Error(name);
  } }).useUIStore;
  const store = build();
  store.getState().setLang('es');
  store.getState().setMeasurementSystem('imperial');
  assert.equal(store.getState().currentLang, 'es');
  store.getState().setLang('en');
  assert.equal(store.getState().measurementSystem, 'imperial');
  assert.equal(build().getState().currentLang, 'en');
  assert.equal(build().getState().measurementSystem, 'imperial');
  saved.set('aeropath_units', 'bad');
  assert.equal(build().getState().measurementSystem, 'metric');
});

test('all existing French and English illustrations have both unit variants', () => {
  const root = new URL('../public/illustrations/', import.meta.url);
  for (const lang of ['', 'en/']) {
    for (const name of readdirSync(new URL(lang, root)).filter((name) => name.endsWith('.svg'))) {
      for (const system of ['us', 'imperial']) {
        assert.ok(existsSync(new URL(`units/${system}/${lang}${name}`, root)));
      }
    }
  }
  const us = readFileSync(new URL('units/us/en/ppl-fp-001.svg', root), 'utf8');
  const uk = readFileSync(new URL('units/imperial/en/ppl-fp-001.svg', root), 'utf8');
  assert.match(us, /US gal/); assert.match(uk, /imp gal/);
});
