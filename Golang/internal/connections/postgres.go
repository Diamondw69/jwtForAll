package connections

import (
	"database/sql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var Path string = "./internal/connections/env/.env"

func EnvLoader(pathToEnv string) {
	err := godotenv.Load(pathToEnv)
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
}

func DbConnection() *sql.DB {

	EnvLoader(Path)

	psqlInfo := os.Getenv("POSTGRESQL")

	db, er := sql.Open("postgres", psqlInfo)
	if er != nil {
		log.Fatalf("postgres doesnt work : %s", er)
	}

	return db
}
