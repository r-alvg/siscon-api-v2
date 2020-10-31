package valueObjects

type HealthPlan struct {
	Id     string 	`json:"id"`
	Name   string	`json:"name"`
	Status float64	`json:"status"`
}
