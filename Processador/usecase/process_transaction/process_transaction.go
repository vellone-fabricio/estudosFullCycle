package process_transaction

import (
	"fmt"

	"github.com/vellone-fabricio/imersao5-gateway/adapter/broker"
	"github.com/vellone-fabricio/imersao5-gateway/domain/entity"
	"github.com/vellone-fabricio/imersao5-gateway/domain/repository"
)

type ProcessTransaction struct {
	Repository repository.TransactionRepository
	Producer   broker.ProducerInterface
	Topic      string
}

func NewProcessTransaction(repository repository.TransactionRepository, producerInterface broker.ProducerInterface, topic string) *ProcessTransaction {
	return &ProcessTransaction{Repository: repository, Producer: producerInterface, Topic: topic}
}

func (p *ProcessTransaction) Execute(input TransactionDtoInput) (TransactionDtoOutput, error) {
	transaction := entity.NewTransaction()
	transaction.ID = input.ID
	transaction.AccountID = input.AccountID
	transaction.Amount = input.Amount

	cc, invalidCC := entity.NewCreditCard(input.CreditCardNumber, input.CreditCardName, input.CreditCardExpirationMonth, input.CreditCardExpirationYear, input.CreditCardCVV)
	fmt.Println(invalidCC)

	if invalidCC != nil {
		return p.rejectTransaction(transaction, invalidCC)
	}

	transaction.SetCreditCard(*cc)
	invalidTransaction := transaction.IsValid()
	if invalidTransaction != nil {
		return p.rejectTransaction(transaction, invalidTransaction)
	}

	return p.approveTransaction(transaction, input)
}

func (p *ProcessTransaction) approveTransaction(transaction *entity.Transaction, input TransactionDtoInput) (TransactionDtoOutput, error) {
	err := p.Repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount, entity.APPROVED, "")
	if err != nil {
		return TransactionDtoOutput{}, nil
	}
	output := TransactionDtoOutput{
		ID:           transaction.ID,
		Status:       entity.APPROVED,
		ErrorMessage: "",
	}

	err = p.publish(output, []byte(transaction.ID))
	if err != nil {
		return TransactionDtoOutput{}, err
	}

	return output, nil
}

func (p *ProcessTransaction) rejectTransaction(transaction *entity.Transaction, invalidTransaction error) (TransactionDtoOutput, error) {
	err := p.Repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount, entity.REJECTED, invalidTransaction.Error())
	if err != nil {
		return TransactionDtoOutput{}, nil
	}
	output := TransactionDtoOutput{
		ID:           transaction.ID,
		Status:       entity.REJECTED,
		ErrorMessage: invalidTransaction.Error(),
	}

	err = p.publish(output, []byte(transaction.ID))
	if err != nil {
		return TransactionDtoOutput{}, err
	}

	return output, nil
}

func (p *ProcessTransaction) publish(output TransactionDtoOutput, key []byte) error {
	err := p.Producer.Publish(output, key, p.Topic)
	if err != nil {
		return err
	}

	return nil
}
