package handlers

import (
	"database/sql"
	"fmt"
	"go-mysql/models"
	"log"
)

func GetContacts(db *sql.DB) {
	query := "SELECT id, name, email, phone FROM contact"

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	fmt.Println("Contacts:")
	fmt.Println("--------------------------------------")
	for rows.Next() {
		contact := models.Contact{}
		var vPhone sql.NullString
		err := rows.Scan(&contact.Id, &contact.Name, &contact.Email, &vPhone)

		if vPhone.Valid {
			contact.Phone = vPhone.String
		} else {
			contact.Phone = "N/A"
		}

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Email: %s, Phone: %s\n", contact.Id, contact.Name, contact.Email, contact.Phone)
		fmt.Println("--------------------------------------")
	}
}

func GetContactById(db *sql.DB, id int) {
	query := "SELECT id, name, email, phone FROM contact WHERE id = ?"

	row := db.QueryRow(query, id)

	contact := models.Contact{}
	var vPhone sql.NullString
	err := row.Scan(&contact.Id, &contact.Name, &contact.Email, &vPhone)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("No contact found with ID %d\n", id)
		}
	}

	if vPhone.Valid {
		contact.Phone = vPhone.String
	} else {
		contact.Phone = "N/A"
	}

	fmt.Println("Contact:")
	fmt.Println("--------------------------------------")
	fmt.Printf("ID: %d, Name: %s, Email: %s, Phone: %s\n", contact.Id, contact.Name, contact.Email, contact.Phone)
	fmt.Println("--------------------------------------")
}

func CreateContact(db *sql.DB, c models.Contact) {
	query := "INSERT INTO contact (name, email, phone) VALUES (?, ?, ?)"

	_, err := db.Exec(query, c.Name, c.Email, c.Phone)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Contact created successfully")
}

func DeleteContact(db *sql.DB, id int) {
	query := "DELETE FROM contact WHERE id = ?"

	_, err := db.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Contact deleted successfully")
}

func UpdateContact(db *sql.DB, c models.Contact) {
	query := "UPDATE contact SET name = ?, email = ?, phone = ? WHERE id = ?"

	_, err := db.Exec(query, c.Name, c.Email, c.Phone, c.Id)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Contact updated successfully")
}
