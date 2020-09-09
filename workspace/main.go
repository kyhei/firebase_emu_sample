package main

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
)

func main() {
	log.Println(os.Getenv("GCP_PROJECT_ID"))
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, os.Getenv("GCP_PROJECT_ID"))
	defer client.Close()
	if err != nil {
		log.Fatal(err.Error())
	}

	doc, _, err := client.Collection("users").Add(ctx, map[string]interface{}{
		"name": "kyhei",
		"age":  123,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(doc.ID)

	q := client.Collection("users").Select("name", "age")
	docs, err := q.Documents(ctx).GetAll()
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, doc := range docs {
		log.Println(doc.Data())
	}

}
