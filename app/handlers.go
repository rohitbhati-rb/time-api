package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func getTime(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string, 0)
	tz := r.URL.Query().Get("tz")
	timezones := strings.Split(tz, ",")

	if len(timezones) <= 1 {
		location, error := time.LoadLocation(tz)
		if error != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(fmt.Sprintf("Invalid Timezone %s", tz)))
		} else {
			response["current_time"] = time.Now().In(location).String()
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		}
	} else {
		for _, timezone := range timezones {
			location, error := time.LoadLocation(timezone)
			if error != nil {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(fmt.Sprintf("Invalid Timezone %s in input", timezone)))
				return
			}
			response[timezone] = time.Now().In(location).String()
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
