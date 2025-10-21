package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api. somewhereintheinternet. om/",
		"https://graph.microsoft.com",
	}
	chanel := make(chan string)
	for _, api := range apis {
		go CheckApi(api, chanel)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Println(<-chanel)
	}

	// time.Sleep(1 * time.Second) // Wait for goroutines to finish (for demo purposes)

	alapsed := time.Since(start)
	fmt.Printf("API checks completed in %s\n", alapsed)
}

func CheckApi(api string, chanel chan string) {
	// Placeholder for actual API checking logic
	if _, err := http.Get(api); err != nil {
		chanel <- fmt.Sprintf("API %s is down: %v\n", api, err)
		return
	}

	chanel <- fmt.Sprintf("API %s is up\n", api)

}
