"""Select the quantitative diagrams for the September PPL application lessons."""
import json
from pathlib import Path

HERE=Path(__file__).resolve().parent
def main():
    p=HERE/'ppl_diagrams.json';specs=json.loads(p.read_text(encoding='utf-8'))
    for module,kind in [('PPL_MB_002','cg_envelope'),('PPL_NAV_009','wind_vectors'),('PPL_POF_003','stability_curves')]:specs[module]['kind']=kind
    p.write_text(json.dumps(specs,ensure_ascii=False,indent=2)+'\n',encoding='utf-8')
    pairs='''1,00|1.00
1,05|1.05
1,20|1.20
700|700
850|850
Bras (m)|Arm (m)
Masse (kg)|Mass (kg)
830 kg ; 1,096 m|830 kg; 1.096 m
800 kg ; 1,081 m|800 kg; 1.081 m
Enveloppe fictive|Fictional envelope
Interpolation linéaire de la limite avant.|Linear interpolation of the forward limit.
Départ|Departure
830 kg et 910 kg·m.|830 kg and 910 kg·m.
Après consommation|After fuel burn
800 kg et 865 kg·m.|800 kg and 865 kg·m.
Air : 100 kt|Air: 100 kt
Cap : 348,5° vrai|Heading: 348.5° true
Vent du 270° : 20 kt|Wind from 270°: 20 kt
Sol : ≈ 98 kt|Ground: ≈ 98 kt
Route : 000° vrai|Track: 000° true
Vecteurs à la même échelle ; vent vers l’est.|Vectors at the same scale; wind towards east.
Toutes les directions sont vraies.|All directions are true.
Composante ouest de 20 kt dans la vitesse air.|20 kt westward component in air velocity.
La somme des vecteurs donne la vitesse sol.|Adding the vectors gives ground velocity.
Oscillations amorties|Damped oscillations
Oscillations divergentes|Divergent oscillations
Temps →|Time →
Statique|Static
La tendance initiale après perturbation.|Initial tendency after a disturbance.
Dynamique|Dynamic
L’évolution de l’amplitude au fil du temps.|How amplitude develops over time.
Résultat|Result
Interprétation|Interpretation
Un retour initial ne garantit pas l’amortissement.|An initial return does not guarantee damping.'''
    p=HERE/'ppl_svg_en.json';trans=json.loads(p.read_text(encoding='utf-8'));trans.update(dict(row.split('|') for row in pairs.splitlines()))
    p.write_text(json.dumps(trans,ensure_ascii=False,indent=2)+'\n',encoding='utf-8')
if __name__=='__main__':main()
