// Reuse the frontend converter so text and diagrams have identical equivalents.
import { readFileSync, writeFileSync, readdirSync, mkdirSync } from 'node:fs';
import { createRequire } from 'node:module';
import { fileURLToPath } from 'node:url';
import path from 'node:path';
import vm from 'node:vm';
const root = fileURLToPath(new URL('../../', import.meta.url));
const require = createRequire(path.join(root, 'web/frontend-react/package.json'));
const ts = require('typescript');
const context = { exports: {}, Intl };
vm.runInNewContext(ts.transpileModule(readFileSync(path.join(root, 'web/frontend-react/src/utils/measurementSystem.ts'), 'utf8'),
  { compilerOptions: { module: ts.ModuleKind.CommonJS, target: ts.ScriptTarget.ES2020 } }).outputText, context);
const { measurementText } = context.exports;
const assets = path.join(root, 'web/frontend-react/public/illustrations');
const escape = (text) => text.replaceAll('&', '&amp;').replaceAll('<', '&lt;').replaceAll('>', '&gt;');
let count = 0;
for (const lang of ['fr', 'en', 'es']) {
  const directory = lang === 'fr' ? assets : path.join(assets, lang);
  let files;
  try { files = readdirSync(directory).filter((name) => name.endsWith('.svg')); }
  catch (error) { if (error.code === 'ENOENT') continue; throw error; }
  for (const system of ['us', 'imperial']) {
    const output = path.join(assets, 'units', system, lang === 'fr' ? '' : lang);
    mkdirSync(output, { recursive: true });
    for (const name of files) {
      let svg = readFileSync(path.join(directory, name), 'utf8');
      const equivalents = new Set();
      for (const match of svg.matchAll(/<(?:text|tspan)\b[^>]*>([^<]+)</g)) {
        const text = match[1].replaceAll('&amp;', '&').replaceAll('&lt;', '<').replaceAll('&gt;', '>');
        const converted = measurementText(text, system, lang);
        for (const item of converted.matchAll(/(-?\d[\d., \u00a0\u202f]*\s*(?:kg[·⋅]m|kg|L\/h|L|cm|m)) \(≈ ([^)]+)\)/g)) {
          equivalents.add(`${item[1].trim()} ≈ ${item[2]}`);
        }
      }
      if (equivalents.size) {
        const rows = Math.ceil(equivalents.size / 3);
        const height = 792 + rows * 30;
        svg = svg.replace('height="720"', `height="${height}"`).replace('viewBox="0 0 1200 720"', `viewBox="0 0 1200 ${height}"`);
        const heading = lang === 'fr' ? 'Équivalences arrondies · valeurs d’origine conservées pour les calculs'
          : lang === 'es' ? 'Equivalencias redondeadas · valores originales para los cálculos'
          : 'Rounded equivalents · original values retained for calculations';
        let footer = `<rect x="0" y="720" width="1200" height="${height - 720}" fill="#10243d"/>`;
        footer += `<text x="42" y="754" fill="#58dfed" font-family="Segoe UI, Arial, sans-serif" font-size="18">${escape(heading)}</text>`;
        [...equivalents].forEach((value, index) => {
          footer += `<text x="${42 + index % 3 * 390}" y="${786 + Math.floor(index / 3) * 30}" fill="#edf5ff" font-family="Segoe UI, Arial, sans-serif" font-size="17">${escape(value)}</text>`;
        });
        svg = svg.replace('</svg>', footer + '</svg>');
      }
      writeFileSync(path.join(output, name), svg);
      count++;
    }
  }
}
console.log(`${count} SVG unit variants generated.`);
