package handlers

import (
	"encoding/json"
	"net/http"
	"stoic/types"
	"stoic/utils"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	// set the required headers
	utils.SetHeaders(w, r)

	response := types.DailyQuote{
		Author: "Marcus Aurelius",
		Quote: "You have power over your mind - not outside events. Realize this, and you will find strength",
		DateWritten: "175 AD",
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", 500)
		return
	}
}