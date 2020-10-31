package domain

import "siscon/patient/domain/valueObjects"

type Medical struct {
	Cnpj     string                  `json:"cnpj"`
	Adhesion string                  `json:"adhesion"`
	Bank     string                  `json:"bank"`
	Agency   string                  `json:"agency"`
	Account  string                  `json:"account"`
	Plans    valueObjects.HealthPlan `json:"plans"`
}
