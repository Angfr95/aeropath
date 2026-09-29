"""Independent checks of sign, datum, scenario constraints and bilingual keys."""
import json,unittest
from pathlib import Path
import xml.etree.ElementTree as ET
ROOT=Path(__file__).resolve().parents[2]
D=json.loads((ROOT/'content/fr/ppl/meteo/PPL_MET_002_pressure.json').read_text(encoding='utf-8'))


class Altimetry(unittest.TestCase):
    def key(self,n,fr,en):
        q=D['questions'][n-1]
        self.assertEqual(q['correct_answer'],fr)
        self.assertEqual(q['options_en'][q['options'].index(fr)],en)

    def test_setting_sign_and_inverse(self):
        # Reference-offset model: pressure stays constant as the reference changes.
        for n,start,old,new,expected in [(2,2400,1002,1010,2616),(11,3200,1020,1008,2876),(13,2600,1000,1012,2924)]:
            offset=27*(new-old);self.assertEqual(start+offset,expected)
            self.assertEqual(expected+27*(old-new),start)
            self.key(n,f'{expected:,} ft'.replace(',',' '),f'{expected:,} ft')

    def test_qfe_does_not_measure_local_terrain_clearance(self):
        altitude=900+1200;terrain=1400
        self.key(4,f'{altitude-terrain} ft',f'{altitude-terrain} ft')
        self.key(7,f'{650+1450:,} ft'.replace(',',' '),f'{650+1450:,} ft')

    def test_clearance_requires_both_correction_and_obstacle(self):
        corrected=3100-27*(1022-1010);clearance=corrected-2300
        self.assertEqual(corrected,2776);self.assertEqual(clearance,476)
        self.assertLess(clearance,500)
        self.key(12,f'{clearance} ft disponibles, déficit de {500-clearance} ft',f'{clearance} ft available, {500-clearance} ft shortfall')
        for lang in ['', 'en']:
            text=' '.join(ET.parse(ROOT/'web/frontend-react/public/illustrations'/lang/'quiz-ppl-met-002-clearance.svg').getroot().itertext())
            for answer in ['2776','2 776','2,776','476','24 ft']:self.assertNotIn(answer,text)
            for value in ['1022','1010','500 ft','27 ft/hPa']:self.assertIn(value,text)


if __name__=='__main__':unittest.main()
