"""LAPL navigation diagrams, sharing the PPL theoretical syllabus geometry."""
import sys
import json
sys.dont_write_bytecode = True
from svg_common import ROOT


def main():
    assets = ROOT / 'web/frontend-react/public/illustrations'
    sources={d['metadata']['module_id']:d for p in (ROOT/'content/fr/ppl').rglob('*.json') for d in [json.loads(p.read_text(encoding='utf-8'))]}
    for path in (ROOT/'content/fr/lapl').rglob('*.json'):
        meta=json.loads(path.read_text(encoding='utf-8'))['metadata']
        original=meta['derived_from'];source_name=sources[original]['metadata']['illustration']['src'].rsplit('/',1)[1]
        target=meta['illustration']['src'].rsplit('/',1)[1]
        for directory in [assets, assets / 'en']:
            source=(directory/source_name).read_text(encoding='utf-8')
            (directory/target).write_text(source.replace(original,meta['module_id']).replace('AEROPATH / PPL /','AEROPATH / LAPL /'),encoding='utf-8')


if __name__ == '__main__':
    main()
