"""Shared bilingual diagram output, without licence-specific teaching content."""
from svg_common import SVG, ROOT


def table(s, headers, rows, note):
    """Three-column exercise table; caller supplies bilingual subject data."""
    from svg_common import CYAN, AMBER, MUTED
    for x, label in zip([85, 485, 845], headers):
        s.text(x, 235, label, 21, CYAN)
    for i, row in enumerate(rows):
        y = 295 + i * 55
        for x, value in zip([85, 485, 845], row):
            s.text(x, y, value, 21, AMBER if x == 845 else MUTED)
        s.line(85, y + 14, 1115, y + 14, '#304961', 1)
    s.wrap(85, 590, note, 1020, 18, line=23)


def render(program, module, titles, draw):
    assets = ROOT / 'web/frontend-react/public/illustrations'
    for en in [False, True]:
        out = assets / 'en' if en else assets
        out.mkdir(exist_ok=True)
        t = lambda fr, eng: eng if en else fr
        title = t(*titles)
        s = SVG(module, title, t('Exercice fictif · Schéma non à l’échelle', 'Fictional exercise · Diagram not to scale'), title, program=program)
        draw(s, t, en)
        text = s.finish()
        if en:
            text = text.replace('COMPRENDRE EN IMAGE', 'VISUAL LEARNING').replace('Schéma pédagogique · Données d’exercice, sauf mention contraire', 'Training diagram · Exercise data unless stated otherwise').replace('Relier le schéma aux explications et aux références de la leçon.', 'Read the diagram with the lesson explanations and references.')
        (out / (module.lower().replace('_', '-') + '.svg')).write_text(text + '\n', encoding='utf-8')
