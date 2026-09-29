"""Verify cross-licence content integrity and independently calculated answers."""
import json
import math
import re
import sys
import unittest
from pathlib import Path
sys.dont_write_bytecode=True
from shared.content import ROOT, stable_id

LESSONS={}
for program in ['LAPL','CPL','ATPL','IR']:
    for p in (ROOT/'content/fr'/program.lower()).rglob('*.json'):
        d=json.loads(p.read_text(encoding='utf-8'));LESSONS[d['metadata']['module_id']]=d


class LicenceContent(unittest.TestCase):
    def test_licences_options_sources_and_images(self):
        ids=set()
        for module,d in LESSONS.items():
            m=d['metadata'];program=m['program']
            self.assertTrue(module.startswith(program+'_'))
            self.assertEqual(len(d['questions']),15)
            self.assertEqual(m['quiz_format'],'chapter_practice')
            lesson_id=stable_id('lesson-'+program+'-'+module)
            self.assertNotIn(lesson_id,ids);ids.add(lesson_id)
            self.assertTrue(m['sources'])
            for i,q in enumerate(d['questions']):
                revision = program == 'LAPL' and q['practice_format'] == 'true_false_revision'
                count = 2 if revision else 4
                self.assertEqual(len(set(q['options'])),count)
                self.assertEqual(len(set(q['options_en'])),count)
                if revision:self.assertNotIn('exam-mcq',q['tags'])
                self.assertIn(q['correct_answer'],q['options'])
                self.assertTrue(q['syllabus_reference'] and q['source_references'])
                self.assertIn(program,q['tags']);self.assertNotIn('PPL',q['tags'])
                for key in ['explanation','explanation_en']:
                    paths=re.findall(r'!\[[^\]]*\]\(([^)]+)\)',q[key])
                    self.assertEqual(len(paths),1)
                    for path in paths:
                        for prefix in ['', 'units/us/', 'units/imperial/']:
                            target=ROOT/'web/frontend-react/public/illustrations'/prefix/path.removeprefix('/illustrations/')
                            self.assertTrue(target.exists(),target)
                for key in ['question','question_en']:
                    for path in re.findall(r'!\[[^\]]*\]\(([^)]+)\)',q[key]):
                        self.assertTrue((ROOT/'web/frontend-react/public'/path.lstrip('/')).exists())
                qid=stable_id(f'q-{module}-{q["concept_id"]}-{i}')
                self.assertNotIn(qid,ids);ids.add(qid)
        self.assertEqual(len(LESSONS),85)

    def test_lapl_shared_syllabus_and_exam_separation(self):
        lapl=[d for d in LESSONS.values() if d['metadata']['program']=='LAPL']
        self.assertEqual(len(lapl),63)
        self.assertEqual(len({d['metadata']['derived_from'] for d in lapl}),63)
        self.assertEqual(len({d['metadata']['exam_subject'] for d in lapl}),9)
        questions=[q for d in lapl for q in d['questions']]
        self.assertEqual(sum('exam-mcq' in q['tags'] for q in questions),455)
        self.assertEqual(sum(q['practice_format']=='true_false_revision' for q in questions),490)
        # Licence comparisons must still name the licence whose rule is taught.
        source=next((ROOT/'content/fr/ppl/air_law').glob('PPL_AIRLAW_002*.json'))
        original=json.loads(source.read_text(encoding='utf-8'))
        adapted=LESSONS['LAPL_AIRLAW_002']
        for a,b in zip(original['questions'],adapted['questions']):
            self.assertEqual(a['correct_answer'],b['correct_answer'])
            self.assertEqual(a['options'],b['options'])

    def test_expansion_calculations(self):
        def answer(module,n):return LESSONS[module]['questions'][n-1]['correct_answer']
        interpolated=((500+600)/2+(650+780)/2)/2
        self.assertEqual(answer('CPL_032_001',3),f'{interpolated:.1f}'.replace('.',',')+' m')
        self.assertEqual(answer('CPL_032_001',6),f'{math.ceil(interpolated*1.15)} m')
        self.assertEqual(answer('CPL_081_001',5),f'{50*math.sqrt(2):.1f}'.replace('.',',')+' kt')
        self.assertEqual(answer('CPL_033_001',10),f'{180/100*40:.0f} L')
        self.assertEqual(answer('ATPL_022_001',1),f'{200/math.sqrt(.64):.0f} kt')
        self.assertEqual(answer('ATPL_061_001',3),f'{60*4*math.cos(math.radians(60)):.0f} NM')
        self.assertEqual(answer('IR_062_002',8),f'{120/60*318:.0f} ft/min')
        self.assertEqual(answer('IR_022_001',1),f'+{10*27} ft')
        self.assertEqual(answer('ATPL_021_002',1),f'{75-30} kPa')
        self.assertEqual(answer('ATPL_021_002',8),f'{75-25} kPa')
        self.assertEqual(answer('ATPL_021_002',11),'−'+str(65-60)+' kPa')
        trip_hours=120/120+90/180
        total=5+trip_hours*40+6+18+20
        self.assertEqual(answer('IR_033_001',6),f'{total:.0f} L')
        self.assertEqual(answer('IR_033_001',10),f'{120/100*40:.0f} L')
        self.assertEqual(answer('IR_033_001',11),f'{total+(120/100-1)*40:.0f} L')
        self.assertEqual(answer('IR_033_001',12),f'{12/60*30:.0f} L')
        self.assertEqual(answer('IR_010_001',8),f'{2400+1000:,}'.replace(',',' ')+' ft AMSL')
        self.assertEqual(answer('IR_010_001',9),f'{6200+2000:,}'.replace(',',' ')+' ft AMSL')
        self.assertEqual(answer('CPL_022_001',2),f'{850-830} hPa')
        landing=28000+6500+9000-150-5500
        self.assertEqual(answer('ATPL_031_001',4),f'{landing:,}'.replace(',',' ')+' kg')
        self.assertEqual(answer('ATPL_031_001',8),f'{min(36000-34500,45000-43350,39000-landing):,}'.replace(',',' ')+' kg')
        self.assertEqual(answer('ATPL_031_001',12),'Masse admise, CG un point de pourcentage trop avant')
        self.assertEqual(answer('ATPL_031_001',14),f'{100*(3.65-3)/2.6:.0f} %')

    def test_calculated_answer_keys(self):
        def answer(module,n):return LESSONS[module]['questions'][n-1]['correct_answer']
        mass=sum([1800,300,200]);moment=sum(m*a for m,a in [(1800,2),(300,3),(200,2.5)])
        self.assertEqual(answer('CPL_031_001',3),f'{moment/mass:.3f}'.replace('.',',')+' m')
        self.assertEqual(answer('CPL_031_001',8),f'{(moment+100*(1-3))/mass:.3f}'.replace('.',',')+' m')
        self.assertEqual(answer('CPL_031_001',10),f'{(moment-100*2.5)/(mass-100):.3f}'.replace('.',',')+' m')
        etp=600*180/(180+120);psr=4/(1/120+1/180)
        self.assertEqual(answer('ATPL_033_001',1),f'{etp:.0f} NM')
        self.assertEqual(answer('ATPL_033_001',8),f'{psr:.0f} NM')
        self.assertAlmostEqual(etp/180,(600-etp)/120)
        self.assertAlmostEqual(psr/120+psr/180,4)
        self.assertEqual(answer('IR_062_001',1),f'{math.hypot(3,4):.0f} NM')
        self.assertEqual(answer('IR_062_001',9),f'{math.sqrt(5**2-3**2):.0f} NM')

    def test_ppl_ids_and_scope_remain_separate(self):
        shared=LESSONS['LAPL_NAV_001']
        self.assertEqual(shared['metadata']['derived_from'],'PPL_NAV_009')
        self.assertNotEqual(stable_id('lesson-LAPL-LAPL_NAV_001'),stable_id('lesson-PPL-PPL_NAV_009'))
        ids=set()
        for p in (ROOT/'content/fr').rglob('*.json'):
            data=json.loads(p.read_text(encoding='utf-8'));meta=data['metadata']
            if '/ppl/' in p.as_posix():self.assertEqual(meta['program'],'PPL')
            module=meta['module_id']
            current=[stable_id('lesson-'+meta['program']+'-'+module)]
            current += [stable_id(f'q-{module}-{q["concept_id"]}-{i}') for i,q in enumerate(data['questions'])]
            for identifier in current:
                self.assertNotIn(identifier,ids)
                ids.add(identifier)


if __name__=='__main__':unittest.main()
