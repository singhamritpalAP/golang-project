package models

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Phone struct {
	ID       int    `json:"id"`
	Number   string `json:"number"`
	PersonID int    `json:"person_id"`
}

type Address struct {
	ID      int    `json:"id"`
	City    string `json:"city"`
	State   string `json:"state"`
	Street1 string `json:"street1"`
	Street2 string `json:"street2,omitempty"`
	ZipCode string `json:"zip_code"`
}

type UserData struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	City        string `json:"city"`
	State       string `json:"state"`
	Street1     string `json:"street1"`
	Street2     string `json:"street2,omitempty"` // Use omitempty to omit if empty
	ZipCode     string `json:"zip_code"`
}
