package valueObjects

type Address struct {
	Street       string `json:"street"`
	Number       int    `json:"number"`
	Neighborhood string `json:"neighborhood"`
	Cep          string `json:"cep"`
	City         string `json:"city"`
	State        string `json:"state"`
}
