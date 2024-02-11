// TODO: Add documentation
package main

import (
	"log"
	"net/http"

	"githib.com/mohitudupa/rosella/utils"
)

func main() {
	ac, err := utils.NewApplicationConfig("config.json")
	if err != nil {
		log.Printf("WARN: could not load application configs. %s", err.Error())
	}

	repo := getRepository(ac)
	server := getServer(repo, ac.Server.ReadOnly)
	serverLocation := getServerLocation(ac)

	err = repo.Connect()
	if err != nil {
		log.Fatalf("ERROR: could not connect to repository. %s", err.Error())
	}
	defer repo.Close()

	log.Printf("INFO: starting HTTP server on %s", serverLocation)
	err = http.ListenAndServe(serverLocation, server)
	if err != nil {
		log.Fatalf("ERROR: could not start http server. %s", err.Error())
	}
}
