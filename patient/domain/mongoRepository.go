package domain

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"siscon/db"
)

type Repository struct {
	db.Repository
}

const paymentCollection = "patient"

func (r Repository) FindByID(ID string) (Patient, error) {
	var patient Patient

	objId, _ := primitive.ObjectIDFromHex(ID)

	if err := r.FindOne(context.TODO(), bson.M{"_id": objId}).Decode(&patient); err != nil {
		logrus.Errorf("Error on finding payment cause: %v", err)
		return Patient{}, errors.New("patient not found")
	}

	return patient, nil
}

func (r Repository) Insert(patient Patient) (string, error) {
	result, err := r.InsertOne(context.Background(), patient)
	if err != nil {
		logrus.Error(err)
		return "", err
	}
	objId, _ := result.InsertedID.(primitive.ObjectID)
	return objId.Hex(), nil
}

func NewRepository() Repository {
	return Repository{db.NewRepository(paymentCollection)}
}
