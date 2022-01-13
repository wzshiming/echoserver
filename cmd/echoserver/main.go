package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/wzshiming/echoserver"
)

func main() {
	port := ":8080"
	envs := []string{"HOSTNAME", "MESSAGE"}

	tmp := "SERVER VALUE:"
	for _, env := range envs {
		v := os.Getenv(env)
		tmp += fmt.Sprintf("\n%s: %s", env, v)
	}
	var handler http.Handler = echoserver.Handler{
		Message: tmp,
	}
	handler = handlers.LoggingHandler(os.Stderr, handler)

	log.Printf("Starting server on %s", port)
	err := http.ListenAndServe(port, handler)
	if err != nil {
		log.Fatal(err)
	}
}
