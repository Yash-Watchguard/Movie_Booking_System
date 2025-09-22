package utills

import (
	"testing"
)

// unit test for validator function

func TestCheckPhoneNumber(t *testing.T) {
	tests := []struct {
		name    string
		number  string
		wantErr bool
	}{
		{"Valid number starting with 6", "6123456789", false},
		{"Valid number starting with 7", "7123456789", false},
		{"Valid number starting with 8", "8123456789", false},
		{"Valid number starting with 9", "9123456789", false},
		{"Invalid number less than 10 digits", "912345678", true},
		{"Invalid number more than 10 digits", "91234567890", true},
		{"Invalid number starting with 5", "5123456789", true},
		{"Invalid number starting with 0", "0123456789", true},
		{"Invalid number with letters", "91234abcde", true},
		{"Empty string", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckPhoneNumber(tt.number)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckPhoneNumber() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckEmail(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		wantErr bool
	}{
		{"Valid email lowercase", "test@example.com", false},
		{"Valid email uppercase", "TEST@EXAMPLE.COM", false},
		{"Valid email mixed case", "Test.Email+alias@example.co.uk", false},
		{"Invalid email missing @", "testexample.com", true},
		{"Invalid email missing domain", "test@", true},
		{"Invalid email missing username", "@example.com", true},
		{"Invalid email with spaces", "test @example.com", true},
		{"Invalid email with special chars", "test!@example.com", true},
		{"Empty string", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckEmail(tt.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
