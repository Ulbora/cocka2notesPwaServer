package main

import (
	"fmt"
	"net/http"
	"os"

	b64 "encoding/base64"

	"github.com/gorilla/mux"
)

// go mod init github.com/Ulbora/cocka2notesPwaServer

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/rs/config/get", config).Methods("GET")

	//site pages
	//router.HandleFunc("/", index).Methods("GET")
	fmt.Println("Running on port 8080")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	http.ListenAndServe(":8080", router)
}

func config(w http.ResponseWriter, r *http.Request) {
	var data string
	if os.Getenv("COCKA2NOTE_CONFIG") != "" {
		data = os.Getenv("COCKA2NOTE_CONFIG")
	} else {
		data = "ewogICAidXJsIjoiaHR0cDovL2xvY2FsaG9zdDozMDAwIiwKICAgImFwaUtleSI6IkdERzY1MUdGRDY2RkQxNjE1MXNzczY1MWY2NTFmZjY1NTU1ZGRmaGprbHl5NSIKfQ=="
	}

	rtn, _ := b64.StdEncoding.DecodeString(data)
	fmt.Println(rtn)

	fmt.Fprint(w, string(rtn))

}
