"""Independent arithmetic and constraint checks for the CPL application cases."""
import json
import math
import unittest
from fractions import Fraction as F
from pathlib import Path
import xml.etree.ElementTree as ET

ROOT=Path(__file__).resolve().parents[2]


class PerformanceFuelPedagogy(unittest.TestCase):
    def lesson(self,module):
        return json.loads(next((ROOT/'content/fr/cpl').rglob(module+'*.json')).read_text(encoding='utf-8'))

    def answer(self,module,number,fr,en):
        q=self.lesson(module)['questions'][number-1]
        self.assertEqual(q['correct_answer'],fr)
        self.assertEqual(q['options_en'][q['options'].index(fr)],en)

    def test_interpolation_and_rounding(self):
        # Weighted corner contributions, independent of the row-first worked method.
        temp=F(10-5,25-5);alt=F(1500,2000)
        raw=600*(1-temp)*(1-alt)+720*temp*(1-alt)+800*(1-temp)*alt+960*temp*alt
        self.assertEqual(raw,F(1575,2))
        margin=900-math.ceil(raw*F(112,100))
        deficit=math.ceil(raw*F(112,100)*F(108,100))-950
        self.answer('CPL_032_001',9,f'{margin} m',f'{margin} m')
        self.answer('CPL_032_001',10,f'Déficit de {deficit} m',f'{deficit} m shortfall')
        self.assertEqual(900-math.ceil(F(600+720+800+960,4)*F(112,100)),37)
        self.assertEqual(900-math.ceil(raw),112)
        self.assertEqual(math.ceil(810*F(112,100)),908)
        value=math.ceil(810*F(108,100))
        self.answer('CPL_032_001',8,f'{value} m',f'{value} m')

    def test_fuel_phase_and_reference_balances(self):
        times=[F(90,120),F(40,80)]
        self.assertEqual(sum(times)*60,75)
        burn=sum(t*f for t,f in zip(times,[32,28]))
        self.answer('CPL_033_001',2,f'{burn} L',f'{burn} L')
        required=sum([4,8,38,6,9,24]);usable=104-4
        self.answer('CPL_033_001',4,f'{required} L',f'{required} L')
        self.answer('CPL_033_001',5,f'{usable-required} L',f'{usable-required} L')
        self.answer('CPL_033_001',6,f'{usable-required} L',f'{usable-required} L')
        margin=64-F(45,60)*40-5-24
        self.answer('CPL_033_001',9,f'{margin} L',f'{margin} L')
        self.assertEqual(F(21.6).limit_denominator()/F('.72'),30)

    def test_suitable_arrival_requires_both_constraints(self):
        candidates=[('A',40,True),('B',30,False),('C',50,True)]
        arrivals={name:58-F(minutes,60)*42-4 for name,minutes,_ in candidates}
        self.assertEqual(arrivals,{'A':26,'B':33,'C':19})
        viable=[name for name,_,suitable in candidates if suitable and arrivals[name]>=24]
        self.assertEqual(viable,['A'])
        self.answer('CPL_033_001',13,'A, arrivée 26 L','A, arrival 26 L')

    def test_prompt_diagrams_supply_inputs_without_solution(self):
        for lang in ['', 'en']:
            assets=ROOT/'web/frontend-react/public/illustrations'/lang
            perf=' '.join(ET.parse(assets/'cpl-032-001.svg').getroot().itertext())
            for value in ['600 m','720 m','800 m','960 m','5 °C','25 °C']:self.assertIn(value,perf)
            for value in ['787','882','953','18 m']:self.assertNotIn(value,perf)
            fuel=' '.join(ET.parse(assets/'cpl-033-001.svg').getroot().itertext())
            for value in ['58 L','24 L','40 min','30 min','50 min','42 L/h','4 L']:self.assertIn(value,fuel)
            for value in ['26 L','33 L','19 L']:self.assertNotIn(value,fuel)


if __name__=='__main__':unittest.main()
