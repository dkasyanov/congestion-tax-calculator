package db

import (
	"congestion-calculator/config"
	"congestion-calculator/entity"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Repository struct {
	config *config.Config
	client *mongo.Client
}

const collectionName = "CityTaxRules"

func New(c *config.Config) *Repository {
	connStr := fmt.Sprintf("mongodb://%s:%s@%s:%s", c.Env.DbUser, c.Env.DbPassword, c.Env.DbHost, c.Env.DbPort)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connStr))
	if err != nil {
		panic(fmt.Sprintf("Cannot connect to DB: %+v", err))
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(fmt.Sprintf("DB is unreachable: %+v", err))
	}

	return &Repository{
		config: c,
		client: client,
	}
}

func (r *Repository) GetCityTaxRule(ctx context.Context, city string) (*entity.CityTaxRule, error) {
	collection := r.client.Database(r.config.Env.DbName).Collection(collectionName)

	var result entity.CityTaxRule
	err := collection.FindOne(ctx, bson.D{{"city", city}}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
