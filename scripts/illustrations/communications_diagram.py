"""Shared readback-loop illustration for the common 090 objectives."""
from svg_common import CYAN, AMBER, MUTED


def draw(s, t, en):
    for x, title in [(95,t('CONTRÔLEUR','CONTROLLER')),(770,t('PILOTE','PILOT'))]:
        s.rect(x,225,330,240,'#203a50',stroke='#46dce8',radius=20)
        s.text(x+165,275,title,25,CYAN,anchor='middle')
    s.arrow(435,315,760,315,CYAN,3)
    s.text(600,300,t('Clairance','Clearance'),23,anchor='middle')
    s.arrow(760,395,435,395,AMBER,3)
    s.text(600,380,t('Collationnement','Readback'),23,anchor='middle')
    s.text(260,340,t('Écouter','Listen'),23,anchor='middle')
    s.text(260,385,t('Corriger les écarts','Correct discrepancies'),21,anchor='middle')
    s.text(935,340,t('Vérifier l’indicatif','Check callsign'),21,anchor='middle')
    s.text(935,385,t('Clarifier tout doute','Clarify any doubt'),21,anchor='middle')
    s.wrap(95,530,t('Niveau, cap, vitesse, QNH, piste, code SSR : vérifier les éléments de sécurité.','Level, heading, speed, QNH, runway, SSR code: verify safety-related items.'),1000,24,line=32)
    s.text(95,620,t('SERA.8015(e) • Schéma de principe, pas une transcription radio complète.','SERA.8015(e) • Principle diagram, not a complete radio transcript.'),18,MUTED)
