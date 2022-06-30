package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id             primitive.ObjectID `json:"_id,omitempty"`
	ProductName    string             `json:"productName,omitempty" bson:"productname,omitempty"`
	Port           string             `json:"port,omitempty" bson:"port,omitempty"`
	Price          string             `json:"price,omitempty" bson:"price,omitempty"`
	TypeCoin       string             `json:"typeCoin,omitempty" bson:"typecoin,omitempty"`
	TypeBusiness   string             `json:"typeBusiness,omitempty" bson:"typebusiness,omitempty"`
	Observations   string             `json:"observations,omitempty" bson:"observations,omitempty"`
	Status         string             `json:"status,omitempty" bson:"status,omitempty" `
	DeliveryPeriod string             `json:"deliveryPeriod,omitempty" bson:"deliveryperiod,omitempty"`
	Quality        string             `json:"quality,omitempty" bson:"quality,omitempty" `
	WayPay         string             `json:"wayPay,omitempty" bson:"waypay,omitempty"`
	IsPort         string             `json:"isPort,omitempty" bson:"isport,omitempty"`
	Expiration     string             `json:"expiration,omitempty" bson:"expiration,omitempty"`
}

type Order struct {

	//Produc Product
	ProductName    string             `json:"productName,omitempty" bson:"productname,omitempty"`
	Price          string             `json:"price,omitempty" bson:"price,omitempty"`
	TypeCoin       string             `json:"typeCoin,omitempty" bson:"typecoin,omitempty"`
	StatusOrder    string             `json:"statusOrder,omitempty" bson:"statusorder,omitempty"`
	DeliveryPeriod string             `json:"deliveryPeriod,omitempty" bson:"deliveryperiod,omitempty"`
	SellerName     string             `json:"sellerName,omitempty" bson:"sellername,omitempty"`
	Tons           string             `json:"tons,omitempty" bson:"tons,omitempty"`
	Id             primitive.ObjectID `json:"_id,omitempty"`
}

type Buscar struct {
	Email  string `json:"email,omitempty" bson:"email,omitempty"`
	Status string `json:"status,omitempty" bson:"status,omitempty"`
}

type BuscarProd struct {
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	Id   string `json:"id,omitempty" bson:"id,omitempty"`
}
