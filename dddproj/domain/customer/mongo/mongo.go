package mongo

import (
	"context"
	"dddproj/aggregate"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoRepository struct {
	db       *mongo.Database
	customer *mongo.Collection
}

func (mr *MongoRepository) Add(customer aggregate.Customer) error {
	//TODO implement me
	panic("implement me")
}
func (mr *MongoRepository) Update(customer aggregate.Customer) error {
	//TODO implement me
	panic("implement me")
}

// the internal type that is used to store a CustomerAffregate inside
type mongoCustomer struct {
	ID   uuid.UUID `bson:"_id,omitempty"`
	Name string    `bson:"name"`
}

func NewFromCustomer(c aggregate.Customer) mongoCustomer {
	return mongoCustomer{
		ID:   c.GetID(),
		Name: c.GetName(),
	}
}
func (m mongoCustomer) ToAggragte() aggregate.Customer {
	c := aggregate.Customer{}
	c.SetID(m.ID)
	c.SetName(m.Name)
	return c
}

func New(ctx context.Context, connectionString string) (*MongoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}
	db := client.Database("ddd")
	customers := db.Collection("customers")
	return &MongoRepository{db: db, customer: customers}, nil
}
func (mr *MongoRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result := mr.customer.FindOne(ctx, bson.M{"_id": id})

	var c mongoCustomer
	if err := result.Decode(&c); err != nil {
		return aggregate.Customer{}, err
	}
	return c.ToAggragte(), nil
}
func Add(mr *MongoRepository, c aggregate.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	internal := NewFromCustomer(c)
	_, err := mr.customer.InsertOne(ctx, internal)
	if err != nil {
		return err
	}
	return nil
}
