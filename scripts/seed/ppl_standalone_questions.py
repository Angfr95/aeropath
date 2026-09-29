"""Make contextual revision questions self-contained when drawn into a paper."""
import json
from pathlib import Path
ROOT=Path(__file__).resolve().parents[2]
CHANGES={
 ('PPL_AIRLAW_002',7):('Quel minimum de double commande prévoit le parcours standard PPL(A) FCL.210.A sans crédit particulier ?','What minimum dual instruction does the standard FCL.210.A PPL(A) route require without special credit?'),
 ('PPL_AIRLAW_002',8):('Quel minimum de solo supervisé prévoit le parcours standard PPL(A) FCL.210.A sans crédit particulier ?','What minimum supervised solo time does the standard FCL.210.A PPL(A) route require without special credit?'),
 ('PPL_AIRLAW_002',9):('Dans le parcours standard PPL(A) FCL.210.A, quel minimum de navigation solo est requis ?','In the standard FCL.210.A PPL(A) route, what minimum solo cross-country time is required?'),
 ('PPL_AIRLAW_002',10):('Quelle navigation répond à l’exigence de distance du parcours standard PPL(A) FCL.210.A ?','Which cross-country flight meets the distance requirement of the standard FCL.210.A PPL(A) route?'),
 ('PPL_AIRLAW_011',3):('En classe D, sous le seuil supérieur applicable de SERA.5001, quelle distance verticale aux nuages s’applique en VFR normal de jour ?','In Class D below the applicable SERA.5001 upper threshold, what vertical cloud distance applies in normal daytime VFR?'),
 ('PPL_AIRLAW_011',9):('En classe G dans les basses couches définies par SERA.5001, en VFR de jour sans réduction nationale, quelle condition s’applique aux nuages et à la surface ?','In the SERA.5001 lower band in Class G, daytime VFR with no national reduction, what cloud and surface condition applies?'),
 ('PPL_NAV_009',3):('Route souhaitée au nord, vent de l’ouest : vers où corriger le cap pour maintenir la route ?','Desired northbound track, westerly wind: which way should heading be corrected to maintain track?'),
 ('PPL_NAV_009',8):('Route vraie 000°, TAS 100 kt, vent du 270° vrai à 20 kt, dérive corrigée : vitesse sol approchée ?','True track 000°, TAS 100 kt, wind from 270° true at 20 kt, drift corrected: approximate groundspeed?'),
 ('PPL_NAV_009',9):('Route au nord, TAS 100 kt, vent du 270° à 20 kt : pourquoi ne pas soustraire simplement 20 kt au TAS ?','Northbound track, TAS 100 kt, wind from 270° at 20 kt: why not simply subtract 20 kt from TAS?'),
 ('PPL_NAV_009',12):('Le contrôle visuel confirme un écart à droite de la route prévue. Vers où corriger pour le réduire ?','Visual cross-check confirms displacement to the right of planned track. Which way should you correct to reduce it?'),
 ('PPL_MB_002',7):('Enveloppe fictive définie uniquement entre 700 et 850 kg : quelle masse maximale est couverte ?','A fictional envelope is defined only between 700 and 850 kg: what maximum mass is covered?'),
 ('PPL_MB_002',10):('Une enveloppe fictive autorise une interpolation uniquement entre 700 et 850 kg. Peut-on l’extrapoler à 900 kg ?','A fictional envelope permits interpolation only between 700 and 850 kg. May it be extrapolated to 900 kg?'),
 ('PPL_MB_002',15):('On déplace un bagage de 20 kg du bras 2 m au bras 1 m, sans rien ajouter ni retirer. La masse totale change-t-elle ?','A 20 kg bag is moved from arm 2 m to arm 1 m without adding or removing anything. Does total mass change?'),
 ('PPL_INS_002',3):('Source statique totalement bloquée sur un variomètre classique : après le transitoire, son indication tend vers…','With a completely blocked static source on a conventional vertical speed indicator, after the transient its indication tends towards…'),
}
def main():
    for path in (ROOT/'content/fr/ppl').rglob('*.json'):
        d=json.loads(path.read_text(encoding='utf-8'));module=d['metadata']['module_id'];changed=False
        for (m,index),(fr,en) in CHANGES.items():
            if m==module:d['questions'][index-1].update(question=fr,question_en=en);changed=True
        if changed:path.write_text(json.dumps(d,ensure_ascii=False,indent=2)+'\n',encoding='utf-8')
if __name__=='__main__':main()
