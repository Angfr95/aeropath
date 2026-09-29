"""Bilingual correction diagram for combined loading changes, without case data."""
from licence_diagrams import render
from svg_common import CYAN, AMBER, WHITE, MUTED


def draw(s, t, en):
    s.text(65,215,t('1. Fixer l’état initial', '1. Define the initial state'),23,CYAN,weight=700)
    s.text(65,249,t('Masse M ; moment J ; bras depuis le même datum.',
                    'Mass M; moment J; arms from the same datum.'),21,WHITE)
    for x,fr,eng,formula,note_fr,note_en in [
        (65,'Déplacement interne','Internal movement','ΔM = 0',
         'ΔJ = m × (bras final − bras initial)', 'ΔJ = m × (final arm − initial arm)'),
        (620,'Ajout / retrait','Addition / removal',
         t('ΔM = +m (ajout) ou −m (retrait)', 'ΔM = +m (add) or −m (remove)'),
         'ΔJ = ΔM × bras', 'ΔJ = ΔM × arm')]:
        s.rect(x,277,515,145,'#1c344b','#304961',12)
        s.text(x+18,307,t(fr,eng),23,CYAN,weight=700)
        s.text(x+18,345,formula,22,AMBER)
        s.text(x+18,385,t(note_fr,note_en),21,WHITE)
    s.text(65,459,t('2. Mettre à jour les deux totaux', '2. Update both totals'),23,CYAN,weight=700)
    s.text(65,493,t('M finale = M + ΣΔM ; J final = J + ΣΔJ',
                    'Final M = M + ΣΔM; final J = J + ΣΔJ'),23,WHITE)
    s.text(65,537,t('3. CG = J final / M finale', '3. CG = final J / final M'),23,AMBER,weight=700)
    s.text(65,574,t('Contrôler le sens, les unités et l’enveloppe à la masse finale.',
                    'Check direction, units and the envelope at final mass.'),21,MUTED)


if __name__ == '__main__':
    render('CPL / ATPL','MASS_BALANCE_METHOD',
           ('Chargement : combiner les modifications', 'Loading: combine the changes'),draw)
