package database

import (
	"context"
	"fmt"
	"linebot/infras"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	err        error
	collection *mongo.Collection
)

// new mongodb instance
func NewDb(Opt *infras.Options) (*mongo.Collection, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@localhost:27017/",
		Opt.Config.MongoDB.User,
		Opt.Config.MongoDB.Password,
	)
	//1.建立連線
	if client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri).SetConnectTimeout(5*time.Second)); err != nil {
		log.Println(err)
		return nil, err
	}
	// 检查連接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//2.選擇數據庫
	collection = client.Database("local").Collection("test")

	return collection, nil
}
