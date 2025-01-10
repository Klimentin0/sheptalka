package auth_test

import (
	"fmt"
	"testing"

	"github.com/Klimentin0/sheptalka/internal/auth"
)

// TestHashPassword проверяет корректность хеширования пароля.
func TestHashPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{
			name:     "Valid password",
			password: "password123",
			wantErr:  false,
		},
		{
			name:     "Empty password",
			password: "",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHashedPassword, err := auth.HashPassword(tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotHashedPassword == "" && !tt.wantErr {
				t.Error("got empty hashed password but no error")
			}

			fmt.Printf("Original Password: %s\nHashed Password: %s\n", tt.password, gotHashedPassword)
		})
	}
}
