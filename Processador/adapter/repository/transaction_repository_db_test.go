package repository

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vellone-fabricio/imersao5-gateway/adapter/repository/fixture"
	"github.com/vellone-fabricio/imersao5-gateway/domain/entity"
)

func TestTransactionRepositoryDbInsert(t *testing.T) {
	migrationsDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationsDir)
	defer fixture.Down(db, migrationsDir)

	repository := NewTransactionRepositoryDb(db)
	err := repository.Insert("1", "1", 12.1, entity.APPROVED, "")
	assert.Nil(t, err)

}
