package greeter_test

import (
	"go-testing/greeter"
	"go-testing/greeter/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreetWithMockReader(t *testing.T) {

	mockReader := &mocks.Reader{}
	mockReader.On("GetMessage", "en").Return("hello", nil).Once()

	testService := greeter.NewGreeterService(mockReader)
	msg := testService.Greet("en")

	expectedMsg := "hello"

	assert.Equal(t, expectedMsg, msg)
	mockReader.AssertNumberOfCalls(t, "GetMessage", 1) // we can assert how many times the mocked method will be called
}
