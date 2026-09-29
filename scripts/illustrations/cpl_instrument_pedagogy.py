"""CPL diagrams: pressure comparisons and common-source failures."""
from licence_diagrams import render
from svg_common import CYAN, AMBER, WHITE, MUTED


def pressure(s,t,en):
    s.text(65,215,t('Écrire les pressions reçues avant de conclure',
                    'Write the received pressures before concluding'),24,CYAN,weight=700)
    for x in [65,620]:s.rect(x,242,515,272,'#1c344b','#304961',12)
    s.text(83,282,t('Pt piégée, Ps libre','Trapped Pt, open Ps'),24,AMBER)
    s.text(83,327,t('Différence = Pt piégée − Ps actuelle',
                    'Difference = trapped Pt − current Ps'),22,WHITE)
    s.text(83,376,t('Ps baisse → différence augmente',
                    'Ps falls → difference increases'),22,CYAN)
    s.text(83,417,t('Ps augmente → différence baisse',
                    'Ps rises → difference decreases'),22,CYAN)
    s.wrap(83,468,t('Entrée ET drainage pitot obstrués.',
                     'Pitot inlet AND drain blocked.'),478,20,MUTED,line=25)
    s.text(638,282,t('Ps piégée, pitot libre','Trapped Ps, open pitot'),24,AMBER)
    s.text(638,327,t('Comparer au même instant :',
                     'Compare at the same instant:'),22,WHITE)
    s.text(638,369,t('Reçue : Pt − Ps piégée',
                     'Received: Pt − trapped Ps'),22,CYAN)
    s.text(638,409,t('Correcte : Pt − Ps extérieure',
                     'Correct: Pt − external Ps'),22,CYAN)
    s.wrap(638,461,t('Le sens du mouvement seul ne donne pas le signe de l’erreur.',
                      'Direction of movement alone does not give the error sign.'),470,20,MUTED,line=25)
    s.text(65,554,t('Pt ET Ps piégées → différence constante dans le modèle idéal.',
                    'BOTH Pt and Ps trapped → constant difference in the ideal model.'),22,AMBER)
    s.text(65,592,t('Hypothèses : sans fuite ni effet thermique ; hPa ≠ kt.',
                    'Assumptions: no leaks or thermal effects; hPa ≠ kt.'),20,MUTED)


def independence(s,t,en):
    s.text(65,215,t('Trois affichages ne garantissent pas trois sources',
                    'Three displays do not guarantee three sources'),24,CYAN,weight=700)
    s.rect(65,320,240,60,'#1c344b',AMBER,10)
    s.text(185,357,t('Prise Ps commune','Shared Ps port'),22,AMBER,anchor='middle')
    s.line(305,350,345,350,AMBER,3)
    s.line(345,290,345,410,AMBER,3)
    for y,label in [(260,'A'),(380,'B'),(490,'C')]:
        s.rect(410,y,280,60,'#1c344b','#304961',10)
        s.text(550,y+37,t('Calculateur ','Computer ')+label,22,WHITE,anchor='middle')
        s.rect(835,y,265,60,'#1c344b','#304961',10)
        s.text(967,y+37,t('Affichage ','Display ')+label,22,WHITE,anchor='middle')
        s.arrow(690,y+30,830,y+30,CYAN,3)
        if label != 'C':s.arrow(345,y+30,405,y+30,AMBER,3)
    s.rect(65,490,240,60,'#1c344b',CYAN,10)
    s.text(185,527,t('Autre prise Ps','Separate Ps port'),22,CYAN,anchor='middle')
    s.arrow(305,520,405,520,CYAN,3)
    s.text(65,466,t('Même défaut possible sur A et B malgré deux calculateurs.',
                    'A and B may share a fault despite separate computers.'),21,AMBER)
    s.text(65,593,t('C permet un recoupement ; son indépendance ne garantit pas son exactitude.',
                    'C provides a cross-check; independence does not guarantee accuracy.'),20,MUTED)


if __name__ == '__main__':
    render('CPL','CPL_PITOT_REASONING',('Pitot-statique : raisonner avec les pressions',
                                     'Pitot-static: reason from pressures'),pressure)
    render('CPL','CPL_SENSOR_INDEPENDENCE',('Capteurs : vérifier la source commune',
                                         'Sensors: check the shared source'),independence)
