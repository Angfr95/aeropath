"""Regression checks for answer leakage, independent calculations and assets."""
import json
import math
import unittest
from pathlib import Path
import xml.etree.ElementTree as ET

ROOT = Path(__file__).resolve().parents[2]
LESSONS = {d['metadata']['module_id']: d for p in (ROOT/'content/fr').rglob('*.json') for d in [json.loads(p.read_text(encoding='utf-8'))]}


class CourseQuality(unittest.TestCase):
    def test_truth_choices_do_not_encode_correctness(self):
        for program in ['PPL', 'LAPL']:
            qs = [q for m,d in LESSONS.items() if m.startswith(program+'_') for q in d['questions'] if len(q['options']) == 2]
            self.assertTrue(qs)
            for q in qs:
                self.assertEqual(q['options'], ['Vrai', 'Faux'])
                self.assertEqual(q['options_en'], ['True', 'False'])
            fraction = sum(q['correct_answer'] == q['options'][0] for q in qs) / len(qs)
            self.assertLess(abs(fraction - .5), .1, 'A constant-position strategy must not achieve near-perfect revision scores')

    def test_static_pressure_counterexample(self):
        # Independent pressure-difference model: descent does not imply overreading.
        # At 6,000 ft, current static is below static trapped at 5,000 ft.
        trapped_static, current_static, impact_pressure = 843, 812, 50
        total = current_static + impact_pressure
        faulty = total - trapped_static
        self.assertLess(faulty, impact_pressure)
        q = LESSONS['IR_022_001']['questions'][11]
        self.assertEqual(q['correct_answer'], 'Sous-indication')
        for marker in ['5 000', '7 000', '6 000']: self.assertIn(marker, q['question'])
        self.assertIn('descend', q['question'])
        self.assertIn('Underreading', q['options_en'])
        self.assertEqual(q['options_en'][q['options'].index(q['correct_answer'])], 'Underreading')

    def test_performance_worked_examples(self):
        def bilinear(a, t):
            return (1-a)*((1-t)*500+t*600) + a*((1-t)*650+t*780)
        self.assertEqual(bilinear(0, .5), 550)
        self.assertEqual(bilinear(.5, .5), 632.5)
        self.assertEqual(bilinear(.25, .25), 564.375)
        d = LESSONS['CPL_032_001']
        self.assertEqual(d['questions'][0]['correct_answer'], '550 m')
        self.assertEqual(d['questions'][2]['correct_answer'], '632,5 m')
        self.assertEqual(d['questions'][5]['correct_answer'], str(math.ceil(bilinear(.5,.5)*1.15))+' m')
        self.assertIn('564,375', d['concepts'][1]['content'])
        self.assertIn('564.375', d['concepts'][1]['content_en'])

    def test_shared_table_does_not_impose_a_conflicting_target(self):
        paths = list((ROOT/'web/frontend-react/public/illustrations').rglob('cpl-032-001.svg'))
        self.assertEqual(len(paths), 6)
        for p in paths:
            text = ' '.join(ET.parse(p).getroot().itertext())
            self.assertNotIn('Point à calculer', text)
            self.assertNotIn('Point to calculate', text)
            self.assertTrue('énoncé' in text or 'question' in text, str(p))


if __name__ == '__main__': unittest.main()
