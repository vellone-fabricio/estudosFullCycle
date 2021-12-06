package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreditCardNumber(t *testing.T) {
	_, err := NewCreditCard("248928482000000", "Jose Delfinos", 12, 2024, 123)
	assert.Equal(t, "invalid credit card number", err.Error())

	_, err = NewCreditCard("5520885788779456", "Jose Delfinos", 12, 2024, 123)
	assert.Nil(t, err)
}

func TestCreditCardExpirationMonth(t *testing.T) {
	_, err := NewCreditCard("5520885788779456", "Jose Delfinos", 13, 2024, 123)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("5520885788779456", "Jose Delfinos", 0, 2024, 123)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("5520885788779456", "Jose Delfinos", 5, 2024, 123)
	assert.Nil(t, err)
}

func TestCredtiCardExpirationYear(t *testing.T) {
	lastYear := time.Now().AddDate(-1, 0, 0)
	_, err := NewCreditCard("5520885788779456", "Jose Delfinos", 6, lastYear.Year(), 123)
	assert.Equal(t, "invalid expiration year", err.Error())
}
