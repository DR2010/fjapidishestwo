// Package models is a dish for packages
// -------------------------------------
// .../restauranteapi/models/dishes.go
// -------------------------------------
package models

import (
	"database/sql"

	"gopkg.in/mgo.v2/bson"

	"fmt"
	"log"

	"gopkg.in/mgo.v2"

	helper "dmapicomplex/helper"
)

type DishModel struct {
	Db *sql.DB
}

// Dish is to be exported
type Dish struct {
	SystemID         bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	Name             string        // name of the dish - this is the KEY, must be unique
	Type             string        // type of dish, includes drinks and deserts
	Price            string        // preco do prato multiplicar por 100 e nao ter digits
	GlutenFree       string        // Gluten free dishes
	DairyFree        string        // Dairy Free dishes
	Vegetarian       string        // Vegeterian dishes
	InitialAvailable string        // Number of items initially available
	CurrentAvailable string        // Currently available
	ImageName        string        // Image Name
	Description      string        // Description of the plate
	Descricao        string        // Descricao do prato
	ActivityType     string        // Descricao do activity
	ImageBase64      string        // Descricao do activity

}

// Dishadd is for export
func Dishadd(dishInsert Dish) helper.Resultado {

	database := helper.GetDBParmFromCache("CollectionDishes")

	session, err := mgo.Dial(database.Location)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	collection := session.DB(database.Database).C(database.Collection)

	err = collection.Insert(dishInsert)

	if err != nil {
		log.Fatal(err)
	}

	var res helper.Resultado
	res.ErrorCode = "0001"
	res.ErrorDescription = "Dish added"
	res.IsSuccessful = "Y"

	return res
}

// Find is to find stuff
func DishFind(dishFind string) (Dish, string) {

	database := helper.GetDBParmFromCache("CollectionDishes")

	dishName := dishFind
	dishnull := Dish{}

	session, err := mgo.Dial(database.Location)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(database.Database).C(database.Collection)

	result := []Dish{}
	err1 := c.Find(bson.M{"name": dishName}).All(&result)
	if err1 != nil {
		log.Fatal(err1)
	}

	var numrecsel = len(result)

	if numrecsel <= 0 {
		return dishnull, "404 Not found"
	}

	return result[0], "200 OK"
}

// Getall works
func DishGetall() []Dish {

	// Get database variables
	database := helper.GetDBParmFromCache("CollectionDishes")

	session, err := mgo.Dial(database.Location)

	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(database.Database).C(database.Collection)

	var results []Dish

	err = c.Find(nil).All(&results)
	if err != nil {
		// TODO: Do something about the error
	} else {
		return results
	}

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// GetDishesByActivity works
func DishGetByActivity(activity string) []Dish {

	// Get database variables
	database := helper.GetDBParmFromCache("CollectionDishes")

	session, err := mgo.Dial(database.Location)

	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(database.Database).C(database.Collection)

	var results []Dish

	err = c.Find(bson.M{"activitytype": activity}).All(&results)
	if err != nil {
		// TODO: Do something about the error
	} else {
		return results
	}

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// GetAvailable works
func DishGetAvailable() []Dish {

	database := helper.GetDBParmFromCache("CollectionDishes")

	fmt.Println("database.Location")
	fmt.Println(database.Location)

	session, err := mgo.Dial(database.Location)

	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(database.Database).C(database.Collection)

	var results []Dish

	err = c.Find(bson.M{"currentavailable": bson.M{"$ne": "0"}}).All(&results)

	if err != nil {
		// TODO: Do something about the error
		log.Println("GetAvailable Error http.NewRequest(GET, url, nil): ", err)

		var res helper.Resultado
		res.ErrorCode = "0021"
		res.ErrorDescription = "Something Bad Happened"
		res.IsSuccessful = "N"

	} else {
		return results
	}

	return nil
}

// Dishupdate is
func Dishupdate(dishUpdate Dish) helper.Resultado {

	database := new(helper.DatabaseX)
	database.Collection = helper.Getvaluefromcache("CollectionDishes")
	database.Database = helper.Getvaluefromcache("API.MongoDB.Database")
	database.Location = helper.Getvaluefromcache("API.MongoDB.Location")

	session, err := mgo.Dial(database.Location)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	collection := session.DB(database.Database).C(database.Collection)

	err = collection.Update(bson.M{"name": dishUpdate.Name}, dishUpdate)

	var res helper.Resultado

	if err != nil {
		// log.Fatal(err)

		log.Println("Dishupdate Error http.NewRequest(GET, url, nil): ", err)

		res.ErrorCode = "0021"
		res.ErrorDescription = "Something Bad Happened"
		res.IsSuccessful = "N"

		return res

	}

	res.ErrorCode = "0001"
	res.ErrorDescription = "Something Happened"
	res.IsSuccessful = "Y"

	return res
}

// Dishdelete is
func Dishdelete(dishDelete Dish) helper.Resultado {

	database := helper.GetDBParmFromCache("CollectionDishes")

	session, err := mgo.Dial(database.Location)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	collection := session.DB(database.Database).C(database.Collection)

	err = collection.Remove(bson.M{"name": dishDelete.Name})

	var res helper.Resultado

	if err != nil {
		log.Fatal(err)

		log.Println("Dishdelete Error http.NewRequest(GET, url, nil): ", err)

		res.ErrorCode = "0021"
		res.ErrorDescription = "Something Bad Happened"
		res.IsSuccessful = "N"

		return res
	}

	res.ErrorCode = "0001"
	res.ErrorDescription = "Dish deleted successfully"
	res.IsSuccessful = "Y"

	return res
}
