package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

var (
	mongoURI     = "mongodb+srv://thuduyenhocmsc:D051199d@cluster0.dzmppq0.mongodb.net/"
	dbName       = "bank"
	collection   = "accounts"
	transaction1 = "user1"
	transaction2 = "user2"
)

type Transaction struct {
	From   string
	To     string
	Amount float64
}

// WithTransactionExample is an example of using the Session.WithTransaction function.
func WithTransactionExample(ctx context.Context) error {
	// For a replica set, include the replica set name and a seedlist of the members in the URI string; e.g.
	//uri := "mongodb://mongodb0.example.com:27017,mongodb1.example.com:27017/?replicaSet=myRepl"
	// For a sharded cluster, connect to the mongos instances; e.g.
	// uri := "mongodb://mongos0.example.com:27017,mongos1.example.com:27017/"
	// uri := mtest.ClusterURI()
	uri := "mongodb+srv://thuduyenhocmsc:D051199d@cluster0.dzmppq0.mongodb.net/"

	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return err
	}
	defer func() { _ = client.Disconnect(ctx) }()

	// Prereq: Create collections.
	wcMajority := writeconcern.New(writeconcern.WMajority(), writeconcern.WTimeout(1*time.Second))
	wcMajorityCollectionOpts := options.Collection().SetWriteConcern(wcMajority)
	fooColl := client.Database("mydb1").Collection("foo", wcMajorityCollectionOpts)
	barColl := client.Database("mydb1").Collection("bar", wcMajorityCollectionOpts)

	// Step 1: Define the callback that specifies the sequence of operations to perform inside the transaction.
	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		// Important: You must pass sessCtx as the Context parameter to the operations for them to be executed in the
		// transaction.
		if _, err := fooColl.InsertOne(sessCtx, bson.D{{"abc", 1}}); err != nil {
			return nil, err
		}
		if _, err := barColl.InsertOne(sessCtx, bson.D{{"xyz", 999}}); err != nil {
			return nil, err
		}

		return nil, nil
	}

	// Step 2: Start a session and run the callback using WithTransaction.
	session, err := client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	result, err := session.WithTransaction(ctx, callback)
	if err != nil {
		return err
	}
	log.Printf("result: %v\n", result)
	return nil
}

// Path: lectures/lecture06/go-quickstart/TransactionExample.go

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	coll := client.Database(dbName).Collection(collection)

	// Tạo dữ liệu ngân hàng
	_, err = coll.InsertMany(ctx, []interface{}{
		bson.M{"user": transaction1, "balance": 1000.0},
		bson.M{"user": transaction2, "balance": 2000.0},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Thực hiện giao dịch
	transactionAmount := 500.0

	session, err := client.StartSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.EndSession(ctx)

	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		err := sc.StartTransaction()
		if err != nil {
			return err
		}

		// Trừ tiền từ tài khoản 1 và cộng vào tài khoản 2
		_, err = coll.UpdateOne(sc, bson.M{"user": transaction1}, bson.M{"$inc": bson.M{"balance": -transactionAmount}})
		if err != nil {
			sc.AbortTransaction(sc)
			return err
		}

		_, err = coll.UpdateOne(sc, bson.M{"user": transaction2}, bson.M{"$inc": bson.M{"balance": transactionAmount}})
		if err != nil {
			sc.AbortTransaction(sc)
			return err
		}

		err = sc.CommitTransaction(sc)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	// Kiểm tra số dư sau giao dịch
	user1Balance := getUserBalance(ctx, coll, transaction1)
	user2Balance := getUserBalance(ctx, coll, transaction2)

	fmt.Printf("Số dư sau giao dịch - User1: %.2f, User2: %.2f\n", user1Balance, user2Balance)
}

func getUserBalance(ctx context.Context, coll *mongo.Collection, user string) float64 {
	var result struct {
		Balance float64 `bson:"balance"`
	}
	filter := bson.M{"user": user}
	err := coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result.Balance
}
