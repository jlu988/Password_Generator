package main

import (
	"PasswordGenerator/domain"
	"fmt"
	"net/http"
)

func main() {
	router := domain.RouterInitialize()
	fmt.Println("Listen and serve")
	http.ListenAndServe(":8080", router)
}
