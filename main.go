package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello")

	addr := "localhost:8083"

	http.ListenAndServe(addr, nil)
}

func handleAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
