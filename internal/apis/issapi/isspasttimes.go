package issapi

type IssPassTimes struct {
	Message string `json:"message"`
	Request struct {
		Altitude  int     `json:"altitude"`
		Datetime  int     `json:"datetime"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Passes    int     `json:"passes"`
	} `json:"request"`
	Response []struct {
		Duration int `json:"duration"`
		Risetime int `json:"risetime"`
	} `json:"response"`
}
