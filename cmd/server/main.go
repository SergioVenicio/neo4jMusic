package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/SergioVenicio/neo4jMusic/handlers"
	"github.com/SergioVenicio/neo4jMusic/repositories"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)


func main() {
	godotenv.Load()

	dbUri := os.Getenv("NEO4J_HOST")
	user := os.Getenv("NEO4J_USER")
	pwd := os.Getenv("NEO4J_PWD")

	driver, err := neo4j.NewDriverWithContext(dbUri, neo4j.BasicAuth(user, pwd, ""))
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	defer driver.Close(ctx)

	albumRepo := repositories.NewAlbumRepository(driver)
	albumHandler := handlers.NewAlbumHandler(albumRepo)
	
	r := chi.NewRouter()
	r.Route("/album", func(r chi.Router)  {
		r.Get("/{name}", albumHandler.FindByName)
		r.Get("/", albumHandler.FindAll)
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}