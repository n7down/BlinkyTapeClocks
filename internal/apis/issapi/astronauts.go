package issapi

type Astronauts struct {
	People []struct {
		Craft string `json:"craft"`
		Name  string `json:"name"`
	} `json:"people"`
	Number  int    `json:"number"`
	Message string `json:"message"`
}
