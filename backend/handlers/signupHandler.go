package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"stoic/types"
	"stoic/utils"
)

var users []types.User

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	utils.SetHeaders(w, r)
	response := make(map[string]string)
	var user types.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response["error"] = "Internal Server error"
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Internal Server error", 500)
			return
		}
		return
	}
	users = append(users, user)
	for _, user := range users {
		fmt.Printf("\nusername: %s\nemail: %s\npassword: %s", user.Name, user.Email, user.Password)
	}
	response["ok"] = "true"
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal Server error", 500)
		return
	}
}


func UserHandler(w http.ResponseWriter, r *http.Request) {
	utils.SetHeaders(w, r)

	w.WriteHeader(http.StatusOK)
	lastIndex := len(users) - 1

	response := make(map[string]string)
	response["name"] = users[lastIndex].Name

	log.Println(response)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}