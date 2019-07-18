package spacexapi

import "time"

type Payload struct {
	PayloadID      string   `json:"payload_id"`
	NoradID        []int    `json:"norad_id"`
	Reused         bool     `json:"reused"`
	Customers      []string `json:"customers"`
	Nationality    string   `json:"nationality"`
	Manufacturer   string   `json:"manufacturer"`
	PayloadType    string   `json:"payload_type"`
	PayloadMassKg  int      `json:"payload_mass_kg"`
	PayloadMassLbs float64  `json:"payload_mass_lbs"`
	Orbit          string   `json:"orbit"`
	OrbitParams    struct {
		ReferenceSystem string    `json:"reference_system"`
		Regime          string    `json:"regime"`
		Longitude       int       `json:"longitude"`
		SemiMajorAxisKm float64   `json:"semi_major_axis_km"`
		Eccentricity    float64   `json:"eccentricity"`
		PeriapsisKm     float64   `json:"periapsis_km"`
		ApoapsisKm      float64   `json:"apoapsis_km"`
		InclinationDeg  float64   `json:"inclination_deg"`
		PeriodMin       float64   `json:"period_min"`
		LifespanYears   int       `json:"lifespan_years"`
		Epoch           time.Time `json:"epoch"`
		MeanMotion      float64   `json:"mean_motion"`
		Raan            float64   `json:"raan"`
	} `json:"orbit_params"`
}
