package main 

import (
	"fmt"
	"net/http"
	"encoding/json"
	"unicode"
)

type RegisterRequest struct {
	Phone string `json:"phone"`
	Password string `json:"password"`
}


func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/register", registerHandler)

	fmt.Println("server started on : 8000")

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		fmt.Println(err)
	}
}


func rootHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, map[string]string{"message": "ok"})
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var request RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	err = validatePhone(request.Phone)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	err = validatePassword(request.Password)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}


	fmt.Fprintln(w, "successfully")

	
}

func validatePhone(phone string) error {
	if len(phone) != 11 {
		return fmt.Errorf("phone number must be 11 digits")
	}
	return nil
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}

	hasDigit := false

	for _, r := range password {
		if unicode.IsDigit(r){
			hasDigit = true
		}
	}

	if hasDigit == false {
		return fmt.Errorf("There must be at least one digit.")
	}

	return nil
}
