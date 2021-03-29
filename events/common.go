package events

import envvar "github.com/hieronimusbudi/go-fiber-bookstore-order-api/env"

var (
	KafkaURL      = envvar.KafkaURL
	KafkaURLLocal = "localhost:9092"

	TopicItemCreated      = envvar.KafkaTopic
	TopicItemCreatedLocal = "test-item-created"

	TopicItemUpdated      = envvar.KafkaTopic
	TopicItemUpdatedLocal = "test-item-updated"

	TopicItemDeleted      = envvar.KafkaTopic
	TopicItemDeletedLocal = "test-item-deleted"

	GroupIDItemCreated = "test-item-created-group"
	GroupIDItemUpdated = "test-item-updated-group"
	GroupIDItemDeleted = "test-item-deleted-group"
)
