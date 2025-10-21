# Saludos en go

## Instalación

```bash
go get -u github.com/litoos11/greetings
```

## Uso

```go
package main

import (
	"fmt"
	"log"

	"github.com/litoos11/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
```
