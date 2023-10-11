package mongo

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackduh/goDomainDriven/aggregate"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	db      *mongo.Database
	patient *mongo.Collection
}

// mongoPatient is an internal type user to store PatientAggregate inside this repo
// no coupling to the aggregate
type mongoPatient struct {
	ID   uuid.UUID `bson:"id"`
	Name string    `bson:"name"`
}

func NewFromPatient(p aggregate.Patient) mongoPatient {
	return mongoPatient{
		ID:   p.GetID(),
		Name: p.GetName(),
	}
}

func (m mongoPatient) ToAggregate() aggregate.Patient {
	p := aggregate.Patient{}

	p.SetID(m.ID)
	p.SetName(m.Name)
	return p
}

func New(ctx context.Context, connectionString string) (*MongoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}

	db := client.Database("ddd")
	patients := db.Collection("patients")

	return &MongoRepository{
		db:      db,
		patient: patients,
	}, nil
}

func (mr *MongoRepository) Get(id uuid.UUID) (aggregate.Patient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := mr.patient.FindOne(ctx, bson.M{"id": id})

	var p mongoPatient
	if err := result.Decode(&p); err != nil {
		return aggregate.Patient{}, err
	}

	return p.ToAggregate(), nil
}

func (mr *MongoRepository) Add(p aggregate.Patient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	internal := NewFromPatient(p)

	_, err := mr.patient.InsertOne(ctx, internal)
	if err != nil {
		return err
	}
	return nil
}

func (mr *MongoRepository) Update(p aggregate.Patient) error {
	panic("to implement")
}
