package events

import (
	"encoding/json"
	"log"

	eventsutils "github.com/hieronimusbudi/go-bookstore-utils/events"
	"github.com/hieronimusbudi/go-fiber-bookstore-order-api/domain/items"
	"github.com/hieronimusbudi/go-fiber-bookstore-order-api/events"
	"github.com/segmentio/kafka-go"
)

func itemDeletedEventMessageHandler(message *kafka.Message) error {
	responseMessage := new(eventsutils.Message)
	parsingErr := json.Unmarshal([]byte(string(message.Value)), responseMessage)
	if parsingErr != nil {
		return parsingErr
	}

	item := new(items.Item)
	contextInByte, mErr := json.Marshal(responseMessage.Context)
	if mErr != nil {
		return mErr
	}
	parsingContextErr := json.Unmarshal(contextInByte, item)
	if parsingContextErr != nil {
		return parsingContextErr
	}

	if err := item.Delete(); err != nil {
		return err
	}
	// log.Printf("message at topic:%v partition:%v offset:%v	%s = %s | %s | %+v\n", message.Topic, message.Partition, message.Offset, string(message.Key), string(message.Value), responseMessage.Subject, responseMessage.Context)
	log.Printf("topic: %s - message: %+v\n", message.Topic, item)
	return nil
}

func ConsumeItemDeletedEvent() {

	eventsutils.RunConsumer(
		itemDeletedEventMessageHandler,
		eventsutils.ConsumerConfig{
			Brokers: events.KafkaURLLocal,
			Topic:   events.TopicItemDeletedLocal,
			GroupID: events.GroupIDItemDeleted,
		},
	)
}
