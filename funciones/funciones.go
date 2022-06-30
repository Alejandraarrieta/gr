package funciones

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	m "gr/models"
)

var client *mongo.Client

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var prod m.Product
	_ = json.NewDecoder(r.Body).Decode(&prod)
	collection := client.Database("productsDB").Collection("products")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, prod)
	json.NewEncoder(w).Encode(result)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var ord m.Order
	_ = json.NewDecoder(r.Body).Decode(&ord)
	collection := client.Database("ordersDB").Collection("orders")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, ord)
	json.NewEncoder(w).Encode(result)
}

func FindProd(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var busca m.BuscarProd
	_ = json.NewDecoder(r.Body).Decode(&busca)
	name := busca.Name
	id := busca.Id
	var products []m.Product
	collection := client.Database("productsDB").Collection("products")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{"productname": name, "_id": id}
	options := options.Find()
	cursor, err := collection.Find(ctx, filter, options)
	if err != nil {
		fmt.Println("Error", err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var product m.Product
		cursor.Decode(&product)
		products = append(products, product)
	}
	if err := cursor.Err(); err != nil {
		fmt.Println("SearchUser", err)
		return
	}
	jsonResult, err := json.Marshal(products)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResult)
}

func FindOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var busca m.Buscar
	var orders []m.Order
	_ = json.NewDecoder(r.Body).Decode(&busca)
	email := busca.Email
	status := busca.Status
	collection := client.Database("ordersDB").Collection("orders")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{"sellername": email, "statusorder": status}
	options := options.Find()
	cursor, err := collection.Find(ctx, filter, options)
	if err != nil {
		fmt.Println("Err", err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var order m.Order
		cursor.Decode(&order)
		orders = append(orders, order)
	}
	if err := cursor.Err(); err != nil {
		fmt.Println("err", err)
		return
	}
	jsonResult, err := json.Marshal(orders)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResult)
}
