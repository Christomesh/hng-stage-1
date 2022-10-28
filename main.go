package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type jsonResponse struct {
	SlackUsername string `json:"slackusername"`
	Backend       bool   `json:"backend"`
	Age           int    `json:"age"`
	Bio           string `json:"bio"`
}

func main() {
	port := "3000"
	mux := http.NewServeMux()
	mux.HandleFunc("/", giveResponse)
	fmt.Println("Server listening on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, mux))

	//.Println("Server listening on port: " + port)
}

func giveResponse(w http.ResponseWriter, r *http.Request) {
	myRes := &jsonResponse{
		SlackUsername: "Meshach",
		Backend:       true,
		Age:           22,
		Bio:           "A junior backend developer",
	}
	resp, err := json.Marshal(myRes)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintln(w, string(resp))

}
