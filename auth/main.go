package main 

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type RegisterRequest struct {
	Phone int `json:"phone"`
	Password string `json:"password"`
}


func main() {

	http.HandleFunc("/register", registerHandler)

	fmt.Println("server started on : 8000")

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		fmt.Println(err)
	}
}



func registerHandler(w http.ResponseWriter, r *http.Request) {
	var request RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		fmt.Fprintln(w, "invalid JSON")
		return
	}

	fmt.Fprintln(w, "Phone", request.Phone)
	fmt.Fprintln(w, "Password", request.Password)
	
}