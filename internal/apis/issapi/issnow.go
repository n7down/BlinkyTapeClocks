package issapi

type IssNow struct {
	IssPosition struct {
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
	} `json:"iss_position"`
	Message   string `json:"message"`
	Timestamp int    `json:"timestamp"`
}
