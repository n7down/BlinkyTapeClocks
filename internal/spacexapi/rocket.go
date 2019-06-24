package spacexapi

type Rocket struct {
	ID             int    `json:"id"`
	Active         bool   `json:"active"`
	Stages         int    `json:"stages"`
	Boosters       int    `json:"boosters"`
	CostPerLaunch  int    `json:"cost_per_launch"`
	SuccessRatePct int    `json:"success_rate_pct"`
	FirstFlight    string `json:"first_flight"`
	Country        string `json:"country"`
	Company        string `json:"company"`
	Height         struct {
		Meters int     `json:"meters"`
		Feet   float64 `json:"feet"`
	} `json:"height"`
	Diameter struct {
		Meters float64 `json:"meters"`
		Feet   float64 `json:"feet"`
	} `json:"diameter"`
	Mass struct {
		Kg int `json:"kg"`
		Lb int `json:"lb"`
	} `json:"mass"`
	PayloadWeights []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Kg   int    `json:"kg"`
		Lb   int    `json:"lb"`
	} `json:"payload_weights"`
	FirstStage struct {
		Reusable       bool `json:"reusable"`
		Engines        int  `json:"engines"`
		FuelAmountTons int  `json:"fuel_amount_tons"`
		Cores          int  `json:"cores"`
		BurnTimeSec    int  `json:"burn_time_sec"`
		ThrustSeaLevel struct {
			KN  int `json:"kN"`
			Lbf int `json:"lbf"`
		} `json:"thrust_sea_level"`
		ThrustVacuum struct {
			KN  int `json:"kN"`
			Lbf int `json:"lbf"`
		} `json:"thrust_vacuum"`
	} `json:"first_stage"`
	SecondStage struct {
		Reusable       bool `json:"reusable"`
		Engines        int  `json:"engines"`
		FuelAmountTons int  `json:"fuel_amount_tons"`
		BurnTimeSec    int  `json:"burn_time_sec"`
		Thrust         struct {
			KN  int `json:"kN"`
			Lbf int `json:"lbf"`
		} `json:"thrust"`
		Payloads struct {
			Option1          string `json:"option_1"`
			Option2          string `json:"option_2"`
			CompositeFairing struct {
				Height struct {
					Meters float64 `json:"meters"`
					Feet   int     `json:"feet"`
				} `json:"height"`
				Diameter struct {
					Meters float64 `json:"meters"`
					Feet   float64 `json:"feet"`
				} `json:"diameter"`
			} `json:"composite_fairing"`
		} `json:"payloads"`
	} `json:"second_stage"`
	Engines struct {
		Number         int    `json:"number"`
		Type           string `json:"type"`
		Version        string `json:"version"`
		Layout         string `json:"layout"`
		EngineLossMax  int    `json:"engine_loss_max"`
		Propellant1    string `json:"propellant_1"`
		Propellant2    string `json:"propellant_2"`
		ThrustSeaLevel struct {
			KN  int `json:"kN"`
			Lbf int `json:"lbf"`
		} `json:"thrust_sea_level"`
		ThrustVacuum struct {
			KN  int `json:"kN"`
			Lbf int `json:"lbf"`
		} `json:"thrust_vacuum"`
		ThrustToWeight float64 `json:"thrust_to_weight"`
	} `json:"engines"`
	LandingLegs struct {
		Number   int    `json:"number"`
		Material string `json:"material"`
	} `json:"landing_legs"`
	FlickrImages []string `json:"flickr_images"`
	Wikipedia    string   `json:"wikipedia"`
	Description  string   `json:"description"`
	RocketID     string   `json:"rocket_id"`
	RocketName   string   `json:"rocket_name"`
	RocketType   string   `json:"rocket_type"`
}
