package mongodb

import (
	"context"
	models_rep "linebot/models/repository"
	rep_interface "linebot/repository/interface"
	"sync"

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

// insert one
func (rep *MemberRep) Insert(ctx context.Context, param *models_rep.Member) error {
	rep.mutex.Lock()
	defer rep.mutex.Unlock()
	if _, err := rep.db.InsertOne(context.Background(), *param); err != nil {
		return err
	}
	return nil
}

// find one
func (rep *MemberRep) Find(ctx context.Context, phone string) (*models_rep.Member, error) {
	rep.mutex.Lock()
	defer rep.mutex.Unlock()
	var result *models_rep.Member
	filter := bson.D{{Key: "phone", Value: phone}}
	if err := rep.db.FindOne(context.TODO(), filter).Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}

// update one
func (rep *MemberRep) Update(ctx context.Context, param *models_rep.Member) error {
	rep.mutex.Lock()
	defer rep.mutex.Unlock()
	updateFilter := bson.D{{Key: "phone", Value: param.Phone}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: param.Name}, {Key: "age", Value: param.Age}}}}
	_, err := rep.db.UpdateOne(context.TODO(), updateFilter, update)
	if err != nil {
		return err
	}
	return nil
}

// delete one
func (rep *MemberRep) Delete(ctx context.Context, phone string) error {
	rep.mutex.Lock()
	defer rep.mutex.Unlock()
	_, err := rep.db.DeleteOne(ctx, bson.M{"phone": phone})
	if err != nil {
		return err
	}
	return nil
}
