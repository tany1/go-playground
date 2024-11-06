package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Booke struct {
	Title, Author string
}

func ReadSampleMflix(client *mongo.Client, title string) {
	coll := client.Database("sample_mflix").Collection("movies")

	var result bson.M

	err := coll.FindOne(context.TODO(), bson.D{{"title", title}}).Decode(&result)

	if err == mongo.ErrNoDocuments {
		log.Println("no documents found with title: ", title)
		return
	}
	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))
}

func AddBook(client *mongo.Client, title, author string) {
	coll := client.Database("sample_db").Collection("books")

	doc := Booke{title, author}

	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		panic(err)
	}
	fmt.Println("inserted document with _id: ", result.InsertedID)
}

func DeleteSampleMflix(client *mongo.Client, title string) {
	coll := client.Database("sample_mflix").Collection("movies")

	result, err := coll.DeleteOne(context.TODO(), bson.D{{"title", title}})
	if err != nil {
		panic(err)
	}
	fmt.Println("deleted document with title: ", title, " deleted count: ", result.DeletedCount)
}

func UpdateSampleMflix(client *mongo.Client) {
	coll := client.Database("sample_mflix").Collection("movies")

	filter := bson.D{{"_id", 2158}}
	update := bson.D{{"$set", bson.D{{"name", "Mary Wollstonecraft Shelley"},
		{"role", "Marketing Director"}}}, {"$inc", bson.D{{"bonus", 2000}}}}

	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Documents matched: %v\n", result.MatchedCount)
	fmt.Printf("Documents updated: %v\n", result.ModifiedCount)
}

type Podcast struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Title  string             `bson:"title,omitempty"`
	Author string             `bson:"author,omitempty"`
	Tags   []string           `bson:"tags,omitempty"`
}

type Episode struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Podcast     primitive.ObjectID `bson:"podcast,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Description string             `bson:"description,omitempty"`
	Duration    int                `bson:"duration,omitempty"`
}

func InsertPodcast(client *mongo.Client) {
	database := client.Database("sample_db")
	podcastCollection := database.Collection("podcasts")
	episodeCollection := database.Collection("episodes")

	podcasts := []interface{}{
		Podcast{Title: "The Reader Podcast", Author: "John Doe"},
		Podcast{Title: "The Polyglot Developer Podcast", Author: "Nic Raboy", Tags: []string{"development", "programming", "coding"}},
	}

	podcastResult, err := podcastCollection.InsertMany(context.TODO(), podcasts)
	if err != nil {
		panic(err)
	}
	fmt.Println("inserted documents with _id: ", podcastResult.InsertedIDs)

	readerID := podcastResult.InsertedIDs[0].(primitive.ObjectID)
	polyglotID := podcastResult.InsertedIDs[1].(primitive.ObjectID)

	readerEpisodes := []interface{}{
		Episode{Title: "Reader Episode #1", Description: "This is the first episode", Duration: 25, Podcast: readerID},
		Episode{Title: "Reader Episode #2", Description: "This is the second episode", Duration: 32, Podcast: readerID},
	}
	polyglotEpisodes := []interface{}{
		Episode{Title: "Polyglot Episode #1", Description: "This is the first episode", Duration: 25, Podcast: polyglotID},
		Episode{Title: "Polyglot Episode #2", Description: "This is the second episode", Duration: 32, Podcast: polyglotID},
	}

	_, err = episodeCollection.InsertMany(context.TODO(), readerEpisodes)
	if err != nil {
		panic(err)
	}
	_, err = episodeCollection.InsertMany(context.TODO(), polyglotEpisodes)
	if err != nil {
		panic(err)
	}
}

func FindAllEpisodes(client *mongo.Client) {
	database := client.Database("sample_db")
	podcastCollection := database.Collection("podcasts")
	ctx := context.TODO()

	pipeline := mongo.Pipeline{
		{{
			"$lookup", bson.D{
				{"from", "episodes"},
				{"localField", "_id"},
				{"foreignField", "podcast"},
				{"as", "episodes"},
			},
		}},
	}

	var results []struct {
		ID       primitive.ObjectID `bson:"_id,omitempty"`
		Title    string             `bson:"title,omitempty"`
		Author   string             `bson:"author,omitempty"`
		Tags     []string           `bson:"tags,omitempty"`
		Episodes []Episode          `bson:"episodes,omitempty"`
	}
	// var results []bson.M

	cursor, err := podcastCollection.Aggregate(ctx, pipeline)
	defer cursor.Close(ctx)
	if err != nil {
		panic(err)
	}

	if err = cursor.All(ctx, &results); err != nil {
		panic(err)
	}

	for _, result := range results {
		fmt.Println(result)
	}
}
