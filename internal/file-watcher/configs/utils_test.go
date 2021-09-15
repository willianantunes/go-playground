package configs

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldRetrieveDefaultValueAsEnvIsNotSet(t *testing.T) {
	// Arrange
	defaultValue := "IAGO"
	// Act
	value := getEnvValueOrUseDefault("JAFAR_PARTNER", defaultValue)
	// Assert
	assert.Equal(t, defaultValue, value)
}

func TestShouldPanicGivenMustHaveEnvIsNotAvailable(t *testing.T) {
	// Arrange
	envName := "I_AM_NOT_HERE"
	// Assert
	expectedMessage := fmt.Sprintf("Must have variable is not set: %s", envName)
	assert.PanicsWithValue(t, expectedMessage, func() {
		getEnvValueOrPanic(envName)
	})

}
