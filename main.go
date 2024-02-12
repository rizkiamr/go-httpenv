package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

// Define a home handler function which writes a byte slice containing
// all current environment variables as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	// Get all environment variables
	envVariables := make(map[string]string)
	for _, env := range os.Environ() {
		keyValue := strings.SplitN(env, "=", 2)
		envVariables[keyValue[0]] = keyValue[1]
	}

	// Encode the environment variables as JSON
	jsonData, err := json.Marshal(envVariables)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(jsonData)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the response
	w.Write([]byte("{\"name\":\"httpenv\",\"healthy\":true}"))
}

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/env", home)
	mux.HandleFunc("/healthz", healthCheck)

	// Use the http.ListenAndServe() function to start a new web server.
	// We pass in two parameters: the TCP network address to listen on
	// (in this case ":8080") and the servemux we just created.
	// If http.ListenAndServe() returns an error, we use the log.Fatal() function
	// to log the error message and exit.
	// Note that any error returned by http.ListenAndServe() is always non-nil.
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
