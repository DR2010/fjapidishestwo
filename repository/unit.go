// Package complexes is a complex for packages
// -------------------------------------
// .../restauranteapi/complexes/complexes.go
// -------------------------------------
package repository

import (
	"context"
	"database/sql"
	helper "dmapicomplex/helper"
	"fmt"
	"log"

	complexes "dmapicomplex/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Complexadd is for export
func Add(complexInsert complexes.Complex) helper.Resultado {

	// MySQL Connection
	MySQLConnString := helper.Getvaluefromcache("API.MySQL.ConnString")

	db, err := sql.Open("mysql", MySQLConnString)
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
	defer db.Close()

	// SQL Insert
	//

	// query := "INSERT INTO `teacher` (`create_time`, `firstname`, `lastname`) VALUES (NOW(), ?, ?)"
	query := "insert into younit.complex (name, streetnumber, streetname, upnumber, recordversion, createdby, updatedby) values (?, ?, ?, ?, ?, ?, ?);"
	insertResult, err := db.ExecContext(context.Background(), query, complexInsert.Name, complexInsert.Streetname, complexInsert.UPNumber, 1, complexInsert.CreatedBy, complexInsert.UpdatedBy)
	if err != nil {
		fmt.Printf("impossible insert complex: %s", err)
		log.Fatalf("impossible insert complex: %s", err)
	}
	id, err := insertResult.LastInsertId()
	if err != nil {
		fmt.Printf("impossible to retrieve last inserted id: %s", err)
		log.Fatalf("impossible to retrieve last inserted id: %s", err)
	}
	fmt.Printf("inserted id: %d", id)

	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}

	fmt.Println("Yay, values added!")

	var res helper.Resultado
	res.ErrorCode = "0001"
	res.ErrorDescription = "Complex added"
	res.IsSuccessful = "Y"

	return res
}

// Find is to find stuff
func Find(complexFind string) (complexes.Complex, string) {

	database := helper.GetDBParmFromCache("Collectioncomplexes")

	complexName := complexFind
	complexnull := complexes.Complex{}

	session, err := mgo.Dial(database.Location)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(database.Database).C(database.Collection)

	result := []complexes.Complex{}
	err1 := c.Find(bson.M{"name": complexName}).All(&result)
	if err1 != nil {
		log.Fatal(err1)
	}

	var numrecsel = len(result)

	if numrecsel <= 0 {
		return complexnull, "404 Not found"
	}

	return result[0], "200 OK"
}

// Getall works
func Getall() []complexes.Complex {

	// Get database variables
	database := helper.GetDBParmFromCache("Collectioncomplexes")

	session, err := mgo.Dial(database.Location)

	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(database.Database).C(database.Collection)

	var results []complexes.Complex

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

// GetcomplexesByActivity works
func GetcomplexesByActivity(activity string) []complexes.Complex {

	// Get database variables
	database := helper.GetDBParmFromCache("Collectioncomplexes")

	session, err := mgo.Dial(database.Location)

	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(database.Database).C(database.Collection)

	var results []complexes.Complex

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
func GetAvailable() []complexes.Dish {

	database := helper.GetDBParmFromCache("Collectioncomplexes")

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

	var results []complexes.Dish

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
func Dishupdate(dishUpdate complexes.Dish) helper.Resultado {

	database := new(helper.DatabaseX)
	database.Collection = helper.Getvaluefromcache("Collectioncomplexes")
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
func Dishdelete(dishDelete complexes.Dish) helper.Resultado {

	database := helper.GetDBParmFromCache("Collectioncomplexes")

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
