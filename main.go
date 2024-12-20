package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/juexuanl/rssagg/internal/database"

	_ "github.com/lib/pq"
)

//create apiConfig struct to store database connection pool
type apiConfig struct {
	DB *database.Queries
}

func main() {
	fmt.Println("Hello, World!")

	godotenv.Load(".env")
	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT is not found in the environment")		
	}

	//load db_url from environment variable and check if it is empty
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	//create database connection pool
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Fail to connect to database:", err)
	}

	//create new apiConfig
	apiCfg := apiConfig{
		DB: database.New(conn),
	}


	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	//create new router v1Router
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", errorHandler)
	v1Router.Post("/users", apiCfg.handlerCreateUser)

	//mount v1Router to router
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server started on port %s", portString)
	//listen to server
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

