package infrastructure

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"net/http"
	"siscon/helpers"
	"siscon/patient/infrastructure/request"
	"siscon/patient/infrastructure/response"
	"siscon/patient/useCases"
)

// Patiente is a router for payment endpoints
func PatientController() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", Create)
	r.Get("/{id}", FindById)
	return r
}

var validate = validator.New()

func FindById(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "id")

	logrus.Infof("finding payment with id: %v", ID)
	payment, err := useCases.Service().Find(ID)
	if err != nil {
		helpers.RespondWithoutPayload(w, http.StatusNotFound)
		return
	}
	helpers.RespondwithJSON(w, http.StatusOK, payment)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var paymentRequest request.Patient

	logrus.Infof("creating new patient.")
	err := json.NewDecoder(r.Body).Decode(&paymentRequest)
	if err != nil {
		logrus.Error(err)
	}

	er := validate.Struct(paymentRequest)
	if er != nil {
		helpers.RespondwithJSON(w, http.StatusBadRequest, helpers.ValidateErros(er.(validator.ValidationErrors)))
		return
	}

	logrus.Info("validated request, creating patient.")
	pId, err := useCases.Service().Create(paymentRequest)

	if err != nil {
		logrus.Errorf("Error on creating patient, cause: %v", err)
		helpers.RespondWithError(w, http.StatusUnprocessableEntity, "Error on creating patient")
	}

	helpers.RespondwithJSON(w, http.StatusOK, response.Model{
		Payload: struct { Id string `json:"id"` }{ Id: pId},
		Links:   []response.Link{{Href: fmt.Sprintf("siscon/v2/patient/%s", pId ), Rel: "patient.get"}},
	})
}
