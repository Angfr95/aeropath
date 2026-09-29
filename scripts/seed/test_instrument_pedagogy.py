"""Independent pressure models for CPL failure scenarios and bilingual answers."""
import json
import unittest
from pathlib import Path

ROOT = Path(__file__).resolve().parents[2]
D = json.loads((ROOT/'content/fr/cpl/instrumentation/CPL_022_001_application.json').read_text(encoding='utf-8'))


class InstrumentPedagogy(unittest.TestCase):
    def answer(self,n,fr,en):
        q=D['questions'][n-1]
        self.assertEqual(q['correct_answer'],fr)
        self.assertEqual(q['options_en'][q['options'].index(fr)],en)

    def test_trapped_total_pressure_and_drain_case(self):
        pt=800
        values=[pt-ps for ps in [780,760,790]]
        self.assertEqual(values,[20,40,10])
        self.answer(5,f'De {values[0]} à {values[1]} hPa',f'From {values[0]} to {values[1]} hPa')
        self.answer(6,f'{values[2]} hPa ; indication en baisse',f'{values[2]} hPa; indication decreases')
        self.assertIn('drainage et statique libres',D['questions'][3]['question'])
        self.assertIn('Vers zéro',D['questions'][3]['correct_answer'])

    def test_double_blockage_does_not_track_external_changes(self):
        trapped_pt,trapped_ps=800,780
        for external_ps,impact_pressure in [(780,20),(760,40),(790,15)]:
            received=trapped_pt-trapped_ps
            self.assertEqual(received,20)
            self.assertGreater(external_ps+impact_pressure,external_ps)
        self.answer(7,'Elle reste à 20 hPa','It remains at 20 hPa')

    def test_static_error_is_relative_to_blockage_not_direction(self):
        pt,external,trapped=850,810,840
        measured,correct=pt-trapped,pt-external
        self.assertEqual((measured,correct),(10,40))
        self.answer(10,'10 hPa reçus contre 40 hPa corrects : sous-indication',
                       '10 hPa received versus 40 hPa correct: underreading')
        # New transfer case: climbing but still below the altitude of blockage.
        self.assertGreater(950-850,950-880)
        q=D['questions'][10]
        for marker in ['4 000','2 000','3 000']:self.assertIn(marker,q['question'])
        self.assertTrue(q['correct_answer'].startswith('Surindication'))
        self.assertTrue(q['options_en'][q['options'].index(q['correct_answer'])].startswith('Overreading'))

    def test_alternate_source_changes_pressure_difference(self):
        correct=900-870
        measured=900-865
        self.assertGreater(measured,correct)
        self.answer(12,f'{measured} hPa',f'{measured} hPa')
        # Worked case uses different values: 920 - 885 = 35, versus 920 - 890 = 30.
        self.assertEqual(920-885,35)
        self.assertIn('920 hPa',D['concepts'][0]['content'])

    def test_feedback_diagrams_do_not_leak_into_questions(self):
        mapping={**{n:'cpl-pitot-reasoning.svg' for n in [4,5,6,7,10,11,12]},
                 13:'cpl-sensor-independence.svg'}
        for n,slug in mapping.items():
            q=D['questions'][n-1]
            for suffix,lang in [('', ''),('_en','en/')]:
                self.assertNotIn(slug,q['question'+suffix])
                self.assertIn('/illustrations/'+lang+slug,q['explanation'+suffix])
                for units in ['', 'units/us/', 'units/imperial/']:
                    self.assertTrue((ROOT/'web/frontend-react/public/illustrations'/units/lang/slug).exists())


if __name__ == '__main__':unittest.main()
