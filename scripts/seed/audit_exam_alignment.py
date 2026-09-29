"""Compare every licence/subject with published exam formats; never certify coverage.

Run from any directory. Writes a reproducible inventory and fails on structural
errors or conflicting chapter-format metadata, not on openly incomplete courses.
"""
import json
import re
from datetime import date
from collections import defaultdict
from pathlib import Path

ROOT = Path(__file__).resolve().parents[2]
DGAC = 'https://www.ecologie.gouv.fr/sites/default/files/documents/nouveau_format_examens_theoriques_PPL_LAPL_BPL_SPL.pdf'
EASA = 'https://www.easa.europa.eu/en/document-library/easy-access-rules/online-publications/easy-access-rules-aircrew-regulation-eu-no?kw=ppl&page=67'
ECQB = 'https://www.easa.europa.eu/en/domains/aircrew-and-medical/european-central-question-bank-ecqb'
AMENDMENT = 'https://www.easa.europa.eu/sites/default/files/dfu/amc_gm_to_part-ara_-_issue_1_amendment_13.pdf'

# Independently transcribed from DSAC 03/10/2023 and AMC1 ARA.FCL.300(b).
# Professional tuples: ATPL(A), CPL(A), IR(A) full modular; NOT CB-IR or BIR.
PRIVATE = {
    'airlaw': ('Droit aérien',20,30),
    'human_performance': ('Performance humaine',12,18),
    'meteorology': ('Météorologie',20,30),
    'communications': ('Communications',8,12),
    'navigation': ('Navigation',20,50),
    'principles_of_flight': ('Principes du vol',12,18),
    'operational_procedures': ('Procédures opérationnelles',12,18),
    'flight_performance_planning': ('Performances et préparation du vol',12,18),
    'aircraft_general': ('Connaissance générale de l’aéronef',12,18),
}
PROFESSIONAL = {
    '010': ('Droit aérien',(44,60),(33,45),(30,45)),
    '021': ('Cellule, systèmes et moteurs',(80,120),(60,90),None),
    '022': ('Instrumentation',(60,90),(40,60),(20,30)),
    '031': ('Masse et centrage',(25,75),(21,60),None),
    '032': ('Performances avion',(45,120),(28,75),None),
    '033': ('Préparation et suivi du vol',(42,120),(33,90),(27,75)),
    '040': ('Performance humaine',(48,90),(35,60),(35,60)),
    '050': ('Météorologie',(84,120),(63,90),(63,90)),
    '061': ('Navigation générale',(55,135),(45,120),None),
    '062': ('Radionavigation',(66,90),(22,30),(44,60)),
    '070': ('Procédures opérationnelles',(42,75),(30,60),None),
    '081': ('Principes du vol avion',(46,90),(37,75),None),
    '090': ('Communications',(34,60),(34,60),(34,60)),
}
REVIEW = {
    'airlaw': 'Vérifier les signaux, les cas locaux et la traçabilité de chaque règle ; conserver les corrections SERA déjà réalisées.',
    'human_performance': 'Évaluer les scénarios physiologiques, illusions et décisions ; les chapitres ajoutés ne prouvent pas une couverture exhaustive.',
    'meteorology': 'Contrôler le décodage complet et des décisions à partir de cartes et bulletins, avec distracteurs plausibles.',
    'communications': 'Compléter la pratique contextualisée de phraséologie, urgence et panne ; un QCM ne valide pas la compétence orale.',
    'navigation': 'Contrôler les exercices de carte, dérive, radionavigation et limites GNSS ; vérifier les annexes utilisables sans réponse visible.',
    'principles_of_flight': 'Vérifier les courbes, commandes, stabilité et limites sur des situations variées, au-delà des exercices introductifs.',
    'operational_procedures': 'Contrôler turbulence de sillage, risques au sol, cisaillement et procédures anormales avec leurs conditions d’application.',
    'flight_performance_planning': 'Diversifier tableaux AFM fictifs, enveloppes et bilans ; distinguer hypothèses d’exercice et exigences de carburant.',
    'aircraft_general': 'Vérifier systèmes, limitations, hélices et diagnostic instrumental avec des annexes représentatives.',
    '010': 'Développer le parcours réglementaire propre à la licence/qualification : Part-FCL, SERA, ATS et procédures applicables. Les seuls minima IFR et changements de régime ne couvrent pas toute la matière.',
    '021': 'Développer cellule, électricité, hydraulique, carburant, moteurs et protections ; turbine et pressurisation ne suffisent pas.',
    '022': 'Compléter instruments, capteurs, gyroscopes, automatismes et alertes selon les objectifs applicables à chaque parcours.',
    '031': 'Développer limitations, distribution des charges et documents de chargement, pas seulement le barycentre.',
    '032': 'Développer les domaines et classes de performances applicables, montée, croisière et atterrissage avec annexes.',
    '033': 'Compléter plans de vol, routes/cartes, carburant réglementaire et suivi ; des calculs fictifs seuls ne couvrent pas la matière.',
    '040': 'Compléter physiologie, perception et psychologie ; un chapitre ciblé sur le TEM ou la désorientation ne couvre pas toutes les sous-matières.',
    '050': 'Développer atmosphère, fronts, vents, nuages, dangers et exploitation des produits météo ; le givrage seul ne suffit pas.',
    '061': 'Développer magnétisme, projections, navigation à l’estime et navigation en vol, avec lecture de cartes.',
    '062': 'Compléter principes radio, radar, GNSS et navigation de surface/PBN selon la colonne du parcours ; VOR/DME/ILS sont partiels.',
    '070': 'Créer les procédures normales/spéciales, risques et contraintes opérationnelles du parcours.',
    '081': 'Développer aérodynamique, stabilité, commandes, hélices et limites ; adapter la compressibilité aux objectifs de la licence.',
    '090': 'Compléter les sept rubriques du tableau 090 ; le seul collationnement ne représente pas toute l’épreuve.',
}


