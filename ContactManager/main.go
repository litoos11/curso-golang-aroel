package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func setContact(contacts []Contact) error {
	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(contacts)
	if err != nil {
		return err
	}
	return nil
}

func getContacts(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")
	if err != nil {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	var contacts []Contact

	err := getContacts(&contacts)
	if err != nil {
		fmt.Println("No existing contacts found, starting fresh.")
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("=== GESTOR DE CONTACTOS ===\n",
			"1. Añadir contacto\n",
			"2. Listar contactos\n",
			"3. Salir\n",
			"Seleccione una opción: ")

		var option int
		_, err = fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Entrada inválida, por favor intente de nuevo.")
			return
		}

		switch option {
		case 1:
			var c Contact
			fmt.Print("Nombre: ")
			c.Name, _ = reader.ReadString('\n')
			fmt.Print("Email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Print("Teléfono: ")
			c.Phone, _ = reader.ReadString('\n')

			contacts = append(contacts, c)

			if err := setContact(contacts); err != nil {
				fmt.Println("Error al guardar el contacto:", err)
			}
		case 2:
			fmt.Println("=== LISTA DE CONTACTOS ===")
			for index, contac := range contacts {
				fmt.Printf("%d. Nomber: %s - Email: %s - Telefono: %s\n", index+1, contac.Name, contac.Email, contac.Phone)
			}
			fmt.Println("==========================")
		case 3:
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción inválida, por favor intente de nuevo.")
		}

	}
}
