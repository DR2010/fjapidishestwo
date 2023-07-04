// Package complex is a dish for packages
// -------------------------------------
// .../restauranteapi/models/dishes.go
// -------------------------------------
package models

import (
	"context"
	"database/sql"
	"dmapicomplex/helper"
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ComplexModel struct {
	Db *sql.DB
}

// Dish is to be exported
type Complex struct {
	ID            int
	Name          string // name of the dish - this is the KEY, must be unique
	Streetnumber  int    // type of dish, includes drinks and deserts
	Streetname    string // preco do prato multiplicar por 100 e nao ter digits
	UPNumber      string // Gluten free dishes
	Recordversion int
	CreatedBy     string
	UpdatedBy     string
	UpdatedOn     time.Time
	CreatedOn     time.Time
}

// ComplexGetall - lists all complexes
func (complexModel ComplexModel) ComplexGetall() ([]Complex, error) {

	rows, err := complexModel.Db.Query("select * from younit.complex")

	if err != nil {
		return nil, err
	}

	complexes := []Complex{}
	for rows.Next() {
		var ID int
		var Name string
		var Streetnumber int
		var Streetname string
		var UPNumber string
		var Recordversion int
		var CreatedBy string
		var UpdatedBy string
		var CreatedOn time.Time
		var UpdatedOn time.Time

		err2 := rows.Scan(&ID, &Name, &Streetnumber, &Streetname, &UPNumber, &Recordversion, &CreatedBy, &UpdatedBy, &CreatedOn, &UpdatedOn)
		if err2 != nil {
			return nil, err2
		} else {
			complex := Complex{ID, Name, Streetnumber, Streetname, UPNumber, Recordversion, CreatedBy, UpdatedBy, CreatedOn, UpdatedOn}
			complexes = append(complexes, complex)
		}
	}
	return complexes, nil
}

// Complexadd is for export
func (complexModel ComplexModel) ComplexAdd(complexInsert Complex) (int64, error) {

	query := "insert into younit.complex (name, streetnumber, streetname, upnumber, recordversion, createdby, updatedby) values (?, ?, ?, ?, ?, ?, ?);"
	insertResult, err := complexModel.Db.ExecContext(context.Background(), query, complexInsert.Name, complexInsert.Streetname, complexInsert.UPNumber, 1, complexInsert.CreatedBy, complexInsert.UpdatedBy)

	newid, err := insertResult.LastInsertId()

	if err != nil {
		fmt.Printf("impossible insert complex: %s", err)
		return 0, err
	}

	id, err := insertResult.LastInsertId()
	if err != nil {
		fmt.Printf("impossible to retrieve last inserted id: %s", err)
		return 0, err
	}
	fmt.Printf("inserted id: %d", id)

	fmt.Println("Yay, values added!")

	return newid, nil
}

// Find is to find stuff
func ComplexFind(complexFind string) (Complex, string) {

	database := helper.GetDBParmFromCache("Collectioncomplexes")

	complexName := complexFind
	complexnull := Complex{}

	session, err := mgo.Dial(database.Location)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(database.Database).C(database.Collection)

	result := []Complex{}
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

// GetcomplexesByActivity works
func ComplexGetByActivity(activity string) []Complex {

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

	var results []Complex

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
func ComplexGetAvailable() []Complex {

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

	var results []Complex

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

// ComplexUpdate is
func ComplexUpdate(dishUpdate Complex) helper.Resultado {

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
func Complexdelete(dishDelete Dish) helper.Resultado {

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
