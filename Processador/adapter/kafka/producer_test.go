package kafka

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vellone-fabricio/imersao5-gateway/adapter/presenter/transaction"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/vellone-fabricio/imersao5-gateway/domain/entity"
	"github.com/vellone-fabricio/imersao5-gateway/usecase/process_transaction"
)

func TestProducerPublish(t *testing.T) {
	expectedOutput := process_transaction.TransactionDtoOutput{
		ID:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "you dont have limit for this transaction",
	}
	// outputJson, _ := json.Marshal(expectedOutput)

	configMap := ckafka.ConfigMap{
		"test.mock.num.brokers": 3,
	}
	producer := NewKafkaProducer(&configMap, transaction.NewTransactionKafkaPresenter())
	err := producer.Publish(expectedOutput, []byte("1"), "test")
	assert.Nil(t, err)
}
