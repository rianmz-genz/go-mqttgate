package test

import (
	"adriandidimqttgate/app"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDB(t *testing.T) {
	db := app.NewDBConnection()
	assert.NotNil(t, db.DB)
}
