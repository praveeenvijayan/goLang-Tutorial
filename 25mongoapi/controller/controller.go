package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/praveeenvijayan/goLang-Tutorial/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Add MongoDB Connection string

const connectionString = ""
const dbName = "netflix"
const colName = "watchlist"

//MOST  IMPORTANT
var collection *mongo.Collection

func init() {
	//client options
	clientOptions := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)

	checkErr(err)

	fmt.Println("MongoDb Connection Successfull")

	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collection instance is Ready")

}

// MOngoDB Helper -FILE

// inseret 1 record

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	checkErr(err)
	fmt.Println("Inserted 1 movie on DB with id:", inserted.InsertedID)
}

// update 1 record

func updateOneMovie(movieId string) {

	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	checkErr(err)

	fmt.Println("modified count: ", result.ModifiedCount)

}

func deleteOneMovie(movieId string) {

	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	DeleteCount, err := collection.DeleteOne(context.Background(), filter)

	checkErr(err)

	fmt.Println("Movie got Deleted: ", DeleteCount)
}

func deleteAllMovie() int64 {

	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}})

	checkErr(err)

	fmt.Println("NUmber of movies delete: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// get all movies from database

func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	checkErr(err)

	var movies []primitive.M

	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		checkErr(err)
		movies = append(movies, movie)
	}

	defer cur.Close(context.Background())
	return movies
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Actual controller - file
func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)

}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}
