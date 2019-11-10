// Package dishes is a dish for packages
// -------------------------------------
// .../restauranteapi/dishes/dishes.go
// -------------------------------------
package dishes

import (
	helper "fjapidishes/helper"
	"fmt"
	"log"

	dishes "fjapidishes/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Dishadd is for export
func Dishadd(dishInsert dishes.Dish) helper.Resultado {

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
func Find(dishFind string) (dishes.Dish, string) {

	database := helper.GetDBParmFromCache("CollectionDishes")

	dishName := dishFind
	dishnull := dishes.Dish{}

	session, err := mgo.Dial(database.Location)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(database.Database).C(database.Collection)

	result := []dishes.Dish{}
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
func Getall() []dishes.Dish {

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

	var results []dishes.Dish

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
func GetDishesByActivity(activity string) []dishes.Dish {

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

	var results []dishes.Dish

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
func GetAvailable() []dishes.Dish {

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

	var results []dishes.Dish

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
func Dishupdate(dishUpdate dishes.Dish) helper.Resultado {

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
func Dishdelete(dishDelete dishes.Dish) helper.Resultado {

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
