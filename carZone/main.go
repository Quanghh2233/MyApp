package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Quanghh2233/MyApp/driver"
	carHandler "github.com/Quanghh2233/MyApp/handler/car"
	engineHandler "github.com/Quanghh2233/MyApp/handler/engine"
	carService "github.com/Quanghh2233/MyApp/service/car"
	engineService "github.com/Quanghh2233/MyApp/service/engine"
	carStore "github.com/Quanghh2233/MyApp/store/car"
	engineStore "github.com/Quanghh2233/MyApp/store/engine"
	"github.com/gorilla/mux"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	driver.InitDB()
	defer driver.CloseDB()

	db := driver.GetDB()
	carStore := carStore.New(db)
	carService := carService.NewCarService(carStore)

	engineStore := engineStore.New(db)
	engineService := engineService.NewEngineService(engineStore)

	carHandler := carHandler.NewCarHandler(carService)
	engineHandler := engineHandler.NewEngineHandler(engineService)

	router := mux.NewRouter()

	schemaFile := "store/schema.sql"
	if err := executeSchemaFile(db, schemaFile); err != nil {
		log.Fatal("Error while executing the schema file", err)
	}

	router.HandleFunc("/cars/{id}", carHandler.GetCarByID).Methods("GET")
	router.HandleFunc("/cars", carHandler.GetCarByBrand).Methods("GET")
	router.HandleFunc("/cars", carHandler.CreateCar).Methods("POST")
	router.HandleFunc("/cars/{id}", carHandler.UpdateCar).Methods("PUT")
	router.HandleFunc("/cars/{id}", carHandler.DeleteCar).Methods("DELETE")

	router.HandleFunc("/engine/{id}", engineHandler.GetEngineByID).Methods("GET")
	router.HandleFunc("/engine", engineHandler.CreateEngine).Methods("POST")
	router.HandleFunc("/engine/{id}", engineHandler.UpdateEngine).Methods("PUT")
	router.HandleFunc("/engine/{id}", engineHandler.DeleteEngine).Methods("DELETE")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := fmt.Sprintf(":%s", port)
	log.Printf("Server Listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}

func executeSchemaFile(db *sql.DB, fileName string) error {
	sqlFile, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	_, err = db.Exec(string(sqlFile))
	if err != nil {
		return err
	}
	return nil
}
