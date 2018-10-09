package Mongo

import (
	"database/sql"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

type MUser struct {
	Id          bson.ObjectId                         `json:"id" bson:"_id"`
	Username    string                                `json:"username" bson:"username"`
	Password    string                                `json:"password" bson:"password"`
	Dob         sql.NullString                        `json:"dob" bson:"dob"`
	*MAddresses `json:"maddresses" bson:"maddresses"` //composition declare
}

type MAddresses struct {
	Addr []MAddress
}

func (u *MUser) GetAddresses() int {
	for index, ad := range u.MAddresses.Addr {
		fmt.Println("address ", index, ad.FullAddress)
	}
	return len(u.MAddresses.Addr)
}
