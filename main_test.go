package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassword_Validate(t *testing.T) {
	p := Password("1231")
	err := p.Validate()

	assert.Error(t, err)
}
