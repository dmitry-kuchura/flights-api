package api

import (
	"net/http"

	u "flights-api/utils"
)

func GetFlights(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["airport"]

	collection := u.GetConnection()

	key := ""

	if !ok || len(keys[0]) < 1 {
		key = ""
	} else {
		key = keys[0]
	}

	data, _ := u.FindMany(collection, key)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

func GetDepartedFlights(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["airport"]

	collection := u.GetConnection()

	key := ""

	if !ok || len(keys[0]) < 1 {
		key = ""
	} else {
		key = keys[0]
	}

	data, _ := u.FindDeparted(collection, key)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

func GetArrivalFlights(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["airport"]

	collection := u.GetConnection()

	key := ""

	if !ok || len(keys[0]) < 1 {
		key = ""
	} else {
		key = keys[0]
	}

	data, _ := u.FindArriving(collection, key)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
