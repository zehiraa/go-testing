package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var (
	a = 25.0
	b = 4.0
	c = 0.0
	expectedSumResult = 29.0
	expectedDivisionResult = 6.25
)

func TestSumWithTestify(t *testing.T){
	assert.Equal(t, Sum(a, b), expectedSumResult)
}

func TestDivideWithTestify(t *testing.T){
	t.Log("divide with non zero decimal")
	division, _ := Divide(a, b)
	assert.Equal(t, division, expectedDivisionResult)

	t.Log("divide with zero decimal")
	divide, err := Divide(a,c)
	assert.NotNil(t, err)
	assert.Equal(t, divide, 0.0)
}