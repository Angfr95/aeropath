"""Independently verify the FR/EN mass-and-balance transfer exercises."""
import json
import unittest
from pathlib import Path

ROOT = Path(__file__).resolve().parents[2]
MODULES = {'PPL_MB_002', 'LAPL_MB_002', 'CPL_031_001', 'ATPL_031_001'}
LESSONS = {d['metadata']['module_id']:d for p in (ROOT/'content/fr').rglob('*.json') for d in [json.loads(p.read_text(encoding='utf-8'))] if d['metadata']['module_id'] in MODULES}


def cg(items):
    return sum(m*a for m,a in items) / sum(m for m,a in items)


class MassAndBalanceQuality(unittest.TestCase):
    def assert_bilingual_answer(self, module, number, fr, en):
        q = LESSONS[module]['questions'][number-1]
        self.assertEqual(q['correct_answer'], fr)
        self.assertEqual(q['options_en'][q['options'].index(fr)], en)

    def test_cpl_inverse_shift_question(self):
        moved = 2400 * (2.20-2.10) / 3
        self.assert_bilingual_answer('CPL_031_001', 7, f'{moved:.0f} kg', f'{moved:.0f} kg')

    def test_cpl_combined_changes_and_distractors(self):
        mass = 2300 - 100 + 50
        moment = 5000 - 100*2.5 + 50*1
        final = moment/mass
        self.assert_bilingual_answer('CPL_031_001', 11,
                                     f'{final:.3f}'.replace('.', ',')+' m', f'{final:.3f} m')
        q = LESSONS['CPL_031_001']['questions'][10]
        # Each distractor must actually represent the mistake named in feedback.
        mistakes = [(5000-250)/2200, moment/2300, (5000+250+50)/mass]
        for value in mistakes:
            en = f'{value:.3f}'
            fr = en.replace('.', ',')
            self.assertIn(en+' m', q['options_en'])
            self.assertIn(fr+' m', q['options'])
            self.assertIn(en, q['explanation_en'])
            self.assertIn(fr, q['explanation'])

    def test_atpl_mass_stage_questions(self):
        ramp = 28000+6500+9000
        for number, mass in [(2, ramp), (3, ramp-150)]:
            self.assert_bilingual_answer('ATPL_031_001', number,
                                         f'{mass:,}'.replace(',', ' ')+' kg', f'{mass:,} kg')

    def test_atpl_payload_decisions_include_all_limits(self):
        baseline = [34500, 43350, 37850]
        limits = [36000, 45000, 39000]
        excesses = [mass+1400-limit for mass, limit in zip(baseline, limits)]
        self.assertEqual([i for i, excess in enumerate(excesses) if excess > 0], [2])
        self.assert_bilingual_answer('ATPL_031_001', 5,
                                     f'Atterrissage de {excesses[2]} kg', f'Landing by {excesses[2]} kg')
        zfm = 28000+6500
        takeoff = zfm+9000-150
        landing = takeoff-5500
        allowed = min(36000-zfm, 45000-takeoff, 39000-landing, 43800-takeoff)
        self.assert_bilingual_answer('ATPL_031_001', 6, f'{allowed} kg', f'{allowed} kg')
        extra_fuel_masses = [34500, 34500+10600-150, 34500+10600-150-5500]
        offload = max(0, *(mass-limit for mass, limit in zip(extra_fuel_masses, limits)))
        self.assertTrue(all(mass-offload <= limit for mass, limit in zip(extra_fuel_masses, limits)))
        self.assertTrue(any(mass-(offload-1) > limit for mass, limit in zip(extra_fuel_masses, limits)))
        self.assert_bilingual_answer('ATPL_031_001', 7, f'{offload} kg', f'{offload} kg')

    def test_atpl_combined_moments_and_mac_questions(self):
        final_mass = 10000-1000
        final_moment = 10000*5 + 200*(3-7) - 1000*6
        shifted = final_moment/final_mass
        balanced = (10000*5 - 500*6 - 500*4)/(10000-500-500)
        for number, value in [(10, shifted), (11, balanced)]:
            self.assert_bilingual_answer('ATPL_031_001', number,
                                         f'{value:.3f}'.replace('.', ',')+' m', f'{value:.3f} m')
        mac = 100*(41400/9000-4)/2
        self.assertAlmostEqual(mac,30)
        self.assert_bilingual_answer('ATPL_031_001', 15,
            'CG 30 %MAC, dans les limites 17,5–35 % de cette masse',
            "CG 30% MAC, within this mass's 17.5–35% limits")
        q = LESSONS['ATPL_031_001']['questions'][9]
        for wrong in [(50000-6000)/9000, final_moment/10000, (50000+800-6000)/9000]:
            self.assertIn(f'{wrong:.3f} m', q['options_en'])

    def test_method_diagram_is_in_feedback_only_in_both_languages(self):
        expected = {'CPL_031_001': {5,6,7,8,9,10,11,15}, 'ATPL_031_001': {10,11}}
        for module, numbers in expected.items():
            for number, q in enumerate(LESSONS[module]['questions'], start=1):
                for suffix, language in [('', ''), ('_en', 'en/')]:
                    self.assertNotIn('mass-balance-method.svg', q['question'+suffix])
                    self.assertEqual('mass-balance-method.svg' in q['explanation'+suffix], number in numbers)
                    if number in numbers:
                        self.assertIn('/illustrations/'+language+'mass-balance-method.svg', q['explanation'+suffix])
        for language in ['', 'en/']:
            for units in ['', 'units/us/', 'units/imperial/']:
                self.assertTrue((ROOT/'web/frontend-react/public/illustrations'/units/language/'mass-balance-method.svg').exists())

    def test_ppl_added_load_and_envelope(self):
        initial = [(600,1),(150,1.2),(60,1.5),(20,2)]
        final = initial + [(10,2)]
        mass = sum(m for m,a in final)
        forward = 1 + (mass-700)/(850-700)*.05
        self.assertGreater(cg(final), cg(initial))
        self.assertLess(forward, cg(final))
        self.assertLess(cg(final), 1.2)
        for module in ['PPL_MB_002','LAPL_MB_002']:
            c = LESSONS[module]['concepts'][2]
            self.assertIn(f'{cg(final):.4f}'.replace('.',','), c['content'])
            self.assertIn(f'{cg(final):.4f}', c['content_en'])
            self.assertIn('700', LESSONS[module]['questions'][7]['question'])
            self.assertIn('850', LESSONS[module]['questions'][7]['question_en'])

    def test_cpl_shift_inverse_and_datum_invariance(self):
        total = 2300
        moved = total * .05 / 2
        self.assertEqual(moved,57.5)
        c = LESSONS['CPL_031_001']['concepts'][1]
        self.assertIn('57,5',c['content']);self.assertIn('57.5',c['content_en'])
        items = [(1800,2),(300,3),(200,2.5)]
        self.assertAlmostEqual(cg([(m,a+1) for m,a in items])-cg(items),1)
        mass=sum(m for m,a in items);moment=sum(m*a for m,a in items)
        removed=100;arm=2.5
        direct=(moment-removed*arm)/(mass-removed)-cg(items)
        formula=removed*(cg(items)-arm)/(mass-removed)
        self.assertAlmostEqual(direct,formula)

    def test_atpl_mass_stages_and_limiting_payload(self):
        dow,load,block,taxi,trip=30000,7500,10000,200,6000
        zfm=dow+load;takeoff=zfm+block-taxi;landing=takeoff-trip
        self.assertEqual(landing,zfm+(block-taxi-trip))
        allowed=min(39000-zfm,49000-takeoff,42000-landing)
        self.assertEqual(allowed,700)
        c=LESSONS['ATPL_031_001']['concepts'][1]
        self.assertIn('700 kg',c['content']);self.assertIn('700 kg',c['content_en'])
        # A separate extra-fuel case changes landing mass without changing ZFM.
        self.assertEqual(38000+10400-200-6000,42200)
        self.assertIn('42 200',c['content']);self.assertIn('42,200',c['content_en'])

    def test_fuel_burn_direction_depends_on_arm(self):
        original=5
        forward_tank=(10000*5-1000*4)/9000
        aft_tank=(10000*5-1000*6)/9000
        self.assertGreater(forward_tank,original)
        self.assertLess(aft_tank,original)
        q=LESSONS['ATPL_031_001']['questions'][12]
        states=[(10000,46000),(9000,44000),(8000,37000)]
        outside=[]
        for i,(mass,moment) in enumerate(states):
            pct=100*(moment/mass-4)/2
            lower=15+(mass-8000)/4000*10
            if not lower <= pct <= 35:outside.append(i)
        self.assertEqual(outside,[1])
        self.assertEqual(q['correct_answer'],'Seulement après le premier réservoir')
        self.assertEqual(q['options_en'][q['options'].index(q['correct_answer'])],'Only after the first tank')

    def test_mac_is_relative_to_leading_edge(self):
        percentage=100*(3.7-3.2)/2
        self.assertAlmostEqual(percentage,25)
        self.assertAlmostEqual(3.2+percentage/100*2,3.7)
        d=LESSONS['ATPL_031_001']
        self.assertIn('3,70',d['concepts'][3]['content'])
        self.assertIn('3.70',d['concepts'][3]['content_en'])
        self.assertTrue(d['questions'][14]['correct_answer'].startswith('CG 30 %MAC'))


if __name__=='__main__':unittest.main()
