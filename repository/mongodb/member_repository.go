package mongodb

import (
	"context"
	models_rep "linebot/models/repository"
	rep_interface "linebot/repository/interface"
	"sync"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MemberRep struct {
	mutex sync.Mutex
	db    *mongo.Collection
}

func NewMemberRep(db *mongo.Collection) rep_interface.IMemberRep {
	return &MemberRep{
		db: db,
	}
}

// insert user info
func (rep *MemberRep) Insert(ctx context.Context, param *linebot.Event) error {
	rep.mutex.Lock()
	defer rep.mutex.Unlock()
	if _, err := rep.db.InsertOne(context.Background(), *param); err != nil {
		return err
	}
	return nil
}

// find user infos
func (rep *MemberRep) FindAll(ctx context.Context, userId string) (*[]models_rep.LineEvent, error) {
	rep.mutex.Lock()
	defer rep.mutex.Unlock()
	var lineEvent models_rep.LineEvent
	var result []models_rep.LineEvent
	filter := bson.M{"source.userid": userId, "type": "message"}
	cur, err := rep.db.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		if err := cur.Decode(&lineEvent); err != nil {
			return nil, err
		}
		result = append(result, lineEvent)
	}
	return &result, nil
}
