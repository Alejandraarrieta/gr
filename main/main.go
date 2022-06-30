package main

import (
	"context"
	//"encoding/json"
	"fmt"

	"github.com/gorilla/mux"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	//"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"

	f "./funciones"
)

var client *mongo.Client

func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	router := mux.NewRouter()
	router.HandleFunc("/product", f.CreateProduct).Methods("POST")
	router.HandleFunc("/order", f.CreateOrder).Methods("POST")
	router.HandleFunc("/buscarorder", f.FindOrder).Methods("POST")
	router.HandleFunc("/buscarprod", f.FindProd).Methods("POST")
	http.ListenAndServe(":9000", router)
}
