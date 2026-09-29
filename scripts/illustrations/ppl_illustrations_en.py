"""Generate English SVGs from the same geometry and an explicit text catalogue."""
import json
import xml.etree.ElementTree as ET
from pathlib import Path
import ppl_illustrations as base
import ppl_quiz_illustrations as quiz

TRANSLATIONS = json.loads(Path(__file__).with_name('ppl_svg_en.json').read_text(encoding='utf-8'))
ENGLISH = set(TRANSLATIONS.values())

def translate(value):
    value = str(value)
    if value in TRANSLATIONS: return TRANSLATIONS[value]
    if value in ENGLISH: return value
    raise ValueError(f'Missing English SVG translation: {value}')

class EnglishSVG(base.SVG):
    def __init__(self, module, title, subtitle, description):
        self.wrapping = False
        super().__init__(module, translate(title), translate(subtitle), translate(description))

    def text(self, x, y, text, *args, **kwargs):
        super().text(x, y, str(text) if self.wrapping else translate(text), *args, **kwargs)

    def wrap(self, x, y, text, *args, **kwargs):
        self.wrapping = True
        try: return super().wrap(x, y, translate(text), *args, **kwargs)
        finally: self.wrapping = False

def main():
    output = base.ASSETS/'en'
    output.mkdir(exist_ok=True)
    specs = json.loads(Path(__file__).with_name('ppl_diagrams.json').read_text(encoding='utf-8'))
    for module, spec in specs.items():
        drawing = EnglishSVG(module, spec['title'], spec['subtitle'], spec['description'])
        base.RENDERERS[spec['kind']](drawing, spec)
        (output/(module.lower().replace('_','-')+'.svg')).write_text(drawing.finish()+'\n',encoding='utf-8')
    quiz.SVG = EnglishSVG
    for (module, index), (kind, a, b) in quiz.EXERCISES.items():
        svg, _ = quiz.draw(module, kind, a, b)
        name = f'quiz-{module.lower().replace("_", "-")}-{index+1}.svg'
        (output/name).write_text(svg+'\n',encoding='utf-8')
    tree = ET.parse(base.ASSETS/'commandes-de-vol.svg')
    ET.register_namespace('', 'http://www.w3.org/2000/svg')
    for element in tree.getroot().iter():
        if element.tag.rsplit('}',1)[-1] in {'text','title','desc'} and element.text:
            element.text = translate(element.text)
    tree.write(output/'commandes-de-vol.svg',encoding='utf-8',xml_declaration=False)
    print(f'{len(specs)+len(quiz.EXERCISES)+1} English SVGs generated.')

if __name__ == '__main__': main()
