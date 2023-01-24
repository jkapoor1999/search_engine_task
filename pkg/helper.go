package pkg

import (
	"context"
	"os"
	"search_engine_task/pkg/dbconn"

	"go.mongodb.org/mongo-driver/bson"
)


func GetAllCollection() []bson.M {

	cur, err := dbconn.Dbconn().Find(context.Background(), bson.D{})

	if err != nil {

		os.Exit(1)

	}

	var temp []bson.M

	for cur.Next(context.Background()) {

		var p bson.M

		err := cur.Decode(&p)

		if err != nil {

			os.Exit(1)

		}

		temp = append(temp, p)

	}

	defer cur.Close(context.Background())

	return temp

}
