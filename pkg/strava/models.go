package strava

import "time"

type Athlete struct {
	Username      string    `json:"username"`
	Bio           string    `json:"bio"`
	Firstname     string    `json:"firstname"`
	Lastname      string    `json:"lastname"`
	Sex           string    `json:"sex"`
	Weight        float64   `json:"weight"`
	City          string    `json:"city"`
	State         string    `json:"state"`
	Country       string    `json:"country"`
	Premium       bool      `json:"premium"`
	CreatedAt     time.Time `json:"created_at"`
	Follower      *int      `json:"follower"`
	Friend        *int      `json:"friend"`
	ID            int       `json:"id"`
	Profile       string    `json:"profile"`
	ProfileMedium string    `json:"profile_medium"`
	ResourceState int       `json:"resource_state"`
	Summit        bool      `json:"summit"`
	UpdatedAt     time.Time `json:"updated_at"`
	BadgeTypeID   int       `json:"badge_type_id"`
}

type Totals struct {
	Distance         float64 `json:"distance"`
	AchievementCount int     `json:"achievement_count"`
	Count            int     `json:"count"`
	ElapsedTime      float64 `json:"elapsed_time"`
	ElevationGain    float64 `json:"elevation_gain"`
	MovingTime       float64 `json:"moving_time"`
}

type AthleteStats struct {
	BiggestRideDistance       float64 `json:"biggest_ride_distance"`
	BiggestClimbElevationGain float64 `json:"biggest_climb_elevation_gain"`
	RecentRideTotals          Totals  `json:"recent_ride_totals"`
	AllRideTotals             Totals  `json:"all_ride_totals"`
	RecentRunTotals           Totals  `json:"recent_run_totals"`
	AllRunTotals              Totals  `json:"all_run_totals"`
	RecentSwimTotals          Totals  `json:"recent_swim_totals"`
	AllSwimTotals             Totals  `json:"all_swim_totals"`
	YtdRideTotals             Totals  `json:"ytd_ride_totals"`
	YtdRunTotals              Totals  `json:"ytd_run_totals"`
	YtdSwimTotals             Totals  `json:"ytd_swim_totals"`
}
