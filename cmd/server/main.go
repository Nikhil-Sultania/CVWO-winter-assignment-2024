package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/router"
)

func main() {
	r := router.Setup()
	fmt.Print("Listening on port 8080 at http://localhost:8080!")
	database.Init()
	defer database.Db.Close()
	r.Run("localhost:8080")
	log.Fatalln(http.ListenAndServe(":8080", r))
}
