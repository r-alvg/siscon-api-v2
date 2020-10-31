package routes
import (
	"context"
	"encoding/json"
	"net/http"
	"siscon/db"
	"time"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type healthResponse struct {
	Status       string                     `json:"status"`
	Message      string                     `json:"message"`
	Dependencies []healthResponseDependency `json:"dependencies"`
}

type healthResponseDependency struct {
	Name      string `json:"name"`
	Status    string `json:"status"`
	Reference string `json:"reference,omitempty"`
}

func Health() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", GetHealth)
	return r
}

// GetHealth implements `/health` endpoint handler
func GetHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dependencies := getDependencies()
	fullyFunctional := true

	for _, dependency := range dependencies {
		if dependency.Status != "OK" {
			fullyFunctional = false
		}
	}

	if fullyFunctional {
		_ = json.NewEncoder(w).Encode(healthResponse{
			Status:       "OK",
			Message:      "The application is fully functional.",
			Dependencies: dependencies,
		})
	} else {
		_ = json.NewEncoder(w).Encode(healthResponse{
			Status:       "FAIL",
			Message:      "The application is not fully functional.",
			Dependencies: dependencies,
		})
	}
}

func getDependencies() []healthResponseDependency {
	var dependenciesResponse []healthResponseDependency

	dependenciesResponse = append(dependenciesResponse, getMongoDbDependency())

	return dependenciesResponse
}

func getMongoDbDependency() healthResponseDependency {
	var status string

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	err := db.DB.Client().Ping(ctx, readpref.Primary())

	if err != nil {
		status = "FAIL"
		logrus.Error(err)
	} else {
		status = "OK"
	}

	return healthResponseDependency{
		Name:   "MongoDB",
		Status: status,
	}
}

