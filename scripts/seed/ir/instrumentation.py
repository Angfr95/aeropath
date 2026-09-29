"""IR 022: pressure references, setting errors and static-system diagnosis."""
from shared.authoring import author


def build():
    author('IR','instrumentation','022','Altimétrie et circuit statique','Altimetry and the static system',[
      ('Choisir la référence','Select the reference',
       '''### Identifier la grandeur avant de calculer

L’altimètre transforme une pression statique en indication selon une loi d’atmosphère et un calage. Il ne mesure pas directement la distance au sol. Distinguer trois références :

- **QNH :** altitude AMSL approximative, sans suppression des erreurs instrumentales ou de température.
- **QFE :** hauteur par rapport au repère associé au QFE, pas au terrain instantanément survolé.
- **1013,25 hPa :** altitude-pression servant aux niveaux de vol. Un même niveau ne correspond pas partout à une altitude vraie constante.

**Exemple résolu.** Aérodrome à 600 ft AMSL, avion 900 ft au-dessus de son repère. Erreurs négligées, QFE donne 900 ft et QNH donne 600 + 900 = **1 500 ft**. Un relief voisin à 1 200 ft AMSL ne laisserait que 300 ft de hauteur. L’indication QNH seule ne décrit donc pas la marge terrain.

**Vérifier sa compréhension.** Au QFE de cet aérodrome, que représente l’indication au-dessus d’une vallée ? Toujours la hauteur par rapport au repère QFE, pas au fond de vallée. Employer les transitions et calages publiés pour le vol concerné.''',
       '''### Identify the quantity before calculating

An altimeter converts static pressure into an indication using an atmospheric model and a setting. It does not directly measure distance above ground. Distinguish three references:

- **QNH:** approximate AMSL altitude, without eliminating instrument or temperature errors.
- **QFE:** height above the associated QFE datum, not terrain immediately below.
- **1013.25 hPa:** pressure altitude used for flight levels. A given level does not represent a constant true altitude everywhere.

**Worked example.** Aerodrome 600 ft AMSL, aircraft 900 ft above its datum. Neglecting errors, QFE gives 900 ft and QNH gives 600 + 900 = **1,500 ft**. Nearby terrain at 1,200 ft AMSL would leave only 300 ft of height. QNH indication alone therefore does not describe terrain margin.

**Check understanding.** Using that aerodrome’s QFE above a valley, what does the indication represent? Height above the QFE datum, not the valley floor. Use the transition and setting procedures published for the flight.'''),
      ('Quantifier une erreur de calage','Quantify a setting error',
       '''### Calculer, puis contrôler le signe

À pression statique constante, augmenter le calage augmente l’indication. Pour les exercices près du niveau de la mer, on impose **27 ft par hPa**. Ce coefficient n’est pas une correction universelle en altitude.

1. Écart de calage = calage affiché − QNH correct.
2. Erreur d’indication = écart signé × 27.
3. Indication au bon QNH = indication initiale − erreur.

**Exemple résolu.** QNH correct 1000 hPa, calage 1010 hPa, indication 3 000 ft. Écart : +10 hPa ; erreur : +270 ft ; indication corrigée : 3 000 − 270 = **2 730 ft**. Ce résultat n’est pas encore une altitude vraie corrigée de la température.

**Contrôle inverse.** Passer de 1000 à 1010 hPa à position inchangée ferait monter l’indication. Il est donc cohérent de retrancher 270 ft lorsque l’ancien calage était trop élevé. Ajouter 270 ft aggraverait l’erreur.

**Application autonome.** QNH correct 1008 hPa, calage 1000 hPa, indication 2 400 ft. Avant de lire la solution, prévoir si le bon calage donnera plus ou moins. Solution : écart −8 hPa, erreur −216 ft, indication corrigée **2 616 ft**.''',
       '''### Calculate, then check the sign

At constant static pressure, increasing the setting increases the indication. For these near-sea-level exercises, use **27 ft per hPa**. This coefficient is not a universal altitude correction.

1. Setting difference = selected setting minus correct QNH.
2. Indication error = signed difference × 27.
3. Correct-QNH indication = original indication minus error.

**Worked example.** Correct QNH 1000 hPa, setting 1010 hPa, indication 3,000 ft. Difference: +10 hPa; error: +270 ft; corrected indication: 3,000 − 270 = **2,730 ft**. This is not yet a temperature-corrected true altitude.

**Reverse check.** Selecting 1010 instead of 1000 hPa at unchanged position would raise the indication. Subtracting 270 ft when the old setting was too high is therefore consistent. Adding 270 ft would worsen the error.

**Independent application.** Correct QNH 1008 hPa, setting 1000 hPa, indication 2,400 ft. Before reading the solution, predict whether the correct setting gives more or less. Solution: difference −8 hPa, error −216 ft, corrected indication **2,616 ft**.'''),
      ('Pression et température réelles','Actual pressure and temperature',
       '''### Séparer deux causes d’erreur

Une variation de pression et une température différente du modèle sont deux problèmes distincts. Vers une pression plus basse sans recalage, l’indication peut surestimer l’altitude. Avec un QNH correct mais de l’air plus froid que le modèle, les surfaces de pression sont plus rapprochées : au-dessus du repère, l’altitude vraie peut aussi être inférieure à l’indiquée.

**Méthode.** Vérifier successivement le QNH pertinent, le repère de l’indication et la nécessité d’une correction de température selon la procédure applicable. Un bon QNH ne répond qu’à la première question.

**Situation.** Le QNH a été confirmé, mais une correction de basse température est requise sur une altitude publiée. Augmenter arbitrairement le QNH changerait la référence de toutes les indications. Utiliser la méthode publiée, identifier les altitudes concernées et respecter les exigences de coordination associées. Aucun coefficient de température universel n’est fourni ici.

**Erreur à éviter.** Les 2 730 ft de l’exemple précédent ne sont pas une hauteur disponible au-dessus de n’importe quel obstacle : le calcul corrige un calage, pas la température, le terrain ou une marge réglementaire.''',
       '''### Separate two causes of error

A pressure change and a temperature departure from the model are different problems. Moving towards lower pressure without resetting may cause altitude overreading. With correct QNH but air colder than the model, pressure surfaces are closer together: above the datum, true altitude may also be below indicated altitude.

**Method.** Check the relevant QNH, the indication’s datum, and whether the applicable procedure requires a temperature correction. Correct QNH answers only the first question.

**Scenario.** QNH has been confirmed, but a published altitude requires cold-temperature correction. Arbitrarily increasing QNH would change the reference for every indication. Use the published method, identify the affected altitudes and comply with associated coordination requirements. No universal temperature coefficient is supplied here.

**Avoid this mistake.** The preceding example’s 2,730 ft is not clearance above any obstacle: the calculation corrects a setting, not temperature, terrain or a regulatory margin.'''),
      ('Identifier une obstruction statique','Identify static blockage',
       '''### Raisonner à partir des pressions

Dans ce modèle classique, la statique est entièrement piégée et le Pitot reste libre. L’anémomètre compare la pression totale présente à la statique piégée. Comparer l’altitude actuelle à **l’altitude du blocage**, pas seulement constater une montée ou une descente. On suppose une atmosphère stable où la pression diminue avec l’altitude.

- **Au-dessus du blocage :** statique piégée trop élevée, différence de pression trop faible, sous-indication de vitesse par rapport à une statique correcte.
- **Au-dessous :** statique piégée trop faible, différence trop grande, surindication.
- **À la même altitude, pression extérieure inchangée :** pas d’erreur due à ce seul écart.

L’altimètre reste à l’indication du blocage si son calage est inchangé. Le variomètre revient vers zéro après son transitoire. Les trois instruments peuvent partager la même panne : trois indications ne sont pas nécessairement trois sources indépendantes.

**Contre-exemple.** Blocage à 5 000 ft, montée à 7 000 ft puis descente à 6 000 ft. Malgré la descente, l’anémomètre **sous-indique encore** : l’avion reste au-dessus du blocage. À 4 000 ft, il surindiquerait dans le même modèle.

**Diagnostic.** Recouper avec les sources indépendantes disponibles et appliquer la procédure de l’avion. Une statique de secours peut avoir ses propres corrections. Le radioaltimètre mesure une hauteur terrain ; il ne remplace pas automatiquement l’altitude barométrique publiée. Ce chapitre enseigne un raisonnement, pas une procédure d’urgence universelle.''',
       '''### Reason from the pressures

In this conventional model, static pressure is fully trapped and the pitot remains open. The ASI compares current total pressure with trapped static pressure. Compare current altitude with **blockage altitude**, rather than just noting climb or descent. Assume a stable atmosphere with pressure decreasing as altitude increases.

- **Above blockage altitude:** trapped static pressure too high, measured difference too small, airspeed underreading compared with correct static pressure.
- **Below:** trapped static pressure too low, difference too large, overreading.
- **Same altitude, unchanged outside pressure:** no error from this difference alone.

The altimeter stays at the blockage indication if its setting is unchanged. The VSI returns towards zero after its transient. All three may share the same fault: three indications do not necessarily represent three independent sources.

**Counterexample.** Blockage at 5,000 ft, climb to 7,000 ft, then descent to 6,000 ft. Despite descending, the ASI **still underreads**: the aircraft remains above blockage altitude. At 4,000 ft it would overread in this model.

**Diagnosis.** Cross-check available independent sources and follow the aircraft procedure. An alternate static source may require its own corrections. Radio altimetry measures terrain height; it does not automatically replace published barometric altitude. This chapter teaches reasoning, not a universal emergency procedure.''')],
'''2|À pression statique constante, calage de 1000 à 1010 hPa. Modèle près du niveau de la mer : 27 ft/hPa. Variation d’indication ?|At constant static pressure, setting changes from 1000 to 1010 hPa. Near-sea-level model: 27 ft/hPa. Indication change?|+270 ft~−270 ft~+27 ft~−27 ft|+270 ft~−270 ft~+27 ft~−27 ft|10 × 27 = 270 ft. Augmenter le calage augmente l’indication ; −270 inverse le sens et 27 oublie le facteur 10.|10 × 27 = 270 ft. Raising the setting raises the indication; −270 reverses the sign and 27 misses the factor of 10.
1|Aérodrome à 600 ft AMSL, avion 900 ft au-dessus de son repère. Erreurs négligées : quelle paire QNH/QFE attend-on ?|Aerodrome 600 ft AMSL, aircraft 900 ft above its datum. Neglecting errors, which QNH/QFE pair is expected?|1 500 ft / 900 ft~900 ft / 1 500 ft~900 ft / 300 ft~1 500 ft / 600 ft|1,500 ft / 900 ft~900 ft / 1,500 ft~900 ft / 300 ft~1,500 ft / 600 ft|QNH : 600 + 900 = 1 500 ft AMSL ; QFE : 900 ft au-dessus du repère. Inverser les deux confond altitude et hauteur.|QNH: 600 + 900 = 1,500 ft AMSL; QFE: 900 ft above the datum. Swapping them confuses altitude and height.
1|Au calage standard, FL 70 représente…|At standard setting, FL 70 represents…|7 000 ft d’altitude-pression à 1013,25 hPa~7 000 ft de hauteur terrain~Toujours 7 000 ft d’altitude vraie AMSL~7 000 ft au-dessus du repère QFE|7,000 ft pressure altitude at 1013.25 hPa~7,000 ft terrain height~Always 7,000 ft true AMSL altitude~7,000 ft above the QFE datum|Le niveau utilise la référence standard, pas une hauteur terrain ni une altitude vraie constante.|A flight level uses the standard reference, not terrain height or constant true altitude.
1|Au QFE d’un aérodrome, l’avion survole une vallée plus basse. À quoi l’indication se rapporte-t-elle ?|Using an aerodrome QFE, the aircraft crosses a lower valley. What is the indication referenced to?|Au repère de ce QFE~Au terrain sous l’avion~Au niveau moyen de la mer~À la référence 1013,25 hPa|That QFE datum~Terrain below the aircraft~Mean sea level~The 1013.25 hPa reference|Le repère QFE ne se déplace pas avec le relief. Cette indication n’est pas une mesure directe de hauteur terrain.|The QFE datum does not move with terrain. This indication is not a direct terrain-height measurement.
2|QNH réel 1000 hPa, calage 1010 hPa. Avec 27 ft/hPa, l’indication comparée au bon calage est…|Actual QNH 1000 hPa, setting 1010 hPa. At 27 ft/hPa, indication compared with correct setting is…|Environ 270 ft trop haute~Environ 270 ft trop basse~Environ 27 ft trop haute~Environ 27 ft trop basse|About 270 ft too high~About 270 ft too low~About 27 ft too high~About 27 ft too low|Écart : 1010 − 1000 = +10 hPa. Le calage trop élevé produit une surestimation de 270 ft ; 27 ft ne correspond qu’à 1 hPa.|Difference: 1010 − 1000 = +10 hPa. The excessive setting produces 270 ft overreading; 27 ft corresponds to only 1 hPa.
2|Une indication de 3 000 ft surestime de 270 ft celle au bon QNH. Indication corrigée du seul calage ?|An indication of 3,000 ft overreads the correct-QNH indication by 270 ft. Setting-corrected indication?|2 730 ft~3 270 ft~2 970 ft~3 000 ft|2,730 ft~3,270 ft~2,970 ft~3,000 ft|Retrancher la surestimation : 3 000 − 270 = 2 730 ft. Ajouter 270 aggrave l’erreur ; cela ne corrige pas la température.|Subtract the overreading: 3,000 − 270 = 2,730 ft. Adding 270 worsens the error; this does not correct temperature.
3|L’avion conserve la même indication et le même calage vers une pression plus basse. Quel risque anticiper ?|The aircraft keeps the same indication and setting towards lower pressure. Which risk should be anticipated?|Une altitude vraie plus basse que ne le suggère l’indication~Une altitude vraie nécessairement plus haute~Une hauteur terrain nécessairement constante~Une compensation automatique par l’altimètre|True altitude lower than the indication suggests~True altitude necessarily higher~Terrain height necessarily constant~Automatic compensation by the altimeter|Le calage devenu trop élevé surestime l’altitude. Une indication constante ne garantit ni altitude vraie ni marge terrain constantes.|The now-excessive setting overreads altitude. Constant indication guarantees neither constant true altitude nor terrain clearance.
3|QNH correct, air plus froid que le modèle, avion au-dessus du repère : pourquoi une correction publiée peut-elle rester nécessaire ?|Correct QNH, air colder than the model, aircraft above the datum: why may a published correction still be needed?|L’altitude vraie peut être inférieure à l’indiquée~Le QNH correct impose une altitude vraie identique~L’altitude vraie est nécessairement supérieure~Le QNH devient automatiquement un QFE|True altitude may be below indicated altitude~Correct QNH forces true and indicated altitude to match~True altitude must be higher~QNH automatically becomes QFE|Le QNH correct n’annule pas l’effet de la température sur l’épaisseur des couches de pression et ne transforme pas la référence en QFE.|Correct QNH does not remove temperature effects on pressure-layer thickness or turn the reference into QFE.
3|Une procédure requiert une correction de basse température. Quelle méthode retenir ?|A procedure requires cold-temperature correction. Which method should be used?|La méthode publiée pour les altitudes concernées~Augmenter arbitrairement le QNH de toute l’approche~Appliquer 27 ft par degré Celsius~Remplacer les altitudes par des hauteurs radioaltimétriques|The published method for the affected altitudes~Arbitrarily raise QNH for the whole approach~Apply 27 ft per degree Celsius~Replace altitudes with radio-altimeter heights|27 ft/hPa concerne le calage, pas la température. Changer de repère ou de grandeur ne remplace pas la correction publiée.|27 ft/hPa concerns setting, not temperature. Changing datum or quantity does not replace the published correction.
4|Statique piégée à 5 000 ft, calage inchangé, avion maintenant à 6 000 ft. Dans ce modèle classique, que fait l’altimètre ?|Static trapped at 5,000 ft, unchanged setting, aircraft now at 6,000 ft. In this conventional model, what does the altimeter do?|Il reste à l’indication du blocage~Il suit la montée avec un retard constant~Il revient progressivement à zéro~Il affiche l’altitude-pression correcte malgré le blocage|It stays at the blockage indication~It follows the climb with constant lag~It gradually returns to zero~It shows correct pressure altitude despite blockage|La pression appliquée et le calage restent fixes. Une indication figée ne doit pas être confondue avec un décalage qui continuerait à suivre la montée.|Applied pressure and setting remain fixed. A frozen indication is not an offset that continues to follow the climb.
4|La statique d’un variomètre classique est entièrement piégée. Après le transitoire, que peut-il indiquer malgré une montée réelle ?|A conventional VSI has fully trapped static pressure. After the transient, what may it indicate despite an actual climb?|Zéro~Le taux réel de montée~Le dernier taux de montée indéfiniment~Un taux de descente égal au taux réel de montée|Zero~Actual climb rate~The last climb rate indefinitely~A descent rate equal to actual climb rate|Après égalisation interne, il n’y a plus de variation de pression à mesurer. Ne pas confondre ce comportement avec l’altimètre figé.|After internal equalisation, no pressure change remains to measure. Do not confuse this behaviour with the frozen altimeter.
4|Statique piégée à 5 000 ft, Pitot libre. Après 7 000 ft, l’avion descend à 6 000 ft. Atmosphère stable : erreur de vitesse par rapport à une statique correcte ?|Static trapped at 5,000 ft, pitot open. After reaching 7,000 ft, the aircraft descends to 6,000 ft. Stable atmosphere: airspeed error compared with correct static pressure?|Sous-indication~Surindication~Aucune erreur car l’avion descend~Erreur indépendante de l’altitude du blocage|Underreading~Overreading~No error because the aircraft descends~Error independent of blockage altitude|6 000 ft reste au-dessus de 5 000 ft : statique piégée trop élevée, différence mesurée trop faible. Le sens de déplacement ne suffit pas à conclure.|6,000 ft remains above 5,000 ft: trapped static too high, measured difference too small. Direction of motion alone cannot determine the error.
4|Statique piégée à 5 000 ft, Pitot libre, avion à 4 000 ft. Atmosphère stable : erreur de vitesse par rapport à une statique correcte ?|Static trapped at 5,000 ft, pitot open, aircraft at 4,000 ft. Stable atmosphere: airspeed error compared with correct static pressure?|Surindication~Sous-indication~Aucune erreur car le Pitot est libre~Vitesse nécessairement figée à celle du blocage|Overreading~Underreading~No error because the pitot is open~Airspeed necessarily frozen at blockage value|Au-dessous du blocage, statique piégée trop faible et différence mesurée trop grande. Un Pitot libre ne corrige pas une référence statique erronée.|Below blockage altitude, trapped static is too low and measured difference too large. An open pitot cannot correct an erroneous static reference.
4|Après sélection d’une statique de secours, les indications changent légèrement. Quelle interprétation retenir ?|After selecting alternate static, indications change slightly. Which interpretation is appropriate?|La nouvelle source peut nécessiter les corrections du manuel~Tout changement prouve une panne de la source de secours~Les indications deviennent exactes sans correction~Il faut toujours modifier le QNH d’une valeur fixe|The new source may require manual corrections~Any change proves the alternate source has failed~Indications become exact without correction~QNH must always be changed by a fixed amount|La pression dépend de l’installation : utiliser ses corrections et sa procédure. Le changement seul ne prouve ni exactitude parfaite ni panne certaine.|Pressure depends on the installation: use its corrections and procedure. Change alone proves neither perfect accuracy nor definite failure.
4|Une procédure impose une altitude barométrique AMSL. Peut-on y substituer directement une hauteur radioaltimétrique fiable ?|A procedure specifies barometric AMSL altitude. Can a reliable radio-altimeter height be substituted directly?|Non, les repères diffèrent~Oui, une mesure fiable a toujours le même repère~Oui, si le QNH a été confirmé~Oui, en descente stabilisée|No, the datums differ~Yes, a reliable measurement always has the same datum~Yes, if QNH has been confirmed~Yes, in a steady descent|La radioaltimétrie mesure une hauteur terrain. Fiabilité de la mesure et pertinence de la grandeur sont deux vérifications différentes.|Radio altimetry measures terrain height. Measurement reliability and relevance of the quantity are separate checks.''',20,30,
      reference_checked_on='2026-09-27',
      learning_objective='Distinguer altitude et hauteur ; contrôler le signe d’une erreur de calage ; expliquer les effets du froid ; diagnostiquer une statique piégée par rapport à l’altitude du blocage.',
      learning_objective_en='Distinguish altitude from height; check the sign of a setting error; explain cold-air effects; diagnose trapped static pressure relative to blockage altitude.',
      prerequisites=['Pression statique et pression totale', 'Altitude, hauteur et niveau de vol', 'Calcul signé avec un coefficient fourni'],duration_minutes=45,
      extra_sources=[dict(title='FAA PHAK, chapitre 8 — Flight Instruments',url='https://www.faa.gov/sites/faa.gov/files/regulations_policies/handbooks_manuals/aviation/phak/10_phak_ch8.pdf'),
                     dict(title='FAA AIM 7-2 — erreurs altimétriques ; procédures US hors périmètre',url='https://www.faa.gov/air_traffic/publications/aim_html/chap7_section_2.html')])
