/*
 * user-service
 *
 * Use this file to generate server stub
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	sw "github.com/eriksterneberg/go-user-service/src/go"
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")

	router := sw.NewRouter()

	//http.HandleFunc("/health", sw.Health)
	log.Fatal(http.ListenAndServe(":8080", router))
}