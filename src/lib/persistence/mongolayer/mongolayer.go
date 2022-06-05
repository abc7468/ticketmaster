package mongolayer

import (
	"context"
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
		panic(err)
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
