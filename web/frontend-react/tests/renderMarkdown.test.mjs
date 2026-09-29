import test from 'node:test';
import assert from 'node:assert/strict';
import { readFileSync } from 'node:fs';
import vm from 'node:vm';
import ts from 'typescript';

const source = readFileSync(new URL('../src/utils/renderMarkdown.ts', import.meta.url), 'utf8');
const compiled = ts.transpileModule(source, { compilerOptions: { module: ts.ModuleKind.CommonJS } });
const context = { exports: {} };
vm.runInNewContext(compiled.outputText, context);
const { renderMarkdown } = context.exports;

test('untrusted HTML cannot create executable elements', () => {
  for (const input of [
    '<img src=x onerror=alert(1)>',
    '<svg onload=alert(1)>',
    '</p><script>alert(1)</script>',
    '## <iframe srcdoc="<script>alert(1)</script>">',
    '**<img src=x onerror=alert(1)>**',
  ]) {
    const html = renderMarkdown(input);
    assert.doesNotMatch(html, /<\/?(?:img|svg|script|iframe)\b/i);
    assert.match(html, /&lt;/);
  }
});

test('entities stay text and Markdown formatting remains available', () => {
  assert.match(renderMarkdown('&lt;img&gt;'), /&amp;lt;img&amp;gt;/);
  assert.match(renderMarkdown('## Heading\n\n**bold**'), /<h2[^>]*>Heading<\/h2>/);
  assert.match(renderMarkdown('**bold**'), /<strong[^>]*>bold<\/strong>/);
});

test('lesson illustrations allow only local SVG assets and escape alt attributes', () => {
  const html = renderMarkdown('![Axes " du vol](/illustrations/commandes-de-vol.svg)');
  assert.match(html, /<img src="\/illustrations\/commandes-de-vol.svg"/);
  assert.match(html, /alt="Axes &quot; du vol"/);
  for (const input of [
    '![x](javascript:alert(1))',
    '![x](https://example.com/tracker.svg)',
    '![x](/illustrations/../untrusted.svg)',
    '![x](/illustrations/x.svg" onerror="alert(1))',
  ]) assert.doesNotMatch(renderMarkdown(input), /<img\b/);
});

test('source links allow HTTPS without enabling executable URLs or attributes', () => {
  assert.match(renderMarkdown('[Reference](https://example.com/manual?q=1&v=2)'), /href="https:\/\/example.com\/manual\?q=1&amp;v=2"/);
  for (const input of [
    '[unsafe](javascript:alert(1))',
    '[unsafe](data:text/html,test)',
    '[unsafe](https://example.com/" onclick="alert(1))',
  ]) assert.doesNotMatch(renderMarkdown(input), /<a\b/);
});
