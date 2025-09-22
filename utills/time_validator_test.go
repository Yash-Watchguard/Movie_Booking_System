package utills

import (
	"testing"
	"time"
)


// unit test for validatetime
func TestValidateTimeValidRFC3339(t *testing.T) {
	validTimeString := "2023-10-01T12:00:00Z"
	expectedTime, _ := time.Parse(time.RFC3339, validTimeString)

	parsedTime, ok := ValidateTime(validTimeString)

	if !ok {
		t.Errorf("Expected validation to succeed, but it failed")
	}

	if !parsedTime.Equal(expectedTime) {
		t.Errorf("Expected time %v, got %v", expectedTime, parsedTime)
	}
}

func TestValidateTimeInvalidFormat(t *testing.T) {
	invalidTimeString := "invalid-time-string"

	parsedTime, ok := ValidateTime(invalidTimeString)

	if ok {
		t.Errorf("Expected validation to fail, but it succeeded")
	}

	now := time.Now()
	if parsedTime.Sub(now) > time.Minute || now.Sub(parsedTime) > time.Minute {
		t.Errorf("Expected time close to now on failure, got %v", parsedTime)
	}
}

func TestValidateTimeEmptyString(t *testing.T) {
	emptyString := ""

	_, ok := ValidateTime(emptyString)

	if ok {
		t.Errorf("Expected validation to fail for empty string, but it succeeded")
	}
}

func TestValidateTimeWrongFormat(t *testing.T) {
	wrongFormat := "2023-10-01 12:00:00" 

	_, ok := ValidateTime(wrongFormat)

	if ok {
		t.Errorf("Expected validation to fail for wrong format, but it succeeded")
	}
}
