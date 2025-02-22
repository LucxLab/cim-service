package publishers

import (
	"context"
	"encoding/json"
	"github.com/LucxLab/cim-service/internal/publishers"
	"github.com/LucxLab/cim-service/internal/rabbitmq"
	"github.com/LucxLab/cim-service/internal/rabbitmq/message"
	"github.com/rabbitmq/amqp091-go"
)

const CdrFilesExchange = "cdr_files"
const CorrelationIdKey = "correlation_id"

type rabbitmqCdr struct {
	publisher *rabbitmq.Publisher
}

func (r *rabbitmqCdr) PublishFileProcessCommand(ctx context.Context, fileMetadataId string) error {
	cdrFileProcessCommand := message.NewCdrFileProcessCommand(fileMetadataId)
	rawMessage, err := json.Marshal(cdrFileProcessCommand)
	if err != nil {
		return err
	}

	correlationId := ctx.Value(CorrelationIdKey).(string)
	publishingMessage := amqp091.Publishing{
		ContentType:   "text/json",
		Body:          rawMessage,
		CorrelationId: correlationId,
	}
	err = r.publisher.Channel.Publish(
		CdrFilesExchange,
		message.CdrFileProcessCommandType,
		false,
		false,
		publishingMessage,
	)
	if err != nil {
		return err
	}
	return nil
}

func NewRabbitmqCdr(publisher *rabbitmq.Publisher) publishers.Cdr {
	return &rabbitmqCdr{
		publisher: publisher,
	}
}
