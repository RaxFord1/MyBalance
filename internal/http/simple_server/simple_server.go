package simple_server

import (
	"io"
	"log"
	"net/http"
	"os"
)

func SimpleServer() {
	port := os.Getenv("PORT")
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}
	http.HandleFunc("/", helloHandler)
	log.Println("Listing for" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
