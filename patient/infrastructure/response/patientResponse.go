package response

import "siscon/patient/domain/valueObjects"

type Patient struct {
	ID          int                        `json:"id"`
	Name        string                     `json:"name"`
	Cpf         string                     `json:"cpf"`
	Rg          string                     `json:"rg"`
	OrgEmitter  string                     `json:"org_emitter"`
	DateOfBirth string                     `json:"dateOfBirth"`
	Sex         string                    `json:"sex"`
	Ethnicity   string                    `json:"ethnicity"`
	Nationality string                    `json:"nationality"`
	Naturalness string                    `json:"naturalness"`
	Adress      valueObjects.Address      `json:"adress"`
	Schooling   string                    `json:"schooling"`
	Tel         string                    `json:"tel"`
	CellPhone   string                    `json:"cellPhone"`
	Email       string                    `json:"email"`
	Profession  interface{}               `json:"profession"`
	Plans       []valueObjects.HealthPlan `json:"plans"`
	Status      string                    `json:"status"`
	CreateDate  interface{}               `json:"createDate"`
	UpdateDate  interface{}               `json:"updateDate"`
}
