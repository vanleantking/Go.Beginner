package Mongo

import (
	"gopkg.in/mgo.v2/bson"
)

type MAddress struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	Street      string        `json:"street" bson:"street"`
	Ward        string        `json:"ward" bson:"ward"`
	City        string        `json:"city" bson:"city"`
	Country     string        `json:"country" bson:"country"`
	FullAddress string        `json:"fulladdress" bson:"fulladdress"`
}

// Receiver base type
func (addr *MAddress) GetFullAddress() string {
	return addr.FullAddress
}
