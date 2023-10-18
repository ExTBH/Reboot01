package webapp

import (
	"fmt"
	"log"
	"net/http"
)

// ListenOn starts the server and listens on the specified port.
// Parameters:
//
//	port (uint16): The port number on which to listen.
//
// Returns:
//
//	error: An error if there was a problem starting the server.
func ListenOn(port uint16) error {
	log.Printf("Listening on http://localhost:%d\n", port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/band/", bandHanlder)
	mux.HandleFunc("/styles/", stylesHandler)
	mux.HandleFunc("/js/", jsHandler)
	mux.HandleFunc("/search", searchHandler)
	mux.HandleFunc("/filter", filterHandler)
	addr := fmt.Sprintf(":%d", port)
	return http.ListenAndServe(addr, mux)
}