def run():
    errors, rows = [], []
    report = ['# Audit par matière des examens théoriques', '',
        f'Inventaire exécuté le {date.today().isoformat()} — avions, France / EASA.', '',
        '**Conclusion : format PPL vérifié ; conformité pédagogique complète non établie pour aucun parcours.**', '',
        'Cet audit contrôle tous les fichiers de contenu, les catégories d’examen, les choix, traductions, références et quotas déclarés. '
        'Il inventorie les chapitres et les priorités par matière ; il ne prétend pas valider individuellement toutes les affirmations ni tous les objectifs EASA. '
        'Une quantité suffisante de QCM ne constitue pas une preuve de représentativité.', '',
        'La [contre-vérification réglementaire et pédagogique du 28 septembre 2026](regulatory-level-review-2026-09-28.md) '
        'précise les sources actuelles, les lacunes confirmées et les limites de cette comparaison. '
        'La date d’exécution de ce script ne vaut pas nouvelle validation des textes réglementaires.', '',
        '## Références et limites', '',
        f'- [DGAC : quotas PPL/LAPL]({DGAC}) : 9 matières, 128 questions, 212 minutes, 75 % par matière.',
        f'- [EASA : quotas par matière]({EASA}) : tableau public édition août 2023 ; comparaison ATPL(A), CPL(A), IR(A) modulaire complet.',
        f'- [Amendement 2025/002/R]({AMENDMENT}) : correction de durée du 070 CPL(H), sans modification des colonnes avion contrôlées ici.',
        f'- [ECQB 2026 et objectifs]({ECQB}) : banque officielle non publique ; aucune comparaison question par question avec cette banque. '
        'Les objectifs publiés et les annexes de référence sont la base de la revue pédagogique.',
        '- IR désigne ici la qualification complète, pas le CB-IR ni le BIR. Le choix de langue/unités ne transforme pas le cours en préparation FAA ou UK CAA.', '',
        '## Mode examen et couverture', '',
        'Le code `internal/learning/ppl_exam.go` propose uniquement les neuf épreuves PPL. '
        'Les quotas et durées sont comparés ci-dessous à la note DGAC. Le store PPL applique 75 % par matière, '
        'zéro aux blancs, aucun point négatif, un délai conservé au rechargement et des corrections après remise. '
        'LAPL, CPL, ATPL et IR disposent de quiz de chapitre, mais pas de simulateur complet propre à leur parcours. '
        'Les limites de tentatives, sessions et validité administrative ne sont pas simulées.', '',
        'Les 530 vrai/faux de chaque parcours PPL/LAPL sont des révisions et sont exclus de la sélection PPL. '
        'Le marquage `exam-mcq` indique un type de question éligible ; il ne constitue pas une validation par une autorité. '
        'La sélection PPL répartit les questions entre chapitres, sans matrice exhaustive d’objectifs. '
        'Les références génériques à une matière doivent être complétées par des objectifs précis et une revue des distracteurs. '
        'Les SVG de correction ne remplacent pas des exercices sur annexes dans les énoncés.', '']
    go = (ROOT/'internal/learning/ppl_exam.go').read_text(encoding='utf-8')
    implemented = {m[0]:(int(m[1]),int(m[2])) for m in re.findall(r'\{"([a-z_]+)",[^\n]+?, (\d+), (\d+)\}',go)}
    for program in ['PPL','LAPL','CPL','ATPL','IR']:
        expected = PRIVATE if program in ['PPL','LAPL'] else {
            code:(v[0],*v[['ATPL','CPL','IR'].index(program)+1]) for code,v in PROFESSIONAL.items()
            if v[['ATPL','CPL','IR'].index(program)+1] is not None}
        grouped=defaultdict(list)
        for path in sorted((ROOT/'content/fr'/program.lower()).rglob('*.json')):
            d=json.loads(path.read_text(encoding='utf-8'));m=d['metadata'];subject=m['exam_subject']
            if subject not in expected: errors.append(f'{m["module_id"]}: unexpected subject {subject}')
            grouped[subject].append(d)
        report += [f'## {program}(A)', '', '| Matière | Officiel Q / min | Chapitres | QCM / révision | Déficit numérique QCM | Format déclaré |',
                   '|---|---:|---:|---:|---:|---|']
        details=[]
        for subject,(name,count,minutes) in expected.items():
            lessons=grouped[subject];mcq=revision=0
            for d in lessons:
                m=d['metadata'];module=m['module_id']
                fmt=m.get('official_subject_format')
                if fmt and (fmt['questions'],fmt['minutes'],fmt['pass_percent'])!=(count,minutes,75):
                    errors.append(f'{module}: conflicting official_subject_format {fmt}')
                for q in d['questions']:
                    eligible='exam-mcq' in q.get('tags',[])
                    mcq+=eligible;revision+=not eligible
                    if eligible and (len(q['options'])!=4 or len(set(q['options']))!=4):errors.append(f'{module}: invalid MCQ')
                    if eligible and 'exam-subject:'+subject not in q['tags']:errors.append(f'{module}: subject tag mismatch')
                    if q['correct_answer'] not in q['options']:errors.append(f'{module}: missing answer')
                    if len(q.get('options_en',[]))!=len(q['options']) or not q.get('question_en') or not q.get('explanation_en'):errors.append(f'{module}: missing English')
                    if not q.get('source_references'):errors.append(f'{module}: missing source')
                if program=='PPL' and implemented.get(subject)!=(count,minutes): errors.append(f'PPL {subject}: engine quota mismatch')
            deficit=max(0,count-mcq)
            status='Concordant' if lessons and all(d['metadata'].get('official_subject_format') for d in lessons) else ('Moteur PPL concordant' if program=='PPL' else 'Non déclaré' if lessons else 'Matière absente')
            report.append(f'| {subject} — {name} | {count} / {minutes} | {len(lessons)} | {mcq} / {revision} | {deficit} | {status} |')
            rows.append(dict(program=program,subject=subject,name=name,official_questions=count,official_minutes=minutes,
                             chapters=len(lessons),mcq=mcq,revision=revision,minimum_numeric_shortfall=deficit,
                             coverage='partial_unverified' if lessons else 'absent',format_status=status))
            details += [f'### {name}', '', 'Chapitres : '+('; '.join('`'+d['metadata']['module_id']+'` — '+d['metadata']['chapter_title'] for d in lessons) or '**aucun**')+'.', '',
                        '**Revue à mener / approfondir :** '+REVIEW[subject], '']
        report += ['', 'Le déficit est seulement le nombre de QCM manquants pour un tirage sans répétition ; zéro ne prouve ni la couverture ni la difficulté attendue.', '']+details
    report += ['## Contrôles reproductibles', '',
               '```text', 'python scripts/seed/audit_exam_alignment.py', 'python scripts/seed/test_licence_content.py',
               'python scripts/seed/test_ppl_alignment.py', 'node --test web/frontend-react/tests/*.test.mjs', '```', '',
               f'{len(rows)} couples parcours/matière examinés. {len(errors)} incohérence(s) structurelle(s) ou de quota détectée(s).', '',
               '## Suite nécessaire', '',
               '1. Rattacher concepts et questions aux objectifs précis du programme applicable.',
               '2. Relire les faits, calculs, distracteurs et traductions ; traiter les priorités de chaque matière.',
               '3. Compléter les matières absentes et les exercices sur annexes avant les examens blancs professionnels.',
               '4. Valider la répartition par sous-matière, puis étendre le simulateur aux parcours disposant d’une banque suffisante.', '']
    if errors: report += ['## Incohérences détectées', '']+['- '+e for e in errors]
    directory=ROOT/'docs/licences'
    (directory/'exam-alignment-audit.md').write_text('\n'.join(report)+'\n',encoding='utf-8')
    (directory/'exam-alignment-audit.json').write_text(json.dumps(dict(checked_on=date.today().isoformat(),scope='France/EASA aeroplanes',
        regulatory_review='regulatory-level-review-2026-09-28.md',
        full_pedagogical_compliance_established=False,rows=rows,errors=errors),ensure_ascii=False,indent=2)+'\n',encoding='utf-8')
    print(json.dumps(dict(subjects_checked=len(rows),errors=errors,summary={p:dict(subjects_present=sum(r['chapters']>0 for r in rows if r['program']==p),subjects_required=sum(r['program']==p for r in rows)) for p in ['PPL','LAPL','CPL','ATPL','IR']})))
    return bool(errors)


if __name__=='__main__':
    raise SystemExit(run())
