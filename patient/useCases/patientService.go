package useCases

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"siscon/patient/domain"
	"siscon/patient/infrastructure/request"
	"time"
)

type service struct {
	repository domain.Repository
}

func Service() *service {
	return &service{
		repository: domain.NewRepository(),
	}
}

func (s service)Create(request request.Patient) (string, error) {
	patient := bindToPatient(request)

	patient.ID = primitive.NewObjectIDFromTimestamp(time.Now())

	id , err := s.repository.Insert(patient)
	return id, err
}

func (s service)Find(id string) (domain.Patient, error) {
	return s.repository.FindByID(id)
}

func bindToPatient(request request.Patient) domain.Patient {
	return domain.Patient{
		Name:        request.Name,
		Cpf:         request.Cpf,
		Rg:          request.Rg,
		OrgEmitter:  request.OrgEmitter,
		DateOfBirth: request.DateOfBirth,
		Sex:         request.Sex,
		Ethnicity:   request.Ethnicity,
		Nationality: request.Nationality,
		Naturalness: request.Naturalness,
		Adress:      request.Adress,
		Schooling:   request.Schooling,
		Tel:         request.Tel,
		CellPhone:   request.CellPhone,
		Email:       request.Email,
		Profession:  request.Profession,
		Plans:       request.Plans,
		Status:      request.Status,
		CreateDate:  time.Now().Format("2006-01-02 15:04:05"),
		UpdateDate:  time.Now().Format("2006-01-02 15:04:05"),
	}

}

