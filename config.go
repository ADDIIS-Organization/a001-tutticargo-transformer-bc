package main

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func initDB() {
	var err error

	// Leer variables de entorno
	dbUser := goDotEnvVariable("DB_USER")
	dbPassword := goDotEnvVariable("DB_PASSWORD")
	dbName := goDotEnvVariable("DB_NAME")
	dbHost := goDotEnvVariable("DB_HOST")
	dbPort := goDotEnvVariable("DB_PORT")

	log.Println("DB_USER:", dbUser)
	log.Println("DB_PASSWORD:", dbPassword)
	log.Println("DB_NAME:", dbName)
	log.Println("DB_HOST:", dbHost)
	log.Println("DB_PORT:", dbPort)

	// Verificar que todas las variables de entorno estén configuradas
	if dbUser == "" || dbPassword == "" || dbName == "" || dbHost == "" || dbPort == "" {
		log.Fatalln("One or more required environment variables are missing")
	}

	connStr := "user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " host=" + dbHost + " port=" + dbPort + " sslmode=disable options='-c extra_float_digits=0'"
	db, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	db.SetMaxOpenConns(400) // Aumentar el número de conexiones abiertas
	db.SetMaxIdleConns(50)  // Aumentar el número de conexiones en espera
}
