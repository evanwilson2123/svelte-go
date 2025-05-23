package handlers

import (
	"encoding/json"
	"fmt"
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
		return
	}

	users = append(users, user)
	for _, user := range users {
		fmt.Printf("\nusername: %s\nemail: %s\npassword: %s", user.Name, user.Email, user.Password)
	}
	response["ok"] = "true"
}