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

	user := Mongo.MUser{Id: bson.NewObjectId(), Username: "Le Van", Password: "123456", MAddresses: addresses, Dob: ""}
	user.MAddresses.Addr[2].FullAddress = "Phan Huy Ich, Phuong 9, Tp Tuy Hoa"
	u_er := m_user.Insert(user)
	if u_er != nil {
		panic(u_er.Error())
	}

	// Declare by composition literal
	user2 := &Mongo.MUser{
		Id:         bson.NewObjectId(),
		Username:   "Hong Van",
		Password:   "lvhv",
		MAddresses: &Mongo.MAddresses{[]Mongo.MAddress{Mongo.MAddress{Id: bson.NewObjectId(), Street: "93 Phan Huy Ich", City: "Quy Nhon", Country: "Viet Nam"}}}}
	u_er2 := m_user.Insert(user2)
	if u_er2 != nil {
		panic(u_er2.Error())
	}

	fmt.Println("Insert successful")
	u_iters := m_user.Find(bson.M{"username": "Le Van"}).Iter()
	users := []Mongo.MUser{}
	u := Mongo.MUser{}
	for u_iters.Next(&u) {
		users = append(users, u)
		fmt.Println("User name: ", u.Username)
	}
	for _, uiter := range users {
		uiter.Username = "updated"
		uiter.Password = "updated_pwd"

		uiter.MAddresses.Addr[1].FullAddress = "Chung cu Binh Minh, phuong Binh An, Quan 2"
		a_er := m_address.Update(bson.M{"_id": uiter.MAddresses.Addr[1].Id}, uiter.MAddresses.Addr[1])
		if a_er != nil {
			panic(a_er.Error())
		}
		u_er := m_user.Update(bson.M{"_id": uiter.Id}, uiter)
		if u_er != nil {
			panic(u_er.Error())
		}
	}

}
