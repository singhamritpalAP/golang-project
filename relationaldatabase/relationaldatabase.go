package relationaldatabase

import (
	"database/sql"
	"errors"
	"golang-project/golang-project/constants"
	"golang-project/golang-project/models"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DbWrapper struct {
	db *sql.DB
}

// getDb for getting db todo db connection must be done while starting server
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

// Get returns user details for received user id
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

// CreateUser stores new user data in respective tables
func CreateUser(data models.UserData) error {
	db := getDb()
	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		log.Println("failed to start transaction")
		return err
	}
	// tx rollback in case of failure
	defer func() {
		if err != nil {
			err = tx.Rollback()
			if err != nil {
				log.Println("Failed to rollback transaction:", err)
			}
		}
	}()

	// 1. Insert into address table
	result, err := tx.Exec("INSERT INTO address (city, state, street1, street2, zip_code) VALUES (?, ?, ?, ?, ?)",
		data.City, data.State, data.Street1, data.Street2, data.ZipCode)
	if err != nil {
		log.Println("Failed to create address")
		return err
	}
	addressID, err := result.LastInsertId()
	if err != nil {
		log.Println("Failed to get last insert ID for address")
		return err
	}

	// 2. Insert into person table
	result, err = tx.Exec("INSERT INTO person (name) VALUES (?)", data.Name) // age can be updated later using update
	if err != nil {
		log.Println("Failed to create person")
		return err
	}

	personID, err := result.LastInsertId()
	if err != nil {
		log.Println("Failed to get last insert ID for person")
		return err
	}

	// 3. Insert into phone table
	_, err = tx.Exec("INSERT INTO phone (person_id, number) VALUES (?, ?)", personID, data.PhoneNumber)
	if err != nil {
		log.Println("Failed to create phone")
		return err
	}

	// 4. Insert into address_join table
	_, err = tx.Exec("INSERT INTO address_join (person_id, address_id) VALUES (?, ?)", personID, addressID)
	if err != nil {
		log.Println("Failed to create address_join")
		return err
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Println("Failed to commit transaction")
		return err
	}
	return nil
}
