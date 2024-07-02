package main

import (
	"github.com/dubirajara/go-expert/fcutils/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	rabbitmq.Publish(ch, "hello World!", "amq.direct")

}
