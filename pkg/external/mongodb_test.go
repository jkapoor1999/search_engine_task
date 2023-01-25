package external

import (
	"reflect"
	"search_engine_task/pkg/models"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

func TestNewMongoDB(t *testing.T) {
	type args struct {
		coll *mongo.Collection
	}
	tests := []struct {
		name string
		args args
		want *MongoDB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMongoDB(tt.args.coll); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMongoDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMongoDB_GetAllCollection(t *testing.T) {
	tests := []struct {
		name string
		c    *MongoDB
		want []bson.M
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.GetAllCollection(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MongoDB.GetAllCollection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMongoDB_InsertOnePage(t *testing.T) {
	type args struct {
		ctx     context.Context
		newPage models.Page
	}
	tests := []struct {
		name    string
		c       *MongoDB
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.InsertOnePage(tt.args.ctx, tt.args.newPage); (err != nil) != tt.wantErr {
				t.Errorf("MongoDB.InsertOnePage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
