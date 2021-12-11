package main

import (
	"database/sql"
	"log"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/rapando/hash-consumer/consumers"
	"github.com/rapando/hash-consumer/utils"
	"github.com/streadway/amqp"
)

var (
	q  *amqp.Connection
	db *sql.DB
)

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Printf("unable to read dotenv because %v", err)
		os.Exit(3)
	}

	q, err = utils.QConnect(os.Getenv("Q_URI"))
	if err != nil {
		log.Printf("unable to connect to rabbitmq because %v", err)
		os.Exit(3)
	}

	db, err = utils.DbConnect(os.Getenv("DB_URI"))
	if err != nil {
		log.Printf("unable to connect to db because %v", err)
		return
	}

	md5Channel, _ := q.Channel()

	var wg sync.WaitGroup
	wg.Add(1)
	go consumers.MD5Consumer(&wg, md5Channel, db)
	wg.Wait()
}
