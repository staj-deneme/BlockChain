package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type User struct {
	UserName  string
	Password  string
	IsAdmin   bool
	CreatedAt time.Time
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/test", testHandler)

	http.ListenAndServe(":3000", mux)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	user := User{}
	/*
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			panic(err)
		}
	*/
	user.UserName = "ahmet"
	user.Password = "mehmet"
	user.IsAdmin = true
	user.CreatedAt = time.Now().Local()

	userJson, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(userJson)

}
