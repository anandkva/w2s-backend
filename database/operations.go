package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

func InsertOne(collection string, document interface{}) (*mongo.InsertOneResult, error) {
	col := DB.Collection(collection)
	ctx, cancel := getCtx()
	defer cancel()
	return col.InsertOne(ctx, document)
}

func InsertMany(collection string, documents []interface{}) (*mongo.InsertManyResult, error) {
	col := DB.Collection(collection)
	ctx, cancel := getCtx()
	defer cancel()
	return col.InsertMany(ctx, documents)
}

func FindOne(collection string, filter interface{}, opts *options.FindOneOptions) *mongo.SingleResult {
	col := DB.Collection(collection)
	ctx, _ := getCtx()
	return col.FindOne(ctx, filter, opts)
}

func FindMany(collection string, filter interface{}, opts *options.FindOptions) (*mongo.Cursor, error) {
	col := DB.Collection(collection)
	ctx, cancel := getCtx()
	defer cancel()
	return col.Find(ctx, filter, opts)
}

func UpdateOne(collection string, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	col := DB.Collection(collection)
	ctx, cancel := getCtx()
	defer cancel()
	var uoOpts *options.UpdateOptions
	if len(opts) > 0 {
		uoOpts = opts[0]
	}
	return col.UpdateOne(ctx, filter, update, uoOpts)
}

func DeleteOne(collection string, filter interface{}) (*mongo.DeleteResult, error) {
	col := DB.Collection(collection)
	ctx, cancel := getCtx()
	defer cancel()
	return col.DeleteOne(ctx, filter)
}
