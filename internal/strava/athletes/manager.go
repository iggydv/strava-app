package athletes

import (
	"strava-app/internal/strava/web/models"
)

func CalculateTotalDistance(stats *models.AthleteStats) float64 {
	return stats.AllRideTotals.Distance + stats.AllRunTotals.Distance + stats.AllSwimTotals.Distance
}
