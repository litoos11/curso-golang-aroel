package main

import (
	"bufio"
	"fmt"
	"go-mysql/database"
	"go-mysql/handlers"
	"go-mysql/models"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// handlers.GetContacts(db)
	// handlers.GetContactById(db, 1)

	// contact := models.Contact{
	// 	Name:  "John Doe",
	// 	Email: "jdow@x.com",
	// 	Phone: "123-456-7890",
	// }
	// handlers.CreateContact(db, contact)
	// uContact := models.Contact{
	// 	Id:    4,
	// 	Name:  "John Doe Edited",
	// 	Email: "jdow@x.com",
	// 	Phone: "123-456-7890",
	// }
	// handlers.CreateContact(db, contact)
	// handlers.UpdateContact(db, uContact)
	// handlers.DeleteContact(db, 4)

	for {
		fmt.Println("\nMenú:")
		fmt.Println("1. Listar Contactos")
		fmt.Println("2. Obtener contacto por Id")
		fmt.Println("3. Crear nuevo contacto")
		fmt.Println("4. Actualizar contacto")
		fmt.Println("5. Eliminar contacto")
		fmt.Println("6. Salir")
		fmt.Println("Seleccione una opción: ")

		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			handlers.GetContacts(db)
		case 2:
			fmt.Println("Ingrese el id del contacto que quiere buscar: ")
			var id int
			fmt.Scanln(&id)
			handlers.GetContactById(db, id)
		case 3:
			nContact := inputContactDetails(option)
			handlers.CreateContact(db, nContact)
		case 4:
			nContact := inputContactDetails(option)
			handlers.UpdateContact(db, nContact)
		case 5:
			fmt.Println("Ingrese el id del contacto que quiere eliminar: ")
			var id int
			fmt.Scanln(&id)
			handlers.DeleteContact(db, id)
		case 6:
			fmt.Println("Saliendo de la aplicacion...")
			return
			// os.Exit(0)
		default:
			fmt.Println("Opcion no valida...")
		}
	}
}

func inputContactDetails(option int) models.Contact {
	reader := bufio.NewReader(os.Stdin)

	var contact models.Contact
	if option == 4 {
		fmt.Print("Ingrese el id del contacto que quiere editar: ")
		var id int
		fmt.Scanln(&id)

		contact.Id = id
	}

	fmt.Print("Ingrese el nombre del contacto: ")
	name, _ := reader.ReadString('\n')
	contact.Name = strings.TrimSpace(name)

	fmt.Print("Ingrese el correo del contacto: ")
	email, _ := reader.ReadString('\n')
	contact.Email = strings.TrimSpace(email)

	fmt.Print("Ingrese el telefono del contacto: ")
	phone, _ := reader.ReadString('\n')
	contact.Phone = strings.TrimSpace(phone)

	return contact
}
