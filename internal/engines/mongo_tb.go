package engines

import (
	"context"
	"crawlab-sdk/internal/interfaces"
	"github.com/crawlab-team/go-trace"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

type MongoTb struct {
	interfaces.MongoTb
	_TB *qmgo.Collection
}

func (my *MongoTb) Insert(data interface{}) error {
	ctx := context.Background()
	_, err := my._TB.InsertOne(ctx, data)
	return trace.Error(err)
}


func (my *MongoTb) Delete(condi interface{}) error {
	ctx := context.Background()
	_,err := my._TB.RemoveAll(ctx,condi)
	return trace.Error(err)
}

func (my *MongoTb) Update(replacement, condi interface{}) error {
	m := bson.M{"$set": replacement}
	ctx := context.Background()
	_,err := my._TB.UpdateAll(ctx,condi,m)
	return trace.Error(err)
}

func (my *MongoTb) Upsert(replacement, condi interface{}) error {
	ctx := context.Background()
	_, err := my._TB.Upsert(ctx, condi, replacement)
	return trace.Error(err)
}

func (my *MongoTb) FindOne(result interface{},condi interface{}) error {
	ctx := context.Background()
	err := my._TB.Find(ctx,condi).One(result)
	return trace.Error(err)
}

func (my *MongoTb) FindALL(result interface{},condi interface{}) error {
	ctx := context.Background()
	err := my._TB.Find(ctx,condi).All(result)
	return trace.Error(err)
}
