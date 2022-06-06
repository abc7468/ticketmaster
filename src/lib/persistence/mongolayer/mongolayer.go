package mongolayer

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"ticketmaster/src/lib/persistence"
)

const (
	DB     = "myevents"
	USERS  = "users"
	EVENTS = "events"
)

type MongoDBLayer struct {
	client *mongo.Client
}

func NewMongoDBLayer(connection string) (*MongoDBLayer, error) {
	clientOptions := options.Client().ApplyURI(connection)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}
	return &MongoDBLayer{
		client: client,
	}, nil
}

func (mgoLayer *MongoDBLayer) AddEvent(e persistence.Event) ([]byte, error) {
	// session pool에서 데이터베이스 세션을 가져오기 위함
	sess, err := mgoLayer.client.StartSession()
	if err != nil {
		return nil, err
	}
	defer sess.EndSession(context.TODO())
	result, err := sess.WithTransaction(
		context.TODO(),
		func(sessCtx mongo.SessionContext) (interface{}, error) {
			coll := mgoLayer.client.Database(DB).Collection(EVENTS)
			e.ID = primitive.NewObjectID()
			res, err := coll.InsertOne(sessCtx, e)
			if err != nil {
				return nil, err
			}
			return res, err
		},
	)
	if err != nil {
		return nil, err
	}
	return result.([]byte), nil
}

func (mgoLayer *MongoDBLayer) FindEvent(id []byte) (persistence.Event, error) {
	sess, _ := mgoLayer.client.StartSession()
	defer sess.EndSession(context.TODO())
	result, err := sess.WithTransaction(
		context.TODO(),
		func(sessCtx mongo.SessionContext) (interface{}, error) {
			e := persistence.Event{}
			err := mgoLayer.client.Database(DB).Collection(EVENTS).FindOne(sessCtx, bson.D{{"_id", string(id)}}).Decode(&e)
			if err != nil {
				return nil, err
			}
			return e, nil
		},
	)
	return result.(persistence.Event), err
}

func (mgoLayer *MongoDBLayer) FindEventByName(name string) (persistence.Event, error) {
	sess, _ := mgoLayer.client.StartSession()
	defer sess.EndSession(context.TODO())
	result, err := sess.WithTransaction(
		context.TODO(),
		func(sessCtx mongo.SessionContext) (interface{}, error) {
			e := persistence.Event{}
			err := mgoLayer.client.Database(DB).Collection(EVENTS).FindOne(sessCtx, bson.D{{"name", name}}).Decode(&e)
			if err != nil {
				return nil, err
			}
			return e, nil
		},
	)
	return result.(persistence.Event), err
}
func (mgoLayer *MongoDBLayer) FindAllAvailableEvents() ([]persistence.Event, error) {
	sess, _ := mgoLayer.client.StartSession()
	defer sess.EndSession(context.TODO())
	result, err := sess.WithTransaction(
		context.TODO(),
		func(sessCtx mongo.SessionContext) (interface{}, error) {
			es := []persistence.Event{}
			cur, err := mgoLayer.client.Database(DB).Collection(EVENTS).Find(sessCtx, bson.D{{}})
			for cur.Next(context.TODO()) {
				e := persistence.Event{}
				err := cur.Decode(&e)
				if err != nil {
					return nil, err
				}
				es = append(es, e)
			}
			if err != nil {
				return nil, err
			}
			return es, nil
		},
	)
	return result.([]persistence.Event), err
}
