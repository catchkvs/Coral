package repo

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/catchkvs/Coral/pkg/config"
	"google.golang.org/api/iterator"
	"log"
)

var client *firestore.Client
var inMemoryStore map[string] interface{}
var ctx = context.Background()
func init() {
	connector := config.GetProperty("coral.store.connector")
	if connector == "inmemory" {
		// TODO: implement inmemory connector
	} else {

		projectId := config.GetProperty("coral.googlecloud.projectid")
		// Get a Firestore client.
		client, _ = firestore.NewClient(ctx, projectId)
	}

}

func GetClient() *firestore.Client {
	return client
}


func Add(collection string, data interface{}) string {
	docRef, wr, err := client.Collection(collection).Add(ctx, data)
	if err != nil {
		log.Fatalf("Failed adding aturing: %v", err)
	}
	log.Println("update time: ", wr.UpdateTime)
	return docRef.ID
}

func AddWithId(collection string, id string, data interface{}) {
	wr, err := client.Collection(collection).Doc(id).Set(ctx, data)
	if err != nil {
		log.Fatalf("Failed adding aturing: %v", err)
	}
	log.Println("update time: ", wr.UpdateTime)
}

func UpdateDoc(collection string, id string, fields map[string]interface{}) {
	_, err := client.Collection(collection).Doc(id).Set(ctx, fields, firestore.MergeAll)
	if err != nil {
		log.Fatalf("Failed adding aturing: %v", err)
	}
}

func Get(collection string, id string) *firestore.DocumentSnapshot {
	docSnap, err := client.Collection(collection).Doc(id).Get(ctx)
	if err != nil {

		log.Println("Failed to get item: %v", err)
		return nil
	}
	return docSnap
}

func FindOneByField(collection, fieldName, fieldValue string) *firestore.DocumentSnapshot {
	query := client.Collection(collection).Where(fieldName, "==", fieldValue)
	iter := query.Documents(ctx)
	doc, err := iter.Next()
	if err == iterator.Done {
		return doc
	}
	if err != nil {
		return nil
	}
	return doc
}

