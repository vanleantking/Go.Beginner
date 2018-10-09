/** For insert MySQL
*/
package main

import (
	"fmt"

	"./models/Mongo"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	ADDRESS = "localhost"
)

func main() {
	sess, err := mgo.Dial(ADDRESS)
	if err != nil {
		panic(err.Error())
	}
	defer sess.Close()
	m_user := sess.DB("test").C("user")
	m_address := sess.DB("test").C("address")

	addr1 := Mongo.MAddress{Id: bson.NewObjectId(), Street: "No.44", Country: "Viet Nam", FullAddress: "Duong so 44, phuong Thao Dien, Quan 2, tp HCM"}
	addr2 := Mongo.MAddress{Id: bson.NewObjectId(), Street: "Luong Dinh Cua", City: "Ho Chi Minh", Country: "Viet Nam"}
	addr3 := Mongo.MAddress{Id: bson.NewObjectId(), Ward: "9", City: "Tuy Hoa", Country: "Viet Nam"}

	er := m_address.Insert(addr1, addr2, addr3)
	if er != nil {
		panic(er.Error())
	}
	// var addresses Addresses
	addresses := &Mongo.MAddresses{}
	var adds = []Mongo.MAddress{addr1, addr2, addr3}
	addresses.Addr = adds

	user := Mongo.MUser{Id: bson.NewObjectId(), Username: "Le Van", Password: "123456", MAddresses: addresses}
	user.MAddresses.Addr[2].FullAddress = "Phan Huy Ich, Phuong 9, Tp Tuy Hoa"
	fmt.Println("address: ", user.MAddresses.Addr[2].GetFullAddress())
	fmt.Println("This user has ...")
	ad := user.GetAddresses()
	fmt.Println(ad)
	u_er := m_user.Insert(user)
	if u_er != nil {
		panic(u_er.Error())
	}
	fmt.Println("Insert successful")
}
