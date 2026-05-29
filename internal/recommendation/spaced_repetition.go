package recommendation

import (
	"math"
	"time"
)

// Card représente une carte de répétition espacée pour une question.
//
// 📖 DDIA Chapitre 1 : "Reliability"
//    L'algorithme SM-2 (SuperMemo 2) est utilisé depuis 1987.
//    Il est simple, efficace, et ne nécessite pas de base de données.
//    Chaque carte est indépendante : si une carte est corrompue,
//    les autres continuent de fonctionner (isolation des pannes).
//
// 📖 DDIA Chapitre 12 : "Derived Data"
//    Les paramètres SM-2 (EaseFactor, Interval, Repetitions) sont
//    des "données dérivées" calculées à partir de l'historique.
//    On pourrait les stocker dans Redis pour les récupérer plus vite.
type Card struct {
	QuestionID string    `json:"question_id"`
	EaseFactor  float64   `json:"ease_factor"`  // Facteur de facilité (1.3 - 3.0)
	Interval    int       `json:"interval"`     // Jours avant la prochaine révision
	Repetitions int       `json:"repetitions"`  // Nombre de révisions réussies consécutives
	NextReview  time.Time `json:"next_review"`  // Prochaine date de révision
	LastReview  time.Time `json:"last_review"`  // Dernière révision
}

// NewCard crée une nouvelle carte avec les valeurs par défaut.
func NewCard(questionID string) *Card {
	return &Card{
		QuestionID: questionID,
		EaseFactor:  2.5, // Valeur par défaut SM-2
		Interval:    0,
		Repetitions: 0,
		NextReview:  time.Now(),
	}
}

// Review enregistre une révision et met à jour les paramètres SM-2.
// quality : 0 (oubli total) à 5 (réponse parfaite)
//
// 📖 DDIA Chapitre 11 : "Stream Processing"
//    Chaque appel à Review() est un événement.
//    Si on stockait ces événements dans un log (NATS/Kafka),
//    on pourrait rejouer l'historique complet pour recalculer
//    l'état d'une carte en cas de bug.
func (c *Card) Review(quality int) {
	if quality < 0 {
		quality = 0
	}
	if quality > 5 {
		quality = 5
	}

	c.LastReview = time.Now()

	if quality < 3 {
		// La réponse est insuffisante → on réinitialise
		c.Repetitions = 0
		c.Interval = 1
	} else {
		// La réponse est correcte
		switch c.Repetitions {
		case 0:
			c.Interval = 1
		case 1:
			c.Interval = 6
		default:
			c.Interval = int(math.Round(float64(c.Interval) * c.EaseFactor))
		}
		c.Repetitions++
	}

	// Mise à jour du facteur de facilité (formule SM-2)
	c.EaseFactor = c.EaseFactor + (0.1 - float64(5-quality)*(0.08+float64(5-quality)*0.02))
	if c.EaseFactor < 1.3 {
		c.EaseFactor = 1.3
	}

	c.NextReview = time.Now().AddDate(0, 0, c.Interval)
}

// DueCards retourne les cartes dont la révision est due (NextReview <= now).
func DueCards(cards []*Card, now time.Time) []*Card {
	var due []*Card
	for _, c := range cards {
		if !c.NextReview.After(now) {
			due = append(due, c)
		}
	}
	return due
}

// SortByPriority trie les cartes par priorité de révision.
// Les cartes les plus en retard passent en premier.
func SortByPriority(cards []*Card, now time.Time) {
	// Simple tri par NextReview croissant (les plus urgentes d'abord)
	for i := 0; i < len(cards); i++ {
		for j := i + 1; j < len(cards); j++ {
			if cards[j].NextReview.Before(cards[i].NextReview) {
				cards[i], cards[j] = cards[j], cards[i]
			}
		}
	}
}
