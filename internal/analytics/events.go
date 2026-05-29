package analytics

// AnalyticsService gère l'ingestion et l'analyse des événements.
type AnalyticsService struct{}

func New() *AnalyticsService {
	return &AnalyticsService{}
}
