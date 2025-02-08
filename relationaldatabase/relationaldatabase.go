package relationaldatabase

import (
	"database/sql"
	"errors"
	"golang-project/golang-project/constants"
	"golang-project/golang-project/models"
	"log"
)

type DbWrapper struct {
	db *sql.DB
}

func Get(personId int) (models.UserData, error) {
	var data models.UserData
	db := getDb()
	err := db.QueryRow(`
  SELECT p.name, ph.number, a.city, a.state, a.street1, a.street2, a.zip_code 
  FROM person p
  INNER JOIN phone ph ON p.id = ph.person_id
  INNER JOIN address_join aj ON p.id = aj.person_id
  INNER JOIN address a ON aj.address_id = a.id
  WHERE p.id = ? `, personId).Scan(
		&data.Name,
		&data.PhoneNumber,
		&data.City,
		&data.State,
		&data.Street1,
		&data.Street2,
		&data.ZipCode,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.UserData{}, constants.ErrUnableToFetchPerson
		}
		log.Println(err)
		return models.UserData{}, constants.ErrFailedToFetchPerson
	}
	return data, nil
}

func getDb() *sql.DB {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/cetec")
	if err != nil {
		log.Fatal(err)
	}

	// Check if the database connection is working
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
