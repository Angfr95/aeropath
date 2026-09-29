"""Check the independent transfer case and bilingual correction illustrations."""
import json
import unittest
from pathlib import Path
import xml.etree.ElementTree as ET

ROOT = Path(__file__).resolve().parents[2]
LESSON = json.loads((ROOT/'content/fr/ir/meteo/IR_050_001_application.json').read_text(encoding='utf-8'))


class IcingPedagogy(unittest.TestCase):
    def test_constraints_before_shortest_time(self):
        routes = [('A',12,120,6500,0),('B',21,140,8000,7),('C',16,120,8000,10)]
        feasible = [(name,60*distance/gs) for name,distance,gs,altitude,fuel in routes if altitude >= 7000 and fuel <= 8]
        self.assertEqual(feasible,[('B',9)])
        self.assertEqual(60*routes[0][1]/routes[0][2],6)
        q = LESSON['questions'][13]
        self.assertEqual(q['correct_answer'],'B, 9 minutes')
        self.assertEqual(q['options_en'][q['options'].index(q['correct_answer'])],
                         'B, 9 minutes')
        for text,markers in [(q['question'],['12 NM','21 NM','6 500','8 000','7 000','140 kt','10 kg']),
                             (q['question_en'],['12 NM','21 NM','6,500','8,000','7,000','140 kt','10 kg'])]:
            for marker in markers:self.assertIn(marker,text)

    def test_transfer_values_differ_from_worked_example(self):
        c = LESSON['concepts'][3]
        self.assertEqual(60*6/120,3)
        self.assertEqual(60*18/120,9)
        for field in ['content','content_en']:
            self.assertIn('18 NM',c[field])
            self.assertIn('9 minutes',c[field])
            self.assertNotIn('24 NM',c[field])

    def test_diagram_matches_case_without_showing_correction_in_prompt(self):
        q = LESSON['questions'][13]
        for suffix,lang in [('', ''),('_en','en/')]:
            self.assertNotIn('ir-icing-exit.svg',q['question'+suffix])
            self.assertIn('/illustrations/'+lang+'ir-050-001.svg',q['explanation'+suffix])
            path = ROOT/'web/frontend-react/public/illustrations'/lang/'ir-050-001.svg'
            text = ' '.join(ET.parse(path).getroot().itertext())
            for marker in ['21','140','10','8 kg']:self.assertIn(marker,text)
            self.assertNotIn('9 min',text)  # Inputs shown, calculated answer withheld.
            for slug in ['ir-icing-exit.svg','ir-icing-effects.svg']:
                for units in ['','units/us/','units/imperial/']:
                    self.assertTrue((ROOT/'web/frontend-react/public/illustrations'/units/lang/slug).exists())


if __name__ == '__main__':
    unittest.main()
