package main

import (
	"database/sql"
	"encoding/json"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/mattn/go-sqlite3"
	"github.com/vellone-fabricio/imersao5-gateway/adapter/broker/kafka"
	"github.com/vellone-fabricio/imersao5-gateway/adapter/factory"
	"github.com/vellone-fabricio/imersao5-gateway/adapter/presenter/transaction"
	"github.com/vellone-fabricio/imersao5-gateway/usecase/process_transaction"
)

func main() {
	// db
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	// repository
	repositoryFactory := factory.NewRepositoryDatabaseFactory(db)
	repository := repositoryFactory.CreateTransactionRepository()

	// configMapProducer
	configMapProducer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
	}

	// Producer
	kafkaPresenter := transaction.NewTransactionKafkaPresenter()
	producer := kafka.NewKafkaProducer(configMapProducer, kafkaPresenter)

	var msgChan = make(chan *ckafka.Message)
	//Configmap Consumer
	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"client.id":         "goapp",
		"group.id":          "goapp",
	}
	//topic
	topics := []string{"transactions"}

	//consumer
	consumer := kafka.NewConsumer(configMapConsumer, topics)
	go consumer.Consume(msgChan)

	//usecase
	usecase := process_transaction.NewProcessTransaction(repository, producer, "transaction_result")

	for msg := range msgChan {
		var input process_transaction.TransactionDtoInput
		json.Unmarshal(msg.Value, &input)

		usecase.Execute(input)
	}
}
