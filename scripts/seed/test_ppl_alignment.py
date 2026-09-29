"""Regression checks for corrected teaching facts and worked examples."""
import json
import math
import unittest
from pathlib import Path

ROOT=Path(__file__).resolve().parents[2]
LESSONS={d['metadata']['module_id']:d for p in (ROOT/'content/fr/ppl').rglob('*.json') for d in [json.loads(p.read_text(encoding='utf-8'))]}

class PedagogyRegression(unittest.TestCase):
    def answer(self,module,index):return LESSONS[module]['questions'][index-1]['correct_answer']

    def test_corrected_legal_facts(self):
        self.assertEqual(self.answer('PPL_AIRLAW_001',1),'1947')
        self.assertEqual(self.answer('PPL_AIRLAW_002',4),'31')
        self.assertEqual(self.answer('PPL_AIRLAW_002',6),'45 h')
        self.assertEqual(self.answer('PPL_AIRLAW_002',7),'25 h')
        self.assertEqual(self.answer('PPL_AIRLAW_002',8),'10 h')
        self.assertEqual(self.answer('PPL_AIRLAW_002',9),'5 h')
        self.assertEqual(self.answer('PPL_AIRLAW_011',1),'8 km')
        self.assertEqual(self.answer('PPL_AIRLAW_011',8),'5 km')
        for module in ['PPL_AIRLAW_002','PPL_AIRLAW_003']:
            content=' '.join(c['content']+' '+c['content_en'] for c in LESSONS[module]['concepts'])
            self.assertNotIn('20h solo',content)
            self.assertNotIn('**3 ans**',content)

    def test_worked_wind_triangle(self):
        heading=360-math.degrees(math.asin(20/100))
        speed=math.sqrt(100**2-20**2)
        self.assertAlmostEqual(heading,348.5,delta=.05)
        self.assertEqual(round(speed),98)
        self.assertEqual(self.answer('PPL_NAV_009',7),'348,5°')
        self.assertEqual(self.answer('PPL_NAV_009',8),'98 kt')
        self.assertEqual(self.answer('PPL_NAV_009',15),'8° gauche')

    def test_loading_and_envelope(self):
        masses=[600,150,60,20];arms=[1,1.2,1.5,2]
        moment=sum(m*a for m,a in zip(masses,arms));mass=sum(masses)
        self.assertEqual(mass,830);self.assertEqual(moment,910)
        self.assertAlmostEqual(moment/mass,1.096,delta=.001)
        remaining=(moment-30*1.5)/(mass-30)
        self.assertAlmostEqual(remaining,1.081,delta=.001)
        forward=1+(800-700)/(850-700)*.05
        self.assertTrue(forward<remaining<1.2)
        self.assertEqual(self.answer('PPL_MB_002',4),'Environ 1,096 m')
        self.assertEqual(self.answer('PPL_MB_002',12),'Environ 1,081 m')

    def test_bank_separates_revision_and_exam(self):
        mcq=revision=0
        for d in LESSONS.values():
            self.assertTrue(d['metadata'].get('sources'))
            for q in d['questions']:
                self.assertTrue(q.get('source_references'))
                if 'exam-mcq' in q['tags']:
                    mcq+=1;self.assertEqual(len(q['options']),4)
                    self.assertIn('exam-subject:'+d['metadata']['exam_subject'],q['tags'])
                else:revision+=1;self.assertEqual(q['practice_format'],'true_false_revision')
        self.assertEqual((mcq,revision),(455,490))

    def test_priority_and_airspace_scenarios(self):
        self.assertEqual(self.answer('PPL_AIRLAW_012',5),'Votre avion')
        self.assertEqual(self.answer('PPL_AIRLAW_012',14),'Rester arrêté et demander clarification')
        self.assertEqual(self.answer('PPL_AIRLAW_012',15),'Non, obtenir aussi la clairance')
        self.assertEqual(self.answer('PPL_AIRLAW_013',4),'Des IFR')
        self.assertEqual(self.answer('PPL_AIRLAW_013',6),'Information de trafic et conseil d’évitement sur demande')
        self.assertEqual(self.answer('PPL_AIRLAW_013',14),'Non, modes A et C ou S requis')
        for module in ['PPL_AIRLAW_012','PPL_AIRLAW_013']:
            self.assertEqual(LESSONS[module]['metadata']['reference_checked_on'],'2026-09-22')
            self.assertTrue(all(q.get('regulatory_reference') for q in LESSONS[module]['questions']))

if __name__=='__main__':unittest.main()
