package config

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/lib/pq"
	"log"
	"github.com/joho/godotenv"
)

func CreateConnection() *sql.DB{
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error when load .env file")
	}
	db,err:=sql.Open("postgres",os.Getenv("POSTGRES_URL"))
	
	if err!=nil{
		panic(err)
	}

	fmt.Println("Connected")
	

	return db

}