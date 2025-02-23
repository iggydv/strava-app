package athletes

import (
	"fmt"
	"strava-app/internal/strava/web/models"
)

func CalculateTotalDistance(stats *models.AthleteStats) float64 {
	fmt.Println("Calculating total distance")
	fmt.Println(stats.AllRideTotals.Distance)
	fmt.Println(stats.AllRunTotals.Distance)
	fmt.Println(stats.AllSwimTotals.Distance)
	return stats.AllRideTotals.Distance + stats.AllRunTotals.Distance + stats.AllSwimTotals.Distance
}
