package test

import (
	"adriandidimqttgate/model"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestNewDB(t *testing.T) {
	model.OpenConnection()
	assert.NotNil(t, model.DB)
}