package publisher

import (
	"encoding/json"
	"github.com/LucxLab/cim-service/internal/cdr"
	"github.com/LucxLab/cim-service/internal/rabbitmq"
	"github.com/LucxLab/cim-service/internal/rabbitmq/message"
	"github.com/rabbitmq/amqp091-go"
)

type rabbitmqCdr struct {
	publisher *rabbitmq.Publisher
}

func (r *rabbitmqCdr) PublishUploadCreated(upload *cdr.Upload) error {
	exchange := "cim"
	routingKey := "cdr_file_uploaded"

	cdrFileUploadedMessage := message.ToCdrFileUploaded(upload)
	jsonMessage, err := json.Marshal(cdrFileUploadedMessage)
	if err != nil {
		return err
	}

	err = r.publisher.Channel.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/json",
			Body:        jsonMessage,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func NewRabbitmqCdr(publisher *rabbitmq.Publisher) cdr.Publisher {
	return &rabbitmqCdr{
		publisher: publisher,
	}
}
