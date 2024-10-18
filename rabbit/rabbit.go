package rabbit

import "log"

type Rabbit struct {
}

func NewRabbit() *Rabbit {
	return &Rabbit{}
}

func (r *Rabbit) PublishMessage(message string) {
	log.Printf("Publishing to the queue; message: %s\n", message)
}
