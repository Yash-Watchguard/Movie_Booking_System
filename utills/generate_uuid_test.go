package utills

import (
	"testing"

	"github.com/google/uuid"
)

func TestGenerateUuid(t *testing.T) {
	uuidString := GenerateUuid()

	if uuidString == "" {
		t.Errorf("Expected non-empty UUID string, got empty")
	}
	parsedUUID, err := uuid.Parse(uuidString)
	if err != nil {
		t.Errorf("Expected valid UUID, got error: %v", err)
	}
	if parsedUUID == uuid.Nil {
		t.Errorf("Expected non-nil UUID, got nil")
	}
}

