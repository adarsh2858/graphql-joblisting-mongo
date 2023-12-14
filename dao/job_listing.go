package dao

import (
	"context"
	"log"

	"github.com/adarsh2858/graphql-joblisting-mongo/graph/model"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	collectionName = "jobs"
)

// type JobListing struct {
// 	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
// 	Name    string             `json:"name"`
// 	Company string             `json:"company"`
// 	URL     string             `json:"url"`
// }

func init() {
	Connect()
}

func GetJobListings() []*model.JobListing {
	query := bson.D{{}}
	cursor, _ := mg.db.Collection(collectionName).Find(context.Background(), query)

	var jobListings []*model.JobListing
	if err := cursor.All(context.Background(), &jobListings); err != nil {
		log.Fatalf(err.Error())
	}

	return jobListings
}
