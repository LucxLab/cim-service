package publisher

import (
	"context"
	"encoding/json"
	"github.com/LucxLab/cim-service/internal/cdr"
	"github.com/LucxLab/cim-service/internal/constant"
	"github.com/LucxLab/cim-service/internal/rabbitmq"
	"github.com/LucxLab/cim-service/internal/rabbitmq/message"
	"github.com/rabbitmq/amqp091-go"
)

const CdrFilesExchange = "cdr_files"

type rabbitmqCdr struct {
	publisher *rabbitmq.Publisher
}

func (r *rabbitmqCdr) PublishCdrFileUploaded(ctx context.Context, fileMetadataId string) error {
	cdrFileUploadedMessage := message.ToCdrFileUploadedMessage(fileMetadataId)
	rawMessage, err := json.Marshal(cdrFileUploadedMessage)
	if err != nil {
		return err
	}

	correlationId := ctx.Value(constant.CorrelationIdContextKey).(string)
	publishingMessage := amqp091.Publishing{
		ContentType:   "text/json",
		Body:          rawMessage,
		CorrelationId: correlationId,
	}
	err = r.publisher.Channel.Publish(
		CdrFilesExchange,
		message.CdrFileUploadedType,
		false,
		false,
		publishingMessage,
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
