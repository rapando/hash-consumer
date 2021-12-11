package utils

import (
	"github.com/streadway/amqp"
)

func QConnect(qURI string) (conn *amqp.Connection, err error) {

	conn, err = amqp.Dial(qURI)
	if err != nil {
		return
	}
	return
}
