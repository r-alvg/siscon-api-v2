package helpers

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, msg interface{}) {
	RespondwithJSON(w, code, map[string]interface{}{"message": msg})
}
func RespondWithoutPayload(w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
}

func RespondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		logrus.Error("Error ao parsear payload para json cause: ", err)
		RespondwithJSON(w, http.StatusInternalServerError, "")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}

