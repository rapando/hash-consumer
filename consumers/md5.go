package consumers

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"log"
	"sync"

	"github.com/streadway/amqp"
)

func MD5Consumer(wg *sync.WaitGroup, q *amqp.Channel, db *sql.DB) {
	defer wg.Done()
	log.Printf("consuming from md5 queue")
	_ = q.Qos(10, 0, false)
	msgChan, _ := q.Consume("md5", "md5-consumer", true, false, false, false, nil)

	stopChan := make(chan bool)
	go func() {
		for msg := range msgChan {
			password := string(msg.Body)
			log.Printf("[md5] consumed %v", password)
			hash := md5.Sum([]byte(password))
			hashed := hex.EncodeToString(hash[:])
			save("md5", password, hashed, db)
		}
	}()
	<-stopChan
}
